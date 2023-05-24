package requests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"patient/initializers"
)

/*
La sṕecialit ́e de ḿedecins pour laquelle les m ́edecins prescrivent le plus de m ́edicaments.
*/

func DoRequest3(c *gin.Context) {

	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "3",
		"subject": "La spécialité de médecins pour laquelle les médecins prescrivent le plus de médicaments.\n",
		"result":  result,
	}

	c.HTML(http.StatusOK, "request.html", data)

}
