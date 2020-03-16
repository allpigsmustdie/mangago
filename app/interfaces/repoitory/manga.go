package repoitory

import (
	"github.com/jinzhu/gorm"

	"github.com/allpigsmustdie/mangago/app/domain/models"
	"github.com/allpigsmustdie/mangago/app/domain/repository"
	"github.com/allpigsmustdie/mangago/app/interfaces/repoitory/gormodel"
)

type Manga struct {
	db *gorm.DB
}

func NewManga(db *gorm.DB) *Manga {
	db.AutoMigrate(new(gormodel.MangaRow)) 	//TODO: error handling
	return &Manga{db: db}
}

func (m Manga) Create(manga models.Manga) (id int, err error) {
	model := gormodel.NewMangaRow(manga)
	m.db.Create(model)
	return int(model.ID), m.db.Error
}

func (m Manga) Get(id int) (models.Manga, error) {
	model := gormodel.NewMangaRow(models.Manga{})
	result := m.db.First(model, id)

	if result.RecordNotFound() || int(model.ID) != id {
		return models.Manga{}, repository.ErrNotFound
	}
	return model.Manga, m.db.Error
}
