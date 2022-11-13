package golang

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GongAstGenerator(modelPkg *models.ModelPkg, pkgPath string) {

	codeGO := GongAstTemplate

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace4(codeGO,
		"{{PkgName}}", modelPkg.Name,
		"{{TitlePkgName}}", caserEnglish.String(modelPkg.Name),
		"{{pkgname}}", strings.ToLower(modelPkg.Name),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, "gong_ast.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
