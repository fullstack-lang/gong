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

// _ point for meta package dummy declaration
var ___dummy__ref_models_Default ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
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

	"ref_models.Astruct.TextArea": (ref_models.Astruct{}).TextArea,

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

	"ref_models.Dstruct.Anarrayofb": (ref_models.Dstruct{}).Anarrayofb,

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
// but it has to properly reference the _ gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = _
// }

// _ will stage objects of database "Default"
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape

	// Declarations of staged instances of Link

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	// Setup of pointers
}


