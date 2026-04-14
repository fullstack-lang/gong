package diagrams

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/form/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time
var _ = slices.Index[[]int, int]

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__AttributeShape__00000000_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2026-04-13T17:37:02Z`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-Form2`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `Name`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.Form{}.Name
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `Form2`
	__AttributeShape__00000000_.Fieldtypename = `string`

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = false
	__Classdiagram__00000000_.ShowMultiplicity = false
	__Classdiagram__00000000_.ShowLinkNames = false
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2026-04-13T17:37:02Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-Form2`
	__GongStructShape__00000000_.X = 63.000000
	__GongStructShape__00000000_.Y = 82.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Form{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 83.000000
	__GongStructShape__00000000_.IsSelected = false

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000000_)
}
