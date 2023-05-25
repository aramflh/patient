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

	/* HOME */
	r.GET("/", controllers.IndexViewer)

	/* SQL requests */
	requestRoutes := r.Group("/requests")

	requestRoutes.GET("/1", requests.DoRequest1)
	requestRoutes.POST("/1", requests.DoRequest1)
	requestRoutes.GET("/2", requests.DoRequest2)
	requestRoutes.GET("/3", requests.DoRequest3)
	requestRoutes.GET("/4", requests.DoRequest4)
	requestRoutes.POST("/4", requests.DoRequest4)
	requestRoutes.GET("/5", requests.DoRequest5)
	requestRoutes.GET("/6", requests.DoRequest6)
	requestRoutes.GET("/7", requests.DoRequest7)
	requestRoutes.GET("/8", requests.DoRequest8)
	requestRoutes.GET("/9", requests.DoRequest9)
	requestRoutes.GET("/10", requests.DoRequest10)

	/* Add 'patient' */
	r.GET("/patients", controllers.AddPatientViewer)
	r.POST("/patients", controllers.AddPatient)

	/* Add 'pharamcien' */
	r.GET("/pharmaciens", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addPharma.html", gin.H{})
	})
	r.POST("/pharmaciens", controllers.PharmaciensCreate)

	/* Add 'medecin' */
	r.GET("/medecins", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addDoctor.html", gin.H{})
	})
	r.POST("/medecins", controllers.MedecinsCreate)

	/* Choose a password for a patient */
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "choosePwd.html", gin.H{})
	})
	r.POST("/signup", controllers.ChoosePwd)

	/* LOGIN */
	r.GET("/login", controllers.LoginViewer)
	r.POST("/login", controllers.Login)

	/* LOGOUT */
	r.GET("/logout", controllers.Logout)

	/* MY ACCOUNT */
	r.GET("/account", middleware.RequireAuth, controllers.ManageAccountViewer)
	r.POST("/account", middleware.RequireAuth, controllers.ManageAccount)

	/* MY TREATEMENTS */
	r.GET("/traitements", middleware.RequireAuth, controllers.TraitementViewer)

	/* MY MED INFO */
	//r.GET("/info-med", middleware.RequireAuth, controllers.InfoMedViewer)

	// Run the server
	if r.Run() != nil {
		log.Fatal("Unable to run the server")
		return
	}
}
