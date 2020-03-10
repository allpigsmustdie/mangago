package gormodel

import (
	"github.com/jinzhu/gorm"

	"github.com/allpigsmustdie/mangago/app/domain/models"
)

type MangaRow struct {
	gorm.Model
	models.Manga
}

func NewMangaRow(manga models.Manga) *MangaRow {
	return &MangaRow{Manga: manga}
}