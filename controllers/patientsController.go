package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"patient/initializers"
)

func SignUp(c *gin.Context) {
	// Get Email/Password of request body
	var requestData struct {
		INSS          string
		Nom           string
		Prenom        string
		Genre         string // TODO: Get only 1 char
		DateNaissance string // TODO: Change to Date
		Email         string
		Password      string
		Num           string
		INAMIMed      string
		INAMIPha      string
	}

	if c.Bind(&requestData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		// Stop
		return
	}
	// Hash the password
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(requestData.Password), 10)

	// Check for an error
	if hashErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		// Stop
		return
	}

	// Create the user
	querry := fmt.Sprintf("INSERT INTO \"Patient\" (n_niss , nom, prenom, genre, date_naissance, a_mail, pwd, n_telephone, n_inami_med, n_inami_pha) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');",
		requestData.INSS,
		requestData.Nom,
		requestData.Prenom,
		requestData.Genre,
		requestData.DateNaissance,
		requestData.Email,
		string(hash),
		requestData.Num,
		requestData.INAMIMed,
		requestData.INAMIPha)

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

/*
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

*/
