package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_Content := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_Base64EncodedContent := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000004_Name := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_21T06_22_15Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Message := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_FileToUpload := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_FileToDownload := (&models.GongStructShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.FileToDownload.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.FileToDownload.Name`
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `FileToDownload`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_Content.Name = `Content`

	//gong:ident [ref_models.FileToDownload.Content] comment added to overcome the problem with the comment map association
	__AttributeShape__000001_Content.Identifier = `ref_models.FileToDownload.Content`
	__AttributeShape__000001_Content.FieldTypeAsString = ``
	__AttributeShape__000001_Content.Structname = `FileToDownload`
	__AttributeShape__000001_Content.Fieldtypename = `string`

	__AttributeShape__000002_Name.Name = `Name`

	//gong:ident [ref_models.FileToUpload.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000002_Name.Identifier = `ref_models.FileToUpload.Name`
	__AttributeShape__000002_Name.FieldTypeAsString = ``
	__AttributeShape__000002_Name.Structname = `FileToUpload`
	__AttributeShape__000002_Name.Fieldtypename = `string`

	__AttributeShape__000003_Base64EncodedContent.Name = `Base64EncodedContent`

	//gong:ident [ref_models.FileToUpload.Base64EncodedContent] comment added to overcome the problem with the comment map association
	__AttributeShape__000003_Base64EncodedContent.Identifier = `ref_models.FileToUpload.Base64EncodedContent`
	__AttributeShape__000003_Base64EncodedContent.FieldTypeAsString = ``
	__AttributeShape__000003_Base64EncodedContent.Structname = `FileToUpload`
	__AttributeShape__000003_Base64EncodedContent.Fieldtypename = `string`

	__AttributeShape__000004_Name.Name = `Name`

	//gong:ident [ref_models.Message.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000004_Name.Identifier = `ref_models.Message.Name`
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

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_21T06_22_15Z.Name = `Diagram Package created the 2025-05-21T06:22:15Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_21T06_22_15Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_21T06_22_15Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_21T06_22_15Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Message.Name = `Default-Message`
	__GongStructShape__000000_Default_Message.X = 82.000000
	__GongStructShape__000000_Default_Message.Y = 59.000000

	//gong:ident [ref_models.Message] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Message.Identifier = `ref_models.Message`
	__GongStructShape__000000_Default_Message.IdentifierMeta = ref_models.Message{}
	__GongStructShape__000000_Default_Message.ShowNbInstances = false
	__GongStructShape__000000_Default_Message.NbInstances = 0
	__GongStructShape__000000_Default_Message.Width = 240.000000
	__GongStructShape__000000_Default_Message.Height = 83.000000
	__GongStructShape__000000_Default_Message.IsSelected = false

	__GongStructShape__000001_Default_FileToUpload.Name = `Default-FileToUpload`
	__GongStructShape__000001_Default_FileToUpload.X = 865.000000
	__GongStructShape__000001_Default_FileToUpload.Y = 55.000000

	//gong:ident [ref_models.FileToUpload] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_FileToUpload.Identifier = `ref_models.FileToUpload`
	__GongStructShape__000001_Default_FileToUpload.IdentifierMeta = ref_models.FileToUpload{}
	__GongStructShape__000001_Default_FileToUpload.ShowNbInstances = false
	__GongStructShape__000001_Default_FileToUpload.NbInstances = 0
	__GongStructShape__000001_Default_FileToUpload.Width = 301.000000
	__GongStructShape__000001_Default_FileToUpload.Height = 103.000000
	__GongStructShape__000001_Default_FileToUpload.IsSelected = false

	__GongStructShape__000002_Default_FileToDownload.Name = `Default-FileToDownload`
	__GongStructShape__000002_Default_FileToDownload.X = 463.000000
	__GongStructShape__000002_Default_FileToDownload.Y = 55.000000

	//gong:ident [ref_models.FileToDownload] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_FileToDownload.Identifier = `ref_models.FileToDownload`
	__GongStructShape__000002_Default_FileToDownload.IdentifierMeta = ref_models.FileToDownload{}
	__GongStructShape__000002_Default_FileToDownload.ShowNbInstances = false
	__GongStructShape__000002_Default_FileToDownload.NbInstances = 0
	__GongStructShape__000002_Default_FileToDownload.Width = 240.000000
	__GongStructShape__000002_Default_FileToDownload.Height = 103.000000
	__GongStructShape__000002_Default_FileToDownload.IsSelected = false

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Message)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_FileToUpload)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_FileToDownload)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_21T06_22_15Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_21T06_22_15Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_21T06_22_15Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Message.AttributeShapes = append(__GongStructShape__000000_Default_Message.AttributeShapes, __AttributeShape__000004_Name)
	__GongStructShape__000001_Default_FileToUpload.AttributeShapes = append(__GongStructShape__000001_Default_FileToUpload.AttributeShapes, __AttributeShape__000002_Name)
	__GongStructShape__000001_Default_FileToUpload.AttributeShapes = append(__GongStructShape__000001_Default_FileToUpload.AttributeShapes, __AttributeShape__000003_Base64EncodedContent)
	__GongStructShape__000002_Default_FileToDownload.AttributeShapes = append(__GongStructShape__000002_Default_FileToDownload.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000002_Default_FileToDownload.AttributeShapes = append(__GongStructShape__000002_Default_FileToDownload.AttributeShapes, __AttributeShape__000001_Content)
}
