package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println(fetch("http://www.gopl.io/")) // index.html 4154 <nil>
}

func fetch(url string) (local string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	local = path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return
	}
	// defer f.Close() // 姜还是老的辣，Windows 的 NTFS 文件系统，如果写出错了可能在文件关闭时才返回，所以手动 close

	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); closeErr != nil {
		return
	}

	return local, n, err
}
