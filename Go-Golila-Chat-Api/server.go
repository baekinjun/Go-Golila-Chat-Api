package main

import (
	router "wooah/router"
)

func main() {
	e := router.Init()
	// Server
	e.Logger.Fatal(e.Start(":1213"))
}
