package main

import (
	"ngepet-yuk/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoute(r)
	routes.UserDetailRoute(r)

	r.Run(":8080")
}
