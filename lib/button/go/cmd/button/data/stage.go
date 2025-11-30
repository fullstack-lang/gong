package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/button/go/models"
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

	const __write__local_time = "2025-11-30 14:51:33.987428 CET"
	const __write__utc_time__ = "2025-11-30 13:51:33.987428 UTC"

	const __commitId__ = "0000000140"

	// Declaration of instances to stage

	__Button__000000_One := (&models.Button{}).Stage(stage)
	__Button__000001_Two := (&models.Button{}).Stage(stage)
	__Button__000002_Three := (&models.Button{}).Stage(stage)
	__Button__000003_Four := (&models.Button{}).Stage(stage)
	__Button__000004_Five := (&models.Button{}).Stage(stage)
	__Button__000005_Un := (&models.Button{}).Stage(stage)

	__ButtonToggle__000000_small := (&models.ButtonToggle{}).Stage(stage)
	__ButtonToggle__000001_medium := (&models.ButtonToggle{}).Stage(stage)
	__ButtonToggle__000002_large := (&models.ButtonToggle{}).Stage(stage)
	__ButtonToggle__000003_flour := (&models.ButtonToggle{}).Stage(stage)
	__ButtonToggle__000004_eggs := (&models.ButtonToggle{}).Stage(stage)
	__ButtonToggle__000005_sugar := (&models.ButtonToggle{}).Stage(stage)

	__Group__000000_Group_1 := (&models.Group{}).Stage(stage)
	__Group__000001_Groupe_2 := (&models.Group{}).Stage(stage)

	__GroupToogle__000000_Group_toggle_single_selector := (&models.GroupToogle{}).Stage(stage)
	__GroupToogle__000001_Group_toggle_multi_selector := (&models.GroupToogle{}).Stage(stage)

	__Layout__000000_layout := (&models.Layout{}).Stage(stage)

	// Setup of values

	__Button__000000_One.Name = `One`
	__Button__000000_One.Label = `One`
	__Button__000000_One.Icon = `draw`
	__Button__000000_One.IsDisabled = false
	__Button__000000_One.Color = models.MatButtonPaletteTypePrimary
	__Button__000000_One.MatButtonType = models.MatButtonTypeBasic
	__Button__000000_One.MatButtonAppearance = models.MatButtonAppearanceElevated

	__Button__000001_Two.Name = `Two`
	__Button__000001_Two.Label = `Two`
	__Button__000001_Two.Icon = `add`
	__Button__000001_Two.IsDisabled = false
	__Button__000001_Two.Color = models.MatButtonPaletteTypeWarn
	__Button__000001_Two.MatButtonType = models.MatButtonTypeExtendedFab
	__Button__000001_Two.MatButtonAppearance = models.MatButtonAppearanceFilled

	__Button__000002_Three.Name = `Three`
	__Button__000002_Three.Label = `Three`
	__Button__000002_Three.Icon = `add_box`
	__Button__000002_Three.IsDisabled = false
	__Button__000002_Three.Color = models.MatButtonPaletteTypeAccent
	__Button__000002_Three.MatButtonType = models.MatButtonTypeBasic
	__Button__000002_Three.MatButtonAppearance = models.MatButtonAppearanceFilled

	__Button__000003_Four.Name = `Four`
	__Button__000003_Four.Label = `Four`
	__Button__000003_Four.Icon = ``
	__Button__000003_Four.IsDisabled = false
	__Button__000003_Four.Color = models.MatButtonPaletteTypePrimary
	__Button__000003_Four.MatButtonType = models.MatButtonTypeExtendedFab
	__Button__000003_Four.MatButtonAppearance = models.MatButtonAppearanceElevated

	__Button__000004_Five.Name = `Five`
	__Button__000004_Five.Label = `Five`
	__Button__000004_Five.Icon = ``
	__Button__000004_Five.IsDisabled = false

	__Button__000005_Un.Name = `Un`
	__Button__000005_Un.Label = `Un`
	__Button__000005_Un.Icon = `add`
	__Button__000005_Un.IsDisabled = false
	__Button__000005_Un.MatButtonType = models.MatButtonTypeMiniFab
	__Button__000005_Un.MatButtonAppearance = models.MatButtonAppearanceOutlined

	__ButtonToggle__000000_small.Name = `small`
	__ButtonToggle__000000_small.Label = `small`
	__ButtonToggle__000000_small.Icon = ``
	__ButtonToggle__000000_small.IsDisabled = false
	__ButtonToggle__000000_small.IsChecked = false

	__ButtonToggle__000001_medium.Name = `medium`
	__ButtonToggle__000001_medium.Label = `medium`
	__ButtonToggle__000001_medium.Icon = ``
	__ButtonToggle__000001_medium.IsDisabled = false
	__ButtonToggle__000001_medium.IsChecked = false

	__ButtonToggle__000002_large.Name = `large`
	__ButtonToggle__000002_large.Label = `large`
	__ButtonToggle__000002_large.Icon = ``
	__ButtonToggle__000002_large.IsDisabled = false
	__ButtonToggle__000002_large.IsChecked = false

	__ButtonToggle__000003_flour.Name = `flour`
	__ButtonToggle__000003_flour.Label = `flour`
	__ButtonToggle__000003_flour.Icon = ``
	__ButtonToggle__000003_flour.IsDisabled = false
	__ButtonToggle__000003_flour.IsChecked = false

	__ButtonToggle__000004_eggs.Name = `eggs`
	__ButtonToggle__000004_eggs.Label = `eggs`
	__ButtonToggle__000004_eggs.Icon = ``
	__ButtonToggle__000004_eggs.IsDisabled = false
	__ButtonToggle__000004_eggs.IsChecked = false

	__ButtonToggle__000005_sugar.Name = `sugar`
	__ButtonToggle__000005_sugar.Label = `sugar`
	__ButtonToggle__000005_sugar.Icon = ``
	__ButtonToggle__000005_sugar.IsDisabled = false
	__ButtonToggle__000005_sugar.IsChecked = false

	__Group__000000_Group_1.Name = `Group 1`
	__Group__000000_Group_1.Percentage = 50.000000
	__Group__000000_Group_1.NbColumns = 15

	__Group__000001_Groupe_2.Name = `Groupe 2`
	__Group__000001_Groupe_2.Percentage = 50.000000
	__Group__000001_Groupe_2.NbColumns = 12

	__GroupToogle__000000_Group_toggle_single_selector.Name = `Group toggle single selector`
	__GroupToogle__000000_Group_toggle_single_selector.Percentage = 40.000000
	__GroupToogle__000000_Group_toggle_single_selector.IsSingleSelector = false

	__GroupToogle__000001_Group_toggle_multi_selector.Name = `Group toggle multi selector`
	__GroupToogle__000001_Group_toggle_multi_selector.Percentage = 50.000000
	__GroupToogle__000001_Group_toggle_multi_selector.IsSingleSelector = false

	__Layout__000000_layout.Name = `layout`

	// Setup of pointers
	// setup of Button instances pointers
	// setup of ButtonToggle instances pointers
	// setup of Group instances pointers
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000000_One)
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000001_Two)
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000002_Three)
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000003_Four)
	__Group__000000_Group_1.Buttons = append(__Group__000000_Group_1.Buttons, __Button__000004_Five)
	__Group__000001_Groupe_2.Buttons = append(__Group__000001_Groupe_2.Buttons, __Button__000005_Un)
	// setup of GroupToogle instances pointers
	__GroupToogle__000000_Group_toggle_single_selector.ButtonToggles = append(__GroupToogle__000000_Group_toggle_single_selector.ButtonToggles, __ButtonToggle__000000_small)
	__GroupToogle__000000_Group_toggle_single_selector.ButtonToggles = append(__GroupToogle__000000_Group_toggle_single_selector.ButtonToggles, __ButtonToggle__000001_medium)
	__GroupToogle__000000_Group_toggle_single_selector.ButtonToggles = append(__GroupToogle__000000_Group_toggle_single_selector.ButtonToggles, __ButtonToggle__000002_large)
	__GroupToogle__000001_Group_toggle_multi_selector.ButtonToggles = append(__GroupToogle__000001_Group_toggle_multi_selector.ButtonToggles, __ButtonToggle__000003_flour)
	__GroupToogle__000001_Group_toggle_multi_selector.ButtonToggles = append(__GroupToogle__000001_Group_toggle_multi_selector.ButtonToggles, __ButtonToggle__000004_eggs)
	__GroupToogle__000001_Group_toggle_multi_selector.ButtonToggles = append(__GroupToogle__000001_Group_toggle_multi_selector.ButtonToggles, __ButtonToggle__000005_sugar)
	// setup of Layout instances pointers
	__Layout__000000_layout.Groups = append(__Layout__000000_layout.Groups, __Group__000000_Group_1)
	__Layout__000000_layout.Groups = append(__Layout__000000_layout.Groups, __Group__000001_Groupe_2)
	__Layout__000000_layout.GroupToogles = append(__Layout__000000_layout.GroupToogles, __GroupToogle__000000_Group_toggle_single_selector)
	__Layout__000000_layout.GroupToogles = append(__Layout__000000_layout.GroupToogles, __GroupToogle__000001_Group_toggle_multi_selector)
}

