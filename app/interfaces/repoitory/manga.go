package repoitory

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/allpigsmustdie/mangago/app/domain/models"
)

type mangaRow struct {
	gorm.Model
	models.Manga
}

func newMangaRow(manga models.Manga) *mangaRow {
	return &mangaRow{Manga: manga}
}

type Manga struct {
	db *gorm.DB
}

func NewManga(db *gorm.DB) *Manga {
	db.AutoMigrate(&mangaRow{})
	return &Manga{db: db}
}

func (m Manga) Create(manga models.Manga) (id int, err error) {
	model := newMangaRow(manga)
	m.db.Create(model)
	return int(model.ID), m.db.Error
}

func (m Manga) Get(id int) (models.Manga, error) {
	model := newMangaRow(models.Manga{})
	m.db.First(model, id)
	err := m.db.Error
	if  err != nil {
		return models.Manga{}, err
	}
	return model.Manga, err
}
