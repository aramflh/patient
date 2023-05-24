package main

import (
	"github.com/gin-gonic/gin" // Go web frameworks
	"log"
	"patient/controllers"  // This package contains functions managing patient account, medecin and pharamacien
	"patient/initializers" // This package contains functions enabling the initialization of the DB and the env var
	"patient/requests"     // This package contains the SQL requests to be executed
)

// init loads before the main function
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	//initializers.GenStructFromDB()
}

func main() {
	r := gin.Default()

	/**********************
	 *   ROUTES
	 *********************/

	/*
	 * SQL requests
	 */
	r.GET("/requests/1", requests.DoRequest1)
	r.GET("/requests/2", requests.DoRequest2)
	r.GET("/requests/3", requests.DoRequest3)
	r.GET("/requests/4", requests.DoRequest4)
	r.GET("/requests/5", requests.DoRequest5)
	r.GET("/requests/6", requests.DoRequest6)
	r.GET("/requests/7", requests.DoRequest7)
	r.GET("/requests/8", requests.DoRequest8)
	r.GET("/requests/9", requests.DoRequest9)
	r.GET("/requests/10", requests.DoRequest10)

	/*
	 * SIGN UP
	 */

	r.POST("/signup", controllers.SignUp)

	// Create a pharmacien
	r.POST("/pharmaciens", controllers.PharmaciensCreate)
	// Create a medecin
	r.POST("/medecins", controllers.MedecinsCreate)

	// Run the server

	if r.Run() != nil {
		log.Fatal("Unable to run the server")
		return
	}
}
