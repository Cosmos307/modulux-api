package middleware

import (
	"context"
	"modulux/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if the user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Validate the token
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store the user information in the context
		c.Set("user_id", claims["sub"])

		c.Next()
	}
}

// Authorize überprüft, ob der Benutzer die erforderlichen Berechtigungen hat
func Authorize(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
			c.Abort()
			return
		}

		modulKuerzel := c.Param("kuerzel")
		modulVersion := c.Param("version")

		if modulKuerzel == "" || modulVersion == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Module identifier missing"})
			c.Abort()
			return
		}

		// SQL query to check if the user has the required permissions
		query := `
            SELECT r.bezeichnung
            FROM modul_person_rolle mpr
            JOIN rolle r ON mpr.rolle_id = r.rolle_id
            JOIN rolle_berechtigung br ON r.rolle_id = br.rolle_id
            JOIN berechtigung b ON br.berechtigung_id = b.berechtigung_id
            WHERE mpr.person_id = $1 AND mpr.modul_kuerzel = $2 AND mpr.modul_version = $3 AND b.bezeichnung = $4
        `

		var role string
		err := database.DB.QueryRow(context.Background(), query, personID, modulKuerzel, modulVersion, requiredPermission).Scan(&role)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
