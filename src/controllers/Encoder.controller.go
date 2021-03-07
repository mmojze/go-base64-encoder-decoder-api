package controllers

import (
	b64 "encoding/base64"

	"github.com/gin-gonic/gin"
)

type EncoderRequest struct {
	Value string `json:"value" binding:"required"`
}

/*
@path: /api/v1/encode/base64
@method: POST
*/

func EncodeStringToBase(c *gin.Context) {

	var encoder EncoderRequest

	if err := c.ShouldBindJSON(&encoder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	encodedString := b64.StdEncoding.EncodeToString([]byte(encoder.Value))

	c.JSON(200, gin.H{
		"EncodedValue": encodedString,
	})

}
