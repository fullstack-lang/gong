package models

type DataFlowType string

var (
	DataFlow_Task2Task                DataFlowType = "DataFlow_Task2Task"
	DataFlow_ExternalParticipant2Task DataFlowType = "DataFlow_ExternalParticipant2Task"
	DataFlow_Task2ExternalParticipant DataFlowType = "DataFlow_Task2ExternalParticipant"
)

type DataFlow struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	Type DataFlowType

	StartTask *Task // non nil if
	EndTask   *Task

	StartExternalParticipant *Participant
	EndExternalParticipant   *Participant

	IsDatasNodeExpanded bool
	Datas               []*Data
}

var _ AbstractType = (*DataFlow)(nil)
