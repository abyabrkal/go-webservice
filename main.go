package main

import (
	"net/http"
	"github.com/abyabrkal/webservice/controllers"
)

func main() {
	//fmt.Println("Hello Go! Hi from module")

	// u := models.User{
	// 	ID: 2,
	// 	FirstName: "Tony",
	// 	LastName: "Stank",
	// }

	// fmt.Println(u);

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}