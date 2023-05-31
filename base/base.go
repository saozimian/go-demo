package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 整型 有符号
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5

	fmt.Printf("int    : %v\n", unsafe.Sizeof(i1))
	fmt.Printf("int8   : %v\n", unsafe.Sizeof(i2))
	fmt.Printf("int16  : %v\n", unsafe.Sizeof(i3))
	fmt.Printf("int32  : %v\n", unsafe.Sizeof(i4))
	fmt.Printf("int64  : %v\n", unsafe.Sizeof(i5))

	fmt.Println("--------------------------------")

	// 整型 无符号
	var u1 uint = 1
	var u2 uint8 = 2
	var u3 uint16 = 3
	var u4 uint32 = 4
	var u5 uint64 = 5

	fmt.Printf("uint    : %v\n", unsafe.Sizeof(u1))
	fmt.Printf("uint8   : %v\n", unsafe.Sizeof(u2))
	fmt.Printf("uint16  : %v\n", unsafe.Sizeof(u3))
	fmt.Printf("uint32  : %v\n", unsafe.Sizeof(u4))
	fmt.Printf("uint64  : %v\n", unsafe.Sizeof(u5))

	// 浮点型  float32 4字节 float64 8字节
	var f1 float32 = 1
	var f2 float64 = 2
	fmt.Printf("float32  : %v\n", unsafe.Sizeof(f1))
	fmt.Printf("float64  : %v\n", unsafe.Sizeof(f2))

	fmt.Println("--------------------------------")

	// 数组
	// 全局
	var arr0 [5]int = [5]int{1, 2, 3}
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4, 5, 6}
	var str = [5]string{3: "hello world", 4: "tom"}
	fmt.Printf("arr0  : %v\n", arr0)
	fmt.Printf("arr1  : %v\n", arr1)
	fmt.Printf("arr2  : %v\n", arr2)
	fmt.Printf("str   : %v\n", str)
	// 局部
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}

	fmt.Printf("a  : %v\n", a)
	fmt.Printf("b  : %v\n", b)
	fmt.Printf("c  : %v\n", c)
	fmt.Printf("d  : %v\n", d)

	fmt.Println("--------------------------------")

	// 多维数组 第一个值表示数组的长度，第二个表示每一个里面有几个元素
	var arr3 [5][3]int
	var arr4 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	e := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	f := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Println(arr3, arr4)
	fmt.Println(e, f)

	fmt.Println("--------------------------------")

	// 切片 Slice
	//1.声明切片
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)

	fmt.Println("--------------------------------")

	// 指针
	i := 10
	g := &i
	fmt.Printf("type of b:%T\n", g)
	h := *g
	fmt.Printf("type of c:%T\n", h)
	fmt.Printf("value of c:%v\n", h)

	fmt.Println("-------------Map----------------")

	map1 := make(map[string]int, 8)
	map1["张三"] = 90
	map1["李四"] = 100
	fmt.Println(map1)
	fmt.Println(map1["张三"])

	userInfo := map[string]string{
		"username": "hello",
		"password": "123456",
	}
	fmt.Println("-------------判断Map中的key是否存在----------------")
	v, ok := userInfo["username"]
	fmt.Println(userInfo)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	fmt.Println("-------------map的遍历----------------")
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println("-------------只遍历Key----------------")
	for k := range map1 {
		fmt.Println(k)
	}

	fmt.Println("-------------删除Map中的元素----------------")
	delete(map1, "张三")
	fmt.Println(map1)

	fmt.Println("-------------元素为map类型的切片----------------")
	// 类似 java 的 lsit<Map<String,String>>
	mapSlice := make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	mapSlice[1] = make(map[string]string, 10)
	mapSlice[1]["name"] = "李四"
	mapSlice[1]["password"] = "1234567"
	mapSlice[1]["address"] = "深圳"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// 类似 java Map<key,List<>>
	fmt.Println("-------------值为切片类型的map----------------")

	sliceMap := make(map[string][]string, 0)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

	fmt.Println("-------------类型别名和自定义类型-------------")

	// 类型定义
	type NewInt int

	//类型别名
	type MyInt = int

	var aa1 NewInt
	var bb1 MyInt

	fmt.Printf("type of aa1:%T\n", aa1) //type of a:main.NewInt
	fmt.Printf("type of bb1:%T\n", bb1) //type of b:int

	fmt.Println("-------------结构体-------------")

	var p1 person
	p1.name = "test"
	p1.city = "sz"
	p1.age = 23
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p1: %#v\n", p1)

	fmt.Println("-------------匿名结构体-------------")

	// 在定义一些临时数据结构等场景下还可以使用匿名结构体。
	var user struct {
		Name string
		Age  int
	}
	user.Name = "test"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	fmt.Println("-------------创建指针类型结构体-------------")
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2: %T\n", p2)
	fmt.Printf("p2=%#v\n", p2)

	fmt.Println("-------------取结构体的地址实例化-------------")
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}

	fmt.Println("-------------使用键值对初始化-------------")
	// 使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
	p5 := person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"pprof.cn", city:"北京", age:18}

	fmt.Println("-------------结构体指针进行键值对初始化-------------")
	p6 := &person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"pprof.cn", city:"北京", age:18}

	// 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}

	fmt.Println("-------------使用值的列表初始化-------------")
	// 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	// 使用这种格式初始化时，需要注意：
	// 1.必须初始化结构体的所有字段。
	// 2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 3.该方式不能和键值初始化方式混用。
	p8 := &person{
		"pprof.cn",
		"北京",
		18,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}

	fmt.Println("-------------面试题-------------")

	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// 定义结构体 类似java的类
type Person struct {
	phone    string
	password string
}

// 同样类型的字段也可以写在一行
type Person1 struct {
	phone, password string
}
type person struct {
	name string
	city string
	age  int8
}
type student struct {
	name string
	age  int
}
