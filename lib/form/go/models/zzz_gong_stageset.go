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
		checkboxOrdered := []*CheckBox{}
		for checkbox := range stageSet.Stage.CheckBoxs {
			checkboxOrdered = append(checkboxOrdered, checkbox)
		}
		sort.Slice(checkboxOrdered, func(i, j int) bool {
			return stageSet.Stage.CheckBox_stagedOrder[checkboxOrdered[i]] < stageSet.Stage.CheckBox_stagedOrder[checkboxOrdered[j]]
		})
		for _, checkbox := range checkboxOrdered {
			checkboxIdent := "__stage_0" + checkbox.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.CheckBox{Name: %s}).Stage(stageSet.Stage)", checkboxIdent, __gong__toRawStringLiteral(checkbox.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", checkboxIdent, __gong__toRawStringLiteral(checkbox.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %t", checkboxIdent, checkbox.Value))
		}
	}
	if stageSet.Stage != nil {
		formdivOrdered := []*FormDiv{}
		for formdiv := range stageSet.Stage.FormDivs {
			formdivOrdered = append(formdivOrdered, formdiv)
		}
		sort.Slice(formdivOrdered, func(i, j int) bool {
			return stageSet.Stage.FormDiv_stagedOrder[formdivOrdered[i]] < stageSet.Stage.FormDiv_stagedOrder[formdivOrdered[j]]
		})
		for _, formdiv := range formdivOrdered {
			formdivIdent := "__stage_0" + formdiv.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormDiv{Name: %s}).Stage(stageSet.Stage)", formdivIdent, __gong__toRawStringLiteral(formdiv.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formdivIdent, __gong__toRawStringLiteral(formdiv.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsADivider = %t", formdivIdent, formdiv.IsADivider))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAStartAccordionGroup = %t", formdivIdent, formdiv.IsAStartAccordionGroup))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionGroupName = %s", formdivIdent, __gong__toRawStringLiteral(formdiv.AccordionGroupName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAEndAccordionGroup = %t", formdivIdent, formdiv.IsAEndAccordionGroup))
			for _, elem := range formdiv.FormFields {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFields = append(%s.FormFields, %s)", formdivIdent, formdivIdent, targetIdent))
			}
			for _, elem := range formdiv.CheckBoxs {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CheckBoxs = append(%s.CheckBoxs, %s)", formdivIdent, formdivIdent, targetIdent))
			}
			if formdiv.FormEditAssocButton != nil {
				targetIdent := "__stage_0" + formdiv.FormEditAssocButton.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormEditAssocButton = %s", formdivIdent, targetIdent))
			}
			if formdiv.FormSortAssocButton != nil {
				targetIdent := "__stage_0" + formdiv.FormSortAssocButton.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormSortAssocButton = %s", formdivIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		formeditassocbuttonOrdered := []*FormEditAssocButton{}
		for formeditassocbutton := range stageSet.Stage.FormEditAssocButtons {
			formeditassocbuttonOrdered = append(formeditassocbuttonOrdered, formeditassocbutton)
		}
		sort.Slice(formeditassocbuttonOrdered, func(i, j int) bool {
			return stageSet.Stage.FormEditAssocButton_stagedOrder[formeditassocbuttonOrdered[i]] < stageSet.Stage.FormEditAssocButton_stagedOrder[formeditassocbuttonOrdered[j]]
		})
		for _, formeditassocbutton := range formeditassocbuttonOrdered {
			formeditassocbuttonIdent := "__stage_0" + formeditassocbutton.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormEditAssocButton{Name: %s}).Stage(stageSet.Stage)", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.AssociationStorage = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.AssociationStorage)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasChanged = %t", formeditassocbuttonIdent, formeditassocbutton.HasChanged))
			values.WriteString(fmt.Sprintf("\n\t%s.IsForSavePurpose = %t", formeditassocbuttonIdent, formeditassocbutton.IsForSavePurpose))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", formeditassocbuttonIdent, formeditassocbutton.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.MatTooltipShowDelay = %s", formeditassocbuttonIdent, __gong__toRawStringLiteral(formeditassocbutton.MatTooltipShowDelay)))
		}
	}
	if stageSet.Stage != nil {
		formfieldOrdered := []*FormField{}
		for formfield := range stageSet.Stage.FormFields {
			formfieldOrdered = append(formfieldOrdered, formfield)
		}
		sort.Slice(formfieldOrdered, func(i, j int) bool {
			return stageSet.Stage.FormField_stagedOrder[formfieldOrdered[i]] < stageSet.Stage.FormField_stagedOrder[formfieldOrdered[j]]
		})
		for _, formfield := range formfieldOrdered {
			formfieldIdent := "__stage_0" + formfield.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormField{Name: %s}).Stage(stageSet.Stage)", formfieldIdent, __gong__toRawStringLiteral(formfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldIdent, __gong__toRawStringLiteral(formfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.InputTypeEnum = %s", formfieldIdent, __gong__toRawStringLiteral(string(formfield.InputTypeEnum))))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", formfieldIdent, __gong__toRawStringLiteral(formfield.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placeholder = %s", formfieldIdent, __gong__toRawStringLiteral(formfield.Placeholder)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasBespokeWidth = %t", formfieldIdent, formfield.HasBespokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeWidthPx = %d", formfieldIdent, formfield.BespokeWidthPx))
			values.WriteString(fmt.Sprintf("\n\t%s.HasBespokeHeight = %t", formfieldIdent, formfield.HasBespokeHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeHeightPx = %d", formfieldIdent, formfield.BespokeHeightPx))
			if formfield.FormFieldString != nil {
				targetIdent := "__stage_0" + formfield.FormFieldString.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldString = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldFloat64 != nil {
				targetIdent := "__stage_0" + formfield.FormFieldFloat64.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldFloat64 = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldInt != nil {
				targetIdent := "__stage_0" + formfield.FormFieldInt.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldInt = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldDate != nil {
				targetIdent := "__stage_0" + formfield.FormFieldDate.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldDate = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldTime != nil {
				targetIdent := "__stage_0" + formfield.FormFieldTime.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldTime = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldDateTime != nil {
				targetIdent := "__stage_0" + formfield.FormFieldDateTime.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldDateTime = %s", formfieldIdent, targetIdent))
			}
			if formfield.FormFieldSelect != nil {
				targetIdent := "__stage_0" + formfield.FormFieldSelect.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormFieldSelect = %s", formfieldIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		formfielddateOrdered := []*FormFieldDate{}
		for formfielddate := range stageSet.Stage.FormFieldDates {
			formfielddateOrdered = append(formfielddateOrdered, formfielddate)
		}
		sort.Slice(formfielddateOrdered, func(i, j int) bool {
			return stageSet.Stage.FormFieldDate_stagedOrder[formfielddateOrdered[i]] < stageSet.Stage.FormFieldDate_stagedOrder[formfielddateOrdered[j]]
		})
		for _, formfielddate := range formfielddateOrdered {
			formfielddateIdent := "__stage_0" + formfielddate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormFieldDate{Name: %s}).Stage(stageSet.Stage)", formfielddateIdent, __gong__toRawStringLiteral(formfielddate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfielddateIdent, __gong__toRawStringLiteral(formfielddate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", formfielddateIdent, formfielddate.Value.String()))
		}
	}
	if stageSet.Stage != nil {
		formfielddatetimeOrdered := []*FormFieldDateTime{}
		for formfielddatetime := range stageSet.Stage.FormFieldDateTimes {
			formfielddatetimeOrdered = append(formfielddatetimeOrdered, formfielddatetime)
		}
		sort.Slice(formfielddatetimeOrdered, func(i, j int) bool {
			return stageSet.Stage.FormFieldDateTime_stagedOrder[formfielddatetimeOrdered[i]] < stageSet.Stage.FormFieldDateTime_stagedOrder[formfielddatetimeOrdered[j]]
		})
		for _, formfielddatetime := range formfielddatetimeOrdered {
			formfielddatetimeIdent := "__stage_0" + formfielddatetime.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormFieldDateTime{Name: %s}).Stage(stageSet.Stage)", formfielddatetimeIdent, __gong__toRawStringLiteral(formfielddatetime.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfielddatetimeIdent, __gong__toRawStringLiteral(formfielddatetime.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", formfielddatetimeIdent, formfielddatetime.Value.String()))
		}
	}
	if stageSet.Stage != nil {
		formfieldfloat64Ordered := []*FormFieldFloat64{}
		for formfieldfloat64 := range stageSet.Stage.FormFieldFloat64s {
			formfieldfloat64Ordered = append(formfieldfloat64Ordered, formfieldfloat64)
		}
		sort.Slice(formfieldfloat64Ordered, func(i, j int) bool {
			return stageSet.Stage.FormFieldFloat64_stagedOrder[formfieldfloat64Ordered[i]] < stageSet.Stage.FormFieldFloat64_stagedOrder[formfieldfloat64Ordered[j]]
		})
		for _, formfieldfloat64 := range formfieldfloat64Ordered {
			formfieldfloat64Ident := "__stage_0" + formfieldfloat64.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormFieldFloat64{Name: %s}).Stage(stageSet.Stage)", formfieldfloat64Ident, __gong__toRawStringLiteral(formfieldfloat64.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldfloat64Ident, __gong__toRawStringLiteral(formfieldfloat64.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %f", formfieldfloat64Ident, formfieldfloat64.Value))
			values.WriteString(fmt.Sprintf("\n\t%s.HasMinValidator = %t", formfieldfloat64Ident, formfieldfloat64.HasMinValidator))
			values.WriteString(fmt.Sprintf("\n\t%s.MinValue = %f", formfieldfloat64Ident, formfieldfloat64.MinValue))
			values.WriteString(fmt.Sprintf("\n\t%s.HasMaxValidator = %t", formfieldfloat64Ident, formfieldfloat64.HasMaxValidator))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxValue = %f", formfieldfloat64Ident, formfieldfloat64.MaxValue))
		}
	}
	if stageSet.Stage != nil {
		formfieldintOrdered := []*FormFieldInt{}
		for formfieldint := range stageSet.Stage.FormFieldInts {
			formfieldintOrdered = append(formfieldintOrdered, formfieldint)
		}
		sort.Slice(formfieldintOrdered, func(i, j int) bool {
			return stageSet.Stage.FormFieldInt_stagedOrder[formfieldintOrdered[i]] < stageSet.Stage.FormFieldInt_stagedOrder[formfieldintOrdered[j]]
		})
		for _, formfieldint := range formfieldintOrdered {
			formfieldintIdent := "__stage_0" + formfieldint.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormFieldInt{Name: %s}).Stage(stageSet.Stage)", formfieldintIdent, __gong__toRawStringLiteral(formfieldint.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldintIdent, __gong__toRawStringLiteral(formfieldint.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", formfieldintIdent, formfieldint.Value))
			values.WriteString(fmt.Sprintf("\n\t%s.HasMinValidator = %t", formfieldintIdent, formfieldint.HasMinValidator))
			values.WriteString(fmt.Sprintf("\n\t%s.MinValue = %d", formfieldintIdent, formfieldint.MinValue))
			values.WriteString(fmt.Sprintf("\n\t%s.HasMaxValidator = %t", formfieldintIdent, formfieldint.HasMaxValidator))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxValue = %d", formfieldintIdent, formfieldint.MaxValue))
		}
	}
	if stageSet.Stage != nil {
		formfieldselectOrdered := []*FormFieldSelect{}
		for formfieldselect := range stageSet.Stage.FormFieldSelects {
			formfieldselectOrdered = append(formfieldselectOrdered, formfieldselect)
		}
		sort.Slice(formfieldselectOrdered, func(i, j int) bool {
			return stageSet.Stage.FormFieldSelect_stagedOrder[formfieldselectOrdered[i]] < stageSet.Stage.FormFieldSelect_stagedOrder[formfieldselectOrdered[j]]
		})
		for _, formfieldselect := range formfieldselectOrdered {
			formfieldselectIdent := "__stage_0" + formfieldselect.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormFieldSelect{Name: %s}).Stage(stageSet.Stage)", formfieldselectIdent, __gong__toRawStringLiteral(formfieldselect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldselectIdent, __gong__toRawStringLiteral(formfieldselect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CanBeEmpty = %t", formfieldselectIdent, formfieldselect.CanBeEmpty))
			values.WriteString(fmt.Sprintf("\n\t%s.PreserveInitialOrder = %t", formfieldselectIdent, formfieldselect.PreserveInitialOrder))
			if formfieldselect.Value != nil {
				targetIdent := "__stage_0" + formfieldselect.Value.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Value = %s", formfieldselectIdent, targetIdent))
			}
			for _, elem := range formfieldselect.Options {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Options = append(%s.Options, %s)", formfieldselectIdent, formfieldselectIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		formfieldstringOrdered := []*FormFieldString{}
		for formfieldstring := range stageSet.Stage.FormFieldStrings {
			formfieldstringOrdered = append(formfieldstringOrdered, formfieldstring)
		}
		sort.Slice(formfieldstringOrdered, func(i, j int) bool {
			return stageSet.Stage.FormFieldString_stagedOrder[formfieldstringOrdered[i]] < stageSet.Stage.FormFieldString_stagedOrder[formfieldstringOrdered[j]]
		})
		for _, formfieldstring := range formfieldstringOrdered {
			formfieldstringIdent := "__stage_0" + formfieldstring.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormFieldString{Name: %s}).Stage(stageSet.Stage)", formfieldstringIdent, __gong__toRawStringLiteral(formfieldstring.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldstringIdent, __gong__toRawStringLiteral(formfieldstring.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", formfieldstringIdent, __gong__toRawStringLiteral(formfieldstring.Value)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTextArea = %t", formfieldstringIdent, formfieldstring.IsTextArea))
		}
	}
	if stageSet.Stage != nil {
		formfieldtimeOrdered := []*FormFieldTime{}
		for formfieldtime := range stageSet.Stage.FormFieldTimes {
			formfieldtimeOrdered = append(formfieldtimeOrdered, formfieldtime)
		}
		sort.Slice(formfieldtimeOrdered, func(i, j int) bool {
			return stageSet.Stage.FormFieldTime_stagedOrder[formfieldtimeOrdered[i]] < stageSet.Stage.FormFieldTime_stagedOrder[formfieldtimeOrdered[j]]
		})
		for _, formfieldtime := range formfieldtimeOrdered {
			formfieldtimeIdent := "__stage_0" + formfieldtime.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormFieldTime{Name: %s}).Stage(stageSet.Stage)", formfieldtimeIdent, __gong__toRawStringLiteral(formfieldtime.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formfieldtimeIdent, __gong__toRawStringLiteral(formfieldtime.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", formfieldtimeIdent, formfieldtime.Value.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Step = %f", formfieldtimeIdent, formfieldtime.Step))
		}
	}
	if stageSet.Stage != nil {
		formgroupOrdered := []*FormGroup{}
		for formgroup := range stageSet.Stage.FormGroups {
			formgroupOrdered = append(formgroupOrdered, formgroup)
		}
		sort.Slice(formgroupOrdered, func(i, j int) bool {
			return stageSet.Stage.FormGroup_stagedOrder[formgroupOrdered[i]] < stageSet.Stage.FormGroup_stagedOrder[formgroupOrdered[j]]
		})
		for _, formgroup := range formgroupOrdered {
			formgroupIdent := "__stage_0" + formgroup.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormGroup{Name: %s}).Stage(stageSet.Stage)", formgroupIdent, __gong__toRawStringLiteral(formgroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formgroupIdent, __gong__toRawStringLiteral(formgroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", formgroupIdent, __gong__toRawStringLiteral(formgroup.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.TypeLabel = %s", formgroupIdent, __gong__toRawStringLiteral(formgroup.TypeLabel)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasSuppressButton = %t", formgroupIdent, formgroup.HasSuppressButton))
			values.WriteString(fmt.Sprintf("\n\t%s.HasSuppressButtonBeenPressed = %t", formgroupIdent, formgroup.HasSuppressButtonBeenPressed))
			for _, elem := range formgroup.FormDivs {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormDivs = append(%s.FormDivs, %s)", formgroupIdent, formgroupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		formsortassocbuttonOrdered := []*FormSortAssocButton{}
		for formsortassocbutton := range stageSet.Stage.FormSortAssocButtons {
			formsortassocbuttonOrdered = append(formsortassocbuttonOrdered, formsortassocbutton)
		}
		sort.Slice(formsortassocbuttonOrdered, func(i, j int) bool {
			return stageSet.Stage.FormSortAssocButton_stagedOrder[formsortassocbuttonOrdered[i]] < stageSet.Stage.FormSortAssocButton_stagedOrder[formsortassocbuttonOrdered[j]]
		})
		for _, formsortassocbutton := range formsortassocbuttonOrdered {
			formsortassocbuttonIdent := "__stage_0" + formsortassocbutton.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.FormSortAssocButton{Name: %s}).Stage(stageSet.Stage)", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", formsortassocbuttonIdent, formsortassocbutton.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.MatTooltipShowDelay = %s", formsortassocbuttonIdent, __gong__toRawStringLiteral(formsortassocbutton.MatTooltipShowDelay)))
			if formsortassocbutton.FormEditAssocButton != nil {
				targetIdent := "__stage_0" + formsortassocbutton.FormEditAssocButton.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FormEditAssocButton = %s", formsortassocbuttonIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		optionOrdered := []*Option{}
		for option := range stageSet.Stage.Options {
			optionOrdered = append(optionOrdered, option)
		}
		sort.Slice(optionOrdered, func(i, j int) bool {
			return stageSet.Stage.Option_stagedOrder[optionOrdered[i]] < stageSet.Stage.Option_stagedOrder[optionOrdered[j]]
		})
		for _, option := range optionOrdered {
			optionIdent := "__stage_0" + option.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Option{Name: %s}).Stage(stageSet.Stage)", optionIdent, __gong__toRawStringLiteral(option.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", optionIdent, __gong__toRawStringLiteral(option.Name)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/lib/form/go/models"
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
				case "CheckBox":
					if !preserveOrder {
						inst := (&CheckBox{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(CheckBox)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormDiv":
					if !preserveOrder {
						inst := (&FormDiv{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormDiv)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormEditAssocButton":
					if !preserveOrder {
						inst := (&FormEditAssocButton{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormEditAssocButton)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormField":
					if !preserveOrder {
						inst := (&FormField{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormField)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormFieldDate":
					if !preserveOrder {
						inst := (&FormFieldDate{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormFieldDate)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormFieldDateTime":
					if !preserveOrder {
						inst := (&FormFieldDateTime{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormFieldDateTime)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormFieldFloat64":
					if !preserveOrder {
						inst := (&FormFieldFloat64{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormFieldFloat64)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormFieldInt":
					if !preserveOrder {
						inst := (&FormFieldInt{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormFieldInt)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormFieldSelect":
					if !preserveOrder {
						inst := (&FormFieldSelect{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormFieldSelect)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormFieldString":
					if !preserveOrder {
						inst := (&FormFieldString{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormFieldString)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormFieldTime":
					if !preserveOrder {
						inst := (&FormFieldTime{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormFieldTime)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormGroup":
					if !preserveOrder {
						inst := (&FormGroup{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormGroup)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "FormSortAssocButton":
					if !preserveOrder {
						inst := (&FormSortAssocButton{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(FormSortAssocButton)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Option":
					if !preserveOrder {
						inst := (&Option{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Option)
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
				case *CheckBox:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractBool(rhs)
					}
				case *FormDiv:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FormFields":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormField); ok {
										inst.FormFields = append(inst.FormFields, typedTarget)
									}
								}
							}
						}
					case "CheckBoxs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CheckBox); ok {
										inst.CheckBoxs = append(inst.CheckBoxs, typedTarget)
									}
								}
							}
						}
					case "FormEditAssocButton":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormEditAssocButton); ok {
									inst.FormEditAssocButton = typedTarget
								}
							}
						}
					case "FormSortAssocButton":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormSortAssocButton); ok {
									inst.FormSortAssocButton = typedTarget
								}
							}
						}
					case "IsADivider":
						inst.IsADivider = GongExtractBool(rhs)
					case "IsAStartAccordionGroup":
						inst.IsAStartAccordionGroup = GongExtractBool(rhs)
					case "AccordionGroupName":
						inst.AccordionGroupName = GongExtractString(rhs)
					case "IsAEndAccordionGroup":
						inst.IsAEndAccordionGroup = GongExtractBool(rhs)
					}
				case *FormEditAssocButton:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "AssociationStorage":
						inst.AssociationStorage = GongExtractString(rhs)
					case "HasChanged":
						inst.HasChanged = GongExtractBool(rhs)
					case "IsForSavePurpose":
						inst.IsForSavePurpose = GongExtractBool(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "MatTooltipShowDelay":
						inst.MatTooltipShowDelay = GongExtractString(rhs)
					}
				case *FormField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "InputTypeEnum":
						inst.InputTypeEnum = InputTypeEnum(GongExtractString(rhs))
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Placeholder":
						inst.Placeholder = GongExtractString(rhs)
					case "FormFieldString":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormFieldString); ok {
									inst.FormFieldString = typedTarget
								}
							}
						}
					case "FormFieldFloat64":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormFieldFloat64); ok {
									inst.FormFieldFloat64 = typedTarget
								}
							}
						}
					case "FormFieldInt":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormFieldInt); ok {
									inst.FormFieldInt = typedTarget
								}
							}
						}
					case "FormFieldDate":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormFieldDate); ok {
									inst.FormFieldDate = typedTarget
								}
							}
						}
					case "FormFieldTime":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormFieldTime); ok {
									inst.FormFieldTime = typedTarget
								}
							}
						}
					case "FormFieldDateTime":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormFieldDateTime); ok {
									inst.FormFieldDateTime = typedTarget
								}
							}
						}
					case "FormFieldSelect":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormFieldSelect); ok {
									inst.FormFieldSelect = typedTarget
								}
							}
						}
					case "HasBespokeWidth":
						inst.HasBespokeWidth = GongExtractBool(rhs)
					case "BespokeWidthPx":
						inst.BespokeWidthPx = GongExtractInt(rhs)
					case "HasBespokeHeight":
						inst.HasBespokeHeight = GongExtractBool(rhs)
					case "BespokeHeightPx":
						inst.BespokeHeightPx = GongExtractInt(rhs)
					}
				case *FormFieldDate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					}
				case *FormFieldDateTime:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					}
				case *FormFieldFloat64:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractFloat(rhs)
					case "HasMinValidator":
						inst.HasMinValidator = GongExtractBool(rhs)
					case "MinValue":
						inst.MinValue = GongExtractFloat(rhs)
					case "HasMaxValidator":
						inst.HasMaxValidator = GongExtractBool(rhs)
					case "MaxValue":
						inst.MaxValue = GongExtractFloat(rhs)
					}
				case *FormFieldInt:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractInt(rhs)
					case "HasMinValidator":
						inst.HasMinValidator = GongExtractBool(rhs)
					case "MinValue":
						inst.MinValue = GongExtractInt(rhs)
					case "HasMaxValidator":
						inst.HasMaxValidator = GongExtractBool(rhs)
					case "MaxValue":
						inst.MaxValue = GongExtractInt(rhs)
					}
				case *FormFieldSelect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Option); ok {
									inst.Value = typedTarget
								}
							}
						}
					case "Options":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Option); ok {
										inst.Options = append(inst.Options, typedTarget)
									}
								}
							}
						}
					case "CanBeEmpty":
						inst.CanBeEmpty = GongExtractBool(rhs)
					case "PreserveInitialOrder":
						inst.PreserveInitialOrder = GongExtractBool(rhs)
					}
				case *FormFieldString:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					case "IsTextArea":
						inst.IsTextArea = GongExtractBool(rhs)
					}
				case *FormFieldTime:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Value, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "Step":
						inst.Step = GongExtractFloat(rhs)
					}
				case *FormGroup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "TypeLabel":
						inst.TypeLabel = GongExtractString(rhs)
					case "FormDivs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormDiv); ok {
										inst.FormDivs = append(inst.FormDivs, typedTarget)
									}
								}
							}
						}
					case "HasSuppressButton":
						inst.HasSuppressButton = GongExtractBool(rhs)
					case "HasSuppressButtonBeenPressed":
						inst.HasSuppressButtonBeenPressed = GongExtractBool(rhs)
					}
				case *FormSortAssocButton:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "MatTooltipShowDelay":
						inst.MatTooltipShowDelay = GongExtractString(rhs)
					case "FormEditAssocButton":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*FormEditAssocButton); ok {
									inst.FormEditAssocButton = typedTarget
								}
							}
						}
					}
				case *Option:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
