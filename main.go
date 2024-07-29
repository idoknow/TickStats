package main

import (
	"github.com/soulter/tickstats/routes"
)

func main() {
	routes.SetupRouter().Run(":8080")
}
