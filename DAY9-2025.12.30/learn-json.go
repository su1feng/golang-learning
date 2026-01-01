package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Gender string
	ID uint8
}

type Class struct {
	Title string
	Students []Student
}


func main(){
	//json序列化
	c := Class{
		Title: "101",
		Students: make([]Student, 0 , 200),
	}

	for i := range 10{
		stu := Student{
			Name: fmt.Sprintf("stu%02d" , i),
			Gender: "男",
			ID: uint8(i),
		}
		c.Students = append(c.Students, stu)
	}

	//序列化
	data , err := json.Marshal(c)

	if err != nil{
		fmt.Println("json Marshal failed")
		return
	}
	fmt.Printf("data = %s\n",data)
	
	fmt.Println()

	//反序列化
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}//必须用&取地址
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
	fmt.Println()

	//结构体标签Tag 由键值对构成 用` `包裹
	//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	dat, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", dat) //json str:{"id":1,"Gender":"男"}

}