// generated code - do not edit
package probe

import (
	"fmt"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/xlink"
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_15s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_15 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_15 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_15 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_15",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_15_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_15_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_15]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_16s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_16 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_16 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_16 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_16",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_16_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_16_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_16]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_17s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_17 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_17 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_17 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_17",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_17_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_17_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_17]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_18s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_18 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_18 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_18 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_18",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_18_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_18_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_18]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_19s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_19 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_19 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_19 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_19",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_19_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_19_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_19]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_20s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_20 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_20 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_20 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_20",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_20_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_20_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_20]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_21s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_21 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_21 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_21 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_21",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_21_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_21_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_21]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_22s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_22 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_22 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_22 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_22",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_22_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_22_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_22]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_23s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_23 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_23 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_23 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_23",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_23_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_23_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_23]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_24s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_24 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_24 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_24 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_24",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_24_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_24_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_24]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_25s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_25 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_25 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_25 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_25",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_25_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_25_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_25]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_26s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_26 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_26 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_26 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_26",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_26_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_26_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_26]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_27s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_27 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_27 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_27 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_27",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_27_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_27_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_27]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_28s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_28 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_28 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_28 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_28",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_28_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_28_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_28]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_29s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_29 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_29 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_29 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_29",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_29_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_29_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_29]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_30s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_30 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_30 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_30 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_30",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_30_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_30_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_30]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_31s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_31 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_31 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_31 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_31",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_31_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_31_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_31]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_32s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_32 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_32 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_32 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_32",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_32_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_32_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_32]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_33s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_33 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_33 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_33 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_33",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_33_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_33_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_33]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_34s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_34 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_34 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_34 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_34",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_34_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_34_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_34]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_35s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_35 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_35 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_35 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_35",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_35_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_35_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_35]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_36s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_36 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_36 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_36 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_36",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_36_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_36_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_36]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_37s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_37 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_37 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_37 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_37",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_37_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_37_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_37]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_38s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_38 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_38 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_38 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_38",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_38_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_38_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_38]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_39s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_39 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_39 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_39 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_39",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_39_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_39_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_39]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_40s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_40 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_40 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_40 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_40",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_40_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_40_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_40]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_41s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_41 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_41 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_41 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_41",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_41_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_41_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_41]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AbstractNavaidEquipmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractNavaidEquipmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractNavaidEquipmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractNavaidEquipmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractNavaidEquipmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractNavaidEquipmentType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractNavaidEquipmentType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AbstractNavaidEquipmentType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AbstractNavigationSystemCheckpointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractNavigationSystemCheckpointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractNavigationSystemCheckpointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractNavigationSystemCheckpointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractNavigationSystemCheckpointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractNavigationSystemCheckpointType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractNavigationSystemCheckpointType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AbstractNavigationSystemCheckpointType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AbstractRadarEquipmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractRadarEquipmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractRadarEquipmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractRadarEquipmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractRadarEquipmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractRadarEquipmentType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractRadarEquipmentType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AbstractRadarEquipmentType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AbstractSurveillanceRadarTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractSurveillanceRadarType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractSurveillanceRadarType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractSurveillanceRadarType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractSurveillanceRadarType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractSurveillanceRadarType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractSurveillanceRadarType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AbstractSurveillanceRadarType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirportHeliportPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirportHeliportPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirportHeliportPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirportHeliportPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirportHeliportPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirportHeliportPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirportHeliportPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirportHeliportPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AngleIndicationTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AngleIndicationTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AngleIndicationTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AngleIndicationTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AngleIndicationTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AngleIndicationTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AngleIndicationTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AngleIndicationTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AngleIndicationTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AngleIndicationTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AngleIndicationTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AngleIndicationTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AngleIndicationTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AngleIndicationTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AngleIndicationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AngleIndicationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AngleIndicationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AngleIndicationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AngleIndicationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AngleIndicationType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AngleIndicationType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AngleIndicationType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForNavaidEquipmentPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForNavaidEquipmentPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForNavaidEquipmentPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForNavaidEquipmentPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForNavaidEquipmentPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForNavaidEquipmentPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForNavaidEquipmentPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForNavaidEquipmentPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForNavaidEquipmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForNavaidEquipmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForNavaidEquipmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForNavaidEquipmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForNavaidEquipmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForNavaidEquipmentType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForNavaidEquipmentType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForNavaidEquipmentType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationStationPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForSpecialNavigationStationPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForSpecialNavigationStationPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForSpecialNavigationStationPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForSpecialNavigationStationPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForSpecialNavigationStationPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForSpecialNavigationStationPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForSpecialNavigationStationPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationStationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForSpecialNavigationStationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForSpecialNavigationStationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForSpecialNavigationStationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForSpecialNavigationStationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForSpecialNavigationStationType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForSpecialNavigationStationType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForSpecialNavigationStationType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationSystemPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForSpecialNavigationSystemPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForSpecialNavigationSystemPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForSpecialNavigationSystemPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForSpecialNavigationSystemPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForSpecialNavigationSystemPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForSpecialNavigationSystemPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForSpecialNavigationSystemPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForSpecialNavigationSystemTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForSpecialNavigationSystemType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForSpecialNavigationSystemType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForSpecialNavigationSystemType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForSpecialNavigationSystemType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForSpecialNavigationSystemType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForSpecialNavigationSystemType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForSpecialNavigationSystemType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AzimuthTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AzimuthTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AzimuthTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AzimuthTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AzimuthTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AzimuthTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AzimuthTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AzimuthTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AzimuthTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AzimuthTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AzimuthTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AzimuthTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AzimuthTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AzimuthTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AzimuthTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AzimuthTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AzimuthType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AzimuthType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AzimuthType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AzimuthType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AzimuthType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AzimuthType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AzimuthType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CheckpointINSTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CheckpointINSTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CheckpointINSTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CheckpointINSTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CheckpointINSTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CheckpointINSTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.CheckpointINSTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CheckpointINSTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CheckpointINSTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CheckpointINSTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CheckpointINSTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CheckpointINSTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CheckpointINSTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.CheckpointINSTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.CheckpointINSTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CheckpointINSType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CheckpointINSType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CheckpointINSType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CheckpointINSType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CheckpointINSType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CheckpointINSType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.CheckpointINSType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CheckpointVORTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CheckpointVORTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CheckpointVORTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CheckpointVORTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CheckpointVORTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CheckpointVORTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.CheckpointVORTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CheckpointVORTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CheckpointVORTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CheckpointVORTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CheckpointVORTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CheckpointVORTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CheckpointVORTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.CheckpointVORTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.CheckpointVORTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CheckpointVORType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CheckpointVORType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CheckpointVORType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CheckpointVORType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CheckpointVORType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CheckpointVORType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.CheckpointVORType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ContactInformationPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ContactInformationPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ContactInformationPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ContactInformationPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ContactInformationPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ContactInformationPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ContactInformationPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ContactInformationPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ContactInformationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ContactInformationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ContactInformationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ContactInformationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ContactInformationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ContactInformationType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ContactInformationType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ContactInformationType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DMETimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DMETimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DMETimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DMETimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DMETimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DMETimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DMETimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DMETimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DMETimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DMETimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DMETimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DMETimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DMETimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DMETimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DMETimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DMETimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DMETypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DMEType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DMEType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DMEType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DMEType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DMEType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DMEType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DMEType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DesignatedPointPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DesignatedPointPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DesignatedPointPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DesignatedPointPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DesignatedPointPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DesignatedPointPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DesignatedPointPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DesignatedPointPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DesignatedPointTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DesignatedPointTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DesignatedPointTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DesignatedPointTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DesignatedPointTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DesignatedPointTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DesignatedPointTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DesignatedPointTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DesignatedPointTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DesignatedPointTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DesignatedPointTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DesignatedPointTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DesignatedPointTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DesignatedPointTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DesignatedPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DesignatedPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DesignatedPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DesignatedPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DesignatedPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DesignatedPointType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DesignatedPointType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DesignatedPointType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DirectionFinderTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DirectionFinderTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DirectionFinderTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DirectionFinderTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DirectionFinderTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DirectionFinderTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DirectionFinderTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DirectionFinderTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DirectionFinderTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DirectionFinderTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DirectionFinderTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DirectionFinderTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DirectionFinderTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DirectionFinderTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DirectionFinderTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DirectionFinderType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DirectionFinderType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DirectionFinderType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DirectionFinderType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DirectionFinderType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DirectionFinderType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DirectionFinderType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DistanceIndicationTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DistanceIndicationTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DistanceIndicationTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DistanceIndicationTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DistanceIndicationTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DistanceIndicationTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DistanceIndicationTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DistanceIndicationTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DistanceIndicationTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DistanceIndicationTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DistanceIndicationTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DistanceIndicationTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DistanceIndicationTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DistanceIndicationTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DistanceIndicationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DistanceIndicationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DistanceIndicationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DistanceIndicationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DistanceIndicationType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DistanceIndicationType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DistanceIndicationType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ElevatedPointPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ElevatedPointPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ElevatedPointPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ElevatedPointPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ElevatedPointPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ElevatedPointPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ElevatedPointPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedPointPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ElevatedPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ElevatedPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ElevatedPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ElevatedPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ElevatedPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ElevatedPointType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ElevatedPointType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevatedPointType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ElevationTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ElevationTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ElevationTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ElevationTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ElevationTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ElevationTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevationTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ElevationTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ElevationTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ElevationTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ElevationTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ElevationTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ElevationTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ElevationTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevationTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ElevationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ElevationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ElevationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ElevationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ElevationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ElevationType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ElevationType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ElevationType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("GlidepathTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all GlidepathTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "GlidepathTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of GlidepathTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_GlidepathTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_GlidepathTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.GlidepathTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.GlidepathTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("GlidepathTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all GlidepathTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "GlidepathTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of GlidepathTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_GlidepathTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_GlidepathTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.GlidepathTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.GlidepathTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("GlidepathType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all GlidepathType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "GlidepathType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of GlidepathType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_GlidepathType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_GlidepathType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.GlidepathType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.InformationServicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("InformationServicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all InformationServicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "InformationServicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of InformationServicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_InformationServicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_InformationServicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.InformationServicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("LocalizerTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all LocalizerTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "LocalizerTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of LocalizerTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_LocalizerTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_LocalizerTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LocalizerTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.LocalizerTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("LocalizerTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all LocalizerTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "LocalizerTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of LocalizerTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_LocalizerTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_LocalizerTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LocalizerTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.LocalizerTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("LocalizerType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all LocalizerType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "LocalizerType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of LocalizerType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_LocalizerType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_LocalizerType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.LocalizerType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("MarkerBeaconTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all MarkerBeaconTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "MarkerBeaconTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of MarkerBeaconTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_MarkerBeaconTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_MarkerBeaconTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.MarkerBeaconTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("MarkerBeaconTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all MarkerBeaconTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "MarkerBeaconTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of MarkerBeaconTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_MarkerBeaconTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_MarkerBeaconTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.MarkerBeaconTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.MarkerBeaconTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("MarkerBeaconType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all MarkerBeaconType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "MarkerBeaconType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of MarkerBeaconType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_MarkerBeaconType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_MarkerBeaconType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.MarkerBeaconType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NDBTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NDBTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NDBTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NDBTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NDBTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NDBTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NDBTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NDBTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NDBTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NDBTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NDBTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NDBTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NDBTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NDBTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NDBTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NDBTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NDBTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NDBType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NDBType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NDBType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NDBType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NDBType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NDBType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NDBType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidComponentPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidComponentPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidComponentPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidComponentPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidComponentPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidComponentPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidComponentPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidComponentPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidComponentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidComponentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidComponentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidComponentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidComponentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidComponentType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidComponentType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidComponentType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidEquipmentMonitoringPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidEquipmentMonitoringPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidEquipmentMonitoringPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidEquipmentMonitoringPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidEquipmentMonitoringPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidEquipmentMonitoringPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidEquipmentMonitoringPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentMonitoringTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidEquipmentMonitoringType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidEquipmentMonitoringType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidEquipmentMonitoringType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidEquipmentMonitoringType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidEquipmentMonitoringType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidEquipmentMonitoringType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidEquipmentMonitoringType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidEquipmentPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidEquipmentPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidEquipmentPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidEquipmentPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidEquipmentPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidEquipmentPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidEquipmentPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidEquipmentPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidOperationalStatusPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidOperationalStatusPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidOperationalStatusPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidOperationalStatusPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidOperationalStatusPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidOperationalStatusPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidOperationalStatusPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidOperationalStatusTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidOperationalStatusType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidOperationalStatusType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidOperationalStatusType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidOperationalStatusType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidOperationalStatusType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidOperationalStatusType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidOperationalStatusType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.NavaidTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("NavaidType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all NavaidType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "NavaidType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of NavaidType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_NavaidType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_NavaidType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.NavaidType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.OnlineContactPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("OnlineContactPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all OnlineContactPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "OnlineContactPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of OnlineContactPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_OnlineContactPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_OnlineContactPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.OnlineContactPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.OnlineContactTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("OnlineContactType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all OnlineContactType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "OnlineContactType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of OnlineContactType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_OnlineContactType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_OnlineContactType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.OnlineContactType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PointPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PointPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PointPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PointPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PointPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PointPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PointPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PointPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PointType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PointType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PointType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PostalAddressPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PostalAddressPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PostalAddressPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PostalAddressPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PostalAddressPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PostalAddressPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PostalAddressPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PostalAddressPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PostalAddressTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PostalAddressType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PostalAddressType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PostalAddressType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PostalAddressType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PostalAddressType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PostalAddressType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PostalAddressType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PrecisionApproachRadarTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PrecisionApproachRadarTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PrecisionApproachRadarTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PrecisionApproachRadarTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PrecisionApproachRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PrecisionApproachRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PrecisionApproachRadarTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PrecisionApproachRadarTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PrecisionApproachRadarTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PrecisionApproachRadarTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PrecisionApproachRadarTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PrecisionApproachRadarTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PrecisionApproachRadarTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PrecisionApproachRadarTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PrecisionApproachRadarTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PrecisionApproachRadarType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PrecisionApproachRadarType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PrecisionApproachRadarType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PrecisionApproachRadarType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PrecisionApproachRadarType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PrecisionApproachRadarType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PrecisionApproachRadarType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PrimarySurveillanceRadarTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PrimarySurveillanceRadarTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PrimarySurveillanceRadarTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PrimarySurveillanceRadarTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PrimarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PrimarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PrimarySurveillanceRadarTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PrimarySurveillanceRadarTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PrimarySurveillanceRadarTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PrimarySurveillanceRadarTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PrimarySurveillanceRadarTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PrimarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PrimarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PrimarySurveillanceRadarTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PrimarySurveillanceRadarTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PrimarySurveillanceRadarType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PrimarySurveillanceRadarType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PrimarySurveillanceRadarType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PrimarySurveillanceRadarType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PrimarySurveillanceRadarType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PrimarySurveillanceRadarType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PrimarySurveillanceRadarType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RadarComponentPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RadarComponentPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RadarComponentPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RadarComponentPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RadarComponentPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RadarComponentPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RadarComponentPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RadarComponentPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RadarComponentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RadarComponentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RadarComponentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RadarComponentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RadarComponentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RadarComponentType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RadarComponentType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RadarComponentType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RadarEquipmentPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RadarEquipmentPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RadarEquipmentPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RadarEquipmentPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RadarEquipmentPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RadarEquipmentPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RadarEquipmentPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RadarEquipmentPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RadarSystemTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RadarSystemTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RadarSystemTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RadarSystemTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RadarSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RadarSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RadarSystemTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RadarSystemTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RadarSystemTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RadarSystemTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RadarSystemTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RadarSystemTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RadarSystemTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RadarSystemTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RadarSystemTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RadarSystemTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RadarSystemType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RadarSystemType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RadarSystemType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RadarSystemType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RadarSystemType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RadarSystemType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RadarSystemType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ReflectorPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ReflectorPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ReflectorPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ReflectorPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ReflectorPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ReflectorPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ReflectorPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ReflectorPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ReflectorTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ReflectorType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ReflectorType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ReflectorType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ReflectorType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ReflectorType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ReflectorType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ReflectorType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RunwayCentrelinePointPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RunwayCentrelinePointPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RunwayCentrelinePointPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RunwayCentrelinePointPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RunwayCentrelinePointPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RunwayCentrelinePointPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RunwayCentrelinePointPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RunwayCentrelinePointPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RunwayDirectionPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RunwayDirectionPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RunwayDirectionPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RunwayDirectionPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RunwayDirectionPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RunwayDirectionPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RunwayDirectionPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RunwayDirectionPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RunwayPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RunwayPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RunwayPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RunwayPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RunwayPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RunwayPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RunwayPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RunwayPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SDFTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SDFTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SDFTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SDFTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SDFTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SDFTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SDFTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SDFTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SDFTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SDFTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SDFTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SDFTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SDFTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SDFTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SDFTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SDFTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SDFTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SDFType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SDFType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SDFType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SDFType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SDFType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SDFType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SDFType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SecondarySurveillanceRadarTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SecondarySurveillanceRadarTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SecondarySurveillanceRadarTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SecondarySurveillanceRadarTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SecondarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SecondarySurveillanceRadarTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SecondarySurveillanceRadarTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SecondarySurveillanceRadarTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SecondarySurveillanceRadarTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SecondarySurveillanceRadarTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SecondarySurveillanceRadarTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SecondarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SecondarySurveillanceRadarTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SecondarySurveillanceRadarTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SecondarySurveillanceRadarTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SecondarySurveillanceRadarType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SecondarySurveillanceRadarType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SecondarySurveillanceRadarType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SecondarySurveillanceRadarType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SecondarySurveillanceRadarType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SecondarySurveillanceRadarType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SecondarySurveillanceRadarType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SignificantPointInAirspaceTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SignificantPointInAirspaceTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SignificantPointInAirspaceTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SignificantPointInAirspaceTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SignificantPointInAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SignificantPointInAirspaceTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SignificantPointInAirspaceTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SignificantPointInAirspaceTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SignificantPointInAirspaceTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SignificantPointInAirspaceTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SignificantPointInAirspaceTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SignificantPointInAirspaceTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SignificantPointInAirspaceTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SignificantPointInAirspaceTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SignificantPointInAirspaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SignificantPointInAirspaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SignificantPointInAirspaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SignificantPointInAirspaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SignificantPointInAirspaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SignificantPointInAirspaceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SignificantPointInAirspaceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SignificantPointInAirspaceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SpecialNavigationStationStatusPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SpecialNavigationStationStatusPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SpecialNavigationStationStatusPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SpecialNavigationStationStatusPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SpecialNavigationStationStatusPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SpecialNavigationStationStatusPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationStationStatusPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationStatusTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SpecialNavigationStationStatusType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SpecialNavigationStationStatusType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SpecialNavigationStationStatusType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SpecialNavigationStationStatusType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SpecialNavigationStationStatusType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SpecialNavigationStationStatusType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationStationStatusType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SpecialNavigationStationTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SpecialNavigationStationTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SpecialNavigationStationTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SpecialNavigationStationTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SpecialNavigationStationTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SpecialNavigationStationTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationStationTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SpecialNavigationStationTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SpecialNavigationStationTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SpecialNavigationStationTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SpecialNavigationStationTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SpecialNavigationStationTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SpecialNavigationStationTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationStationTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationStationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SpecialNavigationStationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SpecialNavigationStationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SpecialNavigationStationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SpecialNavigationStationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SpecialNavigationStationType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SpecialNavigationStationType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationStationType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SpecialNavigationSystemPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SpecialNavigationSystemPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SpecialNavigationSystemPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SpecialNavigationSystemPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SpecialNavigationSystemPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SpecialNavigationSystemPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationSystemPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SpecialNavigationSystemTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SpecialNavigationSystemTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SpecialNavigationSystemTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SpecialNavigationSystemTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SpecialNavigationSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SpecialNavigationSystemTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationSystemTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SpecialNavigationSystemTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SpecialNavigationSystemTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SpecialNavigationSystemTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SpecialNavigationSystemTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SpecialNavigationSystemTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SpecialNavigationSystemTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationSystemTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SpecialNavigationSystemTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SpecialNavigationSystemType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SpecialNavigationSystemType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SpecialNavigationSystemType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SpecialNavigationSystemType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SpecialNavigationSystemType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SpecialNavigationSystemType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SpecialNavigationSystemType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SurveillanceGroundStationPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SurveillanceGroundStationPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SurveillanceGroundStationPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SurveillanceGroundStationPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SurveillanceGroundStationPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SurveillanceGroundStationPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SurveillanceGroundStationPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SurveillanceGroundStationPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.SurveillanceGroundStationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("SurveillanceGroundStationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all SurveillanceGroundStationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "SurveillanceGroundStationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of SurveillanceGroundStationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_SurveillanceGroundStationType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_SurveillanceGroundStationType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.SurveillanceGroundStationType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.TACANTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TACANTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TACANTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TACANTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TACANTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TACANTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TACANTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TACANTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.TACANTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TACANTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TACANTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TACANTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TACANTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TACANTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TACANTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TACANTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.TACANTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TACANType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TACANType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TACANType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TACANType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TACANType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TACANType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TACANType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.TelephoneContactPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TelephoneContactPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TelephoneContactPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TelephoneContactPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TelephoneContactPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TelephoneContactPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TelephoneContactPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TelephoneContactPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.TelephoneContactTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TelephoneContactType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TelephoneContactType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TelephoneContactType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TelephoneContactType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TelephoneContactType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TelephoneContactType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TelephoneContactType]() {
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

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.TouchDownLiftOffPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TouchDownLiftOffPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TouchDownLiftOffPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TouchDownLiftOffPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TouchDownLiftOffPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TouchDownLiftOffPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TouchDownLiftOffPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.TouchDownLiftOffPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.UnitPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("UnitPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all UnitPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "UnitPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of UnitPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_UnitPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_UnitPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.UnitPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.VORPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("VORPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all VORPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "VORPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of VORPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_VORPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_VORPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VORPropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.VORTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("VORTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all VORTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "VORTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of VORTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_VORTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_VORTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VORTimeSlicePropertyType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.VORTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("VORTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all VORTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "VORTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of VORTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_VORTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_VORTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VORTimeSliceType]() {
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.VORTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("VORType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all VORType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "VORType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of VORType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_VORType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_VORType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VORType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAirspacePointPositionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAirspacePointPositionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAirspacePointPositionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAirspacePointPositionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAirspacePointPositionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAirspacePointPositionType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAirspacePointPositionType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAirspacePointPositionType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAirspacePointRoleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAirspacePointRoleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAirspacePointRoleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAirspacePointRoleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAirspacePointRoleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAirspacePointRoleType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAirspacePointRoleType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAirspacePointRoleType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAuralMorseTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAuralMorseType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAuralMorseType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAuralMorseType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAuralMorseType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAuralMorseType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAuralMorseType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAuralMorseType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAuthorityRoleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAuthorityRoleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAuthorityRoleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAuthorityRoleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAuthorityRoleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAuthorityRoleType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAuthorityRoleType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAuthorityRoleType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeBearingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeBearingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeBearingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeBearingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeBearingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeBearingType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeBearingType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeBearingType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeCardinalDirectionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeCardinalDirectionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeCardinalDirectionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeCardinalDirectionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeCardinalDirectionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeCardinalDirectionType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeCardinalDirectionType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeCardinalDirectionType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeCheckpointCategoryTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeCheckpointCategoryType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeCheckpointCategoryType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeCheckpointCategoryType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeCheckpointCategoryType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeCheckpointCategoryType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeCheckpointCategoryType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeCheckpointCategoryType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeCourseQualityILSTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeCourseQualityILSType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeCourseQualityILSType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeCourseQualityILSType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeCourseQualityILSType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeCourseQualityILSType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeCourseQualityILSType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeCourseQualityILSType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeDMEChannelTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeDMEChannelType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeDMEChannelType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeDMEChannelType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeDMEChannelType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeDMEChannelType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeDMEChannelType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDMEChannelType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeDMETypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeDMEType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeDMEType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeDMEType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeDMEType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeDMEType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeDMEType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDMEType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeDesignatedPointDesignatorTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeDesignatedPointDesignatorType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeDesignatedPointDesignatorType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeDesignatedPointDesignatorType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeDesignatedPointDesignatorType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeDesignatedPointDesignatorType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeDesignatedPointDesignatorType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDesignatedPointDesignatorType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeDesignatedPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeDesignatedPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeDesignatedPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeDesignatedPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeDesignatedPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeDesignatedPointType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeDesignatedPointType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDesignatedPointType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeDirectionReferenceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeDirectionReferenceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeDirectionReferenceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeDirectionReferenceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeDirectionReferenceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeDirectionReferenceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeDirectionReferenceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDirectionReferenceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeDistanceIndicationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeDistanceIndicationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeDistanceIndicationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeDistanceIndicationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeDistanceIndicationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeDistanceIndicationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeDistanceIndicationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDistanceIndicationType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeEmissionBandTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeEmissionBandType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeEmissionBandType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeEmissionBandType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeEmissionBandType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeEmissionBandType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeEmissionBandType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeEmissionBandType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeILSBackCourseTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeILSBackCourseType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeILSBackCourseType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeILSBackCourseType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeILSBackCourseType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeILSBackCourseType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeILSBackCourseType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeILSBackCourseType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeIntegrityLevelILSTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeIntegrityLevelILSType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeIntegrityLevelILSType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeIntegrityLevelILSType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeIntegrityLevelILSType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeIntegrityLevelILSType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeIntegrityLevelILSType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeIntegrityLevelILSType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeMLSAzimuthTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeMLSAzimuthType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeMLSAzimuthType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeMLSAzimuthType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeMLSAzimuthType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeMLSAzimuthType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeMLSAzimuthType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeMLSAzimuthType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeMLSChannelTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeMLSChannelType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeMLSChannelType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeMLSChannelType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeMLSChannelType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeMLSChannelType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeMLSChannelType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeMLSChannelType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeMarkerBeaconSignalTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeMarkerBeaconSignalType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeMarkerBeaconSignalType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeMarkerBeaconSignalType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeMarkerBeaconSignalType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeMarkerBeaconSignalType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeMarkerBeaconSignalType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeMarkerBeaconSignalType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeNDBUsageTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeNDBUsageType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeNDBUsageType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeNDBUsageType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeNDBUsageType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeNDBUsageType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeNDBUsageType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeNDBUsageType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeNavaidDesignatorTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeNavaidDesignatorType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeNavaidDesignatorType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeNavaidDesignatorType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeNavaidDesignatorType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeNavaidDesignatorType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeNavaidDesignatorType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeNavaidDesignatorType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeNavaidPurposeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeNavaidPurposeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeNavaidPurposeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeNavaidPurposeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeNavaidPurposeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeNavaidPurposeType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeNavaidPurposeType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeNavaidPurposeType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeNavaidServiceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeNavaidServiceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeNavaidServiceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeNavaidServiceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeNavaidServiceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeNavaidServiceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeNavaidServiceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeNavaidServiceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeNorthReferenceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeNorthReferenceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeNorthReferenceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeNorthReferenceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeNorthReferenceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeNorthReferenceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeNorthReferenceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeNorthReferenceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodePARTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodePARType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodePARType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodePARType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodePARType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodePARType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodePARType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodePARType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodePositionInILSTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodePositionInILSType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodePositionInILSType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodePositionInILSType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodePositionInILSType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodePositionInILSType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodePositionInILSType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodePositionInILSType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodePrimaryRadarTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodePrimaryRadarType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodePrimaryRadarType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodePrimaryRadarType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodePrimaryRadarType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodePrimaryRadarType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodePrimaryRadarType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodePrimaryRadarType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRadarServiceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRadarServiceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRadarServiceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRadarServiceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRadarServiceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRadarServiceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRadarServiceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRadarServiceType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRadioEmissionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRadioEmissionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRadioEmissionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRadioEmissionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRadioEmissionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRadioEmissionType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRadioEmissionType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRadioEmissionType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRadioSignalTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRadioSignalType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRadioSignalType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRadioSignalType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRadioSignalType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRadioSignalType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRadioSignalType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRadioSignalType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeReflectorTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeReflectorType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeReflectorType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeReflectorType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeReflectorType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeReflectorType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeReflectorType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeReflectorType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeSignalPerformanceILSTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeSignalPerformanceILSType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeSignalPerformanceILSType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeSignalPerformanceILSType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeSignalPerformanceILSType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeSignalPerformanceILSType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeSignalPerformanceILSType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeSignalPerformanceILSType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeSpecialNavigationChainDesignatorTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeSpecialNavigationChainDesignatorType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeSpecialNavigationChainDesignatorType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeSpecialNavigationChainDesignatorType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeSpecialNavigationChainDesignatorType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeSpecialNavigationChainDesignatorType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeSpecialNavigationChainDesignatorType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeSpecialNavigationChainDesignatorType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeSpecialNavigationStationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeSpecialNavigationStationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeSpecialNavigationStationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeSpecialNavigationStationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeSpecialNavigationStationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeSpecialNavigationStationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeSpecialNavigationStationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeSpecialNavigationStationType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeSpecialNavigationSystemTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeSpecialNavigationSystemType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeSpecialNavigationSystemType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeSpecialNavigationSystemType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeSpecialNavigationSystemType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeSpecialNavigationSystemType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeSpecialNavigationSystemType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeSpecialNavigationSystemType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeStandbyPowerTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeStandbyPowerType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeStandbyPowerType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeStandbyPowerType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeStandbyPowerType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeStandbyPowerType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeStandbyPowerType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeStandbyPowerType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeStatusNavaidTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeStatusNavaidType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeStatusNavaidType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeStatusNavaidType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeStatusNavaidType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeStatusNavaidType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeStatusNavaidType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeStatusNavaidType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeTACANChannelTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeTACANChannelType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeTACANChannelType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeTACANChannelType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeTACANChannelType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeTACANChannelType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeTACANChannelType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeTACANChannelType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeTelecomNetworkTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeTelecomNetworkType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeTelecomNetworkType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeTelecomNetworkType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeTelecomNetworkType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeTelecomNetworkType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeTelecomNetworkType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeTelecomNetworkType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeVORTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeVORType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeVORType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeVORType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeVORType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeVORType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeVORType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeVORType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.DateYearTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DateYearType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DateYearType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DateYearType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DateYearType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DateYearType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DateYearType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.DateYearType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.TextAddressTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TextAddressType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TextAddressType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TextAddressType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TextAddressType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TextAddressType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TextAddressType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextAddressType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.TextDesignatorTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TextDesignatorType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TextDesignatorType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TextDesignatorType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TextDesignatorType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TextDesignatorType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TextDesignatorType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextDesignatorType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.TextPhoneTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TextPhoneType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TextPhoneType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TextPhoneType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TextPhoneType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TextPhoneType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TextPhoneType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextPhoneType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValAngleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValAngleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValAngleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValAngleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValAngleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValAngleType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValAngleType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValAngleType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValBearingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValBearingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValBearingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValBearingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValBearingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValBearingType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValBearingType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValBearingType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValFrequencyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValFrequencyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValFrequencyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValFrequencyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValFrequencyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValFrequencyType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValFrequencyType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValFrequencyType]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValMagneticVariationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValMagneticVariationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValMagneticVariationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValMagneticVariationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValMagneticVariationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValMagneticVariationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValMagneticVariationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValMagneticVariationType]() {
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
