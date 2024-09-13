package adapter

import "github.com/fullstack-lang/gongdoc/go/diagrammer"

type ModelCategoryNodeBase struct {
	portfolioAdapter *PortfolioAdapter
	children         []diagrammer.ModelNode
	Name             string
	isExpanded       bool
}

func (base *ModelCategoryNodeBase) IsExpanded() bool {
	return true
}

func (base *ModelCategoryNodeBase) SetIsExpanded(isExpanded bool) {
	base.isExpanded = isExpanded
}

func (base *ModelCategoryNodeBase) IsNameEditable() bool {
	return false
}

func (base *ModelCategoryNodeBase) GetName() string {
	return base.Name
}

func (base *ModelCategoryNodeBase) GetChildren() []diagrammer.ModelNode {
	return base.children
}

func (base *ModelCategoryNodeBase) GetParent() diagrammer.ModelNode {
	return nil
}

func (base *ModelCategoryNodeBase) HasSecondCheckbox() bool {
	return false
}

func (base *ModelCategoryNodeBase) HasSecondName() bool {
	return false
}

func (base *ModelCategoryNodeBase) GetSecondName() string {
	return ""
}

func (base *ModelCategoryNodeBase) GetLinkedNode() diagrammer.ModelElementNode {
	return nil
}
