package method

type Person struct {
	name string
}

func Say3() (p *Person) {
	p = &Person{name: "na"}
	return
}
func Say4() (p *Person) {
	p = &Person{name: "na"}
	return p
}
