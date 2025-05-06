package main

import (
	"fmt"
	"net"
	"sync"
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

// 监听广播信息channel（Message）的goroutine，有消息就发送给全部在线User
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
	user := NewUser(conn)
	//用户上线，将用户加入map
	this.mapLock.Lock() //加写锁，加读锁是RLock
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock() //解写锁，解读锁是RUnlock

	//广播该用户上线信息
	this.BroadCast(user, "已上线")

	//阻塞
	select {}
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

	//启动监听Message的goroutine
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
