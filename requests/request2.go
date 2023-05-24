package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La liste des pathologies qui peuvent être prise en charge par un seul type de spécialistes.
*/

func DoRequest2(c *gin.Context) {

	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "2",
		"subject": "La liste des pathologies qui peuvent être prise en charge par un seul type de spécialistes.\n",
		"result":  result,
		"command": "SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;",
	}

	c.HTML(http.StatusOK, "request.html", data)

}
