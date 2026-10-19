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
		case "github.com/fullstack-lang/gong/app/xsd/go/models":
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
				case "All":
					identifierMap[ident.Name] = __gong__stageSetInit(new(All), stage, ident.Name, instanceName, preserveOrder)
				case "Annotation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Annotation), stage, ident.Name, instanceName, preserveOrder)
				case "Attribute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Attribute), stage, ident.Name, instanceName, preserveOrder)
				case "AttributeGroup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AttributeGroup), stage, ident.Name, instanceName, preserveOrder)
				case "Choice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Choice), stage, ident.Name, instanceName, preserveOrder)
				case "ComplexContent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ComplexContent), stage, ident.Name, instanceName, preserveOrder)
				case "ComplexType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ComplexType), stage, ident.Name, instanceName, preserveOrder)
				case "Documentation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Documentation), stage, ident.Name, instanceName, preserveOrder)
				case "Element":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Element), stage, ident.Name, instanceName, preserveOrder)
				case "Enumeration":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Enumeration), stage, ident.Name, instanceName, preserveOrder)
				case "Extension":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Extension), stage, ident.Name, instanceName, preserveOrder)
				case "Group":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group), stage, ident.Name, instanceName, preserveOrder)
				case "Length":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Length), stage, ident.Name, instanceName, preserveOrder)
				case "MaxInclusive":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MaxInclusive), stage, ident.Name, instanceName, preserveOrder)
				case "MaxLength":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MaxLength), stage, ident.Name, instanceName, preserveOrder)
				case "MinInclusive":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MinInclusive), stage, ident.Name, instanceName, preserveOrder)
				case "MinLength":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MinLength), stage, ident.Name, instanceName, preserveOrder)
				case "Pattern":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pattern), stage, ident.Name, instanceName, preserveOrder)
				case "Restriction":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Restriction), stage, ident.Name, instanceName, preserveOrder)
				case "Schema":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Schema), stage, ident.Name, instanceName, preserveOrder)
				case "Sequence":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Sequence), stage, ident.Name, instanceName, preserveOrder)
				case "SimpleContent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SimpleContent), stage, ident.Name, instanceName, preserveOrder)
				case "SimpleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SimpleType), stage, ident.Name, instanceName, preserveOrder)
				case "TotalDigit":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TotalDigit), stage, ident.Name, instanceName, preserveOrder)
				case "Union":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Union), stage, ident.Name, instanceName, preserveOrder)
				case "WhiteSpace":
					identifierMap[ident.Name] = __gong__stageSetInit(new(WhiteSpace), stage, ident.Name, instanceName, preserveOrder)
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
				case *All:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					}
				case *Annotation:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Documentations":
						__gong__assignSliceOfPointers(&inst.Documentations, rhs, identifierMap)
					}
				case *Attribute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "Default":
						inst.Default = GongExtractString(rhs)
					case "Use":
						inst.Use = GongExtractString(rhs)
					case "Form":
						inst.Form = GongExtractString(rhs)
					case "Fixed":
						inst.Fixed = GongExtractString(rhs)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "TargetNamespace":
						inst.TargetNamespace = GongExtractString(rhs)
					case "SimpleType":
						inst.SimpleType = GongExtractString(rhs)
					case "IDXSD":
						inst.IDXSD = GongExtractString(rhs)
					}
				case *AttributeGroup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "AttributeGroups":
						__gong__assignSliceOfPointers(&inst.AttributeGroups, rhs, identifierMap)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					}
				case *Choice:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					case "IsDuplicatedInXSD":
						inst.IsDuplicatedInXSD = GongExtractBool(rhs)
					}
				case *ComplexContent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *ComplexType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "IsAnonymous":
						inst.IsAnonymous = GongExtractBool(rhs)
					case "OuterElement":
						__gong__assignPointer(&inst.OuterElement, rhs, identifierMap)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					case "Extension":
						__gong__assignPointer(&inst.Extension, rhs, identifierMap)
					case "SimpleContent":
						__gong__assignPointer(&inst.SimpleContent, rhs, identifierMap)
					case "ComplexContent":
						__gong__assignPointer(&inst.ComplexContent, rhs, identifierMap)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "AttributeGroups":
						__gong__assignSliceOfPointers(&inst.AttributeGroups, rhs, identifierMap)
					case "IsDuplicatedInXSD":
						inst.IsDuplicatedInXSD = GongExtractBool(rhs)
					}
				case *Documentation:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Source":
						inst.Source = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					}
				case *Element:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					case "Default":
						inst.Default = GongExtractString(rhs)
					case "Fixed":
						inst.Fixed = GongExtractString(rhs)
					case "Nillable":
						inst.Nillable = GongExtractString(rhs)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "Abstract":
						inst.Abstract = GongExtractString(rhs)
					case "Form":
						inst.Form = GongExtractString(rhs)
					case "Block":
						inst.Block = GongExtractString(rhs)
					case "Final":
						inst.Final = GongExtractString(rhs)
					case "SimpleType":
						__gong__assignPointer(&inst.SimpleType, rhs, identifierMap)
					case "ComplexType":
						__gong__assignPointer(&inst.ComplexType, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "IsDuplicatedInXSD":
						inst.IsDuplicatedInXSD = GongExtractBool(rhs)
					}
				case *Enumeration:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Extension:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					case "Base":
						inst.Base = GongExtractString(rhs)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "AttributeGroups":
						__gong__assignSliceOfPointers(&inst.AttributeGroups, rhs, identifierMap)
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "IsAnonymous":
						inst.IsAnonymous = GongExtractBool(rhs)
					case "OuterElement":
						__gong__assignPointer(&inst.OuterElement, rhs, identifierMap)
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					}
				case *Length:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MaxInclusive:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MaxLength:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MinInclusive:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MinLength:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Pattern:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Restriction:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Base":
						inst.Base = GongExtractString(rhs)
					case "Enumerations":
						__gong__assignSliceOfPointers(&inst.Enumerations, rhs, identifierMap)
					case "MinInclusive":
						__gong__assignPointer(&inst.MinInclusive, rhs, identifierMap)
					case "MaxInclusive":
						__gong__assignPointer(&inst.MaxInclusive, rhs, identifierMap)
					case "Pattern":
						__gong__assignPointer(&inst.Pattern, rhs, identifierMap)
					case "WhiteSpace":
						__gong__assignPointer(&inst.WhiteSpace, rhs, identifierMap)
					case "MinLength":
						__gong__assignPointer(&inst.MinLength, rhs, identifierMap)
					case "MaxLength":
						__gong__assignPointer(&inst.MaxLength, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "TotalDigit":
						__gong__assignPointer(&inst.TotalDigit, rhs, identifierMap)
					}
				case *Schema:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Xs":
						inst.Xs = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
					case "SimpleTypes":
						__gong__assignSliceOfPointers(&inst.SimpleTypes, rhs, identifierMap)
					case "ComplexTypes":
						__gong__assignSliceOfPointers(&inst.ComplexTypes, rhs, identifierMap)
					case "AttributeGroups":
						__gong__assignSliceOfPointers(&inst.AttributeGroups, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					}
				case *Sequence:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					}
				case *SimpleContent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Extension":
						__gong__assignPointer(&inst.Extension, rhs, identifierMap)
					case "Restriction":
						__gong__assignPointer(&inst.Restriction, rhs, identifierMap)
					}
				case *SimpleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Restriction":
						__gong__assignPointer(&inst.Restriction, rhs, identifierMap)
					case "Union":
						__gong__assignPointer(&inst.Union, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					}
				case *TotalDigit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Union:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "MemberTypes":
						inst.MemberTypes = GongExtractString(rhs)
					}
				case *WhiteSpace:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
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
type AllUnmarshaller struct{}

func (u *AllUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(All), stage, identifier, instanceName, preserveOrder)
}

func (u *AllUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*All)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	}
	return nil
}

type AnnotationUnmarshaller struct{}

func (u *AnnotationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Annotation), stage, identifier, instanceName, preserveOrder)
}

func (u *AnnotationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Annotation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Documentations":
		GongUnmarshallSliceOfPointers(&instance.Documentations, valueExpr, identifierMap)
	}
	return nil
}

type AttributeUnmarshaller struct{}

func (u *AttributeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Attribute), stage, identifier, instanceName, preserveOrder)
}

func (u *AttributeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Attribute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "Default":
		instance.Default = GongExtractString(valueExpr)
	case "Use":
		instance.Use = GongExtractString(valueExpr)
	case "Form":
		instance.Form = GongExtractString(valueExpr)
	case "Fixed":
		instance.Fixed = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "TargetNamespace":
		instance.TargetNamespace = GongExtractString(valueExpr)
	case "SimpleType":
		instance.SimpleType = GongExtractString(valueExpr)
	case "IDXSD":
		instance.IDXSD = GongExtractString(valueExpr)
	}
	return nil
}

type AttributeGroupUnmarshaller struct{}

func (u *AttributeGroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AttributeGroup), stage, identifier, instanceName, preserveOrder)
}

func (u *AttributeGroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AttributeGroup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	}
	return nil
}

type ChoiceUnmarshaller struct{}

func (u *ChoiceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Choice), stage, identifier, instanceName, preserveOrder)
}

func (u *ChoiceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Choice)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "IsDuplicatedInXSD":
		instance.IsDuplicatedInXSD = GongExtractBool(valueExpr)
	}
	return nil
}

type ComplexContentUnmarshaller struct{}

func (u *ComplexContentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ComplexContent), stage, identifier, instanceName, preserveOrder)
}

func (u *ComplexContentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ComplexContent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ComplexTypeUnmarshaller struct{}

func (u *ComplexTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ComplexType), stage, identifier, instanceName, preserveOrder)
}

func (u *ComplexTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ComplexType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "IsAnonymous":
		instance.IsAnonymous = GongExtractBool(valueExpr)
	case "OuterElement":
		GongUnmarshallPointer(&instance.OuterElement, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "Extension":
		GongUnmarshallPointer(&instance.Extension, valueExpr, identifierMap)
	case "SimpleContent":
		GongUnmarshallPointer(&instance.SimpleContent, valueExpr, identifierMap)
	case "ComplexContent":
		GongUnmarshallPointer(&instance.ComplexContent, valueExpr, identifierMap)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	case "IsDuplicatedInXSD":
		instance.IsDuplicatedInXSD = GongExtractBool(valueExpr)
	}
	return nil
}

type DocumentationUnmarshaller struct{}

func (u *DocumentationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Documentation), stage, identifier, instanceName, preserveOrder)
}

func (u *DocumentationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Documentation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Source":
		instance.Source = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	}
	return nil
}

type ElementUnmarshaller struct{}

func (u *ElementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Element), stage, identifier, instanceName, preserveOrder)
}

func (u *ElementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Element)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Type":
		instance.Type = GongExtractString(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "Default":
		instance.Default = GongExtractString(valueExpr)
	case "Fixed":
		instance.Fixed = GongExtractString(valueExpr)
	case "Nillable":
		instance.Nillable = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "Abstract":
		instance.Abstract = GongExtractString(valueExpr)
	case "Form":
		instance.Form = GongExtractString(valueExpr)
	case "Block":
		instance.Block = GongExtractString(valueExpr)
	case "Final":
		instance.Final = GongExtractString(valueExpr)
	case "SimpleType":
		GongUnmarshallPointer(&instance.SimpleType, valueExpr, identifierMap)
	case "ComplexType":
		GongUnmarshallPointer(&instance.ComplexType, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "IsDuplicatedInXSD":
		instance.IsDuplicatedInXSD = GongExtractBool(valueExpr)
	}
	return nil
}

type EnumerationUnmarshaller struct{}

func (u *EnumerationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Enumeration), stage, identifier, instanceName, preserveOrder)
}

func (u *EnumerationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Enumeration)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type ExtensionUnmarshaller struct{}

func (u *ExtensionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Extension), stage, identifier, instanceName, preserveOrder)
}

func (u *ExtensionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Extension)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	case "Base":
		instance.Base = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	}
	return nil
}

type GroupUnmarshaller struct{}

func (u *GroupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Group), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Ref":
		instance.Ref = GongExtractString(valueExpr)
	case "IsAnonymous":
		instance.IsAnonymous = GongExtractBool(valueExpr)
	case "OuterElement":
		GongUnmarshallPointer(&instance.OuterElement, valueExpr, identifierMap)
	case "HasNameConflict":
		instance.HasNameConflict = GongExtractBool(valueExpr)
	case "GoIdentifier":
		instance.GoIdentifier = GongExtractString(valueExpr)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	}
	return nil
}

type LengthUnmarshaller struct{}

func (u *LengthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Length), stage, identifier, instanceName, preserveOrder)
}

func (u *LengthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Length)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MaxInclusiveUnmarshaller struct{}

func (u *MaxInclusiveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MaxInclusive), stage, identifier, instanceName, preserveOrder)
}

func (u *MaxInclusiveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MaxInclusive)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MaxLengthUnmarshaller struct{}

func (u *MaxLengthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MaxLength), stage, identifier, instanceName, preserveOrder)
}

func (u *MaxLengthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MaxLength)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MinInclusiveUnmarshaller struct{}

func (u *MinInclusiveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MinInclusive), stage, identifier, instanceName, preserveOrder)
}

func (u *MinInclusiveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MinInclusive)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type MinLengthUnmarshaller struct{}

func (u *MinLengthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MinLength), stage, identifier, instanceName, preserveOrder)
}

func (u *MinLengthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MinLength)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type PatternUnmarshaller struct{}

func (u *PatternUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Pattern), stage, identifier, instanceName, preserveOrder)
}

func (u *PatternUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pattern)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type RestrictionUnmarshaller struct{}

func (u *RestrictionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Restriction), stage, identifier, instanceName, preserveOrder)
}

func (u *RestrictionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Restriction)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Base":
		instance.Base = GongExtractString(valueExpr)
	case "Enumerations":
		GongUnmarshallSliceOfPointers(&instance.Enumerations, valueExpr, identifierMap)
	case "MinInclusive":
		GongUnmarshallPointer(&instance.MinInclusive, valueExpr, identifierMap)
	case "MaxInclusive":
		GongUnmarshallPointer(&instance.MaxInclusive, valueExpr, identifierMap)
	case "Pattern":
		GongUnmarshallPointer(&instance.Pattern, valueExpr, identifierMap)
	case "WhiteSpace":
		GongUnmarshallPointer(&instance.WhiteSpace, valueExpr, identifierMap)
	case "MinLength":
		GongUnmarshallPointer(&instance.MinLength, valueExpr, identifierMap)
	case "MaxLength":
		GongUnmarshallPointer(&instance.MaxLength, valueExpr, identifierMap)
	case "Length":
		GongUnmarshallPointer(&instance.Length, valueExpr, identifierMap)
	case "TotalDigit":
		GongUnmarshallPointer(&instance.TotalDigit, valueExpr, identifierMap)
	}
	return nil
}

type SchemaUnmarshaller struct{}

func (u *SchemaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Schema), stage, identifier, instanceName, preserveOrder)
}

func (u *SchemaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Schema)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Xs":
		instance.Xs = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "SimpleTypes":
		GongUnmarshallSliceOfPointers(&instance.SimpleTypes, valueExpr, identifierMap)
	case "ComplexTypes":
		GongUnmarshallSliceOfPointers(&instance.ComplexTypes, valueExpr, identifierMap)
	case "AttributeGroups":
		GongUnmarshallSliceOfPointers(&instance.AttributeGroups, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	}
	return nil
}

type SequenceUnmarshaller struct{}

func (u *SequenceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Sequence), stage, identifier, instanceName, preserveOrder)
}

func (u *SequenceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Sequence)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "OuterElementName":
		instance.OuterElementName = GongExtractString(valueExpr)
	case "Sequences":
		GongUnmarshallSliceOfPointers(&instance.Sequences, valueExpr, identifierMap)
	case "Alls":
		GongUnmarshallSliceOfPointers(&instance.Alls, valueExpr, identifierMap)
	case "Choices":
		GongUnmarshallSliceOfPointers(&instance.Choices, valueExpr, identifierMap)
	case "Groups":
		GongUnmarshallSliceOfPointers(&instance.Groups, valueExpr, identifierMap)
	case "Elements":
		GongUnmarshallSliceOfPointers(&instance.Elements, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	case "MinOccurs":
		instance.MinOccurs = GongExtractString(valueExpr)
	case "MaxOccurs":
		instance.MaxOccurs = GongExtractString(valueExpr)
	}
	return nil
}

type SimpleContentUnmarshaller struct{}

func (u *SimpleContentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SimpleContent), stage, identifier, instanceName, preserveOrder)
}

func (u *SimpleContentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SimpleContent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Extension":
		GongUnmarshallPointer(&instance.Extension, valueExpr, identifierMap)
	case "Restriction":
		GongUnmarshallPointer(&instance.Restriction, valueExpr, identifierMap)
	}
	return nil
}

type SimpleTypeUnmarshaller struct{}

func (u *SimpleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SimpleType), stage, identifier, instanceName, preserveOrder)
}

func (u *SimpleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SimpleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Restriction":
		GongUnmarshallPointer(&instance.Restriction, valueExpr, identifierMap)
	case "Union":
		GongUnmarshallPointer(&instance.Union, valueExpr, identifierMap)
	case "Order":
		instance.Order = GongExtractInt(valueExpr)
	case "Depth":
		instance.Depth = GongExtractInt(valueExpr)
	}
	return nil
}

type TotalDigitUnmarshaller struct{}

func (u *TotalDigitUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TotalDigit), stage, identifier, instanceName, preserveOrder)
}

func (u *TotalDigitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TotalDigit)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type UnionUnmarshaller struct{}

func (u *UnionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Union), stage, identifier, instanceName, preserveOrder)
}

func (u *UnionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Union)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "MemberTypes":
		instance.MemberTypes = GongExtractString(valueExpr)
	}
	return nil
}

type WhiteSpaceUnmarshaller struct{}

func (u *WhiteSpaceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(WhiteSpace), stage, identifier, instanceName, preserveOrder)
}

func (u *WhiteSpaceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*WhiteSpace)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallPointer(&instance.Annotation, valueExpr, identifierMap)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}
