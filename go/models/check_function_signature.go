package models

import "go/ast"

func checkFunctionSignature(file *ast.File, modelPkg *ModelPkg) {
	targetFuncNameOnAfterUpdate := "OnAfterUpdate"

	for i, decl := range file.Decls {
		_ = i
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if funcDecl.Name.Name == targetFuncNameOnAfterUpdate {
			recv := funcDecl.Recv.List[0]
			ptrType, ok := recv.Type.(*ast.StarExpr)
			if !ok {
				continue
			}

			structType, ok := ptrType.X.(*ast.Ident)
			if !ok {
				continue
			}

			// is the receiver signature a match for a gong struct ?
			gongstruct, ok := modelPkg.GongStructs[modelPkg.PkgPath+"."+structType.Name]

			if !ok {
				continue
			}

			params := funcDecl.Type.Params.List

			if len(params) < 2 {
				continue
			}

			param1 := params[0].Type.(*ast.StarExpr)
			param2 := params[1].Type.(*ast.StarExpr)

			if funcDecl.Name.Name == targetFuncNameOnAfterUpdate &&
				param1.X.(*ast.Ident).Name == "Stage" &&
				param2.X.(*ast.Ident).Name == gongstruct.Name {

				gongstruct.HasOnAfterUpdateSignature = true
			}
		}
	}
}
