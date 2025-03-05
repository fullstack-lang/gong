package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/go/models"
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

	"ref_models.GongBasicField": &(ref_models.GongBasicField{}),

	"ref_models.GongBasicField.BasicKindName": (ref_models.GongBasicField{}).BasicKindName,

	"ref_models.GongBasicField.BespokeHeight": (ref_models.GongBasicField{}).BespokeHeight,

	"ref_models.GongBasicField.BespokeWidth": (ref_models.GongBasicField{}).BespokeWidth,

	"ref_models.GongBasicField.CompositeStructName": (ref_models.GongBasicField{}).CompositeStructName,

	"ref_models.GongBasicField.DeclaredType": (ref_models.GongBasicField{}).DeclaredType,

	"ref_models.GongBasicField.GongEnum": (ref_models.GongBasicField{}).GongEnum,

	"ref_models.GongBasicField.Index": (ref_models.GongBasicField{}).Index,

	"ref_models.GongBasicField.IsBespokeHeight": (ref_models.GongBasicField{}).IsBespokeHeight,

	"ref_models.GongBasicField.IsBespokeWidth": (ref_models.GongBasicField{}).IsBespokeWidth,

	"ref_models.GongBasicField.IsDocLink": (ref_models.GongBasicField{}).IsDocLink,

	"ref_models.GongBasicField.IsTextArea": (ref_models.GongBasicField{}).IsTextArea,

	"ref_models.GongBasicField.Name": (ref_models.GongBasicField{}).Name,

	"ref_models.GongEnum": &(ref_models.GongEnum{}),

	"ref_models.GongEnum.GongEnumValues": (ref_models.GongEnum{}).GongEnumValues,

	"ref_models.GongEnum.Name": (ref_models.GongEnum{}).Name,

	"ref_models.GongEnum.Type": (ref_models.GongEnum{}).Type,

	"ref_models.GongEnumType": ref_models.GongEnumType(0),

	"ref_models.GongEnumValue": &(ref_models.GongEnumValue{}),

	"ref_models.GongEnumValue.Name": (ref_models.GongEnumValue{}).Name,

	"ref_models.GongEnumValue.Value": (ref_models.GongEnumValue{}).Value,

	"ref_models.GongLink": &(ref_models.GongLink{}),

	"ref_models.GongLink.ImportPath": (ref_models.GongLink{}).ImportPath,

	"ref_models.GongLink.Name": (ref_models.GongLink{}).Name,

	"ref_models.GongLink.Recv": (ref_models.GongLink{}).Recv,

	"ref_models.GongNote": &(ref_models.GongNote{}),

	"ref_models.GongNote.Body": (ref_models.GongNote{}).Body,

	"ref_models.GongNote.BodyHTML": (ref_models.GongNote{}).BodyHTML,

	"ref_models.GongNote.Links": (ref_models.GongNote{}).Links,

	"ref_models.GongNote.Name": (ref_models.GongNote{}).Name,

	"ref_models.GongStruct": &(ref_models.GongStruct{}),

	"ref_models.GongStruct.GongBasicFields": (ref_models.GongStruct{}).GongBasicFields,

	"ref_models.GongStruct.GongTimeFields": (ref_models.GongStruct{}).GongTimeFields,

	"ref_models.GongStruct.HasOnAfterUpdateSignature": (ref_models.GongStruct{}).HasOnAfterUpdateSignature,

	"ref_models.GongStruct.IsIgnoredForFront": (ref_models.GongStruct{}).IsIgnoredForFront,

	"ref_models.GongStruct.Name": (ref_models.GongStruct{}).Name,

	"ref_models.GongStruct.PointerToGongStructFields": (ref_models.GongStruct{}).PointerToGongStructFields,

	"ref_models.GongStruct.SliceOfPointerToGongStructFields": (ref_models.GongStruct{}).SliceOfPointerToGongStructFields,

	"ref_models.GongTimeField": &(ref_models.GongTimeField{}),

	"ref_models.GongTimeField.CompositeStructName": (ref_models.GongTimeField{}).CompositeStructName,

	"ref_models.GongTimeField.Index": (ref_models.GongTimeField{}).Index,

	"ref_models.GongTimeField.Name": (ref_models.GongTimeField{}).Name,

	"ref_models.Int": ref_models.Int,

	"ref_models.Meta": &(ref_models.Meta{}),

	"ref_models.Meta.MetaReferences": (ref_models.Meta{}).MetaReferences,

	"ref_models.Meta.Name": (ref_models.Meta{}).Name,

	"ref_models.Meta.Text": (ref_models.Meta{}).Text,

	"ref_models.MetaReference": &(ref_models.MetaReference{}),

	"ref_models.MetaReference.Name": (ref_models.MetaReference{}).Name,

	"ref_models.ModelPkg": &(ref_models.ModelPkg{}),

	"ref_models.ModelPkg.ControllersPkgGenPath": (ref_models.ModelPkg{}).ControllersPkgGenPath,

	"ref_models.ModelPkg.FullstackPkgGenPath": (ref_models.ModelPkg{}).FullstackPkgGenPath,

	"ref_models.ModelPkg.MaterialLibDatamodelTargetPath": (ref_models.ModelPkg{}).MaterialLibDatamodelTargetPath,

	"ref_models.ModelPkg.Name": (ref_models.ModelPkg{}).Name,

	"ref_models.ModelPkg.NgDataLibrarySourceCodeDirectory": (ref_models.ModelPkg{}).NgDataLibrarySourceCodeDirectory,

	"ref_models.ModelPkg.NgSpecificLibrarySourceCodeDirectory": (ref_models.ModelPkg{}).NgSpecificLibrarySourceCodeDirectory,

	"ref_models.ModelPkg.NgWorkspaceName": (ref_models.ModelPkg{}).NgWorkspaceName,

	"ref_models.ModelPkg.NgWorkspacePath": (ref_models.ModelPkg{}).NgWorkspacePath,

	"ref_models.ModelPkg.OrmPkgGenPath": (ref_models.ModelPkg{}).OrmPkgGenPath,

	"ref_models.ModelPkg.PathToGoSubDirectory": (ref_models.ModelPkg{}).PathToGoSubDirectory,

	"ref_models.ModelPkg.PkgPath": (ref_models.ModelPkg{}).PkgPath,

	"ref_models.ModelPkg.ProbePkgGenPath": (ref_models.ModelPkg{}).ProbePkgGenPath,

	"ref_models.ModelPkg.StackPkgGenPath": (ref_models.ModelPkg{}).StackPkgGenPath,

	"ref_models.ModelPkg.StaticPkgGenPath": (ref_models.ModelPkg{}).StaticPkgGenPath,

	"ref_models.PointerToGongStructField": &(ref_models.PointerToGongStructField{}),

	"ref_models.PointerToGongStructField.CompositeStructName": (ref_models.PointerToGongStructField{}).CompositeStructName,

	"ref_models.PointerToGongStructField.GongStruct": (ref_models.PointerToGongStructField{}).GongStruct,

	"ref_models.PointerToGongStructField.Index": (ref_models.PointerToGongStructField{}).Index,

	"ref_models.PointerToGongStructField.IsType": (ref_models.PointerToGongStructField{}).IsType,

	"ref_models.PointerToGongStructField.Name": (ref_models.PointerToGongStructField{}).Name,

	"ref_models.SliceOfPointerToGongStructField": &(ref_models.SliceOfPointerToGongStructField{}),

	"ref_models.SliceOfPointerToGongStructField.CompositeStructName": (ref_models.SliceOfPointerToGongStructField{}).CompositeStructName,

	"ref_models.SliceOfPointerToGongStructField.GongStruct": (ref_models.SliceOfPointerToGongStructField{}).GongStruct,

	"ref_models.SliceOfPointerToGongStructField.Index": (ref_models.SliceOfPointerToGongStructField{}).Index,

	"ref_models.SliceOfPointerToGongStructField.Name": (ref_models.SliceOfPointerToGongStructField{}).Name,

	"ref_models.String": ref_models.String,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_BasicKindName := (&models.Field{Name: `BasicKindName`}).Stage(stage)
	__Field__000001_BespokeHeight := (&models.Field{Name: `BespokeHeight`}).Stage(stage)
	__Field__000002_BespokeWidth := (&models.Field{Name: `BespokeWidth`}).Stage(stage)
	__Field__000003_Body := (&models.Field{Name: `Body`}).Stage(stage)
	__Field__000004_BodyHTML := (&models.Field{Name: `BodyHTML`}).Stage(stage)
	__Field__000005_CompositeStructName := (&models.Field{Name: `CompositeStructName`}).Stage(stage)
	__Field__000006_DeclaredType := (&models.Field{Name: `DeclaredType`}).Stage(stage)
	__Field__000007_HasOnAfterUpdateSignature := (&models.Field{Name: `HasOnAfterUpdateSignature`}).Stage(stage)
	__Field__000008_ImportPath := (&models.Field{Name: `ImportPath`}).Stage(stage)
	__Field__000009_Index := (&models.Field{Name: `Index`}).Stage(stage)
	__Field__000010_IsBespokeHeight := (&models.Field{Name: `IsBespokeHeight`}).Stage(stage)
	__Field__000011_IsBespokeWidth := (&models.Field{Name: `IsBespokeWidth`}).Stage(stage)
	__Field__000012_IsDocLink := (&models.Field{Name: `IsDocLink`}).Stage(stage)
	__Field__000013_IsTextArea := (&models.Field{Name: `IsTextArea`}).Stage(stage)
	__Field__000014_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000015_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000016_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000017_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000018_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000019_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000020_Recv := (&models.Field{Name: `Recv`}).Stage(stage)
	__Field__000021_Type := (&models.Field{Name: `Type`}).Stage(stage)
	__Field__000022_Value := (&models.Field{Name: `Value`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_GongBasicField := (&models.GongStructShape{Name: `Default-GongBasicField`}).Stage(stage)
	__GongStructShape__000001_Default_GongEnum := (&models.GongStructShape{Name: `Default-GongEnum`}).Stage(stage)
	__GongStructShape__000002_Default_GongEnumValue := (&models.GongStructShape{Name: `Default-GongEnumValue`}).Stage(stage)
	__GongStructShape__000003_Default_GongLink := (&models.GongStructShape{Name: `Default-GongLink`}).Stage(stage)
	__GongStructShape__000004_Default_GongNote := (&models.GongStructShape{Name: `Default-GongNote`}).Stage(stage)
	__GongStructShape__000005_Default_GongStruct := (&models.GongStructShape{Name: `Default-GongStruct`}).Stage(stage)
	__GongStructShape__000006_Default_GongTimeField := (&models.GongStructShape{Name: `Default-GongTimeField`}).Stage(stage)
	__GongStructShape__000007_Default_ModelPkg := (&models.GongStructShape{Name: `Default-ModelPkg`}).Stage(stage)
	__GongStructShape__000008_Default_PointerToGongStructField := (&models.GongStructShape{Name: `Default-PointerToGongStructField`}).Stage(stage)
	__GongStructShape__000009_Default_SliceOfPointerToGongStructField := (&models.GongStructShape{Name: `Default-SliceOfPointerToGongStructField`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_GongBasicFields := (&models.Link{Name: `GongBasicFields`}).Stage(stage)
	__Link__000001_GongEnum := (&models.Link{Name: `GongEnum`}).Stage(stage)
	__Link__000002_GongEnumValues := (&models.Link{Name: `GongEnumValues`}).Stage(stage)
	__Link__000003_Links := (&models.Link{Name: `Links`}).Stage(stage)
	__Link__000004_PointerToGongStructFields := (&models.Link{Name: `PointerToGongStructFields`}).Stage(stage)
	__Link__000005_SliceOfPointerToGongStructFields := (&models.Link{Name: `SliceOfPointerToGongStructFields`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_GongBasicField := (&models.Position{Name: `Pos-Default-GongBasicField`}).Stage(stage)
	__Position__000001_Pos_Default_GongEnum := (&models.Position{Name: `Pos-Default-GongEnum`}).Stage(stage)
	__Position__000002_Pos_Default_GongEnumValue := (&models.Position{Name: `Pos-Default-GongEnumValue`}).Stage(stage)
	__Position__000003_Pos_Default_GongLink := (&models.Position{Name: `Pos-Default-GongLink`}).Stage(stage)
	__Position__000004_Pos_Default_GongNote := (&models.Position{Name: `Pos-Default-GongNote`}).Stage(stage)
	__Position__000005_Pos_Default_GongStruct := (&models.Position{Name: `Pos-Default-GongStruct`}).Stage(stage)
	__Position__000006_Pos_Default_GongTimeField := (&models.Position{Name: `Pos-Default-GongTimeField`}).Stage(stage)
	__Position__000007_Pos_Default_ModelPkg := (&models.Position{Name: `Pos-Default-ModelPkg`}).Stage(stage)
	__Position__000008_Pos_Default_PointerToGongStructField := (&models.Position{Name: `Pos-Default-PointerToGongStructField`}).Stage(stage)
	__Position__000009_Pos_Default_SliceOfPointerToGongStructField := (&models.Position{Name: `Pos-Default-SliceOfPointerToGongStructField`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_GongBasicField_and_Default_GongEnum := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-GongBasicField and Default-GongEnum`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_GongEnum_and_Default_GongEnumValue := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-GongEnum and Default-GongEnumValue`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_GongNote_and_Default_GongLink := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-GongNote and Default-GongLink`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_GongBasicField := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-GongStruct and Default-GongBasicField`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_PointerToGongStructField := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-GongStruct and Default-PointerToGongStructField`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_SliceOfPointerToGongStructField := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-GongStruct and Default-SliceOfPointerToGongStructField`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	// Field values setup
	__Field__000000_BasicKindName.Name = `BasicKindName`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.BasicKindName]
	__Field__000000_BasicKindName.Identifier = `ref_models.GongBasicField.BasicKindName`
	__Field__000000_BasicKindName.FieldTypeAsString = ``
	__Field__000000_BasicKindName.Structname = `GongBasicField`
	__Field__000000_BasicKindName.Fieldtypename = `string`

	// Field values setup
	__Field__000001_BespokeHeight.Name = `BespokeHeight`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.BespokeHeight]
	__Field__000001_BespokeHeight.Identifier = `ref_models.GongBasicField.BespokeHeight`
	__Field__000001_BespokeHeight.FieldTypeAsString = ``
	__Field__000001_BespokeHeight.Structname = `GongBasicField`
	__Field__000001_BespokeHeight.Fieldtypename = `int`

	// Field values setup
	__Field__000002_BespokeWidth.Name = `BespokeWidth`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.BespokeWidth]
	__Field__000002_BespokeWidth.Identifier = `ref_models.GongBasicField.BespokeWidth`
	__Field__000002_BespokeWidth.FieldTypeAsString = ``
	__Field__000002_BespokeWidth.Structname = `GongBasicField`
	__Field__000002_BespokeWidth.Fieldtypename = `int`

	// Field values setup
	__Field__000003_Body.Name = `Body`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongNote.Body]
	__Field__000003_Body.Identifier = `ref_models.GongNote.Body`
	__Field__000003_Body.FieldTypeAsString = ``
	__Field__000003_Body.Structname = `GongNote`
	__Field__000003_Body.Fieldtypename = `string`

	// Field values setup
	__Field__000004_BodyHTML.Name = `BodyHTML`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongNote.BodyHTML]
	__Field__000004_BodyHTML.Identifier = `ref_models.GongNote.BodyHTML`
	__Field__000004_BodyHTML.FieldTypeAsString = ``
	__Field__000004_BodyHTML.Structname = `GongNote`
	__Field__000004_BodyHTML.Fieldtypename = `string`

	// Field values setup
	__Field__000005_CompositeStructName.Name = `CompositeStructName`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.CompositeStructName]
	__Field__000005_CompositeStructName.Identifier = `ref_models.GongBasicField.CompositeStructName`
	__Field__000005_CompositeStructName.FieldTypeAsString = ``
	__Field__000005_CompositeStructName.Structname = `GongBasicField`
	__Field__000005_CompositeStructName.Fieldtypename = `string`

	// Field values setup
	__Field__000006_DeclaredType.Name = `DeclaredType`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.DeclaredType]
	__Field__000006_DeclaredType.Identifier = `ref_models.GongBasicField.DeclaredType`
	__Field__000006_DeclaredType.FieldTypeAsString = ``
	__Field__000006_DeclaredType.Structname = `GongBasicField`
	__Field__000006_DeclaredType.Fieldtypename = `string`

	// Field values setup
	__Field__000007_HasOnAfterUpdateSignature.Name = `HasOnAfterUpdateSignature`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStruct.HasOnAfterUpdateSignature]
	__Field__000007_HasOnAfterUpdateSignature.Identifier = `ref_models.GongStruct.HasOnAfterUpdateSignature`
	__Field__000007_HasOnAfterUpdateSignature.FieldTypeAsString = ``
	__Field__000007_HasOnAfterUpdateSignature.Structname = `GongStruct`
	__Field__000007_HasOnAfterUpdateSignature.Fieldtypename = `bool`

	// Field values setup
	__Field__000008_ImportPath.Name = `ImportPath`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongLink.ImportPath]
	__Field__000008_ImportPath.Identifier = `ref_models.GongLink.ImportPath`
	__Field__000008_ImportPath.FieldTypeAsString = ``
	__Field__000008_ImportPath.Structname = `GongLink`
	__Field__000008_ImportPath.Fieldtypename = `string`

	// Field values setup
	__Field__000009_Index.Name = `Index`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.Index]
	__Field__000009_Index.Identifier = `ref_models.GongBasicField.Index`
	__Field__000009_Index.FieldTypeAsString = ``
	__Field__000009_Index.Structname = `GongBasicField`
	__Field__000009_Index.Fieldtypename = `int`

	// Field values setup
	__Field__000010_IsBespokeHeight.Name = `IsBespokeHeight`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.IsBespokeHeight]
	__Field__000010_IsBespokeHeight.Identifier = `ref_models.GongBasicField.IsBespokeHeight`
	__Field__000010_IsBespokeHeight.FieldTypeAsString = ``
	__Field__000010_IsBespokeHeight.Structname = `GongBasicField`
	__Field__000010_IsBespokeHeight.Fieldtypename = `bool`

	// Field values setup
	__Field__000011_IsBespokeWidth.Name = `IsBespokeWidth`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.IsBespokeWidth]
	__Field__000011_IsBespokeWidth.Identifier = `ref_models.GongBasicField.IsBespokeWidth`
	__Field__000011_IsBespokeWidth.FieldTypeAsString = ``
	__Field__000011_IsBespokeWidth.Structname = `GongBasicField`
	__Field__000011_IsBespokeWidth.Fieldtypename = `bool`

	// Field values setup
	__Field__000012_IsDocLink.Name = `IsDocLink`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.IsDocLink]
	__Field__000012_IsDocLink.Identifier = `ref_models.GongBasicField.IsDocLink`
	__Field__000012_IsDocLink.FieldTypeAsString = ``
	__Field__000012_IsDocLink.Structname = `GongBasicField`
	__Field__000012_IsDocLink.Fieldtypename = `bool`

	// Field values setup
	__Field__000013_IsTextArea.Name = `IsTextArea`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.IsTextArea]
	__Field__000013_IsTextArea.Identifier = `ref_models.GongBasicField.IsTextArea`
	__Field__000013_IsTextArea.FieldTypeAsString = ``
	__Field__000013_IsTextArea.Structname = `GongBasicField`
	__Field__000013_IsTextArea.Fieldtypename = `bool`

	// Field values setup
	__Field__000014_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongNote.Name]
	__Field__000014_Name.Identifier = `ref_models.GongNote.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `GongNote`
	__Field__000014_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000015_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.Name]
	__Field__000015_Name.Identifier = `ref_models.GongBasicField.Name`
	__Field__000015_Name.FieldTypeAsString = ``
	__Field__000015_Name.Structname = `GongBasicField`
	__Field__000015_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000016_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStruct.Name]
	__Field__000016_Name.Identifier = `ref_models.GongStruct.Name`
	__Field__000016_Name.FieldTypeAsString = ``
	__Field__000016_Name.Structname = `GongStruct`
	__Field__000016_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000017_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongEnum.Name]
	__Field__000017_Name.Identifier = `ref_models.GongEnum.Name`
	__Field__000017_Name.FieldTypeAsString = ``
	__Field__000017_Name.Structname = `GongEnum`
	__Field__000017_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000018_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongLink.Name]
	__Field__000018_Name.Identifier = `ref_models.GongLink.Name`
	__Field__000018_Name.FieldTypeAsString = ``
	__Field__000018_Name.Structname = `GongLink`
	__Field__000018_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000019_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongEnumValue.Name]
	__Field__000019_Name.Identifier = `ref_models.GongEnumValue.Name`
	__Field__000019_Name.FieldTypeAsString = ``
	__Field__000019_Name.Structname = `GongEnumValue`
	__Field__000019_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000020_Recv.Name = `Recv`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongLink.Recv]
	__Field__000020_Recv.Identifier = `ref_models.GongLink.Recv`
	__Field__000020_Recv.FieldTypeAsString = ``
	__Field__000020_Recv.Structname = `GongLink`
	__Field__000020_Recv.Fieldtypename = `string`

	// Field values setup
	__Field__000021_Type.Name = `Type`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongEnum.Type]
	__Field__000021_Type.Identifier = `ref_models.GongEnum.Type`
	__Field__000021_Type.FieldTypeAsString = ``
	__Field__000021_Type.Structname = `GongEnum`
	__Field__000021_Type.Fieldtypename = `GongEnumType`

	// Field values setup
	__Field__000022_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongEnumValue.Value]
	__Field__000022_Value.Identifier = `ref_models.GongEnumValue.Value`
	__Field__000022_Value.FieldTypeAsString = ``
	__Field__000022_Value.Structname = `GongEnumValue`
	__Field__000022_Value.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_Default_GongBasicField.Name = `Default-GongBasicField`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField]
	__GongStructShape__000000_Default_GongBasicField.Identifier = `ref_models.GongBasicField`
	__GongStructShape__000000_Default_GongBasicField.ShowNbInstances = false
	__GongStructShape__000000_Default_GongBasicField.NbInstances = 0
	__GongStructShape__000000_Default_GongBasicField.Width = 284.000000
	__GongStructShape__000000_Default_GongBasicField.Height = 228.000000
	__GongStructShape__000000_Default_GongBasicField.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Default_GongEnum.Name = `Default-GongEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongEnum]
	__GongStructShape__000001_Default_GongEnum.Identifier = `ref_models.GongEnum`
	__GongStructShape__000001_Default_GongEnum.ShowNbInstances = false
	__GongStructShape__000001_Default_GongEnum.NbInstances = 0
	__GongStructShape__000001_Default_GongEnum.Width = 240.000000
	__GongStructShape__000001_Default_GongEnum.Height = 93.000000
	__GongStructShape__000001_Default_GongEnum.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Default_GongEnumValue.Name = `Default-GongEnumValue`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongEnumValue]
	__GongStructShape__000002_Default_GongEnumValue.Identifier = `ref_models.GongEnumValue`
	__GongStructShape__000002_Default_GongEnumValue.ShowNbInstances = false
	__GongStructShape__000002_Default_GongEnumValue.NbInstances = 0
	__GongStructShape__000002_Default_GongEnumValue.Width = 240.000000
	__GongStructShape__000002_Default_GongEnumValue.Height = 93.000000
	__GongStructShape__000002_Default_GongEnumValue.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Default_GongLink.Name = `Default-GongLink`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongLink]
	__GongStructShape__000003_Default_GongLink.Identifier = `ref_models.GongLink`
	__GongStructShape__000003_Default_GongLink.ShowNbInstances = false
	__GongStructShape__000003_Default_GongLink.NbInstances = 0
	__GongStructShape__000003_Default_GongLink.Width = 293.000000
	__GongStructShape__000003_Default_GongLink.Height = 108.000000
	__GongStructShape__000003_Default_GongLink.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_Default_GongNote.Name = `Default-GongNote`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongNote]
	__GongStructShape__000004_Default_GongNote.Identifier = `ref_models.GongNote`
	__GongStructShape__000004_Default_GongNote.ShowNbInstances = false
	__GongStructShape__000004_Default_GongNote.NbInstances = 0
	__GongStructShape__000004_Default_GongNote.Width = 240.000000
	__GongStructShape__000004_Default_GongNote.Height = 108.000000
	__GongStructShape__000004_Default_GongNote.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_Default_GongStruct.Name = `Default-GongStruct`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStruct]
	__GongStructShape__000005_Default_GongStruct.Identifier = `ref_models.GongStruct`
	__GongStructShape__000005_Default_GongStruct.ShowNbInstances = false
	__GongStructShape__000005_Default_GongStruct.NbInstances = 0
	__GongStructShape__000005_Default_GongStruct.Width = 240.000000
	__GongStructShape__000005_Default_GongStruct.Height = 93.000000
	__GongStructShape__000005_Default_GongStruct.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_Default_GongTimeField.Name = `Default-GongTimeField`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongTimeField]
	__GongStructShape__000006_Default_GongTimeField.Identifier = `ref_models.GongTimeField`
	__GongStructShape__000006_Default_GongTimeField.ShowNbInstances = false
	__GongStructShape__000006_Default_GongTimeField.NbInstances = 0
	__GongStructShape__000006_Default_GongTimeField.Width = 240.000000
	__GongStructShape__000006_Default_GongTimeField.Height = 63.000000
	__GongStructShape__000006_Default_GongTimeField.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_Default_ModelPkg.Name = `Default-ModelPkg`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ModelPkg]
	__GongStructShape__000007_Default_ModelPkg.Identifier = `ref_models.ModelPkg`
	__GongStructShape__000007_Default_ModelPkg.ShowNbInstances = false
	__GongStructShape__000007_Default_ModelPkg.NbInstances = 0
	__GongStructShape__000007_Default_ModelPkg.Width = 240.000000
	__GongStructShape__000007_Default_ModelPkg.Height = 63.000000
	__GongStructShape__000007_Default_ModelPkg.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000008_Default_PointerToGongStructField.Name = `Default-PointerToGongStructField`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.PointerToGongStructField]
	__GongStructShape__000008_Default_PointerToGongStructField.Identifier = `ref_models.PointerToGongStructField`
	__GongStructShape__000008_Default_PointerToGongStructField.ShowNbInstances = false
	__GongStructShape__000008_Default_PointerToGongStructField.NbInstances = 0
	__GongStructShape__000008_Default_PointerToGongStructField.Width = 286.000000
	__GongStructShape__000008_Default_PointerToGongStructField.Height = 63.000000
	__GongStructShape__000008_Default_PointerToGongStructField.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000009_Default_SliceOfPointerToGongStructField.Name = `Default-SliceOfPointerToGongStructField`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SliceOfPointerToGongStructField]
	__GongStructShape__000009_Default_SliceOfPointerToGongStructField.Identifier = `ref_models.SliceOfPointerToGongStructField`
	__GongStructShape__000009_Default_SliceOfPointerToGongStructField.ShowNbInstances = false
	__GongStructShape__000009_Default_SliceOfPointerToGongStructField.NbInstances = 0
	__GongStructShape__000009_Default_SliceOfPointerToGongStructField.Width = 287.000000
	__GongStructShape__000009_Default_SliceOfPointerToGongStructField.Height = 63.000000
	__GongStructShape__000009_Default_SliceOfPointerToGongStructField.IsSelected = false

	// Link values setup
	__Link__000000_GongBasicFields.Name = `GongBasicFields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStruct.GongBasicFields]
	__Link__000000_GongBasicFields.Identifier = `ref_models.GongStruct.GongBasicFields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField]
	__Link__000000_GongBasicFields.Fieldtypename = `ref_models.GongBasicField`
	__Link__000000_GongBasicFields.FieldOffsetX = -125.000000
	__Link__000000_GongBasicFields.FieldOffsetY = -18.000000
	__Link__000000_GongBasicFields.TargetMultiplicity = models.MANY
	__Link__000000_GongBasicFields.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_GongBasicFields.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_GongBasicFields.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_GongBasicFields.SourceMultiplicityOffsetX = -67.000000
	__Link__000000_GongBasicFields.SourceMultiplicityOffsetY = 22.000000
	__Link__000000_GongBasicFields.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_GongBasicFields.StartRatio = 0.885286
	__Link__000000_GongBasicFields.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_GongBasicFields.EndRatio = 0.523872
	__Link__000000_GongBasicFields.CornerOffsetRatio = 2.258107

	// Link values setup
	__Link__000001_GongEnum.Name = `GongEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongBasicField.GongEnum]
	__Link__000001_GongEnum.Identifier = `ref_models.GongBasicField.GongEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongEnum]
	__Link__000001_GongEnum.Fieldtypename = `ref_models.GongEnum`
	__Link__000001_GongEnum.FieldOffsetX = -50.000000
	__Link__000001_GongEnum.FieldOffsetY = -16.000000
	__Link__000001_GongEnum.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_GongEnum.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_GongEnum.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_GongEnum.SourceMultiplicity = models.MANY
	__Link__000001_GongEnum.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_GongEnum.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_GongEnum.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_GongEnum.StartRatio = 0.500000
	__Link__000001_GongEnum.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_GongEnum.EndRatio = 0.500000
	__Link__000001_GongEnum.CornerOffsetRatio = 1.380000

	// Link values setup
	__Link__000002_GongEnumValues.Name = `GongEnumValues`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongEnum.GongEnumValues]
	__Link__000002_GongEnumValues.Identifier = `ref_models.GongEnum.GongEnumValues`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongEnumValue]
	__Link__000002_GongEnumValues.Fieldtypename = `ref_models.GongEnumValue`
	__Link__000002_GongEnumValues.FieldOffsetX = -50.000000
	__Link__000002_GongEnumValues.FieldOffsetY = -16.000000
	__Link__000002_GongEnumValues.TargetMultiplicity = models.MANY
	__Link__000002_GongEnumValues.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_GongEnumValues.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_GongEnumValues.SourceMultiplicity = models.MANY
	__Link__000002_GongEnumValues.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_GongEnumValues.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_GongEnumValues.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_GongEnumValues.StartRatio = 0.500000
	__Link__000002_GongEnumValues.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_GongEnumValues.EndRatio = 0.500000
	__Link__000002_GongEnumValues.CornerOffsetRatio = 1.380000

	// Link values setup
	__Link__000003_Links.Name = `Links`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongNote.Links]
	__Link__000003_Links.Identifier = `ref_models.GongNote.Links`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongLink]
	__Link__000003_Links.Fieldtypename = `ref_models.GongLink`
	__Link__000003_Links.FieldOffsetX = -50.000000
	__Link__000003_Links.FieldOffsetY = -16.000000
	__Link__000003_Links.TargetMultiplicity = models.MANY
	__Link__000003_Links.TargetMultiplicityOffsetX = -26.000000
	__Link__000003_Links.TargetMultiplicityOffsetY = 23.000000
	__Link__000003_Links.SourceMultiplicity = models.ZERO_ONE
	__Link__000003_Links.SourceMultiplicityOffsetX = 18.000000
	__Link__000003_Links.SourceMultiplicityOffsetY = -20.000000
	__Link__000003_Links.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_Links.StartRatio = 0.500000
	__Link__000003_Links.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_Links.EndRatio = 0.537073
	__Link__000003_Links.CornerOffsetRatio = 1.556120

	// Link values setup
	__Link__000004_PointerToGongStructFields.Name = `PointerToGongStructFields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStruct.PointerToGongStructFields]
	__Link__000004_PointerToGongStructFields.Identifier = `ref_models.GongStruct.PointerToGongStructFields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.PointerToGongStructField]
	__Link__000004_PointerToGongStructFields.Fieldtypename = `ref_models.PointerToGongStructField`
	__Link__000004_PointerToGongStructFields.FieldOffsetX = -193.000000
	__Link__000004_PointerToGongStructFields.FieldOffsetY = -17.000000
	__Link__000004_PointerToGongStructFields.TargetMultiplicity = models.MANY
	__Link__000004_PointerToGongStructFields.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_PointerToGongStructFields.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_PointerToGongStructFields.SourceMultiplicity = models.ZERO_ONE
	__Link__000004_PointerToGongStructFields.SourceMultiplicityOffsetX = 8.000000
	__Link__000004_PointerToGongStructFields.SourceMultiplicityOffsetY = -22.000000
	__Link__000004_PointerToGongStructFields.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_PointerToGongStructFields.StartRatio = 0.500000
	__Link__000004_PointerToGongStructFields.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_PointerToGongStructFields.EndRatio = 0.500000
	__Link__000004_PointerToGongStructFields.CornerOffsetRatio = 1.218620

	// Link values setup
	__Link__000005_SliceOfPointerToGongStructFields.Name = `SliceOfPointerToGongStructFields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongStruct.SliceOfPointerToGongStructFields]
	__Link__000005_SliceOfPointerToGongStructFields.Identifier = `ref_models.GongStruct.SliceOfPointerToGongStructFields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SliceOfPointerToGongStructField]
	__Link__000005_SliceOfPointerToGongStructFields.Fieldtypename = `ref_models.SliceOfPointerToGongStructField`
	__Link__000005_SliceOfPointerToGongStructFields.FieldOffsetX = -50.000000
	__Link__000005_SliceOfPointerToGongStructFields.FieldOffsetY = -16.000000
	__Link__000005_SliceOfPointerToGongStructFields.TargetMultiplicity = models.MANY
	__Link__000005_SliceOfPointerToGongStructFields.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_SliceOfPointerToGongStructFields.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_SliceOfPointerToGongStructFields.SourceMultiplicity = models.MANY
	__Link__000005_SliceOfPointerToGongStructFields.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_SliceOfPointerToGongStructFields.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_SliceOfPointerToGongStructFields.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000005_SliceOfPointerToGongStructFields.StartRatio = 0.888472
	__Link__000005_SliceOfPointerToGongStructFields.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_SliceOfPointerToGongStructFields.EndRatio = 0.500000
	__Link__000005_SliceOfPointerToGongStructFields.CornerOffsetRatio = 1.000898

	// Position values setup
	__Position__000000_Pos_Default_GongBasicField.X = 650.000000
	__Position__000000_Pos_Default_GongBasicField.Y = 318.000000
	__Position__000000_Pos_Default_GongBasicField.Name = `Pos-Default-GongBasicField`

	// Position values setup
	__Position__000001_Pos_Default_GongEnum.X = 1084.000061
	__Position__000001_Pos_Default_GongEnum.Y = 380.000000
	__Position__000001_Pos_Default_GongEnum.Name = `Pos-Default-GongEnum`

	// Position values setup
	__Position__000002_Pos_Default_GongEnumValue.X = 647.000000
	__Position__000002_Pos_Default_GongEnumValue.Y = 689.000000
	__Position__000002_Pos_Default_GongEnumValue.Name = `Pos-Default-GongEnumValue`

	// Position values setup
	__Position__000003_Pos_Default_GongLink.X = 652.000000
	__Position__000003_Pos_Default_GongLink.Y = 563.000000
	__Position__000003_Pos_Default_GongLink.Name = `Pos-Default-GongLink`

	// Position values setup
	__Position__000004_Pos_Default_GongNote.X = 28.000000
	__Position__000004_Pos_Default_GongNote.Y = 556.000000
	__Position__000004_Pos_Default_GongNote.Name = `Pos-Default-GongNote`

	// Position values setup
	__Position__000005_Pos_Default_GongStruct.X = 28.000000
	__Position__000005_Pos_Default_GongStruct.Y = 82.000000
	__Position__000005_Pos_Default_GongStruct.Name = `Pos-Default-GongStruct`

	// Position values setup
	__Position__000006_Pos_Default_GongTimeField.X = 647.000000
	__Position__000006_Pos_Default_GongTimeField.Y = 807.000000
	__Position__000006_Pos_Default_GongTimeField.Name = `Pos-Default-GongTimeField`

	// Position values setup
	__Position__000007_Pos_Default_ModelPkg.X = 49.000000
	__Position__000007_Pos_Default_ModelPkg.Y = 736.000000
	__Position__000007_Pos_Default_ModelPkg.Name = `Pos-Default-ModelPkg`

	// Position values setup
	__Position__000008_Pos_Default_PointerToGongStructField.X = 650.000000
	__Position__000008_Pos_Default_PointerToGongStructField.Y = 95.000000
	__Position__000008_Pos_Default_PointerToGongStructField.Name = `Pos-Default-PointerToGongStructField`

	// Position values setup
	__Position__000009_Pos_Default_SliceOfPointerToGongStructField.X = 651.000000
	__Position__000009_Pos_Default_SliceOfPointerToGongStructField.Y = 199.000000
	__Position__000009_Pos_Default_SliceOfPointerToGongStructField.Name = `Pos-Default-SliceOfPointerToGongStructField`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_GongBasicField_and_Default_GongEnum.X = 770.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_GongBasicField_and_Default_GongEnum.Y = 484.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_GongBasicField_and_Default_GongEnum.Name = `Verticle in class diagram Default in middle between Default-GongBasicField and Default-GongEnum`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_GongEnum_and_Default_GongEnumValue.X = 423.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_GongEnum_and_Default_GongEnumValue.Y = 43.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_GongEnum_and_Default_GongEnumValue.Name = `Verticle in class diagram Default in middle between Default-GongEnum and Default-GongEnumValue`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_GongNote_and_Default_GongLink.X = 743.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_GongNote_and_Default_GongLink.Y = 500.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_GongNote_and_Default_GongLink.Name = `Verticle in class diagram Default in middle between Default-GongNote and Default-GongLink`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_GongBasicField.X = 466.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_GongBasicField.Y = 259.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_GongBasicField.Name = `Verticle in class diagram Default in middle between Default-GongStruct and Default-GongBasicField`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_PointerToGongStructField.X = 415.000000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_PointerToGongStructField.Y = 103.000000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_PointerToGongStructField.Name = `Verticle in class diagram Default in middle between Default-GongStruct and Default-PointerToGongStructField`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_SliceOfPointerToGongStructField.X = 699.500000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_SliceOfPointerToGongStructField.Y = 187.000000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_SliceOfPointerToGongStructField.Name = `Verticle in class diagram Default in middle between Default-GongStruct and Default-SliceOfPointerToGongStructField`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_GongStruct)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_GongBasicField)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_GongLink)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000008_Default_PointerToGongStructField)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_GongNote)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_GongTimeField)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000007_Default_ModelPkg)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000009_Default_SliceOfPointerToGongStructField)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_GongEnum)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_GongEnumValue)
	__GongStructShape__000000_Default_GongBasicField.Position = __Position__000000_Pos_Default_GongBasicField
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000015_Name)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000000_BasicKindName)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000006_DeclaredType)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000005_CompositeStructName)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000009_Index)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000012_IsDocLink)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000013_IsTextArea)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000011_IsBespokeWidth)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000002_BespokeWidth)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000010_IsBespokeHeight)
	__GongStructShape__000000_Default_GongBasicField.Fields = append(__GongStructShape__000000_Default_GongBasicField.Fields, __Field__000001_BespokeHeight)
	__GongStructShape__000000_Default_GongBasicField.Links = append(__GongStructShape__000000_Default_GongBasicField.Links, __Link__000001_GongEnum)
	__GongStructShape__000001_Default_GongEnum.Position = __Position__000001_Pos_Default_GongEnum
	__GongStructShape__000001_Default_GongEnum.Fields = append(__GongStructShape__000001_Default_GongEnum.Fields, __Field__000017_Name)
	__GongStructShape__000001_Default_GongEnum.Fields = append(__GongStructShape__000001_Default_GongEnum.Fields, __Field__000021_Type)
	__GongStructShape__000001_Default_GongEnum.Links = append(__GongStructShape__000001_Default_GongEnum.Links, __Link__000002_GongEnumValues)
	__GongStructShape__000002_Default_GongEnumValue.Position = __Position__000002_Pos_Default_GongEnumValue
	__GongStructShape__000002_Default_GongEnumValue.Fields = append(__GongStructShape__000002_Default_GongEnumValue.Fields, __Field__000019_Name)
	__GongStructShape__000002_Default_GongEnumValue.Fields = append(__GongStructShape__000002_Default_GongEnumValue.Fields, __Field__000022_Value)
	__GongStructShape__000003_Default_GongLink.Position = __Position__000003_Pos_Default_GongLink
	__GongStructShape__000003_Default_GongLink.Fields = append(__GongStructShape__000003_Default_GongLink.Fields, __Field__000018_Name)
	__GongStructShape__000003_Default_GongLink.Fields = append(__GongStructShape__000003_Default_GongLink.Fields, __Field__000020_Recv)
	__GongStructShape__000003_Default_GongLink.Fields = append(__GongStructShape__000003_Default_GongLink.Fields, __Field__000008_ImportPath)
	__GongStructShape__000004_Default_GongNote.Position = __Position__000004_Pos_Default_GongNote
	__GongStructShape__000004_Default_GongNote.Fields = append(__GongStructShape__000004_Default_GongNote.Fields, __Field__000014_Name)
	__GongStructShape__000004_Default_GongNote.Fields = append(__GongStructShape__000004_Default_GongNote.Fields, __Field__000003_Body)
	__GongStructShape__000004_Default_GongNote.Fields = append(__GongStructShape__000004_Default_GongNote.Fields, __Field__000004_BodyHTML)
	__GongStructShape__000004_Default_GongNote.Links = append(__GongStructShape__000004_Default_GongNote.Links, __Link__000003_Links)
	__GongStructShape__000005_Default_GongStruct.Position = __Position__000005_Pos_Default_GongStruct
	__GongStructShape__000005_Default_GongStruct.Fields = append(__GongStructShape__000005_Default_GongStruct.Fields, __Field__000016_Name)
	__GongStructShape__000005_Default_GongStruct.Fields = append(__GongStructShape__000005_Default_GongStruct.Fields, __Field__000007_HasOnAfterUpdateSignature)
	__GongStructShape__000005_Default_GongStruct.Links = append(__GongStructShape__000005_Default_GongStruct.Links, __Link__000004_PointerToGongStructFields)
	__GongStructShape__000005_Default_GongStruct.Links = append(__GongStructShape__000005_Default_GongStruct.Links, __Link__000000_GongBasicFields)
	__GongStructShape__000005_Default_GongStruct.Links = append(__GongStructShape__000005_Default_GongStruct.Links, __Link__000005_SliceOfPointerToGongStructFields)
	__GongStructShape__000006_Default_GongTimeField.Position = __Position__000006_Pos_Default_GongTimeField
	__GongStructShape__000007_Default_ModelPkg.Position = __Position__000007_Pos_Default_ModelPkg
	__GongStructShape__000008_Default_PointerToGongStructField.Position = __Position__000008_Pos_Default_PointerToGongStructField
	__GongStructShape__000009_Default_SliceOfPointerToGongStructField.Position = __Position__000009_Pos_Default_SliceOfPointerToGongStructField
	__Link__000000_GongBasicFields.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_GongBasicField
	__Link__000001_GongEnum.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_GongBasicField_and_Default_GongEnum
	__Link__000002_GongEnumValues.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_GongEnum_and_Default_GongEnumValue
	__Link__000003_Links.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_GongNote_and_Default_GongLink
	__Link__000004_PointerToGongStructFields.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_PointerToGongStructField
	__Link__000005_SliceOfPointerToGongStructFields.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_GongStruct_and_Default_SliceOfPointerToGongStructField
}
