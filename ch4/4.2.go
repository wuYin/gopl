// 命令行标记控制输出散列值
package main

import (
	"flag"
	"crypto/sha512"
	"crypto/sha256"
	"fmt"
)

func main() {
	sha := flag.Int("sha", 256, "sha hash length (256, 384 or 512)")
	s := flag.String("s", "", "string to sha hash")
	flag.Parse()
	var shaFunc func(b []byte) []byte
	switch *sha {
	case 384:
		shaFunc = func(b []byte) []byte {
			s := sha512.Sum384(b)
			return s[:]
		}
	case 512:
		shaFunc = func(b []byte) []byte {
			s := sha512.Sum512(b)
			return s[:]
		}
	default:
		shaFunc = func(b []byte) []byte {
			s := sha256.Sum256(b)
			return s[:]
		}
	}
	fmt.Printf("sha%d(%q) = %x\n", *sha, *s, shaFunc([]byte(*s)))
}
