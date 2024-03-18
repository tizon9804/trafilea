package main

import "trafilea/routes"

func main() {
	router := routes.NewRouter()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
