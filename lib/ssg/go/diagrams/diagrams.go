package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/ssg/go/models"
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

	const __write__local_time = "2025-12-03 00:12:14.658179 CET"
	const __write__utc_time__ = "2025-12-02 23:12:14.658179 UTC"

	const __commitId__ = "0000000046"

	// Declaration of instances to stage

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_MardownContent := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_MardownContent := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000004_ContentPath := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000005_OutputPath := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000006_LayoutPath := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000007_StaticPath := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000008_Target := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000009_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000010_MardownContent := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000011_LogoFileName := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_09T00_55_33Z := (&models.DiagramPackage{}).Stage(stage)

	__GongEnumShape__000000_Default_Target := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueShape__000000_FILE := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000001_WEB := (&models.GongEnumValueShape{}).Stage(stage)

	__GongStructShape__000000_Default_Chapter := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Content := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Page := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Pages := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Chapters := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.Chapter{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Chapter`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_MardownContent.Name = `MardownContent`
	__AttributeShape__000001_MardownContent.IdentifierMeta = ref_models.Chapter{}.MardownContent
	__AttributeShape__000001_MardownContent.FieldTypeAsString = ``
	__AttributeShape__000001_MardownContent.Structname = `Chapter`
	__AttributeShape__000001_MardownContent.Fieldtypename = `string`

	__AttributeShape__000002_Name.Name = `Name`
	__AttributeShape__000002_Name.IdentifierMeta = ref_models.Content{}.Name
	__AttributeShape__000002_Name.FieldTypeAsString = ``
	__AttributeShape__000002_Name.Structname = `Content`
	__AttributeShape__000002_Name.Fieldtypename = `string`

	__AttributeShape__000003_MardownContent.Name = `MardownContent`
	__AttributeShape__000003_MardownContent.IdentifierMeta = ref_models.Content{}.MardownContent
	__AttributeShape__000003_MardownContent.FieldTypeAsString = ``
	__AttributeShape__000003_MardownContent.Structname = `Content`
	__AttributeShape__000003_MardownContent.Fieldtypename = `string`

	__AttributeShape__000004_ContentPath.Name = `ContentPath`
	__AttributeShape__000004_ContentPath.IdentifierMeta = ref_models.Content{}.ContentPath
	__AttributeShape__000004_ContentPath.FieldTypeAsString = ``
	__AttributeShape__000004_ContentPath.Structname = `Content`
	__AttributeShape__000004_ContentPath.Fieldtypename = `string`

	__AttributeShape__000005_OutputPath.Name = `OutputPath`
	__AttributeShape__000005_OutputPath.IdentifierMeta = ref_models.Content{}.OutputPath
	__AttributeShape__000005_OutputPath.FieldTypeAsString = ``
	__AttributeShape__000005_OutputPath.Structname = `Content`
	__AttributeShape__000005_OutputPath.Fieldtypename = `string`

	__AttributeShape__000006_LayoutPath.Name = `LayoutPath`
	__AttributeShape__000006_LayoutPath.IdentifierMeta = ref_models.Content{}.LayoutPath
	__AttributeShape__000006_LayoutPath.FieldTypeAsString = ``
	__AttributeShape__000006_LayoutPath.Structname = `Content`
	__AttributeShape__000006_LayoutPath.Fieldtypename = `string`

	__AttributeShape__000007_StaticPath.Name = `StaticPath`
	__AttributeShape__000007_StaticPath.IdentifierMeta = ref_models.Content{}.StaticPath
	__AttributeShape__000007_StaticPath.FieldTypeAsString = ``
	__AttributeShape__000007_StaticPath.Structname = `Content`
	__AttributeShape__000007_StaticPath.Fieldtypename = `string`

	__AttributeShape__000008_Target.Name = `Target`
	__AttributeShape__000008_Target.IdentifierMeta = ref_models.Content{}.Target
	__AttributeShape__000008_Target.FieldTypeAsString = ``
	__AttributeShape__000008_Target.Structname = `Content`
	__AttributeShape__000008_Target.Fieldtypename = `Target`

	__AttributeShape__000009_Name.Name = `Name`
	__AttributeShape__000009_Name.IdentifierMeta = ref_models.Page{}.Name
	__AttributeShape__000009_Name.FieldTypeAsString = ``
	__AttributeShape__000009_Name.Structname = `Page`
	__AttributeShape__000009_Name.Fieldtypename = `string`

	__AttributeShape__000010_MardownContent.Name = `MardownContent`
	__AttributeShape__000010_MardownContent.IdentifierMeta = ref_models.Page{}.MardownContent
	__AttributeShape__000010_MardownContent.FieldTypeAsString = ``
	__AttributeShape__000010_MardownContent.Structname = `Page`
	__AttributeShape__000010_MardownContent.Fieldtypename = `string`

	__AttributeShape__000011_LogoFileName.Name = `LogoFileName`
	__AttributeShape__000011_LogoFileName.IdentifierMeta = ref_models.Content{}.LogoFileName
	__AttributeShape__000011_LogoFileName.FieldTypeAsString = ``
	__AttributeShape__000011_LogoFileName.Structname = `Content`
	__AttributeShape__000011_LogoFileName.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.ShowNbInstances = false
	__Classdiagram__000000_Default.ShowMultiplicity = false
	__Classdiagram__000000_Default.ShowLinkNames = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,false,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = `[true]`
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_09T00_55_33Z.Name = `Diagram Package created the 2025-05-09T00:55:33Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_09T00_55_33Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_09T00_55_33Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_09T00_55_33Z.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__000000_Default_Target.Name = `Default-Target`
	__GongEnumShape__000000_Default_Target.X = 47.000000
	__GongEnumShape__000000_Default_Target.Y = 272.000000
	__GongEnumShape__000000_Default_Target.IdentifierMeta = new(ref_models.Target)
	__GongEnumShape__000000_Default_Target.Width = 240.000000
	__GongEnumShape__000000_Default_Target.Height = 103.000000
	__GongEnumShape__000000_Default_Target.IsExpanded = false

	__GongEnumValueShape__000000_FILE.Name = `FILE`
	__GongEnumValueShape__000000_FILE.IdentifierMeta = ref_models.FILE

	__GongEnumValueShape__000001_WEB.Name = `WEB`
	__GongEnumValueShape__000001_WEB.IdentifierMeta = ref_models.WEB

	__GongStructShape__000000_Default_Chapter.Name = `Default-Chapter`
	__GongStructShape__000000_Default_Chapter.X = 497.000000
	__GongStructShape__000000_Default_Chapter.Y = 98.000000
	__GongStructShape__000000_Default_Chapter.IdentifierMeta = ref_models.Chapter{}
	__GongStructShape__000000_Default_Chapter.Width = 240.000000
	__GongStructShape__000000_Default_Chapter.Height = 103.000000
	__GongStructShape__000000_Default_Chapter.IsSelected = false

	__GongStructShape__000001_Default_Content.Name = `Default-Content`
	__GongStructShape__000001_Default_Content.X = 46.000000
	__GongStructShape__000001_Default_Content.Y = 44.000000
	__GongStructShape__000001_Default_Content.IdentifierMeta = ref_models.Content{}
	__GongStructShape__000001_Default_Content.Width = 240.000000
	__GongStructShape__000001_Default_Content.Height = 223.000000
	__GongStructShape__000001_Default_Content.IsSelected = false

	__GongStructShape__000002_Default_Page.Name = `Default-Page`
	__GongStructShape__000002_Default_Page.X = 941.000000
	__GongStructShape__000002_Default_Page.Y = 107.000000
	__GongStructShape__000002_Default_Page.IdentifierMeta = ref_models.Page{}
	__GongStructShape__000002_Default_Page.Width = 240.000000
	__GongStructShape__000002_Default_Page.Height = 103.000000
	__GongStructShape__000002_Default_Page.IsSelected = false

	__LinkShape__000000_Pages.Name = `Pages`
	__LinkShape__000000_Pages.IdentifierMeta = ref_models.Chapter{}.Pages
	__LinkShape__000000_Pages.FieldTypeIdentifierMeta = ref_models.Page{}
	__LinkShape__000000_Pages.FieldOffsetX = 0.000000
	__LinkShape__000000_Pages.FieldOffsetY = 0.000000
	__LinkShape__000000_Pages.TargetMultiplicity = models.MANY
	__LinkShape__000000_Pages.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Pages.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Pages.SourceMultiplicity = models.MANY
	__LinkShape__000000_Pages.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Pages.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Pages.X = 798.500000
	__LinkShape__000000_Pages.Y = 107.000000
	__LinkShape__000000_Pages.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Pages.StartRatio = 0.500000
	__LinkShape__000000_Pages.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Pages.EndRatio = 0.500000
	__LinkShape__000000_Pages.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Chapters.Name = `Chapters`
	__LinkShape__000001_Chapters.IdentifierMeta = ref_models.Content{}.Chapters
	__LinkShape__000001_Chapters.FieldTypeIdentifierMeta = ref_models.Chapter{}
	__LinkShape__000001_Chapters.FieldOffsetX = 0.000000
	__LinkShape__000001_Chapters.FieldOffsetY = 0.000000
	__LinkShape__000001_Chapters.TargetMultiplicity = models.MANY
	__LinkShape__000001_Chapters.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Chapters.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Chapters.SourceMultiplicity = models.MANY
	__LinkShape__000001_Chapters.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Chapters.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Chapters.X = 574.000000
	__LinkShape__000001_Chapters.Y = 147.500000
	__LinkShape__000001_Chapters.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Chapters.StartRatio = 0.500000
	__LinkShape__000001_Chapters.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Chapters.EndRatio = 0.500000
	__LinkShape__000001_Chapters.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Chapter)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Content)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Page)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_Target)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_09T00_55_33Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_09T00_55_33Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_09T00_55_33Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongEnumShape instances pointers
	__GongEnumShape__000000_Default_Target.GongEnumValueShapes = append(__GongEnumShape__000000_Default_Target.GongEnumValueShapes, __GongEnumValueShape__000000_FILE)
	__GongEnumShape__000000_Default_Target.GongEnumValueShapes = append(__GongEnumShape__000000_Default_Target.GongEnumValueShapes, __GongEnumValueShape__000001_WEB)
	// setup of GongEnumValueShape instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Chapter.AttributeShapes = append(__GongStructShape__000000_Default_Chapter.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_Chapter.AttributeShapes = append(__GongStructShape__000000_Default_Chapter.AttributeShapes, __AttributeShape__000001_MardownContent)
	__GongStructShape__000000_Default_Chapter.LinkShapes = append(__GongStructShape__000000_Default_Chapter.LinkShapes, __LinkShape__000000_Pages)
	__GongStructShape__000001_Default_Content.AttributeShapes = append(__GongStructShape__000001_Default_Content.AttributeShapes, __AttributeShape__000002_Name)
	__GongStructShape__000001_Default_Content.AttributeShapes = append(__GongStructShape__000001_Default_Content.AttributeShapes, __AttributeShape__000003_MardownContent)
	__GongStructShape__000001_Default_Content.AttributeShapes = append(__GongStructShape__000001_Default_Content.AttributeShapes, __AttributeShape__000004_ContentPath)
	__GongStructShape__000001_Default_Content.AttributeShapes = append(__GongStructShape__000001_Default_Content.AttributeShapes, __AttributeShape__000005_OutputPath)
	__GongStructShape__000001_Default_Content.AttributeShapes = append(__GongStructShape__000001_Default_Content.AttributeShapes, __AttributeShape__000006_LayoutPath)
	__GongStructShape__000001_Default_Content.AttributeShapes = append(__GongStructShape__000001_Default_Content.AttributeShapes, __AttributeShape__000007_StaticPath)
	__GongStructShape__000001_Default_Content.AttributeShapes = append(__GongStructShape__000001_Default_Content.AttributeShapes, __AttributeShape__000011_LogoFileName)
	__GongStructShape__000001_Default_Content.AttributeShapes = append(__GongStructShape__000001_Default_Content.AttributeShapes, __AttributeShape__000008_Target)
	__GongStructShape__000001_Default_Content.LinkShapes = append(__GongStructShape__000001_Default_Content.LinkShapes, __LinkShape__000001_Chapters)
	__GongStructShape__000002_Default_Page.AttributeShapes = append(__GongStructShape__000002_Default_Page.AttributeShapes, __AttributeShape__000009_Name)
	__GongStructShape__000002_Default_Page.AttributeShapes = append(__GongStructShape__000002_Default_Page.AttributeShapes, __AttributeShape__000010_MardownContent)
	// setup of LinkShape instances pointers
}

