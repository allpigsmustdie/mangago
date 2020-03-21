package manga

import (
	"log"
	"os"

	"github.com/allpigsmustdie/mangago/app/domain/service"
)

//TODO: DI logger
var logger = log.New(os.Stdout, "[manga router]", log.Lshortfile | log.Ltime)

type Controller struct {
	service service.Manga
}

func NewController(service service.Manga) *Controller {
	return &Controller{service}
}