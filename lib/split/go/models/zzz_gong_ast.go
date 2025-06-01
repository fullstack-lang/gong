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
		return errors.New("Unable to read embedded file '" + pathToFile + "': " + err.Error())
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
var __gong__map_AsSplit = make(map[string]*AsSplit)
var __gong__map_AsSplitArea = make(map[string]*AsSplitArea)
var __gong__map_Button = make(map[string]*Button)
var __gong__map_Cursor = make(map[string]*Cursor)
var __gong__map_Doc = make(map[string]*Doc)
var __gong__map_Form = make(map[string]*Form)
var __gong__map_Load = make(map[string]*Load)
var __gong__map_Slider = make(map[string]*Slider)
var __gong__map_Split = make(map[string]*Split)
var __gong__map_Svg = make(map[string]*Svg)
var __gong__map_Table = make(map[string]*Table)
var __gong__map_Tone = make(map[string]*Tone)
var __gong__map_Tree = make(map[string]*Tree)
var __gong__map_View = make(map[string]*View)
var __gong__map_Xlsx = make(map[string]*Xlsx)

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
									case "AsSplit":
										instanceAsSplit := new(AsSplit)
										instanceAsSplit.Name = instanceName
										instanceAsSplit.Stage(stage)
										instance = any(instanceAsSplit)
										__gong__map_AsSplit[identifier] = instanceAsSplit
									case "AsSplitArea":
										instanceAsSplitArea := new(AsSplitArea)
										instanceAsSplitArea.Name = instanceName
										instanceAsSplitArea.Stage(stage)
										instance = any(instanceAsSplitArea)
										__gong__map_AsSplitArea[identifier] = instanceAsSplitArea
									case "Button":
										instanceButton := new(Button)
										instanceButton.Name = instanceName
										instanceButton.Stage(stage)
										instance = any(instanceButton)
										__gong__map_Button[identifier] = instanceButton
									case "Cursor":
										instanceCursor := new(Cursor)
										instanceCursor.Name = instanceName
										instanceCursor.Stage(stage)
										instance = any(instanceCursor)
										__gong__map_Cursor[identifier] = instanceCursor
									case "Doc":
										instanceDoc := new(Doc)
										instanceDoc.Name = instanceName
										instanceDoc.Stage(stage)
										instance = any(instanceDoc)
										__gong__map_Doc[identifier] = instanceDoc
									case "Form":
										instanceForm := new(Form)
										instanceForm.Name = instanceName
										instanceForm.Stage(stage)
										instance = any(instanceForm)
										__gong__map_Form[identifier] = instanceForm
									case "Load":
										instanceLoad := new(Load)
										instanceLoad.Name = instanceName
										instanceLoad.Stage(stage)
										instance = any(instanceLoad)
										__gong__map_Load[identifier] = instanceLoad
									case "Slider":
										instanceSlider := new(Slider)
										instanceSlider.Name = instanceName
										instanceSlider.Stage(stage)
										instance = any(instanceSlider)
										__gong__map_Slider[identifier] = instanceSlider
									case "Split":
										instanceSplit := new(Split)
										instanceSplit.Name = instanceName
										instanceSplit.Stage(stage)
										instance = any(instanceSplit)
										__gong__map_Split[identifier] = instanceSplit
									case "Svg":
										instanceSvg := new(Svg)
										instanceSvg.Name = instanceName
										instanceSvg.Stage(stage)
										instance = any(instanceSvg)
										__gong__map_Svg[identifier] = instanceSvg
									case "Table":
										instanceTable := new(Table)
										instanceTable.Name = instanceName
										instanceTable.Stage(stage)
										instance = any(instanceTable)
										__gong__map_Table[identifier] = instanceTable
									case "Tone":
										instanceTone := new(Tone)
										instanceTone.Name = instanceName
										instanceTone.Stage(stage)
										instance = any(instanceTone)
										__gong__map_Tone[identifier] = instanceTone
									case "Tree":
										instanceTree := new(Tree)
										instanceTree.Name = instanceName
										instanceTree.Stage(stage)
										instance = any(instanceTree)
										__gong__map_Tree[identifier] = instanceTree
									case "View":
										instanceView := new(View)
										instanceView.Name = instanceName
										instanceView.Stage(stage)
										instance = any(instanceView)
										__gong__map_View[identifier] = instanceView
									case "Xlsx":
										instanceXlsx := new(Xlsx)
										instanceXlsx.Name = instanceName
										instanceXlsx.Stage(stage)
										instance = any(instanceXlsx)
										__gong__map_Xlsx[identifier] = instanceXlsx
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
						case "AsSplit":
							switch fieldName {
							// insertion point for date assign code
							}
						case "AsSplitArea":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Button":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Cursor":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Doc":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Form":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Load":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Slider":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Split":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Svg":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Table":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tone":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Tree":
							switch fieldName {
							// insertion point for date assign code
							}
						case "View":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Xlsx":
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
					case "AsSplit":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "AsSplitAreas":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_AsSplitArea[targetIdentifier]
							__gong__map_AsSplit[identifier].AsSplitAreas =
								append(__gong__map_AsSplit[identifier].AsSplitAreas, target)
						}
					case "AsSplitArea":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Button":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Cursor":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Doc":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Form":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Load":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Slider":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Split":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Svg":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Table":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tone":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Tree":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "View":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "RootAsSplitAreas":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_AsSplitArea[targetIdentifier]
							__gong__map_View[identifier].RootAsSplitAreas =
								append(__gong__map_View[identifier].RootAsSplitAreas, target)
						}
					case "Xlsx":
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
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "AsSplit":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AsSplit[identifier].Name = fielValue
				}
			case "AsSplitArea":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AsSplitArea[identifier].Name = fielValue
				case "Size":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AsSplitArea[identifier].Size = exprSign * fielValue
				case "DivStyle":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AsSplitArea[identifier].DivStyle = fielValue
				}
			case "Button":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Button[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Button[identifier].StackName = fielValue
				}
			case "Cursor":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Cursor[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Cursor[identifier].StackName = fielValue
				case "Style":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Cursor[identifier].Style = fielValue
				}
			case "Doc":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Doc[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Doc[identifier].StackName = fielValue
				}
			case "Form":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Form[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Form[identifier].StackName = fielValue
				case "FormName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Form[identifier].FormName = fielValue
				}
			case "Load":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Load[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Load[identifier].StackName = fielValue
				}
			case "Slider":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Slider[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Slider[identifier].StackName = fielValue
				}
			case "Split":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Split[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Split[identifier].StackName = fielValue
				}
			case "Svg":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Svg[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Svg[identifier].StackName = fielValue
				case "Style":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Svg[identifier].Style = fielValue
				}
			case "Table":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Table[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Table[identifier].StackName = fielValue
				case "TableName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Table[identifier].TableName = fielValue
				}
			case "Tone":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tone[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tone[identifier].StackName = fielValue
				}
			case "Tree":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tree[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tree[identifier].StackName = fielValue
				case "TreeName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Tree[identifier].TreeName = fielValue
				}
			case "View":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_View[identifier].Name = fielValue
				}
			case "Xlsx":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Xlsx[identifier].Name = fielValue
				case "StackName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Xlsx[identifier].StackName = fielValue
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
			case "AsSplit":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "AsSplitArea":
				switch fieldName {
				// insertion point for field dependant code
				case "ShowNameInHeader":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AsSplitArea[identifier].ShowNameInHeader = fielValue
				case "IsAny":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AsSplitArea[identifier].IsAny = fielValue
				case "AsSplit":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].AsSplit = __gong__map_AsSplit[targetIdentifier]
				case "Button":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Button = __gong__map_Button[targetIdentifier]
				case "Cursor":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Cursor = __gong__map_Cursor[targetIdentifier]
				case "Doc":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Doc = __gong__map_Doc[targetIdentifier]
				case "Form":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Form = __gong__map_Form[targetIdentifier]
				case "Load":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Load = __gong__map_Load[targetIdentifier]
				case "Slider":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Slider = __gong__map_Slider[targetIdentifier]
				case "Split":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Split = __gong__map_Split[targetIdentifier]
				case "Svg":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Svg = __gong__map_Svg[targetIdentifier]
				case "Table":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Table = __gong__map_Table[targetIdentifier]
				case "Tone":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Tone = __gong__map_Tone[targetIdentifier]
				case "Tree":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Tree = __gong__map_Tree[targetIdentifier]
				case "Xlsx":
					targetIdentifier := ident.Name
					__gong__map_AsSplitArea[identifier].Xlsx = __gong__map_Xlsx[targetIdentifier]
				case "HasDiv":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AsSplitArea[identifier].HasDiv = fielValue
				}
			case "Button":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Cursor":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Doc":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Form":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Load":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Slider":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Split":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Svg":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Table":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tone":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Tree":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "View":
				switch fieldName {
				// insertion point for field dependant code
				case "ShowViewName":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_View[identifier].ShowViewName = fielValue
				}
			case "Xlsx":
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
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for selector expr assignments
				case "AsSplit":
					switch fieldName {
					// insertion point for selector expr assign code
					case "Direction":
						var val Direction
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_AsSplit[identifier].Direction = Direction(val)
					}
				case "AsSplitArea":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Button":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Cursor":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Doc":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Form":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Load":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Slider":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Split":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Svg":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Table":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Tone":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Tree":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "View":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Xlsx":
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
