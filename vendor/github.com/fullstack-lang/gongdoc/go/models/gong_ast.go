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
var __gong__map_Classdiagram = make(map[string]*Classdiagram)
var __gong__map_Classshape = make(map[string]*Classshape)
var __gong__map_DiagramPackage = make(map[string]*DiagramPackage)
var __gong__map_Field = make(map[string]*Field)
var __gong__map_Link = make(map[string]*Link)
var __gong__map_Node = make(map[string]*Node)
var __gong__map_NoteLink = make(map[string]*NoteLink)
var __gong__map_NoteShape = make(map[string]*NoteShape)
var __gong__map_Position = make(map[string]*Position)
var __gong__map_Reference = make(map[string]*Reference)
var __gong__map_Tree = make(map[string]*Tree)
var __gong__map_UmlState = make(map[string]*UmlState)
var __gong__map_Umlsc = make(map[string]*Umlsc)
var __gong__map_Vertice = make(map[string]*Vertice)

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
									case "Classdiagram":
										instanceClassdiagram := (&Classdiagram{Name: instanceName}).Stage()
										instance = any(instanceClassdiagram)
										__gong__map_Classdiagram[identifier] = instanceClassdiagram
									case "Classshape":
										instanceClassshape := (&Classshape{Name: instanceName}).Stage()
										instance = any(instanceClassshape)
										__gong__map_Classshape[identifier] = instanceClassshape
									case "DiagramPackage":
										instanceDiagramPackage := (&DiagramPackage{Name: instanceName}).Stage()
										instance = any(instanceDiagramPackage)
										__gong__map_DiagramPackage[identifier] = instanceDiagramPackage
									case "Field":
										instanceField := (&Field{Name: instanceName}).Stage()
										instance = any(instanceField)
										__gong__map_Field[identifier] = instanceField
									case "Link":
										instanceLink := (&Link{Name: instanceName}).Stage()
										instance = any(instanceLink)
										__gong__map_Link[identifier] = instanceLink
									case "Node":
										instanceNode := (&Node{Name: instanceName}).Stage()
										instance = any(instanceNode)
										__gong__map_Node[identifier] = instanceNode
									case "NoteLink":
										instanceNoteLink := (&NoteLink{Name: instanceName}).Stage()
										instance = any(instanceNoteLink)
										__gong__map_NoteLink[identifier] = instanceNoteLink
									case "NoteShape":
										instanceNoteShape := (&NoteShape{Name: instanceName}).Stage()
										instance = any(instanceNoteShape)
										__gong__map_NoteShape[identifier] = instanceNoteShape
									case "Position":
										instancePosition := (&Position{Name: instanceName}).Stage()
										instance = any(instancePosition)
										__gong__map_Position[identifier] = instancePosition
									case "Reference":
										instanceReference := (&Reference{Name: instanceName}).Stage()
										instance = any(instanceReference)
										__gong__map_Reference[identifier] = instanceReference
									case "Tree":
										instanceTree := (&Tree{Name: instanceName}).Stage()
										instance = any(instanceTree)
										__gong__map_Tree[identifier] = instanceTree
									case "UmlState":
										instanceUmlState := (&UmlState{Name: instanceName}).Stage()
										instance = any(instanceUmlState)
										__gong__map_UmlState[identifier] = instanceUmlState
									case "Umlsc":
										instanceUmlsc := (&Umlsc{Name: instanceName}).Stage()
										instance = any(instanceUmlsc)
										__gong__map_Umlsc[identifier] = instanceUmlsc
									case "Vertice":
										instanceVertice := (&Vertice{Name: instanceName}).Stage()
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
						case "Classshape":
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
						case "Link":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Node":
							switch fieldName {
							// insertion point for date assign code
							}
						case "NoteLink":
							switch fieldName {
							// insertion point for date assign code
							}
						case "NoteShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Position":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Reference":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tree":
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
						case "Classshapes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Classshape[targetIdentifier]
							__gong__map_Classdiagram[identifier].Classshapes =
								append(__gong__map_Classdiagram[identifier].Classshapes, target)
						case "Notes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_NoteShape[targetIdentifier]
							__gong__map_Classdiagram[identifier].Notes =
								append(__gong__map_Classdiagram[identifier].Notes, target)
						}
					case "Classshape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Fields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Field[targetIdentifier]
							__gong__map_Classshape[identifier].Fields =
								append(__gong__map_Classshape[identifier].Fields, target)
						case "Links":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Link[targetIdentifier]
							__gong__map_Classshape[identifier].Links =
								append(__gong__map_Classshape[identifier].Links, target)
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
					case "Link":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Node":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Children":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Node[targetIdentifier]
							__gong__map_Node[identifier].Children =
								append(__gong__map_Node[identifier].Children, target)
						}
					case "NoteLink":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "NoteShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "NoteLinks":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_NoteLink[targetIdentifier]
							__gong__map_NoteShape[identifier].NoteLinks =
								append(__gong__map_NoteShape[identifier].NoteLinks, target)
						}
					case "Position":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Reference":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tree":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "RootNodes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Node[targetIdentifier]
							__gong__map_Tree[identifier].RootNodes =
								append(__gong__map_Tree[identifier].RootNodes, target)
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
			case "Classdiagram":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Classdiagram[identifier].Name = fielValue
				}
			case "Classshape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Classshape[identifier].Name = fielValue
				case "ReferenceName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Classshape[identifier].ReferenceName = fielValue
				case "NbInstances":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Classshape[identifier].NbInstances = int(fielValue)
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Classshape[identifier].Width = fielValue
				case "Heigth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Classshape[identifier].Heigth = fielValue
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
				case "Fieldname":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Field[identifier].Fieldname = fielValue
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
			case "Link":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Name = fielValue
				case "Fieldname":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Fieldname = fielValue
				case "Structname":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Structname = fielValue
				case "Fieldtypename":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Fieldtypename = fielValue
				}
			case "Node":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Node[identifier].Name = fielValue
				}
			case "NoteLink":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteLink[identifier].Name = fielValue
				}
			case "NoteShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteShape[identifier].Name = fielValue
				case "Body":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteShape[identifier].Body = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].X = fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].Y = fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].Width = fielValue
				case "Heigth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].Heigth = fielValue
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
					__gong__map_Position[identifier].X = fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Position[identifier].Y = fielValue
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Position[identifier].Name = fielValue
				}
			case "Reference":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Reference[identifier].Name = fielValue
				case "NbInstances":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Reference[identifier].NbInstances = int(fielValue)
				}
			case "Tree":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tree[identifier].Name = fielValue
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
					__gong__map_UmlState[identifier].X = fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_UmlState[identifier].Y = fielValue
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
					__gong__map_Vertice[identifier].X = fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Vertice[identifier].Y = fielValue
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
			case "Classshape":
				switch fieldName {
				// insertion point for field dependant code
				case "Position":
					targetIdentifier := ident.Name
					__gong__map_Classshape[identifier].Position = __gong__map_Position[targetIdentifier]
				case "Reference":
					targetIdentifier := ident.Name
					__gong__map_Classshape[identifier].Reference = __gong__map_Reference[targetIdentifier]
				case "ShowNbInstances":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Classshape[identifier].ShowNbInstances = fielValue
				case "IsSelected":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Classshape[identifier].IsSelected = fielValue
				}
			case "DiagramPackage":
				switch fieldName {
				// insertion point for field dependant code
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
			case "Link":
				switch fieldName {
				// insertion point for field dependant code
				case "Middlevertice":
					targetIdentifier := ident.Name
					__gong__map_Link[identifier].Middlevertice = __gong__map_Vertice[targetIdentifier]
				}
			case "Node":
				switch fieldName {
				// insertion point for field dependant code
				case "Classdiagram":
					targetIdentifier := ident.Name
					__gong__map_Node[identifier].Classdiagram = __gong__map_Classdiagram[targetIdentifier]
				case "Umlsc":
					targetIdentifier := ident.Name
					__gong__map_Node[identifier].Umlsc = __gong__map_Umlsc[targetIdentifier]
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].IsExpanded = fielValue
				case "HasCheckboxButton":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].HasCheckboxButton = fielValue
				case "IsChecked":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].IsChecked = fielValue
				case "IsCheckboxDisabled":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].IsCheckboxDisabled = fielValue
				case "HasAddChildButton":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].HasAddChildButton = fielValue
				case "HasEditButton":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].HasEditButton = fielValue
				case "IsInEditMode":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].IsInEditMode = fielValue
				case "HasDrawButton":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].HasDrawButton = fielValue
				case "HasDrawOffButton":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].HasDrawOffButton = fielValue
				case "IsInDrawMode":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].IsInDrawMode = fielValue
				case "IsSaved":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].IsSaved = fielValue
				case "HasDeleteButton":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Node[identifier].HasDeleteButton = fielValue
				}
			case "NoteLink":
				switch fieldName {
				// insertion point for field dependant code
				case "Classshape":
					targetIdentifier := ident.Name
					__gong__map_NoteLink[identifier].Classshape = __gong__map_Classshape[targetIdentifier]
				case "Link":
					targetIdentifier := ident.Name
					__gong__map_NoteLink[identifier].Link = __gong__map_Link[targetIdentifier]
				case "Middlevertice":
					targetIdentifier := ident.Name
					__gong__map_NoteLink[identifier].Middlevertice = __gong__map_Vertice[targetIdentifier]
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
			case "Position":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Reference":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tree":
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
				case "Classshape":
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
					}
				case "Node":
					switch fieldName {
					// insertion point for enum assign code
					case "Type":
						var val GongdocNodeType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Node[identifier].Type = GongdocNodeType(val)
					}
				case "NoteLink":
					switch fieldName {
					// insertion point for enum assign code
					case "Type":
						var val ReferenceType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_NoteLink[identifier].Type = ReferenceType(val)
					}
				case "NoteShape":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Position":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Reference":
					switch fieldName {
					// insertion point for enum assign code
					case "Type":
						var val ReferenceType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Reference[identifier].Type = ReferenceType(val)
					}
				case "Tree":
					switch fieldName {
					// insertion point for enum assign code
					case "Type":
						var val TreeType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Tree[identifier].Type = TreeType(val)
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
