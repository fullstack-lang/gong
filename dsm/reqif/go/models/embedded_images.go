package models

type EmbeddedSvgImage struct {
	Name string // path to the image

	// gong:text gong:width 600 gong:height 400
	Content string // the svg
}

type EmbeddedJpgImage struct {
	Name string // path to the image

	// gong:text gong:width 600 gong:height 400
	Base64Content string // the svg
}

type EmbeddedPngImage struct {
	Name string // path to the image

	// gong:text gong:width 600 gong:height 400
	Base64Content string // the svg
}
