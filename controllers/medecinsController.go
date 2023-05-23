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
	var med_data struct {
		INAMI      string
		Nom        string
		Prenom     string
		Email      string
		Num        string
		Specialite string
		SystemAna  string
	}

	c.Bind(&med_data)

	// Create a medecin
	// TODO: Change the query
	querry := fmt.Sprintf("INSERT INTO \"Systeme_ana\" (nom_sys_ana, nom_pathologie) VALUES ('%s%s%s', '%s%s%s%s');",
		med_data.INAMI,
		med_data.Nom,
		med_data.Prenom,
		med_data.Email,
		med_data.Num,
		med_data.Specialite,
		med_data.SystemAna)

	// Executes the query and get error if exist
	err := initializers.DB.Exec(querry).Error

	// Check if an error occurred
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("An error occured: %s", err),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Medecins created !",
		})
	}

}
