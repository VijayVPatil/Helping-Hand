package main

import (
	"fmt"

	"github.com/VijayVPatil/Helping-Hand/api"
	"github.com/VijayVPatil/Helping-Hand/controller"
)

func main() {
	fmt.Println("Starting Database connection.......")

	api := api.ApiRoute{}

	api.StartApp(controller.DbServer{})

	fmt.Printf("The main Server : %v\n", api)

}
