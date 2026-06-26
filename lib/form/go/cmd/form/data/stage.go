package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/form/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__CheckBox__00000000_ := (&models.CheckBox{Name: `Boolean`}).Stage(stage)

	__FormDiv__00000000_ := (&models.FormDiv{Name: `CheckBoxs`}).Stage(stage)
	__FormDiv__00000001_ := (&models.FormDiv{Name: `Date - Time`}).Stage(stage)
	__FormDiv__00000002_ := (&models.FormDiv{Name: `DateTime`}).Stage(stage)
	__FormDiv__00000003_ := (&models.FormDiv{Name: `Edit Assoc`}).Stage(stage)
	__FormDiv__00000004_ := (&models.FormDiv{Name: `Float64 - level of x`}).Stage(stage)
	__FormDiv__00000005_ := (&models.FormDiv{Name: `Form Div First Name  - Name`}).Stage(stage)
	__FormDiv__00000006_ := (&models.FormDiv{Name: `Int - Age`}).Stage(stage)
	__FormDiv__00000007_ := (&models.FormDiv{Name: `Select (can be set empty)`}).Stage(stage)
	__FormDiv__00000008_ := (&models.FormDiv{Name: `Select (cannot be set empty)`}).Stage(stage)
	__FormDiv__00000009_ := (&models.FormDiv{Name: `Text Area`}).Stage(stage)
	__FormDiv__00000010_ := (&models.FormDiv{Name: `Divider`}).Stage(stage)
	__FormDiv__00000011_ := (&models.FormDiv{Name: `Accordion Group Start`}).Stage(stage)
	__FormDiv__00000012_ := (&models.FormDiv{Name: `End of Accordion Group`}).Stage(stage)

	__FormEditAssocButton__00000000_ := (&models.FormEditAssocButton{Name: `Edit Assoc`}).Stage(stage)

	__FormField__00000000_ := (&models.FormField{Name: `Age`}).Stage(stage)
	__FormField__00000001_ := (&models.FormField{Name: `Date`}).Stage(stage)
	__FormField__00000002_ := (&models.FormField{Name: `DateTime`}).Stage(stage)
	__FormField__00000003_ := (&models.FormField{Name: `FirstName`}).Stage(stage)
	__FormField__00000004_ := (&models.FormField{Name: `Float64 - level of x`}).Stage(stage)
	__FormField__00000005_ := (&models.FormField{Name: `LastName`}).Stage(stage)
	__FormField__00000006_ := (&models.FormField{Name: `Select (can be set empty)`}).Stage(stage)
	__FormField__00000007_ := (&models.FormField{Name: `Select (cannot be set empty)`}).Stage(stage)
	__FormField__00000008_ := (&models.FormField{Name: `Text Area`}).Stage(stage)
	__FormField__00000009_ := (&models.FormField{Name: `Time`}).Stage(stage)
	__FormField__00000010_ := (&models.FormField{Name: `Time`}).Stage(stage)

	__FormFieldDate__00000000_ := (&models.FormFieldDate{Name: `Time`}).Stage(stage)

	__FormFieldDateTime__00000000_ := (&models.FormFieldDateTime{Name: `DateTime`}).Stage(stage)

	__FormFieldFloat64__00000000_ := (&models.FormFieldFloat64{Name: `Float64 - level of x`}).Stage(stage)

	__FormFieldInt__00000000_ := (&models.FormFieldInt{Name: `Age`}).Stage(stage)

	__FormFieldSelect__00000000_ := (&models.FormFieldSelect{Name: `Select (can be set empty)`}).Stage(stage)
	__FormFieldSelect__00000001_ := (&models.FormFieldSelect{Name: `Select (cannot be set empty)`}).Stage(stage)

	__FormFieldString__00000000_ := (&models.FormFieldString{Name: `FirstName`}).Stage(stage)
	__FormFieldString__00000001_ := (&models.FormFieldString{Name: `LastName`}).Stage(stage)
	__FormFieldString__00000002_ := (&models.FormFieldString{Name: `Text Area`}).Stage(stage)

	__FormFieldTime__00000000_ := (&models.FormFieldTime{Name: `Time`}).Stage(stage)
	__FormFieldTime__00000001_ := (&models.FormFieldTime{Name: `Time 2`}).Stage(stage)

	__FormGroup__00000000_ := (&models.FormGroup{Name: `Form 1`}).Stage(stage)

	__FormSortAssocButton__00000000_ := (&models.FormSortAssocButton{Name: `Sort Button`}).Stage(stage)

	__Option__00000000_ := (&models.Option{Name: `A - Select (cannot be set empty)`}).Stage(stage)
	__Option__00000001_ := (&models.Option{Name: `B- Select (cannot be set empty)`}).Stage(stage)
	__Option__00000002_ := (&models.Option{Name: `Option 1 - Select (can be set empty)`}).Stage(stage)
	__Option__00000003_ := (&models.Option{Name: `Option 2 - Select (can be set empty)`}).Stage(stage)
	__Option__00000004_ := (&models.Option{Name: `Option 3 - Select (can be set empty)`}).Stage(stage)
	__Option__00000005_ := (&models.Option{Name: `Option 4 - Select (can be set empty)`}).Stage(stage)

	// insertion point for initialization of values

	__CheckBox__00000000_.Name = `Boolean`
	__CheckBox__00000000_.Value = true

	__FormDiv__00000000_.Name = `CheckBoxs`
	__FormDiv__00000000_.IsADivider = false
	__FormDiv__00000000_.IsAStartAccordionGroup = false
	__FormDiv__00000000_.AccordionGroupName = ``
	__FormDiv__00000000_.IsAEndAccordionGroup = false

	__FormDiv__00000001_.Name = `Date - Time`
	__FormDiv__00000001_.IsADivider = false
	__FormDiv__00000001_.IsAStartAccordionGroup = false
	__FormDiv__00000001_.AccordionGroupName = ``
	__FormDiv__00000001_.IsAEndAccordionGroup = false

	__FormDiv__00000002_.Name = `DateTime`
	__FormDiv__00000002_.IsADivider = false
	__FormDiv__00000002_.IsAStartAccordionGroup = false
	__FormDiv__00000002_.AccordionGroupName = ``
	__FormDiv__00000002_.IsAEndAccordionGroup = false

	__FormDiv__00000003_.Name = `Edit Assoc`
	__FormDiv__00000003_.IsADivider = false
	__FormDiv__00000003_.IsAStartAccordionGroup = false
	__FormDiv__00000003_.AccordionGroupName = ``
	__FormDiv__00000003_.IsAEndAccordionGroup = false

	__FormDiv__00000004_.Name = `Float64 - level of x`
	__FormDiv__00000004_.IsADivider = false
	__FormDiv__00000004_.IsAStartAccordionGroup = false
	__FormDiv__00000004_.AccordionGroupName = ``
	__FormDiv__00000004_.IsAEndAccordionGroup = false

	__FormDiv__00000005_.Name = `Form Div First Name  - Name`
	__FormDiv__00000005_.IsADivider = false
	__FormDiv__00000005_.IsAStartAccordionGroup = false
	__FormDiv__00000005_.AccordionGroupName = ``
	__FormDiv__00000005_.IsAEndAccordionGroup = false

	__FormDiv__00000006_.Name = `Int - Age`
	__FormDiv__00000006_.IsADivider = false
	__FormDiv__00000006_.IsAStartAccordionGroup = false
	__FormDiv__00000006_.AccordionGroupName = ``
	__FormDiv__00000006_.IsAEndAccordionGroup = false

	__FormDiv__00000007_.Name = `Select (can be set empty)`
	__FormDiv__00000007_.IsADivider = false
	__FormDiv__00000007_.IsAStartAccordionGroup = false
	__FormDiv__00000007_.AccordionGroupName = ``
	__FormDiv__00000007_.IsAEndAccordionGroup = false

	__FormDiv__00000008_.Name = `Select (cannot be set empty)`
	__FormDiv__00000008_.IsADivider = false
	__FormDiv__00000008_.IsAStartAccordionGroup = false
	__FormDiv__00000008_.AccordionGroupName = ``
	__FormDiv__00000008_.IsAEndAccordionGroup = false

	__FormDiv__00000009_.Name = `Text Area`
	__FormDiv__00000009_.IsADivider = false
	__FormDiv__00000009_.IsAStartAccordionGroup = false
	__FormDiv__00000009_.AccordionGroupName = ``
	__FormDiv__00000009_.IsAEndAccordionGroup = false

	__FormDiv__00000010_.Name = `Divider`
	__FormDiv__00000010_.IsADivider = true
	__FormDiv__00000010_.IsAStartAccordionGroup = false
	__FormDiv__00000010_.AccordionGroupName = ``
	__FormDiv__00000010_.IsAEndAccordionGroup = false

	__FormDiv__00000011_.Name = `Accordion Group Start`
	__FormDiv__00000011_.IsADivider = false
	__FormDiv__00000011_.IsAStartAccordionGroup = true
	__FormDiv__00000011_.AccordionGroupName = `The Time Fields`
	__FormDiv__00000011_.IsAEndAccordionGroup = false

	__FormDiv__00000012_.Name = `End of Accordion Group`
	__FormDiv__00000012_.IsADivider = false
	__FormDiv__00000012_.IsAStartAccordionGroup = false
	__FormDiv__00000012_.AccordionGroupName = ``
	__FormDiv__00000012_.IsAEndAccordionGroup = true

	__FormEditAssocButton__00000000_.Name = `Edit Assoc`
	__FormEditAssocButton__00000000_.Label = `Association to B`
	__FormEditAssocButton__00000000_.AssociationStorage = ``
	__FormEditAssocButton__00000000_.HasChanged = false
	__FormEditAssocButton__00000000_.IsForSavePurpose = false
	__FormEditAssocButton__00000000_.HasToolTip = false
	__FormEditAssocButton__00000000_.ToolTipText = ``
	__FormEditAssocButton__00000000_.MatTooltipShowDelay = ``

	__FormField__00000000_.Name = `Age`
	__FormField__00000000_.InputTypeEnum = models.Number
	__FormField__00000000_.Label = `Age`
	__FormField__00000000_.Placeholder = `18`
	__FormField__00000000_.HasBespokeWidth = true
	__FormField__00000000_.BespokeWidthPx = 90
	__FormField__00000000_.HasBespokeHeight = false
	__FormField__00000000_.BespokeHeightPx = 0

	__FormField__00000001_.Name = `Date`
	__FormField__00000001_.InputTypeEnum = models.Date
	__FormField__00000001_.Label = `Date`
	__FormField__00000001_.Placeholder = ``
	__FormField__00000001_.HasBespokeWidth = false
	__FormField__00000001_.BespokeWidthPx = 0
	__FormField__00000001_.HasBespokeHeight = false
	__FormField__00000001_.BespokeHeightPx = 0

	__FormField__00000002_.Name = `DateTime`
	__FormField__00000002_.InputTypeEnum = ""
	__FormField__00000002_.Label = ``
	__FormField__00000002_.Placeholder = ``
	__FormField__00000002_.HasBespokeWidth = false
	__FormField__00000002_.BespokeWidthPx = 0
	__FormField__00000002_.HasBespokeHeight = false
	__FormField__00000002_.BespokeHeightPx = 0

	__FormField__00000003_.Name = `FirstName`
	__FormField__00000003_.InputTypeEnum = models.Text
	__FormField__00000003_.Label = `First Name (label)`
	__FormField__00000003_.Placeholder = ``
	__FormField__00000003_.HasBespokeWidth = false
	__FormField__00000003_.BespokeWidthPx = 0
	__FormField__00000003_.HasBespokeHeight = false
	__FormField__00000003_.BespokeHeightPx = 0

	__FormField__00000004_.Name = `Float64 - level of x`
	__FormField__00000004_.InputTypeEnum = ""
	__FormField__00000004_.Label = `Float`
	__FormField__00000004_.Placeholder = ``
	__FormField__00000004_.HasBespokeWidth = false
	__FormField__00000004_.BespokeWidthPx = 0
	__FormField__00000004_.HasBespokeHeight = false
	__FormField__00000004_.BespokeHeightPx = 0

	__FormField__00000005_.Name = `LastName`
	__FormField__00000005_.InputTypeEnum = models.Text
	__FormField__00000005_.Label = `Last Name (label)`
	__FormField__00000005_.Placeholder = ``
	__FormField__00000005_.HasBespokeWidth = false
	__FormField__00000005_.BespokeWidthPx = 0
	__FormField__00000005_.HasBespokeHeight = false
	__FormField__00000005_.BespokeHeightPx = 0

	__FormField__00000006_.Name = `Select (can be set empty)`
	__FormField__00000006_.InputTypeEnum = ""
	__FormField__00000006_.Label = ``
	__FormField__00000006_.Placeholder = ``
	__FormField__00000006_.HasBespokeWidth = false
	__FormField__00000006_.BespokeWidthPx = 0
	__FormField__00000006_.HasBespokeHeight = false
	__FormField__00000006_.BespokeHeightPx = 0

	__FormField__00000007_.Name = `Select (cannot be set empty)`
	__FormField__00000007_.InputTypeEnum = ""
	__FormField__00000007_.Label = ``
	__FormField__00000007_.Placeholder = ``
	__FormField__00000007_.HasBespokeWidth = false
	__FormField__00000007_.BespokeWidthPx = 0
	__FormField__00000007_.HasBespokeHeight = false
	__FormField__00000007_.BespokeHeightPx = 0

	__FormField__00000008_.Name = `Text Area`
	__FormField__00000008_.InputTypeEnum = ""
	__FormField__00000008_.Label = `Text Area (long multi line text)`
	__FormField__00000008_.Placeholder = ``
	__FormField__00000008_.HasBespokeWidth = false
	__FormField__00000008_.BespokeWidthPx = 0
	__FormField__00000008_.HasBespokeHeight = false
	__FormField__00000008_.BespokeHeightPx = 0

	__FormField__00000009_.Name = `Time`
	__FormField__00000009_.InputTypeEnum = models.Time
	__FormField__00000009_.Label = `Time`
	__FormField__00000009_.Placeholder = ``
	__FormField__00000009_.HasBespokeWidth = false
	__FormField__00000009_.BespokeWidthPx = 0
	__FormField__00000009_.HasBespokeHeight = false
	__FormField__00000009_.BespokeHeightPx = 0

	__FormField__00000010_.Name = `Time`
	__FormField__00000010_.InputTypeEnum = models.Time
	__FormField__00000010_.Label = `Time 2`
	__FormField__00000010_.Placeholder = `Time 2`
	__FormField__00000010_.HasBespokeWidth = false
	__FormField__00000010_.BespokeWidthPx = 0
	__FormField__00000010_.HasBespokeHeight = false
	__FormField__00000010_.BespokeHeightPx = 0

	__FormFieldDate__00000000_.Name = `Time`
	__FormFieldDate__00000000_.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-03-02 00:00:00 +0000 UTC")

	__FormFieldDateTime__00000000_.Name = `DateTime`
	__FormFieldDateTime__00000000_.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2001-01-01 01:02:03 +0000 UTC")

	__FormFieldFloat64__00000000_.Name = `Float64 - level of x`
	__FormFieldFloat64__00000000_.Value = 74.000000
	__FormFieldFloat64__00000000_.HasMinValidator = false
	__FormFieldFloat64__00000000_.MinValue = 0.000000
	__FormFieldFloat64__00000000_.HasMaxValidator = false
	__FormFieldFloat64__00000000_.MaxValue = 0.000000

	__FormFieldInt__00000000_.Name = `Age`
	__FormFieldInt__00000000_.Value = 5
	__FormFieldInt__00000000_.HasMinValidator = true
	__FormFieldInt__00000000_.MinValue = 0
	__FormFieldInt__00000000_.HasMaxValidator = true
	__FormFieldInt__00000000_.MaxValue = 59

	__FormFieldSelect__00000000_.Name = `Select (can be set empty)`
	__FormFieldSelect__00000000_.CanBeEmpty = true
	__FormFieldSelect__00000000_.PreserveInitialOrder = true

	__FormFieldSelect__00000001_.Name = `Select (cannot be set empty)`
	__FormFieldSelect__00000001_.CanBeEmpty = false
	__FormFieldSelect__00000001_.PreserveInitialOrder = false

	__FormFieldString__00000000_.Name = `FirstName`
	__FormFieldString__00000000_.Value = `charles`
	__FormFieldString__00000000_.IsTextArea = false

	__FormFieldString__00000001_.Name = `LastName`
	__FormFieldString__00000001_.Value = `Baudelaires`
	__FormFieldString__00000001_.IsTextArea = false

	__FormFieldString__00000002_.Name = `Text Area`
	__FormFieldString__00000002_.Value = `This is a long text
second line`
	__FormFieldString__00000002_.IsTextArea = true

	__FormFieldTime__00000000_.Name = `Time`
	__FormFieldTime__00000000_.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1970-01-01 09:12:10 +0000 UTC")
	__FormFieldTime__00000000_.Step = 1.000000

	__FormFieldTime__00000001_.Name = `Time 2`
	__FormFieldTime__00000001_.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1970-01-01 15:24:33 +0000 UTC")
	__FormFieldTime__00000001_.Step = 1.000000

	__FormGroup__00000000_.Name = `Form 1`
	__FormGroup__00000000_.Label = `Instance`
	__FormGroup__00000000_.TypeLabel = `TypeLabel`
	__FormGroup__00000000_.HasSuppressButton = true
	__FormGroup__00000000_.HasSuppressButtonBeenPressed = true

	__FormSortAssocButton__00000000_.Name = `Sort Button`
	__FormSortAssocButton__00000000_.Label = `Sort Button`
	__FormSortAssocButton__00000000_.HasToolTip = false
	__FormSortAssocButton__00000000_.ToolTipText = ``
	__FormSortAssocButton__00000000_.MatTooltipShowDelay = ``

	__Option__00000000_.Name = `A - Select (cannot be set empty)`

	__Option__00000001_.Name = `B- Select (cannot be set empty)`

	__Option__00000002_.Name = `Option 1 - Select (can be set empty)`

	__Option__00000003_.Name = `Option 2 - Select (can be set empty)`

	__Option__00000004_.Name = `Option 3 - Select (can be set empty)`

	__Option__00000005_.Name = `Option 4 - Select (can be set empty)`

	// insertion point for setup of pointers
	__FormDiv__00000000_.CheckBoxs = append(__FormDiv__00000000_.CheckBoxs, __CheckBox__00000000_)
	__FormDiv__00000000_.FormEditAssocButton = nil
	__FormDiv__00000000_.FormSortAssocButton = nil
	__FormDiv__00000001_.FormFields = append(__FormDiv__00000001_.FormFields, __FormField__00000001_)
	__FormDiv__00000001_.FormFields = append(__FormDiv__00000001_.FormFields, __FormField__00000009_)
	__FormDiv__00000001_.FormFields = append(__FormDiv__00000001_.FormFields, __FormField__00000010_)
	__FormDiv__00000001_.FormEditAssocButton = nil
	__FormDiv__00000001_.FormSortAssocButton = nil
	__FormDiv__00000002_.FormFields = append(__FormDiv__00000002_.FormFields, __FormField__00000002_)
	__FormDiv__00000002_.FormEditAssocButton = nil
	__FormDiv__00000002_.FormSortAssocButton = nil
	__FormDiv__00000003_.FormEditAssocButton = __FormEditAssocButton__00000000_
	__FormDiv__00000003_.FormSortAssocButton = __FormSortAssocButton__00000000_
	__FormDiv__00000004_.FormFields = append(__FormDiv__00000004_.FormFields, __FormField__00000004_)
	__FormDiv__00000004_.FormEditAssocButton = nil
	__FormDiv__00000004_.FormSortAssocButton = nil
	__FormDiv__00000005_.FormFields = append(__FormDiv__00000005_.FormFields, __FormField__00000003_)
	__FormDiv__00000005_.FormFields = append(__FormDiv__00000005_.FormFields, __FormField__00000005_)
	__FormDiv__00000005_.FormEditAssocButton = nil
	__FormDiv__00000005_.FormSortAssocButton = nil
	__FormDiv__00000006_.FormFields = append(__FormDiv__00000006_.FormFields, __FormField__00000000_)
	__FormDiv__00000006_.FormEditAssocButton = nil
	__FormDiv__00000006_.FormSortAssocButton = nil
	__FormDiv__00000007_.FormFields = append(__FormDiv__00000007_.FormFields, __FormField__00000006_)
	__FormDiv__00000007_.FormEditAssocButton = nil
	__FormDiv__00000007_.FormSortAssocButton = nil
	__FormDiv__00000008_.FormFields = append(__FormDiv__00000008_.FormFields, __FormField__00000007_)
	__FormDiv__00000008_.FormEditAssocButton = nil
	__FormDiv__00000008_.FormSortAssocButton = nil
	__FormDiv__00000009_.FormFields = append(__FormDiv__00000009_.FormFields, __FormField__00000008_)
	__FormDiv__00000009_.FormEditAssocButton = nil
	__FormDiv__00000009_.FormSortAssocButton = nil
	__FormDiv__00000010_.FormEditAssocButton = nil
	__FormDiv__00000010_.FormSortAssocButton = nil
	__FormDiv__00000011_.FormEditAssocButton = nil
	__FormDiv__00000011_.FormSortAssocButton = nil
	__FormDiv__00000012_.FormEditAssocButton = nil
	__FormDiv__00000012_.FormSortAssocButton = nil
	__FormField__00000000_.FormFieldString = nil
	__FormField__00000000_.FormFieldFloat64 = nil
	__FormField__00000000_.FormFieldInt = __FormFieldInt__00000000_
	__FormField__00000000_.FormFieldDate = nil
	__FormField__00000000_.FormFieldTime = nil
	__FormField__00000000_.FormFieldDateTime = nil
	__FormField__00000000_.FormFieldSelect = nil
	__FormField__00000001_.FormFieldString = nil
	__FormField__00000001_.FormFieldFloat64 = nil
	__FormField__00000001_.FormFieldInt = nil
	__FormField__00000001_.FormFieldDate = __FormFieldDate__00000000_
	__FormField__00000001_.FormFieldTime = nil
	__FormField__00000001_.FormFieldDateTime = nil
	__FormField__00000001_.FormFieldSelect = nil
	__FormField__00000002_.FormFieldString = nil
	__FormField__00000002_.FormFieldFloat64 = nil
	__FormField__00000002_.FormFieldInt = nil
	__FormField__00000002_.FormFieldDate = nil
	__FormField__00000002_.FormFieldTime = nil
	__FormField__00000002_.FormFieldDateTime = __FormFieldDateTime__00000000_
	__FormField__00000002_.FormFieldSelect = nil
	__FormField__00000003_.FormFieldString = __FormFieldString__00000000_
	__FormField__00000003_.FormFieldFloat64 = nil
	__FormField__00000003_.FormFieldInt = nil
	__FormField__00000003_.FormFieldDate = nil
	__FormField__00000003_.FormFieldTime = nil
	__FormField__00000003_.FormFieldDateTime = nil
	__FormField__00000003_.FormFieldSelect = nil
	__FormField__00000004_.FormFieldString = nil
	__FormField__00000004_.FormFieldFloat64 = __FormFieldFloat64__00000000_
	__FormField__00000004_.FormFieldInt = nil
	__FormField__00000004_.FormFieldDate = nil
	__FormField__00000004_.FormFieldTime = nil
	__FormField__00000004_.FormFieldDateTime = nil
	__FormField__00000004_.FormFieldSelect = nil
	__FormField__00000005_.FormFieldString = __FormFieldString__00000001_
	__FormField__00000005_.FormFieldFloat64 = nil
	__FormField__00000005_.FormFieldInt = nil
	__FormField__00000005_.FormFieldDate = nil
	__FormField__00000005_.FormFieldTime = nil
	__FormField__00000005_.FormFieldDateTime = nil
	__FormField__00000005_.FormFieldSelect = nil
	__FormField__00000006_.FormFieldString = nil
	__FormField__00000006_.FormFieldFloat64 = nil
	__FormField__00000006_.FormFieldInt = nil
	__FormField__00000006_.FormFieldDate = nil
	__FormField__00000006_.FormFieldTime = nil
	__FormField__00000006_.FormFieldDateTime = nil
	__FormField__00000006_.FormFieldSelect = __FormFieldSelect__00000000_
	__FormField__00000007_.FormFieldString = nil
	__FormField__00000007_.FormFieldFloat64 = nil
	__FormField__00000007_.FormFieldInt = nil
	__FormField__00000007_.FormFieldDate = nil
	__FormField__00000007_.FormFieldTime = nil
	__FormField__00000007_.FormFieldDateTime = nil
	__FormField__00000007_.FormFieldSelect = __FormFieldSelect__00000001_
	__FormField__00000008_.FormFieldString = __FormFieldString__00000002_
	__FormField__00000008_.FormFieldFloat64 = nil
	__FormField__00000008_.FormFieldInt = nil
	__FormField__00000008_.FormFieldDate = nil
	__FormField__00000008_.FormFieldTime = nil
	__FormField__00000008_.FormFieldDateTime = nil
	__FormField__00000008_.FormFieldSelect = nil
	__FormField__00000009_.FormFieldString = nil
	__FormField__00000009_.FormFieldFloat64 = nil
	__FormField__00000009_.FormFieldInt = nil
	__FormField__00000009_.FormFieldDate = nil
	__FormField__00000009_.FormFieldTime = __FormFieldTime__00000000_
	__FormField__00000009_.FormFieldDateTime = nil
	__FormField__00000009_.FormFieldSelect = nil
	__FormField__00000010_.FormFieldString = nil
	__FormField__00000010_.FormFieldFloat64 = nil
	__FormField__00000010_.FormFieldInt = nil
	__FormField__00000010_.FormFieldDate = nil
	__FormField__00000010_.FormFieldTime = __FormFieldTime__00000001_
	__FormField__00000010_.FormFieldDateTime = nil
	__FormField__00000010_.FormFieldSelect = nil
	__FormFieldSelect__00000000_.Value = __Option__00000005_
	__FormFieldSelect__00000000_.Options = append(__FormFieldSelect__00000000_.Options, __Option__00000004_)
	__FormFieldSelect__00000000_.Options = append(__FormFieldSelect__00000000_.Options, __Option__00000005_)
	__FormFieldSelect__00000000_.Options = append(__FormFieldSelect__00000000_.Options, __Option__00000002_)
	__FormFieldSelect__00000000_.Options = append(__FormFieldSelect__00000000_.Options, __Option__00000003_)
	__FormFieldSelect__00000001_.Value = __Option__00000000_
	__FormFieldSelect__00000001_.Options = append(__FormFieldSelect__00000001_.Options, __Option__00000000_)
	__FormFieldSelect__00000001_.Options = append(__FormFieldSelect__00000001_.Options, __Option__00000001_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000009_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000000_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000010_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000011_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000001_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000002_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000012_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000004_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000003_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000006_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000005_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000007_)
	__FormGroup__00000000_.FormDivs = append(__FormGroup__00000000_.FormDivs, __FormDiv__00000008_)
	__FormSortAssocButton__00000000_.FormEditAssocButton = nil
}
