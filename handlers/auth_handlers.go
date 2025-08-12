package handlers

import (
	"PracticalProject/dto"
	"PracticalProject/services"
	"PracticalProject/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) AuthHandler {
	return AuthHandler{service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input dto.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := h.service.Register(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User creation filed",
			"Error":   err.Error(),
		})
		return
	}
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"Error": "Password hashing filed"})
	//	return
	//}
	//user := models.User{
	//	Name:     input.Name,
	//	Email:    input.Email,
	//	Password: string(hashedPassword),
	//	Phone:    input.Phone,
	//	Address:  input.Address,
	//}

	//if err := config.DB.Create(&user).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"Error": "User creation filed"})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!!!"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input dto.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	user, err := h.service.Login(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email and password",
			"Error":   err.Error(),
		})
		return
	}

	//if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid email"})
	//	return
	//}
	//
	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"Error": "Incorrect password"})
	//	return
	//}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Filed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}
