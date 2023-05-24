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

	// Get the result of the query
	type querryResult []struct {
		NomMedic string
	}
	var result querryResult
	var query string

	query = "SELECT DISTINCT nom_medic " +
		"FROM \"Medicament\" " +
		"WHERE nom_medic NOT IN ( " +
		"SELECT DISTINCT nom_medic " +
		"FROM \"Prescription\" " +
		"WHERE date_emission > 'Date_specifique' );"

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message": "",
		"number":  "10",
		"subject": "La liste de médicament n’étant plus prescrit depuis une date spécifique.\n",
		"result":  result,
		"command": query,
	}

	c.HTML(http.StatusOK, "request10.html", data)

}
