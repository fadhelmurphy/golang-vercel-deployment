package main

import (
	"golang-vercel-deployment/api"
)

func main() {
	// Start REST Server on main thread
	router := api.RunRoute()
	router.Run(":3000")
}