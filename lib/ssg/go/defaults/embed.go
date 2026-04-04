package defaults

import "embed"

//go:embed all:layouts
var LayoutsFS embed.FS

//go:embed all:static
var StaticFS embed.FS
