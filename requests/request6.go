package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La liste des m ́edecins ayant prescrit des m ́edicaments ne relevant pas de leur sp ́ecialit ́e.
*/

func DoRequest6(c *gin.Context) {

	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "6",
		"subject": "La liste des médecins ayant prescrit des médicaments ne relevant pas de leur spécialité.",
		"result":  result,
	}

	c.HTML(http.StatusOK, "request.html", data)
}
