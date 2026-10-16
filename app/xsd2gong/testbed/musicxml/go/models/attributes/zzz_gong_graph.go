// generated code - do not edit
package attributes

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (attributes *Attributes) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Attributess[attributes]
	return ok
}

func (attributesdirective *AttributesDirective) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AttributesDirectives[attributesdirective]
	return ok
}

func (beatrepeat *BeatRepeat) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BeatRepeats[beatrepeat]
	return ok
}

func (cancel *Cancel) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Cancels[cancel]
	return ok
}

func (clef *Clef) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Clefs[clef]
	return ok
}

func (double *Double) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Doubles[double]
	return ok
}

func (forpart *ForPart) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ForParts[forpart]
	return ok
}

func (interchangeable *Interchangeable) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Interchangeables[interchangeable]
	return ok
}

func (key *Key) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Keys[key]
	return ok
}

func (keyaccidental *KeyAccidental) GongIsStaged(stage *Stage) bool {
	_, ok := stage.KeyAccidentals[keyaccidental]
	return ok
}

func (keyoctave *KeyOctave) GongIsStaged(stage *Stage) bool {
	_, ok := stage.KeyOctaves[keyoctave]
	return ok
}

func (linedetail *LineDetail) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LineDetails[linedetail]
	return ok
}

func (measurerepeat *MeasureRepeat) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeasureRepeats[measurerepeat]
	return ok
}

func (measurestyle *MeasureStyle) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeasureStyles[measurestyle]
	return ok
}

func (multiplerest *MultipleRest) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MultipleRests[multiplerest]
	return ok
}

func (partclef *PartClef) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartClefs[partclef]
	return ok
}

func (partsymbol *PartSymbol) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartSymbols[partsymbol]
	return ok
}

func (parttranspose *PartTranspose) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartTransposes[parttranspose]
	return ok
}

func (slash *Slash) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Slashs[slash]
	return ok
}

func (staffdetails *StaffDetails) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StaffDetailss[staffdetails]
	return ok
}

func (staffsize *StaffSize) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StaffSizes[staffsize]
	return ok
}

func (stafftuning *StaffTuning) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StaffTunings[stafftuning]
	return ok
}

func (time *Time) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Times[time]
	return ok
}

func (transpose *Transpose) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Transposes[transpose]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (attributes *Attributes) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(attributes) {
		return
	}

	attributes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributes.PartSymbol != nil {
		stage.StageBranch(attributes.PartSymbol)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key := range attributes.Key {
		stage.StageBranch(_key)
	}
	for _, _time := range attributes.Time {
		stage.StageBranch(_time)
	}
	for _, _clef := range attributes.Clef {
		stage.StageBranch(_clef)
	}
	for _, _staffdetails := range attributes.StaffDetails {
		stage.StageBranch(_staffdetails)
	}
	for _, _transpose := range attributes.Transpose {
		stage.StageBranch(_transpose)
	}
	for _, _forpart := range attributes.ForPart {
		stage.StageBranch(_forpart)
	}
	for _, _attributesdirective := range attributes.Directive {
		stage.StageBranch(_attributesdirective)
	}
	for _, _measurestyle := range attributes.MeasureStyle {
		stage.StageBranch(_measurestyle)
	}

}

func (attributesdirective *AttributesDirective) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(attributesdirective) {
		return
	}

	attributesdirective.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beatrepeat *BeatRepeat) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(beatrepeat) {
		return
	}

	beatrepeat.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cancel *Cancel) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(cancel) {
		return
	}

	cancel.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clef *Clef) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(clef) {
		return
	}

	clef.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (double *Double) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(double) {
		return
	}

	double.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (forpart *ForPart) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(forpart) {
		return
	}

	forpart.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if forpart.PartClef != nil {
		stage.StageBranch(forpart.PartClef)
	}
	if forpart.PartTranspose != nil {
		stage.StageBranch(forpart.PartTranspose)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (interchangeable *Interchangeable) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(interchangeable) {
		return
	}

	interchangeable.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key *Key) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(key) {
		return
	}

	key.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if key.Cancel != nil {
		stage.StageBranch(key.Cancel)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _keyaccidental := range key.KeyAccidental {
		stage.StageBranch(_keyaccidental)
	}
	for _, _keyoctave := range key.KeyOctave {
		stage.StageBranch(_keyoctave)
	}

}

func (keyaccidental *KeyAccidental) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(keyaccidental) {
		return
	}

	keyaccidental.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyoctave *KeyOctave) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(keyoctave) {
		return
	}

	keyoctave.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linedetail *LineDetail) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(linedetail) {
		return
	}

	linedetail.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measurerepeat *MeasureRepeat) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measurerepeat) {
		return
	}

	measurerepeat.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measurestyle *MeasureStyle) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measurestyle) {
		return
	}

	measurestyle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if measurestyle.MultipleRest != nil {
		stage.StageBranch(measurestyle.MultipleRest)
	}
	if measurestyle.MeasureRepeat != nil {
		stage.StageBranch(measurestyle.MeasureRepeat)
	}
	if measurestyle.BeatRepeat != nil {
		stage.StageBranch(measurestyle.BeatRepeat)
	}
	if measurestyle.Slash != nil {
		stage.StageBranch(measurestyle.Slash)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (multiplerest *MultipleRest) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(multiplerest) {
		return
	}

	multiplerest.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partclef *PartClef) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partclef) {
		return
	}

	partclef.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partsymbol *PartSymbol) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partsymbol) {
		return
	}

	partsymbol.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parttranspose *PartTranspose) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(parttranspose) {
		return
	}

	parttranspose.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parttranspose.Double != nil {
		stage.StageBranch(parttranspose.Double)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slash *Slash) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(slash) {
		return
	}

	slash.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (staffdetails *StaffDetails) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(staffdetails) {
		return
	}

	staffdetails.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if staffdetails.StaffSize != nil {
		stage.StageBranch(staffdetails.StaffSize)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linedetail := range staffdetails.LineDetail {
		stage.StageBranch(_linedetail)
	}
	for _, _stafftuning := range staffdetails.StaffTuning {
		stage.StageBranch(_stafftuning)
	}

}

func (staffsize *StaffSize) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(staffsize) {
		return
	}

	staffsize.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stafftuning *StaffTuning) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stafftuning) {
		return
	}

	stafftuning.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (time *Time) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(time) {
		return
	}

	time.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if time.Interchangeable != nil {
		stage.StageBranch(time.Interchangeable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (transpose *Transpose) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(transpose) {
		return
	}

	transpose.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transpose.Double != nil {
		stage.StageBranch(transpose.Double)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Attributes:
		toT := GongCopyBranchAttributes(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AttributesDirective:
		toT := GongCopyBranchAttributesDirective(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BeatRepeat:
		toT := GongCopyBranchBeatRepeat(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Cancel:
		toT := GongCopyBranchCancel(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Clef:
		toT := GongCopyBranchClef(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Double:
		toT := GongCopyBranchDouble(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ForPart:
		toT := GongCopyBranchForPart(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Interchangeable:
		toT := GongCopyBranchInterchangeable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key:
		toT := GongCopyBranchKey(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *KeyAccidental:
		toT := GongCopyBranchKeyAccidental(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *KeyOctave:
		toT := GongCopyBranchKeyOctave(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LineDetail:
		toT := GongCopyBranchLineDetail(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeasureRepeat:
		toT := GongCopyBranchMeasureRepeat(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeasureStyle:
		toT := GongCopyBranchMeasureStyle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MultipleRest:
		toT := GongCopyBranchMultipleRest(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartClef:
		toT := GongCopyBranchPartClef(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartSymbol:
		toT := GongCopyBranchPartSymbol(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartTranspose:
		toT := GongCopyBranchPartTranspose(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slash:
		toT := GongCopyBranchSlash(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StaffDetails:
		toT := GongCopyBranchStaffDetails(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StaffSize:
		toT := GongCopyBranchStaffSize(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StaffTuning:
		toT := GongCopyBranchStaffTuning(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Time:
		toT := GongCopyBranchTime(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Transpose:
		toT := GongCopyBranchTranspose(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAttributes(mapOrigCopy map[any]any, attributesFrom *Attributes) (attributesTo *Attributes) {
	var alreadyCopied bool
	attributesTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, attributesFrom)
	if alreadyCopied {
		return
	}
	attributesFrom.GongCopyBasicFields(attributesTo)

	//insertion point for the staging of instances referenced by pointers
	if attributesFrom.PartSymbol != nil {
		attributesTo.PartSymbol = GongCopyBranchPartSymbol(mapOrigCopy, attributesFrom.PartSymbol)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key := range attributesFrom.Key {
		attributesTo.Key = append(attributesTo.Key, GongCopyBranchKey(mapOrigCopy, _key))
	}
	for _, _time := range attributesFrom.Time {
		attributesTo.Time = append(attributesTo.Time, GongCopyBranchTime(mapOrigCopy, _time))
	}
	for _, _clef := range attributesFrom.Clef {
		attributesTo.Clef = append(attributesTo.Clef, GongCopyBranchClef(mapOrigCopy, _clef))
	}
	for _, _staffdetails := range attributesFrom.StaffDetails {
		attributesTo.StaffDetails = append(attributesTo.StaffDetails, GongCopyBranchStaffDetails(mapOrigCopy, _staffdetails))
	}
	for _, _transpose := range attributesFrom.Transpose {
		attributesTo.Transpose = append(attributesTo.Transpose, GongCopyBranchTranspose(mapOrigCopy, _transpose))
	}
	for _, _forpart := range attributesFrom.ForPart {
		attributesTo.ForPart = append(attributesTo.ForPart, GongCopyBranchForPart(mapOrigCopy, _forpart))
	}
	for _, _attributesdirective := range attributesFrom.Directive {
		attributesTo.Directive = append(attributesTo.Directive, GongCopyBranchAttributesDirective(mapOrigCopy, _attributesdirective))
	}
	for _, _measurestyle := range attributesFrom.MeasureStyle {
		attributesTo.MeasureStyle = append(attributesTo.MeasureStyle, GongCopyBranchMeasureStyle(mapOrigCopy, _measurestyle))
	}

	return
}

func GongCopyBranchAttributesDirective(mapOrigCopy map[any]any, attributesdirectiveFrom *AttributesDirective) (attributesdirectiveTo *AttributesDirective) {
	var alreadyCopied bool
	attributesdirectiveTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, attributesdirectiveFrom)
	if alreadyCopied {
		return
	}
	attributesdirectiveFrom.GongCopyBasicFields(attributesdirectiveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBeatRepeat(mapOrigCopy map[any]any, beatrepeatFrom *BeatRepeat) (beatrepeatTo *BeatRepeat) {
	var alreadyCopied bool
	beatrepeatTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, beatrepeatFrom)
	if alreadyCopied {
		return
	}
	beatrepeatFrom.GongCopyBasicFields(beatrepeatTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCancel(mapOrigCopy map[any]any, cancelFrom *Cancel) (cancelTo *Cancel) {
	var alreadyCopied bool
	cancelTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cancelFrom)
	if alreadyCopied {
		return
	}
	cancelFrom.GongCopyBasicFields(cancelTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchClef(mapOrigCopy map[any]any, clefFrom *Clef) (clefTo *Clef) {
	var alreadyCopied bool
	clefTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, clefFrom)
	if alreadyCopied {
		return
	}
	clefFrom.GongCopyBasicFields(clefTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDouble(mapOrigCopy map[any]any, doubleFrom *Double) (doubleTo *Double) {
	var alreadyCopied bool
	doubleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, doubleFrom)
	if alreadyCopied {
		return
	}
	doubleFrom.GongCopyBasicFields(doubleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchForPart(mapOrigCopy map[any]any, forpartFrom *ForPart) (forpartTo *ForPart) {
	var alreadyCopied bool
	forpartTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, forpartFrom)
	if alreadyCopied {
		return
	}
	forpartFrom.GongCopyBasicFields(forpartTo)

	//insertion point for the staging of instances referenced by pointers
	if forpartFrom.PartClef != nil {
		forpartTo.PartClef = GongCopyBranchPartClef(mapOrigCopy, forpartFrom.PartClef)
	}
	if forpartFrom.PartTranspose != nil {
		forpartTo.PartTranspose = GongCopyBranchPartTranspose(mapOrigCopy, forpartFrom.PartTranspose)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInterchangeable(mapOrigCopy map[any]any, interchangeableFrom *Interchangeable) (interchangeableTo *Interchangeable) {
	var alreadyCopied bool
	interchangeableTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, interchangeableFrom)
	if alreadyCopied {
		return
	}
	interchangeableFrom.GongCopyBasicFields(interchangeableTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKey(mapOrigCopy map[any]any, keyFrom *Key) (keyTo *Key) {
	var alreadyCopied bool
	keyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, keyFrom)
	if alreadyCopied {
		return
	}
	keyFrom.GongCopyBasicFields(keyTo)

	//insertion point for the staging of instances referenced by pointers
	if keyFrom.Cancel != nil {
		keyTo.Cancel = GongCopyBranchCancel(mapOrigCopy, keyFrom.Cancel)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _keyaccidental := range keyFrom.KeyAccidental {
		keyTo.KeyAccidental = append(keyTo.KeyAccidental, GongCopyBranchKeyAccidental(mapOrigCopy, _keyaccidental))
	}
	for _, _keyoctave := range keyFrom.KeyOctave {
		keyTo.KeyOctave = append(keyTo.KeyOctave, GongCopyBranchKeyOctave(mapOrigCopy, _keyoctave))
	}

	return
}

func GongCopyBranchKeyAccidental(mapOrigCopy map[any]any, keyaccidentalFrom *KeyAccidental) (keyaccidentalTo *KeyAccidental) {
	var alreadyCopied bool
	keyaccidentalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, keyaccidentalFrom)
	if alreadyCopied {
		return
	}
	keyaccidentalFrom.GongCopyBasicFields(keyaccidentalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKeyOctave(mapOrigCopy map[any]any, keyoctaveFrom *KeyOctave) (keyoctaveTo *KeyOctave) {
	var alreadyCopied bool
	keyoctaveTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, keyoctaveFrom)
	if alreadyCopied {
		return
	}
	keyoctaveFrom.GongCopyBasicFields(keyoctaveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLineDetail(mapOrigCopy map[any]any, linedetailFrom *LineDetail) (linedetailTo *LineDetail) {
	var alreadyCopied bool
	linedetailTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linedetailFrom)
	if alreadyCopied {
		return
	}
	linedetailFrom.GongCopyBasicFields(linedetailTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasureRepeat(mapOrigCopy map[any]any, measurerepeatFrom *MeasureRepeat) (measurerepeatTo *MeasureRepeat) {
	var alreadyCopied bool
	measurerepeatTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measurerepeatFrom)
	if alreadyCopied {
		return
	}
	measurerepeatFrom.GongCopyBasicFields(measurerepeatTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasureStyle(mapOrigCopy map[any]any, measurestyleFrom *MeasureStyle) (measurestyleTo *MeasureStyle) {
	var alreadyCopied bool
	measurestyleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measurestyleFrom)
	if alreadyCopied {
		return
	}
	measurestyleFrom.GongCopyBasicFields(measurestyleTo)

	//insertion point for the staging of instances referenced by pointers
	if measurestyleFrom.MultipleRest != nil {
		measurestyleTo.MultipleRest = GongCopyBranchMultipleRest(mapOrigCopy, measurestyleFrom.MultipleRest)
	}
	if measurestyleFrom.MeasureRepeat != nil {
		measurestyleTo.MeasureRepeat = GongCopyBranchMeasureRepeat(mapOrigCopy, measurestyleFrom.MeasureRepeat)
	}
	if measurestyleFrom.BeatRepeat != nil {
		measurestyleTo.BeatRepeat = GongCopyBranchBeatRepeat(mapOrigCopy, measurestyleFrom.BeatRepeat)
	}
	if measurestyleFrom.Slash != nil {
		measurestyleTo.Slash = GongCopyBranchSlash(mapOrigCopy, measurestyleFrom.Slash)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMultipleRest(mapOrigCopy map[any]any, multiplerestFrom *MultipleRest) (multiplerestTo *MultipleRest) {
	var alreadyCopied bool
	multiplerestTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, multiplerestFrom)
	if alreadyCopied {
		return
	}
	multiplerestFrom.GongCopyBasicFields(multiplerestTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartClef(mapOrigCopy map[any]any, partclefFrom *PartClef) (partclefTo *PartClef) {
	var alreadyCopied bool
	partclefTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partclefFrom)
	if alreadyCopied {
		return
	}
	partclefFrom.GongCopyBasicFields(partclefTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartSymbol(mapOrigCopy map[any]any, partsymbolFrom *PartSymbol) (partsymbolTo *PartSymbol) {
	var alreadyCopied bool
	partsymbolTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partsymbolFrom)
	if alreadyCopied {
		return
	}
	partsymbolFrom.GongCopyBasicFields(partsymbolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartTranspose(mapOrigCopy map[any]any, parttransposeFrom *PartTranspose) (parttransposeTo *PartTranspose) {
	var alreadyCopied bool
	parttransposeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, parttransposeFrom)
	if alreadyCopied {
		return
	}
	parttransposeFrom.GongCopyBasicFields(parttransposeTo)

	//insertion point for the staging of instances referenced by pointers
	if parttransposeFrom.Double != nil {
		parttransposeTo.Double = GongCopyBranchDouble(mapOrigCopy, parttransposeFrom.Double)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSlash(mapOrigCopy map[any]any, slashFrom *Slash) (slashTo *Slash) {
	var alreadyCopied bool
	slashTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, slashFrom)
	if alreadyCopied {
		return
	}
	slashFrom.GongCopyBasicFields(slashTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStaffDetails(mapOrigCopy map[any]any, staffdetailsFrom *StaffDetails) (staffdetailsTo *StaffDetails) {
	var alreadyCopied bool
	staffdetailsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, staffdetailsFrom)
	if alreadyCopied {
		return
	}
	staffdetailsFrom.GongCopyBasicFields(staffdetailsTo)

	//insertion point for the staging of instances referenced by pointers
	if staffdetailsFrom.StaffSize != nil {
		staffdetailsTo.StaffSize = GongCopyBranchStaffSize(mapOrigCopy, staffdetailsFrom.StaffSize)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linedetail := range staffdetailsFrom.LineDetail {
		staffdetailsTo.LineDetail = append(staffdetailsTo.LineDetail, GongCopyBranchLineDetail(mapOrigCopy, _linedetail))
	}
	for _, _stafftuning := range staffdetailsFrom.StaffTuning {
		staffdetailsTo.StaffTuning = append(staffdetailsTo.StaffTuning, GongCopyBranchStaffTuning(mapOrigCopy, _stafftuning))
	}

	return
}

func GongCopyBranchStaffSize(mapOrigCopy map[any]any, staffsizeFrom *StaffSize) (staffsizeTo *StaffSize) {
	var alreadyCopied bool
	staffsizeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, staffsizeFrom)
	if alreadyCopied {
		return
	}
	staffsizeFrom.GongCopyBasicFields(staffsizeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStaffTuning(mapOrigCopy map[any]any, stafftuningFrom *StaffTuning) (stafftuningTo *StaffTuning) {
	var alreadyCopied bool
	stafftuningTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stafftuningFrom)
	if alreadyCopied {
		return
	}
	stafftuningFrom.GongCopyBasicFields(stafftuningTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTime(mapOrigCopy map[any]any, timeFrom *Time) (timeTo *Time) {
	var alreadyCopied bool
	timeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timeFrom)
	if alreadyCopied {
		return
	}
	timeFrom.GongCopyBasicFields(timeTo)

	//insertion point for the staging of instances referenced by pointers
	if timeFrom.Interchangeable != nil {
		timeTo.Interchangeable = GongCopyBranchInterchangeable(mapOrigCopy, timeFrom.Interchangeable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTranspose(mapOrigCopy map[any]any, transposeFrom *Transpose) (transposeTo *Transpose) {
	var alreadyCopied bool
	transposeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, transposeFrom)
	if alreadyCopied {
		return
	}
	transposeFrom.GongCopyBasicFields(transposeTo)

	//insertion point for the staging of instances referenced by pointers
	if transposeFrom.Double != nil {
		transposeTo.Double = GongCopyBranchDouble(mapOrigCopy, transposeFrom.Double)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (attributes *Attributes) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(attributes) {
		return
	}

	attributes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributes.PartSymbol != nil {
		stage.UnstageBranch(attributes.PartSymbol)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key := range attributes.Key {
		stage.UnstageBranch(_key)
	}
	for _, _time := range attributes.Time {
		stage.UnstageBranch(_time)
	}
	for _, _clef := range attributes.Clef {
		stage.UnstageBranch(_clef)
	}
	for _, _staffdetails := range attributes.StaffDetails {
		stage.UnstageBranch(_staffdetails)
	}
	for _, _transpose := range attributes.Transpose {
		stage.UnstageBranch(_transpose)
	}
	for _, _forpart := range attributes.ForPart {
		stage.UnstageBranch(_forpart)
	}
	for _, _attributesdirective := range attributes.Directive {
		stage.UnstageBranch(_attributesdirective)
	}
	for _, _measurestyle := range attributes.MeasureStyle {
		stage.UnstageBranch(_measurestyle)
	}

}

func (attributesdirective *AttributesDirective) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(attributesdirective) {
		return
	}

	attributesdirective.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beatrepeat *BeatRepeat) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(beatrepeat) {
		return
	}

	beatrepeat.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cancel *Cancel) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(cancel) {
		return
	}

	cancel.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clef *Clef) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(clef) {
		return
	}

	clef.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (double *Double) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(double) {
		return
	}

	double.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (forpart *ForPart) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(forpart) {
		return
	}

	forpart.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if forpart.PartClef != nil {
		stage.UnstageBranch(forpart.PartClef)
	}
	if forpart.PartTranspose != nil {
		stage.UnstageBranch(forpart.PartTranspose)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (interchangeable *Interchangeable) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(interchangeable) {
		return
	}

	interchangeable.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key *Key) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(key) {
		return
	}

	key.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if key.Cancel != nil {
		stage.UnstageBranch(key.Cancel)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _keyaccidental := range key.KeyAccidental {
		stage.UnstageBranch(_keyaccidental)
	}
	for _, _keyoctave := range key.KeyOctave {
		stage.UnstageBranch(_keyoctave)
	}

}

func (keyaccidental *KeyAccidental) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(keyaccidental) {
		return
	}

	keyaccidental.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (keyoctave *KeyOctave) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(keyoctave) {
		return
	}

	keyoctave.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linedetail *LineDetail) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(linedetail) {
		return
	}

	linedetail.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measurerepeat *MeasureRepeat) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measurerepeat) {
		return
	}

	measurerepeat.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measurestyle *MeasureStyle) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measurestyle) {
		return
	}

	measurestyle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if measurestyle.MultipleRest != nil {
		stage.UnstageBranch(measurestyle.MultipleRest)
	}
	if measurestyle.MeasureRepeat != nil {
		stage.UnstageBranch(measurestyle.MeasureRepeat)
	}
	if measurestyle.BeatRepeat != nil {
		stage.UnstageBranch(measurestyle.BeatRepeat)
	}
	if measurestyle.Slash != nil {
		stage.UnstageBranch(measurestyle.Slash)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (multiplerest *MultipleRest) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(multiplerest) {
		return
	}

	multiplerest.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partclef *PartClef) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partclef) {
		return
	}

	partclef.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partsymbol *PartSymbol) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partsymbol) {
		return
	}

	partsymbol.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (parttranspose *PartTranspose) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(parttranspose) {
		return
	}

	parttranspose.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if parttranspose.Double != nil {
		stage.UnstageBranch(parttranspose.Double)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slash *Slash) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(slash) {
		return
	}

	slash.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (staffdetails *StaffDetails) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(staffdetails) {
		return
	}

	staffdetails.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if staffdetails.StaffSize != nil {
		stage.UnstageBranch(staffdetails.StaffSize)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linedetail := range staffdetails.LineDetail {
		stage.UnstageBranch(_linedetail)
	}
	for _, _stafftuning := range staffdetails.StaffTuning {
		stage.UnstageBranch(_stafftuning)
	}

}

func (staffsize *StaffSize) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(staffsize) {
		return
	}

	staffsize.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stafftuning *StaffTuning) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stafftuning) {
		return
	}

	stafftuning.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (time *Time) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(time) {
		return
	}

	time.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if time.Interchangeable != nil {
		stage.UnstageBranch(time.Interchangeable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (transpose *Transpose) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(transpose) {
		return
	}

	transpose.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transpose.Double != nil {
		stage.UnstageBranch(transpose.Double)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Attributes) GongReconstructPointersFromReferences(stage *Stage, instance *Attributes) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.PartSymbol, stage.PartSymbols_reference, instance.PartSymbol)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Key, stage.Keys_reference, instance.Key)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Time, stage.Times_reference, instance.Time)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Clef, stage.Clefs_reference, instance.Clef)
	__gong__reconstructSliceOfPointersFromReferences(&reference.StaffDetails, stage.StaffDetailss_reference, instance.StaffDetails)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Transpose, stage.Transposes_reference, instance.Transpose)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ForPart, stage.ForParts_reference, instance.ForPart)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Directive, stage.AttributesDirectives_reference, instance.Directive)
	__gong__reconstructSliceOfPointersFromReferences(&reference.MeasureStyle, stage.MeasureStyles_reference, instance.MeasureStyle)
}

func (reference *AttributesDirective) GongReconstructPointersFromReferences(stage *Stage, instance *AttributesDirective) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BeatRepeat) GongReconstructPointersFromReferences(stage *Stage, instance *BeatRepeat) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Cancel) GongReconstructPointersFromReferences(stage *Stage, instance *Cancel) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Clef) GongReconstructPointersFromReferences(stage *Stage, instance *Clef) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Double) GongReconstructPointersFromReferences(stage *Stage, instance *Double) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ForPart) GongReconstructPointersFromReferences(stage *Stage, instance *ForPart) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.PartClef, stage.PartClefs_reference, instance.PartClef)
	__gong__reconstructPointer(&reference.PartTranspose, stage.PartTransposes_reference, instance.PartTranspose)
	// insertion point for slice of pointers field
}

func (reference *Interchangeable) GongReconstructPointersFromReferences(stage *Stage, instance *Interchangeable) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Key) GongReconstructPointersFromReferences(stage *Stage, instance *Key) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Cancel, stage.Cancels_reference, instance.Cancel)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.KeyAccidental, stage.KeyAccidentals_reference, instance.KeyAccidental)
	__gong__reconstructSliceOfPointersFromReferences(&reference.KeyOctave, stage.KeyOctaves_reference, instance.KeyOctave)
}

func (reference *KeyAccidental) GongReconstructPointersFromReferences(stage *Stage, instance *KeyAccidental) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *KeyOctave) GongReconstructPointersFromReferences(stage *Stage, instance *KeyOctave) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LineDetail) GongReconstructPointersFromReferences(stage *Stage, instance *LineDetail) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MeasureRepeat) GongReconstructPointersFromReferences(stage *Stage, instance *MeasureRepeat) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MeasureStyle) GongReconstructPointersFromReferences(stage *Stage, instance *MeasureStyle) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MultipleRest, stage.MultipleRests_reference, instance.MultipleRest)
	__gong__reconstructPointer(&reference.MeasureRepeat, stage.MeasureRepeats_reference, instance.MeasureRepeat)
	__gong__reconstructPointer(&reference.BeatRepeat, stage.BeatRepeats_reference, instance.BeatRepeat)
	__gong__reconstructPointer(&reference.Slash, stage.Slashs_reference, instance.Slash)
	// insertion point for slice of pointers field
}

func (reference *MultipleRest) GongReconstructPointersFromReferences(stage *Stage, instance *MultipleRest) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartClef) GongReconstructPointersFromReferences(stage *Stage, instance *PartClef) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartSymbol) GongReconstructPointersFromReferences(stage *Stage, instance *PartSymbol) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartTranspose) GongReconstructPointersFromReferences(stage *Stage, instance *PartTranspose) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Double, stage.Doubles_reference, instance.Double)
	// insertion point for slice of pointers field
}

func (reference *Slash) GongReconstructPointersFromReferences(stage *Stage, instance *Slash) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StaffDetails) GongReconstructPointersFromReferences(stage *Stage, instance *StaffDetails) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.StaffSize, stage.StaffSizes_reference, instance.StaffSize)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.LineDetail, stage.LineDetails_reference, instance.LineDetail)
	__gong__reconstructSliceOfPointersFromReferences(&reference.StaffTuning, stage.StaffTunings_reference, instance.StaffTuning)
}

func (reference *StaffSize) GongReconstructPointersFromReferences(stage *Stage, instance *StaffSize) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StaffTuning) GongReconstructPointersFromReferences(stage *Stage, instance *StaffTuning) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Time) GongReconstructPointersFromReferences(stage *Stage, instance *Time) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Interchangeable, stage.Interchangeables_reference, instance.Interchangeable)
	// insertion point for slice of pointers field
}

func (reference *Transpose) GongReconstructPointersFromReferences(stage *Stage, instance *Transpose) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Double, stage.Doubles_reference, instance.Double)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Attributes) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.PartSymbol, stage.PartSymbols_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Key, stage.Keys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Time, stage.Times_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Clef, stage.Clefs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.StaffDetails, stage.StaffDetailss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Transpose, stage.Transposes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ForPart, stage.ForParts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Directive, stage.AttributesDirectives_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.MeasureStyle, stage.MeasureStyles_instance)
}

func (reference *AttributesDirective) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BeatRepeat) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Cancel) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Clef) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Double) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ForPart) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.PartClef, stage.PartClefs_instance)
	__gong__reconstructPointerFromInstance(&reference.PartTranspose, stage.PartTransposes_instance)
	// insertion point for slice of pointers fields
}

func (reference *Interchangeable) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Key) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Cancel, stage.Cancels_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.KeyAccidental, stage.KeyAccidentals_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.KeyOctave, stage.KeyOctaves_instance)
}

func (reference *KeyAccidental) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *KeyOctave) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LineDetail) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MeasureRepeat) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MeasureStyle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MultipleRest, stage.MultipleRests_instance)
	__gong__reconstructPointerFromInstance(&reference.MeasureRepeat, stage.MeasureRepeats_instance)
	__gong__reconstructPointerFromInstance(&reference.BeatRepeat, stage.BeatRepeats_instance)
	__gong__reconstructPointerFromInstance(&reference.Slash, stage.Slashs_instance)
	// insertion point for slice of pointers fields
}

func (reference *MultipleRest) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartClef) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartSymbol) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartTranspose) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Double, stage.Doubles_instance)
	// insertion point for slice of pointers fields
}

func (reference *Slash) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StaffDetails) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.StaffSize, stage.StaffSizes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.LineDetail, stage.LineDetails_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.StaffTuning, stage.StaffTunings_instance)
}

func (reference *StaffSize) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StaffTuning) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Time) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Interchangeable, stage.Interchangeables_instance)
	// insertion point for slice of pointers fields
}

func (reference *Transpose) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Double, stage.Doubles_instance)
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attributes *Attributes) GongDiff(stage *Stage, attributesOther *Attributes) (diffs []string) {
	// insertion point for field diffs
	if attributes.Name != attributesOther.Name {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Name"))
	}
	if attributes.Divisions != attributesOther.Divisions {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Divisions"))
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Key", attributesOther.Key, attributes.Key); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Time", attributesOther.Time, attributes.Time); ops != "" {
		diffs = append(diffs, ops)
	}
	if attributes.Staves != attributesOther.Staves {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Staves"))
	}
	if attributes.PartSymbol != attributesOther.PartSymbol {
		diffs = append(diffs, attributes.GongMarshallField(stage, "PartSymbol"))
	}
	if attributes.Instruments != attributesOther.Instruments {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Instruments"))
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Clef", attributesOther.Clef, attributes.Clef); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "StaffDetails", attributesOther.StaffDetails, attributes.StaffDetails); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Transpose", attributesOther.Transpose, attributes.Transpose); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "ForPart", attributesOther.ForPart, attributes.ForPart); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Directive", attributesOther.Directive, attributes.Directive); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "MeasureStyle", attributesOther.MeasureStyle, attributes.MeasureStyle); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attributesdirective *AttributesDirective) GongDiff(stage *Stage, attributesdirectiveOther *AttributesDirective) (diffs []string) {
	// insertion point for field diffs
	if attributesdirective.Name != attributesdirectiveOther.Name {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "Name"))
	}
	if attributesdirective.Lang != attributesdirectiveOther.Lang {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "Lang"))
	}
	if attributesdirective.DefaultX != attributesdirectiveOther.DefaultX {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "DefaultX"))
	}
	if attributesdirective.DefaultY != attributesdirectiveOther.DefaultY {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "DefaultY"))
	}
	if attributesdirective.RelativeX != attributesdirectiveOther.RelativeX {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "RelativeX"))
	}
	if attributesdirective.RelativeY != attributesdirectiveOther.RelativeY {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "RelativeY"))
	}
	if attributesdirective.FontFamily != attributesdirectiveOther.FontFamily {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "FontFamily"))
	}
	if attributesdirective.FontSize != attributesdirectiveOther.FontSize {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "FontSize"))
	}
	if attributesdirective.Color != attributesdirectiveOther.Color {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "Color"))
	}
	if attributesdirective.Value != attributesdirectiveOther.Value {
		diffs = append(diffs, attributesdirective.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beatrepeat *BeatRepeat) GongDiff(stage *Stage, beatrepeatOther *BeatRepeat) (diffs []string) {
	// insertion point for field diffs
	if beatrepeat.Name != beatrepeatOther.Name {
		diffs = append(diffs, beatrepeat.GongMarshallField(stage, "Name"))
	}
	if beatrepeat.Slashes != beatrepeatOther.Slashes {
		diffs = append(diffs, beatrepeat.GongMarshallField(stage, "Slashes"))
	}
	if beatrepeat.ExceptVoice != beatrepeatOther.ExceptVoice {
		diffs = append(diffs, beatrepeat.GongMarshallField(stage, "ExceptVoice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cancel *Cancel) GongDiff(stage *Stage, cancelOther *Cancel) (diffs []string) {
	// insertion point for field diffs
	if cancel.Name != cancelOther.Name {
		diffs = append(diffs, cancel.GongMarshallField(stage, "Name"))
	}
	if cancel.Location != cancelOther.Location {
		diffs = append(diffs, cancel.GongMarshallField(stage, "Location"))
	}
	if cancel.Value != cancelOther.Value {
		diffs = append(diffs, cancel.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (clef *Clef) GongDiff(stage *Stage, clefOther *Clef) (diffs []string) {
	// insertion point for field diffs
	if clef.Name != clefOther.Name {
		diffs = append(diffs, clef.GongMarshallField(stage, "Name"))
	}
	if clef.Number != clefOther.Number {
		diffs = append(diffs, clef.GongMarshallField(stage, "Number"))
	}
	if clef.DefaultX != clefOther.DefaultX {
		diffs = append(diffs, clef.GongMarshallField(stage, "DefaultX"))
	}
	if clef.DefaultY != clefOther.DefaultY {
		diffs = append(diffs, clef.GongMarshallField(stage, "DefaultY"))
	}
	if clef.RelativeX != clefOther.RelativeX {
		diffs = append(diffs, clef.GongMarshallField(stage, "RelativeX"))
	}
	if clef.RelativeY != clefOther.RelativeY {
		diffs = append(diffs, clef.GongMarshallField(stage, "RelativeY"))
	}
	if clef.FontFamily != clefOther.FontFamily {
		diffs = append(diffs, clef.GongMarshallField(stage, "FontFamily"))
	}
	if clef.FontSize != clefOther.FontSize {
		diffs = append(diffs, clef.GongMarshallField(stage, "FontSize"))
	}
	if clef.Color != clefOther.Color {
		diffs = append(diffs, clef.GongMarshallField(stage, "Color"))
	}
	if clef.Id != clefOther.Id {
		diffs = append(diffs, clef.GongMarshallField(stage, "Id"))
	}
	if clef.Sign != clefOther.Sign {
		diffs = append(diffs, clef.GongMarshallField(stage, "Sign"))
	}
	if clef.Line != clefOther.Line {
		diffs = append(diffs, clef.GongMarshallField(stage, "Line"))
	}
	if clef.ClefOctaveChange != clefOther.ClefOctaveChange {
		diffs = append(diffs, clef.GongMarshallField(stage, "ClefOctaveChange"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (double *Double) GongDiff(stage *Stage, doubleOther *Double) (diffs []string) {
	// insertion point for field diffs
	if double.Name != doubleOther.Name {
		diffs = append(diffs, double.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (forpart *ForPart) GongDiff(stage *Stage, forpartOther *ForPart) (diffs []string) {
	// insertion point for field diffs
	if forpart.Name != forpartOther.Name {
		diffs = append(diffs, forpart.GongMarshallField(stage, "Name"))
	}
	if forpart.Number != forpartOther.Number {
		diffs = append(diffs, forpart.GongMarshallField(stage, "Number"))
	}
	if forpart.Id != forpartOther.Id {
		diffs = append(diffs, forpart.GongMarshallField(stage, "Id"))
	}
	if forpart.PartClef != forpartOther.PartClef {
		diffs = append(diffs, forpart.GongMarshallField(stage, "PartClef"))
	}
	if forpart.PartTranspose != forpartOther.PartTranspose {
		diffs = append(diffs, forpart.GongMarshallField(stage, "PartTranspose"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (interchangeable *Interchangeable) GongDiff(stage *Stage, interchangeableOther *Interchangeable) (diffs []string) {
	// insertion point for field diffs
	if interchangeable.Name != interchangeableOther.Name {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Name"))
	}
	if interchangeable.Symbol != interchangeableOther.Symbol {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Symbol"))
	}
	if interchangeable.Separator != interchangeableOther.Separator {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Separator"))
	}
	if interchangeable.TimeRelation != interchangeableOther.TimeRelation {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "TimeRelation"))
	}
	if interchangeable.Beats != interchangeableOther.Beats {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Beats"))
	}
	if interchangeable.BeatType != interchangeableOther.BeatType {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "BeatType"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (key *Key) GongDiff(stage *Stage, keyOther *Key) (diffs []string) {
	// insertion point for field diffs
	if key.Name != keyOther.Name {
		diffs = append(diffs, key.GongMarshallField(stage, "Name"))
	}
	if key.Number != keyOther.Number {
		diffs = append(diffs, key.GongMarshallField(stage, "Number"))
	}
	if key.DefaultX != keyOther.DefaultX {
		diffs = append(diffs, key.GongMarshallField(stage, "DefaultX"))
	}
	if key.DefaultY != keyOther.DefaultY {
		diffs = append(diffs, key.GongMarshallField(stage, "DefaultY"))
	}
	if key.RelativeX != keyOther.RelativeX {
		diffs = append(diffs, key.GongMarshallField(stage, "RelativeX"))
	}
	if key.RelativeY != keyOther.RelativeY {
		diffs = append(diffs, key.GongMarshallField(stage, "RelativeY"))
	}
	if key.FontFamily != keyOther.FontFamily {
		diffs = append(diffs, key.GongMarshallField(stage, "FontFamily"))
	}
	if key.FontSize != keyOther.FontSize {
		diffs = append(diffs, key.GongMarshallField(stage, "FontSize"))
	}
	if key.Color != keyOther.Color {
		diffs = append(diffs, key.GongMarshallField(stage, "Color"))
	}
	if key.Id != keyOther.Id {
		diffs = append(diffs, key.GongMarshallField(stage, "Id"))
	}
	if key.Cancel != keyOther.Cancel {
		diffs = append(diffs, key.GongMarshallField(stage, "Cancel"))
	}
	if key.Fifths != keyOther.Fifths {
		diffs = append(diffs, key.GongMarshallField(stage, "Fifths"))
	}
	if key.Mode != keyOther.Mode {
		diffs = append(diffs, key.GongMarshallField(stage, "Mode"))
	}
	if key.KeyStep != keyOther.KeyStep {
		diffs = append(diffs, key.GongMarshallField(stage, "KeyStep"))
	}
	if key.KeyAlter != keyOther.KeyAlter {
		diffs = append(diffs, key.GongMarshallField(stage, "KeyAlter"))
	}
	if ops := __gong__diffSliceOfPointers(stage, key, "KeyAccidental", keyOther.KeyAccidental, key.KeyAccidental); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, key, "KeyOctave", keyOther.KeyOctave, key.KeyOctave); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (keyaccidental *KeyAccidental) GongDiff(stage *Stage, keyaccidentalOther *KeyAccidental) (diffs []string) {
	// insertion point for field diffs
	if keyaccidental.Name != keyaccidentalOther.Name {
		diffs = append(diffs, keyaccidental.GongMarshallField(stage, "Name"))
	}
	if keyaccidental.Smufl != keyaccidentalOther.Smufl {
		diffs = append(diffs, keyaccidental.GongMarshallField(stage, "Smufl"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (keyoctave *KeyOctave) GongDiff(stage *Stage, keyoctaveOther *KeyOctave) (diffs []string) {
	// insertion point for field diffs
	if keyoctave.Name != keyoctaveOther.Name {
		diffs = append(diffs, keyoctave.GongMarshallField(stage, "Name"))
	}
	if keyoctave.Number != keyoctaveOther.Number {
		diffs = append(diffs, keyoctave.GongMarshallField(stage, "Number"))
	}
	if keyoctave.Value != keyoctaveOther.Value {
		diffs = append(diffs, keyoctave.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (linedetail *LineDetail) GongDiff(stage *Stage, linedetailOther *LineDetail) (diffs []string) {
	// insertion point for field diffs
	if linedetail.Name != linedetailOther.Name {
		diffs = append(diffs, linedetail.GongMarshallField(stage, "Name"))
	}
	if linedetail.Line != linedetailOther.Line {
		diffs = append(diffs, linedetail.GongMarshallField(stage, "Line"))
	}
	if linedetail.Width != linedetailOther.Width {
		diffs = append(diffs, linedetail.GongMarshallField(stage, "Width"))
	}
	if linedetail.Color != linedetailOther.Color {
		diffs = append(diffs, linedetail.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measurerepeat *MeasureRepeat) GongDiff(stage *Stage, measurerepeatOther *MeasureRepeat) (diffs []string) {
	// insertion point for field diffs
	if measurerepeat.Name != measurerepeatOther.Name {
		diffs = append(diffs, measurerepeat.GongMarshallField(stage, "Name"))
	}
	if measurerepeat.Slashes != measurerepeatOther.Slashes {
		diffs = append(diffs, measurerepeat.GongMarshallField(stage, "Slashes"))
	}
	if measurerepeat.Value != measurerepeatOther.Value {
		diffs = append(diffs, measurerepeat.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measurestyle *MeasureStyle) GongDiff(stage *Stage, measurestyleOther *MeasureStyle) (diffs []string) {
	// insertion point for field diffs
	if measurestyle.Name != measurestyleOther.Name {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "Name"))
	}
	if measurestyle.Number != measurestyleOther.Number {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "Number"))
	}
	if measurestyle.FontFamily != measurestyleOther.FontFamily {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "FontFamily"))
	}
	if measurestyle.FontSize != measurestyleOther.FontSize {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "FontSize"))
	}
	if measurestyle.Color != measurestyleOther.Color {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "Color"))
	}
	if measurestyle.Id != measurestyleOther.Id {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "Id"))
	}
	if measurestyle.MultipleRest != measurestyleOther.MultipleRest {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "MultipleRest"))
	}
	if measurestyle.MeasureRepeat != measurestyleOther.MeasureRepeat {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "MeasureRepeat"))
	}
	if measurestyle.BeatRepeat != measurestyleOther.BeatRepeat {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "BeatRepeat"))
	}
	if measurestyle.Slash != measurestyleOther.Slash {
		diffs = append(diffs, measurestyle.GongMarshallField(stage, "Slash"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (multiplerest *MultipleRest) GongDiff(stage *Stage, multiplerestOther *MultipleRest) (diffs []string) {
	// insertion point for field diffs
	if multiplerest.Name != multiplerestOther.Name {
		diffs = append(diffs, multiplerest.GongMarshallField(stage, "Name"))
	}
	if multiplerest.Value != multiplerestOther.Value {
		diffs = append(diffs, multiplerest.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partclef *PartClef) GongDiff(stage *Stage, partclefOther *PartClef) (diffs []string) {
	// insertion point for field diffs
	if partclef.Name != partclefOther.Name {
		diffs = append(diffs, partclef.GongMarshallField(stage, "Name"))
	}
	if partclef.Sign != partclefOther.Sign {
		diffs = append(diffs, partclef.GongMarshallField(stage, "Sign"))
	}
	if partclef.Line != partclefOther.Line {
		diffs = append(diffs, partclef.GongMarshallField(stage, "Line"))
	}
	if partclef.ClefOctaveChange != partclefOther.ClefOctaveChange {
		diffs = append(diffs, partclef.GongMarshallField(stage, "ClefOctaveChange"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partsymbol *PartSymbol) GongDiff(stage *Stage, partsymbolOther *PartSymbol) (diffs []string) {
	// insertion point for field diffs
	if partsymbol.Name != partsymbolOther.Name {
		diffs = append(diffs, partsymbol.GongMarshallField(stage, "Name"))
	}
	if partsymbol.TopStaff != partsymbolOther.TopStaff {
		diffs = append(diffs, partsymbol.GongMarshallField(stage, "TopStaff"))
	}
	if partsymbol.BottomStaff != partsymbolOther.BottomStaff {
		diffs = append(diffs, partsymbol.GongMarshallField(stage, "BottomStaff"))
	}
	if partsymbol.DefaultX != partsymbolOther.DefaultX {
		diffs = append(diffs, partsymbol.GongMarshallField(stage, "DefaultX"))
	}
	if partsymbol.DefaultY != partsymbolOther.DefaultY {
		diffs = append(diffs, partsymbol.GongMarshallField(stage, "DefaultY"))
	}
	if partsymbol.RelativeX != partsymbolOther.RelativeX {
		diffs = append(diffs, partsymbol.GongMarshallField(stage, "RelativeX"))
	}
	if partsymbol.RelativeY != partsymbolOther.RelativeY {
		diffs = append(diffs, partsymbol.GongMarshallField(stage, "RelativeY"))
	}
	if partsymbol.Color != partsymbolOther.Color {
		diffs = append(diffs, partsymbol.GongMarshallField(stage, "Color"))
	}
	if partsymbol.Value != partsymbolOther.Value {
		diffs = append(diffs, partsymbol.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (parttranspose *PartTranspose) GongDiff(stage *Stage, parttransposeOther *PartTranspose) (diffs []string) {
	// insertion point for field diffs
	if parttranspose.Name != parttransposeOther.Name {
		diffs = append(diffs, parttranspose.GongMarshallField(stage, "Name"))
	}
	if parttranspose.Diatonic != parttransposeOther.Diatonic {
		diffs = append(diffs, parttranspose.GongMarshallField(stage, "Diatonic"))
	}
	if parttranspose.Chromatic != parttransposeOther.Chromatic {
		diffs = append(diffs, parttranspose.GongMarshallField(stage, "Chromatic"))
	}
	if parttranspose.OctaveChange != parttransposeOther.OctaveChange {
		diffs = append(diffs, parttranspose.GongMarshallField(stage, "OctaveChange"))
	}
	if parttranspose.Double != parttransposeOther.Double {
		diffs = append(diffs, parttranspose.GongMarshallField(stage, "Double"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (slash *Slash) GongDiff(stage *Stage, slashOther *Slash) (diffs []string) {
	// insertion point for field diffs
	if slash.Name != slashOther.Name {
		diffs = append(diffs, slash.GongMarshallField(stage, "Name"))
	}
	if slash.ExceptVoice != slashOther.ExceptVoice {
		diffs = append(diffs, slash.GongMarshallField(stage, "ExceptVoice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (staffdetails *StaffDetails) GongDiff(stage *Stage, staffdetailsOther *StaffDetails) (diffs []string) {
	// insertion point for field diffs
	if staffdetails.Name != staffdetailsOther.Name {
		diffs = append(diffs, staffdetails.GongMarshallField(stage, "Name"))
	}
	if staffdetails.Number != staffdetailsOther.Number {
		diffs = append(diffs, staffdetails.GongMarshallField(stage, "Number"))
	}
	if staffdetails.ShowFrets != staffdetailsOther.ShowFrets {
		diffs = append(diffs, staffdetails.GongMarshallField(stage, "ShowFrets"))
	}
	if staffdetails.StaffType != staffdetailsOther.StaffType {
		diffs = append(diffs, staffdetails.GongMarshallField(stage, "StaffType"))
	}
	if staffdetails.StaffLines != staffdetailsOther.StaffLines {
		diffs = append(diffs, staffdetails.GongMarshallField(stage, "StaffLines"))
	}
	if ops := __gong__diffSliceOfPointers(stage, staffdetails, "LineDetail", staffdetailsOther.LineDetail, staffdetails.LineDetail); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, staffdetails, "StaffTuning", staffdetailsOther.StaffTuning, staffdetails.StaffTuning); ops != "" {
		diffs = append(diffs, ops)
	}
	if staffdetails.Capo != staffdetailsOther.Capo {
		diffs = append(diffs, staffdetails.GongMarshallField(stage, "Capo"))
	}
	if staffdetails.StaffSize != staffdetailsOther.StaffSize {
		diffs = append(diffs, staffdetails.GongMarshallField(stage, "StaffSize"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (staffsize *StaffSize) GongDiff(stage *Stage, staffsizeOther *StaffSize) (diffs []string) {
	// insertion point for field diffs
	if staffsize.Name != staffsizeOther.Name {
		diffs = append(diffs, staffsize.GongMarshallField(stage, "Name"))
	}
	if staffsize.Scaling != staffsizeOther.Scaling {
		diffs = append(diffs, staffsize.GongMarshallField(stage, "Scaling"))
	}
	if staffsize.Value != staffsizeOther.Value {
		diffs = append(diffs, staffsize.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stafftuning *StaffTuning) GongDiff(stage *Stage, stafftuningOther *StaffTuning) (diffs []string) {
	// insertion point for field diffs
	if stafftuning.Name != stafftuningOther.Name {
		diffs = append(diffs, stafftuning.GongMarshallField(stage, "Name"))
	}
	if stafftuning.Line != stafftuningOther.Line {
		diffs = append(diffs, stafftuning.GongMarshallField(stage, "Line"))
	}
	if stafftuning.TuningStep != stafftuningOther.TuningStep {
		diffs = append(diffs, stafftuning.GongMarshallField(stage, "TuningStep"))
	}
	if stafftuning.TuningAlter != stafftuningOther.TuningAlter {
		diffs = append(diffs, stafftuning.GongMarshallField(stage, "TuningAlter"))
	}
	if stafftuning.TuningOctave != stafftuningOther.TuningOctave {
		diffs = append(diffs, stafftuning.GongMarshallField(stage, "TuningOctave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (time *Time) GongDiff(stage *Stage, timeOther *Time) (diffs []string) {
	// insertion point for field diffs
	if time.Name != timeOther.Name {
		diffs = append(diffs, time.GongMarshallField(stage, "Name"))
	}
	if time.Number != timeOther.Number {
		diffs = append(diffs, time.GongMarshallField(stage, "Number"))
	}
	if time.Symbol != timeOther.Symbol {
		diffs = append(diffs, time.GongMarshallField(stage, "Symbol"))
	}
	if time.Separator != timeOther.Separator {
		diffs = append(diffs, time.GongMarshallField(stage, "Separator"))
	}
	if time.DefaultX != timeOther.DefaultX {
		diffs = append(diffs, time.GongMarshallField(stage, "DefaultX"))
	}
	if time.DefaultY != timeOther.DefaultY {
		diffs = append(diffs, time.GongMarshallField(stage, "DefaultY"))
	}
	if time.RelativeX != timeOther.RelativeX {
		diffs = append(diffs, time.GongMarshallField(stage, "RelativeX"))
	}
	if time.RelativeY != timeOther.RelativeY {
		diffs = append(diffs, time.GongMarshallField(stage, "RelativeY"))
	}
	if time.FontFamily != timeOther.FontFamily {
		diffs = append(diffs, time.GongMarshallField(stage, "FontFamily"))
	}
	if time.FontSize != timeOther.FontSize {
		diffs = append(diffs, time.GongMarshallField(stage, "FontSize"))
	}
	if time.Color != timeOther.Color {
		diffs = append(diffs, time.GongMarshallField(stage, "Color"))
	}
	if time.Id != timeOther.Id {
		diffs = append(diffs, time.GongMarshallField(stage, "Id"))
	}
	if time.Beats != timeOther.Beats {
		diffs = append(diffs, time.GongMarshallField(stage, "Beats"))
	}
	if time.BeatType != timeOther.BeatType {
		diffs = append(diffs, time.GongMarshallField(stage, "BeatType"))
	}
	if time.Interchangeable != timeOther.Interchangeable {
		diffs = append(diffs, time.GongMarshallField(stage, "Interchangeable"))
	}
	if time.SenzaMisura != timeOther.SenzaMisura {
		diffs = append(diffs, time.GongMarshallField(stage, "SenzaMisura"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (transpose *Transpose) GongDiff(stage *Stage, transposeOther *Transpose) (diffs []string) {
	// insertion point for field diffs
	if transpose.Name != transposeOther.Name {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Name"))
	}
	if transpose.Number != transposeOther.Number {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Number"))
	}
	if transpose.Id != transposeOther.Id {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Id"))
	}
	if transpose.Diatonic != transposeOther.Diatonic {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Diatonic"))
	}
	if transpose.Chromatic != transposeOther.Chromatic {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Chromatic"))
	}
	if transpose.OctaveChange != transposeOther.OctaveChange {
		diffs = append(diffs, transpose.GongMarshallField(stage, "OctaveChange"))
	}
	if transpose.Double != transposeOther.Double {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Double"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
