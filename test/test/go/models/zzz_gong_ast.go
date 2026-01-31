// generated code - do not edit
package models

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError
var _ = dummy_strconv_import
var dummy_time_import time.Time
var _ = dummy_time_import
var dummy_slices_import = slices.Insert([]int{0}, 0)
var _ = dummy_slices_import

// ParseAstFileLegacy Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileLegacy(stage *Stage, pathToFile string, preserveOrder bool) error {

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

	return ParseAstFileLegacyFromAst(stage, inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFileLegacy parses the Go source code from an embedded file
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
//	An error if reading or parsing the file fails, or if ParseAstFileLegacyFromAst fails.
func ParseAstEmbeddedFileLegacy(stage *Stage, directory embed.FS, pathToFile string) error {

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
	return ParseAstFileLegacyFromAst(stage, inFile, fset, false)
}

// ParseAstFileLegacy Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileLegacyFromAst(stage *Stage, inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	// Robust parsing of imports to identify the meta package.
	// We ignore standard imports and the primary models package import.
	stage.MetaPackageImportAlias = ""
	stage.MetaPackageImportPath = ""
	for _, imp := range inFile.Imports {
		path := strings.Trim(imp.Path.Value, "\"")

		// Skip known standard packages and the main models package for this stage
		if path == "time" || path == "slices" || path == stage.GetType() {
			continue
		}

		// The remaining import is the meta package used for docLink renaming.
		// Capture the alias if it exists.
		if imp.Name != nil {
			stage.MetaPackageImportAlias = imp.Name.Name
		}
		// Store the path (including quotes for the marshaller)
		stage.MetaPackageImportPath = imp.Path.Value
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
						// cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						var cmap ast.CommentMap
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								stage, &cmap, assignStmt, astCoordinate, preserveOrder)
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
var __gong__map_Astruct = make(map[string]*Astruct)
var __gong__map_AstructBstruct2Use = make(map[string]*AstructBstruct2Use)
var __gong__map_AstructBstructUse = make(map[string]*AstructBstructUse)
var __gong__map_Bstruct = make(map[string]*Bstruct)
var __gong__map_Dstruct = make(map[string]*Dstruct)
var __gong__map_F0123456789012345678901234567890 = make(map[string]*F0123456789012345678901234567890)
var __gong__map_Gstruct = make(map[string]*Gstruct)

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
func UnmarshallGongstructStaging(stage *Stage, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string, preserveOrder bool) (
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

			// 1. Detect the function call type
			var isAppend bool
			_ = isAppend
			var isSlicesInsert bool
			var isSlicesDelete bool

			if id, ok := callExpr.Fun.(*ast.Ident); ok && id.Name == "append" {
				isAppend = true
			} else if se, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				if id, ok := se.X.(*ast.Ident); ok && id.Name == "slices" {
					switch se.Sel.Name {
					case "Insert":
						isSlicesInsert = true
					case "Delete":
						isSlicesDelete = true
					}
				}
			}

			// 2. Handle slices.Delete immediately
			if isSlicesDelete {
				var start, end int
				_, _ = start, end
				if bl, ok := callExpr.Args[1].(*ast.BasicLit); ok && bl.Kind == token.INT {
					start, _ = strconv.Atoi(bl.Value)
				}
				if bl, ok := callExpr.Args[2].(*ast.BasicLit); ok && bl.Kind == token.INT {
					end, _ = strconv.Atoi(bl.Value)
				}
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if ok {
					switch gongstructName {
					// insertion point for slices.Delete code
					case "Astruct":
						switch fieldName {
						// insertion point for slices.Delete field code
						case "Anarrayofb":
							instance := __gong__map_Astruct[identifier]
							if start < len(instance.Anarrayofb) && end <= len(instance.Anarrayofb) && start < end {
								instance.Anarrayofb = slices.Delete(instance.Anarrayofb, start, end)
							}
						case "Dstruct4s":
							instance := __gong__map_Astruct[identifier]
							if start < len(instance.Dstruct4s) && end <= len(instance.Dstruct4s) && start < end {
								instance.Dstruct4s = slices.Delete(instance.Dstruct4s, start, end)
							}
						case "Anarrayofa":
							instance := __gong__map_Astruct[identifier]
							if start < len(instance.Anarrayofa) && end <= len(instance.Anarrayofa) && start < end {
								instance.Anarrayofa = slices.Delete(instance.Anarrayofa, start, end)
							}
						case "Anotherarrayofb":
							instance := __gong__map_Astruct[identifier]
							if start < len(instance.Anotherarrayofb) && end <= len(instance.Anotherarrayofb) && start < end {
								instance.Anotherarrayofb = slices.Delete(instance.Anotherarrayofb, start, end)
							}
						case "AnarrayofbUse":
							instance := __gong__map_Astruct[identifier]
							if start < len(instance.AnarrayofbUse) && end <= len(instance.AnarrayofbUse) && start < end {
								instance.AnarrayofbUse = slices.Delete(instance.AnarrayofbUse, start, end)
							}
						case "Anarrayofb2Use":
							instance := __gong__map_Astruct[identifier]
							if start < len(instance.Anarrayofb2Use) && end <= len(instance.Anarrayofb2Use) && start < end {
								instance.Anarrayofb2Use = slices.Delete(instance.Anarrayofb2Use, start, end)
							}
						}
					case "AstructBstruct2Use":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "AstructBstructUse":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "Bstruct":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "Dstruct":
						switch fieldName {
						// insertion point for slices.Delete field code
						case "Anarrayofb":
							instance := __gong__map_Dstruct[identifier]
							if start < len(instance.Anarrayofb) && end <= len(instance.Anarrayofb) && start < end {
								instance.Anarrayofb = slices.Delete(instance.Anarrayofb, start, end)
							}
						case "Gstructs":
							instance := __gong__map_Dstruct[identifier]
							if start < len(instance.Gstructs) && end <= len(instance.Gstructs) && start < end {
								instance.Gstructs = slices.Delete(instance.Gstructs, start, end)
							}
						}
					case "F0123456789012345678901234567890":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "Gstruct":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					}
				}
				return
			}

			// 3. Prepare index for slices.Insert
			var insertIndex int
			_ = insertIndex
			if isSlicesInsert {
				if bl, ok := callExpr.Args[1].(*ast.BasicLit); ok && bl.Kind == token.INT {
					insertIndex, _ = strconv.Atoi(bl.Value)
				}
			}

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
										instanceAstruct := new(Astruct)
										instanceAstruct.Name = instanceName
										if !preserveOrder {
											instanceAstruct.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceAstruct.Stage(stage)
											} else {
												instanceAstruct.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceAstruct)
										__gong__map_Astruct[identifier] = instanceAstruct
									case "AstructBstruct2Use":
										instanceAstructBstruct2Use := new(AstructBstruct2Use)
										instanceAstructBstruct2Use.Name = instanceName
										if !preserveOrder {
											instanceAstructBstruct2Use.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceAstructBstruct2Use.Stage(stage)
											} else {
												instanceAstructBstruct2Use.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceAstructBstruct2Use)
										__gong__map_AstructBstruct2Use[identifier] = instanceAstructBstruct2Use
									case "AstructBstructUse":
										instanceAstructBstructUse := new(AstructBstructUse)
										instanceAstructBstructUse.Name = instanceName
										if !preserveOrder {
											instanceAstructBstructUse.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceAstructBstructUse.Stage(stage)
											} else {
												instanceAstructBstructUse.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceAstructBstructUse)
										__gong__map_AstructBstructUse[identifier] = instanceAstructBstructUse
									case "Bstruct":
										instanceBstruct := new(Bstruct)
										instanceBstruct.Name = instanceName
										if !preserveOrder {
											instanceBstruct.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceBstruct.Stage(stage)
											} else {
												instanceBstruct.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceBstruct)
										__gong__map_Bstruct[identifier] = instanceBstruct
									case "Dstruct":
										instanceDstruct := new(Dstruct)
										instanceDstruct.Name = instanceName
										if !preserveOrder {
											instanceDstruct.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceDstruct.Stage(stage)
											} else {
												instanceDstruct.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceDstruct)
										__gong__map_Dstruct[identifier] = instanceDstruct
									case "F0123456789012345678901234567890":
										instanceF0123456789012345678901234567890 := new(F0123456789012345678901234567890)
										instanceF0123456789012345678901234567890.Name = instanceName
										if !preserveOrder {
											instanceF0123456789012345678901234567890.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceF0123456789012345678901234567890.Stage(stage)
											} else {
												instanceF0123456789012345678901234567890.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceF0123456789012345678901234567890)
										__gong__map_F0123456789012345678901234567890[identifier] = instanceF0123456789012345678901234567890
									case "Gstruct":
										instanceGstruct := new(Gstruct)
										instanceGstruct.Name = instanceName
										if !preserveOrder {
											instanceGstruct.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceGstruct.Stage(stage)
											} else {
												instanceGstruct.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceGstruct)
										__gong__map_Gstruct[identifier] = instanceGstruct
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
						// Only remove first and last char if it is a STRING literal
						// Indices in slices.Insert/Delete are INT literals and must not be trimmed
						var date string
						if basicLit.Kind == token.STRING {
							date = basicLit.Value[1 : len(basicLit.Value)-1]
						} else {
							date = basicLit.Value
						}
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Println("gongstructName not found for identifier", identifier)
							break
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
							case "Date2":
								__gong__map_Astruct[identifier].Date2, _ = time.Parse(
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
						case "F0123456789012345678901234567890":
							switch fieldName {
							// insertion point for date assign code
							case "Date":
								__gong__map_F0123456789012345678901234567890[identifier].Date, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						case "Gstruct":
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
					_ = arg
					if ident, ok = arg.(*ast.Ident); !ok {
						_ = ident
						_ = ok
						// log.Println("we are in the case of new(....)")
					}

					var se *ast.SelectorExpr
					if se, ok = arg.(*ast.SelectorExpr); ok {
						if ident, ok = se.X.(*ast.Ident); !ok {
							_ = ident
							_ = ok
							// log.Println("we are in the case of append(....)")
						}
					}

					if ident == nil {
						continue
					}

					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Println("gongstructName not found for identifier", identifier)
						break
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "Astruct":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Field":
							__gong__map_Astruct[identifier].Field = basicLit.Value
						case "Anarrayofb":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Bstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									instanceWhoseFieldIsAppended.Anarrayofb = append(instanceWhoseFieldIsAppended.Anarrayofb, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Bstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Anarrayofb) {
										instanceWhoseFieldIsAppended.Anarrayofb = slices.Insert(instanceWhoseFieldIsAppended.Anarrayofb, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Dstruct4s":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Dstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									instanceWhoseFieldIsAppended.Dstruct4s = append(instanceWhoseFieldIsAppended.Dstruct4s, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Dstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Dstruct4s) {
										instanceWhoseFieldIsAppended.Dstruct4s = slices.Insert(instanceWhoseFieldIsAppended.Dstruct4s, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Anarrayofa":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Astruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									instanceWhoseFieldIsAppended.Anarrayofa = append(instanceWhoseFieldIsAppended.Anarrayofa, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Astruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Anarrayofa) {
										instanceWhoseFieldIsAppended.Anarrayofa = slices.Insert(instanceWhoseFieldIsAppended.Anarrayofa, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Anotherarrayofb":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Bstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									instanceWhoseFieldIsAppended.Anotherarrayofb = append(instanceWhoseFieldIsAppended.Anotherarrayofb, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Bstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Anotherarrayofb) {
										instanceWhoseFieldIsAppended.Anotherarrayofb = slices.Insert(instanceWhoseFieldIsAppended.Anotherarrayofb, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "AnarrayofbUse":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_AstructBstructUse[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									instanceWhoseFieldIsAppended.AnarrayofbUse = append(instanceWhoseFieldIsAppended.AnarrayofbUse, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_AstructBstructUse[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.AnarrayofbUse) {
										instanceWhoseFieldIsAppended.AnarrayofbUse = slices.Insert(instanceWhoseFieldIsAppended.AnarrayofbUse, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Anarrayofb2Use":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_AstructBstruct2Use[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									instanceWhoseFieldIsAppended.Anarrayofb2Use = append(instanceWhoseFieldIsAppended.Anarrayofb2Use, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_AstructBstruct2Use[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Astruct[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Anarrayofb2Use) {
										instanceWhoseFieldIsAppended.Anarrayofb2Use = slices.Insert(instanceWhoseFieldIsAppended.Anarrayofb2Use, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
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
						case "Anarrayofb":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Bstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Dstruct[identifier]
									instanceWhoseFieldIsAppended.Anarrayofb = append(instanceWhoseFieldIsAppended.Anarrayofb, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Bstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Dstruct[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Anarrayofb) {
										instanceWhoseFieldIsAppended.Anarrayofb = slices.Insert(instanceWhoseFieldIsAppended.Anarrayofb, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Gstructs":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Gstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Dstruct[identifier]
									instanceWhoseFieldIsAppended.Gstructs = append(instanceWhoseFieldIsAppended.Gstructs, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Gstruct[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Dstruct[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Gstructs) {
										instanceWhoseFieldIsAppended.Gstructs = slices.Insert(instanceWhoseFieldIsAppended.Gstructs, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						}
					case "F0123456789012345678901234567890":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Gstruct":
						switch fieldName {
						// insertion point for slice of pointers assign code
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
					switch v.Op {
					case token.SUB: // token.SUB is for '-'
						exprSign = -1
					case token.ADD: // token.ADD is for '+'
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
			case "Astruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Astruct[identifier].Name = fielValue
				case "Field":
					__gong__map_Astruct[identifier].Field = basicLit.Value
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
					__gong__map_Astruct[identifier].CFloatfield = exprSign * fielValue
				case "Floatfield":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Floatfield = exprSign * fielValue
				case "Intfield":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Intfield = int(exprSign) * int(fielValue)
				case "Duration1":
					// convert string to duration
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Astruct[identifier].Duration1 = time.Duration(int(exprSign) * int(fielValue))
				case "TextFieldBespokeSize":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Astruct[identifier].TextFieldBespokeSize = fielValue
				case "TextArea":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Astruct[identifier].TextArea = fielValue
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
					__gong__map_Bstruct[identifier].Floatfield = exprSign * fielValue
				case "Floatfield2":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bstruct[identifier].Floatfield2 = exprSign * fielValue
				case "Intfield":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bstruct[identifier].Intfield = int(exprSign) * int(fielValue)
				}
			case "Dstruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Dstruct[identifier].Name = fielValue
				}
			case "F0123456789012345678901234567890":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_F0123456789012345678901234567890[identifier].Name = fielValue
				}
			case "Gstruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Gstruct[identifier].Name = fielValue
				case "Floatfield":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gstruct[identifier].Floatfield = exprSign * fielValue
				case "Floatfield2":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gstruct[identifier].Floatfield2 = exprSign * fielValue
				case "Intfield":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gstruct[identifier].Intfield = int(exprSign) * int(fielValue)
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
			case "Astruct":
				switch fieldName {
				// insertion point for field dependant code
				case "Associationtob":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Associationtob = __gong__map_Bstruct[targetIdentifier]
				case "Anotherassociationtob_2":
					targetIdentifier := ident.Name
					__gong__map_Astruct[identifier].Anotherassociationtob_2 = __gong__map_Bstruct[targetIdentifier]
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
				case "Gstruct":
					targetIdentifier := ident.Name
					__gong__map_Dstruct[identifier].Gstruct = __gong__map_Gstruct[targetIdentifier]
				}
			case "F0123456789012345678901234567890":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Gstruct":
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
				case "Astruct":
					switch fieldName {
					// insertion point for selector expr assign code
					case "Field":
						__gong__map_Astruct[identifier].Field = basicLit.Value
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
					// insertion point for selector expr assign code
					}
				case "AstructBstructUse":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Bstruct":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Dstruct":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "F0123456789012345678901234567890":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Gstruct":
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

// ExtractMiddleUint takes a formatted string and returns the extracted integer.
func ExtractMiddleUint(input string) (uint, error) {
	// Compile the Regex Pattern
	re := regexp.MustCompile(`__.*?__(\d+)_.*`)

	// Find the matches
	matches := re.FindStringSubmatch(input)

	// Validate that we found the pattern and the capture group
	if len(matches) < 2 {
		return 0, fmt.Errorf("pattern not found in string: %s", input)
	}

	// matches[0] is the whole string, matches[1] is the capture group (\d+)
	numberStr := matches[1]

	// Convert string to integer (handles leading zeros automatically)
	result, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %s to int: %v", numberStr, err)
	}

	return uint(result), nil
}
