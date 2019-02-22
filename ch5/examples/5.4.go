package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Ctrl+D 中断输入，返回特殊的 EOF io 错误
	charCount()
}

func charCount() error {
	in := bufio.NewReader(os.Stdin)
	count := 0
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		count++
		fmt.Printf("%d\t%q\n", count, r)
	}

	return nil
}
