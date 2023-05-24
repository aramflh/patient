package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La liste des pathologies qui peuvent être prise en charge par un seul type de spécialistes.
*/

func DoRequest2(c *gin.Context) {

	// Get the result of the query
	type querryResult []struct {
		NomPatho  string
		NomSpecia string
	}
	var result querryResult
	var query string

	query = "SELECT p.nom_pathologie, m.specialite " +
		"FROM \"Pathologie\" p " +
		"INNER JOIN \"Medecin\" m ON p.nom_sys_ana = m.nom_sys_ana " +
		"GROUP BY p.nom_pathologie, m.specialite " +
		"HAVING COUNT(DISTINCT m.specialite) = 1;"

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message": "",
		"number":  "2",
		"subject": "La liste des pathologies qui peuvent être prise en charge par un seul type de spécialistes.\n",
		"result":  result,
		"command": query,
	}

	c.HTML(http.StatusOK, "request2.html", data)

}
