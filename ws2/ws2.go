package ws2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: username`
	Password string `json: password`
}

var router *gin.Engine

func Iniciar(router1 *gin.Engine) {
	router = router1
	initializeRoutes()

}

func initializeRoutes() {
	router.POST("/api2", handleVerification)
	router.OPTIONS("/api2", handleVerification)
	router.GET("/api2", handleGet)
}

func handleGet(c *gin.Context) {
	message, _ := c.GetQuery("m")
	c.String(http.StatusOK, "Get works! you sent: "+message)
}

func handleVerification(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		// setup headers
		c.Header("Allow", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
		c.Header("Content-Type", "application/json")
		c.Status(http.StatusOK)
	} else if c.Request.Method == "POST" {
		var u User
		c.BindJSON(&u)
		c.JSON(http.StatusOK, gin.H{
			"user": u.Username,
			"pass": u.Password,
		})
	}
}
