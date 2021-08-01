package main

import (
	"alta-store/config"
	"alta-store/routes"
)

func main() {
	config.InitDb()
	e := routes.Start()
	e.Logger.Fatal(e.Start(":8000"))
}
