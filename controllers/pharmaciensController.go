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
	var pharma_data struct {
		INAMI  string
		Nom    string
		Prenom string
		Email  string
		Num    string
	}

	c.Bind(&pharma_data)

	// Create a pharmacien
	querry := fmt.Sprintf("INSERT INTO \"Pharmacien\" (inami, nom, prenom, a_mail, n_telephone) VALUES ('%s', '%s', '%s', '%s', '%s');",
		pharma_data.INAMI,
		pharma_data.Nom,
		pharma_data.Prenom,
		pharma_data.Email,
		pharma_data.Num)

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
