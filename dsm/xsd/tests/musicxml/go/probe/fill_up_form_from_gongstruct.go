// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/musicxml/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.A_directive:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "A_directive",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_directiveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_measure:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "A_measure",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_measureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_measure_1:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "A_measure_1",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_measure_1FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_part:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "A_part",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_partFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.A_part_1:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "A_part_1",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_part_1FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Accidental:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Accidental",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AccidentalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Accidental_mark:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Accidental_mark",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accidental_markFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Accidental_text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Accidental_text",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accidental_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Accord:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Accord",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AccordFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Accordion_registration:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Accordion_registration",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accordion_registrationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Appearance:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Appearance",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AppearanceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Arpeggiate:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Arpeggiate",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArpeggiateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Arrow:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Arrow",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArrowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Articulations:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Articulations",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArticulationsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Assess:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Assess",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AssessFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Attributes:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Attributes",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Backup:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Backup",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BackupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bar_style_color:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Bar_style_color",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Bar_style_colorFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Barline:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Barline",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarlineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Barre:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Barre",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarreFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bass:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Bass",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BassFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bass_step:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Bass_step",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Bass_stepFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Beam:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Beam",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BeamFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Beat_repeat:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Beat_repeat",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Beat_repeatFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Beat_unit_tied:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Beat_unit_tied",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Beat_unit_tiedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Beater:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Beater",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BeaterFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bend:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Bend",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BendFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bookmark:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Bookmark",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BookmarkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bracket:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Bracket",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BracketFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Breath_mark:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Breath_mark",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Breath_markFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Caesura:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Caesura",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CaesuraFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Cancel:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Cancel",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CancelFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Clef:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Clef",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ClefFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Coda:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Coda",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CodaFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Credit:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Credit",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CreditFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Dashes:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Dashes",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DashesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Defaults:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Defaults",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DefaultsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Degree:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Degree",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DegreeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Degree_alter:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Degree_alter",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_alterFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Degree_type:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Degree_type",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Degree_value:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Degree_value",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_valueFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Direction:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Direction",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DirectionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Direction_type:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Direction_type",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Direction_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Distance:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Distance",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DistanceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Double:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Double",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DoubleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Dynamics:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Dynamics",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DynamicsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Effect:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Effect",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EffectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Elision:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Elision",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElisionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EmptyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_font:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty_font",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_fontFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_line:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty_line",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_lineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_placement:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty_placement",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_placementFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_placement_smufl:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty_placement_smufl",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_placement_smuflFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_print_object_style_align:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty_print_object_style_align",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_object_style_alignFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_print_style:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty_print_style",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_styleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_print_style_align:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty_print_style_align",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_style_alignFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_print_style_align_id:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty_print_style_align_id",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_style_align_idFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_trill_sound:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Empty_trill_sound",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_trill_soundFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Encoding:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Encoding",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EncodingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Ending:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Ending",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Extend:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Extend",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExtendFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Feature:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Feature",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FeatureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Fermata:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Fermata",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FermataFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Figure:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Figure",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FigureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Figured_bass:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Figured_bass",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Figured_bassFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Fingering:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Fingering",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FingeringFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.First_fret:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "First_fret",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__First_fretFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.For_part:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "For_part",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__For_partFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Formatted_symbol:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Formatted_symbol",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_symbolFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Formatted_symbol_id:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Formatted_symbol_id",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_symbol_idFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Formatted_text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Formatted_text",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Formatted_text_id:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Formatted_text_id",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_text_idFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Forward:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Forward",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ForwardFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Frame:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Frame",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FrameFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Frame_note:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Frame_note",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Frame_noteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Fret:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Fret",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FretFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Glass:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Glass",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlassFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Glissando:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Glissando",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlissandoFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Glyph:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Glyph",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlyphFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Grace:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Grace",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GraceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group_barline:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Group_barline",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_barlineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group_name:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Group_name",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_nameFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group_symbol:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Group_symbol",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_symbolFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Grouping:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Grouping",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Hammer_on_pull_off:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Hammer_on_pull_off",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Hammer_on_pull_offFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Handbell:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Handbell",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HandbellFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmon_closed:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Harmon_closed",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmon_closedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmon_mute:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Harmon_mute",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmon_muteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmonic:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Harmonic",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HarmonicFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmony:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Harmony",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HarmonyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmony_alter:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Harmony_alter",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmony_alterFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harp_pedals:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Harp_pedals",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harp_pedalsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Heel_toe:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Heel_toe",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Heel_toeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Hole:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Hole",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HoleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Hole_closed:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Hole_closed",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Hole_closedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Horizontal_turn:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Horizontal_turn",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Horizontal_turnFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Identification:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Identification",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__IdentificationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Image:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Image",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ImageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Instrument:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Instrument",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InstrumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Instrument_change:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Instrument_change",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Instrument_changeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Instrument_link:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Instrument_link",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Instrument_linkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Interchangeable:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Interchangeable",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InterchangeableFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Inversion:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Inversion",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InversionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Key:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Key",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KeyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Key_accidental:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Key_accidental",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Key_accidentalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Key_octave:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Key_octave",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Key_octaveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Kind:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Kind",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KindFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Level:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Level",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LevelFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Line_detail:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Line_detail",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Line_detailFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Line_width:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Line_width",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Line_widthFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Link:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Link",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Listen:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Listen",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ListenFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Listening:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Listening",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ListeningFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lyric:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Lyric",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LyricFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lyric_font:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Lyric_font",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_fontFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lyric_language:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Lyric_language",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_languageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Measure_layout:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Measure_layout",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_layoutFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Measure_numbering:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Measure_numbering",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_numberingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Measure_repeat:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Measure_repeat",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_repeatFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Measure_style:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Measure_style",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_styleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Membrane:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Membrane",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MembraneFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metal:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Metal",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Metronome",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetronomeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome_beam:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Metronome_beam",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_beamFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome_note:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Metronome_note",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_noteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome_tied:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Metronome_tied",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_tiedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome_tuplet:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Metronome_tuplet",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_tupletFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Midi_device:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Midi_device",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Midi_deviceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Midi_instrument:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Midi_instrument",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Midi_instrumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Miscellaneous:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Miscellaneous",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MiscellaneousFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Miscellaneous_field:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Miscellaneous_field",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Miscellaneous_fieldFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Mordent:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Mordent",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MordentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Multiple_rest:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Multiple_rest",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Multiple_restFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Name_display:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Name_display",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Name_displayFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Non_arpeggiate:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Non_arpeggiate",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Non_arpeggiateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Notations:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Notations",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NotationsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Note:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Note",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Note_size:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Note_size",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Note_sizeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Note_type:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Note_type",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Note_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Notehead:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Notehead",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteheadFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Notehead_text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Notehead_text",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Notehead_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Numeral:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Numeral",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NumeralFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Numeral_key:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Numeral_key",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Numeral_keyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Numeral_root:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Numeral_root",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Numeral_rootFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Octave_shift:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Octave_shift",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Octave_shiftFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Offset:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Offset",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OffsetFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Opus:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Opus",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OpusFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Ornaments:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Ornaments",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OrnamentsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_appearance:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Other_appearance",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_appearanceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_direction:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Other_direction",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_directionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_listening:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Other_listening",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_listeningFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_notation:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Other_notation",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_notationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_placement_text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Other_placement_text",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_placement_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_play:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Other_play",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_playFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Other_text",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Page_layout:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Page_layout",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Page_layoutFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Page_margins:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Page_margins",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Page_marginsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_clef:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Part_clef",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_clefFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_group:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Part_group",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_groupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_link:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Part_link",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_linkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_list:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Part_list",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_listFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_name:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Part_name",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_nameFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_symbol:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Part_symbol",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_symbolFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_transpose:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Part_transpose",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_transposeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pedal:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Pedal",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PedalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pedal_tuning:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Pedal_tuning",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Pedal_tuningFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Per_minute:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Per_minute",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Per_minuteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Percussion:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Percussion",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PercussionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pitch:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Pitch",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PitchFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pitched:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Pitched",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PitchedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Placement_text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Placement_text",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Placement_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Play:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Play",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlayFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Player:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Player",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlayerFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Principal_voice:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Principal_voice",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Principal_voiceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Print:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Print",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PrintFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Release:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Release",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ReleaseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Repeat:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Repeat",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RepeatFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Rest:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Rest",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RestFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Root:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Root",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RootFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Root_step:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Root_step",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Root_stepFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Scaling:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Scaling",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScalingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Scordatura:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Scordatura",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScordaturaFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Score_instrument:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Score_instrument",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_instrumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Score_part:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Score_part",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_partFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Score_partwise:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Score_partwise",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_partwiseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Score_timewise:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Score_timewise",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_timewiseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Segno:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Segno",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SegnoFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Slash:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Slash",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlashFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Slide:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Slide",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlideFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Slur:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Slur",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlurFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Sound:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Sound",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SoundFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_details:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Staff_details",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_detailsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_divide:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Staff_divide",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_divideFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_layout:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Staff_layout",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_layoutFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_size:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Staff_size",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_sizeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_tuning:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Staff_tuning",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_tuningFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Stem:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Stem",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StemFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Stick:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Stick",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StickFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.String_mute:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "String_mute",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__String_muteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.String_type:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "String_type",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__String_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Strong_accent:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Strong_accent",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Strong_accentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Style_text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Style_text",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Style_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Supports:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Supports",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SupportsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Swing:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Swing",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SwingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Sync:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Sync",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SyncFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.System_dividers:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "System_dividers",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_dividersFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.System_layout:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "System_layout",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_layoutFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.System_margins:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "System_margins",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_marginsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tap:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tap",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TapFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Technical:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Technical",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TechnicalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Text_element_data:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Text_element_data",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Text_element_dataFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tie:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tie",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TieFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tied:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tied",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TiedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Time:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Time",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TimeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Time_modification:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Time_modification",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Time_modificationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Timpani:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Timpani",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TimpaniFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Transpose:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Transpose",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TransposeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tremolo:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tremolo",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TremoloFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tuplet",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TupletFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet_dot:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tuplet_dot",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_dotFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet_number:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tuplet_number",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_numberFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet_portion:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tuplet_portion",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_portionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet_type:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tuplet_type",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Typed_text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Typed_text",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Typed_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Unpitched:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Unpitched",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UnpitchedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Virtual_instrument:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Virtual_instrument",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Virtual_instrumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Wait:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Wait",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WaitFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Wavy_line:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Wavy_line",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Wavy_lineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Wedge:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Wedge",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WedgeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Wood:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Wood",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WoodFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Work:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Work",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WorkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
