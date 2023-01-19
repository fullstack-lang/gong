package models

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// LoadDiagramPackage fill up the stage with the diagrams elements
func LoadDiagramPackage(pkgPath string, modelPkg *gong_models.ModelPkg, editable bool) (diagramPackage *DiagramPackage, err error) {

	gongdocStage := Stage
	_ = gongdocStage

	diagramPackage = (&DiagramPackage{}).Stage()
	diagramPackage.IsEditable = editable
	diagramPackage.ModelPkg = modelPkg
	diagramPackage.Name = modelPkg.Name + "_diagrams"

	// parse the diagram package
	diagramPkgPath := filepath.Join(pkgPath, "../diagrams")
	diagramPackage.AbsolutePathToDiagramPackage, _ = filepath.Abs(diagramPkgPath)
	diagramPackage.Path = diagramPkgPath
	diagramPackage.GongModelPath = modelPkg.PkgPath

	// diagram package, when marshalled, will reference identifiers in the
	// model package. Both of the variable need to be set up for the
	// generic marshalling/unmarshalling to work
	Stage.MetaPackageImportAlias = "ref_" + filepath.Base(diagramPackage.GongModelPath)
	Stage.MetaPackageImportPath = `"` + diagramPackage.GongModelPath + `"`

	// if diagrams directory does not exist create it
	_, errd := os.Stat(diagramPkgPath)
	if os.IsNotExist(errd) {
		log.Printf(diagramPkgPath, " does not exist, gongdoc creates it")

		errd := os.MkdirAll(diagramPkgPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + diagramPkgPath)
		}
		if os.IsExist(errd) {
			log.Println("directory " + diagramPkgPath + " allready exists")
		}
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	pkgsParser, errParser := parser.ParseDir(fset, diagramPkgPath, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		log.Panic("Unable to parser ")
	}
	diagramPackageAst, ok := pkgsParser["diagrams"]
	if !ok {
		diagramPackage.IsEditable = editable
		FillUpNodeTree(diagramPackage)
		Stage.Commit()
		return diagramPackage, nil
	}
	diagramPackage.ast = diagramPackageAst
	diagramPackage.fset = fset

	// TO BE REMOVED AFTER THE TRANSITION PHASE
	//
	// check wether the file is in the legacy format
	// convert it to the new format if it is the case
	// if one legacy format is found, exit

	var oneLegacyFormatFound bool

	for fileName := range diagramPackageAst.Files {
		diagramName := strings.TrimSuffix(filepath.Base(fileName), ".go")
		diagramPackage.Files = append(diagramPackage.Files, diagramName)

		pathToDiagramFile, err := filepath.Abs(fileName)
		if err != nil {
			log.Fatalln("Path does not exist ", pathToDiagramFile)
		}

		fset := token.NewFileSet()
		startParser := time.Now()
		inFile, errParser := parser.ParseFile(fset, pathToDiagramFile, nil, parser.ParseComments)
		log.Printf("Parser took %s", time.Since(startParser))
		if errParser != nil {
			log.Fatalln("Unable to parse ", pathToDiagramFile, errParser.Error())
		}

		var isGenericFormat bool
		for _, decl := range inFile.Decls {
			switch decl := decl.(type) {
			case *ast.FuncDecl:
				funcDecl := decl
				if name := funcDecl.Name; name != nil {
					isOfInterest := strings.Contains(funcDecl.Name.Name, "Injection")
					if !isOfInterest {
						continue
					}
					isGenericFormat = true
				}
			}
		}
		if !isGenericFormat {
			log.Println("File", diagramName, "is in the legacy format")

			// checkout in order to get the latest version of the diagram before
			// modifying it updated by the front
			stage := Stage
			_ = stage
			Stage.Checkout()
			Stage.Unstage() // clean the stage

			// load the file to the stage
			classdiagram := diagramPackage.Unmarshall(
				diagramPackage.ModelPkg,
				diagramPackage.ast,
				diagramPackage.fset,
				diagramPackage.GongModelPath,
				diagramName)

			// some files, loke docs.go, are not in generic format and
			// and are of no interest
			if classdiagram == nil {
				log.Println("File", diagramName, "is neither in the generic format or the generic format")
				continue
			}

			oneLegacyFormatFound = true
			diagramPackage.SerializeToStage()

			// marshall the stage to the new format.
			file, err := os.Create(pathToDiagramFile)
			if err != nil {
				log.Fatal("Cannot open diagram file" + err.Error())
			}
			defer file.Close()
			Stage.Marshall(file, "github.com/fullstack-lang/gongdoc/go/models", "diagrams")

			// restore the original stage
			diagramPackage.Classdiagrams = make([]*Classdiagram, 0)
			Stage.Unstage()
			Stage.Checkout()
		} else {
			log.Println("File", diagramName, "is in the generic format")
		}

		// now the file is in the generic format. It can be load
	}

	if oneLegacyFormatFound {
		log.Fatalln("Found at least one legacy format. Restart the application")
	} else {
		// the number of instances per classshape has to be restored.
		Stage.Unstage()
		Stage.Checkout()
	}

	// End of TO BE REMOVED AFTER TRANSITION

	// load all diagram files
	for fileName := range diagramPackageAst.Files {
		diagramName := strings.TrimSuffix(filepath.Base(fileName), ".go")
		diagramPackage.UnmarshallOneDiagram(diagramName)
	}

	diagramPackage.IsEditable = editable
	FillUpNodeTree(diagramPackage)
	Stage.Commit()
	return diagramPackage, nil
}
