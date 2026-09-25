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
		for _, astruct := range __gong__sortStageSetInstances(stageSet.Stage.Astructs, stageSet.Stage.Astruct_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			astructIdent := "__models" + astruct.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Astruct{Name: %s}).Stage(stageSet.Stage)", astructIdent, __gong__toRawStringLiteral(astruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astruct.Associationtob.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Associationtob = %s", astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anarrayofb {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofb = append(%s.Anarrayofb, %s)", astructIdent, astructIdent, targetIdent))
			}
			if astruct.Anotherassociationtob_2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astruct.Anotherassociationtob_2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anotherassociationtob_2 = %s", astructIdent, targetIdent))
			}
			if astruct.Bstruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astruct.Bstruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstruct = %s", astructIdent, targetIdent))
			}
			if astruct.Bstruct2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astruct.Bstruct2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstruct2 = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astruct.Dstruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astruct.Dstruct2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct2 = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct3 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astruct.Dstruct3.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct3 = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct4 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astruct.Dstruct4.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct4 = %s", astructIdent, targetIdent))
			}
			for _, elem := range astruct.Dstruct4s {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct4s = append(%s.Dstruct4s, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anarrayofa {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofa = append(%s.Anarrayofa, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anotherarrayofb {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anotherarrayofb = append(%s.Anotherarrayofb, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.AnarrayofbUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnarrayofbUse = append(%s.AnarrayofbUse, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anarrayofb2Use {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofb2Use = append(%s.Anarrayofb2Use, %s)", astructIdent, astructIdent, targetIdent))
			}
			if astruct.AnAstruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astruct.AnAstruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnAstruct = %s", astructIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, astructbstruct2use := range __gong__sortStageSetInstances(stageSet.Stage.AstructBstruct2Uses, stageSet.Stage.AstructBstruct2Use_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			astructbstruct2useIdent := "__models" + astructbstruct2use.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AstructBstruct2Use{Name: %s}).Stage(stageSet.Stage)", astructbstruct2useIdent, __gong__toRawStringLiteral(astructbstruct2use.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", astructbstruct2useIdent, __gong__toRawStringLiteral(astructbstruct2use.Name)))
			if astructbstruct2use.Bstrcut2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astructbstruct2use.Bstrcut2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstrcut2 = %s", astructbstruct2useIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, astructbstructuse := range __gong__sortStageSetInstances(stageSet.Stage.AstructBstructUses, stageSet.Stage.AstructBstructUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			astructbstructuseIdent := "__models" + astructbstructuse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AstructBstructUse{Name: %s}).Stage(stageSet.Stage)", astructbstructuseIdent, __gong__toRawStringLiteral(astructbstructuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", astructbstructuseIdent, __gong__toRawStringLiteral(astructbstructuse.Name)))
			if astructbstructuse.Bstruct2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + astructbstructuse.Bstruct2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstruct2 = %s", astructbstructuseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, bstruct := range __gong__sortStageSetInstances(stageSet.Stage.Bstructs, stageSet.Stage.Bstruct_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bstructIdent := "__models" + bstruct.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bstruct{Name: %s}).Stage(stageSet.Stage)", bstructIdent, __gong__toRawStringLiteral(bstruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bstructIdent, __gong__toRawStringLiteral(bstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield = %f", bstructIdent, bstruct.Floatfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield2 = %f", bstructIdent, bstruct.Floatfield2))
			values.WriteString(fmt.Sprintf("\n\t%s.Intfield = %d", bstructIdent, bstruct.Intfield))
			values.WriteString(fmt.Sprintf("\n\t%s.ToBeIgnored = %d", bstructIdent, bstruct.ToBeIgnored))
		}
	}
	if stageSet.Stage != nil {
		for _, dstruct := range __gong__sortStageSetInstances(stageSet.Stage.Dstructs, stageSet.Stage.Dstruct_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dstructIdent := "__models" + dstruct.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Dstruct{Name: %s}).Stage(stageSet.Stage)", dstructIdent, __gong__toRawStringLiteral(dstruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dstructIdent, __gong__toRawStringLiteral(dstruct.Name)))
			for _, elem := range dstruct.Anarrayofb {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofb = append(%s.Anarrayofb, %s)", dstructIdent, dstructIdent, targetIdent))
			}
			if dstruct.Gstruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + dstruct.Gstruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Gstruct = %s", dstructIdent, targetIdent))
			}
			for _, elem := range dstruct.Gstructs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Gstructs = append(%s.Gstructs, %s)", dstructIdent, dstructIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, f0123456789012345678901234567890 := range __gong__sortStageSetInstances(stageSet.Stage.F0123456789012345678901234567890s, stageSet.Stage.F0123456789012345678901234567890_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			f0123456789012345678901234567890Ident := "__models" + f0123456789012345678901234567890.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.F0123456789012345678901234567890{Name: %s}).Stage(stageSet.Stage)", f0123456789012345678901234567890Ident, __gong__toRawStringLiteral(f0123456789012345678901234567890.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", f0123456789012345678901234567890Ident, __gong__toRawStringLiteral(f0123456789012345678901234567890.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", f0123456789012345678901234567890Ident, f0123456789012345678901234567890.Date.String()))
		}
	}
	if stageSet.Stage != nil {
		for _, gstruct := range __gong__sortStageSetInstances(stageSet.Stage.Gstructs, stageSet.Stage.Gstruct_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gstructIdent := "__models" + gstruct.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Gstruct{Name: %s}).Stage(stageSet.Stage)", gstructIdent, __gong__toRawStringLiteral(gstruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
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

	"github.com/fullstack-lang/gong/test/test1/go/models"
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
		case "github.com/fullstack-lang/gong/test/test1/go/models":
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
				case "Astruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Astruct), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "AstructBstruct2Use":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AstructBstruct2Use), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "AstructBstructUse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AstructBstructUse), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bstruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bstruct), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Dstruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Dstruct), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "F0123456789012345678901234567890":
					identifierMap[ident.Name] = __gong__stageSetInit(new(F0123456789012345678901234567890), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Gstruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Gstruct), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
						__gong__assignPointer(&inst.Associationtob, rhs, identifierMap)
					case "Anarrayofb":
						__gong__assignSliceOfPointers(&inst.Anarrayofb, rhs, identifierMap)
					case "Anotherassociationtob_2":
						__gong__assignPointer(&inst.Anotherassociationtob_2, rhs, identifierMap)
					case "Date":
						inst.Date = GongExtractDate(rhs)
					case "Date2":
						inst.Date2 = GongExtractDate(rhs)
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
						__gong__assignPointer(&inst.Bstruct, rhs, identifierMap)
					case "Bstruct2":
						__gong__assignPointer(&inst.Bstruct2, rhs, identifierMap)
					case "Dstruct":
						__gong__assignPointer(&inst.Dstruct, rhs, identifierMap)
					case "Dstruct2":
						__gong__assignPointer(&inst.Dstruct2, rhs, identifierMap)
					case "Dstruct3":
						__gong__assignPointer(&inst.Dstruct3, rhs, identifierMap)
					case "Dstruct4":
						__gong__assignPointer(&inst.Dstruct4, rhs, identifierMap)
					case "Dstruct4s":
						__gong__assignSliceOfPointers(&inst.Dstruct4s, rhs, identifierMap)
					case "Floatfield":
						inst.Floatfield = GongExtractFloat(rhs)
					case "Intfield":
						inst.Intfield = GongExtractInt(rhs)
					case "Anotherbooleanfield":
						inst.Anotherbooleanfield = GongExtractBool(rhs)
					case "Duration1":
						inst.Duration1 = time.Duration(GongExtractInt(rhs))
					case "Anarrayofa":
						__gong__assignSliceOfPointers(&inst.Anarrayofa, rhs, identifierMap)
					case "Anotherarrayofb":
						__gong__assignSliceOfPointers(&inst.Anotherarrayofb, rhs, identifierMap)
					case "AnarrayofbUse":
						__gong__assignSliceOfPointers(&inst.AnarrayofbUse, rhs, identifierMap)
					case "Anarrayofb2Use":
						__gong__assignSliceOfPointers(&inst.Anarrayofb2Use, rhs, identifierMap)
					case "AnAstruct":
						__gong__assignPointer(&inst.AnAstruct, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Bstrcut2, rhs, identifierMap)
					}
				case *AstructBstructUse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Bstruct2":
						__gong__assignPointer(&inst.Bstruct2, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.Anarrayofb, rhs, identifierMap)
					case "Gstruct":
						__gong__assignPointer(&inst.Gstruct, rhs, identifierMap)
					case "Gstructs":
						__gong__assignSliceOfPointers(&inst.Gstructs, rhs, identifierMap)
					}
				case *F0123456789012345678901234567890:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Date":
						inst.Date = GongExtractDate(rhs)
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
