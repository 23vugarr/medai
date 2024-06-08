package handlers

import (
	"backend/src/models"
	"backend/src/repository"
	"backend/src/schemas"
	"backend/src/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"response": "ok",
		})
	}
}

func (u *UserHandler) Register(userRepo *repository.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.User

		err := json.NewDecoder(c.Request.Body).Decode(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": err})
			return
		}

		err = userRepo.CreateUser(req.Name, req.Surname, utils.HashPassword(req.Password), req.PhoneNumber, req.Birthday)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"response": "user registered successfully",
		})
	}
}

func (u *UserHandler) Login(userRepo *repository.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req schemas.UserLogin

		err := json.NewDecoder(c.Request.Body).Decode(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": "error while decoding the body"})
			return
		}

		err = userRepo.LoginUser(req.PhoneNumber, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"response": err})
			return
		}

		jwt, _ := utils.GenerateJWT(req.PhoneNumber)

		c.JSON(http.StatusOK, gin.H{
			"response": "ok",
			"payload": gin.H{
				"jwt": jwt,
			},
		})
	}
}
