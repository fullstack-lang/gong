package models

// GongdocStatus is the struct of the instance that is updated by the back for ackonleding commands
type GongdocStatus struct {
	Name                  string
	Status                GongdocCommandType
	CommandCompletionDate string
}

var GongdocStatusSingloton = (&GongdocStatus{
	Name:                  "Gongdoc Status Singloton",
	Status:                MARSHALL_DIAGRAM,
	CommandCompletionDate: "",
}).Stage()
