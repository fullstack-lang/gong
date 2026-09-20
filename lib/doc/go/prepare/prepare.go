package prepare

import (
	"embed"
	"fmt"
	"go/ast"
	"go/token"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	form_fullstack "github.com/fullstack-lang/gong/lib/form/go/fullstack"
	svg_fullstack "github.com/fullstack-lang/gong/lib/svg/go/fullstack"
	tree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"

	gong "github.com/fullstack-lang/gong/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	splitlite "github.com/fullstack-lang/gong/lib/splitlite/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

// hook marhalling to stage
type beforeCommitImplementation struct {
	marshallOnCommit string

	packageName string
}

func (beforeCommitImplementation *beforeCommitImplementation) BeforeCommit(stage *models.Stage) {

	// the ".go" is not provided
	filename := beforeCommitImplementation.marshallOnCommit
	if !strings.HasSuffix(filename, ".go") {
		filename = filename + ".go"
	}

	file, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	packageName := beforeCommitImplementation.packageName
	if packageName == "" {
		packageName = "main"
	}

	stage.Marshall(file, "github.com/fullstack-lang/gong/lib/doc/go/models", packageName)
}

func prepareStages(
	r *http.ServeMux,
	embeddedDiagrams bool,
	docStackName string,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
) (
	stage *models.Stage,
	treeStage *tree.Stage,
	svgStage *svg.Stage,
	gongStage *gong.Stage,
	formStage *form.Stage,
	treeNavigationStage *tree.Stage,
) {
	stage = models.NewStage(docStackName)

	stage.MetaPackageImportAlias = "ref_models"

	splits := strings.Split(docStackName, ":")
	stage.MetaPackageImportPath = `"` + splits[0] + `/models"`

	if !embeddedDiagrams {
		diagramsPath := "../../models/diagrams/diagrams.go"
		if _, err := os.Stat(diagramsPath); os.IsNotExist(err) {
			if _, errOld := os.Stat("../../diagrams/diagrams.go"); errOld == nil {
				diagramsPath = "../../diagrams/diagrams.go"
			}
		}

		err := stage.ParseAstFile(diagramsPath, true)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		BeforeCommitImplementation := &beforeCommitImplementation{
			marshallOnCommit: diagramsPath,
			packageName:      "diagrams", // necessity because the diagram file is in a diagrams package
		}
		stage.OnInitCommitCallback = BeforeCommitImplementation

		// use delta mode
		stage.SetDeltaMode(true)
		stage.ComputeReferenceAndOrders() // from which the delta are computed

	} else {
		err := stage.ParseAstEmbeddedFile(goDiagramsDir, "models/diagrams/diagrams.go")
		if err != nil {
			err = stage.ParseAstEmbeddedFile(goDiagramsDir, "diagrams/diagrams.go")
		}

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}
	}

	treeStage, _ = tree_fullstack.NewStackInstance(r, docStackName+":doc-sidebar", "", "")
	svgStage, _ = svg_fullstack.NewStackInstance(r, docStackName+":doc-svg", "", "", "")
	gongStage = gong.NewStage(docStackName + ":doc-gong")
	formStage, _ = form_fullstack.NewStackInstance(r, docStackName+":doc-diagramForm", "", "")
	treeNavigationStage, _ = tree_fullstack.NewStackInstance(r, docStackName+":doc-sidebar-navigation", "", "")

	// load the code of the model of interest into the gongStage
	gong.LoadEmbedded(gongStage, goModelsDir)

	return
}

func Prepare(
	r *http.ServeMux,
	embeddedDiagrams bool,
	docStackName string,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	receivingAsSplitArea *split.AsSplitArea, // split area that will receive the doc areas
	map_GongStructName_InstancesNb map[string]int,
) (stager *models.Stager) {
	stage, treeStage, svgStage, gongStage, formStage, treeNavigationStage := prepareStages(r, embeddedDiagrams, docStackName, goModelsDir, goDiagramsDir)

	return models.NewStager(
		r,
		receivingAsSplitArea,
		stage,
		treeStage,
		svgStage,
		gongStage,
		formStage,
		treeNavigationStage,
		embeddedDiagrams,
		map_GongStructName_InstancesNb)
}

func PrepareSplitlite(
	r *http.ServeMux,
	embeddedDiagrams bool,
	docStackName string,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	receivingAsSplitArea *splitlite.AsSplitArea, // split area that will receive the doc areas
	map_GongStructName_InstancesNb map[string]int,
) (stager *models.Stager) {
	stage, treeStage, svgStage, gongStage, formStage, treeNavigationStage := prepareStages(r, embeddedDiagrams, docStackName, goModelsDir, goDiagramsDir)

	return models.NewStagerSplitlite(
		r,
		receivingAsSplitArea,
		stage,
		treeStage,
		svgStage,
		gongStage,
		formStage,
		treeNavigationStage,
		embeddedDiagrams,
		map_GongStructName_InstancesNb)
}

func loadEmbeddedPackages(stage *gong.Stage, goModelsDir embed.FS) {
	dirsWithGoFiles := make(map[string]bool)
	_ = fs.WalkDir(goModelsDir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() && filepath.Ext(path) == ".go" {
			dir := filepath.Dir(path)
			// skip diagrams directory and probe directory
			if !strings.Contains(dir, "diagrams") && !strings.Contains(dir, "probe") {
				dirsWithGoFiles[dir] = true
			}
		}
		return nil
	})

	if len(dirsWithGoFiles) == 0 {
		return
	}

	allPkgs := make(map[string]map[string]*ast.Package)
	for dir := range dirsWithGoFiles {
		modelPkg := gong.NewModelPkg(stage)
		modelPkg.Fset = token.NewFileSet()
		pkgs := gong.ParseEmbedModelWithFset(goModelsDir, dir, modelPkg.Fset)
		if len(pkgs) > 0 {
			gong.WalkParser(pkgs, modelPkg, nil)
			modelPkg.SerializeToStage()
			allPkgs[dir] = pkgs
		}
	}

	// Index all staged GongStructs by both package-qualified and unqualified names
	structsByPkgAndName := make(map[string]*gong.GongStruct)
	structsByName := make(map[string]*gong.GongStruct)
	for gs := range *stage.GetInstancesSet[*gong.GongStruct]() {
		structsByName[gs.Name] = gs
		if gs.ModelPkg != nil && gs.ModelPkg.PkgGoName != "" {
			structsByPkgAndName[gs.ModelPkg.PkgGoName+"."+gs.Name] = gs
		}
	}

	// Scan ASTs for cross-package pointer and slice fields that single-package parser skips
	for _, pkgs := range allPkgs {
		for _, astPkg := range pkgs {
			for _, file := range astPkg.Files {
				for _, decl := range file.Decls {
					genDecl, ok := decl.(*ast.GenDecl)
					if !ok || genDecl.Tok != token.TYPE {
						continue
					}
					for _, spec := range genDecl.Specs {
						typeSpec, ok := spec.(*ast.TypeSpec)
						if !ok {
							continue
						}
						structType, ok := typeSpec.Type.(*ast.StructType)
						if !ok {
							continue
						}
						owningGS, ok := structsByPkgAndName[astPkg.Name+"."+typeSpec.Name.Name]
						if !ok {
							owningGS = structsByName[typeSpec.Name.Name]
						}
						if owningGS == nil {
							continue
						}

						existingFields := make(map[string]bool)
						for _, f := range owningGS.Fields {
							existingFields[f.GetName()] = true
						}

						for _, field := range structType.Fields.List {
							if len(field.Names) == 0 {
								continue
							}
							fieldName := field.Names[0].Name
							if existingFields[fieldName] {
								continue
							}

							switch ft := field.Type.(type) {
							case *ast.StarExpr:
								if sel, ok := ft.X.(*ast.SelectorExpr); ok {
									pkgIdent, isIdent := sel.X.(*ast.Ident)
									targetName := sel.Sel.Name
									var targetGS *gong.GongStruct
									if isIdent {
										targetGS = structsByPkgAndName[pkgIdent.Name+"."+targetName]
									}
									if targetGS == nil {
										targetGS = structsByName[targetName]
									}
									if targetGS != nil {
										ptrField := (&gong.PointerToGongStructField{
											Name:       fieldName,
											GongStruct: targetGS,
											Index:      len(owningGS.Fields),
										}).Stage(stage)
										owningGS.Fields = append(owningGS.Fields, ptrField)
										owningGS.PointerToGongStructFields = append(owningGS.PointerToGongStructFields, ptrField)
										existingFields[fieldName] = true
									}
								}
							case *ast.ArrayType:
								if star, ok := ft.Elt.(*ast.StarExpr); ok {
									if sel, ok := star.X.(*ast.SelectorExpr); ok {
										pkgIdent, isIdent := sel.X.(*ast.Ident)
										targetName := sel.Sel.Name
										var targetGS *gong.GongStruct
										if isIdent {
											targetGS = structsByPkgAndName[pkgIdent.Name+"."+targetName]
										}
										if targetGS == nil {
											targetGS = structsByName[targetName]
										}
										if targetGS != nil {
											sliceField := (&gong.SliceOfPointerToGongStructField{
												Name:       fieldName,
												GongStruct: targetGS,
												Index:      len(owningGS.Fields),
											}).Stage(stage)
											owningGS.Fields = append(owningGS.Fields, sliceField)
											owningGS.SliceOfPointerToGongStructFields = append(owningGS.SliceOfPointerToGongStructFields, sliceField)
											existingFields[fieldName] = true
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// Link intra-package pointer targets if any are placeholder instances
	for gs := range *stage.GetInstancesSet[*gong.GongStruct]() {
		for _, ptrField := range gs.PointerToGongStructFields {
			if ptrField.GongStruct != nil {
				if actual, ok := structsByName[ptrField.GongStruct.Name]; ok && ptrField.GongStruct != actual {
					ptrField.GongStruct = actual
				}
			}
		}
		for _, sliceField := range gs.SliceOfPointerToGongStructFields {
			if sliceField.GongStruct != nil {
				if actual, ok := structsByName[sliceField.GongStruct.Name]; ok && sliceField.GongStruct != actual {
					sliceField.GongStruct = actual
				}
			}
		}
	}
	stage.Commit()
}

func prepareStagesSet(
	r *http.ServeMux,
	embeddedDiagrams bool,
	docStackName string,
	metaPackageImports []*models.MetaPackageImport,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
) (
	stage *models.Stage,
	treeStage *tree.Stage,
	svgStage *svg.Stage,
	gongStage *gong.Stage,
	formStage *form.Stage,
	treeNavigationStage *tree.Stage,
) {
	stage = models.NewStage(docStackName)
	stage.MetaPackageImports = metaPackageImports

	if !embeddedDiagrams {
		diagramsPath := "../../models/diagrams/diagrams_set.go"
		if _, err := os.Stat(diagramsPath); os.IsNotExist(err) {
			if _, errOld := os.Stat("../../diagrams/diagrams_set.go"); errOld == nil {
				diagramsPath = "../../diagrams/diagrams_set.go"
			}
		}

		err := stage.ParseAstFile(diagramsPath, true)
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		BeforeCommitImplementation := &beforeCommitImplementation{
			marshallOnCommit: diagramsPath,
			packageName:      "diagrams",
		}
		stage.OnInitCommitCallback = BeforeCommitImplementation

		stage.SetDeltaMode(true)
		stage.ComputeReferenceAndOrders()
	} else {
		err := stage.ParseAstEmbeddedFile(goDiagramsDir, "models/diagrams/diagrams_set.go")
		if err != nil {
			err = stage.ParseAstEmbeddedFile(goDiagramsDir, "diagrams/diagrams_set.go")
		}
		if err != nil {
			log.Println("no file to read " + err.Error())
		}
	}

	treeStage, _ = tree_fullstack.NewStackInstance(r, docStackName+":doc-sidebar", "", "")
	svgStage, _ = svg_fullstack.NewStackInstance(r, docStackName+":doc-svg", "", "", "")
	gongStage = gong.NewStage(docStackName + ":doc-gong")
	formStage, _ = form_fullstack.NewStackInstance(r, docStackName+":doc-diagramForm", "", "")
	treeNavigationStage, _ = tree_fullstack.NewStackInstance(r, docStackName+":doc-sidebar-navigation", "", "")

	loadEmbeddedPackages(gongStage, goModelsDir)

	return
}

func PrepareStageSet(
	r *http.ServeMux,
	embeddedDiagrams bool,
	docStackName string,
	metaPackageImports []*models.MetaPackageImport,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	receivingAsSplitArea *split.AsSplitArea,
	map_GongStructName_InstancesNb map[string]int,
) (stager *models.Stager) {
	stage, treeStage, svgStage, gongStage, formStage, treeNavigationStage := prepareStagesSet(
		r,
		embeddedDiagrams,
		docStackName,
		metaPackageImports,
		goModelsDir,
		goDiagramsDir,
	)

	return models.NewStager(
		r,
		receivingAsSplitArea,
		stage,
		treeStage,
		svgStage,
		gongStage,
		formStage,
		treeNavigationStage,
		embeddedDiagrams,
		map_GongStructName_InstancesNb,
	)
}

func PrepareStageSetSplitlite(
	r *http.ServeMux,
	embeddedDiagrams bool,
	docStackName string,
	metaPackageImports []*models.MetaPackageImport,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	receivingAsSplitArea *splitlite.AsSplitArea,
	map_GongStructName_InstancesNb map[string]int,
) (stager *models.Stager) {
	stage, treeStage, svgStage, gongStage, formStage, treeNavigationStage := prepareStagesSet(
		r,
		embeddedDiagrams,
		docStackName,
		metaPackageImports,
		goModelsDir,
		goDiagramsDir,
	)

	return models.NewStagerSplitlite(
		r,
		receivingAsSplitArea,
		stage,
		treeStage,
		svgStage,
		gongStage,
		formStage,
		treeNavigationStage,
		embeddedDiagrams,
		map_GongStructName_InstancesNb,
	)
}


