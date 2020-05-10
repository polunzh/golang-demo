package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var age int32 = 30
	fmt.Println(unsafe.Sizeof(age))

	var age1 int64 = 30
	fmt.Println(unsafe.Sizeof(age1))

	var miles float32 = 1.23
	fmt.Println(unsafe.Sizeof(miles))

	var miles1 float64 = 1.23
	fmt.Println(unsafe.Sizeof(miles1))

	var name string = "zzhangzhenqiangzhangzhenqiangzhangzhenqiangzhangzhenqianghangzhenqiang"
	fmt.Println(unsafe.Sizeof(name))
}
