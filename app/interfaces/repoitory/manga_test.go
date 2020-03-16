package repoitory

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/allpigsmustdie/mangago/app/domain/models"
	"github.com/allpigsmustdie/mangago/app/domain/repository"
)

var _ repository.Manga = new(Manga)

func newDb() (db *gorm.DB, err error) {
	return gorm.Open("sqlite3", ":memory:")
}

func TestMangaRepo(t *testing.T) {
	db, err := newDb()
	if err != nil {
		t.Fatalf("failed to connect database: %v\n", err)
	}
	defer db.Close()
	repo := NewManga(db)

	model := models.NewManga(
		"Neon Genesis Evangelion",
		"Yoshiyuki Sadamoto",
		time.Date(1994, time.December, 26, 0, 0, 0, 0, time.UTC),
		`Neon Genesis Evangelion is a manga series by Evangelion character designer Yoshiyuki Sadamoto. It is an alternate retelling of the story of the Neon Genesis Evangelion TV series, constituting a separate continuity, with many subtle and obvious changes from the series on which it is based. The manga ran for almost two decades, and was serialized from December 26, 1994 to June 4, 2013 in Shōnen Ace (1994-2009) and Young Ace (2009-2013), and consisted of a total of 96 chapters and a bonus chapter spread across 14 Tankōbon volumes. The manga also ran before the series premiered and well after the series had ended, with its debut some 10 months ahead of the show's release in October 1995 thanks to delays in its production. However, after the early chapters Sadamoto, who prioritized other projects, maintained a very inconsistent release schedule, which caused multiple years-long hiatuses and significant portions of the story to be cut. As a result of this, the manga was only concluded some 17 years after the series' end in March 1996. Links to the summaries of each Tankōbon volume of the manga are listed below.

Unlike it is sometimes assumed, the manga is in fact not the original version of the story, but instead it was created as a supplemental designed to promote the TV series and was continued thanks to the series' popularity. Sadamoto has also repeatedly stated that the manga is his own individual work and should never be used as a reference for anything in the anime. He has also denied links to the Rebuild movies.`,
	)

	id, err := repo.Create(model)

	if err != nil {
		t.Fatalf("failed to create: %v\n", err)
	}

	result, err := repo.Get(id)

	if err != nil {
		t.Fatalf("failed to get by id: %v\n", err)
	}

	if model != result {
		t.Fatalf("model != result\n%#v\n%#v\n", model, result)
	}

}

func TestManga_Get(t *testing.T) {
	db, err := newDb()
	if err != nil {
		t.Fatalf("failed to connect database: %v\n", err)
	}
	defer db.Close()
	repo := NewManga(db)

	_, err = repo.Get(100500)

	if err != repository.ErrNotFound {
		t.Fatalf("err != repository.ErrNotFound, got %v", err)
	}

}
