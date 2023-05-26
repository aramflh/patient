package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Tous les patients ayant été traités par un médicament (sous sa DCI) à une date antérieure mais qui ne le sont plus, pour vérifier qu’un patients suive bien un traitement chronique.*/

func DoRequest5(c *gin.Context) {
	// Get data from the POST request
	var requestData struct {
		DCI string
	}

	// Check if data received
	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "request1.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}
	// Get all the DCI name from DB
	type listString []string
	var AllDCIname listString
	initializers.DB.Raw("SELECT dci " +
		"FROM \"Medicament\" ;").Scan(&AllDCIname)

	// Get the result of the query
	type querryResult struct {
		Nom    string `gorm:"column:nom"`
		Prenom string `gorm:"column:prenom"`
	}
	var result []querryResult
	var query string

	query = fmt.Sprintf("SELECT DISTINCT p.nom, p.prenom "+
		"FROM \"Patient\" p "+
		"INNER JOIN \"Prescription\" pr on pr.n_niss = p.n_niss "+
		"INNER JOIN \"Traitement\" t ON t.id_prescription = pr.id "+
		"INNER JOIN \"Medicament\" m ON pr.nom_commercial = m.nom_commercial "+
		"WHERE m.dci = '%s' "+
		"AND t.date_vente < CURRENT_DATE - t.duree_traitement;",
		requestData.DCI)

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message":    "",
		"AllDCIname": AllDCIname,
		"number":     "5",
		"subject":    "Tous les patients ayant été traités par un médicament (sous sa DCI) à une date antérieure mais qui ne le sont plus, pour vérifier qu’un patients suive bien un traitement chronique.\n",
		"result":     result,
		"command":    query,
	}

	c.HTML(http.StatusOK, "request5.html", data)
}
