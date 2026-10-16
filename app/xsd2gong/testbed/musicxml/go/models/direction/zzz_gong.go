// generated code - do not edit
package direction

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
)

// can be used for
//
//	days := __gong__abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __gong__abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	_ = __gong__abs
	_ = strings.Clone("")
)

type GongMarshallingMode = gong_runtime.GongMarshallingMode

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal = gong_runtime.GongMarshallingNormal

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit = gong_runtime.GongMarshallingAppendCommit
)

type gongStageNavigationMode = gong_runtime.StageNavigationMode

const (
	GongNavigationModeNormal     = gong_runtime.GongNavigationModeNormal
	GongNavigationModeNavigating = gong_runtime.GongNavigationModeNavigating
)

const (
	GongProbeTreeSidebarSuffix           = gong_runtime.GongProbeTreeSidebarSuffix
	GongProbeNavigationTreeSidebarSuffix = gong_runtime.GongProbeNavigationTreeSidebarSuffix
	GongProbeTableSuffix                 = gong_runtime.GongProbeTableSuffix
	GongProbeNotificationTableSuffix     = gong_runtime.GongProbeNotificationTableSuffix
	GongProbeFormSuffix                  = gong_runtime.GongProbeFormSuffix
	GongProbeSplitSuffix                 = gong_runtime.GongProbeSplitSuffix
	GongProbeLoadSuffix                  = gong_runtime.GongProbeLoadSuffix
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeSplitSuffix
}

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeLoadSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var _ = fmt.Sprintf

// idem for math package when not need by generated code
var _ = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

// MetaPackageImport represents a package import needed by a meta/diagram file
type MetaPackageImport struct {
	Alias string
	Path  string
}

// Stage enables storage of staged instances
type Stage struct {
	gong_runtime.StageCore

	// insertion point for definition of arrays registering instances
	Accords                map[*Accord]struct{}
	Accords_instance       map[*Accord]*Accord
	Accords_mapString      map[string]*Accord
	AccordOrder            uint
	Accord_stagedOrder     map[*Accord]uint
	Accord_orderStaged     map[uint]*Accord
	Accords_reference      map[*Accord]*Accord
	Accords_referenceOrder map[*Accord]uint

	// insertion point for slice of pointers maps
	OnAfterAccordCreateCallback GongOnAfterCreateInterface[Accord]
	OnAfterAccordUpdateCallback GongOnAfterUpdateInterface[Accord]
	OnAfterAccordDeleteCallback GongOnAfterDeleteInterface[Accord]

	AccordionRegistrations                map[*AccordionRegistration]struct{}
	AccordionRegistrations_instance       map[*AccordionRegistration]*AccordionRegistration
	AccordionRegistrations_mapString      map[string]*AccordionRegistration
	AccordionRegistrationOrder            uint
	AccordionRegistration_stagedOrder     map[*AccordionRegistration]uint
	AccordionRegistration_orderStaged     map[uint]*AccordionRegistration
	AccordionRegistrations_reference      map[*AccordionRegistration]*AccordionRegistration
	AccordionRegistrations_referenceOrder map[*AccordionRegistration]uint

	// insertion point for slice of pointers maps
	OnAfterAccordionRegistrationCreateCallback GongOnAfterCreateInterface[AccordionRegistration]
	OnAfterAccordionRegistrationUpdateCallback GongOnAfterUpdateInterface[AccordionRegistration]
	OnAfterAccordionRegistrationDeleteCallback GongOnAfterDeleteInterface[AccordionRegistration]

	Barres                map[*Barre]struct{}
	Barres_instance       map[*Barre]*Barre
	Barres_mapString      map[string]*Barre
	BarreOrder            uint
	Barre_stagedOrder     map[*Barre]uint
	Barre_orderStaged     map[uint]*Barre
	Barres_reference      map[*Barre]*Barre
	Barres_referenceOrder map[*Barre]uint

	// insertion point for slice of pointers maps
	OnAfterBarreCreateCallback GongOnAfterCreateInterface[Barre]
	OnAfterBarreUpdateCallback GongOnAfterUpdateInterface[Barre]
	OnAfterBarreDeleteCallback GongOnAfterDeleteInterface[Barre]

	Basss                map[*Bass]struct{}
	Basss_instance       map[*Bass]*Bass
	Basss_mapString      map[string]*Bass
	BassOrder            uint
	Bass_stagedOrder     map[*Bass]uint
	Bass_orderStaged     map[uint]*Bass
	Basss_reference      map[*Bass]*Bass
	Basss_referenceOrder map[*Bass]uint

	// insertion point for slice of pointers maps
	OnAfterBassCreateCallback GongOnAfterCreateInterface[Bass]
	OnAfterBassUpdateCallback GongOnAfterUpdateInterface[Bass]
	OnAfterBassDeleteCallback GongOnAfterDeleteInterface[Bass]

	BassSteps                map[*BassStep]struct{}
	BassSteps_instance       map[*BassStep]*BassStep
	BassSteps_mapString      map[string]*BassStep
	BassStepOrder            uint
	BassStep_stagedOrder     map[*BassStep]uint
	BassStep_orderStaged     map[uint]*BassStep
	BassSteps_reference      map[*BassStep]*BassStep
	BassSteps_referenceOrder map[*BassStep]uint

	// insertion point for slice of pointers maps
	OnAfterBassStepCreateCallback GongOnAfterCreateInterface[BassStep]
	OnAfterBassStepUpdateCallback GongOnAfterUpdateInterface[BassStep]
	OnAfterBassStepDeleteCallback GongOnAfterDeleteInterface[BassStep]

	BeatUnitTieds                map[*BeatUnitTied]struct{}
	BeatUnitTieds_instance       map[*BeatUnitTied]*BeatUnitTied
	BeatUnitTieds_mapString      map[string]*BeatUnitTied
	BeatUnitTiedOrder            uint
	BeatUnitTied_stagedOrder     map[*BeatUnitTied]uint
	BeatUnitTied_orderStaged     map[uint]*BeatUnitTied
	BeatUnitTieds_reference      map[*BeatUnitTied]*BeatUnitTied
	BeatUnitTieds_referenceOrder map[*BeatUnitTied]uint

	// insertion point for slice of pointers maps
	OnAfterBeatUnitTiedCreateCallback GongOnAfterCreateInterface[BeatUnitTied]
	OnAfterBeatUnitTiedUpdateCallback GongOnAfterUpdateInterface[BeatUnitTied]
	OnAfterBeatUnitTiedDeleteCallback GongOnAfterDeleteInterface[BeatUnitTied]

	Beaters                map[*Beater]struct{}
	Beaters_instance       map[*Beater]*Beater
	Beaters_mapString      map[string]*Beater
	BeaterOrder            uint
	Beater_stagedOrder     map[*Beater]uint
	Beater_orderStaged     map[uint]*Beater
	Beaters_reference      map[*Beater]*Beater
	Beaters_referenceOrder map[*Beater]uint

	// insertion point for slice of pointers maps
	OnAfterBeaterCreateCallback GongOnAfterCreateInterface[Beater]
	OnAfterBeaterUpdateCallback GongOnAfterUpdateInterface[Beater]
	OnAfterBeaterDeleteCallback GongOnAfterDeleteInterface[Beater]

	Brackets                map[*Bracket]struct{}
	Brackets_instance       map[*Bracket]*Bracket
	Brackets_mapString      map[string]*Bracket
	BracketOrder            uint
	Bracket_stagedOrder     map[*Bracket]uint
	Bracket_orderStaged     map[uint]*Bracket
	Brackets_reference      map[*Bracket]*Bracket
	Brackets_referenceOrder map[*Bracket]uint

	// insertion point for slice of pointers maps
	OnAfterBracketCreateCallback GongOnAfterCreateInterface[Bracket]
	OnAfterBracketUpdateCallback GongOnAfterUpdateInterface[Bracket]
	OnAfterBracketDeleteCallback GongOnAfterDeleteInterface[Bracket]

	Dashess                map[*Dashes]struct{}
	Dashess_instance       map[*Dashes]*Dashes
	Dashess_mapString      map[string]*Dashes
	DashesOrder            uint
	Dashes_stagedOrder     map[*Dashes]uint
	Dashes_orderStaged     map[uint]*Dashes
	Dashess_reference      map[*Dashes]*Dashes
	Dashess_referenceOrder map[*Dashes]uint

	// insertion point for slice of pointers maps
	OnAfterDashesCreateCallback GongOnAfterCreateInterface[Dashes]
	OnAfterDashesUpdateCallback GongOnAfterUpdateInterface[Dashes]
	OnAfterDashesDeleteCallback GongOnAfterDeleteInterface[Dashes]

	Degrees                map[*Degree]struct{}
	Degrees_instance       map[*Degree]*Degree
	Degrees_mapString      map[string]*Degree
	DegreeOrder            uint
	Degree_stagedOrder     map[*Degree]uint
	Degree_orderStaged     map[uint]*Degree
	Degrees_reference      map[*Degree]*Degree
	Degrees_referenceOrder map[*Degree]uint

	// insertion point for slice of pointers maps
	OnAfterDegreeCreateCallback GongOnAfterCreateInterface[Degree]
	OnAfterDegreeUpdateCallback GongOnAfterUpdateInterface[Degree]
	OnAfterDegreeDeleteCallback GongOnAfterDeleteInterface[Degree]

	DegreeAlters                map[*DegreeAlter]struct{}
	DegreeAlters_instance       map[*DegreeAlter]*DegreeAlter
	DegreeAlters_mapString      map[string]*DegreeAlter
	DegreeAlterOrder            uint
	DegreeAlter_stagedOrder     map[*DegreeAlter]uint
	DegreeAlter_orderStaged     map[uint]*DegreeAlter
	DegreeAlters_reference      map[*DegreeAlter]*DegreeAlter
	DegreeAlters_referenceOrder map[*DegreeAlter]uint

	// insertion point for slice of pointers maps
	OnAfterDegreeAlterCreateCallback GongOnAfterCreateInterface[DegreeAlter]
	OnAfterDegreeAlterUpdateCallback GongOnAfterUpdateInterface[DegreeAlter]
	OnAfterDegreeAlterDeleteCallback GongOnAfterDeleteInterface[DegreeAlter]

	DegreeTypes                map[*DegreeType]struct{}
	DegreeTypes_instance       map[*DegreeType]*DegreeType
	DegreeTypes_mapString      map[string]*DegreeType
	DegreeTypeOrder            uint
	DegreeType_stagedOrder     map[*DegreeType]uint
	DegreeType_orderStaged     map[uint]*DegreeType
	DegreeTypes_reference      map[*DegreeType]*DegreeType
	DegreeTypes_referenceOrder map[*DegreeType]uint

	// insertion point for slice of pointers maps
	OnAfterDegreeTypeCreateCallback GongOnAfterCreateInterface[DegreeType]
	OnAfterDegreeTypeUpdateCallback GongOnAfterUpdateInterface[DegreeType]
	OnAfterDegreeTypeDeleteCallback GongOnAfterDeleteInterface[DegreeType]

	DegreeValues                map[*DegreeValue]struct{}
	DegreeValues_instance       map[*DegreeValue]*DegreeValue
	DegreeValues_mapString      map[string]*DegreeValue
	DegreeValueOrder            uint
	DegreeValue_stagedOrder     map[*DegreeValue]uint
	DegreeValue_orderStaged     map[uint]*DegreeValue
	DegreeValues_reference      map[*DegreeValue]*DegreeValue
	DegreeValues_referenceOrder map[*DegreeValue]uint

	// insertion point for slice of pointers maps
	OnAfterDegreeValueCreateCallback GongOnAfterCreateInterface[DegreeValue]
	OnAfterDegreeValueUpdateCallback GongOnAfterUpdateInterface[DegreeValue]
	OnAfterDegreeValueDeleteCallback GongOnAfterDeleteInterface[DegreeValue]

	Directions                map[*Direction]struct{}
	Directions_instance       map[*Direction]*Direction
	Directions_mapString      map[string]*Direction
	DirectionOrder            uint
	Direction_stagedOrder     map[*Direction]uint
	Direction_orderStaged     map[uint]*Direction
	Directions_reference      map[*Direction]*Direction
	Directions_referenceOrder map[*Direction]uint

	// insertion point for slice of pointers maps
	Direction_DirectionType_reverseMap map[*DirectionType]*Direction

	OnAfterDirectionCreateCallback GongOnAfterCreateInterface[Direction]
	OnAfterDirectionUpdateCallback GongOnAfterUpdateInterface[Direction]
	OnAfterDirectionDeleteCallback GongOnAfterDeleteInterface[Direction]

	DirectionTypes                map[*DirectionType]struct{}
	DirectionTypes_instance       map[*DirectionType]*DirectionType
	DirectionTypes_mapString      map[string]*DirectionType
	DirectionTypeOrder            uint
	DirectionType_stagedOrder     map[*DirectionType]uint
	DirectionType_orderStaged     map[uint]*DirectionType
	DirectionTypes_reference      map[*DirectionType]*DirectionType
	DirectionTypes_referenceOrder map[*DirectionType]uint

	// insertion point for slice of pointers maps
	DirectionType_Percussion_reverseMap map[*Percussion]*DirectionType

	OnAfterDirectionTypeCreateCallback GongOnAfterCreateInterface[DirectionType]
	OnAfterDirectionTypeUpdateCallback GongOnAfterUpdateInterface[DirectionType]
	OnAfterDirectionTypeDeleteCallback GongOnAfterDeleteInterface[DirectionType]

	Effects                map[*Effect]struct{}
	Effects_instance       map[*Effect]*Effect
	Effects_mapString      map[string]*Effect
	EffectOrder            uint
	Effect_stagedOrder     map[*Effect]uint
	Effect_orderStaged     map[uint]*Effect
	Effects_reference      map[*Effect]*Effect
	Effects_referenceOrder map[*Effect]uint

	// insertion point for slice of pointers maps
	OnAfterEffectCreateCallback GongOnAfterCreateInterface[Effect]
	OnAfterEffectUpdateCallback GongOnAfterUpdateInterface[Effect]
	OnAfterEffectDeleteCallback GongOnAfterDeleteInterface[Effect]

	Features                map[*Feature]struct{}
	Features_instance       map[*Feature]*Feature
	Features_mapString      map[string]*Feature
	FeatureOrder            uint
	Feature_stagedOrder     map[*Feature]uint
	Feature_orderStaged     map[uint]*Feature
	Features_reference      map[*Feature]*Feature
	Features_referenceOrder map[*Feature]uint

	// insertion point for slice of pointers maps
	OnAfterFeatureCreateCallback GongOnAfterCreateInterface[Feature]
	OnAfterFeatureUpdateCallback GongOnAfterUpdateInterface[Feature]
	OnAfterFeatureDeleteCallback GongOnAfterDeleteInterface[Feature]

	FirstFrets                map[*FirstFret]struct{}
	FirstFrets_instance       map[*FirstFret]*FirstFret
	FirstFrets_mapString      map[string]*FirstFret
	FirstFretOrder            uint
	FirstFret_stagedOrder     map[*FirstFret]uint
	FirstFret_orderStaged     map[uint]*FirstFret
	FirstFrets_reference      map[*FirstFret]*FirstFret
	FirstFrets_referenceOrder map[*FirstFret]uint

	// insertion point for slice of pointers maps
	OnAfterFirstFretCreateCallback GongOnAfterCreateInterface[FirstFret]
	OnAfterFirstFretUpdateCallback GongOnAfterUpdateInterface[FirstFret]
	OnAfterFirstFretDeleteCallback GongOnAfterDeleteInterface[FirstFret]

	Frames                map[*Frame]struct{}
	Frames_instance       map[*Frame]*Frame
	Frames_mapString      map[string]*Frame
	FrameOrder            uint
	Frame_stagedOrder     map[*Frame]uint
	Frame_orderStaged     map[uint]*Frame
	Frames_reference      map[*Frame]*Frame
	Frames_referenceOrder map[*Frame]uint

	// insertion point for slice of pointers maps
	Frame_FrameNote_reverseMap map[*FrameNote]*Frame

	OnAfterFrameCreateCallback GongOnAfterCreateInterface[Frame]
	OnAfterFrameUpdateCallback GongOnAfterUpdateInterface[Frame]
	OnAfterFrameDeleteCallback GongOnAfterDeleteInterface[Frame]

	FrameNotes                map[*FrameNote]struct{}
	FrameNotes_instance       map[*FrameNote]*FrameNote
	FrameNotes_mapString      map[string]*FrameNote
	FrameNoteOrder            uint
	FrameNote_stagedOrder     map[*FrameNote]uint
	FrameNote_orderStaged     map[uint]*FrameNote
	FrameNotes_reference      map[*FrameNote]*FrameNote
	FrameNotes_referenceOrder map[*FrameNote]uint

	// insertion point for slice of pointers maps
	OnAfterFrameNoteCreateCallback GongOnAfterCreateInterface[FrameNote]
	OnAfterFrameNoteUpdateCallback GongOnAfterUpdateInterface[FrameNote]
	OnAfterFrameNoteDeleteCallback GongOnAfterDeleteInterface[FrameNote]

	Glasss                map[*Glass]struct{}
	Glasss_instance       map[*Glass]*Glass
	Glasss_mapString      map[string]*Glass
	GlassOrder            uint
	Glass_stagedOrder     map[*Glass]uint
	Glass_orderStaged     map[uint]*Glass
	Glasss_reference      map[*Glass]*Glass
	Glasss_referenceOrder map[*Glass]uint

	// insertion point for slice of pointers maps
	OnAfterGlassCreateCallback GongOnAfterCreateInterface[Glass]
	OnAfterGlassUpdateCallback GongOnAfterUpdateInterface[Glass]
	OnAfterGlassDeleteCallback GongOnAfterDeleteInterface[Glass]

	Groupings                map[*Grouping]struct{}
	Groupings_instance       map[*Grouping]*Grouping
	Groupings_mapString      map[string]*Grouping
	GroupingOrder            uint
	Grouping_stagedOrder     map[*Grouping]uint
	Grouping_orderStaged     map[uint]*Grouping
	Groupings_reference      map[*Grouping]*Grouping
	Groupings_referenceOrder map[*Grouping]uint

	// insertion point for slice of pointers maps
	Grouping_Feature_reverseMap map[*Feature]*Grouping

	OnAfterGroupingCreateCallback GongOnAfterCreateInterface[Grouping]
	OnAfterGroupingUpdateCallback GongOnAfterUpdateInterface[Grouping]
	OnAfterGroupingDeleteCallback GongOnAfterDeleteInterface[Grouping]

	Harmonys                map[*Harmony]struct{}
	Harmonys_instance       map[*Harmony]*Harmony
	Harmonys_mapString      map[string]*Harmony
	HarmonyOrder            uint
	Harmony_stagedOrder     map[*Harmony]uint
	Harmony_orderStaged     map[uint]*Harmony
	Harmonys_reference      map[*Harmony]*Harmony
	Harmonys_referenceOrder map[*Harmony]uint

	// insertion point for slice of pointers maps
	Harmony_Root_reverseMap map[*Root]*Harmony

	Harmony_Numeral_reverseMap map[*Numeral]*Harmony

	Harmony_Function_reverseMap map[*StyleText]*Harmony

	Harmony_Kind_reverseMap map[*Kind]*Harmony

	Harmony_Inversion_reverseMap map[*Inversion]*Harmony

	Harmony_Bass_reverseMap map[*Bass]*Harmony

	Harmony_Degree_reverseMap map[*Degree]*Harmony

	OnAfterHarmonyCreateCallback GongOnAfterCreateInterface[Harmony]
	OnAfterHarmonyUpdateCallback GongOnAfterUpdateInterface[Harmony]
	OnAfterHarmonyDeleteCallback GongOnAfterDeleteInterface[Harmony]

	HarmonyAlters                map[*HarmonyAlter]struct{}
	HarmonyAlters_instance       map[*HarmonyAlter]*HarmonyAlter
	HarmonyAlters_mapString      map[string]*HarmonyAlter
	HarmonyAlterOrder            uint
	HarmonyAlter_stagedOrder     map[*HarmonyAlter]uint
	HarmonyAlter_orderStaged     map[uint]*HarmonyAlter
	HarmonyAlters_reference      map[*HarmonyAlter]*HarmonyAlter
	HarmonyAlters_referenceOrder map[*HarmonyAlter]uint

	// insertion point for slice of pointers maps
	OnAfterHarmonyAlterCreateCallback GongOnAfterCreateInterface[HarmonyAlter]
	OnAfterHarmonyAlterUpdateCallback GongOnAfterUpdateInterface[HarmonyAlter]
	OnAfterHarmonyAlterDeleteCallback GongOnAfterDeleteInterface[HarmonyAlter]

	HarpPedalss                map[*HarpPedals]struct{}
	HarpPedalss_instance       map[*HarpPedals]*HarpPedals
	HarpPedalss_mapString      map[string]*HarpPedals
	HarpPedalsOrder            uint
	HarpPedals_stagedOrder     map[*HarpPedals]uint
	HarpPedals_orderStaged     map[uint]*HarpPedals
	HarpPedalss_reference      map[*HarpPedals]*HarpPedals
	HarpPedalss_referenceOrder map[*HarpPedals]uint

	// insertion point for slice of pointers maps
	HarpPedals_PedalTuning_reverseMap map[*PedalTuning]*HarpPedals

	OnAfterHarpPedalsCreateCallback GongOnAfterCreateInterface[HarpPedals]
	OnAfterHarpPedalsUpdateCallback GongOnAfterUpdateInterface[HarpPedals]
	OnAfterHarpPedalsDeleteCallback GongOnAfterDeleteInterface[HarpPedals]

	Images                map[*Image]struct{}
	Images_instance       map[*Image]*Image
	Images_mapString      map[string]*Image
	ImageOrder            uint
	Image_stagedOrder     map[*Image]uint
	Image_orderStaged     map[uint]*Image
	Images_reference      map[*Image]*Image
	Images_referenceOrder map[*Image]uint

	// insertion point for slice of pointers maps
	OnAfterImageCreateCallback GongOnAfterCreateInterface[Image]
	OnAfterImageUpdateCallback GongOnAfterUpdateInterface[Image]
	OnAfterImageDeleteCallback GongOnAfterDeleteInterface[Image]

	InstrumentChanges                map[*InstrumentChange]struct{}
	InstrumentChanges_instance       map[*InstrumentChange]*InstrumentChange
	InstrumentChanges_mapString      map[string]*InstrumentChange
	InstrumentChangeOrder            uint
	InstrumentChange_stagedOrder     map[*InstrumentChange]uint
	InstrumentChange_orderStaged     map[uint]*InstrumentChange
	InstrumentChanges_reference      map[*InstrumentChange]*InstrumentChange
	InstrumentChanges_referenceOrder map[*InstrumentChange]uint

	// insertion point for slice of pointers maps
	OnAfterInstrumentChangeCreateCallback GongOnAfterCreateInterface[InstrumentChange]
	OnAfterInstrumentChangeUpdateCallback GongOnAfterUpdateInterface[InstrumentChange]
	OnAfterInstrumentChangeDeleteCallback GongOnAfterDeleteInterface[InstrumentChange]

	Inversions                map[*Inversion]struct{}
	Inversions_instance       map[*Inversion]*Inversion
	Inversions_mapString      map[string]*Inversion
	InversionOrder            uint
	Inversion_stagedOrder     map[*Inversion]uint
	Inversion_orderStaged     map[uint]*Inversion
	Inversions_reference      map[*Inversion]*Inversion
	Inversions_referenceOrder map[*Inversion]uint

	// insertion point for slice of pointers maps
	OnAfterInversionCreateCallback GongOnAfterCreateInterface[Inversion]
	OnAfterInversionUpdateCallback GongOnAfterUpdateInterface[Inversion]
	OnAfterInversionDeleteCallback GongOnAfterDeleteInterface[Inversion]

	Kinds                map[*Kind]struct{}
	Kinds_instance       map[*Kind]*Kind
	Kinds_mapString      map[string]*Kind
	KindOrder            uint
	Kind_stagedOrder     map[*Kind]uint
	Kind_orderStaged     map[uint]*Kind
	Kinds_reference      map[*Kind]*Kind
	Kinds_referenceOrder map[*Kind]uint

	// insertion point for slice of pointers maps
	OnAfterKindCreateCallback GongOnAfterCreateInterface[Kind]
	OnAfterKindUpdateCallback GongOnAfterUpdateInterface[Kind]
	OnAfterKindDeleteCallback GongOnAfterDeleteInterface[Kind]

	Listenings                map[*Listening]struct{}
	Listenings_instance       map[*Listening]*Listening
	Listenings_mapString      map[string]*Listening
	ListeningOrder            uint
	Listening_stagedOrder     map[*Listening]uint
	Listening_orderStaged     map[uint]*Listening
	Listenings_reference      map[*Listening]*Listening
	Listenings_referenceOrder map[*Listening]uint

	// insertion point for slice of pointers maps
	Listening_Sync_reverseMap map[*Sync]*Listening

	Listening_OtherListening_reverseMap map[*OtherListening]*Listening

	OnAfterListeningCreateCallback GongOnAfterCreateInterface[Listening]
	OnAfterListeningUpdateCallback GongOnAfterUpdateInterface[Listening]
	OnAfterListeningDeleteCallback GongOnAfterDeleteInterface[Listening]

	MeasureNumberings                map[*MeasureNumbering]struct{}
	MeasureNumberings_instance       map[*MeasureNumbering]*MeasureNumbering
	MeasureNumberings_mapString      map[string]*MeasureNumbering
	MeasureNumberingOrder            uint
	MeasureNumbering_stagedOrder     map[*MeasureNumbering]uint
	MeasureNumbering_orderStaged     map[uint]*MeasureNumbering
	MeasureNumberings_reference      map[*MeasureNumbering]*MeasureNumbering
	MeasureNumberings_referenceOrder map[*MeasureNumbering]uint

	// insertion point for slice of pointers maps
	OnAfterMeasureNumberingCreateCallback GongOnAfterCreateInterface[MeasureNumbering]
	OnAfterMeasureNumberingUpdateCallback GongOnAfterUpdateInterface[MeasureNumbering]
	OnAfterMeasureNumberingDeleteCallback GongOnAfterDeleteInterface[MeasureNumbering]

	Membranes                map[*Membrane]struct{}
	Membranes_instance       map[*Membrane]*Membrane
	Membranes_mapString      map[string]*Membrane
	MembraneOrder            uint
	Membrane_stagedOrder     map[*Membrane]uint
	Membrane_orderStaged     map[uint]*Membrane
	Membranes_reference      map[*Membrane]*Membrane
	Membranes_referenceOrder map[*Membrane]uint

	// insertion point for slice of pointers maps
	OnAfterMembraneCreateCallback GongOnAfterCreateInterface[Membrane]
	OnAfterMembraneUpdateCallback GongOnAfterUpdateInterface[Membrane]
	OnAfterMembraneDeleteCallback GongOnAfterDeleteInterface[Membrane]

	Metals                map[*Metal]struct{}
	Metals_instance       map[*Metal]*Metal
	Metals_mapString      map[string]*Metal
	MetalOrder            uint
	Metal_stagedOrder     map[*Metal]uint
	Metal_orderStaged     map[uint]*Metal
	Metals_reference      map[*Metal]*Metal
	Metals_referenceOrder map[*Metal]uint

	// insertion point for slice of pointers maps
	OnAfterMetalCreateCallback GongOnAfterCreateInterface[Metal]
	OnAfterMetalUpdateCallback GongOnAfterUpdateInterface[Metal]
	OnAfterMetalDeleteCallback GongOnAfterDeleteInterface[Metal]

	Metronomes                map[*Metronome]struct{}
	Metronomes_instance       map[*Metronome]*Metronome
	Metronomes_mapString      map[string]*Metronome
	MetronomeOrder            uint
	Metronome_stagedOrder     map[*Metronome]uint
	Metronome_orderStaged     map[uint]*Metronome
	Metronomes_reference      map[*Metronome]*Metronome
	Metronomes_referenceOrder map[*Metronome]uint

	// insertion point for slice of pointers maps
	Metronome_BeatUnitTied_reverseMap map[*BeatUnitTied]*Metronome

	Metronome_MetronomeNote_reverseMap map[*MetronomeNote]*Metronome

	OnAfterMetronomeCreateCallback GongOnAfterCreateInterface[Metronome]
	OnAfterMetronomeUpdateCallback GongOnAfterUpdateInterface[Metronome]
	OnAfterMetronomeDeleteCallback GongOnAfterDeleteInterface[Metronome]

	MetronomeBeams                map[*MetronomeBeam]struct{}
	MetronomeBeams_instance       map[*MetronomeBeam]*MetronomeBeam
	MetronomeBeams_mapString      map[string]*MetronomeBeam
	MetronomeBeamOrder            uint
	MetronomeBeam_stagedOrder     map[*MetronomeBeam]uint
	MetronomeBeam_orderStaged     map[uint]*MetronomeBeam
	MetronomeBeams_reference      map[*MetronomeBeam]*MetronomeBeam
	MetronomeBeams_referenceOrder map[*MetronomeBeam]uint

	// insertion point for slice of pointers maps
	OnAfterMetronomeBeamCreateCallback GongOnAfterCreateInterface[MetronomeBeam]
	OnAfterMetronomeBeamUpdateCallback GongOnAfterUpdateInterface[MetronomeBeam]
	OnAfterMetronomeBeamDeleteCallback GongOnAfterDeleteInterface[MetronomeBeam]

	MetronomeNotes                map[*MetronomeNote]struct{}
	MetronomeNotes_instance       map[*MetronomeNote]*MetronomeNote
	MetronomeNotes_mapString      map[string]*MetronomeNote
	MetronomeNoteOrder            uint
	MetronomeNote_stagedOrder     map[*MetronomeNote]uint
	MetronomeNote_orderStaged     map[uint]*MetronomeNote
	MetronomeNotes_reference      map[*MetronomeNote]*MetronomeNote
	MetronomeNotes_referenceOrder map[*MetronomeNote]uint

	// insertion point for slice of pointers maps
	MetronomeNote_MetronomeBeam_reverseMap map[*MetronomeBeam]*MetronomeNote

	OnAfterMetronomeNoteCreateCallback GongOnAfterCreateInterface[MetronomeNote]
	OnAfterMetronomeNoteUpdateCallback GongOnAfterUpdateInterface[MetronomeNote]
	OnAfterMetronomeNoteDeleteCallback GongOnAfterDeleteInterface[MetronomeNote]

	MetronomeTieds                map[*MetronomeTied]struct{}
	MetronomeTieds_instance       map[*MetronomeTied]*MetronomeTied
	MetronomeTieds_mapString      map[string]*MetronomeTied
	MetronomeTiedOrder            uint
	MetronomeTied_stagedOrder     map[*MetronomeTied]uint
	MetronomeTied_orderStaged     map[uint]*MetronomeTied
	MetronomeTieds_reference      map[*MetronomeTied]*MetronomeTied
	MetronomeTieds_referenceOrder map[*MetronomeTied]uint

	// insertion point for slice of pointers maps
	OnAfterMetronomeTiedCreateCallback GongOnAfterCreateInterface[MetronomeTied]
	OnAfterMetronomeTiedUpdateCallback GongOnAfterUpdateInterface[MetronomeTied]
	OnAfterMetronomeTiedDeleteCallback GongOnAfterDeleteInterface[MetronomeTied]

	MetronomeTuplets                map[*MetronomeTuplet]struct{}
	MetronomeTuplets_instance       map[*MetronomeTuplet]*MetronomeTuplet
	MetronomeTuplets_mapString      map[string]*MetronomeTuplet
	MetronomeTupletOrder            uint
	MetronomeTuplet_stagedOrder     map[*MetronomeTuplet]uint
	MetronomeTuplet_orderStaged     map[uint]*MetronomeTuplet
	MetronomeTuplets_reference      map[*MetronomeTuplet]*MetronomeTuplet
	MetronomeTuplets_referenceOrder map[*MetronomeTuplet]uint

	// insertion point for slice of pointers maps
	OnAfterMetronomeTupletCreateCallback GongOnAfterCreateInterface[MetronomeTuplet]
	OnAfterMetronomeTupletUpdateCallback GongOnAfterUpdateInterface[MetronomeTuplet]
	OnAfterMetronomeTupletDeleteCallback GongOnAfterDeleteInterface[MetronomeTuplet]

	Numerals                map[*Numeral]struct{}
	Numerals_instance       map[*Numeral]*Numeral
	Numerals_mapString      map[string]*Numeral
	NumeralOrder            uint
	Numeral_stagedOrder     map[*Numeral]uint
	Numeral_orderStaged     map[uint]*Numeral
	Numerals_reference      map[*Numeral]*Numeral
	Numerals_referenceOrder map[*Numeral]uint

	// insertion point for slice of pointers maps
	OnAfterNumeralCreateCallback GongOnAfterCreateInterface[Numeral]
	OnAfterNumeralUpdateCallback GongOnAfterUpdateInterface[Numeral]
	OnAfterNumeralDeleteCallback GongOnAfterDeleteInterface[Numeral]

	NumeralKeys                map[*NumeralKey]struct{}
	NumeralKeys_instance       map[*NumeralKey]*NumeralKey
	NumeralKeys_mapString      map[string]*NumeralKey
	NumeralKeyOrder            uint
	NumeralKey_stagedOrder     map[*NumeralKey]uint
	NumeralKey_orderStaged     map[uint]*NumeralKey
	NumeralKeys_reference      map[*NumeralKey]*NumeralKey
	NumeralKeys_referenceOrder map[*NumeralKey]uint

	// insertion point for slice of pointers maps
	OnAfterNumeralKeyCreateCallback GongOnAfterCreateInterface[NumeralKey]
	OnAfterNumeralKeyUpdateCallback GongOnAfterUpdateInterface[NumeralKey]
	OnAfterNumeralKeyDeleteCallback GongOnAfterDeleteInterface[NumeralKey]

	NumeralRoots                map[*NumeralRoot]struct{}
	NumeralRoots_instance       map[*NumeralRoot]*NumeralRoot
	NumeralRoots_mapString      map[string]*NumeralRoot
	NumeralRootOrder            uint
	NumeralRoot_stagedOrder     map[*NumeralRoot]uint
	NumeralRoot_orderStaged     map[uint]*NumeralRoot
	NumeralRoots_reference      map[*NumeralRoot]*NumeralRoot
	NumeralRoots_referenceOrder map[*NumeralRoot]uint

	// insertion point for slice of pointers maps
	OnAfterNumeralRootCreateCallback GongOnAfterCreateInterface[NumeralRoot]
	OnAfterNumeralRootUpdateCallback GongOnAfterUpdateInterface[NumeralRoot]
	OnAfterNumeralRootDeleteCallback GongOnAfterDeleteInterface[NumeralRoot]

	OctaveShifts                map[*OctaveShift]struct{}
	OctaveShifts_instance       map[*OctaveShift]*OctaveShift
	OctaveShifts_mapString      map[string]*OctaveShift
	OctaveShiftOrder            uint
	OctaveShift_stagedOrder     map[*OctaveShift]uint
	OctaveShift_orderStaged     map[uint]*OctaveShift
	OctaveShifts_reference      map[*OctaveShift]*OctaveShift
	OctaveShifts_referenceOrder map[*OctaveShift]uint

	// insertion point for slice of pointers maps
	OnAfterOctaveShiftCreateCallback GongOnAfterCreateInterface[OctaveShift]
	OnAfterOctaveShiftUpdateCallback GongOnAfterUpdateInterface[OctaveShift]
	OnAfterOctaveShiftDeleteCallback GongOnAfterDeleteInterface[OctaveShift]

	Offsets                map[*Offset]struct{}
	Offsets_instance       map[*Offset]*Offset
	Offsets_mapString      map[string]*Offset
	OffsetOrder            uint
	Offset_stagedOrder     map[*Offset]uint
	Offset_orderStaged     map[uint]*Offset
	Offsets_reference      map[*Offset]*Offset
	Offsets_referenceOrder map[*Offset]uint

	// insertion point for slice of pointers maps
	OnAfterOffsetCreateCallback GongOnAfterCreateInterface[Offset]
	OnAfterOffsetUpdateCallback GongOnAfterUpdateInterface[Offset]
	OnAfterOffsetDeleteCallback GongOnAfterDeleteInterface[Offset]

	OtherDirections                map[*OtherDirection]struct{}
	OtherDirections_instance       map[*OtherDirection]*OtherDirection
	OtherDirections_mapString      map[string]*OtherDirection
	OtherDirectionOrder            uint
	OtherDirection_stagedOrder     map[*OtherDirection]uint
	OtherDirection_orderStaged     map[uint]*OtherDirection
	OtherDirections_reference      map[*OtherDirection]*OtherDirection
	OtherDirections_referenceOrder map[*OtherDirection]uint

	// insertion point for slice of pointers maps
	OnAfterOtherDirectionCreateCallback GongOnAfterCreateInterface[OtherDirection]
	OnAfterOtherDirectionUpdateCallback GongOnAfterUpdateInterface[OtherDirection]
	OnAfterOtherDirectionDeleteCallback GongOnAfterDeleteInterface[OtherDirection]

	OtherListenings                map[*OtherListening]struct{}
	OtherListenings_instance       map[*OtherListening]*OtherListening
	OtherListenings_mapString      map[string]*OtherListening
	OtherListeningOrder            uint
	OtherListening_stagedOrder     map[*OtherListening]uint
	OtherListening_orderStaged     map[uint]*OtherListening
	OtherListenings_reference      map[*OtherListening]*OtherListening
	OtherListenings_referenceOrder map[*OtherListening]uint

	// insertion point for slice of pointers maps
	OnAfterOtherListeningCreateCallback GongOnAfterCreateInterface[OtherListening]
	OnAfterOtherListeningUpdateCallback GongOnAfterUpdateInterface[OtherListening]
	OnAfterOtherListeningDeleteCallback GongOnAfterDeleteInterface[OtherListening]

	Pedals                map[*Pedal]struct{}
	Pedals_instance       map[*Pedal]*Pedal
	Pedals_mapString      map[string]*Pedal
	PedalOrder            uint
	Pedal_stagedOrder     map[*Pedal]uint
	Pedal_orderStaged     map[uint]*Pedal
	Pedals_reference      map[*Pedal]*Pedal
	Pedals_referenceOrder map[*Pedal]uint

	// insertion point for slice of pointers maps
	OnAfterPedalCreateCallback GongOnAfterCreateInterface[Pedal]
	OnAfterPedalUpdateCallback GongOnAfterUpdateInterface[Pedal]
	OnAfterPedalDeleteCallback GongOnAfterDeleteInterface[Pedal]

	PedalTunings                map[*PedalTuning]struct{}
	PedalTunings_instance       map[*PedalTuning]*PedalTuning
	PedalTunings_mapString      map[string]*PedalTuning
	PedalTuningOrder            uint
	PedalTuning_stagedOrder     map[*PedalTuning]uint
	PedalTuning_orderStaged     map[uint]*PedalTuning
	PedalTunings_reference      map[*PedalTuning]*PedalTuning
	PedalTunings_referenceOrder map[*PedalTuning]uint

	// insertion point for slice of pointers maps
	OnAfterPedalTuningCreateCallback GongOnAfterCreateInterface[PedalTuning]
	OnAfterPedalTuningUpdateCallback GongOnAfterUpdateInterface[PedalTuning]
	OnAfterPedalTuningDeleteCallback GongOnAfterDeleteInterface[PedalTuning]

	PerMinutes                map[*PerMinute]struct{}
	PerMinutes_instance       map[*PerMinute]*PerMinute
	PerMinutes_mapString      map[string]*PerMinute
	PerMinuteOrder            uint
	PerMinute_stagedOrder     map[*PerMinute]uint
	PerMinute_orderStaged     map[uint]*PerMinute
	PerMinutes_reference      map[*PerMinute]*PerMinute
	PerMinutes_referenceOrder map[*PerMinute]uint

	// insertion point for slice of pointers maps
	OnAfterPerMinuteCreateCallback GongOnAfterCreateInterface[PerMinute]
	OnAfterPerMinuteUpdateCallback GongOnAfterUpdateInterface[PerMinute]
	OnAfterPerMinuteDeleteCallback GongOnAfterDeleteInterface[PerMinute]

	Percussions                map[*Percussion]struct{}
	Percussions_instance       map[*Percussion]*Percussion
	Percussions_mapString      map[string]*Percussion
	PercussionOrder            uint
	Percussion_stagedOrder     map[*Percussion]uint
	Percussion_orderStaged     map[uint]*Percussion
	Percussions_reference      map[*Percussion]*Percussion
	Percussions_referenceOrder map[*Percussion]uint

	// insertion point for slice of pointers maps
	OnAfterPercussionCreateCallback GongOnAfterCreateInterface[Percussion]
	OnAfterPercussionUpdateCallback GongOnAfterUpdateInterface[Percussion]
	OnAfterPercussionDeleteCallback GongOnAfterDeleteInterface[Percussion]

	Pitcheds                map[*Pitched]struct{}
	Pitcheds_instance       map[*Pitched]*Pitched
	Pitcheds_mapString      map[string]*Pitched
	PitchedOrder            uint
	Pitched_stagedOrder     map[*Pitched]uint
	Pitched_orderStaged     map[uint]*Pitched
	Pitcheds_reference      map[*Pitched]*Pitched
	Pitcheds_referenceOrder map[*Pitched]uint

	// insertion point for slice of pointers maps
	OnAfterPitchedCreateCallback GongOnAfterCreateInterface[Pitched]
	OnAfterPitchedUpdateCallback GongOnAfterUpdateInterface[Pitched]
	OnAfterPitchedDeleteCallback GongOnAfterDeleteInterface[Pitched]

	PrincipalVoices                map[*PrincipalVoice]struct{}
	PrincipalVoices_instance       map[*PrincipalVoice]*PrincipalVoice
	PrincipalVoices_mapString      map[string]*PrincipalVoice
	PrincipalVoiceOrder            uint
	PrincipalVoice_stagedOrder     map[*PrincipalVoice]uint
	PrincipalVoice_orderStaged     map[uint]*PrincipalVoice
	PrincipalVoices_reference      map[*PrincipalVoice]*PrincipalVoice
	PrincipalVoices_referenceOrder map[*PrincipalVoice]uint

	// insertion point for slice of pointers maps
	OnAfterPrincipalVoiceCreateCallback GongOnAfterCreateInterface[PrincipalVoice]
	OnAfterPrincipalVoiceUpdateCallback GongOnAfterUpdateInterface[PrincipalVoice]
	OnAfterPrincipalVoiceDeleteCallback GongOnAfterDeleteInterface[PrincipalVoice]

	Prints                map[*Print]struct{}
	Prints_instance       map[*Print]*Print
	Prints_mapString      map[string]*Print
	PrintOrder            uint
	Print_stagedOrder     map[*Print]uint
	Print_orderStaged     map[uint]*Print
	Prints_reference      map[*Print]*Print
	Prints_referenceOrder map[*Print]uint

	// insertion point for slice of pointers maps
	OnAfterPrintCreateCallback GongOnAfterCreateInterface[Print]
	OnAfterPrintUpdateCallback GongOnAfterUpdateInterface[Print]
	OnAfterPrintDeleteCallback GongOnAfterDeleteInterface[Print]

	Roots                map[*Root]struct{}
	Roots_instance       map[*Root]*Root
	Roots_mapString      map[string]*Root
	RootOrder            uint
	Root_stagedOrder     map[*Root]uint
	Root_orderStaged     map[uint]*Root
	Roots_reference      map[*Root]*Root
	Roots_referenceOrder map[*Root]uint

	// insertion point for slice of pointers maps
	OnAfterRootCreateCallback GongOnAfterCreateInterface[Root]
	OnAfterRootUpdateCallback GongOnAfterUpdateInterface[Root]
	OnAfterRootDeleteCallback GongOnAfterDeleteInterface[Root]

	RootSteps                map[*RootStep]struct{}
	RootSteps_instance       map[*RootStep]*RootStep
	RootSteps_mapString      map[string]*RootStep
	RootStepOrder            uint
	RootStep_stagedOrder     map[*RootStep]uint
	RootStep_orderStaged     map[uint]*RootStep
	RootSteps_reference      map[*RootStep]*RootStep
	RootSteps_referenceOrder map[*RootStep]uint

	// insertion point for slice of pointers maps
	OnAfterRootStepCreateCallback GongOnAfterCreateInterface[RootStep]
	OnAfterRootStepUpdateCallback GongOnAfterUpdateInterface[RootStep]
	OnAfterRootStepDeleteCallback GongOnAfterDeleteInterface[RootStep]

	Scordaturas                map[*Scordatura]struct{}
	Scordaturas_instance       map[*Scordatura]*Scordatura
	Scordaturas_mapString      map[string]*Scordatura
	ScordaturaOrder            uint
	Scordatura_stagedOrder     map[*Scordatura]uint
	Scordatura_orderStaged     map[uint]*Scordatura
	Scordaturas_reference      map[*Scordatura]*Scordatura
	Scordaturas_referenceOrder map[*Scordatura]uint

	// insertion point for slice of pointers maps
	Scordatura_Accord_reverseMap map[*Accord]*Scordatura

	OnAfterScordaturaCreateCallback GongOnAfterCreateInterface[Scordatura]
	OnAfterScordaturaUpdateCallback GongOnAfterUpdateInterface[Scordatura]
	OnAfterScordaturaDeleteCallback GongOnAfterDeleteInterface[Scordatura]

	Sounds                map[*Sound]struct{}
	Sounds_instance       map[*Sound]*Sound
	Sounds_mapString      map[string]*Sound
	SoundOrder            uint
	Sound_stagedOrder     map[*Sound]uint
	Sound_orderStaged     map[uint]*Sound
	Sounds_reference      map[*Sound]*Sound
	Sounds_referenceOrder map[*Sound]uint

	// insertion point for slice of pointers maps
	Sound_InstrumentChange_reverseMap map[*InstrumentChange]*Sound

	OnAfterSoundCreateCallback GongOnAfterCreateInterface[Sound]
	OnAfterSoundUpdateCallback GongOnAfterUpdateInterface[Sound]
	OnAfterSoundDeleteCallback GongOnAfterDeleteInterface[Sound]

	StaffDivides                map[*StaffDivide]struct{}
	StaffDivides_instance       map[*StaffDivide]*StaffDivide
	StaffDivides_mapString      map[string]*StaffDivide
	StaffDivideOrder            uint
	StaffDivide_stagedOrder     map[*StaffDivide]uint
	StaffDivide_orderStaged     map[uint]*StaffDivide
	StaffDivides_reference      map[*StaffDivide]*StaffDivide
	StaffDivides_referenceOrder map[*StaffDivide]uint

	// insertion point for slice of pointers maps
	OnAfterStaffDivideCreateCallback GongOnAfterCreateInterface[StaffDivide]
	OnAfterStaffDivideUpdateCallback GongOnAfterUpdateInterface[StaffDivide]
	OnAfterStaffDivideDeleteCallback GongOnAfterDeleteInterface[StaffDivide]

	Sticks                map[*Stick]struct{}
	Sticks_instance       map[*Stick]*Stick
	Sticks_mapString      map[string]*Stick
	StickOrder            uint
	Stick_stagedOrder     map[*Stick]uint
	Stick_orderStaged     map[uint]*Stick
	Sticks_reference      map[*Stick]*Stick
	Sticks_referenceOrder map[*Stick]uint

	// insertion point for slice of pointers maps
	OnAfterStickCreateCallback GongOnAfterCreateInterface[Stick]
	OnAfterStickUpdateCallback GongOnAfterUpdateInterface[Stick]
	OnAfterStickDeleteCallback GongOnAfterDeleteInterface[Stick]

	StringMutes                map[*StringMute]struct{}
	StringMutes_instance       map[*StringMute]*StringMute
	StringMutes_mapString      map[string]*StringMute
	StringMuteOrder            uint
	StringMute_stagedOrder     map[*StringMute]uint
	StringMute_orderStaged     map[uint]*StringMute
	StringMutes_reference      map[*StringMute]*StringMute
	StringMutes_referenceOrder map[*StringMute]uint

	// insertion point for slice of pointers maps
	OnAfterStringMuteCreateCallback GongOnAfterCreateInterface[StringMute]
	OnAfterStringMuteUpdateCallback GongOnAfterUpdateInterface[StringMute]
	OnAfterStringMuteDeleteCallback GongOnAfterDeleteInterface[StringMute]

	StyleTexts                map[*StyleText]struct{}
	StyleTexts_instance       map[*StyleText]*StyleText
	StyleTexts_mapString      map[string]*StyleText
	StyleTextOrder            uint
	StyleText_stagedOrder     map[*StyleText]uint
	StyleText_orderStaged     map[uint]*StyleText
	StyleTexts_reference      map[*StyleText]*StyleText
	StyleTexts_referenceOrder map[*StyleText]uint

	// insertion point for slice of pointers maps
	OnAfterStyleTextCreateCallback GongOnAfterCreateInterface[StyleText]
	OnAfterStyleTextUpdateCallback GongOnAfterUpdateInterface[StyleText]
	OnAfterStyleTextDeleteCallback GongOnAfterDeleteInterface[StyleText]

	Swings                map[*Swing]struct{}
	Swings_instance       map[*Swing]*Swing
	Swings_mapString      map[string]*Swing
	SwingOrder            uint
	Swing_stagedOrder     map[*Swing]uint
	Swing_orderStaged     map[uint]*Swing
	Swings_reference      map[*Swing]*Swing
	Swings_referenceOrder map[*Swing]uint

	// insertion point for slice of pointers maps
	OnAfterSwingCreateCallback GongOnAfterCreateInterface[Swing]
	OnAfterSwingUpdateCallback GongOnAfterUpdateInterface[Swing]
	OnAfterSwingDeleteCallback GongOnAfterDeleteInterface[Swing]

	Syncs                map[*Sync]struct{}
	Syncs_instance       map[*Sync]*Sync
	Syncs_mapString      map[string]*Sync
	SyncOrder            uint
	Sync_stagedOrder     map[*Sync]uint
	Sync_orderStaged     map[uint]*Sync
	Syncs_reference      map[*Sync]*Sync
	Syncs_referenceOrder map[*Sync]uint

	// insertion point for slice of pointers maps
	OnAfterSyncCreateCallback GongOnAfterCreateInterface[Sync]
	OnAfterSyncUpdateCallback GongOnAfterUpdateInterface[Sync]
	OnAfterSyncDeleteCallback GongOnAfterDeleteInterface[Sync]

	TimeModifications                map[*TimeModification]struct{}
	TimeModifications_instance       map[*TimeModification]*TimeModification
	TimeModifications_mapString      map[string]*TimeModification
	TimeModificationOrder            uint
	TimeModification_stagedOrder     map[*TimeModification]uint
	TimeModification_orderStaged     map[uint]*TimeModification
	TimeModifications_reference      map[*TimeModification]*TimeModification
	TimeModifications_referenceOrder map[*TimeModification]uint

	// insertion point for slice of pointers maps
	OnAfterTimeModificationCreateCallback GongOnAfterCreateInterface[TimeModification]
	OnAfterTimeModificationUpdateCallback GongOnAfterUpdateInterface[TimeModification]
	OnAfterTimeModificationDeleteCallback GongOnAfterDeleteInterface[TimeModification]

	Timpanis                map[*Timpani]struct{}
	Timpanis_instance       map[*Timpani]*Timpani
	Timpanis_mapString      map[string]*Timpani
	TimpaniOrder            uint
	Timpani_stagedOrder     map[*Timpani]uint
	Timpani_orderStaged     map[uint]*Timpani
	Timpanis_reference      map[*Timpani]*Timpani
	Timpanis_referenceOrder map[*Timpani]uint

	// insertion point for slice of pointers maps
	OnAfterTimpaniCreateCallback GongOnAfterCreateInterface[Timpani]
	OnAfterTimpaniUpdateCallback GongOnAfterUpdateInterface[Timpani]
	OnAfterTimpaniDeleteCallback GongOnAfterDeleteInterface[Timpani]

	VirtualInstruments                map[*VirtualInstrument]struct{}
	VirtualInstruments_instance       map[*VirtualInstrument]*VirtualInstrument
	VirtualInstruments_mapString      map[string]*VirtualInstrument
	VirtualInstrumentOrder            uint
	VirtualInstrument_stagedOrder     map[*VirtualInstrument]uint
	VirtualInstrument_orderStaged     map[uint]*VirtualInstrument
	VirtualInstruments_reference      map[*VirtualInstrument]*VirtualInstrument
	VirtualInstruments_referenceOrder map[*VirtualInstrument]uint

	// insertion point for slice of pointers maps
	OnAfterVirtualInstrumentCreateCallback GongOnAfterCreateInterface[VirtualInstrument]
	OnAfterVirtualInstrumentUpdateCallback GongOnAfterUpdateInterface[VirtualInstrument]
	OnAfterVirtualInstrumentDeleteCallback GongOnAfterDeleteInterface[VirtualInstrument]

	Wedges                map[*Wedge]struct{}
	Wedges_instance       map[*Wedge]*Wedge
	Wedges_mapString      map[string]*Wedge
	WedgeOrder            uint
	Wedge_stagedOrder     map[*Wedge]uint
	Wedge_orderStaged     map[uint]*Wedge
	Wedges_reference      map[*Wedge]*Wedge
	Wedges_referenceOrder map[*Wedge]uint

	// insertion point for slice of pointers maps
	OnAfterWedgeCreateCallback GongOnAfterCreateInterface[Wedge]
	OnAfterWedgeUpdateCallback GongOnAfterUpdateInterface[Wedge]
	OnAfterWedgeDeleteCallback GongOnAfterDeleteInterface[Wedge]

	Woods                map[*Wood]struct{}
	Woods_instance       map[*Wood]*Wood
	Woods_mapString      map[string]*Wood
	WoodOrder            uint
	Wood_stagedOrder     map[*Wood]uint
	Wood_orderStaged     map[uint]*Wood
	Woods_reference      map[*Wood]*Wood
	Woods_referenceOrder map[*Wood]uint

	// insertion point for slice of pointers maps
	OnAfterWoodCreateCallback GongOnAfterCreateInterface[Wood]
	OnAfterWoodUpdateCallback GongOnAfterUpdateInterface[Wood]
	OnAfterWoodDeleteCallback GongOnAfterDeleteInterface[Wood]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string
	MetaPackageImports     []*MetaPackageImport

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	// end of insertion point

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]GongModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF GongProbeIF
}

type GongStage = Stage

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterBeforeCommitHook(func() { hook(s) })
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterAfterCommitHook(func() { hook(s) })
}

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {
	return stage.StageCore.ApplyBackwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ApplyForwardCommit applies the commit after the current one
func (stage *Stage) ApplyForwardCommit() error {
	return stage.StageCore.ApplyForwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {
	stage.StageCore.ResetHardCore()

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.StageCore.SquashCore()

	// insertion point for clear references
	__gong__clearReferences(&stage.Accords_reference, &stage.Accords_instance, &stage.Accords_referenceOrder)

	__gong__clearReferences(&stage.AccordionRegistrations_reference, &stage.AccordionRegistrations_instance, &stage.AccordionRegistrations_referenceOrder)

	__gong__clearReferences(&stage.Barres_reference, &stage.Barres_instance, &stage.Barres_referenceOrder)

	__gong__clearReferences(&stage.Basss_reference, &stage.Basss_instance, &stage.Basss_referenceOrder)

	__gong__clearReferences(&stage.BassSteps_reference, &stage.BassSteps_instance, &stage.BassSteps_referenceOrder)

	__gong__clearReferences(&stage.BeatUnitTieds_reference, &stage.BeatUnitTieds_instance, &stage.BeatUnitTieds_referenceOrder)

	__gong__clearReferences(&stage.Beaters_reference, &stage.Beaters_instance, &stage.Beaters_referenceOrder)

	__gong__clearReferences(&stage.Brackets_reference, &stage.Brackets_instance, &stage.Brackets_referenceOrder)

	__gong__clearReferences(&stage.Dashess_reference, &stage.Dashess_instance, &stage.Dashess_referenceOrder)

	__gong__clearReferences(&stage.Degrees_reference, &stage.Degrees_instance, &stage.Degrees_referenceOrder)

	__gong__clearReferences(&stage.DegreeAlters_reference, &stage.DegreeAlters_instance, &stage.DegreeAlters_referenceOrder)

	__gong__clearReferences(&stage.DegreeTypes_reference, &stage.DegreeTypes_instance, &stage.DegreeTypes_referenceOrder)

	__gong__clearReferences(&stage.DegreeValues_reference, &stage.DegreeValues_instance, &stage.DegreeValues_referenceOrder)

	__gong__clearReferences(&stage.Directions_reference, &stage.Directions_instance, &stage.Directions_referenceOrder)

	__gong__clearReferences(&stage.DirectionTypes_reference, &stage.DirectionTypes_instance, &stage.DirectionTypes_referenceOrder)

	__gong__clearReferences(&stage.Effects_reference, &stage.Effects_instance, &stage.Effects_referenceOrder)

	__gong__clearReferences(&stage.Features_reference, &stage.Features_instance, &stage.Features_referenceOrder)

	__gong__clearReferences(&stage.FirstFrets_reference, &stage.FirstFrets_instance, &stage.FirstFrets_referenceOrder)

	__gong__clearReferences(&stage.Frames_reference, &stage.Frames_instance, &stage.Frames_referenceOrder)

	__gong__clearReferences(&stage.FrameNotes_reference, &stage.FrameNotes_instance, &stage.FrameNotes_referenceOrder)

	__gong__clearReferences(&stage.Glasss_reference, &stage.Glasss_instance, &stage.Glasss_referenceOrder)

	__gong__clearReferences(&stage.Groupings_reference, &stage.Groupings_instance, &stage.Groupings_referenceOrder)

	__gong__clearReferences(&stage.Harmonys_reference, &stage.Harmonys_instance, &stage.Harmonys_referenceOrder)

	__gong__clearReferences(&stage.HarmonyAlters_reference, &stage.HarmonyAlters_instance, &stage.HarmonyAlters_referenceOrder)

	__gong__clearReferences(&stage.HarpPedalss_reference, &stage.HarpPedalss_instance, &stage.HarpPedalss_referenceOrder)

	__gong__clearReferences(&stage.Images_reference, &stage.Images_instance, &stage.Images_referenceOrder)

	__gong__clearReferences(&stage.InstrumentChanges_reference, &stage.InstrumentChanges_instance, &stage.InstrumentChanges_referenceOrder)

	__gong__clearReferences(&stage.Inversions_reference, &stage.Inversions_instance, &stage.Inversions_referenceOrder)

	__gong__clearReferences(&stage.Kinds_reference, &stage.Kinds_instance, &stage.Kinds_referenceOrder)

	__gong__clearReferences(&stage.Listenings_reference, &stage.Listenings_instance, &stage.Listenings_referenceOrder)

	__gong__clearReferences(&stage.MeasureNumberings_reference, &stage.MeasureNumberings_instance, &stage.MeasureNumberings_referenceOrder)

	__gong__clearReferences(&stage.Membranes_reference, &stage.Membranes_instance, &stage.Membranes_referenceOrder)

	__gong__clearReferences(&stage.Metals_reference, &stage.Metals_instance, &stage.Metals_referenceOrder)

	__gong__clearReferences(&stage.Metronomes_reference, &stage.Metronomes_instance, &stage.Metronomes_referenceOrder)

	__gong__clearReferences(&stage.MetronomeBeams_reference, &stage.MetronomeBeams_instance, &stage.MetronomeBeams_referenceOrder)

	__gong__clearReferences(&stage.MetronomeNotes_reference, &stage.MetronomeNotes_instance, &stage.MetronomeNotes_referenceOrder)

	__gong__clearReferences(&stage.MetronomeTieds_reference, &stage.MetronomeTieds_instance, &stage.MetronomeTieds_referenceOrder)

	__gong__clearReferences(&stage.MetronomeTuplets_reference, &stage.MetronomeTuplets_instance, &stage.MetronomeTuplets_referenceOrder)

	__gong__clearReferences(&stage.Numerals_reference, &stage.Numerals_instance, &stage.Numerals_referenceOrder)

	__gong__clearReferences(&stage.NumeralKeys_reference, &stage.NumeralKeys_instance, &stage.NumeralKeys_referenceOrder)

	__gong__clearReferences(&stage.NumeralRoots_reference, &stage.NumeralRoots_instance, &stage.NumeralRoots_referenceOrder)

	__gong__clearReferences(&stage.OctaveShifts_reference, &stage.OctaveShifts_instance, &stage.OctaveShifts_referenceOrder)

	__gong__clearReferences(&stage.Offsets_reference, &stage.Offsets_instance, &stage.Offsets_referenceOrder)

	__gong__clearReferences(&stage.OtherDirections_reference, &stage.OtherDirections_instance, &stage.OtherDirections_referenceOrder)

	__gong__clearReferences(&stage.OtherListenings_reference, &stage.OtherListenings_instance, &stage.OtherListenings_referenceOrder)

	__gong__clearReferences(&stage.Pedals_reference, &stage.Pedals_instance, &stage.Pedals_referenceOrder)

	__gong__clearReferences(&stage.PedalTunings_reference, &stage.PedalTunings_instance, &stage.PedalTunings_referenceOrder)

	__gong__clearReferences(&stage.PerMinutes_reference, &stage.PerMinutes_instance, &stage.PerMinutes_referenceOrder)

	__gong__clearReferences(&stage.Percussions_reference, &stage.Percussions_instance, &stage.Percussions_referenceOrder)

	__gong__clearReferences(&stage.Pitcheds_reference, &stage.Pitcheds_instance, &stage.Pitcheds_referenceOrder)

	__gong__clearReferences(&stage.PrincipalVoices_reference, &stage.PrincipalVoices_instance, &stage.PrincipalVoices_referenceOrder)

	__gong__clearReferences(&stage.Prints_reference, &stage.Prints_instance, &stage.Prints_referenceOrder)

	__gong__clearReferences(&stage.Roots_reference, &stage.Roots_instance, &stage.Roots_referenceOrder)

	__gong__clearReferences(&stage.RootSteps_reference, &stage.RootSteps_instance, &stage.RootSteps_referenceOrder)

	__gong__clearReferences(&stage.Scordaturas_reference, &stage.Scordaturas_instance, &stage.Scordaturas_referenceOrder)

	__gong__clearReferences(&stage.Sounds_reference, &stage.Sounds_instance, &stage.Sounds_referenceOrder)

	__gong__clearReferences(&stage.StaffDivides_reference, &stage.StaffDivides_instance, &stage.StaffDivides_referenceOrder)

	__gong__clearReferences(&stage.Sticks_reference, &stage.Sticks_instance, &stage.Sticks_referenceOrder)

	__gong__clearReferences(&stage.StringMutes_reference, &stage.StringMutes_instance, &stage.StringMutes_referenceOrder)

	__gong__clearReferences(&stage.StyleTexts_reference, &stage.StyleTexts_instance, &stage.StyleTexts_referenceOrder)

	__gong__clearReferences(&stage.Swings_reference, &stage.Swings_instance, &stage.Swings_referenceOrder)

	__gong__clearReferences(&stage.Syncs_reference, &stage.Syncs_instance, &stage.Syncs_referenceOrder)

	__gong__clearReferences(&stage.TimeModifications_reference, &stage.TimeModifications_instance, &stage.TimeModifications_referenceOrder)

	__gong__clearReferences(&stage.Timpanis_reference, &stage.Timpanis_instance, &stage.Timpanis_referenceOrder)

	__gong__clearReferences(&stage.VirtualInstruments_reference, &stage.VirtualInstruments_instance, &stage.VirtualInstruments_referenceOrder)

	__gong__clearReferences(&stage.Wedges_reference, &stage.Wedges_instance, &stage.Wedges_referenceOrder)

	__gong__clearReferences(&stage.Woods_reference, &stage.Woods_instance, &stage.Woods_referenceOrder)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	stage.StageCore.EndSquash()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	stage.AccordOrder = __gong__recomputeOrder(stage.Accord_stagedOrder)

	stage.AccordionRegistrationOrder = __gong__recomputeOrder(stage.AccordionRegistration_stagedOrder)

	stage.BarreOrder = __gong__recomputeOrder(stage.Barre_stagedOrder)

	stage.BassOrder = __gong__recomputeOrder(stage.Bass_stagedOrder)

	stage.BassStepOrder = __gong__recomputeOrder(stage.BassStep_stagedOrder)

	stage.BeatUnitTiedOrder = __gong__recomputeOrder(stage.BeatUnitTied_stagedOrder)

	stage.BeaterOrder = __gong__recomputeOrder(stage.Beater_stagedOrder)

	stage.BracketOrder = __gong__recomputeOrder(stage.Bracket_stagedOrder)

	stage.DashesOrder = __gong__recomputeOrder(stage.Dashes_stagedOrder)

	stage.DegreeOrder = __gong__recomputeOrder(stage.Degree_stagedOrder)

	stage.DegreeAlterOrder = __gong__recomputeOrder(stage.DegreeAlter_stagedOrder)

	stage.DegreeTypeOrder = __gong__recomputeOrder(stage.DegreeType_stagedOrder)

	stage.DegreeValueOrder = __gong__recomputeOrder(stage.DegreeValue_stagedOrder)

	stage.DirectionOrder = __gong__recomputeOrder(stage.Direction_stagedOrder)

	stage.DirectionTypeOrder = __gong__recomputeOrder(stage.DirectionType_stagedOrder)

	stage.EffectOrder = __gong__recomputeOrder(stage.Effect_stagedOrder)

	stage.FeatureOrder = __gong__recomputeOrder(stage.Feature_stagedOrder)

	stage.FirstFretOrder = __gong__recomputeOrder(stage.FirstFret_stagedOrder)

	stage.FrameOrder = __gong__recomputeOrder(stage.Frame_stagedOrder)

	stage.FrameNoteOrder = __gong__recomputeOrder(stage.FrameNote_stagedOrder)

	stage.GlassOrder = __gong__recomputeOrder(stage.Glass_stagedOrder)

	stage.GroupingOrder = __gong__recomputeOrder(stage.Grouping_stagedOrder)

	stage.HarmonyOrder = __gong__recomputeOrder(stage.Harmony_stagedOrder)

	stage.HarmonyAlterOrder = __gong__recomputeOrder(stage.HarmonyAlter_stagedOrder)

	stage.HarpPedalsOrder = __gong__recomputeOrder(stage.HarpPedals_stagedOrder)

	stage.ImageOrder = __gong__recomputeOrder(stage.Image_stagedOrder)

	stage.InstrumentChangeOrder = __gong__recomputeOrder(stage.InstrumentChange_stagedOrder)

	stage.InversionOrder = __gong__recomputeOrder(stage.Inversion_stagedOrder)

	stage.KindOrder = __gong__recomputeOrder(stage.Kind_stagedOrder)

	stage.ListeningOrder = __gong__recomputeOrder(stage.Listening_stagedOrder)

	stage.MeasureNumberingOrder = __gong__recomputeOrder(stage.MeasureNumbering_stagedOrder)

	stage.MembraneOrder = __gong__recomputeOrder(stage.Membrane_stagedOrder)

	stage.MetalOrder = __gong__recomputeOrder(stage.Metal_stagedOrder)

	stage.MetronomeOrder = __gong__recomputeOrder(stage.Metronome_stagedOrder)

	stage.MetronomeBeamOrder = __gong__recomputeOrder(stage.MetronomeBeam_stagedOrder)

	stage.MetronomeNoteOrder = __gong__recomputeOrder(stage.MetronomeNote_stagedOrder)

	stage.MetronomeTiedOrder = __gong__recomputeOrder(stage.MetronomeTied_stagedOrder)

	stage.MetronomeTupletOrder = __gong__recomputeOrder(stage.MetronomeTuplet_stagedOrder)

	stage.NumeralOrder = __gong__recomputeOrder(stage.Numeral_stagedOrder)

	stage.NumeralKeyOrder = __gong__recomputeOrder(stage.NumeralKey_stagedOrder)

	stage.NumeralRootOrder = __gong__recomputeOrder(stage.NumeralRoot_stagedOrder)

	stage.OctaveShiftOrder = __gong__recomputeOrder(stage.OctaveShift_stagedOrder)

	stage.OffsetOrder = __gong__recomputeOrder(stage.Offset_stagedOrder)

	stage.OtherDirectionOrder = __gong__recomputeOrder(stage.OtherDirection_stagedOrder)

	stage.OtherListeningOrder = __gong__recomputeOrder(stage.OtherListening_stagedOrder)

	stage.PedalOrder = __gong__recomputeOrder(stage.Pedal_stagedOrder)

	stage.PedalTuningOrder = __gong__recomputeOrder(stage.PedalTuning_stagedOrder)

	stage.PerMinuteOrder = __gong__recomputeOrder(stage.PerMinute_stagedOrder)

	stage.PercussionOrder = __gong__recomputeOrder(stage.Percussion_stagedOrder)

	stage.PitchedOrder = __gong__recomputeOrder(stage.Pitched_stagedOrder)

	stage.PrincipalVoiceOrder = __gong__recomputeOrder(stage.PrincipalVoice_stagedOrder)

	stage.PrintOrder = __gong__recomputeOrder(stage.Print_stagedOrder)

	stage.RootOrder = __gong__recomputeOrder(stage.Root_stagedOrder)

	stage.RootStepOrder = __gong__recomputeOrder(stage.RootStep_stagedOrder)

	stage.ScordaturaOrder = __gong__recomputeOrder(stage.Scordatura_stagedOrder)

	stage.SoundOrder = __gong__recomputeOrder(stage.Sound_stagedOrder)

	stage.StaffDivideOrder = __gong__recomputeOrder(stage.StaffDivide_stagedOrder)

	stage.StickOrder = __gong__recomputeOrder(stage.Stick_stagedOrder)

	stage.StringMuteOrder = __gong__recomputeOrder(stage.StringMute_stagedOrder)

	stage.StyleTextOrder = __gong__recomputeOrder(stage.StyleText_stagedOrder)

	stage.SwingOrder = __gong__recomputeOrder(stage.Swing_stagedOrder)

	stage.SyncOrder = __gong__recomputeOrder(stage.Sync_stagedOrder)

	stage.TimeModificationOrder = __gong__recomputeOrder(stage.TimeModification_stagedOrder)

	stage.TimpaniOrder = __gong__recomputeOrder(stage.Timpani_stagedOrder)

	stage.VirtualInstrumentOrder = __gong__recomputeOrder(stage.VirtualInstrument_stagedOrder)

	stage.WedgeOrder = __gong__recomputeOrder(stage.Wedge_stagedOrder)

	stage.WoodOrder = __gong__recomputeOrder(stage.Wood_stagedOrder)

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetProbeIF(probeIF GongProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() GongProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// insertion point for stage ops
func (*Accord) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Accords, stage.Accord_stagedOrder)
}

func (*Accord) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Accord_orderStaged[order]
}

func (*Accord) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Accords_mapString
}

func (*Accord) GongGetInstancesSet(stage *Stage) any {
	return &stage.Accords
}

func (*Accord) GongNewInstance() any {
	return new(Accord)
}

func (*AccordionRegistration) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AccordionRegistrations, stage.AccordionRegistration_stagedOrder)
}

func (*AccordionRegistration) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AccordionRegistration_orderStaged[order]
}

func (*AccordionRegistration) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AccordionRegistrations_mapString
}

func (*AccordionRegistration) GongGetInstancesSet(stage *Stage) any {
	return &stage.AccordionRegistrations
}

func (*AccordionRegistration) GongNewInstance() any {
	return new(AccordionRegistration)
}

func (*Barre) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Barres, stage.Barre_stagedOrder)
}

func (*Barre) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Barre_orderStaged[order]
}

func (*Barre) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Barres_mapString
}

func (*Barre) GongGetInstancesSet(stage *Stage) any {
	return &stage.Barres
}

func (*Barre) GongNewInstance() any {
	return new(Barre)
}

func (*Bass) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Basss, stage.Bass_stagedOrder)
}

func (*Bass) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Bass_orderStaged[order]
}

func (*Bass) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Basss_mapString
}

func (*Bass) GongGetInstancesSet(stage *Stage) any {
	return &stage.Basss
}

func (*Bass) GongNewInstance() any {
	return new(Bass)
}

func (*BassStep) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.BassSteps, stage.BassStep_stagedOrder)
}

func (*BassStep) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.BassStep_orderStaged[order]
}

func (*BassStep) GongGetInstancesMapByName(stage *Stage) any {
	return stage.BassSteps_mapString
}

func (*BassStep) GongGetInstancesSet(stage *Stage) any {
	return &stage.BassSteps
}

func (*BassStep) GongNewInstance() any {
	return new(BassStep)
}

func (*BeatUnitTied) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.BeatUnitTieds, stage.BeatUnitTied_stagedOrder)
}

func (*BeatUnitTied) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.BeatUnitTied_orderStaged[order]
}

func (*BeatUnitTied) GongGetInstancesMapByName(stage *Stage) any {
	return stage.BeatUnitTieds_mapString
}

func (*BeatUnitTied) GongGetInstancesSet(stage *Stage) any {
	return &stage.BeatUnitTieds
}

func (*BeatUnitTied) GongNewInstance() any {
	return new(BeatUnitTied)
}

func (*Beater) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Beaters, stage.Beater_stagedOrder)
}

func (*Beater) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Beater_orderStaged[order]
}

func (*Beater) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Beaters_mapString
}

func (*Beater) GongGetInstancesSet(stage *Stage) any {
	return &stage.Beaters
}

func (*Beater) GongNewInstance() any {
	return new(Beater)
}

func (*Bracket) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Brackets, stage.Bracket_stagedOrder)
}

func (*Bracket) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Bracket_orderStaged[order]
}

func (*Bracket) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Brackets_mapString
}

func (*Bracket) GongGetInstancesSet(stage *Stage) any {
	return &stage.Brackets
}

func (*Bracket) GongNewInstance() any {
	return new(Bracket)
}

func (*Dashes) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Dashess, stage.Dashes_stagedOrder)
}

func (*Dashes) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Dashes_orderStaged[order]
}

func (*Dashes) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Dashess_mapString
}

func (*Dashes) GongGetInstancesSet(stage *Stage) any {
	return &stage.Dashess
}

func (*Dashes) GongNewInstance() any {
	return new(Dashes)
}

func (*Degree) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Degrees, stage.Degree_stagedOrder)
}

func (*Degree) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Degree_orderStaged[order]
}

func (*Degree) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Degrees_mapString
}

func (*Degree) GongGetInstancesSet(stage *Stage) any {
	return &stage.Degrees
}

func (*Degree) GongNewInstance() any {
	return new(Degree)
}

func (*DegreeAlter) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DegreeAlters, stage.DegreeAlter_stagedOrder)
}

func (*DegreeAlter) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DegreeAlter_orderStaged[order]
}

func (*DegreeAlter) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DegreeAlters_mapString
}

func (*DegreeAlter) GongGetInstancesSet(stage *Stage) any {
	return &stage.DegreeAlters
}

func (*DegreeAlter) GongNewInstance() any {
	return new(DegreeAlter)
}

func (*DegreeType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DegreeTypes, stage.DegreeType_stagedOrder)
}

func (*DegreeType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DegreeType_orderStaged[order]
}

func (*DegreeType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DegreeTypes_mapString
}

func (*DegreeType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DegreeTypes
}

func (*DegreeType) GongNewInstance() any {
	return new(DegreeType)
}

func (*DegreeValue) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DegreeValues, stage.DegreeValue_stagedOrder)
}

func (*DegreeValue) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DegreeValue_orderStaged[order]
}

func (*DegreeValue) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DegreeValues_mapString
}

func (*DegreeValue) GongGetInstancesSet(stage *Stage) any {
	return &stage.DegreeValues
}

func (*DegreeValue) GongNewInstance() any {
	return new(DegreeValue)
}

func (*Direction) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Directions, stage.Direction_stagedOrder)
}

func (*Direction) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Direction_orderStaged[order]
}

func (*Direction) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Directions_mapString
}

func (*Direction) GongGetInstancesSet(stage *Stage) any {
	return &stage.Directions
}

func (*Direction) GongNewInstance() any {
	return new(Direction)
}

func (*DirectionType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DirectionTypes, stage.DirectionType_stagedOrder)
}

func (*DirectionType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DirectionType_orderStaged[order]
}

func (*DirectionType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DirectionTypes_mapString
}

func (*DirectionType) GongGetInstancesSet(stage *Stage) any {
	return &stage.DirectionTypes
}

func (*DirectionType) GongNewInstance() any {
	return new(DirectionType)
}

func (*Effect) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Effects, stage.Effect_stagedOrder)
}

func (*Effect) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Effect_orderStaged[order]
}

func (*Effect) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Effects_mapString
}

func (*Effect) GongGetInstancesSet(stage *Stage) any {
	return &stage.Effects
}

func (*Effect) GongNewInstance() any {
	return new(Effect)
}

func (*Feature) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Features, stage.Feature_stagedOrder)
}

func (*Feature) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Feature_orderStaged[order]
}

func (*Feature) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Features_mapString
}

func (*Feature) GongGetInstancesSet(stage *Stage) any {
	return &stage.Features
}

func (*Feature) GongNewInstance() any {
	return new(Feature)
}

func (*FirstFret) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.FirstFrets, stage.FirstFret_stagedOrder)
}

func (*FirstFret) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.FirstFret_orderStaged[order]
}

func (*FirstFret) GongGetInstancesMapByName(stage *Stage) any {
	return stage.FirstFrets_mapString
}

func (*FirstFret) GongGetInstancesSet(stage *Stage) any {
	return &stage.FirstFrets
}

func (*FirstFret) GongNewInstance() any {
	return new(FirstFret)
}

func (*Frame) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Frames, stage.Frame_stagedOrder)
}

func (*Frame) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Frame_orderStaged[order]
}

func (*Frame) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Frames_mapString
}

func (*Frame) GongGetInstancesSet(stage *Stage) any {
	return &stage.Frames
}

func (*Frame) GongNewInstance() any {
	return new(Frame)
}

func (*FrameNote) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.FrameNotes, stage.FrameNote_stagedOrder)
}

func (*FrameNote) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.FrameNote_orderStaged[order]
}

func (*FrameNote) GongGetInstancesMapByName(stage *Stage) any {
	return stage.FrameNotes_mapString
}

func (*FrameNote) GongGetInstancesSet(stage *Stage) any {
	return &stage.FrameNotes
}

func (*FrameNote) GongNewInstance() any {
	return new(FrameNote)
}

func (*Glass) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Glasss, stage.Glass_stagedOrder)
}

func (*Glass) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Glass_orderStaged[order]
}

func (*Glass) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Glasss_mapString
}

func (*Glass) GongGetInstancesSet(stage *Stage) any {
	return &stage.Glasss
}

func (*Glass) GongNewInstance() any {
	return new(Glass)
}

func (*Grouping) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Groupings, stage.Grouping_stagedOrder)
}

func (*Grouping) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Grouping_orderStaged[order]
}

func (*Grouping) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Groupings_mapString
}

func (*Grouping) GongGetInstancesSet(stage *Stage) any {
	return &stage.Groupings
}

func (*Grouping) GongNewInstance() any {
	return new(Grouping)
}

func (*Harmony) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Harmonys, stage.Harmony_stagedOrder)
}

func (*Harmony) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Harmony_orderStaged[order]
}

func (*Harmony) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Harmonys_mapString
}

func (*Harmony) GongGetInstancesSet(stage *Stage) any {
	return &stage.Harmonys
}

func (*Harmony) GongNewInstance() any {
	return new(Harmony)
}

func (*HarmonyAlter) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.HarmonyAlters, stage.HarmonyAlter_stagedOrder)
}

func (*HarmonyAlter) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.HarmonyAlter_orderStaged[order]
}

func (*HarmonyAlter) GongGetInstancesMapByName(stage *Stage) any {
	return stage.HarmonyAlters_mapString
}

func (*HarmonyAlter) GongGetInstancesSet(stage *Stage) any {
	return &stage.HarmonyAlters
}

func (*HarmonyAlter) GongNewInstance() any {
	return new(HarmonyAlter)
}

func (*HarpPedals) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.HarpPedalss, stage.HarpPedals_stagedOrder)
}

func (*HarpPedals) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.HarpPedals_orderStaged[order]
}

func (*HarpPedals) GongGetInstancesMapByName(stage *Stage) any {
	return stage.HarpPedalss_mapString
}

func (*HarpPedals) GongGetInstancesSet(stage *Stage) any {
	return &stage.HarpPedalss
}

func (*HarpPedals) GongNewInstance() any {
	return new(HarpPedals)
}

func (*Image) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Images, stage.Image_stagedOrder)
}

func (*Image) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Image_orderStaged[order]
}

func (*Image) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Images_mapString
}

func (*Image) GongGetInstancesSet(stage *Stage) any {
	return &stage.Images
}

func (*Image) GongNewInstance() any {
	return new(Image)
}

func (*InstrumentChange) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.InstrumentChanges, stage.InstrumentChange_stagedOrder)
}

func (*InstrumentChange) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.InstrumentChange_orderStaged[order]
}

func (*InstrumentChange) GongGetInstancesMapByName(stage *Stage) any {
	return stage.InstrumentChanges_mapString
}

func (*InstrumentChange) GongGetInstancesSet(stage *Stage) any {
	return &stage.InstrumentChanges
}

func (*InstrumentChange) GongNewInstance() any {
	return new(InstrumentChange)
}

func (*Inversion) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Inversions, stage.Inversion_stagedOrder)
}

func (*Inversion) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Inversion_orderStaged[order]
}

func (*Inversion) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Inversions_mapString
}

func (*Inversion) GongGetInstancesSet(stage *Stage) any {
	return &stage.Inversions
}

func (*Inversion) GongNewInstance() any {
	return new(Inversion)
}

func (*Kind) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Kinds, stage.Kind_stagedOrder)
}

func (*Kind) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Kind_orderStaged[order]
}

func (*Kind) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Kinds_mapString
}

func (*Kind) GongGetInstancesSet(stage *Stage) any {
	return &stage.Kinds
}

func (*Kind) GongNewInstance() any {
	return new(Kind)
}

func (*Listening) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Listenings, stage.Listening_stagedOrder)
}

func (*Listening) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Listening_orderStaged[order]
}

func (*Listening) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Listenings_mapString
}

func (*Listening) GongGetInstancesSet(stage *Stage) any {
	return &stage.Listenings
}

func (*Listening) GongNewInstance() any {
	return new(Listening)
}

func (*MeasureNumbering) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MeasureNumberings, stage.MeasureNumbering_stagedOrder)
}

func (*MeasureNumbering) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MeasureNumbering_orderStaged[order]
}

func (*MeasureNumbering) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MeasureNumberings_mapString
}

func (*MeasureNumbering) GongGetInstancesSet(stage *Stage) any {
	return &stage.MeasureNumberings
}

func (*MeasureNumbering) GongNewInstance() any {
	return new(MeasureNumbering)
}

func (*Membrane) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Membranes, stage.Membrane_stagedOrder)
}

func (*Membrane) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Membrane_orderStaged[order]
}

func (*Membrane) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Membranes_mapString
}

func (*Membrane) GongGetInstancesSet(stage *Stage) any {
	return &stage.Membranes
}

func (*Membrane) GongNewInstance() any {
	return new(Membrane)
}

func (*Metal) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Metals, stage.Metal_stagedOrder)
}

func (*Metal) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Metal_orderStaged[order]
}

func (*Metal) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Metals_mapString
}

func (*Metal) GongGetInstancesSet(stage *Stage) any {
	return &stage.Metals
}

func (*Metal) GongNewInstance() any {
	return new(Metal)
}

func (*Metronome) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Metronomes, stage.Metronome_stagedOrder)
}

func (*Metronome) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Metronome_orderStaged[order]
}

func (*Metronome) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Metronomes_mapString
}

func (*Metronome) GongGetInstancesSet(stage *Stage) any {
	return &stage.Metronomes
}

func (*Metronome) GongNewInstance() any {
	return new(Metronome)
}

func (*MetronomeBeam) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MetronomeBeams, stage.MetronomeBeam_stagedOrder)
}

func (*MetronomeBeam) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MetronomeBeam_orderStaged[order]
}

func (*MetronomeBeam) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MetronomeBeams_mapString
}

func (*MetronomeBeam) GongGetInstancesSet(stage *Stage) any {
	return &stage.MetronomeBeams
}

func (*MetronomeBeam) GongNewInstance() any {
	return new(MetronomeBeam)
}

func (*MetronomeNote) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MetronomeNotes, stage.MetronomeNote_stagedOrder)
}

func (*MetronomeNote) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MetronomeNote_orderStaged[order]
}

func (*MetronomeNote) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MetronomeNotes_mapString
}

func (*MetronomeNote) GongGetInstancesSet(stage *Stage) any {
	return &stage.MetronomeNotes
}

func (*MetronomeNote) GongNewInstance() any {
	return new(MetronomeNote)
}

func (*MetronomeTied) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MetronomeTieds, stage.MetronomeTied_stagedOrder)
}

func (*MetronomeTied) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MetronomeTied_orderStaged[order]
}

func (*MetronomeTied) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MetronomeTieds_mapString
}

func (*MetronomeTied) GongGetInstancesSet(stage *Stage) any {
	return &stage.MetronomeTieds
}

func (*MetronomeTied) GongNewInstance() any {
	return new(MetronomeTied)
}

func (*MetronomeTuplet) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MetronomeTuplets, stage.MetronomeTuplet_stagedOrder)
}

func (*MetronomeTuplet) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MetronomeTuplet_orderStaged[order]
}

func (*MetronomeTuplet) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MetronomeTuplets_mapString
}

func (*MetronomeTuplet) GongGetInstancesSet(stage *Stage) any {
	return &stage.MetronomeTuplets
}

func (*MetronomeTuplet) GongNewInstance() any {
	return new(MetronomeTuplet)
}

func (*Numeral) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Numerals, stage.Numeral_stagedOrder)
}

func (*Numeral) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Numeral_orderStaged[order]
}

func (*Numeral) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Numerals_mapString
}

func (*Numeral) GongGetInstancesSet(stage *Stage) any {
	return &stage.Numerals
}

func (*Numeral) GongNewInstance() any {
	return new(Numeral)
}

func (*NumeralKey) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NumeralKeys, stage.NumeralKey_stagedOrder)
}

func (*NumeralKey) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NumeralKey_orderStaged[order]
}

func (*NumeralKey) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NumeralKeys_mapString
}

func (*NumeralKey) GongGetInstancesSet(stage *Stage) any {
	return &stage.NumeralKeys
}

func (*NumeralKey) GongNewInstance() any {
	return new(NumeralKey)
}

func (*NumeralRoot) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NumeralRoots, stage.NumeralRoot_stagedOrder)
}

func (*NumeralRoot) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NumeralRoot_orderStaged[order]
}

func (*NumeralRoot) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NumeralRoots_mapString
}

func (*NumeralRoot) GongGetInstancesSet(stage *Stage) any {
	return &stage.NumeralRoots
}

func (*NumeralRoot) GongNewInstance() any {
	return new(NumeralRoot)
}

func (*OctaveShift) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OctaveShifts, stage.OctaveShift_stagedOrder)
}

func (*OctaveShift) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OctaveShift_orderStaged[order]
}

func (*OctaveShift) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OctaveShifts_mapString
}

func (*OctaveShift) GongGetInstancesSet(stage *Stage) any {
	return &stage.OctaveShifts
}

func (*OctaveShift) GongNewInstance() any {
	return new(OctaveShift)
}

func (*Offset) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Offsets, stage.Offset_stagedOrder)
}

func (*Offset) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Offset_orderStaged[order]
}

func (*Offset) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Offsets_mapString
}

func (*Offset) GongGetInstancesSet(stage *Stage) any {
	return &stage.Offsets
}

func (*Offset) GongNewInstance() any {
	return new(Offset)
}

func (*OtherDirection) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OtherDirections, stage.OtherDirection_stagedOrder)
}

func (*OtherDirection) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OtherDirection_orderStaged[order]
}

func (*OtherDirection) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OtherDirections_mapString
}

func (*OtherDirection) GongGetInstancesSet(stage *Stage) any {
	return &stage.OtherDirections
}

func (*OtherDirection) GongNewInstance() any {
	return new(OtherDirection)
}

func (*OtherListening) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OtherListenings, stage.OtherListening_stagedOrder)
}

func (*OtherListening) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OtherListening_orderStaged[order]
}

func (*OtherListening) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OtherListenings_mapString
}

func (*OtherListening) GongGetInstancesSet(stage *Stage) any {
	return &stage.OtherListenings
}

func (*OtherListening) GongNewInstance() any {
	return new(OtherListening)
}

func (*Pedal) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Pedals, stage.Pedal_stagedOrder)
}

func (*Pedal) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Pedal_orderStaged[order]
}

func (*Pedal) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Pedals_mapString
}

func (*Pedal) GongGetInstancesSet(stage *Stage) any {
	return &stage.Pedals
}

func (*Pedal) GongNewInstance() any {
	return new(Pedal)
}

func (*PedalTuning) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PedalTunings, stage.PedalTuning_stagedOrder)
}

func (*PedalTuning) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PedalTuning_orderStaged[order]
}

func (*PedalTuning) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PedalTunings_mapString
}

func (*PedalTuning) GongGetInstancesSet(stage *Stage) any {
	return &stage.PedalTunings
}

func (*PedalTuning) GongNewInstance() any {
	return new(PedalTuning)
}

func (*PerMinute) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PerMinutes, stage.PerMinute_stagedOrder)
}

func (*PerMinute) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PerMinute_orderStaged[order]
}

func (*PerMinute) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PerMinutes_mapString
}

func (*PerMinute) GongGetInstancesSet(stage *Stage) any {
	return &stage.PerMinutes
}

func (*PerMinute) GongNewInstance() any {
	return new(PerMinute)
}

func (*Percussion) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Percussions, stage.Percussion_stagedOrder)
}

func (*Percussion) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Percussion_orderStaged[order]
}

func (*Percussion) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Percussions_mapString
}

func (*Percussion) GongGetInstancesSet(stage *Stage) any {
	return &stage.Percussions
}

func (*Percussion) GongNewInstance() any {
	return new(Percussion)
}

func (*Pitched) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Pitcheds, stage.Pitched_stagedOrder)
}

func (*Pitched) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Pitched_orderStaged[order]
}

func (*Pitched) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Pitcheds_mapString
}

func (*Pitched) GongGetInstancesSet(stage *Stage) any {
	return &stage.Pitcheds
}

func (*Pitched) GongNewInstance() any {
	return new(Pitched)
}

func (*PrincipalVoice) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PrincipalVoices, stage.PrincipalVoice_stagedOrder)
}

func (*PrincipalVoice) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PrincipalVoice_orderStaged[order]
}

func (*PrincipalVoice) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PrincipalVoices_mapString
}

func (*PrincipalVoice) GongGetInstancesSet(stage *Stage) any {
	return &stage.PrincipalVoices
}

func (*PrincipalVoice) GongNewInstance() any {
	return new(PrincipalVoice)
}

func (*Print) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Prints, stage.Print_stagedOrder)
}

func (*Print) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Print_orderStaged[order]
}

func (*Print) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Prints_mapString
}

func (*Print) GongGetInstancesSet(stage *Stage) any {
	return &stage.Prints
}

func (*Print) GongNewInstance() any {
	return new(Print)
}

func (*Root) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Roots, stage.Root_stagedOrder)
}

func (*Root) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Root_orderStaged[order]
}

func (*Root) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Roots_mapString
}

func (*Root) GongGetInstancesSet(stage *Stage) any {
	return &stage.Roots
}

func (*Root) GongNewInstance() any {
	return new(Root)
}

func (*RootStep) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.RootSteps, stage.RootStep_stagedOrder)
}

func (*RootStep) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.RootStep_orderStaged[order]
}

func (*RootStep) GongGetInstancesMapByName(stage *Stage) any {
	return stage.RootSteps_mapString
}

func (*RootStep) GongGetInstancesSet(stage *Stage) any {
	return &stage.RootSteps
}

func (*RootStep) GongNewInstance() any {
	return new(RootStep)
}

func (*Scordatura) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Scordaturas, stage.Scordatura_stagedOrder)
}

func (*Scordatura) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Scordatura_orderStaged[order]
}

func (*Scordatura) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Scordaturas_mapString
}

func (*Scordatura) GongGetInstancesSet(stage *Stage) any {
	return &stage.Scordaturas
}

func (*Scordatura) GongNewInstance() any {
	return new(Scordatura)
}

func (*Sound) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Sounds, stage.Sound_stagedOrder)
}

func (*Sound) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Sound_orderStaged[order]
}

func (*Sound) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Sounds_mapString
}

func (*Sound) GongGetInstancesSet(stage *Stage) any {
	return &stage.Sounds
}

func (*Sound) GongNewInstance() any {
	return new(Sound)
}

func (*StaffDivide) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StaffDivides, stage.StaffDivide_stagedOrder)
}

func (*StaffDivide) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StaffDivide_orderStaged[order]
}

func (*StaffDivide) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StaffDivides_mapString
}

func (*StaffDivide) GongGetInstancesSet(stage *Stage) any {
	return &stage.StaffDivides
}

func (*StaffDivide) GongNewInstance() any {
	return new(StaffDivide)
}

func (*Stick) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Sticks, stage.Stick_stagedOrder)
}

func (*Stick) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Stick_orderStaged[order]
}

func (*Stick) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Sticks_mapString
}

func (*Stick) GongGetInstancesSet(stage *Stage) any {
	return &stage.Sticks
}

func (*Stick) GongNewInstance() any {
	return new(Stick)
}

func (*StringMute) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StringMutes, stage.StringMute_stagedOrder)
}

func (*StringMute) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StringMute_orderStaged[order]
}

func (*StringMute) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StringMutes_mapString
}

func (*StringMute) GongGetInstancesSet(stage *Stage) any {
	return &stage.StringMutes
}

func (*StringMute) GongNewInstance() any {
	return new(StringMute)
}

func (*StyleText) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StyleTexts, stage.StyleText_stagedOrder)
}

func (*StyleText) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StyleText_orderStaged[order]
}

func (*StyleText) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StyleTexts_mapString
}

func (*StyleText) GongGetInstancesSet(stage *Stage) any {
	return &stage.StyleTexts
}

func (*StyleText) GongNewInstance() any {
	return new(StyleText)
}

func (*Swing) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Swings, stage.Swing_stagedOrder)
}

func (*Swing) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Swing_orderStaged[order]
}

func (*Swing) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Swings_mapString
}

func (*Swing) GongGetInstancesSet(stage *Stage) any {
	return &stage.Swings
}

func (*Swing) GongNewInstance() any {
	return new(Swing)
}

func (*Sync) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Syncs, stage.Sync_stagedOrder)
}

func (*Sync) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Sync_orderStaged[order]
}

func (*Sync) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Syncs_mapString
}

func (*Sync) GongGetInstancesSet(stage *Stage) any {
	return &stage.Syncs
}

func (*Sync) GongNewInstance() any {
	return new(Sync)
}

func (*TimeModification) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TimeModifications, stage.TimeModification_stagedOrder)
}

func (*TimeModification) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TimeModification_orderStaged[order]
}

func (*TimeModification) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TimeModifications_mapString
}

func (*TimeModification) GongGetInstancesSet(stage *Stage) any {
	return &stage.TimeModifications
}

func (*TimeModification) GongNewInstance() any {
	return new(TimeModification)
}

func (*Timpani) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Timpanis, stage.Timpani_stagedOrder)
}

func (*Timpani) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Timpani_orderStaged[order]
}

func (*Timpani) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Timpanis_mapString
}

func (*Timpani) GongGetInstancesSet(stage *Stage) any {
	return &stage.Timpanis
}

func (*Timpani) GongNewInstance() any {
	return new(Timpani)
}

func (*VirtualInstrument) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.VirtualInstruments, stage.VirtualInstrument_stagedOrder)
}

func (*VirtualInstrument) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.VirtualInstrument_orderStaged[order]
}

func (*VirtualInstrument) GongGetInstancesMapByName(stage *Stage) any {
	return stage.VirtualInstruments_mapString
}

func (*VirtualInstrument) GongGetInstancesSet(stage *Stage) any {
	return &stage.VirtualInstruments
}

func (*VirtualInstrument) GongNewInstance() any {
	return new(VirtualInstrument)
}

func (*Wedge) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Wedges, stage.Wedge_stagedOrder)
}

func (*Wedge) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Wedge_orderStaged[order]
}

func (*Wedge) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Wedges_mapString
}

func (*Wedge) GongGetInstancesSet(stage *Stage) any {
	return &stage.Wedges
}

func (*Wedge) GongNewInstance() any {
	return new(Wedge)
}

func (*Wood) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Woods, stage.Wood_stagedOrder)
}

func (*Wood) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Wood_orderStaged[order]
}

func (*Wood) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Woods_mapString
}

func (*Wood) GongGetInstancesSet(stage *Stage) any {
	return &stage.Woods
}

func (*Wood) GongNewInstance() any {
	return new(Wood)
}


// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	if stage == nil {
		return nil
	}
	var t T
	return t.GongGetInstancesByOrder(stage).([]T)
}

func __gong__getStructInstancesByOrder[T GongstructPtr](set map[T]struct{}, order map[T]uint) (res []T) {
	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("getStructInstancesByOrder: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func __gong__castSlice[T any, S any](s []S) []T {
	res := make([]T, len(s))
	for i, v := range s {
		res[i] = any(v).(T)
	}
	return res
}

func __gong__stage[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	order *uint,
	mapString map[string]T,
	instance T,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		stagedOrder[instance] = *order
		orderStaged[*order] = instance
		*order++
	}
	mapString[name] = instance
}

func __gong__stagePreserveOrder[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	currentOrder *uint,
	mapString map[string]T,
	instance T,
	order uint,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		if order > *currentOrder {
			*currentOrder = order
		}
		stagedOrder[instance] = order
		orderStaged[order] = instance
		*currentOrder++
	}
	mapString[name] = instance
}

func __gong__unstage[T comparable](
	instances map[T]struct{},
	mapString map[string]T,
	instance T,
	name string,
) {
	delete(instances, instance)
	delete(mapString, name)
}

func __gong__recomputeOrder[T comparable](stagedOrder map[T]uint) uint {
	var maxOrder uint
	var found bool
	for _, order := range stagedOrder {
		if !found || order > maxOrder {
			maxOrder = order
			found = true
		}
	}
	if found {
		return maxOrder + 1
	}
	return 0
}

func __gong__rebuildMapString[T interface {
	comparable
	GetName() string
}](staged map[T]struct{}, mapString *map[string]T) {
	*mapString = make(map[string]T, len(staged))
	for instance := range staged {
		(*mapString)[instance.GetName()] = instance
	}
}

func __gong__clearReferences[T comparable](ref *map[T]T, inst *map[T]T, refOrder *map[T]uint) {
	*ref = make(map[T]T)
	*inst = make(map[T]T)
	*refOrder = make(map[T]uint)
}

func __gong__resetStageType[T comparable](staged *map[T]struct{}, mapString *map[string]T, stagedOrder *map[T]uint, order *uint) {
	*staged = make(map[T]struct{})
	*mapString = make(map[string]T)
	*stagedOrder = make(map[T]uint)
	*order = 0
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type GongOnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

type OnInitCommitInterface = GongOnInitCommitInterface

// GongOnAfterCreateInterface callback when an instance is updated from the front
type GongOnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

type OnAfterCreateInterface[Type Gongstruct] = GongOnAfterCreateInterface[Type]

// GongOnAfterUpdateInterface callback when an instance is updated from the front
type GongOnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

type OnAfterUpdateInterface[Type Gongstruct] = GongOnAfterUpdateInterface[Type]

// GongOnAfterDeleteInterface callback when an instance is updated from the front
type GongOnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type OnAfterDeleteInterface[Type Gongstruct] = GongOnAfterDeleteInterface[Type]

type GongBackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

type BackRepoInterface = GongBackRepoInterface

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		Accords:           make(map[*Accord]struct{}),
		Accords_mapString: make(map[string]*Accord),

		AccordionRegistrations:           make(map[*AccordionRegistration]struct{}),
		AccordionRegistrations_mapString: make(map[string]*AccordionRegistration),

		Barres:           make(map[*Barre]struct{}),
		Barres_mapString: make(map[string]*Barre),

		Basss:           make(map[*Bass]struct{}),
		Basss_mapString: make(map[string]*Bass),

		BassSteps:           make(map[*BassStep]struct{}),
		BassSteps_mapString: make(map[string]*BassStep),

		BeatUnitTieds:           make(map[*BeatUnitTied]struct{}),
		BeatUnitTieds_mapString: make(map[string]*BeatUnitTied),

		Beaters:           make(map[*Beater]struct{}),
		Beaters_mapString: make(map[string]*Beater),

		Brackets:           make(map[*Bracket]struct{}),
		Brackets_mapString: make(map[string]*Bracket),

		Dashess:           make(map[*Dashes]struct{}),
		Dashess_mapString: make(map[string]*Dashes),

		Degrees:           make(map[*Degree]struct{}),
		Degrees_mapString: make(map[string]*Degree),

		DegreeAlters:           make(map[*DegreeAlter]struct{}),
		DegreeAlters_mapString: make(map[string]*DegreeAlter),

		DegreeTypes:           make(map[*DegreeType]struct{}),
		DegreeTypes_mapString: make(map[string]*DegreeType),

		DegreeValues:           make(map[*DegreeValue]struct{}),
		DegreeValues_mapString: make(map[string]*DegreeValue),

		Directions:           make(map[*Direction]struct{}),
		Directions_mapString: make(map[string]*Direction),

		DirectionTypes:           make(map[*DirectionType]struct{}),
		DirectionTypes_mapString: make(map[string]*DirectionType),

		Effects:           make(map[*Effect]struct{}),
		Effects_mapString: make(map[string]*Effect),

		Features:           make(map[*Feature]struct{}),
		Features_mapString: make(map[string]*Feature),

		FirstFrets:           make(map[*FirstFret]struct{}),
		FirstFrets_mapString: make(map[string]*FirstFret),

		Frames:           make(map[*Frame]struct{}),
		Frames_mapString: make(map[string]*Frame),

		FrameNotes:           make(map[*FrameNote]struct{}),
		FrameNotes_mapString: make(map[string]*FrameNote),

		Glasss:           make(map[*Glass]struct{}),
		Glasss_mapString: make(map[string]*Glass),

		Groupings:           make(map[*Grouping]struct{}),
		Groupings_mapString: make(map[string]*Grouping),

		Harmonys:           make(map[*Harmony]struct{}),
		Harmonys_mapString: make(map[string]*Harmony),

		HarmonyAlters:           make(map[*HarmonyAlter]struct{}),
		HarmonyAlters_mapString: make(map[string]*HarmonyAlter),

		HarpPedalss:           make(map[*HarpPedals]struct{}),
		HarpPedalss_mapString: make(map[string]*HarpPedals),

		Images:           make(map[*Image]struct{}),
		Images_mapString: make(map[string]*Image),

		InstrumentChanges:           make(map[*InstrumentChange]struct{}),
		InstrumentChanges_mapString: make(map[string]*InstrumentChange),

		Inversions:           make(map[*Inversion]struct{}),
		Inversions_mapString: make(map[string]*Inversion),

		Kinds:           make(map[*Kind]struct{}),
		Kinds_mapString: make(map[string]*Kind),

		Listenings:           make(map[*Listening]struct{}),
		Listenings_mapString: make(map[string]*Listening),

		MeasureNumberings:           make(map[*MeasureNumbering]struct{}),
		MeasureNumberings_mapString: make(map[string]*MeasureNumbering),

		Membranes:           make(map[*Membrane]struct{}),
		Membranes_mapString: make(map[string]*Membrane),

		Metals:           make(map[*Metal]struct{}),
		Metals_mapString: make(map[string]*Metal),

		Metronomes:           make(map[*Metronome]struct{}),
		Metronomes_mapString: make(map[string]*Metronome),

		MetronomeBeams:           make(map[*MetronomeBeam]struct{}),
		MetronomeBeams_mapString: make(map[string]*MetronomeBeam),

		MetronomeNotes:           make(map[*MetronomeNote]struct{}),
		MetronomeNotes_mapString: make(map[string]*MetronomeNote),

		MetronomeTieds:           make(map[*MetronomeTied]struct{}),
		MetronomeTieds_mapString: make(map[string]*MetronomeTied),

		MetronomeTuplets:           make(map[*MetronomeTuplet]struct{}),
		MetronomeTuplets_mapString: make(map[string]*MetronomeTuplet),

		Numerals:           make(map[*Numeral]struct{}),
		Numerals_mapString: make(map[string]*Numeral),

		NumeralKeys:           make(map[*NumeralKey]struct{}),
		NumeralKeys_mapString: make(map[string]*NumeralKey),

		NumeralRoots:           make(map[*NumeralRoot]struct{}),
		NumeralRoots_mapString: make(map[string]*NumeralRoot),

		OctaveShifts:           make(map[*OctaveShift]struct{}),
		OctaveShifts_mapString: make(map[string]*OctaveShift),

		Offsets:           make(map[*Offset]struct{}),
		Offsets_mapString: make(map[string]*Offset),

		OtherDirections:           make(map[*OtherDirection]struct{}),
		OtherDirections_mapString: make(map[string]*OtherDirection),

		OtherListenings:           make(map[*OtherListening]struct{}),
		OtherListenings_mapString: make(map[string]*OtherListening),

		Pedals:           make(map[*Pedal]struct{}),
		Pedals_mapString: make(map[string]*Pedal),

		PedalTunings:           make(map[*PedalTuning]struct{}),
		PedalTunings_mapString: make(map[string]*PedalTuning),

		PerMinutes:           make(map[*PerMinute]struct{}),
		PerMinutes_mapString: make(map[string]*PerMinute),

		Percussions:           make(map[*Percussion]struct{}),
		Percussions_mapString: make(map[string]*Percussion),

		Pitcheds:           make(map[*Pitched]struct{}),
		Pitcheds_mapString: make(map[string]*Pitched),

		PrincipalVoices:           make(map[*PrincipalVoice]struct{}),
		PrincipalVoices_mapString: make(map[string]*PrincipalVoice),

		Prints:           make(map[*Print]struct{}),
		Prints_mapString: make(map[string]*Print),

		Roots:           make(map[*Root]struct{}),
		Roots_mapString: make(map[string]*Root),

		RootSteps:           make(map[*RootStep]struct{}),
		RootSteps_mapString: make(map[string]*RootStep),

		Scordaturas:           make(map[*Scordatura]struct{}),
		Scordaturas_mapString: make(map[string]*Scordatura),

		Sounds:           make(map[*Sound]struct{}),
		Sounds_mapString: make(map[string]*Sound),

		StaffDivides:           make(map[*StaffDivide]struct{}),
		StaffDivides_mapString: make(map[string]*StaffDivide),

		Sticks:           make(map[*Stick]struct{}),
		Sticks_mapString: make(map[string]*Stick),

		StringMutes:           make(map[*StringMute]struct{}),
		StringMutes_mapString: make(map[string]*StringMute),

		StyleTexts:           make(map[*StyleText]struct{}),
		StyleTexts_mapString: make(map[string]*StyleText),

		Swings:           make(map[*Swing]struct{}),
		Swings_mapString: make(map[string]*Swing),

		Syncs:           make(map[*Sync]struct{}),
		Syncs_mapString: make(map[string]*Sync),

		TimeModifications:           make(map[*TimeModification]struct{}),
		TimeModifications_mapString: make(map[string]*TimeModification),

		Timpanis:           make(map[*Timpani]struct{}),
		Timpanis_mapString: make(map[string]*Timpani),

		VirtualInstruments:           make(map[*VirtualInstrument]struct{}),
		VirtualInstruments_mapString: make(map[string]*VirtualInstrument),

		Wedges:           make(map[*Wedge]struct{}),
		Wedges_mapString: make(map[string]*Wedge),

		Woods:           make(map[*Wood]struct{}),
		Woods_mapString: make(map[string]*Wood),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Accord_stagedOrder: make(map[*Accord]uint),
		Accord_orderStaged: make(map[uint]*Accord),
		Accords_reference:  make(map[*Accord]*Accord),

		AccordionRegistration_stagedOrder: make(map[*AccordionRegistration]uint),
		AccordionRegistration_orderStaged: make(map[uint]*AccordionRegistration),
		AccordionRegistrations_reference:  make(map[*AccordionRegistration]*AccordionRegistration),

		Barre_stagedOrder: make(map[*Barre]uint),
		Barre_orderStaged: make(map[uint]*Barre),
		Barres_reference:  make(map[*Barre]*Barre),

		Bass_stagedOrder: make(map[*Bass]uint),
		Bass_orderStaged: make(map[uint]*Bass),
		Basss_reference:  make(map[*Bass]*Bass),

		BassStep_stagedOrder: make(map[*BassStep]uint),
		BassStep_orderStaged: make(map[uint]*BassStep),
		BassSteps_reference:  make(map[*BassStep]*BassStep),

		BeatUnitTied_stagedOrder: make(map[*BeatUnitTied]uint),
		BeatUnitTied_orderStaged: make(map[uint]*BeatUnitTied),
		BeatUnitTieds_reference:  make(map[*BeatUnitTied]*BeatUnitTied),

		Beater_stagedOrder: make(map[*Beater]uint),
		Beater_orderStaged: make(map[uint]*Beater),
		Beaters_reference:  make(map[*Beater]*Beater),

		Bracket_stagedOrder: make(map[*Bracket]uint),
		Bracket_orderStaged: make(map[uint]*Bracket),
		Brackets_reference:  make(map[*Bracket]*Bracket),

		Dashes_stagedOrder: make(map[*Dashes]uint),
		Dashes_orderStaged: make(map[uint]*Dashes),
		Dashess_reference:  make(map[*Dashes]*Dashes),

		Degree_stagedOrder: make(map[*Degree]uint),
		Degree_orderStaged: make(map[uint]*Degree),
		Degrees_reference:  make(map[*Degree]*Degree),

		DegreeAlter_stagedOrder: make(map[*DegreeAlter]uint),
		DegreeAlter_orderStaged: make(map[uint]*DegreeAlter),
		DegreeAlters_reference:  make(map[*DegreeAlter]*DegreeAlter),

		DegreeType_stagedOrder: make(map[*DegreeType]uint),
		DegreeType_orderStaged: make(map[uint]*DegreeType),
		DegreeTypes_reference:  make(map[*DegreeType]*DegreeType),

		DegreeValue_stagedOrder: make(map[*DegreeValue]uint),
		DegreeValue_orderStaged: make(map[uint]*DegreeValue),
		DegreeValues_reference:  make(map[*DegreeValue]*DegreeValue),

		Direction_stagedOrder: make(map[*Direction]uint),
		Direction_orderStaged: make(map[uint]*Direction),
		Directions_reference:  make(map[*Direction]*Direction),

		DirectionType_stagedOrder: make(map[*DirectionType]uint),
		DirectionType_orderStaged: make(map[uint]*DirectionType),
		DirectionTypes_reference:  make(map[*DirectionType]*DirectionType),

		Effect_stagedOrder: make(map[*Effect]uint),
		Effect_orderStaged: make(map[uint]*Effect),
		Effects_reference:  make(map[*Effect]*Effect),

		Feature_stagedOrder: make(map[*Feature]uint),
		Feature_orderStaged: make(map[uint]*Feature),
		Features_reference:  make(map[*Feature]*Feature),

		FirstFret_stagedOrder: make(map[*FirstFret]uint),
		FirstFret_orderStaged: make(map[uint]*FirstFret),
		FirstFrets_reference:  make(map[*FirstFret]*FirstFret),

		Frame_stagedOrder: make(map[*Frame]uint),
		Frame_orderStaged: make(map[uint]*Frame),
		Frames_reference:  make(map[*Frame]*Frame),

		FrameNote_stagedOrder: make(map[*FrameNote]uint),
		FrameNote_orderStaged: make(map[uint]*FrameNote),
		FrameNotes_reference:  make(map[*FrameNote]*FrameNote),

		Glass_stagedOrder: make(map[*Glass]uint),
		Glass_orderStaged: make(map[uint]*Glass),
		Glasss_reference:  make(map[*Glass]*Glass),

		Grouping_stagedOrder: make(map[*Grouping]uint),
		Grouping_orderStaged: make(map[uint]*Grouping),
		Groupings_reference:  make(map[*Grouping]*Grouping),

		Harmony_stagedOrder: make(map[*Harmony]uint),
		Harmony_orderStaged: make(map[uint]*Harmony),
		Harmonys_reference:  make(map[*Harmony]*Harmony),

		HarmonyAlter_stagedOrder: make(map[*HarmonyAlter]uint),
		HarmonyAlter_orderStaged: make(map[uint]*HarmonyAlter),
		HarmonyAlters_reference:  make(map[*HarmonyAlter]*HarmonyAlter),

		HarpPedals_stagedOrder: make(map[*HarpPedals]uint),
		HarpPedals_orderStaged: make(map[uint]*HarpPedals),
		HarpPedalss_reference:  make(map[*HarpPedals]*HarpPedals),

		Image_stagedOrder: make(map[*Image]uint),
		Image_orderStaged: make(map[uint]*Image),
		Images_reference:  make(map[*Image]*Image),

		InstrumentChange_stagedOrder: make(map[*InstrumentChange]uint),
		InstrumentChange_orderStaged: make(map[uint]*InstrumentChange),
		InstrumentChanges_reference:  make(map[*InstrumentChange]*InstrumentChange),

		Inversion_stagedOrder: make(map[*Inversion]uint),
		Inversion_orderStaged: make(map[uint]*Inversion),
		Inversions_reference:  make(map[*Inversion]*Inversion),

		Kind_stagedOrder: make(map[*Kind]uint),
		Kind_orderStaged: make(map[uint]*Kind),
		Kinds_reference:  make(map[*Kind]*Kind),

		Listening_stagedOrder: make(map[*Listening]uint),
		Listening_orderStaged: make(map[uint]*Listening),
		Listenings_reference:  make(map[*Listening]*Listening),

		MeasureNumbering_stagedOrder: make(map[*MeasureNumbering]uint),
		MeasureNumbering_orderStaged: make(map[uint]*MeasureNumbering),
		MeasureNumberings_reference:  make(map[*MeasureNumbering]*MeasureNumbering),

		Membrane_stagedOrder: make(map[*Membrane]uint),
		Membrane_orderStaged: make(map[uint]*Membrane),
		Membranes_reference:  make(map[*Membrane]*Membrane),

		Metal_stagedOrder: make(map[*Metal]uint),
		Metal_orderStaged: make(map[uint]*Metal),
		Metals_reference:  make(map[*Metal]*Metal),

		Metronome_stagedOrder: make(map[*Metronome]uint),
		Metronome_orderStaged: make(map[uint]*Metronome),
		Metronomes_reference:  make(map[*Metronome]*Metronome),

		MetronomeBeam_stagedOrder: make(map[*MetronomeBeam]uint),
		MetronomeBeam_orderStaged: make(map[uint]*MetronomeBeam),
		MetronomeBeams_reference:  make(map[*MetronomeBeam]*MetronomeBeam),

		MetronomeNote_stagedOrder: make(map[*MetronomeNote]uint),
		MetronomeNote_orderStaged: make(map[uint]*MetronomeNote),
		MetronomeNotes_reference:  make(map[*MetronomeNote]*MetronomeNote),

		MetronomeTied_stagedOrder: make(map[*MetronomeTied]uint),
		MetronomeTied_orderStaged: make(map[uint]*MetronomeTied),
		MetronomeTieds_reference:  make(map[*MetronomeTied]*MetronomeTied),

		MetronomeTuplet_stagedOrder: make(map[*MetronomeTuplet]uint),
		MetronomeTuplet_orderStaged: make(map[uint]*MetronomeTuplet),
		MetronomeTuplets_reference:  make(map[*MetronomeTuplet]*MetronomeTuplet),

		Numeral_stagedOrder: make(map[*Numeral]uint),
		Numeral_orderStaged: make(map[uint]*Numeral),
		Numerals_reference:  make(map[*Numeral]*Numeral),

		NumeralKey_stagedOrder: make(map[*NumeralKey]uint),
		NumeralKey_orderStaged: make(map[uint]*NumeralKey),
		NumeralKeys_reference:  make(map[*NumeralKey]*NumeralKey),

		NumeralRoot_stagedOrder: make(map[*NumeralRoot]uint),
		NumeralRoot_orderStaged: make(map[uint]*NumeralRoot),
		NumeralRoots_reference:  make(map[*NumeralRoot]*NumeralRoot),

		OctaveShift_stagedOrder: make(map[*OctaveShift]uint),
		OctaveShift_orderStaged: make(map[uint]*OctaveShift),
		OctaveShifts_reference:  make(map[*OctaveShift]*OctaveShift),

		Offset_stagedOrder: make(map[*Offset]uint),
		Offset_orderStaged: make(map[uint]*Offset),
		Offsets_reference:  make(map[*Offset]*Offset),

		OtherDirection_stagedOrder: make(map[*OtherDirection]uint),
		OtherDirection_orderStaged: make(map[uint]*OtherDirection),
		OtherDirections_reference:  make(map[*OtherDirection]*OtherDirection),

		OtherListening_stagedOrder: make(map[*OtherListening]uint),
		OtherListening_orderStaged: make(map[uint]*OtherListening),
		OtherListenings_reference:  make(map[*OtherListening]*OtherListening),

		Pedal_stagedOrder: make(map[*Pedal]uint),
		Pedal_orderStaged: make(map[uint]*Pedal),
		Pedals_reference:  make(map[*Pedal]*Pedal),

		PedalTuning_stagedOrder: make(map[*PedalTuning]uint),
		PedalTuning_orderStaged: make(map[uint]*PedalTuning),
		PedalTunings_reference:  make(map[*PedalTuning]*PedalTuning),

		PerMinute_stagedOrder: make(map[*PerMinute]uint),
		PerMinute_orderStaged: make(map[uint]*PerMinute),
		PerMinutes_reference:  make(map[*PerMinute]*PerMinute),

		Percussion_stagedOrder: make(map[*Percussion]uint),
		Percussion_orderStaged: make(map[uint]*Percussion),
		Percussions_reference:  make(map[*Percussion]*Percussion),

		Pitched_stagedOrder: make(map[*Pitched]uint),
		Pitched_orderStaged: make(map[uint]*Pitched),
		Pitcheds_reference:  make(map[*Pitched]*Pitched),

		PrincipalVoice_stagedOrder: make(map[*PrincipalVoice]uint),
		PrincipalVoice_orderStaged: make(map[uint]*PrincipalVoice),
		PrincipalVoices_reference:  make(map[*PrincipalVoice]*PrincipalVoice),

		Print_stagedOrder: make(map[*Print]uint),
		Print_orderStaged: make(map[uint]*Print),
		Prints_reference:  make(map[*Print]*Print),

		Root_stagedOrder: make(map[*Root]uint),
		Root_orderStaged: make(map[uint]*Root),
		Roots_reference:  make(map[*Root]*Root),

		RootStep_stagedOrder: make(map[*RootStep]uint),
		RootStep_orderStaged: make(map[uint]*RootStep),
		RootSteps_reference:  make(map[*RootStep]*RootStep),

		Scordatura_stagedOrder: make(map[*Scordatura]uint),
		Scordatura_orderStaged: make(map[uint]*Scordatura),
		Scordaturas_reference:  make(map[*Scordatura]*Scordatura),

		Sound_stagedOrder: make(map[*Sound]uint),
		Sound_orderStaged: make(map[uint]*Sound),
		Sounds_reference:  make(map[*Sound]*Sound),

		StaffDivide_stagedOrder: make(map[*StaffDivide]uint),
		StaffDivide_orderStaged: make(map[uint]*StaffDivide),
		StaffDivides_reference:  make(map[*StaffDivide]*StaffDivide),

		Stick_stagedOrder: make(map[*Stick]uint),
		Stick_orderStaged: make(map[uint]*Stick),
		Sticks_reference:  make(map[*Stick]*Stick),

		StringMute_stagedOrder: make(map[*StringMute]uint),
		StringMute_orderStaged: make(map[uint]*StringMute),
		StringMutes_reference:  make(map[*StringMute]*StringMute),

		StyleText_stagedOrder: make(map[*StyleText]uint),
		StyleText_orderStaged: make(map[uint]*StyleText),
		StyleTexts_reference:  make(map[*StyleText]*StyleText),

		Swing_stagedOrder: make(map[*Swing]uint),
		Swing_orderStaged: make(map[uint]*Swing),
		Swings_reference:  make(map[*Swing]*Swing),

		Sync_stagedOrder: make(map[*Sync]uint),
		Sync_orderStaged: make(map[uint]*Sync),
		Syncs_reference:  make(map[*Sync]*Sync),

		TimeModification_stagedOrder: make(map[*TimeModification]uint),
		TimeModification_orderStaged: make(map[uint]*TimeModification),
		TimeModifications_reference:  make(map[*TimeModification]*TimeModification),

		Timpani_stagedOrder: make(map[*Timpani]uint),
		Timpani_orderStaged: make(map[uint]*Timpani),
		Timpanis_reference:  make(map[*Timpani]*Timpani),

		VirtualInstrument_stagedOrder: make(map[*VirtualInstrument]uint),
		VirtualInstrument_orderStaged: make(map[uint]*VirtualInstrument),
		VirtualInstruments_reference:  make(map[*VirtualInstrument]*VirtualInstrument),

		Wedge_stagedOrder: make(map[*Wedge]uint),
		Wedge_orderStaged: make(map[uint]*Wedge),
		Wedges_reference:  make(map[*Wedge]*Wedge),

		Wood_stagedOrder: make(map[*Wood]uint),
		Wood_orderStaged: make(map[uint]*Wood),
		Woods_reference:  make(map[*Wood]*Wood),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Accord": &AccordUnmarshaller{},

			"AccordionRegistration": &AccordionRegistrationUnmarshaller{},

			"Barre": &BarreUnmarshaller{},

			"Bass": &BassUnmarshaller{},

			"BassStep": &BassStepUnmarshaller{},

			"BeatUnitTied": &BeatUnitTiedUnmarshaller{},

			"Beater": &BeaterUnmarshaller{},

			"Bracket": &BracketUnmarshaller{},

			"Dashes": &DashesUnmarshaller{},

			"Degree": &DegreeUnmarshaller{},

			"DegreeAlter": &DegreeAlterUnmarshaller{},

			"DegreeType": &DegreeTypeUnmarshaller{},

			"DegreeValue": &DegreeValueUnmarshaller{},

			"Direction": &DirectionUnmarshaller{},

			"DirectionType": &DirectionTypeUnmarshaller{},

			"Effect": &EffectUnmarshaller{},

			"Feature": &FeatureUnmarshaller{},

			"FirstFret": &FirstFretUnmarshaller{},

			"Frame": &FrameUnmarshaller{},

			"FrameNote": &FrameNoteUnmarshaller{},

			"Glass": &GlassUnmarshaller{},

			"Grouping": &GroupingUnmarshaller{},

			"Harmony": &HarmonyUnmarshaller{},

			"HarmonyAlter": &HarmonyAlterUnmarshaller{},

			"HarpPedals": &HarpPedalsUnmarshaller{},

			"Image": &ImageUnmarshaller{},

			"InstrumentChange": &InstrumentChangeUnmarshaller{},

			"Inversion": &InversionUnmarshaller{},

			"Kind": &KindUnmarshaller{},

			"Listening": &ListeningUnmarshaller{},

			"MeasureNumbering": &MeasureNumberingUnmarshaller{},

			"Membrane": &MembraneUnmarshaller{},

			"Metal": &MetalUnmarshaller{},

			"Metronome": &MetronomeUnmarshaller{},

			"MetronomeBeam": &MetronomeBeamUnmarshaller{},

			"MetronomeNote": &MetronomeNoteUnmarshaller{},

			"MetronomeTied": &MetronomeTiedUnmarshaller{},

			"MetronomeTuplet": &MetronomeTupletUnmarshaller{},

			"Numeral": &NumeralUnmarshaller{},

			"NumeralKey": &NumeralKeyUnmarshaller{},

			"NumeralRoot": &NumeralRootUnmarshaller{},

			"OctaveShift": &OctaveShiftUnmarshaller{},

			"Offset": &OffsetUnmarshaller{},

			"OtherDirection": &OtherDirectionUnmarshaller{},

			"OtherListening": &OtherListeningUnmarshaller{},

			"Pedal": &PedalUnmarshaller{},

			"PedalTuning": &PedalTuningUnmarshaller{},

			"PerMinute": &PerMinuteUnmarshaller{},

			"Percussion": &PercussionUnmarshaller{},

			"Pitched": &PitchedUnmarshaller{},

			"PrincipalVoice": &PrincipalVoiceUnmarshaller{},

			"Print": &PrintUnmarshaller{},

			"Root": &RootUnmarshaller{},

			"RootStep": &RootStepUnmarshaller{},

			"Scordatura": &ScordaturaUnmarshaller{},

			"Sound": &SoundUnmarshaller{},

			"StaffDivide": &StaffDivideUnmarshaller{},

			"Stick": &StickUnmarshaller{},

			"StringMute": &StringMuteUnmarshaller{},

			"StyleText": &StyleTextUnmarshaller{},

			"Swing": &SwingUnmarshaller{},

			"Sync": &SyncUnmarshaller{},

			"TimeModification": &TimeModificationUnmarshaller{},

			"Timpani": &TimpaniUnmarshaller{},

			"VirtualInstrument": &VirtualInstrumentUnmarshaller{},

			"Wedge": &WedgeUnmarshaller{},

			"Wood": &WoodUnmarshaller{},

			// end of insertion point
		},
	}
	stage.StageCore.SetName(name)
	stage.StageCore.SetNavigationMode(GongNavigationModeNormal)

	return
}

// GetOrder is the Stage method returning the order of a gongstruct instance.
func (stage *Stage) GetOrder(instance GongstructIF) uint {
	if instance != nil {
		return instance.GongGetOrder(stage)
	}
	return 0
}

// GetInstanceFromOrder is the Stage method returning a gongstruct instance from its order.
func (stage *Stage) GetInstanceFromOrder[Type GongstructPtr](order uint) (res Type) {
	if stage == nil {
		return
	}
	var t Type
	val := t.GongGetInstanceFromOrder(stage, order)
	if val != nil {
		res = val.(Type)
	}
	return
}

func (stage *Stage) CommitWithSuspendedCallbacks() {
	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	tmp2 := stage.StageCore.GetBeforeCommitHooks()
	stage.StageCore.ClearBeforeCommitHooks()
	tmp3 := stage.StageCore.GetAfterCommitHooks()
	stage.StageCore.ClearAfterCommitHooks()
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.StageCore.SetBeforeCommitHooks(tmp2)
	stage.StageCore.SetAfterCommitHooks(tmp3)
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.IsInDeltaMode() && stage.GetNavigationMode() == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
		stage.ResetHard()
	}

	if stage.IsInDeltaMode() {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
		if stage.probeIF != nil {
			stage.probeIF.RefreshNavigationTree()
		}
	}

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Accord"] = len(stage.Accords)
	stage.Map_GongStructName_InstancesNb["AccordionRegistration"] = len(stage.AccordionRegistrations)
	stage.Map_GongStructName_InstancesNb["Barre"] = len(stage.Barres)
	stage.Map_GongStructName_InstancesNb["Bass"] = len(stage.Basss)
	stage.Map_GongStructName_InstancesNb["BassStep"] = len(stage.BassSteps)
	stage.Map_GongStructName_InstancesNb["BeatUnitTied"] = len(stage.BeatUnitTieds)
	stage.Map_GongStructName_InstancesNb["Beater"] = len(stage.Beaters)
	stage.Map_GongStructName_InstancesNb["Bracket"] = len(stage.Brackets)
	stage.Map_GongStructName_InstancesNb["Dashes"] = len(stage.Dashess)
	stage.Map_GongStructName_InstancesNb["Degree"] = len(stage.Degrees)
	stage.Map_GongStructName_InstancesNb["DegreeAlter"] = len(stage.DegreeAlters)
	stage.Map_GongStructName_InstancesNb["DegreeType"] = len(stage.DegreeTypes)
	stage.Map_GongStructName_InstancesNb["DegreeValue"] = len(stage.DegreeValues)
	stage.Map_GongStructName_InstancesNb["Direction"] = len(stage.Directions)
	stage.Map_GongStructName_InstancesNb["DirectionType"] = len(stage.DirectionTypes)
	stage.Map_GongStructName_InstancesNb["Effect"] = len(stage.Effects)
	stage.Map_GongStructName_InstancesNb["Feature"] = len(stage.Features)
	stage.Map_GongStructName_InstancesNb["FirstFret"] = len(stage.FirstFrets)
	stage.Map_GongStructName_InstancesNb["Frame"] = len(stage.Frames)
	stage.Map_GongStructName_InstancesNb["FrameNote"] = len(stage.FrameNotes)
	stage.Map_GongStructName_InstancesNb["Glass"] = len(stage.Glasss)
	stage.Map_GongStructName_InstancesNb["Grouping"] = len(stage.Groupings)
	stage.Map_GongStructName_InstancesNb["Harmony"] = len(stage.Harmonys)
	stage.Map_GongStructName_InstancesNb["HarmonyAlter"] = len(stage.HarmonyAlters)
	stage.Map_GongStructName_InstancesNb["HarpPedals"] = len(stage.HarpPedalss)
	stage.Map_GongStructName_InstancesNb["Image"] = len(stage.Images)
	stage.Map_GongStructName_InstancesNb["InstrumentChange"] = len(stage.InstrumentChanges)
	stage.Map_GongStructName_InstancesNb["Inversion"] = len(stage.Inversions)
	stage.Map_GongStructName_InstancesNb["Kind"] = len(stage.Kinds)
	stage.Map_GongStructName_InstancesNb["Listening"] = len(stage.Listenings)
	stage.Map_GongStructName_InstancesNb["MeasureNumbering"] = len(stage.MeasureNumberings)
	stage.Map_GongStructName_InstancesNb["Membrane"] = len(stage.Membranes)
	stage.Map_GongStructName_InstancesNb["Metal"] = len(stage.Metals)
	stage.Map_GongStructName_InstancesNb["Metronome"] = len(stage.Metronomes)
	stage.Map_GongStructName_InstancesNb["MetronomeBeam"] = len(stage.MetronomeBeams)
	stage.Map_GongStructName_InstancesNb["MetronomeNote"] = len(stage.MetronomeNotes)
	stage.Map_GongStructName_InstancesNb["MetronomeTied"] = len(stage.MetronomeTieds)
	stage.Map_GongStructName_InstancesNb["MetronomeTuplet"] = len(stage.MetronomeTuplets)
	stage.Map_GongStructName_InstancesNb["Numeral"] = len(stage.Numerals)
	stage.Map_GongStructName_InstancesNb["NumeralKey"] = len(stage.NumeralKeys)
	stage.Map_GongStructName_InstancesNb["NumeralRoot"] = len(stage.NumeralRoots)
	stage.Map_GongStructName_InstancesNb["OctaveShift"] = len(stage.OctaveShifts)
	stage.Map_GongStructName_InstancesNb["Offset"] = len(stage.Offsets)
	stage.Map_GongStructName_InstancesNb["OtherDirection"] = len(stage.OtherDirections)
	stage.Map_GongStructName_InstancesNb["OtherListening"] = len(stage.OtherListenings)
	stage.Map_GongStructName_InstancesNb["Pedal"] = len(stage.Pedals)
	stage.Map_GongStructName_InstancesNb["PedalTuning"] = len(stage.PedalTunings)
	stage.Map_GongStructName_InstancesNb["PerMinute"] = len(stage.PerMinutes)
	stage.Map_GongStructName_InstancesNb["Percussion"] = len(stage.Percussions)
	stage.Map_GongStructName_InstancesNb["Pitched"] = len(stage.Pitcheds)
	stage.Map_GongStructName_InstancesNb["PrincipalVoice"] = len(stage.PrincipalVoices)
	stage.Map_GongStructName_InstancesNb["Print"] = len(stage.Prints)
	stage.Map_GongStructName_InstancesNb["Root"] = len(stage.Roots)
	stage.Map_GongStructName_InstancesNb["RootStep"] = len(stage.RootSteps)
	stage.Map_GongStructName_InstancesNb["Scordatura"] = len(stage.Scordaturas)
	stage.Map_GongStructName_InstancesNb["Sound"] = len(stage.Sounds)
	stage.Map_GongStructName_InstancesNb["StaffDivide"] = len(stage.StaffDivides)
	stage.Map_GongStructName_InstancesNb["Stick"] = len(stage.Sticks)
	stage.Map_GongStructName_InstancesNb["StringMute"] = len(stage.StringMutes)
	stage.Map_GongStructName_InstancesNb["StyleText"] = len(stage.StyleTexts)
	stage.Map_GongStructName_InstancesNb["Swing"] = len(stage.Swings)
	stage.Map_GongStructName_InstancesNb["Sync"] = len(stage.Syncs)
	stage.Map_GongStructName_InstancesNb["TimeModification"] = len(stage.TimeModifications)
	stage.Map_GongStructName_InstancesNb["Timpani"] = len(stage.Timpanis)
	stage.Map_GongStructName_InstancesNb["VirtualInstrument"] = len(stage.VirtualInstruments)
	stage.Map_GongStructName_InstancesNb["Wedge"] = len(stage.Wedges)
	stage.Map_GongStructName_InstancesNb["Wood"] = len(stage.Woods)
}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
}

// backup generates backup files in the dirPath
func (stage *Stage) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *Stage) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts accord to the model stage
func (accord *Accord) Stage(stage *Stage) *Accord {
	__gong__stage(stage.Accords, stage.Accord_stagedOrder, stage.Accord_orderStaged, &stage.AccordOrder, stage.Accords_mapString, accord, accord.Name)
	return accord
}

// StagePreserveOrder puts accord to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AccordOrder
// - update stage.AccordOrder accordingly
func (accord *Accord) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Accords, stage.Accord_stagedOrder, stage.Accord_orderStaged, &stage.AccordOrder, stage.Accords_mapString, accord, order, accord.Name)
}

// Unstage removes accord off the model stage
func (accord *Accord) Unstage(stage *Stage) *Accord {
	__gong__unstage(stage.Accords, stage.Accords_mapString, accord, accord.Name)
	return accord
}

// UnstageVoid removes accord off the model stage
func (accord *Accord) UnstageVoid(stage *Stage) {
	accord.Unstage(stage)
}

func (accord *Accord) StageVoid(stage *Stage) {
	accord.Stage(stage)
}

// for satisfaction of GongStruct interface
func (accord *Accord) GetName() (res string) {
	return accord.Name
}

// for satisfaction of GongStruct interface
func (accord *Accord) SetName(name string) {
	accord.Name = name
}

// Stage puts accordionregistration to the model stage
func (accordionregistration *AccordionRegistration) Stage(stage *Stage) *AccordionRegistration {
	__gong__stage(stage.AccordionRegistrations, stage.AccordionRegistration_stagedOrder, stage.AccordionRegistration_orderStaged, &stage.AccordionRegistrationOrder, stage.AccordionRegistrations_mapString, accordionregistration, accordionregistration.Name)
	return accordionregistration
}

// StagePreserveOrder puts accordionregistration to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AccordionRegistrationOrder
// - update stage.AccordionRegistrationOrder accordingly
func (accordionregistration *AccordionRegistration) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AccordionRegistrations, stage.AccordionRegistration_stagedOrder, stage.AccordionRegistration_orderStaged, &stage.AccordionRegistrationOrder, stage.AccordionRegistrations_mapString, accordionregistration, order, accordionregistration.Name)
}

// Unstage removes accordionregistration off the model stage
func (accordionregistration *AccordionRegistration) Unstage(stage *Stage) *AccordionRegistration {
	__gong__unstage(stage.AccordionRegistrations, stage.AccordionRegistrations_mapString, accordionregistration, accordionregistration.Name)
	return accordionregistration
}

// UnstageVoid removes accordionregistration off the model stage
func (accordionregistration *AccordionRegistration) UnstageVoid(stage *Stage) {
	accordionregistration.Unstage(stage)
}

func (accordionregistration *AccordionRegistration) StageVoid(stage *Stage) {
	accordionregistration.Stage(stage)
}

// for satisfaction of GongStruct interface
func (accordionregistration *AccordionRegistration) GetName() (res string) {
	return accordionregistration.Name
}

// for satisfaction of GongStruct interface
func (accordionregistration *AccordionRegistration) SetName(name string) {
	accordionregistration.Name = name
}

// Stage puts barre to the model stage
func (barre *Barre) Stage(stage *Stage) *Barre {
	__gong__stage(stage.Barres, stage.Barre_stagedOrder, stage.Barre_orderStaged, &stage.BarreOrder, stage.Barres_mapString, barre, barre.Name)
	return barre
}

// StagePreserveOrder puts barre to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BarreOrder
// - update stage.BarreOrder accordingly
func (barre *Barre) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Barres, stage.Barre_stagedOrder, stage.Barre_orderStaged, &stage.BarreOrder, stage.Barres_mapString, barre, order, barre.Name)
}

// Unstage removes barre off the model stage
func (barre *Barre) Unstage(stage *Stage) *Barre {
	__gong__unstage(stage.Barres, stage.Barres_mapString, barre, barre.Name)
	return barre
}

// UnstageVoid removes barre off the model stage
func (barre *Barre) UnstageVoid(stage *Stage) {
	barre.Unstage(stage)
}

func (barre *Barre) StageVoid(stage *Stage) {
	barre.Stage(stage)
}

// for satisfaction of GongStruct interface
func (barre *Barre) GetName() (res string) {
	return barre.Name
}

// for satisfaction of GongStruct interface
func (barre *Barre) SetName(name string) {
	barre.Name = name
}

// Stage puts bass to the model stage
func (bass *Bass) Stage(stage *Stage) *Bass {
	__gong__stage(stage.Basss, stage.Bass_stagedOrder, stage.Bass_orderStaged, &stage.BassOrder, stage.Basss_mapString, bass, bass.Name)
	return bass
}

// StagePreserveOrder puts bass to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BassOrder
// - update stage.BassOrder accordingly
func (bass *Bass) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Basss, stage.Bass_stagedOrder, stage.Bass_orderStaged, &stage.BassOrder, stage.Basss_mapString, bass, order, bass.Name)
}

// Unstage removes bass off the model stage
func (bass *Bass) Unstage(stage *Stage) *Bass {
	__gong__unstage(stage.Basss, stage.Basss_mapString, bass, bass.Name)
	return bass
}

// UnstageVoid removes bass off the model stage
func (bass *Bass) UnstageVoid(stage *Stage) {
	bass.Unstage(stage)
}

func (bass *Bass) StageVoid(stage *Stage) {
	bass.Stage(stage)
}

// for satisfaction of GongStruct interface
func (bass *Bass) GetName() (res string) {
	return bass.Name
}

// for satisfaction of GongStruct interface
func (bass *Bass) SetName(name string) {
	bass.Name = name
}

// Stage puts bassstep to the model stage
func (bassstep *BassStep) Stage(stage *Stage) *BassStep {
	__gong__stage(stage.BassSteps, stage.BassStep_stagedOrder, stage.BassStep_orderStaged, &stage.BassStepOrder, stage.BassSteps_mapString, bassstep, bassstep.Name)
	return bassstep
}

// StagePreserveOrder puts bassstep to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BassStepOrder
// - update stage.BassStepOrder accordingly
func (bassstep *BassStep) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.BassSteps, stage.BassStep_stagedOrder, stage.BassStep_orderStaged, &stage.BassStepOrder, stage.BassSteps_mapString, bassstep, order, bassstep.Name)
}

// Unstage removes bassstep off the model stage
func (bassstep *BassStep) Unstage(stage *Stage) *BassStep {
	__gong__unstage(stage.BassSteps, stage.BassSteps_mapString, bassstep, bassstep.Name)
	return bassstep
}

// UnstageVoid removes bassstep off the model stage
func (bassstep *BassStep) UnstageVoid(stage *Stage) {
	bassstep.Unstage(stage)
}

func (bassstep *BassStep) StageVoid(stage *Stage) {
	bassstep.Stage(stage)
}

// for satisfaction of GongStruct interface
func (bassstep *BassStep) GetName() (res string) {
	return bassstep.Name
}

// for satisfaction of GongStruct interface
func (bassstep *BassStep) SetName(name string) {
	bassstep.Name = name
}

// Stage puts beatunittied to the model stage
func (beatunittied *BeatUnitTied) Stage(stage *Stage) *BeatUnitTied {
	__gong__stage(stage.BeatUnitTieds, stage.BeatUnitTied_stagedOrder, stage.BeatUnitTied_orderStaged, &stage.BeatUnitTiedOrder, stage.BeatUnitTieds_mapString, beatunittied, beatunittied.Name)
	return beatunittied
}

// StagePreserveOrder puts beatunittied to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BeatUnitTiedOrder
// - update stage.BeatUnitTiedOrder accordingly
func (beatunittied *BeatUnitTied) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.BeatUnitTieds, stage.BeatUnitTied_stagedOrder, stage.BeatUnitTied_orderStaged, &stage.BeatUnitTiedOrder, stage.BeatUnitTieds_mapString, beatunittied, order, beatunittied.Name)
}

// Unstage removes beatunittied off the model stage
func (beatunittied *BeatUnitTied) Unstage(stage *Stage) *BeatUnitTied {
	__gong__unstage(stage.BeatUnitTieds, stage.BeatUnitTieds_mapString, beatunittied, beatunittied.Name)
	return beatunittied
}

// UnstageVoid removes beatunittied off the model stage
func (beatunittied *BeatUnitTied) UnstageVoid(stage *Stage) {
	beatunittied.Unstage(stage)
}

func (beatunittied *BeatUnitTied) StageVoid(stage *Stage) {
	beatunittied.Stage(stage)
}

// for satisfaction of GongStruct interface
func (beatunittied *BeatUnitTied) GetName() (res string) {
	return beatunittied.Name
}

// for satisfaction of GongStruct interface
func (beatunittied *BeatUnitTied) SetName(name string) {
	beatunittied.Name = name
}

// Stage puts beater to the model stage
func (beater *Beater) Stage(stage *Stage) *Beater {
	__gong__stage(stage.Beaters, stage.Beater_stagedOrder, stage.Beater_orderStaged, &stage.BeaterOrder, stage.Beaters_mapString, beater, beater.Name)
	return beater
}

// StagePreserveOrder puts beater to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BeaterOrder
// - update stage.BeaterOrder accordingly
func (beater *Beater) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Beaters, stage.Beater_stagedOrder, stage.Beater_orderStaged, &stage.BeaterOrder, stage.Beaters_mapString, beater, order, beater.Name)
}

// Unstage removes beater off the model stage
func (beater *Beater) Unstage(stage *Stage) *Beater {
	__gong__unstage(stage.Beaters, stage.Beaters_mapString, beater, beater.Name)
	return beater
}

// UnstageVoid removes beater off the model stage
func (beater *Beater) UnstageVoid(stage *Stage) {
	beater.Unstage(stage)
}

func (beater *Beater) StageVoid(stage *Stage) {
	beater.Stage(stage)
}

// for satisfaction of GongStruct interface
func (beater *Beater) GetName() (res string) {
	return beater.Name
}

// for satisfaction of GongStruct interface
func (beater *Beater) SetName(name string) {
	beater.Name = name
}

// Stage puts bracket to the model stage
func (bracket *Bracket) Stage(stage *Stage) *Bracket {
	__gong__stage(stage.Brackets, stage.Bracket_stagedOrder, stage.Bracket_orderStaged, &stage.BracketOrder, stage.Brackets_mapString, bracket, bracket.Name)
	return bracket
}

// StagePreserveOrder puts bracket to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BracketOrder
// - update stage.BracketOrder accordingly
func (bracket *Bracket) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Brackets, stage.Bracket_stagedOrder, stage.Bracket_orderStaged, &stage.BracketOrder, stage.Brackets_mapString, bracket, order, bracket.Name)
}

// Unstage removes bracket off the model stage
func (bracket *Bracket) Unstage(stage *Stage) *Bracket {
	__gong__unstage(stage.Brackets, stage.Brackets_mapString, bracket, bracket.Name)
	return bracket
}

// UnstageVoid removes bracket off the model stage
func (bracket *Bracket) UnstageVoid(stage *Stage) {
	bracket.Unstage(stage)
}

func (bracket *Bracket) StageVoid(stage *Stage) {
	bracket.Stage(stage)
}

// for satisfaction of GongStruct interface
func (bracket *Bracket) GetName() (res string) {
	return bracket.Name
}

// for satisfaction of GongStruct interface
func (bracket *Bracket) SetName(name string) {
	bracket.Name = name
}

// Stage puts dashes to the model stage
func (dashes *Dashes) Stage(stage *Stage) *Dashes {
	__gong__stage(stage.Dashess, stage.Dashes_stagedOrder, stage.Dashes_orderStaged, &stage.DashesOrder, stage.Dashess_mapString, dashes, dashes.Name)
	return dashes
}

// StagePreserveOrder puts dashes to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DashesOrder
// - update stage.DashesOrder accordingly
func (dashes *Dashes) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Dashess, stage.Dashes_stagedOrder, stage.Dashes_orderStaged, &stage.DashesOrder, stage.Dashess_mapString, dashes, order, dashes.Name)
}

// Unstage removes dashes off the model stage
func (dashes *Dashes) Unstage(stage *Stage) *Dashes {
	__gong__unstage(stage.Dashess, stage.Dashess_mapString, dashes, dashes.Name)
	return dashes
}

// UnstageVoid removes dashes off the model stage
func (dashes *Dashes) UnstageVoid(stage *Stage) {
	dashes.Unstage(stage)
}

func (dashes *Dashes) StageVoid(stage *Stage) {
	dashes.Stage(stage)
}

// for satisfaction of GongStruct interface
func (dashes *Dashes) GetName() (res string) {
	return dashes.Name
}

// for satisfaction of GongStruct interface
func (dashes *Dashes) SetName(name string) {
	dashes.Name = name
}

// Stage puts degree to the model stage
func (degree *Degree) Stage(stage *Stage) *Degree {
	__gong__stage(stage.Degrees, stage.Degree_stagedOrder, stage.Degree_orderStaged, &stage.DegreeOrder, stage.Degrees_mapString, degree, degree.Name)
	return degree
}

// StagePreserveOrder puts degree to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DegreeOrder
// - update stage.DegreeOrder accordingly
func (degree *Degree) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Degrees, stage.Degree_stagedOrder, stage.Degree_orderStaged, &stage.DegreeOrder, stage.Degrees_mapString, degree, order, degree.Name)
}

// Unstage removes degree off the model stage
func (degree *Degree) Unstage(stage *Stage) *Degree {
	__gong__unstage(stage.Degrees, stage.Degrees_mapString, degree, degree.Name)
	return degree
}

// UnstageVoid removes degree off the model stage
func (degree *Degree) UnstageVoid(stage *Stage) {
	degree.Unstage(stage)
}

func (degree *Degree) StageVoid(stage *Stage) {
	degree.Stage(stage)
}

// for satisfaction of GongStruct interface
func (degree *Degree) GetName() (res string) {
	return degree.Name
}

// for satisfaction of GongStruct interface
func (degree *Degree) SetName(name string) {
	degree.Name = name
}

// Stage puts degreealter to the model stage
func (degreealter *DegreeAlter) Stage(stage *Stage) *DegreeAlter {
	__gong__stage(stage.DegreeAlters, stage.DegreeAlter_stagedOrder, stage.DegreeAlter_orderStaged, &stage.DegreeAlterOrder, stage.DegreeAlters_mapString, degreealter, degreealter.Name)
	return degreealter
}

// StagePreserveOrder puts degreealter to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DegreeAlterOrder
// - update stage.DegreeAlterOrder accordingly
func (degreealter *DegreeAlter) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DegreeAlters, stage.DegreeAlter_stagedOrder, stage.DegreeAlter_orderStaged, &stage.DegreeAlterOrder, stage.DegreeAlters_mapString, degreealter, order, degreealter.Name)
}

// Unstage removes degreealter off the model stage
func (degreealter *DegreeAlter) Unstage(stage *Stage) *DegreeAlter {
	__gong__unstage(stage.DegreeAlters, stage.DegreeAlters_mapString, degreealter, degreealter.Name)
	return degreealter
}

// UnstageVoid removes degreealter off the model stage
func (degreealter *DegreeAlter) UnstageVoid(stage *Stage) {
	degreealter.Unstage(stage)
}

func (degreealter *DegreeAlter) StageVoid(stage *Stage) {
	degreealter.Stage(stage)
}

// for satisfaction of GongStruct interface
func (degreealter *DegreeAlter) GetName() (res string) {
	return degreealter.Name
}

// for satisfaction of GongStruct interface
func (degreealter *DegreeAlter) SetName(name string) {
	degreealter.Name = name
}

// Stage puts degreetype to the model stage
func (degreetype *DegreeType) Stage(stage *Stage) *DegreeType {
	__gong__stage(stage.DegreeTypes, stage.DegreeType_stagedOrder, stage.DegreeType_orderStaged, &stage.DegreeTypeOrder, stage.DegreeTypes_mapString, degreetype, degreetype.Name)
	return degreetype
}

// StagePreserveOrder puts degreetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DegreeTypeOrder
// - update stage.DegreeTypeOrder accordingly
func (degreetype *DegreeType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DegreeTypes, stage.DegreeType_stagedOrder, stage.DegreeType_orderStaged, &stage.DegreeTypeOrder, stage.DegreeTypes_mapString, degreetype, order, degreetype.Name)
}

// Unstage removes degreetype off the model stage
func (degreetype *DegreeType) Unstage(stage *Stage) *DegreeType {
	__gong__unstage(stage.DegreeTypes, stage.DegreeTypes_mapString, degreetype, degreetype.Name)
	return degreetype
}

// UnstageVoid removes degreetype off the model stage
func (degreetype *DegreeType) UnstageVoid(stage *Stage) {
	degreetype.Unstage(stage)
}

func (degreetype *DegreeType) StageVoid(stage *Stage) {
	degreetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (degreetype *DegreeType) GetName() (res string) {
	return degreetype.Name
}

// for satisfaction of GongStruct interface
func (degreetype *DegreeType) SetName(name string) {
	degreetype.Name = name
}

// Stage puts degreevalue to the model stage
func (degreevalue *DegreeValue) Stage(stage *Stage) *DegreeValue {
	__gong__stage(stage.DegreeValues, stage.DegreeValue_stagedOrder, stage.DegreeValue_orderStaged, &stage.DegreeValueOrder, stage.DegreeValues_mapString, degreevalue, degreevalue.Name)
	return degreevalue
}

// StagePreserveOrder puts degreevalue to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DegreeValueOrder
// - update stage.DegreeValueOrder accordingly
func (degreevalue *DegreeValue) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DegreeValues, stage.DegreeValue_stagedOrder, stage.DegreeValue_orderStaged, &stage.DegreeValueOrder, stage.DegreeValues_mapString, degreevalue, order, degreevalue.Name)
}

// Unstage removes degreevalue off the model stage
func (degreevalue *DegreeValue) Unstage(stage *Stage) *DegreeValue {
	__gong__unstage(stage.DegreeValues, stage.DegreeValues_mapString, degreevalue, degreevalue.Name)
	return degreevalue
}

// UnstageVoid removes degreevalue off the model stage
func (degreevalue *DegreeValue) UnstageVoid(stage *Stage) {
	degreevalue.Unstage(stage)
}

func (degreevalue *DegreeValue) StageVoid(stage *Stage) {
	degreevalue.Stage(stage)
}

// for satisfaction of GongStruct interface
func (degreevalue *DegreeValue) GetName() (res string) {
	return degreevalue.Name
}

// for satisfaction of GongStruct interface
func (degreevalue *DegreeValue) SetName(name string) {
	degreevalue.Name = name
}

// Stage puts direction to the model stage
func (direction *Direction) Stage(stage *Stage) *Direction {
	__gong__stage(stage.Directions, stage.Direction_stagedOrder, stage.Direction_orderStaged, &stage.DirectionOrder, stage.Directions_mapString, direction, direction.Name)
	return direction
}

// StagePreserveOrder puts direction to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DirectionOrder
// - update stage.DirectionOrder accordingly
func (direction *Direction) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Directions, stage.Direction_stagedOrder, stage.Direction_orderStaged, &stage.DirectionOrder, stage.Directions_mapString, direction, order, direction.Name)
}

// Unstage removes direction off the model stage
func (direction *Direction) Unstage(stage *Stage) *Direction {
	__gong__unstage(stage.Directions, stage.Directions_mapString, direction, direction.Name)
	return direction
}

// UnstageVoid removes direction off the model stage
func (direction *Direction) UnstageVoid(stage *Stage) {
	direction.Unstage(stage)
}

func (direction *Direction) StageVoid(stage *Stage) {
	direction.Stage(stage)
}

// for satisfaction of GongStruct interface
func (direction *Direction) GetName() (res string) {
	return direction.Name
}

// for satisfaction of GongStruct interface
func (direction *Direction) SetName(name string) {
	direction.Name = name
}

// Stage puts directiontype to the model stage
func (directiontype *DirectionType) Stage(stage *Stage) *DirectionType {
	__gong__stage(stage.DirectionTypes, stage.DirectionType_stagedOrder, stage.DirectionType_orderStaged, &stage.DirectionTypeOrder, stage.DirectionTypes_mapString, directiontype, directiontype.Name)
	return directiontype
}

// StagePreserveOrder puts directiontype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DirectionTypeOrder
// - update stage.DirectionTypeOrder accordingly
func (directiontype *DirectionType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DirectionTypes, stage.DirectionType_stagedOrder, stage.DirectionType_orderStaged, &stage.DirectionTypeOrder, stage.DirectionTypes_mapString, directiontype, order, directiontype.Name)
}

// Unstage removes directiontype off the model stage
func (directiontype *DirectionType) Unstage(stage *Stage) *DirectionType {
	__gong__unstage(stage.DirectionTypes, stage.DirectionTypes_mapString, directiontype, directiontype.Name)
	return directiontype
}

// UnstageVoid removes directiontype off the model stage
func (directiontype *DirectionType) UnstageVoid(stage *Stage) {
	directiontype.Unstage(stage)
}

func (directiontype *DirectionType) StageVoid(stage *Stage) {
	directiontype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (directiontype *DirectionType) GetName() (res string) {
	return directiontype.Name
}

// for satisfaction of GongStruct interface
func (directiontype *DirectionType) SetName(name string) {
	directiontype.Name = name
}

// Stage puts effect to the model stage
func (effect *Effect) Stage(stage *Stage) *Effect {
	__gong__stage(stage.Effects, stage.Effect_stagedOrder, stage.Effect_orderStaged, &stage.EffectOrder, stage.Effects_mapString, effect, effect.Name)
	return effect
}

// StagePreserveOrder puts effect to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EffectOrder
// - update stage.EffectOrder accordingly
func (effect *Effect) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Effects, stage.Effect_stagedOrder, stage.Effect_orderStaged, &stage.EffectOrder, stage.Effects_mapString, effect, order, effect.Name)
}

// Unstage removes effect off the model stage
func (effect *Effect) Unstage(stage *Stage) *Effect {
	__gong__unstage(stage.Effects, stage.Effects_mapString, effect, effect.Name)
	return effect
}

// UnstageVoid removes effect off the model stage
func (effect *Effect) UnstageVoid(stage *Stage) {
	effect.Unstage(stage)
}

func (effect *Effect) StageVoid(stage *Stage) {
	effect.Stage(stage)
}

// for satisfaction of GongStruct interface
func (effect *Effect) GetName() (res string) {
	return effect.Name
}

// for satisfaction of GongStruct interface
func (effect *Effect) SetName(name string) {
	effect.Name = name
}

// Stage puts feature to the model stage
func (feature *Feature) Stage(stage *Stage) *Feature {
	__gong__stage(stage.Features, stage.Feature_stagedOrder, stage.Feature_orderStaged, &stage.FeatureOrder, stage.Features_mapString, feature, feature.Name)
	return feature
}

// StagePreserveOrder puts feature to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FeatureOrder
// - update stage.FeatureOrder accordingly
func (feature *Feature) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Features, stage.Feature_stagedOrder, stage.Feature_orderStaged, &stage.FeatureOrder, stage.Features_mapString, feature, order, feature.Name)
}

// Unstage removes feature off the model stage
func (feature *Feature) Unstage(stage *Stage) *Feature {
	__gong__unstage(stage.Features, stage.Features_mapString, feature, feature.Name)
	return feature
}

// UnstageVoid removes feature off the model stage
func (feature *Feature) UnstageVoid(stage *Stage) {
	feature.Unstage(stage)
}

func (feature *Feature) StageVoid(stage *Stage) {
	feature.Stage(stage)
}

// for satisfaction of GongStruct interface
func (feature *Feature) GetName() (res string) {
	return feature.Name
}

// for satisfaction of GongStruct interface
func (feature *Feature) SetName(name string) {
	feature.Name = name
}

// Stage puts firstfret to the model stage
func (firstfret *FirstFret) Stage(stage *Stage) *FirstFret {
	__gong__stage(stage.FirstFrets, stage.FirstFret_stagedOrder, stage.FirstFret_orderStaged, &stage.FirstFretOrder, stage.FirstFrets_mapString, firstfret, firstfret.Name)
	return firstfret
}

// StagePreserveOrder puts firstfret to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FirstFretOrder
// - update stage.FirstFretOrder accordingly
func (firstfret *FirstFret) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FirstFrets, stage.FirstFret_stagedOrder, stage.FirstFret_orderStaged, &stage.FirstFretOrder, stage.FirstFrets_mapString, firstfret, order, firstfret.Name)
}

// Unstage removes firstfret off the model stage
func (firstfret *FirstFret) Unstage(stage *Stage) *FirstFret {
	__gong__unstage(stage.FirstFrets, stage.FirstFrets_mapString, firstfret, firstfret.Name)
	return firstfret
}

// UnstageVoid removes firstfret off the model stage
func (firstfret *FirstFret) UnstageVoid(stage *Stage) {
	firstfret.Unstage(stage)
}

func (firstfret *FirstFret) StageVoid(stage *Stage) {
	firstfret.Stage(stage)
}

// for satisfaction of GongStruct interface
func (firstfret *FirstFret) GetName() (res string) {
	return firstfret.Name
}

// for satisfaction of GongStruct interface
func (firstfret *FirstFret) SetName(name string) {
	firstfret.Name = name
}

// Stage puts frame to the model stage
func (frame *Frame) Stage(stage *Stage) *Frame {
	__gong__stage(stage.Frames, stage.Frame_stagedOrder, stage.Frame_orderStaged, &stage.FrameOrder, stage.Frames_mapString, frame, frame.Name)
	return frame
}

// StagePreserveOrder puts frame to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FrameOrder
// - update stage.FrameOrder accordingly
func (frame *Frame) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Frames, stage.Frame_stagedOrder, stage.Frame_orderStaged, &stage.FrameOrder, stage.Frames_mapString, frame, order, frame.Name)
}

// Unstage removes frame off the model stage
func (frame *Frame) Unstage(stage *Stage) *Frame {
	__gong__unstage(stage.Frames, stage.Frames_mapString, frame, frame.Name)
	return frame
}

// UnstageVoid removes frame off the model stage
func (frame *Frame) UnstageVoid(stage *Stage) {
	frame.Unstage(stage)
}

func (frame *Frame) StageVoid(stage *Stage) {
	frame.Stage(stage)
}

// for satisfaction of GongStruct interface
func (frame *Frame) GetName() (res string) {
	return frame.Name
}

// for satisfaction of GongStruct interface
func (frame *Frame) SetName(name string) {
	frame.Name = name
}

// Stage puts framenote to the model stage
func (framenote *FrameNote) Stage(stage *Stage) *FrameNote {
	__gong__stage(stage.FrameNotes, stage.FrameNote_stagedOrder, stage.FrameNote_orderStaged, &stage.FrameNoteOrder, stage.FrameNotes_mapString, framenote, framenote.Name)
	return framenote
}

// StagePreserveOrder puts framenote to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FrameNoteOrder
// - update stage.FrameNoteOrder accordingly
func (framenote *FrameNote) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FrameNotes, stage.FrameNote_stagedOrder, stage.FrameNote_orderStaged, &stage.FrameNoteOrder, stage.FrameNotes_mapString, framenote, order, framenote.Name)
}

// Unstage removes framenote off the model stage
func (framenote *FrameNote) Unstage(stage *Stage) *FrameNote {
	__gong__unstage(stage.FrameNotes, stage.FrameNotes_mapString, framenote, framenote.Name)
	return framenote
}

// UnstageVoid removes framenote off the model stage
func (framenote *FrameNote) UnstageVoid(stage *Stage) {
	framenote.Unstage(stage)
}

func (framenote *FrameNote) StageVoid(stage *Stage) {
	framenote.Stage(stage)
}

// for satisfaction of GongStruct interface
func (framenote *FrameNote) GetName() (res string) {
	return framenote.Name
}

// for satisfaction of GongStruct interface
func (framenote *FrameNote) SetName(name string) {
	framenote.Name = name
}

// Stage puts glass to the model stage
func (glass *Glass) Stage(stage *Stage) *Glass {
	__gong__stage(stage.Glasss, stage.Glass_stagedOrder, stage.Glass_orderStaged, &stage.GlassOrder, stage.Glasss_mapString, glass, glass.Name)
	return glass
}

// StagePreserveOrder puts glass to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GlassOrder
// - update stage.GlassOrder accordingly
func (glass *Glass) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Glasss, stage.Glass_stagedOrder, stage.Glass_orderStaged, &stage.GlassOrder, stage.Glasss_mapString, glass, order, glass.Name)
}

// Unstage removes glass off the model stage
func (glass *Glass) Unstage(stage *Stage) *Glass {
	__gong__unstage(stage.Glasss, stage.Glasss_mapString, glass, glass.Name)
	return glass
}

// UnstageVoid removes glass off the model stage
func (glass *Glass) UnstageVoid(stage *Stage) {
	glass.Unstage(stage)
}

func (glass *Glass) StageVoid(stage *Stage) {
	glass.Stage(stage)
}

// for satisfaction of GongStruct interface
func (glass *Glass) GetName() (res string) {
	return glass.Name
}

// for satisfaction of GongStruct interface
func (glass *Glass) SetName(name string) {
	glass.Name = name
}

// Stage puts grouping to the model stage
func (grouping *Grouping) Stage(stage *Stage) *Grouping {
	__gong__stage(stage.Groupings, stage.Grouping_stagedOrder, stage.Grouping_orderStaged, &stage.GroupingOrder, stage.Groupings_mapString, grouping, grouping.Name)
	return grouping
}

// StagePreserveOrder puts grouping to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupingOrder
// - update stage.GroupingOrder accordingly
func (grouping *Grouping) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Groupings, stage.Grouping_stagedOrder, stage.Grouping_orderStaged, &stage.GroupingOrder, stage.Groupings_mapString, grouping, order, grouping.Name)
}

// Unstage removes grouping off the model stage
func (grouping *Grouping) Unstage(stage *Stage) *Grouping {
	__gong__unstage(stage.Groupings, stage.Groupings_mapString, grouping, grouping.Name)
	return grouping
}

// UnstageVoid removes grouping off the model stage
func (grouping *Grouping) UnstageVoid(stage *Stage) {
	grouping.Unstage(stage)
}

func (grouping *Grouping) StageVoid(stage *Stage) {
	grouping.Stage(stage)
}

// for satisfaction of GongStruct interface
func (grouping *Grouping) GetName() (res string) {
	return grouping.Name
}

// for satisfaction of GongStruct interface
func (grouping *Grouping) SetName(name string) {
	grouping.Name = name
}

// Stage puts harmony to the model stage
func (harmony *Harmony) Stage(stage *Stage) *Harmony {
	__gong__stage(stage.Harmonys, stage.Harmony_stagedOrder, stage.Harmony_orderStaged, &stage.HarmonyOrder, stage.Harmonys_mapString, harmony, harmony.Name)
	return harmony
}

// StagePreserveOrder puts harmony to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HarmonyOrder
// - update stage.HarmonyOrder accordingly
func (harmony *Harmony) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Harmonys, stage.Harmony_stagedOrder, stage.Harmony_orderStaged, &stage.HarmonyOrder, stage.Harmonys_mapString, harmony, order, harmony.Name)
}

// Unstage removes harmony off the model stage
func (harmony *Harmony) Unstage(stage *Stage) *Harmony {
	__gong__unstage(stage.Harmonys, stage.Harmonys_mapString, harmony, harmony.Name)
	return harmony
}

// UnstageVoid removes harmony off the model stage
func (harmony *Harmony) UnstageVoid(stage *Stage) {
	harmony.Unstage(stage)
}

func (harmony *Harmony) StageVoid(stage *Stage) {
	harmony.Stage(stage)
}

// for satisfaction of GongStruct interface
func (harmony *Harmony) GetName() (res string) {
	return harmony.Name
}

// for satisfaction of GongStruct interface
func (harmony *Harmony) SetName(name string) {
	harmony.Name = name
}

// Stage puts harmonyalter to the model stage
func (harmonyalter *HarmonyAlter) Stage(stage *Stage) *HarmonyAlter {
	__gong__stage(stage.HarmonyAlters, stage.HarmonyAlter_stagedOrder, stage.HarmonyAlter_orderStaged, &stage.HarmonyAlterOrder, stage.HarmonyAlters_mapString, harmonyalter, harmonyalter.Name)
	return harmonyalter
}

// StagePreserveOrder puts harmonyalter to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HarmonyAlterOrder
// - update stage.HarmonyAlterOrder accordingly
func (harmonyalter *HarmonyAlter) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.HarmonyAlters, stage.HarmonyAlter_stagedOrder, stage.HarmonyAlter_orderStaged, &stage.HarmonyAlterOrder, stage.HarmonyAlters_mapString, harmonyalter, order, harmonyalter.Name)
}

// Unstage removes harmonyalter off the model stage
func (harmonyalter *HarmonyAlter) Unstage(stage *Stage) *HarmonyAlter {
	__gong__unstage(stage.HarmonyAlters, stage.HarmonyAlters_mapString, harmonyalter, harmonyalter.Name)
	return harmonyalter
}

// UnstageVoid removes harmonyalter off the model stage
func (harmonyalter *HarmonyAlter) UnstageVoid(stage *Stage) {
	harmonyalter.Unstage(stage)
}

func (harmonyalter *HarmonyAlter) StageVoid(stage *Stage) {
	harmonyalter.Stage(stage)
}

// for satisfaction of GongStruct interface
func (harmonyalter *HarmonyAlter) GetName() (res string) {
	return harmonyalter.Name
}

// for satisfaction of GongStruct interface
func (harmonyalter *HarmonyAlter) SetName(name string) {
	harmonyalter.Name = name
}

// Stage puts harppedals to the model stage
func (harppedals *HarpPedals) Stage(stage *Stage) *HarpPedals {
	__gong__stage(stage.HarpPedalss, stage.HarpPedals_stagedOrder, stage.HarpPedals_orderStaged, &stage.HarpPedalsOrder, stage.HarpPedalss_mapString, harppedals, harppedals.Name)
	return harppedals
}

// StagePreserveOrder puts harppedals to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HarpPedalsOrder
// - update stage.HarpPedalsOrder accordingly
func (harppedals *HarpPedals) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.HarpPedalss, stage.HarpPedals_stagedOrder, stage.HarpPedals_orderStaged, &stage.HarpPedalsOrder, stage.HarpPedalss_mapString, harppedals, order, harppedals.Name)
}

// Unstage removes harppedals off the model stage
func (harppedals *HarpPedals) Unstage(stage *Stage) *HarpPedals {
	__gong__unstage(stage.HarpPedalss, stage.HarpPedalss_mapString, harppedals, harppedals.Name)
	return harppedals
}

// UnstageVoid removes harppedals off the model stage
func (harppedals *HarpPedals) UnstageVoid(stage *Stage) {
	harppedals.Unstage(stage)
}

func (harppedals *HarpPedals) StageVoid(stage *Stage) {
	harppedals.Stage(stage)
}

// for satisfaction of GongStruct interface
func (harppedals *HarpPedals) GetName() (res string) {
	return harppedals.Name
}

// for satisfaction of GongStruct interface
func (harppedals *HarpPedals) SetName(name string) {
	harppedals.Name = name
}

// Stage puts image to the model stage
func (image *Image) Stage(stage *Stage) *Image {
	__gong__stage(stage.Images, stage.Image_stagedOrder, stage.Image_orderStaged, &stage.ImageOrder, stage.Images_mapString, image, image.Name)
	return image
}

// StagePreserveOrder puts image to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ImageOrder
// - update stage.ImageOrder accordingly
func (image *Image) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Images, stage.Image_stagedOrder, stage.Image_orderStaged, &stage.ImageOrder, stage.Images_mapString, image, order, image.Name)
}

// Unstage removes image off the model stage
func (image *Image) Unstage(stage *Stage) *Image {
	__gong__unstage(stage.Images, stage.Images_mapString, image, image.Name)
	return image
}

// UnstageVoid removes image off the model stage
func (image *Image) UnstageVoid(stage *Stage) {
	image.Unstage(stage)
}

func (image *Image) StageVoid(stage *Stage) {
	image.Stage(stage)
}

// for satisfaction of GongStruct interface
func (image *Image) GetName() (res string) {
	return image.Name
}

// for satisfaction of GongStruct interface
func (image *Image) SetName(name string) {
	image.Name = name
}

// Stage puts instrumentchange to the model stage
func (instrumentchange *InstrumentChange) Stage(stage *Stage) *InstrumentChange {
	__gong__stage(stage.InstrumentChanges, stage.InstrumentChange_stagedOrder, stage.InstrumentChange_orderStaged, &stage.InstrumentChangeOrder, stage.InstrumentChanges_mapString, instrumentchange, instrumentchange.Name)
	return instrumentchange
}

// StagePreserveOrder puts instrumentchange to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InstrumentChangeOrder
// - update stage.InstrumentChangeOrder accordingly
func (instrumentchange *InstrumentChange) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.InstrumentChanges, stage.InstrumentChange_stagedOrder, stage.InstrumentChange_orderStaged, &stage.InstrumentChangeOrder, stage.InstrumentChanges_mapString, instrumentchange, order, instrumentchange.Name)
}

// Unstage removes instrumentchange off the model stage
func (instrumentchange *InstrumentChange) Unstage(stage *Stage) *InstrumentChange {
	__gong__unstage(stage.InstrumentChanges, stage.InstrumentChanges_mapString, instrumentchange, instrumentchange.Name)
	return instrumentchange
}

// UnstageVoid removes instrumentchange off the model stage
func (instrumentchange *InstrumentChange) UnstageVoid(stage *Stage) {
	instrumentchange.Unstage(stage)
}

func (instrumentchange *InstrumentChange) StageVoid(stage *Stage) {
	instrumentchange.Stage(stage)
}

// for satisfaction of GongStruct interface
func (instrumentchange *InstrumentChange) GetName() (res string) {
	return instrumentchange.Name
}

// for satisfaction of GongStruct interface
func (instrumentchange *InstrumentChange) SetName(name string) {
	instrumentchange.Name = name
}

// Stage puts inversion to the model stage
func (inversion *Inversion) Stage(stage *Stage) *Inversion {
	__gong__stage(stage.Inversions, stage.Inversion_stagedOrder, stage.Inversion_orderStaged, &stage.InversionOrder, stage.Inversions_mapString, inversion, inversion.Name)
	return inversion
}

// StagePreserveOrder puts inversion to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InversionOrder
// - update stage.InversionOrder accordingly
func (inversion *Inversion) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Inversions, stage.Inversion_stagedOrder, stage.Inversion_orderStaged, &stage.InversionOrder, stage.Inversions_mapString, inversion, order, inversion.Name)
}

// Unstage removes inversion off the model stage
func (inversion *Inversion) Unstage(stage *Stage) *Inversion {
	__gong__unstage(stage.Inversions, stage.Inversions_mapString, inversion, inversion.Name)
	return inversion
}

// UnstageVoid removes inversion off the model stage
func (inversion *Inversion) UnstageVoid(stage *Stage) {
	inversion.Unstage(stage)
}

func (inversion *Inversion) StageVoid(stage *Stage) {
	inversion.Stage(stage)
}

// for satisfaction of GongStruct interface
func (inversion *Inversion) GetName() (res string) {
	return inversion.Name
}

// for satisfaction of GongStruct interface
func (inversion *Inversion) SetName(name string) {
	inversion.Name = name
}

// Stage puts kind to the model stage
func (kind *Kind) Stage(stage *Stage) *Kind {
	__gong__stage(stage.Kinds, stage.Kind_stagedOrder, stage.Kind_orderStaged, &stage.KindOrder, stage.Kinds_mapString, kind, kind.Name)
	return kind
}

// StagePreserveOrder puts kind to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.KindOrder
// - update stage.KindOrder accordingly
func (kind *Kind) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Kinds, stage.Kind_stagedOrder, stage.Kind_orderStaged, &stage.KindOrder, stage.Kinds_mapString, kind, order, kind.Name)
}

// Unstage removes kind off the model stage
func (kind *Kind) Unstage(stage *Stage) *Kind {
	__gong__unstage(stage.Kinds, stage.Kinds_mapString, kind, kind.Name)
	return kind
}

// UnstageVoid removes kind off the model stage
func (kind *Kind) UnstageVoid(stage *Stage) {
	kind.Unstage(stage)
}

func (kind *Kind) StageVoid(stage *Stage) {
	kind.Stage(stage)
}

// for satisfaction of GongStruct interface
func (kind *Kind) GetName() (res string) {
	return kind.Name
}

// for satisfaction of GongStruct interface
func (kind *Kind) SetName(name string) {
	kind.Name = name
}

// Stage puts listening to the model stage
func (listening *Listening) Stage(stage *Stage) *Listening {
	__gong__stage(stage.Listenings, stage.Listening_stagedOrder, stage.Listening_orderStaged, &stage.ListeningOrder, stage.Listenings_mapString, listening, listening.Name)
	return listening
}

// StagePreserveOrder puts listening to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ListeningOrder
// - update stage.ListeningOrder accordingly
func (listening *Listening) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Listenings, stage.Listening_stagedOrder, stage.Listening_orderStaged, &stage.ListeningOrder, stage.Listenings_mapString, listening, order, listening.Name)
}

// Unstage removes listening off the model stage
func (listening *Listening) Unstage(stage *Stage) *Listening {
	__gong__unstage(stage.Listenings, stage.Listenings_mapString, listening, listening.Name)
	return listening
}

// UnstageVoid removes listening off the model stage
func (listening *Listening) UnstageVoid(stage *Stage) {
	listening.Unstage(stage)
}

func (listening *Listening) StageVoid(stage *Stage) {
	listening.Stage(stage)
}

// for satisfaction of GongStruct interface
func (listening *Listening) GetName() (res string) {
	return listening.Name
}

// for satisfaction of GongStruct interface
func (listening *Listening) SetName(name string) {
	listening.Name = name
}

// Stage puts measurenumbering to the model stage
func (measurenumbering *MeasureNumbering) Stage(stage *Stage) *MeasureNumbering {
	__gong__stage(stage.MeasureNumberings, stage.MeasureNumbering_stagedOrder, stage.MeasureNumbering_orderStaged, &stage.MeasureNumberingOrder, stage.MeasureNumberings_mapString, measurenumbering, measurenumbering.Name)
	return measurenumbering
}

// StagePreserveOrder puts measurenumbering to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeasureNumberingOrder
// - update stage.MeasureNumberingOrder accordingly
func (measurenumbering *MeasureNumbering) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeasureNumberings, stage.MeasureNumbering_stagedOrder, stage.MeasureNumbering_orderStaged, &stage.MeasureNumberingOrder, stage.MeasureNumberings_mapString, measurenumbering, order, measurenumbering.Name)
}

// Unstage removes measurenumbering off the model stage
func (measurenumbering *MeasureNumbering) Unstage(stage *Stage) *MeasureNumbering {
	__gong__unstage(stage.MeasureNumberings, stage.MeasureNumberings_mapString, measurenumbering, measurenumbering.Name)
	return measurenumbering
}

// UnstageVoid removes measurenumbering off the model stage
func (measurenumbering *MeasureNumbering) UnstageVoid(stage *Stage) {
	measurenumbering.Unstage(stage)
}

func (measurenumbering *MeasureNumbering) StageVoid(stage *Stage) {
	measurenumbering.Stage(stage)
}

// for satisfaction of GongStruct interface
func (measurenumbering *MeasureNumbering) GetName() (res string) {
	return measurenumbering.Name
}

// for satisfaction of GongStruct interface
func (measurenumbering *MeasureNumbering) SetName(name string) {
	measurenumbering.Name = name
}

// Stage puts membrane to the model stage
func (membrane *Membrane) Stage(stage *Stage) *Membrane {
	__gong__stage(stage.Membranes, stage.Membrane_stagedOrder, stage.Membrane_orderStaged, &stage.MembraneOrder, stage.Membranes_mapString, membrane, membrane.Name)
	return membrane
}

// StagePreserveOrder puts membrane to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MembraneOrder
// - update stage.MembraneOrder accordingly
func (membrane *Membrane) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Membranes, stage.Membrane_stagedOrder, stage.Membrane_orderStaged, &stage.MembraneOrder, stage.Membranes_mapString, membrane, order, membrane.Name)
}

// Unstage removes membrane off the model stage
func (membrane *Membrane) Unstage(stage *Stage) *Membrane {
	__gong__unstage(stage.Membranes, stage.Membranes_mapString, membrane, membrane.Name)
	return membrane
}

// UnstageVoid removes membrane off the model stage
func (membrane *Membrane) UnstageVoid(stage *Stage) {
	membrane.Unstage(stage)
}

func (membrane *Membrane) StageVoid(stage *Stage) {
	membrane.Stage(stage)
}

// for satisfaction of GongStruct interface
func (membrane *Membrane) GetName() (res string) {
	return membrane.Name
}

// for satisfaction of GongStruct interface
func (membrane *Membrane) SetName(name string) {
	membrane.Name = name
}

// Stage puts metal to the model stage
func (metal *Metal) Stage(stage *Stage) *Metal {
	__gong__stage(stage.Metals, stage.Metal_stagedOrder, stage.Metal_orderStaged, &stage.MetalOrder, stage.Metals_mapString, metal, metal.Name)
	return metal
}

// StagePreserveOrder puts metal to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MetalOrder
// - update stage.MetalOrder accordingly
func (metal *Metal) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Metals, stage.Metal_stagedOrder, stage.Metal_orderStaged, &stage.MetalOrder, stage.Metals_mapString, metal, order, metal.Name)
}

// Unstage removes metal off the model stage
func (metal *Metal) Unstage(stage *Stage) *Metal {
	__gong__unstage(stage.Metals, stage.Metals_mapString, metal, metal.Name)
	return metal
}

// UnstageVoid removes metal off the model stage
func (metal *Metal) UnstageVoid(stage *Stage) {
	metal.Unstage(stage)
}

func (metal *Metal) StageVoid(stage *Stage) {
	metal.Stage(stage)
}

// for satisfaction of GongStruct interface
func (metal *Metal) GetName() (res string) {
	return metal.Name
}

// for satisfaction of GongStruct interface
func (metal *Metal) SetName(name string) {
	metal.Name = name
}

// Stage puts metronome to the model stage
func (metronome *Metronome) Stage(stage *Stage) *Metronome {
	__gong__stage(stage.Metronomes, stage.Metronome_stagedOrder, stage.Metronome_orderStaged, &stage.MetronomeOrder, stage.Metronomes_mapString, metronome, metronome.Name)
	return metronome
}

// StagePreserveOrder puts metronome to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MetronomeOrder
// - update stage.MetronomeOrder accordingly
func (metronome *Metronome) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Metronomes, stage.Metronome_stagedOrder, stage.Metronome_orderStaged, &stage.MetronomeOrder, stage.Metronomes_mapString, metronome, order, metronome.Name)
}

// Unstage removes metronome off the model stage
func (metronome *Metronome) Unstage(stage *Stage) *Metronome {
	__gong__unstage(stage.Metronomes, stage.Metronomes_mapString, metronome, metronome.Name)
	return metronome
}

// UnstageVoid removes metronome off the model stage
func (metronome *Metronome) UnstageVoid(stage *Stage) {
	metronome.Unstage(stage)
}

func (metronome *Metronome) StageVoid(stage *Stage) {
	metronome.Stage(stage)
}

// for satisfaction of GongStruct interface
func (metronome *Metronome) GetName() (res string) {
	return metronome.Name
}

// for satisfaction of GongStruct interface
func (metronome *Metronome) SetName(name string) {
	metronome.Name = name
}

// Stage puts metronomebeam to the model stage
func (metronomebeam *MetronomeBeam) Stage(stage *Stage) *MetronomeBeam {
	__gong__stage(stage.MetronomeBeams, stage.MetronomeBeam_stagedOrder, stage.MetronomeBeam_orderStaged, &stage.MetronomeBeamOrder, stage.MetronomeBeams_mapString, metronomebeam, metronomebeam.Name)
	return metronomebeam
}

// StagePreserveOrder puts metronomebeam to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MetronomeBeamOrder
// - update stage.MetronomeBeamOrder accordingly
func (metronomebeam *MetronomeBeam) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MetronomeBeams, stage.MetronomeBeam_stagedOrder, stage.MetronomeBeam_orderStaged, &stage.MetronomeBeamOrder, stage.MetronomeBeams_mapString, metronomebeam, order, metronomebeam.Name)
}

// Unstage removes metronomebeam off the model stage
func (metronomebeam *MetronomeBeam) Unstage(stage *Stage) *MetronomeBeam {
	__gong__unstage(stage.MetronomeBeams, stage.MetronomeBeams_mapString, metronomebeam, metronomebeam.Name)
	return metronomebeam
}

// UnstageVoid removes metronomebeam off the model stage
func (metronomebeam *MetronomeBeam) UnstageVoid(stage *Stage) {
	metronomebeam.Unstage(stage)
}

func (metronomebeam *MetronomeBeam) StageVoid(stage *Stage) {
	metronomebeam.Stage(stage)
}

// for satisfaction of GongStruct interface
func (metronomebeam *MetronomeBeam) GetName() (res string) {
	return metronomebeam.Name
}

// for satisfaction of GongStruct interface
func (metronomebeam *MetronomeBeam) SetName(name string) {
	metronomebeam.Name = name
}

// Stage puts metronomenote to the model stage
func (metronomenote *MetronomeNote) Stage(stage *Stage) *MetronomeNote {
	__gong__stage(stage.MetronomeNotes, stage.MetronomeNote_stagedOrder, stage.MetronomeNote_orderStaged, &stage.MetronomeNoteOrder, stage.MetronomeNotes_mapString, metronomenote, metronomenote.Name)
	return metronomenote
}

// StagePreserveOrder puts metronomenote to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MetronomeNoteOrder
// - update stage.MetronomeNoteOrder accordingly
func (metronomenote *MetronomeNote) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MetronomeNotes, stage.MetronomeNote_stagedOrder, stage.MetronomeNote_orderStaged, &stage.MetronomeNoteOrder, stage.MetronomeNotes_mapString, metronomenote, order, metronomenote.Name)
}

// Unstage removes metronomenote off the model stage
func (metronomenote *MetronomeNote) Unstage(stage *Stage) *MetronomeNote {
	__gong__unstage(stage.MetronomeNotes, stage.MetronomeNotes_mapString, metronomenote, metronomenote.Name)
	return metronomenote
}

// UnstageVoid removes metronomenote off the model stage
func (metronomenote *MetronomeNote) UnstageVoid(stage *Stage) {
	metronomenote.Unstage(stage)
}

func (metronomenote *MetronomeNote) StageVoid(stage *Stage) {
	metronomenote.Stage(stage)
}

// for satisfaction of GongStruct interface
func (metronomenote *MetronomeNote) GetName() (res string) {
	return metronomenote.Name
}

// for satisfaction of GongStruct interface
func (metronomenote *MetronomeNote) SetName(name string) {
	metronomenote.Name = name
}

// Stage puts metronometied to the model stage
func (metronometied *MetronomeTied) Stage(stage *Stage) *MetronomeTied {
	__gong__stage(stage.MetronomeTieds, stage.MetronomeTied_stagedOrder, stage.MetronomeTied_orderStaged, &stage.MetronomeTiedOrder, stage.MetronomeTieds_mapString, metronometied, metronometied.Name)
	return metronometied
}

// StagePreserveOrder puts metronometied to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MetronomeTiedOrder
// - update stage.MetronomeTiedOrder accordingly
func (metronometied *MetronomeTied) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MetronomeTieds, stage.MetronomeTied_stagedOrder, stage.MetronomeTied_orderStaged, &stage.MetronomeTiedOrder, stage.MetronomeTieds_mapString, metronometied, order, metronometied.Name)
}

// Unstage removes metronometied off the model stage
func (metronometied *MetronomeTied) Unstage(stage *Stage) *MetronomeTied {
	__gong__unstage(stage.MetronomeTieds, stage.MetronomeTieds_mapString, metronometied, metronometied.Name)
	return metronometied
}

// UnstageVoid removes metronometied off the model stage
func (metronometied *MetronomeTied) UnstageVoid(stage *Stage) {
	metronometied.Unstage(stage)
}

func (metronometied *MetronomeTied) StageVoid(stage *Stage) {
	metronometied.Stage(stage)
}

// for satisfaction of GongStruct interface
func (metronometied *MetronomeTied) GetName() (res string) {
	return metronometied.Name
}

// for satisfaction of GongStruct interface
func (metronometied *MetronomeTied) SetName(name string) {
	metronometied.Name = name
}

// Stage puts metronometuplet to the model stage
func (metronometuplet *MetronomeTuplet) Stage(stage *Stage) *MetronomeTuplet {
	__gong__stage(stage.MetronomeTuplets, stage.MetronomeTuplet_stagedOrder, stage.MetronomeTuplet_orderStaged, &stage.MetronomeTupletOrder, stage.MetronomeTuplets_mapString, metronometuplet, metronometuplet.Name)
	return metronometuplet
}

// StagePreserveOrder puts metronometuplet to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MetronomeTupletOrder
// - update stage.MetronomeTupletOrder accordingly
func (metronometuplet *MetronomeTuplet) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MetronomeTuplets, stage.MetronomeTuplet_stagedOrder, stage.MetronomeTuplet_orderStaged, &stage.MetronomeTupletOrder, stage.MetronomeTuplets_mapString, metronometuplet, order, metronometuplet.Name)
}

// Unstage removes metronometuplet off the model stage
func (metronometuplet *MetronomeTuplet) Unstage(stage *Stage) *MetronomeTuplet {
	__gong__unstage(stage.MetronomeTuplets, stage.MetronomeTuplets_mapString, metronometuplet, metronometuplet.Name)
	return metronometuplet
}

// UnstageVoid removes metronometuplet off the model stage
func (metronometuplet *MetronomeTuplet) UnstageVoid(stage *Stage) {
	metronometuplet.Unstage(stage)
}

func (metronometuplet *MetronomeTuplet) StageVoid(stage *Stage) {
	metronometuplet.Stage(stage)
}

// for satisfaction of GongStruct interface
func (metronometuplet *MetronomeTuplet) GetName() (res string) {
	return metronometuplet.Name
}

// for satisfaction of GongStruct interface
func (metronometuplet *MetronomeTuplet) SetName(name string) {
	metronometuplet.Name = name
}

// Stage puts numeral to the model stage
func (numeral *Numeral) Stage(stage *Stage) *Numeral {
	__gong__stage(stage.Numerals, stage.Numeral_stagedOrder, stage.Numeral_orderStaged, &stage.NumeralOrder, stage.Numerals_mapString, numeral, numeral.Name)
	return numeral
}

// StagePreserveOrder puts numeral to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NumeralOrder
// - update stage.NumeralOrder accordingly
func (numeral *Numeral) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Numerals, stage.Numeral_stagedOrder, stage.Numeral_orderStaged, &stage.NumeralOrder, stage.Numerals_mapString, numeral, order, numeral.Name)
}

// Unstage removes numeral off the model stage
func (numeral *Numeral) Unstage(stage *Stage) *Numeral {
	__gong__unstage(stage.Numerals, stage.Numerals_mapString, numeral, numeral.Name)
	return numeral
}

// UnstageVoid removes numeral off the model stage
func (numeral *Numeral) UnstageVoid(stage *Stage) {
	numeral.Unstage(stage)
}

func (numeral *Numeral) StageVoid(stage *Stage) {
	numeral.Stage(stage)
}

// for satisfaction of GongStruct interface
func (numeral *Numeral) GetName() (res string) {
	return numeral.Name
}

// for satisfaction of GongStruct interface
func (numeral *Numeral) SetName(name string) {
	numeral.Name = name
}

// Stage puts numeralkey to the model stage
func (numeralkey *NumeralKey) Stage(stage *Stage) *NumeralKey {
	__gong__stage(stage.NumeralKeys, stage.NumeralKey_stagedOrder, stage.NumeralKey_orderStaged, &stage.NumeralKeyOrder, stage.NumeralKeys_mapString, numeralkey, numeralkey.Name)
	return numeralkey
}

// StagePreserveOrder puts numeralkey to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NumeralKeyOrder
// - update stage.NumeralKeyOrder accordingly
func (numeralkey *NumeralKey) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NumeralKeys, stage.NumeralKey_stagedOrder, stage.NumeralKey_orderStaged, &stage.NumeralKeyOrder, stage.NumeralKeys_mapString, numeralkey, order, numeralkey.Name)
}

// Unstage removes numeralkey off the model stage
func (numeralkey *NumeralKey) Unstage(stage *Stage) *NumeralKey {
	__gong__unstage(stage.NumeralKeys, stage.NumeralKeys_mapString, numeralkey, numeralkey.Name)
	return numeralkey
}

// UnstageVoid removes numeralkey off the model stage
func (numeralkey *NumeralKey) UnstageVoid(stage *Stage) {
	numeralkey.Unstage(stage)
}

func (numeralkey *NumeralKey) StageVoid(stage *Stage) {
	numeralkey.Stage(stage)
}

// for satisfaction of GongStruct interface
func (numeralkey *NumeralKey) GetName() (res string) {
	return numeralkey.Name
}

// for satisfaction of GongStruct interface
func (numeralkey *NumeralKey) SetName(name string) {
	numeralkey.Name = name
}

// Stage puts numeralroot to the model stage
func (numeralroot *NumeralRoot) Stage(stage *Stage) *NumeralRoot {
	__gong__stage(stage.NumeralRoots, stage.NumeralRoot_stagedOrder, stage.NumeralRoot_orderStaged, &stage.NumeralRootOrder, stage.NumeralRoots_mapString, numeralroot, numeralroot.Name)
	return numeralroot
}

// StagePreserveOrder puts numeralroot to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NumeralRootOrder
// - update stage.NumeralRootOrder accordingly
func (numeralroot *NumeralRoot) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NumeralRoots, stage.NumeralRoot_stagedOrder, stage.NumeralRoot_orderStaged, &stage.NumeralRootOrder, stage.NumeralRoots_mapString, numeralroot, order, numeralroot.Name)
}

// Unstage removes numeralroot off the model stage
func (numeralroot *NumeralRoot) Unstage(stage *Stage) *NumeralRoot {
	__gong__unstage(stage.NumeralRoots, stage.NumeralRoots_mapString, numeralroot, numeralroot.Name)
	return numeralroot
}

// UnstageVoid removes numeralroot off the model stage
func (numeralroot *NumeralRoot) UnstageVoid(stage *Stage) {
	numeralroot.Unstage(stage)
}

func (numeralroot *NumeralRoot) StageVoid(stage *Stage) {
	numeralroot.Stage(stage)
}

// for satisfaction of GongStruct interface
func (numeralroot *NumeralRoot) GetName() (res string) {
	return numeralroot.Name
}

// for satisfaction of GongStruct interface
func (numeralroot *NumeralRoot) SetName(name string) {
	numeralroot.Name = name
}

// Stage puts octaveshift to the model stage
func (octaveshift *OctaveShift) Stage(stage *Stage) *OctaveShift {
	__gong__stage(stage.OctaveShifts, stage.OctaveShift_stagedOrder, stage.OctaveShift_orderStaged, &stage.OctaveShiftOrder, stage.OctaveShifts_mapString, octaveshift, octaveshift.Name)
	return octaveshift
}

// StagePreserveOrder puts octaveshift to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OctaveShiftOrder
// - update stage.OctaveShiftOrder accordingly
func (octaveshift *OctaveShift) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OctaveShifts, stage.OctaveShift_stagedOrder, stage.OctaveShift_orderStaged, &stage.OctaveShiftOrder, stage.OctaveShifts_mapString, octaveshift, order, octaveshift.Name)
}

// Unstage removes octaveshift off the model stage
func (octaveshift *OctaveShift) Unstage(stage *Stage) *OctaveShift {
	__gong__unstage(stage.OctaveShifts, stage.OctaveShifts_mapString, octaveshift, octaveshift.Name)
	return octaveshift
}

// UnstageVoid removes octaveshift off the model stage
func (octaveshift *OctaveShift) UnstageVoid(stage *Stage) {
	octaveshift.Unstage(stage)
}

func (octaveshift *OctaveShift) StageVoid(stage *Stage) {
	octaveshift.Stage(stage)
}

// for satisfaction of GongStruct interface
func (octaveshift *OctaveShift) GetName() (res string) {
	return octaveshift.Name
}

// for satisfaction of GongStruct interface
func (octaveshift *OctaveShift) SetName(name string) {
	octaveshift.Name = name
}

// Stage puts offset to the model stage
func (offset *Offset) Stage(stage *Stage) *Offset {
	__gong__stage(stage.Offsets, stage.Offset_stagedOrder, stage.Offset_orderStaged, &stage.OffsetOrder, stage.Offsets_mapString, offset, offset.Name)
	return offset
}

// StagePreserveOrder puts offset to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OffsetOrder
// - update stage.OffsetOrder accordingly
func (offset *Offset) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Offsets, stage.Offset_stagedOrder, stage.Offset_orderStaged, &stage.OffsetOrder, stage.Offsets_mapString, offset, order, offset.Name)
}

// Unstage removes offset off the model stage
func (offset *Offset) Unstage(stage *Stage) *Offset {
	__gong__unstage(stage.Offsets, stage.Offsets_mapString, offset, offset.Name)
	return offset
}

// UnstageVoid removes offset off the model stage
func (offset *Offset) UnstageVoid(stage *Stage) {
	offset.Unstage(stage)
}

func (offset *Offset) StageVoid(stage *Stage) {
	offset.Stage(stage)
}

// for satisfaction of GongStruct interface
func (offset *Offset) GetName() (res string) {
	return offset.Name
}

// for satisfaction of GongStruct interface
func (offset *Offset) SetName(name string) {
	offset.Name = name
}

// Stage puts otherdirection to the model stage
func (otherdirection *OtherDirection) Stage(stage *Stage) *OtherDirection {
	__gong__stage(stage.OtherDirections, stage.OtherDirection_stagedOrder, stage.OtherDirection_orderStaged, &stage.OtherDirectionOrder, stage.OtherDirections_mapString, otherdirection, otherdirection.Name)
	return otherdirection
}

// StagePreserveOrder puts otherdirection to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OtherDirectionOrder
// - update stage.OtherDirectionOrder accordingly
func (otherdirection *OtherDirection) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OtherDirections, stage.OtherDirection_stagedOrder, stage.OtherDirection_orderStaged, &stage.OtherDirectionOrder, stage.OtherDirections_mapString, otherdirection, order, otherdirection.Name)
}

// Unstage removes otherdirection off the model stage
func (otherdirection *OtherDirection) Unstage(stage *Stage) *OtherDirection {
	__gong__unstage(stage.OtherDirections, stage.OtherDirections_mapString, otherdirection, otherdirection.Name)
	return otherdirection
}

// UnstageVoid removes otherdirection off the model stage
func (otherdirection *OtherDirection) UnstageVoid(stage *Stage) {
	otherdirection.Unstage(stage)
}

func (otherdirection *OtherDirection) StageVoid(stage *Stage) {
	otherdirection.Stage(stage)
}

// for satisfaction of GongStruct interface
func (otherdirection *OtherDirection) GetName() (res string) {
	return otherdirection.Name
}

// for satisfaction of GongStruct interface
func (otherdirection *OtherDirection) SetName(name string) {
	otherdirection.Name = name
}

// Stage puts otherlistening to the model stage
func (otherlistening *OtherListening) Stage(stage *Stage) *OtherListening {
	__gong__stage(stage.OtherListenings, stage.OtherListening_stagedOrder, stage.OtherListening_orderStaged, &stage.OtherListeningOrder, stage.OtherListenings_mapString, otherlistening, otherlistening.Name)
	return otherlistening
}

// StagePreserveOrder puts otherlistening to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OtherListeningOrder
// - update stage.OtherListeningOrder accordingly
func (otherlistening *OtherListening) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OtherListenings, stage.OtherListening_stagedOrder, stage.OtherListening_orderStaged, &stage.OtherListeningOrder, stage.OtherListenings_mapString, otherlistening, order, otherlistening.Name)
}

// Unstage removes otherlistening off the model stage
func (otherlistening *OtherListening) Unstage(stage *Stage) *OtherListening {
	__gong__unstage(stage.OtherListenings, stage.OtherListenings_mapString, otherlistening, otherlistening.Name)
	return otherlistening
}

// UnstageVoid removes otherlistening off the model stage
func (otherlistening *OtherListening) UnstageVoid(stage *Stage) {
	otherlistening.Unstage(stage)
}

func (otherlistening *OtherListening) StageVoid(stage *Stage) {
	otherlistening.Stage(stage)
}

// for satisfaction of GongStruct interface
func (otherlistening *OtherListening) GetName() (res string) {
	return otherlistening.Name
}

// for satisfaction of GongStruct interface
func (otherlistening *OtherListening) SetName(name string) {
	otherlistening.Name = name
}

// Stage puts pedal to the model stage
func (pedal *Pedal) Stage(stage *Stage) *Pedal {
	__gong__stage(stage.Pedals, stage.Pedal_stagedOrder, stage.Pedal_orderStaged, &stage.PedalOrder, stage.Pedals_mapString, pedal, pedal.Name)
	return pedal
}

// StagePreserveOrder puts pedal to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PedalOrder
// - update stage.PedalOrder accordingly
func (pedal *Pedal) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Pedals, stage.Pedal_stagedOrder, stage.Pedal_orderStaged, &stage.PedalOrder, stage.Pedals_mapString, pedal, order, pedal.Name)
}

// Unstage removes pedal off the model stage
func (pedal *Pedal) Unstage(stage *Stage) *Pedal {
	__gong__unstage(stage.Pedals, stage.Pedals_mapString, pedal, pedal.Name)
	return pedal
}

// UnstageVoid removes pedal off the model stage
func (pedal *Pedal) UnstageVoid(stage *Stage) {
	pedal.Unstage(stage)
}

func (pedal *Pedal) StageVoid(stage *Stage) {
	pedal.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pedal *Pedal) GetName() (res string) {
	return pedal.Name
}

// for satisfaction of GongStruct interface
func (pedal *Pedal) SetName(name string) {
	pedal.Name = name
}

// Stage puts pedaltuning to the model stage
func (pedaltuning *PedalTuning) Stage(stage *Stage) *PedalTuning {
	__gong__stage(stage.PedalTunings, stage.PedalTuning_stagedOrder, stage.PedalTuning_orderStaged, &stage.PedalTuningOrder, stage.PedalTunings_mapString, pedaltuning, pedaltuning.Name)
	return pedaltuning
}

// StagePreserveOrder puts pedaltuning to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PedalTuningOrder
// - update stage.PedalTuningOrder accordingly
func (pedaltuning *PedalTuning) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PedalTunings, stage.PedalTuning_stagedOrder, stage.PedalTuning_orderStaged, &stage.PedalTuningOrder, stage.PedalTunings_mapString, pedaltuning, order, pedaltuning.Name)
}

// Unstage removes pedaltuning off the model stage
func (pedaltuning *PedalTuning) Unstage(stage *Stage) *PedalTuning {
	__gong__unstage(stage.PedalTunings, stage.PedalTunings_mapString, pedaltuning, pedaltuning.Name)
	return pedaltuning
}

// UnstageVoid removes pedaltuning off the model stage
func (pedaltuning *PedalTuning) UnstageVoid(stage *Stage) {
	pedaltuning.Unstage(stage)
}

func (pedaltuning *PedalTuning) StageVoid(stage *Stage) {
	pedaltuning.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pedaltuning *PedalTuning) GetName() (res string) {
	return pedaltuning.Name
}

// for satisfaction of GongStruct interface
func (pedaltuning *PedalTuning) SetName(name string) {
	pedaltuning.Name = name
}

// Stage puts perminute to the model stage
func (perminute *PerMinute) Stage(stage *Stage) *PerMinute {
	__gong__stage(stage.PerMinutes, stage.PerMinute_stagedOrder, stage.PerMinute_orderStaged, &stage.PerMinuteOrder, stage.PerMinutes_mapString, perminute, perminute.Name)
	return perminute
}

// StagePreserveOrder puts perminute to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PerMinuteOrder
// - update stage.PerMinuteOrder accordingly
func (perminute *PerMinute) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PerMinutes, stage.PerMinute_stagedOrder, stage.PerMinute_orderStaged, &stage.PerMinuteOrder, stage.PerMinutes_mapString, perminute, order, perminute.Name)
}

// Unstage removes perminute off the model stage
func (perminute *PerMinute) Unstage(stage *Stage) *PerMinute {
	__gong__unstage(stage.PerMinutes, stage.PerMinutes_mapString, perminute, perminute.Name)
	return perminute
}

// UnstageVoid removes perminute off the model stage
func (perminute *PerMinute) UnstageVoid(stage *Stage) {
	perminute.Unstage(stage)
}

func (perminute *PerMinute) StageVoid(stage *Stage) {
	perminute.Stage(stage)
}

// for satisfaction of GongStruct interface
func (perminute *PerMinute) GetName() (res string) {
	return perminute.Name
}

// for satisfaction of GongStruct interface
func (perminute *PerMinute) SetName(name string) {
	perminute.Name = name
}

// Stage puts percussion to the model stage
func (percussion *Percussion) Stage(stage *Stage) *Percussion {
	__gong__stage(stage.Percussions, stage.Percussion_stagedOrder, stage.Percussion_orderStaged, &stage.PercussionOrder, stage.Percussions_mapString, percussion, percussion.Name)
	return percussion
}

// StagePreserveOrder puts percussion to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PercussionOrder
// - update stage.PercussionOrder accordingly
func (percussion *Percussion) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Percussions, stage.Percussion_stagedOrder, stage.Percussion_orderStaged, &stage.PercussionOrder, stage.Percussions_mapString, percussion, order, percussion.Name)
}

// Unstage removes percussion off the model stage
func (percussion *Percussion) Unstage(stage *Stage) *Percussion {
	__gong__unstage(stage.Percussions, stage.Percussions_mapString, percussion, percussion.Name)
	return percussion
}

// UnstageVoid removes percussion off the model stage
func (percussion *Percussion) UnstageVoid(stage *Stage) {
	percussion.Unstage(stage)
}

func (percussion *Percussion) StageVoid(stage *Stage) {
	percussion.Stage(stage)
}

// for satisfaction of GongStruct interface
func (percussion *Percussion) GetName() (res string) {
	return percussion.Name
}

// for satisfaction of GongStruct interface
func (percussion *Percussion) SetName(name string) {
	percussion.Name = name
}

// Stage puts pitched to the model stage
func (pitched *Pitched) Stage(stage *Stage) *Pitched {
	__gong__stage(stage.Pitcheds, stage.Pitched_stagedOrder, stage.Pitched_orderStaged, &stage.PitchedOrder, stage.Pitcheds_mapString, pitched, pitched.Name)
	return pitched
}

// StagePreserveOrder puts pitched to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PitchedOrder
// - update stage.PitchedOrder accordingly
func (pitched *Pitched) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Pitcheds, stage.Pitched_stagedOrder, stage.Pitched_orderStaged, &stage.PitchedOrder, stage.Pitcheds_mapString, pitched, order, pitched.Name)
}

// Unstage removes pitched off the model stage
func (pitched *Pitched) Unstage(stage *Stage) *Pitched {
	__gong__unstage(stage.Pitcheds, stage.Pitcheds_mapString, pitched, pitched.Name)
	return pitched
}

// UnstageVoid removes pitched off the model stage
func (pitched *Pitched) UnstageVoid(stage *Stage) {
	pitched.Unstage(stage)
}

func (pitched *Pitched) StageVoid(stage *Stage) {
	pitched.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pitched *Pitched) GetName() (res string) {
	return pitched.Name
}

// for satisfaction of GongStruct interface
func (pitched *Pitched) SetName(name string) {
	pitched.Name = name
}

// Stage puts principalvoice to the model stage
func (principalvoice *PrincipalVoice) Stage(stage *Stage) *PrincipalVoice {
	__gong__stage(stage.PrincipalVoices, stage.PrincipalVoice_stagedOrder, stage.PrincipalVoice_orderStaged, &stage.PrincipalVoiceOrder, stage.PrincipalVoices_mapString, principalvoice, principalvoice.Name)
	return principalvoice
}

// StagePreserveOrder puts principalvoice to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PrincipalVoiceOrder
// - update stage.PrincipalVoiceOrder accordingly
func (principalvoice *PrincipalVoice) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PrincipalVoices, stage.PrincipalVoice_stagedOrder, stage.PrincipalVoice_orderStaged, &stage.PrincipalVoiceOrder, stage.PrincipalVoices_mapString, principalvoice, order, principalvoice.Name)
}

// Unstage removes principalvoice off the model stage
func (principalvoice *PrincipalVoice) Unstage(stage *Stage) *PrincipalVoice {
	__gong__unstage(stage.PrincipalVoices, stage.PrincipalVoices_mapString, principalvoice, principalvoice.Name)
	return principalvoice
}

// UnstageVoid removes principalvoice off the model stage
func (principalvoice *PrincipalVoice) UnstageVoid(stage *Stage) {
	principalvoice.Unstage(stage)
}

func (principalvoice *PrincipalVoice) StageVoid(stage *Stage) {
	principalvoice.Stage(stage)
}

// for satisfaction of GongStruct interface
func (principalvoice *PrincipalVoice) GetName() (res string) {
	return principalvoice.Name
}

// for satisfaction of GongStruct interface
func (principalvoice *PrincipalVoice) SetName(name string) {
	principalvoice.Name = name
}

// Stage puts print to the model stage
func (print *Print) Stage(stage *Stage) *Print {
	__gong__stage(stage.Prints, stage.Print_stagedOrder, stage.Print_orderStaged, &stage.PrintOrder, stage.Prints_mapString, print, print.Name)
	return print
}

// StagePreserveOrder puts print to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PrintOrder
// - update stage.PrintOrder accordingly
func (print *Print) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Prints, stage.Print_stagedOrder, stage.Print_orderStaged, &stage.PrintOrder, stage.Prints_mapString, print, order, print.Name)
}

// Unstage removes print off the model stage
func (print *Print) Unstage(stage *Stage) *Print {
	__gong__unstage(stage.Prints, stage.Prints_mapString, print, print.Name)
	return print
}

// UnstageVoid removes print off the model stage
func (print *Print) UnstageVoid(stage *Stage) {
	print.Unstage(stage)
}

func (print *Print) StageVoid(stage *Stage) {
	print.Stage(stage)
}

// for satisfaction of GongStruct interface
func (print *Print) GetName() (res string) {
	return print.Name
}

// for satisfaction of GongStruct interface
func (print *Print) SetName(name string) {
	print.Name = name
}

// Stage puts root to the model stage
func (root *Root) Stage(stage *Stage) *Root {
	__gong__stage(stage.Roots, stage.Root_stagedOrder, stage.Root_orderStaged, &stage.RootOrder, stage.Roots_mapString, root, root.Name)
	return root
}

// StagePreserveOrder puts root to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RootOrder
// - update stage.RootOrder accordingly
func (root *Root) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Roots, stage.Root_stagedOrder, stage.Root_orderStaged, &stage.RootOrder, stage.Roots_mapString, root, order, root.Name)
}

// Unstage removes root off the model stage
func (root *Root) Unstage(stage *Stage) *Root {
	__gong__unstage(stage.Roots, stage.Roots_mapString, root, root.Name)
	return root
}

// UnstageVoid removes root off the model stage
func (root *Root) UnstageVoid(stage *Stage) {
	root.Unstage(stage)
}

func (root *Root) StageVoid(stage *Stage) {
	root.Stage(stage)
}

// for satisfaction of GongStruct interface
func (root *Root) GetName() (res string) {
	return root.Name
}

// for satisfaction of GongStruct interface
func (root *Root) SetName(name string) {
	root.Name = name
}

// Stage puts rootstep to the model stage
func (rootstep *RootStep) Stage(stage *Stage) *RootStep {
	__gong__stage(stage.RootSteps, stage.RootStep_stagedOrder, stage.RootStep_orderStaged, &stage.RootStepOrder, stage.RootSteps_mapString, rootstep, rootstep.Name)
	return rootstep
}

// StagePreserveOrder puts rootstep to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RootStepOrder
// - update stage.RootStepOrder accordingly
func (rootstep *RootStep) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RootSteps, stage.RootStep_stagedOrder, stage.RootStep_orderStaged, &stage.RootStepOrder, stage.RootSteps_mapString, rootstep, order, rootstep.Name)
}

// Unstage removes rootstep off the model stage
func (rootstep *RootStep) Unstage(stage *Stage) *RootStep {
	__gong__unstage(stage.RootSteps, stage.RootSteps_mapString, rootstep, rootstep.Name)
	return rootstep
}

// UnstageVoid removes rootstep off the model stage
func (rootstep *RootStep) UnstageVoid(stage *Stage) {
	rootstep.Unstage(stage)
}

func (rootstep *RootStep) StageVoid(stage *Stage) {
	rootstep.Stage(stage)
}

// for satisfaction of GongStruct interface
func (rootstep *RootStep) GetName() (res string) {
	return rootstep.Name
}

// for satisfaction of GongStruct interface
func (rootstep *RootStep) SetName(name string) {
	rootstep.Name = name
}

// Stage puts scordatura to the model stage
func (scordatura *Scordatura) Stage(stage *Stage) *Scordatura {
	__gong__stage(stage.Scordaturas, stage.Scordatura_stagedOrder, stage.Scordatura_orderStaged, &stage.ScordaturaOrder, stage.Scordaturas_mapString, scordatura, scordatura.Name)
	return scordatura
}

// StagePreserveOrder puts scordatura to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScordaturaOrder
// - update stage.ScordaturaOrder accordingly
func (scordatura *Scordatura) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Scordaturas, stage.Scordatura_stagedOrder, stage.Scordatura_orderStaged, &stage.ScordaturaOrder, stage.Scordaturas_mapString, scordatura, order, scordatura.Name)
}

// Unstage removes scordatura off the model stage
func (scordatura *Scordatura) Unstage(stage *Stage) *Scordatura {
	__gong__unstage(stage.Scordaturas, stage.Scordaturas_mapString, scordatura, scordatura.Name)
	return scordatura
}

// UnstageVoid removes scordatura off the model stage
func (scordatura *Scordatura) UnstageVoid(stage *Stage) {
	scordatura.Unstage(stage)
}

func (scordatura *Scordatura) StageVoid(stage *Stage) {
	scordatura.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scordatura *Scordatura) GetName() (res string) {
	return scordatura.Name
}

// for satisfaction of GongStruct interface
func (scordatura *Scordatura) SetName(name string) {
	scordatura.Name = name
}

// Stage puts sound to the model stage
func (sound *Sound) Stage(stage *Stage) *Sound {
	__gong__stage(stage.Sounds, stage.Sound_stagedOrder, stage.Sound_orderStaged, &stage.SoundOrder, stage.Sounds_mapString, sound, sound.Name)
	return sound
}

// StagePreserveOrder puts sound to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SoundOrder
// - update stage.SoundOrder accordingly
func (sound *Sound) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Sounds, stage.Sound_stagedOrder, stage.Sound_orderStaged, &stage.SoundOrder, stage.Sounds_mapString, sound, order, sound.Name)
}

// Unstage removes sound off the model stage
func (sound *Sound) Unstage(stage *Stage) *Sound {
	__gong__unstage(stage.Sounds, stage.Sounds_mapString, sound, sound.Name)
	return sound
}

// UnstageVoid removes sound off the model stage
func (sound *Sound) UnstageVoid(stage *Stage) {
	sound.Unstage(stage)
}

func (sound *Sound) StageVoid(stage *Stage) {
	sound.Stage(stage)
}

// for satisfaction of GongStruct interface
func (sound *Sound) GetName() (res string) {
	return sound.Name
}

// for satisfaction of GongStruct interface
func (sound *Sound) SetName(name string) {
	sound.Name = name
}

// Stage puts staffdivide to the model stage
func (staffdivide *StaffDivide) Stage(stage *Stage) *StaffDivide {
	__gong__stage(stage.StaffDivides, stage.StaffDivide_stagedOrder, stage.StaffDivide_orderStaged, &stage.StaffDivideOrder, stage.StaffDivides_mapString, staffdivide, staffdivide.Name)
	return staffdivide
}

// StagePreserveOrder puts staffdivide to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StaffDivideOrder
// - update stage.StaffDivideOrder accordingly
func (staffdivide *StaffDivide) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StaffDivides, stage.StaffDivide_stagedOrder, stage.StaffDivide_orderStaged, &stage.StaffDivideOrder, stage.StaffDivides_mapString, staffdivide, order, staffdivide.Name)
}

// Unstage removes staffdivide off the model stage
func (staffdivide *StaffDivide) Unstage(stage *Stage) *StaffDivide {
	__gong__unstage(stage.StaffDivides, stage.StaffDivides_mapString, staffdivide, staffdivide.Name)
	return staffdivide
}

// UnstageVoid removes staffdivide off the model stage
func (staffdivide *StaffDivide) UnstageVoid(stage *Stage) {
	staffdivide.Unstage(stage)
}

func (staffdivide *StaffDivide) StageVoid(stage *Stage) {
	staffdivide.Stage(stage)
}

// for satisfaction of GongStruct interface
func (staffdivide *StaffDivide) GetName() (res string) {
	return staffdivide.Name
}

// for satisfaction of GongStruct interface
func (staffdivide *StaffDivide) SetName(name string) {
	staffdivide.Name = name
}

// Stage puts stick to the model stage
func (stick *Stick) Stage(stage *Stage) *Stick {
	__gong__stage(stage.Sticks, stage.Stick_stagedOrder, stage.Stick_orderStaged, &stage.StickOrder, stage.Sticks_mapString, stick, stick.Name)
	return stick
}

// StagePreserveOrder puts stick to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StickOrder
// - update stage.StickOrder accordingly
func (stick *Stick) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Sticks, stage.Stick_stagedOrder, stage.Stick_orderStaged, &stage.StickOrder, stage.Sticks_mapString, stick, order, stick.Name)
}

// Unstage removes stick off the model stage
func (stick *Stick) Unstage(stage *Stage) *Stick {
	__gong__unstage(stage.Sticks, stage.Sticks_mapString, stick, stick.Name)
	return stick
}

// UnstageVoid removes stick off the model stage
func (stick *Stick) UnstageVoid(stage *Stage) {
	stick.Unstage(stage)
}

func (stick *Stick) StageVoid(stage *Stage) {
	stick.Stage(stage)
}

// for satisfaction of GongStruct interface
func (stick *Stick) GetName() (res string) {
	return stick.Name
}

// for satisfaction of GongStruct interface
func (stick *Stick) SetName(name string) {
	stick.Name = name
}

// Stage puts stringmute to the model stage
func (stringmute *StringMute) Stage(stage *Stage) *StringMute {
	__gong__stage(stage.StringMutes, stage.StringMute_stagedOrder, stage.StringMute_orderStaged, &stage.StringMuteOrder, stage.StringMutes_mapString, stringmute, stringmute.Name)
	return stringmute
}

// StagePreserveOrder puts stringmute to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StringMuteOrder
// - update stage.StringMuteOrder accordingly
func (stringmute *StringMute) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StringMutes, stage.StringMute_stagedOrder, stage.StringMute_orderStaged, &stage.StringMuteOrder, stage.StringMutes_mapString, stringmute, order, stringmute.Name)
}

// Unstage removes stringmute off the model stage
func (stringmute *StringMute) Unstage(stage *Stage) *StringMute {
	__gong__unstage(stage.StringMutes, stage.StringMutes_mapString, stringmute, stringmute.Name)
	return stringmute
}

// UnstageVoid removes stringmute off the model stage
func (stringmute *StringMute) UnstageVoid(stage *Stage) {
	stringmute.Unstage(stage)
}

func (stringmute *StringMute) StageVoid(stage *Stage) {
	stringmute.Stage(stage)
}

// for satisfaction of GongStruct interface
func (stringmute *StringMute) GetName() (res string) {
	return stringmute.Name
}

// for satisfaction of GongStruct interface
func (stringmute *StringMute) SetName(name string) {
	stringmute.Name = name
}

// Stage puts styletext to the model stage
func (styletext *StyleText) Stage(stage *Stage) *StyleText {
	__gong__stage(stage.StyleTexts, stage.StyleText_stagedOrder, stage.StyleText_orderStaged, &stage.StyleTextOrder, stage.StyleTexts_mapString, styletext, styletext.Name)
	return styletext
}

// StagePreserveOrder puts styletext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StyleTextOrder
// - update stage.StyleTextOrder accordingly
func (styletext *StyleText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StyleTexts, stage.StyleText_stagedOrder, stage.StyleText_orderStaged, &stage.StyleTextOrder, stage.StyleTexts_mapString, styletext, order, styletext.Name)
}

// Unstage removes styletext off the model stage
func (styletext *StyleText) Unstage(stage *Stage) *StyleText {
	__gong__unstage(stage.StyleTexts, stage.StyleTexts_mapString, styletext, styletext.Name)
	return styletext
}

// UnstageVoid removes styletext off the model stage
func (styletext *StyleText) UnstageVoid(stage *Stage) {
	styletext.Unstage(stage)
}

func (styletext *StyleText) StageVoid(stage *Stage) {
	styletext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (styletext *StyleText) GetName() (res string) {
	return styletext.Name
}

// for satisfaction of GongStruct interface
func (styletext *StyleText) SetName(name string) {
	styletext.Name = name
}

// Stage puts swing to the model stage
func (swing *Swing) Stage(stage *Stage) *Swing {
	__gong__stage(stage.Swings, stage.Swing_stagedOrder, stage.Swing_orderStaged, &stage.SwingOrder, stage.Swings_mapString, swing, swing.Name)
	return swing
}

// StagePreserveOrder puts swing to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SwingOrder
// - update stage.SwingOrder accordingly
func (swing *Swing) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Swings, stage.Swing_stagedOrder, stage.Swing_orderStaged, &stage.SwingOrder, stage.Swings_mapString, swing, order, swing.Name)
}

// Unstage removes swing off the model stage
func (swing *Swing) Unstage(stage *Stage) *Swing {
	__gong__unstage(stage.Swings, stage.Swings_mapString, swing, swing.Name)
	return swing
}

// UnstageVoid removes swing off the model stage
func (swing *Swing) UnstageVoid(stage *Stage) {
	swing.Unstage(stage)
}

func (swing *Swing) StageVoid(stage *Stage) {
	swing.Stage(stage)
}

// for satisfaction of GongStruct interface
func (swing *Swing) GetName() (res string) {
	return swing.Name
}

// for satisfaction of GongStruct interface
func (swing *Swing) SetName(name string) {
	swing.Name = name
}

// Stage puts sync to the model stage
func (sync *Sync) Stage(stage *Stage) *Sync {
	__gong__stage(stage.Syncs, stage.Sync_stagedOrder, stage.Sync_orderStaged, &stage.SyncOrder, stage.Syncs_mapString, sync, sync.Name)
	return sync
}

// StagePreserveOrder puts sync to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SyncOrder
// - update stage.SyncOrder accordingly
func (sync *Sync) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Syncs, stage.Sync_stagedOrder, stage.Sync_orderStaged, &stage.SyncOrder, stage.Syncs_mapString, sync, order, sync.Name)
}

// Unstage removes sync off the model stage
func (sync *Sync) Unstage(stage *Stage) *Sync {
	__gong__unstage(stage.Syncs, stage.Syncs_mapString, sync, sync.Name)
	return sync
}

// UnstageVoid removes sync off the model stage
func (sync *Sync) UnstageVoid(stage *Stage) {
	sync.Unstage(stage)
}

func (sync *Sync) StageVoid(stage *Stage) {
	sync.Stage(stage)
}

// for satisfaction of GongStruct interface
func (sync *Sync) GetName() (res string) {
	return sync.Name
}

// for satisfaction of GongStruct interface
func (sync *Sync) SetName(name string) {
	sync.Name = name
}

// Stage puts timemodification to the model stage
func (timemodification *TimeModification) Stage(stage *Stage) *TimeModification {
	__gong__stage(stage.TimeModifications, stage.TimeModification_stagedOrder, stage.TimeModification_orderStaged, &stage.TimeModificationOrder, stage.TimeModifications_mapString, timemodification, timemodification.Name)
	return timemodification
}

// StagePreserveOrder puts timemodification to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimeModificationOrder
// - update stage.TimeModificationOrder accordingly
func (timemodification *TimeModification) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TimeModifications, stage.TimeModification_stagedOrder, stage.TimeModification_orderStaged, &stage.TimeModificationOrder, stage.TimeModifications_mapString, timemodification, order, timemodification.Name)
}

// Unstage removes timemodification off the model stage
func (timemodification *TimeModification) Unstage(stage *Stage) *TimeModification {
	__gong__unstage(stage.TimeModifications, stage.TimeModifications_mapString, timemodification, timemodification.Name)
	return timemodification
}

// UnstageVoid removes timemodification off the model stage
func (timemodification *TimeModification) UnstageVoid(stage *Stage) {
	timemodification.Unstage(stage)
}

func (timemodification *TimeModification) StageVoid(stage *Stage) {
	timemodification.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timemodification *TimeModification) GetName() (res string) {
	return timemodification.Name
}

// for satisfaction of GongStruct interface
func (timemodification *TimeModification) SetName(name string) {
	timemodification.Name = name
}

// Stage puts timpani to the model stage
func (timpani *Timpani) Stage(stage *Stage) *Timpani {
	__gong__stage(stage.Timpanis, stage.Timpani_stagedOrder, stage.Timpani_orderStaged, &stage.TimpaniOrder, stage.Timpanis_mapString, timpani, timpani.Name)
	return timpani
}

// StagePreserveOrder puts timpani to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TimpaniOrder
// - update stage.TimpaniOrder accordingly
func (timpani *Timpani) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Timpanis, stage.Timpani_stagedOrder, stage.Timpani_orderStaged, &stage.TimpaniOrder, stage.Timpanis_mapString, timpani, order, timpani.Name)
}

// Unstage removes timpani off the model stage
func (timpani *Timpani) Unstage(stage *Stage) *Timpani {
	__gong__unstage(stage.Timpanis, stage.Timpanis_mapString, timpani, timpani.Name)
	return timpani
}

// UnstageVoid removes timpani off the model stage
func (timpani *Timpani) UnstageVoid(stage *Stage) {
	timpani.Unstage(stage)
}

func (timpani *Timpani) StageVoid(stage *Stage) {
	timpani.Stage(stage)
}

// for satisfaction of GongStruct interface
func (timpani *Timpani) GetName() (res string) {
	return timpani.Name
}

// for satisfaction of GongStruct interface
func (timpani *Timpani) SetName(name string) {
	timpani.Name = name
}

// Stage puts virtualinstrument to the model stage
func (virtualinstrument *VirtualInstrument) Stage(stage *Stage) *VirtualInstrument {
	__gong__stage(stage.VirtualInstruments, stage.VirtualInstrument_stagedOrder, stage.VirtualInstrument_orderStaged, &stage.VirtualInstrumentOrder, stage.VirtualInstruments_mapString, virtualinstrument, virtualinstrument.Name)
	return virtualinstrument
}

// StagePreserveOrder puts virtualinstrument to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.VirtualInstrumentOrder
// - update stage.VirtualInstrumentOrder accordingly
func (virtualinstrument *VirtualInstrument) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.VirtualInstruments, stage.VirtualInstrument_stagedOrder, stage.VirtualInstrument_orderStaged, &stage.VirtualInstrumentOrder, stage.VirtualInstruments_mapString, virtualinstrument, order, virtualinstrument.Name)
}

// Unstage removes virtualinstrument off the model stage
func (virtualinstrument *VirtualInstrument) Unstage(stage *Stage) *VirtualInstrument {
	__gong__unstage(stage.VirtualInstruments, stage.VirtualInstruments_mapString, virtualinstrument, virtualinstrument.Name)
	return virtualinstrument
}

// UnstageVoid removes virtualinstrument off the model stage
func (virtualinstrument *VirtualInstrument) UnstageVoid(stage *Stage) {
	virtualinstrument.Unstage(stage)
}

func (virtualinstrument *VirtualInstrument) StageVoid(stage *Stage) {
	virtualinstrument.Stage(stage)
}

// for satisfaction of GongStruct interface
func (virtualinstrument *VirtualInstrument) GetName() (res string) {
	return virtualinstrument.Name
}

// for satisfaction of GongStruct interface
func (virtualinstrument *VirtualInstrument) SetName(name string) {
	virtualinstrument.Name = name
}

// Stage puts wedge to the model stage
func (wedge *Wedge) Stage(stage *Stage) *Wedge {
	__gong__stage(stage.Wedges, stage.Wedge_stagedOrder, stage.Wedge_orderStaged, &stage.WedgeOrder, stage.Wedges_mapString, wedge, wedge.Name)
	return wedge
}

// StagePreserveOrder puts wedge to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.WedgeOrder
// - update stage.WedgeOrder accordingly
func (wedge *Wedge) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Wedges, stage.Wedge_stagedOrder, stage.Wedge_orderStaged, &stage.WedgeOrder, stage.Wedges_mapString, wedge, order, wedge.Name)
}

// Unstage removes wedge off the model stage
func (wedge *Wedge) Unstage(stage *Stage) *Wedge {
	__gong__unstage(stage.Wedges, stage.Wedges_mapString, wedge, wedge.Name)
	return wedge
}

// UnstageVoid removes wedge off the model stage
func (wedge *Wedge) UnstageVoid(stage *Stage) {
	wedge.Unstage(stage)
}

func (wedge *Wedge) StageVoid(stage *Stage) {
	wedge.Stage(stage)
}

// for satisfaction of GongStruct interface
func (wedge *Wedge) GetName() (res string) {
	return wedge.Name
}

// for satisfaction of GongStruct interface
func (wedge *Wedge) SetName(name string) {
	wedge.Name = name
}

// Stage puts wood to the model stage
func (wood *Wood) Stage(stage *Stage) *Wood {
	__gong__stage(stage.Woods, stage.Wood_stagedOrder, stage.Wood_orderStaged, &stage.WoodOrder, stage.Woods_mapString, wood, wood.Name)
	return wood
}

// StagePreserveOrder puts wood to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.WoodOrder
// - update stage.WoodOrder accordingly
func (wood *Wood) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Woods, stage.Wood_stagedOrder, stage.Wood_orderStaged, &stage.WoodOrder, stage.Woods_mapString, wood, order, wood.Name)
}

// Unstage removes wood off the model stage
func (wood *Wood) Unstage(stage *Stage) *Wood {
	__gong__unstage(stage.Woods, stage.Woods_mapString, wood, wood.Name)
	return wood
}

// UnstageVoid removes wood off the model stage
func (wood *Wood) UnstageVoid(stage *Stage) {
	wood.Unstage(stage)
}

func (wood *Wood) StageVoid(stage *Stage) {
	wood.Stage(stage)
}

// for satisfaction of GongStruct interface
func (wood *Wood) GetName() (res string) {
	return wood.Name
}

// for satisfaction of GongStruct interface
func (wood *Wood) SetName(name string) {
	wood.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Accords, &stage.Accords_mapString, &stage.Accord_stagedOrder, &stage.AccordOrder)

	__gong__resetStageType(&stage.AccordionRegistrations, &stage.AccordionRegistrations_mapString, &stage.AccordionRegistration_stagedOrder, &stage.AccordionRegistrationOrder)

	__gong__resetStageType(&stage.Barres, &stage.Barres_mapString, &stage.Barre_stagedOrder, &stage.BarreOrder)

	__gong__resetStageType(&stage.Basss, &stage.Basss_mapString, &stage.Bass_stagedOrder, &stage.BassOrder)

	__gong__resetStageType(&stage.BassSteps, &stage.BassSteps_mapString, &stage.BassStep_stagedOrder, &stage.BassStepOrder)

	__gong__resetStageType(&stage.BeatUnitTieds, &stage.BeatUnitTieds_mapString, &stage.BeatUnitTied_stagedOrder, &stage.BeatUnitTiedOrder)

	__gong__resetStageType(&stage.Beaters, &stage.Beaters_mapString, &stage.Beater_stagedOrder, &stage.BeaterOrder)

	__gong__resetStageType(&stage.Brackets, &stage.Brackets_mapString, &stage.Bracket_stagedOrder, &stage.BracketOrder)

	__gong__resetStageType(&stage.Dashess, &stage.Dashess_mapString, &stage.Dashes_stagedOrder, &stage.DashesOrder)

	__gong__resetStageType(&stage.Degrees, &stage.Degrees_mapString, &stage.Degree_stagedOrder, &stage.DegreeOrder)

	__gong__resetStageType(&stage.DegreeAlters, &stage.DegreeAlters_mapString, &stage.DegreeAlter_stagedOrder, &stage.DegreeAlterOrder)

	__gong__resetStageType(&stage.DegreeTypes, &stage.DegreeTypes_mapString, &stage.DegreeType_stagedOrder, &stage.DegreeTypeOrder)

	__gong__resetStageType(&stage.DegreeValues, &stage.DegreeValues_mapString, &stage.DegreeValue_stagedOrder, &stage.DegreeValueOrder)

	__gong__resetStageType(&stage.Directions, &stage.Directions_mapString, &stage.Direction_stagedOrder, &stage.DirectionOrder)

	__gong__resetStageType(&stage.DirectionTypes, &stage.DirectionTypes_mapString, &stage.DirectionType_stagedOrder, &stage.DirectionTypeOrder)

	__gong__resetStageType(&stage.Effects, &stage.Effects_mapString, &stage.Effect_stagedOrder, &stage.EffectOrder)

	__gong__resetStageType(&stage.Features, &stage.Features_mapString, &stage.Feature_stagedOrder, &stage.FeatureOrder)

	__gong__resetStageType(&stage.FirstFrets, &stage.FirstFrets_mapString, &stage.FirstFret_stagedOrder, &stage.FirstFretOrder)

	__gong__resetStageType(&stage.Frames, &stage.Frames_mapString, &stage.Frame_stagedOrder, &stage.FrameOrder)

	__gong__resetStageType(&stage.FrameNotes, &stage.FrameNotes_mapString, &stage.FrameNote_stagedOrder, &stage.FrameNoteOrder)

	__gong__resetStageType(&stage.Glasss, &stage.Glasss_mapString, &stage.Glass_stagedOrder, &stage.GlassOrder)

	__gong__resetStageType(&stage.Groupings, &stage.Groupings_mapString, &stage.Grouping_stagedOrder, &stage.GroupingOrder)

	__gong__resetStageType(&stage.Harmonys, &stage.Harmonys_mapString, &stage.Harmony_stagedOrder, &stage.HarmonyOrder)

	__gong__resetStageType(&stage.HarmonyAlters, &stage.HarmonyAlters_mapString, &stage.HarmonyAlter_stagedOrder, &stage.HarmonyAlterOrder)

	__gong__resetStageType(&stage.HarpPedalss, &stage.HarpPedalss_mapString, &stage.HarpPedals_stagedOrder, &stage.HarpPedalsOrder)

	__gong__resetStageType(&stage.Images, &stage.Images_mapString, &stage.Image_stagedOrder, &stage.ImageOrder)

	__gong__resetStageType(&stage.InstrumentChanges, &stage.InstrumentChanges_mapString, &stage.InstrumentChange_stagedOrder, &stage.InstrumentChangeOrder)

	__gong__resetStageType(&stage.Inversions, &stage.Inversions_mapString, &stage.Inversion_stagedOrder, &stage.InversionOrder)

	__gong__resetStageType(&stage.Kinds, &stage.Kinds_mapString, &stage.Kind_stagedOrder, &stage.KindOrder)

	__gong__resetStageType(&stage.Listenings, &stage.Listenings_mapString, &stage.Listening_stagedOrder, &stage.ListeningOrder)

	__gong__resetStageType(&stage.MeasureNumberings, &stage.MeasureNumberings_mapString, &stage.MeasureNumbering_stagedOrder, &stage.MeasureNumberingOrder)

	__gong__resetStageType(&stage.Membranes, &stage.Membranes_mapString, &stage.Membrane_stagedOrder, &stage.MembraneOrder)

	__gong__resetStageType(&stage.Metals, &stage.Metals_mapString, &stage.Metal_stagedOrder, &stage.MetalOrder)

	__gong__resetStageType(&stage.Metronomes, &stage.Metronomes_mapString, &stage.Metronome_stagedOrder, &stage.MetronomeOrder)

	__gong__resetStageType(&stage.MetronomeBeams, &stage.MetronomeBeams_mapString, &stage.MetronomeBeam_stagedOrder, &stage.MetronomeBeamOrder)

	__gong__resetStageType(&stage.MetronomeNotes, &stage.MetronomeNotes_mapString, &stage.MetronomeNote_stagedOrder, &stage.MetronomeNoteOrder)

	__gong__resetStageType(&stage.MetronomeTieds, &stage.MetronomeTieds_mapString, &stage.MetronomeTied_stagedOrder, &stage.MetronomeTiedOrder)

	__gong__resetStageType(&stage.MetronomeTuplets, &stage.MetronomeTuplets_mapString, &stage.MetronomeTuplet_stagedOrder, &stage.MetronomeTupletOrder)

	__gong__resetStageType(&stage.Numerals, &stage.Numerals_mapString, &stage.Numeral_stagedOrder, &stage.NumeralOrder)

	__gong__resetStageType(&stage.NumeralKeys, &stage.NumeralKeys_mapString, &stage.NumeralKey_stagedOrder, &stage.NumeralKeyOrder)

	__gong__resetStageType(&stage.NumeralRoots, &stage.NumeralRoots_mapString, &stage.NumeralRoot_stagedOrder, &stage.NumeralRootOrder)

	__gong__resetStageType(&stage.OctaveShifts, &stage.OctaveShifts_mapString, &stage.OctaveShift_stagedOrder, &stage.OctaveShiftOrder)

	__gong__resetStageType(&stage.Offsets, &stage.Offsets_mapString, &stage.Offset_stagedOrder, &stage.OffsetOrder)

	__gong__resetStageType(&stage.OtherDirections, &stage.OtherDirections_mapString, &stage.OtherDirection_stagedOrder, &stage.OtherDirectionOrder)

	__gong__resetStageType(&stage.OtherListenings, &stage.OtherListenings_mapString, &stage.OtherListening_stagedOrder, &stage.OtherListeningOrder)

	__gong__resetStageType(&stage.Pedals, &stage.Pedals_mapString, &stage.Pedal_stagedOrder, &stage.PedalOrder)

	__gong__resetStageType(&stage.PedalTunings, &stage.PedalTunings_mapString, &stage.PedalTuning_stagedOrder, &stage.PedalTuningOrder)

	__gong__resetStageType(&stage.PerMinutes, &stage.PerMinutes_mapString, &stage.PerMinute_stagedOrder, &stage.PerMinuteOrder)

	__gong__resetStageType(&stage.Percussions, &stage.Percussions_mapString, &stage.Percussion_stagedOrder, &stage.PercussionOrder)

	__gong__resetStageType(&stage.Pitcheds, &stage.Pitcheds_mapString, &stage.Pitched_stagedOrder, &stage.PitchedOrder)

	__gong__resetStageType(&stage.PrincipalVoices, &stage.PrincipalVoices_mapString, &stage.PrincipalVoice_stagedOrder, &stage.PrincipalVoiceOrder)

	__gong__resetStageType(&stage.Prints, &stage.Prints_mapString, &stage.Print_stagedOrder, &stage.PrintOrder)

	__gong__resetStageType(&stage.Roots, &stage.Roots_mapString, &stage.Root_stagedOrder, &stage.RootOrder)

	__gong__resetStageType(&stage.RootSteps, &stage.RootSteps_mapString, &stage.RootStep_stagedOrder, &stage.RootStepOrder)

	__gong__resetStageType(&stage.Scordaturas, &stage.Scordaturas_mapString, &stage.Scordatura_stagedOrder, &stage.ScordaturaOrder)

	__gong__resetStageType(&stage.Sounds, &stage.Sounds_mapString, &stage.Sound_stagedOrder, &stage.SoundOrder)

	__gong__resetStageType(&stage.StaffDivides, &stage.StaffDivides_mapString, &stage.StaffDivide_stagedOrder, &stage.StaffDivideOrder)

	__gong__resetStageType(&stage.Sticks, &stage.Sticks_mapString, &stage.Stick_stagedOrder, &stage.StickOrder)

	__gong__resetStageType(&stage.StringMutes, &stage.StringMutes_mapString, &stage.StringMute_stagedOrder, &stage.StringMuteOrder)

	__gong__resetStageType(&stage.StyleTexts, &stage.StyleTexts_mapString, &stage.StyleText_stagedOrder, &stage.StyleTextOrder)

	__gong__resetStageType(&stage.Swings, &stage.Swings_mapString, &stage.Swing_stagedOrder, &stage.SwingOrder)

	__gong__resetStageType(&stage.Syncs, &stage.Syncs_mapString, &stage.Sync_stagedOrder, &stage.SyncOrder)

	__gong__resetStageType(&stage.TimeModifications, &stage.TimeModifications_mapString, &stage.TimeModification_stagedOrder, &stage.TimeModificationOrder)

	__gong__resetStageType(&stage.Timpanis, &stage.Timpanis_mapString, &stage.Timpani_stagedOrder, &stage.TimpaniOrder)

	__gong__resetStageType(&stage.VirtualInstruments, &stage.VirtualInstruments_mapString, &stage.VirtualInstrument_stagedOrder, &stage.VirtualInstrumentOrder)

	__gong__resetStageType(&stage.Wedges, &stage.Wedges_mapString, &stage.Wedge_stagedOrder, &stage.WedgeOrder)

	__gong__resetStageType(&stage.Woods, &stage.Woods_mapString, &stage.Wood_stagedOrder, &stage.WoodOrder)

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	GongGetAssociationName() any
}

type GongstructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

type GongtructBasicField = GongstructBasicField

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) string
	GongGetUUID(stage *Stage) string
	GongAfterCreateFromFront(stage *Stage)
	GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF)
	GongAfterDeleteFromFront(stage *Stage, front GongstructIF)
	GongIsStaged(stage *Stage) bool
	GongStageBranch(stage *Stage)
	GongUnstageBranch(stage *Stage)

	GongGetInstancesByOrder(stage *Stage) any
	GongGetInstanceFromOrder(stage *Stage, order uint) any
	GongGetInstancesMapByName(stage *Stage) any
	GongGetInstancesSet(stage *Stage) any
	GongNewInstance() any
	GongGetReverseFields() []GongReverseField
}
type GongstructPtr interface {
	GongstructIF
	comparable
}

type PointerToGongstruct = GongstructPtr

func GongCompareGongstructByName[T GongstructPtr](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func GongSortGongstructSetByName[T GongstructPtr](set map[T]struct{}) (sortedSlice []T) {
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, GongCompareGongstructByName)

	return
}

// GetInstancesSorted is the Stage method returning sorted instances of a gongstruct.
func (stage *Stage) GetInstancesSorted[T GongstructPtr]() (sortedSlice []T) {
	set := stage.GetInstancesSet[T]()
	sortedSlice = GongSortGongstructSetByName(*set)

	return
}

// GetInstancesMapByName is the Stage method returning a map of staged instances by their name.
func (stage *Stage) GetInstancesMapByName[Type GongstructIF]() map[string]Type {
	if stage == nil {
		return nil
	}
	var t Type
	if any(t) == nil {
		return nil
	}
	return t.GongGetInstancesMapByName(stage).(map[string]Type)
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	if stage == nil {
		return nil
	}
	var t Type
	return t.GongGetInstancesSet(stage).(*map[Type]struct{})
}

// insertion point for instance with special fields
func (Accord) GongGetAssociationName() any {
	return &Accord{
	}
}

func (AccordionRegistration) GongGetAssociationName() any {
	return &AccordionRegistration{
	}
}

func (Barre) GongGetAssociationName() any {
	return &Barre{
	}
}

func (Bass) GongGetAssociationName() any {
	return &Bass{
			BassSeparator: &StyleText{Name: "BassSeparator"},
			BassStep: &BassStep{Name: "BassStep"},
			BassAlter: &HarmonyAlter{Name: "BassAlter"},
	}
}

func (BassStep) GongGetAssociationName() any {
	return &BassStep{
	}
}

func (BeatUnitTied) GongGetAssociationName() any {
	return &BeatUnitTied{
	}
}

func (Beater) GongGetAssociationName() any {
	return &Beater{
	}
}

func (Bracket) GongGetAssociationName() any {
	return &Bracket{
	}
}

func (Dashes) GongGetAssociationName() any {
	return &Dashes{
	}
}

func (Degree) GongGetAssociationName() any {
	return &Degree{
			DegreeValue: &DegreeValue{Name: "DegreeValue"},
			DegreeAlter: &DegreeAlter{Name: "DegreeAlter"},
			DegreeType: &DegreeType{Name: "DegreeType"},
	}
}

func (DegreeAlter) GongGetAssociationName() any {
	return &DegreeAlter{
	}
}

func (DegreeType) GongGetAssociationName() any {
	return &DegreeType{
	}
}

func (DegreeValue) GongGetAssociationName() any {
	return &DegreeValue{
	}
}

func (Direction) GongGetAssociationName() any {
	return &Direction{
			DirectionType: []*DirectionType{{Name: "DirectionType"}},
			Offset: &Offset{Name: "Offset"},
			Sound: &Sound{Name: "Sound"},
			Listening: &Listening{Name: "Listening"},
	}
}

func (DirectionType) GongGetAssociationName() any {
	return &DirectionType{
			Wedge: &Wedge{Name: "Wedge"},
			Dashes: &Dashes{Name: "Dashes"},
			Bracket: &Bracket{Name: "Bracket"},
			Pedal: &Pedal{Name: "Pedal"},
			Metronome: &Metronome{Name: "Metronome"},
			OctaveShift: &OctaveShift{Name: "OctaveShift"},
			HarpPedals: &HarpPedals{Name: "HarpPedals"},
			StringMute: &StringMute{Name: "StringMute"},
			Scordatura: &Scordatura{Name: "Scordatura"},
			Image: &Image{Name: "Image"},
			PrincipalVoice: &PrincipalVoice{Name: "PrincipalVoice"},
			Percussion: []*Percussion{{Name: "Percussion"}},
			AccordionRegistration: &AccordionRegistration{Name: "AccordionRegistration"},
			StaffDivide: &StaffDivide{Name: "StaffDivide"},
			OtherDirection: &OtherDirection{Name: "OtherDirection"},
	}
}

func (Effect) GongGetAssociationName() any {
	return &Effect{
	}
}

func (Feature) GongGetAssociationName() any {
	return &Feature{
	}
}

func (FirstFret) GongGetAssociationName() any {
	return &FirstFret{
	}
}

func (Frame) GongGetAssociationName() any {
	return &Frame{
			FirstFret: &FirstFret{Name: "FirstFret"},
			FrameNote: []*FrameNote{{Name: "FrameNote"}},
	}
}

func (FrameNote) GongGetAssociationName() any {
	return &FrameNote{
			Barre: &Barre{Name: "Barre"},
	}
}

func (Glass) GongGetAssociationName() any {
	return &Glass{
	}
}

func (Grouping) GongGetAssociationName() any {
	return &Grouping{
			Feature: []*Feature{{Name: "Feature"}},
	}
}

func (Harmony) GongGetAssociationName() any {
	return &Harmony{
			Root: []*Root{{Name: "Root"}},
			Numeral: []*Numeral{{Name: "Numeral"}},
			Function: []*StyleText{{Name: "Function"}},
			Kind: []*Kind{{Name: "Kind"}},
			Inversion: []*Inversion{{Name: "Inversion"}},
			Bass: []*Bass{{Name: "Bass"}},
			Degree: []*Degree{{Name: "Degree"}},
			Frame: &Frame{Name: "Frame"},
			Offset: &Offset{Name: "Offset"},
	}
}

func (HarmonyAlter) GongGetAssociationName() any {
	return &HarmonyAlter{
	}
}

func (HarpPedals) GongGetAssociationName() any {
	return &HarpPedals{
			PedalTuning: []*PedalTuning{{Name: "PedalTuning"}},
	}
}

func (Image) GongGetAssociationName() any {
	return &Image{
	}
}

func (InstrumentChange) GongGetAssociationName() any {
	return &InstrumentChange{
			VirtualInstrument: &VirtualInstrument{Name: "VirtualInstrument"},
	}
}

func (Inversion) GongGetAssociationName() any {
	return &Inversion{
	}
}

func (Kind) GongGetAssociationName() any {
	return &Kind{
	}
}

func (Listening) GongGetAssociationName() any {
	return &Listening{
			Sync: []*Sync{{Name: "Sync"}},
			OtherListening: []*OtherListening{{Name: "OtherListening"}},
			Offset: &Offset{Name: "Offset"},
	}
}

func (MeasureNumbering) GongGetAssociationName() any {
	return &MeasureNumbering{
	}
}

func (Membrane) GongGetAssociationName() any {
	return &Membrane{
	}
}

func (Metal) GongGetAssociationName() any {
	return &Metal{
	}
}

func (Metronome) GongGetAssociationName() any {
	return &Metronome{
			BeatUnitTied: []*BeatUnitTied{{Name: "BeatUnitTied"}},
			PerMinute: &PerMinute{Name: "PerMinute"},
			MetronomeNote: []*MetronomeNote{{Name: "MetronomeNote"}},
	}
}

func (MetronomeBeam) GongGetAssociationName() any {
	return &MetronomeBeam{
	}
}

func (MetronomeNote) GongGetAssociationName() any {
	return &MetronomeNote{
			MetronomeBeam: []*MetronomeBeam{{Name: "MetronomeBeam"}},
			MetronomeTied: &MetronomeTied{Name: "MetronomeTied"},
			MetronomeTuplet: &MetronomeTuplet{Name: "MetronomeTuplet"},
	}
}

func (MetronomeTied) GongGetAssociationName() any {
	return &MetronomeTied{
	}
}

func (MetronomeTuplet) GongGetAssociationName() any {
	return &MetronomeTuplet{
	}
}

func (Numeral) GongGetAssociationName() any {
	return &Numeral{
			NumeralRoot: &NumeralRoot{Name: "NumeralRoot"},
			NumeralAlter: &HarmonyAlter{Name: "NumeralAlter"},
			NumeralKey: &NumeralKey{Name: "NumeralKey"},
	}
}

func (NumeralKey) GongGetAssociationName() any {
	return &NumeralKey{
	}
}

func (NumeralRoot) GongGetAssociationName() any {
	return &NumeralRoot{
	}
}

func (OctaveShift) GongGetAssociationName() any {
	return &OctaveShift{
	}
}

func (Offset) GongGetAssociationName() any {
	return &Offset{
	}
}

func (OtherDirection) GongGetAssociationName() any {
	return &OtherDirection{
	}
}

func (OtherListening) GongGetAssociationName() any {
	return &OtherListening{
	}
}

func (Pedal) GongGetAssociationName() any {
	return &Pedal{
	}
}

func (PedalTuning) GongGetAssociationName() any {
	return &PedalTuning{
	}
}

func (PerMinute) GongGetAssociationName() any {
	return &PerMinute{
	}
}

func (Percussion) GongGetAssociationName() any {
	return &Percussion{
			Glass: &Glass{Name: "Glass"},
			Metal: &Metal{Name: "Metal"},
			Wood: &Wood{Name: "Wood"},
			Pitched: &Pitched{Name: "Pitched"},
			Membrane: &Membrane{Name: "Membrane"},
			Effect: &Effect{Name: "Effect"},
			Timpani: &Timpani{Name: "Timpani"},
			Beater: &Beater{Name: "Beater"},
			Stick: &Stick{Name: "Stick"},
	}
}

func (Pitched) GongGetAssociationName() any {
	return &Pitched{
	}
}

func (PrincipalVoice) GongGetAssociationName() any {
	return &PrincipalVoice{
	}
}

func (Print) GongGetAssociationName() any {
	return &Print{
			MeasureNumbering: &MeasureNumbering{Name: "MeasureNumbering"},
	}
}

func (Root) GongGetAssociationName() any {
	return &Root{
			RootStep: &RootStep{Name: "RootStep"},
			RootAlter: &HarmonyAlter{Name: "RootAlter"},
	}
}

func (RootStep) GongGetAssociationName() any {
	return &RootStep{
	}
}

func (Scordatura) GongGetAssociationName() any {
	return &Scordatura{
			Accord: []*Accord{{Name: "Accord"}},
	}
}

func (Sound) GongGetAssociationName() any {
	return &Sound{
			InstrumentChange: []*InstrumentChange{{Name: "InstrumentChange"}},
			Swing: &Swing{Name: "Swing"},
			Offset: &Offset{Name: "Offset"},
	}
}

func (StaffDivide) GongGetAssociationName() any {
	return &StaffDivide{
	}
}

func (Stick) GongGetAssociationName() any {
	return &Stick{
	}
}

func (StringMute) GongGetAssociationName() any {
	return &StringMute{
	}
}

func (StyleText) GongGetAssociationName() any {
	return &StyleText{
	}
}

func (Swing) GongGetAssociationName() any {
	return &Swing{
	}
}

func (Sync) GongGetAssociationName() any {
	return &Sync{
	}
}

func (TimeModification) GongGetAssociationName() any {
	return &TimeModification{
	}
}

func (Timpani) GongGetAssociationName() any {
	return &Timpani{
	}
}

func (VirtualInstrument) GongGetAssociationName() any {
	return &VirtualInstrument{
	}
}

func (Wedge) GongGetAssociationName() any {
	return &Wedge{
	}
}

func (Wood) GongGetAssociationName() any {
	return &Wood{
	}
}


// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var t Type
	return t.GongGetAssociationName().(*Type)
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
// GetPointerReverseMap is the Stage method for backtrack navigation of pointer associations.
func (stage *Stage) GetPointerReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Accord
	case Accord:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AccordionRegistration
	case AccordionRegistration:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Barre
	case Barre:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bass
	case Bass:
		switch fieldname {
		// insertion point for per direct association field
		case "BassSeparator":
			res := make(map[*StyleText][]*Bass)
			for bass := range stage.Basss {
				if bass.BassSeparator != nil {
					styletext_ := bass.BassSeparator
					var basss []*Bass
					_, ok := res[styletext_]
					if ok {
						basss = res[styletext_]
					} else {
						basss = make([]*Bass, 0)
					}
					basss = append(basss, bass)
					res[styletext_] = basss
				}
			}
			return any(res).(map[*End][]*Start)
		case "BassStep":
			res := make(map[*BassStep][]*Bass)
			for bass := range stage.Basss {
				if bass.BassStep != nil {
					bassstep_ := bass.BassStep
					var basss []*Bass
					_, ok := res[bassstep_]
					if ok {
						basss = res[bassstep_]
					} else {
						basss = make([]*Bass, 0)
					}
					basss = append(basss, bass)
					res[bassstep_] = basss
				}
			}
			return any(res).(map[*End][]*Start)
		case "BassAlter":
			res := make(map[*HarmonyAlter][]*Bass)
			for bass := range stage.Basss {
				if bass.BassAlter != nil {
					harmonyalter_ := bass.BassAlter
					var basss []*Bass
					_, ok := res[harmonyalter_]
					if ok {
						basss = res[harmonyalter_]
					} else {
						basss = make([]*Bass, 0)
					}
					basss = append(basss, bass)
					res[harmonyalter_] = basss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BassStep
	case BassStep:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BeatUnitTied
	case BeatUnitTied:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Beater
	case Beater:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bracket
	case Bracket:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Dashes
	case Dashes:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Degree
	case Degree:
		switch fieldname {
		// insertion point for per direct association field
		case "DegreeValue":
			res := make(map[*DegreeValue][]*Degree)
			for degree := range stage.Degrees {
				if degree.DegreeValue != nil {
					degreevalue_ := degree.DegreeValue
					var degrees []*Degree
					_, ok := res[degreevalue_]
					if ok {
						degrees = res[degreevalue_]
					} else {
						degrees = make([]*Degree, 0)
					}
					degrees = append(degrees, degree)
					res[degreevalue_] = degrees
				}
			}
			return any(res).(map[*End][]*Start)
		case "DegreeAlter":
			res := make(map[*DegreeAlter][]*Degree)
			for degree := range stage.Degrees {
				if degree.DegreeAlter != nil {
					degreealter_ := degree.DegreeAlter
					var degrees []*Degree
					_, ok := res[degreealter_]
					if ok {
						degrees = res[degreealter_]
					} else {
						degrees = make([]*Degree, 0)
					}
					degrees = append(degrees, degree)
					res[degreealter_] = degrees
				}
			}
			return any(res).(map[*End][]*Start)
		case "DegreeType":
			res := make(map[*DegreeType][]*Degree)
			for degree := range stage.Degrees {
				if degree.DegreeType != nil {
					degreetype_ := degree.DegreeType
					var degrees []*Degree
					_, ok := res[degreetype_]
					if ok {
						degrees = res[degreetype_]
					} else {
						degrees = make([]*Degree, 0)
					}
					degrees = append(degrees, degree)
					res[degreetype_] = degrees
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DegreeAlter
	case DegreeAlter:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DegreeType
	case DegreeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DegreeValue
	case DegreeValue:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Direction
	case Direction:
		switch fieldname {
		// insertion point for per direct association field
		case "Offset":
			res := make(map[*Offset][]*Direction)
			for direction := range stage.Directions {
				if direction.Offset != nil {
					offset_ := direction.Offset
					var directions []*Direction
					_, ok := res[offset_]
					if ok {
						directions = res[offset_]
					} else {
						directions = make([]*Direction, 0)
					}
					directions = append(directions, direction)
					res[offset_] = directions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Sound":
			res := make(map[*Sound][]*Direction)
			for direction := range stage.Directions {
				if direction.Sound != nil {
					sound_ := direction.Sound
					var directions []*Direction
					_, ok := res[sound_]
					if ok {
						directions = res[sound_]
					} else {
						directions = make([]*Direction, 0)
					}
					directions = append(directions, direction)
					res[sound_] = directions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Listening":
			res := make(map[*Listening][]*Direction)
			for direction := range stage.Directions {
				if direction.Listening != nil {
					listening_ := direction.Listening
					var directions []*Direction
					_, ok := res[listening_]
					if ok {
						directions = res[listening_]
					} else {
						directions = make([]*Direction, 0)
					}
					directions = append(directions, direction)
					res[listening_] = directions
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DirectionType
	case DirectionType:
		switch fieldname {
		// insertion point for per direct association field
		case "Wedge":
			res := make(map[*Wedge][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.Wedge != nil {
					wedge_ := directiontype.Wedge
					var directiontypes []*DirectionType
					_, ok := res[wedge_]
					if ok {
						directiontypes = res[wedge_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[wedge_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Dashes":
			res := make(map[*Dashes][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.Dashes != nil {
					dashes_ := directiontype.Dashes
					var directiontypes []*DirectionType
					_, ok := res[dashes_]
					if ok {
						directiontypes = res[dashes_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[dashes_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Bracket":
			res := make(map[*Bracket][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.Bracket != nil {
					bracket_ := directiontype.Bracket
					var directiontypes []*DirectionType
					_, ok := res[bracket_]
					if ok {
						directiontypes = res[bracket_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[bracket_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pedal":
			res := make(map[*Pedal][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.Pedal != nil {
					pedal_ := directiontype.Pedal
					var directiontypes []*DirectionType
					_, ok := res[pedal_]
					if ok {
						directiontypes = res[pedal_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[pedal_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Metronome":
			res := make(map[*Metronome][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.Metronome != nil {
					metronome_ := directiontype.Metronome
					var directiontypes []*DirectionType
					_, ok := res[metronome_]
					if ok {
						directiontypes = res[metronome_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[metronome_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "OctaveShift":
			res := make(map[*OctaveShift][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.OctaveShift != nil {
					octaveshift_ := directiontype.OctaveShift
					var directiontypes []*DirectionType
					_, ok := res[octaveshift_]
					if ok {
						directiontypes = res[octaveshift_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[octaveshift_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "HarpPedals":
			res := make(map[*HarpPedals][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.HarpPedals != nil {
					harppedals_ := directiontype.HarpPedals
					var directiontypes []*DirectionType
					_, ok := res[harppedals_]
					if ok {
						directiontypes = res[harppedals_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[harppedals_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "StringMute":
			res := make(map[*StringMute][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.StringMute != nil {
					stringmute_ := directiontype.StringMute
					var directiontypes []*DirectionType
					_, ok := res[stringmute_]
					if ok {
						directiontypes = res[stringmute_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[stringmute_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Scordatura":
			res := make(map[*Scordatura][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.Scordatura != nil {
					scordatura_ := directiontype.Scordatura
					var directiontypes []*DirectionType
					_, ok := res[scordatura_]
					if ok {
						directiontypes = res[scordatura_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[scordatura_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Image":
			res := make(map[*Image][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.Image != nil {
					image_ := directiontype.Image
					var directiontypes []*DirectionType
					_, ok := res[image_]
					if ok {
						directiontypes = res[image_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[image_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "PrincipalVoice":
			res := make(map[*PrincipalVoice][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.PrincipalVoice != nil {
					principalvoice_ := directiontype.PrincipalVoice
					var directiontypes []*DirectionType
					_, ok := res[principalvoice_]
					if ok {
						directiontypes = res[principalvoice_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[principalvoice_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "AccordionRegistration":
			res := make(map[*AccordionRegistration][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.AccordionRegistration != nil {
					accordionregistration_ := directiontype.AccordionRegistration
					var directiontypes []*DirectionType
					_, ok := res[accordionregistration_]
					if ok {
						directiontypes = res[accordionregistration_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[accordionregistration_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "StaffDivide":
			res := make(map[*StaffDivide][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.StaffDivide != nil {
					staffdivide_ := directiontype.StaffDivide
					var directiontypes []*DirectionType
					_, ok := res[staffdivide_]
					if ok {
						directiontypes = res[staffdivide_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[staffdivide_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "OtherDirection":
			res := make(map[*OtherDirection][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				if directiontype.OtherDirection != nil {
					otherdirection_ := directiontype.OtherDirection
					var directiontypes []*DirectionType
					_, ok := res[otherdirection_]
					if ok {
						directiontypes = res[otherdirection_]
					} else {
						directiontypes = make([]*DirectionType, 0)
					}
					directiontypes = append(directiontypes, directiontype)
					res[otherdirection_] = directiontypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Effect
	case Effect:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Feature
	case Feature:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FirstFret
	case FirstFret:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Frame
	case Frame:
		switch fieldname {
		// insertion point for per direct association field
		case "FirstFret":
			res := make(map[*FirstFret][]*Frame)
			for frame := range stage.Frames {
				if frame.FirstFret != nil {
					firstfret_ := frame.FirstFret
					var frames []*Frame
					_, ok := res[firstfret_]
					if ok {
						frames = res[firstfret_]
					} else {
						frames = make([]*Frame, 0)
					}
					frames = append(frames, frame)
					res[firstfret_] = frames
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FrameNote
	case FrameNote:
		switch fieldname {
		// insertion point for per direct association field
		case "Barre":
			res := make(map[*Barre][]*FrameNote)
			for framenote := range stage.FrameNotes {
				if framenote.Barre != nil {
					barre_ := framenote.Barre
					var framenotes []*FrameNote
					_, ok := res[barre_]
					if ok {
						framenotes = res[barre_]
					} else {
						framenotes = make([]*FrameNote, 0)
					}
					framenotes = append(framenotes, framenote)
					res[barre_] = framenotes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Glass
	case Glass:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Grouping
	case Grouping:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Harmony
	case Harmony:
		switch fieldname {
		// insertion point for per direct association field
		case "Frame":
			res := make(map[*Frame][]*Harmony)
			for harmony := range stage.Harmonys {
				if harmony.Frame != nil {
					frame_ := harmony.Frame
					var harmonys []*Harmony
					_, ok := res[frame_]
					if ok {
						harmonys = res[frame_]
					} else {
						harmonys = make([]*Harmony, 0)
					}
					harmonys = append(harmonys, harmony)
					res[frame_] = harmonys
				}
			}
			return any(res).(map[*End][]*Start)
		case "Offset":
			res := make(map[*Offset][]*Harmony)
			for harmony := range stage.Harmonys {
				if harmony.Offset != nil {
					offset_ := harmony.Offset
					var harmonys []*Harmony
					_, ok := res[offset_]
					if ok {
						harmonys = res[offset_]
					} else {
						harmonys = make([]*Harmony, 0)
					}
					harmonys = append(harmonys, harmony)
					res[offset_] = harmonys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of HarmonyAlter
	case HarmonyAlter:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HarpPedals
	case HarpPedals:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Image
	case Image:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InstrumentChange
	case InstrumentChange:
		switch fieldname {
		// insertion point for per direct association field
		case "VirtualInstrument":
			res := make(map[*VirtualInstrument][]*InstrumentChange)
			for instrumentchange := range stage.InstrumentChanges {
				if instrumentchange.VirtualInstrument != nil {
					virtualinstrument_ := instrumentchange.VirtualInstrument
					var instrumentchanges []*InstrumentChange
					_, ok := res[virtualinstrument_]
					if ok {
						instrumentchanges = res[virtualinstrument_]
					} else {
						instrumentchanges = make([]*InstrumentChange, 0)
					}
					instrumentchanges = append(instrumentchanges, instrumentchange)
					res[virtualinstrument_] = instrumentchanges
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Inversion
	case Inversion:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Kind
	case Kind:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Listening
	case Listening:
		switch fieldname {
		// insertion point for per direct association field
		case "Offset":
			res := make(map[*Offset][]*Listening)
			for listening := range stage.Listenings {
				if listening.Offset != nil {
					offset_ := listening.Offset
					var listenings []*Listening
					_, ok := res[offset_]
					if ok {
						listenings = res[offset_]
					} else {
						listenings = make([]*Listening, 0)
					}
					listenings = append(listenings, listening)
					res[offset_] = listenings
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MeasureNumbering
	case MeasureNumbering:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Membrane
	case Membrane:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Metal
	case Metal:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Metronome
	case Metronome:
		switch fieldname {
		// insertion point for per direct association field
		case "PerMinute":
			res := make(map[*PerMinute][]*Metronome)
			for metronome := range stage.Metronomes {
				if metronome.PerMinute != nil {
					perminute_ := metronome.PerMinute
					var metronomes []*Metronome
					_, ok := res[perminute_]
					if ok {
						metronomes = res[perminute_]
					} else {
						metronomes = make([]*Metronome, 0)
					}
					metronomes = append(metronomes, metronome)
					res[perminute_] = metronomes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MetronomeBeam
	case MetronomeBeam:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MetronomeNote
	case MetronomeNote:
		switch fieldname {
		// insertion point for per direct association field
		case "MetronomeTied":
			res := make(map[*MetronomeTied][]*MetronomeNote)
			for metronomenote := range stage.MetronomeNotes {
				if metronomenote.MetronomeTied != nil {
					metronometied_ := metronomenote.MetronomeTied
					var metronomenotes []*MetronomeNote
					_, ok := res[metronometied_]
					if ok {
						metronomenotes = res[metronometied_]
					} else {
						metronomenotes = make([]*MetronomeNote, 0)
					}
					metronomenotes = append(metronomenotes, metronomenote)
					res[metronometied_] = metronomenotes
				}
			}
			return any(res).(map[*End][]*Start)
		case "MetronomeTuplet":
			res := make(map[*MetronomeTuplet][]*MetronomeNote)
			for metronomenote := range stage.MetronomeNotes {
				if metronomenote.MetronomeTuplet != nil {
					metronometuplet_ := metronomenote.MetronomeTuplet
					var metronomenotes []*MetronomeNote
					_, ok := res[metronometuplet_]
					if ok {
						metronomenotes = res[metronometuplet_]
					} else {
						metronomenotes = make([]*MetronomeNote, 0)
					}
					metronomenotes = append(metronomenotes, metronomenote)
					res[metronometuplet_] = metronomenotes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MetronomeTied
	case MetronomeTied:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MetronomeTuplet
	case MetronomeTuplet:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Numeral
	case Numeral:
		switch fieldname {
		// insertion point for per direct association field
		case "NumeralRoot":
			res := make(map[*NumeralRoot][]*Numeral)
			for numeral := range stage.Numerals {
				if numeral.NumeralRoot != nil {
					numeralroot_ := numeral.NumeralRoot
					var numerals []*Numeral
					_, ok := res[numeralroot_]
					if ok {
						numerals = res[numeralroot_]
					} else {
						numerals = make([]*Numeral, 0)
					}
					numerals = append(numerals, numeral)
					res[numeralroot_] = numerals
				}
			}
			return any(res).(map[*End][]*Start)
		case "NumeralAlter":
			res := make(map[*HarmonyAlter][]*Numeral)
			for numeral := range stage.Numerals {
				if numeral.NumeralAlter != nil {
					harmonyalter_ := numeral.NumeralAlter
					var numerals []*Numeral
					_, ok := res[harmonyalter_]
					if ok {
						numerals = res[harmonyalter_]
					} else {
						numerals = make([]*Numeral, 0)
					}
					numerals = append(numerals, numeral)
					res[harmonyalter_] = numerals
				}
			}
			return any(res).(map[*End][]*Start)
		case "NumeralKey":
			res := make(map[*NumeralKey][]*Numeral)
			for numeral := range stage.Numerals {
				if numeral.NumeralKey != nil {
					numeralkey_ := numeral.NumeralKey
					var numerals []*Numeral
					_, ok := res[numeralkey_]
					if ok {
						numerals = res[numeralkey_]
					} else {
						numerals = make([]*Numeral, 0)
					}
					numerals = append(numerals, numeral)
					res[numeralkey_] = numerals
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NumeralKey
	case NumeralKey:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NumeralRoot
	case NumeralRoot:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OctaveShift
	case OctaveShift:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Offset
	case Offset:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherDirection
	case OtherDirection:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherListening
	case OtherListening:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Pedal
	case Pedal:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PedalTuning
	case PedalTuning:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PerMinute
	case PerMinute:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Percussion
	case Percussion:
		switch fieldname {
		// insertion point for per direct association field
		case "Glass":
			res := make(map[*Glass][]*Percussion)
			for percussion := range stage.Percussions {
				if percussion.Glass != nil {
					glass_ := percussion.Glass
					var percussions []*Percussion
					_, ok := res[glass_]
					if ok {
						percussions = res[glass_]
					} else {
						percussions = make([]*Percussion, 0)
					}
					percussions = append(percussions, percussion)
					res[glass_] = percussions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Metal":
			res := make(map[*Metal][]*Percussion)
			for percussion := range stage.Percussions {
				if percussion.Metal != nil {
					metal_ := percussion.Metal
					var percussions []*Percussion
					_, ok := res[metal_]
					if ok {
						percussions = res[metal_]
					} else {
						percussions = make([]*Percussion, 0)
					}
					percussions = append(percussions, percussion)
					res[metal_] = percussions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Wood":
			res := make(map[*Wood][]*Percussion)
			for percussion := range stage.Percussions {
				if percussion.Wood != nil {
					wood_ := percussion.Wood
					var percussions []*Percussion
					_, ok := res[wood_]
					if ok {
						percussions = res[wood_]
					} else {
						percussions = make([]*Percussion, 0)
					}
					percussions = append(percussions, percussion)
					res[wood_] = percussions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pitched":
			res := make(map[*Pitched][]*Percussion)
			for percussion := range stage.Percussions {
				if percussion.Pitched != nil {
					pitched_ := percussion.Pitched
					var percussions []*Percussion
					_, ok := res[pitched_]
					if ok {
						percussions = res[pitched_]
					} else {
						percussions = make([]*Percussion, 0)
					}
					percussions = append(percussions, percussion)
					res[pitched_] = percussions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Membrane":
			res := make(map[*Membrane][]*Percussion)
			for percussion := range stage.Percussions {
				if percussion.Membrane != nil {
					membrane_ := percussion.Membrane
					var percussions []*Percussion
					_, ok := res[membrane_]
					if ok {
						percussions = res[membrane_]
					} else {
						percussions = make([]*Percussion, 0)
					}
					percussions = append(percussions, percussion)
					res[membrane_] = percussions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Effect":
			res := make(map[*Effect][]*Percussion)
			for percussion := range stage.Percussions {
				if percussion.Effect != nil {
					effect_ := percussion.Effect
					var percussions []*Percussion
					_, ok := res[effect_]
					if ok {
						percussions = res[effect_]
					} else {
						percussions = make([]*Percussion, 0)
					}
					percussions = append(percussions, percussion)
					res[effect_] = percussions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Timpani":
			res := make(map[*Timpani][]*Percussion)
			for percussion := range stage.Percussions {
				if percussion.Timpani != nil {
					timpani_ := percussion.Timpani
					var percussions []*Percussion
					_, ok := res[timpani_]
					if ok {
						percussions = res[timpani_]
					} else {
						percussions = make([]*Percussion, 0)
					}
					percussions = append(percussions, percussion)
					res[timpani_] = percussions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Beater":
			res := make(map[*Beater][]*Percussion)
			for percussion := range stage.Percussions {
				if percussion.Beater != nil {
					beater_ := percussion.Beater
					var percussions []*Percussion
					_, ok := res[beater_]
					if ok {
						percussions = res[beater_]
					} else {
						percussions = make([]*Percussion, 0)
					}
					percussions = append(percussions, percussion)
					res[beater_] = percussions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Stick":
			res := make(map[*Stick][]*Percussion)
			for percussion := range stage.Percussions {
				if percussion.Stick != nil {
					stick_ := percussion.Stick
					var percussions []*Percussion
					_, ok := res[stick_]
					if ok {
						percussions = res[stick_]
					} else {
						percussions = make([]*Percussion, 0)
					}
					percussions = append(percussions, percussion)
					res[stick_] = percussions
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Pitched
	case Pitched:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PrincipalVoice
	case PrincipalVoice:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Print
	case Print:
		switch fieldname {
		// insertion point for per direct association field
		case "MeasureNumbering":
			res := make(map[*MeasureNumbering][]*Print)
			for print := range stage.Prints {
				if print.MeasureNumbering != nil {
					measurenumbering_ := print.MeasureNumbering
					var prints []*Print
					_, ok := res[measurenumbering_]
					if ok {
						prints = res[measurenumbering_]
					} else {
						prints = make([]*Print, 0)
					}
					prints = append(prints, print)
					res[measurenumbering_] = prints
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Root
	case Root:
		switch fieldname {
		// insertion point for per direct association field
		case "RootStep":
			res := make(map[*RootStep][]*Root)
			for root := range stage.Roots {
				if root.RootStep != nil {
					rootstep_ := root.RootStep
					var roots []*Root
					_, ok := res[rootstep_]
					if ok {
						roots = res[rootstep_]
					} else {
						roots = make([]*Root, 0)
					}
					roots = append(roots, root)
					res[rootstep_] = roots
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootAlter":
			res := make(map[*HarmonyAlter][]*Root)
			for root := range stage.Roots {
				if root.RootAlter != nil {
					harmonyalter_ := root.RootAlter
					var roots []*Root
					_, ok := res[harmonyalter_]
					if ok {
						roots = res[harmonyalter_]
					} else {
						roots = make([]*Root, 0)
					}
					roots = append(roots, root)
					res[harmonyalter_] = roots
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RootStep
	case RootStep:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Scordatura
	case Scordatura:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Sound
	case Sound:
		switch fieldname {
		// insertion point for per direct association field
		case "Swing":
			res := make(map[*Swing][]*Sound)
			for sound := range stage.Sounds {
				if sound.Swing != nil {
					swing_ := sound.Swing
					var sounds []*Sound
					_, ok := res[swing_]
					if ok {
						sounds = res[swing_]
					} else {
						sounds = make([]*Sound, 0)
					}
					sounds = append(sounds, sound)
					res[swing_] = sounds
				}
			}
			return any(res).(map[*End][]*Start)
		case "Offset":
			res := make(map[*Offset][]*Sound)
			for sound := range stage.Sounds {
				if sound.Offset != nil {
					offset_ := sound.Offset
					var sounds []*Sound
					_, ok := res[offset_]
					if ok {
						sounds = res[offset_]
					} else {
						sounds = make([]*Sound, 0)
					}
					sounds = append(sounds, sound)
					res[offset_] = sounds
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StaffDivide
	case StaffDivide:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Stick
	case Stick:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StringMute
	case StringMute:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StyleText
	case StyleText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Swing
	case Swing:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Sync
	case Sync:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimeModification
	case TimeModification:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Timpani
	case Timpani:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VirtualInstrument
	case VirtualInstrument:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Wedge
	case Wedge:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Wood
	case Wood:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Accord
	case Accord:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AccordionRegistration
	case AccordionRegistration:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Barre
	case Barre:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bass
	case Bass:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BassStep
	case BassStep:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BeatUnitTied
	case BeatUnitTied:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Beater
	case Beater:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bracket
	case Bracket:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Dashes
	case Dashes:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Degree
	case Degree:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DegreeAlter
	case DegreeAlter:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DegreeType
	case DegreeType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DegreeValue
	case DegreeValue:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Direction
	case Direction:
		switch fieldname {
		// insertion point for per direct association field
		case "DirectionType":
			res := make(map[*DirectionType][]*Direction)
			for direction := range stage.Directions {
				for _, directiontype_ := range direction.DirectionType {
					res[directiontype_] = append(res[directiontype_], direction)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DirectionType
	case DirectionType:
		switch fieldname {
		// insertion point for per direct association field
		case "Percussion":
			res := make(map[*Percussion][]*DirectionType)
			for directiontype := range stage.DirectionTypes {
				for _, percussion_ := range directiontype.Percussion {
					res[percussion_] = append(res[percussion_], directiontype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Effect
	case Effect:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Feature
	case Feature:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FirstFret
	case FirstFret:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Frame
	case Frame:
		switch fieldname {
		// insertion point for per direct association field
		case "FrameNote":
			res := make(map[*FrameNote][]*Frame)
			for frame := range stage.Frames {
				for _, framenote_ := range frame.FrameNote {
					res[framenote_] = append(res[framenote_], frame)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FrameNote
	case FrameNote:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Glass
	case Glass:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Grouping
	case Grouping:
		switch fieldname {
		// insertion point for per direct association field
		case "Feature":
			res := make(map[*Feature][]*Grouping)
			for grouping := range stage.Groupings {
				for _, feature_ := range grouping.Feature {
					res[feature_] = append(res[feature_], grouping)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Harmony
	case Harmony:
		switch fieldname {
		// insertion point for per direct association field
		case "Root":
			res := make(map[*Root][]*Harmony)
			for harmony := range stage.Harmonys {
				for _, root_ := range harmony.Root {
					res[root_] = append(res[root_], harmony)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Numeral":
			res := make(map[*Numeral][]*Harmony)
			for harmony := range stage.Harmonys {
				for _, numeral_ := range harmony.Numeral {
					res[numeral_] = append(res[numeral_], harmony)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Function":
			res := make(map[*StyleText][]*Harmony)
			for harmony := range stage.Harmonys {
				for _, styletext_ := range harmony.Function {
					res[styletext_] = append(res[styletext_], harmony)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Kind":
			res := make(map[*Kind][]*Harmony)
			for harmony := range stage.Harmonys {
				for _, kind_ := range harmony.Kind {
					res[kind_] = append(res[kind_], harmony)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Inversion":
			res := make(map[*Inversion][]*Harmony)
			for harmony := range stage.Harmonys {
				for _, inversion_ := range harmony.Inversion {
					res[inversion_] = append(res[inversion_], harmony)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Bass":
			res := make(map[*Bass][]*Harmony)
			for harmony := range stage.Harmonys {
				for _, bass_ := range harmony.Bass {
					res[bass_] = append(res[bass_], harmony)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Degree":
			res := make(map[*Degree][]*Harmony)
			for harmony := range stage.Harmonys {
				for _, degree_ := range harmony.Degree {
					res[degree_] = append(res[degree_], harmony)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of HarmonyAlter
	case HarmonyAlter:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HarpPedals
	case HarpPedals:
		switch fieldname {
		// insertion point for per direct association field
		case "PedalTuning":
			res := make(map[*PedalTuning][]*HarpPedals)
			for harppedals := range stage.HarpPedalss {
				for _, pedaltuning_ := range harppedals.PedalTuning {
					res[pedaltuning_] = append(res[pedaltuning_], harppedals)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Image
	case Image:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of InstrumentChange
	case InstrumentChange:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Inversion
	case Inversion:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Kind
	case Kind:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Listening
	case Listening:
		switch fieldname {
		// insertion point for per direct association field
		case "Sync":
			res := make(map[*Sync][]*Listening)
			for listening := range stage.Listenings {
				for _, sync_ := range listening.Sync {
					res[sync_] = append(res[sync_], listening)
				}
			}
			return any(res).(map[*End][]*Start)
		case "OtherListening":
			res := make(map[*OtherListening][]*Listening)
			for listening := range stage.Listenings {
				for _, otherlistening_ := range listening.OtherListening {
					res[otherlistening_] = append(res[otherlistening_], listening)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MeasureNumbering
	case MeasureNumbering:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Membrane
	case Membrane:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Metal
	case Metal:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Metronome
	case Metronome:
		switch fieldname {
		// insertion point for per direct association field
		case "BeatUnitTied":
			res := make(map[*BeatUnitTied][]*Metronome)
			for metronome := range stage.Metronomes {
				for _, beatunittied_ := range metronome.BeatUnitTied {
					res[beatunittied_] = append(res[beatunittied_], metronome)
				}
			}
			return any(res).(map[*End][]*Start)
		case "MetronomeNote":
			res := make(map[*MetronomeNote][]*Metronome)
			for metronome := range stage.Metronomes {
				for _, metronomenote_ := range metronome.MetronomeNote {
					res[metronomenote_] = append(res[metronomenote_], metronome)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MetronomeBeam
	case MetronomeBeam:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MetronomeNote
	case MetronomeNote:
		switch fieldname {
		// insertion point for per direct association field
		case "MetronomeBeam":
			res := make(map[*MetronomeBeam][]*MetronomeNote)
			for metronomenote := range stage.MetronomeNotes {
				for _, metronomebeam_ := range metronomenote.MetronomeBeam {
					res[metronomebeam_] = append(res[metronomebeam_], metronomenote)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MetronomeTied
	case MetronomeTied:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MetronomeTuplet
	case MetronomeTuplet:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Numeral
	case Numeral:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NumeralKey
	case NumeralKey:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NumeralRoot
	case NumeralRoot:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OctaveShift
	case OctaveShift:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Offset
	case Offset:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherDirection
	case OtherDirection:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherListening
	case OtherListening:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Pedal
	case Pedal:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PedalTuning
	case PedalTuning:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PerMinute
	case PerMinute:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Percussion
	case Percussion:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Pitched
	case Pitched:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PrincipalVoice
	case PrincipalVoice:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Print
	case Print:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Root
	case Root:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RootStep
	case RootStep:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Scordatura
	case Scordatura:
		switch fieldname {
		// insertion point for per direct association field
		case "Accord":
			res := make(map[*Accord][]*Scordatura)
			for scordatura := range stage.Scordaturas {
				for _, accord_ := range scordatura.Accord {
					res[accord_] = append(res[accord_], scordatura)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Sound
	case Sound:
		switch fieldname {
		// insertion point for per direct association field
		case "InstrumentChange":
			res := make(map[*InstrumentChange][]*Sound)
			for sound := range stage.Sounds {
				for _, instrumentchange_ := range sound.InstrumentChange {
					res[instrumentchange_] = append(res[instrumentchange_], sound)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StaffDivide
	case StaffDivide:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Stick
	case Stick:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StringMute
	case StringMute:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StyleText
	case StyleText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Swing
	case Swing:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Sync
	case Sync:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TimeModification
	case TimeModification:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Timpani
	case Timpani:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of VirtualInstrument
	case VirtualInstrument:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Wedge
	case Wedge:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Wood
	case Wood:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var t Type
	return t.GongNewInstance().(Type)
}

func NewInstance[Type GongstructPtr]() (res Type) {
	return GongNewInstance[Type]()
}

func (stage *Stage) GongNewInstance[Type GongstructPtr]() (res Type) {
	res = GongNewInstance[Type]()
	var zero Type
	if res != zero {
		res.StageVoid(stage)
	}
	return res
}

func (stage *Stage) NewInstance[Type GongstructPtr]() (res Type) {
	return stage.GongNewInstance[Type]()
}

// GongGetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GongGetPointerToGongstructName[Type GongstructIF]() (res string) {
	var t Type
	if any(t) == nil {
		return ""
	}
	return t.GongGetGongstructName()
}

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

// insertion point for generic get reverse fields
func (*Accord) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Scordatura",
			Fieldname: "Accord",
		},
	}
}

func (*AccordionRegistration) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Barre) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Bass) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Harmony",
			Fieldname: "Bass",
		},
	}
}

func (*BassStep) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*BeatUnitTied) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Metronome",
			Fieldname: "BeatUnitTied",
		},
	}
}

func (*Beater) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Bracket) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Dashes) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Degree) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Harmony",
			Fieldname: "Degree",
		},
	}
}

func (*DegreeAlter) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DegreeType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DegreeValue) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Direction) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*DirectionType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Direction",
			Fieldname: "DirectionType",
		},
	}
}

func (*Effect) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Feature) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Grouping",
			Fieldname: "Feature",
		},
	}
}

func (*FirstFret) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Frame) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*FrameNote) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Frame",
			Fieldname: "FrameNote",
		},
	}
}

func (*Glass) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Grouping) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Harmony) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*HarmonyAlter) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*HarpPedals) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Image) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*InstrumentChange) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Sound",
			Fieldname: "InstrumentChange",
		},
	}
}

func (*Inversion) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Harmony",
			Fieldname: "Inversion",
		},
	}
}

func (*Kind) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Harmony",
			Fieldname: "Kind",
		},
	}
}

func (*Listening) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MeasureNumbering) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Membrane) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Metal) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Metronome) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MetronomeBeam) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "MetronomeNote",
			Fieldname: "MetronomeBeam",
		},
	}
}

func (*MetronomeNote) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Metronome",
			Fieldname: "MetronomeNote",
		},
	}
}

func (*MetronomeTied) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MetronomeTuplet) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Numeral) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Harmony",
			Fieldname: "Numeral",
		},
	}
}

func (*NumeralKey) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*NumeralRoot) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*OctaveShift) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Offset) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*OtherDirection) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*OtherListening) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Listening",
			Fieldname: "OtherListening",
		},
	}
}

func (*Pedal) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PedalTuning) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "HarpPedals",
			Fieldname: "PedalTuning",
		},
	}
}

func (*PerMinute) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Percussion) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "DirectionType",
			Fieldname: "Percussion",
		},
	}
}

func (*Pitched) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PrincipalVoice) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Print) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Root) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Harmony",
			Fieldname: "Root",
		},
	}
}

func (*RootStep) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Scordatura) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Sound) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*StaffDivide) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Stick) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*StringMute) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*StyleText) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Harmony",
			Fieldname: "Function",
		},
	}
}

func (*Swing) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Sync) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Listening",
			Fieldname: "Sync",
		},
	}
}

func (*TimeModification) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Timpani) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*VirtualInstrument) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Wedge) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Wood) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}


func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	var t Type
	if any(t) == nil {
		return nil
	}
	return t.GongGetReverseFields()
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (accord *Accord) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "String",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "TuningAlter",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TuningOctave",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (accordionregistration *AccordionRegistration) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "AccordionMiddle",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (barre *Barre) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (bass *Bass) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Arrangement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumHarmonyArrangement",
		},
		{
			Name:                 "BassSeparator",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StyleText",
		},
		{
			Name:                 "BassStep",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "BassStep",
		},
		{
			Name:                 "BassAlter",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "HarmonyAlter",
		},
	}
	return
}

func (bassstep *BassStep) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (beatunittied *BeatUnitTied) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (beater *Beater) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Tip",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTipDirection",
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumBeaterValue",
		},
	}
	return
}

func (bracket *Bracket) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "LineEnd",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLineEnd",
		},
		{
			Name:               "EndLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DashLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SpaceLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (dashes *Dashes) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "DashLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SpaceLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (degree *Degree) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "DegreeValue",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DegreeValue",
		},
		{
			Name:                 "DegreeAlter",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DegreeAlter",
		},
		{
			Name:                 "DegreeType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DegreeType",
		},
	}
	return
}

func (degreealter *DegreeAlter) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (degreetype *DegreeType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumDegreeTypeValue",
		},
	}
	return
}

func (degreevalue *DegreeValue) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Symbol",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumDegreeSymbolValue",
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (direction *Direction) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "DirectionType",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DirectionType",
		},
		{
			Name:                 "Offset",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Offset",
		},
		{
			Name:               "Voice",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Staff",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Sound",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Sound",
		},
		{
			Name:                 "Listening",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Listening",
		},
	}
	return
}

func (directiontype *DirectionType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Wedge",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Wedge",
		},
		{
			Name:                 "Dashes",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Dashes",
		},
		{
			Name:                 "Bracket",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bracket",
		},
		{
			Name:                 "Pedal",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Pedal",
		},
		{
			Name:                 "Metronome",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Metronome",
		},
		{
			Name:                 "OctaveShift",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "OctaveShift",
		},
		{
			Name:                 "HarpPedals",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "HarpPedals",
		},
		{
			Name:                 "StringMute",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StringMute",
		},
		{
			Name:                 "Scordatura",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Scordatura",
		},
		{
			Name:                 "Image",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Image",
		},
		{
			Name:                 "PrincipalVoice",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PrincipalVoice",
		},
		{
			Name:                 "Percussion",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Percussion",
		},
		{
			Name:                 "AccordionRegistration",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "AccordionRegistration",
		},
		{
			Name:                 "StaffDivide",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "StaffDivide",
		},
		{
			Name:                 "OtherDirection",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "OtherDirection",
		},
	}
	return
}

func (effect *Effect) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumEffectValue",
		},
	}
	return
}

func (feature *Feature) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (firstfret *FirstFret) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (frame *Frame) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Unplayed",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FrameStrings",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "FrameFrets",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "FirstFret",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "FirstFret",
		},
		{
			Name:                 "FrameNote",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "FrameNote",
		},
	}
	return
}

func (framenote *FrameNote) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Barre",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Barre",
		},
	}
	return
}

func (glass *Glass) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumGlassValue",
		},
	}
	return
}

func (grouping *Grouping) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MemberOf",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Feature",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Feature",
		},
	}
	return
}

func (harmony *Harmony) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumHarmonyType",
		},
		{
			Name:                 "Arrangement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumHarmonyArrangement",
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Root",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Root",
		},
		{
			Name:                 "Numeral",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Numeral",
		},
		{
			Name:                 "Function",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StyleText",
		},
		{
			Name:                 "Kind",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Kind",
		},
		{
			Name:                 "Inversion",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Inversion",
		},
		{
			Name:                 "Bass",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Bass",
		},
		{
			Name:                 "Degree",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Degree",
		},
		{
			Name:                 "Frame",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Frame",
		},
		{
			Name:                 "Offset",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Offset",
		},
		{
			Name:               "Staff",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (harmonyalter *HarmonyAlter) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (harppedals *HarpPedals) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PedalTuning",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PedalTuning",
		},
	}
	return
}

func (image *Image) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Source",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (instrumentchange *InstrumentChange) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InstrumentSound",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Ensemble",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "VirtualInstrument",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "VirtualInstrument",
		},
	}
	return
}

func (inversion *Inversion) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (kind *Kind) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumKindValue",
		},
	}
	return
}

func (listening *Listening) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sync",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Sync",
		},
		{
			Name:                 "OtherListening",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OtherListening",
		},
		{
			Name:                 "Offset",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Offset",
		},
	}
	return
}

func (measurenumbering *MeasureNumbering) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Staff",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumMeasureNumberingValue",
		},
	}
	return
}

func (membrane *Membrane) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumMembraneValue",
		},
	}
	return
}

func (metal *Metal) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumMetalValue",
		},
	}
	return
}

func (metronome *Metronome) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "BeatUnitTied",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BeatUnitTied",
		},
		{
			Name:                 "PerMinute",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PerMinute",
		},
		{
			Name:                 "MetronomeNote",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MetronomeNote",
		},
		{
			Name:               "MetronomeRelation",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (metronomebeam *MetronomeBeam) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumBeamValue",
		},
	}
	return
}

func (metronomenote *MetronomeNote) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "MetronomeBeam",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MetronomeBeam",
		},
		{
			Name:                 "MetronomeTied",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MetronomeTied",
		},
		{
			Name:                 "MetronomeTuplet",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MetronomeTuplet",
		},
	}
	return
}

func (metronometied *MetronomeTied) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (metronometuplet *MetronomeTuplet) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ShowNumber",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumShowTuplet",
		},
		{
			Name:               "ActualNotes",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NormalNotes",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (numeral *Numeral) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "NumeralRoot",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "NumeralRoot",
		},
		{
			Name:                 "NumeralAlter",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "HarmonyAlter",
		},
		{
			Name:                 "NumeralKey",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "NumeralKey",
		},
	}
	return
}

func (numeralkey *NumeralKey) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NumeralFifths",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "NumeralMode",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumNumeralMode",
		},
	}
	return
}

func (numeralroot *NumeralRoot) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (octaveshift *OctaveShift) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumUpDownStopContinue",
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Size",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "DashLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SpaceLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (offset *Offset) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (otherdirection *OtherDirection) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (otherlistening *OtherListening) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Player",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeOnly",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (pedal *Pedal) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumPedalType",
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (pedaltuning *PedalTuning) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PedalAlter",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (perminute *PerMinute) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (percussion *Percussion) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Glass",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Glass",
		},
		{
			Name:                 "Metal",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Metal",
		},
		{
			Name:                 "Wood",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Wood",
		},
		{
			Name:                 "Pitched",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Pitched",
		},
		{
			Name:                 "Membrane",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Membrane",
		},
		{
			Name:                 "Effect",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Effect",
		},
		{
			Name:                 "Timpani",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Timpani",
		},
		{
			Name:                 "Beater",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Beater",
		},
		{
			Name:                 "Stick",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Stick",
		},
		{
			Name:                 "StickLocation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStickLocation",
		},
	}
	return
}

func (pitched *Pitched) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumPitchedValue",
		},
	}
	return
}

func (principalvoice *PrincipalVoice) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Symbol",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumPrincipalVoiceSymbol",
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (print *Print) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StaffSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BlankPage",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "PageNumber",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "MeasureNumbering",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MeasureNumbering",
		},
	}
	return
}

func (root *Root) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "RootStep",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RootStep",
		},
		{
			Name:                 "RootAlter",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "HarmonyAlter",
		},
	}
	return
}

func (rootstep *RootStep) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (scordatura *Scordatura) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Accord",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Accord",
		},
	}
	return
}

func (sound *Sound) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Tempo",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Dynamics",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Segno",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Dalsegno",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Coda",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Tocoda",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Divisions",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Fine",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeOnly",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Pan",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Elevation",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DamperPedal",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SoftPedal",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SostenutoPedal",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "InstrumentChange",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "InstrumentChange",
		},
		{
			Name:                 "Swing",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Swing",
		},
		{
			Name:                 "Offset",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Offset",
		},
	}
	return
}

func (staffdivide *StaffDivide) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStaffDivideSymbol",
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (stick *Stick) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Tip",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTipDirection",
		},
		{
			Name:                 "StickType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStickType",
		},
		{
			Name:                 "StickMaterial",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStickMaterial",
		},
	}
	return
}

func (stringmute *StringMute) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumOnOff",
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (styletext *StyleText) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontFamily",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (swing *Swing) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "First",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Second",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "SwingStyle",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (sync *Sync) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumSyncType",
		},
		{
			Name:               "Latency",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Player",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeOnly",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (timemodification *TimeModification) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ActualNotes",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NormalNotes",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (timpani *Timpani) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (virtualinstrument *VirtualInstrument) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "VirtualLibrary",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "VirtualName",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (wedge *Wedge) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumWedgeType",
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Spread",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DashLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SpaceLength",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DefaultY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RelativeY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (wood *Wood) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumWoodValue",
		},
	}
	return
}

// GongGetFieldsFromPointer return the array of the fields
func GongGetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
}

func GetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	return GongGetFieldsFromPointer[Type]()
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt             GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeIntDuration     GongFieldValueType = "GongFieldValueTypeIntDuration"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeDate            GongFieldValueType = "GongFieldValueTypeDate"
	GongFieldValueTypeBasicKind       GongFieldValueType = "GongFieldValueTypeBasicKind"
	GongFieldValueTypePointer         GongFieldValueType = "GongFieldValueTypePointer"
	GongFieldValueTypeSliceOfPointers GongFieldValueType = "GongFieldValueTypeSliceOfPointers"
)

type GongFieldValue struct {
	GongFieldValueType
	valueString string
	valueInt    int
	valueFloat  float64
	valueBool   bool

	// in case of a pointer, the ID of the pointed element
	// in case of a slice of pointers, the IDs, separated by semi columbs
	ids string
}

type GongFieldHeader struct {
	Name string
	GongFieldValueType
	TargetGongstructName string
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}

func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}

func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

// insertion point for generic get gongstruct field value
func (accord *Accord) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = accord.Name
	case "String":
		res.valueString = fmt.Sprintf("%d", accord.String)
		res.valueInt = accord.String
		res.GongFieldValueType = GongFieldValueTypeInt
	case "TuningAlter":
		res.valueString = accord.TuningAlter
	case "TuningOctave":
		res.valueString = fmt.Sprintf("%d", accord.TuningOctave)
		res.valueInt = accord.TuningOctave
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (accordionregistration *AccordionRegistration) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = accordionregistration.Name
	case "DefaultX":
		res.valueString = accordionregistration.DefaultX
	case "DefaultY":
		res.valueString = accordionregistration.DefaultY
	case "RelativeX":
		res.valueString = accordionregistration.RelativeX
	case "RelativeY":
		res.valueString = accordionregistration.RelativeY
	case "FontFamily":
		res.valueString = accordionregistration.FontFamily
	case "FontSize":
		res.valueString = accordionregistration.FontSize
	case "Color":
		res.valueString = accordionregistration.Color
	case "Id":
		res.valueString = accordionregistration.Id
	case "AccordionMiddle":
		res.valueString = fmt.Sprintf("%d", accordionregistration.AccordionMiddle)
		res.valueInt = accordionregistration.AccordionMiddle
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (barre *Barre) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = barre.Name
	case "Color":
		res.valueString = barre.Color
	}
	return
}

func (bass *Bass) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bass.Name
	case "Arrangement":
		enum := bass.Arrangement
		res.valueString = enum.ToCodeString()
	case "BassSeparator":
		res.GongFieldValueType = GongFieldValueTypePointer
		if bass.BassSeparator != nil {
			res.valueString = bass.BassSeparator.Name
			res.ids = bass.BassSeparator.GongGetUUID(stage)
		}
	case "BassStep":
		res.GongFieldValueType = GongFieldValueTypePointer
		if bass.BassStep != nil {
			res.valueString = bass.BassStep.Name
			res.ids = bass.BassStep.GongGetUUID(stage)
		}
	case "BassAlter":
		res.GongFieldValueType = GongFieldValueTypePointer
		if bass.BassAlter != nil {
			res.valueString = bass.BassAlter.Name
			res.ids = bass.BassAlter.GongGetUUID(stage)
		}
	}
	return
}

func (bassstep *BassStep) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bassstep.Name
	case "Text":
		res.valueString = bassstep.Text
	case "DefaultX":
		res.valueString = bassstep.DefaultX
	case "DefaultY":
		res.valueString = bassstep.DefaultY
	case "RelativeX":
		res.valueString = bassstep.RelativeX
	case "RelativeY":
		res.valueString = bassstep.RelativeY
	case "FontFamily":
		res.valueString = bassstep.FontFamily
	case "FontSize":
		res.valueString = bassstep.FontSize
	case "Color":
		res.valueString = bassstep.Color
	}
	return
}

func (beatunittied *BeatUnitTied) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = beatunittied.Name
	}
	return
}

func (beater *Beater) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = beater.Name
	case "Tip":
		enum := beater.Tip
		res.valueString = enum.ToCodeString()
	case "Value":
		enum := beater.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (bracket *Bracket) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bracket.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", bracket.Number)
		res.valueInt = bracket.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "LineEnd":
		enum := bracket.LineEnd
		res.valueString = enum.ToCodeString()
	case "EndLength":
		res.valueString = bracket.EndLength
	case "DashLength":
		res.valueString = bracket.DashLength
	case "SpaceLength":
		res.valueString = bracket.SpaceLength
	case "DefaultX":
		res.valueString = bracket.DefaultX
	case "DefaultY":
		res.valueString = bracket.DefaultY
	case "RelativeX":
		res.valueString = bracket.RelativeX
	case "RelativeY":
		res.valueString = bracket.RelativeY
	case "Color":
		res.valueString = bracket.Color
	case "Id":
		res.valueString = bracket.Id
	}
	return
}

func (dashes *Dashes) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dashes.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", dashes.Number)
		res.valueInt = dashes.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DashLength":
		res.valueString = dashes.DashLength
	case "SpaceLength":
		res.valueString = dashes.SpaceLength
	case "DefaultX":
		res.valueString = dashes.DefaultX
	case "DefaultY":
		res.valueString = dashes.DefaultY
	case "RelativeX":
		res.valueString = dashes.RelativeX
	case "RelativeY":
		res.valueString = dashes.RelativeY
	case "Color":
		res.valueString = dashes.Color
	case "Id":
		res.valueString = dashes.Id
	}
	return
}

func (degree *Degree) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = degree.Name
	case "DegreeValue":
		res.GongFieldValueType = GongFieldValueTypePointer
		if degree.DegreeValue != nil {
			res.valueString = degree.DegreeValue.Name
			res.ids = degree.DegreeValue.GongGetUUID(stage)
		}
	case "DegreeAlter":
		res.GongFieldValueType = GongFieldValueTypePointer
		if degree.DegreeAlter != nil {
			res.valueString = degree.DegreeAlter.Name
			res.ids = degree.DegreeAlter.GongGetUUID(stage)
		}
	case "DegreeType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if degree.DegreeType != nil {
			res.valueString = degree.DegreeType.Name
			res.ids = degree.DegreeType.GongGetUUID(stage)
		}
	}
	return
}

func (degreealter *DegreeAlter) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = degreealter.Name
	case "DefaultX":
		res.valueString = degreealter.DefaultX
	case "DefaultY":
		res.valueString = degreealter.DefaultY
	case "RelativeX":
		res.valueString = degreealter.RelativeX
	case "RelativeY":
		res.valueString = degreealter.RelativeY
	case "FontFamily":
		res.valueString = degreealter.FontFamily
	case "FontSize":
		res.valueString = degreealter.FontSize
	case "Color":
		res.valueString = degreealter.Color
	case "Value":
		res.valueString = degreealter.Value
	}
	return
}

func (degreetype *DegreeType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = degreetype.Name
	case "Text":
		res.valueString = degreetype.Text
	case "DefaultX":
		res.valueString = degreetype.DefaultX
	case "DefaultY":
		res.valueString = degreetype.DefaultY
	case "RelativeX":
		res.valueString = degreetype.RelativeX
	case "RelativeY":
		res.valueString = degreetype.RelativeY
	case "FontFamily":
		res.valueString = degreetype.FontFamily
	case "FontSize":
		res.valueString = degreetype.FontSize
	case "Color":
		res.valueString = degreetype.Color
	case "Value":
		enum := degreetype.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (degreevalue *DegreeValue) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = degreevalue.Name
	case "Symbol":
		enum := degreevalue.Symbol
		res.valueString = enum.ToCodeString()
	case "Text":
		res.valueString = degreevalue.Text
	case "DefaultX":
		res.valueString = degreevalue.DefaultX
	case "DefaultY":
		res.valueString = degreevalue.DefaultY
	case "RelativeX":
		res.valueString = degreevalue.RelativeX
	case "RelativeY":
		res.valueString = degreevalue.RelativeY
	case "FontFamily":
		res.valueString = degreevalue.FontFamily
	case "FontSize":
		res.valueString = degreevalue.FontSize
	case "Color":
		res.valueString = degreevalue.Color
	case "Value":
		res.valueString = fmt.Sprintf("%d", degreevalue.Value)
		res.valueInt = degreevalue.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (direction *Direction) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = direction.Name
	case "Id":
		res.valueString = direction.Id
	case "DirectionType":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range direction.DirectionType {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Offset":
		res.GongFieldValueType = GongFieldValueTypePointer
		if direction.Offset != nil {
			res.valueString = direction.Offset.Name
			res.ids = direction.Offset.GongGetUUID(stage)
		}
	case "Voice":
		res.valueString = direction.Voice
	case "Staff":
		res.valueString = fmt.Sprintf("%d", direction.Staff)
		res.valueInt = direction.Staff
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Sound":
		res.GongFieldValueType = GongFieldValueTypePointer
		if direction.Sound != nil {
			res.valueString = direction.Sound.Name
			res.ids = direction.Sound.GongGetUUID(stage)
		}
	case "Listening":
		res.GongFieldValueType = GongFieldValueTypePointer
		if direction.Listening != nil {
			res.valueString = direction.Listening.Name
			res.ids = direction.Listening.GongGetUUID(stage)
		}
	}
	return
}

func (directiontype *DirectionType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = directiontype.Name
	case "Id":
		res.valueString = directiontype.Id
	case "Wedge":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.Wedge != nil {
			res.valueString = directiontype.Wedge.Name
			res.ids = directiontype.Wedge.GongGetUUID(stage)
		}
	case "Dashes":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.Dashes != nil {
			res.valueString = directiontype.Dashes.Name
			res.ids = directiontype.Dashes.GongGetUUID(stage)
		}
	case "Bracket":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.Bracket != nil {
			res.valueString = directiontype.Bracket.Name
			res.ids = directiontype.Bracket.GongGetUUID(stage)
		}
	case "Pedal":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.Pedal != nil {
			res.valueString = directiontype.Pedal.Name
			res.ids = directiontype.Pedal.GongGetUUID(stage)
		}
	case "Metronome":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.Metronome != nil {
			res.valueString = directiontype.Metronome.Name
			res.ids = directiontype.Metronome.GongGetUUID(stage)
		}
	case "OctaveShift":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.OctaveShift != nil {
			res.valueString = directiontype.OctaveShift.Name
			res.ids = directiontype.OctaveShift.GongGetUUID(stage)
		}
	case "HarpPedals":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.HarpPedals != nil {
			res.valueString = directiontype.HarpPedals.Name
			res.ids = directiontype.HarpPedals.GongGetUUID(stage)
		}
	case "StringMute":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.StringMute != nil {
			res.valueString = directiontype.StringMute.Name
			res.ids = directiontype.StringMute.GongGetUUID(stage)
		}
	case "Scordatura":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.Scordatura != nil {
			res.valueString = directiontype.Scordatura.Name
			res.ids = directiontype.Scordatura.GongGetUUID(stage)
		}
	case "Image":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.Image != nil {
			res.valueString = directiontype.Image.Name
			res.ids = directiontype.Image.GongGetUUID(stage)
		}
	case "PrincipalVoice":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.PrincipalVoice != nil {
			res.valueString = directiontype.PrincipalVoice.Name
			res.ids = directiontype.PrincipalVoice.GongGetUUID(stage)
		}
	case "Percussion":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range directiontype.Percussion {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AccordionRegistration":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.AccordionRegistration != nil {
			res.valueString = directiontype.AccordionRegistration.Name
			res.ids = directiontype.AccordionRegistration.GongGetUUID(stage)
		}
	case "StaffDivide":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.StaffDivide != nil {
			res.valueString = directiontype.StaffDivide.Name
			res.ids = directiontype.StaffDivide.GongGetUUID(stage)
		}
	case "OtherDirection":
		res.GongFieldValueType = GongFieldValueTypePointer
		if directiontype.OtherDirection != nil {
			res.valueString = directiontype.OtherDirection.Name
			res.ids = directiontype.OtherDirection.GongGetUUID(stage)
		}
	}
	return
}

func (effect *Effect) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = effect.Name
	case "Smufl":
		res.valueString = effect.Smufl
	case "Value":
		enum := effect.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (feature *Feature) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = feature.Name
	case "Type":
		res.valueString = feature.Type
	case "Value":
		res.valueString = feature.Value
	}
	return
}

func (firstfret *FirstFret) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = firstfret.Name
	case "Text":
		res.valueString = firstfret.Text
	case "Value":
		res.valueString = fmt.Sprintf("%d", firstfret.Value)
		res.valueInt = firstfret.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (frame *Frame) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = frame.Name
	case "Height":
		res.valueString = frame.Height
	case "Width":
		res.valueString = frame.Width
	case "Unplayed":
		res.valueString = frame.Unplayed
	case "DefaultX":
		res.valueString = frame.DefaultX
	case "DefaultY":
		res.valueString = frame.DefaultY
	case "RelativeX":
		res.valueString = frame.RelativeX
	case "RelativeY":
		res.valueString = frame.RelativeY
	case "Color":
		res.valueString = frame.Color
	case "Id":
		res.valueString = frame.Id
	case "FrameStrings":
		res.valueString = fmt.Sprintf("%d", frame.FrameStrings)
		res.valueInt = frame.FrameStrings
		res.GongFieldValueType = GongFieldValueTypeInt
	case "FrameFrets":
		res.valueString = fmt.Sprintf("%d", frame.FrameFrets)
		res.valueInt = frame.FrameFrets
		res.GongFieldValueType = GongFieldValueTypeInt
	case "FirstFret":
		res.GongFieldValueType = GongFieldValueTypePointer
		if frame.FirstFret != nil {
			res.valueString = frame.FirstFret.Name
			res.ids = frame.FirstFret.GongGetUUID(stage)
		}
	case "FrameNote":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range frame.FrameNote {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (framenote *FrameNote) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = framenote.Name
	case "Barre":
		res.GongFieldValueType = GongFieldValueTypePointer
		if framenote.Barre != nil {
			res.valueString = framenote.Barre.Name
			res.ids = framenote.Barre.GongGetUUID(stage)
		}
	}
	return
}

func (glass *Glass) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = glass.Name
	case "Smufl":
		res.valueString = glass.Smufl
	case "Value":
		enum := glass.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (grouping *Grouping) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = grouping.Name
	case "Number":
		res.valueString = grouping.Number
	case "MemberOf":
		res.valueString = grouping.MemberOf
	case "Id":
		res.valueString = grouping.Id
	case "Feature":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range grouping.Feature {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (harmony *Harmony) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = harmony.Name
	case "Type":
		enum := harmony.Type
		res.valueString = enum.ToCodeString()
	case "Arrangement":
		enum := harmony.Arrangement
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = harmony.DefaultX
	case "DefaultY":
		res.valueString = harmony.DefaultY
	case "RelativeX":
		res.valueString = harmony.RelativeX
	case "RelativeY":
		res.valueString = harmony.RelativeY
	case "FontFamily":
		res.valueString = harmony.FontFamily
	case "FontSize":
		res.valueString = harmony.FontSize
	case "Color":
		res.valueString = harmony.Color
	case "Id":
		res.valueString = harmony.Id
	case "Root":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range harmony.Root {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Numeral":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range harmony.Numeral {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Function":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range harmony.Function {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Kind":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range harmony.Kind {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Inversion":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range harmony.Inversion {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Bass":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range harmony.Bass {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Degree":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range harmony.Degree {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Frame":
		res.GongFieldValueType = GongFieldValueTypePointer
		if harmony.Frame != nil {
			res.valueString = harmony.Frame.Name
			res.ids = harmony.Frame.GongGetUUID(stage)
		}
	case "Offset":
		res.GongFieldValueType = GongFieldValueTypePointer
		if harmony.Offset != nil {
			res.valueString = harmony.Offset.Name
			res.ids = harmony.Offset.GongGetUUID(stage)
		}
	case "Staff":
		res.valueString = fmt.Sprintf("%d", harmony.Staff)
		res.valueInt = harmony.Staff
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (harmonyalter *HarmonyAlter) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = harmonyalter.Name
	case "DefaultX":
		res.valueString = harmonyalter.DefaultX
	case "DefaultY":
		res.valueString = harmonyalter.DefaultY
	case "RelativeX":
		res.valueString = harmonyalter.RelativeX
	case "RelativeY":
		res.valueString = harmonyalter.RelativeY
	case "FontFamily":
		res.valueString = harmonyalter.FontFamily
	case "FontSize":
		res.valueString = harmonyalter.FontSize
	case "Color":
		res.valueString = harmonyalter.Color
	case "Value":
		res.valueString = harmonyalter.Value
	}
	return
}

func (harppedals *HarpPedals) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = harppedals.Name
	case "DefaultX":
		res.valueString = harppedals.DefaultX
	case "DefaultY":
		res.valueString = harppedals.DefaultY
	case "RelativeX":
		res.valueString = harppedals.RelativeX
	case "RelativeY":
		res.valueString = harppedals.RelativeY
	case "FontFamily":
		res.valueString = harppedals.FontFamily
	case "FontSize":
		res.valueString = harppedals.FontSize
	case "Color":
		res.valueString = harppedals.Color
	case "Id":
		res.valueString = harppedals.Id
	case "PedalTuning":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range harppedals.PedalTuning {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (image *Image) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = image.Name
	case "Source":
		res.valueString = image.Source
	case "Type":
		res.valueString = image.Type
	case "Height":
		res.valueString = image.Height
	case "Width":
		res.valueString = image.Width
	case "DefaultX":
		res.valueString = image.DefaultX
	case "DefaultY":
		res.valueString = image.DefaultY
	case "RelativeX":
		res.valueString = image.RelativeX
	case "RelativeY":
		res.valueString = image.RelativeY
	case "Id":
		res.valueString = image.Id
	}
	return
}

func (instrumentchange *InstrumentChange) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = instrumentchange.Name
	case "Id":
		res.valueString = instrumentchange.Id
	case "InstrumentSound":
		res.valueString = instrumentchange.InstrumentSound
	case "Ensemble":
		res.valueString = instrumentchange.Ensemble
	case "VirtualInstrument":
		res.GongFieldValueType = GongFieldValueTypePointer
		if instrumentchange.VirtualInstrument != nil {
			res.valueString = instrumentchange.VirtualInstrument.Name
			res.ids = instrumentchange.VirtualInstrument.GongGetUUID(stage)
		}
	}
	return
}

func (inversion *Inversion) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = inversion.Name
	case "Text":
		res.valueString = inversion.Text
	case "DefaultX":
		res.valueString = inversion.DefaultX
	case "DefaultY":
		res.valueString = inversion.DefaultY
	case "RelativeX":
		res.valueString = inversion.RelativeX
	case "RelativeY":
		res.valueString = inversion.RelativeY
	case "FontFamily":
		res.valueString = inversion.FontFamily
	case "FontSize":
		res.valueString = inversion.FontSize
	case "Color":
		res.valueString = inversion.Color
	case "Value":
		res.valueString = fmt.Sprintf("%d", inversion.Value)
		res.valueInt = inversion.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (kind *Kind) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = kind.Name
	case "Text":
		res.valueString = kind.Text
	case "DefaultX":
		res.valueString = kind.DefaultX
	case "DefaultY":
		res.valueString = kind.DefaultY
	case "RelativeX":
		res.valueString = kind.RelativeX
	case "RelativeY":
		res.valueString = kind.RelativeY
	case "FontFamily":
		res.valueString = kind.FontFamily
	case "FontSize":
		res.valueString = kind.FontSize
	case "Color":
		res.valueString = kind.Color
	case "Value":
		enum := kind.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (listening *Listening) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = listening.Name
	case "Sync":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range listening.Sync {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "OtherListening":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range listening.OtherListening {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Offset":
		res.GongFieldValueType = GongFieldValueTypePointer
		if listening.Offset != nil {
			res.valueString = listening.Offset.Name
			res.ids = listening.Offset.GongGetUUID(stage)
		}
	}
	return
}

func (measurenumbering *MeasureNumbering) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = measurenumbering.Name
	case "Staff":
		res.valueString = fmt.Sprintf("%d", measurenumbering.Staff)
		res.valueInt = measurenumbering.Staff
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = measurenumbering.DefaultX
	case "DefaultY":
		res.valueString = measurenumbering.DefaultY
	case "RelativeX":
		res.valueString = measurenumbering.RelativeX
	case "RelativeY":
		res.valueString = measurenumbering.RelativeY
	case "FontFamily":
		res.valueString = measurenumbering.FontFamily
	case "FontSize":
		res.valueString = measurenumbering.FontSize
	case "Color":
		res.valueString = measurenumbering.Color
	case "Value":
		enum := measurenumbering.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (membrane *Membrane) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = membrane.Name
	case "Smufl":
		res.valueString = membrane.Smufl
	case "Value":
		enum := membrane.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (metal *Metal) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = metal.Name
	case "Smufl":
		res.valueString = metal.Smufl
	case "Value":
		enum := metal.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (metronome *Metronome) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = metronome.Name
	case "DefaultX":
		res.valueString = metronome.DefaultX
	case "DefaultY":
		res.valueString = metronome.DefaultY
	case "RelativeX":
		res.valueString = metronome.RelativeX
	case "RelativeY":
		res.valueString = metronome.RelativeY
	case "FontFamily":
		res.valueString = metronome.FontFamily
	case "FontSize":
		res.valueString = metronome.FontSize
	case "Color":
		res.valueString = metronome.Color
	case "Id":
		res.valueString = metronome.Id
	case "BeatUnitTied":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range metronome.BeatUnitTied {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PerMinute":
		res.GongFieldValueType = GongFieldValueTypePointer
		if metronome.PerMinute != nil {
			res.valueString = metronome.PerMinute.Name
			res.ids = metronome.PerMinute.GongGetUUID(stage)
		}
	case "MetronomeNote":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range metronome.MetronomeNote {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "MetronomeRelation":
		res.valueString = metronome.MetronomeRelation
	}
	return
}

func (metronomebeam *MetronomeBeam) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = metronomebeam.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", metronomebeam.Number)
		res.valueInt = metronomebeam.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Value":
		enum := metronomebeam.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (metronomenote *MetronomeNote) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = metronomenote.Name
	case "MetronomeBeam":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range metronomenote.MetronomeBeam {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "MetronomeTied":
		res.GongFieldValueType = GongFieldValueTypePointer
		if metronomenote.MetronomeTied != nil {
			res.valueString = metronomenote.MetronomeTied.Name
			res.ids = metronomenote.MetronomeTied.GongGetUUID(stage)
		}
	case "MetronomeTuplet":
		res.GongFieldValueType = GongFieldValueTypePointer
		if metronomenote.MetronomeTuplet != nil {
			res.valueString = metronomenote.MetronomeTuplet.Name
			res.ids = metronomenote.MetronomeTuplet.GongGetUUID(stage)
		}
	}
	return
}

func (metronometied *MetronomeTied) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = metronometied.Name
	}
	return
}

func (metronometuplet *MetronomeTuplet) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = metronometuplet.Name
	case "ShowNumber":
		enum := metronometuplet.ShowNumber
		res.valueString = enum.ToCodeString()
	case "ActualNotes":
		res.valueString = fmt.Sprintf("%d", metronometuplet.ActualNotes)
		res.valueInt = metronometuplet.ActualNotes
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NormalNotes":
		res.valueString = fmt.Sprintf("%d", metronometuplet.NormalNotes)
		res.valueInt = metronometuplet.NormalNotes
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (numeral *Numeral) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = numeral.Name
	case "NumeralRoot":
		res.GongFieldValueType = GongFieldValueTypePointer
		if numeral.NumeralRoot != nil {
			res.valueString = numeral.NumeralRoot.Name
			res.ids = numeral.NumeralRoot.GongGetUUID(stage)
		}
	case "NumeralAlter":
		res.GongFieldValueType = GongFieldValueTypePointer
		if numeral.NumeralAlter != nil {
			res.valueString = numeral.NumeralAlter.Name
			res.ids = numeral.NumeralAlter.GongGetUUID(stage)
		}
	case "NumeralKey":
		res.GongFieldValueType = GongFieldValueTypePointer
		if numeral.NumeralKey != nil {
			res.valueString = numeral.NumeralKey.Name
			res.ids = numeral.NumeralKey.GongGetUUID(stage)
		}
	}
	return
}

func (numeralkey *NumeralKey) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = numeralkey.Name
	case "NumeralFifths":
		res.valueString = fmt.Sprintf("%d", numeralkey.NumeralFifths)
		res.valueInt = numeralkey.NumeralFifths
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NumeralMode":
		enum := numeralkey.NumeralMode
		res.valueString = enum.ToCodeString()
	}
	return
}

func (numeralroot *NumeralRoot) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = numeralroot.Name
	case "Text":
		res.valueString = numeralroot.Text
	case "DefaultX":
		res.valueString = numeralroot.DefaultX
	case "DefaultY":
		res.valueString = numeralroot.DefaultY
	case "RelativeX":
		res.valueString = numeralroot.RelativeX
	case "RelativeY":
		res.valueString = numeralroot.RelativeY
	case "FontFamily":
		res.valueString = numeralroot.FontFamily
	case "FontSize":
		res.valueString = numeralroot.FontSize
	case "Color":
		res.valueString = numeralroot.Color
	case "Value":
		res.valueString = fmt.Sprintf("%d", numeralroot.Value)
		res.valueInt = numeralroot.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (octaveshift *OctaveShift) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = octaveshift.Name
	case "Type":
		enum := octaveshift.Type
		res.valueString = enum.ToCodeString()
	case "Number":
		res.valueString = fmt.Sprintf("%d", octaveshift.Number)
		res.valueInt = octaveshift.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Size":
		res.valueString = fmt.Sprintf("%d", octaveshift.Size)
		res.valueInt = octaveshift.Size
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DashLength":
		res.valueString = octaveshift.DashLength
	case "SpaceLength":
		res.valueString = octaveshift.SpaceLength
	case "DefaultX":
		res.valueString = octaveshift.DefaultX
	case "DefaultY":
		res.valueString = octaveshift.DefaultY
	case "RelativeX":
		res.valueString = octaveshift.RelativeX
	case "RelativeY":
		res.valueString = octaveshift.RelativeY
	case "FontFamily":
		res.valueString = octaveshift.FontFamily
	case "FontSize":
		res.valueString = octaveshift.FontSize
	case "Color":
		res.valueString = octaveshift.Color
	case "Id":
		res.valueString = octaveshift.Id
	}
	return
}

func (offset *Offset) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = offset.Name
	case "Value":
		res.valueString = offset.Value
	}
	return
}

func (otherdirection *OtherDirection) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = otherdirection.Name
	case "DefaultX":
		res.valueString = otherdirection.DefaultX
	case "DefaultY":
		res.valueString = otherdirection.DefaultY
	case "RelativeX":
		res.valueString = otherdirection.RelativeX
	case "RelativeY":
		res.valueString = otherdirection.RelativeY
	case "FontFamily":
		res.valueString = otherdirection.FontFamily
	case "FontSize":
		res.valueString = otherdirection.FontSize
	case "Color":
		res.valueString = otherdirection.Color
	case "Smufl":
		res.valueString = otherdirection.Smufl
	case "Id":
		res.valueString = otherdirection.Id
	case "Value":
		res.valueString = otherdirection.Value
	}
	return
}

func (otherlistening *OtherListening) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = otherlistening.Name
	case "Type":
		res.valueString = otherlistening.Type
	case "Player":
		res.valueString = otherlistening.Player
	case "TimeOnly":
		res.valueString = otherlistening.TimeOnly
	case "Value":
		res.valueString = otherlistening.Value
	}
	return
}

func (pedal *Pedal) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pedal.Name
	case "Type":
		enum := pedal.Type
		res.valueString = enum.ToCodeString()
	case "Number":
		res.valueString = fmt.Sprintf("%d", pedal.Number)
		res.valueInt = pedal.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = pedal.DefaultX
	case "DefaultY":
		res.valueString = pedal.DefaultY
	case "RelativeX":
		res.valueString = pedal.RelativeX
	case "RelativeY":
		res.valueString = pedal.RelativeY
	case "FontFamily":
		res.valueString = pedal.FontFamily
	case "FontSize":
		res.valueString = pedal.FontSize
	case "Color":
		res.valueString = pedal.Color
	case "Id":
		res.valueString = pedal.Id
	}
	return
}

func (pedaltuning *PedalTuning) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pedaltuning.Name
	case "PedalAlter":
		res.valueString = pedaltuning.PedalAlter
	}
	return
}

func (perminute *PerMinute) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = perminute.Name
	case "FontFamily":
		res.valueString = perminute.FontFamily
	case "FontSize":
		res.valueString = perminute.FontSize
	case "Value":
		res.valueString = perminute.Value
	}
	return
}

func (percussion *Percussion) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = percussion.Name
	case "DefaultX":
		res.valueString = percussion.DefaultX
	case "DefaultY":
		res.valueString = percussion.DefaultY
	case "RelativeX":
		res.valueString = percussion.RelativeX
	case "RelativeY":
		res.valueString = percussion.RelativeY
	case "FontFamily":
		res.valueString = percussion.FontFamily
	case "FontSize":
		res.valueString = percussion.FontSize
	case "Color":
		res.valueString = percussion.Color
	case "Id":
		res.valueString = percussion.Id
	case "Glass":
		res.GongFieldValueType = GongFieldValueTypePointer
		if percussion.Glass != nil {
			res.valueString = percussion.Glass.Name
			res.ids = percussion.Glass.GongGetUUID(stage)
		}
	case "Metal":
		res.GongFieldValueType = GongFieldValueTypePointer
		if percussion.Metal != nil {
			res.valueString = percussion.Metal.Name
			res.ids = percussion.Metal.GongGetUUID(stage)
		}
	case "Wood":
		res.GongFieldValueType = GongFieldValueTypePointer
		if percussion.Wood != nil {
			res.valueString = percussion.Wood.Name
			res.ids = percussion.Wood.GongGetUUID(stage)
		}
	case "Pitched":
		res.GongFieldValueType = GongFieldValueTypePointer
		if percussion.Pitched != nil {
			res.valueString = percussion.Pitched.Name
			res.ids = percussion.Pitched.GongGetUUID(stage)
		}
	case "Membrane":
		res.GongFieldValueType = GongFieldValueTypePointer
		if percussion.Membrane != nil {
			res.valueString = percussion.Membrane.Name
			res.ids = percussion.Membrane.GongGetUUID(stage)
		}
	case "Effect":
		res.GongFieldValueType = GongFieldValueTypePointer
		if percussion.Effect != nil {
			res.valueString = percussion.Effect.Name
			res.ids = percussion.Effect.GongGetUUID(stage)
		}
	case "Timpani":
		res.GongFieldValueType = GongFieldValueTypePointer
		if percussion.Timpani != nil {
			res.valueString = percussion.Timpani.Name
			res.ids = percussion.Timpani.GongGetUUID(stage)
		}
	case "Beater":
		res.GongFieldValueType = GongFieldValueTypePointer
		if percussion.Beater != nil {
			res.valueString = percussion.Beater.Name
			res.ids = percussion.Beater.GongGetUUID(stage)
		}
	case "Stick":
		res.GongFieldValueType = GongFieldValueTypePointer
		if percussion.Stick != nil {
			res.valueString = percussion.Stick.Name
			res.ids = percussion.Stick.GongGetUUID(stage)
		}
	case "StickLocation":
		enum := percussion.StickLocation
		res.valueString = enum.ToCodeString()
	}
	return
}

func (pitched *Pitched) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pitched.Name
	case "Smufl":
		res.valueString = pitched.Smufl
	case "Value":
		enum := pitched.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (principalvoice *PrincipalVoice) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = principalvoice.Name
	case "Symbol":
		enum := principalvoice.Symbol
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = principalvoice.DefaultX
	case "DefaultY":
		res.valueString = principalvoice.DefaultY
	case "RelativeX":
		res.valueString = principalvoice.RelativeX
	case "RelativeY":
		res.valueString = principalvoice.RelativeY
	case "FontFamily":
		res.valueString = principalvoice.FontFamily
	case "FontSize":
		res.valueString = principalvoice.FontSize
	case "Color":
		res.valueString = principalvoice.Color
	case "Id":
		res.valueString = principalvoice.Id
	case "Value":
		res.valueString = principalvoice.Value
	}
	return
}

func (print *Print) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = print.Name
	case "StaffSpacing":
		res.valueString = print.StaffSpacing
	case "BlankPage":
		res.valueString = fmt.Sprintf("%d", print.BlankPage)
		res.valueInt = print.BlankPage
		res.GongFieldValueType = GongFieldValueTypeInt
	case "PageNumber":
		res.valueString = print.PageNumber
	case "Id":
		res.valueString = print.Id
	case "MeasureNumbering":
		res.GongFieldValueType = GongFieldValueTypePointer
		if print.MeasureNumbering != nil {
			res.valueString = print.MeasureNumbering.Name
			res.ids = print.MeasureNumbering.GongGetUUID(stage)
		}
	}
	return
}

func (root *Root) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = root.Name
	case "RootStep":
		res.GongFieldValueType = GongFieldValueTypePointer
		if root.RootStep != nil {
			res.valueString = root.RootStep.Name
			res.ids = root.RootStep.GongGetUUID(stage)
		}
	case "RootAlter":
		res.GongFieldValueType = GongFieldValueTypePointer
		if root.RootAlter != nil {
			res.valueString = root.RootAlter.Name
			res.ids = root.RootAlter.GongGetUUID(stage)
		}
	}
	return
}

func (rootstep *RootStep) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rootstep.Name
	case "Text":
		res.valueString = rootstep.Text
	case "DefaultX":
		res.valueString = rootstep.DefaultX
	case "DefaultY":
		res.valueString = rootstep.DefaultY
	case "RelativeX":
		res.valueString = rootstep.RelativeX
	case "RelativeY":
		res.valueString = rootstep.RelativeY
	case "FontFamily":
		res.valueString = rootstep.FontFamily
	case "FontSize":
		res.valueString = rootstep.FontSize
	case "Color":
		res.valueString = rootstep.Color
	}
	return
}

func (scordatura *Scordatura) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scordatura.Name
	case "Id":
		res.valueString = scordatura.Id
	case "Accord":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range scordatura.Accord {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (sound *Sound) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = sound.Name
	case "Tempo":
		res.valueString = sound.Tempo
	case "Dynamics":
		res.valueString = sound.Dynamics
	case "Segno":
		res.valueString = sound.Segno
	case "Dalsegno":
		res.valueString = sound.Dalsegno
	case "Coda":
		res.valueString = sound.Coda
	case "Tocoda":
		res.valueString = sound.Tocoda
	case "Divisions":
		res.valueString = sound.Divisions
	case "Fine":
		res.valueString = sound.Fine
	case "TimeOnly":
		res.valueString = sound.TimeOnly
	case "Pan":
		res.valueString = sound.Pan
	case "Elevation":
		res.valueString = sound.Elevation
	case "DamperPedal":
		res.valueString = sound.DamperPedal
	case "SoftPedal":
		res.valueString = sound.SoftPedal
	case "SostenutoPedal":
		res.valueString = sound.SostenutoPedal
	case "Id":
		res.valueString = sound.Id
	case "InstrumentChange":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range sound.InstrumentChange {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Swing":
		res.GongFieldValueType = GongFieldValueTypePointer
		if sound.Swing != nil {
			res.valueString = sound.Swing.Name
			res.ids = sound.Swing.GongGetUUID(stage)
		}
	case "Offset":
		res.GongFieldValueType = GongFieldValueTypePointer
		if sound.Offset != nil {
			res.valueString = sound.Offset.Name
			res.ids = sound.Offset.GongGetUUID(stage)
		}
	}
	return
}

func (staffdivide *StaffDivide) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = staffdivide.Name
	case "Type":
		enum := staffdivide.Type
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = staffdivide.DefaultX
	case "DefaultY":
		res.valueString = staffdivide.DefaultY
	case "RelativeX":
		res.valueString = staffdivide.RelativeX
	case "RelativeY":
		res.valueString = staffdivide.RelativeY
	case "FontFamily":
		res.valueString = staffdivide.FontFamily
	case "FontSize":
		res.valueString = staffdivide.FontSize
	case "Color":
		res.valueString = staffdivide.Color
	case "Id":
		res.valueString = staffdivide.Id
	}
	return
}

func (stick *Stick) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stick.Name
	case "Tip":
		enum := stick.Tip
		res.valueString = enum.ToCodeString()
	case "StickType":
		enum := stick.StickType
		res.valueString = enum.ToCodeString()
	case "StickMaterial":
		enum := stick.StickMaterial
		res.valueString = enum.ToCodeString()
	}
	return
}

func (stringmute *StringMute) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stringmute.Name
	case "Type":
		enum := stringmute.Type
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = stringmute.DefaultX
	case "DefaultY":
		res.valueString = stringmute.DefaultY
	case "RelativeX":
		res.valueString = stringmute.RelativeX
	case "RelativeY":
		res.valueString = stringmute.RelativeY
	case "FontFamily":
		res.valueString = stringmute.FontFamily
	case "FontSize":
		res.valueString = stringmute.FontSize
	case "Color":
		res.valueString = stringmute.Color
	case "Id":
		res.valueString = stringmute.Id
	}
	return
}

func (styletext *StyleText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = styletext.Name
	case "DefaultX":
		res.valueString = styletext.DefaultX
	case "DefaultY":
		res.valueString = styletext.DefaultY
	case "RelativeX":
		res.valueString = styletext.RelativeX
	case "RelativeY":
		res.valueString = styletext.RelativeY
	case "FontFamily":
		res.valueString = styletext.FontFamily
	case "FontSize":
		res.valueString = styletext.FontSize
	case "Color":
		res.valueString = styletext.Color
	case "Value":
		res.valueString = styletext.Value
	}
	return
}

func (swing *Swing) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = swing.Name
	case "First":
		res.valueString = fmt.Sprintf("%d", swing.First)
		res.valueInt = swing.First
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Second":
		res.valueString = fmt.Sprintf("%d", swing.Second)
		res.valueInt = swing.Second
		res.GongFieldValueType = GongFieldValueTypeInt
	case "SwingStyle":
		res.valueString = swing.SwingStyle
	}
	return
}

func (sync *Sync) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = sync.Name
	case "Type":
		enum := sync.Type
		res.valueString = enum.ToCodeString()
	case "Latency":
		res.valueString = fmt.Sprintf("%d", sync.Latency)
		res.valueInt = sync.Latency
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Player":
		res.valueString = sync.Player
	case "TimeOnly":
		res.valueString = sync.TimeOnly
	}
	return
}

func (timemodification *TimeModification) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timemodification.Name
	case "ActualNotes":
		res.valueString = fmt.Sprintf("%d", timemodification.ActualNotes)
		res.valueInt = timemodification.ActualNotes
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NormalNotes":
		res.valueString = fmt.Sprintf("%d", timemodification.NormalNotes)
		res.valueInt = timemodification.NormalNotes
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (timpani *Timpani) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = timpani.Name
	case "Smufl":
		res.valueString = timpani.Smufl
	}
	return
}

func (virtualinstrument *VirtualInstrument) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = virtualinstrument.Name
	case "VirtualLibrary":
		res.valueString = virtualinstrument.VirtualLibrary
	case "VirtualName":
		res.valueString = virtualinstrument.VirtualName
	}
	return
}

func (wedge *Wedge) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = wedge.Name
	case "Type":
		enum := wedge.Type
		res.valueString = enum.ToCodeString()
	case "Number":
		res.valueString = fmt.Sprintf("%d", wedge.Number)
		res.valueInt = wedge.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Spread":
		res.valueString = wedge.Spread
	case "DashLength":
		res.valueString = wedge.DashLength
	case "SpaceLength":
		res.valueString = wedge.SpaceLength
	case "DefaultX":
		res.valueString = wedge.DefaultX
	case "DefaultY":
		res.valueString = wedge.DefaultY
	case "RelativeX":
		res.valueString = wedge.RelativeX
	case "RelativeY":
		res.valueString = wedge.RelativeY
	case "Color":
		res.valueString = wedge.Color
	case "Id":
		res.valueString = wedge.Id
	}
	return
}

func (wood *Wood) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = wood.Name
	case "Smufl":
		res.valueString = wood.Smufl
	case "Value":
		enum := wood.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (accord *Accord) GongGetGongstructName() string {
	return "Accord"
}

func (accordionregistration *AccordionRegistration) GongGetGongstructName() string {
	return "AccordionRegistration"
}

func (barre *Barre) GongGetGongstructName() string {
	return "Barre"
}

func (bass *Bass) GongGetGongstructName() string {
	return "Bass"
}

func (bassstep *BassStep) GongGetGongstructName() string {
	return "BassStep"
}

func (beatunittied *BeatUnitTied) GongGetGongstructName() string {
	return "BeatUnitTied"
}

func (beater *Beater) GongGetGongstructName() string {
	return "Beater"
}

func (bracket *Bracket) GongGetGongstructName() string {
	return "Bracket"
}

func (dashes *Dashes) GongGetGongstructName() string {
	return "Dashes"
}

func (degree *Degree) GongGetGongstructName() string {
	return "Degree"
}

func (degreealter *DegreeAlter) GongGetGongstructName() string {
	return "DegreeAlter"
}

func (degreetype *DegreeType) GongGetGongstructName() string {
	return "DegreeType"
}

func (degreevalue *DegreeValue) GongGetGongstructName() string {
	return "DegreeValue"
}

func (direction *Direction) GongGetGongstructName() string {
	return "Direction"
}

func (directiontype *DirectionType) GongGetGongstructName() string {
	return "DirectionType"
}

func (effect *Effect) GongGetGongstructName() string {
	return "Effect"
}

func (feature *Feature) GongGetGongstructName() string {
	return "Feature"
}

func (firstfret *FirstFret) GongGetGongstructName() string {
	return "FirstFret"
}

func (frame *Frame) GongGetGongstructName() string {
	return "Frame"
}

func (framenote *FrameNote) GongGetGongstructName() string {
	return "FrameNote"
}

func (glass *Glass) GongGetGongstructName() string {
	return "Glass"
}

func (grouping *Grouping) GongGetGongstructName() string {
	return "Grouping"
}

func (harmony *Harmony) GongGetGongstructName() string {
	return "Harmony"
}

func (harmonyalter *HarmonyAlter) GongGetGongstructName() string {
	return "HarmonyAlter"
}

func (harppedals *HarpPedals) GongGetGongstructName() string {
	return "HarpPedals"
}

func (image *Image) GongGetGongstructName() string {
	return "Image"
}

func (instrumentchange *InstrumentChange) GongGetGongstructName() string {
	return "InstrumentChange"
}

func (inversion *Inversion) GongGetGongstructName() string {
	return "Inversion"
}

func (kind *Kind) GongGetGongstructName() string {
	return "Kind"
}

func (listening *Listening) GongGetGongstructName() string {
	return "Listening"
}

func (measurenumbering *MeasureNumbering) GongGetGongstructName() string {
	return "MeasureNumbering"
}

func (membrane *Membrane) GongGetGongstructName() string {
	return "Membrane"
}

func (metal *Metal) GongGetGongstructName() string {
	return "Metal"
}

func (metronome *Metronome) GongGetGongstructName() string {
	return "Metronome"
}

func (metronomebeam *MetronomeBeam) GongGetGongstructName() string {
	return "MetronomeBeam"
}

func (metronomenote *MetronomeNote) GongGetGongstructName() string {
	return "MetronomeNote"
}

func (metronometied *MetronomeTied) GongGetGongstructName() string {
	return "MetronomeTied"
}

func (metronometuplet *MetronomeTuplet) GongGetGongstructName() string {
	return "MetronomeTuplet"
}

func (numeral *Numeral) GongGetGongstructName() string {
	return "Numeral"
}

func (numeralkey *NumeralKey) GongGetGongstructName() string {
	return "NumeralKey"
}

func (numeralroot *NumeralRoot) GongGetGongstructName() string {
	return "NumeralRoot"
}

func (octaveshift *OctaveShift) GongGetGongstructName() string {
	return "OctaveShift"
}

func (offset *Offset) GongGetGongstructName() string {
	return "Offset"
}

func (otherdirection *OtherDirection) GongGetGongstructName() string {
	return "OtherDirection"
}

func (otherlistening *OtherListening) GongGetGongstructName() string {
	return "OtherListening"
}

func (pedal *Pedal) GongGetGongstructName() string {
	return "Pedal"
}

func (pedaltuning *PedalTuning) GongGetGongstructName() string {
	return "PedalTuning"
}

func (perminute *PerMinute) GongGetGongstructName() string {
	return "PerMinute"
}

func (percussion *Percussion) GongGetGongstructName() string {
	return "Percussion"
}

func (pitched *Pitched) GongGetGongstructName() string {
	return "Pitched"
}

func (principalvoice *PrincipalVoice) GongGetGongstructName() string {
	return "PrincipalVoice"
}

func (print *Print) GongGetGongstructName() string {
	return "Print"
}

func (root *Root) GongGetGongstructName() string {
	return "Root"
}

func (rootstep *RootStep) GongGetGongstructName() string {
	return "RootStep"
}

func (scordatura *Scordatura) GongGetGongstructName() string {
	return "Scordatura"
}

func (sound *Sound) GongGetGongstructName() string {
	return "Sound"
}

func (staffdivide *StaffDivide) GongGetGongstructName() string {
	return "StaffDivide"
}

func (stick *Stick) GongGetGongstructName() string {
	return "Stick"
}

func (stringmute *StringMute) GongGetGongstructName() string {
	return "StringMute"
}

func (styletext *StyleText) GongGetGongstructName() string {
	return "StyleText"
}

func (swing *Swing) GongGetGongstructName() string {
	return "Swing"
}

func (sync *Sync) GongGetGongstructName() string {
	return "Sync"
}

func (timemodification *TimeModification) GongGetGongstructName() string {
	return "TimeModification"
}

func (timpani *Timpani) GongGetGongstructName() string {
	return "Timpani"
}

func (virtualinstrument *VirtualInstrument) GongGetGongstructName() string {
	return "VirtualInstrument"
}

func (wedge *Wedge) GongGetGongstructName() string {
	return "Wedge"
}

func (wood *Wood) GongGetGongstructName() string {
	return "Wood"
}

func GongGetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	return GongGetGongstructNameFromPointer(instance)
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	__gong__rebuildMapString(stage.Accords, &stage.Accords_mapString)

	__gong__rebuildMapString(stage.AccordionRegistrations, &stage.AccordionRegistrations_mapString)

	__gong__rebuildMapString(stage.Barres, &stage.Barres_mapString)

	__gong__rebuildMapString(stage.Basss, &stage.Basss_mapString)

	__gong__rebuildMapString(stage.BassSteps, &stage.BassSteps_mapString)

	__gong__rebuildMapString(stage.BeatUnitTieds, &stage.BeatUnitTieds_mapString)

	__gong__rebuildMapString(stage.Beaters, &stage.Beaters_mapString)

	__gong__rebuildMapString(stage.Brackets, &stage.Brackets_mapString)

	__gong__rebuildMapString(stage.Dashess, &stage.Dashess_mapString)

	__gong__rebuildMapString(stage.Degrees, &stage.Degrees_mapString)

	__gong__rebuildMapString(stage.DegreeAlters, &stage.DegreeAlters_mapString)

	__gong__rebuildMapString(stage.DegreeTypes, &stage.DegreeTypes_mapString)

	__gong__rebuildMapString(stage.DegreeValues, &stage.DegreeValues_mapString)

	__gong__rebuildMapString(stage.Directions, &stage.Directions_mapString)

	__gong__rebuildMapString(stage.DirectionTypes, &stage.DirectionTypes_mapString)

	__gong__rebuildMapString(stage.Effects, &stage.Effects_mapString)

	__gong__rebuildMapString(stage.Features, &stage.Features_mapString)

	__gong__rebuildMapString(stage.FirstFrets, &stage.FirstFrets_mapString)

	__gong__rebuildMapString(stage.Frames, &stage.Frames_mapString)

	__gong__rebuildMapString(stage.FrameNotes, &stage.FrameNotes_mapString)

	__gong__rebuildMapString(stage.Glasss, &stage.Glasss_mapString)

	__gong__rebuildMapString(stage.Groupings, &stage.Groupings_mapString)

	__gong__rebuildMapString(stage.Harmonys, &stage.Harmonys_mapString)

	__gong__rebuildMapString(stage.HarmonyAlters, &stage.HarmonyAlters_mapString)

	__gong__rebuildMapString(stage.HarpPedalss, &stage.HarpPedalss_mapString)

	__gong__rebuildMapString(stage.Images, &stage.Images_mapString)

	__gong__rebuildMapString(stage.InstrumentChanges, &stage.InstrumentChanges_mapString)

	__gong__rebuildMapString(stage.Inversions, &stage.Inversions_mapString)

	__gong__rebuildMapString(stage.Kinds, &stage.Kinds_mapString)

	__gong__rebuildMapString(stage.Listenings, &stage.Listenings_mapString)

	__gong__rebuildMapString(stage.MeasureNumberings, &stage.MeasureNumberings_mapString)

	__gong__rebuildMapString(stage.Membranes, &stage.Membranes_mapString)

	__gong__rebuildMapString(stage.Metals, &stage.Metals_mapString)

	__gong__rebuildMapString(stage.Metronomes, &stage.Metronomes_mapString)

	__gong__rebuildMapString(stage.MetronomeBeams, &stage.MetronomeBeams_mapString)

	__gong__rebuildMapString(stage.MetronomeNotes, &stage.MetronomeNotes_mapString)

	__gong__rebuildMapString(stage.MetronomeTieds, &stage.MetronomeTieds_mapString)

	__gong__rebuildMapString(stage.MetronomeTuplets, &stage.MetronomeTuplets_mapString)

	__gong__rebuildMapString(stage.Numerals, &stage.Numerals_mapString)

	__gong__rebuildMapString(stage.NumeralKeys, &stage.NumeralKeys_mapString)

	__gong__rebuildMapString(stage.NumeralRoots, &stage.NumeralRoots_mapString)

	__gong__rebuildMapString(stage.OctaveShifts, &stage.OctaveShifts_mapString)

	__gong__rebuildMapString(stage.Offsets, &stage.Offsets_mapString)

	__gong__rebuildMapString(stage.OtherDirections, &stage.OtherDirections_mapString)

	__gong__rebuildMapString(stage.OtherListenings, &stage.OtherListenings_mapString)

	__gong__rebuildMapString(stage.Pedals, &stage.Pedals_mapString)

	__gong__rebuildMapString(stage.PedalTunings, &stage.PedalTunings_mapString)

	__gong__rebuildMapString(stage.PerMinutes, &stage.PerMinutes_mapString)

	__gong__rebuildMapString(stage.Percussions, &stage.Percussions_mapString)

	__gong__rebuildMapString(stage.Pitcheds, &stage.Pitcheds_mapString)

	__gong__rebuildMapString(stage.PrincipalVoices, &stage.PrincipalVoices_mapString)

	__gong__rebuildMapString(stage.Prints, &stage.Prints_mapString)

	__gong__rebuildMapString(stage.Roots, &stage.Roots_mapString)

	__gong__rebuildMapString(stage.RootSteps, &stage.RootSteps_mapString)

	__gong__rebuildMapString(stage.Scordaturas, &stage.Scordaturas_mapString)

	__gong__rebuildMapString(stage.Sounds, &stage.Sounds_mapString)

	__gong__rebuildMapString(stage.StaffDivides, &stage.StaffDivides_mapString)

	__gong__rebuildMapString(stage.Sticks, &stage.Sticks_mapString)

	__gong__rebuildMapString(stage.StringMutes, &stage.StringMutes_mapString)

	__gong__rebuildMapString(stage.StyleTexts, &stage.StyleTexts_mapString)

	__gong__rebuildMapString(stage.Swings, &stage.Swings_mapString)

	__gong__rebuildMapString(stage.Syncs, &stage.Syncs_mapString)

	__gong__rebuildMapString(stage.TimeModifications, &stage.TimeModifications_mapString)

	__gong__rebuildMapString(stage.Timpanis, &stage.Timpanis_mapString)

	__gong__rebuildMapString(stage.VirtualInstruments, &stage.VirtualInstruments_mapString)

	__gong__rebuildMapString(stage.Wedges, &stage.Wedges_mapString)

	__gong__rebuildMapString(stage.Woods, &stage.Woods_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
