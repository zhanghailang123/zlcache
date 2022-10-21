package zcache

type PeerPicker interface {
	PickPeer(key string) (PeerGetter, bool)
}

//对应上述流程中的http客户端
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
