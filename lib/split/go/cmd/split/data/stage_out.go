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
	__AsSplitArea__000006_split_probe := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000007_Slider_Area := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000008_doc_of_slider := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000009_slider_1_probe := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000010_slider_2 := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000011_slider_2_probe := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000012_button_area := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000013_button_probe_area := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000014_tone_area := (&models.AsSplitArea{}).Stage(stage)
	__AsSplitArea__000015_tone_probe_area := (&models.AsSplitArea{}).Stage(stage)

	__Button__000000_button := (&models.Button{}).Stage(stage)

	__Doc__000000_doc_in_panel := (&models.Doc{}).Stage(stage)

	__Form__000000_Form_for_Panel_Form := (&models.Form{}).Stage(stage)

	__Slider__000000_slider_1 := (&models.Slider{}).Stage(stage)
	__Slider__000001_slider_2 := (&models.Slider{}).Stage(stage)

	__Split__000000_split_probe := (&models.Split{}).Stage(stage)
	__Split__000001_slider_1_probe := (&models.Split{}).Stage(stage)
	__Split__000002_slider_2_probe := (&models.Split{}).Stage(stage)
	__Split__000003_button_probe := (&models.Split{}).Stage(stage)
	__Split__000004_tone_probe := (&models.Split{}).Stage(stage)

	__Svg__000000_svg_component := (&models.Svg{}).Stage(stage)

	__Table__000000_Table_within_panel := (&models.Table{}).Stage(stage)

	__Tone__000000_tone := (&models.Tone{}).Stage(stage)

	__Tree__000000_Tree := (&models.Tree{}).Stage(stage)

	__View__000000_doc_reconstructed := (&models.View{}).Stage(stage)
	__View__000001_view_of_split_probe := (&models.View{}).Stage(stage)
	__View__000002_view_of_slider_1 := (&models.View{}).Stage(stage)
	__View__000003_view_of_slider_1_probe := (&models.View{}).Stage(stage)
	__View__000004_view_of_slider_2 := (&models.View{}).Stage(stage)
	__View__000005_view_of_slider_2_probe := (&models.View{}).Stage(stage)
	__View__000006_view_of_button := (&models.View{}).Stage(stage)
	__View__000007_view_of_button_probe := (&models.View{}).Stage(stage)
	__View__000008_view_of_tone := (&models.View{}).Stage(stage)
	__View__000009_view_of_tone_probe := (&models.View{}).Stage(stage)

	// Setup of values

	__AsSplit__000000_planel_root.Name = `planel root`
	__AsSplit__000000_planel_root.Direction = models.Horizontal

	__AsSplitArea__000000_root_for_tree_table_form.Name = `root for tree/table/form`
	__AsSplitArea__000000_root_for_tree_table_form.ShowNameInHeader = false
	__AsSplitArea__000000_root_for_tree_table_form.Size = 80.000000
	__AsSplitArea__000000_root_for_tree_table_form.IsAny = false

	__AsSplitArea__000001_doc_component.Name = `doc component`
	__AsSplitArea__000001_doc_component.ShowNameInHeader = false
	__AsSplitArea__000001_doc_component.Size = 20.000000
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

	__AsSplitArea__000006_split_probe.Name = `split-probe`
	__AsSplitArea__000006_split_probe.ShowNameInHeader = true
	__AsSplitArea__000006_split_probe.Size = 100.000000
	__AsSplitArea__000006_split_probe.IsAny = false

	__AsSplitArea__000007_Slider_Area.Name = `Slider Area`
	__AsSplitArea__000007_Slider_Area.ShowNameInHeader = true
	__AsSplitArea__000007_Slider_Area.Size = 100.000000
	__AsSplitArea__000007_Slider_Area.IsAny = false

	__AsSplitArea__000008_doc_of_slider.Name = `doc of slider`
	__AsSplitArea__000008_doc_of_slider.ShowNameInHeader = true
	__AsSplitArea__000008_doc_of_slider.Size = 100.000000
	__AsSplitArea__000008_doc_of_slider.IsAny = false

	__AsSplitArea__000009_slider_1_probe.Name = `slider 1 probe`
	__AsSplitArea__000009_slider_1_probe.ShowNameInHeader = true
	__AsSplitArea__000009_slider_1_probe.Size = 0.000000
	__AsSplitArea__000009_slider_1_probe.IsAny = false

	__AsSplitArea__000010_slider_2.Name = `slider 2`
	__AsSplitArea__000010_slider_2.ShowNameInHeader = true
	__AsSplitArea__000010_slider_2.Size = 100.000000
	__AsSplitArea__000010_slider_2.IsAny = false

	__AsSplitArea__000011_slider_2_probe.Name = `slider 2 probe`
	__AsSplitArea__000011_slider_2_probe.ShowNameInHeader = true
	__AsSplitArea__000011_slider_2_probe.Size = 100.000000
	__AsSplitArea__000011_slider_2_probe.IsAny = false

	__AsSplitArea__000012_button_area.Name = `button area`
	__AsSplitArea__000012_button_area.ShowNameInHeader = true
	__AsSplitArea__000012_button_area.Size = 100.000000
	__AsSplitArea__000012_button_area.IsAny = false

	__AsSplitArea__000013_button_probe_area.Name = `button probe area`
	__AsSplitArea__000013_button_probe_area.ShowNameInHeader = true
	__AsSplitArea__000013_button_probe_area.Size = 100.000000
	__AsSplitArea__000013_button_probe_area.IsAny = false

	__AsSplitArea__000014_tone_area.Name = `tone area`
	__AsSplitArea__000014_tone_area.ShowNameInHeader = true
	__AsSplitArea__000014_tone_area.Size = 100.000000
	__AsSplitArea__000014_tone_area.IsAny = false

	__AsSplitArea__000015_tone_probe_area.Name = `tone probe area`
	__AsSplitArea__000015_tone_probe_area.ShowNameInHeader = true
	__AsSplitArea__000015_tone_probe_area.Size = 100.000000
	__AsSplitArea__000015_tone_probe_area.IsAny = false

	__Button__000000_button.Name = `button`
	__Button__000000_button.StackName = `button`

	__Doc__000000_doc_in_panel.Name = `doc in panel`
	__Doc__000000_doc_in_panel.StackName = `github.com/fullstack-lang/gong/lib/split/go/models`

	__Form__000000_Form_for_Panel_Form.Name = `Form for Panel Form`
	__Form__000000_Form_for_Panel_Form.StackName = `split-form`
	__Form__000000_Form_for_Panel_Form.FormName = `Form`

	__Slider__000000_slider_1.Name = `slider 1`
	__Slider__000000_slider_1.StackName = `slider 1`

	__Slider__000001_slider_2.Name = `slider 2`
	__Slider__000001_slider_2.StackName = `slider 2`

	__Split__000000_split_probe.Name = `split-probe`
	__Split__000000_split_probe.StackName = `split-probe`

	__Split__000001_slider_1_probe.Name = `slider 1-probe`
	__Split__000001_slider_1_probe.StackName = `slider 1-probe`

	__Split__000002_slider_2_probe.Name = `slider 2 probe`
	__Split__000002_slider_2_probe.StackName = `slider 2-probe`

	__Split__000003_button_probe.Name = `button probe`
	__Split__000003_button_probe.StackName = `button-probe`

	__Split__000004_tone_probe.Name = `tone probe`
	__Split__000004_tone_probe.StackName = `tone-probe`

	__Svg__000000_svg_component.Name = `svg component`
	__Svg__000000_svg_component.StackName = `github.com/fullstack-lang/gong/lib/split/go/models`

	__Table__000000_Table_within_panel.Name = `Table within panel`
	__Table__000000_Table_within_panel.StackName = `split-table`
	__Table__000000_Table_within_panel.TableName = `Table`

	__Tone__000000_tone.Name = `tone`
	__Tone__000000_tone.StackName = `tone`

	__Tree__000000_Tree.Name = `Tree`
	__Tree__000000_Tree.StackName = `split-sidebar`
	__Tree__000000_Tree.TreeName = `gong`

	__View__000000_doc_reconstructed.Name = `doc reconstructed`

	__View__000001_view_of_split_probe.Name = `view of split-probe`

	__View__000002_view_of_slider_1.Name = `view of slider 1`

	__View__000003_view_of_slider_1_probe.Name = `view of slider 1-probe`

	__View__000004_view_of_slider_2.Name = `view of slider 2`

	__View__000005_view_of_slider_2_probe.Name = `view of slider 2 probe`

	__View__000006_view_of_button.Name = `view of button`

	__View__000007_view_of_button_probe.Name = `view of button probe`

	__View__000008_view_of_tone.Name = `view of tone`

	__View__000009_view_of_tone_probe.Name = `view of tone probe`

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
	__AsSplitArea__000006_split_probe.Split = __Split__000000_split_probe
	__AsSplitArea__000007_Slider_Area.Slider = __Slider__000000_slider_1
	__AsSplitArea__000009_slider_1_probe.Split = __Split__000001_slider_1_probe
	__AsSplitArea__000010_slider_2.Slider = __Slider__000001_slider_2
	__AsSplitArea__000011_slider_2_probe.Split = __Split__000002_slider_2_probe
	__AsSplitArea__000012_button_area.Button = __Button__000000_button
	__AsSplitArea__000013_button_probe_area.Split = __Split__000003_button_probe
	__AsSplitArea__000014_tone_area.Tone = __Tone__000000_tone
	__AsSplitArea__000015_tone_probe_area.Split = __Split__000004_tone_probe
	// setup of Button instances pointers
	// setup of Doc instances pointers
	// setup of Form instances pointers
	// setup of Slider instances pointers
	// setup of Split instances pointers
	// setup of Svg instances pointers
	// setup of Table instances pointers
	// setup of Tone instances pointers
	// setup of Tree instances pointers
	// setup of View instances pointers
	__View__000000_doc_reconstructed.RootAsSplitAreas = append(__View__000000_doc_reconstructed.RootAsSplitAreas, __AsSplitArea__000000_root_for_tree_table_form)
	__View__000000_doc_reconstructed.RootAsSplitAreas = append(__View__000000_doc_reconstructed.RootAsSplitAreas, __AsSplitArea__000001_doc_component)
	__View__000001_view_of_split_probe.RootAsSplitAreas = append(__View__000001_view_of_split_probe.RootAsSplitAreas, __AsSplitArea__000006_split_probe)
	__View__000002_view_of_slider_1.RootAsSplitAreas = append(__View__000002_view_of_slider_1.RootAsSplitAreas, __AsSplitArea__000007_Slider_Area)
	__View__000003_view_of_slider_1_probe.RootAsSplitAreas = append(__View__000003_view_of_slider_1_probe.RootAsSplitAreas, __AsSplitArea__000009_slider_1_probe)
	__View__000004_view_of_slider_2.RootAsSplitAreas = append(__View__000004_view_of_slider_2.RootAsSplitAreas, __AsSplitArea__000010_slider_2)
	__View__000005_view_of_slider_2_probe.RootAsSplitAreas = append(__View__000005_view_of_slider_2_probe.RootAsSplitAreas, __AsSplitArea__000011_slider_2_probe)
	__View__000006_view_of_button.RootAsSplitAreas = append(__View__000006_view_of_button.RootAsSplitAreas, __AsSplitArea__000012_button_area)
	__View__000007_view_of_button_probe.RootAsSplitAreas = append(__View__000007_view_of_button_probe.RootAsSplitAreas, __AsSplitArea__000013_button_probe_area)
	__View__000008_view_of_tone.RootAsSplitAreas = append(__View__000008_view_of_tone.RootAsSplitAreas, __AsSplitArea__000014_tone_area)
	__View__000009_view_of_tone_probe.RootAsSplitAreas = append(__View__000009_view_of_tone_probe.RootAsSplitAreas, __AsSplitArea__000015_tone_probe_area)
}
