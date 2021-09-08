package main

import (
	"fmt"
	"net/http"

	"github.com/pluralsight/webservice/controllers"
	"github.com/pluralsight/webservice/models"
)

func main() {

	u1, _ := models.AddUser(models.User{FirstName: "Tricia", LastName: "McMillan"})
	fmt.Println(u1)

	u2, _ := models.AddUser(models.User{FirstName: "Ekaterina", LastName: "Wexler"})
	fmt.Println(u2)

	fmt.Println(models.GetUsers())

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
