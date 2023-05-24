package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Quelle est la pathologie la plus diagnostiqu ́ee ?
*/

func DoRequest8(c *gin.Context) {
	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom_medic FROM \"Medicament\" ORDER BY  conditionnement, nom_medic;").Scan(&result)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
