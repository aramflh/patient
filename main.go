package main

import (
	"github.com/gin-gonic/gin"
	"patient/controllers"
	"patient/initializers"
)

// init loads before the main function
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	//initializers.GenStructFromDB()
}

func main() {
	r := gin.Default()

	// Create a pharmacien
	r.POST("/pharmaciens", controllers.PharmaciensCreate)
	// Create medecin
	r.POST("/medecins", controllers.MedecinsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
