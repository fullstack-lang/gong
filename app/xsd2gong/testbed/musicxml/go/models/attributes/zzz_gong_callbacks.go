// generated code - do not edit
package attributes

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
func (attributes *Attributes) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAttributesCreateCallback != nil {
		stage.OnAfterAttributesCreateCallback.OnAfterCreate(stage, attributes)
	}
}

func (attributes *Attributes) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributesUpdateCallback != nil {
		var frontAttributes *Attributes
		if front != nil {
			frontAttributes, _ = front.(*Attributes)
		}
		stage.OnAfterAttributesUpdateCallback.OnAfterUpdate(stage, attributes, frontAttributes)
	}
}

func (attributes *Attributes) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributesDeleteCallback != nil {
		var frontAttributes *Attributes
		if front != nil {
			frontAttributes, _ = front.(*Attributes)
		}
		stage.OnAfterAttributesDeleteCallback.OnAfterDelete(stage, attributes, frontAttributes)
	}
}

func (attributesdirective *AttributesDirective) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAttributesDirectiveCreateCallback != nil {
		stage.OnAfterAttributesDirectiveCreateCallback.OnAfterCreate(stage, attributesdirective)
	}
}

func (attributesdirective *AttributesDirective) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributesDirectiveUpdateCallback != nil {
		var frontAttributesDirective *AttributesDirective
		if front != nil {
			frontAttributesDirective, _ = front.(*AttributesDirective)
		}
		stage.OnAfterAttributesDirectiveUpdateCallback.OnAfterUpdate(stage, attributesdirective, frontAttributesDirective)
	}
}

func (attributesdirective *AttributesDirective) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributesDirectiveDeleteCallback != nil {
		var frontAttributesDirective *AttributesDirective
		if front != nil {
			frontAttributesDirective, _ = front.(*AttributesDirective)
		}
		stage.OnAfterAttributesDirectiveDeleteCallback.OnAfterDelete(stage, attributesdirective, frontAttributesDirective)
	}
}

func (beatrepeat *BeatRepeat) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBeatRepeatCreateCallback != nil {
		stage.OnAfterBeatRepeatCreateCallback.OnAfterCreate(stage, beatrepeat)
	}
}

func (beatrepeat *BeatRepeat) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeatRepeatUpdateCallback != nil {
		var frontBeatRepeat *BeatRepeat
		if front != nil {
			frontBeatRepeat, _ = front.(*BeatRepeat)
		}
		stage.OnAfterBeatRepeatUpdateCallback.OnAfterUpdate(stage, beatrepeat, frontBeatRepeat)
	}
}

func (beatrepeat *BeatRepeat) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeatRepeatDeleteCallback != nil {
		var frontBeatRepeat *BeatRepeat
		if front != nil {
			frontBeatRepeat, _ = front.(*BeatRepeat)
		}
		stage.OnAfterBeatRepeatDeleteCallback.OnAfterDelete(stage, beatrepeat, frontBeatRepeat)
	}
}

func (cancel *Cancel) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCancelCreateCallback != nil {
		stage.OnAfterCancelCreateCallback.OnAfterCreate(stage, cancel)
	}
}

func (cancel *Cancel) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCancelUpdateCallback != nil {
		var frontCancel *Cancel
		if front != nil {
			frontCancel, _ = front.(*Cancel)
		}
		stage.OnAfterCancelUpdateCallback.OnAfterUpdate(stage, cancel, frontCancel)
	}
}

func (cancel *Cancel) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCancelDeleteCallback != nil {
		var frontCancel *Cancel
		if front != nil {
			frontCancel, _ = front.(*Cancel)
		}
		stage.OnAfterCancelDeleteCallback.OnAfterDelete(stage, cancel, frontCancel)
	}
}

func (clef *Clef) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterClefCreateCallback != nil {
		stage.OnAfterClefCreateCallback.OnAfterCreate(stage, clef)
	}
}

func (clef *Clef) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClefUpdateCallback != nil {
		var frontClef *Clef
		if front != nil {
			frontClef, _ = front.(*Clef)
		}
		stage.OnAfterClefUpdateCallback.OnAfterUpdate(stage, clef, frontClef)
	}
}

func (clef *Clef) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClefDeleteCallback != nil {
		var frontClef *Clef
		if front != nil {
			frontClef, _ = front.(*Clef)
		}
		stage.OnAfterClefDeleteCallback.OnAfterDelete(stage, clef, frontClef)
	}
}

func (double *Double) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDoubleCreateCallback != nil {
		stage.OnAfterDoubleCreateCallback.OnAfterCreate(stage, double)
	}
}

func (double *Double) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDoubleUpdateCallback != nil {
		var frontDouble *Double
		if front != nil {
			frontDouble, _ = front.(*Double)
		}
		stage.OnAfterDoubleUpdateCallback.OnAfterUpdate(stage, double, frontDouble)
	}
}

func (double *Double) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDoubleDeleteCallback != nil {
		var frontDouble *Double
		if front != nil {
			frontDouble, _ = front.(*Double)
		}
		stage.OnAfterDoubleDeleteCallback.OnAfterDelete(stage, double, frontDouble)
	}
}

func (forpart *ForPart) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterForPartCreateCallback != nil {
		stage.OnAfterForPartCreateCallback.OnAfterCreate(stage, forpart)
	}
}

func (forpart *ForPart) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterForPartUpdateCallback != nil {
		var frontForPart *ForPart
		if front != nil {
			frontForPart, _ = front.(*ForPart)
		}
		stage.OnAfterForPartUpdateCallback.OnAfterUpdate(stage, forpart, frontForPart)
	}
}

func (forpart *ForPart) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterForPartDeleteCallback != nil {
		var frontForPart *ForPart
		if front != nil {
			frontForPart, _ = front.(*ForPart)
		}
		stage.OnAfterForPartDeleteCallback.OnAfterDelete(stage, forpart, frontForPart)
	}
}

func (interchangeable *Interchangeable) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInterchangeableCreateCallback != nil {
		stage.OnAfterInterchangeableCreateCallback.OnAfterCreate(stage, interchangeable)
	}
}

func (interchangeable *Interchangeable) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInterchangeableUpdateCallback != nil {
		var frontInterchangeable *Interchangeable
		if front != nil {
			frontInterchangeable, _ = front.(*Interchangeable)
		}
		stage.OnAfterInterchangeableUpdateCallback.OnAfterUpdate(stage, interchangeable, frontInterchangeable)
	}
}

func (interchangeable *Interchangeable) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInterchangeableDeleteCallback != nil {
		var frontInterchangeable *Interchangeable
		if front != nil {
			frontInterchangeable, _ = front.(*Interchangeable)
		}
		stage.OnAfterInterchangeableDeleteCallback.OnAfterDelete(stage, interchangeable, frontInterchangeable)
	}
}

func (key *Key) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKeyCreateCallback != nil {
		stage.OnAfterKeyCreateCallback.OnAfterCreate(stage, key)
	}
}

func (key *Key) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyUpdateCallback != nil {
		var frontKey *Key
		if front != nil {
			frontKey, _ = front.(*Key)
		}
		stage.OnAfterKeyUpdateCallback.OnAfterUpdate(stage, key, frontKey)
	}
}

func (key *Key) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyDeleteCallback != nil {
		var frontKey *Key
		if front != nil {
			frontKey, _ = front.(*Key)
		}
		stage.OnAfterKeyDeleteCallback.OnAfterDelete(stage, key, frontKey)
	}
}

func (keyaccidental *KeyAccidental) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKeyAccidentalCreateCallback != nil {
		stage.OnAfterKeyAccidentalCreateCallback.OnAfterCreate(stage, keyaccidental)
	}
}

func (keyaccidental *KeyAccidental) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyAccidentalUpdateCallback != nil {
		var frontKeyAccidental *KeyAccidental
		if front != nil {
			frontKeyAccidental, _ = front.(*KeyAccidental)
		}
		stage.OnAfterKeyAccidentalUpdateCallback.OnAfterUpdate(stage, keyaccidental, frontKeyAccidental)
	}
}

func (keyaccidental *KeyAccidental) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyAccidentalDeleteCallback != nil {
		var frontKeyAccidental *KeyAccidental
		if front != nil {
			frontKeyAccidental, _ = front.(*KeyAccidental)
		}
		stage.OnAfterKeyAccidentalDeleteCallback.OnAfterDelete(stage, keyaccidental, frontKeyAccidental)
	}
}

func (keyoctave *KeyOctave) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKeyOctaveCreateCallback != nil {
		stage.OnAfterKeyOctaveCreateCallback.OnAfterCreate(stage, keyoctave)
	}
}

func (keyoctave *KeyOctave) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyOctaveUpdateCallback != nil {
		var frontKeyOctave *KeyOctave
		if front != nil {
			frontKeyOctave, _ = front.(*KeyOctave)
		}
		stage.OnAfterKeyOctaveUpdateCallback.OnAfterUpdate(stage, keyoctave, frontKeyOctave)
	}
}

func (keyoctave *KeyOctave) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyOctaveDeleteCallback != nil {
		var frontKeyOctave *KeyOctave
		if front != nil {
			frontKeyOctave, _ = front.(*KeyOctave)
		}
		stage.OnAfterKeyOctaveDeleteCallback.OnAfterDelete(stage, keyoctave, frontKeyOctave)
	}
}

func (linedetail *LineDetail) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLineDetailCreateCallback != nil {
		stage.OnAfterLineDetailCreateCallback.OnAfterCreate(stage, linedetail)
	}
}

func (linedetail *LineDetail) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineDetailUpdateCallback != nil {
		var frontLineDetail *LineDetail
		if front != nil {
			frontLineDetail, _ = front.(*LineDetail)
		}
		stage.OnAfterLineDetailUpdateCallback.OnAfterUpdate(stage, linedetail, frontLineDetail)
	}
}

func (linedetail *LineDetail) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineDetailDeleteCallback != nil {
		var frontLineDetail *LineDetail
		if front != nil {
			frontLineDetail, _ = front.(*LineDetail)
		}
		stage.OnAfterLineDetailDeleteCallback.OnAfterDelete(stage, linedetail, frontLineDetail)
	}
}

func (measurerepeat *MeasureRepeat) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureRepeatCreateCallback != nil {
		stage.OnAfterMeasureRepeatCreateCallback.OnAfterCreate(stage, measurerepeat)
	}
}

func (measurerepeat *MeasureRepeat) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureRepeatUpdateCallback != nil {
		var frontMeasureRepeat *MeasureRepeat
		if front != nil {
			frontMeasureRepeat, _ = front.(*MeasureRepeat)
		}
		stage.OnAfterMeasureRepeatUpdateCallback.OnAfterUpdate(stage, measurerepeat, frontMeasureRepeat)
	}
}

func (measurerepeat *MeasureRepeat) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureRepeatDeleteCallback != nil {
		var frontMeasureRepeat *MeasureRepeat
		if front != nil {
			frontMeasureRepeat, _ = front.(*MeasureRepeat)
		}
		stage.OnAfterMeasureRepeatDeleteCallback.OnAfterDelete(stage, measurerepeat, frontMeasureRepeat)
	}
}

func (measurestyle *MeasureStyle) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureStyleCreateCallback != nil {
		stage.OnAfterMeasureStyleCreateCallback.OnAfterCreate(stage, measurestyle)
	}
}

func (measurestyle *MeasureStyle) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureStyleUpdateCallback != nil {
		var frontMeasureStyle *MeasureStyle
		if front != nil {
			frontMeasureStyle, _ = front.(*MeasureStyle)
		}
		stage.OnAfterMeasureStyleUpdateCallback.OnAfterUpdate(stage, measurestyle, frontMeasureStyle)
	}
}

func (measurestyle *MeasureStyle) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureStyleDeleteCallback != nil {
		var frontMeasureStyle *MeasureStyle
		if front != nil {
			frontMeasureStyle, _ = front.(*MeasureStyle)
		}
		stage.OnAfterMeasureStyleDeleteCallback.OnAfterDelete(stage, measurestyle, frontMeasureStyle)
	}
}

func (multiplerest *MultipleRest) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultipleRestCreateCallback != nil {
		stage.OnAfterMultipleRestCreateCallback.OnAfterCreate(stage, multiplerest)
	}
}

func (multiplerest *MultipleRest) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultipleRestUpdateCallback != nil {
		var frontMultipleRest *MultipleRest
		if front != nil {
			frontMultipleRest, _ = front.(*MultipleRest)
		}
		stage.OnAfterMultipleRestUpdateCallback.OnAfterUpdate(stage, multiplerest, frontMultipleRest)
	}
}

func (multiplerest *MultipleRest) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultipleRestDeleteCallback != nil {
		var frontMultipleRest *MultipleRest
		if front != nil {
			frontMultipleRest, _ = front.(*MultipleRest)
		}
		stage.OnAfterMultipleRestDeleteCallback.OnAfterDelete(stage, multiplerest, frontMultipleRest)
	}
}

func (partclef *PartClef) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartClefCreateCallback != nil {
		stage.OnAfterPartClefCreateCallback.OnAfterCreate(stage, partclef)
	}
}

func (partclef *PartClef) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartClefUpdateCallback != nil {
		var frontPartClef *PartClef
		if front != nil {
			frontPartClef, _ = front.(*PartClef)
		}
		stage.OnAfterPartClefUpdateCallback.OnAfterUpdate(stage, partclef, frontPartClef)
	}
}

func (partclef *PartClef) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartClefDeleteCallback != nil {
		var frontPartClef *PartClef
		if front != nil {
			frontPartClef, _ = front.(*PartClef)
		}
		stage.OnAfterPartClefDeleteCallback.OnAfterDelete(stage, partclef, frontPartClef)
	}
}

func (partsymbol *PartSymbol) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartSymbolCreateCallback != nil {
		stage.OnAfterPartSymbolCreateCallback.OnAfterCreate(stage, partsymbol)
	}
}

func (partsymbol *PartSymbol) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartSymbolUpdateCallback != nil {
		var frontPartSymbol *PartSymbol
		if front != nil {
			frontPartSymbol, _ = front.(*PartSymbol)
		}
		stage.OnAfterPartSymbolUpdateCallback.OnAfterUpdate(stage, partsymbol, frontPartSymbol)
	}
}

func (partsymbol *PartSymbol) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartSymbolDeleteCallback != nil {
		var frontPartSymbol *PartSymbol
		if front != nil {
			frontPartSymbol, _ = front.(*PartSymbol)
		}
		stage.OnAfterPartSymbolDeleteCallback.OnAfterDelete(stage, partsymbol, frontPartSymbol)
	}
}

func (parttranspose *PartTranspose) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartTransposeCreateCallback != nil {
		stage.OnAfterPartTransposeCreateCallback.OnAfterCreate(stage, parttranspose)
	}
}

func (parttranspose *PartTranspose) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartTransposeUpdateCallback != nil {
		var frontPartTranspose *PartTranspose
		if front != nil {
			frontPartTranspose, _ = front.(*PartTranspose)
		}
		stage.OnAfterPartTransposeUpdateCallback.OnAfterUpdate(stage, parttranspose, frontPartTranspose)
	}
}

func (parttranspose *PartTranspose) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartTransposeDeleteCallback != nil {
		var frontPartTranspose *PartTranspose
		if front != nil {
			frontPartTranspose, _ = front.(*PartTranspose)
		}
		stage.OnAfterPartTransposeDeleteCallback.OnAfterDelete(stage, parttranspose, frontPartTranspose)
	}
}

func (slash *Slash) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSlashCreateCallback != nil {
		stage.OnAfterSlashCreateCallback.OnAfterCreate(stage, slash)
	}
}

func (slash *Slash) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSlashUpdateCallback != nil {
		var frontSlash *Slash
		if front != nil {
			frontSlash, _ = front.(*Slash)
		}
		stage.OnAfterSlashUpdateCallback.OnAfterUpdate(stage, slash, frontSlash)
	}
}

func (slash *Slash) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSlashDeleteCallback != nil {
		var frontSlash *Slash
		if front != nil {
			frontSlash, _ = front.(*Slash)
		}
		stage.OnAfterSlashDeleteCallback.OnAfterDelete(stage, slash, frontSlash)
	}
}

func (staffdetails *StaffDetails) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaffDetailsCreateCallback != nil {
		stage.OnAfterStaffDetailsCreateCallback.OnAfterCreate(stage, staffdetails)
	}
}

func (staffdetails *StaffDetails) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffDetailsUpdateCallback != nil {
		var frontStaffDetails *StaffDetails
		if front != nil {
			frontStaffDetails, _ = front.(*StaffDetails)
		}
		stage.OnAfterStaffDetailsUpdateCallback.OnAfterUpdate(stage, staffdetails, frontStaffDetails)
	}
}

func (staffdetails *StaffDetails) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffDetailsDeleteCallback != nil {
		var frontStaffDetails *StaffDetails
		if front != nil {
			frontStaffDetails, _ = front.(*StaffDetails)
		}
		stage.OnAfterStaffDetailsDeleteCallback.OnAfterDelete(stage, staffdetails, frontStaffDetails)
	}
}

func (staffsize *StaffSize) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaffSizeCreateCallback != nil {
		stage.OnAfterStaffSizeCreateCallback.OnAfterCreate(stage, staffsize)
	}
}

func (staffsize *StaffSize) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffSizeUpdateCallback != nil {
		var frontStaffSize *StaffSize
		if front != nil {
			frontStaffSize, _ = front.(*StaffSize)
		}
		stage.OnAfterStaffSizeUpdateCallback.OnAfterUpdate(stage, staffsize, frontStaffSize)
	}
}

func (staffsize *StaffSize) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffSizeDeleteCallback != nil {
		var frontStaffSize *StaffSize
		if front != nil {
			frontStaffSize, _ = front.(*StaffSize)
		}
		stage.OnAfterStaffSizeDeleteCallback.OnAfterDelete(stage, staffsize, frontStaffSize)
	}
}

func (stafftuning *StaffTuning) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaffTuningCreateCallback != nil {
		stage.OnAfterStaffTuningCreateCallback.OnAfterCreate(stage, stafftuning)
	}
}

func (stafftuning *StaffTuning) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffTuningUpdateCallback != nil {
		var frontStaffTuning *StaffTuning
		if front != nil {
			frontStaffTuning, _ = front.(*StaffTuning)
		}
		stage.OnAfterStaffTuningUpdateCallback.OnAfterUpdate(stage, stafftuning, frontStaffTuning)
	}
}

func (stafftuning *StaffTuning) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffTuningDeleteCallback != nil {
		var frontStaffTuning *StaffTuning
		if front != nil {
			frontStaffTuning, _ = front.(*StaffTuning)
		}
		stage.OnAfterStaffTuningDeleteCallback.OnAfterDelete(stage, stafftuning, frontStaffTuning)
	}
}

func (time *Time) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeCreateCallback != nil {
		stage.OnAfterTimeCreateCallback.OnAfterCreate(stage, time)
	}
}

func (time *Time) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeUpdateCallback != nil {
		var frontTime *Time
		if front != nil {
			frontTime, _ = front.(*Time)
		}
		stage.OnAfterTimeUpdateCallback.OnAfterUpdate(stage, time, frontTime)
	}
}

func (time *Time) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeDeleteCallback != nil {
		var frontTime *Time
		if front != nil {
			frontTime, _ = front.(*Time)
		}
		stage.OnAfterTimeDeleteCallback.OnAfterDelete(stage, time, frontTime)
	}
}

func (transpose *Transpose) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTransposeCreateCallback != nil {
		stage.OnAfterTransposeCreateCallback.OnAfterCreate(stage, transpose)
	}
}

func (transpose *Transpose) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransposeUpdateCallback != nil {
		var frontTranspose *Transpose
		if front != nil {
			frontTranspose, _ = front.(*Transpose)
		}
		stage.OnAfterTransposeUpdateCallback.OnAfterUpdate(stage, transpose, frontTranspose)
	}
}

func (transpose *Transpose) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransposeDeleteCallback != nil {
		var frontTranspose *Transpose
		if front != nil {
			frontTranspose, _ = front.(*Transpose)
		}
		stage.OnAfterTransposeDeleteCallback.OnAfterDelete(stage, transpose, frontTranspose)
	}
}

