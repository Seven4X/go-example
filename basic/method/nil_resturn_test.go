package method

import "testing"

//两种方法都work
func TestSay(t *testing.T) {
	{
		p := Say3()
		println(p)

	}

	{
		p := Say4()

		println(p)
	}
}
