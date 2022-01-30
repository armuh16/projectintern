package main

import (
	"fmt"
	"sejutacita/api"
	"sejutacita/api/middlewares"
	config "sejutacita/util"
)

func main() {
	fmt.Println("Hello World")
	// Configuration to Database
	config.InitDB()
	// Call the router
	e := api.New()
	middlewares.LogMiddlewares(e)
	// Logger to run server with port 8000
	e.Logger.Fatal(e.Start(":8080"))
}
