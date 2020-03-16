package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/allpigsmustdie/mangago/app/domain/service"
	v1 "github.com/allpigsmustdie/mangago/app/interfaces/rest/v1"
)

func NewHandler(mangaService service.Manga) http.Handler {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	{
		apiV1 := engine.Group("/v1")
		{
			mangaController := v1.NewMangaController(mangaService)

			apiV1.Group("/manga").
				GET(":id", mangaController.GetById).
				POST("/", mangaController.Create)
		}
	}

	return engine
}