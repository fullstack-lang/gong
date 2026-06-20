package models

type DataFlowType string

var (
	DataFlow_Port2Port                DataFlowType = "DataFlow_Port2Port"
	DataFlow_ExternalPart2Port DataFlowType = "DataFlow_ExternalPart2Port"
	DataFlow_Port2ExternalPart DataFlowType = "DataFlow_Port2ExternalPart"
)

type DataFlow struct {
	Name string

	Datas []*Data

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	Type DataFlowType

	StartPort *Port // non nil if
	EndPort   *Port

	StartExternalPart *Part
	EndExternalPart   *Part

	IsDatasNodeExpanded bool
}

var _ AbstractType = (*DataFlow)(nil)
