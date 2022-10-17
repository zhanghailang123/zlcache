package lru

import "container/list"

type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
}

type Value interface {
	Len() int
}
