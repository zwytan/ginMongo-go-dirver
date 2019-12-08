package api

import "github.com/gin-gonic/gin"

import "net/http"

// Test for test
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "test"})
}
