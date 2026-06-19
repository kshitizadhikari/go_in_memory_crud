package routes

import (
	"go_in_memory_crud/internal/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	v1 := api.Group("/v1")

	userService := user.NewUserService()
	userHandler := user.NewUserHandler(userService)

	v1.GET("/users", userHandler.GetUsers)
	v1.GET("/users/:id", userHandler.GetUserById)
	v1.POST("/users", userHandler.CreateUser)
	v1.PUT("/users/:id", userHandler.UpdateUser)
	v1.PATCH("/users/:id", userHandler.PatchUser)

}
