package angular

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
)

const NgGongselectionServiceTs = `import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable()
export class GongstructSelectionService {

    // Observable string sources
    private gongstructSelection = new Subject<string>();

    // Observable string streams
    gongtructSelected$ = this.gongstructSelection.asObservable();

    // Service message commands
    gongstructSelected(gongstructName: string) {
        this.gongstructSelection.next(gongstructName);
    }

}
`

// CodeGeneratorNgGongselectionServiceTs
func CodeGeneratorNgGongselectionServiceTs(
	mdlPkg *models.ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string,
	apiPath string) {

	codeTS := NgGongselectionServiceTs
	codeTS = strings.ReplaceAll(codeTS, "{{addr}}", apiPath)

	file, err := os.Create(filepath.Join(matTargetPath, "gongstruct-selection.service.ts"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeTS)

}
