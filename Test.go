package main

import "fmt"

func main() {
	var m []interface{}
	m = append(m, "123")
	m = append(m, "341")
	fmt.Println(m[1])
}
