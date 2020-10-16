package main

import (
	"fmt"
)

type test struct {
	name string
}

/*
妙啊！！！！
Golang 中没有“对象”
可以正常输出。Go 本质上不是面向对象的语言，Go 中是不存在 object 的含义的，Go 语言书籍中的对象也和 Java、PHP 中的对象有区别，不是真正的”对象”，是 Go 中 struct 的实体。

调用 getName 方法，在 Go 中还可以转换，转换为：Type.method(t Type, arguments)
所以，以上代码 main 函数中还可以写成：
*/
func (t *test) getName() {
	fmt.Println("hello world")
}
func main() {
	var t *test
	//t = nil
	t.getName()

	newMap := make(map[string]int)
	fmt.Println(newMap["a"])
}

//只可以发送值
func setData1(ch chan<- string) {
	//TODO
}

func setData2(ch <-chan string) {
	//TODO
}
