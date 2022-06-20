package models

import (
	"fmt"
	"go/constant"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/tools/go/packages"
)

const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

// Walk parses go files in the `models` directory pointed by relativePathToModel and
// fills up modelPkg with its Gongstruct & Gongenum
//
// if useParse is used, Walk uses go/parser.Parser
//
// Walk leverages go Parser capabilities to fetch identifiers in go files
//
// The algo is in several steps:
// - First pass gather Gongstruct & Gongenums identifiers
// - Second pass parses fields and link them to other Gongstructs
func Walk(relativePathToModel string, modelPkg *ModelPkg, useParse ...bool) {

	var useParser bool
	if len(useParse) > 1 {
		log.Fatal("Too many args to useParse")
	}
	if len(useParse) == 1 {
		useParser = useParse[0]
	}

	directory, err := filepath.Abs(relativePathToModel)
	if err != nil {
		log.Panic("Path does not exist %s ;" + directory)
	}
	// log.Println("Loading package " + directory)

	// pkgParser
	var pkgParser ModelPkg
	_ = pkgParser
	if useParser {
		fset := token.NewFileSet()
		startParser := time.Now()
		pkgsParser, errParser := parser.ParseDir(fset, directory, nil, parser.ParseComments)
		log.Printf("Parser took %s", time.Since(startParser))

		if errParser != nil {
			log.Panic("Unable to paser ")
		}
		if len(pkgsParser) != 1 {
			log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
		}

		WalkParser(pkgsParser, &pkgParser)
	}

	//
	// prepare package load
	//
	cfg := &packages.Config{
		Dir:   directory,
		Mode:  pkgLoadMode,
		Tests: false,
	}
	start := time.Now()
	pkgs, err := packages.Load(cfg, "./...")
	log.Printf("package Load took %s", time.Since(start))

	if err != nil {
		s := fmt.Sprintf("cannot process package at path %s, err %s", relativePathToModel, err.Error())
		log.Panic(s)
	}

	if len(pkgs) != 1 {
		log.Panicf("Expected 1 package to scope, found %d", len(pkgs))
	}
	pkg := pkgs[0]

	// compute root package path name
	PkgGoPath = pkg.PkgPath
	PkgName = filepath.Base(filepath.Join(pkg.PkgPath, "../.."))

	modelPkg.Name = PkgName
	modelPkg.PkgPath = pkg.PkgPath
	modelPkg.GongEnums = make(map[string]*GongEnum)
	modelPkg.GongStructs = make(map[string]*GongStruct)

	if useParser {
		if pkgParser.Name != modelPkg.Name {
			log.Fatal("Parsing problem")
		}
		if pkgParser.PkgPath != modelPkg.PkgPath {
			log.Fatal("Parsing problem")
		}

	}

	// fetch all pkg names ...
	scope := pkg.Types.Scope()

	// Traverse is in 2 steps
	//
	// first traverse is to gather all gong types
	//
	// second taverse again all gong types fields and link them together
	//

	//
	// First traverse
	//
	for _, eltName := range scope.Names() {

		obj := scope.Lookup(eltName)
		// log.Printf("obj name is %s", obj.Name())

		//
		// collect gong Enum from go const declaration
		//
		switch obj.(type) {
		case *types.TypeName:
			// log.Printf("obj is a Type declation %s", obj.Name())

		// a types.Const is a gong Enum
		case *types.Const:
			cst, _ := obj.(*types.Const)

			// get the type of the const
			var named *types.Named
			var ok bool
			if named, ok = cst.Type().(*types.Named); !ok {
				// it must be some sort of other const
				continue
			}

			// log.Printf("%s is a Const declation with type %s of package path %s", cst.Name(), named.Obj().Name(), named.Obj().Pkg().Path())

			// if type of the const is not a gong type (a type of the package), do not take into account
			if named.Obj().Pkg().Path() != PkgGoPath {
				continue
			}

			// fetch the enum, if it does not exist, create it
			var modelEnum *GongEnum
			if modelEnum = modelPkg.GongEnums[modelPkg.PkgPath+"."+named.Obj().Name()]; modelEnum == nil {

				var enumType GongEnumType
				if cst.Val().Kind() == constant.Kind(constant.Int) {
					enumType = Int
				} else {
					enumType = String
				}

				modelEnum = &GongEnum{
					Name: named.Obj().Name(),
					Type: enumType,
				}
				modelPkg.GongEnums[modelPkg.PkgPath+"."+named.Obj().Name()] = modelEnum
			}

			modelEnum.GongEnumValues = append(modelEnum.GongEnumValues, &GongEnumValue{
				Name:  cst.Name(),
				Value: cst.Val().String(),
			})

		default:
			// we are not interested
			continue
		}

		//
		// collect gong Struct from go Struct
		//

		// check can be of type *type.Struct
		underlying := obj.Type().Underlying()

		switch underlying.(type) {

		// we are only interested in struct
		case *types.Struct:
			// longName := obj.Type().String()
			// log.Println("name : " + longName)

			cmt, hasComments := FindComments(pkg, obj.Name())
			if !hasComments {
				// log.Printf("no comment")
			} else {

				// do not generage something for struct with swwager:ignore
				if strings.Contains(cmt.Text(), "swagger:ignore") {
					// log.Printf("swagger:ignore \n\n%s\n", cmt.Text())
					continue
				}
			}

			modelPkg.GongStructs[modelPkg.PkgPath+"."+eltName] = &GongStruct{
				Name: eltName,
			}

		case *types.Basic:
			// probably a struct
			// log.Printf("Detected a typedef with basic underlying %s\n", obj.Type().String())
		default:
		}
	}

	//
	// second traverse
	//
	for _, eltName := range scope.Names() {

		// fetch object
		obj := scope.Lookup(eltName)

		switch obj.(type) {
		case *types.TypeName:
			// log.Printf("obj is a Type declation, therefore a Struct, hence %s", obj.Name())

		default:
			// we are not interested
			continue
		}

		underlying := obj.Type().Underlying()

		switch underlyingType := underlying.(type) {

		case *types.Named:
			// log.Printf("named %s\n", underlyingType.String())

		case *types.Signature:
			// log.Printf("signature %s\n", underlyingType.String())

		case *types.Tuple:
			// log.Printf("tuple %s\n", underlyingType.String())

		case *types.Struct:
			__struct := underlying.(*types.Struct)

			cmt, hasComments := FindComments(pkg, obj.Name())
			if !hasComments {
				// log.Printf("no comment")
			} else {

				// do not generage something for struct with swwager:ignore
				if strings.Contains(cmt.Text(), "swagger:ignore") {
					// log.Printf("swagger:ignore \n\n%s\n", cmt.Text())
					continue
				}
			}

			structName := obj.Type().String()
			GenerateFields(structName, __struct, pkg, modelPkg, false, "")

		default:
			_ = underlyingType
		}
	}
	// log.Printf("%T", modelPkg)
}
