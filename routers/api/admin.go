package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xyfll7/login/models"
)

// AdminDatabase interface for encapsulating database access.
type AdminDatabase interface {
	InsertAdmin(admin *models.Admin) (*models.Admin, error)
}

// AdminAPI provides handlers for managing admin.
type AdminAPI struct {
	DB AdminDatabase
}

// InsertAdmin creates a admin.
func (a *AdminAPI) InsertAdmin(ctx *gin.Context) {
	var admin = models.Admin{}
	if err := ctx.ShouldBind(&admin); err == nil {
		result, err := a.DB.InsertAdmin(admin.New())
		if err != nil {
			ctx.JSON(203, err)
		}
		ctx.JSON(201, result)
	} else {
		ctx.AbortWithError(500, errors.New("ShouldBind error"))
	}
}
