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
	defer func() {
		// 优先报告 io.Copy 的错误
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
