package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Tous les utilisateurs ayant consomm ́e un m ́edicament sp ́ecifique (sous son nom commercial) apr`es une date donn ́ee,
par exemple en cas de rappel de produit pour lot contamin ́e.
*/

func DoRequest4(c *gin.Context) {
	// Get data from the POST request
	var requestData struct {
		MedName string
	}
	// Check if data received
	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "request1.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}

	// Get all the medicamen name from DB
	type listString []string
	var AllMedicaname listString
	initializers.DB.Raw("SELECT nom_medic " +
		"FROM \"Medicament\" ;").Scan(&AllMedicaname)

	// Get the result of the query
	type querryResult []struct {
		Nom    string
		Prenom string
	}
	var result querryResult
	var query string

	query = fmt.Sprintf("SELECT DISTINCT p.nom, p.prenom "+
		"FROM \"Patient\" p "+
		"INNER JOIN \"Prescription\" pr ON p.n_niss = pr.n_niss "+
		"INNER JOIN \"Medicament\" m ON pr.nom_medic = m.nom_medic "+
		"WHERE m.nom_medic = '%s' "+
		"AND pr.date_emission > 'Date_donnee'; ",
		requestData.MedName)

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message":        "",
		"AllMedicaname":  AllMedicaname,
		"currentMedName": requestData.MedName,
		"number":         "4",
		"subject":        "Tous les utilisateurs ayant consommé un médicament spécifique (sous son nom commercial) après une date donnée, par exemple en cas de rappel de produit pour lot contaminé.\n",
		"result":         result,
		"command":        query,
	}

	c.HTML(http.StatusOK, "request4.html", data)

}
