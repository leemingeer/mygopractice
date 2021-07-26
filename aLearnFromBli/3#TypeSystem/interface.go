package main

import "fmt"

// why we use interface, because there may be multi object and there methods in this package
type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

// user type who has notify method
type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("user %s notified\n", u.name)
}

func (u user) printName() {
}

// must be pointer receiver
func (u *user) changeName(newName string) {

}

// another type who has notify method
type Admin struct {
	Name string
}

func (a *Admin) notify() {
	fmt.Printf("Admin %s notified\n", a.Name)
}

func main() {

	user := user{"xiaoming", "123@qq.com"}
	admin := Admin{"daming"}

	// by user object
	user.notify()
	(&user).notify()
	// by interface
	sendNotification(user)
	sendNotification(&user)

	// by admin o
	admin.notify()
	(&admin).notify()
	// sendNotification(admin) // wrong, receiver is pointer type
	sendNotification(&admin)

}
