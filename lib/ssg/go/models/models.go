package models

// the following blank import will force the vendoring of all the front end code
// this allows the front end compilation of the split front end
import (
	_ "github.com/fullstack-lang/gong/lib/ssg/go/defaults"
)

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

type Target string

// values for EnumType
const (
	FILE Target = "FILE"
	WEB  Target = "WEB"
)

type Chapter struct {
	Name string

	//gong:text
	//gong:width 600 gong:height 300
	MardownContent string

	Pages []*Page
}

type Page struct {
	Name string

	//gong:text
	//gong:width 600 gong:height 300
	MardownContent string
}
