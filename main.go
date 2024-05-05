/*
AUTHOR - Arunava Basu <arunava.basu@arunavabasu.com>
*/

package main

import (
	"hy/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupPingRoutes(r)
	  r.POST("/merge", routes.MergePDFHandler)
	r.Run()
}
