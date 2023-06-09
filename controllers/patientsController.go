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

const COOKIE_AGE int = 3600 // Cookie expire after 1 hour
const JWT_AGE = time.Hour   // JWT token expire after 1 hour

func IndexViewer(c *gin.Context) {
	var isConnected bool

	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}
	data := gin.H{
		"message":     "",
		"isConnected": isConnected,
	}
	c.HTML(http.StatusOK, "index.html", data)
}

func AddPatientViewer(c *gin.Context) {
	type Result []string
	var INAMIMedList Result
	var INAMIPhaList Result
	var isConnected bool

	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}

	initializers.DB.Raw("SELECT inami FROM \"Medecin\" ;").Scan(&INAMIMedList)
	initializers.DB.Raw("SELECT inami FROM \"Pharmacien\" ;").Scan(&INAMIPhaList)

	data := gin.H{
		"INAMIMedList": INAMIMedList,
		"INAMIPhaList": INAMIPhaList,
		"isConnected":  isConnected,
	}
	c.HTML(http.StatusOK, "addPatient.html", data)
}

// AddPatient enables the user to create a patient account
func AddPatient(c *gin.Context) {
	var isConnected bool
	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}
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
			"message":     "Failed to read request",
			"isConnected": isConnected,
		})
		// Stop
		return
	}
	// Hash the password
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(requestData.Password), 10)

	// Check for an error
	if hashErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     "Failed to hash password",
			"isConnected": isConnected,
		})
		// Stop
		return
	}

	// Create the user
	querry := fmt.Sprintf("INSERT INTO \"Patient\" (n_niss , nom, prenom, genre, date_naissance, a_mail, pwd, n_telephone, inami, n_inami_pha) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');",
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
			"message":     queryErr,
			"isConnected": isConnected,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":     "Compte crée avec succès !",
			"isConnected": isConnected,
		})
	}
}

// ChoosePwd enable a patient to link a pwd to its niss number
func ChoosePwd(c *gin.Context) {
	// Get data off requests
	var patientData struct {
		NISS     string
		Password string
	}
	if c.Bind(&patientData) != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "Failed to read request",
		})
		// Stop
		return
	}

	// Check if the session cookie exist
	var isConnected bool

	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}

	// Hash the password
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(patientData.Password), 10)

	// Check for an error
	if hashErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     "Failed to hash password",
			"isConnected": isConnected,
		})
		// Stop
		return
	}

	// Change the pwd for the selected niss
	querry := fmt.Sprintf("UPDATE \"Patient\" SET pwd = '%s' WHERE n_niss = '%s';",
		string(hash),
		patientData.NISS)

	// Executes the query and get error if exist
	queryErr := initializers.DB.Exec(querry).Error

	// Check if an error occurred
	if queryErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     queryErr,
			"isConnected": isConnected,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":     "Mot de passe associé avec succès",
			"isConnected": isConnected,
		})
	}
}

func LoginViewer(c *gin.Context) {
	var isConnected bool

	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}
	data := gin.H{
		"isConnected": isConnected,
	}
	c.HTML(http.StatusOK, "login.html", data)
}

// Login enables the user to connect to its patient account
func Login(c *gin.Context) {
	var isConnected bool
	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}
	// Get Email/Password of request body
	var requestData struct {
		NISS     string
		Password string
	}

	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     "Failed to read request",
			"isConnected": isConnected,
		})
		// Stop
		return
	}

	// Look up the requested user

	var currentNISS string

	initializers.DB.Raw("SELECT n_niss FROM \"Patient\" WHERE  n_niss = $1;", requestData.NISS).Scan(&currentNISS)

	if currentNISS == "" {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     "Invalid NISS or passsword",
			"isConnected": isConnected,
		})
		return
	}

	// Compare sent in pwd with saved hashed pwd

	var hashedPwd string
	initializers.DB.Raw("SELECT pwd FROM \"Patient\" WHERE  n_niss = $1;", requestData.NISS).Scan(&hashedPwd)

	hashErr := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(requestData.Password))

	if hashErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     "Invalid NISS or passsword",
			"isConnected": isConnected,
		})
		return
	}

	// At this point the passwords match
	// Generate a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"niss": currentNISS,                    // Subjec of the token
		"exp":  time.Now().Add(JWT_AGE).Unix(), // Expiration of the token
	})

	tokenString, tokenErr := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if tokenErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     "Failed to create token !",
			"isConnected": isConnected,
		})
		return
	}

	// Set the cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("PatientAuthorization", tokenString, COOKIE_AGE, "", "", false, true)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message":     "Connecté !",
		"isConnected": isConnected,
	})
}

// Logout deletes the cookies and logouts the patient
func Logout(c *gin.Context) {
	var isConnected bool
	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}
	cookie, err := c.Cookie("PatientAuthorization")
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":     "Déja déconnecté !",
			"isConnected": isConnected,
		})
		return
	}
	// Deleting the cookie
	c.SetCookie("PatientAuthorization", cookie, -1, "", "", false, true)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message":     "Déconnecté !",
		"isConnected": isConnected,
	})
}

// ManageAccountViewer  offers a html views for the page account
func ManageAccountViewer(c *gin.Context) {

	type listString []string
	var INAMIMedList listString
	var INAMIPhaList listString

	initializers.DB.Raw("SELECT inami FROM \"Medecin\" ;").Scan(&INAMIMedList)
	initializers.DB.Raw("SELECT inami FROM \"Pharmacien\" ;").Scan(&INAMIPhaList)

	data := gin.H{
		"INAMIMedList": INAMIMedList,
		"INAMIPhaList": INAMIPhaList,
	}
	c.HTML(http.StatusOK, "myAccount.html", data)
}

func TraitementViewer(c *gin.Context) {
	activeNiss, _ := c.Get("activePatientNiss")

	// Get the result of the query
	type querryResult struct {
		DateDebut string `gorm:"column:date_vente"`
		Duree     string `gorm:"column:duree_traitement"`
		NomMedic  string `gorm:"column:nom_commercial"`
		NomMed    string `gorm:"column:nom"`
		INAMIMed  string `gorm:"column:inami"`
	}
	var traitements []querryResult
	var query string

	query = fmt.Sprintf("SELECT CONCAT(EXTRACT(DAY FROM t.date_vente), '-', (EXTRACT(MONTH FROM t.date_vente)) , '-', (EXTRACT(YEAR FROM t.date_vente)) ) as date_vente, t.duree_traitement as duree_traitement, p.nom_commercial as nom_commercial, m.nom as nom, p.inami_med as inami  FROM \"Traitement\" t, \"Prescription\" p, \"Medecin\" m WHERE p.id=t.id_prescription AND p.inami_med=m.inami AND p.n_niss = '%s';",
		activeNiss)

	initializers.DB.Raw(query).Scan(&traitements)

	data := gin.H{
		"message": "",
		"result":  traitements,
	}

	c.HTML(http.StatusOK, "traitements.html", data)
}

func InfoMedViewer(c *gin.Context) {
	activeNiss, _ := c.Get("activePatientNiss")

	// Get the result of the query
	type querryResult struct {
		NISS          string `gorm:"column:niss"`
		Nom           string `gorm:"column:nom"`
		Prenom        string `gorm:"column:prenom"`
		Genre         string `gorm:"column:genre"`
		DateNaissance string `gorm:"column:date_naissance"`
		Mail          string `gorm:"column:mail"`
		Tel           string `gorm:"column:num"`
		INAMIMed      string `gorm:"column:inami_med"`
		NomMed        string `gorm:"column:med_nom"`
		TelMed        string `gorm:"column:med_tel"`
		MailMed       string `gorm:"column:med_mail"`
		INAMIPha      string `gorm:"column:inami_pha"`
		NomPha        string `gorm:"column:pha_nom"`
		TelPha        string `gorm:"column:pha_tel"`
		MailPha       string `gorm:"column:pha_mail"`
	}
	var info []querryResult
	var query string

	query = fmt.Sprintf("SELECT p.n_niss as niss, p.nom as nom, p.prenom as prenom, p.genre as genre, "+
		" CONCAT(EXTRACT(DAY FROM p.date_naissance), '-', (EXTRACT(MONTH FROM p.date_naissance)) , '-', (EXTRACT(YEAR FROM p.date_naissance)) ) "+
		"as date_naissance, p.a_mail as mail, p.n_telephone as num, p.inami as inami_med, p.n_inami_pha as inami_pha, "+
		"m.nom as med_nom, m.n_telephone as med_tel, m.a_mail as med_mail, ph.nom as pha_nom, ph.n_telephone as pha_tel, "+
		"ph.a_mail as pha_mail FROM \"Patient\" p, \"Medecin\" m, \"Pharmacien\" ph WHERE p.inami=m.inami AND "+
		"p.n_inami_pha=ph.inami AND  p.n_niss = '%s';",
		activeNiss)

	initializers.DB.Raw(query).Scan(&info)

	data := gin.H{
		"message": "",
		"result":  info,
	}

	c.HTML(http.StatusOK, "infoMed.html", data)
}

// UpdateMed enables to changes doctor inami
func UpdateMed(c *gin.Context) {
	var isConnected bool
	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}
	// Get the active patient
	activeNiss, _ := c.Get("activePatientNiss")
	// Get patient data request body
	var requestData struct {
		INAMIMed string
	}

	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     "Failed to read request",
			"isConnected": isConnected,
		})
		// Stop
		return
	}
	// Update the user
	querry := fmt.Sprintf("UPDATE \"Patient\" SET inami = '%s' WHERE n_niss = '%s';",
		requestData.INAMIMed,
		activeNiss)

	// Executes the query and get error if exist
	queryErr := initializers.DB.Exec(querry).Error

	// Check if an error occurred
	if queryErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     queryErr,
			"isConnected": isConnected,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":     "Compte modifié avec succès !",
			"isConnected": isConnected,
		})
	}
}

func UpdateMedViewer(c *gin.Context) {
	type Result []string
	var INAMIMedList Result
	var isConnected bool
	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}

	initializers.DB.Raw("SELECT inami FROM \"Medecin\" ;").Scan(&INAMIMedList)
	data := gin.H{
		"INAMIMedList": INAMIMedList,
		"isConnected":  isConnected,
	}
	c.HTML(http.StatusOK, "updateMed.html", data)

}

// UpdateMed enables to changes doctor inami
func UpdatePha(c *gin.Context) {
	var isConnected bool
	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}
	// Get the active patient
	activeNiss, _ := c.Get("activePatientNiss")
	// Get patient data request body
	var requestData struct {
		INAMIPha string
	}

	if c.Bind(&requestData) != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     "Failed to read request",
			"isConnected": isConnected,
		})
		// Stop
		return
	}
	// Update the user
	querry := fmt.Sprintf("UPDATE \"Patient\" SET n_inami_pha = '%s' WHERE n_niss = '%s';",
		requestData.INAMIPha,
		activeNiss)

	// Executes the query and get error if exist
	queryErr := initializers.DB.Exec(querry).Error

	// Check if an error occurred
	if queryErr != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message":     queryErr,
			"isConnected": isConnected,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":     "Compte modifié avec succès !",
			"isConnected": isConnected,
		})
	}
}

func UpdatePhaViewer(c *gin.Context) {
	type Result []string
	var INAMIPhaList Result
	var isConnected bool
	// Check if the session cookie exist
	_, cookieErr := c.Cookie("PatientAuthorization")
	if cookieErr != nil {
		isConnected = false
	} else {
		isConnected = true
	}

	initializers.DB.Raw("SELECT inami FROM \"Pharmacien\" ;").Scan(&INAMIPhaList)
	data := gin.H{
		"INAMIPhaList": INAMIPhaList,
		"isConnected":  isConnected,
	}
	c.HTML(http.StatusOK, "updatePha.html", data)

}
