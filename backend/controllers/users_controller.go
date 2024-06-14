package controllers

import (
	"backend/domain"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginRequest domain.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, userID, userType, err := services.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userId": userID,
		"type":   userType,
	})
}
