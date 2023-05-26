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
	type Result struct {
		Specialite string `gorm:"column:specialite"`
		Quantite   int    `gorm:"column:total_prescriptions"`
	}
	var result []Result
	var query string

	query = "SELECT m.specialite, COUNT(*) AS total_prescriptions " +
		"FROM \"Medecin\" m " +
		"INNER JOIN \"Prescription\" p ON m.inami = p.inami_med " +
		"GROUP BY m.specialite " +
		"ORDER BY total_prescriptions DESC " +
		"LIMIT 1;"

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message": "",
		"number":  "3",
		"subject": "La spécialité de médecins pour laquelle les médecins prescrivent le plus de médicaments.\n",
		"result":  result,
		"command": query,
	}

	c.HTML(http.StatusOK, "request3.html", data)

}
