package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

// MedecinsCreate creates a medecin record on the "Medecin" table of the DB
func MedecinsCreate(c *gin.Context) {
	// Get data off requests
	var medData struct {
		INAMI      string
		Nom        string
		Prenom     string
		Email      string
		Num        string
		Specialite string
	}

	if c.Bind(&medData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		// Stop
		return
	}

	// Create a medecin
	querry := fmt.Sprintf("INSERT INTO \"Medecin\" (inami , nom, prenom, a_mail, n_telephone, specialite) VALUES ('%s', '%s', '%s', '%s', '%s', '%s');",
		medData.INAMI,
		medData.Nom,
		medData.Prenom,
		medData.Email,
		medData.Num,
		medData.Specialite)

	// Executes the query and get error if exist
	err := initializers.DB.Exec(querry).Error

	// Check if an error occurred
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("An error occured: %s", err),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "OK",
		})
	}

}
