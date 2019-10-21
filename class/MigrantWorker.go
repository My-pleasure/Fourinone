package class

import (
	"fmt"
	"net"
	"os"
)

//农民工
type Workers struct {
}

//RPC服务
type WorkRPC struct {
}

func (workRPC WorkRPC) DoTask(a int, ret *WareHouse) error {
	return nil
}

//向职介者注册
func logInToPark(ip string, port string) {
	//创建TCP连接,连接职介者
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("logInToPark net.Dial() err:", err)
		return
	}
	defer conn.Close()

	//发送注册的RPC服务地址
	conn.Write([]byte(ip + ":" + port))

	//接收来自ParkServer的注册成功信息
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("logInToPark conn.Read err:", err)
		return
	}
	if string(buf[:n]) == "ok" {
		fmt.Println("log in successfully , waiting work .....")
	}
}

//开启RPC服务
/*func startRPC(ip string,port string) {
	//注册RPC服务
	workRPC:=new(WorkRPC)
	rpc.Register(workRPC)

	//创建监听端口
	tcpAddr,err:=net.ResolveTCPAddr("tcp",ip+":"+port)
	if err!=nil{
		fmt.Println("startRPC net.ResolveTCPAddr err:",err)
		return
	}
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	if err!=nil{
		fmt.Println("startRPC net.ListenTCP err:",err)
		return
	}
	defer listener.Close()

	//循环监听服务
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("startRPC listener.Accept err:",err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}*/

//农民工启动
func (workers Workers) StartWork() {
	//获取命令行参数
	list := os.Args
	if len(list) != 3 {
		fmt.Println("Oh! Baby, your form is wrong! (⊙ˍ⊙)")
		fmt.Println("Please input your hostIP and running Port!")
		fmt.Println("The right form is ：go run xxx.go IP Port")
		return
	}
	ip := list[1]
	port := list[2]

	//开启RPC服务
	//go startRPC(ip,port)

	//向职介者注册
	go logInToPark(ip, port)

	//农民工头发-1-1-1...
	for {

	}
}
