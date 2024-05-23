package main

import (
	"clicks/config"
	"clicks/routers"
)

func main() {
	config.InitDB()
	config.GetDB()

	routers.StartServer().Run(":1111")
}
