package main

import (
	"fmt"
	"sync"
)

// Visited用于记网页是否被访问过
// 他的方法可以在多个goroutine中并发使用
type Visited struct {
	// mu用于保护visited映射
	mu sync.Mutex // 声明一个互斥锁
	visited map[string]int
}

// VisitLink 会记录本次针对给定网址的访问，然后返回更新之后的连接统计值
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func visitWebsite(downstream chan int, v Visited, url string) {
	for i := 0; i < 10; i++ {
		count := v.VisitLink(url)
		fmt.Printf("Website %s has been visited %d times\n", url, count)
		downstream <- count
	}
	close(downstream)
}

func main() {
	var mu sync.Mutex

	m := make(map[string]int)
	c := make(chan int)
	v := Visited{mu, m}
	go visitWebsite(c, v, "https://example.com")
	go visitWebsite(c, v, "https://github.com")
	for i := range c {
		fmt.Println(i)
	}
}
