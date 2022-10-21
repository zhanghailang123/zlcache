package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

//keys : 真实节点地址的集合
func (m *Map) Add(keys ...string) {
	//循环添加key
	for _, key := range keys {
		//虚拟节点数量 每个节点都得添加这么多
		for i := 0; i < m.replicas; i++ {
			//i + 节点地址 调用hash函数 生成hash码
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			//keys + hash节点的地址 往虚拟节点数组集合里面添加刚刚生成的hash码（）
			m.keys = append(m.keys, hash)
			//虚拟节点与真实节点的映射表
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	//调用同一hash 函数对key计算hash
	hash := int(m.hash([]byte(key)))

	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
