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
func (a_directive *A_directive) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_directiveCreateCallback != nil {
		stage.OnAfterA_directiveCreateCallback.OnAfterCreate(stage, a_directive)
	}
}

func (a_directive *A_directive) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_directiveUpdateCallback != nil {
		var frontA_directive *A_directive
		if front != nil {
			frontA_directive, _ = front.(*A_directive)
		}
		stage.OnAfterA_directiveUpdateCallback.OnAfterUpdate(stage, a_directive, frontA_directive)
	}
}

func (a_directive *A_directive) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_directiveDeleteCallback != nil {
		var frontA_directive *A_directive
		if front != nil {
			frontA_directive, _ = front.(*A_directive)
		}
		stage.OnAfterA_directiveDeleteCallback.OnAfterDelete(stage, a_directive, frontA_directive)
	}
}

func (a_measure *A_measure) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_measureCreateCallback != nil {
		stage.OnAfterA_measureCreateCallback.OnAfterCreate(stage, a_measure)
	}
}

func (a_measure *A_measure) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_measureUpdateCallback != nil {
		var frontA_measure *A_measure
		if front != nil {
			frontA_measure, _ = front.(*A_measure)
		}
		stage.OnAfterA_measureUpdateCallback.OnAfterUpdate(stage, a_measure, frontA_measure)
	}
}

func (a_measure *A_measure) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_measureDeleteCallback != nil {
		var frontA_measure *A_measure
		if front != nil {
			frontA_measure, _ = front.(*A_measure)
		}
		stage.OnAfterA_measureDeleteCallback.OnAfterDelete(stage, a_measure, frontA_measure)
	}
}

func (a_measure_1 *A_measure_1) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_measure_1CreateCallback != nil {
		stage.OnAfterA_measure_1CreateCallback.OnAfterCreate(stage, a_measure_1)
	}
}

func (a_measure_1 *A_measure_1) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_measure_1UpdateCallback != nil {
		var frontA_measure_1 *A_measure_1
		if front != nil {
			frontA_measure_1, _ = front.(*A_measure_1)
		}
		stage.OnAfterA_measure_1UpdateCallback.OnAfterUpdate(stage, a_measure_1, frontA_measure_1)
	}
}

func (a_measure_1 *A_measure_1) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_measure_1DeleteCallback != nil {
		var frontA_measure_1 *A_measure_1
		if front != nil {
			frontA_measure_1, _ = front.(*A_measure_1)
		}
		stage.OnAfterA_measure_1DeleteCallback.OnAfterDelete(stage, a_measure_1, frontA_measure_1)
	}
}

func (a_part *A_part) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_partCreateCallback != nil {
		stage.OnAfterA_partCreateCallback.OnAfterCreate(stage, a_part)
	}
}

func (a_part *A_part) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_partUpdateCallback != nil {
		var frontA_part *A_part
		if front != nil {
			frontA_part, _ = front.(*A_part)
		}
		stage.OnAfterA_partUpdateCallback.OnAfterUpdate(stage, a_part, frontA_part)
	}
}

func (a_part *A_part) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_partDeleteCallback != nil {
		var frontA_part *A_part
		if front != nil {
			frontA_part, _ = front.(*A_part)
		}
		stage.OnAfterA_partDeleteCallback.OnAfterDelete(stage, a_part, frontA_part)
	}
}

func (a_part_1 *A_part_1) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_part_1CreateCallback != nil {
		stage.OnAfterA_part_1CreateCallback.OnAfterCreate(stage, a_part_1)
	}
}

func (a_part_1 *A_part_1) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_part_1UpdateCallback != nil {
		var frontA_part_1 *A_part_1
		if front != nil {
			frontA_part_1, _ = front.(*A_part_1)
		}
		stage.OnAfterA_part_1UpdateCallback.OnAfterUpdate(stage, a_part_1, frontA_part_1)
	}
}

func (a_part_1 *A_part_1) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_part_1DeleteCallback != nil {
		var frontA_part_1 *A_part_1
		if front != nil {
			frontA_part_1, _ = front.(*A_part_1)
		}
		stage.OnAfterA_part_1DeleteCallback.OnAfterDelete(stage, a_part_1, frontA_part_1)
	}
}

func (accidental *Accidental) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAccidentalCreateCallback != nil {
		stage.OnAfterAccidentalCreateCallback.OnAfterCreate(stage, accidental)
	}
}

func (accidental *Accidental) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidentalUpdateCallback != nil {
		var frontAccidental *Accidental
		if front != nil {
			frontAccidental, _ = front.(*Accidental)
		}
		stage.OnAfterAccidentalUpdateCallback.OnAfterUpdate(stage, accidental, frontAccidental)
	}
}

func (accidental *Accidental) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidentalDeleteCallback != nil {
		var frontAccidental *Accidental
		if front != nil {
			frontAccidental, _ = front.(*Accidental)
		}
		stage.OnAfterAccidentalDeleteCallback.OnAfterDelete(stage, accidental, frontAccidental)
	}
}

func (accidental_mark *Accidental_mark) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAccidental_markCreateCallback != nil {
		stage.OnAfterAccidental_markCreateCallback.OnAfterCreate(stage, accidental_mark)
	}
}

func (accidental_mark *Accidental_mark) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidental_markUpdateCallback != nil {
		var frontAccidental_mark *Accidental_mark
		if front != nil {
			frontAccidental_mark, _ = front.(*Accidental_mark)
		}
		stage.OnAfterAccidental_markUpdateCallback.OnAfterUpdate(stage, accidental_mark, frontAccidental_mark)
	}
}

func (accidental_mark *Accidental_mark) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidental_markDeleteCallback != nil {
		var frontAccidental_mark *Accidental_mark
		if front != nil {
			frontAccidental_mark, _ = front.(*Accidental_mark)
		}
		stage.OnAfterAccidental_markDeleteCallback.OnAfterDelete(stage, accidental_mark, frontAccidental_mark)
	}
}

func (accidental_text *Accidental_text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAccidental_textCreateCallback != nil {
		stage.OnAfterAccidental_textCreateCallback.OnAfterCreate(stage, accidental_text)
	}
}

func (accidental_text *Accidental_text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidental_textUpdateCallback != nil {
		var frontAccidental_text *Accidental_text
		if front != nil {
			frontAccidental_text, _ = front.(*Accidental_text)
		}
		stage.OnAfterAccidental_textUpdateCallback.OnAfterUpdate(stage, accidental_text, frontAccidental_text)
	}
}

func (accidental_text *Accidental_text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidental_textDeleteCallback != nil {
		var frontAccidental_text *Accidental_text
		if front != nil {
			frontAccidental_text, _ = front.(*Accidental_text)
		}
		stage.OnAfterAccidental_textDeleteCallback.OnAfterDelete(stage, accidental_text, frontAccidental_text)
	}
}

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

func (accordion_registration *Accordion_registration) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAccordion_registrationCreateCallback != nil {
		stage.OnAfterAccordion_registrationCreateCallback.OnAfterCreate(stage, accordion_registration)
	}
}

func (accordion_registration *Accordion_registration) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccordion_registrationUpdateCallback != nil {
		var frontAccordion_registration *Accordion_registration
		if front != nil {
			frontAccordion_registration, _ = front.(*Accordion_registration)
		}
		stage.OnAfterAccordion_registrationUpdateCallback.OnAfterUpdate(stage, accordion_registration, frontAccordion_registration)
	}
}

func (accordion_registration *Accordion_registration) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccordion_registrationDeleteCallback != nil {
		var frontAccordion_registration *Accordion_registration
		if front != nil {
			frontAccordion_registration, _ = front.(*Accordion_registration)
		}
		stage.OnAfterAccordion_registrationDeleteCallback.OnAfterDelete(stage, accordion_registration, frontAccordion_registration)
	}
}

func (appearance *Appearance) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAppearanceCreateCallback != nil {
		stage.OnAfterAppearanceCreateCallback.OnAfterCreate(stage, appearance)
	}
}

func (appearance *Appearance) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAppearanceUpdateCallback != nil {
		var frontAppearance *Appearance
		if front != nil {
			frontAppearance, _ = front.(*Appearance)
		}
		stage.OnAfterAppearanceUpdateCallback.OnAfterUpdate(stage, appearance, frontAppearance)
	}
}

func (appearance *Appearance) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAppearanceDeleteCallback != nil {
		var frontAppearance *Appearance
		if front != nil {
			frontAppearance, _ = front.(*Appearance)
		}
		stage.OnAfterAppearanceDeleteCallback.OnAfterDelete(stage, appearance, frontAppearance)
	}
}

func (arpeggiate *Arpeggiate) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArpeggiateCreateCallback != nil {
		stage.OnAfterArpeggiateCreateCallback.OnAfterCreate(stage, arpeggiate)
	}
}

func (arpeggiate *Arpeggiate) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArpeggiateUpdateCallback != nil {
		var frontArpeggiate *Arpeggiate
		if front != nil {
			frontArpeggiate, _ = front.(*Arpeggiate)
		}
		stage.OnAfterArpeggiateUpdateCallback.OnAfterUpdate(stage, arpeggiate, frontArpeggiate)
	}
}

func (arpeggiate *Arpeggiate) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArpeggiateDeleteCallback != nil {
		var frontArpeggiate *Arpeggiate
		if front != nil {
			frontArpeggiate, _ = front.(*Arpeggiate)
		}
		stage.OnAfterArpeggiateDeleteCallback.OnAfterDelete(stage, arpeggiate, frontArpeggiate)
	}
}

func (arrow *Arrow) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrowCreateCallback != nil {
		stage.OnAfterArrowCreateCallback.OnAfterCreate(stage, arrow)
	}
}

func (arrow *Arrow) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrowUpdateCallback != nil {
		var frontArrow *Arrow
		if front != nil {
			frontArrow, _ = front.(*Arrow)
		}
		stage.OnAfterArrowUpdateCallback.OnAfterUpdate(stage, arrow, frontArrow)
	}
}

func (arrow *Arrow) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrowDeleteCallback != nil {
		var frontArrow *Arrow
		if front != nil {
			frontArrow, _ = front.(*Arrow)
		}
		stage.OnAfterArrowDeleteCallback.OnAfterDelete(stage, arrow, frontArrow)
	}
}

func (articulations *Articulations) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArticulationsCreateCallback != nil {
		stage.OnAfterArticulationsCreateCallback.OnAfterCreate(stage, articulations)
	}
}

func (articulations *Articulations) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArticulationsUpdateCallback != nil {
		var frontArticulations *Articulations
		if front != nil {
			frontArticulations, _ = front.(*Articulations)
		}
		stage.OnAfterArticulationsUpdateCallback.OnAfterUpdate(stage, articulations, frontArticulations)
	}
}

func (articulations *Articulations) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArticulationsDeleteCallback != nil {
		var frontArticulations *Articulations
		if front != nil {
			frontArticulations, _ = front.(*Articulations)
		}
		stage.OnAfterArticulationsDeleteCallback.OnAfterDelete(stage, articulations, frontArticulations)
	}
}

func (assess *Assess) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAssessCreateCallback != nil {
		stage.OnAfterAssessCreateCallback.OnAfterCreate(stage, assess)
	}
}

func (assess *Assess) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAssessUpdateCallback != nil {
		var frontAssess *Assess
		if front != nil {
			frontAssess, _ = front.(*Assess)
		}
		stage.OnAfterAssessUpdateCallback.OnAfterUpdate(stage, assess, frontAssess)
	}
}

func (assess *Assess) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAssessDeleteCallback != nil {
		var frontAssess *Assess
		if front != nil {
			frontAssess, _ = front.(*Assess)
		}
		stage.OnAfterAssessDeleteCallback.OnAfterDelete(stage, assess, frontAssess)
	}
}

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

func (backup *Backup) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBackupCreateCallback != nil {
		stage.OnAfterBackupCreateCallback.OnAfterCreate(stage, backup)
	}
}

func (backup *Backup) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBackupUpdateCallback != nil {
		var frontBackup *Backup
		if front != nil {
			frontBackup, _ = front.(*Backup)
		}
		stage.OnAfterBackupUpdateCallback.OnAfterUpdate(stage, backup, frontBackup)
	}
}

func (backup *Backup) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBackupDeleteCallback != nil {
		var frontBackup *Backup
		if front != nil {
			frontBackup, _ = front.(*Backup)
		}
		stage.OnAfterBackupDeleteCallback.OnAfterDelete(stage, backup, frontBackup)
	}
}

func (bar_style_color *Bar_style_color) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBar_style_colorCreateCallback != nil {
		stage.OnAfterBar_style_colorCreateCallback.OnAfterCreate(stage, bar_style_color)
	}
}

func (bar_style_color *Bar_style_color) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBar_style_colorUpdateCallback != nil {
		var frontBar_style_color *Bar_style_color
		if front != nil {
			frontBar_style_color, _ = front.(*Bar_style_color)
		}
		stage.OnAfterBar_style_colorUpdateCallback.OnAfterUpdate(stage, bar_style_color, frontBar_style_color)
	}
}

func (bar_style_color *Bar_style_color) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBar_style_colorDeleteCallback != nil {
		var frontBar_style_color *Bar_style_color
		if front != nil {
			frontBar_style_color, _ = front.(*Bar_style_color)
		}
		stage.OnAfterBar_style_colorDeleteCallback.OnAfterDelete(stage, bar_style_color, frontBar_style_color)
	}
}

func (barline *Barline) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBarlineCreateCallback != nil {
		stage.OnAfterBarlineCreateCallback.OnAfterCreate(stage, barline)
	}
}

func (barline *Barline) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarlineUpdateCallback != nil {
		var frontBarline *Barline
		if front != nil {
			frontBarline, _ = front.(*Barline)
		}
		stage.OnAfterBarlineUpdateCallback.OnAfterUpdate(stage, barline, frontBarline)
	}
}

func (barline *Barline) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarlineDeleteCallback != nil {
		var frontBarline *Barline
		if front != nil {
			frontBarline, _ = front.(*Barline)
		}
		stage.OnAfterBarlineDeleteCallback.OnAfterDelete(stage, barline, frontBarline)
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

func (bass_step *Bass_step) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBass_stepCreateCallback != nil {
		stage.OnAfterBass_stepCreateCallback.OnAfterCreate(stage, bass_step)
	}
}

func (bass_step *Bass_step) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBass_stepUpdateCallback != nil {
		var frontBass_step *Bass_step
		if front != nil {
			frontBass_step, _ = front.(*Bass_step)
		}
		stage.OnAfterBass_stepUpdateCallback.OnAfterUpdate(stage, bass_step, frontBass_step)
	}
}

func (bass_step *Bass_step) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBass_stepDeleteCallback != nil {
		var frontBass_step *Bass_step
		if front != nil {
			frontBass_step, _ = front.(*Bass_step)
		}
		stage.OnAfterBass_stepDeleteCallback.OnAfterDelete(stage, bass_step, frontBass_step)
	}
}

func (beam *Beam) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBeamCreateCallback != nil {
		stage.OnAfterBeamCreateCallback.OnAfterCreate(stage, beam)
	}
}

func (beam *Beam) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeamUpdateCallback != nil {
		var frontBeam *Beam
		if front != nil {
			frontBeam, _ = front.(*Beam)
		}
		stage.OnAfterBeamUpdateCallback.OnAfterUpdate(stage, beam, frontBeam)
	}
}

func (beam *Beam) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeamDeleteCallback != nil {
		var frontBeam *Beam
		if front != nil {
			frontBeam, _ = front.(*Beam)
		}
		stage.OnAfterBeamDeleteCallback.OnAfterDelete(stage, beam, frontBeam)
	}
}

func (beat_repeat *Beat_repeat) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBeat_repeatCreateCallback != nil {
		stage.OnAfterBeat_repeatCreateCallback.OnAfterCreate(stage, beat_repeat)
	}
}

func (beat_repeat *Beat_repeat) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeat_repeatUpdateCallback != nil {
		var frontBeat_repeat *Beat_repeat
		if front != nil {
			frontBeat_repeat, _ = front.(*Beat_repeat)
		}
		stage.OnAfterBeat_repeatUpdateCallback.OnAfterUpdate(stage, beat_repeat, frontBeat_repeat)
	}
}

func (beat_repeat *Beat_repeat) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeat_repeatDeleteCallback != nil {
		var frontBeat_repeat *Beat_repeat
		if front != nil {
			frontBeat_repeat, _ = front.(*Beat_repeat)
		}
		stage.OnAfterBeat_repeatDeleteCallback.OnAfterDelete(stage, beat_repeat, frontBeat_repeat)
	}
}

func (beat_unit_tied *Beat_unit_tied) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBeat_unit_tiedCreateCallback != nil {
		stage.OnAfterBeat_unit_tiedCreateCallback.OnAfterCreate(stage, beat_unit_tied)
	}
}

func (beat_unit_tied *Beat_unit_tied) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeat_unit_tiedUpdateCallback != nil {
		var frontBeat_unit_tied *Beat_unit_tied
		if front != nil {
			frontBeat_unit_tied, _ = front.(*Beat_unit_tied)
		}
		stage.OnAfterBeat_unit_tiedUpdateCallback.OnAfterUpdate(stage, beat_unit_tied, frontBeat_unit_tied)
	}
}

func (beat_unit_tied *Beat_unit_tied) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBeat_unit_tiedDeleteCallback != nil {
		var frontBeat_unit_tied *Beat_unit_tied
		if front != nil {
			frontBeat_unit_tied, _ = front.(*Beat_unit_tied)
		}
		stage.OnAfterBeat_unit_tiedDeleteCallback.OnAfterDelete(stage, beat_unit_tied, frontBeat_unit_tied)
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

func (bend *Bend) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBendCreateCallback != nil {
		stage.OnAfterBendCreateCallback.OnAfterCreate(stage, bend)
	}
}

func (bend *Bend) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBendUpdateCallback != nil {
		var frontBend *Bend
		if front != nil {
			frontBend, _ = front.(*Bend)
		}
		stage.OnAfterBendUpdateCallback.OnAfterUpdate(stage, bend, frontBend)
	}
}

func (bend *Bend) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBendDeleteCallback != nil {
		var frontBend *Bend
		if front != nil {
			frontBend, _ = front.(*Bend)
		}
		stage.OnAfterBendDeleteCallback.OnAfterDelete(stage, bend, frontBend)
	}
}

func (bookmark *Bookmark) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBookmarkCreateCallback != nil {
		stage.OnAfterBookmarkCreateCallback.OnAfterCreate(stage, bookmark)
	}
}

func (bookmark *Bookmark) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBookmarkUpdateCallback != nil {
		var frontBookmark *Bookmark
		if front != nil {
			frontBookmark, _ = front.(*Bookmark)
		}
		stage.OnAfterBookmarkUpdateCallback.OnAfterUpdate(stage, bookmark, frontBookmark)
	}
}

func (bookmark *Bookmark) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBookmarkDeleteCallback != nil {
		var frontBookmark *Bookmark
		if front != nil {
			frontBookmark, _ = front.(*Bookmark)
		}
		stage.OnAfterBookmarkDeleteCallback.OnAfterDelete(stage, bookmark, frontBookmark)
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

func (breath_mark *Breath_mark) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBreath_markCreateCallback != nil {
		stage.OnAfterBreath_markCreateCallback.OnAfterCreate(stage, breath_mark)
	}
}

func (breath_mark *Breath_mark) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBreath_markUpdateCallback != nil {
		var frontBreath_mark *Breath_mark
		if front != nil {
			frontBreath_mark, _ = front.(*Breath_mark)
		}
		stage.OnAfterBreath_markUpdateCallback.OnAfterUpdate(stage, breath_mark, frontBreath_mark)
	}
}

func (breath_mark *Breath_mark) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBreath_markDeleteCallback != nil {
		var frontBreath_mark *Breath_mark
		if front != nil {
			frontBreath_mark, _ = front.(*Breath_mark)
		}
		stage.OnAfterBreath_markDeleteCallback.OnAfterDelete(stage, breath_mark, frontBreath_mark)
	}
}

func (caesura *Caesura) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCaesuraCreateCallback != nil {
		stage.OnAfterCaesuraCreateCallback.OnAfterCreate(stage, caesura)
	}
}

func (caesura *Caesura) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCaesuraUpdateCallback != nil {
		var frontCaesura *Caesura
		if front != nil {
			frontCaesura, _ = front.(*Caesura)
		}
		stage.OnAfterCaesuraUpdateCallback.OnAfterUpdate(stage, caesura, frontCaesura)
	}
}

func (caesura *Caesura) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCaesuraDeleteCallback != nil {
		var frontCaesura *Caesura
		if front != nil {
			frontCaesura, _ = front.(*Caesura)
		}
		stage.OnAfterCaesuraDeleteCallback.OnAfterDelete(stage, caesura, frontCaesura)
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

func (coda *Coda) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodaCreateCallback != nil {
		stage.OnAfterCodaCreateCallback.OnAfterCreate(stage, coda)
	}
}

func (coda *Coda) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodaUpdateCallback != nil {
		var frontCoda *Coda
		if front != nil {
			frontCoda, _ = front.(*Coda)
		}
		stage.OnAfterCodaUpdateCallback.OnAfterUpdate(stage, coda, frontCoda)
	}
}

func (coda *Coda) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodaDeleteCallback != nil {
		var frontCoda *Coda
		if front != nil {
			frontCoda, _ = front.(*Coda)
		}
		stage.OnAfterCodaDeleteCallback.OnAfterDelete(stage, coda, frontCoda)
	}
}

func (credit *Credit) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCreditCreateCallback != nil {
		stage.OnAfterCreditCreateCallback.OnAfterCreate(stage, credit)
	}
}

func (credit *Credit) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCreditUpdateCallback != nil {
		var frontCredit *Credit
		if front != nil {
			frontCredit, _ = front.(*Credit)
		}
		stage.OnAfterCreditUpdateCallback.OnAfterUpdate(stage, credit, frontCredit)
	}
}

func (credit *Credit) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCreditDeleteCallback != nil {
		var frontCredit *Credit
		if front != nil {
			frontCredit, _ = front.(*Credit)
		}
		stage.OnAfterCreditDeleteCallback.OnAfterDelete(stage, credit, frontCredit)
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

func (defaults *Defaults) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDefaultsCreateCallback != nil {
		stage.OnAfterDefaultsCreateCallback.OnAfterCreate(stage, defaults)
	}
}

func (defaults *Defaults) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefaultsUpdateCallback != nil {
		var frontDefaults *Defaults
		if front != nil {
			frontDefaults, _ = front.(*Defaults)
		}
		stage.OnAfterDefaultsUpdateCallback.OnAfterUpdate(stage, defaults, frontDefaults)
	}
}

func (defaults *Defaults) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefaultsDeleteCallback != nil {
		var frontDefaults *Defaults
		if front != nil {
			frontDefaults, _ = front.(*Defaults)
		}
		stage.OnAfterDefaultsDeleteCallback.OnAfterDelete(stage, defaults, frontDefaults)
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

func (degree_alter *Degree_alter) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDegree_alterCreateCallback != nil {
		stage.OnAfterDegree_alterCreateCallback.OnAfterCreate(stage, degree_alter)
	}
}

func (degree_alter *Degree_alter) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegree_alterUpdateCallback != nil {
		var frontDegree_alter *Degree_alter
		if front != nil {
			frontDegree_alter, _ = front.(*Degree_alter)
		}
		stage.OnAfterDegree_alterUpdateCallback.OnAfterUpdate(stage, degree_alter, frontDegree_alter)
	}
}

func (degree_alter *Degree_alter) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegree_alterDeleteCallback != nil {
		var frontDegree_alter *Degree_alter
		if front != nil {
			frontDegree_alter, _ = front.(*Degree_alter)
		}
		stage.OnAfterDegree_alterDeleteCallback.OnAfterDelete(stage, degree_alter, frontDegree_alter)
	}
}

func (degree_type *Degree_type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDegree_typeCreateCallback != nil {
		stage.OnAfterDegree_typeCreateCallback.OnAfterCreate(stage, degree_type)
	}
}

func (degree_type *Degree_type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegree_typeUpdateCallback != nil {
		var frontDegree_type *Degree_type
		if front != nil {
			frontDegree_type, _ = front.(*Degree_type)
		}
		stage.OnAfterDegree_typeUpdateCallback.OnAfterUpdate(stage, degree_type, frontDegree_type)
	}
}

func (degree_type *Degree_type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegree_typeDeleteCallback != nil {
		var frontDegree_type *Degree_type
		if front != nil {
			frontDegree_type, _ = front.(*Degree_type)
		}
		stage.OnAfterDegree_typeDeleteCallback.OnAfterDelete(stage, degree_type, frontDegree_type)
	}
}

func (degree_value *Degree_value) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDegree_valueCreateCallback != nil {
		stage.OnAfterDegree_valueCreateCallback.OnAfterCreate(stage, degree_value)
	}
}

func (degree_value *Degree_value) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegree_valueUpdateCallback != nil {
		var frontDegree_value *Degree_value
		if front != nil {
			frontDegree_value, _ = front.(*Degree_value)
		}
		stage.OnAfterDegree_valueUpdateCallback.OnAfterUpdate(stage, degree_value, frontDegree_value)
	}
}

func (degree_value *Degree_value) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegree_valueDeleteCallback != nil {
		var frontDegree_value *Degree_value
		if front != nil {
			frontDegree_value, _ = front.(*Degree_value)
		}
		stage.OnAfterDegree_valueDeleteCallback.OnAfterDelete(stage, degree_value, frontDegree_value)
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

func (direction_type *Direction_type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirection_typeCreateCallback != nil {
		stage.OnAfterDirection_typeCreateCallback.OnAfterCreate(stage, direction_type)
	}
}

func (direction_type *Direction_type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirection_typeUpdateCallback != nil {
		var frontDirection_type *Direction_type
		if front != nil {
			frontDirection_type, _ = front.(*Direction_type)
		}
		stage.OnAfterDirection_typeUpdateCallback.OnAfterUpdate(stage, direction_type, frontDirection_type)
	}
}

func (direction_type *Direction_type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirection_typeDeleteCallback != nil {
		var frontDirection_type *Direction_type
		if front != nil {
			frontDirection_type, _ = front.(*Direction_type)
		}
		stage.OnAfterDirection_typeDeleteCallback.OnAfterDelete(stage, direction_type, frontDirection_type)
	}
}

func (distance *Distance) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDistanceCreateCallback != nil {
		stage.OnAfterDistanceCreateCallback.OnAfterCreate(stage, distance)
	}
}

func (distance *Distance) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceUpdateCallback != nil {
		var frontDistance *Distance
		if front != nil {
			frontDistance, _ = front.(*Distance)
		}
		stage.OnAfterDistanceUpdateCallback.OnAfterUpdate(stage, distance, frontDistance)
	}
}

func (distance *Distance) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceDeleteCallback != nil {
		var frontDistance *Distance
		if front != nil {
			frontDistance, _ = front.(*Distance)
		}
		stage.OnAfterDistanceDeleteCallback.OnAfterDelete(stage, distance, frontDistance)
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

func (dynamics *Dynamics) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDynamicsCreateCallback != nil {
		stage.OnAfterDynamicsCreateCallback.OnAfterCreate(stage, dynamics)
	}
}

func (dynamics *Dynamics) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicsUpdateCallback != nil {
		var frontDynamics *Dynamics
		if front != nil {
			frontDynamics, _ = front.(*Dynamics)
		}
		stage.OnAfterDynamicsUpdateCallback.OnAfterUpdate(stage, dynamics, frontDynamics)
	}
}

func (dynamics *Dynamics) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicsDeleteCallback != nil {
		var frontDynamics *Dynamics
		if front != nil {
			frontDynamics, _ = front.(*Dynamics)
		}
		stage.OnAfterDynamicsDeleteCallback.OnAfterDelete(stage, dynamics, frontDynamics)
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

func (elision *Elision) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElisionCreateCallback != nil {
		stage.OnAfterElisionCreateCallback.OnAfterCreate(stage, elision)
	}
}

func (elision *Elision) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElisionUpdateCallback != nil {
		var frontElision *Elision
		if front != nil {
			frontElision, _ = front.(*Elision)
		}
		stage.OnAfterElisionUpdateCallback.OnAfterUpdate(stage, elision, frontElision)
	}
}

func (elision *Elision) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElisionDeleteCallback != nil {
		var frontElision *Elision
		if front != nil {
			frontElision, _ = front.(*Elision)
		}
		stage.OnAfterElisionDeleteCallback.OnAfterDelete(stage, elision, frontElision)
	}
}

func (empty *Empty) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyCreateCallback != nil {
		stage.OnAfterEmptyCreateCallback.OnAfterCreate(stage, empty)
	}
}

func (empty *Empty) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyUpdateCallback != nil {
		var frontEmpty *Empty
		if front != nil {
			frontEmpty, _ = front.(*Empty)
		}
		stage.OnAfterEmptyUpdateCallback.OnAfterUpdate(stage, empty, frontEmpty)
	}
}

func (empty *Empty) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyDeleteCallback != nil {
		var frontEmpty *Empty
		if front != nil {
			frontEmpty, _ = front.(*Empty)
		}
		stage.OnAfterEmptyDeleteCallback.OnAfterDelete(stage, empty, frontEmpty)
	}
}

func (empty_font *Empty_font) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmpty_fontCreateCallback != nil {
		stage.OnAfterEmpty_fontCreateCallback.OnAfterCreate(stage, empty_font)
	}
}

func (empty_font *Empty_font) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_fontUpdateCallback != nil {
		var frontEmpty_font *Empty_font
		if front != nil {
			frontEmpty_font, _ = front.(*Empty_font)
		}
		stage.OnAfterEmpty_fontUpdateCallback.OnAfterUpdate(stage, empty_font, frontEmpty_font)
	}
}

func (empty_font *Empty_font) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_fontDeleteCallback != nil {
		var frontEmpty_font *Empty_font
		if front != nil {
			frontEmpty_font, _ = front.(*Empty_font)
		}
		stage.OnAfterEmpty_fontDeleteCallback.OnAfterDelete(stage, empty_font, frontEmpty_font)
	}
}

func (empty_line *Empty_line) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmpty_lineCreateCallback != nil {
		stage.OnAfterEmpty_lineCreateCallback.OnAfterCreate(stage, empty_line)
	}
}

func (empty_line *Empty_line) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_lineUpdateCallback != nil {
		var frontEmpty_line *Empty_line
		if front != nil {
			frontEmpty_line, _ = front.(*Empty_line)
		}
		stage.OnAfterEmpty_lineUpdateCallback.OnAfterUpdate(stage, empty_line, frontEmpty_line)
	}
}

func (empty_line *Empty_line) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_lineDeleteCallback != nil {
		var frontEmpty_line *Empty_line
		if front != nil {
			frontEmpty_line, _ = front.(*Empty_line)
		}
		stage.OnAfterEmpty_lineDeleteCallback.OnAfterDelete(stage, empty_line, frontEmpty_line)
	}
}

func (empty_placement *Empty_placement) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmpty_placementCreateCallback != nil {
		stage.OnAfterEmpty_placementCreateCallback.OnAfterCreate(stage, empty_placement)
	}
}

func (empty_placement *Empty_placement) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_placementUpdateCallback != nil {
		var frontEmpty_placement *Empty_placement
		if front != nil {
			frontEmpty_placement, _ = front.(*Empty_placement)
		}
		stage.OnAfterEmpty_placementUpdateCallback.OnAfterUpdate(stage, empty_placement, frontEmpty_placement)
	}
}

func (empty_placement *Empty_placement) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_placementDeleteCallback != nil {
		var frontEmpty_placement *Empty_placement
		if front != nil {
			frontEmpty_placement, _ = front.(*Empty_placement)
		}
		stage.OnAfterEmpty_placementDeleteCallback.OnAfterDelete(stage, empty_placement, frontEmpty_placement)
	}
}

func (empty_placement_smufl *Empty_placement_smufl) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmpty_placement_smuflCreateCallback != nil {
		stage.OnAfterEmpty_placement_smuflCreateCallback.OnAfterCreate(stage, empty_placement_smufl)
	}
}

func (empty_placement_smufl *Empty_placement_smufl) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_placement_smuflUpdateCallback != nil {
		var frontEmpty_placement_smufl *Empty_placement_smufl
		if front != nil {
			frontEmpty_placement_smufl, _ = front.(*Empty_placement_smufl)
		}
		stage.OnAfterEmpty_placement_smuflUpdateCallback.OnAfterUpdate(stage, empty_placement_smufl, frontEmpty_placement_smufl)
	}
}

func (empty_placement_smufl *Empty_placement_smufl) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_placement_smuflDeleteCallback != nil {
		var frontEmpty_placement_smufl *Empty_placement_smufl
		if front != nil {
			frontEmpty_placement_smufl, _ = front.(*Empty_placement_smufl)
		}
		stage.OnAfterEmpty_placement_smuflDeleteCallback.OnAfterDelete(stage, empty_placement_smufl, frontEmpty_placement_smufl)
	}
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmpty_print_object_style_alignCreateCallback != nil {
		stage.OnAfterEmpty_print_object_style_alignCreateCallback.OnAfterCreate(stage, empty_print_object_style_align)
	}
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_print_object_style_alignUpdateCallback != nil {
		var frontEmpty_print_object_style_align *Empty_print_object_style_align
		if front != nil {
			frontEmpty_print_object_style_align, _ = front.(*Empty_print_object_style_align)
		}
		stage.OnAfterEmpty_print_object_style_alignUpdateCallback.OnAfterUpdate(stage, empty_print_object_style_align, frontEmpty_print_object_style_align)
	}
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_print_object_style_alignDeleteCallback != nil {
		var frontEmpty_print_object_style_align *Empty_print_object_style_align
		if front != nil {
			frontEmpty_print_object_style_align, _ = front.(*Empty_print_object_style_align)
		}
		stage.OnAfterEmpty_print_object_style_alignDeleteCallback.OnAfterDelete(stage, empty_print_object_style_align, frontEmpty_print_object_style_align)
	}
}

func (empty_print_style *Empty_print_style) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmpty_print_styleCreateCallback != nil {
		stage.OnAfterEmpty_print_styleCreateCallback.OnAfterCreate(stage, empty_print_style)
	}
}

func (empty_print_style *Empty_print_style) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_print_styleUpdateCallback != nil {
		var frontEmpty_print_style *Empty_print_style
		if front != nil {
			frontEmpty_print_style, _ = front.(*Empty_print_style)
		}
		stage.OnAfterEmpty_print_styleUpdateCallback.OnAfterUpdate(stage, empty_print_style, frontEmpty_print_style)
	}
}

func (empty_print_style *Empty_print_style) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_print_styleDeleteCallback != nil {
		var frontEmpty_print_style *Empty_print_style
		if front != nil {
			frontEmpty_print_style, _ = front.(*Empty_print_style)
		}
		stage.OnAfterEmpty_print_styleDeleteCallback.OnAfterDelete(stage, empty_print_style, frontEmpty_print_style)
	}
}

func (empty_print_style_align *Empty_print_style_align) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmpty_print_style_alignCreateCallback != nil {
		stage.OnAfterEmpty_print_style_alignCreateCallback.OnAfterCreate(stage, empty_print_style_align)
	}
}

func (empty_print_style_align *Empty_print_style_align) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_print_style_alignUpdateCallback != nil {
		var frontEmpty_print_style_align *Empty_print_style_align
		if front != nil {
			frontEmpty_print_style_align, _ = front.(*Empty_print_style_align)
		}
		stage.OnAfterEmpty_print_style_alignUpdateCallback.OnAfterUpdate(stage, empty_print_style_align, frontEmpty_print_style_align)
	}
}

func (empty_print_style_align *Empty_print_style_align) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_print_style_alignDeleteCallback != nil {
		var frontEmpty_print_style_align *Empty_print_style_align
		if front != nil {
			frontEmpty_print_style_align, _ = front.(*Empty_print_style_align)
		}
		stage.OnAfterEmpty_print_style_alignDeleteCallback.OnAfterDelete(stage, empty_print_style_align, frontEmpty_print_style_align)
	}
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmpty_print_style_align_idCreateCallback != nil {
		stage.OnAfterEmpty_print_style_align_idCreateCallback.OnAfterCreate(stage, empty_print_style_align_id)
	}
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_print_style_align_idUpdateCallback != nil {
		var frontEmpty_print_style_align_id *Empty_print_style_align_id
		if front != nil {
			frontEmpty_print_style_align_id, _ = front.(*Empty_print_style_align_id)
		}
		stage.OnAfterEmpty_print_style_align_idUpdateCallback.OnAfterUpdate(stage, empty_print_style_align_id, frontEmpty_print_style_align_id)
	}
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_print_style_align_idDeleteCallback != nil {
		var frontEmpty_print_style_align_id *Empty_print_style_align_id
		if front != nil {
			frontEmpty_print_style_align_id, _ = front.(*Empty_print_style_align_id)
		}
		stage.OnAfterEmpty_print_style_align_idDeleteCallback.OnAfterDelete(stage, empty_print_style_align_id, frontEmpty_print_style_align_id)
	}
}

func (empty_trill_sound *Empty_trill_sound) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmpty_trill_soundCreateCallback != nil {
		stage.OnAfterEmpty_trill_soundCreateCallback.OnAfterCreate(stage, empty_trill_sound)
	}
}

func (empty_trill_sound *Empty_trill_sound) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_trill_soundUpdateCallback != nil {
		var frontEmpty_trill_sound *Empty_trill_sound
		if front != nil {
			frontEmpty_trill_sound, _ = front.(*Empty_trill_sound)
		}
		stage.OnAfterEmpty_trill_soundUpdateCallback.OnAfterUpdate(stage, empty_trill_sound, frontEmpty_trill_sound)
	}
}

func (empty_trill_sound *Empty_trill_sound) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmpty_trill_soundDeleteCallback != nil {
		var frontEmpty_trill_sound *Empty_trill_sound
		if front != nil {
			frontEmpty_trill_sound, _ = front.(*Empty_trill_sound)
		}
		stage.OnAfterEmpty_trill_soundDeleteCallback.OnAfterDelete(stage, empty_trill_sound, frontEmpty_trill_sound)
	}
}

func (encoding *Encoding) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEncodingCreateCallback != nil {
		stage.OnAfterEncodingCreateCallback.OnAfterCreate(stage, encoding)
	}
}

func (encoding *Encoding) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEncodingUpdateCallback != nil {
		var frontEncoding *Encoding
		if front != nil {
			frontEncoding, _ = front.(*Encoding)
		}
		stage.OnAfterEncodingUpdateCallback.OnAfterUpdate(stage, encoding, frontEncoding)
	}
}

func (encoding *Encoding) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEncodingDeleteCallback != nil {
		var frontEncoding *Encoding
		if front != nil {
			frontEncoding, _ = front.(*Encoding)
		}
		stage.OnAfterEncodingDeleteCallback.OnAfterDelete(stage, encoding, frontEncoding)
	}
}

func (ending *Ending) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEndingCreateCallback != nil {
		stage.OnAfterEndingCreateCallback.OnAfterCreate(stage, ending)
	}
}

func (ending *Ending) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndingUpdateCallback != nil {
		var frontEnding *Ending
		if front != nil {
			frontEnding, _ = front.(*Ending)
		}
		stage.OnAfterEndingUpdateCallback.OnAfterUpdate(stage, ending, frontEnding)
	}
}

func (ending *Ending) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndingDeleteCallback != nil {
		var frontEnding *Ending
		if front != nil {
			frontEnding, _ = front.(*Ending)
		}
		stage.OnAfterEndingDeleteCallback.OnAfterDelete(stage, ending, frontEnding)
	}
}

func (extend *Extend) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterExtendCreateCallback != nil {
		stage.OnAfterExtendCreateCallback.OnAfterCreate(stage, extend)
	}
}

func (extend *Extend) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExtendUpdateCallback != nil {
		var frontExtend *Extend
		if front != nil {
			frontExtend, _ = front.(*Extend)
		}
		stage.OnAfterExtendUpdateCallback.OnAfterUpdate(stage, extend, frontExtend)
	}
}

func (extend *Extend) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExtendDeleteCallback != nil {
		var frontExtend *Extend
		if front != nil {
			frontExtend, _ = front.(*Extend)
		}
		stage.OnAfterExtendDeleteCallback.OnAfterDelete(stage, extend, frontExtend)
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

func (fermata *Fermata) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFermataCreateCallback != nil {
		stage.OnAfterFermataCreateCallback.OnAfterCreate(stage, fermata)
	}
}

func (fermata *Fermata) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFermataUpdateCallback != nil {
		var frontFermata *Fermata
		if front != nil {
			frontFermata, _ = front.(*Fermata)
		}
		stage.OnAfterFermataUpdateCallback.OnAfterUpdate(stage, fermata, frontFermata)
	}
}

func (fermata *Fermata) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFermataDeleteCallback != nil {
		var frontFermata *Fermata
		if front != nil {
			frontFermata, _ = front.(*Fermata)
		}
		stage.OnAfterFermataDeleteCallback.OnAfterDelete(stage, fermata, frontFermata)
	}
}

func (figure *Figure) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFigureCreateCallback != nil {
		stage.OnAfterFigureCreateCallback.OnAfterCreate(stage, figure)
	}
}

func (figure *Figure) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFigureUpdateCallback != nil {
		var frontFigure *Figure
		if front != nil {
			frontFigure, _ = front.(*Figure)
		}
		stage.OnAfterFigureUpdateCallback.OnAfterUpdate(stage, figure, frontFigure)
	}
}

func (figure *Figure) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFigureDeleteCallback != nil {
		var frontFigure *Figure
		if front != nil {
			frontFigure, _ = front.(*Figure)
		}
		stage.OnAfterFigureDeleteCallback.OnAfterDelete(stage, figure, frontFigure)
	}
}

func (figured_bass *Figured_bass) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFigured_bassCreateCallback != nil {
		stage.OnAfterFigured_bassCreateCallback.OnAfterCreate(stage, figured_bass)
	}
}

func (figured_bass *Figured_bass) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFigured_bassUpdateCallback != nil {
		var frontFigured_bass *Figured_bass
		if front != nil {
			frontFigured_bass, _ = front.(*Figured_bass)
		}
		stage.OnAfterFigured_bassUpdateCallback.OnAfterUpdate(stage, figured_bass, frontFigured_bass)
	}
}

func (figured_bass *Figured_bass) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFigured_bassDeleteCallback != nil {
		var frontFigured_bass *Figured_bass
		if front != nil {
			frontFigured_bass, _ = front.(*Figured_bass)
		}
		stage.OnAfterFigured_bassDeleteCallback.OnAfterDelete(stage, figured_bass, frontFigured_bass)
	}
}

func (fingering *Fingering) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFingeringCreateCallback != nil {
		stage.OnAfterFingeringCreateCallback.OnAfterCreate(stage, fingering)
	}
}

func (fingering *Fingering) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFingeringUpdateCallback != nil {
		var frontFingering *Fingering
		if front != nil {
			frontFingering, _ = front.(*Fingering)
		}
		stage.OnAfterFingeringUpdateCallback.OnAfterUpdate(stage, fingering, frontFingering)
	}
}

func (fingering *Fingering) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFingeringDeleteCallback != nil {
		var frontFingering *Fingering
		if front != nil {
			frontFingering, _ = front.(*Fingering)
		}
		stage.OnAfterFingeringDeleteCallback.OnAfterDelete(stage, fingering, frontFingering)
	}
}

func (first_fret *First_fret) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFirst_fretCreateCallback != nil {
		stage.OnAfterFirst_fretCreateCallback.OnAfterCreate(stage, first_fret)
	}
}

func (first_fret *First_fret) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFirst_fretUpdateCallback != nil {
		var frontFirst_fret *First_fret
		if front != nil {
			frontFirst_fret, _ = front.(*First_fret)
		}
		stage.OnAfterFirst_fretUpdateCallback.OnAfterUpdate(stage, first_fret, frontFirst_fret)
	}
}

func (first_fret *First_fret) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFirst_fretDeleteCallback != nil {
		var frontFirst_fret *First_fret
		if front != nil {
			frontFirst_fret, _ = front.(*First_fret)
		}
		stage.OnAfterFirst_fretDeleteCallback.OnAfterDelete(stage, first_fret, frontFirst_fret)
	}
}

func (for_part *For_part) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFor_partCreateCallback != nil {
		stage.OnAfterFor_partCreateCallback.OnAfterCreate(stage, for_part)
	}
}

func (for_part *For_part) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFor_partUpdateCallback != nil {
		var frontFor_part *For_part
		if front != nil {
			frontFor_part, _ = front.(*For_part)
		}
		stage.OnAfterFor_partUpdateCallback.OnAfterUpdate(stage, for_part, frontFor_part)
	}
}

func (for_part *For_part) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFor_partDeleteCallback != nil {
		var frontFor_part *For_part
		if front != nil {
			frontFor_part, _ = front.(*For_part)
		}
		stage.OnAfterFor_partDeleteCallback.OnAfterDelete(stage, for_part, frontFor_part)
	}
}

func (formatted_symbol *Formatted_symbol) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormatted_symbolCreateCallback != nil {
		stage.OnAfterFormatted_symbolCreateCallback.OnAfterCreate(stage, formatted_symbol)
	}
}

func (formatted_symbol *Formatted_symbol) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormatted_symbolUpdateCallback != nil {
		var frontFormatted_symbol *Formatted_symbol
		if front != nil {
			frontFormatted_symbol, _ = front.(*Formatted_symbol)
		}
		stage.OnAfterFormatted_symbolUpdateCallback.OnAfterUpdate(stage, formatted_symbol, frontFormatted_symbol)
	}
}

func (formatted_symbol *Formatted_symbol) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormatted_symbolDeleteCallback != nil {
		var frontFormatted_symbol *Formatted_symbol
		if front != nil {
			frontFormatted_symbol, _ = front.(*Formatted_symbol)
		}
		stage.OnAfterFormatted_symbolDeleteCallback.OnAfterDelete(stage, formatted_symbol, frontFormatted_symbol)
	}
}

func (formatted_symbol_id *Formatted_symbol_id) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormatted_symbol_idCreateCallback != nil {
		stage.OnAfterFormatted_symbol_idCreateCallback.OnAfterCreate(stage, formatted_symbol_id)
	}
}

func (formatted_symbol_id *Formatted_symbol_id) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormatted_symbol_idUpdateCallback != nil {
		var frontFormatted_symbol_id *Formatted_symbol_id
		if front != nil {
			frontFormatted_symbol_id, _ = front.(*Formatted_symbol_id)
		}
		stage.OnAfterFormatted_symbol_idUpdateCallback.OnAfterUpdate(stage, formatted_symbol_id, frontFormatted_symbol_id)
	}
}

func (formatted_symbol_id *Formatted_symbol_id) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormatted_symbol_idDeleteCallback != nil {
		var frontFormatted_symbol_id *Formatted_symbol_id
		if front != nil {
			frontFormatted_symbol_id, _ = front.(*Formatted_symbol_id)
		}
		stage.OnAfterFormatted_symbol_idDeleteCallback.OnAfterDelete(stage, formatted_symbol_id, frontFormatted_symbol_id)
	}
}

func (formatted_text *Formatted_text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormatted_textCreateCallback != nil {
		stage.OnAfterFormatted_textCreateCallback.OnAfterCreate(stage, formatted_text)
	}
}

func (formatted_text *Formatted_text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormatted_textUpdateCallback != nil {
		var frontFormatted_text *Formatted_text
		if front != nil {
			frontFormatted_text, _ = front.(*Formatted_text)
		}
		stage.OnAfterFormatted_textUpdateCallback.OnAfterUpdate(stage, formatted_text, frontFormatted_text)
	}
}

func (formatted_text *Formatted_text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormatted_textDeleteCallback != nil {
		var frontFormatted_text *Formatted_text
		if front != nil {
			frontFormatted_text, _ = front.(*Formatted_text)
		}
		stage.OnAfterFormatted_textDeleteCallback.OnAfterDelete(stage, formatted_text, frontFormatted_text)
	}
}

func (formatted_text_id *Formatted_text_id) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormatted_text_idCreateCallback != nil {
		stage.OnAfterFormatted_text_idCreateCallback.OnAfterCreate(stage, formatted_text_id)
	}
}

func (formatted_text_id *Formatted_text_id) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormatted_text_idUpdateCallback != nil {
		var frontFormatted_text_id *Formatted_text_id
		if front != nil {
			frontFormatted_text_id, _ = front.(*Formatted_text_id)
		}
		stage.OnAfterFormatted_text_idUpdateCallback.OnAfterUpdate(stage, formatted_text_id, frontFormatted_text_id)
	}
}

func (formatted_text_id *Formatted_text_id) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormatted_text_idDeleteCallback != nil {
		var frontFormatted_text_id *Formatted_text_id
		if front != nil {
			frontFormatted_text_id, _ = front.(*Formatted_text_id)
		}
		stage.OnAfterFormatted_text_idDeleteCallback.OnAfterDelete(stage, formatted_text_id, frontFormatted_text_id)
	}
}

func (forward *Forward) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterForwardCreateCallback != nil {
		stage.OnAfterForwardCreateCallback.OnAfterCreate(stage, forward)
	}
}

func (forward *Forward) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterForwardUpdateCallback != nil {
		var frontForward *Forward
		if front != nil {
			frontForward, _ = front.(*Forward)
		}
		stage.OnAfterForwardUpdateCallback.OnAfterUpdate(stage, forward, frontForward)
	}
}

func (forward *Forward) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterForwardDeleteCallback != nil {
		var frontForward *Forward
		if front != nil {
			frontForward, _ = front.(*Forward)
		}
		stage.OnAfterForwardDeleteCallback.OnAfterDelete(stage, forward, frontForward)
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

func (frame_note *Frame_note) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFrame_noteCreateCallback != nil {
		stage.OnAfterFrame_noteCreateCallback.OnAfterCreate(stage, frame_note)
	}
}

func (frame_note *Frame_note) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFrame_noteUpdateCallback != nil {
		var frontFrame_note *Frame_note
		if front != nil {
			frontFrame_note, _ = front.(*Frame_note)
		}
		stage.OnAfterFrame_noteUpdateCallback.OnAfterUpdate(stage, frame_note, frontFrame_note)
	}
}

func (frame_note *Frame_note) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFrame_noteDeleteCallback != nil {
		var frontFrame_note *Frame_note
		if front != nil {
			frontFrame_note, _ = front.(*Frame_note)
		}
		stage.OnAfterFrame_noteDeleteCallback.OnAfterDelete(stage, frame_note, frontFrame_note)
	}
}

func (fret *Fret) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFretCreateCallback != nil {
		stage.OnAfterFretCreateCallback.OnAfterCreate(stage, fret)
	}
}

func (fret *Fret) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFretUpdateCallback != nil {
		var frontFret *Fret
		if front != nil {
			frontFret, _ = front.(*Fret)
		}
		stage.OnAfterFretUpdateCallback.OnAfterUpdate(stage, fret, frontFret)
	}
}

func (fret *Fret) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFretDeleteCallback != nil {
		var frontFret *Fret
		if front != nil {
			frontFret, _ = front.(*Fret)
		}
		stage.OnAfterFretDeleteCallback.OnAfterDelete(stage, fret, frontFret)
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

func (glissando *Glissando) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGlissandoCreateCallback != nil {
		stage.OnAfterGlissandoCreateCallback.OnAfterCreate(stage, glissando)
	}
}

func (glissando *Glissando) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlissandoUpdateCallback != nil {
		var frontGlissando *Glissando
		if front != nil {
			frontGlissando, _ = front.(*Glissando)
		}
		stage.OnAfterGlissandoUpdateCallback.OnAfterUpdate(stage, glissando, frontGlissando)
	}
}

func (glissando *Glissando) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlissandoDeleteCallback != nil {
		var frontGlissando *Glissando
		if front != nil {
			frontGlissando, _ = front.(*Glissando)
		}
		stage.OnAfterGlissandoDeleteCallback.OnAfterDelete(stage, glissando, frontGlissando)
	}
}

func (glyph *Glyph) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGlyphCreateCallback != nil {
		stage.OnAfterGlyphCreateCallback.OnAfterCreate(stage, glyph)
	}
}

func (glyph *Glyph) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlyphUpdateCallback != nil {
		var frontGlyph *Glyph
		if front != nil {
			frontGlyph, _ = front.(*Glyph)
		}
		stage.OnAfterGlyphUpdateCallback.OnAfterUpdate(stage, glyph, frontGlyph)
	}
}

func (glyph *Glyph) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlyphDeleteCallback != nil {
		var frontGlyph *Glyph
		if front != nil {
			frontGlyph, _ = front.(*Glyph)
		}
		stage.OnAfterGlyphDeleteCallback.OnAfterDelete(stage, glyph, frontGlyph)
	}
}

func (grace *Grace) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGraceCreateCallback != nil {
		stage.OnAfterGraceCreateCallback.OnAfterCreate(stage, grace)
	}
}

func (grace *Grace) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGraceUpdateCallback != nil {
		var frontGrace *Grace
		if front != nil {
			frontGrace, _ = front.(*Grace)
		}
		stage.OnAfterGraceUpdateCallback.OnAfterUpdate(stage, grace, frontGrace)
	}
}

func (grace *Grace) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGraceDeleteCallback != nil {
		var frontGrace *Grace
		if front != nil {
			frontGrace, _ = front.(*Grace)
		}
		stage.OnAfterGraceDeleteCallback.OnAfterDelete(stage, grace, frontGrace)
	}
}

func (group_barline *Group_barline) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroup_barlineCreateCallback != nil {
		stage.OnAfterGroup_barlineCreateCallback.OnAfterCreate(stage, group_barline)
	}
}

func (group_barline *Group_barline) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroup_barlineUpdateCallback != nil {
		var frontGroup_barline *Group_barline
		if front != nil {
			frontGroup_barline, _ = front.(*Group_barline)
		}
		stage.OnAfterGroup_barlineUpdateCallback.OnAfterUpdate(stage, group_barline, frontGroup_barline)
	}
}

func (group_barline *Group_barline) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroup_barlineDeleteCallback != nil {
		var frontGroup_barline *Group_barline
		if front != nil {
			frontGroup_barline, _ = front.(*Group_barline)
		}
		stage.OnAfterGroup_barlineDeleteCallback.OnAfterDelete(stage, group_barline, frontGroup_barline)
	}
}

func (group_name *Group_name) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroup_nameCreateCallback != nil {
		stage.OnAfterGroup_nameCreateCallback.OnAfterCreate(stage, group_name)
	}
}

func (group_name *Group_name) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroup_nameUpdateCallback != nil {
		var frontGroup_name *Group_name
		if front != nil {
			frontGroup_name, _ = front.(*Group_name)
		}
		stage.OnAfterGroup_nameUpdateCallback.OnAfterUpdate(stage, group_name, frontGroup_name)
	}
}

func (group_name *Group_name) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroup_nameDeleteCallback != nil {
		var frontGroup_name *Group_name
		if front != nil {
			frontGroup_name, _ = front.(*Group_name)
		}
		stage.OnAfterGroup_nameDeleteCallback.OnAfterDelete(stage, group_name, frontGroup_name)
	}
}

func (group_symbol *Group_symbol) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroup_symbolCreateCallback != nil {
		stage.OnAfterGroup_symbolCreateCallback.OnAfterCreate(stage, group_symbol)
	}
}

func (group_symbol *Group_symbol) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroup_symbolUpdateCallback != nil {
		var frontGroup_symbol *Group_symbol
		if front != nil {
			frontGroup_symbol, _ = front.(*Group_symbol)
		}
		stage.OnAfterGroup_symbolUpdateCallback.OnAfterUpdate(stage, group_symbol, frontGroup_symbol)
	}
}

func (group_symbol *Group_symbol) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroup_symbolDeleteCallback != nil {
		var frontGroup_symbol *Group_symbol
		if front != nil {
			frontGroup_symbol, _ = front.(*Group_symbol)
		}
		stage.OnAfterGroup_symbolDeleteCallback.OnAfterDelete(stage, group_symbol, frontGroup_symbol)
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

func (hammer_on_pull_off *Hammer_on_pull_off) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHammer_on_pull_offCreateCallback != nil {
		stage.OnAfterHammer_on_pull_offCreateCallback.OnAfterCreate(stage, hammer_on_pull_off)
	}
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHammer_on_pull_offUpdateCallback != nil {
		var frontHammer_on_pull_off *Hammer_on_pull_off
		if front != nil {
			frontHammer_on_pull_off, _ = front.(*Hammer_on_pull_off)
		}
		stage.OnAfterHammer_on_pull_offUpdateCallback.OnAfterUpdate(stage, hammer_on_pull_off, frontHammer_on_pull_off)
	}
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHammer_on_pull_offDeleteCallback != nil {
		var frontHammer_on_pull_off *Hammer_on_pull_off
		if front != nil {
			frontHammer_on_pull_off, _ = front.(*Hammer_on_pull_off)
		}
		stage.OnAfterHammer_on_pull_offDeleteCallback.OnAfterDelete(stage, hammer_on_pull_off, frontHammer_on_pull_off)
	}
}

func (handbell *Handbell) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHandbellCreateCallback != nil {
		stage.OnAfterHandbellCreateCallback.OnAfterCreate(stage, handbell)
	}
}

func (handbell *Handbell) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHandbellUpdateCallback != nil {
		var frontHandbell *Handbell
		if front != nil {
			frontHandbell, _ = front.(*Handbell)
		}
		stage.OnAfterHandbellUpdateCallback.OnAfterUpdate(stage, handbell, frontHandbell)
	}
}

func (handbell *Handbell) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHandbellDeleteCallback != nil {
		var frontHandbell *Handbell
		if front != nil {
			frontHandbell, _ = front.(*Handbell)
		}
		stage.OnAfterHandbellDeleteCallback.OnAfterDelete(stage, handbell, frontHandbell)
	}
}

func (harmon_closed *Harmon_closed) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarmon_closedCreateCallback != nil {
		stage.OnAfterHarmon_closedCreateCallback.OnAfterCreate(stage, harmon_closed)
	}
}

func (harmon_closed *Harmon_closed) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmon_closedUpdateCallback != nil {
		var frontHarmon_closed *Harmon_closed
		if front != nil {
			frontHarmon_closed, _ = front.(*Harmon_closed)
		}
		stage.OnAfterHarmon_closedUpdateCallback.OnAfterUpdate(stage, harmon_closed, frontHarmon_closed)
	}
}

func (harmon_closed *Harmon_closed) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmon_closedDeleteCallback != nil {
		var frontHarmon_closed *Harmon_closed
		if front != nil {
			frontHarmon_closed, _ = front.(*Harmon_closed)
		}
		stage.OnAfterHarmon_closedDeleteCallback.OnAfterDelete(stage, harmon_closed, frontHarmon_closed)
	}
}

func (harmon_mute *Harmon_mute) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarmon_muteCreateCallback != nil {
		stage.OnAfterHarmon_muteCreateCallback.OnAfterCreate(stage, harmon_mute)
	}
}

func (harmon_mute *Harmon_mute) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmon_muteUpdateCallback != nil {
		var frontHarmon_mute *Harmon_mute
		if front != nil {
			frontHarmon_mute, _ = front.(*Harmon_mute)
		}
		stage.OnAfterHarmon_muteUpdateCallback.OnAfterUpdate(stage, harmon_mute, frontHarmon_mute)
	}
}

func (harmon_mute *Harmon_mute) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmon_muteDeleteCallback != nil {
		var frontHarmon_mute *Harmon_mute
		if front != nil {
			frontHarmon_mute, _ = front.(*Harmon_mute)
		}
		stage.OnAfterHarmon_muteDeleteCallback.OnAfterDelete(stage, harmon_mute, frontHarmon_mute)
	}
}

func (harmonic *Harmonic) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarmonicCreateCallback != nil {
		stage.OnAfterHarmonicCreateCallback.OnAfterCreate(stage, harmonic)
	}
}

func (harmonic *Harmonic) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonicUpdateCallback != nil {
		var frontHarmonic *Harmonic
		if front != nil {
			frontHarmonic, _ = front.(*Harmonic)
		}
		stage.OnAfterHarmonicUpdateCallback.OnAfterUpdate(stage, harmonic, frontHarmonic)
	}
}

func (harmonic *Harmonic) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonicDeleteCallback != nil {
		var frontHarmonic *Harmonic
		if front != nil {
			frontHarmonic, _ = front.(*Harmonic)
		}
		stage.OnAfterHarmonicDeleteCallback.OnAfterDelete(stage, harmonic, frontHarmonic)
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

func (harmony_alter *Harmony_alter) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarmony_alterCreateCallback != nil {
		stage.OnAfterHarmony_alterCreateCallback.OnAfterCreate(stage, harmony_alter)
	}
}

func (harmony_alter *Harmony_alter) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmony_alterUpdateCallback != nil {
		var frontHarmony_alter *Harmony_alter
		if front != nil {
			frontHarmony_alter, _ = front.(*Harmony_alter)
		}
		stage.OnAfterHarmony_alterUpdateCallback.OnAfterUpdate(stage, harmony_alter, frontHarmony_alter)
	}
}

func (harmony_alter *Harmony_alter) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmony_alterDeleteCallback != nil {
		var frontHarmony_alter *Harmony_alter
		if front != nil {
			frontHarmony_alter, _ = front.(*Harmony_alter)
		}
		stage.OnAfterHarmony_alterDeleteCallback.OnAfterDelete(stage, harmony_alter, frontHarmony_alter)
	}
}

func (harp_pedals *Harp_pedals) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarp_pedalsCreateCallback != nil {
		stage.OnAfterHarp_pedalsCreateCallback.OnAfterCreate(stage, harp_pedals)
	}
}

func (harp_pedals *Harp_pedals) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarp_pedalsUpdateCallback != nil {
		var frontHarp_pedals *Harp_pedals
		if front != nil {
			frontHarp_pedals, _ = front.(*Harp_pedals)
		}
		stage.OnAfterHarp_pedalsUpdateCallback.OnAfterUpdate(stage, harp_pedals, frontHarp_pedals)
	}
}

func (harp_pedals *Harp_pedals) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarp_pedalsDeleteCallback != nil {
		var frontHarp_pedals *Harp_pedals
		if front != nil {
			frontHarp_pedals, _ = front.(*Harp_pedals)
		}
		stage.OnAfterHarp_pedalsDeleteCallback.OnAfterDelete(stage, harp_pedals, frontHarp_pedals)
	}
}

func (heel_toe *Heel_toe) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHeel_toeCreateCallback != nil {
		stage.OnAfterHeel_toeCreateCallback.OnAfterCreate(stage, heel_toe)
	}
}

func (heel_toe *Heel_toe) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHeel_toeUpdateCallback != nil {
		var frontHeel_toe *Heel_toe
		if front != nil {
			frontHeel_toe, _ = front.(*Heel_toe)
		}
		stage.OnAfterHeel_toeUpdateCallback.OnAfterUpdate(stage, heel_toe, frontHeel_toe)
	}
}

func (heel_toe *Heel_toe) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHeel_toeDeleteCallback != nil {
		var frontHeel_toe *Heel_toe
		if front != nil {
			frontHeel_toe, _ = front.(*Heel_toe)
		}
		stage.OnAfterHeel_toeDeleteCallback.OnAfterDelete(stage, heel_toe, frontHeel_toe)
	}
}

func (hole *Hole) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoleCreateCallback != nil {
		stage.OnAfterHoleCreateCallback.OnAfterCreate(stage, hole)
	}
}

func (hole *Hole) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoleUpdateCallback != nil {
		var frontHole *Hole
		if front != nil {
			frontHole, _ = front.(*Hole)
		}
		stage.OnAfterHoleUpdateCallback.OnAfterUpdate(stage, hole, frontHole)
	}
}

func (hole *Hole) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoleDeleteCallback != nil {
		var frontHole *Hole
		if front != nil {
			frontHole, _ = front.(*Hole)
		}
		stage.OnAfterHoleDeleteCallback.OnAfterDelete(stage, hole, frontHole)
	}
}

func (hole_closed *Hole_closed) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHole_closedCreateCallback != nil {
		stage.OnAfterHole_closedCreateCallback.OnAfterCreate(stage, hole_closed)
	}
}

func (hole_closed *Hole_closed) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHole_closedUpdateCallback != nil {
		var frontHole_closed *Hole_closed
		if front != nil {
			frontHole_closed, _ = front.(*Hole_closed)
		}
		stage.OnAfterHole_closedUpdateCallback.OnAfterUpdate(stage, hole_closed, frontHole_closed)
	}
}

func (hole_closed *Hole_closed) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHole_closedDeleteCallback != nil {
		var frontHole_closed *Hole_closed
		if front != nil {
			frontHole_closed, _ = front.(*Hole_closed)
		}
		stage.OnAfterHole_closedDeleteCallback.OnAfterDelete(stage, hole_closed, frontHole_closed)
	}
}

func (horizontal_turn *Horizontal_turn) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHorizontal_turnCreateCallback != nil {
		stage.OnAfterHorizontal_turnCreateCallback.OnAfterCreate(stage, horizontal_turn)
	}
}

func (horizontal_turn *Horizontal_turn) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHorizontal_turnUpdateCallback != nil {
		var frontHorizontal_turn *Horizontal_turn
		if front != nil {
			frontHorizontal_turn, _ = front.(*Horizontal_turn)
		}
		stage.OnAfterHorizontal_turnUpdateCallback.OnAfterUpdate(stage, horizontal_turn, frontHorizontal_turn)
	}
}

func (horizontal_turn *Horizontal_turn) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHorizontal_turnDeleteCallback != nil {
		var frontHorizontal_turn *Horizontal_turn
		if front != nil {
			frontHorizontal_turn, _ = front.(*Horizontal_turn)
		}
		stage.OnAfterHorizontal_turnDeleteCallback.OnAfterDelete(stage, horizontal_turn, frontHorizontal_turn)
	}
}

func (identification *Identification) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterIdentificationCreateCallback != nil {
		stage.OnAfterIdentificationCreateCallback.OnAfterCreate(stage, identification)
	}
}

func (identification *Identification) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIdentificationUpdateCallback != nil {
		var frontIdentification *Identification
		if front != nil {
			frontIdentification, _ = front.(*Identification)
		}
		stage.OnAfterIdentificationUpdateCallback.OnAfterUpdate(stage, identification, frontIdentification)
	}
}

func (identification *Identification) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIdentificationDeleteCallback != nil {
		var frontIdentification *Identification
		if front != nil {
			frontIdentification, _ = front.(*Identification)
		}
		stage.OnAfterIdentificationDeleteCallback.OnAfterDelete(stage, identification, frontIdentification)
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

func (instrument *Instrument) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInstrumentCreateCallback != nil {
		stage.OnAfterInstrumentCreateCallback.OnAfterCreate(stage, instrument)
	}
}

func (instrument *Instrument) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentUpdateCallback != nil {
		var frontInstrument *Instrument
		if front != nil {
			frontInstrument, _ = front.(*Instrument)
		}
		stage.OnAfterInstrumentUpdateCallback.OnAfterUpdate(stage, instrument, frontInstrument)
	}
}

func (instrument *Instrument) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentDeleteCallback != nil {
		var frontInstrument *Instrument
		if front != nil {
			frontInstrument, _ = front.(*Instrument)
		}
		stage.OnAfterInstrumentDeleteCallback.OnAfterDelete(stage, instrument, frontInstrument)
	}
}

func (instrument_change *Instrument_change) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInstrument_changeCreateCallback != nil {
		stage.OnAfterInstrument_changeCreateCallback.OnAfterCreate(stage, instrument_change)
	}
}

func (instrument_change *Instrument_change) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrument_changeUpdateCallback != nil {
		var frontInstrument_change *Instrument_change
		if front != nil {
			frontInstrument_change, _ = front.(*Instrument_change)
		}
		stage.OnAfterInstrument_changeUpdateCallback.OnAfterUpdate(stage, instrument_change, frontInstrument_change)
	}
}

func (instrument_change *Instrument_change) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrument_changeDeleteCallback != nil {
		var frontInstrument_change *Instrument_change
		if front != nil {
			frontInstrument_change, _ = front.(*Instrument_change)
		}
		stage.OnAfterInstrument_changeDeleteCallback.OnAfterDelete(stage, instrument_change, frontInstrument_change)
	}
}

func (instrument_link *Instrument_link) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInstrument_linkCreateCallback != nil {
		stage.OnAfterInstrument_linkCreateCallback.OnAfterCreate(stage, instrument_link)
	}
}

func (instrument_link *Instrument_link) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrument_linkUpdateCallback != nil {
		var frontInstrument_link *Instrument_link
		if front != nil {
			frontInstrument_link, _ = front.(*Instrument_link)
		}
		stage.OnAfterInstrument_linkUpdateCallback.OnAfterUpdate(stage, instrument_link, frontInstrument_link)
	}
}

func (instrument_link *Instrument_link) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrument_linkDeleteCallback != nil {
		var frontInstrument_link *Instrument_link
		if front != nil {
			frontInstrument_link, _ = front.(*Instrument_link)
		}
		stage.OnAfterInstrument_linkDeleteCallback.OnAfterDelete(stage, instrument_link, frontInstrument_link)
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

func (key_accidental *Key_accidental) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKey_accidentalCreateCallback != nil {
		stage.OnAfterKey_accidentalCreateCallback.OnAfterCreate(stage, key_accidental)
	}
}

func (key_accidental *Key_accidental) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKey_accidentalUpdateCallback != nil {
		var frontKey_accidental *Key_accidental
		if front != nil {
			frontKey_accidental, _ = front.(*Key_accidental)
		}
		stage.OnAfterKey_accidentalUpdateCallback.OnAfterUpdate(stage, key_accidental, frontKey_accidental)
	}
}

func (key_accidental *Key_accidental) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKey_accidentalDeleteCallback != nil {
		var frontKey_accidental *Key_accidental
		if front != nil {
			frontKey_accidental, _ = front.(*Key_accidental)
		}
		stage.OnAfterKey_accidentalDeleteCallback.OnAfterDelete(stage, key_accidental, frontKey_accidental)
	}
}

func (key_octave *Key_octave) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKey_octaveCreateCallback != nil {
		stage.OnAfterKey_octaveCreateCallback.OnAfterCreate(stage, key_octave)
	}
}

func (key_octave *Key_octave) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKey_octaveUpdateCallback != nil {
		var frontKey_octave *Key_octave
		if front != nil {
			frontKey_octave, _ = front.(*Key_octave)
		}
		stage.OnAfterKey_octaveUpdateCallback.OnAfterUpdate(stage, key_octave, frontKey_octave)
	}
}

func (key_octave *Key_octave) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKey_octaveDeleteCallback != nil {
		var frontKey_octave *Key_octave
		if front != nil {
			frontKey_octave, _ = front.(*Key_octave)
		}
		stage.OnAfterKey_octaveDeleteCallback.OnAfterDelete(stage, key_octave, frontKey_octave)
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

func (level *Level) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLevelCreateCallback != nil {
		stage.OnAfterLevelCreateCallback.OnAfterCreate(stage, level)
	}
}

func (level *Level) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLevelUpdateCallback != nil {
		var frontLevel *Level
		if front != nil {
			frontLevel, _ = front.(*Level)
		}
		stage.OnAfterLevelUpdateCallback.OnAfterUpdate(stage, level, frontLevel)
	}
}

func (level *Level) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLevelDeleteCallback != nil {
		var frontLevel *Level
		if front != nil {
			frontLevel, _ = front.(*Level)
		}
		stage.OnAfterLevelDeleteCallback.OnAfterDelete(stage, level, frontLevel)
	}
}

func (line_detail *Line_detail) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLine_detailCreateCallback != nil {
		stage.OnAfterLine_detailCreateCallback.OnAfterCreate(stage, line_detail)
	}
}

func (line_detail *Line_detail) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLine_detailUpdateCallback != nil {
		var frontLine_detail *Line_detail
		if front != nil {
			frontLine_detail, _ = front.(*Line_detail)
		}
		stage.OnAfterLine_detailUpdateCallback.OnAfterUpdate(stage, line_detail, frontLine_detail)
	}
}

func (line_detail *Line_detail) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLine_detailDeleteCallback != nil {
		var frontLine_detail *Line_detail
		if front != nil {
			frontLine_detail, _ = front.(*Line_detail)
		}
		stage.OnAfterLine_detailDeleteCallback.OnAfterDelete(stage, line_detail, frontLine_detail)
	}
}

func (line_width *Line_width) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLine_widthCreateCallback != nil {
		stage.OnAfterLine_widthCreateCallback.OnAfterCreate(stage, line_width)
	}
}

func (line_width *Line_width) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLine_widthUpdateCallback != nil {
		var frontLine_width *Line_width
		if front != nil {
			frontLine_width, _ = front.(*Line_width)
		}
		stage.OnAfterLine_widthUpdateCallback.OnAfterUpdate(stage, line_width, frontLine_width)
	}
}

func (line_width *Line_width) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLine_widthDeleteCallback != nil {
		var frontLine_width *Line_width
		if front != nil {
			frontLine_width, _ = front.(*Line_width)
		}
		stage.OnAfterLine_widthDeleteCallback.OnAfterDelete(stage, line_width, frontLine_width)
	}
}

func (link *Link) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinkCreateCallback != nil {
		stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, link)
	}
}

func (link *Link) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkUpdateCallback != nil {
		var frontLink *Link
		if front != nil {
			frontLink, _ = front.(*Link)
		}
		stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, link, frontLink)
	}
}

func (link *Link) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkDeleteCallback != nil {
		var frontLink *Link
		if front != nil {
			frontLink, _ = front.(*Link)
		}
		stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, link, frontLink)
	}
}

func (listen *Listen) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterListenCreateCallback != nil {
		stage.OnAfterListenCreateCallback.OnAfterCreate(stage, listen)
	}
}

func (listen *Listen) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterListenUpdateCallback != nil {
		var frontListen *Listen
		if front != nil {
			frontListen, _ = front.(*Listen)
		}
		stage.OnAfterListenUpdateCallback.OnAfterUpdate(stage, listen, frontListen)
	}
}

func (listen *Listen) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterListenDeleteCallback != nil {
		var frontListen *Listen
		if front != nil {
			frontListen, _ = front.(*Listen)
		}
		stage.OnAfterListenDeleteCallback.OnAfterDelete(stage, listen, frontListen)
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

func (lyric *Lyric) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLyricCreateCallback != nil {
		stage.OnAfterLyricCreateCallback.OnAfterCreate(stage, lyric)
	}
}

func (lyric *Lyric) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyricUpdateCallback != nil {
		var frontLyric *Lyric
		if front != nil {
			frontLyric, _ = front.(*Lyric)
		}
		stage.OnAfterLyricUpdateCallback.OnAfterUpdate(stage, lyric, frontLyric)
	}
}

func (lyric *Lyric) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyricDeleteCallback != nil {
		var frontLyric *Lyric
		if front != nil {
			frontLyric, _ = front.(*Lyric)
		}
		stage.OnAfterLyricDeleteCallback.OnAfterDelete(stage, lyric, frontLyric)
	}
}

func (lyric_font *Lyric_font) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLyric_fontCreateCallback != nil {
		stage.OnAfterLyric_fontCreateCallback.OnAfterCreate(stage, lyric_font)
	}
}

func (lyric_font *Lyric_font) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyric_fontUpdateCallback != nil {
		var frontLyric_font *Lyric_font
		if front != nil {
			frontLyric_font, _ = front.(*Lyric_font)
		}
		stage.OnAfterLyric_fontUpdateCallback.OnAfterUpdate(stage, lyric_font, frontLyric_font)
	}
}

func (lyric_font *Lyric_font) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyric_fontDeleteCallback != nil {
		var frontLyric_font *Lyric_font
		if front != nil {
			frontLyric_font, _ = front.(*Lyric_font)
		}
		stage.OnAfterLyric_fontDeleteCallback.OnAfterDelete(stage, lyric_font, frontLyric_font)
	}
}

func (lyric_language *Lyric_language) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLyric_languageCreateCallback != nil {
		stage.OnAfterLyric_languageCreateCallback.OnAfterCreate(stage, lyric_language)
	}
}

func (lyric_language *Lyric_language) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyric_languageUpdateCallback != nil {
		var frontLyric_language *Lyric_language
		if front != nil {
			frontLyric_language, _ = front.(*Lyric_language)
		}
		stage.OnAfterLyric_languageUpdateCallback.OnAfterUpdate(stage, lyric_language, frontLyric_language)
	}
}

func (lyric_language *Lyric_language) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLyric_languageDeleteCallback != nil {
		var frontLyric_language *Lyric_language
		if front != nil {
			frontLyric_language, _ = front.(*Lyric_language)
		}
		stage.OnAfterLyric_languageDeleteCallback.OnAfterDelete(stage, lyric_language, frontLyric_language)
	}
}

func (measure_layout *Measure_layout) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasure_layoutCreateCallback != nil {
		stage.OnAfterMeasure_layoutCreateCallback.OnAfterCreate(stage, measure_layout)
	}
}

func (measure_layout *Measure_layout) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_layoutUpdateCallback != nil {
		var frontMeasure_layout *Measure_layout
		if front != nil {
			frontMeasure_layout, _ = front.(*Measure_layout)
		}
		stage.OnAfterMeasure_layoutUpdateCallback.OnAfterUpdate(stage, measure_layout, frontMeasure_layout)
	}
}

func (measure_layout *Measure_layout) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_layoutDeleteCallback != nil {
		var frontMeasure_layout *Measure_layout
		if front != nil {
			frontMeasure_layout, _ = front.(*Measure_layout)
		}
		stage.OnAfterMeasure_layoutDeleteCallback.OnAfterDelete(stage, measure_layout, frontMeasure_layout)
	}
}

func (measure_numbering *Measure_numbering) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasure_numberingCreateCallback != nil {
		stage.OnAfterMeasure_numberingCreateCallback.OnAfterCreate(stage, measure_numbering)
	}
}

func (measure_numbering *Measure_numbering) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_numberingUpdateCallback != nil {
		var frontMeasure_numbering *Measure_numbering
		if front != nil {
			frontMeasure_numbering, _ = front.(*Measure_numbering)
		}
		stage.OnAfterMeasure_numberingUpdateCallback.OnAfterUpdate(stage, measure_numbering, frontMeasure_numbering)
	}
}

func (measure_numbering *Measure_numbering) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_numberingDeleteCallback != nil {
		var frontMeasure_numbering *Measure_numbering
		if front != nil {
			frontMeasure_numbering, _ = front.(*Measure_numbering)
		}
		stage.OnAfterMeasure_numberingDeleteCallback.OnAfterDelete(stage, measure_numbering, frontMeasure_numbering)
	}
}

func (measure_repeat *Measure_repeat) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasure_repeatCreateCallback != nil {
		stage.OnAfterMeasure_repeatCreateCallback.OnAfterCreate(stage, measure_repeat)
	}
}

func (measure_repeat *Measure_repeat) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_repeatUpdateCallback != nil {
		var frontMeasure_repeat *Measure_repeat
		if front != nil {
			frontMeasure_repeat, _ = front.(*Measure_repeat)
		}
		stage.OnAfterMeasure_repeatUpdateCallback.OnAfterUpdate(stage, measure_repeat, frontMeasure_repeat)
	}
}

func (measure_repeat *Measure_repeat) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_repeatDeleteCallback != nil {
		var frontMeasure_repeat *Measure_repeat
		if front != nil {
			frontMeasure_repeat, _ = front.(*Measure_repeat)
		}
		stage.OnAfterMeasure_repeatDeleteCallback.OnAfterDelete(stage, measure_repeat, frontMeasure_repeat)
	}
}

func (measure_style *Measure_style) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasure_styleCreateCallback != nil {
		stage.OnAfterMeasure_styleCreateCallback.OnAfterCreate(stage, measure_style)
	}
}

func (measure_style *Measure_style) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_styleUpdateCallback != nil {
		var frontMeasure_style *Measure_style
		if front != nil {
			frontMeasure_style, _ = front.(*Measure_style)
		}
		stage.OnAfterMeasure_styleUpdateCallback.OnAfterUpdate(stage, measure_style, frontMeasure_style)
	}
}

func (measure_style *Measure_style) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_styleDeleteCallback != nil {
		var frontMeasure_style *Measure_style
		if front != nil {
			frontMeasure_style, _ = front.(*Measure_style)
		}
		stage.OnAfterMeasure_styleDeleteCallback.OnAfterDelete(stage, measure_style, frontMeasure_style)
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

func (metronome_beam *Metronome_beam) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetronome_beamCreateCallback != nil {
		stage.OnAfterMetronome_beamCreateCallback.OnAfterCreate(stage, metronome_beam)
	}
}

func (metronome_beam *Metronome_beam) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronome_beamUpdateCallback != nil {
		var frontMetronome_beam *Metronome_beam
		if front != nil {
			frontMetronome_beam, _ = front.(*Metronome_beam)
		}
		stage.OnAfterMetronome_beamUpdateCallback.OnAfterUpdate(stage, metronome_beam, frontMetronome_beam)
	}
}

func (metronome_beam *Metronome_beam) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronome_beamDeleteCallback != nil {
		var frontMetronome_beam *Metronome_beam
		if front != nil {
			frontMetronome_beam, _ = front.(*Metronome_beam)
		}
		stage.OnAfterMetronome_beamDeleteCallback.OnAfterDelete(stage, metronome_beam, frontMetronome_beam)
	}
}

func (metronome_note *Metronome_note) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetronome_noteCreateCallback != nil {
		stage.OnAfterMetronome_noteCreateCallback.OnAfterCreate(stage, metronome_note)
	}
}

func (metronome_note *Metronome_note) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronome_noteUpdateCallback != nil {
		var frontMetronome_note *Metronome_note
		if front != nil {
			frontMetronome_note, _ = front.(*Metronome_note)
		}
		stage.OnAfterMetronome_noteUpdateCallback.OnAfterUpdate(stage, metronome_note, frontMetronome_note)
	}
}

func (metronome_note *Metronome_note) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronome_noteDeleteCallback != nil {
		var frontMetronome_note *Metronome_note
		if front != nil {
			frontMetronome_note, _ = front.(*Metronome_note)
		}
		stage.OnAfterMetronome_noteDeleteCallback.OnAfterDelete(stage, metronome_note, frontMetronome_note)
	}
}

func (metronome_tied *Metronome_tied) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetronome_tiedCreateCallback != nil {
		stage.OnAfterMetronome_tiedCreateCallback.OnAfterCreate(stage, metronome_tied)
	}
}

func (metronome_tied *Metronome_tied) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronome_tiedUpdateCallback != nil {
		var frontMetronome_tied *Metronome_tied
		if front != nil {
			frontMetronome_tied, _ = front.(*Metronome_tied)
		}
		stage.OnAfterMetronome_tiedUpdateCallback.OnAfterUpdate(stage, metronome_tied, frontMetronome_tied)
	}
}

func (metronome_tied *Metronome_tied) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronome_tiedDeleteCallback != nil {
		var frontMetronome_tied *Metronome_tied
		if front != nil {
			frontMetronome_tied, _ = front.(*Metronome_tied)
		}
		stage.OnAfterMetronome_tiedDeleteCallback.OnAfterDelete(stage, metronome_tied, frontMetronome_tied)
	}
}

func (metronome_tuplet *Metronome_tuplet) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetronome_tupletCreateCallback != nil {
		stage.OnAfterMetronome_tupletCreateCallback.OnAfterCreate(stage, metronome_tuplet)
	}
}

func (metronome_tuplet *Metronome_tuplet) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronome_tupletUpdateCallback != nil {
		var frontMetronome_tuplet *Metronome_tuplet
		if front != nil {
			frontMetronome_tuplet, _ = front.(*Metronome_tuplet)
		}
		stage.OnAfterMetronome_tupletUpdateCallback.OnAfterUpdate(stage, metronome_tuplet, frontMetronome_tuplet)
	}
}

func (metronome_tuplet *Metronome_tuplet) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetronome_tupletDeleteCallback != nil {
		var frontMetronome_tuplet *Metronome_tuplet
		if front != nil {
			frontMetronome_tuplet, _ = front.(*Metronome_tuplet)
		}
		stage.OnAfterMetronome_tupletDeleteCallback.OnAfterDelete(stage, metronome_tuplet, frontMetronome_tuplet)
	}
}

func (midi_device *Midi_device) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMidi_deviceCreateCallback != nil {
		stage.OnAfterMidi_deviceCreateCallback.OnAfterCreate(stage, midi_device)
	}
}

func (midi_device *Midi_device) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidi_deviceUpdateCallback != nil {
		var frontMidi_device *Midi_device
		if front != nil {
			frontMidi_device, _ = front.(*Midi_device)
		}
		stage.OnAfterMidi_deviceUpdateCallback.OnAfterUpdate(stage, midi_device, frontMidi_device)
	}
}

func (midi_device *Midi_device) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidi_deviceDeleteCallback != nil {
		var frontMidi_device *Midi_device
		if front != nil {
			frontMidi_device, _ = front.(*Midi_device)
		}
		stage.OnAfterMidi_deviceDeleteCallback.OnAfterDelete(stage, midi_device, frontMidi_device)
	}
}

func (midi_instrument *Midi_instrument) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMidi_instrumentCreateCallback != nil {
		stage.OnAfterMidi_instrumentCreateCallback.OnAfterCreate(stage, midi_instrument)
	}
}

func (midi_instrument *Midi_instrument) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidi_instrumentUpdateCallback != nil {
		var frontMidi_instrument *Midi_instrument
		if front != nil {
			frontMidi_instrument, _ = front.(*Midi_instrument)
		}
		stage.OnAfterMidi_instrumentUpdateCallback.OnAfterUpdate(stage, midi_instrument, frontMidi_instrument)
	}
}

func (midi_instrument *Midi_instrument) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidi_instrumentDeleteCallback != nil {
		var frontMidi_instrument *Midi_instrument
		if front != nil {
			frontMidi_instrument, _ = front.(*Midi_instrument)
		}
		stage.OnAfterMidi_instrumentDeleteCallback.OnAfterDelete(stage, midi_instrument, frontMidi_instrument)
	}
}

func (miscellaneous *Miscellaneous) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMiscellaneousCreateCallback != nil {
		stage.OnAfterMiscellaneousCreateCallback.OnAfterCreate(stage, miscellaneous)
	}
}

func (miscellaneous *Miscellaneous) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMiscellaneousUpdateCallback != nil {
		var frontMiscellaneous *Miscellaneous
		if front != nil {
			frontMiscellaneous, _ = front.(*Miscellaneous)
		}
		stage.OnAfterMiscellaneousUpdateCallback.OnAfterUpdate(stage, miscellaneous, frontMiscellaneous)
	}
}

func (miscellaneous *Miscellaneous) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMiscellaneousDeleteCallback != nil {
		var frontMiscellaneous *Miscellaneous
		if front != nil {
			frontMiscellaneous, _ = front.(*Miscellaneous)
		}
		stage.OnAfterMiscellaneousDeleteCallback.OnAfterDelete(stage, miscellaneous, frontMiscellaneous)
	}
}

func (miscellaneous_field *Miscellaneous_field) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMiscellaneous_fieldCreateCallback != nil {
		stage.OnAfterMiscellaneous_fieldCreateCallback.OnAfterCreate(stage, miscellaneous_field)
	}
}

func (miscellaneous_field *Miscellaneous_field) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMiscellaneous_fieldUpdateCallback != nil {
		var frontMiscellaneous_field *Miscellaneous_field
		if front != nil {
			frontMiscellaneous_field, _ = front.(*Miscellaneous_field)
		}
		stage.OnAfterMiscellaneous_fieldUpdateCallback.OnAfterUpdate(stage, miscellaneous_field, frontMiscellaneous_field)
	}
}

func (miscellaneous_field *Miscellaneous_field) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMiscellaneous_fieldDeleteCallback != nil {
		var frontMiscellaneous_field *Miscellaneous_field
		if front != nil {
			frontMiscellaneous_field, _ = front.(*Miscellaneous_field)
		}
		stage.OnAfterMiscellaneous_fieldDeleteCallback.OnAfterDelete(stage, miscellaneous_field, frontMiscellaneous_field)
	}
}

func (mordent *Mordent) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMordentCreateCallback != nil {
		stage.OnAfterMordentCreateCallback.OnAfterCreate(stage, mordent)
	}
}

func (mordent *Mordent) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMordentUpdateCallback != nil {
		var frontMordent *Mordent
		if front != nil {
			frontMordent, _ = front.(*Mordent)
		}
		stage.OnAfterMordentUpdateCallback.OnAfterUpdate(stage, mordent, frontMordent)
	}
}

func (mordent *Mordent) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMordentDeleteCallback != nil {
		var frontMordent *Mordent
		if front != nil {
			frontMordent, _ = front.(*Mordent)
		}
		stage.OnAfterMordentDeleteCallback.OnAfterDelete(stage, mordent, frontMordent)
	}
}

func (multiple_rest *Multiple_rest) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiple_restCreateCallback != nil {
		stage.OnAfterMultiple_restCreateCallback.OnAfterCreate(stage, multiple_rest)
	}
}

func (multiple_rest *Multiple_rest) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiple_restUpdateCallback != nil {
		var frontMultiple_rest *Multiple_rest
		if front != nil {
			frontMultiple_rest, _ = front.(*Multiple_rest)
		}
		stage.OnAfterMultiple_restUpdateCallback.OnAfterUpdate(stage, multiple_rest, frontMultiple_rest)
	}
}

func (multiple_rest *Multiple_rest) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiple_restDeleteCallback != nil {
		var frontMultiple_rest *Multiple_rest
		if front != nil {
			frontMultiple_rest, _ = front.(*Multiple_rest)
		}
		stage.OnAfterMultiple_restDeleteCallback.OnAfterDelete(stage, multiple_rest, frontMultiple_rest)
	}
}

func (name_display *Name_display) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterName_displayCreateCallback != nil {
		stage.OnAfterName_displayCreateCallback.OnAfterCreate(stage, name_display)
	}
}

func (name_display *Name_display) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterName_displayUpdateCallback != nil {
		var frontName_display *Name_display
		if front != nil {
			frontName_display, _ = front.(*Name_display)
		}
		stage.OnAfterName_displayUpdateCallback.OnAfterUpdate(stage, name_display, frontName_display)
	}
}

func (name_display *Name_display) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterName_displayDeleteCallback != nil {
		var frontName_display *Name_display
		if front != nil {
			frontName_display, _ = front.(*Name_display)
		}
		stage.OnAfterName_displayDeleteCallback.OnAfterDelete(stage, name_display, frontName_display)
	}
}

func (non_arpeggiate *Non_arpeggiate) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNon_arpeggiateCreateCallback != nil {
		stage.OnAfterNon_arpeggiateCreateCallback.OnAfterCreate(stage, non_arpeggiate)
	}
}

func (non_arpeggiate *Non_arpeggiate) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNon_arpeggiateUpdateCallback != nil {
		var frontNon_arpeggiate *Non_arpeggiate
		if front != nil {
			frontNon_arpeggiate, _ = front.(*Non_arpeggiate)
		}
		stage.OnAfterNon_arpeggiateUpdateCallback.OnAfterUpdate(stage, non_arpeggiate, frontNon_arpeggiate)
	}
}

func (non_arpeggiate *Non_arpeggiate) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNon_arpeggiateDeleteCallback != nil {
		var frontNon_arpeggiate *Non_arpeggiate
		if front != nil {
			frontNon_arpeggiate, _ = front.(*Non_arpeggiate)
		}
		stage.OnAfterNon_arpeggiateDeleteCallback.OnAfterDelete(stage, non_arpeggiate, frontNon_arpeggiate)
	}
}

func (notations *Notations) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNotationsCreateCallback != nil {
		stage.OnAfterNotationsCreateCallback.OnAfterCreate(stage, notations)
	}
}

func (notations *Notations) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotationsUpdateCallback != nil {
		var frontNotations *Notations
		if front != nil {
			frontNotations, _ = front.(*Notations)
		}
		stage.OnAfterNotationsUpdateCallback.OnAfterUpdate(stage, notations, frontNotations)
	}
}

func (notations *Notations) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotationsDeleteCallback != nil {
		var frontNotations *Notations
		if front != nil {
			frontNotations, _ = front.(*Notations)
		}
		stage.OnAfterNotationsDeleteCallback.OnAfterDelete(stage, notations, frontNotations)
	}
}

func (note *Note) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteCreateCallback != nil {
		stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, note)
	}
}

func (note *Note) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteUpdateCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, note, frontNote)
	}
}

func (note *Note) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteDeleteCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, note, frontNote)
	}
}

func (note_size *Note_size) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNote_sizeCreateCallback != nil {
		stage.OnAfterNote_sizeCreateCallback.OnAfterCreate(stage, note_size)
	}
}

func (note_size *Note_size) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNote_sizeUpdateCallback != nil {
		var frontNote_size *Note_size
		if front != nil {
			frontNote_size, _ = front.(*Note_size)
		}
		stage.OnAfterNote_sizeUpdateCallback.OnAfterUpdate(stage, note_size, frontNote_size)
	}
}

func (note_size *Note_size) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNote_sizeDeleteCallback != nil {
		var frontNote_size *Note_size
		if front != nil {
			frontNote_size, _ = front.(*Note_size)
		}
		stage.OnAfterNote_sizeDeleteCallback.OnAfterDelete(stage, note_size, frontNote_size)
	}
}

func (note_type *Note_type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNote_typeCreateCallback != nil {
		stage.OnAfterNote_typeCreateCallback.OnAfterCreate(stage, note_type)
	}
}

func (note_type *Note_type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNote_typeUpdateCallback != nil {
		var frontNote_type *Note_type
		if front != nil {
			frontNote_type, _ = front.(*Note_type)
		}
		stage.OnAfterNote_typeUpdateCallback.OnAfterUpdate(stage, note_type, frontNote_type)
	}
}

func (note_type *Note_type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNote_typeDeleteCallback != nil {
		var frontNote_type *Note_type
		if front != nil {
			frontNote_type, _ = front.(*Note_type)
		}
		stage.OnAfterNote_typeDeleteCallback.OnAfterDelete(stage, note_type, frontNote_type)
	}
}

func (notehead *Notehead) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteheadCreateCallback != nil {
		stage.OnAfterNoteheadCreateCallback.OnAfterCreate(stage, notehead)
	}
}

func (notehead *Notehead) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteheadUpdateCallback != nil {
		var frontNotehead *Notehead
		if front != nil {
			frontNotehead, _ = front.(*Notehead)
		}
		stage.OnAfterNoteheadUpdateCallback.OnAfterUpdate(stage, notehead, frontNotehead)
	}
}

func (notehead *Notehead) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteheadDeleteCallback != nil {
		var frontNotehead *Notehead
		if front != nil {
			frontNotehead, _ = front.(*Notehead)
		}
		stage.OnAfterNoteheadDeleteCallback.OnAfterDelete(stage, notehead, frontNotehead)
	}
}

func (notehead_text *Notehead_text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNotehead_textCreateCallback != nil {
		stage.OnAfterNotehead_textCreateCallback.OnAfterCreate(stage, notehead_text)
	}
}

func (notehead_text *Notehead_text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotehead_textUpdateCallback != nil {
		var frontNotehead_text *Notehead_text
		if front != nil {
			frontNotehead_text, _ = front.(*Notehead_text)
		}
		stage.OnAfterNotehead_textUpdateCallback.OnAfterUpdate(stage, notehead_text, frontNotehead_text)
	}
}

func (notehead_text *Notehead_text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotehead_textDeleteCallback != nil {
		var frontNotehead_text *Notehead_text
		if front != nil {
			frontNotehead_text, _ = front.(*Notehead_text)
		}
		stage.OnAfterNotehead_textDeleteCallback.OnAfterDelete(stage, notehead_text, frontNotehead_text)
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

func (numeral_key *Numeral_key) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNumeral_keyCreateCallback != nil {
		stage.OnAfterNumeral_keyCreateCallback.OnAfterCreate(stage, numeral_key)
	}
}

func (numeral_key *Numeral_key) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeral_keyUpdateCallback != nil {
		var frontNumeral_key *Numeral_key
		if front != nil {
			frontNumeral_key, _ = front.(*Numeral_key)
		}
		stage.OnAfterNumeral_keyUpdateCallback.OnAfterUpdate(stage, numeral_key, frontNumeral_key)
	}
}

func (numeral_key *Numeral_key) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeral_keyDeleteCallback != nil {
		var frontNumeral_key *Numeral_key
		if front != nil {
			frontNumeral_key, _ = front.(*Numeral_key)
		}
		stage.OnAfterNumeral_keyDeleteCallback.OnAfterDelete(stage, numeral_key, frontNumeral_key)
	}
}

func (numeral_root *Numeral_root) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNumeral_rootCreateCallback != nil {
		stage.OnAfterNumeral_rootCreateCallback.OnAfterCreate(stage, numeral_root)
	}
}

func (numeral_root *Numeral_root) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeral_rootUpdateCallback != nil {
		var frontNumeral_root *Numeral_root
		if front != nil {
			frontNumeral_root, _ = front.(*Numeral_root)
		}
		stage.OnAfterNumeral_rootUpdateCallback.OnAfterUpdate(stage, numeral_root, frontNumeral_root)
	}
}

func (numeral_root *Numeral_root) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumeral_rootDeleteCallback != nil {
		var frontNumeral_root *Numeral_root
		if front != nil {
			frontNumeral_root, _ = front.(*Numeral_root)
		}
		stage.OnAfterNumeral_rootDeleteCallback.OnAfterDelete(stage, numeral_root, frontNumeral_root)
	}
}

func (octave_shift *Octave_shift) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOctave_shiftCreateCallback != nil {
		stage.OnAfterOctave_shiftCreateCallback.OnAfterCreate(stage, octave_shift)
	}
}

func (octave_shift *Octave_shift) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOctave_shiftUpdateCallback != nil {
		var frontOctave_shift *Octave_shift
		if front != nil {
			frontOctave_shift, _ = front.(*Octave_shift)
		}
		stage.OnAfterOctave_shiftUpdateCallback.OnAfterUpdate(stage, octave_shift, frontOctave_shift)
	}
}

func (octave_shift *Octave_shift) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOctave_shiftDeleteCallback != nil {
		var frontOctave_shift *Octave_shift
		if front != nil {
			frontOctave_shift, _ = front.(*Octave_shift)
		}
		stage.OnAfterOctave_shiftDeleteCallback.OnAfterDelete(stage, octave_shift, frontOctave_shift)
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

func (opus *Opus) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOpusCreateCallback != nil {
		stage.OnAfterOpusCreateCallback.OnAfterCreate(stage, opus)
	}
}

func (opus *Opus) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOpusUpdateCallback != nil {
		var frontOpus *Opus
		if front != nil {
			frontOpus, _ = front.(*Opus)
		}
		stage.OnAfterOpusUpdateCallback.OnAfterUpdate(stage, opus, frontOpus)
	}
}

func (opus *Opus) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOpusDeleteCallback != nil {
		var frontOpus *Opus
		if front != nil {
			frontOpus, _ = front.(*Opus)
		}
		stage.OnAfterOpusDeleteCallback.OnAfterDelete(stage, opus, frontOpus)
	}
}

func (ornaments *Ornaments) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrnamentsCreateCallback != nil {
		stage.OnAfterOrnamentsCreateCallback.OnAfterCreate(stage, ornaments)
	}
}

func (ornaments *Ornaments) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrnamentsUpdateCallback != nil {
		var frontOrnaments *Ornaments
		if front != nil {
			frontOrnaments, _ = front.(*Ornaments)
		}
		stage.OnAfterOrnamentsUpdateCallback.OnAfterUpdate(stage, ornaments, frontOrnaments)
	}
}

func (ornaments *Ornaments) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrnamentsDeleteCallback != nil {
		var frontOrnaments *Ornaments
		if front != nil {
			frontOrnaments, _ = front.(*Ornaments)
		}
		stage.OnAfterOrnamentsDeleteCallback.OnAfterDelete(stage, ornaments, frontOrnaments)
	}
}

func (other_appearance *Other_appearance) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOther_appearanceCreateCallback != nil {
		stage.OnAfterOther_appearanceCreateCallback.OnAfterCreate(stage, other_appearance)
	}
}

func (other_appearance *Other_appearance) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_appearanceUpdateCallback != nil {
		var frontOther_appearance *Other_appearance
		if front != nil {
			frontOther_appearance, _ = front.(*Other_appearance)
		}
		stage.OnAfterOther_appearanceUpdateCallback.OnAfterUpdate(stage, other_appearance, frontOther_appearance)
	}
}

func (other_appearance *Other_appearance) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_appearanceDeleteCallback != nil {
		var frontOther_appearance *Other_appearance
		if front != nil {
			frontOther_appearance, _ = front.(*Other_appearance)
		}
		stage.OnAfterOther_appearanceDeleteCallback.OnAfterDelete(stage, other_appearance, frontOther_appearance)
	}
}

func (other_direction *Other_direction) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOther_directionCreateCallback != nil {
		stage.OnAfterOther_directionCreateCallback.OnAfterCreate(stage, other_direction)
	}
}

func (other_direction *Other_direction) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_directionUpdateCallback != nil {
		var frontOther_direction *Other_direction
		if front != nil {
			frontOther_direction, _ = front.(*Other_direction)
		}
		stage.OnAfterOther_directionUpdateCallback.OnAfterUpdate(stage, other_direction, frontOther_direction)
	}
}

func (other_direction *Other_direction) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_directionDeleteCallback != nil {
		var frontOther_direction *Other_direction
		if front != nil {
			frontOther_direction, _ = front.(*Other_direction)
		}
		stage.OnAfterOther_directionDeleteCallback.OnAfterDelete(stage, other_direction, frontOther_direction)
	}
}

func (other_listening *Other_listening) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOther_listeningCreateCallback != nil {
		stage.OnAfterOther_listeningCreateCallback.OnAfterCreate(stage, other_listening)
	}
}

func (other_listening *Other_listening) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_listeningUpdateCallback != nil {
		var frontOther_listening *Other_listening
		if front != nil {
			frontOther_listening, _ = front.(*Other_listening)
		}
		stage.OnAfterOther_listeningUpdateCallback.OnAfterUpdate(stage, other_listening, frontOther_listening)
	}
}

func (other_listening *Other_listening) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_listeningDeleteCallback != nil {
		var frontOther_listening *Other_listening
		if front != nil {
			frontOther_listening, _ = front.(*Other_listening)
		}
		stage.OnAfterOther_listeningDeleteCallback.OnAfterDelete(stage, other_listening, frontOther_listening)
	}
}

func (other_notation *Other_notation) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOther_notationCreateCallback != nil {
		stage.OnAfterOther_notationCreateCallback.OnAfterCreate(stage, other_notation)
	}
}

func (other_notation *Other_notation) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_notationUpdateCallback != nil {
		var frontOther_notation *Other_notation
		if front != nil {
			frontOther_notation, _ = front.(*Other_notation)
		}
		stage.OnAfterOther_notationUpdateCallback.OnAfterUpdate(stage, other_notation, frontOther_notation)
	}
}

func (other_notation *Other_notation) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_notationDeleteCallback != nil {
		var frontOther_notation *Other_notation
		if front != nil {
			frontOther_notation, _ = front.(*Other_notation)
		}
		stage.OnAfterOther_notationDeleteCallback.OnAfterDelete(stage, other_notation, frontOther_notation)
	}
}

func (other_placement_text *Other_placement_text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOther_placement_textCreateCallback != nil {
		stage.OnAfterOther_placement_textCreateCallback.OnAfterCreate(stage, other_placement_text)
	}
}

func (other_placement_text *Other_placement_text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_placement_textUpdateCallback != nil {
		var frontOther_placement_text *Other_placement_text
		if front != nil {
			frontOther_placement_text, _ = front.(*Other_placement_text)
		}
		stage.OnAfterOther_placement_textUpdateCallback.OnAfterUpdate(stage, other_placement_text, frontOther_placement_text)
	}
}

func (other_placement_text *Other_placement_text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_placement_textDeleteCallback != nil {
		var frontOther_placement_text *Other_placement_text
		if front != nil {
			frontOther_placement_text, _ = front.(*Other_placement_text)
		}
		stage.OnAfterOther_placement_textDeleteCallback.OnAfterDelete(stage, other_placement_text, frontOther_placement_text)
	}
}

func (other_play *Other_play) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOther_playCreateCallback != nil {
		stage.OnAfterOther_playCreateCallback.OnAfterCreate(stage, other_play)
	}
}

func (other_play *Other_play) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_playUpdateCallback != nil {
		var frontOther_play *Other_play
		if front != nil {
			frontOther_play, _ = front.(*Other_play)
		}
		stage.OnAfterOther_playUpdateCallback.OnAfterUpdate(stage, other_play, frontOther_play)
	}
}

func (other_play *Other_play) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_playDeleteCallback != nil {
		var frontOther_play *Other_play
		if front != nil {
			frontOther_play, _ = front.(*Other_play)
		}
		stage.OnAfterOther_playDeleteCallback.OnAfterDelete(stage, other_play, frontOther_play)
	}
}

func (other_text *Other_text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOther_textCreateCallback != nil {
		stage.OnAfterOther_textCreateCallback.OnAfterCreate(stage, other_text)
	}
}

func (other_text *Other_text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_textUpdateCallback != nil {
		var frontOther_text *Other_text
		if front != nil {
			frontOther_text, _ = front.(*Other_text)
		}
		stage.OnAfterOther_textUpdateCallback.OnAfterUpdate(stage, other_text, frontOther_text)
	}
}

func (other_text *Other_text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOther_textDeleteCallback != nil {
		var frontOther_text *Other_text
		if front != nil {
			frontOther_text, _ = front.(*Other_text)
		}
		stage.OnAfterOther_textDeleteCallback.OnAfterDelete(stage, other_text, frontOther_text)
	}
}

func (page_layout *Page_layout) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPage_layoutCreateCallback != nil {
		stage.OnAfterPage_layoutCreateCallback.OnAfterCreate(stage, page_layout)
	}
}

func (page_layout *Page_layout) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPage_layoutUpdateCallback != nil {
		var frontPage_layout *Page_layout
		if front != nil {
			frontPage_layout, _ = front.(*Page_layout)
		}
		stage.OnAfterPage_layoutUpdateCallback.OnAfterUpdate(stage, page_layout, frontPage_layout)
	}
}

func (page_layout *Page_layout) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPage_layoutDeleteCallback != nil {
		var frontPage_layout *Page_layout
		if front != nil {
			frontPage_layout, _ = front.(*Page_layout)
		}
		stage.OnAfterPage_layoutDeleteCallback.OnAfterDelete(stage, page_layout, frontPage_layout)
	}
}

func (page_margins *Page_margins) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPage_marginsCreateCallback != nil {
		stage.OnAfterPage_marginsCreateCallback.OnAfterCreate(stage, page_margins)
	}
}

func (page_margins *Page_margins) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPage_marginsUpdateCallback != nil {
		var frontPage_margins *Page_margins
		if front != nil {
			frontPage_margins, _ = front.(*Page_margins)
		}
		stage.OnAfterPage_marginsUpdateCallback.OnAfterUpdate(stage, page_margins, frontPage_margins)
	}
}

func (page_margins *Page_margins) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPage_marginsDeleteCallback != nil {
		var frontPage_margins *Page_margins
		if front != nil {
			frontPage_margins, _ = front.(*Page_margins)
		}
		stage.OnAfterPage_marginsDeleteCallback.OnAfterDelete(stage, page_margins, frontPage_margins)
	}
}

func (part_clef *Part_clef) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPart_clefCreateCallback != nil {
		stage.OnAfterPart_clefCreateCallback.OnAfterCreate(stage, part_clef)
	}
}

func (part_clef *Part_clef) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_clefUpdateCallback != nil {
		var frontPart_clef *Part_clef
		if front != nil {
			frontPart_clef, _ = front.(*Part_clef)
		}
		stage.OnAfterPart_clefUpdateCallback.OnAfterUpdate(stage, part_clef, frontPart_clef)
	}
}

func (part_clef *Part_clef) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_clefDeleteCallback != nil {
		var frontPart_clef *Part_clef
		if front != nil {
			frontPart_clef, _ = front.(*Part_clef)
		}
		stage.OnAfterPart_clefDeleteCallback.OnAfterDelete(stage, part_clef, frontPart_clef)
	}
}

func (part_group *Part_group) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPart_groupCreateCallback != nil {
		stage.OnAfterPart_groupCreateCallback.OnAfterCreate(stage, part_group)
	}
}

func (part_group *Part_group) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_groupUpdateCallback != nil {
		var frontPart_group *Part_group
		if front != nil {
			frontPart_group, _ = front.(*Part_group)
		}
		stage.OnAfterPart_groupUpdateCallback.OnAfterUpdate(stage, part_group, frontPart_group)
	}
}

func (part_group *Part_group) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_groupDeleteCallback != nil {
		var frontPart_group *Part_group
		if front != nil {
			frontPart_group, _ = front.(*Part_group)
		}
		stage.OnAfterPart_groupDeleteCallback.OnAfterDelete(stage, part_group, frontPart_group)
	}
}

func (part_link *Part_link) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPart_linkCreateCallback != nil {
		stage.OnAfterPart_linkCreateCallback.OnAfterCreate(stage, part_link)
	}
}

func (part_link *Part_link) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_linkUpdateCallback != nil {
		var frontPart_link *Part_link
		if front != nil {
			frontPart_link, _ = front.(*Part_link)
		}
		stage.OnAfterPart_linkUpdateCallback.OnAfterUpdate(stage, part_link, frontPart_link)
	}
}

func (part_link *Part_link) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_linkDeleteCallback != nil {
		var frontPart_link *Part_link
		if front != nil {
			frontPart_link, _ = front.(*Part_link)
		}
		stage.OnAfterPart_linkDeleteCallback.OnAfterDelete(stage, part_link, frontPart_link)
	}
}

func (part_list *Part_list) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPart_listCreateCallback != nil {
		stage.OnAfterPart_listCreateCallback.OnAfterCreate(stage, part_list)
	}
}

func (part_list *Part_list) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_listUpdateCallback != nil {
		var frontPart_list *Part_list
		if front != nil {
			frontPart_list, _ = front.(*Part_list)
		}
		stage.OnAfterPart_listUpdateCallback.OnAfterUpdate(stage, part_list, frontPart_list)
	}
}

func (part_list *Part_list) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_listDeleteCallback != nil {
		var frontPart_list *Part_list
		if front != nil {
			frontPart_list, _ = front.(*Part_list)
		}
		stage.OnAfterPart_listDeleteCallback.OnAfterDelete(stage, part_list, frontPart_list)
	}
}

func (part_name *Part_name) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPart_nameCreateCallback != nil {
		stage.OnAfterPart_nameCreateCallback.OnAfterCreate(stage, part_name)
	}
}

func (part_name *Part_name) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_nameUpdateCallback != nil {
		var frontPart_name *Part_name
		if front != nil {
			frontPart_name, _ = front.(*Part_name)
		}
		stage.OnAfterPart_nameUpdateCallback.OnAfterUpdate(stage, part_name, frontPart_name)
	}
}

func (part_name *Part_name) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_nameDeleteCallback != nil {
		var frontPart_name *Part_name
		if front != nil {
			frontPart_name, _ = front.(*Part_name)
		}
		stage.OnAfterPart_nameDeleteCallback.OnAfterDelete(stage, part_name, frontPart_name)
	}
}

func (part_symbol *Part_symbol) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPart_symbolCreateCallback != nil {
		stage.OnAfterPart_symbolCreateCallback.OnAfterCreate(stage, part_symbol)
	}
}

func (part_symbol *Part_symbol) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_symbolUpdateCallback != nil {
		var frontPart_symbol *Part_symbol
		if front != nil {
			frontPart_symbol, _ = front.(*Part_symbol)
		}
		stage.OnAfterPart_symbolUpdateCallback.OnAfterUpdate(stage, part_symbol, frontPart_symbol)
	}
}

func (part_symbol *Part_symbol) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_symbolDeleteCallback != nil {
		var frontPart_symbol *Part_symbol
		if front != nil {
			frontPart_symbol, _ = front.(*Part_symbol)
		}
		stage.OnAfterPart_symbolDeleteCallback.OnAfterDelete(stage, part_symbol, frontPart_symbol)
	}
}

func (part_transpose *Part_transpose) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPart_transposeCreateCallback != nil {
		stage.OnAfterPart_transposeCreateCallback.OnAfterCreate(stage, part_transpose)
	}
}

func (part_transpose *Part_transpose) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_transposeUpdateCallback != nil {
		var frontPart_transpose *Part_transpose
		if front != nil {
			frontPart_transpose, _ = front.(*Part_transpose)
		}
		stage.OnAfterPart_transposeUpdateCallback.OnAfterUpdate(stage, part_transpose, frontPart_transpose)
	}
}

func (part_transpose *Part_transpose) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPart_transposeDeleteCallback != nil {
		var frontPart_transpose *Part_transpose
		if front != nil {
			frontPart_transpose, _ = front.(*Part_transpose)
		}
		stage.OnAfterPart_transposeDeleteCallback.OnAfterDelete(stage, part_transpose, frontPart_transpose)
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

func (pedal_tuning *Pedal_tuning) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPedal_tuningCreateCallback != nil {
		stage.OnAfterPedal_tuningCreateCallback.OnAfterCreate(stage, pedal_tuning)
	}
}

func (pedal_tuning *Pedal_tuning) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPedal_tuningUpdateCallback != nil {
		var frontPedal_tuning *Pedal_tuning
		if front != nil {
			frontPedal_tuning, _ = front.(*Pedal_tuning)
		}
		stage.OnAfterPedal_tuningUpdateCallback.OnAfterUpdate(stage, pedal_tuning, frontPedal_tuning)
	}
}

func (pedal_tuning *Pedal_tuning) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPedal_tuningDeleteCallback != nil {
		var frontPedal_tuning *Pedal_tuning
		if front != nil {
			frontPedal_tuning, _ = front.(*Pedal_tuning)
		}
		stage.OnAfterPedal_tuningDeleteCallback.OnAfterDelete(stage, pedal_tuning, frontPedal_tuning)
	}
}

func (per_minute *Per_minute) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPer_minuteCreateCallback != nil {
		stage.OnAfterPer_minuteCreateCallback.OnAfterCreate(stage, per_minute)
	}
}

func (per_minute *Per_minute) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPer_minuteUpdateCallback != nil {
		var frontPer_minute *Per_minute
		if front != nil {
			frontPer_minute, _ = front.(*Per_minute)
		}
		stage.OnAfterPer_minuteUpdateCallback.OnAfterUpdate(stage, per_minute, frontPer_minute)
	}
}

func (per_minute *Per_minute) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPer_minuteDeleteCallback != nil {
		var frontPer_minute *Per_minute
		if front != nil {
			frontPer_minute, _ = front.(*Per_minute)
		}
		stage.OnAfterPer_minuteDeleteCallback.OnAfterDelete(stage, per_minute, frontPer_minute)
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

func (pitch *Pitch) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPitchCreateCallback != nil {
		stage.OnAfterPitchCreateCallback.OnAfterCreate(stage, pitch)
	}
}

func (pitch *Pitch) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPitchUpdateCallback != nil {
		var frontPitch *Pitch
		if front != nil {
			frontPitch, _ = front.(*Pitch)
		}
		stage.OnAfterPitchUpdateCallback.OnAfterUpdate(stage, pitch, frontPitch)
	}
}

func (pitch *Pitch) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPitchDeleteCallback != nil {
		var frontPitch *Pitch
		if front != nil {
			frontPitch, _ = front.(*Pitch)
		}
		stage.OnAfterPitchDeleteCallback.OnAfterDelete(stage, pitch, frontPitch)
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

func (placement_text *Placement_text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlacement_textCreateCallback != nil {
		stage.OnAfterPlacement_textCreateCallback.OnAfterCreate(stage, placement_text)
	}
}

func (placement_text *Placement_text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlacement_textUpdateCallback != nil {
		var frontPlacement_text *Placement_text
		if front != nil {
			frontPlacement_text, _ = front.(*Placement_text)
		}
		stage.OnAfterPlacement_textUpdateCallback.OnAfterUpdate(stage, placement_text, frontPlacement_text)
	}
}

func (placement_text *Placement_text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlacement_textDeleteCallback != nil {
		var frontPlacement_text *Placement_text
		if front != nil {
			frontPlacement_text, _ = front.(*Placement_text)
		}
		stage.OnAfterPlacement_textDeleteCallback.OnAfterDelete(stage, placement_text, frontPlacement_text)
	}
}

func (play *Play) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlayCreateCallback != nil {
		stage.OnAfterPlayCreateCallback.OnAfterCreate(stage, play)
	}
}

func (play *Play) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlayUpdateCallback != nil {
		var frontPlay *Play
		if front != nil {
			frontPlay, _ = front.(*Play)
		}
		stage.OnAfterPlayUpdateCallback.OnAfterUpdate(stage, play, frontPlay)
	}
}

func (play *Play) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlayDeleteCallback != nil {
		var frontPlay *Play
		if front != nil {
			frontPlay, _ = front.(*Play)
		}
		stage.OnAfterPlayDeleteCallback.OnAfterDelete(stage, play, frontPlay)
	}
}

func (player *Player) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlayerCreateCallback != nil {
		stage.OnAfterPlayerCreateCallback.OnAfterCreate(stage, player)
	}
}

func (player *Player) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlayerUpdateCallback != nil {
		var frontPlayer *Player
		if front != nil {
			frontPlayer, _ = front.(*Player)
		}
		stage.OnAfterPlayerUpdateCallback.OnAfterUpdate(stage, player, frontPlayer)
	}
}

func (player *Player) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlayerDeleteCallback != nil {
		var frontPlayer *Player
		if front != nil {
			frontPlayer, _ = front.(*Player)
		}
		stage.OnAfterPlayerDeleteCallback.OnAfterDelete(stage, player, frontPlayer)
	}
}

func (principal_voice *Principal_voice) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrincipal_voiceCreateCallback != nil {
		stage.OnAfterPrincipal_voiceCreateCallback.OnAfterCreate(stage, principal_voice)
	}
}

func (principal_voice *Principal_voice) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrincipal_voiceUpdateCallback != nil {
		var frontPrincipal_voice *Principal_voice
		if front != nil {
			frontPrincipal_voice, _ = front.(*Principal_voice)
		}
		stage.OnAfterPrincipal_voiceUpdateCallback.OnAfterUpdate(stage, principal_voice, frontPrincipal_voice)
	}
}

func (principal_voice *Principal_voice) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrincipal_voiceDeleteCallback != nil {
		var frontPrincipal_voice *Principal_voice
		if front != nil {
			frontPrincipal_voice, _ = front.(*Principal_voice)
		}
		stage.OnAfterPrincipal_voiceDeleteCallback.OnAfterDelete(stage, principal_voice, frontPrincipal_voice)
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

func (release *Release) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterReleaseCreateCallback != nil {
		stage.OnAfterReleaseCreateCallback.OnAfterCreate(stage, release)
	}
}

func (release *Release) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReleaseUpdateCallback != nil {
		var frontRelease *Release
		if front != nil {
			frontRelease, _ = front.(*Release)
		}
		stage.OnAfterReleaseUpdateCallback.OnAfterUpdate(stage, release, frontRelease)
	}
}

func (release *Release) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReleaseDeleteCallback != nil {
		var frontRelease *Release
		if front != nil {
			frontRelease, _ = front.(*Release)
		}
		stage.OnAfterReleaseDeleteCallback.OnAfterDelete(stage, release, frontRelease)
	}
}

func (repeat *Repeat) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRepeatCreateCallback != nil {
		stage.OnAfterRepeatCreateCallback.OnAfterCreate(stage, repeat)
	}
}

func (repeat *Repeat) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRepeatUpdateCallback != nil {
		var frontRepeat *Repeat
		if front != nil {
			frontRepeat, _ = front.(*Repeat)
		}
		stage.OnAfterRepeatUpdateCallback.OnAfterUpdate(stage, repeat, frontRepeat)
	}
}

func (repeat *Repeat) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRepeatDeleteCallback != nil {
		var frontRepeat *Repeat
		if front != nil {
			frontRepeat, _ = front.(*Repeat)
		}
		stage.OnAfterRepeatDeleteCallback.OnAfterDelete(stage, repeat, frontRepeat)
	}
}

func (rest *Rest) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRestCreateCallback != nil {
		stage.OnAfterRestCreateCallback.OnAfterCreate(stage, rest)
	}
}

func (rest *Rest) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRestUpdateCallback != nil {
		var frontRest *Rest
		if front != nil {
			frontRest, _ = front.(*Rest)
		}
		stage.OnAfterRestUpdateCallback.OnAfterUpdate(stage, rest, frontRest)
	}
}

func (rest *Rest) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRestDeleteCallback != nil {
		var frontRest *Rest
		if front != nil {
			frontRest, _ = front.(*Rest)
		}
		stage.OnAfterRestDeleteCallback.OnAfterDelete(stage, rest, frontRest)
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

func (root_step *Root_step) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRoot_stepCreateCallback != nil {
		stage.OnAfterRoot_stepCreateCallback.OnAfterCreate(stage, root_step)
	}
}

func (root_step *Root_step) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoot_stepUpdateCallback != nil {
		var frontRoot_step *Root_step
		if front != nil {
			frontRoot_step, _ = front.(*Root_step)
		}
		stage.OnAfterRoot_stepUpdateCallback.OnAfterUpdate(stage, root_step, frontRoot_step)
	}
}

func (root_step *Root_step) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoot_stepDeleteCallback != nil {
		var frontRoot_step *Root_step
		if front != nil {
			frontRoot_step, _ = front.(*Root_step)
		}
		stage.OnAfterRoot_stepDeleteCallback.OnAfterDelete(stage, root_step, frontRoot_step)
	}
}

func (scaling *Scaling) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScalingCreateCallback != nil {
		stage.OnAfterScalingCreateCallback.OnAfterCreate(stage, scaling)
	}
}

func (scaling *Scaling) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScalingUpdateCallback != nil {
		var frontScaling *Scaling
		if front != nil {
			frontScaling, _ = front.(*Scaling)
		}
		stage.OnAfterScalingUpdateCallback.OnAfterUpdate(stage, scaling, frontScaling)
	}
}

func (scaling *Scaling) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScalingDeleteCallback != nil {
		var frontScaling *Scaling
		if front != nil {
			frontScaling, _ = front.(*Scaling)
		}
		stage.OnAfterScalingDeleteCallback.OnAfterDelete(stage, scaling, frontScaling)
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

func (score_instrument *Score_instrument) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScore_instrumentCreateCallback != nil {
		stage.OnAfterScore_instrumentCreateCallback.OnAfterCreate(stage, score_instrument)
	}
}

func (score_instrument *Score_instrument) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScore_instrumentUpdateCallback != nil {
		var frontScore_instrument *Score_instrument
		if front != nil {
			frontScore_instrument, _ = front.(*Score_instrument)
		}
		stage.OnAfterScore_instrumentUpdateCallback.OnAfterUpdate(stage, score_instrument, frontScore_instrument)
	}
}

func (score_instrument *Score_instrument) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScore_instrumentDeleteCallback != nil {
		var frontScore_instrument *Score_instrument
		if front != nil {
			frontScore_instrument, _ = front.(*Score_instrument)
		}
		stage.OnAfterScore_instrumentDeleteCallback.OnAfterDelete(stage, score_instrument, frontScore_instrument)
	}
}

func (score_part *Score_part) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScore_partCreateCallback != nil {
		stage.OnAfterScore_partCreateCallback.OnAfterCreate(stage, score_part)
	}
}

func (score_part *Score_part) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScore_partUpdateCallback != nil {
		var frontScore_part *Score_part
		if front != nil {
			frontScore_part, _ = front.(*Score_part)
		}
		stage.OnAfterScore_partUpdateCallback.OnAfterUpdate(stage, score_part, frontScore_part)
	}
}

func (score_part *Score_part) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScore_partDeleteCallback != nil {
		var frontScore_part *Score_part
		if front != nil {
			frontScore_part, _ = front.(*Score_part)
		}
		stage.OnAfterScore_partDeleteCallback.OnAfterDelete(stage, score_part, frontScore_part)
	}
}

func (score_partwise *Score_partwise) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScore_partwiseCreateCallback != nil {
		stage.OnAfterScore_partwiseCreateCallback.OnAfterCreate(stage, score_partwise)
	}
}

func (score_partwise *Score_partwise) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScore_partwiseUpdateCallback != nil {
		var frontScore_partwise *Score_partwise
		if front != nil {
			frontScore_partwise, _ = front.(*Score_partwise)
		}
		stage.OnAfterScore_partwiseUpdateCallback.OnAfterUpdate(stage, score_partwise, frontScore_partwise)
	}
}

func (score_partwise *Score_partwise) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScore_partwiseDeleteCallback != nil {
		var frontScore_partwise *Score_partwise
		if front != nil {
			frontScore_partwise, _ = front.(*Score_partwise)
		}
		stage.OnAfterScore_partwiseDeleteCallback.OnAfterDelete(stage, score_partwise, frontScore_partwise)
	}
}

func (score_timewise *Score_timewise) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScore_timewiseCreateCallback != nil {
		stage.OnAfterScore_timewiseCreateCallback.OnAfterCreate(stage, score_timewise)
	}
}

func (score_timewise *Score_timewise) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScore_timewiseUpdateCallback != nil {
		var frontScore_timewise *Score_timewise
		if front != nil {
			frontScore_timewise, _ = front.(*Score_timewise)
		}
		stage.OnAfterScore_timewiseUpdateCallback.OnAfterUpdate(stage, score_timewise, frontScore_timewise)
	}
}

func (score_timewise *Score_timewise) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScore_timewiseDeleteCallback != nil {
		var frontScore_timewise *Score_timewise
		if front != nil {
			frontScore_timewise, _ = front.(*Score_timewise)
		}
		stage.OnAfterScore_timewiseDeleteCallback.OnAfterDelete(stage, score_timewise, frontScore_timewise)
	}
}

func (segno *Segno) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSegnoCreateCallback != nil {
		stage.OnAfterSegnoCreateCallback.OnAfterCreate(stage, segno)
	}
}

func (segno *Segno) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSegnoUpdateCallback != nil {
		var frontSegno *Segno
		if front != nil {
			frontSegno, _ = front.(*Segno)
		}
		stage.OnAfterSegnoUpdateCallback.OnAfterUpdate(stage, segno, frontSegno)
	}
}

func (segno *Segno) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSegnoDeleteCallback != nil {
		var frontSegno *Segno
		if front != nil {
			frontSegno, _ = front.(*Segno)
		}
		stage.OnAfterSegnoDeleteCallback.OnAfterDelete(stage, segno, frontSegno)
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

func (slide *Slide) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSlideCreateCallback != nil {
		stage.OnAfterSlideCreateCallback.OnAfterCreate(stage, slide)
	}
}

func (slide *Slide) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSlideUpdateCallback != nil {
		var frontSlide *Slide
		if front != nil {
			frontSlide, _ = front.(*Slide)
		}
		stage.OnAfterSlideUpdateCallback.OnAfterUpdate(stage, slide, frontSlide)
	}
}

func (slide *Slide) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSlideDeleteCallback != nil {
		var frontSlide *Slide
		if front != nil {
			frontSlide, _ = front.(*Slide)
		}
		stage.OnAfterSlideDeleteCallback.OnAfterDelete(stage, slide, frontSlide)
	}
}

func (slur *Slur) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSlurCreateCallback != nil {
		stage.OnAfterSlurCreateCallback.OnAfterCreate(stage, slur)
	}
}

func (slur *Slur) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSlurUpdateCallback != nil {
		var frontSlur *Slur
		if front != nil {
			frontSlur, _ = front.(*Slur)
		}
		stage.OnAfterSlurUpdateCallback.OnAfterUpdate(stage, slur, frontSlur)
	}
}

func (slur *Slur) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSlurDeleteCallback != nil {
		var frontSlur *Slur
		if front != nil {
			frontSlur, _ = front.(*Slur)
		}
		stage.OnAfterSlurDeleteCallback.OnAfterDelete(stage, slur, frontSlur)
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

func (staff_details *Staff_details) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaff_detailsCreateCallback != nil {
		stage.OnAfterStaff_detailsCreateCallback.OnAfterCreate(stage, staff_details)
	}
}

func (staff_details *Staff_details) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_detailsUpdateCallback != nil {
		var frontStaff_details *Staff_details
		if front != nil {
			frontStaff_details, _ = front.(*Staff_details)
		}
		stage.OnAfterStaff_detailsUpdateCallback.OnAfterUpdate(stage, staff_details, frontStaff_details)
	}
}

func (staff_details *Staff_details) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_detailsDeleteCallback != nil {
		var frontStaff_details *Staff_details
		if front != nil {
			frontStaff_details, _ = front.(*Staff_details)
		}
		stage.OnAfterStaff_detailsDeleteCallback.OnAfterDelete(stage, staff_details, frontStaff_details)
	}
}

func (staff_divide *Staff_divide) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaff_divideCreateCallback != nil {
		stage.OnAfterStaff_divideCreateCallback.OnAfterCreate(stage, staff_divide)
	}
}

func (staff_divide *Staff_divide) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_divideUpdateCallback != nil {
		var frontStaff_divide *Staff_divide
		if front != nil {
			frontStaff_divide, _ = front.(*Staff_divide)
		}
		stage.OnAfterStaff_divideUpdateCallback.OnAfterUpdate(stage, staff_divide, frontStaff_divide)
	}
}

func (staff_divide *Staff_divide) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_divideDeleteCallback != nil {
		var frontStaff_divide *Staff_divide
		if front != nil {
			frontStaff_divide, _ = front.(*Staff_divide)
		}
		stage.OnAfterStaff_divideDeleteCallback.OnAfterDelete(stage, staff_divide, frontStaff_divide)
	}
}

func (staff_layout *Staff_layout) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaff_layoutCreateCallback != nil {
		stage.OnAfterStaff_layoutCreateCallback.OnAfterCreate(stage, staff_layout)
	}
}

func (staff_layout *Staff_layout) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_layoutUpdateCallback != nil {
		var frontStaff_layout *Staff_layout
		if front != nil {
			frontStaff_layout, _ = front.(*Staff_layout)
		}
		stage.OnAfterStaff_layoutUpdateCallback.OnAfterUpdate(stage, staff_layout, frontStaff_layout)
	}
}

func (staff_layout *Staff_layout) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_layoutDeleteCallback != nil {
		var frontStaff_layout *Staff_layout
		if front != nil {
			frontStaff_layout, _ = front.(*Staff_layout)
		}
		stage.OnAfterStaff_layoutDeleteCallback.OnAfterDelete(stage, staff_layout, frontStaff_layout)
	}
}

func (staff_size *Staff_size) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaff_sizeCreateCallback != nil {
		stage.OnAfterStaff_sizeCreateCallback.OnAfterCreate(stage, staff_size)
	}
}

func (staff_size *Staff_size) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_sizeUpdateCallback != nil {
		var frontStaff_size *Staff_size
		if front != nil {
			frontStaff_size, _ = front.(*Staff_size)
		}
		stage.OnAfterStaff_sizeUpdateCallback.OnAfterUpdate(stage, staff_size, frontStaff_size)
	}
}

func (staff_size *Staff_size) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_sizeDeleteCallback != nil {
		var frontStaff_size *Staff_size
		if front != nil {
			frontStaff_size, _ = front.(*Staff_size)
		}
		stage.OnAfterStaff_sizeDeleteCallback.OnAfterDelete(stage, staff_size, frontStaff_size)
	}
}

func (staff_tuning *Staff_tuning) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStaff_tuningCreateCallback != nil {
		stage.OnAfterStaff_tuningCreateCallback.OnAfterCreate(stage, staff_tuning)
	}
}

func (staff_tuning *Staff_tuning) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_tuningUpdateCallback != nil {
		var frontStaff_tuning *Staff_tuning
		if front != nil {
			frontStaff_tuning, _ = front.(*Staff_tuning)
		}
		stage.OnAfterStaff_tuningUpdateCallback.OnAfterUpdate(stage, staff_tuning, frontStaff_tuning)
	}
}

func (staff_tuning *Staff_tuning) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStaff_tuningDeleteCallback != nil {
		var frontStaff_tuning *Staff_tuning
		if front != nil {
			frontStaff_tuning, _ = front.(*Staff_tuning)
		}
		stage.OnAfterStaff_tuningDeleteCallback.OnAfterDelete(stage, staff_tuning, frontStaff_tuning)
	}
}

func (stem *Stem) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStemCreateCallback != nil {
		stage.OnAfterStemCreateCallback.OnAfterCreate(stage, stem)
	}
}

func (stem *Stem) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStemUpdateCallback != nil {
		var frontStem *Stem
		if front != nil {
			frontStem, _ = front.(*Stem)
		}
		stage.OnAfterStemUpdateCallback.OnAfterUpdate(stage, stem, frontStem)
	}
}

func (stem *Stem) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStemDeleteCallback != nil {
		var frontStem *Stem
		if front != nil {
			frontStem, _ = front.(*Stem)
		}
		stage.OnAfterStemDeleteCallback.OnAfterDelete(stage, stem, frontStem)
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

func (string_mute *String_mute) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterString_muteCreateCallback != nil {
		stage.OnAfterString_muteCreateCallback.OnAfterCreate(stage, string_mute)
	}
}

func (string_mute *String_mute) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterString_muteUpdateCallback != nil {
		var frontString_mute *String_mute
		if front != nil {
			frontString_mute, _ = front.(*String_mute)
		}
		stage.OnAfterString_muteUpdateCallback.OnAfterUpdate(stage, string_mute, frontString_mute)
	}
}

func (string_mute *String_mute) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterString_muteDeleteCallback != nil {
		var frontString_mute *String_mute
		if front != nil {
			frontString_mute, _ = front.(*String_mute)
		}
		stage.OnAfterString_muteDeleteCallback.OnAfterDelete(stage, string_mute, frontString_mute)
	}
}

func (string_type *String_type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterString_typeCreateCallback != nil {
		stage.OnAfterString_typeCreateCallback.OnAfterCreate(stage, string_type)
	}
}

func (string_type *String_type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterString_typeUpdateCallback != nil {
		var frontString_type *String_type
		if front != nil {
			frontString_type, _ = front.(*String_type)
		}
		stage.OnAfterString_typeUpdateCallback.OnAfterUpdate(stage, string_type, frontString_type)
	}
}

func (string_type *String_type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterString_typeDeleteCallback != nil {
		var frontString_type *String_type
		if front != nil {
			frontString_type, _ = front.(*String_type)
		}
		stage.OnAfterString_typeDeleteCallback.OnAfterDelete(stage, string_type, frontString_type)
	}
}

func (strong_accent *Strong_accent) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStrong_accentCreateCallback != nil {
		stage.OnAfterStrong_accentCreateCallback.OnAfterCreate(stage, strong_accent)
	}
}

func (strong_accent *Strong_accent) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStrong_accentUpdateCallback != nil {
		var frontStrong_accent *Strong_accent
		if front != nil {
			frontStrong_accent, _ = front.(*Strong_accent)
		}
		stage.OnAfterStrong_accentUpdateCallback.OnAfterUpdate(stage, strong_accent, frontStrong_accent)
	}
}

func (strong_accent *Strong_accent) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStrong_accentDeleteCallback != nil {
		var frontStrong_accent *Strong_accent
		if front != nil {
			frontStrong_accent, _ = front.(*Strong_accent)
		}
		stage.OnAfterStrong_accentDeleteCallback.OnAfterDelete(stage, strong_accent, frontStrong_accent)
	}
}

func (style_text *Style_text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStyle_textCreateCallback != nil {
		stage.OnAfterStyle_textCreateCallback.OnAfterCreate(stage, style_text)
	}
}

func (style_text *Style_text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStyle_textUpdateCallback != nil {
		var frontStyle_text *Style_text
		if front != nil {
			frontStyle_text, _ = front.(*Style_text)
		}
		stage.OnAfterStyle_textUpdateCallback.OnAfterUpdate(stage, style_text, frontStyle_text)
	}
}

func (style_text *Style_text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStyle_textDeleteCallback != nil {
		var frontStyle_text *Style_text
		if front != nil {
			frontStyle_text, _ = front.(*Style_text)
		}
		stage.OnAfterStyle_textDeleteCallback.OnAfterDelete(stage, style_text, frontStyle_text)
	}
}

func (supports *Supports) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSupportsCreateCallback != nil {
		stage.OnAfterSupportsCreateCallback.OnAfterCreate(stage, supports)
	}
}

func (supports *Supports) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSupportsUpdateCallback != nil {
		var frontSupports *Supports
		if front != nil {
			frontSupports, _ = front.(*Supports)
		}
		stage.OnAfterSupportsUpdateCallback.OnAfterUpdate(stage, supports, frontSupports)
	}
}

func (supports *Supports) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSupportsDeleteCallback != nil {
		var frontSupports *Supports
		if front != nil {
			frontSupports, _ = front.(*Supports)
		}
		stage.OnAfterSupportsDeleteCallback.OnAfterDelete(stage, supports, frontSupports)
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

func (system_dividers *System_dividers) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSystem_dividersCreateCallback != nil {
		stage.OnAfterSystem_dividersCreateCallback.OnAfterCreate(stage, system_dividers)
	}
}

func (system_dividers *System_dividers) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystem_dividersUpdateCallback != nil {
		var frontSystem_dividers *System_dividers
		if front != nil {
			frontSystem_dividers, _ = front.(*System_dividers)
		}
		stage.OnAfterSystem_dividersUpdateCallback.OnAfterUpdate(stage, system_dividers, frontSystem_dividers)
	}
}

func (system_dividers *System_dividers) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystem_dividersDeleteCallback != nil {
		var frontSystem_dividers *System_dividers
		if front != nil {
			frontSystem_dividers, _ = front.(*System_dividers)
		}
		stage.OnAfterSystem_dividersDeleteCallback.OnAfterDelete(stage, system_dividers, frontSystem_dividers)
	}
}

func (system_layout *System_layout) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSystem_layoutCreateCallback != nil {
		stage.OnAfterSystem_layoutCreateCallback.OnAfterCreate(stage, system_layout)
	}
}

func (system_layout *System_layout) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystem_layoutUpdateCallback != nil {
		var frontSystem_layout *System_layout
		if front != nil {
			frontSystem_layout, _ = front.(*System_layout)
		}
		stage.OnAfterSystem_layoutUpdateCallback.OnAfterUpdate(stage, system_layout, frontSystem_layout)
	}
}

func (system_layout *System_layout) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystem_layoutDeleteCallback != nil {
		var frontSystem_layout *System_layout
		if front != nil {
			frontSystem_layout, _ = front.(*System_layout)
		}
		stage.OnAfterSystem_layoutDeleteCallback.OnAfterDelete(stage, system_layout, frontSystem_layout)
	}
}

func (system_margins *System_margins) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSystem_marginsCreateCallback != nil {
		stage.OnAfterSystem_marginsCreateCallback.OnAfterCreate(stage, system_margins)
	}
}

func (system_margins *System_margins) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystem_marginsUpdateCallback != nil {
		var frontSystem_margins *System_margins
		if front != nil {
			frontSystem_margins, _ = front.(*System_margins)
		}
		stage.OnAfterSystem_marginsUpdateCallback.OnAfterUpdate(stage, system_margins, frontSystem_margins)
	}
}

func (system_margins *System_margins) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystem_marginsDeleteCallback != nil {
		var frontSystem_margins *System_margins
		if front != nil {
			frontSystem_margins, _ = front.(*System_margins)
		}
		stage.OnAfterSystem_marginsDeleteCallback.OnAfterDelete(stage, system_margins, frontSystem_margins)
	}
}

func (tap *Tap) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTapCreateCallback != nil {
		stage.OnAfterTapCreateCallback.OnAfterCreate(stage, tap)
	}
}

func (tap *Tap) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTapUpdateCallback != nil {
		var frontTap *Tap
		if front != nil {
			frontTap, _ = front.(*Tap)
		}
		stage.OnAfterTapUpdateCallback.OnAfterUpdate(stage, tap, frontTap)
	}
}

func (tap *Tap) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTapDeleteCallback != nil {
		var frontTap *Tap
		if front != nil {
			frontTap, _ = front.(*Tap)
		}
		stage.OnAfterTapDeleteCallback.OnAfterDelete(stage, tap, frontTap)
	}
}

func (technical *Technical) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTechnicalCreateCallback != nil {
		stage.OnAfterTechnicalCreateCallback.OnAfterCreate(stage, technical)
	}
}

func (technical *Technical) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTechnicalUpdateCallback != nil {
		var frontTechnical *Technical
		if front != nil {
			frontTechnical, _ = front.(*Technical)
		}
		stage.OnAfterTechnicalUpdateCallback.OnAfterUpdate(stage, technical, frontTechnical)
	}
}

func (technical *Technical) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTechnicalDeleteCallback != nil {
		var frontTechnical *Technical
		if front != nil {
			frontTechnical, _ = front.(*Technical)
		}
		stage.OnAfterTechnicalDeleteCallback.OnAfterDelete(stage, technical, frontTechnical)
	}
}

func (text_element_data *Text_element_data) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterText_element_dataCreateCallback != nil {
		stage.OnAfterText_element_dataCreateCallback.OnAfterCreate(stage, text_element_data)
	}
}

func (text_element_data *Text_element_data) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterText_element_dataUpdateCallback != nil {
		var frontText_element_data *Text_element_data
		if front != nil {
			frontText_element_data, _ = front.(*Text_element_data)
		}
		stage.OnAfterText_element_dataUpdateCallback.OnAfterUpdate(stage, text_element_data, frontText_element_data)
	}
}

func (text_element_data *Text_element_data) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterText_element_dataDeleteCallback != nil {
		var frontText_element_data *Text_element_data
		if front != nil {
			frontText_element_data, _ = front.(*Text_element_data)
		}
		stage.OnAfterText_element_dataDeleteCallback.OnAfterDelete(stage, text_element_data, frontText_element_data)
	}
}

func (tie *Tie) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTieCreateCallback != nil {
		stage.OnAfterTieCreateCallback.OnAfterCreate(stage, tie)
	}
}

func (tie *Tie) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTieUpdateCallback != nil {
		var frontTie *Tie
		if front != nil {
			frontTie, _ = front.(*Tie)
		}
		stage.OnAfterTieUpdateCallback.OnAfterUpdate(stage, tie, frontTie)
	}
}

func (tie *Tie) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTieDeleteCallback != nil {
		var frontTie *Tie
		if front != nil {
			frontTie, _ = front.(*Tie)
		}
		stage.OnAfterTieDeleteCallback.OnAfterDelete(stage, tie, frontTie)
	}
}

func (tied *Tied) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTiedCreateCallback != nil {
		stage.OnAfterTiedCreateCallback.OnAfterCreate(stage, tied)
	}
}

func (tied *Tied) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTiedUpdateCallback != nil {
		var frontTied *Tied
		if front != nil {
			frontTied, _ = front.(*Tied)
		}
		stage.OnAfterTiedUpdateCallback.OnAfterUpdate(stage, tied, frontTied)
	}
}

func (tied *Tied) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTiedDeleteCallback != nil {
		var frontTied *Tied
		if front != nil {
			frontTied, _ = front.(*Tied)
		}
		stage.OnAfterTiedDeleteCallback.OnAfterDelete(stage, tied, frontTied)
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

func (time_modification *Time_modification) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTime_modificationCreateCallback != nil {
		stage.OnAfterTime_modificationCreateCallback.OnAfterCreate(stage, time_modification)
	}
}

func (time_modification *Time_modification) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTime_modificationUpdateCallback != nil {
		var frontTime_modification *Time_modification
		if front != nil {
			frontTime_modification, _ = front.(*Time_modification)
		}
		stage.OnAfterTime_modificationUpdateCallback.OnAfterUpdate(stage, time_modification, frontTime_modification)
	}
}

func (time_modification *Time_modification) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTime_modificationDeleteCallback != nil {
		var frontTime_modification *Time_modification
		if front != nil {
			frontTime_modification, _ = front.(*Time_modification)
		}
		stage.OnAfterTime_modificationDeleteCallback.OnAfterDelete(stage, time_modification, frontTime_modification)
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

func (tremolo *Tremolo) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTremoloCreateCallback != nil {
		stage.OnAfterTremoloCreateCallback.OnAfterCreate(stage, tremolo)
	}
}

func (tremolo *Tremolo) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTremoloUpdateCallback != nil {
		var frontTremolo *Tremolo
		if front != nil {
			frontTremolo, _ = front.(*Tremolo)
		}
		stage.OnAfterTremoloUpdateCallback.OnAfterUpdate(stage, tremolo, frontTremolo)
	}
}

func (tremolo *Tremolo) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTremoloDeleteCallback != nil {
		var frontTremolo *Tremolo
		if front != nil {
			frontTremolo, _ = front.(*Tremolo)
		}
		stage.OnAfterTremoloDeleteCallback.OnAfterDelete(stage, tremolo, frontTremolo)
	}
}

func (tuplet *Tuplet) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTupletCreateCallback != nil {
		stage.OnAfterTupletCreateCallback.OnAfterCreate(stage, tuplet)
	}
}

func (tuplet *Tuplet) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletUpdateCallback != nil {
		var frontTuplet *Tuplet
		if front != nil {
			frontTuplet, _ = front.(*Tuplet)
		}
		stage.OnAfterTupletUpdateCallback.OnAfterUpdate(stage, tuplet, frontTuplet)
	}
}

func (tuplet *Tuplet) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletDeleteCallback != nil {
		var frontTuplet *Tuplet
		if front != nil {
			frontTuplet, _ = front.(*Tuplet)
		}
		stage.OnAfterTupletDeleteCallback.OnAfterDelete(stage, tuplet, frontTuplet)
	}
}

func (tuplet_dot *Tuplet_dot) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTuplet_dotCreateCallback != nil {
		stage.OnAfterTuplet_dotCreateCallback.OnAfterCreate(stage, tuplet_dot)
	}
}

func (tuplet_dot *Tuplet_dot) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTuplet_dotUpdateCallback != nil {
		var frontTuplet_dot *Tuplet_dot
		if front != nil {
			frontTuplet_dot, _ = front.(*Tuplet_dot)
		}
		stage.OnAfterTuplet_dotUpdateCallback.OnAfterUpdate(stage, tuplet_dot, frontTuplet_dot)
	}
}

func (tuplet_dot *Tuplet_dot) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTuplet_dotDeleteCallback != nil {
		var frontTuplet_dot *Tuplet_dot
		if front != nil {
			frontTuplet_dot, _ = front.(*Tuplet_dot)
		}
		stage.OnAfterTuplet_dotDeleteCallback.OnAfterDelete(stage, tuplet_dot, frontTuplet_dot)
	}
}

func (tuplet_number *Tuplet_number) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTuplet_numberCreateCallback != nil {
		stage.OnAfterTuplet_numberCreateCallback.OnAfterCreate(stage, tuplet_number)
	}
}

func (tuplet_number *Tuplet_number) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTuplet_numberUpdateCallback != nil {
		var frontTuplet_number *Tuplet_number
		if front != nil {
			frontTuplet_number, _ = front.(*Tuplet_number)
		}
		stage.OnAfterTuplet_numberUpdateCallback.OnAfterUpdate(stage, tuplet_number, frontTuplet_number)
	}
}

func (tuplet_number *Tuplet_number) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTuplet_numberDeleteCallback != nil {
		var frontTuplet_number *Tuplet_number
		if front != nil {
			frontTuplet_number, _ = front.(*Tuplet_number)
		}
		stage.OnAfterTuplet_numberDeleteCallback.OnAfterDelete(stage, tuplet_number, frontTuplet_number)
	}
}

func (tuplet_portion *Tuplet_portion) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTuplet_portionCreateCallback != nil {
		stage.OnAfterTuplet_portionCreateCallback.OnAfterCreate(stage, tuplet_portion)
	}
}

func (tuplet_portion *Tuplet_portion) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTuplet_portionUpdateCallback != nil {
		var frontTuplet_portion *Tuplet_portion
		if front != nil {
			frontTuplet_portion, _ = front.(*Tuplet_portion)
		}
		stage.OnAfterTuplet_portionUpdateCallback.OnAfterUpdate(stage, tuplet_portion, frontTuplet_portion)
	}
}

func (tuplet_portion *Tuplet_portion) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTuplet_portionDeleteCallback != nil {
		var frontTuplet_portion *Tuplet_portion
		if front != nil {
			frontTuplet_portion, _ = front.(*Tuplet_portion)
		}
		stage.OnAfterTuplet_portionDeleteCallback.OnAfterDelete(stage, tuplet_portion, frontTuplet_portion)
	}
}

func (tuplet_type *Tuplet_type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTuplet_typeCreateCallback != nil {
		stage.OnAfterTuplet_typeCreateCallback.OnAfterCreate(stage, tuplet_type)
	}
}

func (tuplet_type *Tuplet_type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTuplet_typeUpdateCallback != nil {
		var frontTuplet_type *Tuplet_type
		if front != nil {
			frontTuplet_type, _ = front.(*Tuplet_type)
		}
		stage.OnAfterTuplet_typeUpdateCallback.OnAfterUpdate(stage, tuplet_type, frontTuplet_type)
	}
}

func (tuplet_type *Tuplet_type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTuplet_typeDeleteCallback != nil {
		var frontTuplet_type *Tuplet_type
		if front != nil {
			frontTuplet_type, _ = front.(*Tuplet_type)
		}
		stage.OnAfterTuplet_typeDeleteCallback.OnAfterDelete(stage, tuplet_type, frontTuplet_type)
	}
}

func (typed_text *Typed_text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTyped_textCreateCallback != nil {
		stage.OnAfterTyped_textCreateCallback.OnAfterCreate(stage, typed_text)
	}
}

func (typed_text *Typed_text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTyped_textUpdateCallback != nil {
		var frontTyped_text *Typed_text
		if front != nil {
			frontTyped_text, _ = front.(*Typed_text)
		}
		stage.OnAfterTyped_textUpdateCallback.OnAfterUpdate(stage, typed_text, frontTyped_text)
	}
}

func (typed_text *Typed_text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTyped_textDeleteCallback != nil {
		var frontTyped_text *Typed_text
		if front != nil {
			frontTyped_text, _ = front.(*Typed_text)
		}
		stage.OnAfterTyped_textDeleteCallback.OnAfterDelete(stage, typed_text, frontTyped_text)
	}
}

func (unpitched *Unpitched) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnpitchedCreateCallback != nil {
		stage.OnAfterUnpitchedCreateCallback.OnAfterCreate(stage, unpitched)
	}
}

func (unpitched *Unpitched) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnpitchedUpdateCallback != nil {
		var frontUnpitched *Unpitched
		if front != nil {
			frontUnpitched, _ = front.(*Unpitched)
		}
		stage.OnAfterUnpitchedUpdateCallback.OnAfterUpdate(stage, unpitched, frontUnpitched)
	}
}

func (unpitched *Unpitched) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnpitchedDeleteCallback != nil {
		var frontUnpitched *Unpitched
		if front != nil {
			frontUnpitched, _ = front.(*Unpitched)
		}
		stage.OnAfterUnpitchedDeleteCallback.OnAfterDelete(stage, unpitched, frontUnpitched)
	}
}

func (virtual_instrument *Virtual_instrument) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVirtual_instrumentCreateCallback != nil {
		stage.OnAfterVirtual_instrumentCreateCallback.OnAfterCreate(stage, virtual_instrument)
	}
}

func (virtual_instrument *Virtual_instrument) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVirtual_instrumentUpdateCallback != nil {
		var frontVirtual_instrument *Virtual_instrument
		if front != nil {
			frontVirtual_instrument, _ = front.(*Virtual_instrument)
		}
		stage.OnAfterVirtual_instrumentUpdateCallback.OnAfterUpdate(stage, virtual_instrument, frontVirtual_instrument)
	}
}

func (virtual_instrument *Virtual_instrument) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVirtual_instrumentDeleteCallback != nil {
		var frontVirtual_instrument *Virtual_instrument
		if front != nil {
			frontVirtual_instrument, _ = front.(*Virtual_instrument)
		}
		stage.OnAfterVirtual_instrumentDeleteCallback.OnAfterDelete(stage, virtual_instrument, frontVirtual_instrument)
	}
}

func (wait *Wait) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWaitCreateCallback != nil {
		stage.OnAfterWaitCreateCallback.OnAfterCreate(stage, wait)
	}
}

func (wait *Wait) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWaitUpdateCallback != nil {
		var frontWait *Wait
		if front != nil {
			frontWait, _ = front.(*Wait)
		}
		stage.OnAfterWaitUpdateCallback.OnAfterUpdate(stage, wait, frontWait)
	}
}

func (wait *Wait) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWaitDeleteCallback != nil {
		var frontWait *Wait
		if front != nil {
			frontWait, _ = front.(*Wait)
		}
		stage.OnAfterWaitDeleteCallback.OnAfterDelete(stage, wait, frontWait)
	}
}

func (wavy_line *Wavy_line) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWavy_lineCreateCallback != nil {
		stage.OnAfterWavy_lineCreateCallback.OnAfterCreate(stage, wavy_line)
	}
}

func (wavy_line *Wavy_line) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWavy_lineUpdateCallback != nil {
		var frontWavy_line *Wavy_line
		if front != nil {
			frontWavy_line, _ = front.(*Wavy_line)
		}
		stage.OnAfterWavy_lineUpdateCallback.OnAfterUpdate(stage, wavy_line, frontWavy_line)
	}
}

func (wavy_line *Wavy_line) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWavy_lineDeleteCallback != nil {
		var frontWavy_line *Wavy_line
		if front != nil {
			frontWavy_line, _ = front.(*Wavy_line)
		}
		stage.OnAfterWavy_lineDeleteCallback.OnAfterDelete(stage, wavy_line, frontWavy_line)
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

func (work *Work) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWorkCreateCallback != nil {
		stage.OnAfterWorkCreateCallback.OnAfterCreate(stage, work)
	}
}

func (work *Work) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkUpdateCallback != nil {
		var frontWork *Work
		if front != nil {
			frontWork, _ = front.(*Work)
		}
		stage.OnAfterWorkUpdateCallback.OnAfterUpdate(stage, work, frontWork)
	}
}

func (work *Work) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkDeleteCallback != nil {
		var frontWork *Work
		if front != nil {
			frontWork, _ = front.(*Work)
		}
		stage.OnAfterWorkDeleteCallback.OnAfterDelete(stage, work, frontWork)
	}
}

