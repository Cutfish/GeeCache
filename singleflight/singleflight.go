package singleflight

import "sync"

// Call代表正在进行中或者已经结束的请求
// 并发协程之间不需要消息传递，非常适合 sync.WaitGroup
type Call struct {
	wg  sync.WaitGroup //避免重入
	val interface{}
	err error
}

// 最重要的数据结构，管理不同key的请求
type Group struct {
	mu sync.Mutex // 保护m
	m  map[string]*Call
}

// 针对相同的 key，无论 Do 被调用多少次，函数 fn 都只会被调用一次，
// 等待 fn 调用结束了，返回返回值或错误。
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*Call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()         // 请求在进行中，就等待
		return c.val, c.err //请求结束就返回
	}

	c := new(Call)
	c.wg.Add(1)  // 发起请求前加锁
	g.m[key] = c // 添加到 g.m，表明 key 已经有对应的请求在处理
	g.mu.Unlock()

	c.val, c.err = fn() // 调用 fn，发起请求
	c.wg.Done()         // 请求结束

	g.mu.Lock() // 更新 g.m
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err // 返回结果
}
