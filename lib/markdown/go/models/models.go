package models

// Content is the content awaited by the ngx-markdown front
type Content struct {
	Name string

	// gong:text
	Content string
}

type SvgImage struct {
	Name string // path to the image

	// gong:text gong:width 600 gong:height 400
	Content string // the svg
}

type PngImage struct {
	Name string // path to the image

	// gong:text gong:width 600 gong:height 400
	Base64Content string // the svg
}

type JpgImage struct {
	Name string // path to the image

	// gong:text gong:width 600 gong:height 400
	Base64Content string // the svg
}
