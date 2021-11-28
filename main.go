package main

import (
	"justrun/config"
)

func main() {

	config.InitDB()
	config.AutoMigrate()

	e := routes.RouteShoes()
	e.Start(":3000")
}
