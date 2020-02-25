package token

import (
	// "fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller  implement gin context
type Controller struct {
}

// Check check token private claims information.
func (cr *Controller) Check(c *gin.Context) {

	claims := c.MustGet("claims")

	c.JSON(http.StatusOK, gin.H{
		"private claims": claims})
}

// NewController create a new User Controller.
func NewController() *Controller {
	return &Controller{}
}
