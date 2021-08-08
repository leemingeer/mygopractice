package main

type Person struct {
	Age int
}

func (p *Person) GrowUp() {
	p.Age++
}

func main() {
	m := map[string]Person{
		"zhangsan": Person{Age: 20},
	}
	// map的value是不可寻址的,map扩容时，即rehashing，value内存迁移，地址会变。map作为形参或返回值时，是值拷贝，新的内存地址是Persion对象的拷贝。因为无法修改原对象
	//m["zhangsan"].Age = 23
	//m["zhangsan"].GrowUp()

	// 直接操作Persion对象本身
	mp := m["zhangsan"]
	mp.Age = 23
	mp.GrowUp()

	// 使用指针，因为value是地址，rehashing后，取到的value是Persion实例的地址，这个对象地址不会变
	mpointer := map[string]*Person{
		"zhangsan": &Person{Age: 20},
	}
	mpointer["zhangsan"].Age = 23
	mpointer["zhangsan"].GrowUp()
}