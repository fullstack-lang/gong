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
		alternative_idOrdered := []*ALTERNATIVE_ID{}
		for alternative_id := range stageSet.Stage.ALTERNATIVE_IDs {
			alternative_idOrdered = append(alternative_idOrdered, alternative_id)
		}
		sort.Slice(alternative_idOrdered, func(i, j int) bool {
			return stageSet.Stage.ALTERNATIVE_ID_stagedOrder[alternative_idOrdered[i]] < stageSet.Stage.ALTERNATIVE_ID_stagedOrder[alternative_idOrdered[j]]
		})
		for _, alternative_id := range alternative_idOrdered {
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
		attribute_definition_booleanOrdered := []*ATTRIBUTE_DEFINITION_BOOLEAN{}
		for attribute_definition_boolean := range stageSet.Stage.ATTRIBUTE_DEFINITION_BOOLEANs {
			attribute_definition_booleanOrdered = append(attribute_definition_booleanOrdered, attribute_definition_boolean)
		}
		sort.Slice(attribute_definition_booleanOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_BOOLEAN_stagedOrder[attribute_definition_booleanOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_BOOLEAN_stagedOrder[attribute_definition_booleanOrdered[j]]
		})
		for _, attribute_definition_boolean := range attribute_definition_booleanOrdered {
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
		attribute_definition_boolean_renderingOrdered := []*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering{}
		for attribute_definition_boolean_rendering := range stageSet.Stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings {
			attribute_definition_boolean_renderingOrdered = append(attribute_definition_boolean_renderingOrdered, attribute_definition_boolean_rendering)
		}
		sort.Slice(attribute_definition_boolean_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_stagedOrder[attribute_definition_boolean_renderingOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_stagedOrder[attribute_definition_boolean_renderingOrdered[j]]
		})
		for _, attribute_definition_boolean_rendering := range attribute_definition_boolean_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_boolean_renderingIdent := "__models" + attribute_definition_boolean_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering{Name: %s}).Stage(stageSet.Stage)", attribute_definition_boolean_renderingIdent, __gong__toRawStringLiteral(attribute_definition_boolean_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_boolean_renderingIdent, __gong__toRawStringLiteral(attribute_definition_boolean_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_boolean_renderingIdent, attribute_definition_boolean_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_boolean_renderingIdent, attribute_definition_boolean_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_boolean_renderingIdent, attribute_definition_boolean_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_boolean_renderingIdent, attribute_definition_boolean_rendering.Rank))
		}
	}
	if stageSet.Stage != nil {
		attribute_definition_dateOrdered := []*ATTRIBUTE_DEFINITION_DATE{}
		for attribute_definition_date := range stageSet.Stage.ATTRIBUTE_DEFINITION_DATEs {
			attribute_definition_dateOrdered = append(attribute_definition_dateOrdered, attribute_definition_date)
		}
		sort.Slice(attribute_definition_dateOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_DATE_stagedOrder[attribute_definition_dateOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_DATE_stagedOrder[attribute_definition_dateOrdered[j]]
		})
		for _, attribute_definition_date := range attribute_definition_dateOrdered {
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
		attribute_definition_date_renderingOrdered := []*ATTRIBUTE_DEFINITION_DATE_Rendering{}
		for attribute_definition_date_rendering := range stageSet.Stage.ATTRIBUTE_DEFINITION_DATE_Renderings {
			attribute_definition_date_renderingOrdered = append(attribute_definition_date_renderingOrdered, attribute_definition_date_rendering)
		}
		sort.Slice(attribute_definition_date_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_DATE_Rendering_stagedOrder[attribute_definition_date_renderingOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_DATE_Rendering_stagedOrder[attribute_definition_date_renderingOrdered[j]]
		})
		for _, attribute_definition_date_rendering := range attribute_definition_date_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_date_renderingIdent := "__models" + attribute_definition_date_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_DATE_Rendering{Name: %s}).Stage(stageSet.Stage)", attribute_definition_date_renderingIdent, __gong__toRawStringLiteral(attribute_definition_date_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_date_renderingIdent, __gong__toRawStringLiteral(attribute_definition_date_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_date_renderingIdent, attribute_definition_date_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_date_renderingIdent, attribute_definition_date_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_date_renderingIdent, attribute_definition_date_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_date_renderingIdent, attribute_definition_date_rendering.Rank))
		}
	}
	if stageSet.Stage != nil {
		attribute_definition_enumerationOrdered := []*ATTRIBUTE_DEFINITION_ENUMERATION{}
		for attribute_definition_enumeration := range stageSet.Stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
			attribute_definition_enumerationOrdered = append(attribute_definition_enumerationOrdered, attribute_definition_enumeration)
		}
		sort.Slice(attribute_definition_enumerationOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_ENUMERATION_stagedOrder[attribute_definition_enumerationOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_ENUMERATION_stagedOrder[attribute_definition_enumerationOrdered[j]]
		})
		for _, attribute_definition_enumeration := range attribute_definition_enumerationOrdered {
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
		attribute_definition_enumeration_renderingOrdered := []*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering{}
		for attribute_definition_enumeration_rendering := range stageSet.Stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings {
			attribute_definition_enumeration_renderingOrdered = append(attribute_definition_enumeration_renderingOrdered, attribute_definition_enumeration_rendering)
		}
		sort.Slice(attribute_definition_enumeration_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_stagedOrder[attribute_definition_enumeration_renderingOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_stagedOrder[attribute_definition_enumeration_renderingOrdered[j]]
		})
		for _, attribute_definition_enumeration_rendering := range attribute_definition_enumeration_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_enumeration_renderingIdent := "__models" + attribute_definition_enumeration_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering{Name: %s}).Stage(stageSet.Stage)", attribute_definition_enumeration_renderingIdent, __gong__toRawStringLiteral(attribute_definition_enumeration_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_enumeration_renderingIdent, __gong__toRawStringLiteral(attribute_definition_enumeration_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_enumeration_renderingIdent, attribute_definition_enumeration_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_enumeration_renderingIdent, attribute_definition_enumeration_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_enumeration_renderingIdent, attribute_definition_enumeration_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_enumeration_renderingIdent, attribute_definition_enumeration_rendering.Rank))
		}
	}
	if stageSet.Stage != nil {
		attribute_definition_integerOrdered := []*ATTRIBUTE_DEFINITION_INTEGER{}
		for attribute_definition_integer := range stageSet.Stage.ATTRIBUTE_DEFINITION_INTEGERs {
			attribute_definition_integerOrdered = append(attribute_definition_integerOrdered, attribute_definition_integer)
		}
		sort.Slice(attribute_definition_integerOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_INTEGER_stagedOrder[attribute_definition_integerOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_INTEGER_stagedOrder[attribute_definition_integerOrdered[j]]
		})
		for _, attribute_definition_integer := range attribute_definition_integerOrdered {
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
		attribute_definition_integer_renderingOrdered := []*ATTRIBUTE_DEFINITION_INTEGER_Rendering{}
		for attribute_definition_integer_rendering := range stageSet.Stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings {
			attribute_definition_integer_renderingOrdered = append(attribute_definition_integer_renderingOrdered, attribute_definition_integer_rendering)
		}
		sort.Slice(attribute_definition_integer_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_INTEGER_Rendering_stagedOrder[attribute_definition_integer_renderingOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_INTEGER_Rendering_stagedOrder[attribute_definition_integer_renderingOrdered[j]]
		})
		for _, attribute_definition_integer_rendering := range attribute_definition_integer_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_integer_renderingIdent := "__models" + attribute_definition_integer_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_INTEGER_Rendering{Name: %s}).Stage(stageSet.Stage)", attribute_definition_integer_renderingIdent, __gong__toRawStringLiteral(attribute_definition_integer_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_integer_renderingIdent, __gong__toRawStringLiteral(attribute_definition_integer_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_integer_renderingIdent, attribute_definition_integer_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_integer_renderingIdent, attribute_definition_integer_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_integer_renderingIdent, attribute_definition_integer_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_integer_renderingIdent, attribute_definition_integer_rendering.Rank))
		}
	}
	if stageSet.Stage != nil {
		attribute_definition_realOrdered := []*ATTRIBUTE_DEFINITION_REAL{}
		for attribute_definition_real := range stageSet.Stage.ATTRIBUTE_DEFINITION_REALs {
			attribute_definition_realOrdered = append(attribute_definition_realOrdered, attribute_definition_real)
		}
		sort.Slice(attribute_definition_realOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_REAL_stagedOrder[attribute_definition_realOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_REAL_stagedOrder[attribute_definition_realOrdered[j]]
		})
		for _, attribute_definition_real := range attribute_definition_realOrdered {
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
		attribute_definition_real_renderingOrdered := []*ATTRIBUTE_DEFINITION_REAL_Rendering{}
		for attribute_definition_real_rendering := range stageSet.Stage.ATTRIBUTE_DEFINITION_REAL_Renderings {
			attribute_definition_real_renderingOrdered = append(attribute_definition_real_renderingOrdered, attribute_definition_real_rendering)
		}
		sort.Slice(attribute_definition_real_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_REAL_Rendering_stagedOrder[attribute_definition_real_renderingOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_REAL_Rendering_stagedOrder[attribute_definition_real_renderingOrdered[j]]
		})
		for _, attribute_definition_real_rendering := range attribute_definition_real_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_real_renderingIdent := "__models" + attribute_definition_real_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_REAL_Rendering{Name: %s}).Stage(stageSet.Stage)", attribute_definition_real_renderingIdent, __gong__toRawStringLiteral(attribute_definition_real_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_real_renderingIdent, __gong__toRawStringLiteral(attribute_definition_real_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_real_renderingIdent, attribute_definition_real_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_real_renderingIdent, attribute_definition_real_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_real_renderingIdent, attribute_definition_real_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_real_renderingIdent, attribute_definition_real_rendering.Rank))
		}
	}
	if stageSet.Stage != nil {
		attribute_definition_renderingOrdered := []*ATTRIBUTE_DEFINITION_Rendering{}
		for attribute_definition_rendering := range stageSet.Stage.ATTRIBUTE_DEFINITION_Renderings {
			attribute_definition_renderingOrdered = append(attribute_definition_renderingOrdered, attribute_definition_rendering)
		}
		sort.Slice(attribute_definition_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_Rendering_stagedOrder[attribute_definition_renderingOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_Rendering_stagedOrder[attribute_definition_renderingOrdered[j]]
		})
		for _, attribute_definition_rendering := range attribute_definition_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_renderingIdent := "__models" + attribute_definition_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_Rendering{Name: %s}).Stage(stageSet.Stage)", attribute_definition_renderingIdent, __gong__toRawStringLiteral(attribute_definition_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_renderingIdent, __gong__toRawStringLiteral(attribute_definition_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_renderingIdent, attribute_definition_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_renderingIdent, attribute_definition_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_renderingIdent, attribute_definition_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_renderingIdent, attribute_definition_rendering.Rank))
		}
	}
	if stageSet.Stage != nil {
		attribute_definition_stringOrdered := []*ATTRIBUTE_DEFINITION_STRING{}
		for attribute_definition_string := range stageSet.Stage.ATTRIBUTE_DEFINITION_STRINGs {
			attribute_definition_stringOrdered = append(attribute_definition_stringOrdered, attribute_definition_string)
		}
		sort.Slice(attribute_definition_stringOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_STRING_stagedOrder[attribute_definition_stringOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_STRING_stagedOrder[attribute_definition_stringOrdered[j]]
		})
		for _, attribute_definition_string := range attribute_definition_stringOrdered {
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
		attribute_definition_string_renderingOrdered := []*ATTRIBUTE_DEFINITION_STRING_Rendering{}
		for attribute_definition_string_rendering := range stageSet.Stage.ATTRIBUTE_DEFINITION_STRING_Renderings {
			attribute_definition_string_renderingOrdered = append(attribute_definition_string_renderingOrdered, attribute_definition_string_rendering)
		}
		sort.Slice(attribute_definition_string_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_STRING_Rendering_stagedOrder[attribute_definition_string_renderingOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_STRING_Rendering_stagedOrder[attribute_definition_string_renderingOrdered[j]]
		})
		for _, attribute_definition_string_rendering := range attribute_definition_string_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_string_renderingIdent := "__models" + attribute_definition_string_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{Name: %s}).Stage(stageSet.Stage)", attribute_definition_string_renderingIdent, __gong__toRawStringLiteral(attribute_definition_string_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_string_renderingIdent, __gong__toRawStringLiteral(attribute_definition_string_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_string_renderingIdent, attribute_definition_string_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_string_renderingIdent, attribute_definition_string_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_string_renderingIdent, attribute_definition_string_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_string_renderingIdent, attribute_definition_string_rendering.Rank))
		}
	}
	if stageSet.Stage != nil {
		attribute_definition_xhtmlOrdered := []*ATTRIBUTE_DEFINITION_XHTML{}
		for attribute_definition_xhtml := range stageSet.Stage.ATTRIBUTE_DEFINITION_XHTMLs {
			attribute_definition_xhtmlOrdered = append(attribute_definition_xhtmlOrdered, attribute_definition_xhtml)
		}
		sort.Slice(attribute_definition_xhtmlOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_XHTML_stagedOrder[attribute_definition_xhtmlOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_XHTML_stagedOrder[attribute_definition_xhtmlOrdered[j]]
		})
		for _, attribute_definition_xhtml := range attribute_definition_xhtmlOrdered {
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
		attribute_definition_xhtml_renderingOrdered := []*ATTRIBUTE_DEFINITION_XHTML_Rendering{}
		for attribute_definition_xhtml_rendering := range stageSet.Stage.ATTRIBUTE_DEFINITION_XHTML_Renderings {
			attribute_definition_xhtml_renderingOrdered = append(attribute_definition_xhtml_renderingOrdered, attribute_definition_xhtml_rendering)
		}
		sort.Slice(attribute_definition_xhtml_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_DEFINITION_XHTML_Rendering_stagedOrder[attribute_definition_xhtml_renderingOrdered[i]] < stageSet.Stage.ATTRIBUTE_DEFINITION_XHTML_Rendering_stagedOrder[attribute_definition_xhtml_renderingOrdered[j]]
		})
		for _, attribute_definition_xhtml_rendering := range attribute_definition_xhtml_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_xhtml_renderingIdent := "__models" + attribute_definition_xhtml_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_XHTML_Rendering{Name: %s}).Stage(stageSet.Stage)", attribute_definition_xhtml_renderingIdent, __gong__toRawStringLiteral(attribute_definition_xhtml_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_xhtml_renderingIdent, __gong__toRawStringLiteral(attribute_definition_xhtml_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_xhtml_renderingIdent, attribute_definition_xhtml_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_xhtml_renderingIdent, attribute_definition_xhtml_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_xhtml_renderingIdent, attribute_definition_xhtml_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_xhtml_renderingIdent, attribute_definition_xhtml_rendering.Rank))
		}
	}
	if stageSet.Stage != nil {
		attribute_value_booleanOrdered := []*ATTRIBUTE_VALUE_BOOLEAN{}
		for attribute_value_boolean := range stageSet.Stage.ATTRIBUTE_VALUE_BOOLEANs {
			attribute_value_booleanOrdered = append(attribute_value_booleanOrdered, attribute_value_boolean)
		}
		sort.Slice(attribute_value_booleanOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_VALUE_BOOLEAN_stagedOrder[attribute_value_booleanOrdered[i]] < stageSet.Stage.ATTRIBUTE_VALUE_BOOLEAN_stagedOrder[attribute_value_booleanOrdered[j]]
		})
		for _, attribute_value_boolean := range attribute_value_booleanOrdered {
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
		attribute_value_dateOrdered := []*ATTRIBUTE_VALUE_DATE{}
		for attribute_value_date := range stageSet.Stage.ATTRIBUTE_VALUE_DATEs {
			attribute_value_dateOrdered = append(attribute_value_dateOrdered, attribute_value_date)
		}
		sort.Slice(attribute_value_dateOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_VALUE_DATE_stagedOrder[attribute_value_dateOrdered[i]] < stageSet.Stage.ATTRIBUTE_VALUE_DATE_stagedOrder[attribute_value_dateOrdered[j]]
		})
		for _, attribute_value_date := range attribute_value_dateOrdered {
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
		attribute_value_enumerationOrdered := []*ATTRIBUTE_VALUE_ENUMERATION{}
		for attribute_value_enumeration := range stageSet.Stage.ATTRIBUTE_VALUE_ENUMERATIONs {
			attribute_value_enumerationOrdered = append(attribute_value_enumerationOrdered, attribute_value_enumeration)
		}
		sort.Slice(attribute_value_enumerationOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_VALUE_ENUMERATION_stagedOrder[attribute_value_enumerationOrdered[i]] < stageSet.Stage.ATTRIBUTE_VALUE_ENUMERATION_stagedOrder[attribute_value_enumerationOrdered[j]]
		})
		for _, attribute_value_enumeration := range attribute_value_enumerationOrdered {
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
		attribute_value_integerOrdered := []*ATTRIBUTE_VALUE_INTEGER{}
		for attribute_value_integer := range stageSet.Stage.ATTRIBUTE_VALUE_INTEGERs {
			attribute_value_integerOrdered = append(attribute_value_integerOrdered, attribute_value_integer)
		}
		sort.Slice(attribute_value_integerOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_VALUE_INTEGER_stagedOrder[attribute_value_integerOrdered[i]] < stageSet.Stage.ATTRIBUTE_VALUE_INTEGER_stagedOrder[attribute_value_integerOrdered[j]]
		})
		for _, attribute_value_integer := range attribute_value_integerOrdered {
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
		attribute_value_realOrdered := []*ATTRIBUTE_VALUE_REAL{}
		for attribute_value_real := range stageSet.Stage.ATTRIBUTE_VALUE_REALs {
			attribute_value_realOrdered = append(attribute_value_realOrdered, attribute_value_real)
		}
		sort.Slice(attribute_value_realOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_VALUE_REAL_stagedOrder[attribute_value_realOrdered[i]] < stageSet.Stage.ATTRIBUTE_VALUE_REAL_stagedOrder[attribute_value_realOrdered[j]]
		})
		for _, attribute_value_real := range attribute_value_realOrdered {
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
		attribute_value_stringOrdered := []*ATTRIBUTE_VALUE_STRING{}
		for attribute_value_string := range stageSet.Stage.ATTRIBUTE_VALUE_STRINGs {
			attribute_value_stringOrdered = append(attribute_value_stringOrdered, attribute_value_string)
		}
		sort.Slice(attribute_value_stringOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_VALUE_STRING_stagedOrder[attribute_value_stringOrdered[i]] < stageSet.Stage.ATTRIBUTE_VALUE_STRING_stagedOrder[attribute_value_stringOrdered[j]]
		})
		for _, attribute_value_string := range attribute_value_stringOrdered {
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
		attribute_value_xhtmlOrdered := []*ATTRIBUTE_VALUE_XHTML{}
		for attribute_value_xhtml := range stageSet.Stage.ATTRIBUTE_VALUE_XHTMLs {
			attribute_value_xhtmlOrdered = append(attribute_value_xhtmlOrdered, attribute_value_xhtml)
		}
		sort.Slice(attribute_value_xhtmlOrdered, func(i, j int) bool {
			return stageSet.Stage.ATTRIBUTE_VALUE_XHTML_stagedOrder[attribute_value_xhtmlOrdered[i]] < stageSet.Stage.ATTRIBUTE_VALUE_XHTML_stagedOrder[attribute_value_xhtmlOrdered[j]]
		})
		for _, attribute_value_xhtml := range attribute_value_xhtmlOrdered {
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
		}
	}
	if stageSet.Stage != nil {
		a_alternative_idOrdered := []*A_ALTERNATIVE_ID{}
		for a_alternative_id := range stageSet.Stage.A_ALTERNATIVE_IDs {
			a_alternative_idOrdered = append(a_alternative_idOrdered, a_alternative_id)
		}
		sort.Slice(a_alternative_idOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ALTERNATIVE_ID_stagedOrder[a_alternative_idOrdered[i]] < stageSet.Stage.A_ALTERNATIVE_ID_stagedOrder[a_alternative_idOrdered[j]]
		})
		for _, a_alternative_id := range a_alternative_idOrdered {
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
		a_attribute_definition_boolean_refOrdered := []*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF{}
		for a_attribute_definition_boolean_ref := range stageSet.Stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
			a_attribute_definition_boolean_refOrdered = append(a_attribute_definition_boolean_refOrdered, a_attribute_definition_boolean_ref)
		}
		sort.Slice(a_attribute_definition_boolean_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_stagedOrder[a_attribute_definition_boolean_refOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_stagedOrder[a_attribute_definition_boolean_refOrdered[j]]
		})
		for _, a_attribute_definition_boolean_ref := range a_attribute_definition_boolean_refOrdered {
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
		a_attribute_definition_date_refOrdered := []*A_ATTRIBUTE_DEFINITION_DATE_REF{}
		for a_attribute_definition_date_ref := range stageSet.Stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
			a_attribute_definition_date_refOrdered = append(a_attribute_definition_date_refOrdered, a_attribute_definition_date_ref)
		}
		sort.Slice(a_attribute_definition_date_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_DEFINITION_DATE_REF_stagedOrder[a_attribute_definition_date_refOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_DEFINITION_DATE_REF_stagedOrder[a_attribute_definition_date_refOrdered[j]]
		})
		for _, a_attribute_definition_date_ref := range a_attribute_definition_date_refOrdered {
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
		a_attribute_definition_enumeration_refOrdered := []*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF{}
		for a_attribute_definition_enumeration_ref := range stageSet.Stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
			a_attribute_definition_enumeration_refOrdered = append(a_attribute_definition_enumeration_refOrdered, a_attribute_definition_enumeration_ref)
		}
		sort.Slice(a_attribute_definition_enumeration_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_stagedOrder[a_attribute_definition_enumeration_refOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_stagedOrder[a_attribute_definition_enumeration_refOrdered[j]]
		})
		for _, a_attribute_definition_enumeration_ref := range a_attribute_definition_enumeration_refOrdered {
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
		a_attribute_definition_integer_refOrdered := []*A_ATTRIBUTE_DEFINITION_INTEGER_REF{}
		for a_attribute_definition_integer_ref := range stageSet.Stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
			a_attribute_definition_integer_refOrdered = append(a_attribute_definition_integer_refOrdered, a_attribute_definition_integer_ref)
		}
		sort.Slice(a_attribute_definition_integer_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_DEFINITION_INTEGER_REF_stagedOrder[a_attribute_definition_integer_refOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_DEFINITION_INTEGER_REF_stagedOrder[a_attribute_definition_integer_refOrdered[j]]
		})
		for _, a_attribute_definition_integer_ref := range a_attribute_definition_integer_refOrdered {
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
		a_attribute_definition_real_refOrdered := []*A_ATTRIBUTE_DEFINITION_REAL_REF{}
		for a_attribute_definition_real_ref := range stageSet.Stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
			a_attribute_definition_real_refOrdered = append(a_attribute_definition_real_refOrdered, a_attribute_definition_real_ref)
		}
		sort.Slice(a_attribute_definition_real_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_DEFINITION_REAL_REF_stagedOrder[a_attribute_definition_real_refOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_DEFINITION_REAL_REF_stagedOrder[a_attribute_definition_real_refOrdered[j]]
		})
		for _, a_attribute_definition_real_ref := range a_attribute_definition_real_refOrdered {
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
		a_attribute_definition_string_refOrdered := []*A_ATTRIBUTE_DEFINITION_STRING_REF{}
		for a_attribute_definition_string_ref := range stageSet.Stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
			a_attribute_definition_string_refOrdered = append(a_attribute_definition_string_refOrdered, a_attribute_definition_string_ref)
		}
		sort.Slice(a_attribute_definition_string_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_DEFINITION_STRING_REF_stagedOrder[a_attribute_definition_string_refOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_DEFINITION_STRING_REF_stagedOrder[a_attribute_definition_string_refOrdered[j]]
		})
		for _, a_attribute_definition_string_ref := range a_attribute_definition_string_refOrdered {
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
		a_attribute_definition_xhtml_refOrdered := []*A_ATTRIBUTE_DEFINITION_XHTML_REF{}
		for a_attribute_definition_xhtml_ref := range stageSet.Stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
			a_attribute_definition_xhtml_refOrdered = append(a_attribute_definition_xhtml_refOrdered, a_attribute_definition_xhtml_ref)
		}
		sort.Slice(a_attribute_definition_xhtml_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_DEFINITION_XHTML_REF_stagedOrder[a_attribute_definition_xhtml_refOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_DEFINITION_XHTML_REF_stagedOrder[a_attribute_definition_xhtml_refOrdered[j]]
		})
		for _, a_attribute_definition_xhtml_ref := range a_attribute_definition_xhtml_refOrdered {
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
		a_attribute_value_booleanOrdered := []*A_ATTRIBUTE_VALUE_BOOLEAN{}
		for a_attribute_value_boolean := range stageSet.Stage.A_ATTRIBUTE_VALUE_BOOLEANs {
			a_attribute_value_booleanOrdered = append(a_attribute_value_booleanOrdered, a_attribute_value_boolean)
		}
		sort.Slice(a_attribute_value_booleanOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_VALUE_BOOLEAN_stagedOrder[a_attribute_value_booleanOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_VALUE_BOOLEAN_stagedOrder[a_attribute_value_booleanOrdered[j]]
		})
		for _, a_attribute_value_boolean := range a_attribute_value_booleanOrdered {
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
		a_attribute_value_dateOrdered := []*A_ATTRIBUTE_VALUE_DATE{}
		for a_attribute_value_date := range stageSet.Stage.A_ATTRIBUTE_VALUE_DATEs {
			a_attribute_value_dateOrdered = append(a_attribute_value_dateOrdered, a_attribute_value_date)
		}
		sort.Slice(a_attribute_value_dateOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_VALUE_DATE_stagedOrder[a_attribute_value_dateOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_VALUE_DATE_stagedOrder[a_attribute_value_dateOrdered[j]]
		})
		for _, a_attribute_value_date := range a_attribute_value_dateOrdered {
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
		a_attribute_value_enumerationOrdered := []*A_ATTRIBUTE_VALUE_ENUMERATION{}
		for a_attribute_value_enumeration := range stageSet.Stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
			a_attribute_value_enumerationOrdered = append(a_attribute_value_enumerationOrdered, a_attribute_value_enumeration)
		}
		sort.Slice(a_attribute_value_enumerationOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_VALUE_ENUMERATION_stagedOrder[a_attribute_value_enumerationOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_VALUE_ENUMERATION_stagedOrder[a_attribute_value_enumerationOrdered[j]]
		})
		for _, a_attribute_value_enumeration := range a_attribute_value_enumerationOrdered {
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
		a_attribute_value_integerOrdered := []*A_ATTRIBUTE_VALUE_INTEGER{}
		for a_attribute_value_integer := range stageSet.Stage.A_ATTRIBUTE_VALUE_INTEGERs {
			a_attribute_value_integerOrdered = append(a_attribute_value_integerOrdered, a_attribute_value_integer)
		}
		sort.Slice(a_attribute_value_integerOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_VALUE_INTEGER_stagedOrder[a_attribute_value_integerOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_VALUE_INTEGER_stagedOrder[a_attribute_value_integerOrdered[j]]
		})
		for _, a_attribute_value_integer := range a_attribute_value_integerOrdered {
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
		a_attribute_value_realOrdered := []*A_ATTRIBUTE_VALUE_REAL{}
		for a_attribute_value_real := range stageSet.Stage.A_ATTRIBUTE_VALUE_REALs {
			a_attribute_value_realOrdered = append(a_attribute_value_realOrdered, a_attribute_value_real)
		}
		sort.Slice(a_attribute_value_realOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_VALUE_REAL_stagedOrder[a_attribute_value_realOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_VALUE_REAL_stagedOrder[a_attribute_value_realOrdered[j]]
		})
		for _, a_attribute_value_real := range a_attribute_value_realOrdered {
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
		a_attribute_value_stringOrdered := []*A_ATTRIBUTE_VALUE_STRING{}
		for a_attribute_value_string := range stageSet.Stage.A_ATTRIBUTE_VALUE_STRINGs {
			a_attribute_value_stringOrdered = append(a_attribute_value_stringOrdered, a_attribute_value_string)
		}
		sort.Slice(a_attribute_value_stringOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_VALUE_STRING_stagedOrder[a_attribute_value_stringOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_VALUE_STRING_stagedOrder[a_attribute_value_stringOrdered[j]]
		})
		for _, a_attribute_value_string := range a_attribute_value_stringOrdered {
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
		a_attribute_value_xhtmlOrdered := []*A_ATTRIBUTE_VALUE_XHTML{}
		for a_attribute_value_xhtml := range stageSet.Stage.A_ATTRIBUTE_VALUE_XHTMLs {
			a_attribute_value_xhtmlOrdered = append(a_attribute_value_xhtmlOrdered, a_attribute_value_xhtml)
		}
		sort.Slice(a_attribute_value_xhtmlOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_VALUE_XHTML_stagedOrder[a_attribute_value_xhtmlOrdered[i]] < stageSet.Stage.A_ATTRIBUTE_VALUE_XHTML_stagedOrder[a_attribute_value_xhtmlOrdered[j]]
		})
		for _, a_attribute_value_xhtml := range a_attribute_value_xhtmlOrdered {
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
		a_attribute_value_xhtml_1Ordered := []*A_ATTRIBUTE_VALUE_XHTML_1{}
		for a_attribute_value_xhtml_1 := range stageSet.Stage.A_ATTRIBUTE_VALUE_XHTML_1s {
			a_attribute_value_xhtml_1Ordered = append(a_attribute_value_xhtml_1Ordered, a_attribute_value_xhtml_1)
		}
		sort.Slice(a_attribute_value_xhtml_1Ordered, func(i, j int) bool {
			return stageSet.Stage.A_ATTRIBUTE_VALUE_XHTML_1_stagedOrder[a_attribute_value_xhtml_1Ordered[i]] < stageSet.Stage.A_ATTRIBUTE_VALUE_XHTML_1_stagedOrder[a_attribute_value_xhtml_1Ordered[j]]
		})
		for _, a_attribute_value_xhtml_1 := range a_attribute_value_xhtml_1Ordered {
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
		a_childrenOrdered := []*A_CHILDREN{}
		for a_children := range stageSet.Stage.A_CHILDRENs {
			a_childrenOrdered = append(a_childrenOrdered, a_children)
		}
		sort.Slice(a_childrenOrdered, func(i, j int) bool {
			return stageSet.Stage.A_CHILDREN_stagedOrder[a_childrenOrdered[i]] < stageSet.Stage.A_CHILDREN_stagedOrder[a_childrenOrdered[j]]
		})
		for _, a_children := range a_childrenOrdered {
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
		a_core_contentOrdered := []*A_CORE_CONTENT{}
		for a_core_content := range stageSet.Stage.A_CORE_CONTENTs {
			a_core_contentOrdered = append(a_core_contentOrdered, a_core_content)
		}
		sort.Slice(a_core_contentOrdered, func(i, j int) bool {
			return stageSet.Stage.A_CORE_CONTENT_stagedOrder[a_core_contentOrdered[i]] < stageSet.Stage.A_CORE_CONTENT_stagedOrder[a_core_contentOrdered[j]]
		})
		for _, a_core_content := range a_core_contentOrdered {
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
		a_datatypesOrdered := []*A_DATATYPES{}
		for a_datatypes := range stageSet.Stage.A_DATATYPESs {
			a_datatypesOrdered = append(a_datatypesOrdered, a_datatypes)
		}
		sort.Slice(a_datatypesOrdered, func(i, j int) bool {
			return stageSet.Stage.A_DATATYPES_stagedOrder[a_datatypesOrdered[i]] < stageSet.Stage.A_DATATYPES_stagedOrder[a_datatypesOrdered[j]]
		})
		for _, a_datatypes := range a_datatypesOrdered {
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
		a_datatype_definition_boolean_refOrdered := []*A_DATATYPE_DEFINITION_BOOLEAN_REF{}
		for a_datatype_definition_boolean_ref := range stageSet.Stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
			a_datatype_definition_boolean_refOrdered = append(a_datatype_definition_boolean_refOrdered, a_datatype_definition_boolean_ref)
		}
		sort.Slice(a_datatype_definition_boolean_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_DATATYPE_DEFINITION_BOOLEAN_REF_stagedOrder[a_datatype_definition_boolean_refOrdered[i]] < stageSet.Stage.A_DATATYPE_DEFINITION_BOOLEAN_REF_stagedOrder[a_datatype_definition_boolean_refOrdered[j]]
		})
		for _, a_datatype_definition_boolean_ref := range a_datatype_definition_boolean_refOrdered {
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
		a_datatype_definition_date_refOrdered := []*A_DATATYPE_DEFINITION_DATE_REF{}
		for a_datatype_definition_date_ref := range stageSet.Stage.A_DATATYPE_DEFINITION_DATE_REFs {
			a_datatype_definition_date_refOrdered = append(a_datatype_definition_date_refOrdered, a_datatype_definition_date_ref)
		}
		sort.Slice(a_datatype_definition_date_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_DATATYPE_DEFINITION_DATE_REF_stagedOrder[a_datatype_definition_date_refOrdered[i]] < stageSet.Stage.A_DATATYPE_DEFINITION_DATE_REF_stagedOrder[a_datatype_definition_date_refOrdered[j]]
		})
		for _, a_datatype_definition_date_ref := range a_datatype_definition_date_refOrdered {
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
		a_datatype_definition_enumeration_refOrdered := []*A_DATATYPE_DEFINITION_ENUMERATION_REF{}
		for a_datatype_definition_enumeration_ref := range stageSet.Stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
			a_datatype_definition_enumeration_refOrdered = append(a_datatype_definition_enumeration_refOrdered, a_datatype_definition_enumeration_ref)
		}
		sort.Slice(a_datatype_definition_enumeration_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_DATATYPE_DEFINITION_ENUMERATION_REF_stagedOrder[a_datatype_definition_enumeration_refOrdered[i]] < stageSet.Stage.A_DATATYPE_DEFINITION_ENUMERATION_REF_stagedOrder[a_datatype_definition_enumeration_refOrdered[j]]
		})
		for _, a_datatype_definition_enumeration_ref := range a_datatype_definition_enumeration_refOrdered {
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
		a_datatype_definition_integer_refOrdered := []*A_DATATYPE_DEFINITION_INTEGER_REF{}
		for a_datatype_definition_integer_ref := range stageSet.Stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
			a_datatype_definition_integer_refOrdered = append(a_datatype_definition_integer_refOrdered, a_datatype_definition_integer_ref)
		}
		sort.Slice(a_datatype_definition_integer_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_DATATYPE_DEFINITION_INTEGER_REF_stagedOrder[a_datatype_definition_integer_refOrdered[i]] < stageSet.Stage.A_DATATYPE_DEFINITION_INTEGER_REF_stagedOrder[a_datatype_definition_integer_refOrdered[j]]
		})
		for _, a_datatype_definition_integer_ref := range a_datatype_definition_integer_refOrdered {
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
		a_datatype_definition_real_refOrdered := []*A_DATATYPE_DEFINITION_REAL_REF{}
		for a_datatype_definition_real_ref := range stageSet.Stage.A_DATATYPE_DEFINITION_REAL_REFs {
			a_datatype_definition_real_refOrdered = append(a_datatype_definition_real_refOrdered, a_datatype_definition_real_ref)
		}
		sort.Slice(a_datatype_definition_real_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_DATATYPE_DEFINITION_REAL_REF_stagedOrder[a_datatype_definition_real_refOrdered[i]] < stageSet.Stage.A_DATATYPE_DEFINITION_REAL_REF_stagedOrder[a_datatype_definition_real_refOrdered[j]]
		})
		for _, a_datatype_definition_real_ref := range a_datatype_definition_real_refOrdered {
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
		a_datatype_definition_string_refOrdered := []*A_DATATYPE_DEFINITION_STRING_REF{}
		for a_datatype_definition_string_ref := range stageSet.Stage.A_DATATYPE_DEFINITION_STRING_REFs {
			a_datatype_definition_string_refOrdered = append(a_datatype_definition_string_refOrdered, a_datatype_definition_string_ref)
		}
		sort.Slice(a_datatype_definition_string_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_DATATYPE_DEFINITION_STRING_REF_stagedOrder[a_datatype_definition_string_refOrdered[i]] < stageSet.Stage.A_DATATYPE_DEFINITION_STRING_REF_stagedOrder[a_datatype_definition_string_refOrdered[j]]
		})
		for _, a_datatype_definition_string_ref := range a_datatype_definition_string_refOrdered {
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
		a_datatype_definition_xhtml_refOrdered := []*A_DATATYPE_DEFINITION_XHTML_REF{}
		for a_datatype_definition_xhtml_ref := range stageSet.Stage.A_DATATYPE_DEFINITION_XHTML_REFs {
			a_datatype_definition_xhtml_refOrdered = append(a_datatype_definition_xhtml_refOrdered, a_datatype_definition_xhtml_ref)
		}
		sort.Slice(a_datatype_definition_xhtml_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_DATATYPE_DEFINITION_XHTML_REF_stagedOrder[a_datatype_definition_xhtml_refOrdered[i]] < stageSet.Stage.A_DATATYPE_DEFINITION_XHTML_REF_stagedOrder[a_datatype_definition_xhtml_refOrdered[j]]
		})
		for _, a_datatype_definition_xhtml_ref := range a_datatype_definition_xhtml_refOrdered {
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
		a_editable_attsOrdered := []*A_EDITABLE_ATTS{}
		for a_editable_atts := range stageSet.Stage.A_EDITABLE_ATTSs {
			a_editable_attsOrdered = append(a_editable_attsOrdered, a_editable_atts)
		}
		sort.Slice(a_editable_attsOrdered, func(i, j int) bool {
			return stageSet.Stage.A_EDITABLE_ATTS_stagedOrder[a_editable_attsOrdered[i]] < stageSet.Stage.A_EDITABLE_ATTS_stagedOrder[a_editable_attsOrdered[j]]
		})
		for _, a_editable_atts := range a_editable_attsOrdered {
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
		a_enum_value_refOrdered := []*A_ENUM_VALUE_REF{}
		for a_enum_value_ref := range stageSet.Stage.A_ENUM_VALUE_REFs {
			a_enum_value_refOrdered = append(a_enum_value_refOrdered, a_enum_value_ref)
		}
		sort.Slice(a_enum_value_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_ENUM_VALUE_REF_stagedOrder[a_enum_value_refOrdered[i]] < stageSet.Stage.A_ENUM_VALUE_REF_stagedOrder[a_enum_value_refOrdered[j]]
		})
		for _, a_enum_value_ref := range a_enum_value_refOrdered {
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
		a_objectOrdered := []*A_OBJECT{}
		for a_object := range stageSet.Stage.A_OBJECTs {
			a_objectOrdered = append(a_objectOrdered, a_object)
		}
		sort.Slice(a_objectOrdered, func(i, j int) bool {
			return stageSet.Stage.A_OBJECT_stagedOrder[a_objectOrdered[i]] < stageSet.Stage.A_OBJECT_stagedOrder[a_objectOrdered[j]]
		})
		for _, a_object := range a_objectOrdered {
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
		a_propertiesOrdered := []*A_PROPERTIES{}
		for a_properties := range stageSet.Stage.A_PROPERTIESs {
			a_propertiesOrdered = append(a_propertiesOrdered, a_properties)
		}
		sort.Slice(a_propertiesOrdered, func(i, j int) bool {
			return stageSet.Stage.A_PROPERTIES_stagedOrder[a_propertiesOrdered[i]] < stageSet.Stage.A_PROPERTIES_stagedOrder[a_propertiesOrdered[j]]
		})
		for _, a_properties := range a_propertiesOrdered {
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
		a_relation_group_type_refOrdered := []*A_RELATION_GROUP_TYPE_REF{}
		for a_relation_group_type_ref := range stageSet.Stage.A_RELATION_GROUP_TYPE_REFs {
			a_relation_group_type_refOrdered = append(a_relation_group_type_refOrdered, a_relation_group_type_ref)
		}
		sort.Slice(a_relation_group_type_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_RELATION_GROUP_TYPE_REF_stagedOrder[a_relation_group_type_refOrdered[i]] < stageSet.Stage.A_RELATION_GROUP_TYPE_REF_stagedOrder[a_relation_group_type_refOrdered[j]]
		})
		for _, a_relation_group_type_ref := range a_relation_group_type_refOrdered {
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
		a_source_1Ordered := []*A_SOURCE_1{}
		for a_source_1 := range stageSet.Stage.A_SOURCE_1s {
			a_source_1Ordered = append(a_source_1Ordered, a_source_1)
		}
		sort.Slice(a_source_1Ordered, func(i, j int) bool {
			return stageSet.Stage.A_SOURCE_1_stagedOrder[a_source_1Ordered[i]] < stageSet.Stage.A_SOURCE_1_stagedOrder[a_source_1Ordered[j]]
		})
		for _, a_source_1 := range a_source_1Ordered {
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
		a_source_specification_1Ordered := []*A_SOURCE_SPECIFICATION_1{}
		for a_source_specification_1 := range stageSet.Stage.A_SOURCE_SPECIFICATION_1s {
			a_source_specification_1Ordered = append(a_source_specification_1Ordered, a_source_specification_1)
		}
		sort.Slice(a_source_specification_1Ordered, func(i, j int) bool {
			return stageSet.Stage.A_SOURCE_SPECIFICATION_1_stagedOrder[a_source_specification_1Ordered[i]] < stageSet.Stage.A_SOURCE_SPECIFICATION_1_stagedOrder[a_source_specification_1Ordered[j]]
		})
		for _, a_source_specification_1 := range a_source_specification_1Ordered {
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
		a_specificationsOrdered := []*A_SPECIFICATIONS{}
		for a_specifications := range stageSet.Stage.A_SPECIFICATIONSs {
			a_specificationsOrdered = append(a_specificationsOrdered, a_specifications)
		}
		sort.Slice(a_specificationsOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPECIFICATIONS_stagedOrder[a_specificationsOrdered[i]] < stageSet.Stage.A_SPECIFICATIONS_stagedOrder[a_specificationsOrdered[j]]
		})
		for _, a_specifications := range a_specificationsOrdered {
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
		a_specification_type_refOrdered := []*A_SPECIFICATION_TYPE_REF{}
		for a_specification_type_ref := range stageSet.Stage.A_SPECIFICATION_TYPE_REFs {
			a_specification_type_refOrdered = append(a_specification_type_refOrdered, a_specification_type_ref)
		}
		sort.Slice(a_specification_type_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPECIFICATION_TYPE_REF_stagedOrder[a_specification_type_refOrdered[i]] < stageSet.Stage.A_SPECIFICATION_TYPE_REF_stagedOrder[a_specification_type_refOrdered[j]]
		})
		for _, a_specification_type_ref := range a_specification_type_refOrdered {
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
		a_specified_valuesOrdered := []*A_SPECIFIED_VALUES{}
		for a_specified_values := range stageSet.Stage.A_SPECIFIED_VALUESs {
			a_specified_valuesOrdered = append(a_specified_valuesOrdered, a_specified_values)
		}
		sort.Slice(a_specified_valuesOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPECIFIED_VALUES_stagedOrder[a_specified_valuesOrdered[i]] < stageSet.Stage.A_SPECIFIED_VALUES_stagedOrder[a_specified_valuesOrdered[j]]
		})
		for _, a_specified_values := range a_specified_valuesOrdered {
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
		a_spec_attributesOrdered := []*A_SPEC_ATTRIBUTES{}
		for a_spec_attributes := range stageSet.Stage.A_SPEC_ATTRIBUTESs {
			a_spec_attributesOrdered = append(a_spec_attributesOrdered, a_spec_attributes)
		}
		sort.Slice(a_spec_attributesOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPEC_ATTRIBUTES_stagedOrder[a_spec_attributesOrdered[i]] < stageSet.Stage.A_SPEC_ATTRIBUTES_stagedOrder[a_spec_attributesOrdered[j]]
		})
		for _, a_spec_attributes := range a_spec_attributesOrdered {
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
		a_spec_objectsOrdered := []*A_SPEC_OBJECTS{}
		for a_spec_objects := range stageSet.Stage.A_SPEC_OBJECTSs {
			a_spec_objectsOrdered = append(a_spec_objectsOrdered, a_spec_objects)
		}
		sort.Slice(a_spec_objectsOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPEC_OBJECTS_stagedOrder[a_spec_objectsOrdered[i]] < stageSet.Stage.A_SPEC_OBJECTS_stagedOrder[a_spec_objectsOrdered[j]]
		})
		for _, a_spec_objects := range a_spec_objectsOrdered {
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
		a_spec_object_type_refOrdered := []*A_SPEC_OBJECT_TYPE_REF{}
		for a_spec_object_type_ref := range stageSet.Stage.A_SPEC_OBJECT_TYPE_REFs {
			a_spec_object_type_refOrdered = append(a_spec_object_type_refOrdered, a_spec_object_type_ref)
		}
		sort.Slice(a_spec_object_type_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPEC_OBJECT_TYPE_REF_stagedOrder[a_spec_object_type_refOrdered[i]] < stageSet.Stage.A_SPEC_OBJECT_TYPE_REF_stagedOrder[a_spec_object_type_refOrdered[j]]
		})
		for _, a_spec_object_type_ref := range a_spec_object_type_refOrdered {
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
		a_spec_relationsOrdered := []*A_SPEC_RELATIONS{}
		for a_spec_relations := range stageSet.Stage.A_SPEC_RELATIONSs {
			a_spec_relationsOrdered = append(a_spec_relationsOrdered, a_spec_relations)
		}
		sort.Slice(a_spec_relationsOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPEC_RELATIONS_stagedOrder[a_spec_relationsOrdered[i]] < stageSet.Stage.A_SPEC_RELATIONS_stagedOrder[a_spec_relationsOrdered[j]]
		})
		for _, a_spec_relations := range a_spec_relationsOrdered {
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
		a_spec_relation_groupsOrdered := []*A_SPEC_RELATION_GROUPS{}
		for a_spec_relation_groups := range stageSet.Stage.A_SPEC_RELATION_GROUPSs {
			a_spec_relation_groupsOrdered = append(a_spec_relation_groupsOrdered, a_spec_relation_groups)
		}
		sort.Slice(a_spec_relation_groupsOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPEC_RELATION_GROUPS_stagedOrder[a_spec_relation_groupsOrdered[i]] < stageSet.Stage.A_SPEC_RELATION_GROUPS_stagedOrder[a_spec_relation_groupsOrdered[j]]
		})
		for _, a_spec_relation_groups := range a_spec_relation_groupsOrdered {
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
		a_spec_relation_refOrdered := []*A_SPEC_RELATION_REF{}
		for a_spec_relation_ref := range stageSet.Stage.A_SPEC_RELATION_REFs {
			a_spec_relation_refOrdered = append(a_spec_relation_refOrdered, a_spec_relation_ref)
		}
		sort.Slice(a_spec_relation_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPEC_RELATION_REF_stagedOrder[a_spec_relation_refOrdered[i]] < stageSet.Stage.A_SPEC_RELATION_REF_stagedOrder[a_spec_relation_refOrdered[j]]
		})
		for _, a_spec_relation_ref := range a_spec_relation_refOrdered {
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
		a_spec_relation_type_refOrdered := []*A_SPEC_RELATION_TYPE_REF{}
		for a_spec_relation_type_ref := range stageSet.Stage.A_SPEC_RELATION_TYPE_REFs {
			a_spec_relation_type_refOrdered = append(a_spec_relation_type_refOrdered, a_spec_relation_type_ref)
		}
		sort.Slice(a_spec_relation_type_refOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPEC_RELATION_TYPE_REF_stagedOrder[a_spec_relation_type_refOrdered[i]] < stageSet.Stage.A_SPEC_RELATION_TYPE_REF_stagedOrder[a_spec_relation_type_refOrdered[j]]
		})
		for _, a_spec_relation_type_ref := range a_spec_relation_type_refOrdered {
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
		a_spec_typesOrdered := []*A_SPEC_TYPES{}
		for a_spec_types := range stageSet.Stage.A_SPEC_TYPESs {
			a_spec_typesOrdered = append(a_spec_typesOrdered, a_spec_types)
		}
		sort.Slice(a_spec_typesOrdered, func(i, j int) bool {
			return stageSet.Stage.A_SPEC_TYPES_stagedOrder[a_spec_typesOrdered[i]] < stageSet.Stage.A_SPEC_TYPES_stagedOrder[a_spec_typesOrdered[j]]
		})
		for _, a_spec_types := range a_spec_typesOrdered {
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
		a_the_headerOrdered := []*A_THE_HEADER{}
		for a_the_header := range stageSet.Stage.A_THE_HEADERs {
			a_the_headerOrdered = append(a_the_headerOrdered, a_the_header)
		}
		sort.Slice(a_the_headerOrdered, func(i, j int) bool {
			return stageSet.Stage.A_THE_HEADER_stagedOrder[a_the_headerOrdered[i]] < stageSet.Stage.A_THE_HEADER_stagedOrder[a_the_headerOrdered[j]]
		})
		for _, a_the_header := range a_the_headerOrdered {
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
		a_tool_extensionsOrdered := []*A_TOOL_EXTENSIONS{}
		for a_tool_extensions := range stageSet.Stage.A_TOOL_EXTENSIONSs {
			a_tool_extensionsOrdered = append(a_tool_extensionsOrdered, a_tool_extensions)
		}
		sort.Slice(a_tool_extensionsOrdered, func(i, j int) bool {
			return stageSet.Stage.A_TOOL_EXTENSIONS_stagedOrder[a_tool_extensionsOrdered[i]] < stageSet.Stage.A_TOOL_EXTENSIONS_stagedOrder[a_tool_extensionsOrdered[j]]
		})
		for _, a_tool_extensions := range a_tool_extensionsOrdered {
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
		datatype_definition_booleanOrdered := []*DATATYPE_DEFINITION_BOOLEAN{}
		for datatype_definition_boolean := range stageSet.Stage.DATATYPE_DEFINITION_BOOLEANs {
			datatype_definition_booleanOrdered = append(datatype_definition_booleanOrdered, datatype_definition_boolean)
		}
		sort.Slice(datatype_definition_booleanOrdered, func(i, j int) bool {
			return stageSet.Stage.DATATYPE_DEFINITION_BOOLEAN_stagedOrder[datatype_definition_booleanOrdered[i]] < stageSet.Stage.DATATYPE_DEFINITION_BOOLEAN_stagedOrder[datatype_definition_booleanOrdered[j]]
		})
		for _, datatype_definition_boolean := range datatype_definition_booleanOrdered {
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
		datatype_definition_dateOrdered := []*DATATYPE_DEFINITION_DATE{}
		for datatype_definition_date := range stageSet.Stage.DATATYPE_DEFINITION_DATEs {
			datatype_definition_dateOrdered = append(datatype_definition_dateOrdered, datatype_definition_date)
		}
		sort.Slice(datatype_definition_dateOrdered, func(i, j int) bool {
			return stageSet.Stage.DATATYPE_DEFINITION_DATE_stagedOrder[datatype_definition_dateOrdered[i]] < stageSet.Stage.DATATYPE_DEFINITION_DATE_stagedOrder[datatype_definition_dateOrdered[j]]
		})
		for _, datatype_definition_date := range datatype_definition_dateOrdered {
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
		datatype_definition_enumerationOrdered := []*DATATYPE_DEFINITION_ENUMERATION{}
		for datatype_definition_enumeration := range stageSet.Stage.DATATYPE_DEFINITION_ENUMERATIONs {
			datatype_definition_enumerationOrdered = append(datatype_definition_enumerationOrdered, datatype_definition_enumeration)
		}
		sort.Slice(datatype_definition_enumerationOrdered, func(i, j int) bool {
			return stageSet.Stage.DATATYPE_DEFINITION_ENUMERATION_stagedOrder[datatype_definition_enumerationOrdered[i]] < stageSet.Stage.DATATYPE_DEFINITION_ENUMERATION_stagedOrder[datatype_definition_enumerationOrdered[j]]
		})
		for _, datatype_definition_enumeration := range datatype_definition_enumerationOrdered {
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
		datatype_definition_integerOrdered := []*DATATYPE_DEFINITION_INTEGER{}
		for datatype_definition_integer := range stageSet.Stage.DATATYPE_DEFINITION_INTEGERs {
			datatype_definition_integerOrdered = append(datatype_definition_integerOrdered, datatype_definition_integer)
		}
		sort.Slice(datatype_definition_integerOrdered, func(i, j int) bool {
			return stageSet.Stage.DATATYPE_DEFINITION_INTEGER_stagedOrder[datatype_definition_integerOrdered[i]] < stageSet.Stage.DATATYPE_DEFINITION_INTEGER_stagedOrder[datatype_definition_integerOrdered[j]]
		})
		for _, datatype_definition_integer := range datatype_definition_integerOrdered {
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
		datatype_definition_realOrdered := []*DATATYPE_DEFINITION_REAL{}
		for datatype_definition_real := range stageSet.Stage.DATATYPE_DEFINITION_REALs {
			datatype_definition_realOrdered = append(datatype_definition_realOrdered, datatype_definition_real)
		}
		sort.Slice(datatype_definition_realOrdered, func(i, j int) bool {
			return stageSet.Stage.DATATYPE_DEFINITION_REAL_stagedOrder[datatype_definition_realOrdered[i]] < stageSet.Stage.DATATYPE_DEFINITION_REAL_stagedOrder[datatype_definition_realOrdered[j]]
		})
		for _, datatype_definition_real := range datatype_definition_realOrdered {
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
		datatype_definition_stringOrdered := []*DATATYPE_DEFINITION_STRING{}
		for datatype_definition_string := range stageSet.Stage.DATATYPE_DEFINITION_STRINGs {
			datatype_definition_stringOrdered = append(datatype_definition_stringOrdered, datatype_definition_string)
		}
		sort.Slice(datatype_definition_stringOrdered, func(i, j int) bool {
			return stageSet.Stage.DATATYPE_DEFINITION_STRING_stagedOrder[datatype_definition_stringOrdered[i]] < stageSet.Stage.DATATYPE_DEFINITION_STRING_stagedOrder[datatype_definition_stringOrdered[j]]
		})
		for _, datatype_definition_string := range datatype_definition_stringOrdered {
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
		datatype_definition_xhtmlOrdered := []*DATATYPE_DEFINITION_XHTML{}
		for datatype_definition_xhtml := range stageSet.Stage.DATATYPE_DEFINITION_XHTMLs {
			datatype_definition_xhtmlOrdered = append(datatype_definition_xhtmlOrdered, datatype_definition_xhtml)
		}
		sort.Slice(datatype_definition_xhtmlOrdered, func(i, j int) bool {
			return stageSet.Stage.DATATYPE_DEFINITION_XHTML_stagedOrder[datatype_definition_xhtmlOrdered[i]] < stageSet.Stage.DATATYPE_DEFINITION_XHTML_stagedOrder[datatype_definition_xhtmlOrdered[j]]
		})
		for _, datatype_definition_xhtml := range datatype_definition_xhtmlOrdered {
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
		embedded_valueOrdered := []*EMBEDDED_VALUE{}
		for embedded_value := range stageSet.Stage.EMBEDDED_VALUEs {
			embedded_valueOrdered = append(embedded_valueOrdered, embedded_value)
		}
		sort.Slice(embedded_valueOrdered, func(i, j int) bool {
			return stageSet.Stage.EMBEDDED_VALUE_stagedOrder[embedded_valueOrdered[i]] < stageSet.Stage.EMBEDDED_VALUE_stagedOrder[embedded_valueOrdered[j]]
		})
		for _, embedded_value := range embedded_valueOrdered {
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
		enum_valueOrdered := []*ENUM_VALUE{}
		for enum_value := range stageSet.Stage.ENUM_VALUEs {
			enum_valueOrdered = append(enum_valueOrdered, enum_value)
		}
		sort.Slice(enum_valueOrdered, func(i, j int) bool {
			return stageSet.Stage.ENUM_VALUE_stagedOrder[enum_valueOrdered[i]] < stageSet.Stage.ENUM_VALUE_stagedOrder[enum_valueOrdered[j]]
		})
		for _, enum_value := range enum_valueOrdered {
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
		embeddedjpgimageOrdered := []*EmbeddedJpgImage{}
		for embeddedjpgimage := range stageSet.Stage.EmbeddedJpgImages {
			embeddedjpgimageOrdered = append(embeddedjpgimageOrdered, embeddedjpgimage)
		}
		sort.Slice(embeddedjpgimageOrdered, func(i, j int) bool {
			return stageSet.Stage.EmbeddedJpgImage_stagedOrder[embeddedjpgimageOrdered[i]] < stageSet.Stage.EmbeddedJpgImage_stagedOrder[embeddedjpgimageOrdered[j]]
		})
		for _, embeddedjpgimage := range embeddedjpgimageOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			embeddedjpgimageIdent := "__models" + embeddedjpgimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EmbeddedJpgImage{Name: %s}).Stage(stageSet.Stage)", embeddedjpgimageIdent, __gong__toRawStringLiteral(embeddedjpgimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", embeddedjpgimageIdent, __gong__toRawStringLiteral(embeddedjpgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", embeddedjpgimageIdent, __gong__toRawStringLiteral(embeddedjpgimage.Base64Content)))
		}
	}
	if stageSet.Stage != nil {
		embeddedpngimageOrdered := []*EmbeddedPngImage{}
		for embeddedpngimage := range stageSet.Stage.EmbeddedPngImages {
			embeddedpngimageOrdered = append(embeddedpngimageOrdered, embeddedpngimage)
		}
		sort.Slice(embeddedpngimageOrdered, func(i, j int) bool {
			return stageSet.Stage.EmbeddedPngImage_stagedOrder[embeddedpngimageOrdered[i]] < stageSet.Stage.EmbeddedPngImage_stagedOrder[embeddedpngimageOrdered[j]]
		})
		for _, embeddedpngimage := range embeddedpngimageOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			embeddedpngimageIdent := "__models" + embeddedpngimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EmbeddedPngImage{Name: %s}).Stage(stageSet.Stage)", embeddedpngimageIdent, __gong__toRawStringLiteral(embeddedpngimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", embeddedpngimageIdent, __gong__toRawStringLiteral(embeddedpngimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", embeddedpngimageIdent, __gong__toRawStringLiteral(embeddedpngimage.Base64Content)))
		}
	}
	if stageSet.Stage != nil {
		embeddedsvgimageOrdered := []*EmbeddedSvgImage{}
		for embeddedsvgimage := range stageSet.Stage.EmbeddedSvgImages {
			embeddedsvgimageOrdered = append(embeddedsvgimageOrdered, embeddedsvgimage)
		}
		sort.Slice(embeddedsvgimageOrdered, func(i, j int) bool {
			return stageSet.Stage.EmbeddedSvgImage_stagedOrder[embeddedsvgimageOrdered[i]] < stageSet.Stage.EmbeddedSvgImage_stagedOrder[embeddedsvgimageOrdered[j]]
		})
		for _, embeddedsvgimage := range embeddedsvgimageOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			embeddedsvgimageIdent := "__models" + embeddedsvgimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EmbeddedSvgImage{Name: %s}).Stage(stageSet.Stage)", embeddedsvgimageIdent, __gong__toRawStringLiteral(embeddedsvgimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", embeddedsvgimageIdent, __gong__toRawStringLiteral(embeddedsvgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", embeddedsvgimageIdent, __gong__toRawStringLiteral(embeddedsvgimage.Content)))
		}
	}
	if stageSet.Stage != nil {
		killOrdered := []*Kill{}
		for kill := range stageSet.Stage.Kills {
			killOrdered = append(killOrdered, kill)
		}
		sort.Slice(killOrdered, func(i, j int) bool {
			return stageSet.Stage.Kill_stagedOrder[killOrdered[i]] < stageSet.Stage.Kill_stagedOrder[killOrdered[j]]
		})
		for _, kill := range killOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			killIdent := "__models" + kill.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Kill{Name: %s}).Stage(stageSet.Stage)", killIdent, __gong__toRawStringLiteral(kill.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", killIdent, __gong__toRawStringLiteral(kill.Name)))
		}
	}
	if stageSet.Stage != nil {
		map_identifier_boolOrdered := []*Map_identifier_bool{}
		for map_identifier_bool := range stageSet.Stage.Map_identifier_bools {
			map_identifier_boolOrdered = append(map_identifier_boolOrdered, map_identifier_bool)
		}
		sort.Slice(map_identifier_boolOrdered, func(i, j int) bool {
			return stageSet.Stage.Map_identifier_bool_stagedOrder[map_identifier_boolOrdered[i]] < stageSet.Stage.Map_identifier_bool_stagedOrder[map_identifier_boolOrdered[j]]
		})
		for _, map_identifier_bool := range map_identifier_boolOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			map_identifier_boolIdent := "__models" + map_identifier_bool.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Map_identifier_bool{Name: %s}).Stage(stageSet.Stage)", map_identifier_boolIdent, __gong__toRawStringLiteral(map_identifier_bool.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", map_identifier_boolIdent, __gong__toRawStringLiteral(map_identifier_bool.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %t", map_identifier_boolIdent, map_identifier_bool.Value))
		}
	}
	if stageSet.Stage != nil {
		relation_groupOrdered := []*RELATION_GROUP{}
		for relation_group := range stageSet.Stage.RELATION_GROUPs {
			relation_groupOrdered = append(relation_groupOrdered, relation_group)
		}
		sort.Slice(relation_groupOrdered, func(i, j int) bool {
			return stageSet.Stage.RELATION_GROUP_stagedOrder[relation_groupOrdered[i]] < stageSet.Stage.RELATION_GROUP_stagedOrder[relation_groupOrdered[j]]
		})
		for _, relation_group := range relation_groupOrdered {
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
		relation_group_typeOrdered := []*RELATION_GROUP_TYPE{}
		for relation_group_type := range stageSet.Stage.RELATION_GROUP_TYPEs {
			relation_group_typeOrdered = append(relation_group_typeOrdered, relation_group_type)
		}
		sort.Slice(relation_group_typeOrdered, func(i, j int) bool {
			return stageSet.Stage.RELATION_GROUP_TYPE_stagedOrder[relation_group_typeOrdered[i]] < stageSet.Stage.RELATION_GROUP_TYPE_stagedOrder[relation_group_typeOrdered[j]]
		})
		for _, relation_group_type := range relation_group_typeOrdered {
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
		req_ifOrdered := []*REQ_IF{}
		for req_if := range stageSet.Stage.REQ_IFs {
			req_ifOrdered = append(req_ifOrdered, req_if)
		}
		sort.Slice(req_ifOrdered, func(i, j int) bool {
			return stageSet.Stage.REQ_IF_stagedOrder[req_ifOrdered[i]] < stageSet.Stage.REQ_IF_stagedOrder[req_ifOrdered[j]]
		})
		for _, req_if := range req_ifOrdered {
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
		req_if_contentOrdered := []*REQ_IF_CONTENT{}
		for req_if_content := range stageSet.Stage.REQ_IF_CONTENTs {
			req_if_contentOrdered = append(req_if_contentOrdered, req_if_content)
		}
		sort.Slice(req_if_contentOrdered, func(i, j int) bool {
			return stageSet.Stage.REQ_IF_CONTENT_stagedOrder[req_if_contentOrdered[i]] < stageSet.Stage.REQ_IF_CONTENT_stagedOrder[req_if_contentOrdered[j]]
		})
		for _, req_if_content := range req_if_contentOrdered {
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
		req_if_headerOrdered := []*REQ_IF_HEADER{}
		for req_if_header := range stageSet.Stage.REQ_IF_HEADERs {
			req_if_headerOrdered = append(req_if_headerOrdered, req_if_header)
		}
		sort.Slice(req_if_headerOrdered, func(i, j int) bool {
			return stageSet.Stage.REQ_IF_HEADER_stagedOrder[req_if_headerOrdered[i]] < stageSet.Stage.REQ_IF_HEADER_stagedOrder[req_if_headerOrdered[j]]
		})
		for _, req_if_header := range req_if_headerOrdered {
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
		req_if_tool_extensionOrdered := []*REQ_IF_TOOL_EXTENSION{}
		for req_if_tool_extension := range stageSet.Stage.REQ_IF_TOOL_EXTENSIONs {
			req_if_tool_extensionOrdered = append(req_if_tool_extensionOrdered, req_if_tool_extension)
		}
		sort.Slice(req_if_tool_extensionOrdered, func(i, j int) bool {
			return stageSet.Stage.REQ_IF_TOOL_EXTENSION_stagedOrder[req_if_tool_extensionOrdered[i]] < stageSet.Stage.REQ_IF_TOOL_EXTENSION_stagedOrder[req_if_tool_extensionOrdered[j]]
		})
		for _, req_if_tool_extension := range req_if_tool_extensionOrdered {
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
		specificationOrdered := []*SPECIFICATION{}
		for specification := range stageSet.Stage.SPECIFICATIONs {
			specificationOrdered = append(specificationOrdered, specification)
		}
		sort.Slice(specificationOrdered, func(i, j int) bool {
			return stageSet.Stage.SPECIFICATION_stagedOrder[specificationOrdered[i]] < stageSet.Stage.SPECIFICATION_stagedOrder[specificationOrdered[j]]
		})
		for _, specification := range specificationOrdered {
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
		}
	}
	if stageSet.Stage != nil {
		specification_renderingOrdered := []*SPECIFICATION_Rendering{}
		for specification_rendering := range stageSet.Stage.SPECIFICATION_Renderings {
			specification_renderingOrdered = append(specification_renderingOrdered, specification_rendering)
		}
		sort.Slice(specification_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.SPECIFICATION_Rendering_stagedOrder[specification_renderingOrdered[i]] < stageSet.Stage.SPECIFICATION_Rendering_stagedOrder[specification_renderingOrdered[j]]
		})
		for _, specification_rendering := range specification_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specification_renderingIdent := "__models" + specification_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPECIFICATION_Rendering{Name: %s}).Stage(stageSet.Stage)", specification_renderingIdent, __gong__toRawStringLiteral(specification_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specification_renderingIdent, __gong__toRawStringLiteral(specification_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNodeExpanded = %t", specification_renderingIdent, specification_rendering.IsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", specification_renderingIdent, specification_rendering.IsSelected))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithHeadingNumbering = %t", specification_renderingIdent, specification_rendering.IsWithHeadingNumbering))
		}
	}
	if stageSet.Stage != nil {
		specification_typeOrdered := []*SPECIFICATION_TYPE{}
		for specification_type := range stageSet.Stage.SPECIFICATION_TYPEs {
			specification_typeOrdered = append(specification_typeOrdered, specification_type)
		}
		sort.Slice(specification_typeOrdered, func(i, j int) bool {
			return stageSet.Stage.SPECIFICATION_TYPE_stagedOrder[specification_typeOrdered[i]] < stageSet.Stage.SPECIFICATION_TYPE_stagedOrder[specification_typeOrdered[j]]
		})
		for _, specification_type := range specification_typeOrdered {
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
		spec_hierarchyOrdered := []*SPEC_HIERARCHY{}
		for spec_hierarchy := range stageSet.Stage.SPEC_HIERARCHYs {
			spec_hierarchyOrdered = append(spec_hierarchyOrdered, spec_hierarchy)
		}
		sort.Slice(spec_hierarchyOrdered, func(i, j int) bool {
			return stageSet.Stage.SPEC_HIERARCHY_stagedOrder[spec_hierarchyOrdered[i]] < stageSet.Stage.SPEC_HIERARCHY_stagedOrder[spec_hierarchyOrdered[j]]
		})
		for _, spec_hierarchy := range spec_hierarchyOrdered {
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
		}
	}
	if stageSet.Stage != nil {
		spec_objectOrdered := []*SPEC_OBJECT{}
		for spec_object := range stageSet.Stage.SPEC_OBJECTs {
			spec_objectOrdered = append(spec_objectOrdered, spec_object)
		}
		sort.Slice(spec_objectOrdered, func(i, j int) bool {
			return stageSet.Stage.SPEC_OBJECT_stagedOrder[spec_objectOrdered[i]] < stageSet.Stage.SPEC_OBJECT_stagedOrder[spec_objectOrdered[j]]
		})
		for _, spec_object := range spec_objectOrdered {
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
		spec_object_typeOrdered := []*SPEC_OBJECT_TYPE{}
		for spec_object_type := range stageSet.Stage.SPEC_OBJECT_TYPEs {
			spec_object_typeOrdered = append(spec_object_typeOrdered, spec_object_type)
		}
		sort.Slice(spec_object_typeOrdered, func(i, j int) bool {
			return stageSet.Stage.SPEC_OBJECT_TYPE_stagedOrder[spec_object_typeOrdered[i]] < stageSet.Stage.SPEC_OBJECT_TYPE_stagedOrder[spec_object_typeOrdered[j]]
		})
		for _, spec_object_type := range spec_object_typeOrdered {
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
		spec_object_type_renderingOrdered := []*SPEC_OBJECT_TYPE_Rendering{}
		for spec_object_type_rendering := range stageSet.Stage.SPEC_OBJECT_TYPE_Renderings {
			spec_object_type_renderingOrdered = append(spec_object_type_renderingOrdered, spec_object_type_rendering)
		}
		sort.Slice(spec_object_type_renderingOrdered, func(i, j int) bool {
			return stageSet.Stage.SPEC_OBJECT_TYPE_Rendering_stagedOrder[spec_object_type_renderingOrdered[i]] < stageSet.Stage.SPEC_OBJECT_TYPE_Rendering_stagedOrder[spec_object_type_renderingOrdered[j]]
		})
		for _, spec_object_type_rendering := range spec_object_type_renderingOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_object_type_renderingIdent := "__models" + spec_object_type_rendering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_OBJECT_TYPE_Rendering{Name: %s}).Stage(stageSet.Stage)", spec_object_type_renderingIdent, __gong__toRawStringLiteral(spec_object_type_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_object_type_renderingIdent, __gong__toRawStringLiteral(spec_object_type_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNodeExpanded = %t", spec_object_type_renderingIdent, spec_object_type_rendering.IsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowIdentifier = %t", spec_object_type_renderingIdent, spec_object_type_rendering.ShowIdentifier))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowName = %t", spec_object_type_renderingIdent, spec_object_type_rendering.ShowName))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowRelations = %t", spec_object_type_renderingIdent, spec_object_type_rendering.ShowRelations))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHeading = %t", spec_object_type_renderingIdent, spec_object_type_rendering.IsHeading))
		}
	}
	if stageSet.Stage != nil {
		spec_relationOrdered := []*SPEC_RELATION{}
		for spec_relation := range stageSet.Stage.SPEC_RELATIONs {
			spec_relationOrdered = append(spec_relationOrdered, spec_relation)
		}
		sort.Slice(spec_relationOrdered, func(i, j int) bool {
			return stageSet.Stage.SPEC_RELATION_stagedOrder[spec_relationOrdered[i]] < stageSet.Stage.SPEC_RELATION_stagedOrder[spec_relationOrdered[j]]
		})
		for _, spec_relation := range spec_relationOrdered {
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
		spec_relation_typeOrdered := []*SPEC_RELATION_TYPE{}
		for spec_relation_type := range stageSet.Stage.SPEC_RELATION_TYPEs {
			spec_relation_typeOrdered = append(spec_relation_typeOrdered, spec_relation_type)
		}
		sort.Slice(spec_relation_typeOrdered, func(i, j int) bool {
			return stageSet.Stage.SPEC_RELATION_TYPE_stagedOrder[spec_relation_typeOrdered[i]] < stageSet.Stage.SPEC_RELATION_TYPE_stagedOrder[spec_relation_typeOrdered[j]]
		})
		for _, spec_relation_type := range spec_relation_typeOrdered {
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
		staticwebsiteOrdered := []*StaticWebSite{}
		for staticwebsite := range stageSet.Stage.StaticWebSites {
			staticwebsiteOrdered = append(staticwebsiteOrdered, staticwebsite)
		}
		sort.Slice(staticwebsiteOrdered, func(i, j int) bool {
			return stageSet.Stage.StaticWebSite_stagedOrder[staticwebsiteOrdered[i]] < stageSet.Stage.StaticWebSite_stagedOrder[staticwebsiteOrdered[j]]
		})
		for _, staticwebsite := range staticwebsiteOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsiteIdent := "__models" + staticwebsite.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSite{Name: %s}).Stage(stageSet.Stage)", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MarkdownContent = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.MarkdownContent)))
			values.WriteString(fmt.Sprintf("\n\t%s.InputImagesDir = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.InputImagesDir)))
			values.WriteString(fmt.Sprintf("\n\t%s.OutputStaticWebDir = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.OutputStaticWebDir)))
			values.WriteString(fmt.Sprintf("\n\t%s.VersionInfo = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.VersionInfo)))
			for _, elem := range staticwebsite.Chapters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Chapters = append(%s.Chapters, %s)", staticwebsiteIdent, staticwebsiteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		staticwebsitechapterOrdered := []*StaticWebSiteChapter{}
		for staticwebsitechapter := range stageSet.Stage.StaticWebSiteChapters {
			staticwebsitechapterOrdered = append(staticwebsitechapterOrdered, staticwebsitechapter)
		}
		sort.Slice(staticwebsitechapterOrdered, func(i, j int) bool {
			return stageSet.Stage.StaticWebSiteChapter_stagedOrder[staticwebsitechapterOrdered[i]] < stageSet.Stage.StaticWebSiteChapter_stagedOrder[staticwebsitechapterOrdered[j]]
		})
		for _, staticwebsitechapter := range staticwebsitechapterOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsitechapterIdent := "__models" + staticwebsitechapter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSiteChapter{Name: %s}).Stage(stageSet.Stage)", staticwebsitechapterIdent, __gong__toRawStringLiteral(staticwebsitechapter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsitechapterIdent, __gong__toRawStringLiteral(staticwebsitechapter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MarkdownContent = %s", staticwebsitechapterIdent, __gong__toRawStringLiteral(staticwebsitechapter.MarkdownContent)))
			for _, elem := range staticwebsitechapter.Paragraphs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paragraphs = append(%s.Paragraphs, %s)", staticwebsitechapterIdent, staticwebsitechapterIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		staticwebsitegeneratedimageOrdered := []*StaticWebSiteGeneratedImage{}
		for staticwebsitegeneratedimage := range stageSet.Stage.StaticWebSiteGeneratedImages {
			staticwebsitegeneratedimageOrdered = append(staticwebsitegeneratedimageOrdered, staticwebsitegeneratedimage)
		}
		sort.Slice(staticwebsitegeneratedimageOrdered, func(i, j int) bool {
			return stageSet.Stage.StaticWebSiteGeneratedImage_stagedOrder[staticwebsitegeneratedimageOrdered[i]] < stageSet.Stage.StaticWebSiteGeneratedImage_stagedOrder[staticwebsitegeneratedimageOrdered[j]]
		})
		for _, staticwebsitegeneratedimage := range staticwebsitegeneratedimageOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsitegeneratedimageIdent := "__models" + staticwebsitegeneratedimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSiteGeneratedImage{Name: %s}).Stage(stageSet.Stage)", staticwebsitegeneratedimageIdent, __gong__toRawStringLiteral(staticwebsitegeneratedimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsitegeneratedimageIdent, __gong__toRawStringLiteral(staticwebsitegeneratedimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SourceDirectoryPath = %s", staticwebsitegeneratedimageIdent, __gong__toRawStringLiteral(staticwebsitegeneratedimage.SourceDirectoryPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", staticwebsitegeneratedimageIdent, staticwebsitegeneratedimage.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", staticwebsitegeneratedimageIdent, staticwebsitegeneratedimage.Height))
		}
	}
	if stageSet.Stage != nil {
		staticwebsiteimageOrdered := []*StaticWebSiteImage{}
		for staticwebsiteimage := range stageSet.Stage.StaticWebSiteImages {
			staticwebsiteimageOrdered = append(staticwebsiteimageOrdered, staticwebsiteimage)
		}
		sort.Slice(staticwebsiteimageOrdered, func(i, j int) bool {
			return stageSet.Stage.StaticWebSiteImage_stagedOrder[staticwebsiteimageOrdered[i]] < stageSet.Stage.StaticWebSiteImage_stagedOrder[staticwebsiteimageOrdered[j]]
		})
		for _, staticwebsiteimage := range staticwebsiteimageOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsiteimageIdent := "__models" + staticwebsiteimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSiteImage{Name: %s}).Stage(stageSet.Stage)", staticwebsiteimageIdent, __gong__toRawStringLiteral(staticwebsiteimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsiteimageIdent, __gong__toRawStringLiteral(staticwebsiteimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SourceDirectoryPath = %s", staticwebsiteimageIdent, __gong__toRawStringLiteral(staticwebsiteimage.SourceDirectoryPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", staticwebsiteimageIdent, staticwebsiteimage.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", staticwebsiteimageIdent, staticwebsiteimage.Height))
		}
	}
	if stageSet.Stage != nil {
		staticwebsiteparagraphOrdered := []*StaticWebSiteParagraph{}
		for staticwebsiteparagraph := range stageSet.Stage.StaticWebSiteParagraphs {
			staticwebsiteparagraphOrdered = append(staticwebsiteparagraphOrdered, staticwebsiteparagraph)
		}
		sort.Slice(staticwebsiteparagraphOrdered, func(i, j int) bool {
			return stageSet.Stage.StaticWebSiteParagraph_stagedOrder[staticwebsiteparagraphOrdered[i]] < stageSet.Stage.StaticWebSiteParagraph_stagedOrder[staticwebsiteparagraphOrdered[j]]
		})
		for _, staticwebsiteparagraph := range staticwebsiteparagraphOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsiteparagraphIdent := "__models" + staticwebsiteparagraph.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSiteParagraph{Name: %s}).Stage(stageSet.Stage)", staticwebsiteparagraphIdent, __gong__toRawStringLiteral(staticwebsiteparagraph.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsiteparagraphIdent, __gong__toRawStringLiteral(staticwebsiteparagraph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.LegendMarkdownContent = %s", staticwebsiteparagraphIdent, __gong__toRawStringLiteral(staticwebsiteparagraph.LegendMarkdownContent)))
			if staticwebsiteparagraph.Image != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + staticwebsiteparagraph.Image.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Image = %s", staticwebsiteparagraphIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		xhtml_contentOrdered := []*XHTML_CONTENT{}
		for xhtml_content := range stageSet.Stage.XHTML_CONTENTs {
			xhtml_contentOrdered = append(xhtml_contentOrdered, xhtml_content)
		}
		sort.Slice(xhtml_contentOrdered, func(i, j int) bool {
			return stageSet.Stage.XHTML_CONTENT_stagedOrder[xhtml_contentOrdered[i]] < stageSet.Stage.XHTML_CONTENT_stagedOrder[xhtml_contentOrdered[j]]
		})
		for _, xhtml_content := range xhtml_contentOrdered {
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
			values.WriteString(fmt.Sprintf("\n\t%s.PureText = %s", xhtml_contentIdent, __gong__toRawStringLiteral(xhtml_content.PureText)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/reqif/go/models"
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
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "ALTERNATIVE_ID":
					if !preserveOrder {
						inst := (&ALTERNATIVE_ID{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ALTERNATIVE_ID)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_BOOLEAN":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_BOOLEAN{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_BOOLEAN)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_BOOLEAN_Rendering":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_BOOLEAN_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_BOOLEAN_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_DATE":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_DATE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_DATE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_DATE_Rendering":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_DATE_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_DATE_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_ENUMERATION":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_ENUMERATION{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_ENUMERATION)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_ENUMERATION_Rendering":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_ENUMERATION_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_ENUMERATION_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_INTEGER":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_INTEGER{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_INTEGER)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_INTEGER_Rendering":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_INTEGER_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_INTEGER_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_REAL":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_REAL{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_REAL)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_REAL_Rendering":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_REAL_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_REAL_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_Rendering":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_STRING":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_STRING{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_STRING)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_STRING_Rendering":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_STRING_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_STRING_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_XHTML":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_XHTML{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_XHTML)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_DEFINITION_XHTML_Rendering":
					if !preserveOrder {
						inst := (&ATTRIBUTE_DEFINITION_XHTML_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_DEFINITION_XHTML_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_VALUE_BOOLEAN":
					if !preserveOrder {
						inst := (&ATTRIBUTE_VALUE_BOOLEAN{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_VALUE_BOOLEAN)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_VALUE_DATE":
					if !preserveOrder {
						inst := (&ATTRIBUTE_VALUE_DATE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_VALUE_DATE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_VALUE_ENUMERATION":
					if !preserveOrder {
						inst := (&ATTRIBUTE_VALUE_ENUMERATION{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_VALUE_ENUMERATION)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_VALUE_INTEGER":
					if !preserveOrder {
						inst := (&ATTRIBUTE_VALUE_INTEGER{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_VALUE_INTEGER)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_VALUE_REAL":
					if !preserveOrder {
						inst := (&ATTRIBUTE_VALUE_REAL{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_VALUE_REAL)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_VALUE_STRING":
					if !preserveOrder {
						inst := (&ATTRIBUTE_VALUE_STRING{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_VALUE_STRING)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ATTRIBUTE_VALUE_XHTML":
					if !preserveOrder {
						inst := (&ATTRIBUTE_VALUE_XHTML{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ATTRIBUTE_VALUE_XHTML)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ALTERNATIVE_ID":
					if !preserveOrder {
						inst := (&A_ALTERNATIVE_ID{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ALTERNATIVE_ID)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_DEFINITION_BOOLEAN_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_DEFINITION_DATE_REF":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_DEFINITION_DATE_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_DEFINITION_DATE_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_DEFINITION_ENUMERATION_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_DEFINITION_INTEGER_REF":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_DEFINITION_INTEGER_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_DEFINITION_INTEGER_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_DEFINITION_REAL_REF":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_DEFINITION_REAL_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_DEFINITION_REAL_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_DEFINITION_STRING_REF":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_DEFINITION_STRING_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_DEFINITION_STRING_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_DEFINITION_XHTML_REF":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_DEFINITION_XHTML_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_DEFINITION_XHTML_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_VALUE_BOOLEAN":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_VALUE_BOOLEAN{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_VALUE_BOOLEAN)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_VALUE_DATE":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_VALUE_DATE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_VALUE_DATE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_VALUE_ENUMERATION":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_VALUE_ENUMERATION{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_VALUE_ENUMERATION)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_VALUE_INTEGER":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_VALUE_INTEGER{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_VALUE_INTEGER)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_VALUE_REAL":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_VALUE_REAL{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_VALUE_REAL)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_VALUE_STRING":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_VALUE_STRING{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_VALUE_STRING)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_VALUE_XHTML":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_VALUE_XHTML{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_VALUE_XHTML)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ATTRIBUTE_VALUE_XHTML_1":
					if !preserveOrder {
						inst := (&A_ATTRIBUTE_VALUE_XHTML_1{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ATTRIBUTE_VALUE_XHTML_1)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_CHILDREN":
					if !preserveOrder {
						inst := (&A_CHILDREN{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_CHILDREN)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_CORE_CONTENT":
					if !preserveOrder {
						inst := (&A_CORE_CONTENT{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_CORE_CONTENT)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_DATATYPES":
					if !preserveOrder {
						inst := (&A_DATATYPES{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_DATATYPES)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_DATATYPE_DEFINITION_BOOLEAN_REF":
					if !preserveOrder {
						inst := (&A_DATATYPE_DEFINITION_BOOLEAN_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_DATATYPE_DEFINITION_BOOLEAN_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_DATATYPE_DEFINITION_DATE_REF":
					if !preserveOrder {
						inst := (&A_DATATYPE_DEFINITION_DATE_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_DATATYPE_DEFINITION_DATE_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_DATATYPE_DEFINITION_ENUMERATION_REF":
					if !preserveOrder {
						inst := (&A_DATATYPE_DEFINITION_ENUMERATION_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_DATATYPE_DEFINITION_ENUMERATION_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_DATATYPE_DEFINITION_INTEGER_REF":
					if !preserveOrder {
						inst := (&A_DATATYPE_DEFINITION_INTEGER_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_DATATYPE_DEFINITION_INTEGER_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_DATATYPE_DEFINITION_REAL_REF":
					if !preserveOrder {
						inst := (&A_DATATYPE_DEFINITION_REAL_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_DATATYPE_DEFINITION_REAL_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_DATATYPE_DEFINITION_STRING_REF":
					if !preserveOrder {
						inst := (&A_DATATYPE_DEFINITION_STRING_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_DATATYPE_DEFINITION_STRING_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_DATATYPE_DEFINITION_XHTML_REF":
					if !preserveOrder {
						inst := (&A_DATATYPE_DEFINITION_XHTML_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_DATATYPE_DEFINITION_XHTML_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_EDITABLE_ATTS":
					if !preserveOrder {
						inst := (&A_EDITABLE_ATTS{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_EDITABLE_ATTS)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_ENUM_VALUE_REF":
					if !preserveOrder {
						inst := (&A_ENUM_VALUE_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_ENUM_VALUE_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_OBJECT":
					if !preserveOrder {
						inst := (&A_OBJECT{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_OBJECT)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_PROPERTIES":
					if !preserveOrder {
						inst := (&A_PROPERTIES{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_PROPERTIES)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_RELATION_GROUP_TYPE_REF":
					if !preserveOrder {
						inst := (&A_RELATION_GROUP_TYPE_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_RELATION_GROUP_TYPE_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SOURCE_1":
					if !preserveOrder {
						inst := (&A_SOURCE_1{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SOURCE_1)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SOURCE_SPECIFICATION_1":
					if !preserveOrder {
						inst := (&A_SOURCE_SPECIFICATION_1{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SOURCE_SPECIFICATION_1)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPECIFICATIONS":
					if !preserveOrder {
						inst := (&A_SPECIFICATIONS{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPECIFICATIONS)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPECIFICATION_TYPE_REF":
					if !preserveOrder {
						inst := (&A_SPECIFICATION_TYPE_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPECIFICATION_TYPE_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPECIFIED_VALUES":
					if !preserveOrder {
						inst := (&A_SPECIFIED_VALUES{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPECIFIED_VALUES)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPEC_ATTRIBUTES":
					if !preserveOrder {
						inst := (&A_SPEC_ATTRIBUTES{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPEC_ATTRIBUTES)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPEC_OBJECTS":
					if !preserveOrder {
						inst := (&A_SPEC_OBJECTS{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPEC_OBJECTS)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPEC_OBJECT_TYPE_REF":
					if !preserveOrder {
						inst := (&A_SPEC_OBJECT_TYPE_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPEC_OBJECT_TYPE_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPEC_RELATIONS":
					if !preserveOrder {
						inst := (&A_SPEC_RELATIONS{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPEC_RELATIONS)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPEC_RELATION_GROUPS":
					if !preserveOrder {
						inst := (&A_SPEC_RELATION_GROUPS{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPEC_RELATION_GROUPS)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPEC_RELATION_REF":
					if !preserveOrder {
						inst := (&A_SPEC_RELATION_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPEC_RELATION_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPEC_RELATION_TYPE_REF":
					if !preserveOrder {
						inst := (&A_SPEC_RELATION_TYPE_REF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPEC_RELATION_TYPE_REF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_SPEC_TYPES":
					if !preserveOrder {
						inst := (&A_SPEC_TYPES{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_SPEC_TYPES)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_THE_HEADER":
					if !preserveOrder {
						inst := (&A_THE_HEADER{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_THE_HEADER)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_TOOL_EXTENSIONS":
					if !preserveOrder {
						inst := (&A_TOOL_EXTENSIONS{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_TOOL_EXTENSIONS)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DATATYPE_DEFINITION_BOOLEAN":
					if !preserveOrder {
						inst := (&DATATYPE_DEFINITION_BOOLEAN{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DATATYPE_DEFINITION_BOOLEAN)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DATATYPE_DEFINITION_DATE":
					if !preserveOrder {
						inst := (&DATATYPE_DEFINITION_DATE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DATATYPE_DEFINITION_DATE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DATATYPE_DEFINITION_ENUMERATION":
					if !preserveOrder {
						inst := (&DATATYPE_DEFINITION_ENUMERATION{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DATATYPE_DEFINITION_ENUMERATION)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DATATYPE_DEFINITION_INTEGER":
					if !preserveOrder {
						inst := (&DATATYPE_DEFINITION_INTEGER{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DATATYPE_DEFINITION_INTEGER)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DATATYPE_DEFINITION_REAL":
					if !preserveOrder {
						inst := (&DATATYPE_DEFINITION_REAL{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DATATYPE_DEFINITION_REAL)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DATATYPE_DEFINITION_STRING":
					if !preserveOrder {
						inst := (&DATATYPE_DEFINITION_STRING{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DATATYPE_DEFINITION_STRING)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DATATYPE_DEFINITION_XHTML":
					if !preserveOrder {
						inst := (&DATATYPE_DEFINITION_XHTML{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DATATYPE_DEFINITION_XHTML)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "EMBEDDED_VALUE":
					if !preserveOrder {
						inst := (&EMBEDDED_VALUE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(EMBEDDED_VALUE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ENUM_VALUE":
					if !preserveOrder {
						inst := (&ENUM_VALUE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ENUM_VALUE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "EmbeddedJpgImage":
					if !preserveOrder {
						inst := (&EmbeddedJpgImage{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(EmbeddedJpgImage)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "EmbeddedPngImage":
					if !preserveOrder {
						inst := (&EmbeddedPngImage{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(EmbeddedPngImage)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "EmbeddedSvgImage":
					if !preserveOrder {
						inst := (&EmbeddedSvgImage{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(EmbeddedSvgImage)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Kill":
					if !preserveOrder {
						inst := (&Kill{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Kill)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Map_identifier_bool":
					if !preserveOrder {
						inst := (&Map_identifier_bool{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Map_identifier_bool)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "RELATION_GROUP":
					if !preserveOrder {
						inst := (&RELATION_GROUP{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(RELATION_GROUP)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "RELATION_GROUP_TYPE":
					if !preserveOrder {
						inst := (&RELATION_GROUP_TYPE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(RELATION_GROUP_TYPE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "REQ_IF":
					if !preserveOrder {
						inst := (&REQ_IF{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(REQ_IF)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "REQ_IF_CONTENT":
					if !preserveOrder {
						inst := (&REQ_IF_CONTENT{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(REQ_IF_CONTENT)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "REQ_IF_HEADER":
					if !preserveOrder {
						inst := (&REQ_IF_HEADER{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(REQ_IF_HEADER)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "REQ_IF_TOOL_EXTENSION":
					if !preserveOrder {
						inst := (&REQ_IF_TOOL_EXTENSION{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(REQ_IF_TOOL_EXTENSION)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SPECIFICATION":
					if !preserveOrder {
						inst := (&SPECIFICATION{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SPECIFICATION)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SPECIFICATION_Rendering":
					if !preserveOrder {
						inst := (&SPECIFICATION_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SPECIFICATION_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SPECIFICATION_TYPE":
					if !preserveOrder {
						inst := (&SPECIFICATION_TYPE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SPECIFICATION_TYPE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SPEC_HIERARCHY":
					if !preserveOrder {
						inst := (&SPEC_HIERARCHY{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SPEC_HIERARCHY)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SPEC_OBJECT":
					if !preserveOrder {
						inst := (&SPEC_OBJECT{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SPEC_OBJECT)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SPEC_OBJECT_TYPE":
					if !preserveOrder {
						inst := (&SPEC_OBJECT_TYPE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SPEC_OBJECT_TYPE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SPEC_OBJECT_TYPE_Rendering":
					if !preserveOrder {
						inst := (&SPEC_OBJECT_TYPE_Rendering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SPEC_OBJECT_TYPE_Rendering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SPEC_RELATION":
					if !preserveOrder {
						inst := (&SPEC_RELATION{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SPEC_RELATION)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SPEC_RELATION_TYPE":
					if !preserveOrder {
						inst := (&SPEC_RELATION_TYPE{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SPEC_RELATION_TYPE)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StaticWebSite":
					if !preserveOrder {
						inst := (&StaticWebSite{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StaticWebSite)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StaticWebSiteChapter":
					if !preserveOrder {
						inst := (&StaticWebSiteChapter{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StaticWebSiteChapter)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StaticWebSiteGeneratedImage":
					if !preserveOrder {
						inst := (&StaticWebSiteGeneratedImage{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StaticWebSiteGeneratedImage)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StaticWebSiteImage":
					if !preserveOrder {
						inst := (&StaticWebSiteImage{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StaticWebSiteImage)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StaticWebSiteParagraph":
					if !preserveOrder {
						inst := (&StaticWebSiteParagraph{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StaticWebSiteParagraph)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "XHTML_CONTENT":
					if !preserveOrder {
						inst := (&XHTML_CONTENT{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(XHTML_CONTENT)
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "DEFAULT_VALUE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_BOOLEAN); ok {
									inst.DEFAULT_VALUE = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_DATATYPE_DEFINITION_BOOLEAN_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "DEFAULT_VALUE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_DATE); ok {
									inst.DEFAULT_VALUE = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_DATATYPE_DEFINITION_DATE_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "DEFAULT_VALUE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_ENUMERATION); ok {
									inst.DEFAULT_VALUE = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_DATATYPE_DEFINITION_ENUMERATION_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "DEFAULT_VALUE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_INTEGER); ok {
									inst.DEFAULT_VALUE = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_DATATYPE_DEFINITION_INTEGER_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "DEFAULT_VALUE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_REAL); ok {
									inst.DEFAULT_VALUE = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_DATATYPE_DEFINITION_REAL_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "DEFAULT_VALUE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_STRING); ok {
									inst.DEFAULT_VALUE = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_DATATYPE_DEFINITION_STRING_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "DEFAULT_VALUE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_XHTML); ok {
									inst.DEFAULT_VALUE = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_DATATYPE_DEFINITION_XHTML_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF); ok {
									inst.DEFINITION = typedTarget
								}
							}
						}
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractBool(rhs)
					}
				case *ATTRIBUTE_VALUE_DATE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_DEFINITION_DATE_REF); ok {
									inst.DEFINITION = typedTarget
								}
							}
						}
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractString(rhs)
					}
				case *ATTRIBUTE_VALUE_ENUMERATION:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF); ok {
									inst.DEFINITION = typedTarget
								}
							}
						}
					case "VALUES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ENUM_VALUE_REF); ok {
									inst.VALUES = typedTarget
								}
							}
						}
					}
				case *ATTRIBUTE_VALUE_INTEGER:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_DEFINITION_INTEGER_REF); ok {
									inst.DEFINITION = typedTarget
								}
							}
						}
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractInt(rhs)
					}
				case *ATTRIBUTE_VALUE_REAL:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_DEFINITION_REAL_REF); ok {
									inst.DEFINITION = typedTarget
								}
							}
						}
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractFloat(rhs)
					}
				case *ATTRIBUTE_VALUE_STRING:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_DEFINITION_STRING_REF); ok {
									inst.DEFINITION = typedTarget
								}
							}
						}
					case "THE_VALUE":
						inst.THE_VALUE = GongExtractString(rhs)
					}
				case *ATTRIBUTE_VALUE_XHTML:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DEFINITION":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_DEFINITION_XHTML_REF); ok {
									inst.DEFINITION = typedTarget
								}
							}
						}
					case "IS_SIMPLIFIED":
						inst.IS_SIMPLIFIED = GongExtractBool(rhs)
					case "THE_VALUE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*XHTML_CONTENT); ok {
									inst.THE_VALUE = typedTarget
								}
							}
						}
					case "THE_ORIGINAL_VALUE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*XHTML_CONTENT); ok {
									inst.THE_ORIGINAL_VALUE = typedTarget
								}
							}
						}
					}
				case *A_ALTERNATIVE_ID:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ALTERNATIVE_ID":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_BOOLEAN); ok {
										inst.ATTRIBUTE_VALUE_BOOLEAN = append(inst.ATTRIBUTE_VALUE_BOOLEAN, typedTarget)
									}
								}
							}
						}
					}
				case *A_ATTRIBUTE_VALUE_DATE:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_DATE":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_DATE); ok {
										inst.ATTRIBUTE_VALUE_DATE = append(inst.ATTRIBUTE_VALUE_DATE, typedTarget)
									}
								}
							}
						}
					}
				case *A_ATTRIBUTE_VALUE_ENUMERATION:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_ENUMERATION":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_ENUMERATION); ok {
										inst.ATTRIBUTE_VALUE_ENUMERATION = append(inst.ATTRIBUTE_VALUE_ENUMERATION, typedTarget)
									}
								}
							}
						}
					}
				case *A_ATTRIBUTE_VALUE_INTEGER:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_INTEGER":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_INTEGER); ok {
										inst.ATTRIBUTE_VALUE_INTEGER = append(inst.ATTRIBUTE_VALUE_INTEGER, typedTarget)
									}
								}
							}
						}
					}
				case *A_ATTRIBUTE_VALUE_REAL:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_REAL":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_REAL); ok {
										inst.ATTRIBUTE_VALUE_REAL = append(inst.ATTRIBUTE_VALUE_REAL, typedTarget)
									}
								}
							}
						}
					}
				case *A_ATTRIBUTE_VALUE_STRING:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_STRING":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_STRING); ok {
										inst.ATTRIBUTE_VALUE_STRING = append(inst.ATTRIBUTE_VALUE_STRING, typedTarget)
									}
								}
							}
						}
					}
				case *A_ATTRIBUTE_VALUE_XHTML:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_XHTML":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_XHTML); ok {
										inst.ATTRIBUTE_VALUE_XHTML = append(inst.ATTRIBUTE_VALUE_XHTML, typedTarget)
									}
								}
							}
						}
					}
				case *A_ATTRIBUTE_VALUE_XHTML_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_VALUE_BOOLEAN":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_BOOLEAN); ok {
										inst.ATTRIBUTE_VALUE_BOOLEAN = append(inst.ATTRIBUTE_VALUE_BOOLEAN, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_VALUE_DATE":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_DATE); ok {
										inst.ATTRIBUTE_VALUE_DATE = append(inst.ATTRIBUTE_VALUE_DATE, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_VALUE_ENUMERATION":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_ENUMERATION); ok {
										inst.ATTRIBUTE_VALUE_ENUMERATION = append(inst.ATTRIBUTE_VALUE_ENUMERATION, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_VALUE_INTEGER":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_INTEGER); ok {
										inst.ATTRIBUTE_VALUE_INTEGER = append(inst.ATTRIBUTE_VALUE_INTEGER, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_VALUE_REAL":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_REAL); ok {
										inst.ATTRIBUTE_VALUE_REAL = append(inst.ATTRIBUTE_VALUE_REAL, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_VALUE_STRING":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_STRING); ok {
										inst.ATTRIBUTE_VALUE_STRING = append(inst.ATTRIBUTE_VALUE_STRING, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_VALUE_XHTML":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_VALUE_XHTML); ok {
										inst.ATTRIBUTE_VALUE_XHTML = append(inst.ATTRIBUTE_VALUE_XHTML, typedTarget)
									}
								}
							}
						}
					}
				case *A_CHILDREN:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_HIERARCHY":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*SPEC_HIERARCHY); ok {
										inst.SPEC_HIERARCHY = append(inst.SPEC_HIERARCHY, typedTarget)
									}
								}
							}
						}
					}
				case *A_CORE_CONTENT:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "REQ_IF_CONTENT":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*REQ_IF_CONTENT); ok {
									inst.REQ_IF_CONTENT = typedTarget
								}
							}
						}
					}
				case *A_DATATYPES:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPE_DEFINITION_BOOLEAN":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DATATYPE_DEFINITION_BOOLEAN); ok {
										inst.DATATYPE_DEFINITION_BOOLEAN = append(inst.DATATYPE_DEFINITION_BOOLEAN, typedTarget)
									}
								}
							}
						}
					case "DATATYPE_DEFINITION_DATE":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DATATYPE_DEFINITION_DATE); ok {
										inst.DATATYPE_DEFINITION_DATE = append(inst.DATATYPE_DEFINITION_DATE, typedTarget)
									}
								}
							}
						}
					case "DATATYPE_DEFINITION_ENUMERATION":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DATATYPE_DEFINITION_ENUMERATION); ok {
										inst.DATATYPE_DEFINITION_ENUMERATION = append(inst.DATATYPE_DEFINITION_ENUMERATION, typedTarget)
									}
								}
							}
						}
					case "DATATYPE_DEFINITION_INTEGER":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DATATYPE_DEFINITION_INTEGER); ok {
										inst.DATATYPE_DEFINITION_INTEGER = append(inst.DATATYPE_DEFINITION_INTEGER, typedTarget)
									}
								}
							}
						}
					case "DATATYPE_DEFINITION_REAL":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DATATYPE_DEFINITION_REAL); ok {
										inst.DATATYPE_DEFINITION_REAL = append(inst.DATATYPE_DEFINITION_REAL, typedTarget)
									}
								}
							}
						}
					case "DATATYPE_DEFINITION_STRING":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DATATYPE_DEFINITION_STRING); ok {
										inst.DATATYPE_DEFINITION_STRING = append(inst.DATATYPE_DEFINITION_STRING, typedTarget)
									}
								}
							}
						}
					case "DATATYPE_DEFINITION_XHTML":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*DATATYPE_DEFINITION_XHTML); ok {
										inst.DATATYPE_DEFINITION_XHTML = append(inst.DATATYPE_DEFINITION_XHTML, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EMBEDDED_VALUE); ok {
									inst.EMBEDDED_VALUE = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*SPECIFICATION); ok {
										inst.SPECIFICATION = append(inst.SPECIFICATION, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ENUM_VALUE); ok {
										inst.ENUM_VALUE = append(inst.ENUM_VALUE, typedTarget)
									}
								}
							}
						}
					}
				case *A_SPEC_ATTRIBUTES:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ATTRIBUTE_DEFINITION_BOOLEAN":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_DEFINITION_BOOLEAN); ok {
										inst.ATTRIBUTE_DEFINITION_BOOLEAN = append(inst.ATTRIBUTE_DEFINITION_BOOLEAN, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_DEFINITION_DATE":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_DEFINITION_DATE); ok {
										inst.ATTRIBUTE_DEFINITION_DATE = append(inst.ATTRIBUTE_DEFINITION_DATE, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_DEFINITION_ENUMERATION":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_DEFINITION_ENUMERATION); ok {
										inst.ATTRIBUTE_DEFINITION_ENUMERATION = append(inst.ATTRIBUTE_DEFINITION_ENUMERATION, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_DEFINITION_INTEGER":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_DEFINITION_INTEGER); ok {
										inst.ATTRIBUTE_DEFINITION_INTEGER = append(inst.ATTRIBUTE_DEFINITION_INTEGER, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_DEFINITION_REAL":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_DEFINITION_REAL); ok {
										inst.ATTRIBUTE_DEFINITION_REAL = append(inst.ATTRIBUTE_DEFINITION_REAL, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_DEFINITION_STRING":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_DEFINITION_STRING); ok {
										inst.ATTRIBUTE_DEFINITION_STRING = append(inst.ATTRIBUTE_DEFINITION_STRING, typedTarget)
									}
								}
							}
						}
					case "ATTRIBUTE_DEFINITION_XHTML":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*ATTRIBUTE_DEFINITION_XHTML); ok {
										inst.ATTRIBUTE_DEFINITION_XHTML = append(inst.ATTRIBUTE_DEFINITION_XHTML, typedTarget)
									}
								}
							}
						}
					}
				case *A_SPEC_OBJECTS:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SPEC_OBJECT":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*SPEC_OBJECT); ok {
										inst.SPEC_OBJECT = append(inst.SPEC_OBJECT, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*SPEC_RELATION); ok {
										inst.SPEC_RELATION = append(inst.SPEC_RELATION, typedTarget)
									}
								}
							}
						}
					}
				case *A_SPEC_RELATION_GROUPS:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RELATION_GROUP":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*RELATION_GROUP); ok {
										inst.RELATION_GROUP = append(inst.RELATION_GROUP, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*RELATION_GROUP_TYPE); ok {
										inst.RELATION_GROUP_TYPE = append(inst.RELATION_GROUP_TYPE, typedTarget)
									}
								}
							}
						}
					case "SPEC_OBJECT_TYPE":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*SPEC_OBJECT_TYPE); ok {
										inst.SPEC_OBJECT_TYPE = append(inst.SPEC_OBJECT_TYPE, typedTarget)
									}
								}
							}
						}
					case "SPEC_RELATION_TYPE":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*SPEC_RELATION_TYPE); ok {
										inst.SPEC_RELATION_TYPE = append(inst.SPEC_RELATION_TYPE, typedTarget)
									}
								}
							}
						}
					case "SPECIFICATION_TYPE":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*SPECIFICATION_TYPE); ok {
										inst.SPECIFICATION_TYPE = append(inst.SPECIFICATION_TYPE, typedTarget)
									}
								}
							}
						}
					}
				case *A_THE_HEADER:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "REQ_IF_HEADER":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*REQ_IF_HEADER); ok {
									inst.REQ_IF_HEADER = typedTarget
								}
							}
						}
					}
				case *A_TOOL_EXTENSIONS:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "REQ_IF_TOOL_EXTENSION":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*REQ_IF_TOOL_EXTENSION); ok {
										inst.REQ_IF_TOOL_EXTENSION = append(inst.REQ_IF_TOOL_EXTENSION, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "SPECIFIED_VALUES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPECIFIED_VALUES); ok {
									inst.SPECIFIED_VALUES = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "PROPERTIES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_PROPERTIES); ok {
									inst.PROPERTIES = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "SOURCE_SPECIFICATION":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SOURCE_SPECIFICATION_1); ok {
									inst.SOURCE_SPECIFICATION = typedTarget
								}
							}
						}
					case "SPEC_RELATIONS":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_RELATION_REF); ok {
									inst.SPEC_RELATIONS = typedTarget
								}
							}
						}
					case "TARGET_SPECIFICATION":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SOURCE_SPECIFICATION_1); ok {
									inst.TARGET_SPECIFICATION = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_RELATION_GROUP_TYPE_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "SPEC_ATTRIBUTES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_ATTRIBUTES); ok {
									inst.SPEC_ATTRIBUTES = typedTarget
								}
							}
						}
					}
				case *REQ_IF:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "THE_HEADER":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_THE_HEADER); ok {
									inst.THE_HEADER = typedTarget
								}
							}
						}
					case "CORE_CONTENT":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_CORE_CONTENT); ok {
									inst.CORE_CONTENT = typedTarget
								}
							}
						}
					case "TOOL_EXTENSIONS":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_TOOL_EXTENSIONS); ok {
									inst.TOOL_EXTENSIONS = typedTarget
								}
							}
						}
					}
				case *REQ_IF_CONTENT:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DATATYPES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_DATATYPES); ok {
									inst.DATATYPES = typedTarget
								}
							}
						}
					case "SPEC_TYPES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_TYPES); ok {
									inst.SPEC_TYPES = typedTarget
								}
							}
						}
					case "SPEC_OBJECTS":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_OBJECTS); ok {
									inst.SPEC_OBJECTS = typedTarget
								}
							}
						}
					case "SPEC_RELATIONS":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_RELATIONS); ok {
									inst.SPEC_RELATIONS = typedTarget
								}
							}
						}
					case "SPECIFICATIONS":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPECIFICATIONS); ok {
									inst.SPECIFICATIONS = typedTarget
								}
							}
						}
					case "SPEC_RELATION_GROUPS":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_RELATION_GROUPS); ok {
									inst.SPEC_RELATION_GROUPS = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPECIFICATION_TYPE_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
					case "CHILDREN":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_CHILDREN); ok {
									inst.CHILDREN = typedTarget
								}
							}
						}
					case "VALUES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_XHTML_1); ok {
									inst.VALUES = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "SPEC_ATTRIBUTES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_ATTRIBUTES); ok {
									inst.SPEC_ATTRIBUTES = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "OBJECT":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_OBJECT); ok {
									inst.OBJECT = typedTarget
								}
							}
						}
					case "CHILDREN":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_CHILDREN); ok {
									inst.CHILDREN = typedTarget
								}
							}
						}
					case "EDITABLE_ATTS":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_EDITABLE_ATTS); ok {
									inst.EDITABLE_ATTS = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "VALUES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_XHTML_1); ok {
									inst.VALUES = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_OBJECT_TYPE_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "SPEC_ATTRIBUTES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_ATTRIBUTES); ok {
									inst.SPEC_ATTRIBUTES = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "VALUES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ATTRIBUTE_VALUE_XHTML_1); ok {
									inst.VALUES = typedTarget
								}
							}
						}
					case "SOURCE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SOURCE_1); ok {
									inst.SOURCE = typedTarget
								}
							}
						}
					case "TARGET":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SOURCE_1); ok {
									inst.TARGET = typedTarget
								}
							}
						}
					case "TYPE":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_RELATION_TYPE_REF); ok {
									inst.TYPE = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_ALTERNATIVE_ID); ok {
									inst.ALTERNATIVE_ID = typedTarget
								}
							}
						}
					case "SPEC_ATTRIBUTES":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_SPEC_ATTRIBUTES); ok {
									inst.SPEC_ATTRIBUTES = typedTarget
								}
							}
						}
					}
				case *StaticWebSite:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MarkdownContent":
						inst.MarkdownContent = GongExtractString(rhs)
					case "Chapters":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*StaticWebSiteChapter); ok {
										inst.Chapters = append(inst.Chapters, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*StaticWebSiteParagraph); ok {
										inst.Paragraphs = append(inst.Paragraphs, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StaticWebSiteImage); ok {
									inst.Image = typedTarget
								}
							}
						}
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
		}
		return true
	})

	return nil
}
