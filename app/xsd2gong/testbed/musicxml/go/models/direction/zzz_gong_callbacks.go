// generated code - do not edit
package direction

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
func (accord *Accord) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAccordCreateCallback != nil {
		stage.OnAfterAccordCreateCallback.OnAfterCreate(stage, accord)
	}
}

func (accord *Accord) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccordUpdateCallback != nil {
		var frontAccord *Accord
		if front != nil {
			frontAccord, _ = front.(*Accord)
		}
		stage.OnAfterAccordUpdateCallback.OnAfterUpdate(stage, accord, frontAccord)
	}
}

func (accord *Accord) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccordDeleteCallback != nil {
		var frontAccord *Accord
		if front != nil {
			frontAccord, _ = front.(*Accord)
		}
		stage.OnAfterAccordDeleteCallback.OnAfterDelete(stage, accord, frontAccord)
	}
}

func (accordionregistration *AccordionRegistration) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAccordionRegistrationCreateCallback != nil {
		stage.OnAfterAccordionRegistrationCreateCallback.OnAfterCreate(stage, accordionregistration)
	}
}

func (accordionregistration *AccordionRegistration) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccordionRegistrationUpdateCallback != nil {
		var frontAccordionRegistration *AccordionRegistration
		if front != nil {
			frontAccordionRegistration, _ = front.(*AccordionRegistration)
		}
		stage.OnAfterAccordionRegistrationUpdateCallback.OnAfterUpdate(stage, accordionregistration, frontAccordionRegistration)
	}
}

func (accordionregistration *AccordionRegistration) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccordionRegistrationDeleteCallback != nil {
		var frontAccordionRegistration *AccordionRegistration
		if front != nil {
			frontAccordionRegistration, _ = front.(*AccordionRegistration)
		}
		stage.OnAfterAccordionRegistrationDeleteCallback.OnAfterDelete(stage, accordionregistration, frontAccordionRegistration)
	}
}

func (barre *Barre) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBarreCreateCallback != nil {
		stage.OnAfterBarreCreateCallback.OnAfterCreate(stage, barre)
	}
}

func (barre *Barre) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarreUpdateCallback != nil {
		var frontBarre *Barre
		if front != nil {
			frontBarre, _ = front.(*Barre)
		}
		stage.OnAfterBarreUpdateCallback.OnAfterUpdate(stage, barre, frontBarre)
	}
}

func (barre *Barre) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarreDeleteCallback != nil {
		var frontBarre *Barre
		if front != nil {
			frontBarre, _ = front.(*Barre)
		}
		stage.OnAfterBarreDeleteCallback.OnAfterDelete(stage, barre, frontBarre)
	}
}

func (bass *Bass) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBassCreateCallback != nil {
		stage.OnAfterBassCreateCallback.OnAfterCreate(stage, bass)
	}
}

func (bass *Bass) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBassUpdateCallback != nil {
		var frontBass *Bass
		if front != nil {
			frontBass, _ = front.(*Bass)
		}
		stage.OnAfterBassUpdateCallback.OnAfterUpdate(stage, bass, frontBass)
	}
}

func (bass *Bass) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBassDeleteCallback != nil {
		var frontBass *Bass
		if front != nil {
			frontBass, _ = front.(*Bass)
		}
		stage.OnAfterBassDeleteCallback.OnAfterDelete(stage, bass, frontBass)
	}
}

func (bassstep *BassStep) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBassStepCreateCallback != nil {
		stage.OnAfterBassStepCreateCallback.OnAfterCreate(stage, bassstep)
	}
}

func (bassstep *BassStep) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBassStepUpdateCallback != nil {
		var frontBassStep *BassStep
		if front != nil {
			frontBassStep, _ = front.(*BassStep)
		}
		stage.OnAfterBassStepUpdateCallback.OnAfterUpdate(stage, bassstep, frontBassStep)
	}
}

func (bassstep *BassStep) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBassStepDeleteCallback != nil {
		var frontBassStep *BassStep
		if front != nil {
			frontBassStep, _ = front.(*BassStep)
		}
		stage.OnAfterBassStepDeleteCallback.OnAfterDelete(stage, bassstep, frontBassStep)
	}
}

func (beatunittied *BeatUnitTied) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBeatUnitTiedCreateCallback != nil {
		stage.OnAfterBeatUnitTiedCreateCallback.OnAfterCreate(stage, beatunittied)
	}
}

func (beatunittied *BeatUnitTied) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeatUnitTiedUpdateCallback != nil {
		var frontBeatUnitTied *BeatUnitTied
		if front != nil {
			frontBeatUnitTied, _ = front.(*BeatUnitTied)
		}
		stage.OnAfterBeatUnitTiedUpdateCallback.OnAfterUpdate(stage, beatunittied, frontBeatUnitTied)
	}
}

func (beatunittied *BeatUnitTied) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeatUnitTiedDeleteCallback != nil {
		var frontBeatUnitTied *BeatUnitTied
		if front != nil {
			frontBeatUnitTied, _ = front.(*BeatUnitTied)
		}
		stage.OnAfterBeatUnitTiedDeleteCallback.OnAfterDelete(stage, beatunittied, frontBeatUnitTied)
	}
}

func (beater *Beater) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBeaterCreateCallback != nil {
		stage.OnAfterBeaterCreateCallback.OnAfterCreate(stage, beater)
	}
}

func (beater *Beater) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeaterUpdateCallback != nil {
		var frontBeater *Beater
		if front != nil {
			frontBeater, _ = front.(*Beater)
		}
		stage.OnAfterBeaterUpdateCallback.OnAfterUpdate(stage, beater, frontBeater)
	}
}

func (beater *Beater) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeaterDeleteCallback != nil {
		var frontBeater *Beater
		if front != nil {
			frontBeater, _ = front.(*Beater)
		}
		stage.OnAfterBeaterDeleteCallback.OnAfterDelete(stage, beater, frontBeater)
	}
}

func (bracket *Bracket) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBracketCreateCallback != nil {
		stage.OnAfterBracketCreateCallback.OnAfterCreate(stage, bracket)
	}
}

func (bracket *Bracket) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBracketUpdateCallback != nil {
		var frontBracket *Bracket
		if front != nil {
			frontBracket, _ = front.(*Bracket)
		}
		stage.OnAfterBracketUpdateCallback.OnAfterUpdate(stage, bracket, frontBracket)
	}
}

func (bracket *Bracket) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBracketDeleteCallback != nil {
		var frontBracket *Bracket
		if front != nil {
			frontBracket, _ = front.(*Bracket)
		}
		stage.OnAfterBracketDeleteCallback.OnAfterDelete(stage, bracket, frontBracket)
	}
}

func (dashes *Dashes) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDashesCreateCallback != nil {
		stage.OnAfterDashesCreateCallback.OnAfterCreate(stage, dashes)
	}
}

func (dashes *Dashes) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDashesUpdateCallback != nil {
		var frontDashes *Dashes
		if front != nil {
			frontDashes, _ = front.(*Dashes)
		}
		stage.OnAfterDashesUpdateCallback.OnAfterUpdate(stage, dashes, frontDashes)
	}
}

func (dashes *Dashes) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDashesDeleteCallback != nil {
		var frontDashes *Dashes
		if front != nil {
			frontDashes, _ = front.(*Dashes)
		}
		stage.OnAfterDashesDeleteCallback.OnAfterDelete(stage, dashes, frontDashes)
	}
}

func (degree *Degree) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDegreeCreateCallback != nil {
		stage.OnAfterDegreeCreateCallback.OnAfterCreate(stage, degree)
	}
}

func (degree *Degree) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreeUpdateCallback != nil {
		var frontDegree *Degree
		if front != nil {
			frontDegree, _ = front.(*Degree)
		}
		stage.OnAfterDegreeUpdateCallback.OnAfterUpdate(stage, degree, frontDegree)
	}
}

func (degree *Degree) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreeDeleteCallback != nil {
		var frontDegree *Degree
		if front != nil {
			frontDegree, _ = front.(*Degree)
		}
		stage.OnAfterDegreeDeleteCallback.OnAfterDelete(stage, degree, frontDegree)
	}
}

func (degreealter *DegreeAlter) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDegreeAlterCreateCallback != nil {
		stage.OnAfterDegreeAlterCreateCallback.OnAfterCreate(stage, degreealter)
	}
}

func (degreealter *DegreeAlter) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreeAlterUpdateCallback != nil {
		var frontDegreeAlter *DegreeAlter
		if front != nil {
			frontDegreeAlter, _ = front.(*DegreeAlter)
		}
		stage.OnAfterDegreeAlterUpdateCallback.OnAfterUpdate(stage, degreealter, frontDegreeAlter)
	}
}

func (degreealter *DegreeAlter) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreeAlterDeleteCallback != nil {
		var frontDegreeAlter *DegreeAlter
		if front != nil {
			frontDegreeAlter, _ = front.(*DegreeAlter)
		}
		stage.OnAfterDegreeAlterDeleteCallback.OnAfterDelete(stage, degreealter, frontDegreeAlter)
	}
}

func (degreetype *DegreeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDegreeTypeCreateCallback != nil {
		stage.OnAfterDegreeTypeCreateCallback.OnAfterCreate(stage, degreetype)
	}
}

func (degreetype *DegreeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreeTypeUpdateCallback != nil {
		var frontDegreeType *DegreeType
		if front != nil {
			frontDegreeType, _ = front.(*DegreeType)
		}
		stage.OnAfterDegreeTypeUpdateCallback.OnAfterUpdate(stage, degreetype, frontDegreeType)
	}
}

func (degreetype *DegreeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreeTypeDeleteCallback != nil {
		var frontDegreeType *DegreeType
		if front != nil {
			frontDegreeType, _ = front.(*DegreeType)
		}
		stage.OnAfterDegreeTypeDeleteCallback.OnAfterDelete(stage, degreetype, frontDegreeType)
	}
}

func (degreevalue *DegreeValue) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDegreeValueCreateCallback != nil {
		stage.OnAfterDegreeValueCreateCallback.OnAfterCreate(stage, degreevalue)
	}
}

func (degreevalue *DegreeValue) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreeValueUpdateCallback != nil {
		var frontDegreeValue *DegreeValue
		if front != nil {
			frontDegreeValue, _ = front.(*DegreeValue)
		}
		stage.OnAfterDegreeValueUpdateCallback.OnAfterUpdate(stage, degreevalue, frontDegreeValue)
	}
}

func (degreevalue *DegreeValue) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreeValueDeleteCallback != nil {
		var frontDegreeValue *DegreeValue
		if front != nil {
			frontDegreeValue, _ = front.(*DegreeValue)
		}
		stage.OnAfterDegreeValueDeleteCallback.OnAfterDelete(stage, degreevalue, frontDegreeValue)
	}
}

func (direction *Direction) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionCreateCallback != nil {
		stage.OnAfterDirectionCreateCallback.OnAfterCreate(stage, direction)
	}
}

func (direction *Direction) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionUpdateCallback != nil {
		var frontDirection *Direction
		if front != nil {
			frontDirection, _ = front.(*Direction)
		}
		stage.OnAfterDirectionUpdateCallback.OnAfterUpdate(stage, direction, frontDirection)
	}
}

func (direction *Direction) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionDeleteCallback != nil {
		var frontDirection *Direction
		if front != nil {
			frontDirection, _ = front.(*Direction)
		}
		stage.OnAfterDirectionDeleteCallback.OnAfterDelete(stage, direction, frontDirection)
	}
}

func (directiontype *DirectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionTypeCreateCallback != nil {
		stage.OnAfterDirectionTypeCreateCallback.OnAfterCreate(stage, directiontype)
	}
}

func (directiontype *DirectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionTypeUpdateCallback != nil {
		var frontDirectionType *DirectionType
		if front != nil {
			frontDirectionType, _ = front.(*DirectionType)
		}
		stage.OnAfterDirectionTypeUpdateCallback.OnAfterUpdate(stage, directiontype, frontDirectionType)
	}
}

func (directiontype *DirectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionTypeDeleteCallback != nil {
		var frontDirectionType *DirectionType
		if front != nil {
			frontDirectionType, _ = front.(*DirectionType)
		}
		stage.OnAfterDirectionTypeDeleteCallback.OnAfterDelete(stage, directiontype, frontDirectionType)
	}
}

func (effect *Effect) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEffectCreateCallback != nil {
		stage.OnAfterEffectCreateCallback.OnAfterCreate(stage, effect)
	}
}

func (effect *Effect) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEffectUpdateCallback != nil {
		var frontEffect *Effect
		if front != nil {
			frontEffect, _ = front.(*Effect)
		}
		stage.OnAfterEffectUpdateCallback.OnAfterUpdate(stage, effect, frontEffect)
	}
}

func (effect *Effect) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEffectDeleteCallback != nil {
		var frontEffect *Effect
		if front != nil {
			frontEffect, _ = front.(*Effect)
		}
		stage.OnAfterEffectDeleteCallback.OnAfterDelete(stage, effect, frontEffect)
	}
}

func (feature *Feature) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFeatureCreateCallback != nil {
		stage.OnAfterFeatureCreateCallback.OnAfterCreate(stage, feature)
	}
}

func (feature *Feature) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureUpdateCallback != nil {
		var frontFeature *Feature
		if front != nil {
			frontFeature, _ = front.(*Feature)
		}
		stage.OnAfterFeatureUpdateCallback.OnAfterUpdate(stage, feature, frontFeature)
	}
}

func (feature *Feature) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureDeleteCallback != nil {
		var frontFeature *Feature
		if front != nil {
			frontFeature, _ = front.(*Feature)
		}
		stage.OnAfterFeatureDeleteCallback.OnAfterDelete(stage, feature, frontFeature)
	}
}

func (firstfret *FirstFret) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFirstFretCreateCallback != nil {
		stage.OnAfterFirstFretCreateCallback.OnAfterCreate(stage, firstfret)
	}
}

func (firstfret *FirstFret) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFirstFretUpdateCallback != nil {
		var frontFirstFret *FirstFret
		if front != nil {
			frontFirstFret, _ = front.(*FirstFret)
		}
		stage.OnAfterFirstFretUpdateCallback.OnAfterUpdate(stage, firstfret, frontFirstFret)
	}
}

func (firstfret *FirstFret) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFirstFretDeleteCallback != nil {
		var frontFirstFret *FirstFret
		if front != nil {
			frontFirstFret, _ = front.(*FirstFret)
		}
		stage.OnAfterFirstFretDeleteCallback.OnAfterDelete(stage, firstfret, frontFirstFret)
	}
}

func (frame *Frame) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFrameCreateCallback != nil {
		stage.OnAfterFrameCreateCallback.OnAfterCreate(stage, frame)
	}
}

func (frame *Frame) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFrameUpdateCallback != nil {
		var frontFrame *Frame
		if front != nil {
			frontFrame, _ = front.(*Frame)
		}
		stage.OnAfterFrameUpdateCallback.OnAfterUpdate(stage, frame, frontFrame)
	}
}

func (frame *Frame) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFrameDeleteCallback != nil {
		var frontFrame *Frame
		if front != nil {
			frontFrame, _ = front.(*Frame)
		}
		stage.OnAfterFrameDeleteCallback.OnAfterDelete(stage, frame, frontFrame)
	}
}

func (framenote *FrameNote) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFrameNoteCreateCallback != nil {
		stage.OnAfterFrameNoteCreateCallback.OnAfterCreate(stage, framenote)
	}
}

func (framenote *FrameNote) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFrameNoteUpdateCallback != nil {
		var frontFrameNote *FrameNote
		if front != nil {
			frontFrameNote, _ = front.(*FrameNote)
		}
		stage.OnAfterFrameNoteUpdateCallback.OnAfterUpdate(stage, framenote, frontFrameNote)
	}
}

func (framenote *FrameNote) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFrameNoteDeleteCallback != nil {
		var frontFrameNote *FrameNote
		if front != nil {
			frontFrameNote, _ = front.(*FrameNote)
		}
		stage.OnAfterFrameNoteDeleteCallback.OnAfterDelete(stage, framenote, frontFrameNote)
	}
}

func (glass *Glass) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGlassCreateCallback != nil {
		stage.OnAfterGlassCreateCallback.OnAfterCreate(stage, glass)
	}
}

func (glass *Glass) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlassUpdateCallback != nil {
		var frontGlass *Glass
		if front != nil {
			frontGlass, _ = front.(*Glass)
		}
		stage.OnAfterGlassUpdateCallback.OnAfterUpdate(stage, glass, frontGlass)
	}
}

func (glass *Glass) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlassDeleteCallback != nil {
		var frontGlass *Glass
		if front != nil {
			frontGlass, _ = front.(*Glass)
		}
		stage.OnAfterGlassDeleteCallback.OnAfterDelete(stage, glass, frontGlass)
	}
}

func (grouping *Grouping) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupingCreateCallback != nil {
		stage.OnAfterGroupingCreateCallback.OnAfterCreate(stage, grouping)
	}
}

func (grouping *Grouping) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupingUpdateCallback != nil {
		var frontGrouping *Grouping
		if front != nil {
			frontGrouping, _ = front.(*Grouping)
		}
		stage.OnAfterGroupingUpdateCallback.OnAfterUpdate(stage, grouping, frontGrouping)
	}
}

func (grouping *Grouping) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupingDeleteCallback != nil {
		var frontGrouping *Grouping
		if front != nil {
			frontGrouping, _ = front.(*Grouping)
		}
		stage.OnAfterGroupingDeleteCallback.OnAfterDelete(stage, grouping, frontGrouping)
	}
}

func (harmony *Harmony) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarmonyCreateCallback != nil {
		stage.OnAfterHarmonyCreateCallback.OnAfterCreate(stage, harmony)
	}
}

func (harmony *Harmony) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonyUpdateCallback != nil {
		var frontHarmony *Harmony
		if front != nil {
			frontHarmony, _ = front.(*Harmony)
		}
		stage.OnAfterHarmonyUpdateCallback.OnAfterUpdate(stage, harmony, frontHarmony)
	}
}

func (harmony *Harmony) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonyDeleteCallback != nil {
		var frontHarmony *Harmony
		if front != nil {
			frontHarmony, _ = front.(*Harmony)
		}
		stage.OnAfterHarmonyDeleteCallback.OnAfterDelete(stage, harmony, frontHarmony)
	}
}

func (harmonyalter *HarmonyAlter) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarmonyAlterCreateCallback != nil {
		stage.OnAfterHarmonyAlterCreateCallback.OnAfterCreate(stage, harmonyalter)
	}
}

func (harmonyalter *HarmonyAlter) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonyAlterUpdateCallback != nil {
		var frontHarmonyAlter *HarmonyAlter
		if front != nil {
			frontHarmonyAlter, _ = front.(*HarmonyAlter)
		}
		stage.OnAfterHarmonyAlterUpdateCallback.OnAfterUpdate(stage, harmonyalter, frontHarmonyAlter)
	}
}

func (harmonyalter *HarmonyAlter) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonyAlterDeleteCallback != nil {
		var frontHarmonyAlter *HarmonyAlter
		if front != nil {
			frontHarmonyAlter, _ = front.(*HarmonyAlter)
		}
		stage.OnAfterHarmonyAlterDeleteCallback.OnAfterDelete(stage, harmonyalter, frontHarmonyAlter)
	}
}

func (harppedals *HarpPedals) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarpPedalsCreateCallback != nil {
		stage.OnAfterHarpPedalsCreateCallback.OnAfterCreate(stage, harppedals)
	}
}

func (harppedals *HarpPedals) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarpPedalsUpdateCallback != nil {
		var frontHarpPedals *HarpPedals
		if front != nil {
			frontHarpPedals, _ = front.(*HarpPedals)
		}
		stage.OnAfterHarpPedalsUpdateCallback.OnAfterUpdate(stage, harppedals, frontHarpPedals)
	}
}

func (harppedals *HarpPedals) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarpPedalsDeleteCallback != nil {
		var frontHarpPedals *HarpPedals
		if front != nil {
			frontHarpPedals, _ = front.(*HarpPedals)
		}
		stage.OnAfterHarpPedalsDeleteCallback.OnAfterDelete(stage, harppedals, frontHarpPedals)
	}
}

func (image *Image) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterImageCreateCallback != nil {
		stage.OnAfterImageCreateCallback.OnAfterCreate(stage, image)
	}
}

func (image *Image) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageUpdateCallback != nil {
		var frontImage *Image
		if front != nil {
			frontImage, _ = front.(*Image)
		}
		stage.OnAfterImageUpdateCallback.OnAfterUpdate(stage, image, frontImage)
	}
}

func (image *Image) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageDeleteCallback != nil {
		var frontImage *Image
		if front != nil {
			frontImage, _ = front.(*Image)
		}
		stage.OnAfterImageDeleteCallback.OnAfterDelete(stage, image, frontImage)
	}
}

func (instrumentchange *InstrumentChange) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInstrumentChangeCreateCallback != nil {
		stage.OnAfterInstrumentChangeCreateCallback.OnAfterCreate(stage, instrumentchange)
	}
}

func (instrumentchange *InstrumentChange) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentChangeUpdateCallback != nil {
		var frontInstrumentChange *InstrumentChange
		if front != nil {
			frontInstrumentChange, _ = front.(*InstrumentChange)
		}
		stage.OnAfterInstrumentChangeUpdateCallback.OnAfterUpdate(stage, instrumentchange, frontInstrumentChange)
	}
}

func (instrumentchange *InstrumentChange) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentChangeDeleteCallback != nil {
		var frontInstrumentChange *InstrumentChange
		if front != nil {
			frontInstrumentChange, _ = front.(*InstrumentChange)
		}
		stage.OnAfterInstrumentChangeDeleteCallback.OnAfterDelete(stage, instrumentchange, frontInstrumentChange)
	}
}

func (inversion *Inversion) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInversionCreateCallback != nil {
		stage.OnAfterInversionCreateCallback.OnAfterCreate(stage, inversion)
	}
}

func (inversion *Inversion) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInversionUpdateCallback != nil {
		var frontInversion *Inversion
		if front != nil {
			frontInversion, _ = front.(*Inversion)
		}
		stage.OnAfterInversionUpdateCallback.OnAfterUpdate(stage, inversion, frontInversion)
	}
}

func (inversion *Inversion) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInversionDeleteCallback != nil {
		var frontInversion *Inversion
		if front != nil {
			frontInversion, _ = front.(*Inversion)
		}
		stage.OnAfterInversionDeleteCallback.OnAfterDelete(stage, inversion, frontInversion)
	}
}

func (kind *Kind) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKindCreateCallback != nil {
		stage.OnAfterKindCreateCallback.OnAfterCreate(stage, kind)
	}
}

func (kind *Kind) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKindUpdateCallback != nil {
		var frontKind *Kind
		if front != nil {
			frontKind, _ = front.(*Kind)
		}
		stage.OnAfterKindUpdateCallback.OnAfterUpdate(stage, kind, frontKind)
	}
}

func (kind *Kind) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKindDeleteCallback != nil {
		var frontKind *Kind
		if front != nil {
			frontKind, _ = front.(*Kind)
		}
		stage.OnAfterKindDeleteCallback.OnAfterDelete(stage, kind, frontKind)
	}
}

func (listening *Listening) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterListeningCreateCallback != nil {
		stage.OnAfterListeningCreateCallback.OnAfterCreate(stage, listening)
	}
}

func (listening *Listening) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterListeningUpdateCallback != nil {
		var frontListening *Listening
		if front != nil {
			frontListening, _ = front.(*Listening)
		}
		stage.OnAfterListeningUpdateCallback.OnAfterUpdate(stage, listening, frontListening)
	}
}

func (listening *Listening) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterListeningDeleteCallback != nil {
		var frontListening *Listening
		if front != nil {
			frontListening, _ = front.(*Listening)
		}
		stage.OnAfterListeningDeleteCallback.OnAfterDelete(stage, listening, frontListening)
	}
}

func (measurenumbering *MeasureNumbering) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureNumberingCreateCallback != nil {
		stage.OnAfterMeasureNumberingCreateCallback.OnAfterCreate(stage, measurenumbering)
	}
}

func (measurenumbering *MeasureNumbering) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureNumberingUpdateCallback != nil {
		var frontMeasureNumbering *MeasureNumbering
		if front != nil {
			frontMeasureNumbering, _ = front.(*MeasureNumbering)
		}
		stage.OnAfterMeasureNumberingUpdateCallback.OnAfterUpdate(stage, measurenumbering, frontMeasureNumbering)
	}
}

func (measurenumbering *MeasureNumbering) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureNumberingDeleteCallback != nil {
		var frontMeasureNumbering *MeasureNumbering
		if front != nil {
			frontMeasureNumbering, _ = front.(*MeasureNumbering)
		}
		stage.OnAfterMeasureNumberingDeleteCallback.OnAfterDelete(stage, measurenumbering, frontMeasureNumbering)
	}
}

func (membrane *Membrane) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMembraneCreateCallback != nil {
		stage.OnAfterMembraneCreateCallback.OnAfterCreate(stage, membrane)
	}
}

func (membrane *Membrane) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMembraneUpdateCallback != nil {
		var frontMembrane *Membrane
		if front != nil {
			frontMembrane, _ = front.(*Membrane)
		}
		stage.OnAfterMembraneUpdateCallback.OnAfterUpdate(stage, membrane, frontMembrane)
	}
}

func (membrane *Membrane) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMembraneDeleteCallback != nil {
		var frontMembrane *Membrane
		if front != nil {
			frontMembrane, _ = front.(*Membrane)
		}
		stage.OnAfterMembraneDeleteCallback.OnAfterDelete(stage, membrane, frontMembrane)
	}
}

func (metal *Metal) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetalCreateCallback != nil {
		stage.OnAfterMetalCreateCallback.OnAfterCreate(stage, metal)
	}
}

func (metal *Metal) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetalUpdateCallback != nil {
		var frontMetal *Metal
		if front != nil {
			frontMetal, _ = front.(*Metal)
		}
		stage.OnAfterMetalUpdateCallback.OnAfterUpdate(stage, metal, frontMetal)
	}
}

func (metal *Metal) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetalDeleteCallback != nil {
		var frontMetal *Metal
		if front != nil {
			frontMetal, _ = front.(*Metal)
		}
		stage.OnAfterMetalDeleteCallback.OnAfterDelete(stage, metal, frontMetal)
	}
}

func (metronome *Metronome) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetronomeCreateCallback != nil {
		stage.OnAfterMetronomeCreateCallback.OnAfterCreate(stage, metronome)
	}
}

func (metronome *Metronome) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeUpdateCallback != nil {
		var frontMetronome *Metronome
		if front != nil {
			frontMetronome, _ = front.(*Metronome)
		}
		stage.OnAfterMetronomeUpdateCallback.OnAfterUpdate(stage, metronome, frontMetronome)
	}
}

func (metronome *Metronome) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeDeleteCallback != nil {
		var frontMetronome *Metronome
		if front != nil {
			frontMetronome, _ = front.(*Metronome)
		}
		stage.OnAfterMetronomeDeleteCallback.OnAfterDelete(stage, metronome, frontMetronome)
	}
}

func (metronomebeam *MetronomeBeam) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetronomeBeamCreateCallback != nil {
		stage.OnAfterMetronomeBeamCreateCallback.OnAfterCreate(stage, metronomebeam)
	}
}

func (metronomebeam *MetronomeBeam) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeBeamUpdateCallback != nil {
		var frontMetronomeBeam *MetronomeBeam
		if front != nil {
			frontMetronomeBeam, _ = front.(*MetronomeBeam)
		}
		stage.OnAfterMetronomeBeamUpdateCallback.OnAfterUpdate(stage, metronomebeam, frontMetronomeBeam)
	}
}

func (metronomebeam *MetronomeBeam) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeBeamDeleteCallback != nil {
		var frontMetronomeBeam *MetronomeBeam
		if front != nil {
			frontMetronomeBeam, _ = front.(*MetronomeBeam)
		}
		stage.OnAfterMetronomeBeamDeleteCallback.OnAfterDelete(stage, metronomebeam, frontMetronomeBeam)
	}
}

func (metronomenote *MetronomeNote) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetronomeNoteCreateCallback != nil {
		stage.OnAfterMetronomeNoteCreateCallback.OnAfterCreate(stage, metronomenote)
	}
}

func (metronomenote *MetronomeNote) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeNoteUpdateCallback != nil {
		var frontMetronomeNote *MetronomeNote
		if front != nil {
			frontMetronomeNote, _ = front.(*MetronomeNote)
		}
		stage.OnAfterMetronomeNoteUpdateCallback.OnAfterUpdate(stage, metronomenote, frontMetronomeNote)
	}
}

func (metronomenote *MetronomeNote) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeNoteDeleteCallback != nil {
		var frontMetronomeNote *MetronomeNote
		if front != nil {
			frontMetronomeNote, _ = front.(*MetronomeNote)
		}
		stage.OnAfterMetronomeNoteDeleteCallback.OnAfterDelete(stage, metronomenote, frontMetronomeNote)
	}
}

func (metronometied *MetronomeTied) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetronomeTiedCreateCallback != nil {
		stage.OnAfterMetronomeTiedCreateCallback.OnAfterCreate(stage, metronometied)
	}
}

func (metronometied *MetronomeTied) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeTiedUpdateCallback != nil {
		var frontMetronomeTied *MetronomeTied
		if front != nil {
			frontMetronomeTied, _ = front.(*MetronomeTied)
		}
		stage.OnAfterMetronomeTiedUpdateCallback.OnAfterUpdate(stage, metronometied, frontMetronomeTied)
	}
}

func (metronometied *MetronomeTied) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeTiedDeleteCallback != nil {
		var frontMetronomeTied *MetronomeTied
		if front != nil {
			frontMetronomeTied, _ = front.(*MetronomeTied)
		}
		stage.OnAfterMetronomeTiedDeleteCallback.OnAfterDelete(stage, metronometied, frontMetronomeTied)
	}
}

func (metronometuplet *MetronomeTuplet) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetronomeTupletCreateCallback != nil {
		stage.OnAfterMetronomeTupletCreateCallback.OnAfterCreate(stage, metronometuplet)
	}
}

func (metronometuplet *MetronomeTuplet) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeTupletUpdateCallback != nil {
		var frontMetronomeTuplet *MetronomeTuplet
		if front != nil {
			frontMetronomeTuplet, _ = front.(*MetronomeTuplet)
		}
		stage.OnAfterMetronomeTupletUpdateCallback.OnAfterUpdate(stage, metronometuplet, frontMetronomeTuplet)
	}
}

func (metronometuplet *MetronomeTuplet) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronomeTupletDeleteCallback != nil {
		var frontMetronomeTuplet *MetronomeTuplet
		if front != nil {
			frontMetronomeTuplet, _ = front.(*MetronomeTuplet)
		}
		stage.OnAfterMetronomeTupletDeleteCallback.OnAfterDelete(stage, metronometuplet, frontMetronomeTuplet)
	}
}

func (numeral *Numeral) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNumeralCreateCallback != nil {
		stage.OnAfterNumeralCreateCallback.OnAfterCreate(stage, numeral)
	}
}

func (numeral *Numeral) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeralUpdateCallback != nil {
		var frontNumeral *Numeral
		if front != nil {
			frontNumeral, _ = front.(*Numeral)
		}
		stage.OnAfterNumeralUpdateCallback.OnAfterUpdate(stage, numeral, frontNumeral)
	}
}

func (numeral *Numeral) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeralDeleteCallback != nil {
		var frontNumeral *Numeral
		if front != nil {
			frontNumeral, _ = front.(*Numeral)
		}
		stage.OnAfterNumeralDeleteCallback.OnAfterDelete(stage, numeral, frontNumeral)
	}
}

func (numeralkey *NumeralKey) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNumeralKeyCreateCallback != nil {
		stage.OnAfterNumeralKeyCreateCallback.OnAfterCreate(stage, numeralkey)
	}
}

func (numeralkey *NumeralKey) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeralKeyUpdateCallback != nil {
		var frontNumeralKey *NumeralKey
		if front != nil {
			frontNumeralKey, _ = front.(*NumeralKey)
		}
		stage.OnAfterNumeralKeyUpdateCallback.OnAfterUpdate(stage, numeralkey, frontNumeralKey)
	}
}

func (numeralkey *NumeralKey) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeralKeyDeleteCallback != nil {
		var frontNumeralKey *NumeralKey
		if front != nil {
			frontNumeralKey, _ = front.(*NumeralKey)
		}
		stage.OnAfterNumeralKeyDeleteCallback.OnAfterDelete(stage, numeralkey, frontNumeralKey)
	}
}

func (numeralroot *NumeralRoot) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNumeralRootCreateCallback != nil {
		stage.OnAfterNumeralRootCreateCallback.OnAfterCreate(stage, numeralroot)
	}
}

func (numeralroot *NumeralRoot) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeralRootUpdateCallback != nil {
		var frontNumeralRoot *NumeralRoot
		if front != nil {
			frontNumeralRoot, _ = front.(*NumeralRoot)
		}
		stage.OnAfterNumeralRootUpdateCallback.OnAfterUpdate(stage, numeralroot, frontNumeralRoot)
	}
}

func (numeralroot *NumeralRoot) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeralRootDeleteCallback != nil {
		var frontNumeralRoot *NumeralRoot
		if front != nil {
			frontNumeralRoot, _ = front.(*NumeralRoot)
		}
		stage.OnAfterNumeralRootDeleteCallback.OnAfterDelete(stage, numeralroot, frontNumeralRoot)
	}
}

func (octaveshift *OctaveShift) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOctaveShiftCreateCallback != nil {
		stage.OnAfterOctaveShiftCreateCallback.OnAfterCreate(stage, octaveshift)
	}
}

func (octaveshift *OctaveShift) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOctaveShiftUpdateCallback != nil {
		var frontOctaveShift *OctaveShift
		if front != nil {
			frontOctaveShift, _ = front.(*OctaveShift)
		}
		stage.OnAfterOctaveShiftUpdateCallback.OnAfterUpdate(stage, octaveshift, frontOctaveShift)
	}
}

func (octaveshift *OctaveShift) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOctaveShiftDeleteCallback != nil {
		var frontOctaveShift *OctaveShift
		if front != nil {
			frontOctaveShift, _ = front.(*OctaveShift)
		}
		stage.OnAfterOctaveShiftDeleteCallback.OnAfterDelete(stage, octaveshift, frontOctaveShift)
	}
}

func (offset *Offset) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOffsetCreateCallback != nil {
		stage.OnAfterOffsetCreateCallback.OnAfterCreate(stage, offset)
	}
}

func (offset *Offset) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOffsetUpdateCallback != nil {
		var frontOffset *Offset
		if front != nil {
			frontOffset, _ = front.(*Offset)
		}
		stage.OnAfterOffsetUpdateCallback.OnAfterUpdate(stage, offset, frontOffset)
	}
}

func (offset *Offset) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOffsetDeleteCallback != nil {
		var frontOffset *Offset
		if front != nil {
			frontOffset, _ = front.(*Offset)
		}
		stage.OnAfterOffsetDeleteCallback.OnAfterDelete(stage, offset, frontOffset)
	}
}

func (otherdirection *OtherDirection) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOtherDirectionCreateCallback != nil {
		stage.OnAfterOtherDirectionCreateCallback.OnAfterCreate(stage, otherdirection)
	}
}

func (otherdirection *OtherDirection) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherDirectionUpdateCallback != nil {
		var frontOtherDirection *OtherDirection
		if front != nil {
			frontOtherDirection, _ = front.(*OtherDirection)
		}
		stage.OnAfterOtherDirectionUpdateCallback.OnAfterUpdate(stage, otherdirection, frontOtherDirection)
	}
}

func (otherdirection *OtherDirection) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherDirectionDeleteCallback != nil {
		var frontOtherDirection *OtherDirection
		if front != nil {
			frontOtherDirection, _ = front.(*OtherDirection)
		}
		stage.OnAfterOtherDirectionDeleteCallback.OnAfterDelete(stage, otherdirection, frontOtherDirection)
	}
}

func (otherlistening *OtherListening) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOtherListeningCreateCallback != nil {
		stage.OnAfterOtherListeningCreateCallback.OnAfterCreate(stage, otherlistening)
	}
}

func (otherlistening *OtherListening) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherListeningUpdateCallback != nil {
		var frontOtherListening *OtherListening
		if front != nil {
			frontOtherListening, _ = front.(*OtherListening)
		}
		stage.OnAfterOtherListeningUpdateCallback.OnAfterUpdate(stage, otherlistening, frontOtherListening)
	}
}

func (otherlistening *OtherListening) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherListeningDeleteCallback != nil {
		var frontOtherListening *OtherListening
		if front != nil {
			frontOtherListening, _ = front.(*OtherListening)
		}
		stage.OnAfterOtherListeningDeleteCallback.OnAfterDelete(stage, otherlistening, frontOtherListening)
	}
}

func (pedal *Pedal) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPedalCreateCallback != nil {
		stage.OnAfterPedalCreateCallback.OnAfterCreate(stage, pedal)
	}
}

func (pedal *Pedal) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPedalUpdateCallback != nil {
		var frontPedal *Pedal
		if front != nil {
			frontPedal, _ = front.(*Pedal)
		}
		stage.OnAfterPedalUpdateCallback.OnAfterUpdate(stage, pedal, frontPedal)
	}
}

func (pedal *Pedal) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPedalDeleteCallback != nil {
		var frontPedal *Pedal
		if front != nil {
			frontPedal, _ = front.(*Pedal)
		}
		stage.OnAfterPedalDeleteCallback.OnAfterDelete(stage, pedal, frontPedal)
	}
}

func (pedaltuning *PedalTuning) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPedalTuningCreateCallback != nil {
		stage.OnAfterPedalTuningCreateCallback.OnAfterCreate(stage, pedaltuning)
	}
}

func (pedaltuning *PedalTuning) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPedalTuningUpdateCallback != nil {
		var frontPedalTuning *PedalTuning
		if front != nil {
			frontPedalTuning, _ = front.(*PedalTuning)
		}
		stage.OnAfterPedalTuningUpdateCallback.OnAfterUpdate(stage, pedaltuning, frontPedalTuning)
	}
}

func (pedaltuning *PedalTuning) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPedalTuningDeleteCallback != nil {
		var frontPedalTuning *PedalTuning
		if front != nil {
			frontPedalTuning, _ = front.(*PedalTuning)
		}
		stage.OnAfterPedalTuningDeleteCallback.OnAfterDelete(stage, pedaltuning, frontPedalTuning)
	}
}

func (perminute *PerMinute) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPerMinuteCreateCallback != nil {
		stage.OnAfterPerMinuteCreateCallback.OnAfterCreate(stage, perminute)
	}
}

func (perminute *PerMinute) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerMinuteUpdateCallback != nil {
		var frontPerMinute *PerMinute
		if front != nil {
			frontPerMinute, _ = front.(*PerMinute)
		}
		stage.OnAfterPerMinuteUpdateCallback.OnAfterUpdate(stage, perminute, frontPerMinute)
	}
}

func (perminute *PerMinute) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerMinuteDeleteCallback != nil {
		var frontPerMinute *PerMinute
		if front != nil {
			frontPerMinute, _ = front.(*PerMinute)
		}
		stage.OnAfterPerMinuteDeleteCallback.OnAfterDelete(stage, perminute, frontPerMinute)
	}
}

func (percussion *Percussion) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPercussionCreateCallback != nil {
		stage.OnAfterPercussionCreateCallback.OnAfterCreate(stage, percussion)
	}
}

func (percussion *Percussion) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPercussionUpdateCallback != nil {
		var frontPercussion *Percussion
		if front != nil {
			frontPercussion, _ = front.(*Percussion)
		}
		stage.OnAfterPercussionUpdateCallback.OnAfterUpdate(stage, percussion, frontPercussion)
	}
}

func (percussion *Percussion) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPercussionDeleteCallback != nil {
		var frontPercussion *Percussion
		if front != nil {
			frontPercussion, _ = front.(*Percussion)
		}
		stage.OnAfterPercussionDeleteCallback.OnAfterDelete(stage, percussion, frontPercussion)
	}
}

func (pitched *Pitched) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPitchedCreateCallback != nil {
		stage.OnAfterPitchedCreateCallback.OnAfterCreate(stage, pitched)
	}
}

func (pitched *Pitched) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPitchedUpdateCallback != nil {
		var frontPitched *Pitched
		if front != nil {
			frontPitched, _ = front.(*Pitched)
		}
		stage.OnAfterPitchedUpdateCallback.OnAfterUpdate(stage, pitched, frontPitched)
	}
}

func (pitched *Pitched) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPitchedDeleteCallback != nil {
		var frontPitched *Pitched
		if front != nil {
			frontPitched, _ = front.(*Pitched)
		}
		stage.OnAfterPitchedDeleteCallback.OnAfterDelete(stage, pitched, frontPitched)
	}
}

func (principalvoice *PrincipalVoice) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrincipalVoiceCreateCallback != nil {
		stage.OnAfterPrincipalVoiceCreateCallback.OnAfterCreate(stage, principalvoice)
	}
}

func (principalvoice *PrincipalVoice) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrincipalVoiceUpdateCallback != nil {
		var frontPrincipalVoice *PrincipalVoice
		if front != nil {
			frontPrincipalVoice, _ = front.(*PrincipalVoice)
		}
		stage.OnAfterPrincipalVoiceUpdateCallback.OnAfterUpdate(stage, principalvoice, frontPrincipalVoice)
	}
}

func (principalvoice *PrincipalVoice) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrincipalVoiceDeleteCallback != nil {
		var frontPrincipalVoice *PrincipalVoice
		if front != nil {
			frontPrincipalVoice, _ = front.(*PrincipalVoice)
		}
		stage.OnAfterPrincipalVoiceDeleteCallback.OnAfterDelete(stage, principalvoice, frontPrincipalVoice)
	}
}

func (print *Print) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrintCreateCallback != nil {
		stage.OnAfterPrintCreateCallback.OnAfterCreate(stage, print)
	}
}

func (print *Print) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrintUpdateCallback != nil {
		var frontPrint *Print
		if front != nil {
			frontPrint, _ = front.(*Print)
		}
		stage.OnAfterPrintUpdateCallback.OnAfterUpdate(stage, print, frontPrint)
	}
}

func (print *Print) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrintDeleteCallback != nil {
		var frontPrint *Print
		if front != nil {
			frontPrint, _ = front.(*Print)
		}
		stage.OnAfterPrintDeleteCallback.OnAfterDelete(stage, print, frontPrint)
	}
}

func (root *Root) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRootCreateCallback != nil {
		stage.OnAfterRootCreateCallback.OnAfterCreate(stage, root)
	}
}

func (root *Root) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRootUpdateCallback != nil {
		var frontRoot *Root
		if front != nil {
			frontRoot, _ = front.(*Root)
		}
		stage.OnAfterRootUpdateCallback.OnAfterUpdate(stage, root, frontRoot)
	}
}

func (root *Root) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRootDeleteCallback != nil {
		var frontRoot *Root
		if front != nil {
			frontRoot, _ = front.(*Root)
		}
		stage.OnAfterRootDeleteCallback.OnAfterDelete(stage, root, frontRoot)
	}
}

func (rootstep *RootStep) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRootStepCreateCallback != nil {
		stage.OnAfterRootStepCreateCallback.OnAfterCreate(stage, rootstep)
	}
}

func (rootstep *RootStep) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRootStepUpdateCallback != nil {
		var frontRootStep *RootStep
		if front != nil {
			frontRootStep, _ = front.(*RootStep)
		}
		stage.OnAfterRootStepUpdateCallback.OnAfterUpdate(stage, rootstep, frontRootStep)
	}
}

func (rootstep *RootStep) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRootStepDeleteCallback != nil {
		var frontRootStep *RootStep
		if front != nil {
			frontRootStep, _ = front.(*RootStep)
		}
		stage.OnAfterRootStepDeleteCallback.OnAfterDelete(stage, rootstep, frontRootStep)
	}
}

func (scordatura *Scordatura) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScordaturaCreateCallback != nil {
		stage.OnAfterScordaturaCreateCallback.OnAfterCreate(stage, scordatura)
	}
}

func (scordatura *Scordatura) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScordaturaUpdateCallback != nil {
		var frontScordatura *Scordatura
		if front != nil {
			frontScordatura, _ = front.(*Scordatura)
		}
		stage.OnAfterScordaturaUpdateCallback.OnAfterUpdate(stage, scordatura, frontScordatura)
	}
}

func (scordatura *Scordatura) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScordaturaDeleteCallback != nil {
		var frontScordatura *Scordatura
		if front != nil {
			frontScordatura, _ = front.(*Scordatura)
		}
		stage.OnAfterScordaturaDeleteCallback.OnAfterDelete(stage, scordatura, frontScordatura)
	}
}

func (sound *Sound) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSoundCreateCallback != nil {
		stage.OnAfterSoundCreateCallback.OnAfterCreate(stage, sound)
	}
}

func (sound *Sound) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSoundUpdateCallback != nil {
		var frontSound *Sound
		if front != nil {
			frontSound, _ = front.(*Sound)
		}
		stage.OnAfterSoundUpdateCallback.OnAfterUpdate(stage, sound, frontSound)
	}
}

func (sound *Sound) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSoundDeleteCallback != nil {
		var frontSound *Sound
		if front != nil {
			frontSound, _ = front.(*Sound)
		}
		stage.OnAfterSoundDeleteCallback.OnAfterDelete(stage, sound, frontSound)
	}
}

func (staffdivide *StaffDivide) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaffDivideCreateCallback != nil {
		stage.OnAfterStaffDivideCreateCallback.OnAfterCreate(stage, staffdivide)
	}
}

func (staffdivide *StaffDivide) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffDivideUpdateCallback != nil {
		var frontStaffDivide *StaffDivide
		if front != nil {
			frontStaffDivide, _ = front.(*StaffDivide)
		}
		stage.OnAfterStaffDivideUpdateCallback.OnAfterUpdate(stage, staffdivide, frontStaffDivide)
	}
}

func (staffdivide *StaffDivide) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaffDivideDeleteCallback != nil {
		var frontStaffDivide *StaffDivide
		if front != nil {
			frontStaffDivide, _ = front.(*StaffDivide)
		}
		stage.OnAfterStaffDivideDeleteCallback.OnAfterDelete(stage, staffdivide, frontStaffDivide)
	}
}

func (stick *Stick) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStickCreateCallback != nil {
		stage.OnAfterStickCreateCallback.OnAfterCreate(stage, stick)
	}
}

func (stick *Stick) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStickUpdateCallback != nil {
		var frontStick *Stick
		if front != nil {
			frontStick, _ = front.(*Stick)
		}
		stage.OnAfterStickUpdateCallback.OnAfterUpdate(stage, stick, frontStick)
	}
}

func (stick *Stick) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStickDeleteCallback != nil {
		var frontStick *Stick
		if front != nil {
			frontStick, _ = front.(*Stick)
		}
		stage.OnAfterStickDeleteCallback.OnAfterDelete(stage, stick, frontStick)
	}
}

func (stringmute *StringMute) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStringMuteCreateCallback != nil {
		stage.OnAfterStringMuteCreateCallback.OnAfterCreate(stage, stringmute)
	}
}

func (stringmute *StringMute) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStringMuteUpdateCallback != nil {
		var frontStringMute *StringMute
		if front != nil {
			frontStringMute, _ = front.(*StringMute)
		}
		stage.OnAfterStringMuteUpdateCallback.OnAfterUpdate(stage, stringmute, frontStringMute)
	}
}

func (stringmute *StringMute) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStringMuteDeleteCallback != nil {
		var frontStringMute *StringMute
		if front != nil {
			frontStringMute, _ = front.(*StringMute)
		}
		stage.OnAfterStringMuteDeleteCallback.OnAfterDelete(stage, stringmute, frontStringMute)
	}
}

func (styletext *StyleText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStyleTextCreateCallback != nil {
		stage.OnAfterStyleTextCreateCallback.OnAfterCreate(stage, styletext)
	}
}

func (styletext *StyleText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStyleTextUpdateCallback != nil {
		var frontStyleText *StyleText
		if front != nil {
			frontStyleText, _ = front.(*StyleText)
		}
		stage.OnAfterStyleTextUpdateCallback.OnAfterUpdate(stage, styletext, frontStyleText)
	}
}

func (styletext *StyleText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStyleTextDeleteCallback != nil {
		var frontStyleText *StyleText
		if front != nil {
			frontStyleText, _ = front.(*StyleText)
		}
		stage.OnAfterStyleTextDeleteCallback.OnAfterDelete(stage, styletext, frontStyleText)
	}
}

func (swing *Swing) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSwingCreateCallback != nil {
		stage.OnAfterSwingCreateCallback.OnAfterCreate(stage, swing)
	}
}

func (swing *Swing) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSwingUpdateCallback != nil {
		var frontSwing *Swing
		if front != nil {
			frontSwing, _ = front.(*Swing)
		}
		stage.OnAfterSwingUpdateCallback.OnAfterUpdate(stage, swing, frontSwing)
	}
}

func (swing *Swing) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSwingDeleteCallback != nil {
		var frontSwing *Swing
		if front != nil {
			frontSwing, _ = front.(*Swing)
		}
		stage.OnAfterSwingDeleteCallback.OnAfterDelete(stage, swing, frontSwing)
	}
}

func (sync *Sync) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSyncCreateCallback != nil {
		stage.OnAfterSyncCreateCallback.OnAfterCreate(stage, sync)
	}
}

func (sync *Sync) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSyncUpdateCallback != nil {
		var frontSync *Sync
		if front != nil {
			frontSync, _ = front.(*Sync)
		}
		stage.OnAfterSyncUpdateCallback.OnAfterUpdate(stage, sync, frontSync)
	}
}

func (sync *Sync) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSyncDeleteCallback != nil {
		var frontSync *Sync
		if front != nil {
			frontSync, _ = front.(*Sync)
		}
		stage.OnAfterSyncDeleteCallback.OnAfterDelete(stage, sync, frontSync)
	}
}

func (timemodification *TimeModification) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeModificationCreateCallback != nil {
		stage.OnAfterTimeModificationCreateCallback.OnAfterCreate(stage, timemodification)
	}
}

func (timemodification *TimeModification) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeModificationUpdateCallback != nil {
		var frontTimeModification *TimeModification
		if front != nil {
			frontTimeModification, _ = front.(*TimeModification)
		}
		stage.OnAfterTimeModificationUpdateCallback.OnAfterUpdate(stage, timemodification, frontTimeModification)
	}
}

func (timemodification *TimeModification) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeModificationDeleteCallback != nil {
		var frontTimeModification *TimeModification
		if front != nil {
			frontTimeModification, _ = front.(*TimeModification)
		}
		stage.OnAfterTimeModificationDeleteCallback.OnAfterDelete(stage, timemodification, frontTimeModification)
	}
}

func (timpani *Timpani) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimpaniCreateCallback != nil {
		stage.OnAfterTimpaniCreateCallback.OnAfterCreate(stage, timpani)
	}
}

func (timpani *Timpani) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimpaniUpdateCallback != nil {
		var frontTimpani *Timpani
		if front != nil {
			frontTimpani, _ = front.(*Timpani)
		}
		stage.OnAfterTimpaniUpdateCallback.OnAfterUpdate(stage, timpani, frontTimpani)
	}
}

func (timpani *Timpani) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimpaniDeleteCallback != nil {
		var frontTimpani *Timpani
		if front != nil {
			frontTimpani, _ = front.(*Timpani)
		}
		stage.OnAfterTimpaniDeleteCallback.OnAfterDelete(stage, timpani, frontTimpani)
	}
}

func (virtualinstrument *VirtualInstrument) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVirtualInstrumentCreateCallback != nil {
		stage.OnAfterVirtualInstrumentCreateCallback.OnAfterCreate(stage, virtualinstrument)
	}
}

func (virtualinstrument *VirtualInstrument) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVirtualInstrumentUpdateCallback != nil {
		var frontVirtualInstrument *VirtualInstrument
		if front != nil {
			frontVirtualInstrument, _ = front.(*VirtualInstrument)
		}
		stage.OnAfterVirtualInstrumentUpdateCallback.OnAfterUpdate(stage, virtualinstrument, frontVirtualInstrument)
	}
}

func (virtualinstrument *VirtualInstrument) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVirtualInstrumentDeleteCallback != nil {
		var frontVirtualInstrument *VirtualInstrument
		if front != nil {
			frontVirtualInstrument, _ = front.(*VirtualInstrument)
		}
		stage.OnAfterVirtualInstrumentDeleteCallback.OnAfterDelete(stage, virtualinstrument, frontVirtualInstrument)
	}
}

func (wedge *Wedge) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWedgeCreateCallback != nil {
		stage.OnAfterWedgeCreateCallback.OnAfterCreate(stage, wedge)
	}
}

func (wedge *Wedge) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWedgeUpdateCallback != nil {
		var frontWedge *Wedge
		if front != nil {
			frontWedge, _ = front.(*Wedge)
		}
		stage.OnAfterWedgeUpdateCallback.OnAfterUpdate(stage, wedge, frontWedge)
	}
}

func (wedge *Wedge) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWedgeDeleteCallback != nil {
		var frontWedge *Wedge
		if front != nil {
			frontWedge, _ = front.(*Wedge)
		}
		stage.OnAfterWedgeDeleteCallback.OnAfterDelete(stage, wedge, frontWedge)
	}
}

func (wood *Wood) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWoodCreateCallback != nil {
		stage.OnAfterWoodCreateCallback.OnAfterCreate(stage, wood)
	}
}

func (wood *Wood) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWoodUpdateCallback != nil {
		var frontWood *Wood
		if front != nil {
			frontWood, _ = front.(*Wood)
		}
		stage.OnAfterWoodUpdateCallback.OnAfterUpdate(stage, wood, frontWood)
	}
}

func (wood *Wood) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWoodDeleteCallback != nil {
		var frontWood *Wood
		if front != nil {
			frontWood, _ = front.(*Wood)
		}
		stage.OnAfterWoodDeleteCallback.OnAfterDelete(stage, wood, frontWood)
	}
}

