package main


import "fmt"

func main() {
	// 选择工厂类型，创建具体工厂
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike") // 返回nike工厂，变体已经确定下来了

	adidasShoe := adidasFactory.makeShoe() // 返回具体产品，赋值给抽象产品
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	println("nike shoe")
	printShoeDetails(nikeShoe) // // 通过抽象产品调用产品的方法，注意此时变体已经确定下来是nike了，只能生产nike的Shoe
	println("nike shirt")
	printShirtDetails(nikeShirt)

	println("adidas shoe")
	printShoeDetails(adidasShoe)
	println("adidas shirt")
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}