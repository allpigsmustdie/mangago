package models

import "time"

type Manga struct {
	Name string
	Author string
	ReleaseDate time.Time
	Description string
}

func NewManga(name string, author string, releaseDate time.Time, description string) Manga {
	return Manga{Name: name, Author: author, ReleaseDate: releaseDate, Description: description}
}
