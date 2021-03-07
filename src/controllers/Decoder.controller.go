package controllers

import (
	b64 "encoding/base64"

	"github.com/gin-gonic/gin"
)

type DecoderRequest struct {
	Value string `json:"value" binding:"required"`
}

/*
@path: /api/v1/decode/base64
@method: POST
*/

func DecodeBaseToString(c *gin.Context) {

	var decoder DecoderRequest

	if err := c.ShouldBindJSON(&decoder); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	decodedString, err := b64.StdEncoding.DecodeString(decoder.Value)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"DecodedValue": string(decodedString),
	})

}
