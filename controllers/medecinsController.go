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
		Email      string
		Num        string
		Specialite string
	}

	if c.Bind(&medData) != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}

	// Create a medecin
	querry := fmt.Sprintf("INSERT INTO \"Medecin\" (inami , nom, a_mail, n_telephone, specialite) VALUES ('%s', '%s', '%s', '%s', '%s');",
		medData.INAMI,
		medData.Nom,
		medData.Email,
		medData.Num,
		medData.Specialite)

	// Executes the query and get error if exist
	err := initializers.DB.Exec(querry).Error

	// Check if an error occurred
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": err,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Médecin ajouté avec succès !",
		})
	}

}
