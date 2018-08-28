// dup2 打印输入中多次出现的行的个数和文本
// 它从 stdin 或指定的文件列表读取
// go run dup2 demo.txt
package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// 从 stdin 读取
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		return
	}
	// 从文件列表读取
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open file %s error: %v\n", file, err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}
	for line, n := range counts {
		if n >= 2 {
			fmt.Printf("%q\t%d\n", line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Err() != nil {
			fmt.Fprintf(os.Stderr, "%v\n", input.Err())
			break
		}
		counts[input.Text()]++
	}
}
