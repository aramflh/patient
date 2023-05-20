package main

import (
	"dossier-patient/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
	// "database/sql"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

// init loads before the main function
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.GenStructFromDB()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
