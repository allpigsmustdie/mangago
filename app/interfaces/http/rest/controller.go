package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/allpigsmustdie/mangago/app/domain/service"
	"github.com/allpigsmustdie/mangago/app/interfaces/http"
	"github.com/allpigsmustdie/mangago/app/interfaces/http/rest/v1/manga"
)

func NewHandler(mangaService service.Manga) http.RESTHandler {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	{
		apiV1 := engine.Group("/v1")
		{
			mangaController := manga.NewController(mangaService)

			apiV1.Group("/manga").
				GET("/:id", mangaController.GetById).
				POST("/", mangaController.Create)
		}
	}

	return engine
}