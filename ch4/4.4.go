// 实现一次遍历即可完成元素旋转
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	rotate(nums)
	fmt.Println(nums)
}

func rotate(nums []int) {
	head := nums[0]
	copy(nums, nums[1:])
	nums[len(nums)-1] = head
}
