package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/table/go/models"
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

	const __write__local_time = "2025-10-04 03:45:16.159826 CEST"
	const __write__utc_time__ = "2025-10-04 01:45:16.159826 UTC"

	const __commitId__ = "0000000013"

	// Declaration of instances to stage

	__CellString__000000_A := (&models.CellString{}).Stage(stage)

	__CheckBox__000000_Boolean := (&models.CheckBox{}).Stage(stage)

	__FormDiv__000000_CheckBoxs := (&models.FormDiv{}).Stage(stage)
	__FormDiv__000001_Date_Time := (&models.FormDiv{}).Stage(stage)
	__FormDiv__000002_DateTime := (&models.FormDiv{}).Stage(stage)
	__FormDiv__000003_Edit_Assoc := (&models.FormDiv{}).Stage(stage)
	__FormDiv__000004_Float64_level_of_x := (&models.FormDiv{}).Stage(stage)
	__FormDiv__000005_Form_Div_First_Name_Name := (&models.FormDiv{}).Stage(stage)
	__FormDiv__000006_Int_Age := (&models.FormDiv{}).Stage(stage)
	__FormDiv__000007_Select_can_be_set_empty_ := (&models.FormDiv{}).Stage(stage)
	__FormDiv__000008_Select_cannot_be_set_empty_ := (&models.FormDiv{}).Stage(stage)
	__FormDiv__000009_Text_Area := (&models.FormDiv{}).Stage(stage)

	__FormEditAssocButton__000000_Edit_Assoc := (&models.FormEditAssocButton{}).Stage(stage)

	__FormField__000000_Age := (&models.FormField{}).Stage(stage)
	__FormField__000001_Date := (&models.FormField{}).Stage(stage)
	__FormField__000002_DateTime := (&models.FormField{}).Stage(stage)
	__FormField__000003_FirstName := (&models.FormField{}).Stage(stage)
	__FormField__000004_Float64_level_of_x := (&models.FormField{}).Stage(stage)
	__FormField__000005_LastName := (&models.FormField{}).Stage(stage)
	__FormField__000006_Select_can_be_set_empty_ := (&models.FormField{}).Stage(stage)
	__FormField__000007_Select_cannot_be_set_empty_ := (&models.FormField{}).Stage(stage)
	__FormField__000008_Text_Area := (&models.FormField{}).Stage(stage)
	__FormField__000009_Time := (&models.FormField{}).Stage(stage)
	__FormField__000010_Time := (&models.FormField{}).Stage(stage)

	__FormFieldDate__000000_Time := (&models.FormFieldDate{}).Stage(stage)

	__FormFieldDateTime__000000_DateTime := (&models.FormFieldDateTime{}).Stage(stage)

	__FormFieldFloat64__000000_Float64_level_of_x := (&models.FormFieldFloat64{}).Stage(stage)

	__FormFieldInt__000000_Age := (&models.FormFieldInt{}).Stage(stage)

	__FormFieldSelect__000000_Select_can_be_set_empty_ := (&models.FormFieldSelect{}).Stage(stage)
	__FormFieldSelect__000001_Select_cannot_be_set_empty_ := (&models.FormFieldSelect{}).Stage(stage)

	__FormFieldString__000000_FirstName := (&models.FormFieldString{}).Stage(stage)
	__FormFieldString__000001_LastName := (&models.FormFieldString{}).Stage(stage)
	__FormFieldString__000002_Text_Area := (&models.FormFieldString{}).Stage(stage)

	__FormFieldTime__000000_Time := (&models.FormFieldTime{}).Stage(stage)
	__FormFieldTime__000001_Time_2 := (&models.FormFieldTime{}).Stage(stage)

	__FormGroup__000000_Form_1 := (&models.FormGroup{}).Stage(stage)

	__FormSortAssocButton__000000_Sort_Button := (&models.FormSortAssocButton{}).Stage(stage)

	__Option__000000_A_Select_cannot_be_set_empty_ := (&models.Option{}).Stage(stage)
	__Option__000001_B_Select_cannot_be_set_empty_ := (&models.Option{}).Stage(stage)
	__Option__000002_Option_1_Select_can_be_set_empty_ := (&models.Option{}).Stage(stage)
	__Option__000003_Option_2_Select_can_be_set_empty_ := (&models.Option{}).Stage(stage)
	__Option__000004_Option_3_Select_can_be_set_empty_ := (&models.Option{}).Stage(stage)
	__Option__000005_Option_4_Select_can_be_set_empty_ := (&models.Option{}).Stage(stage)

	// Setup of values

	__CellString__000000_A.Name = `A`
	__CellString__000000_A.Value = `A`

	__CheckBox__000000_Boolean.Name = `Boolean`
	__CheckBox__000000_Boolean.Value = true

	__FormDiv__000000_CheckBoxs.Name = `CheckBoxs`

	__FormDiv__000001_Date_Time.Name = `Date - Time`

	__FormDiv__000002_DateTime.Name = `DateTime`

	__FormDiv__000003_Edit_Assoc.Name = `Edit Assoc`

	__FormDiv__000004_Float64_level_of_x.Name = `Float64 - level of x`

	__FormDiv__000005_Form_Div_First_Name_Name.Name = `Form Div First Name  - Name`

	__FormDiv__000006_Int_Age.Name = `Int - Age`

	__FormDiv__000007_Select_can_be_set_empty_.Name = `Select (can be set empty)`

	__FormDiv__000008_Select_cannot_be_set_empty_.Name = `Select (cannot be set empty)`

	__FormDiv__000009_Text_Area.Name = `Text Area`

	__FormEditAssocButton__000000_Edit_Assoc.Name = `Edit Assoc`
	__FormEditAssocButton__000000_Edit_Assoc.Label = `Association to B`
	__FormEditAssocButton__000000_Edit_Assoc.AssociationStorage = ``
	__FormEditAssocButton__000000_Edit_Assoc.HasChanged = false
	__FormEditAssocButton__000000_Edit_Assoc.IsForSavePurpose = false
	__FormEditAssocButton__000000_Edit_Assoc.HasToolTip = false
	__FormEditAssocButton__000000_Edit_Assoc.ToolTipText = ``

	__FormField__000000_Age.Name = `Age`
	__FormField__000000_Age.InputTypeEnum = models.Number
	__FormField__000000_Age.Label = `Age`
	__FormField__000000_Age.Placeholder = `18`
	__FormField__000000_Age.HasBespokeWidth = true
	__FormField__000000_Age.BespokeWidthPx = 90
	__FormField__000000_Age.HasBespokeHeight = false
	__FormField__000000_Age.BespokeHeightPx = 0

	__FormField__000001_Date.Name = `Date`
	__FormField__000001_Date.InputTypeEnum = models.Date
	__FormField__000001_Date.Label = `Date`
	__FormField__000001_Date.Placeholder = ``
	__FormField__000001_Date.HasBespokeWidth = false
	__FormField__000001_Date.BespokeWidthPx = 0
	__FormField__000001_Date.HasBespokeHeight = false
	__FormField__000001_Date.BespokeHeightPx = 0

	__FormField__000002_DateTime.Name = `DateTime`
	__FormField__000002_DateTime.Label = ``
	__FormField__000002_DateTime.Placeholder = ``
	__FormField__000002_DateTime.HasBespokeWidth = false
	__FormField__000002_DateTime.BespokeWidthPx = 0
	__FormField__000002_DateTime.HasBespokeHeight = false
	__FormField__000002_DateTime.BespokeHeightPx = 0

	__FormField__000003_FirstName.Name = `FirstName`
	__FormField__000003_FirstName.InputTypeEnum = models.Text
	__FormField__000003_FirstName.Label = `First Name (label)`
	__FormField__000003_FirstName.Placeholder = ``
	__FormField__000003_FirstName.HasBespokeWidth = false
	__FormField__000003_FirstName.BespokeWidthPx = 0
	__FormField__000003_FirstName.HasBespokeHeight = false
	__FormField__000003_FirstName.BespokeHeightPx = 0

	__FormField__000004_Float64_level_of_x.Name = `Float64 - level of x`
	__FormField__000004_Float64_level_of_x.Label = `Float`
	__FormField__000004_Float64_level_of_x.Placeholder = ``
	__FormField__000004_Float64_level_of_x.HasBespokeWidth = false
	__FormField__000004_Float64_level_of_x.BespokeWidthPx = 0
	__FormField__000004_Float64_level_of_x.HasBespokeHeight = false
	__FormField__000004_Float64_level_of_x.BespokeHeightPx = 0

	__FormField__000005_LastName.Name = `LastName`
	__FormField__000005_LastName.InputTypeEnum = models.Text
	__FormField__000005_LastName.Label = `Last Name (label)`
	__FormField__000005_LastName.Placeholder = ``
	__FormField__000005_LastName.HasBespokeWidth = false
	__FormField__000005_LastName.BespokeWidthPx = 0
	__FormField__000005_LastName.HasBespokeHeight = false
	__FormField__000005_LastName.BespokeHeightPx = 0

	__FormField__000006_Select_can_be_set_empty_.Name = `Select (can be set empty)`
	__FormField__000006_Select_can_be_set_empty_.Label = ``
	__FormField__000006_Select_can_be_set_empty_.Placeholder = ``
	__FormField__000006_Select_can_be_set_empty_.HasBespokeWidth = false
	__FormField__000006_Select_can_be_set_empty_.BespokeWidthPx = 0
	__FormField__000006_Select_can_be_set_empty_.HasBespokeHeight = false
	__FormField__000006_Select_can_be_set_empty_.BespokeHeightPx = 0

	__FormField__000007_Select_cannot_be_set_empty_.Name = `Select (cannot be set empty)`
	__FormField__000007_Select_cannot_be_set_empty_.Label = ``
	__FormField__000007_Select_cannot_be_set_empty_.Placeholder = ``
	__FormField__000007_Select_cannot_be_set_empty_.HasBespokeWidth = false
	__FormField__000007_Select_cannot_be_set_empty_.BespokeWidthPx = 0
	__FormField__000007_Select_cannot_be_set_empty_.HasBespokeHeight = false
	__FormField__000007_Select_cannot_be_set_empty_.BespokeHeightPx = 0

	__FormField__000008_Text_Area.Name = `Text Area`
	__FormField__000008_Text_Area.Label = `Text Area (long multi line text)`
	__FormField__000008_Text_Area.Placeholder = ``
	__FormField__000008_Text_Area.HasBespokeWidth = false
	__FormField__000008_Text_Area.BespokeWidthPx = 0
	__FormField__000008_Text_Area.HasBespokeHeight = false
	__FormField__000008_Text_Area.BespokeHeightPx = 0

	__FormField__000009_Time.Name = `Time`
	__FormField__000009_Time.InputTypeEnum = models.Time
	__FormField__000009_Time.Label = `Time`
	__FormField__000009_Time.Placeholder = ``
	__FormField__000009_Time.HasBespokeWidth = false
	__FormField__000009_Time.BespokeWidthPx = 0
	__FormField__000009_Time.HasBespokeHeight = false
	__FormField__000009_Time.BespokeHeightPx = 0

	__FormField__000010_Time.Name = `Time`
	__FormField__000010_Time.InputTypeEnum = models.Time
	__FormField__000010_Time.Label = `Time 2`
	__FormField__000010_Time.Placeholder = `Time 2`
	__FormField__000010_Time.HasBespokeWidth = false
	__FormField__000010_Time.BespokeWidthPx = 0
	__FormField__000010_Time.HasBespokeHeight = false
	__FormField__000010_Time.BespokeHeightPx = 0

	__FormFieldDate__000000_Time.Name = `Time`
	__FormFieldDate__000000_Time.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-03-02 00:00:00 +0000 UTC")

	__FormFieldDateTime__000000_DateTime.Name = `DateTime`
	__FormFieldDateTime__000000_DateTime.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2001-01-01 01:02:03 +0000 UTC")

	__FormFieldFloat64__000000_Float64_level_of_x.Name = `Float64 - level of x`
	__FormFieldFloat64__000000_Float64_level_of_x.Value = 74.000000
	__FormFieldFloat64__000000_Float64_level_of_x.HasMinValidator = false
	__FormFieldFloat64__000000_Float64_level_of_x.MinValue = 0.000000
	__FormFieldFloat64__000000_Float64_level_of_x.HasMaxValidator = false
	__FormFieldFloat64__000000_Float64_level_of_x.MaxValue = 0.000000

	__FormFieldInt__000000_Age.Name = `Age`
	__FormFieldInt__000000_Age.Value = 5
	__FormFieldInt__000000_Age.HasMinValidator = true
	__FormFieldInt__000000_Age.MinValue = 0
	__FormFieldInt__000000_Age.HasMaxValidator = true
	__FormFieldInt__000000_Age.MaxValue = 59

	__FormFieldSelect__000000_Select_can_be_set_empty_.Name = `Select (can be set empty)`
	__FormFieldSelect__000000_Select_can_be_set_empty_.CanBeEmpty = true
	__FormFieldSelect__000000_Select_can_be_set_empty_.PreserveInitialOrder = true

	__FormFieldSelect__000001_Select_cannot_be_set_empty_.Name = `Select (cannot be set empty)`
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.CanBeEmpty = false
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.PreserveInitialOrder = false

	__FormFieldString__000000_FirstName.Name = `FirstName`
	__FormFieldString__000000_FirstName.Value = `charles`
	__FormFieldString__000000_FirstName.IsTextArea = false

	__FormFieldString__000001_LastName.Name = `LastName`
	__FormFieldString__000001_LastName.Value = `Baudelaire`
	__FormFieldString__000001_LastName.IsTextArea = false

	__FormFieldString__000002_Text_Area.Name = `Text Area`
	__FormFieldString__000002_Text_Area.Value = `This is a long text
second line`
	__FormFieldString__000002_Text_Area.IsTextArea = true

	__FormFieldTime__000000_Time.Name = `Time`
	__FormFieldTime__000000_Time.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1970-01-01 09:12:10 +0000 UTC")
	__FormFieldTime__000000_Time.Step = 1.000000

	__FormFieldTime__000001_Time_2.Name = `Time 2`
	__FormFieldTime__000001_Time_2.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1970-01-01 15:24:33 +0000 UTC")
	__FormFieldTime__000001_Time_2.Step = 1.000000

	__FormGroup__000000_Form_1.Name = `Form 1`
	__FormGroup__000000_Form_1.Label = `<FormGroup.Label>`
	__FormGroup__000000_Form_1.HasSuppressButton = true
	__FormGroup__000000_Form_1.HasSuppressButtonBeenPressed = true

	__FormSortAssocButton__000000_Sort_Button.Name = `Sort Button`
	__FormSortAssocButton__000000_Sort_Button.Label = `Sort Button`
	__FormSortAssocButton__000000_Sort_Button.HasToolTip = false
	__FormSortAssocButton__000000_Sort_Button.ToolTipText = ``

	__Option__000000_A_Select_cannot_be_set_empty_.Name = `A - Select (cannot be set empty)`

	__Option__000001_B_Select_cannot_be_set_empty_.Name = `B- Select (cannot be set empty)`

	__Option__000002_Option_1_Select_can_be_set_empty_.Name = `Option 1 - Select (can be set empty)`

	__Option__000003_Option_2_Select_can_be_set_empty_.Name = `Option 2 - Select (can be set empty)`

	__Option__000004_Option_3_Select_can_be_set_empty_.Name = `Option 3 - Select (can be set empty)`

	__Option__000005_Option_4_Select_can_be_set_empty_.Name = `Option 4 - Select (can be set empty)`

	// Setup of pointers
	// setup of CellString instances pointers
	// setup of CheckBox instances pointers
	// setup of FormDiv instances pointers
	__FormDiv__000000_CheckBoxs.CheckBoxs = append(__FormDiv__000000_CheckBoxs.CheckBoxs, __CheckBox__000000_Boolean)
	__FormDiv__000001_Date_Time.FormFields = append(__FormDiv__000001_Date_Time.FormFields, __FormField__000001_Date)
	__FormDiv__000001_Date_Time.FormFields = append(__FormDiv__000001_Date_Time.FormFields, __FormField__000009_Time)
	__FormDiv__000001_Date_Time.FormFields = append(__FormDiv__000001_Date_Time.FormFields, __FormField__000010_Time)
	__FormDiv__000002_DateTime.FormFields = append(__FormDiv__000002_DateTime.FormFields, __FormField__000002_DateTime)
	__FormDiv__000003_Edit_Assoc.FormEditAssocButton = __FormEditAssocButton__000000_Edit_Assoc
	__FormDiv__000003_Edit_Assoc.FormSortAssocButton = __FormSortAssocButton__000000_Sort_Button
	__FormDiv__000004_Float64_level_of_x.FormFields = append(__FormDiv__000004_Float64_level_of_x.FormFields, __FormField__000004_Float64_level_of_x)
	__FormDiv__000005_Form_Div_First_Name_Name.FormFields = append(__FormDiv__000005_Form_Div_First_Name_Name.FormFields, __FormField__000003_FirstName)
	__FormDiv__000005_Form_Div_First_Name_Name.FormFields = append(__FormDiv__000005_Form_Div_First_Name_Name.FormFields, __FormField__000005_LastName)
	__FormDiv__000006_Int_Age.FormFields = append(__FormDiv__000006_Int_Age.FormFields, __FormField__000000_Age)
	__FormDiv__000007_Select_can_be_set_empty_.FormFields = append(__FormDiv__000007_Select_can_be_set_empty_.FormFields, __FormField__000006_Select_can_be_set_empty_)
	__FormDiv__000008_Select_cannot_be_set_empty_.FormFields = append(__FormDiv__000008_Select_cannot_be_set_empty_.FormFields, __FormField__000007_Select_cannot_be_set_empty_)
	__FormDiv__000009_Text_Area.FormFields = append(__FormDiv__000009_Text_Area.FormFields, __FormField__000008_Text_Area)
	// setup of FormEditAssocButton instances pointers
	// setup of FormField instances pointers
	__FormField__000000_Age.FormFieldInt = __FormFieldInt__000000_Age
	__FormField__000001_Date.FormFieldDate = __FormFieldDate__000000_Time
	__FormField__000002_DateTime.FormFieldDateTime = __FormFieldDateTime__000000_DateTime
	__FormField__000003_FirstName.FormFieldString = __FormFieldString__000000_FirstName
	__FormField__000004_Float64_level_of_x.FormFieldFloat64 = __FormFieldFloat64__000000_Float64_level_of_x
	__FormField__000005_LastName.FormFieldString = __FormFieldString__000001_LastName
	__FormField__000006_Select_can_be_set_empty_.FormFieldSelect = __FormFieldSelect__000000_Select_can_be_set_empty_
	__FormField__000007_Select_cannot_be_set_empty_.FormFieldSelect = __FormFieldSelect__000001_Select_cannot_be_set_empty_
	__FormField__000008_Text_Area.FormFieldString = __FormFieldString__000002_Text_Area
	__FormField__000009_Time.FormFieldTime = __FormFieldTime__000000_Time
	__FormField__000010_Time.FormFieldTime = __FormFieldTime__000001_Time_2
	// setup of FormFieldDate instances pointers
	// setup of FormFieldDateTime instances pointers
	// setup of FormFieldFloat64 instances pointers
	// setup of FormFieldInt instances pointers
	// setup of FormFieldSelect instances pointers
	__FormFieldSelect__000000_Select_can_be_set_empty_.Value = __Option__000005_Option_4_Select_can_be_set_empty_
	__FormFieldSelect__000000_Select_can_be_set_empty_.Options = append(__FormFieldSelect__000000_Select_can_be_set_empty_.Options, __Option__000004_Option_3_Select_can_be_set_empty_)
	__FormFieldSelect__000000_Select_can_be_set_empty_.Options = append(__FormFieldSelect__000000_Select_can_be_set_empty_.Options, __Option__000005_Option_4_Select_can_be_set_empty_)
	__FormFieldSelect__000000_Select_can_be_set_empty_.Options = append(__FormFieldSelect__000000_Select_can_be_set_empty_.Options, __Option__000002_Option_1_Select_can_be_set_empty_)
	__FormFieldSelect__000000_Select_can_be_set_empty_.Options = append(__FormFieldSelect__000000_Select_can_be_set_empty_.Options, __Option__000003_Option_2_Select_can_be_set_empty_)
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.Value = __Option__000000_A_Select_cannot_be_set_empty_
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.Options = append(__FormFieldSelect__000001_Select_cannot_be_set_empty_.Options, __Option__000000_A_Select_cannot_be_set_empty_)
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.Options = append(__FormFieldSelect__000001_Select_cannot_be_set_empty_.Options, __Option__000001_B_Select_cannot_be_set_empty_)
	// setup of FormFieldString instances pointers
	// setup of FormFieldTime instances pointers
	// setup of FormGroup instances pointers
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000009_Text_Area)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000000_CheckBoxs)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000001_Date_Time)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000003_Edit_Assoc)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000002_DateTime)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000004_Float64_level_of_x)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000006_Int_Age)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000005_Form_Div_First_Name_Name)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000007_Select_can_be_set_empty_)
	__FormGroup__000000_Form_1.FormDivs = append(__FormGroup__000000_Form_1.FormDivs, __FormDiv__000008_Select_cannot_be_set_empty_)
	// setup of FormSortAssocButton instances pointers
	// setup of Option instances pointers
}

