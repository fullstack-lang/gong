package diagrammer

import (
	"log"

	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/maticons/maticons"
)

type Diagrammer struct {
	model     Model
	portfolio Portfolio
	treeStage *gongtree_models.StageStruct

	map_portfolioNode_treeNode    map[PortfolioNode]*gongtree_models.Node
	map_modelElementNode_treeNode map[ModelElementNode]*gongtree_models.Node
	map_modelElementNode_Shape    map[ModelElementNode]Shape
	map_modelElements_Shape       map[any]Shape
}

func NewDiagrammer(
	model Model,
	portfolio Portfolio,
	treeStage *gongtree_models.StageStruct,

) (diagrammer *Diagrammer) {
	diagrammer = new(Diagrammer)

	diagrammer.model = model
	diagrammer.portfolio = portfolio
	diagrammer.treeStage = treeStage

	diagrammer.map_portfolioNode_treeNode = make(map[PortfolioNode]*gongtree_models.Node)
	diagrammer.map_modelElementNode_treeNode = make(map[ModelElementNode]*gongtree_models.Node)

	return
}

func (diagrammer *Diagrammer) AddPortfiolioNodeTreeNodeEntry(portfolioNode PortfolioNode, treeNode *gongtree_models.Node) {
	diagrammer.map_portfolioNode_treeNode[portfolioNode] = treeNode
}

func (diagrammer *Diagrammer) RemovePortfiolioNodeTreeNodeEntry(portfolioNode PortfolioNode) {
	delete(diagrammer.map_portfolioNode_treeNode, portfolioNode)
}

func (diagrammer *Diagrammer) GetPortfiolioNodeFromTreeNode(portfolioNode PortfolioNode) (treeNode *gongtree_models.Node) {

	var ok bool
	treeNode, ok = diagrammer.map_portfolioNode_treeNode[portfolioNode]
	if !ok {
		log.Println("unknown node", portfolioNode.GetName())
		for tn, pn := range diagrammer.map_portfolioNode_treeNode {
			log.Println(tn.GetName(), pn.GetName())
			log.Printf("%p %p\n", pn, tn)
		}
		log.Printf("entry node %s, %p\n", portfolioNode.GetName(), portfolioNode)

		return
	} else {
		return
	}
}

// FillUpModelTree ranges over Model Root Nodes
// and recursively fill up the Model Tree
func (diagrammer *Diagrammer) FillUpModelTree(modelTree *gongtree_models.Tree) {
	diagrammer.model.GenerateProgeny()

	for _, node := range diagrammer.model.GetChildren() {
		treeNode := diagrammer.modelNode2ModelTreeNode(node, diagrammer.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
	}
}

func (diagrammer *Diagrammer) modelNode2ModelTreeNode(modelNode ModelNode, treeStage *gongtree_models.StageStruct) (
	modelTreeNode *gongtree_models.Node) {
	modelTreeNode = (&gongtree_models.Node{Name: modelNode.GetName()}).Stage(treeStage)
	modelTreeNode.IsExpanded = modelNode.IsExpanded()

	if elementNode, ok := modelNode.(ModelElementNode); ok {
		modelTreeNode.HasCheckboxButton = true
		modelTreeNode.IsCheckboxDisabled = true
		diagrammer.map_modelElementNode_treeNode[elementNode] = modelTreeNode
	}

	modelTreeNode.Impl = &ModelNodeImpl{
		diagrammer: diagrammer,
		modelNode:  modelNode}

	for _, childrenModelNode := range modelNode.GetChildren() {
		childrenTreeNode := diagrammer.modelNode2ModelTreeNode(childrenModelNode, treeStage)
		modelTreeNode.Children = append(modelTreeNode.Children, childrenTreeNode)
	}

	return
}

// FillUpPortfolioUITree ranges over Portfolio Root Nodes
// and recursively fill up the Tree UI from the Portfolio tree
func (diagrammer *Diagrammer) FillUpPortfolioUITree(portfolioUITree *gongtree_models.Tree) {

	for _, portfolioNode := range diagrammer.portfolio.GeneratesProgeny() {
		// log.Printf("FillUpPortfolioTree %s %p\n", portfolioNode.GetName(), portfolioNode)
		treeNode := diagrammer.portfolioNode2NodeTree(portfolioNode, diagrammer.treeStage)
		portfolioUITree.RootNodes = append(portfolioUITree.RootNodes, treeNode)
	}
	diagrammer.generatePortfolioNodesStatusAndButtons()
}

func (diagrammer *Diagrammer) portfolioNode2NodeTree(portfolioNode PortfolioNode, treeStage *gongtree_models.StageStruct) (portfolioTreeNode *gongtree_models.Node) {
	portfolioTreeNode = (&gongtree_models.Node{Name: portfolioNode.GetName()}).Stage(treeStage)
	portfolioTreeNode.IsExpanded = portfolioNode.IsExpanded()

	// log.Printf("portfolioNode2NodeTree %s %p\n", portfolioNode.GetName(), portfolioNode)
	diagrammer.AddPortfiolioNodeTreeNodeEntry(portfolioNode, portfolioTreeNode)

	if diagramNode, ok := portfolioNode.(PortfolioDiagramNode); ok {
		_ = diagramNode
		portfolioTreeNode.HasCheckboxButton = true
		portfolioTreeNode.Impl = &PortfolioDiagramNodeImpl{
			diagrammer:           diagrammer,
			portfolioDiagramNode: diagramNode}
	}

	portfolioTreeNode.IsCheckboxDisabled = !diagrammer.portfolio.IsInSelectionMode()

	for _, childrenPortfolioNode := range portfolioNode.GetChildren() {
		childrenTreeNode := diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, treeStage)
		portfolioTreeNode.Children = append(portfolioTreeNode.Children, childrenTreeNode)
	}

	return
}

func (diagrammer *Diagrammer) generatePortfolioNodesStatusAndButtons() {

	for _, portfolioNode := range diagrammer.portfolio.GetChildren() {
		// log.Printf("generatePortfolioNodesButtons %s %p\n", portfolioNode.GetName(), portfolioNode)
		diagrammer.generatePortfolioNodesStatusAndButtonsRecursive(portfolioNode)
	}
}

func (diagrammer *Diagrammer) generatePortfolioNodesStatusAndButtonsRecursive(portfolioNode PortfolioNode) {

	// log.Printf("generatePortfolioNodesButtonsRecursive %s %p\n", portfolioNode.GetName(), portfolioNode)

	treeNode := diagrammer.GetPortfiolioNodeFromTreeNode(portfolioNode)

	// remove all buttons
	for _, _button := range treeNode.Buttons {
		_button.Unstage(diagrammer.treeStage)
	}
	treeNode.Buttons = make([]*gongtree_models.Button, 0)

	if portfolioCategoryNode, ok := portfolioNode.(PortfolioCategoryNode); ok {

		if portfolioCategoryNode.HasAddDiagramButton() {
			addDiagramButton := (&gongtree_models.Button{
				Name: portfolioNode.GetName() + " " + string(maticons.BUTTON_add),
				Icon: string(maticons.BUTTON_add)}).Stage(diagrammer.treeStage)
			treeNode.Buttons = append(treeNode.Buttons, addDiagramButton)
			addDiagramButton.Impl = NewDiagramButtonAddImpl(
				portfolioCategoryNode,
				diagrammer,
				treeNode,
				diagrammer.treeStage,
			)
		}
	}

	if portfolioDiagramNode, ok := portfolioNode.(PortfolioDiagramNode); ok {

		// disable the check button if the end user edit the diagram
		treeNode.IsCheckboxDisabled = diagrammer.portfolio.IsInDrawingMode()

		if portfolioDiagramNode == diagrammer.portfolio.GetSelectedPortfolioDiagramNode() {
			treeNode.IsChecked = true

			type nodeStateEnum int

			const (
				IDLE nodeStateEnum = iota
				IN_RENAME_MODE
				IN_EDIT_MODE
			)
			nodeState := IDLE

			if portfolioDiagramNode.IsInDrawingMode() {
				nodeState = IN_EDIT_MODE
			} else {
				if portfolioDiagramNode.IsInRenameMode() {
					nodeState = IN_RENAME_MODE
				}
			}

			switch nodeState {
			case IDLE:
				if portfolioDiagramNode.HasDrawButton() {
					editDiagramButton := (&gongtree_models.Button{
						Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_draw),
						Icon: string(maticons.BUTTON_draw)}).Stage(diagrammer.treeStage)
					treeNode.Buttons = append(treeNode.Buttons, editDiagramButton)
					editDiagramButton.Impl = NewPortfolioDiagramNodeButtonImpl(
						portfolioDiagramNode,
						diagrammer,
						treeNode,
						diagrammer.treeStage,
						EDIT,
					)
				}
				if portfolioDiagramNode.HasDiagramRenameButton() {
					renameDiagramButton := (&gongtree_models.Button{
						Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_edit),
						Icon: string(maticons.BUTTON_edit)}).Stage(diagrammer.treeStage)
					treeNode.Buttons = append(treeNode.Buttons, renameDiagramButton)
					renameDiagramButton.Impl = NewPortfolioDiagramNodeButtonImpl(
						portfolioDiagramNode,
						diagrammer,
						treeNode,
						diagrammer.treeStage,
						RENAME,
					)
				}
				if portfolioDiagramNode.HasDuplicateButton() {
					button := (&gongtree_models.Button{
						Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_file_copy),
						Icon: string(maticons.BUTTON_file_copy)}).Stage(diagrammer.treeStage)
					treeNode.Buttons = append(treeNode.Buttons, button)
					button.Impl = NewPortfolioDiagramNodeButtonImpl(
						portfolioDiagramNode,
						diagrammer,
						treeNode,
						diagrammer.treeStage,
						DUPLICATE,
					)
				}
				if portfolioDiagramNode.HasDeleteButton() {
					button := (&gongtree_models.Button{
						Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_delete),
						Icon: string(maticons.BUTTON_delete)}).Stage(diagrammer.treeStage)
					treeNode.Buttons = append(treeNode.Buttons, button)
					button.Impl = NewPortfolioDiagramNodeButtonImpl(
						portfolioDiagramNode,
						diagrammer,
						treeNode,
						diagrammer.treeStage,
						REMOVE,
					)
				}
			case IN_EDIT_MODE:
				saveDiagramButton := (&gongtree_models.Button{
					Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_save),
					Icon: string(maticons.BUTTON_save)}).Stage(diagrammer.treeStage)
				treeNode.Buttons = append(treeNode.Buttons, saveDiagramButton)
				saveDiagramButton.Impl = NewPortfolioDiagramNodeButtonImpl(
					portfolioDiagramNode,
					diagrammer,
					treeNode,
					diagrammer.treeStage,
					SAVE,
				)
				editCancelDiagramButton := (&gongtree_models.Button{
					Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_edit_off),
					Icon: string(maticons.BUTTON_edit_off)}).Stage(diagrammer.treeStage)
				treeNode.Buttons = append(treeNode.Buttons, editCancelDiagramButton)
				editCancelDiagramButton.Impl = NewPortfolioDiagramNodeButtonImpl(
					portfolioDiagramNode,
					diagrammer,
					treeNode,
					diagrammer.treeStage,
					EDIT_CANCEL,
				)
			case IN_RENAME_MODE:
				renameCancelDiagramButton := (&gongtree_models.Button{
					Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_edit_off),
					Icon: string(maticons.BUTTON_edit_off)}).Stage(diagrammer.treeStage)
				treeNode.Buttons = append(treeNode.Buttons, renameCancelDiagramButton)
				renameCancelDiagramButton.Impl = NewPortfolioDiagramNodeButtonImpl(
					portfolioDiagramNode,
					diagrammer,
					treeNode,
					diagrammer.treeStage,
					RENAME_CANCEL,
				)
			}
		} else {
			treeNode.IsChecked = false
		}
	}

	for _, childrenPortfolioNode := range portfolioNode.GetChildren() {
		diagrammer.generatePortfolioNodesStatusAndButtonsRecursive(childrenPortfolioNode)
	}
}

func (diagrammer *Diagrammer) CommitTreeStage() {
	diagrammer.treeStage.Commit()
}

func (diagrammer *Diagrammer) GetMap_elementNode_treeNode() map[ModelElementNode]*gongtree_models.Node {
	return diagrammer.map_modelElementNode_treeNode
}

// computeModelNodeStatus parses all nodes in the Model Tree
// and unchecks all nodes unless nodes matches an element in the diagram
func (diagrammer *Diagrammer) computeModelNodeStatus(map_ModelElementNode_Shape map[ModelElementNode]Shape) {
	isInDrawingMode := diagrammer.portfolio.IsInDrawingMode()
	diagrammer.map_modelElementNode_Shape = map_ModelElementNode_Shape

	diagrammer.map_modelElements_Shape = make(map[any]Shape)
	for elementNode, shape := range diagrammer.map_modelElementNode_Shape {
		diagrammer.map_modelElements_Shape[elementNode.GetElement()] = shape
	}

	for modelElementNode, treeNode := range diagrammer.map_modelElementNode_treeNode {
		treeNode.IsChecked = false
		treeNode.IsCheckboxDisabled = !isInDrawingMode || !modelElementNode.CanBeAddedToDiagram()

		if _, ok := map_ModelElementNode_Shape[modelElementNode]; ok {
			treeNode.IsChecked = true
		}

		if modelElementNode.HasSecondCheckbox() {
			treeNode.HasSecondCheckboxButton = true
			treeNode.IsSecondCheckboxDisabled = !isInDrawingMode
		}
		if modelElementNode.HasSecondName() {
			treeNode.TextAfterSecondCheckbox = modelElementNode.GetSecondName()
		}

	}
}

func (diagrammer *Diagrammer) IsElementNodeDisplayed(modelElementNode ModelElementNode) (ok bool) {

	_, ok = diagrammer.map_modelElementNode_Shape[modelElementNode]

	return
}

func (diagrammer *Diagrammer) GetElementNodeDisplayed(modelElementNode ModelElementNode) (shape Shape, ok bool) {

	shape, ok = diagrammer.map_modelElementNode_Shape[modelElementNode]

	return
}

func (diagrammer *Diagrammer) IsElementDisplayed(modelElement any) (ok bool) {

	_, ok = diagrammer.map_modelElements_Shape[modelElement]

	return
}
