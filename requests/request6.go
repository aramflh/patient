package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La liste des médecins ayant prescrit des médicaments ne relevant pas de leur spécialité.
*/

func DoRequest6(c *gin.Context) {

	// Get the result of the query
	type querryResult struct {
		Nom   string `gorm:"column:nom"`
		INAMI string `gorm:"column:inami"`
	}
	var result []querryResult
	var query string

	query = "SELECT DISTINCT m.nom, m.inami " +
		"FROM \"Medecin\" m " +
		"INNER JOIN \"Prescription\" pr ON m.inami = pr.inami_med " +
		"INNER JOIN \"Medicament\" med ON pr.nom_commercial = med.nom_commercial " +
		"WHERE med.nom_sys_ana NOT IN ( " +
		"SELECT ss.nom_sys_ana " +
		"FROM \"Specialite_systeme_ana\" ss " +
		"WHERE ss.specialite = m.specialite );"

	initializers.DB.Raw(query).Scan(&result)

	data := gin.H{
		"message": "",
		"number":  "6",
		"subject": "La liste des médecins ayant prescrit des médicaments ne relevant pas de leur spécialité.\n",
		"result":  result,
		"command": query,
	}

	c.HTML(http.StatusOK, "request6.html", data)
}
