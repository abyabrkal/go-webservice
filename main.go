package main

import (
	"fmt"
	"github.com/abyabrkal/webservice/models"
)

func main() {
	//fmt.Println("Hello Go! Hi from module")

	u := models.User{
		ID: 2,
		FirstName: "Tony",
		LastName: "Stank",
	}

	fmt.Println(u);
}