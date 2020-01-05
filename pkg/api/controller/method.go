package controller

import (
	"github.com/gin-gonic/gin"
)

// GetHandler handler implements
// @Summary GetHandler
// @Description Ping server
// @Tags Ping
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Router /ping [get]
func GetHandler(c *gin.Context) {}

// PostHandler handler implements
// @Summary PostHandler
// @Description Ping server
// @Tags Ping
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Router /ping [post]
func PostHandler(c *gin.Context) {}

// PostFileHandler handler implements
// @Summary PostFileHandler
// @Description Ping server
// @Tags Ping
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Router /ping [post]
func PostFileHandler(c *gin.Context) {}

// PutHandler handler implements
// @Summary PutHandler
// @Description Ping server
// @Tags Ping
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Router /ping [put]
func PutHandler(c *gin.Context) {}

// HeaderHandler handler implements
// @Summary HeaderHandler
// @Description Ping server
// @Tags Ping
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Router /ping [header]
func HeaderHandler(c *gin.Context) {}

// DeleteHandler handler implements
// @Summary DeleteHandler
// @Description Ping server
// @Tags Ping
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Router /ping [delete]
func DeleteHandler(c *gin.Context) {}
