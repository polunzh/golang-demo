package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	t := l.PushFront(101)
	t = l.InsertAfter(102, t)
	l.InsertAfter(103, t)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
