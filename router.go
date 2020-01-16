package main

import (
	"github.com/gin-gonic/gin"
	"goserviceJenkinsDocker/controllers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}

type Routes []Route

var publicRoutes = Routes{
	{"SaveUserData", "POST", "/user", controllers.SaveUserData},
}

func NewRouter() {
	router := gin.Default()

	/* public routes */
	public := router.Group("/api")
	for _, route := range publicRoutes {
		switch route.Method {
		case "GET":
			public.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			public.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			public.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			public.DELETE(route.Pattern, route.HandlerFunc)
		default:
			public.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route.",
				})
			})
		}
	}
	router.Run(":9090")
}
