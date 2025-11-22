package orm

const GetInstanceDBFromInstanceTemplateCode = `// generated code - do not edit
package orm

type GongstructDB interface {
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
