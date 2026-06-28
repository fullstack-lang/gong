package models

type StaticWebSite struct {
	Name string

	//gong:text
	//gong:width 900 gong:height 600
	MarkdownContent string

	Chapters []*StaticWebSiteChapter

	InputImagesDir     string
	OutputStaticWebDir string

	VersionInfo string
}

type StaticWebSiteChapter struct {
	Name string

	//gong:text
	//gong:width 900 gong:height 600
	MarkdownContent string

	Paragraphs []*StaticWebSiteParagraph
}

type StaticWebSiteParagraph struct {
	Name string

	//gong:text
	//gong:width 300 gong:height 300
	LegendMarkdownContent string

	Image *StaticWebSiteImage
}

// Iamge needs to be copied to "/static/images"
type StaticWebSiteImage struct {
	Name                string
	SourceDirectoryPath string // for instance "../../../images/logo.svg"

	Width  int
	Height int
}

// the generated image will be unstaged at the end
type StaticWebSiteGeneratedImage struct {
	Name                string
	SourceDirectoryPath string // for instance "../../../images/logo.svg"
	Width               int
	Height              int
}

// type GeneratedImageMetamodel struct {
// 	Name string

// 	ImageName   string
// 	IsMetamodel bool

// 	//gong:text
// 	//gong:width 300 gong:height 300
// 	LegendMarkdownContent string
// }
