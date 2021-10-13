package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const NgNullInt64TemplateTS = `// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
    Int64: number = 0
    Valid: boolean = true // exception to the golang convention that boolean default value is false
}
`

// MultiCodeGeneratorNgCommitNb parses mdlPkg and generates the code for the
// CommitNb components
func CodeGeneratorNgNullInt64(
	mdlPkg *ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string,
	apiPath string) {

	codeTS := NgNullInt64TemplateTS

	file, err := os.Create(filepath.Join(matTargetPath, "null-int64.ts"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeTS)
}
