// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"slices"
	"strings"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
)

var (
	_ = errors.New
	_ = parser.ParseFile
	_ = filepath.Base
	_ = slices.Index[[]int, int]
	_ = strings.Trim
	_ = time.Hour
)

// swagger:ignore
type GONG__ExpressionType = gong_runtime.GONG__ExpressionType

const (
	GONG__STRUCT_INSTANCE      = gong_runtime.GONG__STRUCT_INSTANCE
	GONG__FIELD_OR_CONST_VALUE = gong_runtime.GONG__FIELD_OR_CONST_VALUE
	GONG__FIELD_VALUE          = gong_runtime.GONG__FIELD_VALUE
	GONG__ENUM_CAST_INT        = gong_runtime.GONG__ENUM_CAST_INT
	GONG__ENUM_CAST_STRING     = gong_runtime.GONG__ENUM_CAST_STRING
	GONG__IDENTIFIER_CONST     = gong_runtime.GONG__IDENTIFIER_CONST
)

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// GongModelUnmarshaller abstracts the logic for setting fields on a staged instance
type GongModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

type ModelUnmarshaller = GongModelUnmarshaller

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func (stage *Stage) ParseAstFile(pathToFile string, preserveOrder bool) error {
	inFile, fset, err := gong_runtime.ParseAstFile(pathToFile)
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
func (stage *Stage) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	inFile, fset, err := gong_runtime.ParseAstEmbeddedFile(directory, pathToFile, stage.GetName())
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string
func (stage *Stage) ParseAstString(blob string, preserveOrder bool) error {
	var inFile *ast.File
	var fset *token.FileSet
	var err error
	if strings.HasPrefix(strings.TrimSpace(blob), "package ") || strings.Contains(blob, "\npackage ") {
		fset = token.NewFileSet()
		inFile, err = parser.ParseFile(fset, "", blob, parser.ParseComments)
		if err != nil {
			return errors.New("Unable to parse " + err.Error())
		}
	} else {
		inFile, fset, err = gong_runtime.ParseAstString(blob)
		if err != nil {
			return err
		}
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances across stage and its sub-stages
func (stage *Stage) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	gong_runtime.CheckModuleVersion(stage.GetProbeIF(), inFile)

	identifierMap := make(map[string]any)

	if stage != nil {
		for _, instance := range stage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage)] = instance
			identifierMap["__models"+instance.GongGetIdentifier(stage)] = instance
		}
	}

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/app/reqif/go/models":
			aliasToCanonical[alias] = "models"
		}
	}

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					} else if pkgAlias == "" || pkgAlias == "models" || len(inFile.Imports) == 0 {
						pkgAlias = "models"
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "ALTERNATIVE_ID":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ALTERNATIVE_ID), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_BOOLEAN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_BOOLEAN), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_BOOLEAN_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_BOOLEAN_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_DATE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_DATE), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_DATE_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_DATE_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_ENUMERATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_ENUMERATION), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_ENUMERATION_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_ENUMERATION_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_INTEGER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_INTEGER), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_INTEGER_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_INTEGER_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_REAL":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_REAL), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_REAL_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_REAL_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_STRING":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_STRING), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_STRING_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_STRING_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_XHTML":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_XHTML), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_XHTML_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_XHTML_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_BOOLEAN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_BOOLEAN), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_DATE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_DATE), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_ENUMERATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_ENUMERATION), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_INTEGER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_INTEGER), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_REAL":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_REAL), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_STRING":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_STRING), stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_XHTML":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_XHTML), stage, ident.Name, instanceName, preserveOrder)
				case "A_ALTERNATIVE_ID":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ALTERNATIVE_ID), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_DATE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_DATE_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_INTEGER_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_REAL_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_REAL_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_STRING_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_STRING_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_XHTML_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_BOOLEAN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_BOOLEAN), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_DATE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_DATE), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_ENUMERATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_ENUMERATION), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_INTEGER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_INTEGER), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_REAL":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_REAL), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_STRING":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_STRING), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_XHTML":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_XHTML), stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_XHTML_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_XHTML_1), stage, ident.Name, instanceName, preserveOrder)
				case "A_CHILDREN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_CHILDREN), stage, ident.Name, instanceName, preserveOrder)
				case "A_CORE_CONTENT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_CORE_CONTENT), stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPES), stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_BOOLEAN_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_DATE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_DATE_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_ENUMERATION_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_INTEGER_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_INTEGER_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_REAL_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_REAL_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_STRING_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_STRING_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_XHTML_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_XHTML_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_EDITABLE_ATTS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_EDITABLE_ATTS), stage, ident.Name, instanceName, preserveOrder)
				case "A_ENUM_VALUE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ENUM_VALUE_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_OBJECT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_OBJECT), stage, ident.Name, instanceName, preserveOrder)
				case "A_PROPERTIES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_PROPERTIES), stage, ident.Name, instanceName, preserveOrder)
				case "A_RELATION_GROUP_TYPE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_RELATION_GROUP_TYPE_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_SOURCE_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SOURCE_1), stage, ident.Name, instanceName, preserveOrder)
				case "A_SOURCE_SPECIFICATION_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SOURCE_SPECIFICATION_1), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPECIFICATIONS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPECIFICATIONS), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPECIFICATION_TYPE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPECIFICATION_TYPE_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPECIFIED_VALUES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPECIFIED_VALUES), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_ATTRIBUTES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_ATTRIBUTES), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_OBJECTS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_OBJECTS), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_OBJECT_TYPE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_OBJECT_TYPE_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_RELATIONS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_RELATIONS), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_RELATION_GROUPS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_RELATION_GROUPS), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_RELATION_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_RELATION_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_RELATION_TYPE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_RELATION_TYPE_REF), stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_TYPES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_TYPES), stage, ident.Name, instanceName, preserveOrder)
				case "A_THE_HEADER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_THE_HEADER), stage, ident.Name, instanceName, preserveOrder)
				case "A_TOOL_EXTENSIONS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_TOOL_EXTENSIONS), stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_BOOLEAN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_BOOLEAN), stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_DATE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_DATE), stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_ENUMERATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_ENUMERATION), stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_INTEGER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_INTEGER), stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_REAL":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_REAL), stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_STRING":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_STRING), stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_XHTML":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_XHTML), stage, ident.Name, instanceName, preserveOrder)
				case "EMBEDDED_VALUE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EMBEDDED_VALUE), stage, ident.Name, instanceName, preserveOrder)
				case "ENUM_VALUE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ENUM_VALUE), stage, ident.Name, instanceName, preserveOrder)
				case "EmbeddedJpgImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmbeddedJpgImage), stage, ident.Name, instanceName, preserveOrder)
				case "EmbeddedPngImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmbeddedPngImage), stage, ident.Name, instanceName, preserveOrder)
				case "EmbeddedSvgImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmbeddedSvgImage), stage, ident.Name, instanceName, preserveOrder)
				case "Kill":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Kill), stage, ident.Name, instanceName, preserveOrder)
				case "Map_identifier_bool":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Map_identifier_bool), stage, ident.Name, instanceName, preserveOrder)
				case "RELATION_GROUP":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RELATION_GROUP), stage, ident.Name, instanceName, preserveOrder)
				case "RELATION_GROUP_TYPE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RELATION_GROUP_TYPE), stage, ident.Name, instanceName, preserveOrder)
				case "REQ_IF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(REQ_IF), stage, ident.Name, instanceName, preserveOrder)
				case "REQ_IF_CONTENT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(REQ_IF_CONTENT), stage, ident.Name, instanceName, preserveOrder)
				case "REQ_IF_HEADER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(REQ_IF_HEADER), stage, ident.Name, instanceName, preserveOrder)
				case "REQ_IF_TOOL_EXTENSION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(REQ_IF_TOOL_EXTENSION), stage, ident.Name, instanceName, preserveOrder)
				case "SPECIFICATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPECIFICATION), stage, ident.Name, instanceName, preserveOrder)
				case "SPECIFICATION_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPECIFICATION_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "SPECIFICATION_TYPE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPECIFICATION_TYPE), stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_HIERARCHY":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_HIERARCHY), stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_OBJECT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_OBJECT), stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_OBJECT_TYPE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_OBJECT_TYPE), stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_OBJECT_TYPE_Rendering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_OBJECT_TYPE_Rendering), stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_RELATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_RELATION), stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_RELATION_TYPE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_RELATION_TYPE), stage, ident.Name, instanceName, preserveOrder)
				case "StaticWebSite":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaticWebSite), stage, ident.Name, instanceName, preserveOrder)
				case "StaticWebSiteChapter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaticWebSiteChapter), stage, ident.Name, instanceName, preserveOrder)
				case "StaticWebSiteGeneratedImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaticWebSiteGeneratedImage), stage, ident.Name, instanceName, preserveOrder)
				case "StaticWebSiteImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaticWebSiteImage), stage, ident.Name, instanceName, preserveOrder)
				case "StaticWebSiteParagraph":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaticWebSiteParagraph), stage, ident.Name, instanceName, preserveOrder)
				case "XHTML_CONTENT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(XHTML_CONTENT), stage, ident.Name, instanceName, preserveOrder)
				}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *ALTERNATIVE_ID:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					}
				case *ATTRIBUTE_DEFINITION_BOOLEAN:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "IS_EDITABLE":
						inst.IS_EDITABLE = GongExtractBool(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "DEFAULT_VALUE":
						__gong__assignPointer(&inst.DEFAULT_VALUE, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowInTable":
						inst.ShowInTable = GongExtractBool(rhs)
					case "ShowInTitle":
						inst.ShowInTitle = GongExtractBool(rhs)
					case "ShowInSubject":
						inst.ShowInSubject = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_DEFINITION_DATE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "IS_EDITABLE":
						inst.IS_EDITABLE = GongExtractBool(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "DEFAULT_VALUE":
						__gong__assignPointer(&inst.DEFAULT_VALUE, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *ATTRIBUTE_DEFINITION_DATE_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowInTable":
						inst.ShowInTable = GongExtractBool(rhs)
					case "ShowInTitle":
						inst.ShowInTitle = GongExtractBool(rhs)
					case "ShowInSubject":
						inst.ShowInSubject = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_DEFINITION_ENUMERATION:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "IS_EDITABLE":
						inst.IS_EDITABLE = GongExtractBool(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "MULTI_VALUED":
						inst.MULTI_VALUED = GongExtractBool(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "DEFAULT_VALUE":
						__gong__assignPointer(&inst.DEFAULT_VALUE, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowInTable":
						inst.ShowInTable = GongExtractBool(rhs)
					case "ShowInTitle":
						inst.ShowInTitle = GongExtractBool(rhs)
					case "ShowInSubject":
						inst.ShowInSubject = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_DEFINITION_INTEGER:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "IS_EDITABLE":
						inst.IS_EDITABLE = GongExtractBool(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "DEFAULT_VALUE":
						__gong__assignPointer(&inst.DEFAULT_VALUE, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *ATTRIBUTE_DEFINITION_INTEGER_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowInTable":
						inst.ShowInTable = GongExtractBool(rhs)
					case "ShowInTitle":
						inst.ShowInTitle = GongExtractBool(rhs)
					case "ShowInSubject":
						inst.ShowInSubject = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_DEFINITION_REAL:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "IS_EDITABLE":
						inst.IS_EDITABLE = GongExtractBool(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "DEFAULT_VALUE":
						__gong__assignPointer(&inst.DEFAULT_VALUE, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *ATTRIBUTE_DEFINITION_REAL_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowInTable":
						inst.ShowInTable = GongExtractBool(rhs)
					case "ShowInTitle":
						inst.ShowInTitle = GongExtractBool(rhs)
					case "ShowInSubject":
						inst.ShowInSubject = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_DEFINITION_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowInTable":
						inst.ShowInTable = GongExtractBool(rhs)
					case "ShowInTitle":
						inst.ShowInTitle = GongExtractBool(rhs)
					case "ShowInSubject":
						inst.ShowInSubject = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_DEFINITION_STRING:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "IS_EDITABLE":
						inst.IS_EDITABLE = GongExtractBool(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "DEFAULT_VALUE":
						__gong__assignPointer(&inst.DEFAULT_VALUE, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *ATTRIBUTE_DEFINITION_STRING_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowInTable":
						inst.ShowInTable = GongExtractBool(rhs)
					case "ShowInTitle":
						inst.ShowInTitle = GongExtractBool(rhs)
					case "ShowInSubject":
						inst.ShowInSubject = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_DEFINITION_XHTML:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "IS_EDITABLE":
						inst.IS_EDITABLE = GongExtractBool(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "DEFAULT_VALUE":
						__gong__assignPointer(&inst.DEFAULT_VALUE, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *ATTRIBUTE_DEFINITION_XHTML_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowInTable":
						inst.ShowInTable = GongExtractBool(rhs)
					case "ShowInTitle":
						inst.ShowInTitle = GongExtractBool(rhs)
					case "ShowInSubject":
						inst.ShowInSubject = GongExtractBool(rhs)
					case "Rank":
						inst.Rank = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_VALUE_BOOLEAN:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractBool(rhs)
					}
				case *ATTRIBUTE_VALUE_DATE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractString(rhs)
					}
				case *ATTRIBUTE_VALUE_ENUMERATION:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					case "VALUES":
						__gong__assignPointer(&inst.VALUES, rhs, identifierMap)
					}
				case *ATTRIBUTE_VALUE_INTEGER:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_VALUE_REAL:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractFloat(rhs)
					}
				case *ATTRIBUTE_VALUE_STRING:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractString(rhs)
					}
				case *ATTRIBUTE_VALUE_XHTML:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					case "IS_SIMPLIFIED":
						inst.IS_SIMPLIFIED = GongExtractBool(rhs)
					case "THE_VALUE":
						__gong__assignPointer(&inst.THE_VALUE, rhs, identifierMap)
					case "THE_ORIGINAL_VALUE":
						__gong__assignPointer(&inst.THE_ORIGINAL_VALUE, rhs, identifierMap)
					}
				case *A_ALTERNATIVE_ID:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					}
				case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
						inst.ATTRIBUTE_DEFINITION_BOOLEAN_REF = GongExtractString(rhs)
					}
				case *A_ATTRIBUTE_DEFINITION_DATE_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_DATE_REF":
						inst.ATTRIBUTE_DEFINITION_DATE_REF = GongExtractString(rhs)
					}
				case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
						inst.ATTRIBUTE_DEFINITION_ENUMERATION_REF = GongExtractString(rhs)
					}
				case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_INTEGER_REF":
						inst.ATTRIBUTE_DEFINITION_INTEGER_REF = GongExtractString(rhs)
					}
				case *A_ATTRIBUTE_DEFINITION_REAL_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_REAL_REF":
						inst.ATTRIBUTE_DEFINITION_REAL_REF = GongExtractString(rhs)
					}
				case *A_ATTRIBUTE_DEFINITION_STRING_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_STRING_REF":
						inst.ATTRIBUTE_DEFINITION_STRING_REF = GongExtractString(rhs)
					}
				case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_XHTML_REF":
						inst.ATTRIBUTE_DEFINITION_XHTML_REF = GongExtractString(rhs)
					}
				case *A_ATTRIBUTE_VALUE_BOOLEAN:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_BOOLEAN":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_BOOLEAN, rhs, identifierMap)
					}
				case *A_ATTRIBUTE_VALUE_DATE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_DATE":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_DATE, rhs, identifierMap)
					}
				case *A_ATTRIBUTE_VALUE_ENUMERATION:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_ENUMERATION":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_ENUMERATION, rhs, identifierMap)
					}
				case *A_ATTRIBUTE_VALUE_INTEGER:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_INTEGER":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_INTEGER, rhs, identifierMap)
					}
				case *A_ATTRIBUTE_VALUE_REAL:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_REAL":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_REAL, rhs, identifierMap)
					}
				case *A_ATTRIBUTE_VALUE_STRING:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_STRING":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_STRING, rhs, identifierMap)
					}
				case *A_ATTRIBUTE_VALUE_XHTML:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_XHTML":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_XHTML, rhs, identifierMap)
					}
				case *A_ATTRIBUTE_VALUE_XHTML_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_BOOLEAN":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_BOOLEAN, rhs, identifierMap)
					case "ATTRIBUTE_VALUE_DATE":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_DATE, rhs, identifierMap)
					case "ATTRIBUTE_VALUE_ENUMERATION":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_ENUMERATION, rhs, identifierMap)
					case "ATTRIBUTE_VALUE_INTEGER":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_INTEGER, rhs, identifierMap)
					case "ATTRIBUTE_VALUE_REAL":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_REAL, rhs, identifierMap)
					case "ATTRIBUTE_VALUE_STRING":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_STRING, rhs, identifierMap)
					case "ATTRIBUTE_VALUE_XHTML":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_VALUE_XHTML, rhs, identifierMap)
					}
				case *A_CHILDREN:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_HIERARCHY":
						__gong__assignSliceOfPointers(&inst.SPEC_HIERARCHY, rhs, identifierMap)
					}
				case *A_CORE_CONTENT:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "REQ_IF_CONTENT":
						__gong__assignPointer(&inst.REQ_IF_CONTENT, rhs, identifierMap)
					}
				case *A_DATATYPES:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPE_DEFINITION_BOOLEAN":
						__gong__assignSliceOfPointers(&inst.DATATYPE_DEFINITION_BOOLEAN, rhs, identifierMap)
					case "DATATYPE_DEFINITION_DATE":
						__gong__assignSliceOfPointers(&inst.DATATYPE_DEFINITION_DATE, rhs, identifierMap)
					case "DATATYPE_DEFINITION_ENUMERATION":
						__gong__assignSliceOfPointers(&inst.DATATYPE_DEFINITION_ENUMERATION, rhs, identifierMap)
					case "DATATYPE_DEFINITION_INTEGER":
						__gong__assignSliceOfPointers(&inst.DATATYPE_DEFINITION_INTEGER, rhs, identifierMap)
					case "DATATYPE_DEFINITION_REAL":
						__gong__assignSliceOfPointers(&inst.DATATYPE_DEFINITION_REAL, rhs, identifierMap)
					case "DATATYPE_DEFINITION_STRING":
						__gong__assignSliceOfPointers(&inst.DATATYPE_DEFINITION_STRING, rhs, identifierMap)
					case "DATATYPE_DEFINITION_XHTML":
						__gong__assignSliceOfPointers(&inst.DATATYPE_DEFINITION_XHTML, rhs, identifierMap)
					}
				case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPE_DEFINITION_BOOLEAN_REF":
						inst.DATATYPE_DEFINITION_BOOLEAN_REF = GongExtractString(rhs)
					}
				case *A_DATATYPE_DEFINITION_DATE_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPE_DEFINITION_DATE_REF":
						inst.DATATYPE_DEFINITION_DATE_REF = GongExtractString(rhs)
					}
				case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPE_DEFINITION_ENUMERATION_REF":
						inst.DATATYPE_DEFINITION_ENUMERATION_REF = GongExtractString(rhs)
					}
				case *A_DATATYPE_DEFINITION_INTEGER_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPE_DEFINITION_INTEGER_REF":
						inst.DATATYPE_DEFINITION_INTEGER_REF = GongExtractString(rhs)
					}
				case *A_DATATYPE_DEFINITION_REAL_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPE_DEFINITION_REAL_REF":
						inst.DATATYPE_DEFINITION_REAL_REF = GongExtractString(rhs)
					}
				case *A_DATATYPE_DEFINITION_STRING_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPE_DEFINITION_STRING_REF":
						inst.DATATYPE_DEFINITION_STRING_REF = GongExtractString(rhs)
					}
				case *A_DATATYPE_DEFINITION_XHTML_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPE_DEFINITION_XHTML_REF":
						inst.DATATYPE_DEFINITION_XHTML_REF = GongExtractString(rhs)
					}
				case *A_EDITABLE_ATTS:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
						inst.ATTRIBUTE_DEFINITION_BOOLEAN_REF = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_DATE_REF":
						inst.ATTRIBUTE_DEFINITION_DATE_REF = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
						inst.ATTRIBUTE_DEFINITION_ENUMERATION_REF = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_INTEGER_REF":
						inst.ATTRIBUTE_DEFINITION_INTEGER_REF = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_REAL_REF":
						inst.ATTRIBUTE_DEFINITION_REAL_REF = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_STRING_REF":
						inst.ATTRIBUTE_DEFINITION_STRING_REF = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_XHTML_REF":
						inst.ATTRIBUTE_DEFINITION_XHTML_REF = GongExtractString(rhs)
					}
				case *A_ENUM_VALUE_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ENUM_VALUE_REF":
						inst.ENUM_VALUE_REF = GongExtractString(rhs)
					}
				case *A_OBJECT:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_OBJECT_REF":
						inst.SPEC_OBJECT_REF = GongExtractString(rhs)
					}
				case *A_PROPERTIES:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EMBEDDED_VALUE":
						__gong__assignPointer(&inst.EMBEDDED_VALUE, rhs, identifierMap)
					}
				case *A_RELATION_GROUP_TYPE_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RELATION_GROUP_TYPE_REF":
						inst.RELATION_GROUP_TYPE_REF = GongExtractString(rhs)
					}
				case *A_SOURCE_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_OBJECT_REF":
						inst.SPEC_OBJECT_REF = GongExtractString(rhs)
					}
				case *A_SOURCE_SPECIFICATION_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPECIFICATION_REF":
						GongUnmarshallEnum(&inst.SPECIFICATION_REF, rhs)
					}
				case *A_SPECIFICATIONS:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPECIFICATION":
						__gong__assignSliceOfPointers(&inst.SPECIFICATION, rhs, identifierMap)
					}
				case *A_SPECIFICATION_TYPE_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPECIFICATION_TYPE_REF":
						inst.SPECIFICATION_TYPE_REF = GongExtractString(rhs)
					}
				case *A_SPECIFIED_VALUES:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ENUM_VALUE":
						__gong__assignSliceOfPointers(&inst.ENUM_VALUE, rhs, identifierMap)
					}
				case *A_SPEC_ATTRIBUTES:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_BOOLEAN":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_DEFINITION_BOOLEAN, rhs, identifierMap)
					case "ATTRIBUTE_DEFINITION_DATE":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_DEFINITION_DATE, rhs, identifierMap)
					case "ATTRIBUTE_DEFINITION_ENUMERATION":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_DEFINITION_ENUMERATION, rhs, identifierMap)
					case "ATTRIBUTE_DEFINITION_INTEGER":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_DEFINITION_INTEGER, rhs, identifierMap)
					case "ATTRIBUTE_DEFINITION_REAL":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_DEFINITION_REAL, rhs, identifierMap)
					case "ATTRIBUTE_DEFINITION_STRING":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_DEFINITION_STRING, rhs, identifierMap)
					case "ATTRIBUTE_DEFINITION_XHTML":
						__gong__assignSliceOfPointers(&inst.ATTRIBUTE_DEFINITION_XHTML, rhs, identifierMap)
					}
				case *A_SPEC_OBJECTS:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_OBJECT":
						__gong__assignSliceOfPointers(&inst.SPEC_OBJECT, rhs, identifierMap)
					}
				case *A_SPEC_OBJECT_TYPE_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_OBJECT_TYPE_REF":
						inst.SPEC_OBJECT_TYPE_REF = GongExtractString(rhs)
					}
				case *A_SPEC_RELATIONS:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_RELATION":
						__gong__assignSliceOfPointers(&inst.SPEC_RELATION, rhs, identifierMap)
					}
				case *A_SPEC_RELATION_GROUPS:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RELATION_GROUP":
						__gong__assignSliceOfPointers(&inst.RELATION_GROUP, rhs, identifierMap)
					}
				case *A_SPEC_RELATION_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_RELATION_REF":
						inst.SPEC_RELATION_REF = GongExtractString(rhs)
					}
				case *A_SPEC_RELATION_TYPE_REF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_RELATION_TYPE_REF":
						inst.SPEC_RELATION_TYPE_REF = GongExtractString(rhs)
					}
				case *A_SPEC_TYPES:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RELATION_GROUP_TYPE":
						__gong__assignSliceOfPointers(&inst.RELATION_GROUP_TYPE, rhs, identifierMap)
					case "SPEC_OBJECT_TYPE":
						__gong__assignSliceOfPointers(&inst.SPEC_OBJECT_TYPE, rhs, identifierMap)
					case "SPEC_RELATION_TYPE":
						__gong__assignSliceOfPointers(&inst.SPEC_RELATION_TYPE, rhs, identifierMap)
					case "SPECIFICATION_TYPE":
						__gong__assignSliceOfPointers(&inst.SPECIFICATION_TYPE, rhs, identifierMap)
					}
				case *A_THE_HEADER:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "REQ_IF_HEADER":
						__gong__assignPointer(&inst.REQ_IF_HEADER, rhs, identifierMap)
					}
				case *A_TOOL_EXTENSIONS:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "REQ_IF_TOOL_EXTENSION":
						__gong__assignSliceOfPointers(&inst.REQ_IF_TOOL_EXTENSION, rhs, identifierMap)
					}
				case *DATATYPE_DEFINITION_BOOLEAN:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					}
				case *DATATYPE_DEFINITION_DATE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					}
				case *DATATYPE_DEFINITION_ENUMERATION:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "SPECIFIED_VALUES":
						__gong__assignPointer(&inst.SPECIFIED_VALUES, rhs, identifierMap)
					}
				case *DATATYPE_DEFINITION_INTEGER:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "MAX":
						inst.MAX = GongExtractInt(rhs)
					case "MIN":
						inst.MIN = GongExtractInt(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					}
				case *DATATYPE_DEFINITION_REAL:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ACCURACY":
						inst.ACCURACY = GongExtractInt(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "MAX":
						inst.MAX = GongExtractFloat(rhs)
					case "MIN":
						inst.MIN = GongExtractFloat(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					}
				case *DATATYPE_DEFINITION_STRING:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "MAX_LENGTH":
						inst.MAX_LENGTH = GongExtractInt(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					}
				case *DATATYPE_DEFINITION_XHTML:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					}
				case *EMBEDDED_VALUE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "KEY":
						inst.KEY = GongExtractInt(rhs)
					case "OTHER_CONTENT":
						inst.OTHER_CONTENT = GongExtractString(rhs)
					}
				case *ENUM_VALUE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "PROPERTIES":
						__gong__assignPointer(&inst.PROPERTIES, rhs, identifierMap)
					}
				case *EmbeddedJpgImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *EmbeddedPngImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *EmbeddedSvgImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					}
				case *Kill:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Map_identifier_bool:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractBool(rhs)
					}
				case *RELATION_GROUP:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "SOURCE_SPECIFICATION":
						__gong__assignPointer(&inst.SOURCE_SPECIFICATION, rhs, identifierMap)
					case "SPEC_RELATIONS":
						__gong__assignPointer(&inst.SPEC_RELATIONS, rhs, identifierMap)
					case "TARGET_SPECIFICATION":
						__gong__assignPointer(&inst.TARGET_SPECIFICATION, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *RELATION_GROUP_TYPE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "SPEC_ATTRIBUTES":
						__gong__assignPointer(&inst.SPEC_ATTRIBUTES, rhs, identifierMap)
					}
				case *REQ_IF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "THE_HEADER":
						__gong__assignPointer(&inst.THE_HEADER, rhs, identifierMap)
					case "CORE_CONTENT":
						__gong__assignPointer(&inst.CORE_CONTENT, rhs, identifierMap)
					case "TOOL_EXTENSIONS":
						__gong__assignPointer(&inst.TOOL_EXTENSIONS, rhs, identifierMap)
					}
				case *REQ_IF_CONTENT:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPES":
						__gong__assignPointer(&inst.DATATYPES, rhs, identifierMap)
					case "SPEC_TYPES":
						__gong__assignPointer(&inst.SPEC_TYPES, rhs, identifierMap)
					case "SPEC_OBJECTS":
						__gong__assignPointer(&inst.SPEC_OBJECTS, rhs, identifierMap)
					case "SPEC_RELATIONS":
						__gong__assignPointer(&inst.SPEC_RELATIONS, rhs, identifierMap)
					case "SPECIFICATIONS":
						__gong__assignPointer(&inst.SPECIFICATIONS, rhs, identifierMap)
					case "SPEC_RELATION_GROUPS":
						__gong__assignPointer(&inst.SPEC_RELATION_GROUPS, rhs, identifierMap)
					}
				case *REQ_IF_HEADER:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "COMMENT":
						inst.COMMENT = GongExtractString(rhs)
					case "CREATION_TIME":
						inst.CREATION_TIME = GongExtractString(rhs)
					case "REPOSITORY_ID":
						inst.REPOSITORY_ID = GongExtractString(rhs)
					case "REQ_IF_TOOL_ID":
						inst.REQ_IF_TOOL_ID = GongExtractString(rhs)
					case "REQ_IF_VERSION":
						inst.REQ_IF_VERSION = GongExtractString(rhs)
					case "SOURCE_TOOL_ID":
						inst.SOURCE_TOOL_ID = GongExtractString(rhs)
					case "TITLE":
						inst.TITLE = GongExtractString(rhs)
					}
				case *REQ_IF_TOOL_EXTENSION:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *SPECIFICATION:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					case "CHILDREN":
						__gong__assignPointer(&inst.CHILDREN, rhs, identifierMap)
					case "VALUES":
						__gong__assignPointer(&inst.VALUES, rhs, identifierMap)
					}
				case *SPECIFICATION_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsNodeExpanded":
						inst.IsNodeExpanded = GongExtractBool(rhs)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "IsWithHeadingNumbering":
						inst.IsWithHeadingNumbering = GongExtractBool(rhs)
					}
				case *SPECIFICATION_TYPE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "SPEC_ATTRIBUTES":
						__gong__assignPointer(&inst.SPEC_ATTRIBUTES, rhs, identifierMap)
					}
				case *SPEC_HIERARCHY:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "IS_EDITABLE":
						inst.IS_EDITABLE = GongExtractBool(rhs)
					case "IS_TABLE_INTERNAL":
						inst.IS_TABLE_INTERNAL = GongExtractBool(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "OBJECT":
						__gong__assignPointer(&inst.OBJECT, rhs, identifierMap)
					case "CHILDREN":
						__gong__assignPointer(&inst.CHILDREN, rhs, identifierMap)
					case "EDITABLE_ATTS":
						__gong__assignPointer(&inst.EDITABLE_ATTS, rhs, identifierMap)
					}
				case *SPEC_OBJECT:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "VALUES":
						__gong__assignPointer(&inst.VALUES, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *SPEC_OBJECT_TYPE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "SPEC_ATTRIBUTES":
						__gong__assignPointer(&inst.SPEC_ATTRIBUTES, rhs, identifierMap)
					}
				case *SPEC_OBJECT_TYPE_Rendering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsNodeExpanded":
						inst.IsNodeExpanded = GongExtractBool(rhs)
					case "ShowIdentifier":
						inst.ShowIdentifier = GongExtractBool(rhs)
					case "ShowName":
						inst.ShowName = GongExtractBool(rhs)
					case "ShowRelations":
						inst.ShowRelations = GongExtractBool(rhs)
					case "IsHeading":
						inst.IsHeading = GongExtractBool(rhs)
					}
				case *SPEC_RELATION:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "VALUES":
						__gong__assignPointer(&inst.VALUES, rhs, identifierMap)
					case "SOURCE":
						__gong__assignPointer(&inst.SOURCE, rhs, identifierMap)
					case "TARGET":
						__gong__assignPointer(&inst.TARGET, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
					}
				case *SPEC_RELATION_TYPE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DESC":
						inst.DESC = GongExtractString(rhs)
					case "IDENTIFIER":
						inst.IDENTIFIER = GongExtractString(rhs)
					case "LAST_CHANGE":
						inst.LAST_CHANGE = GongExtractString(rhs)
					case "LONG_NAME":
						inst.LONG_NAME = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						__gong__assignPointer(&inst.ALTERNATIVE_ID, rhs, identifierMap)
					case "SPEC_ATTRIBUTES":
						__gong__assignPointer(&inst.SPEC_ATTRIBUTES, rhs, identifierMap)
					}
				case *StaticWebSite:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MarkdownContent":
						inst.MarkdownContent = GongExtractString(rhs)
					case "Chapters":
						__gong__assignSliceOfPointers(&inst.Chapters, rhs, identifierMap)
					case "InputImagesDir":
						inst.InputImagesDir = GongExtractString(rhs)
					case "OutputStaticWebDir":
						inst.OutputStaticWebDir = GongExtractString(rhs)
					case "VersionInfo":
						inst.VersionInfo = GongExtractString(rhs)
					}
				case *StaticWebSiteChapter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MarkdownContent":
						inst.MarkdownContent = GongExtractString(rhs)
					case "Paragraphs":
						__gong__assignSliceOfPointers(&inst.Paragraphs, rhs, identifierMap)
					}
				case *StaticWebSiteGeneratedImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SourceDirectoryPath":
						inst.SourceDirectoryPath = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractInt(rhs)
					case "Height":
						inst.Height = GongExtractInt(rhs)
					}
				case *StaticWebSiteImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SourceDirectoryPath":
						inst.SourceDirectoryPath = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractInt(rhs)
					case "Height":
						inst.Height = GongExtractInt(rhs)
					}
				case *StaticWebSiteParagraph:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "LegendMarkdownContent":
						inst.LegendMarkdownContent = GongExtractString(rhs)
					case "Image":
						__gong__assignPointer(&inst.Image, rhs, identifierMap)
					}
				case *XHTML_CONTENT:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					case "PureText":
						inst.PureText = GongExtractString(rhs)
					}
							}
						}
					}
				}
			}
		case *ast.ExprStmt:
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "Unstage" {
						if ident, ok := sel.X.(*ast.Ident); ok {
							if instance, ok := identifierMap[ident.Name]; ok {
								if inst, ok := instance.(GongstructIF); ok {
									inst.UnstageVoid(stage)
								}
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
							if stage.IsInDeltaMode() && stage.GetNavigationMode() != GongNavigationModeNavigating {
								stage.Commit()
							} else {
								stage.ComputeInstancesNb()
								stage.ComputeReferenceAndOrders()
								if stage.OnInitCommitCallback != nil {
									stage.OnInitCommitCallback.BeforeCommit(stage)
								}
								if stage.OnInitCommitFromBackCallback != nil {
									stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
								}
								stage.RunBeforeCommitHooks()
								stage.RunAfterCommitHooks()
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}

// --- Generic Helpers for Unmarshallers (delegating to gong_runtime) ---

var GongExtractString = gong_runtime.ExtractString
var GongExtractInt = gong_runtime.ExtractInt
var GongExtractFloat = gong_runtime.ExtractFloat
var GongExtractBool = gong_runtime.ExtractBool
var GongExtractExpr = gong_runtime.ExtractExpr
var GongExtractDate = gong_runtime.ExtractDate

func __gong__extractMiddleUint(input string) (uint, error) {
	return gong_runtime.ExtractMiddleUint(input)
}

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok {
		funcName := ""
		var isSlices bool
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "slices" {
				isSlices = true
				funcName = sel.Sel.Name
			}
		} else if ident, ok := call.Fun.(*ast.Ident); ok {
			funcName = ident.Name
		}

		if isSlices {
			if funcName == "Delete" && len(call.Args) == 3 {
				start := GongExtractInt(call.Args[1])
				end := GongExtractInt(call.Args[2])
				if end <= len(*slice) {
					*slice = slices.Delete(*slice, start, end)
				}
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if index <= len(*slice) {
					if rIdent, ok := call.Args[2].(*ast.Ident); ok {
						if target, ok := identifierMap[rIdent.Name]; ok {
							if typedTarget, ok := target.(*T); ok {
								*slice = slices.Insert(*slice, index, typedTarget)
							}
						}
					}
				}
			}
		} else if funcName == "append" && len(call.Args) >= 2 {
			for i := 1; i < len(call.Args); i++ {
				if rIdent, ok := call.Args[i].(*ast.Ident); ok {
					if target, ok := identifierMap[rIdent.Name]; ok {
						if typedTarget, ok := target.(*T); ok {
							*slice = append(*slice, typedTarget)
						}
					}
				}
			}
		}
	}
}

// GongUnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func GongUnmarshallSliceOfPointers[T GongstructPtr](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) (err error) {
	return gong_runtime.UnmarshallSliceOfPointers(slice, valueExpr, identifierMap)
}

// GongUnmarshallPointer handles assignment of a single pointer field
func GongUnmarshallPointer[T GongstructPtr](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {
	gong_runtime.UnmarshallPointer(ptr, valueExpr, identifierMap)
}

// GongUnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func GongUnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr) {
	gong_runtime.UnmarshallEnum(ptr, valueExpr)
}

// GongInitialize initializes a staged instance, sets its name, and stages it
func GongInitialize[P interface {
	GongstructIF
	StagePreserveOrder(stage *Stage, order uint)
}](instance P, stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if newOrder, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

// insertion point per named struct
type ALTERNATIVE_IDUnmarshaller struct{}

func (u *ALTERNATIVE_IDUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ALTERNATIVE_ID), stage, identifier, instanceName, preserveOrder)
}

func (u *ALTERNATIVE_IDUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ALTERNATIVE_ID)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_BOOLEANUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_BOOLEANUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_BOOLEAN), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_BOOLEANUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_BOOLEAN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_BOOLEAN_RenderingUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_BOOLEAN_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_BOOLEAN_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_BOOLEAN_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowInTable":
		instance.ShowInTable = GongExtractBool(valueExpr)
	case "ShowInTitle":
		instance.ShowInTitle = GongExtractBool(valueExpr)
	case "ShowInSubject":
		instance.ShowInSubject = GongExtractBool(valueExpr)
	case "Rank":
		instance.Rank = GongExtractInt(valueExpr)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_DATEUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_DATEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_DATE), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_DATEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_DATE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_DATE_RenderingUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_DATE_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_DATE_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_DATE_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_DATE_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowInTable":
		instance.ShowInTable = GongExtractBool(valueExpr)
	case "ShowInTitle":
		instance.ShowInTitle = GongExtractBool(valueExpr)
	case "ShowInSubject":
		instance.ShowInSubject = GongExtractBool(valueExpr)
	case "Rank":
		instance.Rank = GongExtractInt(valueExpr)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_ENUMERATIONUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_ENUMERATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_ENUMERATION), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_ENUMERATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_ENUMERATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "MULTI_VALUED":
		instance.MULTI_VALUED = GongExtractBool(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_ENUMERATION_RenderingUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_ENUMERATION_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_ENUMERATION_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_ENUMERATION_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowInTable":
		instance.ShowInTable = GongExtractBool(valueExpr)
	case "ShowInTitle":
		instance.ShowInTitle = GongExtractBool(valueExpr)
	case "ShowInSubject":
		instance.ShowInSubject = GongExtractBool(valueExpr)
	case "Rank":
		instance.Rank = GongExtractInt(valueExpr)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_INTEGERUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_INTEGERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_INTEGER), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_INTEGERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_INTEGER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_INTEGER_RenderingUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_INTEGER_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_INTEGER_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_INTEGER_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_INTEGER_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowInTable":
		instance.ShowInTable = GongExtractBool(valueExpr)
	case "ShowInTitle":
		instance.ShowInTitle = GongExtractBool(valueExpr)
	case "ShowInSubject":
		instance.ShowInSubject = GongExtractBool(valueExpr)
	case "Rank":
		instance.Rank = GongExtractInt(valueExpr)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_REALUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_REALUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_REAL), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_REALUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_REAL)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_REAL_RenderingUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_REAL_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_REAL_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_REAL_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_REAL_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowInTable":
		instance.ShowInTable = GongExtractBool(valueExpr)
	case "ShowInTitle":
		instance.ShowInTitle = GongExtractBool(valueExpr)
	case "ShowInSubject":
		instance.ShowInSubject = GongExtractBool(valueExpr)
	case "Rank":
		instance.Rank = GongExtractInt(valueExpr)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_RenderingUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowInTable":
		instance.ShowInTable = GongExtractBool(valueExpr)
	case "ShowInTitle":
		instance.ShowInTitle = GongExtractBool(valueExpr)
	case "ShowInSubject":
		instance.ShowInSubject = GongExtractBool(valueExpr)
	case "Rank":
		instance.Rank = GongExtractInt(valueExpr)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_STRINGUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_STRINGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_STRING), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_STRINGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_STRING)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_STRING_RenderingUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_STRING_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_STRING_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_STRING_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_STRING_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowInTable":
		instance.ShowInTable = GongExtractBool(valueExpr)
	case "ShowInTitle":
		instance.ShowInTitle = GongExtractBool(valueExpr)
	case "ShowInSubject":
		instance.ShowInSubject = GongExtractBool(valueExpr)
	case "Rank":
		instance.Rank = GongExtractInt(valueExpr)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_XHTMLUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_XHTMLUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_XHTML), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_XHTMLUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_XHTML)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "DEFAULT_VALUE":
		GongUnmarshallPointer(&instance.DEFAULT_VALUE, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_DEFINITION_XHTML_RenderingUnmarshaller struct{}

func (u *ATTRIBUTE_DEFINITION_XHTML_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_DEFINITION_XHTML_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_DEFINITION_XHTML_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_DEFINITION_XHTML_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShowInTable":
		instance.ShowInTable = GongExtractBool(valueExpr)
	case "ShowInTitle":
		instance.ShowInTitle = GongExtractBool(valueExpr)
	case "ShowInSubject":
		instance.ShowInSubject = GongExtractBool(valueExpr)
	case "Rank":
		instance.Rank = GongExtractInt(valueExpr)
	}
	return nil
}

type ATTRIBUTE_VALUE_BOOLEANUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_BOOLEANUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_BOOLEAN), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_BOOLEANUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_BOOLEAN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractBool(valueExpr)
	}
	return nil
}

type ATTRIBUTE_VALUE_DATEUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_DATEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_DATE), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_DATEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_DATE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractString(valueExpr)
	}
	return nil
}

type ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_ENUMERATION), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_ENUMERATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	case "VALUES":
		GongUnmarshallPointer(&instance.VALUES, valueExpr, identifierMap)
	}
	return nil
}

type ATTRIBUTE_VALUE_INTEGERUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_INTEGERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_INTEGER), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_INTEGERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_INTEGER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractInt(valueExpr)
	}
	return nil
}

type ATTRIBUTE_VALUE_REALUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_REALUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_REAL), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_REALUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_REAL)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractFloat(valueExpr)
	}
	return nil
}

type ATTRIBUTE_VALUE_STRINGUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_STRINGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_STRING), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_STRINGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_STRING)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	case "THE_VALUE":
		instance.THE_VALUE = GongExtractString(valueExpr)
	}
	return nil
}

type ATTRIBUTE_VALUE_XHTMLUnmarshaller struct{}

func (u *ATTRIBUTE_VALUE_XHTMLUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ATTRIBUTE_VALUE_XHTML), stage, identifier, instanceName, preserveOrder)
}

func (u *ATTRIBUTE_VALUE_XHTMLUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ATTRIBUTE_VALUE_XHTML)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DEFINITION":
		GongUnmarshallPointer(&instance.DEFINITION, valueExpr, identifierMap)
	case "IS_SIMPLIFIED":
		instance.IS_SIMPLIFIED = GongExtractBool(valueExpr)
	case "THE_VALUE":
		GongUnmarshallPointer(&instance.THE_VALUE, valueExpr, identifierMap)
	case "THE_ORIGINAL_VALUE":
		GongUnmarshallPointer(&instance.THE_ORIGINAL_VALUE, valueExpr, identifierMap)
	}
	return nil
}

type A_ALTERNATIVE_IDUnmarshaller struct{}

func (u *A_ALTERNATIVE_IDUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ALTERNATIVE_ID), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ALTERNATIVE_IDUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ALTERNATIVE_ID)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_BOOLEAN_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_BOOLEAN_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
		instance.ATTRIBUTE_DEFINITION_BOOLEAN_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_DATE_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_DATE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_DATE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_DATE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_DATE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_DATE_REF":
		instance.ATTRIBUTE_DEFINITION_DATE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_ENUMERATION_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_ENUMERATION_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
		instance.ATTRIBUTE_DEFINITION_ENUMERATION_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_INTEGER_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_INTEGER_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_INTEGER_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_INTEGER_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_INTEGER_REF":
		instance.ATTRIBUTE_DEFINITION_INTEGER_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_REAL_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_REAL_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_REAL_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_REAL_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_REAL_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_REAL_REF":
		instance.ATTRIBUTE_DEFINITION_REAL_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_STRING_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_STRING_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_STRING_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_STRING_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_STRING_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_STRING_REF":
		instance.ATTRIBUTE_DEFINITION_STRING_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_DEFINITION_XHTML_REFUnmarshaller struct{}

func (u *A_ATTRIBUTE_DEFINITION_XHTML_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_DEFINITION_XHTML_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_DEFINITION_XHTML_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_XHTML_REF":
		instance.ATTRIBUTE_DEFINITION_XHTML_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_BOOLEANUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_BOOLEANUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_BOOLEAN), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_BOOLEANUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_BOOLEAN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_BOOLEAN":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_BOOLEAN, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_DATEUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_DATEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_DATE), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_DATEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_DATE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_DATE":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_DATE, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_ENUMERATION), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_ENUMERATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_ENUMERATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_ENUMERATION":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_ENUMERATION, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_INTEGERUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_INTEGERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_INTEGER), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_INTEGERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_INTEGER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_INTEGER":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_INTEGER, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_REALUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_REALUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_REAL), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_REALUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_REAL)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_REAL":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_REAL, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_STRINGUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_STRINGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_STRING), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_STRINGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_STRING)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_STRING":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_STRING, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_XHTMLUnmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_XHTMLUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_XHTML), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_XHTMLUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_XHTML)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_XHTML":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_XHTML, valueExpr, identifierMap)
	}
	return nil
}

type A_ATTRIBUTE_VALUE_XHTML_1Unmarshaller struct{}

func (u *A_ATTRIBUTE_VALUE_XHTML_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ATTRIBUTE_VALUE_XHTML_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ATTRIBUTE_VALUE_XHTML_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ATTRIBUTE_VALUE_XHTML_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_VALUE_BOOLEAN":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_BOOLEAN, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_DATE":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_DATE, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_ENUMERATION":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_ENUMERATION, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_INTEGER":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_INTEGER, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_REAL":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_REAL, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_STRING":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_STRING, valueExpr, identifierMap)
	case "ATTRIBUTE_VALUE_XHTML":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_VALUE_XHTML, valueExpr, identifierMap)
	}
	return nil
}

type A_CHILDRENUnmarshaller struct{}

func (u *A_CHILDRENUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_CHILDREN), stage, identifier, instanceName, preserveOrder)
}

func (u *A_CHILDRENUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_CHILDREN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_HIERARCHY":
		GongUnmarshallSliceOfPointers(&instance.SPEC_HIERARCHY, valueExpr, identifierMap)
	}
	return nil
}

type A_CORE_CONTENTUnmarshaller struct{}

func (u *A_CORE_CONTENTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_CORE_CONTENT), stage, identifier, instanceName, preserveOrder)
}

func (u *A_CORE_CONTENTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_CORE_CONTENT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "REQ_IF_CONTENT":
		GongUnmarshallPointer(&instance.REQ_IF_CONTENT, valueExpr, identifierMap)
	}
	return nil
}

type A_DATATYPESUnmarshaller struct{}

func (u *A_DATATYPESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_BOOLEAN":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_BOOLEAN, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_DATE":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_DATE, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_ENUMERATION":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_ENUMERATION, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_INTEGER":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_INTEGER, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_REAL":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_REAL, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_STRING":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_STRING, valueExpr, identifierMap)
	case "DATATYPE_DEFINITION_XHTML":
		GongUnmarshallSliceOfPointers(&instance.DATATYPE_DEFINITION_XHTML, valueExpr, identifierMap)
	}
	return nil
}

type A_DATATYPE_DEFINITION_BOOLEAN_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_BOOLEAN_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_BOOLEAN_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_BOOLEAN_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_BOOLEAN_REF":
		instance.DATATYPE_DEFINITION_BOOLEAN_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_DATE_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_DATE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_DATE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_DATE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_DATE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_DATE_REF":
		instance.DATATYPE_DEFINITION_DATE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_ENUMERATION_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_ENUMERATION_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_ENUMERATION_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_ENUMERATION_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_ENUMERATION_REF":
		instance.DATATYPE_DEFINITION_ENUMERATION_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_INTEGER_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_INTEGER_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_INTEGER_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_INTEGER_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_INTEGER_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_INTEGER_REF":
		instance.DATATYPE_DEFINITION_INTEGER_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_REAL_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_REAL_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_REAL_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_REAL_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_REAL_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_REAL_REF":
		instance.DATATYPE_DEFINITION_REAL_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_STRING_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_STRING_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_STRING_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_STRING_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_STRING_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_STRING_REF":
		instance.DATATYPE_DEFINITION_STRING_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_DATATYPE_DEFINITION_XHTML_REFUnmarshaller struct{}

func (u *A_DATATYPE_DEFINITION_XHTML_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_DATATYPE_DEFINITION_XHTML_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_DATATYPE_DEFINITION_XHTML_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_DATATYPE_DEFINITION_XHTML_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPE_DEFINITION_XHTML_REF":
		instance.DATATYPE_DEFINITION_XHTML_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_EDITABLE_ATTSUnmarshaller struct{}

func (u *A_EDITABLE_ATTSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_EDITABLE_ATTS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_EDITABLE_ATTSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_EDITABLE_ATTS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
		instance.ATTRIBUTE_DEFINITION_BOOLEAN_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_DATE_REF":
		instance.ATTRIBUTE_DEFINITION_DATE_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
		instance.ATTRIBUTE_DEFINITION_ENUMERATION_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_INTEGER_REF":
		instance.ATTRIBUTE_DEFINITION_INTEGER_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_REAL_REF":
		instance.ATTRIBUTE_DEFINITION_REAL_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_STRING_REF":
		instance.ATTRIBUTE_DEFINITION_STRING_REF = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_XHTML_REF":
		instance.ATTRIBUTE_DEFINITION_XHTML_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_ENUM_VALUE_REFUnmarshaller struct{}

func (u *A_ENUM_VALUE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_ENUM_VALUE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_ENUM_VALUE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_ENUM_VALUE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ENUM_VALUE_REF":
		instance.ENUM_VALUE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_OBJECTUnmarshaller struct{}

func (u *A_OBJECTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_OBJECT), stage, identifier, instanceName, preserveOrder)
}

func (u *A_OBJECTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_OBJECT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_OBJECT_REF":
		instance.SPEC_OBJECT_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_PROPERTIESUnmarshaller struct{}

func (u *A_PROPERTIESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_PROPERTIES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_PROPERTIESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_PROPERTIES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EMBEDDED_VALUE":
		GongUnmarshallPointer(&instance.EMBEDDED_VALUE, valueExpr, identifierMap)
	}
	return nil
}

type A_RELATION_GROUP_TYPE_REFUnmarshaller struct{}

func (u *A_RELATION_GROUP_TYPE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_RELATION_GROUP_TYPE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_RELATION_GROUP_TYPE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_RELATION_GROUP_TYPE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RELATION_GROUP_TYPE_REF":
		instance.RELATION_GROUP_TYPE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SOURCE_1Unmarshaller struct{}

func (u *A_SOURCE_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SOURCE_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SOURCE_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SOURCE_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_OBJECT_REF":
		instance.SPEC_OBJECT_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SOURCE_SPECIFICATION_1Unmarshaller struct{}

func (u *A_SOURCE_SPECIFICATION_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SOURCE_SPECIFICATION_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SOURCE_SPECIFICATION_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SOURCE_SPECIFICATION_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPECIFICATION_REF":
		GongUnmarshallEnum(&instance.SPECIFICATION_REF, valueExpr)
	}
	return nil
}

type A_SPECIFICATIONSUnmarshaller struct{}

func (u *A_SPECIFICATIONSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPECIFICATIONS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPECIFICATIONSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPECIFICATIONS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPECIFICATION":
		GongUnmarshallSliceOfPointers(&instance.SPECIFICATION, valueExpr, identifierMap)
	}
	return nil
}

type A_SPECIFICATION_TYPE_REFUnmarshaller struct{}

func (u *A_SPECIFICATION_TYPE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPECIFICATION_TYPE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPECIFICATION_TYPE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPECIFICATION_TYPE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPECIFICATION_TYPE_REF":
		instance.SPECIFICATION_TYPE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SPECIFIED_VALUESUnmarshaller struct{}

func (u *A_SPECIFIED_VALUESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPECIFIED_VALUES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPECIFIED_VALUESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPECIFIED_VALUES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ENUM_VALUE":
		GongUnmarshallSliceOfPointers(&instance.ENUM_VALUE, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_ATTRIBUTESUnmarshaller struct{}

func (u *A_SPEC_ATTRIBUTESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_ATTRIBUTES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_ATTRIBUTESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_ATTRIBUTES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ATTRIBUTE_DEFINITION_BOOLEAN":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_BOOLEAN, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_DATE":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_DATE, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_ENUMERATION":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_ENUMERATION, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_INTEGER":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_INTEGER, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_REAL":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_REAL, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_STRING":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_STRING, valueExpr, identifierMap)
	case "ATTRIBUTE_DEFINITION_XHTML":
		GongUnmarshallSliceOfPointers(&instance.ATTRIBUTE_DEFINITION_XHTML, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_OBJECTSUnmarshaller struct{}

func (u *A_SPEC_OBJECTSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_OBJECTS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_OBJECTSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_OBJECTS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_OBJECT":
		GongUnmarshallSliceOfPointers(&instance.SPEC_OBJECT, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_OBJECT_TYPE_REFUnmarshaller struct{}

func (u *A_SPEC_OBJECT_TYPE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_OBJECT_TYPE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_OBJECT_TYPE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_OBJECT_TYPE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_OBJECT_TYPE_REF":
		instance.SPEC_OBJECT_TYPE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SPEC_RELATIONSUnmarshaller struct{}

func (u *A_SPEC_RELATIONSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_RELATIONS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_RELATIONSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_RELATIONS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_RELATION":
		GongUnmarshallSliceOfPointers(&instance.SPEC_RELATION, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_RELATION_GROUPSUnmarshaller struct{}

func (u *A_SPEC_RELATION_GROUPSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_RELATION_GROUPS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_RELATION_GROUPSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_RELATION_GROUPS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RELATION_GROUP":
		GongUnmarshallSliceOfPointers(&instance.RELATION_GROUP, valueExpr, identifierMap)
	}
	return nil
}

type A_SPEC_RELATION_REFUnmarshaller struct{}

func (u *A_SPEC_RELATION_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_RELATION_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_RELATION_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_RELATION_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_RELATION_REF":
		instance.SPEC_RELATION_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SPEC_RELATION_TYPE_REFUnmarshaller struct{}

func (u *A_SPEC_RELATION_TYPE_REFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_RELATION_TYPE_REF), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_RELATION_TYPE_REFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_RELATION_TYPE_REF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SPEC_RELATION_TYPE_REF":
		instance.SPEC_RELATION_TYPE_REF = GongExtractString(valueExpr)
	}
	return nil
}

type A_SPEC_TYPESUnmarshaller struct{}

func (u *A_SPEC_TYPESUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_SPEC_TYPES), stage, identifier, instanceName, preserveOrder)
}

func (u *A_SPEC_TYPESUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_SPEC_TYPES)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RELATION_GROUP_TYPE":
		GongUnmarshallSliceOfPointers(&instance.RELATION_GROUP_TYPE, valueExpr, identifierMap)
	case "SPEC_OBJECT_TYPE":
		GongUnmarshallSliceOfPointers(&instance.SPEC_OBJECT_TYPE, valueExpr, identifierMap)
	case "SPEC_RELATION_TYPE":
		GongUnmarshallSliceOfPointers(&instance.SPEC_RELATION_TYPE, valueExpr, identifierMap)
	case "SPECIFICATION_TYPE":
		GongUnmarshallSliceOfPointers(&instance.SPECIFICATION_TYPE, valueExpr, identifierMap)
	}
	return nil
}

type A_THE_HEADERUnmarshaller struct{}

func (u *A_THE_HEADERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_THE_HEADER), stage, identifier, instanceName, preserveOrder)
}

func (u *A_THE_HEADERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_THE_HEADER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "REQ_IF_HEADER":
		GongUnmarshallPointer(&instance.REQ_IF_HEADER, valueExpr, identifierMap)
	}
	return nil
}

type A_TOOL_EXTENSIONSUnmarshaller struct{}

func (u *A_TOOL_EXTENSIONSUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_TOOL_EXTENSIONS), stage, identifier, instanceName, preserveOrder)
}

func (u *A_TOOL_EXTENSIONSUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_TOOL_EXTENSIONS)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "REQ_IF_TOOL_EXTENSION":
		GongUnmarshallSliceOfPointers(&instance.REQ_IF_TOOL_EXTENSION, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_BOOLEANUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_BOOLEANUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_BOOLEAN), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_BOOLEANUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_BOOLEAN)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_DATEUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_DATEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_DATE), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_DATEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_DATE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_ENUMERATIONUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_ENUMERATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_ENUMERATION), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_ENUMERATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_ENUMERATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPECIFIED_VALUES":
		GongUnmarshallPointer(&instance.SPECIFIED_VALUES, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_INTEGERUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_INTEGERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_INTEGER), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_INTEGERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_INTEGER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "MAX":
		instance.MAX = GongExtractInt(valueExpr)
	case "MIN":
		instance.MIN = GongExtractInt(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_REALUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_REALUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_REAL), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_REALUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_REAL)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ACCURACY":
		instance.ACCURACY = GongExtractInt(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "MAX":
		instance.MAX = GongExtractFloat(valueExpr)
	case "MIN":
		instance.MIN = GongExtractFloat(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_STRINGUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_STRINGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_STRING), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_STRINGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_STRING)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "MAX_LENGTH":
		instance.MAX_LENGTH = GongExtractInt(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type DATATYPE_DEFINITION_XHTMLUnmarshaller struct{}

func (u *DATATYPE_DEFINITION_XHTMLUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DATATYPE_DEFINITION_XHTML), stage, identifier, instanceName, preserveOrder)
}

func (u *DATATYPE_DEFINITION_XHTMLUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DATATYPE_DEFINITION_XHTML)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	}
	return nil
}

type EMBEDDED_VALUEUnmarshaller struct{}

func (u *EMBEDDED_VALUEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EMBEDDED_VALUE), stage, identifier, instanceName, preserveOrder)
}

func (u *EMBEDDED_VALUEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EMBEDDED_VALUE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "KEY":
		instance.KEY = GongExtractInt(valueExpr)
	case "OTHER_CONTENT":
		instance.OTHER_CONTENT = GongExtractString(valueExpr)
	}
	return nil
}

type ENUM_VALUEUnmarshaller struct{}

func (u *ENUM_VALUEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ENUM_VALUE), stage, identifier, instanceName, preserveOrder)
}

func (u *ENUM_VALUEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ENUM_VALUE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "PROPERTIES":
		GongUnmarshallPointer(&instance.PROPERTIES, valueExpr, identifierMap)
	}
	return nil
}

type EmbeddedJpgImageUnmarshaller struct{}

func (u *EmbeddedJpgImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmbeddedJpgImage), stage, identifier, instanceName, preserveOrder)
}

func (u *EmbeddedJpgImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmbeddedJpgImage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Base64Content":
		instance.Base64Content = GongExtractString(valueExpr)
	}
	return nil
}

type EmbeddedPngImageUnmarshaller struct{}

func (u *EmbeddedPngImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmbeddedPngImage), stage, identifier, instanceName, preserveOrder)
}

func (u *EmbeddedPngImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmbeddedPngImage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Base64Content":
		instance.Base64Content = GongExtractString(valueExpr)
	}
	return nil
}

type EmbeddedSvgImageUnmarshaller struct{}

func (u *EmbeddedSvgImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EmbeddedSvgImage), stage, identifier, instanceName, preserveOrder)
}

func (u *EmbeddedSvgImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EmbeddedSvgImage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	}
	return nil
}

type KillUnmarshaller struct{}

func (u *KillUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Kill), stage, identifier, instanceName, preserveOrder)
}

func (u *KillUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Kill)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Map_identifier_boolUnmarshaller struct{}

func (u *Map_identifier_boolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Map_identifier_bool), stage, identifier, instanceName, preserveOrder)
}

func (u *Map_identifier_boolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Map_identifier_bool)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractBool(valueExpr)
	}
	return nil
}

type RELATION_GROUPUnmarshaller struct{}

func (u *RELATION_GROUPUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RELATION_GROUP), stage, identifier, instanceName, preserveOrder)
}

func (u *RELATION_GROUPUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RELATION_GROUP)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SOURCE_SPECIFICATION":
		GongUnmarshallPointer(&instance.SOURCE_SPECIFICATION, valueExpr, identifierMap)
	case "SPEC_RELATIONS":
		GongUnmarshallPointer(&instance.SPEC_RELATIONS, valueExpr, identifierMap)
	case "TARGET_SPECIFICATION":
		GongUnmarshallPointer(&instance.TARGET_SPECIFICATION, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type RELATION_GROUP_TYPEUnmarshaller struct{}

func (u *RELATION_GROUP_TYPEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RELATION_GROUP_TYPE), stage, identifier, instanceName, preserveOrder)
}

func (u *RELATION_GROUP_TYPEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RELATION_GROUP_TYPE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPEC_ATTRIBUTES":
		GongUnmarshallPointer(&instance.SPEC_ATTRIBUTES, valueExpr, identifierMap)
	}
	return nil
}

type REQ_IFUnmarshaller struct{}

func (u *REQ_IFUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(REQ_IF), stage, identifier, instanceName, preserveOrder)
}

func (u *REQ_IFUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*REQ_IF)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "THE_HEADER":
		GongUnmarshallPointer(&instance.THE_HEADER, valueExpr, identifierMap)
	case "CORE_CONTENT":
		GongUnmarshallPointer(&instance.CORE_CONTENT, valueExpr, identifierMap)
	case "TOOL_EXTENSIONS":
		GongUnmarshallPointer(&instance.TOOL_EXTENSIONS, valueExpr, identifierMap)
	}
	return nil
}

type REQ_IF_CONTENTUnmarshaller struct{}

func (u *REQ_IF_CONTENTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(REQ_IF_CONTENT), stage, identifier, instanceName, preserveOrder)
}

func (u *REQ_IF_CONTENTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*REQ_IF_CONTENT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DATATYPES":
		GongUnmarshallPointer(&instance.DATATYPES, valueExpr, identifierMap)
	case "SPEC_TYPES":
		GongUnmarshallPointer(&instance.SPEC_TYPES, valueExpr, identifierMap)
	case "SPEC_OBJECTS":
		GongUnmarshallPointer(&instance.SPEC_OBJECTS, valueExpr, identifierMap)
	case "SPEC_RELATIONS":
		GongUnmarshallPointer(&instance.SPEC_RELATIONS, valueExpr, identifierMap)
	case "SPECIFICATIONS":
		GongUnmarshallPointer(&instance.SPECIFICATIONS, valueExpr, identifierMap)
	case "SPEC_RELATION_GROUPS":
		GongUnmarshallPointer(&instance.SPEC_RELATION_GROUPS, valueExpr, identifierMap)
	}
	return nil
}

type REQ_IF_HEADERUnmarshaller struct{}

func (u *REQ_IF_HEADERUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(REQ_IF_HEADER), stage, identifier, instanceName, preserveOrder)
}

func (u *REQ_IF_HEADERUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*REQ_IF_HEADER)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "COMMENT":
		instance.COMMENT = GongExtractString(valueExpr)
	case "CREATION_TIME":
		instance.CREATION_TIME = GongExtractString(valueExpr)
	case "REPOSITORY_ID":
		instance.REPOSITORY_ID = GongExtractString(valueExpr)
	case "REQ_IF_TOOL_ID":
		instance.REQ_IF_TOOL_ID = GongExtractString(valueExpr)
	case "REQ_IF_VERSION":
		instance.REQ_IF_VERSION = GongExtractString(valueExpr)
	case "SOURCE_TOOL_ID":
		instance.SOURCE_TOOL_ID = GongExtractString(valueExpr)
	case "TITLE":
		instance.TITLE = GongExtractString(valueExpr)
	}
	return nil
}

type REQ_IF_TOOL_EXTENSIONUnmarshaller struct{}

func (u *REQ_IF_TOOL_EXTENSIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(REQ_IF_TOOL_EXTENSION), stage, identifier, instanceName, preserveOrder)
}

func (u *REQ_IF_TOOL_EXTENSIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*REQ_IF_TOOL_EXTENSION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SPECIFICATIONUnmarshaller struct{}

func (u *SPECIFICATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPECIFICATION), stage, identifier, instanceName, preserveOrder)
}

func (u *SPECIFICATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPECIFICATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	case "CHILDREN":
		GongUnmarshallPointer(&instance.CHILDREN, valueExpr, identifierMap)
	case "VALUES":
		GongUnmarshallPointer(&instance.VALUES, valueExpr, identifierMap)
	}
	return nil
}

type SPECIFICATION_RenderingUnmarshaller struct{}

func (u *SPECIFICATION_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPECIFICATION_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *SPECIFICATION_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPECIFICATION_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsNodeExpanded":
		instance.IsNodeExpanded = GongExtractBool(valueExpr)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	case "IsWithHeadingNumbering":
		instance.IsWithHeadingNumbering = GongExtractBool(valueExpr)
	}
	return nil
}

type SPECIFICATION_TYPEUnmarshaller struct{}

func (u *SPECIFICATION_TYPEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPECIFICATION_TYPE), stage, identifier, instanceName, preserveOrder)
}

func (u *SPECIFICATION_TYPEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPECIFICATION_TYPE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPEC_ATTRIBUTES":
		GongUnmarshallPointer(&instance.SPEC_ATTRIBUTES, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_HIERARCHYUnmarshaller struct{}

func (u *SPEC_HIERARCHYUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_HIERARCHY), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_HIERARCHYUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_HIERARCHY)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "IS_EDITABLE":
		instance.IS_EDITABLE = GongExtractBool(valueExpr)
	case "IS_TABLE_INTERNAL":
		instance.IS_TABLE_INTERNAL = GongExtractBool(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "OBJECT":
		GongUnmarshallPointer(&instance.OBJECT, valueExpr, identifierMap)
	case "CHILDREN":
		GongUnmarshallPointer(&instance.CHILDREN, valueExpr, identifierMap)
	case "EDITABLE_ATTS":
		GongUnmarshallPointer(&instance.EDITABLE_ATTS, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_OBJECTUnmarshaller struct{}

func (u *SPEC_OBJECTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_OBJECT), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_OBJECTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_OBJECT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "VALUES":
		GongUnmarshallPointer(&instance.VALUES, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_OBJECT_TYPEUnmarshaller struct{}

func (u *SPEC_OBJECT_TYPEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_OBJECT_TYPE), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_OBJECT_TYPEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_OBJECT_TYPE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPEC_ATTRIBUTES":
		GongUnmarshallPointer(&instance.SPEC_ATTRIBUTES, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_OBJECT_TYPE_RenderingUnmarshaller struct{}

func (u *SPEC_OBJECT_TYPE_RenderingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_OBJECT_TYPE_Rendering), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_OBJECT_TYPE_RenderingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_OBJECT_TYPE_Rendering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsNodeExpanded":
		instance.IsNodeExpanded = GongExtractBool(valueExpr)
	case "ShowIdentifier":
		instance.ShowIdentifier = GongExtractBool(valueExpr)
	case "ShowName":
		instance.ShowName = GongExtractBool(valueExpr)
	case "ShowRelations":
		instance.ShowRelations = GongExtractBool(valueExpr)
	case "IsHeading":
		instance.IsHeading = GongExtractBool(valueExpr)
	}
	return nil
}

type SPEC_RELATIONUnmarshaller struct{}

func (u *SPEC_RELATIONUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_RELATION), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_RELATIONUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_RELATION)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "VALUES":
		GongUnmarshallPointer(&instance.VALUES, valueExpr, identifierMap)
	case "SOURCE":
		GongUnmarshallPointer(&instance.SOURCE, valueExpr, identifierMap)
	case "TARGET":
		GongUnmarshallPointer(&instance.TARGET, valueExpr, identifierMap)
	case "TYPE":
		GongUnmarshallPointer(&instance.TYPE, valueExpr, identifierMap)
	}
	return nil
}

type SPEC_RELATION_TYPEUnmarshaller struct{}

func (u *SPEC_RELATION_TYPEUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SPEC_RELATION_TYPE), stage, identifier, instanceName, preserveOrder)
}

func (u *SPEC_RELATION_TYPEUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SPEC_RELATION_TYPE)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DESC":
		instance.DESC = GongExtractString(valueExpr)
	case "IDENTIFIER":
		instance.IDENTIFIER = GongExtractString(valueExpr)
	case "LAST_CHANGE":
		instance.LAST_CHANGE = GongExtractString(valueExpr)
	case "LONG_NAME":
		instance.LONG_NAME = GongExtractString(valueExpr)
	case "ALTERNATIVE_ID":
		GongUnmarshallPointer(&instance.ALTERNATIVE_ID, valueExpr, identifierMap)
	case "SPEC_ATTRIBUTES":
		GongUnmarshallPointer(&instance.SPEC_ATTRIBUTES, valueExpr, identifierMap)
	}
	return nil
}

type StaticWebSiteUnmarshaller struct{}

func (u *StaticWebSiteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaticWebSite), stage, identifier, instanceName, preserveOrder)
}

func (u *StaticWebSiteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaticWebSite)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MarkdownContent":
		instance.MarkdownContent = GongExtractString(valueExpr)
	case "Chapters":
		GongUnmarshallSliceOfPointers(&instance.Chapters, valueExpr, identifierMap)
	case "InputImagesDir":
		instance.InputImagesDir = GongExtractString(valueExpr)
	case "OutputStaticWebDir":
		instance.OutputStaticWebDir = GongExtractString(valueExpr)
	case "VersionInfo":
		instance.VersionInfo = GongExtractString(valueExpr)
	}
	return nil
}

type StaticWebSiteChapterUnmarshaller struct{}

func (u *StaticWebSiteChapterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaticWebSiteChapter), stage, identifier, instanceName, preserveOrder)
}

func (u *StaticWebSiteChapterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaticWebSiteChapter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MarkdownContent":
		instance.MarkdownContent = GongExtractString(valueExpr)
	case "Paragraphs":
		GongUnmarshallSliceOfPointers(&instance.Paragraphs, valueExpr, identifierMap)
	}
	return nil
}

type StaticWebSiteGeneratedImageUnmarshaller struct{}

func (u *StaticWebSiteGeneratedImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaticWebSiteGeneratedImage), stage, identifier, instanceName, preserveOrder)
}

func (u *StaticWebSiteGeneratedImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaticWebSiteGeneratedImage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SourceDirectoryPath":
		instance.SourceDirectoryPath = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractInt(valueExpr)
	case "Height":
		instance.Height = GongExtractInt(valueExpr)
	}
	return nil
}

type StaticWebSiteImageUnmarshaller struct{}

func (u *StaticWebSiteImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaticWebSiteImage), stage, identifier, instanceName, preserveOrder)
}

func (u *StaticWebSiteImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaticWebSiteImage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SourceDirectoryPath":
		instance.SourceDirectoryPath = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractInt(valueExpr)
	case "Height":
		instance.Height = GongExtractInt(valueExpr)
	}
	return nil
}

type StaticWebSiteParagraphUnmarshaller struct{}

func (u *StaticWebSiteParagraphUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StaticWebSiteParagraph), stage, identifier, instanceName, preserveOrder)
}

func (u *StaticWebSiteParagraphUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StaticWebSiteParagraph)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "LegendMarkdownContent":
		instance.LegendMarkdownContent = GongExtractString(valueExpr)
	case "Image":
		GongUnmarshallPointer(&instance.Image, valueExpr, identifierMap)
	}
	return nil
}

type XHTML_CONTENTUnmarshaller struct{}

func (u *XHTML_CONTENTUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(XHTML_CONTENT), stage, identifier, instanceName, preserveOrder)
}

func (u *XHTML_CONTENTUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*XHTML_CONTENT)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	case "PureText":
		instance.PureText = GongExtractString(valueExpr)
	}
	return nil
}
