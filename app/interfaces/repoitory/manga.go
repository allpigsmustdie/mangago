package repoitory

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/allpigsmustdie/mangago/app/domain/models"
	"github.com/allpigsmustdie/mangago/app/interfaces/repoitory/gormodel"
)

type Manga struct {
	db *gorm.DB
}

func NewManga(db *gorm.DB) *Manga {
	db.AutoMigrate(new(gormodel.MangaRow))
	return &Manga{db: db}
}

func (m Manga) Create(manga models.Manga) (id int, err error) {
	model := gormodel.NewMangaRow(manga)
	m.db.Create(model)
	return int(model.ID), m.db.Error
}

func (m Manga) Get(id int) (models.Manga, error) {
	model := gormodel.NewMangaRow(models.Manga{})
	m.db.First(model, id)
	err := m.db.Error
	if  err != nil {
		return models.Manga{}, err
	}
	return model.Manga, err
}
