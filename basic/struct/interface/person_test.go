package person

import (
	"testing"
)

func TestNewMethod(t *testing.T) {
	{
		var p = New1()
		println(p.name)
		println(p)
		// p是引用 New1方法内和调用后是同一个对象
		println("----------------------->")
	}

	{
		var p = New3()
		println(p.name)
		println(&p)
		//p是值copy，方法内和方法外不是同一个地址
		println("----------------------->")

	}
	{
		var p = New2()
		println(p.name)
		println(p)
		println("----------------------->")

	}
	{
		var p = New4()
		println(p.name)
		println(&p)
		println("----------------------->")

	}

	{
		var p = New1()
		println(p.name)
		println(p)
		p = New1()
		println(p.name)
		println(p)
		// p是引用 New1方法内和调用后是同一个对象
		println("----------------------->")
	}

}
