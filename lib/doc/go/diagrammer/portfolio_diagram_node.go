package diagrammer

type PortfolioDiagramNode interface {
	PortfolioNode

	// DisplayDiagram request the portfolio to display the diagram
	DisplayDiagram() map[ModelElementNode]Shape

	HasDrawButton() bool
	DrawDiagram() map[ModelElementNode]Shape
	CancelEdit() map[ModelElementNode]Shape
	SaveDiagram() map[ModelElementNode]Shape
	IsInDrawingMode() bool

	HasDiagramRenameButton() bool
	RenameDiagram(newName string) error
	IsInRenameMode() bool
	SetIsInRenameMode(isInRenameMode bool)

	HasDuplicateButton() bool
	DuplicateDiagram() PortfolioDiagramNode // returns the new diagram

	HasDeleteButton() bool
	DeleteDiagram()
}
