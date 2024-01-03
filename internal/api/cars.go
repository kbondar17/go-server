package api

import (
	"fmt"
	db "yandex/second/package/database"

	"github.com/gin-gonic/gin"
)

func CarsHandlers(r *gin.Engine, database *db.Database) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/cars", func(c *gin.Context) {
		res, ok := database.GetAll()
		if !ok {
			c.JSON(404, gin.H{
				"message": "Not found",
			})
			return
		}
		c.JSON(200, res)
	})

	r.GET("/cars/:brand", func(c *gin.Context) {
		brand := c.Param("brand")
		fmt.Println("brand::", brand)
		if brand == "" {
			c.JSON(404, gin.H{
				"message": "Brand must be specified",
			})
			return
		}
		res, ok := database.GetByBrand(brand)
		if !ok {
			c.JSON(404, gin.H{
				"message": "Cars with this brand not found",
			})
			return
		}
		c.JSON(200, res)
	})
}
