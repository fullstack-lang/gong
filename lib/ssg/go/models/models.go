package models

type Content struct {
	Name string

	//gong:text
	//gong:width 600 gong:height 300
	Text string

	// path to the content generation
	ContentPath string

	Chapters []*Chapter
}

type Chapter struct {
	Name   string
	Weigth float64

	//gong:text
	//gong:width 600 gong:height 300
	Description string
}
