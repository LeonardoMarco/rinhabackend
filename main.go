package main

import (
	"rinhabackendleo/src/routes"
)

func main() {
	router := routes.InitRoutes()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
