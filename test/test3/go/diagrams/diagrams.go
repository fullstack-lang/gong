package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/test3/go/models"
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

	// insertion point for declaration of instances to stage

	__AttributeShape__00000000_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-06-13T23:34:45Z`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-A`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-B`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `B`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `Bs`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `Name`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.A{}.Name
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `A`
	__AttributeShape__00000000_.Fieldtypename = `string`

	__AttributeShape__00000001_.Name = `Name`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.B{}.Name
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `B`
	__AttributeShape__00000001_.Fieldtypename = `string`

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[true,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2025-06-13T23:34:45Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-A`
	__GongStructShape__00000000_.X = 77.000000
	__GongStructShape__00000000_.Y = 64.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.A{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 83.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-B`
	__GongStructShape__00000001_.X = 381.000000
	__GongStructShape__00000001_.Y = 245.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.B{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 83.000000
	__GongStructShape__00000001_.IsSelected = false

	__LinkShape__00000000_.Name = `B`
	__LinkShape__00000000_.IdentifierMeta = ref_models.A{}.B
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.B{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 617.500000
	__LinkShape__00000000_.Y = 128.000000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000000_.StartRatio = 0.219824
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.500000
	__LinkShape__00000000_.CornerOffsetRatio = 1.060288

	__LinkShape__00000001_.Name = `Bs`
	__LinkShape__00000001_.IdentifierMeta = ref_models.A{}.Bs
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.B{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 617.500000
	__LinkShape__00000001_.Y = 128.000000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.500000
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.500000
	__LinkShape__00000001_.CornerOffsetRatio = 2.740658

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000001_)
}
