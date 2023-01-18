package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default models.StageStruct
var ___dummy__Time_Default time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default map[string]any = map[string]any{
	// injection point for docLink to identifiers
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0000 := (&models.Classshape{Name: `Classshape0000`}).Stage()
	__Classshape__000001_Classshape0001 := (&models.Classshape{Name: `Classshape0001`}).Stage()

	// Declarations of staged instances of DiagramPackage
	__DiagramPackage__000000_test_diagrams := (&models.DiagramPackage{Name: `test_diagrams`}).Stage()

	// Declarations of staged instances of Field
	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Anotherarrayofb := (&models.Link{Name: `Anotherarrayofb`}).Stage()
	__Link__000001_Associationtob := (&models.Link{Name: `Associationtob`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Position_0000 := (&models.Position{Name: `Position-0000`}).Stage()
	__Position__000001_Position_0001 := (&models.Position{Name: `Position-0001`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Astruct := (&models.Reference{Name: `Astruct`}).Stage()
	__Reference__000001_Bstruct := (&models.Reference{Name: `Bstruct`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0000 := (&models.Vertice{Name: `Vertice-0000`}).Stage()
	__Vertice__000001_Vertice_0001 := (&models.Vertice{Name: `Vertice-0001`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	// Classshape values setup
	__Classshape__000000_Classshape0000.Name = `Classshape0000`
	__Classshape__000000_Classshape0000.ReferenceName = `Astruct`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct]
	__Classshape__000000_Classshape0000.Identifier = `ref_models.Astruct`
	__Classshape__000000_Classshape0000.ShowNbInstances = true
	__Classshape__000000_Classshape0000.NbInstances = 0
	__Classshape__000000_Classshape0000.Width = 240.000000
	__Classshape__000000_Classshape0000.Heigth = 78.000000
	__Classshape__000000_Classshape0000.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0001.Name = `Classshape0001`
	__Classshape__000001_Classshape0001.ReferenceName = `Bstruct`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bstruct]
	__Classshape__000001_Classshape0001.Identifier = `ref_models.Bstruct`
	__Classshape__000001_Classshape0001.ShowNbInstances = true
	__Classshape__000001_Classshape0001.NbInstances = 0
	__Classshape__000001_Classshape0001.Width = 240.000000
	__Classshape__000001_Classshape0001.Heigth = 63.000000
	__Classshape__000001_Classshape0001.IsSelected = false

	// DiagramPackage values setup
	__DiagramPackage__000000_test_diagrams.Name = `test_diagrams`
	__DiagramPackage__000000_test_diagrams.Path = `github.com/fullstack-lang/gong/test/go/models`
	__DiagramPackage__000000_test_diagrams.GongModelPath = `github.com/fullstack-lang/gong/test/go/models`
	__DiagramPackage__000000_test_diagrams.IsEditable = true
	__DiagramPackage__000000_test_diagrams.IsReloaded = false
	__DiagramPackage__000000_test_diagrams.AbsolutePathToDiagramPackage = `/Users/thomaspeugeot/go/src/github.com/fullstack-lang/gong/test/go/diagrams`

	// Field values setup
	__Field__000000_Name.Name = `Name`
	__Field__000000_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct.Name]
	__Field__000000_Name.Identifier = `ref_models.Astruct.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Astruct`
	__Field__000000_Name.Fieldtypename = `string`

	// Link values setup
	__Link__000000_Anotherarrayofb.Name = `Anotherarrayofb`
	__Link__000000_Anotherarrayofb.Fieldname = `Anotherarrayofb`
	__Link__000000_Anotherarrayofb.Structname = `Astruct`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct.Anotherarrayofb]
	__Link__000000_Anotherarrayofb.Identifier = `ref_models.Astruct.Anotherarrayofb`
	__Link__000000_Anotherarrayofb.Fieldtypename = `Bstruct`
	__Link__000000_Anotherarrayofb.TargetMultiplicity = models.MANY
	__Link__000000_Anotherarrayofb.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000001_Associationtob.Name = `Associationtob`
	__Link__000001_Associationtob.Fieldname = `Associationtob`
	__Link__000001_Associationtob.Structname = `Astruct`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct.Associationtob]
	__Link__000001_Associationtob.Identifier = `ref_models.Astruct.Associationtob`
	__Link__000001_Associationtob.Fieldtypename = `Bstruct`
	__Link__000001_Associationtob.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Associationtob.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Position_0000.X = 52.000000
	__Position__000000_Position_0000.Y = 78.000000
	__Position__000000_Position_0000.Name = `Position-0000`

	// Position values setup
	__Position__000001_Position_0001.X = 460.000000
	__Position__000001_Position_0001.Y = 130.000000
	__Position__000001_Position_0001.Name = `Position-0001`

	// Reference values setup
	__Reference__000000_Astruct.Name = `Astruct`
	__Reference__000000_Astruct.NbInstances = 0
	__Reference__000000_Astruct.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_Bstruct.Name = `Bstruct`
	__Reference__000001_Bstruct.NbInstances = 0
	__Reference__000001_Bstruct.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Vertice_0000.X = 551.000000
	__Vertice__000000_Vertice_0000.Y = 23.000000
	__Vertice__000000_Vertice_0000.Name = `Vertice-0000`

	// Vertice values setup
	__Vertice__000001_Vertice_0001.X = 291.000000
	__Vertice__000001_Vertice_0001.Y = 233.000000
	__Vertice__000001_Vertice_0001.Name = `Vertice-0001`

	// Setup of pointers
	__Classdiagram__000000_Default.Classshapes = append(__Classdiagram__000000_Default.Classshapes, __Classshape__000000_Classshape0000)
	__Classdiagram__000000_Default.Classshapes = append(__Classdiagram__000000_Default.Classshapes, __Classshape__000001_Classshape0001)
	__Classshape__000000_Classshape0000.Position = __Position__000000_Position_0000
	__Classshape__000000_Classshape0000.Reference = __Reference__000000_Astruct
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000000_Name)
	__Classshape__000000_Classshape0000.Links = append(__Classshape__000000_Classshape0000.Links, __Link__000000_Anotherarrayofb)
	__Classshape__000000_Classshape0000.Links = append(__Classshape__000000_Classshape0000.Links, __Link__000001_Associationtob)
	__Classshape__000001_Classshape0001.Position = __Position__000001_Position_0001
	__Classshape__000001_Classshape0001.Reference = __Reference__000001_Bstruct
	__DiagramPackage__000000_test_diagrams.Classdiagrams = append(__DiagramPackage__000000_test_diagrams.Classdiagrams, __Classdiagram__000000_Default)
	__Link__000000_Anotherarrayofb.Middlevertice = __Vertice__000000_Vertice_0000
	__Link__000001_Associationtob.Middlevertice = __Vertice__000001_Vertice_0001
}


