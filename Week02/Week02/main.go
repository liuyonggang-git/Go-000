package main

import (
	"Week02/service"
	"fmt"
)

func main() {
	username1, err1 := service.GetUsernameById1("1")
	fmt.Println(username1)
	fmt.Println(err1)

	username2, err2 := service.GetUsernameById2("4")
	fmt.Println(username2)
	fmt.Println(err2)
}
