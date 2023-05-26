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
	type querryResult struct {
		NomPatho  string `gorm:"column:nom_pathologie"`
		NomSpecia string `gorm:"column:specialite"`
	}
	var result []querryResult
	var query string

	query = "SELECT p.nom_pathologie, s.specialite " +
		"FROM \"Pathologie\" p " +
		"INNER JOIN \"Pathologie_specialite\" s ON p.nom_pathologie = s.nom_pathologie " +
		"GROUP BY p.nom_pathologie, s.specialite " +
		"HAVING COUNT(DISTINCT s.specialite) = 1;"

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
