package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()

	l.PushBack("zhl first")
	l.PushBack("zhangsan a")
	l.PushFront(67)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	map1 := make(map[string]string)
	map1["zhl"] = "zy"
	map1["hai"] = "love"

	for k, v := range map1 {
		fmt.Printf("key=%s, value=%s\n", k, v)
	}

}
