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
var __gong__map_Animate = make(map[string]*Animate)
var __gong__map_Circle = make(map[string]*Circle)
var __gong__map_Ellipse = make(map[string]*Ellipse)
var __gong__map_Layer = make(map[string]*Layer)
var __gong__map_Line = make(map[string]*Line)
var __gong__map_Link = make(map[string]*Link)
var __gong__map_LinkAnchoredText = make(map[string]*LinkAnchoredText)
var __gong__map_Path = make(map[string]*Path)
var __gong__map_Point = make(map[string]*Point)
var __gong__map_Polygone = make(map[string]*Polygone)
var __gong__map_Polyline = make(map[string]*Polyline)
var __gong__map_Rect = make(map[string]*Rect)
var __gong__map_RectAnchoredRect = make(map[string]*RectAnchoredRect)
var __gong__map_RectAnchoredText = make(map[string]*RectAnchoredText)
var __gong__map_RectLinkLink = make(map[string]*RectLinkLink)
var __gong__map_SVG = make(map[string]*SVG)
var __gong__map_Text = make(map[string]*Text)

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
									case "Animate":
										instanceAnimate := (&Animate{Name: instanceName}).Stage(stage)
										instance = any(instanceAnimate)
										__gong__map_Animate[identifier] = instanceAnimate
									case "Circle":
										instanceCircle := (&Circle{Name: instanceName}).Stage(stage)
										instance = any(instanceCircle)
										__gong__map_Circle[identifier] = instanceCircle
									case "Ellipse":
										instanceEllipse := (&Ellipse{Name: instanceName}).Stage(stage)
										instance = any(instanceEllipse)
										__gong__map_Ellipse[identifier] = instanceEllipse
									case "Layer":
										instanceLayer := (&Layer{Name: instanceName}).Stage(stage)
										instance = any(instanceLayer)
										__gong__map_Layer[identifier] = instanceLayer
									case "Line":
										instanceLine := (&Line{Name: instanceName}).Stage(stage)
										instance = any(instanceLine)
										__gong__map_Line[identifier] = instanceLine
									case "Link":
										instanceLink := (&Link{Name: instanceName}).Stage(stage)
										instance = any(instanceLink)
										__gong__map_Link[identifier] = instanceLink
									case "LinkAnchoredText":
										instanceLinkAnchoredText := (&LinkAnchoredText{Name: instanceName}).Stage(stage)
										instance = any(instanceLinkAnchoredText)
										__gong__map_LinkAnchoredText[identifier] = instanceLinkAnchoredText
									case "Path":
										instancePath := (&Path{Name: instanceName}).Stage(stage)
										instance = any(instancePath)
										__gong__map_Path[identifier] = instancePath
									case "Point":
										instancePoint := (&Point{Name: instanceName}).Stage(stage)
										instance = any(instancePoint)
										__gong__map_Point[identifier] = instancePoint
									case "Polygone":
										instancePolygone := (&Polygone{Name: instanceName}).Stage(stage)
										instance = any(instancePolygone)
										__gong__map_Polygone[identifier] = instancePolygone
									case "Polyline":
										instancePolyline := (&Polyline{Name: instanceName}).Stage(stage)
										instance = any(instancePolyline)
										__gong__map_Polyline[identifier] = instancePolyline
									case "Rect":
										instanceRect := (&Rect{Name: instanceName}).Stage(stage)
										instance = any(instanceRect)
										__gong__map_Rect[identifier] = instanceRect
									case "RectAnchoredRect":
										instanceRectAnchoredRect := (&RectAnchoredRect{Name: instanceName}).Stage(stage)
										instance = any(instanceRectAnchoredRect)
										__gong__map_RectAnchoredRect[identifier] = instanceRectAnchoredRect
									case "RectAnchoredText":
										instanceRectAnchoredText := (&RectAnchoredText{Name: instanceName}).Stage(stage)
										instance = any(instanceRectAnchoredText)
										__gong__map_RectAnchoredText[identifier] = instanceRectAnchoredText
									case "RectLinkLink":
										instanceRectLinkLink := (&RectLinkLink{Name: instanceName}).Stage(stage)
										instance = any(instanceRectLinkLink)
										__gong__map_RectLinkLink[identifier] = instanceRectLinkLink
									case "SVG":
										instanceSVG := (&SVG{Name: instanceName}).Stage(stage)
										instance = any(instanceSVG)
										__gong__map_SVG[identifier] = instanceSVG
									case "Text":
										instanceText := (&Text{Name: instanceName}).Stage(stage)
										instance = any(instanceText)
										__gong__map_Text[identifier] = instanceText
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
						case "Animate":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Circle":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Ellipse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Layer":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Line":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Link":
							switch fieldName {
							// insertion point for date assign code
							}
						case "LinkAnchoredText":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Path":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Point":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Polygone":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Polyline":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Rect":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RectAnchoredRect":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RectAnchoredText":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RectLinkLink":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SVG":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Text":
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
					case "Animate":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Circle":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animations":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Circle[identifier].Animations =
								append(__gong__map_Circle[identifier].Animations, target)
						}
					case "Ellipse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Ellipse[identifier].Animates =
								append(__gong__map_Ellipse[identifier].Animates, target)
						}
					case "Layer":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Rects":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Rect[targetIdentifier]
							__gong__map_Layer[identifier].Rects =
								append(__gong__map_Layer[identifier].Rects, target)
						case "Texts":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Text[targetIdentifier]
							__gong__map_Layer[identifier].Texts =
								append(__gong__map_Layer[identifier].Texts, target)
						case "Circles":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Circle[targetIdentifier]
							__gong__map_Layer[identifier].Circles =
								append(__gong__map_Layer[identifier].Circles, target)
						case "Lines":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Line[targetIdentifier]
							__gong__map_Layer[identifier].Lines =
								append(__gong__map_Layer[identifier].Lines, target)
						case "Ellipses":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Ellipse[targetIdentifier]
							__gong__map_Layer[identifier].Ellipses =
								append(__gong__map_Layer[identifier].Ellipses, target)
						case "Polylines":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Polyline[targetIdentifier]
							__gong__map_Layer[identifier].Polylines =
								append(__gong__map_Layer[identifier].Polylines, target)
						case "Polygones":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Polygone[targetIdentifier]
							__gong__map_Layer[identifier].Polygones =
								append(__gong__map_Layer[identifier].Polygones, target)
						case "Paths":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Path[targetIdentifier]
							__gong__map_Layer[identifier].Paths =
								append(__gong__map_Layer[identifier].Paths, target)
						case "Links":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Link[targetIdentifier]
							__gong__map_Layer[identifier].Links =
								append(__gong__map_Layer[identifier].Links, target)
						case "RectLinkLinks":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_RectLinkLink[targetIdentifier]
							__gong__map_Layer[identifier].RectLinkLinks =
								append(__gong__map_Layer[identifier].RectLinkLinks, target)
						}
					case "Line":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Line[identifier].Animates =
								append(__gong__map_Line[identifier].Animates, target)
						}
					case "Link":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "TextAtArrowEnd":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_LinkAnchoredText[targetIdentifier]
							__gong__map_Link[identifier].TextAtArrowEnd =
								append(__gong__map_Link[identifier].TextAtArrowEnd, target)
						case "TextAtArrowStart":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_LinkAnchoredText[targetIdentifier]
							__gong__map_Link[identifier].TextAtArrowStart =
								append(__gong__map_Link[identifier].TextAtArrowStart, target)
						case "ControlPoints":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Point[targetIdentifier]
							__gong__map_Link[identifier].ControlPoints =
								append(__gong__map_Link[identifier].ControlPoints, target)
						}
					case "LinkAnchoredText":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_LinkAnchoredText[identifier].Animates =
								append(__gong__map_LinkAnchoredText[identifier].Animates, target)
						}
					case "Path":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Path[identifier].Animates =
								append(__gong__map_Path[identifier].Animates, target)
						}
					case "Point":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Polygone":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Polygone[identifier].Animates =
								append(__gong__map_Polygone[identifier].Animates, target)
						}
					case "Polyline":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Polyline[identifier].Animates =
								append(__gong__map_Polyline[identifier].Animates, target)
						}
					case "Rect":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animations":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Rect[identifier].Animations =
								append(__gong__map_Rect[identifier].Animations, target)
						case "RectAnchoredTexts":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_RectAnchoredText[targetIdentifier]
							__gong__map_Rect[identifier].RectAnchoredTexts =
								append(__gong__map_Rect[identifier].RectAnchoredTexts, target)
						case "RectAnchoredRects":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_RectAnchoredRect[targetIdentifier]
							__gong__map_Rect[identifier].RectAnchoredRects =
								append(__gong__map_Rect[identifier].RectAnchoredRects, target)
						}
					case "RectAnchoredRect":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RectAnchoredText":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_RectAnchoredText[identifier].Animates =
								append(__gong__map_RectAnchoredText[identifier].Animates, target)
						}
					case "RectLinkLink":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SVG":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Layers":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Layer[targetIdentifier]
							__gong__map_SVG[identifier].Layers =
								append(__gong__map_SVG[identifier].Layers, target)
						}
					case "Text":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Text[identifier].Animates =
								append(__gong__map_Text[identifier].Animates, target)
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
			case "Animate":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].Name = fielValue
				case "AttributeName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].AttributeName = fielValue
				case "Values":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].Values = fielValue
				case "Dur":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].Dur = fielValue
				case "RepeatCount":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].RepeatCount = fielValue
				}
			case "Circle":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Name = fielValue
				case "CX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].CX = exprSign * fielValue
				case "CY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].CY = exprSign * fielValue
				case "Radius":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].Radius = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Transform = fielValue
				}
			case "Ellipse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].Name = fielValue
				case "CX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].CX = exprSign * fielValue
				case "CY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].CY = exprSign * fielValue
				case "RX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].RX = exprSign * fielValue
				case "RY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].RY = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].Transform = fielValue
				}
			case "Layer":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Layer[identifier].Name = fielValue
				}
			case "Line":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].Name = fielValue
				case "X1":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].X1 = exprSign * fielValue
				case "Y1":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].Y1 = exprSign * fielValue
				case "X2":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].X2 = exprSign * fielValue
				case "Y2":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].Y2 = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].Transform = fielValue
				case "MouseClickX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].MouseClickX = exprSign * fielValue
				case "MouseClickY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].MouseClickY = exprSign * fielValue
				}
			case "Link":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Name = fielValue
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
				case "CornerRadius":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].CornerRadius = exprSign * fielValue
				case "EndArrowSize":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].EndArrowSize = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Link[identifier].Transform = fielValue
				}
			case "LinkAnchoredText":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].Content = fielValue
				case "X_Offset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_LinkAnchoredText[identifier].X_Offset = exprSign * fielValue
				case "Y_Offset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_LinkAnchoredText[identifier].Y_Offset = exprSign * fielValue
				case "FontWeight":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].FontWeight = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_LinkAnchoredText[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_LinkAnchoredText[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].Transform = fielValue
				}
			case "Path":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Name = fielValue
				case "Definition":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Definition = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Path[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Path[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Transform = fielValue
				}
			case "Point":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Point[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Point[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Point[identifier].Y = exprSign * fielValue
				}
			case "Polygone":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Name = fielValue
				case "Points":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Points = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polygone[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polygone[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Transform = fielValue
				}
			case "Polyline":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Name = fielValue
				case "Points":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Points = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polyline[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polyline[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Transform = fielValue
				}
			case "Rect":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].Y = exprSign * fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].Width = exprSign * fielValue
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].Height = exprSign * fielValue
				case "RX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].RX = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].Transform = fielValue
				}
			case "RectAnchoredRect":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredRect[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].Y = exprSign * fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].Width = exprSign * fielValue
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].Height = exprSign * fielValue
				case "RX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].RX = exprSign * fielValue
				case "X_Offset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].X_Offset = exprSign * fielValue
				case "Y_Offset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].Y_Offset = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredRect[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredRect[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredRect[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredRect[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredRect[identifier].Transform = fielValue
				}
			case "RectAnchoredText":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].Name = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].Content = fielValue
				case "FontWeight":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].FontWeight = fielValue
				case "FontSize":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredText[identifier].FontSize = int(exprSign) * int(fielValue)
				case "X_Offset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredText[identifier].X_Offset = exprSign * fielValue
				case "Y_Offset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredText[identifier].Y_Offset = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredText[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredText[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].Transform = fielValue
				}
			case "RectLinkLink":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectLinkLink[identifier].Name = fielValue
				case "TargetAnchorPosition":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectLinkLink[identifier].TargetAnchorPosition = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectLinkLink[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectLinkLink[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectLinkLink[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectLinkLink[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectLinkLink[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectLinkLink[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectLinkLink[identifier].Transform = fielValue
				}
			case "SVG":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SVG[identifier].Name = fielValue
				}
			case "Text":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].Y = exprSign * fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Content = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Transform = fielValue
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
			case "Animate":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Circle":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Ellipse":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Layer":
				switch fieldName {
				// insertion point for field dependant code
				case "Display":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Layer[identifier].Display = fielValue
				}
			case "Line":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Link":
				switch fieldName {
				// insertion point for field dependant code
				case "Start":
					targetIdentifier := ident.Name
					__gong__map_Link[identifier].Start = __gong__map_Rect[targetIdentifier]
				case "End":
					targetIdentifier := ident.Name
					__gong__map_Link[identifier].End = __gong__map_Rect[targetIdentifier]
				case "HasEndArrow":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].HasEndArrow = fielValue
				}
			case "LinkAnchoredText":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Path":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Point":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Polygone":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Polyline":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Rect":
				switch fieldName {
				// insertion point for field dependant code
				case "IsSelectable":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].IsSelectable = fielValue
				case "IsSelected":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].IsSelected = fielValue
				case "CanHaveLeftHandle":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].CanHaveLeftHandle = fielValue
				case "HasLeftHandle":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].HasLeftHandle = fielValue
				case "CanHaveRightHandle":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].CanHaveRightHandle = fielValue
				case "HasRightHandle":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].HasRightHandle = fielValue
				case "CanHaveTopHandle":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].CanHaveTopHandle = fielValue
				case "HasTopHandle":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].HasTopHandle = fielValue
				case "CanHaveBottomHandle":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].CanHaveBottomHandle = fielValue
				case "HasBottomHandle":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].HasBottomHandle = fielValue
				case "CanMoveHorizontaly":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].CanMoveHorizontaly = fielValue
				case "CanMoveVerticaly":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].CanMoveVerticaly = fielValue
				}
			case "RectAnchoredRect":
				switch fieldName {
				// insertion point for field dependant code
				case "WidthFollowRect":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].WidthFollowRect = fielValue
				case "HeightFollowRect":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].HeightFollowRect = fielValue
				}
			case "RectAnchoredText":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "RectLinkLink":
				switch fieldName {
				// insertion point for field dependant code
				case "Start":
					targetIdentifier := ident.Name
					__gong__map_RectLinkLink[identifier].Start = __gong__map_Rect[targetIdentifier]
				case "End":
					targetIdentifier := ident.Name
					__gong__map_RectLinkLink[identifier].End = __gong__map_Link[targetIdentifier]
				}
			case "SVG":
				switch fieldName {
				// insertion point for field dependant code
				case "StartRect":
					targetIdentifier := ident.Name
					__gong__map_SVG[identifier].StartRect = __gong__map_Rect[targetIdentifier]
				case "EndRect":
					targetIdentifier := ident.Name
					__gong__map_SVG[identifier].EndRect = __gong__map_Rect[targetIdentifier]
				case "IsEditable":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SVG[identifier].IsEditable = fielValue
				}
			case "Text":
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
				case "Animate":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Circle":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Ellipse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Layer":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Line":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Link":
					switch fieldName {
					// insertion point for enum assign code
					case "Type":
						var val LinkType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Link[identifier].Type = LinkType(val)
					case "StartAnchorType":
						var val AnchorType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Link[identifier].StartAnchorType = AnchorType(val)
					case "EndAnchorType":
						var val AnchorType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Link[identifier].EndAnchorType = AnchorType(val)
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
				case "LinkAnchoredText":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Path":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Point":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Polygone":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Polyline":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Rect":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "RectAnchoredRect":
					switch fieldName {
					// insertion point for enum assign code
					case "RectAnchorType":
						var val RectAnchorType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_RectAnchoredRect[identifier].RectAnchorType = RectAnchorType(val)
					}
				case "RectAnchoredText":
					switch fieldName {
					// insertion point for enum assign code
					case "RectAnchorType":
						var val RectAnchorType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_RectAnchoredText[identifier].RectAnchorType = RectAnchorType(val)
					case "TextAnchorType":
						var val TextAnchorType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_RectAnchoredText[identifier].TextAnchorType = TextAnchorType(val)
					}
				case "RectLinkLink":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SVG":
					switch fieldName {
					// insertion point for enum assign code
					case "DrawingState":
						var val DrawingState
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_SVG[identifier].DrawingState = DrawingState(val)
					}
				case "Text":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}
