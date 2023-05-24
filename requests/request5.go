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

	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "5",
		"subject": "Tous les patients ayant  été traités par un médicament (sous sa DCI) à une date antérieure mais qui ne le sont plus, pour vérifier qu’un patients suive bien un traitement chronique.",
		"result":  result,
		"command": "SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;",
	}

	c.HTML(http.StatusOK, "request.html", data)
}
