package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

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
	// TODO: Change the query
	querry := fmt.Sprintf("INSERT INTO \"Systeme_ana\" (nom_sys_ana, nom_pathologie) VALUES ('%s%s%s', '%s%s');",
		pharma_data.INAMI,
		pharma_data.Nom,
		pharma_data.Prenom,
		pharma_data.Email,
		pharma_data.Num)
	initializers.DB.Exec(querry)

	// Return a patient
	c.JSON(http.StatusOK, gin.H{
		"message": "Pharmacien created !",
	})
}
