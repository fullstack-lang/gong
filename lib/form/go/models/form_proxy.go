package models

type FormElementProxyInterface interface {
	Updated()
}

type SliderNumber interface {
	int | float64
}

type intSliderProxy struct {
	field    *FormFieldInt
	valueRef *int
	onUpdate func()
}

func (p *intSliderProxy) Updated() {
	if p.valueRef != nil {
		*p.valueRef = p.field.Value
	}
	if p.onUpdate != nil {
		p.onUpdate()
	}
}

type float64SliderProxy struct {
	field    *FormFieldFloat64
	valueRef *float64
	onUpdate func()
}

func (p *float64SliderProxy) Updated() {
	if p.valueRef != nil {
		*p.valueRef = p.field.Value
	}
	if p.onUpdate != nil {
		p.onUpdate()
	}
}

type checkBoxProxy struct {
	checkBox *CheckBox
	valueRef *bool
	onUpdate func()
}

func (p *checkBoxProxy) Updated() {
	if p.valueRef != nil {
		*p.valueRef = p.checkBox.Value
	}
	if p.onUpdate != nil {
		p.onUpdate()
	}
}

type selectProxy struct {
	selectField *FormFieldSelect
	onSelect    func(selected string)
}

func (p *selectProxy) Updated() {
	if p.onSelect != nil {
		selected := ""
		if p.selectField.Value != nil {
			selected = p.selectField.Value.Name
		}
		p.onSelect(selected)
	}
}

// NewLiveFormGroup creates and stages a FormGroup without a Submit button for live reactive controls.
func NewLiveFormGroup(stage *Stage, name string) *FormGroup {
	return (&FormGroup{
		Name:             name,
		HideSubmitButton: true,
	}).Stage(stage)
}

// StartAccordion appends and stages a FormDiv starting a Material Accordion panel.
func StartAccordion(stage *Stage, formGroup *FormGroup, accordionName string, isExpandedInitially bool) *FormDiv {
	div := (&FormDiv{
		Name:                   accordionName,
		IsAStartAccordionGroup: true,
		AccordionGroupName:     accordionName,
		IsExpandedInitially:    isExpandedInitially,
	}).Stage(stage)
	formGroup.FormDivs = append(formGroup.FormDivs, div)
	return div
}

// EndAccordion appends and stages a FormDiv ending the current Material Accordion group.
func EndAccordion(stage *Stage, formGroup *FormGroup) *FormDiv {
	div := (&FormDiv{
		Name:                 "end-accordion",
		IsAEndAccordionGroup: true,
	}).Stage(stage)
	formGroup.FormDivs = append(formGroup.FormDivs, div)
	return div
}

// AppendDivider appends and stages a visual divider FormDiv.
func AppendDivider(stage *Stage, formGroup *FormGroup) *FormDiv {
	div := (&FormDiv{
		Name:       "divider",
		IsADivider: true,
	}).Stage(stage)
	formGroup.FormDivs = append(formGroup.FormDivs, div)
	return div
}

// AppendSlider creates, stages, and appends a live-updating int or float64 slider field to formGroup.
func AppendSlider[T SliderNumber](
	stage *Stage,
	formGroup *FormGroup,
	name string,
	min T,
	max T,
	step T,
	valueRef *T,
	onUpdate func(),
) *FormField {
	div := (&FormDiv{
		Name: name,
	}).Stage(stage)
	formGroup.FormDivs = append(formGroup.FormDivs, div)

	formField := (&FormField{
		Name:  name,
		Label: name,
	}).Stage(stage)
	div.FormFields = append(div.FormFields, formField)

	switch vRef := any(valueRef).(type) {
	case *int:
		fieldInt := (&FormFieldInt{
			Name:            name,
			Value:           *vRef,
			HasMinValidator: true,
			MinValue:        any(min).(int),
			HasMaxValidator: true,
			MaxValue:        any(max).(int),
			IsSlider:        true,
			Step:            any(step).(int),
			IsLiveUpdate:    true,
		}).Stage(stage)
		fieldInt.Proxy = &intSliderProxy{
			field:    fieldInt,
			valueRef: vRef,
			onUpdate: onUpdate,
		}
		formField.FormFieldInt = fieldInt
	case *float64:
		fieldFloat := (&FormFieldFloat64{
			Name:            name,
			Value:           *vRef,
			HasMinValidator: true,
			MinValue:        any(min).(float64),
			HasMaxValidator: true,
			MaxValue:        any(max).(float64),
			IsSlider:        true,
			Step:            any(step).(float64),
			IsLiveUpdate:    true,
		}).Stage(stage)
		fieldFloat.Proxy = &float64SliderProxy{
			field:    fieldFloat,
			valueRef: vRef,
			onUpdate: onUpdate,
		}
		formField.FormFieldFloat64 = fieldFloat
	}

	return formField
}

// AppendCheckBox creates, stages, and appends a live-updating boolean CheckBox to formGroup.
func AppendCheckBox(
	stage *Stage,
	formGroup *FormGroup,
	name string,
	valueRef *bool,
	onUpdate func(),
) *CheckBox {
	div := (&FormDiv{
		Name: name,
	}).Stage(stage)
	formGroup.FormDivs = append(formGroup.FormDivs, div)

	cb := (&CheckBox{
		Name:         name,
		Value:        *valueRef,
		IsLiveUpdate: true,
	}).Stage(stage)
	cb.Proxy = &checkBoxProxy{
		checkBox: cb,
		valueRef: valueRef,
		onUpdate: onUpdate,
	}
	div.CheckBoxs = append(div.CheckBoxs, cb)
	return cb
}

// AppendRadioCheckBox creates, stages, and appends a live-updating boolean radio group (LabelForFalse / LabelForTrue) to formGroup.
func AppendRadioCheckBox(
	stage *Stage,
	formGroup *FormGroup,
	name string,
	labelForTrue string,
	labelForFalse string,
	valueRef *bool,
	onUpdate func(),
) *CheckBox {
	cb := AppendCheckBox(stage, formGroup, name, valueRef, onUpdate)
	cb.LabelForTrue = labelForTrue
	cb.LabelForFalse = labelForFalse
	return cb
}

// AppendSelect creates, stages, and appends a live-updating dropdown select field to formGroup.
func AppendSelect(
	stage *Stage,
	formGroup *FormGroup,
	name string,
	label string,
	optionNames []string,
	currentOptionName string,
	onSelect func(selected string),
) *FormField {
	div := (&FormDiv{
		Name: name,
	}).Stage(stage)
	formGroup.FormDivs = append(formGroup.FormDivs, div)

	formField := (&FormField{
		Name:  name,
		Label: label,
	}).Stage(stage)
	div.FormFields = append(div.FormFields, formField)

	sel := (&FormFieldSelect{
		Name:                 name,
		PreserveInitialOrder: true,
		IsLiveUpdate:         true,
	}).Stage(stage)
	sel.Proxy = &selectProxy{
		selectField: sel,
		onSelect:    onSelect,
	}

	for _, optName := range optionNames {
		opt := (&Option{
			Name: optName,
		}).Stage(stage)
		if optName == currentOptionName {
			sel.Value = opt
		}
		sel.Options = append(sel.Options, opt)
	}

	formField.FormFieldSelect = sel
	return formField
}
