package main

import "sync"

var mu sync.Mutex

func main() {
	mu.Lock()
	defer mu.Unlock()
	// 在函数结束之前，互斥锁始终处于锁定状态
}
