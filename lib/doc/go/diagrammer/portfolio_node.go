package diagrammer

// PortfolioNode has to be implemented by the Adapter Nodes
type PortfolioNode interface {

	// GetName() returns the Name that is displayed on the node
	GetName() string

	// GetParent returns the node above
	GetParent() PortfolioNode

	// GeneratesProgeny generates and returns the nodes below the node
	// it recursively
	GeneratesProgeny() []PortfolioNode
	GetChildren() []PortfolioNode

	// AppendChildren allows for dynmicaly adding a node below the node
	AppendChildren(PortfolioNode)

	// RemoveChildren allows for dynmicaly adding a node below the node
	RemoveChildren(PortfolioNode)

	// IsExpanded is true if the the node is visualy expanded
	IsExpanded() bool
}
