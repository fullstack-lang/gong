package diagrams

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/form/go/models"
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

	__AttributeShape__00000000_ := (&models.AttributeShape{Name: `Value`}).Stage(stage)
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `InputTypeEnum`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `HasBespokeWidth`}).Stage(stage)
	__AttributeShape__00000003_ := (&models.AttributeShape{Name: `Placeholder`}).Stage(stage)
	__AttributeShape__00000004_ := (&models.AttributeShape{Name: `Label`}).Stage(stage)
	__AttributeShape__00000005_ := (&models.AttributeShape{Name: `CanBeEmpty`}).Stage(stage)
	__AttributeShape__00000006_ := (&models.AttributeShape{Name: `PreserveInitialOrder`}).Stage(stage)
	__AttributeShape__00000007_ := (&models.AttributeShape{Name: `HasSuppressButton`}).Stage(stage)
	__AttributeShape__00000008_ := (&models.AttributeShape{Name: `HasSuppressButtonBeenPressed`}).Stage(stage)
	__AttributeShape__00000010_ := (&models.AttributeShape{Name: `AssociationStorage`}).Stage(stage)
	__AttributeShape__00000011_ := (&models.AttributeShape{Name: `HasChanged`}).Stage(stage)
	__AttributeShape__00000012_ := (&models.AttributeShape{Name: `IsForSavePurpose`}).Stage(stage)
	__AttributeShape__00000013_ := (&models.AttributeShape{Name: `HasToolTip`}).Stage(stage)
	__AttributeShape__00000014_ := (&models.AttributeShape{Name: `ToolTipText`}).Stage(stage)
	__AttributeShape__00000015_ := (&models.AttributeShape{Name: `MatTooltipShowDelay`}).Stage(stage)
	__AttributeShape__00000016_ := (&models.AttributeShape{Name: `Label`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2026-06-12T01:50:45Z`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-FormDiv`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-CheckBox`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-FormEditAssocButton`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Default-FormField`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Default-FormFieldDate`}).Stage(stage)
	__GongStructShape__00000005_ := (&models.GongStructShape{Name: `Default-FormFieldDateTime`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `Default-FormFieldFloat64`}).Stage(stage)
	__GongStructShape__00000007_ := (&models.GongStructShape{Name: `Default-FormFieldInt`}).Stage(stage)
	__GongStructShape__00000008_ := (&models.GongStructShape{Name: `Default-FormFieldSelect`}).Stage(stage)
	__GongStructShape__00000009_ := (&models.GongStructShape{Name: `Default-FormFieldString`}).Stage(stage)
	__GongStructShape__00000010_ := (&models.GongStructShape{Name: `Default-FormFieldTime`}).Stage(stage)
	__GongStructShape__00000011_ := (&models.GongStructShape{Name: `Default-FormGroup`}).Stage(stage)
	__GongStructShape__00000012_ := (&models.GongStructShape{Name: `Default-FormSortAssocButton`}).Stage(stage)
	__GongStructShape__00000013_ := (&models.GongStructShape{Name: `Default-Option`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `FormFields`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `CheckBoxs`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `FormEditAssocButton`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `FormSortAssocButton`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `FormFieldString`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `FormFieldFloat64`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `FormFieldInt`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `FormFieldDate`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `FormFieldTime`}).Stage(stage)
	__LinkShape__00000009_ := (&models.LinkShape{Name: `FormFieldDateTime`}).Stage(stage)
	__LinkShape__00000010_ := (&models.LinkShape{Name: `FormFieldSelect`}).Stage(stage)
	__LinkShape__00000011_ := (&models.LinkShape{Name: `Options`}).Stage(stage)
	__LinkShape__00000012_ := (&models.LinkShape{Name: `FormDivs`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `Value`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.CheckBox{}.Value
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `CheckBox`
	__AttributeShape__00000000_.Fieldtypename = `bool`

	__AttributeShape__00000001_.Name = `InputTypeEnum`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.FormField{}.InputTypeEnum
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `FormField`
	__AttributeShape__00000001_.Fieldtypename = `InputTypeEnum`

	__AttributeShape__00000002_.Name = `HasBespokeWidth`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.FormField{}.HasBespokeWidth
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `FormField`
	__AttributeShape__00000002_.Fieldtypename = `bool`

	__AttributeShape__00000003_.Name = `Placeholder`
	__AttributeShape__00000003_.IdentifierMeta = ref_models.FormField{}.Placeholder
	__AttributeShape__00000003_.FieldTypeAsString = ``
	__AttributeShape__00000003_.Structname = `FormField`
	__AttributeShape__00000003_.Fieldtypename = `string`

	__AttributeShape__00000004_.Name = `Label`
	__AttributeShape__00000004_.IdentifierMeta = ref_models.FormField{}.Label
	__AttributeShape__00000004_.FieldTypeAsString = ``
	__AttributeShape__00000004_.Structname = `FormField`
	__AttributeShape__00000004_.Fieldtypename = `string`

	__AttributeShape__00000005_.Name = `CanBeEmpty`
	__AttributeShape__00000005_.IdentifierMeta = ref_models.FormFieldSelect{}.CanBeEmpty
	__AttributeShape__00000005_.FieldTypeAsString = ``
	__AttributeShape__00000005_.Structname = `FormFieldSelect`
	__AttributeShape__00000005_.Fieldtypename = `bool`

	__AttributeShape__00000006_.Name = `PreserveInitialOrder`
	__AttributeShape__00000006_.IdentifierMeta = ref_models.FormFieldSelect{}.PreserveInitialOrder
	__AttributeShape__00000006_.FieldTypeAsString = ``
	__AttributeShape__00000006_.Structname = `FormFieldSelect`
	__AttributeShape__00000006_.Fieldtypename = `bool`

	__AttributeShape__00000007_.Name = `HasSuppressButton`
	__AttributeShape__00000007_.IdentifierMeta = ref_models.FormGroup{}.HasSuppressButton
	__AttributeShape__00000007_.FieldTypeAsString = ``
	__AttributeShape__00000007_.Structname = `FormGroup`
	__AttributeShape__00000007_.Fieldtypename = `bool`

	__AttributeShape__00000008_.Name = `HasSuppressButtonBeenPressed`
	__AttributeShape__00000008_.IdentifierMeta = ref_models.FormGroup{}.HasSuppressButtonBeenPressed
	__AttributeShape__00000008_.FieldTypeAsString = ``
	__AttributeShape__00000008_.Structname = `FormGroup`
	__AttributeShape__00000008_.Fieldtypename = `bool`

	__AttributeShape__00000010_.Name = `AssociationStorage`
	__AttributeShape__00000010_.IdentifierMeta = ref_models.FormEditAssocButton{}.AssociationStorage
	__AttributeShape__00000010_.FieldTypeAsString = ``
	__AttributeShape__00000010_.Structname = `FormEditAssocButton`
	__AttributeShape__00000010_.Fieldtypename = `string`

	__AttributeShape__00000011_.Name = `HasChanged`
	__AttributeShape__00000011_.IdentifierMeta = ref_models.FormEditAssocButton{}.HasChanged
	__AttributeShape__00000011_.FieldTypeAsString = ``
	__AttributeShape__00000011_.Structname = `FormEditAssocButton`
	__AttributeShape__00000011_.Fieldtypename = `bool`

	__AttributeShape__00000012_.Name = `IsForSavePurpose`
	__AttributeShape__00000012_.IdentifierMeta = ref_models.FormEditAssocButton{}.IsForSavePurpose
	__AttributeShape__00000012_.FieldTypeAsString = ``
	__AttributeShape__00000012_.Structname = `FormEditAssocButton`
	__AttributeShape__00000012_.Fieldtypename = `bool`

	__AttributeShape__00000013_.Name = `HasToolTip`
	__AttributeShape__00000013_.IdentifierMeta = ref_models.FormEditAssocButton{}.HasToolTip
	__AttributeShape__00000013_.FieldTypeAsString = ``
	__AttributeShape__00000013_.Structname = `FormEditAssocButton`
	__AttributeShape__00000013_.Fieldtypename = `bool`

	__AttributeShape__00000014_.Name = `ToolTipText`
	__AttributeShape__00000014_.IdentifierMeta = ref_models.FormEditAssocButton{}.ToolTipText
	__AttributeShape__00000014_.FieldTypeAsString = ``
	__AttributeShape__00000014_.Structname = `FormEditAssocButton`
	__AttributeShape__00000014_.Fieldtypename = `string`

	__AttributeShape__00000015_.Name = `MatTooltipShowDelay`
	__AttributeShape__00000015_.IdentifierMeta = ref_models.FormEditAssocButton{}.MatTooltipShowDelay
	__AttributeShape__00000015_.FieldTypeAsString = ``
	__AttributeShape__00000015_.Structname = `FormEditAssocButton`
	__AttributeShape__00000015_.Fieldtypename = `string`

	__AttributeShape__00000016_.Name = `Label`
	__AttributeShape__00000016_.IdentifierMeta = ref_models.FormEditAssocButton{}.Label
	__AttributeShape__00000016_.FieldTypeAsString = ``
	__AttributeShape__00000016_.Structname = `FormEditAssocButton`
	__AttributeShape__00000016_.Fieldtypename = `string`

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,false,false,true,false,false,false,false,false,false,false,false,false,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2026-06-12T01:50:45Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_.Name = `Default-FormDiv`
	__GongStructShape__00000000_.X = 101.000000
	__GongStructShape__00000000_.Y = 18.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.FormDiv{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 63.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-CheckBox`
	__GongStructShape__00000001_.X = 553.000000
	__GongStructShape__00000001_.Y = 198.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.CheckBox{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 83.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-FormEditAssocButton`
	__GongStructShape__00000002_.X = 551.000000
	__GongStructShape__00000002_.Y = 385.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.FormEditAssocButton{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 203.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Default-FormField`
	__GongStructShape__00000003_.X = 559.000000
	__GongStructShape__00000003_.Y = 20.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.FormField{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 143.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Default-FormFieldDate`
	__GongStructShape__00000004_.X = 966.000000
	__GongStructShape__00000004_.Y = 531.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.FormFieldDate{}
	__GongStructShape__00000004_.Width = 240.000000
	__GongStructShape__00000004_.Height = 63.000000
	__GongStructShape__00000004_.IsSelected = false

	__GongStructShape__00000005_.Name = `Default-FormFieldDateTime`
	__GongStructShape__00000005_.X = 975.000000
	__GongStructShape__00000005_.Y = 357.000000
	__GongStructShape__00000005_.IdentifierMeta = ref_models.FormFieldDateTime{}
	__GongStructShape__00000005_.Width = 240.000000
	__GongStructShape__00000005_.Height = 63.000000
	__GongStructShape__00000005_.IsSelected = false

	__GongStructShape__00000006_.Name = `Default-FormFieldFloat64`
	__GongStructShape__00000006_.X = 976.000000
	__GongStructShape__00000006_.Y = 604.000000
	__GongStructShape__00000006_.IdentifierMeta = ref_models.FormFieldFloat64{}
	__GongStructShape__00000006_.Width = 240.000000
	__GongStructShape__00000006_.Height = 63.000000
	__GongStructShape__00000006_.IsSelected = false

	__GongStructShape__00000007_.Name = `Default-FormFieldInt`
	__GongStructShape__00000007_.X = 970.000000
	__GongStructShape__00000007_.Y = 437.000000
	__GongStructShape__00000007_.IdentifierMeta = ref_models.FormFieldInt{}
	__GongStructShape__00000007_.Width = 240.000000
	__GongStructShape__00000007_.Height = 63.000000
	__GongStructShape__00000007_.IsSelected = false

	__GongStructShape__00000008_.Name = `Default-FormFieldSelect`
	__GongStructShape__00000008_.X = 960.000000
	__GongStructShape__00000008_.Y = 220.000000
	__GongStructShape__00000008_.IdentifierMeta = ref_models.FormFieldSelect{}
	__GongStructShape__00000008_.Width = 240.000000
	__GongStructShape__00000008_.Height = 103.000000
	__GongStructShape__00000008_.IsSelected = false

	__GongStructShape__00000009_.Name = `Default-FormFieldString`
	__GongStructShape__00000009_.X = 956.000000
	__GongStructShape__00000009_.Y = 140.000000
	__GongStructShape__00000009_.IdentifierMeta = ref_models.FormFieldString{}
	__GongStructShape__00000009_.Width = 240.000000
	__GongStructShape__00000009_.Height = 63.000000
	__GongStructShape__00000009_.IsSelected = false

	__GongStructShape__00000010_.Name = `Default-FormFieldTime`
	__GongStructShape__00000010_.X = 955.000000
	__GongStructShape__00000010_.Y = 44.000000
	__GongStructShape__00000010_.IdentifierMeta = ref_models.FormFieldTime{}
	__GongStructShape__00000010_.Width = 240.000000
	__GongStructShape__00000010_.Height = 63.000000
	__GongStructShape__00000010_.IsSelected = false

	__GongStructShape__00000011_.Name = `Default-FormGroup`
	__GongStructShape__00000011_.X = 74.000000
	__GongStructShape__00000011_.Y = 237.000000
	__GongStructShape__00000011_.IdentifierMeta = ref_models.FormGroup{}
	__GongStructShape__00000011_.Width = 296.000000
	__GongStructShape__00000011_.Height = 103.000000
	__GongStructShape__00000011_.IsSelected = false

	__GongStructShape__00000012_.Name = `Default-FormSortAssocButton`
	__GongStructShape__00000012_.X = 550.000000
	__GongStructShape__00000012_.Y = 307.000000
	__GongStructShape__00000012_.IdentifierMeta = ref_models.FormSortAssocButton{}
	__GongStructShape__00000012_.Width = 240.000000
	__GongStructShape__00000012_.Height = 63.000000
	__GongStructShape__00000012_.IsSelected = false

	__GongStructShape__00000013_.Name = `Default-Option`
	__GongStructShape__00000013_.X = 1321.000000
	__GongStructShape__00000013_.Y = 235.000000
	__GongStructShape__00000013_.IdentifierMeta = ref_models.Option{}
	__GongStructShape__00000013_.Width = 240.000000
	__GongStructShape__00000013_.Height = 63.000000
	__GongStructShape__00000013_.IsSelected = false

	__LinkShape__00000000_.Name = `FormFields`
	__LinkShape__00000000_.IdentifierMeta = ref_models.FormDiv{}.FormFields
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.FormField{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.MANY
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 455.000000
	__LinkShape__00000000_.Y = 78.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.500000
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.500000
	__LinkShape__00000000_.CornerOffsetRatio = 1.380000

	__LinkShape__00000001_.Name = `CheckBoxs`
	__LinkShape__00000001_.IdentifierMeta = ref_models.FormDiv{}.CheckBoxs
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.CheckBox{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 449.000000
	__LinkShape__00000001_.Y = 77.000000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.500000
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.500000
	__LinkShape__00000001_.CornerOffsetRatio = 1.380000

	__LinkShape__00000002_.Name = `FormEditAssocButton`
	__LinkShape__00000002_.IdentifierMeta = ref_models.FormDiv{}.FormEditAssocButton
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.FormEditAssocButton{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 426.000000
	__LinkShape__00000002_.Y = 69.500000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.500000
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.500000
	__LinkShape__00000002_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `FormSortAssocButton`
	__LinkShape__00000003_.IdentifierMeta = ref_models.FormDiv{}.FormSortAssocButton
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.FormSortAssocButton{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 601.500000
	__LinkShape__00000003_.Y = 161.500000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.StartRatio = 0.500000
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.EndRatio = 0.500000
	__LinkShape__00000003_.CornerOffsetRatio = 1.380000

	__LinkShape__00000004_.Name = `FormFieldString`
	__LinkShape__00000004_.IdentifierMeta = ref_models.FormField{}.FormFieldString
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.FormFieldString{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 1121.500000
	__LinkShape__00000004_.Y = 164.000000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.StartRatio = 0.500000
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.EndRatio = 0.500000
	__LinkShape__00000004_.CornerOffsetRatio = 1.380000

	__LinkShape__00000005_.Name = `FormFieldFloat64`
	__LinkShape__00000005_.IdentifierMeta = ref_models.FormField{}.FormFieldFloat64
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.FormFieldFloat64{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 1127.000000
	__LinkShape__00000005_.Y = 365.500000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.StartRatio = 0.500000
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.EndRatio = 0.500000
	__LinkShape__00000005_.CornerOffsetRatio = 1.380000

	__LinkShape__00000006_.Name = `FormFieldInt`
	__LinkShape__00000006_.IdentifierMeta = ref_models.FormField{}.FormFieldInt
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.FormFieldInt{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 1124.000000
	__LinkShape__00000006_.Y = 282.000000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.StartRatio = 0.500000
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.EndRatio = 0.500000
	__LinkShape__00000006_.CornerOffsetRatio = 1.380000

	__LinkShape__00000007_.Name = `FormFieldDate`
	__LinkShape__00000007_.IdentifierMeta = ref_models.FormField{}.FormFieldDate
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.FormFieldDate{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 1122.000000
	__LinkShape__00000007_.Y = 329.000000
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.StartRatio = 0.500000
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.EndRatio = 0.500000
	__LinkShape__00000007_.CornerOffsetRatio = 1.380000

	__LinkShape__00000008_.Name = `FormFieldTime`
	__LinkShape__00000008_.IdentifierMeta = ref_models.FormField{}.FormFieldTime
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.FormFieldTime{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 1112.500000
	__LinkShape__00000008_.Y = 120.500000
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.StartRatio = 0.500000
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.EndRatio = 0.500000
	__LinkShape__00000008_.CornerOffsetRatio = 1.380000

	__LinkShape__00000009_.Name = `FormFieldDateTime`
	__LinkShape__00000009_.IdentifierMeta = ref_models.FormField{}.FormFieldDateTime
	__LinkShape__00000009_.FieldTypeIdentifierMeta = ref_models.FormFieldDateTime{}
	__LinkShape__00000009_.FieldOffsetX = 0.000000
	__LinkShape__00000009_.FieldOffsetY = 0.000000
	__LinkShape__00000009_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000009_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.SourceMultiplicity = models.MANY
	__LinkShape__00000009_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.X = 1126.500000
	__LinkShape__00000009_.Y = 242.000000
	__LinkShape__00000009_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.StartRatio = 0.500000
	__LinkShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.EndRatio = 0.500000
	__LinkShape__00000009_.CornerOffsetRatio = 1.380000

	__LinkShape__00000010_.Name = `FormFieldSelect`
	__LinkShape__00000010_.IdentifierMeta = ref_models.FormField{}.FormFieldSelect
	__LinkShape__00000010_.FieldTypeIdentifierMeta = ref_models.FormFieldSelect{}
	__LinkShape__00000010_.FieldOffsetX = 0.000000
	__LinkShape__00000010_.FieldOffsetY = 0.000000
	__LinkShape__00000010_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000010_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.SourceMultiplicity = models.MANY
	__LinkShape__00000010_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.X = 1124.500000
	__LinkShape__00000010_.Y = 203.500000
	__LinkShape__00000010_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000010_.StartRatio = 0.500000
	__LinkShape__00000010_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000010_.EndRatio = 0.500000
	__LinkShape__00000010_.CornerOffsetRatio = 1.380000

	__LinkShape__00000011_.Name = `Options`
	__LinkShape__00000011_.IdentifierMeta = ref_models.FormFieldSelect{}.Options
	__LinkShape__00000011_.FieldTypeIdentifierMeta = ref_models.Option{}
	__LinkShape__00000011_.FieldOffsetX = 0.000000
	__LinkShape__00000011_.FieldOffsetY = 0.000000
	__LinkShape__00000011_.TargetMultiplicity = models.MANY
	__LinkShape__00000011_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.SourceMultiplicity = models.MANY
	__LinkShape__00000011_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.X = 1498.000000
	__LinkShape__00000011_.Y = 313.500000
	__LinkShape__00000011_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.StartRatio = 0.500000
	__LinkShape__00000011_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.EndRatio = 0.500000
	__LinkShape__00000011_.CornerOffsetRatio = 1.380000

	__LinkShape__00000012_.Name = `FormDivs`
	__LinkShape__00000012_.IdentifierMeta = ref_models.FormGroup{}.FormDivs
	__LinkShape__00000012_.FieldTypeIdentifierMeta = ref_models.FormDiv{}
	__LinkShape__00000012_.FieldOffsetX = 0.000000
	__LinkShape__00000012_.FieldOffsetY = 0.000000
	__LinkShape__00000012_.TargetMultiplicity = models.MANY
	__LinkShape__00000012_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.SourceMultiplicity = models.MANY
	__LinkShape__00000012_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.X = 447.500000
	__LinkShape__00000012_.Y = 159.500000
	__LinkShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000012_.StartRatio = 0.296875
	__LinkShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000012_.EndRatio = 0.226042
	__LinkShape__00000012_.CornerOffsetRatio = -0.853175

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
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000011_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000012_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000013_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000016_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000010_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000011_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000012_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000013_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000014_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000015_)
	__GongStructShape__00000003_.AttributeShapes = append(__GongStructShape__00000003_.AttributeShapes, __AttributeShape__00000001_)
	__GongStructShape__00000003_.AttributeShapes = append(__GongStructShape__00000003_.AttributeShapes, __AttributeShape__00000004_)
	__GongStructShape__00000003_.AttributeShapes = append(__GongStructShape__00000003_.AttributeShapes, __AttributeShape__00000003_)
	__GongStructShape__00000003_.AttributeShapes = append(__GongStructShape__00000003_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000008_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000009_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000010_)
	__GongStructShape__00000008_.AttributeShapes = append(__GongStructShape__00000008_.AttributeShapes, __AttributeShape__00000005_)
	__GongStructShape__00000008_.AttributeShapes = append(__GongStructShape__00000008_.AttributeShapes, __AttributeShape__00000006_)
	__GongStructShape__00000008_.LinkShapes = append(__GongStructShape__00000008_.LinkShapes, __LinkShape__00000011_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000007_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000008_)
	__GongStructShape__00000011_.LinkShapes = append(__GongStructShape__00000011_.LinkShapes, __LinkShape__00000012_)
}
