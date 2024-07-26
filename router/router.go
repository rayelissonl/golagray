package router

import  (
	"github.com/gin-gonic/gin"
)

func Initializer() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}