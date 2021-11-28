package main

import (
	"justrun/config"
	"justrun/route"
)

func main() {

	config.InitDB()
	config.AutoMigrate()

	e := route.RouteShoes()
	e.Start(":3000")
}
