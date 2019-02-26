package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

var v = flag.Bool("v", false, "-v: show more detail")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	now := time.Now()
	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSize)
			close(fileSize)
		}
	}()

	var fileCount, totalSize int64
	ticker := make(<-chan time.Time) // ticker 本质上还是一个时间值接收 channel
	if *v {
		ticker = time.Tick(200 * time.Millisecond)
	}

loop:
	for {
		select {
		case size, ok := <-fileSize: // 在 channel 上不进行 range 后必须判断当前 channel 是否已关闭。否则可能 panic
			if !ok {
				break loop // 需要使用 label 跳出 select 和 for
			}
			fileCount++
			totalSize += size
		case <-ticker:
			printUshage(fileCount, totalSize)
		}
	}

	printUshage(fileCount, totalSize) // ok
	fmt.Printf("cost: %dms", time.Now().Sub(now).Nanoseconds()/int64(time.Millisecond))
}

func printUshage(files, size int64) {
	fmt.Printf("%d files: %.2fGB\n", files, float64(size)/1e9)
}

func walkDir(dir string, fileSize chan<- int64) {
	for _, f := range read(dir) {
		if f.IsDir() {
			subDir := path.Join(dir, f.Name())
			walkDir(subDir, fileSize)
		} else {
			fileSize <- f.Size()
		}
	}
}

func read(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return files
}
