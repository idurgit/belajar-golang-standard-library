package main

import "fmt"
import "container/list"

func main() {
	var data *list.List = list.New()
	data.PushBack("Eko")
	data.PushBack("Jaka")
	data.PushBack("Kelana")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) 

	next := head.Next()
	fmt.Println(next.Value) 

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) 
	}
}