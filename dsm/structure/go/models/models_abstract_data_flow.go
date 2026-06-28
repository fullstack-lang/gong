package models

type DataFlowType string

var (
	DataFlow_Port2Port         DataFlowType = "Port2Port"
	DataFlow_ExternalPart2Port DataFlowType = "ExternalPart2Port"
	DataFlow_Port2ExternalPart DataFlowType = "Port2ExternalPart"
)

type DataFlowDirection string

var (
	DataFlow_Forward  DataFlowDirection = "Forward"
	DataFlow_Backward DataFlowDirection = "Backward"
	DataFlow_BothWays DataFlowDirection = "BothWays"
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

	Type      DataFlowType
	Direction DataFlowDirection

	IsDatasNodeExpanded bool

	LibraryAbstractFields
	AbstractTypeFields
}

var _ AbstractType = (*DataFlow)(nil)
