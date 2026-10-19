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
		case "github.com/fullstack-lang/gong/lib/svg/go/models":
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
				case "Animate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Animate), stage, ident.Name, instanceName, preserveOrder)
				case "Circle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Circle), stage, ident.Name, instanceName, preserveOrder)
				case "Condition":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Condition), stage, ident.Name, instanceName, preserveOrder)
				case "ControlPoint":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlPoint), stage, ident.Name, instanceName, preserveOrder)
				case "Ellipse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Ellipse), stage, ident.Name, instanceName, preserveOrder)
				case "FileToDownload":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FileToDownload), stage, ident.Name, instanceName, preserveOrder)
				case "Layer":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Layer), stage, ident.Name, instanceName, preserveOrder)
				case "Line":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Line), stage, ident.Name, instanceName, preserveOrder)
				case "Link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Link), stage, ident.Name, instanceName, preserveOrder)
				case "LinkAnchoredPath":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinkAnchoredPath), stage, ident.Name, instanceName, preserveOrder)
				case "LinkAnchoredText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinkAnchoredText), stage, ident.Name, instanceName, preserveOrder)
				case "Path":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Path), stage, ident.Name, instanceName, preserveOrder)
				case "Point":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Point), stage, ident.Name, instanceName, preserveOrder)
				case "Polygone":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Polygone), stage, ident.Name, instanceName, preserveOrder)
				case "Polyline":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Polyline), stage, ident.Name, instanceName, preserveOrder)
				case "Rect":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Rect), stage, ident.Name, instanceName, preserveOrder)
				case "RectAnchoredPath":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectAnchoredPath), stage, ident.Name, instanceName, preserveOrder)
				case "RectAnchoredPngImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectAnchoredPngImage), stage, ident.Name, instanceName, preserveOrder)
				case "RectAnchoredRect":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectAnchoredRect), stage, ident.Name, instanceName, preserveOrder)
				case "RectAnchoredText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectAnchoredText), stage, ident.Name, instanceName, preserveOrder)
				case "RectLinkLink":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectLinkLink), stage, ident.Name, instanceName, preserveOrder)
				case "SVG":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SVG), stage, ident.Name, instanceName, preserveOrder)
				case "SvgText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SvgText), stage, ident.Name, instanceName, preserveOrder)
				case "Text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Text), stage, ident.Name, instanceName, preserveOrder)
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
				case *Animate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "AttributeName":
						inst.AttributeName = GongExtractString(rhs)
					case "Values":
						inst.Values = GongExtractString(rhs)
					case "From":
						inst.From = GongExtractString(rhs)
					case "To":
						inst.To = GongExtractString(rhs)
					case "Dur":
						inst.Dur = GongExtractString(rhs)
					case "RepeatCount":
						inst.RepeatCount = GongExtractString(rhs)
					}
				case *Circle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CX":
						inst.CX = GongExtractFloat(rhs)
					case "CY":
						inst.CY = GongExtractFloat(rhs)
					case "Radius":
						inst.Radius = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animations":
						__gong__assignSliceOfPointers(&inst.Animations, rhs, identifierMap)
					}
				case *Condition:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *ControlPoint:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X_Relative":
						inst.X_Relative = GongExtractFloat(rhs)
					case "Y_Relative":
						inst.Y_Relative = GongExtractFloat(rhs)
					case "ClosestRect":
						__gong__assignPointer(&inst.ClosestRect, rhs, identifierMap)
					}
				case *Ellipse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CX":
						inst.CX = GongExtractFloat(rhs)
					case "CY":
						inst.CY = GongExtractFloat(rhs)
					case "RX":
						inst.RX = GongExtractFloat(rhs)
					case "RY":
						inst.RY = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *FileToDownload:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64EncodedContent":
						inst.Base64EncodedContent = GongExtractString(rhs)
					}
				case *Layer:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Rects":
						__gong__assignSliceOfPointers(&inst.Rects, rhs, identifierMap)
					case "Texts":
						__gong__assignSliceOfPointers(&inst.Texts, rhs, identifierMap)
					case "Circles":
						__gong__assignSliceOfPointers(&inst.Circles, rhs, identifierMap)
					case "Lines":
						__gong__assignSliceOfPointers(&inst.Lines, rhs, identifierMap)
					case "Ellipses":
						__gong__assignSliceOfPointers(&inst.Ellipses, rhs, identifierMap)
					case "Polylines":
						__gong__assignSliceOfPointers(&inst.Polylines, rhs, identifierMap)
					case "Polygones":
						__gong__assignSliceOfPointers(&inst.Polygones, rhs, identifierMap)
					case "Paths":
						__gong__assignSliceOfPointers(&inst.Paths, rhs, identifierMap)
					case "Links":
						__gong__assignSliceOfPointers(&inst.Links, rhs, identifierMap)
					case "RectLinkLinks":
						__gong__assignSliceOfPointers(&inst.RectLinkLinks, rhs, identifierMap)
					}
				case *Line:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X1":
						inst.X1 = GongExtractFloat(rhs)
					case "Y1":
						inst.Y1 = GongExtractFloat(rhs)
					case "X2":
						inst.X2 = GongExtractFloat(rhs)
					case "Y2":
						inst.Y2 = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					case "MouseClickX":
						inst.MouseClickX = GongExtractFloat(rhs)
					case "MouseClickY":
						inst.MouseClickY = GongExtractFloat(rhs)
					}
				case *Link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						GongUnmarshallEnum(&inst.Type, rhs)
					case "IsBezierCurve":
						inst.IsBezierCurve = GongExtractBool(rhs)
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "StartAnchorType":
						GongUnmarshallEnum(&inst.StartAnchorType, rhs)
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
					case "EndAnchorType":
						GongUnmarshallEnum(&inst.EndAnchorType, rhs)
					case "StartOrientation":
						GongUnmarshallEnum(&inst.StartOrientation, rhs)
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndOrientation":
						GongUnmarshallEnum(&inst.EndOrientation, rhs)
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "CornerRadius":
						inst.CornerRadius = GongExtractFloat(rhs)
					case "HasEndArrow":
						inst.HasEndArrow = GongExtractBool(rhs)
					case "EndArrowSize":
						inst.EndArrowSize = GongExtractFloat(rhs)
					case "EndArrowOffset":
						inst.EndArrowOffset = GongExtractFloat(rhs)
					case "HasStartArrow":
						inst.HasStartArrow = GongExtractBool(rhs)
					case "StartArrowSize":
						inst.StartArrowSize = GongExtractFloat(rhs)
					case "StartArrowOffset":
						inst.StartArrowOffset = GongExtractFloat(rhs)
					case "TextAtArrowStart":
						__gong__assignSliceOfPointers(&inst.TextAtArrowStart, rhs, identifierMap)
					case "TextAtArrowEnd":
						__gong__assignSliceOfPointers(&inst.TextAtArrowEnd, rhs, identifierMap)
					case "TextAtCorner":
						__gong__assignSliceOfPointers(&inst.TextAtCorner, rhs, identifierMap)
					case "PathAtArrowStart":
						__gong__assignSliceOfPointers(&inst.PathAtArrowStart, rhs, identifierMap)
					case "PathAtArrowEnd":
						__gong__assignSliceOfPointers(&inst.PathAtArrowEnd, rhs, identifierMap)
					case "PathAtCorner":
						__gong__assignSliceOfPointers(&inst.PathAtCorner, rhs, identifierMap)
					case "ControlPoints":
						__gong__assignSliceOfPointers(&inst.ControlPoints, rhs, identifierMap)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "MouseX":
						inst.MouseX = GongExtractFloat(rhs)
					case "MouseY":
						inst.MouseY = GongExtractFloat(rhs)
					case "MouseEventKey":
						GongUnmarshallEnum(&inst.MouseEventKey, rhs)
					}
				case *LinkAnchoredPath:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Definition":
						inst.Definition = GongExtractString(rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "ScalePropotionnally":
						inst.ScalePropotionnally = GongExtractBool(rhs)
					case "AppliedScaling":
						inst.AppliedScaling = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					}
				case *LinkAnchoredText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "AutomaticLayout":
						inst.AutomaticLayout = GongExtractBool(rhs)
					case "LinkAnchorType":
						GongUnmarshallEnum(&inst.LinkAnchorType, rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "FontWeight":
						inst.FontWeight = GongExtractString(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = GongExtractString(rhs)
					case "LetterSpacing":
						inst.LetterSpacing = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "WhiteSpace":
						GongUnmarshallEnum(&inst.WhiteSpace, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *Path:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Definition":
						inst.Definition = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *Point:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					}
				case *Polygone:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Points":
						inst.Points = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *Polyline:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Points":
						inst.Points = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *Rect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "RX":
						inst.RX = GongExtractFloat(rhs)
					case "Peers":
						__gong__assignSliceOfPointers(&inst.Peers, rhs, identifierMap)
					case "EnclosingRect":
						__gong__assignPointer(&inst.EnclosingRect, rhs, identifierMap)
					case "Obstacles":
						__gong__assignSliceOfPointers(&inst.Obstacles, rhs, identifierMap)
					case "AnchoredTo":
						__gong__assignPointer(&inst.AnchoredTo, rhs, identifierMap)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "HoveringTrigger":
						__gong__assignSliceOfPointers(&inst.HoveringTrigger, rhs, identifierMap)
					case "DisplayConditions":
						__gong__assignSliceOfPointers(&inst.DisplayConditions, rhs, identifierMap)
					case "Animations":
						__gong__assignSliceOfPointers(&inst.Animations, rhs, identifierMap)
					case "IsSelectable":
						inst.IsSelectable = GongExtractBool(rhs)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "CanHaveLeftHandle":
						inst.CanHaveLeftHandle = GongExtractBool(rhs)
					case "HasLeftHandle":
						inst.HasLeftHandle = GongExtractBool(rhs)
					case "CanHaveRightHandle":
						inst.CanHaveRightHandle = GongExtractBool(rhs)
					case "HasRightHandle":
						inst.HasRightHandle = GongExtractBool(rhs)
					case "CanHaveTopHandle":
						inst.CanHaveTopHandle = GongExtractBool(rhs)
					case "HasTopHandle":
						inst.HasTopHandle = GongExtractBool(rhs)
					case "IsScalingProportionally":
						inst.IsScalingProportionally = GongExtractBool(rhs)
					case "CanHaveBottomHandle":
						inst.CanHaveBottomHandle = GongExtractBool(rhs)
					case "HasBottomHandle":
						inst.HasBottomHandle = GongExtractBool(rhs)
					case "CanMoveHorizontaly":
						inst.CanMoveHorizontaly = GongExtractBool(rhs)
					case "CanMoveVerticaly":
						inst.CanMoveVerticaly = GongExtractBool(rhs)
					case "RectAnchoredTexts":
						__gong__assignSliceOfPointers(&inst.RectAnchoredTexts, rhs, identifierMap)
					case "RectAnchoredRects":
						__gong__assignSliceOfPointers(&inst.RectAnchoredRects, rhs, identifierMap)
					case "RectAnchoredPaths":
						__gong__assignSliceOfPointers(&inst.RectAnchoredPaths, rhs, identifierMap)
					case "RectAnchoredPngImages":
						__gong__assignSliceOfPointers(&inst.RectAnchoredPngImages, rhs, identifierMap)
					case "ChangeColorWhenHovered":
						inst.ChangeColorWhenHovered = GongExtractBool(rhs)
					case "ColorWhenHovered":
						inst.ColorWhenHovered = GongExtractString(rhs)
					case "OriginalColor":
						inst.OriginalColor = GongExtractString(rhs)
					case "FillOpacityWhenHovered":
						inst.FillOpacityWhenHovered = GongExtractFloat(rhs)
					case "OriginalFillOpacity":
						inst.OriginalFillOpacity = GongExtractFloat(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "ToolTipPosition":
						GongUnmarshallEnum(&inst.ToolTipPosition, rhs)
					case "MouseX":
						inst.MouseX = GongExtractFloat(rhs)
					case "MouseY":
						inst.MouseY = GongExtractFloat(rhs)
					case "MouseEventKey":
						GongUnmarshallEnum(&inst.MouseEventKey, rhs)
					case "URLPath":
						inst.URLPath = GongExtractString(rhs)
					case "URLTarget":
						GongUnmarshallEnum(&inst.URLTarget, rhs)
					}
				case *RectAnchoredPath:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Definition":
						inst.Definition = GongExtractString(rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "RectAnchorType":
						GongUnmarshallEnum(&inst.RectAnchorType, rhs)
					case "ScalePropotionnally":
						inst.ScalePropotionnally = GongExtractBool(rhs)
					case "AppliedScaling":
						inst.AppliedScaling = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					}
				case *RectAnchoredPngImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "RX":
						inst.RX = GongExtractFloat(rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "RectAnchorType":
						GongUnmarshallEnum(&inst.RectAnchorType, rhs)
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *RectAnchoredRect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "RX":
						inst.RX = GongExtractFloat(rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "RectAnchorType":
						GongUnmarshallEnum(&inst.RectAnchorType, rhs)
					case "WidthFollowRect":
						inst.WidthFollowRect = GongExtractBool(rhs)
					case "HeightFollowRect":
						inst.HeightFollowRect = GongExtractBool(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					}
				case *RectAnchoredText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = GongExtractString(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = GongExtractString(rhs)
					case "LetterSpacing":
						inst.LetterSpacing = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "WhiteSpace":
						GongUnmarshallEnum(&inst.WhiteSpace, rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "RectAnchorType":
						GongUnmarshallEnum(&inst.RectAnchorType, rhs)
					case "TextAnchorType":
						GongUnmarshallEnum(&inst.TextAnchorType, rhs)
					case "DominantBaseline":
						GongUnmarshallEnum(&inst.DominantBaseline, rhs)
					case "WritingMode":
						GongUnmarshallEnum(&inst.WritingMode, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					case "URLPath":
						inst.URLPath = GongExtractString(rhs)
					case "URLTarget":
						GongUnmarshallEnum(&inst.URLTarget, rhs)
					}
				case *RectLinkLink:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
					case "TargetAnchorPosition":
						inst.TargetAnchorPosition = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					}
				case *SVG:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Layers":
						__gong__assignSliceOfPointers(&inst.Layers, rhs, identifierMap)
					case "DrawingState":
						GongUnmarshallEnum(&inst.DrawingState, rhs)
					case "StartRect":
						__gong__assignPointer(&inst.StartRect, rhs, identifierMap)
					case "EndRect":
						__gong__assignPointer(&inst.EndRect, rhs, identifierMap)
					case "IsEditable":
						inst.IsEditable = GongExtractBool(rhs)
					case "IsSVGFrontEndFileGenerated":
						inst.IsSVGFrontEndFileGenerated = GongExtractBool(rhs)
					case "IsSVGBackEndFileGenerated":
						inst.IsSVGBackEndFileGenerated = GongExtractBool(rhs)
					case "DefaultDirectoryForGeneratedImages":
						inst.DefaultDirectoryForGeneratedImages = GongExtractString(rhs)
					case "IsControlBannerHidden":
						inst.IsControlBannerHidden = GongExtractBool(rhs)
					case "PanX":
						inst.PanX = GongExtractFloat(rhs)
					case "PanY":
						inst.PanY = GongExtractFloat(rhs)
					case "Zoom":
						inst.Zoom = GongExtractFloat(rhs)
					case "OverrideWidth":
						inst.OverrideWidth = GongExtractBool(rhs)
					case "OverriddenWidth":
						inst.OverriddenWidth = GongExtractFloat(rhs)
					case "OverrideHeight":
						inst.OverrideHeight = GongExtractBool(rhs)
					case "OverriddenHeight":
						inst.OverriddenHeight = GongExtractFloat(rhs)
					}
				case *SvgText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					}
				case *Text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = GongExtractString(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = GongExtractString(rhs)
					case "LetterSpacing":
						inst.LetterSpacing = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "WhiteSpace":
						GongUnmarshallEnum(&inst.WhiteSpace, rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
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
type AnimateUnmarshaller struct{}

func (u *AnimateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Animate), stage, identifier, instanceName, preserveOrder)
}

func (u *AnimateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Animate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "AttributeName":
		instance.AttributeName = GongExtractString(valueExpr)
	case "Values":
		instance.Values = GongExtractString(valueExpr)
	case "From":
		instance.From = GongExtractString(valueExpr)
	case "To":
		instance.To = GongExtractString(valueExpr)
	case "Dur":
		instance.Dur = GongExtractString(valueExpr)
	case "RepeatCount":
		instance.RepeatCount = GongExtractString(valueExpr)
	}
	return nil
}

type CircleUnmarshaller struct{}

func (u *CircleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Circle), stage, identifier, instanceName, preserveOrder)
}

func (u *CircleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Circle)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "CX":
		instance.CX = GongExtractFloat(valueExpr)
	case "CY":
		instance.CY = GongExtractFloat(valueExpr)
	case "Radius":
		instance.Radius = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animations":
		GongUnmarshallSliceOfPointers(&instance.Animations, valueExpr, identifierMap)
	}
	return nil
}

type ConditionUnmarshaller struct{}

func (u *ConditionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Condition), stage, identifier, instanceName, preserveOrder)
}

func (u *ConditionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Condition)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ControlPointUnmarshaller struct{}

func (u *ControlPointUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ControlPoint), stage, identifier, instanceName, preserveOrder)
}

func (u *ControlPointUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlPoint)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X_Relative":
		instance.X_Relative = GongExtractFloat(valueExpr)
	case "Y_Relative":
		instance.Y_Relative = GongExtractFloat(valueExpr)
	case "ClosestRect":
		GongUnmarshallPointer(&instance.ClosestRect, valueExpr, identifierMap)
	}
	return nil
}

type EllipseUnmarshaller struct{}

func (u *EllipseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Ellipse), stage, identifier, instanceName, preserveOrder)
}

func (u *EllipseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Ellipse)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "CX":
		instance.CX = GongExtractFloat(valueExpr)
	case "CY":
		instance.CY = GongExtractFloat(valueExpr)
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "RY":
		instance.RY = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type FileToDownloadUnmarshaller struct{}

func (u *FileToDownloadUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FileToDownload), stage, identifier, instanceName, preserveOrder)
}

func (u *FileToDownloadUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FileToDownload)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Base64EncodedContent":
		instance.Base64EncodedContent = GongExtractString(valueExpr)
	}
	return nil
}

type LayerUnmarshaller struct{}

func (u *LayerUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Layer), stage, identifier, instanceName, preserveOrder)
}

func (u *LayerUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Layer)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Rects":
		GongUnmarshallSliceOfPointers(&instance.Rects, valueExpr, identifierMap)
	case "Texts":
		GongUnmarshallSliceOfPointers(&instance.Texts, valueExpr, identifierMap)
	case "Circles":
		GongUnmarshallSliceOfPointers(&instance.Circles, valueExpr, identifierMap)
	case "Lines":
		GongUnmarshallSliceOfPointers(&instance.Lines, valueExpr, identifierMap)
	case "Ellipses":
		GongUnmarshallSliceOfPointers(&instance.Ellipses, valueExpr, identifierMap)
	case "Polylines":
		GongUnmarshallSliceOfPointers(&instance.Polylines, valueExpr, identifierMap)
	case "Polygones":
		GongUnmarshallSliceOfPointers(&instance.Polygones, valueExpr, identifierMap)
	case "Paths":
		GongUnmarshallSliceOfPointers(&instance.Paths, valueExpr, identifierMap)
	case "Links":
		GongUnmarshallSliceOfPointers(&instance.Links, valueExpr, identifierMap)
	case "RectLinkLinks":
		GongUnmarshallSliceOfPointers(&instance.RectLinkLinks, valueExpr, identifierMap)
	}
	return nil
}

type LineUnmarshaller struct{}

func (u *LineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Line), stage, identifier, instanceName, preserveOrder)
}

func (u *LineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Line)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X1":
		instance.X1 = GongExtractFloat(valueExpr)
	case "Y1":
		instance.Y1 = GongExtractFloat(valueExpr)
	case "X2":
		instance.X2 = GongExtractFloat(valueExpr)
	case "Y2":
		instance.Y2 = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	case "MouseClickX":
		instance.MouseClickX = GongExtractFloat(valueExpr)
	case "MouseClickY":
		instance.MouseClickY = GongExtractFloat(valueExpr)
	}
	return nil
}

type LinkUnmarshaller struct{}

func (u *LinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Link), stage, identifier, instanceName, preserveOrder)
}

func (u *LinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Link)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type":
		GongUnmarshallEnum(&instance.Type, valueExpr)
	case "IsBezierCurve":
		instance.IsBezierCurve = GongExtractBool(valueExpr)
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "StartAnchorType":
		GongUnmarshallEnum(&instance.StartAnchorType, valueExpr)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	case "EndAnchorType":
		GongUnmarshallEnum(&instance.EndAnchorType, valueExpr)
	case "StartOrientation":
		GongUnmarshallEnum(&instance.StartOrientation, valueExpr)
	case "StartRatio":
		instance.StartRatio = GongExtractFloat(valueExpr)
	case "EndOrientation":
		GongUnmarshallEnum(&instance.EndOrientation, valueExpr)
	case "EndRatio":
		instance.EndRatio = GongExtractFloat(valueExpr)
	case "CornerOffsetRatio":
		instance.CornerOffsetRatio = GongExtractFloat(valueExpr)
	case "CornerRadius":
		instance.CornerRadius = GongExtractFloat(valueExpr)
	case "HasEndArrow":
		instance.HasEndArrow = GongExtractBool(valueExpr)
	case "EndArrowSize":
		instance.EndArrowSize = GongExtractFloat(valueExpr)
	case "EndArrowOffset":
		instance.EndArrowOffset = GongExtractFloat(valueExpr)
	case "HasStartArrow":
		instance.HasStartArrow = GongExtractBool(valueExpr)
	case "StartArrowSize":
		instance.StartArrowSize = GongExtractFloat(valueExpr)
	case "StartArrowOffset":
		instance.StartArrowOffset = GongExtractFloat(valueExpr)
	case "TextAtArrowStart":
		GongUnmarshallSliceOfPointers(&instance.TextAtArrowStart, valueExpr, identifierMap)
	case "TextAtArrowEnd":
		GongUnmarshallSliceOfPointers(&instance.TextAtArrowEnd, valueExpr, identifierMap)
	case "TextAtCorner":
		GongUnmarshallSliceOfPointers(&instance.TextAtCorner, valueExpr, identifierMap)
	case "PathAtArrowStart":
		GongUnmarshallSliceOfPointers(&instance.PathAtArrowStart, valueExpr, identifierMap)
	case "PathAtArrowEnd":
		GongUnmarshallSliceOfPointers(&instance.PathAtArrowEnd, valueExpr, identifierMap)
	case "PathAtCorner":
		GongUnmarshallSliceOfPointers(&instance.PathAtCorner, valueExpr, identifierMap)
	case "ControlPoints":
		GongUnmarshallSliceOfPointers(&instance.ControlPoints, valueExpr, identifierMap)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "MouseX":
		instance.MouseX = GongExtractFloat(valueExpr)
	case "MouseY":
		instance.MouseY = GongExtractFloat(valueExpr)
	case "MouseEventKey":
		GongUnmarshallEnum(&instance.MouseEventKey, valueExpr)
	}
	return nil
}

type LinkAnchoredPathUnmarshaller struct{}

func (u *LinkAnchoredPathUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinkAnchoredPath), stage, identifier, instanceName, preserveOrder)
}

func (u *LinkAnchoredPathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinkAnchoredPath)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Definition":
		instance.Definition = GongExtractString(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "ScalePropotionnally":
		instance.ScalePropotionnally = GongExtractBool(valueExpr)
	case "AppliedScaling":
		instance.AppliedScaling = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type LinkAnchoredTextUnmarshaller struct{}

func (u *LinkAnchoredTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinkAnchoredText), stage, identifier, instanceName, preserveOrder)
}

func (u *LinkAnchoredTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinkAnchoredText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "AutomaticLayout":
		instance.AutomaticLayout = GongExtractBool(valueExpr)
	case "LinkAnchorType":
		GongUnmarshallEnum(&instance.LinkAnchorType, valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "FontWeight":
		instance.FontWeight = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontStyle":
		instance.FontStyle = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.LetterSpacing = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "WhiteSpace":
		GongUnmarshallEnum(&instance.WhiteSpace, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type PathUnmarshaller struct{}

func (u *PathUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Path), stage, identifier, instanceName, preserveOrder)
}

func (u *PathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Path)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Definition":
		instance.Definition = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type PointUnmarshaller struct{}

func (u *PointUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Point), stage, identifier, instanceName, preserveOrder)
}

func (u *PointUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Point)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type PolygoneUnmarshaller struct{}

func (u *PolygoneUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Polygone), stage, identifier, instanceName, preserveOrder)
}

func (u *PolygoneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Polygone)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		instance.Points = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type PolylineUnmarshaller struct{}

func (u *PolylineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Polyline), stage, identifier, instanceName, preserveOrder)
}

func (u *PolylineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Polyline)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		instance.Points = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}

type RectUnmarshaller struct{}

func (u *RectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Rect), stage, identifier, instanceName, preserveOrder)
}

func (u *RectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rect)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "Peers":
		GongUnmarshallSliceOfPointers(&instance.Peers, valueExpr, identifierMap)
	case "EnclosingRect":
		GongUnmarshallPointer(&instance.EnclosingRect, valueExpr, identifierMap)
	case "Obstacles":
		GongUnmarshallSliceOfPointers(&instance.Obstacles, valueExpr, identifierMap)
	case "AnchoredTo":
		GongUnmarshallPointer(&instance.AnchoredTo, valueExpr, identifierMap)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "HoveringTrigger":
		GongUnmarshallSliceOfPointers(&instance.HoveringTrigger, valueExpr, identifierMap)
	case "DisplayConditions":
		GongUnmarshallSliceOfPointers(&instance.DisplayConditions, valueExpr, identifierMap)
	case "Animations":
		GongUnmarshallSliceOfPointers(&instance.Animations, valueExpr, identifierMap)
	case "IsSelectable":
		instance.IsSelectable = GongExtractBool(valueExpr)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	case "CanHaveLeftHandle":
		instance.CanHaveLeftHandle = GongExtractBool(valueExpr)
	case "HasLeftHandle":
		instance.HasLeftHandle = GongExtractBool(valueExpr)
	case "CanHaveRightHandle":
		instance.CanHaveRightHandle = GongExtractBool(valueExpr)
	case "HasRightHandle":
		instance.HasRightHandle = GongExtractBool(valueExpr)
	case "CanHaveTopHandle":
		instance.CanHaveTopHandle = GongExtractBool(valueExpr)
	case "HasTopHandle":
		instance.HasTopHandle = GongExtractBool(valueExpr)
	case "IsScalingProportionally":
		instance.IsScalingProportionally = GongExtractBool(valueExpr)
	case "CanHaveBottomHandle":
		instance.CanHaveBottomHandle = GongExtractBool(valueExpr)
	case "HasBottomHandle":
		instance.HasBottomHandle = GongExtractBool(valueExpr)
	case "CanMoveHorizontaly":
		instance.CanMoveHorizontaly = GongExtractBool(valueExpr)
	case "CanMoveVerticaly":
		instance.CanMoveVerticaly = GongExtractBool(valueExpr)
	case "RectAnchoredTexts":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredTexts, valueExpr, identifierMap)
	case "RectAnchoredRects":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredRects, valueExpr, identifierMap)
	case "RectAnchoredPaths":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredPaths, valueExpr, identifierMap)
	case "RectAnchoredPngImages":
		GongUnmarshallSliceOfPointers(&instance.RectAnchoredPngImages, valueExpr, identifierMap)
	case "ChangeColorWhenHovered":
		instance.ChangeColorWhenHovered = GongExtractBool(valueExpr)
	case "ColorWhenHovered":
		instance.ColorWhenHovered = GongExtractString(valueExpr)
	case "OriginalColor":
		instance.OriginalColor = GongExtractString(valueExpr)
	case "FillOpacityWhenHovered":
		instance.FillOpacityWhenHovered = GongExtractFloat(valueExpr)
	case "OriginalFillOpacity":
		instance.OriginalFillOpacity = GongExtractFloat(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.ToolTipText = GongExtractString(valueExpr)
	case "ToolTipPosition":
		GongUnmarshallEnum(&instance.ToolTipPosition, valueExpr)
	case "MouseX":
		instance.MouseX = GongExtractFloat(valueExpr)
	case "MouseY":
		instance.MouseY = GongExtractFloat(valueExpr)
	case "MouseEventKey":
		GongUnmarshallEnum(&instance.MouseEventKey, valueExpr)
	case "URLPath":
		instance.URLPath = GongExtractString(valueExpr)
	case "URLTarget":
		GongUnmarshallEnum(&instance.URLTarget, valueExpr)
	}
	return nil
}

type RectAnchoredPathUnmarshaller struct{}

func (u *RectAnchoredPathUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectAnchoredPath), stage, identifier, instanceName, preserveOrder)
}

func (u *RectAnchoredPathUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredPath)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Definition":
		instance.Definition = GongExtractString(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "ScalePropotionnally":
		instance.ScalePropotionnally = GongExtractBool(valueExpr)
	case "AppliedScaling":
		instance.AppliedScaling = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type RectAnchoredPngImageUnmarshaller struct{}

func (u *RectAnchoredPngImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectAnchoredPngImage), stage, identifier, instanceName, preserveOrder)
}

func (u *RectAnchoredPngImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredPngImage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "Base64Content":
		instance.Base64Content = GongExtractString(valueExpr)
	}
	return nil
}

type RectAnchoredRectUnmarshaller struct{}

func (u *RectAnchoredRectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectAnchoredRect), stage, identifier, instanceName, preserveOrder)
}

func (u *RectAnchoredRectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredRect)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "RX":
		instance.RX = GongExtractFloat(valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "WidthFollowRect":
		instance.WidthFollowRect = GongExtractBool(valueExpr)
	case "HeightFollowRect":
		instance.HeightFollowRect = GongExtractBool(valueExpr)
	case "HasToolTip":
		instance.HasToolTip = GongExtractBool(valueExpr)
	case "ToolTipText":
		instance.ToolTipText = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type RectAnchoredTextUnmarshaller struct{}

func (u *RectAnchoredTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectAnchoredText), stage, identifier, instanceName, preserveOrder)
}

func (u *RectAnchoredTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectAnchoredText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "FontWeight":
		instance.FontWeight = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontStyle":
		instance.FontStyle = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.LetterSpacing = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "WhiteSpace":
		GongUnmarshallEnum(&instance.WhiteSpace, valueExpr)
	case "X_Offset":
		instance.X_Offset = GongExtractFloat(valueExpr)
	case "Y_Offset":
		instance.Y_Offset = GongExtractFloat(valueExpr)
	case "RectAnchorType":
		GongUnmarshallEnum(&instance.RectAnchorType, valueExpr)
	case "TextAnchorType":
		GongUnmarshallEnum(&instance.TextAnchorType, valueExpr)
	case "DominantBaseline":
		GongUnmarshallEnum(&instance.DominantBaseline, valueExpr)
	case "WritingMode":
		GongUnmarshallEnum(&instance.WritingMode, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	case "URLPath":
		instance.URLPath = GongExtractString(valueExpr)
	case "URLTarget":
		GongUnmarshallEnum(&instance.URLTarget, valueExpr)
	}
	return nil
}

type RectLinkLinkUnmarshaller struct{}

func (u *RectLinkLinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RectLinkLink), stage, identifier, instanceName, preserveOrder)
}

func (u *RectLinkLinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RectLinkLink)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Start":
		GongUnmarshallPointer(&instance.Start, valueExpr, identifierMap)
	case "End":
		GongUnmarshallPointer(&instance.End, valueExpr, identifierMap)
	case "TargetAnchorPosition":
		instance.TargetAnchorPosition = GongExtractFloat(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	}
	return nil
}

type SVGUnmarshaller struct{}

func (u *SVGUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SVG), stage, identifier, instanceName, preserveOrder)
}

func (u *SVGUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SVG)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Layers":
		GongUnmarshallSliceOfPointers(&instance.Layers, valueExpr, identifierMap)
	case "DrawingState":
		GongUnmarshallEnum(&instance.DrawingState, valueExpr)
	case "StartRect":
		GongUnmarshallPointer(&instance.StartRect, valueExpr, identifierMap)
	case "EndRect":
		GongUnmarshallPointer(&instance.EndRect, valueExpr, identifierMap)
	case "IsEditable":
		instance.IsEditable = GongExtractBool(valueExpr)
	case "IsSVGFrontEndFileGenerated":
		instance.IsSVGFrontEndFileGenerated = GongExtractBool(valueExpr)
	case "IsSVGBackEndFileGenerated":
		instance.IsSVGBackEndFileGenerated = GongExtractBool(valueExpr)
	case "DefaultDirectoryForGeneratedImages":
		instance.DefaultDirectoryForGeneratedImages = GongExtractString(valueExpr)
	case "IsControlBannerHidden":
		instance.IsControlBannerHidden = GongExtractBool(valueExpr)
	case "PanX":
		instance.PanX = GongExtractFloat(valueExpr)
	case "PanY":
		instance.PanY = GongExtractFloat(valueExpr)
	case "Zoom":
		instance.Zoom = GongExtractFloat(valueExpr)
	case "OverrideWidth":
		instance.OverrideWidth = GongExtractBool(valueExpr)
	case "OverriddenWidth":
		instance.OverriddenWidth = GongExtractFloat(valueExpr)
	case "OverrideHeight":
		instance.OverrideHeight = GongExtractBool(valueExpr)
	case "OverriddenHeight":
		instance.OverriddenHeight = GongExtractFloat(valueExpr)
	}
	return nil
}

type SvgTextUnmarshaller struct{}

func (u *SvgTextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SvgText), stage, identifier, instanceName, preserveOrder)
}

func (u *SvgTextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SvgText)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	}
	return nil
}

type TextUnmarshaller struct{}

func (u *TextUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Text), stage, identifier, instanceName, preserveOrder)
}

func (u *TextUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Content":
		instance.Content = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "FillOpacity":
		instance.FillOpacity = GongExtractFloat(valueExpr)
	case "Stroke":
		instance.Stroke = GongExtractString(valueExpr)
	case "StrokeOpacity":
		instance.StrokeOpacity = GongExtractFloat(valueExpr)
	case "StrokeWidth":
		instance.StrokeWidth = GongExtractFloat(valueExpr)
	case "StrokeDashArray":
		instance.StrokeDashArray = GongExtractString(valueExpr)
	case "StrokeDashArrayWhenSelected":
		instance.StrokeDashArrayWhenSelected = GongExtractString(valueExpr)
	case "Transform":
		instance.Transform = GongExtractString(valueExpr)
	case "FontWeight":
		instance.FontWeight = GongExtractString(valueExpr)
	case "FontSize":
		instance.FontSize = GongExtractString(valueExpr)
	case "FontStyle":
		instance.FontStyle = GongExtractString(valueExpr)
	case "LetterSpacing":
		instance.LetterSpacing = GongExtractString(valueExpr)
	case "FontFamily":
		instance.FontFamily = GongExtractString(valueExpr)
	case "WhiteSpace":
		GongUnmarshallEnum(&instance.WhiteSpace, valueExpr)
	case "Animates":
		GongUnmarshallSliceOfPointers(&instance.Animates, valueExpr, identifierMap)
	}
	return nil
}
