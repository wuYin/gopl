package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("xx.txt")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)

	v, ok := err.(*os.PathError)
	fmt.Println(v)  // open xx.txt: no such file or directory
	fmt.Println(ok) // true
}
