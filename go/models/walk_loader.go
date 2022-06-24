package models

import (
	"go/constant"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

func WalkLoader(pkg *packages.Package, modelPkg *ModelPkg) {

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
			if named.Obj().Pkg().Path() != modelPkg.PkgPath {
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
