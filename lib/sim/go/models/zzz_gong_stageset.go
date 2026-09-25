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
	"os"
	"path/filepath"
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
		for _, command := range __gong__sortStageSetInstances(stageSet.Stage.Commands, stageSet.Stage.Command_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			commandIdent := "__models" + command.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Command{Name: %s}).Stage(stageSet.Stage)", commandIdent, __gong__toRawStringLiteral(command.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", commandIdent, __gong__toRawStringLiteral(command.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Command = %s", commandIdent, __gong__toRawStringLiteral(string(command.Command))))
			values.WriteString(fmt.Sprintf("\n\t%s.CommandDate = %s", commandIdent, __gong__toRawStringLiteral(command.CommandDate)))
			if command.Engine != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + command.Engine.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Engine = %s", commandIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, dummyagent := range __gong__sortStageSetInstances(stageSet.Stage.DummyAgents, stageSet.Stage.DummyAgent_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dummyagentIdent := "__models" + dummyagent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DummyAgent{Name: %s}).Stage(stageSet.Stage)", dummyagentIdent, __gong__toRawStringLiteral(dummyagent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.TechName = %s", dummyagentIdent, __gong__toRawStringLiteral(dummyagent.TechName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dummyagentIdent, __gong__toRawStringLiteral(dummyagent.Name)))
			if dummyagent.Engine != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dummyagent.Engine.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Engine = %s", dummyagentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, engine := range __gong__sortStageSetInstances(stageSet.Stage.Engines, stageSet.Stage.Engine_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			engineIdent := "__models" + engine.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Engine{Name: %s}).Stage(stageSet.Stage)", engineIdent, __gong__toRawStringLiteral(engine.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", engineIdent, __gong__toRawStringLiteral(engine.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EndTime = %s", engineIdent, __gong__toRawStringLiteral(engine.EndTime)))
			values.WriteString(fmt.Sprintf("\n\t%s.CurrentTime = %s", engineIdent, __gong__toRawStringLiteral(engine.CurrentTime)))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayFormat = %s", engineIdent, __gong__toRawStringLiteral(engine.DisplayFormat)))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondsSinceStart = %f", engineIdent, engine.SecondsSinceStart))
			values.WriteString(fmt.Sprintf("\n\t%s.Fired = %d", engineIdent, engine.Fired))
			values.WriteString(fmt.Sprintf("\n\t%s.ControlMode = %s", engineIdent, __gong__toRawStringLiteral(string(engine.ControlMode))))
			values.WriteString(fmt.Sprintf("\n\t%s.State = %s", engineIdent, __gong__toRawStringLiteral(string(engine.State))))
			values.WriteString(fmt.Sprintf("\n\t%s.Speed = %f", engineIdent, engine.Speed))
		}
	}
	if stageSet.Stage != nil {
		for _, event := range __gong__sortStageSetInstances(stageSet.Stage.Events, stageSet.Stage.Event_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			eventIdent := "__models" + event.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Event{Name: %s}).Stage(stageSet.Stage)", eventIdent, __gong__toRawStringLiteral(event.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", eventIdent, __gong__toRawStringLiteral(event.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = time.Duration(%d)", eventIdent, int64(event.Duration)))
		}
	}
	if stageSet.Stage != nil {
		for _, status := range __gong__sortStageSetInstances(stageSet.Stage.Statuss, stageSet.Stage.Status_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			statusIdent := "__models" + status.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Status{Name: %s}).Stage(stageSet.Stage)", statusIdent, __gong__toRawStringLiteral(status.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", statusIdent, __gong__toRawStringLiteral(status.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CurrentCommand = %s", statusIdent, __gong__toRawStringLiteral(string(status.CurrentCommand))))
			values.WriteString(fmt.Sprintf("\n\t%s.CompletionDate = %s", statusIdent, __gong__toRawStringLiteral(status.CompletionDate)))
			values.WriteString(fmt.Sprintf("\n\t%s.CurrentSpeedCommand = %s", statusIdent, __gong__toRawStringLiteral(string(status.CurrentSpeedCommand))))
			values.WriteString(fmt.Sprintf("\n\t%s.SpeedCommandCompletionDate = %s", statusIdent, __gong__toRawStringLiteral(status.SpeedCommandCompletionDate)))
		}
	}
	if stageSet.Stage != nil {
		for _, updatestate := range __gong__sortStageSetInstances(stageSet.Stage.UpdateStates, stageSet.Stage.UpdateState_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			updatestateIdent := "__models" + updatestate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.UpdateState{Name: %s}).Stage(stageSet.Stage)", updatestateIdent, __gong__toRawStringLiteral(updatestate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", updatestateIdent, __gong__toRawStringLiteral(updatestate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = time.Duration(%d)", updatestateIdent, int64(updatestate.Duration)))
			values.WriteString(fmt.Sprintf("\n\t%s.Period = time.Duration(%d)", updatestateIdent, int64(updatestate.Period)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *models.StageSet) {

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
`, packageName, declarations.String(), values.String(), pointers.String())

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
		case "github.com/fullstack-lang/gong/lib/sim/go/models":
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
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "Command":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Command), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DummyAgent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DummyAgent), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Engine":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Engine), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Event":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Event), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Status":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Status), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "UpdateState":
					identifierMap[ident.Name] = __gong__stageSetInit(new(UpdateState), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *Command:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Command":
						inst.Command = CommandType(GongExtractString(rhs))
					case "CommandDate":
						inst.CommandDate = GongExtractString(rhs)
					case "Engine":
						__gong__assignPointer(&inst.Engine, rhs, identifierMap)
					}
				case *DummyAgent:
					switch fieldName {
					case "TechName":
						inst.TechName = GongExtractString(rhs)
					case "Engine":
						__gong__assignPointer(&inst.Engine, rhs, identifierMap)
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Engine:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EndTime":
						inst.EndTime = GongExtractString(rhs)
					case "CurrentTime":
						inst.CurrentTime = GongExtractString(rhs)
					case "DisplayFormat":
						inst.DisplayFormat = GongExtractString(rhs)
					case "SecondsSinceStart":
						inst.SecondsSinceStart = GongExtractFloat(rhs)
					case "Fired":
						inst.Fired = GongExtractInt(rhs)
					case "ControlMode":
						inst.ControlMode = ControlMode(GongExtractString(rhs))
					case "State":
						inst.State = EngineState(GongExtractString(rhs))
					case "Speed":
						inst.Speed = GongExtractFloat(rhs)
					}
				case *Event:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Duration":
						inst.Duration = time.Duration(GongExtractInt(rhs))
					}
				case *Status:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CurrentCommand":
						inst.CurrentCommand = CommandType(GongExtractString(rhs))
					case "CompletionDate":
						inst.CompletionDate = GongExtractString(rhs)
					case "CurrentSpeedCommand":
						inst.CurrentSpeedCommand = SpeedCommandType(GongExtractString(rhs))
					case "SpeedCommandCompletionDate":
						inst.SpeedCommandCompletionDate = GongExtractString(rhs)
					}
				case *UpdateState:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Duration":
						inst.Duration = time.Duration(GongExtractInt(rhs))
					case "Period":
						inst.Period = time.Duration(GongExtractInt(rhs))
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
