// generated code - do not edit
package models

import (
	"bufio"
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
var dummy_time_import time.Time

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
var __gong__map_Cell = make(map[string]*Cell)
var __gong__map_CellBoolean = make(map[string]*CellBoolean)
var __gong__map_CellFloat64 = make(map[string]*CellFloat64)
var __gong__map_CellIcon = make(map[string]*CellIcon)
var __gong__map_CellInt = make(map[string]*CellInt)
var __gong__map_CellString = make(map[string]*CellString)
var __gong__map_CheckBox = make(map[string]*CheckBox)
var __gong__map_DisplayedColumn = make(map[string]*DisplayedColumn)
var __gong__map_FormDiv = make(map[string]*FormDiv)
var __gong__map_FormEditAssocButton = make(map[string]*FormEditAssocButton)
var __gong__map_FormField = make(map[string]*FormField)
var __gong__map_FormFieldDate = make(map[string]*FormFieldDate)
var __gong__map_FormFieldDateTime = make(map[string]*FormFieldDateTime)
var __gong__map_FormFieldFloat64 = make(map[string]*FormFieldFloat64)
var __gong__map_FormFieldInt = make(map[string]*FormFieldInt)
var __gong__map_FormFieldSelect = make(map[string]*FormFieldSelect)
var __gong__map_FormFieldString = make(map[string]*FormFieldString)
var __gong__map_FormFieldTime = make(map[string]*FormFieldTime)
var __gong__map_FormGroup = make(map[string]*FormGroup)
var __gong__map_FormSortAssocButton = make(map[string]*FormSortAssocButton)
var __gong__map_Option = make(map[string]*Option)
var __gong__map_Row = make(map[string]*Row)
var __gong__map_Table = make(map[string]*Table)

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
									case "Cell":
										instanceCell := new(Cell)
										instanceCell.Name = instanceName
										instanceCell.Stage(stage)
										instance = any(instanceCell)
										__gong__map_Cell[identifier] = instanceCell
									case "CellBoolean":
										instanceCellBoolean := new(CellBoolean)
										instanceCellBoolean.Name = instanceName
										instanceCellBoolean.Stage(stage)
										instance = any(instanceCellBoolean)
										__gong__map_CellBoolean[identifier] = instanceCellBoolean
									case "CellFloat64":
										instanceCellFloat64 := new(CellFloat64)
										instanceCellFloat64.Name = instanceName
										instanceCellFloat64.Stage(stage)
										instance = any(instanceCellFloat64)
										__gong__map_CellFloat64[identifier] = instanceCellFloat64
									case "CellIcon":
										instanceCellIcon := new(CellIcon)
										instanceCellIcon.Name = instanceName
										instanceCellIcon.Stage(stage)
										instance = any(instanceCellIcon)
										__gong__map_CellIcon[identifier] = instanceCellIcon
									case "CellInt":
										instanceCellInt := new(CellInt)
										instanceCellInt.Name = instanceName
										instanceCellInt.Stage(stage)
										instance = any(instanceCellInt)
										__gong__map_CellInt[identifier] = instanceCellInt
									case "CellString":
										instanceCellString := new(CellString)
										instanceCellString.Name = instanceName
										instanceCellString.Stage(stage)
										instance = any(instanceCellString)
										__gong__map_CellString[identifier] = instanceCellString
									case "CheckBox":
										instanceCheckBox := new(CheckBox)
										instanceCheckBox.Name = instanceName
										instanceCheckBox.Stage(stage)
										instance = any(instanceCheckBox)
										__gong__map_CheckBox[identifier] = instanceCheckBox
									case "DisplayedColumn":
										instanceDisplayedColumn := new(DisplayedColumn)
										instanceDisplayedColumn.Name = instanceName
										instanceDisplayedColumn.Stage(stage)
										instance = any(instanceDisplayedColumn)
										__gong__map_DisplayedColumn[identifier] = instanceDisplayedColumn
									case "FormDiv":
										instanceFormDiv := new(FormDiv)
										instanceFormDiv.Name = instanceName
										instanceFormDiv.Stage(stage)
										instance = any(instanceFormDiv)
										__gong__map_FormDiv[identifier] = instanceFormDiv
									case "FormEditAssocButton":
										instanceFormEditAssocButton := new(FormEditAssocButton)
										instanceFormEditAssocButton.Name = instanceName
										instanceFormEditAssocButton.Stage(stage)
										instance = any(instanceFormEditAssocButton)
										__gong__map_FormEditAssocButton[identifier] = instanceFormEditAssocButton
									case "FormField":
										instanceFormField := new(FormField)
										instanceFormField.Name = instanceName
										instanceFormField.Stage(stage)
										instance = any(instanceFormField)
										__gong__map_FormField[identifier] = instanceFormField
									case "FormFieldDate":
										instanceFormFieldDate := new(FormFieldDate)
										instanceFormFieldDate.Name = instanceName
										instanceFormFieldDate.Stage(stage)
										instance = any(instanceFormFieldDate)
										__gong__map_FormFieldDate[identifier] = instanceFormFieldDate
									case "FormFieldDateTime":
										instanceFormFieldDateTime := new(FormFieldDateTime)
										instanceFormFieldDateTime.Name = instanceName
										instanceFormFieldDateTime.Stage(stage)
										instance = any(instanceFormFieldDateTime)
										__gong__map_FormFieldDateTime[identifier] = instanceFormFieldDateTime
									case "FormFieldFloat64":
										instanceFormFieldFloat64 := new(FormFieldFloat64)
										instanceFormFieldFloat64.Name = instanceName
										instanceFormFieldFloat64.Stage(stage)
										instance = any(instanceFormFieldFloat64)
										__gong__map_FormFieldFloat64[identifier] = instanceFormFieldFloat64
									case "FormFieldInt":
										instanceFormFieldInt := new(FormFieldInt)
										instanceFormFieldInt.Name = instanceName
										instanceFormFieldInt.Stage(stage)
										instance = any(instanceFormFieldInt)
										__gong__map_FormFieldInt[identifier] = instanceFormFieldInt
									case "FormFieldSelect":
										instanceFormFieldSelect := new(FormFieldSelect)
										instanceFormFieldSelect.Name = instanceName
										instanceFormFieldSelect.Stage(stage)
										instance = any(instanceFormFieldSelect)
										__gong__map_FormFieldSelect[identifier] = instanceFormFieldSelect
									case "FormFieldString":
										instanceFormFieldString := new(FormFieldString)
										instanceFormFieldString.Name = instanceName
										instanceFormFieldString.Stage(stage)
										instance = any(instanceFormFieldString)
										__gong__map_FormFieldString[identifier] = instanceFormFieldString
									case "FormFieldTime":
										instanceFormFieldTime := new(FormFieldTime)
										instanceFormFieldTime.Name = instanceName
										instanceFormFieldTime.Stage(stage)
										instance = any(instanceFormFieldTime)
										__gong__map_FormFieldTime[identifier] = instanceFormFieldTime
									case "FormGroup":
										instanceFormGroup := new(FormGroup)
										instanceFormGroup.Name = instanceName
										instanceFormGroup.Stage(stage)
										instance = any(instanceFormGroup)
										__gong__map_FormGroup[identifier] = instanceFormGroup
									case "FormSortAssocButton":
										instanceFormSortAssocButton := new(FormSortAssocButton)
										instanceFormSortAssocButton.Name = instanceName
										instanceFormSortAssocButton.Stage(stage)
										instance = any(instanceFormSortAssocButton)
										__gong__map_FormSortAssocButton[identifier] = instanceFormSortAssocButton
									case "Option":
										instanceOption := new(Option)
										instanceOption.Name = instanceName
										instanceOption.Stage(stage)
										instance = any(instanceOption)
										__gong__map_Option[identifier] = instanceOption
									case "Row":
										instanceRow := new(Row)
										instanceRow.Name = instanceName
										instanceRow.Stage(stage)
										instance = any(instanceRow)
										__gong__map_Row[identifier] = instanceRow
									case "Table":
										instanceTable := new(Table)
										instanceTable.Name = instanceName
										instanceTable.Stage(stage)
										instance = any(instanceTable)
										__gong__map_Table[identifier] = instanceTable
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
						case "Cell":
							switch fieldName {
							// insertion point for date assign code
							}
						case "CellBoolean":
							switch fieldName {
							// insertion point for date assign code
							}
						case "CellFloat64":
							switch fieldName {
							// insertion point for date assign code
							}
						case "CellIcon":
							switch fieldName {
							// insertion point for date assign code
							}
						case "CellInt":
							switch fieldName {
							// insertion point for date assign code
							}
						case "CellString":
							switch fieldName {
							// insertion point for date assign code
							}
						case "CheckBox":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DisplayedColumn":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FormDiv":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FormEditAssocButton":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FormField":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FormFieldDate":
							switch fieldName {
							// insertion point for date assign code
							case "Value":
								__gong__map_FormFieldDate[identifier].Value, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						case "FormFieldDateTime":
							switch fieldName {
							// insertion point for date assign code
							case "Value":
								__gong__map_FormFieldDateTime[identifier].Value, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						case "FormFieldFloat64":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FormFieldInt":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FormFieldSelect":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FormFieldString":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FormFieldTime":
							switch fieldName {
							// insertion point for date assign code
							case "Value":
								__gong__map_FormFieldTime[identifier].Value, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						case "FormGroup":
							switch fieldName {
							// insertion point for date assign code
							}
						case "FormSortAssocButton":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Option":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Row":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Table":
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
					case "Cell":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "CellBoolean":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "CellFloat64":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "CellIcon":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "CellInt":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "CellString":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "CheckBox":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DisplayedColumn":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FormDiv":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "FormFields":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_FormField[targetIdentifier]
							__gong__map_FormDiv[identifier].FormFields =
								append(__gong__map_FormDiv[identifier].FormFields, target)
						case "CheckBoxs":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_CheckBox[targetIdentifier]
							__gong__map_FormDiv[identifier].CheckBoxs =
								append(__gong__map_FormDiv[identifier].CheckBoxs, target)
						}
					case "FormEditAssocButton":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FormField":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FormFieldDate":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FormFieldDateTime":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FormFieldFloat64":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FormFieldInt":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FormFieldSelect":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Options":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Option[targetIdentifier]
							__gong__map_FormFieldSelect[identifier].Options =
								append(__gong__map_FormFieldSelect[identifier].Options, target)
						}
					case "FormFieldString":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FormFieldTime":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "FormGroup":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "FormDivs":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_FormDiv[targetIdentifier]
							__gong__map_FormGroup[identifier].FormDivs =
								append(__gong__map_FormGroup[identifier].FormDivs, target)
						}
					case "FormSortAssocButton":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Option":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Row":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Cells":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Cell[targetIdentifier]
							__gong__map_Row[identifier].Cells =
								append(__gong__map_Row[identifier].Cells, target)
						}
					case "Table":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "DisplayedColumns":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_DisplayedColumn[targetIdentifier]
							__gong__map_Table[identifier].DisplayedColumns =
								append(__gong__map_Table[identifier].DisplayedColumns, target)
						case "Rows":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Row[targetIdentifier]
							__gong__map_Table[identifier].Rows =
								append(__gong__map_Table[identifier].Rows, target)
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
			case "Cell":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Cell[identifier].Name = fielValue
				}
			case "CellBoolean":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CellBoolean[identifier].Name = fielValue
				}
			case "CellFloat64":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CellFloat64[identifier].Name = fielValue
				case "Value":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_CellFloat64[identifier].Value = exprSign * fielValue
				}
			case "CellIcon":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CellIcon[identifier].Name = fielValue
				case "Icon":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CellIcon[identifier].Icon = fielValue
				}
			case "CellInt":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CellInt[identifier].Name = fielValue
				case "Value":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_CellInt[identifier].Value = int(exprSign) * int(fielValue)
				}
			case "CellString":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CellString[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CellString[identifier].Value = fielValue
				}
			case "CheckBox":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_CheckBox[identifier].Name = fielValue
				}
			case "DisplayedColumn":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DisplayedColumn[identifier].Name = fielValue
				}
			case "FormDiv":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormDiv[identifier].Name = fielValue
				}
			case "FormEditAssocButton":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormEditAssocButton[identifier].Name = fielValue
				case "Label":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormEditAssocButton[identifier].Label = fielValue
				}
			case "FormField":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormField[identifier].Name = fielValue
				case "Label":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormField[identifier].Label = fielValue
				case "Placeholder":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormField[identifier].Placeholder = fielValue
				case "BespokeWidthPx":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormField[identifier].BespokeWidthPx = int(exprSign) * int(fielValue)
				case "BespokeHeightPx":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormField[identifier].BespokeHeightPx = int(exprSign) * int(fielValue)
				}
			case "FormFieldDate":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormFieldDate[identifier].Name = fielValue
				}
			case "FormFieldDateTime":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormFieldDateTime[identifier].Name = fielValue
				}
			case "FormFieldFloat64":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormFieldFloat64[identifier].Name = fielValue
				case "Value":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldFloat64[identifier].Value = exprSign * fielValue
				case "MinValue":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldFloat64[identifier].MinValue = exprSign * fielValue
				case "MaxValue":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldFloat64[identifier].MaxValue = exprSign * fielValue
				}
			case "FormFieldInt":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormFieldInt[identifier].Name = fielValue
				case "Value":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldInt[identifier].Value = int(exprSign) * int(fielValue)
				case "MinValue":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldInt[identifier].MinValue = int(exprSign) * int(fielValue)
				case "MaxValue":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldInt[identifier].MaxValue = int(exprSign) * int(fielValue)
				}
			case "FormFieldSelect":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormFieldSelect[identifier].Name = fielValue
				}
			case "FormFieldString":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormFieldString[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormFieldString[identifier].Value = fielValue
				}
			case "FormFieldTime":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormFieldTime[identifier].Name = fielValue
				case "Step":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldTime[identifier].Step = exprSign * fielValue
				}
			case "FormGroup":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormGroup[identifier].Name = fielValue
				case "Label":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormGroup[identifier].Label = fielValue
				}
			case "FormSortAssocButton":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormSortAssocButton[identifier].Name = fielValue
				case "Label":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_FormSortAssocButton[identifier].Label = fielValue
				}
			case "Option":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Option[identifier].Name = fielValue
				}
			case "Row":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Row[identifier].Name = fielValue
				}
			case "Table":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Table[identifier].Name = fielValue
				case "NbOfStickyColumns":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Table[identifier].NbOfStickyColumns = int(exprSign) * int(fielValue)
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
			case "Cell":
				switch fieldName {
				// insertion point for field dependant code
				case "CellString":
					targetIdentifier := ident.Name
					__gong__map_Cell[identifier].CellString = __gong__map_CellString[targetIdentifier]
				case "CellFloat64":
					targetIdentifier := ident.Name
					__gong__map_Cell[identifier].CellFloat64 = __gong__map_CellFloat64[targetIdentifier]
				case "CellInt":
					targetIdentifier := ident.Name
					__gong__map_Cell[identifier].CellInt = __gong__map_CellInt[targetIdentifier]
				case "CellBool":
					targetIdentifier := ident.Name
					__gong__map_Cell[identifier].CellBool = __gong__map_CellBoolean[targetIdentifier]
				case "CellIcon":
					targetIdentifier := ident.Name
					__gong__map_Cell[identifier].CellIcon = __gong__map_CellIcon[targetIdentifier]
				}
			case "CellBoolean":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_CellBoolean[identifier].Value = fielValue
				}
			case "CellFloat64":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "CellIcon":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "CellInt":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "CellString":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "CheckBox":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_CheckBox[identifier].Value = fielValue
				}
			case "DisplayedColumn":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "FormDiv":
				switch fieldName {
				// insertion point for field dependant code
				case "FormEditAssocButton":
					targetIdentifier := ident.Name
					__gong__map_FormDiv[identifier].FormEditAssocButton = __gong__map_FormEditAssocButton[targetIdentifier]
				case "FormSortAssocButton":
					targetIdentifier := ident.Name
					__gong__map_FormDiv[identifier].FormSortAssocButton = __gong__map_FormSortAssocButton[targetIdentifier]
				}
			case "FormEditAssocButton":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "FormField":
				switch fieldName {
				// insertion point for field dependant code
				case "FormFieldString":
					targetIdentifier := ident.Name
					__gong__map_FormField[identifier].FormFieldString = __gong__map_FormFieldString[targetIdentifier]
				case "FormFieldFloat64":
					targetIdentifier := ident.Name
					__gong__map_FormField[identifier].FormFieldFloat64 = __gong__map_FormFieldFloat64[targetIdentifier]
				case "FormFieldInt":
					targetIdentifier := ident.Name
					__gong__map_FormField[identifier].FormFieldInt = __gong__map_FormFieldInt[targetIdentifier]
				case "FormFieldDate":
					targetIdentifier := ident.Name
					__gong__map_FormField[identifier].FormFieldDate = __gong__map_FormFieldDate[targetIdentifier]
				case "FormFieldTime":
					targetIdentifier := ident.Name
					__gong__map_FormField[identifier].FormFieldTime = __gong__map_FormFieldTime[targetIdentifier]
				case "FormFieldDateTime":
					targetIdentifier := ident.Name
					__gong__map_FormField[identifier].FormFieldDateTime = __gong__map_FormFieldDateTime[targetIdentifier]
				case "FormFieldSelect":
					targetIdentifier := ident.Name
					__gong__map_FormField[identifier].FormFieldSelect = __gong__map_FormFieldSelect[targetIdentifier]
				case "HasBespokeWidth":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormField[identifier].HasBespokeWidth = fielValue
				case "HasBespokeHeight":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormField[identifier].HasBespokeHeight = fielValue
				}
			case "FormFieldDate":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "FormFieldDateTime":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "FormFieldFloat64":
				switch fieldName {
				// insertion point for field dependant code
				case "HasMinValidator":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldFloat64[identifier].HasMinValidator = fielValue
				case "HasMaxValidator":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldFloat64[identifier].HasMaxValidator = fielValue
				}
			case "FormFieldInt":
				switch fieldName {
				// insertion point for field dependant code
				case "HasMinValidator":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldInt[identifier].HasMinValidator = fielValue
				case "HasMaxValidator":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldInt[identifier].HasMaxValidator = fielValue
				}
			case "FormFieldSelect":
				switch fieldName {
				// insertion point for field dependant code
				case "Value":
					targetIdentifier := ident.Name
					__gong__map_FormFieldSelect[identifier].Value = __gong__map_Option[targetIdentifier]
				case "CanBeEmpty":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldSelect[identifier].CanBeEmpty = fielValue
				}
			case "FormFieldString":
				switch fieldName {
				// insertion point for field dependant code
				case "IsTextArea":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormFieldString[identifier].IsTextArea = fielValue
				}
			case "FormFieldTime":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "FormGroup":
				switch fieldName {
				// insertion point for field dependant code
				case "HasSuppressButton":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormGroup[identifier].HasSuppressButton = fielValue
				case "HasSuppressButtonBeenPressed":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_FormGroup[identifier].HasSuppressButtonBeenPressed = fielValue
				}
			case "FormSortAssocButton":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Option":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Row":
				switch fieldName {
				// insertion point for field dependant code
				case "IsChecked":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Row[identifier].IsChecked = fielValue
				}
			case "Table":
				switch fieldName {
				// insertion point for field dependant code
				case "HasFiltering":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Table[identifier].HasFiltering = fielValue
				case "HasColumnSorting":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Table[identifier].HasColumnSorting = fielValue
				case "HasPaginator":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Table[identifier].HasPaginator = fielValue
				case "HasCheckableRows":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Table[identifier].HasCheckableRows = fielValue
				case "HasSaveButton":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Table[identifier].HasSaveButton = fielValue
				case "CanDragDropRows":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Table[identifier].CanDragDropRows = fielValue
				case "HasCloseButton":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Table[identifier].HasCloseButton = fielValue
				case "SavingInProgress":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Table[identifier].SavingInProgress = fielValue
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
				case "Cell":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "CellBoolean":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "CellFloat64":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "CellIcon":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "CellInt":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "CellString":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "CheckBox":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DisplayedColumn":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormDiv":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormEditAssocButton":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormField":
					switch fieldName {
					// insertion point for enum assign code
					case "InputTypeEnum":
						var val InputTypeEnum
						err := (&val).FromCodeString(enumValue)
						if err != nil {
							log.Fatalln(err)
						}
						__gong__map_FormField[identifier].InputTypeEnum = InputTypeEnum(val)
					}
				case "FormFieldDate":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormFieldDateTime":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormFieldFloat64":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormFieldInt":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormFieldSelect":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormFieldString":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormFieldTime":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormGroup":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "FormSortAssocButton":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Option":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Row":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Table":
					switch fieldName {
					// insertion point for enum assign code
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
