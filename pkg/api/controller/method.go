package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHandler handler implements
// @Summary GetHandler
// @Description method sample get
// @Tags Method
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Router /sample/get_handler [get]
func GetHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// An PostSampleParams for PostHandler params
type PostSampleParams struct {
	SampleKey string `json:"sample_Key" binding:"required"` // sample key is sample and must be valued when post to server
}

// PostHandler handler implements
// @Summary PostHandler
// @Description method sample post
// @Tags Method
// @Accept application/json
// @Param body body controller.PostSampleParams true "PostSampleParams the sample_key is required"
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Router /sample/post_handler [post]
func PostHandler(c *gin.Context) {
	var params PostSampleParams
	if err := c.BindJSON(&params); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("sample key:%s", params.SampleKey))
}

// PostFileHandler handler implements
// @Summary PostFileHandler
// @Description method sample post file
// @Accept multipart/form-data
// @Param file formData file true "the file to upload"
// @Tags Method
// @Produce text/plain
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "bad request"
// @Router /sample/post_file_handler [post]
func PostFileHandler(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("sample file name:%s", fh.Filename))
}

// An PutSampleParams for PostHandler params
type PutSampleParams struct {
	SampleKey string `json:"sample_Key" binding:"required"` // sample key is sample and must be valued when post to server
}

// PutHandler handler implements
// @Summary PutHandler
// @Description method sample put
// @Tags Method
// @Accept application/json
// @Param body body controller.PutSampleParams true "PutSampleParams the sample_key is required"
// @Produce text/plain
// @Success 200 {string} string	"ok"
// Failure 400 {string} string	"bad request"
// @Router /sample/put_handler[put]
func PutHandler(c *gin.Context) {
	var params PutSampleParams
	if err := c.BindJSON(&params); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("sample key:%s", params.SampleKey))
}

// DeleteHandler handler implements
// @Summary DeleteHandler
// @Description method sample delete
// @Tags Method
// @Success 204
// @Router /sample/delete_handler [delete]
func DeleteHandler(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
