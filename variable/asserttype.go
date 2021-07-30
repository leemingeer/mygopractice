package main

import (
	"fmt"
	_ "github.com/golang/glog"
)
type rest struct {
	name string
	result interface{}
}

var s interface{}= []map[string]map[string]string{{"x": {"y":"z"}}}
var s1 []interface{}

//var s = rest{
//	name: "ming",
//	result: res,
//}




func main(){
    // interface{} 是匿名接口类型, '=', append等操作，是在运行时发生的，编译时都是各自在栈上的变量，编译器检查类型，不一致的话，编译器就报错了。编译器保证语法和类型等层面没有问题。
	var a interface{}   = []map[string]interface{}{{"x": "y"}}
	//var aa []interface{} = []map[string]interface{}{{"x": "y"}} // aa type is not interface but slice， element happends to interface{}
	fmt.Printf("a type1: %T\n", a) //[]map[string]interface {}

	var a2 interface{} = []map[string]string{{"x": "y"}}
	fmt.Printf("a2 type: %T\n", a2) //[]map[string]string

	b := map[string]interface{}{"z": 1}
	//b2 := map[string]string{"z": "1"}

	fmt.Printf("b type: %T\n", b) //[]map[string]interface {}
	//a = append(a.([]map[string]interface{}), b2) //编译就报错，a类型是[]map[string]interface{}， b类型是map[string]string, 说interface{}类型的变量可以接受任意类型变量，也只是在运行时，要给interface{}赋值动作是在运行时发生的，在编译时interface{}和int，string等一样，都是不同的类型，他们是在同一等级的。
	//a = append(a.([]map[string]string), b2) // 虽然编译不会报错，因为已经将a断言（人为告诉编译器按这种类型解析）为b2的类型了，但在运行时会panic，因为运行时接口a的实际类型是[]map[string]interface{}, not []map[string]string）。
	a = append(a.([]map[string]interface{}), b)
	fmt.Printf("a type2: %T\n", a)

}

