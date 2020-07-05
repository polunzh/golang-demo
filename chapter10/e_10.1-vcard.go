package main

import "fmt"

type Address struct {
	value string
}

type VCard struct {
	name        string
	address     *Address
	addressCode string
	birthday    string
	avatar      string
}

func main() {
	address := Address{value: "BeiJing"}
	vCard := &VCard{
		name:        "zhangzhenqiang",
		address:     &address,
		addressCode: "10231",
		birthday:    "1989-10-14",
		avatar:      "/avatar/zhang.png",
	}

	fmt.Printf("%v", vCard)
}
