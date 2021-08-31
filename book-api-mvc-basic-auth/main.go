package main

import (
	"book-api-mvc/config"
	m "book-api-mvc/middlewares"
	"book-api-mvc/routes"
	"os"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":"+os.Getenv("SERV_PORT")))
}
