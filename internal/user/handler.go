package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.service.GetUsers()
	fmt.Println(users)
	c.JSON(200, users)
}
