package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/test/go/models"
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

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z := (&models.DiagramPackage{}).Stage(stage)

	__GongEnumValueShape__000000_NoName_yet := (&models.GongEnumValueShape{}).Stage(stage)

	// Setup of values

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Name = `Diagram Package created the 2025-06-04T05:37:56Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.AbsolutePathToDiagramPackage = ``

	__GongEnumValueShape__000000_NoName_yet.Name = `NoName yet`
	__GongEnumValueShape__000000_NoName_yet.IdentifierMeta = ref_models.ENUM_VAL1

	// Setup of pointers
	// setup of DiagramPackage instances pointers
	// setup of GongEnumValueShape instances pointers
}
