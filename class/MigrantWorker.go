package class

import (
	"fmt"
	"net"
	"os"
)

type Workers struct {
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
		fmt.Println("注册成功，等待打黑工中。。。")
	}
}

//开启RPC服务
func startRPC() {

}

//农民工启动
func (workers Workers) StartWork() {
	//获取命令行参数
	list := os.Args
	if len(list) != 3 {
		fmt.Println("我的天呐！大兄弟，你的格式错误了(⊙ˍ⊙)")
		fmt.Println("正确格式为：go run xxx.go IP Port")
		return
	}
	ip := list[1]
	port := list[2]

	//向职介者注册
	go logInToPark(ip, port)

	//开启RPC服务
	go startRPC()

	//农民工头发-1-1-1...
	for {

	}
}
