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
		DateMed string
	}
	// Check if data received
	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}
	fmt.Println(requestData)

	// Get all the medicamen name from DB
	type listString []string
	var AllMedicaname listString
	initializers.DB.Raw("SELECT nom_commercial " +
		"FROM \"Medicament\" ;").Scan(&AllMedicaname)

	// Get the result of the query
	type querryResult struct {
		Nom    string `gorm:"column:nom"`
		Prenom string `gorm:"column:prenom"`
	}
	var result []querryResult
	var query string

	query = fmt.Sprintf("SELECT DISTINCT p.nom, p.prenom "+
		"FROM \"Patient\" p "+
		"INNER JOIN \"Prescription\" pr ON p.n_niss = pr.n_niss "+
		"INNER JOIN \"Medicament\" m ON pr.nom_commercial = m.nom_commercial "+
		"INNER JOIN \"Traitement\" t ON t.id_prescription = pr.id "+
		"WHERE m.nom_commercial = '%s' "+
		"AND t.date_vente > '%s';",
		requestData.MedName,
		requestData.DateMed)

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
