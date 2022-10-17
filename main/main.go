package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

func main1() {

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

	m := make(map[int]int)

	go func() {
		for true {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
		}
	}()

}

var m sync.Mutex

var set = make(map[int]bool, 0)

func printOnce(num int) {
	m.Lock()
	if _, exist := set[num]; exist {
		fmt.Println(num)
	}
	set[num] = true
	m.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
