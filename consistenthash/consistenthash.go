package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// bytes转换为uint32
type Hash func(data []byte) uint32 //只要是长这个样子的函数都行

// 包含所有的hashed keys
type Map struct {
	hash     Hash
	replicas int            // 虚拟节点倍数 一个真实节点映射为replicas个虚拟节点
	keys     []int          // 排序（哈希环）
	hashMap  map[int]string // 键是虚拟节点的哈希值，值是真实节点的名称
}

// 构造函数
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE // 默认的哈希函数
	}
	return m
}

// 添加缓存节点到hash里面
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key))) // 计算hash值，编号+key
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// 选择节点
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))

	// 二分查找合适的replica
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx%len(m.keys)]] // 如果 idx == len(m.keys)，说明应选择 m.keys[0]
}
