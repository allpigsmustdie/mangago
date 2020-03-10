package service

type Manga interface {
	Create(manga Manga) (id int, err error)
	Get(id int) (*Manga, error)
}