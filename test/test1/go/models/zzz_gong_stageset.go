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
		astructOrdered := []*Astruct{}
		for astruct := range stageSet.Stage.Astructs {
			astructOrdered = append(astructOrdered, astruct)
		}
		sort.Slice(astructOrdered, func(i, j int) bool {
			return stageSet.Stage.Astruct_stagedOrder[astructOrdered[i]] < stageSet.Stage.Astruct_stagedOrder[astructOrdered[j]]
		})
		for _, astruct := range astructOrdered {
			astructIdent := "__stage_0" + astruct.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Astruct{Name: %s}).Stage(stageSet.Stage)", astructIdent, __gong__toRawStringLiteral(astruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", astructIdent, __gong__toRawStringLiteral(astruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", astructIdent, astruct.Date.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Date2, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", astructIdent, astruct.Date2.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Booleanfield = %t", astructIdent, astruct.Booleanfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Aenum = %s", astructIdent, __gong__toRawStringLiteral(string(astruct.Aenum))))
			values.WriteString(fmt.Sprintf("\n\t%s.Aenum_2 = %s", astructIdent, __gong__toRawStringLiteral(string(astruct.Aenum_2))))
			values.WriteString(fmt.Sprintf("\n\t%s.Benum = %s", astructIdent, __gong__toRawStringLiteral(string(astruct.Benum))))
			values.WriteString(fmt.Sprintf("\n\t%s.CEnum = %d", astructIdent, int(astruct.CEnum)))
			values.WriteString(fmt.Sprintf("\n\t%s.CName = %s", astructIdent, __gong__toRawStringLiteral(astruct.CName)))
			values.WriteString(fmt.Sprintf("\n\t%s.CFloatfield = %f", astructIdent, astruct.CFloatfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield = %f", astructIdent, astruct.Floatfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Intfield = %d", astructIdent, astruct.Intfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Anotherbooleanfield = %t", astructIdent, astruct.Anotherbooleanfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration1 = time.Duration(%d)", astructIdent, int64(astruct.Duration1)))
			values.WriteString(fmt.Sprintf("\n\t%s.TextFieldBespokeSize = %s", astructIdent, __gong__toRawStringLiteral(astruct.TextFieldBespokeSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.TextArea = %s", astructIdent, __gong__toRawStringLiteral(astruct.TextArea)))
			if astruct.Associationtob != nil {
				targetIdent := "__stage_0" + astruct.Associationtob.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Associationtob = %s", astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anarrayofb {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofb = append(%s.Anarrayofb, %s)", astructIdent, astructIdent, targetIdent))
			}
			if astruct.Anotherassociationtob_2 != nil {
				targetIdent := "__stage_0" + astruct.Anotherassociationtob_2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anotherassociationtob_2 = %s", astructIdent, targetIdent))
			}
			if astruct.Bstruct != nil {
				targetIdent := "__stage_0" + astruct.Bstruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstruct = %s", astructIdent, targetIdent))
			}
			if astruct.Bstruct2 != nil {
				targetIdent := "__stage_0" + astruct.Bstruct2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstruct2 = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct != nil {
				targetIdent := "__stage_0" + astruct.Dstruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct2 != nil {
				targetIdent := "__stage_0" + astruct.Dstruct2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct2 = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct3 != nil {
				targetIdent := "__stage_0" + astruct.Dstruct3.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct3 = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct4 != nil {
				targetIdent := "__stage_0" + astruct.Dstruct4.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct4 = %s", astructIdent, targetIdent))
			}
			for _, elem := range astruct.Dstruct4s {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct4s = append(%s.Dstruct4s, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anarrayofa {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofa = append(%s.Anarrayofa, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anotherarrayofb {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anotherarrayofb = append(%s.Anotherarrayofb, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.AnarrayofbUse {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnarrayofbUse = append(%s.AnarrayofbUse, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anarrayofb2Use {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofb2Use = append(%s.Anarrayofb2Use, %s)", astructIdent, astructIdent, targetIdent))
			}
			if astruct.AnAstruct != nil {
				targetIdent := "__stage_0" + astruct.AnAstruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnAstruct = %s", astructIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		astructbstruct2useOrdered := []*AstructBstruct2Use{}
		for astructbstruct2use := range stageSet.Stage.AstructBstruct2Uses {
			astructbstruct2useOrdered = append(astructbstruct2useOrdered, astructbstruct2use)
		}
		sort.Slice(astructbstruct2useOrdered, func(i, j int) bool {
			return stageSet.Stage.AstructBstruct2Use_stagedOrder[astructbstruct2useOrdered[i]] < stageSet.Stage.AstructBstruct2Use_stagedOrder[astructbstruct2useOrdered[j]]
		})
		for _, astructbstruct2use := range astructbstruct2useOrdered {
			astructbstruct2useIdent := "__stage_0" + astructbstruct2use.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.AstructBstruct2Use{Name: %s}).Stage(stageSet.Stage)", astructbstruct2useIdent, __gong__toRawStringLiteral(astructbstruct2use.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", astructbstruct2useIdent, __gong__toRawStringLiteral(astructbstruct2use.Name)))
			if astructbstruct2use.Bstrcut2 != nil {
				targetIdent := "__stage_0" + astructbstruct2use.Bstrcut2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstrcut2 = %s", astructbstruct2useIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		astructbstructuseOrdered := []*AstructBstructUse{}
		for astructbstructuse := range stageSet.Stage.AstructBstructUses {
			astructbstructuseOrdered = append(astructbstructuseOrdered, astructbstructuse)
		}
		sort.Slice(astructbstructuseOrdered, func(i, j int) bool {
			return stageSet.Stage.AstructBstructUse_stagedOrder[astructbstructuseOrdered[i]] < stageSet.Stage.AstructBstructUse_stagedOrder[astructbstructuseOrdered[j]]
		})
		for _, astructbstructuse := range astructbstructuseOrdered {
			astructbstructuseIdent := "__stage_0" + astructbstructuse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.AstructBstructUse{Name: %s}).Stage(stageSet.Stage)", astructbstructuseIdent, __gong__toRawStringLiteral(astructbstructuse.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", astructbstructuseIdent, __gong__toRawStringLiteral(astructbstructuse.Name)))
			if astructbstructuse.Bstruct2 != nil {
				targetIdent := "__stage_0" + astructbstructuse.Bstruct2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstruct2 = %s", astructbstructuseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		bstructOrdered := []*Bstruct{}
		for bstruct := range stageSet.Stage.Bstructs {
			bstructOrdered = append(bstructOrdered, bstruct)
		}
		sort.Slice(bstructOrdered, func(i, j int) bool {
			return stageSet.Stage.Bstruct_stagedOrder[bstructOrdered[i]] < stageSet.Stage.Bstruct_stagedOrder[bstructOrdered[j]]
		})
		for _, bstruct := range bstructOrdered {
			bstructIdent := "__stage_0" + bstruct.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Bstruct{Name: %s}).Stage(stageSet.Stage)", bstructIdent, __gong__toRawStringLiteral(bstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bstructIdent, __gong__toRawStringLiteral(bstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield = %f", bstructIdent, bstruct.Floatfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield2 = %f", bstructIdent, bstruct.Floatfield2))
			values.WriteString(fmt.Sprintf("\n\t%s.Intfield = %d", bstructIdent, bstruct.Intfield))
			values.WriteString(fmt.Sprintf("\n\t%s.ToBeIgnored = %d", bstructIdent, bstruct.ToBeIgnored))
		}
	}
	if stageSet.Stage != nil {
		dstructOrdered := []*Dstruct{}
		for dstruct := range stageSet.Stage.Dstructs {
			dstructOrdered = append(dstructOrdered, dstruct)
		}
		sort.Slice(dstructOrdered, func(i, j int) bool {
			return stageSet.Stage.Dstruct_stagedOrder[dstructOrdered[i]] < stageSet.Stage.Dstruct_stagedOrder[dstructOrdered[j]]
		})
		for _, dstruct := range dstructOrdered {
			dstructIdent := "__stage_0" + dstruct.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Dstruct{Name: %s}).Stage(stageSet.Stage)", dstructIdent, __gong__toRawStringLiteral(dstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dstructIdent, __gong__toRawStringLiteral(dstruct.Name)))
			for _, elem := range dstruct.Anarrayofb {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofb = append(%s.Anarrayofb, %s)", dstructIdent, dstructIdent, targetIdent))
			}
			if dstruct.Gstruct != nil {
				targetIdent := "__stage_0" + dstruct.Gstruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Gstruct = %s", dstructIdent, targetIdent))
			}
			for _, elem := range dstruct.Gstructs {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Gstructs = append(%s.Gstructs, %s)", dstructIdent, dstructIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		f0123456789012345678901234567890Ordered := []*F0123456789012345678901234567890{}
		for f0123456789012345678901234567890 := range stageSet.Stage.F0123456789012345678901234567890s {
			f0123456789012345678901234567890Ordered = append(f0123456789012345678901234567890Ordered, f0123456789012345678901234567890)
		}
		sort.Slice(f0123456789012345678901234567890Ordered, func(i, j int) bool {
			return stageSet.Stage.F0123456789012345678901234567890_stagedOrder[f0123456789012345678901234567890Ordered[i]] < stageSet.Stage.F0123456789012345678901234567890_stagedOrder[f0123456789012345678901234567890Ordered[j]]
		})
		for _, f0123456789012345678901234567890 := range f0123456789012345678901234567890Ordered {
			f0123456789012345678901234567890Ident := "__stage_0" + f0123456789012345678901234567890.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.F0123456789012345678901234567890{Name: %s}).Stage(stageSet.Stage)", f0123456789012345678901234567890Ident, __gong__toRawStringLiteral(f0123456789012345678901234567890.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", f0123456789012345678901234567890Ident, __gong__toRawStringLiteral(f0123456789012345678901234567890.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", f0123456789012345678901234567890Ident, f0123456789012345678901234567890.Date.String()))
		}
	}
	if stageSet.Stage != nil {
		gstructOrdered := []*Gstruct{}
		for gstruct := range stageSet.Stage.Gstructs {
			gstructOrdered = append(gstructOrdered, gstruct)
		}
		sort.Slice(gstructOrdered, func(i, j int) bool {
			return stageSet.Stage.Gstruct_stagedOrder[gstructOrdered[i]] < stageSet.Stage.Gstruct_stagedOrder[gstructOrdered[j]]
		})
		for _, gstruct := range gstructOrdered {
			gstructIdent := "__stage_0" + gstruct.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Gstruct{Name: %s}).Stage(stageSet.Stage)", gstructIdent, __gong__toRawStringLiteral(gstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gstructIdent, __gong__toRawStringLiteral(gstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield = %f", gstructIdent, gstruct.Floatfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield2 = %f", gstructIdent, gstruct.Floatfield2))
			values.WriteString(fmt.Sprintf("\n\t%s.Intfield = %d", gstructIdent, gstruct.Intfield))
			values.WriteString(fmt.Sprintf("\n\t%s.ToBeIgnored = %d", gstructIdent, gstruct.ToBeIgnored))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/test/test1/go/models"
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
				case "Astruct":
					if !preserveOrder {
						inst := (&Astruct{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Astruct)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "AstructBstruct2Use":
					if !preserveOrder {
						inst := (&AstructBstruct2Use{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AstructBstruct2Use)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "AstructBstructUse":
					if !preserveOrder {
						inst := (&AstructBstructUse{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AstructBstructUse)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Bstruct":
					if !preserveOrder {
						inst := (&Bstruct{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Bstruct)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Dstruct":
					if !preserveOrder {
						inst := (&Dstruct{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Dstruct)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "F0123456789012345678901234567890":
					if !preserveOrder {
						inst := (&F0123456789012345678901234567890{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(F0123456789012345678901234567890)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Gstruct":
					if !preserveOrder {
						inst := (&Gstruct{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Gstruct)
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
				case *Astruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Associationtob":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bstruct); ok {
									inst.Associationtob = typedTarget
								}
							}
						}
					case "Anarrayofb":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bstruct); ok {
										inst.Anarrayofb = append(inst.Anarrayofb, typedTarget)
									}
								}
							}
						}
					case "Anotherassociationtob_2":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bstruct); ok {
									inst.Anotherassociationtob_2 = typedTarget
								}
							}
						}
					case "Date":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "Date2":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Date2, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "Booleanfield":
						inst.Booleanfield = GongExtractBool(rhs)
					case "Aenum":
						inst.Aenum = AEnumType(GongExtractString(rhs))
					case "Aenum_2":
						inst.Aenum_2 = AEnumType(GongExtractString(rhs))
					case "Benum":
						inst.Benum = BEnumType(GongExtractString(rhs))
					case "CEnum":
						inst.CEnum = CEnumTypeInt(GongExtractInt(rhs))
					case "CName":
						inst.CName = GongExtractString(rhs)
					case "CFloatfield":
						inst.CFloatfield = GongExtractFloat(rhs)
					case "Bstruct":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bstruct); ok {
									inst.Bstruct = typedTarget
								}
							}
						}
					case "Bstruct2":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bstruct); ok {
									inst.Bstruct2 = typedTarget
								}
							}
						}
					case "Dstruct":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Dstruct); ok {
									inst.Dstruct = typedTarget
								}
							}
						}
					case "Dstruct2":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Dstruct); ok {
									inst.Dstruct2 = typedTarget
								}
							}
						}
					case "Dstruct3":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Dstruct); ok {
									inst.Dstruct3 = typedTarget
								}
							}
						}
					case "Dstruct4":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Dstruct); ok {
									inst.Dstruct4 = typedTarget
								}
							}
						}
					case "Dstruct4s":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Dstruct); ok {
										inst.Dstruct4s = append(inst.Dstruct4s, typedTarget)
									}
								}
							}
						}
					case "Floatfield":
						inst.Floatfield = GongExtractFloat(rhs)
					case "Intfield":
						inst.Intfield = GongExtractInt(rhs)
					case "Anotherbooleanfield":
						inst.Anotherbooleanfield = GongExtractBool(rhs)
					case "Duration1":
						inst.Duration1 = time.Duration(GongExtractInt(rhs))
					case "Anarrayofa":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Astruct); ok {
										inst.Anarrayofa = append(inst.Anarrayofa, typedTarget)
									}
								}
							}
						}
					case "Anotherarrayofb":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bstruct); ok {
										inst.Anotherarrayofb = append(inst.Anotherarrayofb, typedTarget)
									}
								}
							}
						}
					case "AnarrayofbUse":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AstructBstructUse); ok {
										inst.AnarrayofbUse = append(inst.AnarrayofbUse, typedTarget)
									}
								}
							}
						}
					case "Anarrayofb2Use":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AstructBstruct2Use); ok {
										inst.Anarrayofb2Use = append(inst.Anarrayofb2Use, typedTarget)
									}
								}
							}
						}
					case "AnAstruct":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Astruct); ok {
									inst.AnAstruct = typedTarget
								}
							}
						}
					case "TextFieldBespokeSize":
						inst.TextFieldBespokeSize = GongExtractString(rhs)
					case "TextArea":
						inst.TextArea = GongExtractString(rhs)
					}
				case *AstructBstruct2Use:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Bstrcut2":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bstruct); ok {
									inst.Bstrcut2 = typedTarget
								}
							}
						}
					}
				case *AstructBstructUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Bstruct2":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bstruct); ok {
									inst.Bstruct2 = typedTarget
								}
							}
						}
					}
				case *Bstruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Floatfield":
						inst.Floatfield = GongExtractFloat(rhs)
					case "Floatfield2":
						inst.Floatfield2 = GongExtractFloat(rhs)
					case "Intfield":
						inst.Intfield = GongExtractInt(rhs)
					case "ToBeIgnored":
						inst.ToBeIgnored = GongExtractInt(rhs)
					}
				case *Dstruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Anarrayofb":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bstruct); ok {
										inst.Anarrayofb = append(inst.Anarrayofb, typedTarget)
									}
								}
							}
						}
					case "Gstruct":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Gstruct); ok {
									inst.Gstruct = typedTarget
								}
							}
						}
					case "Gstructs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Gstruct); ok {
										inst.Gstructs = append(inst.Gstructs, typedTarget)
									}
								}
							}
						}
					}
				case *F0123456789012345678901234567890:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Date":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					}
				case *Gstruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Floatfield":
						inst.Floatfield = GongExtractFloat(rhs)
					case "Floatfield2":
						inst.Floatfield2 = GongExtractFloat(rhs)
					case "Intfield":
						inst.Intfield = GongExtractInt(rhs)
					case "ToBeIgnored":
						inst.ToBeIgnored = GongExtractInt(rhs)
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
