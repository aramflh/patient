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
	type querryResult struct {
		Decade   string `gorm:"column:decade"`
		NomMedic string `gorm:"column:medicament"`
		Conso    int    `gorm:"column:consommation"`
	}
	var result []querryResult
	var query string

	query = "SELECT DISTINCT ON (decade) CONCAT(EXTRACT(DECADE FROM " +
		"p.date_naissance)*10::INT, ' - ', (EXTRACT(DECADE FROM " +
		"p.date_naissance)*10 + 9)::INT) AS decade,m.nom_commercial AS  " +
		"medicament, COUNT(*) AS consommation " +
		"FROM \"Patient\" p " +
		"INNER JOIN \"Prescription\" pr ON p.n_niss = pr.n_niss " +
		"INNER JOIN \"Medicament\" m ON pr.nom_commercial =  " +
		"m.nom_commercial " +
		"WHERE EXTRACT(YEAR from p.date_naissance) BETWEEN 1950 AND 2019 " +
		"GROUP BY decade,medicament " +
		"ORDER BY decade, consommation desc; "

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
