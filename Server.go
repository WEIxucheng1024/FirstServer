package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip		string
	Port	int
}

// 传建一个Server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip: ip,
		Port: port,
	}

	return server
}

func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	fmt.Println("连接建立成功")
}

// 启动服务器的接口
func (this *Server) Start() {
	// socker listen
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen ERROR:",err)
	}
	defer listen.Close()

	// accept
	for{
		//do handler
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("listen ccept ERROR :",err)
			continue
		}

		//close listen socket
		go this.Handler(conn)
	}


}
