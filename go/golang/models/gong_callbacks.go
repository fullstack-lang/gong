package models

const ModelGongCallbacksFileTemplate = `// generated code - do not edit
package {{PkgGoName}}

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point{{` + string(rune(ModelGongCallbacksPerStruct)) + `}}
`

type ModelGongCallbacksStructInsertionId int

const (
	ModelGongCallbacksPerStruct ModelGongCallbacksStructInsertionId = iota
)

var ModelGongCallbacksStructSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(ModelGongCallbacksPerStruct)): `
func ({{structname}} *{{Structname}}) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfter{{Structname}}CreateCallback != nil {
		stage.OnAfter{{Structname}}CreateCallback.OnAfterCreate(stage, {{structname}})
	}
}

func ({{structname}} *{{Structname}}) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfter{{Structname}}UpdateCallback != nil {
		var front{{Structname}} *{{Structname}}
		if front != nil {
			front{{Structname}}, _ = front.(*{{Structname}})
		}
		stage.OnAfter{{Structname}}UpdateCallback.OnAfterUpdate(stage, {{structname}}, front{{Structname}})
	}
}

func ({{structname}} *{{Structname}}) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfter{{Structname}}DeleteCallback != nil {
		var front{{Structname}} *{{Structname}}
		if front != nil {
			front{{Structname}}, _ = front.(*{{Structname}})
		}
		stage.OnAfter{{Structname}}DeleteCallback.OnAfterDelete(stage, {{structname}}, front{{Structname}})
	}
}
`,
}
