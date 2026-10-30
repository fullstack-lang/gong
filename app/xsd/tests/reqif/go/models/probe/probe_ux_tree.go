// generated code - do not edit
package probe

import (
	"fmt"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/reqif/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/reqif/go/models/xmlxsd"
)

func (probe *Probe) ux_navigation_tree() {
	stageOfInterest := probe.stageOfInterest
	if !stageOfInterest.IsInDeltaMode() {
		return
	}
	probe.treeNavigationStage.Reset()
	sidebar := &tree_models.Tree{Name: "Sidebar"}
	probe.AddCommitNavigationNode(func(node models.GongNodeIF) {
		sidebar.RootNodes = append(sidebar.RootNodes, node.(*tree_models.Node))
	})
	probe.treeNavigationStage.StageBranch(sidebar)
	probe.treeNavigationStage.Commit()
}

func (probe *Probe) AddCommitNavigationNode(appendChildrenNodeFunc func(models.GongNodeIF)) {
	gongprobe.AddCommitNavigationNode(
		probe.stageOfInterest,
		probe.Refresh,
		probe.AddNotification,
		probe.CommitNotificationTable,
		func(node any) {
			appendChildrenNodeFunc(node.(models.GongNodeIF))
		},
	)
}

func (probe *Probe) ux_tree() {
	probe.ux_navigation_tree()
	probe.treeStage.Reset()

	sidebar := &tree_models.Tree{Name: "Sidebar"}
	topNode := &tree_models.Node{
		Name:       probe.stageOfInterest.GetName(),
		IsExpanded: true,
	}
	topNode.OnIsExpandedChange = func(isExpanded bool) {
		topNode.IsExpanded = isExpanded
	}
	sidebar.RootNodes = append(sidebar.RootNodes, topNode)

	refreshButton := &tree_models.Button{
		Name:            "RefreshButton " + string(tree_buttons.BUTTON_refresh),
		Icon:            string(tree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.Refresh()
		},
	}
	topNode.Buttons = append(topNode.Buttons, refreshButton)

	exportExcelButton := &tree_models.Button{
		Name:            "ExportExcelButton " + string(tree_buttons.BUTTON_download),
		Icon:            string(tree_buttons.BUTTON_download),
		HasToolTip:      true,
		ToolTipText:     "Export XL",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.ExportStageExcel()
		},
	}
	topNode.Buttons = append(topNode.Buttons, exportExcelButton)

	resetButton := &tree_models.Button{
		Name:            "ResetButton " + string(tree_buttons.BUTTON_reset_tv),
		Icon:            string(tree_buttons.BUTTON_reset_tv),
		HasToolTip:      true,
		ToolTipText:     "Reset stage",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.stageOfInterest.Reset()
			probe.stageOfInterest.Commit()
			probe.Refresh()
		},
	}
	topNode.Buttons = append(topNode.Buttons, resetButton)

	exportGoButton := &tree_models.Button{
		Name:            "ExportGoButton " + string(tree_buttons.BUTTON_file_download),
		Icon:            string(tree_buttons.BUTTON_file_download),
		HasToolTip:      true,
		ToolTipText:     "Export Go code for stage",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.ExportStage()
		},
	}
	topNode.Buttons = append(topNode.Buttons, exportGoButton)

	{
		count := len(probe.stageOfInterest.ALTERNATIVE_IDs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ALTERNATIVE_ID (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ALTERNATIVE_ID instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ALTERNATIVE_ID " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ALTERNATIVE_ID",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ALTERNATIVE_ID_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ALTERNATIVE_ID_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ALTERNATIVE_ID]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_DEFINITION_BOOLEANs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_DEFINITION_BOOLEAN (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_DEFINITION_BOOLEAN instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_DEFINITION_BOOLEAN " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_DEFINITION_BOOLEAN",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_DEFINITION_BOOLEAN]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_DEFINITION_DATEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_DEFINITION_DATE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_DEFINITION_DATE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_DEFINITION_DATE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_DEFINITION_DATE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_DEFINITION_DATE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_DEFINITION_DATE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_DEFINITION_DATE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_DEFINITION_ENUMERATIONs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_DEFINITION_ENUMERATION (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_DEFINITION_ENUMERATION instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_DEFINITION_ENUMERATION " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_DEFINITION_ENUMERATION",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_DEFINITION_ENUMERATION]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_DEFINITION_INTEGERs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_DEFINITION_INTEGER (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_DEFINITION_INTEGER instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_DEFINITION_INTEGER " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_DEFINITION_INTEGER",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_DEFINITION_INTEGER_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_DEFINITION_INTEGER_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_DEFINITION_INTEGER]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_DEFINITION_REALs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_DEFINITION_REAL (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_DEFINITION_REAL instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_DEFINITION_REAL " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_DEFINITION_REAL",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_DEFINITION_REAL_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_DEFINITION_REAL_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_DEFINITION_REAL]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_DEFINITION_STRINGs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_DEFINITION_STRING (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_DEFINITION_STRING instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_DEFINITION_STRING " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_DEFINITION_STRING",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_DEFINITION_STRING_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_DEFINITION_STRING_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_DEFINITION_STRING]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_DEFINITION_XHTMLs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_DEFINITION_XHTML (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_DEFINITION_XHTML instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_DEFINITION_XHTML " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_DEFINITION_XHTML",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_DEFINITION_XHTML_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_DEFINITION_XHTML_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_DEFINITION_XHTML]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_VALUE_BOOLEANs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_VALUE_BOOLEAN (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_VALUE_BOOLEAN instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_VALUE_BOOLEAN " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_VALUE_BOOLEAN",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_VALUE_BOOLEAN]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_VALUE_DATEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_VALUE_DATE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_VALUE_DATE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_VALUE_DATE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_VALUE_DATE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_VALUE_DATE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_VALUE_DATE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_VALUE_DATE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_VALUE_ENUMERATIONs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_VALUE_ENUMERATION (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_VALUE_ENUMERATION instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_VALUE_ENUMERATION " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_VALUE_ENUMERATION",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_VALUE_ENUMERATION]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_VALUE_INTEGERs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_VALUE_INTEGER (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_VALUE_INTEGER instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_VALUE_INTEGER " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_VALUE_INTEGER",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_VALUE_INTEGER]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_VALUE_REALs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_VALUE_REAL (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_VALUE_REAL instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_VALUE_REAL " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_VALUE_REAL",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_VALUE_REAL_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_VALUE_REAL_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_VALUE_REAL]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_VALUE_STRINGs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_VALUE_STRING (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_VALUE_STRING instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_VALUE_STRING " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_VALUE_STRING",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_VALUE_STRING_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_VALUE_STRING_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_VALUE_STRING]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ATTRIBUTE_VALUE_XHTMLs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ATTRIBUTE_VALUE_XHTML (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ATTRIBUTE_VALUE_XHTML instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ATTRIBUTE_VALUE_XHTML " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ATTRIBUTE_VALUE_XHTML",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ATTRIBUTE_VALUE_XHTML_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ATTRIBUTE_VALUE_XHTML_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ATTRIBUTE_VALUE_XHTML]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ALTERNATIVE_IDs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ALTERNATIVE_ID (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ALTERNATIVE_ID instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ALTERNATIVE_ID " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ALTERNATIVE_ID",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ALTERNATIVE_ID_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ALTERNATIVE_ID_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ALTERNATIVE_ID]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_DEFINITION_BOOLEAN_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_DEFINITION_BOOLEAN_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_DEFINITION_BOOLEAN_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_DEFINITION_DATE_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_DEFINITION_DATE_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_DEFINITION_DATE_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_DEFINITION_DATE_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_DEFINITION_DATE_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_DEFINITION_DATE_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_DEFINITION_ENUMERATION_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_DEFINITION_ENUMERATION_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_DEFINITION_ENUMERATION_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_DEFINITION_INTEGER_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_DEFINITION_INTEGER_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_DEFINITION_INTEGER_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_DEFINITION_INTEGER_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_DEFINITION_INTEGER_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_DEFINITION_REAL_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_DEFINITION_REAL_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_DEFINITION_REAL_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_DEFINITION_REAL_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_DEFINITION_REAL_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_DEFINITION_REAL_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_DEFINITION_STRING_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_DEFINITION_STRING_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_DEFINITION_STRING_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_DEFINITION_STRING_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_DEFINITION_STRING_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_DEFINITION_STRING_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_DEFINITION_XHTML_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_DEFINITION_XHTML_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_DEFINITION_XHTML_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_DEFINITION_XHTML_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_DEFINITION_XHTML_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_VALUE_BOOLEANs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_VALUE_BOOLEAN (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_VALUE_BOOLEAN instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_VALUE_BOOLEAN " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_VALUE_BOOLEAN",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_VALUE_BOOLEAN]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_VALUE_DATEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_VALUE_DATE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_VALUE_DATE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_VALUE_DATE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_VALUE_DATE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_VALUE_DATE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_VALUE_DATE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_VALUE_DATE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_VALUE_ENUMERATIONs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_VALUE_ENUMERATION (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_VALUE_ENUMERATION instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_VALUE_ENUMERATION " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_VALUE_ENUMERATION",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_VALUE_ENUMERATION]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_VALUE_INTEGERs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_VALUE_INTEGER (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_VALUE_INTEGER instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_VALUE_INTEGER " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_VALUE_INTEGER",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_VALUE_INTEGER]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_VALUE_REALs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_VALUE_REAL (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_VALUE_REAL instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_VALUE_REAL " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_VALUE_REAL",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_VALUE_REAL_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_VALUE_REAL_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_VALUE_REAL]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_VALUE_STRINGs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_VALUE_STRING (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_VALUE_STRING instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_VALUE_STRING " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_VALUE_STRING",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_VALUE_STRING_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_VALUE_STRING_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_VALUE_STRING]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTMLs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_VALUE_XHTML (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_VALUE_XHTML instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_VALUE_XHTML " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_VALUE_XHTML",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_VALUE_XHTML_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_VALUE_XHTML_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_VALUE_XHTML]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTML_1s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ATTRIBUTE_VALUE_XHTML_1 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ATTRIBUTE_VALUE_XHTML_1 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ATTRIBUTE_VALUE_XHTML_1 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ATTRIBUTE_VALUE_XHTML_1",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ATTRIBUTE_VALUE_XHTML_1_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ATTRIBUTE_VALUE_XHTML_1_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ATTRIBUTE_VALUE_XHTML_1]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_CHILDRENs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_CHILDREN (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_CHILDREN instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_CHILDREN " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_CHILDREN",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_CHILDREN_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_CHILDREN_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_CHILDREN]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_CORE_CONTENTs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_CORE_CONTENT (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_CORE_CONTENT instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_CORE_CONTENT " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_CORE_CONTENT",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_CORE_CONTENT_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_CORE_CONTENT_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_CORE_CONTENT]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_DATATYPESs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_DATATYPES (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_DATATYPES instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_DATATYPES " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_DATATYPES",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_DATATYPES_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_DATATYPES_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_DATATYPES]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_DATATYPE_DEFINITION_BOOLEAN_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_DATATYPE_DEFINITION_BOOLEAN_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_DATATYPE_DEFINITION_BOOLEAN_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_DATATYPE_DEFINITION_BOOLEAN_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_DATATYPE_DEFINITION_BOOLEAN_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_DATATYPE_DEFINITION_DATE_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_DATATYPE_DEFINITION_DATE_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_DATATYPE_DEFINITION_DATE_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_DATATYPE_DEFINITION_DATE_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_DATATYPE_DEFINITION_DATE_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_DATATYPE_DEFINITION_DATE_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_DATATYPE_DEFINITION_DATE_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_DATATYPE_DEFINITION_DATE_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_DATATYPE_DEFINITION_ENUMERATION_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_DATATYPE_DEFINITION_ENUMERATION_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_DATATYPE_DEFINITION_ENUMERATION_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_DATATYPE_DEFINITION_ENUMERATION_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_DATATYPE_DEFINITION_ENUMERATION_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_DATATYPE_DEFINITION_INTEGER_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_DATATYPE_DEFINITION_INTEGER_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_DATATYPE_DEFINITION_INTEGER_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_DATATYPE_DEFINITION_INTEGER_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_DATATYPE_DEFINITION_INTEGER_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_DATATYPE_DEFINITION_INTEGER_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_DATATYPE_DEFINITION_REAL_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_DATATYPE_DEFINITION_REAL_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_DATATYPE_DEFINITION_REAL_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_DATATYPE_DEFINITION_REAL_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_DATATYPE_DEFINITION_REAL_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_DATATYPE_DEFINITION_REAL_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_DATATYPE_DEFINITION_REAL_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_DATATYPE_DEFINITION_REAL_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_DATATYPE_DEFINITION_STRING_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_DATATYPE_DEFINITION_STRING_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_DATATYPE_DEFINITION_STRING_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_DATATYPE_DEFINITION_STRING_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_DATATYPE_DEFINITION_STRING_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_DATATYPE_DEFINITION_STRING_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_DATATYPE_DEFINITION_STRING_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_DATATYPE_DEFINITION_STRING_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_DATATYPE_DEFINITION_XHTML_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_DATATYPE_DEFINITION_XHTML_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_DATATYPE_DEFINITION_XHTML_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_DATATYPE_DEFINITION_XHTML_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_DATATYPE_DEFINITION_XHTML_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_DATATYPE_DEFINITION_XHTML_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_DATATYPE_DEFINITION_XHTML_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_DATATYPE_DEFINITION_XHTML_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_EDITABLE_ATTSs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_EDITABLE_ATTS (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_EDITABLE_ATTS instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_EDITABLE_ATTS " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_EDITABLE_ATTS",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_EDITABLE_ATTS_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_EDITABLE_ATTS_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_EDITABLE_ATTS]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_ENUM_VALUE_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_ENUM_VALUE_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_ENUM_VALUE_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_ENUM_VALUE_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_ENUM_VALUE_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_ENUM_VALUE_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_ENUM_VALUE_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_ENUM_VALUE_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_OBJECTs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_OBJECT (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_OBJECT instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_OBJECT " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_OBJECT",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_OBJECT_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_OBJECT_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_OBJECT]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_PROPERTIESs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_PROPERTIES (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_PROPERTIES instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_PROPERTIES " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_PROPERTIES",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_PROPERTIES_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_PROPERTIES_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_PROPERTIES]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_RELATION_GROUP_TYPE_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_RELATION_GROUP_TYPE_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_RELATION_GROUP_TYPE_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_RELATION_GROUP_TYPE_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_RELATION_GROUP_TYPE_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_RELATION_GROUP_TYPE_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_RELATION_GROUP_TYPE_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_RELATION_GROUP_TYPE_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SOURCEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SOURCE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SOURCE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SOURCE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SOURCE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SOURCE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SOURCE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SOURCE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SOURCE_SPECIFICATIONs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SOURCE_SPECIFICATION (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SOURCE_SPECIFICATION instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SOURCE_SPECIFICATION " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SOURCE_SPECIFICATION",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SOURCE_SPECIFICATION_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SOURCE_SPECIFICATION_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SOURCE_SPECIFICATION]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPECIFICATIONSs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPECIFICATIONS (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPECIFICATIONS instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPECIFICATIONS " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPECIFICATIONS",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPECIFICATIONS_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPECIFICATIONS_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPECIFICATIONS]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPECIFICATION_TYPE_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPECIFICATION_TYPE_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPECIFICATION_TYPE_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPECIFICATION_TYPE_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPECIFICATION_TYPE_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPECIFICATION_TYPE_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPECIFICATION_TYPE_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPECIFICATION_TYPE_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPECIFIED_VALUESs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPECIFIED_VALUES (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPECIFIED_VALUES instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPECIFIED_VALUES " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPECIFIED_VALUES",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPECIFIED_VALUES_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPECIFIED_VALUES_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPECIFIED_VALUES]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPEC_ATTRIBUTESs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPEC_ATTRIBUTES (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPEC_ATTRIBUTES instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPEC_ATTRIBUTES " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPEC_ATTRIBUTES",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPEC_ATTRIBUTES_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPEC_ATTRIBUTES_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPEC_ATTRIBUTES]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPEC_OBJECTSs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPEC_OBJECTS (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPEC_OBJECTS instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPEC_OBJECTS " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPEC_OBJECTS",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPEC_OBJECTS_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPEC_OBJECTS_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPEC_OBJECTS]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPEC_OBJECT_TYPE_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPEC_OBJECT_TYPE_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPEC_OBJECT_TYPE_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPEC_OBJECT_TYPE_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPEC_OBJECT_TYPE_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPEC_OBJECT_TYPE_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPEC_OBJECT_TYPE_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPEC_OBJECT_TYPE_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPEC_RELATIONSs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPEC_RELATIONS (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPEC_RELATIONS instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPEC_RELATIONS " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPEC_RELATIONS",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPEC_RELATIONS_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPEC_RELATIONS_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPEC_RELATIONS]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPEC_RELATION_GROUPSs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPEC_RELATION_GROUPS (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPEC_RELATION_GROUPS instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPEC_RELATION_GROUPS " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPEC_RELATION_GROUPS",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPEC_RELATION_GROUPS_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPEC_RELATION_GROUPS_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPEC_RELATION_GROUPS]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPEC_RELATION_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPEC_RELATION_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPEC_RELATION_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPEC_RELATION_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPEC_RELATION_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPEC_RELATION_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPEC_RELATION_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPEC_RELATION_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPEC_RELATION_TYPE_REFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPEC_RELATION_TYPE_REF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPEC_RELATION_TYPE_REF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPEC_RELATION_TYPE_REF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPEC_RELATION_TYPE_REF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPEC_RELATION_TYPE_REF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPEC_RELATION_TYPE_REF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPEC_RELATION_TYPE_REF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_SPEC_TYPESs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_SPEC_TYPES (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_SPEC_TYPES instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_SPEC_TYPES " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_SPEC_TYPES",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_SPEC_TYPES_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_SPEC_TYPES_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_SPEC_TYPES]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_THE_HEADERs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_THE_HEADER (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_THE_HEADER instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_THE_HEADER " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_THE_HEADER",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_THE_HEADER_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_THE_HEADER_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_THE_HEADER]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.A_TOOL_EXTENSIONSs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_TOOL_EXTENSIONS (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_TOOL_EXTENSIONS instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_TOOL_EXTENSIONS " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_TOOL_EXTENSIONS",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_TOOL_EXTENSIONS_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_TOOL_EXTENSIONS_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.A_TOOL_EXTENSIONS]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.DATATYPE_DEFINITION_BOOLEANs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DATATYPE_DEFINITION_BOOLEAN (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DATATYPE_DEFINITION_BOOLEAN instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DATATYPE_DEFINITION_BOOLEAN " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DATATYPE_DEFINITION_BOOLEAN",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DATATYPE_DEFINITION_BOOLEAN_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DATATYPE_DEFINITION_BOOLEAN_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.DATATYPE_DEFINITION_BOOLEAN]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.DATATYPE_DEFINITION_DATEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DATATYPE_DEFINITION_DATE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DATATYPE_DEFINITION_DATE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DATATYPE_DEFINITION_DATE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DATATYPE_DEFINITION_DATE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DATATYPE_DEFINITION_DATE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DATATYPE_DEFINITION_DATE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.DATATYPE_DEFINITION_DATE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.DATATYPE_DEFINITION_ENUMERATIONs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DATATYPE_DEFINITION_ENUMERATION (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DATATYPE_DEFINITION_ENUMERATION instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DATATYPE_DEFINITION_ENUMERATION " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DATATYPE_DEFINITION_ENUMERATION",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DATATYPE_DEFINITION_ENUMERATION_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DATATYPE_DEFINITION_ENUMERATION_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.DATATYPE_DEFINITION_ENUMERATION]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.DATATYPE_DEFINITION_INTEGERs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DATATYPE_DEFINITION_INTEGER (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DATATYPE_DEFINITION_INTEGER instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DATATYPE_DEFINITION_INTEGER " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DATATYPE_DEFINITION_INTEGER",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DATATYPE_DEFINITION_INTEGER_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DATATYPE_DEFINITION_INTEGER_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.DATATYPE_DEFINITION_INTEGER]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.DATATYPE_DEFINITION_REALs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DATATYPE_DEFINITION_REAL (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DATATYPE_DEFINITION_REAL instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DATATYPE_DEFINITION_REAL " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DATATYPE_DEFINITION_REAL",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DATATYPE_DEFINITION_REAL_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DATATYPE_DEFINITION_REAL_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.DATATYPE_DEFINITION_REAL]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.DATATYPE_DEFINITION_STRINGs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DATATYPE_DEFINITION_STRING (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DATATYPE_DEFINITION_STRING instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DATATYPE_DEFINITION_STRING " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DATATYPE_DEFINITION_STRING",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DATATYPE_DEFINITION_STRING_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DATATYPE_DEFINITION_STRING_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.DATATYPE_DEFINITION_STRING]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.DATATYPE_DEFINITION_XHTMLs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DATATYPE_DEFINITION_XHTML (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DATATYPE_DEFINITION_XHTML instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DATATYPE_DEFINITION_XHTML " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DATATYPE_DEFINITION_XHTML",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DATATYPE_DEFINITION_XHTML_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DATATYPE_DEFINITION_XHTML_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.DATATYPE_DEFINITION_XHTML]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.EMBEDDED_VALUEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("EMBEDDED_VALUE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all EMBEDDED_VALUE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "EMBEDDED_VALUE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of EMBEDDED_VALUE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_EMBEDDED_VALUE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_EMBEDDED_VALUE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.EMBEDDED_VALUE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.ENUM_VALUEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ENUM_VALUE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ENUM_VALUE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ENUM_VALUE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ENUM_VALUE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ENUM_VALUE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ENUM_VALUE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.ENUM_VALUE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.RELATION_GROUPs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RELATION_GROUP (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RELATION_GROUP instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RELATION_GROUP " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RELATION_GROUP",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RELATION_GROUP_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RELATION_GROUP_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.RELATION_GROUP]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.RELATION_GROUP_TYPEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RELATION_GROUP_TYPE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RELATION_GROUP_TYPE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RELATION_GROUP_TYPE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RELATION_GROUP_TYPE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RELATION_GROUP_TYPE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RELATION_GROUP_TYPE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.RELATION_GROUP_TYPE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.REQ_IFs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("REQ_IF (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all REQ_IF instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "REQ_IF " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of REQ_IF",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_REQ_IF_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_REQ_IF_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.REQ_IF]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.REQ_IF_CONTENTs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("REQ_IF_CONTENT (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all REQ_IF_CONTENT instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "REQ_IF_CONTENT " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of REQ_IF_CONTENT",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_REQ_IF_CONTENT_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_REQ_IF_CONTENT_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.REQ_IF_CONTENT]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.REQ_IF_HEADERs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("REQ_IF_HEADER (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all REQ_IF_HEADER instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "REQ_IF_HEADER " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of REQ_IF_HEADER",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_REQ_IF_HEADER_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_REQ_IF_HEADER_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.REQ_IF_HEADER]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.REQ_IF_TOOL_EXTENSIONs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("REQ_IF_TOOL_EXTENSION (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all REQ_IF_TOOL_EXTENSION instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "REQ_IF_TOOL_EXTENSION " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of REQ_IF_TOOL_EXTENSION",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_REQ_IF_TOOL_EXTENSION_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_REQ_IF_TOOL_EXTENSION_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.REQ_IF_TOOL_EXTENSION]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.SPECIFICATIONs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SPECIFICATION (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SPECIFICATION instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SPECIFICATION " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SPECIFICATION",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SPECIFICATION_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SPECIFICATION_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.SPECIFICATION]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.SPECIFICATION_TYPEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SPECIFICATION_TYPE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SPECIFICATION_TYPE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SPECIFICATION_TYPE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SPECIFICATION_TYPE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SPECIFICATION_TYPE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SPECIFICATION_TYPE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.SPECIFICATION_TYPE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.SPEC_HIERARCHYs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SPEC_HIERARCHY (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SPEC_HIERARCHY instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SPEC_HIERARCHY " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SPEC_HIERARCHY",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SPEC_HIERARCHY_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SPEC_HIERARCHY_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.SPEC_HIERARCHY]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.SPEC_OBJECTs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SPEC_OBJECT (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SPEC_OBJECT instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SPEC_OBJECT " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SPEC_OBJECT",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SPEC_OBJECT_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SPEC_OBJECT_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.SPEC_OBJECT]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.SPEC_OBJECT_TYPEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SPEC_OBJECT_TYPE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SPEC_OBJECT_TYPE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SPEC_OBJECT_TYPE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SPEC_OBJECT_TYPE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SPEC_OBJECT_TYPE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SPEC_OBJECT_TYPE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.SPEC_OBJECT_TYPE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.SPEC_RELATIONs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SPEC_RELATION (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SPEC_RELATION instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SPEC_RELATION " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SPEC_RELATION",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SPEC_RELATION_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SPEC_RELATION_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.SPEC_RELATION]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.SPEC_RELATION_TYPEs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SPEC_RELATION_TYPE (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SPEC_RELATION_TYPE instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SPEC_RELATION_TYPE " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SPEC_RELATION_TYPE",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SPEC_RELATION_TYPE_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SPEC_RELATION_TYPE_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.SPEC_RELATION_TYPE]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.XHTML_CONTENTs)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("XHTML_CONTENT (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all XHTML_CONTENT instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "XHTML_CONTENT " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of XHTML_CONTENT",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_XHTML_CONTENT_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_XHTML_CONTENT_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.XHTML_CONTENT]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	// Package node: xmlxsd
	pkgNode_XmlxsdStage := &tree_models.Node{
		Name:       "xmlxsd",
		IsExpanded: true,
	}
	pkgNode_XmlxsdStage.OnIsExpandedChange = func(isExpanded bool) {
		pkgNode_XmlxsdStage.IsExpanded = isExpanded
	}
	topNode.Children = append(topNode.Children, pkgNode_XmlxsdStage)

	{
		count := len(probe.stageOfInterest.XmlxsdStage.Xmlxsd_SchemaElements)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("Xmlxsd_SchemaElement (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all Xmlxsd_SchemaElement instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_XmlxsdStage.Children = append(pkgNode_XmlxsdStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "Xmlxsd_SchemaElement " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of Xmlxsd_SchemaElement",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_Xmlxsd_SchemaElement_XmlxsdStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_Xmlxsd_SchemaElement_XmlxsdStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.XmlxsdStage.GetInstancesByOrder[*xmlxsd.Xmlxsd_SchemaElement]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	probe.treeStage.StageBranch(sidebar)
	probe.treeStage.Commit()
}
