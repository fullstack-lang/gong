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
	ContentPath string // path to the markdown content
	OutputPath  string // path to the output directory
	LayoutPath  string // path to the layout directory
	StaticPath  string // path to the static directory

	// BespokeLogoFileName is the filename is the file located in the StaticPath/images
	// default "logo.svg"
	IsBespokeLogoFileName bool
	BespokeLogoFileName   string

	IsBespokePageTileLogoFileName bool
	BespokePageTileLogoFileName   string

	Target Target

	Chapters []*Chapter

	VersionInfo string // will appear on the top right
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
