package zcache

import (
	"reflect"
	"testing"
)

/**
ww
*/
func TestGetterFunc_Get(t *testing.T) {
	//将一个匿名回调函数转换成接口
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")

	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("callback failed")
	}
}
