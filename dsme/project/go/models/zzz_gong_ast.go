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
var __gong__map_Diagram = make(map[string]*Diagram)
var __gong__map_Note = make(map[string]*Note)
var __gong__map_NoteProductShape = make(map[string]*NoteProductShape)
var __gong__map_NoteShape = make(map[string]*NoteShape)
var __gong__map_NoteTaskShape = make(map[string]*NoteTaskShape)
var __gong__map_Product = make(map[string]*Product)
var __gong__map_ProductCompositionShape = make(map[string]*ProductCompositionShape)
var __gong__map_ProductShape = make(map[string]*ProductShape)
var __gong__map_Project = make(map[string]*Project)
var __gong__map_Root = make(map[string]*Root)
var __gong__map_Task = make(map[string]*Task)
var __gong__map_TaskCompositionShape = make(map[string]*TaskCompositionShape)
var __gong__map_TaskInputShape = make(map[string]*TaskInputShape)
var __gong__map_TaskOutputShape = make(map[string]*TaskOutputShape)
var __gong__map_TaskShape = make(map[string]*TaskShape)

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
					case "Diagram":
						switch fieldName {
						// insertion point for slices.Delete field code
						case "Product_Shapes":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.Product_Shapes) && end <= len(instance.Product_Shapes) && start < end {
								instance.Product_Shapes = slices.Delete(instance.Product_Shapes, start, end)
							}
						case "ProductsWhoseNodeIsExpanded":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.ProductsWhoseNodeIsExpanded) && end <= len(instance.ProductsWhoseNodeIsExpanded) && start < end {
								instance.ProductsWhoseNodeIsExpanded = slices.Delete(instance.ProductsWhoseNodeIsExpanded, start, end)
							}
						case "ProductComposition_Shapes":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.ProductComposition_Shapes) && end <= len(instance.ProductComposition_Shapes) && start < end {
								instance.ProductComposition_Shapes = slices.Delete(instance.ProductComposition_Shapes, start, end)
							}
						case "Task_Shapes":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.Task_Shapes) && end <= len(instance.Task_Shapes) && start < end {
								instance.Task_Shapes = slices.Delete(instance.Task_Shapes, start, end)
							}
						case "TasksWhoseNodeIsExpanded":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.TasksWhoseNodeIsExpanded) && end <= len(instance.TasksWhoseNodeIsExpanded) && start < end {
								instance.TasksWhoseNodeIsExpanded = slices.Delete(instance.TasksWhoseNodeIsExpanded, start, end)
							}
						case "TasksWhoseInputNodeIsExpanded":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.TasksWhoseInputNodeIsExpanded) && end <= len(instance.TasksWhoseInputNodeIsExpanded) && start < end {
								instance.TasksWhoseInputNodeIsExpanded = slices.Delete(instance.TasksWhoseInputNodeIsExpanded, start, end)
							}
						case "TasksWhoseOutputNodeIsExpanded":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.TasksWhoseOutputNodeIsExpanded) && end <= len(instance.TasksWhoseOutputNodeIsExpanded) && start < end {
								instance.TasksWhoseOutputNodeIsExpanded = slices.Delete(instance.TasksWhoseOutputNodeIsExpanded, start, end)
							}
						case "TaskComposition_Shapes":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.TaskComposition_Shapes) && end <= len(instance.TaskComposition_Shapes) && start < end {
								instance.TaskComposition_Shapes = slices.Delete(instance.TaskComposition_Shapes, start, end)
							}
						case "TaskInputShapes":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.TaskInputShapes) && end <= len(instance.TaskInputShapes) && start < end {
								instance.TaskInputShapes = slices.Delete(instance.TaskInputShapes, start, end)
							}
						case "TaskOutputShapes":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.TaskOutputShapes) && end <= len(instance.TaskOutputShapes) && start < end {
								instance.TaskOutputShapes = slices.Delete(instance.TaskOutputShapes, start, end)
							}
						case "Note_Shapes":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.Note_Shapes) && end <= len(instance.Note_Shapes) && start < end {
								instance.Note_Shapes = slices.Delete(instance.Note_Shapes, start, end)
							}
						case "NotesWhoseNodeIsExpanded":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.NotesWhoseNodeIsExpanded) && end <= len(instance.NotesWhoseNodeIsExpanded) && start < end {
								instance.NotesWhoseNodeIsExpanded = slices.Delete(instance.NotesWhoseNodeIsExpanded, start, end)
							}
						case "NoteProductShapes":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.NoteProductShapes) && end <= len(instance.NoteProductShapes) && start < end {
								instance.NoteProductShapes = slices.Delete(instance.NoteProductShapes, start, end)
							}
						case "NoteTaskShapes":
							instance := __gong__map_Diagram[identifier]
							if start < len(instance.NoteTaskShapes) && end <= len(instance.NoteTaskShapes) && start < end {
								instance.NoteTaskShapes = slices.Delete(instance.NoteTaskShapes, start, end)
							}
						}
					case "Note":
						switch fieldName {
						// insertion point for slices.Delete field code
						case "Products":
							instance := __gong__map_Note[identifier]
							if start < len(instance.Products) && end <= len(instance.Products) && start < end {
								instance.Products = slices.Delete(instance.Products, start, end)
							}
						case "Tasks":
							instance := __gong__map_Note[identifier]
							if start < len(instance.Tasks) && end <= len(instance.Tasks) && start < end {
								instance.Tasks = slices.Delete(instance.Tasks, start, end)
							}
						}
					case "NoteProductShape":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "NoteShape":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "NoteTaskShape":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "Product":
						switch fieldName {
						// insertion point for slices.Delete field code
						case "SubProducts":
							instance := __gong__map_Product[identifier]
							if start < len(instance.SubProducts) && end <= len(instance.SubProducts) && start < end {
								instance.SubProducts = slices.Delete(instance.SubProducts, start, end)
							}
						}
					case "ProductCompositionShape":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "ProductShape":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "Project":
						switch fieldName {
						// insertion point for slices.Delete field code
						case "RootProducts":
							instance := __gong__map_Project[identifier]
							if start < len(instance.RootProducts) && end <= len(instance.RootProducts) && start < end {
								instance.RootProducts = slices.Delete(instance.RootProducts, start, end)
							}
						case "RootTasks":
							instance := __gong__map_Project[identifier]
							if start < len(instance.RootTasks) && end <= len(instance.RootTasks) && start < end {
								instance.RootTasks = slices.Delete(instance.RootTasks, start, end)
							}
						case "Notes":
							instance := __gong__map_Project[identifier]
							if start < len(instance.Notes) && end <= len(instance.Notes) && start < end {
								instance.Notes = slices.Delete(instance.Notes, start, end)
							}
						case "Diagrams":
							instance := __gong__map_Project[identifier]
							if start < len(instance.Diagrams) && end <= len(instance.Diagrams) && start < end {
								instance.Diagrams = slices.Delete(instance.Diagrams, start, end)
							}
						}
					case "Root":
						switch fieldName {
						// insertion point for slices.Delete field code
						case "Projects":
							instance := __gong__map_Root[identifier]
							if start < len(instance.Projects) && end <= len(instance.Projects) && start < end {
								instance.Projects = slices.Delete(instance.Projects, start, end)
							}
						}
					case "Task":
						switch fieldName {
						// insertion point for slices.Delete field code
						case "SubTasks":
							instance := __gong__map_Task[identifier]
							if start < len(instance.SubTasks) && end <= len(instance.SubTasks) && start < end {
								instance.SubTasks = slices.Delete(instance.SubTasks, start, end)
							}
						case "Inputs":
							instance := __gong__map_Task[identifier]
							if start < len(instance.Inputs) && end <= len(instance.Inputs) && start < end {
								instance.Inputs = slices.Delete(instance.Inputs, start, end)
							}
						case "Outputs":
							instance := __gong__map_Task[identifier]
							if start < len(instance.Outputs) && end <= len(instance.Outputs) && start < end {
								instance.Outputs = slices.Delete(instance.Outputs, start, end)
							}
						}
					case "TaskCompositionShape":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "TaskInputShape":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "TaskOutputShape":
						switch fieldName {
						// insertion point for slices.Delete field code
						}
					case "TaskShape":
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
									case "Diagram":
										instanceDiagram := new(Diagram)
										instanceDiagram.Name = instanceName
										if !preserveOrder {
											instanceDiagram.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceDiagram.Stage(stage)
											} else {
												instanceDiagram.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceDiagram)
										__gong__map_Diagram[identifier] = instanceDiagram
									case "Note":
										instanceNote := new(Note)
										instanceNote.Name = instanceName
										if !preserveOrder {
											instanceNote.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceNote.Stage(stage)
											} else {
												instanceNote.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceNote)
										__gong__map_Note[identifier] = instanceNote
									case "NoteProductShape":
										instanceNoteProductShape := new(NoteProductShape)
										instanceNoteProductShape.Name = instanceName
										if !preserveOrder {
											instanceNoteProductShape.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceNoteProductShape.Stage(stage)
											} else {
												instanceNoteProductShape.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceNoteProductShape)
										__gong__map_NoteProductShape[identifier] = instanceNoteProductShape
									case "NoteShape":
										instanceNoteShape := new(NoteShape)
										instanceNoteShape.Name = instanceName
										if !preserveOrder {
											instanceNoteShape.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceNoteShape.Stage(stage)
											} else {
												instanceNoteShape.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceNoteShape)
										__gong__map_NoteShape[identifier] = instanceNoteShape
									case "NoteTaskShape":
										instanceNoteTaskShape := new(NoteTaskShape)
										instanceNoteTaskShape.Name = instanceName
										if !preserveOrder {
											instanceNoteTaskShape.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceNoteTaskShape.Stage(stage)
											} else {
												instanceNoteTaskShape.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceNoteTaskShape)
										__gong__map_NoteTaskShape[identifier] = instanceNoteTaskShape
									case "Product":
										instanceProduct := new(Product)
										instanceProduct.Name = instanceName
										if !preserveOrder {
											instanceProduct.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceProduct.Stage(stage)
											} else {
												instanceProduct.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceProduct)
										__gong__map_Product[identifier] = instanceProduct
									case "ProductCompositionShape":
										instanceProductCompositionShape := new(ProductCompositionShape)
										instanceProductCompositionShape.Name = instanceName
										if !preserveOrder {
											instanceProductCompositionShape.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceProductCompositionShape.Stage(stage)
											} else {
												instanceProductCompositionShape.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceProductCompositionShape)
										__gong__map_ProductCompositionShape[identifier] = instanceProductCompositionShape
									case "ProductShape":
										instanceProductShape := new(ProductShape)
										instanceProductShape.Name = instanceName
										if !preserveOrder {
											instanceProductShape.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceProductShape.Stage(stage)
											} else {
												instanceProductShape.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceProductShape)
										__gong__map_ProductShape[identifier] = instanceProductShape
									case "Project":
										instanceProject := new(Project)
										instanceProject.Name = instanceName
										if !preserveOrder {
											instanceProject.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceProject.Stage(stage)
											} else {
												instanceProject.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceProject)
										__gong__map_Project[identifier] = instanceProject
									case "Root":
										instanceRoot := new(Root)
										instanceRoot.Name = instanceName
										if !preserveOrder {
											instanceRoot.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceRoot.Stage(stage)
											} else {
												instanceRoot.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceRoot)
										__gong__map_Root[identifier] = instanceRoot
									case "Task":
										instanceTask := new(Task)
										instanceTask.Name = instanceName
										if !preserveOrder {
											instanceTask.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceTask.Stage(stage)
											} else {
												instanceTask.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceTask)
										__gong__map_Task[identifier] = instanceTask
									case "TaskCompositionShape":
										instanceTaskCompositionShape := new(TaskCompositionShape)
										instanceTaskCompositionShape.Name = instanceName
										if !preserveOrder {
											instanceTaskCompositionShape.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceTaskCompositionShape.Stage(stage)
											} else {
												instanceTaskCompositionShape.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceTaskCompositionShape)
										__gong__map_TaskCompositionShape[identifier] = instanceTaskCompositionShape
									case "TaskInputShape":
										instanceTaskInputShape := new(TaskInputShape)
										instanceTaskInputShape.Name = instanceName
										if !preserveOrder {
											instanceTaskInputShape.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceTaskInputShape.Stage(stage)
											} else {
												instanceTaskInputShape.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceTaskInputShape)
										__gong__map_TaskInputShape[identifier] = instanceTaskInputShape
									case "TaskOutputShape":
										instanceTaskOutputShape := new(TaskOutputShape)
										instanceTaskOutputShape.Name = instanceName
										if !preserveOrder {
											instanceTaskOutputShape.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceTaskOutputShape.Stage(stage)
											} else {
												instanceTaskOutputShape.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceTaskOutputShape)
										__gong__map_TaskOutputShape[identifier] = instanceTaskOutputShape
									case "TaskShape":
										instanceTaskShape := new(TaskShape)
										instanceTaskShape.Name = instanceName
										if !preserveOrder {
											instanceTaskShape.Stage(stage)
										} else {
											if newOrder, err := ExtractMiddleUint(identifier); err != nil {
												log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
												instanceTaskShape.Stage(stage)
											} else {
												instanceTaskShape.StagePreserveOrder(stage, newOrder)
											}
										}
										instance = any(instanceTaskShape)
										__gong__map_TaskShape[identifier] = instanceTaskShape
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
						case "Diagram":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Note":
							switch fieldName {
							// insertion point for date assign code
							}
						case "NoteProductShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "NoteShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "NoteTaskShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Product":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ProductCompositionShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ProductShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Project":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Root":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Task":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TaskCompositionShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TaskInputShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TaskOutputShape":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TaskShape":
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
					case "Diagram":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Product_Shapes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_ProductShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.Product_Shapes = append(instanceWhoseFieldIsAppended.Product_Shapes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_ProductShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Product_Shapes) {
										instanceWhoseFieldIsAppended.Product_Shapes = slices.Insert(instanceWhoseFieldIsAppended.Product_Shapes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "ProductsWhoseNodeIsExpanded":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.ProductsWhoseNodeIsExpanded = append(instanceWhoseFieldIsAppended.ProductsWhoseNodeIsExpanded, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.ProductsWhoseNodeIsExpanded) {
										instanceWhoseFieldIsAppended.ProductsWhoseNodeIsExpanded = slices.Insert(instanceWhoseFieldIsAppended.ProductsWhoseNodeIsExpanded, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "ProductComposition_Shapes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_ProductCompositionShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.ProductComposition_Shapes = append(instanceWhoseFieldIsAppended.ProductComposition_Shapes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_ProductCompositionShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.ProductComposition_Shapes) {
										instanceWhoseFieldIsAppended.ProductComposition_Shapes = slices.Insert(instanceWhoseFieldIsAppended.ProductComposition_Shapes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Task_Shapes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_TaskShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.Task_Shapes = append(instanceWhoseFieldIsAppended.Task_Shapes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_TaskShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Task_Shapes) {
										instanceWhoseFieldIsAppended.Task_Shapes = slices.Insert(instanceWhoseFieldIsAppended.Task_Shapes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "TasksWhoseNodeIsExpanded":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.TasksWhoseNodeIsExpanded = append(instanceWhoseFieldIsAppended.TasksWhoseNodeIsExpanded, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.TasksWhoseNodeIsExpanded) {
										instanceWhoseFieldIsAppended.TasksWhoseNodeIsExpanded = slices.Insert(instanceWhoseFieldIsAppended.TasksWhoseNodeIsExpanded, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "TasksWhoseInputNodeIsExpanded":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.TasksWhoseInputNodeIsExpanded = append(instanceWhoseFieldIsAppended.TasksWhoseInputNodeIsExpanded, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.TasksWhoseInputNodeIsExpanded) {
										instanceWhoseFieldIsAppended.TasksWhoseInputNodeIsExpanded = slices.Insert(instanceWhoseFieldIsAppended.TasksWhoseInputNodeIsExpanded, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "TasksWhoseOutputNodeIsExpanded":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.TasksWhoseOutputNodeIsExpanded = append(instanceWhoseFieldIsAppended.TasksWhoseOutputNodeIsExpanded, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.TasksWhoseOutputNodeIsExpanded) {
										instanceWhoseFieldIsAppended.TasksWhoseOutputNodeIsExpanded = slices.Insert(instanceWhoseFieldIsAppended.TasksWhoseOutputNodeIsExpanded, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "TaskComposition_Shapes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_TaskCompositionShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.TaskComposition_Shapes = append(instanceWhoseFieldIsAppended.TaskComposition_Shapes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_TaskCompositionShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.TaskComposition_Shapes) {
										instanceWhoseFieldIsAppended.TaskComposition_Shapes = slices.Insert(instanceWhoseFieldIsAppended.TaskComposition_Shapes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "TaskInputShapes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_TaskInputShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.TaskInputShapes = append(instanceWhoseFieldIsAppended.TaskInputShapes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_TaskInputShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.TaskInputShapes) {
										instanceWhoseFieldIsAppended.TaskInputShapes = slices.Insert(instanceWhoseFieldIsAppended.TaskInputShapes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "TaskOutputShapes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_TaskOutputShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.TaskOutputShapes = append(instanceWhoseFieldIsAppended.TaskOutputShapes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_TaskOutputShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.TaskOutputShapes) {
										instanceWhoseFieldIsAppended.TaskOutputShapes = slices.Insert(instanceWhoseFieldIsAppended.TaskOutputShapes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Note_Shapes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_NoteShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.Note_Shapes = append(instanceWhoseFieldIsAppended.Note_Shapes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_NoteShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Note_Shapes) {
										instanceWhoseFieldIsAppended.Note_Shapes = slices.Insert(instanceWhoseFieldIsAppended.Note_Shapes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "NotesWhoseNodeIsExpanded":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Note[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.NotesWhoseNodeIsExpanded = append(instanceWhoseFieldIsAppended.NotesWhoseNodeIsExpanded, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Note[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.NotesWhoseNodeIsExpanded) {
										instanceWhoseFieldIsAppended.NotesWhoseNodeIsExpanded = slices.Insert(instanceWhoseFieldIsAppended.NotesWhoseNodeIsExpanded, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "NoteProductShapes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_NoteProductShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.NoteProductShapes = append(instanceWhoseFieldIsAppended.NoteProductShapes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_NoteProductShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.NoteProductShapes) {
										instanceWhoseFieldIsAppended.NoteProductShapes = slices.Insert(instanceWhoseFieldIsAppended.NoteProductShapes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "NoteTaskShapes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_NoteTaskShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									instanceWhoseFieldIsAppended.NoteTaskShapes = append(instanceWhoseFieldIsAppended.NoteTaskShapes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_NoteTaskShape[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Diagram[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.NoteTaskShapes) {
										instanceWhoseFieldIsAppended.NoteTaskShapes = slices.Insert(instanceWhoseFieldIsAppended.NoteTaskShapes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						}
					case "Note":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Products":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Note[identifier]
									instanceWhoseFieldIsAppended.Products = append(instanceWhoseFieldIsAppended.Products, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Note[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Products) {
										instanceWhoseFieldIsAppended.Products = slices.Insert(instanceWhoseFieldIsAppended.Products, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Tasks":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Note[identifier]
									instanceWhoseFieldIsAppended.Tasks = append(instanceWhoseFieldIsAppended.Tasks, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Note[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Tasks) {
										instanceWhoseFieldIsAppended.Tasks = slices.Insert(instanceWhoseFieldIsAppended.Tasks, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						}
					case "NoteProductShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "NoteShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "NoteTaskShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Product":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SubProducts":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Product[identifier]
									instanceWhoseFieldIsAppended.SubProducts = append(instanceWhoseFieldIsAppended.SubProducts, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Product[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.SubProducts) {
										instanceWhoseFieldIsAppended.SubProducts = slices.Insert(instanceWhoseFieldIsAppended.SubProducts, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						}
					case "ProductCompositionShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ProductShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Project":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "RootProducts":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Project[identifier]
									instanceWhoseFieldIsAppended.RootProducts = append(instanceWhoseFieldIsAppended.RootProducts, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Project[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.RootProducts) {
										instanceWhoseFieldIsAppended.RootProducts = slices.Insert(instanceWhoseFieldIsAppended.RootProducts, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "RootTasks":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Project[identifier]
									instanceWhoseFieldIsAppended.RootTasks = append(instanceWhoseFieldIsAppended.RootTasks, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Project[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.RootTasks) {
										instanceWhoseFieldIsAppended.RootTasks = slices.Insert(instanceWhoseFieldIsAppended.RootTasks, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Notes":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Note[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Project[identifier]
									instanceWhoseFieldIsAppended.Notes = append(instanceWhoseFieldIsAppended.Notes, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Note[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Project[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Notes) {
										instanceWhoseFieldIsAppended.Notes = slices.Insert(instanceWhoseFieldIsAppended.Notes, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Diagrams":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Diagram[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Project[identifier]
									instanceWhoseFieldIsAppended.Diagrams = append(instanceWhoseFieldIsAppended.Diagrams, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Diagram[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Project[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Diagrams) {
										instanceWhoseFieldIsAppended.Diagrams = slices.Insert(instanceWhoseFieldIsAppended.Diagrams, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						}
					case "Root":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Projects":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Project[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Root[identifier]
									instanceWhoseFieldIsAppended.Projects = append(instanceWhoseFieldIsAppended.Projects, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Project[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Root[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Projects) {
										instanceWhoseFieldIsAppended.Projects = slices.Insert(instanceWhoseFieldIsAppended.Projects, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						}
					case "Task":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "SubTasks":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Task[identifier]
									instanceWhoseFieldIsAppended.SubTasks = append(instanceWhoseFieldIsAppended.SubTasks, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Task[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Task[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.SubTasks) {
										instanceWhoseFieldIsAppended.SubTasks = slices.Insert(instanceWhoseFieldIsAppended.SubTasks, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Inputs":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Task[identifier]
									instanceWhoseFieldIsAppended.Inputs = append(instanceWhoseFieldIsAppended.Inputs, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Task[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Inputs) {
										instanceWhoseFieldIsAppended.Inputs = slices.Insert(instanceWhoseFieldIsAppended.Inputs, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						case "Outputs":
							// Handle append: elements start at argNb 1
							if isAppend && argNb > 0 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Task[identifier]
									instanceWhoseFieldIsAppended.Outputs = append(instanceWhoseFieldIsAppended.Outputs, instanceToAppend)
								}
							}
							// Handle slices.Insert: elements start at argNb 2
							if isSlicesInsert && argNb > 1 {
								identifierOfInstanceToAppend := ident.Name
								if instanceToAppend, ok := __gong__map_Product[identifierOfInstanceToAppend]; ok {
									instanceWhoseFieldIsAppended := __gong__map_Task[identifier]
									if insertIndex <= len(instanceWhoseFieldIsAppended.Outputs) {
										instanceWhoseFieldIsAppended.Outputs = slices.Insert(instanceWhoseFieldIsAppended.Outputs, insertIndex, instanceToAppend)
										insertIndex++ // Increment for subsequent elements in the same call
									}
								}
							}
						}
					case "TaskCompositionShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "TaskInputShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "TaskOutputShape":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "TaskShape":
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
			case "Diagram":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Diagram[identifier].Name = fielValue
				case "DefaultBoxWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].DefaultBoxWidth = exprSign * fielValue
				case "DefaultBoxHeigth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].DefaultBoxHeigth = exprSign * fielValue
				case "ComputedPrefix":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Diagram[identifier].ComputedPrefix = fielValue
				}
			case "Note":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Note[identifier].Name = fielValue
				}
			case "NoteProductShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteProductShape[identifier].Name = fielValue
				case "StartRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteProductShape[identifier].StartRatio = exprSign * fielValue
				case "EndRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteProductShape[identifier].EndRatio = exprSign * fielValue
				case "CornerOffsetRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteProductShape[identifier].CornerOffsetRatio = exprSign * fielValue
				}
			case "NoteShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteShape[identifier].Name = fielValue
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
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].Height = exprSign * fielValue
				}
			case "NoteTaskShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_NoteTaskShape[identifier].Name = fielValue
				case "StartRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteTaskShape[identifier].StartRatio = exprSign * fielValue
				case "EndRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteTaskShape[identifier].EndRatio = exprSign * fielValue
				case "CornerOffsetRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteTaskShape[identifier].CornerOffsetRatio = exprSign * fielValue
				}
			case "Product":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Product[identifier].Name = fielValue
				case "Description":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Product[identifier].Description = fielValue
				case "ComputedPrefix":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Product[identifier].ComputedPrefix = fielValue
				}
			case "ProductCompositionShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ProductCompositionShape[identifier].Name = fielValue
				case "StartRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ProductCompositionShape[identifier].StartRatio = exprSign * fielValue
				case "EndRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ProductCompositionShape[identifier].EndRatio = exprSign * fielValue
				case "CornerOffsetRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ProductCompositionShape[identifier].CornerOffsetRatio = exprSign * fielValue
				}
			case "ProductShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ProductShape[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ProductShape[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ProductShape[identifier].Y = exprSign * fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ProductShape[identifier].Width = exprSign * fielValue
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ProductShape[identifier].Height = exprSign * fielValue
				}
			case "Project":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Project[identifier].Name = fielValue
				case "ComputedPrefix":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Project[identifier].ComputedPrefix = fielValue
				}
			case "Root":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Root[identifier].Name = fielValue
				case "NbPixPerCharacter":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Root[identifier].NbPixPerCharacter = exprSign * fielValue
				}
			case "Task":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Task[identifier].Name = fielValue
				case "Description":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Task[identifier].Description = fielValue
				case "ComputedPrefix":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Task[identifier].ComputedPrefix = fielValue
				}
			case "TaskCompositionShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TaskCompositionShape[identifier].Name = fielValue
				case "StartRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskCompositionShape[identifier].StartRatio = exprSign * fielValue
				case "EndRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskCompositionShape[identifier].EndRatio = exprSign * fielValue
				case "CornerOffsetRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskCompositionShape[identifier].CornerOffsetRatio = exprSign * fielValue
				}
			case "TaskInputShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TaskInputShape[identifier].Name = fielValue
				case "StartRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskInputShape[identifier].StartRatio = exprSign * fielValue
				case "EndRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskInputShape[identifier].EndRatio = exprSign * fielValue
				case "CornerOffsetRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskInputShape[identifier].CornerOffsetRatio = exprSign * fielValue
				}
			case "TaskOutputShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TaskOutputShape[identifier].Name = fielValue
				case "StartRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskOutputShape[identifier].StartRatio = exprSign * fielValue
				case "EndRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskOutputShape[identifier].EndRatio = exprSign * fielValue
				case "CornerOffsetRatio":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskOutputShape[identifier].CornerOffsetRatio = exprSign * fielValue
				}
			case "TaskShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TaskShape[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskShape[identifier].X = exprSign * fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskShape[identifier].Y = exprSign * fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskShape[identifier].Width = exprSign * fielValue
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskShape[identifier].Height = exprSign * fielValue
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
			case "Diagram":
				switch fieldName {
				// insertion point for field dependant code
				case "IsChecked":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].IsChecked = fielValue
				case "IsEditable_":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].IsEditable_ = fielValue
				case "IsInRenameMode":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].IsInRenameMode = fielValue
				case "ShowPrefix":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].ShowPrefix = fielValue
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].IsExpanded = fielValue
				case "IsPBSNodeExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].IsPBSNodeExpanded = fielValue
				case "IsWBSNodeExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].IsWBSNodeExpanded = fielValue
				case "IsNotesNodeExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Diagram[identifier].IsNotesNodeExpanded = fielValue
				}
			case "Note":
				switch fieldName {
				// insertion point for field dependant code
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Note[identifier].IsExpanded = fielValue
				}
			case "NoteProductShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Note":
					targetIdentifier := ident.Name
					__gong__map_NoteProductShape[identifier].Note = __gong__map_Note[targetIdentifier]
				case "Product":
					targetIdentifier := ident.Name
					__gong__map_NoteProductShape[identifier].Product = __gong__map_Product[targetIdentifier]
				}
			case "NoteShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Note":
					targetIdentifier := ident.Name
					__gong__map_NoteShape[identifier].Note = __gong__map_Note[targetIdentifier]
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_NoteShape[identifier].IsExpanded = fielValue
				}
			case "NoteTaskShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Note":
					targetIdentifier := ident.Name
					__gong__map_NoteTaskShape[identifier].Note = __gong__map_Note[targetIdentifier]
				case "Task":
					targetIdentifier := ident.Name
					__gong__map_NoteTaskShape[identifier].Task = __gong__map_Task[targetIdentifier]
				}
			case "Product":
				switch fieldName {
				// insertion point for field dependant code
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Product[identifier].IsExpanded = fielValue
				case "IsProducersNodeExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Product[identifier].IsProducersNodeExpanded = fielValue
				case "IsConsumersNodeExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Product[identifier].IsConsumersNodeExpanded = fielValue
				}
			case "ProductCompositionShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Product":
					targetIdentifier := ident.Name
					__gong__map_ProductCompositionShape[identifier].Product = __gong__map_Product[targetIdentifier]
				}
			case "ProductShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Product":
					targetIdentifier := ident.Name
					__gong__map_ProductShape[identifier].Product = __gong__map_Product[targetIdentifier]
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ProductShape[identifier].IsExpanded = fielValue
				}
			case "Project":
				switch fieldName {
				// insertion point for field dependant code
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Project[identifier].IsExpanded = fielValue
				}
			case "Root":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Task":
				switch fieldName {
				// insertion point for field dependant code
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Task[identifier].IsExpanded = fielValue
				case "IsInputsNodeExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Task[identifier].IsInputsNodeExpanded = fielValue
				case "IsOutputsNodeExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Task[identifier].IsOutputsNodeExpanded = fielValue
				case "IsWithCompletion":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Task[identifier].IsWithCompletion = fielValue
				}
			case "TaskCompositionShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Task":
					targetIdentifier := ident.Name
					__gong__map_TaskCompositionShape[identifier].Task = __gong__map_Task[targetIdentifier]
				}
			case "TaskInputShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Task":
					targetIdentifier := ident.Name
					__gong__map_TaskInputShape[identifier].Task = __gong__map_Task[targetIdentifier]
				case "Product":
					targetIdentifier := ident.Name
					__gong__map_TaskInputShape[identifier].Product = __gong__map_Product[targetIdentifier]
				}
			case "TaskOutputShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Task":
					targetIdentifier := ident.Name
					__gong__map_TaskOutputShape[identifier].Task = __gong__map_Task[targetIdentifier]
				case "Product":
					targetIdentifier := ident.Name
					__gong__map_TaskOutputShape[identifier].Product = __gong__map_Product[targetIdentifier]
				}
			case "TaskShape":
				switch fieldName {
				// insertion point for field dependant code
				case "Task":
					targetIdentifier := ident.Name
					__gong__map_TaskShape[identifier].Task = __gong__map_Task[targetIdentifier]
				case "IsExpanded":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_TaskShape[identifier].IsExpanded = fielValue
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
				case "Diagram":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Note":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "NoteProductShape":
					switch fieldName {
					// insertion point for selector expr assign code
					case "StartOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_NoteProductShape[identifier].StartOrientation = OrientationType(val)
					case "EndOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_NoteProductShape[identifier].EndOrientation = OrientationType(val)
					}
				case "NoteShape":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "NoteTaskShape":
					switch fieldName {
					// insertion point for selector expr assign code
					case "StartOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_NoteTaskShape[identifier].StartOrientation = OrientationType(val)
					case "EndOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_NoteTaskShape[identifier].EndOrientation = OrientationType(val)
					}
				case "Product":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ProductCompositionShape":
					switch fieldName {
					// insertion point for selector expr assign code
					case "StartOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_ProductCompositionShape[identifier].StartOrientation = OrientationType(val)
					case "EndOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_ProductCompositionShape[identifier].EndOrientation = OrientationType(val)
					}
				case "ProductShape":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Project":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Root":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Task":
					switch fieldName {
					// insertion point for selector expr assign code
					case "Completion":
						var val CompletionEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_Task[identifier].Completion = CompletionEnum(val)
					}
				case "TaskCompositionShape":
					switch fieldName {
					// insertion point for selector expr assign code
					case "StartOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_TaskCompositionShape[identifier].StartOrientation = OrientationType(val)
					case "EndOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_TaskCompositionShape[identifier].EndOrientation = OrientationType(val)
					}
				case "TaskInputShape":
					switch fieldName {
					// insertion point for selector expr assign code
					case "StartOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_TaskInputShape[identifier].StartOrientation = OrientationType(val)
					case "EndOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_TaskInputShape[identifier].EndOrientation = OrientationType(val)
					}
				case "TaskOutputShape":
					switch fieldName {
					// insertion point for selector expr assign code
					case "StartOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_TaskOutputShape[identifier].StartOrientation = OrientationType(val)
					case "EndOrientation":
						var val OrientationType
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_TaskOutputShape[identifier].EndOrientation = OrientationType(val)
					}
				case "TaskShape":
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
