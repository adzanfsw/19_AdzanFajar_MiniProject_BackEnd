package main

import (
	"justrun/config"
	"justrun/routes"
)

func main() {
	//initiateDB
	config.InitDB()
	config.AutoMigrate()

	//initRoutes
	e := routes.New()

	//starting the server
	e.Start(":3000")
}
