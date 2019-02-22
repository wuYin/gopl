package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var (
	oneTitle = `
<!DOCTYPE html>
<html lang="zh-cn">
 <head>
  <title> See</title>
 </head>
 <body>
 </body>
</html>
`
	twoTitle = `
<!DOCTYPE html>
<html lang="zh-cn">
 <head>
  <title> See</title>
  <title> See You Again</title>
 </head>
 <body>
 </body>
</html>
`
)

func main() {
	for _, s := range []string{oneTitle, twoTitle,} {
		r := strings.NewReader(s)
		doc, err := html.Parse(r)
		if err != nil {
			return
		}
		fmt.Println(soleTitle(doc)) //  See should panic: has many titles
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch res := recover(); res {
		case nil:
			return
		case bailout{}: // 预料中的 panic，当做错误处理。一般可能会有资源泄漏的风险等等
			err = fmt.Errorf("should panic: has many titles")
		default:
			panic(res) // 预料之外的错误，不予处理
		}
	}()

	var foreachNode func(node *html.Node, pre, post func(node *html.Node))
	foreachNode = func(node *html.Node, pre, post func(node *html.Node)) {
		if pre != nil {
			pre(node)
		}

		for cur := node.FirstChild; cur != nil; cur = cur.NextSibling {
			foreachNode(cur, pre, post) // 提前 var 申明
		}

		if post != nil {
			post(node)
		}
	}
	foreachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			if n.FirstChild != nil {
				if title != "" {
					// 有多个 title 不正常，但处理了
					panic(bailout{})
				}
				title = n.FirstChild.Data
			}
		}
	}, nil)

	return title, nil
}
