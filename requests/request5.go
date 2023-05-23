package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Tous les patients ayant  ́et ́e trait ́es par un m ́edicament (sous sa DCI) `a une date ant ́erieure mais qui ne le sont plus,
pour v ́erifier qu’un patients suive bien un traitement chronique.*/

func DoRequest5(c *gin.Context) {
	// Get data off requests body => Active user data (Login or JWT)

	type Result struct {
		NomMedicaments string
	}
	var result Result

	initializers.DB.Raw("SELECT nom_medic FROM \"Medicament\" ORDER BY conditionnement;").Scan(&result)
	c.JSON(http.StatusCreated, gin.H{
		"result": result.NomMedicaments,
	})

}
