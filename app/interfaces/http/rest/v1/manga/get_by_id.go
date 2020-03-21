package manga

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/allpigsmustdie/mangago/app/domain/repository"
)

func (c Controller) GetById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		logger.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	manga, err := c.service.Get(id)

	if err == repository.ErrNotFound {
		context.JSON(http.StatusNotFound, gin.H{"status": "NotFound"})
		return
	}

	if err != nil {
		logger.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "OK", "manga": manga})
}
