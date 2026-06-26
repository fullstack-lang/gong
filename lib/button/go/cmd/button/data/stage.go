package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
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

	__Button__00000000_ := (&models.Button{Name: `One`}).Stage(stage)
	__Button__00000001_ := (&models.Button{Name: `Two`}).Stage(stage)
	__Button__00000002_ := (&models.Button{Name: `Three`}).Stage(stage)
	__Button__00000003_ := (&models.Button{Name: `Four`}).Stage(stage)
	__Button__00000004_ := (&models.Button{Name: `Five`}).Stage(stage)
	__Button__00000005_ := (&models.Button{Name: `Un`}).Stage(stage)

	__ButtonToggle__00000000_ := (&models.ButtonToggle{Name: `small`}).Stage(stage)
	__ButtonToggle__00000001_ := (&models.ButtonToggle{Name: `medium`}).Stage(stage)
	__ButtonToggle__00000002_ := (&models.ButtonToggle{Name: `large`}).Stage(stage)
	__ButtonToggle__00000003_ := (&models.ButtonToggle{Name: `flour`}).Stage(stage)
	__ButtonToggle__00000004_ := (&models.ButtonToggle{Name: `eggs`}).Stage(stage)
	__ButtonToggle__00000005_ := (&models.ButtonToggle{Name: `sugar`}).Stage(stage)

	__Group__00000000_ := (&models.Group{Name: `Group 1`}).Stage(stage)
	__Group__00000001_ := (&models.Group{Name: `Groupe 2`}).Stage(stage)

	__GroupToogle__00000000_ := (&models.GroupToogle{Name: `Group toggle single selector`}).Stage(stage)
	__GroupToogle__00000001_ := (&models.GroupToogle{Name: `Group toggle multi selector`}).Stage(stage)

	__Layout__00000000_ := (&models.Layout{Name: `layout`}).Stage(stage)

	// insertion point for initialization of values

	__Button__00000000_.Name = `One`
	__Button__00000000_.Label = `One`
	__Button__00000000_.Icon = `draw`
	__Button__00000000_.IsDisabled = false
	__Button__00000000_.Color = models.MatButtonPaletteTypePrimary
	__Button__00000000_.MatButtonType = models.MatButtonTypeBasic
	__Button__00000000_.MatButtonAppearance = models.MatButtonAppearanceElevated
	__Button__00000000_.HasToolTip = true
	__Button__00000000_.ToolTipText = `This is a tool tip`
	__Button__00000000_.ToolTipPosition = models.Below

	__Button__00000001_.Name = `Two`
	__Button__00000001_.Label = `Two`
	__Button__00000001_.Icon = `add`
	__Button__00000001_.IsDisabled = false
	__Button__00000001_.Color = models.MatButtonPaletteTypeWarn
	__Button__00000001_.MatButtonType = models.MatButtonTypeExtendedFab
	__Button__00000001_.MatButtonAppearance = models.MatButtonAppearanceFilled
	__Button__00000001_.HasToolTip = false
	__Button__00000001_.ToolTipText = ``
	__Button__00000001_.ToolTipPosition = ""

	__Button__00000002_.Name = `Three`
	__Button__00000002_.Label = `Three`
	__Button__00000002_.Icon = `add_box`
	__Button__00000002_.IsDisabled = false
	__Button__00000002_.Color = models.MatButtonPaletteTypeAccent
	__Button__00000002_.MatButtonType = models.MatButtonTypeBasic
	__Button__00000002_.MatButtonAppearance = models.MatButtonAppearanceFilled
	__Button__00000002_.HasToolTip = false
	__Button__00000002_.ToolTipText = ``
	__Button__00000002_.ToolTipPosition = ""

	__Button__00000003_.Name = `Four`
	__Button__00000003_.Label = `Four`
	__Button__00000003_.Icon = ``
	__Button__00000003_.IsDisabled = false
	__Button__00000003_.Color = models.MatButtonPaletteTypePrimary
	__Button__00000003_.MatButtonType = models.MatButtonTypeExtendedFab
	__Button__00000003_.MatButtonAppearance = models.MatButtonAppearanceElevated
	__Button__00000003_.HasToolTip = false
	__Button__00000003_.ToolTipText = ``
	__Button__00000003_.ToolTipPosition = ""

	__Button__00000004_.Name = `Five`
	__Button__00000004_.Label = `Five`
	__Button__00000004_.Icon = ``
	__Button__00000004_.IsDisabled = false
	__Button__00000004_.Color = ""
	__Button__00000004_.MatButtonType = ""
	__Button__00000004_.MatButtonAppearance = ""
	__Button__00000004_.HasToolTip = false
	__Button__00000004_.ToolTipText = ``
	__Button__00000004_.ToolTipPosition = ""

	__Button__00000005_.Name = `Un`
	__Button__00000005_.Label = `Un`
	__Button__00000005_.Icon = `add`
	__Button__00000005_.IsDisabled = false
	__Button__00000005_.Color = ""
	__Button__00000005_.MatButtonType = models.MatButtonTypeMiniFab
	__Button__00000005_.MatButtonAppearance = models.MatButtonAppearanceOutlined
	__Button__00000005_.HasToolTip = false
	__Button__00000005_.ToolTipText = ``
	__Button__00000005_.ToolTipPosition = ""

	__ButtonToggle__00000000_.Name = `small`
	__ButtonToggle__00000000_.Label = `small`
	__ButtonToggle__00000000_.Icon = ``
	__ButtonToggle__00000000_.IsDisabled = false
	__ButtonToggle__00000000_.IsChecked = false

	__ButtonToggle__00000001_.Name = `medium`
	__ButtonToggle__00000001_.Label = `medium`
	__ButtonToggle__00000001_.Icon = ``
	__ButtonToggle__00000001_.IsDisabled = false
	__ButtonToggle__00000001_.IsChecked = false

	__ButtonToggle__00000002_.Name = `large`
	__ButtonToggle__00000002_.Label = `large`
	__ButtonToggle__00000002_.Icon = ``
	__ButtonToggle__00000002_.IsDisabled = false
	__ButtonToggle__00000002_.IsChecked = false

	__ButtonToggle__00000003_.Name = `flour`
	__ButtonToggle__00000003_.Label = `flour`
	__ButtonToggle__00000003_.Icon = ``
	__ButtonToggle__00000003_.IsDisabled = false
	__ButtonToggle__00000003_.IsChecked = false

	__ButtonToggle__00000004_.Name = `eggs`
	__ButtonToggle__00000004_.Label = `eggs`
	__ButtonToggle__00000004_.Icon = ``
	__ButtonToggle__00000004_.IsDisabled = false
	__ButtonToggle__00000004_.IsChecked = false

	__ButtonToggle__00000005_.Name = `sugar`
	__ButtonToggle__00000005_.Label = `sugar`
	__ButtonToggle__00000005_.Icon = ``
	__ButtonToggle__00000005_.IsDisabled = false
	__ButtonToggle__00000005_.IsChecked = false

	__Group__00000000_.Name = `Group 1`
	__Group__00000000_.Percentage = 50.000000
	__Group__00000000_.NbColumns = 15

	__Group__00000001_.Name = `Groupe 2`
	__Group__00000001_.Percentage = 50.000000
	__Group__00000001_.NbColumns = 12

	__GroupToogle__00000000_.Name = `Group toggle single selector`
	__GroupToogle__00000000_.Percentage = 40.000000
	__GroupToogle__00000000_.IsSingleSelector = false

	__GroupToogle__00000001_.Name = `Group toggle multi selector`
	__GroupToogle__00000001_.Percentage = 50.000000
	__GroupToogle__00000001_.IsSingleSelector = false

	__Layout__00000000_.Name = `layout`

	// insertion point for setup of pointers
	__Group__00000000_.Buttons = append(__Group__00000000_.Buttons, __Button__00000000_)
	__Group__00000000_.Buttons = append(__Group__00000000_.Buttons, __Button__00000001_)
	__Group__00000000_.Buttons = append(__Group__00000000_.Buttons, __Button__00000002_)
	__Group__00000000_.Buttons = append(__Group__00000000_.Buttons, __Button__00000003_)
	__Group__00000000_.Buttons = append(__Group__00000000_.Buttons, __Button__00000004_)
	__Group__00000001_.Buttons = append(__Group__00000001_.Buttons, __Button__00000005_)
	__GroupToogle__00000000_.ButtonToggles = append(__GroupToogle__00000000_.ButtonToggles, __ButtonToggle__00000000_)
	__GroupToogle__00000000_.ButtonToggles = append(__GroupToogle__00000000_.ButtonToggles, __ButtonToggle__00000001_)
	__GroupToogle__00000000_.ButtonToggles = append(__GroupToogle__00000000_.ButtonToggles, __ButtonToggle__00000002_)
	__GroupToogle__00000001_.ButtonToggles = append(__GroupToogle__00000001_.ButtonToggles, __ButtonToggle__00000003_)
	__GroupToogle__00000001_.ButtonToggles = append(__GroupToogle__00000001_.ButtonToggles, __ButtonToggle__00000004_)
	__GroupToogle__00000001_.ButtonToggles = append(__GroupToogle__00000001_.ButtonToggles, __ButtonToggle__00000005_)
	__Layout__00000000_.Groups = append(__Layout__00000000_.Groups, __Group__00000000_)
	__Layout__00000000_.Groups = append(__Layout__00000000_.Groups, __Group__00000001_)
	__Layout__00000000_.GroupToogles = append(__Layout__00000000_.GroupToogles, __GroupToogle__00000000_)
	__Layout__00000000_.GroupToogles = append(__Layout__00000000_.GroupToogles, __GroupToogle__00000001_)
}
