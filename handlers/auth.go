package handlers

import (
	"net/http"

	"backend/database"
	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	NationalID string `json:"national_id" binding:"required"`
	Password   string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	NationalID string `json:"national_id" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !utils.IsValidIranianNationalID(req.NationalID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "national_id is invalid",
		})
		return
	}

	var existingUser models.User
	if err := database.DB.Where("national_id = ?", req.NationalID).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "user already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := models.User{
		NationalID:   req.NationalID,
		PasswordHash: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.NationalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"access": token,
		"user": gin.H{
			"id":          user.ID,
			"national_id": user.NationalID,
		},
	})
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !utils.IsValidIranianNationalID(req.NationalID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "national_id is invalid",
		})
		return
	}

	var user models.User
	if err := database.DB.Where("national_id = ?", req.NationalID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "national_id or password is incorrect",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "national_id or password is incorrect",
		})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.NationalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access": token,
		"user": gin.H{
			"id":          user.ID,
			"national_id": user.NationalID,
		},
	})
}