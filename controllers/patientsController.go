package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"patient/initializers"
	"time"
)

const COOKIE_AGE int = 3600 * 24 * 30 // Cookie expire after 1 month
const JWT_AGE = time.Hour * 24 * 30   // JWT token expire after 1 month

// SignUp enables the user to create a patient account
func SignUp(c *gin.Context) {
	// Get patient data request body
	var requestData struct {
		NISS          string
		Nom           string
		Prenom        string
		Genre         string
		DateNaissance string
		Email         string
		Password      string
		Num           string
		INAMIMed      string
		INAMIPha      string
	}

	if c.Bind(&requestData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		// Stop
		return
	}
	// Hash the password
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(requestData.Password), 10)

	// Check for an error
	if hashErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		// Stop
		return
	}

	// Create the user
	querry := fmt.Sprintf("INSERT INTO \"Patient\" (n_niss , nom, prenom, genre, date_naissance, a_mail, pwd, n_telephone, n_inami_med, n_inami_pha) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');",
		requestData.NISS,
		requestData.Nom,
		requestData.Prenom,
		requestData.Genre,
		requestData.DateNaissance,
		requestData.Email,
		string(hash),
		requestData.Num,
		requestData.INAMIMed,
		requestData.INAMIPha)

	// Executes the query and get error if exist
	queryErr := initializers.DB.Exec(querry).Error

	// Check if an error occurred
	if queryErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("An error occured: %s", queryErr),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "OK",
		})
	}
}

// Login enables the user to connect to its patient account
func Login(c *gin.Context) {
	// Get Email/Password of request body
	var requestData struct {
		Email    string
		Password string
	}

	if c.Bind(&requestData) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		// Stop
		return
	}

	// Look up the requested user

	var currentNISS string

	initializers.DB.Raw("SELECT n_niss FROM \"Patient\" WHERE  a_mail = $1;", requestData.Email).Scan(&currentNISS)

	if currentNISS == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or passsword",
		})
		return
	}

	// Compare sent in pwd with saved hashed pwd

	var hashedPwd string
	initializers.DB.Raw("SELECT pwd FROM \"Patient\" WHERE  a_mail = $1;", requestData.Email).Scan(&hashedPwd)

	hashErr := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(requestData.Password))

	if hashErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or passsword",
		})
		return
	}

	// At this point the passwords match
	// Generate a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"niss": currentNISS,                    // Subjec of the token
		"exp":  time.Now().Add(JWT_AGE).Unix(), // Expiration of the token: 1 month
	})

	tokenString, tokenErr := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if tokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Set the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, COOKIE_AGE, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

// Logout deletes the cookies and logouts the patient
func Logout(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to disable cookie",
		})
		return
	}
	c.SetCookie("Authorization", cookie, -1, "", "", false, true)
}

func Validate(c *gin.Context) {
	patient, _ := c.Get("activePatientNiss")
	c.JSON(http.StatusOK, gin.H{
		"message": patient,
	})
}
