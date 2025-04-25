package models

type Content struct {
	Name string

	//gong:text
	//gong:width 600 gong:height 300
	MardownContent string

	// path to the content generation
	ContentPath string
	OutputPath  string
	LayoutPath  string
	StaticPath  string

	Target Target

	Chapters []*Chapter
}

type Chapter struct {
	Name   string
	Weigth float64

	//gong:text
	//gong:width 600 gong:height 300
	MardownContent string
}

type Target string

// values for EnumType
const (
	FILE Target = "FILE"
	WEB  Target = "WEB"
)
