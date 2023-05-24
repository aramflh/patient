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

	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "4",
		"subject": "Tous les utilisateurs ayant consommé un médicament spécifique (sous son nom commercial) après une date donnée, par exemple en cas de rappel de produit pour lot contaminé.",
		"result":  result,
	}

	c.HTML(http.StatusOK, "request.html", data)
}
