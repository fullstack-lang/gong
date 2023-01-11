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
	map_DocLink_Renaming := make(map[string]string, 0)

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
								&map_DocLink_Renaming,
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
					if ident.Name != "map_DocLink_Identifier" {
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
							// file
							if docLink == key {
								continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							map_DocLink_Renaming[key] = docLink
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
var __gong__map_Astruct = make(map[string]*Astruct)
var __gong__map_AstructBstruct2Use = make(map[string]*AstructBstruct2Use)
var __gong__map_AstructBstructUse = make(map[string]*AstructBstructUse)
var __gong__map_Bstruct = make(map[string]*Bstruct)
var __gong__map_Dstruct = make(map[string]*Dstruct)

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
func UnmarshallGongstructStaging(map_DocLink_Renaming *map[string]string, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
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
						if renamed, ok := (*map_DocLink_Renaming)[docLinkText]; ok {
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
									case "Astruct":
										instanceAstruct := (&Astruct{Name: instanceName}).Stage()
										instance = any(instanceAstruct)
										__gong__map_Astruct[identifier] = instanceAstruct
									case "AstructBstruct2Use":
										instanceAstructBstruct2Use := (&AstructBstruct2Use{Name: instanceName}).Stage()
										instance = any(instanceAstructBstruct2Use)
										__gong__map_AstructBstruct2Use[identifier] = instanceAstructBstruct2Use
									case "AstructBstructUse":
										instanceAstructBstructUse := (&AstructBstructUse{Name: instanceName}).Stage()
										instance = any(instanceAstructBstructUse)
										__gong__map_AstructBstructUse[identifier] = instanceAstructBstructUse
									case "Bstruct":
										instanceBstruct := (&Bstruct{Name: instanceName}).Stage()
										instance = any(instanceBstruct)
										__gong__map_Bstruct[identifier] = instanceBstruct
									case "Dstruct":
										instanceDstruct := (&Dstruct{Name: instanceName}).Stage()
										instance = any(instanceDstruct)
										__gong__map_Dstruct[identifier] = instanceDstruct
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
						case "Astruct":
							switch fieldName {
							// insertion point for date assign code
							case "Date":
								__gong__map_Astruct[identifier].Date, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						case "AstructBstruct2Use":
							switch fieldName {
							// insertion point for date assign code
							}
						case "AstructBstructUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Bstruct":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Dstruct":
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
					case "Astruct":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Anarrayofb":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Bstruct[targetIdentifier]
							__gong__map_Astruct[identifier].Anarrayofb =
								append(__gong__map_Astruct[identifier].Anarrayofb, target)
						case "Anotherarrayofb":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Bstruct[targetIdentifier]
							__gong__map_Astruct[identifier].Anotherarrayofb =
								append(__gong__map_Astruct[identifier].Anotherarrayofb, target)
						case "Anarrayofa":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Astruct[targetIdentifier]
							__gong__map_Astruct[identifier].Anarrayofa =
								append(__gong__map_Astruct[identifier].Anarrayofa, target)
						case "AnarrayofbUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_AstructBstructUse[targetIdentifier]
							__gong__map_Astruct[identifier].AnarrayofbUse =
								append(__gong__map_Astruct[identifier].AnarrayofbUse, target)
						case "Anarrayofb2Use":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_AstructBstruct2Use[targetIdentifier]
							__gong__map_Astruct[identifier].Anarrayofb2Use =
								append(__gong__map_Astruct[identifier].Anarrayofb2Use, target)
						}
					case "AstructBstruct2Use":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "AstructBstructUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Bstruct":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Dstruct":
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
			case "Astruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Astruct[identifier].Name = fielValue
				case "CName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Astruct[identifier].CName = fielValue
				case "CFloatfield":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].CFloatfield = fielValue
				case "Floatfield":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Floatfield = fielValue
				case "Intfield":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Intfield = int(fielValue)
				case "Duration1":
					// convert string to duration
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Duration1 = time.Duration(fielValue)
				case "StructRef":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Astruct[identifier].StructRef = fielValue
				case "FieldRef":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Astruct[identifier].FieldRef = fielValue
				}
			case "AstructBstruct2Use":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AstructBstruct2Use[identifier].Name = fielValue
				}
			case "AstructBstructUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AstructBstructUse[identifier].Name = fielValue
				}
			case "Bstruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bstruct[identifier].Name = fielValue
				case "Floatfield":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bstruct[identifier].Floatfield = fielValue
				case "Floatfield2":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bstruct[identifier].Floatfield2 = fielValue
				case "Intfield":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bstruct[identifier].Intfield = int(fielValue)
				}
			case "Dstruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Dstruct[identifier].Name = fielValue
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
			case "Astruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Booleanfield":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Booleanfield = fielValue
				case "Bstruct":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Bstruct = __gong__map_Bstruct[targetIdentifier]
				case "Bstruct2":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Bstruct2 = __gong__map_Bstruct[targetIdentifier]
				case "Dstruct":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Dstruct = __gong__map_Dstruct[targetIdentifier]
				case "Dstruct2":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Dstruct2 = __gong__map_Dstruct[targetIdentifier]
				case "Dstruct3":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Dstruct3 = __gong__map_Dstruct[targetIdentifier]
				case "Dstruct4":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Dstruct4 = __gong__map_Dstruct[targetIdentifier]
				case "Anotherbooleanfield":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Anotherbooleanfield = fielValue
				case "Associationtob":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Associationtob = __gong__map_Bstruct[targetIdentifier]
				case "Anotherassociationtob_2":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Anotherassociationtob_2 = __gong__map_Bstruct[targetIdentifier]
				case "AnAstruct":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].AnAstruct = __gong__map_Astruct[targetIdentifier]
				}
			case "AstructBstruct2Use":
				switch fieldName {
				// insertion point for field dependant code
				case "Bstrcut2":
					targetIdentifier := ident.Name
					__gong__map_AstructBstruct2Use[identifier].Bstrcut2 = __gong__map_Bstruct[targetIdentifier]
				}
			case "AstructBstructUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Bstruct2":
					targetIdentifier := ident.Name
					__gong__map_AstructBstructUse[identifier].Bstruct2 = __gong__map_Bstruct[targetIdentifier]
				}
			case "Bstruct":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Dstruct":
				switch fieldName {
				// insertion point for field dependant code
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
				case "Astruct":
					switch fieldName {
					// insertion point for enum assign code
					case "Aenum":
						var val AEnumType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Astruct[identifier].Aenum = AEnumType(val)
					case "Aenum_2":
						var val AEnumType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Astruct[identifier].Aenum_2 = AEnumType(val)
					case "Benum":
						var val BEnumType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Astruct[identifier].Benum = BEnumType(val)
					case "CEnum":
						var val CEnumTypeInt
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Astruct[identifier].CEnum = CEnumTypeInt(val)
					}
				case "AstructBstruct2Use":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "AstructBstructUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Bstruct":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Dstruct":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}
