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
	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;").Scan(&result)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}
