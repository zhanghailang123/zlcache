package main

import (
	"ZCache/zcache"
	"container/list"
	"fmt"
	"log"
	"net/http"
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

//var m sync.Mutex
//
//var set = make(map[int]bool, 0)
//
//func printOnce(num int) {
//	m.Lock()
//	if _, exist := set[num]; exist {
//		fmt.Println(num)
//	}
//	set[num] = true
//	m.Unlock()
//}
//
//func main() {
//	for i := 0; i < 10; i++ {
//		go printOnce(100)
//	}
//	time.Sleep(time.Second)
//}

//type server int
//
//func (s server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	log.Println(request.URL.Path)
//	writer.Write([]byte("Hello world"))
//}
//
//func main() {
//	var s server
//
//	http.ListenAndServe("localhost:8080", &s)
//}
//
//type OpError struct {
//	op string
//}
//
//func (e *OpError) Error() string {
//	return fmt.Sprintf("无权执行%s操作", e.op)
//}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "631",
	"Sam":  "567",
}

func main() {
	zcache.NewGroup("scores", 2<<10, zcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)

			if v, ok := db[key]; ok {
				return []byte(v), nil
			}

			return nil, fmt.Errorf("%s not exists", key)
		}))
	addr := "localhost:9999"
	peers := zcache.NewHTTPPool(addr)
	log.Println("zcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
