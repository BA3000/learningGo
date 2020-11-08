package main

import "fmt"

func main() {
	// 三种文本输出方法
	// Println会自动在输出文本之后加多一个空行，Print则不会
	fmt.Println("你好，golang")
	fmt.Print("你好 golang")
	fmt.Printf("你好 狗狼")

	// 一次输出多个值的时候Print出来的文本之间不会存在空格
	fmt.Print("A", "B", "C")
	// 一次输出多个值的时候Println出来的文本之间会存在空格
	fmt.Println("A", "B", "C")

	// Println和Printf区别
	// 注意golang中变量定义了之后必须使用，不然会报错
	// 通常的定义方式就是var + 变量名 + 类型 = 表达式
	var a = "test"
	fmt.Println(a)

	// 如果是Printf的话必须要格式化的方式来输出，和C/C++差不多
	fmt.Printf("%v\n", a)

	// 此外还可以用类型推导方式来进行定义，变量名 := 表达式
	var b int = 10
	var c int = 3
	d := 20 // 这个就是类型推导方式

	fmt.Println("b=", b, "c=", c, "d=", d)
	fmt.Printf("b = %v, c = %v, d = %v\n", b, c, d)
	// 除了%v之外还有%d输出十进制，%o输出八进制，%b输出二进制等，具体的可以看文档

	// 使用Printf打印一个变量的类型
	fmt.Printf("d = %v type of d is %T\n", d, d)
}
