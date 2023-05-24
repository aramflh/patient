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

	type Result []string
	var result Result

	initializers.DB.Raw("SELECT nom FROM \"Medecin\" ;").Scan(&result)

	data := gin.H{
		"number":  "8",
		"subject": "Quelle est la pathologie la plus diagnostiquée ?\n",
		"result":  result,
	}

	c.HTML(http.StatusOK, "request.html", data)
}
