package split

import (
	"embed"

	// those are imported in order for client of split to have all
	// front code vendored

	_ "github.com/fullstack-lang/gong/lib/button/go/stack"
	_ "github.com/fullstack-lang/gong/lib/cursor/go/stack"
	_ "github.com/fullstack-lang/gong/lib/slider/go/stack"
	_ "github.com/fullstack-lang/gong/lib/svg/go/stack"
	_ "github.com/fullstack-lang/gong/lib/tone/go/stack"
	_ "github.com/fullstack-lang/gong/lib/tree/go/stack"
)

//go:embed models
var GoModelsDir embed.FS

//go:embed diagrams
var GoDiagramsDir embed.FS
