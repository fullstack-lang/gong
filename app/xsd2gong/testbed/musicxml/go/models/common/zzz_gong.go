// generated code - do not edit
package common

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
	AccidentalTexts                map[*AccidentalText]struct{}
	AccidentalTexts_instance       map[*AccidentalText]*AccidentalText
	AccidentalTexts_mapString      map[string]*AccidentalText
	AccidentalTextOrder            uint
	AccidentalText_stagedOrder     map[*AccidentalText]uint
	AccidentalText_orderStaged     map[uint]*AccidentalText
	AccidentalTexts_reference      map[*AccidentalText]*AccidentalText
	AccidentalTexts_referenceOrder map[*AccidentalText]uint

	// insertion point for slice of pointers maps
	OnAfterAccidentalTextCreateCallback GongOnAfterCreateInterface[AccidentalText]
	OnAfterAccidentalTextUpdateCallback GongOnAfterUpdateInterface[AccidentalText]
	OnAfterAccidentalTextDeleteCallback GongOnAfterDeleteInterface[AccidentalText]

	Appearances                map[*Appearance]struct{}
	Appearances_instance       map[*Appearance]*Appearance
	Appearances_mapString      map[string]*Appearance
	AppearanceOrder            uint
	Appearance_stagedOrder     map[*Appearance]uint
	Appearance_orderStaged     map[uint]*Appearance
	Appearances_reference      map[*Appearance]*Appearance
	Appearances_referenceOrder map[*Appearance]uint

	// insertion point for slice of pointers maps
	Appearance_LineWidth_reverseMap map[*LineWidth]*Appearance

	Appearance_NoteSize_reverseMap map[*NoteSize]*Appearance

	Appearance_Distance_reverseMap map[*Distance]*Appearance

	Appearance_Glyph_reverseMap map[*Glyph]*Appearance

	Appearance_OtherAppearance_reverseMap map[*OtherAppearance]*Appearance

	OnAfterAppearanceCreateCallback GongOnAfterCreateInterface[Appearance]
	OnAfterAppearanceUpdateCallback GongOnAfterUpdateInterface[Appearance]
	OnAfterAppearanceDeleteCallback GongOnAfterDeleteInterface[Appearance]

	Bookmarks                map[*Bookmark]struct{}
	Bookmarks_instance       map[*Bookmark]*Bookmark
	Bookmarks_mapString      map[string]*Bookmark
	BookmarkOrder            uint
	Bookmark_stagedOrder     map[*Bookmark]uint
	Bookmark_orderStaged     map[uint]*Bookmark
	Bookmarks_reference      map[*Bookmark]*Bookmark
	Bookmarks_referenceOrder map[*Bookmark]uint

	// insertion point for slice of pointers maps
	OnAfterBookmarkCreateCallback GongOnAfterCreateInterface[Bookmark]
	OnAfterBookmarkUpdateCallback GongOnAfterUpdateInterface[Bookmark]
	OnAfterBookmarkDeleteCallback GongOnAfterDeleteInterface[Bookmark]

	Codas                map[*Coda]struct{}
	Codas_instance       map[*Coda]*Coda
	Codas_mapString      map[string]*Coda
	CodaOrder            uint
	Coda_stagedOrder     map[*Coda]uint
	Coda_orderStaged     map[uint]*Coda
	Codas_reference      map[*Coda]*Coda
	Codas_referenceOrder map[*Coda]uint

	// insertion point for slice of pointers maps
	OnAfterCodaCreateCallback GongOnAfterCreateInterface[Coda]
	OnAfterCodaUpdateCallback GongOnAfterUpdateInterface[Coda]
	OnAfterCodaDeleteCallback GongOnAfterDeleteInterface[Coda]

	Distances                map[*Distance]struct{}
	Distances_instance       map[*Distance]*Distance
	Distances_mapString      map[string]*Distance
	DistanceOrder            uint
	Distance_stagedOrder     map[*Distance]uint
	Distance_orderStaged     map[uint]*Distance
	Distances_reference      map[*Distance]*Distance
	Distances_referenceOrder map[*Distance]uint

	// insertion point for slice of pointers maps
	OnAfterDistanceCreateCallback GongOnAfterCreateInterface[Distance]
	OnAfterDistanceUpdateCallback GongOnAfterUpdateInterface[Distance]
	OnAfterDistanceDeleteCallback GongOnAfterDeleteInterface[Distance]

	Dynamicss                map[*Dynamics]struct{}
	Dynamicss_instance       map[*Dynamics]*Dynamics
	Dynamicss_mapString      map[string]*Dynamics
	DynamicsOrder            uint
	Dynamics_stagedOrder     map[*Dynamics]uint
	Dynamics_orderStaged     map[uint]*Dynamics
	Dynamicss_reference      map[*Dynamics]*Dynamics
	Dynamicss_referenceOrder map[*Dynamics]uint

	// insertion point for slice of pointers maps
	Dynamics_P_reverseMap map[*Empty]*Dynamics

	Dynamics_Pp_reverseMap map[*Empty]*Dynamics

	Dynamics_Ppp_reverseMap map[*Empty]*Dynamics

	Dynamics_Pppp_reverseMap map[*Empty]*Dynamics

	Dynamics_Ppppp_reverseMap map[*Empty]*Dynamics

	Dynamics_Pppppp_reverseMap map[*Empty]*Dynamics

	Dynamics_F_reverseMap map[*Empty]*Dynamics

	Dynamics_Ff_reverseMap map[*Empty]*Dynamics

	Dynamics_Fff_reverseMap map[*Empty]*Dynamics

	Dynamics_Ffff_reverseMap map[*Empty]*Dynamics

	Dynamics_Fffff_reverseMap map[*Empty]*Dynamics

	Dynamics_Ffffff_reverseMap map[*Empty]*Dynamics

	Dynamics_Mp_reverseMap map[*Empty]*Dynamics

	Dynamics_Mf_reverseMap map[*Empty]*Dynamics

	Dynamics_Sf_reverseMap map[*Empty]*Dynamics

	Dynamics_Sfp_reverseMap map[*Empty]*Dynamics

	Dynamics_Sfpp_reverseMap map[*Empty]*Dynamics

	Dynamics_Fp_reverseMap map[*Empty]*Dynamics

	Dynamics_Rf_reverseMap map[*Empty]*Dynamics

	Dynamics_Rfz_reverseMap map[*Empty]*Dynamics

	Dynamics_Sfz_reverseMap map[*Empty]*Dynamics

	Dynamics_Sffz_reverseMap map[*Empty]*Dynamics

	Dynamics_Fz_reverseMap map[*Empty]*Dynamics

	Dynamics_N_reverseMap map[*Empty]*Dynamics

	Dynamics_Pf_reverseMap map[*Empty]*Dynamics

	Dynamics_Sfzp_reverseMap map[*Empty]*Dynamics

	Dynamics_OtherDynamics_reverseMap map[*OtherText]*Dynamics

	OnAfterDynamicsCreateCallback GongOnAfterCreateInterface[Dynamics]
	OnAfterDynamicsUpdateCallback GongOnAfterUpdateInterface[Dynamics]
	OnAfterDynamicsDeleteCallback GongOnAfterDeleteInterface[Dynamics]

	Emptys                map[*Empty]struct{}
	Emptys_instance       map[*Empty]*Empty
	Emptys_mapString      map[string]*Empty
	EmptyOrder            uint
	Empty_stagedOrder     map[*Empty]uint
	Empty_orderStaged     map[uint]*Empty
	Emptys_reference      map[*Empty]*Empty
	Emptys_referenceOrder map[*Empty]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyCreateCallback GongOnAfterCreateInterface[Empty]
	OnAfterEmptyUpdateCallback GongOnAfterUpdateInterface[Empty]
	OnAfterEmptyDeleteCallback GongOnAfterDeleteInterface[Empty]

	EmptyPlacements                map[*EmptyPlacement]struct{}
	EmptyPlacements_instance       map[*EmptyPlacement]*EmptyPlacement
	EmptyPlacements_mapString      map[string]*EmptyPlacement
	EmptyPlacementOrder            uint
	EmptyPlacement_stagedOrder     map[*EmptyPlacement]uint
	EmptyPlacement_orderStaged     map[uint]*EmptyPlacement
	EmptyPlacements_reference      map[*EmptyPlacement]*EmptyPlacement
	EmptyPlacements_referenceOrder map[*EmptyPlacement]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyPlacementCreateCallback GongOnAfterCreateInterface[EmptyPlacement]
	OnAfterEmptyPlacementUpdateCallback GongOnAfterUpdateInterface[EmptyPlacement]
	OnAfterEmptyPlacementDeleteCallback GongOnAfterDeleteInterface[EmptyPlacement]

	EmptyPlacementSmufls                map[*EmptyPlacementSmufl]struct{}
	EmptyPlacementSmufls_instance       map[*EmptyPlacementSmufl]*EmptyPlacementSmufl
	EmptyPlacementSmufls_mapString      map[string]*EmptyPlacementSmufl
	EmptyPlacementSmuflOrder            uint
	EmptyPlacementSmufl_stagedOrder     map[*EmptyPlacementSmufl]uint
	EmptyPlacementSmufl_orderStaged     map[uint]*EmptyPlacementSmufl
	EmptyPlacementSmufls_reference      map[*EmptyPlacementSmufl]*EmptyPlacementSmufl
	EmptyPlacementSmufls_referenceOrder map[*EmptyPlacementSmufl]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyPlacementSmuflCreateCallback GongOnAfterCreateInterface[EmptyPlacementSmufl]
	OnAfterEmptyPlacementSmuflUpdateCallback GongOnAfterUpdateInterface[EmptyPlacementSmufl]
	OnAfterEmptyPlacementSmuflDeleteCallback GongOnAfterDeleteInterface[EmptyPlacementSmufl]

	EmptyPrintObjectStyleAligns                map[*EmptyPrintObjectStyleAlign]struct{}
	EmptyPrintObjectStyleAligns_instance       map[*EmptyPrintObjectStyleAlign]*EmptyPrintObjectStyleAlign
	EmptyPrintObjectStyleAligns_mapString      map[string]*EmptyPrintObjectStyleAlign
	EmptyPrintObjectStyleAlignOrder            uint
	EmptyPrintObjectStyleAlign_stagedOrder     map[*EmptyPrintObjectStyleAlign]uint
	EmptyPrintObjectStyleAlign_orderStaged     map[uint]*EmptyPrintObjectStyleAlign
	EmptyPrintObjectStyleAligns_reference      map[*EmptyPrintObjectStyleAlign]*EmptyPrintObjectStyleAlign
	EmptyPrintObjectStyleAligns_referenceOrder map[*EmptyPrintObjectStyleAlign]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyPrintObjectStyleAlignCreateCallback GongOnAfterCreateInterface[EmptyPrintObjectStyleAlign]
	OnAfterEmptyPrintObjectStyleAlignUpdateCallback GongOnAfterUpdateInterface[EmptyPrintObjectStyleAlign]
	OnAfterEmptyPrintObjectStyleAlignDeleteCallback GongOnAfterDeleteInterface[EmptyPrintObjectStyleAlign]

	EmptyPrintStyles                map[*EmptyPrintStyle]struct{}
	EmptyPrintStyles_instance       map[*EmptyPrintStyle]*EmptyPrintStyle
	EmptyPrintStyles_mapString      map[string]*EmptyPrintStyle
	EmptyPrintStyleOrder            uint
	EmptyPrintStyle_stagedOrder     map[*EmptyPrintStyle]uint
	EmptyPrintStyle_orderStaged     map[uint]*EmptyPrintStyle
	EmptyPrintStyles_reference      map[*EmptyPrintStyle]*EmptyPrintStyle
	EmptyPrintStyles_referenceOrder map[*EmptyPrintStyle]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyPrintStyleCreateCallback GongOnAfterCreateInterface[EmptyPrintStyle]
	OnAfterEmptyPrintStyleUpdateCallback GongOnAfterUpdateInterface[EmptyPrintStyle]
	OnAfterEmptyPrintStyleDeleteCallback GongOnAfterDeleteInterface[EmptyPrintStyle]

	EmptyPrintStyleAligns                map[*EmptyPrintStyleAlign]struct{}
	EmptyPrintStyleAligns_instance       map[*EmptyPrintStyleAlign]*EmptyPrintStyleAlign
	EmptyPrintStyleAligns_mapString      map[string]*EmptyPrintStyleAlign
	EmptyPrintStyleAlignOrder            uint
	EmptyPrintStyleAlign_stagedOrder     map[*EmptyPrintStyleAlign]uint
	EmptyPrintStyleAlign_orderStaged     map[uint]*EmptyPrintStyleAlign
	EmptyPrintStyleAligns_reference      map[*EmptyPrintStyleAlign]*EmptyPrintStyleAlign
	EmptyPrintStyleAligns_referenceOrder map[*EmptyPrintStyleAlign]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyPrintStyleAlignCreateCallback GongOnAfterCreateInterface[EmptyPrintStyleAlign]
	OnAfterEmptyPrintStyleAlignUpdateCallback GongOnAfterUpdateInterface[EmptyPrintStyleAlign]
	OnAfterEmptyPrintStyleAlignDeleteCallback GongOnAfterDeleteInterface[EmptyPrintStyleAlign]

	EmptyPrintStyleAlignIds                map[*EmptyPrintStyleAlignId]struct{}
	EmptyPrintStyleAlignIds_instance       map[*EmptyPrintStyleAlignId]*EmptyPrintStyleAlignId
	EmptyPrintStyleAlignIds_mapString      map[string]*EmptyPrintStyleAlignId
	EmptyPrintStyleAlignIdOrder            uint
	EmptyPrintStyleAlignId_stagedOrder     map[*EmptyPrintStyleAlignId]uint
	EmptyPrintStyleAlignId_orderStaged     map[uint]*EmptyPrintStyleAlignId
	EmptyPrintStyleAlignIds_reference      map[*EmptyPrintStyleAlignId]*EmptyPrintStyleAlignId
	EmptyPrintStyleAlignIds_referenceOrder map[*EmptyPrintStyleAlignId]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyPrintStyleAlignIdCreateCallback GongOnAfterCreateInterface[EmptyPrintStyleAlignId]
	OnAfterEmptyPrintStyleAlignIdUpdateCallback GongOnAfterUpdateInterface[EmptyPrintStyleAlignId]
	OnAfterEmptyPrintStyleAlignIdDeleteCallback GongOnAfterDeleteInterface[EmptyPrintStyleAlignId]

	EmptyTrillSounds                map[*EmptyTrillSound]struct{}
	EmptyTrillSounds_instance       map[*EmptyTrillSound]*EmptyTrillSound
	EmptyTrillSounds_mapString      map[string]*EmptyTrillSound
	EmptyTrillSoundOrder            uint
	EmptyTrillSound_stagedOrder     map[*EmptyTrillSound]uint
	EmptyTrillSound_orderStaged     map[uint]*EmptyTrillSound
	EmptyTrillSounds_reference      map[*EmptyTrillSound]*EmptyTrillSound
	EmptyTrillSounds_referenceOrder map[*EmptyTrillSound]uint

	// insertion point for slice of pointers maps
	OnAfterEmptyTrillSoundCreateCallback GongOnAfterCreateInterface[EmptyTrillSound]
	OnAfterEmptyTrillSoundUpdateCallback GongOnAfterUpdateInterface[EmptyTrillSound]
	OnAfterEmptyTrillSoundDeleteCallback GongOnAfterDeleteInterface[EmptyTrillSound]

	Encodings                map[*Encoding]struct{}
	Encodings_instance       map[*Encoding]*Encoding
	Encodings_mapString      map[string]*Encoding
	EncodingOrder            uint
	Encoding_stagedOrder     map[*Encoding]uint
	Encoding_orderStaged     map[uint]*Encoding
	Encodings_reference      map[*Encoding]*Encoding
	Encodings_referenceOrder map[*Encoding]uint

	// insertion point for slice of pointers maps
	Encoding_Encoder_reverseMap map[*TypedText]*Encoding

	Encoding_Supports_reverseMap map[*Supports]*Encoding

	OnAfterEncodingCreateCallback GongOnAfterCreateInterface[Encoding]
	OnAfterEncodingUpdateCallback GongOnAfterUpdateInterface[Encoding]
	OnAfterEncodingDeleteCallback GongOnAfterDeleteInterface[Encoding]

	Fermatas                map[*Fermata]struct{}
	Fermatas_instance       map[*Fermata]*Fermata
	Fermatas_mapString      map[string]*Fermata
	FermataOrder            uint
	Fermata_stagedOrder     map[*Fermata]uint
	Fermata_orderStaged     map[uint]*Fermata
	Fermatas_reference      map[*Fermata]*Fermata
	Fermatas_referenceOrder map[*Fermata]uint

	// insertion point for slice of pointers maps
	OnAfterFermataCreateCallback GongOnAfterCreateInterface[Fermata]
	OnAfterFermataUpdateCallback GongOnAfterUpdateInterface[Fermata]
	OnAfterFermataDeleteCallback GongOnAfterDeleteInterface[Fermata]

	Fingerings                map[*Fingering]struct{}
	Fingerings_instance       map[*Fingering]*Fingering
	Fingerings_mapString      map[string]*Fingering
	FingeringOrder            uint
	Fingering_stagedOrder     map[*Fingering]uint
	Fingering_orderStaged     map[uint]*Fingering
	Fingerings_reference      map[*Fingering]*Fingering
	Fingerings_referenceOrder map[*Fingering]uint

	// insertion point for slice of pointers maps
	OnAfterFingeringCreateCallback GongOnAfterCreateInterface[Fingering]
	OnAfterFingeringUpdateCallback GongOnAfterUpdateInterface[Fingering]
	OnAfterFingeringDeleteCallback GongOnAfterDeleteInterface[Fingering]

	FormattedSymbols                map[*FormattedSymbol]struct{}
	FormattedSymbols_instance       map[*FormattedSymbol]*FormattedSymbol
	FormattedSymbols_mapString      map[string]*FormattedSymbol
	FormattedSymbolOrder            uint
	FormattedSymbol_stagedOrder     map[*FormattedSymbol]uint
	FormattedSymbol_orderStaged     map[uint]*FormattedSymbol
	FormattedSymbols_reference      map[*FormattedSymbol]*FormattedSymbol
	FormattedSymbols_referenceOrder map[*FormattedSymbol]uint

	// insertion point for slice of pointers maps
	OnAfterFormattedSymbolCreateCallback GongOnAfterCreateInterface[FormattedSymbol]
	OnAfterFormattedSymbolUpdateCallback GongOnAfterUpdateInterface[FormattedSymbol]
	OnAfterFormattedSymbolDeleteCallback GongOnAfterDeleteInterface[FormattedSymbol]

	FormattedSymbolIds                map[*FormattedSymbolId]struct{}
	FormattedSymbolIds_instance       map[*FormattedSymbolId]*FormattedSymbolId
	FormattedSymbolIds_mapString      map[string]*FormattedSymbolId
	FormattedSymbolIdOrder            uint
	FormattedSymbolId_stagedOrder     map[*FormattedSymbolId]uint
	FormattedSymbolId_orderStaged     map[uint]*FormattedSymbolId
	FormattedSymbolIds_reference      map[*FormattedSymbolId]*FormattedSymbolId
	FormattedSymbolIds_referenceOrder map[*FormattedSymbolId]uint

	// insertion point for slice of pointers maps
	OnAfterFormattedSymbolIdCreateCallback GongOnAfterCreateInterface[FormattedSymbolId]
	OnAfterFormattedSymbolIdUpdateCallback GongOnAfterUpdateInterface[FormattedSymbolId]
	OnAfterFormattedSymbolIdDeleteCallback GongOnAfterDeleteInterface[FormattedSymbolId]

	FormattedTexts                map[*FormattedText]struct{}
	FormattedTexts_instance       map[*FormattedText]*FormattedText
	FormattedTexts_mapString      map[string]*FormattedText
	FormattedTextOrder            uint
	FormattedText_stagedOrder     map[*FormattedText]uint
	FormattedText_orderStaged     map[uint]*FormattedText
	FormattedTexts_reference      map[*FormattedText]*FormattedText
	FormattedTexts_referenceOrder map[*FormattedText]uint

	// insertion point for slice of pointers maps
	OnAfterFormattedTextCreateCallback GongOnAfterCreateInterface[FormattedText]
	OnAfterFormattedTextUpdateCallback GongOnAfterUpdateInterface[FormattedText]
	OnAfterFormattedTextDeleteCallback GongOnAfterDeleteInterface[FormattedText]

	FormattedTextIds                map[*FormattedTextId]struct{}
	FormattedTextIds_instance       map[*FormattedTextId]*FormattedTextId
	FormattedTextIds_mapString      map[string]*FormattedTextId
	FormattedTextIdOrder            uint
	FormattedTextId_stagedOrder     map[*FormattedTextId]uint
	FormattedTextId_orderStaged     map[uint]*FormattedTextId
	FormattedTextIds_reference      map[*FormattedTextId]*FormattedTextId
	FormattedTextIds_referenceOrder map[*FormattedTextId]uint

	// insertion point for slice of pointers maps
	OnAfterFormattedTextIdCreateCallback GongOnAfterCreateInterface[FormattedTextId]
	OnAfterFormattedTextIdUpdateCallback GongOnAfterUpdateInterface[FormattedTextId]
	OnAfterFormattedTextIdDeleteCallback GongOnAfterDeleteInterface[FormattedTextId]

	Frets                map[*Fret]struct{}
	Frets_instance       map[*Fret]*Fret
	Frets_mapString      map[string]*Fret
	FretOrder            uint
	Fret_stagedOrder     map[*Fret]uint
	Fret_orderStaged     map[uint]*Fret
	Frets_reference      map[*Fret]*Fret
	Frets_referenceOrder map[*Fret]uint

	// insertion point for slice of pointers maps
	OnAfterFretCreateCallback GongOnAfterCreateInterface[Fret]
	OnAfterFretUpdateCallback GongOnAfterUpdateInterface[Fret]
	OnAfterFretDeleteCallback GongOnAfterDeleteInterface[Fret]

	Glyphs                map[*Glyph]struct{}
	Glyphs_instance       map[*Glyph]*Glyph
	Glyphs_mapString      map[string]*Glyph
	GlyphOrder            uint
	Glyph_stagedOrder     map[*Glyph]uint
	Glyph_orderStaged     map[uint]*Glyph
	Glyphs_reference      map[*Glyph]*Glyph
	Glyphs_referenceOrder map[*Glyph]uint

	// insertion point for slice of pointers maps
	OnAfterGlyphCreateCallback GongOnAfterCreateInterface[Glyph]
	OnAfterGlyphUpdateCallback GongOnAfterUpdateInterface[Glyph]
	OnAfterGlyphDeleteCallback GongOnAfterDeleteInterface[Glyph]

	HorizontalTurns                map[*HorizontalTurn]struct{}
	HorizontalTurns_instance       map[*HorizontalTurn]*HorizontalTurn
	HorizontalTurns_mapString      map[string]*HorizontalTurn
	HorizontalTurnOrder            uint
	HorizontalTurn_stagedOrder     map[*HorizontalTurn]uint
	HorizontalTurn_orderStaged     map[uint]*HorizontalTurn
	HorizontalTurns_reference      map[*HorizontalTurn]*HorizontalTurn
	HorizontalTurns_referenceOrder map[*HorizontalTurn]uint

	// insertion point for slice of pointers maps
	OnAfterHorizontalTurnCreateCallback GongOnAfterCreateInterface[HorizontalTurn]
	OnAfterHorizontalTurnUpdateCallback GongOnAfterUpdateInterface[HorizontalTurn]
	OnAfterHorizontalTurnDeleteCallback GongOnAfterDeleteInterface[HorizontalTurn]

	Identifications                map[*Identification]struct{}
	Identifications_instance       map[*Identification]*Identification
	Identifications_mapString      map[string]*Identification
	IdentificationOrder            uint
	Identification_stagedOrder     map[*Identification]uint
	Identification_orderStaged     map[uint]*Identification
	Identifications_reference      map[*Identification]*Identification
	Identifications_referenceOrder map[*Identification]uint

	// insertion point for slice of pointers maps
	Identification_Creator_reverseMap map[*TypedText]*Identification

	Identification_Rights_reverseMap map[*TypedText]*Identification

	Identification_Relation_reverseMap map[*TypedText]*Identification

	OnAfterIdentificationCreateCallback GongOnAfterCreateInterface[Identification]
	OnAfterIdentificationUpdateCallback GongOnAfterUpdateInterface[Identification]
	OnAfterIdentificationDeleteCallback GongOnAfterDeleteInterface[Identification]

	Levels                map[*Level]struct{}
	Levels_instance       map[*Level]*Level
	Levels_mapString      map[string]*Level
	LevelOrder            uint
	Level_stagedOrder     map[*Level]uint
	Level_orderStaged     map[uint]*Level
	Levels_reference      map[*Level]*Level
	Levels_referenceOrder map[*Level]uint

	// insertion point for slice of pointers maps
	OnAfterLevelCreateCallback GongOnAfterCreateInterface[Level]
	OnAfterLevelUpdateCallback GongOnAfterUpdateInterface[Level]
	OnAfterLevelDeleteCallback GongOnAfterDeleteInterface[Level]

	LineWidths                map[*LineWidth]struct{}
	LineWidths_instance       map[*LineWidth]*LineWidth
	LineWidths_mapString      map[string]*LineWidth
	LineWidthOrder            uint
	LineWidth_stagedOrder     map[*LineWidth]uint
	LineWidth_orderStaged     map[uint]*LineWidth
	LineWidths_reference      map[*LineWidth]*LineWidth
	LineWidths_referenceOrder map[*LineWidth]uint

	// insertion point for slice of pointers maps
	OnAfterLineWidthCreateCallback GongOnAfterCreateInterface[LineWidth]
	OnAfterLineWidthUpdateCallback GongOnAfterUpdateInterface[LineWidth]
	OnAfterLineWidthDeleteCallback GongOnAfterDeleteInterface[LineWidth]

	Links                map[*Link]struct{}
	Links_instance       map[*Link]*Link
	Links_mapString      map[string]*Link
	LinkOrder            uint
	Link_stagedOrder     map[*Link]uint
	Link_orderStaged     map[uint]*Link
	Links_reference      map[*Link]*Link
	Links_referenceOrder map[*Link]uint

	// insertion point for slice of pointers maps
	OnAfterLinkCreateCallback GongOnAfterCreateInterface[Link]
	OnAfterLinkUpdateCallback GongOnAfterUpdateInterface[Link]
	OnAfterLinkDeleteCallback GongOnAfterDeleteInterface[Link]

	MeasureLayouts                map[*MeasureLayout]struct{}
	MeasureLayouts_instance       map[*MeasureLayout]*MeasureLayout
	MeasureLayouts_mapString      map[string]*MeasureLayout
	MeasureLayoutOrder            uint
	MeasureLayout_stagedOrder     map[*MeasureLayout]uint
	MeasureLayout_orderStaged     map[uint]*MeasureLayout
	MeasureLayouts_reference      map[*MeasureLayout]*MeasureLayout
	MeasureLayouts_referenceOrder map[*MeasureLayout]uint

	// insertion point for slice of pointers maps
	OnAfterMeasureLayoutCreateCallback GongOnAfterCreateInterface[MeasureLayout]
	OnAfterMeasureLayoutUpdateCallback GongOnAfterUpdateInterface[MeasureLayout]
	OnAfterMeasureLayoutDeleteCallback GongOnAfterDeleteInterface[MeasureLayout]

	MidiDevices                map[*MidiDevice]struct{}
	MidiDevices_instance       map[*MidiDevice]*MidiDevice
	MidiDevices_mapString      map[string]*MidiDevice
	MidiDeviceOrder            uint
	MidiDevice_stagedOrder     map[*MidiDevice]uint
	MidiDevice_orderStaged     map[uint]*MidiDevice
	MidiDevices_reference      map[*MidiDevice]*MidiDevice
	MidiDevices_referenceOrder map[*MidiDevice]uint

	// insertion point for slice of pointers maps
	OnAfterMidiDeviceCreateCallback GongOnAfterCreateInterface[MidiDevice]
	OnAfterMidiDeviceUpdateCallback GongOnAfterUpdateInterface[MidiDevice]
	OnAfterMidiDeviceDeleteCallback GongOnAfterDeleteInterface[MidiDevice]

	MidiInstruments                map[*MidiInstrument]struct{}
	MidiInstruments_instance       map[*MidiInstrument]*MidiInstrument
	MidiInstruments_mapString      map[string]*MidiInstrument
	MidiInstrumentOrder            uint
	MidiInstrument_stagedOrder     map[*MidiInstrument]uint
	MidiInstrument_orderStaged     map[uint]*MidiInstrument
	MidiInstruments_reference      map[*MidiInstrument]*MidiInstrument
	MidiInstruments_referenceOrder map[*MidiInstrument]uint

	// insertion point for slice of pointers maps
	OnAfterMidiInstrumentCreateCallback GongOnAfterCreateInterface[MidiInstrument]
	OnAfterMidiInstrumentUpdateCallback GongOnAfterUpdateInterface[MidiInstrument]
	OnAfterMidiInstrumentDeleteCallback GongOnAfterDeleteInterface[MidiInstrument]

	Miscellaneouss                map[*Miscellaneous]struct{}
	Miscellaneouss_instance       map[*Miscellaneous]*Miscellaneous
	Miscellaneouss_mapString      map[string]*Miscellaneous
	MiscellaneousOrder            uint
	Miscellaneous_stagedOrder     map[*Miscellaneous]uint
	Miscellaneous_orderStaged     map[uint]*Miscellaneous
	Miscellaneouss_reference      map[*Miscellaneous]*Miscellaneous
	Miscellaneouss_referenceOrder map[*Miscellaneous]uint

	// insertion point for slice of pointers maps
	Miscellaneous_MiscellaneousField_reverseMap map[*MiscellaneousField]*Miscellaneous

	OnAfterMiscellaneousCreateCallback GongOnAfterCreateInterface[Miscellaneous]
	OnAfterMiscellaneousUpdateCallback GongOnAfterUpdateInterface[Miscellaneous]
	OnAfterMiscellaneousDeleteCallback GongOnAfterDeleteInterface[Miscellaneous]

	MiscellaneousFields                map[*MiscellaneousField]struct{}
	MiscellaneousFields_instance       map[*MiscellaneousField]*MiscellaneousField
	MiscellaneousFields_mapString      map[string]*MiscellaneousField
	MiscellaneousFieldOrder            uint
	MiscellaneousField_stagedOrder     map[*MiscellaneousField]uint
	MiscellaneousField_orderStaged     map[uint]*MiscellaneousField
	MiscellaneousFields_reference      map[*MiscellaneousField]*MiscellaneousField
	MiscellaneousFields_referenceOrder map[*MiscellaneousField]uint

	// insertion point for slice of pointers maps
	OnAfterMiscellaneousFieldCreateCallback GongOnAfterCreateInterface[MiscellaneousField]
	OnAfterMiscellaneousFieldUpdateCallback GongOnAfterUpdateInterface[MiscellaneousField]
	OnAfterMiscellaneousFieldDeleteCallback GongOnAfterDeleteInterface[MiscellaneousField]

	NameDisplays                map[*NameDisplay]struct{}
	NameDisplays_instance       map[*NameDisplay]*NameDisplay
	NameDisplays_mapString      map[string]*NameDisplay
	NameDisplayOrder            uint
	NameDisplay_stagedOrder     map[*NameDisplay]uint
	NameDisplay_orderStaged     map[uint]*NameDisplay
	NameDisplays_reference      map[*NameDisplay]*NameDisplay
	NameDisplays_referenceOrder map[*NameDisplay]uint

	// insertion point for slice of pointers maps
	NameDisplay_DisplayText_reverseMap map[*FormattedText]*NameDisplay

	NameDisplay_AccidentalText_reverseMap map[*AccidentalText]*NameDisplay

	OnAfterNameDisplayCreateCallback GongOnAfterCreateInterface[NameDisplay]
	OnAfterNameDisplayUpdateCallback GongOnAfterUpdateInterface[NameDisplay]
	OnAfterNameDisplayDeleteCallback GongOnAfterDeleteInterface[NameDisplay]

	NoteSizes                map[*NoteSize]struct{}
	NoteSizes_instance       map[*NoteSize]*NoteSize
	NoteSizes_mapString      map[string]*NoteSize
	NoteSizeOrder            uint
	NoteSize_stagedOrder     map[*NoteSize]uint
	NoteSize_orderStaged     map[uint]*NoteSize
	NoteSizes_reference      map[*NoteSize]*NoteSize
	NoteSizes_referenceOrder map[*NoteSize]uint

	// insertion point for slice of pointers maps
	OnAfterNoteSizeCreateCallback GongOnAfterCreateInterface[NoteSize]
	OnAfterNoteSizeUpdateCallback GongOnAfterUpdateInterface[NoteSize]
	OnAfterNoteSizeDeleteCallback GongOnAfterDeleteInterface[NoteSize]

	OtherAppearances                map[*OtherAppearance]struct{}
	OtherAppearances_instance       map[*OtherAppearance]*OtherAppearance
	OtherAppearances_mapString      map[string]*OtherAppearance
	OtherAppearanceOrder            uint
	OtherAppearance_stagedOrder     map[*OtherAppearance]uint
	OtherAppearance_orderStaged     map[uint]*OtherAppearance
	OtherAppearances_reference      map[*OtherAppearance]*OtherAppearance
	OtherAppearances_referenceOrder map[*OtherAppearance]uint

	// insertion point for slice of pointers maps
	OnAfterOtherAppearanceCreateCallback GongOnAfterCreateInterface[OtherAppearance]
	OnAfterOtherAppearanceUpdateCallback GongOnAfterUpdateInterface[OtherAppearance]
	OnAfterOtherAppearanceDeleteCallback GongOnAfterDeleteInterface[OtherAppearance]

	OtherPlays                map[*OtherPlay]struct{}
	OtherPlays_instance       map[*OtherPlay]*OtherPlay
	OtherPlays_mapString      map[string]*OtherPlay
	OtherPlayOrder            uint
	OtherPlay_stagedOrder     map[*OtherPlay]uint
	OtherPlay_orderStaged     map[uint]*OtherPlay
	OtherPlays_reference      map[*OtherPlay]*OtherPlay
	OtherPlays_referenceOrder map[*OtherPlay]uint

	// insertion point for slice of pointers maps
	OnAfterOtherPlayCreateCallback GongOnAfterCreateInterface[OtherPlay]
	OnAfterOtherPlayUpdateCallback GongOnAfterUpdateInterface[OtherPlay]
	OnAfterOtherPlayDeleteCallback GongOnAfterDeleteInterface[OtherPlay]

	OtherTexts                map[*OtherText]struct{}
	OtherTexts_instance       map[*OtherText]*OtherText
	OtherTexts_mapString      map[string]*OtherText
	OtherTextOrder            uint
	OtherText_stagedOrder     map[*OtherText]uint
	OtherText_orderStaged     map[uint]*OtherText
	OtherTexts_reference      map[*OtherText]*OtherText
	OtherTexts_referenceOrder map[*OtherText]uint

	// insertion point for slice of pointers maps
	OnAfterOtherTextCreateCallback GongOnAfterCreateInterface[OtherText]
	OnAfterOtherTextUpdateCallback GongOnAfterUpdateInterface[OtherText]
	OnAfterOtherTextDeleteCallback GongOnAfterDeleteInterface[OtherText]

	PageLayouts                map[*PageLayout]struct{}
	PageLayouts_instance       map[*PageLayout]*PageLayout
	PageLayouts_mapString      map[string]*PageLayout
	PageLayoutOrder            uint
	PageLayout_stagedOrder     map[*PageLayout]uint
	PageLayout_orderStaged     map[uint]*PageLayout
	PageLayouts_reference      map[*PageLayout]*PageLayout
	PageLayouts_referenceOrder map[*PageLayout]uint

	// insertion point for slice of pointers maps
	PageLayout_PageMargins_reverseMap map[*PageMargins]*PageLayout

	OnAfterPageLayoutCreateCallback GongOnAfterCreateInterface[PageLayout]
	OnAfterPageLayoutUpdateCallback GongOnAfterUpdateInterface[PageLayout]
	OnAfterPageLayoutDeleteCallback GongOnAfterDeleteInterface[PageLayout]

	PageMarginss                map[*PageMargins]struct{}
	PageMarginss_instance       map[*PageMargins]*PageMargins
	PageMarginss_mapString      map[string]*PageMargins
	PageMarginsOrder            uint
	PageMargins_stagedOrder     map[*PageMargins]uint
	PageMargins_orderStaged     map[uint]*PageMargins
	PageMarginss_reference      map[*PageMargins]*PageMargins
	PageMarginss_referenceOrder map[*PageMargins]uint

	// insertion point for slice of pointers maps
	OnAfterPageMarginsCreateCallback GongOnAfterCreateInterface[PageMargins]
	OnAfterPageMarginsUpdateCallback GongOnAfterUpdateInterface[PageMargins]
	OnAfterPageMarginsDeleteCallback GongOnAfterDeleteInterface[PageMargins]

	Plays                map[*Play]struct{}
	Plays_instance       map[*Play]*Play
	Plays_mapString      map[string]*Play
	PlayOrder            uint
	Play_stagedOrder     map[*Play]uint
	Play_orderStaged     map[uint]*Play
	Plays_reference      map[*Play]*Play
	Plays_referenceOrder map[*Play]uint

	// insertion point for slice of pointers maps
	Play_OtherPlay_reverseMap map[*OtherPlay]*Play

	OnAfterPlayCreateCallback GongOnAfterCreateInterface[Play]
	OnAfterPlayUpdateCallback GongOnAfterUpdateInterface[Play]
	OnAfterPlayDeleteCallback GongOnAfterDeleteInterface[Play]

	Scalings                map[*Scaling]struct{}
	Scalings_instance       map[*Scaling]*Scaling
	Scalings_mapString      map[string]*Scaling
	ScalingOrder            uint
	Scaling_stagedOrder     map[*Scaling]uint
	Scaling_orderStaged     map[uint]*Scaling
	Scalings_reference      map[*Scaling]*Scaling
	Scalings_referenceOrder map[*Scaling]uint

	// insertion point for slice of pointers maps
	OnAfterScalingCreateCallback GongOnAfterCreateInterface[Scaling]
	OnAfterScalingUpdateCallback GongOnAfterUpdateInterface[Scaling]
	OnAfterScalingDeleteCallback GongOnAfterDeleteInterface[Scaling]

	Segnos                map[*Segno]struct{}
	Segnos_instance       map[*Segno]*Segno
	Segnos_mapString      map[string]*Segno
	SegnoOrder            uint
	Segno_stagedOrder     map[*Segno]uint
	Segno_orderStaged     map[uint]*Segno
	Segnos_reference      map[*Segno]*Segno
	Segnos_referenceOrder map[*Segno]uint

	// insertion point for slice of pointers maps
	OnAfterSegnoCreateCallback GongOnAfterCreateInterface[Segno]
	OnAfterSegnoUpdateCallback GongOnAfterUpdateInterface[Segno]
	OnAfterSegnoDeleteCallback GongOnAfterDeleteInterface[Segno]

	StaffLayouts                map[*StaffLayout]struct{}
	StaffLayouts_instance       map[*StaffLayout]*StaffLayout
	StaffLayouts_mapString      map[string]*StaffLayout
	StaffLayoutOrder            uint
	StaffLayout_stagedOrder     map[*StaffLayout]uint
	StaffLayout_orderStaged     map[uint]*StaffLayout
	StaffLayouts_reference      map[*StaffLayout]*StaffLayout
	StaffLayouts_referenceOrder map[*StaffLayout]uint

	// insertion point for slice of pointers maps
	OnAfterStaffLayoutCreateCallback GongOnAfterCreateInterface[StaffLayout]
	OnAfterStaffLayoutUpdateCallback GongOnAfterUpdateInterface[StaffLayout]
	OnAfterStaffLayoutDeleteCallback GongOnAfterDeleteInterface[StaffLayout]

	StringTypes                map[*StringType]struct{}
	StringTypes_instance       map[*StringType]*StringType
	StringTypes_mapString      map[string]*StringType
	StringTypeOrder            uint
	StringType_stagedOrder     map[*StringType]uint
	StringType_orderStaged     map[uint]*StringType
	StringTypes_reference      map[*StringType]*StringType
	StringTypes_referenceOrder map[*StringType]uint

	// insertion point for slice of pointers maps
	OnAfterStringTypeCreateCallback GongOnAfterCreateInterface[StringType]
	OnAfterStringTypeUpdateCallback GongOnAfterUpdateInterface[StringType]
	OnAfterStringTypeDeleteCallback GongOnAfterDeleteInterface[StringType]

	Supportss                map[*Supports]struct{}
	Supportss_instance       map[*Supports]*Supports
	Supportss_mapString      map[string]*Supports
	SupportsOrder            uint
	Supports_stagedOrder     map[*Supports]uint
	Supports_orderStaged     map[uint]*Supports
	Supportss_reference      map[*Supports]*Supports
	Supportss_referenceOrder map[*Supports]uint

	// insertion point for slice of pointers maps
	OnAfterSupportsCreateCallback GongOnAfterCreateInterface[Supports]
	OnAfterSupportsUpdateCallback GongOnAfterUpdateInterface[Supports]
	OnAfterSupportsDeleteCallback GongOnAfterDeleteInterface[Supports]

	SystemDividerss                map[*SystemDividers]struct{}
	SystemDividerss_instance       map[*SystemDividers]*SystemDividers
	SystemDividerss_mapString      map[string]*SystemDividers
	SystemDividersOrder            uint
	SystemDividers_stagedOrder     map[*SystemDividers]uint
	SystemDividers_orderStaged     map[uint]*SystemDividers
	SystemDividerss_reference      map[*SystemDividers]*SystemDividers
	SystemDividerss_referenceOrder map[*SystemDividers]uint

	// insertion point for slice of pointers maps
	OnAfterSystemDividersCreateCallback GongOnAfterCreateInterface[SystemDividers]
	OnAfterSystemDividersUpdateCallback GongOnAfterUpdateInterface[SystemDividers]
	OnAfterSystemDividersDeleteCallback GongOnAfterDeleteInterface[SystemDividers]

	SystemLayouts                map[*SystemLayout]struct{}
	SystemLayouts_instance       map[*SystemLayout]*SystemLayout
	SystemLayouts_mapString      map[string]*SystemLayout
	SystemLayoutOrder            uint
	SystemLayout_stagedOrder     map[*SystemLayout]uint
	SystemLayout_orderStaged     map[uint]*SystemLayout
	SystemLayouts_reference      map[*SystemLayout]*SystemLayout
	SystemLayouts_referenceOrder map[*SystemLayout]uint

	// insertion point for slice of pointers maps
	OnAfterSystemLayoutCreateCallback GongOnAfterCreateInterface[SystemLayout]
	OnAfterSystemLayoutUpdateCallback GongOnAfterUpdateInterface[SystemLayout]
	OnAfterSystemLayoutDeleteCallback GongOnAfterDeleteInterface[SystemLayout]

	SystemMarginss                map[*SystemMargins]struct{}
	SystemMarginss_instance       map[*SystemMargins]*SystemMargins
	SystemMarginss_mapString      map[string]*SystemMargins
	SystemMarginsOrder            uint
	SystemMargins_stagedOrder     map[*SystemMargins]uint
	SystemMargins_orderStaged     map[uint]*SystemMargins
	SystemMarginss_reference      map[*SystemMargins]*SystemMargins
	SystemMarginss_referenceOrder map[*SystemMargins]uint

	// insertion point for slice of pointers maps
	OnAfterSystemMarginsCreateCallback GongOnAfterCreateInterface[SystemMargins]
	OnAfterSystemMarginsUpdateCallback GongOnAfterUpdateInterface[SystemMargins]
	OnAfterSystemMarginsDeleteCallback GongOnAfterDeleteInterface[SystemMargins]

	TypedTexts                map[*TypedText]struct{}
	TypedTexts_instance       map[*TypedText]*TypedText
	TypedTexts_mapString      map[string]*TypedText
	TypedTextOrder            uint
	TypedText_stagedOrder     map[*TypedText]uint
	TypedText_orderStaged     map[uint]*TypedText
	TypedTexts_reference      map[*TypedText]*TypedText
	TypedTexts_referenceOrder map[*TypedText]uint

	// insertion point for slice of pointers maps
	OnAfterTypedTextCreateCallback GongOnAfterCreateInterface[TypedText]
	OnAfterTypedTextUpdateCallback GongOnAfterUpdateInterface[TypedText]
	OnAfterTypedTextDeleteCallback GongOnAfterDeleteInterface[TypedText]

	WavyLines                map[*WavyLine]struct{}
	WavyLines_instance       map[*WavyLine]*WavyLine
	WavyLines_mapString      map[string]*WavyLine
	WavyLineOrder            uint
	WavyLine_stagedOrder     map[*WavyLine]uint
	WavyLine_orderStaged     map[uint]*WavyLine
	WavyLines_reference      map[*WavyLine]*WavyLine
	WavyLines_referenceOrder map[*WavyLine]uint

	// insertion point for slice of pointers maps
	OnAfterWavyLineCreateCallback GongOnAfterCreateInterface[WavyLine]
	OnAfterWavyLineUpdateCallback GongOnAfterUpdateInterface[WavyLine]
	OnAfterWavyLineDeleteCallback GongOnAfterDeleteInterface[WavyLine]

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
	__gong__clearReferences(&stage.AccidentalTexts_reference, &stage.AccidentalTexts_instance, &stage.AccidentalTexts_referenceOrder)

	__gong__clearReferences(&stage.Appearances_reference, &stage.Appearances_instance, &stage.Appearances_referenceOrder)

	__gong__clearReferences(&stage.Bookmarks_reference, &stage.Bookmarks_instance, &stage.Bookmarks_referenceOrder)

	__gong__clearReferences(&stage.Codas_reference, &stage.Codas_instance, &stage.Codas_referenceOrder)

	__gong__clearReferences(&stage.Distances_reference, &stage.Distances_instance, &stage.Distances_referenceOrder)

	__gong__clearReferences(&stage.Dynamicss_reference, &stage.Dynamicss_instance, &stage.Dynamicss_referenceOrder)

	__gong__clearReferences(&stage.Emptys_reference, &stage.Emptys_instance, &stage.Emptys_referenceOrder)

	__gong__clearReferences(&stage.EmptyPlacements_reference, &stage.EmptyPlacements_instance, &stage.EmptyPlacements_referenceOrder)

	__gong__clearReferences(&stage.EmptyPlacementSmufls_reference, &stage.EmptyPlacementSmufls_instance, &stage.EmptyPlacementSmufls_referenceOrder)

	__gong__clearReferences(&stage.EmptyPrintObjectStyleAligns_reference, &stage.EmptyPrintObjectStyleAligns_instance, &stage.EmptyPrintObjectStyleAligns_referenceOrder)

	__gong__clearReferences(&stage.EmptyPrintStyles_reference, &stage.EmptyPrintStyles_instance, &stage.EmptyPrintStyles_referenceOrder)

	__gong__clearReferences(&stage.EmptyPrintStyleAligns_reference, &stage.EmptyPrintStyleAligns_instance, &stage.EmptyPrintStyleAligns_referenceOrder)

	__gong__clearReferences(&stage.EmptyPrintStyleAlignIds_reference, &stage.EmptyPrintStyleAlignIds_instance, &stage.EmptyPrintStyleAlignIds_referenceOrder)

	__gong__clearReferences(&stage.EmptyTrillSounds_reference, &stage.EmptyTrillSounds_instance, &stage.EmptyTrillSounds_referenceOrder)

	__gong__clearReferences(&stage.Encodings_reference, &stage.Encodings_instance, &stage.Encodings_referenceOrder)

	__gong__clearReferences(&stage.Fermatas_reference, &stage.Fermatas_instance, &stage.Fermatas_referenceOrder)

	__gong__clearReferences(&stage.Fingerings_reference, &stage.Fingerings_instance, &stage.Fingerings_referenceOrder)

	__gong__clearReferences(&stage.FormattedSymbols_reference, &stage.FormattedSymbols_instance, &stage.FormattedSymbols_referenceOrder)

	__gong__clearReferences(&stage.FormattedSymbolIds_reference, &stage.FormattedSymbolIds_instance, &stage.FormattedSymbolIds_referenceOrder)

	__gong__clearReferences(&stage.FormattedTexts_reference, &stage.FormattedTexts_instance, &stage.FormattedTexts_referenceOrder)

	__gong__clearReferences(&stage.FormattedTextIds_reference, &stage.FormattedTextIds_instance, &stage.FormattedTextIds_referenceOrder)

	__gong__clearReferences(&stage.Frets_reference, &stage.Frets_instance, &stage.Frets_referenceOrder)

	__gong__clearReferences(&stage.Glyphs_reference, &stage.Glyphs_instance, &stage.Glyphs_referenceOrder)

	__gong__clearReferences(&stage.HorizontalTurns_reference, &stage.HorizontalTurns_instance, &stage.HorizontalTurns_referenceOrder)

	__gong__clearReferences(&stage.Identifications_reference, &stage.Identifications_instance, &stage.Identifications_referenceOrder)

	__gong__clearReferences(&stage.Levels_reference, &stage.Levels_instance, &stage.Levels_referenceOrder)

	__gong__clearReferences(&stage.LineWidths_reference, &stage.LineWidths_instance, &stage.LineWidths_referenceOrder)

	__gong__clearReferences(&stage.Links_reference, &stage.Links_instance, &stage.Links_referenceOrder)

	__gong__clearReferences(&stage.MeasureLayouts_reference, &stage.MeasureLayouts_instance, &stage.MeasureLayouts_referenceOrder)

	__gong__clearReferences(&stage.MidiDevices_reference, &stage.MidiDevices_instance, &stage.MidiDevices_referenceOrder)

	__gong__clearReferences(&stage.MidiInstruments_reference, &stage.MidiInstruments_instance, &stage.MidiInstruments_referenceOrder)

	__gong__clearReferences(&stage.Miscellaneouss_reference, &stage.Miscellaneouss_instance, &stage.Miscellaneouss_referenceOrder)

	__gong__clearReferences(&stage.MiscellaneousFields_reference, &stage.MiscellaneousFields_instance, &stage.MiscellaneousFields_referenceOrder)

	__gong__clearReferences(&stage.NameDisplays_reference, &stage.NameDisplays_instance, &stage.NameDisplays_referenceOrder)

	__gong__clearReferences(&stage.NoteSizes_reference, &stage.NoteSizes_instance, &stage.NoteSizes_referenceOrder)

	__gong__clearReferences(&stage.OtherAppearances_reference, &stage.OtherAppearances_instance, &stage.OtherAppearances_referenceOrder)

	__gong__clearReferences(&stage.OtherPlays_reference, &stage.OtherPlays_instance, &stage.OtherPlays_referenceOrder)

	__gong__clearReferences(&stage.OtherTexts_reference, &stage.OtherTexts_instance, &stage.OtherTexts_referenceOrder)

	__gong__clearReferences(&stage.PageLayouts_reference, &stage.PageLayouts_instance, &stage.PageLayouts_referenceOrder)

	__gong__clearReferences(&stage.PageMarginss_reference, &stage.PageMarginss_instance, &stage.PageMarginss_referenceOrder)

	__gong__clearReferences(&stage.Plays_reference, &stage.Plays_instance, &stage.Plays_referenceOrder)

	__gong__clearReferences(&stage.Scalings_reference, &stage.Scalings_instance, &stage.Scalings_referenceOrder)

	__gong__clearReferences(&stage.Segnos_reference, &stage.Segnos_instance, &stage.Segnos_referenceOrder)

	__gong__clearReferences(&stage.StaffLayouts_reference, &stage.StaffLayouts_instance, &stage.StaffLayouts_referenceOrder)

	__gong__clearReferences(&stage.StringTypes_reference, &stage.StringTypes_instance, &stage.StringTypes_referenceOrder)

	__gong__clearReferences(&stage.Supportss_reference, &stage.Supportss_instance, &stage.Supportss_referenceOrder)

	__gong__clearReferences(&stage.SystemDividerss_reference, &stage.SystemDividerss_instance, &stage.SystemDividerss_referenceOrder)

	__gong__clearReferences(&stage.SystemLayouts_reference, &stage.SystemLayouts_instance, &stage.SystemLayouts_referenceOrder)

	__gong__clearReferences(&stage.SystemMarginss_reference, &stage.SystemMarginss_instance, &stage.SystemMarginss_referenceOrder)

	__gong__clearReferences(&stage.TypedTexts_reference, &stage.TypedTexts_instance, &stage.TypedTexts_referenceOrder)

	__gong__clearReferences(&stage.WavyLines_reference, &stage.WavyLines_instance, &stage.WavyLines_referenceOrder)

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
	stage.AccidentalTextOrder = __gong__recomputeOrder(stage.AccidentalText_stagedOrder)

	stage.AppearanceOrder = __gong__recomputeOrder(stage.Appearance_stagedOrder)

	stage.BookmarkOrder = __gong__recomputeOrder(stage.Bookmark_stagedOrder)

	stage.CodaOrder = __gong__recomputeOrder(stage.Coda_stagedOrder)

	stage.DistanceOrder = __gong__recomputeOrder(stage.Distance_stagedOrder)

	stage.DynamicsOrder = __gong__recomputeOrder(stage.Dynamics_stagedOrder)

	stage.EmptyOrder = __gong__recomputeOrder(stage.Empty_stagedOrder)

	stage.EmptyPlacementOrder = __gong__recomputeOrder(stage.EmptyPlacement_stagedOrder)

	stage.EmptyPlacementSmuflOrder = __gong__recomputeOrder(stage.EmptyPlacementSmufl_stagedOrder)

	stage.EmptyPrintObjectStyleAlignOrder = __gong__recomputeOrder(stage.EmptyPrintObjectStyleAlign_stagedOrder)

	stage.EmptyPrintStyleOrder = __gong__recomputeOrder(stage.EmptyPrintStyle_stagedOrder)

	stage.EmptyPrintStyleAlignOrder = __gong__recomputeOrder(stage.EmptyPrintStyleAlign_stagedOrder)

	stage.EmptyPrintStyleAlignIdOrder = __gong__recomputeOrder(stage.EmptyPrintStyleAlignId_stagedOrder)

	stage.EmptyTrillSoundOrder = __gong__recomputeOrder(stage.EmptyTrillSound_stagedOrder)

	stage.EncodingOrder = __gong__recomputeOrder(stage.Encoding_stagedOrder)

	stage.FermataOrder = __gong__recomputeOrder(stage.Fermata_stagedOrder)

	stage.FingeringOrder = __gong__recomputeOrder(stage.Fingering_stagedOrder)

	stage.FormattedSymbolOrder = __gong__recomputeOrder(stage.FormattedSymbol_stagedOrder)

	stage.FormattedSymbolIdOrder = __gong__recomputeOrder(stage.FormattedSymbolId_stagedOrder)

	stage.FormattedTextOrder = __gong__recomputeOrder(stage.FormattedText_stagedOrder)

	stage.FormattedTextIdOrder = __gong__recomputeOrder(stage.FormattedTextId_stagedOrder)

	stage.FretOrder = __gong__recomputeOrder(stage.Fret_stagedOrder)

	stage.GlyphOrder = __gong__recomputeOrder(stage.Glyph_stagedOrder)

	stage.HorizontalTurnOrder = __gong__recomputeOrder(stage.HorizontalTurn_stagedOrder)

	stage.IdentificationOrder = __gong__recomputeOrder(stage.Identification_stagedOrder)

	stage.LevelOrder = __gong__recomputeOrder(stage.Level_stagedOrder)

	stage.LineWidthOrder = __gong__recomputeOrder(stage.LineWidth_stagedOrder)

	stage.LinkOrder = __gong__recomputeOrder(stage.Link_stagedOrder)

	stage.MeasureLayoutOrder = __gong__recomputeOrder(stage.MeasureLayout_stagedOrder)

	stage.MidiDeviceOrder = __gong__recomputeOrder(stage.MidiDevice_stagedOrder)

	stage.MidiInstrumentOrder = __gong__recomputeOrder(stage.MidiInstrument_stagedOrder)

	stage.MiscellaneousOrder = __gong__recomputeOrder(stage.Miscellaneous_stagedOrder)

	stage.MiscellaneousFieldOrder = __gong__recomputeOrder(stage.MiscellaneousField_stagedOrder)

	stage.NameDisplayOrder = __gong__recomputeOrder(stage.NameDisplay_stagedOrder)

	stage.NoteSizeOrder = __gong__recomputeOrder(stage.NoteSize_stagedOrder)

	stage.OtherAppearanceOrder = __gong__recomputeOrder(stage.OtherAppearance_stagedOrder)

	stage.OtherPlayOrder = __gong__recomputeOrder(stage.OtherPlay_stagedOrder)

	stage.OtherTextOrder = __gong__recomputeOrder(stage.OtherText_stagedOrder)

	stage.PageLayoutOrder = __gong__recomputeOrder(stage.PageLayout_stagedOrder)

	stage.PageMarginsOrder = __gong__recomputeOrder(stage.PageMargins_stagedOrder)

	stage.PlayOrder = __gong__recomputeOrder(stage.Play_stagedOrder)

	stage.ScalingOrder = __gong__recomputeOrder(stage.Scaling_stagedOrder)

	stage.SegnoOrder = __gong__recomputeOrder(stage.Segno_stagedOrder)

	stage.StaffLayoutOrder = __gong__recomputeOrder(stage.StaffLayout_stagedOrder)

	stage.StringTypeOrder = __gong__recomputeOrder(stage.StringType_stagedOrder)

	stage.SupportsOrder = __gong__recomputeOrder(stage.Supports_stagedOrder)

	stage.SystemDividersOrder = __gong__recomputeOrder(stage.SystemDividers_stagedOrder)

	stage.SystemLayoutOrder = __gong__recomputeOrder(stage.SystemLayout_stagedOrder)

	stage.SystemMarginsOrder = __gong__recomputeOrder(stage.SystemMargins_stagedOrder)

	stage.TypedTextOrder = __gong__recomputeOrder(stage.TypedText_stagedOrder)

	stage.WavyLineOrder = __gong__recomputeOrder(stage.WavyLine_stagedOrder)

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
func (*AccidentalText) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AccidentalTexts, stage.AccidentalText_stagedOrder)
}

func (*AccidentalText) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AccidentalText_orderStaged[order]
}

func (*AccidentalText) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AccidentalTexts_mapString
}

func (*AccidentalText) GongGetInstancesSet(stage *Stage) any {
	return &stage.AccidentalTexts
}

func (*AccidentalText) GongNewInstance() any {
	return new(AccidentalText)
}

func (*Appearance) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Appearances, stage.Appearance_stagedOrder)
}

func (*Appearance) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Appearance_orderStaged[order]
}

func (*Appearance) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Appearances_mapString
}

func (*Appearance) GongGetInstancesSet(stage *Stage) any {
	return &stage.Appearances
}

func (*Appearance) GongNewInstance() any {
	return new(Appearance)
}

func (*Bookmark) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Bookmarks, stage.Bookmark_stagedOrder)
}

func (*Bookmark) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Bookmark_orderStaged[order]
}

func (*Bookmark) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Bookmarks_mapString
}

func (*Bookmark) GongGetInstancesSet(stage *Stage) any {
	return &stage.Bookmarks
}

func (*Bookmark) GongNewInstance() any {
	return new(Bookmark)
}

func (*Coda) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Codas, stage.Coda_stagedOrder)
}

func (*Coda) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Coda_orderStaged[order]
}

func (*Coda) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Codas_mapString
}

func (*Coda) GongGetInstancesSet(stage *Stage) any {
	return &stage.Codas
}

func (*Coda) GongNewInstance() any {
	return new(Coda)
}

func (*Distance) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Distances, stage.Distance_stagedOrder)
}

func (*Distance) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Distance_orderStaged[order]
}

func (*Distance) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Distances_mapString
}

func (*Distance) GongGetInstancesSet(stage *Stage) any {
	return &stage.Distances
}

func (*Distance) GongNewInstance() any {
	return new(Distance)
}

func (*Dynamics) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Dynamicss, stage.Dynamics_stagedOrder)
}

func (*Dynamics) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Dynamics_orderStaged[order]
}

func (*Dynamics) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Dynamicss_mapString
}

func (*Dynamics) GongGetInstancesSet(stage *Stage) any {
	return &stage.Dynamicss
}

func (*Dynamics) GongNewInstance() any {
	return new(Dynamics)
}

func (*Empty) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Emptys, stage.Empty_stagedOrder)
}

func (*Empty) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Empty_orderStaged[order]
}

func (*Empty) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Emptys_mapString
}

func (*Empty) GongGetInstancesSet(stage *Stage) any {
	return &stage.Emptys
}

func (*Empty) GongNewInstance() any {
	return new(Empty)
}

func (*EmptyPlacement) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EmptyPlacements, stage.EmptyPlacement_stagedOrder)
}

func (*EmptyPlacement) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EmptyPlacement_orderStaged[order]
}

func (*EmptyPlacement) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EmptyPlacements_mapString
}

func (*EmptyPlacement) GongGetInstancesSet(stage *Stage) any {
	return &stage.EmptyPlacements
}

func (*EmptyPlacement) GongNewInstance() any {
	return new(EmptyPlacement)
}

func (*EmptyPlacementSmufl) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EmptyPlacementSmufls, stage.EmptyPlacementSmufl_stagedOrder)
}

func (*EmptyPlacementSmufl) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EmptyPlacementSmufl_orderStaged[order]
}

func (*EmptyPlacementSmufl) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EmptyPlacementSmufls_mapString
}

func (*EmptyPlacementSmufl) GongGetInstancesSet(stage *Stage) any {
	return &stage.EmptyPlacementSmufls
}

func (*EmptyPlacementSmufl) GongNewInstance() any {
	return new(EmptyPlacementSmufl)
}

func (*EmptyPrintObjectStyleAlign) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EmptyPrintObjectStyleAligns, stage.EmptyPrintObjectStyleAlign_stagedOrder)
}

func (*EmptyPrintObjectStyleAlign) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EmptyPrintObjectStyleAlign_orderStaged[order]
}

func (*EmptyPrintObjectStyleAlign) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EmptyPrintObjectStyleAligns_mapString
}

func (*EmptyPrintObjectStyleAlign) GongGetInstancesSet(stage *Stage) any {
	return &stage.EmptyPrintObjectStyleAligns
}

func (*EmptyPrintObjectStyleAlign) GongNewInstance() any {
	return new(EmptyPrintObjectStyleAlign)
}

func (*EmptyPrintStyle) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EmptyPrintStyles, stage.EmptyPrintStyle_stagedOrder)
}

func (*EmptyPrintStyle) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EmptyPrintStyle_orderStaged[order]
}

func (*EmptyPrintStyle) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EmptyPrintStyles_mapString
}

func (*EmptyPrintStyle) GongGetInstancesSet(stage *Stage) any {
	return &stage.EmptyPrintStyles
}

func (*EmptyPrintStyle) GongNewInstance() any {
	return new(EmptyPrintStyle)
}

func (*EmptyPrintStyleAlign) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EmptyPrintStyleAligns, stage.EmptyPrintStyleAlign_stagedOrder)
}

func (*EmptyPrintStyleAlign) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EmptyPrintStyleAlign_orderStaged[order]
}

func (*EmptyPrintStyleAlign) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EmptyPrintStyleAligns_mapString
}

func (*EmptyPrintStyleAlign) GongGetInstancesSet(stage *Stage) any {
	return &stage.EmptyPrintStyleAligns
}

func (*EmptyPrintStyleAlign) GongNewInstance() any {
	return new(EmptyPrintStyleAlign)
}

func (*EmptyPrintStyleAlignId) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EmptyPrintStyleAlignIds, stage.EmptyPrintStyleAlignId_stagedOrder)
}

func (*EmptyPrintStyleAlignId) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EmptyPrintStyleAlignId_orderStaged[order]
}

func (*EmptyPrintStyleAlignId) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EmptyPrintStyleAlignIds_mapString
}

func (*EmptyPrintStyleAlignId) GongGetInstancesSet(stage *Stage) any {
	return &stage.EmptyPrintStyleAlignIds
}

func (*EmptyPrintStyleAlignId) GongNewInstance() any {
	return new(EmptyPrintStyleAlignId)
}

func (*EmptyTrillSound) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.EmptyTrillSounds, stage.EmptyTrillSound_stagedOrder)
}

func (*EmptyTrillSound) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.EmptyTrillSound_orderStaged[order]
}

func (*EmptyTrillSound) GongGetInstancesMapByName(stage *Stage) any {
	return stage.EmptyTrillSounds_mapString
}

func (*EmptyTrillSound) GongGetInstancesSet(stage *Stage) any {
	return &stage.EmptyTrillSounds
}

func (*EmptyTrillSound) GongNewInstance() any {
	return new(EmptyTrillSound)
}

func (*Encoding) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Encodings, stage.Encoding_stagedOrder)
}

func (*Encoding) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Encoding_orderStaged[order]
}

func (*Encoding) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Encodings_mapString
}

func (*Encoding) GongGetInstancesSet(stage *Stage) any {
	return &stage.Encodings
}

func (*Encoding) GongNewInstance() any {
	return new(Encoding)
}

func (*Fermata) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Fermatas, stage.Fermata_stagedOrder)
}

func (*Fermata) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Fermata_orderStaged[order]
}

func (*Fermata) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Fermatas_mapString
}

func (*Fermata) GongGetInstancesSet(stage *Stage) any {
	return &stage.Fermatas
}

func (*Fermata) GongNewInstance() any {
	return new(Fermata)
}

func (*Fingering) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Fingerings, stage.Fingering_stagedOrder)
}

func (*Fingering) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Fingering_orderStaged[order]
}

func (*Fingering) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Fingerings_mapString
}

func (*Fingering) GongGetInstancesSet(stage *Stage) any {
	return &stage.Fingerings
}

func (*Fingering) GongNewInstance() any {
	return new(Fingering)
}

func (*FormattedSymbol) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.FormattedSymbols, stage.FormattedSymbol_stagedOrder)
}

func (*FormattedSymbol) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.FormattedSymbol_orderStaged[order]
}

func (*FormattedSymbol) GongGetInstancesMapByName(stage *Stage) any {
	return stage.FormattedSymbols_mapString
}

func (*FormattedSymbol) GongGetInstancesSet(stage *Stage) any {
	return &stage.FormattedSymbols
}

func (*FormattedSymbol) GongNewInstance() any {
	return new(FormattedSymbol)
}

func (*FormattedSymbolId) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.FormattedSymbolIds, stage.FormattedSymbolId_stagedOrder)
}

func (*FormattedSymbolId) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.FormattedSymbolId_orderStaged[order]
}

func (*FormattedSymbolId) GongGetInstancesMapByName(stage *Stage) any {
	return stage.FormattedSymbolIds_mapString
}

func (*FormattedSymbolId) GongGetInstancesSet(stage *Stage) any {
	return &stage.FormattedSymbolIds
}

func (*FormattedSymbolId) GongNewInstance() any {
	return new(FormattedSymbolId)
}

func (*FormattedText) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.FormattedTexts, stage.FormattedText_stagedOrder)
}

func (*FormattedText) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.FormattedText_orderStaged[order]
}

func (*FormattedText) GongGetInstancesMapByName(stage *Stage) any {
	return stage.FormattedTexts_mapString
}

func (*FormattedText) GongGetInstancesSet(stage *Stage) any {
	return &stage.FormattedTexts
}

func (*FormattedText) GongNewInstance() any {
	return new(FormattedText)
}

func (*FormattedTextId) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.FormattedTextIds, stage.FormattedTextId_stagedOrder)
}

func (*FormattedTextId) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.FormattedTextId_orderStaged[order]
}

func (*FormattedTextId) GongGetInstancesMapByName(stage *Stage) any {
	return stage.FormattedTextIds_mapString
}

func (*FormattedTextId) GongGetInstancesSet(stage *Stage) any {
	return &stage.FormattedTextIds
}

func (*FormattedTextId) GongNewInstance() any {
	return new(FormattedTextId)
}

func (*Fret) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Frets, stage.Fret_stagedOrder)
}

func (*Fret) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Fret_orderStaged[order]
}

func (*Fret) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Frets_mapString
}

func (*Fret) GongGetInstancesSet(stage *Stage) any {
	return &stage.Frets
}

func (*Fret) GongNewInstance() any {
	return new(Fret)
}

func (*Glyph) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Glyphs, stage.Glyph_stagedOrder)
}

func (*Glyph) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Glyph_orderStaged[order]
}

func (*Glyph) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Glyphs_mapString
}

func (*Glyph) GongGetInstancesSet(stage *Stage) any {
	return &stage.Glyphs
}

func (*Glyph) GongNewInstance() any {
	return new(Glyph)
}

func (*HorizontalTurn) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.HorizontalTurns, stage.HorizontalTurn_stagedOrder)
}

func (*HorizontalTurn) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.HorizontalTurn_orderStaged[order]
}

func (*HorizontalTurn) GongGetInstancesMapByName(stage *Stage) any {
	return stage.HorizontalTurns_mapString
}

func (*HorizontalTurn) GongGetInstancesSet(stage *Stage) any {
	return &stage.HorizontalTurns
}

func (*HorizontalTurn) GongNewInstance() any {
	return new(HorizontalTurn)
}

func (*Identification) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Identifications, stage.Identification_stagedOrder)
}

func (*Identification) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Identification_orderStaged[order]
}

func (*Identification) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Identifications_mapString
}

func (*Identification) GongGetInstancesSet(stage *Stage) any {
	return &stage.Identifications
}

func (*Identification) GongNewInstance() any {
	return new(Identification)
}

func (*Level) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Levels, stage.Level_stagedOrder)
}

func (*Level) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Level_orderStaged[order]
}

func (*Level) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Levels_mapString
}

func (*Level) GongGetInstancesSet(stage *Stage) any {
	return &stage.Levels
}

func (*Level) GongNewInstance() any {
	return new(Level)
}

func (*LineWidth) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.LineWidths, stage.LineWidth_stagedOrder)
}

func (*LineWidth) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.LineWidth_orderStaged[order]
}

func (*LineWidth) GongGetInstancesMapByName(stage *Stage) any {
	return stage.LineWidths_mapString
}

func (*LineWidth) GongGetInstancesSet(stage *Stage) any {
	return &stage.LineWidths
}

func (*LineWidth) GongNewInstance() any {
	return new(LineWidth)
}

func (*Link) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Links, stage.Link_stagedOrder)
}

func (*Link) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Link_orderStaged[order]
}

func (*Link) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Links_mapString
}

func (*Link) GongGetInstancesSet(stage *Stage) any {
	return &stage.Links
}

func (*Link) GongNewInstance() any {
	return new(Link)
}

func (*MeasureLayout) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MeasureLayouts, stage.MeasureLayout_stagedOrder)
}

func (*MeasureLayout) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MeasureLayout_orderStaged[order]
}

func (*MeasureLayout) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MeasureLayouts_mapString
}

func (*MeasureLayout) GongGetInstancesSet(stage *Stage) any {
	return &stage.MeasureLayouts
}

func (*MeasureLayout) GongNewInstance() any {
	return new(MeasureLayout)
}

func (*MidiDevice) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MidiDevices, stage.MidiDevice_stagedOrder)
}

func (*MidiDevice) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MidiDevice_orderStaged[order]
}

func (*MidiDevice) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MidiDevices_mapString
}

func (*MidiDevice) GongGetInstancesSet(stage *Stage) any {
	return &stage.MidiDevices
}

func (*MidiDevice) GongNewInstance() any {
	return new(MidiDevice)
}

func (*MidiInstrument) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MidiInstruments, stage.MidiInstrument_stagedOrder)
}

func (*MidiInstrument) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MidiInstrument_orderStaged[order]
}

func (*MidiInstrument) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MidiInstruments_mapString
}

func (*MidiInstrument) GongGetInstancesSet(stage *Stage) any {
	return &stage.MidiInstruments
}

func (*MidiInstrument) GongNewInstance() any {
	return new(MidiInstrument)
}

func (*Miscellaneous) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Miscellaneouss, stage.Miscellaneous_stagedOrder)
}

func (*Miscellaneous) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Miscellaneous_orderStaged[order]
}

func (*Miscellaneous) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Miscellaneouss_mapString
}

func (*Miscellaneous) GongGetInstancesSet(stage *Stage) any {
	return &stage.Miscellaneouss
}

func (*Miscellaneous) GongNewInstance() any {
	return new(Miscellaneous)
}

func (*MiscellaneousField) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.MiscellaneousFields, stage.MiscellaneousField_stagedOrder)
}

func (*MiscellaneousField) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.MiscellaneousField_orderStaged[order]
}

func (*MiscellaneousField) GongGetInstancesMapByName(stage *Stage) any {
	return stage.MiscellaneousFields_mapString
}

func (*MiscellaneousField) GongGetInstancesSet(stage *Stage) any {
	return &stage.MiscellaneousFields
}

func (*MiscellaneousField) GongNewInstance() any {
	return new(MiscellaneousField)
}

func (*NameDisplay) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NameDisplays, stage.NameDisplay_stagedOrder)
}

func (*NameDisplay) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NameDisplay_orderStaged[order]
}

func (*NameDisplay) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NameDisplays_mapString
}

func (*NameDisplay) GongGetInstancesSet(stage *Stage) any {
	return &stage.NameDisplays
}

func (*NameDisplay) GongNewInstance() any {
	return new(NameDisplay)
}

func (*NoteSize) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.NoteSizes, stage.NoteSize_stagedOrder)
}

func (*NoteSize) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.NoteSize_orderStaged[order]
}

func (*NoteSize) GongGetInstancesMapByName(stage *Stage) any {
	return stage.NoteSizes_mapString
}

func (*NoteSize) GongGetInstancesSet(stage *Stage) any {
	return &stage.NoteSizes
}

func (*NoteSize) GongNewInstance() any {
	return new(NoteSize)
}

func (*OtherAppearance) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OtherAppearances, stage.OtherAppearance_stagedOrder)
}

func (*OtherAppearance) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OtherAppearance_orderStaged[order]
}

func (*OtherAppearance) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OtherAppearances_mapString
}

func (*OtherAppearance) GongGetInstancesSet(stage *Stage) any {
	return &stage.OtherAppearances
}

func (*OtherAppearance) GongNewInstance() any {
	return new(OtherAppearance)
}

func (*OtherPlay) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OtherPlays, stage.OtherPlay_stagedOrder)
}

func (*OtherPlay) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OtherPlay_orderStaged[order]
}

func (*OtherPlay) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OtherPlays_mapString
}

func (*OtherPlay) GongGetInstancesSet(stage *Stage) any {
	return &stage.OtherPlays
}

func (*OtherPlay) GongNewInstance() any {
	return new(OtherPlay)
}

func (*OtherText) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.OtherTexts, stage.OtherText_stagedOrder)
}

func (*OtherText) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.OtherText_orderStaged[order]
}

func (*OtherText) GongGetInstancesMapByName(stage *Stage) any {
	return stage.OtherTexts_mapString
}

func (*OtherText) GongGetInstancesSet(stage *Stage) any {
	return &stage.OtherTexts
}

func (*OtherText) GongNewInstance() any {
	return new(OtherText)
}

func (*PageLayout) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PageLayouts, stage.PageLayout_stagedOrder)
}

func (*PageLayout) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PageLayout_orderStaged[order]
}

func (*PageLayout) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PageLayouts_mapString
}

func (*PageLayout) GongGetInstancesSet(stage *Stage) any {
	return &stage.PageLayouts
}

func (*PageLayout) GongNewInstance() any {
	return new(PageLayout)
}

func (*PageMargins) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.PageMarginss, stage.PageMargins_stagedOrder)
}

func (*PageMargins) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.PageMargins_orderStaged[order]
}

func (*PageMargins) GongGetInstancesMapByName(stage *Stage) any {
	return stage.PageMarginss_mapString
}

func (*PageMargins) GongGetInstancesSet(stage *Stage) any {
	return &stage.PageMarginss
}

func (*PageMargins) GongNewInstance() any {
	return new(PageMargins)
}

func (*Play) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Plays, stage.Play_stagedOrder)
}

func (*Play) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Play_orderStaged[order]
}

func (*Play) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Plays_mapString
}

func (*Play) GongGetInstancesSet(stage *Stage) any {
	return &stage.Plays
}

func (*Play) GongNewInstance() any {
	return new(Play)
}

func (*Scaling) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Scalings, stage.Scaling_stagedOrder)
}

func (*Scaling) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Scaling_orderStaged[order]
}

func (*Scaling) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Scalings_mapString
}

func (*Scaling) GongGetInstancesSet(stage *Stage) any {
	return &stage.Scalings
}

func (*Scaling) GongNewInstance() any {
	return new(Scaling)
}

func (*Segno) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Segnos, stage.Segno_stagedOrder)
}

func (*Segno) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Segno_orderStaged[order]
}

func (*Segno) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Segnos_mapString
}

func (*Segno) GongGetInstancesSet(stage *Stage) any {
	return &stage.Segnos
}

func (*Segno) GongNewInstance() any {
	return new(Segno)
}

func (*StaffLayout) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StaffLayouts, stage.StaffLayout_stagedOrder)
}

func (*StaffLayout) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StaffLayout_orderStaged[order]
}

func (*StaffLayout) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StaffLayouts_mapString
}

func (*StaffLayout) GongGetInstancesSet(stage *Stage) any {
	return &stage.StaffLayouts
}

func (*StaffLayout) GongNewInstance() any {
	return new(StaffLayout)
}

func (*StringType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.StringTypes, stage.StringType_stagedOrder)
}

func (*StringType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.StringType_orderStaged[order]
}

func (*StringType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.StringTypes_mapString
}

func (*StringType) GongGetInstancesSet(stage *Stage) any {
	return &stage.StringTypes
}

func (*StringType) GongNewInstance() any {
	return new(StringType)
}

func (*Supports) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.Supportss, stage.Supports_stagedOrder)
}

func (*Supports) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.Supports_orderStaged[order]
}

func (*Supports) GongGetInstancesMapByName(stage *Stage) any {
	return stage.Supportss_mapString
}

func (*Supports) GongGetInstancesSet(stage *Stage) any {
	return &stage.Supportss
}

func (*Supports) GongNewInstance() any {
	return new(Supports)
}

func (*SystemDividers) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.SystemDividerss, stage.SystemDividers_stagedOrder)
}

func (*SystemDividers) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.SystemDividers_orderStaged[order]
}

func (*SystemDividers) GongGetInstancesMapByName(stage *Stage) any {
	return stage.SystemDividerss_mapString
}

func (*SystemDividers) GongGetInstancesSet(stage *Stage) any {
	return &stage.SystemDividerss
}

func (*SystemDividers) GongNewInstance() any {
	return new(SystemDividers)
}

func (*SystemLayout) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.SystemLayouts, stage.SystemLayout_stagedOrder)
}

func (*SystemLayout) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.SystemLayout_orderStaged[order]
}

func (*SystemLayout) GongGetInstancesMapByName(stage *Stage) any {
	return stage.SystemLayouts_mapString
}

func (*SystemLayout) GongGetInstancesSet(stage *Stage) any {
	return &stage.SystemLayouts
}

func (*SystemLayout) GongNewInstance() any {
	return new(SystemLayout)
}

func (*SystemMargins) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.SystemMarginss, stage.SystemMargins_stagedOrder)
}

func (*SystemMargins) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.SystemMargins_orderStaged[order]
}

func (*SystemMargins) GongGetInstancesMapByName(stage *Stage) any {
	return stage.SystemMarginss_mapString
}

func (*SystemMargins) GongGetInstancesSet(stage *Stage) any {
	return &stage.SystemMarginss
}

func (*SystemMargins) GongNewInstance() any {
	return new(SystemMargins)
}

func (*TypedText) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.TypedTexts, stage.TypedText_stagedOrder)
}

func (*TypedText) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.TypedText_orderStaged[order]
}

func (*TypedText) GongGetInstancesMapByName(stage *Stage) any {
	return stage.TypedTexts_mapString
}

func (*TypedText) GongGetInstancesSet(stage *Stage) any {
	return &stage.TypedTexts
}

func (*TypedText) GongNewInstance() any {
	return new(TypedText)
}

func (*WavyLine) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.WavyLines, stage.WavyLine_stagedOrder)
}

func (*WavyLine) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.WavyLine_orderStaged[order]
}

func (*WavyLine) GongGetInstancesMapByName(stage *Stage) any {
	return stage.WavyLines_mapString
}

func (*WavyLine) GongGetInstancesSet(stage *Stage) any {
	return &stage.WavyLines
}

func (*WavyLine) GongNewInstance() any {
	return new(WavyLine)
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
	return "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
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
		AccidentalTexts:           make(map[*AccidentalText]struct{}),
		AccidentalTexts_mapString: make(map[string]*AccidentalText),

		Appearances:           make(map[*Appearance]struct{}),
		Appearances_mapString: make(map[string]*Appearance),

		Bookmarks:           make(map[*Bookmark]struct{}),
		Bookmarks_mapString: make(map[string]*Bookmark),

		Codas:           make(map[*Coda]struct{}),
		Codas_mapString: make(map[string]*Coda),

		Distances:           make(map[*Distance]struct{}),
		Distances_mapString: make(map[string]*Distance),

		Dynamicss:           make(map[*Dynamics]struct{}),
		Dynamicss_mapString: make(map[string]*Dynamics),

		Emptys:           make(map[*Empty]struct{}),
		Emptys_mapString: make(map[string]*Empty),

		EmptyPlacements:           make(map[*EmptyPlacement]struct{}),
		EmptyPlacements_mapString: make(map[string]*EmptyPlacement),

		EmptyPlacementSmufls:           make(map[*EmptyPlacementSmufl]struct{}),
		EmptyPlacementSmufls_mapString: make(map[string]*EmptyPlacementSmufl),

		EmptyPrintObjectStyleAligns:           make(map[*EmptyPrintObjectStyleAlign]struct{}),
		EmptyPrintObjectStyleAligns_mapString: make(map[string]*EmptyPrintObjectStyleAlign),

		EmptyPrintStyles:           make(map[*EmptyPrintStyle]struct{}),
		EmptyPrintStyles_mapString: make(map[string]*EmptyPrintStyle),

		EmptyPrintStyleAligns:           make(map[*EmptyPrintStyleAlign]struct{}),
		EmptyPrintStyleAligns_mapString: make(map[string]*EmptyPrintStyleAlign),

		EmptyPrintStyleAlignIds:           make(map[*EmptyPrintStyleAlignId]struct{}),
		EmptyPrintStyleAlignIds_mapString: make(map[string]*EmptyPrintStyleAlignId),

		EmptyTrillSounds:           make(map[*EmptyTrillSound]struct{}),
		EmptyTrillSounds_mapString: make(map[string]*EmptyTrillSound),

		Encodings:           make(map[*Encoding]struct{}),
		Encodings_mapString: make(map[string]*Encoding),

		Fermatas:           make(map[*Fermata]struct{}),
		Fermatas_mapString: make(map[string]*Fermata),

		Fingerings:           make(map[*Fingering]struct{}),
		Fingerings_mapString: make(map[string]*Fingering),

		FormattedSymbols:           make(map[*FormattedSymbol]struct{}),
		FormattedSymbols_mapString: make(map[string]*FormattedSymbol),

		FormattedSymbolIds:           make(map[*FormattedSymbolId]struct{}),
		FormattedSymbolIds_mapString: make(map[string]*FormattedSymbolId),

		FormattedTexts:           make(map[*FormattedText]struct{}),
		FormattedTexts_mapString: make(map[string]*FormattedText),

		FormattedTextIds:           make(map[*FormattedTextId]struct{}),
		FormattedTextIds_mapString: make(map[string]*FormattedTextId),

		Frets:           make(map[*Fret]struct{}),
		Frets_mapString: make(map[string]*Fret),

		Glyphs:           make(map[*Glyph]struct{}),
		Glyphs_mapString: make(map[string]*Glyph),

		HorizontalTurns:           make(map[*HorizontalTurn]struct{}),
		HorizontalTurns_mapString: make(map[string]*HorizontalTurn),

		Identifications:           make(map[*Identification]struct{}),
		Identifications_mapString: make(map[string]*Identification),

		Levels:           make(map[*Level]struct{}),
		Levels_mapString: make(map[string]*Level),

		LineWidths:           make(map[*LineWidth]struct{}),
		LineWidths_mapString: make(map[string]*LineWidth),

		Links:           make(map[*Link]struct{}),
		Links_mapString: make(map[string]*Link),

		MeasureLayouts:           make(map[*MeasureLayout]struct{}),
		MeasureLayouts_mapString: make(map[string]*MeasureLayout),

		MidiDevices:           make(map[*MidiDevice]struct{}),
		MidiDevices_mapString: make(map[string]*MidiDevice),

		MidiInstruments:           make(map[*MidiInstrument]struct{}),
		MidiInstruments_mapString: make(map[string]*MidiInstrument),

		Miscellaneouss:           make(map[*Miscellaneous]struct{}),
		Miscellaneouss_mapString: make(map[string]*Miscellaneous),

		MiscellaneousFields:           make(map[*MiscellaneousField]struct{}),
		MiscellaneousFields_mapString: make(map[string]*MiscellaneousField),

		NameDisplays:           make(map[*NameDisplay]struct{}),
		NameDisplays_mapString: make(map[string]*NameDisplay),

		NoteSizes:           make(map[*NoteSize]struct{}),
		NoteSizes_mapString: make(map[string]*NoteSize),

		OtherAppearances:           make(map[*OtherAppearance]struct{}),
		OtherAppearances_mapString: make(map[string]*OtherAppearance),

		OtherPlays:           make(map[*OtherPlay]struct{}),
		OtherPlays_mapString: make(map[string]*OtherPlay),

		OtherTexts:           make(map[*OtherText]struct{}),
		OtherTexts_mapString: make(map[string]*OtherText),

		PageLayouts:           make(map[*PageLayout]struct{}),
		PageLayouts_mapString: make(map[string]*PageLayout),

		PageMarginss:           make(map[*PageMargins]struct{}),
		PageMarginss_mapString: make(map[string]*PageMargins),

		Plays:           make(map[*Play]struct{}),
		Plays_mapString: make(map[string]*Play),

		Scalings:           make(map[*Scaling]struct{}),
		Scalings_mapString: make(map[string]*Scaling),

		Segnos:           make(map[*Segno]struct{}),
		Segnos_mapString: make(map[string]*Segno),

		StaffLayouts:           make(map[*StaffLayout]struct{}),
		StaffLayouts_mapString: make(map[string]*StaffLayout),

		StringTypes:           make(map[*StringType]struct{}),
		StringTypes_mapString: make(map[string]*StringType),

		Supportss:           make(map[*Supports]struct{}),
		Supportss_mapString: make(map[string]*Supports),

		SystemDividerss:           make(map[*SystemDividers]struct{}),
		SystemDividerss_mapString: make(map[string]*SystemDividers),

		SystemLayouts:           make(map[*SystemLayout]struct{}),
		SystemLayouts_mapString: make(map[string]*SystemLayout),

		SystemMarginss:           make(map[*SystemMargins]struct{}),
		SystemMarginss_mapString: make(map[string]*SystemMargins),

		TypedTexts:           make(map[*TypedText]struct{}),
		TypedTexts_mapString: make(map[string]*TypedText),

		WavyLines:           make(map[*WavyLine]struct{}),
		WavyLines_mapString: make(map[string]*WavyLine),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AccidentalText_stagedOrder: make(map[*AccidentalText]uint),
		AccidentalText_orderStaged: make(map[uint]*AccidentalText),
		AccidentalTexts_reference:  make(map[*AccidentalText]*AccidentalText),

		Appearance_stagedOrder: make(map[*Appearance]uint),
		Appearance_orderStaged: make(map[uint]*Appearance),
		Appearances_reference:  make(map[*Appearance]*Appearance),

		Bookmark_stagedOrder: make(map[*Bookmark]uint),
		Bookmark_orderStaged: make(map[uint]*Bookmark),
		Bookmarks_reference:  make(map[*Bookmark]*Bookmark),

		Coda_stagedOrder: make(map[*Coda]uint),
		Coda_orderStaged: make(map[uint]*Coda),
		Codas_reference:  make(map[*Coda]*Coda),

		Distance_stagedOrder: make(map[*Distance]uint),
		Distance_orderStaged: make(map[uint]*Distance),
		Distances_reference:  make(map[*Distance]*Distance),

		Dynamics_stagedOrder: make(map[*Dynamics]uint),
		Dynamics_orderStaged: make(map[uint]*Dynamics),
		Dynamicss_reference:  make(map[*Dynamics]*Dynamics),

		Empty_stagedOrder: make(map[*Empty]uint),
		Empty_orderStaged: make(map[uint]*Empty),
		Emptys_reference:  make(map[*Empty]*Empty),

		EmptyPlacement_stagedOrder: make(map[*EmptyPlacement]uint),
		EmptyPlacement_orderStaged: make(map[uint]*EmptyPlacement),
		EmptyPlacements_reference:  make(map[*EmptyPlacement]*EmptyPlacement),

		EmptyPlacementSmufl_stagedOrder: make(map[*EmptyPlacementSmufl]uint),
		EmptyPlacementSmufl_orderStaged: make(map[uint]*EmptyPlacementSmufl),
		EmptyPlacementSmufls_reference:  make(map[*EmptyPlacementSmufl]*EmptyPlacementSmufl),

		EmptyPrintObjectStyleAlign_stagedOrder: make(map[*EmptyPrintObjectStyleAlign]uint),
		EmptyPrintObjectStyleAlign_orderStaged: make(map[uint]*EmptyPrintObjectStyleAlign),
		EmptyPrintObjectStyleAligns_reference:  make(map[*EmptyPrintObjectStyleAlign]*EmptyPrintObjectStyleAlign),

		EmptyPrintStyle_stagedOrder: make(map[*EmptyPrintStyle]uint),
		EmptyPrintStyle_orderStaged: make(map[uint]*EmptyPrintStyle),
		EmptyPrintStyles_reference:  make(map[*EmptyPrintStyle]*EmptyPrintStyle),

		EmptyPrintStyleAlign_stagedOrder: make(map[*EmptyPrintStyleAlign]uint),
		EmptyPrintStyleAlign_orderStaged: make(map[uint]*EmptyPrintStyleAlign),
		EmptyPrintStyleAligns_reference:  make(map[*EmptyPrintStyleAlign]*EmptyPrintStyleAlign),

		EmptyPrintStyleAlignId_stagedOrder: make(map[*EmptyPrintStyleAlignId]uint),
		EmptyPrintStyleAlignId_orderStaged: make(map[uint]*EmptyPrintStyleAlignId),
		EmptyPrintStyleAlignIds_reference:  make(map[*EmptyPrintStyleAlignId]*EmptyPrintStyleAlignId),

		EmptyTrillSound_stagedOrder: make(map[*EmptyTrillSound]uint),
		EmptyTrillSound_orderStaged: make(map[uint]*EmptyTrillSound),
		EmptyTrillSounds_reference:  make(map[*EmptyTrillSound]*EmptyTrillSound),

		Encoding_stagedOrder: make(map[*Encoding]uint),
		Encoding_orderStaged: make(map[uint]*Encoding),
		Encodings_reference:  make(map[*Encoding]*Encoding),

		Fermata_stagedOrder: make(map[*Fermata]uint),
		Fermata_orderStaged: make(map[uint]*Fermata),
		Fermatas_reference:  make(map[*Fermata]*Fermata),

		Fingering_stagedOrder: make(map[*Fingering]uint),
		Fingering_orderStaged: make(map[uint]*Fingering),
		Fingerings_reference:  make(map[*Fingering]*Fingering),

		FormattedSymbol_stagedOrder: make(map[*FormattedSymbol]uint),
		FormattedSymbol_orderStaged: make(map[uint]*FormattedSymbol),
		FormattedSymbols_reference:  make(map[*FormattedSymbol]*FormattedSymbol),

		FormattedSymbolId_stagedOrder: make(map[*FormattedSymbolId]uint),
		FormattedSymbolId_orderStaged: make(map[uint]*FormattedSymbolId),
		FormattedSymbolIds_reference:  make(map[*FormattedSymbolId]*FormattedSymbolId),

		FormattedText_stagedOrder: make(map[*FormattedText]uint),
		FormattedText_orderStaged: make(map[uint]*FormattedText),
		FormattedTexts_reference:  make(map[*FormattedText]*FormattedText),

		FormattedTextId_stagedOrder: make(map[*FormattedTextId]uint),
		FormattedTextId_orderStaged: make(map[uint]*FormattedTextId),
		FormattedTextIds_reference:  make(map[*FormattedTextId]*FormattedTextId),

		Fret_stagedOrder: make(map[*Fret]uint),
		Fret_orderStaged: make(map[uint]*Fret),
		Frets_reference:  make(map[*Fret]*Fret),

		Glyph_stagedOrder: make(map[*Glyph]uint),
		Glyph_orderStaged: make(map[uint]*Glyph),
		Glyphs_reference:  make(map[*Glyph]*Glyph),

		HorizontalTurn_stagedOrder: make(map[*HorizontalTurn]uint),
		HorizontalTurn_orderStaged: make(map[uint]*HorizontalTurn),
		HorizontalTurns_reference:  make(map[*HorizontalTurn]*HorizontalTurn),

		Identification_stagedOrder: make(map[*Identification]uint),
		Identification_orderStaged: make(map[uint]*Identification),
		Identifications_reference:  make(map[*Identification]*Identification),

		Level_stagedOrder: make(map[*Level]uint),
		Level_orderStaged: make(map[uint]*Level),
		Levels_reference:  make(map[*Level]*Level),

		LineWidth_stagedOrder: make(map[*LineWidth]uint),
		LineWidth_orderStaged: make(map[uint]*LineWidth),
		LineWidths_reference:  make(map[*LineWidth]*LineWidth),

		Link_stagedOrder: make(map[*Link]uint),
		Link_orderStaged: make(map[uint]*Link),
		Links_reference:  make(map[*Link]*Link),

		MeasureLayout_stagedOrder: make(map[*MeasureLayout]uint),
		MeasureLayout_orderStaged: make(map[uint]*MeasureLayout),
		MeasureLayouts_reference:  make(map[*MeasureLayout]*MeasureLayout),

		MidiDevice_stagedOrder: make(map[*MidiDevice]uint),
		MidiDevice_orderStaged: make(map[uint]*MidiDevice),
		MidiDevices_reference:  make(map[*MidiDevice]*MidiDevice),

		MidiInstrument_stagedOrder: make(map[*MidiInstrument]uint),
		MidiInstrument_orderStaged: make(map[uint]*MidiInstrument),
		MidiInstruments_reference:  make(map[*MidiInstrument]*MidiInstrument),

		Miscellaneous_stagedOrder: make(map[*Miscellaneous]uint),
		Miscellaneous_orderStaged: make(map[uint]*Miscellaneous),
		Miscellaneouss_reference:  make(map[*Miscellaneous]*Miscellaneous),

		MiscellaneousField_stagedOrder: make(map[*MiscellaneousField]uint),
		MiscellaneousField_orderStaged: make(map[uint]*MiscellaneousField),
		MiscellaneousFields_reference:  make(map[*MiscellaneousField]*MiscellaneousField),

		NameDisplay_stagedOrder: make(map[*NameDisplay]uint),
		NameDisplay_orderStaged: make(map[uint]*NameDisplay),
		NameDisplays_reference:  make(map[*NameDisplay]*NameDisplay),

		NoteSize_stagedOrder: make(map[*NoteSize]uint),
		NoteSize_orderStaged: make(map[uint]*NoteSize),
		NoteSizes_reference:  make(map[*NoteSize]*NoteSize),

		OtherAppearance_stagedOrder: make(map[*OtherAppearance]uint),
		OtherAppearance_orderStaged: make(map[uint]*OtherAppearance),
		OtherAppearances_reference:  make(map[*OtherAppearance]*OtherAppearance),

		OtherPlay_stagedOrder: make(map[*OtherPlay]uint),
		OtherPlay_orderStaged: make(map[uint]*OtherPlay),
		OtherPlays_reference:  make(map[*OtherPlay]*OtherPlay),

		OtherText_stagedOrder: make(map[*OtherText]uint),
		OtherText_orderStaged: make(map[uint]*OtherText),
		OtherTexts_reference:  make(map[*OtherText]*OtherText),

		PageLayout_stagedOrder: make(map[*PageLayout]uint),
		PageLayout_orderStaged: make(map[uint]*PageLayout),
		PageLayouts_reference:  make(map[*PageLayout]*PageLayout),

		PageMargins_stagedOrder: make(map[*PageMargins]uint),
		PageMargins_orderStaged: make(map[uint]*PageMargins),
		PageMarginss_reference:  make(map[*PageMargins]*PageMargins),

		Play_stagedOrder: make(map[*Play]uint),
		Play_orderStaged: make(map[uint]*Play),
		Plays_reference:  make(map[*Play]*Play),

		Scaling_stagedOrder: make(map[*Scaling]uint),
		Scaling_orderStaged: make(map[uint]*Scaling),
		Scalings_reference:  make(map[*Scaling]*Scaling),

		Segno_stagedOrder: make(map[*Segno]uint),
		Segno_orderStaged: make(map[uint]*Segno),
		Segnos_reference:  make(map[*Segno]*Segno),

		StaffLayout_stagedOrder: make(map[*StaffLayout]uint),
		StaffLayout_orderStaged: make(map[uint]*StaffLayout),
		StaffLayouts_reference:  make(map[*StaffLayout]*StaffLayout),

		StringType_stagedOrder: make(map[*StringType]uint),
		StringType_orderStaged: make(map[uint]*StringType),
		StringTypes_reference:  make(map[*StringType]*StringType),

		Supports_stagedOrder: make(map[*Supports]uint),
		Supports_orderStaged: make(map[uint]*Supports),
		Supportss_reference:  make(map[*Supports]*Supports),

		SystemDividers_stagedOrder: make(map[*SystemDividers]uint),
		SystemDividers_orderStaged: make(map[uint]*SystemDividers),
		SystemDividerss_reference:  make(map[*SystemDividers]*SystemDividers),

		SystemLayout_stagedOrder: make(map[*SystemLayout]uint),
		SystemLayout_orderStaged: make(map[uint]*SystemLayout),
		SystemLayouts_reference:  make(map[*SystemLayout]*SystemLayout),

		SystemMargins_stagedOrder: make(map[*SystemMargins]uint),
		SystemMargins_orderStaged: make(map[uint]*SystemMargins),
		SystemMarginss_reference:  make(map[*SystemMargins]*SystemMargins),

		TypedText_stagedOrder: make(map[*TypedText]uint),
		TypedText_orderStaged: make(map[uint]*TypedText),
		TypedTexts_reference:  make(map[*TypedText]*TypedText),

		WavyLine_stagedOrder: make(map[*WavyLine]uint),
		WavyLine_orderStaged: make(map[uint]*WavyLine),
		WavyLines_reference:  make(map[*WavyLine]*WavyLine),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AccidentalText": &AccidentalTextUnmarshaller{},

			"Appearance": &AppearanceUnmarshaller{},

			"Bookmark": &BookmarkUnmarshaller{},

			"Coda": &CodaUnmarshaller{},

			"Distance": &DistanceUnmarshaller{},

			"Dynamics": &DynamicsUnmarshaller{},

			"Empty": &EmptyUnmarshaller{},

			"EmptyPlacement": &EmptyPlacementUnmarshaller{},

			"EmptyPlacementSmufl": &EmptyPlacementSmuflUnmarshaller{},

			"EmptyPrintObjectStyleAlign": &EmptyPrintObjectStyleAlignUnmarshaller{},

			"EmptyPrintStyle": &EmptyPrintStyleUnmarshaller{},

			"EmptyPrintStyleAlign": &EmptyPrintStyleAlignUnmarshaller{},

			"EmptyPrintStyleAlignId": &EmptyPrintStyleAlignIdUnmarshaller{},

			"EmptyTrillSound": &EmptyTrillSoundUnmarshaller{},

			"Encoding": &EncodingUnmarshaller{},

			"Fermata": &FermataUnmarshaller{},

			"Fingering": &FingeringUnmarshaller{},

			"FormattedSymbol": &FormattedSymbolUnmarshaller{},

			"FormattedSymbolId": &FormattedSymbolIdUnmarshaller{},

			"FormattedText": &FormattedTextUnmarshaller{},

			"FormattedTextId": &FormattedTextIdUnmarshaller{},

			"Fret": &FretUnmarshaller{},

			"Glyph": &GlyphUnmarshaller{},

			"HorizontalTurn": &HorizontalTurnUnmarshaller{},

			"Identification": &IdentificationUnmarshaller{},

			"Level": &LevelUnmarshaller{},

			"LineWidth": &LineWidthUnmarshaller{},

			"Link": &LinkUnmarshaller{},

			"MeasureLayout": &MeasureLayoutUnmarshaller{},

			"MidiDevice": &MidiDeviceUnmarshaller{},

			"MidiInstrument": &MidiInstrumentUnmarshaller{},

			"Miscellaneous": &MiscellaneousUnmarshaller{},

			"MiscellaneousField": &MiscellaneousFieldUnmarshaller{},

			"NameDisplay": &NameDisplayUnmarshaller{},

			"NoteSize": &NoteSizeUnmarshaller{},

			"OtherAppearance": &OtherAppearanceUnmarshaller{},

			"OtherPlay": &OtherPlayUnmarshaller{},

			"OtherText": &OtherTextUnmarshaller{},

			"PageLayout": &PageLayoutUnmarshaller{},

			"PageMargins": &PageMarginsUnmarshaller{},

			"Play": &PlayUnmarshaller{},

			"Scaling": &ScalingUnmarshaller{},

			"Segno": &SegnoUnmarshaller{},

			"StaffLayout": &StaffLayoutUnmarshaller{},

			"StringType": &StringTypeUnmarshaller{},

			"Supports": &SupportsUnmarshaller{},

			"SystemDividers": &SystemDividersUnmarshaller{},

			"SystemLayout": &SystemLayoutUnmarshaller{},

			"SystemMargins": &SystemMarginsUnmarshaller{},

			"TypedText": &TypedTextUnmarshaller{},

			"WavyLine": &WavyLineUnmarshaller{},

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
	stage.Map_GongStructName_InstancesNb["AccidentalText"] = len(stage.AccidentalTexts)
	stage.Map_GongStructName_InstancesNb["Appearance"] = len(stage.Appearances)
	stage.Map_GongStructName_InstancesNb["Bookmark"] = len(stage.Bookmarks)
	stage.Map_GongStructName_InstancesNb["Coda"] = len(stage.Codas)
	stage.Map_GongStructName_InstancesNb["Distance"] = len(stage.Distances)
	stage.Map_GongStructName_InstancesNb["Dynamics"] = len(stage.Dynamicss)
	stage.Map_GongStructName_InstancesNb["Empty"] = len(stage.Emptys)
	stage.Map_GongStructName_InstancesNb["EmptyPlacement"] = len(stage.EmptyPlacements)
	stage.Map_GongStructName_InstancesNb["EmptyPlacementSmufl"] = len(stage.EmptyPlacementSmufls)
	stage.Map_GongStructName_InstancesNb["EmptyPrintObjectStyleAlign"] = len(stage.EmptyPrintObjectStyleAligns)
	stage.Map_GongStructName_InstancesNb["EmptyPrintStyle"] = len(stage.EmptyPrintStyles)
	stage.Map_GongStructName_InstancesNb["EmptyPrintStyleAlign"] = len(stage.EmptyPrintStyleAligns)
	stage.Map_GongStructName_InstancesNb["EmptyPrintStyleAlignId"] = len(stage.EmptyPrintStyleAlignIds)
	stage.Map_GongStructName_InstancesNb["EmptyTrillSound"] = len(stage.EmptyTrillSounds)
	stage.Map_GongStructName_InstancesNb["Encoding"] = len(stage.Encodings)
	stage.Map_GongStructName_InstancesNb["Fermata"] = len(stage.Fermatas)
	stage.Map_GongStructName_InstancesNb["Fingering"] = len(stage.Fingerings)
	stage.Map_GongStructName_InstancesNb["FormattedSymbol"] = len(stage.FormattedSymbols)
	stage.Map_GongStructName_InstancesNb["FormattedSymbolId"] = len(stage.FormattedSymbolIds)
	stage.Map_GongStructName_InstancesNb["FormattedText"] = len(stage.FormattedTexts)
	stage.Map_GongStructName_InstancesNb["FormattedTextId"] = len(stage.FormattedTextIds)
	stage.Map_GongStructName_InstancesNb["Fret"] = len(stage.Frets)
	stage.Map_GongStructName_InstancesNb["Glyph"] = len(stage.Glyphs)
	stage.Map_GongStructName_InstancesNb["HorizontalTurn"] = len(stage.HorizontalTurns)
	stage.Map_GongStructName_InstancesNb["Identification"] = len(stage.Identifications)
	stage.Map_GongStructName_InstancesNb["Level"] = len(stage.Levels)
	stage.Map_GongStructName_InstancesNb["LineWidth"] = len(stage.LineWidths)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
	stage.Map_GongStructName_InstancesNb["MeasureLayout"] = len(stage.MeasureLayouts)
	stage.Map_GongStructName_InstancesNb["MidiDevice"] = len(stage.MidiDevices)
	stage.Map_GongStructName_InstancesNb["MidiInstrument"] = len(stage.MidiInstruments)
	stage.Map_GongStructName_InstancesNb["Miscellaneous"] = len(stage.Miscellaneouss)
	stage.Map_GongStructName_InstancesNb["MiscellaneousField"] = len(stage.MiscellaneousFields)
	stage.Map_GongStructName_InstancesNb["NameDisplay"] = len(stage.NameDisplays)
	stage.Map_GongStructName_InstancesNb["NoteSize"] = len(stage.NoteSizes)
	stage.Map_GongStructName_InstancesNb["OtherAppearance"] = len(stage.OtherAppearances)
	stage.Map_GongStructName_InstancesNb["OtherPlay"] = len(stage.OtherPlays)
	stage.Map_GongStructName_InstancesNb["OtherText"] = len(stage.OtherTexts)
	stage.Map_GongStructName_InstancesNb["PageLayout"] = len(stage.PageLayouts)
	stage.Map_GongStructName_InstancesNb["PageMargins"] = len(stage.PageMarginss)
	stage.Map_GongStructName_InstancesNb["Play"] = len(stage.Plays)
	stage.Map_GongStructName_InstancesNb["Scaling"] = len(stage.Scalings)
	stage.Map_GongStructName_InstancesNb["Segno"] = len(stage.Segnos)
	stage.Map_GongStructName_InstancesNb["StaffLayout"] = len(stage.StaffLayouts)
	stage.Map_GongStructName_InstancesNb["StringType"] = len(stage.StringTypes)
	stage.Map_GongStructName_InstancesNb["Supports"] = len(stage.Supportss)
	stage.Map_GongStructName_InstancesNb["SystemDividers"] = len(stage.SystemDividerss)
	stage.Map_GongStructName_InstancesNb["SystemLayout"] = len(stage.SystemLayouts)
	stage.Map_GongStructName_InstancesNb["SystemMargins"] = len(stage.SystemMarginss)
	stage.Map_GongStructName_InstancesNb["TypedText"] = len(stage.TypedTexts)
	stage.Map_GongStructName_InstancesNb["WavyLine"] = len(stage.WavyLines)
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
// Stage puts accidentaltext to the model stage
func (accidentaltext *AccidentalText) Stage(stage *Stage) *AccidentalText {
	__gong__stage(stage.AccidentalTexts, stage.AccidentalText_stagedOrder, stage.AccidentalText_orderStaged, &stage.AccidentalTextOrder, stage.AccidentalTexts_mapString, accidentaltext, accidentaltext.Name)
	return accidentaltext
}

// StagePreserveOrder puts accidentaltext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AccidentalTextOrder
// - update stage.AccidentalTextOrder accordingly
func (accidentaltext *AccidentalText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AccidentalTexts, stage.AccidentalText_stagedOrder, stage.AccidentalText_orderStaged, &stage.AccidentalTextOrder, stage.AccidentalTexts_mapString, accidentaltext, order, accidentaltext.Name)
}

// Unstage removes accidentaltext off the model stage
func (accidentaltext *AccidentalText) Unstage(stage *Stage) *AccidentalText {
	__gong__unstage(stage.AccidentalTexts, stage.AccidentalTexts_mapString, accidentaltext, accidentaltext.Name)
	return accidentaltext
}

// UnstageVoid removes accidentaltext off the model stage
func (accidentaltext *AccidentalText) UnstageVoid(stage *Stage) {
	accidentaltext.Unstage(stage)
}

func (accidentaltext *AccidentalText) StageVoid(stage *Stage) {
	accidentaltext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (accidentaltext *AccidentalText) GetName() (res string) {
	return accidentaltext.Name
}

// for satisfaction of GongStruct interface
func (accidentaltext *AccidentalText) SetName(name string) {
	accidentaltext.Name = name
}

// Stage puts appearance to the model stage
func (appearance *Appearance) Stage(stage *Stage) *Appearance {
	__gong__stage(stage.Appearances, stage.Appearance_stagedOrder, stage.Appearance_orderStaged, &stage.AppearanceOrder, stage.Appearances_mapString, appearance, appearance.Name)
	return appearance
}

// StagePreserveOrder puts appearance to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AppearanceOrder
// - update stage.AppearanceOrder accordingly
func (appearance *Appearance) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Appearances, stage.Appearance_stagedOrder, stage.Appearance_orderStaged, &stage.AppearanceOrder, stage.Appearances_mapString, appearance, order, appearance.Name)
}

// Unstage removes appearance off the model stage
func (appearance *Appearance) Unstage(stage *Stage) *Appearance {
	__gong__unstage(stage.Appearances, stage.Appearances_mapString, appearance, appearance.Name)
	return appearance
}

// UnstageVoid removes appearance off the model stage
func (appearance *Appearance) UnstageVoid(stage *Stage) {
	appearance.Unstage(stage)
}

func (appearance *Appearance) StageVoid(stage *Stage) {
	appearance.Stage(stage)
}

// for satisfaction of GongStruct interface
func (appearance *Appearance) GetName() (res string) {
	return appearance.Name
}

// for satisfaction of GongStruct interface
func (appearance *Appearance) SetName(name string) {
	appearance.Name = name
}

// Stage puts bookmark to the model stage
func (bookmark *Bookmark) Stage(stage *Stage) *Bookmark {
	__gong__stage(stage.Bookmarks, stage.Bookmark_stagedOrder, stage.Bookmark_orderStaged, &stage.BookmarkOrder, stage.Bookmarks_mapString, bookmark, bookmark.Name)
	return bookmark
}

// StagePreserveOrder puts bookmark to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BookmarkOrder
// - update stage.BookmarkOrder accordingly
func (bookmark *Bookmark) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Bookmarks, stage.Bookmark_stagedOrder, stage.Bookmark_orderStaged, &stage.BookmarkOrder, stage.Bookmarks_mapString, bookmark, order, bookmark.Name)
}

// Unstage removes bookmark off the model stage
func (bookmark *Bookmark) Unstage(stage *Stage) *Bookmark {
	__gong__unstage(stage.Bookmarks, stage.Bookmarks_mapString, bookmark, bookmark.Name)
	return bookmark
}

// UnstageVoid removes bookmark off the model stage
func (bookmark *Bookmark) UnstageVoid(stage *Stage) {
	bookmark.Unstage(stage)
}

func (bookmark *Bookmark) StageVoid(stage *Stage) {
	bookmark.Stage(stage)
}

// for satisfaction of GongStruct interface
func (bookmark *Bookmark) GetName() (res string) {
	return bookmark.Name
}

// for satisfaction of GongStruct interface
func (bookmark *Bookmark) SetName(name string) {
	bookmark.Name = name
}

// Stage puts coda to the model stage
func (coda *Coda) Stage(stage *Stage) *Coda {
	__gong__stage(stage.Codas, stage.Coda_stagedOrder, stage.Coda_orderStaged, &stage.CodaOrder, stage.Codas_mapString, coda, coda.Name)
	return coda
}

// StagePreserveOrder puts coda to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CodaOrder
// - update stage.CodaOrder accordingly
func (coda *Coda) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Codas, stage.Coda_stagedOrder, stage.Coda_orderStaged, &stage.CodaOrder, stage.Codas_mapString, coda, order, coda.Name)
}

// Unstage removes coda off the model stage
func (coda *Coda) Unstage(stage *Stage) *Coda {
	__gong__unstage(stage.Codas, stage.Codas_mapString, coda, coda.Name)
	return coda
}

// UnstageVoid removes coda off the model stage
func (coda *Coda) UnstageVoid(stage *Stage) {
	coda.Unstage(stage)
}

func (coda *Coda) StageVoid(stage *Stage) {
	coda.Stage(stage)
}

// for satisfaction of GongStruct interface
func (coda *Coda) GetName() (res string) {
	return coda.Name
}

// for satisfaction of GongStruct interface
func (coda *Coda) SetName(name string) {
	coda.Name = name
}

// Stage puts distance to the model stage
func (distance *Distance) Stage(stage *Stage) *Distance {
	__gong__stage(stage.Distances, stage.Distance_stagedOrder, stage.Distance_orderStaged, &stage.DistanceOrder, stage.Distances_mapString, distance, distance.Name)
	return distance
}

// StagePreserveOrder puts distance to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DistanceOrder
// - update stage.DistanceOrder accordingly
func (distance *Distance) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Distances, stage.Distance_stagedOrder, stage.Distance_orderStaged, &stage.DistanceOrder, stage.Distances_mapString, distance, order, distance.Name)
}

// Unstage removes distance off the model stage
func (distance *Distance) Unstage(stage *Stage) *Distance {
	__gong__unstage(stage.Distances, stage.Distances_mapString, distance, distance.Name)
	return distance
}

// UnstageVoid removes distance off the model stage
func (distance *Distance) UnstageVoid(stage *Stage) {
	distance.Unstage(stage)
}

func (distance *Distance) StageVoid(stage *Stage) {
	distance.Stage(stage)
}

// for satisfaction of GongStruct interface
func (distance *Distance) GetName() (res string) {
	return distance.Name
}

// for satisfaction of GongStruct interface
func (distance *Distance) SetName(name string) {
	distance.Name = name
}

// Stage puts dynamics to the model stage
func (dynamics *Dynamics) Stage(stage *Stage) *Dynamics {
	__gong__stage(stage.Dynamicss, stage.Dynamics_stagedOrder, stage.Dynamics_orderStaged, &stage.DynamicsOrder, stage.Dynamicss_mapString, dynamics, dynamics.Name)
	return dynamics
}

// StagePreserveOrder puts dynamics to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DynamicsOrder
// - update stage.DynamicsOrder accordingly
func (dynamics *Dynamics) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Dynamicss, stage.Dynamics_stagedOrder, stage.Dynamics_orderStaged, &stage.DynamicsOrder, stage.Dynamicss_mapString, dynamics, order, dynamics.Name)
}

// Unstage removes dynamics off the model stage
func (dynamics *Dynamics) Unstage(stage *Stage) *Dynamics {
	__gong__unstage(stage.Dynamicss, stage.Dynamicss_mapString, dynamics, dynamics.Name)
	return dynamics
}

// UnstageVoid removes dynamics off the model stage
func (dynamics *Dynamics) UnstageVoid(stage *Stage) {
	dynamics.Unstage(stage)
}

func (dynamics *Dynamics) StageVoid(stage *Stage) {
	dynamics.Stage(stage)
}

// for satisfaction of GongStruct interface
func (dynamics *Dynamics) GetName() (res string) {
	return dynamics.Name
}

// for satisfaction of GongStruct interface
func (dynamics *Dynamics) SetName(name string) {
	dynamics.Name = name
}

// Stage puts empty to the model stage
func (empty *Empty) Stage(stage *Stage) *Empty {
	__gong__stage(stage.Emptys, stage.Empty_stagedOrder, stage.Empty_orderStaged, &stage.EmptyOrder, stage.Emptys_mapString, empty, empty.Name)
	return empty
}

// StagePreserveOrder puts empty to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyOrder
// - update stage.EmptyOrder accordingly
func (empty *Empty) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Emptys, stage.Empty_stagedOrder, stage.Empty_orderStaged, &stage.EmptyOrder, stage.Emptys_mapString, empty, order, empty.Name)
}

// Unstage removes empty off the model stage
func (empty *Empty) Unstage(stage *Stage) *Empty {
	__gong__unstage(stage.Emptys, stage.Emptys_mapString, empty, empty.Name)
	return empty
}

// UnstageVoid removes empty off the model stage
func (empty *Empty) UnstageVoid(stage *Stage) {
	empty.Unstage(stage)
}

func (empty *Empty) StageVoid(stage *Stage) {
	empty.Stage(stage)
}

// for satisfaction of GongStruct interface
func (empty *Empty) GetName() (res string) {
	return empty.Name
}

// for satisfaction of GongStruct interface
func (empty *Empty) SetName(name string) {
	empty.Name = name
}

// Stage puts emptyplacement to the model stage
func (emptyplacement *EmptyPlacement) Stage(stage *Stage) *EmptyPlacement {
	__gong__stage(stage.EmptyPlacements, stage.EmptyPlacement_stagedOrder, stage.EmptyPlacement_orderStaged, &stage.EmptyPlacementOrder, stage.EmptyPlacements_mapString, emptyplacement, emptyplacement.Name)
	return emptyplacement
}

// StagePreserveOrder puts emptyplacement to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyPlacementOrder
// - update stage.EmptyPlacementOrder accordingly
func (emptyplacement *EmptyPlacement) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EmptyPlacements, stage.EmptyPlacement_stagedOrder, stage.EmptyPlacement_orderStaged, &stage.EmptyPlacementOrder, stage.EmptyPlacements_mapString, emptyplacement, order, emptyplacement.Name)
}

// Unstage removes emptyplacement off the model stage
func (emptyplacement *EmptyPlacement) Unstage(stage *Stage) *EmptyPlacement {
	__gong__unstage(stage.EmptyPlacements, stage.EmptyPlacements_mapString, emptyplacement, emptyplacement.Name)
	return emptyplacement
}

// UnstageVoid removes emptyplacement off the model stage
func (emptyplacement *EmptyPlacement) UnstageVoid(stage *Stage) {
	emptyplacement.Unstage(stage)
}

func (emptyplacement *EmptyPlacement) StageVoid(stage *Stage) {
	emptyplacement.Stage(stage)
}

// for satisfaction of GongStruct interface
func (emptyplacement *EmptyPlacement) GetName() (res string) {
	return emptyplacement.Name
}

// for satisfaction of GongStruct interface
func (emptyplacement *EmptyPlacement) SetName(name string) {
	emptyplacement.Name = name
}

// Stage puts emptyplacementsmufl to the model stage
func (emptyplacementsmufl *EmptyPlacementSmufl) Stage(stage *Stage) *EmptyPlacementSmufl {
	__gong__stage(stage.EmptyPlacementSmufls, stage.EmptyPlacementSmufl_stagedOrder, stage.EmptyPlacementSmufl_orderStaged, &stage.EmptyPlacementSmuflOrder, stage.EmptyPlacementSmufls_mapString, emptyplacementsmufl, emptyplacementsmufl.Name)
	return emptyplacementsmufl
}

// StagePreserveOrder puts emptyplacementsmufl to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyPlacementSmuflOrder
// - update stage.EmptyPlacementSmuflOrder accordingly
func (emptyplacementsmufl *EmptyPlacementSmufl) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EmptyPlacementSmufls, stage.EmptyPlacementSmufl_stagedOrder, stage.EmptyPlacementSmufl_orderStaged, &stage.EmptyPlacementSmuflOrder, stage.EmptyPlacementSmufls_mapString, emptyplacementsmufl, order, emptyplacementsmufl.Name)
}

// Unstage removes emptyplacementsmufl off the model stage
func (emptyplacementsmufl *EmptyPlacementSmufl) Unstage(stage *Stage) *EmptyPlacementSmufl {
	__gong__unstage(stage.EmptyPlacementSmufls, stage.EmptyPlacementSmufls_mapString, emptyplacementsmufl, emptyplacementsmufl.Name)
	return emptyplacementsmufl
}

// UnstageVoid removes emptyplacementsmufl off the model stage
func (emptyplacementsmufl *EmptyPlacementSmufl) UnstageVoid(stage *Stage) {
	emptyplacementsmufl.Unstage(stage)
}

func (emptyplacementsmufl *EmptyPlacementSmufl) StageVoid(stage *Stage) {
	emptyplacementsmufl.Stage(stage)
}

// for satisfaction of GongStruct interface
func (emptyplacementsmufl *EmptyPlacementSmufl) GetName() (res string) {
	return emptyplacementsmufl.Name
}

// for satisfaction of GongStruct interface
func (emptyplacementsmufl *EmptyPlacementSmufl) SetName(name string) {
	emptyplacementsmufl.Name = name
}

// Stage puts emptyprintobjectstylealign to the model stage
func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) Stage(stage *Stage) *EmptyPrintObjectStyleAlign {
	__gong__stage(stage.EmptyPrintObjectStyleAligns, stage.EmptyPrintObjectStyleAlign_stagedOrder, stage.EmptyPrintObjectStyleAlign_orderStaged, &stage.EmptyPrintObjectStyleAlignOrder, stage.EmptyPrintObjectStyleAligns_mapString, emptyprintobjectstylealign, emptyprintobjectstylealign.Name)
	return emptyprintobjectstylealign
}

// StagePreserveOrder puts emptyprintobjectstylealign to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyPrintObjectStyleAlignOrder
// - update stage.EmptyPrintObjectStyleAlignOrder accordingly
func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EmptyPrintObjectStyleAligns, stage.EmptyPrintObjectStyleAlign_stagedOrder, stage.EmptyPrintObjectStyleAlign_orderStaged, &stage.EmptyPrintObjectStyleAlignOrder, stage.EmptyPrintObjectStyleAligns_mapString, emptyprintobjectstylealign, order, emptyprintobjectstylealign.Name)
}

// Unstage removes emptyprintobjectstylealign off the model stage
func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) Unstage(stage *Stage) *EmptyPrintObjectStyleAlign {
	__gong__unstage(stage.EmptyPrintObjectStyleAligns, stage.EmptyPrintObjectStyleAligns_mapString, emptyprintobjectstylealign, emptyprintobjectstylealign.Name)
	return emptyprintobjectstylealign
}

// UnstageVoid removes emptyprintobjectstylealign off the model stage
func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) UnstageVoid(stage *Stage) {
	emptyprintobjectstylealign.Unstage(stage)
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) StageVoid(stage *Stage) {
	emptyprintobjectstylealign.Stage(stage)
}

// for satisfaction of GongStruct interface
func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GetName() (res string) {
	return emptyprintobjectstylealign.Name
}

// for satisfaction of GongStruct interface
func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) SetName(name string) {
	emptyprintobjectstylealign.Name = name
}

// Stage puts emptyprintstyle to the model stage
func (emptyprintstyle *EmptyPrintStyle) Stage(stage *Stage) *EmptyPrintStyle {
	__gong__stage(stage.EmptyPrintStyles, stage.EmptyPrintStyle_stagedOrder, stage.EmptyPrintStyle_orderStaged, &stage.EmptyPrintStyleOrder, stage.EmptyPrintStyles_mapString, emptyprintstyle, emptyprintstyle.Name)
	return emptyprintstyle
}

// StagePreserveOrder puts emptyprintstyle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyPrintStyleOrder
// - update stage.EmptyPrintStyleOrder accordingly
func (emptyprintstyle *EmptyPrintStyle) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EmptyPrintStyles, stage.EmptyPrintStyle_stagedOrder, stage.EmptyPrintStyle_orderStaged, &stage.EmptyPrintStyleOrder, stage.EmptyPrintStyles_mapString, emptyprintstyle, order, emptyprintstyle.Name)
}

// Unstage removes emptyprintstyle off the model stage
func (emptyprintstyle *EmptyPrintStyle) Unstage(stage *Stage) *EmptyPrintStyle {
	__gong__unstage(stage.EmptyPrintStyles, stage.EmptyPrintStyles_mapString, emptyprintstyle, emptyprintstyle.Name)
	return emptyprintstyle
}

// UnstageVoid removes emptyprintstyle off the model stage
func (emptyprintstyle *EmptyPrintStyle) UnstageVoid(stage *Stage) {
	emptyprintstyle.Unstage(stage)
}

func (emptyprintstyle *EmptyPrintStyle) StageVoid(stage *Stage) {
	emptyprintstyle.Stage(stage)
}

// for satisfaction of GongStruct interface
func (emptyprintstyle *EmptyPrintStyle) GetName() (res string) {
	return emptyprintstyle.Name
}

// for satisfaction of GongStruct interface
func (emptyprintstyle *EmptyPrintStyle) SetName(name string) {
	emptyprintstyle.Name = name
}

// Stage puts emptyprintstylealign to the model stage
func (emptyprintstylealign *EmptyPrintStyleAlign) Stage(stage *Stage) *EmptyPrintStyleAlign {
	__gong__stage(stage.EmptyPrintStyleAligns, stage.EmptyPrintStyleAlign_stagedOrder, stage.EmptyPrintStyleAlign_orderStaged, &stage.EmptyPrintStyleAlignOrder, stage.EmptyPrintStyleAligns_mapString, emptyprintstylealign, emptyprintstylealign.Name)
	return emptyprintstylealign
}

// StagePreserveOrder puts emptyprintstylealign to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyPrintStyleAlignOrder
// - update stage.EmptyPrintStyleAlignOrder accordingly
func (emptyprintstylealign *EmptyPrintStyleAlign) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EmptyPrintStyleAligns, stage.EmptyPrintStyleAlign_stagedOrder, stage.EmptyPrintStyleAlign_orderStaged, &stage.EmptyPrintStyleAlignOrder, stage.EmptyPrintStyleAligns_mapString, emptyprintstylealign, order, emptyprintstylealign.Name)
}

// Unstage removes emptyprintstylealign off the model stage
func (emptyprintstylealign *EmptyPrintStyleAlign) Unstage(stage *Stage) *EmptyPrintStyleAlign {
	__gong__unstage(stage.EmptyPrintStyleAligns, stage.EmptyPrintStyleAligns_mapString, emptyprintstylealign, emptyprintstylealign.Name)
	return emptyprintstylealign
}

// UnstageVoid removes emptyprintstylealign off the model stage
func (emptyprintstylealign *EmptyPrintStyleAlign) UnstageVoid(stage *Stage) {
	emptyprintstylealign.Unstage(stage)
}

func (emptyprintstylealign *EmptyPrintStyleAlign) StageVoid(stage *Stage) {
	emptyprintstylealign.Stage(stage)
}

// for satisfaction of GongStruct interface
func (emptyprintstylealign *EmptyPrintStyleAlign) GetName() (res string) {
	return emptyprintstylealign.Name
}

// for satisfaction of GongStruct interface
func (emptyprintstylealign *EmptyPrintStyleAlign) SetName(name string) {
	emptyprintstylealign.Name = name
}

// Stage puts emptyprintstylealignid to the model stage
func (emptyprintstylealignid *EmptyPrintStyleAlignId) Stage(stage *Stage) *EmptyPrintStyleAlignId {
	__gong__stage(stage.EmptyPrintStyleAlignIds, stage.EmptyPrintStyleAlignId_stagedOrder, stage.EmptyPrintStyleAlignId_orderStaged, &stage.EmptyPrintStyleAlignIdOrder, stage.EmptyPrintStyleAlignIds_mapString, emptyprintstylealignid, emptyprintstylealignid.Name)
	return emptyprintstylealignid
}

// StagePreserveOrder puts emptyprintstylealignid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyPrintStyleAlignIdOrder
// - update stage.EmptyPrintStyleAlignIdOrder accordingly
func (emptyprintstylealignid *EmptyPrintStyleAlignId) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EmptyPrintStyleAlignIds, stage.EmptyPrintStyleAlignId_stagedOrder, stage.EmptyPrintStyleAlignId_orderStaged, &stage.EmptyPrintStyleAlignIdOrder, stage.EmptyPrintStyleAlignIds_mapString, emptyprintstylealignid, order, emptyprintstylealignid.Name)
}

// Unstage removes emptyprintstylealignid off the model stage
func (emptyprintstylealignid *EmptyPrintStyleAlignId) Unstage(stage *Stage) *EmptyPrintStyleAlignId {
	__gong__unstage(stage.EmptyPrintStyleAlignIds, stage.EmptyPrintStyleAlignIds_mapString, emptyprintstylealignid, emptyprintstylealignid.Name)
	return emptyprintstylealignid
}

// UnstageVoid removes emptyprintstylealignid off the model stage
func (emptyprintstylealignid *EmptyPrintStyleAlignId) UnstageVoid(stage *Stage) {
	emptyprintstylealignid.Unstage(stage)
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) StageVoid(stage *Stage) {
	emptyprintstylealignid.Stage(stage)
}

// for satisfaction of GongStruct interface
func (emptyprintstylealignid *EmptyPrintStyleAlignId) GetName() (res string) {
	return emptyprintstylealignid.Name
}

// for satisfaction of GongStruct interface
func (emptyprintstylealignid *EmptyPrintStyleAlignId) SetName(name string) {
	emptyprintstylealignid.Name = name
}

// Stage puts emptytrillsound to the model stage
func (emptytrillsound *EmptyTrillSound) Stage(stage *Stage) *EmptyTrillSound {
	__gong__stage(stage.EmptyTrillSounds, stage.EmptyTrillSound_stagedOrder, stage.EmptyTrillSound_orderStaged, &stage.EmptyTrillSoundOrder, stage.EmptyTrillSounds_mapString, emptytrillsound, emptytrillsound.Name)
	return emptytrillsound
}

// StagePreserveOrder puts emptytrillsound to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EmptyTrillSoundOrder
// - update stage.EmptyTrillSoundOrder accordingly
func (emptytrillsound *EmptyTrillSound) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.EmptyTrillSounds, stage.EmptyTrillSound_stagedOrder, stage.EmptyTrillSound_orderStaged, &stage.EmptyTrillSoundOrder, stage.EmptyTrillSounds_mapString, emptytrillsound, order, emptytrillsound.Name)
}

// Unstage removes emptytrillsound off the model stage
func (emptytrillsound *EmptyTrillSound) Unstage(stage *Stage) *EmptyTrillSound {
	__gong__unstage(stage.EmptyTrillSounds, stage.EmptyTrillSounds_mapString, emptytrillsound, emptytrillsound.Name)
	return emptytrillsound
}

// UnstageVoid removes emptytrillsound off the model stage
func (emptytrillsound *EmptyTrillSound) UnstageVoid(stage *Stage) {
	emptytrillsound.Unstage(stage)
}

func (emptytrillsound *EmptyTrillSound) StageVoid(stage *Stage) {
	emptytrillsound.Stage(stage)
}

// for satisfaction of GongStruct interface
func (emptytrillsound *EmptyTrillSound) GetName() (res string) {
	return emptytrillsound.Name
}

// for satisfaction of GongStruct interface
func (emptytrillsound *EmptyTrillSound) SetName(name string) {
	emptytrillsound.Name = name
}

// Stage puts encoding to the model stage
func (encoding *Encoding) Stage(stage *Stage) *Encoding {
	__gong__stage(stage.Encodings, stage.Encoding_stagedOrder, stage.Encoding_orderStaged, &stage.EncodingOrder, stage.Encodings_mapString, encoding, encoding.Name)
	return encoding
}

// StagePreserveOrder puts encoding to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EncodingOrder
// - update stage.EncodingOrder accordingly
func (encoding *Encoding) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Encodings, stage.Encoding_stagedOrder, stage.Encoding_orderStaged, &stage.EncodingOrder, stage.Encodings_mapString, encoding, order, encoding.Name)
}

// Unstage removes encoding off the model stage
func (encoding *Encoding) Unstage(stage *Stage) *Encoding {
	__gong__unstage(stage.Encodings, stage.Encodings_mapString, encoding, encoding.Name)
	return encoding
}

// UnstageVoid removes encoding off the model stage
func (encoding *Encoding) UnstageVoid(stage *Stage) {
	encoding.Unstage(stage)
}

func (encoding *Encoding) StageVoid(stage *Stage) {
	encoding.Stage(stage)
}

// for satisfaction of GongStruct interface
func (encoding *Encoding) GetName() (res string) {
	return encoding.Name
}

// for satisfaction of GongStruct interface
func (encoding *Encoding) SetName(name string) {
	encoding.Name = name
}

// Stage puts fermata to the model stage
func (fermata *Fermata) Stage(stage *Stage) *Fermata {
	__gong__stage(stage.Fermatas, stage.Fermata_stagedOrder, stage.Fermata_orderStaged, &stage.FermataOrder, stage.Fermatas_mapString, fermata, fermata.Name)
	return fermata
}

// StagePreserveOrder puts fermata to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FermataOrder
// - update stage.FermataOrder accordingly
func (fermata *Fermata) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Fermatas, stage.Fermata_stagedOrder, stage.Fermata_orderStaged, &stage.FermataOrder, stage.Fermatas_mapString, fermata, order, fermata.Name)
}

// Unstage removes fermata off the model stage
func (fermata *Fermata) Unstage(stage *Stage) *Fermata {
	__gong__unstage(stage.Fermatas, stage.Fermatas_mapString, fermata, fermata.Name)
	return fermata
}

// UnstageVoid removes fermata off the model stage
func (fermata *Fermata) UnstageVoid(stage *Stage) {
	fermata.Unstage(stage)
}

func (fermata *Fermata) StageVoid(stage *Stage) {
	fermata.Stage(stage)
}

// for satisfaction of GongStruct interface
func (fermata *Fermata) GetName() (res string) {
	return fermata.Name
}

// for satisfaction of GongStruct interface
func (fermata *Fermata) SetName(name string) {
	fermata.Name = name
}

// Stage puts fingering to the model stage
func (fingering *Fingering) Stage(stage *Stage) *Fingering {
	__gong__stage(stage.Fingerings, stage.Fingering_stagedOrder, stage.Fingering_orderStaged, &stage.FingeringOrder, stage.Fingerings_mapString, fingering, fingering.Name)
	return fingering
}

// StagePreserveOrder puts fingering to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FingeringOrder
// - update stage.FingeringOrder accordingly
func (fingering *Fingering) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Fingerings, stage.Fingering_stagedOrder, stage.Fingering_orderStaged, &stage.FingeringOrder, stage.Fingerings_mapString, fingering, order, fingering.Name)
}

// Unstage removes fingering off the model stage
func (fingering *Fingering) Unstage(stage *Stage) *Fingering {
	__gong__unstage(stage.Fingerings, stage.Fingerings_mapString, fingering, fingering.Name)
	return fingering
}

// UnstageVoid removes fingering off the model stage
func (fingering *Fingering) UnstageVoid(stage *Stage) {
	fingering.Unstage(stage)
}

func (fingering *Fingering) StageVoid(stage *Stage) {
	fingering.Stage(stage)
}

// for satisfaction of GongStruct interface
func (fingering *Fingering) GetName() (res string) {
	return fingering.Name
}

// for satisfaction of GongStruct interface
func (fingering *Fingering) SetName(name string) {
	fingering.Name = name
}

// Stage puts formattedsymbol to the model stage
func (formattedsymbol *FormattedSymbol) Stage(stage *Stage) *FormattedSymbol {
	__gong__stage(stage.FormattedSymbols, stage.FormattedSymbol_stagedOrder, stage.FormattedSymbol_orderStaged, &stage.FormattedSymbolOrder, stage.FormattedSymbols_mapString, formattedsymbol, formattedsymbol.Name)
	return formattedsymbol
}

// StagePreserveOrder puts formattedsymbol to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormattedSymbolOrder
// - update stage.FormattedSymbolOrder accordingly
func (formattedsymbol *FormattedSymbol) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormattedSymbols, stage.FormattedSymbol_stagedOrder, stage.FormattedSymbol_orderStaged, &stage.FormattedSymbolOrder, stage.FormattedSymbols_mapString, formattedsymbol, order, formattedsymbol.Name)
}

// Unstage removes formattedsymbol off the model stage
func (formattedsymbol *FormattedSymbol) Unstage(stage *Stage) *FormattedSymbol {
	__gong__unstage(stage.FormattedSymbols, stage.FormattedSymbols_mapString, formattedsymbol, formattedsymbol.Name)
	return formattedsymbol
}

// UnstageVoid removes formattedsymbol off the model stage
func (formattedsymbol *FormattedSymbol) UnstageVoid(stage *Stage) {
	formattedsymbol.Unstage(stage)
}

func (formattedsymbol *FormattedSymbol) StageVoid(stage *Stage) {
	formattedsymbol.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formattedsymbol *FormattedSymbol) GetName() (res string) {
	return formattedsymbol.Name
}

// for satisfaction of GongStruct interface
func (formattedsymbol *FormattedSymbol) SetName(name string) {
	formattedsymbol.Name = name
}

// Stage puts formattedsymbolid to the model stage
func (formattedsymbolid *FormattedSymbolId) Stage(stage *Stage) *FormattedSymbolId {
	__gong__stage(stage.FormattedSymbolIds, stage.FormattedSymbolId_stagedOrder, stage.FormattedSymbolId_orderStaged, &stage.FormattedSymbolIdOrder, stage.FormattedSymbolIds_mapString, formattedsymbolid, formattedsymbolid.Name)
	return formattedsymbolid
}

// StagePreserveOrder puts formattedsymbolid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormattedSymbolIdOrder
// - update stage.FormattedSymbolIdOrder accordingly
func (formattedsymbolid *FormattedSymbolId) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormattedSymbolIds, stage.FormattedSymbolId_stagedOrder, stage.FormattedSymbolId_orderStaged, &stage.FormattedSymbolIdOrder, stage.FormattedSymbolIds_mapString, formattedsymbolid, order, formattedsymbolid.Name)
}

// Unstage removes formattedsymbolid off the model stage
func (formattedsymbolid *FormattedSymbolId) Unstage(stage *Stage) *FormattedSymbolId {
	__gong__unstage(stage.FormattedSymbolIds, stage.FormattedSymbolIds_mapString, formattedsymbolid, formattedsymbolid.Name)
	return formattedsymbolid
}

// UnstageVoid removes formattedsymbolid off the model stage
func (formattedsymbolid *FormattedSymbolId) UnstageVoid(stage *Stage) {
	formattedsymbolid.Unstage(stage)
}

func (formattedsymbolid *FormattedSymbolId) StageVoid(stage *Stage) {
	formattedsymbolid.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formattedsymbolid *FormattedSymbolId) GetName() (res string) {
	return formattedsymbolid.Name
}

// for satisfaction of GongStruct interface
func (formattedsymbolid *FormattedSymbolId) SetName(name string) {
	formattedsymbolid.Name = name
}

// Stage puts formattedtext to the model stage
func (formattedtext *FormattedText) Stage(stage *Stage) *FormattedText {
	__gong__stage(stage.FormattedTexts, stage.FormattedText_stagedOrder, stage.FormattedText_orderStaged, &stage.FormattedTextOrder, stage.FormattedTexts_mapString, formattedtext, formattedtext.Name)
	return formattedtext
}

// StagePreserveOrder puts formattedtext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormattedTextOrder
// - update stage.FormattedTextOrder accordingly
func (formattedtext *FormattedText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormattedTexts, stage.FormattedText_stagedOrder, stage.FormattedText_orderStaged, &stage.FormattedTextOrder, stage.FormattedTexts_mapString, formattedtext, order, formattedtext.Name)
}

// Unstage removes formattedtext off the model stage
func (formattedtext *FormattedText) Unstage(stage *Stage) *FormattedText {
	__gong__unstage(stage.FormattedTexts, stage.FormattedTexts_mapString, formattedtext, formattedtext.Name)
	return formattedtext
}

// UnstageVoid removes formattedtext off the model stage
func (formattedtext *FormattedText) UnstageVoid(stage *Stage) {
	formattedtext.Unstage(stage)
}

func (formattedtext *FormattedText) StageVoid(stage *Stage) {
	formattedtext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formattedtext *FormattedText) GetName() (res string) {
	return formattedtext.Name
}

// for satisfaction of GongStruct interface
func (formattedtext *FormattedText) SetName(name string) {
	formattedtext.Name = name
}

// Stage puts formattedtextid to the model stage
func (formattedtextid *FormattedTextId) Stage(stage *Stage) *FormattedTextId {
	__gong__stage(stage.FormattedTextIds, stage.FormattedTextId_stagedOrder, stage.FormattedTextId_orderStaged, &stage.FormattedTextIdOrder, stage.FormattedTextIds_mapString, formattedtextid, formattedtextid.Name)
	return formattedtextid
}

// StagePreserveOrder puts formattedtextid to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FormattedTextIdOrder
// - update stage.FormattedTextIdOrder accordingly
func (formattedtextid *FormattedTextId) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.FormattedTextIds, stage.FormattedTextId_stagedOrder, stage.FormattedTextId_orderStaged, &stage.FormattedTextIdOrder, stage.FormattedTextIds_mapString, formattedtextid, order, formattedtextid.Name)
}

// Unstage removes formattedtextid off the model stage
func (formattedtextid *FormattedTextId) Unstage(stage *Stage) *FormattedTextId {
	__gong__unstage(stage.FormattedTextIds, stage.FormattedTextIds_mapString, formattedtextid, formattedtextid.Name)
	return formattedtextid
}

// UnstageVoid removes formattedtextid off the model stage
func (formattedtextid *FormattedTextId) UnstageVoid(stage *Stage) {
	formattedtextid.Unstage(stage)
}

func (formattedtextid *FormattedTextId) StageVoid(stage *Stage) {
	formattedtextid.Stage(stage)
}

// for satisfaction of GongStruct interface
func (formattedtextid *FormattedTextId) GetName() (res string) {
	return formattedtextid.Name
}

// for satisfaction of GongStruct interface
func (formattedtextid *FormattedTextId) SetName(name string) {
	formattedtextid.Name = name
}

// Stage puts fret to the model stage
func (fret *Fret) Stage(stage *Stage) *Fret {
	__gong__stage(stage.Frets, stage.Fret_stagedOrder, stage.Fret_orderStaged, &stage.FretOrder, stage.Frets_mapString, fret, fret.Name)
	return fret
}

// StagePreserveOrder puts fret to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FretOrder
// - update stage.FretOrder accordingly
func (fret *Fret) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Frets, stage.Fret_stagedOrder, stage.Fret_orderStaged, &stage.FretOrder, stage.Frets_mapString, fret, order, fret.Name)
}

// Unstage removes fret off the model stage
func (fret *Fret) Unstage(stage *Stage) *Fret {
	__gong__unstage(stage.Frets, stage.Frets_mapString, fret, fret.Name)
	return fret
}

// UnstageVoid removes fret off the model stage
func (fret *Fret) UnstageVoid(stage *Stage) {
	fret.Unstage(stage)
}

func (fret *Fret) StageVoid(stage *Stage) {
	fret.Stage(stage)
}

// for satisfaction of GongStruct interface
func (fret *Fret) GetName() (res string) {
	return fret.Name
}

// for satisfaction of GongStruct interface
func (fret *Fret) SetName(name string) {
	fret.Name = name
}

// Stage puts glyph to the model stage
func (glyph *Glyph) Stage(stage *Stage) *Glyph {
	__gong__stage(stage.Glyphs, stage.Glyph_stagedOrder, stage.Glyph_orderStaged, &stage.GlyphOrder, stage.Glyphs_mapString, glyph, glyph.Name)
	return glyph
}

// StagePreserveOrder puts glyph to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GlyphOrder
// - update stage.GlyphOrder accordingly
func (glyph *Glyph) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Glyphs, stage.Glyph_stagedOrder, stage.Glyph_orderStaged, &stage.GlyphOrder, stage.Glyphs_mapString, glyph, order, glyph.Name)
}

// Unstage removes glyph off the model stage
func (glyph *Glyph) Unstage(stage *Stage) *Glyph {
	__gong__unstage(stage.Glyphs, stage.Glyphs_mapString, glyph, glyph.Name)
	return glyph
}

// UnstageVoid removes glyph off the model stage
func (glyph *Glyph) UnstageVoid(stage *Stage) {
	glyph.Unstage(stage)
}

func (glyph *Glyph) StageVoid(stage *Stage) {
	glyph.Stage(stage)
}

// for satisfaction of GongStruct interface
func (glyph *Glyph) GetName() (res string) {
	return glyph.Name
}

// for satisfaction of GongStruct interface
func (glyph *Glyph) SetName(name string) {
	glyph.Name = name
}

// Stage puts horizontalturn to the model stage
func (horizontalturn *HorizontalTurn) Stage(stage *Stage) *HorizontalTurn {
	__gong__stage(stage.HorizontalTurns, stage.HorizontalTurn_stagedOrder, stage.HorizontalTurn_orderStaged, &stage.HorizontalTurnOrder, stage.HorizontalTurns_mapString, horizontalturn, horizontalturn.Name)
	return horizontalturn
}

// StagePreserveOrder puts horizontalturn to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.HorizontalTurnOrder
// - update stage.HorizontalTurnOrder accordingly
func (horizontalturn *HorizontalTurn) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.HorizontalTurns, stage.HorizontalTurn_stagedOrder, stage.HorizontalTurn_orderStaged, &stage.HorizontalTurnOrder, stage.HorizontalTurns_mapString, horizontalturn, order, horizontalturn.Name)
}

// Unstage removes horizontalturn off the model stage
func (horizontalturn *HorizontalTurn) Unstage(stage *Stage) *HorizontalTurn {
	__gong__unstage(stage.HorizontalTurns, stage.HorizontalTurns_mapString, horizontalturn, horizontalturn.Name)
	return horizontalturn
}

// UnstageVoid removes horizontalturn off the model stage
func (horizontalturn *HorizontalTurn) UnstageVoid(stage *Stage) {
	horizontalturn.Unstage(stage)
}

func (horizontalturn *HorizontalTurn) StageVoid(stage *Stage) {
	horizontalturn.Stage(stage)
}

// for satisfaction of GongStruct interface
func (horizontalturn *HorizontalTurn) GetName() (res string) {
	return horizontalturn.Name
}

// for satisfaction of GongStruct interface
func (horizontalturn *HorizontalTurn) SetName(name string) {
	horizontalturn.Name = name
}

// Stage puts identification to the model stage
func (identification *Identification) Stage(stage *Stage) *Identification {
	__gong__stage(stage.Identifications, stage.Identification_stagedOrder, stage.Identification_orderStaged, &stage.IdentificationOrder, stage.Identifications_mapString, identification, identification.Name)
	return identification
}

// StagePreserveOrder puts identification to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.IdentificationOrder
// - update stage.IdentificationOrder accordingly
func (identification *Identification) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Identifications, stage.Identification_stagedOrder, stage.Identification_orderStaged, &stage.IdentificationOrder, stage.Identifications_mapString, identification, order, identification.Name)
}

// Unstage removes identification off the model stage
func (identification *Identification) Unstage(stage *Stage) *Identification {
	__gong__unstage(stage.Identifications, stage.Identifications_mapString, identification, identification.Name)
	return identification
}

// UnstageVoid removes identification off the model stage
func (identification *Identification) UnstageVoid(stage *Stage) {
	identification.Unstage(stage)
}

func (identification *Identification) StageVoid(stage *Stage) {
	identification.Stage(stage)
}

// for satisfaction of GongStruct interface
func (identification *Identification) GetName() (res string) {
	return identification.Name
}

// for satisfaction of GongStruct interface
func (identification *Identification) SetName(name string) {
	identification.Name = name
}

// Stage puts level to the model stage
func (level *Level) Stage(stage *Stage) *Level {
	__gong__stage(stage.Levels, stage.Level_stagedOrder, stage.Level_orderStaged, &stage.LevelOrder, stage.Levels_mapString, level, level.Name)
	return level
}

// StagePreserveOrder puts level to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LevelOrder
// - update stage.LevelOrder accordingly
func (level *Level) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Levels, stage.Level_stagedOrder, stage.Level_orderStaged, &stage.LevelOrder, stage.Levels_mapString, level, order, level.Name)
}

// Unstage removes level off the model stage
func (level *Level) Unstage(stage *Stage) *Level {
	__gong__unstage(stage.Levels, stage.Levels_mapString, level, level.Name)
	return level
}

// UnstageVoid removes level off the model stage
func (level *Level) UnstageVoid(stage *Stage) {
	level.Unstage(stage)
}

func (level *Level) StageVoid(stage *Stage) {
	level.Stage(stage)
}

// for satisfaction of GongStruct interface
func (level *Level) GetName() (res string) {
	return level.Name
}

// for satisfaction of GongStruct interface
func (level *Level) SetName(name string) {
	level.Name = name
}

// Stage puts linewidth to the model stage
func (linewidth *LineWidth) Stage(stage *Stage) *LineWidth {
	__gong__stage(stage.LineWidths, stage.LineWidth_stagedOrder, stage.LineWidth_orderStaged, &stage.LineWidthOrder, stage.LineWidths_mapString, linewidth, linewidth.Name)
	return linewidth
}

// StagePreserveOrder puts linewidth to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LineWidthOrder
// - update stage.LineWidthOrder accordingly
func (linewidth *LineWidth) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.LineWidths, stage.LineWidth_stagedOrder, stage.LineWidth_orderStaged, &stage.LineWidthOrder, stage.LineWidths_mapString, linewidth, order, linewidth.Name)
}

// Unstage removes linewidth off the model stage
func (linewidth *LineWidth) Unstage(stage *Stage) *LineWidth {
	__gong__unstage(stage.LineWidths, stage.LineWidths_mapString, linewidth, linewidth.Name)
	return linewidth
}

// UnstageVoid removes linewidth off the model stage
func (linewidth *LineWidth) UnstageVoid(stage *Stage) {
	linewidth.Unstage(stage)
}

func (linewidth *LineWidth) StageVoid(stage *Stage) {
	linewidth.Stage(stage)
}

// for satisfaction of GongStruct interface
func (linewidth *LineWidth) GetName() (res string) {
	return linewidth.Name
}

// for satisfaction of GongStruct interface
func (linewidth *LineWidth) SetName(name string) {
	linewidth.Name = name
}

// Stage puts link to the model stage
func (link *Link) Stage(stage *Stage) *Link {
	__gong__stage(stage.Links, stage.Link_stagedOrder, stage.Link_orderStaged, &stage.LinkOrder, stage.Links_mapString, link, link.Name)
	return link
}

// StagePreserveOrder puts link to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LinkOrder
// - update stage.LinkOrder accordingly
func (link *Link) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Links, stage.Link_stagedOrder, stage.Link_orderStaged, &stage.LinkOrder, stage.Links_mapString, link, order, link.Name)
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *Stage) *Link {
	__gong__unstage(stage.Links, stage.Links_mapString, link, link.Name)
	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *Stage) {
	link.Unstage(stage)
}

func (link *Link) StageVoid(stage *Stage) {
	link.Stage(stage)
}

// for satisfaction of GongStruct interface
func (link *Link) GetName() (res string) {
	return link.Name
}

// for satisfaction of GongStruct interface
func (link *Link) SetName(name string) {
	link.Name = name
}

// Stage puts measurelayout to the model stage
func (measurelayout *MeasureLayout) Stage(stage *Stage) *MeasureLayout {
	__gong__stage(stage.MeasureLayouts, stage.MeasureLayout_stagedOrder, stage.MeasureLayout_orderStaged, &stage.MeasureLayoutOrder, stage.MeasureLayouts_mapString, measurelayout, measurelayout.Name)
	return measurelayout
}

// StagePreserveOrder puts measurelayout to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MeasureLayoutOrder
// - update stage.MeasureLayoutOrder accordingly
func (measurelayout *MeasureLayout) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MeasureLayouts, stage.MeasureLayout_stagedOrder, stage.MeasureLayout_orderStaged, &stage.MeasureLayoutOrder, stage.MeasureLayouts_mapString, measurelayout, order, measurelayout.Name)
}

// Unstage removes measurelayout off the model stage
func (measurelayout *MeasureLayout) Unstage(stage *Stage) *MeasureLayout {
	__gong__unstage(stage.MeasureLayouts, stage.MeasureLayouts_mapString, measurelayout, measurelayout.Name)
	return measurelayout
}

// UnstageVoid removes measurelayout off the model stage
func (measurelayout *MeasureLayout) UnstageVoid(stage *Stage) {
	measurelayout.Unstage(stage)
}

func (measurelayout *MeasureLayout) StageVoid(stage *Stage) {
	measurelayout.Stage(stage)
}

// for satisfaction of GongStruct interface
func (measurelayout *MeasureLayout) GetName() (res string) {
	return measurelayout.Name
}

// for satisfaction of GongStruct interface
func (measurelayout *MeasureLayout) SetName(name string) {
	measurelayout.Name = name
}

// Stage puts mididevice to the model stage
func (mididevice *MidiDevice) Stage(stage *Stage) *MidiDevice {
	__gong__stage(stage.MidiDevices, stage.MidiDevice_stagedOrder, stage.MidiDevice_orderStaged, &stage.MidiDeviceOrder, stage.MidiDevices_mapString, mididevice, mididevice.Name)
	return mididevice
}

// StagePreserveOrder puts mididevice to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MidiDeviceOrder
// - update stage.MidiDeviceOrder accordingly
func (mididevice *MidiDevice) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MidiDevices, stage.MidiDevice_stagedOrder, stage.MidiDevice_orderStaged, &stage.MidiDeviceOrder, stage.MidiDevices_mapString, mididevice, order, mididevice.Name)
}

// Unstage removes mididevice off the model stage
func (mididevice *MidiDevice) Unstage(stage *Stage) *MidiDevice {
	__gong__unstage(stage.MidiDevices, stage.MidiDevices_mapString, mididevice, mididevice.Name)
	return mididevice
}

// UnstageVoid removes mididevice off the model stage
func (mididevice *MidiDevice) UnstageVoid(stage *Stage) {
	mididevice.Unstage(stage)
}

func (mididevice *MidiDevice) StageVoid(stage *Stage) {
	mididevice.Stage(stage)
}

// for satisfaction of GongStruct interface
func (mididevice *MidiDevice) GetName() (res string) {
	return mididevice.Name
}

// for satisfaction of GongStruct interface
func (mididevice *MidiDevice) SetName(name string) {
	mididevice.Name = name
}

// Stage puts midiinstrument to the model stage
func (midiinstrument *MidiInstrument) Stage(stage *Stage) *MidiInstrument {
	__gong__stage(stage.MidiInstruments, stage.MidiInstrument_stagedOrder, stage.MidiInstrument_orderStaged, &stage.MidiInstrumentOrder, stage.MidiInstruments_mapString, midiinstrument, midiinstrument.Name)
	return midiinstrument
}

// StagePreserveOrder puts midiinstrument to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MidiInstrumentOrder
// - update stage.MidiInstrumentOrder accordingly
func (midiinstrument *MidiInstrument) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MidiInstruments, stage.MidiInstrument_stagedOrder, stage.MidiInstrument_orderStaged, &stage.MidiInstrumentOrder, stage.MidiInstruments_mapString, midiinstrument, order, midiinstrument.Name)
}

// Unstage removes midiinstrument off the model stage
func (midiinstrument *MidiInstrument) Unstage(stage *Stage) *MidiInstrument {
	__gong__unstage(stage.MidiInstruments, stage.MidiInstruments_mapString, midiinstrument, midiinstrument.Name)
	return midiinstrument
}

// UnstageVoid removes midiinstrument off the model stage
func (midiinstrument *MidiInstrument) UnstageVoid(stage *Stage) {
	midiinstrument.Unstage(stage)
}

func (midiinstrument *MidiInstrument) StageVoid(stage *Stage) {
	midiinstrument.Stage(stage)
}

// for satisfaction of GongStruct interface
func (midiinstrument *MidiInstrument) GetName() (res string) {
	return midiinstrument.Name
}

// for satisfaction of GongStruct interface
func (midiinstrument *MidiInstrument) SetName(name string) {
	midiinstrument.Name = name
}

// Stage puts miscellaneous to the model stage
func (miscellaneous *Miscellaneous) Stage(stage *Stage) *Miscellaneous {
	__gong__stage(stage.Miscellaneouss, stage.Miscellaneous_stagedOrder, stage.Miscellaneous_orderStaged, &stage.MiscellaneousOrder, stage.Miscellaneouss_mapString, miscellaneous, miscellaneous.Name)
	return miscellaneous
}

// StagePreserveOrder puts miscellaneous to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MiscellaneousOrder
// - update stage.MiscellaneousOrder accordingly
func (miscellaneous *Miscellaneous) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Miscellaneouss, stage.Miscellaneous_stagedOrder, stage.Miscellaneous_orderStaged, &stage.MiscellaneousOrder, stage.Miscellaneouss_mapString, miscellaneous, order, miscellaneous.Name)
}

// Unstage removes miscellaneous off the model stage
func (miscellaneous *Miscellaneous) Unstage(stage *Stage) *Miscellaneous {
	__gong__unstage(stage.Miscellaneouss, stage.Miscellaneouss_mapString, miscellaneous, miscellaneous.Name)
	return miscellaneous
}

// UnstageVoid removes miscellaneous off the model stage
func (miscellaneous *Miscellaneous) UnstageVoid(stage *Stage) {
	miscellaneous.Unstage(stage)
}

func (miscellaneous *Miscellaneous) StageVoid(stage *Stage) {
	miscellaneous.Stage(stage)
}

// for satisfaction of GongStruct interface
func (miscellaneous *Miscellaneous) GetName() (res string) {
	return miscellaneous.Name
}

// for satisfaction of GongStruct interface
func (miscellaneous *Miscellaneous) SetName(name string) {
	miscellaneous.Name = name
}

// Stage puts miscellaneousfield to the model stage
func (miscellaneousfield *MiscellaneousField) Stage(stage *Stage) *MiscellaneousField {
	__gong__stage(stage.MiscellaneousFields, stage.MiscellaneousField_stagedOrder, stage.MiscellaneousField_orderStaged, &stage.MiscellaneousFieldOrder, stage.MiscellaneousFields_mapString, miscellaneousfield, miscellaneousfield.Name)
	return miscellaneousfield
}

// StagePreserveOrder puts miscellaneousfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MiscellaneousFieldOrder
// - update stage.MiscellaneousFieldOrder accordingly
func (miscellaneousfield *MiscellaneousField) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MiscellaneousFields, stage.MiscellaneousField_stagedOrder, stage.MiscellaneousField_orderStaged, &stage.MiscellaneousFieldOrder, stage.MiscellaneousFields_mapString, miscellaneousfield, order, miscellaneousfield.Name)
}

// Unstage removes miscellaneousfield off the model stage
func (miscellaneousfield *MiscellaneousField) Unstage(stage *Stage) *MiscellaneousField {
	__gong__unstage(stage.MiscellaneousFields, stage.MiscellaneousFields_mapString, miscellaneousfield, miscellaneousfield.Name)
	return miscellaneousfield
}

// UnstageVoid removes miscellaneousfield off the model stage
func (miscellaneousfield *MiscellaneousField) UnstageVoid(stage *Stage) {
	miscellaneousfield.Unstage(stage)
}

func (miscellaneousfield *MiscellaneousField) StageVoid(stage *Stage) {
	miscellaneousfield.Stage(stage)
}

// for satisfaction of GongStruct interface
func (miscellaneousfield *MiscellaneousField) GetName() (res string) {
	return miscellaneousfield.Name
}

// for satisfaction of GongStruct interface
func (miscellaneousfield *MiscellaneousField) SetName(name string) {
	miscellaneousfield.Name = name
}

// Stage puts namedisplay to the model stage
func (namedisplay *NameDisplay) Stage(stage *Stage) *NameDisplay {
	__gong__stage(stage.NameDisplays, stage.NameDisplay_stagedOrder, stage.NameDisplay_orderStaged, &stage.NameDisplayOrder, stage.NameDisplays_mapString, namedisplay, namedisplay.Name)
	return namedisplay
}

// StagePreserveOrder puts namedisplay to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NameDisplayOrder
// - update stage.NameDisplayOrder accordingly
func (namedisplay *NameDisplay) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NameDisplays, stage.NameDisplay_stagedOrder, stage.NameDisplay_orderStaged, &stage.NameDisplayOrder, stage.NameDisplays_mapString, namedisplay, order, namedisplay.Name)
}

// Unstage removes namedisplay off the model stage
func (namedisplay *NameDisplay) Unstage(stage *Stage) *NameDisplay {
	__gong__unstage(stage.NameDisplays, stage.NameDisplays_mapString, namedisplay, namedisplay.Name)
	return namedisplay
}

// UnstageVoid removes namedisplay off the model stage
func (namedisplay *NameDisplay) UnstageVoid(stage *Stage) {
	namedisplay.Unstage(stage)
}

func (namedisplay *NameDisplay) StageVoid(stage *Stage) {
	namedisplay.Stage(stage)
}

// for satisfaction of GongStruct interface
func (namedisplay *NameDisplay) GetName() (res string) {
	return namedisplay.Name
}

// for satisfaction of GongStruct interface
func (namedisplay *NameDisplay) SetName(name string) {
	namedisplay.Name = name
}

// Stage puts notesize to the model stage
func (notesize *NoteSize) Stage(stage *Stage) *NoteSize {
	__gong__stage(stage.NoteSizes, stage.NoteSize_stagedOrder, stage.NoteSize_orderStaged, &stage.NoteSizeOrder, stage.NoteSizes_mapString, notesize, notesize.Name)
	return notesize
}

// StagePreserveOrder puts notesize to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteSizeOrder
// - update stage.NoteSizeOrder accordingly
func (notesize *NoteSize) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteSizes, stage.NoteSize_stagedOrder, stage.NoteSize_orderStaged, &stage.NoteSizeOrder, stage.NoteSizes_mapString, notesize, order, notesize.Name)
}

// Unstage removes notesize off the model stage
func (notesize *NoteSize) Unstage(stage *Stage) *NoteSize {
	__gong__unstage(stage.NoteSizes, stage.NoteSizes_mapString, notesize, notesize.Name)
	return notesize
}

// UnstageVoid removes notesize off the model stage
func (notesize *NoteSize) UnstageVoid(stage *Stage) {
	notesize.Unstage(stage)
}

func (notesize *NoteSize) StageVoid(stage *Stage) {
	notesize.Stage(stage)
}

// for satisfaction of GongStruct interface
func (notesize *NoteSize) GetName() (res string) {
	return notesize.Name
}

// for satisfaction of GongStruct interface
func (notesize *NoteSize) SetName(name string) {
	notesize.Name = name
}

// Stage puts otherappearance to the model stage
func (otherappearance *OtherAppearance) Stage(stage *Stage) *OtherAppearance {
	__gong__stage(stage.OtherAppearances, stage.OtherAppearance_stagedOrder, stage.OtherAppearance_orderStaged, &stage.OtherAppearanceOrder, stage.OtherAppearances_mapString, otherappearance, otherappearance.Name)
	return otherappearance
}

// StagePreserveOrder puts otherappearance to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OtherAppearanceOrder
// - update stage.OtherAppearanceOrder accordingly
func (otherappearance *OtherAppearance) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OtherAppearances, stage.OtherAppearance_stagedOrder, stage.OtherAppearance_orderStaged, &stage.OtherAppearanceOrder, stage.OtherAppearances_mapString, otherappearance, order, otherappearance.Name)
}

// Unstage removes otherappearance off the model stage
func (otherappearance *OtherAppearance) Unstage(stage *Stage) *OtherAppearance {
	__gong__unstage(stage.OtherAppearances, stage.OtherAppearances_mapString, otherappearance, otherappearance.Name)
	return otherappearance
}

// UnstageVoid removes otherappearance off the model stage
func (otherappearance *OtherAppearance) UnstageVoid(stage *Stage) {
	otherappearance.Unstage(stage)
}

func (otherappearance *OtherAppearance) StageVoid(stage *Stage) {
	otherappearance.Stage(stage)
}

// for satisfaction of GongStruct interface
func (otherappearance *OtherAppearance) GetName() (res string) {
	return otherappearance.Name
}

// for satisfaction of GongStruct interface
func (otherappearance *OtherAppearance) SetName(name string) {
	otherappearance.Name = name
}

// Stage puts otherplay to the model stage
func (otherplay *OtherPlay) Stage(stage *Stage) *OtherPlay {
	__gong__stage(stage.OtherPlays, stage.OtherPlay_stagedOrder, stage.OtherPlay_orderStaged, &stage.OtherPlayOrder, stage.OtherPlays_mapString, otherplay, otherplay.Name)
	return otherplay
}

// StagePreserveOrder puts otherplay to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OtherPlayOrder
// - update stage.OtherPlayOrder accordingly
func (otherplay *OtherPlay) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OtherPlays, stage.OtherPlay_stagedOrder, stage.OtherPlay_orderStaged, &stage.OtherPlayOrder, stage.OtherPlays_mapString, otherplay, order, otherplay.Name)
}

// Unstage removes otherplay off the model stage
func (otherplay *OtherPlay) Unstage(stage *Stage) *OtherPlay {
	__gong__unstage(stage.OtherPlays, stage.OtherPlays_mapString, otherplay, otherplay.Name)
	return otherplay
}

// UnstageVoid removes otherplay off the model stage
func (otherplay *OtherPlay) UnstageVoid(stage *Stage) {
	otherplay.Unstage(stage)
}

func (otherplay *OtherPlay) StageVoid(stage *Stage) {
	otherplay.Stage(stage)
}

// for satisfaction of GongStruct interface
func (otherplay *OtherPlay) GetName() (res string) {
	return otherplay.Name
}

// for satisfaction of GongStruct interface
func (otherplay *OtherPlay) SetName(name string) {
	otherplay.Name = name
}

// Stage puts othertext to the model stage
func (othertext *OtherText) Stage(stage *Stage) *OtherText {
	__gong__stage(stage.OtherTexts, stage.OtherText_stagedOrder, stage.OtherText_orderStaged, &stage.OtherTextOrder, stage.OtherTexts_mapString, othertext, othertext.Name)
	return othertext
}

// StagePreserveOrder puts othertext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.OtherTextOrder
// - update stage.OtherTextOrder accordingly
func (othertext *OtherText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.OtherTexts, stage.OtherText_stagedOrder, stage.OtherText_orderStaged, &stage.OtherTextOrder, stage.OtherTexts_mapString, othertext, order, othertext.Name)
}

// Unstage removes othertext off the model stage
func (othertext *OtherText) Unstage(stage *Stage) *OtherText {
	__gong__unstage(stage.OtherTexts, stage.OtherTexts_mapString, othertext, othertext.Name)
	return othertext
}

// UnstageVoid removes othertext off the model stage
func (othertext *OtherText) UnstageVoid(stage *Stage) {
	othertext.Unstage(stage)
}

func (othertext *OtherText) StageVoid(stage *Stage) {
	othertext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (othertext *OtherText) GetName() (res string) {
	return othertext.Name
}

// for satisfaction of GongStruct interface
func (othertext *OtherText) SetName(name string) {
	othertext.Name = name
}

// Stage puts pagelayout to the model stage
func (pagelayout *PageLayout) Stage(stage *Stage) *PageLayout {
	__gong__stage(stage.PageLayouts, stage.PageLayout_stagedOrder, stage.PageLayout_orderStaged, &stage.PageLayoutOrder, stage.PageLayouts_mapString, pagelayout, pagelayout.Name)
	return pagelayout
}

// StagePreserveOrder puts pagelayout to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PageLayoutOrder
// - update stage.PageLayoutOrder accordingly
func (pagelayout *PageLayout) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PageLayouts, stage.PageLayout_stagedOrder, stage.PageLayout_orderStaged, &stage.PageLayoutOrder, stage.PageLayouts_mapString, pagelayout, order, pagelayout.Name)
}

// Unstage removes pagelayout off the model stage
func (pagelayout *PageLayout) Unstage(stage *Stage) *PageLayout {
	__gong__unstage(stage.PageLayouts, stage.PageLayouts_mapString, pagelayout, pagelayout.Name)
	return pagelayout
}

// UnstageVoid removes pagelayout off the model stage
func (pagelayout *PageLayout) UnstageVoid(stage *Stage) {
	pagelayout.Unstage(stage)
}

func (pagelayout *PageLayout) StageVoid(stage *Stage) {
	pagelayout.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pagelayout *PageLayout) GetName() (res string) {
	return pagelayout.Name
}

// for satisfaction of GongStruct interface
func (pagelayout *PageLayout) SetName(name string) {
	pagelayout.Name = name
}

// Stage puts pagemargins to the model stage
func (pagemargins *PageMargins) Stage(stage *Stage) *PageMargins {
	__gong__stage(stage.PageMarginss, stage.PageMargins_stagedOrder, stage.PageMargins_orderStaged, &stage.PageMarginsOrder, stage.PageMarginss_mapString, pagemargins, pagemargins.Name)
	return pagemargins
}

// StagePreserveOrder puts pagemargins to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PageMarginsOrder
// - update stage.PageMarginsOrder accordingly
func (pagemargins *PageMargins) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.PageMarginss, stage.PageMargins_stagedOrder, stage.PageMargins_orderStaged, &stage.PageMarginsOrder, stage.PageMarginss_mapString, pagemargins, order, pagemargins.Name)
}

// Unstage removes pagemargins off the model stage
func (pagemargins *PageMargins) Unstage(stage *Stage) *PageMargins {
	__gong__unstage(stage.PageMarginss, stage.PageMarginss_mapString, pagemargins, pagemargins.Name)
	return pagemargins
}

// UnstageVoid removes pagemargins off the model stage
func (pagemargins *PageMargins) UnstageVoid(stage *Stage) {
	pagemargins.Unstage(stage)
}

func (pagemargins *PageMargins) StageVoid(stage *Stage) {
	pagemargins.Stage(stage)
}

// for satisfaction of GongStruct interface
func (pagemargins *PageMargins) GetName() (res string) {
	return pagemargins.Name
}

// for satisfaction of GongStruct interface
func (pagemargins *PageMargins) SetName(name string) {
	pagemargins.Name = name
}

// Stage puts play to the model stage
func (play *Play) Stage(stage *Stage) *Play {
	__gong__stage(stage.Plays, stage.Play_stagedOrder, stage.Play_orderStaged, &stage.PlayOrder, stage.Plays_mapString, play, play.Name)
	return play
}

// StagePreserveOrder puts play to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlayOrder
// - update stage.PlayOrder accordingly
func (play *Play) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Plays, stage.Play_stagedOrder, stage.Play_orderStaged, &stage.PlayOrder, stage.Plays_mapString, play, order, play.Name)
}

// Unstage removes play off the model stage
func (play *Play) Unstage(stage *Stage) *Play {
	__gong__unstage(stage.Plays, stage.Plays_mapString, play, play.Name)
	return play
}

// UnstageVoid removes play off the model stage
func (play *Play) UnstageVoid(stage *Stage) {
	play.Unstage(stage)
}

func (play *Play) StageVoid(stage *Stage) {
	play.Stage(stage)
}

// for satisfaction of GongStruct interface
func (play *Play) GetName() (res string) {
	return play.Name
}

// for satisfaction of GongStruct interface
func (play *Play) SetName(name string) {
	play.Name = name
}

// Stage puts scaling to the model stage
func (scaling *Scaling) Stage(stage *Stage) *Scaling {
	__gong__stage(stage.Scalings, stage.Scaling_stagedOrder, stage.Scaling_orderStaged, &stage.ScalingOrder, stage.Scalings_mapString, scaling, scaling.Name)
	return scaling
}

// StagePreserveOrder puts scaling to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ScalingOrder
// - update stage.ScalingOrder accordingly
func (scaling *Scaling) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Scalings, stage.Scaling_stagedOrder, stage.Scaling_orderStaged, &stage.ScalingOrder, stage.Scalings_mapString, scaling, order, scaling.Name)
}

// Unstage removes scaling off the model stage
func (scaling *Scaling) Unstage(stage *Stage) *Scaling {
	__gong__unstage(stage.Scalings, stage.Scalings_mapString, scaling, scaling.Name)
	return scaling
}

// UnstageVoid removes scaling off the model stage
func (scaling *Scaling) UnstageVoid(stage *Stage) {
	scaling.Unstage(stage)
}

func (scaling *Scaling) StageVoid(stage *Stage) {
	scaling.Stage(stage)
}

// for satisfaction of GongStruct interface
func (scaling *Scaling) GetName() (res string) {
	return scaling.Name
}

// for satisfaction of GongStruct interface
func (scaling *Scaling) SetName(name string) {
	scaling.Name = name
}

// Stage puts segno to the model stage
func (segno *Segno) Stage(stage *Stage) *Segno {
	__gong__stage(stage.Segnos, stage.Segno_stagedOrder, stage.Segno_orderStaged, &stage.SegnoOrder, stage.Segnos_mapString, segno, segno.Name)
	return segno
}

// StagePreserveOrder puts segno to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SegnoOrder
// - update stage.SegnoOrder accordingly
func (segno *Segno) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Segnos, stage.Segno_stagedOrder, stage.Segno_orderStaged, &stage.SegnoOrder, stage.Segnos_mapString, segno, order, segno.Name)
}

// Unstage removes segno off the model stage
func (segno *Segno) Unstage(stage *Stage) *Segno {
	__gong__unstage(stage.Segnos, stage.Segnos_mapString, segno, segno.Name)
	return segno
}

// UnstageVoid removes segno off the model stage
func (segno *Segno) UnstageVoid(stage *Stage) {
	segno.Unstage(stage)
}

func (segno *Segno) StageVoid(stage *Stage) {
	segno.Stage(stage)
}

// for satisfaction of GongStruct interface
func (segno *Segno) GetName() (res string) {
	return segno.Name
}

// for satisfaction of GongStruct interface
func (segno *Segno) SetName(name string) {
	segno.Name = name
}

// Stage puts stafflayout to the model stage
func (stafflayout *StaffLayout) Stage(stage *Stage) *StaffLayout {
	__gong__stage(stage.StaffLayouts, stage.StaffLayout_stagedOrder, stage.StaffLayout_orderStaged, &stage.StaffLayoutOrder, stage.StaffLayouts_mapString, stafflayout, stafflayout.Name)
	return stafflayout
}

// StagePreserveOrder puts stafflayout to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StaffLayoutOrder
// - update stage.StaffLayoutOrder accordingly
func (stafflayout *StaffLayout) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StaffLayouts, stage.StaffLayout_stagedOrder, stage.StaffLayout_orderStaged, &stage.StaffLayoutOrder, stage.StaffLayouts_mapString, stafflayout, order, stafflayout.Name)
}

// Unstage removes stafflayout off the model stage
func (stafflayout *StaffLayout) Unstage(stage *Stage) *StaffLayout {
	__gong__unstage(stage.StaffLayouts, stage.StaffLayouts_mapString, stafflayout, stafflayout.Name)
	return stafflayout
}

// UnstageVoid removes stafflayout off the model stage
func (stafflayout *StaffLayout) UnstageVoid(stage *Stage) {
	stafflayout.Unstage(stage)
}

func (stafflayout *StaffLayout) StageVoid(stage *Stage) {
	stafflayout.Stage(stage)
}

// for satisfaction of GongStruct interface
func (stafflayout *StaffLayout) GetName() (res string) {
	return stafflayout.Name
}

// for satisfaction of GongStruct interface
func (stafflayout *StaffLayout) SetName(name string) {
	stafflayout.Name = name
}

// Stage puts stringtype to the model stage
func (stringtype *StringType) Stage(stage *Stage) *StringType {
	__gong__stage(stage.StringTypes, stage.StringType_stagedOrder, stage.StringType_orderStaged, &stage.StringTypeOrder, stage.StringTypes_mapString, stringtype, stringtype.Name)
	return stringtype
}

// StagePreserveOrder puts stringtype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StringTypeOrder
// - update stage.StringTypeOrder accordingly
func (stringtype *StringType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StringTypes, stage.StringType_stagedOrder, stage.StringType_orderStaged, &stage.StringTypeOrder, stage.StringTypes_mapString, stringtype, order, stringtype.Name)
}

// Unstage removes stringtype off the model stage
func (stringtype *StringType) Unstage(stage *Stage) *StringType {
	__gong__unstage(stage.StringTypes, stage.StringTypes_mapString, stringtype, stringtype.Name)
	return stringtype
}

// UnstageVoid removes stringtype off the model stage
func (stringtype *StringType) UnstageVoid(stage *Stage) {
	stringtype.Unstage(stage)
}

func (stringtype *StringType) StageVoid(stage *Stage) {
	stringtype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (stringtype *StringType) GetName() (res string) {
	return stringtype.Name
}

// for satisfaction of GongStruct interface
func (stringtype *StringType) SetName(name string) {
	stringtype.Name = name
}

// Stage puts supports to the model stage
func (supports *Supports) Stage(stage *Stage) *Supports {
	__gong__stage(stage.Supportss, stage.Supports_stagedOrder, stage.Supports_orderStaged, &stage.SupportsOrder, stage.Supportss_mapString, supports, supports.Name)
	return supports
}

// StagePreserveOrder puts supports to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SupportsOrder
// - update stage.SupportsOrder accordingly
func (supports *Supports) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Supportss, stage.Supports_stagedOrder, stage.Supports_orderStaged, &stage.SupportsOrder, stage.Supportss_mapString, supports, order, supports.Name)
}

// Unstage removes supports off the model stage
func (supports *Supports) Unstage(stage *Stage) *Supports {
	__gong__unstage(stage.Supportss, stage.Supportss_mapString, supports, supports.Name)
	return supports
}

// UnstageVoid removes supports off the model stage
func (supports *Supports) UnstageVoid(stage *Stage) {
	supports.Unstage(stage)
}

func (supports *Supports) StageVoid(stage *Stage) {
	supports.Stage(stage)
}

// for satisfaction of GongStruct interface
func (supports *Supports) GetName() (res string) {
	return supports.Name
}

// for satisfaction of GongStruct interface
func (supports *Supports) SetName(name string) {
	supports.Name = name
}

// Stage puts systemdividers to the model stage
func (systemdividers *SystemDividers) Stage(stage *Stage) *SystemDividers {
	__gong__stage(stage.SystemDividerss, stage.SystemDividers_stagedOrder, stage.SystemDividers_orderStaged, &stage.SystemDividersOrder, stage.SystemDividerss_mapString, systemdividers, systemdividers.Name)
	return systemdividers
}

// StagePreserveOrder puts systemdividers to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemDividersOrder
// - update stage.SystemDividersOrder accordingly
func (systemdividers *SystemDividers) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SystemDividerss, stage.SystemDividers_stagedOrder, stage.SystemDividers_orderStaged, &stage.SystemDividersOrder, stage.SystemDividerss_mapString, systemdividers, order, systemdividers.Name)
}

// Unstage removes systemdividers off the model stage
func (systemdividers *SystemDividers) Unstage(stage *Stage) *SystemDividers {
	__gong__unstage(stage.SystemDividerss, stage.SystemDividerss_mapString, systemdividers, systemdividers.Name)
	return systemdividers
}

// UnstageVoid removes systemdividers off the model stage
func (systemdividers *SystemDividers) UnstageVoid(stage *Stage) {
	systemdividers.Unstage(stage)
}

func (systemdividers *SystemDividers) StageVoid(stage *Stage) {
	systemdividers.Stage(stage)
}

// for satisfaction of GongStruct interface
func (systemdividers *SystemDividers) GetName() (res string) {
	return systemdividers.Name
}

// for satisfaction of GongStruct interface
func (systemdividers *SystemDividers) SetName(name string) {
	systemdividers.Name = name
}

// Stage puts systemlayout to the model stage
func (systemlayout *SystemLayout) Stage(stage *Stage) *SystemLayout {
	__gong__stage(stage.SystemLayouts, stage.SystemLayout_stagedOrder, stage.SystemLayout_orderStaged, &stage.SystemLayoutOrder, stage.SystemLayouts_mapString, systemlayout, systemlayout.Name)
	return systemlayout
}

// StagePreserveOrder puts systemlayout to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemLayoutOrder
// - update stage.SystemLayoutOrder accordingly
func (systemlayout *SystemLayout) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SystemLayouts, stage.SystemLayout_stagedOrder, stage.SystemLayout_orderStaged, &stage.SystemLayoutOrder, stage.SystemLayouts_mapString, systemlayout, order, systemlayout.Name)
}

// Unstage removes systemlayout off the model stage
func (systemlayout *SystemLayout) Unstage(stage *Stage) *SystemLayout {
	__gong__unstage(stage.SystemLayouts, stage.SystemLayouts_mapString, systemlayout, systemlayout.Name)
	return systemlayout
}

// UnstageVoid removes systemlayout off the model stage
func (systemlayout *SystemLayout) UnstageVoid(stage *Stage) {
	systemlayout.Unstage(stage)
}

func (systemlayout *SystemLayout) StageVoid(stage *Stage) {
	systemlayout.Stage(stage)
}

// for satisfaction of GongStruct interface
func (systemlayout *SystemLayout) GetName() (res string) {
	return systemlayout.Name
}

// for satisfaction of GongStruct interface
func (systemlayout *SystemLayout) SetName(name string) {
	systemlayout.Name = name
}

// Stage puts systemmargins to the model stage
func (systemmargins *SystemMargins) Stage(stage *Stage) *SystemMargins {
	__gong__stage(stage.SystemMarginss, stage.SystemMargins_stagedOrder, stage.SystemMargins_orderStaged, &stage.SystemMarginsOrder, stage.SystemMarginss_mapString, systemmargins, systemmargins.Name)
	return systemmargins
}

// StagePreserveOrder puts systemmargins to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemMarginsOrder
// - update stage.SystemMarginsOrder accordingly
func (systemmargins *SystemMargins) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SystemMarginss, stage.SystemMargins_stagedOrder, stage.SystemMargins_orderStaged, &stage.SystemMarginsOrder, stage.SystemMarginss_mapString, systemmargins, order, systemmargins.Name)
}

// Unstage removes systemmargins off the model stage
func (systemmargins *SystemMargins) Unstage(stage *Stage) *SystemMargins {
	__gong__unstage(stage.SystemMarginss, stage.SystemMarginss_mapString, systemmargins, systemmargins.Name)
	return systemmargins
}

// UnstageVoid removes systemmargins off the model stage
func (systemmargins *SystemMargins) UnstageVoid(stage *Stage) {
	systemmargins.Unstage(stage)
}

func (systemmargins *SystemMargins) StageVoid(stage *Stage) {
	systemmargins.Stage(stage)
}

// for satisfaction of GongStruct interface
func (systemmargins *SystemMargins) GetName() (res string) {
	return systemmargins.Name
}

// for satisfaction of GongStruct interface
func (systemmargins *SystemMargins) SetName(name string) {
	systemmargins.Name = name
}

// Stage puts typedtext to the model stage
func (typedtext *TypedText) Stage(stage *Stage) *TypedText {
	__gong__stage(stage.TypedTexts, stage.TypedText_stagedOrder, stage.TypedText_orderStaged, &stage.TypedTextOrder, stage.TypedTexts_mapString, typedtext, typedtext.Name)
	return typedtext
}

// StagePreserveOrder puts typedtext to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TypedTextOrder
// - update stage.TypedTextOrder accordingly
func (typedtext *TypedText) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TypedTexts, stage.TypedText_stagedOrder, stage.TypedText_orderStaged, &stage.TypedTextOrder, stage.TypedTexts_mapString, typedtext, order, typedtext.Name)
}

// Unstage removes typedtext off the model stage
func (typedtext *TypedText) Unstage(stage *Stage) *TypedText {
	__gong__unstage(stage.TypedTexts, stage.TypedTexts_mapString, typedtext, typedtext.Name)
	return typedtext
}

// UnstageVoid removes typedtext off the model stage
func (typedtext *TypedText) UnstageVoid(stage *Stage) {
	typedtext.Unstage(stage)
}

func (typedtext *TypedText) StageVoid(stage *Stage) {
	typedtext.Stage(stage)
}

// for satisfaction of GongStruct interface
func (typedtext *TypedText) GetName() (res string) {
	return typedtext.Name
}

// for satisfaction of GongStruct interface
func (typedtext *TypedText) SetName(name string) {
	typedtext.Name = name
}

// Stage puts wavyline to the model stage
func (wavyline *WavyLine) Stage(stage *Stage) *WavyLine {
	__gong__stage(stage.WavyLines, stage.WavyLine_stagedOrder, stage.WavyLine_orderStaged, &stage.WavyLineOrder, stage.WavyLines_mapString, wavyline, wavyline.Name)
	return wavyline
}

// StagePreserveOrder puts wavyline to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.WavyLineOrder
// - update stage.WavyLineOrder accordingly
func (wavyline *WavyLine) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.WavyLines, stage.WavyLine_stagedOrder, stage.WavyLine_orderStaged, &stage.WavyLineOrder, stage.WavyLines_mapString, wavyline, order, wavyline.Name)
}

// Unstage removes wavyline off the model stage
func (wavyline *WavyLine) Unstage(stage *Stage) *WavyLine {
	__gong__unstage(stage.WavyLines, stage.WavyLines_mapString, wavyline, wavyline.Name)
	return wavyline
}

// UnstageVoid removes wavyline off the model stage
func (wavyline *WavyLine) UnstageVoid(stage *Stage) {
	wavyline.Unstage(stage)
}

func (wavyline *WavyLine) StageVoid(stage *Stage) {
	wavyline.Stage(stage)
}

// for satisfaction of GongStruct interface
func (wavyline *WavyLine) GetName() (res string) {
	return wavyline.Name
}

// for satisfaction of GongStruct interface
func (wavyline *WavyLine) SetName(name string) {
	wavyline.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AccidentalTexts, &stage.AccidentalTexts_mapString, &stage.AccidentalText_stagedOrder, &stage.AccidentalTextOrder)

	__gong__resetStageType(&stage.Appearances, &stage.Appearances_mapString, &stage.Appearance_stagedOrder, &stage.AppearanceOrder)

	__gong__resetStageType(&stage.Bookmarks, &stage.Bookmarks_mapString, &stage.Bookmark_stagedOrder, &stage.BookmarkOrder)

	__gong__resetStageType(&stage.Codas, &stage.Codas_mapString, &stage.Coda_stagedOrder, &stage.CodaOrder)

	__gong__resetStageType(&stage.Distances, &stage.Distances_mapString, &stage.Distance_stagedOrder, &stage.DistanceOrder)

	__gong__resetStageType(&stage.Dynamicss, &stage.Dynamicss_mapString, &stage.Dynamics_stagedOrder, &stage.DynamicsOrder)

	__gong__resetStageType(&stage.Emptys, &stage.Emptys_mapString, &stage.Empty_stagedOrder, &stage.EmptyOrder)

	__gong__resetStageType(&stage.EmptyPlacements, &stage.EmptyPlacements_mapString, &stage.EmptyPlacement_stagedOrder, &stage.EmptyPlacementOrder)

	__gong__resetStageType(&stage.EmptyPlacementSmufls, &stage.EmptyPlacementSmufls_mapString, &stage.EmptyPlacementSmufl_stagedOrder, &stage.EmptyPlacementSmuflOrder)

	__gong__resetStageType(&stage.EmptyPrintObjectStyleAligns, &stage.EmptyPrintObjectStyleAligns_mapString, &stage.EmptyPrintObjectStyleAlign_stagedOrder, &stage.EmptyPrintObjectStyleAlignOrder)

	__gong__resetStageType(&stage.EmptyPrintStyles, &stage.EmptyPrintStyles_mapString, &stage.EmptyPrintStyle_stagedOrder, &stage.EmptyPrintStyleOrder)

	__gong__resetStageType(&stage.EmptyPrintStyleAligns, &stage.EmptyPrintStyleAligns_mapString, &stage.EmptyPrintStyleAlign_stagedOrder, &stage.EmptyPrintStyleAlignOrder)

	__gong__resetStageType(&stage.EmptyPrintStyleAlignIds, &stage.EmptyPrintStyleAlignIds_mapString, &stage.EmptyPrintStyleAlignId_stagedOrder, &stage.EmptyPrintStyleAlignIdOrder)

	__gong__resetStageType(&stage.EmptyTrillSounds, &stage.EmptyTrillSounds_mapString, &stage.EmptyTrillSound_stagedOrder, &stage.EmptyTrillSoundOrder)

	__gong__resetStageType(&stage.Encodings, &stage.Encodings_mapString, &stage.Encoding_stagedOrder, &stage.EncodingOrder)

	__gong__resetStageType(&stage.Fermatas, &stage.Fermatas_mapString, &stage.Fermata_stagedOrder, &stage.FermataOrder)

	__gong__resetStageType(&stage.Fingerings, &stage.Fingerings_mapString, &stage.Fingering_stagedOrder, &stage.FingeringOrder)

	__gong__resetStageType(&stage.FormattedSymbols, &stage.FormattedSymbols_mapString, &stage.FormattedSymbol_stagedOrder, &stage.FormattedSymbolOrder)

	__gong__resetStageType(&stage.FormattedSymbolIds, &stage.FormattedSymbolIds_mapString, &stage.FormattedSymbolId_stagedOrder, &stage.FormattedSymbolIdOrder)

	__gong__resetStageType(&stage.FormattedTexts, &stage.FormattedTexts_mapString, &stage.FormattedText_stagedOrder, &stage.FormattedTextOrder)

	__gong__resetStageType(&stage.FormattedTextIds, &stage.FormattedTextIds_mapString, &stage.FormattedTextId_stagedOrder, &stage.FormattedTextIdOrder)

	__gong__resetStageType(&stage.Frets, &stage.Frets_mapString, &stage.Fret_stagedOrder, &stage.FretOrder)

	__gong__resetStageType(&stage.Glyphs, &stage.Glyphs_mapString, &stage.Glyph_stagedOrder, &stage.GlyphOrder)

	__gong__resetStageType(&stage.HorizontalTurns, &stage.HorizontalTurns_mapString, &stage.HorizontalTurn_stagedOrder, &stage.HorizontalTurnOrder)

	__gong__resetStageType(&stage.Identifications, &stage.Identifications_mapString, &stage.Identification_stagedOrder, &stage.IdentificationOrder)

	__gong__resetStageType(&stage.Levels, &stage.Levels_mapString, &stage.Level_stagedOrder, &stage.LevelOrder)

	__gong__resetStageType(&stage.LineWidths, &stage.LineWidths_mapString, &stage.LineWidth_stagedOrder, &stage.LineWidthOrder)

	__gong__resetStageType(&stage.Links, &stage.Links_mapString, &stage.Link_stagedOrder, &stage.LinkOrder)

	__gong__resetStageType(&stage.MeasureLayouts, &stage.MeasureLayouts_mapString, &stage.MeasureLayout_stagedOrder, &stage.MeasureLayoutOrder)

	__gong__resetStageType(&stage.MidiDevices, &stage.MidiDevices_mapString, &stage.MidiDevice_stagedOrder, &stage.MidiDeviceOrder)

	__gong__resetStageType(&stage.MidiInstruments, &stage.MidiInstruments_mapString, &stage.MidiInstrument_stagedOrder, &stage.MidiInstrumentOrder)

	__gong__resetStageType(&stage.Miscellaneouss, &stage.Miscellaneouss_mapString, &stage.Miscellaneous_stagedOrder, &stage.MiscellaneousOrder)

	__gong__resetStageType(&stage.MiscellaneousFields, &stage.MiscellaneousFields_mapString, &stage.MiscellaneousField_stagedOrder, &stage.MiscellaneousFieldOrder)

	__gong__resetStageType(&stage.NameDisplays, &stage.NameDisplays_mapString, &stage.NameDisplay_stagedOrder, &stage.NameDisplayOrder)

	__gong__resetStageType(&stage.NoteSizes, &stage.NoteSizes_mapString, &stage.NoteSize_stagedOrder, &stage.NoteSizeOrder)

	__gong__resetStageType(&stage.OtherAppearances, &stage.OtherAppearances_mapString, &stage.OtherAppearance_stagedOrder, &stage.OtherAppearanceOrder)

	__gong__resetStageType(&stage.OtherPlays, &stage.OtherPlays_mapString, &stage.OtherPlay_stagedOrder, &stage.OtherPlayOrder)

	__gong__resetStageType(&stage.OtherTexts, &stage.OtherTexts_mapString, &stage.OtherText_stagedOrder, &stage.OtherTextOrder)

	__gong__resetStageType(&stage.PageLayouts, &stage.PageLayouts_mapString, &stage.PageLayout_stagedOrder, &stage.PageLayoutOrder)

	__gong__resetStageType(&stage.PageMarginss, &stage.PageMarginss_mapString, &stage.PageMargins_stagedOrder, &stage.PageMarginsOrder)

	__gong__resetStageType(&stage.Plays, &stage.Plays_mapString, &stage.Play_stagedOrder, &stage.PlayOrder)

	__gong__resetStageType(&stage.Scalings, &stage.Scalings_mapString, &stage.Scaling_stagedOrder, &stage.ScalingOrder)

	__gong__resetStageType(&stage.Segnos, &stage.Segnos_mapString, &stage.Segno_stagedOrder, &stage.SegnoOrder)

	__gong__resetStageType(&stage.StaffLayouts, &stage.StaffLayouts_mapString, &stage.StaffLayout_stagedOrder, &stage.StaffLayoutOrder)

	__gong__resetStageType(&stage.StringTypes, &stage.StringTypes_mapString, &stage.StringType_stagedOrder, &stage.StringTypeOrder)

	__gong__resetStageType(&stage.Supportss, &stage.Supportss_mapString, &stage.Supports_stagedOrder, &stage.SupportsOrder)

	__gong__resetStageType(&stage.SystemDividerss, &stage.SystemDividerss_mapString, &stage.SystemDividers_stagedOrder, &stage.SystemDividersOrder)

	__gong__resetStageType(&stage.SystemLayouts, &stage.SystemLayouts_mapString, &stage.SystemLayout_stagedOrder, &stage.SystemLayoutOrder)

	__gong__resetStageType(&stage.SystemMarginss, &stage.SystemMarginss_mapString, &stage.SystemMargins_stagedOrder, &stage.SystemMarginsOrder)

	__gong__resetStageType(&stage.TypedTexts, &stage.TypedTexts_mapString, &stage.TypedText_stagedOrder, &stage.TypedTextOrder)

	__gong__resetStageType(&stage.WavyLines, &stage.WavyLines_mapString, &stage.WavyLine_stagedOrder, &stage.WavyLineOrder)

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
func (AccidentalText) GongGetAssociationName() any {
	return &AccidentalText{
	}
}

func (Appearance) GongGetAssociationName() any {
	return &Appearance{
			LineWidth: []*LineWidth{{Name: "LineWidth"}},
			NoteSize: []*NoteSize{{Name: "NoteSize"}},
			Distance: []*Distance{{Name: "Distance"}},
			Glyph: []*Glyph{{Name: "Glyph"}},
			OtherAppearance: []*OtherAppearance{{Name: "OtherAppearance"}},
	}
}

func (Bookmark) GongGetAssociationName() any {
	return &Bookmark{
	}
}

func (Coda) GongGetAssociationName() any {
	return &Coda{
	}
}

func (Distance) GongGetAssociationName() any {
	return &Distance{
	}
}

func (Dynamics) GongGetAssociationName() any {
	return &Dynamics{
			P: []*Empty{{Name: "P"}},
			Pp: []*Empty{{Name: "Pp"}},
			Ppp: []*Empty{{Name: "Ppp"}},
			Pppp: []*Empty{{Name: "Pppp"}},
			Ppppp: []*Empty{{Name: "Ppppp"}},
			Pppppp: []*Empty{{Name: "Pppppp"}},
			F: []*Empty{{Name: "F"}},
			Ff: []*Empty{{Name: "Ff"}},
			Fff: []*Empty{{Name: "Fff"}},
			Ffff: []*Empty{{Name: "Ffff"}},
			Fffff: []*Empty{{Name: "Fffff"}},
			Ffffff: []*Empty{{Name: "Ffffff"}},
			Mp: []*Empty{{Name: "Mp"}},
			Mf: []*Empty{{Name: "Mf"}},
			Sf: []*Empty{{Name: "Sf"}},
			Sfp: []*Empty{{Name: "Sfp"}},
			Sfpp: []*Empty{{Name: "Sfpp"}},
			Fp: []*Empty{{Name: "Fp"}},
			Rf: []*Empty{{Name: "Rf"}},
			Rfz: []*Empty{{Name: "Rfz"}},
			Sfz: []*Empty{{Name: "Sfz"}},
			Sffz: []*Empty{{Name: "Sffz"}},
			Fz: []*Empty{{Name: "Fz"}},
			N: []*Empty{{Name: "N"}},
			Pf: []*Empty{{Name: "Pf"}},
			Sfzp: []*Empty{{Name: "Sfzp"}},
			OtherDynamics: []*OtherText{{Name: "OtherDynamics"}},
	}
}

func (Empty) GongGetAssociationName() any {
	return &Empty{
	}
}

func (EmptyPlacement) GongGetAssociationName() any {
	return &EmptyPlacement{
	}
}

func (EmptyPlacementSmufl) GongGetAssociationName() any {
	return &EmptyPlacementSmufl{
	}
}

func (EmptyPrintObjectStyleAlign) GongGetAssociationName() any {
	return &EmptyPrintObjectStyleAlign{
	}
}

func (EmptyPrintStyle) GongGetAssociationName() any {
	return &EmptyPrintStyle{
	}
}

func (EmptyPrintStyleAlign) GongGetAssociationName() any {
	return &EmptyPrintStyleAlign{
	}
}

func (EmptyPrintStyleAlignId) GongGetAssociationName() any {
	return &EmptyPrintStyleAlignId{
	}
}

func (EmptyTrillSound) GongGetAssociationName() any {
	return &EmptyTrillSound{
	}
}

func (Encoding) GongGetAssociationName() any {
	return &Encoding{
			Encoder: []*TypedText{{Name: "Encoder"}},
			Supports: []*Supports{{Name: "Supports"}},
	}
}

func (Fermata) GongGetAssociationName() any {
	return &Fermata{
	}
}

func (Fingering) GongGetAssociationName() any {
	return &Fingering{
	}
}

func (FormattedSymbol) GongGetAssociationName() any {
	return &FormattedSymbol{
	}
}

func (FormattedSymbolId) GongGetAssociationName() any {
	return &FormattedSymbolId{
	}
}

func (FormattedText) GongGetAssociationName() any {
	return &FormattedText{
	}
}

func (FormattedTextId) GongGetAssociationName() any {
	return &FormattedTextId{
	}
}

func (Fret) GongGetAssociationName() any {
	return &Fret{
	}
}

func (Glyph) GongGetAssociationName() any {
	return &Glyph{
	}
}

func (HorizontalTurn) GongGetAssociationName() any {
	return &HorizontalTurn{
	}
}

func (Identification) GongGetAssociationName() any {
	return &Identification{
			Creator: []*TypedText{{Name: "Creator"}},
			Rights: []*TypedText{{Name: "Rights"}},
			Encoding: &Encoding{Name: "Encoding"},
			Relation: []*TypedText{{Name: "Relation"}},
			Miscellaneous: &Miscellaneous{Name: "Miscellaneous"},
	}
}

func (Level) GongGetAssociationName() any {
	return &Level{
	}
}

func (LineWidth) GongGetAssociationName() any {
	return &LineWidth{
	}
}

func (Link) GongGetAssociationName() any {
	return &Link{
	}
}

func (MeasureLayout) GongGetAssociationName() any {
	return &MeasureLayout{
	}
}

func (MidiDevice) GongGetAssociationName() any {
	return &MidiDevice{
	}
}

func (MidiInstrument) GongGetAssociationName() any {
	return &MidiInstrument{
	}
}

func (Miscellaneous) GongGetAssociationName() any {
	return &Miscellaneous{
			MiscellaneousField: []*MiscellaneousField{{Name: "MiscellaneousField"}},
	}
}

func (MiscellaneousField) GongGetAssociationName() any {
	return &MiscellaneousField{
	}
}

func (NameDisplay) GongGetAssociationName() any {
	return &NameDisplay{
			DisplayText: []*FormattedText{{Name: "DisplayText"}},
			AccidentalText: []*AccidentalText{{Name: "AccidentalText"}},
	}
}

func (NoteSize) GongGetAssociationName() any {
	return &NoteSize{
	}
}

func (OtherAppearance) GongGetAssociationName() any {
	return &OtherAppearance{
	}
}

func (OtherPlay) GongGetAssociationName() any {
	return &OtherPlay{
	}
}

func (OtherText) GongGetAssociationName() any {
	return &OtherText{
	}
}

func (PageLayout) GongGetAssociationName() any {
	return &PageLayout{
			PageMargins: []*PageMargins{{Name: "PageMargins"}},
	}
}

func (PageMargins) GongGetAssociationName() any {
	return &PageMargins{
	}
}

func (Play) GongGetAssociationName() any {
	return &Play{
			OtherPlay: []*OtherPlay{{Name: "OtherPlay"}},
	}
}

func (Scaling) GongGetAssociationName() any {
	return &Scaling{
	}
}

func (Segno) GongGetAssociationName() any {
	return &Segno{
	}
}

func (StaffLayout) GongGetAssociationName() any {
	return &StaffLayout{
	}
}

func (StringType) GongGetAssociationName() any {
	return &StringType{
	}
}

func (Supports) GongGetAssociationName() any {
	return &Supports{
	}
}

func (SystemDividers) GongGetAssociationName() any {
	return &SystemDividers{
			LeftDivider: &EmptyPrintObjectStyleAlign{Name: "LeftDivider"},
			RightDivider: &EmptyPrintObjectStyleAlign{Name: "RightDivider"},
	}
}

func (SystemLayout) GongGetAssociationName() any {
	return &SystemLayout{
			SystemMargins: &SystemMargins{Name: "SystemMargins"},
			SystemDividers: &SystemDividers{Name: "SystemDividers"},
	}
}

func (SystemMargins) GongGetAssociationName() any {
	return &SystemMargins{
	}
}

func (TypedText) GongGetAssociationName() any {
	return &TypedText{
	}
}

func (WavyLine) GongGetAssociationName() any {
	return &WavyLine{
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
	// reverse maps of direct associations of AccidentalText
	case AccidentalText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Appearance
	case Appearance:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bookmark
	case Bookmark:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Coda
	case Coda:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Distance
	case Distance:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Dynamics
	case Dynamics:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Empty
	case Empty:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPlacement
	case EmptyPlacement:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPlacementSmufl
	case EmptyPlacementSmufl:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPrintObjectStyleAlign
	case EmptyPrintObjectStyleAlign:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPrintStyle
	case EmptyPrintStyle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPrintStyleAlign
	case EmptyPrintStyleAlign:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPrintStyleAlignId
	case EmptyPrintStyleAlignId:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyTrillSound
	case EmptyTrillSound:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Encoding
	case Encoding:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Fermata
	case Fermata:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Fingering
	case Fingering:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormattedSymbol
	case FormattedSymbol:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormattedSymbolId
	case FormattedSymbolId:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormattedText
	case FormattedText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormattedTextId
	case FormattedTextId:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Fret
	case Fret:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Glyph
	case Glyph:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HorizontalTurn
	case HorizontalTurn:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Identification
	case Identification:
		switch fieldname {
		// insertion point for per direct association field
		case "Encoding":
			res := make(map[*Encoding][]*Identification)
			for identification := range stage.Identifications {
				if identification.Encoding != nil {
					encoding_ := identification.Encoding
					var identifications []*Identification
					_, ok := res[encoding_]
					if ok {
						identifications = res[encoding_]
					} else {
						identifications = make([]*Identification, 0)
					}
					identifications = append(identifications, identification)
					res[encoding_] = identifications
				}
			}
			return any(res).(map[*End][]*Start)
		case "Miscellaneous":
			res := make(map[*Miscellaneous][]*Identification)
			for identification := range stage.Identifications {
				if identification.Miscellaneous != nil {
					miscellaneous_ := identification.Miscellaneous
					var identifications []*Identification
					_, ok := res[miscellaneous_]
					if ok {
						identifications = res[miscellaneous_]
					} else {
						identifications = make([]*Identification, 0)
					}
					identifications = append(identifications, identification)
					res[miscellaneous_] = identifications
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Level
	case Level:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LineWidth
	case LineWidth:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureLayout
	case MeasureLayout:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MidiDevice
	case MidiDevice:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MidiInstrument
	case MidiInstrument:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Miscellaneous
	case Miscellaneous:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MiscellaneousField
	case MiscellaneousField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NameDisplay
	case NameDisplay:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteSize
	case NoteSize:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherAppearance
	case OtherAppearance:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherPlay
	case OtherPlay:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherText
	case OtherText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PageLayout
	case PageLayout:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PageMargins
	case PageMargins:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Play
	case Play:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Scaling
	case Scaling:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Segno
	case Segno:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StaffLayout
	case StaffLayout:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StringType
	case StringType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Supports
	case Supports:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SystemDividers
	case SystemDividers:
		switch fieldname {
		// insertion point for per direct association field
		case "LeftDivider":
			res := make(map[*EmptyPrintObjectStyleAlign][]*SystemDividers)
			for systemdividers := range stage.SystemDividerss {
				if systemdividers.LeftDivider != nil {
					emptyprintobjectstylealign_ := systemdividers.LeftDivider
					var systemdividerss []*SystemDividers
					_, ok := res[emptyprintobjectstylealign_]
					if ok {
						systemdividerss = res[emptyprintobjectstylealign_]
					} else {
						systemdividerss = make([]*SystemDividers, 0)
					}
					systemdividerss = append(systemdividerss, systemdividers)
					res[emptyprintobjectstylealign_] = systemdividerss
				}
			}
			return any(res).(map[*End][]*Start)
		case "RightDivider":
			res := make(map[*EmptyPrintObjectStyleAlign][]*SystemDividers)
			for systemdividers := range stage.SystemDividerss {
				if systemdividers.RightDivider != nil {
					emptyprintobjectstylealign_ := systemdividers.RightDivider
					var systemdividerss []*SystemDividers
					_, ok := res[emptyprintobjectstylealign_]
					if ok {
						systemdividerss = res[emptyprintobjectstylealign_]
					} else {
						systemdividerss = make([]*SystemDividers, 0)
					}
					systemdividerss = append(systemdividerss, systemdividers)
					res[emptyprintobjectstylealign_] = systemdividerss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SystemLayout
	case SystemLayout:
		switch fieldname {
		// insertion point for per direct association field
		case "SystemMargins":
			res := make(map[*SystemMargins][]*SystemLayout)
			for systemlayout := range stage.SystemLayouts {
				if systemlayout.SystemMargins != nil {
					systemmargins_ := systemlayout.SystemMargins
					var systemlayouts []*SystemLayout
					_, ok := res[systemmargins_]
					if ok {
						systemlayouts = res[systemmargins_]
					} else {
						systemlayouts = make([]*SystemLayout, 0)
					}
					systemlayouts = append(systemlayouts, systemlayout)
					res[systemmargins_] = systemlayouts
				}
			}
			return any(res).(map[*End][]*Start)
		case "SystemDividers":
			res := make(map[*SystemDividers][]*SystemLayout)
			for systemlayout := range stage.SystemLayouts {
				if systemlayout.SystemDividers != nil {
					systemdividers_ := systemlayout.SystemDividers
					var systemlayouts []*SystemLayout
					_, ok := res[systemdividers_]
					if ok {
						systemlayouts = res[systemdividers_]
					} else {
						systemlayouts = make([]*SystemLayout, 0)
					}
					systemlayouts = append(systemlayouts, systemlayout)
					res[systemdividers_] = systemlayouts
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SystemMargins
	case SystemMargins:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TypedText
	case TypedText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of WavyLine
	case WavyLine:
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
	// reverse maps of direct associations of AccidentalText
	case AccidentalText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Appearance
	case Appearance:
		switch fieldname {
		// insertion point for per direct association field
		case "LineWidth":
			res := make(map[*LineWidth][]*Appearance)
			for appearance := range stage.Appearances {
				for _, linewidth_ := range appearance.LineWidth {
					res[linewidth_] = append(res[linewidth_], appearance)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteSize":
			res := make(map[*NoteSize][]*Appearance)
			for appearance := range stage.Appearances {
				for _, notesize_ := range appearance.NoteSize {
					res[notesize_] = append(res[notesize_], appearance)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Distance":
			res := make(map[*Distance][]*Appearance)
			for appearance := range stage.Appearances {
				for _, distance_ := range appearance.Distance {
					res[distance_] = append(res[distance_], appearance)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Glyph":
			res := make(map[*Glyph][]*Appearance)
			for appearance := range stage.Appearances {
				for _, glyph_ := range appearance.Glyph {
					res[glyph_] = append(res[glyph_], appearance)
				}
			}
			return any(res).(map[*End][]*Start)
		case "OtherAppearance":
			res := make(map[*OtherAppearance][]*Appearance)
			for appearance := range stage.Appearances {
				for _, otherappearance_ := range appearance.OtherAppearance {
					res[otherappearance_] = append(res[otherappearance_], appearance)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Bookmark
	case Bookmark:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Coda
	case Coda:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Distance
	case Distance:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Dynamics
	case Dynamics:
		switch fieldname {
		// insertion point for per direct association field
		case "P":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.P {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Pp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Ppp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Ppp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pppp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Pppp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Ppppp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Ppppp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pppppp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Pppppp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "F":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.F {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Ff":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Ff {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Fff":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Fff {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Ffff":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Ffff {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Fffff":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Fffff {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Ffffff":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Ffffff {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Mp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Mp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Mf":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Mf {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Sf":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Sf {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Sfp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Sfp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Sfpp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Sfpp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Fp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Fp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Rf":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Rf {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Rfz":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Rfz {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Sfz":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Sfz {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Sffz":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Sffz {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Fz":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Fz {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "N":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.N {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pf":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Pf {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Sfzp":
			res := make(map[*Empty][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, empty_ := range dynamics.Sfzp {
					res[empty_] = append(res[empty_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		case "OtherDynamics":
			res := make(map[*OtherText][]*Dynamics)
			for dynamics := range stage.Dynamicss {
				for _, othertext_ := range dynamics.OtherDynamics {
					res[othertext_] = append(res[othertext_], dynamics)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Empty
	case Empty:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPlacement
	case EmptyPlacement:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPlacementSmufl
	case EmptyPlacementSmufl:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPrintObjectStyleAlign
	case EmptyPrintObjectStyleAlign:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPrintStyle
	case EmptyPrintStyle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPrintStyleAlign
	case EmptyPrintStyleAlign:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyPrintStyleAlignId
	case EmptyPrintStyleAlignId:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EmptyTrillSound
	case EmptyTrillSound:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Encoding
	case Encoding:
		switch fieldname {
		// insertion point for per direct association field
		case "Encoder":
			res := make(map[*TypedText][]*Encoding)
			for encoding := range stage.Encodings {
				for _, typedtext_ := range encoding.Encoder {
					res[typedtext_] = append(res[typedtext_], encoding)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Supports":
			res := make(map[*Supports][]*Encoding)
			for encoding := range stage.Encodings {
				for _, supports_ := range encoding.Supports {
					res[supports_] = append(res[supports_], encoding)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Fermata
	case Fermata:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Fingering
	case Fingering:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormattedSymbol
	case FormattedSymbol:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormattedSymbolId
	case FormattedSymbolId:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormattedText
	case FormattedText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormattedTextId
	case FormattedTextId:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Fret
	case Fret:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Glyph
	case Glyph:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of HorizontalTurn
	case HorizontalTurn:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Identification
	case Identification:
		switch fieldname {
		// insertion point for per direct association field
		case "Creator":
			res := make(map[*TypedText][]*Identification)
			for identification := range stage.Identifications {
				for _, typedtext_ := range identification.Creator {
					res[typedtext_] = append(res[typedtext_], identification)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Rights":
			res := make(map[*TypedText][]*Identification)
			for identification := range stage.Identifications {
				for _, typedtext_ := range identification.Rights {
					res[typedtext_] = append(res[typedtext_], identification)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Relation":
			res := make(map[*TypedText][]*Identification)
			for identification := range stage.Identifications {
				for _, typedtext_ := range identification.Relation {
					res[typedtext_] = append(res[typedtext_], identification)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Level
	case Level:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LineWidth
	case LineWidth:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MeasureLayout
	case MeasureLayout:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MidiDevice
	case MidiDevice:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MidiInstrument
	case MidiInstrument:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Miscellaneous
	case Miscellaneous:
		switch fieldname {
		// insertion point for per direct association field
		case "MiscellaneousField":
			res := make(map[*MiscellaneousField][]*Miscellaneous)
			for miscellaneous := range stage.Miscellaneouss {
				for _, miscellaneousfield_ := range miscellaneous.MiscellaneousField {
					res[miscellaneousfield_] = append(res[miscellaneousfield_], miscellaneous)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MiscellaneousField
	case MiscellaneousField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NameDisplay
	case NameDisplay:
		switch fieldname {
		// insertion point for per direct association field
		case "DisplayText":
			res := make(map[*FormattedText][]*NameDisplay)
			for namedisplay := range stage.NameDisplays {
				for _, formattedtext_ := range namedisplay.DisplayText {
					res[formattedtext_] = append(res[formattedtext_], namedisplay)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AccidentalText":
			res := make(map[*AccidentalText][]*NameDisplay)
			for namedisplay := range stage.NameDisplays {
				for _, accidentaltext_ := range namedisplay.AccidentalText {
					res[accidentaltext_] = append(res[accidentaltext_], namedisplay)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteSize
	case NoteSize:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherAppearance
	case OtherAppearance:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherPlay
	case OtherPlay:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of OtherText
	case OtherText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PageLayout
	case PageLayout:
		switch fieldname {
		// insertion point for per direct association field
		case "PageMargins":
			res := make(map[*PageMargins][]*PageLayout)
			for pagelayout := range stage.PageLayouts {
				for _, pagemargins_ := range pagelayout.PageMargins {
					res[pagemargins_] = append(res[pagemargins_], pagelayout)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PageMargins
	case PageMargins:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Play
	case Play:
		switch fieldname {
		// insertion point for per direct association field
		case "OtherPlay":
			res := make(map[*OtherPlay][]*Play)
			for play := range stage.Plays {
				for _, otherplay_ := range play.OtherPlay {
					res[otherplay_] = append(res[otherplay_], play)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Scaling
	case Scaling:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Segno
	case Segno:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StaffLayout
	case StaffLayout:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StringType
	case StringType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Supports
	case Supports:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SystemDividers
	case SystemDividers:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SystemLayout
	case SystemLayout:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SystemMargins
	case SystemMargins:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TypedText
	case TypedText:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of WavyLine
	case WavyLine:
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
func (*AccidentalText) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "NameDisplay",
			Fieldname: "AccidentalText",
		},
	}
}

func (*Appearance) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Bookmark) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Coda) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Distance) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Appearance",
			Fieldname: "Distance",
		},
	}
}

func (*Dynamics) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Empty) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Dynamics",
			Fieldname: "P",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Pp",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Ppp",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Pppp",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Ppppp",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Pppppp",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "F",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Ff",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Fff",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Ffff",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Fffff",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Ffffff",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Mp",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Mf",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Sf",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Sfp",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Sfpp",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Fp",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Rf",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Rfz",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Sfz",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Sffz",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Fz",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "N",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Pf",
		},
		{
			GongstructName: "Dynamics",
			Fieldname: "Sfzp",
		},
	}
}

func (*EmptyPlacement) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*EmptyPlacementSmufl) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*EmptyPrintObjectStyleAlign) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*EmptyPrintStyle) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*EmptyPrintStyleAlign) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*EmptyPrintStyleAlignId) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*EmptyTrillSound) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Encoding) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Fermata) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Fingering) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*FormattedSymbol) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*FormattedSymbolId) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*FormattedText) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "NameDisplay",
			Fieldname: "DisplayText",
		},
	}
}

func (*FormattedTextId) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Fret) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Glyph) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Appearance",
			Fieldname: "Glyph",
		},
	}
}

func (*HorizontalTurn) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Identification) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Level) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*LineWidth) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Appearance",
			Fieldname: "LineWidth",
		},
	}
}

func (*Link) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MeasureLayout) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MidiDevice) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MidiInstrument) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Miscellaneous) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*MiscellaneousField) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Miscellaneous",
			Fieldname: "MiscellaneousField",
		},
	}
}

func (*NameDisplay) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*NoteSize) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Appearance",
			Fieldname: "NoteSize",
		},
	}
}

func (*OtherAppearance) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Appearance",
			Fieldname: "OtherAppearance",
		},
	}
}

func (*OtherPlay) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Play",
			Fieldname: "OtherPlay",
		},
	}
}

func (*OtherText) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Dynamics",
			Fieldname: "OtherDynamics",
		},
	}
}

func (*PageLayout) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*PageMargins) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "PageLayout",
			Fieldname: "PageMargins",
		},
	}
}

func (*Play) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Scaling) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Segno) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*StaffLayout) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*StringType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*Supports) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Encoding",
			Fieldname: "Supports",
		},
	}
}

func (*SystemDividers) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*SystemLayout) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*SystemMargins) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*TypedText) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "Encoding",
			Fieldname: "Encoder",
		},
		{
			GongstructName: "Identification",
			Fieldname: "Creator",
		},
		{
			GongstructName: "Identification",
			Fieldname: "Rights",
		},
		{
			GongstructName: "Identification",
			Fieldname: "Relation",
		},
	}
}

func (*WavyLine) GongGetReverseFields() []GongReverseField {
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
func (accidentaltext *AccidentalText) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Lang",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Space",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Justify",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
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
			Name:               "LineHeight",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Dir",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTextDirection",
		},
		{
			Name:                 "Enclosure",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumEnclosureShape",
		},
		{
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumAccidentalValue",
		},
	}
	return
}

func (appearance *Appearance) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "LineWidth",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "LineWidth",
		},
		{
			Name:                 "NoteSize",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteSize",
		},
		{
			Name:                 "Distance",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Distance",
		},
		{
			Name:                 "Glyph",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Glyph",
		},
		{
			Name:                 "OtherAppearance",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OtherAppearance",
		},
	}
	return
}

func (bookmark *Bookmark) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "NameAttr",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Element",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Position",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (coda *Coda) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (distance *Distance) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (dynamics *Dynamics) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
		},
		{
			Name:                 "Placement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumAboveBelow",
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
			Name:                 "Enclosure",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumEnclosureShape",
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "P",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Pp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Ppp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Pppp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Ppppp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Pppppp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "F",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Ff",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Fff",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Ffff",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Fffff",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Ffffff",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Mp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Mf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Sf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Sfp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Sfpp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Fp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Rf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Rfz",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Sfz",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Sffz",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Fz",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "N",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Pf",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "Sfzp",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Empty",
		},
		{
			Name:                 "OtherDynamics",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OtherText",
		},
	}
	return
}

func (empty *Empty) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (emptyplacement *EmptyPlacement) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Placement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumAboveBelow",
		},
	}
	return
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Placement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumAboveBelow",
		},
		{
			Name:               "Smufl",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PrintObject",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
		},
	}
	return
}

func (emptyprintstyle *EmptyPrintStyle) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
		},
	}
	return
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (emptytrillsound *EmptyTrillSound) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Placement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumAboveBelow",
		},
		{
			Name:                 "StartNote",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStartNote",
		},
		{
			Name:                 "TrillStep",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTrillStep",
		},
		{
			Name:                 "TwoNoteTurn",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTwoNoteTurn",
		},
		{
			Name:                 "Accelerate",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
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

func (encoding *Encoding) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EncodingDate",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Encoder",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TypedText",
		},
		{
			Name:               "Software",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EncodingDescription",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Supports",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Supports",
		},
	}
	return
}

func (fermata *Fermata) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumUprightInverted",
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
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
			Name:                 "Value",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFermataShape",
		},
	}
	return
}

func (fingering *Fingering) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Substitution",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
		},
		{
			Name:                 "Alternate",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Placement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumAboveBelow",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (formattedsymbol *FormattedSymbol) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Justify",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
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
			Name:               "LineHeight",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Dir",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTextDirection",
		},
		{
			Name:                 "Enclosure",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumEnclosureShape",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (formattedsymbolid *FormattedSymbolId) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Justify",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
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
			Name:               "LineHeight",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Dir",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTextDirection",
		},
		{
			Name:                 "Enclosure",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumEnclosureShape",
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

func (formattedtext *FormattedText) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Space",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Justify",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
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
			Name:               "LineHeight",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Dir",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTextDirection",
		},
		{
			Name:                 "Enclosure",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumEnclosureShape",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (formattedtextid *FormattedTextId) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Space",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Justify",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
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
			Name:               "LineHeight",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Dir",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTextDirection",
		},
		{
			Name:                 "Enclosure",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumEnclosureShape",
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

func (fret *Fret) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
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

func (glyph *Glyph) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (horizontalturn *HorizontalTurn) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Slash",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Placement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumAboveBelow",
		},
		{
			Name:                 "StartNote",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStartNote",
		},
		{
			Name:                 "TrillStep",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTrillStep",
		},
		{
			Name:                 "TwoNoteTurn",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTwoNoteTurn",
		},
		{
			Name:                 "Accelerate",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
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

func (identification *Identification) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Creator",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TypedText",
		},
		{
			Name:                 "Rights",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TypedText",
		},
		{
			Name:                 "Encoding",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Encoding",
		},
		{
			Name:               "Source",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Relation",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TypedText",
		},
		{
			Name:                 "Miscellaneous",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Miscellaneous",
		},
	}
	return
}

func (level *Level) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Reference",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStartStopSingle",
		},
		{
			Name:                 "Parentheses",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
		},
		{
			Name:                 "Bracket",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
		},
		{
			Name:                 "Size",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumSymbolSize",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (linewidth *LineWidth) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (link *Link) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NameAttr",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Href",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Role",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Element",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Position",
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
	}
	return
}

func (measurelayout *MeasureLayout) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MeasureDistance",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (mididevice *MidiDevice) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Port",
			GongFieldValueType: GongFieldValueTypeInt,
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

func (midiinstrument *MidiInstrument) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "MidiChannel",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MidiName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MidiBank",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MidiProgram",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MidiUnpitched",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Volume",
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
	}
	return
}

func (miscellaneous *Miscellaneous) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "MiscellaneousField",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MiscellaneousField",
		},
	}
	return
}

func (miscellaneousfield *MiscellaneousField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NameAttr",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (namedisplay *NameDisplay) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PrintObject",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
		},
		{
			Name:                 "DisplayText",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "FormattedText",
		},
		{
			Name:                 "AccidentalText",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AccidentalText",
		},
	}
	return
}

func (notesize *NoteSize) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumNoteSizeType",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (otherappearance *OtherAppearance) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (otherplay *OtherPlay) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (othertext *OtherText) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (pagelayout *PageLayout) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PageHeight",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "PageWidth",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "PageMargins",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PageMargins",
		},
	}
	return
}

func (pagemargins *PageMargins) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumMarginType",
		},
		{
			Name:               "LeftMargin",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RightMargin",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TopMargin",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "BottomMargin",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (play *Play) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Ipa",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Mute",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumMute",
		},
		{
			Name:                 "SemiPitched",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumSemiPitched",
		},
		{
			Name:                 "OtherPlay",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "OtherPlay",
		},
	}
	return
}

func (scaling *Scaling) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Millimeters",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Tenths",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (segno *Segno) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Halign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumLeftCenterRight",
		},
		{
			Name:                 "Valign",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumValign",
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (stafflayout *StaffLayout) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "StaffDistance",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (stringtype *StringType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "FontStyle",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontStyle",
		},
		{
			Name:               "FontSize",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FontWeight",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumFontWeight",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Placement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumAboveBelow",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (supports *Supports) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
		},
		{
			Name:               "Element",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Attribute",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (systemdividers *SystemDividers) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "LeftDivider",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EmptyPrintObjectStyleAlign",
		},
		{
			Name:                 "RightDivider",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "EmptyPrintObjectStyleAlign",
		},
	}
	return
}

func (systemlayout *SystemLayout) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SystemMargins",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SystemMargins",
		},
		{
			Name:               "SystemDistance",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TopSystemDistance",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SystemDividers",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SystemDividers",
		},
	}
	return
}

func (systemmargins *SystemMargins) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LeftMargin",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RightMargin",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (typedtext *TypedText) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (wavyline *WavyLine) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStartStopContinue",
		},
		{
			Name:               "Number",
			GongFieldValueType: GongFieldValueTypeInt,
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
			Name:                 "Placement",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumAboveBelow",
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "StartNote",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumStartNote",
		},
		{
			Name:                 "TrillStep",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTrillStep",
		},
		{
			Name:                 "TwoNoteTurn",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumTwoNoteTurn",
		},
		{
			Name:                 "Accelerate",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EnumYesNo",
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
func (accidentaltext *AccidentalText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = accidentaltext.Name
	case "Smufl":
		res.valueString = accidentaltext.Smufl
	case "Lang":
		res.valueString = accidentaltext.Lang
	case "Space":
		res.valueString = accidentaltext.Space
	case "Justify":
		enum := accidentaltext.Justify
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = accidentaltext.DefaultX
	case "DefaultY":
		res.valueString = accidentaltext.DefaultY
	case "RelativeX":
		res.valueString = accidentaltext.RelativeX
	case "RelativeY":
		res.valueString = accidentaltext.RelativeY
	case "FontFamily":
		res.valueString = accidentaltext.FontFamily
	case "FontStyle":
		enum := accidentaltext.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = accidentaltext.FontSize
	case "FontWeight":
		enum := accidentaltext.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = accidentaltext.Color
	case "Halign":
		enum := accidentaltext.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := accidentaltext.Valign
		res.valueString = enum.ToCodeString()
	case "Underline":
		res.valueString = fmt.Sprintf("%d", accidentaltext.Underline)
		res.valueInt = accidentaltext.Underline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Overline":
		res.valueString = fmt.Sprintf("%d", accidentaltext.Overline)
		res.valueInt = accidentaltext.Overline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "LineThrough":
		res.valueString = fmt.Sprintf("%d", accidentaltext.LineThrough)
		res.valueInt = accidentaltext.LineThrough
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Rotation":
		res.valueString = accidentaltext.Rotation
	case "LetterSpacing":
		res.valueString = accidentaltext.LetterSpacing
	case "LineHeight":
		res.valueString = accidentaltext.LineHeight
	case "Dir":
		enum := accidentaltext.Dir
		res.valueString = enum.ToCodeString()
	case "Enclosure":
		enum := accidentaltext.Enclosure
		res.valueString = enum.ToCodeString()
	case "Value":
		enum := accidentaltext.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (appearance *Appearance) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = appearance.Name
	case "LineWidth":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range appearance.LineWidth {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NoteSize":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range appearance.NoteSize {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Distance":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range appearance.Distance {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Glyph":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range appearance.Glyph {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "OtherAppearance":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range appearance.OtherAppearance {
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

func (bookmark *Bookmark) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bookmark.Name
	case "Id":
		res.valueString = bookmark.Id
	case "NameAttr":
		res.valueString = bookmark.NameAttr
	case "Element":
		res.valueString = bookmark.Element
	case "Position":
		res.valueString = fmt.Sprintf("%d", bookmark.Position)
		res.valueInt = bookmark.Position
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (coda *Coda) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = coda.Name
	case "Smufl":
		res.valueString = coda.Smufl
	case "DefaultX":
		res.valueString = coda.DefaultX
	case "DefaultY":
		res.valueString = coda.DefaultY
	case "RelativeX":
		res.valueString = coda.RelativeX
	case "RelativeY":
		res.valueString = coda.RelativeY
	case "FontFamily":
		res.valueString = coda.FontFamily
	case "FontStyle":
		enum := coda.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = coda.FontSize
	case "FontWeight":
		enum := coda.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = coda.Color
	case "Halign":
		enum := coda.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := coda.Valign
		res.valueString = enum.ToCodeString()
	case "Id":
		res.valueString = coda.Id
	}
	return
}

func (distance *Distance) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = distance.Name
	case "Type":
		res.valueString = distance.Type
	case "Value":
		res.valueString = distance.Value
	}
	return
}

func (dynamics *Dynamics) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dynamics.Name
	case "DefaultX":
		res.valueString = dynamics.DefaultX
	case "DefaultY":
		res.valueString = dynamics.DefaultY
	case "RelativeX":
		res.valueString = dynamics.RelativeX
	case "RelativeY":
		res.valueString = dynamics.RelativeY
	case "FontFamily":
		res.valueString = dynamics.FontFamily
	case "FontStyle":
		enum := dynamics.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = dynamics.FontSize
	case "FontWeight":
		enum := dynamics.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = dynamics.Color
	case "Halign":
		enum := dynamics.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := dynamics.Valign
		res.valueString = enum.ToCodeString()
	case "Placement":
		enum := dynamics.Placement
		res.valueString = enum.ToCodeString()
	case "Underline":
		res.valueString = fmt.Sprintf("%d", dynamics.Underline)
		res.valueInt = dynamics.Underline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Overline":
		res.valueString = fmt.Sprintf("%d", dynamics.Overline)
		res.valueInt = dynamics.Overline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "LineThrough":
		res.valueString = fmt.Sprintf("%d", dynamics.LineThrough)
		res.valueInt = dynamics.LineThrough
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Enclosure":
		enum := dynamics.Enclosure
		res.valueString = enum.ToCodeString()
	case "Id":
		res.valueString = dynamics.Id
	case "P":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.P {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Pp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Pp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Ppp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Ppp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Pppp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Pppp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Ppppp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Ppppp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Pppppp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Pppppp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "F":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.F {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Ff":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Ff {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Fff":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Fff {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Ffff":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Ffff {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Fffff":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Fffff {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Ffffff":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Ffffff {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Mp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Mp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Mf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Mf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Sf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Sf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Sfp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Sfp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Sfpp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Sfpp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Fp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Fp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Rf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Rf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Rfz":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Rfz {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Sfz":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Sfz {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Sffz":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Sffz {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Fz":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Fz {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "N":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.N {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Pf":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Pf {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Sfzp":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.Sfzp {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "OtherDynamics":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dynamics.OtherDynamics {
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

func (empty *Empty) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = empty.Name
	}
	return
}

func (emptyplacement *EmptyPlacement) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = emptyplacement.Name
	case "DefaultX":
		res.valueString = emptyplacement.DefaultX
	case "DefaultY":
		res.valueString = emptyplacement.DefaultY
	case "RelativeX":
		res.valueString = emptyplacement.RelativeX
	case "RelativeY":
		res.valueString = emptyplacement.RelativeY
	case "FontFamily":
		res.valueString = emptyplacement.FontFamily
	case "FontStyle":
		enum := emptyplacement.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = emptyplacement.FontSize
	case "FontWeight":
		enum := emptyplacement.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = emptyplacement.Color
	case "Placement":
		enum := emptyplacement.Placement
		res.valueString = enum.ToCodeString()
	}
	return
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = emptyplacementsmufl.Name
	case "DefaultX":
		res.valueString = emptyplacementsmufl.DefaultX
	case "DefaultY":
		res.valueString = emptyplacementsmufl.DefaultY
	case "RelativeX":
		res.valueString = emptyplacementsmufl.RelativeX
	case "RelativeY":
		res.valueString = emptyplacementsmufl.RelativeY
	case "FontFamily":
		res.valueString = emptyplacementsmufl.FontFamily
	case "FontStyle":
		enum := emptyplacementsmufl.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = emptyplacementsmufl.FontSize
	case "FontWeight":
		enum := emptyplacementsmufl.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = emptyplacementsmufl.Color
	case "Placement":
		enum := emptyplacementsmufl.Placement
		res.valueString = enum.ToCodeString()
	case "Smufl":
		res.valueString = emptyplacementsmufl.Smufl
	}
	return
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = emptyprintobjectstylealign.Name
	case "PrintObject":
		enum := emptyprintobjectstylealign.PrintObject
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = emptyprintobjectstylealign.DefaultX
	case "DefaultY":
		res.valueString = emptyprintobjectstylealign.DefaultY
	case "RelativeX":
		res.valueString = emptyprintobjectstylealign.RelativeX
	case "RelativeY":
		res.valueString = emptyprintobjectstylealign.RelativeY
	case "FontFamily":
		res.valueString = emptyprintobjectstylealign.FontFamily
	case "FontStyle":
		enum := emptyprintobjectstylealign.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = emptyprintobjectstylealign.FontSize
	case "FontWeight":
		enum := emptyprintobjectstylealign.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = emptyprintobjectstylealign.Color
	case "Halign":
		enum := emptyprintobjectstylealign.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := emptyprintobjectstylealign.Valign
		res.valueString = enum.ToCodeString()
	}
	return
}

func (emptyprintstyle *EmptyPrintStyle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = emptyprintstyle.Name
	case "DefaultX":
		res.valueString = emptyprintstyle.DefaultX
	case "DefaultY":
		res.valueString = emptyprintstyle.DefaultY
	case "RelativeX":
		res.valueString = emptyprintstyle.RelativeX
	case "RelativeY":
		res.valueString = emptyprintstyle.RelativeY
	case "FontFamily":
		res.valueString = emptyprintstyle.FontFamily
	case "FontStyle":
		enum := emptyprintstyle.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = emptyprintstyle.FontSize
	case "FontWeight":
		enum := emptyprintstyle.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = emptyprintstyle.Color
	}
	return
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = emptyprintstylealign.Name
	case "DefaultX":
		res.valueString = emptyprintstylealign.DefaultX
	case "DefaultY":
		res.valueString = emptyprintstylealign.DefaultY
	case "RelativeX":
		res.valueString = emptyprintstylealign.RelativeX
	case "RelativeY":
		res.valueString = emptyprintstylealign.RelativeY
	case "FontFamily":
		res.valueString = emptyprintstylealign.FontFamily
	case "FontStyle":
		enum := emptyprintstylealign.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = emptyprintstylealign.FontSize
	case "FontWeight":
		enum := emptyprintstylealign.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = emptyprintstylealign.Color
	case "Halign":
		enum := emptyprintstylealign.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := emptyprintstylealign.Valign
		res.valueString = enum.ToCodeString()
	}
	return
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = emptyprintstylealignid.Name
	case "DefaultX":
		res.valueString = emptyprintstylealignid.DefaultX
	case "DefaultY":
		res.valueString = emptyprintstylealignid.DefaultY
	case "RelativeX":
		res.valueString = emptyprintstylealignid.RelativeX
	case "RelativeY":
		res.valueString = emptyprintstylealignid.RelativeY
	case "FontFamily":
		res.valueString = emptyprintstylealignid.FontFamily
	case "FontStyle":
		enum := emptyprintstylealignid.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = emptyprintstylealignid.FontSize
	case "FontWeight":
		enum := emptyprintstylealignid.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = emptyprintstylealignid.Color
	case "Halign":
		enum := emptyprintstylealignid.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := emptyprintstylealignid.Valign
		res.valueString = enum.ToCodeString()
	case "Id":
		res.valueString = emptyprintstylealignid.Id
	}
	return
}

func (emptytrillsound *EmptyTrillSound) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = emptytrillsound.Name
	case "DefaultX":
		res.valueString = emptytrillsound.DefaultX
	case "DefaultY":
		res.valueString = emptytrillsound.DefaultY
	case "RelativeX":
		res.valueString = emptytrillsound.RelativeX
	case "RelativeY":
		res.valueString = emptytrillsound.RelativeY
	case "FontFamily":
		res.valueString = emptytrillsound.FontFamily
	case "FontStyle":
		enum := emptytrillsound.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = emptytrillsound.FontSize
	case "FontWeight":
		enum := emptytrillsound.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = emptytrillsound.Color
	case "Placement":
		enum := emptytrillsound.Placement
		res.valueString = enum.ToCodeString()
	case "StartNote":
		enum := emptytrillsound.StartNote
		res.valueString = enum.ToCodeString()
	case "TrillStep":
		enum := emptytrillsound.TrillStep
		res.valueString = enum.ToCodeString()
	case "TwoNoteTurn":
		enum := emptytrillsound.TwoNoteTurn
		res.valueString = enum.ToCodeString()
	case "Accelerate":
		enum := emptytrillsound.Accelerate
		res.valueString = enum.ToCodeString()
	case "Beats":
		res.valueString = emptytrillsound.Beats
	case "SecondBeat":
		res.valueString = emptytrillsound.SecondBeat
	case "LastBeat":
		res.valueString = emptytrillsound.LastBeat
	}
	return
}

func (encoding *Encoding) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = encoding.Name
	case "EncodingDate":
		res.valueString = encoding.EncodingDate
	case "Encoder":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range encoding.Encoder {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Software":
		res.valueString = encoding.Software
	case "EncodingDescription":
		res.valueString = encoding.EncodingDescription
	case "Supports":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range encoding.Supports {
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

func (fermata *Fermata) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = fermata.Name
	case "Type":
		enum := fermata.Type
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = fermata.DefaultX
	case "DefaultY":
		res.valueString = fermata.DefaultY
	case "RelativeX":
		res.valueString = fermata.RelativeX
	case "RelativeY":
		res.valueString = fermata.RelativeY
	case "FontFamily":
		res.valueString = fermata.FontFamily
	case "FontStyle":
		enum := fermata.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = fermata.FontSize
	case "FontWeight":
		enum := fermata.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = fermata.Color
	case "Id":
		res.valueString = fermata.Id
	case "Value":
		enum := fermata.Value
		res.valueString = enum.ToCodeString()
	}
	return
}

func (fingering *Fingering) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = fingering.Name
	case "Substitution":
		enum := fingering.Substitution
		res.valueString = enum.ToCodeString()
	case "Alternate":
		enum := fingering.Alternate
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = fingering.DefaultX
	case "DefaultY":
		res.valueString = fingering.DefaultY
	case "RelativeX":
		res.valueString = fingering.RelativeX
	case "RelativeY":
		res.valueString = fingering.RelativeY
	case "FontFamily":
		res.valueString = fingering.FontFamily
	case "FontStyle":
		enum := fingering.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = fingering.FontSize
	case "FontWeight":
		enum := fingering.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = fingering.Color
	case "Placement":
		enum := fingering.Placement
		res.valueString = enum.ToCodeString()
	case "Value":
		res.valueString = fingering.Value
	}
	return
}

func (formattedsymbol *FormattedSymbol) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formattedsymbol.Name
	case "Justify":
		enum := formattedsymbol.Justify
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = formattedsymbol.DefaultX
	case "DefaultY":
		res.valueString = formattedsymbol.DefaultY
	case "RelativeX":
		res.valueString = formattedsymbol.RelativeX
	case "RelativeY":
		res.valueString = formattedsymbol.RelativeY
	case "FontFamily":
		res.valueString = formattedsymbol.FontFamily
	case "FontStyle":
		enum := formattedsymbol.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = formattedsymbol.FontSize
	case "FontWeight":
		enum := formattedsymbol.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = formattedsymbol.Color
	case "Halign":
		enum := formattedsymbol.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := formattedsymbol.Valign
		res.valueString = enum.ToCodeString()
	case "Underline":
		res.valueString = fmt.Sprintf("%d", formattedsymbol.Underline)
		res.valueInt = formattedsymbol.Underline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Overline":
		res.valueString = fmt.Sprintf("%d", formattedsymbol.Overline)
		res.valueInt = formattedsymbol.Overline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "LineThrough":
		res.valueString = fmt.Sprintf("%d", formattedsymbol.LineThrough)
		res.valueInt = formattedsymbol.LineThrough
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Rotation":
		res.valueString = formattedsymbol.Rotation
	case "LetterSpacing":
		res.valueString = formattedsymbol.LetterSpacing
	case "LineHeight":
		res.valueString = formattedsymbol.LineHeight
	case "Dir":
		enum := formattedsymbol.Dir
		res.valueString = enum.ToCodeString()
	case "Enclosure":
		enum := formattedsymbol.Enclosure
		res.valueString = enum.ToCodeString()
	case "Value":
		res.valueString = formattedsymbol.Value
	}
	return
}

func (formattedsymbolid *FormattedSymbolId) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formattedsymbolid.Name
	case "Justify":
		enum := formattedsymbolid.Justify
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = formattedsymbolid.DefaultX
	case "DefaultY":
		res.valueString = formattedsymbolid.DefaultY
	case "RelativeX":
		res.valueString = formattedsymbolid.RelativeX
	case "RelativeY":
		res.valueString = formattedsymbolid.RelativeY
	case "FontFamily":
		res.valueString = formattedsymbolid.FontFamily
	case "FontStyle":
		enum := formattedsymbolid.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = formattedsymbolid.FontSize
	case "FontWeight":
		enum := formattedsymbolid.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = formattedsymbolid.Color
	case "Halign":
		enum := formattedsymbolid.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := formattedsymbolid.Valign
		res.valueString = enum.ToCodeString()
	case "Underline":
		res.valueString = fmt.Sprintf("%d", formattedsymbolid.Underline)
		res.valueInt = formattedsymbolid.Underline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Overline":
		res.valueString = fmt.Sprintf("%d", formattedsymbolid.Overline)
		res.valueInt = formattedsymbolid.Overline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "LineThrough":
		res.valueString = fmt.Sprintf("%d", formattedsymbolid.LineThrough)
		res.valueInt = formattedsymbolid.LineThrough
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Rotation":
		res.valueString = formattedsymbolid.Rotation
	case "LetterSpacing":
		res.valueString = formattedsymbolid.LetterSpacing
	case "LineHeight":
		res.valueString = formattedsymbolid.LineHeight
	case "Dir":
		enum := formattedsymbolid.Dir
		res.valueString = enum.ToCodeString()
	case "Enclosure":
		enum := formattedsymbolid.Enclosure
		res.valueString = enum.ToCodeString()
	case "Id":
		res.valueString = formattedsymbolid.Id
	case "Value":
		res.valueString = formattedsymbolid.Value
	}
	return
}

func (formattedtext *FormattedText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formattedtext.Name
	case "Lang":
		res.valueString = formattedtext.Lang
	case "Space":
		res.valueString = formattedtext.Space
	case "Justify":
		enum := formattedtext.Justify
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = formattedtext.DefaultX
	case "DefaultY":
		res.valueString = formattedtext.DefaultY
	case "RelativeX":
		res.valueString = formattedtext.RelativeX
	case "RelativeY":
		res.valueString = formattedtext.RelativeY
	case "FontFamily":
		res.valueString = formattedtext.FontFamily
	case "FontStyle":
		enum := formattedtext.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = formattedtext.FontSize
	case "FontWeight":
		enum := formattedtext.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = formattedtext.Color
	case "Halign":
		enum := formattedtext.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := formattedtext.Valign
		res.valueString = enum.ToCodeString()
	case "Underline":
		res.valueString = fmt.Sprintf("%d", formattedtext.Underline)
		res.valueInt = formattedtext.Underline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Overline":
		res.valueString = fmt.Sprintf("%d", formattedtext.Overline)
		res.valueInt = formattedtext.Overline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "LineThrough":
		res.valueString = fmt.Sprintf("%d", formattedtext.LineThrough)
		res.valueInt = formattedtext.LineThrough
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Rotation":
		res.valueString = formattedtext.Rotation
	case "LetterSpacing":
		res.valueString = formattedtext.LetterSpacing
	case "LineHeight":
		res.valueString = formattedtext.LineHeight
	case "Dir":
		enum := formattedtext.Dir
		res.valueString = enum.ToCodeString()
	case "Enclosure":
		enum := formattedtext.Enclosure
		res.valueString = enum.ToCodeString()
	case "Value":
		res.valueString = formattedtext.Value
	}
	return
}

func (formattedtextid *FormattedTextId) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = formattedtextid.Name
	case "Lang":
		res.valueString = formattedtextid.Lang
	case "Space":
		res.valueString = formattedtextid.Space
	case "Justify":
		enum := formattedtextid.Justify
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = formattedtextid.DefaultX
	case "DefaultY":
		res.valueString = formattedtextid.DefaultY
	case "RelativeX":
		res.valueString = formattedtextid.RelativeX
	case "RelativeY":
		res.valueString = formattedtextid.RelativeY
	case "FontFamily":
		res.valueString = formattedtextid.FontFamily
	case "FontStyle":
		enum := formattedtextid.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = formattedtextid.FontSize
	case "FontWeight":
		enum := formattedtextid.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = formattedtextid.Color
	case "Halign":
		enum := formattedtextid.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := formattedtextid.Valign
		res.valueString = enum.ToCodeString()
	case "Underline":
		res.valueString = fmt.Sprintf("%d", formattedtextid.Underline)
		res.valueInt = formattedtextid.Underline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Overline":
		res.valueString = fmt.Sprintf("%d", formattedtextid.Overline)
		res.valueInt = formattedtextid.Overline
		res.GongFieldValueType = GongFieldValueTypeInt
	case "LineThrough":
		res.valueString = fmt.Sprintf("%d", formattedtextid.LineThrough)
		res.valueInt = formattedtextid.LineThrough
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Rotation":
		res.valueString = formattedtextid.Rotation
	case "LetterSpacing":
		res.valueString = formattedtextid.LetterSpacing
	case "LineHeight":
		res.valueString = formattedtextid.LineHeight
	case "Dir":
		enum := formattedtextid.Dir
		res.valueString = enum.ToCodeString()
	case "Enclosure":
		enum := formattedtextid.Enclosure
		res.valueString = enum.ToCodeString()
	case "Id":
		res.valueString = formattedtextid.Id
	case "Value":
		res.valueString = formattedtextid.Value
	}
	return
}

func (fret *Fret) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = fret.Name
	case "FontFamily":
		res.valueString = fret.FontFamily
	case "FontStyle":
		enum := fret.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = fret.FontSize
	case "FontWeight":
		enum := fret.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = fret.Color
	case "Value":
		res.valueString = fmt.Sprintf("%d", fret.Value)
		res.valueInt = fret.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (glyph *Glyph) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = glyph.Name
	case "Type":
		res.valueString = glyph.Type
	case "Value":
		res.valueString = glyph.Value
	}
	return
}

func (horizontalturn *HorizontalTurn) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = horizontalturn.Name
	case "Slash":
		enum := horizontalturn.Slash
		res.valueString = enum.ToCodeString()
	case "DefaultX":
		res.valueString = horizontalturn.DefaultX
	case "DefaultY":
		res.valueString = horizontalturn.DefaultY
	case "RelativeX":
		res.valueString = horizontalturn.RelativeX
	case "RelativeY":
		res.valueString = horizontalturn.RelativeY
	case "FontFamily":
		res.valueString = horizontalturn.FontFamily
	case "FontStyle":
		enum := horizontalturn.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = horizontalturn.FontSize
	case "FontWeight":
		enum := horizontalturn.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = horizontalturn.Color
	case "Placement":
		enum := horizontalturn.Placement
		res.valueString = enum.ToCodeString()
	case "StartNote":
		enum := horizontalturn.StartNote
		res.valueString = enum.ToCodeString()
	case "TrillStep":
		enum := horizontalturn.TrillStep
		res.valueString = enum.ToCodeString()
	case "TwoNoteTurn":
		enum := horizontalturn.TwoNoteTurn
		res.valueString = enum.ToCodeString()
	case "Accelerate":
		enum := horizontalturn.Accelerate
		res.valueString = enum.ToCodeString()
	case "Beats":
		res.valueString = horizontalturn.Beats
	case "SecondBeat":
		res.valueString = horizontalturn.SecondBeat
	case "LastBeat":
		res.valueString = horizontalturn.LastBeat
	}
	return
}

func (identification *Identification) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = identification.Name
	case "Creator":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range identification.Creator {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Rights":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range identification.Rights {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Encoding":
		res.GongFieldValueType = GongFieldValueTypePointer
		if identification.Encoding != nil {
			res.valueString = identification.Encoding.Name
			res.ids = identification.Encoding.GongGetUUID(stage)
		}
	case "Source":
		res.valueString = identification.Source
	case "Relation":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range identification.Relation {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Miscellaneous":
		res.GongFieldValueType = GongFieldValueTypePointer
		if identification.Miscellaneous != nil {
			res.valueString = identification.Miscellaneous.Name
			res.ids = identification.Miscellaneous.GongGetUUID(stage)
		}
	}
	return
}

func (level *Level) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = level.Name
	case "Reference":
		enum := level.Reference
		res.valueString = enum.ToCodeString()
	case "Type":
		enum := level.Type
		res.valueString = enum.ToCodeString()
	case "Parentheses":
		enum := level.Parentheses
		res.valueString = enum.ToCodeString()
	case "Bracket":
		enum := level.Bracket
		res.valueString = enum.ToCodeString()
	case "Size":
		enum := level.Size
		res.valueString = enum.ToCodeString()
	case "Value":
		res.valueString = level.Value
	}
	return
}

func (linewidth *LineWidth) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = linewidth.Name
	case "Type":
		res.valueString = linewidth.Type
	case "Value":
		res.valueString = linewidth.Value
	}
	return
}

func (link *Link) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = link.Name
	case "NameAttr":
		res.valueString = link.NameAttr
	case "Href":
		res.valueString = link.Href
	case "Role":
		res.valueString = link.Role
	case "Title":
		res.valueString = link.Title
	case "Element":
		res.valueString = link.Element
	case "Position":
		res.valueString = fmt.Sprintf("%d", link.Position)
		res.valueInt = link.Position
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DefaultX":
		res.valueString = link.DefaultX
	case "DefaultY":
		res.valueString = link.DefaultY
	case "RelativeX":
		res.valueString = link.RelativeX
	case "RelativeY":
		res.valueString = link.RelativeY
	}
	return
}

func (measurelayout *MeasureLayout) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = measurelayout.Name
	case "MeasureDistance":
		res.valueString = measurelayout.MeasureDistance
	}
	return
}

func (mididevice *MidiDevice) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = mididevice.Name
	case "Port":
		res.valueString = fmt.Sprintf("%d", mididevice.Port)
		res.valueInt = mididevice.Port
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Id":
		res.valueString = mididevice.Id
	case "Value":
		res.valueString = mididevice.Value
	}
	return
}

func (midiinstrument *MidiInstrument) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = midiinstrument.Name
	case "Id":
		res.valueString = midiinstrument.Id
	case "MidiChannel":
		res.valueString = fmt.Sprintf("%d", midiinstrument.MidiChannel)
		res.valueInt = midiinstrument.MidiChannel
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MidiName":
		res.valueString = midiinstrument.MidiName
	case "MidiBank":
		res.valueString = fmt.Sprintf("%d", midiinstrument.MidiBank)
		res.valueInt = midiinstrument.MidiBank
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MidiProgram":
		res.valueString = fmt.Sprintf("%d", midiinstrument.MidiProgram)
		res.valueInt = midiinstrument.MidiProgram
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MidiUnpitched":
		res.valueString = fmt.Sprintf("%d", midiinstrument.MidiUnpitched)
		res.valueInt = midiinstrument.MidiUnpitched
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Volume":
		res.valueString = midiinstrument.Volume
	case "Pan":
		res.valueString = midiinstrument.Pan
	case "Elevation":
		res.valueString = midiinstrument.Elevation
	}
	return
}

func (miscellaneous *Miscellaneous) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = miscellaneous.Name
	case "MiscellaneousField":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range miscellaneous.MiscellaneousField {
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

func (miscellaneousfield *MiscellaneousField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = miscellaneousfield.Name
	case "NameAttr":
		res.valueString = miscellaneousfield.NameAttr
	case "Value":
		res.valueString = miscellaneousfield.Value
	}
	return
}

func (namedisplay *NameDisplay) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = namedisplay.Name
	case "PrintObject":
		enum := namedisplay.PrintObject
		res.valueString = enum.ToCodeString()
	case "DisplayText":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range namedisplay.DisplayText {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AccidentalText":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range namedisplay.AccidentalText {
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

func (notesize *NoteSize) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notesize.Name
	case "Type":
		enum := notesize.Type
		res.valueString = enum.ToCodeString()
	case "Value":
		res.valueString = notesize.Value
	}
	return
}

func (otherappearance *OtherAppearance) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = otherappearance.Name
	case "Type":
		res.valueString = otherappearance.Type
	case "Value":
		res.valueString = otherappearance.Value
	}
	return
}

func (otherplay *OtherPlay) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = otherplay.Name
	case "Type":
		res.valueString = otherplay.Type
	case "Value":
		res.valueString = otherplay.Value
	}
	return
}

func (othertext *OtherText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = othertext.Name
	case "Smufl":
		res.valueString = othertext.Smufl
	case "Value":
		res.valueString = othertext.Value
	}
	return
}

func (pagelayout *PageLayout) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pagelayout.Name
	case "PageHeight":
		res.valueString = pagelayout.PageHeight
	case "PageWidth":
		res.valueString = pagelayout.PageWidth
	case "PageMargins":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range pagelayout.PageMargins {
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

func (pagemargins *PageMargins) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pagemargins.Name
	case "Type":
		enum := pagemargins.Type
		res.valueString = enum.ToCodeString()
	case "LeftMargin":
		res.valueString = pagemargins.LeftMargin
	case "RightMargin":
		res.valueString = pagemargins.RightMargin
	case "TopMargin":
		res.valueString = pagemargins.TopMargin
	case "BottomMargin":
		res.valueString = pagemargins.BottomMargin
	}
	return
}

func (play *Play) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = play.Name
	case "Id":
		res.valueString = play.Id
	case "Ipa":
		res.valueString = play.Ipa
	case "Mute":
		enum := play.Mute
		res.valueString = enum.ToCodeString()
	case "SemiPitched":
		enum := play.SemiPitched
		res.valueString = enum.ToCodeString()
	case "OtherPlay":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range play.OtherPlay {
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

func (scaling *Scaling) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = scaling.Name
	case "Millimeters":
		res.valueString = scaling.Millimeters
	case "Tenths":
		res.valueString = scaling.Tenths
	}
	return
}

func (segno *Segno) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = segno.Name
	case "Smufl":
		res.valueString = segno.Smufl
	case "DefaultX":
		res.valueString = segno.DefaultX
	case "DefaultY":
		res.valueString = segno.DefaultY
	case "RelativeX":
		res.valueString = segno.RelativeX
	case "RelativeY":
		res.valueString = segno.RelativeY
	case "FontFamily":
		res.valueString = segno.FontFamily
	case "FontStyle":
		enum := segno.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = segno.FontSize
	case "FontWeight":
		enum := segno.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = segno.Color
	case "Halign":
		enum := segno.Halign
		res.valueString = enum.ToCodeString()
	case "Valign":
		enum := segno.Valign
		res.valueString = enum.ToCodeString()
	case "Id":
		res.valueString = segno.Id
	}
	return
}

func (stafflayout *StaffLayout) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stafflayout.Name
	case "Number":
		res.valueString = fmt.Sprintf("%d", stafflayout.Number)
		res.valueInt = stafflayout.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "StaffDistance":
		res.valueString = stafflayout.StaffDistance
	}
	return
}

func (stringtype *StringType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stringtype.Name
	case "DefaultX":
		res.valueString = stringtype.DefaultX
	case "DefaultY":
		res.valueString = stringtype.DefaultY
	case "RelativeX":
		res.valueString = stringtype.RelativeX
	case "RelativeY":
		res.valueString = stringtype.RelativeY
	case "FontFamily":
		res.valueString = stringtype.FontFamily
	case "FontStyle":
		enum := stringtype.FontStyle
		res.valueString = enum.ToCodeString()
	case "FontSize":
		res.valueString = stringtype.FontSize
	case "FontWeight":
		enum := stringtype.FontWeight
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = stringtype.Color
	case "Placement":
		enum := stringtype.Placement
		res.valueString = enum.ToCodeString()
	case "Value":
		res.valueString = fmt.Sprintf("%d", stringtype.Value)
		res.valueInt = stringtype.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (supports *Supports) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = supports.Name
	case "Type":
		enum := supports.Type
		res.valueString = enum.ToCodeString()
	case "Element":
		res.valueString = supports.Element
	case "Attribute":
		res.valueString = supports.Attribute
	case "Value":
		res.valueString = supports.Value
	}
	return
}

func (systemdividers *SystemDividers) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = systemdividers.Name
	case "LeftDivider":
		res.GongFieldValueType = GongFieldValueTypePointer
		if systemdividers.LeftDivider != nil {
			res.valueString = systemdividers.LeftDivider.Name
			res.ids = systemdividers.LeftDivider.GongGetUUID(stage)
		}
	case "RightDivider":
		res.GongFieldValueType = GongFieldValueTypePointer
		if systemdividers.RightDivider != nil {
			res.valueString = systemdividers.RightDivider.Name
			res.ids = systemdividers.RightDivider.GongGetUUID(stage)
		}
	}
	return
}

func (systemlayout *SystemLayout) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = systemlayout.Name
	case "SystemMargins":
		res.GongFieldValueType = GongFieldValueTypePointer
		if systemlayout.SystemMargins != nil {
			res.valueString = systemlayout.SystemMargins.Name
			res.ids = systemlayout.SystemMargins.GongGetUUID(stage)
		}
	case "SystemDistance":
		res.valueString = systemlayout.SystemDistance
	case "TopSystemDistance":
		res.valueString = systemlayout.TopSystemDistance
	case "SystemDividers":
		res.GongFieldValueType = GongFieldValueTypePointer
		if systemlayout.SystemDividers != nil {
			res.valueString = systemlayout.SystemDividers.Name
			res.ids = systemlayout.SystemDividers.GongGetUUID(stage)
		}
	}
	return
}

func (systemmargins *SystemMargins) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = systemmargins.Name
	case "LeftMargin":
		res.valueString = systemmargins.LeftMargin
	case "RightMargin":
		res.valueString = systemmargins.RightMargin
	}
	return
}

func (typedtext *TypedText) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = typedtext.Name
	case "Type":
		res.valueString = typedtext.Type
	case "Value":
		res.valueString = typedtext.Value
	}
	return
}

func (wavyline *WavyLine) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = wavyline.Name
	case "Type":
		enum := wavyline.Type
		res.valueString = enum.ToCodeString()
	case "Number":
		res.valueString = fmt.Sprintf("%d", wavyline.Number)
		res.valueInt = wavyline.Number
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Smufl":
		res.valueString = wavyline.Smufl
	case "DefaultX":
		res.valueString = wavyline.DefaultX
	case "DefaultY":
		res.valueString = wavyline.DefaultY
	case "RelativeX":
		res.valueString = wavyline.RelativeX
	case "RelativeY":
		res.valueString = wavyline.RelativeY
	case "Placement":
		enum := wavyline.Placement
		res.valueString = enum.ToCodeString()
	case "Color":
		res.valueString = wavyline.Color
	case "StartNote":
		enum := wavyline.StartNote
		res.valueString = enum.ToCodeString()
	case "TrillStep":
		enum := wavyline.TrillStep
		res.valueString = enum.ToCodeString()
	case "TwoNoteTurn":
		enum := wavyline.TwoNoteTurn
		res.valueString = enum.ToCodeString()
	case "Accelerate":
		enum := wavyline.Accelerate
		res.valueString = enum.ToCodeString()
	case "Beats":
		res.valueString = wavyline.Beats
	case "SecondBeat":
		res.valueString = wavyline.SecondBeat
	case "LastBeat":
		res.valueString = wavyline.LastBeat
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
func (accidentaltext *AccidentalText) GongGetGongstructName() string {
	return "AccidentalText"
}

func (appearance *Appearance) GongGetGongstructName() string {
	return "Appearance"
}

func (bookmark *Bookmark) GongGetGongstructName() string {
	return "Bookmark"
}

func (coda *Coda) GongGetGongstructName() string {
	return "Coda"
}

func (distance *Distance) GongGetGongstructName() string {
	return "Distance"
}

func (dynamics *Dynamics) GongGetGongstructName() string {
	return "Dynamics"
}

func (empty *Empty) GongGetGongstructName() string {
	return "Empty"
}

func (emptyplacement *EmptyPlacement) GongGetGongstructName() string {
	return "EmptyPlacement"
}

func (emptyplacementsmufl *EmptyPlacementSmufl) GongGetGongstructName() string {
	return "EmptyPlacementSmufl"
}

func (emptyprintobjectstylealign *EmptyPrintObjectStyleAlign) GongGetGongstructName() string {
	return "EmptyPrintObjectStyleAlign"
}

func (emptyprintstyle *EmptyPrintStyle) GongGetGongstructName() string {
	return "EmptyPrintStyle"
}

func (emptyprintstylealign *EmptyPrintStyleAlign) GongGetGongstructName() string {
	return "EmptyPrintStyleAlign"
}

func (emptyprintstylealignid *EmptyPrintStyleAlignId) GongGetGongstructName() string {
	return "EmptyPrintStyleAlignId"
}

func (emptytrillsound *EmptyTrillSound) GongGetGongstructName() string {
	return "EmptyTrillSound"
}

func (encoding *Encoding) GongGetGongstructName() string {
	return "Encoding"
}

func (fermata *Fermata) GongGetGongstructName() string {
	return "Fermata"
}

func (fingering *Fingering) GongGetGongstructName() string {
	return "Fingering"
}

func (formattedsymbol *FormattedSymbol) GongGetGongstructName() string {
	return "FormattedSymbol"
}

func (formattedsymbolid *FormattedSymbolId) GongGetGongstructName() string {
	return "FormattedSymbolId"
}

func (formattedtext *FormattedText) GongGetGongstructName() string {
	return "FormattedText"
}

func (formattedtextid *FormattedTextId) GongGetGongstructName() string {
	return "FormattedTextId"
}

func (fret *Fret) GongGetGongstructName() string {
	return "Fret"
}

func (glyph *Glyph) GongGetGongstructName() string {
	return "Glyph"
}

func (horizontalturn *HorizontalTurn) GongGetGongstructName() string {
	return "HorizontalTurn"
}

func (identification *Identification) GongGetGongstructName() string {
	return "Identification"
}

func (level *Level) GongGetGongstructName() string {
	return "Level"
}

func (linewidth *LineWidth) GongGetGongstructName() string {
	return "LineWidth"
}

func (link *Link) GongGetGongstructName() string {
	return "Link"
}

func (measurelayout *MeasureLayout) GongGetGongstructName() string {
	return "MeasureLayout"
}

func (mididevice *MidiDevice) GongGetGongstructName() string {
	return "MidiDevice"
}

func (midiinstrument *MidiInstrument) GongGetGongstructName() string {
	return "MidiInstrument"
}

func (miscellaneous *Miscellaneous) GongGetGongstructName() string {
	return "Miscellaneous"
}

func (miscellaneousfield *MiscellaneousField) GongGetGongstructName() string {
	return "MiscellaneousField"
}

func (namedisplay *NameDisplay) GongGetGongstructName() string {
	return "NameDisplay"
}

func (notesize *NoteSize) GongGetGongstructName() string {
	return "NoteSize"
}

func (otherappearance *OtherAppearance) GongGetGongstructName() string {
	return "OtherAppearance"
}

func (otherplay *OtherPlay) GongGetGongstructName() string {
	return "OtherPlay"
}

func (othertext *OtherText) GongGetGongstructName() string {
	return "OtherText"
}

func (pagelayout *PageLayout) GongGetGongstructName() string {
	return "PageLayout"
}

func (pagemargins *PageMargins) GongGetGongstructName() string {
	return "PageMargins"
}

func (play *Play) GongGetGongstructName() string {
	return "Play"
}

func (scaling *Scaling) GongGetGongstructName() string {
	return "Scaling"
}

func (segno *Segno) GongGetGongstructName() string {
	return "Segno"
}

func (stafflayout *StaffLayout) GongGetGongstructName() string {
	return "StaffLayout"
}

func (stringtype *StringType) GongGetGongstructName() string {
	return "StringType"
}

func (supports *Supports) GongGetGongstructName() string {
	return "Supports"
}

func (systemdividers *SystemDividers) GongGetGongstructName() string {
	return "SystemDividers"
}

func (systemlayout *SystemLayout) GongGetGongstructName() string {
	return "SystemLayout"
}

func (systemmargins *SystemMargins) GongGetGongstructName() string {
	return "SystemMargins"
}

func (typedtext *TypedText) GongGetGongstructName() string {
	return "TypedText"
}

func (wavyline *WavyLine) GongGetGongstructName() string {
	return "WavyLine"
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
	__gong__rebuildMapString(stage.AccidentalTexts, &stage.AccidentalTexts_mapString)

	__gong__rebuildMapString(stage.Appearances, &stage.Appearances_mapString)

	__gong__rebuildMapString(stage.Bookmarks, &stage.Bookmarks_mapString)

	__gong__rebuildMapString(stage.Codas, &stage.Codas_mapString)

	__gong__rebuildMapString(stage.Distances, &stage.Distances_mapString)

	__gong__rebuildMapString(stage.Dynamicss, &stage.Dynamicss_mapString)

	__gong__rebuildMapString(stage.Emptys, &stage.Emptys_mapString)

	__gong__rebuildMapString(stage.EmptyPlacements, &stage.EmptyPlacements_mapString)

	__gong__rebuildMapString(stage.EmptyPlacementSmufls, &stage.EmptyPlacementSmufls_mapString)

	__gong__rebuildMapString(stage.EmptyPrintObjectStyleAligns, &stage.EmptyPrintObjectStyleAligns_mapString)

	__gong__rebuildMapString(stage.EmptyPrintStyles, &stage.EmptyPrintStyles_mapString)

	__gong__rebuildMapString(stage.EmptyPrintStyleAligns, &stage.EmptyPrintStyleAligns_mapString)

	__gong__rebuildMapString(stage.EmptyPrintStyleAlignIds, &stage.EmptyPrintStyleAlignIds_mapString)

	__gong__rebuildMapString(stage.EmptyTrillSounds, &stage.EmptyTrillSounds_mapString)

	__gong__rebuildMapString(stage.Encodings, &stage.Encodings_mapString)

	__gong__rebuildMapString(stage.Fermatas, &stage.Fermatas_mapString)

	__gong__rebuildMapString(stage.Fingerings, &stage.Fingerings_mapString)

	__gong__rebuildMapString(stage.FormattedSymbols, &stage.FormattedSymbols_mapString)

	__gong__rebuildMapString(stage.FormattedSymbolIds, &stage.FormattedSymbolIds_mapString)

	__gong__rebuildMapString(stage.FormattedTexts, &stage.FormattedTexts_mapString)

	__gong__rebuildMapString(stage.FormattedTextIds, &stage.FormattedTextIds_mapString)

	__gong__rebuildMapString(stage.Frets, &stage.Frets_mapString)

	__gong__rebuildMapString(stage.Glyphs, &stage.Glyphs_mapString)

	__gong__rebuildMapString(stage.HorizontalTurns, &stage.HorizontalTurns_mapString)

	__gong__rebuildMapString(stage.Identifications, &stage.Identifications_mapString)

	__gong__rebuildMapString(stage.Levels, &stage.Levels_mapString)

	__gong__rebuildMapString(stage.LineWidths, &stage.LineWidths_mapString)

	__gong__rebuildMapString(stage.Links, &stage.Links_mapString)

	__gong__rebuildMapString(stage.MeasureLayouts, &stage.MeasureLayouts_mapString)

	__gong__rebuildMapString(stage.MidiDevices, &stage.MidiDevices_mapString)

	__gong__rebuildMapString(stage.MidiInstruments, &stage.MidiInstruments_mapString)

	__gong__rebuildMapString(stage.Miscellaneouss, &stage.Miscellaneouss_mapString)

	__gong__rebuildMapString(stage.MiscellaneousFields, &stage.MiscellaneousFields_mapString)

	__gong__rebuildMapString(stage.NameDisplays, &stage.NameDisplays_mapString)

	__gong__rebuildMapString(stage.NoteSizes, &stage.NoteSizes_mapString)

	__gong__rebuildMapString(stage.OtherAppearances, &stage.OtherAppearances_mapString)

	__gong__rebuildMapString(stage.OtherPlays, &stage.OtherPlays_mapString)

	__gong__rebuildMapString(stage.OtherTexts, &stage.OtherTexts_mapString)

	__gong__rebuildMapString(stage.PageLayouts, &stage.PageLayouts_mapString)

	__gong__rebuildMapString(stage.PageMarginss, &stage.PageMarginss_mapString)

	__gong__rebuildMapString(stage.Plays, &stage.Plays_mapString)

	__gong__rebuildMapString(stage.Scalings, &stage.Scalings_mapString)

	__gong__rebuildMapString(stage.Segnos, &stage.Segnos_mapString)

	__gong__rebuildMapString(stage.StaffLayouts, &stage.StaffLayouts_mapString)

	__gong__rebuildMapString(stage.StringTypes, &stage.StringTypes_mapString)

	__gong__rebuildMapString(stage.Supportss, &stage.Supportss_mapString)

	__gong__rebuildMapString(stage.SystemDividerss, &stage.SystemDividerss_mapString)

	__gong__rebuildMapString(stage.SystemLayouts, &stage.SystemLayouts_mapString)

	__gong__rebuildMapString(stage.SystemMarginss, &stage.SystemMarginss_mapString)

	__gong__rebuildMapString(stage.TypedTexts, &stage.TypedTexts_mapString)

	__gong__rebuildMapString(stage.WavyLines, &stage.WavyLines_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
