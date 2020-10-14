package method

import "testing"

func buildToken(username string) (tokenstr string) {
	println(&tokenstr)

	tokenstr = username + "123"
	println(&tokenstr)

	return
}

func TestMethod(t *testing.T) {
	s := buildToken("abc")
	println(&s)
	println(s)

}
