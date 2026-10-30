// generated code - do not edit
package probe

import (
	"fmt"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/xlink"
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.A_extension_42s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("A_extension_42 (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all A_extension_42 instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "A_extension_42 " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of A_extension_42",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_A_extension_42_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_A_extension_42_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.A_extension_42]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AbstractDirectFlightTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractDirectFlightType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractDirectFlightType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractDirectFlightType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractDirectFlightType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractDirectFlightType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractDirectFlightType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AbstractDirectFlightType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AbstractSegmentPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AbstractSegmentPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AbstractSegmentPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AbstractSegmentPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AbstractSegmentPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AbstractSegmentPointType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AbstractSegmentPointType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AbstractSegmentPointType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingAnchorPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingAnchorPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingAnchorPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingAnchorPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingAnchorPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingAnchorPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingAnchorPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingAnchorTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingAnchorType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingAnchorType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingAnchorType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingAnchorType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingAnchorType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingAnchorType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingAnchorType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingPointPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingPointPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingPointPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingPointPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingPointPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingPointPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingPointPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingPointType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingPointType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingPointType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingTimeSlicePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingTimeSliceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingTrackPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingTrackPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingTrackPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingTrackPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingTrackPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingTrackPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingTrackPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTrackTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingTrackType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingTrackType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingTrackType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingTrackType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingTrackType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingTrackType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingTrackType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AerialRefuellingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AerialRefuellingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AerialRefuellingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AerialRefuellingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AerialRefuellingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AerialRefuellingType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AerialRefuellingType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AerialRefuellingType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceBorderCrossingPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceBorderCrossingPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceBorderCrossingPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceBorderCrossingPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceBorderCrossingPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceBorderCrossingPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceBorderCrossingPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceBorderCrossingTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceBorderCrossingTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceBorderCrossingTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceBorderCrossingTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceBorderCrossingTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceBorderCrossingTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceBorderCrossingTimeSlicePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceBorderCrossingTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceBorderCrossingTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceBorderCrossingTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceBorderCrossingTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceBorderCrossingTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceBorderCrossingTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceBorderCrossingTimeSliceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AirspaceBorderCrossingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AirspaceBorderCrossingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AirspaceBorderCrossingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AirspaceBorderCrossingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AirspaceBorderCrossingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AirspaceBorderCrossingType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AirspaceBorderCrossingType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AirspaceBorderCrossingType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AltitudeAdjustmentPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AltitudeAdjustmentPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AltitudeAdjustmentPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AltitudeAdjustmentPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AltitudeAdjustmentPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AltitudeAdjustmentPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AltitudeAdjustmentPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AltitudeAdjustmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AltitudeAdjustmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AltitudeAdjustmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AltitudeAdjustmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AltitudeAdjustmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AltitudeAdjustmentType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AltitudeAdjustmentType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AltitudeAdjustmentType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AngleIndicationPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AngleIndicationPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AngleIndicationPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AngleIndicationPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AngleIndicationPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AngleIndicationPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AngleIndicationPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AngleIndicationPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AngleUsePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AngleUsePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AngleUsePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AngleUsePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AngleUsePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AngleUsePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AngleUsePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AngleUsePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AngleUseTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AngleUseType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AngleUseType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AngleUseType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AngleUseType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AngleUseType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AngleUseType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AngleUseType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForAerialRefuellingPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForAerialRefuellingPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForAerialRefuellingPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForAerialRefuellingPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForAerialRefuellingPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForAerialRefuellingPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForAerialRefuellingPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForAerialRefuellingPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.AuthorityForAerialRefuellingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("AuthorityForAerialRefuellingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all AuthorityForAerialRefuellingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "AuthorityForAerialRefuellingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of AuthorityForAerialRefuellingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_AuthorityForAerialRefuellingType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_AuthorityForAerialRefuellingType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.AuthorityForAerialRefuellingType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ChangeOverPointTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ChangeOverPointTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ChangeOverPointTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ChangeOverPointTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ChangeOverPointTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ChangeOverPointTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ChangeOverPointTimeSlicePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ChangeOverPointTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ChangeOverPointTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ChangeOverPointTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ChangeOverPointTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ChangeOverPointTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ChangeOverPointTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ChangeOverPointTimeSliceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.ChangeOverPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ChangeOverPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ChangeOverPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ChangeOverPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ChangeOverPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ChangeOverPointType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ChangeOverPointType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ChangeOverPointType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DMEPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DMEPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DMEPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DMEPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DMEPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DMEPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DMEPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DMEPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.DirectFlightClassTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DirectFlightClassType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DirectFlightClassType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DirectFlightClassType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DirectFlightClassType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DirectFlightClassType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DirectFlightClassType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DirectFlightClassType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.DirectFlightPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DirectFlightPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DirectFlightPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DirectFlightPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DirectFlightPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DirectFlightPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DirectFlightPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DirectFlightPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DirectFlightSegmentPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DirectFlightSegmentPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DirectFlightSegmentPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DirectFlightSegmentPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DirectFlightSegmentPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DirectFlightSegmentPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DirectFlightSegmentPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.DirectFlightSegmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DirectFlightSegmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DirectFlightSegmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DirectFlightSegmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DirectFlightSegmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DirectFlightSegmentType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DirectFlightSegmentType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DirectFlightSegmentType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.DistanceIndicationPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("DistanceIndicationPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all DistanceIndicationPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "DistanceIndicationPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of DistanceIndicationPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_DistanceIndicationPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_DistanceIndicationPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.DistanceIndicationPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("EnRouteSegmentPointPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all EnRouteSegmentPointPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "EnRouteSegmentPointPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of EnRouteSegmentPointPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_EnRouteSegmentPointPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_EnRouteSegmentPointPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.EnRouteSegmentPointPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.EnRouteSegmentPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("EnRouteSegmentPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all EnRouteSegmentPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "EnRouteSegmentPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of EnRouteSegmentPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_EnRouteSegmentPointType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_EnRouteSegmentPointType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.EnRouteSegmentPointType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightCharacteristicPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightCharacteristicPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightCharacteristicPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightCharacteristicPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightCharacteristicPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightCharacteristicPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightCharacteristicPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightCharacteristicTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightCharacteristicType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightCharacteristicType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightCharacteristicType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightCharacteristicType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightCharacteristicType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightCharacteristicType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightCharacteristicType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightConditionCircumstancePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightConditionCircumstancePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightConditionCircumstancePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightConditionCircumstancePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightConditionCircumstancePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightConditionCircumstancePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightConditionCircumstancePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightConditionCircumstancePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightConditionCircumstanceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightConditionCircumstanceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightConditionCircumstanceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightConditionCircumstanceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightConditionCircumstanceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightConditionCircumstanceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightConditionCircumstanceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightConditionCircumstanceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightConditionCombinationPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightConditionCombinationPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightConditionCombinationPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightConditionCombinationPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightConditionCombinationPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightConditionCombinationPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightConditionCombinationPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightConditionCombinationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightConditionCombinationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightConditionCombinationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightConditionCombinationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightConditionCombinationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightConditionCombinationType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightConditionCombinationType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightConditionCombinationType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightConditionElementPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightConditionElementPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightConditionElementPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightConditionElementPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightConditionElementPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightConditionElementPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightConditionElementPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightConditionElementTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightConditionElementType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightConditionElementType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightConditionElementType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightConditionElementType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightConditionElementType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightConditionElementType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightConditionElementType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightRestrictionLevelPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightRestrictionLevelPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightRestrictionLevelPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightRestrictionLevelPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightRestrictionLevelPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightRestrictionLevelPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightRestrictionLevelPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionLevelTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightRestrictionLevelType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightRestrictionLevelType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightRestrictionLevelType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightRestrictionLevelType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightRestrictionLevelType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightRestrictionLevelType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightRestrictionLevelType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRoutePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightRestrictionRoutePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightRestrictionRoutePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightRestrictionRoutePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightRestrictionRoutePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightRestrictionRoutePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightRestrictionRoutePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightRestrictionRoutePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionRouteTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightRestrictionRouteType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightRestrictionRouteType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightRestrictionRouteType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightRestrictionRouteType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightRestrictionRouteType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightRestrictionRouteType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightRestrictionRouteType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightRestrictionTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightRestrictionTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightRestrictionTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightRestrictionTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightRestrictionTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightRestrictionTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightRestrictionTimeSlicePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightRestrictionTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightRestrictionTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightRestrictionTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightRestrictionTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightRestrictionTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightRestrictionTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightRestrictionTimeSliceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightRestrictionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightRestrictionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightRestrictionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightRestrictionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightRestrictionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightRestrictionType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightRestrictionType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightRestrictionType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightRoutingElementPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightRoutingElementPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightRoutingElementPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightRoutingElementPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightRoutingElementPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightRoutingElementPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightRoutingElementPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.FlightRoutingElementTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("FlightRoutingElementType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all FlightRoutingElementType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "FlightRoutingElementType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of FlightRoutingElementType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_FlightRoutingElementType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_FlightRoutingElementType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.FlightRoutingElementType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.MeteorologyPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("MeteorologyPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all MeteorologyPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "MeteorologyPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of MeteorologyPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_MeteorologyPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_MeteorologyPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.MeteorologyPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.MeteorologyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("MeteorologyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all MeteorologyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "MeteorologyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of MeteorologyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_MeteorologyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_MeteorologyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.MeteorologyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ObstacleAssessmentAreaPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ObstacleAssessmentAreaPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ObstacleAssessmentAreaPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ObstacleAssessmentAreaPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ObstacleAssessmentAreaPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ObstacleAssessmentAreaPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ObstacleAssessmentAreaPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.ObstacleAssessmentAreaTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ObstacleAssessmentAreaType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ObstacleAssessmentAreaType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ObstacleAssessmentAreaType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ObstacleAssessmentAreaType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ObstacleAssessmentAreaType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ObstacleAssessmentAreaType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ObstacleAssessmentAreaType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ObstaclePlacementPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ObstaclePlacementPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ObstaclePlacementPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ObstaclePlacementPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ObstaclePlacementPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ObstaclePlacementPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ObstaclePlacementPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.ObstaclePlacementTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ObstaclePlacementType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ObstaclePlacementType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ObstaclePlacementType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ObstaclePlacementType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ObstaclePlacementType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ObstaclePlacementType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ObstaclePlacementType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.ObstructionPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ObstructionPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ObstructionPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ObstructionPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ObstructionPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ObstructionPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ObstructionPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ObstructionPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.ObstructionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ObstructionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ObstructionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ObstructionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ObstructionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ObstructionType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ObstructionType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.ObstructionType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.PointReferencePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PointReferencePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PointReferencePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PointReferencePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PointReferencePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PointReferencePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PointReferencePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PointReferencePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.PointReferenceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("PointReferenceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all PointReferenceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "PointReferenceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of PointReferenceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_PointReferenceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_PointReferenceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.PointReferenceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RadioFrequencyAreaPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RadioFrequencyAreaPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RadioFrequencyAreaPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RadioFrequencyAreaPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RadioFrequencyAreaPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RadioFrequencyAreaPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RadioFrequencyAreaPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RadioFrequencyAreaPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteAvailabilityPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteAvailabilityPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteAvailabilityPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteAvailabilityPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteAvailabilityPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteAvailabilityPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteAvailabilityPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteAvailabilityTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteAvailabilityType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteAvailabilityType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteAvailabilityType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteAvailabilityType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteAvailabilityType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteAvailabilityType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteAvailabilityType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteDMETimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteDMETimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteDMETimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteDMETimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteDMETimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteDMETimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteDMETimeSlicePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteDMETimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteDMETimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteDMETimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteDMETimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteDMETimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteDMETimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteDMETimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteDMETimeSliceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteDMETypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteDMEType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteDMEType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteDMEType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteDMEType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteDMEType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteDMEType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteDMEType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RoutePortionPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RoutePortionPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RoutePortionPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RoutePortionPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RoutePortionPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RoutePortionPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RoutePortionPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RoutePortionPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RoutePortionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RoutePortionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RoutePortionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RoutePortionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RoutePortionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RoutePortionType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RoutePortionType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RoutePortionType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteSegmentTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteSegmentTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteSegmentTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteSegmentTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteSegmentTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteSegmentTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteSegmentTimeSlicePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteSegmentTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteSegmentTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteSegmentTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteSegmentTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteSegmentTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteSegmentTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteSegmentTimeSliceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteSegmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteSegmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteSegmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteSegmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteSegmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteSegmentType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteSegmentType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteSegmentType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteTimeSlicePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteTimeSliceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RouteTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RouteType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RouteType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RouteType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RouteType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RouteType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RouteType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RouteType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSlicePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RulesProceduresTimeSlicePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RulesProceduresTimeSlicePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RulesProceduresTimeSlicePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RulesProceduresTimeSlicePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RulesProceduresTimeSlicePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RulesProceduresTimeSlicePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RulesProceduresTimeSlicePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTimeSliceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RulesProceduresTimeSliceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RulesProceduresTimeSliceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RulesProceduresTimeSliceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RulesProceduresTimeSliceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RulesProceduresTimeSliceType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RulesProceduresTimeSliceType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RulesProceduresTimeSliceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.RulesProceduresTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("RulesProceduresType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all RulesProceduresType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "RulesProceduresType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of RulesProceduresType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_RulesProceduresType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_RulesProceduresType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.RulesProceduresType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.StandardInstrumentArrivalPropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("StandardInstrumentArrivalPropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all StandardInstrumentArrivalPropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "StandardInstrumentArrivalPropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of StandardInstrumentArrivalPropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_StandardInstrumentArrivalPropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_StandardInstrumentArrivalPropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.StandardInstrumentArrivalPropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmfeaturesStage.StandardInstrumentDeparturePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("StandardInstrumentDeparturePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all StandardInstrumentDeparturePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "StandardInstrumentDeparturePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of StandardInstrumentDeparturePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_StandardInstrumentDeparturePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_StandardInstrumentDeparturePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.StandardInstrumentDeparturePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmfeaturesStage.VerticalStructurePropertyTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("VerticalStructurePropertyType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all VerticalStructurePropertyType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmfeaturesStage.Children = append(pkgNode_AixmfeaturesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "VerticalStructurePropertyType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of VerticalStructurePropertyType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_VerticalStructurePropertyType_AixmfeaturesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_VerticalStructurePropertyType_AixmfeaturesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmfeaturesStage.GetInstancesByOrder[*aixmfeatures.VerticalStructurePropertyType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeATCReportingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeATCReportingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeATCReportingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeATCReportingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeATCReportingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeATCReportingType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeATCReportingType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeATCReportingType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAerialRefuellingPointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAerialRefuellingPointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAerialRefuellingPointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAerialRefuellingPointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAerialRefuellingPointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAerialRefuellingPointType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAerialRefuellingPointType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAerialRefuellingPointType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAerialRefuellingPrefixTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAerialRefuellingPrefixType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAerialRefuellingPrefixType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAerialRefuellingPrefixType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAerialRefuellingPrefixType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAerialRefuellingPrefixType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAerialRefuellingPrefixType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAerialRefuellingPrefixType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAerialRefuellingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAerialRefuellingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAerialRefuellingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAerialRefuellingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAerialRefuellingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAerialRefuellingType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAerialRefuellingType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAerialRefuellingType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeAltitudeAdjustmentTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeAltitudeAdjustmentType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeAltitudeAdjustmentType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeAltitudeAdjustmentType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeAltitudeAdjustmentType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeAltitudeAdjustmentType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeAltitudeAdjustmentType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeAltitudeAdjustmentType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeComparisonTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeComparisonType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeComparisonType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeComparisonType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeComparisonType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeComparisonType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeComparisonType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeComparisonType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeCourseTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeCourseType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeCourseType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeCourseType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeCourseType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeCourseType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeCourseType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeCourseType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeDirectionTurnTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeDirectionTurnType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeDirectionTurnType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeDirectionTurnType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeDirectionTurnType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeDirectionTurnType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeDirectionTurnType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDirectionTurnType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeDirectionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeDirectionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeDirectionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeDirectionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeDirectionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeDirectionType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeDirectionType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeDirectionType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeFlightOriginTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeFlightOriginType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeFlightOriginType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeFlightOriginType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeFlightOriginType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeFlightOriginType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeFlightOriginType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeFlightOriginType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeFlightPurposeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeFlightPurposeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeFlightPurposeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeFlightPurposeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeFlightPurposeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeFlightPurposeType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeFlightPurposeType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeFlightPurposeType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeFlightRestrictionDesignatorTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeFlightRestrictionDesignatorType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeFlightRestrictionDesignatorType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeFlightRestrictionDesignatorType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeFlightRestrictionDesignatorType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeFlightRestrictionDesignatorType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeFlightRestrictionDesignatorType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeFlightRestrictionDesignatorType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeFlightRestrictionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeFlightRestrictionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeFlightRestrictionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeFlightRestrictionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeFlightRestrictionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeFlightRestrictionType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeFlightRestrictionType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeFlightRestrictionType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeFlightRuleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeFlightRuleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeFlightRuleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeFlightRuleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeFlightRuleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeFlightRuleType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeFlightRuleType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeFlightRuleType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeFlightStatusTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeFlightStatusType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeFlightStatusType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeFlightStatusType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeFlightStatusType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeFlightStatusType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeFlightStatusType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeFlightStatusType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeFlightTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeFlightType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeFlightType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeFlightType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeFlightType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeFlightType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeFlightType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeFlightType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeFlowConditionOperationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeFlowConditionOperationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeFlowConditionOperationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeFlowConditionOperationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeFlowConditionOperationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeFlowConditionOperationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeFlowConditionOperationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeFlowConditionOperationType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeFreeFlightTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeFreeFlightType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeFreeFlightType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeFreeFlightType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeFreeFlightType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeFreeFlightType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeFreeFlightType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeFreeFlightType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeLevelTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeLevelType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeLevelType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeLevelType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeLevelType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeLevelType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeLevelType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeLevelType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeLocationQualifierTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeLocationQualifierType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeLocationQualifierType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeLocationQualifierType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeLocationQualifierType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeLocationQualifierType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeLocationQualifierType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeLocationQualifierType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeMeteoConditionsTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeMeteoConditionsType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeMeteoConditionsType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeMeteoConditionsType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeMeteoConditionsType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeMeteoConditionsType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeMeteoConditionsType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeMeteoConditionsType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeMilitaryRoutePointTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeMilitaryRoutePointType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeMilitaryRoutePointType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeMilitaryRoutePointType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeMilitaryRoutePointType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeMilitaryRoutePointType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeMilitaryRoutePointType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeMilitaryRoutePointType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeMilitaryStatusTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeMilitaryStatusType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeMilitaryStatusType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeMilitaryStatusType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeMilitaryStatusType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeMilitaryStatusType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeMilitaryStatusType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeMilitaryStatusType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeMilitaryTrainingTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeMilitaryTrainingType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeMilitaryTrainingType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeMilitaryTrainingType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeMilitaryTrainingType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeMilitaryTrainingType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeMilitaryTrainingType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeMilitaryTrainingType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeObstacleAssessmentSurfaceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeObstacleAssessmentSurfaceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeObstacleAssessmentSurfaceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeObstacleAssessmentSurfaceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeObstacleAssessmentSurfaceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeObstacleAssessmentSurfaceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeObstacleAssessmentSurfaceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeObstructionIdSurfaceZoneTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeObstructionIdSurfaceZoneType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeObstructionIdSurfaceZoneType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeObstructionIdSurfaceZoneType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeObstructionIdSurfaceZoneType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeObstructionIdSurfaceZoneType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeObstructionIdSurfaceZoneType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeObstructionIdSurfaceZoneType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRNPTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRNPType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRNPType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRNPType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRNPType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRNPType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRNPType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRNPType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRVSMPointRoleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRVSMPointRoleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRVSMPointRoleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRVSMPointRoleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRVSMPointRoleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRVSMPointRoleType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRVSMPointRoleType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRVSMPointRoleType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeReferenceRoleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeReferenceRoleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeReferenceRoleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeReferenceRoleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeReferenceRoleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeReferenceRoleType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeReferenceRoleType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeReferenceRoleType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRouteAvailabilityTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRouteAvailabilityType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRouteAvailabilityType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRouteAvailabilityType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRouteAvailabilityType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRouteAvailabilityType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRouteAvailabilityType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRouteAvailabilityType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRouteDesignatorLetterTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRouteDesignatorLetterType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRouteDesignatorLetterType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRouteDesignatorLetterType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRouteDesignatorLetterType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRouteDesignatorLetterType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRouteDesignatorLetterType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRouteDesignatorLetterType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRouteDesignatorPrefixTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRouteDesignatorPrefixType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRouteDesignatorPrefixType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRouteDesignatorPrefixType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRouteDesignatorPrefixType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRouteDesignatorPrefixType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRouteDesignatorPrefixType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRouteDesignatorPrefixType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRouteDesignatorSuffixTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRouteDesignatorSuffixType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRouteDesignatorSuffixType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRouteDesignatorSuffixType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRouteDesignatorSuffixType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRouteDesignatorSuffixType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRouteDesignatorSuffixType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRouteDesignatorSuffixType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRouteNavigationTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRouteNavigationType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRouteNavigationType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRouteNavigationType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRouteNavigationType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRouteNavigationType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRouteNavigationType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRouteNavigationType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRouteOriginTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRouteOriginType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRouteOriginType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRouteOriginType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRouteOriginType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRouteOriginType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRouteOriginType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRouteOriginType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRouteSegmentPathTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRouteSegmentPathType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRouteSegmentPathType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRouteSegmentPathType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRouteSegmentPathType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRouteSegmentPathType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRouteSegmentPathType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRouteSegmentPathType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRouteTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRouteType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRouteType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRouteType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRouteType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRouteType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRouteType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRouteType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRuleProcedureTitleTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRuleProcedureTitleType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRuleProcedureTitleType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRuleProcedureTitleType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRuleProcedureTitleType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRuleProcedureTitleType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRuleProcedureTitleType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRuleProcedureTitleType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeRuleProcedureTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeRuleProcedureType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeRuleProcedureType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeRuleProcedureType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeRuleProcedureType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeRuleProcedureType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeRuleProcedureType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeRuleProcedureType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeSideTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeSideType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeSideType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeSideType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeSideType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeSideType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeSideType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeSideType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeSpeedReferenceTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeSpeedReferenceType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeSpeedReferenceType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeSpeedReferenceType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeSpeedReferenceType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeSpeedReferenceType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeSpeedReferenceType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeSpeedReferenceType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.CodeUpperAlphaTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("CodeUpperAlphaType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all CodeUpperAlphaType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "CodeUpperAlphaType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of CodeUpperAlphaType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_CodeUpperAlphaType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_CodeUpperAlphaType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.CodeUpperAlphaType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.TextInstructionTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("TextInstructionType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all TextInstructionType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "TextInstructionType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of TextInstructionType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_TextInstructionType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_TextInstructionType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.TextInstructionType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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
		count := len(probe.stageOfInterest.AixmdatatypesStage.ValSlopeTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("ValSlopeType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all ValSlopeType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "ValSlopeType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of ValSlopeType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_ValSlopeType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_ValSlopeType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.ValSlopeType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
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

	{
		count := len(probe.stageOfInterest.AixmdatatypesStage.XHTMLTypes)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("XHTMLType (%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all XHTMLType instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_AixmdatatypesStage.Children = append(pkgNode_AixmdatatypesStage.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "XHTMLType " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of XHTMLType",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_XHTMLType_AixmdatatypesStage(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_XHTMLType_AixmdatatypesStage(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _, _inst := range probe.stageOfInterest.AixmdatatypesStage.GetInstancesByOrder[*aixmdatatypes.XHTMLType]() {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					FillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
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
