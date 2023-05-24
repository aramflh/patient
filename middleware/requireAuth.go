package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
)

// RequireAuth checks if the user is authenticated
func RequireAuth(c *gin.Context) {
	// Get the cookie off the request
	tokenString, cookieErr := c.Cookie("PatientAuthorization")

	if cookieErr != nil {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{
			"message": "Connectez-vous pour continuer",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// Decode/Validate the cookie

	// Parse takes the token string and a function for looking up the key.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect: SigningMethodHS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check the cookie expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("Token expired")
			c.HTML(http.StatusUnauthorized, "index.html", gin.H{
				"message": "Session expiré veuillez-vous reconnecter",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with the token sub
		var activePatinentNISS string

		// Transform the value of type []interface to type string (Specific to Go)
		for index, value := range claims {
			if index == "niss" {
				activePatinentNISS = fmt.Sprint(value)
			}
		}

		// Attach to request
		c.Set("activePatientNiss", strings.TrimSpace(activePatinentNISS))

		// User is authenticated continue to next function
		c.Next()
	} else {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{
			"message": "Connectez-vous pour continuer",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
