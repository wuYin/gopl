package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get error: %v\n", err)
			os.Exit(1)
		}
		// 无需调用 ioutil.ReadAll 将响应数据流全部读入缓冲区，直接输出
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy error: %v\n", err)
			os.Exit(1)
		}
	}
}
