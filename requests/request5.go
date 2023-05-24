package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Tous les patients ayant été traités par un médicament (sous sa DCI) à une date antérieure mais qui ne le sont plus, pour vérifier qu’un patients suive bien un traitement chronique.*/

func DoRequest5(c *gin.Context) {

	// Get the result of the query
	type querryResult []struct {
		Nom    string
		Prenom string
	}
	var result querryResult
	var query string

	query = "SELECT DISTINCT p.nom, p.prenom " +
		"FROM \"Patient\" p " +
		"INNER JOIN \"Traitement\" t ON p.n_niss = t.n_niss " +
		"INNER JOIN \"Medicament\" m ON t.nom_medic = m.nom_medic " +
		"WHERE m.dci = 'Nom_en_DCI' " +
		"AND t.date_debut < CURDATE() " +
		"AND t.date_fin IS NOT NULL;"

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message": "",
		"number":  "5",
		"subject": "Tous les patients ayant été traités par un médicament (sous sa DCI) à une date antérieure mais qui ne le sont plus, pour vérifier qu’un patients suive bien un traitement chronique.\n",
		"result":  result,
		"command": query,
	}

	c.HTML(http.StatusOK, "request5.html", data)
}
