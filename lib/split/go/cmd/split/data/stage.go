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

	__AsSplitArea__000000_root_for_tree_table_form := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000001_doc_component := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000002_sidebar_tree := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000003_table := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000004_form := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000005_dynmicaly_load_component := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000006_doc_of_split := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000007_Slider_Area := (&models.AsSplitArea{}).Stage(stage)

	__Doc__000000_doc_in_panel := (&models.Doc{}).Stage(stage)

	__Form__000000_Form_for_Panel_Form := (&models.Form{}).Stage(stage)

	__Slider__000000_Slider := (&models.Slider{}).Stage(stage)

	__Split__000000_doc_of_split_probe := (&models.Split{}).Stage(stage)

	__Svg__000000_svg_component := (&models.Svg{}).Stage(stage)

	__Table__000000_Table_within_panel := (&models.Table{}).Stage(stage)

	__Tree__000000_Tree := (&models.Tree{}).Stage(stage)

	__View__000000_doc_reconstructed := (&models.View{}).Stage(stage)
	__View__000001_slider := (&models.View{}).Stage(stage)
	__View__000002_doc_of_split := (&models.View{}).Stage(stage)

	// Setup of values

	__AsSplit__000000_planel_root.Name = `planel root`
	__AsSplit__000000_planel_root.Direction = models.Horizontal

	__AsSplitArea__000000_root_for_tree_table_form.Name = `root for tree/table/form`
	__AsSplitArea__000000_root_for_tree_table_form.ShowNameInHeader = false
	__AsSplitArea__000000_root_for_tree_table_form.Size = 40.000000
	__AsSplitArea__000000_root_for_tree_table_form.IsAny = false

	__AsSplitArea__000001_doc_component.Name = `doc component`
	__AsSplitArea__000001_doc_component.ShowNameInHeader = false
	__AsSplitArea__000001_doc_component.Size = 30.000000
	__AsSplitArea__000001_doc_component.IsAny = false

	__AsSplitArea__000002_sidebar_tree.Name = `sidebar tree`
	__AsSplitArea__000002_sidebar_tree.ShowNameInHeader = false
	__AsSplitArea__000002_sidebar_tree.Size = 10.000000
	__AsSplitArea__000002_sidebar_tree.IsAny = false

	__AsSplitArea__000003_table.Name = `table`
	__AsSplitArea__000003_table.ShowNameInHeader = false
	__AsSplitArea__000003_table.Size = 50.000000
	__AsSplitArea__000003_table.IsAny = false

	__AsSplitArea__000004_form.Name = `form`
	__AsSplitArea__000004_form.ShowNameInHeader = false
	__AsSplitArea__000004_form.Size = 30.000000
	__AsSplitArea__000004_form.IsAny = false

	__AsSplitArea__000005_dynmicaly_load_component.Name = `dynmicaly load component`
	__AsSplitArea__000005_dynmicaly_load_component.ShowNameInHeader = true
	__AsSplitArea__000005_dynmicaly_load_component.Size = 10.000000
	__AsSplitArea__000005_dynmicaly_load_component.IsAny = false

	__AsSplitArea__000006_doc_of_split.Name = `doc of split`
	__AsSplitArea__000006_doc_of_split.ShowNameInHeader = true
	__AsSplitArea__000006_doc_of_split.Size = 100.000000
	__AsSplitArea__000006_doc_of_split.IsAny = false

	__AsSplitArea__000007_Slider_Area.Name = `Slider Area`
	__AsSplitArea__000007_Slider_Area.ShowNameInHeader = true
	__AsSplitArea__000007_Slider_Area.Size = 100.000000
	__AsSplitArea__000007_Slider_Area.IsAny = false

	__Doc__000000_doc_in_panel.Name = `doc in panel`
	__Doc__000000_doc_in_panel.StackName = `github.com/fullstack-lang/gong/lib/split/go/models`

	__Form__000000_Form_for_Panel_Form.Name = `Form for Panel Form`
	__Form__000000_Form_for_Panel_Form.StackName = `split-form`
	__Form__000000_Form_for_Panel_Form.FormName = `Form`

	__Slider__000000_Slider.Name = `Slider`
	__Slider__000000_Slider.StackName = `slider`

	__Split__000000_doc_of_split_probe.Name = `doc of split-probe`
	__Split__000000_doc_of_split_probe.StackName = `split-probe`

	__Svg__000000_svg_component.Name = `svg component`
	__Svg__000000_svg_component.StackName = `github.com/fullstack-lang/gong/lib/split/go/models`

	__Table__000000_Table_within_panel.Name = `Table within panel`
	__Table__000000_Table_within_panel.StackName = `split-table`
	__Table__000000_Table_within_panel.TableName = `Table`

	__Tree__000000_Tree.Name = `Tree`
	__Tree__000000_Tree.StackName = `split-sidebar`
	__Tree__000000_Tree.TreeName = `gong`

	__View__000000_doc_reconstructed.Name = `doc reconstructed`

	__View__000001_slider.Name = `slider`

	__View__000002_doc_of_split.Name = `doc of split`

	// Setup of pointers
	// setup of AsSplit instances pointers
	__AsSplit__000000_planel_root.AsSplitAreas = append(__AsSplit__000000_planel_root.AsSplitAreas, __AsSplitArea__000002_sidebar_tree)
	__AsSplit__000000_planel_root.AsSplitAreas = append(__AsSplit__000000_planel_root.AsSplitAreas, __AsSplitArea__000003_table)
	__AsSplit__000000_planel_root.AsSplitAreas = append(__AsSplit__000000_planel_root.AsSplitAreas, __AsSplitArea__000004_form)
	__AsSplit__000000_planel_root.AsSplitAreas = append(__AsSplit__000000_planel_root.AsSplitAreas, __AsSplitArea__000005_dynmicaly_load_component)
	// setup of AsSplitArea instances pointers
	__AsSplitArea__000000_root_for_tree_table_form.AsSplits = append(__AsSplitArea__000000_root_for_tree_table_form.AsSplits, __AsSplit__000000_planel_root)
	__AsSplitArea__000001_doc_component.Doc = __Doc__000000_doc_in_panel
	__AsSplitArea__000002_sidebar_tree.Tree = __Tree__000000_Tree
	__AsSplitArea__000003_table.Table = __Table__000000_Table_within_panel
	__AsSplitArea__000004_form.Form = __Form__000000_Form_for_Panel_Form
	__AsSplitArea__000006_doc_of_split.Split = __Split__000000_doc_of_split_probe
	__AsSplitArea__000007_Slider_Area.Slider = __Slider__000000_Slider
	// setup of Doc instances pointers
	// setup of Form instances pointers
	// setup of Slider instances pointers
	// setup of Split instances pointers
	// setup of Svg instances pointers
	// setup of Table instances pointers
	// setup of Tree instances pointers
	// setup of View instances pointers
	__View__000000_doc_reconstructed.RootAsSplitAreas = append(__View__000000_doc_reconstructed.RootAsSplitAreas, __AsSplitArea__000000_root_for_tree_table_form)
	__View__000000_doc_reconstructed.RootAsSplitAreas = append(__View__000000_doc_reconstructed.RootAsSplitAreas, __AsSplitArea__000001_doc_component)
	__View__000001_slider.RootAsSplitAreas = append(__View__000001_slider.RootAsSplitAreas, __AsSplitArea__000007_Slider_Area)
	__View__000002_doc_of_split.RootAsSplitAreas = append(__View__000002_doc_of_split.RootAsSplitAreas, __AsSplitArea__000006_doc_of_split)
}
