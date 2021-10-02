package routes

import (
	"go/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside main middleware.")
		next.ServeHTTP(w, r)
	})
}

func ConfingRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("fastfood")
	{
		lojas := main.Group("lojas")
		{
			lojas.GET("/:id", controllers.ShowLojas)
			lojas.GET("/list", controllers.ListLojas)
			lojas.POST("/register", controllers.CreateLoja)
			lojas.PUT("/:id", controllers.UpdateLoja)
			lojas.DELETE("/:id", controllers.DeleteLoja)
		}
	}

	return router
}
