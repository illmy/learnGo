# 第一个GO程序

### HelloWorld程序

[代码地址](./../code/go_learning/src/ch1/main/hello_world.go)
```GO
package main   // 包 声明代码所在的模块

import "fmt"   // 引入代码依赖

// 功能实现
func main()  {
	fmt.Println("Hello World")
}
```

运行代码
> go run hello_world.go

编译代码
> go build hello_world.go
> ./hello_world

### 应用程序入口

1. 必须是main包：package main
2. 必须是main方法：func main()
3. 文件名不一定是main.go

### 退出返回值

与其他主要编程语言的差异
* Go中main函数不支持任何返回值
* 通过os.Exit来返回状态 os.Exit(0)正常返回 os.Exit(-1)返回异常

### 获取命令行参数

与其他主要编程语言的差异
* main函数不支持传入参数
  比如func main(arg []string)这样是不被允许的
* 在程序中直接使用os.Args获取命令行参数