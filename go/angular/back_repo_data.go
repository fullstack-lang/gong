package angular

const BackRepoTemplateTS = `// generated code - do not edit

//insertion point for imports{{` + string(rune(BackRepoDataImports)) + `}}

export class BackRepoData {
	// insertion point for declarations{{` + string(rune(BackRepoInsertionEnumsExportDeclaration)) + `}}

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies{{` + string(rune(BackRepoInsertionEnumsExportCopies)) + `}}
		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
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
import { {{Structname}}API } from './{{structname}}-api'
`,

	string(rune(BackRepoInsertionEnumsExportDeclaration)): `
	{{Structname}}APIs = new Array<{{Structname}}API>()
`,
	string(rune(BackRepoInsertionEnumsExportCopies)): `
		this.{{Structname}}APIs = data?.{{Structname}}APIs || [];
`,
}
