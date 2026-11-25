package models

type DataFlowType string

var (
	DataFlow_Task2Task                DataFlowType = "DataFlow_Task2Task"
	DataFlow_ExternalParticipant2Task DataFlowType = "DataFlow_ExternalParticipant2Task"
	DataFlow_Task2ExternalParticipant DataFlowType = "DataFlow_Task2ExternalParticipant"
	DataFlow_DataStore2Task           DataFlowType = "DataFlow_DataStore2Task"
	DataFlow_Task2DataStore           DataFlowType = "DataFlow_Task2DataStore"
)

type DataFlow struct {
	Name string

	Datas []*Data

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	Type DataFlowType

	StartTask *Task // non nil if
	EndTask   *Task

	StartParticipant *ExternalParticipant
	EndParticipant   *ExternalParticipant

	StartDataStore *DataStore
	EndDataStore   *DataStore

	IsDatasNodeExpanded bool
}

var _ AbstractType = (*DataFlow)(nil)
