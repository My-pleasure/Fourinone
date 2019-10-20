package ParkServer

import (
	"fmt"
	"net"
)

//农民工注册信息
type waitingWorkers struct {
	//workerIP   string
	//workerPort string
	//Ready      bool
	workAddr string
}

//职介者信息
type Park struct {
}

var m map[waitingWorkers]int

//用于处理注册的GO程
func HandleConn(conn net.Conn) {
	defer conn.Close()

	//获取农民工的地址
	addr := conn.RemoteAddr()
	fmt.Println("农民工 ", addr, " 连接成功(p≧w≦q)")

	//ParkServer接收注册信息
	var wks waitingWorkers
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("HandleConn conn.Read err:", err)
		return
	}
	wks.workAddr = string(buf[:n])
	m[wks] = 1
	fmt.Println("成功收到来自农民工 ", addr, " 的注册信息：", buf[:n])

	//回写表示接收成功
	conn.Write([]byte("ok"))
}

//监听MigrantWorker的注册
func logInServer() {
	//创建监听socket;固定IP+端口为"127.0.0.1:8000"
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("logInServer net.Listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("ParkServer is waiting connection ...")

	//循环监听注册
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("logInServer listener.Accept err:", err)
			return
		}
		go HandleConn(conn)
	}
}

//监听包工头雇工服务
func waitingWorkersServer() {

}

//启动ParkServer
func (park Park) ParkStart() {
	m=make(map[waitingWorkers]int)

	//启动监听农民工注册服务
	go logInServer()

	//启动监听包工头雇工服务
	go waitingWorkersServer()

	//不要停o(*￣▽￣*)ブ
	for {

	}
}
