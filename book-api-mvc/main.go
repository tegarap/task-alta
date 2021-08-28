package main

import (
	"book-api-mvc/config"
	"book-api-mvc/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
