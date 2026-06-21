package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"backend/database"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "price is invalid",
		})
		return
	}

	product := models.Product{
		Title:       title,
		Description: description,
		Price:       uint(price),
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "images are required",
		})
		return
	}

	files := form.File["images"]

	os.MkdirAll("uploads/products", os.ModePerm)

	var images []models.ProductImage

	for _, file := range files {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
		filePath := filepath.Join("uploads/products", filename)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to save image",
			})
			return
		}

		image := models.ProductImage{
			ProductID: product.ID,
			ImageURL:  "/" + filePath,
		}

		images = append(images, image)
	}

	if len(images) > 0 {
		if err := database.DB.Create(&images).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	product.Images = images

	c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
	var products []models.Product

	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

