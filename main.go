package main

import (
	"fmt"
	_ "react_fiesta_backend/database"
	"react_fiesta_backend/routes"
)

func main() {
	fmt.Println("\033[31mError: \033[0mSomething went wrong")
	router := routes.InitRouter()
	router.Run()
}
