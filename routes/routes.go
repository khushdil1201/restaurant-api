package routes

import (
    "restaurant-api/controllers"
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
    r.POST("/register", controllers.RegisterUser)

    authorized := r.Group("/")
    authorized.Use(controllers.AuthMiddleware())
    {
        authorized.GET("/menu", controllers.GetMenu)
        authorized.POST("/menu", controllers.AddMenuItem)
        authorized.DELETE("/menu/:id", controllers.DeleteMenuItem)
    }
}
