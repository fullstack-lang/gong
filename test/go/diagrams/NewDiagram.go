package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.AEnumType": ref_models.AEnumType(""),

	"ref_models.Astruct": &(ref_models.Astruct{}),

	"ref_models.Astruct.Aenum": (ref_models.Astruct{}).Aenum,

	"ref_models.Astruct.Aenum_2": (ref_models.Astruct{}).Aenum_2,

	"ref_models.Astruct.AnAstruct": (ref_models.Astruct{}).AnAstruct,

	"ref_models.Astruct.Anarrayofa": (ref_models.Astruct{}).Anarrayofa,

	"ref_models.Astruct.Anarrayofb": (ref_models.Astruct{}).Anarrayofb,

	"ref_models.Astruct.Anarrayofb2Use": (ref_models.Astruct{}).Anarrayofb2Use,

	"ref_models.Astruct.AnarrayofbUse": (ref_models.Astruct{}).AnarrayofbUse,

	"ref_models.Astruct.Anotherarrayofb": (ref_models.Astruct{}).Anotherarrayofb,

	"ref_models.Astruct.Anotherassociationtob_2": (ref_models.Astruct{}).Anotherassociationtob_2,

	"ref_models.Astruct.Anotherbooleanfield": (ref_models.Astruct{}).Anotherbooleanfield,

	"ref_models.Astruct.Associationtob": (ref_models.Astruct{}).Associationtob,

	"ref_models.Astruct.Benum": (ref_models.Astruct{}).Benum,

	"ref_models.Astruct.Booleanfield": (ref_models.Astruct{}).Booleanfield,

	"ref_models.Astruct.Bstruct": (ref_models.Astruct{}).Bstruct,

	"ref_models.Astruct.Bstruct2": (ref_models.Astruct{}).Bstruct2,

	"ref_models.Astruct.CEnum": (ref_models.Astruct{}).CEnum,

	"ref_models.Astruct.CFloatfield": (ref_models.Astruct{}).CFloatfield,

	"ref_models.Astruct.CName": (ref_models.Astruct{}).CName,

	"ref_models.Astruct.ConstIdentifierValue": (ref_models.Astruct{}).ConstIdentifierValue,

	"ref_models.Astruct.Date": (ref_models.Astruct{}).Date,

	"ref_models.Astruct.Dstruct": (ref_models.Astruct{}).Dstruct,

	"ref_models.Astruct.Dstruct2": (ref_models.Astruct{}).Dstruct2,

	"ref_models.Astruct.Dstruct3": (ref_models.Astruct{}).Dstruct3,

	"ref_models.Astruct.Dstruct4": (ref_models.Astruct{}).Dstruct4,

	"ref_models.Astruct.Duration1": (ref_models.Astruct{}).Duration1,

	"ref_models.Astruct.EnumIntRef": (ref_models.Astruct{}).EnumIntRef,

	"ref_models.Astruct.EnumStringRef": (ref_models.Astruct{}).EnumStringRef,

	"ref_models.Astruct.EnumValue": (ref_models.Astruct{}).EnumValue,

	"ref_models.Astruct.FieldRef": (ref_models.Astruct{}).FieldRef,

	"ref_models.Astruct.Floatfield": (ref_models.Astruct{}).Floatfield,

	"ref_models.Astruct.Intfield": (ref_models.Astruct{}).Intfield,

	"ref_models.Astruct.Name": (ref_models.Astruct{}).Name,

	"ref_models.Astruct.StructRef": (ref_models.Astruct{}).StructRef,

	"ref_models.AstructBstruct2Use": &(ref_models.AstructBstruct2Use{}),

	"ref_models.AstructBstruct2Use.Bstrcut2": (ref_models.AstructBstruct2Use{}).Bstrcut2,

	"ref_models.AstructBstruct2Use.Name": (ref_models.AstructBstruct2Use{}).Name,

	"ref_models.AstructBstructUse": &(ref_models.AstructBstructUse{}),

	"ref_models.AstructBstructUse.Bstruct2": (ref_models.AstructBstructUse{}).Bstruct2,

	"ref_models.AstructBstructUse.Name": (ref_models.AstructBstructUse{}).Name,

	"ref_models.BENUM_VAL1": ref_models.BENUM_VAL1,

	"ref_models.BENUM_VAL2": ref_models.BENUM_VAL2,

	"ref_models.BEnumType": ref_models.BEnumType(""),

	"ref_models.Bstruct": &(ref_models.Bstruct{}),

	"ref_models.Bstruct.Floatfield": (ref_models.Bstruct{}).Floatfield,

	"ref_models.Bstruct.Floatfield2": (ref_models.Bstruct{}).Floatfield2,

	"ref_models.Bstruct.Intfield": (ref_models.Bstruct{}).Intfield,

	"ref_models.Bstruct.Name": (ref_models.Bstruct{}).Name,

	"ref_models.CENUM_VAL1": ref_models.CENUM_VAL1,

	"ref_models.CENUM_VAL2": ref_models.CENUM_VAL2,

	"ref_models.CEnumTypeInt": ref_models.CEnumTypeInt(0),

	"ref_models.Dstruct": &(ref_models.Dstruct{}),

	"ref_models.Dstruct.Name": (ref_models.Dstruct{}).Name,

	"ref_models.ENUM_VAL1": ref_models.ENUM_VAL1,

	"ref_models.ENUM_VAL2": ref_models.ENUM_VAL2,

	"ref_models.GONG__ENUM_CAST_INT": ref_models.GONG__ENUM_CAST_INT,

	"ref_models.GONG__ENUM_CAST_STRING": ref_models.GONG__ENUM_CAST_STRING,

	"ref_models.GONG__ExpressionType": ref_models.GONG__ExpressionType(""),

	"ref_models.GONG__FIELD_OR_CONST_VALUE": ref_models.GONG__FIELD_OR_CONST_VALUE,

	"ref_models.GONG__FIELD_VALUE": ref_models.GONG__FIELD_VALUE,

	"ref_models.GONG__IDENTIFIER_CONST": ref_models.GONG__IDENTIFIER_CONST,

	"ref_models.GONG__STRUCT_INSTANCE": ref_models.GONG__STRUCT_INSTANCE,

	"ref_models.NoteOnOrganisation": ref_models.NoteOnOrganisation,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Aenum := (&models.Field{Name: `Aenum`}).Stage(stage)
	__Field__000001_Aenum_2 := (&models.Field{Name: `Aenum_2`}).Stage(stage)
	__Field__000002_Date := (&models.Field{Name: `Date`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Astruct := (&models.GongStructShape{Name: `NewDiagram-Astruct`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Bstruct := (&models.GongStructShape{Name: `NewDiagram-Bstruct`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Anarrayofb := (&models.Link{Name: `Anarrayofb`}).Stage(stage)
	__Link__000001_Bstruct := (&models.Link{Name: `Bstruct`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Astruct := (&models.Position{Name: `Pos-NewDiagram-Astruct`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_Bstruct := (&models.Position{Name: `Pos-NewDiagram-Bstruct`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Astruct and NewDiagram-Bstruct`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Astruct and NewDiagram-Bstruct`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_Aenum.Name = `Aenum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct.Aenum]
	__Field__000000_Aenum.Identifier = `ref_models.Astruct.Aenum`
	__Field__000000_Aenum.FieldTypeAsString = ``
	__Field__000000_Aenum.Structname = `Astruct`
	__Field__000000_Aenum.Fieldtypename = `AEnumType`

	// Field values setup
	__Field__000001_Aenum_2.Name = `Aenum_2`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct.Aenum_2]
	__Field__000001_Aenum_2.Identifier = `ref_models.Astruct.Aenum_2`
	__Field__000001_Aenum_2.FieldTypeAsString = ``
	__Field__000001_Aenum_2.Structname = `Astruct`
	__Field__000001_Aenum_2.Fieldtypename = `AEnumType`

	// Field values setup
	__Field__000002_Date.Name = `Date`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct.Date]
	__Field__000002_Date.Identifier = `ref_models.Astruct.Date`
	__Field__000002_Date.FieldTypeAsString = ``
	__Field__000002_Date.Structname = `Astruct`
	__Field__000002_Date.Fieldtypename = `Time`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Astruct.Name = `NewDiagram-Astruct`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct]
	__GongStructShape__000000_NewDiagram_Astruct.Identifier = `ref_models.Astruct`
	__GongStructShape__000000_NewDiagram_Astruct.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Astruct.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Astruct.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Astruct.Heigth = 108.000000
	__GongStructShape__000000_NewDiagram_Astruct.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Bstruct.Name = `NewDiagram-Bstruct`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bstruct]
	__GongStructShape__000001_NewDiagram_Bstruct.Identifier = `ref_models.Bstruct`
	__GongStructShape__000001_NewDiagram_Bstruct.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Bstruct.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Bstruct.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Bstruct.Heigth = 63.000000
	__GongStructShape__000001_NewDiagram_Bstruct.IsSelected = false

	// Link values setup
	__Link__000000_Anarrayofb.Name = `Anarrayofb`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct.Anarrayofb]
	__Link__000000_Anarrayofb.Identifier = `ref_models.Astruct.Anarrayofb`
	__Link__000000_Anarrayofb.Fieldtypename = `Bstruct`
	__Link__000000_Anarrayofb.TargetMultiplicity = models.MANY
	__Link__000000_Anarrayofb.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000001_Bstruct.Name = `Bstruct`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Astruct.Bstruct]
	__Link__000001_Bstruct.Identifier = `ref_models.Astruct.Bstruct`
	__Link__000001_Bstruct.Fieldtypename = `Bstruct`
	__Link__000001_Bstruct.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Bstruct.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Pos_NewDiagram_Astruct.X = 70.000000
	__Position__000000_Pos_NewDiagram_Astruct.Y = 104.000000
	__Position__000000_Pos_NewDiagram_Astruct.Name = `Pos-NewDiagram-Astruct`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Bstruct.X = 430.000000
	__Position__000001_Pos_NewDiagram_Bstruct.Y = 180.000000
	__Position__000001_Pos_NewDiagram_Bstruct.Name = `Pos-NewDiagram-Bstruct`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.X = 433.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.Y = 110.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Astruct and NewDiagram-Bstruct`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.X = 350.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.Y = 298.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Astruct and NewDiagram-Bstruct`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Astruct)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Bstruct)
	__GongStructShape__000000_NewDiagram_Astruct.Position = __Position__000000_Pos_NewDiagram_Astruct
	__GongStructShape__000000_NewDiagram_Astruct.Fields = append(__GongStructShape__000000_NewDiagram_Astruct.Fields, __Field__000002_Date)
	__GongStructShape__000000_NewDiagram_Astruct.Fields = append(__GongStructShape__000000_NewDiagram_Astruct.Fields, __Field__000000_Aenum)
	__GongStructShape__000000_NewDiagram_Astruct.Fields = append(__GongStructShape__000000_NewDiagram_Astruct.Fields, __Field__000001_Aenum_2)
	__GongStructShape__000000_NewDiagram_Astruct.Links = append(__GongStructShape__000000_NewDiagram_Astruct.Links, __Link__000000_Anarrayofb)
	__GongStructShape__000000_NewDiagram_Astruct.Links = append(__GongStructShape__000000_NewDiagram_Astruct.Links, __Link__000001_Bstruct)
	__GongStructShape__000001_NewDiagram_Bstruct.Position = __Position__000001_Pos_NewDiagram_Bstruct
	__Link__000000_Anarrayofb.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct
	__Link__000001_Bstruct.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Astruct_and_NewDiagram_Bstruct
}
