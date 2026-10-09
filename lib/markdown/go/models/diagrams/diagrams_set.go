package diagrams

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/markdown/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

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
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000003_ := (&models.AttributeShape{Name: `Base64Content`}).Stage(stage)
	__AttributeShape__00000004_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000005_ := (&models.AttributeShape{Name: `Base64Content`}).Stage(stage)
	__AttributeShape__00000006_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000007_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-07-07T03:02:10Z`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-Content`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-JpgImage`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-PngImage`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Default-SvgImage`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `Name`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.Content{}.Name
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `Content`
	__AttributeShape__00000000_.Fieldtypename = `string`

	__AttributeShape__00000001_.Name = `Content`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.Content{}.Content
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `Content`
	__AttributeShape__00000001_.Fieldtypename = `string`

	__AttributeShape__00000002_.Name = `Name`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.JpgImage{}.Name
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `JpgImage`
	__AttributeShape__00000002_.Fieldtypename = `string`

	__AttributeShape__00000003_.Name = `Base64Content`
	__AttributeShape__00000003_.IdentifierMeta = ref_models.JpgImage{}.Base64Content
	__AttributeShape__00000003_.FieldTypeAsString = ``
	__AttributeShape__00000003_.Structname = `JpgImage`
	__AttributeShape__00000003_.Fieldtypename = `string`

	__AttributeShape__00000004_.Name = `Name`
	__AttributeShape__00000004_.IdentifierMeta = ref_models.PngImage{}.Name
	__AttributeShape__00000004_.FieldTypeAsString = ``
	__AttributeShape__00000004_.Structname = `PngImage`
	__AttributeShape__00000004_.Fieldtypename = `string`

	__AttributeShape__00000005_.Name = `Base64Content`
	__AttributeShape__00000005_.IdentifierMeta = ref_models.PngImage{}.Base64Content
	__AttributeShape__00000005_.FieldTypeAsString = ``
	__AttributeShape__00000005_.Structname = `PngImage`
	__AttributeShape__00000005_.Fieldtypename = `string`

	__AttributeShape__00000006_.Name = `Content`
	__AttributeShape__00000006_.IdentifierMeta = ref_models.SvgImage{}.Content
	__AttributeShape__00000006_.FieldTypeAsString = ``
	__AttributeShape__00000006_.Structname = `SvgImage`
	__AttributeShape__00000006_.Fieldtypename = `string`

	__AttributeShape__00000007_.Name = `Name`
	__AttributeShape__00000007_.IdentifierMeta = ref_models.SvgImage{}.Name
	__AttributeShape__00000007_.FieldTypeAsString = ``
	__AttributeShape__00000007_.Structname = `SvgImage`
	__AttributeShape__00000007_.Fieldtypename = `string`

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = false
	__Classdiagram__00000000_.ShowMultiplicity = false
	__Classdiagram__00000000_.ShowLinkNames = false
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[true,true,true,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2025-07-07T03:02:10Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-Content`
	__GongStructShape__00000000_.X = 33.000000
	__GongStructShape__00000000_.Y = 88.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Content{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 103.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-JpgImage`
	__GongStructShape__00000001_.X = 424.000000
	__GongStructShape__00000001_.Y = 323.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.JpgImage{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 103.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-PngImage`
	__GongStructShape__00000002_.X = 440.000000
	__GongStructShape__00000002_.Y = 193.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.PngImage{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 103.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Default-SvgImage`
	__GongStructShape__00000003_.X = 441.000000
	__GongStructShape__00000003_.Y = 49.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.SvgImage{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 103.000000
	__GongStructShape__00000003_.IsSelected = false

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000001_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000003_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000004_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000005_)
	__GongStructShape__00000003_.AttributeShapes = append(__GongStructShape__00000003_.AttributeShapes, __AttributeShape__00000007_)
	__GongStructShape__00000003_.AttributeShapes = append(__GongStructShape__00000003_.AttributeShapes, __AttributeShape__00000006_)
}
