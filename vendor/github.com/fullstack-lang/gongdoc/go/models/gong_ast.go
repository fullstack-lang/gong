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

// swagger:ignore
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(stage *StageStruct, pathToFile string) error {

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

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *StageStruct, inFile *ast.File, fset *token.FileSet) error {
	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", inFile.Name)
	}
	if len(inFile.Imports) == 3 {
		stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
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
								stage, &cmap, assignStmt, astCoordinate)
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
							var expressionType GONG__ExpressionType = GONG__STRUCT_INSTANCE
							var docLink GONG__Identifier

							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								expressionType = GONG__FIELD_OR_CONST_VALUE
							}

							var callExpr *ast.CallExpr
							if callExpr, ok = kve.Value.(*ast.CallExpr); ok {

								var se *ast.SelectorExpr
								if se, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(callExpr.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}

								// check the arg type to select wether this is a int or a string enum
								var bl *ast.BasicLit
								if bl, ok = callExpr.Args[0].(*ast.BasicLit); ok {
									switch bl.Kind {
									case token.STRING:
										expressionType = GONG__ENUM_CAST_STRING
									case token.INT:
										expressionType = GONG__ENUM_CAST_INT
									}
								} else {
									log.Fatal("Expression should be a basic lit" +
										fset.Position(se.Pos()).String())
								}

								docLink.Ident = id.Name + "." + se.Sel.Name
								_ = callExpr
							}

							var se2 *ast.SelectorExpr
							switch expressionType {
							case GONG__FIELD_OR_CONST_VALUE:
								if se2, ok = kve.Value.(*ast.SelectorExpr); ok {

									var ident *ast.Ident
									if _, ok = se2.X.(*ast.ParenExpr); ok {
										expressionType = GONG__FIELD_VALUE
										fieldName = se2.Sel.Name
									} else if ident, ok = se2.X.(*ast.Ident); ok {
										expressionType = GONG__IDENTIFIER_CONST
										docLink.Ident = ident.Name + "." + se2.Sel.Name
									} else {
										log.Fatal("Expression should be a selector expression or an ident" +
											fset.Position(kve.Pos()).String())
									}
								} else {

								}
							}

							var pe *ast.ParenExpr
							switch expressionType {
							case GONG__STRUCT_INSTANCE:
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							case GONG__FIELD_VALUE:
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}
							switch expressionType {
							case GONG__FIELD_VALUE, GONG__STRUCT_INSTANCE:
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
								docLink.Ident = id.Name + "." + se.Sel.Name
							}

							switch expressionType {
							case GONG__FIELD_VALUE:
								docLink.Ident += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink.Ident == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							docLink.Type = expressionType
							stage.Map_DocLink_Renaming[key] = docLink
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
var __gong__map_Classdiagram = make(map[string]*Classdiagram)
var __gong__map_DiagramPackage = make(map[string]*DiagramPackage)
var __gong__map_Field = make(map[string]*Field)
var __gong__map_GongEnumShape = make(map[string]*GongEnumShape)
var __gong__map_GongEnumValueEntry = make(map[string]*GongEnumValueEntry)
var __gong__map_GongStructShape = make(map[string]*GongStructShape)
var __gong__map_Link = make(map[string]*Link)
var __gong__map_NoteShape = make(map[string]*NoteShape)
var __gong__map_NoteShapeLink = make(map[string]*NoteShapeLink)
var __gong__map_Position = make(map[string]*Position)
var __gong__map_UmlState = make(map[string]*UmlState)
var __gong__map_Umlsc = make(map[string]*Umlsc)
var __gong__map_Vertice = make(map[string]*Vertice)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}
func lookupSym(recv, name string) (ok bool) {
	if recv == "" {
		return true
	}
	return false
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *StageStruct, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
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
						if renamed, ok := (stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed.Ident
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
									case "Classdiagram":
										instanceClassdiagram := (&Classdiagram{Name: instanceName}).Stage(stage)
										instance = any(instanceClassdiagram)
										__gong__map_Classdiagram[identifier] = instanceClassdiagram
									case "DiagramPackage":
										instanceDiagramPackage := (&DiagramPackage{Name: instanceName}).Stage(stage)
										instance = any(instanceDiagramPackage)
										__gong__map_DiagramPackage[identifier] = instanceDiagramPackage
									case "Field":
										instanceField := (&Field{Name: instanceName}).Stage(stage)
										instance = any(instanceField)
										__gong__map_Field[identifier] = instanceField
									case "GongEnumShape":
										instanceGongEnumShape := (&GongEnumShape{Name: instanceName}).Stage(stage)
										instance = any(instanceGongEnumShape)
										__gong__map_GongEnumShape[identifier] = instanceGongEnumShape
									case "GongEnumValueEntry":
										instanceGongEnumValueEntry := (&GongEnumValueEntry{Name: instanceName}).Stage(stage)
										instance = any(instanceGongEnumValueEntry)
										__gong__map_GongEnumValueEntry[identifier] = instanceGongEnumValueEntry
									case "GongStructShape":
										instanceGongStructShape := (&GongStructShape{Name: instanceName}).Stage(stage)
										instance = any(instanceGongStructShape)
										__gong__map_GongStructShape[identifier] = instanceGongStructShape
									case "Link":
										instanceLink := (&Link{Name: instanceName}).Stage(stage)
										instance = any(instanceLink)
										__gong__map_Link[identifier] = instanceLink
									case "NoteShape":
										instanceNoteShape := (&NoteShape{Name: instanceName}).Stage(stage)
										instance = any(instanceNoteShape)
										__gong__map_NoteShape[identifier] = instanceNoteShape
									case "NoteShapeLink":
										instanceNoteShapeLink := (&NoteShapeLink{Name: instanceName}).Stage(stage)
										instance = any(instanceNoteShapeLink)
										__gong__map_NoteShapeLink[identifier] = instanceNoteShapeLink
									case "Position":
										instancePosition := (&Position{Name: instanceName}).Stage(stage)
										instance = any(instancePosition)
										__gong__map_Position[identifier] = instancePosition
									case "UmlState":
										instanceUmlState := (&UmlState{Name: instanceName}).Stage(stage)
										instance = any(instanceUmlState)
										__gong__map_UmlState[identifier] = instanceUmlState
									case "Umlsc":
										instanceUmlsc := (&Umlsc{Name: instanceName}).Stage(stage)
										instance = any(instanceUmlsc)
										__gong__map_Umlsc[identifier] = instanceUmlsc
									case "Vertice":
										instanceVertice := (&Vertice{Name: instanceName}).Stage(stage)
										instance = any(instanceVertice)
										__gong__map_Vertice[identifier] = instanceVertice
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
						case "Classdiagram":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DiagramPackage":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Field":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongEnumShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongEnumValueEntry":
							switch fieldName {
							// insertion point for date assign code
							}
						case "GongStructShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Link":
							switch fieldName {
							// insertion point for date assign code
							}
						case "NoteShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "NoteShapeLink":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Position":
							switch fieldName {
							// insertion point for date assign code
							}
						case "UmlState":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Umlsc":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Vertice":
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
					case "Classdiagram":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "GongStructShapes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongStructShape[targetIdentifier]
							__gong__map_Classdiagram[identifier].GongStructShapes =
								append(__gong__map_Classdiagram[identifier].GongStructShapes, target)
						case "GongEnumShapes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongEnumShape[targetIdentifier]
							__gong__map_Classdiagram[identifier].GongEnumShapes =
								append(__gong__map_Classdiagram[identifier].GongEnumShapes, target)
						case "NoteShapes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_NoteShape[targetIdentifier]
							__gong__map_Classdiagram[identifier].NoteShapes =
								append(__gong__map_Classdiagram[identifier].NoteShapes, target)
						}
					case "DiagramPackage":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Classdiagrams":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Classdiagram[targetIdentifier]
							__gong__map_DiagramPackage[identifier].Classdiagrams =
								append(__gong__map_DiagramPackage[identifier].Classdiagrams, target)
						case "Umlscs":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Umlsc[targetIdentifier]
							__gong__map_DiagramPackage[identifier].Umlscs =
								append(__gong__map_DiagramPackage[identifier].Umlscs, target)
						}
					case "Field":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GongEnumShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "GongEnumValueEntrys":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_GongEnumValueEntry[targetIdentifier]
							__gong__map_GongEnumShape[identifier].GongEnumValueEntrys =
								append(__gong__map_GongEnumShape[identifier].GongEnumValueEntrys, target)
						}
					case "GongEnumValueEntry":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "GongStructShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Fields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Field[targetIdentifier]
							__gong__map_GongStructShape[identifier].Fields =
								append(__gong__map_GongStructShape[identifier].Fields, target)
						case "Links":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Link[targetIdentifier]
							__gong__map_GongStructShape[identifier].Links =
								append(__gong__map_GongStructShape[identifier].Links, target)
						}
					case "Link":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "NoteShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "NoteShapeLinks":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_NoteShapeLink[targetIdentifier]
							__gong__map_NoteShape[identifier].NoteShapeLinks =
								append(__gong__map_NoteShape[identifier].NoteShapeLinks, target)
						}
					case "NoteShapeLink":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Position":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "UmlState":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Umlsc":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "States":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_UmlState[targetIdentifier]
							__gong__map_Umlsc[identifier].States =
								append(__gong__map_Umlsc[identifier].States, target)
						}
					case "Vertice":
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
		case *ast.BasicLit, *ast.UnaryExpr:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used

			if bl, ok := expr.(*ast.BasicLit); ok {
				// expression is  for instance ... = 18.000
				basicLit = bl
			} else if ue, ok := expr.(*ast.UnaryExpr); ok {
				// expression is  for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				basicLit = ue.X.(*ast.BasicLit)
				exprSign = -1
			}

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
			case "Classdiagram":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Classdiagram[identifier].Name = fielValue
				}
			case "DiagramPackage":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DiagramPackage[identifier].Name = fielValue
				case "Path":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DiagramPackage[identifier].Path = fielValue
				case "GongModelPath":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DiagramPackage[identifier].GongModelPath = fielValue
				case "AbsolutePathToDiagramPackage":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DiagramPackage[identifier].AbsolutePathToDiagramPackage = fielValue
				}
			case "Field":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Field[identifier].Name = fielValue
				case "Identifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Field[identifier].Identifier = fielValue
				case "FieldTypeAsString":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Field[identifier].FieldTypeAsString = fielValue
				case "Structname":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Field[identifier].Structname = fielValue
				case "Fieldtypename":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Field[identifier].Fieldtypename = fielValue
				}
			case "GongEnumShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnumShape[identifier].Name = fielValue
				case "Identifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnumShape[identifier].Identifier = fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongEnumShape[identifier].Width = exprSign * fielValue
				case "Heigth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongEnumShape[identifier].Heigth = exprSign * fielValue
				}
			case "GongEnumValueEntry":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnumValueEntry[identifier].Name = fielValue
				case "Identifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongEnumValueEntry[identifier].Identifier = fielValue
				}
			case "GongStructShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongStructShape[identifier].Name = fielValue
				case "Identifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_GongStructShape[identifier].Identifier = fielValue
				case "NbInstances":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongStructShape[identifier].NbInstances = int(exprSign) * int(fielValue)
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongStructShape[identifier].Width = exprSign * fielValue
				case "Heigth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongStructShape[identifier].Heigth = exprSign * fielValue
				}
			case "Link":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Name = fielValue
				case "Identifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Identifier = fielValue
				case "Fieldtypename":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Fieldtypename = fielValue
				case "FieldOffsetX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].FieldOffsetX = exprSign * fielValue
				case "FieldOffsetY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].FieldOffsetY = exprSign * fielValue
				case "TargetMultiplicityOffsetX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].TargetMultiplicityOffsetX = exprSign * fielValue
				case "TargetMultiplicityOffsetY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].TargetMultiplicityOffsetY = exprSign * fielValue
				case "SourceMultiplicityOffsetX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].SourceMultiplicityOffsetX = exprSign * fielValue
				case "SourceMultiplicityOffsetY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].SourceMultiplicityOffsetY = exprSign * fielValue
				case "StartRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].StartRatio = exprSign * fielValue
				case "EndRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].EndRatio = exprSign * fielValue
				case "CornerOffsetRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].CornerOffsetRatio = exprSign * fielValue
				}
			case "NoteShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteShape[identifier].Name = fielValue
				case "Identifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteShape[identifier].Identifier = fielValue
				case "Body":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteShape[identifier].Body = fielValue
				case "BodyHTML":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteShape[identifier].BodyHTML = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].Y = exprSign * fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].Width = exprSign * fielValue
				case "Heigth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].Heigth = exprSign * fielValue
				}
			case "NoteShapeLink":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteShapeLink[identifier].Name = fielValue
				case "Identifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteShapeLink[identifier].Identifier = fielValue
				}
			case "Position":
				switch fieldName {
				// insertion point for field dependant code
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Position[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Position[identifier].Y = exprSign * fielValue
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Position[identifier].Name = fielValue
				}
			case "UmlState":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_UmlState[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_UmlState[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_UmlState[identifier].Y = exprSign * fielValue
				}
			case "Umlsc":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Umlsc[identifier].Name = fielValue
				case "Activestate":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Umlsc[identifier].Activestate = fielValue
				}
			case "Vertice":
				switch fieldName {
				// insertion point for field dependant code
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Vertice[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Vertice[identifier].Y = exprSign * fielValue
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Vertice[identifier].Name = fielValue
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
			case "Classdiagram":
				switch fieldName {
				// insertion point for field dependant code
				case "IsInDrawMode":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Classdiagram[identifier].IsInDrawMode = fielValue
				}
			case "DiagramPackage":
				switch fieldName {
				// insertion point for field dependant code
				case "SelectedClassdiagram":
					targetIdentifier := ident.Name
					__gong__map_DiagramPackage[identifier].SelectedClassdiagram = __gong__map_Classdiagram[targetIdentifier]
				case "IsEditable":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DiagramPackage[identifier].IsEditable = fielValue
				case "IsReloaded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DiagramPackage[identifier].IsReloaded = fielValue
				}
			case "Field":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongEnumShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Position":
					targetIdentifier := ident.Name
					__gong__map_GongEnumShape[identifier].Position = __gong__map_Position[targetIdentifier]
				}
			case "GongEnumValueEntry":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "GongStructShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Position":
					targetIdentifier := ident.Name
					__gong__map_GongStructShape[identifier].Position = __gong__map_Position[targetIdentifier]
				case "ShowNbInstances":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongStructShape[identifier].ShowNbInstances = fielValue
				case "IsSelected":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_GongStructShape[identifier].IsSelected = fielValue
				}
			case "Link":
				switch fieldName {
				// insertion point for field dependant code
				case "Middlevertice":
					targetIdentifier := ident.Name
					__gong__map_Link[identifier].Middlevertice = __gong__map_Vertice[targetIdentifier]
				}
			case "NoteShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Matched":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].Matched = fielValue
				}
			case "NoteShapeLink":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Position":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "UmlState":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Umlsc":
				switch fieldName {
				// insertion point for field dependant code
				case "IsInDrawMode":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Umlsc[identifier].IsInDrawMode = fielValue
				}
			case "Vertice":
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
				case "Classdiagram":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DiagramPackage":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Field":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongEnumShape":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongEnumValueEntry":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "GongStructShape":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Link":
					switch fieldName {
					// insertion point for enum assign code
					case "TargetMultiplicity":
						var val MultiplicityType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Link[identifier].TargetMultiplicity = MultiplicityType(val)
					case "SourceMultiplicity":
						var val MultiplicityType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Link[identifier].SourceMultiplicity = MultiplicityType(val)
					case "StartOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Link[identifier].StartOrientation = OrientationType(val)
					case "EndOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Link[identifier].EndOrientation = OrientationType(val)
					}
				case "NoteShape":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "NoteShapeLink":
					switch fieldName {
					// insertion point for enum assign code
					case "Type":
						var val NoteShapeLinkType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_NoteShapeLink[identifier].Type = NoteShapeLinkType(val)
					}
				case "Position":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "UmlState":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Umlsc":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Vertice":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}
