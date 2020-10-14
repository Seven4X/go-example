package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//https://github.com/golang/go/issues/6858
func TestConvert(t *testing.T) {

	msg := Message{1, "Hello"}
	//如何Message和Message2d定义不同没法通过()转化
	msg2 := Message2(msg)
	_ = msg2
	assert.Equal(t, msg, msg2, "struct直接比较")

	println("0-----------------")
	//同类型struct能够通过==判断
	msg3 := Message{2, "Hello"}
	if &msg == &msg3 {
		println("==")
	} else {
		println("!=")
	}

	if msg == msg3 {
		println("==")
	} else {
		println("!=")
	}

	println(msg2.Id)
	println(msg2.Text)
}
