package main

var sa string //空值 是空串

func main() {
	//作用域
	sa := say()

	println(sa)

	global()
}

func global() {
	println(sa)
}
func say() string {
	return "hh"
}
