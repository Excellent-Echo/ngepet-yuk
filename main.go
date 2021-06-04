package main

import (
	"ngepet-yuk/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	routes.UserRoute(r)
	routes.UserDetailRoute(r)
	routes.UserTransactionRoute(r)
	routes.SubTypeRoute(r)
	routes.CategoryRoute(r)
	routes.MasteryRoute(r)
	routes.CoursesRoute(r)
	routes.CryptoRoute(r)

	r.Run(":8080")
}
