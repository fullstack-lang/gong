package angular

const BackRepoTemplateTS = `// generated code - do not edit

//insertion point for imports{{` + string(rune(BackRepoDataImports)) + `}}

export class BackRepoData {
	// insertion point for declarations{{` + string(rune(BackRepoInsertionEnumsExportDeclaration)) + `}}

}

constructor(data?: Partial<BackRepoData>) {
	// insertion point for copies{{` + string(rune(BackRepoInsertionEnumsExportCopies)) + `}}
}
`

// insertion points
type BackRepoDataTSInsertionPoint int

const (
	BackRepoDataImports BackRepoDataTSInsertionPoint = iota
	BackRepoInsertionEnumsExportDeclaration
	BackRepoInsertionEnumsExportCopies
	BackRepoNbInsertionPoints
)

// Sub Templates
type BackRepoDataTSSubTemplate int

const (
	BackRepoDataTSImports BackRepoDataTSSubTemplate = iota
	BackRepoDEnumsExportDeclaration
)

var BackRepoHtmlSubTemplateCode map[string]string = map[string]string{
	string(rune(BackRepoDataTSImports)): `
import { {{Structname}}DB } from './{{structname}}-db'
`,
}
