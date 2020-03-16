package v1

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/allpigsmustdie/mangago/app/domain/models"
	"github.com/allpigsmustdie/mangago/app/domain/repository"
	"github.com/allpigsmustdie/mangago/app/domain/service"
)

//TODO: DI logger
var logger = log.New(os.Stdout, "[manga router]", log.Lshortfile | log.Ltime)

type MangaController struct {
	service service.Manga
}

func NewMangaController(manga service.Manga) *MangaController {
	return &MangaController{manga}
}

func (r MangaController) GetById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		logger.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	manga, err := r.service.Get(id)

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

func (r MangaController) Create(context *gin.Context) {
	manga := models.Manga{}

	if 	err := context.BindJSON(&manga); err != nil {
		logger.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	id, err := r.service.Create(manga)
	if err != nil {
		logger.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "OK", "id": id})
}
