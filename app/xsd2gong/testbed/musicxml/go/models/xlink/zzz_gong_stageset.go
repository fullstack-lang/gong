// generated code - do not edit
package xlink

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

)

var (
	_ = time.Hour
	_ = slices.Index[[]int, int]
	_ = sort.Slice
	_ = strconv.Itoa
)

// StageSet coordinates multiple stages across packages
type StageSet struct {
	Stage *Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	return stageSet
}

// GetProbeSplitStageName returns the split stage name for the StageSet probe
func (stageSet *StageSet) GetProbeSplitStageName() string {
	if stageSet.Stage != nil {
		return stageSet.Stage.GetProbeSplitStageName() + "_stageset"
	}
	return "stageset_probe_split"
}

// MarshallFile marshalls all stages into a file
func (stageSet *StageSet) MarshallFile(filename, packageName string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stageSet.Marshall(file, packageName)
}

// Marshall marshalls all stages into an open file
func (stageSet *StageSet) Marshall(file *os.File, packageName string) {
	res, err := stageSet.MarshallToString(packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}
	fmt.Fprintln(file, res)
}

// MarshallToString marshalls all stages into a Go code string
func (stageSet *StageSet) MarshallToString(packageName string) (res string, err error) {
	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stageSet.Stage != nil {
		for _, arctype := range __gong__sortStageSetInstances(stageSet.Stage.ArcTypes, stageSet.Stage.ArcType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arctypeIdent := "__xlink" + arctype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.ArcType{Name: %s}).Stage(stageSet.Stage)", arctypeIdent, __gong__toRawStringLiteral(arctype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", arctypeIdent, __gong__toRawStringLiteral(string(arctype.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", arctypeIdent, __gong__toRawStringLiteral(string(arctype.Show))))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", arctypeIdent, __gong__toRawStringLiteral(string(arctype.Actuate))))
			values.WriteString(fmt.Sprintf("\n\t%s.From = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.From)))
			values.WriteString(fmt.Sprintf("\n\t%s.To = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.To)))
			for _, elem := range arctype.Title_1 {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title_1 = append(%s.Title_1, %s)", arctypeIdent, arctypeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, extended := range __gong__sortStageSetInstances(stageSet.Stage.Extendeds, stageSet.Stage.Extended_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			extendedIdent := "__xlink" + extended.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.Extended{Name: %s}).Stage(stageSet.Stage)", extendedIdent, __gong__toRawStringLiteral(extended.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extendedIdent, __gong__toRawStringLiteral(extended.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", extendedIdent, __gong__toRawStringLiteral(string(extended.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", extendedIdent, __gong__toRawStringLiteral(extended.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", extendedIdent, __gong__toRawStringLiteral(extended.Title)))
			for _, elem := range extended.Title_1 {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title_1 = append(%s.Title_1, %s)", extendedIdent, extendedIdent, targetIdent))
			}
			for _, elem := range extended.Resource {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = append(%s.Resource, %s)", extendedIdent, extendedIdent, targetIdent))
			}
			for _, elem := range extended.Locator {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Locator = append(%s.Locator, %s)", extendedIdent, extendedIdent, targetIdent))
			}
			for _, elem := range extended.Arc {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arc = append(%s.Arc, %s)", extendedIdent, extendedIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, locatortype := range __gong__sortStageSetInstances(stageSet.Stage.LocatorTypes, stageSet.Stage.LocatorType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			locatortypeIdent := "__xlink" + locatortype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.LocatorType{Name: %s}).Stage(stageSet.Stage)", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", locatortypeIdent, __gong__toRawStringLiteral(string(locatortype.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Label)))
			for _, elem := range locatortype.Title_1 {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title_1 = append(%s.Title_1, %s)", locatortypeIdent, locatortypeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, resourcetype := range __gong__sortStageSetInstances(stageSet.Stage.ResourceTypes, stageSet.Stage.ResourceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			resourcetypeIdent := "__xlink" + resourcetype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.ResourceType{Name: %s}).Stage(stageSet.Stage)", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", resourcetypeIdent, __gong__toRawStringLiteral(string(resourcetype.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, simple := range __gong__sortStageSetInstances(stageSet.Stage.Simples, stageSet.Stage.Simple_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			simpleIdent := "__xlink" + simple.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.Simple{Name: %s}).Stage(stageSet.Stage)", simpleIdent, __gong__toRawStringLiteral(simple.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", simpleIdent, __gong__toRawStringLiteral(simple.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", simpleIdent, __gong__toRawStringLiteral(string(simple.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", simpleIdent, __gong__toRawStringLiteral(simple.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", simpleIdent, __gong__toRawStringLiteral(simple.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", simpleIdent, __gong__toRawStringLiteral(simple.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", simpleIdent, __gong__toRawStringLiteral(simple.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", simpleIdent, __gong__toRawStringLiteral(string(simple.Show))))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", simpleIdent, __gong__toRawStringLiteral(string(simple.Actuate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", simpleIdent, __gong__toRawStringLiteral(simple.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, titleelttype := range __gong__sortStageSetInstances(stageSet.Stage.TitleEltTypes, stageSet.Stage.TitleEltType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			titleelttypeIdent := "__xlink" + titleelttype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.TitleEltType{Name: %s}).Stage(stageSet.Stage)", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", titleelttypeIdent, __gong__toRawStringLiteral(string(titleelttype.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Value)))
		}
	}


	var goModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		goModuleVersion = buildInfo.Main.Version
	}

	goModuleVersionWithoutDirty := strings.ReplaceAll(goModuleVersion, "+dirty", "")
	if goModuleVersionWithoutDirty == "" || goModuleVersionWithoutDirty == "(devel)" {
		goModuleVersionWithoutDirty = "latest"
	}

	res = fmt.Sprintf(`// file generated by gong
// go module version: %s
/* to edit this file with the default probe, use the command
go run github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *xlink.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *xlink.StageSet) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, goModuleVersion, goModuleVersionWithoutDirty, packageName, declarations.String(), values.String(), pointers.String())

	return res, nil
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file into stageSet
func (stageSet *StageSet) ParseAstFile(pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist " + pathToFile + " ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file into stageSet
func (stageSet *StageSet) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New("Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string into stageSet
func (stageSet *StageSet) ParseAstString(blob string, preserveOrder bool) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", blob, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances into stageSet
func (stageSet *StageSet) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	identifierMap := make(map[string]any)

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink":
			aliasToCanonical[alias] = "xlink"
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
					}

					switch pkgAlias {
			case "xlink":
				switch typeName {
				case "ArcType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArcType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Extended":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Extended), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LocatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LocatorType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ResourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ResourceType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Simple":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Simple), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TitleEltType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TitleEltType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *ArcType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumTypeType(GongExtractString(rhs))
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = EnumShowType(GongExtractString(rhs))
					case "Actuate":
						inst.Actuate = EnumActuateType(GongExtractString(rhs))
					case "From":
						inst.From = GongExtractString(rhs)
					case "To":
						inst.To = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					}
				case *Extended:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumTypeType(GongExtractString(rhs))
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					case "Resource":
						__gong__assignSliceOfPointers(&inst.Resource, rhs, identifierMap)
					case "Locator":
						__gong__assignSliceOfPointers(&inst.Locator, rhs, identifierMap)
					case "Arc":
						__gong__assignSliceOfPointers(&inst.Arc, rhs, identifierMap)
					}
				case *LocatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumTypeType(GongExtractString(rhs))
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					}
				case *ResourceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumTypeType(GongExtractString(rhs))
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Simple:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumTypeType(GongExtractString(rhs))
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = EnumShowType(GongExtractString(rhs))
					case "Actuate":
						inst.Actuate = EnumActuateType(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *TitleEltType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumTypeType(GongExtractString(rhs))
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
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

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
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
	if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
		if rIdent, ok := call.Args[1].(*ast.Ident); ok {
			if target, ok := identifierMap[rIdent.Name]; ok {
				if typedTarget, ok := target.(*T); ok {
					*slice = append(*slice, typedTarget)
				}
			}
		}
	}
}
