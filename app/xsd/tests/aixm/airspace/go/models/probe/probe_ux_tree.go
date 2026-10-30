// generated code - do not edit
package probe

import (
	"fmt"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/xlink"
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
		count := len(probe.stageOfInterest.AIXMBasicMessages)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AIXMBasicMessage (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AIXMBasicMessage instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AIXMBasicMessage " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AIXMBasicMessage",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AIXMBasicMessage_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AIXMBasicMessage_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.AIXMBasicMessage]() {
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
		count := len(probe.stageOfInterest.AIXMBasicMessageTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AIXMBasicMessageType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AIXMBasicMessageType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AIXMBasicMessageType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AIXMBasicMessageType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AIXMBasicMessageType_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AIXMBasicMessageType_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.AIXMBasicMessageType]() {
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
		count := len(probe.stageOfInterest.BasicMessageMemberAIXMPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("BasicMessageMemberAIXMPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all BasicMessageMemberAIXMPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		topNode.Children = append(topNode.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "BasicMessageMemberAIXMPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of BasicMessageMemberAIXMPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_BasicMessageMemberAIXMPropertyType_Stage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_BasicMessageMemberAIXMPropertyType_Stage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GetInstancesByOrder[*models.BasicMessageMemberAIXMPropertyType]() {
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

	// Package node: aixmfeatures
	pkgNode_AixmfeaturesStage := &tree_models.Node{
		Name:       "aixmfeatures",
		IsExpanded: true,
	}
	pkgNode_AixmfeaturesStage.OnIsExpandedChange = func(isExpanded bool) {
		pkgNode_AixmfeaturesStage.IsExpanded = isExpanded
	}
	topNode.Children = append(topNode.Children, pkgNode_AixmfeaturesStage)

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extensions)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_1s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_1 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_1 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_1 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_1",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_1_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_1_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_1]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_10s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_10 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_10 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_10 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_10",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_10_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_10_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_10]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_11s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_11 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_11 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_11 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_11",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_11_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_11_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_11]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_12s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_12 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_12 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_12 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_12",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_12_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_12_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_12]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_13s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_13 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_13 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_13 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_13",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_13_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_13_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_13]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_14s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_14 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_14 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_14 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_14",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_14_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_14_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_14]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_2s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_2 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_2 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_2 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_2",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_2_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_2_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_2]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_3s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_3 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_3 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_3 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_3",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_3_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_3_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_3]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_4s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_4 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_4 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_4 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_4",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_4_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_4_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_4]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_5s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_5 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_5 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_5 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_5",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_5_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_5_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_5]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_6s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_6 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_6 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_6 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_6",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_6_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_6_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_6]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_7s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_7 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_7 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_7 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_7",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_7_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_7_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_7]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_8s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_8 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_8 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_8 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_8",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_8_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_8_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_8]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_9s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_9 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_9 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_9 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_9",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_9_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_9_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_9]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AbstractPropertiesWithScheduleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractPropertiesWithScheduleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractPropertiesWithScheduleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractPropertiesWithScheduleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractPropertiesWithScheduleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractPropertiesWithScheduleType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractPropertiesWithScheduleType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AbstractPropertiesWithScheduleType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AircraftCharacteristicPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AircraftCharacteristicPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AircraftCharacteristicPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AircraftCharacteristicPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AircraftCharacteristicPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AircraftCharacteristicPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AircraftCharacteristicPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AircraftCharacteristicTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AircraftCharacteristicType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AircraftCharacteristicType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AircraftCharacteristicType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AircraftCharacteristicType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AircraftCharacteristicType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AircraftCharacteristicType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AircraftCharacteristicType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceActivationPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceActivationPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceActivationPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceActivationPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceActivationPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceActivationPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceActivationPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceActivationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceActivationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceActivationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceActivationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceActivationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceActivationType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceActivationType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceActivationType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceGeometryComponentPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceGeometryComponentPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceGeometryComponentPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceGeometryComponentPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceGeometryComponentPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceGeometryComponentPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceGeometryComponentPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceGeometryComponentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceGeometryComponentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceGeometryComponentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceGeometryComponentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceGeometryComponentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceGeometryComponentType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceGeometryComponentType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceGeometryComponentType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceLayerClassPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceLayerClassPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceLayerClassPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceLayerClassPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceLayerClassPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceLayerClassPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceLayerClassPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerClassTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceLayerClassType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceLayerClassType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceLayerClassType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceLayerClassType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceLayerClassType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceLayerClassType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceLayerClassType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceLayerPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceLayerPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceLayerPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceLayerPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceLayerPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceLayerPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceLayerPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceLayerTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceLayerType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceLayerType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceLayerType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceLayerType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceLayerType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceLayerType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceLayerType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspacePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspacePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspacePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspacePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspacePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspacePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspacePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspacePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeDependencyPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceVolumeDependencyPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceVolumeDependencyPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceVolumeDependencyPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceVolumeDependencyPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceVolumeDependencyPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceVolumeDependencyPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceVolumeDependencyPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeDependencyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceVolumeDependencyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceVolumeDependencyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceVolumeDependencyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceVolumeDependencyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceVolumeDependencyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceVolumeDependencyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceVolumeDependencyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceVolumePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceVolumePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceVolumePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceVolumePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceVolumePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceVolumePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceVolumePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceVolumeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceVolumeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceVolumeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceVolumeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceVolumeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceVolumeType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceVolumeType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceVolumeType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForAirspaceTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForAirspaceTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForAirspaceTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForAirspaceTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForAirspaceTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForAirspaceTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForAirspaceTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForAirspaceTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForAirspaceTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForAirspaceTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForAirspaceTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForAirspaceTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForAirspaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForAirspaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForAirspaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForAirspaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForAirspaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForAirspaceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForAirspaceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForAirspaceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.CurvePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CurvePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CurvePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CurvePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CurvePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CurvePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CurvePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.CurvePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.CurveTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CurveType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CurveType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CurveType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CurveType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CurveType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CurveType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.CurveType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ElevatedCurveTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ElevatedCurveType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ElevatedCurveType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ElevatedCurveType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ElevatedCurveType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ElevatedCurveType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ElevatedCurveType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedCurveType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ElevatedSurfaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ElevatedSurfaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ElevatedSurfaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ElevatedSurfaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ElevatedSurfaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ElevatedSurfaceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ElevatedSurfaceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedSurfaceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("GeoBorderTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all GeoBorderTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "GeoBorderTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of GeoBorderTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_GeoBorderTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_GeoBorderTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.GeoBorderTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.GeoBorderTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("GeoBorderTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all GeoBorderTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "GeoBorderTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of GeoBorderTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_GeoBorderTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_GeoBorderTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.GeoBorderTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.GeoBorderTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("GeoBorderType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all GeoBorderType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "GeoBorderType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of GeoBorderType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_GeoBorderType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_GeoBorderType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.GeoBorderType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.LinguisticNotePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("LinguisticNotePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all LinguisticNotePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "LinguisticNotePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of LinguisticNotePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_LinguisticNotePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_LinguisticNotePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LinguisticNotePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.LinguisticNoteTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("LinguisticNoteType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all LinguisticNoteType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "LinguisticNoteType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of LinguisticNoteType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_LinguisticNoteType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_LinguisticNoteType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LinguisticNoteType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NotePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NotePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NotePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NotePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NotePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NotePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NotePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NotePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NoteTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NoteType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NoteType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NoteType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NoteType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NoteType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NoteType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NoteType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.OrganisationAuthorityPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("OrganisationAuthorityPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all OrganisationAuthorityPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "OrganisationAuthorityPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of OrganisationAuthorityPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_OrganisationAuthorityPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_OrganisationAuthorityPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.OrganisationAuthorityPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RoutePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RoutePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RoutePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RoutePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RoutePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RoutePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RoutePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RoutePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.StandardLevelColumnPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("StandardLevelColumnPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all StandardLevelColumnPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "StandardLevelColumnPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of StandardLevelColumnPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_StandardLevelColumnPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_StandardLevelColumnPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.StandardLevelColumnPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SurfacePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SurfacePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SurfacePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SurfacePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SurfacePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SurfacePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SurfacePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SurfacePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SurfaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SurfaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SurfaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SurfaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SurfaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SurfaceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SurfaceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SurfaceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.TimesheetPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TimesheetPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TimesheetPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TimesheetPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TimesheetPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TimesheetPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TimesheetPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TimesheetPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.TimesheetTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TimesheetType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TimesheetType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TimesheetType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TimesheetType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TimesheetType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TimesheetType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TimesheetType]() {
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

	// Package node: aixmdatatypes
	pkgNode_AixmdatatypesStage := &tree_models.Node{
		Name:       "aixmdatatypes",
		IsExpanded: true,
	}
	pkgNode_AixmdatatypesStage.OnIsExpandedChange = func(isExpanded bool) {
		pkgNode_AixmdatatypesStage.IsExpanded = isExpanded
	}
	topNode.Children = append(topNode.Children, pkgNode_AixmdatatypesStage)

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAircraftCategoryTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAircraftCategoryType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAircraftCategoryType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAircraftCategoryType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAircraftCategoryType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAircraftCategoryType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAircraftCategoryType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAircraftCategoryType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAircraftEngineNumberTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAircraftEngineNumberType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAircraftEngineNumberType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAircraftEngineNumberType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAircraftEngineNumberType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAircraftEngineNumberType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAircraftEngineNumberType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAircraftEngineNumberType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAircraftEngineTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAircraftEngineType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAircraftEngineType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAircraftEngineType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAircraftEngineType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAircraftEngineType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAircraftEngineType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAircraftEngineType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAircraftICAOTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAircraftICAOType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAircraftICAOType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAircraftICAOType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAircraftICAOType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAircraftICAOType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAircraftICAOType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAircraftICAOType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAircraftTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAircraftType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAircraftType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAircraftType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAircraftType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAircraftType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAircraftType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAircraftType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAircraftWingspanClassTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAircraftWingspanClassType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAircraftWingspanClassType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAircraftWingspanClassType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAircraftWingspanClassType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAircraftWingspanClassType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAircraftWingspanClassType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAircraftWingspanClassType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAirspaceActivityTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAirspaceActivityType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAirspaceActivityType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAirspaceActivityType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAirspaceActivityType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAirspaceActivityType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAirspaceActivityType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAirspaceActivityType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAirspaceAggregationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAirspaceAggregationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAirspaceAggregationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAirspaceAggregationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAirspaceAggregationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAirspaceAggregationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAirspaceAggregationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAirspaceAggregationType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAirspaceClassificationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAirspaceClassificationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAirspaceClassificationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAirspaceClassificationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAirspaceClassificationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAirspaceClassificationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAirspaceClassificationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAirspaceClassificationType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAirspaceDependencyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAirspaceDependencyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAirspaceDependencyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAirspaceDependencyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAirspaceDependencyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAirspaceDependencyType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAirspaceDependencyType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAirspaceDependencyType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAirspaceDesignatorTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAirspaceDesignatorType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAirspaceDesignatorType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAirspaceDesignatorType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAirspaceDesignatorType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAirspaceDesignatorType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAirspaceDesignatorType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAirspaceDesignatorType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAirspaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAirspaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAirspaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAirspaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAirspaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAirspaceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAirspaceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAirspaceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAltitudeUseTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAltitudeUseType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAltitudeUseType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAltitudeUseType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAltitudeUseType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAltitudeUseType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAltitudeUseType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAltitudeUseType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAuthorityTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAuthorityType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAuthorityType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAuthorityType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAuthorityType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAuthorityType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAuthorityType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAuthorityType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeCommunicationModeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeCommunicationModeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeCommunicationModeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeCommunicationModeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeCommunicationModeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeCommunicationModeType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeCommunicationModeType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeCommunicationModeType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeDayTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeDayType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeDayType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeDayType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeDayType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeDayType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeDayType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDayType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeEquipmentAntiCollisionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeEquipmentAntiCollisionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeEquipmentAntiCollisionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeEquipmentAntiCollisionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeEquipmentAntiCollisionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeEquipmentAntiCollisionType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeEquipmentAntiCollisionType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeEquipmentAntiCollisionType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeGeoBorderTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeGeoBorderType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeGeoBorderType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeGeoBorderType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeGeoBorderType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeGeoBorderType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeGeoBorderType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeGeoBorderType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeMilitaryOperationsTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeMilitaryOperationsType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeMilitaryOperationsType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeMilitaryOperationsType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeMilitaryOperationsType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeMilitaryOperationsType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeMilitaryOperationsType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeMilitaryOperationsType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeNavigationEquipmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeNavigationEquipmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeNavigationEquipmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeNavigationEquipmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeNavigationEquipmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeNavigationEquipmentType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeNavigationEquipmentType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeNavigationEquipmentType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeNavigationSpecificationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeNavigationSpecificationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeNavigationSpecificationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeNavigationSpecificationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeNavigationSpecificationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeNavigationSpecificationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeNavigationSpecificationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeNavigationSpecificationType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeNotePurposeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeNotePurposeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeNotePurposeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeNotePurposeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeNotePurposeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeNotePurposeType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeNotePurposeType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeNotePurposeType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRVSMTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRVSMType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRVSMType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRVSMType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRVSMType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRVSMType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRVSMType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRVSMType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeStatusAirspaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeStatusAirspaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeStatusAirspaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeStatusAirspaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeStatusAirspaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeStatusAirspaceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeStatusAirspaceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeStatusAirspaceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeTimeEventCombinationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeTimeEventCombinationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeTimeEventCombinationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeTimeEventCombinationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeTimeEventCombinationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeTimeEventCombinationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeTimeEventCombinationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeTimeEventCombinationType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeTimeEventTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeTimeEventType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeTimeEventType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeTimeEventType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeTimeEventType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeTimeEventType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeTimeEventType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeTimeEventType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeTimeReferenceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeTimeReferenceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeTimeReferenceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeTimeReferenceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeTimeReferenceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeTimeReferenceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeTimeReferenceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeTimeReferenceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeTransponderTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeTransponderType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeTransponderType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeTransponderType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeTransponderType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeTransponderType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeTransponderType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeTransponderType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeValueInterpretationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeValueInterpretationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeValueInterpretationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeValueInterpretationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeValueInterpretationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeValueInterpretationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeValueInterpretationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeValueInterpretationType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeVerticalDatumTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeVerticalDatumType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeVerticalDatumType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeVerticalDatumType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeVerticalDatumType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeVerticalDatumType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeVerticalDatumType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeVerticalDatumType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeVerticalReferenceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeVerticalReferenceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeVerticalReferenceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeVerticalReferenceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeVerticalReferenceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeVerticalReferenceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeVerticalReferenceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeVerticalReferenceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeWakeTurbulenceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeWakeTurbulenceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeWakeTurbulenceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeWakeTurbulenceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeWakeTurbulenceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeWakeTurbulenceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeWakeTurbulenceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeWakeTurbulenceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeYesNoTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeYesNoType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeYesNoType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeYesNoType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeYesNoType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeYesNoType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeYesNoType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeYesNoType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.DateMonthDayTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DateMonthDayType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DateMonthDayType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DateMonthDayType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DateMonthDayType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DateMonthDayType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DateMonthDayType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.DateMonthDayType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.NoNumberTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NoNumberType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NoNumberType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NoNumberType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NoNumberType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NoNumberType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NoNumberType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.NoNumberType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.NoSequenceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NoSequenceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NoSequenceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NoSequenceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NoSequenceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NoSequenceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NoSequenceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.NoSequenceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.TextNameTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TextNameType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TextNameType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TextNameType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TextNameType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TextNameType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TextNameType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextNameType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.TextNoteTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TextNoteType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TextNoteType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TextNoteType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TextNoteType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TextNoteType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TextNoteType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextNoteType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.TextPropertyNameTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TextPropertyNameType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TextPropertyNameType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TextPropertyNameType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TextPropertyNameType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TextPropertyNameType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TextPropertyNameType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextPropertyNameType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.TimeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TimeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TimeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TimeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TimeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TimeType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TimeType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TimeType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValDistanceSignedTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValDistanceSignedType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValDistanceSignedType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValDistanceSignedType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValDistanceSignedType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValDistanceSignedType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValDistanceSignedType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValDistanceSignedType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValDistanceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValDistanceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValDistanceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValDistanceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValDistanceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValDistanceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValDistanceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValDistanceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValDistanceVerticalTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValDistanceVerticalType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValDistanceVerticalType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValDistanceVerticalType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValDistanceVerticalType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValDistanceVerticalType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValDistanceVerticalType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValDistanceVerticalType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValDurationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValDurationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValDurationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValDurationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValDurationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValDurationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValDurationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValDurationType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValFLTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValFLType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValFLType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValFLType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValFLType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValFLType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValFLType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValFLType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValSpeedTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValSpeedType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValSpeedType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValSpeedType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValSpeedType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValSpeedType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValSpeedType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValSpeedType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValWeightTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValWeightType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValWeightType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValWeightType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValWeightType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValWeightType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValWeightType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValWeightType]() {
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

	// Package node: aixmabstractgmlobjecttypes
	pkgNode_AixmabstractgmlobjecttypesStage := &tree_models.Node{
		Name:       "aixmabstractgmlobjecttypes",
		IsExpanded: true,
	}
	pkgNode_AixmabstractgmlobjecttypesStage.OnIsExpandedChange = func(isExpanded bool) {
		pkgNode_AixmabstractgmlobjecttypesStage.IsExpanded = isExpanded
	}
	topNode.Children = append(topNode.Children, pkgNode_AixmabstractgmlobjecttypesStage)

	{
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureBaseTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractAIXMFeatureBaseType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractAIXMFeatureBaseType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractAIXMFeatureBaseType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractAIXMFeatureBaseType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractAIXMFeatureBaseType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMFeatureTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractAIXMFeatureType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractAIXMFeatureType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractAIXMFeatureType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractAIXMFeatureType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractAIXMFeatureType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMFeatureType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageBaseTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractAIXMMessageBaseType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractAIXMMessageBaseType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractAIXMMessageBaseType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractAIXMMessageBaseType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractAIXMMessageBaseType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMMessageTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractAIXMMessageType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractAIXMMessageType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractAIXMMessageType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractAIXMMessageType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractAIXMMessageType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMMessageType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMObjectTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractAIXMObjectType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractAIXMObjectType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractAIXMObjectType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractAIXMObjectType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractAIXMObjectType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMObjectType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractAIXMPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractAIXMPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractAIXMPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractAIXMPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractAIXMPropertyType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceBaseTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractAIXMTimeSliceBaseType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractAIXMTimeSliceBaseType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractAIXMTimeSliceBaseType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractAIXMTimeSliceBaseType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractAIXMTimeSliceBaseType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractAIXMTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractAIXMTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractAIXMTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractAIXMTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractAIXMTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractAIXMTimeSliceType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.AbstractExtensionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractExtensionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractExtensionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractExtensionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractExtensionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractExtensionType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractExtensionType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.AbstractExtensionType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.FeatureMetadataPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FeatureMetadataPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FeatureMetadataPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FeatureMetadataPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FeatureMetadataPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FeatureMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.FeatureMetadataPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.FeatureTimeSliceMetadataPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FeatureTimeSliceMetadataPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FeatureTimeSliceMetadataPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FeatureTimeSliceMetadataPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FeatureTimeSliceMetadataPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FeatureTimeSliceMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmabstractgmlobjecttypesStage.MessageMetadataPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("MessageMetadataPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all MessageMetadataPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmabstractgmlobjecttypesStage.Children = append(pkgNode_AixmabstractgmlobjecttypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "MessageMetadataPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of MessageMetadataPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_MessageMetadataPropertyType_AixmabstractgmlobjecttypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmabstractgmlobjecttypesStage.GetInstancesByOrder[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType]() {
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

	// Package node: gml
	pkgNode_GmlStage := &tree_models.Node{
		Name:       "gml",
		IsExpanded: true,
	}
	pkgNode_GmlStage.OnIsExpandedChange = func(isExpanded bool) {
		pkgNode_GmlStage.IsExpanded = isExpanded
	}
	topNode.Children = append(topNode.Children, pkgNode_GmlStage)

	{
		count := len(probe.stageOfInterest.GmlStage.AbstractCurveSegmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractCurveSegmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractCurveSegmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractCurveSegmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractCurveSegmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractCurveSegmentType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractCurveSegmentType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractCurveSegmentType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractCurveTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractCurveType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractCurveType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractCurveType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractCurveType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractCurveType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractCurveType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractCurveType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractFeatureMemberTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractFeatureMemberType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractFeatureMemberType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractFeatureMemberType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractFeatureMemberType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractFeatureMemberType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractFeatureMemberType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractFeatureMemberType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractFeatureTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractFeatureType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractFeatureType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractFeatureType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractFeatureType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractFeatureType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractFeatureType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractFeatureType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractGMLTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractGMLType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractGMLType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractGMLType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractGMLType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractGMLType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractGMLType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractGMLType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractGeometricPrimitiveTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractGeometricPrimitiveType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractGeometricPrimitiveType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractGeometricPrimitiveType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractGeometricPrimitiveType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractGeometricPrimitiveType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractGeometricPrimitiveType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractGeometricPrimitiveType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractGeometryTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractGeometryType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractGeometryType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractGeometryType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractGeometryType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractGeometryType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractGeometryType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractGeometryType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractMetadataPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractMetadataPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractMetadataPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractMetadataPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractMetadataPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractMetadataPropertyType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractMetadataPropertyType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractMetadataPropertyType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractRingPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractRingPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractRingPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractRingPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractRingPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractRingPropertyType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractRingPropertyType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractRingPropertyType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractRingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractRingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractRingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractRingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractRingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractRingType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractRingType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractRingType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractSurfacePatchTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractSurfacePatchType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractSurfacePatchType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractSurfacePatchType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractSurfacePatchType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractSurfacePatchType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractSurfacePatchType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractSurfacePatchType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractSurfaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractSurfaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractSurfaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractSurfaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractSurfaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractSurfaceType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractSurfaceType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractSurfaceType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractTimeGeometricPrimitiveTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractTimeGeometricPrimitiveType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractTimeGeometricPrimitiveType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractTimeGeometricPrimitiveType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractTimeGeometricPrimitiveType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractTimeGeometricPrimitiveType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractTimeGeometricPrimitiveType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractTimeGeometricPrimitiveType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractTimeObjectTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractTimeObjectType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractTimeObjectType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractTimeObjectType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractTimeObjectType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractTimeObjectType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractTimeObjectType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractTimeObjectType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractTimePrimitiveTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractTimePrimitiveType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractTimePrimitiveType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractTimePrimitiveType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractTimePrimitiveType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractTimePrimitiveType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractTimePrimitiveType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractTimePrimitiveType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AbstractTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractTimeSliceType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractTimeSliceType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AbstractTimeSliceType]() {
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
		count := len(probe.stageOfInterest.GmlStage.AngleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AngleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AngleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AngleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AngleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AngleType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AngleType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.AngleType]() {
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
		count := len(probe.stageOfInterest.GmlStage.ArcByCenterPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ArcByCenterPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ArcByCenterPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ArcByCenterPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ArcByCenterPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ArcByCenterPointType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ArcByCenterPointType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.ArcByCenterPointType]() {
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
		count := len(probe.stageOfInterest.GmlStage.ArcStringTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ArcStringType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ArcStringType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ArcStringType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ArcStringType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ArcStringType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ArcStringType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.ArcStringType]() {
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
		count := len(probe.stageOfInterest.GmlStage.ArcTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ArcType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ArcType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ArcType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ArcType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ArcType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ArcType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.ArcType]() {
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
		count := len(probe.stageOfInterest.GmlStage.BoundingShapeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("BoundingShapeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all BoundingShapeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "BoundingShapeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of BoundingShapeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_BoundingShapeType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_BoundingShapeType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.BoundingShapeType]() {
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
		count := len(probe.stageOfInterest.GmlStage.CircleByCenterPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CircleByCenterPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CircleByCenterPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CircleByCenterPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CircleByCenterPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CircleByCenterPointType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CircleByCenterPointType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CircleByCenterPointType]() {
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
		count := len(probe.stageOfInterest.GmlStage.CircleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CircleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CircleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CircleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CircleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CircleType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CircleType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CircleType]() {
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
		count := len(probe.stageOfInterest.GmlStage.CodeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CodeType]() {
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
		count := len(probe.stageOfInterest.GmlStage.CodeWithAuthorityTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeWithAuthorityType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeWithAuthorityType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeWithAuthorityType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeWithAuthorityType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeWithAuthorityType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeWithAuthorityType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CodeWithAuthorityType]() {
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
		count := len(probe.stageOfInterest.GmlStage.CompositeCurveTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CompositeCurveType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CompositeCurveType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CompositeCurveType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CompositeCurveType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CompositeCurveType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CompositeCurveType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CompositeCurveType]() {
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
		count := len(probe.stageOfInterest.GmlStage.CurvePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CurvePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CurvePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CurvePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CurvePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CurvePropertyType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CurvePropertyType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CurvePropertyType]() {
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
		count := len(probe.stageOfInterest.GmlStage.CurveSegmentArrayPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CurveSegmentArrayPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CurveSegmentArrayPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CurveSegmentArrayPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CurveSegmentArrayPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CurveSegmentArrayPropertyType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CurveSegmentArrayPropertyType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CurveSegmentArrayPropertyType]() {
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
		count := len(probe.stageOfInterest.GmlStage.CurveTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CurveType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CurveType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CurveType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CurveType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CurveType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CurveType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.CurveType]() {
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
		count := len(probe.stageOfInterest.GmlStage.DirectPositionListTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DirectPositionListType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DirectPositionListType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DirectPositionListType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DirectPositionListType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DirectPositionListType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DirectPositionListType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.DirectPositionListType]() {
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
		count := len(probe.stageOfInterest.GmlStage.DirectPositionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DirectPositionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DirectPositionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DirectPositionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DirectPositionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DirectPositionType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DirectPositionType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.DirectPositionType]() {
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
		count := len(probe.stageOfInterest.GmlStage.DynamicFeatureTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DynamicFeatureType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DynamicFeatureType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DynamicFeatureType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DynamicFeatureType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DynamicFeatureType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DynamicFeatureType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.DynamicFeatureType]() {
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
		count := len(probe.stageOfInterest.GmlStage.EnvelopeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("EnvelopeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all EnvelopeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "EnvelopeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of EnvelopeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_EnvelopeType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_EnvelopeType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.EnvelopeType]() {
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
		count := len(probe.stageOfInterest.GmlStage.GeodesicStringTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("GeodesicStringType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all GeodesicStringType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "GeodesicStringType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of GeodesicStringType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_GeodesicStringType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_GeodesicStringType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.GeodesicStringType]() {
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
		count := len(probe.stageOfInterest.GmlStage.GeodesicTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("GeodesicType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all GeodesicType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "GeodesicType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of GeodesicType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_GeodesicType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_GeodesicType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.GeodesicType]() {
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
		count := len(probe.stageOfInterest.GmlStage.LengthTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("LengthType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all LengthType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "LengthType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of LengthType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_LengthType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_LengthType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.LengthType]() {
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
		count := len(probe.stageOfInterest.GmlStage.LineStringSegmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("LineStringSegmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all LineStringSegmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "LineStringSegmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of LineStringSegmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_LineStringSegmentType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_LineStringSegmentType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.LineStringSegmentType]() {
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
		count := len(probe.stageOfInterest.GmlStage.MeasureTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("MeasureType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all MeasureType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "MeasureType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of MeasureType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_MeasureType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_MeasureType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.MeasureType]() {
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
		count := len(probe.stageOfInterest.GmlStage.OrientableCurveTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("OrientableCurveType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all OrientableCurveType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "OrientableCurveType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of OrientableCurveType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_OrientableCurveType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_OrientableCurveType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.OrientableCurveType]() {
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
		count := len(probe.stageOfInterest.GmlStage.PointPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PointPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PointPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PointPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PointPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PointPropertyType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PointPropertyType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.PointPropertyType]() {
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
		count := len(probe.stageOfInterest.GmlStage.PointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PointType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PointType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.PointType]() {
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
		count := len(probe.stageOfInterest.GmlStage.PolygonPatchTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PolygonPatchType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PolygonPatchType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PolygonPatchType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PolygonPatchType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PolygonPatchType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PolygonPatchType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.PolygonPatchType]() {
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
		count := len(probe.stageOfInterest.GmlStage.RingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RingType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RingType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.RingType]() {
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
		count := len(probe.stageOfInterest.GmlStage.SurfacePatchArrayPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SurfacePatchArrayPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SurfacePatchArrayPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SurfacePatchArrayPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SurfacePatchArrayPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SurfacePatchArrayPropertyType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SurfacePatchArrayPropertyType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.SurfacePatchArrayPropertyType]() {
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
		count := len(probe.stageOfInterest.GmlStage.SurfaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SurfaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SurfaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SurfaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SurfaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SurfaceType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SurfaceType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.SurfaceType]() {
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
		count := len(probe.stageOfInterest.GmlStage.TimeInstantTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TimeInstantType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TimeInstantType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TimeInstantType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TimeInstantType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TimeInstantType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TimeInstantType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.TimeInstantType]() {
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
		count := len(probe.stageOfInterest.GmlStage.TimePeriodTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TimePeriodType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TimePeriodType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TimePeriodType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TimePeriodType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TimePeriodType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TimePeriodType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.TimePeriodType]() {
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
		count := len(probe.stageOfInterest.GmlStage.TimePositionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TimePositionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TimePositionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TimePositionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TimePositionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TimePositionType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TimePositionType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.TimePositionType]() {
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
		count := len(probe.stageOfInterest.GmlStage.TimePrimitivePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TimePrimitivePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TimePrimitivePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_GmlStage.Children = append(pkgNode_GmlStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TimePrimitivePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TimePrimitivePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TimePrimitivePropertyType_GmlStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TimePrimitivePropertyType_GmlStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.GmlStage.GetInstancesByOrder[*gml.TimePrimitivePropertyType]() {
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

	// Package node: xlink
	pkgNode_XlinkStage := &tree_models.Node{
		Name:       "xlink",
		IsExpanded: true,
	}
	pkgNode_XlinkStage.OnIsExpandedChange = func(isExpanded bool) {
		pkgNode_XlinkStage.IsExpanded = isExpanded
	}
	topNode.Children = append(topNode.Children, pkgNode_XlinkStage)

	{
		count := len(probe.stageOfInterest.XlinkStage.Xlink_SchemaElements)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("Xlink_SchemaElement (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all Xlink_SchemaElement instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_XlinkStage.Children = append(pkgNode_XlinkStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "Xlink_SchemaElement " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of Xlink_SchemaElement",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_Xlink_SchemaElement_XlinkStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_Xlink_SchemaElement_XlinkStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.XlinkStage.GetInstancesByOrder[*xlink.Xlink_SchemaElement]() {
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
