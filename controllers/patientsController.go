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

// SignUpViewer offers a html views for signing up
func SignUpViewer(c *gin.Context) {
	type Result []string
	var INAMIMedList Result
	var INAMIPhaList Result

	initializers.DB.Raw("SELECT inami FROM \"Medecin\" ;").Scan(&INAMIMedList)
	initializers.DB.Raw("SELECT inami FROM \"Pharmacien\" ;").Scan(&INAMIPhaList)

	data := gin.H{
		"INAMIMedList": INAMIMedList,
		"INAMIPhaList": INAMIPhaList,
	}
	c.HTML(http.StatusOK, "signUp.html", data)
}

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
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}
	// Hash the password
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(requestData.Password), 10)

	// Check for an error
	if hashErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": "Failed to hash password",
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
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": queryErr,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Compte crée avec succès !",
		})
	}
}

// LoginViewer offers a html views for logging in
func LoginViewer(c *gin.Context) {
	data := gin.H{}
	c.HTML(http.StatusOK, "login.html", data)
}

// Login enables the user to connect to its patient account
func Login(c *gin.Context) {
	// Get Email/Password of request body
	var requestData struct {
		Email    string
		Password string
	}

	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}

	// Look up the requested user

	var currentNISS string

	initializers.DB.Raw("SELECT n_niss FROM \"Patient\" WHERE  a_mail = $1;", requestData.Email).Scan(&currentNISS)

	if currentNISS == "" {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": "Invalid email or passsword",
		})
		return
	}

	// Compare sent in pwd with saved hashed pwd

	var hashedPwd string
	initializers.DB.Raw("SELECT pwd FROM \"Patient\" WHERE  a_mail = $1;", requestData.Email).Scan(&hashedPwd)

	hashErr := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(requestData.Password))

	if hashErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": "Invalid email or passsword",
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
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": "Failed to create token !",
		})
		return
	}

	// Set the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("PatientAuthorization", tokenString, COOKIE_AGE, "", "", false, true)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Connecté !",
	})
}

// Logout deletes the cookies and logouts the patient
func Logout(c *gin.Context) {
	cookie, err := c.Cookie("PatientAuthorization")
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "",
		})
		return
	}
	c.SetCookie("PatientAuthorization", cookie, -1, "", "", false, true)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Déconnecté !",
	})
}

// ManageAccount enables the patient to modify medecin and/or pharmacien, and view its treatment and medical information
func ManageAccount(c *gin.Context) {
	// Get the active patient
	activeNiss, _ := c.Get("activePatientNiss")
	// Get patient data request body
	var requestData struct {
		INAMIMed string
		INAMIPha string
	}

	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}
	// Update the user
	querry := fmt.Sprintf("UPDATE \"Patient\" SET n_inami_med = '%s', n_inami_pha = '%s' WHERE n_niss = '%s';",
		requestData.INAMIMed,
		requestData.INAMIPha,
		activeNiss)

	// Executes the query and get error if exist
	queryErr := initializers.DB.Exec(querry).Error

	// Check if an error occurred
	if queryErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": queryErr,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Compte modifié avec succès !",
		})
	}
}

// ManageAccountViewer  offers a html views for the page account
func ManageAccountViewer(c *gin.Context) {

	type Result []string
	var INAMIMedList Result
	var INAMIPhaList Result
	var traitements Result
	var infoMed Result

	initializers.DB.Raw("SELECT inami FROM \"Medecin\" ;").Scan(&INAMIMedList)
	initializers.DB.Raw("SELECT inami FROM \"Pharmacien\" ;").Scan(&INAMIPhaList)

	//TODO: Créer les requetes traitements et informations médicales

	data := gin.H{
		"INAMIMedList": INAMIMedList,
		"INAMIPhaList": INAMIPhaList,
		"traitements":  traitements,
		"infoMed":      infoMed,
	}
	c.HTML(http.StatusOK, "myAccount.html", data)
}
