package tests

import (
	"io"
	"os"
	"path/filepath"

	"testing"

	"github.com/fullstack-lang/gong/go/models"
)

func TestTSConfigModification(t *testing.T) {

	filename := filepath.Join(".", "tsconfig.json")

	pkgName := "gong"

	models.InsertStringToFile(filename,
		"        \"projects/"+pkgName+"/src/public-api.ts\",",
		pkgName+"\": [")

	models.InsertStringToFile(filename,
		models.TsConfigInsertForPaths,
		"\"paths\": {")
}

func TestAppModuleModification(t *testing.T) {

	filenameOrig := filepath.Join(".", "app.module.orig.txt")
	filenameResult := filepath.Join(".", "app.module.result.txt")

	Copy(filenameOrig, filenameResult)

	models.InsertStringToFile(filenameResult,
		models.AppModuleImport,
		"import { BrowserAnimationsModule } from '@angular/platform-browser/animations';")

	models.InsertStringToFile(filenameResult,
		models.AppModuleImport2,
		"        HttpClientModule,")

}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
