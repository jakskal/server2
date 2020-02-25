package storages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller wraps an user service and implement gin context
type Controller struct {
	service Service
}

// ListStorage find loged in user profile.
func (cr *Controller) ListStorage(c *gin.Context) {
	storageList, err := cr.service.StorageList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch data",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, storageList)
}

// ListSizeOption find loged in user profile.
func (cr *Controller) ListSizeOption(c *gin.Context) {
	storageList, err := cr.service.ListSizeOption()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch data",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, storageList)
}

// ListAreaOption find loged in user profile.
func (cr *Controller) ListAreaOption(c *gin.Context) {
	AreaOptions, err := cr.service.ListAreaOption()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch data",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, AreaOptions)
}

// NewController create a new User Controller.
func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}
