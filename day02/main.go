package main

import "fmt"

func main() {
	// 三种文本输出方法
	// Println会自动在输出文本之后加多一个空行，Print则不会
	fmt.Println("你好，golang")
	fmt.Print("你好 golang")
	fmt.Printf("你好 狗狼")
}
