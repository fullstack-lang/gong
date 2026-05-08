// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

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

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

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

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

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

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *A_directive:
		if stage.OnAfterA_directiveReadCallback != nil {
			stage.OnAfterA_directiveReadCallback.OnAfterRead(stage, target)
		}
	case *A_measure:
		if stage.OnAfterA_measureReadCallback != nil {
			stage.OnAfterA_measureReadCallback.OnAfterRead(stage, target)
		}
	case *A_measure_1:
		if stage.OnAfterA_measure_1ReadCallback != nil {
			stage.OnAfterA_measure_1ReadCallback.OnAfterRead(stage, target)
		}
	case *A_part:
		if stage.OnAfterA_partReadCallback != nil {
			stage.OnAfterA_partReadCallback.OnAfterRead(stage, target)
		}
	case *A_part_1:
		if stage.OnAfterA_part_1ReadCallback != nil {
			stage.OnAfterA_part_1ReadCallback.OnAfterRead(stage, target)
		}
	case *Accidental:
		if stage.OnAfterAccidentalReadCallback != nil {
			stage.OnAfterAccidentalReadCallback.OnAfterRead(stage, target)
		}
	case *Accidental_mark:
		if stage.OnAfterAccidental_markReadCallback != nil {
			stage.OnAfterAccidental_markReadCallback.OnAfterRead(stage, target)
		}
	case *Accidental_text:
		if stage.OnAfterAccidental_textReadCallback != nil {
			stage.OnAfterAccidental_textReadCallback.OnAfterRead(stage, target)
		}
	case *Accord:
		if stage.OnAfterAccordReadCallback != nil {
			stage.OnAfterAccordReadCallback.OnAfterRead(stage, target)
		}
	case *Accordion_registration:
		if stage.OnAfterAccordion_registrationReadCallback != nil {
			stage.OnAfterAccordion_registrationReadCallback.OnAfterRead(stage, target)
		}
	case *Appearance:
		if stage.OnAfterAppearanceReadCallback != nil {
			stage.OnAfterAppearanceReadCallback.OnAfterRead(stage, target)
		}
	case *Arpeggiate:
		if stage.OnAfterArpeggiateReadCallback != nil {
			stage.OnAfterArpeggiateReadCallback.OnAfterRead(stage, target)
		}
	case *Arrow:
		if stage.OnAfterArrowReadCallback != nil {
			stage.OnAfterArrowReadCallback.OnAfterRead(stage, target)
		}
	case *Articulations:
		if stage.OnAfterArticulationsReadCallback != nil {
			stage.OnAfterArticulationsReadCallback.OnAfterRead(stage, target)
		}
	case *Assess:
		if stage.OnAfterAssessReadCallback != nil {
			stage.OnAfterAssessReadCallback.OnAfterRead(stage, target)
		}
	case *Attributes:
		if stage.OnAfterAttributesReadCallback != nil {
			stage.OnAfterAttributesReadCallback.OnAfterRead(stage, target)
		}
	case *Backup:
		if stage.OnAfterBackupReadCallback != nil {
			stage.OnAfterBackupReadCallback.OnAfterRead(stage, target)
		}
	case *Bar_style_color:
		if stage.OnAfterBar_style_colorReadCallback != nil {
			stage.OnAfterBar_style_colorReadCallback.OnAfterRead(stage, target)
		}
	case *Barline:
		if stage.OnAfterBarlineReadCallback != nil {
			stage.OnAfterBarlineReadCallback.OnAfterRead(stage, target)
		}
	case *Barre:
		if stage.OnAfterBarreReadCallback != nil {
			stage.OnAfterBarreReadCallback.OnAfterRead(stage, target)
		}
	case *Bass:
		if stage.OnAfterBassReadCallback != nil {
			stage.OnAfterBassReadCallback.OnAfterRead(stage, target)
		}
	case *Bass_step:
		if stage.OnAfterBass_stepReadCallback != nil {
			stage.OnAfterBass_stepReadCallback.OnAfterRead(stage, target)
		}
	case *Beam:
		if stage.OnAfterBeamReadCallback != nil {
			stage.OnAfterBeamReadCallback.OnAfterRead(stage, target)
		}
	case *Beat_repeat:
		if stage.OnAfterBeat_repeatReadCallback != nil {
			stage.OnAfterBeat_repeatReadCallback.OnAfterRead(stage, target)
		}
	case *Beat_unit_tied:
		if stage.OnAfterBeat_unit_tiedReadCallback != nil {
			stage.OnAfterBeat_unit_tiedReadCallback.OnAfterRead(stage, target)
		}
	case *Beater:
		if stage.OnAfterBeaterReadCallback != nil {
			stage.OnAfterBeaterReadCallback.OnAfterRead(stage, target)
		}
	case *Bend:
		if stage.OnAfterBendReadCallback != nil {
			stage.OnAfterBendReadCallback.OnAfterRead(stage, target)
		}
	case *Bookmark:
		if stage.OnAfterBookmarkReadCallback != nil {
			stage.OnAfterBookmarkReadCallback.OnAfterRead(stage, target)
		}
	case *Bracket:
		if stage.OnAfterBracketReadCallback != nil {
			stage.OnAfterBracketReadCallback.OnAfterRead(stage, target)
		}
	case *Breath_mark:
		if stage.OnAfterBreath_markReadCallback != nil {
			stage.OnAfterBreath_markReadCallback.OnAfterRead(stage, target)
		}
	case *Caesura:
		if stage.OnAfterCaesuraReadCallback != nil {
			stage.OnAfterCaesuraReadCallback.OnAfterRead(stage, target)
		}
	case *Cancel:
		if stage.OnAfterCancelReadCallback != nil {
			stage.OnAfterCancelReadCallback.OnAfterRead(stage, target)
		}
	case *Clef:
		if stage.OnAfterClefReadCallback != nil {
			stage.OnAfterClefReadCallback.OnAfterRead(stage, target)
		}
	case *Coda:
		if stage.OnAfterCodaReadCallback != nil {
			stage.OnAfterCodaReadCallback.OnAfterRead(stage, target)
		}
	case *Credit:
		if stage.OnAfterCreditReadCallback != nil {
			stage.OnAfterCreditReadCallback.OnAfterRead(stage, target)
		}
	case *Dashes:
		if stage.OnAfterDashesReadCallback != nil {
			stage.OnAfterDashesReadCallback.OnAfterRead(stage, target)
		}
	case *Defaults:
		if stage.OnAfterDefaultsReadCallback != nil {
			stage.OnAfterDefaultsReadCallback.OnAfterRead(stage, target)
		}
	case *Degree:
		if stage.OnAfterDegreeReadCallback != nil {
			stage.OnAfterDegreeReadCallback.OnAfterRead(stage, target)
		}
	case *Degree_alter:
		if stage.OnAfterDegree_alterReadCallback != nil {
			stage.OnAfterDegree_alterReadCallback.OnAfterRead(stage, target)
		}
	case *Degree_type:
		if stage.OnAfterDegree_typeReadCallback != nil {
			stage.OnAfterDegree_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Degree_value:
		if stage.OnAfterDegree_valueReadCallback != nil {
			stage.OnAfterDegree_valueReadCallback.OnAfterRead(stage, target)
		}
	case *Direction:
		if stage.OnAfterDirectionReadCallback != nil {
			stage.OnAfterDirectionReadCallback.OnAfterRead(stage, target)
		}
	case *Direction_type:
		if stage.OnAfterDirection_typeReadCallback != nil {
			stage.OnAfterDirection_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Distance:
		if stage.OnAfterDistanceReadCallback != nil {
			stage.OnAfterDistanceReadCallback.OnAfterRead(stage, target)
		}
	case *Double:
		if stage.OnAfterDoubleReadCallback != nil {
			stage.OnAfterDoubleReadCallback.OnAfterRead(stage, target)
		}
	case *Dynamics:
		if stage.OnAfterDynamicsReadCallback != nil {
			stage.OnAfterDynamicsReadCallback.OnAfterRead(stage, target)
		}
	case *Effect:
		if stage.OnAfterEffectReadCallback != nil {
			stage.OnAfterEffectReadCallback.OnAfterRead(stage, target)
		}
	case *Elision:
		if stage.OnAfterElisionReadCallback != nil {
			stage.OnAfterElisionReadCallback.OnAfterRead(stage, target)
		}
	case *Empty:
		if stage.OnAfterEmptyReadCallback != nil {
			stage.OnAfterEmptyReadCallback.OnAfterRead(stage, target)
		}
	case *Empty_font:
		if stage.OnAfterEmpty_fontReadCallback != nil {
			stage.OnAfterEmpty_fontReadCallback.OnAfterRead(stage, target)
		}
	case *Empty_line:
		if stage.OnAfterEmpty_lineReadCallback != nil {
			stage.OnAfterEmpty_lineReadCallback.OnAfterRead(stage, target)
		}
	case *Empty_placement:
		if stage.OnAfterEmpty_placementReadCallback != nil {
			stage.OnAfterEmpty_placementReadCallback.OnAfterRead(stage, target)
		}
	case *Empty_placement_smufl:
		if stage.OnAfterEmpty_placement_smuflReadCallback != nil {
			stage.OnAfterEmpty_placement_smuflReadCallback.OnAfterRead(stage, target)
		}
	case *Empty_print_object_style_align:
		if stage.OnAfterEmpty_print_object_style_alignReadCallback != nil {
			stage.OnAfterEmpty_print_object_style_alignReadCallback.OnAfterRead(stage, target)
		}
	case *Empty_print_style:
		if stage.OnAfterEmpty_print_styleReadCallback != nil {
			stage.OnAfterEmpty_print_styleReadCallback.OnAfterRead(stage, target)
		}
	case *Empty_print_style_align:
		if stage.OnAfterEmpty_print_style_alignReadCallback != nil {
			stage.OnAfterEmpty_print_style_alignReadCallback.OnAfterRead(stage, target)
		}
	case *Empty_print_style_align_id:
		if stage.OnAfterEmpty_print_style_align_idReadCallback != nil {
			stage.OnAfterEmpty_print_style_align_idReadCallback.OnAfterRead(stage, target)
		}
	case *Empty_trill_sound:
		if stage.OnAfterEmpty_trill_soundReadCallback != nil {
			stage.OnAfterEmpty_trill_soundReadCallback.OnAfterRead(stage, target)
		}
	case *Encoding:
		if stage.OnAfterEncodingReadCallback != nil {
			stage.OnAfterEncodingReadCallback.OnAfterRead(stage, target)
		}
	case *Ending:
		if stage.OnAfterEndingReadCallback != nil {
			stage.OnAfterEndingReadCallback.OnAfterRead(stage, target)
		}
	case *Extend:
		if stage.OnAfterExtendReadCallback != nil {
			stage.OnAfterExtendReadCallback.OnAfterRead(stage, target)
		}
	case *Feature:
		if stage.OnAfterFeatureReadCallback != nil {
			stage.OnAfterFeatureReadCallback.OnAfterRead(stage, target)
		}
	case *Fermata:
		if stage.OnAfterFermataReadCallback != nil {
			stage.OnAfterFermataReadCallback.OnAfterRead(stage, target)
		}
	case *Figure:
		if stage.OnAfterFigureReadCallback != nil {
			stage.OnAfterFigureReadCallback.OnAfterRead(stage, target)
		}
	case *Figured_bass:
		if stage.OnAfterFigured_bassReadCallback != nil {
			stage.OnAfterFigured_bassReadCallback.OnAfterRead(stage, target)
		}
	case *Fingering:
		if stage.OnAfterFingeringReadCallback != nil {
			stage.OnAfterFingeringReadCallback.OnAfterRead(stage, target)
		}
	case *First_fret:
		if stage.OnAfterFirst_fretReadCallback != nil {
			stage.OnAfterFirst_fretReadCallback.OnAfterRead(stage, target)
		}
	case *For_part:
		if stage.OnAfterFor_partReadCallback != nil {
			stage.OnAfterFor_partReadCallback.OnAfterRead(stage, target)
		}
	case *Formatted_symbol:
		if stage.OnAfterFormatted_symbolReadCallback != nil {
			stage.OnAfterFormatted_symbolReadCallback.OnAfterRead(stage, target)
		}
	case *Formatted_symbol_id:
		if stage.OnAfterFormatted_symbol_idReadCallback != nil {
			stage.OnAfterFormatted_symbol_idReadCallback.OnAfterRead(stage, target)
		}
	case *Formatted_text:
		if stage.OnAfterFormatted_textReadCallback != nil {
			stage.OnAfterFormatted_textReadCallback.OnAfterRead(stage, target)
		}
	case *Formatted_text_id:
		if stage.OnAfterFormatted_text_idReadCallback != nil {
			stage.OnAfterFormatted_text_idReadCallback.OnAfterRead(stage, target)
		}
	case *Forward:
		if stage.OnAfterForwardReadCallback != nil {
			stage.OnAfterForwardReadCallback.OnAfterRead(stage, target)
		}
	case *Frame:
		if stage.OnAfterFrameReadCallback != nil {
			stage.OnAfterFrameReadCallback.OnAfterRead(stage, target)
		}
	case *Frame_note:
		if stage.OnAfterFrame_noteReadCallback != nil {
			stage.OnAfterFrame_noteReadCallback.OnAfterRead(stage, target)
		}
	case *Fret:
		if stage.OnAfterFretReadCallback != nil {
			stage.OnAfterFretReadCallback.OnAfterRead(stage, target)
		}
	case *Glass:
		if stage.OnAfterGlassReadCallback != nil {
			stage.OnAfterGlassReadCallback.OnAfterRead(stage, target)
		}
	case *Glissando:
		if stage.OnAfterGlissandoReadCallback != nil {
			stage.OnAfterGlissandoReadCallback.OnAfterRead(stage, target)
		}
	case *Glyph:
		if stage.OnAfterGlyphReadCallback != nil {
			stage.OnAfterGlyphReadCallback.OnAfterRead(stage, target)
		}
	case *Grace:
		if stage.OnAfterGraceReadCallback != nil {
			stage.OnAfterGraceReadCallback.OnAfterRead(stage, target)
		}
	case *Group_barline:
		if stage.OnAfterGroup_barlineReadCallback != nil {
			stage.OnAfterGroup_barlineReadCallback.OnAfterRead(stage, target)
		}
	case *Group_name:
		if stage.OnAfterGroup_nameReadCallback != nil {
			stage.OnAfterGroup_nameReadCallback.OnAfterRead(stage, target)
		}
	case *Group_symbol:
		if stage.OnAfterGroup_symbolReadCallback != nil {
			stage.OnAfterGroup_symbolReadCallback.OnAfterRead(stage, target)
		}
	case *Grouping:
		if stage.OnAfterGroupingReadCallback != nil {
			stage.OnAfterGroupingReadCallback.OnAfterRead(stage, target)
		}
	case *Hammer_on_pull_off:
		if stage.OnAfterHammer_on_pull_offReadCallback != nil {
			stage.OnAfterHammer_on_pull_offReadCallback.OnAfterRead(stage, target)
		}
	case *Handbell:
		if stage.OnAfterHandbellReadCallback != nil {
			stage.OnAfterHandbellReadCallback.OnAfterRead(stage, target)
		}
	case *Harmon_closed:
		if stage.OnAfterHarmon_closedReadCallback != nil {
			stage.OnAfterHarmon_closedReadCallback.OnAfterRead(stage, target)
		}
	case *Harmon_mute:
		if stage.OnAfterHarmon_muteReadCallback != nil {
			stage.OnAfterHarmon_muteReadCallback.OnAfterRead(stage, target)
		}
	case *Harmonic:
		if stage.OnAfterHarmonicReadCallback != nil {
			stage.OnAfterHarmonicReadCallback.OnAfterRead(stage, target)
		}
	case *Harmony:
		if stage.OnAfterHarmonyReadCallback != nil {
			stage.OnAfterHarmonyReadCallback.OnAfterRead(stage, target)
		}
	case *Harmony_alter:
		if stage.OnAfterHarmony_alterReadCallback != nil {
			stage.OnAfterHarmony_alterReadCallback.OnAfterRead(stage, target)
		}
	case *Harp_pedals:
		if stage.OnAfterHarp_pedalsReadCallback != nil {
			stage.OnAfterHarp_pedalsReadCallback.OnAfterRead(stage, target)
		}
	case *Heel_toe:
		if stage.OnAfterHeel_toeReadCallback != nil {
			stage.OnAfterHeel_toeReadCallback.OnAfterRead(stage, target)
		}
	case *Hole:
		if stage.OnAfterHoleReadCallback != nil {
			stage.OnAfterHoleReadCallback.OnAfterRead(stage, target)
		}
	case *Hole_closed:
		if stage.OnAfterHole_closedReadCallback != nil {
			stage.OnAfterHole_closedReadCallback.OnAfterRead(stage, target)
		}
	case *Horizontal_turn:
		if stage.OnAfterHorizontal_turnReadCallback != nil {
			stage.OnAfterHorizontal_turnReadCallback.OnAfterRead(stage, target)
		}
	case *Identification:
		if stage.OnAfterIdentificationReadCallback != nil {
			stage.OnAfterIdentificationReadCallback.OnAfterRead(stage, target)
		}
	case *Image:
		if stage.OnAfterImageReadCallback != nil {
			stage.OnAfterImageReadCallback.OnAfterRead(stage, target)
		}
	case *Instrument:
		if stage.OnAfterInstrumentReadCallback != nil {
			stage.OnAfterInstrumentReadCallback.OnAfterRead(stage, target)
		}
	case *Instrument_change:
		if stage.OnAfterInstrument_changeReadCallback != nil {
			stage.OnAfterInstrument_changeReadCallback.OnAfterRead(stage, target)
		}
	case *Instrument_link:
		if stage.OnAfterInstrument_linkReadCallback != nil {
			stage.OnAfterInstrument_linkReadCallback.OnAfterRead(stage, target)
		}
	case *Interchangeable:
		if stage.OnAfterInterchangeableReadCallback != nil {
			stage.OnAfterInterchangeableReadCallback.OnAfterRead(stage, target)
		}
	case *Inversion:
		if stage.OnAfterInversionReadCallback != nil {
			stage.OnAfterInversionReadCallback.OnAfterRead(stage, target)
		}
	case *Key:
		if stage.OnAfterKeyReadCallback != nil {
			stage.OnAfterKeyReadCallback.OnAfterRead(stage, target)
		}
	case *Key_accidental:
		if stage.OnAfterKey_accidentalReadCallback != nil {
			stage.OnAfterKey_accidentalReadCallback.OnAfterRead(stage, target)
		}
	case *Key_octave:
		if stage.OnAfterKey_octaveReadCallback != nil {
			stage.OnAfterKey_octaveReadCallback.OnAfterRead(stage, target)
		}
	case *Kind:
		if stage.OnAfterKindReadCallback != nil {
			stage.OnAfterKindReadCallback.OnAfterRead(stage, target)
		}
	case *Level:
		if stage.OnAfterLevelReadCallback != nil {
			stage.OnAfterLevelReadCallback.OnAfterRead(stage, target)
		}
	case *Line_detail:
		if stage.OnAfterLine_detailReadCallback != nil {
			stage.OnAfterLine_detailReadCallback.OnAfterRead(stage, target)
		}
	case *Line_width:
		if stage.OnAfterLine_widthReadCallback != nil {
			stage.OnAfterLine_widthReadCallback.OnAfterRead(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkReadCallback != nil {
			stage.OnAfterLinkReadCallback.OnAfterRead(stage, target)
		}
	case *Listen:
		if stage.OnAfterListenReadCallback != nil {
			stage.OnAfterListenReadCallback.OnAfterRead(stage, target)
		}
	case *Listening:
		if stage.OnAfterListeningReadCallback != nil {
			stage.OnAfterListeningReadCallback.OnAfterRead(stage, target)
		}
	case *Lyric:
		if stage.OnAfterLyricReadCallback != nil {
			stage.OnAfterLyricReadCallback.OnAfterRead(stage, target)
		}
	case *Lyric_font:
		if stage.OnAfterLyric_fontReadCallback != nil {
			stage.OnAfterLyric_fontReadCallback.OnAfterRead(stage, target)
		}
	case *Lyric_language:
		if stage.OnAfterLyric_languageReadCallback != nil {
			stage.OnAfterLyric_languageReadCallback.OnAfterRead(stage, target)
		}
	case *Measure_layout:
		if stage.OnAfterMeasure_layoutReadCallback != nil {
			stage.OnAfterMeasure_layoutReadCallback.OnAfterRead(stage, target)
		}
	case *Measure_numbering:
		if stage.OnAfterMeasure_numberingReadCallback != nil {
			stage.OnAfterMeasure_numberingReadCallback.OnAfterRead(stage, target)
		}
	case *Measure_repeat:
		if stage.OnAfterMeasure_repeatReadCallback != nil {
			stage.OnAfterMeasure_repeatReadCallback.OnAfterRead(stage, target)
		}
	case *Measure_style:
		if stage.OnAfterMeasure_styleReadCallback != nil {
			stage.OnAfterMeasure_styleReadCallback.OnAfterRead(stage, target)
		}
	case *Membrane:
		if stage.OnAfterMembraneReadCallback != nil {
			stage.OnAfterMembraneReadCallback.OnAfterRead(stage, target)
		}
	case *Metal:
		if stage.OnAfterMetalReadCallback != nil {
			stage.OnAfterMetalReadCallback.OnAfterRead(stage, target)
		}
	case *Metronome:
		if stage.OnAfterMetronomeReadCallback != nil {
			stage.OnAfterMetronomeReadCallback.OnAfterRead(stage, target)
		}
	case *Metronome_beam:
		if stage.OnAfterMetronome_beamReadCallback != nil {
			stage.OnAfterMetronome_beamReadCallback.OnAfterRead(stage, target)
		}
	case *Metronome_note:
		if stage.OnAfterMetronome_noteReadCallback != nil {
			stage.OnAfterMetronome_noteReadCallback.OnAfterRead(stage, target)
		}
	case *Metronome_tied:
		if stage.OnAfterMetronome_tiedReadCallback != nil {
			stage.OnAfterMetronome_tiedReadCallback.OnAfterRead(stage, target)
		}
	case *Metronome_tuplet:
		if stage.OnAfterMetronome_tupletReadCallback != nil {
			stage.OnAfterMetronome_tupletReadCallback.OnAfterRead(stage, target)
		}
	case *Midi_device:
		if stage.OnAfterMidi_deviceReadCallback != nil {
			stage.OnAfterMidi_deviceReadCallback.OnAfterRead(stage, target)
		}
	case *Midi_instrument:
		if stage.OnAfterMidi_instrumentReadCallback != nil {
			stage.OnAfterMidi_instrumentReadCallback.OnAfterRead(stage, target)
		}
	case *Miscellaneous:
		if stage.OnAfterMiscellaneousReadCallback != nil {
			stage.OnAfterMiscellaneousReadCallback.OnAfterRead(stage, target)
		}
	case *Miscellaneous_field:
		if stage.OnAfterMiscellaneous_fieldReadCallback != nil {
			stage.OnAfterMiscellaneous_fieldReadCallback.OnAfterRead(stage, target)
		}
	case *Mordent:
		if stage.OnAfterMordentReadCallback != nil {
			stage.OnAfterMordentReadCallback.OnAfterRead(stage, target)
		}
	case *Multiple_rest:
		if stage.OnAfterMultiple_restReadCallback != nil {
			stage.OnAfterMultiple_restReadCallback.OnAfterRead(stage, target)
		}
	case *Name_display:
		if stage.OnAfterName_displayReadCallback != nil {
			stage.OnAfterName_displayReadCallback.OnAfterRead(stage, target)
		}
	case *Non_arpeggiate:
		if stage.OnAfterNon_arpeggiateReadCallback != nil {
			stage.OnAfterNon_arpeggiateReadCallback.OnAfterRead(stage, target)
		}
	case *Notations:
		if stage.OnAfterNotationsReadCallback != nil {
			stage.OnAfterNotationsReadCallback.OnAfterRead(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteReadCallback != nil {
			stage.OnAfterNoteReadCallback.OnAfterRead(stage, target)
		}
	case *Note_size:
		if stage.OnAfterNote_sizeReadCallback != nil {
			stage.OnAfterNote_sizeReadCallback.OnAfterRead(stage, target)
		}
	case *Note_type:
		if stage.OnAfterNote_typeReadCallback != nil {
			stage.OnAfterNote_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Notehead:
		if stage.OnAfterNoteheadReadCallback != nil {
			stage.OnAfterNoteheadReadCallback.OnAfterRead(stage, target)
		}
	case *Notehead_text:
		if stage.OnAfterNotehead_textReadCallback != nil {
			stage.OnAfterNotehead_textReadCallback.OnAfterRead(stage, target)
		}
	case *Numeral:
		if stage.OnAfterNumeralReadCallback != nil {
			stage.OnAfterNumeralReadCallback.OnAfterRead(stage, target)
		}
	case *Numeral_key:
		if stage.OnAfterNumeral_keyReadCallback != nil {
			stage.OnAfterNumeral_keyReadCallback.OnAfterRead(stage, target)
		}
	case *Numeral_root:
		if stage.OnAfterNumeral_rootReadCallback != nil {
			stage.OnAfterNumeral_rootReadCallback.OnAfterRead(stage, target)
		}
	case *Octave_shift:
		if stage.OnAfterOctave_shiftReadCallback != nil {
			stage.OnAfterOctave_shiftReadCallback.OnAfterRead(stage, target)
		}
	case *Offset:
		if stage.OnAfterOffsetReadCallback != nil {
			stage.OnAfterOffsetReadCallback.OnAfterRead(stage, target)
		}
	case *Opus:
		if stage.OnAfterOpusReadCallback != nil {
			stage.OnAfterOpusReadCallback.OnAfterRead(stage, target)
		}
	case *Ornaments:
		if stage.OnAfterOrnamentsReadCallback != nil {
			stage.OnAfterOrnamentsReadCallback.OnAfterRead(stage, target)
		}
	case *Other_appearance:
		if stage.OnAfterOther_appearanceReadCallback != nil {
			stage.OnAfterOther_appearanceReadCallback.OnAfterRead(stage, target)
		}
	case *Other_direction:
		if stage.OnAfterOther_directionReadCallback != nil {
			stage.OnAfterOther_directionReadCallback.OnAfterRead(stage, target)
		}
	case *Other_listening:
		if stage.OnAfterOther_listeningReadCallback != nil {
			stage.OnAfterOther_listeningReadCallback.OnAfterRead(stage, target)
		}
	case *Other_notation:
		if stage.OnAfterOther_notationReadCallback != nil {
			stage.OnAfterOther_notationReadCallback.OnAfterRead(stage, target)
		}
	case *Other_placement_text:
		if stage.OnAfterOther_placement_textReadCallback != nil {
			stage.OnAfterOther_placement_textReadCallback.OnAfterRead(stage, target)
		}
	case *Other_play:
		if stage.OnAfterOther_playReadCallback != nil {
			stage.OnAfterOther_playReadCallback.OnAfterRead(stage, target)
		}
	case *Other_text:
		if stage.OnAfterOther_textReadCallback != nil {
			stage.OnAfterOther_textReadCallback.OnAfterRead(stage, target)
		}
	case *Page_layout:
		if stage.OnAfterPage_layoutReadCallback != nil {
			stage.OnAfterPage_layoutReadCallback.OnAfterRead(stage, target)
		}
	case *Page_margins:
		if stage.OnAfterPage_marginsReadCallback != nil {
			stage.OnAfterPage_marginsReadCallback.OnAfterRead(stage, target)
		}
	case *Part_clef:
		if stage.OnAfterPart_clefReadCallback != nil {
			stage.OnAfterPart_clefReadCallback.OnAfterRead(stage, target)
		}
	case *Part_group:
		if stage.OnAfterPart_groupReadCallback != nil {
			stage.OnAfterPart_groupReadCallback.OnAfterRead(stage, target)
		}
	case *Part_link:
		if stage.OnAfterPart_linkReadCallback != nil {
			stage.OnAfterPart_linkReadCallback.OnAfterRead(stage, target)
		}
	case *Part_list:
		if stage.OnAfterPart_listReadCallback != nil {
			stage.OnAfterPart_listReadCallback.OnAfterRead(stage, target)
		}
	case *Part_name:
		if stage.OnAfterPart_nameReadCallback != nil {
			stage.OnAfterPart_nameReadCallback.OnAfterRead(stage, target)
		}
	case *Part_symbol:
		if stage.OnAfterPart_symbolReadCallback != nil {
			stage.OnAfterPart_symbolReadCallback.OnAfterRead(stage, target)
		}
	case *Part_transpose:
		if stage.OnAfterPart_transposeReadCallback != nil {
			stage.OnAfterPart_transposeReadCallback.OnAfterRead(stage, target)
		}
	case *Pedal:
		if stage.OnAfterPedalReadCallback != nil {
			stage.OnAfterPedalReadCallback.OnAfterRead(stage, target)
		}
	case *Pedal_tuning:
		if stage.OnAfterPedal_tuningReadCallback != nil {
			stage.OnAfterPedal_tuningReadCallback.OnAfterRead(stage, target)
		}
	case *Per_minute:
		if stage.OnAfterPer_minuteReadCallback != nil {
			stage.OnAfterPer_minuteReadCallback.OnAfterRead(stage, target)
		}
	case *Percussion:
		if stage.OnAfterPercussionReadCallback != nil {
			stage.OnAfterPercussionReadCallback.OnAfterRead(stage, target)
		}
	case *Pitch:
		if stage.OnAfterPitchReadCallback != nil {
			stage.OnAfterPitchReadCallback.OnAfterRead(stage, target)
		}
	case *Pitched:
		if stage.OnAfterPitchedReadCallback != nil {
			stage.OnAfterPitchedReadCallback.OnAfterRead(stage, target)
		}
	case *Placement_text:
		if stage.OnAfterPlacement_textReadCallback != nil {
			stage.OnAfterPlacement_textReadCallback.OnAfterRead(stage, target)
		}
	case *Play:
		if stage.OnAfterPlayReadCallback != nil {
			stage.OnAfterPlayReadCallback.OnAfterRead(stage, target)
		}
	case *Player:
		if stage.OnAfterPlayerReadCallback != nil {
			stage.OnAfterPlayerReadCallback.OnAfterRead(stage, target)
		}
	case *Principal_voice:
		if stage.OnAfterPrincipal_voiceReadCallback != nil {
			stage.OnAfterPrincipal_voiceReadCallback.OnAfterRead(stage, target)
		}
	case *Print:
		if stage.OnAfterPrintReadCallback != nil {
			stage.OnAfterPrintReadCallback.OnAfterRead(stage, target)
		}
	case *Release:
		if stage.OnAfterReleaseReadCallback != nil {
			stage.OnAfterReleaseReadCallback.OnAfterRead(stage, target)
		}
	case *Repeat:
		if stage.OnAfterRepeatReadCallback != nil {
			stage.OnAfterRepeatReadCallback.OnAfterRead(stage, target)
		}
	case *Rest:
		if stage.OnAfterRestReadCallback != nil {
			stage.OnAfterRestReadCallback.OnAfterRead(stage, target)
		}
	case *Root:
		if stage.OnAfterRootReadCallback != nil {
			stage.OnAfterRootReadCallback.OnAfterRead(stage, target)
		}
	case *Root_step:
		if stage.OnAfterRoot_stepReadCallback != nil {
			stage.OnAfterRoot_stepReadCallback.OnAfterRead(stage, target)
		}
	case *Scaling:
		if stage.OnAfterScalingReadCallback != nil {
			stage.OnAfterScalingReadCallback.OnAfterRead(stage, target)
		}
	case *Scordatura:
		if stage.OnAfterScordaturaReadCallback != nil {
			stage.OnAfterScordaturaReadCallback.OnAfterRead(stage, target)
		}
	case *Score_instrument:
		if stage.OnAfterScore_instrumentReadCallback != nil {
			stage.OnAfterScore_instrumentReadCallback.OnAfterRead(stage, target)
		}
	case *Score_part:
		if stage.OnAfterScore_partReadCallback != nil {
			stage.OnAfterScore_partReadCallback.OnAfterRead(stage, target)
		}
	case *Score_partwise:
		if stage.OnAfterScore_partwiseReadCallback != nil {
			stage.OnAfterScore_partwiseReadCallback.OnAfterRead(stage, target)
		}
	case *Score_timewise:
		if stage.OnAfterScore_timewiseReadCallback != nil {
			stage.OnAfterScore_timewiseReadCallback.OnAfterRead(stage, target)
		}
	case *Segno:
		if stage.OnAfterSegnoReadCallback != nil {
			stage.OnAfterSegnoReadCallback.OnAfterRead(stage, target)
		}
	case *Slash:
		if stage.OnAfterSlashReadCallback != nil {
			stage.OnAfterSlashReadCallback.OnAfterRead(stage, target)
		}
	case *Slide:
		if stage.OnAfterSlideReadCallback != nil {
			stage.OnAfterSlideReadCallback.OnAfterRead(stage, target)
		}
	case *Slur:
		if stage.OnAfterSlurReadCallback != nil {
			stage.OnAfterSlurReadCallback.OnAfterRead(stage, target)
		}
	case *Sound:
		if stage.OnAfterSoundReadCallback != nil {
			stage.OnAfterSoundReadCallback.OnAfterRead(stage, target)
		}
	case *Staff_details:
		if stage.OnAfterStaff_detailsReadCallback != nil {
			stage.OnAfterStaff_detailsReadCallback.OnAfterRead(stage, target)
		}
	case *Staff_divide:
		if stage.OnAfterStaff_divideReadCallback != nil {
			stage.OnAfterStaff_divideReadCallback.OnAfterRead(stage, target)
		}
	case *Staff_layout:
		if stage.OnAfterStaff_layoutReadCallback != nil {
			stage.OnAfterStaff_layoutReadCallback.OnAfterRead(stage, target)
		}
	case *Staff_size:
		if stage.OnAfterStaff_sizeReadCallback != nil {
			stage.OnAfterStaff_sizeReadCallback.OnAfterRead(stage, target)
		}
	case *Staff_tuning:
		if stage.OnAfterStaff_tuningReadCallback != nil {
			stage.OnAfterStaff_tuningReadCallback.OnAfterRead(stage, target)
		}
	case *Stem:
		if stage.OnAfterStemReadCallback != nil {
			stage.OnAfterStemReadCallback.OnAfterRead(stage, target)
		}
	case *Stick:
		if stage.OnAfterStickReadCallback != nil {
			stage.OnAfterStickReadCallback.OnAfterRead(stage, target)
		}
	case *String_mute:
		if stage.OnAfterString_muteReadCallback != nil {
			stage.OnAfterString_muteReadCallback.OnAfterRead(stage, target)
		}
	case *String_type:
		if stage.OnAfterString_typeReadCallback != nil {
			stage.OnAfterString_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Strong_accent:
		if stage.OnAfterStrong_accentReadCallback != nil {
			stage.OnAfterStrong_accentReadCallback.OnAfterRead(stage, target)
		}
	case *Style_text:
		if stage.OnAfterStyle_textReadCallback != nil {
			stage.OnAfterStyle_textReadCallback.OnAfterRead(stage, target)
		}
	case *Supports:
		if stage.OnAfterSupportsReadCallback != nil {
			stage.OnAfterSupportsReadCallback.OnAfterRead(stage, target)
		}
	case *Swing:
		if stage.OnAfterSwingReadCallback != nil {
			stage.OnAfterSwingReadCallback.OnAfterRead(stage, target)
		}
	case *Sync:
		if stage.OnAfterSyncReadCallback != nil {
			stage.OnAfterSyncReadCallback.OnAfterRead(stage, target)
		}
	case *System_dividers:
		if stage.OnAfterSystem_dividersReadCallback != nil {
			stage.OnAfterSystem_dividersReadCallback.OnAfterRead(stage, target)
		}
	case *System_layout:
		if stage.OnAfterSystem_layoutReadCallback != nil {
			stage.OnAfterSystem_layoutReadCallback.OnAfterRead(stage, target)
		}
	case *System_margins:
		if stage.OnAfterSystem_marginsReadCallback != nil {
			stage.OnAfterSystem_marginsReadCallback.OnAfterRead(stage, target)
		}
	case *Tap:
		if stage.OnAfterTapReadCallback != nil {
			stage.OnAfterTapReadCallback.OnAfterRead(stage, target)
		}
	case *Technical:
		if stage.OnAfterTechnicalReadCallback != nil {
			stage.OnAfterTechnicalReadCallback.OnAfterRead(stage, target)
		}
	case *Text_element_data:
		if stage.OnAfterText_element_dataReadCallback != nil {
			stage.OnAfterText_element_dataReadCallback.OnAfterRead(stage, target)
		}
	case *Tie:
		if stage.OnAfterTieReadCallback != nil {
			stage.OnAfterTieReadCallback.OnAfterRead(stage, target)
		}
	case *Tied:
		if stage.OnAfterTiedReadCallback != nil {
			stage.OnAfterTiedReadCallback.OnAfterRead(stage, target)
		}
	case *Time:
		if stage.OnAfterTimeReadCallback != nil {
			stage.OnAfterTimeReadCallback.OnAfterRead(stage, target)
		}
	case *Time_modification:
		if stage.OnAfterTime_modificationReadCallback != nil {
			stage.OnAfterTime_modificationReadCallback.OnAfterRead(stage, target)
		}
	case *Timpani:
		if stage.OnAfterTimpaniReadCallback != nil {
			stage.OnAfterTimpaniReadCallback.OnAfterRead(stage, target)
		}
	case *Transpose:
		if stage.OnAfterTransposeReadCallback != nil {
			stage.OnAfterTransposeReadCallback.OnAfterRead(stage, target)
		}
	case *Tremolo:
		if stage.OnAfterTremoloReadCallback != nil {
			stage.OnAfterTremoloReadCallback.OnAfterRead(stage, target)
		}
	case *Tuplet:
		if stage.OnAfterTupletReadCallback != nil {
			stage.OnAfterTupletReadCallback.OnAfterRead(stage, target)
		}
	case *Tuplet_dot:
		if stage.OnAfterTuplet_dotReadCallback != nil {
			stage.OnAfterTuplet_dotReadCallback.OnAfterRead(stage, target)
		}
	case *Tuplet_number:
		if stage.OnAfterTuplet_numberReadCallback != nil {
			stage.OnAfterTuplet_numberReadCallback.OnAfterRead(stage, target)
		}
	case *Tuplet_portion:
		if stage.OnAfterTuplet_portionReadCallback != nil {
			stage.OnAfterTuplet_portionReadCallback.OnAfterRead(stage, target)
		}
	case *Tuplet_type:
		if stage.OnAfterTuplet_typeReadCallback != nil {
			stage.OnAfterTuplet_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Typed_text:
		if stage.OnAfterTyped_textReadCallback != nil {
			stage.OnAfterTyped_textReadCallback.OnAfterRead(stage, target)
		}
	case *Unpitched:
		if stage.OnAfterUnpitchedReadCallback != nil {
			stage.OnAfterUnpitchedReadCallback.OnAfterRead(stage, target)
		}
	case *Virtual_instrument:
		if stage.OnAfterVirtual_instrumentReadCallback != nil {
			stage.OnAfterVirtual_instrumentReadCallback.OnAfterRead(stage, target)
		}
	case *Wait:
		if stage.OnAfterWaitReadCallback != nil {
			stage.OnAfterWaitReadCallback.OnAfterRead(stage, target)
		}
	case *Wavy_line:
		if stage.OnAfterWavy_lineReadCallback != nil {
			stage.OnAfterWavy_lineReadCallback.OnAfterRead(stage, target)
		}
	case *Wedge:
		if stage.OnAfterWedgeReadCallback != nil {
			stage.OnAfterWedgeReadCallback.OnAfterRead(stage, target)
		}
	case *Wood:
		if stage.OnAfterWoodReadCallback != nil {
			stage.OnAfterWoodReadCallback.OnAfterRead(stage, target)
		}
	case *Work:
		if stage.OnAfterWorkReadCallback != nil {
			stage.OnAfterWorkReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *A_directive:
		stage.OnAfterA_directiveUpdateCallback = any(callback).(OnAfterUpdateInterface[A_directive])
	case *A_measure:
		stage.OnAfterA_measureUpdateCallback = any(callback).(OnAfterUpdateInterface[A_measure])
	case *A_measure_1:
		stage.OnAfterA_measure_1UpdateCallback = any(callback).(OnAfterUpdateInterface[A_measure_1])
	case *A_part:
		stage.OnAfterA_partUpdateCallback = any(callback).(OnAfterUpdateInterface[A_part])
	case *A_part_1:
		stage.OnAfterA_part_1UpdateCallback = any(callback).(OnAfterUpdateInterface[A_part_1])
	case *Accidental:
		stage.OnAfterAccidentalUpdateCallback = any(callback).(OnAfterUpdateInterface[Accidental])
	case *Accidental_mark:
		stage.OnAfterAccidental_markUpdateCallback = any(callback).(OnAfterUpdateInterface[Accidental_mark])
	case *Accidental_text:
		stage.OnAfterAccidental_textUpdateCallback = any(callback).(OnAfterUpdateInterface[Accidental_text])
	case *Accord:
		stage.OnAfterAccordUpdateCallback = any(callback).(OnAfterUpdateInterface[Accord])
	case *Accordion_registration:
		stage.OnAfterAccordion_registrationUpdateCallback = any(callback).(OnAfterUpdateInterface[Accordion_registration])
	case *Appearance:
		stage.OnAfterAppearanceUpdateCallback = any(callback).(OnAfterUpdateInterface[Appearance])
	case *Arpeggiate:
		stage.OnAfterArpeggiateUpdateCallback = any(callback).(OnAfterUpdateInterface[Arpeggiate])
	case *Arrow:
		stage.OnAfterArrowUpdateCallback = any(callback).(OnAfterUpdateInterface[Arrow])
	case *Articulations:
		stage.OnAfterArticulationsUpdateCallback = any(callback).(OnAfterUpdateInterface[Articulations])
	case *Assess:
		stage.OnAfterAssessUpdateCallback = any(callback).(OnAfterUpdateInterface[Assess])
	case *Attributes:
		stage.OnAfterAttributesUpdateCallback = any(callback).(OnAfterUpdateInterface[Attributes])
	case *Backup:
		stage.OnAfterBackupUpdateCallback = any(callback).(OnAfterUpdateInterface[Backup])
	case *Bar_style_color:
		stage.OnAfterBar_style_colorUpdateCallback = any(callback).(OnAfterUpdateInterface[Bar_style_color])
	case *Barline:
		stage.OnAfterBarlineUpdateCallback = any(callback).(OnAfterUpdateInterface[Barline])
	case *Barre:
		stage.OnAfterBarreUpdateCallback = any(callback).(OnAfterUpdateInterface[Barre])
	case *Bass:
		stage.OnAfterBassUpdateCallback = any(callback).(OnAfterUpdateInterface[Bass])
	case *Bass_step:
		stage.OnAfterBass_stepUpdateCallback = any(callback).(OnAfterUpdateInterface[Bass_step])
	case *Beam:
		stage.OnAfterBeamUpdateCallback = any(callback).(OnAfterUpdateInterface[Beam])
	case *Beat_repeat:
		stage.OnAfterBeat_repeatUpdateCallback = any(callback).(OnAfterUpdateInterface[Beat_repeat])
	case *Beat_unit_tied:
		stage.OnAfterBeat_unit_tiedUpdateCallback = any(callback).(OnAfterUpdateInterface[Beat_unit_tied])
	case *Beater:
		stage.OnAfterBeaterUpdateCallback = any(callback).(OnAfterUpdateInterface[Beater])
	case *Bend:
		stage.OnAfterBendUpdateCallback = any(callback).(OnAfterUpdateInterface[Bend])
	case *Bookmark:
		stage.OnAfterBookmarkUpdateCallback = any(callback).(OnAfterUpdateInterface[Bookmark])
	case *Bracket:
		stage.OnAfterBracketUpdateCallback = any(callback).(OnAfterUpdateInterface[Bracket])
	case *Breath_mark:
		stage.OnAfterBreath_markUpdateCallback = any(callback).(OnAfterUpdateInterface[Breath_mark])
	case *Caesura:
		stage.OnAfterCaesuraUpdateCallback = any(callback).(OnAfterUpdateInterface[Caesura])
	case *Cancel:
		stage.OnAfterCancelUpdateCallback = any(callback).(OnAfterUpdateInterface[Cancel])
	case *Clef:
		stage.OnAfterClefUpdateCallback = any(callback).(OnAfterUpdateInterface[Clef])
	case *Coda:
		stage.OnAfterCodaUpdateCallback = any(callback).(OnAfterUpdateInterface[Coda])
	case *Credit:
		stage.OnAfterCreditUpdateCallback = any(callback).(OnAfterUpdateInterface[Credit])
	case *Dashes:
		stage.OnAfterDashesUpdateCallback = any(callback).(OnAfterUpdateInterface[Dashes])
	case *Defaults:
		stage.OnAfterDefaultsUpdateCallback = any(callback).(OnAfterUpdateInterface[Defaults])
	case *Degree:
		stage.OnAfterDegreeUpdateCallback = any(callback).(OnAfterUpdateInterface[Degree])
	case *Degree_alter:
		stage.OnAfterDegree_alterUpdateCallback = any(callback).(OnAfterUpdateInterface[Degree_alter])
	case *Degree_type:
		stage.OnAfterDegree_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Degree_type])
	case *Degree_value:
		stage.OnAfterDegree_valueUpdateCallback = any(callback).(OnAfterUpdateInterface[Degree_value])
	case *Direction:
		stage.OnAfterDirectionUpdateCallback = any(callback).(OnAfterUpdateInterface[Direction])
	case *Direction_type:
		stage.OnAfterDirection_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Direction_type])
	case *Distance:
		stage.OnAfterDistanceUpdateCallback = any(callback).(OnAfterUpdateInterface[Distance])
	case *Double:
		stage.OnAfterDoubleUpdateCallback = any(callback).(OnAfterUpdateInterface[Double])
	case *Dynamics:
		stage.OnAfterDynamicsUpdateCallback = any(callback).(OnAfterUpdateInterface[Dynamics])
	case *Effect:
		stage.OnAfterEffectUpdateCallback = any(callback).(OnAfterUpdateInterface[Effect])
	case *Elision:
		stage.OnAfterElisionUpdateCallback = any(callback).(OnAfterUpdateInterface[Elision])
	case *Empty:
		stage.OnAfterEmptyUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty])
	case *Empty_font:
		stage.OnAfterEmpty_fontUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty_font])
	case *Empty_line:
		stage.OnAfterEmpty_lineUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty_line])
	case *Empty_placement:
		stage.OnAfterEmpty_placementUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty_placement])
	case *Empty_placement_smufl:
		stage.OnAfterEmpty_placement_smuflUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty_placement_smufl])
	case *Empty_print_object_style_align:
		stage.OnAfterEmpty_print_object_style_alignUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty_print_object_style_align])
	case *Empty_print_style:
		stage.OnAfterEmpty_print_styleUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty_print_style])
	case *Empty_print_style_align:
		stage.OnAfterEmpty_print_style_alignUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty_print_style_align])
	case *Empty_print_style_align_id:
		stage.OnAfterEmpty_print_style_align_idUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty_print_style_align_id])
	case *Empty_trill_sound:
		stage.OnAfterEmpty_trill_soundUpdateCallback = any(callback).(OnAfterUpdateInterface[Empty_trill_sound])
	case *Encoding:
		stage.OnAfterEncodingUpdateCallback = any(callback).(OnAfterUpdateInterface[Encoding])
	case *Ending:
		stage.OnAfterEndingUpdateCallback = any(callback).(OnAfterUpdateInterface[Ending])
	case *Extend:
		stage.OnAfterExtendUpdateCallback = any(callback).(OnAfterUpdateInterface[Extend])
	case *Feature:
		stage.OnAfterFeatureUpdateCallback = any(callback).(OnAfterUpdateInterface[Feature])
	case *Fermata:
		stage.OnAfterFermataUpdateCallback = any(callback).(OnAfterUpdateInterface[Fermata])
	case *Figure:
		stage.OnAfterFigureUpdateCallback = any(callback).(OnAfterUpdateInterface[Figure])
	case *Figured_bass:
		stage.OnAfterFigured_bassUpdateCallback = any(callback).(OnAfterUpdateInterface[Figured_bass])
	case *Fingering:
		stage.OnAfterFingeringUpdateCallback = any(callback).(OnAfterUpdateInterface[Fingering])
	case *First_fret:
		stage.OnAfterFirst_fretUpdateCallback = any(callback).(OnAfterUpdateInterface[First_fret])
	case *For_part:
		stage.OnAfterFor_partUpdateCallback = any(callback).(OnAfterUpdateInterface[For_part])
	case *Formatted_symbol:
		stage.OnAfterFormatted_symbolUpdateCallback = any(callback).(OnAfterUpdateInterface[Formatted_symbol])
	case *Formatted_symbol_id:
		stage.OnAfterFormatted_symbol_idUpdateCallback = any(callback).(OnAfterUpdateInterface[Formatted_symbol_id])
	case *Formatted_text:
		stage.OnAfterFormatted_textUpdateCallback = any(callback).(OnAfterUpdateInterface[Formatted_text])
	case *Formatted_text_id:
		stage.OnAfterFormatted_text_idUpdateCallback = any(callback).(OnAfterUpdateInterface[Formatted_text_id])
	case *Forward:
		stage.OnAfterForwardUpdateCallback = any(callback).(OnAfterUpdateInterface[Forward])
	case *Frame:
		stage.OnAfterFrameUpdateCallback = any(callback).(OnAfterUpdateInterface[Frame])
	case *Frame_note:
		stage.OnAfterFrame_noteUpdateCallback = any(callback).(OnAfterUpdateInterface[Frame_note])
	case *Fret:
		stage.OnAfterFretUpdateCallback = any(callback).(OnAfterUpdateInterface[Fret])
	case *Glass:
		stage.OnAfterGlassUpdateCallback = any(callback).(OnAfterUpdateInterface[Glass])
	case *Glissando:
		stage.OnAfterGlissandoUpdateCallback = any(callback).(OnAfterUpdateInterface[Glissando])
	case *Glyph:
		stage.OnAfterGlyphUpdateCallback = any(callback).(OnAfterUpdateInterface[Glyph])
	case *Grace:
		stage.OnAfterGraceUpdateCallback = any(callback).(OnAfterUpdateInterface[Grace])
	case *Group_barline:
		stage.OnAfterGroup_barlineUpdateCallback = any(callback).(OnAfterUpdateInterface[Group_barline])
	case *Group_name:
		stage.OnAfterGroup_nameUpdateCallback = any(callback).(OnAfterUpdateInterface[Group_name])
	case *Group_symbol:
		stage.OnAfterGroup_symbolUpdateCallback = any(callback).(OnAfterUpdateInterface[Group_symbol])
	case *Grouping:
		stage.OnAfterGroupingUpdateCallback = any(callback).(OnAfterUpdateInterface[Grouping])
	case *Hammer_on_pull_off:
		stage.OnAfterHammer_on_pull_offUpdateCallback = any(callback).(OnAfterUpdateInterface[Hammer_on_pull_off])
	case *Handbell:
		stage.OnAfterHandbellUpdateCallback = any(callback).(OnAfterUpdateInterface[Handbell])
	case *Harmon_closed:
		stage.OnAfterHarmon_closedUpdateCallback = any(callback).(OnAfterUpdateInterface[Harmon_closed])
	case *Harmon_mute:
		stage.OnAfterHarmon_muteUpdateCallback = any(callback).(OnAfterUpdateInterface[Harmon_mute])
	case *Harmonic:
		stage.OnAfterHarmonicUpdateCallback = any(callback).(OnAfterUpdateInterface[Harmonic])
	case *Harmony:
		stage.OnAfterHarmonyUpdateCallback = any(callback).(OnAfterUpdateInterface[Harmony])
	case *Harmony_alter:
		stage.OnAfterHarmony_alterUpdateCallback = any(callback).(OnAfterUpdateInterface[Harmony_alter])
	case *Harp_pedals:
		stage.OnAfterHarp_pedalsUpdateCallback = any(callback).(OnAfterUpdateInterface[Harp_pedals])
	case *Heel_toe:
		stage.OnAfterHeel_toeUpdateCallback = any(callback).(OnAfterUpdateInterface[Heel_toe])
	case *Hole:
		stage.OnAfterHoleUpdateCallback = any(callback).(OnAfterUpdateInterface[Hole])
	case *Hole_closed:
		stage.OnAfterHole_closedUpdateCallback = any(callback).(OnAfterUpdateInterface[Hole_closed])
	case *Horizontal_turn:
		stage.OnAfterHorizontal_turnUpdateCallback = any(callback).(OnAfterUpdateInterface[Horizontal_turn])
	case *Identification:
		stage.OnAfterIdentificationUpdateCallback = any(callback).(OnAfterUpdateInterface[Identification])
	case *Image:
		stage.OnAfterImageUpdateCallback = any(callback).(OnAfterUpdateInterface[Image])
	case *Instrument:
		stage.OnAfterInstrumentUpdateCallback = any(callback).(OnAfterUpdateInterface[Instrument])
	case *Instrument_change:
		stage.OnAfterInstrument_changeUpdateCallback = any(callback).(OnAfterUpdateInterface[Instrument_change])
	case *Instrument_link:
		stage.OnAfterInstrument_linkUpdateCallback = any(callback).(OnAfterUpdateInterface[Instrument_link])
	case *Interchangeable:
		stage.OnAfterInterchangeableUpdateCallback = any(callback).(OnAfterUpdateInterface[Interchangeable])
	case *Inversion:
		stage.OnAfterInversionUpdateCallback = any(callback).(OnAfterUpdateInterface[Inversion])
	case *Key:
		stage.OnAfterKeyUpdateCallback = any(callback).(OnAfterUpdateInterface[Key])
	case *Key_accidental:
		stage.OnAfterKey_accidentalUpdateCallback = any(callback).(OnAfterUpdateInterface[Key_accidental])
	case *Key_octave:
		stage.OnAfterKey_octaveUpdateCallback = any(callback).(OnAfterUpdateInterface[Key_octave])
	case *Kind:
		stage.OnAfterKindUpdateCallback = any(callback).(OnAfterUpdateInterface[Kind])
	case *Level:
		stage.OnAfterLevelUpdateCallback = any(callback).(OnAfterUpdateInterface[Level])
	case *Line_detail:
		stage.OnAfterLine_detailUpdateCallback = any(callback).(OnAfterUpdateInterface[Line_detail])
	case *Line_width:
		stage.OnAfterLine_widthUpdateCallback = any(callback).(OnAfterUpdateInterface[Line_width])
	case *Link:
		stage.OnAfterLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[Link])
	case *Listen:
		stage.OnAfterListenUpdateCallback = any(callback).(OnAfterUpdateInterface[Listen])
	case *Listening:
		stage.OnAfterListeningUpdateCallback = any(callback).(OnAfterUpdateInterface[Listening])
	case *Lyric:
		stage.OnAfterLyricUpdateCallback = any(callback).(OnAfterUpdateInterface[Lyric])
	case *Lyric_font:
		stage.OnAfterLyric_fontUpdateCallback = any(callback).(OnAfterUpdateInterface[Lyric_font])
	case *Lyric_language:
		stage.OnAfterLyric_languageUpdateCallback = any(callback).(OnAfterUpdateInterface[Lyric_language])
	case *Measure_layout:
		stage.OnAfterMeasure_layoutUpdateCallback = any(callback).(OnAfterUpdateInterface[Measure_layout])
	case *Measure_numbering:
		stage.OnAfterMeasure_numberingUpdateCallback = any(callback).(OnAfterUpdateInterface[Measure_numbering])
	case *Measure_repeat:
		stage.OnAfterMeasure_repeatUpdateCallback = any(callback).(OnAfterUpdateInterface[Measure_repeat])
	case *Measure_style:
		stage.OnAfterMeasure_styleUpdateCallback = any(callback).(OnAfterUpdateInterface[Measure_style])
	case *Membrane:
		stage.OnAfterMembraneUpdateCallback = any(callback).(OnAfterUpdateInterface[Membrane])
	case *Metal:
		stage.OnAfterMetalUpdateCallback = any(callback).(OnAfterUpdateInterface[Metal])
	case *Metronome:
		stage.OnAfterMetronomeUpdateCallback = any(callback).(OnAfterUpdateInterface[Metronome])
	case *Metronome_beam:
		stage.OnAfterMetronome_beamUpdateCallback = any(callback).(OnAfterUpdateInterface[Metronome_beam])
	case *Metronome_note:
		stage.OnAfterMetronome_noteUpdateCallback = any(callback).(OnAfterUpdateInterface[Metronome_note])
	case *Metronome_tied:
		stage.OnAfterMetronome_tiedUpdateCallback = any(callback).(OnAfterUpdateInterface[Metronome_tied])
	case *Metronome_tuplet:
		stage.OnAfterMetronome_tupletUpdateCallback = any(callback).(OnAfterUpdateInterface[Metronome_tuplet])
	case *Midi_device:
		stage.OnAfterMidi_deviceUpdateCallback = any(callback).(OnAfterUpdateInterface[Midi_device])
	case *Midi_instrument:
		stage.OnAfterMidi_instrumentUpdateCallback = any(callback).(OnAfterUpdateInterface[Midi_instrument])
	case *Miscellaneous:
		stage.OnAfterMiscellaneousUpdateCallback = any(callback).(OnAfterUpdateInterface[Miscellaneous])
	case *Miscellaneous_field:
		stage.OnAfterMiscellaneous_fieldUpdateCallback = any(callback).(OnAfterUpdateInterface[Miscellaneous_field])
	case *Mordent:
		stage.OnAfterMordentUpdateCallback = any(callback).(OnAfterUpdateInterface[Mordent])
	case *Multiple_rest:
		stage.OnAfterMultiple_restUpdateCallback = any(callback).(OnAfterUpdateInterface[Multiple_rest])
	case *Name_display:
		stage.OnAfterName_displayUpdateCallback = any(callback).(OnAfterUpdateInterface[Name_display])
	case *Non_arpeggiate:
		stage.OnAfterNon_arpeggiateUpdateCallback = any(callback).(OnAfterUpdateInterface[Non_arpeggiate])
	case *Notations:
		stage.OnAfterNotationsUpdateCallback = any(callback).(OnAfterUpdateInterface[Notations])
	case *Note:
		stage.OnAfterNoteUpdateCallback = any(callback).(OnAfterUpdateInterface[Note])
	case *Note_size:
		stage.OnAfterNote_sizeUpdateCallback = any(callback).(OnAfterUpdateInterface[Note_size])
	case *Note_type:
		stage.OnAfterNote_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Note_type])
	case *Notehead:
		stage.OnAfterNoteheadUpdateCallback = any(callback).(OnAfterUpdateInterface[Notehead])
	case *Notehead_text:
		stage.OnAfterNotehead_textUpdateCallback = any(callback).(OnAfterUpdateInterface[Notehead_text])
	case *Numeral:
		stage.OnAfterNumeralUpdateCallback = any(callback).(OnAfterUpdateInterface[Numeral])
	case *Numeral_key:
		stage.OnAfterNumeral_keyUpdateCallback = any(callback).(OnAfterUpdateInterface[Numeral_key])
	case *Numeral_root:
		stage.OnAfterNumeral_rootUpdateCallback = any(callback).(OnAfterUpdateInterface[Numeral_root])
	case *Octave_shift:
		stage.OnAfterOctave_shiftUpdateCallback = any(callback).(OnAfterUpdateInterface[Octave_shift])
	case *Offset:
		stage.OnAfterOffsetUpdateCallback = any(callback).(OnAfterUpdateInterface[Offset])
	case *Opus:
		stage.OnAfterOpusUpdateCallback = any(callback).(OnAfterUpdateInterface[Opus])
	case *Ornaments:
		stage.OnAfterOrnamentsUpdateCallback = any(callback).(OnAfterUpdateInterface[Ornaments])
	case *Other_appearance:
		stage.OnAfterOther_appearanceUpdateCallback = any(callback).(OnAfterUpdateInterface[Other_appearance])
	case *Other_direction:
		stage.OnAfterOther_directionUpdateCallback = any(callback).(OnAfterUpdateInterface[Other_direction])
	case *Other_listening:
		stage.OnAfterOther_listeningUpdateCallback = any(callback).(OnAfterUpdateInterface[Other_listening])
	case *Other_notation:
		stage.OnAfterOther_notationUpdateCallback = any(callback).(OnAfterUpdateInterface[Other_notation])
	case *Other_placement_text:
		stage.OnAfterOther_placement_textUpdateCallback = any(callback).(OnAfterUpdateInterface[Other_placement_text])
	case *Other_play:
		stage.OnAfterOther_playUpdateCallback = any(callback).(OnAfterUpdateInterface[Other_play])
	case *Other_text:
		stage.OnAfterOther_textUpdateCallback = any(callback).(OnAfterUpdateInterface[Other_text])
	case *Page_layout:
		stage.OnAfterPage_layoutUpdateCallback = any(callback).(OnAfterUpdateInterface[Page_layout])
	case *Page_margins:
		stage.OnAfterPage_marginsUpdateCallback = any(callback).(OnAfterUpdateInterface[Page_margins])
	case *Part_clef:
		stage.OnAfterPart_clefUpdateCallback = any(callback).(OnAfterUpdateInterface[Part_clef])
	case *Part_group:
		stage.OnAfterPart_groupUpdateCallback = any(callback).(OnAfterUpdateInterface[Part_group])
	case *Part_link:
		stage.OnAfterPart_linkUpdateCallback = any(callback).(OnAfterUpdateInterface[Part_link])
	case *Part_list:
		stage.OnAfterPart_listUpdateCallback = any(callback).(OnAfterUpdateInterface[Part_list])
	case *Part_name:
		stage.OnAfterPart_nameUpdateCallback = any(callback).(OnAfterUpdateInterface[Part_name])
	case *Part_symbol:
		stage.OnAfterPart_symbolUpdateCallback = any(callback).(OnAfterUpdateInterface[Part_symbol])
	case *Part_transpose:
		stage.OnAfterPart_transposeUpdateCallback = any(callback).(OnAfterUpdateInterface[Part_transpose])
	case *Pedal:
		stage.OnAfterPedalUpdateCallback = any(callback).(OnAfterUpdateInterface[Pedal])
	case *Pedal_tuning:
		stage.OnAfterPedal_tuningUpdateCallback = any(callback).(OnAfterUpdateInterface[Pedal_tuning])
	case *Per_minute:
		stage.OnAfterPer_minuteUpdateCallback = any(callback).(OnAfterUpdateInterface[Per_minute])
	case *Percussion:
		stage.OnAfterPercussionUpdateCallback = any(callback).(OnAfterUpdateInterface[Percussion])
	case *Pitch:
		stage.OnAfterPitchUpdateCallback = any(callback).(OnAfterUpdateInterface[Pitch])
	case *Pitched:
		stage.OnAfterPitchedUpdateCallback = any(callback).(OnAfterUpdateInterface[Pitched])
	case *Placement_text:
		stage.OnAfterPlacement_textUpdateCallback = any(callback).(OnAfterUpdateInterface[Placement_text])
	case *Play:
		stage.OnAfterPlayUpdateCallback = any(callback).(OnAfterUpdateInterface[Play])
	case *Player:
		stage.OnAfterPlayerUpdateCallback = any(callback).(OnAfterUpdateInterface[Player])
	case *Principal_voice:
		stage.OnAfterPrincipal_voiceUpdateCallback = any(callback).(OnAfterUpdateInterface[Principal_voice])
	case *Print:
		stage.OnAfterPrintUpdateCallback = any(callback).(OnAfterUpdateInterface[Print])
	case *Release:
		stage.OnAfterReleaseUpdateCallback = any(callback).(OnAfterUpdateInterface[Release])
	case *Repeat:
		stage.OnAfterRepeatUpdateCallback = any(callback).(OnAfterUpdateInterface[Repeat])
	case *Rest:
		stage.OnAfterRestUpdateCallback = any(callback).(OnAfterUpdateInterface[Rest])
	case *Root:
		stage.OnAfterRootUpdateCallback = any(callback).(OnAfterUpdateInterface[Root])
	case *Root_step:
		stage.OnAfterRoot_stepUpdateCallback = any(callback).(OnAfterUpdateInterface[Root_step])
	case *Scaling:
		stage.OnAfterScalingUpdateCallback = any(callback).(OnAfterUpdateInterface[Scaling])
	case *Scordatura:
		stage.OnAfterScordaturaUpdateCallback = any(callback).(OnAfterUpdateInterface[Scordatura])
	case *Score_instrument:
		stage.OnAfterScore_instrumentUpdateCallback = any(callback).(OnAfterUpdateInterface[Score_instrument])
	case *Score_part:
		stage.OnAfterScore_partUpdateCallback = any(callback).(OnAfterUpdateInterface[Score_part])
	case *Score_partwise:
		stage.OnAfterScore_partwiseUpdateCallback = any(callback).(OnAfterUpdateInterface[Score_partwise])
	case *Score_timewise:
		stage.OnAfterScore_timewiseUpdateCallback = any(callback).(OnAfterUpdateInterface[Score_timewise])
	case *Segno:
		stage.OnAfterSegnoUpdateCallback = any(callback).(OnAfterUpdateInterface[Segno])
	case *Slash:
		stage.OnAfterSlashUpdateCallback = any(callback).(OnAfterUpdateInterface[Slash])
	case *Slide:
		stage.OnAfterSlideUpdateCallback = any(callback).(OnAfterUpdateInterface[Slide])
	case *Slur:
		stage.OnAfterSlurUpdateCallback = any(callback).(OnAfterUpdateInterface[Slur])
	case *Sound:
		stage.OnAfterSoundUpdateCallback = any(callback).(OnAfterUpdateInterface[Sound])
	case *Staff_details:
		stage.OnAfterStaff_detailsUpdateCallback = any(callback).(OnAfterUpdateInterface[Staff_details])
	case *Staff_divide:
		stage.OnAfterStaff_divideUpdateCallback = any(callback).(OnAfterUpdateInterface[Staff_divide])
	case *Staff_layout:
		stage.OnAfterStaff_layoutUpdateCallback = any(callback).(OnAfterUpdateInterface[Staff_layout])
	case *Staff_size:
		stage.OnAfterStaff_sizeUpdateCallback = any(callback).(OnAfterUpdateInterface[Staff_size])
	case *Staff_tuning:
		stage.OnAfterStaff_tuningUpdateCallback = any(callback).(OnAfterUpdateInterface[Staff_tuning])
	case *Stem:
		stage.OnAfterStemUpdateCallback = any(callback).(OnAfterUpdateInterface[Stem])
	case *Stick:
		stage.OnAfterStickUpdateCallback = any(callback).(OnAfterUpdateInterface[Stick])
	case *String_mute:
		stage.OnAfterString_muteUpdateCallback = any(callback).(OnAfterUpdateInterface[String_mute])
	case *String_type:
		stage.OnAfterString_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[String_type])
	case *Strong_accent:
		stage.OnAfterStrong_accentUpdateCallback = any(callback).(OnAfterUpdateInterface[Strong_accent])
	case *Style_text:
		stage.OnAfterStyle_textUpdateCallback = any(callback).(OnAfterUpdateInterface[Style_text])
	case *Supports:
		stage.OnAfterSupportsUpdateCallback = any(callback).(OnAfterUpdateInterface[Supports])
	case *Swing:
		stage.OnAfterSwingUpdateCallback = any(callback).(OnAfterUpdateInterface[Swing])
	case *Sync:
		stage.OnAfterSyncUpdateCallback = any(callback).(OnAfterUpdateInterface[Sync])
	case *System_dividers:
		stage.OnAfterSystem_dividersUpdateCallback = any(callback).(OnAfterUpdateInterface[System_dividers])
	case *System_layout:
		stage.OnAfterSystem_layoutUpdateCallback = any(callback).(OnAfterUpdateInterface[System_layout])
	case *System_margins:
		stage.OnAfterSystem_marginsUpdateCallback = any(callback).(OnAfterUpdateInterface[System_margins])
	case *Tap:
		stage.OnAfterTapUpdateCallback = any(callback).(OnAfterUpdateInterface[Tap])
	case *Technical:
		stage.OnAfterTechnicalUpdateCallback = any(callback).(OnAfterUpdateInterface[Technical])
	case *Text_element_data:
		stage.OnAfterText_element_dataUpdateCallback = any(callback).(OnAfterUpdateInterface[Text_element_data])
	case *Tie:
		stage.OnAfterTieUpdateCallback = any(callback).(OnAfterUpdateInterface[Tie])
	case *Tied:
		stage.OnAfterTiedUpdateCallback = any(callback).(OnAfterUpdateInterface[Tied])
	case *Time:
		stage.OnAfterTimeUpdateCallback = any(callback).(OnAfterUpdateInterface[Time])
	case *Time_modification:
		stage.OnAfterTime_modificationUpdateCallback = any(callback).(OnAfterUpdateInterface[Time_modification])
	case *Timpani:
		stage.OnAfterTimpaniUpdateCallback = any(callback).(OnAfterUpdateInterface[Timpani])
	case *Transpose:
		stage.OnAfterTransposeUpdateCallback = any(callback).(OnAfterUpdateInterface[Transpose])
	case *Tremolo:
		stage.OnAfterTremoloUpdateCallback = any(callback).(OnAfterUpdateInterface[Tremolo])
	case *Tuplet:
		stage.OnAfterTupletUpdateCallback = any(callback).(OnAfterUpdateInterface[Tuplet])
	case *Tuplet_dot:
		stage.OnAfterTuplet_dotUpdateCallback = any(callback).(OnAfterUpdateInterface[Tuplet_dot])
	case *Tuplet_number:
		stage.OnAfterTuplet_numberUpdateCallback = any(callback).(OnAfterUpdateInterface[Tuplet_number])
	case *Tuplet_portion:
		stage.OnAfterTuplet_portionUpdateCallback = any(callback).(OnAfterUpdateInterface[Tuplet_portion])
	case *Tuplet_type:
		stage.OnAfterTuplet_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Tuplet_type])
	case *Typed_text:
		stage.OnAfterTyped_textUpdateCallback = any(callback).(OnAfterUpdateInterface[Typed_text])
	case *Unpitched:
		stage.OnAfterUnpitchedUpdateCallback = any(callback).(OnAfterUpdateInterface[Unpitched])
	case *Virtual_instrument:
		stage.OnAfterVirtual_instrumentUpdateCallback = any(callback).(OnAfterUpdateInterface[Virtual_instrument])
	case *Wait:
		stage.OnAfterWaitUpdateCallback = any(callback).(OnAfterUpdateInterface[Wait])
	case *Wavy_line:
		stage.OnAfterWavy_lineUpdateCallback = any(callback).(OnAfterUpdateInterface[Wavy_line])
	case *Wedge:
		stage.OnAfterWedgeUpdateCallback = any(callback).(OnAfterUpdateInterface[Wedge])
	case *Wood:
		stage.OnAfterWoodUpdateCallback = any(callback).(OnAfterUpdateInterface[Wood])
	case *Work:
		stage.OnAfterWorkUpdateCallback = any(callback).(OnAfterUpdateInterface[Work])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *A_directive:
		stage.OnAfterA_directiveCreateCallback = any(callback).(OnAfterCreateInterface[A_directive])
	case *A_measure:
		stage.OnAfterA_measureCreateCallback = any(callback).(OnAfterCreateInterface[A_measure])
	case *A_measure_1:
		stage.OnAfterA_measure_1CreateCallback = any(callback).(OnAfterCreateInterface[A_measure_1])
	case *A_part:
		stage.OnAfterA_partCreateCallback = any(callback).(OnAfterCreateInterface[A_part])
	case *A_part_1:
		stage.OnAfterA_part_1CreateCallback = any(callback).(OnAfterCreateInterface[A_part_1])
	case *Accidental:
		stage.OnAfterAccidentalCreateCallback = any(callback).(OnAfterCreateInterface[Accidental])
	case *Accidental_mark:
		stage.OnAfterAccidental_markCreateCallback = any(callback).(OnAfterCreateInterface[Accidental_mark])
	case *Accidental_text:
		stage.OnAfterAccidental_textCreateCallback = any(callback).(OnAfterCreateInterface[Accidental_text])
	case *Accord:
		stage.OnAfterAccordCreateCallback = any(callback).(OnAfterCreateInterface[Accord])
	case *Accordion_registration:
		stage.OnAfterAccordion_registrationCreateCallback = any(callback).(OnAfterCreateInterface[Accordion_registration])
	case *Appearance:
		stage.OnAfterAppearanceCreateCallback = any(callback).(OnAfterCreateInterface[Appearance])
	case *Arpeggiate:
		stage.OnAfterArpeggiateCreateCallback = any(callback).(OnAfterCreateInterface[Arpeggiate])
	case *Arrow:
		stage.OnAfterArrowCreateCallback = any(callback).(OnAfterCreateInterface[Arrow])
	case *Articulations:
		stage.OnAfterArticulationsCreateCallback = any(callback).(OnAfterCreateInterface[Articulations])
	case *Assess:
		stage.OnAfterAssessCreateCallback = any(callback).(OnAfterCreateInterface[Assess])
	case *Attributes:
		stage.OnAfterAttributesCreateCallback = any(callback).(OnAfterCreateInterface[Attributes])
	case *Backup:
		stage.OnAfterBackupCreateCallback = any(callback).(OnAfterCreateInterface[Backup])
	case *Bar_style_color:
		stage.OnAfterBar_style_colorCreateCallback = any(callback).(OnAfterCreateInterface[Bar_style_color])
	case *Barline:
		stage.OnAfterBarlineCreateCallback = any(callback).(OnAfterCreateInterface[Barline])
	case *Barre:
		stage.OnAfterBarreCreateCallback = any(callback).(OnAfterCreateInterface[Barre])
	case *Bass:
		stage.OnAfterBassCreateCallback = any(callback).(OnAfterCreateInterface[Bass])
	case *Bass_step:
		stage.OnAfterBass_stepCreateCallback = any(callback).(OnAfterCreateInterface[Bass_step])
	case *Beam:
		stage.OnAfterBeamCreateCallback = any(callback).(OnAfterCreateInterface[Beam])
	case *Beat_repeat:
		stage.OnAfterBeat_repeatCreateCallback = any(callback).(OnAfterCreateInterface[Beat_repeat])
	case *Beat_unit_tied:
		stage.OnAfterBeat_unit_tiedCreateCallback = any(callback).(OnAfterCreateInterface[Beat_unit_tied])
	case *Beater:
		stage.OnAfterBeaterCreateCallback = any(callback).(OnAfterCreateInterface[Beater])
	case *Bend:
		stage.OnAfterBendCreateCallback = any(callback).(OnAfterCreateInterface[Bend])
	case *Bookmark:
		stage.OnAfterBookmarkCreateCallback = any(callback).(OnAfterCreateInterface[Bookmark])
	case *Bracket:
		stage.OnAfterBracketCreateCallback = any(callback).(OnAfterCreateInterface[Bracket])
	case *Breath_mark:
		stage.OnAfterBreath_markCreateCallback = any(callback).(OnAfterCreateInterface[Breath_mark])
	case *Caesura:
		stage.OnAfterCaesuraCreateCallback = any(callback).(OnAfterCreateInterface[Caesura])
	case *Cancel:
		stage.OnAfterCancelCreateCallback = any(callback).(OnAfterCreateInterface[Cancel])
	case *Clef:
		stage.OnAfterClefCreateCallback = any(callback).(OnAfterCreateInterface[Clef])
	case *Coda:
		stage.OnAfterCodaCreateCallback = any(callback).(OnAfterCreateInterface[Coda])
	case *Credit:
		stage.OnAfterCreditCreateCallback = any(callback).(OnAfterCreateInterface[Credit])
	case *Dashes:
		stage.OnAfterDashesCreateCallback = any(callback).(OnAfterCreateInterface[Dashes])
	case *Defaults:
		stage.OnAfterDefaultsCreateCallback = any(callback).(OnAfterCreateInterface[Defaults])
	case *Degree:
		stage.OnAfterDegreeCreateCallback = any(callback).(OnAfterCreateInterface[Degree])
	case *Degree_alter:
		stage.OnAfterDegree_alterCreateCallback = any(callback).(OnAfterCreateInterface[Degree_alter])
	case *Degree_type:
		stage.OnAfterDegree_typeCreateCallback = any(callback).(OnAfterCreateInterface[Degree_type])
	case *Degree_value:
		stage.OnAfterDegree_valueCreateCallback = any(callback).(OnAfterCreateInterface[Degree_value])
	case *Direction:
		stage.OnAfterDirectionCreateCallback = any(callback).(OnAfterCreateInterface[Direction])
	case *Direction_type:
		stage.OnAfterDirection_typeCreateCallback = any(callback).(OnAfterCreateInterface[Direction_type])
	case *Distance:
		stage.OnAfterDistanceCreateCallback = any(callback).(OnAfterCreateInterface[Distance])
	case *Double:
		stage.OnAfterDoubleCreateCallback = any(callback).(OnAfterCreateInterface[Double])
	case *Dynamics:
		stage.OnAfterDynamicsCreateCallback = any(callback).(OnAfterCreateInterface[Dynamics])
	case *Effect:
		stage.OnAfterEffectCreateCallback = any(callback).(OnAfterCreateInterface[Effect])
	case *Elision:
		stage.OnAfterElisionCreateCallback = any(callback).(OnAfterCreateInterface[Elision])
	case *Empty:
		stage.OnAfterEmptyCreateCallback = any(callback).(OnAfterCreateInterface[Empty])
	case *Empty_font:
		stage.OnAfterEmpty_fontCreateCallback = any(callback).(OnAfterCreateInterface[Empty_font])
	case *Empty_line:
		stage.OnAfterEmpty_lineCreateCallback = any(callback).(OnAfterCreateInterface[Empty_line])
	case *Empty_placement:
		stage.OnAfterEmpty_placementCreateCallback = any(callback).(OnAfterCreateInterface[Empty_placement])
	case *Empty_placement_smufl:
		stage.OnAfterEmpty_placement_smuflCreateCallback = any(callback).(OnAfterCreateInterface[Empty_placement_smufl])
	case *Empty_print_object_style_align:
		stage.OnAfterEmpty_print_object_style_alignCreateCallback = any(callback).(OnAfterCreateInterface[Empty_print_object_style_align])
	case *Empty_print_style:
		stage.OnAfterEmpty_print_styleCreateCallback = any(callback).(OnAfterCreateInterface[Empty_print_style])
	case *Empty_print_style_align:
		stage.OnAfterEmpty_print_style_alignCreateCallback = any(callback).(OnAfterCreateInterface[Empty_print_style_align])
	case *Empty_print_style_align_id:
		stage.OnAfterEmpty_print_style_align_idCreateCallback = any(callback).(OnAfterCreateInterface[Empty_print_style_align_id])
	case *Empty_trill_sound:
		stage.OnAfterEmpty_trill_soundCreateCallback = any(callback).(OnAfterCreateInterface[Empty_trill_sound])
	case *Encoding:
		stage.OnAfterEncodingCreateCallback = any(callback).(OnAfterCreateInterface[Encoding])
	case *Ending:
		stage.OnAfterEndingCreateCallback = any(callback).(OnAfterCreateInterface[Ending])
	case *Extend:
		stage.OnAfterExtendCreateCallback = any(callback).(OnAfterCreateInterface[Extend])
	case *Feature:
		stage.OnAfterFeatureCreateCallback = any(callback).(OnAfterCreateInterface[Feature])
	case *Fermata:
		stage.OnAfterFermataCreateCallback = any(callback).(OnAfterCreateInterface[Fermata])
	case *Figure:
		stage.OnAfterFigureCreateCallback = any(callback).(OnAfterCreateInterface[Figure])
	case *Figured_bass:
		stage.OnAfterFigured_bassCreateCallback = any(callback).(OnAfterCreateInterface[Figured_bass])
	case *Fingering:
		stage.OnAfterFingeringCreateCallback = any(callback).(OnAfterCreateInterface[Fingering])
	case *First_fret:
		stage.OnAfterFirst_fretCreateCallback = any(callback).(OnAfterCreateInterface[First_fret])
	case *For_part:
		stage.OnAfterFor_partCreateCallback = any(callback).(OnAfterCreateInterface[For_part])
	case *Formatted_symbol:
		stage.OnAfterFormatted_symbolCreateCallback = any(callback).(OnAfterCreateInterface[Formatted_symbol])
	case *Formatted_symbol_id:
		stage.OnAfterFormatted_symbol_idCreateCallback = any(callback).(OnAfterCreateInterface[Formatted_symbol_id])
	case *Formatted_text:
		stage.OnAfterFormatted_textCreateCallback = any(callback).(OnAfterCreateInterface[Formatted_text])
	case *Formatted_text_id:
		stage.OnAfterFormatted_text_idCreateCallback = any(callback).(OnAfterCreateInterface[Formatted_text_id])
	case *Forward:
		stage.OnAfterForwardCreateCallback = any(callback).(OnAfterCreateInterface[Forward])
	case *Frame:
		stage.OnAfterFrameCreateCallback = any(callback).(OnAfterCreateInterface[Frame])
	case *Frame_note:
		stage.OnAfterFrame_noteCreateCallback = any(callback).(OnAfterCreateInterface[Frame_note])
	case *Fret:
		stage.OnAfterFretCreateCallback = any(callback).(OnAfterCreateInterface[Fret])
	case *Glass:
		stage.OnAfterGlassCreateCallback = any(callback).(OnAfterCreateInterface[Glass])
	case *Glissando:
		stage.OnAfterGlissandoCreateCallback = any(callback).(OnAfterCreateInterface[Glissando])
	case *Glyph:
		stage.OnAfterGlyphCreateCallback = any(callback).(OnAfterCreateInterface[Glyph])
	case *Grace:
		stage.OnAfterGraceCreateCallback = any(callback).(OnAfterCreateInterface[Grace])
	case *Group_barline:
		stage.OnAfterGroup_barlineCreateCallback = any(callback).(OnAfterCreateInterface[Group_barline])
	case *Group_name:
		stage.OnAfterGroup_nameCreateCallback = any(callback).(OnAfterCreateInterface[Group_name])
	case *Group_symbol:
		stage.OnAfterGroup_symbolCreateCallback = any(callback).(OnAfterCreateInterface[Group_symbol])
	case *Grouping:
		stage.OnAfterGroupingCreateCallback = any(callback).(OnAfterCreateInterface[Grouping])
	case *Hammer_on_pull_off:
		stage.OnAfterHammer_on_pull_offCreateCallback = any(callback).(OnAfterCreateInterface[Hammer_on_pull_off])
	case *Handbell:
		stage.OnAfterHandbellCreateCallback = any(callback).(OnAfterCreateInterface[Handbell])
	case *Harmon_closed:
		stage.OnAfterHarmon_closedCreateCallback = any(callback).(OnAfterCreateInterface[Harmon_closed])
	case *Harmon_mute:
		stage.OnAfterHarmon_muteCreateCallback = any(callback).(OnAfterCreateInterface[Harmon_mute])
	case *Harmonic:
		stage.OnAfterHarmonicCreateCallback = any(callback).(OnAfterCreateInterface[Harmonic])
	case *Harmony:
		stage.OnAfterHarmonyCreateCallback = any(callback).(OnAfterCreateInterface[Harmony])
	case *Harmony_alter:
		stage.OnAfterHarmony_alterCreateCallback = any(callback).(OnAfterCreateInterface[Harmony_alter])
	case *Harp_pedals:
		stage.OnAfterHarp_pedalsCreateCallback = any(callback).(OnAfterCreateInterface[Harp_pedals])
	case *Heel_toe:
		stage.OnAfterHeel_toeCreateCallback = any(callback).(OnAfterCreateInterface[Heel_toe])
	case *Hole:
		stage.OnAfterHoleCreateCallback = any(callback).(OnAfterCreateInterface[Hole])
	case *Hole_closed:
		stage.OnAfterHole_closedCreateCallback = any(callback).(OnAfterCreateInterface[Hole_closed])
	case *Horizontal_turn:
		stage.OnAfterHorizontal_turnCreateCallback = any(callback).(OnAfterCreateInterface[Horizontal_turn])
	case *Identification:
		stage.OnAfterIdentificationCreateCallback = any(callback).(OnAfterCreateInterface[Identification])
	case *Image:
		stage.OnAfterImageCreateCallback = any(callback).(OnAfterCreateInterface[Image])
	case *Instrument:
		stage.OnAfterInstrumentCreateCallback = any(callback).(OnAfterCreateInterface[Instrument])
	case *Instrument_change:
		stage.OnAfterInstrument_changeCreateCallback = any(callback).(OnAfterCreateInterface[Instrument_change])
	case *Instrument_link:
		stage.OnAfterInstrument_linkCreateCallback = any(callback).(OnAfterCreateInterface[Instrument_link])
	case *Interchangeable:
		stage.OnAfterInterchangeableCreateCallback = any(callback).(OnAfterCreateInterface[Interchangeable])
	case *Inversion:
		stage.OnAfterInversionCreateCallback = any(callback).(OnAfterCreateInterface[Inversion])
	case *Key:
		stage.OnAfterKeyCreateCallback = any(callback).(OnAfterCreateInterface[Key])
	case *Key_accidental:
		stage.OnAfterKey_accidentalCreateCallback = any(callback).(OnAfterCreateInterface[Key_accidental])
	case *Key_octave:
		stage.OnAfterKey_octaveCreateCallback = any(callback).(OnAfterCreateInterface[Key_octave])
	case *Kind:
		stage.OnAfterKindCreateCallback = any(callback).(OnAfterCreateInterface[Kind])
	case *Level:
		stage.OnAfterLevelCreateCallback = any(callback).(OnAfterCreateInterface[Level])
	case *Line_detail:
		stage.OnAfterLine_detailCreateCallback = any(callback).(OnAfterCreateInterface[Line_detail])
	case *Line_width:
		stage.OnAfterLine_widthCreateCallback = any(callback).(OnAfterCreateInterface[Line_width])
	case *Link:
		stage.OnAfterLinkCreateCallback = any(callback).(OnAfterCreateInterface[Link])
	case *Listen:
		stage.OnAfterListenCreateCallback = any(callback).(OnAfterCreateInterface[Listen])
	case *Listening:
		stage.OnAfterListeningCreateCallback = any(callback).(OnAfterCreateInterface[Listening])
	case *Lyric:
		stage.OnAfterLyricCreateCallback = any(callback).(OnAfterCreateInterface[Lyric])
	case *Lyric_font:
		stage.OnAfterLyric_fontCreateCallback = any(callback).(OnAfterCreateInterface[Lyric_font])
	case *Lyric_language:
		stage.OnAfterLyric_languageCreateCallback = any(callback).(OnAfterCreateInterface[Lyric_language])
	case *Measure_layout:
		stage.OnAfterMeasure_layoutCreateCallback = any(callback).(OnAfterCreateInterface[Measure_layout])
	case *Measure_numbering:
		stage.OnAfterMeasure_numberingCreateCallback = any(callback).(OnAfterCreateInterface[Measure_numbering])
	case *Measure_repeat:
		stage.OnAfterMeasure_repeatCreateCallback = any(callback).(OnAfterCreateInterface[Measure_repeat])
	case *Measure_style:
		stage.OnAfterMeasure_styleCreateCallback = any(callback).(OnAfterCreateInterface[Measure_style])
	case *Membrane:
		stage.OnAfterMembraneCreateCallback = any(callback).(OnAfterCreateInterface[Membrane])
	case *Metal:
		stage.OnAfterMetalCreateCallback = any(callback).(OnAfterCreateInterface[Metal])
	case *Metronome:
		stage.OnAfterMetronomeCreateCallback = any(callback).(OnAfterCreateInterface[Metronome])
	case *Metronome_beam:
		stage.OnAfterMetronome_beamCreateCallback = any(callback).(OnAfterCreateInterface[Metronome_beam])
	case *Metronome_note:
		stage.OnAfterMetronome_noteCreateCallback = any(callback).(OnAfterCreateInterface[Metronome_note])
	case *Metronome_tied:
		stage.OnAfterMetronome_tiedCreateCallback = any(callback).(OnAfterCreateInterface[Metronome_tied])
	case *Metronome_tuplet:
		stage.OnAfterMetronome_tupletCreateCallback = any(callback).(OnAfterCreateInterface[Metronome_tuplet])
	case *Midi_device:
		stage.OnAfterMidi_deviceCreateCallback = any(callback).(OnAfterCreateInterface[Midi_device])
	case *Midi_instrument:
		stage.OnAfterMidi_instrumentCreateCallback = any(callback).(OnAfterCreateInterface[Midi_instrument])
	case *Miscellaneous:
		stage.OnAfterMiscellaneousCreateCallback = any(callback).(OnAfterCreateInterface[Miscellaneous])
	case *Miscellaneous_field:
		stage.OnAfterMiscellaneous_fieldCreateCallback = any(callback).(OnAfterCreateInterface[Miscellaneous_field])
	case *Mordent:
		stage.OnAfterMordentCreateCallback = any(callback).(OnAfterCreateInterface[Mordent])
	case *Multiple_rest:
		stage.OnAfterMultiple_restCreateCallback = any(callback).(OnAfterCreateInterface[Multiple_rest])
	case *Name_display:
		stage.OnAfterName_displayCreateCallback = any(callback).(OnAfterCreateInterface[Name_display])
	case *Non_arpeggiate:
		stage.OnAfterNon_arpeggiateCreateCallback = any(callback).(OnAfterCreateInterface[Non_arpeggiate])
	case *Notations:
		stage.OnAfterNotationsCreateCallback = any(callback).(OnAfterCreateInterface[Notations])
	case *Note:
		stage.OnAfterNoteCreateCallback = any(callback).(OnAfterCreateInterface[Note])
	case *Note_size:
		stage.OnAfterNote_sizeCreateCallback = any(callback).(OnAfterCreateInterface[Note_size])
	case *Note_type:
		stage.OnAfterNote_typeCreateCallback = any(callback).(OnAfterCreateInterface[Note_type])
	case *Notehead:
		stage.OnAfterNoteheadCreateCallback = any(callback).(OnAfterCreateInterface[Notehead])
	case *Notehead_text:
		stage.OnAfterNotehead_textCreateCallback = any(callback).(OnAfterCreateInterface[Notehead_text])
	case *Numeral:
		stage.OnAfterNumeralCreateCallback = any(callback).(OnAfterCreateInterface[Numeral])
	case *Numeral_key:
		stage.OnAfterNumeral_keyCreateCallback = any(callback).(OnAfterCreateInterface[Numeral_key])
	case *Numeral_root:
		stage.OnAfterNumeral_rootCreateCallback = any(callback).(OnAfterCreateInterface[Numeral_root])
	case *Octave_shift:
		stage.OnAfterOctave_shiftCreateCallback = any(callback).(OnAfterCreateInterface[Octave_shift])
	case *Offset:
		stage.OnAfterOffsetCreateCallback = any(callback).(OnAfterCreateInterface[Offset])
	case *Opus:
		stage.OnAfterOpusCreateCallback = any(callback).(OnAfterCreateInterface[Opus])
	case *Ornaments:
		stage.OnAfterOrnamentsCreateCallback = any(callback).(OnAfterCreateInterface[Ornaments])
	case *Other_appearance:
		stage.OnAfterOther_appearanceCreateCallback = any(callback).(OnAfterCreateInterface[Other_appearance])
	case *Other_direction:
		stage.OnAfterOther_directionCreateCallback = any(callback).(OnAfterCreateInterface[Other_direction])
	case *Other_listening:
		stage.OnAfterOther_listeningCreateCallback = any(callback).(OnAfterCreateInterface[Other_listening])
	case *Other_notation:
		stage.OnAfterOther_notationCreateCallback = any(callback).(OnAfterCreateInterface[Other_notation])
	case *Other_placement_text:
		stage.OnAfterOther_placement_textCreateCallback = any(callback).(OnAfterCreateInterface[Other_placement_text])
	case *Other_play:
		stage.OnAfterOther_playCreateCallback = any(callback).(OnAfterCreateInterface[Other_play])
	case *Other_text:
		stage.OnAfterOther_textCreateCallback = any(callback).(OnAfterCreateInterface[Other_text])
	case *Page_layout:
		stage.OnAfterPage_layoutCreateCallback = any(callback).(OnAfterCreateInterface[Page_layout])
	case *Page_margins:
		stage.OnAfterPage_marginsCreateCallback = any(callback).(OnAfterCreateInterface[Page_margins])
	case *Part_clef:
		stage.OnAfterPart_clefCreateCallback = any(callback).(OnAfterCreateInterface[Part_clef])
	case *Part_group:
		stage.OnAfterPart_groupCreateCallback = any(callback).(OnAfterCreateInterface[Part_group])
	case *Part_link:
		stage.OnAfterPart_linkCreateCallback = any(callback).(OnAfterCreateInterface[Part_link])
	case *Part_list:
		stage.OnAfterPart_listCreateCallback = any(callback).(OnAfterCreateInterface[Part_list])
	case *Part_name:
		stage.OnAfterPart_nameCreateCallback = any(callback).(OnAfterCreateInterface[Part_name])
	case *Part_symbol:
		stage.OnAfterPart_symbolCreateCallback = any(callback).(OnAfterCreateInterface[Part_symbol])
	case *Part_transpose:
		stage.OnAfterPart_transposeCreateCallback = any(callback).(OnAfterCreateInterface[Part_transpose])
	case *Pedal:
		stage.OnAfterPedalCreateCallback = any(callback).(OnAfterCreateInterface[Pedal])
	case *Pedal_tuning:
		stage.OnAfterPedal_tuningCreateCallback = any(callback).(OnAfterCreateInterface[Pedal_tuning])
	case *Per_minute:
		stage.OnAfterPer_minuteCreateCallback = any(callback).(OnAfterCreateInterface[Per_minute])
	case *Percussion:
		stage.OnAfterPercussionCreateCallback = any(callback).(OnAfterCreateInterface[Percussion])
	case *Pitch:
		stage.OnAfterPitchCreateCallback = any(callback).(OnAfterCreateInterface[Pitch])
	case *Pitched:
		stage.OnAfterPitchedCreateCallback = any(callback).(OnAfterCreateInterface[Pitched])
	case *Placement_text:
		stage.OnAfterPlacement_textCreateCallback = any(callback).(OnAfterCreateInterface[Placement_text])
	case *Play:
		stage.OnAfterPlayCreateCallback = any(callback).(OnAfterCreateInterface[Play])
	case *Player:
		stage.OnAfterPlayerCreateCallback = any(callback).(OnAfterCreateInterface[Player])
	case *Principal_voice:
		stage.OnAfterPrincipal_voiceCreateCallback = any(callback).(OnAfterCreateInterface[Principal_voice])
	case *Print:
		stage.OnAfterPrintCreateCallback = any(callback).(OnAfterCreateInterface[Print])
	case *Release:
		stage.OnAfterReleaseCreateCallback = any(callback).(OnAfterCreateInterface[Release])
	case *Repeat:
		stage.OnAfterRepeatCreateCallback = any(callback).(OnAfterCreateInterface[Repeat])
	case *Rest:
		stage.OnAfterRestCreateCallback = any(callback).(OnAfterCreateInterface[Rest])
	case *Root:
		stage.OnAfterRootCreateCallback = any(callback).(OnAfterCreateInterface[Root])
	case *Root_step:
		stage.OnAfterRoot_stepCreateCallback = any(callback).(OnAfterCreateInterface[Root_step])
	case *Scaling:
		stage.OnAfterScalingCreateCallback = any(callback).(OnAfterCreateInterface[Scaling])
	case *Scordatura:
		stage.OnAfterScordaturaCreateCallback = any(callback).(OnAfterCreateInterface[Scordatura])
	case *Score_instrument:
		stage.OnAfterScore_instrumentCreateCallback = any(callback).(OnAfterCreateInterface[Score_instrument])
	case *Score_part:
		stage.OnAfterScore_partCreateCallback = any(callback).(OnAfterCreateInterface[Score_part])
	case *Score_partwise:
		stage.OnAfterScore_partwiseCreateCallback = any(callback).(OnAfterCreateInterface[Score_partwise])
	case *Score_timewise:
		stage.OnAfterScore_timewiseCreateCallback = any(callback).(OnAfterCreateInterface[Score_timewise])
	case *Segno:
		stage.OnAfterSegnoCreateCallback = any(callback).(OnAfterCreateInterface[Segno])
	case *Slash:
		stage.OnAfterSlashCreateCallback = any(callback).(OnAfterCreateInterface[Slash])
	case *Slide:
		stage.OnAfterSlideCreateCallback = any(callback).(OnAfterCreateInterface[Slide])
	case *Slur:
		stage.OnAfterSlurCreateCallback = any(callback).(OnAfterCreateInterface[Slur])
	case *Sound:
		stage.OnAfterSoundCreateCallback = any(callback).(OnAfterCreateInterface[Sound])
	case *Staff_details:
		stage.OnAfterStaff_detailsCreateCallback = any(callback).(OnAfterCreateInterface[Staff_details])
	case *Staff_divide:
		stage.OnAfterStaff_divideCreateCallback = any(callback).(OnAfterCreateInterface[Staff_divide])
	case *Staff_layout:
		stage.OnAfterStaff_layoutCreateCallback = any(callback).(OnAfterCreateInterface[Staff_layout])
	case *Staff_size:
		stage.OnAfterStaff_sizeCreateCallback = any(callback).(OnAfterCreateInterface[Staff_size])
	case *Staff_tuning:
		stage.OnAfterStaff_tuningCreateCallback = any(callback).(OnAfterCreateInterface[Staff_tuning])
	case *Stem:
		stage.OnAfterStemCreateCallback = any(callback).(OnAfterCreateInterface[Stem])
	case *Stick:
		stage.OnAfterStickCreateCallback = any(callback).(OnAfterCreateInterface[Stick])
	case *String_mute:
		stage.OnAfterString_muteCreateCallback = any(callback).(OnAfterCreateInterface[String_mute])
	case *String_type:
		stage.OnAfterString_typeCreateCallback = any(callback).(OnAfterCreateInterface[String_type])
	case *Strong_accent:
		stage.OnAfterStrong_accentCreateCallback = any(callback).(OnAfterCreateInterface[Strong_accent])
	case *Style_text:
		stage.OnAfterStyle_textCreateCallback = any(callback).(OnAfterCreateInterface[Style_text])
	case *Supports:
		stage.OnAfterSupportsCreateCallback = any(callback).(OnAfterCreateInterface[Supports])
	case *Swing:
		stage.OnAfterSwingCreateCallback = any(callback).(OnAfterCreateInterface[Swing])
	case *Sync:
		stage.OnAfterSyncCreateCallback = any(callback).(OnAfterCreateInterface[Sync])
	case *System_dividers:
		stage.OnAfterSystem_dividersCreateCallback = any(callback).(OnAfterCreateInterface[System_dividers])
	case *System_layout:
		stage.OnAfterSystem_layoutCreateCallback = any(callback).(OnAfterCreateInterface[System_layout])
	case *System_margins:
		stage.OnAfterSystem_marginsCreateCallback = any(callback).(OnAfterCreateInterface[System_margins])
	case *Tap:
		stage.OnAfterTapCreateCallback = any(callback).(OnAfterCreateInterface[Tap])
	case *Technical:
		stage.OnAfterTechnicalCreateCallback = any(callback).(OnAfterCreateInterface[Technical])
	case *Text_element_data:
		stage.OnAfterText_element_dataCreateCallback = any(callback).(OnAfterCreateInterface[Text_element_data])
	case *Tie:
		stage.OnAfterTieCreateCallback = any(callback).(OnAfterCreateInterface[Tie])
	case *Tied:
		stage.OnAfterTiedCreateCallback = any(callback).(OnAfterCreateInterface[Tied])
	case *Time:
		stage.OnAfterTimeCreateCallback = any(callback).(OnAfterCreateInterface[Time])
	case *Time_modification:
		stage.OnAfterTime_modificationCreateCallback = any(callback).(OnAfterCreateInterface[Time_modification])
	case *Timpani:
		stage.OnAfterTimpaniCreateCallback = any(callback).(OnAfterCreateInterface[Timpani])
	case *Transpose:
		stage.OnAfterTransposeCreateCallback = any(callback).(OnAfterCreateInterface[Transpose])
	case *Tremolo:
		stage.OnAfterTremoloCreateCallback = any(callback).(OnAfterCreateInterface[Tremolo])
	case *Tuplet:
		stage.OnAfterTupletCreateCallback = any(callback).(OnAfterCreateInterface[Tuplet])
	case *Tuplet_dot:
		stage.OnAfterTuplet_dotCreateCallback = any(callback).(OnAfterCreateInterface[Tuplet_dot])
	case *Tuplet_number:
		stage.OnAfterTuplet_numberCreateCallback = any(callback).(OnAfterCreateInterface[Tuplet_number])
	case *Tuplet_portion:
		stage.OnAfterTuplet_portionCreateCallback = any(callback).(OnAfterCreateInterface[Tuplet_portion])
	case *Tuplet_type:
		stage.OnAfterTuplet_typeCreateCallback = any(callback).(OnAfterCreateInterface[Tuplet_type])
	case *Typed_text:
		stage.OnAfterTyped_textCreateCallback = any(callback).(OnAfterCreateInterface[Typed_text])
	case *Unpitched:
		stage.OnAfterUnpitchedCreateCallback = any(callback).(OnAfterCreateInterface[Unpitched])
	case *Virtual_instrument:
		stage.OnAfterVirtual_instrumentCreateCallback = any(callback).(OnAfterCreateInterface[Virtual_instrument])
	case *Wait:
		stage.OnAfterWaitCreateCallback = any(callback).(OnAfterCreateInterface[Wait])
	case *Wavy_line:
		stage.OnAfterWavy_lineCreateCallback = any(callback).(OnAfterCreateInterface[Wavy_line])
	case *Wedge:
		stage.OnAfterWedgeCreateCallback = any(callback).(OnAfterCreateInterface[Wedge])
	case *Wood:
		stage.OnAfterWoodCreateCallback = any(callback).(OnAfterCreateInterface[Wood])
	case *Work:
		stage.OnAfterWorkCreateCallback = any(callback).(OnAfterCreateInterface[Work])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *A_directive:
		stage.OnAfterA_directiveDeleteCallback = any(callback).(OnAfterDeleteInterface[A_directive])
	case *A_measure:
		stage.OnAfterA_measureDeleteCallback = any(callback).(OnAfterDeleteInterface[A_measure])
	case *A_measure_1:
		stage.OnAfterA_measure_1DeleteCallback = any(callback).(OnAfterDeleteInterface[A_measure_1])
	case *A_part:
		stage.OnAfterA_partDeleteCallback = any(callback).(OnAfterDeleteInterface[A_part])
	case *A_part_1:
		stage.OnAfterA_part_1DeleteCallback = any(callback).(OnAfterDeleteInterface[A_part_1])
	case *Accidental:
		stage.OnAfterAccidentalDeleteCallback = any(callback).(OnAfterDeleteInterface[Accidental])
	case *Accidental_mark:
		stage.OnAfterAccidental_markDeleteCallback = any(callback).(OnAfterDeleteInterface[Accidental_mark])
	case *Accidental_text:
		stage.OnAfterAccidental_textDeleteCallback = any(callback).(OnAfterDeleteInterface[Accidental_text])
	case *Accord:
		stage.OnAfterAccordDeleteCallback = any(callback).(OnAfterDeleteInterface[Accord])
	case *Accordion_registration:
		stage.OnAfterAccordion_registrationDeleteCallback = any(callback).(OnAfterDeleteInterface[Accordion_registration])
	case *Appearance:
		stage.OnAfterAppearanceDeleteCallback = any(callback).(OnAfterDeleteInterface[Appearance])
	case *Arpeggiate:
		stage.OnAfterArpeggiateDeleteCallback = any(callback).(OnAfterDeleteInterface[Arpeggiate])
	case *Arrow:
		stage.OnAfterArrowDeleteCallback = any(callback).(OnAfterDeleteInterface[Arrow])
	case *Articulations:
		stage.OnAfterArticulationsDeleteCallback = any(callback).(OnAfterDeleteInterface[Articulations])
	case *Assess:
		stage.OnAfterAssessDeleteCallback = any(callback).(OnAfterDeleteInterface[Assess])
	case *Attributes:
		stage.OnAfterAttributesDeleteCallback = any(callback).(OnAfterDeleteInterface[Attributes])
	case *Backup:
		stage.OnAfterBackupDeleteCallback = any(callback).(OnAfterDeleteInterface[Backup])
	case *Bar_style_color:
		stage.OnAfterBar_style_colorDeleteCallback = any(callback).(OnAfterDeleteInterface[Bar_style_color])
	case *Barline:
		stage.OnAfterBarlineDeleteCallback = any(callback).(OnAfterDeleteInterface[Barline])
	case *Barre:
		stage.OnAfterBarreDeleteCallback = any(callback).(OnAfterDeleteInterface[Barre])
	case *Bass:
		stage.OnAfterBassDeleteCallback = any(callback).(OnAfterDeleteInterface[Bass])
	case *Bass_step:
		stage.OnAfterBass_stepDeleteCallback = any(callback).(OnAfterDeleteInterface[Bass_step])
	case *Beam:
		stage.OnAfterBeamDeleteCallback = any(callback).(OnAfterDeleteInterface[Beam])
	case *Beat_repeat:
		stage.OnAfterBeat_repeatDeleteCallback = any(callback).(OnAfterDeleteInterface[Beat_repeat])
	case *Beat_unit_tied:
		stage.OnAfterBeat_unit_tiedDeleteCallback = any(callback).(OnAfterDeleteInterface[Beat_unit_tied])
	case *Beater:
		stage.OnAfterBeaterDeleteCallback = any(callback).(OnAfterDeleteInterface[Beater])
	case *Bend:
		stage.OnAfterBendDeleteCallback = any(callback).(OnAfterDeleteInterface[Bend])
	case *Bookmark:
		stage.OnAfterBookmarkDeleteCallback = any(callback).(OnAfterDeleteInterface[Bookmark])
	case *Bracket:
		stage.OnAfterBracketDeleteCallback = any(callback).(OnAfterDeleteInterface[Bracket])
	case *Breath_mark:
		stage.OnAfterBreath_markDeleteCallback = any(callback).(OnAfterDeleteInterface[Breath_mark])
	case *Caesura:
		stage.OnAfterCaesuraDeleteCallback = any(callback).(OnAfterDeleteInterface[Caesura])
	case *Cancel:
		stage.OnAfterCancelDeleteCallback = any(callback).(OnAfterDeleteInterface[Cancel])
	case *Clef:
		stage.OnAfterClefDeleteCallback = any(callback).(OnAfterDeleteInterface[Clef])
	case *Coda:
		stage.OnAfterCodaDeleteCallback = any(callback).(OnAfterDeleteInterface[Coda])
	case *Credit:
		stage.OnAfterCreditDeleteCallback = any(callback).(OnAfterDeleteInterface[Credit])
	case *Dashes:
		stage.OnAfterDashesDeleteCallback = any(callback).(OnAfterDeleteInterface[Dashes])
	case *Defaults:
		stage.OnAfterDefaultsDeleteCallback = any(callback).(OnAfterDeleteInterface[Defaults])
	case *Degree:
		stage.OnAfterDegreeDeleteCallback = any(callback).(OnAfterDeleteInterface[Degree])
	case *Degree_alter:
		stage.OnAfterDegree_alterDeleteCallback = any(callback).(OnAfterDeleteInterface[Degree_alter])
	case *Degree_type:
		stage.OnAfterDegree_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Degree_type])
	case *Degree_value:
		stage.OnAfterDegree_valueDeleteCallback = any(callback).(OnAfterDeleteInterface[Degree_value])
	case *Direction:
		stage.OnAfterDirectionDeleteCallback = any(callback).(OnAfterDeleteInterface[Direction])
	case *Direction_type:
		stage.OnAfterDirection_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Direction_type])
	case *Distance:
		stage.OnAfterDistanceDeleteCallback = any(callback).(OnAfterDeleteInterface[Distance])
	case *Double:
		stage.OnAfterDoubleDeleteCallback = any(callback).(OnAfterDeleteInterface[Double])
	case *Dynamics:
		stage.OnAfterDynamicsDeleteCallback = any(callback).(OnAfterDeleteInterface[Dynamics])
	case *Effect:
		stage.OnAfterEffectDeleteCallback = any(callback).(OnAfterDeleteInterface[Effect])
	case *Elision:
		stage.OnAfterElisionDeleteCallback = any(callback).(OnAfterDeleteInterface[Elision])
	case *Empty:
		stage.OnAfterEmptyDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty])
	case *Empty_font:
		stage.OnAfterEmpty_fontDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty_font])
	case *Empty_line:
		stage.OnAfterEmpty_lineDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty_line])
	case *Empty_placement:
		stage.OnAfterEmpty_placementDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty_placement])
	case *Empty_placement_smufl:
		stage.OnAfterEmpty_placement_smuflDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty_placement_smufl])
	case *Empty_print_object_style_align:
		stage.OnAfterEmpty_print_object_style_alignDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty_print_object_style_align])
	case *Empty_print_style:
		stage.OnAfterEmpty_print_styleDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty_print_style])
	case *Empty_print_style_align:
		stage.OnAfterEmpty_print_style_alignDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty_print_style_align])
	case *Empty_print_style_align_id:
		stage.OnAfterEmpty_print_style_align_idDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty_print_style_align_id])
	case *Empty_trill_sound:
		stage.OnAfterEmpty_trill_soundDeleteCallback = any(callback).(OnAfterDeleteInterface[Empty_trill_sound])
	case *Encoding:
		stage.OnAfterEncodingDeleteCallback = any(callback).(OnAfterDeleteInterface[Encoding])
	case *Ending:
		stage.OnAfterEndingDeleteCallback = any(callback).(OnAfterDeleteInterface[Ending])
	case *Extend:
		stage.OnAfterExtendDeleteCallback = any(callback).(OnAfterDeleteInterface[Extend])
	case *Feature:
		stage.OnAfterFeatureDeleteCallback = any(callback).(OnAfterDeleteInterface[Feature])
	case *Fermata:
		stage.OnAfterFermataDeleteCallback = any(callback).(OnAfterDeleteInterface[Fermata])
	case *Figure:
		stage.OnAfterFigureDeleteCallback = any(callback).(OnAfterDeleteInterface[Figure])
	case *Figured_bass:
		stage.OnAfterFigured_bassDeleteCallback = any(callback).(OnAfterDeleteInterface[Figured_bass])
	case *Fingering:
		stage.OnAfterFingeringDeleteCallback = any(callback).(OnAfterDeleteInterface[Fingering])
	case *First_fret:
		stage.OnAfterFirst_fretDeleteCallback = any(callback).(OnAfterDeleteInterface[First_fret])
	case *For_part:
		stage.OnAfterFor_partDeleteCallback = any(callback).(OnAfterDeleteInterface[For_part])
	case *Formatted_symbol:
		stage.OnAfterFormatted_symbolDeleteCallback = any(callback).(OnAfterDeleteInterface[Formatted_symbol])
	case *Formatted_symbol_id:
		stage.OnAfterFormatted_symbol_idDeleteCallback = any(callback).(OnAfterDeleteInterface[Formatted_symbol_id])
	case *Formatted_text:
		stage.OnAfterFormatted_textDeleteCallback = any(callback).(OnAfterDeleteInterface[Formatted_text])
	case *Formatted_text_id:
		stage.OnAfterFormatted_text_idDeleteCallback = any(callback).(OnAfterDeleteInterface[Formatted_text_id])
	case *Forward:
		stage.OnAfterForwardDeleteCallback = any(callback).(OnAfterDeleteInterface[Forward])
	case *Frame:
		stage.OnAfterFrameDeleteCallback = any(callback).(OnAfterDeleteInterface[Frame])
	case *Frame_note:
		stage.OnAfterFrame_noteDeleteCallback = any(callback).(OnAfterDeleteInterface[Frame_note])
	case *Fret:
		stage.OnAfterFretDeleteCallback = any(callback).(OnAfterDeleteInterface[Fret])
	case *Glass:
		stage.OnAfterGlassDeleteCallback = any(callback).(OnAfterDeleteInterface[Glass])
	case *Glissando:
		stage.OnAfterGlissandoDeleteCallback = any(callback).(OnAfterDeleteInterface[Glissando])
	case *Glyph:
		stage.OnAfterGlyphDeleteCallback = any(callback).(OnAfterDeleteInterface[Glyph])
	case *Grace:
		stage.OnAfterGraceDeleteCallback = any(callback).(OnAfterDeleteInterface[Grace])
	case *Group_barline:
		stage.OnAfterGroup_barlineDeleteCallback = any(callback).(OnAfterDeleteInterface[Group_barline])
	case *Group_name:
		stage.OnAfterGroup_nameDeleteCallback = any(callback).(OnAfterDeleteInterface[Group_name])
	case *Group_symbol:
		stage.OnAfterGroup_symbolDeleteCallback = any(callback).(OnAfterDeleteInterface[Group_symbol])
	case *Grouping:
		stage.OnAfterGroupingDeleteCallback = any(callback).(OnAfterDeleteInterface[Grouping])
	case *Hammer_on_pull_off:
		stage.OnAfterHammer_on_pull_offDeleteCallback = any(callback).(OnAfterDeleteInterface[Hammer_on_pull_off])
	case *Handbell:
		stage.OnAfterHandbellDeleteCallback = any(callback).(OnAfterDeleteInterface[Handbell])
	case *Harmon_closed:
		stage.OnAfterHarmon_closedDeleteCallback = any(callback).(OnAfterDeleteInterface[Harmon_closed])
	case *Harmon_mute:
		stage.OnAfterHarmon_muteDeleteCallback = any(callback).(OnAfterDeleteInterface[Harmon_mute])
	case *Harmonic:
		stage.OnAfterHarmonicDeleteCallback = any(callback).(OnAfterDeleteInterface[Harmonic])
	case *Harmony:
		stage.OnAfterHarmonyDeleteCallback = any(callback).(OnAfterDeleteInterface[Harmony])
	case *Harmony_alter:
		stage.OnAfterHarmony_alterDeleteCallback = any(callback).(OnAfterDeleteInterface[Harmony_alter])
	case *Harp_pedals:
		stage.OnAfterHarp_pedalsDeleteCallback = any(callback).(OnAfterDeleteInterface[Harp_pedals])
	case *Heel_toe:
		stage.OnAfterHeel_toeDeleteCallback = any(callback).(OnAfterDeleteInterface[Heel_toe])
	case *Hole:
		stage.OnAfterHoleDeleteCallback = any(callback).(OnAfterDeleteInterface[Hole])
	case *Hole_closed:
		stage.OnAfterHole_closedDeleteCallback = any(callback).(OnAfterDeleteInterface[Hole_closed])
	case *Horizontal_turn:
		stage.OnAfterHorizontal_turnDeleteCallback = any(callback).(OnAfterDeleteInterface[Horizontal_turn])
	case *Identification:
		stage.OnAfterIdentificationDeleteCallback = any(callback).(OnAfterDeleteInterface[Identification])
	case *Image:
		stage.OnAfterImageDeleteCallback = any(callback).(OnAfterDeleteInterface[Image])
	case *Instrument:
		stage.OnAfterInstrumentDeleteCallback = any(callback).(OnAfterDeleteInterface[Instrument])
	case *Instrument_change:
		stage.OnAfterInstrument_changeDeleteCallback = any(callback).(OnAfterDeleteInterface[Instrument_change])
	case *Instrument_link:
		stage.OnAfterInstrument_linkDeleteCallback = any(callback).(OnAfterDeleteInterface[Instrument_link])
	case *Interchangeable:
		stage.OnAfterInterchangeableDeleteCallback = any(callback).(OnAfterDeleteInterface[Interchangeable])
	case *Inversion:
		stage.OnAfterInversionDeleteCallback = any(callback).(OnAfterDeleteInterface[Inversion])
	case *Key:
		stage.OnAfterKeyDeleteCallback = any(callback).(OnAfterDeleteInterface[Key])
	case *Key_accidental:
		stage.OnAfterKey_accidentalDeleteCallback = any(callback).(OnAfterDeleteInterface[Key_accidental])
	case *Key_octave:
		stage.OnAfterKey_octaveDeleteCallback = any(callback).(OnAfterDeleteInterface[Key_octave])
	case *Kind:
		stage.OnAfterKindDeleteCallback = any(callback).(OnAfterDeleteInterface[Kind])
	case *Level:
		stage.OnAfterLevelDeleteCallback = any(callback).(OnAfterDeleteInterface[Level])
	case *Line_detail:
		stage.OnAfterLine_detailDeleteCallback = any(callback).(OnAfterDeleteInterface[Line_detail])
	case *Line_width:
		stage.OnAfterLine_widthDeleteCallback = any(callback).(OnAfterDeleteInterface[Line_width])
	case *Link:
		stage.OnAfterLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[Link])
	case *Listen:
		stage.OnAfterListenDeleteCallback = any(callback).(OnAfterDeleteInterface[Listen])
	case *Listening:
		stage.OnAfterListeningDeleteCallback = any(callback).(OnAfterDeleteInterface[Listening])
	case *Lyric:
		stage.OnAfterLyricDeleteCallback = any(callback).(OnAfterDeleteInterface[Lyric])
	case *Lyric_font:
		stage.OnAfterLyric_fontDeleteCallback = any(callback).(OnAfterDeleteInterface[Lyric_font])
	case *Lyric_language:
		stage.OnAfterLyric_languageDeleteCallback = any(callback).(OnAfterDeleteInterface[Lyric_language])
	case *Measure_layout:
		stage.OnAfterMeasure_layoutDeleteCallback = any(callback).(OnAfterDeleteInterface[Measure_layout])
	case *Measure_numbering:
		stage.OnAfterMeasure_numberingDeleteCallback = any(callback).(OnAfterDeleteInterface[Measure_numbering])
	case *Measure_repeat:
		stage.OnAfterMeasure_repeatDeleteCallback = any(callback).(OnAfterDeleteInterface[Measure_repeat])
	case *Measure_style:
		stage.OnAfterMeasure_styleDeleteCallback = any(callback).(OnAfterDeleteInterface[Measure_style])
	case *Membrane:
		stage.OnAfterMembraneDeleteCallback = any(callback).(OnAfterDeleteInterface[Membrane])
	case *Metal:
		stage.OnAfterMetalDeleteCallback = any(callback).(OnAfterDeleteInterface[Metal])
	case *Metronome:
		stage.OnAfterMetronomeDeleteCallback = any(callback).(OnAfterDeleteInterface[Metronome])
	case *Metronome_beam:
		stage.OnAfterMetronome_beamDeleteCallback = any(callback).(OnAfterDeleteInterface[Metronome_beam])
	case *Metronome_note:
		stage.OnAfterMetronome_noteDeleteCallback = any(callback).(OnAfterDeleteInterface[Metronome_note])
	case *Metronome_tied:
		stage.OnAfterMetronome_tiedDeleteCallback = any(callback).(OnAfterDeleteInterface[Metronome_tied])
	case *Metronome_tuplet:
		stage.OnAfterMetronome_tupletDeleteCallback = any(callback).(OnAfterDeleteInterface[Metronome_tuplet])
	case *Midi_device:
		stage.OnAfterMidi_deviceDeleteCallback = any(callback).(OnAfterDeleteInterface[Midi_device])
	case *Midi_instrument:
		stage.OnAfterMidi_instrumentDeleteCallback = any(callback).(OnAfterDeleteInterface[Midi_instrument])
	case *Miscellaneous:
		stage.OnAfterMiscellaneousDeleteCallback = any(callback).(OnAfterDeleteInterface[Miscellaneous])
	case *Miscellaneous_field:
		stage.OnAfterMiscellaneous_fieldDeleteCallback = any(callback).(OnAfterDeleteInterface[Miscellaneous_field])
	case *Mordent:
		stage.OnAfterMordentDeleteCallback = any(callback).(OnAfterDeleteInterface[Mordent])
	case *Multiple_rest:
		stage.OnAfterMultiple_restDeleteCallback = any(callback).(OnAfterDeleteInterface[Multiple_rest])
	case *Name_display:
		stage.OnAfterName_displayDeleteCallback = any(callback).(OnAfterDeleteInterface[Name_display])
	case *Non_arpeggiate:
		stage.OnAfterNon_arpeggiateDeleteCallback = any(callback).(OnAfterDeleteInterface[Non_arpeggiate])
	case *Notations:
		stage.OnAfterNotationsDeleteCallback = any(callback).(OnAfterDeleteInterface[Notations])
	case *Note:
		stage.OnAfterNoteDeleteCallback = any(callback).(OnAfterDeleteInterface[Note])
	case *Note_size:
		stage.OnAfterNote_sizeDeleteCallback = any(callback).(OnAfterDeleteInterface[Note_size])
	case *Note_type:
		stage.OnAfterNote_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Note_type])
	case *Notehead:
		stage.OnAfterNoteheadDeleteCallback = any(callback).(OnAfterDeleteInterface[Notehead])
	case *Notehead_text:
		stage.OnAfterNotehead_textDeleteCallback = any(callback).(OnAfterDeleteInterface[Notehead_text])
	case *Numeral:
		stage.OnAfterNumeralDeleteCallback = any(callback).(OnAfterDeleteInterface[Numeral])
	case *Numeral_key:
		stage.OnAfterNumeral_keyDeleteCallback = any(callback).(OnAfterDeleteInterface[Numeral_key])
	case *Numeral_root:
		stage.OnAfterNumeral_rootDeleteCallback = any(callback).(OnAfterDeleteInterface[Numeral_root])
	case *Octave_shift:
		stage.OnAfterOctave_shiftDeleteCallback = any(callback).(OnAfterDeleteInterface[Octave_shift])
	case *Offset:
		stage.OnAfterOffsetDeleteCallback = any(callback).(OnAfterDeleteInterface[Offset])
	case *Opus:
		stage.OnAfterOpusDeleteCallback = any(callback).(OnAfterDeleteInterface[Opus])
	case *Ornaments:
		stage.OnAfterOrnamentsDeleteCallback = any(callback).(OnAfterDeleteInterface[Ornaments])
	case *Other_appearance:
		stage.OnAfterOther_appearanceDeleteCallback = any(callback).(OnAfterDeleteInterface[Other_appearance])
	case *Other_direction:
		stage.OnAfterOther_directionDeleteCallback = any(callback).(OnAfterDeleteInterface[Other_direction])
	case *Other_listening:
		stage.OnAfterOther_listeningDeleteCallback = any(callback).(OnAfterDeleteInterface[Other_listening])
	case *Other_notation:
		stage.OnAfterOther_notationDeleteCallback = any(callback).(OnAfterDeleteInterface[Other_notation])
	case *Other_placement_text:
		stage.OnAfterOther_placement_textDeleteCallback = any(callback).(OnAfterDeleteInterface[Other_placement_text])
	case *Other_play:
		stage.OnAfterOther_playDeleteCallback = any(callback).(OnAfterDeleteInterface[Other_play])
	case *Other_text:
		stage.OnAfterOther_textDeleteCallback = any(callback).(OnAfterDeleteInterface[Other_text])
	case *Page_layout:
		stage.OnAfterPage_layoutDeleteCallback = any(callback).(OnAfterDeleteInterface[Page_layout])
	case *Page_margins:
		stage.OnAfterPage_marginsDeleteCallback = any(callback).(OnAfterDeleteInterface[Page_margins])
	case *Part_clef:
		stage.OnAfterPart_clefDeleteCallback = any(callback).(OnAfterDeleteInterface[Part_clef])
	case *Part_group:
		stage.OnAfterPart_groupDeleteCallback = any(callback).(OnAfterDeleteInterface[Part_group])
	case *Part_link:
		stage.OnAfterPart_linkDeleteCallback = any(callback).(OnAfterDeleteInterface[Part_link])
	case *Part_list:
		stage.OnAfterPart_listDeleteCallback = any(callback).(OnAfterDeleteInterface[Part_list])
	case *Part_name:
		stage.OnAfterPart_nameDeleteCallback = any(callback).(OnAfterDeleteInterface[Part_name])
	case *Part_symbol:
		stage.OnAfterPart_symbolDeleteCallback = any(callback).(OnAfterDeleteInterface[Part_symbol])
	case *Part_transpose:
		stage.OnAfterPart_transposeDeleteCallback = any(callback).(OnAfterDeleteInterface[Part_transpose])
	case *Pedal:
		stage.OnAfterPedalDeleteCallback = any(callback).(OnAfterDeleteInterface[Pedal])
	case *Pedal_tuning:
		stage.OnAfterPedal_tuningDeleteCallback = any(callback).(OnAfterDeleteInterface[Pedal_tuning])
	case *Per_minute:
		stage.OnAfterPer_minuteDeleteCallback = any(callback).(OnAfterDeleteInterface[Per_minute])
	case *Percussion:
		stage.OnAfterPercussionDeleteCallback = any(callback).(OnAfterDeleteInterface[Percussion])
	case *Pitch:
		stage.OnAfterPitchDeleteCallback = any(callback).(OnAfterDeleteInterface[Pitch])
	case *Pitched:
		stage.OnAfterPitchedDeleteCallback = any(callback).(OnAfterDeleteInterface[Pitched])
	case *Placement_text:
		stage.OnAfterPlacement_textDeleteCallback = any(callback).(OnAfterDeleteInterface[Placement_text])
	case *Play:
		stage.OnAfterPlayDeleteCallback = any(callback).(OnAfterDeleteInterface[Play])
	case *Player:
		stage.OnAfterPlayerDeleteCallback = any(callback).(OnAfterDeleteInterface[Player])
	case *Principal_voice:
		stage.OnAfterPrincipal_voiceDeleteCallback = any(callback).(OnAfterDeleteInterface[Principal_voice])
	case *Print:
		stage.OnAfterPrintDeleteCallback = any(callback).(OnAfterDeleteInterface[Print])
	case *Release:
		stage.OnAfterReleaseDeleteCallback = any(callback).(OnAfterDeleteInterface[Release])
	case *Repeat:
		stage.OnAfterRepeatDeleteCallback = any(callback).(OnAfterDeleteInterface[Repeat])
	case *Rest:
		stage.OnAfterRestDeleteCallback = any(callback).(OnAfterDeleteInterface[Rest])
	case *Root:
		stage.OnAfterRootDeleteCallback = any(callback).(OnAfterDeleteInterface[Root])
	case *Root_step:
		stage.OnAfterRoot_stepDeleteCallback = any(callback).(OnAfterDeleteInterface[Root_step])
	case *Scaling:
		stage.OnAfterScalingDeleteCallback = any(callback).(OnAfterDeleteInterface[Scaling])
	case *Scordatura:
		stage.OnAfterScordaturaDeleteCallback = any(callback).(OnAfterDeleteInterface[Scordatura])
	case *Score_instrument:
		stage.OnAfterScore_instrumentDeleteCallback = any(callback).(OnAfterDeleteInterface[Score_instrument])
	case *Score_part:
		stage.OnAfterScore_partDeleteCallback = any(callback).(OnAfterDeleteInterface[Score_part])
	case *Score_partwise:
		stage.OnAfterScore_partwiseDeleteCallback = any(callback).(OnAfterDeleteInterface[Score_partwise])
	case *Score_timewise:
		stage.OnAfterScore_timewiseDeleteCallback = any(callback).(OnAfterDeleteInterface[Score_timewise])
	case *Segno:
		stage.OnAfterSegnoDeleteCallback = any(callback).(OnAfterDeleteInterface[Segno])
	case *Slash:
		stage.OnAfterSlashDeleteCallback = any(callback).(OnAfterDeleteInterface[Slash])
	case *Slide:
		stage.OnAfterSlideDeleteCallback = any(callback).(OnAfterDeleteInterface[Slide])
	case *Slur:
		stage.OnAfterSlurDeleteCallback = any(callback).(OnAfterDeleteInterface[Slur])
	case *Sound:
		stage.OnAfterSoundDeleteCallback = any(callback).(OnAfterDeleteInterface[Sound])
	case *Staff_details:
		stage.OnAfterStaff_detailsDeleteCallback = any(callback).(OnAfterDeleteInterface[Staff_details])
	case *Staff_divide:
		stage.OnAfterStaff_divideDeleteCallback = any(callback).(OnAfterDeleteInterface[Staff_divide])
	case *Staff_layout:
		stage.OnAfterStaff_layoutDeleteCallback = any(callback).(OnAfterDeleteInterface[Staff_layout])
	case *Staff_size:
		stage.OnAfterStaff_sizeDeleteCallback = any(callback).(OnAfterDeleteInterface[Staff_size])
	case *Staff_tuning:
		stage.OnAfterStaff_tuningDeleteCallback = any(callback).(OnAfterDeleteInterface[Staff_tuning])
	case *Stem:
		stage.OnAfterStemDeleteCallback = any(callback).(OnAfterDeleteInterface[Stem])
	case *Stick:
		stage.OnAfterStickDeleteCallback = any(callback).(OnAfterDeleteInterface[Stick])
	case *String_mute:
		stage.OnAfterString_muteDeleteCallback = any(callback).(OnAfterDeleteInterface[String_mute])
	case *String_type:
		stage.OnAfterString_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[String_type])
	case *Strong_accent:
		stage.OnAfterStrong_accentDeleteCallback = any(callback).(OnAfterDeleteInterface[Strong_accent])
	case *Style_text:
		stage.OnAfterStyle_textDeleteCallback = any(callback).(OnAfterDeleteInterface[Style_text])
	case *Supports:
		stage.OnAfterSupportsDeleteCallback = any(callback).(OnAfterDeleteInterface[Supports])
	case *Swing:
		stage.OnAfterSwingDeleteCallback = any(callback).(OnAfterDeleteInterface[Swing])
	case *Sync:
		stage.OnAfterSyncDeleteCallback = any(callback).(OnAfterDeleteInterface[Sync])
	case *System_dividers:
		stage.OnAfterSystem_dividersDeleteCallback = any(callback).(OnAfterDeleteInterface[System_dividers])
	case *System_layout:
		stage.OnAfterSystem_layoutDeleteCallback = any(callback).(OnAfterDeleteInterface[System_layout])
	case *System_margins:
		stage.OnAfterSystem_marginsDeleteCallback = any(callback).(OnAfterDeleteInterface[System_margins])
	case *Tap:
		stage.OnAfterTapDeleteCallback = any(callback).(OnAfterDeleteInterface[Tap])
	case *Technical:
		stage.OnAfterTechnicalDeleteCallback = any(callback).(OnAfterDeleteInterface[Technical])
	case *Text_element_data:
		stage.OnAfterText_element_dataDeleteCallback = any(callback).(OnAfterDeleteInterface[Text_element_data])
	case *Tie:
		stage.OnAfterTieDeleteCallback = any(callback).(OnAfterDeleteInterface[Tie])
	case *Tied:
		stage.OnAfterTiedDeleteCallback = any(callback).(OnAfterDeleteInterface[Tied])
	case *Time:
		stage.OnAfterTimeDeleteCallback = any(callback).(OnAfterDeleteInterface[Time])
	case *Time_modification:
		stage.OnAfterTime_modificationDeleteCallback = any(callback).(OnAfterDeleteInterface[Time_modification])
	case *Timpani:
		stage.OnAfterTimpaniDeleteCallback = any(callback).(OnAfterDeleteInterface[Timpani])
	case *Transpose:
		stage.OnAfterTransposeDeleteCallback = any(callback).(OnAfterDeleteInterface[Transpose])
	case *Tremolo:
		stage.OnAfterTremoloDeleteCallback = any(callback).(OnAfterDeleteInterface[Tremolo])
	case *Tuplet:
		stage.OnAfterTupletDeleteCallback = any(callback).(OnAfterDeleteInterface[Tuplet])
	case *Tuplet_dot:
		stage.OnAfterTuplet_dotDeleteCallback = any(callback).(OnAfterDeleteInterface[Tuplet_dot])
	case *Tuplet_number:
		stage.OnAfterTuplet_numberDeleteCallback = any(callback).(OnAfterDeleteInterface[Tuplet_number])
	case *Tuplet_portion:
		stage.OnAfterTuplet_portionDeleteCallback = any(callback).(OnAfterDeleteInterface[Tuplet_portion])
	case *Tuplet_type:
		stage.OnAfterTuplet_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Tuplet_type])
	case *Typed_text:
		stage.OnAfterTyped_textDeleteCallback = any(callback).(OnAfterDeleteInterface[Typed_text])
	case *Unpitched:
		stage.OnAfterUnpitchedDeleteCallback = any(callback).(OnAfterDeleteInterface[Unpitched])
	case *Virtual_instrument:
		stage.OnAfterVirtual_instrumentDeleteCallback = any(callback).(OnAfterDeleteInterface[Virtual_instrument])
	case *Wait:
		stage.OnAfterWaitDeleteCallback = any(callback).(OnAfterDeleteInterface[Wait])
	case *Wavy_line:
		stage.OnAfterWavy_lineDeleteCallback = any(callback).(OnAfterDeleteInterface[Wavy_line])
	case *Wedge:
		stage.OnAfterWedgeDeleteCallback = any(callback).(OnAfterDeleteInterface[Wedge])
	case *Wood:
		stage.OnAfterWoodDeleteCallback = any(callback).(OnAfterDeleteInterface[Wood])
	case *Work:
		stage.OnAfterWorkDeleteCallback = any(callback).(OnAfterDeleteInterface[Work])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *A_directive:
		stage.OnAfterA_directiveReadCallback = any(callback).(OnAfterReadInterface[A_directive])
	case *A_measure:
		stage.OnAfterA_measureReadCallback = any(callback).(OnAfterReadInterface[A_measure])
	case *A_measure_1:
		stage.OnAfterA_measure_1ReadCallback = any(callback).(OnAfterReadInterface[A_measure_1])
	case *A_part:
		stage.OnAfterA_partReadCallback = any(callback).(OnAfterReadInterface[A_part])
	case *A_part_1:
		stage.OnAfterA_part_1ReadCallback = any(callback).(OnAfterReadInterface[A_part_1])
	case *Accidental:
		stage.OnAfterAccidentalReadCallback = any(callback).(OnAfterReadInterface[Accidental])
	case *Accidental_mark:
		stage.OnAfterAccidental_markReadCallback = any(callback).(OnAfterReadInterface[Accidental_mark])
	case *Accidental_text:
		stage.OnAfterAccidental_textReadCallback = any(callback).(OnAfterReadInterface[Accidental_text])
	case *Accord:
		stage.OnAfterAccordReadCallback = any(callback).(OnAfterReadInterface[Accord])
	case *Accordion_registration:
		stage.OnAfterAccordion_registrationReadCallback = any(callback).(OnAfterReadInterface[Accordion_registration])
	case *Appearance:
		stage.OnAfterAppearanceReadCallback = any(callback).(OnAfterReadInterface[Appearance])
	case *Arpeggiate:
		stage.OnAfterArpeggiateReadCallback = any(callback).(OnAfterReadInterface[Arpeggiate])
	case *Arrow:
		stage.OnAfterArrowReadCallback = any(callback).(OnAfterReadInterface[Arrow])
	case *Articulations:
		stage.OnAfterArticulationsReadCallback = any(callback).(OnAfterReadInterface[Articulations])
	case *Assess:
		stage.OnAfterAssessReadCallback = any(callback).(OnAfterReadInterface[Assess])
	case *Attributes:
		stage.OnAfterAttributesReadCallback = any(callback).(OnAfterReadInterface[Attributes])
	case *Backup:
		stage.OnAfterBackupReadCallback = any(callback).(OnAfterReadInterface[Backup])
	case *Bar_style_color:
		stage.OnAfterBar_style_colorReadCallback = any(callback).(OnAfterReadInterface[Bar_style_color])
	case *Barline:
		stage.OnAfterBarlineReadCallback = any(callback).(OnAfterReadInterface[Barline])
	case *Barre:
		stage.OnAfterBarreReadCallback = any(callback).(OnAfterReadInterface[Barre])
	case *Bass:
		stage.OnAfterBassReadCallback = any(callback).(OnAfterReadInterface[Bass])
	case *Bass_step:
		stage.OnAfterBass_stepReadCallback = any(callback).(OnAfterReadInterface[Bass_step])
	case *Beam:
		stage.OnAfterBeamReadCallback = any(callback).(OnAfterReadInterface[Beam])
	case *Beat_repeat:
		stage.OnAfterBeat_repeatReadCallback = any(callback).(OnAfterReadInterface[Beat_repeat])
	case *Beat_unit_tied:
		stage.OnAfterBeat_unit_tiedReadCallback = any(callback).(OnAfterReadInterface[Beat_unit_tied])
	case *Beater:
		stage.OnAfterBeaterReadCallback = any(callback).(OnAfterReadInterface[Beater])
	case *Bend:
		stage.OnAfterBendReadCallback = any(callback).(OnAfterReadInterface[Bend])
	case *Bookmark:
		stage.OnAfterBookmarkReadCallback = any(callback).(OnAfterReadInterface[Bookmark])
	case *Bracket:
		stage.OnAfterBracketReadCallback = any(callback).(OnAfterReadInterface[Bracket])
	case *Breath_mark:
		stage.OnAfterBreath_markReadCallback = any(callback).(OnAfterReadInterface[Breath_mark])
	case *Caesura:
		stage.OnAfterCaesuraReadCallback = any(callback).(OnAfterReadInterface[Caesura])
	case *Cancel:
		stage.OnAfterCancelReadCallback = any(callback).(OnAfterReadInterface[Cancel])
	case *Clef:
		stage.OnAfterClefReadCallback = any(callback).(OnAfterReadInterface[Clef])
	case *Coda:
		stage.OnAfterCodaReadCallback = any(callback).(OnAfterReadInterface[Coda])
	case *Credit:
		stage.OnAfterCreditReadCallback = any(callback).(OnAfterReadInterface[Credit])
	case *Dashes:
		stage.OnAfterDashesReadCallback = any(callback).(OnAfterReadInterface[Dashes])
	case *Defaults:
		stage.OnAfterDefaultsReadCallback = any(callback).(OnAfterReadInterface[Defaults])
	case *Degree:
		stage.OnAfterDegreeReadCallback = any(callback).(OnAfterReadInterface[Degree])
	case *Degree_alter:
		stage.OnAfterDegree_alterReadCallback = any(callback).(OnAfterReadInterface[Degree_alter])
	case *Degree_type:
		stage.OnAfterDegree_typeReadCallback = any(callback).(OnAfterReadInterface[Degree_type])
	case *Degree_value:
		stage.OnAfterDegree_valueReadCallback = any(callback).(OnAfterReadInterface[Degree_value])
	case *Direction:
		stage.OnAfterDirectionReadCallback = any(callback).(OnAfterReadInterface[Direction])
	case *Direction_type:
		stage.OnAfterDirection_typeReadCallback = any(callback).(OnAfterReadInterface[Direction_type])
	case *Distance:
		stage.OnAfterDistanceReadCallback = any(callback).(OnAfterReadInterface[Distance])
	case *Double:
		stage.OnAfterDoubleReadCallback = any(callback).(OnAfterReadInterface[Double])
	case *Dynamics:
		stage.OnAfterDynamicsReadCallback = any(callback).(OnAfterReadInterface[Dynamics])
	case *Effect:
		stage.OnAfterEffectReadCallback = any(callback).(OnAfterReadInterface[Effect])
	case *Elision:
		stage.OnAfterElisionReadCallback = any(callback).(OnAfterReadInterface[Elision])
	case *Empty:
		stage.OnAfterEmptyReadCallback = any(callback).(OnAfterReadInterface[Empty])
	case *Empty_font:
		stage.OnAfterEmpty_fontReadCallback = any(callback).(OnAfterReadInterface[Empty_font])
	case *Empty_line:
		stage.OnAfterEmpty_lineReadCallback = any(callback).(OnAfterReadInterface[Empty_line])
	case *Empty_placement:
		stage.OnAfterEmpty_placementReadCallback = any(callback).(OnAfterReadInterface[Empty_placement])
	case *Empty_placement_smufl:
		stage.OnAfterEmpty_placement_smuflReadCallback = any(callback).(OnAfterReadInterface[Empty_placement_smufl])
	case *Empty_print_object_style_align:
		stage.OnAfterEmpty_print_object_style_alignReadCallback = any(callback).(OnAfterReadInterface[Empty_print_object_style_align])
	case *Empty_print_style:
		stage.OnAfterEmpty_print_styleReadCallback = any(callback).(OnAfterReadInterface[Empty_print_style])
	case *Empty_print_style_align:
		stage.OnAfterEmpty_print_style_alignReadCallback = any(callback).(OnAfterReadInterface[Empty_print_style_align])
	case *Empty_print_style_align_id:
		stage.OnAfterEmpty_print_style_align_idReadCallback = any(callback).(OnAfterReadInterface[Empty_print_style_align_id])
	case *Empty_trill_sound:
		stage.OnAfterEmpty_trill_soundReadCallback = any(callback).(OnAfterReadInterface[Empty_trill_sound])
	case *Encoding:
		stage.OnAfterEncodingReadCallback = any(callback).(OnAfterReadInterface[Encoding])
	case *Ending:
		stage.OnAfterEndingReadCallback = any(callback).(OnAfterReadInterface[Ending])
	case *Extend:
		stage.OnAfterExtendReadCallback = any(callback).(OnAfterReadInterface[Extend])
	case *Feature:
		stage.OnAfterFeatureReadCallback = any(callback).(OnAfterReadInterface[Feature])
	case *Fermata:
		stage.OnAfterFermataReadCallback = any(callback).(OnAfterReadInterface[Fermata])
	case *Figure:
		stage.OnAfterFigureReadCallback = any(callback).(OnAfterReadInterface[Figure])
	case *Figured_bass:
		stage.OnAfterFigured_bassReadCallback = any(callback).(OnAfterReadInterface[Figured_bass])
	case *Fingering:
		stage.OnAfterFingeringReadCallback = any(callback).(OnAfterReadInterface[Fingering])
	case *First_fret:
		stage.OnAfterFirst_fretReadCallback = any(callback).(OnAfterReadInterface[First_fret])
	case *For_part:
		stage.OnAfterFor_partReadCallback = any(callback).(OnAfterReadInterface[For_part])
	case *Formatted_symbol:
		stage.OnAfterFormatted_symbolReadCallback = any(callback).(OnAfterReadInterface[Formatted_symbol])
	case *Formatted_symbol_id:
		stage.OnAfterFormatted_symbol_idReadCallback = any(callback).(OnAfterReadInterface[Formatted_symbol_id])
	case *Formatted_text:
		stage.OnAfterFormatted_textReadCallback = any(callback).(OnAfterReadInterface[Formatted_text])
	case *Formatted_text_id:
		stage.OnAfterFormatted_text_idReadCallback = any(callback).(OnAfterReadInterface[Formatted_text_id])
	case *Forward:
		stage.OnAfterForwardReadCallback = any(callback).(OnAfterReadInterface[Forward])
	case *Frame:
		stage.OnAfterFrameReadCallback = any(callback).(OnAfterReadInterface[Frame])
	case *Frame_note:
		stage.OnAfterFrame_noteReadCallback = any(callback).(OnAfterReadInterface[Frame_note])
	case *Fret:
		stage.OnAfterFretReadCallback = any(callback).(OnAfterReadInterface[Fret])
	case *Glass:
		stage.OnAfterGlassReadCallback = any(callback).(OnAfterReadInterface[Glass])
	case *Glissando:
		stage.OnAfterGlissandoReadCallback = any(callback).(OnAfterReadInterface[Glissando])
	case *Glyph:
		stage.OnAfterGlyphReadCallback = any(callback).(OnAfterReadInterface[Glyph])
	case *Grace:
		stage.OnAfterGraceReadCallback = any(callback).(OnAfterReadInterface[Grace])
	case *Group_barline:
		stage.OnAfterGroup_barlineReadCallback = any(callback).(OnAfterReadInterface[Group_barline])
	case *Group_name:
		stage.OnAfterGroup_nameReadCallback = any(callback).(OnAfterReadInterface[Group_name])
	case *Group_symbol:
		stage.OnAfterGroup_symbolReadCallback = any(callback).(OnAfterReadInterface[Group_symbol])
	case *Grouping:
		stage.OnAfterGroupingReadCallback = any(callback).(OnAfterReadInterface[Grouping])
	case *Hammer_on_pull_off:
		stage.OnAfterHammer_on_pull_offReadCallback = any(callback).(OnAfterReadInterface[Hammer_on_pull_off])
	case *Handbell:
		stage.OnAfterHandbellReadCallback = any(callback).(OnAfterReadInterface[Handbell])
	case *Harmon_closed:
		stage.OnAfterHarmon_closedReadCallback = any(callback).(OnAfterReadInterface[Harmon_closed])
	case *Harmon_mute:
		stage.OnAfterHarmon_muteReadCallback = any(callback).(OnAfterReadInterface[Harmon_mute])
	case *Harmonic:
		stage.OnAfterHarmonicReadCallback = any(callback).(OnAfterReadInterface[Harmonic])
	case *Harmony:
		stage.OnAfterHarmonyReadCallback = any(callback).(OnAfterReadInterface[Harmony])
	case *Harmony_alter:
		stage.OnAfterHarmony_alterReadCallback = any(callback).(OnAfterReadInterface[Harmony_alter])
	case *Harp_pedals:
		stage.OnAfterHarp_pedalsReadCallback = any(callback).(OnAfterReadInterface[Harp_pedals])
	case *Heel_toe:
		stage.OnAfterHeel_toeReadCallback = any(callback).(OnAfterReadInterface[Heel_toe])
	case *Hole:
		stage.OnAfterHoleReadCallback = any(callback).(OnAfterReadInterface[Hole])
	case *Hole_closed:
		stage.OnAfterHole_closedReadCallback = any(callback).(OnAfterReadInterface[Hole_closed])
	case *Horizontal_turn:
		stage.OnAfterHorizontal_turnReadCallback = any(callback).(OnAfterReadInterface[Horizontal_turn])
	case *Identification:
		stage.OnAfterIdentificationReadCallback = any(callback).(OnAfterReadInterface[Identification])
	case *Image:
		stage.OnAfterImageReadCallback = any(callback).(OnAfterReadInterface[Image])
	case *Instrument:
		stage.OnAfterInstrumentReadCallback = any(callback).(OnAfterReadInterface[Instrument])
	case *Instrument_change:
		stage.OnAfterInstrument_changeReadCallback = any(callback).(OnAfterReadInterface[Instrument_change])
	case *Instrument_link:
		stage.OnAfterInstrument_linkReadCallback = any(callback).(OnAfterReadInterface[Instrument_link])
	case *Interchangeable:
		stage.OnAfterInterchangeableReadCallback = any(callback).(OnAfterReadInterface[Interchangeable])
	case *Inversion:
		stage.OnAfterInversionReadCallback = any(callback).(OnAfterReadInterface[Inversion])
	case *Key:
		stage.OnAfterKeyReadCallback = any(callback).(OnAfterReadInterface[Key])
	case *Key_accidental:
		stage.OnAfterKey_accidentalReadCallback = any(callback).(OnAfterReadInterface[Key_accidental])
	case *Key_octave:
		stage.OnAfterKey_octaveReadCallback = any(callback).(OnAfterReadInterface[Key_octave])
	case *Kind:
		stage.OnAfterKindReadCallback = any(callback).(OnAfterReadInterface[Kind])
	case *Level:
		stage.OnAfterLevelReadCallback = any(callback).(OnAfterReadInterface[Level])
	case *Line_detail:
		stage.OnAfterLine_detailReadCallback = any(callback).(OnAfterReadInterface[Line_detail])
	case *Line_width:
		stage.OnAfterLine_widthReadCallback = any(callback).(OnAfterReadInterface[Line_width])
	case *Link:
		stage.OnAfterLinkReadCallback = any(callback).(OnAfterReadInterface[Link])
	case *Listen:
		stage.OnAfterListenReadCallback = any(callback).(OnAfterReadInterface[Listen])
	case *Listening:
		stage.OnAfterListeningReadCallback = any(callback).(OnAfterReadInterface[Listening])
	case *Lyric:
		stage.OnAfterLyricReadCallback = any(callback).(OnAfterReadInterface[Lyric])
	case *Lyric_font:
		stage.OnAfterLyric_fontReadCallback = any(callback).(OnAfterReadInterface[Lyric_font])
	case *Lyric_language:
		stage.OnAfterLyric_languageReadCallback = any(callback).(OnAfterReadInterface[Lyric_language])
	case *Measure_layout:
		stage.OnAfterMeasure_layoutReadCallback = any(callback).(OnAfterReadInterface[Measure_layout])
	case *Measure_numbering:
		stage.OnAfterMeasure_numberingReadCallback = any(callback).(OnAfterReadInterface[Measure_numbering])
	case *Measure_repeat:
		stage.OnAfterMeasure_repeatReadCallback = any(callback).(OnAfterReadInterface[Measure_repeat])
	case *Measure_style:
		stage.OnAfterMeasure_styleReadCallback = any(callback).(OnAfterReadInterface[Measure_style])
	case *Membrane:
		stage.OnAfterMembraneReadCallback = any(callback).(OnAfterReadInterface[Membrane])
	case *Metal:
		stage.OnAfterMetalReadCallback = any(callback).(OnAfterReadInterface[Metal])
	case *Metronome:
		stage.OnAfterMetronomeReadCallback = any(callback).(OnAfterReadInterface[Metronome])
	case *Metronome_beam:
		stage.OnAfterMetronome_beamReadCallback = any(callback).(OnAfterReadInterface[Metronome_beam])
	case *Metronome_note:
		stage.OnAfterMetronome_noteReadCallback = any(callback).(OnAfterReadInterface[Metronome_note])
	case *Metronome_tied:
		stage.OnAfterMetronome_tiedReadCallback = any(callback).(OnAfterReadInterface[Metronome_tied])
	case *Metronome_tuplet:
		stage.OnAfterMetronome_tupletReadCallback = any(callback).(OnAfterReadInterface[Metronome_tuplet])
	case *Midi_device:
		stage.OnAfterMidi_deviceReadCallback = any(callback).(OnAfterReadInterface[Midi_device])
	case *Midi_instrument:
		stage.OnAfterMidi_instrumentReadCallback = any(callback).(OnAfterReadInterface[Midi_instrument])
	case *Miscellaneous:
		stage.OnAfterMiscellaneousReadCallback = any(callback).(OnAfterReadInterface[Miscellaneous])
	case *Miscellaneous_field:
		stage.OnAfterMiscellaneous_fieldReadCallback = any(callback).(OnAfterReadInterface[Miscellaneous_field])
	case *Mordent:
		stage.OnAfterMordentReadCallback = any(callback).(OnAfterReadInterface[Mordent])
	case *Multiple_rest:
		stage.OnAfterMultiple_restReadCallback = any(callback).(OnAfterReadInterface[Multiple_rest])
	case *Name_display:
		stage.OnAfterName_displayReadCallback = any(callback).(OnAfterReadInterface[Name_display])
	case *Non_arpeggiate:
		stage.OnAfterNon_arpeggiateReadCallback = any(callback).(OnAfterReadInterface[Non_arpeggiate])
	case *Notations:
		stage.OnAfterNotationsReadCallback = any(callback).(OnAfterReadInterface[Notations])
	case *Note:
		stage.OnAfterNoteReadCallback = any(callback).(OnAfterReadInterface[Note])
	case *Note_size:
		stage.OnAfterNote_sizeReadCallback = any(callback).(OnAfterReadInterface[Note_size])
	case *Note_type:
		stage.OnAfterNote_typeReadCallback = any(callback).(OnAfterReadInterface[Note_type])
	case *Notehead:
		stage.OnAfterNoteheadReadCallback = any(callback).(OnAfterReadInterface[Notehead])
	case *Notehead_text:
		stage.OnAfterNotehead_textReadCallback = any(callback).(OnAfterReadInterface[Notehead_text])
	case *Numeral:
		stage.OnAfterNumeralReadCallback = any(callback).(OnAfterReadInterface[Numeral])
	case *Numeral_key:
		stage.OnAfterNumeral_keyReadCallback = any(callback).(OnAfterReadInterface[Numeral_key])
	case *Numeral_root:
		stage.OnAfterNumeral_rootReadCallback = any(callback).(OnAfterReadInterface[Numeral_root])
	case *Octave_shift:
		stage.OnAfterOctave_shiftReadCallback = any(callback).(OnAfterReadInterface[Octave_shift])
	case *Offset:
		stage.OnAfterOffsetReadCallback = any(callback).(OnAfterReadInterface[Offset])
	case *Opus:
		stage.OnAfterOpusReadCallback = any(callback).(OnAfterReadInterface[Opus])
	case *Ornaments:
		stage.OnAfterOrnamentsReadCallback = any(callback).(OnAfterReadInterface[Ornaments])
	case *Other_appearance:
		stage.OnAfterOther_appearanceReadCallback = any(callback).(OnAfterReadInterface[Other_appearance])
	case *Other_direction:
		stage.OnAfterOther_directionReadCallback = any(callback).(OnAfterReadInterface[Other_direction])
	case *Other_listening:
		stage.OnAfterOther_listeningReadCallback = any(callback).(OnAfterReadInterface[Other_listening])
	case *Other_notation:
		stage.OnAfterOther_notationReadCallback = any(callback).(OnAfterReadInterface[Other_notation])
	case *Other_placement_text:
		stage.OnAfterOther_placement_textReadCallback = any(callback).(OnAfterReadInterface[Other_placement_text])
	case *Other_play:
		stage.OnAfterOther_playReadCallback = any(callback).(OnAfterReadInterface[Other_play])
	case *Other_text:
		stage.OnAfterOther_textReadCallback = any(callback).(OnAfterReadInterface[Other_text])
	case *Page_layout:
		stage.OnAfterPage_layoutReadCallback = any(callback).(OnAfterReadInterface[Page_layout])
	case *Page_margins:
		stage.OnAfterPage_marginsReadCallback = any(callback).(OnAfterReadInterface[Page_margins])
	case *Part_clef:
		stage.OnAfterPart_clefReadCallback = any(callback).(OnAfterReadInterface[Part_clef])
	case *Part_group:
		stage.OnAfterPart_groupReadCallback = any(callback).(OnAfterReadInterface[Part_group])
	case *Part_link:
		stage.OnAfterPart_linkReadCallback = any(callback).(OnAfterReadInterface[Part_link])
	case *Part_list:
		stage.OnAfterPart_listReadCallback = any(callback).(OnAfterReadInterface[Part_list])
	case *Part_name:
		stage.OnAfterPart_nameReadCallback = any(callback).(OnAfterReadInterface[Part_name])
	case *Part_symbol:
		stage.OnAfterPart_symbolReadCallback = any(callback).(OnAfterReadInterface[Part_symbol])
	case *Part_transpose:
		stage.OnAfterPart_transposeReadCallback = any(callback).(OnAfterReadInterface[Part_transpose])
	case *Pedal:
		stage.OnAfterPedalReadCallback = any(callback).(OnAfterReadInterface[Pedal])
	case *Pedal_tuning:
		stage.OnAfterPedal_tuningReadCallback = any(callback).(OnAfterReadInterface[Pedal_tuning])
	case *Per_minute:
		stage.OnAfterPer_minuteReadCallback = any(callback).(OnAfterReadInterface[Per_minute])
	case *Percussion:
		stage.OnAfterPercussionReadCallback = any(callback).(OnAfterReadInterface[Percussion])
	case *Pitch:
		stage.OnAfterPitchReadCallback = any(callback).(OnAfterReadInterface[Pitch])
	case *Pitched:
		stage.OnAfterPitchedReadCallback = any(callback).(OnAfterReadInterface[Pitched])
	case *Placement_text:
		stage.OnAfterPlacement_textReadCallback = any(callback).(OnAfterReadInterface[Placement_text])
	case *Play:
		stage.OnAfterPlayReadCallback = any(callback).(OnAfterReadInterface[Play])
	case *Player:
		stage.OnAfterPlayerReadCallback = any(callback).(OnAfterReadInterface[Player])
	case *Principal_voice:
		stage.OnAfterPrincipal_voiceReadCallback = any(callback).(OnAfterReadInterface[Principal_voice])
	case *Print:
		stage.OnAfterPrintReadCallback = any(callback).(OnAfterReadInterface[Print])
	case *Release:
		stage.OnAfterReleaseReadCallback = any(callback).(OnAfterReadInterface[Release])
	case *Repeat:
		stage.OnAfterRepeatReadCallback = any(callback).(OnAfterReadInterface[Repeat])
	case *Rest:
		stage.OnAfterRestReadCallback = any(callback).(OnAfterReadInterface[Rest])
	case *Root:
		stage.OnAfterRootReadCallback = any(callback).(OnAfterReadInterface[Root])
	case *Root_step:
		stage.OnAfterRoot_stepReadCallback = any(callback).(OnAfterReadInterface[Root_step])
	case *Scaling:
		stage.OnAfterScalingReadCallback = any(callback).(OnAfterReadInterface[Scaling])
	case *Scordatura:
		stage.OnAfterScordaturaReadCallback = any(callback).(OnAfterReadInterface[Scordatura])
	case *Score_instrument:
		stage.OnAfterScore_instrumentReadCallback = any(callback).(OnAfterReadInterface[Score_instrument])
	case *Score_part:
		stage.OnAfterScore_partReadCallback = any(callback).(OnAfterReadInterface[Score_part])
	case *Score_partwise:
		stage.OnAfterScore_partwiseReadCallback = any(callback).(OnAfterReadInterface[Score_partwise])
	case *Score_timewise:
		stage.OnAfterScore_timewiseReadCallback = any(callback).(OnAfterReadInterface[Score_timewise])
	case *Segno:
		stage.OnAfterSegnoReadCallback = any(callback).(OnAfterReadInterface[Segno])
	case *Slash:
		stage.OnAfterSlashReadCallback = any(callback).(OnAfterReadInterface[Slash])
	case *Slide:
		stage.OnAfterSlideReadCallback = any(callback).(OnAfterReadInterface[Slide])
	case *Slur:
		stage.OnAfterSlurReadCallback = any(callback).(OnAfterReadInterface[Slur])
	case *Sound:
		stage.OnAfterSoundReadCallback = any(callback).(OnAfterReadInterface[Sound])
	case *Staff_details:
		stage.OnAfterStaff_detailsReadCallback = any(callback).(OnAfterReadInterface[Staff_details])
	case *Staff_divide:
		stage.OnAfterStaff_divideReadCallback = any(callback).(OnAfterReadInterface[Staff_divide])
	case *Staff_layout:
		stage.OnAfterStaff_layoutReadCallback = any(callback).(OnAfterReadInterface[Staff_layout])
	case *Staff_size:
		stage.OnAfterStaff_sizeReadCallback = any(callback).(OnAfterReadInterface[Staff_size])
	case *Staff_tuning:
		stage.OnAfterStaff_tuningReadCallback = any(callback).(OnAfterReadInterface[Staff_tuning])
	case *Stem:
		stage.OnAfterStemReadCallback = any(callback).(OnAfterReadInterface[Stem])
	case *Stick:
		stage.OnAfterStickReadCallback = any(callback).(OnAfterReadInterface[Stick])
	case *String_mute:
		stage.OnAfterString_muteReadCallback = any(callback).(OnAfterReadInterface[String_mute])
	case *String_type:
		stage.OnAfterString_typeReadCallback = any(callback).(OnAfterReadInterface[String_type])
	case *Strong_accent:
		stage.OnAfterStrong_accentReadCallback = any(callback).(OnAfterReadInterface[Strong_accent])
	case *Style_text:
		stage.OnAfterStyle_textReadCallback = any(callback).(OnAfterReadInterface[Style_text])
	case *Supports:
		stage.OnAfterSupportsReadCallback = any(callback).(OnAfterReadInterface[Supports])
	case *Swing:
		stage.OnAfterSwingReadCallback = any(callback).(OnAfterReadInterface[Swing])
	case *Sync:
		stage.OnAfterSyncReadCallback = any(callback).(OnAfterReadInterface[Sync])
	case *System_dividers:
		stage.OnAfterSystem_dividersReadCallback = any(callback).(OnAfterReadInterface[System_dividers])
	case *System_layout:
		stage.OnAfterSystem_layoutReadCallback = any(callback).(OnAfterReadInterface[System_layout])
	case *System_margins:
		stage.OnAfterSystem_marginsReadCallback = any(callback).(OnAfterReadInterface[System_margins])
	case *Tap:
		stage.OnAfterTapReadCallback = any(callback).(OnAfterReadInterface[Tap])
	case *Technical:
		stage.OnAfterTechnicalReadCallback = any(callback).(OnAfterReadInterface[Technical])
	case *Text_element_data:
		stage.OnAfterText_element_dataReadCallback = any(callback).(OnAfterReadInterface[Text_element_data])
	case *Tie:
		stage.OnAfterTieReadCallback = any(callback).(OnAfterReadInterface[Tie])
	case *Tied:
		stage.OnAfterTiedReadCallback = any(callback).(OnAfterReadInterface[Tied])
	case *Time:
		stage.OnAfterTimeReadCallback = any(callback).(OnAfterReadInterface[Time])
	case *Time_modification:
		stage.OnAfterTime_modificationReadCallback = any(callback).(OnAfterReadInterface[Time_modification])
	case *Timpani:
		stage.OnAfterTimpaniReadCallback = any(callback).(OnAfterReadInterface[Timpani])
	case *Transpose:
		stage.OnAfterTransposeReadCallback = any(callback).(OnAfterReadInterface[Transpose])
	case *Tremolo:
		stage.OnAfterTremoloReadCallback = any(callback).(OnAfterReadInterface[Tremolo])
	case *Tuplet:
		stage.OnAfterTupletReadCallback = any(callback).(OnAfterReadInterface[Tuplet])
	case *Tuplet_dot:
		stage.OnAfterTuplet_dotReadCallback = any(callback).(OnAfterReadInterface[Tuplet_dot])
	case *Tuplet_number:
		stage.OnAfterTuplet_numberReadCallback = any(callback).(OnAfterReadInterface[Tuplet_number])
	case *Tuplet_portion:
		stage.OnAfterTuplet_portionReadCallback = any(callback).(OnAfterReadInterface[Tuplet_portion])
	case *Tuplet_type:
		stage.OnAfterTuplet_typeReadCallback = any(callback).(OnAfterReadInterface[Tuplet_type])
	case *Typed_text:
		stage.OnAfterTyped_textReadCallback = any(callback).(OnAfterReadInterface[Typed_text])
	case *Unpitched:
		stage.OnAfterUnpitchedReadCallback = any(callback).(OnAfterReadInterface[Unpitched])
	case *Virtual_instrument:
		stage.OnAfterVirtual_instrumentReadCallback = any(callback).(OnAfterReadInterface[Virtual_instrument])
	case *Wait:
		stage.OnAfterWaitReadCallback = any(callback).(OnAfterReadInterface[Wait])
	case *Wavy_line:
		stage.OnAfterWavy_lineReadCallback = any(callback).(OnAfterReadInterface[Wavy_line])
	case *Wedge:
		stage.OnAfterWedgeReadCallback = any(callback).(OnAfterReadInterface[Wedge])
	case *Wood:
		stage.OnAfterWoodReadCallback = any(callback).(OnAfterReadInterface[Wood])
	case *Work:
		stage.OnAfterWorkReadCallback = any(callback).(OnAfterReadInterface[Work])
	}
}
