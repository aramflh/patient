package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"patient/dal"
	"patient/initializers"
)

// init loads before the main function
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	// Transform the tables in the DB into struct in dal/ directory
	initializers.GenStructFromDB()
}

func main() {

	Q := dal.Use(initializers.DB)

	syst, err := Q.SystemeAna.First()

	// Log an error if err is not nil (null)
	if err != nil {
		log.Fatal("Failed to fetch DB")
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Systeme Ana : %v", syst.NomSysAna),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
