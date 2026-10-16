// generated code - do not edit
package note

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
	Accidentals                map[*Accidental]struct{}
	Accidentals_instance       map[*Accidental]*Accidental
	Accidentals_mapString      map[string]*Accidental
	AccidentalOrder            uint
	Accidental_stagedOrder     map[*Accidental]uint
	Accidental_orderStaged     map[uint]*Accidental
	Accidentals_reference      map[*Accidental]*Accidental
	Accidentals_referenceOrder map[*Accidental]uint

	// insertion point for slice of pointers maps
	OnAfterAccidentalCreateCallback GongOnAfterCreateInterface[Accidental]
	OnAfterAccidentalUpdateCallback GongOnAfterUpdateInterface[Accidental]
	OnAfterAccidentalDeleteCallback GongOnAfterDeleteInterface[Accidental]

	AccidentalMarks                map[*AccidentalMark]struct{}
	AccidentalMarks_instance       map[*AccidentalMark]*AccidentalMark
	AccidentalMarks_mapString      map[string]*AccidentalMark
	AccidentalMarkOrder            uint
	AccidentalMark_stagedOrder     map[*AccidentalMark]uint
	AccidentalMark_orderStaged     map[uint]*AccidentalMark
	AccidentalMarks_reference      map[*AccidentalMark]*AccidentalMark
	AccidentalMarks_referenceOrder map[*AccidentalMark]uint

	// insertion point for slice of pointers maps
	OnAfterAccidentalMarkCreateCallback GongOnAfterCreateInterface[AccidentalMark]
	OnAfterAccidentalMarkUpdateCallback GongOnAfterUpdateInterface[AccidentalMark]
	OnAfterAccidentalMarkDeleteCallback GongOnAfterDeleteInterface[AccidentalMark]

	Arpeggiates                map[*Arpeggiate]struct{}
	Arpeggiates_instance       map[*Arpeggiate]*Arpeggiate
	Arpeggiates_mapString      map[string]*Arpeggiate
	ArpeggiateOrder            uint
	Arpeggiate_stagedOrder     map[*Arpeggiate]uint
	Arpeggiate_orderStaged     map[uint]*Arpeggiate
	Arpeggiates_reference      map[*Arpeggiate]*Arpeggiate
	Arpeggiates_referenceOrder map[*Arpeggiate]uint

	// insertion point for slice of pointers maps
	OnAfterArpeggiateCreateCallback GongOnAfterCreateInterface[Arpeggiate]
	OnAfterArpeggiateUpdateCallback GongOnAfterUpdateInterface[Arpeggiate]
	OnAfterArpeggiateDeleteCallback GongOnAfterDeleteInterface[Arpeggiate]

	Arrows                map[*Arrow]struct{}
	Arrows_instance       map[*Arrow]*Arrow
	Arrows_mapString      map[string]*Arrow
	ArrowOrder            uint
	Arrow_stagedOrder     map[*Arrow]uint
	Arrow_orderStaged     map[uint]*Arrow
	Arrows_reference      map[*Arrow]*Arrow
	Arrows_referenceOrder map[*Arrow]uint

	// insertion point for slice of pointers maps
	OnAfterArrowCreateCallback GongOnAfterCreateInterface[Arrow]
	OnAfterArrowUpdateCallback GongOnAfterUpdateInterface[Arrow]
	OnAfterArrowDeleteCallback GongOnAfterDeleteInterface[Arrow]

	Articulationss                map[*Articulations]struct{}
	Articulationss_instance       map[*Articulations]*Articulations
	Articulationss_mapString      map[string]*Articulations
	ArticulationsOrder            uint
	Articulations_stagedOrder     map[*Articulations]uint
	Articulations_orderStaged     map[uint]*Articulations
	Articulationss_reference      map[*Articulations]*Articulations
	Articulationss_referenceOrder map[*Articulations]uint

	// insertion point for slice of pointers maps
	Articulations_StrongAccent_reverseMap map[*StrongAccent]*Articulations

	Articulations_Scoop_reverseMap map[*EmptyLine]*Articulations

	Articulations_Plop_reverseMap map[*EmptyLine]*Articulations

	Articulations_Doit_reverseMap map[*EmptyLine]*Articulations

	Articulations_Falloff_reverseMap map[*EmptyLine]*Articulations

	Articulations_BreathMark_reverseMap map[*BreathMark]*Articulations

	Articulations_Caesura_reverseMap map[*Caesura]*Articulations

	Articulations_OtherArticulation_reverseMap map[*OtherPlacementText]*Articulations

	OnAfterArticulationsCreateCallback GongOnAfterCreateInterface[Articulations]
	OnAfterArticulationsUpdateCallback GongOnAfterUpdateInterface[Articulations]
	OnAfterArticulationsDeleteCallback GongOnAfterDeleteInterface[Articulations]

	Assesss                map[*Assess]struct{}
	Assesss_instance       map[*Assess]*Assess
	Assesss_mapString      map[string]*Assess
	AssessOrder            uint
	Assess_stagedOrder     map[*Assess]uint
	Assess_orderStaged     map[uint]*Assess
	Assesss_reference      map[*Assess]*Assess
	Assesss_referenceOrder map[*Assess]uint

	// insertion point for slice of pointers maps
	OnAfterAssessCreateCallback GongOnAfterCreateInterface[Assess]
	OnAfterAssessUpdateCallback GongOnAfterUpdateInterface[Assess]
	OnAfterAssessDeleteCallback GongOnAfterDeleteInterface[Assess]

	Backups                map[*Backup]struct{}
	Backups_instance       map[*Backup]*Backup
	Backups_mapString      map[string]*Backup
	BackupOrder            uint
	Backup_stagedOrder     map[*Backup]uint
	Backup_orderStaged     map[uint]*Backup
	Backups_reference      map[*Backup]*Backup
	Backups_referenceOrder map[*Backup]uint

	// insertion point for slice of pointers maps
	OnAfterBackupCreateCallback GongOnAfterCreateInterface[Backup]
	OnAfterBackupUpdateCallback GongOnAfterUpdateInterface[Backup]
	OnAfterBackupDeleteCallback GongOnAfterDeleteInterface[Backup]

	Beams                map[*Beam]struct{}
	Beams_instance       map[*Beam]*Beam
	Beams_mapString      map[string]*Beam
	BeamOrder            uint
	Beam_stagedOrder     map[*Beam]uint
	Beam_orderStaged     map[uint]*Beam
	Beams_reference      map[*Beam]*Beam
	Beams_referenceOrder map[*Beam]uint

	// insertion point for slice of pointers maps
	OnAfterBeamCreateCallback GongOnAfterCreateInterface[Beam]
	OnAfterBeamUpdateCallback GongOnAfterUpdateInterface[Beam]
	OnAfterBeamDeleteCallback GongOnAfterDeleteInterface[Beam]

	Bends                map[*Bend]struct{}
	Bends_instance       map[*Bend]*Bend
	Bends_mapString      map[string]*Bend
	BendOrder            uint
	Bend_stagedOrder     map[*Bend]uint
	Bend_orderStaged     map[uint]*Bend
	Bends_reference      map[*Bend]*Bend
	Bends_referenceOrder map[*Bend]uint

	// insertion point for slice of pointers maps
	OnAfterBendCreateCallback GongOnAfterCreateInterface[Bend]
	OnAfterBendUpdateCallback GongOnAfterUpdateInterface[Bend]
	OnAfterBendDeleteCallback GongOnAfterDeleteInterface[Bend]

	BreathMarks                map[*BreathMark]struct{}
	BreathMarks_instance       map[*BreathMark]*BreathMark
	BreathMarks_mapString      map[string]*BreathMark
	BreathMarkOrder            uint
	BreathMark_stagedOrder     map[*BreathMark]uint
	BreathMark_orderStaged     map[uint]*BreathMark
	BreathMarks_reference      map[*BreathMark]*BreathMark
	BreathMarks_referenceOrder map[*BreathMark]uint

	// insertion point for slice of pointers maps
	OnAfterBreathMarkCreateCallback GongOnAfterCreateInterface[BreathMark]
	OnAfterBreathMarkUpdateCallback GongOnAfterUpdateInterface[BreathMark]
	OnAfterBreathMarkDeleteCallback GongOnAfterDeleteInterface[BreathMark]

	Caesuras                map[*Caesura]struct{}
	Caesuras_instance       map[*Caesura]*Caesura
	Caesuras_mapString      map[string]*Caesura
	CaesuraOrder            uint
	Caesura_stagedOrder     map[*Caesura]uint
	Caesura_orderStaged     map[uint]*Caesura
	Caesuras_reference      map[*Caesura]*Caesura
	Caesuras_referenceOrder map[*Caesura]uint

	// insertion point for slice of pointers maps
	OnAfterCaesuraCreateCallback GongOnAfterCreateInterface[Caesura]
	OnAfterCaesuraUpdateCallback GongOnAfterUpdateInterface[Caesura]
	OnAfterCaesuraDeleteCallback GongOnAfterDeleteInterface[Caesura]

	Elisions                map[*Elision]struct{}
	Elisions_instance       map[*Elision]*Elision
	Elisions_mapString      map[string]*Elision
	ElisionOrder            uint
	Elision_stagedOrder     map[*Elision]uint
	Elision_orderStaged     map[uint]*Elision
	Elisions_reference      map[*Elision]*Elision
	Elisions_referenceOrder map[*Elision]uint

	// insertion point for slice of pointers maps
	OnAfterElisionCreateCallback GongOnAfterCreateInterface[Elision]
	OnAfterElisionUpdateCallback GongOnAfterUpdateInterface[Elision]
	OnAfterElisionDeleteCallback GongOnAfterDeleteInterface[Elision]

	EmptyLines                map[*EmptyLine]struct{}
	EmptyLines_instance       map[*EmptyLine]*EmptyLine
	EmptyLines_mapString      map[string]*EmptyLine
	EmptyLineOrder            uint
	EmptyLine_stagedOrder     map[*EmptyLine]uint
	EmptyLine_orderStaged     map[uint]*EmptyLine
	EmptyLines_reference      map[*EmptyLine]*EmptyLine
	EmptyLines_referenceOrder map[*EmptyLine]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyLineCreateCallback GongOnAfterCreateInterface[EmptyLine]
	OnAfterEmptyLineUpdateCallback GongOnAfterUpdateInterface[EmptyLine]
	OnAfterEmptyLineDeleteCallback GongOnAfterDeleteInterface[EmptyLine]

	Extends                map[*Extend]struct{}
	Extends_instance       map[*Extend]*Extend
	Extends_mapString      map[string]*Extend
	ExtendOrder            uint
	Extend_stagedOrder     map[*Extend]uint
	Extend_orderStaged     map[uint]*Extend
	Extends_reference      map[*Extend]*Extend
	Extends_referenceOrder map[*Extend]uint

	// insertion point for slice of pointers maps
	OnAfterExtendCreateCallback GongOnAfterCreateInterface[Extend]
	OnAfterExtendUpdateCallback GongOnAfterUpdateInterface[Extend]
	OnAfterExtendDeleteCallback GongOnAfterDeleteInterface[Extend]

	Figures                map[*Figure]struct{}
	Figures_instance       map[*Figure]*Figure
	Figures_mapString      map[string]*Figure
	FigureOrder            uint
	Figure_stagedOrder     map[*Figure]uint
	Figure_orderStaged     map[uint]*Figure
	Figures_reference      map[*Figure]*Figure
	Figures_referenceOrder map[*Figure]uint

	// insertion point for slice of pointers maps
	OnAfterFigureCreateCallback GongOnAfterCreateInterface[Figure]
	OnAfterFigureUpdateCallback GongOnAfterUpdateInterface[Figure]
	OnAfterFigureDeleteCallback GongOnAfterDeleteInterface[Figure]

	FiguredBasss                map[*FiguredBass]struct{}
	FiguredBasss_instance       map[*FiguredBass]*FiguredBass
	FiguredBasss_mapString      map[string]*FiguredBass
	FiguredBassOrder            uint
	FiguredBass_stagedOrder     map[*FiguredBass]uint
	FiguredBass_orderStaged     map[uint]*FiguredBass
	FiguredBasss_reference      map[*FiguredBass]*FiguredBass
	FiguredBasss_referenceOrder map[*FiguredBass]uint

	// insertion point for slice of pointers maps
	FiguredBass_Figure_reverseMap map[*Figure]*FiguredBass

	OnAfterFiguredBassCreateCallback GongOnAfterCreateInterface[FiguredBass]
	OnAfterFiguredBassUpdateCallback GongOnAfterUpdateInterface[FiguredBass]
	OnAfterFiguredBassDeleteCallback GongOnAfterDeleteInterface[FiguredBass]

	Forwards                map[*Forward]struct{}
	Forwards_instance       map[*Forward]*Forward
	Forwards_mapString      map[string]*Forward
	ForwardOrder            uint
	Forward_stagedOrder     map[*Forward]uint
	Forward_orderStaged     map[uint]*Forward
	Forwards_reference      map[*Forward]*Forward
	Forwards_referenceOrder map[*Forward]uint

	// insertion point for slice of pointers maps
	OnAfterForwardCreateCallback GongOnAfterCreateInterface[Forward]
	OnAfterForwardUpdateCallback GongOnAfterUpdateInterface[Forward]
	OnAfterForwardDeleteCallback GongOnAfterDeleteInterface[Forward]

	Glissandos                map[*Glissando]struct{}
	Glissandos_instance       map[*Glissando]*Glissando
	Glissandos_mapString      map[string]*Glissando
	GlissandoOrder            uint
	Glissando_stagedOrder     map[*Glissando]uint
	Glissando_orderStaged     map[uint]*Glissando
	Glissandos_reference      map[*Glissando]*Glissando
	Glissandos_referenceOrder map[*Glissando]uint

	// insertion point for slice of pointers maps
	OnAfterGlissandoCreateCallback GongOnAfterCreateInterface[Glissando]
	OnAfterGlissandoUpdateCallback GongOnAfterUpdateInterface[Glissando]
	OnAfterGlissandoDeleteCallback GongOnAfterDeleteInterface[Glissando]

	Graces                map[*Grace]struct{}
	Graces_instance       map[*Grace]*Grace
	Graces_mapString      map[string]*Grace
	GraceOrder            uint
	Grace_stagedOrder     map[*Grace]uint
	Grace_orderStaged     map[uint]*Grace
	Graces_reference      map[*Grace]*Grace
	Graces_referenceOrder map[*Grace]uint

	// insertion point for slice of pointers maps
	OnAfterGraceCreateCallback GongOnAfterCreateInterface[Grace]
	OnAfterGraceUpdateCallback GongOnAfterUpdateInterface[Grace]
	OnAfterGraceDeleteCallback GongOnAfterDeleteInterface[Grace]

	HammerOnPullOffs                map[*HammerOnPullOff]struct{}
	HammerOnPullOffs_instance       map[*HammerOnPullOff]*HammerOnPullOff
	HammerOnPullOffs_mapString      map[string]*HammerOnPullOff
	HammerOnPullOffOrder            uint
	HammerOnPullOff_stagedOrder     map[*HammerOnPullOff]uint
	HammerOnPullOff_orderStaged     map[uint]*HammerOnPullOff
	HammerOnPullOffs_reference      map[*HammerOnPullOff]*HammerOnPullOff
	HammerOnPullOffs_referenceOrder map[*HammerOnPullOff]uint

	// insertion point for slice of pointers maps
	OnAfterHammerOnPullOffCreateCallback GongOnAfterCreateInterface[HammerOnPullOff]
	OnAfterHammerOnPullOffUpdateCallback GongOnAfterUpdateInterface[HammerOnPullOff]
	OnAfterHammerOnPullOffDeleteCallback GongOnAfterDeleteInterface[HammerOnPullOff]

	Handbells                map[*Handbell]struct{}
	Handbells_instance       map[*Handbell]*Handbell
	Handbells_mapString      map[string]*Handbell
	HandbellOrder            uint
	Handbell_stagedOrder     map[*Handbell]uint
	Handbell_orderStaged     map[uint]*Handbell
	Handbells_reference      map[*Handbell]*Handbell
	Handbells_referenceOrder map[*Handbell]uint

	// insertion point for slice of pointers maps
	OnAfterHandbellCreateCallback GongOnAfterCreateInterface[Handbell]
	OnAfterHandbellUpdateCallback GongOnAfterUpdateInterface[Handbell]
	OnAfterHandbellDeleteCallback GongOnAfterDeleteInterface[Handbell]

	HarmonCloseds                map[*HarmonClosed]struct{}
	HarmonCloseds_instance       map[*HarmonClosed]*HarmonClosed
	HarmonCloseds_mapString      map[string]*HarmonClosed
	HarmonClosedOrder            uint
	HarmonClosed_stagedOrder     map[*HarmonClosed]uint
	HarmonClosed_orderStaged     map[uint]*HarmonClosed
	HarmonCloseds_reference      map[*HarmonClosed]*HarmonClosed
	HarmonCloseds_referenceOrder map[*HarmonClosed]uint

	// insertion point for slice of pointers maps
	OnAfterHarmonClosedCreateCallback GongOnAfterCreateInterface[HarmonClosed]
	OnAfterHarmonClosedUpdateCallback GongOnAfterUpdateInterface[HarmonClosed]
	OnAfterHarmonClosedDeleteCallback GongOnAfterDeleteInterface[HarmonClosed]

	HarmonMutes                map[*HarmonMute]struct{}
	HarmonMutes_instance       map[*HarmonMute]*HarmonMute
	HarmonMutes_mapString      map[string]*HarmonMute
	HarmonMuteOrder            uint
	HarmonMute_stagedOrder     map[*HarmonMute]uint
	HarmonMute_orderStaged     map[uint]*HarmonMute
	HarmonMutes_reference      map[*HarmonMute]*HarmonMute
	HarmonMutes_referenceOrder map[*HarmonMute]uint

	// insertion point for slice of pointers maps
	OnAfterHarmonMuteCreateCallback GongOnAfterCreateInterface[HarmonMute]
	OnAfterHarmonMuteUpdateCallback GongOnAfterUpdateInterface[HarmonMute]
	OnAfterHarmonMuteDeleteCallback GongOnAfterDeleteInterface[HarmonMute]

	Harmonics                map[*Harmonic]struct{}
	Harmonics_instance       map[*Harmonic]*Harmonic
	Harmonics_mapString      map[string]*Harmonic
	HarmonicOrder            uint
	Harmonic_stagedOrder     map[*Harmonic]uint
	Harmonic_orderStaged     map[uint]*Harmonic
	Harmonics_reference      map[*Harmonic]*Harmonic
	Harmonics_referenceOrder map[*Harmonic]uint

	// insertion point for slice of pointers maps
	OnAfterHarmonicCreateCallback GongOnAfterCreateInterface[Harmonic]
	OnAfterHarmonicUpdateCallback GongOnAfterUpdateInterface[Harmonic]
	OnAfterHarmonicDeleteCallback GongOnAfterDeleteInterface[Harmonic]

	HeelToes                map[*HeelToe]struct{}
	HeelToes_instance       map[*HeelToe]*HeelToe
	HeelToes_mapString      map[string]*HeelToe
	HeelToeOrder            uint
	HeelToe_stagedOrder     map[*HeelToe]uint
	HeelToe_orderStaged     map[uint]*HeelToe
	HeelToes_reference      map[*HeelToe]*HeelToe
	HeelToes_referenceOrder map[*HeelToe]uint

	// insertion point for slice of pointers maps
	OnAfterHeelToeCreateCallback GongOnAfterCreateInterface[HeelToe]
	OnAfterHeelToeUpdateCallback GongOnAfterUpdateInterface[HeelToe]
	OnAfterHeelToeDeleteCallback GongOnAfterDeleteInterface[HeelToe]

	Holes                map[*Hole]struct{}
	Holes_instance       map[*Hole]*Hole
	Holes_mapString      map[string]*Hole
	HoleOrder            uint
	Hole_stagedOrder     map[*Hole]uint
	Hole_orderStaged     map[uint]*Hole
	Holes_reference      map[*Hole]*Hole
	Holes_referenceOrder map[*Hole]uint

	// insertion point for slice of pointers maps
	OnAfterHoleCreateCallback GongOnAfterCreateInterface[Hole]
	OnAfterHoleUpdateCallback GongOnAfterUpdateInterface[Hole]
	OnAfterHoleDeleteCallback GongOnAfterDeleteInterface[Hole]

	HoleCloseds                map[*HoleClosed]struct{}
	HoleCloseds_instance       map[*HoleClosed]*HoleClosed
	HoleCloseds_mapString      map[string]*HoleClosed
	HoleClosedOrder            uint
	HoleClosed_stagedOrder     map[*HoleClosed]uint
	HoleClosed_orderStaged     map[uint]*HoleClosed
	HoleCloseds_reference      map[*HoleClosed]*HoleClosed
	HoleCloseds_referenceOrder map[*HoleClosed]uint

	// insertion point for slice of pointers maps
	OnAfterHoleClosedCreateCallback GongOnAfterCreateInterface[HoleClosed]
	OnAfterHoleClosedUpdateCallback GongOnAfterUpdateInterface[HoleClosed]
	OnAfterHoleClosedDeleteCallback GongOnAfterDeleteInterface[HoleClosed]

	Instruments                map[*Instrument]struct{}
	Instruments_instance       map[*Instrument]*Instrument
	Instruments_mapString      map[string]*Instrument
	InstrumentOrder            uint
	Instrument_stagedOrder     map[*Instrument]uint
	Instrument_orderStaged     map[uint]*Instrument
	Instruments_reference      map[*Instrument]*Instrument
	Instruments_referenceOrder map[*Instrument]uint

	// insertion point for slice of pointers maps
	OnAfterInstrumentCreateCallback GongOnAfterCreateInterface[Instrument]
	OnAfterInstrumentUpdateCallback GongOnAfterUpdateInterface[Instrument]
	OnAfterInstrumentDeleteCallback GongOnAfterDeleteInterface[Instrument]

	Listens                map[*Listen]struct{}
	Listens_instance       map[*Listen]*Listen
	Listens_mapString      map[string]*Listen
	ListenOrder            uint
	Listen_stagedOrder     map[*Listen]uint
	Listen_orderStaged     map[uint]*Listen
	Listens_reference      map[*Listen]*Listen
	Listens_referenceOrder map[*Listen]uint

	// insertion point for slice of pointers maps
	Listen_Assess_reverseMap map[*Assess]*Listen

	Listen_Wait_reverseMap map[*Wait]*Listen

	OnAfterListenCreateCallback GongOnAfterCreateInterface[Listen]
	OnAfterListenUpdateCallback GongOnAfterUpdateInterface[Listen]
	OnAfterListenDeleteCallback GongOnAfterDeleteInterface[Listen]

	Lyrics                map[*Lyric]struct{}
	Lyrics_instance       map[*Lyric]*Lyric
	Lyrics_mapString      map[string]*Lyric
	LyricOrder            uint
	Lyric_stagedOrder     map[*Lyric]uint
	Lyric_orderStaged     map[uint]*Lyric
	Lyrics_reference      map[*Lyric]*Lyric
	Lyrics_referenceOrder map[*Lyric]uint

	// insertion point for slice of pointers maps
	Lyric_Elision_reverseMap map[*Elision]*Lyric

	OnAfterLyricCreateCallback GongOnAfterCreateInterface[Lyric]
	OnAfterLyricUpdateCallback GongOnAfterUpdateInterface[Lyric]
	OnAfterLyricDeleteCallback GongOnAfterDeleteInterface[Lyric]

	Mordents                map[*Mordent]struct{}
	Mordents_instance       map[*Mordent]*Mordent
	Mordents_mapString      map[string]*Mordent
	MordentOrder            uint
	Mordent_stagedOrder     map[*Mordent]uint
	Mordent_orderStaged     map[uint]*Mordent
	Mordents_reference      map[*Mordent]*Mordent
	Mordents_referenceOrder map[*Mordent]uint

	// insertion point for slice of pointers maps
	OnAfterMordentCreateCallback GongOnAfterCreateInterface[Mordent]
	OnAfterMordentUpdateCallback GongOnAfterUpdateInterface[Mordent]
	OnAfterMordentDeleteCallback GongOnAfterDeleteInterface[Mordent]

	NonArpeggiates                map[*NonArpeggiate]struct{}
	NonArpeggiates_instance       map[*NonArpeggiate]*NonArpeggiate
	NonArpeggiates_mapString      map[string]*NonArpeggiate
	NonArpeggiateOrder            uint
	NonArpeggiate_stagedOrder     map[*NonArpeggiate]uint
	NonArpeggiate_orderStaged     map[uint]*NonArpeggiate
	NonArpeggiates_reference      map[*NonArpeggiate]*NonArpeggiate
	NonArpeggiates_referenceOrder map[*NonArpeggiate]uint

	// insertion point for slice of pointers maps
	OnAfterNonArpeggiateCreateCallback GongOnAfterCreateInterface[NonArpeggiate]
	OnAfterNonArpeggiateUpdateCallback GongOnAfterUpdateInterface[NonArpeggiate]
	OnAfterNonArpeggiateDeleteCallback GongOnAfterDeleteInterface[NonArpeggiate]

	Notationss                map[*Notations]struct{}
	Notationss_instance       map[*Notations]*Notations
	Notationss_mapString      map[string]*Notations
	NotationsOrder            uint
	Notations_stagedOrder     map[*Notations]uint
	Notations_orderStaged     map[uint]*Notations
	Notationss_reference      map[*Notations]*Notations
	Notationss_referenceOrder map[*Notations]uint

	// insertion point for slice of pointers maps
	Notations_Tied_reverseMap map[*Tied]*Notations

	Notations_Slur_reverseMap map[*Slur]*Notations

	Notations_Tuplet_reverseMap map[*Tuplet]*Notations

	Notations_Glissando_reverseMap map[*Glissando]*Notations

	Notations_Slide_reverseMap map[*Slide]*Notations

	Notations_Ornaments_reverseMap map[*Ornaments]*Notations

	Notations_Technical_reverseMap map[*Technical]*Notations

	Notations_Articulations_reverseMap map[*Articulations]*Notations

	Notations_Arpeggiate_reverseMap map[*Arpeggiate]*Notations

	Notations_NonArpeggiate_reverseMap map[*NonArpeggiate]*Notations

	Notations_AccidentalMark_reverseMap map[*AccidentalMark]*Notations

	Notations_OtherNotation_reverseMap map[*OtherNotation]*Notations

	OnAfterNotationsCreateCallback GongOnAfterCreateInterface[Notations]
	OnAfterNotationsUpdateCallback GongOnAfterUpdateInterface[Notations]
	OnAfterNotationsDeleteCallback GongOnAfterDeleteInterface[Notations]

	Notes                map[*Note]struct{}
	Notes_instance       map[*Note]*Note
	Notes_mapString      map[string]*Note
	NoteOrder            uint
	Note_stagedOrder     map[*Note]uint
	Note_orderStaged     map[uint]*Note
	Notes_reference      map[*Note]*Note
	Notes_referenceOrder map[*Note]uint

	// insertion point for slice of pointers maps
	Note_Tie_reverseMap map[*Tie]*Note

	Note_Instrument_reverseMap map[*Instrument]*Note

	Note_Beam_reverseMap map[*Beam]*Note

	Note_Notations_reverseMap map[*Notations]*Note

	Note_Lyric_reverseMap map[*Lyric]*Note

	OnAfterNoteCreateCallback GongOnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback GongOnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback GongOnAfterDeleteInterface[Note]

	NoteTypes                map[*NoteType]struct{}
	NoteTypes_instance       map[*NoteType]*NoteType
	NoteTypes_mapString      map[string]*NoteType
	NoteTypeOrder            uint
	NoteType_stagedOrder     map[*NoteType]uint
	NoteType_orderStaged     map[uint]*NoteType
	NoteTypes_reference      map[*NoteType]*NoteType
	NoteTypes_referenceOrder map[*NoteType]uint

	// insertion point for slice of pointers maps
	OnAfterNoteTypeCreateCallback GongOnAfterCreateInterface[NoteType]
	OnAfterNoteTypeUpdateCallback GongOnAfterUpdateInterface[NoteType]
	OnAfterNoteTypeDeleteCallback GongOnAfterDeleteInterface[NoteType]

	Noteheads                map[*Notehead]struct{}
	Noteheads_instance       map[*Notehead]*Notehead
	Noteheads_mapString      map[string]*Notehead
	NoteheadOrder            uint
	Notehead_stagedOrder     map[*Notehead]uint
	Notehead_orderStaged     map[uint]*Notehead
	Noteheads_reference      map[*Notehead]*Notehead
	Noteheads_referenceOrder map[*Notehead]uint

	// insertion point for slice of pointers maps
	OnAfterNoteheadCreateCallback GongOnAfterCreateInterface[Notehead]
	OnAfterNoteheadUpdateCallback GongOnAfterUpdateInterface[Notehead]
	OnAfterNoteheadDeleteCallback GongOnAfterDeleteInterface[Notehead]

	NoteheadTexts                map[*NoteheadText]struct{}
	NoteheadTexts_instance       map[*NoteheadText]*NoteheadText
	NoteheadTexts_mapString      map[string]*NoteheadText
	NoteheadTextOrder            uint
	NoteheadText_stagedOrder     map[*NoteheadText]uint
	NoteheadText_orderStaged     map[uint]*NoteheadText
	NoteheadTexts_reference      map[*NoteheadText]*NoteheadText
	NoteheadTexts_referenceOrder map[*NoteheadText]uint

	// insertion point for slice of pointers maps
	OnAfterNoteheadTextCreateCallback GongOnAfterCreateInterface[NoteheadText]
	OnAfterNoteheadTextUpdateCallback GongOnAfterUpdateInterface[NoteheadText]
	OnAfterNoteheadTextDeleteCallback GongOnAfterDeleteInterface[NoteheadText]

	Ornamentss                map[*Ornaments]struct{}
	Ornamentss_instance       map[*Ornaments]*Ornaments
	Ornamentss_mapString      map[string]*Ornaments
	OrnamentsOrder            uint
	Ornaments_stagedOrder     map[*Ornaments]uint
	Ornaments_orderStaged     map[uint]*Ornaments
	Ornamentss_reference      map[*Ornaments]*Ornaments
	Ornamentss_referenceOrder map[*Ornaments]uint

	// insertion point for slice of pointers maps
	Ornaments_Mordent_reverseMap map[*Mordent]*Ornaments

	Ornaments_InvertedMordent_reverseMap map[*Mordent]*Ornaments

	Ornaments_Tremolo_reverseMap map[*Tremolo]*Ornaments

	Ornaments_OtherOrnament_reverseMap map[*OtherPlacementText]*Ornaments

	Ornaments_AccidentalMark_reverseMap map[*AccidentalMark]*Ornaments

	OnAfterOrnamentsCreateCallback GongOnAfterCreateInterface[Ornaments]
	OnAfterOrnamentsUpdateCallback GongOnAfterUpdateInterface[Ornaments]
	OnAfterOrnamentsDeleteCallback GongOnAfterDeleteInterface[Ornaments]

	OtherNotations                map[*OtherNotation]struct{}
	OtherNotations_instance       map[*OtherNotation]*OtherNotation
	OtherNotations_mapString      map[string]*OtherNotation
	OtherNotationOrder            uint
	OtherNotation_stagedOrder     map[*OtherNotation]uint
	OtherNotation_orderStaged     map[uint]*OtherNotation
	OtherNotations_reference      map[*OtherNotation]*OtherNotation
	OtherNotations_referenceOrder map[*OtherNotation]uint

	// insertion point for slice of pointers maps
	OnAfterOtherNotationCreateCallback GongOnAfterCreateInterface[OtherNotation]
	OnAfterOtherNotationUpdateCallback GongOnAfterUpdateInterface[OtherNotation]
	OnAfterOtherNotationDeleteCallback GongOnAfterDeleteInterface[OtherNotation]

	OtherPlacementTexts                map[*OtherPlacementText]struct{}
	OtherPlacementTexts_instance       map[*OtherPlacementText]*OtherPlacementText
	OtherPlacementTexts_mapString      map[string]*OtherPlacementText
	OtherPlacementTextOrder            uint
	OtherPlacementText_stagedOrder     map[*OtherPlacementText]uint
	OtherPlacementText_orderStaged     map[uint]*OtherPlacementText
	OtherPlacementTexts_reference      map[*OtherPlacementText]*OtherPlacementText
	OtherPlacementTexts_referenceOrder map[*OtherPlacementText]uint

	// insertion point for slice of pointers maps
	OnAfterOtherPlacementTextCreateCallback GongOnAfterCreateInterface[OtherPlacementText]
	OnAfterOtherPlacementTextUpdateCallback GongOnAfterUpdateInterface[OtherPlacementText]
	OnAfterOtherPlacementTextDeleteCallback GongOnAfterDeleteInterface[OtherPlacementText]

	Pitchs                map[*Pitch]struct{}
	Pitchs_instance       map[*Pitch]*Pitch
	Pitchs_mapString      map[string]*Pitch
	PitchOrder            uint
	Pitch_stagedOrder     map[*Pitch]uint
	Pitch_orderStaged     map[uint]*Pitch
	Pitchs_reference      map[*Pitch]*Pitch
	Pitchs_referenceOrder map[*Pitch]uint

	// insertion point for slice of pointers maps
	OnAfterPitchCreateCallback GongOnAfterCreateInterface[Pitch]
	OnAfterPitchUpdateCallback GongOnAfterUpdateInterface[Pitch]
	OnAfterPitchDeleteCallback GongOnAfterDeleteInterface[Pitch]

	PlacementTexts                map[*PlacementText]struct{}
	PlacementTexts_instance       map[*PlacementText]*PlacementText
	PlacementTexts_mapString      map[string]*PlacementText
	PlacementTextOrder            uint
	PlacementText_stagedOrder     map[*PlacementText]uint
	PlacementText_orderStaged     map[uint]*PlacementText
	PlacementTexts_reference      map[*PlacementText]*PlacementText
	PlacementTexts_referenceOrder map[*PlacementText]uint

	// insertion point for slice of pointers maps
	OnAfterPlacementTextCreateCallback GongOnAfterCreateInterface[PlacementText]
	OnAfterPlacementTextUpdateCallback GongOnAfterUpdateInterface[PlacementText]
	OnAfterPlacementTextDeleteCallback GongOnAfterDeleteInterface[PlacementText]

	Releases                map[*Release]struct{}
	Releases_instance       map[*Release]*Release
	Releases_mapString      map[string]*Release
	ReleaseOrder            uint
	Release_stagedOrder     map[*Release]uint
	Release_orderStaged     map[uint]*Release
	Releases_reference      map[*Release]*Release
	Releases_referenceOrder map[*Release]uint

	// insertion point for slice of pointers maps
	OnAfterReleaseCreateCallback GongOnAfterCreateInterface[Release]
	OnAfterReleaseUpdateCallback GongOnAfterUpdateInterface[Release]
	OnAfterReleaseDeleteCallback GongOnAfterDeleteInterface[Release]

	Rests                map[*Rest]struct{}
	Rests_instance       map[*Rest]*Rest
	Rests_mapString      map[string]*Rest
	RestOrder            uint
	Rest_stagedOrder     map[*Rest]uint
	Rest_orderStaged     map[uint]*Rest
	Rests_reference      map[*Rest]*Rest
	Rests_referenceOrder map[*Rest]uint

	// insertion point for slice of pointers maps
	OnAfterRestCreateCallback GongOnAfterCreateInterface[Rest]
	OnAfterRestUpdateCallback GongOnAfterUpdateInterface[Rest]
	OnAfterRestDeleteCallback GongOnAfterDeleteInterface[Rest]

	Slides                map[*Slide]struct{}
	Slides_instance       map[*Slide]*Slide
	Slides_mapString      map[string]*Slide
	SlideOrder            uint
	Slide_stagedOrder     map[*Slide]uint
	Slide_orderStaged     map[uint]*Slide
	Slides_reference      map[*Slide]*Slide
	Slides_referenceOrder map[*Slide]uint

	// insertion point for slice of pointers maps
	OnAfterSlideCreateCallback GongOnAfterCreateInterface[Slide]
	OnAfterSlideUpdateCallback GongOnAfterUpdateInterface[Slide]
	OnAfterSlideDeleteCallback GongOnAfterDeleteInterface[Slide]

	Slurs                map[*Slur]struct{}
	Slurs_instance       map[*Slur]*Slur
	Slurs_mapString      map[string]*Slur
	SlurOrder            uint
	Slur_stagedOrder     map[*Slur]uint
	Slur_orderStaged     map[uint]*Slur
	Slurs_reference      map[*Slur]*Slur
	Slurs_referenceOrder map[*Slur]uint

	// insertion point for slice of pointers maps
	OnAfterSlurCreateCallback GongOnAfterCreateInterface[Slur]
	OnAfterSlurUpdateCallback GongOnAfterUpdateInterface[Slur]
	OnAfterSlurDeleteCallback GongOnAfterDeleteInterface[Slur]

	Stems                map[*Stem]struct{}
	Stems_instance       map[*Stem]*Stem
	Stems_mapString      map[string]*Stem
	StemOrder            uint
	Stem_stagedOrder     map[*Stem]uint
	Stem_orderStaged     map[uint]*Stem
	Stems_reference      map[*Stem]*Stem
	Stems_referenceOrder map[*Stem]uint

	// insertion point for slice of pointers maps
	OnAfterStemCreateCallback GongOnAfterCreateInterface[Stem]
	OnAfterStemUpdateCallback GongOnAfterUpdateInterface[Stem]
	OnAfterStemDeleteCallback GongOnAfterDeleteInterface[Stem]

	StrongAccents                map[*StrongAccent]struct{}
	StrongAccents_instance       map[*StrongAccent]*StrongAccent
	StrongAccents_mapString      map[string]*StrongAccent
	StrongAccentOrder            uint
	StrongAccent_stagedOrder     map[*StrongAccent]uint
	StrongAccent_orderStaged     map[uint]*StrongAccent
	StrongAccents_reference      map[*StrongAccent]*StrongAccent
	StrongAccents_referenceOrder map[*StrongAccent]uint

	// insertion point for slice of pointers maps
	OnAfterStrongAccentCreateCallback GongOnAfterCreateInterface[StrongAccent]
	OnAfterStrongAccentUpdateCallback GongOnAfterUpdateInterface[StrongAccent]
	OnAfterStrongAccentDeleteCallback GongOnAfterDeleteInterface[StrongAccent]

	Taps                map[*Tap]struct{}
	Taps_instance       map[*Tap]*Tap
	Taps_mapString      map[string]*Tap
	TapOrder            uint
	Tap_stagedOrder     map[*Tap]uint
	Tap_orderStaged     map[uint]*Tap
	Taps_reference      map[*Tap]*Tap
	Taps_referenceOrder map[*Tap]uint

	// insertion point for slice of pointers maps
	OnAfterTapCreateCallback GongOnAfterCreateInterface[Tap]
	OnAfterTapUpdateCallback GongOnAfterUpdateInterface[Tap]
	OnAfterTapDeleteCallback GongOnAfterDeleteInterface[Tap]

	Technicals                map[*Technical]struct{}
	Technicals_instance       map[*Technical]*Technical
	Technicals_mapString      map[string]*Technical
	TechnicalOrder            uint
	Technical_stagedOrder     map[*Technical]uint
	Technical_orderStaged     map[uint]*Technical
	Technicals_reference      map[*Technical]*Technical
	Technicals_referenceOrder map[*Technical]uint

	// insertion point for slice of pointers maps
	Technical_Harmonic_reverseMap map[*Harmonic]*Technical

	Technical_Pluck_reverseMap map[*PlacementText]*Technical

	Technical_HammerOn_reverseMap map[*HammerOnPullOff]*Technical

	Technical_PullOff_reverseMap map[*HammerOnPullOff]*Technical

	Technical_Bend_reverseMap map[*Bend]*Technical

	Technical_Tap_reverseMap map[*Tap]*Technical

	Technical_Heel_reverseMap map[*HeelToe]*Technical

	Technical_Toe_reverseMap map[*HeelToe]*Technical

	Technical_Hole_reverseMap map[*Hole]*Technical

	Technical_Arrow_reverseMap map[*Arrow]*Technical

	Technical_Handbell_reverseMap map[*Handbell]*Technical

	Technical_HarmonMute_reverseMap map[*HarmonMute]*Technical

	Technical_OtherTechnical_reverseMap map[*OtherPlacementText]*Technical

	OnAfterTechnicalCreateCallback GongOnAfterCreateInterface[Technical]
	OnAfterTechnicalUpdateCallback GongOnAfterUpdateInterface[Technical]
	OnAfterTechnicalDeleteCallback GongOnAfterDeleteInterface[Technical]

	TextElementDatas                map[*TextElementData]struct{}
	TextElementDatas_instance       map[*TextElementData]*TextElementData
	TextElementDatas_mapString      map[string]*TextElementData
	TextElementDataOrder            uint
	TextElementData_stagedOrder     map[*TextElementData]uint
	TextElementData_orderStaged     map[uint]*TextElementData
	TextElementDatas_reference      map[*TextElementData]*TextElementData
	TextElementDatas_referenceOrder map[*TextElementData]uint

	// insertion point for slice of pointers maps
	OnAfterTextElementDataCreateCallback GongOnAfterCreateInterface[TextElementData]
	OnAfterTextElementDataUpdateCallback GongOnAfterUpdateInterface[TextElementData]
	OnAfterTextElementDataDeleteCallback GongOnAfterDeleteInterface[TextElementData]

	Ties                map[*Tie]struct{}
	Ties_instance       map[*Tie]*Tie
	Ties_mapString      map[string]*Tie
	TieOrder            uint
	Tie_stagedOrder     map[*Tie]uint
	Tie_orderStaged     map[uint]*Tie
	Ties_reference      map[*Tie]*Tie
	Ties_referenceOrder map[*Tie]uint

	// insertion point for slice of pointers maps
	OnAfterTieCreateCallback GongOnAfterCreateInterface[Tie]
	OnAfterTieUpdateCallback GongOnAfterUpdateInterface[Tie]
	OnAfterTieDeleteCallback GongOnAfterDeleteInterface[Tie]

	Tieds                map[*Tied]struct{}
	Tieds_instance       map[*Tied]*Tied
	Tieds_mapString      map[string]*Tied
	TiedOrder            uint
	Tied_stagedOrder     map[*Tied]uint
	Tied_orderStaged     map[uint]*Tied
	Tieds_reference      map[*Tied]*Tied
	Tieds_referenceOrder map[*Tied]uint

	// insertion point for slice of pointers maps
	OnAfterTiedCreateCallback GongOnAfterCreateInterface[Tied]
	OnAfterTiedUpdateCallback GongOnAfterUpdateInterface[Tied]
	OnAfterTiedDeleteCallback GongOnAfterDeleteInterface[Tied]

	Tremolos                map[*Tremolo]struct{}
	Tremolos_instance       map[*Tremolo]*Tremolo
	Tremolos_mapString      map[string]*Tremolo
	TremoloOrder            uint
	Tremolo_stagedOrder     map[*Tremolo]uint
	Tremolo_orderStaged     map[uint]*Tremolo
	Tremolos_reference      map[*Tremolo]*Tremolo
	Tremolos_referenceOrder map[*Tremolo]uint

	// insertion point for slice of pointers maps
	OnAfterTremoloCreateCallback GongOnAfterCreateInterface[Tremolo]
	OnAfterTremoloUpdateCallback GongOnAfterUpdateInterface[Tremolo]
	OnAfterTremoloDeleteCallback GongOnAfterDeleteInterface[Tremolo]

	Tuplets                map[*Tuplet]struct{}
	Tuplets_instance       map[*Tuplet]*Tuplet
	Tuplets_mapString      map[string]*Tuplet
	TupletOrder            uint
	Tuplet_stagedOrder     map[*Tuplet]uint
	Tuplet_orderStaged     map[uint]*Tuplet
	Tuplets_reference      map[*Tuplet]*Tuplet
	Tuplets_referenceOrder map[*Tuplet]uint

	// insertion point for slice of pointers maps
	OnAfterTupletCreateCallback GongOnAfterCreateInterface[Tuplet]
	OnAfterTupletUpdateCallback GongOnAfterUpdateInterface[Tuplet]
	OnAfterTupletDeleteCallback GongOnAfterDeleteInterface[Tuplet]

	TupletDots                map[*TupletDot]struct{}
	TupletDots_instance       map[*TupletDot]*TupletDot
	TupletDots_mapString      map[string]*TupletDot
	TupletDotOrder            uint
	TupletDot_stagedOrder     map[*TupletDot]uint
	TupletDot_orderStaged     map[uint]*TupletDot
	TupletDots_reference      map[*TupletDot]*TupletDot
	TupletDots_referenceOrder map[*TupletDot]uint

	// insertion point for slice of pointers maps
	OnAfterTupletDotCreateCallback GongOnAfterCreateInterface[TupletDot]
	OnAfterTupletDotUpdateCallback GongOnAfterUpdateInterface[TupletDot]
	OnAfterTupletDotDeleteCallback GongOnAfterDeleteInterface[TupletDot]

	TupletNumbers                map[*TupletNumber]struct{}
	TupletNumbers_instance       map[*TupletNumber]*TupletNumber
	TupletNumbers_mapString      map[string]*TupletNumber
	TupletNumberOrder            uint
	TupletNumber_stagedOrder     map[*TupletNumber]uint
	TupletNumber_orderStaged     map[uint]*TupletNumber
	TupletNumbers_reference      map[*TupletNumber]*TupletNumber
	TupletNumbers_referenceOrder map[*TupletNumber]uint

	// insertion point for slice of pointers maps
	OnAfterTupletNumberCreateCallback GongOnAfterCreateInterface[TupletNumber]
	OnAfterTupletNumberUpdateCallback GongOnAfterUpdateInterface[TupletNumber]
	OnAfterTupletNumberDeleteCallback GongOnAfterDeleteInterface[TupletNumber]

	TupletPortions                map[*TupletPortion]struct{}
	TupletPortions_instance       map[*TupletPortion]*TupletPortion
	TupletPortions_mapString      map[string]*TupletPortion
	TupletPortionOrder            uint
	TupletPortion_stagedOrder     map[*TupletPortion]uint
	TupletPortion_orderStaged     map[uint]*TupletPortion
	TupletPortions_reference      map[*TupletPortion]*TupletPortion
	TupletPortions_referenceOrder map[*TupletPortion]uint

	// insertion point for slice of pointers maps
	TupletPortion_TupletDot_reverseMap map[*TupletDot]*TupletPortion

	OnAfterTupletPortionCreateCallback GongOnAfterCreateInterface[TupletPortion]
	OnAfterTupletPortionUpdateCallback GongOnAfterUpdateInterface[TupletPortion]
	OnAfterTupletPortionDeleteCallback GongOnAfterDeleteInterface[TupletPortion]

	TupletTypes                map[*TupletType]struct{}
	TupletTypes_instance       map[*TupletType]*TupletType
	TupletTypes_mapString      map[string]*TupletType
	TupletTypeOrder            uint
	TupletType_stagedOrder     map[*TupletType]uint
	TupletType_orderStaged     map[uint]*TupletType
	TupletTypes_reference      map[*TupletType]*TupletType
	TupletTypes_referenceOrder map[*TupletType]uint

	// insertion point for slice of pointers maps
	OnAfterTupletTypeCreateCallback GongOnAfterCreateInterface[TupletType]
	OnAfterTupletTypeUpdateCallback GongOnAfterUpdateInterface[TupletType]
	OnAfterTupletTypeDeleteCallback GongOnAfterDeleteInterface[TupletType]

	Unpitcheds                map[*Unpitched]struct{}
	Unpitcheds_instance       map[*Unpitched]*Unpitched
	Unpitcheds_mapString      map[string]*Unpitched
	UnpitchedOrder            uint
	Unpitched_stagedOrder     map[*Unpitched]uint
	Unpitched_orderStaged     map[uint]*Unpitched
	Unpitcheds_reference      map[*Unpitched]*Unpitched
	Unpitcheds_referenceOrder map[*Unpitched]uint

	// insertion point for slice of pointers maps
	OnAfterUnpitchedCreateCallback GongOnAfterCreateInterface[Unpitched]
	OnAfterUnpitchedUpdateCallback GongOnAfterUpdateInterface[Unpitched]
	OnAfterUnpitchedDeleteCallback GongOnAfterDeleteInterface[Unpitched]

	Waits                map[*Wait]struct{}
	Waits_instance       map[*Wait]*Wait
	Waits_mapString      map[string]*Wait
	WaitOrder            uint
	Wait_stagedOrder     map[*Wait]uint
	Wait_orderStaged     map[uint]*Wait
	Waits_reference      map[*Wait]*Wait
	Waits_referenceOrder map[*Wait]uint

	// insertion point for slice of pointers maps
	OnAfterWaitCreateCallback GongOnAfterCreateInterface[Wait]
	OnAfterWaitUpdateCallback GongOnAfterUpdateInterface[Wait]
	OnAfterWaitDeleteCallback GongOnAfterDeleteInterface[Wait]

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
	__gong__clearReferences(&stage.Accidentals_reference, &stage.Accidentals_instance, &stage.Accidentals_referenceOrder)

	__gong__clearReferences(&stage.AccidentalMarks_reference, &stage.AccidentalMarks_instance, &stage.AccidentalMarks_referenceOrder)

	__gong__clearReferences(&stage.Arpeggiates_reference, &stage.Arpeggiates_instance, &stage.Arpeggiates_referenceOrder)

	__gong__clearReferences(&stage.Arrows_reference, &stage.Arrows_instance, &stage.Arrows_referenceOrder)

	__gong__clearReferences(&stage.Articulationss_reference, &stage.Articulationss_instance, &stage.Articulationss_referenceOrder)

	__gong__clearReferences(&stage.Assesss_reference, &stage.Assesss_instance, &stage.Assesss_referenceOrder)

	__gong__clearReferences(&stage.Backups_reference, &stage.Backups_instance, &stage.Backups_referenceOrder)

	__gong__clearReferences(&stage.Beams_reference, &stage.Beams_instance, &stage.Beams_referenceOrder)

	__gong__clearReferences(&stage.Bends_reference, &stage.Bends_instance, &stage.Bends_referenceOrder)

	__gong__clearReferences(&stage.BreathMarks_reference, &stage.BreathMarks_instance, &stage.BreathMarks_referenceOrder)

	__gong__clearReferences(&stage.Caesuras_reference, &stage.Caesuras_instance, &stage.Caesuras_referenceOrder)

	__gong__clearReferences(&stage.Elisions_reference, &stage.Elisions_instance, &stage.Elisions_referenceOrder)

	__gong__clearReferences(&stage.EmptyLines_reference, &stage.EmptyLines_instance, &stage.EmptyLines_referenceOrder)

	__gong__clearReferences(&stage.Extends_reference, &stage.Extends_instance, &stage.Extends_referenceOrder)

	__gong__clearReferences(&stage.Figures_reference, &stage.Figures_instance, &stage.Figures_referenceOrder)

	__gong__clearReferences(&stage.FiguredBasss_reference, &stage.FiguredBasss_instance, &stage.FiguredBasss_referenceOrder)

	__gong__clearReferences(&stage.Forwards_reference, &stage.Forwards_instance, &stage.Forwards_referenceOrder)

	__gong__clearReferences(&stage.Glissandos_reference, &stage.Glissandos_instance, &stage.Glissandos_referenceOrder)

	__gong__clearReferences(&stage.Graces_reference, &stage.Graces_instance, &stage.Graces_referenceOrder)

	__gong__clearReferences(&stage.HammerOnPullOffs_reference, &stage.HammerOnPullOffs_instance, &stage.HammerOnPullOffs_referenceOrder)

	__gong__clearReferences(&stage.Handbells_reference, &stage.Handbells_instance, &stage.Handbells_referenceOrder)

	__gong__clearReferences(&stage.HarmonCloseds_reference, &stage.HarmonCloseds_instance, &stage.HarmonCloseds_referenceOrder)

	__gong__clearReferences(&stage.HarmonMutes_reference, &stage.HarmonMutes_instance, &stage.HarmonMutes_referenceOrder)

	__gong__clearReferences(&stage.Harmonics_reference, &stage.Harmonics_instance, &stage.Harmonics_referenceOrder)

	__gong__clearReferences(&stage.HeelToes_reference, &stage.HeelToes_instance, &stage.HeelToes_referenceOrder)

	__gong__clearReferences(&stage.Holes_reference, &stage.Holes_instance, &stage.Holes_referenceOrder)

	__gong__clearReferences(&stage.HoleCloseds_reference, &stage.HoleCloseds_instance, &stage.HoleCloseds_referenceOrder)

	__gong__clearReferences(&stage.Instruments_reference, &stage.Instruments_instance, &stage.Instruments_referenceOrder)

	__gong__clearReferences(&stage.Listens_reference, &stage.Listens_instance, &stage.Listens_referenceOrder)

	__gong__clearReferences(&stage.Lyrics_reference, &stage.Lyrics_instance, &stage.Lyrics_referenceOrder)

	__gong__clearReferences(&stage.Mordents_reference, &stage.Mordents_instance, &stage.Mordents_referenceOrder)

	__gong__clearReferences(&stage.NonArpeggiates_reference, &stage.NonArpeggiates_instance, &stage.NonArpeggiates_referenceOrder)

	__gong__clearReferences(&stage.Notationss_reference, &stage.Notationss_instance, &stage.Notationss_referenceOrder)

	__gong__clearReferences(&stage.Notes_reference, &stage.Notes_instance, &stage.Notes_referenceOrder)

	__gong__clearReferences(&stage.NoteTypes_reference, &stage.NoteTypes_instance, &stage.NoteTypes_referenceOrder)

	__gong__clearReferences(&stage.Noteheads_reference, &stage.Noteheads_instance, &stage.Noteheads_referenceOrder)

	__gong__clearReferences(&stage.NoteheadTexts_reference, &stage.NoteheadTexts_instance, &stage.NoteheadTexts_referenceOrder)

	__gong__clearReferences(&stage.Ornamentss_reference, &stage.Ornamentss_instance, &stage.Ornamentss_referenceOrder)

	__gong__clearReferences(&stage.OtherNotations_reference, &stage.OtherNotations_instance, &stage.OtherNotations_referenceOrder)

	__gong__clearReferences(&stage.OtherPlacementTexts_reference, &stage.OtherPlacementTexts_instance, &stage.OtherPlacementTexts_referenceOrder)

	__gong__clearReferences(&stage.Pitchs_reference, &stage.Pitchs_instance, &stage.Pitchs_referenceOrder)

	__gong__clearReferences(&stage.PlacementTexts_reference, &stage.PlacementTexts_instance, &stage.PlacementTexts_referenceOrder)

	__gong__clearReferences(&stage.Releases_reference, &stage.Releases_instance, &stage.Releases_referenceOrder)

	__gong__clearReferences(&stage.Rests_reference, &stage.Rests_instance, &stage.Rests_referenceOrder)

	__gong__clearReferences(&stage.Slides_reference, &stage.Slides_instance, &stage.Slides_referenceOrder)

	__gong__clearReferences(&stage.Slurs_reference, &stage.Slurs_instance, &stage.Slurs_referenceOrder)

	__gong__clearReferences(&stage.Stems_reference, &stage.Stems_instance, &stage.Stems_referenceOrder)

	__gong__clearReferences(&stage.StrongAccents_reference, &stage.StrongAccents_instance, &stage.StrongAccents_referenceOrder)

	__gong__clearReferences(&stage.Taps_reference, &stage.Taps_instance, &stage.Taps_referenceOrder)

	__gong__clearReferences(&stage.Technicals_reference, &stage.Technicals_instance, &stage.Technicals_referenceOrder)

	__gong__clearReferences(&stage.TextElementDatas_reference, &stage.TextElementDatas_instance, &stage.TextElementDatas_referenceOrder)

	__gong__clearReferences(&stage.Ties_reference, &stage.Ties_instance, &stage.Ties_referenceOrder)

	__gong__clearReferences(&stage.Tieds_reference, &stage.Tieds_instance, &stage.Tieds_referenceOrder)

	__gong__clearReferences(&stage.Tremolos_reference, &stage.Tremolos_instance, &stage.Tremolos_referenceOrder)

	__gong__clearReferences(&stage.Tuplets_reference, &stage.Tuplets_instance, &stage.Tuplets_referenceOrder)

	__gong__clearReferences(&stage.TupletDots_reference, &stage.TupletDots_instance, &stage.TupletDots_referenceOrder)

	__gong__clearReferences(&stage.TupletNumbers_reference, &stage.TupletNumbers_instance, &stage.TupletNumbers_referenceOrder)

	__gong__clearReferences(&stage.TupletPortions_reference, &stage.TupletPortions_instance, &stage.TupletPortions_referenceOrder)

	__gong__clearReferences(&stage.TupletTypes_reference, &stage.TupletTypes_instance, &stage.TupletTypes_referenceOrder)

	__gong__clearReferences(&stage.Unpitcheds_reference, &stage.Unpitcheds_instance, &stage.Unpitcheds_referenceOrder)

	__gong__clearReferences(&stage.Waits_reference, &stage.Waits_instance, &stage.Waits_referenceOrder)

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
	stage.AccidentalOrder = __gong__recomputeOrder(stage.Accidental_stagedOrder)

	stage.AccidentalMarkOrder = __gong__recomputeOrder(stage.AccidentalMark_stagedOrder)

	stage.ArpeggiateOrder = __gong__recomputeOrder(stage.Arpeggiate_stagedOrder)

	stage.ArrowOrder = __gong__recomputeOrder(stage.Arrow_stagedOrder)

	stage.ArticulationsOrder = __gong__recomputeOrder(stage.Articulations_stagedOrder)

	stage.AssessOrder = __gong__recomputeOrder(stage.Assess_stagedOrder)

	stage.BackupOrder = __gong__recomputeOrder(stage.Backup_stagedOrder)

	stage.BeamOrder = __gong__recomputeOrder(stage.Beam_stagedOrder)

	stage.BendOrder = __gong__recomputeOrder(stage.Bend_stagedOrder)

	stage.BreathMarkOrder = __gong__recomputeOrder(stage.BreathMark_stagedOrder)

	stage.CaesuraOrder = __gong__recomputeOrder(stage.Caesura_stagedOrder)

	stage.ElisionOrder = __gong__recomputeOrder(stage.Elision_stagedOrder)

	stage.EmptyLineOrder = __gong__recomputeOrder(stage.EmptyLine_stagedOrder)

	stage.ExtendOrder = __gong__recomputeOrder(stage.Extend_stagedOrder)

	stage.FigureOrder = __gong__recomputeOrder(stage.Figure_stagedOrder)

	stage.FiguredBassOrder = __gong__recomputeOrder(stage.FiguredBass_stagedOrder)

	stage.ForwardOrder = __gong__recomputeOrder(stage.Forward_stagedOrder)

	stage.GlissandoOrder = __gong__recomputeOrder(stage.Glissando_stagedOrder)

	stage.GraceOrder = __gong__recomputeOrder(stage.Grace_stagedOrder)

	stage.HammerOnPullOffOrder = __gong__recomputeOrder(stage.HammerOnPullOff_stagedOrder)

	stage.HandbellOrder = __gong__recomputeOrder(stage.Handbell_stagedOrder)

	stage.HarmonClosedOrder = __gong__recomputeOrder(stage.HarmonClosed_stagedOrder)

	stage.HarmonMuteOrder = __gong__recomputeOrder(stage.HarmonMute_stagedOrder)

	stage.HarmonicOrder = __gong__recomputeOrder(stage.Harmonic_stagedOrder)

	stage.HeelToeOrder = __gong__recomputeOrder(stage.HeelToe_stagedOrder)

	stage.HoleOrder = __gong__recomputeOrder(stage.Hole_stagedOrder)

	stage.HoleClosedOrder = __gong__recomputeOrder(stage.HoleClosed_stagedOrder)

	stage.InstrumentOrder = __gong__recomputeOrder(stage.Instrument_stagedOrder)

	stage.ListenOrder = __gong__recomputeOrder(stage.Listen_stagedOrder)

	stage.LyricOrder = __gong__recomputeOrder(stage.Lyric_stagedOrder)

	stage.MordentOrder = __gong__recomputeOrder(stage.Mordent_stagedOrder)

	stage.NonArpeggiateOrder = __gong__recomputeOrder(stage.NonArpeggiate_stagedOrder)

	stage.NotationsOrder = __gong__recomputeOrder(stage.Notations_stagedOrder)

	stage.NoteOrder = __gong__recomputeOrder(stage.Note_stagedOrder)

	stage.NoteTypeOrder = __gong__recomputeOrder(stage.NoteType_stagedOrder)

	stage.NoteheadOrder = __gong__recomputeOrder(stage.Notehead_stagedOrder)

	stage.NoteheadTextOrder = __gong__recomputeOrder(stage.NoteheadText_stagedOrder)

	stage.OrnamentsOrder = __gong__recomputeOrder(stage.Ornaments_stagedOrder)

	stage.OtherNotationOrder = __gong__recomputeOrder(stage.OtherNotation_stagedOrder)

	stage.OtherPlacementTextOrder = __gong__recomputeOrder(stage.OtherPlacementText_stagedOrder)

	stage.PitchOrder = __gong__recomputeOrder(stage.Pitch_stagedOrder)

	stage.PlacementTextOrder = __gong__recomputeOrder(stage.PlacementText_stagedOrder)

	stage.ReleaseOrder = __gong__recomputeOrder(stage.Release_stagedOrder)

	stage.RestOrder = __gong__recomputeOrder(stage.Rest_stagedOrder)

	stage.SlideOrder = __gong__recomputeOrder(stage.Slide_stagedOrder)

	stage.SlurOrder = __gong__recomputeOrder(stage.Slur_stagedOrder)

	stage.StemOrder = __gong__recomputeOrder(stage.Stem_stagedOrder)

	stage.StrongAccentOrder = __gong__recomputeOrder(stage.StrongAccent_stagedOrder)

	stage.TapOrder = __gong__recomputeOrder(stage.Tap_stagedOrder)

	stage.TechnicalOrder = __gong__recomputeOrder(stage.Technical_stagedOrder)

	stage.TextElementDataOrder = __gong__recomputeOrder(stage.TextElementData_stagedOrder)

	stage.TieOrder = __gong__recomputeOrder(stage.Tie_stagedOrder)

	stage.TiedOrder = __gong__recomputeOrder(stage.Tied_stagedOrder)

	stage.TremoloOrder = __gong__recomputeOrder(stage.Tremolo_stagedOrder)

	stage.TupletOrder = __gong__recomputeOrder(stage.Tuplet_stagedOrder)

	stage.TupletDotOrder = __gong__recomputeOrder(stage.TupletDot_stagedOrder)

	stage.TupletNumberOrder = __gong__recomputeOrder(stage.TupletNumber_stagedOrder)

	stage.TupletPortionOrder = __gong__recomputeOrder(stage.TupletPortion_stagedOrder)

	stage.TupletTypeOrder = __gong__recomputeOrder(stage.TupletType_stagedOrder)

	stage.UnpitchedOrder = __gong__recomputeOrder(stage.Unpitched_stagedOrder)

	stage.WaitOrder = __gong__recomputeOrder(stage.Wait_stagedOrder)

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
func (*Accidental) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Accidentals, stage.Accidental_stagedOrder)
}

func (*Accidental) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Accidental_orderStaged[order]
}

func (*Accidental) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Accidentals_mapString
}

func (*Accidental) GongGetInstancesSet(stage *Stage) any {
	return &stage.Accidentals
}

func (*Accidental) GongNewInstance() any {
	return new(Accidental)
}

func (*AccidentalMark) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AccidentalMarks, stage.AccidentalMark_stagedOrder)
}

func (*AccidentalMark) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AccidentalMark_orderStaged[order]
}

func (*AccidentalMark) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AccidentalMarks_mapString
}

func (*AccidentalMark) GongGetInstancesSet(stage *Stage) any {
	return &stage.AccidentalMarks
}

func (*AccidentalMark) GongNewInstance() any {
	return new(AccidentalMark)
}

func (*Arpeggiate) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Arpeggiates, stage.Arpeggiate_stagedOrder)
}

func (*Arpeggiate) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Arpeggiate_orderStaged[order]
}

func (*Arpeggiate) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Arpeggiates_mapString
}

func (*Arpeggiate) GongGetInstancesSet(stage *Stage) any {
	return &stage.Arpeggiates
}

func (*Arpeggiate) GongNewInstance() any {
	return new(Arpeggiate)
}

func (*Arrow) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Arrows, stage.Arrow_stagedOrder)
}

func (*Arrow) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Arrow_orderStaged[order]
}

func (*Arrow) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Arrows_mapString
}

func (*Arrow) GongGetInstancesSet(stage *Stage) any {
	return &stage.Arrows
}

func (*Arrow) GongNewInstance() any {
	return new(Arrow)
}

func (*Articulations) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Articulationss, stage.Articulations_stagedOrder)
}

func (*Articulations) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Articulations_orderStaged[order]
}

func (*Articulations) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Articulationss_mapString
}

func (*Articulations) GongGetInstancesSet(stage *Stage) any {
	return &stage.Articulationss
}

func (*Articulations) GongNewInstance() any {
	return new(Articulations)
}

func (*Assess) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Assesss, stage.Assess_stagedOrder)
}

func (*Assess) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Assess_orderStaged[order]
}

func (*Assess) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Assesss_mapString
}

func (*Assess) GongGetInstancesSet(stage *Stage) any {
	return &stage.Assesss
}

func (*Assess) GongNewInstance() any {
	return new(Assess)
}

func (*Backup) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Backups, stage.Backup_stagedOrder)
}

func (*Backup) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Backup_orderStaged[order]
}

func (*Backup) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Backups_mapString
}

func (*Backup) GongGetInstancesSet(stage *Stage) any {
	return &stage.Backups
}

func (*Backup) GongNewInstance() any {
	return new(Backup)
}

func (*Beam) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Beams, stage.Beam_stagedOrder)
}

func (*Beam) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Beam_orderStaged[order]
}

func (*Beam) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Beams_mapString
}

func (*Beam) GongGetInstancesSet(stage *Stage) any {
	return &stage.Beams
}

func (*Beam) GongNewInstance() any {
	return new(Beam)
}

func (*Bend) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Bends, stage.Bend_stagedOrder)
}

func (*Bend) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Bend_orderStaged[order]
}

func (*Bend) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Bends_mapString
}

func (*Bend) GongGetInstancesSet(stage *Stage) any {
	return &stage.Bends
}

func (*Bend) GongNewInstance() any {
	return new(Bend)
}

func (*BreathMark) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.BreathMarks, stage.BreathMark_stagedOrder)
}

func (*BreathMark) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.BreathMark_orderStaged[order]
}

func (*BreathMark) GongGetInstancesMapByName(stage *Stage) any {
	return stage.BreathMarks_mapString
}

func (*BreathMark) GongGetInstancesSet(stage *Stage) any {
	return &stage.BreathMarks
}

func (*BreathMark) GongNewInstance() any {
	return new(BreathMark)
}

func (*Caesura) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Caesuras, stage.Caesura_stagedOrder)
}

func (*Caesura) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Caesura_orderStaged[order]
}

func (*Caesura) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Caesuras_mapString
}

func (*Caesura) GongGetInstancesSet(stage *Stage) any {
	return &stage.Caesuras
}

func (*Caesura) GongNewInstance() any {
	return new(Caesura)
}

func (*Elision) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Elisions, stage.Elision_stagedOrder)
}

func (*Elision) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Elision_orderStaged[order]
}

func (*Elision) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Elisions_mapString
}

func (*Elision) GongGetInstancesSet(stage *Stage) any {
	return &stage.Elisions
}

func (*Elision) GongNewInstance() any {
	return new(Elision)
}

func (*EmptyLine) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EmptyLines, stage.EmptyLine_stagedOrder)
}

func (*EmptyLine) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EmptyLine_orderStaged[order]
}

func (*EmptyLine) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EmptyLines_mapString
}

func (*EmptyLine) GongGetInstancesSet(stage *Stage) any {
	return &stage.EmptyLines
}

func (*EmptyLine) GongNewInstance() any {
	return new(EmptyLine)
}

func (*Extend) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Extends, stage.Extend_stagedOrder)
}

func (*Extend) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Extend_orderStaged[order]
}

func (*Extend) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Extends_mapString
}

func (*Extend) GongGetInstancesSet(stage *Stage) any {
	return &stage.Extends
}

func (*Extend) GongNewInstance() any {
	return new(Extend)
}

func (*Figure) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Figures, stage.Figure_stagedOrder)
}

func (*Figure) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Figure_orderStaged[order]
}

func (*Figure) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Figures_mapString
}

func (*Figure) GongGetInstancesSet(stage *Stage) any {
	return &stage.Figures
}

func (*Figure) GongNewInstance() any {
	return new(Figure)
}

func (*FiguredBass) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.FiguredBasss, stage.FiguredBass_stagedOrder)
}

func (*FiguredBass) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.FiguredBass_orderStaged[order]
}

func (*FiguredBass) GongGetInstancesMapByName(stage *Stage) any {
	return stage.FiguredBasss_mapString
}

func (*FiguredBass) GongGetInstancesSet(stage *Stage) any {
	return &stage.FiguredBasss
}

func (*FiguredBass) GongNewInstance() any {
	return new(FiguredBass)
}

func (*Forward) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Forwards, stage.Forward_stagedOrder)
}

func (*Forward) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Forward_orderStaged[order]
}

func (*Forward) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Forwards_mapString
}

func (*Forward) GongGetInstancesSet(stage *Stage) any {
	return &stage.Forwards
}

func (*Forward) GongNewInstance() any {
	return new(Forward)
}

func (*Glissando) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Glissandos, stage.Glissando_stagedOrder)
}

func (*Glissando) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Glissando_orderStaged[order]
}

func (*Glissando) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Glissandos_mapString
}

func (*Glissando) GongGetInstancesSet(stage *Stage) any {
	return &stage.Glissandos
}

func (*Glissando) GongNewInstance() any {
	return new(Glissando)
}

func (*Grace) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Graces, stage.Grace_stagedOrder)
}

func (*Grace) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Grace_orderStaged[order]
}

func (*Grace) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Graces_mapString
}

func (*Grace) GongGetInstancesSet(stage *Stage) any {
	return &stage.Graces
}

func (*Grace) GongNewInstance() any {
	return new(Grace)
}

func (*HammerOnPullOff) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.HammerOnPullOffs, stage.HammerOnPullOff_stagedOrder)
}

func (*HammerOnPullOff) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.HammerOnPullOff_orderStaged[order]
}

func (*HammerOnPullOff) GongGetInstancesMapByName(stage *Stage) any {
	return stage.HammerOnPullOffs_mapString
}

func (*HammerOnPullOff) GongGetInstancesSet(stage *Stage) any {
	return &stage.HammerOnPullOffs
}

func (*HammerOnPullOff) GongNewInstance() any {
	return new(HammerOnPullOff)
}

func (*Handbell) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Handbells, stage.Handbell_stagedOrder)
}

func (*Handbell) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Handbell_orderStaged[order]
}

func (*Handbell) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Handbells_mapString
}

func (*Handbell) GongGetInstancesSet(stage *Stage) any {
	return &stage.Handbells
}

func (*Handbell) GongNewInstance() any {
	return new(Handbell)
}

func (*HarmonClosed) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.HarmonCloseds, stage.HarmonClosed_stagedOrder)
}

func (*HarmonClosed) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.HarmonClosed_orderStaged[order]
}

func (*HarmonClosed) GongGetInstancesMapByName(stage *Stage) any {
	return stage.HarmonCloseds_mapString
}

func (*HarmonClosed) GongGetInstancesSet(stage *Stage) any {
	return &stage.HarmonCloseds
}

func (*HarmonClosed) GongNewInstance() any {
	return new(HarmonClosed)
}

func (*HarmonMute) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.HarmonMutes, stage.HarmonMute_stagedOrder)
}

func (*HarmonMute) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.HarmonMute_orderStaged[order]
}

func (*HarmonMute) GongGetInstancesMapByName(stage *Stage) any {
	return stage.HarmonMutes_mapString
}

func (*HarmonMute) GongGetInstancesSet(stage *Stage) any {
	return &stage.HarmonMutes
}

func (*HarmonMute) GongNewInstance() any {
	return new(HarmonMute)
}

func (*Harmonic) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Harmonics, stage.Harmonic_stagedOrder)
}

func (*Harmonic) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Harmonic_orderStaged[order]
}

func (*Harmonic) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Harmonics_mapString
}

func (*Harmonic) GongGetInstancesSet(stage *Stage) any {
	return &stage.Harmonics
}

func (*Harmonic) GongNewInstance() any {
	return new(Harmonic)
}

func (*HeelToe) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.HeelToes, stage.HeelToe_stagedOrder)
}

func (*HeelToe) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.HeelToe_orderStaged[order]
}

func (*HeelToe) GongGetInstancesMapByName(stage *Stage) any {
	return stage.HeelToes_mapString
}

func (*HeelToe) GongGetInstancesSet(stage *Stage) any {
	return &stage.HeelToes
}

func (*HeelToe) GongNewInstance() any {
	return new(HeelToe)
}

func (*Hole) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Holes, stage.Hole_stagedOrder)
}

func (*Hole) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Hole_orderStaged[order]
}

func (*Hole) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Holes_mapString
}

func (*Hole) GongGetInstancesSet(stage *Stage) any {
	return &stage.Holes
}

func (*Hole) GongNewInstance() any {
	return new(Hole)
}

func (*HoleClosed) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.HoleCloseds, stage.HoleClosed_stagedOrder)
}

func (*HoleClosed) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.HoleClosed_orderStaged[order]
}

func (*HoleClosed) GongGetInstancesMapByName(stage *Stage) any {
	return stage.HoleCloseds_mapString
}

func (*HoleClosed) GongGetInstancesSet(stage *Stage) any {
	return &stage.HoleCloseds
}

func (*HoleClosed) GongNewInstance() any {
	return new(HoleClosed)
}

func (*Instrument) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Instruments, stage.Instrument_stagedOrder)
}

func (*Instrument) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Instrument_orderStaged[order]
}

func (*Instrument) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Instruments_mapString
}

func (*Instrument) GongGetInstancesSet(stage *Stage) any {
	return &stage.Instruments
}

func (*Instrument) GongNewInstance() any {
	return new(Instrument)
}

func (*Listen) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Listens, stage.Listen_stagedOrder)
}

func (*Listen) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Listen_orderStaged[order]
}

func (*Listen) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Listens_mapString
}

func (*Listen) GongGetInstancesSet(stage *Stage) any {
	return &stage.Listens
}

func (*Listen) GongNewInstance() any {
	return new(Listen)
}

func (*Lyric) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Lyrics, stage.Lyric_stagedOrder)
}

func (*Lyric) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Lyric_orderStaged[order]
}

func (*Lyric) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Lyrics_mapString
}

func (*Lyric) GongGetInstancesSet(stage *Stage) any {
	return &stage.Lyrics
}

func (*Lyric) GongNewInstance() any {
	return new(Lyric)
}

func (*Mordent) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Mordents, stage.Mordent_stagedOrder)
}

func (*Mordent) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Mordent_orderStaged[order]
}

func (*Mordent) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Mordents_mapString
}

func (*Mordent) GongGetInstancesSet(stage *Stage) any {
	return &stage.Mordents
}

func (*Mordent) GongNewInstance() any {
	return new(Mordent)
}

func (*NonArpeggiate) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NonArpeggiates, stage.NonArpeggiate_stagedOrder)
}

func (*NonArpeggiate) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NonArpeggiate_orderStaged[order]
}

func (*NonArpeggiate) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NonArpeggiates_mapString
}

func (*NonArpeggiate) GongGetInstancesSet(stage *Stage) any {
	return &stage.NonArpeggiates
}

func (*NonArpeggiate) GongNewInstance() any {
	return new(NonArpeggiate)
}

func (*Notations) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Notationss, stage.Notations_stagedOrder)
}

func (*Notations) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Notations_orderStaged[order]
}

func (*Notations) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Notationss_mapString
}

func (*Notations) GongGetInstancesSet(stage *Stage) any {
	return &stage.Notationss
}

func (*Notations) GongNewInstance() any {
	return new(Notations)
}

func (*Note) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Notes, stage.Note_stagedOrder)
}

func (*Note) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Note_orderStaged[order]
}

func (*Note) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Notes_mapString
}

func (*Note) GongGetInstancesSet(stage *Stage) any {
	return &stage.Notes
}

func (*Note) GongNewInstance() any {
	return new(Note)
}

func (*NoteType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NoteTypes, stage.NoteType_stagedOrder)
}

func (*NoteType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NoteType_orderStaged[order]
}

func (*NoteType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NoteTypes_mapString
}

func (*NoteType) GongGetInstancesSet(stage *Stage) any {
	return &stage.NoteTypes
}

func (*NoteType) GongNewInstance() any {
	return new(NoteType)
}

func (*Notehead) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Noteheads, stage.Notehead_stagedOrder)
}

func (*Notehead) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Notehead_orderStaged[order]
}

func (*Notehead) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Noteheads_mapString
}

func (*Notehead) GongGetInstancesSet(stage *Stage) any {
	return &stage.Noteheads
}

func (*Notehead) GongNewInstance() any {
	return new(Notehead)
}

func (*NoteheadText) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NoteheadTexts, stage.NoteheadText_stagedOrder)
}

func (*NoteheadText) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NoteheadText_orderStaged[order]
}

func (*NoteheadText) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NoteheadTexts_mapString
}

func (*NoteheadText) GongGetInstancesSet(stage *Stage) any {
	return &stage.NoteheadTexts
}

func (*NoteheadText) GongNewInstance() any {
	return new(NoteheadText)
}

func (*Ornaments) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Ornamentss, stage.Ornaments_stagedOrder)
}

func (*Ornaments) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Ornaments_orderStaged[order]
}

func (*Ornaments) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Ornamentss_mapString
}

func (*Ornaments) GongGetInstancesSet(stage *Stage) any {
	return &stage.Ornamentss
}

func (*Ornaments) GongNewInstance() any {
	return new(Ornaments)
}

func (*OtherNotation) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OtherNotations, stage.OtherNotation_stagedOrder)
}

func (*OtherNotation) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OtherNotation_orderStaged[order]
}

func (*OtherNotation) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OtherNotations_mapString
}

func (*OtherNotation) GongGetInstancesSet(stage *Stage) any {
	return &stage.OtherNotations
}

func (*OtherNotation) GongNewInstance() any {
	return new(OtherNotation)
}

func (*OtherPlacementText) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OtherPlacementTexts, stage.OtherPlacementText_stagedOrder)
}

func (*OtherPlacementText) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OtherPlacementText_orderStaged[order]
}

func (*OtherPlacementText) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OtherPlacementTexts_mapString
}

func (*OtherPlacementText) GongGetInstancesSet(stage *Stage) any {
	return &stage.OtherPlacementTexts
}

func (*OtherPlacementText) GongNewInstance() any {
	return new(OtherPlacementText)
}

func (*Pitch) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Pitchs, stage.Pitch_stagedOrder)
}

func (*Pitch) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Pitch_orderStaged[order]
}

func (*Pitch) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Pitchs_mapString
}

func (*Pitch) GongGetInstancesSet(stage *Stage) any {
	return &stage.Pitchs
}

func (*Pitch) GongNewInstance() any {
	return new(Pitch)
}

func (*PlacementText) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PlacementTexts, stage.PlacementText_stagedOrder)
}

func (*PlacementText) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PlacementText_orderStaged[order]
}

func (*PlacementText) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PlacementTexts_mapString
}

func (*PlacementText) GongGetInstancesSet(stage *Stage) any {
	return &stage.PlacementTexts
}

func (*PlacementText) GongNewInstance() any {
	return new(PlacementText)
}

func (*Release) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Releases, stage.Release_stagedOrder)
}

func (*Release) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Release_orderStaged[order]
}

func (*Release) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Releases_mapString
}

func (*Release) GongGetInstancesSet(stage *Stage) any {
	return &stage.Releases
}

func (*Release) GongNewInstance() any {
	return new(Release)
}

func (*Rest) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Rests, stage.Rest_stagedOrder)
}

func (*Rest) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Rest_orderStaged[order]
}

func (*Rest) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Rests_mapString
}

func (*Rest) GongGetInstancesSet(stage *Stage) any {
	return &stage.Rests
}

func (*Rest) GongNewInstance() any {
	return new(Rest)
}

func (*Slide) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Slides, stage.Slide_stagedOrder)
}

func (*Slide) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Slide_orderStaged[order]
}

func (*Slide) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Slides_mapString
}

func (*Slide) GongGetInstancesSet(stage *Stage) any {
	return &stage.Slides
}

func (*Slide) GongNewInstance() any {
	return new(Slide)
}

func (*Slur) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Slurs, stage.Slur_stagedOrder)
}

func (*Slur) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Slur_orderStaged[order]
}

func (*Slur) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Slurs_mapString
}

func (*Slur) GongGetInstancesSet(stage *Stage) any {
	return &stage.Slurs
}

func (*Slur) GongNewInstance() any {
	return new(Slur)
}

func (*Stem) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Stems, stage.Stem_stagedOrder)
}

func (*Stem) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Stem_orderStaged[order]
}

func (*Stem) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Stems_mapString
}

func (*Stem) GongGetInstancesSet(stage *Stage) any {
	return &stage.Stems
}

func (*Stem) GongNewInstance() any {
	return new(Stem)
}

func (*StrongAccent) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StrongAccents, stage.StrongAccent_stagedOrder)
}

func (*StrongAccent) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StrongAccent_orderStaged[order]
}

func (*StrongAccent) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StrongAccents_mapString
}

func (*StrongAccent) GongGetInstancesSet(stage *Stage) any {
	return &stage.StrongAccents
}

func (*StrongAccent) GongNewInstance() any {
	return new(StrongAccent)
}

func (*Tap) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Taps, stage.Tap_stagedOrder)
}

func (*Tap) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Tap_orderStaged[order]
}

func (*Tap) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Taps_mapString
}

func (*Tap) GongGetInstancesSet(stage *Stage) any {
	return &stage.Taps
}

func (*Tap) GongNewInstance() any {
	return new(Tap)
}

func (*Technical) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Technicals, stage.Technical_stagedOrder)
}

func (*Technical) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Technical_orderStaged[order]
}

func (*Technical) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Technicals_mapString
}

func (*Technical) GongGetInstancesSet(stage *Stage) any {
	return &stage.Technicals
}

func (*Technical) GongNewInstance() any {
	return new(Technical)
}

func (*TextElementData) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TextElementDatas, stage.TextElementData_stagedOrder)
}

func (*TextElementData) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TextElementData_orderStaged[order]
}

func (*TextElementData) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TextElementDatas_mapString
}

func (*TextElementData) GongGetInstancesSet(stage *Stage) any {
	return &stage.TextElementDatas
}

func (*TextElementData) GongNewInstance() any {
	return new(TextElementData)
}

func (*Tie) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Ties, stage.Tie_stagedOrder)
}

func (*Tie) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Tie_orderStaged[order]
}

func (*Tie) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Ties_mapString
}

func (*Tie) GongGetInstancesSet(stage *Stage) any {
	return &stage.Ties
}

func (*Tie) GongNewInstance() any {
	return new(Tie)
}

func (*Tied) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Tieds, stage.Tied_stagedOrder)
}

func (*Tied) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Tied_orderStaged[order]
}

func (*Tied) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Tieds_mapString
}

func (*Tied) GongGetInstancesSet(stage *Stage) any {
	return &stage.Tieds
}

func (*Tied) GongNewInstance() any {
	return new(Tied)
}

func (*Tremolo) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Tremolos, stage.Tremolo_stagedOrder)
}

func (*Tremolo) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Tremolo_orderStaged[order]
}

func (*Tremolo) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Tremolos_mapString
}

func (*Tremolo) GongGetInstancesSet(stage *Stage) any {
	return &stage.Tremolos
}

func (*Tremolo) GongNewInstance() any {
	return new(Tremolo)
}

func (*Tuplet) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Tuplets, stage.Tuplet_stagedOrder)
}

func (*Tuplet) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Tuplet_orderStaged[order]
}

func (*Tuplet) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Tuplets_mapString
}

func (*Tuplet) GongGetInstancesSet(stage *Stage) any {
	return &stage.Tuplets
}

func (*Tuplet) GongNewInstance() any {
	return new(Tuplet)
}

func (*TupletDot) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TupletDots, stage.TupletDot_stagedOrder)
}

func (*TupletDot) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TupletDot_orderStaged[order]
}

func (*TupletDot) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TupletDots_mapString
}

func (*TupletDot) GongGetInstancesSet(stage *Stage) any {
	return &stage.TupletDots
}

func (*TupletDot) GongNewInstance() any {
	return new(TupletDot)
}

func (*TupletNumber) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TupletNumbers, stage.TupletNumber_stagedOrder)
}

func (*TupletNumber) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TupletNumber_orderStaged[order]
}

func (*TupletNumber) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TupletNumbers_mapString
}

func (*TupletNumber) GongGetInstancesSet(stage *Stage) any {
	return &stage.TupletNumbers
}

func (*TupletNumber) GongNewInstance() any {
	return new(TupletNumber)
}

func (*TupletPortion) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TupletPortions, stage.TupletPortion_stagedOrder)
}

func (*TupletPortion) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TupletPortion_orderStaged[order]
}

func (*TupletPortion) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TupletPortions_mapString
}

func (*TupletPortion) GongGetInstancesSet(stage *Stage) any {
	return &stage.TupletPortions
}

func (*TupletPortion) GongNewInstance() any {
	return new(TupletPortion)
}

func (*TupletType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TupletTypes, stage.TupletType_stagedOrder)
}

func (*TupletType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TupletType_orderStaged[order]
}

func (*TupletType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TupletTypes_mapString
}

func (*TupletType) GongGetInstancesSet(stage *Stage) any {
	return &stage.TupletTypes
}

func (*TupletType) GongNewInstance() any {
	return new(TupletType)
}

func (*Unpitched) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Unpitcheds, stage.Unpitched_stagedOrder)
}

func (*Unpitched) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Unpitched_orderStaged[order]
}

func (*Unpitched) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Unpitcheds_mapString
}

func (*Unpitched) GongGetInstancesSet(stage *Stage) any {
	return &stage.Unpitcheds
}

func (*Unpitched) GongNewInstance() any {
	return new(Unpitched)
}

func (*Wait) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Waits, stage.Wait_stagedOrder)
}

func (*Wait) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Wait_orderStaged[order]
}

func (*Wait) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Waits_mapString
}

func (*Wait) GongGetInstancesSet(stage *Stage) any {
	return &stage.Waits
}

func (*Wait) GongNewInstance() any {
	return new(Wait)
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
	return "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
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
		Accidentals:           make(map[*Accidental]struct{}),
		Accidentals_mapString: make(map[string]*Accidental),

		AccidentalMarks:           make(map[*AccidentalMark]struct{}),
		AccidentalMarks_mapString: make(map[string]*AccidentalMark),

		Arpeggiates:           make(map[*Arpeggiate]struct{}),
		Arpeggiates_mapString: make(map[string]*Arpeggiate),

		Arrows:           make(map[*Arrow]struct{}),
		Arrows_mapString: make(map[string]*Arrow),

		Articulationss:           make(map[*Articulations]struct{}),
		Articulationss_mapString: make(map[string]*Articulations),

		Assesss:           make(map[*Assess]struct{}),
		Assesss_mapString: make(map[string]*Assess),

		Backups:           make(map[*Backup]struct{}),
		Backups_mapString: make(map[string]*Backup),

		Beams:           make(map[*Beam]struct{}),
		Beams_mapString: make(map[string]*Beam),

		Bends:           make(map[*Bend]struct{}),
		Bends_mapString: make(map[string]*Bend),

		BreathMarks:           make(map[*BreathMark]struct{}),
		BreathMarks_mapString: make(map[string]*BreathMark),

		Caesuras:           make(map[*Caesura]struct{}),
		Caesuras_mapString: make(map[string]*Caesura),

		Elisions:           make(map[*Elision]struct{}),
		Elisions_mapString: make(map[string]*Elision),

		EmptyLines:           make(map[*EmptyLine]struct{}),
		EmptyLines_mapString: make(map[string]*EmptyLine),

		Extends:           make(map[*Extend]struct{}),
		Extends_mapString: make(map[string]*Extend),

		Figures:           make(map[*Figure]struct{}),
		Figures_mapString: make(map[string]*Figure),

		FiguredBasss:           make(map[*FiguredBass]struct{}),
		FiguredBasss_mapString: make(map[string]*FiguredBass),

		Forwards:           make(map[*Forward]struct{}),
		Forwards_mapString: make(map[string]*Forward),

		Glissandos:           make(map[*Glissando]struct{}),
		Glissandos_mapString: make(map[string]*Glissando),

		Graces:           make(map[*Grace]struct{}),
		Graces_mapString: make(map[string]*Grace),

		HammerOnPullOffs:           make(map[*HammerOnPullOff]struct{}),
		HammerOnPullOffs_mapString: make(map[string]*HammerOnPullOff),

		Handbells:           make(map[*Handbell]struct{}),
		Handbells_mapString: make(map[string]*Handbell),

		HarmonCloseds:           make(map[*HarmonClosed]struct{}),
		HarmonCloseds_mapString: make(map[string]*HarmonClosed),

		HarmonMutes:           make(map[*HarmonMute]struct{}),
		HarmonMutes_mapString: make(map[string]*HarmonMute),

		Harmonics:           make(map[*Harmonic]struct{}),
		Harmonics_mapString: make(map[string]*Harmonic),

		HeelToes:           make(map[*HeelToe]struct{}),
		HeelToes_mapString: make(map[string]*HeelToe),

		Holes:           make(map[*Hole]struct{}),
		Holes_mapString: make(map[string]*Hole),

		HoleCloseds:           make(map[*HoleClosed]struct{}),
		HoleCloseds_mapString: make(map[string]*HoleClosed),

		Instruments:           make(map[*Instrument]struct{}),
		Instruments_mapString: make(map[string]*Instrument),

		Listens:           make(map[*Listen]struct{}),
		Listens_mapString: make(map[string]*Listen),

		Lyrics:           make(map[*Lyric]struct{}),
		Lyrics_mapString: make(map[string]*Lyric),

		Mordents:           make(map[*Mordent]struct{}),
		Mordents_mapString: make(map[string]*Mordent),

		NonArpeggiates:           make(map[*NonArpeggiate]struct{}),
		NonArpeggiates_mapString: make(map[string]*NonArpeggiate),

		Notationss:           make(map[*Notations]struct{}),
		Notationss_mapString: make(map[string]*Notations),

		Notes:           make(map[*Note]struct{}),
		Notes_mapString: make(map[string]*Note),

		NoteTypes:           make(map[*NoteType]struct{}),
		NoteTypes_mapString: make(map[string]*NoteType),

		Noteheads:           make(map[*Notehead]struct{}),
		Noteheads_mapString: make(map[string]*Notehead),

		NoteheadTexts:           make(map[*NoteheadText]struct{}),
		NoteheadTexts_mapString: make(map[string]*NoteheadText),

		Ornamentss:           make(map[*Ornaments]struct{}),
		Ornamentss_mapString: make(map[string]*Ornaments),

		OtherNotations:           make(map[*OtherNotation]struct{}),
		OtherNotations_mapString: make(map[string]*OtherNotation),

		OtherPlacementTexts:           make(map[*OtherPlacementText]struct{}),
		OtherPlacementTexts_mapString: make(map[string]*OtherPlacementText),

		Pitchs:           make(map[*Pitch]struct{}),
		Pitchs_mapString: make(map[string]*Pitch),

		PlacementTexts:           make(map[*PlacementText]struct{}),
		PlacementTexts_mapString: make(map[string]*PlacementText),

		Releases:           make(map[*Release]struct{}),
		Releases_mapString: make(map[string]*Release),

		Rests:           make(map[*Rest]struct{}),
		Rests_mapString: make(map[string]*Rest),

		Slides:           make(map[*Slide]struct{}),
		Slides_mapString: make(map[string]*Slide),

		Slurs:           make(map[*Slur]struct{}),
		Slurs_mapString: make(map[string]*Slur),

		Stems:           make(map[*Stem]struct{}),
		Stems_mapString: make(map[string]*Stem),

		StrongAccents:           make(map[*StrongAccent]struct{}),
		StrongAccents_mapString: make(map[string]*StrongAccent),

		Taps:           make(map[*Tap]struct{}),
		Taps_mapString: make(map[string]*Tap),

		Technicals:           make(map[*Technical]struct{}),
		Technicals_mapString: make(map[string]*Technical),

		TextElementDatas:           make(map[*TextElementData]struct{}),
		TextElementDatas_mapString: make(map[string]*TextElementData),

		Ties:           make(map[*Tie]struct{}),
		Ties_mapString: make(map[string]*Tie),

		Tieds:           make(map[*Tied]struct{}),
		Tieds_mapString: make(map[string]*Tied),

		Tremolos:           make(map[*Tremolo]struct{}),
		Tremolos_mapString: make(map[string]*Tremolo),

		Tuplets:           make(map[*Tuplet]struct{}),
		Tuplets_mapString: make(map[string]*Tuplet),

		TupletDots:           make(map[*TupletDot]struct{}),
		TupletDots_mapString: make(map[string]*TupletDot),

		TupletNumbers:           make(map[*TupletNumber]struct{}),
		TupletNumbers_mapString: make(map[string]*TupletNumber),

		TupletPortions:           make(map[*TupletPortion]struct{}),
		TupletPortions_mapString: make(map[string]*TupletPortion),

		TupletTypes:           make(map[*TupletType]struct{}),
		TupletTypes_mapString: make(map[string]*TupletType),

		Unpitcheds:           make(map[*Unpitched]struct{}),
		Unpitcheds_mapString: make(map[string]*Unpitched),

		Waits:           make(map[*Wait]struct{}),
		Waits_mapString: make(map[string]*Wait),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Accidental_stagedOrder: make(map[*Accidental]uint),
		Accidental_orderStaged: make(map[uint]*Accidental),
		Accidentals_reference:  make(map[*Accidental]*Accidental),

		AccidentalMark_stagedOrder: make(map[*AccidentalMark]uint),
		AccidentalMark_orderStaged: make(map[uint]*AccidentalMark),
		AccidentalMarks_reference:  make(map[*AccidentalMark]*AccidentalMark),

		Arpeggiate_stagedOrder: make(map[*Arpeggiate]uint),
		Arpeggiate_orderStaged: make(map[uint]*Arpeggiate),
		Arpeggiates_reference:  make(map[*Arpeggiate]*Arpeggiate),

		Arrow_stagedOrder: make(map[*Arrow]uint),
		Arrow_orderStaged: make(map[uint]*Arrow),
		Arrows_reference:  make(map[*Arrow]*Arrow),

		Articulations_stagedOrder: make(map[*Articulations]uint),
		Articulations_orderStaged: make(map[uint]*Articulations),
		Articulationss_reference:  make(map[*Articulations]*Articulations),

		Assess_stagedOrder: make(map[*Assess]uint),
		Assess_orderStaged: make(map[uint]*Assess),
		Assesss_reference:  make(map[*Assess]*Assess),

		Backup_stagedOrder: make(map[*Backup]uint),
		Backup_orderStaged: make(map[uint]*Backup),
		Backups_reference:  make(map[*Backup]*Backup),

		Beam_stagedOrder: make(map[*Beam]uint),
		Beam_orderStaged: make(map[uint]*Beam),
		Beams_reference:  make(map[*Beam]*Beam),

		Bend_stagedOrder: make(map[*Bend]uint),
		Bend_orderStaged: make(map[uint]*Bend),
		Bends_reference:  make(map[*Bend]*Bend),

		BreathMark_stagedOrder: make(map[*BreathMark]uint),
		BreathMark_orderStaged: make(map[uint]*BreathMark),
		BreathMarks_reference:  make(map[*BreathMark]*BreathMark),

		Caesura_stagedOrder: make(map[*Caesura]uint),
		Caesura_orderStaged: make(map[uint]*Caesura),
		Caesuras_reference:  make(map[*Caesura]*Caesura),

		Elision_stagedOrder: make(map[*Elision]uint),
		Elision_orderStaged: make(map[uint]*Elision),
		Elisions_reference:  make(map[*Elision]*Elision),

		EmptyLine_stagedOrder: make(map[*EmptyLine]uint),
		EmptyLine_orderStaged: make(map[uint]*EmptyLine),
		EmptyLines_reference:  make(map[*EmptyLine]*EmptyLine),

		Extend_stagedOrder: make(map[*Extend]uint),
		Extend_orderStaged: make(map[uint]*Extend),
		Extends_reference:  make(map[*Extend]*Extend),

		Figure_stagedOrder: make(map[*Figure]uint),
		Figure_orderStaged: make(map[uint]*Figure),
		Figures_reference:  make(map[*Figure]*Figure),

		FiguredBass_stagedOrder: make(map[*FiguredBass]uint),
		FiguredBass_orderStaged: make(map[uint]*FiguredBass),
		FiguredBasss_reference:  make(map[*FiguredBass]*FiguredBass),

		Forward_stagedOrder: make(map[*Forward]uint),
		Forward_orderStaged: make(map[uint]*Forward),
		Forwards_reference:  make(map[*Forward]*Forward),

		Glissando_stagedOrder: make(map[*Glissando]uint),
		Glissando_orderStaged: make(map[uint]*Glissando),
		Glissandos_reference:  make(map[*Glissando]*Glissando),

		Grace_stagedOrder: make(map[*Grace]uint),
		Grace_orderStaged: make(map[uint]*Grace),
		Graces_reference:  make(map[*Grace]*Grace),

		HammerOnPullOff_stagedOrder: make(map[*HammerOnPullOff]uint),
		HammerOnPullOff_orderStaged: make(map[uint]*HammerOnPullOff),
		HammerOnPullOffs_reference:  make(map[*HammerOnPullOff]*HammerOnPullOff),

		Handbell_stagedOrder: make(map[*Handbell]uint),
		Handbell_orderStaged: make(map[uint]*Handbell),
		Handbells_reference:  make(map[*Handbell]*Handbell),

		HarmonClosed_stagedOrder: make(map[*HarmonClosed]uint),
		HarmonClosed_orderStaged: make(map[uint]*HarmonClosed),
		HarmonCloseds_reference:  make(map[*HarmonClosed]*HarmonClosed),

		HarmonMute_stagedOrder: make(map[*HarmonMute]uint),
		HarmonMute_orderStaged: make(map[uint]*HarmonMute),
		HarmonMutes_reference:  make(map[*HarmonMute]*HarmonMute),

		Harmonic_stagedOrder: make(map[*Harmonic]uint),
		Harmonic_orderStaged: make(map[uint]*Harmonic),
		Harmonics_reference:  make(map[*Harmonic]*Harmonic),

		HeelToe_stagedOrder: make(map[*HeelToe]uint),
		HeelToe_orderStaged: make(map[uint]*HeelToe),
		HeelToes_reference:  make(map[*HeelToe]*HeelToe),

		Hole_stagedOrder: make(map[*Hole]uint),
		Hole_orderStaged: make(map[uint]*Hole),
		Holes_reference:  make(map[*Hole]*Hole),

		HoleClosed_stagedOrder: make(map[*HoleClosed]uint),
		HoleClosed_orderStaged: make(map[uint]*HoleClosed),
		HoleCloseds_reference:  make(map[*HoleClosed]*HoleClosed),

		Instrument_stagedOrder: make(map[*Instrument]uint),
		Instrument_orderStaged: make(map[uint]*Instrument),
		Instruments_reference:  make(map[*Instrument]*Instrument),

		Listen_stagedOrder: make(map[*Listen]uint),
		Listen_orderStaged: make(map[uint]*Listen),
		Listens_reference:  make(map[*Listen]*Listen),

		Lyric_stagedOrder: make(map[*Lyric]uint),
		Lyric_orderStaged: make(map[uint]*Lyric),
		Lyrics_reference:  make(map[*Lyric]*Lyric),

		Mordent_stagedOrder: make(map[*Mordent]uint),
		Mordent_orderStaged: make(map[uint]*Mordent),
		Mordents_reference:  make(map[*Mordent]*Mordent),

		NonArpeggiate_stagedOrder: make(map[*NonArpeggiate]uint),
		NonArpeggiate_orderStaged: make(map[uint]*NonArpeggiate),
		NonArpeggiates_reference:  make(map[*NonArpeggiate]*NonArpeggiate),

		Notations_stagedOrder: make(map[*Notations]uint),
		Notations_orderStaged: make(map[uint]*Notations),
		Notationss_reference:  make(map[*Notations]*Notations),

		Note_stagedOrder: make(map[*Note]uint),
		Note_orderStaged: make(map[uint]*Note),
		Notes_reference:  make(map[*Note]*Note),

		NoteType_stagedOrder: make(map[*NoteType]uint),
		NoteType_orderStaged: make(map[uint]*NoteType),
		NoteTypes_reference:  make(map[*NoteType]*NoteType),

		Notehead_stagedOrder: make(map[*Notehead]uint),
		Notehead_orderStaged: make(map[uint]*Notehead),
		Noteheads_reference:  make(map[*Notehead]*Notehead),

		NoteheadText_stagedOrder: make(map[*NoteheadText]uint),
		NoteheadText_orderStaged: make(map[uint]*NoteheadText),
		NoteheadTexts_reference:  make(map[*NoteheadText]*NoteheadText),

		Ornaments_stagedOrder: make(map[*Ornaments]uint),
		Ornaments_orderStaged: make(map[uint]*Ornaments),
		Ornamentss_reference:  make(map[*Ornaments]*Ornaments),

		OtherNotation_stagedOrder: make(map[*OtherNotation]uint),
		OtherNotation_orderStaged: make(map[uint]*OtherNotation),
		OtherNotations_reference:  make(map[*OtherNotation]*OtherNotation),

		OtherPlacementText_stagedOrder: make(map[*OtherPlacementText]uint),
		OtherPlacementText_orderStaged: make(map[uint]*OtherPlacementText),
		OtherPlacementTexts_reference:  make(map[*OtherPlacementText]*OtherPlacementText),

		Pitch_stagedOrder: make(map[*Pitch]uint),
		Pitch_orderStaged: make(map[uint]*Pitch),
		Pitchs_reference:  make(map[*Pitch]*Pitch),

		PlacementText_stagedOrder: make(map[*PlacementText]uint),
		PlacementText_orderStaged: make(map[uint]*PlacementText),
		PlacementTexts_reference:  make(map[*PlacementText]*PlacementText),

		Release_stagedOrder: make(map[*Release]uint),
		Release_orderStaged: make(map[uint]*Release),
		Releases_reference:  make(map[*Release]*Release),

		Rest_stagedOrder: make(map[*Rest]uint),
		Rest_orderStaged: make(map[uint]*Rest),
		Rests_reference:  make(map[*Rest]*Rest),

		Slide_stagedOrder: make(map[*Slide]uint),
		Slide_orderStaged: make(map[uint]*Slide),
		Slides_reference:  make(map[*Slide]*Slide),

		Slur_stagedOrder: make(map[*Slur]uint),
		Slur_orderStaged: make(map[uint]*Slur),
		Slurs_reference:  make(map[*Slur]*Slur),

		Stem_stagedOrder: make(map[*Stem]uint),
		Stem_orderStaged: make(map[uint]*Stem),
		Stems_reference:  make(map[*Stem]*Stem),

		StrongAccent_stagedOrder: make(map[*StrongAccent]uint),
		StrongAccent_orderStaged: make(map[uint]*StrongAccent),
		StrongAccents_reference:  make(map[*StrongAccent]*StrongAccent),

		Tap_stagedOrder: make(map[*Tap]uint),
		Tap_orderStaged: make(map[uint]*Tap),
		Taps_reference:  make(map[*Tap]*Tap),

		Technical_stagedOrder: make(map[*Technical]uint),
		Technical_orderStaged: make(map[uint]*Technical),
		Technicals_reference:  make(map[*Technical]*Technical),

		TextElementData_stagedOrder: make(map[*TextElementData]uint),
		TextElementData_orderStaged: make(map[uint]*TextElementData),
		TextElementDatas_reference:  make(map[*TextElementData]*TextElementData),

		Tie_stagedOrder: make(map[*Tie]uint),
		Tie_orderStaged: make(map[uint]*Tie),
		Ties_reference:  make(map[*Tie]*Tie),

		Tied_stagedOrder: make(map[*Tied]uint),
		Tied_orderStaged: make(map[uint]*Tied),
		Tieds_reference:  make(map[*Tied]*Tied),

		Tremolo_stagedOrder: make(map[*Tremolo]uint),
		Tremolo_orderStaged: make(map[uint]*Tremolo),
		Tremolos_reference:  make(map[*Tremolo]*Tremolo),

		Tuplet_stagedOrder: make(map[*Tuplet]uint),
		Tuplet_orderStaged: make(map[uint]*Tuplet),
		Tuplets_reference:  make(map[*Tuplet]*Tuplet),

		TupletDot_stagedOrder: make(map[*TupletDot]uint),
		TupletDot_orderStaged: make(map[uint]*TupletDot),
		TupletDots_reference:  make(map[*TupletDot]*TupletDot),

		TupletNumber_stagedOrder: make(map[*TupletNumber]uint),
		TupletNumber_orderStaged: make(map[uint]*TupletNumber),
		TupletNumbers_reference:  make(map[*TupletNumber]*TupletNumber),

		TupletPortion_stagedOrder: make(map[*TupletPortion]uint),
		TupletPortion_orderStaged: make(map[uint]*TupletPortion),
		TupletPortions_reference:  make(map[*TupletPortion]*TupletPortion),

		TupletType_stagedOrder: make(map[*TupletType]uint),
		TupletType_orderStaged: make(map[uint]*TupletType),
		TupletTypes_reference:  make(map[*TupletType]*TupletType),

		Unpitched_stagedOrder: make(map[*Unpitched]uint),
		Unpitched_orderStaged: make(map[uint]*Unpitched),
		Unpitcheds_reference:  make(map[*Unpitched]*Unpitched),

		Wait_stagedOrder: make(map[*Wait]uint),
		Wait_orderStaged: make(map[uint]*Wait),
		Waits_reference:  make(map[*Wait]*Wait),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Accidental": &AccidentalUnmarshaller{},

			"AccidentalMark": &AccidentalMarkUnmarshaller{},

			"Arpeggiate": &ArpeggiateUnmarshaller{},

			"Arrow": &ArrowUnmarshaller{},

			"Articulations": &ArticulationsUnmarshaller{},

			"Assess": &AssessUnmarshaller{},

			"Backup": &BackupUnmarshaller{},

			"Beam": &BeamUnmarshaller{},

			"Bend": &BendUnmarshaller{},

			"BreathMark": &BreathMarkUnmarshaller{},

			"Caesura": &CaesuraUnmarshaller{},

			"Elision": &ElisionUnmarshaller{},

			"EmptyLine": &EmptyLineUnmarshaller{},

			"Extend": &ExtendUnmarshaller{},

			"Figure": &FigureUnmarshaller{},

			"FiguredBass": &FiguredBassUnmarshaller{},

			"Forward": &ForwardUnmarshaller{},

			"Glissando": &GlissandoUnmarshaller{},

			"Grace": &GraceUnmarshaller{},

			"HammerOnPullOff": &HammerOnPullOffUnmarshaller{},

			"Handbell": &HandbellUnmarshaller{},

			"HarmonClosed": &HarmonClosedUnmarshaller{},

			"HarmonMute": &HarmonMuteUnmarshaller{},

			"Harmonic": &HarmonicUnmarshaller{},

			"HeelToe": &HeelToeUnmarshaller{},

			"Hole": &HoleUnmarshaller{},

			"HoleClosed": &HoleClosedUnmarshaller{},

			"Instrument": &InstrumentUnmarshaller{},

			"Listen": &ListenUnmarshaller{},

			"Lyric": &LyricUnmarshaller{},

			"Mordent": &MordentUnmarshaller{},

			"NonArpeggiate": &NonArpeggiateUnmarshaller{},

			"Notations": &NotationsUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteType": &NoteTypeUnmarshaller{},

			"Notehead": &NoteheadUnmarshaller{},

			"NoteheadText": &NoteheadTextUnmarshaller{},

			"Ornaments": &OrnamentsUnmarshaller{},

			"OtherNotation": &OtherNotationUnmarshaller{},

			"OtherPlacementText": &OtherPlacementTextUnmarshaller{},

			"Pitch": &PitchUnmarshaller{},

			"PlacementText": &PlacementTextUnmarshaller{},

			"Release": &ReleaseUnmarshaller{},

			"Rest": &RestUnmarshaller{},

			"Slide": &SlideUnmarshaller{},

			"Slur": &SlurUnmarshaller{},

			"Stem": &StemUnmarshaller{},

			"StrongAccent": &StrongAccentUnmarshaller{},

			"Tap": &TapUnmarshaller{},

			"Technical": &TechnicalUnmarshaller{},

			"TextElementData": &TextElementDataUnmarshaller{},

			"Tie": &TieUnmarshaller{},

			"Tied": &TiedUnmarshaller{},

			"Tremolo": &TremoloUnmarshaller{},

			"Tuplet": &TupletUnmarshaller{},

			"TupletDot": &TupletDotUnmarshaller{},

			"TupletNumber": &TupletNumberUnmarshaller{},

			"TupletPortion": &TupletPortionUnmarshaller{},

			"TupletType": &TupletTypeUnmarshaller{},

			"Unpitched": &UnpitchedUnmarshaller{},

			"Wait": &WaitUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["Accidental"] = len(stage.Accidentals)
	stage.Map_GongStructName_InstancesNb["AccidentalMark"] = len(stage.AccidentalMarks)
	stage.Map_GongStructName_InstancesNb["Arpeggiate"] = len(stage.Arpeggiates)
	stage.Map_GongStructName_InstancesNb["Arrow"] = len(stage.Arrows)
	stage.Map_GongStructName_InstancesNb["Articulations"] = len(stage.Articulationss)
	stage.Map_GongStructName_InstancesNb["Assess"] = len(stage.Assesss)
	stage.Map_GongStructName_InstancesNb["Backup"] = len(stage.Backups)
	stage.Map_GongStructName_InstancesNb["Beam"] = len(stage.Beams)
	stage.Map_GongStructName_InstancesNb["Bend"] = len(stage.Bends)
	stage.Map_GongStructName_InstancesNb["BreathMark"] = len(stage.BreathMarks)
	stage.Map_GongStructName_InstancesNb["Caesura"] = len(stage.Caesuras)
	stage.Map_GongStructName_InstancesNb["Elision"] = len(stage.Elisions)
	stage.Map_GongStructName_InstancesNb["EmptyLine"] = len(stage.EmptyLines)
	stage.Map_GongStructName_InstancesNb["Extend"] = len(stage.Extends)
	stage.Map_GongStructName_InstancesNb["Figure"] = len(stage.Figures)
	stage.Map_GongStructName_InstancesNb["FiguredBass"] = len(stage.FiguredBasss)
	stage.Map_GongStructName_InstancesNb["Forward"] = len(stage.Forwards)
	stage.Map_GongStructName_InstancesNb["Glissando"] = len(stage.Glissandos)
	stage.Map_GongStructName_InstancesNb["Grace"] = len(stage.Graces)
	stage.Map_GongStructName_InstancesNb["HammerOnPullOff"] = len(stage.HammerOnPullOffs)
	stage.Map_GongStructName_InstancesNb["Handbell"] = len(stage.Handbells)
	stage.Map_GongStructName_InstancesNb["HarmonClosed"] = len(stage.HarmonCloseds)
	stage.Map_GongStructName_InstancesNb["HarmonMute"] = len(stage.HarmonMutes)
	stage.Map_GongStructName_InstancesNb["Harmonic"] = len(stage.Harmonics)
	stage.Map_GongStructName_InstancesNb["HeelToe"] = len(stage.HeelToes)
	stage.Map_GongStructName_InstancesNb["Hole"] = len(stage.Holes)
	stage.Map_GongStructName_InstancesNb["HoleClosed"] = len(stage.HoleCloseds)
	stage.Map_GongStructName_InstancesNb["Instrument"] = len(stage.Instruments)
	stage.Map_GongStructName_InstancesNb["Listen"] = len(stage.Listens)
	stage.Map_GongStructName_InstancesNb["Lyric"] = len(stage.Lyrics)
	stage.Map_GongStructName_InstancesNb["Mordent"] = len(stage.Mordents)
	stage.Map_GongStructName_InstancesNb["NonArpeggiate"] = len(stage.NonArpeggiates)
	stage.Map_GongStructName_InstancesNb["Notations"] = len(stage.Notationss)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteType"] = len(stage.NoteTypes)
	stage.Map_GongStructName_InstancesNb["Notehead"] = len(stage.Noteheads)
	stage.Map_GongStructName_InstancesNb["NoteheadText"] = len(stage.NoteheadTexts)
	stage.Map_GongStructName_InstancesNb["Ornaments"] = len(stage.Ornamentss)
	stage.Map_GongStructName_InstancesNb["OtherNotation"] = len(stage.OtherNotations)
	stage.Map_GongStructName_InstancesNb["OtherPlacementText"] = len(stage.OtherPlacementTexts)
	stage.Map_GongStructName_InstancesNb["Pitch"] = len(stage.Pitchs)
	stage.Map_GongStructName_InstancesNb["PlacementText"] = len(stage.PlacementTexts)
	stage.Map_GongStructName_InstancesNb["Release"] = len(stage.Releases)
	stage.Map_GongStructName_InstancesNb["Rest"] = len(stage.Rests)
	stage.Map_GongStructName_InstancesNb["Slide"] = len(stage.Slides)
	stage.Map_GongStructName_InstancesNb["Slur"] = len(stage.Slurs)
	stage.Map_GongStructName_InstancesNb["Stem"] = len(stage.Stems)
	stage.Map_GongStructName_InstancesNb["StrongAccent"] = len(stage.StrongAccents)
	stage.Map_GongStructName_InstancesNb["Tap"] = len(stage.Taps)
	stage.Map_GongStructName_InstancesNb["Technical"] = len(stage.Technicals)
	stage.Map_GongStructName_InstancesNb["TextElementData"] = len(stage.TextElementDatas)
	stage.Map_GongStructName_InstancesNb["Tie"] = len(stage.Ties)
	stage.Map_GongStructName_InstancesNb["Tied"] = len(stage.Tieds)
	stage.Map_GongStructName_InstancesNb["Tremolo"] = len(stage.Tremolos)
	stage.Map_GongStructName_InstancesNb["Tuplet"] = len(stage.Tuplets)
	stage.Map_GongStructName_InstancesNb["TupletDot"] = len(stage.TupletDots)
	stage.Map_GongStructName_InstancesNb["TupletNumber"] = len(stage.TupletNumbers)
	stage.Map_GongStructName_InstancesNb["TupletPortion"] = len(stage.TupletPortions)
	stage.Map_GongStructName_InstancesNb["TupletType"] = len(stage.TupletTypes)
	stage.Map_GongStructName_InstancesNb["Unpitched"] = len(stage.Unpitcheds)
	stage.Map_GongStructName_InstancesNb["Wait"] = len(stage.Waits)
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
// Stage puts accidental to the model stage
func (accidental *Accidental) Stage(stage *Stage) *Accidental {
	__gong__stage(stage.Accidentals, stage.Accidental_stagedOrder, stage.Accidental_orderStaged, &stage.AccidentalOrder, stage.Accidentals_mapString, accidental, accidental.Name)
	return accidental
}

// StagePreserveOrder puts accidental to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AccidentalOrder
// - update stage.AccidentalOrder accordingly
func (accidental *Accidental) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Accidentals, stage.Accidental_stagedOrder, stage.Accidental_orderStaged, &stage.AccidentalOrder, stage.Accidentals_mapString, accidental, order, accidental.Name)
}

// Unstage removes accidental off the model stage
func (accidental *Accidental) Unstage(stage *Stage) *Accidental {
	__gong__unstage(stage.Accidentals, stage.Accidentals_mapString, accidental, accidental.Name)
	return accidental
}

// UnstageVoid removes accidental off the model stage
func (accidental *Accidental) UnstageVoid(stage *Stage) {
	accidental.Unstage(stage)
}

func (accidental *Accidental) StageVoid(stage *Stage) {
	accidental.Stage(stage)
}

// for satisfaction of GongStruct interface
func (accidental *Accidental) GetName() (res string) {
	return accidental.Name
}

// for satisfaction of GongStruct interface
func (accidental *Accidental) SetName(name string) {
	accidental.Name = name
}

// Stage puts accidentalmark to the model stage
func (accidentalmark *AccidentalMark) Stage(stage *Stage) *AccidentalMark {
	__gong__stage(stage.AccidentalMarks, stage.AccidentalMark_stagedOrder, stage.AccidentalMark_orderStaged, &stage.AccidentalMarkOrder, stage.AccidentalMarks_mapString, accidentalmark, accidentalmark.Name)
	return accidentalmark
}

// StagePreserveOrder puts accidentalmark to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AccidentalMarkOrder
// - update stage.AccidentalMarkOrder accordingly
func (accidentalmark *AccidentalMark) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AccidentalMarks, stage.AccidentalMark_stagedOrder, stage.AccidentalMark_orderStaged, &stage.AccidentalMarkOrder, stage.AccidentalMarks_mapString, accidentalmark, order, accidentalmark.Name)
}

// Unstage removes accidentalmark off the model stage
func (accidentalmark *AccidentalMark) Unstage(stage *Stage) *AccidentalMark {
	__gong__unstage(stage.AccidentalMarks, stage.AccidentalMarks_mapString, accidentalmark, accidentalmark.Name)
	return accidentalmark
}

// UnstageVoid removes accidentalmark off the model stage
func (accidentalmark *AccidentalMark) UnstageVoid(stage *Stage) {
	accidentalmark.Unstage(stage)
}

func (accidentalmark *AccidentalMark) StageVoid(stage *Stage) {
	accidentalmark.Stage(stage)
}

// for satisfaction of GongStruct interface
func (accidentalmark *AccidentalMark) GetName() (res string) {
	return accidentalmark.Name
}

// for satisfaction of GongStruct interface
func (accidentalmark *AccidentalMark) SetName(name string) {
	accidentalmark.Name = name
}

// Stage puts arpeggiate to the model stage
func (arpeggiate *Arpeggiate) Stage(stage *Stage) *Arpeggiate {
	__gong__stage(stage.Arpeggiates, stage.Arpeggiate_stagedOrder, stage.Arpeggiate_orderStaged, &stage.ArpeggiateOrder, stage.Arpeggiates_mapString, arpeggiate, arpeggiate.Name)
	return arpeggiate
}

// StagePreserveOrder puts arpeggiate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArpeggiateOrder
// - update stage.ArpeggiateOrder accordingly
func (arpeggiate *Arpeggiate) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Arpeggiates, stage.Arpeggiate_stagedOrder, stage.Arpeggiate_orderStaged, &stage.ArpeggiateOrder, stage.Arpeggiates_mapString, arpeggiate, order, arpeggiate.Name)
}

// Unstage removes arpeggiate off the model stage
func (arpeggiate *Arpeggiate) Unstage(stage *Stage) *Arpeggiate {
	__gong__unstage(stage.Arpeggiates, stage.Arpeggiates_mapString, arpeggiate, arpeggiate.Name)
	return arpeggiate
}

// UnstageVoid removes arpeggiate off the model stage
func (arpeggiate *Arpeggiate) UnstageVoid(stage *Stage) {
	arpeggiate.Unstage(stage)
}

func (arpeggiate *Arpeggiate) StageVoid(stage *Stage) {
	arpeggiate.Stage(stage)
}

// for satisfaction of GongStruct interface
func (arpeggiate *Arpeggiate) GetName() (res string) {
	return arpeggiate.Name
}

// for satisfaction of GongStruct interface
func (arpeggiate *Arpeggiate) SetName(name string) {
	arpeggiate.Name = name
}

// Stage puts arrow to the model stage
func (arrow *Arrow) Stage(stage *Stage) *Arrow {
	__gong__stage(stage.Arrows, stage.Arrow_stagedOrder, stage.Arrow_orderStaged, &stage.ArrowOrder, stage.Arrows_mapString, arrow, arrow.Name)
	return arrow
}

// StagePreserveOrder puts arrow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArrowOrder
// - update stage.ArrowOrder accordingly
func (arrow *Arrow) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Arrows, stage.Arrow_stagedOrder, stage.Arrow_orderStaged, &stage.ArrowOrder, stage.Arrows_mapString, arrow, order, arrow.Name)
}

// Unstage removes arrow off the model stage
func (arrow *Arrow) Unstage(stage *Stage) *Arrow {
	__gong__unstage(stage.Arrows, stage.Arrows_mapString, arrow, arrow.Name)
	return arrow
}

// UnstageVoid removes arrow off the model stage
func (arrow *Arrow) UnstageVoid(stage *Stage) {
	arrow.Unstage(stage)
}

func (arrow *Arrow) StageVoid(stage *Stage) {
	arrow.Stage(stage)
}

// for satisfaction of GongStruct interface
func (arrow *Arrow) GetName() (res string) {
	return arrow.Name
}

// for satisfaction of GongStruct interface
func (arrow *Arrow) SetName(name string) {
	arrow.Name = name
}

// Stage puts articulations to the model stage
func (articulations *Articulations) Stage(stage *Stage) *Articulations {
	__gong__stage(stage.Articulationss, stage.Articulations_stagedOrder, stage.Articulations_orderStaged, &stage.ArticulationsOrder, stage.Articulationss_mapString, articulations, articulations.Name)
	return articulations
}

// StagePreserveOrder puts articulations to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArticulationsOrder
// - update stage.ArticulationsOrder accordingly
func (articulations *Articulations) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Articulationss, stage.Articulations_stagedOrder, stage.Articulations_orderStaged, &stage.ArticulationsOrder, stage.Articulationss_mapString, articulations, order, articulations.Name)
}

// Unstage removes articulations off the model stage
func (articulations *Articulations) Unstage(stage *Stage) *Articulations {
	__gong__unstage(stage.Articulationss, stage.Articulationss_mapString, articulations, articulations.Name)
	return articulations
}

// UnstageVoid removes articulations off the model stage
func (articulations *Articulations) UnstageVoid(stage *Stage) {
	articulations.Unstage(stage)
}

func (articulations *Articulations) StageVoid(stage *Stage) {
	articulations.Stage(stage)
}

// for satisfaction of GongStruct interface
func (articulations *Articulations) GetName() (res string) {
	return articulations.Name
}

// for satisfaction of GongStruct interface
func (articulations *Articulations) SetName(name string) {
	articulations.Name = name
}

// Stage puts assess to the model stage
func (assess *Assess) Stage(stage *Stage) *Assess {
	__gong__stage(stage.Assesss, stage.Assess_stagedOrder, stage.Assess_orderStaged, &stage.AssessOrder, stage.Assesss_mapString, assess, assess.Name)
	return assess
}

// StagePreserveOrder puts assess to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AssessOrder
// - update stage.AssessOrder accordingly
func (assess *Assess) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Assesss, stage.Assess_stagedOrder, stage.Assess_orderStaged, &stage.AssessOrder, stage.Assesss_mapString, assess, order, assess.Name)
}

// Unstage removes assess off the model stage
func (assess *Assess) Unstage(stage *Stage) *Assess {
	__gong__unstage(stage.Assesss, stage.Assesss_mapString, assess, assess.Name)
	return assess
}

// UnstageVoid removes assess off the model stage
func (assess *Assess) UnstageVoid(stage *Stage) {
	assess.Unstage(stage)
}

func (assess *Assess) StageVoid(stage *Stage) {
	assess.Stage(stage)
}

// for satisfaction of GongStruct interface
func (assess *Assess) GetName() (res string) {
	return assess.Name
}

// for satisfaction of GongStruct interface
func (assess *Assess) SetName(name string) {
	assess.Name = name
}

// Stage puts backup to the model stage
func (backup *Backup) Stage(stage *Stage) *Backup {
	__gong__stage(stage.Backups, stage.Backup_stagedOrder, stage.Backup_orderStaged, &stage.BackupOrder, stage.Backups_mapString, backup, backup.Name)
	return backup
}

// StagePreserveOrder puts backup to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BackupOrder
// - update stage.BackupOrder accordingly
func (backup *Backup) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Backups, stage.Backup_stagedOrder, stage.Backup_orderStaged, &stage.BackupOrder, stage.Backups_mapString, backup, order, backup.Name)
}

// Unstage removes backup off the model stage
func (backup *Backup) Unstage(stage *Stage) *Backup {
	__gong__unstage(stage.Backups, stage.Backups_mapString, backup, backup.Name)
	return backup
}

// UnstageVoid removes backup off the model stage
func (backup *Backup) UnstageVoid(stage *Stage) {
	backup.Unstage(stage)
}

func (backup *Backup) StageVoid(stage *Stage) {
	backup.Stage(stage)
}

// for satisfaction of GongStruct interface
func (backup *Backup) GetName() (res string) {
	return backup.Name
}

// for satisfaction of GongStruct interface
func (backup *Backup) SetName(name string) {
	backup.Name = name
}

// Stage puts beam to the model stage
func (beam *Beam) Stage(stage *Stage) *Beam {
	__gong__stage(stage.Beams, stage.Beam_stagedOrder, stage.Beam_orderStaged, &stage.BeamOrder, stage.Beams_mapString, beam, beam.Name)
	return beam
}

// StagePreserveOrder puts beam to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BeamOrder
// - update stage.BeamOrder accordingly
func (beam *Beam) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Beams, stage.Beam_stagedOrder, stage.Beam_orderStaged, &stage.BeamOrder, stage.Beams_mapString, beam, order, beam.Name)
}

// Unstage removes beam off the model stage
func (beam *Beam) Unstage(stage *Stage) *Beam {
	__gong__unstage(stage.Beams, stage.Beams_mapString, beam, beam.Name)
	return beam
}

// UnstageVoid removes beam off the model stage
func (beam *Beam) UnstageVoid(stage *Stage) {
	beam.Unstage(stage)
}

func (beam *Beam) StageVoid(stage *Stage) {
	beam.Stage(stage)
}

// for satisfaction of GongStruct interface
func (beam *Beam) GetName() (res string) {
	return beam.Name
}

// for satisfaction of GongStruct interface
func (beam *Beam) SetName(name string) {
	beam.Name = name
}

// Stage puts bend to the model stage
func (bend *Bend) Stage(stage *Stage) *Bend {
	__gong__stage(stage.Bends, stage.Bend_stagedOrder, stage.Bend_orderStaged, &stage.BendOrder, stage.Bends_mapString, bend, bend.Name)
	return bend
}

// StagePreserveOrder puts bend to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BendOrder
// - update stage.BendOrder accordingly
func (bend *Bend) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Bends, stage.Bend_stagedOrder, stage.Bend_orderStaged, &stage.BendOrder, stage.Bends_mapString, bend, order, bend.Name)
}

// Unstage removes bend off the model stage
func (bend *Bend) Unstage(stage *Stage) *Bend {
	__gong__unstage(stage.Bends, stage.Bends_mapString, bend, bend.Name)
	return bend
}

// UnstageVoid removes bend off the model stage
func (bend *Bend) UnstageVoid(stage *Stage) {
	bend.Unstage(stage)
}

func (bend *Bend) StageVoid(stage *Stage) {
	bend.Stage(stage)
}

// for satisfaction of GongStruct interface
func (bend *Bend) GetName() (res string) {
	return bend.Name
}

// for satisfaction of GongStruct interface
func (bend *Bend) SetName(name string) {
	bend.Name = name
}

// Stage puts breathmark to the model stage
func (breathmark *BreathMark) Stage(stage *Stage) *BreathMark {
	__gong__stage(stage.BreathMarks, stage.BreathMark_stagedOrder, stage.BreathMark_orderStaged, &stage.BreathMarkOrder, stage.BreathMarks_mapString, breathmark, breathmark.Name)
	return breathmark
}

// StagePreserveOrder puts breathmark to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BreathMarkOrder
// - update stage.BreathMarkOrder accordingly
func (breathmark *BreathMark) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.BreathMarks, stage.BreathMark_stagedOrder, stage.BreathMark_orderStaged, &stage.BreathMarkOrder, stage.BreathMarks_mapString, breathmark, order, breathmark.Name)
}

// Unstage removes breathmark off the model stage
func (breathmark *BreathMark) Unstage(stage *Stage) *BreathMark {
	__gong__unstage(stage.BreathMarks, stage.BreathMarks_mapString, breathmark, breathmark.Name)
	return breathmark
}

// UnstageVoid removes breathmark off the model stage
func (breathmark *BreathMark) UnstageVoid(stage *Stage) {
	breathmark.Unstage(stage)
}

func (breathmark *BreathMark) StageVoid(stage *Stage) {
	breathmark.Stage(stage)
}

// for satisfaction of GongStruct interface
func (breathmark *BreathMark) GetName() (res string) {
	return breathmark.Name
}

// for satisfaction of GongStruct interface
func (breathmark *BreathMark) SetName(name string) {
	breathmark.Name = name
}

// Stage puts caesura to the model stage
func (caesura *Caesura) Stage(stage *Stage) *Caesura {
	__gong__stage(stage.Caesuras, stage.Caesura_stagedOrder, stage.Caesura_orderStaged, &stage.CaesuraOrder, stage.Caesuras_mapString, caesura, caesura.Name)
	return caesura
}

// StagePreserveOrder puts caesura to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CaesuraOrder
// - update stage.CaesuraOrder accordingly
func (caesura *Caesura) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Caesuras, stage.Caesura_stagedOrder, stage.Caesura_orderStaged, &stage.CaesuraOrder, stage.Caesuras_mapString, caesura, order, caesura.Name)
}

// Unstage removes caesura off the model stage
func (caesura *Caesura) Unstage(stage *Stage) *Caesura {
	__gong__unstage(stage.Caesuras, stage.Caesuras_mapString, caesura, caesura.Name)
	return caesura
}

// UnstageVoid removes caesura off the model stage
func (caesura *Caesura) UnstageVoid(stage *Stage) {
	caesura.Unstage(stage)
}

func (caesura *Caesura) StageVoid(stage *Stage) {
	caesura.Stage(stage)
}

// for satisfaction of GongStruct interface
func (caesura *Caesura) GetName() (res string) {
	return caesura.Name
}

// for satisfaction of GongStruct interface
func (caesura *Caesura) SetName(name string) {
	caesura.Name = name
}

// Stage puts elision to the model stage
func (elision *Elision) Stage(stage *Stage) *Elision {
	__gong__stage(stage.Elisions, stage.Elision_stagedOrder, stage.Elision_orderStaged, &stage.ElisionOrder, stage.Elisions_mapString, elision, elision.Name)
	return elision
}

// StagePreserveOrder puts elision to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ElisionOrder
// - update stage.ElisionOrder accordingly
func (elision *Elision) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Elisions, stage.Elision_stagedOrder, stage.Elision_orderStaged, &stage.ElisionOrder, stage.Elisions_mapString, elision, order, elision.Name)
}

// Unstage removes elision off the model stage
func (elision *Elision) Unstage(stage *Stage) *Elision {
	__gong__unstage(stage.Elisions, stage.Elisions_mapString, elision, elision.Name)
	return elision
}

// UnstageVoid removes elision off the model stage
func (elision *Elision) UnstageVoid(stage *Stage) {
	elision.Unstage(stage)
}

func (elision *Elision) StageVoid(stage *Stage) {
	elision.Stage(stage)
}

// for satisfaction of GongStruct interface
func (elision *Elision) GetName() (res string) {
	return elision.Name
}

// for satisfaction of GongStruct interface
func (elision *Elision) SetName(name string) {
	elision.Name = name
}

// Stage puts emptyline to the model stage
func (emptyline *EmptyLine) Stage(stage *Stage) *EmptyLine {
	__gong__stage(stage.EmptyLines, stage.EmptyLine_stagedOrder, stage.EmptyLine_orderStaged, &stage.EmptyLineOrder, stage.EmptyLines_mapString, emptyline, emptyline.Name)
	return emptyline
}

// StagePreserveOrder puts emptyline to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyLineOrder
// - update stage.EmptyLineOrder accordingly
func (emptyline *EmptyLine) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EmptyLines, stage.EmptyLine_stagedOrder, stage.EmptyLine_orderStaged, &stage.EmptyLineOrder, stage.EmptyLines_mapString, emptyline, order, emptyline.Name)
}

// Unstage removes emptyline off the model stage
func (emptyline *EmptyLine) Unstage(stage *Stage) *EmptyLine {
	__gong__unstage(stage.EmptyLines, stage.EmptyLines_mapString, emptyline, emptyline.Name)
	return emptyline
}

// UnstageVoid removes emptyline off the model stage
func (emptyline *EmptyLine) UnstageVoid(stage *Stage) {
	emptyline.Unstage(stage)
}

func (emptyline *EmptyLine) StageVoid(stage *Stage) {
	emptyline.Stage(stage)
}

// for satisfaction of GongStruct interface
func (emptyline *EmptyLine) GetName() (res string) {
	return emptyline.Name
}

// for satisfaction of GongStruct interface
func (emptyline *EmptyLine) SetName(name string) {
	emptyline.Name = name
}

// Stage puts extend to the model stage
func (extend *Extend) Stage(stage *Stage) *Extend {
	__gong__stage(stage.Extends, stage.Extend_stagedOrder, stage.Extend_orderStaged, &stage.ExtendOrder, stage.Extends_mapString, extend, extend.Name)
	return extend
}

// StagePreserveOrder puts extend to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExtendOrder
// - update stage.ExtendOrder accordingly
func (extend *Extend) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Extends, stage.Extend_stagedOrder, stage.Extend_orderStaged, &stage.ExtendOrder, stage.Extends_mapString, extend, order, extend.Name)
}

// Unstage removes extend off the model stage
func (extend *Extend) Unstage(stage *Stage) *Extend {
	__gong__unstage(stage.Extends, stage.Extends_mapString, extend, extend.Name)
	return extend
}

// UnstageVoid removes extend off the model stage
func (extend *Extend) UnstageVoid(stage *Stage) {
	extend.Unstage(stage)
}

func (extend *Extend) StageVoid(stage *Stage) {
	extend.Stage(stage)
}

// for satisfaction of GongStruct interface
func (extend *Extend) GetName() (res string) {
	return extend.Name
}

// for satisfaction of GongStruct interface
func (extend *Extend) SetName(name string) {
	extend.Name = name
}

// Stage puts figure to the model stage
func (figure *Figure) Stage(stage *Stage) *Figure {
	__gong__stage(stage.Figures, stage.Figure_stagedOrder, stage.Figure_orderStaged, &stage.FigureOrder, stage.Figures_mapString, figure, figure.Name)
	return figure
}

// StagePreserveOrder puts figure to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FigureOrder
// - update stage.FigureOrder accordingly
func (figure *Figure) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Figures, stage.Figure_stagedOrder, stage.Figure_orderStaged, &stage.FigureOrder, stage.Figures_mapString, figure, order, figure.Name)
}

// Unstage removes figure off the model stage
func (figure *Figure) Unstage(stage *Stage) *Figure {
	__gong__unstage(stage.Figures, stage.Figures_mapString, figure, figure.Name)
	return figure
}

// UnstageVoid removes figure off the model stage
func (figure *Figure) UnstageVoid(stage *Stage) {
	figure.Unstage(stage)
}

func (figure *Figure) StageVoid(stage *Stage) {
	figure.Stage(stage)
}

// for satisfaction of GongStruct interface
func (figure *Figure) GetName() (res string) {
	return figure.Name
}

// for satisfaction of GongStruct interface
func (figure *Figure) SetName(name string) {
	figure.Name = name
}

// Stage puts figuredbass to the model stage
func (figuredbass *FiguredBass) Stage(stage *Stage) *FiguredBass {
	__gong__stage(stage.FiguredBasss, stage.FiguredBass_stagedOrder, stage.FiguredBass_orderStaged, &stage.FiguredBassOrder, stage.FiguredBasss_mapString, figuredbass, figuredbass.Name)
	return figuredbass
}

// StagePreserveOrder puts figuredbass to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FiguredBassOrder
// - update stage.FiguredBassOrder accordingly
func (figuredbass *FiguredBass) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FiguredBasss, stage.FiguredBass_stagedOrder, stage.FiguredBass_orderStaged, &stage.FiguredBassOrder, stage.FiguredBasss_mapString, figuredbass, order, figuredbass.Name)
}

// Unstage removes figuredbass off the model stage
func (figuredbass *FiguredBass) Unstage(stage *Stage) *FiguredBass {
	__gong__unstage(stage.FiguredBasss, stage.FiguredBasss_mapString, figuredbass, figuredbass.Name)
	return figuredbass
}

// UnstageVoid removes figuredbass off the model stage
func (figuredbass *FiguredBass) UnstageVoid(stage *Stage) {
	figuredbass.Unstage(stage)
}

func (figuredbass *FiguredBass) StageVoid(stage *Stage) {
	figuredbass.Stage(stage)
}

// for satisfaction of GongStruct interface
func (figuredbass *FiguredBass) GetName() (res string) {
	return figuredbass.Name
}

// for satisfaction of GongStruct interface
func (figuredbass *FiguredBass) SetName(name string) {
	figuredbass.Name = name
}

// Stage puts forward to the model stage
func (forward *Forward) Stage(stage *Stage) *Forward {
	__gong__stage(stage.Forwards, stage.Forward_stagedOrder, stage.Forward_orderStaged, &stage.ForwardOrder, stage.Forwards_mapString, forward, forward.Name)
	return forward
}

// StagePreserveOrder puts forward to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ForwardOrder
// - update stage.ForwardOrder accordingly
func (forward *Forward) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Forwards, stage.Forward_stagedOrder, stage.Forward_orderStaged, &stage.ForwardOrder, stage.Forwards_mapString, forward, order, forward.Name)
}

// Unstage removes forward off the model stage
func (forward *Forward) Unstage(stage *Stage) *Forward {
	__gong__unstage(stage.Forwards, stage.Forwards_mapString, forward, forward.Name)
	return forward
}

// UnstageVoid removes forward off the model stage
func (forward *Forward) UnstageVoid(stage *Stage) {
	forward.Unstage(stage)
}

func (forward *Forward) StageVoid(stage *Stage) {
	forward.Stage(stage)
}

// for satisfaction of GongStruct interface
func (forward *Forward) GetName() (res string) {
	return forward.Name
}

// for satisfaction of GongStruct interface
func (forward *Forward) SetName(name string) {
	forward.Name = name
}

// Stage puts glissando to the model stage
func (glissando *Glissando) Stage(stage *Stage) *Glissando {
	__gong__stage(stage.Glissandos, stage.Glissando_stagedOrder, stage.Glissando_orderStaged, &stage.GlissandoOrder, stage.Glissandos_mapString, glissando, glissando.Name)
	return glissando
}

// StagePreserveOrder puts glissando to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GlissandoOrder
// - update stage.GlissandoOrder accordingly
func (glissando *Glissando) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Glissandos, stage.Glissando_stagedOrder, stage.Glissando_orderStaged, &stage.GlissandoOrder, stage.Glissandos_mapString, glissando, order, glissando.Name)
}

// Unstage removes glissando off the model stage
func (glissando *Glissando) Unstage(stage *Stage) *Glissando {
	__gong__unstage(stage.Glissandos, stage.Glissandos_mapString, glissando, glissando.Name)
	return glissando
}

// UnstageVoid removes glissando off the model stage
func (glissando *Glissando) UnstageVoid(stage *Stage) {
	glissando.Unstage(stage)
}

func (glissando *Glissando) StageVoid(stage *Stage) {
	glissando.Stage(stage)
}

// for satisfaction of GongStruct interface
func (glissando *Glissando) GetName() (res string) {
	return glissando.Name
}

// for satisfaction of GongStruct interface
func (glissando *Glissando) SetName(name string) {
	glissando.Name = name
}

// Stage puts grace to the model stage
func (grace *Grace) Stage(stage *Stage) *Grace {
	__gong__stage(stage.Graces, stage.Grace_stagedOrder, stage.Grace_orderStaged, &stage.GraceOrder, stage.Graces_mapString, grace, grace.Name)
	return grace
}

// StagePreserveOrder puts grace to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GraceOrder
// - update stage.GraceOrder accordingly
func (grace *Grace) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Graces, stage.Grace_stagedOrder, stage.Grace_orderStaged, &stage.GraceOrder, stage.Graces_mapString, grace, order, grace.Name)
}

// Unstage removes grace off the model stage
func (grace *Grace) Unstage(stage *Stage) *Grace {
	__gong__unstage(stage.Graces, stage.Graces_mapString, grace, grace.Name)
	return grace
}

// UnstageVoid removes grace off the model stage
func (grace *Grace) UnstageVoid(stage *Stage) {
	grace.Unstage(stage)
}

func (grace *Grace) StageVoid(stage *Stage) {
	grace.Stage(stage)
}

// for satisfaction of GongStruct interface
func (grace *Grace) GetName() (res string) {
	return grace.Name
}

// for satisfaction of GongStruct interface
func (grace *Grace) SetName(name string) {
	grace.Name = name
}

// Stage puts hammeronpulloff to the model stage
func (hammeronpulloff *HammerOnPullOff) Stage(stage *Stage) *HammerOnPullOff {
	__gong__stage(stage.HammerOnPullOffs, stage.HammerOnPullOff_stagedOrder, stage.HammerOnPullOff_orderStaged, &stage.HammerOnPullOffOrder, stage.HammerOnPullOffs_mapString, hammeronpulloff, hammeronpulloff.Name)
	return hammeronpulloff
}

// StagePreserveOrder puts hammeronpulloff to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HammerOnPullOffOrder
// - update stage.HammerOnPullOffOrder accordingly
func (hammeronpulloff *HammerOnPullOff) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.HammerOnPullOffs, stage.HammerOnPullOff_stagedOrder, stage.HammerOnPullOff_orderStaged, &stage.HammerOnPullOffOrder, stage.HammerOnPullOffs_mapString, hammeronpulloff, order, hammeronpulloff.Name)
}

// Unstage removes hammeronpulloff off the model stage
func (hammeronpulloff *HammerOnPullOff) Unstage(stage *Stage) *HammerOnPullOff {
	__gong__unstage(stage.HammerOnPullOffs, stage.HammerOnPullOffs_mapString, hammeronpulloff, hammeronpulloff.Name)
	return hammeronpulloff
}

// UnstageVoid removes hammeronpulloff off the model stage
func (hammeronpulloff *HammerOnPullOff) UnstageVoid(stage *Stage) {
	hammeronpulloff.Unstage(stage)
}

func (hammeronpulloff *HammerOnPullOff) StageVoid(stage *Stage) {
	hammeronpulloff.Stage(stage)
}

// for satisfaction of GongStruct interface
func (hammeronpulloff *HammerOnPullOff) GetName() (res string) {
	return hammeronpulloff.Name
}

// for satisfaction of GongStruct interface
func (hammeronpulloff *HammerOnPullOff) SetName(name string) {
	hammeronpulloff.Name = name
}

// Stage puts handbell to the model stage
func (handbell *Handbell) Stage(stage *Stage) *Handbell {
	__gong__stage(stage.Handbells, stage.Handbell_stagedOrder, stage.Handbell_orderStaged, &stage.HandbellOrder, stage.Handbells_mapString, handbell, handbell.Name)
	return handbell
}

// StagePreserveOrder puts handbell to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HandbellOrder
// - update stage.HandbellOrder accordingly
func (handbell *Handbell) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Handbells, stage.Handbell_stagedOrder, stage.Handbell_orderStaged, &stage.HandbellOrder, stage.Handbells_mapString, handbell, order, handbell.Name)
}

// Unstage removes handbell off the model stage
func (handbell *Handbell) Unstage(stage *Stage) *Handbell {
	__gong__unstage(stage.Handbells, stage.Handbells_mapString, handbell, handbell.Name)
	return handbell
}

// UnstageVoid removes handbell off the model stage
func (handbell *Handbell) UnstageVoid(stage *Stage) {
	handbell.Unstage(stage)
}

func (handbell *Handbell) StageVoid(stage *Stage) {
	handbell.Stage(stage)
}

// for satisfaction of GongStruct interface
func (handbell *Handbell) GetName() (res string) {
	return handbell.Name
}

// for satisfaction of GongStruct interface
func (handbell *Handbell) SetName(name string) {
	handbell.Name = name
}

// Stage puts harmonclosed to the model stage
func (harmonclosed *HarmonClosed) Stage(stage *Stage) *HarmonClosed {
	__gong__stage(stage.HarmonCloseds, stage.HarmonClosed_stagedOrder, stage.HarmonClosed_orderStaged, &stage.HarmonClosedOrder, stage.HarmonCloseds_mapString, harmonclosed, harmonclosed.Name)
	return harmonclosed
}

// StagePreserveOrder puts harmonclosed to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HarmonClosedOrder
// - update stage.HarmonClosedOrder accordingly
func (harmonclosed *HarmonClosed) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.HarmonCloseds, stage.HarmonClosed_stagedOrder, stage.HarmonClosed_orderStaged, &stage.HarmonClosedOrder, stage.HarmonCloseds_mapString, harmonclosed, order, harmonclosed.Name)
}

// Unstage removes harmonclosed off the model stage
func (harmonclosed *HarmonClosed) Unstage(stage *Stage) *HarmonClosed {
	__gong__unstage(stage.HarmonCloseds, stage.HarmonCloseds_mapString, harmonclosed, harmonclosed.Name)
	return harmonclosed
}

// UnstageVoid removes harmonclosed off the model stage
func (harmonclosed *HarmonClosed) UnstageVoid(stage *Stage) {
	harmonclosed.Unstage(stage)
}

func (harmonclosed *HarmonClosed) StageVoid(stage *Stage) {
	harmonclosed.Stage(stage)
}

// for satisfaction of GongStruct interface
func (harmonclosed *HarmonClosed) GetName() (res string) {
	return harmonclosed.Name
}

// for satisfaction of GongStruct interface
func (harmonclosed *HarmonClosed) SetName(name string) {
	harmonclosed.Name = name
}

// Stage puts harmonmute to the model stage
func (harmonmute *HarmonMute) Stage(stage *Stage) *HarmonMute {
	__gong__stage(stage.HarmonMutes, stage.HarmonMute_stagedOrder, stage.HarmonMute_orderStaged, &stage.HarmonMuteOrder, stage.HarmonMutes_mapString, harmonmute, harmonmute.Name)
	return harmonmute
}

// StagePreserveOrder puts harmonmute to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HarmonMuteOrder
// - update stage.HarmonMuteOrder accordingly
func (harmonmute *HarmonMute) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.HarmonMutes, stage.HarmonMute_stagedOrder, stage.HarmonMute_orderStaged, &stage.HarmonMuteOrder, stage.HarmonMutes_mapString, harmonmute, order, harmonmute.Name)
}

// Unstage removes harmonmute off the model stage
func (harmonmute *HarmonMute) Unstage(stage *Stage) *HarmonMute {
	__gong__unstage(stage.HarmonMutes, stage.HarmonMutes_mapString, harmonmute, harmonmute.Name)
	return harmonmute
}

// UnstageVoid removes harmonmute off the model stage
func (harmonmute *HarmonMute) UnstageVoid(stage *Stage) {
	harmonmute.Unstage(stage)
}

func (harmonmute *HarmonMute) StageVoid(stage *Stage) {
	harmonmute.Stage(stage)
}

// for satisfaction of GongStruct interface
func (harmonmute *HarmonMute) GetName() (res string) {
	return harmonmute.Name
}

// for satisfaction of GongStruct interface
func (harmonmute *HarmonMute) SetName(name string) {
	harmonmute.Name = name
}

// Stage puts harmonic to the model stage
func (harmonic *Harmonic) Stage(stage *Stage) *Harmonic {
	__gong__stage(stage.Harmonics, stage.Harmonic_stagedOrder, stage.Harmonic_orderStaged, &stage.HarmonicOrder, stage.Harmonics_mapString, harmonic, harmonic.Name)
	return harmonic
}

// StagePreserveOrder puts harmonic to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HarmonicOrder
// - update stage.HarmonicOrder accordingly
func (harmonic *Harmonic) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Harmonics, stage.Harmonic_stagedOrder, stage.Harmonic_orderStaged, &stage.HarmonicOrder, stage.Harmonics_mapString, harmonic, order, harmonic.Name)
}

// Unstage removes harmonic off the model stage
func (harmonic *Harmonic) Unstage(stage *Stage) *Harmonic {
	__gong__unstage(stage.Harmonics, stage.Harmonics_mapString, harmonic, harmonic.Name)
	return harmonic
}

// UnstageVoid removes harmonic off the model stage
func (harmonic *Harmonic) UnstageVoid(stage *Stage) {
	harmonic.Unstage(stage)
}

func (harmonic *Harmonic) StageVoid(stage *Stage) {
	harmonic.Stage(stage)
}

// for satisfaction of GongStruct interface
func (harmonic *Harmonic) GetName() (res string) {
	return harmonic.Name
}

// for satisfaction of GongStruct interface
func (harmonic *Harmonic) SetName(name string) {
	harmonic.Name = name
}

// Stage puts heeltoe to the model stage
func (heeltoe *HeelToe) Stage(stage *Stage) *HeelToe {
	__gong__stage(stage.HeelToes, stage.HeelToe_stagedOrder, stage.HeelToe_orderStaged, &stage.HeelToeOrder, stage.HeelToes_mapString, heeltoe, heeltoe.Name)
	return heeltoe
}

// StagePreserveOrder puts heeltoe to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HeelToeOrder
// - update stage.HeelToeOrder accordingly
func (heeltoe *HeelToe) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.HeelToes, stage.HeelToe_stagedOrder, stage.HeelToe_orderStaged, &stage.HeelToeOrder, stage.HeelToes_mapString, heeltoe, order, heeltoe.Name)
}

// Unstage removes heeltoe off the model stage
func (heeltoe *HeelToe) Unstage(stage *Stage) *HeelToe {
	__gong__unstage(stage.HeelToes, stage.HeelToes_mapString, heeltoe, heeltoe.Name)
	return heeltoe
}

// UnstageVoid removes heeltoe off the model stage
func (heeltoe *HeelToe) UnstageVoid(stage *Stage) {
	heeltoe.Unstage(stage)
}

func (heeltoe *HeelToe) StageVoid(stage *Stage) {
	heeltoe.Stage(stage)
}

// for satisfaction of GongStruct interface
func (heeltoe *HeelToe) GetName() (res string) {
	return heeltoe.Name
}

// for satisfaction of GongStruct interface
func (heeltoe *HeelToe) SetName(name string) {
	heeltoe.Name = name
}

// Stage puts hole to the model stage
func (hole *Hole) Stage(stage *Stage) *Hole {
	__gong__stage(stage.Holes, stage.Hole_stagedOrder, stage.Hole_orderStaged, &stage.HoleOrder, stage.Holes_mapString, hole, hole.Name)
	return hole
}

// StagePreserveOrder puts hole to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HoleOrder
// - update stage.HoleOrder accordingly
func (hole *Hole) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Holes, stage.Hole_stagedOrder, stage.Hole_orderStaged, &stage.HoleOrder, stage.Holes_mapString, hole, order, hole.Name)
}

// Unstage removes hole off the model stage
func (hole *Hole) Unstage(stage *Stage) *Hole {
	__gong__unstage(stage.Holes, stage.Holes_mapString, hole, hole.Name)
	return hole
}

// UnstageVoid removes hole off the model stage
func (hole *Hole) UnstageVoid(stage *Stage) {
	hole.Unstage(stage)
}

func (hole *Hole) StageVoid(stage *Stage) {
	hole.Stage(stage)
}

// for satisfaction of GongStruct interface
func (hole *Hole) GetName() (res string) {
	return hole.Name
}

// for satisfaction of GongStruct interface
func (hole *Hole) SetName(name string) {
	hole.Name = name
}

// Stage puts holeclosed to the model stage
func (holeclosed *HoleClosed) Stage(stage *Stage) *HoleClosed {
	__gong__stage(stage.HoleCloseds, stage.HoleClosed_stagedOrder, stage.HoleClosed_orderStaged, &stage.HoleClosedOrder, stage.HoleCloseds_mapString, holeclosed, holeclosed.Name)
	return holeclosed
}

// StagePreserveOrder puts holeclosed to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HoleClosedOrder
// - update stage.HoleClosedOrder accordingly
func (holeclosed *HoleClosed) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.HoleCloseds, stage.HoleClosed_stagedOrder, stage.HoleClosed_orderStaged, &stage.HoleClosedOrder, stage.HoleCloseds_mapString, holeclosed, order, holeclosed.Name)
}

// Unstage removes holeclosed off the model stage
func (holeclosed *HoleClosed) Unstage(stage *Stage) *HoleClosed {
	__gong__unstage(stage.HoleCloseds, stage.HoleCloseds_mapString, holeclosed, holeclosed.Name)
	return holeclosed
}

// UnstageVoid removes holeclosed off the model stage
func (holeclosed *HoleClosed) UnstageVoid(stage *Stage) {
	holeclosed.Unstage(stage)
}

func (holeclosed *HoleClosed) StageVoid(stage *Stage) {
	holeclosed.Stage(stage)
}

// for satisfaction of GongStruct interface
func (holeclosed *HoleClosed) GetName() (res string) {
	return holeclosed.Name
}

// for satisfaction of GongStruct interface
func (holeclosed *HoleClosed) SetName(name string) {
	holeclosed.Name = name
}

// Stage puts instrument to the model stage
func (instrument *Instrument) Stage(stage *Stage) *Instrument {
	__gong__stage(stage.Instruments, stage.Instrument_stagedOrder, stage.Instrument_orderStaged, &stage.InstrumentOrder, stage.Instruments_mapString, instrument, instrument.Name)
	return instrument
}

// StagePreserveOrder puts instrument to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.InstrumentOrder
// - update stage.InstrumentOrder accordingly
func (instrument *Instrument) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Instruments, stage.Instrument_stagedOrder, stage.Instrument_orderStaged, &stage.InstrumentOrder, stage.Instruments_mapString, instrument, order, instrument.Name)
}

// Unstage removes instrument off the model stage
func (instrument *Instrument) Unstage(stage *Stage) *Instrument {
	__gong__unstage(stage.Instruments, stage.Instruments_mapString, instrument, instrument.Name)
	return instrument
}

// UnstageVoid removes instrument off the model stage
func (instrument *Instrument) UnstageVoid(stage *Stage) {
	instrument.Unstage(stage)
}

func (instrument *Instrument) StageVoid(stage *Stage) {
	instrument.Stage(stage)
}

// for satisfaction of GongStruct interface
func (instrument *Instrument) GetName() (res string) {
	return instrument.Name
}

// for satisfaction of GongStruct interface
func (instrument *Instrument) SetName(name string) {
	instrument.Name = name
}

// Stage puts listen to the model stage
func (listen *Listen) Stage(stage *Stage) *Listen {
	__gong__stage(stage.Listens, stage.Listen_stagedOrder, stage.Listen_orderStaged, &stage.ListenOrder, stage.Listens_mapString, listen, listen.Name)
	return listen
}

// StagePreserveOrder puts listen to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ListenOrder
// - update stage.ListenOrder accordingly
func (listen *Listen) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Listens, stage.Listen_stagedOrder, stage.Listen_orderStaged, &stage.ListenOrder, stage.Listens_mapString, listen, order, listen.Name)
}

// Unstage removes listen off the model stage
func (listen *Listen) Unstage(stage *Stage) *Listen {
	__gong__unstage(stage.Listens, stage.Listens_mapString, listen, listen.Name)
	return listen
}

// UnstageVoid removes listen off the model stage
func (listen *Listen) UnstageVoid(stage *Stage) {
	listen.Unstage(stage)
}

func (listen *Listen) StageVoid(stage *Stage) {
	listen.Stage(stage)
}

// for satisfaction of GongStruct interface
func (listen *Listen) GetName() (res string) {
	return listen.Name
}

// for satisfaction of GongStruct interface
func (listen *Listen) SetName(name string) {
	listen.Name = name
}

// Stage puts lyric to the model stage
func (lyric *Lyric) Stage(stage *Stage) *Lyric {
	__gong__stage(stage.Lyrics, stage.Lyric_stagedOrder, stage.Lyric_orderStaged, &stage.LyricOrder, stage.Lyrics_mapString, lyric, lyric.Name)
	return lyric
}

// StagePreserveOrder puts lyric to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LyricOrder
// - update stage.LyricOrder accordingly
func (lyric *Lyric) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Lyrics, stage.Lyric_stagedOrder, stage.Lyric_orderStaged, &stage.LyricOrder, stage.Lyrics_mapString, lyric, order, lyric.Name)
}

// Unstage removes lyric off the model stage
func (lyric *Lyric) Unstage(stage *Stage) *Lyric {
	__gong__unstage(stage.Lyrics, stage.Lyrics_mapString, lyric, lyric.Name)
	return lyric
}

// UnstageVoid removes lyric off the model stage
func (lyric *Lyric) UnstageVoid(stage *Stage) {
	lyric.Unstage(stage)
}

func (lyric *Lyric) StageVoid(stage *Stage) {
	lyric.Stage(stage)
}

// for satisfaction of GongStruct interface
func (lyric *Lyric) GetName() (res string) {
	return lyric.Name
}

// for satisfaction of GongStruct interface
func (lyric *Lyric) SetName(name string) {
	lyric.Name = name
}

// Stage puts mordent to the model stage
func (mordent *Mordent) Stage(stage *Stage) *Mordent {
	__gong__stage(stage.Mordents, stage.Mordent_stagedOrder, stage.Mordent_orderStaged, &stage.MordentOrder, stage.Mordents_mapString, mordent, mordent.Name)
	return mordent
}

// StagePreserveOrder puts mordent to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MordentOrder
// - update stage.MordentOrder accordingly
func (mordent *Mordent) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Mordents, stage.Mordent_stagedOrder, stage.Mordent_orderStaged, &stage.MordentOrder, stage.Mordents_mapString, mordent, order, mordent.Name)
}

// Unstage removes mordent off the model stage
func (mordent *Mordent) Unstage(stage *Stage) *Mordent {
	__gong__unstage(stage.Mordents, stage.Mordents_mapString, mordent, mordent.Name)
	return mordent
}

// UnstageVoid removes mordent off the model stage
func (mordent *Mordent) UnstageVoid(stage *Stage) {
	mordent.Unstage(stage)
}

func (mordent *Mordent) StageVoid(stage *Stage) {
	mordent.Stage(stage)
}

// for satisfaction of GongStruct interface
func (mordent *Mordent) GetName() (res string) {
	return mordent.Name
}

// for satisfaction of GongStruct interface
func (mordent *Mordent) SetName(name string) {
	mordent.Name = name
}

// Stage puts nonarpeggiate to the model stage
func (nonarpeggiate *NonArpeggiate) Stage(stage *Stage) *NonArpeggiate {
	__gong__stage(stage.NonArpeggiates, stage.NonArpeggiate_stagedOrder, stage.NonArpeggiate_orderStaged, &stage.NonArpeggiateOrder, stage.NonArpeggiates_mapString, nonarpeggiate, nonarpeggiate.Name)
	return nonarpeggiate
}

// StagePreserveOrder puts nonarpeggiate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NonArpeggiateOrder
// - update stage.NonArpeggiateOrder accordingly
func (nonarpeggiate *NonArpeggiate) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NonArpeggiates, stage.NonArpeggiate_stagedOrder, stage.NonArpeggiate_orderStaged, &stage.NonArpeggiateOrder, stage.NonArpeggiates_mapString, nonarpeggiate, order, nonarpeggiate.Name)
}

// Unstage removes nonarpeggiate off the model stage
func (nonarpeggiate *NonArpeggiate) Unstage(stage *Stage) *NonArpeggiate {
	__gong__unstage(stage.NonArpeggiates, stage.NonArpeggiates_mapString, nonarpeggiate, nonarpeggiate.Name)
	return nonarpeggiate
}

// UnstageVoid removes nonarpeggiate off the model stage
func (nonarpeggiate *NonArpeggiate) UnstageVoid(stage *Stage) {
	nonarpeggiate.Unstage(stage)
}

func (nonarpeggiate *NonArpeggiate) StageVoid(stage *Stage) {
	nonarpeggiate.Stage(stage)
}

// for satisfaction of GongStruct interface
func (nonarpeggiate *NonArpeggiate) GetName() (res string) {
	return nonarpeggiate.Name
}

// for satisfaction of GongStruct interface
func (nonarpeggiate *NonArpeggiate) SetName(name string) {
	nonarpeggiate.Name = name
}

// Stage puts notations to the model stage
func (notations *Notations) Stage(stage *Stage) *Notations {
	__gong__stage(stage.Notationss, stage.Notations_stagedOrder, stage.Notations_orderStaged, &stage.NotationsOrder, stage.Notationss_mapString, notations, notations.Name)
	return notations
}

// StagePreserveOrder puts notations to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NotationsOrder
// - update stage.NotationsOrder accordingly
func (notations *Notations) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Notationss, stage.Notations_stagedOrder, stage.Notations_orderStaged, &stage.NotationsOrder, stage.Notationss_mapString, notations, order, notations.Name)
}

// Unstage removes notations off the model stage
func (notations *Notations) Unstage(stage *Stage) *Notations {
	__gong__unstage(stage.Notationss, stage.Notationss_mapString, notations, notations.Name)
	return notations
}

// UnstageVoid removes notations off the model stage
func (notations *Notations) UnstageVoid(stage *Stage) {
	notations.Unstage(stage)
}

func (notations *Notations) StageVoid(stage *Stage) {
	notations.Stage(stage)
}

// for satisfaction of GongStruct interface
func (notations *Notations) GetName() (res string) {
	return notations.Name
}

// for satisfaction of GongStruct interface
func (notations *Notations) SetName(name string) {
	notations.Name = name
}

// Stage puts note to the model stage
func (note *Note) Stage(stage *Stage) *Note {
	__gong__stage(stage.Notes, stage.Note_stagedOrder, stage.Note_orderStaged, &stage.NoteOrder, stage.Notes_mapString, note, note.Name)
	return note
}

// StagePreserveOrder puts note to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteOrder
// - update stage.NoteOrder accordingly
func (note *Note) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Notes, stage.Note_stagedOrder, stage.Note_orderStaged, &stage.NoteOrder, stage.Notes_mapString, note, order, note.Name)
}

// Unstage removes note off the model stage
func (note *Note) Unstage(stage *Stage) *Note {
	__gong__unstage(stage.Notes, stage.Notes_mapString, note, note.Name)
	return note
}

// UnstageVoid removes note off the model stage
func (note *Note) UnstageVoid(stage *Stage) {
	note.Unstage(stage)
}

func (note *Note) StageVoid(stage *Stage) {
	note.Stage(stage)
}

// for satisfaction of GongStruct interface
func (note *Note) GetName() (res string) {
	return note.Name
}

// for satisfaction of GongStruct interface
func (note *Note) SetName(name string) {
	note.Name = name
}

// Stage puts notetype to the model stage
func (notetype *NoteType) Stage(stage *Stage) *NoteType {
	__gong__stage(stage.NoteTypes, stage.NoteType_stagedOrder, stage.NoteType_orderStaged, &stage.NoteTypeOrder, stage.NoteTypes_mapString, notetype, notetype.Name)
	return notetype
}

// StagePreserveOrder puts notetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteTypeOrder
// - update stage.NoteTypeOrder accordingly
func (notetype *NoteType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteTypes, stage.NoteType_stagedOrder, stage.NoteType_orderStaged, &stage.NoteTypeOrder, stage.NoteTypes_mapString, notetype, order, notetype.Name)
}

// Unstage removes notetype off the model stage
func (notetype *NoteType) Unstage(stage *Stage) *NoteType {
	__gong__unstage(stage.NoteTypes, stage.NoteTypes_mapString, notetype, notetype.Name)
	return notetype
}

// UnstageVoid removes notetype off the model stage
func (notetype *NoteType) UnstageVoid(stage *Stage) {
	notetype.Unstage(stage)
}

func (notetype *NoteType) StageVoid(stage *Stage) {
	notetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (notetype *NoteType) GetName() (res string) {
	return notetype.Name
}

// for satisfaction of GongStruct interface
func (notetype *NoteType) SetName(name string) {
	notetype.Name = name
}

// Stage puts notehead to the model stage
func (notehead *Notehead) Stage(stage *Stage) *Notehead {
	__gong__stage(stage.Noteheads, stage.Notehead_stagedOrder, stage.Notehead_orderStaged, &stage.NoteheadOrder, stage.Noteheads_mapString, notehead, notehead.Name)
	return notehead
}

// StagePreserveOrder puts notehead to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteheadOrder
// - update stage.NoteheadOrder accordingly
func (notehead *Notehead) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Noteheads, stage.Notehead_stagedOrder, stage.Notehead_orderStaged, &stage.NoteheadOrder, stage.Noteheads_mapString, notehead, order, notehead.Name)
}

// Unstage removes notehead off the model stage
func (notehead *Notehead) Unstage(stage *Stage) *Notehead {
	__gong__unstage(stage.Noteheads, stage.Noteheads_mapString, notehead, notehead.Name)
	return notehead
}

// UnstageVoid removes notehead off the model stage
func (notehead *Notehead) UnstageVoid(stage *Stage) {
	notehead.Unstage(stage)
}

func (notehead *Notehead) StageVoid(stage *Stage) {
	notehead.Stage(stage)
}

// for satisfaction of GongStruct interface
func (notehead *Notehead) GetName() (res string) {
	return notehead.Name
}

// for satisfaction of GongStruct interface
func (notehead *Notehead) SetName(name string) {
	notehead.Name = name
}

// Stage puts noteheadtext to the model stage
func (noteheadtext *NoteheadText) Stage(stage *Stage) *NoteheadText {
	__gong__stage(stage.NoteheadTexts, stage.NoteheadText_stagedOrder, stage.NoteheadText_orderStaged, &stage.NoteheadTextOrder, stage.NoteheadTexts_mapString, noteheadtext, noteheadtext.Name)
	return noteheadtext
}

// StagePreserveOrder puts noteheadtext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteheadTextOrder
// - update stage.NoteheadTextOrder accordingly
func (noteheadtext *NoteheadText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteheadTexts, stage.NoteheadText_stagedOrder, stage.NoteheadText_orderStaged, &stage.NoteheadTextOrder, stage.NoteheadTexts_mapString, noteheadtext, order, noteheadtext.Name)
}

// Unstage removes noteheadtext off the model stage
func (noteheadtext *NoteheadText) Unstage(stage *Stage) *NoteheadText {
	__gong__unstage(stage.NoteheadTexts, stage.NoteheadTexts_mapString, noteheadtext, noteheadtext.Name)
	return noteheadtext
}

// UnstageVoid removes noteheadtext off the model stage
func (noteheadtext *NoteheadText) UnstageVoid(stage *Stage) {
	noteheadtext.Unstage(stage)
}

func (noteheadtext *NoteheadText) StageVoid(stage *Stage) {
	noteheadtext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (noteheadtext *NoteheadText) GetName() (res string) {
	return noteheadtext.Name
}

// for satisfaction of GongStruct interface
func (noteheadtext *NoteheadText) SetName(name string) {
	noteheadtext.Name = name
}

// Stage puts ornaments to the model stage
func (ornaments *Ornaments) Stage(stage *Stage) *Ornaments {
	__gong__stage(stage.Ornamentss, stage.Ornaments_stagedOrder, stage.Ornaments_orderStaged, &stage.OrnamentsOrder, stage.Ornamentss_mapString, ornaments, ornaments.Name)
	return ornaments
}

// StagePreserveOrder puts ornaments to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OrnamentsOrder
// - update stage.OrnamentsOrder accordingly
func (ornaments *Ornaments) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Ornamentss, stage.Ornaments_stagedOrder, stage.Ornaments_orderStaged, &stage.OrnamentsOrder, stage.Ornamentss_mapString, ornaments, order, ornaments.Name)
}

// Unstage removes ornaments off the model stage
func (ornaments *Ornaments) Unstage(stage *Stage) *Ornaments {
	__gong__unstage(stage.Ornamentss, stage.Ornamentss_mapString, ornaments, ornaments.Name)
	return ornaments
}

// UnstageVoid removes ornaments off the model stage
func (ornaments *Ornaments) UnstageVoid(stage *Stage) {
	ornaments.Unstage(stage)
}

func (ornaments *Ornaments) StageVoid(stage *Stage) {
	ornaments.Stage(stage)
}

// for satisfaction of GongStruct interface
func (ornaments *Ornaments) GetName() (res string) {
	return ornaments.Name
}

// for satisfaction of GongStruct interface
func (ornaments *Ornaments) SetName(name string) {
	ornaments.Name = name
}

// Stage puts othernotation to the model stage
func (othernotation *OtherNotation) Stage(stage *Stage) *OtherNotation {
	__gong__stage(stage.OtherNotations, stage.OtherNotation_stagedOrder, stage.OtherNotation_orderStaged, &stage.OtherNotationOrder, stage.OtherNotations_mapString, othernotation, othernotation.Name)
	return othernotation
}

// StagePreserveOrder puts othernotation to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OtherNotationOrder
// - update stage.OtherNotationOrder accordingly
func (othernotation *OtherNotation) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OtherNotations, stage.OtherNotation_stagedOrder, stage.OtherNotation_orderStaged, &stage.OtherNotationOrder, stage.OtherNotations_mapString, othernotation, order, othernotation.Name)
}

// Unstage removes othernotation off the model stage
func (othernotation *OtherNotation) Unstage(stage *Stage) *OtherNotation {
	__gong__unstage(stage.OtherNotations, stage.OtherNotations_mapString, othernotation, othernotation.Name)
	return othernotation
}

// UnstageVoid removes othernotation off the model stage
func (othernotation *OtherNotation) UnstageVoid(stage *Stage) {
	othernotation.Unstage(stage)
}

func (othernotation *OtherNotation) StageVoid(stage *Stage) {
	othernotation.Stage(stage)
}

// for satisfaction of GongStruct interface
func (othernotation *OtherNotation) GetName() (res string) {
	return othernotation.Name
}

// for satisfaction of GongStruct interface
func (othernotation *OtherNotation) SetName(name string) {
	othernotation.Name = name
}

// Stage puts otherplacementtext to the model stage
func (otherplacementtext *OtherPlacementText) Stage(stage *Stage) *OtherPlacementText {
	__gong__stage(stage.OtherPlacementTexts, stage.OtherPlacementText_stagedOrder, stage.OtherPlacementText_orderStaged, &stage.OtherPlacementTextOrder, stage.OtherPlacementTexts_mapString, otherplacementtext, otherplacementtext.Name)
	return otherplacementtext
}

// StagePreserveOrder puts otherplacementtext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OtherPlacementTextOrder
// - update stage.OtherPlacementTextOrder accordingly
func (otherplacementtext *OtherPlacementText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OtherPlacementTexts, stage.OtherPlacementText_stagedOrder, stage.OtherPlacementText_orderStaged, &stage.OtherPlacementTextOrder, stage.OtherPlacementTexts_mapString, otherplacementtext, order, otherplacementtext.Name)
}

// Unstage removes otherplacementtext off the model stage
func (otherplacementtext *OtherPlacementText) Unstage(stage *Stage) *OtherPlacementText {
	__gong__unstage(stage.OtherPlacementTexts, stage.OtherPlacementTexts_mapString, otherplacementtext, otherplacementtext.Name)
	return otherplacementtext
}

// UnstageVoid removes otherplacementtext off the model stage
func (otherplacementtext *OtherPlacementText) UnstageVoid(stage *Stage) {
	otherplacementtext.Unstage(stage)
}

func (otherplacementtext *OtherPlacementText) StageVoid(stage *Stage) {
	otherplacementtext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (otherplacementtext *OtherPlacementText) GetName() (res string) {
	return otherplacementtext.Name
}

// for satisfaction of GongStruct interface
func (otherplacementtext *OtherPlacementText) SetName(name string) {
	otherplacementtext.Name = name
}

// Stage puts pitch to the model stage
func (pitch *Pitch) Stage(stage *Stage) *Pitch {
	__gong__stage(stage.Pitchs, stage.Pitch_stagedOrder, stage.Pitch_orderStaged, &stage.PitchOrder, stage.Pitchs_mapString, pitch, pitch.Name)
	return pitch
}

// StagePreserveOrder puts pitch to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PitchOrder
// - update stage.PitchOrder accordingly
func (pitch *Pitch) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Pitchs, stage.Pitch_stagedOrder, stage.Pitch_orderStaged, &stage.PitchOrder, stage.Pitchs_mapString, pitch, order, pitch.Name)
}

// Unstage removes pitch off the model stage
func (pitch *Pitch) Unstage(stage *Stage) *Pitch {
	__gong__unstage(stage.Pitchs, stage.Pitchs_mapString, pitch, pitch.Name)
	return pitch
}

// UnstageVoid removes pitch off the model stage
func (pitch *Pitch) UnstageVoid(stage *Stage) {
	pitch.Unstage(stage)
}

func (pitch *Pitch) StageVoid(stage *Stage) {
	pitch.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pitch *Pitch) GetName() (res string) {
	return pitch.Name
}

// for satisfaction of GongStruct interface
func (pitch *Pitch) SetName(name string) {
	pitch.Name = name
}

// Stage puts placementtext to the model stage
func (placementtext *PlacementText) Stage(stage *Stage) *PlacementText {
	__gong__stage(stage.PlacementTexts, stage.PlacementText_stagedOrder, stage.PlacementText_orderStaged, &stage.PlacementTextOrder, stage.PlacementTexts_mapString, placementtext, placementtext.Name)
	return placementtext
}

// StagePreserveOrder puts placementtext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlacementTextOrder
// - update stage.PlacementTextOrder accordingly
func (placementtext *PlacementText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PlacementTexts, stage.PlacementText_stagedOrder, stage.PlacementText_orderStaged, &stage.PlacementTextOrder, stage.PlacementTexts_mapString, placementtext, order, placementtext.Name)
}

// Unstage removes placementtext off the model stage
func (placementtext *PlacementText) Unstage(stage *Stage) *PlacementText {
	__gong__unstage(stage.PlacementTexts, stage.PlacementTexts_mapString, placementtext, placementtext.Name)
	return placementtext
}

// UnstageVoid removes placementtext off the model stage
func (placementtext *PlacementText) UnstageVoid(stage *Stage) {
	placementtext.Unstage(stage)
}

func (placementtext *PlacementText) StageVoid(stage *Stage) {
	placementtext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (placementtext *PlacementText) GetName() (res string) {
	return placementtext.Name
}

// for satisfaction of GongStruct interface
func (placementtext *PlacementText) SetName(name string) {
	placementtext.Name = name
}

// Stage puts release to the model stage
func (release *Release) Stage(stage *Stage) *Release {
	__gong__stage(stage.Releases, stage.Release_stagedOrder, stage.Release_orderStaged, &stage.ReleaseOrder, stage.Releases_mapString, release, release.Name)
	return release
}

// StagePreserveOrder puts release to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ReleaseOrder
// - update stage.ReleaseOrder accordingly
func (release *Release) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Releases, stage.Release_stagedOrder, stage.Release_orderStaged, &stage.ReleaseOrder, stage.Releases_mapString, release, order, release.Name)
}

// Unstage removes release off the model stage
func (release *Release) Unstage(stage *Stage) *Release {
	__gong__unstage(stage.Releases, stage.Releases_mapString, release, release.Name)
	return release
}

// UnstageVoid removes release off the model stage
func (release *Release) UnstageVoid(stage *Stage) {
	release.Unstage(stage)
}

func (release *Release) StageVoid(stage *Stage) {
	release.Stage(stage)
}

// for satisfaction of GongStruct interface
func (release *Release) GetName() (res string) {
	return release.Name
}

// for satisfaction of GongStruct interface
func (release *Release) SetName(name string) {
	release.Name = name
}

// Stage puts rest to the model stage
func (rest *Rest) Stage(stage *Stage) *Rest {
	__gong__stage(stage.Rests, stage.Rest_stagedOrder, stage.Rest_orderStaged, &stage.RestOrder, stage.Rests_mapString, rest, rest.Name)
	return rest
}

// StagePreserveOrder puts rest to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RestOrder
// - update stage.RestOrder accordingly
func (rest *Rest) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Rests, stage.Rest_stagedOrder, stage.Rest_orderStaged, &stage.RestOrder, stage.Rests_mapString, rest, order, rest.Name)
}

// Unstage removes rest off the model stage
func (rest *Rest) Unstage(stage *Stage) *Rest {
	__gong__unstage(stage.Rests, stage.Rests_mapString, rest, rest.Name)
	return rest
}

// UnstageVoid removes rest off the model stage
func (rest *Rest) UnstageVoid(stage *Stage) {
	rest.Unstage(stage)
}

func (rest *Rest) StageVoid(stage *Stage) {
	rest.Stage(stage)
}

// for satisfaction of GongStruct interface
func (rest *Rest) GetName() (res string) {
	return rest.Name
}

// for satisfaction of GongStruct interface
func (rest *Rest) SetName(name string) {
	rest.Name = name
}

// Stage puts slide to the model stage
func (slide *Slide) Stage(stage *Stage) *Slide {
	__gong__stage(stage.Slides, stage.Slide_stagedOrder, stage.Slide_orderStaged, &stage.SlideOrder, stage.Slides_mapString, slide, slide.Name)
	return slide
}

// StagePreserveOrder puts slide to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SlideOrder
// - update stage.SlideOrder accordingly
func (slide *Slide) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Slides, stage.Slide_stagedOrder, stage.Slide_orderStaged, &stage.SlideOrder, stage.Slides_mapString, slide, order, slide.Name)
}

// Unstage removes slide off the model stage
func (slide *Slide) Unstage(stage *Stage) *Slide {
	__gong__unstage(stage.Slides, stage.Slides_mapString, slide, slide.Name)
	return slide
}

// UnstageVoid removes slide off the model stage
func (slide *Slide) UnstageVoid(stage *Stage) {
	slide.Unstage(stage)
}

func (slide *Slide) StageVoid(stage *Stage) {
	slide.Stage(stage)
}

// for satisfaction of GongStruct interface
func (slide *Slide) GetName() (res string) {
	return slide.Name
}

// for satisfaction of GongStruct interface
func (slide *Slide) SetName(name string) {
	slide.Name = name
}

// Stage puts slur to the model stage
func (slur *Slur) Stage(stage *Stage) *Slur {
	__gong__stage(stage.Slurs, stage.Slur_stagedOrder, stage.Slur_orderStaged, &stage.SlurOrder, stage.Slurs_mapString, slur, slur.Name)
	return slur
}

// StagePreserveOrder puts slur to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SlurOrder
// - update stage.SlurOrder accordingly
func (slur *Slur) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Slurs, stage.Slur_stagedOrder, stage.Slur_orderStaged, &stage.SlurOrder, stage.Slurs_mapString, slur, order, slur.Name)
}

// Unstage removes slur off the model stage
func (slur *Slur) Unstage(stage *Stage) *Slur {
	__gong__unstage(stage.Slurs, stage.Slurs_mapString, slur, slur.Name)
	return slur
}

// UnstageVoid removes slur off the model stage
func (slur *Slur) UnstageVoid(stage *Stage) {
	slur.Unstage(stage)
}

func (slur *Slur) StageVoid(stage *Stage) {
	slur.Stage(stage)
}

// for satisfaction of GongStruct interface
func (slur *Slur) GetName() (res string) {
	return slur.Name
}

// for satisfaction of GongStruct interface
func (slur *Slur) SetName(name string) {
	slur.Name = name
}

// Stage puts stem to the model stage
func (stem *Stem) Stage(stage *Stage) *Stem {
	__gong__stage(stage.Stems, stage.Stem_stagedOrder, stage.Stem_orderStaged, &stage.StemOrder, stage.Stems_mapString, stem, stem.Name)
	return stem
}

// StagePreserveOrder puts stem to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StemOrder
// - update stage.StemOrder accordingly
func (stem *Stem) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Stems, stage.Stem_stagedOrder, stage.Stem_orderStaged, &stage.StemOrder, stage.Stems_mapString, stem, order, stem.Name)
}

// Unstage removes stem off the model stage
func (stem *Stem) Unstage(stage *Stage) *Stem {
	__gong__unstage(stage.Stems, stage.Stems_mapString, stem, stem.Name)
	return stem
}

// UnstageVoid removes stem off the model stage
func (stem *Stem) UnstageVoid(stage *Stage) {
	stem.Unstage(stage)
}

func (stem *Stem) StageVoid(stage *Stage) {
	stem.Stage(stage)
}

// for satisfaction of GongStruct interface
func (stem *Stem) GetName() (res string) {
	return stem.Name
}

// for satisfaction of GongStruct interface
func (stem *Stem) SetName(name string) {
	stem.Name = name
}

// Stage puts strongaccent to the model stage
func (strongaccent *StrongAccent) Stage(stage *Stage) *StrongAccent {
	__gong__stage(stage.StrongAccents, stage.StrongAccent_stagedOrder, stage.StrongAccent_orderStaged, &stage.StrongAccentOrder, stage.StrongAccents_mapString, strongaccent, strongaccent.Name)
	return strongaccent
}

// StagePreserveOrder puts strongaccent to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StrongAccentOrder
// - update stage.StrongAccentOrder accordingly
func (strongaccent *StrongAccent) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StrongAccents, stage.StrongAccent_stagedOrder, stage.StrongAccent_orderStaged, &stage.StrongAccentOrder, stage.StrongAccents_mapString, strongaccent, order, strongaccent.Name)
}

// Unstage removes strongaccent off the model stage
func (strongaccent *StrongAccent) Unstage(stage *Stage) *StrongAccent {
	__gong__unstage(stage.StrongAccents, stage.StrongAccents_mapString, strongaccent, strongaccent.Name)
	return strongaccent
}

// UnstageVoid removes strongaccent off the model stage
func (strongaccent *StrongAccent) UnstageVoid(stage *Stage) {
	strongaccent.Unstage(stage)
}

func (strongaccent *StrongAccent) StageVoid(stage *Stage) {
	strongaccent.Stage(stage)
}

// for satisfaction of GongStruct interface
func (strongaccent *StrongAccent) GetName() (res string) {
	return strongaccent.Name
}

// for satisfaction of GongStruct interface
func (strongaccent *StrongAccent) SetName(name string) {
	strongaccent.Name = name
}

// Stage puts tap to the model stage
func (tap *Tap) Stage(stage *Stage) *Tap {
	__gong__stage(stage.Taps, stage.Tap_stagedOrder, stage.Tap_orderStaged, &stage.TapOrder, stage.Taps_mapString, tap, tap.Name)
	return tap
}

// StagePreserveOrder puts tap to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TapOrder
// - update stage.TapOrder accordingly
func (tap *Tap) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Taps, stage.Tap_stagedOrder, stage.Tap_orderStaged, &stage.TapOrder, stage.Taps_mapString, tap, order, tap.Name)
}

// Unstage removes tap off the model stage
func (tap *Tap) Unstage(stage *Stage) *Tap {
	__gong__unstage(stage.Taps, stage.Taps_mapString, tap, tap.Name)
	return tap
}

// UnstageVoid removes tap off the model stage
func (tap *Tap) UnstageVoid(stage *Stage) {
	tap.Unstage(stage)
}

func (tap *Tap) StageVoid(stage *Stage) {
	tap.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tap *Tap) GetName() (res string) {
	return tap.Name
}

// for satisfaction of GongStruct interface
func (tap *Tap) SetName(name string) {
	tap.Name = name
}

// Stage puts technical to the model stage
func (technical *Technical) Stage(stage *Stage) *Technical {
	__gong__stage(stage.Technicals, stage.Technical_stagedOrder, stage.Technical_orderStaged, &stage.TechnicalOrder, stage.Technicals_mapString, technical, technical.Name)
	return technical
}

// StagePreserveOrder puts technical to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TechnicalOrder
// - update stage.TechnicalOrder accordingly
func (technical *Technical) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Technicals, stage.Technical_stagedOrder, stage.Technical_orderStaged, &stage.TechnicalOrder, stage.Technicals_mapString, technical, order, technical.Name)
}

// Unstage removes technical off the model stage
func (technical *Technical) Unstage(stage *Stage) *Technical {
	__gong__unstage(stage.Technicals, stage.Technicals_mapString, technical, technical.Name)
	return technical
}

// UnstageVoid removes technical off the model stage
func (technical *Technical) UnstageVoid(stage *Stage) {
	technical.Unstage(stage)
}

func (technical *Technical) StageVoid(stage *Stage) {
	technical.Stage(stage)
}

// for satisfaction of GongStruct interface
func (technical *Technical) GetName() (res string) {
	return technical.Name
}

// for satisfaction of GongStruct interface
func (technical *Technical) SetName(name string) {
	technical.Name = name
}

// Stage puts textelementdata to the model stage
func (textelementdata *TextElementData) Stage(stage *Stage) *TextElementData {
	__gong__stage(stage.TextElementDatas, stage.TextElementData_stagedOrder, stage.TextElementData_orderStaged, &stage.TextElementDataOrder, stage.TextElementDatas_mapString, textelementdata, textelementdata.Name)
	return textelementdata
}

// StagePreserveOrder puts textelementdata to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextElementDataOrder
// - update stage.TextElementDataOrder accordingly
func (textelementdata *TextElementData) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TextElementDatas, stage.TextElementData_stagedOrder, stage.TextElementData_orderStaged, &stage.TextElementDataOrder, stage.TextElementDatas_mapString, textelementdata, order, textelementdata.Name)
}

// Unstage removes textelementdata off the model stage
func (textelementdata *TextElementData) Unstage(stage *Stage) *TextElementData {
	__gong__unstage(stage.TextElementDatas, stage.TextElementDatas_mapString, textelementdata, textelementdata.Name)
	return textelementdata
}

// UnstageVoid removes textelementdata off the model stage
func (textelementdata *TextElementData) UnstageVoid(stage *Stage) {
	textelementdata.Unstage(stage)
}

func (textelementdata *TextElementData) StageVoid(stage *Stage) {
	textelementdata.Stage(stage)
}

// for satisfaction of GongStruct interface
func (textelementdata *TextElementData) GetName() (res string) {
	return textelementdata.Name
}

// for satisfaction of GongStruct interface
func (textelementdata *TextElementData) SetName(name string) {
	textelementdata.Name = name
}

// Stage puts tie to the model stage
func (tie *Tie) Stage(stage *Stage) *Tie {
	__gong__stage(stage.Ties, stage.Tie_stagedOrder, stage.Tie_orderStaged, &stage.TieOrder, stage.Ties_mapString, tie, tie.Name)
	return tie
}

// StagePreserveOrder puts tie to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TieOrder
// - update stage.TieOrder accordingly
func (tie *Tie) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Ties, stage.Tie_stagedOrder, stage.Tie_orderStaged, &stage.TieOrder, stage.Ties_mapString, tie, order, tie.Name)
}

// Unstage removes tie off the model stage
func (tie *Tie) Unstage(stage *Stage) *Tie {
	__gong__unstage(stage.Ties, stage.Ties_mapString, tie, tie.Name)
	return tie
}

// UnstageVoid removes tie off the model stage
func (tie *Tie) UnstageVoid(stage *Stage) {
	tie.Unstage(stage)
}

func (tie *Tie) StageVoid(stage *Stage) {
	tie.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tie *Tie) GetName() (res string) {
	return tie.Name
}

// for satisfaction of GongStruct interface
func (tie *Tie) SetName(name string) {
	tie.Name = name
}

// Stage puts tied to the model stage
func (tied *Tied) Stage(stage *Stage) *Tied {
	__gong__stage(stage.Tieds, stage.Tied_stagedOrder, stage.Tied_orderStaged, &stage.TiedOrder, stage.Tieds_mapString, tied, tied.Name)
	return tied
}

// StagePreserveOrder puts tied to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TiedOrder
// - update stage.TiedOrder accordingly
func (tied *Tied) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Tieds, stage.Tied_stagedOrder, stage.Tied_orderStaged, &stage.TiedOrder, stage.Tieds_mapString, tied, order, tied.Name)
}

// Unstage removes tied off the model stage
func (tied *Tied) Unstage(stage *Stage) *Tied {
	__gong__unstage(stage.Tieds, stage.Tieds_mapString, tied, tied.Name)
	return tied
}

// UnstageVoid removes tied off the model stage
func (tied *Tied) UnstageVoid(stage *Stage) {
	tied.Unstage(stage)
}

func (tied *Tied) StageVoid(stage *Stage) {
	tied.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tied *Tied) GetName() (res string) {
	return tied.Name
}

// for satisfaction of GongStruct interface
func (tied *Tied) SetName(name string) {
	tied.Name = name
}

// Stage puts tremolo to the model stage
func (tremolo *Tremolo) Stage(stage *Stage) *Tremolo {
	__gong__stage(stage.Tremolos, stage.Tremolo_stagedOrder, stage.Tremolo_orderStaged, &stage.TremoloOrder, stage.Tremolos_mapString, tremolo, tremolo.Name)
	return tremolo
}

// StagePreserveOrder puts tremolo to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TremoloOrder
// - update stage.TremoloOrder accordingly
func (tremolo *Tremolo) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Tremolos, stage.Tremolo_stagedOrder, stage.Tremolo_orderStaged, &stage.TremoloOrder, stage.Tremolos_mapString, tremolo, order, tremolo.Name)
}

// Unstage removes tremolo off the model stage
func (tremolo *Tremolo) Unstage(stage *Stage) *Tremolo {
	__gong__unstage(stage.Tremolos, stage.Tremolos_mapString, tremolo, tremolo.Name)
	return tremolo
}

// UnstageVoid removes tremolo off the model stage
func (tremolo *Tremolo) UnstageVoid(stage *Stage) {
	tremolo.Unstage(stage)
}

func (tremolo *Tremolo) StageVoid(stage *Stage) {
	tremolo.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tremolo *Tremolo) GetName() (res string) {
	return tremolo.Name
}

// for satisfaction of GongStruct interface
func (tremolo *Tremolo) SetName(name string) {
	tremolo.Name = name
}

// Stage puts tuplet to the model stage
func (tuplet *Tuplet) Stage(stage *Stage) *Tuplet {
	__gong__stage(stage.Tuplets, stage.Tuplet_stagedOrder, stage.Tuplet_orderStaged, &stage.TupletOrder, stage.Tuplets_mapString, tuplet, tuplet.Name)
	return tuplet
}

// StagePreserveOrder puts tuplet to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TupletOrder
// - update stage.TupletOrder accordingly
func (tuplet *Tuplet) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Tuplets, stage.Tuplet_stagedOrder, stage.Tuplet_orderStaged, &stage.TupletOrder, stage.Tuplets_mapString, tuplet, order, tuplet.Name)
}

// Unstage removes tuplet off the model stage
func (tuplet *Tuplet) Unstage(stage *Stage) *Tuplet {
	__gong__unstage(stage.Tuplets, stage.Tuplets_mapString, tuplet, tuplet.Name)
	return tuplet
}

// UnstageVoid removes tuplet off the model stage
func (tuplet *Tuplet) UnstageVoid(stage *Stage) {
	tuplet.Unstage(stage)
}

func (tuplet *Tuplet) StageVoid(stage *Stage) {
	tuplet.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tuplet *Tuplet) GetName() (res string) {
	return tuplet.Name
}

// for satisfaction of GongStruct interface
func (tuplet *Tuplet) SetName(name string) {
	tuplet.Name = name
}

// Stage puts tupletdot to the model stage
func (tupletdot *TupletDot) Stage(stage *Stage) *TupletDot {
	__gong__stage(stage.TupletDots, stage.TupletDot_stagedOrder, stage.TupletDot_orderStaged, &stage.TupletDotOrder, stage.TupletDots_mapString, tupletdot, tupletdot.Name)
	return tupletdot
}

// StagePreserveOrder puts tupletdot to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TupletDotOrder
// - update stage.TupletDotOrder accordingly
func (tupletdot *TupletDot) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TupletDots, stage.TupletDot_stagedOrder, stage.TupletDot_orderStaged, &stage.TupletDotOrder, stage.TupletDots_mapString, tupletdot, order, tupletdot.Name)
}

// Unstage removes tupletdot off the model stage
func (tupletdot *TupletDot) Unstage(stage *Stage) *TupletDot {
	__gong__unstage(stage.TupletDots, stage.TupletDots_mapString, tupletdot, tupletdot.Name)
	return tupletdot
}

// UnstageVoid removes tupletdot off the model stage
func (tupletdot *TupletDot) UnstageVoid(stage *Stage) {
	tupletdot.Unstage(stage)
}

func (tupletdot *TupletDot) StageVoid(stage *Stage) {
	tupletdot.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tupletdot *TupletDot) GetName() (res string) {
	return tupletdot.Name
}

// for satisfaction of GongStruct interface
func (tupletdot *TupletDot) SetName(name string) {
	tupletdot.Name = name
}

// Stage puts tupletnumber to the model stage
func (tupletnumber *TupletNumber) Stage(stage *Stage) *TupletNumber {
	__gong__stage(stage.TupletNumbers, stage.TupletNumber_stagedOrder, stage.TupletNumber_orderStaged, &stage.TupletNumberOrder, stage.TupletNumbers_mapString, tupletnumber, tupletnumber.Name)
	return tupletnumber
}

// StagePreserveOrder puts tupletnumber to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TupletNumberOrder
// - update stage.TupletNumberOrder accordingly
func (tupletnumber *TupletNumber) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TupletNumbers, stage.TupletNumber_stagedOrder, stage.TupletNumber_orderStaged, &stage.TupletNumberOrder, stage.TupletNumbers_mapString, tupletnumber, order, tupletnumber.Name)
}

// Unstage removes tupletnumber off the model stage
func (tupletnumber *TupletNumber) Unstage(stage *Stage) *TupletNumber {
	__gong__unstage(stage.TupletNumbers, stage.TupletNumbers_mapString, tupletnumber, tupletnumber.Name)
	return tupletnumber
}

// UnstageVoid removes tupletnumber off the model stage
func (tupletnumber *TupletNumber) UnstageVoid(stage *Stage) {
	tupletnumber.Unstage(stage)
}

func (tupletnumber *TupletNumber) StageVoid(stage *Stage) {
	tupletnumber.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tupletnumber *TupletNumber) GetName() (res string) {
	return tupletnumber.Name
}

// for satisfaction of GongStruct interface
func (tupletnumber *TupletNumber) SetName(name string) {
	tupletnumber.Name = name
}

// Stage puts tupletportion to the model stage
func (tupletportion *TupletPortion) Stage(stage *Stage) *TupletPortion {
	__gong__stage(stage.TupletPortions, stage.TupletPortion_stagedOrder, stage.TupletPortion_orderStaged, &stage.TupletPortionOrder, stage.TupletPortions_mapString, tupletportion, tupletportion.Name)
	return tupletportion
}

// StagePreserveOrder puts tupletportion to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TupletPortionOrder
// - update stage.TupletPortionOrder accordingly
func (tupletportion *TupletPortion) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TupletPortions, stage.TupletPortion_stagedOrder, stage.TupletPortion_orderStaged, &stage.TupletPortionOrder, stage.TupletPortions_mapString, tupletportion, order, tupletportion.Name)
}

// Unstage removes tupletportion off the model stage
func (tupletportion *TupletPortion) Unstage(stage *Stage) *TupletPortion {
	__gong__unstage(stage.TupletPortions, stage.TupletPortions_mapString, tupletportion, tupletportion.Name)
	return tupletportion
}

// UnstageVoid removes tupletportion off the model stage
func (tupletportion *TupletPortion) UnstageVoid(stage *Stage) {
	tupletportion.Unstage(stage)
}

func (tupletportion *TupletPortion) StageVoid(stage *Stage) {
	tupletportion.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tupletportion *TupletPortion) GetName() (res string) {
	return tupletportion.Name
}

// for satisfaction of GongStruct interface
func (tupletportion *TupletPortion) SetName(name string) {
	tupletportion.Name = name
}

// Stage puts tuplettype to the model stage
func (tuplettype *TupletType) Stage(stage *Stage) *TupletType {
	__gong__stage(stage.TupletTypes, stage.TupletType_stagedOrder, stage.TupletType_orderStaged, &stage.TupletTypeOrder, stage.TupletTypes_mapString, tuplettype, tuplettype.Name)
	return tuplettype
}

// StagePreserveOrder puts tuplettype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TupletTypeOrder
// - update stage.TupletTypeOrder accordingly
func (tuplettype *TupletType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TupletTypes, stage.TupletType_stagedOrder, stage.TupletType_orderStaged, &stage.TupletTypeOrder, stage.TupletTypes_mapString, tuplettype, order, tuplettype.Name)
}

// Unstage removes tuplettype off the model stage
func (tuplettype *TupletType) Unstage(stage *Stage) *TupletType {
	__gong__unstage(stage.TupletTypes, stage.TupletTypes_mapString, tuplettype, tuplettype.Name)
	return tuplettype
}

// UnstageVoid removes tuplettype off the model stage
func (tuplettype *TupletType) UnstageVoid(stage *Stage) {
	tuplettype.Unstage(stage)
}

func (tuplettype *TupletType) StageVoid(stage *Stage) {
	tuplettype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tuplettype *TupletType) GetName() (res string) {
	return tuplettype.Name
}

// for satisfaction of GongStruct interface
func (tuplettype *TupletType) SetName(name string) {
	tuplettype.Name = name
}

// Stage puts unpitched to the model stage
func (unpitched *Unpitched) Stage(stage *Stage) *Unpitched {
	__gong__stage(stage.Unpitcheds, stage.Unpitched_stagedOrder, stage.Unpitched_orderStaged, &stage.UnpitchedOrder, stage.Unpitcheds_mapString, unpitched, unpitched.Name)
	return unpitched
}

// StagePreserveOrder puts unpitched to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UnpitchedOrder
// - update stage.UnpitchedOrder accordingly
func (unpitched *Unpitched) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Unpitcheds, stage.Unpitched_stagedOrder, stage.Unpitched_orderStaged, &stage.UnpitchedOrder, stage.Unpitcheds_mapString, unpitched, order, unpitched.Name)
}

// Unstage removes unpitched off the model stage
func (unpitched *Unpitched) Unstage(stage *Stage) *Unpitched {
	__gong__unstage(stage.Unpitcheds, stage.Unpitcheds_mapString, unpitched, unpitched.Name)
	return unpitched
}

// UnstageVoid removes unpitched off the model stage
func (unpitched *Unpitched) UnstageVoid(stage *Stage) {
	unpitched.Unstage(stage)
}

func (unpitched *Unpitched) StageVoid(stage *Stage) {
	unpitched.Stage(stage)
}

// for satisfaction of GongStruct interface
func (unpitched *Unpitched) GetName() (res string) {
	return unpitched.Name
}

// for satisfaction of GongStruct interface
func (unpitched *Unpitched) SetName(name string) {
	unpitched.Name = name
}

// Stage puts wait to the model stage
func (wait *Wait) Stage(stage *Stage) *Wait {
	__gong__stage(stage.Waits, stage.Wait_stagedOrder, stage.Wait_orderStaged, &stage.WaitOrder, stage.Waits_mapString, wait, wait.Name)
	return wait
}

// StagePreserveOrder puts wait to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.WaitOrder
// - update stage.WaitOrder accordingly
func (wait *Wait) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Waits, stage.Wait_stagedOrder, stage.Wait_orderStaged, &stage.WaitOrder, stage.Waits_mapString, wait, order, wait.Name)
}

// Unstage removes wait off the model stage
func (wait *Wait) Unstage(stage *Stage) *Wait {
	__gong__unstage(stage.Waits, stage.Waits_mapString, wait, wait.Name)
	return wait
}

// UnstageVoid removes wait off the model stage
func (wait *Wait) UnstageVoid(stage *Stage) {
	wait.Unstage(stage)
}

func (wait *Wait) StageVoid(stage *Stage) {
	wait.Stage(stage)
}

// for satisfaction of GongStruct interface
func (wait *Wait) GetName() (res string) {
	return wait.Name
}

// for satisfaction of GongStruct interface
func (wait *Wait) SetName(name string) {
	wait.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Accidentals, &stage.Accidentals_mapString, &stage.Accidental_stagedOrder, &stage.AccidentalOrder)

	__gong__resetStageType(&stage.AccidentalMarks, &stage.AccidentalMarks_mapString, &stage.AccidentalMark_stagedOrder, &stage.AccidentalMarkOrder)

	__gong__resetStageType(&stage.Arpeggiates, &stage.Arpeggiates_mapString, &stage.Arpeggiate_stagedOrder, &stage.ArpeggiateOrder)

	__gong__resetStageType(&stage.Arrows, &stage.Arrows_mapString, &stage.Arrow_stagedOrder, &stage.ArrowOrder)

	__gong__resetStageType(&stage.Articulationss, &stage.Articulationss_mapString, &stage.Articulations_stagedOrder, &stage.ArticulationsOrder)

	__gong__resetStageType(&stage.Assesss, &stage.Assesss_mapString, &stage.Assess_stagedOrder, &stage.AssessOrder)

	__gong__resetStageType(&stage.Backups, &stage.Backups_mapString, &stage.Backup_stagedOrder, &stage.BackupOrder)

	__gong__resetStageType(&stage.Beams, &stage.Beams_mapString, &stage.Beam_stagedOrder, &stage.BeamOrder)

	__gong__resetStageType(&stage.Bends, &stage.Bends_mapString, &stage.Bend_stagedOrder, &stage.BendOrder)

	__gong__resetStageType(&stage.BreathMarks, &stage.BreathMarks_mapString, &stage.BreathMark_stagedOrder, &stage.BreathMarkOrder)

	__gong__resetStageType(&stage.Caesuras, &stage.Caesuras_mapString, &stage.Caesura_stagedOrder, &stage.CaesuraOrder)

	__gong__resetStageType(&stage.Elisions, &stage.Elisions_mapString, &stage.Elision_stagedOrder, &stage.ElisionOrder)

	__gong__resetStageType(&stage.EmptyLines, &stage.EmptyLines_mapString, &stage.EmptyLine_stagedOrder, &stage.EmptyLineOrder)

	__gong__resetStageType(&stage.Extends, &stage.Extends_mapString, &stage.Extend_stagedOrder, &stage.ExtendOrder)

	__gong__resetStageType(&stage.Figures, &stage.Figures_mapString, &stage.Figure_stagedOrder, &stage.FigureOrder)

	__gong__resetStageType(&stage.FiguredBasss, &stage.FiguredBasss_mapString, &stage.FiguredBass_stagedOrder, &stage.FiguredBassOrder)

	__gong__resetStageType(&stage.Forwards, &stage.Forwards_mapString, &stage.Forward_stagedOrder, &stage.ForwardOrder)

	__gong__resetStageType(&stage.Glissandos, &stage.Glissandos_mapString, &stage.Glissando_stagedOrder, &stage.GlissandoOrder)

	__gong__resetStageType(&stage.Graces, &stage.Graces_mapString, &stage.Grace_stagedOrder, &stage.GraceOrder)

	__gong__resetStageType(&stage.HammerOnPullOffs, &stage.HammerOnPullOffs_mapString, &stage.HammerOnPullOff_stagedOrder, &stage.HammerOnPullOffOrder)

	__gong__resetStageType(&stage.Handbells, &stage.Handbells_mapString, &stage.Handbell_stagedOrder, &stage.HandbellOrder)

	__gong__resetStageType(&stage.HarmonCloseds, &stage.HarmonCloseds_mapString, &stage.HarmonClosed_stagedOrder, &stage.HarmonClosedOrder)

	__gong__resetStageType(&stage.HarmonMutes, &stage.HarmonMutes_mapString, &stage.HarmonMute_stagedOrder, &stage.HarmonMuteOrder)

	__gong__resetStageType(&stage.Harmonics, &stage.Harmonics_mapString, &stage.Harmonic_stagedOrder, &stage.HarmonicOrder)

	__gong__resetStageType(&stage.HeelToes, &stage.HeelToes_mapString, &stage.HeelToe_stagedOrder, &stage.HeelToeOrder)

	__gong__resetStageType(&stage.Holes, &stage.Holes_mapString, &stage.Hole_stagedOrder, &stage.HoleOrder)

	__gong__resetStageType(&stage.HoleCloseds, &stage.HoleCloseds_mapString, &stage.HoleClosed_stagedOrder, &stage.HoleClosedOrder)

	__gong__resetStageType(&stage.Instruments, &stage.Instruments_mapString, &stage.Instrument_stagedOrder, &stage.InstrumentOrder)

	__gong__resetStageType(&stage.Listens, &stage.Listens_mapString, &stage.Listen_stagedOrder, &stage.ListenOrder)

	__gong__resetStageType(&stage.Lyrics, &stage.Lyrics_mapString, &stage.Lyric_stagedOrder, &stage.LyricOrder)

	__gong__resetStageType(&stage.Mordents, &stage.Mordents_mapString, &stage.Mordent_stagedOrder, &stage.MordentOrder)

	__gong__resetStageType(&stage.NonArpeggiates, &stage.NonArpeggiates_mapString, &stage.NonArpeggiate_stagedOrder, &stage.NonArpeggiateOrder)

	__gong__resetStageType(&stage.Notationss, &stage.Notationss_mapString, &stage.Notations_stagedOrder, &stage.NotationsOrder)

	__gong__resetStageType(&stage.Notes, &stage.Notes_mapString, &stage.Note_stagedOrder, &stage.NoteOrder)

	__gong__resetStageType(&stage.NoteTypes, &stage.NoteTypes_mapString, &stage.NoteType_stagedOrder, &stage.NoteTypeOrder)

	__gong__resetStageType(&stage.Noteheads, &stage.Noteheads_mapString, &stage.Notehead_stagedOrder, &stage.NoteheadOrder)

	__gong__resetStageType(&stage.NoteheadTexts, &stage.NoteheadTexts_mapString, &stage.NoteheadText_stagedOrder, &stage.NoteheadTextOrder)

	__gong__resetStageType(&stage.Ornamentss, &stage.Ornamentss_mapString, &stage.Ornaments_stagedOrder, &stage.OrnamentsOrder)

	__gong__resetStageType(&stage.OtherNotations, &stage.OtherNotations_mapString, &stage.OtherNotation_stagedOrder, &stage.OtherNotationOrder)

	__gong__resetStageType(&stage.OtherPlacementTexts, &stage.OtherPlacementTexts_mapString, &stage.OtherPlacementText_stagedOrder, &stage.OtherPlacementTextOrder)

	__gong__resetStageType(&stage.Pitchs, &stage.Pitchs_mapString, &stage.Pitch_stagedOrder, &stage.PitchOrder)

	__gong__resetStageType(&stage.PlacementTexts, &stage.PlacementTexts_mapString, &stage.PlacementText_stagedOrder, &stage.PlacementTextOrder)

	__gong__resetStageType(&stage.Releases, &stage.Releases_mapString, &stage.Release_stagedOrder, &stage.ReleaseOrder)

	__gong__resetStageType(&stage.Rests, &stage.Rests_mapString, &stage.Rest_stagedOrder, &stage.RestOrder)

	__gong__resetStageType(&stage.Slides, &stage.Slides_mapString, &stage.Slide_stagedOrder, &stage.SlideOrder)

	__gong__resetStageType(&stage.Slurs, &stage.Slurs_mapString, &stage.Slur_stagedOrder, &stage.SlurOrder)

	__gong__resetStageType(&stage.Stems, &stage.Stems_mapString, &stage.Stem_stagedOrder, &stage.StemOrder)

	__gong__resetStageType(&stage.StrongAccents, &stage.StrongAccents_mapString, &stage.StrongAccent_stagedOrder, &stage.StrongAccentOrder)

	__gong__resetStageType(&stage.Taps, &stage.Taps_mapString, &stage.Tap_stagedOrder, &stage.TapOrder)

	__gong__resetStageType(&stage.Technicals, &stage.Technicals_mapString, &stage.Technical_stagedOrder, &stage.TechnicalOrder)

	__gong__resetStageType(&stage.TextElementDatas, &stage.TextElementDatas_mapString, &stage.TextElementData_stagedOrder, &stage.TextElementDataOrder)

	__gong__resetStageType(&stage.Ties, &stage.Ties_mapString, &stage.Tie_stagedOrder, &stage.TieOrder)

	__gong__resetStageType(&stage.Tieds, &stage.Tieds_mapString, &stage.Tied_stagedOrder, &stage.TiedOrder)

	__gong__resetStageType(&stage.Tremolos, &stage.Tremolos_mapString, &stage.Tremolo_stagedOrder, &stage.TremoloOrder)

	__gong__resetStageType(&stage.Tuplets, &stage.Tuplets_mapString, &stage.Tuplet_stagedOrder, &stage.TupletOrder)

	__gong__resetStageType(&stage.TupletDots, &stage.TupletDots_mapString, &stage.TupletDot_stagedOrder, &stage.TupletDotOrder)

	__gong__resetStageType(&stage.TupletNumbers, &stage.TupletNumbers_mapString, &stage.TupletNumber_stagedOrder, &stage.TupletNumberOrder)

	__gong__resetStageType(&stage.TupletPortions, &stage.TupletPortions_mapString, &stage.TupletPortion_stagedOrder, &stage.TupletPortionOrder)

	__gong__resetStageType(&stage.TupletTypes, &stage.TupletTypes_mapString, &stage.TupletType_stagedOrder, &stage.TupletTypeOrder)

	__gong__resetStageType(&stage.Unpitcheds, &stage.Unpitcheds_mapString, &stage.Unpitched_stagedOrder, &stage.UnpitchedOrder)

	__gong__resetStageType(&stage.Waits, &stage.Waits_mapString, &stage.Wait_stagedOrder, &stage.WaitOrder)

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
func (Accidental) GongGetAssociationName() any {
	return &Accidental{
	}
}

func (AccidentalMark) GongGetAssociationName() any {
	return &AccidentalMark{
	}
}

func (Arpeggiate) GongGetAssociationName() any {
	return &Arpeggiate{
	}
}

func (Arrow) GongGetAssociationName() any {
	return &Arrow{
	}
}

func (Articulations) GongGetAssociationName() any {
	return &Articulations{
			StrongAccent: []*StrongAccent{{Name: "StrongAccent"}},
			Scoop: []*EmptyLine{{Name: "Scoop"}},
			Plop: []*EmptyLine{{Name: "Plop"}},
			Doit: []*EmptyLine{{Name: "Doit"}},
			Falloff: []*EmptyLine{{Name: "Falloff"}},
			BreathMark: []*BreathMark{{Name: "BreathMark"}},
			Caesura: []*Caesura{{Name: "Caesura"}},
			OtherArticulation: []*OtherPlacementText{{Name: "OtherArticulation"}},
	}
}

func (Assess) GongGetAssociationName() any {
	return &Assess{
	}
}

func (Backup) GongGetAssociationName() any {
	return &Backup{
	}
}

func (Beam) GongGetAssociationName() any {
	return &Beam{
	}
}

func (Bend) GongGetAssociationName() any {
	return &Bend{
			Release: &Release{Name: "Release"},
			WithBar: &PlacementText{Name: "WithBar"},
	}
}

func (BreathMark) GongGetAssociationName() any {
	return &BreathMark{
	}
}

func (Caesura) GongGetAssociationName() any {
	return &Caesura{
	}
}

func (Elision) GongGetAssociationName() any {
	return &Elision{
	}
}

func (EmptyLine) GongGetAssociationName() any {
	return &EmptyLine{
	}
}

func (Extend) GongGetAssociationName() any {
	return &Extend{
	}
}

func (Figure) GongGetAssociationName() any {
	return &Figure{
			Extend: &Extend{Name: "Extend"},
	}
}

func (FiguredBass) GongGetAssociationName() any {
	return &FiguredBass{
			Figure: []*Figure{{Name: "Figure"}},
	}
}

func (Forward) GongGetAssociationName() any {
	return &Forward{
	}
}

func (Glissando) GongGetAssociationName() any {
	return &Glissando{
	}
}

func (Grace) GongGetAssociationName() any {
	return &Grace{
	}
}

func (HammerOnPullOff) GongGetAssociationName() any {
	return &HammerOnPullOff{
	}
}

func (Handbell) GongGetAssociationName() any {
	return &Handbell{
	}
}

func (HarmonClosed) GongGetAssociationName() any {
	return &HarmonClosed{
	}
}

func (HarmonMute) GongGetAssociationName() any {
	return &HarmonMute{
			HarmonClosed: &HarmonClosed{Name: "HarmonClosed"},
	}
}

func (Harmonic) GongGetAssociationName() any {
	return &Harmonic{
	}
}

func (HeelToe) GongGetAssociationName() any {
	return &HeelToe{
	}
}

func (Hole) GongGetAssociationName() any {
	return &Hole{
			HoleClosed: &HoleClosed{Name: "HoleClosed"},
	}
}

func (HoleClosed) GongGetAssociationName() any {
	return &HoleClosed{
	}
}

func (Instrument) GongGetAssociationName() any {
	return &Instrument{
	}
}

func (Listen) GongGetAssociationName() any {
	return &Listen{
			Assess: []*Assess{{Name: "Assess"}},
			Wait: []*Wait{{Name: "Wait"}},
	}
}

func (Lyric) GongGetAssociationName() any {
	return &Lyric{
			Text: &TextElementData{Name: "Text"},
			Elision: []*Elision{{Name: "Elision"}},
			Extend: &Extend{Name: "Extend"},
	}
}

func (Mordent) GongGetAssociationName() any {
	return &Mordent{
	}
}

func (NonArpeggiate) GongGetAssociationName() any {
	return &NonArpeggiate{
	}
}

func (Notations) GongGetAssociationName() any {
	return &Notations{
			Tied: []*Tied{{Name: "Tied"}},
			Slur: []*Slur{{Name: "Slur"}},
			Tuplet: []*Tuplet{{Name: "Tuplet"}},
			Glissando: []*Glissando{{Name: "Glissando"}},
			Slide: []*Slide{{Name: "Slide"}},
			Ornaments: []*Ornaments{{Name: "Ornaments"}},
			Technical: []*Technical{{Name: "Technical"}},
			Articulations: []*Articulations{{Name: "Articulations"}},
			Arpeggiate: []*Arpeggiate{{Name: "Arpeggiate"}},
			NonArpeggiate: []*NonArpeggiate{{Name: "NonArpeggiate"}},
			AccidentalMark: []*AccidentalMark{{Name: "AccidentalMark"}},
			OtherNotation: []*OtherNotation{{Name: "OtherNotation"}},
	}
}

func (Note) GongGetAssociationName() any {
	return &Note{
			Grace: &Grace{Name: "Grace"},
			Pitch: &Pitch{Name: "Pitch"},
			Unpitched: &Unpitched{Name: "Unpitched"},
			Rest: &Rest{Name: "Rest"},
			Tie: []*Tie{{Name: "Tie"}},
			Instrument: []*Instrument{{Name: "Instrument"}},
			Type: &NoteType{Name: "Type"},
			Accidental: &Accidental{Name: "Accidental"},
			Stem: &Stem{Name: "Stem"},
			Notehead: &Notehead{Name: "Notehead"},
			NoteheadText: &NoteheadText{Name: "NoteheadText"},
			Beam: []*Beam{{Name: "Beam"}},
			Notations: []*Notations{{Name: "Notations"}},
			Lyric: []*Lyric{{Name: "Lyric"}},
			Listen: &Listen{Name: "Listen"},
	}
}

func (NoteType) GongGetAssociationName() any {
	return &NoteType{
	}
}

func (Notehead) GongGetAssociationName() any {
	return &Notehead{
	}
}

func (NoteheadText) GongGetAssociationName() any {
	return &NoteheadText{
	}
}

func (Ornaments) GongGetAssociationName() any {
	return &Ornaments{
			Mordent: []*Mordent{{Name: "Mordent"}},
			InvertedMordent: []*Mordent{{Name: "InvertedMordent"}},
			Tremolo: []*Tremolo{{Name: "Tremolo"}},
			OtherOrnament: []*OtherPlacementText{{Name: "OtherOrnament"}},
			AccidentalMark: []*AccidentalMark{{Name: "AccidentalMark"}},
	}
}

func (OtherNotation) GongGetAssociationName() any {
	return &OtherNotation{
	}
}

func (OtherPlacementText) GongGetAssociationName() any {
	return &OtherPlacementText{
	}
}

func (Pitch) GongGetAssociationName() any {
	return &Pitch{
	}
}

func (PlacementText) GongGetAssociationName() any {
	return &PlacementText{
	}
}

func (Release) GongGetAssociationName() any {
	return &Release{
	}
}

func (Rest) GongGetAssociationName() any {
	return &Rest{
	}
}

func (Slide) GongGetAssociationName() any {
	return &Slide{
	}
}

func (Slur) GongGetAssociationName() any {
	return &Slur{
	}
}

func (Stem) GongGetAssociationName() any {
	return &Stem{
	}
}

func (StrongAccent) GongGetAssociationName() any {
	return &StrongAccent{
	}
}

func (Tap) GongGetAssociationName() any {
	return &Tap{
	}
}

func (Technical) GongGetAssociationName() any {
	return &Technical{
			Harmonic: []*Harmonic{{Name: "Harmonic"}},
			Pluck: []*PlacementText{{Name: "Pluck"}},
			HammerOn: []*HammerOnPullOff{{Name: "HammerOn"}},
			PullOff: []*HammerOnPullOff{{Name: "PullOff"}},
			Bend: []*Bend{{Name: "Bend"}},
			Tap: []*Tap{{Name: "Tap"}},
			Heel: []*HeelToe{{Name: "Heel"}},
			Toe: []*HeelToe{{Name: "Toe"}},
			Hole: []*Hole{{Name: "Hole"}},
			Arrow: []*Arrow{{Name: "Arrow"}},
			Handbell: []*Handbell{{Name: "Handbell"}},
			HarmonMute: []*HarmonMute{{Name: "HarmonMute"}},
			OtherTechnical: []*OtherPlacementText{{Name: "OtherTechnical"}},
	}
}

func (TextElementData) GongGetAssociationName() any {
	return &TextElementData{
	}
}

func (Tie) GongGetAssociationName() any {
	return &Tie{
	}
}

func (Tied) GongGetAssociationName() any {
	return &Tied{
	}
}

func (Tremolo) GongGetAssociationName() any {
	return &Tremolo{
	}
}

func (Tuplet) GongGetAssociationName() any {
	return &Tuplet{
			TupletActual: &TupletPortion{Name: "TupletActual"},
			TupletNormal: &TupletPortion{Name: "TupletNormal"},
	}
}

func (TupletDot) GongGetAssociationName() any {
	return &TupletDot{
	}
}

func (TupletNumber) GongGetAssociationName() any {
	return &TupletNumber{
	}
}

func (TupletPortion) GongGetAssociationName() any {
	return &TupletPortion{
			TupletNumber: &TupletNumber{Name: "TupletNumber"},
			TupletType: &TupletType{Name: "TupletType"},
			TupletDot: []*TupletDot{{Name: "TupletDot"}},
	}
}

func (TupletType) GongGetAssociationName() any {
	return &TupletType{
	}
}

func (Unpitched) GongGetAssociationName() any {
	return &Unpitched{
	}
}

func (Wait) GongGetAssociationName() any {
	return &Wait{
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
	// reverse maps of direct associations of Accidental
	case Accidental:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AccidentalMark
	case AccidentalMark:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Arpeggiate
	case Arpeggiate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Arrow
	case Arrow:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Articulations
	case Articulations:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Assess
	case Assess:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Backup
	case Backup:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Beam
	case Beam:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bend
	case Bend:
		switch fieldname {
		// insertion point for per direct association field
		case "Release":
			res := make(map[*Release][]*Bend)
			for bend := range stage.Bends {
				if bend.Release != nil {
					release_ := bend.Release
					var bends []*Bend
					_, ok := res[release_]
					if ok {
						bends = res[release_]
					} else {
						bends = make([]*Bend, 0)
					}
					bends = append(bends, bend)
					res[release_] = bends
				}
			}
			return any(res).(map[*End][]*Start)
		case "WithBar":
			res := make(map[*PlacementText][]*Bend)
			for bend := range stage.Bends {
				if bend.WithBar != nil {
					placementtext_ := bend.WithBar
					var bends []*Bend
					_, ok := res[placementtext_]
					if ok {
						bends = res[placementtext_]
					} else {
						bends = make([]*Bend, 0)
					}
					bends = append(bends, bend)
					res[placementtext_] = bends
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BreathMark
	case BreathMark:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Caesura
	case Caesura:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Elision
	case Elision:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyLine
	case EmptyLine:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Extend
	case Extend:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Figure
	case Figure:
		switch fieldname {
		// insertion point for per direct association field
		case "Extend":
			res := make(map[*Extend][]*Figure)
			for figure := range stage.Figures {
				if figure.Extend != nil {
					extend_ := figure.Extend
					var figures []*Figure
					_, ok := res[extend_]
					if ok {
						figures = res[extend_]
					} else {
						figures = make([]*Figure, 0)
					}
					figures = append(figures, figure)
					res[extend_] = figures
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FiguredBass
	case FiguredBass:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Forward
	case Forward:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Glissando
	case Glissando:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Grace
	case Grace:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HammerOnPullOff
	case HammerOnPullOff:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Handbell
	case Handbell:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HarmonClosed
	case HarmonClosed:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HarmonMute
	case HarmonMute:
		switch fieldname {
		// insertion point for per direct association field
		case "HarmonClosed":
			res := make(map[*HarmonClosed][]*HarmonMute)
			for harmonmute := range stage.HarmonMutes {
				if harmonmute.HarmonClosed != nil {
					harmonclosed_ := harmonmute.HarmonClosed
					var harmonmutes []*HarmonMute
					_, ok := res[harmonclosed_]
					if ok {
						harmonmutes = res[harmonclosed_]
					} else {
						harmonmutes = make([]*HarmonMute, 0)
					}
					harmonmutes = append(harmonmutes, harmonmute)
					res[harmonclosed_] = harmonmutes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Harmonic
	case Harmonic:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HeelToe
	case HeelToe:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Hole
	case Hole:
		switch fieldname {
		// insertion point for per direct association field
		case "HoleClosed":
			res := make(map[*HoleClosed][]*Hole)
			for hole := range stage.Holes {
				if hole.HoleClosed != nil {
					holeclosed_ := hole.HoleClosed
					var holes []*Hole
					_, ok := res[holeclosed_]
					if ok {
						holes = res[holeclosed_]
					} else {
						holes = make([]*Hole, 0)
					}
					holes = append(holes, hole)
					res[holeclosed_] = holes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of HoleClosed
	case HoleClosed:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Instrument
	case Instrument:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Listen
	case Listen:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lyric
	case Lyric:
		switch fieldname {
		// insertion point for per direct association field
		case "Text":
			res := make(map[*TextElementData][]*Lyric)
			for lyric := range stage.Lyrics {
				if lyric.Text != nil {
					textelementdata_ := lyric.Text
					var lyrics []*Lyric
					_, ok := res[textelementdata_]
					if ok {
						lyrics = res[textelementdata_]
					} else {
						lyrics = make([]*Lyric, 0)
					}
					lyrics = append(lyrics, lyric)
					res[textelementdata_] = lyrics
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extend":
			res := make(map[*Extend][]*Lyric)
			for lyric := range stage.Lyrics {
				if lyric.Extend != nil {
					extend_ := lyric.Extend
					var lyrics []*Lyric
					_, ok := res[extend_]
					if ok {
						lyrics = res[extend_]
					} else {
						lyrics = make([]*Lyric, 0)
					}
					lyrics = append(lyrics, lyric)
					res[extend_] = lyrics
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Mordent
	case Mordent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NonArpeggiate
	case NonArpeggiate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Notations
	case Notations:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		case "Grace":
			res := make(map[*Grace][]*Note)
			for note := range stage.Notes {
				if note.Grace != nil {
					grace_ := note.Grace
					var notes []*Note
					_, ok := res[grace_]
					if ok {
						notes = res[grace_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[grace_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pitch":
			res := make(map[*Pitch][]*Note)
			for note := range stage.Notes {
				if note.Pitch != nil {
					pitch_ := note.Pitch
					var notes []*Note
					_, ok := res[pitch_]
					if ok {
						notes = res[pitch_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[pitch_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Unpitched":
			res := make(map[*Unpitched][]*Note)
			for note := range stage.Notes {
				if note.Unpitched != nil {
					unpitched_ := note.Unpitched
					var notes []*Note
					_, ok := res[unpitched_]
					if ok {
						notes = res[unpitched_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[unpitched_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Rest":
			res := make(map[*Rest][]*Note)
			for note := range stage.Notes {
				if note.Rest != nil {
					rest_ := note.Rest
					var notes []*Note
					_, ok := res[rest_]
					if ok {
						notes = res[rest_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[rest_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Type":
			res := make(map[*NoteType][]*Note)
			for note := range stage.Notes {
				if note.Type != nil {
					notetype_ := note.Type
					var notes []*Note
					_, ok := res[notetype_]
					if ok {
						notes = res[notetype_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[notetype_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Accidental":
			res := make(map[*Accidental][]*Note)
			for note := range stage.Notes {
				if note.Accidental != nil {
					accidental_ := note.Accidental
					var notes []*Note
					_, ok := res[accidental_]
					if ok {
						notes = res[accidental_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[accidental_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Stem":
			res := make(map[*Stem][]*Note)
			for note := range stage.Notes {
				if note.Stem != nil {
					stem_ := note.Stem
					var notes []*Note
					_, ok := res[stem_]
					if ok {
						notes = res[stem_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[stem_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Notehead":
			res := make(map[*Notehead][]*Note)
			for note := range stage.Notes {
				if note.Notehead != nil {
					notehead_ := note.Notehead
					var notes []*Note
					_, ok := res[notehead_]
					if ok {
						notes = res[notehead_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[notehead_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteheadText":
			res := make(map[*NoteheadText][]*Note)
			for note := range stage.Notes {
				if note.NoteheadText != nil {
					noteheadtext_ := note.NoteheadText
					var notes []*Note
					_, ok := res[noteheadtext_]
					if ok {
						notes = res[noteheadtext_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[noteheadtext_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Listen":
			res := make(map[*Listen][]*Note)
			for note := range stage.Notes {
				if note.Listen != nil {
					listen_ := note.Listen
					var notes []*Note
					_, ok := res[listen_]
					if ok {
						notes = res[listen_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[listen_] = notes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteType
	case NoteType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Notehead
	case Notehead:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteheadText
	case NoteheadText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Ornaments
	case Ornaments:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherNotation
	case OtherNotation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherPlacementText
	case OtherPlacementText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Pitch
	case Pitch:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PlacementText
	case PlacementText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Release
	case Release:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Rest
	case Rest:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slide
	case Slide:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slur
	case Slur:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Stem
	case Stem:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StrongAccent
	case StrongAccent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tap
	case Tap:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Technical
	case Technical:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TextElementData
	case TextElementData:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tie
	case Tie:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tied
	case Tied:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tremolo
	case Tremolo:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tuplet
	case Tuplet:
		switch fieldname {
		// insertion point for per direct association field
		case "TupletActual":
			res := make(map[*TupletPortion][]*Tuplet)
			for tuplet := range stage.Tuplets {
				if tuplet.TupletActual != nil {
					tupletportion_ := tuplet.TupletActual
					var tuplets []*Tuplet
					_, ok := res[tupletportion_]
					if ok {
						tuplets = res[tupletportion_]
					} else {
						tuplets = make([]*Tuplet, 0)
					}
					tuplets = append(tuplets, tuplet)
					res[tupletportion_] = tuplets
				}
			}
			return any(res).(map[*End][]*Start)
		case "TupletNormal":
			res := make(map[*TupletPortion][]*Tuplet)
			for tuplet := range stage.Tuplets {
				if tuplet.TupletNormal != nil {
					tupletportion_ := tuplet.TupletNormal
					var tuplets []*Tuplet
					_, ok := res[tupletportion_]
					if ok {
						tuplets = res[tupletportion_]
					} else {
						tuplets = make([]*Tuplet, 0)
					}
					tuplets = append(tuplets, tuplet)
					res[tupletportion_] = tuplets
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TupletDot
	case TupletDot:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TupletNumber
	case TupletNumber:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TupletPortion
	case TupletPortion:
		switch fieldname {
		// insertion point for per direct association field
		case "TupletNumber":
			res := make(map[*TupletNumber][]*TupletPortion)
			for tupletportion := range stage.TupletPortions {
				if tupletportion.TupletNumber != nil {
					tupletnumber_ := tupletportion.TupletNumber
					var tupletportions []*TupletPortion
					_, ok := res[tupletnumber_]
					if ok {
						tupletportions = res[tupletnumber_]
					} else {
						tupletportions = make([]*TupletPortion, 0)
					}
					tupletportions = append(tupletportions, tupletportion)
					res[tupletnumber_] = tupletportions
				}
			}
			return any(res).(map[*End][]*Start)
		case "TupletType":
			res := make(map[*TupletType][]*TupletPortion)
			for tupletportion := range stage.TupletPortions {
				if tupletportion.TupletType != nil {
					tuplettype_ := tupletportion.TupletType
					var tupletportions []*TupletPortion
					_, ok := res[tuplettype_]
					if ok {
						tupletportions = res[tuplettype_]
					} else {
						tupletportions = make([]*TupletPortion, 0)
					}
					tupletportions = append(tupletportions, tupletportion)
					res[tuplettype_] = tupletportions
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TupletType
	case TupletType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Unpitched
	case Unpitched:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Wait
	case Wait:
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
	// reverse maps of direct associations of Accidental
	case Accidental:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AccidentalMark
	case AccidentalMark:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Arpeggiate
	case Arpeggiate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Arrow
	case Arrow:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Articulations
	case Articulations:
		switch fieldname {
		// insertion point for per direct association field
		case "StrongAccent":
			res := make(map[*StrongAccent][]*Articulations)
			for articulations := range stage.Articulationss {
				for _, strongaccent_ := range articulations.StrongAccent {
					res[strongaccent_] = append(res[strongaccent_], articulations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Scoop":
			res := make(map[*EmptyLine][]*Articulations)
			for articulations := range stage.Articulationss {
				for _, emptyline_ := range articulations.Scoop {
					res[emptyline_] = append(res[emptyline_], articulations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Plop":
			res := make(map[*EmptyLine][]*Articulations)
			for articulations := range stage.Articulationss {
				for _, emptyline_ := range articulations.Plop {
					res[emptyline_] = append(res[emptyline_], articulations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Doit":
			res := make(map[*EmptyLine][]*Articulations)
			for articulations := range stage.Articulationss {
				for _, emptyline_ := range articulations.Doit {
					res[emptyline_] = append(res[emptyline_], articulations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Falloff":
			res := make(map[*EmptyLine][]*Articulations)
			for articulations := range stage.Articulationss {
				for _, emptyline_ := range articulations.Falloff {
					res[emptyline_] = append(res[emptyline_], articulations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "BreathMark":
			res := make(map[*BreathMark][]*Articulations)
			for articulations := range stage.Articulationss {
				for _, breathmark_ := range articulations.BreathMark {
					res[breathmark_] = append(res[breathmark_], articulations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Caesura":
			res := make(map[*Caesura][]*Articulations)
			for articulations := range stage.Articulationss {
				for _, caesura_ := range articulations.Caesura {
					res[caesura_] = append(res[caesura_], articulations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "OtherArticulation":
			res := make(map[*OtherPlacementText][]*Articulations)
			for articulations := range stage.Articulationss {
				for _, otherplacementtext_ := range articulations.OtherArticulation {
					res[otherplacementtext_] = append(res[otherplacementtext_], articulations)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Assess
	case Assess:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Backup
	case Backup:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Beam
	case Beam:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bend
	case Bend:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BreathMark
	case BreathMark:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Caesura
	case Caesura:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Elision
	case Elision:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyLine
	case EmptyLine:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Extend
	case Extend:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Figure
	case Figure:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FiguredBass
	case FiguredBass:
		switch fieldname {
		// insertion point for per direct association field
		case "Figure":
			res := make(map[*Figure][]*FiguredBass)
			for figuredbass := range stage.FiguredBasss {
				for _, figure_ := range figuredbass.Figure {
					res[figure_] = append(res[figure_], figuredbass)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Forward
	case Forward:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Glissando
	case Glissando:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Grace
	case Grace:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HammerOnPullOff
	case HammerOnPullOff:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Handbell
	case Handbell:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HarmonClosed
	case HarmonClosed:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HarmonMute
	case HarmonMute:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Harmonic
	case Harmonic:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HeelToe
	case HeelToe:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Hole
	case Hole:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HoleClosed
	case HoleClosed:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Instrument
	case Instrument:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Listen
	case Listen:
		switch fieldname {
		// insertion point for per direct association field
		case "Assess":
			res := make(map[*Assess][]*Listen)
			for listen := range stage.Listens {
				for _, assess_ := range listen.Assess {
					res[assess_] = append(res[assess_], listen)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Wait":
			res := make(map[*Wait][]*Listen)
			for listen := range stage.Listens {
				for _, wait_ := range listen.Wait {
					res[wait_] = append(res[wait_], listen)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Lyric
	case Lyric:
		switch fieldname {
		// insertion point for per direct association field
		case "Elision":
			res := make(map[*Elision][]*Lyric)
			for lyric := range stage.Lyrics {
				for _, elision_ := range lyric.Elision {
					res[elision_] = append(res[elision_], lyric)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Mordent
	case Mordent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NonArpeggiate
	case NonArpeggiate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Notations
	case Notations:
		switch fieldname {
		// insertion point for per direct association field
		case "Tied":
			res := make(map[*Tied][]*Notations)
			for notations := range stage.Notationss {
				for _, tied_ := range notations.Tied {
					res[tied_] = append(res[tied_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Slur":
			res := make(map[*Slur][]*Notations)
			for notations := range stage.Notationss {
				for _, slur_ := range notations.Slur {
					res[slur_] = append(res[slur_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tuplet":
			res := make(map[*Tuplet][]*Notations)
			for notations := range stage.Notationss {
				for _, tuplet_ := range notations.Tuplet {
					res[tuplet_] = append(res[tuplet_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Glissando":
			res := make(map[*Glissando][]*Notations)
			for notations := range stage.Notationss {
				for _, glissando_ := range notations.Glissando {
					res[glissando_] = append(res[glissando_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Slide":
			res := make(map[*Slide][]*Notations)
			for notations := range stage.Notationss {
				for _, slide_ := range notations.Slide {
					res[slide_] = append(res[slide_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Ornaments":
			res := make(map[*Ornaments][]*Notations)
			for notations := range stage.Notationss {
				for _, ornaments_ := range notations.Ornaments {
					res[ornaments_] = append(res[ornaments_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Technical":
			res := make(map[*Technical][]*Notations)
			for notations := range stage.Notationss {
				for _, technical_ := range notations.Technical {
					res[technical_] = append(res[technical_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Articulations":
			res := make(map[*Articulations][]*Notations)
			for notations := range stage.Notationss {
				for _, articulations_ := range notations.Articulations {
					res[articulations_] = append(res[articulations_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Arpeggiate":
			res := make(map[*Arpeggiate][]*Notations)
			for notations := range stage.Notationss {
				for _, arpeggiate_ := range notations.Arpeggiate {
					res[arpeggiate_] = append(res[arpeggiate_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NonArpeggiate":
			res := make(map[*NonArpeggiate][]*Notations)
			for notations := range stage.Notationss {
				for _, nonarpeggiate_ := range notations.NonArpeggiate {
					res[nonarpeggiate_] = append(res[nonarpeggiate_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AccidentalMark":
			res := make(map[*AccidentalMark][]*Notations)
			for notations := range stage.Notationss {
				for _, accidentalmark_ := range notations.AccidentalMark {
					res[accidentalmark_] = append(res[accidentalmark_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		case "OtherNotation":
			res := make(map[*OtherNotation][]*Notations)
			for notations := range stage.Notationss {
				for _, othernotation_ := range notations.OtherNotation {
					res[othernotation_] = append(res[othernotation_], notations)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		case "Tie":
			res := make(map[*Tie][]*Note)
			for note := range stage.Notes {
				for _, tie_ := range note.Tie {
					res[tie_] = append(res[tie_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Instrument":
			res := make(map[*Instrument][]*Note)
			for note := range stage.Notes {
				for _, instrument_ := range note.Instrument {
					res[instrument_] = append(res[instrument_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Beam":
			res := make(map[*Beam][]*Note)
			for note := range stage.Notes {
				for _, beam_ := range note.Beam {
					res[beam_] = append(res[beam_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Notations":
			res := make(map[*Notations][]*Note)
			for note := range stage.Notes {
				for _, notations_ := range note.Notations {
					res[notations_] = append(res[notations_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Lyric":
			res := make(map[*Lyric][]*Note)
			for note := range stage.Notes {
				for _, lyric_ := range note.Lyric {
					res[lyric_] = append(res[lyric_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteType
	case NoteType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Notehead
	case Notehead:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteheadText
	case NoteheadText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Ornaments
	case Ornaments:
		switch fieldname {
		// insertion point for per direct association field
		case "Mordent":
			res := make(map[*Mordent][]*Ornaments)
			for ornaments := range stage.Ornamentss {
				for _, mordent_ := range ornaments.Mordent {
					res[mordent_] = append(res[mordent_], ornaments)
				}
			}
			return any(res).(map[*End][]*Start)
		case "InvertedMordent":
			res := make(map[*Mordent][]*Ornaments)
			for ornaments := range stage.Ornamentss {
				for _, mordent_ := range ornaments.InvertedMordent {
					res[mordent_] = append(res[mordent_], ornaments)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tremolo":
			res := make(map[*Tremolo][]*Ornaments)
			for ornaments := range stage.Ornamentss {
				for _, tremolo_ := range ornaments.Tremolo {
					res[tremolo_] = append(res[tremolo_], ornaments)
				}
			}
			return any(res).(map[*End][]*Start)
		case "OtherOrnament":
			res := make(map[*OtherPlacementText][]*Ornaments)
			for ornaments := range stage.Ornamentss {
				for _, otherplacementtext_ := range ornaments.OtherOrnament {
					res[otherplacementtext_] = append(res[otherplacementtext_], ornaments)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AccidentalMark":
			res := make(map[*AccidentalMark][]*Ornaments)
			for ornaments := range stage.Ornamentss {
				for _, accidentalmark_ := range ornaments.AccidentalMark {
					res[accidentalmark_] = append(res[accidentalmark_], ornaments)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of OtherNotation
	case OtherNotation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherPlacementText
	case OtherPlacementText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Pitch
	case Pitch:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PlacementText
	case PlacementText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Release
	case Release:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Rest
	case Rest:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slide
	case Slide:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slur
	case Slur:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Stem
	case Stem:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StrongAccent
	case StrongAccent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tap
	case Tap:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Technical
	case Technical:
		switch fieldname {
		// insertion point for per direct association field
		case "Harmonic":
			res := make(map[*Harmonic][]*Technical)
			for technical := range stage.Technicals {
				for _, harmonic_ := range technical.Harmonic {
					res[harmonic_] = append(res[harmonic_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pluck":
			res := make(map[*PlacementText][]*Technical)
			for technical := range stage.Technicals {
				for _, placementtext_ := range technical.Pluck {
					res[placementtext_] = append(res[placementtext_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "HammerOn":
			res := make(map[*HammerOnPullOff][]*Technical)
			for technical := range stage.Technicals {
				for _, hammeronpulloff_ := range technical.HammerOn {
					res[hammeronpulloff_] = append(res[hammeronpulloff_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PullOff":
			res := make(map[*HammerOnPullOff][]*Technical)
			for technical := range stage.Technicals {
				for _, hammeronpulloff_ := range technical.PullOff {
					res[hammeronpulloff_] = append(res[hammeronpulloff_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Bend":
			res := make(map[*Bend][]*Technical)
			for technical := range stage.Technicals {
				for _, bend_ := range technical.Bend {
					res[bend_] = append(res[bend_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tap":
			res := make(map[*Tap][]*Technical)
			for technical := range stage.Technicals {
				for _, tap_ := range technical.Tap {
					res[tap_] = append(res[tap_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Heel":
			res := make(map[*HeelToe][]*Technical)
			for technical := range stage.Technicals {
				for _, heeltoe_ := range technical.Heel {
					res[heeltoe_] = append(res[heeltoe_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Toe":
			res := make(map[*HeelToe][]*Technical)
			for technical := range stage.Technicals {
				for _, heeltoe_ := range technical.Toe {
					res[heeltoe_] = append(res[heeltoe_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Hole":
			res := make(map[*Hole][]*Technical)
			for technical := range stage.Technicals {
				for _, hole_ := range technical.Hole {
					res[hole_] = append(res[hole_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Arrow":
			res := make(map[*Arrow][]*Technical)
			for technical := range stage.Technicals {
				for _, arrow_ := range technical.Arrow {
					res[arrow_] = append(res[arrow_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Handbell":
			res := make(map[*Handbell][]*Technical)
			for technical := range stage.Technicals {
				for _, handbell_ := range technical.Handbell {
					res[handbell_] = append(res[handbell_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "HarmonMute":
			res := make(map[*HarmonMute][]*Technical)
			for technical := range stage.Technicals {
				for _, harmonmute_ := range technical.HarmonMute {
					res[harmonmute_] = append(res[harmonmute_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		case "OtherTechnical":
			res := make(map[*OtherPlacementText][]*Technical)
			for technical := range stage.Technicals {
				for _, otherplacementtext_ := range technical.OtherTechnical {
					res[otherplacementtext_] = append(res[otherplacementtext_], technical)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TextElementData
	case TextElementData:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tie
	case Tie:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tied
	case Tied:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tremolo
	case Tremolo:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tuplet
	case Tuplet:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TupletDot
	case TupletDot:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TupletNumber
	case TupletNumber:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TupletPortion
	case TupletPortion:
		switch fieldname {
		// insertion point for per direct association field
		case "TupletDot":
			res := make(map[*TupletDot][]*TupletPortion)
			for tupletportion := range stage.TupletPortions {
				for _, tupletdot_ := range tupletportion.TupletDot {
					res[tupletdot_] = append(res[tupletdot_], tupletportion)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TupletType
	case TupletType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Unpitched
	case Unpitched:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Wait
	case Wait:
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
func (*Accidental) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AccidentalMark) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "AccidentalMark",
		},
		{
			GongstructName: "Ornaments",
			Fieldname: "AccidentalMark",
		},
	}
}

func (*Arpeggiate) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "Arpeggiate",
		},
	}
}

func (*Arrow) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "Arrow",
		},
	}
}

func (*Articulations) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "Articulations",
		},
	}
}

func (*Assess) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Listen",
			Fieldname: "Assess",
		},
	}
}

func (*Backup) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Beam) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Note",
			Fieldname: "Beam",
		},
	}
}

func (*Bend) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "Bend",
		},
	}
}

func (*BreathMark) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Articulations",
			Fieldname: "BreathMark",
		},
	}
}

func (*Caesura) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Articulations",
			Fieldname: "Caesura",
		},
	}
}

func (*Elision) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Lyric",
			Fieldname: "Elision",
		},
	}
}

func (*EmptyLine) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Articulations",
			Fieldname: "Scoop",
		},
		{
			GongstructName: "Articulations",
			Fieldname: "Plop",
		},
		{
			GongstructName: "Articulations",
			Fieldname: "Doit",
		},
		{
			GongstructName: "Articulations",
			Fieldname: "Falloff",
		},
	}
}

func (*Extend) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Figure) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "FiguredBass",
			Fieldname: "Figure",
		},
	}
}

func (*FiguredBass) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Forward) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Glissando) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "Glissando",
		},
	}
}

func (*Grace) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*HammerOnPullOff) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "HammerOn",
		},
		{
			GongstructName: "Technical",
			Fieldname: "PullOff",
		},
	}
}

func (*Handbell) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "Handbell",
		},
	}
}

func (*HarmonClosed) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*HarmonMute) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "HarmonMute",
		},
	}
}

func (*Harmonic) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "Harmonic",
		},
	}
}

func (*HeelToe) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "Heel",
		},
		{
			GongstructName: "Technical",
			Fieldname: "Toe",
		},
	}
}

func (*Hole) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "Hole",
		},
	}
}

func (*HoleClosed) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Instrument) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Note",
			Fieldname: "Instrument",
		},
	}
}

func (*Listen) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Lyric) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Note",
			Fieldname: "Lyric",
		},
	}
}

func (*Mordent) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Ornaments",
			Fieldname: "Mordent",
		},
		{
			GongstructName: "Ornaments",
			Fieldname: "InvertedMordent",
		},
	}
}

func (*NonArpeggiate) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "NonArpeggiate",
		},
	}
}

func (*Notations) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Note",
			Fieldname: "Notations",
		},
	}
}

func (*Note) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*NoteType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Notehead) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*NoteheadText) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Ornaments) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "Ornaments",
		},
	}
}

func (*OtherNotation) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "OtherNotation",
		},
	}
}

func (*OtherPlacementText) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Articulations",
			Fieldname: "OtherArticulation",
		},
		{
			GongstructName: "Ornaments",
			Fieldname: "OtherOrnament",
		},
		{
			GongstructName: "Technical",
			Fieldname: "OtherTechnical",
		},
	}
}

func (*Pitch) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PlacementText) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "Pluck",
		},
	}
}

func (*Release) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Rest) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Slide) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "Slide",
		},
	}
}

func (*Slur) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "Slur",
		},
	}
}

func (*Stem) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*StrongAccent) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Articulations",
			Fieldname: "StrongAccent",
		},
	}
}

func (*Tap) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Technical",
			Fieldname: "Tap",
		},
	}
}

func (*Technical) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "Technical",
		},
	}
}

func (*TextElementData) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Tie) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Note",
			Fieldname: "Tie",
		},
	}
}

func (*Tied) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "Tied",
		},
	}
}

func (*Tremolo) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Ornaments",
			Fieldname: "Tremolo",
		},
	}
}

func (*Tuplet) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Notations",
			Fieldname: "Tuplet",
		},
	}
}

func (*TupletDot) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "TupletPortion",
			Fieldname: "TupletDot",
		},
	}
}

func (*TupletNumber) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TupletPortion) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TupletType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Unpitched) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Wait) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Listen",
			Fieldname: "Wait",
		},
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
func (accidental *Accidental) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (accidentalmark *AccidentalMark) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (arpeggiate *Arpeggiate) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (arrow *Arrow) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ArrowDirection",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumArrowDirection",
		},
		{
			Name:                 "ArrowStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumArrowStyle",
		},
		{
			Name:                 "CircularArrow",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumCircularArrow",
		},
	}
	return
}

func (articulations *Articulations) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "StrongAccent",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StrongAccent",
		},
		{
			Name:                 "Scoop",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EmptyLine",
		},
		{
			Name:                 "Plop",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EmptyLine",
		},
		{
			Name:                 "Doit",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EmptyLine",
		},
		{
			Name:                 "Falloff",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EmptyLine",
		},
		{
			Name:                 "BreathMark",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BreathMark",
		},
		{
			Name:                 "Caesura",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Caesura",
		},
		{
			Name:                 "OtherArticulation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OtherPlacementText",
		},
	}
	return
}

func (assess *Assess) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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
	}
	return
}

func (backup *Backup) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (beam *Beam) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Fan",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFan",
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

func (bend *Bend) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Shape",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumBendShape",
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
			Name:               "Beats",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FirstBeat",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LastBeat",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BendAlter",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Release",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Release",
		},
		{
			Name:                 "WithBar",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PlacementText",
		},
	}
	return
}

func (breathmark *BreathMark) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumBreathMarkValue",
		},
	}
	return
}

func (caesura *Caesura) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumCaesuraValue",
		},
	}
	return
}

func (elision *Elision) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (emptyline *EmptyLine) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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

func (extend *Extend) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (figure *Figure) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Extend",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Extend",
		},
	}
	return
}

func (figuredbass *FiguredBass) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Figure",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Figure",
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (forward *Forward) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Voice",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Staff",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (glissando *Glissando) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (grace *Grace) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StealTimePrevious",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StealTimeFollowing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MakeTime",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (hammeronpulloff *HammerOnPullOff) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (handbell *Handbell) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumHandbellValue",
		},
	}
	return
}

func (harmonclosed *HarmonClosed) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Location",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumHarmonClosedLocation",
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumHarmonClosedValue",
		},
	}
	return
}

func (harmonmute *HarmonMute) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "HarmonClosed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "HarmonClosed",
		},
	}
	return
}

func (harmonic *Harmonic) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (heeltoe *HeelToe) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (hole *Hole) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "HoleType",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "HoleClosed",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "HoleClosed",
		},
		{
			Name:               "HoleShape",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (holeclosed *HoleClosed) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Location",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumHoleClosedLocation",
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumHoleClosedValue",
		},
	}
	return
}

func (instrument *Instrument) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (listen *Listen) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Assess",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Assess",
		},
		{
			Name:                 "Wait",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Wait",
		},
	}
	return
}

func (lyric *Lyric) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "NameAttr",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeOnly",
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
			Name:                 "Syllabic",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumSyllabic",
		},
		{
			Name:                 "Text",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TextElementData",
		},
		{
			Name:                 "Elision",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Elision",
		},
		{
			Name:                 "Extend",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Extend",
		},
	}
	return
}

func (mordent *Mordent) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Beats",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SecondBeat",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LastBeat",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (nonarpeggiate *NonArpeggiate) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (notations *Notations) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Tied",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Tied",
		},
		{
			Name:                 "Slur",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Slur",
		},
		{
			Name:                 "Tuplet",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Tuplet",
		},
		{
			Name:                 "Glissando",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Glissando",
		},
		{
			Name:                 "Slide",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Slide",
		},
		{
			Name:                 "Ornaments",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Ornaments",
		},
		{
			Name:                 "Technical",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Technical",
		},
		{
			Name:                 "Articulations",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Articulations",
		},
		{
			Name:                 "Arpeggiate",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Arpeggiate",
		},
		{
			Name:                 "NonArpeggiate",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NonArpeggiate",
		},
		{
			Name:                 "AccidentalMark",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AccidentalMark",
		},
		{
			Name:                 "OtherNotation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OtherNotation",
		},
	}
	return
}

func (note *Note) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Dynamics",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EndDynamics",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Attack",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Release",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeOnly",
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
			Name:                 "Grace",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Grace",
		},
		{
			Name:                 "Pitch",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Pitch",
		},
		{
			Name:                 "Unpitched",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Unpitched",
		},
		{
			Name:                 "Rest",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rest",
		},
		{
			Name:                 "Tie",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Tie",
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Instrument",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Instrument",
		},
		{
			Name:               "Voice",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "NoteType",
		},
		{
			Name:                 "Accidental",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Accidental",
		},
		{
			Name:                 "Stem",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Stem",
		},
		{
			Name:                 "Notehead",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Notehead",
		},
		{
			Name:                 "NoteheadText",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "NoteheadText",
		},
		{
			Name:               "Staff",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Beam",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Beam",
		},
		{
			Name:                 "Notations",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Notations",
		},
		{
			Name:                 "Lyric",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Lyric",
		},
		{
			Name:                 "Listen",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Listen",
		},
	}
	return
}

func (notetype *NoteType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (notehead *Notehead) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumNoteheadValue",
		},
	}
	return
}

func (noteheadtext *NoteheadText) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (ornaments *Ornaments) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Mordent",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Mordent",
		},
		{
			Name:                 "InvertedMordent",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Mordent",
		},
		{
			Name:                 "Tremolo",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Tremolo",
		},
		{
			Name:                 "OtherOrnament",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OtherPlacementText",
		},
		{
			Name:                 "AccidentalMark",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AccidentalMark",
		},
	}
	return
}

func (othernotation *OtherNotation) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (otherplacementtext *OtherPlacementText) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (pitch *Pitch) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Alter",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Octave",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (placementtext *PlacementText) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (release *Release) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Offset",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (rest *Rest) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DisplayOctave",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (slide *Slide) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Beats",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FirstBeat",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LastBeat",
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

func (slur *Slur) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "BezierX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierX2",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierY2",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierOffset",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierOffset2",
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

func (stem *Stem) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStemValue",
		},
	}
	return
}

func (strongaccent *StrongAccent) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (tap *Tap) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Hand",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTapHand",
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

func (technical *Technical) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Harmonic",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Harmonic",
		},
		{
			Name:                 "Pluck",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PlacementText",
		},
		{
			Name:                 "HammerOn",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "HammerOnPullOff",
		},
		{
			Name:                 "PullOff",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "HammerOnPullOff",
		},
		{
			Name:                 "Bend",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Bend",
		},
		{
			Name:                 "Tap",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Tap",
		},
		{
			Name:                 "Heel",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "HeelToe",
		},
		{
			Name:                 "Toe",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "HeelToe",
		},
		{
			Name:                 "Hole",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Hole",
		},
		{
			Name:                 "Arrow",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Arrow",
		},
		{
			Name:                 "Handbell",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Handbell",
		},
		{
			Name:                 "HarmonMute",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "HarmonMute",
		},
		{
			Name:                 "OtherTechnical",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OtherPlacementText",
		},
	}
	return
}

func (textelementdata *TextElementData) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Lang",
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
			Name:               "Underline",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Overline",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "LineThrough",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Rotation",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LetterSpacing",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (tie *Tie) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeOnly",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (tied *Tied) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "BezierX",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierY",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierX2",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierY2",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierOffset",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BezierOffset2",
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

func (tremolo *Tremolo) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (tuplet *Tuplet) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:                 "TupletActual",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TupletPortion",
		},
		{
			Name:                 "TupletNormal",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TupletPortion",
		},
	}
	return
}

func (tupletdot *TupletDot) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (tupletnumber *TupletNumber) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (tupletportion *TupletPortion) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TupletNumber",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TupletNumber",
		},
		{
			Name:                 "TupletType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TupletType",
		},
		{
			Name:                 "TupletDot",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TupletDot",
		},
	}
	return
}

func (tuplettype *TupletType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (unpitched *Unpitched) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DisplayOctave",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (wait *Wait) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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
func (accidental *Accidental) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = accidental.Name
	case "Smufl":
		res.valueString = accidental.Smufl
	case "DefaultX":
		res.valueString = accidental.DefaultX
	case "DefaultY":
		res.valueString = accidental.DefaultY
	case "RelativeX":
		res.valueString = accidental.RelativeX
	case "RelativeY":
		res.valueString = accidental.RelativeY
	case "FontFamily":
		res.valueString = accidental.FontFamily
	case "FontSize":
		res.valueString = accidental.FontSize
	case "Color":
		res.valueString = accidental.Color
	}
	return
}

func (accidentalmark *AccidentalMark) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = accidentalmark.Name
	case "Smufl":
		res.valueString = accidentalmark.Smufl
	case "DefaultX":
		res.valueString = accidentalmark.DefaultX
	case "DefaultY":
		res.valueString = accidentalmark.DefaultY
	case "RelativeX":
		res.valueString = accidentalmark.RelativeX
	case "RelativeY":
		res.valueString = accidentalmark.RelativeY
	case "FontFamily":
		res.valueString = accidentalmark.FontFamily
	case "FontSize":
		res.valueString = accidentalmark.FontSize
	case "Color":
		res.valueString = accidentalmark.Color
	case "Id":
		res.valueString = accidentalmark.Id
	}
	return
}

func (arpeggiate *Arpeggiate) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arpeggiate.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", arpeggiate.Number)
		res.valueInt = arpeggiate.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = arpeggiate.DefaultX
	case "DefaultY":
		res.valueString = arpeggiate.DefaultY
	case "RelativeX":
		res.valueString = arpeggiate.RelativeX
	case "RelativeY":
		res.valueString = arpeggiate.RelativeY
	case "Color":
		res.valueString = arpeggiate.Color
	case "Id":
		res.valueString = arpeggiate.Id
	}
	return
}

func (arrow *Arrow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arrow.Name
	case "DefaultX":
		res.valueString = arrow.DefaultX
	case "DefaultY":
		res.valueString = arrow.DefaultY
	case "RelativeX":
		res.valueString = arrow.RelativeX
	case "RelativeY":
		res.valueString = arrow.RelativeY
	case "FontFamily":
		res.valueString = arrow.FontFamily
	case "FontSize":
		res.valueString = arrow.FontSize
	case "Color":
		res.valueString = arrow.Color
	case "Smufl":
		res.valueString = arrow.Smufl
	case "ArrowDirection":
		enum := arrow.ArrowDirection
		res.valueString = enum.ToCodeString()
	case "ArrowStyle":
		enum := arrow.ArrowStyle
		res.valueString = enum.ToCodeString()
	case "CircularArrow":
		enum := arrow.CircularArrow
		res.valueString = enum.ToCodeString()
	}
	return
}

func (articulations *Articulations) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = articulations.Name
	case "Id":
		res.valueString = articulations.Id
	case "StrongAccent":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range articulations.StrongAccent {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Scoop":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range articulations.Scoop {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Plop":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range articulations.Plop {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Doit":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range articulations.Doit {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Falloff":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range articulations.Falloff {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "BreathMark":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range articulations.BreathMark {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Caesura":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range articulations.Caesura {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "OtherArticulation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range articulations.OtherArticulation {
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

func (assess *Assess) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = assess.Name
	case "Player":
		res.valueString = assess.Player
	case "TimeOnly":
		res.valueString = assess.TimeOnly
	}
	return
}

func (backup *Backup) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = backup.Name
	case "Duration":
		res.valueString = backup.Duration
	}
	return
}

func (beam *Beam) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = beam.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", beam.Number)
		res.valueInt = beam.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Fan":
		enum := beam.Fan
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = beam.Color
	case "Id":
		res.valueString = beam.Id
	}
	return
}

func (bend *Bend) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bend.Name
	case "Shape":
		enum := bend.Shape
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = bend.DefaultX
	case "DefaultY":
		res.valueString = bend.DefaultY
	case "RelativeX":
		res.valueString = bend.RelativeX
	case "RelativeY":
		res.valueString = bend.RelativeY
	case "FontFamily":
		res.valueString = bend.FontFamily
	case "FontSize":
		res.valueString = bend.FontSize
	case "Color":
		res.valueString = bend.Color
	case "Beats":
		res.valueString = bend.Beats
	case "FirstBeat":
		res.valueString = bend.FirstBeat
	case "LastBeat":
		res.valueString = bend.LastBeat
	case "BendAlter":
		res.valueString = bend.BendAlter
	case "Release":
		res.GongFieldValueType = GongFieldValueTypePointer
		if bend.Release != nil {
			res.valueString = bend.Release.Name
			res.ids = bend.Release.GongGetUUID(stage)
		}
	case "WithBar":
		res.GongFieldValueType = GongFieldValueTypePointer
		if bend.WithBar != nil {
			res.valueString = bend.WithBar.Name
			res.ids = bend.WithBar.GongGetUUID(stage)
		}
	}
	return
}

func (breathmark *BreathMark) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = breathmark.Name
	case "DefaultX":
		res.valueString = breathmark.DefaultX
	case "DefaultY":
		res.valueString = breathmark.DefaultY
	case "RelativeX":
		res.valueString = breathmark.RelativeX
	case "RelativeY":
		res.valueString = breathmark.RelativeY
	case "FontFamily":
		res.valueString = breathmark.FontFamily
	case "FontSize":
		res.valueString = breathmark.FontSize
	case "Color":
		res.valueString = breathmark.Color
	case "Value":
		enum := breathmark.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (caesura *Caesura) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = caesura.Name
	case "DefaultX":
		res.valueString = caesura.DefaultX
	case "DefaultY":
		res.valueString = caesura.DefaultY
	case "RelativeX":
		res.valueString = caesura.RelativeX
	case "RelativeY":
		res.valueString = caesura.RelativeY
	case "FontFamily":
		res.valueString = caesura.FontFamily
	case "FontSize":
		res.valueString = caesura.FontSize
	case "Color":
		res.valueString = caesura.Color
	case "Value":
		enum := caesura.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (elision *Elision) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = elision.Name
	case "Smufl":
		res.valueString = elision.Smufl
	case "FontFamily":
		res.valueString = elision.FontFamily
	case "FontSize":
		res.valueString = elision.FontSize
	case "Color":
		res.valueString = elision.Color
	case "Value":
		res.valueString = elision.Value
	}
	return
}

func (emptyline *EmptyLine) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = emptyline.Name
	case "DashLength":
		res.valueString = emptyline.DashLength
	case "SpaceLength":
		res.valueString = emptyline.SpaceLength
	case "DefaultX":
		res.valueString = emptyline.DefaultX
	case "DefaultY":
		res.valueString = emptyline.DefaultY
	case "RelativeX":
		res.valueString = emptyline.RelativeX
	case "RelativeY":
		res.valueString = emptyline.RelativeY
	case "FontFamily":
		res.valueString = emptyline.FontFamily
	case "FontSize":
		res.valueString = emptyline.FontSize
	case "Color":
		res.valueString = emptyline.Color
	}
	return
}

func (extend *Extend) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = extend.Name
	case "DefaultX":
		res.valueString = extend.DefaultX
	case "DefaultY":
		res.valueString = extend.DefaultY
	case "RelativeX":
		res.valueString = extend.RelativeX
	case "RelativeY":
		res.valueString = extend.RelativeY
	case "Color":
		res.valueString = extend.Color
	}
	return
}

func (figure *Figure) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = figure.Name
	case "Extend":
		res.GongFieldValueType = GongFieldValueTypePointer
		if figure.Extend != nil {
			res.valueString = figure.Extend.Name
			res.ids = figure.Extend.GongGetUUID(stage)
		}
	}
	return
}

func (figuredbass *FiguredBass) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = figuredbass.Name
	case "DefaultX":
		res.valueString = figuredbass.DefaultX
	case "DefaultY":
		res.valueString = figuredbass.DefaultY
	case "RelativeX":
		res.valueString = figuredbass.RelativeX
	case "RelativeY":
		res.valueString = figuredbass.RelativeY
	case "FontFamily":
		res.valueString = figuredbass.FontFamily
	case "FontSize":
		res.valueString = figuredbass.FontSize
	case "Color":
		res.valueString = figuredbass.Color
	case "Id":
		res.valueString = figuredbass.Id
	case "Figure":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range figuredbass.Figure {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Duration":
		res.valueString = figuredbass.Duration
	}
	return
}

func (forward *Forward) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = forward.Name
	case "Duration":
		res.valueString = forward.Duration
	case "Voice":
		res.valueString = forward.Voice
	case "Staff":
		res.valueString = fmt.Sprintf("%d", forward.Staff)
		res.valueInt = forward.Staff
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (glissando *Glissando) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = glissando.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", glissando.Number)
		res.valueInt = glissando.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DashLength":
		res.valueString = glissando.DashLength
	case "SpaceLength":
		res.valueString = glissando.SpaceLength
	case "DefaultX":
		res.valueString = glissando.DefaultX
	case "DefaultY":
		res.valueString = glissando.DefaultY
	case "RelativeX":
		res.valueString = glissando.RelativeX
	case "RelativeY":
		res.valueString = glissando.RelativeY
	case "FontFamily":
		res.valueString = glissando.FontFamily
	case "FontSize":
		res.valueString = glissando.FontSize
	case "Color":
		res.valueString = glissando.Color
	case "Id":
		res.valueString = glissando.Id
	case "Value":
		res.valueString = glissando.Value
	}
	return
}

func (grace *Grace) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = grace.Name
	case "StealTimePrevious":
		res.valueString = grace.StealTimePrevious
	case "StealTimeFollowing":
		res.valueString = grace.StealTimeFollowing
	case "MakeTime":
		res.valueString = grace.MakeTime
	}
	return
}

func (hammeronpulloff *HammerOnPullOff) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = hammeronpulloff.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", hammeronpulloff.Number)
		res.valueInt = hammeronpulloff.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = hammeronpulloff.DefaultX
	case "DefaultY":
		res.valueString = hammeronpulloff.DefaultY
	case "RelativeX":
		res.valueString = hammeronpulloff.RelativeX
	case "RelativeY":
		res.valueString = hammeronpulloff.RelativeY
	case "FontFamily":
		res.valueString = hammeronpulloff.FontFamily
	case "FontSize":
		res.valueString = hammeronpulloff.FontSize
	case "Color":
		res.valueString = hammeronpulloff.Color
	case "Value":
		res.valueString = hammeronpulloff.Value
	}
	return
}

func (handbell *Handbell) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = handbell.Name
	case "DefaultX":
		res.valueString = handbell.DefaultX
	case "DefaultY":
		res.valueString = handbell.DefaultY
	case "RelativeX":
		res.valueString = handbell.RelativeX
	case "RelativeY":
		res.valueString = handbell.RelativeY
	case "FontFamily":
		res.valueString = handbell.FontFamily
	case "FontSize":
		res.valueString = handbell.FontSize
	case "Color":
		res.valueString = handbell.Color
	case "Value":
		enum := handbell.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (harmonclosed *HarmonClosed) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = harmonclosed.Name
	case "Location":
		enum := harmonclosed.Location
		res.valueString = enum.ToCodeString()
	case "Value":
		enum := harmonclosed.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (harmonmute *HarmonMute) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = harmonmute.Name
	case "DefaultX":
		res.valueString = harmonmute.DefaultX
	case "DefaultY":
		res.valueString = harmonmute.DefaultY
	case "RelativeX":
		res.valueString = harmonmute.RelativeX
	case "RelativeY":
		res.valueString = harmonmute.RelativeY
	case "FontFamily":
		res.valueString = harmonmute.FontFamily
	case "FontSize":
		res.valueString = harmonmute.FontSize
	case "Color":
		res.valueString = harmonmute.Color
	case "HarmonClosed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if harmonmute.HarmonClosed != nil {
			res.valueString = harmonmute.HarmonClosed.Name
			res.ids = harmonmute.HarmonClosed.GongGetUUID(stage)
		}
	}
	return
}

func (harmonic *Harmonic) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = harmonic.Name
	case "DefaultX":
		res.valueString = harmonic.DefaultX
	case "DefaultY":
		res.valueString = harmonic.DefaultY
	case "RelativeX":
		res.valueString = harmonic.RelativeX
	case "RelativeY":
		res.valueString = harmonic.RelativeY
	case "FontFamily":
		res.valueString = harmonic.FontFamily
	case "FontSize":
		res.valueString = harmonic.FontSize
	case "Color":
		res.valueString = harmonic.Color
	}
	return
}

func (heeltoe *HeelToe) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = heeltoe.Name
	case "DefaultX":
		res.valueString = heeltoe.DefaultX
	case "DefaultY":
		res.valueString = heeltoe.DefaultY
	case "RelativeX":
		res.valueString = heeltoe.RelativeX
	case "RelativeY":
		res.valueString = heeltoe.RelativeY
	case "FontFamily":
		res.valueString = heeltoe.FontFamily
	case "FontSize":
		res.valueString = heeltoe.FontSize
	case "Color":
		res.valueString = heeltoe.Color
	}
	return
}

func (hole *Hole) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = hole.Name
	case "DefaultX":
		res.valueString = hole.DefaultX
	case "DefaultY":
		res.valueString = hole.DefaultY
	case "RelativeX":
		res.valueString = hole.RelativeX
	case "RelativeY":
		res.valueString = hole.RelativeY
	case "FontFamily":
		res.valueString = hole.FontFamily
	case "FontSize":
		res.valueString = hole.FontSize
	case "Color":
		res.valueString = hole.Color
	case "HoleType":
		res.valueString = hole.HoleType
	case "HoleClosed":
		res.GongFieldValueType = GongFieldValueTypePointer
		if hole.HoleClosed != nil {
			res.valueString = hole.HoleClosed.Name
			res.ids = hole.HoleClosed.GongGetUUID(stage)
		}
	case "HoleShape":
		res.valueString = hole.HoleShape
	}
	return
}

func (holeclosed *HoleClosed) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = holeclosed.Name
	case "Location":
		enum := holeclosed.Location
		res.valueString = enum.ToCodeString()
	case "Value":
		enum := holeclosed.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (instrument *Instrument) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = instrument.Name
	case "Id":
		res.valueString = instrument.Id
	}
	return
}

func (listen *Listen) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = listen.Name
	case "Assess":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range listen.Assess {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Wait":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range listen.Wait {
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

func (lyric *Lyric) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lyric.Name
	case "Number":
		res.valueString = lyric.Number
	case "NameAttr":
		res.valueString = lyric.NameAttr
	case "TimeOnly":
		res.valueString = lyric.TimeOnly
	case "DefaultX":
		res.valueString = lyric.DefaultX
	case "DefaultY":
		res.valueString = lyric.DefaultY
	case "RelativeX":
		res.valueString = lyric.RelativeX
	case "RelativeY":
		res.valueString = lyric.RelativeY
	case "Color":
		res.valueString = lyric.Color
	case "Id":
		res.valueString = lyric.Id
	case "Syllabic":
		enum := lyric.Syllabic
		res.valueString = enum.ToCodeString()
	case "Text":
		res.GongFieldValueType = GongFieldValueTypePointer
		if lyric.Text != nil {
			res.valueString = lyric.Text.Name
			res.ids = lyric.Text.GongGetUUID(stage)
		}
	case "Elision":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lyric.Elision {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Extend":
		res.GongFieldValueType = GongFieldValueTypePointer
		if lyric.Extend != nil {
			res.valueString = lyric.Extend.Name
			res.ids = lyric.Extend.GongGetUUID(stage)
		}
	}
	return
}

func (mordent *Mordent) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = mordent.Name
	case "DefaultX":
		res.valueString = mordent.DefaultX
	case "DefaultY":
		res.valueString = mordent.DefaultY
	case "RelativeX":
		res.valueString = mordent.RelativeX
	case "RelativeY":
		res.valueString = mordent.RelativeY
	case "FontFamily":
		res.valueString = mordent.FontFamily
	case "FontSize":
		res.valueString = mordent.FontSize
	case "Color":
		res.valueString = mordent.Color
	case "Beats":
		res.valueString = mordent.Beats
	case "SecondBeat":
		res.valueString = mordent.SecondBeat
	case "LastBeat":
		res.valueString = mordent.LastBeat
	}
	return
}

func (nonarpeggiate *NonArpeggiate) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = nonarpeggiate.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", nonarpeggiate.Number)
		res.valueInt = nonarpeggiate.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = nonarpeggiate.DefaultX
	case "DefaultY":
		res.valueString = nonarpeggiate.DefaultY
	case "RelativeX":
		res.valueString = nonarpeggiate.RelativeX
	case "RelativeY":
		res.valueString = nonarpeggiate.RelativeY
	case "Color":
		res.valueString = nonarpeggiate.Color
	case "Id":
		res.valueString = nonarpeggiate.Id
	}
	return
}

func (notations *Notations) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notations.Name
	case "Id":
		res.valueString = notations.Id
	case "Tied":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.Tied {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Slur":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.Slur {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Tuplet":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.Tuplet {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Glissando":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.Glissando {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Slide":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.Slide {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Ornaments":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.Ornaments {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Technical":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.Technical {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Articulations":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.Articulations {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Arpeggiate":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.Arpeggiate {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NonArpeggiate":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.NonArpeggiate {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AccidentalMark":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.AccidentalMark {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "OtherNotation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notations.OtherNotation {
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

func (note *Note) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = note.Name
	case "Dynamics":
		res.valueString = note.Dynamics
	case "EndDynamics":
		res.valueString = note.EndDynamics
	case "Attack":
		res.valueString = note.Attack
	case "Release":
		res.valueString = note.Release
	case "TimeOnly":
		res.valueString = note.TimeOnly
	case "DefaultX":
		res.valueString = note.DefaultX
	case "DefaultY":
		res.valueString = note.DefaultY
	case "RelativeX":
		res.valueString = note.RelativeX
	case "RelativeY":
		res.valueString = note.RelativeY
	case "FontFamily":
		res.valueString = note.FontFamily
	case "FontSize":
		res.valueString = note.FontSize
	case "Color":
		res.valueString = note.Color
	case "Id":
		res.valueString = note.Id
	case "Grace":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.Grace != nil {
			res.valueString = note.Grace.Name
			res.ids = note.Grace.GongGetUUID(stage)
		}
	case "Pitch":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.Pitch != nil {
			res.valueString = note.Pitch.Name
			res.ids = note.Pitch.GongGetUUID(stage)
		}
	case "Unpitched":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.Unpitched != nil {
			res.valueString = note.Unpitched.Name
			res.ids = note.Unpitched.GongGetUUID(stage)
		}
	case "Rest":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.Rest != nil {
			res.valueString = note.Rest.Name
			res.ids = note.Rest.GongGetUUID(stage)
		}
	case "Tie":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Tie {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Duration":
		res.valueString = note.Duration
	case "Instrument":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Instrument {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Voice":
		res.valueString = note.Voice
	case "Type":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.Type != nil {
			res.valueString = note.Type.Name
			res.ids = note.Type.GongGetUUID(stage)
		}
	case "Accidental":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.Accidental != nil {
			res.valueString = note.Accidental.Name
			res.ids = note.Accidental.GongGetUUID(stage)
		}
	case "Stem":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.Stem != nil {
			res.valueString = note.Stem.Name
			res.ids = note.Stem.GongGetUUID(stage)
		}
	case "Notehead":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.Notehead != nil {
			res.valueString = note.Notehead.Name
			res.ids = note.Notehead.GongGetUUID(stage)
		}
	case "NoteheadText":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.NoteheadText != nil {
			res.valueString = note.NoteheadText.Name
			res.ids = note.NoteheadText.GongGetUUID(stage)
		}
	case "Staff":
		res.valueString = fmt.Sprintf("%d", note.Staff)
		res.valueInt = note.Staff
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Beam":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Beam {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Notations":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Notations {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Lyric":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Lyric {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Listen":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.Listen != nil {
			res.valueString = note.Listen.Name
			res.ids = note.Listen.GongGetUUID(stage)
		}
	}
	return
}

func (notetype *NoteType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notetype.Name
	}
	return
}

func (notehead *Notehead) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notehead.Name
	case "FontFamily":
		res.valueString = notehead.FontFamily
	case "FontSize":
		res.valueString = notehead.FontSize
	case "Color":
		res.valueString = notehead.Color
	case "Smufl":
		res.valueString = notehead.Smufl
	case "Value":
		enum := notehead.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (noteheadtext *NoteheadText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = noteheadtext.Name
	}
	return
}

func (ornaments *Ornaments) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = ornaments.Name
	case "Id":
		res.valueString = ornaments.Id
	case "Mordent":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ornaments.Mordent {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "InvertedMordent":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ornaments.InvertedMordent {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Tremolo":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ornaments.Tremolo {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "OtherOrnament":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ornaments.OtherOrnament {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AccidentalMark":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range ornaments.AccidentalMark {
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

func (othernotation *OtherNotation) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = othernotation.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", othernotation.Number)
		res.valueInt = othernotation.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = othernotation.DefaultX
	case "DefaultY":
		res.valueString = othernotation.DefaultY
	case "RelativeX":
		res.valueString = othernotation.RelativeX
	case "RelativeY":
		res.valueString = othernotation.RelativeY
	case "FontFamily":
		res.valueString = othernotation.FontFamily
	case "FontSize":
		res.valueString = othernotation.FontSize
	case "Color":
		res.valueString = othernotation.Color
	case "Smufl":
		res.valueString = othernotation.Smufl
	case "Id":
		res.valueString = othernotation.Id
	case "Value":
		res.valueString = othernotation.Value
	}
	return
}

func (otherplacementtext *OtherPlacementText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = otherplacementtext.Name
	case "DefaultX":
		res.valueString = otherplacementtext.DefaultX
	case "DefaultY":
		res.valueString = otherplacementtext.DefaultY
	case "RelativeX":
		res.valueString = otherplacementtext.RelativeX
	case "RelativeY":
		res.valueString = otherplacementtext.RelativeY
	case "FontFamily":
		res.valueString = otherplacementtext.FontFamily
	case "FontSize":
		res.valueString = otherplacementtext.FontSize
	case "Color":
		res.valueString = otherplacementtext.Color
	case "Smufl":
		res.valueString = otherplacementtext.Smufl
	case "Value":
		res.valueString = otherplacementtext.Value
	}
	return
}

func (pitch *Pitch) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pitch.Name
	case "Alter":
		res.valueString = pitch.Alter
	case "Octave":
		res.valueString = fmt.Sprintf("%d", pitch.Octave)
		res.valueInt = pitch.Octave
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (placementtext *PlacementText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = placementtext.Name
	case "DefaultX":
		res.valueString = placementtext.DefaultX
	case "DefaultY":
		res.valueString = placementtext.DefaultY
	case "RelativeX":
		res.valueString = placementtext.RelativeX
	case "RelativeY":
		res.valueString = placementtext.RelativeY
	case "FontFamily":
		res.valueString = placementtext.FontFamily
	case "FontSize":
		res.valueString = placementtext.FontSize
	case "Color":
		res.valueString = placementtext.Color
	case "Value":
		res.valueString = placementtext.Value
	}
	return
}

func (release *Release) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = release.Name
	case "Offset":
		res.valueString = release.Offset
	}
	return
}

func (rest *Rest) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rest.Name
	case "DisplayOctave":
		res.valueString = fmt.Sprintf("%d", rest.DisplayOctave)
		res.valueInt = rest.DisplayOctave
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (slide *Slide) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = slide.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", slide.Number)
		res.valueInt = slide.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DashLength":
		res.valueString = slide.DashLength
	case "SpaceLength":
		res.valueString = slide.SpaceLength
	case "DefaultX":
		res.valueString = slide.DefaultX
	case "DefaultY":
		res.valueString = slide.DefaultY
	case "RelativeX":
		res.valueString = slide.RelativeX
	case "RelativeY":
		res.valueString = slide.RelativeY
	case "FontFamily":
		res.valueString = slide.FontFamily
	case "FontSize":
		res.valueString = slide.FontSize
	case "Color":
		res.valueString = slide.Color
	case "Beats":
		res.valueString = slide.Beats
	case "FirstBeat":
		res.valueString = slide.FirstBeat
	case "LastBeat":
		res.valueString = slide.LastBeat
	case "Id":
		res.valueString = slide.Id
	case "Value":
		res.valueString = slide.Value
	}
	return
}

func (slur *Slur) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = slur.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", slur.Number)
		res.valueInt = slur.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DashLength":
		res.valueString = slur.DashLength
	case "SpaceLength":
		res.valueString = slur.SpaceLength
	case "DefaultX":
		res.valueString = slur.DefaultX
	case "DefaultY":
		res.valueString = slur.DefaultY
	case "RelativeX":
		res.valueString = slur.RelativeX
	case "RelativeY":
		res.valueString = slur.RelativeY
	case "BezierX":
		res.valueString = slur.BezierX
	case "BezierY":
		res.valueString = slur.BezierY
	case "BezierX2":
		res.valueString = slur.BezierX2
	case "BezierY2":
		res.valueString = slur.BezierY2
	case "BezierOffset":
		res.valueString = slur.BezierOffset
	case "BezierOffset2":
		res.valueString = slur.BezierOffset2
	case "Color":
		res.valueString = slur.Color
	case "Id":
		res.valueString = slur.Id
	}
	return
}

func (stem *Stem) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stem.Name
	case "DefaultX":
		res.valueString = stem.DefaultX
	case "DefaultY":
		res.valueString = stem.DefaultY
	case "RelativeX":
		res.valueString = stem.RelativeX
	case "RelativeY":
		res.valueString = stem.RelativeY
	case "Color":
		res.valueString = stem.Color
	case "Value":
		enum := stem.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (strongaccent *StrongAccent) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = strongaccent.Name
	case "DefaultX":
		res.valueString = strongaccent.DefaultX
	case "DefaultY":
		res.valueString = strongaccent.DefaultY
	case "RelativeX":
		res.valueString = strongaccent.RelativeX
	case "RelativeY":
		res.valueString = strongaccent.RelativeY
	case "FontFamily":
		res.valueString = strongaccent.FontFamily
	case "FontSize":
		res.valueString = strongaccent.FontSize
	case "Color":
		res.valueString = strongaccent.Color
	}
	return
}

func (tap *Tap) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tap.Name
	case "Hand":
		enum := tap.Hand
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = tap.DefaultX
	case "DefaultY":
		res.valueString = tap.DefaultY
	case "RelativeX":
		res.valueString = tap.RelativeX
	case "RelativeY":
		res.valueString = tap.RelativeY
	case "FontFamily":
		res.valueString = tap.FontFamily
	case "FontSize":
		res.valueString = tap.FontSize
	case "Color":
		res.valueString = tap.Color
	case "Value":
		res.valueString = tap.Value
	}
	return
}

func (technical *Technical) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = technical.Name
	case "Id":
		res.valueString = technical.Id
	case "Harmonic":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.Harmonic {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Pluck":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.Pluck {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HammerOn":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.HammerOn {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PullOff":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.PullOff {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Bend":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.Bend {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Tap":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.Tap {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Heel":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.Heel {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Toe":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.Toe {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Hole":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.Hole {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Arrow":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.Arrow {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Handbell":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.Handbell {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HarmonMute":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.HarmonMute {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "OtherTechnical":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range technical.OtherTechnical {
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

func (textelementdata *TextElementData) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = textelementdata.Name
	case "Lang":
		res.valueString = textelementdata.Lang
	case "FontFamily":
		res.valueString = textelementdata.FontFamily
	case "FontSize":
		res.valueString = textelementdata.FontSize
	case "Color":
		res.valueString = textelementdata.Color
	case "Underline":
		res.valueString = fmt.Sprintf("%d", textelementdata.Underline)
		res.valueInt = textelementdata.Underline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Overline":
		res.valueString = fmt.Sprintf("%d", textelementdata.Overline)
		res.valueInt = textelementdata.Overline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "LineThrough":
		res.valueString = fmt.Sprintf("%d", textelementdata.LineThrough)
		res.valueInt = textelementdata.LineThrough
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Rotation":
		res.valueString = textelementdata.Rotation
	case "LetterSpacing":
		res.valueString = textelementdata.LetterSpacing
	case "Value":
		res.valueString = textelementdata.Value
	}
	return
}

func (tie *Tie) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tie.Name
	case "TimeOnly":
		res.valueString = tie.TimeOnly
	}
	return
}

func (tied *Tied) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tied.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", tied.Number)
		res.valueInt = tied.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DashLength":
		res.valueString = tied.DashLength
	case "SpaceLength":
		res.valueString = tied.SpaceLength
	case "DefaultX":
		res.valueString = tied.DefaultX
	case "DefaultY":
		res.valueString = tied.DefaultY
	case "RelativeX":
		res.valueString = tied.RelativeX
	case "RelativeY":
		res.valueString = tied.RelativeY
	case "BezierX":
		res.valueString = tied.BezierX
	case "BezierY":
		res.valueString = tied.BezierY
	case "BezierX2":
		res.valueString = tied.BezierX2
	case "BezierY2":
		res.valueString = tied.BezierY2
	case "BezierOffset":
		res.valueString = tied.BezierOffset
	case "BezierOffset2":
		res.valueString = tied.BezierOffset2
	case "Color":
		res.valueString = tied.Color
	case "Id":
		res.valueString = tied.Id
	}
	return
}

func (tremolo *Tremolo) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tremolo.Name
	case "DefaultX":
		res.valueString = tremolo.DefaultX
	case "DefaultY":
		res.valueString = tremolo.DefaultY
	case "RelativeX":
		res.valueString = tremolo.RelativeX
	case "RelativeY":
		res.valueString = tremolo.RelativeY
	case "FontFamily":
		res.valueString = tremolo.FontFamily
	case "FontSize":
		res.valueString = tremolo.FontSize
	case "Color":
		res.valueString = tremolo.Color
	case "Smufl":
		res.valueString = tremolo.Smufl
	case "Value":
		res.valueString = fmt.Sprintf("%d", tremolo.Value)
		res.valueInt = tremolo.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (tuplet *Tuplet) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tuplet.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", tuplet.Number)
		res.valueInt = tuplet.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = tuplet.DefaultX
	case "DefaultY":
		res.valueString = tuplet.DefaultY
	case "RelativeX":
		res.valueString = tuplet.RelativeX
	case "RelativeY":
		res.valueString = tuplet.RelativeY
	case "Id":
		res.valueString = tuplet.Id
	case "TupletActual":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tuplet.TupletActual != nil {
			res.valueString = tuplet.TupletActual.Name
			res.ids = tuplet.TupletActual.GongGetUUID(stage)
		}
	case "TupletNormal":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tuplet.TupletNormal != nil {
			res.valueString = tuplet.TupletNormal.Name
			res.ids = tuplet.TupletNormal.GongGetUUID(stage)
		}
	}
	return
}

func (tupletdot *TupletDot) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tupletdot.Name
	case "FontFamily":
		res.valueString = tupletdot.FontFamily
	case "FontSize":
		res.valueString = tupletdot.FontSize
	case "Color":
		res.valueString = tupletdot.Color
	}
	return
}

func (tupletnumber *TupletNumber) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tupletnumber.Name
	case "FontFamily":
		res.valueString = tupletnumber.FontFamily
	case "FontSize":
		res.valueString = tupletnumber.FontSize
	case "Color":
		res.valueString = tupletnumber.Color
	case "Value":
		res.valueString = fmt.Sprintf("%d", tupletnumber.Value)
		res.valueInt = tupletnumber.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (tupletportion *TupletPortion) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tupletportion.Name
	case "TupletNumber":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tupletportion.TupletNumber != nil {
			res.valueString = tupletportion.TupletNumber.Name
			res.ids = tupletportion.TupletNumber.GongGetUUID(stage)
		}
	case "TupletType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tupletportion.TupletType != nil {
			res.valueString = tupletportion.TupletType.Name
			res.ids = tupletportion.TupletType.GongGetUUID(stage)
		}
	case "TupletDot":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range tupletportion.TupletDot {
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

func (tuplettype *TupletType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tuplettype.Name
	case "FontFamily":
		res.valueString = tuplettype.FontFamily
	case "FontSize":
		res.valueString = tuplettype.FontSize
	case "Color":
		res.valueString = tuplettype.Color
	}
	return
}

func (unpitched *Unpitched) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = unpitched.Name
	case "DisplayOctave":
		res.valueString = fmt.Sprintf("%d", unpitched.DisplayOctave)
		res.valueInt = unpitched.DisplayOctave
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (wait *Wait) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = wait.Name
	case "Player":
		res.valueString = wait.Player
	case "TimeOnly":
		res.valueString = wait.TimeOnly
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
func (accidental *Accidental) GongGetGongstructName() string {
	return "Accidental"
}

func (accidentalmark *AccidentalMark) GongGetGongstructName() string {
	return "AccidentalMark"
}

func (arpeggiate *Arpeggiate) GongGetGongstructName() string {
	return "Arpeggiate"
}

func (arrow *Arrow) GongGetGongstructName() string {
	return "Arrow"
}

func (articulations *Articulations) GongGetGongstructName() string {
	return "Articulations"
}

func (assess *Assess) GongGetGongstructName() string {
	return "Assess"
}

func (backup *Backup) GongGetGongstructName() string {
	return "Backup"
}

func (beam *Beam) GongGetGongstructName() string {
	return "Beam"
}

func (bend *Bend) GongGetGongstructName() string {
	return "Bend"
}

func (breathmark *BreathMark) GongGetGongstructName() string {
	return "BreathMark"
}

func (caesura *Caesura) GongGetGongstructName() string {
	return "Caesura"
}

func (elision *Elision) GongGetGongstructName() string {
	return "Elision"
}

func (emptyline *EmptyLine) GongGetGongstructName() string {
	return "EmptyLine"
}

func (extend *Extend) GongGetGongstructName() string {
	return "Extend"
}

func (figure *Figure) GongGetGongstructName() string {
	return "Figure"
}

func (figuredbass *FiguredBass) GongGetGongstructName() string {
	return "FiguredBass"
}

func (forward *Forward) GongGetGongstructName() string {
	return "Forward"
}

func (glissando *Glissando) GongGetGongstructName() string {
	return "Glissando"
}

func (grace *Grace) GongGetGongstructName() string {
	return "Grace"
}

func (hammeronpulloff *HammerOnPullOff) GongGetGongstructName() string {
	return "HammerOnPullOff"
}

func (handbell *Handbell) GongGetGongstructName() string {
	return "Handbell"
}

func (harmonclosed *HarmonClosed) GongGetGongstructName() string {
	return "HarmonClosed"
}

func (harmonmute *HarmonMute) GongGetGongstructName() string {
	return "HarmonMute"
}

func (harmonic *Harmonic) GongGetGongstructName() string {
	return "Harmonic"
}

func (heeltoe *HeelToe) GongGetGongstructName() string {
	return "HeelToe"
}

func (hole *Hole) GongGetGongstructName() string {
	return "Hole"
}

func (holeclosed *HoleClosed) GongGetGongstructName() string {
	return "HoleClosed"
}

func (instrument *Instrument) GongGetGongstructName() string {
	return "Instrument"
}

func (listen *Listen) GongGetGongstructName() string {
	return "Listen"
}

func (lyric *Lyric) GongGetGongstructName() string {
	return "Lyric"
}

func (mordent *Mordent) GongGetGongstructName() string {
	return "Mordent"
}

func (nonarpeggiate *NonArpeggiate) GongGetGongstructName() string {
	return "NonArpeggiate"
}

func (notations *Notations) GongGetGongstructName() string {
	return "Notations"
}

func (note *Note) GongGetGongstructName() string {
	return "Note"
}

func (notetype *NoteType) GongGetGongstructName() string {
	return "NoteType"
}

func (notehead *Notehead) GongGetGongstructName() string {
	return "Notehead"
}

func (noteheadtext *NoteheadText) GongGetGongstructName() string {
	return "NoteheadText"
}

func (ornaments *Ornaments) GongGetGongstructName() string {
	return "Ornaments"
}

func (othernotation *OtherNotation) GongGetGongstructName() string {
	return "OtherNotation"
}

func (otherplacementtext *OtherPlacementText) GongGetGongstructName() string {
	return "OtherPlacementText"
}

func (pitch *Pitch) GongGetGongstructName() string {
	return "Pitch"
}

func (placementtext *PlacementText) GongGetGongstructName() string {
	return "PlacementText"
}

func (release *Release) GongGetGongstructName() string {
	return "Release"
}

func (rest *Rest) GongGetGongstructName() string {
	return "Rest"
}

func (slide *Slide) GongGetGongstructName() string {
	return "Slide"
}

func (slur *Slur) GongGetGongstructName() string {
	return "Slur"
}

func (stem *Stem) GongGetGongstructName() string {
	return "Stem"
}

func (strongaccent *StrongAccent) GongGetGongstructName() string {
	return "StrongAccent"
}

func (tap *Tap) GongGetGongstructName() string {
	return "Tap"
}

func (technical *Technical) GongGetGongstructName() string {
	return "Technical"
}

func (textelementdata *TextElementData) GongGetGongstructName() string {
	return "TextElementData"
}

func (tie *Tie) GongGetGongstructName() string {
	return "Tie"
}

func (tied *Tied) GongGetGongstructName() string {
	return "Tied"
}

func (tremolo *Tremolo) GongGetGongstructName() string {
	return "Tremolo"
}

func (tuplet *Tuplet) GongGetGongstructName() string {
	return "Tuplet"
}

func (tupletdot *TupletDot) GongGetGongstructName() string {
	return "TupletDot"
}

func (tupletnumber *TupletNumber) GongGetGongstructName() string {
	return "TupletNumber"
}

func (tupletportion *TupletPortion) GongGetGongstructName() string {
	return "TupletPortion"
}

func (tuplettype *TupletType) GongGetGongstructName() string {
	return "TupletType"
}

func (unpitched *Unpitched) GongGetGongstructName() string {
	return "Unpitched"
}

func (wait *Wait) GongGetGongstructName() string {
	return "Wait"
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
	__gong__rebuildMapString(stage.Accidentals, &stage.Accidentals_mapString)

	__gong__rebuildMapString(stage.AccidentalMarks, &stage.AccidentalMarks_mapString)

	__gong__rebuildMapString(stage.Arpeggiates, &stage.Arpeggiates_mapString)

	__gong__rebuildMapString(stage.Arrows, &stage.Arrows_mapString)

	__gong__rebuildMapString(stage.Articulationss, &stage.Articulationss_mapString)

	__gong__rebuildMapString(stage.Assesss, &stage.Assesss_mapString)

	__gong__rebuildMapString(stage.Backups, &stage.Backups_mapString)

	__gong__rebuildMapString(stage.Beams, &stage.Beams_mapString)

	__gong__rebuildMapString(stage.Bends, &stage.Bends_mapString)

	__gong__rebuildMapString(stage.BreathMarks, &stage.BreathMarks_mapString)

	__gong__rebuildMapString(stage.Caesuras, &stage.Caesuras_mapString)

	__gong__rebuildMapString(stage.Elisions, &stage.Elisions_mapString)

	__gong__rebuildMapString(stage.EmptyLines, &stage.EmptyLines_mapString)

	__gong__rebuildMapString(stage.Extends, &stage.Extends_mapString)

	__gong__rebuildMapString(stage.Figures, &stage.Figures_mapString)

	__gong__rebuildMapString(stage.FiguredBasss, &stage.FiguredBasss_mapString)

	__gong__rebuildMapString(stage.Forwards, &stage.Forwards_mapString)

	__gong__rebuildMapString(stage.Glissandos, &stage.Glissandos_mapString)

	__gong__rebuildMapString(stage.Graces, &stage.Graces_mapString)

	__gong__rebuildMapString(stage.HammerOnPullOffs, &stage.HammerOnPullOffs_mapString)

	__gong__rebuildMapString(stage.Handbells, &stage.Handbells_mapString)

	__gong__rebuildMapString(stage.HarmonCloseds, &stage.HarmonCloseds_mapString)

	__gong__rebuildMapString(stage.HarmonMutes, &stage.HarmonMutes_mapString)

	__gong__rebuildMapString(stage.Harmonics, &stage.Harmonics_mapString)

	__gong__rebuildMapString(stage.HeelToes, &stage.HeelToes_mapString)

	__gong__rebuildMapString(stage.Holes, &stage.Holes_mapString)

	__gong__rebuildMapString(stage.HoleCloseds, &stage.HoleCloseds_mapString)

	__gong__rebuildMapString(stage.Instruments, &stage.Instruments_mapString)

	__gong__rebuildMapString(stage.Listens, &stage.Listens_mapString)

	__gong__rebuildMapString(stage.Lyrics, &stage.Lyrics_mapString)

	__gong__rebuildMapString(stage.Mordents, &stage.Mordents_mapString)

	__gong__rebuildMapString(stage.NonArpeggiates, &stage.NonArpeggiates_mapString)

	__gong__rebuildMapString(stage.Notationss, &stage.Notationss_mapString)

	__gong__rebuildMapString(stage.Notes, &stage.Notes_mapString)

	__gong__rebuildMapString(stage.NoteTypes, &stage.NoteTypes_mapString)

	__gong__rebuildMapString(stage.Noteheads, &stage.Noteheads_mapString)

	__gong__rebuildMapString(stage.NoteheadTexts, &stage.NoteheadTexts_mapString)

	__gong__rebuildMapString(stage.Ornamentss, &stage.Ornamentss_mapString)

	__gong__rebuildMapString(stage.OtherNotations, &stage.OtherNotations_mapString)

	__gong__rebuildMapString(stage.OtherPlacementTexts, &stage.OtherPlacementTexts_mapString)

	__gong__rebuildMapString(stage.Pitchs, &stage.Pitchs_mapString)

	__gong__rebuildMapString(stage.PlacementTexts, &stage.PlacementTexts_mapString)

	__gong__rebuildMapString(stage.Releases, &stage.Releases_mapString)

	__gong__rebuildMapString(stage.Rests, &stage.Rests_mapString)

	__gong__rebuildMapString(stage.Slides, &stage.Slides_mapString)

	__gong__rebuildMapString(stage.Slurs, &stage.Slurs_mapString)

	__gong__rebuildMapString(stage.Stems, &stage.Stems_mapString)

	__gong__rebuildMapString(stage.StrongAccents, &stage.StrongAccents_mapString)

	__gong__rebuildMapString(stage.Taps, &stage.Taps_mapString)

	__gong__rebuildMapString(stage.Technicals, &stage.Technicals_mapString)

	__gong__rebuildMapString(stage.TextElementDatas, &stage.TextElementDatas_mapString)

	__gong__rebuildMapString(stage.Ties, &stage.Ties_mapString)

	__gong__rebuildMapString(stage.Tieds, &stage.Tieds_mapString)

	__gong__rebuildMapString(stage.Tremolos, &stage.Tremolos_mapString)

	__gong__rebuildMapString(stage.Tuplets, &stage.Tuplets_mapString)

	__gong__rebuildMapString(stage.TupletDots, &stage.TupletDots_mapString)

	__gong__rebuildMapString(stage.TupletNumbers, &stage.TupletNumbers_mapString)

	__gong__rebuildMapString(stage.TupletPortions, &stage.TupletPortions_mapString)

	__gong__rebuildMapString(stage.TupletTypes, &stage.TupletTypes_mapString)

	__gong__rebuildMapString(stage.Unpitcheds, &stage.Unpitcheds_mapString)

	__gong__rebuildMapString(stage.Waits, &stage.Waits_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
