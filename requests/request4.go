package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Tous les utilisateurs ayant consomm ́e un m ́edicament sp ́ecifique (sous son nom commercial) apr`es une date donn ́ee,
par exemple en cas de rappel de produit pour lot contamin ́e.
*/

func DoRequest4(c *gin.Context) {
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
