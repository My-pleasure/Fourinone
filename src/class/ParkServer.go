package ParkServer

import (
	"fmt"
	"net"
)

type waitingWorkers struct {
	workerIP   string
	workerPort string
	Ready      bool
}

type Park struct {
}

//监听MigrantWorker的注册
func logInServer() {
	//创建监听socket;固定IP+端口为"127.0.0.1:8000"
	listener, err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("logInServer net.Listen err:",err)
		return
	}
	defer listener.Close()

	//循环监听注册
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("logInServer listener.Accept err:",err)
			return
		}
		go HandleConn(conn)
	}
}
