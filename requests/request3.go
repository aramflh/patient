package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La sṕecialit ́e de ḿedecins pour laquelle les m ́edecins prescrivent le plus de m ́edicaments.
*/

func DoRequest3(c *gin.Context) {
	// Get data off requests body => Active user data (Login or JWT)

	type Result struct {
		NomMedicaments string
	}
	var result Result

	initializers.DB.Raw("SELECT nom_medic FROM \"Medicament\" ORDER BY conditionnement;").Scan(&result)
	c.JSON(http.StatusCreated, gin.H{
		"result": result.NomMedicaments,
	})

}
