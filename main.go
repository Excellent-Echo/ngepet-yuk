package main

import (
	"ngepet-yuk/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// routes.UserRoute(r)
	// routes.UserDetailRoute(r)
	// routes.UserTransactionRoute(r)
	// routes.SubTypeRoute(r)
	// routes.CategoryRoute(r)
	// routes.MasteryRoute(r)
	// routes.CoursesRoute(r)
	routes.CryptoRoute(r)

	r.Run(":8080")
}
