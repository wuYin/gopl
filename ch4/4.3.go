// 重写函数 reverse，使用数组指针而非 slice
package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}

func reverse(nums *[5]int) {
	l := len(*nums)
	for head := 0; head < l/2; head++ {
		tail := (l - 1) - head
		// 数组和 struct 等聚合结构的元素是可寻址的
		// slice 和 map 等动态结构的元素由于动态增长的特性，元素地址可能变化，地址无意义，不能寻址
		nums[head], nums[tail] = nums[tail], nums[head]
	}
}
