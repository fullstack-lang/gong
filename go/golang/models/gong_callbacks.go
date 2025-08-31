package models

const ModelGongCallbacksFileTemplate = `// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point{{` + string(rune(ModelGongCallbacksCreate)) + `}}
	default:
		_ = target
	}
}

type MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type, mouseEvent *MouseEvent) {

	switch oldTarget := any(old).(type) {
	// insertion point{{` + string(rune(ModelGongCallbacksUpdate)) + `}}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point{{` + string(rune(ModelGongCallbacksDelete)) + `}}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point{{` + string(rune(ModelGongCallbacksRead)) + `}}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point{{` + string(rune(ModelGongCallbacksSetFuncUpdate)) + `}}
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point{{` + string(rune(ModelGongCallbacksSetFuncCreate)) + `}}
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point{{` + string(rune(ModelGongCallbacksSetFuncDelete)) + `}}
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point{{` + string(rune(ModelGongCallbacksSetFuncRead)) + `}}
	}
}
`

type ModelGongCallbacksStructInsertionId int

const (
	ModelGongCallbacksCreate ModelGongCallbacksStructInsertionId = iota
	ModelGongCallbacksUpdate
	ModelGongCallbacksRead
	ModelGongCallbacksDelete
	ModelGongCallbacksSetFuncCreate
	ModelGongCallbacksSetFuncUpdate
	ModelGongCallbacksSetFuncRead
	ModelGongCallbacksSetFuncDelete
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
		if stage.OnAfter{{Structname}}UpdateCallback != nil && mouseEvent == nil {
			stage.OnAfter{{Structname}}UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfter{{Structname}}UpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfter{{Structname}}UpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}`,
	string(rune(ModelGongCallbacksRead)): `
	case *{{Structname}}:
		if stage.OnAfter{{Structname}}ReadCallback != nil {
			stage.OnAfter{{Structname}}ReadCallback.OnAfterRead(stage, target)
		}`,
	string(rune(ModelGongCallbacksDelete)): `
	case *{{Structname}}:
		if stage.OnAfter{{Structname}}DeleteCallback != nil {
			staged := any(staged).(*{{Structname}})
			stage.OnAfter{{Structname}}DeleteCallback.OnAfterDelete(stage, staged, front)
		}`,
	string(rune(ModelGongCallbacksSetFuncCreate)): `
	case *{{Structname}}:
		stage.OnAfter{{Structname}}CreateCallback = any(callback).(OnAfterCreateInterface[{{Structname}}])
	`,
	string(rune(ModelGongCallbacksSetFuncUpdate)): `
	case *{{Structname}}:
		stage.OnAfter{{Structname}}UpdateCallback = any(callback).(OnAfterUpdateInterface[{{Structname}}])
	`,
	string(rune(ModelGongCallbacksSetFuncRead)): `
	case *{{Structname}}:
		stage.OnAfter{{Structname}}ReadCallback = any(callback).(OnAfterReadInterface[{{Structname}}])
	`,
	string(rune(ModelGongCallbacksSetFuncDelete)): `
	case *{{Structname}}:
		stage.OnAfter{{Structname}}DeleteCallback = any(callback).(OnAfterDeleteInterface[{{Structname}}])
	`,
}
