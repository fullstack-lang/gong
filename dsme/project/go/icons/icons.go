package icons

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	_ "embed"
)

//go:embed "percent_000.svg"
var percent_000 string
var Percent_000_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_000",
	SVG:  percent_000,
})

//go:embed "percent_025.svg"
var percent_025 string
var Percent_025_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_025",
	SVG:  percent_025,
})

//go:embed "percent_050.svg"
var percent_050 string
var Percent_050_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_050",
	SVG:  percent_050,
})

//go:embed "percent_075.svg"
var percent_075 string
var Percent_075_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_075",
	SVG:  percent_075,
})

//go:embed "percent_100.svg"
var percent_100 string
var Percent_100_Icon *tree.SVGIcon = (&tree.SVGIcon{
	Name: "percent_100",
	SVG:  percent_100,
})
