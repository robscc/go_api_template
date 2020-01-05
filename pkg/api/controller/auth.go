package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
