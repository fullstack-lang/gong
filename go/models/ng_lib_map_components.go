package models

const NgLibMapComponentsServiceTemplate = `// insertion point sub template for components imports {{` + string(NgLibMapComponentsImports) + `}}

// insertion point sub template for map of components per struct {{` + string(NgLibMapComponentsIndivDecls) + `}}

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components {{` + string(NgLibMapComponentsDecls) + `}}
    ]
  )`

type NgLibMapComponentsServiceSubTemplate string

const (
	NgLibMapComponentsImports    NgLibMapComponentsServiceSubTemplate = "NgLibMapComponentsImports"
	NgLibMapComponentsIndivDecls NgLibMapComponentsServiceSubTemplate = "NgLibMapComponentsIndivDecls"
	NgLibMapComponentsDecls      NgLibMapComponentsServiceSubTemplate = "NgLibMapComponentsDecls"
)

var NgLibMapComponentsSubTemplateCode map[string]string = // new line
map[string]string{

	string(NgLibMapComponentsImports): `
import { {{Structname}}sTableComponent } from './{{structname}}s-table/{{structname}}s-table.component'`,

	string(NgLibMapComponentsIndivDecls): `
export const MapOf{{Structname}}sComponents: Map<string, any> = new Map([["{{Structname}}sTableComponent", {{Structname}}sTableComponent],])`,

	string(NgLibMapComponentsDecls): `
      ["{{Structname}}", MapOf{{Structname}}sComponents],`,
}
