package angular

import (
	"log"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/fullstack-lang/gong/go/models"
)

//go:embed ng_file_splitter.css
var NgFileSplitterCssTmpl string

const NgSplitterTemplateTS = `import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-{{pkgname}}-splitter',
  templateUrl: './splitter.component.html',
  styleUrls: ['./splitter.component.css'],
})
export class SplitterComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }
}
`

// MultiCodeGeneratorNgSplitter parses mdlPkg and generates the code for the
// Splitter components
func CodeGeneratorNgSplitter(
	mdlPkg *models.ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string) {

	// create the component directory
	dirPath := filepath.Join(matTargetPath, "splitter")
	errd := os.MkdirAll(dirPath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + dirPath)
	}

	// generate the css file
	models.VerySimpleCodeGenerator(mdlPkg,
		pkgName,
		pkgGoPath,
		filepath.Join(dirPath, "splitter.component.css"),
		NgFileSplitterCssTmpl,
	)

	// generate the ts file
	models.VerySimpleCodeGenerator(mdlPkg,
		pkgName,
		pkgGoPath,
		filepath.Join(dirPath, "splitter.component.ts"),
		NgSplitterTemplateTS)

	// generate the html file
	models.VerySimpleCodeGenerator(mdlPkg,
		pkgName,
		pkgGoPath,
		filepath.Join(dirPath, "splitter.component.html"),
		NgSplitterTemplateHTML)

}
