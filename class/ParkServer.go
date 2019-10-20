package class

import (
	"fmt"
	"net"
	"net/rpc"
)

//农民工注册信息
type waitingWorkers struct {
	//workerIP   string
	//workerPort string
	//Ready      bool
	workAddr string
}

//民工注册信息	value为1,可打黑工;value为-1,在别的包工头下打黑工;value为0,当前任务完成
var m map[waitingWorkers]int

//职介者信息
type Park struct {
}

//RPC调用
type Service struct{}

//RPC for Contractor
func (s *Service) QueryAllWorkers(a int, ret *WareHouse) error {
	var str string
	for key, value := range m {
		if value == 1 {
			str = "Ready"
		} else {
			str = "Not Ready"
		}
		*ret = key.workAddr + " " + str
	}
	return nil
}

//用于处理注册的GO程
func HandleLogInConn(conn net.Conn) {
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
	fmt.Println("成功收到来自农民工 ", addr, " 的注册信息：", string(buf[:n]))

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
		go HandleLogInConn(conn)
	}
}

//监听包工头雇工服务
func waitingWorkerServer() {
	//注册RPC服务
	s := new(Service)
	rpc.Register(s)

	//创建监听端口
	tcpaddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8001")
	listener, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		fmt.Println("waitingWorkerServer net.Listen err:", err)
		return
	}
	defer listener.Close()

	//循环接收服务
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("waitingWorkerServer listener.Accept err:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

//启动ParkServer
func (park Park) ParkStart() {
	m = make(map[waitingWorkers]int)

	//启动监听农民工注册服务
	go logInServer()

	//启动监听包工头雇工服务
	go waitingWorkerServer()

	//不要停o(*￣▽￣*)ブ
	for {

	}
}
