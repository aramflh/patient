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

	initializers.DB.Raw("SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;").Scan(&result)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}
