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

	if stageSet.Stage != nil {
		arrowOrdered := []*Arrow{}
		for arrow := range stageSet.Stage.Arrows {
			arrowOrdered = append(arrowOrdered, arrow)
		}
		sort.Slice(arrowOrdered, func(i, j int) bool {
			return stageSet.Stage.Arrow_stagedOrder[arrowOrdered[i]] < stageSet.Stage.Arrow_stagedOrder[arrowOrdered[j]]
		})
		for _, arrow := range arrowOrdered {
			arrowIdent := "__stage_0" + arrow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Arrow{Name: %s}).Stage(stageSet.Stage)", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalColor = %s", arrowIdent, __gong__toRawStringLiteral(arrow.OptionnalColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalStroke = %s", arrowIdent, __gong__toRawStringLiteral(arrow.OptionnalStroke)))
			if arrow.From != nil {
				targetIdent := "__stage_0" + arrow.From.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.From = %s", arrowIdent, targetIdent))
			}
			if arrow.To != nil {
				targetIdent := "__stage_0" + arrow.To.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.To = %s", arrowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		barOrdered := []*Bar{}
		for bar := range stageSet.Stage.Bars {
			barOrdered = append(barOrdered, bar)
		}
		sort.Slice(barOrdered, func(i, j int) bool {
			return stageSet.Stage.Bar_stagedOrder[barOrdered[i]] < stageSet.Stage.Bar_stagedOrder[barOrdered[j]]
		})
		for _, bar := range barOrdered {
			barIdent := "__stage_0" + bar.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Bar{Name: %s}).Stage(stageSet.Stage)", barIdent, __gong__toRawStringLiteral(bar.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barIdent, __gong__toRawStringLiteral(bar.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Start, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", barIdent, bar.Start.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.End, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", barIdent, bar.End.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedDuration = time.Duration(%d)", barIdent, int64(bar.ComputedDuration)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalColor = %s", barIdent, __gong__toRawStringLiteral(bar.OptionnalColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.OptionnalStroke = %s", barIdent, __gong__toRawStringLiteral(bar.OptionnalStroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", barIdent, bar.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", barIdent, bar.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", barIdent, __gong__toRawStringLiteral(bar.StrokeDashArray)))
		}
	}
	if stageSet.Stage != nil {
		ganttOrdered := []*Gantt{}
		for gantt := range stageSet.Stage.Gantts {
			ganttOrdered = append(ganttOrdered, gantt)
		}
		sort.Slice(ganttOrdered, func(i, j int) bool {
			return stageSet.Stage.Gantt_stagedOrder[ganttOrdered[i]] < stageSet.Stage.Gantt_stagedOrder[ganttOrdered[j]]
		})
		for _, gantt := range ganttOrdered {
			ganttIdent := "__stage_0" + gantt.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Gantt{Name: %s}).Stage(stageSet.Stage)", ganttIdent, __gong__toRawStringLiteral(gantt.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ganttIdent, __gong__toRawStringLiteral(gantt.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ComputedStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ComputedEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedDuration = time.Duration(%d)", ganttIdent, int64(gantt.ComputedDuration)))
			values.WriteString(fmt.Sprintf("\n\t%s.UseManualStartAndEndDates = %t", ganttIdent, gantt.UseManualStartAndEndDates))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualStart, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ManualStart.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.ManualEnd, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", ganttIdent, gantt.ManualEnd.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.LaneHeight = %f", ganttIdent, gantt.LaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RatioBarToLaneHeight = %f", ganttIdent, gantt.RatioBarToLaneHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.YTopMargin = %f", ganttIdent, gantt.YTopMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftText = %f", ganttIdent, gantt.XLeftText))
			values.WriteString(fmt.Sprintf("\n\t%s.TextHeight = %f", ganttIdent, gantt.TextHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.XLeftLanes = %f", ganttIdent, gantt.XLeftLanes))
			values.WriteString(fmt.Sprintf("\n\t%s.XRightMargin = %f", ganttIdent, gantt.XRightMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowLengthToTheRightOfStartBar = %f", ganttIdent, gantt.ArrowLengthToTheRightOfStartBar))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowTipLenght = %f", ganttIdent, gantt.ArrowTipLenght))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Color = %s", ganttIdent, __gong__toRawStringLiteral(gantt.TimeLine_Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_FillOpacity = %f", ganttIdent, gantt.TimeLine_FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_Stroke = %s", ganttIdent, __gong__toRawStringLiteral(gantt.TimeLine_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeLine_StrokeWidth = %f", ganttIdent, gantt.TimeLine_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_Stroke = %s", ganttIdent, __gong__toRawStringLiteral(gantt.Group_Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeWidth = %f", ganttIdent, gantt.Group_StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_StrokeDashArray = %s", ganttIdent, __gong__toRawStringLiteral(gantt.Group_StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.DateYOffset = %f", ganttIdent, gantt.DateYOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.AlignOnStartEndOnYearStart = %t", ganttIdent, gantt.AlignOnStartEndOnYearStart))
			for _, elem := range gantt.Lanes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lanes = append(%s.Lanes, %s)", ganttIdent, ganttIdent, targetIdent))
			}
			for _, elem := range gantt.Milestones {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Milestones = append(%s.Milestones, %s)", ganttIdent, ganttIdent, targetIdent))
			}
			for _, elem := range gantt.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", ganttIdent, ganttIdent, targetIdent))
			}
			for _, elem := range gantt.Arrows {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrows = append(%s.Arrows, %s)", ganttIdent, ganttIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		groupOrdered := []*Group{}
		for group := range stageSet.Stage.Groups {
			groupOrdered = append(groupOrdered, group)
		}
		sort.Slice(groupOrdered, func(i, j int) bool {
			return stageSet.Stage.Group_stagedOrder[groupOrdered[i]] < stageSet.Stage.Group_stagedOrder[groupOrdered[j]]
		})
		for _, group := range groupOrdered {
			groupIdent := "__stage_0" + group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Group{Name: %s}).Stage(stageSet.Stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			for _, elem := range group.GroupLanes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupLanes = append(%s.GroupLanes, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		laneOrdered := []*Lane{}
		for lane := range stageSet.Stage.Lanes {
			laneOrdered = append(laneOrdered, lane)
		}
		sort.Slice(laneOrdered, func(i, j int) bool {
			return stageSet.Stage.Lane_stagedOrder[laneOrdered[i]] < stageSet.Stage.Lane_stagedOrder[laneOrdered[j]]
		})
		for _, lane := range laneOrdered {
			laneIdent := "__stage_0" + lane.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Lane{Name: %s}).Stage(stageSet.Stage)", laneIdent, __gong__toRawStringLiteral(lane.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", laneIdent, __gong__toRawStringLiteral(lane.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", laneIdent, lane.Order))
			for _, elem := range lane.Bars {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bars = append(%s.Bars, %s)", laneIdent, laneIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		laneuseOrdered := []*LaneUse{}
		for laneuse := range stageSet.Stage.LaneUses {
			laneuseOrdered = append(laneuseOrdered, laneuse)
		}
		sort.Slice(laneuseOrdered, func(i, j int) bool {
			return stageSet.Stage.LaneUse_stagedOrder[laneuseOrdered[i]] < stageSet.Stage.LaneUse_stagedOrder[laneuseOrdered[j]]
		})
		for _, laneuse := range laneuseOrdered {
			laneuseIdent := "__stage_0" + laneuse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.LaneUse{Name: %s}).Stage(stageSet.Stage)", laneuseIdent, __gong__toRawStringLiteral(laneuse.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", laneuseIdent, __gong__toRawStringLiteral(laneuse.Name)))
			if laneuse.Lane != nil {
				targetIdent := "__stage_0" + laneuse.Lane.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lane = %s", laneuseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		milestoneOrdered := []*Milestone{}
		for milestone := range stageSet.Stage.Milestones {
			milestoneOrdered = append(milestoneOrdered, milestone)
		}
		sort.Slice(milestoneOrdered, func(i, j int) bool {
			return stageSet.Stage.Milestone_stagedOrder[milestoneOrdered[i]] < stageSet.Stage.Milestone_stagedOrder[milestoneOrdered[j]]
		})
		for _, milestone := range milestoneOrdered {
			milestoneIdent := "__stage_0" + milestone.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Milestone{Name: %s}).Stage(stageSet.Stage)", milestoneIdent, __gong__toRawStringLiteral(milestone.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", milestoneIdent, __gong__toRawStringLiteral(milestone.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", milestoneIdent, milestone.Date.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayVerticalBar = %t", milestoneIdent, milestone.DisplayVerticalBar))
			for _, elem := range milestone.LanesToDisplay {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LanesToDisplay = append(%s.LanesToDisplay, %s)", milestoneIdent, milestoneIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/lib/gantt/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *__stage_0__.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *__stage_0__.StageSet) {

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

					switch pkgAlias {
			case "__stage_0__":
				switch typeName {
				case "Arrow":
					if !preserveOrder {
						inst := (&Arrow{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Arrow)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Bar":
					if !preserveOrder {
						inst := (&Bar{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Bar)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Gantt":
					if !preserveOrder {
						inst := (&Gantt{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Gantt)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Group":
					if !preserveOrder {
						inst := (&Group{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Group)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Lane":
					if !preserveOrder {
						inst := (&Lane{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Lane)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "LaneUse":
					if !preserveOrder {
						inst := (&LaneUse{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(LaneUse)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Milestone":
					if !preserveOrder {
						inst := (&Milestone{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Milestone)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
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
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *Arrow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "From":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bar); ok {
									inst.From = typedTarget
								}
							}
						}
					case "To":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bar); ok {
									inst.To = typedTarget
								}
							}
						}
					case "OptionnalColor":
						inst.OptionnalColor = GongExtractString(rhs)
					case "OptionnalStroke":
						inst.OptionnalStroke = GongExtractString(rhs)
					}
				case *Bar:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Start":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "End":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "ComputedDuration":
						inst.ComputedDuration = time.Duration(GongExtractInt(rhs))
					case "OptionnalColor":
						inst.OptionnalColor = GongExtractString(rhs)
					case "OptionnalStroke":
						inst.OptionnalStroke = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					}
				case *Gantt:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedStart":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "ComputedEnd":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "ComputedDuration":
						inst.ComputedDuration = time.Duration(GongExtractInt(rhs))
					case "UseManualStartAndEndDates":
						inst.UseManualStartAndEndDates = GongExtractBool(rhs)
					case "ManualStart":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "ManualEnd":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "LaneHeight":
						inst.LaneHeight = GongExtractFloat(rhs)
					case "RatioBarToLaneHeight":
						inst.RatioBarToLaneHeight = GongExtractFloat(rhs)
					case "YTopMargin":
						inst.YTopMargin = GongExtractFloat(rhs)
					case "XLeftText":
						inst.XLeftText = GongExtractFloat(rhs)
					case "TextHeight":
						inst.TextHeight = GongExtractFloat(rhs)
					case "XLeftLanes":
						inst.XLeftLanes = GongExtractFloat(rhs)
					case "XRightMargin":
						inst.XRightMargin = GongExtractFloat(rhs)
					case "ArrowLengthToTheRightOfStartBar":
						inst.ArrowLengthToTheRightOfStartBar = GongExtractFloat(rhs)
					case "ArrowTipLenght":
						inst.ArrowTipLenght = GongExtractFloat(rhs)
					case "TimeLine_Color":
						inst.TimeLine_Color = GongExtractString(rhs)
					case "TimeLine_FillOpacity":
						inst.TimeLine_FillOpacity = GongExtractFloat(rhs)
					case "TimeLine_Stroke":
						inst.TimeLine_Stroke = GongExtractString(rhs)
					case "TimeLine_StrokeWidth":
						inst.TimeLine_StrokeWidth = GongExtractFloat(rhs)
					case "Group_Stroke":
						inst.Group_Stroke = GongExtractString(rhs)
					case "Group_StrokeWidth":
						inst.Group_StrokeWidth = GongExtractFloat(rhs)
					case "Group_StrokeDashArray":
						inst.Group_StrokeDashArray = GongExtractString(rhs)
					case "DateYOffset":
						inst.DateYOffset = GongExtractFloat(rhs)
					case "AlignOnStartEndOnYearStart":
						inst.AlignOnStartEndOnYearStart = GongExtractBool(rhs)
					case "Lanes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Lane); ok {
										inst.Lanes = append(inst.Lanes, typedTarget)
									}
								}
							}
						}
					case "Milestones":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Milestone); ok {
										inst.Milestones = append(inst.Milestones, typedTarget)
									}
								}
							}
						}
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "Arrows":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Arrow); ok {
										inst.Arrows = append(inst.Arrows, typedTarget)
									}
								}
							}
						}
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GroupLanes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Lane); ok {
										inst.GroupLanes = append(inst.GroupLanes, typedTarget)
									}
								}
							}
						}
					}
				case *Lane:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Bars":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bar); ok {
										inst.Bars = append(inst.Bars, typedTarget)
									}
								}
							}
						}
					}
				case *LaneUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lane":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Lane); ok {
									inst.Lane = typedTarget
								}
							}
						}
					}
				case *Milestone:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Date":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "DisplayVerticalBar":
						inst.DisplayVerticalBar = GongExtractBool(rhs)
					case "LanesToDisplay":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Lane); ok {
										inst.LanesToDisplay = append(inst.LanesToDisplay, typedTarget)
									}
								}
							}
						}
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
