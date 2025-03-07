package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"

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
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__AsSplit__000000_planel_root := (&models.AsSplit{}).Stage(stage)

	__AsSplitArea__000000_top := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000001_bottom := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000002_extra := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000003_sidebar_tree := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000004_table := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000005_form := (&models.AsSplitArea{}).Stage(stage)

	__Form__000000_Form_for_Panel_Form := (&models.Form{}).Stage(stage)

	__Table__000000_Table_within_panel := (&models.Table{}).Stage(stage)

	__Tree__000000_Tree := (&models.Tree{}).Stage(stage)

	__View__000000_a := (&models.View{}).Stage(stage)
	__View__000001_b := (&models.View{}).Stage(stage)
	__View__000002_c := (&models.View{}).Stage(stage)

	// Setup of values

	__AsSplit__000000_planel_root.Name = `planel root`
	__AsSplit__000000_planel_root.Direction = models.Vertical

	__AsSplitArea__000000_top.Name = `top`
	__AsSplitArea__000000_top.Size = 100.000000
	__AsSplitArea__000000_top.IsAny = false

	__AsSplitArea__000001_bottom.Name = `bottom`
	__AsSplitArea__000001_bottom.Size = 150.000000
	__AsSplitArea__000001_bottom.IsAny = false

	__AsSplitArea__000002_extra.Name = `extra`
	__AsSplitArea__000002_extra.Size = 0.000000
	__AsSplitArea__000002_extra.IsAny = true

	__AsSplitArea__000003_sidebar_tree.Name = `sidebar tree`
	__AsSplitArea__000003_sidebar_tree.Size = 20.000000
	__AsSplitArea__000003_sidebar_tree.IsAny = false

	__AsSplitArea__000004_table.Name = `table`
	__AsSplitArea__000004_table.Size = 50.000000
	__AsSplitArea__000004_table.IsAny = false

	__AsSplitArea__000005_form.Name = `form`
	__AsSplitArea__000005_form.Size = 30.000000
	__AsSplitArea__000005_form.IsAny = false

	__Form__000000_Form_for_Panel_Form.Name = `Form for Panel Form`
	__Form__000000_Form_for_Panel_Form.StackName = `split-form`
	__Form__000000_Form_for_Panel_Form.FormName = `Form`

	__Table__000000_Table_within_panel.Name = `Table within panel`
	__Table__000000_Table_within_panel.StackName = `split-table`
	__Table__000000_Table_within_panel.TableName = `Table`

	__Tree__000000_Tree.Name = `Tree`
	__Tree__000000_Tree.StackName = `split-sidebar`
	__Tree__000000_Tree.TreeName = `gong`

	__View__000000_a.Name = `a`

	__View__000001_b.Name = `b`

	__View__000002_c.Name = `c`

	// Setup of pointers
	// setup of AsSplit instances pointers
	__AsSplit__000000_planel_root.AsSplitAreas = append(__AsSplit__000000_planel_root.AsSplitAreas, __AsSplitArea__000003_sidebar_tree)
	__AsSplit__000000_planel_root.AsSplitAreas = append(__AsSplit__000000_planel_root.AsSplitAreas, __AsSplitArea__000004_table)
	__AsSplit__000000_planel_root.AsSplitAreas = append(__AsSplit__000000_planel_root.AsSplitAreas, __AsSplitArea__000005_form)
	// setup of AsSplitArea instances pointers
	__AsSplitArea__000002_extra.AsSplits = append(__AsSplitArea__000002_extra.AsSplits, __AsSplit__000000_planel_root)
	__AsSplitArea__000003_sidebar_tree.Tree = __Tree__000000_Tree
	__AsSplitArea__000004_table.Table = __Table__000000_Table_within_panel
	__AsSplitArea__000005_form.Form = __Form__000000_Form_for_Panel_Form
	// setup of Form instances pointers
	// setup of Table instances pointers
	// setup of Tree instances pointers
	// setup of View instances pointers
	__View__000000_a.RootAsSplitAreas = append(__View__000000_a.RootAsSplitAreas, __AsSplitArea__000000_top)
	__View__000000_a.RootAsSplitAreas = append(__View__000000_a.RootAsSplitAreas, __AsSplitArea__000001_bottom)
	__View__000000_a.RootAsSplitAreas = append(__View__000000_a.RootAsSplitAreas, __AsSplitArea__000002_extra)
}
