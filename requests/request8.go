package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
Quelle est la pathologie la plus diagnostiqu ́ee ?
*/

func DoRequest8(c *gin.Context) {

	// Get the result of the query
	var result string
	var query string

	query = "SELECT nom_pathologie, COUNT(*) AS total_diagnosis " +
		"FROM \"Dossier_med\" " +
		"GROUP BY nom_pathologie " +
		"ORDER BY total_diagnosis DESC " +
		"LIMIT 1; "

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message": "",
		"number":  "8",
		"subject": "Quelle est la pathologie la plus diagnostiquée ?\n",
		"result":  result,
		"command": query,
	}

	c.HTML(http.StatusOK, "request8.html", data)
}
