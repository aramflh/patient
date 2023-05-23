package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La liste des noms commerciaux de m ́edicaments correspondant `a un nom en DCI, class ́es par ordre alphab ́etique et taille de conditionnement.
*/

func DoRequest1(c *gin.Context) {
	// Get data off requests body => Active user data (Login or JWT)

	type Result struct {
		NomMedicaments string
	}
	var result Result

	initializers.DB.Raw("SELECT nom_medic FROM \"Medicament\" ORDER BY conditionnement;", 3).Scan(&result)

	// Get results
	c.JSON(http.StatusOK, gin.H{
		"result": result.NomMedicaments,
	})
}
