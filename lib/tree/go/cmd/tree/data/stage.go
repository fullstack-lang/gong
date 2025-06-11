package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
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

	__Button__000000_Button_For_Empty_Space := (&models.Button{}).Stage(stage)
	__Button__000001_PreceedingSVGIcon := (&models.Button{}).Stage(stage)
	__Button__000002_Test_1_2_add := (&models.Button{}).Stage(stage)
	__Button__000003_arrow_circle_left := (&models.Button{}).Stage(stage)
	__Button__000004_dataset := (&models.Button{}).Stage(stage)
	__Button__000005_dynamic_form := (&models.Button{}).Stage(stage)
	__Button__000006_key := (&models.Button{}).Stage(stage)
	__Button__000007_logout := (&models.Button{}).Stage(stage)
	__Button__000008_root_2_add := (&models.Button{}).Stage(stage)
	__Button__000009_settings := (&models.Button{}).Stage(stage)
	__Button__000010_test := (&models.Button{}).Stage(stage)

	__Node__000000_Test_1_2_without_checkbox := (&models.Node{}).Stage(stage)
	__Node__000001_root1 := (&models.Node{}).Stage(stage)
	__Node__000002_root2 := (&models.Node{}).Stage(stage)
	__Node__000003_root3 := (&models.Node{}).Stage(stage)
	__Node__000004_root3_test3 := (&models.Node{}).Stage(stage)
	__Node__000005_test_1_1 := (&models.Node{}).Stage(stage)
	__Node__000006_test_1_2_clickable_ := (&models.Node{}).Stage(stage)
	__Node__000007_test_1_3 := (&models.Node{}).Stage(stage)
	__Node__000008_test_1_3_1_F := (&models.Node{}).Stage(stage)
	__Node__000009_test_1_3_1_svg_icon := (&models.Node{}).Stage(stage)
	__Node__000010_test_1_4 := (&models.Node{}).Stage(stage)
	__Node__000011_test_2_1 := (&models.Node{}).Stage(stage)
	__Node__000012_test3 := (&models.Node{}).Stage(stage)

	__SVGIcon__000000_empty := (&models.SVGIcon{}).Stage(stage)
	__SVGIcon__000001_sync_alt_rotated_90 := (&models.SVGIcon{}).Stage(stage)

	__Tree__000000_test := (&models.Tree{}).Stage(stage)

	// Setup of values

	__Button__000000_Button_For_Empty_Space.Name = `Button For Empty Space`
	__Button__000000_Button_For_Empty_Space.Icon = ``
	__Button__000000_Button_For_Empty_Space.IsDisabled = false
	__Button__000000_Button_For_Empty_Space.HasToolTip = false
	__Button__000000_Button_For_Empty_Space.ToolTipText = ``

	__Button__000001_PreceedingSVGIcon.Name = `PreceedingSVGIcon`
	__Button__000001_PreceedingSVGIcon.Icon = ``
	__Button__000001_PreceedingSVGIcon.IsDisabled = false
	__Button__000001_PreceedingSVGIcon.HasToolTip = false
	__Button__000001_PreceedingSVGIcon.ToolTipText = ``

	__Button__000002_Test_1_2_add.Name = `Test 1.2 add`
	__Button__000002_Test_1_2_add.Icon = `add`
	__Button__000002_Test_1_2_add.IsDisabled = false
	__Button__000002_Test_1_2_add.HasToolTip = false
	__Button__000002_Test_1_2_add.ToolTipText = ``

	__Button__000003_arrow_circle_left.Name = `arrow_circle_left`
	__Button__000003_arrow_circle_left.Icon = `arrow_circle_left`
	__Button__000003_arrow_circle_left.IsDisabled = true
	__Button__000003_arrow_circle_left.HasToolTip = true
	__Button__000003_arrow_circle_left.ToolTipText = `This is a tooltip for the button`
	__Button__000003_arrow_circle_left.ToolTipPosition = models.Above

	__Button__000004_dataset.Name = `dataset`
	__Button__000004_dataset.Icon = `dataset`
	__Button__000004_dataset.IsDisabled = false
	__Button__000004_dataset.HasToolTip = false
	__Button__000004_dataset.ToolTipText = ``

	__Button__000005_dynamic_form.Name = `dynamic_form`
	__Button__000005_dynamic_form.Icon = `dynamic_form`
	__Button__000005_dynamic_form.IsDisabled = false
	__Button__000005_dynamic_form.HasToolTip = false
	__Button__000005_dynamic_form.ToolTipText = ``

	__Button__000006_key.Name = `key`
	__Button__000006_key.Icon = `key`
	__Button__000006_key.IsDisabled = false
	__Button__000006_key.HasToolTip = false
	__Button__000006_key.ToolTipText = ``

	__Button__000007_logout.Name = `logout`
	__Button__000007_logout.Icon = `logout`
	__Button__000007_logout.IsDisabled = false
	__Button__000007_logout.HasToolTip = false
	__Button__000007_logout.ToolTipText = ``

	__Button__000008_root_2_add.Name = `root 2 - add`
	__Button__000008_root_2_add.Icon = `add`
	__Button__000008_root_2_add.IsDisabled = false
	__Button__000008_root_2_add.HasToolTip = false
	__Button__000008_root_2_add.ToolTipText = ``

	__Button__000009_settings.Name = `settings`
	__Button__000009_settings.Icon = `settings`
	__Button__000009_settings.IsDisabled = false
	__Button__000009_settings.HasToolTip = false
	__Button__000009_settings.ToolTipText = ``

	__Button__000010_test.Name = `test`
	__Button__000010_test.Icon = `edit`
	__Button__000010_test.IsDisabled = false
	__Button__000010_test.HasToolTip = false
	__Button__000010_test.ToolTipText = ``

	__Node__000000_Test_1_2_without_checkbox.Name = `Test 1.2 without checkbox`
	__Node__000000_Test_1_2_without_checkbox.FontStyle = models.NORMAL
	__Node__000000_Test_1_2_without_checkbox.BackgroundColor = ``
	__Node__000000_Test_1_2_without_checkbox.IsExpanded = false
	__Node__000000_Test_1_2_without_checkbox.HasCheckboxButton = false
	__Node__000000_Test_1_2_without_checkbox.IsChecked = false
	__Node__000000_Test_1_2_without_checkbox.IsCheckboxDisabled = false
	__Node__000000_Test_1_2_without_checkbox.CheckboxHasToolTip = false
	__Node__000000_Test_1_2_without_checkbox.CheckboxToolTipText = ``
	__Node__000000_Test_1_2_without_checkbox.HasSecondCheckboxButton = false
	__Node__000000_Test_1_2_without_checkbox.IsSecondCheckboxChecked = false
	__Node__000000_Test_1_2_without_checkbox.IsSecondCheckboxDisabled = false
	__Node__000000_Test_1_2_without_checkbox.TextAfterSecondCheckbox = `Text After Second CheckBox`
	__Node__000000_Test_1_2_without_checkbox.IsInEditMode = false
	__Node__000000_Test_1_2_without_checkbox.IsNodeClickable = false
	__Node__000000_Test_1_2_without_checkbox.IsWithPreceedingIcon = false
	__Node__000000_Test_1_2_without_checkbox.PreceedingIcon = ``

	__Node__000001_root1.Name = `root1`
	__Node__000001_root1.BackgroundColor = ``
	__Node__000001_root1.IsExpanded = true
	__Node__000001_root1.HasCheckboxButton = false
	__Node__000001_root1.IsChecked = false
	__Node__000001_root1.IsCheckboxDisabled = false
	__Node__000001_root1.CheckboxHasToolTip = false
	__Node__000001_root1.CheckboxToolTipText = ``
	__Node__000001_root1.HasSecondCheckboxButton = false
	__Node__000001_root1.IsSecondCheckboxChecked = false
	__Node__000001_root1.IsSecondCheckboxDisabled = false
	__Node__000001_root1.TextAfterSecondCheckbox = ``
	__Node__000001_root1.IsInEditMode = false
	__Node__000001_root1.IsNodeClickable = false
	__Node__000001_root1.IsWithPreceedingIcon = false
	__Node__000001_root1.PreceedingIcon = ``

	__Node__000002_root2.Name = `root2`
	__Node__000002_root2.BackgroundColor = ``
	__Node__000002_root2.IsExpanded = true
	__Node__000002_root2.HasCheckboxButton = false
	__Node__000002_root2.IsChecked = false
	__Node__000002_root2.IsCheckboxDisabled = false
	__Node__000002_root2.CheckboxHasToolTip = false
	__Node__000002_root2.CheckboxToolTipText = ``
	__Node__000002_root2.HasSecondCheckboxButton = false
	__Node__000002_root2.IsSecondCheckboxChecked = false
	__Node__000002_root2.IsSecondCheckboxDisabled = false
	__Node__000002_root2.TextAfterSecondCheckbox = ``
	__Node__000002_root2.IsInEditMode = false
	__Node__000002_root2.IsNodeClickable = false
	__Node__000002_root2.IsWithPreceedingIcon = false
	__Node__000002_root2.PreceedingIcon = ``

	__Node__000003_root3.Name = `root3`
	__Node__000003_root3.BackgroundColor = ``
	__Node__000003_root3.IsExpanded = true
	__Node__000003_root3.HasCheckboxButton = false
	__Node__000003_root3.IsChecked = false
	__Node__000003_root3.IsCheckboxDisabled = false
	__Node__000003_root3.CheckboxHasToolTip = false
	__Node__000003_root3.CheckboxToolTipText = ``
	__Node__000003_root3.HasSecondCheckboxButton = false
	__Node__000003_root3.IsSecondCheckboxChecked = false
	__Node__000003_root3.IsSecondCheckboxDisabled = false
	__Node__000003_root3.TextAfterSecondCheckbox = ``
	__Node__000003_root3.IsInEditMode = false
	__Node__000003_root3.IsNodeClickable = false
	__Node__000003_root3.IsWithPreceedingIcon = false
	__Node__000003_root3.PreceedingIcon = ``

	__Node__000004_root3_test3.Name = `root3.test3`
	__Node__000004_root3_test3.BackgroundColor = ``
	__Node__000004_root3_test3.IsExpanded = false
	__Node__000004_root3_test3.HasCheckboxButton = false
	__Node__000004_root3_test3.IsChecked = false
	__Node__000004_root3_test3.IsCheckboxDisabled = false
	__Node__000004_root3_test3.CheckboxHasToolTip = false
	__Node__000004_root3_test3.CheckboxToolTipText = ``
	__Node__000004_root3_test3.HasSecondCheckboxButton = false
	__Node__000004_root3_test3.IsSecondCheckboxChecked = false
	__Node__000004_root3_test3.IsSecondCheckboxDisabled = false
	__Node__000004_root3_test3.TextAfterSecondCheckbox = ``
	__Node__000004_root3_test3.IsInEditMode = false
	__Node__000004_root3_test3.IsNodeClickable = false
	__Node__000004_root3_test3.IsWithPreceedingIcon = false
	__Node__000004_root3_test3.PreceedingIcon = ``

	__Node__000005_test_1_1.Name = `test 1.1`
	__Node__000005_test_1_1.BackgroundColor = ``
	__Node__000005_test_1_1.IsExpanded = false
	__Node__000005_test_1_1.HasCheckboxButton = true
	__Node__000005_test_1_1.IsChecked = true
	__Node__000005_test_1_1.IsCheckboxDisabled = true
	__Node__000005_test_1_1.CheckboxHasToolTip = false
	__Node__000005_test_1_1.CheckboxToolTipText = ``
	__Node__000005_test_1_1.HasSecondCheckboxButton = false
	__Node__000005_test_1_1.IsSecondCheckboxChecked = false
	__Node__000005_test_1_1.IsSecondCheckboxDisabled = false
	__Node__000005_test_1_1.TextAfterSecondCheckbox = ``
	__Node__000005_test_1_1.IsInEditMode = false
	__Node__000005_test_1_1.IsNodeClickable = false
	__Node__000005_test_1_1.IsWithPreceedingIcon = false
	__Node__000005_test_1_1.PreceedingIcon = ``

	__Node__000006_test_1_2_clickable_.Name = `test 1.2 (clickable)`
	__Node__000006_test_1_2_clickable_.BackgroundColor = ``
	__Node__000006_test_1_2_clickable_.IsExpanded = false
	__Node__000006_test_1_2_clickable_.HasCheckboxButton = false
	__Node__000006_test_1_2_clickable_.IsChecked = false
	__Node__000006_test_1_2_clickable_.IsCheckboxDisabled = false
	__Node__000006_test_1_2_clickable_.CheckboxHasToolTip = false
	__Node__000006_test_1_2_clickable_.CheckboxToolTipText = ``
	__Node__000006_test_1_2_clickable_.HasSecondCheckboxButton = false
	__Node__000006_test_1_2_clickable_.IsSecondCheckboxChecked = false
	__Node__000006_test_1_2_clickable_.IsSecondCheckboxDisabled = false
	__Node__000006_test_1_2_clickable_.TextAfterSecondCheckbox = ``
	__Node__000006_test_1_2_clickable_.IsInEditMode = false
	__Node__000006_test_1_2_clickable_.IsNodeClickable = true
	__Node__000006_test_1_2_clickable_.IsWithPreceedingIcon = true
	__Node__000006_test_1_2_clickable_.PreceedingIcon = `directions_walk`

	__Node__000007_test_1_3.Name = `test 1.3`
	__Node__000007_test_1_3.FontStyle = models.ITALIC
	__Node__000007_test_1_3.BackgroundColor = `yellow`
	__Node__000007_test_1_3.IsExpanded = true
	__Node__000007_test_1_3.HasCheckboxButton = false
	__Node__000007_test_1_3.IsChecked = true
	__Node__000007_test_1_3.IsCheckboxDisabled = false
	__Node__000007_test_1_3.CheckboxHasToolTip = false
	__Node__000007_test_1_3.CheckboxToolTipText = ``
	__Node__000007_test_1_3.HasSecondCheckboxButton = false
	__Node__000007_test_1_3.IsSecondCheckboxChecked = false
	__Node__000007_test_1_3.IsSecondCheckboxDisabled = false
	__Node__000007_test_1_3.TextAfterSecondCheckbox = ``
	__Node__000007_test_1_3.IsInEditMode = false
	__Node__000007_test_1_3.IsNodeClickable = false
	__Node__000007_test_1_3.IsWithPreceedingIcon = false
	__Node__000007_test_1_3.PreceedingIcon = ``

	__Node__000008_test_1_3_1_F.Name = `test 1.3.1..F`
	__Node__000008_test_1_3_1_F.FontStyle = models.ITALIC
	__Node__000008_test_1_3_1_F.BackgroundColor = ``
	__Node__000008_test_1_3_1_F.IsExpanded = false
	__Node__000008_test_1_3_1_F.HasCheckboxButton = false
	__Node__000008_test_1_3_1_F.IsChecked = false
	__Node__000008_test_1_3_1_F.IsCheckboxDisabled = false
	__Node__000008_test_1_3_1_F.CheckboxHasToolTip = false
	__Node__000008_test_1_3_1_F.CheckboxToolTipText = ``
	__Node__000008_test_1_3_1_F.HasSecondCheckboxButton = false
	__Node__000008_test_1_3_1_F.IsSecondCheckboxChecked = false
	__Node__000008_test_1_3_1_F.IsSecondCheckboxDisabled = false
	__Node__000008_test_1_3_1_F.TextAfterSecondCheckbox = ``
	__Node__000008_test_1_3_1_F.IsInEditMode = false
	__Node__000008_test_1_3_1_F.IsNodeClickable = false
	__Node__000008_test_1_3_1_F.IsWithPreceedingIcon = true
	__Node__000008_test_1_3_1_F.PreceedingIcon = `query_stats`

	__Node__000009_test_1_3_1_svg_icon.Name = `test 1.3.1 / svg icon`
	__Node__000009_test_1_3_1_svg_icon.BackgroundColor = ``
	__Node__000009_test_1_3_1_svg_icon.IsExpanded = false
	__Node__000009_test_1_3_1_svg_icon.HasCheckboxButton = false
	__Node__000009_test_1_3_1_svg_icon.IsChecked = false
	__Node__000009_test_1_3_1_svg_icon.IsCheckboxDisabled = false
	__Node__000009_test_1_3_1_svg_icon.CheckboxHasToolTip = false
	__Node__000009_test_1_3_1_svg_icon.CheckboxToolTipText = ``
	__Node__000009_test_1_3_1_svg_icon.HasSecondCheckboxButton = false
	__Node__000009_test_1_3_1_svg_icon.IsSecondCheckboxChecked = false
	__Node__000009_test_1_3_1_svg_icon.IsSecondCheckboxDisabled = false
	__Node__000009_test_1_3_1_svg_icon.TextAfterSecondCheckbox = ``
	__Node__000009_test_1_3_1_svg_icon.IsInEditMode = false
	__Node__000009_test_1_3_1_svg_icon.IsNodeClickable = false
	__Node__000009_test_1_3_1_svg_icon.IsWithPreceedingIcon = false
	__Node__000009_test_1_3_1_svg_icon.PreceedingIcon = ``

	__Node__000010_test_1_4.Name = `test 1.4`
	__Node__000010_test_1_4.BackgroundColor = `#4ece92`
	__Node__000010_test_1_4.IsExpanded = false
	__Node__000010_test_1_4.HasCheckboxButton = false
	__Node__000010_test_1_4.IsChecked = false
	__Node__000010_test_1_4.IsCheckboxDisabled = false
	__Node__000010_test_1_4.CheckboxHasToolTip = false
	__Node__000010_test_1_4.CheckboxToolTipText = ``
	__Node__000010_test_1_4.HasSecondCheckboxButton = false
	__Node__000010_test_1_4.IsSecondCheckboxChecked = false
	__Node__000010_test_1_4.IsSecondCheckboxDisabled = false
	__Node__000010_test_1_4.TextAfterSecondCheckbox = ``
	__Node__000010_test_1_4.IsInEditMode = false
	__Node__000010_test_1_4.IsNodeClickable = false
	__Node__000010_test_1_4.IsWithPreceedingIcon = false
	__Node__000010_test_1_4.PreceedingIcon = ``

	__Node__000011_test_2_1.Name = `test 2.1`
	__Node__000011_test_2_1.BackgroundColor = ``
	__Node__000011_test_2_1.IsExpanded = false
	__Node__000011_test_2_1.HasCheckboxButton = true
	__Node__000011_test_2_1.IsChecked = true
	__Node__000011_test_2_1.IsCheckboxDisabled = false
	__Node__000011_test_2_1.CheckboxHasToolTip = false
	__Node__000011_test_2_1.CheckboxToolTipText = ``
	__Node__000011_test_2_1.HasSecondCheckboxButton = true
	__Node__000011_test_2_1.IsSecondCheckboxChecked = false
	__Node__000011_test_2_1.IsSecondCheckboxDisabled = false
	__Node__000011_test_2_1.TextAfterSecondCheckbox = `Text After Second Checkbox`
	__Node__000011_test_2_1.IsInEditMode = false
	__Node__000011_test_2_1.IsNodeClickable = false
	__Node__000011_test_2_1.IsWithPreceedingIcon = false
	__Node__000011_test_2_1.PreceedingIcon = ``

	__Node__000012_test3.Name = `test3`
	__Node__000012_test3.BackgroundColor = `#F8F8F8`
	__Node__000012_test3.IsExpanded = false
	__Node__000012_test3.HasCheckboxButton = true
	__Node__000012_test3.IsChecked = true
	__Node__000012_test3.IsCheckboxDisabled = false
	__Node__000012_test3.CheckboxHasToolTip = true
	__Node__000012_test3.CheckboxToolTipText = `If you check, this will happen`
	__Node__000012_test3.CheckboxToolTipPosition = models.Above
	__Node__000012_test3.HasSecondCheckboxButton = false
	__Node__000012_test3.IsSecondCheckboxChecked = false
	__Node__000012_test3.IsSecondCheckboxDisabled = false
	__Node__000012_test3.TextAfterSecondCheckbox = ``
	__Node__000012_test3.IsInEditMode = false
	__Node__000012_test3.IsNodeClickable = false
	__Node__000012_test3.IsWithPreceedingIcon = false
	__Node__000012_test3.PreceedingIcon = ``

	__SVGIcon__000000_empty.Name = `empty`
	__SVGIcon__000000_empty.SVG = `<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 -960 960 960" width="24">  </svg>`

	__SVGIcon__000001_sync_alt_rotated_90.Name = `sync_alt_rotated_90`
	__SVGIcon__000001_sync_alt_rotated_90.SVG = `<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 -960 960 960" width="24">   <path d="M280-120 80-320l200-200 57 56-104 104h607v80H233l104 104-57 56Zm400-320-57-56 104-104H120v-80h607L623-784l57-56 200 200-200 200Z" transform="rotate(90, 480, -480)"/> </svg>`

	__Tree__000000_test.Name = `test`

	// Setup of pointers
	// setup of Button instances pointers
	__Button__000000_Button_For_Empty_Space.SVGIcon = __SVGIcon__000001_sync_alt_rotated_90
	__Button__000001_PreceedingSVGIcon.SVGIcon = __SVGIcon__000001_sync_alt_rotated_90
	// setup of Node instances pointers
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000005_test_1_1)
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000006_test_1_2_clickable_)
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000007_test_1_3)
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000010_test_1_4)
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000012_test3)
	__Node__000001_root1.Buttons = append(__Node__000001_root1.Buttons, __Button__000006_key)
	__Node__000001_root1.Buttons = append(__Node__000001_root1.Buttons, __Button__000005_dynamic_form)
	__Node__000002_root2.Children = append(__Node__000002_root2.Children, __Node__000000_Test_1_2_without_checkbox)
	__Node__000002_root2.Children = append(__Node__000002_root2.Children, __Node__000011_test_2_1)
	__Node__000002_root2.Buttons = append(__Node__000002_root2.Buttons, __Button__000008_root_2_add)
	__Node__000003_root3.Children = append(__Node__000003_root3.Children, __Node__000004_root3_test3)
	__Node__000005_test_1_1.Buttons = append(__Node__000005_test_1_1.Buttons, __Button__000007_logout)
	__Node__000005_test_1_1.Buttons = append(__Node__000005_test_1_1.Buttons, __Button__000004_dataset)
	__Node__000006_test_1_2_clickable_.Buttons = append(__Node__000006_test_1_2_clickable_.Buttons, __Button__000006_key)
	__Node__000007_test_1_3.Children = append(__Node__000007_test_1_3.Children, __Node__000009_test_1_3_1_svg_icon)
	__Node__000007_test_1_3.Children = append(__Node__000007_test_1_3.Children, __Node__000008_test_1_3_1_F)
	__Node__000009_test_1_3_1_svg_icon.PreceedingSVGIcon = __SVGIcon__000001_sync_alt_rotated_90
	__Node__000009_test_1_3_1_svg_icon.Buttons = append(__Node__000009_test_1_3_1_svg_icon.Buttons, __Button__000001_PreceedingSVGIcon)
	__Node__000011_test_2_1.Buttons = append(__Node__000011_test_2_1.Buttons, __Button__000000_Button_For_Empty_Space)
	__Node__000012_test3.Buttons = append(__Node__000012_test3.Buttons, __Button__000010_test)
	__Node__000012_test3.Buttons = append(__Node__000012_test3.Buttons, __Button__000009_settings)
	__Node__000012_test3.Buttons = append(__Node__000012_test3.Buttons, __Button__000003_arrow_circle_left)
	// setup of SVGIcon instances pointers
	// setup of Tree instances pointers
	__Tree__000000_test.RootNodes = append(__Tree__000000_test.RootNodes, __Node__000001_root1)
	__Tree__000000_test.RootNodes = append(__Tree__000000_test.RootNodes, __Node__000002_root2)
	__Tree__000000_test.RootNodes = append(__Tree__000000_test.RootNodes, __Node__000003_root3)
}
