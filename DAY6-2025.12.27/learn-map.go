package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//map是无序的 key-value数据结构   是引用类型的  必须初始化才能使用
	//map[keytype] valuetype   默认初始值为nil
	//使用make来分配内存  make(map[key type]value type , cap)
	mp1 := make(map[int]int, 8)
	mp2 := map[string]int{
		"su1feng": 23,
		"ustc":    3,
	}
	fmt.Println(mp1)
	mp1[3] = 666
	fmt.Println(mp1)
	fmt.Println(mp2)

	//判断某个键是否存在   存在ok为true value为对应的值;不存在ok为false value为值类型的零值
	//value , ok := map[key]
	value, ok := mp1[3]
	if ok {
		fmt.Println("mp1[3] = ", value)
	} else {
		fmt.Println("mp1[3]不存在")
	}

	//遍历map    遍历顺序与添加顺序无关
	for key, value := range mp2 {
		fmt.Println("key = ", key, "value = ", value)
	}

	for key := range mp2 {
		fmt.Println("key = ", key)
	}

	//delete 删除键值对
	//delete(map , key)
	delete(mp1, 3)
	fmt.Println(mp1)

	fmt.Println("=========================================================")

	//按指定顺序遍历
	var sort_map = make(map[int]int, 100)

	for i := 0; i < 100; i++ {
		key := i
		value := i + 620
		sort_map[key] = value
	}

	//将map中的key存入切片keys 对切片排序 再进行遍历
	keys := make([]int, 100)
	for key := range sort_map {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for key := range keys {
		fmt.Println("key = ", key, "value = ", sort_map[key])
	}

	fmt.Println("=============")

	//元素为map的切片slice
	var mapslice = make([]map[string]string, 3)

	for key, value := range mapslice {
		fmt.Println("key = ", key, "value = ", value)
	}

	mapslice[0] = make(map[string]string, 5)
	mapslice[0]["name"] = "su1feng"
	mapslice[0]["age"] = "23"

	for index, val := range mapslice {
		fmt.Println("index = ", index, "val = ", val)
	}

	fmt.Println("-------------------------------")

	//值为切片slice的map
	var slice_map = make(map[string][]string, 10)
	fmt.Println(slice_map)

	key := "hobby"
	v, ok := slice_map[key]

	if ok {
		fmt.Println("hobby = ", v)
	} else {
		v = make([]string, 0, 3)
	}

	v = append(v, "coding", "music", "games")
	slice_map[key] = v

	fmt.Println(slice_map)

	practice1()
	practice2()
}

func practice1() {
	str := "how do you do think you about"
	//strings.Split 按指定的分隔符将字符串分割成切片
	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)

	countMap := make(map[string]int, 10)
	for _, key := range strSlice {
		_, isReal := countMap[key]
		if !isReal {
			countMap[key] = 1
		} else {
			countMap[key] += 1
		}
	}
	fmt.Println(countMap)
}

func practice2() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
