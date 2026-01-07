package main

import (
	"fmt"
	"reflect"
)

type myint int64

func reflectType(x interface{}) {
	//关于类型分为 type kind(底层类型)
	type_of_x := reflect.TypeOf(x)
	fmt.Printf("type : %v  kind : %v\n", type_of_x.Name(), type_of_x.Kind())
}

func reflectValue(x interface{}) {
	value_of_x := reflect.ValueOf(x)
	k := value_of_x.Kind()

	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Printf("type is Int , value is %d \n", value_of_x.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Printf("type is Uint , value is %d \n", value_of_x.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Printf("type is Float , value is %f \n", value_of_x.Float())
	case reflect.Bool:
		fmt.Printf("type is Bool , value is %v \n", value_of_x.Bool())
	case reflect.String:
		fmt.Printf("type is String , value is %s \n", value_of_x.String())
	case reflect.Ptr:
		fmt.Printf("type is Pointer , value is %v \n", value_of_x.Pointer())
	case reflect.Slice:
		fmt.Printf("type is Slice , length is %d , value is %v \n", value_of_x.Len(), value_of_x.Interface())
	case reflect.Map:
		fmt.Printf("type is Map , length is %d , value is %v \n", value_of_x.Len(), value_of_x.Interface())
	case reflect.Struct:
		fmt.Printf("type is Struct , value is %v \n", value_of_x.Interface())
	case reflect.Array:
		fmt.Printf("type is Array , length is %d , value is %v \n", value_of_x.Len(), value_of_x.Interface())
	case reflect.Interface:
		fmt.Printf("type is Interface , value is %v \n", value_of_x.Interface())
	case reflect.Func:
		fmt.Printf("type is Function , value is %v \n", value_of_x.IsValid())
	default:
		fmt.Printf("unknown type: %v \n", k)
	}
}

func reflectType_test() {
	var a *int
	var b myint
	var c rune

	type person struct {
		name string
		age  int
	}
	var d = person{
		name: "su1feng",
		age:  23,
	}

	reflectType(a)
	reflectType(b)
	reflectType(c)
	reflectType(d)
}

func reflectValue_test() {
	var a int = 100
	var b uint = 200
	var c float64 = 3.14
	var d bool = true
	var e string = "Hello Go"
	var f []int = []int{1, 2, 3}
	var g map[string]int = map[string]int{"a": 1, "b": 2}
	var h *int = &a
	var i [3]int = [3]int{1, 2, 3}
	var j interface{} = "interface value"

	type person struct {
		name string
		age  int
	}
	var k = person{name: "su1feng", age: 23}

	var l func() = func() {}

	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
	reflectValue(d)
	reflectValue(e)
	reflectValue(f)
	reflectValue(g)
	reflectValue(h)
	reflectValue(i)
	reflectValue(j)
	reflectValue(k)
	reflectValue(l)
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	/*
		空接口可以存储任意类型的变量，通过反射可以知道这个空接口保存的数据是什么
		反射就是在运行时动态的获取一个变量的类型信息和值信息。
		任意接口值在反射中都可以理解为reflect.Type  reflect.Value
		通过提供的reflect.TypeOf和reflect.ValueOf函数获取
	*/
	reflectType_test()
	fmt.Println()
	reflectValue_test()
	fmt.Println()

	a := 100
	fmt.Println("before reflectSetValue , a = ", a)
	reflectSetValue(&a)
	fmt.Println("after reflectSetValue , a = ", a)
}
