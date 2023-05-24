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

	initializers.DB.Raw("SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;").Scan(&result)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}
