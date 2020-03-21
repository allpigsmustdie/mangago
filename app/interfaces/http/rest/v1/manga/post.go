package manga

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/allpigsmustdie/mangago/app/domain/models"
)

func (c Controller) Create(context *gin.Context) {
	manga := models.Manga{}

	if err := context.BindJSON(&manga); err != nil {
		logger.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	id, err := c.service.Create(manga)
	if err != nil {
		logger.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "OK", "id": id})
}
