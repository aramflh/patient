package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La liste des noms commerciaux de ḿedicaments correspondant à un nom en DCI, class ́es par ordre alphabétique et taille de conditionnement.

REQUEST
-------
SELECT nom_medic FROM "Medicament" ORDER BY  conditionnement, nom_medic;
*/

func DoRequest1(c *gin.Context) {
	// Get data from the POST request
	var requestData struct {
		DCI string
	}

	// Check if data received
	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "request1.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}

	// Get all the DCI name from DB
	type listString []string
	var AllDCIname listString
	initializers.DB.Raw("SELECT dci " +
		"FROM \"Medicament\" ;").Scan(&AllDCIname)

	// Get the result of the query
	type Result struct {
		Nom      string `gorm:"column:nom_commercial"`
		Quantite int    `gorm:"column:quantite"`
	}
	var result []Result
	var query string

	query = fmt.Sprintf("SELECT m.nom_commercial, mc.quantite "+
		"FROM \"Medicament\" m "+
		"INNER JOIN \"Medicament_conditionnement\" mc on m.nom_commercial = mc.nom_commercial "+
		"WHERE dci = '%s' "+
		"ORDER BY m.nom_commercial ASC, mc.quantite ASC; ",
		requestData.DCI)

	initializers.DB.Raw(query).Scan(&result)

	fmt.Println(result)

	data := gin.H{
		"message":    "",
		"AllDCIname": AllDCIname,
		"currentDCI": requestData.DCI,
		"number":     "1",
		"subject":    "La liste des noms commerciaux de médicaments correspondant à un nom en DCI, classés par ordre alphabétique et taille de conditionnement.\n",
		"result":     result,
	}
	c.HTML(http.StatusOK, "request1.html", data)

}
