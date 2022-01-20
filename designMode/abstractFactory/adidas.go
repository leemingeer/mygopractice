package main

// 具体工厂
// 返回抽象产品
type adidas struct {
}

// 具体工厂返回抽象产品
func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 10,
		},
	}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 15,
		},
	}
}