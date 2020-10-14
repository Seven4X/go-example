package main

import "example/basic/init/demo"

func init() {
	println("init-main")
}

// 深度优先
func main() {

	demo.Say()
	demo.Init()
	println("main")
}
