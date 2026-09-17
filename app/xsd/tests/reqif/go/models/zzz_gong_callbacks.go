// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (alternative_id *ALTERNATIVE_ID) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterALTERNATIVE_IDCreateCallback != nil {
		stage.OnAfterALTERNATIVE_IDCreateCallback.OnAfterCreate(stage, alternative_id)
	}
}

func (alternative_id *ALTERNATIVE_ID) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterALTERNATIVE_IDUpdateCallback != nil {
		var frontALTERNATIVE_ID *ALTERNATIVE_ID
		if front != nil {
			frontALTERNATIVE_ID, _ = front.(*ALTERNATIVE_ID)
		}
		stage.OnAfterALTERNATIVE_IDUpdateCallback.OnAfterUpdate(stage, alternative_id, frontALTERNATIVE_ID)
	}
}

func (alternative_id *ALTERNATIVE_ID) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterALTERNATIVE_IDDeleteCallback != nil {
		var frontALTERNATIVE_ID *ALTERNATIVE_ID
		if front != nil {
			frontALTERNATIVE_ID, _ = front.(*ALTERNATIVE_ID)
		}
		stage.OnAfterALTERNATIVE_IDDeleteCallback.OnAfterDelete(stage, alternative_id, frontALTERNATIVE_ID)
	}
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback != nil {
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback.OnAfterCreate(stage, attribute_definition_boolean)
	}
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback != nil {
		var frontATTRIBUTE_DEFINITION_BOOLEAN *ATTRIBUTE_DEFINITION_BOOLEAN
		if front != nil {
			frontATTRIBUTE_DEFINITION_BOOLEAN, _ = front.(*ATTRIBUTE_DEFINITION_BOOLEAN)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback.OnAfterUpdate(stage, attribute_definition_boolean, frontATTRIBUTE_DEFINITION_BOOLEAN)
	}
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback != nil {
		var frontATTRIBUTE_DEFINITION_BOOLEAN *ATTRIBUTE_DEFINITION_BOOLEAN
		if front != nil {
			frontATTRIBUTE_DEFINITION_BOOLEAN, _ = front.(*ATTRIBUTE_DEFINITION_BOOLEAN)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback.OnAfterDelete(stage, attribute_definition_boolean, frontATTRIBUTE_DEFINITION_BOOLEAN)
	}
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_DEFINITION_DATECreateCallback != nil {
		stage.OnAfterATTRIBUTE_DEFINITION_DATECreateCallback.OnAfterCreate(stage, attribute_definition_date)
	}
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback != nil {
		var frontATTRIBUTE_DEFINITION_DATE *ATTRIBUTE_DEFINITION_DATE
		if front != nil {
			frontATTRIBUTE_DEFINITION_DATE, _ = front.(*ATTRIBUTE_DEFINITION_DATE)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback.OnAfterUpdate(stage, attribute_definition_date, frontATTRIBUTE_DEFINITION_DATE)
	}
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback != nil {
		var frontATTRIBUTE_DEFINITION_DATE *ATTRIBUTE_DEFINITION_DATE
		if front != nil {
			frontATTRIBUTE_DEFINITION_DATE, _ = front.(*ATTRIBUTE_DEFINITION_DATE)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback.OnAfterDelete(stage, attribute_definition_date, frontATTRIBUTE_DEFINITION_DATE)
	}
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback != nil {
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback.OnAfterCreate(stage, attribute_definition_enumeration)
	}
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback != nil {
		var frontATTRIBUTE_DEFINITION_ENUMERATION *ATTRIBUTE_DEFINITION_ENUMERATION
		if front != nil {
			frontATTRIBUTE_DEFINITION_ENUMERATION, _ = front.(*ATTRIBUTE_DEFINITION_ENUMERATION)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, attribute_definition_enumeration, frontATTRIBUTE_DEFINITION_ENUMERATION)
	}
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback != nil {
		var frontATTRIBUTE_DEFINITION_ENUMERATION *ATTRIBUTE_DEFINITION_ENUMERATION
		if front != nil {
			frontATTRIBUTE_DEFINITION_ENUMERATION, _ = front.(*ATTRIBUTE_DEFINITION_ENUMERATION)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback.OnAfterDelete(stage, attribute_definition_enumeration, frontATTRIBUTE_DEFINITION_ENUMERATION)
	}
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback != nil {
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback.OnAfterCreate(stage, attribute_definition_integer)
	}
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback != nil {
		var frontATTRIBUTE_DEFINITION_INTEGER *ATTRIBUTE_DEFINITION_INTEGER
		if front != nil {
			frontATTRIBUTE_DEFINITION_INTEGER, _ = front.(*ATTRIBUTE_DEFINITION_INTEGER)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback.OnAfterUpdate(stage, attribute_definition_integer, frontATTRIBUTE_DEFINITION_INTEGER)
	}
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback != nil {
		var frontATTRIBUTE_DEFINITION_INTEGER *ATTRIBUTE_DEFINITION_INTEGER
		if front != nil {
			frontATTRIBUTE_DEFINITION_INTEGER, _ = front.(*ATTRIBUTE_DEFINITION_INTEGER)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback.OnAfterDelete(stage, attribute_definition_integer, frontATTRIBUTE_DEFINITION_INTEGER)
	}
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_DEFINITION_REALCreateCallback != nil {
		stage.OnAfterATTRIBUTE_DEFINITION_REALCreateCallback.OnAfterCreate(stage, attribute_definition_real)
	}
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback != nil {
		var frontATTRIBUTE_DEFINITION_REAL *ATTRIBUTE_DEFINITION_REAL
		if front != nil {
			frontATTRIBUTE_DEFINITION_REAL, _ = front.(*ATTRIBUTE_DEFINITION_REAL)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback.OnAfterUpdate(stage, attribute_definition_real, frontATTRIBUTE_DEFINITION_REAL)
	}
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback != nil {
		var frontATTRIBUTE_DEFINITION_REAL *ATTRIBUTE_DEFINITION_REAL
		if front != nil {
			frontATTRIBUTE_DEFINITION_REAL, _ = front.(*ATTRIBUTE_DEFINITION_REAL)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback.OnAfterDelete(stage, attribute_definition_real, frontATTRIBUTE_DEFINITION_REAL)
	}
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback != nil {
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback.OnAfterCreate(stage, attribute_definition_string)
	}
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback != nil {
		var frontATTRIBUTE_DEFINITION_STRING *ATTRIBUTE_DEFINITION_STRING
		if front != nil {
			frontATTRIBUTE_DEFINITION_STRING, _ = front.(*ATTRIBUTE_DEFINITION_STRING)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback.OnAfterUpdate(stage, attribute_definition_string, frontATTRIBUTE_DEFINITION_STRING)
	}
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback != nil {
		var frontATTRIBUTE_DEFINITION_STRING *ATTRIBUTE_DEFINITION_STRING
		if front != nil {
			frontATTRIBUTE_DEFINITION_STRING, _ = front.(*ATTRIBUTE_DEFINITION_STRING)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback.OnAfterDelete(stage, attribute_definition_string, frontATTRIBUTE_DEFINITION_STRING)
	}
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback != nil {
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback.OnAfterCreate(stage, attribute_definition_xhtml)
	}
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback != nil {
		var frontATTRIBUTE_DEFINITION_XHTML *ATTRIBUTE_DEFINITION_XHTML
		if front != nil {
			frontATTRIBUTE_DEFINITION_XHTML, _ = front.(*ATTRIBUTE_DEFINITION_XHTML)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback.OnAfterUpdate(stage, attribute_definition_xhtml, frontATTRIBUTE_DEFINITION_XHTML)
	}
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback != nil {
		var frontATTRIBUTE_DEFINITION_XHTML *ATTRIBUTE_DEFINITION_XHTML
		if front != nil {
			frontATTRIBUTE_DEFINITION_XHTML, _ = front.(*ATTRIBUTE_DEFINITION_XHTML)
		}
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback.OnAfterDelete(stage, attribute_definition_xhtml, frontATTRIBUTE_DEFINITION_XHTML)
	}
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback != nil {
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback.OnAfterCreate(stage, attribute_value_boolean)
	}
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback != nil {
		var frontATTRIBUTE_VALUE_BOOLEAN *ATTRIBUTE_VALUE_BOOLEAN
		if front != nil {
			frontATTRIBUTE_VALUE_BOOLEAN, _ = front.(*ATTRIBUTE_VALUE_BOOLEAN)
		}
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback.OnAfterUpdate(stage, attribute_value_boolean, frontATTRIBUTE_VALUE_BOOLEAN)
	}
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback != nil {
		var frontATTRIBUTE_VALUE_BOOLEAN *ATTRIBUTE_VALUE_BOOLEAN
		if front != nil {
			frontATTRIBUTE_VALUE_BOOLEAN, _ = front.(*ATTRIBUTE_VALUE_BOOLEAN)
		}
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback.OnAfterDelete(stage, attribute_value_boolean, frontATTRIBUTE_VALUE_BOOLEAN)
	}
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_VALUE_DATECreateCallback != nil {
		stage.OnAfterATTRIBUTE_VALUE_DATECreateCallback.OnAfterCreate(stage, attribute_value_date)
	}
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback != nil {
		var frontATTRIBUTE_VALUE_DATE *ATTRIBUTE_VALUE_DATE
		if front != nil {
			frontATTRIBUTE_VALUE_DATE, _ = front.(*ATTRIBUTE_VALUE_DATE)
		}
		stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback.OnAfterUpdate(stage, attribute_value_date, frontATTRIBUTE_VALUE_DATE)
	}
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_DATEDeleteCallback != nil {
		var frontATTRIBUTE_VALUE_DATE *ATTRIBUTE_VALUE_DATE
		if front != nil {
			frontATTRIBUTE_VALUE_DATE, _ = front.(*ATTRIBUTE_VALUE_DATE)
		}
		stage.OnAfterATTRIBUTE_VALUE_DATEDeleteCallback.OnAfterDelete(stage, attribute_value_date, frontATTRIBUTE_VALUE_DATE)
	}
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback != nil {
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback.OnAfterCreate(stage, attribute_value_enumeration)
	}
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback != nil {
		var frontATTRIBUTE_VALUE_ENUMERATION *ATTRIBUTE_VALUE_ENUMERATION
		if front != nil {
			frontATTRIBUTE_VALUE_ENUMERATION, _ = front.(*ATTRIBUTE_VALUE_ENUMERATION)
		}
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, attribute_value_enumeration, frontATTRIBUTE_VALUE_ENUMERATION)
	}
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback != nil {
		var frontATTRIBUTE_VALUE_ENUMERATION *ATTRIBUTE_VALUE_ENUMERATION
		if front != nil {
			frontATTRIBUTE_VALUE_ENUMERATION, _ = front.(*ATTRIBUTE_VALUE_ENUMERATION)
		}
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback.OnAfterDelete(stage, attribute_value_enumeration, frontATTRIBUTE_VALUE_ENUMERATION)
	}
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback != nil {
		stage.OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback.OnAfterCreate(stage, attribute_value_integer)
	}
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback != nil {
		var frontATTRIBUTE_VALUE_INTEGER *ATTRIBUTE_VALUE_INTEGER
		if front != nil {
			frontATTRIBUTE_VALUE_INTEGER, _ = front.(*ATTRIBUTE_VALUE_INTEGER)
		}
		stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback.OnAfterUpdate(stage, attribute_value_integer, frontATTRIBUTE_VALUE_INTEGER)
	}
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback != nil {
		var frontATTRIBUTE_VALUE_INTEGER *ATTRIBUTE_VALUE_INTEGER
		if front != nil {
			frontATTRIBUTE_VALUE_INTEGER, _ = front.(*ATTRIBUTE_VALUE_INTEGER)
		}
		stage.OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback.OnAfterDelete(stage, attribute_value_integer, frontATTRIBUTE_VALUE_INTEGER)
	}
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_VALUE_REALCreateCallback != nil {
		stage.OnAfterATTRIBUTE_VALUE_REALCreateCallback.OnAfterCreate(stage, attribute_value_real)
	}
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback != nil {
		var frontATTRIBUTE_VALUE_REAL *ATTRIBUTE_VALUE_REAL
		if front != nil {
			frontATTRIBUTE_VALUE_REAL, _ = front.(*ATTRIBUTE_VALUE_REAL)
		}
		stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback.OnAfterUpdate(stage, attribute_value_real, frontATTRIBUTE_VALUE_REAL)
	}
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_REALDeleteCallback != nil {
		var frontATTRIBUTE_VALUE_REAL *ATTRIBUTE_VALUE_REAL
		if front != nil {
			frontATTRIBUTE_VALUE_REAL, _ = front.(*ATTRIBUTE_VALUE_REAL)
		}
		stage.OnAfterATTRIBUTE_VALUE_REALDeleteCallback.OnAfterDelete(stage, attribute_value_real, frontATTRIBUTE_VALUE_REAL)
	}
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_VALUE_STRINGCreateCallback != nil {
		stage.OnAfterATTRIBUTE_VALUE_STRINGCreateCallback.OnAfterCreate(stage, attribute_value_string)
	}
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback != nil {
		var frontATTRIBUTE_VALUE_STRING *ATTRIBUTE_VALUE_STRING
		if front != nil {
			frontATTRIBUTE_VALUE_STRING, _ = front.(*ATTRIBUTE_VALUE_STRING)
		}
		stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback.OnAfterUpdate(stage, attribute_value_string, frontATTRIBUTE_VALUE_STRING)
	}
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback != nil {
		var frontATTRIBUTE_VALUE_STRING *ATTRIBUTE_VALUE_STRING
		if front != nil {
			frontATTRIBUTE_VALUE_STRING, _ = front.(*ATTRIBUTE_VALUE_STRING)
		}
		stage.OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback.OnAfterDelete(stage, attribute_value_string, frontATTRIBUTE_VALUE_STRING)
	}
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback != nil {
		stage.OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback.OnAfterCreate(stage, attribute_value_xhtml)
	}
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback != nil {
		var frontATTRIBUTE_VALUE_XHTML *ATTRIBUTE_VALUE_XHTML
		if front != nil {
			frontATTRIBUTE_VALUE_XHTML, _ = front.(*ATTRIBUTE_VALUE_XHTML)
		}
		stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback.OnAfterUpdate(stage, attribute_value_xhtml, frontATTRIBUTE_VALUE_XHTML)
	}
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback != nil {
		var frontATTRIBUTE_VALUE_XHTML *ATTRIBUTE_VALUE_XHTML
		if front != nil {
			frontATTRIBUTE_VALUE_XHTML, _ = front.(*ATTRIBUTE_VALUE_XHTML)
		}
		stage.OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback.OnAfterDelete(stage, attribute_value_xhtml, frontATTRIBUTE_VALUE_XHTML)
	}
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ALTERNATIVE_IDCreateCallback != nil {
		stage.OnAfterA_ALTERNATIVE_IDCreateCallback.OnAfterCreate(stage, a_alternative_id)
	}
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ALTERNATIVE_IDUpdateCallback != nil {
		var frontA_ALTERNATIVE_ID *A_ALTERNATIVE_ID
		if front != nil {
			frontA_ALTERNATIVE_ID, _ = front.(*A_ALTERNATIVE_ID)
		}
		stage.OnAfterA_ALTERNATIVE_IDUpdateCallback.OnAfterUpdate(stage, a_alternative_id, frontA_ALTERNATIVE_ID)
	}
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ALTERNATIVE_IDDeleteCallback != nil {
		var frontA_ALTERNATIVE_ID *A_ALTERNATIVE_ID
		if front != nil {
			frontA_ALTERNATIVE_ID, _ = front.(*A_ALTERNATIVE_ID)
		}
		stage.OnAfterA_ALTERNATIVE_IDDeleteCallback.OnAfterDelete(stage, a_alternative_id, frontA_ALTERNATIVE_ID)
	}
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFCreateCallback.OnAfterCreate(stage, a_attribute_definition_boolean_ref)
	}
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_BOOLEAN_REF *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_BOOLEAN_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateCallback.OnAfterUpdate(stage, a_attribute_definition_boolean_ref, frontA_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	}
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDeleteCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_BOOLEAN_REF *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_BOOLEAN_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDeleteCallback.OnAfterDelete(stage, a_attribute_definition_boolean_ref, frontA_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	}
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFCreateCallback.OnAfterCreate(stage, a_attribute_definition_date_ref)
	}
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_DATE_REF *A_ATTRIBUTE_DEFINITION_DATE_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_DATE_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_DATE_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateCallback.OnAfterUpdate(stage, a_attribute_definition_date_ref, frontA_ATTRIBUTE_DEFINITION_DATE_REF)
	}
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFDeleteCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_DATE_REF *A_ATTRIBUTE_DEFINITION_DATE_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_DATE_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_DATE_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFDeleteCallback.OnAfterDelete(stage, a_attribute_definition_date_ref, frontA_ATTRIBUTE_DEFINITION_DATE_REF)
	}
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFCreateCallback.OnAfterCreate(stage, a_attribute_definition_enumeration_ref)
	}
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_ENUMERATION_REF *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_ENUMERATION_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateCallback.OnAfterUpdate(stage, a_attribute_definition_enumeration_ref, frontA_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	}
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDeleteCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_ENUMERATION_REF *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_ENUMERATION_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDeleteCallback.OnAfterDelete(stage, a_attribute_definition_enumeration_ref, frontA_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	}
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFCreateCallback.OnAfterCreate(stage, a_attribute_definition_integer_ref)
	}
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_INTEGER_REF *A_ATTRIBUTE_DEFINITION_INTEGER_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_INTEGER_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateCallback.OnAfterUpdate(stage, a_attribute_definition_integer_ref, frontA_ATTRIBUTE_DEFINITION_INTEGER_REF)
	}
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFDeleteCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_INTEGER_REF *A_ATTRIBUTE_DEFINITION_INTEGER_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_INTEGER_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFDeleteCallback.OnAfterDelete(stage, a_attribute_definition_integer_ref, frontA_ATTRIBUTE_DEFINITION_INTEGER_REF)
	}
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFCreateCallback.OnAfterCreate(stage, a_attribute_definition_real_ref)
	}
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_REAL_REF *A_ATTRIBUTE_DEFINITION_REAL_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_REAL_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_REAL_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateCallback.OnAfterUpdate(stage, a_attribute_definition_real_ref, frontA_ATTRIBUTE_DEFINITION_REAL_REF)
	}
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFDeleteCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_REAL_REF *A_ATTRIBUTE_DEFINITION_REAL_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_REAL_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_REAL_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFDeleteCallback.OnAfterDelete(stage, a_attribute_definition_real_ref, frontA_ATTRIBUTE_DEFINITION_REAL_REF)
	}
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFCreateCallback.OnAfterCreate(stage, a_attribute_definition_string_ref)
	}
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_STRING_REF *A_ATTRIBUTE_DEFINITION_STRING_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_STRING_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_STRING_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateCallback.OnAfterUpdate(stage, a_attribute_definition_string_ref, frontA_ATTRIBUTE_DEFINITION_STRING_REF)
	}
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFDeleteCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_STRING_REF *A_ATTRIBUTE_DEFINITION_STRING_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_STRING_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_STRING_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFDeleteCallback.OnAfterDelete(stage, a_attribute_definition_string_ref, frontA_ATTRIBUTE_DEFINITION_STRING_REF)
	}
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFCreateCallback.OnAfterCreate(stage, a_attribute_definition_xhtml_ref)
	}
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_XHTML_REF *A_ATTRIBUTE_DEFINITION_XHTML_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_XHTML_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateCallback.OnAfterUpdate(stage, a_attribute_definition_xhtml_ref, frontA_ATTRIBUTE_DEFINITION_XHTML_REF)
	}
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFDeleteCallback != nil {
		var frontA_ATTRIBUTE_DEFINITION_XHTML_REF *A_ATTRIBUTE_DEFINITION_XHTML_REF
		if front != nil {
			frontA_ATTRIBUTE_DEFINITION_XHTML_REF, _ = front.(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
		}
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFDeleteCallback.OnAfterDelete(stage, a_attribute_definition_xhtml_ref, frontA_ATTRIBUTE_DEFINITION_XHTML_REF)
	}
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANCreateCallback.OnAfterCreate(stage, a_attribute_value_boolean)
	}
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateCallback != nil {
		var frontA_ATTRIBUTE_VALUE_BOOLEAN *A_ATTRIBUTE_VALUE_BOOLEAN
		if front != nil {
			frontA_ATTRIBUTE_VALUE_BOOLEAN, _ = front.(*A_ATTRIBUTE_VALUE_BOOLEAN)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateCallback.OnAfterUpdate(stage, a_attribute_value_boolean, frontA_ATTRIBUTE_VALUE_BOOLEAN)
	}
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANDeleteCallback != nil {
		var frontA_ATTRIBUTE_VALUE_BOOLEAN *A_ATTRIBUTE_VALUE_BOOLEAN
		if front != nil {
			frontA_ATTRIBUTE_VALUE_BOOLEAN, _ = front.(*A_ATTRIBUTE_VALUE_BOOLEAN)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANDeleteCallback.OnAfterDelete(stage, a_attribute_value_boolean, frontA_ATTRIBUTE_VALUE_BOOLEAN)
	}
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_DATECreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_VALUE_DATECreateCallback.OnAfterCreate(stage, a_attribute_value_date)
	}
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateCallback != nil {
		var frontA_ATTRIBUTE_VALUE_DATE *A_ATTRIBUTE_VALUE_DATE
		if front != nil {
			frontA_ATTRIBUTE_VALUE_DATE, _ = front.(*A_ATTRIBUTE_VALUE_DATE)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateCallback.OnAfterUpdate(stage, a_attribute_value_date, frontA_ATTRIBUTE_VALUE_DATE)
	}
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_DATEDeleteCallback != nil {
		var frontA_ATTRIBUTE_VALUE_DATE *A_ATTRIBUTE_VALUE_DATE
		if front != nil {
			frontA_ATTRIBUTE_VALUE_DATE, _ = front.(*A_ATTRIBUTE_VALUE_DATE)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_DATEDeleteCallback.OnAfterDelete(stage, a_attribute_value_date, frontA_ATTRIBUTE_VALUE_DATE)
	}
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONCreateCallback.OnAfterCreate(stage, a_attribute_value_enumeration)
	}
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateCallback != nil {
		var frontA_ATTRIBUTE_VALUE_ENUMERATION *A_ATTRIBUTE_VALUE_ENUMERATION
		if front != nil {
			frontA_ATTRIBUTE_VALUE_ENUMERATION, _ = front.(*A_ATTRIBUTE_VALUE_ENUMERATION)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, a_attribute_value_enumeration, frontA_ATTRIBUTE_VALUE_ENUMERATION)
	}
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONDeleteCallback != nil {
		var frontA_ATTRIBUTE_VALUE_ENUMERATION *A_ATTRIBUTE_VALUE_ENUMERATION
		if front != nil {
			frontA_ATTRIBUTE_VALUE_ENUMERATION, _ = front.(*A_ATTRIBUTE_VALUE_ENUMERATION)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONDeleteCallback.OnAfterDelete(stage, a_attribute_value_enumeration, frontA_ATTRIBUTE_VALUE_ENUMERATION)
	}
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERCreateCallback.OnAfterCreate(stage, a_attribute_value_integer)
	}
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateCallback != nil {
		var frontA_ATTRIBUTE_VALUE_INTEGER *A_ATTRIBUTE_VALUE_INTEGER
		if front != nil {
			frontA_ATTRIBUTE_VALUE_INTEGER, _ = front.(*A_ATTRIBUTE_VALUE_INTEGER)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateCallback.OnAfterUpdate(stage, a_attribute_value_integer, frontA_ATTRIBUTE_VALUE_INTEGER)
	}
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERDeleteCallback != nil {
		var frontA_ATTRIBUTE_VALUE_INTEGER *A_ATTRIBUTE_VALUE_INTEGER
		if front != nil {
			frontA_ATTRIBUTE_VALUE_INTEGER, _ = front.(*A_ATTRIBUTE_VALUE_INTEGER)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERDeleteCallback.OnAfterDelete(stage, a_attribute_value_integer, frontA_ATTRIBUTE_VALUE_INTEGER)
	}
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_REALCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_VALUE_REALCreateCallback.OnAfterCreate(stage, a_attribute_value_real)
	}
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateCallback != nil {
		var frontA_ATTRIBUTE_VALUE_REAL *A_ATTRIBUTE_VALUE_REAL
		if front != nil {
			frontA_ATTRIBUTE_VALUE_REAL, _ = front.(*A_ATTRIBUTE_VALUE_REAL)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateCallback.OnAfterUpdate(stage, a_attribute_value_real, frontA_ATTRIBUTE_VALUE_REAL)
	}
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_REALDeleteCallback != nil {
		var frontA_ATTRIBUTE_VALUE_REAL *A_ATTRIBUTE_VALUE_REAL
		if front != nil {
			frontA_ATTRIBUTE_VALUE_REAL, _ = front.(*A_ATTRIBUTE_VALUE_REAL)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_REALDeleteCallback.OnAfterDelete(stage, a_attribute_value_real, frontA_ATTRIBUTE_VALUE_REAL)
	}
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGCreateCallback.OnAfterCreate(stage, a_attribute_value_string)
	}
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateCallback != nil {
		var frontA_ATTRIBUTE_VALUE_STRING *A_ATTRIBUTE_VALUE_STRING
		if front != nil {
			frontA_ATTRIBUTE_VALUE_STRING, _ = front.(*A_ATTRIBUTE_VALUE_STRING)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateCallback.OnAfterUpdate(stage, a_attribute_value_string, frontA_ATTRIBUTE_VALUE_STRING)
	}
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGDeleteCallback != nil {
		var frontA_ATTRIBUTE_VALUE_STRING *A_ATTRIBUTE_VALUE_STRING
		if front != nil {
			frontA_ATTRIBUTE_VALUE_STRING, _ = front.(*A_ATTRIBUTE_VALUE_STRING)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGDeleteCallback.OnAfterDelete(stage, a_attribute_value_string, frontA_ATTRIBUTE_VALUE_STRING)
	}
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLCreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLCreateCallback.OnAfterCreate(stage, a_attribute_value_xhtml)
	}
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateCallback != nil {
		var frontA_ATTRIBUTE_VALUE_XHTML *A_ATTRIBUTE_VALUE_XHTML
		if front != nil {
			frontA_ATTRIBUTE_VALUE_XHTML, _ = front.(*A_ATTRIBUTE_VALUE_XHTML)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateCallback.OnAfterUpdate(stage, a_attribute_value_xhtml, frontA_ATTRIBUTE_VALUE_XHTML)
	}
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLDeleteCallback != nil {
		var frontA_ATTRIBUTE_VALUE_XHTML *A_ATTRIBUTE_VALUE_XHTML
		if front != nil {
			frontA_ATTRIBUTE_VALUE_XHTML, _ = front.(*A_ATTRIBUTE_VALUE_XHTML)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLDeleteCallback.OnAfterDelete(stage, a_attribute_value_xhtml, frontA_ATTRIBUTE_VALUE_XHTML)
	}
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1CreateCallback != nil {
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1CreateCallback.OnAfterCreate(stage, a_attribute_value_xhtml_1)
	}
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateCallback != nil {
		var frontA_ATTRIBUTE_VALUE_XHTML_1 *A_ATTRIBUTE_VALUE_XHTML_1
		if front != nil {
			frontA_ATTRIBUTE_VALUE_XHTML_1, _ = front.(*A_ATTRIBUTE_VALUE_XHTML_1)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateCallback.OnAfterUpdate(stage, a_attribute_value_xhtml_1, frontA_ATTRIBUTE_VALUE_XHTML_1)
	}
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1DeleteCallback != nil {
		var frontA_ATTRIBUTE_VALUE_XHTML_1 *A_ATTRIBUTE_VALUE_XHTML_1
		if front != nil {
			frontA_ATTRIBUTE_VALUE_XHTML_1, _ = front.(*A_ATTRIBUTE_VALUE_XHTML_1)
		}
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1DeleteCallback.OnAfterDelete(stage, a_attribute_value_xhtml_1, frontA_ATTRIBUTE_VALUE_XHTML_1)
	}
}

func (a_children *A_CHILDREN) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_CHILDRENCreateCallback != nil {
		stage.OnAfterA_CHILDRENCreateCallback.OnAfterCreate(stage, a_children)
	}
}

func (a_children *A_CHILDREN) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_CHILDRENUpdateCallback != nil {
		var frontA_CHILDREN *A_CHILDREN
		if front != nil {
			frontA_CHILDREN, _ = front.(*A_CHILDREN)
		}
		stage.OnAfterA_CHILDRENUpdateCallback.OnAfterUpdate(stage, a_children, frontA_CHILDREN)
	}
}

func (a_children *A_CHILDREN) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_CHILDRENDeleteCallback != nil {
		var frontA_CHILDREN *A_CHILDREN
		if front != nil {
			frontA_CHILDREN, _ = front.(*A_CHILDREN)
		}
		stage.OnAfterA_CHILDRENDeleteCallback.OnAfterDelete(stage, a_children, frontA_CHILDREN)
	}
}

func (a_core_content *A_CORE_CONTENT) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_CORE_CONTENTCreateCallback != nil {
		stage.OnAfterA_CORE_CONTENTCreateCallback.OnAfterCreate(stage, a_core_content)
	}
}

func (a_core_content *A_CORE_CONTENT) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_CORE_CONTENTUpdateCallback != nil {
		var frontA_CORE_CONTENT *A_CORE_CONTENT
		if front != nil {
			frontA_CORE_CONTENT, _ = front.(*A_CORE_CONTENT)
		}
		stage.OnAfterA_CORE_CONTENTUpdateCallback.OnAfterUpdate(stage, a_core_content, frontA_CORE_CONTENT)
	}
}

func (a_core_content *A_CORE_CONTENT) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_CORE_CONTENTDeleteCallback != nil {
		var frontA_CORE_CONTENT *A_CORE_CONTENT
		if front != nil {
			frontA_CORE_CONTENT, _ = front.(*A_CORE_CONTENT)
		}
		stage.OnAfterA_CORE_CONTENTDeleteCallback.OnAfterDelete(stage, a_core_content, frontA_CORE_CONTENT)
	}
}

func (a_datatypes *A_DATATYPES) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_DATATYPESCreateCallback != nil {
		stage.OnAfterA_DATATYPESCreateCallback.OnAfterCreate(stage, a_datatypes)
	}
}

func (a_datatypes *A_DATATYPES) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPESUpdateCallback != nil {
		var frontA_DATATYPES *A_DATATYPES
		if front != nil {
			frontA_DATATYPES, _ = front.(*A_DATATYPES)
		}
		stage.OnAfterA_DATATYPESUpdateCallback.OnAfterUpdate(stage, a_datatypes, frontA_DATATYPES)
	}
}

func (a_datatypes *A_DATATYPES) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPESDeleteCallback != nil {
		var frontA_DATATYPES *A_DATATYPES
		if front != nil {
			frontA_DATATYPES, _ = front.(*A_DATATYPES)
		}
		stage.OnAfterA_DATATYPESDeleteCallback.OnAfterDelete(stage, a_datatypes, frontA_DATATYPES)
	}
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFCreateCallback != nil {
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFCreateCallback.OnAfterCreate(stage, a_datatype_definition_boolean_ref)
	}
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateCallback != nil {
		var frontA_DATATYPE_DEFINITION_BOOLEAN_REF *A_DATATYPE_DEFINITION_BOOLEAN_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_BOOLEAN_REF, _ = front.(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateCallback.OnAfterUpdate(stage, a_datatype_definition_boolean_ref, frontA_DATATYPE_DEFINITION_BOOLEAN_REF)
	}
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFDeleteCallback != nil {
		var frontA_DATATYPE_DEFINITION_BOOLEAN_REF *A_DATATYPE_DEFINITION_BOOLEAN_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_BOOLEAN_REF, _ = front.(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFDeleteCallback.OnAfterDelete(stage, a_datatype_definition_boolean_ref, frontA_DATATYPE_DEFINITION_BOOLEAN_REF)
	}
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFCreateCallback != nil {
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFCreateCallback.OnAfterCreate(stage, a_datatype_definition_date_ref)
	}
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateCallback != nil {
		var frontA_DATATYPE_DEFINITION_DATE_REF *A_DATATYPE_DEFINITION_DATE_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_DATE_REF, _ = front.(*A_DATATYPE_DEFINITION_DATE_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateCallback.OnAfterUpdate(stage, a_datatype_definition_date_ref, frontA_DATATYPE_DEFINITION_DATE_REF)
	}
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFDeleteCallback != nil {
		var frontA_DATATYPE_DEFINITION_DATE_REF *A_DATATYPE_DEFINITION_DATE_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_DATE_REF, _ = front.(*A_DATATYPE_DEFINITION_DATE_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFDeleteCallback.OnAfterDelete(stage, a_datatype_definition_date_ref, frontA_DATATYPE_DEFINITION_DATE_REF)
	}
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFCreateCallback != nil {
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFCreateCallback.OnAfterCreate(stage, a_datatype_definition_enumeration_ref)
	}
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateCallback != nil {
		var frontA_DATATYPE_DEFINITION_ENUMERATION_REF *A_DATATYPE_DEFINITION_ENUMERATION_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_ENUMERATION_REF, _ = front.(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateCallback.OnAfterUpdate(stage, a_datatype_definition_enumeration_ref, frontA_DATATYPE_DEFINITION_ENUMERATION_REF)
	}
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFDeleteCallback != nil {
		var frontA_DATATYPE_DEFINITION_ENUMERATION_REF *A_DATATYPE_DEFINITION_ENUMERATION_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_ENUMERATION_REF, _ = front.(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFDeleteCallback.OnAfterDelete(stage, a_datatype_definition_enumeration_ref, frontA_DATATYPE_DEFINITION_ENUMERATION_REF)
	}
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFCreateCallback != nil {
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFCreateCallback.OnAfterCreate(stage, a_datatype_definition_integer_ref)
	}
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateCallback != nil {
		var frontA_DATATYPE_DEFINITION_INTEGER_REF *A_DATATYPE_DEFINITION_INTEGER_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_INTEGER_REF, _ = front.(*A_DATATYPE_DEFINITION_INTEGER_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateCallback.OnAfterUpdate(stage, a_datatype_definition_integer_ref, frontA_DATATYPE_DEFINITION_INTEGER_REF)
	}
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFDeleteCallback != nil {
		var frontA_DATATYPE_DEFINITION_INTEGER_REF *A_DATATYPE_DEFINITION_INTEGER_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_INTEGER_REF, _ = front.(*A_DATATYPE_DEFINITION_INTEGER_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFDeleteCallback.OnAfterDelete(stage, a_datatype_definition_integer_ref, frontA_DATATYPE_DEFINITION_INTEGER_REF)
	}
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFCreateCallback != nil {
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFCreateCallback.OnAfterCreate(stage, a_datatype_definition_real_ref)
	}
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateCallback != nil {
		var frontA_DATATYPE_DEFINITION_REAL_REF *A_DATATYPE_DEFINITION_REAL_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_REAL_REF, _ = front.(*A_DATATYPE_DEFINITION_REAL_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateCallback.OnAfterUpdate(stage, a_datatype_definition_real_ref, frontA_DATATYPE_DEFINITION_REAL_REF)
	}
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFDeleteCallback != nil {
		var frontA_DATATYPE_DEFINITION_REAL_REF *A_DATATYPE_DEFINITION_REAL_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_REAL_REF, _ = front.(*A_DATATYPE_DEFINITION_REAL_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFDeleteCallback.OnAfterDelete(stage, a_datatype_definition_real_ref, frontA_DATATYPE_DEFINITION_REAL_REF)
	}
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFCreateCallback != nil {
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFCreateCallback.OnAfterCreate(stage, a_datatype_definition_string_ref)
	}
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateCallback != nil {
		var frontA_DATATYPE_DEFINITION_STRING_REF *A_DATATYPE_DEFINITION_STRING_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_STRING_REF, _ = front.(*A_DATATYPE_DEFINITION_STRING_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateCallback.OnAfterUpdate(stage, a_datatype_definition_string_ref, frontA_DATATYPE_DEFINITION_STRING_REF)
	}
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFDeleteCallback != nil {
		var frontA_DATATYPE_DEFINITION_STRING_REF *A_DATATYPE_DEFINITION_STRING_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_STRING_REF, _ = front.(*A_DATATYPE_DEFINITION_STRING_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFDeleteCallback.OnAfterDelete(stage, a_datatype_definition_string_ref, frontA_DATATYPE_DEFINITION_STRING_REF)
	}
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFCreateCallback != nil {
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFCreateCallback.OnAfterCreate(stage, a_datatype_definition_xhtml_ref)
	}
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateCallback != nil {
		var frontA_DATATYPE_DEFINITION_XHTML_REF *A_DATATYPE_DEFINITION_XHTML_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_XHTML_REF, _ = front.(*A_DATATYPE_DEFINITION_XHTML_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateCallback.OnAfterUpdate(stage, a_datatype_definition_xhtml_ref, frontA_DATATYPE_DEFINITION_XHTML_REF)
	}
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFDeleteCallback != nil {
		var frontA_DATATYPE_DEFINITION_XHTML_REF *A_DATATYPE_DEFINITION_XHTML_REF
		if front != nil {
			frontA_DATATYPE_DEFINITION_XHTML_REF, _ = front.(*A_DATATYPE_DEFINITION_XHTML_REF)
		}
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFDeleteCallback.OnAfterDelete(stage, a_datatype_definition_xhtml_ref, frontA_DATATYPE_DEFINITION_XHTML_REF)
	}
}

func (a_editable_atts *A_EDITABLE_ATTS) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_EDITABLE_ATTSCreateCallback != nil {
		stage.OnAfterA_EDITABLE_ATTSCreateCallback.OnAfterCreate(stage, a_editable_atts)
	}
}

func (a_editable_atts *A_EDITABLE_ATTS) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_EDITABLE_ATTSUpdateCallback != nil {
		var frontA_EDITABLE_ATTS *A_EDITABLE_ATTS
		if front != nil {
			frontA_EDITABLE_ATTS, _ = front.(*A_EDITABLE_ATTS)
		}
		stage.OnAfterA_EDITABLE_ATTSUpdateCallback.OnAfterUpdate(stage, a_editable_atts, frontA_EDITABLE_ATTS)
	}
}

func (a_editable_atts *A_EDITABLE_ATTS) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_EDITABLE_ATTSDeleteCallback != nil {
		var frontA_EDITABLE_ATTS *A_EDITABLE_ATTS
		if front != nil {
			frontA_EDITABLE_ATTS, _ = front.(*A_EDITABLE_ATTS)
		}
		stage.OnAfterA_EDITABLE_ATTSDeleteCallback.OnAfterDelete(stage, a_editable_atts, frontA_EDITABLE_ATTS)
	}
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_ENUM_VALUE_REFCreateCallback != nil {
		stage.OnAfterA_ENUM_VALUE_REFCreateCallback.OnAfterCreate(stage, a_enum_value_ref)
	}
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ENUM_VALUE_REFUpdateCallback != nil {
		var frontA_ENUM_VALUE_REF *A_ENUM_VALUE_REF
		if front != nil {
			frontA_ENUM_VALUE_REF, _ = front.(*A_ENUM_VALUE_REF)
		}
		stage.OnAfterA_ENUM_VALUE_REFUpdateCallback.OnAfterUpdate(stage, a_enum_value_ref, frontA_ENUM_VALUE_REF)
	}
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_ENUM_VALUE_REFDeleteCallback != nil {
		var frontA_ENUM_VALUE_REF *A_ENUM_VALUE_REF
		if front != nil {
			frontA_ENUM_VALUE_REF, _ = front.(*A_ENUM_VALUE_REF)
		}
		stage.OnAfterA_ENUM_VALUE_REFDeleteCallback.OnAfterDelete(stage, a_enum_value_ref, frontA_ENUM_VALUE_REF)
	}
}

func (a_object *A_OBJECT) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_OBJECTCreateCallback != nil {
		stage.OnAfterA_OBJECTCreateCallback.OnAfterCreate(stage, a_object)
	}
}

func (a_object *A_OBJECT) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_OBJECTUpdateCallback != nil {
		var frontA_OBJECT *A_OBJECT
		if front != nil {
			frontA_OBJECT, _ = front.(*A_OBJECT)
		}
		stage.OnAfterA_OBJECTUpdateCallback.OnAfterUpdate(stage, a_object, frontA_OBJECT)
	}
}

func (a_object *A_OBJECT) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_OBJECTDeleteCallback != nil {
		var frontA_OBJECT *A_OBJECT
		if front != nil {
			frontA_OBJECT, _ = front.(*A_OBJECT)
		}
		stage.OnAfterA_OBJECTDeleteCallback.OnAfterDelete(stage, a_object, frontA_OBJECT)
	}
}

func (a_properties *A_PROPERTIES) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_PROPERTIESCreateCallback != nil {
		stage.OnAfterA_PROPERTIESCreateCallback.OnAfterCreate(stage, a_properties)
	}
}

func (a_properties *A_PROPERTIES) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_PROPERTIESUpdateCallback != nil {
		var frontA_PROPERTIES *A_PROPERTIES
		if front != nil {
			frontA_PROPERTIES, _ = front.(*A_PROPERTIES)
		}
		stage.OnAfterA_PROPERTIESUpdateCallback.OnAfterUpdate(stage, a_properties, frontA_PROPERTIES)
	}
}

func (a_properties *A_PROPERTIES) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_PROPERTIESDeleteCallback != nil {
		var frontA_PROPERTIES *A_PROPERTIES
		if front != nil {
			frontA_PROPERTIES, _ = front.(*A_PROPERTIES)
		}
		stage.OnAfterA_PROPERTIESDeleteCallback.OnAfterDelete(stage, a_properties, frontA_PROPERTIES)
	}
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_RELATION_GROUP_TYPE_REFCreateCallback != nil {
		stage.OnAfterA_RELATION_GROUP_TYPE_REFCreateCallback.OnAfterCreate(stage, a_relation_group_type_ref)
	}
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateCallback != nil {
		var frontA_RELATION_GROUP_TYPE_REF *A_RELATION_GROUP_TYPE_REF
		if front != nil {
			frontA_RELATION_GROUP_TYPE_REF, _ = front.(*A_RELATION_GROUP_TYPE_REF)
		}
		stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateCallback.OnAfterUpdate(stage, a_relation_group_type_ref, frontA_RELATION_GROUP_TYPE_REF)
	}
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_RELATION_GROUP_TYPE_REFDeleteCallback != nil {
		var frontA_RELATION_GROUP_TYPE_REF *A_RELATION_GROUP_TYPE_REF
		if front != nil {
			frontA_RELATION_GROUP_TYPE_REF, _ = front.(*A_RELATION_GROUP_TYPE_REF)
		}
		stage.OnAfterA_RELATION_GROUP_TYPE_REFDeleteCallback.OnAfterDelete(stage, a_relation_group_type_ref, frontA_RELATION_GROUP_TYPE_REF)
	}
}

func (a_source_1 *A_SOURCE_1) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SOURCE_1CreateCallback != nil {
		stage.OnAfterA_SOURCE_1CreateCallback.OnAfterCreate(stage, a_source_1)
	}
}

func (a_source_1 *A_SOURCE_1) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SOURCE_1UpdateCallback != nil {
		var frontA_SOURCE_1 *A_SOURCE_1
		if front != nil {
			frontA_SOURCE_1, _ = front.(*A_SOURCE_1)
		}
		stage.OnAfterA_SOURCE_1UpdateCallback.OnAfterUpdate(stage, a_source_1, frontA_SOURCE_1)
	}
}

func (a_source_1 *A_SOURCE_1) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SOURCE_1DeleteCallback != nil {
		var frontA_SOURCE_1 *A_SOURCE_1
		if front != nil {
			frontA_SOURCE_1, _ = front.(*A_SOURCE_1)
		}
		stage.OnAfterA_SOURCE_1DeleteCallback.OnAfterDelete(stage, a_source_1, frontA_SOURCE_1)
	}
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SOURCE_SPECIFICATION_1CreateCallback != nil {
		stage.OnAfterA_SOURCE_SPECIFICATION_1CreateCallback.OnAfterCreate(stage, a_source_specification_1)
	}
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateCallback != nil {
		var frontA_SOURCE_SPECIFICATION_1 *A_SOURCE_SPECIFICATION_1
		if front != nil {
			frontA_SOURCE_SPECIFICATION_1, _ = front.(*A_SOURCE_SPECIFICATION_1)
		}
		stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateCallback.OnAfterUpdate(stage, a_source_specification_1, frontA_SOURCE_SPECIFICATION_1)
	}
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SOURCE_SPECIFICATION_1DeleteCallback != nil {
		var frontA_SOURCE_SPECIFICATION_1 *A_SOURCE_SPECIFICATION_1
		if front != nil {
			frontA_SOURCE_SPECIFICATION_1, _ = front.(*A_SOURCE_SPECIFICATION_1)
		}
		stage.OnAfterA_SOURCE_SPECIFICATION_1DeleteCallback.OnAfterDelete(stage, a_source_specification_1, frontA_SOURCE_SPECIFICATION_1)
	}
}

func (a_specifications *A_SPECIFICATIONS) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPECIFICATIONSCreateCallback != nil {
		stage.OnAfterA_SPECIFICATIONSCreateCallback.OnAfterCreate(stage, a_specifications)
	}
}

func (a_specifications *A_SPECIFICATIONS) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPECIFICATIONSUpdateCallback != nil {
		var frontA_SPECIFICATIONS *A_SPECIFICATIONS
		if front != nil {
			frontA_SPECIFICATIONS, _ = front.(*A_SPECIFICATIONS)
		}
		stage.OnAfterA_SPECIFICATIONSUpdateCallback.OnAfterUpdate(stage, a_specifications, frontA_SPECIFICATIONS)
	}
}

func (a_specifications *A_SPECIFICATIONS) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPECIFICATIONSDeleteCallback != nil {
		var frontA_SPECIFICATIONS *A_SPECIFICATIONS
		if front != nil {
			frontA_SPECIFICATIONS, _ = front.(*A_SPECIFICATIONS)
		}
		stage.OnAfterA_SPECIFICATIONSDeleteCallback.OnAfterDelete(stage, a_specifications, frontA_SPECIFICATIONS)
	}
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPECIFICATION_TYPE_REFCreateCallback != nil {
		stage.OnAfterA_SPECIFICATION_TYPE_REFCreateCallback.OnAfterCreate(stage, a_specification_type_ref)
	}
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateCallback != nil {
		var frontA_SPECIFICATION_TYPE_REF *A_SPECIFICATION_TYPE_REF
		if front != nil {
			frontA_SPECIFICATION_TYPE_REF, _ = front.(*A_SPECIFICATION_TYPE_REF)
		}
		stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateCallback.OnAfterUpdate(stage, a_specification_type_ref, frontA_SPECIFICATION_TYPE_REF)
	}
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPECIFICATION_TYPE_REFDeleteCallback != nil {
		var frontA_SPECIFICATION_TYPE_REF *A_SPECIFICATION_TYPE_REF
		if front != nil {
			frontA_SPECIFICATION_TYPE_REF, _ = front.(*A_SPECIFICATION_TYPE_REF)
		}
		stage.OnAfterA_SPECIFICATION_TYPE_REFDeleteCallback.OnAfterDelete(stage, a_specification_type_ref, frontA_SPECIFICATION_TYPE_REF)
	}
}

func (a_specified_values *A_SPECIFIED_VALUES) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPECIFIED_VALUESCreateCallback != nil {
		stage.OnAfterA_SPECIFIED_VALUESCreateCallback.OnAfterCreate(stage, a_specified_values)
	}
}

func (a_specified_values *A_SPECIFIED_VALUES) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPECIFIED_VALUESUpdateCallback != nil {
		var frontA_SPECIFIED_VALUES *A_SPECIFIED_VALUES
		if front != nil {
			frontA_SPECIFIED_VALUES, _ = front.(*A_SPECIFIED_VALUES)
		}
		stage.OnAfterA_SPECIFIED_VALUESUpdateCallback.OnAfterUpdate(stage, a_specified_values, frontA_SPECIFIED_VALUES)
	}
}

func (a_specified_values *A_SPECIFIED_VALUES) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPECIFIED_VALUESDeleteCallback != nil {
		var frontA_SPECIFIED_VALUES *A_SPECIFIED_VALUES
		if front != nil {
			frontA_SPECIFIED_VALUES, _ = front.(*A_SPECIFIED_VALUES)
		}
		stage.OnAfterA_SPECIFIED_VALUESDeleteCallback.OnAfterDelete(stage, a_specified_values, frontA_SPECIFIED_VALUES)
	}
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPEC_ATTRIBUTESCreateCallback != nil {
		stage.OnAfterA_SPEC_ATTRIBUTESCreateCallback.OnAfterCreate(stage, a_spec_attributes)
	}
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_ATTRIBUTESUpdateCallback != nil {
		var frontA_SPEC_ATTRIBUTES *A_SPEC_ATTRIBUTES
		if front != nil {
			frontA_SPEC_ATTRIBUTES, _ = front.(*A_SPEC_ATTRIBUTES)
		}
		stage.OnAfterA_SPEC_ATTRIBUTESUpdateCallback.OnAfterUpdate(stage, a_spec_attributes, frontA_SPEC_ATTRIBUTES)
	}
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_ATTRIBUTESDeleteCallback != nil {
		var frontA_SPEC_ATTRIBUTES *A_SPEC_ATTRIBUTES
		if front != nil {
			frontA_SPEC_ATTRIBUTES, _ = front.(*A_SPEC_ATTRIBUTES)
		}
		stage.OnAfterA_SPEC_ATTRIBUTESDeleteCallback.OnAfterDelete(stage, a_spec_attributes, frontA_SPEC_ATTRIBUTES)
	}
}

func (a_spec_objects *A_SPEC_OBJECTS) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPEC_OBJECTSCreateCallback != nil {
		stage.OnAfterA_SPEC_OBJECTSCreateCallback.OnAfterCreate(stage, a_spec_objects)
	}
}

func (a_spec_objects *A_SPEC_OBJECTS) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_OBJECTSUpdateCallback != nil {
		var frontA_SPEC_OBJECTS *A_SPEC_OBJECTS
		if front != nil {
			frontA_SPEC_OBJECTS, _ = front.(*A_SPEC_OBJECTS)
		}
		stage.OnAfterA_SPEC_OBJECTSUpdateCallback.OnAfterUpdate(stage, a_spec_objects, frontA_SPEC_OBJECTS)
	}
}

func (a_spec_objects *A_SPEC_OBJECTS) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_OBJECTSDeleteCallback != nil {
		var frontA_SPEC_OBJECTS *A_SPEC_OBJECTS
		if front != nil {
			frontA_SPEC_OBJECTS, _ = front.(*A_SPEC_OBJECTS)
		}
		stage.OnAfterA_SPEC_OBJECTSDeleteCallback.OnAfterDelete(stage, a_spec_objects, frontA_SPEC_OBJECTS)
	}
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPEC_OBJECT_TYPE_REFCreateCallback != nil {
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFCreateCallback.OnAfterCreate(stage, a_spec_object_type_ref)
	}
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateCallback != nil {
		var frontA_SPEC_OBJECT_TYPE_REF *A_SPEC_OBJECT_TYPE_REF
		if front != nil {
			frontA_SPEC_OBJECT_TYPE_REF, _ = front.(*A_SPEC_OBJECT_TYPE_REF)
		}
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateCallback.OnAfterUpdate(stage, a_spec_object_type_ref, frontA_SPEC_OBJECT_TYPE_REF)
	}
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_OBJECT_TYPE_REFDeleteCallback != nil {
		var frontA_SPEC_OBJECT_TYPE_REF *A_SPEC_OBJECT_TYPE_REF
		if front != nil {
			frontA_SPEC_OBJECT_TYPE_REF, _ = front.(*A_SPEC_OBJECT_TYPE_REF)
		}
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFDeleteCallback.OnAfterDelete(stage, a_spec_object_type_ref, frontA_SPEC_OBJECT_TYPE_REF)
	}
}

func (a_spec_relations *A_SPEC_RELATIONS) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPEC_RELATIONSCreateCallback != nil {
		stage.OnAfterA_SPEC_RELATIONSCreateCallback.OnAfterCreate(stage, a_spec_relations)
	}
}

func (a_spec_relations *A_SPEC_RELATIONS) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_RELATIONSUpdateCallback != nil {
		var frontA_SPEC_RELATIONS *A_SPEC_RELATIONS
		if front != nil {
			frontA_SPEC_RELATIONS, _ = front.(*A_SPEC_RELATIONS)
		}
		stage.OnAfterA_SPEC_RELATIONSUpdateCallback.OnAfterUpdate(stage, a_spec_relations, frontA_SPEC_RELATIONS)
	}
}

func (a_spec_relations *A_SPEC_RELATIONS) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_RELATIONSDeleteCallback != nil {
		var frontA_SPEC_RELATIONS *A_SPEC_RELATIONS
		if front != nil {
			frontA_SPEC_RELATIONS, _ = front.(*A_SPEC_RELATIONS)
		}
		stage.OnAfterA_SPEC_RELATIONSDeleteCallback.OnAfterDelete(stage, a_spec_relations, frontA_SPEC_RELATIONS)
	}
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPEC_RELATION_GROUPSCreateCallback != nil {
		stage.OnAfterA_SPEC_RELATION_GROUPSCreateCallback.OnAfterCreate(stage, a_spec_relation_groups)
	}
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_RELATION_GROUPSUpdateCallback != nil {
		var frontA_SPEC_RELATION_GROUPS *A_SPEC_RELATION_GROUPS
		if front != nil {
			frontA_SPEC_RELATION_GROUPS, _ = front.(*A_SPEC_RELATION_GROUPS)
		}
		stage.OnAfterA_SPEC_RELATION_GROUPSUpdateCallback.OnAfterUpdate(stage, a_spec_relation_groups, frontA_SPEC_RELATION_GROUPS)
	}
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_RELATION_GROUPSDeleteCallback != nil {
		var frontA_SPEC_RELATION_GROUPS *A_SPEC_RELATION_GROUPS
		if front != nil {
			frontA_SPEC_RELATION_GROUPS, _ = front.(*A_SPEC_RELATION_GROUPS)
		}
		stage.OnAfterA_SPEC_RELATION_GROUPSDeleteCallback.OnAfterDelete(stage, a_spec_relation_groups, frontA_SPEC_RELATION_GROUPS)
	}
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPEC_RELATION_REFCreateCallback != nil {
		stage.OnAfterA_SPEC_RELATION_REFCreateCallback.OnAfterCreate(stage, a_spec_relation_ref)
	}
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_RELATION_REFUpdateCallback != nil {
		var frontA_SPEC_RELATION_REF *A_SPEC_RELATION_REF
		if front != nil {
			frontA_SPEC_RELATION_REF, _ = front.(*A_SPEC_RELATION_REF)
		}
		stage.OnAfterA_SPEC_RELATION_REFUpdateCallback.OnAfterUpdate(stage, a_spec_relation_ref, frontA_SPEC_RELATION_REF)
	}
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_RELATION_REFDeleteCallback != nil {
		var frontA_SPEC_RELATION_REF *A_SPEC_RELATION_REF
		if front != nil {
			frontA_SPEC_RELATION_REF, _ = front.(*A_SPEC_RELATION_REF)
		}
		stage.OnAfterA_SPEC_RELATION_REFDeleteCallback.OnAfterDelete(stage, a_spec_relation_ref, frontA_SPEC_RELATION_REF)
	}
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPEC_RELATION_TYPE_REFCreateCallback != nil {
		stage.OnAfterA_SPEC_RELATION_TYPE_REFCreateCallback.OnAfterCreate(stage, a_spec_relation_type_ref)
	}
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateCallback != nil {
		var frontA_SPEC_RELATION_TYPE_REF *A_SPEC_RELATION_TYPE_REF
		if front != nil {
			frontA_SPEC_RELATION_TYPE_REF, _ = front.(*A_SPEC_RELATION_TYPE_REF)
		}
		stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateCallback.OnAfterUpdate(stage, a_spec_relation_type_ref, frontA_SPEC_RELATION_TYPE_REF)
	}
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_RELATION_TYPE_REFDeleteCallback != nil {
		var frontA_SPEC_RELATION_TYPE_REF *A_SPEC_RELATION_TYPE_REF
		if front != nil {
			frontA_SPEC_RELATION_TYPE_REF, _ = front.(*A_SPEC_RELATION_TYPE_REF)
		}
		stage.OnAfterA_SPEC_RELATION_TYPE_REFDeleteCallback.OnAfterDelete(stage, a_spec_relation_type_ref, frontA_SPEC_RELATION_TYPE_REF)
	}
}

func (a_spec_types *A_SPEC_TYPES) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_SPEC_TYPESCreateCallback != nil {
		stage.OnAfterA_SPEC_TYPESCreateCallback.OnAfterCreate(stage, a_spec_types)
	}
}

func (a_spec_types *A_SPEC_TYPES) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_TYPESUpdateCallback != nil {
		var frontA_SPEC_TYPES *A_SPEC_TYPES
		if front != nil {
			frontA_SPEC_TYPES, _ = front.(*A_SPEC_TYPES)
		}
		stage.OnAfterA_SPEC_TYPESUpdateCallback.OnAfterUpdate(stage, a_spec_types, frontA_SPEC_TYPES)
	}
}

func (a_spec_types *A_SPEC_TYPES) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_SPEC_TYPESDeleteCallback != nil {
		var frontA_SPEC_TYPES *A_SPEC_TYPES
		if front != nil {
			frontA_SPEC_TYPES, _ = front.(*A_SPEC_TYPES)
		}
		stage.OnAfterA_SPEC_TYPESDeleteCallback.OnAfterDelete(stage, a_spec_types, frontA_SPEC_TYPES)
	}
}

func (a_the_header *A_THE_HEADER) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_THE_HEADERCreateCallback != nil {
		stage.OnAfterA_THE_HEADERCreateCallback.OnAfterCreate(stage, a_the_header)
	}
}

func (a_the_header *A_THE_HEADER) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_THE_HEADERUpdateCallback != nil {
		var frontA_THE_HEADER *A_THE_HEADER
		if front != nil {
			frontA_THE_HEADER, _ = front.(*A_THE_HEADER)
		}
		stage.OnAfterA_THE_HEADERUpdateCallback.OnAfterUpdate(stage, a_the_header, frontA_THE_HEADER)
	}
}

func (a_the_header *A_THE_HEADER) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_THE_HEADERDeleteCallback != nil {
		var frontA_THE_HEADER *A_THE_HEADER
		if front != nil {
			frontA_THE_HEADER, _ = front.(*A_THE_HEADER)
		}
		stage.OnAfterA_THE_HEADERDeleteCallback.OnAfterDelete(stage, a_the_header, frontA_THE_HEADER)
	}
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_TOOL_EXTENSIONSCreateCallback != nil {
		stage.OnAfterA_TOOL_EXTENSIONSCreateCallback.OnAfterCreate(stage, a_tool_extensions)
	}
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_TOOL_EXTENSIONSUpdateCallback != nil {
		var frontA_TOOL_EXTENSIONS *A_TOOL_EXTENSIONS
		if front != nil {
			frontA_TOOL_EXTENSIONS, _ = front.(*A_TOOL_EXTENSIONS)
		}
		stage.OnAfterA_TOOL_EXTENSIONSUpdateCallback.OnAfterUpdate(stage, a_tool_extensions, frontA_TOOL_EXTENSIONS)
	}
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_TOOL_EXTENSIONSDeleteCallback != nil {
		var frontA_TOOL_EXTENSIONS *A_TOOL_EXTENSIONS
		if front != nil {
			frontA_TOOL_EXTENSIONS, _ = front.(*A_TOOL_EXTENSIONS)
		}
		stage.OnAfterA_TOOL_EXTENSIONSDeleteCallback.OnAfterDelete(stage, a_tool_extensions, frontA_TOOL_EXTENSIONS)
	}
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback != nil {
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback.OnAfterCreate(stage, datatype_definition_boolean)
	}
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback != nil {
		var frontDATATYPE_DEFINITION_BOOLEAN *DATATYPE_DEFINITION_BOOLEAN
		if front != nil {
			frontDATATYPE_DEFINITION_BOOLEAN, _ = front.(*DATATYPE_DEFINITION_BOOLEAN)
		}
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback.OnAfterUpdate(stage, datatype_definition_boolean, frontDATATYPE_DEFINITION_BOOLEAN)
	}
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback != nil {
		var frontDATATYPE_DEFINITION_BOOLEAN *DATATYPE_DEFINITION_BOOLEAN
		if front != nil {
			frontDATATYPE_DEFINITION_BOOLEAN, _ = front.(*DATATYPE_DEFINITION_BOOLEAN)
		}
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback.OnAfterDelete(stage, datatype_definition_boolean, frontDATATYPE_DEFINITION_BOOLEAN)
	}
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDATATYPE_DEFINITION_DATECreateCallback != nil {
		stage.OnAfterDATATYPE_DEFINITION_DATECreateCallback.OnAfterCreate(stage, datatype_definition_date)
	}
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback != nil {
		var frontDATATYPE_DEFINITION_DATE *DATATYPE_DEFINITION_DATE
		if front != nil {
			frontDATATYPE_DEFINITION_DATE, _ = front.(*DATATYPE_DEFINITION_DATE)
		}
		stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback.OnAfterUpdate(stage, datatype_definition_date, frontDATATYPE_DEFINITION_DATE)
	}
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_DATEDeleteCallback != nil {
		var frontDATATYPE_DEFINITION_DATE *DATATYPE_DEFINITION_DATE
		if front != nil {
			frontDATATYPE_DEFINITION_DATE, _ = front.(*DATATYPE_DEFINITION_DATE)
		}
		stage.OnAfterDATATYPE_DEFINITION_DATEDeleteCallback.OnAfterDelete(stage, datatype_definition_date, frontDATATYPE_DEFINITION_DATE)
	}
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback != nil {
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback.OnAfterCreate(stage, datatype_definition_enumeration)
	}
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback != nil {
		var frontDATATYPE_DEFINITION_ENUMERATION *DATATYPE_DEFINITION_ENUMERATION
		if front != nil {
			frontDATATYPE_DEFINITION_ENUMERATION, _ = front.(*DATATYPE_DEFINITION_ENUMERATION)
		}
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, datatype_definition_enumeration, frontDATATYPE_DEFINITION_ENUMERATION)
	}
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback != nil {
		var frontDATATYPE_DEFINITION_ENUMERATION *DATATYPE_DEFINITION_ENUMERATION
		if front != nil {
			frontDATATYPE_DEFINITION_ENUMERATION, _ = front.(*DATATYPE_DEFINITION_ENUMERATION)
		}
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback.OnAfterDelete(stage, datatype_definition_enumeration, frontDATATYPE_DEFINITION_ENUMERATION)
	}
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback != nil {
		stage.OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback.OnAfterCreate(stage, datatype_definition_integer)
	}
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback != nil {
		var frontDATATYPE_DEFINITION_INTEGER *DATATYPE_DEFINITION_INTEGER
		if front != nil {
			frontDATATYPE_DEFINITION_INTEGER, _ = front.(*DATATYPE_DEFINITION_INTEGER)
		}
		stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback.OnAfterUpdate(stage, datatype_definition_integer, frontDATATYPE_DEFINITION_INTEGER)
	}
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback != nil {
		var frontDATATYPE_DEFINITION_INTEGER *DATATYPE_DEFINITION_INTEGER
		if front != nil {
			frontDATATYPE_DEFINITION_INTEGER, _ = front.(*DATATYPE_DEFINITION_INTEGER)
		}
		stage.OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback.OnAfterDelete(stage, datatype_definition_integer, frontDATATYPE_DEFINITION_INTEGER)
	}
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDATATYPE_DEFINITION_REALCreateCallback != nil {
		stage.OnAfterDATATYPE_DEFINITION_REALCreateCallback.OnAfterCreate(stage, datatype_definition_real)
	}
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback != nil {
		var frontDATATYPE_DEFINITION_REAL *DATATYPE_DEFINITION_REAL
		if front != nil {
			frontDATATYPE_DEFINITION_REAL, _ = front.(*DATATYPE_DEFINITION_REAL)
		}
		stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback.OnAfterUpdate(stage, datatype_definition_real, frontDATATYPE_DEFINITION_REAL)
	}
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_REALDeleteCallback != nil {
		var frontDATATYPE_DEFINITION_REAL *DATATYPE_DEFINITION_REAL
		if front != nil {
			frontDATATYPE_DEFINITION_REAL, _ = front.(*DATATYPE_DEFINITION_REAL)
		}
		stage.OnAfterDATATYPE_DEFINITION_REALDeleteCallback.OnAfterDelete(stage, datatype_definition_real, frontDATATYPE_DEFINITION_REAL)
	}
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDATATYPE_DEFINITION_STRINGCreateCallback != nil {
		stage.OnAfterDATATYPE_DEFINITION_STRINGCreateCallback.OnAfterCreate(stage, datatype_definition_string)
	}
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback != nil {
		var frontDATATYPE_DEFINITION_STRING *DATATYPE_DEFINITION_STRING
		if front != nil {
			frontDATATYPE_DEFINITION_STRING, _ = front.(*DATATYPE_DEFINITION_STRING)
		}
		stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback.OnAfterUpdate(stage, datatype_definition_string, frontDATATYPE_DEFINITION_STRING)
	}
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback != nil {
		var frontDATATYPE_DEFINITION_STRING *DATATYPE_DEFINITION_STRING
		if front != nil {
			frontDATATYPE_DEFINITION_STRING, _ = front.(*DATATYPE_DEFINITION_STRING)
		}
		stage.OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback.OnAfterDelete(stage, datatype_definition_string, frontDATATYPE_DEFINITION_STRING)
	}
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback != nil {
		stage.OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback.OnAfterCreate(stage, datatype_definition_xhtml)
	}
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback != nil {
		var frontDATATYPE_DEFINITION_XHTML *DATATYPE_DEFINITION_XHTML
		if front != nil {
			frontDATATYPE_DEFINITION_XHTML, _ = front.(*DATATYPE_DEFINITION_XHTML)
		}
		stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback.OnAfterUpdate(stage, datatype_definition_xhtml, frontDATATYPE_DEFINITION_XHTML)
	}
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback != nil {
		var frontDATATYPE_DEFINITION_XHTML *DATATYPE_DEFINITION_XHTML
		if front != nil {
			frontDATATYPE_DEFINITION_XHTML, _ = front.(*DATATYPE_DEFINITION_XHTML)
		}
		stage.OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback.OnAfterDelete(stage, datatype_definition_xhtml, frontDATATYPE_DEFINITION_XHTML)
	}
}

func (embedded_value *EMBEDDED_VALUE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEMBEDDED_VALUECreateCallback != nil {
		stage.OnAfterEMBEDDED_VALUECreateCallback.OnAfterCreate(stage, embedded_value)
	}
}

func (embedded_value *EMBEDDED_VALUE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEMBEDDED_VALUEUpdateCallback != nil {
		var frontEMBEDDED_VALUE *EMBEDDED_VALUE
		if front != nil {
			frontEMBEDDED_VALUE, _ = front.(*EMBEDDED_VALUE)
		}
		stage.OnAfterEMBEDDED_VALUEUpdateCallback.OnAfterUpdate(stage, embedded_value, frontEMBEDDED_VALUE)
	}
}

func (embedded_value *EMBEDDED_VALUE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEMBEDDED_VALUEDeleteCallback != nil {
		var frontEMBEDDED_VALUE *EMBEDDED_VALUE
		if front != nil {
			frontEMBEDDED_VALUE, _ = front.(*EMBEDDED_VALUE)
		}
		stage.OnAfterEMBEDDED_VALUEDeleteCallback.OnAfterDelete(stage, embedded_value, frontEMBEDDED_VALUE)
	}
}

func (enum_value *ENUM_VALUE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterENUM_VALUECreateCallback != nil {
		stage.OnAfterENUM_VALUECreateCallback.OnAfterCreate(stage, enum_value)
	}
}

func (enum_value *ENUM_VALUE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterENUM_VALUEUpdateCallback != nil {
		var frontENUM_VALUE *ENUM_VALUE
		if front != nil {
			frontENUM_VALUE, _ = front.(*ENUM_VALUE)
		}
		stage.OnAfterENUM_VALUEUpdateCallback.OnAfterUpdate(stage, enum_value, frontENUM_VALUE)
	}
}

func (enum_value *ENUM_VALUE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterENUM_VALUEDeleteCallback != nil {
		var frontENUM_VALUE *ENUM_VALUE
		if front != nil {
			frontENUM_VALUE, _ = front.(*ENUM_VALUE)
		}
		stage.OnAfterENUM_VALUEDeleteCallback.OnAfterDelete(stage, enum_value, frontENUM_VALUE)
	}
}

func (relation_group *RELATION_GROUP) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRELATION_GROUPCreateCallback != nil {
		stage.OnAfterRELATION_GROUPCreateCallback.OnAfterCreate(stage, relation_group)
	}
}

func (relation_group *RELATION_GROUP) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRELATION_GROUPUpdateCallback != nil {
		var frontRELATION_GROUP *RELATION_GROUP
		if front != nil {
			frontRELATION_GROUP, _ = front.(*RELATION_GROUP)
		}
		stage.OnAfterRELATION_GROUPUpdateCallback.OnAfterUpdate(stage, relation_group, frontRELATION_GROUP)
	}
}

func (relation_group *RELATION_GROUP) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRELATION_GROUPDeleteCallback != nil {
		var frontRELATION_GROUP *RELATION_GROUP
		if front != nil {
			frontRELATION_GROUP, _ = front.(*RELATION_GROUP)
		}
		stage.OnAfterRELATION_GROUPDeleteCallback.OnAfterDelete(stage, relation_group, frontRELATION_GROUP)
	}
}

func (relation_group_type *RELATION_GROUP_TYPE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRELATION_GROUP_TYPECreateCallback != nil {
		stage.OnAfterRELATION_GROUP_TYPECreateCallback.OnAfterCreate(stage, relation_group_type)
	}
}

func (relation_group_type *RELATION_GROUP_TYPE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRELATION_GROUP_TYPEUpdateCallback != nil {
		var frontRELATION_GROUP_TYPE *RELATION_GROUP_TYPE
		if front != nil {
			frontRELATION_GROUP_TYPE, _ = front.(*RELATION_GROUP_TYPE)
		}
		stage.OnAfterRELATION_GROUP_TYPEUpdateCallback.OnAfterUpdate(stage, relation_group_type, frontRELATION_GROUP_TYPE)
	}
}

func (relation_group_type *RELATION_GROUP_TYPE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRELATION_GROUP_TYPEDeleteCallback != nil {
		var frontRELATION_GROUP_TYPE *RELATION_GROUP_TYPE
		if front != nil {
			frontRELATION_GROUP_TYPE, _ = front.(*RELATION_GROUP_TYPE)
		}
		stage.OnAfterRELATION_GROUP_TYPEDeleteCallback.OnAfterDelete(stage, relation_group_type, frontRELATION_GROUP_TYPE)
	}
}

func (req_if *REQ_IF) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterREQ_IFCreateCallback != nil {
		stage.OnAfterREQ_IFCreateCallback.OnAfterCreate(stage, req_if)
	}
}

func (req_if *REQ_IF) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterREQ_IFUpdateCallback != nil {
		var frontREQ_IF *REQ_IF
		if front != nil {
			frontREQ_IF, _ = front.(*REQ_IF)
		}
		stage.OnAfterREQ_IFUpdateCallback.OnAfterUpdate(stage, req_if, frontREQ_IF)
	}
}

func (req_if *REQ_IF) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterREQ_IFDeleteCallback != nil {
		var frontREQ_IF *REQ_IF
		if front != nil {
			frontREQ_IF, _ = front.(*REQ_IF)
		}
		stage.OnAfterREQ_IFDeleteCallback.OnAfterDelete(stage, req_if, frontREQ_IF)
	}
}

func (req_if_content *REQ_IF_CONTENT) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterREQ_IF_CONTENTCreateCallback != nil {
		stage.OnAfterREQ_IF_CONTENTCreateCallback.OnAfterCreate(stage, req_if_content)
	}
}

func (req_if_content *REQ_IF_CONTENT) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterREQ_IF_CONTENTUpdateCallback != nil {
		var frontREQ_IF_CONTENT *REQ_IF_CONTENT
		if front != nil {
			frontREQ_IF_CONTENT, _ = front.(*REQ_IF_CONTENT)
		}
		stage.OnAfterREQ_IF_CONTENTUpdateCallback.OnAfterUpdate(stage, req_if_content, frontREQ_IF_CONTENT)
	}
}

func (req_if_content *REQ_IF_CONTENT) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterREQ_IF_CONTENTDeleteCallback != nil {
		var frontREQ_IF_CONTENT *REQ_IF_CONTENT
		if front != nil {
			frontREQ_IF_CONTENT, _ = front.(*REQ_IF_CONTENT)
		}
		stage.OnAfterREQ_IF_CONTENTDeleteCallback.OnAfterDelete(stage, req_if_content, frontREQ_IF_CONTENT)
	}
}

func (req_if_header *REQ_IF_HEADER) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterREQ_IF_HEADERCreateCallback != nil {
		stage.OnAfterREQ_IF_HEADERCreateCallback.OnAfterCreate(stage, req_if_header)
	}
}

func (req_if_header *REQ_IF_HEADER) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterREQ_IF_HEADERUpdateCallback != nil {
		var frontREQ_IF_HEADER *REQ_IF_HEADER
		if front != nil {
			frontREQ_IF_HEADER, _ = front.(*REQ_IF_HEADER)
		}
		stage.OnAfterREQ_IF_HEADERUpdateCallback.OnAfterUpdate(stage, req_if_header, frontREQ_IF_HEADER)
	}
}

func (req_if_header *REQ_IF_HEADER) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterREQ_IF_HEADERDeleteCallback != nil {
		var frontREQ_IF_HEADER *REQ_IF_HEADER
		if front != nil {
			frontREQ_IF_HEADER, _ = front.(*REQ_IF_HEADER)
		}
		stage.OnAfterREQ_IF_HEADERDeleteCallback.OnAfterDelete(stage, req_if_header, frontREQ_IF_HEADER)
	}
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback != nil {
		stage.OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback.OnAfterCreate(stage, req_if_tool_extension)
	}
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback != nil {
		var frontREQ_IF_TOOL_EXTENSION *REQ_IF_TOOL_EXTENSION
		if front != nil {
			frontREQ_IF_TOOL_EXTENSION, _ = front.(*REQ_IF_TOOL_EXTENSION)
		}
		stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback.OnAfterUpdate(stage, req_if_tool_extension, frontREQ_IF_TOOL_EXTENSION)
	}
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback != nil {
		var frontREQ_IF_TOOL_EXTENSION *REQ_IF_TOOL_EXTENSION
		if front != nil {
			frontREQ_IF_TOOL_EXTENSION, _ = front.(*REQ_IF_TOOL_EXTENSION)
		}
		stage.OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback.OnAfterDelete(stage, req_if_tool_extension, frontREQ_IF_TOOL_EXTENSION)
	}
}

func (specification *SPECIFICATION) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSPECIFICATIONCreateCallback != nil {
		stage.OnAfterSPECIFICATIONCreateCallback.OnAfterCreate(stage, specification)
	}
}

func (specification *SPECIFICATION) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPECIFICATIONUpdateCallback != nil {
		var frontSPECIFICATION *SPECIFICATION
		if front != nil {
			frontSPECIFICATION, _ = front.(*SPECIFICATION)
		}
		stage.OnAfterSPECIFICATIONUpdateCallback.OnAfterUpdate(stage, specification, frontSPECIFICATION)
	}
}

func (specification *SPECIFICATION) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPECIFICATIONDeleteCallback != nil {
		var frontSPECIFICATION *SPECIFICATION
		if front != nil {
			frontSPECIFICATION, _ = front.(*SPECIFICATION)
		}
		stage.OnAfterSPECIFICATIONDeleteCallback.OnAfterDelete(stage, specification, frontSPECIFICATION)
	}
}

func (specification_type *SPECIFICATION_TYPE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSPECIFICATION_TYPECreateCallback != nil {
		stage.OnAfterSPECIFICATION_TYPECreateCallback.OnAfterCreate(stage, specification_type)
	}
}

func (specification_type *SPECIFICATION_TYPE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPECIFICATION_TYPEUpdateCallback != nil {
		var frontSPECIFICATION_TYPE *SPECIFICATION_TYPE
		if front != nil {
			frontSPECIFICATION_TYPE, _ = front.(*SPECIFICATION_TYPE)
		}
		stage.OnAfterSPECIFICATION_TYPEUpdateCallback.OnAfterUpdate(stage, specification_type, frontSPECIFICATION_TYPE)
	}
}

func (specification_type *SPECIFICATION_TYPE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPECIFICATION_TYPEDeleteCallback != nil {
		var frontSPECIFICATION_TYPE *SPECIFICATION_TYPE
		if front != nil {
			frontSPECIFICATION_TYPE, _ = front.(*SPECIFICATION_TYPE)
		}
		stage.OnAfterSPECIFICATION_TYPEDeleteCallback.OnAfterDelete(stage, specification_type, frontSPECIFICATION_TYPE)
	}
}

func (spec_hierarchy *SPEC_HIERARCHY) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSPEC_HIERARCHYCreateCallback != nil {
		stage.OnAfterSPEC_HIERARCHYCreateCallback.OnAfterCreate(stage, spec_hierarchy)
	}
}

func (spec_hierarchy *SPEC_HIERARCHY) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_HIERARCHYUpdateCallback != nil {
		var frontSPEC_HIERARCHY *SPEC_HIERARCHY
		if front != nil {
			frontSPEC_HIERARCHY, _ = front.(*SPEC_HIERARCHY)
		}
		stage.OnAfterSPEC_HIERARCHYUpdateCallback.OnAfterUpdate(stage, spec_hierarchy, frontSPEC_HIERARCHY)
	}
}

func (spec_hierarchy *SPEC_HIERARCHY) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_HIERARCHYDeleteCallback != nil {
		var frontSPEC_HIERARCHY *SPEC_HIERARCHY
		if front != nil {
			frontSPEC_HIERARCHY, _ = front.(*SPEC_HIERARCHY)
		}
		stage.OnAfterSPEC_HIERARCHYDeleteCallback.OnAfterDelete(stage, spec_hierarchy, frontSPEC_HIERARCHY)
	}
}

func (spec_object *SPEC_OBJECT) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSPEC_OBJECTCreateCallback != nil {
		stage.OnAfterSPEC_OBJECTCreateCallback.OnAfterCreate(stage, spec_object)
	}
}

func (spec_object *SPEC_OBJECT) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_OBJECTUpdateCallback != nil {
		var frontSPEC_OBJECT *SPEC_OBJECT
		if front != nil {
			frontSPEC_OBJECT, _ = front.(*SPEC_OBJECT)
		}
		stage.OnAfterSPEC_OBJECTUpdateCallback.OnAfterUpdate(stage, spec_object, frontSPEC_OBJECT)
	}
}

func (spec_object *SPEC_OBJECT) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_OBJECTDeleteCallback != nil {
		var frontSPEC_OBJECT *SPEC_OBJECT
		if front != nil {
			frontSPEC_OBJECT, _ = front.(*SPEC_OBJECT)
		}
		stage.OnAfterSPEC_OBJECTDeleteCallback.OnAfterDelete(stage, spec_object, frontSPEC_OBJECT)
	}
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSPEC_OBJECT_TYPECreateCallback != nil {
		stage.OnAfterSPEC_OBJECT_TYPECreateCallback.OnAfterCreate(stage, spec_object_type)
	}
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback != nil {
		var frontSPEC_OBJECT_TYPE *SPEC_OBJECT_TYPE
		if front != nil {
			frontSPEC_OBJECT_TYPE, _ = front.(*SPEC_OBJECT_TYPE)
		}
		stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback.OnAfterUpdate(stage, spec_object_type, frontSPEC_OBJECT_TYPE)
	}
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_OBJECT_TYPEDeleteCallback != nil {
		var frontSPEC_OBJECT_TYPE *SPEC_OBJECT_TYPE
		if front != nil {
			frontSPEC_OBJECT_TYPE, _ = front.(*SPEC_OBJECT_TYPE)
		}
		stage.OnAfterSPEC_OBJECT_TYPEDeleteCallback.OnAfterDelete(stage, spec_object_type, frontSPEC_OBJECT_TYPE)
	}
}

func (spec_relation *SPEC_RELATION) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSPEC_RELATIONCreateCallback != nil {
		stage.OnAfterSPEC_RELATIONCreateCallback.OnAfterCreate(stage, spec_relation)
	}
}

func (spec_relation *SPEC_RELATION) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_RELATIONUpdateCallback != nil {
		var frontSPEC_RELATION *SPEC_RELATION
		if front != nil {
			frontSPEC_RELATION, _ = front.(*SPEC_RELATION)
		}
		stage.OnAfterSPEC_RELATIONUpdateCallback.OnAfterUpdate(stage, spec_relation, frontSPEC_RELATION)
	}
}

func (spec_relation *SPEC_RELATION) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_RELATIONDeleteCallback != nil {
		var frontSPEC_RELATION *SPEC_RELATION
		if front != nil {
			frontSPEC_RELATION, _ = front.(*SPEC_RELATION)
		}
		stage.OnAfterSPEC_RELATIONDeleteCallback.OnAfterDelete(stage, spec_relation, frontSPEC_RELATION)
	}
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSPEC_RELATION_TYPECreateCallback != nil {
		stage.OnAfterSPEC_RELATION_TYPECreateCallback.OnAfterCreate(stage, spec_relation_type)
	}
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_RELATION_TYPEUpdateCallback != nil {
		var frontSPEC_RELATION_TYPE *SPEC_RELATION_TYPE
		if front != nil {
			frontSPEC_RELATION_TYPE, _ = front.(*SPEC_RELATION_TYPE)
		}
		stage.OnAfterSPEC_RELATION_TYPEUpdateCallback.OnAfterUpdate(stage, spec_relation_type, frontSPEC_RELATION_TYPE)
	}
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSPEC_RELATION_TYPEDeleteCallback != nil {
		var frontSPEC_RELATION_TYPE *SPEC_RELATION_TYPE
		if front != nil {
			frontSPEC_RELATION_TYPE, _ = front.(*SPEC_RELATION_TYPE)
		}
		stage.OnAfterSPEC_RELATION_TYPEDeleteCallback.OnAfterDelete(stage, spec_relation_type, frontSPEC_RELATION_TYPE)
	}
}

func (xhtml_content *XHTML_CONTENT) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXHTML_CONTENTCreateCallback != nil {
		stage.OnAfterXHTML_CONTENTCreateCallback.OnAfterCreate(stage, xhtml_content)
	}
}

func (xhtml_content *XHTML_CONTENT) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXHTML_CONTENTUpdateCallback != nil {
		var frontXHTML_CONTENT *XHTML_CONTENT
		if front != nil {
			frontXHTML_CONTENT, _ = front.(*XHTML_CONTENT)
		}
		stage.OnAfterXHTML_CONTENTUpdateCallback.OnAfterUpdate(stage, xhtml_content, frontXHTML_CONTENT)
	}
}

func (xhtml_content *XHTML_CONTENT) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXHTML_CONTENTDeleteCallback != nil {
		var frontXHTML_CONTENT *XHTML_CONTENT
		if front != nil {
			frontXHTML_CONTENT, _ = front.(*XHTML_CONTENT)
		}
		stage.OnAfterXHTML_CONTENTDeleteCallback.OnAfterDelete(stage, xhtml_content, frontXHTML_CONTENT)
	}
}

