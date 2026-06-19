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
			"error": "invalid user id",
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

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req User
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request id",
		})
		return
	}

	var req User
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := h.service.UpdateUser(id, req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, user)
}
