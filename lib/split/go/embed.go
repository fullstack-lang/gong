package split

import (
	"embed"
)

//go:embed models
var GoModelsDir embed.FS

//go:embed diagrams
var GoDiagramsDir embed.FS
