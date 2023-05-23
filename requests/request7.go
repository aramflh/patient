package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Pour chaque d ́ecennie entre 1950 et 2020, (1950 − 59, 1960 − 69, ...), le m ́edicament le plus consomm ́
e par des patients n ́es durant cette d ́ecennie.*/

func DoRequest7(c *gin.Context) {
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
