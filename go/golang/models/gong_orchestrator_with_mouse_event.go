package models

const ModelGongOrchestratorMouseEventFileTemplate = `// generated code - do not edit
package models

// insertion point{{` + string(rune(ModelGongOrchestratorStructOnAfterUpdateWithMouseEvent)) + `}}

func SetOrchestratorOnAfterUpdateWithMouseEvent[Type Gongstruct](stage *Stage) {

	var ret Type

	switch any(ret).(type) {
	// insertion point{{` + string(rune(ModelGongOrchestratorOnAfterUpdateWithMouseEventSwitch)) + `}}

	}
}

`

type ModelGongOrchestratorStructWithMouseEventInsertionId int

const (
	ModelGongOrchestratorStructOnAfterUpdateWithMouseEvent ModelGongOrchestratorStructWithMouseEventInsertionId = iota
	ModelGongOrchestratorOnAfterUpdateWithMouseEventSwitch
)

var ModelGongOrchestratorStructWithMouseEventSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(ModelGongOrchestratorStructOnAfterUpdateWithMouseEvent)): `
// {{Structname}}OrchestratorWithMouseEvent
type {{Structname}}OrchestratorWithMouseEvent struct {
}

func (orchestrator *{{Structname}}OrchestratorWithMouseEvent) OnAfterUpdateWithMouseEvent(
	gongsvgStage *Stage,
	staged{{Structname}}, backRepo{{Structname}} *{{Structname}}, mouseEvent *Gong__MouseEvent) {

	staged{{Structname}}.OnAfterUpdateWithMouseEvent(gongsvgStage, backRepo{{Structname}}, mouseEvent)
}`,
	string(rune(ModelGongOrchestratorOnAfterUpdateWithMouseEventSwitch)): `
	case {{Structname}}:
		stage.OnAfter{{Structname}}UpdateWithMouseEventCallback = new({{Structname}}OrchestratorWithMouseEvent)`,
}
