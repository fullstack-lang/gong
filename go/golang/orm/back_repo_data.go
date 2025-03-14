package orm

const BackRepoDataTemplateCode = `// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices{{` + string(rune(BackRepoDataSlice)) + `}}

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies{{` + string(rune(BackRepoDataSliceCopies)) + `}}
}
`

type BackRepoDataSubTemplateInsertion int

const (
	BackRepoDataSlice BackRepoDataSubTemplateInsertion = iota
	BackRepoDataSliceCopies
)

var BackRepoDataSubTemplate map[string]string = // new line
map[string]string{

	string(rune(BackRepoDataSlice)): `

	{{Structname}}APIs []*{{Structname}}API`,

	string(rune(BackRepoDataSliceCopies)): `
	for _, {{structname}}DB := range backRepo.BackRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB {

		var {{structname}}API {{Structname}}API
		{{structname}}API.ID = {{structname}}DB.ID
		{{structname}}API.{{Structname}}PointersEncoding = {{structname}}DB.{{Structname}}PointersEncoding
		{{structname}}DB.CopyBasicFieldsTo{{Structname}}_WOP(&{{structname}}API.{{Structname}}_WOP)

		backRepoData.{{Structname}}APIs = append(backRepoData.{{Structname}}APIs, &{{structname}}API)
	}
`,
}
