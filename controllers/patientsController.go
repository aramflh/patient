package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PatientsCreate(c *gin.Context) {
	// Get data off request body

	// Create a patient

	// Return a patient
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
