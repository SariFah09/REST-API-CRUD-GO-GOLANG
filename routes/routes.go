package routes

import (
	"testgo/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/animals", controllers.FindAnimals)
	r.POST("/animals", controllers.CreateAnimal)
	r.GET("/animals/:id", controllers.FindAnimal)
	r.PUT("/animals/:id", controllers.UpdateAnimal)
	r.DELETE("/animals/:id", controllers.DeleteAnimal)

	return r
}
