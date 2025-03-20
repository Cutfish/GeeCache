package geecache

import (
	pb "GeeCache/geecachepb/geecachepb"
)

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}

// 相等于HTTP客户端
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool) // 从group中查找缓存
}
