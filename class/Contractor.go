package class

import (
	"fmt"
	"net/rpc"
)

type Contractor struct {
}

func (contractor Contractor) QueryAllWorkers() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("getAllWorkers rpc.Dial err:", err)
		return
	}
	defer conn.Close()

	//远程调用ParkServer的方法
	var ret WareHouse
	err = conn.Call("Service.QueryAllWorkers", 1, &ret)
	if err != nil {
		fmt.Println("QueryAllWorkers conn.Call err:", err)
		return
	}
	fmt.Println(ret)

	for {

	}
}
