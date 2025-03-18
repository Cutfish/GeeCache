package geecache

// 表示缓存数值
type ByteView struct {
	b []byte
}

// 返回View的长度
func (v ByteView) Len() int {
	return len(v.b)
}

// 返回复制数据byte
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// b是只读的，防止缓存值被修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// 将data返回为string
func (v ByteView) String() string {
	return string(v.b)
}
