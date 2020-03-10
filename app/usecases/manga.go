package usecases

import (
	"github.com/allpigsmustdie/mangago/app/domain/models"
	"github.com/allpigsmustdie/mangago/app/domain/repository"
)

type MangaService struct {
	repo repository.Manga
}

func NewMangaService(repo repository.Manga) *MangaService {
	return &MangaService{repo: repo}
}

func (m *MangaService) Create(manga models.Manga) (id int, err error) {
	// TODO: validation
	return m.repo.Create(manga)
}

func (m *MangaService) Get(id int) (*models.Manga, error) {
	return m.repo.Get(id)
}

