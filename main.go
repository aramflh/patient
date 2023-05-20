package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

// Q initializes a query var
//var Q dal.Query

// init loads before the main function
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	// Transform the tables in the DB into struct in dal/ directory
	initializers.GenStructFromDB()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
