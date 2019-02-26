package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 后台遍历全部目录
	fileSizes := make(chan int64)
	go func() {
		for _, dir := range roots {
			walkDir(dir, fileSizes)
		}
		close(fileSizes)
	}()

	// 主 goroutine 等待结果
	var files, size int64
	for fileSize := range fileSizes {
		files++
		size += fileSize
	}

	fmt.Printf("%d files: %.2fGB\n", files, float64(size)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			// 目录
			subdir := path.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			// 文件
			fileSizes <- entry.Size()
		}
	}
}

// 返回目录下文件信息
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return entries
}
