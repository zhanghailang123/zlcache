package main

import (
	"ZCache/zcache"
	"container/list"
	"flag"
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

//var db = map[string]string{
//	"Tom":  "630",
//	"Jack": "631",
//	"Sam":  "567",
//}
//
//func main2() {
//	zcache.NewGroup("scores", 2<<10, zcache.GetterFunc(
//		func(key string) ([]byte, error) {
//			log.Println("[SlowDB] search key", key)
//
//			if v, ok := db[key]; ok {
//				return []byte(v), nil
//			}
//
//			return nil, fmt.Errorf("%s not exists", key)
//		}))
//	addr := "localhost:9999"
//	peers := zcache.NewHTTPPool(addr)
//	log.Println("zcache is running at", addr)
//	log.Fatal(http.ListenAndServe(addr, peers))
//}
//
//func main() {
//	str := new(string)
//	*str = "GO语言教程"
//	fmt.Printf("输出结果：%s\n", *str)
//
//	str1 := "zzz"
//	fmt.Println(str1)
//}

//10.23 unit test

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func createGroup() *zcache.Group {
	return zcache.NewGroup("scores", 2<<10, zcache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exists", key)
	}))
}

//启动缓存服务器 创建HTTPPool 添加节点信息 注册到zcache中，
func startCacheServer(addr string, addrs []string, z *zcache.Group) {
	peers := zcache.NewHTTPPool(addr)
	peers.Set(addrs...)
	z.RegisterPeers(peers)
	log.Println("zcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

//func startApiServer(apiAddr string, group *zcache.Group) {
//	http.Handle("/api", http.HandlerFunc(
//		func(w http.ResponseWriter, r *http.Request) {
//		key := r.URL.Query().Get("key")
//		view, err := group.Get(key)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.Header().Set("Content-Type", "application/octet-stream")
//		w.Write(view.ByteSlice())
//
//	}))
//}

func startAPIServer(apiAddr string, gee *zcache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())

		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))

}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "zcache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	zcache := createGroup()
	if api {
		go startAPIServer(apiAddr, zcache)
	}
	startCacheServer(addrMap[port], []string(addrs), zcache)
}
