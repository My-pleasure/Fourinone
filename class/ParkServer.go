package class

import (
	"fmt"
	"net"
	"net/rpc"
)

//农民工注册信息
type WaitingWorkers struct {
	//workerIP   string
	//workerPort string
	//Ready      bool
	WorkAddr string
}

//民工注册信息	value为1,可打黑工;value为-1,在别的包工头下打黑工;value为0,当前任务完成
var m map[WaitingWorkers]int

//职介者信息
type Park struct {
}

//RPC调用
type Service struct{}

//RPC for Contractor
func (s *Service) QueryAllWorkers(a int, ret *WareHouse) error {
	str := ""
	answer := ""
	for key, value := range m {
		if value == 1 {
			str = "Ready"
		} else {
			str = "Not Ready"
		}
		answer += key.WorkAddr + " " + str + "\n"
	}
	*ret = answer
	return nil
}

//用于处理注册的GO程
func HandleLogInConn(conn net.Conn) {
	defer conn.Close()

	//获取农民工的地址
	addr := conn.RemoteAddr()
	fmt.Println("MigrantWorker ", addr, " connect successfully!(p≧w≦q)")

	//ParkServer接收注册信息
	var wks WaitingWorkers
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("HandleConn conn.Read err:", err)
		return
	}
	wks.WorkAddr = string(buf[:n])
	m[wks] = 1
	fmt.Println("success to get the ", addr, "'s log message ：", string(buf[:n]))

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
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("waitingWorkerServer net.ResolveTCPAddr err:", err)
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
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
	m = make(map[WaitingWorkers]int)

	//启动监听农民工注册服务
	go logInServer()

	//启动监听包工头雇工服务
	go waitingWorkerServer()

	//不要停o(*￣▽￣*)ブ
	for {

	}
}
