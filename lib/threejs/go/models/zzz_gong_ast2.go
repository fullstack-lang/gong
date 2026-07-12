// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
	"time"
)

var (
	_time__dummyDeclaration2 time.Duration
	_                        = _time__dummyDeclaration2
)

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

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// ModelUnmarshaller abstracts the logic for setting fields on a staged instance
type ModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

func (stage *Stage) UnmarshallFile(pathToFile string, preserveOrder bool) error {
	return ParseAstFile(stage, pathToFile, preserveOrder)
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func ParseAstFile(stage *Stage, pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
func ParseAstEmbeddedFile(stage *Stage, directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New(stage.GetName() + "; Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset, false)
}

// GongParseAstString parses the Go source code from a string
func GongParseAstString(stage *Stage, blob string, preserveOrder bool) error {
	fileString := "package main\nfunc _() {\n" + blob + "\n}"
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", fileString, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances using the Unmarshaller registry
func ParseAstFileFromAst(stage *Stage, inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {

	var fileModuleVersion string
	for _, commentGroup := range inFile.Comments {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "// go module version: ") {
				fileModuleVersion = strings.TrimPrefix(comment.Text, "// go module version: ")
			}
		}
	}

	var parserModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		parserModuleVersion = buildInfo.Main.Version
	}

	if fileModuleVersion != "" && parserModuleVersion != "" && fileModuleVersion != parserModuleVersion {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(time.Now(), fmt.Sprintf("Warning: file module version '%s' does not match parser module version '%s'", fileModuleVersion, parserModuleVersion))
		}
	}

	// 1. Remove Global Variables: Use a local map to track variable names to instances
	identifierMap := make(map[string]GongstructIF)

	for _, instance := range stage.GetInstances() {
		identifierMap[instance.GongGetIdentifier(stage)] = instance
	}

	// 2. Visitor Pattern: Traverse the AST
	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var typeName string
					var instanceName string

					// Inspect RHS to find the Struct Type and Name
					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								typeName = selExpr.Sel.Name
								// Attempt to find Name field in literal
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

					// Dispatch to specific Unmarshaller
					if typeName != "" {
						if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
							instance, err := unmarshaller.Initialize(stage, ident.Name, instanceName, preserveOrder)
							if err == nil {
								identifierMap[ident.Name] = instance
							}
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
							typeName := instance.GongGetGongstructName()
							if unmarshaller, exists := stage.GongUnmarshallers[typeName]; exists {
								// 3. Strategy Pattern: Delegate to Handler
								unmarshaller.UnmarshallField(stage, instance, selExpr.Sel.Name, node.Rhs[0], identifierMap)
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
								instance.UnstageVoid(stage)
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
							if stage.IsInDeltaMode() && stage.navigationMode != GongNavigationModeNavigating {
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
								// 1. Run all Before Commit hooks
								for _, hook := range stage.beforeCommitHooks {
									hook(stage)
								}

								// 2. Run all After Commit hooks
								for _, hook := range stage.afterCommitHooks {
									hook(stage)
								}
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

// --- Generic Helpers for Unmarshallers ---

func GongExtractString(expr ast.Expr) string {
	switch e := expr.(type) {

	// Case 1: It's a standard string literal
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			if unquoted, err := strconv.Unquote(e.Value); err == nil {
				return unquoted
			}
		}

	// Case 2: It's a concatenated string
	case *ast.BinaryExpr:
		// We only care if they are being added together
		if e.Op == token.ADD {
			// Recursively extract the left and right sides
			left := GongExtractString(e.X)
			right := GongExtractString(e.Y)

			// Join them back together
			return left + right
		}
	}

	return ""
}

func GongExtractInt(expr ast.Expr) int {
	if bl, ok := expr.(*ast.BasicLit); ok {
		val, _ := strconv.Atoi(bl.Value)
		return val
	}
	if ue, ok := expr.(*ast.UnaryExpr); ok && ue.Op == token.SUB {
		if bl, ok := ue.X.(*ast.BasicLit); ok {
			val, _ := strconv.Atoi(bl.Value)
			return -val
		}
	}
	return 0
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

func GongExtractFloat(expr ast.Expr) float64 {
	if bl, ok := expr.(*ast.BasicLit); ok {
		val, _ := strconv.ParseFloat(bl.Value, 64)
		return val
	}
	if ue, ok := expr.(*ast.UnaryExpr); ok && ue.Op == token.SUB {
		if bl, ok := ue.X.(*ast.BasicLit); ok {
			val, _ := strconv.ParseFloat(bl.Value, 64)
			return -val
		}
	}
	return 0.0
}

func GongExtractBool(expr ast.Expr) bool {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name == "true"
	}
	return false
}

func GongExtractExpr(expr ast.Expr) any {
	switch v := expr.(type) {
	case *ast.BasicLit:
		return v.Value
	case *ast.CompositeLit:
		// Reconstruct "Package.Struct{}"
		if sel, ok := v.Type.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok {
				return id.Name + "." + sel.Sel.Name + "{}"
			}
		}
	case *ast.SelectorExpr:
		// Reconstruct "Package.Struct{}.Field"
		// X is likely a CompositeLit (Package.Struct{})
		if cl, ok := v.X.(*ast.CompositeLit); ok {
			if sel, ok := cl.Type.(*ast.SelectorExpr); ok {
				if id, ok := sel.X.(*ast.Ident); ok {
					return id.Name + "." + sel.Sel.Name + "{}." + v.Sel.Name
				}
			}
		}
		// Reconstruct "Package.Identifier"
		if id, ok := v.X.(*ast.Ident); ok {
			return id.Name + "." + v.Sel.Name
		}
	case *ast.CallExpr:
		// Reconstruct "new(Package.Struct)"
		if fun, ok := v.Fun.(*ast.Ident); ok && fun.Name == "new" {
			if len(v.Args) == 1 {
				if sel, ok := v.Args[0].(*ast.SelectorExpr); ok {
					if id, ok := sel.X.(*ast.Ident); ok {
						return "new(" + id.Name + "." + sel.Sel.Name + ")"
					}
				}
			}
		}
	}
	return ""
}

// GongUnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func GongUnmarshallSliceOfPointers[T PointerToGongstruct](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) (err error) {

	if call, ok := valueExpr.(*ast.CallExpr); ok {
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
				if end > len(*slice) {
					// Handle error: log warning, resize slice, or return error
					// For example:
					return fmt.Errorf("index out of bounds: %d for len %d", end, len(*slice))
				}
				*slice = slices.Delete(*slice, start, end)
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if index > len(*slice) {
					return fmt.Errorf("index out of bounds: %d for len %d", index, len(*slice))
				}
				if ident, ok := call.Args[2].(*ast.Ident); ok {
					if val, ok := identifierMap[ident.Name]; ok {
						*slice = slices.Insert(*slice, index, val.(T))
					} else {
						log.Println("Ast2 Insert Unkown identifier", ident.Name)
					}
				}
			}
		} else if funcName == "append" {
			if len(call.Args) >= 2 {
				for i := 1; i < len(call.Args); i++ {
					if ident, ok := call.Args[i].(*ast.Ident); ok {
						if val, ok := identifierMap[ident.Name]; ok {
							*slice = append(*slice, val.(T))
						} else {
							log.Println("Ast2 append Unkown identifier", ident.Name)
						}
					}
				}
			}
		}
	}
	return nil
}

// GongUnmarshallPointer handles assignment of a single pointer field
func GongUnmarshallPointer[T PointerToGongstruct](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {

	if ident, ok := valueExpr.(*ast.Ident); ok {
		if ident.Name == "nil" {
			var zero T
			*ptr = zero
			return
		}
		if val, ok := identifierMap[ident.Name]; ok {
			*ptr = val.(T)
		}
	}
}

// GongUnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func GongUnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr) {

	// Case 1: Standard Enum usage (models.EnumType_Value)
	if sel, ok := valueExpr.(*ast.SelectorExpr); ok {
		if err := ptr.FromCodeString(sel.Sel.Name); err != nil {
			log.Printf("UnmarshallEnum: Error parsing code string '%s': %v", sel.Sel.Name, err)
		}
	} else {
		// Case 2: Fallback to string literal
		valStr := GongExtractString(valueExpr)
		if valStr != "" {
			if err := ptr.FromCodeString(valStr); err != nil {
				log.Printf("UnmarshallEnum: Error parsing string literal '%s': %v", valStr, err)
			}
		}
	}
}

// insertion point per named struct
type AmbiantLightUnmarshaller struct{}

func (u *AmbiantLightUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(AmbiantLight)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *AmbiantLightUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AmbiantLight)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Intensity":
		instance.Intensity = GongExtractFloat(valueExpr)
	}
	return nil
}

type BoxGeometryUnmarshaller struct{}

func (u *BoxGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(BoxGeometry)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *BoxGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BoxGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "Depth":
		instance.Depth = GongExtractFloat(valueExpr)
	case "WidthSegments":
		instance.WidthSegments = GongExtractInt(valueExpr)
	case "HeightSegments":
		instance.HeightSegments = GongExtractInt(valueExpr)
	case "DepthSegments":
		instance.DepthSegments = GongExtractInt(valueExpr)
	}
	return nil
}

type CameraUnmarshaller struct{}

func (u *CameraUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Camera)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *CameraUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Camera)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Z":
		instance.Z = GongExtractFloat(valueExpr)
	case "TargetX":
		instance.TargetX = GongExtractFloat(valueExpr)
	case "TargetY":
		instance.TargetY = GongExtractFloat(valueExpr)
	case "TargetZ":
		instance.TargetZ = GongExtractFloat(valueExpr)
	case "Fov":
		instance.Fov = GongExtractFloat(valueExpr)
	}
	return nil
}

type CanvasUnmarshaller struct{}

func (u *CanvasUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Canvas)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *CanvasUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Canvas)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "DirectionalLights":
		GongUnmarshallSliceOfPointers(&instance.DirectionalLights, valueExpr, identifierMap)
	case "AmbiantLight":
		GongUnmarshallPointer(&instance.AmbiantLight, valueExpr, identifierMap)
	case "Meshs":
		GongUnmarshallSliceOfPointers(&instance.Meshs, valueExpr, identifierMap)
	case "Camera":
		GongUnmarshallPointer(&instance.Camera, valueExpr, identifierMap)
	}
	return nil
}

type CurveUnmarshaller struct{}

func (u *CurveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Curve)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *CurveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Curve)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		GongUnmarshallSliceOfPointers(&instance.Points, valueExpr, identifierMap)
	}
	return nil
}

type CylinderGeometryUnmarshaller struct{}

func (u *CylinderGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(CylinderGeometry)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *CylinderGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CylinderGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RadiusTop":
		instance.RadiusTop = GongExtractFloat(valueExpr)
	case "RadiusBottom":
		instance.RadiusBottom = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "RadialSegments":
		instance.RadialSegments = GongExtractInt(valueExpr)
	case "HeightSegments":
		instance.HeightSegments = GongExtractInt(valueExpr)
	case "OpenEnded":
		instance.OpenEnded = GongExtractBool(valueExpr)
	case "ThetaStart":
		instance.ThetaStart = GongExtractFloat(valueExpr)
	case "ThetaLength":
		instance.ThetaLength = GongExtractFloat(valueExpr)
	}
	return nil
}

type DirectionalLightUnmarshaller struct{}

func (u *DirectionalLightUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(DirectionalLight)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *DirectionalLightUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectionalLight)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Z":
		instance.Z = GongExtractFloat(valueExpr)
	case "Intensity":
		instance.Intensity = GongExtractFloat(valueExpr)
	case "IsWithCastShadow":
		instance.IsWithCastShadow = GongExtractBool(valueExpr)
	}
	return nil
}

type ExtrudeGeometryUnmarshaller struct{}

func (u *ExtrudeGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(ExtrudeGeometry)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ExtrudeGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ExtrudeGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Shape":
		GongUnmarshallPointer(&instance.Shape, valueExpr, identifierMap)
	case "ExtrudePath":
		GongUnmarshallPointer(&instance.ExtrudePath, valueExpr, identifierMap)
	case "Steps":
		instance.Steps = GongExtractInt(valueExpr)
	}
	return nil
}

type MeshUnmarshaller struct{}

func (u *MeshUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Mesh)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MeshUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Mesh)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Z":
		instance.Z = GongExtractFloat(valueExpr)
	case "MeshMaterialBasic":
		GongUnmarshallPointer(&instance.MeshMaterialBasic, valueExpr, identifierMap)
	case "MeshPhysicalMaterial":
		GongUnmarshallPointer(&instance.MeshPhysicalMaterial, valueExpr, identifierMap)
	case "CylinderGeometry":
		GongUnmarshallPointer(&instance.CylinderGeometry, valueExpr, identifierMap)
	case "BoxGeometry":
		GongUnmarshallPointer(&instance.BoxGeometry, valueExpr, identifierMap)
	case "SphereGeometry":
		GongUnmarshallPointer(&instance.SphereGeometry, valueExpr, identifierMap)
	case "TorusGeometry":
		GongUnmarshallPointer(&instance.TorusGeometry, valueExpr, identifierMap)
	case "PlaneGeometry":
		GongUnmarshallPointer(&instance.PlaneGeometry, valueExpr, identifierMap)
	case "TubeGeometry":
		GongUnmarshallPointer(&instance.TubeGeometry, valueExpr, identifierMap)
	case "ExtrudeGeometry":
		GongUnmarshallPointer(&instance.ExtrudeGeometry, valueExpr, identifierMap)
	}
	return nil
}

type MeshMaterialBasicUnmarshaller struct{}

func (u *MeshMaterialBasicUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MeshMaterialBasic)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MeshMaterialBasicUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeshMaterialBasic)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type MeshPhysicalMaterialUnmarshaller struct{}

func (u *MeshPhysicalMaterialUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(MeshPhysicalMaterial)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *MeshPhysicalMaterialUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeshPhysicalMaterial)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Wireframe":
		instance.Wireframe = GongExtractBool(valueExpr)
	case "Opacity":
		instance.Opacity = GongExtractFloat(valueExpr)
	case "Transparent":
		instance.Transparent = GongExtractBool(valueExpr)
	case "Visible":
		instance.Visible = GongExtractBool(valueExpr)
	}
	return nil
}

type PlaneGeometryUnmarshaller struct{}

func (u *PlaneGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(PlaneGeometry)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *PlaneGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PlaneGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "WidthSegments":
		instance.WidthSegments = GongExtractInt(valueExpr)
	case "HeightSegments":
		instance.HeightSegments = GongExtractInt(valueExpr)
	}
	return nil
}

type ShapeUnmarshaller struct{}

func (u *ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Shape)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Points":
		GongUnmarshallSliceOfPointers(&instance.Points, valueExpr, identifierMap)
	}
	return nil
}

type SphereGeometryUnmarshaller struct{}

func (u *SphereGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(SphereGeometry)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *SphereGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SphereGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Radius":
		instance.Radius = GongExtractFloat(valueExpr)
	case "WidthSegments":
		instance.WidthSegments = GongExtractInt(valueExpr)
	case "HeightSegments":
		instance.HeightSegments = GongExtractInt(valueExpr)
	case "PhiStart":
		instance.PhiStart = GongExtractFloat(valueExpr)
	case "PhiLength":
		instance.PhiLength = GongExtractFloat(valueExpr)
	case "ThetaStart":
		instance.ThetaStart = GongExtractFloat(valueExpr)
	case "ThetaLength":
		instance.ThetaLength = GongExtractFloat(valueExpr)
	}
	return nil
}

type TorusGeometryUnmarshaller struct{}

func (u *TorusGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TorusGeometry)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TorusGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TorusGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Radius":
		instance.Radius = GongExtractFloat(valueExpr)
	case "Tube":
		instance.Tube = GongExtractFloat(valueExpr)
	case "RadialSegments":
		instance.RadialSegments = GongExtractInt(valueExpr)
	case "TubularSegments":
		instance.TubularSegments = GongExtractInt(valueExpr)
	case "Arc":
		instance.Arc = GongExtractFloat(valueExpr)
	}
	return nil
}

type TubeGeometryUnmarshaller struct{}

func (u *TubeGeometryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(TubeGeometry)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *TubeGeometryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TubeGeometry)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Path":
		GongUnmarshallPointer(&instance.Path, valueExpr, identifierMap)
	case "TubularSegments":
		instance.TubularSegments = GongExtractInt(valueExpr)
	case "Radius":
		instance.Radius = GongExtractFloat(valueExpr)
	case "RadialSegments":
		instance.RadialSegments = GongExtractInt(valueExpr)
	case "Closed":
		instance.Closed = GongExtractBool(valueExpr)
	}
	return nil
}

type Vector2Unmarshaller struct{}

func (u *Vector2Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Vector2)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *Vector2Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Vector2)
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

type Vector3Unmarshaller struct{}

func (u *Vector3Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance := new(Vector3)
	instance.Name = instanceName
	if !preserveOrder {
		instance.Stage(stage)
	} else {
		if newOrder, err := ExtractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.Stage(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

func (u *Vector3Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Vector3)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Z":
		instance.Z = GongExtractFloat(valueExpr)
	}
	return nil
}
