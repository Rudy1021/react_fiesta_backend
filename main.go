package main

import (
	_ "react_fiesta_backend/database"
	"react_fiesta_backend/routes"
)

func main() {
	router := routes.InitRouter()
	router.Run()
}
