package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La liste de m ́edicament n’ ́etant plus prescrit depuis une date sp ́ecifique.
*/

func DoRequest10(c *gin.Context) {
	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "10",
		"subject": "La liste de médicament n’étant plus prescrit depuis une date spécifique.\n",
		"result":  result,
		"command": "SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;",
	}

	c.HTML(http.StatusOK, "request.html", data)

}
