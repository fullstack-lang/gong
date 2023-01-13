package models

import (
	"errors"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(pathToFile string) error {

	// map to store renaming docLink
	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	Stage.Map_DocLink_Renaming = make(map[string]string, 0)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", fileOfInterest)
	}
	stage := &Stage
	_ = stage
	if len(inFile.Imports) == 3 {
		Stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		Stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "Injection")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						// Create an ast.CommentMap from the ast.File's comments.
						// This helps keeping the association between comments
						// and AST nodes.
						cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								&cmap, assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				case *ast.ValueSpec:
					ident := spec.Names[0]
					_ = ident
					if !strings.HasPrefix(ident.Name, "map_DocLink_Identifier") {
						continue
					}
					switch compLit := spec.Values[0].(type) {
					case *ast.CompositeLit:
						var key string
						_ = key
						var value string
						_ = value
						for _, elt := range compLit.Elts {

							// each elt is an expression for struct or for field such as
							// for struct
							//
							//         "dummy.Dummy": &(dummy.Dummy{})
							//
							// or, for field
							//
							//          "dummy.Dummy.Name": (dummy.Dummy{}).Name,
							//
							// first node in the AST is a key value expression
							var ok bool
							var kve *ast.KeyValueExpr
							if kve, ok = elt.(*ast.KeyValueExpr); !ok {
								log.Fatal("Expression should be key value expression" +
									fset.Position(kve.Pos()).String())
							}

							switch bl := kve.Key.(type) {
							case *ast.BasicLit:
								key = bl.Value // "\"dumm.Dummy\"" is the value

								// one remove the ambracing double quotes
								key = strings.TrimPrefix(key, "\"")
								key = strings.TrimSuffix(key, "\"")
							}

							var isFieldEntry bool
							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								isFieldEntry = true
							}

							var se2 *ast.SelectorExpr
							if isFieldEntry {
								if se2, ok = kve.Value.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(kve.Pos()).String())
								}
								fieldName = se2.Sel.Name
							}

							var pe *ast.ParenExpr
							if !isFieldEntry {
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							} else {
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}

							// expect a Composite Litteral with no Element <type>{}
							var cl *ast.CompositeLit
							if cl, ok = pe.X.(*ast.CompositeLit); !ok {
								log.Fatal("Expression should be a composite lit" +
									fset.Position(pe.Pos()).String())
							}

							var se *ast.SelectorExpr
							if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
								log.Fatal("Expression should be a selector" +
									fset.Position(cl.Pos()).String())
							}

							var id *ast.Ident
							if id, ok = se.X.(*ast.Ident); !ok {
								log.Fatal("Expression should be an ident" +
									fset.Position(se.Pos()).String())
							}
							docLink := id.Name + "." + se.Sel.Name

							if isFieldEntry {
								docLink += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							Stage.Map_DocLink_Renaming[key] = docLink
						}
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_GongBasicField = make(map[string]*GongBasicField)
var __gong__map_GongEnum = make(map[string]*GongEnum)
var __gong__map_GongEnumValue = make(map[string]*GongEnumValue)
var __gong__map_GongLink = make(map[string]*GongLink)
var __gong__map_GongNote = make(map[string]*GongNote)
var __gong__map_GongStruct = make(map[string]*GongStruct)
var __gong__map_GongTimeField = make(map[string]*GongTimeField)
var __gong__map_Meta = make(map[string]*Meta)
var __gong__map_MetaReference = make(map[string]*MetaReference)
var __gong__map_ModelPkg = make(map[string]*ModelPkg)
var __gong__map_PointerToGongStructField = make(map[string]*PointerToGongStructField)
var __gong__map_SliceOfPointerToGongStructField = make(map[string]*SliceOfPointerToGongStructField)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	if name == Stage.MetaPackageImportAlias {
		return Stage.MetaPackageImportAlias, true
	}
	return comment.DefaultLookupPackage(name)
}
func lookupSym(recv, name string) (ok bool) {
	if recv == "" {
		return true
	}
	return false
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {

	// used for debug purposes
	astCoordinate := "\tAssignStmt: "

	//
	// First parse all comment groups in the assignement
	// if a comment "//gong:ident [DocLink]" is met and is followed by a string assignement.
	// modify the following AST assignement to assigns the DocLink text to the string value
	//

	// get the comment group of the assigStmt
	commentGroups := (*cmap)[assignStmt]
	// get the the prefix
	var hasGongIdentDirective bool
	var commentText string
	var docLinkText string
	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "//gong:ident") {
				hasGongIdentDirective = true
				commentText = comment.Text
			}
		}
	}
	if hasGongIdentDirective {
		// parser configured to find doclinks
		var docLinkFinder comment.Parser
		docLinkFinder.LookupPackage = lookupPackage
		docLinkFinder.LookupSym = lookupSym
		doc := docLinkFinder.Parse(commentText)

		for _, block := range doc.Content {
			switch paragraph := block.(type) {
			case *comment.Paragraph:
				_ = paragraph
				for _, text := range paragraph.Text {
					switch docLink := text.(type) {
					case *comment.DocLink:
						if docLink.Recv == "" {
							docLinkText = docLink.ImportPath + "." + docLink.Name
						} else {
							docLinkText = docLink.ImportPath + "." + docLink.Recv + "." + docLink.Name
						}

						// we check wether the doc link has been renamed
						// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
						if renamed, ok := (Stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed
						}
					}
				}
			}
		}
	}

	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "GongBasicField":
										instanceGongBasicField := (&GongBasicField{Name: instanceName}).Stage()
										instance = any(instanceGongBasicField)
										__gong__map_GongBasicField[identifier] = instanceGongBasicField
									case "GongEnum":
										instanceGongEnum := (&GongEnum{Name: instanceName}).Stage()
										instance = any(instanceGongEnum)
										__gong__map_GongEnum[identifier] = instanceGongEnum
									case "GongEnumValue":
										instanceGongEnumValue := (&GongEnumValue{Name: instanceName}).Stage()
										instance = any(instanceGongEnumValue)
										__gong__map_GongEnumValue[identifier] = instanceGongEnumValue
									case "GongLink":
										instanceGongLink := (&GongLink{Name: instanceName}).Stage()
										instance = any(instanceGongLink)
										__gong__map_GongLink[identifier] = instanceGongLink
									case "GongNote":
										instanceGongNote := (&GongNote{Name: instanceName}).Stage()
										instance = any(instanceGongNote)
										__gong__map_GongNote[identifier] = instanceGongNote
									case "GongStruct":
										instanceGongStruct := (&GongStruct{Name: instanceName}).Stage()
										instance = any(instanceGongStruct)
										__gong__map_GongStruct[identifier] = instanceGongStruct
									case "GongTimeField":
										instanceGongTimeField := (&GongTimeField{Name: instanceName}).Stage()
										instance = any(instanceGongTimeField)
										__gong__map_GongTimeField[identifier] = instanceGongTimeField
									case "Meta":
										instanceMeta := (&Meta{Name: instanceName}).Stage()
										instance = any(instanceMeta)
										__gong__map_Meta[identifier] = instanceMeta
									case "MetaReference":
										instanceMetaReference := (&MetaReference{Name: instanceName}).Stage()
										instance = any(instanceMetaReference)
										__gong__map_MetaReference[identifier] = instanceMetaReference
									case "ModelPkg":
										instanceModelPkg := (&ModelPkg{Name: instanceName}).Stage()
										instance = any(instanceModelPkg)
										__gong__map_ModelPkg[identifier] = instanceModelPkg
									case "PointerToGongStructField":
										instancePointerToGongStructField := (&PointerToGongStructField{Name: instanceName}).Stage()
										instance = any(instancePointerToGongStructField)
										__gong__map_PointerToGongStructField[identifier] = instancePointerToGongStructField
									case "SliceOfPointerToGongStructField":
										instanceSliceOfPointerToGongStructField := (&SliceOfPointerToGongStructField{Name: instanceName}).Stage()
										instance = any(instanceSliceOfPointerToGongStructField)
										__gong__map_SliceOfPointerToGongStructField[identifier] = instanceSliceOfPointerToGongStructField
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "GongBasicField":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongEnum":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongEnumValue":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongLink":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongNote":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongStruct":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongTimeField":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Meta":
							switch fieldName {
							// insertion point for date assign code
							}
						case "MetaReference":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ModelPkg":
							switch fieldName {
							// insertion point for date assign code
							}
						case "PointerToGongStructField":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SliceOfPointerToGongStructField":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "GongBasicField":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GongEnum":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "GongEnumValues":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongEnumValue[targetIdentifier]
							__gong__map_GongEnum[identifier].GongEnumValues =
								append(__gong__map_GongEnum[identifier].GongEnumValues, target)
						}
					case "GongEnumValue":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GongLink":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GongNote":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Links":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongLink[targetIdentifier]
							__gong__map_GongNote[identifier].Links =
								append(__gong__map_GongNote[identifier].Links, target)
						}
					case "GongStruct":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "GongBasicFields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongBasicField[targetIdentifier]
							__gong__map_GongStruct[identifier].GongBasicFields =
								append(__gong__map_GongStruct[identifier].GongBasicFields, target)
						case "GongTimeFields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongTimeField[targetIdentifier]
							__gong__map_GongStruct[identifier].GongTimeFields =
								append(__gong__map_GongStruct[identifier].GongTimeFields, target)
						case "PointerToGongStructFields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_PointerToGongStructField[targetIdentifier]
							__gong__map_GongStruct[identifier].PointerToGongStructFields =
								append(__gong__map_GongStruct[identifier].PointerToGongStructFields, target)
						case "SliceOfPointerToGongStructFields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SliceOfPointerToGongStructField[targetIdentifier]
							__gong__map_GongStruct[identifier].SliceOfPointerToGongStructFields =
								append(__gong__map_GongStruct[identifier].SliceOfPointerToGongStructFields, target)
						}
					case "GongTimeField":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Meta":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "MetaReferences":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_MetaReference[targetIdentifier]
							__gong__map_Meta[identifier].MetaReferences =
								append(__gong__map_Meta[identifier].MetaReferences, target)
						}
					case "MetaReference":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ModelPkg":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "PointerToGongStructField":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SliceOfPointerToGongStructField":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit:
			// assignment to string field
			basicLit := expr
			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "GongBasicField":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongBasicField[identifier].Name = fielValue
				case "BasicKindName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongBasicField[identifier].BasicKindName = fielValue
				case "DeclaredType":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongBasicField[identifier].DeclaredType = fielValue
				case "CompositeStructName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongBasicField[identifier].CompositeStructName = fielValue
				case "Index":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongBasicField[identifier].Index = int(fielValue)
				}
			case "GongEnum":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnum[identifier].Name = fielValue
				}
			case "GongEnumValue":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnumValue[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnumValue[identifier].Value = fielValue
				}
			case "GongLink":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongLink[identifier].Name = fielValue
				case "ImportPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongLink[identifier].ImportPath = fielValue
				}
			case "GongNote":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongNote[identifier].Name = fielValue
				case "Body":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongNote[identifier].Body = fielValue
				}
			case "GongStruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongStruct[identifier].Name = fielValue
				}
			case "GongTimeField":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongTimeField[identifier].Name = fielValue
				case "Index":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongTimeField[identifier].Index = int(fielValue)
				case "CompositeStructName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongTimeField[identifier].CompositeStructName = fielValue
				}
			case "Meta":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Meta[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Meta[identifier].Text = fielValue
				}
			case "MetaReference":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MetaReference[identifier].Name = fielValue
				}
			case "ModelPkg":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ModelPkg[identifier].Name = fielValue
				case "PkgPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ModelPkg[identifier].PkgPath = fielValue
				}
			case "PointerToGongStructField":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_PointerToGongStructField[identifier].Name = fielValue
				case "Index":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_PointerToGongStructField[identifier].Index = int(fielValue)
				case "CompositeStructName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_PointerToGongStructField[identifier].CompositeStructName = fielValue
				}
			case "SliceOfPointerToGongStructField":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SliceOfPointerToGongStructField[identifier].Name = fielValue
				case "Index":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SliceOfPointerToGongStructField[identifier].Index = int(fielValue)
				case "CompositeStructName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SliceOfPointerToGongStructField[identifier].CompositeStructName = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "GongBasicField":
				switch fieldName {
				// insertion point for field dependant code
				case "GongEnum":
					targetIdentifier := ident.Name
					__gong__map_GongBasicField[identifier].GongEnum = __gong__map_GongEnum[targetIdentifier]
				case "IsDocLink":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongBasicField[identifier].IsDocLink = fielValue
				}
			case "GongEnum":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongEnumValue":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongLink":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongNote":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongStruct":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongTimeField":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Meta":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "MetaReference":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ModelPkg":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "PointerToGongStructField":
				switch fieldName {
				// insertion point for field dependant code
				case "GongStruct":
					targetIdentifier := ident.Name
					__gong__map_PointerToGongStructField[identifier].GongStruct = __gong__map_GongStruct[targetIdentifier]
				}
			case "SliceOfPointerToGongStructField":
				switch fieldName {
				// insertion point for field dependant code
				case "GongStruct":
					targetIdentifier := ident.Name
					__gong__map_SliceOfPointerToGongStructField[identifier].GongStruct = __gong__map_GongStruct[targetIdentifier]
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "GongBasicField":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongEnum":
					switch fieldName {
					// insertion point for enum assign code
					case "Type":
						var val GongEnumType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_GongEnum[identifier].Type = GongEnumType(val)
					}
				case "GongEnumValue":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongLink":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongNote":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongStruct":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongTimeField":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Meta":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "MetaReference":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ModelPkg":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "PointerToGongStructField":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SliceOfPointerToGongStructField":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}
