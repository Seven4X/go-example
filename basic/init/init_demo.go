package main

import "example/basic/init_demo/demo"

func init() {
	println("init-main")
}

// 深度优先
func main() {

	demo.Say()
	demo.Init()
	println("main")
}
