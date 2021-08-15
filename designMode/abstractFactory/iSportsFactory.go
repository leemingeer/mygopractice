// 抽象工厂接口
// 返回具体工厂
package main
import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe // 返回不同抽象产品的方法
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
