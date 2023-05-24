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
	type querryResult []struct {
		INAMI      string
		Specialite string
	}
	var result querryResult
	var query string

	query = "SELECT DISTINCT m.n_inami_med, m.specialite " +
		"FROM \"Medecin\" m " +
		"INNER JOIN \"Prescription\" p ON m.n_inami_med = p.n_inami_med " +
		"INNER JOIN \"Medicament\" med ON p.nom_medic = med.nom_medic " +
		"WHERE med.nom_pathologie NOT IN (SELECT nom_pathologie " +
		"FROM \"Pathologie\" " +
		"WHERE nom_sys_ana = m.nom_sys_ana);"

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
