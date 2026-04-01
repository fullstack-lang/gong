package models

import (
	"fmt"
)

func (a *ATTRIBUTE_VALUE_XHTML) GetAttributeDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF
}
func (a *ATTRIBUTE_VALUE_STRING) GetAttributeDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF
}
func (a *ATTRIBUTE_VALUE_BOOLEAN) GetAttributeDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF
}
func (a *ATTRIBUTE_VALUE_INTEGER) GetAttributeDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF
}
func (a *ATTRIBUTE_VALUE_DATE) GetAttributeDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF
}
func (a *ATTRIBUTE_VALUE_REAL) GetAttributeDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF
}
func (a *ATTRIBUTE_VALUE_ENUMERATION) GetAttributeDefinitionRef() string {
	return a.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF
}

func (a *ATTRIBUTE_VALUE_XHTML) GetValue() string {
	return a.THE_VALUE.EnclosedText
}
func (a *ATTRIBUTE_VALUE_STRING) GetValue() string {
	return a.THE_VALUE
}
func (a *ATTRIBUTE_VALUE_BOOLEAN) GetValue() string {
	res := "false"

	if a.THE_VALUE {
		res = "true"
	}

	return res
}
func (a *ATTRIBUTE_VALUE_INTEGER) GetValue() string {
	return fmt.Sprintf("%d", a.THE_VALUE)
}
func (a *ATTRIBUTE_VALUE_DATE) GetValue() string {
	return fmt.Sprintf("%s", a.THE_VALUE)
}
func (a *ATTRIBUTE_VALUE_REAL) GetValue() string {
	return fmt.Sprintf("%f", a.THE_VALUE)
}
func (a *ATTRIBUTE_VALUE_ENUMERATION) GetValue() string {
	return a.VALUES.ENUM_VALUE_REF
}
