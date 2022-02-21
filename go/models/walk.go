package models

import (
	"fmt"
	"go/types"
	"log"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

// PkgGoPath for generation
var PkgGoPath string

// PkgName is generated package name (for back) and generated project elements for front
var PkgName string

// PathToGoSubDirectory for instance "/tmp"
var PathToGoSubDirectory string

// OrmPkgGenPath is target path for orm package, for instance "/tmp/libraryorm"
var OrmPkgGenPath string

// ApiPkgGenPath is target path for api package
var ApiPkgGenPath string

// ControllersPkgGenPath is target path for controllers package, for instance "/tmp/librarycontrollers"
var ControllersPkgGenPath string

// // ModulesTargetPath is where ng modules are generated
// var ModulesTargetPath string

// MatTargetPath is where the ng components are generated
var MatTargetPath string

// NgWorkspacePath is the path to the Ng Workspace
var NgWorkspacePath string

// ADDR is the network address addr where the angular generated service will lookup the server
var ADDR string

// Walk parse structs
func Walk(relativePathToModel string, modelPkg *ModelPkg) {

	directory, err := filepath.Abs(relativePathToModel)
	if err != nil {
		log.Panic("Path does not exist %s ;" + directory)
	}
	log.Println("Loading package " + directory)

	//
	// prepare package load
	//
	cfg := &packages.Config{
		Dir:   directory,
		Mode:  pkgLoadMode,
		Tests: false,
	}
	pkgs, err := packages.Load(cfg, "./...")
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
		log.Printf("obj name is %s", obj.Name())

		//
		// collect gong Enum from go const declaration
		//
		switch obj.(type) {
		case *types.TypeName:
			log.Printf("obj is a Type declation %s", obj.Name())

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

			log.Printf("%s is a Const declation with type %s of package path %s", cst.Name(), named.Obj().Name(), named.Obj().Pkg().Path())

			// if type of the const is not a gong type (a type of the package), do not take into account
			if named.Obj().Pkg().Path() != PkgGoPath {
				continue
			}

			// fetch the enum, if it does not exist, create it
			var modelEnum *GongEnum
			if modelEnum = modelPkg.GongEnums[modelPkg.PkgPath+"."+named.Obj().Name()]; modelEnum == nil {
				modelEnum = &GongEnum{
					Name: named.Obj().Name(),
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
			longName := obj.Type().String()
			log.Println("name : " + longName)

			cmt, hasComments := FindComments(pkg, obj.Name())
			if !hasComments {
				// log.Printf("no comment")
			} else {

				// do not generage something for struct with swwager:ignore
				if strings.Contains(cmt.Text(), "swagger:ignore") {
					log.Printf("swagger:ignore \n\n%s\n", cmt.Text())
					continue
				}
			}

			modelPkg.GongStructs[modelPkg.PkgPath+"."+eltName] = &GongStruct{
				Name: eltName,
			}

		case *types.Basic:
			// probably a struct
			log.Printf("Detected a typedef with basic underlying %s\n", obj.Type().String())
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
			log.Printf("obj is a Type declation, therefore a Struct, hence %s", obj.Name())

		default:
			// we are not interested
			continue
		}

		underlying := obj.Type().Underlying()

		switch underlyingType := underlying.(type) {

		case *types.Named:
			log.Printf("named %s\n", underlyingType.String())

		case *types.Signature:
			log.Printf("signature %s\n", underlyingType.String())

		case *types.Tuple:
			log.Printf("tuple %s\n", underlyingType.String())

		case *types.Struct:
			__struct := underlying.(*types.Struct)

			cmt, hasComments := FindComments(pkg, obj.Name())
			if !hasComments {
				// log.Printf("no comment")
			} else {

				// do not generage something for struct with swwager:ignore
				if strings.Contains(cmt.Text(), "swagger:ignore") {
					log.Printf("swagger:ignore \n\n%s\n", cmt.Text())
					continue
				}
			}

			structName := obj.Type().String()
			GenerateFields(structName, __struct, pkg, modelPkg)

		default:
		}
	}
	log.Printf("%T", modelPkg)
}
