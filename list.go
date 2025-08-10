package main

import (
	"container/list"
	"fmt"
)

func main(){
	var data *list.List=list.New()

	data.PushBack("Rama")
	data.PushBack("Fajar")
	data.PushBack("Fadhillah")

	var head *list.Element=data.Front()
	fmt.Print(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}