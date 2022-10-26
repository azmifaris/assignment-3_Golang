package main

import (
	"assignment-3_AzmiFarisM/routers"
)

func main() {
	routers.StartServer().Run(":8000")
}
