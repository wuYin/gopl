package main

import (
	"testing"
	"strings"
)

var args = []string{"The", "quick", "brown", "fox", "jump", "sover", "the", "lazy", "dog"}

func concat(args []string) {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}

//
// > go test -v 1.3_test.go -bench=.
//
// goos: darwin
// goarch: amd64
// pkg: gopl/ch1/ch1.3
// BenchmarkConcat-8        3000000               501 ns/op
// BenchmarkJoin-8         10000000               151 ns/op
// PASS
// ok      gopl/ch1/ch1.3  3.650s
