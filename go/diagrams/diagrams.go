package diagrams

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/go/models"
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
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `ImportPath`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `Recv`}).Stage(stage)
	__AttributeShape__00000003_ := (&models.AttributeShape{Name: `Type`}).Stage(stage)
	__AttributeShape__00000004_ := (&models.AttributeShape{Name: `Body`}).Stage(stage)
	__AttributeShape__00000005_ := (&models.AttributeShape{Name: `BodyHTML`}).Stage(stage)
	__AttributeShape__00000006_ := (&models.AttributeShape{Name: `HasOnAfterUpdateSignature`}).Stage(stage)
	__AttributeShape__00000009_ := (&models.AttributeShape{Name: `DeclaredType`}).Stage(stage)
	__AttributeShape__00000010_ := (&models.AttributeShape{Name: `CompositeStructName`}).Stage(stage)
	__AttributeShape__00000011_ := (&models.AttributeShape{Name: `Index`}).Stage(stage)
	__AttributeShape__00000012_ := (&models.AttributeShape{Name: `IsTextArea`}).Stage(stage)
	__AttributeShape__00000013_ := (&models.AttributeShape{Name: `IsBespokeWidth`}).Stage(stage)
	__AttributeShape__00000014_ := (&models.AttributeShape{Name: `BespokeWidth`}).Stage(stage)
	__AttributeShape__00000015_ := (&models.AttributeShape{Name: `BespokeHeight`}).Stage(stage)
	__AttributeShape__00000016_ := (&models.AttributeShape{Name: `Index`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-06-10T02:31:57Z`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-GongStruct`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-GongBasicField`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-GongEnum`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Default-GongEnumValue`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Default-PointerToGongStructField`}).Stage(stage)
	__GongStructShape__00000005_ := (&models.GongStructShape{Name: `Default-SliceOfPointerToGongStructField`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `Default-GongNote`}).Stage(stage)
	__GongStructShape__00000007_ := (&models.GongStructShape{Name: `Default-GongLink`}).Stage(stage)
	__GongStructShape__00000008_ := (&models.GongStructShape{Name: `Default-GongTimeField`}).Stage(stage)
	__GongStructShape__00000009_ := (&models.GongStructShape{Name: `Default-MetaReference`}).Stage(stage)
	__GongStructShape__00000010_ := (&models.GongStructShape{Name: `Default-ModelPkg`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `GongBasicFields`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `SliceOfPointerToGongStructFields`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `PointerToGongStructFields`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `Links`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `GongEnumValues`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `GongTimeFields`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `GongEnum`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `GongStruct`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `GongStruct`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `Name`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.GongStruct{}.Name
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `GongStruct`
	__AttributeShape__00000000_.Fieldtypename = `string`

	__AttributeShape__00000001_.Name = `ImportPath`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.GongLink{}.ImportPath
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `GongLink`
	__AttributeShape__00000001_.Fieldtypename = `string`

	__AttributeShape__00000002_.Name = `Recv`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.GongLink{}.Recv
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `GongLink`
	__AttributeShape__00000002_.Fieldtypename = `string`

	__AttributeShape__00000003_.Name = `Type`
	__AttributeShape__00000003_.IdentifierMeta = ref_models.GongEnum{}.Type
	__AttributeShape__00000003_.FieldTypeAsString = ``
	__AttributeShape__00000003_.Structname = `GongEnum`
	__AttributeShape__00000003_.Fieldtypename = `GongEnumType`

	__AttributeShape__00000004_.Name = `Body`
	__AttributeShape__00000004_.IdentifierMeta = ref_models.GongNote{}.Body
	__AttributeShape__00000004_.FieldTypeAsString = ``
	__AttributeShape__00000004_.Structname = `GongNote`
	__AttributeShape__00000004_.Fieldtypename = `string`

	__AttributeShape__00000005_.Name = `BodyHTML`
	__AttributeShape__00000005_.IdentifierMeta = ref_models.GongNote{}.BodyHTML
	__AttributeShape__00000005_.FieldTypeAsString = ``
	__AttributeShape__00000005_.Structname = `GongNote`
	__AttributeShape__00000005_.Fieldtypename = `string`

	__AttributeShape__00000006_.Name = `HasOnAfterUpdateSignature`
	__AttributeShape__00000006_.IdentifierMeta = ref_models.GongStruct{}.HasOnAfterUpdateSignature
	__AttributeShape__00000006_.FieldTypeAsString = ``
	__AttributeShape__00000006_.Structname = `GongStruct`
	__AttributeShape__00000006_.Fieldtypename = `bool`

	__AttributeShape__00000009_.Name = `DeclaredType`
	__AttributeShape__00000009_.IdentifierMeta = ref_models.GongBasicField{}.DeclaredType
	__AttributeShape__00000009_.FieldTypeAsString = ``
	__AttributeShape__00000009_.Structname = `GongBasicField`
	__AttributeShape__00000009_.Fieldtypename = `string`

	__AttributeShape__00000010_.Name = `CompositeStructName`
	__AttributeShape__00000010_.IdentifierMeta = ref_models.GongBasicField{}.CompositeStructName
	__AttributeShape__00000010_.FieldTypeAsString = ``
	__AttributeShape__00000010_.Structname = `GongBasicField`
	__AttributeShape__00000010_.Fieldtypename = `string`

	__AttributeShape__00000011_.Name = `Index`
	__AttributeShape__00000011_.IdentifierMeta = ref_models.GongBasicField{}.Index
	__AttributeShape__00000011_.FieldTypeAsString = ``
	__AttributeShape__00000011_.Structname = `GongBasicField`
	__AttributeShape__00000011_.Fieldtypename = `int`

	__AttributeShape__00000012_.Name = `IsTextArea`
	__AttributeShape__00000012_.IdentifierMeta = ref_models.GongBasicField{}.IsTextArea
	__AttributeShape__00000012_.FieldTypeAsString = ``
	__AttributeShape__00000012_.Structname = `GongBasicField`
	__AttributeShape__00000012_.Fieldtypename = `bool`

	__AttributeShape__00000013_.Name = `IsBespokeWidth`
	__AttributeShape__00000013_.IdentifierMeta = ref_models.GongBasicField{}.IsBespokeWidth
	__AttributeShape__00000013_.FieldTypeAsString = ``
	__AttributeShape__00000013_.Structname = `GongBasicField`
	__AttributeShape__00000013_.Fieldtypename = `bool`

	__AttributeShape__00000014_.Name = `BespokeWidth`
	__AttributeShape__00000014_.IdentifierMeta = ref_models.GongBasicField{}.BespokeWidth
	__AttributeShape__00000014_.FieldTypeAsString = ``
	__AttributeShape__00000014_.Structname = `GongBasicField`
	__AttributeShape__00000014_.Fieldtypename = `int`

	__AttributeShape__00000015_.Name = `BespokeHeight`
	__AttributeShape__00000015_.IdentifierMeta = ref_models.GongBasicField{}.BespokeHeight
	__AttributeShape__00000015_.FieldTypeAsString = ``
	__AttributeShape__00000015_.Structname = `GongBasicField`
	__AttributeShape__00000015_.Fieldtypename = `int`

	__AttributeShape__00000016_.Name = `Index`
	__AttributeShape__00000016_.IdentifierMeta = ref_models.PointerToGongStructField{}.Index
	__AttributeShape__00000016_.FieldTypeAsString = ``
	__AttributeShape__00000016_.Structname = `PointerToGongStructField`
	__AttributeShape__00000016_.Fieldtypename = `int`

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[true,true,false,true,true,false,false,true,false,true,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2025-06-10T02:31:57Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-GongStruct`
	__GongStructShape__00000000_.X = 59.000000
	__GongStructShape__00000000_.Y = 132.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.GongStruct{}
	__GongStructShape__00000000_.Width = 290.000000
	__GongStructShape__00000000_.Height = 103.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-GongBasicField`
	__GongStructShape__00000001_.X = 772.000000
	__GongStructShape__00000001_.Y = 140.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.GongBasicField{}
	__GongStructShape__00000001_.Width = 311.000000
	__GongStructShape__00000001_.Height = 203.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-GongEnum`
	__GongStructShape__00000002_.X = 501.000000
	__GongStructShape__00000002_.Y = 774.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.GongEnum{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 83.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Default-GongEnumValue`
	__GongStructShape__00000003_.X = 102.000000
	__GongStructShape__00000003_.Y = 773.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.GongEnumValue{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 63.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Default-PointerToGongStructField`
	__GongStructShape__00000004_.X = 776.000000
	__GongStructShape__00000004_.Y = 507.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.PointerToGongStructField{}
	__GongStructShape__00000004_.Width = 306.000000
	__GongStructShape__00000004_.Height = 83.000000
	__GongStructShape__00000004_.IsSelected = false

	__GongStructShape__00000005_.Name = `Default-SliceOfPointerToGongStructField`
	__GongStructShape__00000005_.X = 772.000000
	__GongStructShape__00000005_.Y = 367.000000
	__GongStructShape__00000005_.IdentifierMeta = ref_models.SliceOfPointerToGongStructField{}
	__GongStructShape__00000005_.Width = 311.000000
	__GongStructShape__00000005_.Height = 63.000000
	__GongStructShape__00000005_.IsSelected = false

	__GongStructShape__00000006_.Name = `Default-GongNote`
	__GongStructShape__00000006_.X = 50.000000
	__GongStructShape__00000006_.Y = 929.000000
	__GongStructShape__00000006_.IdentifierMeta = ref_models.GongNote{}
	__GongStructShape__00000006_.Width = 240.000000
	__GongStructShape__00000006_.Height = 103.000000
	__GongStructShape__00000006_.IsSelected = false

	__GongStructShape__00000007_.Name = `Default-GongLink`
	__GongStructShape__00000007_.X = 510.000000
	__GongStructShape__00000007_.Y = 910.000000
	__GongStructShape__00000007_.IdentifierMeta = ref_models.GongLink{}
	__GongStructShape__00000007_.Width = 240.000000
	__GongStructShape__00000007_.Height = 103.000000
	__GongStructShape__00000007_.IsSelected = false

	__GongStructShape__00000008_.Name = `Default-GongTimeField`
	__GongStructShape__00000008_.X = 776.000000
	__GongStructShape__00000008_.Y = 659.000000
	__GongStructShape__00000008_.IdentifierMeta = ref_models.GongTimeField{}
	__GongStructShape__00000008_.Width = 240.000000
	__GongStructShape__00000008_.Height = 63.000000
	__GongStructShape__00000008_.IsSelected = false

	__GongStructShape__00000009_.Name = `Default-MetaReference`
	__GongStructShape__00000009_.X = 570.000000
	__GongStructShape__00000009_.Y = 34.000000
	__GongStructShape__00000009_.IdentifierMeta = ref_models.MetaReference{}
	__GongStructShape__00000009_.Width = 240.000000
	__GongStructShape__00000009_.Height = 63.000000
	__GongStructShape__00000009_.IsSelected = false

	__GongStructShape__00000010_.Name = `Default-ModelPkg`
	__GongStructShape__00000010_.X = 108.000000
	__GongStructShape__00000010_.Y = 21.000000
	__GongStructShape__00000010_.IdentifierMeta = ref_models.ModelPkg{}
	__GongStructShape__00000010_.Width = 240.000000
	__GongStructShape__00000010_.Height = 63.000000
	__GongStructShape__00000010_.IsSelected = false

	__LinkShape__00000000_.Name = `GongBasicFields`
	__LinkShape__00000000_.IdentifierMeta = ref_models.GongStruct{}.GongBasicFields
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.GongBasicField{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.MANY
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 435.000000
	__LinkShape__00000000_.Y = 196.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.500000
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.500000
	__LinkShape__00000000_.CornerOffsetRatio = 1.380000

	__LinkShape__00000001_.Name = `SliceOfPointerToGongStructFields`
	__LinkShape__00000001_.IdentifierMeta = ref_models.GongStruct{}.SliceOfPointerToGongStructFields
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.SliceOfPointerToGongStructField{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 692.500000
	__LinkShape__00000001_.Y = 400.500000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.500000
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.500000
	__LinkShape__00000001_.CornerOffsetRatio = 1.380000

	__LinkShape__00000002_.Name = `PointerToGongStructFields`
	__LinkShape__00000002_.IdentifierMeta = ref_models.GongStruct{}.PointerToGongStructFields
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.PointerToGongStructField{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.MANY
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 694.500000
	__LinkShape__00000002_.Y = 467.500000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.500000
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.500000
	__LinkShape__00000002_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `Links`
	__LinkShape__00000003_.IdentifierMeta = ref_models.GongNote{}.Links
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.GongLink{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.MANY
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 423.000000
	__LinkShape__00000003_.Y = 774.000000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.StartRatio = 0.339890
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.EndRatio = 0.500000
	__LinkShape__00000003_.CornerOffsetRatio = 1.380000

	__LinkShape__00000004_.Name = `GongEnumValues`
	__LinkShape__00000004_.IdentifierMeta = ref_models.GongEnum{}.GongEnumValues
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.GongEnumValue{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.MANY
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 1066.500000
	__LinkShape__00000004_.Y = 154.500000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.StartRatio = 0.500000
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.EndRatio = 0.555693
	__LinkShape__00000004_.CornerOffsetRatio = -0.204036

	__LinkShape__00000005_.Name = `GongTimeFields`
	__LinkShape__00000005_.IdentifierMeta = ref_models.GongStruct{}.GongTimeFields
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.GongTimeField{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.MANY
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 721.500000
	__LinkShape__00000005_.Y = 268.000000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.StartRatio = 0.500000
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.EndRatio = 0.500000
	__LinkShape__00000005_.CornerOffsetRatio = 1.380000

	__LinkShape__00000006_.Name = `GongEnum`
	__LinkShape__00000006_.IdentifierMeta = ref_models.GongBasicField{}.GongEnum
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.GongEnum{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 879.500000
	__LinkShape__00000006_.Y = 314.500000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.StartRatio = 0.396963
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.EndRatio = 0.469984
	__LinkShape__00000006_.CornerOffsetRatio = 1.167303

	__LinkShape__00000007_.Name = `GongStruct`
	__LinkShape__00000007_.IdentifierMeta = ref_models.PointerToGongStructField{}.GongStruct
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.GongStruct{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 859.000000
	__LinkShape__00000007_.Y = 235.500000
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000007_.StartRatio = 0.509906
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000007_.EndRatio = 0.365625
	__LinkShape__00000007_.CornerOffsetRatio = 1.428709

	__LinkShape__00000008_.Name = `GongStruct`
	__LinkShape__00000008_.IdentifierMeta = ref_models.SliceOfPointerToGongStructField{}.GongStruct
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.GongStruct{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 882.000000
	__LinkShape__00000008_.Y = 281.000000
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000008_.StartRatio = 0.537078
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000008_.EndRatio = 0.827694
	__LinkShape__00000008_.CornerOffsetRatio = 1.666804

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000005_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000006_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000007_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000008_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000009_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000010_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000006_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000009_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000010_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000011_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000012_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000013_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000014_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000015_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000003_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000004_.AttributeShapes = append(__GongStructShape__00000004_.AttributeShapes, __AttributeShape__00000016_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000008_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000004_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000005_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000001_)
}
