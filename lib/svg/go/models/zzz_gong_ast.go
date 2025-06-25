// generated code - do not edit
package models

import (
	"bufio"
	"embed"
	"errors"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError
var _ = dummy_strconv_import
var dummy_time_import time.Time
var _ = dummy_time_import

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
func ParseAstFile(stage *Stage, pathToFile string) error {

	ReplaceOldDeclarationsInFile(pathToFile)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	// Read the file content using os.ReadFile
	content, err := os.ReadFile(fileOfInterest)
	if err != nil {
		return errors.New("Unable to read file " + err.Error())
	}

	// Assign the content to stage.contentWhenParsed
	stage.contentWhenParsed = string(content)

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
// specified by pathToFile within the provided embed.FS directory and
// stages instances declared in the file using the provided Stage.
//
// Parameters:
//
//	stage:      The staging area to populate.
//	directory:  The embedded filesystem containing the file.
//	pathToFile: The path to the Go source file within the embedded filesystem.
//
// Returns:
//
//	An error if reading or parsing the file fails, or if ParseAstFileFromAst fails.
func ParseAstEmbeddedFile(stage *Stage, directory embed.FS, pathToFile string) error {

	// 1. Read the content from the embedded filesystem.
	//    We don't need filepath.Abs as embed.FS uses relative paths.
	//    We also skip ReplaceOldDeclarationsInFile as embedded files are read-only.
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		// Return a specific error if the file can't be read from the embed.FS
		return errors.New(stage.GetName() + "; Unable to read embedded file " + err.Error())
	}

	// 2. Create a FileSet to manage position information.
	fset := token.NewFileSet()

	// 3. Parse the file content.
	//    Instead of passing a filename for the OS to read, we pass the pathToFile
	//    as the identifier and the actual file content (fileContentBytes) as the source.
	//    parser.ParseComments is included to match the original function's behavior.
	//    The type *ast.File is returned by parser.ParseFile.
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		// Return a specific error if parsing fails
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	// 4. Call the common AST processing logic.
	//    Pass the parsed AST (*ast.File), the FileSet, and the stage.
	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *Stage, inFile *ast.File, fset *token.FileSet) error {
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
				isOfInterest := strings.Contains(funcDecl.Name.Name, "_")
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
					case *ast.DeclStmt:
						if genDecl, ok := stmt.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.CONST {
							for _, spec := range genDecl.Specs {
								if valueSpec, ok := spec.(*ast.ValueSpec); ok {
									for i, name := range valueSpec.Names {
										if i < len(valueSpec.Values) {
											if basicLit, ok := valueSpec.Values[i].(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
												// Remove quotes from string literal
												value := strings.Trim(basicLit.Value, `"`)

												switch name.Name {
												case "__commitId__":
													if parsedUint, err := strconv.ParseUint(value, 10, 64); err == nil {
														stage.commitId = uint(parsedUint)
														stage.commitIdWhenParsed = stage.commitId
													}
												}
											}
										}
									}
								}
							}
						}
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
						// cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						var cmap ast.CommentMap
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
					if !strings.HasPrefix(ident.Name, "_") {
						continue
					}
					// declaration of a variable without initial value
					if len(spec.Values) == 0 {
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
var __gong__map_RectAnchoredPath = make(map[string]*RectAnchoredPath)
var __gong__map_RectAnchoredRect = make(map[string]*RectAnchoredRect)
var __gong__map_RectAnchoredText = make(map[string]*RectAnchoredText)
var __gong__map_RectLinkLink = make(map[string]*RectLinkLink)
var __gong__map_SVG = make(map[string]*SVG)
var __gong__map_SvgText = make(map[string]*SvgText)
var __gong__map_Text = make(map[string]*Text)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}

func lookupSym(recv, name string) bool {
	return recv == ""
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *Stage, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
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
			var basicLit *ast.BasicLit

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
										instanceAnimate := new(Animate)
										instanceAnimate.Name = instanceName
										instanceAnimate.Stage(stage)
										instance = any(instanceAnimate)
										__gong__map_Animate[identifier] = instanceAnimate
									case "Circle":
										instanceCircle := new(Circle)
										instanceCircle.Name = instanceName
										instanceCircle.Stage(stage)
										instance = any(instanceCircle)
										__gong__map_Circle[identifier] = instanceCircle
									case "Ellipse":
										instanceEllipse := new(Ellipse)
										instanceEllipse.Name = instanceName
										instanceEllipse.Stage(stage)
										instance = any(instanceEllipse)
										__gong__map_Ellipse[identifier] = instanceEllipse
									case "Layer":
										instanceLayer := new(Layer)
										instanceLayer.Name = instanceName
										instanceLayer.Stage(stage)
										instance = any(instanceLayer)
										__gong__map_Layer[identifier] = instanceLayer
									case "Line":
										instanceLine := new(Line)
										instanceLine.Name = instanceName
										instanceLine.Stage(stage)
										instance = any(instanceLine)
										__gong__map_Line[identifier] = instanceLine
									case "Link":
										instanceLink := new(Link)
										instanceLink.Name = instanceName
										instanceLink.Stage(stage)
										instance = any(instanceLink)
										__gong__map_Link[identifier] = instanceLink
									case "LinkAnchoredText":
										instanceLinkAnchoredText := new(LinkAnchoredText)
										instanceLinkAnchoredText.Name = instanceName
										instanceLinkAnchoredText.Stage(stage)
										instance = any(instanceLinkAnchoredText)
										__gong__map_LinkAnchoredText[identifier] = instanceLinkAnchoredText
									case "Path":
										instancePath := new(Path)
										instancePath.Name = instanceName
										instancePath.Stage(stage)
										instance = any(instancePath)
										__gong__map_Path[identifier] = instancePath
									case "Point":
										instancePoint := new(Point)
										instancePoint.Name = instanceName
										instancePoint.Stage(stage)
										instance = any(instancePoint)
										__gong__map_Point[identifier] = instancePoint
									case "Polygone":
										instancePolygone := new(Polygone)
										instancePolygone.Name = instanceName
										instancePolygone.Stage(stage)
										instance = any(instancePolygone)
										__gong__map_Polygone[identifier] = instancePolygone
									case "Polyline":
										instancePolyline := new(Polyline)
										instancePolyline.Name = instanceName
										instancePolyline.Stage(stage)
										instance = any(instancePolyline)
										__gong__map_Polyline[identifier] = instancePolyline
									case "Rect":
										instanceRect := new(Rect)
										instanceRect.Name = instanceName
										instanceRect.Stage(stage)
										instance = any(instanceRect)
										__gong__map_Rect[identifier] = instanceRect
									case "RectAnchoredPath":
										instanceRectAnchoredPath := new(RectAnchoredPath)
										instanceRectAnchoredPath.Name = instanceName
										instanceRectAnchoredPath.Stage(stage)
										instance = any(instanceRectAnchoredPath)
										__gong__map_RectAnchoredPath[identifier] = instanceRectAnchoredPath
									case "RectAnchoredRect":
										instanceRectAnchoredRect := new(RectAnchoredRect)
										instanceRectAnchoredRect.Name = instanceName
										instanceRectAnchoredRect.Stage(stage)
										instance = any(instanceRectAnchoredRect)
										__gong__map_RectAnchoredRect[identifier] = instanceRectAnchoredRect
									case "RectAnchoredText":
										instanceRectAnchoredText := new(RectAnchoredText)
										instanceRectAnchoredText.Name = instanceName
										instanceRectAnchoredText.Stage(stage)
										instance = any(instanceRectAnchoredText)
										__gong__map_RectAnchoredText[identifier] = instanceRectAnchoredText
									case "RectLinkLink":
										instanceRectLinkLink := new(RectLinkLink)
										instanceRectLinkLink.Name = instanceName
										instanceRectLinkLink.Stage(stage)
										instance = any(instanceRectLinkLink)
										__gong__map_RectLinkLink[identifier] = instanceRectLinkLink
									case "SVG":
										instanceSVG := new(SVG)
										instanceSVG.Name = instanceName
										instanceSVG.Stage(stage)
										instance = any(instanceSVG)
										__gong__map_SVG[identifier] = instanceSVG
									case "SvgText":
										instanceSvgText := new(SvgText)
										instanceSvgText.Name = instanceName
										instanceSvgText.Stage(stage)
										instance = any(instanceSvgText)
										__gong__map_SvgText[identifier] = instanceSvgText
									case "Text":
										instanceText := new(Text)
										instanceText.Name = instanceName
										instanceText.Stage(stage)
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
							log.Println("gongstructName not found for identifier", identifier)
							break
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
						case "RectAnchoredPath":
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
						case "SvgText":
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
				// pick up the first arg
				if len(callExpr.Args) != 1 {
					break
				}
				arg0 := callExpr.Args[0]

				var se *ast.SelectorExpr
				var ok bool
				if se, ok = arg0.(*ast.SelectorExpr); !ok {
					break
				}

				var seXident *ast.Ident
				if seXident = se.X.(*ast.Ident); !ok {
					break
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = "new(" + seXident.Name + "." + se.Sel.Name + ")"
				// following lines are here to avoid warning "unused write to field..."
				_ = basicLit.Kind
				_ = basicLit.Value
				_ = basicLit
			}
			for argNb, arg := range callExpr.Args {
				_ = argNb
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident, *ast.SelectorExpr:
					var ident *ast.Ident
					var ok bool
					_ = ok
					if ident, ok = arg.(*ast.Ident); !ok {
						// log.Println("we are in the case of new(....)")
					}

					var se *ast.SelectorExpr
					if se, ok = arg.(*ast.SelectorExpr); ok {
						if ident, ok = se.X.(*ast.Ident); !ok {
							// log.Println("we are in the case of append(....)")
						}
					}
					_ = ident

					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Println("gongstructName not found for identifier", identifier)
						break
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
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Circle[identifier]
								instanceWhoseFieldIsAppended.Animations = append(instanceWhoseFieldIsAppended.Animations, instanceToAppend)
							}
						}
					case "Ellipse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Ellipse[identifier]
								instanceWhoseFieldIsAppended.Animates = append(instanceWhoseFieldIsAppended.Animates, instanceToAppend)
							}
						}
					case "Layer":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Rects":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Rect[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.Rects = append(instanceWhoseFieldIsAppended.Rects, instanceToAppend)
							}
						case "Texts":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Text[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.Texts = append(instanceWhoseFieldIsAppended.Texts, instanceToAppend)
							}
						case "Circles":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Circle[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.Circles = append(instanceWhoseFieldIsAppended.Circles, instanceToAppend)
							}
						case "Lines":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Line[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.Lines = append(instanceWhoseFieldIsAppended.Lines, instanceToAppend)
							}
						case "Ellipses":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Ellipse[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.Ellipses = append(instanceWhoseFieldIsAppended.Ellipses, instanceToAppend)
							}
						case "Polylines":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Polyline[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.Polylines = append(instanceWhoseFieldIsAppended.Polylines, instanceToAppend)
							}
						case "Polygones":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Polygone[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.Polygones = append(instanceWhoseFieldIsAppended.Polygones, instanceToAppend)
							}
						case "Paths":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Path[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.Paths = append(instanceWhoseFieldIsAppended.Paths, instanceToAppend)
							}
						case "Links":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Link[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.Links = append(instanceWhoseFieldIsAppended.Links, instanceToAppend)
							}
						case "RectLinkLinks":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_RectLinkLink[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Layer[identifier]
								instanceWhoseFieldIsAppended.RectLinkLinks = append(instanceWhoseFieldIsAppended.RectLinkLinks, instanceToAppend)
							}
						}
					case "Line":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Line[identifier]
								instanceWhoseFieldIsAppended.Animates = append(instanceWhoseFieldIsAppended.Animates, instanceToAppend)
							}
						}
					case "Link":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "TextAtArrowStart":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_LinkAnchoredText[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Link[identifier]
								instanceWhoseFieldIsAppended.TextAtArrowStart = append(instanceWhoseFieldIsAppended.TextAtArrowStart, instanceToAppend)
							}
						case "TextAtArrowEnd":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_LinkAnchoredText[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Link[identifier]
								instanceWhoseFieldIsAppended.TextAtArrowEnd = append(instanceWhoseFieldIsAppended.TextAtArrowEnd, instanceToAppend)
							}
						case "ControlPoints":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Point[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Link[identifier]
								instanceWhoseFieldIsAppended.ControlPoints = append(instanceWhoseFieldIsAppended.ControlPoints, instanceToAppend)
							}
						}
					case "LinkAnchoredText":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_LinkAnchoredText[identifier]
								instanceWhoseFieldIsAppended.Animates = append(instanceWhoseFieldIsAppended.Animates, instanceToAppend)
							}
						}
					case "Path":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Path[identifier]
								instanceWhoseFieldIsAppended.Animates = append(instanceWhoseFieldIsAppended.Animates, instanceToAppend)
							}
						}
					case "Point":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Polygone":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Polygone[identifier]
								instanceWhoseFieldIsAppended.Animates = append(instanceWhoseFieldIsAppended.Animates, instanceToAppend)
							}
						}
					case "Polyline":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Polyline[identifier]
								instanceWhoseFieldIsAppended.Animates = append(instanceWhoseFieldIsAppended.Animates, instanceToAppend)
							}
						}
					case "Rect":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animations":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Rect[identifier]
								instanceWhoseFieldIsAppended.Animations = append(instanceWhoseFieldIsAppended.Animations, instanceToAppend)
							}
						case "RectAnchoredTexts":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_RectAnchoredText[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Rect[identifier]
								instanceWhoseFieldIsAppended.RectAnchoredTexts = append(instanceWhoseFieldIsAppended.RectAnchoredTexts, instanceToAppend)
							}
						case "RectAnchoredRects":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_RectAnchoredRect[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Rect[identifier]
								instanceWhoseFieldIsAppended.RectAnchoredRects = append(instanceWhoseFieldIsAppended.RectAnchoredRects, instanceToAppend)
							}
						case "RectAnchoredPaths":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_RectAnchoredPath[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Rect[identifier]
								instanceWhoseFieldIsAppended.RectAnchoredPaths = append(instanceWhoseFieldIsAppended.RectAnchoredPaths, instanceToAppend)
							}
						}
					case "RectAnchoredPath":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RectAnchoredRect":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RectAnchoredText":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_RectAnchoredText[identifier]
								instanceWhoseFieldIsAppended.Animates = append(instanceWhoseFieldIsAppended.Animates, instanceToAppend)
							}
						}
					case "RectLinkLink":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SVG":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Layers":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Layer[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_SVG[identifier]
								instanceWhoseFieldIsAppended.Layers = append(instanceWhoseFieldIsAppended.Layers, instanceToAppend)
							}
						}
					case "SvgText":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Text":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Animate[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Text[identifier]
								instanceWhoseFieldIsAppended.Animates = append(instanceWhoseFieldIsAppended.Animates, instanceToAppend)
							}
						}
					}
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr, *ast.CompositeLit:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used
			switch v := expr.(type) {
			case *ast.BasicLit:
				// expression is for instance ... = 18.000
				basicLit = v
			case *ast.UnaryExpr:
				// expression is for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				if bl, ok := v.X.(*ast.BasicLit); ok {
					basicLit = bl
					// Check the operator to set the sign
					if v.Op == token.SUB { // token.SUB is for '-'
						exprSign = -1
					} else if v.Op == token.ADD { // token.ADD is for '+'
						exprSign = 1
					}
				}
			case *ast.CompositeLit:
				var sl *ast.SelectorExpr
				var ident *ast.Ident
				var ok bool

				if sl, ok = v.Type.(*ast.SelectorExpr); !ok {
					break // Exits the switch case
				}

				if ident, ok = sl.X.(*ast.Ident); !ok {
					break // Exits the switch case
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = ident.Name + "." + sl.Sel.Name + "{}"
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Println("gongstructName not found for identifier", identifier)
				break
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
				case "From":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].From = fielValue
				case "To":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].To = fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].StrokeOpacity = exprSign * fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].StrokeOpacity = exprSign * fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].StrokeOpacity = exprSign * fielValue
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
				case "StartArrowSize":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].StartArrowSize = exprSign * fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].StrokeOpacity = exprSign * fielValue
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
				case "FontSize":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].FontSize = fielValue
				case "FontStyle":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].FontStyle = fielValue
				case "LetterSpacing":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LinkAnchoredText[identifier].LetterSpacing = fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_LinkAnchoredText[identifier].StrokeOpacity = exprSign * fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Path[identifier].StrokeOpacity = exprSign * fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polygone[identifier].StrokeOpacity = exprSign * fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polyline[identifier].StrokeOpacity = exprSign * fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].StrokeOpacity = exprSign * fielValue
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
				case "ColorWhenHovered":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].ColorWhenHovered = fielValue
				case "OriginalColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].OriginalColor = fielValue
				case "FillOpacityWhenHovered":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].FillOpacityWhenHovered = exprSign * fielValue
				case "OriginalFillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].OriginalFillOpacity = exprSign * fielValue
				case "ToolTipText":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].ToolTipText = fielValue
				}
			case "RectAnchoredPath":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredPath[identifier].Name = fielValue
				case "Definition":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredPath[identifier].Definition = fielValue
				case "X_Offset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredPath[identifier].X_Offset = exprSign * fielValue
				case "Y_Offset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredPath[identifier].Y_Offset = exprSign * fielValue
				case "AppliedScaling":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredPath[identifier].AppliedScaling = exprSign * fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredPath[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredPath[identifier].FillOpacity = exprSign * fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredPath[identifier].Stroke = fielValue
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredPath[identifier].StrokeOpacity = exprSign * fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredPath[identifier].StrokeWidth = exprSign * fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredPath[identifier].StrokeDashArray = fielValue
				case "StrokeDashArrayWhenSelected":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredPath[identifier].StrokeDashArrayWhenSelected = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredPath[identifier].Transform = fielValue
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
				case "ToolTipText":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredRect[identifier].ToolTipText = fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].StrokeOpacity = exprSign * fielValue
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
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].FontSize = fielValue
				case "FontStyle":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].FontStyle = fielValue
				case "LetterSpacing":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RectAnchoredText[identifier].LetterSpacing = fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredText[identifier].StrokeOpacity = exprSign * fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectLinkLink[identifier].StrokeOpacity = exprSign * fielValue
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
				case "DefaultDirectoryForGeneratedImages":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SVG[identifier].DefaultDirectoryForGeneratedImages = fielValue
				}
			case "SvgText":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SvgText[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SvgText[identifier].Text = fielValue
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
				case "StrokeOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].StrokeOpacity = exprSign * fielValue
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
				case "FontWeight":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].FontWeight = fielValue
				case "FontSize":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].FontSize = fielValue
				case "FontStyle":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].FontStyle = fielValue
				case "LetterSpacing":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].LetterSpacing = fielValue
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
				log.Println("gongstructName not found for identifier", identifier)
				break
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
				}
			case "Line":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Link":
				switch fieldName {
				// insertion point for field dependant code
				case "IsBezierCurve":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].IsBezierCurve = fielValue
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
				case "HasStartArrow":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Link[identifier].HasStartArrow = fielValue
				}
			case "LinkAnchoredText":
				switch fieldName {
				// insertion point for field dependant code
				case "AutomaticLayout":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_LinkAnchoredText[identifier].AutomaticLayout = fielValue
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
				case "IsScalingProportionally":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].IsScalingProportionally = fielValue
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
				case "ChangeColorWhenHovered":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].ChangeColorWhenHovered = fielValue
				case "HasToolTip":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].HasToolTip = fielValue
				}
			case "RectAnchoredPath":
				switch fieldName {
				// insertion point for field dependant code
				case "ScalePropotionnally":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredPath[identifier].ScalePropotionnally = fielValue
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
				case "HasToolTip":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_RectAnchoredRect[identifier].HasToolTip = fielValue
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
				case "IsSVGFrontEndFileGenerated":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SVG[identifier].IsSVGFrontEndFileGenerated = fielValue
				case "IsSVGBackEndFileGenerated":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SVG[identifier].IsSVGBackEndFileGenerated = fielValue
				case "IsControlBannerHidden":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SVG[identifier].IsControlBannerHidden = fielValue
				}
			case "SvgText":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Text":
				switch fieldName {
				// insertion point for field dependant code
				}
			}
		case *ast.SelectorExpr:
			var basicLit *ast.BasicLit
			var ident *ast.Ident

			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			case *ast.CompositeLit:
				var ok bool
				var sl *ast.SelectorExpr

				if sl, ok = X.Type.(*ast.SelectorExpr); !ok {
					break // Exits the switch case
				}

				if ident, ok = sl.X.(*ast.Ident); !ok {
					break // Exits the switch case
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = ident.Name + "." + sl.Sel.Name + "{}." + selectorExpr.Sel.Name
			}

			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Println("gongstructName not found for identifier", identifier)
					break
				}

				if basicLit == nil {
					// for the meta field written as ref_models.ENUM_VALUE1
					basicLit = new(ast.BasicLit)
					basicLit.Kind = token.STRING // Or another appropriate token.Kind
					basicLit.Value = selectorExpr.X.(*ast.Ident).Name + "." + Sel.Name
					_ = basicLit.Kind
					_ = basicLit.Value
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for selector expr assignments
				case "Animate":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Circle":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Ellipse":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Layer":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Line":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Link":
					switch fieldName {
					// insertion point for selector expr assign code
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
					// insertion point for selector expr assign code
					case "LinkAnchorType":
						var val LinkAnchorType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_LinkAnchoredText[identifier].LinkAnchorType = LinkAnchorType(val)
					}
				case "Path":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Point":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Polygone":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Polyline":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Rect":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "RectAnchoredPath":
					switch fieldName {
					// insertion point for selector expr assign code
					case "RectAnchorType":
						var val RectAnchorType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_RectAnchoredPath[identifier].RectAnchorType = RectAnchorType(val)
					}
				case "RectAnchoredRect":
					switch fieldName {
					// insertion point for selector expr assign code
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
					// insertion point for selector expr assign code
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
					case "WritingMode":
						var val WritingMode
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_RectAnchoredText[identifier].WritingMode = WritingMode(val)
					}
				case "RectLinkLink":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SVG":
					switch fieldName {
					// insertion point for selector expr assign code
					case "DrawingState":
						var val DrawingState
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_SVG[identifier].DrawingState = DrawingState(val)
					}
				case "SvgText":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Text":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				}
			}
		}
	}
	return
}

// ReplaceOldDeclarationsInFile replaces specific text in a file at the given path.
func ReplaceOldDeclarationsInFile(pathToFile string) error {
	// Open the file for reading.
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// replacing function with Injection
	pattern := regexp.MustCompile(`\b\w*Injection\b`)
	pattern2 := regexp.MustCompile(`\bmap_DocLink_Identifier_\w*\b`)

	// Temporary slice to hold lines from the file.
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Replace the target text with the desired text.
		line := strings.Replace(scanner.Text(), "var ___dummy__Time_stage time.Time", "var _ time.Time", -1)
		line = pattern.ReplaceAllString(line, "_")
		line = pattern2.ReplaceAllString(line, "_")

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Re-open the file for writing.
	file, err = os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the modified lines back to the file.
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
