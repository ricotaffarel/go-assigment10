package controllers

import (
	"assigment10/database"
	"assigment10/helpers"
	"assigment10/models"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	var User models.User
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))
	err := db.Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid user id",
		})
		return
	}

	if User.Role == "user" {
		Product.UserID = userID
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	err = db.Create(&Product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func UpdatedProduct(c *gin.Context) {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	// userID := uint(userData["id"].(float64))
	// Product.UserID = userID
	Product.ID = uint(productId)
	product := db.Preload("User").First(&Product)

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	err := product.Save(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"messege": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func ViewProduct(c *gin.Context) {
	db := database.GetDB()
	User := models.User{}
	Product := []models.Product{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	productId, _ := strconv.Atoi(c.Query("productId"))
	userID := uint(userData["id"].(float64))
	err := db.Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid user id",
		})
		return
	}

	fmt.Println(productId)
	fmt.Println(User.Role)
	if productId != 0 {
		err = db.Where("id = ?", productId).Find(&Product).Error
	} else if User.Role == "user" {
		err = db.Where("user_id = ?", userID).Find(&Product).Error
	} else {
		err = db.Find(&Product).Error
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"messege": err.Error(),
		})
		return
	}

	if len(Product) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, Product)

}

func DeletedProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	Product.ID = uint(productId)
	err := db.First(&Product).Delete(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"messege": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messege": "Product has be deleted",
	})
}
