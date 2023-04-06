package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"user_service/models"
	"user_service/repository"
)

type UserController struct {
	repo repository.UserRepository
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{repo: repository.NewUserRepository(db)}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	u.repo.CreateUser(&newUser)
	c.JSON(http.StatusOK, newUser)
}
