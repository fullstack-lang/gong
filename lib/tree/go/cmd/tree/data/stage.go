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

	// insertion point for declaration of instances to stage

	__Button__00000000_ := (&models.Button{Name: `Button For Empty Space`}).Stage(stage)
	__Button__00000001_ := (&models.Button{Name: `PreceedingSVGIcon`}).Stage(stage)
	__Button__00000002_ := (&models.Button{Name: `Test 1.2 add`}).Stage(stage)
	__Button__00000003_ := (&models.Button{Name: `arrow_circle_left`}).Stage(stage)
	__Button__00000004_ := (&models.Button{Name: `dataset`}).Stage(stage)
	__Button__00000005_ := (&models.Button{Name: `dynamic_form`}).Stage(stage)
	__Button__00000006_ := (&models.Button{Name: `key`}).Stage(stage)
	__Button__00000007_ := (&models.Button{Name: `logout`}).Stage(stage)
	__Button__00000008_ := (&models.Button{Name: `root 2 - add`}).Stage(stage)
	__Button__00000009_ := (&models.Button{Name: `settings`}).Stage(stage)
	__Button__00000010_ := (&models.Button{Name: `test`}).Stage(stage)
	__Button__00000011_ := (&models.Button{Name: `test 1.3 plus button`}).Stage(stage)

	__Node__00000000_ := (&models.Node{Name: `Test 1.2 without checkbox`}).Stage(stage)
	__Node__00000001_ := (&models.Node{Name: `root1`}).Stage(stage)
	__Node__00000002_ := (&models.Node{Name: `root2`}).Stage(stage)
	__Node__00000003_ := (&models.Node{Name: `root3`}).Stage(stage)
	__Node__00000004_ := (&models.Node{Name: `root3.test3`}).Stage(stage)
	__Node__00000005_ := (&models.Node{Name: `test 1.1`}).Stage(stage)
	__Node__00000006_ := (&models.Node{Name: `test 1.2 (clickable)`}).Stage(stage)
	__Node__00000007_ := (&models.Node{Name: `test 1.3`}).Stage(stage)
	__Node__00000008_ := (&models.Node{Name: `test 1.3.1..F`}).Stage(stage)
	__Node__00000009_ := (&models.Node{Name: `test 1.3.1 / svg icon`}).Stage(stage)
	__Node__00000010_ := (&models.Node{Name: `test 1.4`}).Stage(stage)
	__Node__00000011_ := (&models.Node{Name: `test 2.1`}).Stage(stage)
	__Node__00000012_ := (&models.Node{Name: `test3`}).Stage(stage)

	__SVGIcon__00000000_ := (&models.SVGIcon{Name: `empty`}).Stage(stage)
	__SVGIcon__00000001_ := (&models.SVGIcon{Name: `sync_alt_rotated_90`}).Stage(stage)

	__Tree__00000000_ := (&models.Tree{Name: `test`}).Stage(stage)

	// insertion point for initialization of values

	__Button__00000000_.Name = `Button For Empty Space`
	__Button__00000000_.Icon = ``
	__Button__00000000_.IsDisabled = false
	__Button__00000000_.HasToolTip = false
	__Button__00000000_.ToolTipText = ``
	__Button__00000000_.ToolTipPosition = ""

	__Button__00000001_.Name = `PreceedingSVGIcon`
	__Button__00000001_.Icon = ``
	__Button__00000001_.IsDisabled = false
	__Button__00000001_.HasToolTip = false
	__Button__00000001_.ToolTipText = ``
	__Button__00000001_.ToolTipPosition = ""

	__Button__00000002_.Name = `Test 1.2 add`
	__Button__00000002_.Icon = `add`
	__Button__00000002_.IsDisabled = false
	__Button__00000002_.HasToolTip = false
	__Button__00000002_.ToolTipText = ``
	__Button__00000002_.ToolTipPosition = ""

	__Button__00000003_.Name = `arrow_circle_left`
	__Button__00000003_.Icon = `arrow_circle_left`
	__Button__00000003_.IsDisabled = true
	__Button__00000003_.HasToolTip = true
	__Button__00000003_.ToolTipText = `This is a tooltip for the button`
	__Button__00000003_.ToolTipPosition = models.Right

	__Button__00000004_.Name = `dataset`
	__Button__00000004_.Icon = `dataset`
	__Button__00000004_.IsDisabled = false
	__Button__00000004_.HasToolTip = false
	__Button__00000004_.ToolTipText = ``
	__Button__00000004_.ToolTipPosition = ""

	__Button__00000005_.Name = `dynamic_form`
	__Button__00000005_.Icon = `dynamic_form`
	__Button__00000005_.IsDisabled = false
	__Button__00000005_.HasToolTip = false
	__Button__00000005_.ToolTipText = ``
	__Button__00000005_.ToolTipPosition = ""

	__Button__00000006_.Name = `key`
	__Button__00000006_.Icon = `key`
	__Button__00000006_.IsDisabled = true
	__Button__00000006_.HasToolTip = false
	__Button__00000006_.ToolTipText = ``
	__Button__00000006_.ToolTipPosition = ""

	__Button__00000007_.Name = `logout`
	__Button__00000007_.Icon = `logout`
	__Button__00000007_.IsDisabled = false
	__Button__00000007_.HasToolTip = false
	__Button__00000007_.ToolTipText = ``
	__Button__00000007_.ToolTipPosition = ""

	__Button__00000008_.Name = `root 2 - add`
	__Button__00000008_.Icon = `add`
	__Button__00000008_.IsDisabled = false
	__Button__00000008_.HasToolTip = false
	__Button__00000008_.ToolTipText = ``
	__Button__00000008_.ToolTipPosition = ""

	__Button__00000009_.Name = `settings`
	__Button__00000009_.Icon = `settings`
	__Button__00000009_.IsDisabled = false
	__Button__00000009_.HasToolTip = false
	__Button__00000009_.ToolTipText = ``
	__Button__00000009_.ToolTipPosition = ""

	__Button__00000010_.Name = `test`
	__Button__00000010_.Icon = `edit`
	__Button__00000010_.IsDisabled = false
	__Button__00000010_.HasToolTip = false
	__Button__00000010_.ToolTipText = ``
	__Button__00000010_.ToolTipPosition = ""

	__Button__00000011_.Name = `test 1.3 plus button`
	__Button__00000011_.Icon = `refresh`
	__Button__00000011_.IsDisabled = false
	__Button__00000011_.HasToolTip = true
	__Button__00000011_.ToolTipText = `refresh`
	__Button__00000011_.ToolTipPosition = models.Right

	__Node__00000000_.Name = `Test 1.2 without checkbox`
	__Node__00000000_.IsWithPrefix = false
	__Node__00000000_.Prefix = ``
	__Node__00000000_.FontStyle = models.NORMAL
	__Node__00000000_.BackgroundColor = ``
	__Node__00000000_.IsExpanded = false
	__Node__00000000_.HasCheckboxButton = false
	__Node__00000000_.IsChecked = false
	__Node__00000000_.IsCheckboxDisabled = false
	__Node__00000000_.CheckboxHasToolTip = false
	__Node__00000000_.CheckboxToolTipText = ``
	__Node__00000000_.CheckboxToolTipPosition = ""
	__Node__00000000_.HasSecondCheckboxButton = false
	__Node__00000000_.IsSecondCheckboxChecked = false
	__Node__00000000_.IsSecondCheckboxDisabled = false
	__Node__00000000_.SecondCheckboxHasToolTip = false
	__Node__00000000_.SecondCheckboxToolTipText = ``
	__Node__00000000_.SecondCheckboxToolTipPosition = ""
	__Node__00000000_.TextAfterSecondCheckbox = `Text After Second CheckBox`
	__Node__00000000_.HasToolTip = false
	__Node__00000000_.ToolTipText = ``
	__Node__00000000_.ToolTipPosition = ""
	__Node__00000000_.IsInEditMode = false
	__Node__00000000_.IsNodeClickable = false
	__Node__00000000_.IsWithPreceedingIcon = false
	__Node__00000000_.PreceedingIcon = ``

	__Node__00000001_.Name = `root1`
	__Node__00000001_.IsWithPrefix = false
	__Node__00000001_.Prefix = ``
	__Node__00000001_.FontStyle = ""
	__Node__00000001_.BackgroundColor = ``
	__Node__00000001_.IsExpanded = true
	__Node__00000001_.HasCheckboxButton = false
	__Node__00000001_.IsChecked = false
	__Node__00000001_.IsCheckboxDisabled = false
	__Node__00000001_.CheckboxHasToolTip = false
	__Node__00000001_.CheckboxToolTipText = ``
	__Node__00000001_.CheckboxToolTipPosition = ""
	__Node__00000001_.HasSecondCheckboxButton = false
	__Node__00000001_.IsSecondCheckboxChecked = false
	__Node__00000001_.IsSecondCheckboxDisabled = false
	__Node__00000001_.SecondCheckboxHasToolTip = false
	__Node__00000001_.SecondCheckboxToolTipText = ``
	__Node__00000001_.SecondCheckboxToolTipPosition = ""
	__Node__00000001_.TextAfterSecondCheckbox = ``
	__Node__00000001_.HasToolTip = false
	__Node__00000001_.ToolTipText = ``
	__Node__00000001_.ToolTipPosition = ""
	__Node__00000001_.IsInEditMode = false
	__Node__00000001_.IsNodeClickable = false
	__Node__00000001_.IsWithPreceedingIcon = false
	__Node__00000001_.PreceedingIcon = ``

	__Node__00000002_.Name = `root2`
	__Node__00000002_.IsWithPrefix = false
	__Node__00000002_.Prefix = ``
	__Node__00000002_.FontStyle = ""
	__Node__00000002_.BackgroundColor = ``
	__Node__00000002_.IsExpanded = true
	__Node__00000002_.HasCheckboxButton = false
	__Node__00000002_.IsChecked = false
	__Node__00000002_.IsCheckboxDisabled = false
	__Node__00000002_.CheckboxHasToolTip = false
	__Node__00000002_.CheckboxToolTipText = ``
	__Node__00000002_.CheckboxToolTipPosition = ""
	__Node__00000002_.HasSecondCheckboxButton = false
	__Node__00000002_.IsSecondCheckboxChecked = false
	__Node__00000002_.IsSecondCheckboxDisabled = false
	__Node__00000002_.SecondCheckboxHasToolTip = false
	__Node__00000002_.SecondCheckboxToolTipText = ``
	__Node__00000002_.SecondCheckboxToolTipPosition = ""
	__Node__00000002_.TextAfterSecondCheckbox = ``
	__Node__00000002_.HasToolTip = true
	__Node__00000002_.ToolTipText = `This node has tool tip`
	__Node__00000002_.ToolTipPosition = models.Above
	__Node__00000002_.IsInEditMode = false
	__Node__00000002_.IsNodeClickable = false
	__Node__00000002_.IsWithPreceedingIcon = false
	__Node__00000002_.PreceedingIcon = ``

	__Node__00000003_.Name = `root3`
	__Node__00000003_.IsWithPrefix = false
	__Node__00000003_.Prefix = ``
	__Node__00000003_.FontStyle = ""
	__Node__00000003_.BackgroundColor = ``
	__Node__00000003_.IsExpanded = true
	__Node__00000003_.HasCheckboxButton = false
	__Node__00000003_.IsChecked = false
	__Node__00000003_.IsCheckboxDisabled = false
	__Node__00000003_.CheckboxHasToolTip = false
	__Node__00000003_.CheckboxToolTipText = ``
	__Node__00000003_.CheckboxToolTipPosition = ""
	__Node__00000003_.HasSecondCheckboxButton = false
	__Node__00000003_.IsSecondCheckboxChecked = false
	__Node__00000003_.IsSecondCheckboxDisabled = false
	__Node__00000003_.SecondCheckboxHasToolTip = false
	__Node__00000003_.SecondCheckboxToolTipText = ``
	__Node__00000003_.SecondCheckboxToolTipPosition = ""
	__Node__00000003_.TextAfterSecondCheckbox = ``
	__Node__00000003_.HasToolTip = false
	__Node__00000003_.ToolTipText = ``
	__Node__00000003_.ToolTipPosition = ""
	__Node__00000003_.IsInEditMode = false
	__Node__00000003_.IsNodeClickable = false
	__Node__00000003_.IsWithPreceedingIcon = false
	__Node__00000003_.PreceedingIcon = ``

	__Node__00000004_.Name = `root3.test3`
	__Node__00000004_.IsWithPrefix = false
	__Node__00000004_.Prefix = ``
	__Node__00000004_.FontStyle = ""
	__Node__00000004_.BackgroundColor = ``
	__Node__00000004_.IsExpanded = false
	__Node__00000004_.HasCheckboxButton = false
	__Node__00000004_.IsChecked = false
	__Node__00000004_.IsCheckboxDisabled = false
	__Node__00000004_.CheckboxHasToolTip = false
	__Node__00000004_.CheckboxToolTipText = ``
	__Node__00000004_.CheckboxToolTipPosition = ""
	__Node__00000004_.HasSecondCheckboxButton = false
	__Node__00000004_.IsSecondCheckboxChecked = false
	__Node__00000004_.IsSecondCheckboxDisabled = false
	__Node__00000004_.SecondCheckboxHasToolTip = false
	__Node__00000004_.SecondCheckboxToolTipText = ``
	__Node__00000004_.SecondCheckboxToolTipPosition = ""
	__Node__00000004_.TextAfterSecondCheckbox = ``
	__Node__00000004_.HasToolTip = false
	__Node__00000004_.ToolTipText = ``
	__Node__00000004_.ToolTipPosition = ""
	__Node__00000004_.IsInEditMode = false
	__Node__00000004_.IsNodeClickable = false
	__Node__00000004_.IsWithPreceedingIcon = false
	__Node__00000004_.PreceedingIcon = ``

	__Node__00000005_.Name = `test 1.1`
	__Node__00000005_.IsWithPrefix = true
	__Node__00000005_.Prefix = `1.1`
	__Node__00000005_.FontStyle = ""
	__Node__00000005_.BackgroundColor = ``
	__Node__00000005_.IsExpanded = false
	__Node__00000005_.HasCheckboxButton = true
	__Node__00000005_.IsChecked = true
	__Node__00000005_.IsCheckboxDisabled = true
	__Node__00000005_.CheckboxHasToolTip = false
	__Node__00000005_.CheckboxToolTipText = ``
	__Node__00000005_.CheckboxToolTipPosition = ""
	__Node__00000005_.HasSecondCheckboxButton = false
	__Node__00000005_.IsSecondCheckboxChecked = false
	__Node__00000005_.IsSecondCheckboxDisabled = false
	__Node__00000005_.SecondCheckboxHasToolTip = false
	__Node__00000005_.SecondCheckboxToolTipText = ``
	__Node__00000005_.SecondCheckboxToolTipPosition = ""
	__Node__00000005_.TextAfterSecondCheckbox = ``
	__Node__00000005_.HasToolTip = false
	__Node__00000005_.ToolTipText = ``
	__Node__00000005_.ToolTipPosition = ""
	__Node__00000005_.IsInEditMode = false
	__Node__00000005_.IsNodeClickable = false
	__Node__00000005_.IsWithPreceedingIcon = false
	__Node__00000005_.PreceedingIcon = ``

	__Node__00000006_.Name = `test 1.2 (clickable)`
	__Node__00000006_.IsWithPrefix = false
	__Node__00000006_.Prefix = ``
	__Node__00000006_.FontStyle = ""
	__Node__00000006_.BackgroundColor = ``
	__Node__00000006_.IsExpanded = false
	__Node__00000006_.HasCheckboxButton = false
	__Node__00000006_.IsChecked = false
	__Node__00000006_.IsCheckboxDisabled = false
	__Node__00000006_.CheckboxHasToolTip = false
	__Node__00000006_.CheckboxToolTipText = ``
	__Node__00000006_.CheckboxToolTipPosition = ""
	__Node__00000006_.HasSecondCheckboxButton = false
	__Node__00000006_.IsSecondCheckboxChecked = false
	__Node__00000006_.IsSecondCheckboxDisabled = false
	__Node__00000006_.SecondCheckboxHasToolTip = false
	__Node__00000006_.SecondCheckboxToolTipText = ``
	__Node__00000006_.SecondCheckboxToolTipPosition = ""
	__Node__00000006_.TextAfterSecondCheckbox = ``
	__Node__00000006_.HasToolTip = false
	__Node__00000006_.ToolTipText = ``
	__Node__00000006_.ToolTipPosition = ""
	__Node__00000006_.IsInEditMode = false
	__Node__00000006_.IsNodeClickable = true
	__Node__00000006_.IsWithPreceedingIcon = true
	__Node__00000006_.PreceedingIcon = `directions_walk`

	__Node__00000007_.Name = `test 1.3`
	__Node__00000007_.IsWithPrefix = false
	__Node__00000007_.Prefix = ``
	__Node__00000007_.FontStyle = models.ITALIC
	__Node__00000007_.BackgroundColor = `yellow`
	__Node__00000007_.IsExpanded = true
	__Node__00000007_.HasCheckboxButton = false
	__Node__00000007_.IsChecked = true
	__Node__00000007_.IsCheckboxDisabled = false
	__Node__00000007_.CheckboxHasToolTip = false
	__Node__00000007_.CheckboxToolTipText = ``
	__Node__00000007_.CheckboxToolTipPosition = ""
	__Node__00000007_.HasSecondCheckboxButton = false
	__Node__00000007_.IsSecondCheckboxChecked = false
	__Node__00000007_.IsSecondCheckboxDisabled = false
	__Node__00000007_.SecondCheckboxHasToolTip = false
	__Node__00000007_.SecondCheckboxToolTipText = ``
	__Node__00000007_.SecondCheckboxToolTipPosition = ""
	__Node__00000007_.TextAfterSecondCheckbox = ``
	__Node__00000007_.HasToolTip = false
	__Node__00000007_.ToolTipText = ``
	__Node__00000007_.ToolTipPosition = ""
	__Node__00000007_.IsInEditMode = false
	__Node__00000007_.IsNodeClickable = false
	__Node__00000007_.IsWithPreceedingIcon = false
	__Node__00000007_.PreceedingIcon = ``

	__Node__00000008_.Name = `test 1.3.1..F`
	__Node__00000008_.IsWithPrefix = false
	__Node__00000008_.Prefix = ``
	__Node__00000008_.FontStyle = models.ITALIC
	__Node__00000008_.BackgroundColor = ``
	__Node__00000008_.IsExpanded = false
	__Node__00000008_.HasCheckboxButton = false
	__Node__00000008_.IsChecked = false
	__Node__00000008_.IsCheckboxDisabled = false
	__Node__00000008_.CheckboxHasToolTip = false
	__Node__00000008_.CheckboxToolTipText = ``
	__Node__00000008_.CheckboxToolTipPosition = ""
	__Node__00000008_.HasSecondCheckboxButton = false
	__Node__00000008_.IsSecondCheckboxChecked = false
	__Node__00000008_.IsSecondCheckboxDisabled = false
	__Node__00000008_.SecondCheckboxHasToolTip = false
	__Node__00000008_.SecondCheckboxToolTipText = ``
	__Node__00000008_.SecondCheckboxToolTipPosition = ""
	__Node__00000008_.TextAfterSecondCheckbox = ``
	__Node__00000008_.HasToolTip = false
	__Node__00000008_.ToolTipText = ``
	__Node__00000008_.ToolTipPosition = ""
	__Node__00000008_.IsInEditMode = false
	__Node__00000008_.IsNodeClickable = false
	__Node__00000008_.IsWithPreceedingIcon = true
	__Node__00000008_.PreceedingIcon = `query_stats`

	__Node__00000009_.Name = `test 1.3.1 / svg icon`
	__Node__00000009_.IsWithPrefix = false
	__Node__00000009_.Prefix = ``
	__Node__00000009_.FontStyle = ""
	__Node__00000009_.BackgroundColor = ``
	__Node__00000009_.IsExpanded = false
	__Node__00000009_.HasCheckboxButton = false
	__Node__00000009_.IsChecked = false
	__Node__00000009_.IsCheckboxDisabled = false
	__Node__00000009_.CheckboxHasToolTip = false
	__Node__00000009_.CheckboxToolTipText = ``
	__Node__00000009_.CheckboxToolTipPosition = ""
	__Node__00000009_.HasSecondCheckboxButton = false
	__Node__00000009_.IsSecondCheckboxChecked = false
	__Node__00000009_.IsSecondCheckboxDisabled = false
	__Node__00000009_.SecondCheckboxHasToolTip = false
	__Node__00000009_.SecondCheckboxToolTipText = ``
	__Node__00000009_.SecondCheckboxToolTipPosition = ""
	__Node__00000009_.TextAfterSecondCheckbox = ``
	__Node__00000009_.HasToolTip = false
	__Node__00000009_.ToolTipText = ``
	__Node__00000009_.ToolTipPosition = ""
	__Node__00000009_.IsInEditMode = false
	__Node__00000009_.IsNodeClickable = false
	__Node__00000009_.IsWithPreceedingIcon = false
	__Node__00000009_.PreceedingIcon = ``

	__Node__00000010_.Name = `test 1.4`
	__Node__00000010_.IsWithPrefix = false
	__Node__00000010_.Prefix = ``
	__Node__00000010_.FontStyle = ""
	__Node__00000010_.BackgroundColor = `#4ece92`
	__Node__00000010_.IsExpanded = false
	__Node__00000010_.HasCheckboxButton = false
	__Node__00000010_.IsChecked = false
	__Node__00000010_.IsCheckboxDisabled = false
	__Node__00000010_.CheckboxHasToolTip = false
	__Node__00000010_.CheckboxToolTipText = ``
	__Node__00000010_.CheckboxToolTipPosition = ""
	__Node__00000010_.HasSecondCheckboxButton = false
	__Node__00000010_.IsSecondCheckboxChecked = false
	__Node__00000010_.IsSecondCheckboxDisabled = false
	__Node__00000010_.SecondCheckboxHasToolTip = false
	__Node__00000010_.SecondCheckboxToolTipText = ``
	__Node__00000010_.SecondCheckboxToolTipPosition = ""
	__Node__00000010_.TextAfterSecondCheckbox = ``
	__Node__00000010_.HasToolTip = false
	__Node__00000010_.ToolTipText = ``
	__Node__00000010_.ToolTipPosition = ""
	__Node__00000010_.IsInEditMode = false
	__Node__00000010_.IsNodeClickable = true
	__Node__00000010_.IsWithPreceedingIcon = false
	__Node__00000010_.PreceedingIcon = ``

	__Node__00000011_.Name = `test 2.1`
	__Node__00000011_.IsWithPrefix = false
	__Node__00000011_.Prefix = ``
	__Node__00000011_.FontStyle = ""
	__Node__00000011_.BackgroundColor = ``
	__Node__00000011_.IsExpanded = false
	__Node__00000011_.HasCheckboxButton = true
	__Node__00000011_.IsChecked = true
	__Node__00000011_.IsCheckboxDisabled = false
	__Node__00000011_.CheckboxHasToolTip = false
	__Node__00000011_.CheckboxToolTipText = ``
	__Node__00000011_.CheckboxToolTipPosition = ""
	__Node__00000011_.HasSecondCheckboxButton = true
	__Node__00000011_.IsSecondCheckboxChecked = false
	__Node__00000011_.IsSecondCheckboxDisabled = false
	__Node__00000011_.SecondCheckboxHasToolTip = false
	__Node__00000011_.SecondCheckboxToolTipText = ``
	__Node__00000011_.SecondCheckboxToolTipPosition = ""
	__Node__00000011_.TextAfterSecondCheckbox = `Text After Second Checkbox`
	__Node__00000011_.HasToolTip = false
	__Node__00000011_.ToolTipText = ``
	__Node__00000011_.ToolTipPosition = ""
	__Node__00000011_.IsInEditMode = false
	__Node__00000011_.IsNodeClickable = false
	__Node__00000011_.IsWithPreceedingIcon = false
	__Node__00000011_.PreceedingIcon = ``

	__Node__00000012_.Name = `test3`
	__Node__00000012_.IsWithPrefix = false
	__Node__00000012_.Prefix = ``
	__Node__00000012_.FontStyle = ""
	__Node__00000012_.BackgroundColor = `#F8F8F8`
	__Node__00000012_.IsExpanded = false
	__Node__00000012_.HasCheckboxButton = true
	__Node__00000012_.IsChecked = true
	__Node__00000012_.IsCheckboxDisabled = false
	__Node__00000012_.CheckboxHasToolTip = true
	__Node__00000012_.CheckboxToolTipText = `If you check, this will happen`
	__Node__00000012_.CheckboxToolTipPosition = models.Above
	__Node__00000012_.HasSecondCheckboxButton = false
	__Node__00000012_.IsSecondCheckboxChecked = false
	__Node__00000012_.IsSecondCheckboxDisabled = false
	__Node__00000012_.SecondCheckboxHasToolTip = false
	__Node__00000012_.SecondCheckboxToolTipText = ``
	__Node__00000012_.SecondCheckboxToolTipPosition = ""
	__Node__00000012_.TextAfterSecondCheckbox = ``
	__Node__00000012_.HasToolTip = false
	__Node__00000012_.ToolTipText = ``
	__Node__00000012_.ToolTipPosition = ""
	__Node__00000012_.IsInEditMode = false
	__Node__00000012_.IsNodeClickable = false
	__Node__00000012_.IsWithPreceedingIcon = false
	__Node__00000012_.PreceedingIcon = ``

	__SVGIcon__00000000_.Name = `empty`
	__SVGIcon__00000000_.SVG = `<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 -960 960 960" width="24">  </svg>`

	__SVGIcon__00000001_.Name = `sync_alt_rotated_90`
	__SVGIcon__00000001_.SVG = `<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 -960 960 960" width="24">   <path d="M280-120 80-320l200-200 57 56-104 104h607v80H233l104 104-57 56Zm400-320-57-56 104-104H120v-80h607L623-784l57-56 200 200-200 200Z" transform="rotate(90, 480, -480)"/> </svg>`

	__Tree__00000000_.Name = `test`

	// insertion point for setup of pointers
	__Button__00000000_.SVGIcon = __SVGIcon__00000001_
	__Button__00000001_.SVGIcon = __SVGIcon__00000001_
	__Button__00000002_.SVGIcon = nil
	__Button__00000003_.SVGIcon = nil
	__Button__00000004_.SVGIcon = nil
	__Button__00000005_.SVGIcon = nil
	__Button__00000006_.SVGIcon = nil
	__Button__00000007_.SVGIcon = nil
	__Button__00000008_.SVGIcon = nil
	__Button__00000009_.SVGIcon = nil
	__Button__00000010_.SVGIcon = nil
	__Button__00000011_.SVGIcon = nil
	__Node__00000000_.PreceedingSVGIcon = nil
	__Node__00000001_.PreceedingSVGIcon = nil
	__Node__00000001_.Children = append(__Node__00000001_.Children, __Node__00000005_)
	__Node__00000001_.Children = append(__Node__00000001_.Children, __Node__00000006_)
	__Node__00000001_.Children = append(__Node__00000001_.Children, __Node__00000007_)
	__Node__00000001_.Children = append(__Node__00000001_.Children, __Node__00000010_)
	__Node__00000001_.Children = append(__Node__00000001_.Children, __Node__00000012_)
	__Node__00000001_.Buttons = append(__Node__00000001_.Buttons, __Button__00000006_)
	__Node__00000001_.Buttons = append(__Node__00000001_.Buttons, __Button__00000005_)
	__Node__00000002_.PreceedingSVGIcon = nil
	__Node__00000002_.Children = append(__Node__00000002_.Children, __Node__00000000_)
	__Node__00000002_.Children = append(__Node__00000002_.Children, __Node__00000011_)
	__Node__00000002_.Buttons = append(__Node__00000002_.Buttons, __Button__00000008_)
	__Node__00000003_.PreceedingSVGIcon = nil
	__Node__00000003_.Children = append(__Node__00000003_.Children, __Node__00000004_)
	__Node__00000004_.PreceedingSVGIcon = nil
	__Node__00000005_.PreceedingSVGIcon = nil
	__Node__00000005_.Buttons = append(__Node__00000005_.Buttons, __Button__00000007_)
	__Node__00000005_.Buttons = append(__Node__00000005_.Buttons, __Button__00000004_)
	__Node__00000006_.PreceedingSVGIcon = nil
	__Node__00000006_.Buttons = append(__Node__00000006_.Buttons, __Button__00000006_)
	__Node__00000007_.PreceedingSVGIcon = nil
	__Node__00000007_.Children = append(__Node__00000007_.Children, __Node__00000009_)
	__Node__00000007_.Children = append(__Node__00000007_.Children, __Node__00000008_)
	__Node__00000007_.Buttons = append(__Node__00000007_.Buttons, __Button__00000011_)
	__Node__00000008_.PreceedingSVGIcon = nil
	__Node__00000009_.PreceedingSVGIcon = __SVGIcon__00000001_
	__Node__00000009_.Buttons = append(__Node__00000009_.Buttons, __Button__00000001_)
	__Node__00000010_.PreceedingSVGIcon = nil
	__Node__00000010_.Buttons = append(__Node__00000010_.Buttons, __Button__00000003_)
	__Node__00000011_.PreceedingSVGIcon = nil
	__Node__00000011_.Buttons = append(__Node__00000011_.Buttons, __Button__00000000_)
	__Node__00000012_.PreceedingSVGIcon = nil
	__Node__00000012_.Buttons = append(__Node__00000012_.Buttons, __Button__00000010_)
	__Node__00000012_.Buttons = append(__Node__00000012_.Buttons, __Button__00000009_)
	__Tree__00000000_.RootNodes = append(__Tree__00000000_.RootNodes, __Node__00000001_)
	__Tree__00000000_.RootNodes = append(__Tree__00000000_.RootNodes, __Node__00000002_)
	__Tree__00000000_.RootNodes = append(__Tree__00000000_.RootNodes, __Node__00000003_)
}
