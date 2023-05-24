package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La spécialité de médecins pour laquelle les médecins prescrivent le plus de médicaments.
*/

func DoRequest3(c *gin.Context) {

	// Get the result of the query
	var specialite string
	var query string

	query = "SELECT m.specialite, COUNT(*) AS total_prescriptions " +
		"FROM \"Medecin\" m " +
		"INNER JOIN \"Prescription\" p ON m.n_inami_med = p.n_inami_med " +
		"GROUP BY m.specialite " +
		"ORDER BY total_prescriptions DESC " +
		"LIMIT 1;"

	initializers.DB.Raw(query).Scan(&specialite)

	data := gin.H{
		"message": "",
		"number":  "3",
		"subject": "La spécialité de médecins pour laquelle les médecins prescrivent le plus de médicaments.\n",
		"result":  specialite,
		"command": query,
	}

	c.HTML(http.StatusOK, "request3.html", data)

}
