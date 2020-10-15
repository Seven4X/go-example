package learn

import (
	"sync"
)

var once sync.Once

//
func Validate(context string) {
	once.Do(func() {
		println("once....")
	})
	println(&once)

	println(context)

	var one sync.Once
	one.Do(func() {
		println("只....")
	})
	println(&one)
}

// 注释
func CurrentOnce() {
	once.Do(func() {
		println("once....current")
	})

	println("current")
}
