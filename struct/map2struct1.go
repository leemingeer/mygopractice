// 	将map对象转成struct
package main

import "log"

var data1 = map[string]interface{}{
	"id":    1001,
	"name":  "apple",
	"price": 16.25,
}

type Fruit1 struct {
	ID    int
	Name  string
	Price float64
}

func newFruit(data map[string]interface{}) *Fruit1 {
	s := Fruit1{
		ID:    data1["id"].(int),
		Name:  data1["name"].(string),
		Price: data1["price"].(float64),
	}
	return &s
}

func main() {
	fruit := newFruit(data1)
	log.Println("fruit:", fruit)
}
