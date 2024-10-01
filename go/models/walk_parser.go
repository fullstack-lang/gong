package models

import (
	"embed"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"math"
	"path/filepath"
	"slices"
	"strings"
	"unicode"

	"github.com/fullstack-lang/gong/go/ignore"
	// to parse the .frontignore file
)

func ParseEmbedModel(embeddedDir embed.FS, source string) map[string]*ast.Package {

	pkg := new(ast.Package)
	pkg.Files = make(map[string]*ast.File)
	fs.WalkDir(embeddedDir, source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".go" {
			return nil
		}

		var data, err1 = embeddedDir.ReadFile(path)
		if err1 != nil {
			log.Fatalln(err.Error())
		}
		fset := token.NewFileSet()
		astFile, errParser := parser.ParseFile(fset, path, data, parser.ParseComments)

		pkg.Files[path] = astFile
		if errParser != nil {
			panic(errParser)
		}
		_ = astFile
		return nil

	})
	pkgs := make(map[string]*ast.Package)
	sourcesElements := strings.Split(source, "/")
	pkgName := sourcesElements[len(sourcesElements)-1]
	pkgs[pkgName] = pkg
	return pkgs
}

func WalkParser(parserPkgs map[string]*ast.Package, modelPkg *ModelPkg, goGitignoreEntries *[]ignore.GitignoreEntry) {

	// this is to store struct that are not gongstruct
	// but that can be embedded
	map_Structname_fieldList := make(map[string]*[]*ast.Field)

	var astPackage *ast.Package
	var ok bool
	if astPackage, ok = parserPkgs["models"]; !ok {
		log.Fatal("No package models")
	}

	modelPkg.GongEnums = make(map[string]*GongEnum)
	modelPkg.GongStructs = make(map[string]*GongStruct)
	modelPkg.GongNotes = make(map[string]*GongNote)

	if len(astPackage.Files) == 0 {
		log.Fatal("No go file to parse")
	}

	// parses all comments in the package
	typeDocumentation := doc.New(astPackage, "./", doc.PreserveAST)
	modelPkg.GenerateDocs(typeDocumentation)

	map_StructName_hasIgnoreStatement := make(map[string]bool)
	for _, t := range typeDocumentation.Types {
		map_StructName_hasIgnoreStatement[t.Name] = strings.Contains(t.Doc, "swagger:ignore") || strings.Contains(t.Doc, "gong:ignore")
	}

	// in astPackage.Files is the map of filePath to file
	// the filePath can be absolute of relative
	// in order to compute at what level to the "models" directory we are
	// we need to process all filePath and get the distance to the "models" directory
	// given that "models" directory name is forbiden in the path
	minFilePathLength := math.MaxInt
	for filePath := range astPackage.Files {
		// get directories to the file
		directories := make([]string, 0)
		workingFilePath := filePath
		for {
			dir := filepath.Dir(workingFilePath)
			if dir == workingFilePath {
				break
			}
			directories = append(directories, dir)
			workingFilePath = dir
		}

		if len(directories) < minFilePathLength {
			minFilePathLength = len(directories)
		}
	}

	// pre pass to identity all struct with a NameField
	mapStructWithNameField := make(map[string]any)
	for _, file := range astPackage.Files {
		for _, decl := range file.Decls {
			switch genDecl := decl.(type) {
			case *ast.GenDecl:
				for _, spec := range genDecl.Specs {
					switch typeSpec := spec.(type) {
					case *ast.TypeSpec:
						switch _type := typeSpec.Type.(type) {
						case *ast.StructType:
							hasName := IsNamedStructWithoutEmbedded(_type, nil)
							if hasName {
								mapStructWithNameField[typeSpec.Name.Name] = true
							}
							// log.Println("file ", typeSpec.Name.Name, hasName)
						}

					}
				}
			}
		}
	}

	// first pass : get "type" definition for enum & struct
	//
	// search all files
	for filePath, file := range astPackage.Files {

		var isFileFrontIgnored bool
		fileName := filepath.Base(filePath)

		// get directories to the file
		directories := make([]string, 0)
		workingFilePath := filePath
		for {
			dir := filepath.Dir(workingFilePath)
			if dir == workingFilePath {
				break
			}
			directories = append(directories, dir)
			workingFilePath = dir
		}

		// we do not take the files in the sub directories (yet)
		if len(directories) > minFilePathLength {
			continue
		}

		if slices.Contains(GeneratedModelFiles, filepath.Base(filePath)) {
			continue
		}

		if goGitignoreEntries != nil && ignore.CheckFileMatches(fileName, *goGitignoreEntries) {
			isFileFrontIgnored = true
		}

		// for exploration
		// if fileName != "estruct.go" {
		// 	continue
		// }
		// if fileName != "cenum_int.go" {
		// 	continue
		// }

		// log.Println("Parsing file ", fileName)

		// for structName, scope := range file.Scope.Objects {

		// 	log.Println("Parsing structName ", structName)
		// 	_ = scope
		// }

		// first pass for gathering the "type" definitions
		for _, decl := range file.Decls {

			switch genDecl := decl.(type) {
			case *ast.GenDecl:
				// log.Println("Length of Spec", len(genDecl.Specs))

				for _, spec := range genDecl.Specs {

					switch typeSpec := spec.(type) {
					case *ast.TypeSpec:

						// TypeSpec is for declaration with the "type" keyword
						// in gong, it can be either a GongStruct or a GongEnum

						// we are only interested in exported symbols
						if !unicode.IsUpper(rune(typeSpec.Name.Name[0])) {
							continue
						}

						// log.Println("Type spec name is ", typeSpec.Name.Name)

						// If it is a GongEnum, the typeSpec Type is a int or a string
						switch _type := typeSpec.Type.(type) {
						case *ast.Ident:
							// log.Println("Type spec type is ", _type.Name)

							var enumType GongEnumType
							switch _type.Name {
							case "int":
								enumType = Int
							case "string":
								enumType = String
							default:
								continue
							}

							hasIgnoreStatement := map_StructName_hasIgnoreStatement[typeSpec.Name.Name]
							if hasIgnoreStatement {
								continue
							}

							modelEnum := &GongEnum{
								Name: typeSpec.Name.Name,
								Type: enumType,
							}
							modelPkg.GongEnums[modelPkg.PkgPath+"."+typeSpec.Name.Name] = modelEnum

							// If it is a GongStruct, the ast is a StructType
						case *ast.StructType:

							// fetch the name of the Gongstruct by identifying if there is a field with name "Name"
							IsNamedStruct := IsNamedStructWithoutEmbedded(_type, mapStructWithNameField)
							hasIgnoreStatement := map_StructName_hasIgnoreStatement[typeSpec.Name.Name]

							if IsNamedStruct && !hasIgnoreStatement {
								gongstruct := (&GongStruct{
									Name:              typeSpec.Name.Name,
									IsIgnoredForFront: isFileFrontIgnored}).
									Stage(modelPkg.GetStage())
								modelPkg.GongStructs[modelPkg.PkgPath+"."+typeSpec.Name.Name] = gongstruct
							}
							if !hasIgnoreStatement {
								map_Structname_fieldList[typeSpec.Name.Name] = &_type.Fields.List
							}
						default:
						}
					default:
					}
				}

			default:
				continue
			}
		}

	}

	// second pass
	for filePath, file := range astPackage.Files {

		if slices.Contains(GeneratedModelFiles, filepath.Base(filePath)) {
			continue
		}

		// if fileName != "astruct.go" {
		// 	continue
		// }

		for _, decl := range file.Decls {
			_ = decl

			switch genDecl := decl.(type) {
			case *ast.GenDecl:

				var gongEnum *GongEnum

				for enumIndex, _spec := range genDecl.Specs {

					switch spec := _spec.(type) {
					// values of a enum
					case *ast.ValueSpec:
						if len(spec.Names) > 0 {
							if len(spec.Names) > 1 {
								log.Fatal("To many Names for value spec ", spec.Names[0].Name)
							}
						}

						// for an Enum, the type is only defined at the first element
						switch _type := spec.Type.(type) {
						case *ast.Ident:
							gongEnum, ok = modelPkg.GongEnums[modelPkg.PkgPath+"."+_type.Name]
							if !ok {
								// log.Println("Constant ", spec.Names[0], "of Type", _type.Name, " not an enum")
								continue
							}
							// log.Println("Const ", spec.Names[0].Name, " of type ", gongEnum.Name)

							if gongEnum.Type == Int {
								gongEnumValue := (&GongEnumValue{
									Name:  spec.Names[0].Name,
									Value: fmt.Sprintf("%d", enumIndex)})
								gongEnum.GongEnumValues = append(gongEnum.GongEnumValues, gongEnumValue)
							} else { // string
								// the value is in the BasicLit expression
								if len(spec.Values) != 1 {
									log.Fatal("Wrong def of const ", spec.Names[0].Name)
								}
								_value := spec.Values[0]

								switch value := _value.(type) {
								case *ast.BasicLit:
									gongEnumValue := (&GongEnumValue{
										Name:  spec.Names[0].Name,
										Value: value.Value})
									gongEnum.GongEnumValues = append(gongEnum.GongEnumValues, gongEnumValue)
								default:
									log.Fatal("Wrong def of const ", spec.Names[0].Name)
								}
							}
						default:
						}

						// if gongEnum is nil, this is because of a non Gongstruct const
						if spec.Type == nil && gongEnum != nil {
							gongEnumValue := (&GongEnumValue{
								Name:  spec.Names[0].Name,
								Value: fmt.Sprintf("%d", enumIndex)})
							gongEnum.GongEnumValues = append(gongEnum.GongEnumValues, gongEnumValue)
						}
					case *ast.TypeSpec:

						// TypeSpec is for declaration with the "type" keyword
						// in gong, it can be either a GongStruct or a GongEnum
						switch _type := spec.Type.(type) {
						case *ast.StructType:

							// log.Println("Parsing fields of gongstruct ", spec.Name.Name)

							// fetch the name of the Gongstruct by identifying if there is a field with name "Name"
							IsNamedStruct := IsNamedStructWithoutEmbedded(_type, mapStructWithNameField)
							hasIgnoreStatement := map_StructName_hasIgnoreStatement[spec.Name.Name]
							if IsNamedStruct && !hasIgnoreStatement {
								gongstruct, ok := modelPkg.GongStructs[modelPkg.PkgPath+"."+spec.Name.Name]
								structName := spec.Name.Name
								if !ok {
									log.Fatalln("Unkown struct ", structName)
								}
								_ = gongstruct

								GenerateFieldParser(
									&_type.Fields.List,
									gongstruct,
									&map_Structname_fieldList,
									modelPkg, "", "")
							}
						}
					default:
					}
				}
			default:
				continue
			}
		}
	}

	// pass to detect if there a function that matches the OnAfterUpdate signature
	for filePath, file := range astPackage.Files {

		if slices.Contains(GeneratedModelFiles, filepath.Base(filePath)) {
			continue
		}

		checkFunctionSignature(file, modelPkg)
	}
}
