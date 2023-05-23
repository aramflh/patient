package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

func PatientsCreate(c *gin.Context) {
	// Get data off requests body
	// TODO
	var patient_data struct {
		INSS          string
		Nom           string
		Prenom        string
		Genre         string
		DateNaissance string // TODO: Check correct type for time postgres type
		Email         string
		Num           string
		INAMIMed      string // CONSTRAINT fk_n_inami_med FOREIGN KEY (n_inami_med) REFERENCES "Medecin"(n_inami_med)
		INAMPha       string // CONSTRAINT fk_n_inami_pha FOREIGN KEY (n_inami_pha) REFERENCES "Pharmacien"(n_inami_pha)
	}

	c.Bind(&patient_data)
	// Create a patient
	// TODO: Change the query
	querry := fmt.Sprintf("INSERT INTO \"Systeme_ana\" (nom_sys_ana, nom_pathologie) VALUES ('%s%s%s%s', '%s%s%s%s%s');",
		patient_data.INSS,
		patient_data.Nom,
		patient_data.Prenom,
		patient_data.Genre,
		patient_data.DateNaissance,
		patient_data.Email,
		patient_data.Num,
		patient_data.INAMIMed,
		patient_data.INAMPha)
	initializers.DB.Exec(querry)

	// Return a patient
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
