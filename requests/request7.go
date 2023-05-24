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

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "7",
		"subject": "Pour chaque décennie entre 1950 et 2020, (1950 − 59, 1960 − 69, ...), le médicament le plus consommé par des patients nés durant cette décennie.",
		"result":  result,
		"command": "SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;",
	}

	c.HTML(http.StatusOK, "request.html", data)
}
