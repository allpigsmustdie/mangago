package repository

import "github.com/allpigsmustdie/mangago/app/domain/models"

type Manga interface {
	Create(manga models.Manga) (id int, err error)
	Get(id int) (*models.Manga, error)
}