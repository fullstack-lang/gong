// generated code by xsd2gong - do not edit
package direction

import (
	"encoding/xml"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
)

var _ xml.Attr

// EnumBeaterValue: The beater-value type represents pictograms for beaters, mallets, and
// sticks that do not have different materials represented in the pictogram. The finger
// and hammer values are in addition to Stone's list.
type EnumBeaterValue string

const (
	EnumBeaterValueBow EnumBeaterValue = "bow"
	EnumBeaterValueChimeHammer EnumBeaterValue = "chime hammer"
	EnumBeaterValueCoin EnumBeaterValue = "coin"
	EnumBeaterValueDrumStick EnumBeaterValue = "drum stick"
	EnumBeaterValueFinger EnumBeaterValue = "finger"
	EnumBeaterValueFingernail EnumBeaterValue = "fingernail"
	EnumBeaterValueFist EnumBeaterValue = "fist"
	EnumBeaterValueGuiroScraper EnumBeaterValue = "guiro scraper"
	EnumBeaterValueHammer EnumBeaterValue = "hammer"
	EnumBeaterValueHand EnumBeaterValue = "hand"
	EnumBeaterValueJazzStick EnumBeaterValue = "jazz stick"
	EnumBeaterValueKnittingNeedle EnumBeaterValue = "knitting needle"
	EnumBeaterValueMetalHammer EnumBeaterValue = "metal hammer"
	EnumBeaterValueSlideBrushOnGong EnumBeaterValue = "slide brush on gong"
	EnumBeaterValueSnareStick EnumBeaterValue = "snare stick"
	EnumBeaterValueSpoonMallet EnumBeaterValue = "spoon mallet"
	EnumBeaterValueSuperball EnumBeaterValue = "superball"
	EnumBeaterValueTriangleBeater EnumBeaterValue = "triangle beater"
	EnumBeaterValueTriangleBeaterPlain EnumBeaterValue = "triangle beater plain"
	EnumBeaterValueWireBrush EnumBeaterValue = "wire brush"
)

// EnumDegreeSymbolValue: The degree-symbol-value type indicates which symbol should be used in
// specifying a degree.
type EnumDegreeSymbolValue string

const (
	EnumDegreeSymbolValueMajor EnumDegreeSymbolValue = "major"
	EnumDegreeSymbolValueMinor EnumDegreeSymbolValue = "minor"
	EnumDegreeSymbolValueAugmented EnumDegreeSymbolValue = "augmented"
	EnumDegreeSymbolValueDiminished EnumDegreeSymbolValue = "diminished"
	EnumDegreeSymbolValueHalfDiminished EnumDegreeSymbolValue = "half-diminished"
)

// EnumDegreeTypeValue: The degree-type-value type indicates whether the current degree
// element is an addition, alteration, or subtraction to the kind of the current chord
// in the harmony element.
type EnumDegreeTypeValue string

const (
	EnumDegreeTypeValueAdd EnumDegreeTypeValue = "add"
	EnumDegreeTypeValueAlter EnumDegreeTypeValue = "alter"
	EnumDegreeTypeValueSubtract EnumDegreeTypeValue = "subtract"
)

// EnumEffectValue: The effect-value type represents pictograms for sound effect
// percussion instruments. The cannon, lotus flute, and megaphone values are in
// addition to Stone's list.
type EnumEffectValue string

const (
	EnumEffectValueAnvil EnumEffectValue = "anvil"
	EnumEffectValueAutoHorn EnumEffectValue = "auto horn"
	EnumEffectValueBirdWhistle EnumEffectValue = "bird whistle"
	EnumEffectValueCannon EnumEffectValue = "cannon"
	EnumEffectValueDuckCall EnumEffectValue = "duck call"
	EnumEffectValueGunShot EnumEffectValue = "gun shot"
	EnumEffectValueKlaxonHorn EnumEffectValue = "klaxon horn"
	EnumEffectValueLionsRoar EnumEffectValue = "lions roar"
	EnumEffectValueLotusFlute EnumEffectValue = "lotus flute"
	EnumEffectValueMegaphone EnumEffectValue = "megaphone"
	EnumEffectValuePoliceWhistle EnumEffectValue = "police whistle"
	EnumEffectValueSiren EnumEffectValue = "siren"
	EnumEffectValueSlideWhistle EnumEffectValue = "slide whistle"
	EnumEffectValueThunderSheet EnumEffectValue = "thunder sheet"
	EnumEffectValueWindMachine EnumEffectValue = "wind machine"
	EnumEffectValueWindWhistle EnumEffectValue = "wind whistle"
)

// EnumGlassValue: The glass-value type represents pictograms for glass percussion
// instruments.
type EnumGlassValue string

const (
	EnumGlassValueGlassHarmonica EnumGlassValue = "glass harmonica"
	EnumGlassValueGlassHarp EnumGlassValue = "glass harp"
	EnumGlassValueWindChimes EnumGlassValue = "wind chimes"
)

// EnumHarmonyArrangement: The harmony-arrangement type indicates how stacked chords and bass
// notes are displayed within a harmony element. The vertical value specifies that the
// second element appears below the first. The horizontal value specifies that the
// second element appears to the right of the first. The diagonal value specifies that
// the second element appears both below and to the right of the first.
type EnumHarmonyArrangement string

const (
	EnumHarmonyArrangementVertical EnumHarmonyArrangement = "vertical"
	EnumHarmonyArrangementHorizontal EnumHarmonyArrangement = "horizontal"
	EnumHarmonyArrangementDiagonal EnumHarmonyArrangement = "diagonal"
)

// EnumHarmonyType: The harmony-type type differentiates different types of harmonies when
// alternate harmonies are possible. Explicit harmonies have all note present in the
// music; implied have some notes missing but implied; alternate represents alternate
// analyses.
type EnumHarmonyType string

const (
	EnumHarmonyTypeExplicit EnumHarmonyType = "explicit"
	EnumHarmonyTypeImplied EnumHarmonyType = "implied"
	EnumHarmonyTypeAlternate EnumHarmonyType = "alternate"
)

// EnumKindValue: A kind-value indicates the type of chord. Degree elements can then
// add, subtract, or alter from these starting points. Values include: Triads: major
// (major third, perfect fifth) minor (minor third, perfect fifth) augmented (major
// third, augmented fifth) diminished (minor third, diminished fifth) Sevenths:
// dominant (major triad, minor seventh) major-seventh (major triad, major seventh)
// minor-seventh (minor triad, minor seventh) diminished-seventh (diminished triad,
// diminished seventh) augmented-seventh (augmented triad, minor seventh)
// half-diminished (diminished triad, minor seventh) major-minor (minor triad, major
// seventh) Sixths: major-sixth (major triad, added sixth) minor-sixth (minor triad,
// added sixth) Ninths: dominant-ninth (dominant-seventh, major ninth) major-ninth
// (major-seventh, major ninth) minor-ninth (minor-seventh, major ninth) 11ths (usually
// as the basis for alteration): dominant-11th (dominant-ninth, perfect 11th)
// major-11th (major-ninth, perfect 11th) minor-11th (minor-ninth, perfect 11th) 13ths
// (usually as the basis for alteration): dominant-13th (dominant-11th, major 13th)
// major-13th (major-11th, major 13th) minor-13th (minor-11th, major 13th) Suspended:
// suspended-second (major second, perfect fifth) suspended-fourth (perfect fourth,
// perfect fifth) Functional sixths: Neapolitan Italian French German Other: pedal
// (pedal-point bass) power (perfect fifth) Tristan The "other" kind is used when the
// harmony is entirely composed of add elements. The "none" kind is used to explicitly
// encode absence of chords or functional harmony. In this case, the root, numeral, or
// function element has no meaning. When using the root or numeral element, the
// root-step or numeral-step text attribute should be set to the empty string to keep
// the root or numeral from being displayed.
type EnumKindValue string

const (
	EnumKindValueMajor EnumKindValue = "major"
	EnumKindValueMinor EnumKindValue = "minor"
	EnumKindValueAugmented EnumKindValue = "augmented"
	EnumKindValueDiminished EnumKindValue = "diminished"
	EnumKindValueDominant EnumKindValue = "dominant"
	EnumKindValueMajorSeventh EnumKindValue = "major-seventh"
	EnumKindValueMinorSeventh EnumKindValue = "minor-seventh"
	EnumKindValueDiminishedSeventh EnumKindValue = "diminished-seventh"
	EnumKindValueAugmentedSeventh EnumKindValue = "augmented-seventh"
	EnumKindValueHalfDiminished EnumKindValue = "half-diminished"
	EnumKindValueMajorMinor EnumKindValue = "major-minor"
	EnumKindValueMajorSixth EnumKindValue = "major-sixth"
	EnumKindValueMinorSixth EnumKindValue = "minor-sixth"
	EnumKindValueDominantNinth EnumKindValue = "dominant-ninth"
	EnumKindValueMajorNinth EnumKindValue = "major-ninth"
	EnumKindValueMinorNinth EnumKindValue = "minor-ninth"
	EnumKindValueDominant11th EnumKindValue = "dominant-11th"
	EnumKindValueMajor11th EnumKindValue = "major-11th"
	EnumKindValueMinor11th EnumKindValue = "minor-11th"
	EnumKindValueDominant13th EnumKindValue = "dominant-13th"
	EnumKindValueMajor13th EnumKindValue = "major-13th"
	EnumKindValueMinor13th EnumKindValue = "minor-13th"
	EnumKindValueSuspendedSecond EnumKindValue = "suspended-second"
	EnumKindValueSuspendedFourth EnumKindValue = "suspended-fourth"
	EnumKindValueNeapolitan EnumKindValue = "Neapolitan"
	EnumKindValueItalian EnumKindValue = "Italian"
	EnumKindValueFrench EnumKindValue = "French"
	EnumKindValueGerman EnumKindValue = "German"
	EnumKindValuePedal EnumKindValue = "pedal"
	EnumKindValuePower EnumKindValue = "power"
	EnumKindValueTristan EnumKindValue = "Tristan"
	EnumKindValueOther EnumKindValue = "other"
	EnumKindValueNone EnumKindValue = "none"
)

// EnumLineEnd: The line-end type specifies if there is a jog up or down (or both), an
// arrow, or nothing at the start or end of a bracket.
type EnumLineEnd string

const (
	EnumLineEndUp EnumLineEnd = "up"
	EnumLineEndDown EnumLineEnd = "down"
	EnumLineEndBoth EnumLineEnd = "both"
	EnumLineEndArrow EnumLineEnd = "arrow"
	EnumLineEndNone EnumLineEnd = "none"
)

// EnumMeasureNumberingValue: The measure-numbering-value type describes how measure numbers are
// displayed on this part: no numbers, numbers every measure, or numbers every system.
type EnumMeasureNumberingValue string

const (
	EnumMeasureNumberingValueNone EnumMeasureNumberingValue = "none"
	EnumMeasureNumberingValueMeasure EnumMeasureNumberingValue = "measure"
	EnumMeasureNumberingValueSystem EnumMeasureNumberingValue = "system"
)

// EnumMembraneValue: The membrane-value type represents pictograms for membrane percussion
// instruments.
type EnumMembraneValue string

const (
	EnumMembraneValueBassDrum EnumMembraneValue = "bass drum"
	EnumMembraneValueBassDrumOnSide EnumMembraneValue = "bass drum on side"
	EnumMembraneValueBongos EnumMembraneValue = "bongos"
	EnumMembraneValueChineseTomtom EnumMembraneValue = "Chinese tomtom"
	EnumMembraneValueCongaDrum EnumMembraneValue = "conga drum"
	EnumMembraneValueCuica EnumMembraneValue = "cuica"
	EnumMembraneValueGobletDrum EnumMembraneValue = "goblet drum"
	EnumMembraneValueIndoAmericanTomtom EnumMembraneValue = "Indo-American tomtom"
	EnumMembraneValueJapaneseTomtom EnumMembraneValue = "Japanese tomtom"
	EnumMembraneValueMilitaryDrum EnumMembraneValue = "military drum"
	EnumMembraneValueSnareDrum EnumMembraneValue = "snare drum"
	EnumMembraneValueSnareDrumSnaresOff EnumMembraneValue = "snare drum snares off"
	EnumMembraneValueTabla EnumMembraneValue = "tabla"
	EnumMembraneValueTambourine EnumMembraneValue = "tambourine"
	EnumMembraneValueTenorDrum EnumMembraneValue = "tenor drum"
	EnumMembraneValueTimbales EnumMembraneValue = "timbales"
	EnumMembraneValueTomtom EnumMembraneValue = "tomtom"
)

// EnumMetalValue: The metal-value type represents pictograms for metal percussion
// instruments. The hi-hat value refers to a pictogram like Stone's high-hat cymbals
// but without the long vertical line at the bottom.
type EnumMetalValue string

const (
	EnumMetalValueAgogo EnumMetalValue = "agogo"
	EnumMetalValueAlmglocken EnumMetalValue = "almglocken"
	EnumMetalValueBell EnumMetalValue = "bell"
	EnumMetalValueBellPlate EnumMetalValue = "bell plate"
	EnumMetalValueBellTree EnumMetalValue = "bell tree"
	EnumMetalValueBrakeDrum EnumMetalValue = "brake drum"
	EnumMetalValueCencerro EnumMetalValue = "cencerro"
	EnumMetalValueChainRattle EnumMetalValue = "chain rattle"
	EnumMetalValueChineseCymbal EnumMetalValue = "Chinese cymbal"
	EnumMetalValueCowbell EnumMetalValue = "cowbell"
	EnumMetalValueCrashCymbals EnumMetalValue = "crash cymbals"
	EnumMetalValueCrotale EnumMetalValue = "crotale"
	EnumMetalValueCymbalTongs EnumMetalValue = "cymbal tongs"
	EnumMetalValueDomedGong EnumMetalValue = "domed gong"
	EnumMetalValueFingerCymbals EnumMetalValue = "finger cymbals"
	EnumMetalValueFlexatone EnumMetalValue = "flexatone"
	EnumMetalValueGong EnumMetalValue = "gong"
	EnumMetalValueHiHat EnumMetalValue = "hi-hat"
	EnumMetalValueHighHatCymbals EnumMetalValue = "high-hat cymbals"
	EnumMetalValueHandbell EnumMetalValue = "handbell"
	EnumMetalValueJawHarp EnumMetalValue = "jaw harp"
	EnumMetalValueJingleBells EnumMetalValue = "jingle bells"
	EnumMetalValueMusicalSaw EnumMetalValue = "musical saw"
	EnumMetalValueShellBells EnumMetalValue = "shell bells"
	EnumMetalValueSistrum EnumMetalValue = "sistrum"
	EnumMetalValueSizzleCymbal EnumMetalValue = "sizzle cymbal"
	EnumMetalValueSleighBells EnumMetalValue = "sleigh bells"
	EnumMetalValueSuspendedCymbal EnumMetalValue = "suspended cymbal"
	EnumMetalValueTamTam EnumMetalValue = "tam tam"
	EnumMetalValueTamTamWithBeater EnumMetalValue = "tam tam with beater"
	EnumMetalValueTriangle EnumMetalValue = "triangle"
	EnumMetalValueVietnameseHat EnumMetalValue = "Vietnamese hat"
)

// EnumNumeralMode: The numeral-mode type specifies the mode similar to the mode type, but
// with a restricted set of values. The different minor values are used to interpret
// numeral-root values of 6 and 7 when present in a minor key. The harmonic minor value
// sharpens the 7 and the melodic minor value sharpens both 6 and 7. If a minor mode is
// used without qualification, either in the mode or numeral-mode elements, natural
// minor is used.
type EnumNumeralMode string

const (
	EnumNumeralModeMajor EnumNumeralMode = "major"
	EnumNumeralModeMinor EnumNumeralMode = "minor"
	EnumNumeralModeNaturalMinor EnumNumeralMode = "natural minor"
	EnumNumeralModeMelodicMinor EnumNumeralMode = "melodic minor"
	EnumNumeralModeHarmonicMinor EnumNumeralMode = "harmonic minor"
)

// EnumOnOff: The on-off type is used for notation elements such as string mutes.
type EnumOnOff string

const (
	EnumOnOffOn EnumOnOff = "on"
	EnumOnOffOff EnumOnOff = "off"
)

// EnumPedalType: The pedal-type simple type is used to distinguish types of pedal
// directions. The start value indicates the start of a damper pedal, while the
// sostenuto value indicates the start of a sostenuto pedal. The other values can be
// used with either the damper or sostenuto pedal. The soft pedal is not included here
// because there is no special symbol or graphic used for it beyond what can be
// specified with words and bracket elements. The change, continue, discontinue, and
// resume types are used when the line attribute is yes. The change type indicates a
// pedal lift and retake indicated with an inverted V marking. The continue type allows
// more precise formatting across system breaks and for more complex pedaling lines.
// The discontinue type indicates the end of a pedal line that does not include the
// explicit lift represented by the stop type. The resume type indicates the start of a
// pedal line that does not include the downstroke represented by the start type. It
// can be used when a line resumes after being discontinued, or to start a pedal line
// that is preceded by a text or symbol representation of the pedal.
type EnumPedalType string

const (
	EnumPedalTypeStart EnumPedalType = "start"
	EnumPedalTypeStop EnumPedalType = "stop"
	EnumPedalTypeSostenuto EnumPedalType = "sostenuto"
	EnumPedalTypeChange EnumPedalType = "change"
	EnumPedalTypeContinue EnumPedalType = "continue"
	EnumPedalTypeDiscontinue EnumPedalType = "discontinue"
	EnumPedalTypeResume EnumPedalType = "resume"
)

// EnumPitchedValue: The pitched-value type represents pictograms for pitched percussion
// instruments. The chimes and tubular chimes values distinguish the single-line and
// double-line versions of the pictogram.
type EnumPitchedValue string

const (
	EnumPitchedValueCelesta EnumPitchedValue = "celesta"
	EnumPitchedValueChimes EnumPitchedValue = "chimes"
	EnumPitchedValueGlockenspiel EnumPitchedValue = "glockenspiel"
	EnumPitchedValueLithophone EnumPitchedValue = "lithophone"
	EnumPitchedValueMallet EnumPitchedValue = "mallet"
	EnumPitchedValueMarimba EnumPitchedValue = "marimba"
	EnumPitchedValueSteelDrums EnumPitchedValue = "steel drums"
	EnumPitchedValueTubaphone EnumPitchedValue = "tubaphone"
	EnumPitchedValueTubularChimes EnumPitchedValue = "tubular chimes"
	EnumPitchedValueVibraphone EnumPitchedValue = "vibraphone"
	EnumPitchedValueXylophone EnumPitchedValue = "xylophone"
)

// EnumPrincipalVoiceSymbol: The principal-voice-symbol type represents the type of symbol used to
// indicate a principal or secondary voice. The "plain" value represents a plain square
// bracket. The value of "none" is used for analysis markup when the principal-voice
// element does not have a corresponding appearance in the score.
type EnumPrincipalVoiceSymbol string

const (
	EnumPrincipalVoiceSymbolHauptstimme EnumPrincipalVoiceSymbol = "Hauptstimme"
	EnumPrincipalVoiceSymbolNebenstimme EnumPrincipalVoiceSymbol = "Nebenstimme"
	EnumPrincipalVoiceSymbolPlain EnumPrincipalVoiceSymbol = "plain"
	EnumPrincipalVoiceSymbolNone EnumPrincipalVoiceSymbol = "none"
)

// EnumStaffDivideSymbol: The staff-divide-symbol type is used for staff division symbols. The
// down, up, and up-down values correspond to SMuFL code points U+E00B, U+E00C, and
// U+E00D respectively.
type EnumStaffDivideSymbol string

const (
	EnumStaffDivideSymbolDown EnumStaffDivideSymbol = "down"
	EnumStaffDivideSymbolUp EnumStaffDivideSymbol = "up"
	EnumStaffDivideSymbolUpDown EnumStaffDivideSymbol = "up-down"
)

// EnumStartStopChangeContinue: The start-stop-change-continue type is used to distinguish types of
// pedal directions.
type EnumStartStopChangeContinue string

const (
	EnumStartStopChangeContinueStart EnumStartStopChangeContinue = "start"
	EnumStartStopChangeContinueStop EnumStartStopChangeContinue = "stop"
	EnumStartStopChangeContinueChange EnumStartStopChangeContinue = "change"
	EnumStartStopChangeContinueContinue EnumStartStopChangeContinue = "continue"
)

// EnumSyncType: The sync-type type specifies the style that a score following
// application should use to synchronize an accompaniment with a performer. The none
// type indicates no synchronization to the performer. The tempo type indicates
// synchronization based on the performer tempo rather than individual events in the
// score. The event type indicates synchronization by following the performance of
// individual events in the score rather than the performer tempo. The mostly-tempo and
// mostly-event types combine these two approaches, with mostly-tempo giving more
// weight to tempo and mostly-event giving more weight to performed events. The
// always-event type provides the strictest synchronization by not being forgiving of
// missing performed events.
type EnumSyncType string

const (
	EnumSyncTypeNone EnumSyncType = "none"
	EnumSyncTypeTempo EnumSyncType = "tempo"
	EnumSyncTypeMostlyTempo EnumSyncType = "mostly-tempo"
	EnumSyncTypeMostlyEvent EnumSyncType = "mostly-event"
	EnumSyncTypeEvent EnumSyncType = "event"
	EnumSyncTypeAlwaysEvent EnumSyncType = "always-event"
)

// EnumTipDirection: The tip-direction type represents the direction in which the tip of a
// stick or beater points, using Unicode arrow terminology.
type EnumTipDirection string

const (
	EnumTipDirectionUp EnumTipDirection = "up"
	EnumTipDirectionDown EnumTipDirection = "down"
	EnumTipDirectionLeft EnumTipDirection = "left"
	EnumTipDirectionRight EnumTipDirection = "right"
	EnumTipDirectionNorthwest EnumTipDirection = "northwest"
	EnumTipDirectionNortheast EnumTipDirection = "northeast"
	EnumTipDirectionSoutheast EnumTipDirection = "southeast"
	EnumTipDirectionSouthwest EnumTipDirection = "southwest"
)

// EnumStickLocation: The stick-location type represents pictograms for the location of
// sticks, beaters, or mallets on cymbals, gongs, drums, and other instruments.
type EnumStickLocation string

const (
	EnumStickLocationCenter EnumStickLocation = "center"
	EnumStickLocationRim EnumStickLocation = "rim"
	EnumStickLocationCymbalBell EnumStickLocation = "cymbal bell"
	EnumStickLocationCymbalEdge EnumStickLocation = "cymbal edge"
)

// EnumStickMaterial: The stick-material type represents the material being displayed in a
// stick pictogram.
type EnumStickMaterial string

const (
	EnumStickMaterialSoft EnumStickMaterial = "soft"
	EnumStickMaterialMedium EnumStickMaterial = "medium"
	EnumStickMaterialHard EnumStickMaterial = "hard"
	EnumStickMaterialShaded EnumStickMaterial = "shaded"
	EnumStickMaterialX EnumStickMaterial = "x"
)

// EnumStickType: The stick-type type represents the shape of pictograms where the
// material in the stick, mallet, or beater is represented in the pictogram.
type EnumStickType string

const (
	EnumStickTypeBassDrum EnumStickType = "bass drum"
	EnumStickTypeDoubleBassDrum EnumStickType = "double bass drum"
	EnumStickTypeGlockenspiel EnumStickType = "glockenspiel"
	EnumStickTypeGum EnumStickType = "gum"
	EnumStickTypeHammer EnumStickType = "hammer"
	EnumStickTypeSuperball EnumStickType = "superball"
	EnumStickTypeTimpani EnumStickType = "timpani"
	EnumStickTypeWound EnumStickType = "wound"
	EnumStickTypeXylophone EnumStickType = "xylophone"
	EnumStickTypeYarn EnumStickType = "yarn"
)

// EnumUpDownStopContinue: The up-down-stop-continue type is used for octave-shift elements,
// indicating the direction of the shift from their true pitched values because of
// printing difficulty.
type EnumUpDownStopContinue string

const (
	EnumUpDownStopContinueUp EnumUpDownStopContinue = "up"
	EnumUpDownStopContinueDown EnumUpDownStopContinue = "down"
	EnumUpDownStopContinueStop EnumUpDownStopContinue = "stop"
	EnumUpDownStopContinueContinue EnumUpDownStopContinue = "continue"
)

// EnumWedgeType: The wedge type is crescendo for the start of a wedge that is closed at
// the left side, diminuendo for the start of a wedge that is closed on the right side,
// and stop for the end of a wedge. The continue type is used for formatting wedges
// over a system break, or for other situations where a single wedge is divided into
// multiple segments.
type EnumWedgeType string

const (
	EnumWedgeTypeCrescendo EnumWedgeType = "crescendo"
	EnumWedgeTypeDiminuendo EnumWedgeType = "diminuendo"
	EnumWedgeTypeStop EnumWedgeType = "stop"
	EnumWedgeTypeContinue EnumWedgeType = "continue"
)

// EnumWoodValue: The wood-value type represents pictograms for wood percussion
// instruments. The maraca and maracas values distinguish the one- and two-maraca
// versions of the pictogram.
type EnumWoodValue string

const (
	EnumWoodValueBambooScraper EnumWoodValue = "bamboo scraper"
	EnumWoodValueBoardClapper EnumWoodValue = "board clapper"
	EnumWoodValueCabasa EnumWoodValue = "cabasa"
	EnumWoodValueCastanets EnumWoodValue = "castanets"
	EnumWoodValueCastanetsWithHandle EnumWoodValue = "castanets with handle"
	EnumWoodValueClaves EnumWoodValue = "claves"
	EnumWoodValueFootballRattle EnumWoodValue = "football rattle"
	EnumWoodValueGuiro EnumWoodValue = "guiro"
	EnumWoodValueLogDrum EnumWoodValue = "log drum"
	EnumWoodValueMaraca EnumWoodValue = "maraca"
	EnumWoodValueMaracas EnumWoodValue = "maracas"
	EnumWoodValueQuijada EnumWoodValue = "quijada"
	EnumWoodValueRainstick EnumWoodValue = "rainstick"
	EnumWoodValueRatchet EnumWoodValue = "ratchet"
	EnumWoodValueRecoReco EnumWoodValue = "reco-reco"
	EnumWoodValueSandpaperBlocks EnumWoodValue = "sandpaper blocks"
	EnumWoodValueSlitDrum EnumWoodValue = "slit drum"
	EnumWoodValueTempleBlock EnumWoodValue = "temple block"
	EnumWoodValueVibraslap EnumWoodValue = "vibraslap"
	EnumWoodValueWhip EnumWoodValue = "whip"
	EnumWoodValueWoodBlock EnumWoodValue = "wood block"
)

// EnumBeamValue: The beam-value type represents the type of beam associated with each
// of 8 beam levels (up to 1024th notes) available for each note.
type EnumBeamValue string

const (
	EnumBeamValueBegin EnumBeamValue = "begin"
	EnumBeamValueContinue EnumBeamValue = "continue"
	EnumBeamValueEnd EnumBeamValue = "end"
	EnumBeamValueForwardHook EnumBeamValue = "forward hook"
	EnumBeamValueBackwardHook EnumBeamValue = "backward hook"
)

// EnumShowTuplet: The show-tuplet type indicates whether to show a part of a tuplet
// relating to the tuplet-actual element, both the tuplet-actual and tuplet-normal
// elements, or neither.
type EnumShowTuplet string

const (
	EnumShowTupletActual EnumShowTuplet = "actual"
	EnumShowTupletBoth EnumShowTuplet = "both"
	EnumShowTupletNone EnumShowTuplet = "none"
)

// Accord: The accord type represents the tuning of a single string in the
// scordatura element. It uses the same group of elements as the staff-tuning element.
// Strings are numbered from high to low.
type Accord struct {
	Name string `xml:"-"`
	String int `xml:"string,attr,omitempty"`
	TuningStep attributes.EnumStep `xml:"tuning-step,omitempty"`
	TuningAlter string `xml:"tuning-alter,omitempty"`
	TuningOctave int `xml:"tuning-octave,omitempty"`
}

// AccordionRegistration: The accordion-registration type is used for accordion registration
// symbols. These are circular symbols divided horizontally into high, middle, and low
// sections that correspond to 4', 8', and 16' pipes. Each accordion-high,
// accordion-middle, and accordion-low element represents the presence of one or more
// dots in the registration diagram. An accordion-registration element needs to have at
// least one of the child elements present.
type AccordionRegistration struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	AccordionHigh *common.Empty `xml:"accordion-high,omitempty"`
	AccordionMiddle int `xml:"accordion-middle,omitempty"`
	AccordionLow *common.Empty `xml:"accordion-low,omitempty"`
}

// Barre: The barre element indicates placing a finger over multiple strings on
// a single fret. The type is "start" for the lowest pitched string (e.g., the string
// with the highest MusicXML number) and is "stop" for the highest pitched string.
type Barre struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
}

// Bass: The bass type is used to indicate a bass note in popular music chord
// symbols, e.g. G/C. It is generally not used in functional harmony, as inversion is
// generally not used in pop chord symbols. As with root, it is divided into step and
// alter elements, similar to pitches. The arrangement attribute specifies where the
// bass is displayed relative to what precedes it.
type Bass struct {
	Name string `xml:"-"`
	Arrangement EnumHarmonyArrangement `xml:"arrangement,attr,omitempty"`
	BassSeparator *StyleText `xml:"bass-separator,omitempty"`
	BassStep *BassStep `xml:"bass-step,omitempty"`
	BassAlter *HarmonyAlter `xml:"bass-alter,omitempty"`
}

// HarmonyAlter: The harmony-alter type represents the chromatic alteration of the
// root, numeral, or bass of the current harmony-chord group within the harmony
// element. In some chord styles, the text of the preceding element may include
// alteration information. In that case, the print-object attribute of this type can be
// set to no. The location attribute indicates whether the alteration should appear to
// the left or the right of the preceding element. Its default value varies by element.
type HarmonyAlter struct {
	Name string `xml:"-"`
	Location common.EnumLeftRight `xml:"location,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value string `xml:",chardata"`
}

// BassStep: The bass-step type represents the pitch step of the bass of the
// current chord within the harmony element. The text attribute indicates how the bass
// should appear in a score if not using the element contents.
type BassStep struct {
	Name string `xml:"-"`
	Text string `xml:"text,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value attributes.EnumStep `xml:",chardata"`
}

// Beater: The beater type represents pictograms for beaters, mallets, and sticks
// that do not have different materials represented in the pictogram.
type Beater struct {
	Name string `xml:"-"`
	Tip EnumTipDirection `xml:"tip,attr,omitempty"`
	Value EnumBeaterValue `xml:",chardata"`
}

// BeatUnitTied: The beat-unit-tied type indicates a beat-unit within a metronome mark
// that is tied to the preceding beat-unit. This allows two or more tied notes to be
// associated with a per-minute value in a metronome mark, whereas the metronome-tied
// element is restricted to metric relationship marks.
type BeatUnitTied struct {
	Name string `xml:"-"`
	BeatUnit common.EnumNoteTypeValue `xml:"beat-unit,omitempty"`
	BeatUnitDot []*common.Empty `xml:"beat-unit-dot,omitempty"`
}

// Bracket: Brackets are combined with words in a variety of modern directions.
// The line-end attribute specifies if there is a jog up or down (or both), an arrow,
// or nothing at the start or end of the bracket. If the line-end is up or down, the
// length of the jog can be specified using the end-length attribute. The line-type is
// solid if not specified.
type Bracket struct {
	Name string `xml:"-"`
	Type common.EnumStartStopContinue `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	LineEnd EnumLineEnd `xml:"line-end,attr,omitempty"`
	EndLength string `xml:"end-length,attr,omitempty"`
	LineType common.EnumLineType `xml:"line-type,attr,omitempty"`
	DashLength string `xml:"dash-length,attr,omitempty"`
	SpaceLength string `xml:"space-length,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Dashes: The dashes type represents dashes, used for instance with cresc. and
// dim. marks.
type Dashes struct {
	Name string `xml:"-"`
	Type common.EnumStartStopContinue `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	DashLength string `xml:"dash-length,attr,omitempty"`
	SpaceLength string `xml:"space-length,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Degree: The degree type is used to add, alter, or subtract individual notes in
// the chord. The print-object attribute can be used to keep the degree from printing
// separately when it has already taken into account in the text attribute of the kind
// element. The degree-value and degree-type text attributes specify how the value and
// type of the degree should be displayed. A harmony of kind "other" can be spelled
// explicitly by using a series of degree elements together with a root.
type Degree struct {
	Name string `xml:"-"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	DegreeValue *DegreeValue `xml:"degree-value,omitempty"`
	DegreeAlter *DegreeAlter `xml:"degree-alter,omitempty"`
	DegreeType *DegreeType `xml:"degree-type,omitempty"`
}

// DegreeAlter: The degree-alter type represents the chromatic alteration for the
// current degree. If the degree-type value is alter or subtract, the degree-alter
// value is relative to the degree already in the chord based on its kind element. If
// the degree-type value is add, the degree-alter is relative to a dominant chord
// (major and perfect intervals except for a minor seventh). The plus-minus attribute
// is used to indicate if plus and minus symbols should be used instead of sharp and
// flat symbols to display the degree alteration. It is no if not specified.
type DegreeAlter struct {
	Name string `xml:"-"`
	PlusMinus common.EnumYesNo `xml:"plus-minus,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value string `xml:",chardata"`
}

// DegreeType: The degree-type type indicates if this degree is an addition,
// alteration, or subtraction relative to the kind of the current chord. The value of
// the degree-type element affects the interpretation of the value of the degree-alter
// element. The text attribute specifies how the type of the degree should be
// displayed.
type DegreeType struct {
	Name string `xml:"-"`
	Text string `xml:"text,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value EnumDegreeTypeValue `xml:",chardata"`
}

// DegreeValue: The content of the degree-value type is a number indicating the degree
// of the chord (1 for the root, 3 for third, etc). The text attribute specifies how
// the value of the degree should be displayed. The symbol attribute indicates that a
// symbol should be used in specifying the degree. If the symbol attribute is present,
// the value of the text attribute follows the symbol.
type DegreeValue struct {
	Name string `xml:"-"`
	Symbol EnumDegreeSymbolValue `xml:"symbol,attr,omitempty"`
	Text string `xml:"text,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value int `xml:",chardata"`
}

// Direction: A direction is a musical indication that is not necessarily attached
// to a specific note. Two or more may be combined to indicate words followed by the
// start of a dashed line, the end of a wedge followed by dynamics, etc. For
// applications where a specific direction is indeed attached to a specific note, the
// direction element can be associated with the first note element that follows it in
// score order that is not in a different voice. By default, a series of direction-type
// elements and a series of child elements of a direction-type within a single
// direction element follow one another in sequence visually. For a series of
// direction-type children, non-positional formatting attributes are carried over from
// the previous element by default.
type Direction struct {
	Name string `xml:"-"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	Directive common.EnumYesNo `xml:"directive,attr,omitempty"`
	System barline.EnumSystemRelation `xml:"system,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	DirectionType []*DirectionType `xml:"direction-type,omitempty"`
	Offset *Offset `xml:"offset,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
	Voice string `xml:"voice,omitempty"`
	Staff int `xml:"staff,omitempty"`
	Sound *Sound `xml:"sound,omitempty"`
	Listening *Listening `xml:"listening,omitempty"`
}

// DirectionType: Textual direction types may have more than 1 component due to multiple
// fonts. The dynamics element may also be used in the notations element. Attribute
// groups related to print suggestions apply to the individual direction-type, not to
// the overall direction.
type DirectionType struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	Rehearsal []*common.FormattedTextId `xml:"rehearsal,omitempty"`
	Segno []*common.Segno `xml:"segno,omitempty"`
	Coda []*common.Coda `xml:"coda,omitempty"`
	Words []*common.FormattedTextId `xml:"words,omitempty"`
	Symbol []*common.FormattedSymbolId `xml:"symbol,omitempty"`
	Wedge *Wedge `xml:"wedge,omitempty"`
	Dynamics []*common.Dynamics `xml:"dynamics,omitempty"`
	Dashes *Dashes `xml:"dashes,omitempty"`
	Bracket *Bracket `xml:"bracket,omitempty"`
	Pedal *Pedal `xml:"pedal,omitempty"`
	Metronome *Metronome `xml:"metronome,omitempty"`
	OctaveShift *OctaveShift `xml:"octave-shift,omitempty"`
	HarpPedals *HarpPedals `xml:"harp-pedals,omitempty"`
	Damp *common.EmptyPrintStyleAlignId `xml:"damp,omitempty"`
	DampAll *common.EmptyPrintStyleAlignId `xml:"damp-all,omitempty"`
	Eyeglasses *common.EmptyPrintStyleAlignId `xml:"eyeglasses,omitempty"`
	StringMute *StringMute `xml:"string-mute,omitempty"`
	Scordatura *Scordatura `xml:"scordatura,omitempty"`
	Image *Image `xml:"image,omitempty"`
	PrincipalVoice *PrincipalVoice `xml:"principal-voice,omitempty"`
	Percussion []*Percussion `xml:"percussion,omitempty"`
	AccordionRegistration *AccordionRegistration `xml:"accordion-registration,omitempty"`
	StaffDivide *StaffDivide `xml:"staff-divide,omitempty"`
	OtherDirection *OtherDirection `xml:"other-direction,omitempty"`
}

// Effect: The effect type represents pictograms for sound effect percussion
// instruments. The smufl attribute is used to distinguish different SMuFL stylistic
// alternates.
type Effect struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value EnumEffectValue `xml:",chardata"`
}

// Feature: The feature type is a part of the grouping element used for musical
// analysis. The type attribute represents the type of the feature and the element
// content represents its value. This type is flexible to allow for different analyses.
type Feature struct {
	Name string `xml:"-"`
	Type string `xml:"type,attr,omitempty"`
	Value string `xml:",chardata"`
}

// FirstFret: The first-fret type indicates which fret is shown in the top space of
// the frame; it is fret 1 if the element is not present. The optional text attribute
// indicates how this is represented in the fret diagram, while the location attribute
// indicates whether the text appears to the left or right of the frame.
type FirstFret struct {
	Name string `xml:"-"`
	Text string `xml:"text,attr,omitempty"`
	Location common.EnumLeftRight `xml:"location,attr,omitempty"`
	Value int `xml:",chardata"`
}

// Frame: The frame type represents a frame or fretboard diagram used together
// with a chord symbol. The representation is based on the NIFF guitar grid with
// additional information. The frame type's unplayed attribute indicates what to
// display above a string that has no associated frame-note element. Typical values are
// x and the empty string. If the attribute is not present, the display of the unplayed
// string is application-defined.
type Frame struct {
	Name string `xml:"-"`
	Height string `xml:"height,attr,omitempty"`
	Width string `xml:"width,attr,omitempty"`
	Unplayed string `xml:"unplayed,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValignImage `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	FrameStrings int `xml:"frame-strings,omitempty"`
	FrameFrets int `xml:"frame-frets,omitempty"`
	FirstFret *FirstFret `xml:"first-fret,omitempty"`
	FrameNote []*FrameNote `xml:"frame-note,omitempty"`
}

// FrameNote: The frame-note type represents each note included in the frame. An
// open string will have a fret value of 0, while a muted string will not be associated
// with a frame-note element.
type FrameNote struct {
	Name string `xml:"-"`
	String *common.StringType `xml:"string,omitempty"`
	Fret *common.Fret `xml:"fret,omitempty"`
	Fingering *common.Fingering `xml:"fingering,omitempty"`
	Barre *Barre `xml:"barre,omitempty"`
}

// Glass: The glass type represents pictograms for glass percussion instruments.
// The smufl attribute is used to distinguish different SMuFL glyphs for wind chimes in
// the Chimes pictograms range, including those made of materials other than glass.
type Glass struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value EnumGlassValue `xml:",chardata"`
}

// Grouping: The grouping type is used for musical analysis. When the type
// attribute is "start" or "single", it usually contains one or more feature elements.
// The number attribute is used for distinguishing between overlapping and hierarchical
// groupings. The member-of attribute allows for easy distinguishing of what grouping
// elements are in what hierarchy. Feature elements contained within a "stop" type of
// grouping may be ignored. This element is flexible to allow for different types of
// analyses. Future versions of the MusicXML format may add elements that can represent
// more standardized categories of analysis data, allowing for easier data sharing.
type Grouping struct {
	Name string `xml:"-"`
	Type common.EnumStartStopSingle `xml:"type,attr,omitempty"`
	Number string `xml:"number,attr,omitempty"`
	MemberOf string `xml:"member-of,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Feature []*Feature `xml:"feature,omitempty"`
}

// Harmony: The harmony type represents harmony analysis, including chord symbols
// in popular music as well as functional harmony analysis in classical music. If there
// are alternate harmonies possible, this can be specified using multiple harmony
// elements differentiated by type. Explicit harmonies have all note present in the
// music; implied have some notes missing but implied; alternate represents alternate
// analyses. The print-object attribute controls whether or not anything is printed due
// to the harmony element. The print-frame attribute controls printing of a frame or
// fretboard diagram. The print-style attribute group sets the default for the harmony,
// but individual elements can override this with their own print-style values. The
// arrangement attribute specifies how multiple harmony-chord groups are arranged
// relative to each other. Harmony-chords with vertical arrangement are separated by
// horizontal lines. Harmony-chords with diagonal or horizontal arrangement are
// separated by diagonal lines or slashes.
type Harmony struct {
	Name string `xml:"-"`
	Type EnumHarmonyType `xml:"type,attr,omitempty"`
	PrintFrame common.EnumYesNo `xml:"print-frame,attr,omitempty"`
	Arrangement EnumHarmonyArrangement `xml:"arrangement,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Placement common.EnumAboveBelow `xml:"placement,attr,omitempty"`
	System barline.EnumSystemRelation `xml:"system,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Root []*Root `xml:"root,omitempty"`
	Numeral []*Numeral `xml:"numeral,omitempty"`
	Function []*StyleText `xml:"function,omitempty"`
	Kind []*Kind `xml:"kind,omitempty"`
	Inversion []*Inversion `xml:"inversion,omitempty"`
	Bass []*Bass `xml:"bass,omitempty"`
	Degree []*Degree `xml:"degree,omitempty"`
	Frame *Frame `xml:"frame,omitempty"`
	Offset *Offset `xml:"offset,omitempty"`
	Footnote *common.FormattedText `xml:"footnote,omitempty"`
	Level *common.Level `xml:"level,omitempty"`
	Staff int `xml:"staff,omitempty"`
}

// HarpPedals: The harp-pedals type is used to create harp pedal diagrams. The
// pedal-step and pedal-alter elements use the same values as the step and alter
// elements. For easiest reading, the pedal-tuning elements should follow standard harp
// pedal order, with pedal-step values of D, C, B, E, F, G, and A.
type HarpPedals struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	PedalTuning []*PedalTuning `xml:"pedal-tuning,omitempty"`
}

// Image: The image type is used to include graphical images in a score.
type Image struct {
	Name string `xml:"-"`
	Source string `xml:"source,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
	Height string `xml:"height,attr,omitempty"`
	Width string `xml:"width,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValignImage `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// InstrumentChange: The instrument-change element type represents a change to the virtual
// instrument sound for a given score-instrument. The id attribute refers to the
// score-instrument affected by the change. All instrument-change child elements can
// also be initially specified within the score-instrument element.
type InstrumentChange struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	InstrumentSound string `xml:"instrument-sound,omitempty"`
	Solo *common.Empty `xml:"solo,omitempty"`
	Ensemble string `xml:"ensemble,omitempty"`
	VirtualInstrument *VirtualInstrument `xml:"virtual-instrument,omitempty"`
}

// Inversion: The inversion type represents harmony inversions. The value is a
// number indicating which inversion is used: 0 for root position, 1 for first
// inversion, etc. The text attribute indicates how the inversion should be displayed
// in a score.
type Inversion struct {
	Name string `xml:"-"`
	Text string `xml:"text,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value int `xml:",chardata"`
}

// Kind: Kind indicates the type of chord. Degree elements can then add,
// subtract, or alter from these starting points The attributes are used to indicate
// the formatting of the symbol. Since the kind element is the constant in all the
// harmony-chord groups that can make up a polychord, many formatting attributes are
// here. The use-symbols attribute is yes if the kind should be represented when
// possible with harmony symbols rather than letters and numbers. These symbols
// include: major: a triangle, like Unicode 25B3 minor: -, like Unicode 002D augmented:
// +, like Unicode 002B diminished: °, like Unicode 00B0 half-diminished: ø, like
// Unicode 00F8 For the major-minor kind, only the minor symbol is used when
// use-symbols is yes. The major symbol is set using the symbol attribute in the
// degree-value element. The corresponding degree-alter value will usually be 0 in this
// case. The text attribute describes how the kind should be spelled in a score. If
// use-symbols is yes, the value of the text attribute follows the symbol. The
// stack-degrees attribute is yes if the degree elements should be stacked above each
// other. The parentheses-degrees attribute is yes if all the degrees should be in
// parentheses. The bracket-degrees attribute is yes if all the degrees should be in a
// bracket. If not specified, these values are implementation-specific. The alignment
// attributes are for the entire harmony-chord group of which this kind element is a
// part. The text attribute may use strings such as "13sus" that refer to both the kind
// and one or more degree elements. In this case, the corresponding degree elements
// should have the print-object attribute set to "no" to keep redundant alterations
// from being displayed.
type Kind struct {
	Name string `xml:"-"`
	UseSymbols common.EnumYesNo `xml:"use-symbols,attr,omitempty"`
	Text string `xml:"text,attr,omitempty"`
	StackDegrees common.EnumYesNo `xml:"stack-degrees,attr,omitempty"`
	ParenthesesDegrees common.EnumYesNo `xml:"parentheses-degrees,attr,omitempty"`
	BracketDegrees common.EnumYesNo `xml:"bracket-degrees,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Value EnumKindValue `xml:",chardata"`
}

// Listening: The listen and listening types, new in Version 4.0, specify different
// ways that a score following or machine listening application can interact with a
// performer. The listening type handles interactions that change the state of the
// listening application from the specified point in the performance onward. If
// multiple child elements of the same type are present, they should have distinct
// player and/or time-only attributes. The offset element is used to indicate that the
// listening change takes place offset from the current score position. If the
// listening element is a child of a direction element, the listening offset element
// overrides the direction offset element if both elements are present. Note that the
// offset reflects the intended musical position for the change in state. It should not
// be used to compensate for latency issues in particular hardware configurations.
type Listening struct {
	Name string `xml:"-"`
	Sync []*Sync `xml:"sync,omitempty"`
	OtherListening []*OtherListening `xml:"other-listening,omitempty"`
	Offset *Offset `xml:"offset,omitempty"`
}

// MeasureNumbering: The measure-numbering type describes how frequently measure numbers
// are displayed on this part. The text attribute from the measure element is used for
// display, or the number attribute if the text attribute is not present. Measures with
// an implicit attribute set to "yes" never display a measure number, regardless of the
// measure-numbering setting. The optional staff attribute refers to staff numbers
// within the part, from top to bottom on the system. It indicates which staff is used
// as the reference point for vertical positioning. A value of 1 is assumed if not
// present. The optional multiple-rest-always and multiple-rest-range attributes
// describe how measure numbers are shown on multiple rests when the measure-numbering
// value is not set to none. The multiple-rest-always attribute is set to yes when the
// measure number should always be shown, even if the multiple rest starts midway
// through a system when measure numbering is set to system level. The
// multiple-rest-range attribute is set to yes when measure numbers on multiple rests
// display the range of numbers for the first and last measure, rather than just the
// number of the first measure.
type MeasureNumbering struct {
	Name string `xml:"-"`
	System barline.EnumSystemRelationNumber `xml:"system,attr,omitempty"`
	Staff int `xml:"staff,attr,omitempty"`
	MultipleRestAlways common.EnumYesNo `xml:"multiple-rest-always,attr,omitempty"`
	MultipleRestRange common.EnumYesNo `xml:"multiple-rest-range,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Value EnumMeasureNumberingValue `xml:",chardata"`
}

// Membrane: The membrane type represents pictograms for membrane percussion
// instruments. The smufl attribute is used to distinguish different SMuFL stylistic
// alternates.
type Membrane struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value EnumMembraneValue `xml:",chardata"`
}

// Metal: The metal type represents pictograms for metal percussion instruments.
// The smufl attribute is used to distinguish different SMuFL stylistic alternates.
type Metal struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value EnumMetalValue `xml:",chardata"`
}

// Metronome: The metronome type represents metronome marks and other metric
// relationships. The beat-unit group and per-minute element specify regular metronome
// marks. The metronome-note and metronome-relation elements allow for the
// specification of metric modulations and other metric relationships, such as swing
// tempo marks where two eighths are equated to a quarter note / eighth note triplet.
// Tied notes can be represented in both types of metronome marks by using the
// beat-unit-tied and metronome-tied elements. The parentheses attribute indicates
// whether or not to put the metronome mark in parentheses; its value is no if not
// specified. The print-object attribute is set to no in cases where the metronome
// element represents a relationship or range that is not displayed in the music
// notation.
type Metronome struct {
	Name string `xml:"-"`
	Parentheses common.EnumYesNo `xml:"parentheses,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	Justify common.EnumLeftCenterRight `xml:"justify,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	BeatUnit []common.EnumNoteTypeValue `xml:"beat-unit,omitempty"`
	BeatUnitDot []*common.Empty `xml:"beat-unit-dot,omitempty"`
	BeatUnitTied []*BeatUnitTied `xml:"beat-unit-tied,omitempty"`
	PerMinute *PerMinute `xml:"per-minute,omitempty"`
	MetronomeArrows *common.Empty `xml:"metronome-arrows,omitempty"`
	MetronomeNote []*MetronomeNote `xml:"metronome-note,omitempty"`
	MetronomeRelation string `xml:"metronome-relation,omitempty"`
}

// MetronomeBeam: The metronome-beam type works like the beam type in defining metric
// relationships, but does not include all the attributes available in the beam type.
type MetronomeBeam struct {
	Name string `xml:"-"`
	Number int `xml:"number,attr,omitempty"`
	Value EnumBeamValue `xml:",chardata"`
}

// MetronomeNote: The metronome-note type defines the appearance of a note within a
// metric relationship mark.
type MetronomeNote struct {
	Name string `xml:"-"`
	MetronomeType common.EnumNoteTypeValue `xml:"metronome-type,omitempty"`
	MetronomeDot []*common.Empty `xml:"metronome-dot,omitempty"`
	MetronomeBeam []*MetronomeBeam `xml:"metronome-beam,omitempty"`
	MetronomeTied *MetronomeTied `xml:"metronome-tied,omitempty"`
	MetronomeTuplet *MetronomeTuplet `xml:"metronome-tuplet,omitempty"`
}

// MetronomeTied: The metronome-tied indicates the presence of a tie within a metric
// relationship mark. As with the tied element, both the start and stop of the tie
// should be specified, in this case within separate metronome-note elements.
type MetronomeTied struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
}

// MetronomeTuplet: The metronome-tuplet type uses the same element structure as the
// time-modification element along with some attributes from the tuplet element.
type MetronomeTuplet struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	Bracket common.EnumYesNo `xml:"bracket,attr,omitempty"`
	ShowNumber EnumShowTuplet `xml:"show-number,attr,omitempty"`
	ActualNotes int `xml:"actual-notes,omitempty"`
	NormalNotes int `xml:"normal-notes,omitempty"`
	NormalType common.EnumNoteTypeValue `xml:"normal-type,omitempty"`
	NormalDot []*common.Empty `xml:"normal-dot,omitempty"`
}

// Numeral: The numeral type represents the Roman numeral or Nashville number part
// of a harmony. It requires that the key be specified in the encoding, either with a
// key or numeral-key element.
type Numeral struct {
	Name string `xml:"-"`
	NumeralRoot *NumeralRoot `xml:"numeral-root,omitempty"`
	NumeralAlter *HarmonyAlter `xml:"numeral-alter,omitempty"`
	NumeralKey *NumeralKey `xml:"numeral-key,omitempty"`
}

// NumeralKey: The numeral-key type is used when the key for the numeral is different
// than the key specified by the key signature. The numeral-fifths element specifies
// the key in the same way as the fifths element. The numeral-mode element specifies
// the mode similar to the mode element, but with a restricted set of values
type NumeralKey struct {
	Name string `xml:"-"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	NumeralFifths int `xml:"numeral-fifths,omitempty"`
	NumeralMode EnumNumeralMode `xml:"numeral-mode,omitempty"`
}

// NumeralRoot: The numeral-root type represents the Roman numeral or Nashville number
// as a positive integer from 1 to 7. The text attribute indicates how the numeral
// should appear in the score. A numeral-root value of 5 with a kind of major would
// have a text attribute of "V" if displayed as a Roman numeral, and "5" if displayed
// as a Nashville number. If the text attribute is not specified, the display is
// application-dependent.
type NumeralRoot struct {
	Name string `xml:"-"`
	Text string `xml:"text,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value int `xml:",chardata"`
}

// OctaveShift: The octave shift type indicates where notes are shifted up or down
// from their true pitched values because of printing difficulty. Thus a treble clef
// line noted with 8va will be indicated with an octave-shift down from the pitch data
// indicated in the notes. A size of 8 indicates one octave; a size of 15 indicates two
// octaves.
type OctaveShift struct {
	Name string `xml:"-"`
	Type EnumUpDownStopContinue `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	Size int `xml:"size,attr,omitempty"`
	DashLength string `xml:"dash-length,attr,omitempty"`
	SpaceLength string `xml:"space-length,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Offset: An offset is represented in terms of divisions, and indicates where
// the direction will appear relative to the current musical location. The current
// musical location is always within the current measure, even at the end of a measure.
// The offset affects the visual appearance of the direction. If the sound attribute is
// "yes", then the offset affects playback and listening too. If the sound attribute is
// "no", then any sound or listening associated with the direction takes effect at the
// current location. The sound attribute is "no" by default for compatibility with
// earlier versions of the MusicXML format. If an element within a direction includes a
// default-x attribute, the offset value will be ignored when determining the
// appearance of that element.
type Offset struct {
	Name string `xml:"-"`
	Sound common.EnumYesNo `xml:"sound,attr,omitempty"`
	Value string `xml:",chardata"`
}

// OtherDirection: The other-direction type is used to define any direction symbols not
// yet in the MusicXML format. The smufl attribute can be used to specify a particular
// direction symbol, allowing application interoperability without requiring every
// SMuFL glyph to have a MusicXML element equivalent. Using the other-direction type
// without the smufl attribute allows for extended representation, though without
// application interoperability.
type OtherDirection struct {
	Name string `xml:"-"`
	PrintObject common.EnumYesNo `xml:"print-object,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

// OtherListening: The other-listening type represents other types of listening control
// and interaction. The required type attribute indicates the type of listening to
// which the element content applies. The optional player and time-only attributes
// restrict the element to apply to a single player or set of times through a repeated
// section, respectively.
type OtherListening struct {
	Name string `xml:"-"`
	Type string `xml:"type,attr,omitempty"`
	Player string `xml:"player,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Pedal: The pedal type represents piano pedal marks, including damper and
// sostenuto pedal marks. The line attribute is yes if pedal lines are used. The sign
// attribute is yes if Ped, Sost, and * signs are used. For compatibility with older
// versions, the sign attribute is yes by default if the line attribute is no, and is
// no by default if the line attribute is yes. If the sign attribute is set to yes and
// the type is start or sostenuto, the abbreviated attribute is yes if the short P and
// S signs are used, and no if the full Ped and Sost signs are used. It is no by
// default. Otherwise the abbreviated attribute is ignored. The alignment attributes
// are ignored if the sign attribute is no.
type Pedal struct {
	Name string `xml:"-"`
	Type EnumPedalType `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	Line common.EnumYesNo `xml:"line,attr,omitempty"`
	Sign common.EnumYesNo `xml:"sign,attr,omitempty"`
	Abbreviated common.EnumYesNo `xml:"abbreviated,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// PedalTuning: The pedal-tuning type specifies the tuning of a single harp pedal.
type PedalTuning struct {
	Name string `xml:"-"`
	PedalStep attributes.EnumStep `xml:"pedal-step,omitempty"`
	PedalAlter string `xml:"pedal-alter,omitempty"`
}

// PerMinute: The per-minute type can be a number, or a text description including
// numbers. If a font is specified, it overrides the font specified for the overall
// metronome element. This allows separate specification of a music font for the
// beat-unit and a text font for the numeric value, in cases where a single metronome
// font is not used.
type PerMinute struct {
	Name string `xml:"-"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Percussion: The percussion element is used to define percussion pictogram symbols.
// Definitions for these symbols can be found in Kurt Stone's "Music Notation in the
// Twentieth Century" on pages 206-212 and 223. Some values are added to these based on
// how usage has evolved in the 30 years since Stone's book was published.
type Percussion struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Enclosure common.EnumEnclosureShape `xml:"enclosure,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Glass *Glass `xml:"glass,omitempty"`
	Metal *Metal `xml:"metal,omitempty"`
	Wood *Wood `xml:"wood,omitempty"`
	Pitched *Pitched `xml:"pitched,omitempty"`
	Membrane *Membrane `xml:"membrane,omitempty"`
	Effect *Effect `xml:"effect,omitempty"`
	Timpani *Timpani `xml:"timpani,omitempty"`
	Beater *Beater `xml:"beater,omitempty"`
	Stick *Stick `xml:"stick,omitempty"`
	StickLocation EnumStickLocation `xml:"stick-location,omitempty"`
	OtherPercussion *common.OtherText `xml:"other-percussion,omitempty"`
}

// Pitched: The pitched-value type represents pictograms for pitched percussion
// instruments. The smufl attribute is used to distinguish different SMuFL glyphs for a
// particular pictogram within the Tuned mallet percussion pictograms range.
type Pitched struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value EnumPitchedValue `xml:",chardata"`
}

// PrincipalVoice: The principal-voice type represents principal and secondary voices in
// a score, either for analysis or for square bracket symbols that appear in a score.
// The element content is used for analysis and may be any text value. The symbol
// attribute indicates the type of symbol used. When used for analysis separate from
// any printed score markings, it should be set to none. Otherwise if the type is stop
// it should be set to plain.
type PrincipalVoice struct {
	Name string `xml:"-"`
	Type common.EnumStartStop `xml:"type,attr,omitempty"`
	Symbol EnumPrincipalVoiceSymbol `xml:"symbol,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

// Print: The print type contains general printing parameters, including layout
// elements. The part-name-display and part-abbreviation-display elements may also be
// used here to change how a part name or abbreviation is displayed over the course of
// a piece. They take effect when the current measure or a succeeding measure starts a
// new system. Layout group elements in a print element only apply to the current page,
// system, or staff. Music that follows continues to take the default values from the
// layout determined by the defaults element.
type Print struct {
	Name string `xml:"-"`
	StaffSpacing string `xml:"staff-spacing,attr,omitempty"`
	NewSystem common.EnumYesNo `xml:"new-system,attr,omitempty"`
	NewPage common.EnumYesNo `xml:"new-page,attr,omitempty"`
	BlankPage int `xml:"blank-page,attr,omitempty"`
	PageNumber string `xml:"page-number,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	PageLayout *common.PageLayout `xml:"page-layout,omitempty"`
	SystemLayout *common.SystemLayout `xml:"system-layout,omitempty"`
	StaffLayout []*common.StaffLayout `xml:"staff-layout,omitempty"`
	MeasureLayout *common.MeasureLayout `xml:"measure-layout,omitempty"`
	MeasureNumbering *MeasureNumbering `xml:"measure-numbering,omitempty"`
	PartNameDisplay *common.NameDisplay `xml:"part-name-display,omitempty"`
	PartAbbreviationDisplay *common.NameDisplay `xml:"part-abbreviation-display,omitempty"`
}

// Root: The root type indicates a pitch like C, D, E vs. a scale degree like
// 1, 2, 3. It is used with chord symbols in popular music. The root element has a
// root-step and optional root-alter element similar to the step and alter elements,
// but renamed to distinguish the different musical meanings.
type Root struct {
	Name string `xml:"-"`
	RootStep *RootStep `xml:"root-step,omitempty"`
	RootAlter *HarmonyAlter `xml:"root-alter,omitempty"`
}

// RootStep: The root-step type represents the pitch step of the root of the
// current chord within the harmony element. The text attribute indicates how the root
// should appear in a score if not using the element contents.
type RootStep struct {
	Name string `xml:"-"`
	Text string `xml:"text,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value attributes.EnumStep `xml:",chardata"`
}

// Scordatura: Scordatura string tunings are represented by a series of accord
// elements, similar to the staff-tuning elements. Strings are numbered from high to
// low.
type Scordatura struct {
	Name string `xml:"-"`
	Id string `xml:"id,attr,omitempty"`
	Accord []*Accord `xml:"accord,omitempty"`
}

// Sound: The sound element contains general playback parameters. They can stand
// alone within a part/measure, or be a component element within a direction. Tempo is
// expressed in quarter notes per minute. If 0, the sound-generating program should
// prompt the user at the time of compiling a sound (MIDI) file. Dynamics (or MIDI
// velocity) are expressed as a percentage of the default forte value (90 for MIDI
// 1.0). Dacapo indicates to go back to the beginning of the movement. When used it
// always has the value "yes". Segno and dalsegno are used for backwards jumps to a
// segno sign; coda and tocoda are used for forward jumps to a coda sign. If there are
// multiple jumps, the value of these parameters can be used to name and distinguish
// them. If segno or coda is used, the divisions attribute can also be used to indicate
// the number of divisions per quarter note. Otherwise sound and MIDI generating
// programs may have to recompute this. By default, a dalsegno or dacapo attribute
// indicates that the jump should occur the first time through, while a tocoda
// attribute indicates the jump should occur the second time through. The time that
// jumps occur can be changed by using the time-only attribute. The forward-repeat
// attribute indicates that a forward repeat sign is implied but not displayed. It is
// used for example in two-part forms with repeats, such as a minuet and trio where no
// repeat is displayed at the start of the trio. This usually occurs after a barline.
// When used it always has the value of "yes". The fine attribute follows the final
// note or rest in a movement with a da capo or dal segno direction. If numeric, the
// value represents the actual duration of the final note or rest, which can be
// ambiguous in written notation and different among parts and voices. The value may
// also be "yes" to indicate no change to the final duration. If the sound element
// applies only particular times through a repeat, the time-only attribute indicates
// which times to apply the sound element. Pizzicato in a sound element effects all
// following notes. Yes indicates pizzicato, no indicates arco. The pan and elevation
// attributes are deprecated in Version 2.0. The pan and elevation elements in the
// midi-instrument element should be used instead. The meaning of the pan and elevation
// attributes is the same as for the pan and elevation elements. If both are present,
// the mid-instrument elements take priority. The damper-pedal, soft-pedal, and
// sostenuto-pedal attributes effect playback of the three common piano pedals and
// their MIDI controller equivalents. The yes value indicates the pedal is depressed;
// no indicates the pedal is released. A numeric value from 0 to 100 may also be used
// for half pedaling. This value is the percentage that the pedal is depressed. A value
// of 0 is equivalent to no, and a value of 100 is equivalent to yes. Instrument
// changes, MIDI devices, MIDI instruments, and playback techniques are changed using
// the instrument-change, midi-device, midi-instrument, and play elements. When there
// are multiple instances of these elements, they should be grouped together by
// instrument using the id attribute values. The offset element is used to indicate
// that the sound takes place offset from the current score position. If the sound
// element is a child of a direction element, the sound offset element overrides the
// direction offset element if both elements are present. Note that the offset reflects
// the intended musical position for the change in sound. It should not be used to
// compensate for latency issues in particular hardware configurations.
type Sound struct {
	Name string `xml:"-"`
	Tempo string `xml:"tempo,attr,omitempty"`
	Dynamics string `xml:"dynamics,attr,omitempty"`
	Dacapo common.EnumYesNo `xml:"dacapo,attr,omitempty"`
	Segno string `xml:"segno,attr,omitempty"`
	Dalsegno string `xml:"dalsegno,attr,omitempty"`
	Coda string `xml:"coda,attr,omitempty"`
	Tocoda string `xml:"tocoda,attr,omitempty"`
	Divisions string `xml:"divisions,attr,omitempty"`
	ForwardRepeat common.EnumYesNo `xml:"forward-repeat,attr,omitempty"`
	Fine string `xml:"fine,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
	Pizzicato common.EnumYesNo `xml:"pizzicato,attr,omitempty"`
	Pan string `xml:"pan,attr,omitempty"`
	Elevation string `xml:"elevation,attr,omitempty"`
	DamperPedal string `xml:"damper-pedal,attr,omitempty"`
	SoftPedal string `xml:"soft-pedal,attr,omitempty"`
	SostenutoPedal string `xml:"sostenuto-pedal,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
	InstrumentChange []*InstrumentChange `xml:"instrument-change,omitempty"`
	MidiDevice []*common.MidiDevice `xml:"midi-device,omitempty"`
	MidiInstrument []*common.MidiInstrument `xml:"midi-instrument,omitempty"`
	Play []*common.Play `xml:"play,omitempty"`
	Swing *Swing `xml:"swing,omitempty"`
	Offset *Offset `xml:"offset,omitempty"`
}

// StaffDivide: The staff-divide element represents the staff division arrow symbols
// found at SMuFL code points U+E00B, U+E00C, and U+E00D.
type StaffDivide struct {
	Name string `xml:"-"`
	Type EnumStaffDivideSymbol `xml:"type,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Stick: The stick type represents pictograms where the material of the stick,
// mallet, or beater is included.The parentheses and dashed-circle attributes indicate
// the presence of these marks around the round beater part of a pictogram. Values for
// these attributes are "no" if not present.
type Stick struct {
	Name string `xml:"-"`
	Tip EnumTipDirection `xml:"tip,attr,omitempty"`
	Parentheses common.EnumYesNo `xml:"parentheses,attr,omitempty"`
	DashedCircle common.EnumYesNo `xml:"dashed-circle,attr,omitempty"`
	StickType EnumStickType `xml:"stick-type,omitempty"`
	StickMaterial EnumStickMaterial `xml:"stick-material,omitempty"`
}

// StringMute: The string-mute type represents string mute on and mute off symbols.
type StringMute struct {
	Name string `xml:"-"`
	Type EnumOnOff `xml:"type,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Halign common.EnumLeftCenterRight `xml:"halign,attr,omitempty"`
	Valign common.EnumValign `xml:"valign,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Swing: The swing element specifies whether or not to use swing playback,
// where consecutive on-beat / off-beat eighth or 16th notes are played with unequal
// nominal durations. The straight element specifies that no swing is present, so
// consecutive notes have equal durations. The first and second elements are positive
// integers that specify the ratio between durations of consecutive notes. For example,
// a first element with a value of 2 and a second element with a value of 1 applied to
// eighth notes specifies a quarter note / eighth note tuplet playback, where the first
// note is twice as long as the second note. Ratios should be specified with the
// smallest integers possible. For example, a ratio of 6 to 4 should be specified as 3
// to 2 instead. The optional swing-type element specifies the note type, either eighth
// or 16th, to which the ratio is applied. The value is eighth if this element is not
// present. The optional swing-style element is a string describing the style of swing
// used. The swing element has no effect for playback of grace notes, notes where a
// type element is not present, and notes where the specified duration is different
// than the nominal value associated with the specified type. If a swung note has
// attack and release attributes, those values modify the swung playback.
type Swing struct {
	Name string `xml:"-"`
	Straight *common.Empty `xml:"straight,omitempty"`
	First int `xml:"first,omitempty"`
	Second int `xml:"second,omitempty"`
	SwingType common.EnumSwingTypeValue `xml:"swing-type,omitempty"`
	SwingStyle string `xml:"swing-style,omitempty"`
}

// Sync: The sync type specifies the style that a score following application
// should use the synchronize an accompaniment with a performer. If this type is not
// included in a score, default synchronization depends on the application. The
// optional latency attribute specifies a time in milliseconds that the listening
// application should expect from the performer. The optional player and time-only
// attributes restrict the element to apply to a single player or set of times through
// a repeated section, respectively.
type Sync struct {
	Name string `xml:"-"`
	Type EnumSyncType `xml:"type,attr,omitempty"`
	Latency int `xml:"latency,attr,omitempty"`
	Player string `xml:"player,attr,omitempty"`
	TimeOnly string `xml:"time-only,attr,omitempty"`
}

// Timpani: The timpani type represents the timpani pictogram. The smufl attribute
// is used to distinguish different SMuFL stylistic alternates.
type Timpani struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
}

// Wedge: The wedge type represents crescendo and diminuendo wedge symbols. The
// type attribute is crescendo for the start of a wedge that is closed at the left
// side, and diminuendo for the start of a wedge that is closed on the right side.
// Spread values are measured in tenths; those at the start of a crescendo wedge or end
// of a diminuendo wedge are ignored. The niente attribute is yes if a circle appears
// at the point of the wedge, indicating a crescendo from nothing or diminuendo to
// nothing. It is no by default, and used only when the type is crescendo, or the type
// is stop for a wedge that began with a diminuendo type. The line-type is solid if not
// specified.
type Wedge struct {
	Name string `xml:"-"`
	Type EnumWedgeType `xml:"type,attr,omitempty"`
	Number int `xml:"number,attr,omitempty"`
	Spread string `xml:"spread,attr,omitempty"`
	Niente common.EnumYesNo `xml:"niente,attr,omitempty"`
	LineType common.EnumLineType `xml:"line-type,attr,omitempty"`
	DashLength string `xml:"dash-length,attr,omitempty"`
	SpaceLength string `xml:"space-length,attr,omitempty"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Id string `xml:"id,attr,omitempty"`
}

// Wood: The wood type represents pictograms for wood percussion instruments.
// The smufl attribute is used to distinguish different SMuFL stylistic alternates.
type Wood struct {
	Name string `xml:"-"`
	Smufl string `xml:"smufl,attr,omitempty"`
	Value EnumWoodValue `xml:",chardata"`
}

// StyleText: The style-text type represents a text element with a print-style
// attribute group.
type StyleText struct {
	Name string `xml:"-"`
	DefaultX string `xml:"default-x,attr,omitempty"`
	DefaultY string `xml:"default-y,attr,omitempty"`
	RelativeX string `xml:"relative-x,attr,omitempty"`
	RelativeY string `xml:"relative-y,attr,omitempty"`
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontStyle common.EnumFontStyle `xml:"font-style,attr,omitempty"`
	FontSize string `xml:"font-size,attr,omitempty"`
	FontWeight common.EnumFontWeight `xml:"font-weight,attr,omitempty"`
	Color string `xml:"color,attr,omitempty"`
	Value string `xml:",chardata"`
}

// TimeModification: Time modification indicates tuplets, double-note tremolos, and other
// durational changes. A time-modification element shows how the cumulative, sounding
// effect of tuplets and double-note tremolos compare to the written note type
// represented by the type and dot elements. Nested tuplets and other notations that
// use more detailed information need both the time-modification and tuplet elements to
// be represented accurately.
type TimeModification struct {
	Name string `xml:"-"`
	ActualNotes int `xml:"actual-notes,omitempty"`
	NormalNotes int `xml:"normal-notes,omitempty"`
	NormalType common.EnumNoteTypeValue `xml:"normal-type,omitempty"`
	NormalDot []*common.Empty `xml:"normal-dot,omitempty"`
}

// VirtualInstrument: The virtual-instrument element defines a specific virtual instrument
// used for an instrument sound.
type VirtualInstrument struct {
	Name string `xml:"-"`
	VirtualLibrary string `xml:"virtual-library,omitempty"`
	VirtualName string `xml:"virtual-name,omitempty"`
}

