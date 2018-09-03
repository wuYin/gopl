package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopl/ch4/export"
)

func main() {
	// 4.1 数组
	a1 := [...]int{9: -1} // 未指定索引位置的全为零值
	fmt.Println(a1)
	a2 := [10]int{8: 0, 9: -1}
	fmt.Println(a1 == a2) // true	// 同类型的数组，元素类型可比较且元素值完全相同，则相等

	// 4.2 slice
	// slice 无法直接比较的 2 个原因
	// slice 中元素不如数组直接比如自身包含自身等
	// 底层数组变化时，slice 在不同时期拥有的元素可能不一致
	b1 := []byte("a")
	b2 := []byte("a")
	fmt.Println(bytes.Equal(b1, b2))                       // true
	fmt.Println(stringEqual([]string{"a"}, []string{"a"})) // 比较函数自己写

	// 检验 slice 为空：len(s) == 0, 而非 s == nil
	ages := make([]int, 2)[2:]
	fmt.Println(len(ages), cap(ages), ages == nil) // 0 0 false

	// slice 中移除 Pike
	names := []string{"robe", "pike", "ken", "bear"}
	i := 1
	copy(names[i:], names[i+1:]) // 往前挪
	names = names[:len(names)-1] // 裁剪最后一个位置
	fmt.Println(names)

	// 4.3 map
	// 无法比较值的类型，需要作为 map 的键值时，可使用转换函数
	m1 := make(map[string]int)
	m1[conv(names)] = 1 // slice 作为 key
	fmt.Println(m1)     // map[[robe ken bear]:1]

	// 4.4 struct
	// getPerson(1).Name = "Robe" // not ok // cannot assign to getPerson(1).Name	// 左侧不是变量，无法赋值
	p1 := getPerson(1) // ok
	p1.Name = "Robe"

	getPersonPtr(1).Name = "Robe" // ok

	p2 := Person{"Robe", 20} // 结构体成员均可比较，则结构体可比较，可作为 map 的键值
	fmt.Println(p1 == p2)    // true

	// 嵌套后的结构体，内部成员的可见性是由成员决定
	export.Demo()

	paint := Paint{Year: 1997}
	b, _ := json.Marshal(paint)
	fmt.Println(string(b)) // {"released":1997}

}

type Paint struct {
	Year  int  `json:"released"`
	Color bool `json:"color,omitempty"` // 成员的值是零值或空值，则 Marshal 时不输出到 JSON 字串中
}

var ps = []Person{{"Pike", 20}, {"Ken", 20}}

func getPerson(id int) Person {
	return ps[id]
}

func getPersonPtr(id int) *Person {
	return &ps[id]
}

type Person struct {
	Name string
	Age  int
}

// 将 slice 转换为 string，间接作为 map 的 key
func conv(s []string) (key string) {
	key = fmt.Sprintf("%s", s)
	return
}

// 比较字符串 slice
func stringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
