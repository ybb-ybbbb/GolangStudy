package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// 监听当前user channel，有信息就发送到客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线
func (this *User) Online() {
	//用户上线，将用户加入map
	this.server.mapLock.Lock() //加写锁，加读锁是RLock
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock() //解写锁，解读锁是RUnlock

	//广播该用户上线信息
	this.server.BroadCast(this, "已上线")
}

// 用户下线
func (this *User) Offline() {
	//用户上线，将用户加入map
	this.server.mapLock.Lock() //加写锁，加读锁是RLock
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock() //解写锁，解读锁是RUnlock

	//广播该用户上线信息
	this.server.BroadCast(this, "下线")
}

// 发送消息给用户所在客户端
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 处理用户消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		for _, user := range this.server.OnlineMap {
			OnlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(OnlineMsg)
		}
		//更改用户名
	} else if len(msg) > 7 && msg[:7] == "rename|" {

		newname := strings.Split(msg, "|")[1]
		_, ok := this.server.OnlineMap[newname]
		if ok {
			this.SendMsg("该用户名已经被占用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newname] = this
			this.server.mapLock.Unlock()
			this.Name = newname
			this.SendMsg("您已更新用户名：" + this.Name + "\n")
		}

	} else {
		//广播消息（群聊）
		this.server.BroadCast(this, msg)
	}

}

// 创建一个User接口
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启用监听当前user channel的goroutine
	go user.ListenMessage()

	return user

}
