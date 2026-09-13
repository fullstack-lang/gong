package diagrams

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/ssg/go/models"
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
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `MardownContent`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000003_ := (&models.AttributeShape{Name: `MardownContent`}).Stage(stage)
	__AttributeShape__00000004_ := (&models.AttributeShape{Name: `ContentPath`}).Stage(stage)
	__AttributeShape__00000005_ := (&models.AttributeShape{Name: `OutputPath`}).Stage(stage)
	__AttributeShape__00000007_ := (&models.AttributeShape{Name: `StaticPath`}).Stage(stage)
	__AttributeShape__00000008_ := (&models.AttributeShape{Name: `Target`}).Stage(stage)
	__AttributeShape__00000009_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000010_ := (&models.AttributeShape{Name: `MardownContent`}).Stage(stage)
	__AttributeShape__00000011_ := (&models.AttributeShape{Name: `LogoFileName`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-05-09T00:55:33Z`}).Stage(stage)

	__GongEnumShape__00000000_ := (&models.GongEnumShape{Name: `Default-Target`}).Stage(stage)

	__GongEnumValueShape__00000000_ := (&models.GongEnumValueShape{Name: `FILE`}).Stage(stage)
	__GongEnumValueShape__00000001_ := (&models.GongEnumValueShape{Name: `WEB`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-Chapter`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Content`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-Page`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `Pages`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `Chapters`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `Name`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.Chapter{}.Name
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `Chapter`
	__AttributeShape__00000000_.Fieldtypename = `string`

	__AttributeShape__00000001_.Name = `MardownContent`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.Chapter{}.MardownContent
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `Chapter`
	__AttributeShape__00000001_.Fieldtypename = `string`

	__AttributeShape__00000002_.Name = `Name`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.Content{}.Name
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `Content`
	__AttributeShape__00000002_.Fieldtypename = `string`

	__AttributeShape__00000003_.Name = `MardownContent`
	__AttributeShape__00000003_.IdentifierMeta = ref_models.Content{}.MardownContent
	__AttributeShape__00000003_.FieldTypeAsString = ``
	__AttributeShape__00000003_.Structname = `Content`
	__AttributeShape__00000003_.Fieldtypename = `string`

	__AttributeShape__00000004_.Name = `ContentPath`
	__AttributeShape__00000004_.IdentifierMeta = ref_models.Content{}.ContentPath
	__AttributeShape__00000004_.FieldTypeAsString = ``
	__AttributeShape__00000004_.Structname = `Content`
	__AttributeShape__00000004_.Fieldtypename = `string`

	__AttributeShape__00000005_.Name = `OutputPath`
	__AttributeShape__00000005_.IdentifierMeta = ref_models.Content{}.OutputPath
	__AttributeShape__00000005_.FieldTypeAsString = ``
	__AttributeShape__00000005_.Structname = `Content`
	__AttributeShape__00000005_.Fieldtypename = `string`

	__AttributeShape__00000007_.Name = `StaticPath`
	__AttributeShape__00000007_.IdentifierMeta = ref_models.Content{}.StaticPath
	__AttributeShape__00000007_.FieldTypeAsString = ``
	__AttributeShape__00000007_.Structname = `Content`
	__AttributeShape__00000007_.Fieldtypename = `string`

	__AttributeShape__00000008_.Name = `Target`
	__AttributeShape__00000008_.IdentifierMeta = ref_models.Content{}.Target
	__AttributeShape__00000008_.FieldTypeAsString = ``
	__AttributeShape__00000008_.Structname = `Content`
	__AttributeShape__00000008_.Fieldtypename = `Target`

	__AttributeShape__00000009_.Name = `Name`
	__AttributeShape__00000009_.IdentifierMeta = ref_models.Page{}.Name
	__AttributeShape__00000009_.FieldTypeAsString = ``
	__AttributeShape__00000009_.Structname = `Page`
	__AttributeShape__00000009_.Fieldtypename = `string`

	__AttributeShape__00000010_.Name = `MardownContent`
	__AttributeShape__00000010_.IdentifierMeta = ref_models.Page{}.MardownContent
	__AttributeShape__00000010_.FieldTypeAsString = ``
	__AttributeShape__00000010_.Structname = `Page`
	__AttributeShape__00000010_.Fieldtypename = `string`

	__AttributeShape__00000011_.Name = `LogoFileName`
	__AttributeShape__00000011_.IdentifierMeta = ref_models.Content{}.BespokeLogoFileName
	__AttributeShape__00000011_.FieldTypeAsString = ``
	__AttributeShape__00000011_.Structname = `Content`
	__AttributeShape__00000011_.Fieldtypename = `string`

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = false
	__Classdiagram__00000000_.ShowMultiplicity = false
	__Classdiagram__00000000_.ShowLinkNames = false
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,false,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = true
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = `[true]`
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2025-05-09T00:55:33Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__00000000_.Name = `Default-Target`
	__GongEnumShape__00000000_.X = 47.000000
	__GongEnumShape__00000000_.Y = 272.000000
	__GongEnumShape__00000000_.IdentifierMeta = new(ref_models.Target)
	__GongEnumShape__00000000_.Width = 240.000000
	__GongEnumShape__00000000_.Height = 103.000000
	__GongEnumShape__00000000_.IsExpanded = false

	__GongEnumValueShape__00000000_.Name = `FILE`
	__GongEnumValueShape__00000000_.IdentifierMeta = ref_models.FILE

	__GongEnumValueShape__00000001_.Name = `WEB`
	__GongEnumValueShape__00000001_.IdentifierMeta = ref_models.WEB

	__GongStructShape__00000000_.Name = `Default-Chapter`
	__GongStructShape__00000000_.X = 497.000000
	__GongStructShape__00000000_.Y = 98.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Chapter{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 103.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Content`
	__GongStructShape__00000001_.X = 46.000000
	__GongStructShape__00000001_.Y = 44.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Content{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 223.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-Page`
	__GongStructShape__00000002_.X = 941.000000
	__GongStructShape__00000002_.Y = 107.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.Page{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 103.000000
	__GongStructShape__00000002_.IsSelected = false

	__LinkShape__00000000_.Name = `Pages`
	__LinkShape__00000000_.IdentifierMeta = ref_models.Chapter{}.Pages
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.Page{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.MANY
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 798.500000
	__LinkShape__00000000_.Y = 107.000000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.500000
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.500000
	__LinkShape__00000000_.CornerOffsetRatio = 1.380000

	__LinkShape__00000001_.Name = `Chapters`
	__LinkShape__00000001_.IdentifierMeta = ref_models.Content{}.Chapters
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.Chapter{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 574.000000
	__LinkShape__00000001_.Y = 147.500000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.500000
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.500000
	__LinkShape__00000001_.CornerOffsetRatio = 1.380000

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongEnumShapes = append(__Classdiagram__00000000_.GongEnumShapes, __GongEnumShape__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000000_)
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000001_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000001_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000003_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000004_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000005_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000007_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000011_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000008_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000009_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000010_)
}
