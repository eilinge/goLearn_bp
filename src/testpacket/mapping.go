package main

import (
	"fmt"
	"sort"
)

/*
Go 映射map: make(map[int][string], cap)
映射是一种内置的数据结构，用来保存键值对的无序集合。
map的本质决定了，一旦容量不够，它会自动扩容

类型其他语言中的哈希表或者字典, 一key-value形式存储数据
key必须是支持== 或 != 比较运算的类型, 不可以是函数/map/slice
Map查找比线性搜索快很多, 但比使用索引访问数据的类型慢100倍
Map使用make()创建, 支持 := 简写
len(map)返回长度
*/
func mainMapping() {
	// m := make(map[int]map[int]string)
	// m[1] = make(map[int]string) // map进行操作前, 需要对其初始化: make(map[int]map[int]string)
	// m[1][1] = "OK"
	// m[1] = "OK"
	// delete(m, 1)  // 删除map中的key, value
	// a := m[1][1]

	// a, k := m[2][1] // k 返回该key是否被初始化
	// if !k {
	// 	m[2] = make(map[int]string)
	// }
	// m[2][1] = "GOOD"
	// a, k = m[2][1]
	// fmt.Println(a, k)

	sm := make([]map[int]string, 3)
	// k: index
	// type: map[int]string
	// len: 3
	for _, v := range sm {
		v = make(map[int]string, 2)
		// int=>string: map[int]string
		// cap: 2
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)
	// v 是从sm拷贝过来的, 对v操作不会影响原sm
	// map[1:OK]
	// map[1:OK]
	// map[1:OK]
	// [map[] map[] map[]]
	for k := range sm {
		sm[k] = make(map[int]string, 2)
		// int=>string: map[int]string
		// cap: 2
		sm[k][1] = "OK"
		fmt.Println(sm[k])
	}
	fmt.Println(sm)
	// 进行迭代操作时, 需要对sm直接操作, 才会返回变化后的值
	// map[1:OK]
	// map[1:OK]
	// map[1:OK]
	// [map[1:OK] map[1:OK] map[1:OK]]

	// 实现对Map的排序
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "5"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)   // [1 2 3 4]
	fmt.Println(s) // 集合是无续的[2 3 4 1]

	// map[k]=v => map[v]=k
	m1 := make(map[string]int, len(m))
	for k, v := range m {
		m1[v] = k
	}
	fmt.Println(m1) // map[5:4 a:1 b:2 c:3]
}
