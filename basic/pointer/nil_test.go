package pointer

import "testing"

type Engine struct {
	Name string
}

func (e *Engine) Say() {
	println("say..")
	//debug模式下，发生空指针恐慌时总是无限跳到这行
	println(e.Name)
}

func CallMeth() error {
	var e = &Engine{}
	e = nil
	e.Say()
	return nil
}

func TestNilCall(t *testing.T) {
	err := CallMeth()
	println(err.Error())
}
