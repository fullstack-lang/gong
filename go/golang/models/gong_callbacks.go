package models

const ModelGongCallbacksFileTemplate = `// generated code - do not edit
package {{PkgGoName}}

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point{{` + string(rune(ModelGongCallbacksCreate)) + `}}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point{{` + string(rune(ModelGongCallbacksUpdate)) + `}}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point{{` + string(rune(ModelGongCallbacksDelete)) + `}}
	default:
		_ = front
	}
}
`

type ModelGongCallbacksStructInsertionId int

const (
	ModelGongCallbacksCreate ModelGongCallbacksStructInsertionId = iota
	ModelGongCallbacksUpdate
	ModelGongCallbacksDelete
)

var ModelGongCallbacksStructSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(ModelGongCallbacksCreate)): `
	case *{{Structname}}:
		if stage.OnAfter{{Structname}}CreateCallback != nil {
			stage.OnAfter{{Structname}}CreateCallback.OnAfterCreate(stage, target)
		}`,
	string(rune(ModelGongCallbacksUpdate)): `
	case *{{Structname}}:
		newTarget := any(new).(*{{Structname}})
		if stage.OnAfter{{Structname}}UpdateCallback != nil {
			stage.OnAfter{{Structname}}UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}`,
	string(rune(ModelGongCallbacksDelete)): `
	case *{{Structname}}:
		if stage.OnAfter{{Structname}}DeleteCallback != nil {
			staged := any(staged).(*{{Structname}})
			stage.OnAfter{{Structname}}DeleteCallback.OnAfterDelete(stage, staged, front)
		}`,
}
