package learn

import "testing"

func Test_struct(t *testing.T) {

	{
		var p = Person{18}
		p.grow()
		println(p.age)
		p.grow2()
		println(p.age)
	}
	println("--------")

	{
		var p = &Person{18}
		p.grow()
		println(p.age)
		p.grow2()
		println(p.age)
	}

}

/*
value method 可以被 pointer和value 对象调用，而
pointer method 只能被 pointer 对象调用

不支持pointer method被value对象调用
*/
func Test_Interface(t *testing.T) {
	var canTalk CanTalk
	man := Man{}     // man是个值类型
	woman := Woman{} // man是个值类型
	canTalk = &man
	canTalk = woman
	//canTalk = man
	canTalk = &woman

	println(canTalk)
}
