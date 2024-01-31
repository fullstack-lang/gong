package orm

const BackRepoDataTemplateCode = `// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices{{` + string(rune(BackRepoDataSlice)) + `}}
}
`

type BackRepoDataSubTemplateInsertion int

const (
	BackRepoDataSlice BackRepoDataSubTemplateInsertion = iota
)

var BackRepoDataSubTemplate map[string]string = // new line
map[string]string{

	string(rune(BackRepoDataSlice)): `

	{{Structname}}DBs []*{{Structname}}DB`,
}
