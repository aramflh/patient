package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Pour chaque d ́ecennie entre 1950 et 2020, (1950 − 59, 1960 − 69, ...), le m ́edicament le plus consomm ́
e par des patients n ́es durant cette d ́ecennie.*/

func DoRequest7(c *gin.Context) {

	// Get the result of the query
	type querryResult []struct {
		Decade   string
		NomMedic string
	}
	var result querryResult
	var query string

	query = "SELECT SUBSTRING(YEAR(date_naissance), 1, 3) || '0s' AS decade, nom_medic " +
		"FROM \"Patient\" p " +
		"INNER JOIN \"Prescription\" pr ON p.n_niss = pr.n_niss " +
		"INNER JOIN \"Medicament\" m ON pr.nom_medic = m.nom_medic " +
		"WHERE YEAR(date_naissance) BETWEEN 1950 AND 2020 " +
		"GROUP BY decade " +
		"HAVING COUNT(*) > 0 " +
		"ORDER BY decade; "

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message": "",
		"number":  "7",
		"subject": "Pour chaque décennie entre 1950 et 2020, (1950 − 59, 1960 − 69, ...), le médicament le plus consommé par des patients nés durant cette décennie.\n",
		"result":  result,
		"command": query,
	}

	c.HTML(http.StatusOK, "request7.html", data)
}
