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

	server  *Server
	isClose bool
}

// 监听当前user channel(this.C)，有信息就发送到客户端
func (this *User) ListenMessage() {
	for !this.isClose {
		select {
		case msg := <-this.C:
			this.conn.Write([]byte(msg + "\n"))
		}

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
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	close(this.C) //销毁资源（关闭user的channel）
	this.isClose = true
	//广播该用户上线信息
	this.server.BroadCast(this, "下线")
	//net.Conn中的Close() 方法用于关闭网络连接
	this.conn.Close()
}

// 发送消息给用户所在客户端
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 处理用户消息的业务
func (this *User) DoMessage(msg string) {
	//查询在线用户有谁
	if msg == "who" {
		for _, user := range this.server.OnlineMap {
			OnlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(OnlineMsg)
		}

		//更改用户名，消息格式: rename|张三
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

	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式:  to|张三|消息内容

		//1 获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确，请使用 \"to|张三|你好啊\"格式。\n")
			return
		}

		//2 根据用户名 得到对方User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户名不不存在\n")
			return
		}

		//3 获取消息内容，通过对方的User对象将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，请重发\n")
			return
		}
		remoteUser.SendMsg(this.Name + "对您说:" + content)

	} else {
		//广播消息
		this.server.BroadCast(this, msg)
	}
}

// 创建一个User接口
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		C:       make(chan string),
		conn:    conn,
		server:  server,
		isClose: false,
	}
	//启用监听当前user channel的goroutine
	go user.ListenMessage()

	return user

}
