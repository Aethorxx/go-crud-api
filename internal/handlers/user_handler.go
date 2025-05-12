package handlers

import (
	"net/http"
	"strconv"

	"go-crud-api/internal/models"
	"go-crud-api/internal/services"
	"go-crud-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.LogError("CreateUser", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		utils.LogError("CreateUser", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.LogOperation("CreateUser", user.ID, "User created successfully")
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	var params models.PaginationParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем значения по умолчанию
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Limit == 0 {
		params.Limit = 10
	}

	response, err := h.userService.ListUsers(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.userService.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.LogError("UpdateUser", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.LogError("UpdateUser", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.UpdateUser(uint(id), req)
	if err != nil {
		utils.LogError("UpdateUser", err)
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	utils.LogOperation("UpdateUser", user.ID, "User updated successfully")
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.LogError("DeleteUser", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		utils.LogError("DeleteUser", err)
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	utils.LogOperation("DeleteUser", uint(id), "User deleted successfully")
	c.Status(http.StatusNoContent)
}
