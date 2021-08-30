package main

import (
	"book-api-mvc/config"
	"book-api-mvc/routes"
	"os"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":"+os.Getenv("SERV_PORT")))
}
