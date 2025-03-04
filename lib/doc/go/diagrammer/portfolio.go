package diagrammer

// Portfolio interface to be implemented by the Portfolio adapter
type Portfolio interface {

	// GeneratesProgeny generates and returns the root nodes of the tree for navigating the portfolio
	// root nodes themselves generate their progreny
	GeneratesProgeny() []PortfolioNode

	// GetChildren return the root nodes
	GetChildren() []PortfolioNode

	// IsInSelectionMode is true is at least one diagram has been selected
	IsInSelectionMode() bool // the end user can select a diagram to display
	GetSelectedPortfolioDiagramNode() PortfolioDiagramNode

	IsInDrawingMode() bool
	AddElementToDiagram(ModelElementNode) map[ModelElementNode]Shape
	RemoveElementFromDiagram(ModelElementNode) map[ModelElementNode]Shape
}
