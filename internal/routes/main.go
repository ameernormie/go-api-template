package routes

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"time"

	"github.com/ameernormie/go-api-template/internal/services/encryptionservice"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Format for time (iso input format for javascript ISO string)
const (
	INPUTFORMAT = "2006-01-02T15:04:05.000Z"
)

// Authorization checks the token in every request (auth middleware)
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			unauthorized(c)
			c.Abort()
			return
		}

		decryptedToken, err := encryptionservice.Decrypt(token)
		if err != nil {
			unauthorized(c)
			c.Abort()
			return
		}

		t, err := time.Parse(INPUTFORMAT, decryptedToken)
		if err != nil {
			unauthorized(c)
			c.Abort()
			return
		}
		currentTime := time.Now()
		diff := math.Abs(currentTime.Sub(t).Minutes())

		if diff > 3 {
			unauthorized(c)
			c.Abort()
			return
		}
		c.Next()
	}
}

// GetRouter : return gin instance with routes
func GetRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))
	api := router.Group("/api")
	// api.Use(Authorization())
	{
		api.GET("/ping", ping)
	}
	return router
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "not found",
	})
}

func ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func internalServerError(c *gin.Context, err error, message string) {
	if len(message) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("%s: %v", message, err),
		})

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("internal server error: %v", err),
		})
	}
}

func unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": fmt.Sprintf("Unauthorized - Invalid token"),
	})
}

// GetTestToken : Get test token for apis
func GetTestToken() (string, error) {
	currentTime := time.Now()
	parsedTime := currentTime.UTC().Format(INPUTFORMAT)
	token, err := encryptionservice.Encrypt(parsedTime)

	if err != nil {
		return "", err
	}
	encodedURL := url.QueryEscape(token)
	return encodedURL, nil
}
