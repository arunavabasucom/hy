package main

import (
	"hy/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupPingRoutes(r)
	r.Run()
}
