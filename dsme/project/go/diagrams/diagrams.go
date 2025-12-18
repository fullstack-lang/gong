package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/dsme/project/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z := (&models.DiagramPackage{}).Stage(stage)

	// Setup of values

	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Name = `Diagram Package created the 2025-12-18T17:41:37Z`
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Path = ``
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.GongModelPath = ``
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.AbsolutePathToDiagramPackage = ``

	// Setup of pointers
	// setup of DiagramPackage instances pointers
}

