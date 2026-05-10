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
	StaticPath  string // path to the static directory (if empty, generation uses the default static content embedded in the binary)

	//gong:width 600 gong:height 300
	LogoSVGFile string // the content of the logo file, used generation uses the default static content embedded in the binary)

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

	// sections are added after the markdown content, they are used to generate the content of the page
	Sections []*Section

	Pages []*Page

	SubChapters []*Chapter
}

type Page struct {
	Name string

	//gong:text
	//gong:width 600 gong:height 300
	MardownContent string

	// sections are added after the markdown content, they are used to generate the content of the page
	Sections []*Section
}

type Section struct {
	Name string

	//gong:text
	//gong:width 600 gong:height 300
	MardownContent string

	IsImage  bool
	SvgImage *SvgImage
	PngImage *PngImage
	JpgImage *JpgImage

	IsDownloadableFile bool
	DownloadableFile   *DownloadableFile
}

type DownloadableFile struct {
	Name string // path to the file

	// gong:text gong:width 600 gong:height 400
	Base64Content string // the file content
}

type SvgImage struct {
	Name string // path to the image

	// gong:text gong:width 600 gong:height 400
	Content string // the svg
}

type PngImage struct {
	Name string // path to the image

	// gong:text gong:width 600 gong:height 400
	Base64Content string
}

type JpgImage struct {
	Name string // path to the image

	// gong:text gong:width 600 gong:height 400
	Base64Content string
}
