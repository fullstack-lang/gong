package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/load/go/models"
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

	const __write__local_time = "2025-07-05 18:10:52.292469 CEST"
	const __write__utc_time__ = "2025-07-05 16:10:52.292469 UTC"

	const __commitId__ = "0000000028"

	// Declaration of instances to stage

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_Content := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_Base64EncodedContent := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000004_Name := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_05T16_09_49Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_FileToDownload := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_FileToUpload := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Message := (&models.GongStructShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.FileToDownload{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `FileToDownload`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_Content.Name = `Content`
	__AttributeShape__000001_Content.IdentifierMeta = ref_models.FileToDownload{}.Base64EncodedContent
	__AttributeShape__000001_Content.FieldTypeAsString = ``
	__AttributeShape__000001_Content.Structname = `FileToDownload`
	__AttributeShape__000001_Content.Fieldtypename = `string`

	__AttributeShape__000002_Name.Name = `Name`
	__AttributeShape__000002_Name.IdentifierMeta = ref_models.FileToUpload{}.Name
	__AttributeShape__000002_Name.FieldTypeAsString = ``
	__AttributeShape__000002_Name.Structname = `FileToUpload`
	__AttributeShape__000002_Name.Fieldtypename = `string`

	__AttributeShape__000003_Base64EncodedContent.Name = `Base64EncodedContent`
	__AttributeShape__000003_Base64EncodedContent.IdentifierMeta = ref_models.FileToUpload{}.Base64EncodedContent
	__AttributeShape__000003_Base64EncodedContent.FieldTypeAsString = ``
	__AttributeShape__000003_Base64EncodedContent.Structname = `FileToUpload`
	__AttributeShape__000003_Base64EncodedContent.Fieldtypename = `string`

	__AttributeShape__000004_Name.Name = `Name`
	__AttributeShape__000004_Name.IdentifierMeta = ref_models.Message{}.Name
	__AttributeShape__000004_Name.FieldTypeAsString = ``
	__AttributeShape__000004_Name.Structname = `Message`
	__AttributeShape__000004_Name.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[true,true,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_05T16_09_49Z.Name = `Diagram Package created the 2025-07-05T16:09:49Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_05T16_09_49Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_05T16_09_49Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_05T16_09_49Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_FileToDownload.Name = `Default-FileToDownload`
	__GongStructShape__000000_Default_FileToDownload.X = 27.000000
	__GongStructShape__000000_Default_FileToDownload.Y = 36.000000
	__GongStructShape__000000_Default_FileToDownload.IdentifierMeta = ref_models.FileToDownload{}
	__GongStructShape__000000_Default_FileToDownload.Width = 291.000000
	__GongStructShape__000000_Default_FileToDownload.Height = 103.000000
	__GongStructShape__000000_Default_FileToDownload.IsSelected = false

	__GongStructShape__000001_Default_FileToUpload.Name = `Default-FileToUpload`
	__GongStructShape__000001_Default_FileToUpload.X = 20.000000
	__GongStructShape__000001_Default_FileToUpload.Y = 159.000000
	__GongStructShape__000001_Default_FileToUpload.IdentifierMeta = ref_models.FileToUpload{}
	__GongStructShape__000001_Default_FileToUpload.Width = 304.000000
	__GongStructShape__000001_Default_FileToUpload.Height = 103.000000
	__GongStructShape__000001_Default_FileToUpload.IsSelected = false

	__GongStructShape__000002_Default_Message.Name = `Default-Message`
	__GongStructShape__000002_Default_Message.X = 484.000000
	__GongStructShape__000002_Default_Message.Y = 24.000000
	__GongStructShape__000002_Default_Message.IdentifierMeta = ref_models.Message{}
	__GongStructShape__000002_Default_Message.Width = 240.000000
	__GongStructShape__000002_Default_Message.Height = 83.000000
	__GongStructShape__000002_Default_Message.IsSelected = false

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_FileToDownload)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_FileToUpload)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Message)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_05T16_09_49Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_07_05T16_09_49Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_07_05T16_09_49Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_FileToDownload.AttributeShapes = append(__GongStructShape__000000_Default_FileToDownload.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_FileToDownload.AttributeShapes = append(__GongStructShape__000000_Default_FileToDownload.AttributeShapes, __AttributeShape__000001_Content)
	__GongStructShape__000001_Default_FileToUpload.AttributeShapes = append(__GongStructShape__000001_Default_FileToUpload.AttributeShapes, __AttributeShape__000002_Name)
	__GongStructShape__000001_Default_FileToUpload.AttributeShapes = append(__GongStructShape__000001_Default_FileToUpload.AttributeShapes, __AttributeShape__000003_Base64EncodedContent)
	__GongStructShape__000002_Default_Message.AttributeShapes = append(__GongStructShape__000002_Default_Message.AttributeShapes, __AttributeShape__000004_Name)
}
