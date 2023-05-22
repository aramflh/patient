package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

func MedecinsCreate(c *gin.Context) {
	// Get data off request
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

	// Create a pharmacien
	querry := fmt.Sprintf("INSERT INTO \"Systeme_ana\" (nom_sys_ana, nom_pathologie) VALUES ('%s%s%s', '%s%s%s%s');",
		med_data.INAMI,
		med_data.Nom,
		med_data.Prenom,
		med_data.Email,
		med_data.Num,
		med_data.Specialite,
		med_data.SystemAna)
	initializers.DB.Exec(querry)

	// Return a patient
	c.JSON(http.StatusOK, gin.H{
		"message": "Medecins created !",
	})
}
