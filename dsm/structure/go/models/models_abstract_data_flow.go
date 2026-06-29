package models

type DataFlowType string

var (
	DataFlow_Port2Port         DataFlowType = "Port2Port"
	DataFlow_ExternalPart2Port DataFlowType = "ExternalPart2Port"
	DataFlow_Port2ExternalPart DataFlowType = "Port2ExternalPart"
)

type DataFlowDirection string

var (
	Unspecified DataFlowDirection = "Unspecified"
	Forward     DataFlowDirection = "Forward"
	Backward    DataFlowDirection = "Backward"
	BothWays    DataFlowDirection = "BothWays"
)

type DataFlow struct {
	Name string

	StartPort *Port // non nil if
	EndPort   *Port

	StartExternalPart *Part
	EndExternalPart   *Part

	Datas []*Data

	//gong:text width:300 height:300
	Description string

	Direction DataFlowDirection

	IsDatasNodeExpanded bool

	LibraryAbstractFields
	AbstractTypeFields

	// unused for the moment
	Type DataFlowType
}

var _ AbstractType = (*DataFlow)(nil)
