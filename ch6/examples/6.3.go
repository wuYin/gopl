package main

import "sync"

// lock 和 m 共同组成临界区，读写都是原子操作，也就是加锁的目的
var cache = struct {
	lock sync.Mutex
	m    map[string]string
}{
	m: make(map[string]string),
}

// good
func get(k string) (v string) {
	cache.lock.Lock()
	v = cache.m[k]
	cache.lock.Unlock()
	return
}
