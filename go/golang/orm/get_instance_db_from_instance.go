package orm

const GetInstanceDBFromInstanceTemplateCode = `// generated code - do not edit
package orm

import (
	"{{PkgPathRoot}}/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup{{` + string(rune(GetInstanceDBFromInstanceSwitchCaseGet)) + `}}
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup{{` + string(rune(GetInstanceDBFromInstanceSwitchCaseGetID)) + `}}
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup{{` + string(rune(GetInstanceDBFromInstanceSwitchCaseGetID)) + `}}
	default:
		_ = inst
	}
	return
}
`

type GetInstanceDBFromInstanceSubTemplateInsertion int

const (
	GetInstanceDBFromInstanceSwitchCaseGetID GetInstanceDBFromInstanceSubTemplateInsertion = iota
	GetInstanceDBFromInstanceSwitchCaseGet
)

var GetInstanceDBFromInstanceSubTemplate map[string]string = // new line
map[string]string{

	string(rune(GetInstanceDBFromInstanceSwitchCaseGetID)): `
	case *models.{{Structname}}:
		tmp := GetInstanceDBFromInstance[models.{{Structname}}, {{Structname}}DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)`,
	string(rune(GetInstanceDBFromInstanceSwitchCaseGet)): `
	case *models.{{Structname}}:
		{{structname}}Instance := any(concreteInstance).(*models.{{Structname}})
		ret2 := backRepo.BackRepo{{Structname}}.Get{{Structname}}DBFrom{{Structname}}Ptr({{structname}}Instance)
		ret = any(ret2).(*T2)`,
}
