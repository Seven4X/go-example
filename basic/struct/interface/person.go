package person

type Person struct {
	name string
	age  int
}

//推荐使用 简洁
func New1() (p *Person) {
	p = &Person{}
	p.name = "New1"
	p.age = 13
	println(p)

	return p

}

func New3() (p Person) {
	//p在没设置之前为空值 默认值
	p.name = "New3"
	p.age = 13
	println(&p)

	return p

}

func New2() (p *Person) {
	//这里只是没有使用p指针
	println(p)

	ip := &Person{}
	ip.name = "New2"
	ip.age = 19
	println(ip)
	return ip

}

func New4() (p Person) {
	println(&p) //没有指针

	ip := &Person{}
	ip.name = "New4"
	ip.age = 19
	println(ip)
	return *ip //取值copy

}
