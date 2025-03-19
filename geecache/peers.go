package geecache

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 相等于HTTP客户端
type PeerGetter interface {
	Get(group string, key string) ([]byte, error) // 从group中查找缓存
}
