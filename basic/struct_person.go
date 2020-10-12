package learn

import "strconv"

/*
https://zhuanlan.zhihu.com/p/76384820
*/
type Person struct {
	age int
}

func (receiver Person) grow() {
	receiver.age = receiver.age + 1
	println("inner:" + strconv.Itoa(receiver.age))
}

func (receiver *Person) grow2() {
	receiver.age = receiver.age + 1
}

type CanTalk interface {
	// 说话方法
	Say()
}
type Man struct {
}
type Woman struct {
}

// Say()方法的全名为(*Man).Say()
// 即只有指针类型*Man才有Say()方法
func (*Man) Say() {
}

// Say()方法的全名为(Woman).Say()
// 只有值类型Woman才有Say()方法
func (Woman) Say() {
}
