package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	w, n := CountWriter(os.Stdout)
	fmt.Fprint(w, "demo string")
	fmt.Println(n)
}

type ByteCounter struct {
	w     io.Writer
	count int64
}

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	c.w.Write(p) // 不管 writer 内部干了什么
	n = len(p)
	c.count += int64(n)
	return
}

func CountWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{w, 0}
	return &c, &c.count
}
