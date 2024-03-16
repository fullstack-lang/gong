package diagrammer

type PortfolioCategoryNode interface {
	PortfolioNode

	// HasAddDiagramButton is true if a "Add" button has to be displayed
	// when the add button is pressed, it calls AddDiagram()
	HasAddDiagramButton() bool

	// AddDiagram allows the end user to request the creation a new Diagram/PortfolioNode
	AddDiagram() PortfolioDiagramNode
}
