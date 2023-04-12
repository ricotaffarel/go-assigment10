package middlewares

import (
	"assigment10/database"
	"assigment10/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		// productId, err := strconv.Atoi(c.Param("productId"))
		// if productId == 0 {

		// } else if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 		"error":   "Bad Request",
		// 		"messege": "Invalid parameter",
		// 	})
		// 	return
		// }

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.First(&User, "id = ?", userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusLocked, gin.H{
				"error":   "Data not found",
				"messege": "Data doesn't exist",
			})
			return
		}

		if User.Role != "user" {
			c.AbortWithStatusJSON(http.StatusLocked, gin.H{
				"error":   "Your don't have access",
				"messege": "Please create access",
			})
			return
		}

		c.Next()
	}
}

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.First(&User, "id = ?", userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusLocked, gin.H{
				"error":   "Data not found",
				"messege": "Data doesn't exist",
			})
			return
		}

		if User.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusLocked, gin.H{
				"error":   "Your don't have access",
				"messege": "Please create access",
			})
			return
		}

		c.Next()
	}
}
