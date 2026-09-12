// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *A_directive:
		if stage.OnAfterA_directiveCreateCallback != nil {
			stage.OnAfterA_directiveCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_measure:
		if stage.OnAfterA_measureCreateCallback != nil {
			stage.OnAfterA_measureCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_measure_1:
		if stage.OnAfterA_measure_1CreateCallback != nil {
			stage.OnAfterA_measure_1CreateCallback.OnAfterCreate(stage, target)
		}
	case *A_part:
		if stage.OnAfterA_partCreateCallback != nil {
			stage.OnAfterA_partCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_part_1:
		if stage.OnAfterA_part_1CreateCallback != nil {
			stage.OnAfterA_part_1CreateCallback.OnAfterCreate(stage, target)
		}
	case *Accidental:
		if stage.OnAfterAccidentalCreateCallback != nil {
			stage.OnAfterAccidentalCreateCallback.OnAfterCreate(stage, target)
		}
	case *Accidental_mark:
		if stage.OnAfterAccidental_markCreateCallback != nil {
			stage.OnAfterAccidental_markCreateCallback.OnAfterCreate(stage, target)
		}
	case *Accidental_text:
		if stage.OnAfterAccidental_textCreateCallback != nil {
			stage.OnAfterAccidental_textCreateCallback.OnAfterCreate(stage, target)
		}
	case *Accord:
		if stage.OnAfterAccordCreateCallback != nil {
			stage.OnAfterAccordCreateCallback.OnAfterCreate(stage, target)
		}
	case *Accordion_registration:
		if stage.OnAfterAccordion_registrationCreateCallback != nil {
			stage.OnAfterAccordion_registrationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Appearance:
		if stage.OnAfterAppearanceCreateCallback != nil {
			stage.OnAfterAppearanceCreateCallback.OnAfterCreate(stage, target)
		}
	case *Arpeggiate:
		if stage.OnAfterArpeggiateCreateCallback != nil {
			stage.OnAfterArpeggiateCreateCallback.OnAfterCreate(stage, target)
		}
	case *Arrow:
		if stage.OnAfterArrowCreateCallback != nil {
			stage.OnAfterArrowCreateCallback.OnAfterCreate(stage, target)
		}
	case *Articulations:
		if stage.OnAfterArticulationsCreateCallback != nil {
			stage.OnAfterArticulationsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Assess:
		if stage.OnAfterAssessCreateCallback != nil {
			stage.OnAfterAssessCreateCallback.OnAfterCreate(stage, target)
		}
	case *Attributes:
		if stage.OnAfterAttributesCreateCallback != nil {
			stage.OnAfterAttributesCreateCallback.OnAfterCreate(stage, target)
		}
	case *Backup:
		if stage.OnAfterBackupCreateCallback != nil {
			stage.OnAfterBackupCreateCallback.OnAfterCreate(stage, target)
		}
	case *Bar_style_color:
		if stage.OnAfterBar_style_colorCreateCallback != nil {
			stage.OnAfterBar_style_colorCreateCallback.OnAfterCreate(stage, target)
		}
	case *Barline:
		if stage.OnAfterBarlineCreateCallback != nil {
			stage.OnAfterBarlineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Barre:
		if stage.OnAfterBarreCreateCallback != nil {
			stage.OnAfterBarreCreateCallback.OnAfterCreate(stage, target)
		}
	case *Bass:
		if stage.OnAfterBassCreateCallback != nil {
			stage.OnAfterBassCreateCallback.OnAfterCreate(stage, target)
		}
	case *Bass_step:
		if stage.OnAfterBass_stepCreateCallback != nil {
			stage.OnAfterBass_stepCreateCallback.OnAfterCreate(stage, target)
		}
	case *Beam:
		if stage.OnAfterBeamCreateCallback != nil {
			stage.OnAfterBeamCreateCallback.OnAfterCreate(stage, target)
		}
	case *Beat_repeat:
		if stage.OnAfterBeat_repeatCreateCallback != nil {
			stage.OnAfterBeat_repeatCreateCallback.OnAfterCreate(stage, target)
		}
	case *Beat_unit_tied:
		if stage.OnAfterBeat_unit_tiedCreateCallback != nil {
			stage.OnAfterBeat_unit_tiedCreateCallback.OnAfterCreate(stage, target)
		}
	case *Beater:
		if stage.OnAfterBeaterCreateCallback != nil {
			stage.OnAfterBeaterCreateCallback.OnAfterCreate(stage, target)
		}
	case *Bend:
		if stage.OnAfterBendCreateCallback != nil {
			stage.OnAfterBendCreateCallback.OnAfterCreate(stage, target)
		}
	case *Bookmark:
		if stage.OnAfterBookmarkCreateCallback != nil {
			stage.OnAfterBookmarkCreateCallback.OnAfterCreate(stage, target)
		}
	case *Bracket:
		if stage.OnAfterBracketCreateCallback != nil {
			stage.OnAfterBracketCreateCallback.OnAfterCreate(stage, target)
		}
	case *Breath_mark:
		if stage.OnAfterBreath_markCreateCallback != nil {
			stage.OnAfterBreath_markCreateCallback.OnAfterCreate(stage, target)
		}
	case *Caesura:
		if stage.OnAfterCaesuraCreateCallback != nil {
			stage.OnAfterCaesuraCreateCallback.OnAfterCreate(stage, target)
		}
	case *Cancel:
		if stage.OnAfterCancelCreateCallback != nil {
			stage.OnAfterCancelCreateCallback.OnAfterCreate(stage, target)
		}
	case *Clef:
		if stage.OnAfterClefCreateCallback != nil {
			stage.OnAfterClefCreateCallback.OnAfterCreate(stage, target)
		}
	case *Coda:
		if stage.OnAfterCodaCreateCallback != nil {
			stage.OnAfterCodaCreateCallback.OnAfterCreate(stage, target)
		}
	case *Credit:
		if stage.OnAfterCreditCreateCallback != nil {
			stage.OnAfterCreditCreateCallback.OnAfterCreate(stage, target)
		}
	case *Dashes:
		if stage.OnAfterDashesCreateCallback != nil {
			stage.OnAfterDashesCreateCallback.OnAfterCreate(stage, target)
		}
	case *Defaults:
		if stage.OnAfterDefaultsCreateCallback != nil {
			stage.OnAfterDefaultsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Degree:
		if stage.OnAfterDegreeCreateCallback != nil {
			stage.OnAfterDegreeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Degree_alter:
		if stage.OnAfterDegree_alterCreateCallback != nil {
			stage.OnAfterDegree_alterCreateCallback.OnAfterCreate(stage, target)
		}
	case *Degree_type:
		if stage.OnAfterDegree_typeCreateCallback != nil {
			stage.OnAfterDegree_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Degree_value:
		if stage.OnAfterDegree_valueCreateCallback != nil {
			stage.OnAfterDegree_valueCreateCallback.OnAfterCreate(stage, target)
		}
	case *Direction:
		if stage.OnAfterDirectionCreateCallback != nil {
			stage.OnAfterDirectionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Direction_type:
		if stage.OnAfterDirection_typeCreateCallback != nil {
			stage.OnAfterDirection_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Distance:
		if stage.OnAfterDistanceCreateCallback != nil {
			stage.OnAfterDistanceCreateCallback.OnAfterCreate(stage, target)
		}
	case *Double:
		if stage.OnAfterDoubleCreateCallback != nil {
			stage.OnAfterDoubleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Dynamics:
		if stage.OnAfterDynamicsCreateCallback != nil {
			stage.OnAfterDynamicsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Effect:
		if stage.OnAfterEffectCreateCallback != nil {
			stage.OnAfterEffectCreateCallback.OnAfterCreate(stage, target)
		}
	case *Elision:
		if stage.OnAfterElisionCreateCallback != nil {
			stage.OnAfterElisionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty:
		if stage.OnAfterEmptyCreateCallback != nil {
			stage.OnAfterEmptyCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty_font:
		if stage.OnAfterEmpty_fontCreateCallback != nil {
			stage.OnAfterEmpty_fontCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty_line:
		if stage.OnAfterEmpty_lineCreateCallback != nil {
			stage.OnAfterEmpty_lineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty_placement:
		if stage.OnAfterEmpty_placementCreateCallback != nil {
			stage.OnAfterEmpty_placementCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty_placement_smufl:
		if stage.OnAfterEmpty_placement_smuflCreateCallback != nil {
			stage.OnAfterEmpty_placement_smuflCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty_print_object_style_align:
		if stage.OnAfterEmpty_print_object_style_alignCreateCallback != nil {
			stage.OnAfterEmpty_print_object_style_alignCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty_print_style:
		if stage.OnAfterEmpty_print_styleCreateCallback != nil {
			stage.OnAfterEmpty_print_styleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty_print_style_align:
		if stage.OnAfterEmpty_print_style_alignCreateCallback != nil {
			stage.OnAfterEmpty_print_style_alignCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty_print_style_align_id:
		if stage.OnAfterEmpty_print_style_align_idCreateCallback != nil {
			stage.OnAfterEmpty_print_style_align_idCreateCallback.OnAfterCreate(stage, target)
		}
	case *Empty_trill_sound:
		if stage.OnAfterEmpty_trill_soundCreateCallback != nil {
			stage.OnAfterEmpty_trill_soundCreateCallback.OnAfterCreate(stage, target)
		}
	case *Encoding:
		if stage.OnAfterEncodingCreateCallback != nil {
			stage.OnAfterEncodingCreateCallback.OnAfterCreate(stage, target)
		}
	case *Ending:
		if stage.OnAfterEndingCreateCallback != nil {
			stage.OnAfterEndingCreateCallback.OnAfterCreate(stage, target)
		}
	case *Extend:
		if stage.OnAfterExtendCreateCallback != nil {
			stage.OnAfterExtendCreateCallback.OnAfterCreate(stage, target)
		}
	case *Feature:
		if stage.OnAfterFeatureCreateCallback != nil {
			stage.OnAfterFeatureCreateCallback.OnAfterCreate(stage, target)
		}
	case *Fermata:
		if stage.OnAfterFermataCreateCallback != nil {
			stage.OnAfterFermataCreateCallback.OnAfterCreate(stage, target)
		}
	case *Figure:
		if stage.OnAfterFigureCreateCallback != nil {
			stage.OnAfterFigureCreateCallback.OnAfterCreate(stage, target)
		}
	case *Figured_bass:
		if stage.OnAfterFigured_bassCreateCallback != nil {
			stage.OnAfterFigured_bassCreateCallback.OnAfterCreate(stage, target)
		}
	case *Fingering:
		if stage.OnAfterFingeringCreateCallback != nil {
			stage.OnAfterFingeringCreateCallback.OnAfterCreate(stage, target)
		}
	case *First_fret:
		if stage.OnAfterFirst_fretCreateCallback != nil {
			stage.OnAfterFirst_fretCreateCallback.OnAfterCreate(stage, target)
		}
	case *For_part:
		if stage.OnAfterFor_partCreateCallback != nil {
			stage.OnAfterFor_partCreateCallback.OnAfterCreate(stage, target)
		}
	case *Formatted_symbol:
		if stage.OnAfterFormatted_symbolCreateCallback != nil {
			stage.OnAfterFormatted_symbolCreateCallback.OnAfterCreate(stage, target)
		}
	case *Formatted_symbol_id:
		if stage.OnAfterFormatted_symbol_idCreateCallback != nil {
			stage.OnAfterFormatted_symbol_idCreateCallback.OnAfterCreate(stage, target)
		}
	case *Formatted_text:
		if stage.OnAfterFormatted_textCreateCallback != nil {
			stage.OnAfterFormatted_textCreateCallback.OnAfterCreate(stage, target)
		}
	case *Formatted_text_id:
		if stage.OnAfterFormatted_text_idCreateCallback != nil {
			stage.OnAfterFormatted_text_idCreateCallback.OnAfterCreate(stage, target)
		}
	case *Forward:
		if stage.OnAfterForwardCreateCallback != nil {
			stage.OnAfterForwardCreateCallback.OnAfterCreate(stage, target)
		}
	case *Frame:
		if stage.OnAfterFrameCreateCallback != nil {
			stage.OnAfterFrameCreateCallback.OnAfterCreate(stage, target)
		}
	case *Frame_note:
		if stage.OnAfterFrame_noteCreateCallback != nil {
			stage.OnAfterFrame_noteCreateCallback.OnAfterCreate(stage, target)
		}
	case *Fret:
		if stage.OnAfterFretCreateCallback != nil {
			stage.OnAfterFretCreateCallback.OnAfterCreate(stage, target)
		}
	case *Glass:
		if stage.OnAfterGlassCreateCallback != nil {
			stage.OnAfterGlassCreateCallback.OnAfterCreate(stage, target)
		}
	case *Glissando:
		if stage.OnAfterGlissandoCreateCallback != nil {
			stage.OnAfterGlissandoCreateCallback.OnAfterCreate(stage, target)
		}
	case *Glyph:
		if stage.OnAfterGlyphCreateCallback != nil {
			stage.OnAfterGlyphCreateCallback.OnAfterCreate(stage, target)
		}
	case *Grace:
		if stage.OnAfterGraceCreateCallback != nil {
			stage.OnAfterGraceCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group_barline:
		if stage.OnAfterGroup_barlineCreateCallback != nil {
			stage.OnAfterGroup_barlineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group_name:
		if stage.OnAfterGroup_nameCreateCallback != nil {
			stage.OnAfterGroup_nameCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group_symbol:
		if stage.OnAfterGroup_symbolCreateCallback != nil {
			stage.OnAfterGroup_symbolCreateCallback.OnAfterCreate(stage, target)
		}
	case *Grouping:
		if stage.OnAfterGroupingCreateCallback != nil {
			stage.OnAfterGroupingCreateCallback.OnAfterCreate(stage, target)
		}
	case *Hammer_on_pull_off:
		if stage.OnAfterHammer_on_pull_offCreateCallback != nil {
			stage.OnAfterHammer_on_pull_offCreateCallback.OnAfterCreate(stage, target)
		}
	case *Handbell:
		if stage.OnAfterHandbellCreateCallback != nil {
			stage.OnAfterHandbellCreateCallback.OnAfterCreate(stage, target)
		}
	case *Harmon_closed:
		if stage.OnAfterHarmon_closedCreateCallback != nil {
			stage.OnAfterHarmon_closedCreateCallback.OnAfterCreate(stage, target)
		}
	case *Harmon_mute:
		if stage.OnAfterHarmon_muteCreateCallback != nil {
			stage.OnAfterHarmon_muteCreateCallback.OnAfterCreate(stage, target)
		}
	case *Harmonic:
		if stage.OnAfterHarmonicCreateCallback != nil {
			stage.OnAfterHarmonicCreateCallback.OnAfterCreate(stage, target)
		}
	case *Harmony:
		if stage.OnAfterHarmonyCreateCallback != nil {
			stage.OnAfterHarmonyCreateCallback.OnAfterCreate(stage, target)
		}
	case *Harmony_alter:
		if stage.OnAfterHarmony_alterCreateCallback != nil {
			stage.OnAfterHarmony_alterCreateCallback.OnAfterCreate(stage, target)
		}
	case *Harp_pedals:
		if stage.OnAfterHarp_pedalsCreateCallback != nil {
			stage.OnAfterHarp_pedalsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Heel_toe:
		if stage.OnAfterHeel_toeCreateCallback != nil {
			stage.OnAfterHeel_toeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Hole:
		if stage.OnAfterHoleCreateCallback != nil {
			stage.OnAfterHoleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Hole_closed:
		if stage.OnAfterHole_closedCreateCallback != nil {
			stage.OnAfterHole_closedCreateCallback.OnAfterCreate(stage, target)
		}
	case *Horizontal_turn:
		if stage.OnAfterHorizontal_turnCreateCallback != nil {
			stage.OnAfterHorizontal_turnCreateCallback.OnAfterCreate(stage, target)
		}
	case *Identification:
		if stage.OnAfterIdentificationCreateCallback != nil {
			stage.OnAfterIdentificationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Image:
		if stage.OnAfterImageCreateCallback != nil {
			stage.OnAfterImageCreateCallback.OnAfterCreate(stage, target)
		}
	case *Instrument:
		if stage.OnAfterInstrumentCreateCallback != nil {
			stage.OnAfterInstrumentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Instrument_change:
		if stage.OnAfterInstrument_changeCreateCallback != nil {
			stage.OnAfterInstrument_changeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Instrument_link:
		if stage.OnAfterInstrument_linkCreateCallback != nil {
			stage.OnAfterInstrument_linkCreateCallback.OnAfterCreate(stage, target)
		}
	case *Interchangeable:
		if stage.OnAfterInterchangeableCreateCallback != nil {
			stage.OnAfterInterchangeableCreateCallback.OnAfterCreate(stage, target)
		}
	case *Inversion:
		if stage.OnAfterInversionCreateCallback != nil {
			stage.OnAfterInversionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Key:
		if stage.OnAfterKeyCreateCallback != nil {
			stage.OnAfterKeyCreateCallback.OnAfterCreate(stage, target)
		}
	case *Key_accidental:
		if stage.OnAfterKey_accidentalCreateCallback != nil {
			stage.OnAfterKey_accidentalCreateCallback.OnAfterCreate(stage, target)
		}
	case *Key_octave:
		if stage.OnAfterKey_octaveCreateCallback != nil {
			stage.OnAfterKey_octaveCreateCallback.OnAfterCreate(stage, target)
		}
	case *Kind:
		if stage.OnAfterKindCreateCallback != nil {
			stage.OnAfterKindCreateCallback.OnAfterCreate(stage, target)
		}
	case *Level:
		if stage.OnAfterLevelCreateCallback != nil {
			stage.OnAfterLevelCreateCallback.OnAfterCreate(stage, target)
		}
	case *Line_detail:
		if stage.OnAfterLine_detailCreateCallback != nil {
			stage.OnAfterLine_detailCreateCallback.OnAfterCreate(stage, target)
		}
	case *Line_width:
		if stage.OnAfterLine_widthCreateCallback != nil {
			stage.OnAfterLine_widthCreateCallback.OnAfterCreate(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkCreateCallback != nil {
			stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, target)
		}
	case *Listen:
		if stage.OnAfterListenCreateCallback != nil {
			stage.OnAfterListenCreateCallback.OnAfterCreate(stage, target)
		}
	case *Listening:
		if stage.OnAfterListeningCreateCallback != nil {
			stage.OnAfterListeningCreateCallback.OnAfterCreate(stage, target)
		}
	case *Lyric:
		if stage.OnAfterLyricCreateCallback != nil {
			stage.OnAfterLyricCreateCallback.OnAfterCreate(stage, target)
		}
	case *Lyric_font:
		if stage.OnAfterLyric_fontCreateCallback != nil {
			stage.OnAfterLyric_fontCreateCallback.OnAfterCreate(stage, target)
		}
	case *Lyric_language:
		if stage.OnAfterLyric_languageCreateCallback != nil {
			stage.OnAfterLyric_languageCreateCallback.OnAfterCreate(stage, target)
		}
	case *Measure_layout:
		if stage.OnAfterMeasure_layoutCreateCallback != nil {
			stage.OnAfterMeasure_layoutCreateCallback.OnAfterCreate(stage, target)
		}
	case *Measure_numbering:
		if stage.OnAfterMeasure_numberingCreateCallback != nil {
			stage.OnAfterMeasure_numberingCreateCallback.OnAfterCreate(stage, target)
		}
	case *Measure_repeat:
		if stage.OnAfterMeasure_repeatCreateCallback != nil {
			stage.OnAfterMeasure_repeatCreateCallback.OnAfterCreate(stage, target)
		}
	case *Measure_style:
		if stage.OnAfterMeasure_styleCreateCallback != nil {
			stage.OnAfterMeasure_styleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Membrane:
		if stage.OnAfterMembraneCreateCallback != nil {
			stage.OnAfterMembraneCreateCallback.OnAfterCreate(stage, target)
		}
	case *Metal:
		if stage.OnAfterMetalCreateCallback != nil {
			stage.OnAfterMetalCreateCallback.OnAfterCreate(stage, target)
		}
	case *Metronome:
		if stage.OnAfterMetronomeCreateCallback != nil {
			stage.OnAfterMetronomeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Metronome_beam:
		if stage.OnAfterMetronome_beamCreateCallback != nil {
			stage.OnAfterMetronome_beamCreateCallback.OnAfterCreate(stage, target)
		}
	case *Metronome_note:
		if stage.OnAfterMetronome_noteCreateCallback != nil {
			stage.OnAfterMetronome_noteCreateCallback.OnAfterCreate(stage, target)
		}
	case *Metronome_tied:
		if stage.OnAfterMetronome_tiedCreateCallback != nil {
			stage.OnAfterMetronome_tiedCreateCallback.OnAfterCreate(stage, target)
		}
	case *Metronome_tuplet:
		if stage.OnAfterMetronome_tupletCreateCallback != nil {
			stage.OnAfterMetronome_tupletCreateCallback.OnAfterCreate(stage, target)
		}
	case *Midi_device:
		if stage.OnAfterMidi_deviceCreateCallback != nil {
			stage.OnAfterMidi_deviceCreateCallback.OnAfterCreate(stage, target)
		}
	case *Midi_instrument:
		if stage.OnAfterMidi_instrumentCreateCallback != nil {
			stage.OnAfterMidi_instrumentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Miscellaneous:
		if stage.OnAfterMiscellaneousCreateCallback != nil {
			stage.OnAfterMiscellaneousCreateCallback.OnAfterCreate(stage, target)
		}
	case *Miscellaneous_field:
		if stage.OnAfterMiscellaneous_fieldCreateCallback != nil {
			stage.OnAfterMiscellaneous_fieldCreateCallback.OnAfterCreate(stage, target)
		}
	case *Mordent:
		if stage.OnAfterMordentCreateCallback != nil {
			stage.OnAfterMordentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Multiple_rest:
		if stage.OnAfterMultiple_restCreateCallback != nil {
			stage.OnAfterMultiple_restCreateCallback.OnAfterCreate(stage, target)
		}
	case *Name_display:
		if stage.OnAfterName_displayCreateCallback != nil {
			stage.OnAfterName_displayCreateCallback.OnAfterCreate(stage, target)
		}
	case *Non_arpeggiate:
		if stage.OnAfterNon_arpeggiateCreateCallback != nil {
			stage.OnAfterNon_arpeggiateCreateCallback.OnAfterCreate(stage, target)
		}
	case *Notations:
		if stage.OnAfterNotationsCreateCallback != nil {
			stage.OnAfterNotationsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note_size:
		if stage.OnAfterNote_sizeCreateCallback != nil {
			stage.OnAfterNote_sizeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note_type:
		if stage.OnAfterNote_typeCreateCallback != nil {
			stage.OnAfterNote_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Notehead:
		if stage.OnAfterNoteheadCreateCallback != nil {
			stage.OnAfterNoteheadCreateCallback.OnAfterCreate(stage, target)
		}
	case *Notehead_text:
		if stage.OnAfterNotehead_textCreateCallback != nil {
			stage.OnAfterNotehead_textCreateCallback.OnAfterCreate(stage, target)
		}
	case *Numeral:
		if stage.OnAfterNumeralCreateCallback != nil {
			stage.OnAfterNumeralCreateCallback.OnAfterCreate(stage, target)
		}
	case *Numeral_key:
		if stage.OnAfterNumeral_keyCreateCallback != nil {
			stage.OnAfterNumeral_keyCreateCallback.OnAfterCreate(stage, target)
		}
	case *Numeral_root:
		if stage.OnAfterNumeral_rootCreateCallback != nil {
			stage.OnAfterNumeral_rootCreateCallback.OnAfterCreate(stage, target)
		}
	case *Octave_shift:
		if stage.OnAfterOctave_shiftCreateCallback != nil {
			stage.OnAfterOctave_shiftCreateCallback.OnAfterCreate(stage, target)
		}
	case *Offset:
		if stage.OnAfterOffsetCreateCallback != nil {
			stage.OnAfterOffsetCreateCallback.OnAfterCreate(stage, target)
		}
	case *Opus:
		if stage.OnAfterOpusCreateCallback != nil {
			stage.OnAfterOpusCreateCallback.OnAfterCreate(stage, target)
		}
	case *Ornaments:
		if stage.OnAfterOrnamentsCreateCallback != nil {
			stage.OnAfterOrnamentsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Other_appearance:
		if stage.OnAfterOther_appearanceCreateCallback != nil {
			stage.OnAfterOther_appearanceCreateCallback.OnAfterCreate(stage, target)
		}
	case *Other_direction:
		if stage.OnAfterOther_directionCreateCallback != nil {
			stage.OnAfterOther_directionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Other_listening:
		if stage.OnAfterOther_listeningCreateCallback != nil {
			stage.OnAfterOther_listeningCreateCallback.OnAfterCreate(stage, target)
		}
	case *Other_notation:
		if stage.OnAfterOther_notationCreateCallback != nil {
			stage.OnAfterOther_notationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Other_placement_text:
		if stage.OnAfterOther_placement_textCreateCallback != nil {
			stage.OnAfterOther_placement_textCreateCallback.OnAfterCreate(stage, target)
		}
	case *Other_play:
		if stage.OnAfterOther_playCreateCallback != nil {
			stage.OnAfterOther_playCreateCallback.OnAfterCreate(stage, target)
		}
	case *Other_text:
		if stage.OnAfterOther_textCreateCallback != nil {
			stage.OnAfterOther_textCreateCallback.OnAfterCreate(stage, target)
		}
	case *Page_layout:
		if stage.OnAfterPage_layoutCreateCallback != nil {
			stage.OnAfterPage_layoutCreateCallback.OnAfterCreate(stage, target)
		}
	case *Page_margins:
		if stage.OnAfterPage_marginsCreateCallback != nil {
			stage.OnAfterPage_marginsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Part_clef:
		if stage.OnAfterPart_clefCreateCallback != nil {
			stage.OnAfterPart_clefCreateCallback.OnAfterCreate(stage, target)
		}
	case *Part_group:
		if stage.OnAfterPart_groupCreateCallback != nil {
			stage.OnAfterPart_groupCreateCallback.OnAfterCreate(stage, target)
		}
	case *Part_link:
		if stage.OnAfterPart_linkCreateCallback != nil {
			stage.OnAfterPart_linkCreateCallback.OnAfterCreate(stage, target)
		}
	case *Part_list:
		if stage.OnAfterPart_listCreateCallback != nil {
			stage.OnAfterPart_listCreateCallback.OnAfterCreate(stage, target)
		}
	case *Part_name:
		if stage.OnAfterPart_nameCreateCallback != nil {
			stage.OnAfterPart_nameCreateCallback.OnAfterCreate(stage, target)
		}
	case *Part_symbol:
		if stage.OnAfterPart_symbolCreateCallback != nil {
			stage.OnAfterPart_symbolCreateCallback.OnAfterCreate(stage, target)
		}
	case *Part_transpose:
		if stage.OnAfterPart_transposeCreateCallback != nil {
			stage.OnAfterPart_transposeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Pedal:
		if stage.OnAfterPedalCreateCallback != nil {
			stage.OnAfterPedalCreateCallback.OnAfterCreate(stage, target)
		}
	case *Pedal_tuning:
		if stage.OnAfterPedal_tuningCreateCallback != nil {
			stage.OnAfterPedal_tuningCreateCallback.OnAfterCreate(stage, target)
		}
	case *Per_minute:
		if stage.OnAfterPer_minuteCreateCallback != nil {
			stage.OnAfterPer_minuteCreateCallback.OnAfterCreate(stage, target)
		}
	case *Percussion:
		if stage.OnAfterPercussionCreateCallback != nil {
			stage.OnAfterPercussionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Pitch:
		if stage.OnAfterPitchCreateCallback != nil {
			stage.OnAfterPitchCreateCallback.OnAfterCreate(stage, target)
		}
	case *Pitched:
		if stage.OnAfterPitchedCreateCallback != nil {
			stage.OnAfterPitchedCreateCallback.OnAfterCreate(stage, target)
		}
	case *Placement_text:
		if stage.OnAfterPlacement_textCreateCallback != nil {
			stage.OnAfterPlacement_textCreateCallback.OnAfterCreate(stage, target)
		}
	case *Play:
		if stage.OnAfterPlayCreateCallback != nil {
			stage.OnAfterPlayCreateCallback.OnAfterCreate(stage, target)
		}
	case *Player:
		if stage.OnAfterPlayerCreateCallback != nil {
			stage.OnAfterPlayerCreateCallback.OnAfterCreate(stage, target)
		}
	case *Principal_voice:
		if stage.OnAfterPrincipal_voiceCreateCallback != nil {
			stage.OnAfterPrincipal_voiceCreateCallback.OnAfterCreate(stage, target)
		}
	case *Print:
		if stage.OnAfterPrintCreateCallback != nil {
			stage.OnAfterPrintCreateCallback.OnAfterCreate(stage, target)
		}
	case *Release:
		if stage.OnAfterReleaseCreateCallback != nil {
			stage.OnAfterReleaseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Repeat:
		if stage.OnAfterRepeatCreateCallback != nil {
			stage.OnAfterRepeatCreateCallback.OnAfterCreate(stage, target)
		}
	case *Rest:
		if stage.OnAfterRestCreateCallback != nil {
			stage.OnAfterRestCreateCallback.OnAfterCreate(stage, target)
		}
	case *Root:
		if stage.OnAfterRootCreateCallback != nil {
			stage.OnAfterRootCreateCallback.OnAfterCreate(stage, target)
		}
	case *Root_step:
		if stage.OnAfterRoot_stepCreateCallback != nil {
			stage.OnAfterRoot_stepCreateCallback.OnAfterCreate(stage, target)
		}
	case *Scaling:
		if stage.OnAfterScalingCreateCallback != nil {
			stage.OnAfterScalingCreateCallback.OnAfterCreate(stage, target)
		}
	case *Scordatura:
		if stage.OnAfterScordaturaCreateCallback != nil {
			stage.OnAfterScordaturaCreateCallback.OnAfterCreate(stage, target)
		}
	case *Score_instrument:
		if stage.OnAfterScore_instrumentCreateCallback != nil {
			stage.OnAfterScore_instrumentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Score_part:
		if stage.OnAfterScore_partCreateCallback != nil {
			stage.OnAfterScore_partCreateCallback.OnAfterCreate(stage, target)
		}
	case *Score_partwise:
		if stage.OnAfterScore_partwiseCreateCallback != nil {
			stage.OnAfterScore_partwiseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Score_timewise:
		if stage.OnAfterScore_timewiseCreateCallback != nil {
			stage.OnAfterScore_timewiseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Segno:
		if stage.OnAfterSegnoCreateCallback != nil {
			stage.OnAfterSegnoCreateCallback.OnAfterCreate(stage, target)
		}
	case *Slash:
		if stage.OnAfterSlashCreateCallback != nil {
			stage.OnAfterSlashCreateCallback.OnAfterCreate(stage, target)
		}
	case *Slide:
		if stage.OnAfterSlideCreateCallback != nil {
			stage.OnAfterSlideCreateCallback.OnAfterCreate(stage, target)
		}
	case *Slur:
		if stage.OnAfterSlurCreateCallback != nil {
			stage.OnAfterSlurCreateCallback.OnAfterCreate(stage, target)
		}
	case *Sound:
		if stage.OnAfterSoundCreateCallback != nil {
			stage.OnAfterSoundCreateCallback.OnAfterCreate(stage, target)
		}
	case *Staff_details:
		if stage.OnAfterStaff_detailsCreateCallback != nil {
			stage.OnAfterStaff_detailsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Staff_divide:
		if stage.OnAfterStaff_divideCreateCallback != nil {
			stage.OnAfterStaff_divideCreateCallback.OnAfterCreate(stage, target)
		}
	case *Staff_layout:
		if stage.OnAfterStaff_layoutCreateCallback != nil {
			stage.OnAfterStaff_layoutCreateCallback.OnAfterCreate(stage, target)
		}
	case *Staff_size:
		if stage.OnAfterStaff_sizeCreateCallback != nil {
			stage.OnAfterStaff_sizeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Staff_tuning:
		if stage.OnAfterStaff_tuningCreateCallback != nil {
			stage.OnAfterStaff_tuningCreateCallback.OnAfterCreate(stage, target)
		}
	case *Stem:
		if stage.OnAfterStemCreateCallback != nil {
			stage.OnAfterStemCreateCallback.OnAfterCreate(stage, target)
		}
	case *Stick:
		if stage.OnAfterStickCreateCallback != nil {
			stage.OnAfterStickCreateCallback.OnAfterCreate(stage, target)
		}
	case *String_mute:
		if stage.OnAfterString_muteCreateCallback != nil {
			stage.OnAfterString_muteCreateCallback.OnAfterCreate(stage, target)
		}
	case *String_type:
		if stage.OnAfterString_typeCreateCallback != nil {
			stage.OnAfterString_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Strong_accent:
		if stage.OnAfterStrong_accentCreateCallback != nil {
			stage.OnAfterStrong_accentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Style_text:
		if stage.OnAfterStyle_textCreateCallback != nil {
			stage.OnAfterStyle_textCreateCallback.OnAfterCreate(stage, target)
		}
	case *Supports:
		if stage.OnAfterSupportsCreateCallback != nil {
			stage.OnAfterSupportsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Swing:
		if stage.OnAfterSwingCreateCallback != nil {
			stage.OnAfterSwingCreateCallback.OnAfterCreate(stage, target)
		}
	case *Sync:
		if stage.OnAfterSyncCreateCallback != nil {
			stage.OnAfterSyncCreateCallback.OnAfterCreate(stage, target)
		}
	case *System_dividers:
		if stage.OnAfterSystem_dividersCreateCallback != nil {
			stage.OnAfterSystem_dividersCreateCallback.OnAfterCreate(stage, target)
		}
	case *System_layout:
		if stage.OnAfterSystem_layoutCreateCallback != nil {
			stage.OnAfterSystem_layoutCreateCallback.OnAfterCreate(stage, target)
		}
	case *System_margins:
		if stage.OnAfterSystem_marginsCreateCallback != nil {
			stage.OnAfterSystem_marginsCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tap:
		if stage.OnAfterTapCreateCallback != nil {
			stage.OnAfterTapCreateCallback.OnAfterCreate(stage, target)
		}
	case *Technical:
		if stage.OnAfterTechnicalCreateCallback != nil {
			stage.OnAfterTechnicalCreateCallback.OnAfterCreate(stage, target)
		}
	case *Text_element_data:
		if stage.OnAfterText_element_dataCreateCallback != nil {
			stage.OnAfterText_element_dataCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tie:
		if stage.OnAfterTieCreateCallback != nil {
			stage.OnAfterTieCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tied:
		if stage.OnAfterTiedCreateCallback != nil {
			stage.OnAfterTiedCreateCallback.OnAfterCreate(stage, target)
		}
	case *Time:
		if stage.OnAfterTimeCreateCallback != nil {
			stage.OnAfterTimeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Time_modification:
		if stage.OnAfterTime_modificationCreateCallback != nil {
			stage.OnAfterTime_modificationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Timpani:
		if stage.OnAfterTimpaniCreateCallback != nil {
			stage.OnAfterTimpaniCreateCallback.OnAfterCreate(stage, target)
		}
	case *Transpose:
		if stage.OnAfterTransposeCreateCallback != nil {
			stage.OnAfterTransposeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tremolo:
		if stage.OnAfterTremoloCreateCallback != nil {
			stage.OnAfterTremoloCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tuplet:
		if stage.OnAfterTupletCreateCallback != nil {
			stage.OnAfterTupletCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tuplet_dot:
		if stage.OnAfterTuplet_dotCreateCallback != nil {
			stage.OnAfterTuplet_dotCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tuplet_number:
		if stage.OnAfterTuplet_numberCreateCallback != nil {
			stage.OnAfterTuplet_numberCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tuplet_portion:
		if stage.OnAfterTuplet_portionCreateCallback != nil {
			stage.OnAfterTuplet_portionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tuplet_type:
		if stage.OnAfterTuplet_typeCreateCallback != nil {
			stage.OnAfterTuplet_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Typed_text:
		if stage.OnAfterTyped_textCreateCallback != nil {
			stage.OnAfterTyped_textCreateCallback.OnAfterCreate(stage, target)
		}
	case *Unpitched:
		if stage.OnAfterUnpitchedCreateCallback != nil {
			stage.OnAfterUnpitchedCreateCallback.OnAfterCreate(stage, target)
		}
	case *Virtual_instrument:
		if stage.OnAfterVirtual_instrumentCreateCallback != nil {
			stage.OnAfterVirtual_instrumentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Wait:
		if stage.OnAfterWaitCreateCallback != nil {
			stage.OnAfterWaitCreateCallback.OnAfterCreate(stage, target)
		}
	case *Wavy_line:
		if stage.OnAfterWavy_lineCreateCallback != nil {
			stage.OnAfterWavy_lineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Wedge:
		if stage.OnAfterWedgeCreateCallback != nil {
			stage.OnAfterWedgeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Wood:
		if stage.OnAfterWoodCreateCallback != nil {
			stage.OnAfterWoodCreateCallback.OnAfterCreate(stage, target)
		}
	case *Work:
		if stage.OnAfterWorkCreateCallback != nil {
			stage.OnAfterWorkCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *A_directive:
		newTarget := any(new).(*A_directive)
		if stage.OnAfterA_directiveUpdateCallback != nil {
			stage.OnAfterA_directiveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_measure:
		newTarget := any(new).(*A_measure)
		if stage.OnAfterA_measureUpdateCallback != nil {
			stage.OnAfterA_measureUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_measure_1:
		newTarget := any(new).(*A_measure_1)
		if stage.OnAfterA_measure_1UpdateCallback != nil {
			stage.OnAfterA_measure_1UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_part:
		newTarget := any(new).(*A_part)
		if stage.OnAfterA_partUpdateCallback != nil {
			stage.OnAfterA_partUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_part_1:
		newTarget := any(new).(*A_part_1)
		if stage.OnAfterA_part_1UpdateCallback != nil {
			stage.OnAfterA_part_1UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Accidental:
		newTarget := any(new).(*Accidental)
		if stage.OnAfterAccidentalUpdateCallback != nil {
			stage.OnAfterAccidentalUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Accidental_mark:
		newTarget := any(new).(*Accidental_mark)
		if stage.OnAfterAccidental_markUpdateCallback != nil {
			stage.OnAfterAccidental_markUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Accidental_text:
		newTarget := any(new).(*Accidental_text)
		if stage.OnAfterAccidental_textUpdateCallback != nil {
			stage.OnAfterAccidental_textUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Accord:
		newTarget := any(new).(*Accord)
		if stage.OnAfterAccordUpdateCallback != nil {
			stage.OnAfterAccordUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Accordion_registration:
		newTarget := any(new).(*Accordion_registration)
		if stage.OnAfterAccordion_registrationUpdateCallback != nil {
			stage.OnAfterAccordion_registrationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Appearance:
		newTarget := any(new).(*Appearance)
		if stage.OnAfterAppearanceUpdateCallback != nil {
			stage.OnAfterAppearanceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Arpeggiate:
		newTarget := any(new).(*Arpeggiate)
		if stage.OnAfterArpeggiateUpdateCallback != nil {
			stage.OnAfterArpeggiateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Arrow:
		newTarget := any(new).(*Arrow)
		if stage.OnAfterArrowUpdateCallback != nil {
			stage.OnAfterArrowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Articulations:
		newTarget := any(new).(*Articulations)
		if stage.OnAfterArticulationsUpdateCallback != nil {
			stage.OnAfterArticulationsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Assess:
		newTarget := any(new).(*Assess)
		if stage.OnAfterAssessUpdateCallback != nil {
			stage.OnAfterAssessUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Attributes:
		newTarget := any(new).(*Attributes)
		if stage.OnAfterAttributesUpdateCallback != nil {
			stage.OnAfterAttributesUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Backup:
		newTarget := any(new).(*Backup)
		if stage.OnAfterBackupUpdateCallback != nil {
			stage.OnAfterBackupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Bar_style_color:
		newTarget := any(new).(*Bar_style_color)
		if stage.OnAfterBar_style_colorUpdateCallback != nil {
			stage.OnAfterBar_style_colorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Barline:
		newTarget := any(new).(*Barline)
		if stage.OnAfterBarlineUpdateCallback != nil {
			stage.OnAfterBarlineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Barre:
		newTarget := any(new).(*Barre)
		if stage.OnAfterBarreUpdateCallback != nil {
			stage.OnAfterBarreUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Bass:
		newTarget := any(new).(*Bass)
		if stage.OnAfterBassUpdateCallback != nil {
			stage.OnAfterBassUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Bass_step:
		newTarget := any(new).(*Bass_step)
		if stage.OnAfterBass_stepUpdateCallback != nil {
			stage.OnAfterBass_stepUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Beam:
		newTarget := any(new).(*Beam)
		if stage.OnAfterBeamUpdateCallback != nil {
			stage.OnAfterBeamUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Beat_repeat:
		newTarget := any(new).(*Beat_repeat)
		if stage.OnAfterBeat_repeatUpdateCallback != nil {
			stage.OnAfterBeat_repeatUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Beat_unit_tied:
		newTarget := any(new).(*Beat_unit_tied)
		if stage.OnAfterBeat_unit_tiedUpdateCallback != nil {
			stage.OnAfterBeat_unit_tiedUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Beater:
		newTarget := any(new).(*Beater)
		if stage.OnAfterBeaterUpdateCallback != nil {
			stage.OnAfterBeaterUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Bend:
		newTarget := any(new).(*Bend)
		if stage.OnAfterBendUpdateCallback != nil {
			stage.OnAfterBendUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Bookmark:
		newTarget := any(new).(*Bookmark)
		if stage.OnAfterBookmarkUpdateCallback != nil {
			stage.OnAfterBookmarkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Bracket:
		newTarget := any(new).(*Bracket)
		if stage.OnAfterBracketUpdateCallback != nil {
			stage.OnAfterBracketUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Breath_mark:
		newTarget := any(new).(*Breath_mark)
		if stage.OnAfterBreath_markUpdateCallback != nil {
			stage.OnAfterBreath_markUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Caesura:
		newTarget := any(new).(*Caesura)
		if stage.OnAfterCaesuraUpdateCallback != nil {
			stage.OnAfterCaesuraUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Cancel:
		newTarget := any(new).(*Cancel)
		if stage.OnAfterCancelUpdateCallback != nil {
			stage.OnAfterCancelUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Clef:
		newTarget := any(new).(*Clef)
		if stage.OnAfterClefUpdateCallback != nil {
			stage.OnAfterClefUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Coda:
		newTarget := any(new).(*Coda)
		if stage.OnAfterCodaUpdateCallback != nil {
			stage.OnAfterCodaUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Credit:
		newTarget := any(new).(*Credit)
		if stage.OnAfterCreditUpdateCallback != nil {
			stage.OnAfterCreditUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Dashes:
		newTarget := any(new).(*Dashes)
		if stage.OnAfterDashesUpdateCallback != nil {
			stage.OnAfterDashesUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Defaults:
		newTarget := any(new).(*Defaults)
		if stage.OnAfterDefaultsUpdateCallback != nil {
			stage.OnAfterDefaultsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Degree:
		newTarget := any(new).(*Degree)
		if stage.OnAfterDegreeUpdateCallback != nil {
			stage.OnAfterDegreeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Degree_alter:
		newTarget := any(new).(*Degree_alter)
		if stage.OnAfterDegree_alterUpdateCallback != nil {
			stage.OnAfterDegree_alterUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Degree_type:
		newTarget := any(new).(*Degree_type)
		if stage.OnAfterDegree_typeUpdateCallback != nil {
			stage.OnAfterDegree_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Degree_value:
		newTarget := any(new).(*Degree_value)
		if stage.OnAfterDegree_valueUpdateCallback != nil {
			stage.OnAfterDegree_valueUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Direction:
		newTarget := any(new).(*Direction)
		if stage.OnAfterDirectionUpdateCallback != nil {
			stage.OnAfterDirectionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Direction_type:
		newTarget := any(new).(*Direction_type)
		if stage.OnAfterDirection_typeUpdateCallback != nil {
			stage.OnAfterDirection_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Distance:
		newTarget := any(new).(*Distance)
		if stage.OnAfterDistanceUpdateCallback != nil {
			stage.OnAfterDistanceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Double:
		newTarget := any(new).(*Double)
		if stage.OnAfterDoubleUpdateCallback != nil {
			stage.OnAfterDoubleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Dynamics:
		newTarget := any(new).(*Dynamics)
		if stage.OnAfterDynamicsUpdateCallback != nil {
			stage.OnAfterDynamicsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Effect:
		newTarget := any(new).(*Effect)
		if stage.OnAfterEffectUpdateCallback != nil {
			stage.OnAfterEffectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Elision:
		newTarget := any(new).(*Elision)
		if stage.OnAfterElisionUpdateCallback != nil {
			stage.OnAfterElisionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty:
		newTarget := any(new).(*Empty)
		if stage.OnAfterEmptyUpdateCallback != nil {
			stage.OnAfterEmptyUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty_font:
		newTarget := any(new).(*Empty_font)
		if stage.OnAfterEmpty_fontUpdateCallback != nil {
			stage.OnAfterEmpty_fontUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty_line:
		newTarget := any(new).(*Empty_line)
		if stage.OnAfterEmpty_lineUpdateCallback != nil {
			stage.OnAfterEmpty_lineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty_placement:
		newTarget := any(new).(*Empty_placement)
		if stage.OnAfterEmpty_placementUpdateCallback != nil {
			stage.OnAfterEmpty_placementUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty_placement_smufl:
		newTarget := any(new).(*Empty_placement_smufl)
		if stage.OnAfterEmpty_placement_smuflUpdateCallback != nil {
			stage.OnAfterEmpty_placement_smuflUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty_print_object_style_align:
		newTarget := any(new).(*Empty_print_object_style_align)
		if stage.OnAfterEmpty_print_object_style_alignUpdateCallback != nil {
			stage.OnAfterEmpty_print_object_style_alignUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty_print_style:
		newTarget := any(new).(*Empty_print_style)
		if stage.OnAfterEmpty_print_styleUpdateCallback != nil {
			stage.OnAfterEmpty_print_styleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty_print_style_align:
		newTarget := any(new).(*Empty_print_style_align)
		if stage.OnAfterEmpty_print_style_alignUpdateCallback != nil {
			stage.OnAfterEmpty_print_style_alignUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty_print_style_align_id:
		newTarget := any(new).(*Empty_print_style_align_id)
		if stage.OnAfterEmpty_print_style_align_idUpdateCallback != nil {
			stage.OnAfterEmpty_print_style_align_idUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Empty_trill_sound:
		newTarget := any(new).(*Empty_trill_sound)
		if stage.OnAfterEmpty_trill_soundUpdateCallback != nil {
			stage.OnAfterEmpty_trill_soundUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Encoding:
		newTarget := any(new).(*Encoding)
		if stage.OnAfterEncodingUpdateCallback != nil {
			stage.OnAfterEncodingUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Ending:
		newTarget := any(new).(*Ending)
		if stage.OnAfterEndingUpdateCallback != nil {
			stage.OnAfterEndingUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Extend:
		newTarget := any(new).(*Extend)
		if stage.OnAfterExtendUpdateCallback != nil {
			stage.OnAfterExtendUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Feature:
		newTarget := any(new).(*Feature)
		if stage.OnAfterFeatureUpdateCallback != nil {
			stage.OnAfterFeatureUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Fermata:
		newTarget := any(new).(*Fermata)
		if stage.OnAfterFermataUpdateCallback != nil {
			stage.OnAfterFermataUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Figure:
		newTarget := any(new).(*Figure)
		if stage.OnAfterFigureUpdateCallback != nil {
			stage.OnAfterFigureUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Figured_bass:
		newTarget := any(new).(*Figured_bass)
		if stage.OnAfterFigured_bassUpdateCallback != nil {
			stage.OnAfterFigured_bassUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Fingering:
		newTarget := any(new).(*Fingering)
		if stage.OnAfterFingeringUpdateCallback != nil {
			stage.OnAfterFingeringUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *First_fret:
		newTarget := any(new).(*First_fret)
		if stage.OnAfterFirst_fretUpdateCallback != nil {
			stage.OnAfterFirst_fretUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *For_part:
		newTarget := any(new).(*For_part)
		if stage.OnAfterFor_partUpdateCallback != nil {
			stage.OnAfterFor_partUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Formatted_symbol:
		newTarget := any(new).(*Formatted_symbol)
		if stage.OnAfterFormatted_symbolUpdateCallback != nil {
			stage.OnAfterFormatted_symbolUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Formatted_symbol_id:
		newTarget := any(new).(*Formatted_symbol_id)
		if stage.OnAfterFormatted_symbol_idUpdateCallback != nil {
			stage.OnAfterFormatted_symbol_idUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Formatted_text:
		newTarget := any(new).(*Formatted_text)
		if stage.OnAfterFormatted_textUpdateCallback != nil {
			stage.OnAfterFormatted_textUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Formatted_text_id:
		newTarget := any(new).(*Formatted_text_id)
		if stage.OnAfterFormatted_text_idUpdateCallback != nil {
			stage.OnAfterFormatted_text_idUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Forward:
		newTarget := any(new).(*Forward)
		if stage.OnAfterForwardUpdateCallback != nil {
			stage.OnAfterForwardUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Frame:
		newTarget := any(new).(*Frame)
		if stage.OnAfterFrameUpdateCallback != nil {
			stage.OnAfterFrameUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Frame_note:
		newTarget := any(new).(*Frame_note)
		if stage.OnAfterFrame_noteUpdateCallback != nil {
			stage.OnAfterFrame_noteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Fret:
		newTarget := any(new).(*Fret)
		if stage.OnAfterFretUpdateCallback != nil {
			stage.OnAfterFretUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Glass:
		newTarget := any(new).(*Glass)
		if stage.OnAfterGlassUpdateCallback != nil {
			stage.OnAfterGlassUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Glissando:
		newTarget := any(new).(*Glissando)
		if stage.OnAfterGlissandoUpdateCallback != nil {
			stage.OnAfterGlissandoUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Glyph:
		newTarget := any(new).(*Glyph)
		if stage.OnAfterGlyphUpdateCallback != nil {
			stage.OnAfterGlyphUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Grace:
		newTarget := any(new).(*Grace)
		if stage.OnAfterGraceUpdateCallback != nil {
			stage.OnAfterGraceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Group_barline:
		newTarget := any(new).(*Group_barline)
		if stage.OnAfterGroup_barlineUpdateCallback != nil {
			stage.OnAfterGroup_barlineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Group_name:
		newTarget := any(new).(*Group_name)
		if stage.OnAfterGroup_nameUpdateCallback != nil {
			stage.OnAfterGroup_nameUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Group_symbol:
		newTarget := any(new).(*Group_symbol)
		if stage.OnAfterGroup_symbolUpdateCallback != nil {
			stage.OnAfterGroup_symbolUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Grouping:
		newTarget := any(new).(*Grouping)
		if stage.OnAfterGroupingUpdateCallback != nil {
			stage.OnAfterGroupingUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Hammer_on_pull_off:
		newTarget := any(new).(*Hammer_on_pull_off)
		if stage.OnAfterHammer_on_pull_offUpdateCallback != nil {
			stage.OnAfterHammer_on_pull_offUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Handbell:
		newTarget := any(new).(*Handbell)
		if stage.OnAfterHandbellUpdateCallback != nil {
			stage.OnAfterHandbellUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Harmon_closed:
		newTarget := any(new).(*Harmon_closed)
		if stage.OnAfterHarmon_closedUpdateCallback != nil {
			stage.OnAfterHarmon_closedUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Harmon_mute:
		newTarget := any(new).(*Harmon_mute)
		if stage.OnAfterHarmon_muteUpdateCallback != nil {
			stage.OnAfterHarmon_muteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Harmonic:
		newTarget := any(new).(*Harmonic)
		if stage.OnAfterHarmonicUpdateCallback != nil {
			stage.OnAfterHarmonicUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Harmony:
		newTarget := any(new).(*Harmony)
		if stage.OnAfterHarmonyUpdateCallback != nil {
			stage.OnAfterHarmonyUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Harmony_alter:
		newTarget := any(new).(*Harmony_alter)
		if stage.OnAfterHarmony_alterUpdateCallback != nil {
			stage.OnAfterHarmony_alterUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Harp_pedals:
		newTarget := any(new).(*Harp_pedals)
		if stage.OnAfterHarp_pedalsUpdateCallback != nil {
			stage.OnAfterHarp_pedalsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Heel_toe:
		newTarget := any(new).(*Heel_toe)
		if stage.OnAfterHeel_toeUpdateCallback != nil {
			stage.OnAfterHeel_toeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Hole:
		newTarget := any(new).(*Hole)
		if stage.OnAfterHoleUpdateCallback != nil {
			stage.OnAfterHoleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Hole_closed:
		newTarget := any(new).(*Hole_closed)
		if stage.OnAfterHole_closedUpdateCallback != nil {
			stage.OnAfterHole_closedUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Horizontal_turn:
		newTarget := any(new).(*Horizontal_turn)
		if stage.OnAfterHorizontal_turnUpdateCallback != nil {
			stage.OnAfterHorizontal_turnUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Identification:
		newTarget := any(new).(*Identification)
		if stage.OnAfterIdentificationUpdateCallback != nil {
			stage.OnAfterIdentificationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Image:
		newTarget := any(new).(*Image)
		if stage.OnAfterImageUpdateCallback != nil {
			stage.OnAfterImageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Instrument:
		newTarget := any(new).(*Instrument)
		if stage.OnAfterInstrumentUpdateCallback != nil {
			stage.OnAfterInstrumentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Instrument_change:
		newTarget := any(new).(*Instrument_change)
		if stage.OnAfterInstrument_changeUpdateCallback != nil {
			stage.OnAfterInstrument_changeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Instrument_link:
		newTarget := any(new).(*Instrument_link)
		if stage.OnAfterInstrument_linkUpdateCallback != nil {
			stage.OnAfterInstrument_linkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Interchangeable:
		newTarget := any(new).(*Interchangeable)
		if stage.OnAfterInterchangeableUpdateCallback != nil {
			stage.OnAfterInterchangeableUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Inversion:
		newTarget := any(new).(*Inversion)
		if stage.OnAfterInversionUpdateCallback != nil {
			stage.OnAfterInversionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Key:
		newTarget := any(new).(*Key)
		if stage.OnAfterKeyUpdateCallback != nil {
			stage.OnAfterKeyUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Key_accidental:
		newTarget := any(new).(*Key_accidental)
		if stage.OnAfterKey_accidentalUpdateCallback != nil {
			stage.OnAfterKey_accidentalUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Key_octave:
		newTarget := any(new).(*Key_octave)
		if stage.OnAfterKey_octaveUpdateCallback != nil {
			stage.OnAfterKey_octaveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Kind:
		newTarget := any(new).(*Kind)
		if stage.OnAfterKindUpdateCallback != nil {
			stage.OnAfterKindUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Level:
		newTarget := any(new).(*Level)
		if stage.OnAfterLevelUpdateCallback != nil {
			stage.OnAfterLevelUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Line_detail:
		newTarget := any(new).(*Line_detail)
		if stage.OnAfterLine_detailUpdateCallback != nil {
			stage.OnAfterLine_detailUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Line_width:
		newTarget := any(new).(*Line_width)
		if stage.OnAfterLine_widthUpdateCallback != nil {
			stage.OnAfterLine_widthUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Link:
		newTarget := any(new).(*Link)
		if stage.OnAfterLinkUpdateCallback != nil {
			stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Listen:
		newTarget := any(new).(*Listen)
		if stage.OnAfterListenUpdateCallback != nil {
			stage.OnAfterListenUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Listening:
		newTarget := any(new).(*Listening)
		if stage.OnAfterListeningUpdateCallback != nil {
			stage.OnAfterListeningUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Lyric:
		newTarget := any(new).(*Lyric)
		if stage.OnAfterLyricUpdateCallback != nil {
			stage.OnAfterLyricUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Lyric_font:
		newTarget := any(new).(*Lyric_font)
		if stage.OnAfterLyric_fontUpdateCallback != nil {
			stage.OnAfterLyric_fontUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Lyric_language:
		newTarget := any(new).(*Lyric_language)
		if stage.OnAfterLyric_languageUpdateCallback != nil {
			stage.OnAfterLyric_languageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Measure_layout:
		newTarget := any(new).(*Measure_layout)
		if stage.OnAfterMeasure_layoutUpdateCallback != nil {
			stage.OnAfterMeasure_layoutUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Measure_numbering:
		newTarget := any(new).(*Measure_numbering)
		if stage.OnAfterMeasure_numberingUpdateCallback != nil {
			stage.OnAfterMeasure_numberingUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Measure_repeat:
		newTarget := any(new).(*Measure_repeat)
		if stage.OnAfterMeasure_repeatUpdateCallback != nil {
			stage.OnAfterMeasure_repeatUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Measure_style:
		newTarget := any(new).(*Measure_style)
		if stage.OnAfterMeasure_styleUpdateCallback != nil {
			stage.OnAfterMeasure_styleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Membrane:
		newTarget := any(new).(*Membrane)
		if stage.OnAfterMembraneUpdateCallback != nil {
			stage.OnAfterMembraneUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Metal:
		newTarget := any(new).(*Metal)
		if stage.OnAfterMetalUpdateCallback != nil {
			stage.OnAfterMetalUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Metronome:
		newTarget := any(new).(*Metronome)
		if stage.OnAfterMetronomeUpdateCallback != nil {
			stage.OnAfterMetronomeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Metronome_beam:
		newTarget := any(new).(*Metronome_beam)
		if stage.OnAfterMetronome_beamUpdateCallback != nil {
			stage.OnAfterMetronome_beamUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Metronome_note:
		newTarget := any(new).(*Metronome_note)
		if stage.OnAfterMetronome_noteUpdateCallback != nil {
			stage.OnAfterMetronome_noteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Metronome_tied:
		newTarget := any(new).(*Metronome_tied)
		if stage.OnAfterMetronome_tiedUpdateCallback != nil {
			stage.OnAfterMetronome_tiedUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Metronome_tuplet:
		newTarget := any(new).(*Metronome_tuplet)
		if stage.OnAfterMetronome_tupletUpdateCallback != nil {
			stage.OnAfterMetronome_tupletUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Midi_device:
		newTarget := any(new).(*Midi_device)
		if stage.OnAfterMidi_deviceUpdateCallback != nil {
			stage.OnAfterMidi_deviceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Midi_instrument:
		newTarget := any(new).(*Midi_instrument)
		if stage.OnAfterMidi_instrumentUpdateCallback != nil {
			stage.OnAfterMidi_instrumentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Miscellaneous:
		newTarget := any(new).(*Miscellaneous)
		if stage.OnAfterMiscellaneousUpdateCallback != nil {
			stage.OnAfterMiscellaneousUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Miscellaneous_field:
		newTarget := any(new).(*Miscellaneous_field)
		if stage.OnAfterMiscellaneous_fieldUpdateCallback != nil {
			stage.OnAfterMiscellaneous_fieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Mordent:
		newTarget := any(new).(*Mordent)
		if stage.OnAfterMordentUpdateCallback != nil {
			stage.OnAfterMordentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Multiple_rest:
		newTarget := any(new).(*Multiple_rest)
		if stage.OnAfterMultiple_restUpdateCallback != nil {
			stage.OnAfterMultiple_restUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Name_display:
		newTarget := any(new).(*Name_display)
		if stage.OnAfterName_displayUpdateCallback != nil {
			stage.OnAfterName_displayUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Non_arpeggiate:
		newTarget := any(new).(*Non_arpeggiate)
		if stage.OnAfterNon_arpeggiateUpdateCallback != nil {
			stage.OnAfterNon_arpeggiateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Notations:
		newTarget := any(new).(*Notations)
		if stage.OnAfterNotationsUpdateCallback != nil {
			stage.OnAfterNotationsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Note:
		newTarget := any(new).(*Note)
		if stage.OnAfterNoteUpdateCallback != nil {
			stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Note_size:
		newTarget := any(new).(*Note_size)
		if stage.OnAfterNote_sizeUpdateCallback != nil {
			stage.OnAfterNote_sizeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Note_type:
		newTarget := any(new).(*Note_type)
		if stage.OnAfterNote_typeUpdateCallback != nil {
			stage.OnAfterNote_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Notehead:
		newTarget := any(new).(*Notehead)
		if stage.OnAfterNoteheadUpdateCallback != nil {
			stage.OnAfterNoteheadUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Notehead_text:
		newTarget := any(new).(*Notehead_text)
		if stage.OnAfterNotehead_textUpdateCallback != nil {
			stage.OnAfterNotehead_textUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Numeral:
		newTarget := any(new).(*Numeral)
		if stage.OnAfterNumeralUpdateCallback != nil {
			stage.OnAfterNumeralUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Numeral_key:
		newTarget := any(new).(*Numeral_key)
		if stage.OnAfterNumeral_keyUpdateCallback != nil {
			stage.OnAfterNumeral_keyUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Numeral_root:
		newTarget := any(new).(*Numeral_root)
		if stage.OnAfterNumeral_rootUpdateCallback != nil {
			stage.OnAfterNumeral_rootUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Octave_shift:
		newTarget := any(new).(*Octave_shift)
		if stage.OnAfterOctave_shiftUpdateCallback != nil {
			stage.OnAfterOctave_shiftUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Offset:
		newTarget := any(new).(*Offset)
		if stage.OnAfterOffsetUpdateCallback != nil {
			stage.OnAfterOffsetUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Opus:
		newTarget := any(new).(*Opus)
		if stage.OnAfterOpusUpdateCallback != nil {
			stage.OnAfterOpusUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Ornaments:
		newTarget := any(new).(*Ornaments)
		if stage.OnAfterOrnamentsUpdateCallback != nil {
			stage.OnAfterOrnamentsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Other_appearance:
		newTarget := any(new).(*Other_appearance)
		if stage.OnAfterOther_appearanceUpdateCallback != nil {
			stage.OnAfterOther_appearanceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Other_direction:
		newTarget := any(new).(*Other_direction)
		if stage.OnAfterOther_directionUpdateCallback != nil {
			stage.OnAfterOther_directionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Other_listening:
		newTarget := any(new).(*Other_listening)
		if stage.OnAfterOther_listeningUpdateCallback != nil {
			stage.OnAfterOther_listeningUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Other_notation:
		newTarget := any(new).(*Other_notation)
		if stage.OnAfterOther_notationUpdateCallback != nil {
			stage.OnAfterOther_notationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Other_placement_text:
		newTarget := any(new).(*Other_placement_text)
		if stage.OnAfterOther_placement_textUpdateCallback != nil {
			stage.OnAfterOther_placement_textUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Other_play:
		newTarget := any(new).(*Other_play)
		if stage.OnAfterOther_playUpdateCallback != nil {
			stage.OnAfterOther_playUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Other_text:
		newTarget := any(new).(*Other_text)
		if stage.OnAfterOther_textUpdateCallback != nil {
			stage.OnAfterOther_textUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Page_layout:
		newTarget := any(new).(*Page_layout)
		if stage.OnAfterPage_layoutUpdateCallback != nil {
			stage.OnAfterPage_layoutUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Page_margins:
		newTarget := any(new).(*Page_margins)
		if stage.OnAfterPage_marginsUpdateCallback != nil {
			stage.OnAfterPage_marginsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Part_clef:
		newTarget := any(new).(*Part_clef)
		if stage.OnAfterPart_clefUpdateCallback != nil {
			stage.OnAfterPart_clefUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Part_group:
		newTarget := any(new).(*Part_group)
		if stage.OnAfterPart_groupUpdateCallback != nil {
			stage.OnAfterPart_groupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Part_link:
		newTarget := any(new).(*Part_link)
		if stage.OnAfterPart_linkUpdateCallback != nil {
			stage.OnAfterPart_linkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Part_list:
		newTarget := any(new).(*Part_list)
		if stage.OnAfterPart_listUpdateCallback != nil {
			stage.OnAfterPart_listUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Part_name:
		newTarget := any(new).(*Part_name)
		if stage.OnAfterPart_nameUpdateCallback != nil {
			stage.OnAfterPart_nameUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Part_symbol:
		newTarget := any(new).(*Part_symbol)
		if stage.OnAfterPart_symbolUpdateCallback != nil {
			stage.OnAfterPart_symbolUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Part_transpose:
		newTarget := any(new).(*Part_transpose)
		if stage.OnAfterPart_transposeUpdateCallback != nil {
			stage.OnAfterPart_transposeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Pedal:
		newTarget := any(new).(*Pedal)
		if stage.OnAfterPedalUpdateCallback != nil {
			stage.OnAfterPedalUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Pedal_tuning:
		newTarget := any(new).(*Pedal_tuning)
		if stage.OnAfterPedal_tuningUpdateCallback != nil {
			stage.OnAfterPedal_tuningUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Per_minute:
		newTarget := any(new).(*Per_minute)
		if stage.OnAfterPer_minuteUpdateCallback != nil {
			stage.OnAfterPer_minuteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Percussion:
		newTarget := any(new).(*Percussion)
		if stage.OnAfterPercussionUpdateCallback != nil {
			stage.OnAfterPercussionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Pitch:
		newTarget := any(new).(*Pitch)
		if stage.OnAfterPitchUpdateCallback != nil {
			stage.OnAfterPitchUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Pitched:
		newTarget := any(new).(*Pitched)
		if stage.OnAfterPitchedUpdateCallback != nil {
			stage.OnAfterPitchedUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Placement_text:
		newTarget := any(new).(*Placement_text)
		if stage.OnAfterPlacement_textUpdateCallback != nil {
			stage.OnAfterPlacement_textUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Play:
		newTarget := any(new).(*Play)
		if stage.OnAfterPlayUpdateCallback != nil {
			stage.OnAfterPlayUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Player:
		newTarget := any(new).(*Player)
		if stage.OnAfterPlayerUpdateCallback != nil {
			stage.OnAfterPlayerUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Principal_voice:
		newTarget := any(new).(*Principal_voice)
		if stage.OnAfterPrincipal_voiceUpdateCallback != nil {
			stage.OnAfterPrincipal_voiceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Print:
		newTarget := any(new).(*Print)
		if stage.OnAfterPrintUpdateCallback != nil {
			stage.OnAfterPrintUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Release:
		newTarget := any(new).(*Release)
		if stage.OnAfterReleaseUpdateCallback != nil {
			stage.OnAfterReleaseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Repeat:
		newTarget := any(new).(*Repeat)
		if stage.OnAfterRepeatUpdateCallback != nil {
			stage.OnAfterRepeatUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Rest:
		newTarget := any(new).(*Rest)
		if stage.OnAfterRestUpdateCallback != nil {
			stage.OnAfterRestUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Root:
		newTarget := any(new).(*Root)
		if stage.OnAfterRootUpdateCallback != nil {
			stage.OnAfterRootUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Root_step:
		newTarget := any(new).(*Root_step)
		if stage.OnAfterRoot_stepUpdateCallback != nil {
			stage.OnAfterRoot_stepUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Scaling:
		newTarget := any(new).(*Scaling)
		if stage.OnAfterScalingUpdateCallback != nil {
			stage.OnAfterScalingUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Scordatura:
		newTarget := any(new).(*Scordatura)
		if stage.OnAfterScordaturaUpdateCallback != nil {
			stage.OnAfterScordaturaUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Score_instrument:
		newTarget := any(new).(*Score_instrument)
		if stage.OnAfterScore_instrumentUpdateCallback != nil {
			stage.OnAfterScore_instrumentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Score_part:
		newTarget := any(new).(*Score_part)
		if stage.OnAfterScore_partUpdateCallback != nil {
			stage.OnAfterScore_partUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Score_partwise:
		newTarget := any(new).(*Score_partwise)
		if stage.OnAfterScore_partwiseUpdateCallback != nil {
			stage.OnAfterScore_partwiseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Score_timewise:
		newTarget := any(new).(*Score_timewise)
		if stage.OnAfterScore_timewiseUpdateCallback != nil {
			stage.OnAfterScore_timewiseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Segno:
		newTarget := any(new).(*Segno)
		if stage.OnAfterSegnoUpdateCallback != nil {
			stage.OnAfterSegnoUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Slash:
		newTarget := any(new).(*Slash)
		if stage.OnAfterSlashUpdateCallback != nil {
			stage.OnAfterSlashUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Slide:
		newTarget := any(new).(*Slide)
		if stage.OnAfterSlideUpdateCallback != nil {
			stage.OnAfterSlideUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Slur:
		newTarget := any(new).(*Slur)
		if stage.OnAfterSlurUpdateCallback != nil {
			stage.OnAfterSlurUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Sound:
		newTarget := any(new).(*Sound)
		if stage.OnAfterSoundUpdateCallback != nil {
			stage.OnAfterSoundUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Staff_details:
		newTarget := any(new).(*Staff_details)
		if stage.OnAfterStaff_detailsUpdateCallback != nil {
			stage.OnAfterStaff_detailsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Staff_divide:
		newTarget := any(new).(*Staff_divide)
		if stage.OnAfterStaff_divideUpdateCallback != nil {
			stage.OnAfterStaff_divideUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Staff_layout:
		newTarget := any(new).(*Staff_layout)
		if stage.OnAfterStaff_layoutUpdateCallback != nil {
			stage.OnAfterStaff_layoutUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Staff_size:
		newTarget := any(new).(*Staff_size)
		if stage.OnAfterStaff_sizeUpdateCallback != nil {
			stage.OnAfterStaff_sizeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Staff_tuning:
		newTarget := any(new).(*Staff_tuning)
		if stage.OnAfterStaff_tuningUpdateCallback != nil {
			stage.OnAfterStaff_tuningUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Stem:
		newTarget := any(new).(*Stem)
		if stage.OnAfterStemUpdateCallback != nil {
			stage.OnAfterStemUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Stick:
		newTarget := any(new).(*Stick)
		if stage.OnAfterStickUpdateCallback != nil {
			stage.OnAfterStickUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *String_mute:
		newTarget := any(new).(*String_mute)
		if stage.OnAfterString_muteUpdateCallback != nil {
			stage.OnAfterString_muteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *String_type:
		newTarget := any(new).(*String_type)
		if stage.OnAfterString_typeUpdateCallback != nil {
			stage.OnAfterString_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Strong_accent:
		newTarget := any(new).(*Strong_accent)
		if stage.OnAfterStrong_accentUpdateCallback != nil {
			stage.OnAfterStrong_accentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Style_text:
		newTarget := any(new).(*Style_text)
		if stage.OnAfterStyle_textUpdateCallback != nil {
			stage.OnAfterStyle_textUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Supports:
		newTarget := any(new).(*Supports)
		if stage.OnAfterSupportsUpdateCallback != nil {
			stage.OnAfterSupportsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Swing:
		newTarget := any(new).(*Swing)
		if stage.OnAfterSwingUpdateCallback != nil {
			stage.OnAfterSwingUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Sync:
		newTarget := any(new).(*Sync)
		if stage.OnAfterSyncUpdateCallback != nil {
			stage.OnAfterSyncUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *System_dividers:
		newTarget := any(new).(*System_dividers)
		if stage.OnAfterSystem_dividersUpdateCallback != nil {
			stage.OnAfterSystem_dividersUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *System_layout:
		newTarget := any(new).(*System_layout)
		if stage.OnAfterSystem_layoutUpdateCallback != nil {
			stage.OnAfterSystem_layoutUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *System_margins:
		newTarget := any(new).(*System_margins)
		if stage.OnAfterSystem_marginsUpdateCallback != nil {
			stage.OnAfterSystem_marginsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tap:
		newTarget := any(new).(*Tap)
		if stage.OnAfterTapUpdateCallback != nil {
			stage.OnAfterTapUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Technical:
		newTarget := any(new).(*Technical)
		if stage.OnAfterTechnicalUpdateCallback != nil {
			stage.OnAfterTechnicalUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Text_element_data:
		newTarget := any(new).(*Text_element_data)
		if stage.OnAfterText_element_dataUpdateCallback != nil {
			stage.OnAfterText_element_dataUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tie:
		newTarget := any(new).(*Tie)
		if stage.OnAfterTieUpdateCallback != nil {
			stage.OnAfterTieUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tied:
		newTarget := any(new).(*Tied)
		if stage.OnAfterTiedUpdateCallback != nil {
			stage.OnAfterTiedUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Time:
		newTarget := any(new).(*Time)
		if stage.OnAfterTimeUpdateCallback != nil {
			stage.OnAfterTimeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Time_modification:
		newTarget := any(new).(*Time_modification)
		if stage.OnAfterTime_modificationUpdateCallback != nil {
			stage.OnAfterTime_modificationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Timpani:
		newTarget := any(new).(*Timpani)
		if stage.OnAfterTimpaniUpdateCallback != nil {
			stage.OnAfterTimpaniUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Transpose:
		newTarget := any(new).(*Transpose)
		if stage.OnAfterTransposeUpdateCallback != nil {
			stage.OnAfterTransposeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tremolo:
		newTarget := any(new).(*Tremolo)
		if stage.OnAfterTremoloUpdateCallback != nil {
			stage.OnAfterTremoloUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tuplet:
		newTarget := any(new).(*Tuplet)
		if stage.OnAfterTupletUpdateCallback != nil {
			stage.OnAfterTupletUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tuplet_dot:
		newTarget := any(new).(*Tuplet_dot)
		if stage.OnAfterTuplet_dotUpdateCallback != nil {
			stage.OnAfterTuplet_dotUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tuplet_number:
		newTarget := any(new).(*Tuplet_number)
		if stage.OnAfterTuplet_numberUpdateCallback != nil {
			stage.OnAfterTuplet_numberUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tuplet_portion:
		newTarget := any(new).(*Tuplet_portion)
		if stage.OnAfterTuplet_portionUpdateCallback != nil {
			stage.OnAfterTuplet_portionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tuplet_type:
		newTarget := any(new).(*Tuplet_type)
		if stage.OnAfterTuplet_typeUpdateCallback != nil {
			stage.OnAfterTuplet_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Typed_text:
		newTarget := any(new).(*Typed_text)
		if stage.OnAfterTyped_textUpdateCallback != nil {
			stage.OnAfterTyped_textUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Unpitched:
		newTarget := any(new).(*Unpitched)
		if stage.OnAfterUnpitchedUpdateCallback != nil {
			stage.OnAfterUnpitchedUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Virtual_instrument:
		newTarget := any(new).(*Virtual_instrument)
		if stage.OnAfterVirtual_instrumentUpdateCallback != nil {
			stage.OnAfterVirtual_instrumentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Wait:
		newTarget := any(new).(*Wait)
		if stage.OnAfterWaitUpdateCallback != nil {
			stage.OnAfterWaitUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Wavy_line:
		newTarget := any(new).(*Wavy_line)
		if stage.OnAfterWavy_lineUpdateCallback != nil {
			stage.OnAfterWavy_lineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Wedge:
		newTarget := any(new).(*Wedge)
		if stage.OnAfterWedgeUpdateCallback != nil {
			stage.OnAfterWedgeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Wood:
		newTarget := any(new).(*Wood)
		if stage.OnAfterWoodUpdateCallback != nil {
			stage.OnAfterWoodUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Work:
		newTarget := any(new).(*Work)
		if stage.OnAfterWorkUpdateCallback != nil {
			stage.OnAfterWorkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *A_directive:
		if stage.OnAfterA_directiveDeleteCallback != nil {
			staged := any(staged).(*A_directive)
			stage.OnAfterA_directiveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_measure:
		if stage.OnAfterA_measureDeleteCallback != nil {
			staged := any(staged).(*A_measure)
			stage.OnAfterA_measureDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_measure_1:
		if stage.OnAfterA_measure_1DeleteCallback != nil {
			staged := any(staged).(*A_measure_1)
			stage.OnAfterA_measure_1DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_part:
		if stage.OnAfterA_partDeleteCallback != nil {
			staged := any(staged).(*A_part)
			stage.OnAfterA_partDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_part_1:
		if stage.OnAfterA_part_1DeleteCallback != nil {
			staged := any(staged).(*A_part_1)
			stage.OnAfterA_part_1DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Accidental:
		if stage.OnAfterAccidentalDeleteCallback != nil {
			staged := any(staged).(*Accidental)
			stage.OnAfterAccidentalDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Accidental_mark:
		if stage.OnAfterAccidental_markDeleteCallback != nil {
			staged := any(staged).(*Accidental_mark)
			stage.OnAfterAccidental_markDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Accidental_text:
		if stage.OnAfterAccidental_textDeleteCallback != nil {
			staged := any(staged).(*Accidental_text)
			stage.OnAfterAccidental_textDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Accord:
		if stage.OnAfterAccordDeleteCallback != nil {
			staged := any(staged).(*Accord)
			stage.OnAfterAccordDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Accordion_registration:
		if stage.OnAfterAccordion_registrationDeleteCallback != nil {
			staged := any(staged).(*Accordion_registration)
			stage.OnAfterAccordion_registrationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Appearance:
		if stage.OnAfterAppearanceDeleteCallback != nil {
			staged := any(staged).(*Appearance)
			stage.OnAfterAppearanceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Arpeggiate:
		if stage.OnAfterArpeggiateDeleteCallback != nil {
			staged := any(staged).(*Arpeggiate)
			stage.OnAfterArpeggiateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Arrow:
		if stage.OnAfterArrowDeleteCallback != nil {
			staged := any(staged).(*Arrow)
			stage.OnAfterArrowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Articulations:
		if stage.OnAfterArticulationsDeleteCallback != nil {
			staged := any(staged).(*Articulations)
			stage.OnAfterArticulationsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Assess:
		if stage.OnAfterAssessDeleteCallback != nil {
			staged := any(staged).(*Assess)
			stage.OnAfterAssessDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Attributes:
		if stage.OnAfterAttributesDeleteCallback != nil {
			staged := any(staged).(*Attributes)
			stage.OnAfterAttributesDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Backup:
		if stage.OnAfterBackupDeleteCallback != nil {
			staged := any(staged).(*Backup)
			stage.OnAfterBackupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Bar_style_color:
		if stage.OnAfterBar_style_colorDeleteCallback != nil {
			staged := any(staged).(*Bar_style_color)
			stage.OnAfterBar_style_colorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Barline:
		if stage.OnAfterBarlineDeleteCallback != nil {
			staged := any(staged).(*Barline)
			stage.OnAfterBarlineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Barre:
		if stage.OnAfterBarreDeleteCallback != nil {
			staged := any(staged).(*Barre)
			stage.OnAfterBarreDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Bass:
		if stage.OnAfterBassDeleteCallback != nil {
			staged := any(staged).(*Bass)
			stage.OnAfterBassDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Bass_step:
		if stage.OnAfterBass_stepDeleteCallback != nil {
			staged := any(staged).(*Bass_step)
			stage.OnAfterBass_stepDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Beam:
		if stage.OnAfterBeamDeleteCallback != nil {
			staged := any(staged).(*Beam)
			stage.OnAfterBeamDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Beat_repeat:
		if stage.OnAfterBeat_repeatDeleteCallback != nil {
			staged := any(staged).(*Beat_repeat)
			stage.OnAfterBeat_repeatDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Beat_unit_tied:
		if stage.OnAfterBeat_unit_tiedDeleteCallback != nil {
			staged := any(staged).(*Beat_unit_tied)
			stage.OnAfterBeat_unit_tiedDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Beater:
		if stage.OnAfterBeaterDeleteCallback != nil {
			staged := any(staged).(*Beater)
			stage.OnAfterBeaterDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Bend:
		if stage.OnAfterBendDeleteCallback != nil {
			staged := any(staged).(*Bend)
			stage.OnAfterBendDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Bookmark:
		if stage.OnAfterBookmarkDeleteCallback != nil {
			staged := any(staged).(*Bookmark)
			stage.OnAfterBookmarkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Bracket:
		if stage.OnAfterBracketDeleteCallback != nil {
			staged := any(staged).(*Bracket)
			stage.OnAfterBracketDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Breath_mark:
		if stage.OnAfterBreath_markDeleteCallback != nil {
			staged := any(staged).(*Breath_mark)
			stage.OnAfterBreath_markDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Caesura:
		if stage.OnAfterCaesuraDeleteCallback != nil {
			staged := any(staged).(*Caesura)
			stage.OnAfterCaesuraDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Cancel:
		if stage.OnAfterCancelDeleteCallback != nil {
			staged := any(staged).(*Cancel)
			stage.OnAfterCancelDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Clef:
		if stage.OnAfterClefDeleteCallback != nil {
			staged := any(staged).(*Clef)
			stage.OnAfterClefDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Coda:
		if stage.OnAfterCodaDeleteCallback != nil {
			staged := any(staged).(*Coda)
			stage.OnAfterCodaDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Credit:
		if stage.OnAfterCreditDeleteCallback != nil {
			staged := any(staged).(*Credit)
			stage.OnAfterCreditDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Dashes:
		if stage.OnAfterDashesDeleteCallback != nil {
			staged := any(staged).(*Dashes)
			stage.OnAfterDashesDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Defaults:
		if stage.OnAfterDefaultsDeleteCallback != nil {
			staged := any(staged).(*Defaults)
			stage.OnAfterDefaultsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Degree:
		if stage.OnAfterDegreeDeleteCallback != nil {
			staged := any(staged).(*Degree)
			stage.OnAfterDegreeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Degree_alter:
		if stage.OnAfterDegree_alterDeleteCallback != nil {
			staged := any(staged).(*Degree_alter)
			stage.OnAfterDegree_alterDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Degree_type:
		if stage.OnAfterDegree_typeDeleteCallback != nil {
			staged := any(staged).(*Degree_type)
			stage.OnAfterDegree_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Degree_value:
		if stage.OnAfterDegree_valueDeleteCallback != nil {
			staged := any(staged).(*Degree_value)
			stage.OnAfterDegree_valueDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Direction:
		if stage.OnAfterDirectionDeleteCallback != nil {
			staged := any(staged).(*Direction)
			stage.OnAfterDirectionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Direction_type:
		if stage.OnAfterDirection_typeDeleteCallback != nil {
			staged := any(staged).(*Direction_type)
			stage.OnAfterDirection_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Distance:
		if stage.OnAfterDistanceDeleteCallback != nil {
			staged := any(staged).(*Distance)
			stage.OnAfterDistanceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Double:
		if stage.OnAfterDoubleDeleteCallback != nil {
			staged := any(staged).(*Double)
			stage.OnAfterDoubleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Dynamics:
		if stage.OnAfterDynamicsDeleteCallback != nil {
			staged := any(staged).(*Dynamics)
			stage.OnAfterDynamicsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Effect:
		if stage.OnAfterEffectDeleteCallback != nil {
			staged := any(staged).(*Effect)
			stage.OnAfterEffectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Elision:
		if stage.OnAfterElisionDeleteCallback != nil {
			staged := any(staged).(*Elision)
			stage.OnAfterElisionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty:
		if stage.OnAfterEmptyDeleteCallback != nil {
			staged := any(staged).(*Empty)
			stage.OnAfterEmptyDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty_font:
		if stage.OnAfterEmpty_fontDeleteCallback != nil {
			staged := any(staged).(*Empty_font)
			stage.OnAfterEmpty_fontDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty_line:
		if stage.OnAfterEmpty_lineDeleteCallback != nil {
			staged := any(staged).(*Empty_line)
			stage.OnAfterEmpty_lineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty_placement:
		if stage.OnAfterEmpty_placementDeleteCallback != nil {
			staged := any(staged).(*Empty_placement)
			stage.OnAfterEmpty_placementDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty_placement_smufl:
		if stage.OnAfterEmpty_placement_smuflDeleteCallback != nil {
			staged := any(staged).(*Empty_placement_smufl)
			stage.OnAfterEmpty_placement_smuflDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty_print_object_style_align:
		if stage.OnAfterEmpty_print_object_style_alignDeleteCallback != nil {
			staged := any(staged).(*Empty_print_object_style_align)
			stage.OnAfterEmpty_print_object_style_alignDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty_print_style:
		if stage.OnAfterEmpty_print_styleDeleteCallback != nil {
			staged := any(staged).(*Empty_print_style)
			stage.OnAfterEmpty_print_styleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty_print_style_align:
		if stage.OnAfterEmpty_print_style_alignDeleteCallback != nil {
			staged := any(staged).(*Empty_print_style_align)
			stage.OnAfterEmpty_print_style_alignDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty_print_style_align_id:
		if stage.OnAfterEmpty_print_style_align_idDeleteCallback != nil {
			staged := any(staged).(*Empty_print_style_align_id)
			stage.OnAfterEmpty_print_style_align_idDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Empty_trill_sound:
		if stage.OnAfterEmpty_trill_soundDeleteCallback != nil {
			staged := any(staged).(*Empty_trill_sound)
			stage.OnAfterEmpty_trill_soundDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Encoding:
		if stage.OnAfterEncodingDeleteCallback != nil {
			staged := any(staged).(*Encoding)
			stage.OnAfterEncodingDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Ending:
		if stage.OnAfterEndingDeleteCallback != nil {
			staged := any(staged).(*Ending)
			stage.OnAfterEndingDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Extend:
		if stage.OnAfterExtendDeleteCallback != nil {
			staged := any(staged).(*Extend)
			stage.OnAfterExtendDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Feature:
		if stage.OnAfterFeatureDeleteCallback != nil {
			staged := any(staged).(*Feature)
			stage.OnAfterFeatureDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Fermata:
		if stage.OnAfterFermataDeleteCallback != nil {
			staged := any(staged).(*Fermata)
			stage.OnAfterFermataDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Figure:
		if stage.OnAfterFigureDeleteCallback != nil {
			staged := any(staged).(*Figure)
			stage.OnAfterFigureDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Figured_bass:
		if stage.OnAfterFigured_bassDeleteCallback != nil {
			staged := any(staged).(*Figured_bass)
			stage.OnAfterFigured_bassDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Fingering:
		if stage.OnAfterFingeringDeleteCallback != nil {
			staged := any(staged).(*Fingering)
			stage.OnAfterFingeringDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *First_fret:
		if stage.OnAfterFirst_fretDeleteCallback != nil {
			staged := any(staged).(*First_fret)
			stage.OnAfterFirst_fretDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *For_part:
		if stage.OnAfterFor_partDeleteCallback != nil {
			staged := any(staged).(*For_part)
			stage.OnAfterFor_partDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Formatted_symbol:
		if stage.OnAfterFormatted_symbolDeleteCallback != nil {
			staged := any(staged).(*Formatted_symbol)
			stage.OnAfterFormatted_symbolDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Formatted_symbol_id:
		if stage.OnAfterFormatted_symbol_idDeleteCallback != nil {
			staged := any(staged).(*Formatted_symbol_id)
			stage.OnAfterFormatted_symbol_idDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Formatted_text:
		if stage.OnAfterFormatted_textDeleteCallback != nil {
			staged := any(staged).(*Formatted_text)
			stage.OnAfterFormatted_textDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Formatted_text_id:
		if stage.OnAfterFormatted_text_idDeleteCallback != nil {
			staged := any(staged).(*Formatted_text_id)
			stage.OnAfterFormatted_text_idDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Forward:
		if stage.OnAfterForwardDeleteCallback != nil {
			staged := any(staged).(*Forward)
			stage.OnAfterForwardDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Frame:
		if stage.OnAfterFrameDeleteCallback != nil {
			staged := any(staged).(*Frame)
			stage.OnAfterFrameDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Frame_note:
		if stage.OnAfterFrame_noteDeleteCallback != nil {
			staged := any(staged).(*Frame_note)
			stage.OnAfterFrame_noteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Fret:
		if stage.OnAfterFretDeleteCallback != nil {
			staged := any(staged).(*Fret)
			stage.OnAfterFretDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Glass:
		if stage.OnAfterGlassDeleteCallback != nil {
			staged := any(staged).(*Glass)
			stage.OnAfterGlassDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Glissando:
		if stage.OnAfterGlissandoDeleteCallback != nil {
			staged := any(staged).(*Glissando)
			stage.OnAfterGlissandoDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Glyph:
		if stage.OnAfterGlyphDeleteCallback != nil {
			staged := any(staged).(*Glyph)
			stage.OnAfterGlyphDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Grace:
		if stage.OnAfterGraceDeleteCallback != nil {
			staged := any(staged).(*Grace)
			stage.OnAfterGraceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Group_barline:
		if stage.OnAfterGroup_barlineDeleteCallback != nil {
			staged := any(staged).(*Group_barline)
			stage.OnAfterGroup_barlineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Group_name:
		if stage.OnAfterGroup_nameDeleteCallback != nil {
			staged := any(staged).(*Group_name)
			stage.OnAfterGroup_nameDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Group_symbol:
		if stage.OnAfterGroup_symbolDeleteCallback != nil {
			staged := any(staged).(*Group_symbol)
			stage.OnAfterGroup_symbolDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Grouping:
		if stage.OnAfterGroupingDeleteCallback != nil {
			staged := any(staged).(*Grouping)
			stage.OnAfterGroupingDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Hammer_on_pull_off:
		if stage.OnAfterHammer_on_pull_offDeleteCallback != nil {
			staged := any(staged).(*Hammer_on_pull_off)
			stage.OnAfterHammer_on_pull_offDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Handbell:
		if stage.OnAfterHandbellDeleteCallback != nil {
			staged := any(staged).(*Handbell)
			stage.OnAfterHandbellDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Harmon_closed:
		if stage.OnAfterHarmon_closedDeleteCallback != nil {
			staged := any(staged).(*Harmon_closed)
			stage.OnAfterHarmon_closedDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Harmon_mute:
		if stage.OnAfterHarmon_muteDeleteCallback != nil {
			staged := any(staged).(*Harmon_mute)
			stage.OnAfterHarmon_muteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Harmonic:
		if stage.OnAfterHarmonicDeleteCallback != nil {
			staged := any(staged).(*Harmonic)
			stage.OnAfterHarmonicDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Harmony:
		if stage.OnAfterHarmonyDeleteCallback != nil {
			staged := any(staged).(*Harmony)
			stage.OnAfterHarmonyDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Harmony_alter:
		if stage.OnAfterHarmony_alterDeleteCallback != nil {
			staged := any(staged).(*Harmony_alter)
			stage.OnAfterHarmony_alterDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Harp_pedals:
		if stage.OnAfterHarp_pedalsDeleteCallback != nil {
			staged := any(staged).(*Harp_pedals)
			stage.OnAfterHarp_pedalsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Heel_toe:
		if stage.OnAfterHeel_toeDeleteCallback != nil {
			staged := any(staged).(*Heel_toe)
			stage.OnAfterHeel_toeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Hole:
		if stage.OnAfterHoleDeleteCallback != nil {
			staged := any(staged).(*Hole)
			stage.OnAfterHoleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Hole_closed:
		if stage.OnAfterHole_closedDeleteCallback != nil {
			staged := any(staged).(*Hole_closed)
			stage.OnAfterHole_closedDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Horizontal_turn:
		if stage.OnAfterHorizontal_turnDeleteCallback != nil {
			staged := any(staged).(*Horizontal_turn)
			stage.OnAfterHorizontal_turnDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Identification:
		if stage.OnAfterIdentificationDeleteCallback != nil {
			staged := any(staged).(*Identification)
			stage.OnAfterIdentificationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Image:
		if stage.OnAfterImageDeleteCallback != nil {
			staged := any(staged).(*Image)
			stage.OnAfterImageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Instrument:
		if stage.OnAfterInstrumentDeleteCallback != nil {
			staged := any(staged).(*Instrument)
			stage.OnAfterInstrumentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Instrument_change:
		if stage.OnAfterInstrument_changeDeleteCallback != nil {
			staged := any(staged).(*Instrument_change)
			stage.OnAfterInstrument_changeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Instrument_link:
		if stage.OnAfterInstrument_linkDeleteCallback != nil {
			staged := any(staged).(*Instrument_link)
			stage.OnAfterInstrument_linkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Interchangeable:
		if stage.OnAfterInterchangeableDeleteCallback != nil {
			staged := any(staged).(*Interchangeable)
			stage.OnAfterInterchangeableDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Inversion:
		if stage.OnAfterInversionDeleteCallback != nil {
			staged := any(staged).(*Inversion)
			stage.OnAfterInversionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Key:
		if stage.OnAfterKeyDeleteCallback != nil {
			staged := any(staged).(*Key)
			stage.OnAfterKeyDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Key_accidental:
		if stage.OnAfterKey_accidentalDeleteCallback != nil {
			staged := any(staged).(*Key_accidental)
			stage.OnAfterKey_accidentalDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Key_octave:
		if stage.OnAfterKey_octaveDeleteCallback != nil {
			staged := any(staged).(*Key_octave)
			stage.OnAfterKey_octaveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Kind:
		if stage.OnAfterKindDeleteCallback != nil {
			staged := any(staged).(*Kind)
			stage.OnAfterKindDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Level:
		if stage.OnAfterLevelDeleteCallback != nil {
			staged := any(staged).(*Level)
			stage.OnAfterLevelDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Line_detail:
		if stage.OnAfterLine_detailDeleteCallback != nil {
			staged := any(staged).(*Line_detail)
			stage.OnAfterLine_detailDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Line_width:
		if stage.OnAfterLine_widthDeleteCallback != nil {
			staged := any(staged).(*Line_width)
			stage.OnAfterLine_widthDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Link:
		if stage.OnAfterLinkDeleteCallback != nil {
			staged := any(staged).(*Link)
			stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Listen:
		if stage.OnAfterListenDeleteCallback != nil {
			staged := any(staged).(*Listen)
			stage.OnAfterListenDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Listening:
		if stage.OnAfterListeningDeleteCallback != nil {
			staged := any(staged).(*Listening)
			stage.OnAfterListeningDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Lyric:
		if stage.OnAfterLyricDeleteCallback != nil {
			staged := any(staged).(*Lyric)
			stage.OnAfterLyricDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Lyric_font:
		if stage.OnAfterLyric_fontDeleteCallback != nil {
			staged := any(staged).(*Lyric_font)
			stage.OnAfterLyric_fontDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Lyric_language:
		if stage.OnAfterLyric_languageDeleteCallback != nil {
			staged := any(staged).(*Lyric_language)
			stage.OnAfterLyric_languageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Measure_layout:
		if stage.OnAfterMeasure_layoutDeleteCallback != nil {
			staged := any(staged).(*Measure_layout)
			stage.OnAfterMeasure_layoutDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Measure_numbering:
		if stage.OnAfterMeasure_numberingDeleteCallback != nil {
			staged := any(staged).(*Measure_numbering)
			stage.OnAfterMeasure_numberingDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Measure_repeat:
		if stage.OnAfterMeasure_repeatDeleteCallback != nil {
			staged := any(staged).(*Measure_repeat)
			stage.OnAfterMeasure_repeatDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Measure_style:
		if stage.OnAfterMeasure_styleDeleteCallback != nil {
			staged := any(staged).(*Measure_style)
			stage.OnAfterMeasure_styleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Membrane:
		if stage.OnAfterMembraneDeleteCallback != nil {
			staged := any(staged).(*Membrane)
			stage.OnAfterMembraneDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Metal:
		if stage.OnAfterMetalDeleteCallback != nil {
			staged := any(staged).(*Metal)
			stage.OnAfterMetalDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Metronome:
		if stage.OnAfterMetronomeDeleteCallback != nil {
			staged := any(staged).(*Metronome)
			stage.OnAfterMetronomeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Metronome_beam:
		if stage.OnAfterMetronome_beamDeleteCallback != nil {
			staged := any(staged).(*Metronome_beam)
			stage.OnAfterMetronome_beamDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Metronome_note:
		if stage.OnAfterMetronome_noteDeleteCallback != nil {
			staged := any(staged).(*Metronome_note)
			stage.OnAfterMetronome_noteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Metronome_tied:
		if stage.OnAfterMetronome_tiedDeleteCallback != nil {
			staged := any(staged).(*Metronome_tied)
			stage.OnAfterMetronome_tiedDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Metronome_tuplet:
		if stage.OnAfterMetronome_tupletDeleteCallback != nil {
			staged := any(staged).(*Metronome_tuplet)
			stage.OnAfterMetronome_tupletDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Midi_device:
		if stage.OnAfterMidi_deviceDeleteCallback != nil {
			staged := any(staged).(*Midi_device)
			stage.OnAfterMidi_deviceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Midi_instrument:
		if stage.OnAfterMidi_instrumentDeleteCallback != nil {
			staged := any(staged).(*Midi_instrument)
			stage.OnAfterMidi_instrumentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Miscellaneous:
		if stage.OnAfterMiscellaneousDeleteCallback != nil {
			staged := any(staged).(*Miscellaneous)
			stage.OnAfterMiscellaneousDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Miscellaneous_field:
		if stage.OnAfterMiscellaneous_fieldDeleteCallback != nil {
			staged := any(staged).(*Miscellaneous_field)
			stage.OnAfterMiscellaneous_fieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Mordent:
		if stage.OnAfterMordentDeleteCallback != nil {
			staged := any(staged).(*Mordent)
			stage.OnAfterMordentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Multiple_rest:
		if stage.OnAfterMultiple_restDeleteCallback != nil {
			staged := any(staged).(*Multiple_rest)
			stage.OnAfterMultiple_restDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Name_display:
		if stage.OnAfterName_displayDeleteCallback != nil {
			staged := any(staged).(*Name_display)
			stage.OnAfterName_displayDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Non_arpeggiate:
		if stage.OnAfterNon_arpeggiateDeleteCallback != nil {
			staged := any(staged).(*Non_arpeggiate)
			stage.OnAfterNon_arpeggiateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Notations:
		if stage.OnAfterNotationsDeleteCallback != nil {
			staged := any(staged).(*Notations)
			stage.OnAfterNotationsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Note:
		if stage.OnAfterNoteDeleteCallback != nil {
			staged := any(staged).(*Note)
			stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Note_size:
		if stage.OnAfterNote_sizeDeleteCallback != nil {
			staged := any(staged).(*Note_size)
			stage.OnAfterNote_sizeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Note_type:
		if stage.OnAfterNote_typeDeleteCallback != nil {
			staged := any(staged).(*Note_type)
			stage.OnAfterNote_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Notehead:
		if stage.OnAfterNoteheadDeleteCallback != nil {
			staged := any(staged).(*Notehead)
			stage.OnAfterNoteheadDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Notehead_text:
		if stage.OnAfterNotehead_textDeleteCallback != nil {
			staged := any(staged).(*Notehead_text)
			stage.OnAfterNotehead_textDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Numeral:
		if stage.OnAfterNumeralDeleteCallback != nil {
			staged := any(staged).(*Numeral)
			stage.OnAfterNumeralDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Numeral_key:
		if stage.OnAfterNumeral_keyDeleteCallback != nil {
			staged := any(staged).(*Numeral_key)
			stage.OnAfterNumeral_keyDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Numeral_root:
		if stage.OnAfterNumeral_rootDeleteCallback != nil {
			staged := any(staged).(*Numeral_root)
			stage.OnAfterNumeral_rootDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Octave_shift:
		if stage.OnAfterOctave_shiftDeleteCallback != nil {
			staged := any(staged).(*Octave_shift)
			stage.OnAfterOctave_shiftDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Offset:
		if stage.OnAfterOffsetDeleteCallback != nil {
			staged := any(staged).(*Offset)
			stage.OnAfterOffsetDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Opus:
		if stage.OnAfterOpusDeleteCallback != nil {
			staged := any(staged).(*Opus)
			stage.OnAfterOpusDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Ornaments:
		if stage.OnAfterOrnamentsDeleteCallback != nil {
			staged := any(staged).(*Ornaments)
			stage.OnAfterOrnamentsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Other_appearance:
		if stage.OnAfterOther_appearanceDeleteCallback != nil {
			staged := any(staged).(*Other_appearance)
			stage.OnAfterOther_appearanceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Other_direction:
		if stage.OnAfterOther_directionDeleteCallback != nil {
			staged := any(staged).(*Other_direction)
			stage.OnAfterOther_directionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Other_listening:
		if stage.OnAfterOther_listeningDeleteCallback != nil {
			staged := any(staged).(*Other_listening)
			stage.OnAfterOther_listeningDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Other_notation:
		if stage.OnAfterOther_notationDeleteCallback != nil {
			staged := any(staged).(*Other_notation)
			stage.OnAfterOther_notationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Other_placement_text:
		if stage.OnAfterOther_placement_textDeleteCallback != nil {
			staged := any(staged).(*Other_placement_text)
			stage.OnAfterOther_placement_textDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Other_play:
		if stage.OnAfterOther_playDeleteCallback != nil {
			staged := any(staged).(*Other_play)
			stage.OnAfterOther_playDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Other_text:
		if stage.OnAfterOther_textDeleteCallback != nil {
			staged := any(staged).(*Other_text)
			stage.OnAfterOther_textDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Page_layout:
		if stage.OnAfterPage_layoutDeleteCallback != nil {
			staged := any(staged).(*Page_layout)
			stage.OnAfterPage_layoutDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Page_margins:
		if stage.OnAfterPage_marginsDeleteCallback != nil {
			staged := any(staged).(*Page_margins)
			stage.OnAfterPage_marginsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Part_clef:
		if stage.OnAfterPart_clefDeleteCallback != nil {
			staged := any(staged).(*Part_clef)
			stage.OnAfterPart_clefDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Part_group:
		if stage.OnAfterPart_groupDeleteCallback != nil {
			staged := any(staged).(*Part_group)
			stage.OnAfterPart_groupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Part_link:
		if stage.OnAfterPart_linkDeleteCallback != nil {
			staged := any(staged).(*Part_link)
			stage.OnAfterPart_linkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Part_list:
		if stage.OnAfterPart_listDeleteCallback != nil {
			staged := any(staged).(*Part_list)
			stage.OnAfterPart_listDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Part_name:
		if stage.OnAfterPart_nameDeleteCallback != nil {
			staged := any(staged).(*Part_name)
			stage.OnAfterPart_nameDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Part_symbol:
		if stage.OnAfterPart_symbolDeleteCallback != nil {
			staged := any(staged).(*Part_symbol)
			stage.OnAfterPart_symbolDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Part_transpose:
		if stage.OnAfterPart_transposeDeleteCallback != nil {
			staged := any(staged).(*Part_transpose)
			stage.OnAfterPart_transposeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Pedal:
		if stage.OnAfterPedalDeleteCallback != nil {
			staged := any(staged).(*Pedal)
			stage.OnAfterPedalDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Pedal_tuning:
		if stage.OnAfterPedal_tuningDeleteCallback != nil {
			staged := any(staged).(*Pedal_tuning)
			stage.OnAfterPedal_tuningDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Per_minute:
		if stage.OnAfterPer_minuteDeleteCallback != nil {
			staged := any(staged).(*Per_minute)
			stage.OnAfterPer_minuteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Percussion:
		if stage.OnAfterPercussionDeleteCallback != nil {
			staged := any(staged).(*Percussion)
			stage.OnAfterPercussionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Pitch:
		if stage.OnAfterPitchDeleteCallback != nil {
			staged := any(staged).(*Pitch)
			stage.OnAfterPitchDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Pitched:
		if stage.OnAfterPitchedDeleteCallback != nil {
			staged := any(staged).(*Pitched)
			stage.OnAfterPitchedDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Placement_text:
		if stage.OnAfterPlacement_textDeleteCallback != nil {
			staged := any(staged).(*Placement_text)
			stage.OnAfterPlacement_textDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Play:
		if stage.OnAfterPlayDeleteCallback != nil {
			staged := any(staged).(*Play)
			stage.OnAfterPlayDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Player:
		if stage.OnAfterPlayerDeleteCallback != nil {
			staged := any(staged).(*Player)
			stage.OnAfterPlayerDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Principal_voice:
		if stage.OnAfterPrincipal_voiceDeleteCallback != nil {
			staged := any(staged).(*Principal_voice)
			stage.OnAfterPrincipal_voiceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Print:
		if stage.OnAfterPrintDeleteCallback != nil {
			staged := any(staged).(*Print)
			stage.OnAfterPrintDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Release:
		if stage.OnAfterReleaseDeleteCallback != nil {
			staged := any(staged).(*Release)
			stage.OnAfterReleaseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Repeat:
		if stage.OnAfterRepeatDeleteCallback != nil {
			staged := any(staged).(*Repeat)
			stage.OnAfterRepeatDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Rest:
		if stage.OnAfterRestDeleteCallback != nil {
			staged := any(staged).(*Rest)
			stage.OnAfterRestDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Root:
		if stage.OnAfterRootDeleteCallback != nil {
			staged := any(staged).(*Root)
			stage.OnAfterRootDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Root_step:
		if stage.OnAfterRoot_stepDeleteCallback != nil {
			staged := any(staged).(*Root_step)
			stage.OnAfterRoot_stepDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Scaling:
		if stage.OnAfterScalingDeleteCallback != nil {
			staged := any(staged).(*Scaling)
			stage.OnAfterScalingDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Scordatura:
		if stage.OnAfterScordaturaDeleteCallback != nil {
			staged := any(staged).(*Scordatura)
			stage.OnAfterScordaturaDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Score_instrument:
		if stage.OnAfterScore_instrumentDeleteCallback != nil {
			staged := any(staged).(*Score_instrument)
			stage.OnAfterScore_instrumentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Score_part:
		if stage.OnAfterScore_partDeleteCallback != nil {
			staged := any(staged).(*Score_part)
			stage.OnAfterScore_partDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Score_partwise:
		if stage.OnAfterScore_partwiseDeleteCallback != nil {
			staged := any(staged).(*Score_partwise)
			stage.OnAfterScore_partwiseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Score_timewise:
		if stage.OnAfterScore_timewiseDeleteCallback != nil {
			staged := any(staged).(*Score_timewise)
			stage.OnAfterScore_timewiseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Segno:
		if stage.OnAfterSegnoDeleteCallback != nil {
			staged := any(staged).(*Segno)
			stage.OnAfterSegnoDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Slash:
		if stage.OnAfterSlashDeleteCallback != nil {
			staged := any(staged).(*Slash)
			stage.OnAfterSlashDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Slide:
		if stage.OnAfterSlideDeleteCallback != nil {
			staged := any(staged).(*Slide)
			stage.OnAfterSlideDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Slur:
		if stage.OnAfterSlurDeleteCallback != nil {
			staged := any(staged).(*Slur)
			stage.OnAfterSlurDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Sound:
		if stage.OnAfterSoundDeleteCallback != nil {
			staged := any(staged).(*Sound)
			stage.OnAfterSoundDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Staff_details:
		if stage.OnAfterStaff_detailsDeleteCallback != nil {
			staged := any(staged).(*Staff_details)
			stage.OnAfterStaff_detailsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Staff_divide:
		if stage.OnAfterStaff_divideDeleteCallback != nil {
			staged := any(staged).(*Staff_divide)
			stage.OnAfterStaff_divideDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Staff_layout:
		if stage.OnAfterStaff_layoutDeleteCallback != nil {
			staged := any(staged).(*Staff_layout)
			stage.OnAfterStaff_layoutDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Staff_size:
		if stage.OnAfterStaff_sizeDeleteCallback != nil {
			staged := any(staged).(*Staff_size)
			stage.OnAfterStaff_sizeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Staff_tuning:
		if stage.OnAfterStaff_tuningDeleteCallback != nil {
			staged := any(staged).(*Staff_tuning)
			stage.OnAfterStaff_tuningDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Stem:
		if stage.OnAfterStemDeleteCallback != nil {
			staged := any(staged).(*Stem)
			stage.OnAfterStemDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Stick:
		if stage.OnAfterStickDeleteCallback != nil {
			staged := any(staged).(*Stick)
			stage.OnAfterStickDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *String_mute:
		if stage.OnAfterString_muteDeleteCallback != nil {
			staged := any(staged).(*String_mute)
			stage.OnAfterString_muteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *String_type:
		if stage.OnAfterString_typeDeleteCallback != nil {
			staged := any(staged).(*String_type)
			stage.OnAfterString_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Strong_accent:
		if stage.OnAfterStrong_accentDeleteCallback != nil {
			staged := any(staged).(*Strong_accent)
			stage.OnAfterStrong_accentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Style_text:
		if stage.OnAfterStyle_textDeleteCallback != nil {
			staged := any(staged).(*Style_text)
			stage.OnAfterStyle_textDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Supports:
		if stage.OnAfterSupportsDeleteCallback != nil {
			staged := any(staged).(*Supports)
			stage.OnAfterSupportsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Swing:
		if stage.OnAfterSwingDeleteCallback != nil {
			staged := any(staged).(*Swing)
			stage.OnAfterSwingDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Sync:
		if stage.OnAfterSyncDeleteCallback != nil {
			staged := any(staged).(*Sync)
			stage.OnAfterSyncDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *System_dividers:
		if stage.OnAfterSystem_dividersDeleteCallback != nil {
			staged := any(staged).(*System_dividers)
			stage.OnAfterSystem_dividersDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *System_layout:
		if stage.OnAfterSystem_layoutDeleteCallback != nil {
			staged := any(staged).(*System_layout)
			stage.OnAfterSystem_layoutDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *System_margins:
		if stage.OnAfterSystem_marginsDeleteCallback != nil {
			staged := any(staged).(*System_margins)
			stage.OnAfterSystem_marginsDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tap:
		if stage.OnAfterTapDeleteCallback != nil {
			staged := any(staged).(*Tap)
			stage.OnAfterTapDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Technical:
		if stage.OnAfterTechnicalDeleteCallback != nil {
			staged := any(staged).(*Technical)
			stage.OnAfterTechnicalDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Text_element_data:
		if stage.OnAfterText_element_dataDeleteCallback != nil {
			staged := any(staged).(*Text_element_data)
			stage.OnAfterText_element_dataDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tie:
		if stage.OnAfterTieDeleteCallback != nil {
			staged := any(staged).(*Tie)
			stage.OnAfterTieDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tied:
		if stage.OnAfterTiedDeleteCallback != nil {
			staged := any(staged).(*Tied)
			stage.OnAfterTiedDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Time:
		if stage.OnAfterTimeDeleteCallback != nil {
			staged := any(staged).(*Time)
			stage.OnAfterTimeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Time_modification:
		if stage.OnAfterTime_modificationDeleteCallback != nil {
			staged := any(staged).(*Time_modification)
			stage.OnAfterTime_modificationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Timpani:
		if stage.OnAfterTimpaniDeleteCallback != nil {
			staged := any(staged).(*Timpani)
			stage.OnAfterTimpaniDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Transpose:
		if stage.OnAfterTransposeDeleteCallback != nil {
			staged := any(staged).(*Transpose)
			stage.OnAfterTransposeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tremolo:
		if stage.OnAfterTremoloDeleteCallback != nil {
			staged := any(staged).(*Tremolo)
			stage.OnAfterTremoloDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tuplet:
		if stage.OnAfterTupletDeleteCallback != nil {
			staged := any(staged).(*Tuplet)
			stage.OnAfterTupletDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tuplet_dot:
		if stage.OnAfterTuplet_dotDeleteCallback != nil {
			staged := any(staged).(*Tuplet_dot)
			stage.OnAfterTuplet_dotDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tuplet_number:
		if stage.OnAfterTuplet_numberDeleteCallback != nil {
			staged := any(staged).(*Tuplet_number)
			stage.OnAfterTuplet_numberDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tuplet_portion:
		if stage.OnAfterTuplet_portionDeleteCallback != nil {
			staged := any(staged).(*Tuplet_portion)
			stage.OnAfterTuplet_portionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tuplet_type:
		if stage.OnAfterTuplet_typeDeleteCallback != nil {
			staged := any(staged).(*Tuplet_type)
			stage.OnAfterTuplet_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Typed_text:
		if stage.OnAfterTyped_textDeleteCallback != nil {
			staged := any(staged).(*Typed_text)
			stage.OnAfterTyped_textDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Unpitched:
		if stage.OnAfterUnpitchedDeleteCallback != nil {
			staged := any(staged).(*Unpitched)
			stage.OnAfterUnpitchedDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Virtual_instrument:
		if stage.OnAfterVirtual_instrumentDeleteCallback != nil {
			staged := any(staged).(*Virtual_instrument)
			stage.OnAfterVirtual_instrumentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Wait:
		if stage.OnAfterWaitDeleteCallback != nil {
			staged := any(staged).(*Wait)
			stage.OnAfterWaitDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Wavy_line:
		if stage.OnAfterWavy_lineDeleteCallback != nil {
			staged := any(staged).(*Wavy_line)
			stage.OnAfterWavy_lineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Wedge:
		if stage.OnAfterWedgeDeleteCallback != nil {
			staged := any(staged).(*Wedge)
			stage.OnAfterWedgeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Wood:
		if stage.OnAfterWoodDeleteCallback != nil {
			staged := any(staged).(*Wood)
			stage.OnAfterWoodDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Work:
		if stage.OnAfterWorkDeleteCallback != nil {
			staged := any(staged).(*Work)
			stage.OnAfterWorkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
