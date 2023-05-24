package main

import (
	"github.com/gin-gonic/gin" // Go web frameworks
	"log"
	"net/http"
	"patient/controllers"  // This package contains functions managing patient account, medecin and pharamacien
	"patient/initializers" // This package contains functions enabling the initialization of the DB and the env var
	"patient/middleware"   // This package contains the function verifying the log status
	"patient/requests"     // This package contains the SQL requests to be executed
)

// init loads before the main function
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Creates a gin instance
	r := gin.Default()

	// Loads HTML files
	r.LoadHTMLGlob("templates/*.html")

	/**********************
	 *   ROUTES
	 *********************/

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "",
		})
	})

	requestRoutes := r.Group("/requests")
	/* SQL requests */
	requestRoutes.GET("/1", requests.DoRequest1)
	requestRoutes.GET("/2", requests.DoRequest2)
	requestRoutes.GET("/3", requests.DoRequest3)
	requestRoutes.GET("/4", requests.DoRequest4)
	requestRoutes.GET("/5", requests.DoRequest5)
	requestRoutes.GET("/6", requests.DoRequest6)
	requestRoutes.GET("/7", requests.DoRequest7)
	requestRoutes.GET("/8", requests.DoRequest8)
	requestRoutes.GET("/9", requests.DoRequest9)
	requestRoutes.GET("/10", requests.DoRequest10)

	/* Add 'pharamcien' */
	r.GET("/pharmaciens", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addPharma.html", gin.H{})
		//c.Redirect(http.StatusMovedPermanently, "/medecins")
	})
	r.POST("/pharmaciens", controllers.PharmaciensCreate)

	/* Add 'medecin' */
	r.GET("/medecins", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addDoctor.html", gin.H{})
		//c.Redirect(http.StatusMovedPermanently, "/medecins")
	})
	r.POST("/medecins", controllers.MedecinsCreate)

	/* SIGN UP */
	r.POST("/signup", controllers.SignUp)

	/* LOGIN */
	r.POST("/login", controllers.Login)

	/* LOGIN */
	r.GET("/logout", controllers.Logout)

	/* VALIDATE */
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Run the server
	if r.Run() != nil {
		log.Fatal("Unable to run the server")
		return
	}
}
