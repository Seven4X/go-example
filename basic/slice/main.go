package main

import (
	"fmt"
)

//https://www.infoq.cn/article/bss370GhURBSPCHzUIxA
/*
1）如果切片的容量小于 1024 个元素，那么扩容的时候 slice 的 cap 就翻番，乘以 2；一旦元素个数超过 1024 个元素，增长因子就变成 1.25，即每次增加原来容量的四分之一。

2）如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组，如果扩容之后，超过了原数组的容量，那么，Go 就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
*/
func main() {
	array := [4]int{10, 20, 30, 40}
	slice := array[0:2]
	newSlice := append(append(append(slice, 50), 100), 150)
	newSlice[1] += 1
	fmt.Println(slice)
	fmt.Println(newSlice)
}
