// generated code - do not edit
package note

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

func (accidentalmark *AccidentalMark) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAccidentalMarkCreateCallback != nil {
		stage.OnAfterAccidentalMarkCreateCallback.OnAfterCreate(stage, accidentalmark)
	}
}

func (accidentalmark *AccidentalMark) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidentalMarkUpdateCallback != nil {
		var frontAccidentalMark *AccidentalMark
		if front != nil {
			frontAccidentalMark, _ = front.(*AccidentalMark)
		}
		stage.OnAfterAccidentalMarkUpdateCallback.OnAfterUpdate(stage, accidentalmark, frontAccidentalMark)
	}
}

func (accidentalmark *AccidentalMark) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAccidentalMarkDeleteCallback != nil {
		var frontAccidentalMark *AccidentalMark
		if front != nil {
			frontAccidentalMark, _ = front.(*AccidentalMark)
		}
		stage.OnAfterAccidentalMarkDeleteCallback.OnAfterDelete(stage, accidentalmark, frontAccidentalMark)
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

func (breathmark *BreathMark) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBreathMarkCreateCallback != nil {
		stage.OnAfterBreathMarkCreateCallback.OnAfterCreate(stage, breathmark)
	}
}

func (breathmark *BreathMark) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBreathMarkUpdateCallback != nil {
		var frontBreathMark *BreathMark
		if front != nil {
			frontBreathMark, _ = front.(*BreathMark)
		}
		stage.OnAfterBreathMarkUpdateCallback.OnAfterUpdate(stage, breathmark, frontBreathMark)
	}
}

func (breathmark *BreathMark) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBreathMarkDeleteCallback != nil {
		var frontBreathMark *BreathMark
		if front != nil {
			frontBreathMark, _ = front.(*BreathMark)
		}
		stage.OnAfterBreathMarkDeleteCallback.OnAfterDelete(stage, breathmark, frontBreathMark)
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

func (emptyline *EmptyLine) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEmptyLineCreateCallback != nil {
		stage.OnAfterEmptyLineCreateCallback.OnAfterCreate(stage, emptyline)
	}
}

func (emptyline *EmptyLine) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyLineUpdateCallback != nil {
		var frontEmptyLine *EmptyLine
		if front != nil {
			frontEmptyLine, _ = front.(*EmptyLine)
		}
		stage.OnAfterEmptyLineUpdateCallback.OnAfterUpdate(stage, emptyline, frontEmptyLine)
	}
}

func (emptyline *EmptyLine) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEmptyLineDeleteCallback != nil {
		var frontEmptyLine *EmptyLine
		if front != nil {
			frontEmptyLine, _ = front.(*EmptyLine)
		}
		stage.OnAfterEmptyLineDeleteCallback.OnAfterDelete(stage, emptyline, frontEmptyLine)
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

func (figuredbass *FiguredBass) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFiguredBassCreateCallback != nil {
		stage.OnAfterFiguredBassCreateCallback.OnAfterCreate(stage, figuredbass)
	}
}

func (figuredbass *FiguredBass) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFiguredBassUpdateCallback != nil {
		var frontFiguredBass *FiguredBass
		if front != nil {
			frontFiguredBass, _ = front.(*FiguredBass)
		}
		stage.OnAfterFiguredBassUpdateCallback.OnAfterUpdate(stage, figuredbass, frontFiguredBass)
	}
}

func (figuredbass *FiguredBass) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFiguredBassDeleteCallback != nil {
		var frontFiguredBass *FiguredBass
		if front != nil {
			frontFiguredBass, _ = front.(*FiguredBass)
		}
		stage.OnAfterFiguredBassDeleteCallback.OnAfterDelete(stage, figuredbass, frontFiguredBass)
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

func (hammeronpulloff *HammerOnPullOff) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHammerOnPullOffCreateCallback != nil {
		stage.OnAfterHammerOnPullOffCreateCallback.OnAfterCreate(stage, hammeronpulloff)
	}
}

func (hammeronpulloff *HammerOnPullOff) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHammerOnPullOffUpdateCallback != nil {
		var frontHammerOnPullOff *HammerOnPullOff
		if front != nil {
			frontHammerOnPullOff, _ = front.(*HammerOnPullOff)
		}
		stage.OnAfterHammerOnPullOffUpdateCallback.OnAfterUpdate(stage, hammeronpulloff, frontHammerOnPullOff)
	}
}

func (hammeronpulloff *HammerOnPullOff) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHammerOnPullOffDeleteCallback != nil {
		var frontHammerOnPullOff *HammerOnPullOff
		if front != nil {
			frontHammerOnPullOff, _ = front.(*HammerOnPullOff)
		}
		stage.OnAfterHammerOnPullOffDeleteCallback.OnAfterDelete(stage, hammeronpulloff, frontHammerOnPullOff)
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

func (harmonclosed *HarmonClosed) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarmonClosedCreateCallback != nil {
		stage.OnAfterHarmonClosedCreateCallback.OnAfterCreate(stage, harmonclosed)
	}
}

func (harmonclosed *HarmonClosed) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonClosedUpdateCallback != nil {
		var frontHarmonClosed *HarmonClosed
		if front != nil {
			frontHarmonClosed, _ = front.(*HarmonClosed)
		}
		stage.OnAfterHarmonClosedUpdateCallback.OnAfterUpdate(stage, harmonclosed, frontHarmonClosed)
	}
}

func (harmonclosed *HarmonClosed) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonClosedDeleteCallback != nil {
		var frontHarmonClosed *HarmonClosed
		if front != nil {
			frontHarmonClosed, _ = front.(*HarmonClosed)
		}
		stage.OnAfterHarmonClosedDeleteCallback.OnAfterDelete(stage, harmonclosed, frontHarmonClosed)
	}
}

func (harmonmute *HarmonMute) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHarmonMuteCreateCallback != nil {
		stage.OnAfterHarmonMuteCreateCallback.OnAfterCreate(stage, harmonmute)
	}
}

func (harmonmute *HarmonMute) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonMuteUpdateCallback != nil {
		var frontHarmonMute *HarmonMute
		if front != nil {
			frontHarmonMute, _ = front.(*HarmonMute)
		}
		stage.OnAfterHarmonMuteUpdateCallback.OnAfterUpdate(stage, harmonmute, frontHarmonMute)
	}
}

func (harmonmute *HarmonMute) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHarmonMuteDeleteCallback != nil {
		var frontHarmonMute *HarmonMute
		if front != nil {
			frontHarmonMute, _ = front.(*HarmonMute)
		}
		stage.OnAfterHarmonMuteDeleteCallback.OnAfterDelete(stage, harmonmute, frontHarmonMute)
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

func (heeltoe *HeelToe) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHeelToeCreateCallback != nil {
		stage.OnAfterHeelToeCreateCallback.OnAfterCreate(stage, heeltoe)
	}
}

func (heeltoe *HeelToe) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHeelToeUpdateCallback != nil {
		var frontHeelToe *HeelToe
		if front != nil {
			frontHeelToe, _ = front.(*HeelToe)
		}
		stage.OnAfterHeelToeUpdateCallback.OnAfterUpdate(stage, heeltoe, frontHeelToe)
	}
}

func (heeltoe *HeelToe) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHeelToeDeleteCallback != nil {
		var frontHeelToe *HeelToe
		if front != nil {
			frontHeelToe, _ = front.(*HeelToe)
		}
		stage.OnAfterHeelToeDeleteCallback.OnAfterDelete(stage, heeltoe, frontHeelToe)
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

func (holeclosed *HoleClosed) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoleClosedCreateCallback != nil {
		stage.OnAfterHoleClosedCreateCallback.OnAfterCreate(stage, holeclosed)
	}
}

func (holeclosed *HoleClosed) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoleClosedUpdateCallback != nil {
		var frontHoleClosed *HoleClosed
		if front != nil {
			frontHoleClosed, _ = front.(*HoleClosed)
		}
		stage.OnAfterHoleClosedUpdateCallback.OnAfterUpdate(stage, holeclosed, frontHoleClosed)
	}
}

func (holeclosed *HoleClosed) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoleClosedDeleteCallback != nil {
		var frontHoleClosed *HoleClosed
		if front != nil {
			frontHoleClosed, _ = front.(*HoleClosed)
		}
		stage.OnAfterHoleClosedDeleteCallback.OnAfterDelete(stage, holeclosed, frontHoleClosed)
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

func (nonarpeggiate *NonArpeggiate) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNonArpeggiateCreateCallback != nil {
		stage.OnAfterNonArpeggiateCreateCallback.OnAfterCreate(stage, nonarpeggiate)
	}
}

func (nonarpeggiate *NonArpeggiate) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNonArpeggiateUpdateCallback != nil {
		var frontNonArpeggiate *NonArpeggiate
		if front != nil {
			frontNonArpeggiate, _ = front.(*NonArpeggiate)
		}
		stage.OnAfterNonArpeggiateUpdateCallback.OnAfterUpdate(stage, nonarpeggiate, frontNonArpeggiate)
	}
}

func (nonarpeggiate *NonArpeggiate) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNonArpeggiateDeleteCallback != nil {
		var frontNonArpeggiate *NonArpeggiate
		if front != nil {
			frontNonArpeggiate, _ = front.(*NonArpeggiate)
		}
		stage.OnAfterNonArpeggiateDeleteCallback.OnAfterDelete(stage, nonarpeggiate, frontNonArpeggiate)
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

func (notetype *NoteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteTypeCreateCallback != nil {
		stage.OnAfterNoteTypeCreateCallback.OnAfterCreate(stage, notetype)
	}
}

func (notetype *NoteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTypeUpdateCallback != nil {
		var frontNoteType *NoteType
		if front != nil {
			frontNoteType, _ = front.(*NoteType)
		}
		stage.OnAfterNoteTypeUpdateCallback.OnAfterUpdate(stage, notetype, frontNoteType)
	}
}

func (notetype *NoteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTypeDeleteCallback != nil {
		var frontNoteType *NoteType
		if front != nil {
			frontNoteType, _ = front.(*NoteType)
		}
		stage.OnAfterNoteTypeDeleteCallback.OnAfterDelete(stage, notetype, frontNoteType)
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

func (noteheadtext *NoteheadText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteheadTextCreateCallback != nil {
		stage.OnAfterNoteheadTextCreateCallback.OnAfterCreate(stage, noteheadtext)
	}
}

func (noteheadtext *NoteheadText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteheadTextUpdateCallback != nil {
		var frontNoteheadText *NoteheadText
		if front != nil {
			frontNoteheadText, _ = front.(*NoteheadText)
		}
		stage.OnAfterNoteheadTextUpdateCallback.OnAfterUpdate(stage, noteheadtext, frontNoteheadText)
	}
}

func (noteheadtext *NoteheadText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteheadTextDeleteCallback != nil {
		var frontNoteheadText *NoteheadText
		if front != nil {
			frontNoteheadText, _ = front.(*NoteheadText)
		}
		stage.OnAfterNoteheadTextDeleteCallback.OnAfterDelete(stage, noteheadtext, frontNoteheadText)
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

func (othernotation *OtherNotation) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOtherNotationCreateCallback != nil {
		stage.OnAfterOtherNotationCreateCallback.OnAfterCreate(stage, othernotation)
	}
}

func (othernotation *OtherNotation) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherNotationUpdateCallback != nil {
		var frontOtherNotation *OtherNotation
		if front != nil {
			frontOtherNotation, _ = front.(*OtherNotation)
		}
		stage.OnAfterOtherNotationUpdateCallback.OnAfterUpdate(stage, othernotation, frontOtherNotation)
	}
}

func (othernotation *OtherNotation) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherNotationDeleteCallback != nil {
		var frontOtherNotation *OtherNotation
		if front != nil {
			frontOtherNotation, _ = front.(*OtherNotation)
		}
		stage.OnAfterOtherNotationDeleteCallback.OnAfterDelete(stage, othernotation, frontOtherNotation)
	}
}

func (otherplacementtext *OtherPlacementText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOtherPlacementTextCreateCallback != nil {
		stage.OnAfterOtherPlacementTextCreateCallback.OnAfterCreate(stage, otherplacementtext)
	}
}

func (otherplacementtext *OtherPlacementText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherPlacementTextUpdateCallback != nil {
		var frontOtherPlacementText *OtherPlacementText
		if front != nil {
			frontOtherPlacementText, _ = front.(*OtherPlacementText)
		}
		stage.OnAfterOtherPlacementTextUpdateCallback.OnAfterUpdate(stage, otherplacementtext, frontOtherPlacementText)
	}
}

func (otherplacementtext *OtherPlacementText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOtherPlacementTextDeleteCallback != nil {
		var frontOtherPlacementText *OtherPlacementText
		if front != nil {
			frontOtherPlacementText, _ = front.(*OtherPlacementText)
		}
		stage.OnAfterOtherPlacementTextDeleteCallback.OnAfterDelete(stage, otherplacementtext, frontOtherPlacementText)
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

func (placementtext *PlacementText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlacementTextCreateCallback != nil {
		stage.OnAfterPlacementTextCreateCallback.OnAfterCreate(stage, placementtext)
	}
}

func (placementtext *PlacementText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlacementTextUpdateCallback != nil {
		var frontPlacementText *PlacementText
		if front != nil {
			frontPlacementText, _ = front.(*PlacementText)
		}
		stage.OnAfterPlacementTextUpdateCallback.OnAfterUpdate(stage, placementtext, frontPlacementText)
	}
}

func (placementtext *PlacementText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlacementTextDeleteCallback != nil {
		var frontPlacementText *PlacementText
		if front != nil {
			frontPlacementText, _ = front.(*PlacementText)
		}
		stage.OnAfterPlacementTextDeleteCallback.OnAfterDelete(stage, placementtext, frontPlacementText)
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

func (strongaccent *StrongAccent) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStrongAccentCreateCallback != nil {
		stage.OnAfterStrongAccentCreateCallback.OnAfterCreate(stage, strongaccent)
	}
}

func (strongaccent *StrongAccent) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStrongAccentUpdateCallback != nil {
		var frontStrongAccent *StrongAccent
		if front != nil {
			frontStrongAccent, _ = front.(*StrongAccent)
		}
		stage.OnAfterStrongAccentUpdateCallback.OnAfterUpdate(stage, strongaccent, frontStrongAccent)
	}
}

func (strongaccent *StrongAccent) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStrongAccentDeleteCallback != nil {
		var frontStrongAccent *StrongAccent
		if front != nil {
			frontStrongAccent, _ = front.(*StrongAccent)
		}
		stage.OnAfterStrongAccentDeleteCallback.OnAfterDelete(stage, strongaccent, frontStrongAccent)
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

func (textelementdata *TextElementData) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextElementDataCreateCallback != nil {
		stage.OnAfterTextElementDataCreateCallback.OnAfterCreate(stage, textelementdata)
	}
}

func (textelementdata *TextElementData) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextElementDataUpdateCallback != nil {
		var frontTextElementData *TextElementData
		if front != nil {
			frontTextElementData, _ = front.(*TextElementData)
		}
		stage.OnAfterTextElementDataUpdateCallback.OnAfterUpdate(stage, textelementdata, frontTextElementData)
	}
}

func (textelementdata *TextElementData) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextElementDataDeleteCallback != nil {
		var frontTextElementData *TextElementData
		if front != nil {
			frontTextElementData, _ = front.(*TextElementData)
		}
		stage.OnAfterTextElementDataDeleteCallback.OnAfterDelete(stage, textelementdata, frontTextElementData)
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

func (tupletdot *TupletDot) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTupletDotCreateCallback != nil {
		stage.OnAfterTupletDotCreateCallback.OnAfterCreate(stage, tupletdot)
	}
}

func (tupletdot *TupletDot) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletDotUpdateCallback != nil {
		var frontTupletDot *TupletDot
		if front != nil {
			frontTupletDot, _ = front.(*TupletDot)
		}
		stage.OnAfterTupletDotUpdateCallback.OnAfterUpdate(stage, tupletdot, frontTupletDot)
	}
}

func (tupletdot *TupletDot) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletDotDeleteCallback != nil {
		var frontTupletDot *TupletDot
		if front != nil {
			frontTupletDot, _ = front.(*TupletDot)
		}
		stage.OnAfterTupletDotDeleteCallback.OnAfterDelete(stage, tupletdot, frontTupletDot)
	}
}

func (tupletnumber *TupletNumber) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTupletNumberCreateCallback != nil {
		stage.OnAfterTupletNumberCreateCallback.OnAfterCreate(stage, tupletnumber)
	}
}

func (tupletnumber *TupletNumber) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletNumberUpdateCallback != nil {
		var frontTupletNumber *TupletNumber
		if front != nil {
			frontTupletNumber, _ = front.(*TupletNumber)
		}
		stage.OnAfterTupletNumberUpdateCallback.OnAfterUpdate(stage, tupletnumber, frontTupletNumber)
	}
}

func (tupletnumber *TupletNumber) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletNumberDeleteCallback != nil {
		var frontTupletNumber *TupletNumber
		if front != nil {
			frontTupletNumber, _ = front.(*TupletNumber)
		}
		stage.OnAfterTupletNumberDeleteCallback.OnAfterDelete(stage, tupletnumber, frontTupletNumber)
	}
}

func (tupletportion *TupletPortion) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTupletPortionCreateCallback != nil {
		stage.OnAfterTupletPortionCreateCallback.OnAfterCreate(stage, tupletportion)
	}
}

func (tupletportion *TupletPortion) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletPortionUpdateCallback != nil {
		var frontTupletPortion *TupletPortion
		if front != nil {
			frontTupletPortion, _ = front.(*TupletPortion)
		}
		stage.OnAfterTupletPortionUpdateCallback.OnAfterUpdate(stage, tupletportion, frontTupletPortion)
	}
}

func (tupletportion *TupletPortion) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletPortionDeleteCallback != nil {
		var frontTupletPortion *TupletPortion
		if front != nil {
			frontTupletPortion, _ = front.(*TupletPortion)
		}
		stage.OnAfterTupletPortionDeleteCallback.OnAfterDelete(stage, tupletportion, frontTupletPortion)
	}
}

func (tuplettype *TupletType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTupletTypeCreateCallback != nil {
		stage.OnAfterTupletTypeCreateCallback.OnAfterCreate(stage, tuplettype)
	}
}

func (tuplettype *TupletType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletTypeUpdateCallback != nil {
		var frontTupletType *TupletType
		if front != nil {
			frontTupletType, _ = front.(*TupletType)
		}
		stage.OnAfterTupletTypeUpdateCallback.OnAfterUpdate(stage, tuplettype, frontTupletType)
	}
}

func (tuplettype *TupletType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTupletTypeDeleteCallback != nil {
		var frontTupletType *TupletType
		if front != nil {
			frontTupletType, _ = front.(*TupletType)
		}
		stage.OnAfterTupletTypeDeleteCallback.OnAfterDelete(stage, tuplettype, frontTupletType)
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

