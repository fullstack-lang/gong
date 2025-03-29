package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/table/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_form_stage models.Stage
var ___dummy__Time_form_stage time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the _ gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["form_stage"] = _
// }

// _ will stage objects of database "form_stage"
func _(stage *models.Stage) {

	// Declaration of instances to stage

	// Declarations of staged instances of Cell

	// Declarations of staged instances of CellBoolean

	// Declarations of staged instances of CellFloat64

	// Declarations of staged instances of CellIcon

	// Declarations of staged instances of CellInt

	// Declarations of staged instances of CellString
	__CellString__000000_A := (&models.CellString{Name: `A`}).Stage(stage)

	// Declarations of staged instances of CheckBox
	__CheckBox__000000_Boolean := (&models.CheckBox{Name: `Boolean`}).Stage(stage)

	// Declarations of staged instances of DisplayedColumn

	// Declarations of staged instances of FormDiv
	__FormDiv__000000_CheckBoxs := (&models.FormDiv{Name: `CheckBoxs`}).Stage(stage)
	__FormDiv__000001_Date_Time := (&models.FormDiv{Name: `Date - Time`}).Stage(stage)
	__FormDiv__000002_DateTime := (&models.FormDiv{Name: `DateTime`}).Stage(stage)
	__FormDiv__000003_Edit_Assoc := (&models.FormDiv{Name: `Edit Assoc`}).Stage(stage)
	__FormDiv__000004_Float64_level_of_x := (&models.FormDiv{Name: `Float64 - level of x`}).Stage(stage)
	__FormDiv__000005_Form_Div_First_Name_Name := (&models.FormDiv{Name: `Form Div First Name  - Name`}).Stage(stage)
	__FormDiv__000006_Int_Age := (&models.FormDiv{Name: `Int - Age`}).Stage(stage)
	__FormDiv__000007_Select_can_be_set_empty_ := (&models.FormDiv{Name: `Select (can be set empty)`}).Stage(stage)
	__FormDiv__000008_Select_cannot_be_set_empty_ := (&models.FormDiv{Name: `Select (cannot be set empty)`}).Stage(stage)
	__FormDiv__000009_Text_Area := (&models.FormDiv{Name: `Text Area`}).Stage(stage)

	// Declarations of staged instances of FormEditAssocButton
	__FormEditAssocButton__000000_Edit_Assoc := (&models.FormEditAssocButton{Name: `Edit Assoc`}).Stage(stage)

	// Declarations of staged instances of FormField
	__FormField__000000_Age := (&models.FormField{Name: `Age`}).Stage(stage)
	__FormField__000001_Date := (&models.FormField{Name: `Date`}).Stage(stage)
	__FormField__000002_DateTime := (&models.FormField{Name: `DateTime`}).Stage(stage)
	__FormField__000003_FirstName := (&models.FormField{Name: `FirstName`}).Stage(stage)
	__FormField__000004_Float64_level_of_x := (&models.FormField{Name: `Float64 - level of x`}).Stage(stage)
	__FormField__000005_LastName := (&models.FormField{Name: `LastName`}).Stage(stage)
	__FormField__000006_Select_can_be_set_empty_ := (&models.FormField{Name: `Select (can be set empty)`}).Stage(stage)
	__FormField__000007_Select_cannot_be_set_empty_ := (&models.FormField{Name: `Select (cannot be set empty)`}).Stage(stage)
	__FormField__000008_Text_Area := (&models.FormField{Name: `Text Area`}).Stage(stage)
	__FormField__000009_Time := (&models.FormField{Name: `Time`}).Stage(stage)

	// Declarations of staged instances of FormFieldDate
	__FormFieldDate__000000_Time := (&models.FormFieldDate{Name: `Time`}).Stage(stage)

	// Declarations of staged instances of FormFieldDateTime
	__FormFieldDateTime__000000_DateTime := (&models.FormFieldDateTime{Name: `DateTime`}).Stage(stage)

	// Declarations of staged instances of FormFieldFloat64
	__FormFieldFloat64__000000_Float64_level_of_x := (&models.FormFieldFloat64{Name: `Float64 - level of x`}).Stage(stage)

	// Declarations of staged instances of FormFieldInt
	__FormFieldInt__000000_Age := (&models.FormFieldInt{Name: `Age`}).Stage(stage)

	// Declarations of staged instances of FormFieldSelect
	__FormFieldSelect__000000_Select_can_be_set_empty_ := (&models.FormFieldSelect{Name: `Select (can be set empty)`}).Stage(stage)
	__FormFieldSelect__000001_Select_cannot_be_set_empty_ := (&models.FormFieldSelect{Name: `Select (cannot be set empty)`}).Stage(stage)

	// Declarations of staged instances of FormFieldString
	__FormFieldString__000000_FirstName := (&models.FormFieldString{Name: `FirstName`}).Stage(stage)
	__FormFieldString__000001_LastName := (&models.FormFieldString{Name: `LastName`}).Stage(stage)
	__FormFieldString__000002_Text_Area := (&models.FormFieldString{Name: `Text Area`}).Stage(stage)

	// Declarations of staged instances of FormFieldTime
	__FormFieldTime__000000_Time := (&models.FormFieldTime{Name: `Time`}).Stage(stage)

	// Declarations of staged instances of FormGroup
	__FormGroup__000000_Form_1 := (&models.FormGroup{Name: `Form 1`}).Stage(stage)

	// Declarations of staged instances of FormSortAssocButton
	__FormSortAssocButton__000000_Sort_Button := (&models.FormSortAssocButton{Name: `Sort Button`}).Stage(stage)

	// Declarations of staged instances of Option
	__Option__000000_A_Select_cannot_be_set_empty_ := (&models.Option{Name: `A - Select (cannot be set empty)`}).Stage(stage)
	__Option__000001_B_Select_cannot_be_set_empty_ := (&models.Option{Name: `B- Select (cannot be set empty)`}).Stage(stage)
	__Option__000002_Option_1_Select_can_be_set_empty_ := (&models.Option{Name: `Option 1 - Select (can be set empty)`}).Stage(stage)
	__Option__000003_Option_2_Select_can_be_set_empty_ := (&models.Option{Name: `Option 2 - Select (can be set empty)`}).Stage(stage)
	__Option__000004_Option_3_Select_can_be_set_empty_ := (&models.Option{Name: `Option 3 - Select (can be set empty)`}).Stage(stage)
	__Option__000005_Option_4_Select_can_be_set_empty_ := (&models.Option{Name: `Option 4 - Select (can be set empty)`}).Stage(stage)

	// Declarations of staged instances of Row

	// Declarations of staged instances of Table

	// Setup of values

	// CellString values setup
	__CellString__000000_A.Name = `A`
	__CellString__000000_A.Value = `A`

	// CheckBox values setup
	__CheckBox__000000_Boolean.Name = `Boolean`
	__CheckBox__000000_Boolean.Value = true

	// FormDiv values setup
	__FormDiv__000000_CheckBoxs.Name = `CheckBoxs`

	// FormDiv values setup
	__FormDiv__000001_Date_Time.Name = `Date - Time`

	// FormDiv values setup
	__FormDiv__000002_DateTime.Name = `DateTime`

	// FormDiv values setup
	__FormDiv__000003_Edit_Assoc.Name = `Edit Assoc`

	// FormDiv values setup
	__FormDiv__000004_Float64_level_of_x.Name = `Float64 - level of x`

	// FormDiv values setup
	__FormDiv__000005_Form_Div_First_Name_Name.Name = `Form Div First Name  - Name`

	// FormDiv values setup
	__FormDiv__000006_Int_Age.Name = `Int - Age`

	// FormDiv values setup
	__FormDiv__000007_Select_can_be_set_empty_.Name = `Select (can be set empty)`

	// FormDiv values setup
	__FormDiv__000008_Select_cannot_be_set_empty_.Name = `Select (cannot be set empty)`

	// FormDiv values setup
	__FormDiv__000009_Text_Area.Name = `Text Area`

	// FormEditAssocButton values setup
	__FormEditAssocButton__000000_Edit_Assoc.Name = `Edit Assoc`
	__FormEditAssocButton__000000_Edit_Assoc.Label = `Association to B`

	// FormField values setup
	__FormField__000000_Age.Name = `Age`
	__FormField__000000_Age.InputTypeEnum = models.Number
	__FormField__000000_Age.Label = `Age`
	__FormField__000000_Age.Placeholder = `18`
	__FormField__000000_Age.HasBespokeWidth = true
	__FormField__000000_Age.BespokeWidthPx = 90
	__FormField__000000_Age.HasBespokeHeight = false
	__FormField__000000_Age.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000001_Date.Name = `Date`
	__FormField__000001_Date.InputTypeEnum = models.Date
	__FormField__000001_Date.Label = `Date`
	__FormField__000001_Date.Placeholder = ``
	__FormField__000001_Date.HasBespokeWidth = false
	__FormField__000001_Date.BespokeWidthPx = 0
	__FormField__000001_Date.HasBespokeHeight = false
	__FormField__000001_Date.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000002_DateTime.Name = `DateTime`
	__FormField__000002_DateTime.Label = ``
	__FormField__000002_DateTime.Placeholder = ``
	__FormField__000002_DateTime.HasBespokeWidth = false
	__FormField__000002_DateTime.BespokeWidthPx = 0
	__FormField__000002_DateTime.HasBespokeHeight = false
	__FormField__000002_DateTime.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000003_FirstName.Name = `FirstName`
	__FormField__000003_FirstName.InputTypeEnum = models.Text
	__FormField__000003_FirstName.Label = `First Name (label)`
	__FormField__000003_FirstName.Placeholder = ``
	__FormField__000003_FirstName.HasBespokeWidth = false
	__FormField__000003_FirstName.BespokeWidthPx = 0
	__FormField__000003_FirstName.HasBespokeHeight = false
	__FormField__000003_FirstName.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000004_Float64_level_of_x.Name = `Float64 - level of x`
	__FormField__000004_Float64_level_of_x.Label = `Float`
	__FormField__000004_Float64_level_of_x.Placeholder = ``
	__FormField__000004_Float64_level_of_x.HasBespokeWidth = false
	__FormField__000004_Float64_level_of_x.BespokeWidthPx = 0
	__FormField__000004_Float64_level_of_x.HasBespokeHeight = false
	__FormField__000004_Float64_level_of_x.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000005_LastName.Name = `LastName`
	__FormField__000005_LastName.InputTypeEnum = models.Text
	__FormField__000005_LastName.Label = `Last Name (label)`
	__FormField__000005_LastName.Placeholder = ``
	__FormField__000005_LastName.HasBespokeWidth = false
	__FormField__000005_LastName.BespokeWidthPx = 0
	__FormField__000005_LastName.HasBespokeHeight = false
	__FormField__000005_LastName.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000006_Select_can_be_set_empty_.Name = `Select (can be set empty)`
	__FormField__000006_Select_can_be_set_empty_.Label = ``
	__FormField__000006_Select_can_be_set_empty_.Placeholder = ``
	__FormField__000006_Select_can_be_set_empty_.HasBespokeWidth = false
	__FormField__000006_Select_can_be_set_empty_.BespokeWidthPx = 0
	__FormField__000006_Select_can_be_set_empty_.HasBespokeHeight = false
	__FormField__000006_Select_can_be_set_empty_.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000007_Select_cannot_be_set_empty_.Name = `Select (cannot be set empty)`
	__FormField__000007_Select_cannot_be_set_empty_.Label = ``
	__FormField__000007_Select_cannot_be_set_empty_.Placeholder = ``
	__FormField__000007_Select_cannot_be_set_empty_.HasBespokeWidth = false
	__FormField__000007_Select_cannot_be_set_empty_.BespokeWidthPx = 0
	__FormField__000007_Select_cannot_be_set_empty_.HasBespokeHeight = false
	__FormField__000007_Select_cannot_be_set_empty_.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000008_Text_Area.Name = `Text Area`
	__FormField__000008_Text_Area.Label = `Text Area (long multi line text)`
	__FormField__000008_Text_Area.Placeholder = ``
	__FormField__000008_Text_Area.HasBespokeWidth = false
	__FormField__000008_Text_Area.BespokeWidthPx = 0
	__FormField__000008_Text_Area.HasBespokeHeight = false
	__FormField__000008_Text_Area.BespokeHeightPx = 0

	// FormField values setup
	__FormField__000009_Time.Name = `Time`
	__FormField__000009_Time.InputTypeEnum = models.Time
	__FormField__000009_Time.Label = `Time`
	__FormField__000009_Time.Placeholder = ``
	__FormField__000009_Time.HasBespokeWidth = false
	__FormField__000009_Time.BespokeWidthPx = 0
	__FormField__000009_Time.HasBespokeHeight = false
	__FormField__000009_Time.BespokeHeightPx = 0

	// FormFieldDate values setup
	__FormFieldDate__000000_Time.Name = `Time`
	__FormFieldDate__000000_Time.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-03-02 00:00:00 +0000 UTC")

	// FormFieldDateTime values setup
	__FormFieldDateTime__000000_DateTime.Name = `DateTime`
	__FormFieldDateTime__000000_DateTime.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2001-01-01 01:02:03 +0000 UTC")

	// FormFieldFloat64 values setup
	__FormFieldFloat64__000000_Float64_level_of_x.Name = `Float64 - level of x`
	__FormFieldFloat64__000000_Float64_level_of_x.Value = 74.000000
	__FormFieldFloat64__000000_Float64_level_of_x.HasMinValidator = false
	__FormFieldFloat64__000000_Float64_level_of_x.MinValue = 0.000000
	__FormFieldFloat64__000000_Float64_level_of_x.HasMaxValidator = false
	__FormFieldFloat64__000000_Float64_level_of_x.MaxValue = 0.000000

	// FormFieldInt values setup
	__FormFieldInt__000000_Age.Name = `Age`
	__FormFieldInt__000000_Age.Value = 5
	__FormFieldInt__000000_Age.HasMinValidator = true
	__FormFieldInt__000000_Age.MinValue = 0
	__FormFieldInt__000000_Age.HasMaxValidator = true
	__FormFieldInt__000000_Age.MaxValue = 59

	// FormFieldSelect values setup
	__FormFieldSelect__000000_Select_can_be_set_empty_.Name = `Select (can be set empty)`
	__FormFieldSelect__000000_Select_can_be_set_empty_.CanBeEmpty = true

	// FormFieldSelect values setup
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.Name = `Select (cannot be set empty)`
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.CanBeEmpty = false

	// FormFieldString values setup
	__FormFieldString__000000_FirstName.Name = `FirstName`
	__FormFieldString__000000_FirstName.Value = `charles`
	__FormFieldString__000000_FirstName.IsTextArea = false

	// FormFieldString values setup
	__FormFieldString__000001_LastName.Name = `LastName`
	__FormFieldString__000001_LastName.Value = `Baudelaire`
	__FormFieldString__000001_LastName.IsTextArea = false

	// FormFieldString values setup
	__FormFieldString__000002_Text_Area.Name = `Text Area`
	__FormFieldString__000002_Text_Area.Value = `This is a long text
second line`
	__FormFieldString__000002_Text_Area.IsTextArea = true

	// FormFieldTime values setup
	__FormFieldTime__000000_Time.Name = `Time`
	__FormFieldTime__000000_Time.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "1970-01-01 09:12:10 +0000 UTC")
	__FormFieldTime__000000_Time.Step = 1.000000

	// FormGroup values setup
	__FormGroup__000000_Form_1.Name = `Form 1`
	__FormGroup__000000_Form_1.Label = `<FormGroup.Label>`
	__FormGroup__000000_Form_1.HasSuppressButton = true
	__FormGroup__000000_Form_1.HasSuppressButtonBeenPressed = true

	// FormSortAssocButton values setup
	__FormSortAssocButton__000000_Sort_Button.Name = `Sort Button`
	__FormSortAssocButton__000000_Sort_Button.Label = `Sort Button`

	// Option values setup
	__Option__000000_A_Select_cannot_be_set_empty_.Name = `A - Select (cannot be set empty)`

	// Option values setup
	__Option__000001_B_Select_cannot_be_set_empty_.Name = `B- Select (cannot be set empty)`

	// Option values setup
	__Option__000002_Option_1_Select_can_be_set_empty_.Name = `Option 1 - Select (can be set empty)`

	// Option values setup
	__Option__000003_Option_2_Select_can_be_set_empty_.Name = `Option 2 - Select (can be set empty)`

	// Option values setup
	__Option__000004_Option_3_Select_can_be_set_empty_.Name = `Option 3 - Select (can be set empty)`

	// Option values setup
	__Option__000005_Option_4_Select_can_be_set_empty_.Name = `Option 4 - Select (can be set empty)`

	// Setup of pointers
	__FormDiv__000000_CheckBoxs.CheckBoxs = append(__FormDiv__000000_CheckBoxs.CheckBoxs, __CheckBox__000000_Boolean)
	__FormDiv__000001_Date_Time.FormFields = append(__FormDiv__000001_Date_Time.FormFields, __FormField__000001_Date)
	__FormDiv__000001_Date_Time.FormFields = append(__FormDiv__000001_Date_Time.FormFields, __FormField__000009_Time)
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
	__FormFieldSelect__000000_Select_can_be_set_empty_.Value = __Option__000005_Option_4_Select_can_be_set_empty_
	__FormFieldSelect__000000_Select_can_be_set_empty_.Options = append(__FormFieldSelect__000000_Select_can_be_set_empty_.Options, __Option__000004_Option_3_Select_can_be_set_empty_)
	__FormFieldSelect__000000_Select_can_be_set_empty_.Options = append(__FormFieldSelect__000000_Select_can_be_set_empty_.Options, __Option__000005_Option_4_Select_can_be_set_empty_)
	__FormFieldSelect__000000_Select_can_be_set_empty_.Options = append(__FormFieldSelect__000000_Select_can_be_set_empty_.Options, __Option__000002_Option_1_Select_can_be_set_empty_)
	__FormFieldSelect__000000_Select_can_be_set_empty_.Options = append(__FormFieldSelect__000000_Select_can_be_set_empty_.Options, __Option__000003_Option_2_Select_can_be_set_empty_)
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.Value = __Option__000000_A_Select_cannot_be_set_empty_
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.Options = append(__FormFieldSelect__000001_Select_cannot_be_set_empty_.Options, __Option__000000_A_Select_cannot_be_set_empty_)
	__FormFieldSelect__000001_Select_cannot_be_set_empty_.Options = append(__FormFieldSelect__000001_Select_cannot_be_set_empty_.Options, __Option__000001_B_Select_cannot_be_set_empty_)
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
}
