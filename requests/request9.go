package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Pour chaque patient, le nombre de m ́edecin lui ayant prescrit un m ́edicament.
*/

func DoRequest9(c *gin.Context) {

	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "9",
		"subject": "Pour chaque patient, le nombre de médecin lui ayant prescrit un médicament.\n",
		"result":  result,
		"command": "SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;",
	}

	c.HTML(http.StatusOK, "request.html", data)

}
