package router

import (
	"assigment10/controllers"
	"assigment10/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	adminRouter := r.Group("/admin")
	{
		adminRouter.Use(middlewares.Authentication())
		adminRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "admin ok")
		})

		//product
		productRouter := adminRouter.Group("/product", middlewares.AdminAuthorization())
		productRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "product ok")
		})
		productRouter.POST("/create", controllers.CreateProduct)
		productRouter.PUT("/update/:productId", controllers.UpdatedProduct)
		productRouter.GET("/view", controllers.ViewProduct)
		productRouter.DELETE("/delete/:productId", controllers.DeletedProduct)
	}

	routeUser := r.Group("/user")
	{
		routeUser.Use(middlewares.Authentication())
		//product
		productRouter := routeUser.Group("/product", middlewares.UserAuthorization())
		productRouter.PUT("/update/:productId", controllers.UpdatedProduct)
		productRouter.GET("/view", controllers.ViewProduct)
	}

	return r
}
