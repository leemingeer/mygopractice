package main

import (
	"log"
	"reflect"
)

var data = map[string]interface{}{
	"id":    1001,
	"name":  "apple",
	"price": 16.25,
}

type Fruit struct {
	ID    int     `key:"id"`
	Name  string  `key:"name"`
	Price float64 `key:"price"`
}

// 遍历struct并且自动进行赋值
func structByReflect(data map[string]interface{}, inStructPtr interface{}) {
	rType := reflect.TypeOf(inStructPtr) // Type类型对象
	rVal := reflect.ValueOf(inStructPtr) // Value类型对象
	if rType.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}
	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)  // StructField对象
		f := rVal.Field(i)   // Value对象
		// 得到结构体中tag中的字段名
		key := t.Tag.Get("key")

		if v, ok := data[key]; ok {
			f.Set(reflect.ValueOf(v))
		} else {
			panic(t.Name + " not found")
		}
	}
}
func main() {
	//fruit := newFruit(data)
	fruit := Fruit{}
	structByReflect(data, &fruit)
	log.Println("fruit:", fruit)
}