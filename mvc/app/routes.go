package app

import "microservices/mvc/controllers"

func mapRoutes() {
	router.GET("/users/:user_id", controllers.GetUser)
}