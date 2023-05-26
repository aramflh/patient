package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La liste de m ́edicament n’ ́etant plus prescrit depuis une date sp ́ecifique.
*/

func DoRequest10(c *gin.Context) {
	// Get data from the POST request
	var requestData struct {
		DateChoice string
	}

	// Check if data received
	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "request1.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}

	// Get the result of the query
	type querryResult struct {
		NomMedic string `gorm:"column:nom_commercial"`
	}
	var result []querryResult
	var query string

	query = fmt.Sprintf("SELECT DISTINCT nom_commercial "+
		"FROM \"Medicament\" "+
		"WHERE nom_commercial NOT IN ( "+
		"SELECT DISTINCT nom_commercial "+
		"FROM \"Prescription\" "+
		"WHERE date_prescription > '%s' );",
		requestData.DateChoice)

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
