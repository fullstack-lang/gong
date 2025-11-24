package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

func updateAndCommitTree(
	probe *Probe,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := &tree.Tree{Name: "Sidebar"}

	nodeRefreshButton := &tree.Node{Name: fmt.Sprintf("Stage %s, # %d, %s",
		probe.stageOfInterest.GetName(),
		probe.stageOfInterest.GetCommitId(),
		probe.stageOfInterest.GetCommitTS().Local().Format(time.Kitchen))}
	nodeRefreshButton.Name +=
		fmt.Sprintf(" (%d/%d/%d)", 
			len(probe.stageOfInterest.GetNew()), 
			len(probe.stageOfInterest.GetModified()), 
			len(probe.stageOfInterest.GetDeleted()),
		)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := &tree.Button{
		Name:            "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon:            string(gongtree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree.Left,
	}

	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSetFromPointerType[*gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := &tree.Node{Name: name}
		nodeGongstruct.HasToolTip = true
		nodeGongstruct.ToolTipText = "Display table of all " + name + " instances"
		nodeGongstruct.ToolTipPosition = tree.Right

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "Cell":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Cell](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _cell := range set {
				nodeInstance := &tree.Node{Name: _cell.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cell, "Cell", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_cell]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_cell]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_cell]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "CellBoolean":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellBoolean](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _cellboolean := range set {
				nodeInstance := &tree.Node{Name: _cellboolean.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellboolean, "CellBoolean", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_cellboolean]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_cellboolean]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_cellboolean]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "CellFloat64":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellFloat64](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _cellfloat64 := range set {
				nodeInstance := &tree.Node{Name: _cellfloat64.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellfloat64, "CellFloat64", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_cellfloat64]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_cellfloat64]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_cellfloat64]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "CellIcon":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellIcon](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _cellicon := range set {
				nodeInstance := &tree.Node{Name: _cellicon.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellicon, "CellIcon", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_cellicon]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_cellicon]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_cellicon]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "CellInt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellInt](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _cellint := range set {
				nodeInstance := &tree.Node{Name: _cellint.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellint, "CellInt", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_cellint]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_cellint]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_cellint]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "CellString":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CellString](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _cellstring := range set {
				nodeInstance := &tree.Node{Name: _cellstring.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cellstring, "CellString", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_cellstring]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_cellstring]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_cellstring]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "CheckBox":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.CheckBox](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _checkbox := range set {
				nodeInstance := &tree.Node{Name: _checkbox.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_checkbox, "CheckBox", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_checkbox]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_checkbox]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_checkbox]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "DisplayedColumn":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.DisplayedColumn](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _displayedcolumn := range set {
				nodeInstance := &tree.Node{Name: _displayedcolumn.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_displayedcolumn, "DisplayedColumn", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_displayedcolumn]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_displayedcolumn]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_displayedcolumn]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormDiv":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormDiv](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formdiv := range set {
				nodeInstance := &tree.Node{Name: _formdiv.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formdiv, "FormDiv", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formdiv]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formdiv]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formdiv]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormEditAssocButton":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormEditAssocButton](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formeditassocbutton := range set {
				nodeInstance := &tree.Node{Name: _formeditassocbutton.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formeditassocbutton, "FormEditAssocButton", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formeditassocbutton]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formeditassocbutton]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formeditassocbutton]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormField":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormField](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formfield := range set {
				nodeInstance := &tree.Node{Name: _formfield.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfield, "FormField", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formfield]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formfield]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formfield]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormFieldDate":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldDate](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formfielddate := range set {
				nodeInstance := &tree.Node{Name: _formfielddate.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfielddate, "FormFieldDate", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formfielddate]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formfielddate]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formfielddate]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormFieldDateTime":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldDateTime](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formfielddatetime := range set {
				nodeInstance := &tree.Node{Name: _formfielddatetime.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfielddatetime, "FormFieldDateTime", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formfielddatetime]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formfielddatetime]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formfielddatetime]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormFieldFloat64":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldFloat64](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formfieldfloat64 := range set {
				nodeInstance := &tree.Node{Name: _formfieldfloat64.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldfloat64, "FormFieldFloat64", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formfieldfloat64]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formfieldfloat64]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formfieldfloat64]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormFieldInt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldInt](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formfieldint := range set {
				nodeInstance := &tree.Node{Name: _formfieldint.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldint, "FormFieldInt", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formfieldint]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formfieldint]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formfieldint]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormFieldSelect":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldSelect](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formfieldselect := range set {
				nodeInstance := &tree.Node{Name: _formfieldselect.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldselect, "FormFieldSelect", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formfieldselect]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formfieldselect]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formfieldselect]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormFieldString":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldString](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formfieldstring := range set {
				nodeInstance := &tree.Node{Name: _formfieldstring.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldstring, "FormFieldString", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formfieldstring]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formfieldstring]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formfieldstring]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormFieldTime":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormFieldTime](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formfieldtime := range set {
				nodeInstance := &tree.Node{Name: _formfieldtime.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formfieldtime, "FormFieldTime", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formfieldtime]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formfieldtime]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formfieldtime]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormGroup":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormGroup](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formgroup := range set {
				nodeInstance := &tree.Node{Name: _formgroup.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formgroup, "FormGroup", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formgroup]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formgroup]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formgroup]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "FormSortAssocButton":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.FormSortAssocButton](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _formsortassocbutton := range set {
				nodeInstance := &tree.Node{Name: _formsortassocbutton.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_formsortassocbutton, "FormSortAssocButton", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_formsortassocbutton]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_formsortassocbutton]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_formsortassocbutton]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Option":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Option](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _option := range set {
				nodeInstance := &tree.Node{Name: _option.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_option, "Option", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_option]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_option]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_option]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Row":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Row](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _row := range set {
				nodeInstance := &tree.Node{Name: _row.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_row, "Row", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_row]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_row]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_row]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		case "Table":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSetFromPointerType[*models.Table](probe.stageOfInterest)
			created := 0
			updated := 0
			deleted := 0
			for _table := range set {
				nodeInstance := &tree.Node{Name: _table.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_table, "Table", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
				if _, ok := probe.stageOfInterest.GetNew()[_table]; ok {
					created++
				}
				if _, ok := probe.stageOfInterest.GetModified()[_table]; ok {
					updated++
				}
				if _, ok := probe.stageOfInterest.GetDeleted()[_table]; ok {
					deleted++
				}
			}
			nodeGongstruct.Name += fmt.Sprintf(" (%d/%d/%d)", created, updated, deleted)
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := &tree.Button{
			Name:            gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon:            string(gongtree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of " + gongStruct.GetName(),
			ToolTipPosition: tree.Right,
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	tree.StageBranch(probe.treeStage, sidebar)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.PointerToGongstruct] struct {
	Instance       T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.PointerToGongstruct](
	instance T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.Stage,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
