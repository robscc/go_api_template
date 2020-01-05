package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping handler implements
// @Summary Ping
// @Description Ping server
// @Tags Ping
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// TestAuth handler implements
// @Summary TestAuth
// @Description TestAuth server
// @Tags Ping
// @Security ApiKeyAuth
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Router /test_auth  [get]
func TestAuth(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
