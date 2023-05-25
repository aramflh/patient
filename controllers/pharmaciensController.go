package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

// PharmaciensCreate creates a pharmacien record on the "Pharmacien" table of the DB
func PharmaciensCreate(c *gin.Context) {
	// Get data off requests
	var pharmaData struct {
		INAMI string
		Nom   string
		Email string
		Num   string
	}

	if c.Bind(&pharmaData) != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}

	// Create a pharmacien
	querry := fmt.Sprintf("INSERT INTO \"Pharmacien\" (inami, nom, a_mail, n_telephone) VALUES ( '%s', '%s', '%s', '%s');",
		pharmaData.INAMI,
		pharmaData.Nom,
		pharmaData.Email,
		pharmaData.Num)

	// Executes the query and get error if exist
	err := initializers.DB.Exec(querry).Error

	// Check if an error occurred
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": err,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Pharmacien ajouté avec succès !",
		})
	}
}
