// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/app/xsd/tests/reqif/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/reqif/go/models/xmlxsd"
)

type FormCallbackIF = gongprobe.FormCallbackIF

type functionalStageSetFormCallback struct {
	onSave func()
}

func (cb *functionalStageSetFormCallback) OnSave() {
	if cb.onSave != nil {
		cb.onSave()
	}
}

func saveStageSet_ALTERNATIVE_ID_Stage(
	inst *models.ALTERNATIVE_ID,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(
	inst *models.ATTRIBUTE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_BOOLEAN](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_DATE_Stage(
	inst *models.ATTRIBUTE_DEFINITION_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_DATE](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_DATE_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(
	inst *models.ATTRIBUTE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "MULTI_VALUED":
			FormDivBasicFieldToField(&inst.MULTI_VALUED, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_ENUMERATION](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_INTEGER_Stage(
	inst *models.ATTRIBUTE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_INTEGER](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_INTEGER_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_REAL_Stage(
	inst *models.ATTRIBUTE_DEFINITION_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_REAL](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_REAL_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_STRING_Stage(
	inst *models.ATTRIBUTE_DEFINITION_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_STRING](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_STRING_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_DEFINITION_XHTML_Stage(
	inst *models.ATTRIBUTE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "DEFAULT_VALUE":
			StageSetFormDivSelectFieldToField(&inst.DEFAULT_VALUE, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_XHTML_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_BOOLEAN_Stage(
	inst *models.ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_DATE_Stage(
	inst *models.ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_DATE_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_ENUMERATION_Stage(
	inst *models.ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](), formDiv)
		case "VALUES":
			StageSetFormDivSelectFieldToField(&inst.VALUES, probe.stageOfInterest.GetInstancesSet[*models.A_ENUM_VALUE_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_INTEGER_Stage(
	inst *models.ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_REAL_Stage(
	inst *models.ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_REAL_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_STRING_Stage(
	inst *models.ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "THE_VALUE":
			FormDivBasicFieldToField(&inst.THE_VALUE, formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_STRING_REF](), formDiv)
		}
	}
}

func saveStageSet_ATTRIBUTE_VALUE_XHTML_Stage(
	inst *models.ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IS_SIMPLIFIED":
			FormDivBasicFieldToField(&inst.IS_SIMPLIFIED, formDiv)
		case "THE_VALUE":
			StageSetFormDivSelectFieldToField(&inst.THE_VALUE, probe.stageOfInterest.GetInstancesSet[*models.XHTML_CONTENT](), formDiv)
		case "THE_ORIGINAL_VALUE":
			StageSetFormDivSelectFieldToField(&inst.THE_ORIGINAL_VALUE, probe.stageOfInterest.GetInstancesSet[*models.XHTML_CONTENT](), formDiv)
		case "DEFINITION":
			StageSetFormDivSelectFieldToField(&inst.DEFINITION, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF](), formDiv)
		}
	}
}

func saveStageSet_A_ALTERNATIVE_ID_Stage(
	inst *models.A_ALTERNATIVE_ID,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_BOOLEAN_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_DATE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_DATE_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_ENUMERATION_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_INTEGER_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_REAL_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_REAL_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_STRING_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_STRING_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(
	inst *models.A_ATTRIBUTE_DEFINITION_XHTML_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_XHTML_REF, formDiv)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(
	inst *models.A_ATTRIBUTE_VALUE_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_VALUE_BOOLEAN":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_BOOLEAN", &inst.ATTRIBUTE_VALUE_BOOLEAN, formDiv, probe)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_VALUE_DATE_Stage(
	inst *models.A_ATTRIBUTE_VALUE_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_VALUE_DATE":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_DATE", &inst.ATTRIBUTE_VALUE_DATE, formDiv, probe)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(
	inst *models.A_ATTRIBUTE_VALUE_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_VALUE_ENUMERATION":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_ENUMERATION", &inst.ATTRIBUTE_VALUE_ENUMERATION, formDiv, probe)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_VALUE_INTEGER_Stage(
	inst *models.A_ATTRIBUTE_VALUE_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_VALUE_INTEGER":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_INTEGER", &inst.ATTRIBUTE_VALUE_INTEGER, formDiv, probe)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_VALUE_REAL_Stage(
	inst *models.A_ATTRIBUTE_VALUE_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_VALUE_REAL":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_REAL", &inst.ATTRIBUTE_VALUE_REAL, formDiv, probe)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_VALUE_STRING_Stage(
	inst *models.A_ATTRIBUTE_VALUE_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_VALUE_STRING":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_STRING", &inst.ATTRIBUTE_VALUE_STRING, formDiv, probe)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_VALUE_XHTML_Stage(
	inst *models.A_ATTRIBUTE_VALUE_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_VALUE_XHTML":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_XHTML", &inst.ATTRIBUTE_VALUE_XHTML, formDiv, probe)
		}
	}
}

func saveStageSet_A_ATTRIBUTE_VALUE_XHTML_1_Stage(
	inst *models.A_ATTRIBUTE_VALUE_XHTML_1,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_VALUE_BOOLEAN":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_BOOLEAN", &inst.ATTRIBUTE_VALUE_BOOLEAN, formDiv, probe)
		case "ATTRIBUTE_VALUE_DATE":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_DATE", &inst.ATTRIBUTE_VALUE_DATE, formDiv, probe)
		case "ATTRIBUTE_VALUE_ENUMERATION":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_ENUMERATION", &inst.ATTRIBUTE_VALUE_ENUMERATION, formDiv, probe)
		case "ATTRIBUTE_VALUE_INTEGER":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_INTEGER", &inst.ATTRIBUTE_VALUE_INTEGER, formDiv, probe)
		case "ATTRIBUTE_VALUE_REAL":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_REAL", &inst.ATTRIBUTE_VALUE_REAL, formDiv, probe)
		case "ATTRIBUTE_VALUE_STRING":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_STRING", &inst.ATTRIBUTE_VALUE_STRING, formDiv, probe)
		case "ATTRIBUTE_VALUE_XHTML":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_VALUE_XHTML", &inst.ATTRIBUTE_VALUE_XHTML, formDiv, probe)
		}
	}
}

func saveStageSet_A_CHILDREN_Stage(
	inst *models.A_CHILDREN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_HIERARCHY":
			FormDivSliceOfPointersToField(inst, "SPEC_HIERARCHY", &inst.SPEC_HIERARCHY, formDiv, probe)
		}
	}
}

func saveStageSet_A_CORE_CONTENT_Stage(
	inst *models.A_CORE_CONTENT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "REQ_IF_CONTENT":
			StageSetFormDivSelectFieldToField(&inst.REQ_IF_CONTENT, probe.stageOfInterest.GetInstancesSet[*models.REQ_IF_CONTENT](), formDiv)
		}
	}
}

func saveStageSet_A_DATATYPES_Stage(
	inst *models.A_DATATYPES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_BOOLEAN":
			FormDivSliceOfPointersToField(inst, "DATATYPE_DEFINITION_BOOLEAN", &inst.DATATYPE_DEFINITION_BOOLEAN, formDiv, probe)
		case "DATATYPE_DEFINITION_DATE":
			FormDivSliceOfPointersToField(inst, "DATATYPE_DEFINITION_DATE", &inst.DATATYPE_DEFINITION_DATE, formDiv, probe)
		case "DATATYPE_DEFINITION_ENUMERATION":
			FormDivSliceOfPointersToField(inst, "DATATYPE_DEFINITION_ENUMERATION", &inst.DATATYPE_DEFINITION_ENUMERATION, formDiv, probe)
		case "DATATYPE_DEFINITION_INTEGER":
			FormDivSliceOfPointersToField(inst, "DATATYPE_DEFINITION_INTEGER", &inst.DATATYPE_DEFINITION_INTEGER, formDiv, probe)
		case "DATATYPE_DEFINITION_REAL":
			FormDivSliceOfPointersToField(inst, "DATATYPE_DEFINITION_REAL", &inst.DATATYPE_DEFINITION_REAL, formDiv, probe)
		case "DATATYPE_DEFINITION_STRING":
			FormDivSliceOfPointersToField(inst, "DATATYPE_DEFINITION_STRING", &inst.DATATYPE_DEFINITION_STRING, formDiv, probe)
		case "DATATYPE_DEFINITION_XHTML":
			FormDivSliceOfPointersToField(inst, "DATATYPE_DEFINITION_XHTML", &inst.DATATYPE_DEFINITION_XHTML, formDiv, probe)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_BOOLEAN_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_BOOLEAN_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_DATE_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_DATE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_DATE_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_ENUMERATION_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_ENUMERATION_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_INTEGER_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_INTEGER_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_REAL_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_REAL_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_REAL_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_STRING_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_STRING_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_STRING_REF, formDiv)
		}
	}
}

func saveStageSet_A_DATATYPE_DEFINITION_XHTML_REF_Stage(
	inst *models.A_DATATYPE_DEFINITION_XHTML_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&inst.DATATYPE_DEFINITION_XHTML_REF, formDiv)
		}
	}
}

func saveStageSet_A_EDITABLE_ATTS_Stage(
	inst *models.A_EDITABLE_ATTS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_BOOLEAN_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_DATE_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_ENUMERATION_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_INTEGER_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_REAL_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_STRING_REF, formDiv)
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			FormDivBasicFieldToField(&inst.ATTRIBUTE_DEFINITION_XHTML_REF, formDiv)
		}
	}
}

func saveStageSet_A_ENUM_VALUE_REF_Stage(
	inst *models.A_ENUM_VALUE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ENUM_VALUE_REF":
			FormDivBasicFieldToField(&inst.ENUM_VALUE_REF, formDiv)
		}
	}
}

func saveStageSet_A_OBJECT_Stage(
	inst *models.A_OBJECT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&inst.SPEC_OBJECT_REF, formDiv)
		}
	}
}

func saveStageSet_A_PROPERTIES_Stage(
	inst *models.A_PROPERTIES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "EMBEDDED_VALUE":
			StageSetFormDivSelectFieldToField(&inst.EMBEDDED_VALUE, probe.stageOfInterest.GetInstancesSet[*models.EMBEDDED_VALUE](), formDiv)
		}
	}
}

func saveStageSet_A_RELATION_GROUP_TYPE_REF_Stage(
	inst *models.A_RELATION_GROUP_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RELATION_GROUP_TYPE_REF":
			FormDivBasicFieldToField(&inst.RELATION_GROUP_TYPE_REF, formDiv)
		}
	}
}

func saveStageSet_A_SOURCE_Stage(
	inst *models.A_SOURCE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_OBJECT_REF":
			FormDivBasicFieldToField(&inst.SPEC_OBJECT_REF, formDiv)
		}
	}
}

func saveStageSet_A_SOURCE_SPECIFICATION_Stage(
	inst *models.A_SOURCE_SPECIFICATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPECIFICATION_REF":
			FormDivBasicFieldToField(&inst.SPECIFICATION_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPECIFICATIONS_Stage(
	inst *models.A_SPECIFICATIONS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPECIFICATION":
			FormDivSliceOfPointersToField(inst, "SPECIFICATION", &inst.SPECIFICATION, formDiv, probe)
		}
	}
}

func saveStageSet_A_SPECIFICATION_TYPE_REF_Stage(
	inst *models.A_SPECIFICATION_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPECIFICATION_TYPE_REF":
			FormDivBasicFieldToField(&inst.SPECIFICATION_TYPE_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPECIFIED_VALUES_Stage(
	inst *models.A_SPECIFIED_VALUES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ENUM_VALUE":
			FormDivSliceOfPointersToField(inst, "ENUM_VALUE", &inst.ENUM_VALUE, formDiv, probe)
		}
	}
}

func saveStageSet_A_SPEC_ATTRIBUTES_Stage(
	inst *models.A_SPEC_ATTRIBUTES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_DEFINITION_BOOLEAN", &inst.ATTRIBUTE_DEFINITION_BOOLEAN, formDiv, probe)
		case "ATTRIBUTE_DEFINITION_DATE":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_DEFINITION_DATE", &inst.ATTRIBUTE_DEFINITION_DATE, formDiv, probe)
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_DEFINITION_ENUMERATION", &inst.ATTRIBUTE_DEFINITION_ENUMERATION, formDiv, probe)
		case "ATTRIBUTE_DEFINITION_INTEGER":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_DEFINITION_INTEGER", &inst.ATTRIBUTE_DEFINITION_INTEGER, formDiv, probe)
		case "ATTRIBUTE_DEFINITION_REAL":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_DEFINITION_REAL", &inst.ATTRIBUTE_DEFINITION_REAL, formDiv, probe)
		case "ATTRIBUTE_DEFINITION_STRING":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_DEFINITION_STRING", &inst.ATTRIBUTE_DEFINITION_STRING, formDiv, probe)
		case "ATTRIBUTE_DEFINITION_XHTML":
			FormDivSliceOfPointersToField(inst, "ATTRIBUTE_DEFINITION_XHTML", &inst.ATTRIBUTE_DEFINITION_XHTML, formDiv, probe)
		}
	}
}

func saveStageSet_A_SPEC_OBJECTS_Stage(
	inst *models.A_SPEC_OBJECTS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_OBJECT":
			FormDivSliceOfPointersToField(inst, "SPEC_OBJECT", &inst.SPEC_OBJECT, formDiv, probe)
		}
	}
}

func saveStageSet_A_SPEC_OBJECT_TYPE_REF_Stage(
	inst *models.A_SPEC_OBJECT_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_OBJECT_TYPE_REF":
			FormDivBasicFieldToField(&inst.SPEC_OBJECT_TYPE_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPEC_RELATIONS_Stage(
	inst *models.A_SPEC_RELATIONS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_RELATION":
			FormDivSliceOfPointersToField(inst, "SPEC_RELATION", &inst.SPEC_RELATION, formDiv, probe)
		}
	}
}

func saveStageSet_A_SPEC_RELATION_GROUPS_Stage(
	inst *models.A_SPEC_RELATION_GROUPS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RELATION_GROUP":
			FormDivSliceOfPointersToField(inst, "RELATION_GROUP", &inst.RELATION_GROUP, formDiv, probe)
		}
	}
}

func saveStageSet_A_SPEC_RELATION_REF_Stage(
	inst *models.A_SPEC_RELATION_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_RELATION_REF":
			FormDivBasicFieldToField(&inst.SPEC_RELATION_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPEC_RELATION_TYPE_REF_Stage(
	inst *models.A_SPEC_RELATION_TYPE_REF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "SPEC_RELATION_TYPE_REF":
			FormDivBasicFieldToField(&inst.SPEC_RELATION_TYPE_REF, formDiv)
		}
	}
}

func saveStageSet_A_SPEC_TYPES_Stage(
	inst *models.A_SPEC_TYPES,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "RELATION_GROUP_TYPE":
			FormDivSliceOfPointersToField(inst, "RELATION_GROUP_TYPE", &inst.RELATION_GROUP_TYPE, formDiv, probe)
		case "SPEC_OBJECT_TYPE":
			FormDivSliceOfPointersToField(inst, "SPEC_OBJECT_TYPE", &inst.SPEC_OBJECT_TYPE, formDiv, probe)
		case "SPEC_RELATION_TYPE":
			FormDivSliceOfPointersToField(inst, "SPEC_RELATION_TYPE", &inst.SPEC_RELATION_TYPE, formDiv, probe)
		case "SPECIFICATION_TYPE":
			FormDivSliceOfPointersToField(inst, "SPECIFICATION_TYPE", &inst.SPECIFICATION_TYPE, formDiv, probe)
		}
	}
}

func saveStageSet_A_THE_HEADER_Stage(
	inst *models.A_THE_HEADER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "REQ_IF_HEADER":
			StageSetFormDivSelectFieldToField(&inst.REQ_IF_HEADER, probe.stageOfInterest.GetInstancesSet[*models.REQ_IF_HEADER](), formDiv)
		}
	}
}

func saveStageSet_A_TOOL_EXTENSIONS_Stage(
	inst *models.A_TOOL_EXTENSIONS,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "REQ_IF_TOOL_EXTENSION":
			FormDivSliceOfPointersToField(inst, "REQ_IF_TOOL_EXTENSION", &inst.REQ_IF_TOOL_EXTENSION, formDiv, probe)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_BOOLEAN_Stage(
	inst *models.DATATYPE_DEFINITION_BOOLEAN,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_DATE_Stage(
	inst *models.DATATYPE_DEFINITION_DATE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_ENUMERATION_Stage(
	inst *models.DATATYPE_DEFINITION_ENUMERATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPECIFIED_VALUES":
			StageSetFormDivSelectFieldToField(&inst.SPECIFIED_VALUES, probe.stageOfInterest.GetInstancesSet[*models.A_SPECIFIED_VALUES](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_INTEGER_Stage(
	inst *models.DATATYPE_DEFINITION_INTEGER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "MAX_XSD":
			FormDivBasicFieldToField(&inst.MAX_XSD, formDiv)
		case "MIN_XSD":
			FormDivBasicFieldToField(&inst.MIN_XSD, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_REAL_Stage(
	inst *models.DATATYPE_DEFINITION_REAL,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "ACCURACY":
			FormDivBasicFieldToField(&inst.ACCURACY, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "MAX_XSD":
			FormDivBasicFieldToField(&inst.MAX_XSD, formDiv)
		case "MIN_XSD":
			FormDivBasicFieldToField(&inst.MIN_XSD, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_STRING_Stage(
	inst *models.DATATYPE_DEFINITION_STRING,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "MAX_LENGTH":
			FormDivBasicFieldToField(&inst.MAX_LENGTH, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_DATATYPE_DEFINITION_XHTML_Stage(
	inst *models.DATATYPE_DEFINITION_XHTML,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		}
	}
}

func saveStageSet_EMBEDDED_VALUE_Stage(
	inst *models.EMBEDDED_VALUE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "KEY":
			FormDivBasicFieldToField(&inst.KEY, formDiv)
		case "OTHER_CONTENT":
			FormDivBasicFieldToField(&inst.OTHER_CONTENT, formDiv)
		}
	}
}

func saveStageSet_ENUM_VALUE_Stage(
	inst *models.ENUM_VALUE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "PROPERTIES":
			StageSetFormDivSelectFieldToField(&inst.PROPERTIES, probe.stageOfInterest.GetInstancesSet[*models.A_PROPERTIES](), formDiv)
		}
	}
}

func saveStageSet_RELATION_GROUP_Stage(
	inst *models.RELATION_GROUP,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SOURCE_SPECIFICATION":
			StageSetFormDivSelectFieldToField(&inst.SOURCE_SPECIFICATION, probe.stageOfInterest.GetInstancesSet[*models.A_SOURCE_SPECIFICATION](), formDiv)
		case "SPEC_RELATIONS":
			StageSetFormDivSelectFieldToField(&inst.SPEC_RELATIONS, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_RELATION_REF](), formDiv)
		case "TARGET_SPECIFICATION":
			StageSetFormDivSelectFieldToField(&inst.TARGET_SPECIFICATION, probe.stageOfInterest.GetInstancesSet[*models.A_SOURCE_SPECIFICATION](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_RELATION_GROUP_TYPE_REF](), formDiv)
		}
	}
}

func saveStageSet_RELATION_GROUP_TYPE_Stage(
	inst *models.RELATION_GROUP_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPEC_ATTRIBUTES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_ATTRIBUTES, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), formDiv)
		}
	}
}

func saveStageSet_REQ_IF_Stage(
	inst *models.REQ_IF,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "Lang":
			FormDivBasicFieldToField(&inst.Lang, formDiv)
		case "THE_HEADER":
			StageSetFormDivSelectFieldToField(&inst.THE_HEADER, probe.stageOfInterest.GetInstancesSet[*models.A_THE_HEADER](), formDiv)
		case "CORE_CONTENT":
			StageSetFormDivSelectFieldToField(&inst.CORE_CONTENT, probe.stageOfInterest.GetInstancesSet[*models.A_CORE_CONTENT](), formDiv)
		case "TOOL_EXTENSIONS":
			StageSetFormDivSelectFieldToField(&inst.TOOL_EXTENSIONS, probe.stageOfInterest.GetInstancesSet[*models.A_TOOL_EXTENSIONS](), formDiv)
		}
	}
}

func saveStageSet_REQ_IF_CONTENT_Stage(
	inst *models.REQ_IF_CONTENT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DATATYPES":
			StageSetFormDivSelectFieldToField(&inst.DATATYPES, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPES](), formDiv)
		case "SPEC_TYPES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_TYPES, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_TYPES](), formDiv)
		case "SPEC_OBJECTS":
			StageSetFormDivSelectFieldToField(&inst.SPEC_OBJECTS, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_OBJECTS](), formDiv)
		case "SPEC_RELATIONS":
			StageSetFormDivSelectFieldToField(&inst.SPEC_RELATIONS, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_RELATIONS](), formDiv)
		case "SPECIFICATIONS":
			StageSetFormDivSelectFieldToField(&inst.SPECIFICATIONS, probe.stageOfInterest.GetInstancesSet[*models.A_SPECIFICATIONS](), formDiv)
		case "SPEC_RELATION_GROUPS":
			StageSetFormDivSelectFieldToField(&inst.SPEC_RELATION_GROUPS, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_RELATION_GROUPS](), formDiv)
		}
	}
}

func saveStageSet_REQ_IF_HEADER_Stage(
	inst *models.REQ_IF_HEADER,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "COMMENT":
			FormDivBasicFieldToField(&inst.COMMENT, formDiv)
		case "CREATION_TIME":
			FormDivBasicFieldToField(&inst.CREATION_TIME, formDiv)
		case "REPOSITORY_ID":
			FormDivBasicFieldToField(&inst.REPOSITORY_ID, formDiv)
		case "REQ_IF_TOOL_ID":
			FormDivBasicFieldToField(&inst.REQ_IF_TOOL_ID, formDiv)
		case "REQ_IF_VERSION":
			FormDivBasicFieldToField(&inst.REQ_IF_VERSION, formDiv)
		case "SOURCE_TOOL_ID":
			FormDivBasicFieldToField(&inst.SOURCE_TOOL_ID, formDiv)
		case "TITLE":
			FormDivBasicFieldToField(&inst.TITLE, formDiv)
		}
	}
}

func saveStageSet_REQ_IF_TOOL_EXTENSION_Stage(
	inst *models.REQ_IF_TOOL_EXTENSION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}

func saveStageSet_SPECIFICATION_Stage(
	inst *models.SPECIFICATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "CHILDREN":
			StageSetFormDivSelectFieldToField(&inst.CHILDREN, probe.stageOfInterest.GetInstancesSet[*models.A_CHILDREN](), formDiv)
		case "VALUES":
			StageSetFormDivSelectFieldToField(&inst.VALUES, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_SPECIFICATION_TYPE_REF](), formDiv)
		}
	}
}

func saveStageSet_SPECIFICATION_TYPE_Stage(
	inst *models.SPECIFICATION_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPEC_ATTRIBUTES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_ATTRIBUTES, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), formDiv)
		}
	}
}

func saveStageSet_SPEC_HIERARCHY_Stage(
	inst *models.SPEC_HIERARCHY,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "IS_EDITABLE":
			FormDivBasicFieldToField(&inst.IS_EDITABLE, formDiv)
		case "IS_TABLE_INTERNAL":
			FormDivBasicFieldToField(&inst.IS_TABLE_INTERNAL, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "CHILDREN":
			StageSetFormDivSelectFieldToField(&inst.CHILDREN, probe.stageOfInterest.GetInstancesSet[*models.A_CHILDREN](), formDiv)
		case "EDITABLE_ATTS":
			StageSetFormDivSelectFieldToField(&inst.EDITABLE_ATTS, probe.stageOfInterest.GetInstancesSet[*models.A_EDITABLE_ATTS](), formDiv)
		case "OBJECT":
			StageSetFormDivSelectFieldToField(&inst.OBJECT, probe.stageOfInterest.GetInstancesSet[*models.A_OBJECT](), formDiv)
		}
	}
}

func saveStageSet_SPEC_OBJECT_Stage(
	inst *models.SPEC_OBJECT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "VALUES":
			StageSetFormDivSelectFieldToField(&inst.VALUES, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_OBJECT_TYPE_REF](), formDiv)
		}
	}
}

func saveStageSet_SPEC_OBJECT_TYPE_Stage(
	inst *models.SPEC_OBJECT_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPEC_ATTRIBUTES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_ATTRIBUTES, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), formDiv)
		}
	}
}

func saveStageSet_SPEC_RELATION_Stage(
	inst *models.SPEC_RELATION,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "VALUES":
			StageSetFormDivSelectFieldToField(&inst.VALUES, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1](), formDiv)
		case "SOURCE":
			StageSetFormDivSelectFieldToField(&inst.SOURCE, probe.stageOfInterest.GetInstancesSet[*models.A_SOURCE](), formDiv)
		case "TARGET":
			StageSetFormDivSelectFieldToField(&inst.TARGET, probe.stageOfInterest.GetInstancesSet[*models.A_SOURCE](), formDiv)
		case "TYPE_XSD":
			StageSetFormDivSelectFieldToField(&inst.TYPE_XSD, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_RELATION_TYPE_REF](), formDiv)
		}
	}
}

func saveStageSet_SPEC_RELATION_TYPE_Stage(
	inst *models.SPEC_RELATION_TYPE,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "DESC":
			FormDivBasicFieldToField(&inst.DESC, formDiv)
		case "IDENTIFIER":
			FormDivBasicFieldToField(&inst.IDENTIFIER, formDiv)
		case "LAST_CHANGE":
			FormDivBasicFieldToField(&inst.LAST_CHANGE, formDiv)
		case "LONG_NAME":
			FormDivBasicFieldToField(&inst.LONG_NAME, formDiv)
		case "ALTERNATIVE_ID":
			StageSetFormDivSelectFieldToField(&inst.ALTERNATIVE_ID, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), formDiv)
		case "SPEC_ATTRIBUTES":
			StageSetFormDivSelectFieldToField(&inst.SPEC_ATTRIBUTES, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), formDiv)
		}
	}
}

func saveStageSet_XHTML_CONTENT_Stage(
	inst *models.XHTML_CONTENT,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&inst.EnclosedText, formDiv)
		}
	}
}

func saveStageSet_Xmlxsd_SchemaElement_XmlxsdStage(
	inst *xmlxsd.Xmlxsd_SchemaElement,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		case "Name":
			FormDivBasicFieldToField(&inst.Name, formDiv)
		}
	}
}


func StageSetNewInstance_ALTERNATIVE_ID_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ALTERNATIVE_ID",
	}).Stage(probe.formStage)
	inst := new(models.ALTERNATIVE_ID)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ALTERNATIVE_ID_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ALTERNATIVE_ID_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_DEFINITION_BOOLEAN",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_BOOLEAN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_BOOLEAN_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_DEFINITION_DATE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_DEFINITION_DATE",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_DATE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_DEFINITION_DATE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_DATE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_DEFINITION_ENUMERATION",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_ENUMERATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_ENUMERATION_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_DEFINITION_INTEGER_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_DEFINITION_INTEGER",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_INTEGER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_DEFINITION_INTEGER_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_INTEGER_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_DEFINITION_REAL_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_DEFINITION_REAL",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_REAL)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_DEFINITION_REAL_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_REAL_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_DEFINITION_STRING_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_DEFINITION_STRING",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_STRING)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_DEFINITION_STRING_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_STRING_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_DEFINITION_XHTML_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_DEFINITION_XHTML",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_DEFINITION_XHTML)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_DEFINITION_XHTML_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_DEFINITION_XHTML_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_VALUE_BOOLEAN",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_BOOLEAN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_VALUE_BOOLEAN_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_VALUE_DATE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_VALUE_DATE",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_DATE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_VALUE_DATE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_DATE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_VALUE_ENUMERATION",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_ENUMERATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_VALUE_ENUMERATION_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_VALUE_INTEGER_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_VALUE_INTEGER",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_INTEGER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_VALUE_INTEGER_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_VALUE_REAL_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_VALUE_REAL",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_REAL)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_VALUE_REAL_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_REAL_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_VALUE_STRING_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_VALUE_STRING",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_STRING)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_VALUE_STRING_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_STRING_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ATTRIBUTE_VALUE_XHTML_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ATTRIBUTE_VALUE_XHTML",
	}).Stage(probe.formStage)
	inst := new(models.ATTRIBUTE_VALUE_XHTML)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ATTRIBUTE_VALUE_XHTML_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ATTRIBUTE_VALUE_XHTML_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ALTERNATIVE_ID_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ALTERNATIVE_ID",
	}).Stage(probe.formStage)
	inst := new(models.A_ALTERNATIVE_ID)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ALTERNATIVE_ID_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ALTERNATIVE_ID_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_DEFINITION_BOOLEAN_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_DEFINITION_DATE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_DATE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_DATE_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_DEFINITION_ENUMERATION_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_DEFINITION_INTEGER_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_INTEGER_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_DEFINITION_REAL_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_REAL_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_REAL_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_DEFINITION_STRING_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_STRING_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_STRING_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_DEFINITION_XHTML_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_DEFINITION_XHTML_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_DEFINITION_XHTML_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_VALUE_BOOLEAN",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_BOOLEAN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_BOOLEAN_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_VALUE_DATE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_VALUE_DATE",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_DATE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_VALUE_DATE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_DATE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_VALUE_ENUMERATION",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_ENUMERATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_ENUMERATION_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_VALUE_INTEGER_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_VALUE_INTEGER",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_INTEGER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_VALUE_INTEGER_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_INTEGER_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_VALUE_REAL_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_VALUE_REAL",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_REAL)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_VALUE_REAL_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_REAL_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_VALUE_STRING_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_VALUE_STRING",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_STRING)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_VALUE_STRING_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_STRING_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_VALUE_XHTML_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_VALUE_XHTML",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_XHTML)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_VALUE_XHTML_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_XHTML_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ATTRIBUTE_VALUE_XHTML_1_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ATTRIBUTE_VALUE_XHTML_1",
	}).Stage(probe.formStage)
	inst := new(models.A_ATTRIBUTE_VALUE_XHTML_1)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ATTRIBUTE_VALUE_XHTML_1_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ATTRIBUTE_VALUE_XHTML_1_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_CHILDREN_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_CHILDREN",
	}).Stage(probe.formStage)
	inst := new(models.A_CHILDREN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_CHILDREN_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_CHILDREN_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_CORE_CONTENT_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_CORE_CONTENT",
	}).Stage(probe.formStage)
	inst := new(models.A_CORE_CONTENT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_CORE_CONTENT_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_CORE_CONTENT_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_DATATYPES_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_DATATYPES",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_DATATYPES_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_DATATYPES_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_DATATYPE_DEFINITION_BOOLEAN_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_BOOLEAN_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_BOOLEAN_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_DATATYPE_DEFINITION_DATE_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_DATATYPE_DEFINITION_DATE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_DATE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_DATATYPE_DEFINITION_DATE_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_DATE_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_DATATYPE_DEFINITION_ENUMERATION_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_ENUMERATION_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_ENUMERATION_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_DATATYPE_DEFINITION_INTEGER_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_INTEGER_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_INTEGER_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_DATATYPE_DEFINITION_REAL_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_DATATYPE_DEFINITION_REAL_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_REAL_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_DATATYPE_DEFINITION_REAL_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_REAL_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_DATATYPE_DEFINITION_STRING_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_DATATYPE_DEFINITION_STRING_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_STRING_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_DATATYPE_DEFINITION_STRING_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_STRING_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_DATATYPE_DEFINITION_XHTML_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_DATATYPE_DEFINITION_XHTML_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_DATATYPE_DEFINITION_XHTML_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_DATATYPE_DEFINITION_XHTML_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_DATATYPE_DEFINITION_XHTML_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_EDITABLE_ATTS_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_EDITABLE_ATTS",
	}).Stage(probe.formStage)
	inst := new(models.A_EDITABLE_ATTS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_EDITABLE_ATTS_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_EDITABLE_ATTS_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_ENUM_VALUE_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_ENUM_VALUE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_ENUM_VALUE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_ENUM_VALUE_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_ENUM_VALUE_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_OBJECT_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_OBJECT",
	}).Stage(probe.formStage)
	inst := new(models.A_OBJECT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_OBJECT_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_OBJECT_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_PROPERTIES_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_PROPERTIES",
	}).Stage(probe.formStage)
	inst := new(models.A_PROPERTIES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_PROPERTIES_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_PROPERTIES_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_RELATION_GROUP_TYPE_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_RELATION_GROUP_TYPE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_RELATION_GROUP_TYPE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_RELATION_GROUP_TYPE_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_RELATION_GROUP_TYPE_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SOURCE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SOURCE",
	}).Stage(probe.formStage)
	inst := new(models.A_SOURCE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SOURCE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SOURCE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SOURCE_SPECIFICATION_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SOURCE_SPECIFICATION",
	}).Stage(probe.formStage)
	inst := new(models.A_SOURCE_SPECIFICATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SOURCE_SPECIFICATION_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SOURCE_SPECIFICATION_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPECIFICATIONS_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPECIFICATIONS",
	}).Stage(probe.formStage)
	inst := new(models.A_SPECIFICATIONS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPECIFICATIONS_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPECIFICATIONS_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPECIFICATION_TYPE_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPECIFICATION_TYPE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_SPECIFICATION_TYPE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPECIFICATION_TYPE_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPECIFICATION_TYPE_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPECIFIED_VALUES_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPECIFIED_VALUES",
	}).Stage(probe.formStage)
	inst := new(models.A_SPECIFIED_VALUES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPECIFIED_VALUES_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPECIFIED_VALUES_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPEC_ATTRIBUTES_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPEC_ATTRIBUTES",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_ATTRIBUTES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPEC_ATTRIBUTES_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPEC_ATTRIBUTES_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPEC_OBJECTS_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPEC_OBJECTS",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_OBJECTS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPEC_OBJECTS_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPEC_OBJECTS_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPEC_OBJECT_TYPE_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPEC_OBJECT_TYPE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_OBJECT_TYPE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPEC_OBJECT_TYPE_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPEC_OBJECT_TYPE_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPEC_RELATIONS_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPEC_RELATIONS",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_RELATIONS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPEC_RELATIONS_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPEC_RELATIONS_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPEC_RELATION_GROUPS_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPEC_RELATION_GROUPS",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_RELATION_GROUPS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPEC_RELATION_GROUPS_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPEC_RELATION_GROUPS_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPEC_RELATION_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPEC_RELATION_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_RELATION_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPEC_RELATION_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPEC_RELATION_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPEC_RELATION_TYPE_REF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPEC_RELATION_TYPE_REF",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_RELATION_TYPE_REF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPEC_RELATION_TYPE_REF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPEC_RELATION_TYPE_REF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_SPEC_TYPES_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_SPEC_TYPES",
	}).Stage(probe.formStage)
	inst := new(models.A_SPEC_TYPES)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_SPEC_TYPES_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_SPEC_TYPES_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_THE_HEADER_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_THE_HEADER",
	}).Stage(probe.formStage)
	inst := new(models.A_THE_HEADER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_THE_HEADER_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_THE_HEADER_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_A_TOOL_EXTENSIONS_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New A_TOOL_EXTENSIONS",
	}).Stage(probe.formStage)
	inst := new(models.A_TOOL_EXTENSIONS)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_A_TOOL_EXTENSIONS_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_A_TOOL_EXTENSIONS_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DATATYPE_DEFINITION_BOOLEAN_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DATATYPE_DEFINITION_BOOLEAN",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_BOOLEAN)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_DATATYPE_DEFINITION_BOOLEAN_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_BOOLEAN_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DATATYPE_DEFINITION_DATE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DATATYPE_DEFINITION_DATE",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_DATE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_DATATYPE_DEFINITION_DATE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_DATE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DATATYPE_DEFINITION_ENUMERATION_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DATATYPE_DEFINITION_ENUMERATION",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_ENUMERATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_DATATYPE_DEFINITION_ENUMERATION_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_ENUMERATION_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DATATYPE_DEFINITION_INTEGER_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DATATYPE_DEFINITION_INTEGER",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_INTEGER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_DATATYPE_DEFINITION_INTEGER_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_INTEGER_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DATATYPE_DEFINITION_REAL_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DATATYPE_DEFINITION_REAL",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_REAL)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_DATATYPE_DEFINITION_REAL_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_REAL_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DATATYPE_DEFINITION_STRING_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DATATYPE_DEFINITION_STRING",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_STRING)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_DATATYPE_DEFINITION_STRING_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_STRING_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_DATATYPE_DEFINITION_XHTML_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New DATATYPE_DEFINITION_XHTML",
	}).Stage(probe.formStage)
	inst := new(models.DATATYPE_DEFINITION_XHTML)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_DATATYPE_DEFINITION_XHTML_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_DATATYPE_DEFINITION_XHTML_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_EMBEDDED_VALUE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New EMBEDDED_VALUE",
	}).Stage(probe.formStage)
	inst := new(models.EMBEDDED_VALUE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_EMBEDDED_VALUE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_EMBEDDED_VALUE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_ENUM_VALUE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New ENUM_VALUE",
	}).Stage(probe.formStage)
	inst := new(models.ENUM_VALUE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_ENUM_VALUE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_ENUM_VALUE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RELATION_GROUP_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RELATION_GROUP",
	}).Stage(probe.formStage)
	inst := new(models.RELATION_GROUP)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_RELATION_GROUP_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_RELATION_GROUP_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_RELATION_GROUP_TYPE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New RELATION_GROUP_TYPE",
	}).Stage(probe.formStage)
	inst := new(models.RELATION_GROUP_TYPE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_RELATION_GROUP_TYPE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_RELATION_GROUP_TYPE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_REQ_IF_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New REQ_IF",
	}).Stage(probe.formStage)
	inst := new(models.REQ_IF)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_REQ_IF_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_REQ_IF_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_REQ_IF_CONTENT_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New REQ_IF_CONTENT",
	}).Stage(probe.formStage)
	inst := new(models.REQ_IF_CONTENT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_REQ_IF_CONTENT_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_REQ_IF_CONTENT_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_REQ_IF_HEADER_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New REQ_IF_HEADER",
	}).Stage(probe.formStage)
	inst := new(models.REQ_IF_HEADER)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_REQ_IF_HEADER_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_REQ_IF_HEADER_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_REQ_IF_TOOL_EXTENSION_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New REQ_IF_TOOL_EXTENSION",
	}).Stage(probe.formStage)
	inst := new(models.REQ_IF_TOOL_EXTENSION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_REQ_IF_TOOL_EXTENSION_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_REQ_IF_TOOL_EXTENSION_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SPECIFICATION_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SPECIFICATION",
	}).Stage(probe.formStage)
	inst := new(models.SPECIFICATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_SPECIFICATION_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_SPECIFICATION_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SPECIFICATION_TYPE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SPECIFICATION_TYPE",
	}).Stage(probe.formStage)
	inst := new(models.SPECIFICATION_TYPE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_SPECIFICATION_TYPE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_SPECIFICATION_TYPE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SPEC_HIERARCHY_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SPEC_HIERARCHY",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_HIERARCHY)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_SPEC_HIERARCHY_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_SPEC_HIERARCHY_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SPEC_OBJECT_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SPEC_OBJECT",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_OBJECT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_SPEC_OBJECT_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_SPEC_OBJECT_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SPEC_OBJECT_TYPE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SPEC_OBJECT_TYPE",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_OBJECT_TYPE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_SPEC_OBJECT_TYPE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_SPEC_OBJECT_TYPE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SPEC_RELATION_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SPEC_RELATION",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_RELATION)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_SPEC_RELATION_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_SPEC_RELATION_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_SPEC_RELATION_TYPE_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New SPEC_RELATION_TYPE",
	}).Stage(probe.formStage)
	inst := new(models.SPEC_RELATION_TYPE)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_SPEC_RELATION_TYPE_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_SPEC_RELATION_TYPE_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_XHTML_CONTENT_Stage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New XHTML_CONTENT",
	}).Stage(probe.formStage)
	inst := new(models.XHTML_CONTENT)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.Lock()
			defer probe.stageOfInterest.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest)
			saveStageSet_XHTML_CONTENT_Stage(inst, probe, formGroup)
			probe.stageOfInterest.Commit()
			updateStageSetTable_XHTML_CONTENT_Stage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

func StageSetNewInstance_Xmlxsd_SchemaElement_XmlxsdStage(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New Xmlxsd_SchemaElement",
	}).Stage(probe.formStage)
	inst := new(xmlxsd.Xmlxsd_SchemaElement)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			probe.stageOfInterest.XmlxsdStage.Lock()
			defer probe.stageOfInterest.XmlxsdStage.Unlock()
			probe.formStage.Checkout()
			inst.Stage(probe.stageOfInterest.XmlxsdStage)
			saveStageSet_Xmlxsd_SchemaElement_XmlxsdStage(inst, probe, formGroup)
			probe.stageOfInterest.XmlxsdStage.Commit()
			updateStageSetTable_Xmlxsd_SchemaElement_XmlxsdStage(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}

