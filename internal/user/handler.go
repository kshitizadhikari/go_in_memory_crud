package user

import (
	"fmt"
	"strconv"

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

func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Errorf("invalid user id : %d", id),
		})
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, user)
}
