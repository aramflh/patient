package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Pour chaque patient, le nombre de m ́edecin lui ayant prescrit un m ́edicament.
*/

func DoRequest9(c *gin.Context) {

	// Get the result of the query
	type querryResult struct {
		NISS   string `gorm:"column:n_niss"`
		Nom    string `gorm:"column:nom"`
		Prenom string `gorm:"column:prenom"`
		NbrMed int    `gorm:"column:nombre_medecins"`
	}
	var result []querryResult
	var query string

	query = "SELECT p.n_niss, p.nom, p.prenom, COUNT(DISTINCT pr.inami_med) AS nombre_medecins " +
		"FROM \"Patient\" p " +
		"INNER JOIN \"Prescription\" pr ON p.n_niss = pr.n_niss " +
		"GROUP BY p.n_niss, p.nom, p.prenom; "

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message": "",
		"number":  "9",
		"subject": "Pour chaque patient, le nombre de médecin lui ayant prescrit un médicament.\n",
		"result":  result,
		"command": query,
	}

	c.HTML(http.StatusOK, "request9.html", data)

}
