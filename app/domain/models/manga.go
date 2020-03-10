package models

import "time"

type Manga struct {
	Id int
	Name string
	Author string
	ReleaseDate time.Time
	Description string
}

func NewManga(id int, name string, author string, releaseDate time.Time, description string) *Manga {
	return &Manga{Id: id, Name: name, Author: author, ReleaseDate: releaseDate, Description: description}
}
