package main

import "fmt"

// 打印就用这个语句，fmt.Println，注意要引入对应的fmt包
// 不过如果安装了插件，那么会自动引入用到的包，而且如果某个包没有被用到，插件会自动删除掉这个包
// 此外，注意包里面能够在外部使用的函数必定是大写字母开头的，这是约定

func main() {
	fmt.Println("Hello world")
}

// go run main.go即可运行本代码，会看到命令行里面输出了要打印的内容
// 如果需要编译成可执行文件，那么需要go build main.go，最终会产出main.exe可执行文件（Windows系统下）
