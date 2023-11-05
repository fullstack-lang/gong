package angular

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fullstack-lang/gong/go/models"
)

const NgNullInt64TemplateTS = `// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
    Int64: number = 0
    Valid: boolean = false
}
`

// MultiCodeGeneratorNgCommitNb parses mdlPkg and generates the code for the
// CommitNb components
func CodeGeneratorNgNullInt64(modelPkg *models.ModelPkg) {

	matTargetPath := modelPkg.NgDataLibrarySourceCodeDirectory

	codeTS := NgNullInt64TemplateTS

	file, err := os.Create(filepath.Join(matTargetPath, "null-int64.ts"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeTS)
}
