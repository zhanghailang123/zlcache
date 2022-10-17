package lru

import "testing"

type String string

func (d String) Len() int {
	return len(d)
}
func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))

	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit 1234 failed")
	}

	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache hit failed")
	}
}
