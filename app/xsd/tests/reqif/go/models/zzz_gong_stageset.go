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
		for _, alternative_id := range __gong__sortStageSetInstances(stageSet.Stage.ALTERNATIVE_IDs, stageSet.Stage.ALTERNATIVE_ID_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			alternative_idIdent := "__models" + alternative_id.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ALTERNATIVE_ID{Name: %s}).Stage(stageSet.Stage)", alternative_idIdent, __gong__toRawStringLiteral(alternative_id.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", alternative_idIdent, __gong__toRawStringLiteral(alternative_id.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", alternative_idIdent, __gong__toRawStringLiteral(alternative_id.IDENTIFIER)))
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_definition_boolean := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_DEFINITION_BOOLEANs, stageSet.Stage.ATTRIBUTE_DEFINITION_BOOLEAN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_booleanIdent := "__models" + attribute_definition_boolean.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_BOOLEAN{Name: %s}).Stage(stageSet.Stage)", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_booleanIdent, attribute_definition_boolean.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.LONG_NAME)))
			if attribute_definition_boolean.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_boolean.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_booleanIdent, targetIdent))
			}
			if attribute_definition_boolean.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_boolean.DEFAULT_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_booleanIdent, targetIdent))
			}
			if attribute_definition_boolean.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_boolean.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_booleanIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_definition_date := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_DEFINITION_DATEs, stageSet.Stage.ATTRIBUTE_DEFINITION_DATE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_dateIdent := "__models" + attribute_definition_date.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_DATE{Name: %s}).Stage(stageSet.Stage)", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_dateIdent, attribute_definition_date.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.LONG_NAME)))
			if attribute_definition_date.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_date.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_dateIdent, targetIdent))
			}
			if attribute_definition_date.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_date.DEFAULT_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_dateIdent, targetIdent))
			}
			if attribute_definition_date.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_date.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_dateIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_definition_enumeration := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, stageSet.Stage.ATTRIBUTE_DEFINITION_ENUMERATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_enumerationIdent := "__models" + attribute_definition_enumeration.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_ENUMERATION{Name: %s}).Stage(stageSet.Stage)", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_enumerationIdent, attribute_definition_enumeration.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.LONG_NAME)))
			values.WriteString(fmt.Sprintf("\n\t%s.MULTI_VALUED = %t", attribute_definition_enumerationIdent, attribute_definition_enumeration.MULTI_VALUED))
			if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_enumeration.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_enumerationIdent, targetIdent))
			}
			if attribute_definition_enumeration.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_enumeration.DEFAULT_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_enumerationIdent, targetIdent))
			}
			if attribute_definition_enumeration.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_enumeration.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_enumerationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_definition_integer := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_DEFINITION_INTEGERs, stageSet.Stage.ATTRIBUTE_DEFINITION_INTEGER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_integerIdent := "__models" + attribute_definition_integer.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_INTEGER{Name: %s}).Stage(stageSet.Stage)", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_integerIdent, attribute_definition_integer.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.LONG_NAME)))
			if attribute_definition_integer.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_integer.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_integerIdent, targetIdent))
			}
			if attribute_definition_integer.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_integer.DEFAULT_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_integerIdent, targetIdent))
			}
			if attribute_definition_integer.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_integer.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_integerIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_definition_real := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_DEFINITION_REALs, stageSet.Stage.ATTRIBUTE_DEFINITION_REAL_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_realIdent := "__models" + attribute_definition_real.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_REAL{Name: %s}).Stage(stageSet.Stage)", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_realIdent, attribute_definition_real.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.LONG_NAME)))
			if attribute_definition_real.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_real.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_realIdent, targetIdent))
			}
			if attribute_definition_real.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_real.DEFAULT_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_realIdent, targetIdent))
			}
			if attribute_definition_real.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_real.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_realIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_definition_string := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_DEFINITION_STRINGs, stageSet.Stage.ATTRIBUTE_DEFINITION_STRING_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_stringIdent := "__models" + attribute_definition_string.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_STRING{Name: %s}).Stage(stageSet.Stage)", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_stringIdent, attribute_definition_string.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.LONG_NAME)))
			if attribute_definition_string.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_string.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_stringIdent, targetIdent))
			}
			if attribute_definition_string.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_string.DEFAULT_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_stringIdent, targetIdent))
			}
			if attribute_definition_string.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_string.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_stringIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_definition_xhtml := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_DEFINITION_XHTMLs, stageSet.Stage.ATTRIBUTE_DEFINITION_XHTML_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_xhtmlIdent := "__models" + attribute_definition_xhtml.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_XHTML{Name: %s}).Stage(stageSet.Stage)", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_xhtmlIdent, attribute_definition_xhtml.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.LONG_NAME)))
			if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_xhtml.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_xhtmlIdent, targetIdent))
			}
			if attribute_definition_xhtml.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_xhtml.DEFAULT_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_xhtmlIdent, targetIdent))
			}
			if attribute_definition_xhtml.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_definition_xhtml.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_xhtmlIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_value_boolean := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_VALUE_BOOLEANs, stageSet.Stage.ATTRIBUTE_VALUE_BOOLEAN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_booleanIdent := "__models" + attribute_value_boolean.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_BOOLEAN{Name: %s}).Stage(stageSet.Stage)", attribute_value_booleanIdent, __gong__toRawStringLiteral(attribute_value_boolean.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_booleanIdent, __gong__toRawStringLiteral(attribute_value_boolean.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %t", attribute_value_booleanIdent, attribute_value_boolean.THE_VALUE))
			if attribute_value_boolean.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_boolean.DEFINITION.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_booleanIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_value_date := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_VALUE_DATEs, stageSet.Stage.ATTRIBUTE_VALUE_DATE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_dateIdent := "__models" + attribute_value_date.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_DATE{Name: %s}).Stage(stageSet.Stage)", attribute_value_dateIdent, __gong__toRawStringLiteral(attribute_value_date.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_dateIdent, __gong__toRawStringLiteral(attribute_value_date.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %s", attribute_value_dateIdent, __gong__toRawStringLiteral(attribute_value_date.THE_VALUE)))
			if attribute_value_date.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_date.DEFINITION.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_dateIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_value_enumeration := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_VALUE_ENUMERATIONs, stageSet.Stage.ATTRIBUTE_VALUE_ENUMERATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_enumerationIdent := "__models" + attribute_value_enumeration.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_ENUMERATION{Name: %s}).Stage(stageSet.Stage)", attribute_value_enumerationIdent, __gong__toRawStringLiteral(attribute_value_enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_enumerationIdent, __gong__toRawStringLiteral(attribute_value_enumeration.Name)))
			if attribute_value_enumeration.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_enumeration.DEFINITION.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_enumerationIdent, targetIdent))
			}
			if attribute_value_enumeration.VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_enumeration.VALUES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VALUES = %s", attribute_value_enumerationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_value_integer := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_VALUE_INTEGERs, stageSet.Stage.ATTRIBUTE_VALUE_INTEGER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_integerIdent := "__models" + attribute_value_integer.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_INTEGER{Name: %s}).Stage(stageSet.Stage)", attribute_value_integerIdent, __gong__toRawStringLiteral(attribute_value_integer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_integerIdent, __gong__toRawStringLiteral(attribute_value_integer.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %d", attribute_value_integerIdent, attribute_value_integer.THE_VALUE))
			if attribute_value_integer.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_integer.DEFINITION.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_integerIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_value_real := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_VALUE_REALs, stageSet.Stage.ATTRIBUTE_VALUE_REAL_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_realIdent := "__models" + attribute_value_real.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_REAL{Name: %s}).Stage(stageSet.Stage)", attribute_value_realIdent, __gong__toRawStringLiteral(attribute_value_real.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_realIdent, __gong__toRawStringLiteral(attribute_value_real.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %f", attribute_value_realIdent, attribute_value_real.THE_VALUE))
			if attribute_value_real.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_real.DEFINITION.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_realIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_value_string := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_VALUE_STRINGs, stageSet.Stage.ATTRIBUTE_VALUE_STRING_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_stringIdent := "__models" + attribute_value_string.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_STRING{Name: %s}).Stage(stageSet.Stage)", attribute_value_stringIdent, __gong__toRawStringLiteral(attribute_value_string.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_stringIdent, __gong__toRawStringLiteral(attribute_value_string.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %s", attribute_value_stringIdent, __gong__toRawStringLiteral(attribute_value_string.THE_VALUE)))
			if attribute_value_string.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_string.DEFINITION.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_stringIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute_value_xhtml := range __gong__sortStageSetInstances(stageSet.Stage.ATTRIBUTE_VALUE_XHTMLs, stageSet.Stage.ATTRIBUTE_VALUE_XHTML_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_xhtmlIdent := "__models" + attribute_value_xhtml.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_XHTML{Name: %s}).Stage(stageSet.Stage)", attribute_value_xhtmlIdent, __gong__toRawStringLiteral(attribute_value_xhtml.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_xhtmlIdent, __gong__toRawStringLiteral(attribute_value_xhtml.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_SIMPLIFIED = %t", attribute_value_xhtmlIdent, attribute_value_xhtml.IS_SIMPLIFIED))
			if attribute_value_xhtml.THE_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_xhtml.THE_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %s", attribute_value_xhtmlIdent, targetIdent))
			}
			if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_xhtml.THE_ORIGINAL_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.THE_ORIGINAL_VALUE = %s", attribute_value_xhtmlIdent, targetIdent))
			}
			if attribute_value_xhtml.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute_value_xhtml.DEFINITION.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_xhtmlIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_alternative_id := range __gong__sortStageSetInstances(stageSet.Stage.A_ALTERNATIVE_IDs, stageSet.Stage.A_ALTERNATIVE_ID_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_alternative_idIdent := "__models" + a_alternative_id.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ALTERNATIVE_ID{Name: %s}).Stage(stageSet.Stage)", a_alternative_idIdent, __gong__toRawStringLiteral(a_alternative_id.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_alternative_idIdent, __gong__toRawStringLiteral(a_alternative_id.Name)))
			if a_alternative_id.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + a_alternative_id.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", a_alternative_idIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_definition_boolean_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs, stageSet.Stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_boolean_refIdent := "__models" + a_attribute_definition_boolean_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF{Name: %s}).Stage(stageSet.Stage)", a_attribute_definition_boolean_refIdent, __gong__toRawStringLiteral(a_attribute_definition_boolean_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_boolean_refIdent, __gong__toRawStringLiteral(a_attribute_definition_boolean_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_BOOLEAN_REF = %s", a_attribute_definition_boolean_refIdent, __gong__toRawStringLiteral(a_attribute_definition_boolean_ref.ATTRIBUTE_DEFINITION_BOOLEAN_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_definition_date_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_DEFINITION_DATE_REFs, stageSet.Stage.A_ATTRIBUTE_DEFINITION_DATE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_date_refIdent := "__models" + a_attribute_definition_date_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_DATE_REF{Name: %s}).Stage(stageSet.Stage)", a_attribute_definition_date_refIdent, __gong__toRawStringLiteral(a_attribute_definition_date_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_date_refIdent, __gong__toRawStringLiteral(a_attribute_definition_date_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_DATE_REF = %s", a_attribute_definition_date_refIdent, __gong__toRawStringLiteral(a_attribute_definition_date_ref.ATTRIBUTE_DEFINITION_DATE_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_definition_enumeration_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs, stageSet.Stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_enumeration_refIdent := "__models" + a_attribute_definition_enumeration_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF{Name: %s}).Stage(stageSet.Stage)", a_attribute_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_attribute_definition_enumeration_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_attribute_definition_enumeration_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_ENUMERATION_REF = %s", a_attribute_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_attribute_definition_enumeration_ref.ATTRIBUTE_DEFINITION_ENUMERATION_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_definition_integer_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs, stageSet.Stage.A_ATTRIBUTE_DEFINITION_INTEGER_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_integer_refIdent := "__models" + a_attribute_definition_integer_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_INTEGER_REF{Name: %s}).Stage(stageSet.Stage)", a_attribute_definition_integer_refIdent, __gong__toRawStringLiteral(a_attribute_definition_integer_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_integer_refIdent, __gong__toRawStringLiteral(a_attribute_definition_integer_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_INTEGER_REF = %s", a_attribute_definition_integer_refIdent, __gong__toRawStringLiteral(a_attribute_definition_integer_ref.ATTRIBUTE_DEFINITION_INTEGER_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_definition_real_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_DEFINITION_REAL_REFs, stageSet.Stage.A_ATTRIBUTE_DEFINITION_REAL_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_real_refIdent := "__models" + a_attribute_definition_real_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_REAL_REF{Name: %s}).Stage(stageSet.Stage)", a_attribute_definition_real_refIdent, __gong__toRawStringLiteral(a_attribute_definition_real_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_real_refIdent, __gong__toRawStringLiteral(a_attribute_definition_real_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_REAL_REF = %s", a_attribute_definition_real_refIdent, __gong__toRawStringLiteral(a_attribute_definition_real_ref.ATTRIBUTE_DEFINITION_REAL_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_definition_string_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_DEFINITION_STRING_REFs, stageSet.Stage.A_ATTRIBUTE_DEFINITION_STRING_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_string_refIdent := "__models" + a_attribute_definition_string_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_STRING_REF{Name: %s}).Stage(stageSet.Stage)", a_attribute_definition_string_refIdent, __gong__toRawStringLiteral(a_attribute_definition_string_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_string_refIdent, __gong__toRawStringLiteral(a_attribute_definition_string_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_STRING_REF = %s", a_attribute_definition_string_refIdent, __gong__toRawStringLiteral(a_attribute_definition_string_ref.ATTRIBUTE_DEFINITION_STRING_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_definition_xhtml_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs, stageSet.Stage.A_ATTRIBUTE_DEFINITION_XHTML_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_xhtml_refIdent := "__models" + a_attribute_definition_xhtml_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_XHTML_REF{Name: %s}).Stage(stageSet.Stage)", a_attribute_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_attribute_definition_xhtml_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_attribute_definition_xhtml_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_XHTML_REF = %s", a_attribute_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_attribute_definition_xhtml_ref.ATTRIBUTE_DEFINITION_XHTML_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_value_boolean := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_VALUE_BOOLEANs, stageSet.Stage.A_ATTRIBUTE_VALUE_BOOLEAN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_booleanIdent := "__models" + a_attribute_value_boolean.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_BOOLEAN{Name: %s}).Stage(stageSet.Stage)", a_attribute_value_booleanIdent, __gong__toRawStringLiteral(a_attribute_value_boolean.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_booleanIdent, __gong__toRawStringLiteral(a_attribute_value_boolean.Name)))
			for _, elem := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_BOOLEAN = append(%s.ATTRIBUTE_VALUE_BOOLEAN, %s)", a_attribute_value_booleanIdent, a_attribute_value_booleanIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_value_date := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_VALUE_DATEs, stageSet.Stage.A_ATTRIBUTE_VALUE_DATE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_dateIdent := "__models" + a_attribute_value_date.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_DATE{Name: %s}).Stage(stageSet.Stage)", a_attribute_value_dateIdent, __gong__toRawStringLiteral(a_attribute_value_date.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_dateIdent, __gong__toRawStringLiteral(a_attribute_value_date.Name)))
			for _, elem := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_DATE = append(%s.ATTRIBUTE_VALUE_DATE, %s)", a_attribute_value_dateIdent, a_attribute_value_dateIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_value_enumeration := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_VALUE_ENUMERATIONs, stageSet.Stage.A_ATTRIBUTE_VALUE_ENUMERATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_enumerationIdent := "__models" + a_attribute_value_enumeration.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_ENUMERATION{Name: %s}).Stage(stageSet.Stage)", a_attribute_value_enumerationIdent, __gong__toRawStringLiteral(a_attribute_value_enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_enumerationIdent, __gong__toRawStringLiteral(a_attribute_value_enumeration.Name)))
			for _, elem := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_ENUMERATION = append(%s.ATTRIBUTE_VALUE_ENUMERATION, %s)", a_attribute_value_enumerationIdent, a_attribute_value_enumerationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_value_integer := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_VALUE_INTEGERs, stageSet.Stage.A_ATTRIBUTE_VALUE_INTEGER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_integerIdent := "__models" + a_attribute_value_integer.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_INTEGER{Name: %s}).Stage(stageSet.Stage)", a_attribute_value_integerIdent, __gong__toRawStringLiteral(a_attribute_value_integer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_integerIdent, __gong__toRawStringLiteral(a_attribute_value_integer.Name)))
			for _, elem := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_INTEGER = append(%s.ATTRIBUTE_VALUE_INTEGER, %s)", a_attribute_value_integerIdent, a_attribute_value_integerIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_value_real := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_VALUE_REALs, stageSet.Stage.A_ATTRIBUTE_VALUE_REAL_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_realIdent := "__models" + a_attribute_value_real.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_REAL{Name: %s}).Stage(stageSet.Stage)", a_attribute_value_realIdent, __gong__toRawStringLiteral(a_attribute_value_real.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_realIdent, __gong__toRawStringLiteral(a_attribute_value_real.Name)))
			for _, elem := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_REAL = append(%s.ATTRIBUTE_VALUE_REAL, %s)", a_attribute_value_realIdent, a_attribute_value_realIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_value_string := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_VALUE_STRINGs, stageSet.Stage.A_ATTRIBUTE_VALUE_STRING_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_stringIdent := "__models" + a_attribute_value_string.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_STRING{Name: %s}).Stage(stageSet.Stage)", a_attribute_value_stringIdent, __gong__toRawStringLiteral(a_attribute_value_string.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_stringIdent, __gong__toRawStringLiteral(a_attribute_value_string.Name)))
			for _, elem := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_STRING = append(%s.ATTRIBUTE_VALUE_STRING, %s)", a_attribute_value_stringIdent, a_attribute_value_stringIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_value_xhtml := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_VALUE_XHTMLs, stageSet.Stage.A_ATTRIBUTE_VALUE_XHTML_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_xhtmlIdent := "__models" + a_attribute_value_xhtml.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_XHTML{Name: %s}).Stage(stageSet.Stage)", a_attribute_value_xhtmlIdent, __gong__toRawStringLiteral(a_attribute_value_xhtml.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_xhtmlIdent, __gong__toRawStringLiteral(a_attribute_value_xhtml.Name)))
			for _, elem := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_XHTML = append(%s.ATTRIBUTE_VALUE_XHTML, %s)", a_attribute_value_xhtmlIdent, a_attribute_value_xhtmlIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_attribute_value_xhtml_1 := range __gong__sortStageSetInstances(stageSet.Stage.A_ATTRIBUTE_VALUE_XHTML_1s, stageSet.Stage.A_ATTRIBUTE_VALUE_XHTML_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_xhtml_1Ident := "__models" + a_attribute_value_xhtml_1.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_XHTML_1{Name: %s}).Stage(stageSet.Stage)", a_attribute_value_xhtml_1Ident, __gong__toRawStringLiteral(a_attribute_value_xhtml_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_xhtml_1Ident, __gong__toRawStringLiteral(a_attribute_value_xhtml_1.Name)))
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_BOOLEAN = append(%s.ATTRIBUTE_VALUE_BOOLEAN, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_DATE = append(%s.ATTRIBUTE_VALUE_DATE, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_ENUMERATION = append(%s.ATTRIBUTE_VALUE_ENUMERATION, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_INTEGER = append(%s.ATTRIBUTE_VALUE_INTEGER, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_REAL = append(%s.ATTRIBUTE_VALUE_REAL, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_STRING = append(%s.ATTRIBUTE_VALUE_STRING, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_XHTML = append(%s.ATTRIBUTE_VALUE_XHTML, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_children := range __gong__sortStageSetInstances(stageSet.Stage.A_CHILDRENs, stageSet.Stage.A_CHILDREN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_childrenIdent := "__models" + a_children.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_CHILDREN{Name: %s}).Stage(stageSet.Stage)", a_childrenIdent, __gong__toRawStringLiteral(a_children.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_childrenIdent, __gong__toRawStringLiteral(a_children.Name)))
			for _, elem := range a_children.SPEC_HIERARCHY {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_HIERARCHY = append(%s.SPEC_HIERARCHY, %s)", a_childrenIdent, a_childrenIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_core_content := range __gong__sortStageSetInstances(stageSet.Stage.A_CORE_CONTENTs, stageSet.Stage.A_CORE_CONTENT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_core_contentIdent := "__models" + a_core_content.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_CORE_CONTENT{Name: %s}).Stage(stageSet.Stage)", a_core_contentIdent, __gong__toRawStringLiteral(a_core_content.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_core_contentIdent, __gong__toRawStringLiteral(a_core_content.Name)))
			if a_core_content.REQ_IF_CONTENT != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + a_core_content.REQ_IF_CONTENT.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_CONTENT = %s", a_core_contentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_datatypes := range __gong__sortStageSetInstances(stageSet.Stage.A_DATATYPESs, stageSet.Stage.A_DATATYPES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatypesIdent := "__models" + a_datatypes.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPES{Name: %s}).Stage(stageSet.Stage)", a_datatypesIdent, __gong__toRawStringLiteral(a_datatypes.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatypesIdent, __gong__toRawStringLiteral(a_datatypes.Name)))
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_BOOLEAN = append(%s.DATATYPE_DEFINITION_BOOLEAN, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_DATE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_DATE = append(%s.DATATYPE_DEFINITION_DATE, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_ENUMERATION = append(%s.DATATYPE_DEFINITION_ENUMERATION, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_INTEGER = append(%s.DATATYPE_DEFINITION_INTEGER, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_REAL {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_REAL = append(%s.DATATYPE_DEFINITION_REAL, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_STRING {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_STRING = append(%s.DATATYPE_DEFINITION_STRING, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_XHTML {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_XHTML = append(%s.DATATYPE_DEFINITION_XHTML, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_datatype_definition_boolean_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs, stageSet.Stage.A_DATATYPE_DEFINITION_BOOLEAN_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_boolean_refIdent := "__models" + a_datatype_definition_boolean_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_BOOLEAN_REF{Name: %s}).Stage(stageSet.Stage)", a_datatype_definition_boolean_refIdent, __gong__toRawStringLiteral(a_datatype_definition_boolean_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_boolean_refIdent, __gong__toRawStringLiteral(a_datatype_definition_boolean_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_BOOLEAN_REF = %s", a_datatype_definition_boolean_refIdent, __gong__toRawStringLiteral(a_datatype_definition_boolean_ref.DATATYPE_DEFINITION_BOOLEAN_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_datatype_definition_date_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_DATATYPE_DEFINITION_DATE_REFs, stageSet.Stage.A_DATATYPE_DEFINITION_DATE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_date_refIdent := "__models" + a_datatype_definition_date_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_DATE_REF{Name: %s}).Stage(stageSet.Stage)", a_datatype_definition_date_refIdent, __gong__toRawStringLiteral(a_datatype_definition_date_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_date_refIdent, __gong__toRawStringLiteral(a_datatype_definition_date_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_DATE_REF = %s", a_datatype_definition_date_refIdent, __gong__toRawStringLiteral(a_datatype_definition_date_ref.DATATYPE_DEFINITION_DATE_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_datatype_definition_enumeration_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs, stageSet.Stage.A_DATATYPE_DEFINITION_ENUMERATION_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_enumeration_refIdent := "__models" + a_datatype_definition_enumeration_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_ENUMERATION_REF{Name: %s}).Stage(stageSet.Stage)", a_datatype_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_datatype_definition_enumeration_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_datatype_definition_enumeration_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_ENUMERATION_REF = %s", a_datatype_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_datatype_definition_enumeration_ref.DATATYPE_DEFINITION_ENUMERATION_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_datatype_definition_integer_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_DATATYPE_DEFINITION_INTEGER_REFs, stageSet.Stage.A_DATATYPE_DEFINITION_INTEGER_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_integer_refIdent := "__models" + a_datatype_definition_integer_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_INTEGER_REF{Name: %s}).Stage(stageSet.Stage)", a_datatype_definition_integer_refIdent, __gong__toRawStringLiteral(a_datatype_definition_integer_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_integer_refIdent, __gong__toRawStringLiteral(a_datatype_definition_integer_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_INTEGER_REF = %s", a_datatype_definition_integer_refIdent, __gong__toRawStringLiteral(a_datatype_definition_integer_ref.DATATYPE_DEFINITION_INTEGER_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_datatype_definition_real_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_DATATYPE_DEFINITION_REAL_REFs, stageSet.Stage.A_DATATYPE_DEFINITION_REAL_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_real_refIdent := "__models" + a_datatype_definition_real_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_REAL_REF{Name: %s}).Stage(stageSet.Stage)", a_datatype_definition_real_refIdent, __gong__toRawStringLiteral(a_datatype_definition_real_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_real_refIdent, __gong__toRawStringLiteral(a_datatype_definition_real_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_REAL_REF = %s", a_datatype_definition_real_refIdent, __gong__toRawStringLiteral(a_datatype_definition_real_ref.DATATYPE_DEFINITION_REAL_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_datatype_definition_string_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_DATATYPE_DEFINITION_STRING_REFs, stageSet.Stage.A_DATATYPE_DEFINITION_STRING_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_string_refIdent := "__models" + a_datatype_definition_string_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_STRING_REF{Name: %s}).Stage(stageSet.Stage)", a_datatype_definition_string_refIdent, __gong__toRawStringLiteral(a_datatype_definition_string_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_string_refIdent, __gong__toRawStringLiteral(a_datatype_definition_string_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_STRING_REF = %s", a_datatype_definition_string_refIdent, __gong__toRawStringLiteral(a_datatype_definition_string_ref.DATATYPE_DEFINITION_STRING_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_datatype_definition_xhtml_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_DATATYPE_DEFINITION_XHTML_REFs, stageSet.Stage.A_DATATYPE_DEFINITION_XHTML_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_xhtml_refIdent := "__models" + a_datatype_definition_xhtml_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_XHTML_REF{Name: %s}).Stage(stageSet.Stage)", a_datatype_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_datatype_definition_xhtml_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_datatype_definition_xhtml_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_XHTML_REF = %s", a_datatype_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_datatype_definition_xhtml_ref.DATATYPE_DEFINITION_XHTML_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_editable_atts := range __gong__sortStageSetInstances(stageSet.Stage.A_EDITABLE_ATTSs, stageSet.Stage.A_EDITABLE_ATTS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_editable_attsIdent := "__models" + a_editable_atts.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_EDITABLE_ATTS{Name: %s}).Stage(stageSet.Stage)", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_BOOLEAN_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_BOOLEAN_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_DATE_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_DATE_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_ENUMERATION_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_ENUMERATION_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_INTEGER_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_INTEGER_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_REAL_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_REAL_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_STRING_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_STRING_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_XHTML_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_XHTML_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_enum_value_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_ENUM_VALUE_REFs, stageSet.Stage.A_ENUM_VALUE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_enum_value_refIdent := "__models" + a_enum_value_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ENUM_VALUE_REF{Name: %s}).Stage(stageSet.Stage)", a_enum_value_refIdent, __gong__toRawStringLiteral(a_enum_value_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_enum_value_refIdent, __gong__toRawStringLiteral(a_enum_value_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ENUM_VALUE_REF = %s", a_enum_value_refIdent, __gong__toRawStringLiteral(a_enum_value_ref.ENUM_VALUE_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_object := range __gong__sortStageSetInstances(stageSet.Stage.A_OBJECTs, stageSet.Stage.A_OBJECT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_objectIdent := "__models" + a_object.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_OBJECT{Name: %s}).Stage(stageSet.Stage)", a_objectIdent, __gong__toRawStringLiteral(a_object.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_objectIdent, __gong__toRawStringLiteral(a_object.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT_REF = %s", a_objectIdent, __gong__toRawStringLiteral(a_object.SPEC_OBJECT_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_properties := range __gong__sortStageSetInstances(stageSet.Stage.A_PROPERTIESs, stageSet.Stage.A_PROPERTIES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_propertiesIdent := "__models" + a_properties.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_PROPERTIES{Name: %s}).Stage(stageSet.Stage)", a_propertiesIdent, __gong__toRawStringLiteral(a_properties.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_propertiesIdent, __gong__toRawStringLiteral(a_properties.Name)))
			if a_properties.EMBEDDED_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + a_properties.EMBEDDED_VALUE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EMBEDDED_VALUE = %s", a_propertiesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_relation_group_type_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_RELATION_GROUP_TYPE_REFs, stageSet.Stage.A_RELATION_GROUP_TYPE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_relation_group_type_refIdent := "__models" + a_relation_group_type_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_RELATION_GROUP_TYPE_REF{Name: %s}).Stage(stageSet.Stage)", a_relation_group_type_refIdent, __gong__toRawStringLiteral(a_relation_group_type_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_relation_group_type_refIdent, __gong__toRawStringLiteral(a_relation_group_type_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.RELATION_GROUP_TYPE_REF = %s", a_relation_group_type_refIdent, __gong__toRawStringLiteral(a_relation_group_type_ref.RELATION_GROUP_TYPE_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_source_1 := range __gong__sortStageSetInstances(stageSet.Stage.A_SOURCE_1s, stageSet.Stage.A_SOURCE_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_source_1Ident := "__models" + a_source_1.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SOURCE_1{Name: %s}).Stage(stageSet.Stage)", a_source_1Ident, __gong__toRawStringLiteral(a_source_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_source_1Ident, __gong__toRawStringLiteral(a_source_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT_REF = %s", a_source_1Ident, __gong__toRawStringLiteral(a_source_1.SPEC_OBJECT_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_source_specification_1 := range __gong__sortStageSetInstances(stageSet.Stage.A_SOURCE_SPECIFICATION_1s, stageSet.Stage.A_SOURCE_SPECIFICATION_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_source_specification_1Ident := "__models" + a_source_specification_1.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SOURCE_SPECIFICATION_1{Name: %s}).Stage(stageSet.Stage)", a_source_specification_1Ident, __gong__toRawStringLiteral(a_source_specification_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_source_specification_1Ident, __gong__toRawStringLiteral(a_source_specification_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATION_REF = %s", a_source_specification_1Ident, __gong__toRawStringLiteral(string(a_source_specification_1.SPECIFICATION_REF))))
		}
	}
	if stageSet.Stage != nil {
		for _, a_specifications := range __gong__sortStageSetInstances(stageSet.Stage.A_SPECIFICATIONSs, stageSet.Stage.A_SPECIFICATIONS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_specificationsIdent := "__models" + a_specifications.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPECIFICATIONS{Name: %s}).Stage(stageSet.Stage)", a_specificationsIdent, __gong__toRawStringLiteral(a_specifications.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_specificationsIdent, __gong__toRawStringLiteral(a_specifications.Name)))
			for _, elem := range a_specifications.SPECIFICATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATION = append(%s.SPECIFICATION, %s)", a_specificationsIdent, a_specificationsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_specification_type_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_SPECIFICATION_TYPE_REFs, stageSet.Stage.A_SPECIFICATION_TYPE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_specification_type_refIdent := "__models" + a_specification_type_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPECIFICATION_TYPE_REF{Name: %s}).Stage(stageSet.Stage)", a_specification_type_refIdent, __gong__toRawStringLiteral(a_specification_type_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_specification_type_refIdent, __gong__toRawStringLiteral(a_specification_type_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATION_TYPE_REF = %s", a_specification_type_refIdent, __gong__toRawStringLiteral(a_specification_type_ref.SPECIFICATION_TYPE_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_specified_values := range __gong__sortStageSetInstances(stageSet.Stage.A_SPECIFIED_VALUESs, stageSet.Stage.A_SPECIFIED_VALUES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_specified_valuesIdent := "__models" + a_specified_values.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPECIFIED_VALUES{Name: %s}).Stage(stageSet.Stage)", a_specified_valuesIdent, __gong__toRawStringLiteral(a_specified_values.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_specified_valuesIdent, __gong__toRawStringLiteral(a_specified_values.Name)))
			for _, elem := range a_specified_values.ENUM_VALUE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ENUM_VALUE = append(%s.ENUM_VALUE, %s)", a_specified_valuesIdent, a_specified_valuesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_spec_attributes := range __gong__sortStageSetInstances(stageSet.Stage.A_SPEC_ATTRIBUTESs, stageSet.Stage.A_SPEC_ATTRIBUTES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_attributesIdent := "__models" + a_spec_attributes.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_ATTRIBUTES{Name: %s}).Stage(stageSet.Stage)", a_spec_attributesIdent, __gong__toRawStringLiteral(a_spec_attributes.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_attributesIdent, __gong__toRawStringLiteral(a_spec_attributes.Name)))
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_BOOLEAN = append(%s.ATTRIBUTE_DEFINITION_BOOLEAN, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_DATE = append(%s.ATTRIBUTE_DEFINITION_DATE, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_ENUMERATION = append(%s.ATTRIBUTE_DEFINITION_ENUMERATION, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_INTEGER = append(%s.ATTRIBUTE_DEFINITION_INTEGER, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_REAL = append(%s.ATTRIBUTE_DEFINITION_REAL, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_STRING = append(%s.ATTRIBUTE_DEFINITION_STRING, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_XHTML = append(%s.ATTRIBUTE_DEFINITION_XHTML, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_spec_objects := range __gong__sortStageSetInstances(stageSet.Stage.A_SPEC_OBJECTSs, stageSet.Stage.A_SPEC_OBJECTS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_objectsIdent := "__models" + a_spec_objects.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_OBJECTS{Name: %s}).Stage(stageSet.Stage)", a_spec_objectsIdent, __gong__toRawStringLiteral(a_spec_objects.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_objectsIdent, __gong__toRawStringLiteral(a_spec_objects.Name)))
			for _, elem := range a_spec_objects.SPEC_OBJECT {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT = append(%s.SPEC_OBJECT, %s)", a_spec_objectsIdent, a_spec_objectsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_spec_object_type_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_SPEC_OBJECT_TYPE_REFs, stageSet.Stage.A_SPEC_OBJECT_TYPE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_object_type_refIdent := "__models" + a_spec_object_type_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_OBJECT_TYPE_REF{Name: %s}).Stage(stageSet.Stage)", a_spec_object_type_refIdent, __gong__toRawStringLiteral(a_spec_object_type_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_object_type_refIdent, __gong__toRawStringLiteral(a_spec_object_type_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT_TYPE_REF = %s", a_spec_object_type_refIdent, __gong__toRawStringLiteral(a_spec_object_type_ref.SPEC_OBJECT_TYPE_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_spec_relations := range __gong__sortStageSetInstances(stageSet.Stage.A_SPEC_RELATIONSs, stageSet.Stage.A_SPEC_RELATIONS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_relationsIdent := "__models" + a_spec_relations.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_RELATIONS{Name: %s}).Stage(stageSet.Stage)", a_spec_relationsIdent, __gong__toRawStringLiteral(a_spec_relations.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_relationsIdent, __gong__toRawStringLiteral(a_spec_relations.Name)))
			for _, elem := range a_spec_relations.SPEC_RELATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION = append(%s.SPEC_RELATION, %s)", a_spec_relationsIdent, a_spec_relationsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_spec_relation_groups := range __gong__sortStageSetInstances(stageSet.Stage.A_SPEC_RELATION_GROUPSs, stageSet.Stage.A_SPEC_RELATION_GROUPS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_relation_groupsIdent := "__models" + a_spec_relation_groups.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_RELATION_GROUPS{Name: %s}).Stage(stageSet.Stage)", a_spec_relation_groupsIdent, __gong__toRawStringLiteral(a_spec_relation_groups.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_relation_groupsIdent, __gong__toRawStringLiteral(a_spec_relation_groups.Name)))
			for _, elem := range a_spec_relation_groups.RELATION_GROUP {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RELATION_GROUP = append(%s.RELATION_GROUP, %s)", a_spec_relation_groupsIdent, a_spec_relation_groupsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_spec_relation_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_SPEC_RELATION_REFs, stageSet.Stage.A_SPEC_RELATION_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_relation_refIdent := "__models" + a_spec_relation_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_RELATION_REF{Name: %s}).Stage(stageSet.Stage)", a_spec_relation_refIdent, __gong__toRawStringLiteral(a_spec_relation_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_relation_refIdent, __gong__toRawStringLiteral(a_spec_relation_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION_REF = %s", a_spec_relation_refIdent, __gong__toRawStringLiteral(a_spec_relation_ref.SPEC_RELATION_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_spec_relation_type_ref := range __gong__sortStageSetInstances(stageSet.Stage.A_SPEC_RELATION_TYPE_REFs, stageSet.Stage.A_SPEC_RELATION_TYPE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_relation_type_refIdent := "__models" + a_spec_relation_type_ref.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_RELATION_TYPE_REF{Name: %s}).Stage(stageSet.Stage)", a_spec_relation_type_refIdent, __gong__toRawStringLiteral(a_spec_relation_type_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_relation_type_refIdent, __gong__toRawStringLiteral(a_spec_relation_type_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION_TYPE_REF = %s", a_spec_relation_type_refIdent, __gong__toRawStringLiteral(a_spec_relation_type_ref.SPEC_RELATION_TYPE_REF)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_spec_types := range __gong__sortStageSetInstances(stageSet.Stage.A_SPEC_TYPESs, stageSet.Stage.A_SPEC_TYPES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_typesIdent := "__models" + a_spec_types.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_TYPES{Name: %s}).Stage(stageSet.Stage)", a_spec_typesIdent, __gong__toRawStringLiteral(a_spec_types.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_typesIdent, __gong__toRawStringLiteral(a_spec_types.Name)))
			for _, elem := range a_spec_types.RELATION_GROUP_TYPE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RELATION_GROUP_TYPE = append(%s.RELATION_GROUP_TYPE, %s)", a_spec_typesIdent, a_spec_typesIdent, targetIdent))
			}
			for _, elem := range a_spec_types.SPEC_OBJECT_TYPE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT_TYPE = append(%s.SPEC_OBJECT_TYPE, %s)", a_spec_typesIdent, a_spec_typesIdent, targetIdent))
			}
			for _, elem := range a_spec_types.SPEC_RELATION_TYPE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION_TYPE = append(%s.SPEC_RELATION_TYPE, %s)", a_spec_typesIdent, a_spec_typesIdent, targetIdent))
			}
			for _, elem := range a_spec_types.SPECIFICATION_TYPE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATION_TYPE = append(%s.SPECIFICATION_TYPE, %s)", a_spec_typesIdent, a_spec_typesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_the_header := range __gong__sortStageSetInstances(stageSet.Stage.A_THE_HEADERs, stageSet.Stage.A_THE_HEADER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_the_headerIdent := "__models" + a_the_header.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_THE_HEADER{Name: %s}).Stage(stageSet.Stage)", a_the_headerIdent, __gong__toRawStringLiteral(a_the_header.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_the_headerIdent, __gong__toRawStringLiteral(a_the_header.Name)))
			if a_the_header.REQ_IF_HEADER != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + a_the_header.REQ_IF_HEADER.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_HEADER = %s", a_the_headerIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_tool_extensions := range __gong__sortStageSetInstances(stageSet.Stage.A_TOOL_EXTENSIONSs, stageSet.Stage.A_TOOL_EXTENSIONS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_tool_extensionsIdent := "__models" + a_tool_extensions.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_TOOL_EXTENSIONS{Name: %s}).Stage(stageSet.Stage)", a_tool_extensionsIdent, __gong__toRawStringLiteral(a_tool_extensions.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_tool_extensionsIdent, __gong__toRawStringLiteral(a_tool_extensions.Name)))
			for _, elem := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_TOOL_EXTENSION = append(%s.REQ_IF_TOOL_EXTENSION, %s)", a_tool_extensionsIdent, a_tool_extensionsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, datatype_definition_boolean := range __gong__sortStageSetInstances(stageSet.Stage.DATATYPE_DEFINITION_BOOLEANs, stageSet.Stage.DATATYPE_DEFINITION_BOOLEAN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_booleanIdent := "__models" + datatype_definition_boolean.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_BOOLEAN{Name: %s}).Stage(stageSet.Stage)", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.LONG_NAME)))
			if datatype_definition_boolean.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datatype_definition_boolean.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_booleanIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, datatype_definition_date := range __gong__sortStageSetInstances(stageSet.Stage.DATATYPE_DEFINITION_DATEs, stageSet.Stage.DATATYPE_DEFINITION_DATE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_dateIdent := "__models" + datatype_definition_date.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_DATE{Name: %s}).Stage(stageSet.Stage)", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.LONG_NAME)))
			if datatype_definition_date.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datatype_definition_date.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_dateIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, datatype_definition_enumeration := range __gong__sortStageSetInstances(stageSet.Stage.DATATYPE_DEFINITION_ENUMERATIONs, stageSet.Stage.DATATYPE_DEFINITION_ENUMERATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_enumerationIdent := "__models" + datatype_definition_enumeration.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_ENUMERATION{Name: %s}).Stage(stageSet.Stage)", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.LONG_NAME)))
			if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datatype_definition_enumeration.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_enumerationIdent, targetIdent))
			}
			if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datatype_definition_enumeration.SPECIFIED_VALUES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPECIFIED_VALUES = %s", datatype_definition_enumerationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, datatype_definition_integer := range __gong__sortStageSetInstances(stageSet.Stage.DATATYPE_DEFINITION_INTEGERs, stageSet.Stage.DATATYPE_DEFINITION_INTEGER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_integerIdent := "__models" + datatype_definition_integer.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_INTEGER{Name: %s}).Stage(stageSet.Stage)", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.LONG_NAME)))
			values.WriteString(fmt.Sprintf("\n\t%s.MAX = %d", datatype_definition_integerIdent, datatype_definition_integer.MAX))
			values.WriteString(fmt.Sprintf("\n\t%s.MIN = %d", datatype_definition_integerIdent, datatype_definition_integer.MIN))
			if datatype_definition_integer.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datatype_definition_integer.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_integerIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, datatype_definition_real := range __gong__sortStageSetInstances(stageSet.Stage.DATATYPE_DEFINITION_REALs, stageSet.Stage.DATATYPE_DEFINITION_REAL_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_realIdent := "__models" + datatype_definition_real.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_REAL{Name: %s}).Stage(stageSet.Stage)", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ACCURACY = %d", datatype_definition_realIdent, datatype_definition_real.ACCURACY))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.LONG_NAME)))
			values.WriteString(fmt.Sprintf("\n\t%s.MAX = %f", datatype_definition_realIdent, datatype_definition_real.MAX))
			values.WriteString(fmt.Sprintf("\n\t%s.MIN = %f", datatype_definition_realIdent, datatype_definition_real.MIN))
			if datatype_definition_real.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datatype_definition_real.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_realIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, datatype_definition_string := range __gong__sortStageSetInstances(stageSet.Stage.DATATYPE_DEFINITION_STRINGs, stageSet.Stage.DATATYPE_DEFINITION_STRING_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_stringIdent := "__models" + datatype_definition_string.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_STRING{Name: %s}).Stage(stageSet.Stage)", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.LONG_NAME)))
			values.WriteString(fmt.Sprintf("\n\t%s.MAX_LENGTH = %d", datatype_definition_stringIdent, datatype_definition_string.MAX_LENGTH))
			if datatype_definition_string.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datatype_definition_string.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_stringIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, datatype_definition_xhtml := range __gong__sortStageSetInstances(stageSet.Stage.DATATYPE_DEFINITION_XHTMLs, stageSet.Stage.DATATYPE_DEFINITION_XHTML_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_xhtmlIdent := "__models" + datatype_definition_xhtml.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_XHTML{Name: %s}).Stage(stageSet.Stage)", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.LONG_NAME)))
			if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + datatype_definition_xhtml.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_xhtmlIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, embedded_value := range __gong__sortStageSetInstances(stageSet.Stage.EMBEDDED_VALUEs, stageSet.Stage.EMBEDDED_VALUE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			embedded_valueIdent := "__models" + embedded_value.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EMBEDDED_VALUE{Name: %s}).Stage(stageSet.Stage)", embedded_valueIdent, __gong__toRawStringLiteral(embedded_value.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", embedded_valueIdent, __gong__toRawStringLiteral(embedded_value.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.KEY = %d", embedded_valueIdent, embedded_value.KEY))
			values.WriteString(fmt.Sprintf("\n\t%s.OTHER_CONTENT = %s", embedded_valueIdent, __gong__toRawStringLiteral(embedded_value.OTHER_CONTENT)))
		}
	}
	if stageSet.Stage != nil {
		for _, enum_value := range __gong__sortStageSetInstances(stageSet.Stage.ENUM_VALUEs, stageSet.Stage.ENUM_VALUE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			enum_valueIdent := "__models" + enum_value.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ENUM_VALUE{Name: %s}).Stage(stageSet.Stage)", enum_valueIdent, __gong__toRawStringLiteral(enum_value.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.LONG_NAME)))
			if enum_value.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + enum_value.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", enum_valueIdent, targetIdent))
			}
			if enum_value.PROPERTIES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + enum_value.PROPERTIES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PROPERTIES = %s", enum_valueIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, relation_group := range __gong__sortStageSetInstances(stageSet.Stage.RELATION_GROUPs, stageSet.Stage.RELATION_GROUP_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			relation_groupIdent := "__models" + relation_group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RELATION_GROUP{Name: %s}).Stage(stageSet.Stage)", relation_groupIdent, __gong__toRawStringLiteral(relation_group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.LONG_NAME)))
			if relation_group.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + relation_group.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", relation_groupIdent, targetIdent))
			}
			if relation_group.SOURCE_SPECIFICATION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + relation_group.SOURCE_SPECIFICATION.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SOURCE_SPECIFICATION = %s", relation_groupIdent, targetIdent))
			}
			if relation_group.SPEC_RELATIONS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + relation_group.SPEC_RELATIONS.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATIONS = %s", relation_groupIdent, targetIdent))
			}
			if relation_group.TARGET_SPECIFICATION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + relation_group.TARGET_SPECIFICATION.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TARGET_SPECIFICATION = %s", relation_groupIdent, targetIdent))
			}
			if relation_group.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + relation_group.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", relation_groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, relation_group_type := range __gong__sortStageSetInstances(stageSet.Stage.RELATION_GROUP_TYPEs, stageSet.Stage.RELATION_GROUP_TYPE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			relation_group_typeIdent := "__models" + relation_group_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RELATION_GROUP_TYPE{Name: %s}).Stage(stageSet.Stage)", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.LONG_NAME)))
			if relation_group_type.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + relation_group_type.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", relation_group_typeIdent, targetIdent))
			}
			if relation_group_type.SPEC_ATTRIBUTES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + relation_group_type.SPEC_ATTRIBUTES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_ATTRIBUTES = %s", relation_group_typeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, req_if := range __gong__sortStageSetInstances(stageSet.Stage.REQ_IFs, stageSet.Stage.REQ_IF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			req_ifIdent := "__models" + req_if.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.REQ_IF{Name: %s}).Stage(stageSet.Stage)", req_ifIdent, __gong__toRawStringLiteral(req_if.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", req_ifIdent, __gong__toRawStringLiteral(req_if.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", req_ifIdent, __gong__toRawStringLiteral(req_if.Lang)))
			if req_if.THE_HEADER != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + req_if.THE_HEADER.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.THE_HEADER = %s", req_ifIdent, targetIdent))
			}
			if req_if.CORE_CONTENT != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + req_if.CORE_CONTENT.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CORE_CONTENT = %s", req_ifIdent, targetIdent))
			}
			if req_if.TOOL_EXTENSIONS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + req_if.TOOL_EXTENSIONS.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TOOL_EXTENSIONS = %s", req_ifIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, req_if_content := range __gong__sortStageSetInstances(stageSet.Stage.REQ_IF_CONTENTs, stageSet.Stage.REQ_IF_CONTENT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			req_if_contentIdent := "__models" + req_if_content.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.REQ_IF_CONTENT{Name: %s}).Stage(stageSet.Stage)", req_if_contentIdent, __gong__toRawStringLiteral(req_if_content.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", req_if_contentIdent, __gong__toRawStringLiteral(req_if_content.Name)))
			if req_if_content.DATATYPES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + req_if_content.DATATYPES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPES = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPEC_TYPES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + req_if_content.SPEC_TYPES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_TYPES = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPEC_OBJECTS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + req_if_content.SPEC_OBJECTS.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECTS = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPEC_RELATIONS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + req_if_content.SPEC_RELATIONS.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATIONS = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPECIFICATIONS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + req_if_content.SPECIFICATIONS.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATIONS = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPEC_RELATION_GROUPS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + req_if_content.SPEC_RELATION_GROUPS.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION_GROUPS = %s", req_if_contentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, req_if_header := range __gong__sortStageSetInstances(stageSet.Stage.REQ_IF_HEADERs, stageSet.Stage.REQ_IF_HEADER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			req_if_headerIdent := "__models" + req_if_header.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.REQ_IF_HEADER{Name: %s}).Stage(stageSet.Stage)", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.COMMENT = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.COMMENT)))
			values.WriteString(fmt.Sprintf("\n\t%s.CREATION_TIME = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.CREATION_TIME)))
			values.WriteString(fmt.Sprintf("\n\t%s.REPOSITORY_ID = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.REPOSITORY_ID)))
			values.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_TOOL_ID = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.REQ_IF_TOOL_ID)))
			values.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_VERSION = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.REQ_IF_VERSION)))
			values.WriteString(fmt.Sprintf("\n\t%s.SOURCE_TOOL_ID = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.SOURCE_TOOL_ID)))
			values.WriteString(fmt.Sprintf("\n\t%s.TITLE = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.TITLE)))
		}
	}
	if stageSet.Stage != nil {
		for _, req_if_tool_extension := range __gong__sortStageSetInstances(stageSet.Stage.REQ_IF_TOOL_EXTENSIONs, stageSet.Stage.REQ_IF_TOOL_EXTENSION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			req_if_tool_extensionIdent := "__models" + req_if_tool_extension.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.REQ_IF_TOOL_EXTENSION{Name: %s}).Stage(stageSet.Stage)", req_if_tool_extensionIdent, __gong__toRawStringLiteral(req_if_tool_extension.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", req_if_tool_extensionIdent, __gong__toRawStringLiteral(req_if_tool_extension.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, specification := range __gong__sortStageSetInstances(stageSet.Stage.SPECIFICATIONs, stageSet.Stage.SPECIFICATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specificationIdent := "__models" + specification.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPECIFICATION{Name: %s}).Stage(stageSet.Stage)", specificationIdent, __gong__toRawStringLiteral(specification.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specificationIdent, __gong__toRawStringLiteral(specification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", specificationIdent, __gong__toRawStringLiteral(specification.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", specificationIdent, __gong__toRawStringLiteral(specification.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", specificationIdent, __gong__toRawStringLiteral(specification.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", specificationIdent, __gong__toRawStringLiteral(specification.LONG_NAME)))
			if specification.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + specification.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", specificationIdent, targetIdent))
			}
			if specification.CHILDREN != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + specification.CHILDREN.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CHILDREN = %s", specificationIdent, targetIdent))
			}
			if specification.VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + specification.VALUES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VALUES = %s", specificationIdent, targetIdent))
			}
			if specification.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + specification.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", specificationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, specification_type := range __gong__sortStageSetInstances(stageSet.Stage.SPECIFICATION_TYPEs, stageSet.Stage.SPECIFICATION_TYPE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specification_typeIdent := "__models" + specification_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPECIFICATION_TYPE{Name: %s}).Stage(stageSet.Stage)", specification_typeIdent, __gong__toRawStringLiteral(specification_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.LONG_NAME)))
			if specification_type.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + specification_type.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", specification_typeIdent, targetIdent))
			}
			if specification_type.SPEC_ATTRIBUTES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + specification_type.SPEC_ATTRIBUTES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_ATTRIBUTES = %s", specification_typeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, spec_hierarchy := range __gong__sortStageSetInstances(stageSet.Stage.SPEC_HIERARCHYs, stageSet.Stage.SPEC_HIERARCHY_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_hierarchyIdent := "__models" + spec_hierarchy.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_HIERARCHY{Name: %s}).Stage(stageSet.Stage)", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", spec_hierarchyIdent, spec_hierarchy.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_TABLE_INTERNAL = %t", spec_hierarchyIdent, spec_hierarchy.IS_TABLE_INTERNAL))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.LONG_NAME)))
			if spec_hierarchy.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_hierarchy.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_hierarchyIdent, targetIdent))
			}
			if spec_hierarchy.CHILDREN != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_hierarchy.CHILDREN.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CHILDREN = %s", spec_hierarchyIdent, targetIdent))
			}
			if spec_hierarchy.EDITABLE_ATTS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_hierarchy.EDITABLE_ATTS.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EDITABLE_ATTS = %s", spec_hierarchyIdent, targetIdent))
			}
			if spec_hierarchy.OBJECT != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_hierarchy.OBJECT.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OBJECT = %s", spec_hierarchyIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, spec_object := range __gong__sortStageSetInstances(stageSet.Stage.SPEC_OBJECTs, stageSet.Stage.SPEC_OBJECT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_objectIdent := "__models" + spec_object.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_OBJECT{Name: %s}).Stage(stageSet.Stage)", spec_objectIdent, __gong__toRawStringLiteral(spec_object.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.LONG_NAME)))
			if spec_object.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_object.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_objectIdent, targetIdent))
			}
			if spec_object.VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_object.VALUES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VALUES = %s", spec_objectIdent, targetIdent))
			}
			if spec_object.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_object.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", spec_objectIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, spec_object_type := range __gong__sortStageSetInstances(stageSet.Stage.SPEC_OBJECT_TYPEs, stageSet.Stage.SPEC_OBJECT_TYPE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_object_typeIdent := "__models" + spec_object_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_OBJECT_TYPE{Name: %s}).Stage(stageSet.Stage)", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.LONG_NAME)))
			if spec_object_type.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_object_type.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_object_typeIdent, targetIdent))
			}
			if spec_object_type.SPEC_ATTRIBUTES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_object_type.SPEC_ATTRIBUTES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_ATTRIBUTES = %s", spec_object_typeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, spec_relation := range __gong__sortStageSetInstances(stageSet.Stage.SPEC_RELATIONs, stageSet.Stage.SPEC_RELATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_relationIdent := "__models" + spec_relation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_RELATION{Name: %s}).Stage(stageSet.Stage)", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.LONG_NAME)))
			if spec_relation.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_relation.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_relationIdent, targetIdent))
			}
			if spec_relation.VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_relation.VALUES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VALUES = %s", spec_relationIdent, targetIdent))
			}
			if spec_relation.SOURCE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_relation.SOURCE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SOURCE = %s", spec_relationIdent, targetIdent))
			}
			if spec_relation.TARGET != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_relation.TARGET.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TARGET = %s", spec_relationIdent, targetIdent))
			}
			if spec_relation.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_relation.TYPE.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", spec_relationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, spec_relation_type := range __gong__sortStageSetInstances(stageSet.Stage.SPEC_RELATION_TYPEs, stageSet.Stage.SPEC_RELATION_TYPE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_relation_typeIdent := "__models" + spec_relation_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_RELATION_TYPE{Name: %s}).Stage(stageSet.Stage)", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.LONG_NAME)))
			if spec_relation_type.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_relation_type.ALTERNATIVE_ID.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_relation_typeIdent, targetIdent))
			}
			if spec_relation_type.SPEC_ATTRIBUTES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + spec_relation_type.SPEC_ATTRIBUTES.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_ATTRIBUTES = %s", spec_relation_typeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, xhtml_content := range __gong__sortStageSetInstances(stageSet.Stage.XHTML_CONTENTs, stageSet.Stage.XHTML_CONTENT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			xhtml_contentIdent := "__models" + xhtml_content.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.XHTML_CONTENT{Name: %s}).Stage(stageSet.Stage)", xhtml_contentIdent, __gong__toRawStringLiteral(xhtml_content.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xhtml_contentIdent, __gong__toRawStringLiteral(xhtml_content.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", xhtml_contentIdent, __gong__toRawStringLiteral(xhtml_content.EnclosedText)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/reqif/go/models"
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
		case "github.com/fullstack-lang/gong/app/xsd/tests/reqif/go/models":
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
				case "ALTERNATIVE_ID":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ALTERNATIVE_ID), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_BOOLEAN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_BOOLEAN), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_DATE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_DATE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_ENUMERATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_ENUMERATION), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_INTEGER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_INTEGER), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_REAL":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_REAL), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_STRING":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_STRING), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_DEFINITION_XHTML":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_DEFINITION_XHTML), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_BOOLEAN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_BOOLEAN), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_DATE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_DATE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_ENUMERATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_ENUMERATION), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_INTEGER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_INTEGER), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_REAL":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_REAL), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_STRING":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_STRING), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ATTRIBUTE_VALUE_XHTML":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ATTRIBUTE_VALUE_XHTML), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ALTERNATIVE_ID":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ALTERNATIVE_ID), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_DATE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_DATE_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_INTEGER_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_REAL_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_REAL_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_STRING_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_STRING_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_DEFINITION_XHTML_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_BOOLEAN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_BOOLEAN), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_DATE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_DATE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_ENUMERATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_ENUMERATION), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_INTEGER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_INTEGER), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_REAL":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_REAL), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_STRING":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_STRING), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_XHTML":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_XHTML), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ATTRIBUTE_VALUE_XHTML_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ATTRIBUTE_VALUE_XHTML_1), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_CHILDREN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_CHILDREN), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_CORE_CONTENT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_CORE_CONTENT), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPES), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_BOOLEAN_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_DATE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_DATE_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_ENUMERATION_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_INTEGER_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_INTEGER_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_REAL_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_REAL_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_STRING_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_STRING_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_DATATYPE_DEFINITION_XHTML_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_DATATYPE_DEFINITION_XHTML_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_EDITABLE_ATTS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_EDITABLE_ATTS), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_ENUM_VALUE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_ENUM_VALUE_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_OBJECT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_OBJECT), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_PROPERTIES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_PROPERTIES), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_RELATION_GROUP_TYPE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_RELATION_GROUP_TYPE_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SOURCE_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SOURCE_1), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SOURCE_SPECIFICATION_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SOURCE_SPECIFICATION_1), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPECIFICATIONS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPECIFICATIONS), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPECIFICATION_TYPE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPECIFICATION_TYPE_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPECIFIED_VALUES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPECIFIED_VALUES), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_ATTRIBUTES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_ATTRIBUTES), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_OBJECTS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_OBJECTS), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_OBJECT_TYPE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_OBJECT_TYPE_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_RELATIONS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_RELATIONS), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_RELATION_GROUPS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_RELATION_GROUPS), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_RELATION_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_RELATION_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_RELATION_TYPE_REF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_RELATION_TYPE_REF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_SPEC_TYPES":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_SPEC_TYPES), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_THE_HEADER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_THE_HEADER), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_TOOL_EXTENSIONS":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_TOOL_EXTENSIONS), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_BOOLEAN":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_BOOLEAN), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_DATE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_DATE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_ENUMERATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_ENUMERATION), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_INTEGER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_INTEGER), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_REAL":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_REAL), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_STRING":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_STRING), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DATATYPE_DEFINITION_XHTML":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DATATYPE_DEFINITION_XHTML), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EMBEDDED_VALUE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EMBEDDED_VALUE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ENUM_VALUE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ENUM_VALUE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RELATION_GROUP":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RELATION_GROUP), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RELATION_GROUP_TYPE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RELATION_GROUP_TYPE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "REQ_IF":
					identifierMap[ident.Name] = __gong__stageSetInit(new(REQ_IF), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "REQ_IF_CONTENT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(REQ_IF_CONTENT), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "REQ_IF_HEADER":
					identifierMap[ident.Name] = __gong__stageSetInit(new(REQ_IF_HEADER), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "REQ_IF_TOOL_EXTENSION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(REQ_IF_TOOL_EXTENSION), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SPECIFICATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPECIFICATION), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SPECIFICATION_TYPE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPECIFICATION_TYPE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_HIERARCHY":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_HIERARCHY), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_OBJECT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_OBJECT), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_OBJECT_TYPE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_OBJECT_TYPE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_RELATION":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_RELATION), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SPEC_RELATION_TYPE":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SPEC_RELATION_TYPE), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "XHTML_CONTENT":
					identifierMap[ident.Name] = __gong__stageSetInit(new(XHTML_CONTENT), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *ATTRIBUTE_VALUE_BOOLEAN:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractBool(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					}
				case *ATTRIBUTE_VALUE_DATE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractString(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
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
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractInt(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					}
				case *ATTRIBUTE_VALUE_REAL:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractFloat(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					}
				case *ATTRIBUTE_VALUE_STRING:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractString(rhs)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
					}
				case *ATTRIBUTE_VALUE_XHTML:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IS_SIMPLIFIED":
						inst.IS_SIMPLIFIED = GongExtractBool(rhs)
					case "THE_VALUE":
						__gong__assignPointer(&inst.THE_VALUE, rhs, identifierMap)
					case "THE_ORIGINAL_VALUE":
						__gong__assignPointer(&inst.THE_ORIGINAL_VALUE, rhs, identifierMap)
					case "DEFINITION":
						__gong__assignPointer(&inst.DEFINITION, rhs, identifierMap)
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
						inst.SPECIFICATION_REF = Enum_GLOBAL_REF(GongExtractString(rhs))
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
					case "CHILDREN":
						__gong__assignPointer(&inst.CHILDREN, rhs, identifierMap)
					case "VALUES":
						__gong__assignPointer(&inst.VALUES, rhs, identifierMap)
					case "TYPE":
						__gong__assignPointer(&inst.TYPE, rhs, identifierMap)
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
					case "CHILDREN":
						__gong__assignPointer(&inst.CHILDREN, rhs, identifierMap)
					case "EDITABLE_ATTS":
						__gong__assignPointer(&inst.EDITABLE_ATTS, rhs, identifierMap)
					case "OBJECT":
						__gong__assignPointer(&inst.OBJECT, rhs, identifierMap)
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
				case *XHTML_CONTENT:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
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
