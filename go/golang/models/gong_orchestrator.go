package models

const ModelGongOrchestratorFileTemplate = `// generated code - do not edit
package models

// insertion point{{` + string(rune(ModelGongOrchestratorStruct)) + `}}
// SetOrchestratorOnAfterUpdate is the Stage method for setting orchestrators.
func (stage *Stage) SetOrchestratorOnAfterUpdate[Type Gongstruct]() {

	var ret Type

	switch any(ret).(type) {
	// insertion point{{` + string(rune(ModelGongOrchestratorSwitch)) + `}}

	}

}

// SetOrchestratorOnAfterUpdate is a backward-compatible package-level forwarder.
func SetOrchestratorOnAfterUpdate[Type Gongstruct](stage *Stage) {
	stage.SetOrchestratorOnAfterUpdate[Type]()
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
	gongsvgStage *Stage,
	staged{{Structname}}, backRepo{{Structname}} *{{Structname}}) {

	staged{{Structname}}.OnAfterUpdate(gongsvgStage, staged{{Structname}}, backRepo{{Structname}})
}
`,
	string(rune(ModelGongOrchestratorSwitch)): `
	case {{Structname}}:
		stage.OnAfter{{Structname}}UpdateCallback = new({{Structname}}Orchestrator)`,
}
