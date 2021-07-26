package main

import (
	"1stgoproj/designMode/services"
	"fmt"
)

func main() {
	//var service1 services.IService = services.NewsService{} // Get() method input para is not address
	var service1 services.IService = new(services.NewsService) // point method
	fmt.Println(service1.Get(123))
	var service2 services.IService = new(services.UserService) // or &services.UserService{}
	fmt.Println(service2.Get(123))

	var service3 services.IService = services.ServiceFactory{}.Create("news")
	fmt.Println(service3.Get(123))

	var service4 services.IService = new(services.ServiceFactory).Create("news")
	fmt.Println(service4.Get(123))

	var service5 services.IService = services.NewServiceFactory().Create("news")
	fmt.Println(service5.Get(123))

}
