package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La liste des noms commerciaux de ḿedicaments correspondant à un nom en DCI, class ́es par ordre alphabétique et taille de conditionnement.

REQUEST
-------
SELECT nom_medic FROM "Medicament" ORDER BY  conditionnement, nom_medic;
*/

func DoRequest1(c *gin.Context) {

	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "1",
		"subject": "La liste des noms commerciaux de médicaments correspondant à un nom en DCI, classés par ordre alphabétique et taille de conditionnement.\n",
		"result":  result,
		"command": "SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;",
	}

	c.HTML(http.StatusOK, "request.html", data)

}
