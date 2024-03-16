package adapter

import "github.com/fullstack-lang/gongdoc/go/diagrammer"

type ElementNodeBase struct {
	portfolioAdapter *PortfolioAdapter
	children         []diagrammer.ModelNode
}

func (base *ElementNodeBase) IsExpanded() bool {
	return false
}

func (base *ElementNodeBase) SetIsExpanded(bool) {}

func (base *ElementNodeBase) IsNameEditable() bool {
	return false
}

func (base *ElementNodeBase) CanBeAddedToDiagram() bool {
	return true
}

func (base *ElementNodeBase) GetChildren() []diagrammer.ModelNode {
	return base.children
}

func (base *ElementNodeBase) GetParent() diagrammer.ModelNode {
	return nil
}
