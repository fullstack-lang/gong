package models

type Content struct {
	Name     string
	Chapters []*Chapter
}

type Chapter struct {
	Name string
}
