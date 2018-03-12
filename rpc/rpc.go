package rpc

import (
	"github.com/Darkera524/WinTraceServer/g"
	"net/rpc"
	"net"
	"fmt"
	"net/rpc/jsonrpc"
)

func Start(){
	addr := g.GetConfig().Listen

	server := rpc.NewServer()
	server.Register(new(Trace))
	server.Register(new(Wmi))

	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("listening", addr)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
