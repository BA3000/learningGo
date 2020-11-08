package main

import "fmt"

var g = "全局变量"

func getUserinfo() (string, int) {
	return "Luna Lovegood", 10
}

func main() {
	// var声明变量
	// var 变量名称 类型
	// 变量名由字母、数字、下划线组成，首个字符不能够是数字，此外go的关键词和保留字不能够用作为变量名

	var username string
	// 变量声明之后没有初始化的话那么值默认为""
	fmt.Println(username)
	username = "初始化" // 一种初始化方式
	fmt.Println("username = ", username)

	// 注意下面这里其实并不是必要的，因为go其实可以通过类型推导得出数据的类型
	var username1 string = "第二种初始化方式"
	fmt.Println("username1 = ", username1)

	// 直接这种就行，可以类型推导，另外注意go中同一个作用域里面不能够重复声明同名的变量，否则会报错，不过可以重复赋值
	var username2 = "第三种初始化方式"
	fmt.Println("username2 = ", username2)

	var a123 = "张三"
	fmt.Println("a123 = ", a123)

	var m_a = "例子"
	fmt.Println("m_a = ", m_a)

	/* 一次声明多个变量
	 var 变量名称，变量名称 类型
	 上面这样的话声明出来的变量类型都是一样的

	另一种：
	var(
		变量名称 类型
		变量名称 类型
	)
	上面这种声明出来的可以每个变量类型不一样
	*/

	var a1, a2 string
	a1 = "aaa"
	a2 = "aaaaaaa"
	fmt.Println("a1, a2 = ", a1, a2)

	var (
		username3 string
		age       int
		sex       string
	)
	username3 = "哈利·波特"
	age = 21
	sex = "男"

	fmt.Println("username3, age, sex = ", username3, age, sex)

	// 可以利用类型推导直接省略掉类型
	var (
		username4 = "张飞"
		age4      = 123
	)
	fmt.Println("username4, age4 = ", username4, age4)

	// 短变量声明法
	// 即直接使用:=进行声明变量，注意这种方法只能够用在局部变量的声明上，不能够用在全局变量的声明上
	username5 := "洛夫古德"
	fmt.Println("username5 = ", username5)

	fmt.Printf("username5 类型：%T", username5)

	// 短变量初始化方法其实还可以一次性声明多个变量并进行初始化
	a, b, c := 12, 13, 20
	fmt.Println("a, b, c", a, b, c)

	a3, b3, c3 := 12, 2, "C"
	fmt.Println("a3, b3, c3 = ", a3, b3, c3)

	// 全局变量不能够使用短变量声明法
	fmt.Println("g:", g)

	// 匿名变量
	// 使用多重赋值的时候如果想要忽略掉某个值就可以使用匿名变量
	// 匿名变量用一个下划线表示
	var username6, age6 = getUserinfo()
	fmt.Println("username6, age6:", username6, age6)

	var username7, _ = getUserinfo()
	fmt.Println("username7:", username7)
	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明，可以重复使用

	var _, age7 = getUserinfo()
	fmt.Println("age7:", age7)
}
