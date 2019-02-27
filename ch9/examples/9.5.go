package main

import "sync"

func Icon(name string) int {
	loadOnce.Do(loadIcons) // 内部的互斥锁保证 loadIcons 只会被执行一次
	return icons[name]
}

var loadOnce sync.Once
var icons map[string]int

// 延迟初始化
func loadIcons() {
	icons = make(map[string]int)
	icons["a"] = 1
	icons["b"] = 2
	icons["c"] = 3
}
