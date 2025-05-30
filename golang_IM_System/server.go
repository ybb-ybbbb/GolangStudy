package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	Message   chan string
}

// 创建server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 广播信息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听广播信息channel（Message）的goroutine，有消息就广播
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		//将msg发给所有在线User的user channel
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 业务方法
func (this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, this)
	//用户上线
	user.Online()

	//用一个channel来监控用户是否活跃，用channel主要是为了方便后面使用select
	isAlive := make(chan bool)

	//接受用户发送的信息
	go func() {
		buf := make([]byte, 4096)
		for !user.isClose {
			n, err := conn.Read(buf) //Read 方法会阻塞当前 Goroutine，直到从网络连接中读取到数据，或者发生错误，或者连接关闭。
			if err == io.EOF {       //err == io.EOF这是唯一明确表示连接关闭的标准方式。(此时n=0)
				//用户下线
				user.Offline()
				return
			}
			if err != nil {
				// 检查是否是连接关闭相关的错误
				if err.Error() == " use of closed network connection" {
					return // 如果是连接关闭错误，直接返回，不打印错误信息
				}
				fmt.Println("Conn Read err:", err)
				return
			}
			//提取用户发来的信息,n-1是为了减掉"\n"
			msg := string(buf[:n-1])

			//处理用户消息的业务
			user.DoMessage(msg)

			//更新用户活跃状态
			isAlive <- true
		}
	}()

	//超时踢出服务器
	for {
		select {
		case <-isAlive:
			//如果isAlive为true进入该case，但是不进行任何操作，golang中case执行完会直接跳出select，跳出后外层for循环会让select中的case的判断条件再执行以更新time.After
		case <-time.After(time.Second * 300):
			//time.After()的返回值为channel，到时间后会可读，重新执行该方法会重置时间
			user.SendMsg("您长时间未操作已被踢")

			user.Offline()

			return

		}
	}
}

// 启动sever
func (this *Server) Start() {
	//socket listen
	lisenner, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close listen socket
	defer lisenner.Close()

	//启动监听广播信息channel（Message）的goroutine
	go this.ListenMessager()

	for {
		//accept
		conn, err := lisenner.Accept() //当这个方法被调用时，它会阻塞，直到有客户端连接到服务器。
		//在 TCP 连接的情况下，conn 的类型通常是 *net.TCPConn(既一个socket)，它提供了对 TCP 连接的访问和控制。这个连接对象可以用来读写数据，即发送和接收来自客户端的数据。
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		//do headler
		go this.Handler(conn) //因为accept会阻塞等待，所以需要开辟一个协程来做业务

	}
}
