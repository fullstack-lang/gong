package golang

const ModelGongOrchestratorFileTemplate = `package models

// insertion point{{` + string(rune(ModelGongOrchestratorStruct)) + `}}

func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *StageStruct) {

	var ret Type

	switch any(ret).(type) {
	// insertion point{{` + string(rune(ModelGongOrchestratorSwitch)) + `}}

	}

}
`

type ModelGongOrchestratorStructInsertionId int

const (
	ModelGongOrchestratorStruct ModelGongOrchestratorStructInsertionId = iota
	ModelGongOrchestratorSwitch
)

var ModelGongOrchestratorStructSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(ModelGongOrchestratorStruct)): `
// {{Structname}}Orchestrator
type {{Structname}}Orchestrator struct {
}

func (orchestrator *{{Structname}}Orchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	staged{{Structname}}, backRepo{{Structname}} *{{Structname}}) {

	staged{{Structname}}.OnAfterUpdate(gongsvgStage, staged{{Structname}}, backRepo{{Structname}})
}`,
	string(rune(ModelGongOrchestratorSwitch)): `
	case {{Structname}}:
		stage.OnAfter{{Structname}}UpdateCallback = new({{Structname}}Orchestrator)`,
}
