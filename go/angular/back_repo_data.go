package angular

const BackRepoTemplateTS = `// generated code - do not edit

//insertion point for imports{{` + string(rune(BackRepoDataImports)) + `}}

export class BackRepoData {
	// insertion point for declarations{{` + string(rune(BackRepoInsertionEnumsExportDeclaration)) + `}}


	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies{{` + string(rune(BackRepoInsertionEnumsExportCopies)) + `}}
	}

}`

// insertion points
type BackRepoDataTSInsertionPoint int

const (
	BackRepoDataImports BackRepoDataTSInsertionPoint = iota
	BackRepoInsertionEnumsExportDeclaration
	BackRepoInsertionEnumsExportCopies
	BackRepoNbInsertionPoints
)

var BackRepoHtmlSubTemplateCode map[string]string = map[string]string{
	string(rune(BackRepoDataImports)): `
import { {{Structname}}DB } from './{{structname}}-db'
`,

	string(rune(BackRepoInsertionEnumsExportDeclaration)): `
	{{Structname}}DBs = new Array<{{Structname}}DB>()
`,
	string(rune(BackRepoInsertionEnumsExportCopies)): `
		this.{{Structname}}DBs = data?.{{Structname}}DBs || [];
`,
}
