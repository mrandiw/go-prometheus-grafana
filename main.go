package main

import (
	"net/http"
	"tutor/prometheus-go/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	internalServerError = "Internal Server Error"
	notFound            = "Not Found"
)

func main() {
	r := gin.Default()

	middleware.PrometheusInit()

	// Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Middleware to track request metrics
	r.Use(middleware.TrackMetrics())

	// A simple route that increments the request count
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Prometheus!")
	})

	// Another example route
	r.GET("/get-user", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") // Get the query parameter "param" with a default empty value

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Users")
	})

	r.GET("/get-role", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") // Get the query parameter "param" with a default empty value

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Roles")
	})

	r.GET("/get-level", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") // Get the query parameter "param" with a default empty value

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Levels")
	})

	// Start the Gin server
	r.Run("127.0.0.1:8080")
}
