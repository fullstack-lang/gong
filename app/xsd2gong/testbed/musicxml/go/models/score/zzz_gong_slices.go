// generated code - do not edit
package score

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Defaults
	// insertion point per field
	stage.Defaults_LyricFont_reverseMap = make(map[*LyricFont]*Defaults)
	for defaults := range stage.Defaultss {
		_ = defaults
		for _, _lyricfont := range defaults.LyricFont {
			stage.Defaults_LyricFont_reverseMap[_lyricfont] = defaults
		}
	}
	stage.Defaults_LyricLanguage_reverseMap = make(map[*LyricLanguage]*Defaults)
	for defaults := range stage.Defaultss {
		_ = defaults
		for _, _lyriclanguage := range defaults.LyricLanguage {
			stage.Defaults_LyricLanguage_reverseMap[_lyriclanguage] = defaults
		}
	}

	// Compute reverse map for named struct Measure
	// insertion point per field
	stage.Measure_MusicData_reverseMap = make(map[*MeasureMusicDataChoice]*Measure)
	for measure := range stage.Measures {
		_ = measure
		for _, _measuremusicdatachoice := range measure.MusicData {
			stage.Measure_MusicData_reverseMap[_measuremusicdatachoice] = measure
		}
	}

	// Compute reverse map for named struct Part
	// insertion point per field
	stage.Part_Measure_reverseMap = make(map[*Measure]*Part)
	for part := range stage.Parts {
		_ = part
		for _, _measure := range part.Measure {
			stage.Part_Measure_reverseMap[_measure] = part
		}
	}

	// Compute reverse map for named struct PartLink
	// insertion point per field
	stage.PartLink_InstrumentLink_reverseMap = make(map[*InstrumentLink]*PartLink)
	for partlink := range stage.PartLinks {
		_ = partlink
		for _, _instrumentlink := range partlink.InstrumentLink {
			stage.PartLink_InstrumentLink_reverseMap[_instrumentlink] = partlink
		}
	}

	// Compute reverse map for named struct PartList
	// insertion point per field
	stage.PartList_PartGroup_reverseMap = make(map[*PartGroup]*PartList)
	for partlist := range stage.PartLists {
		_ = partlist
		for _, _partgroup := range partlist.PartGroup {
			stage.PartList_PartGroup_reverseMap[_partgroup] = partlist
		}
	}
	stage.PartList_ScorePart_reverseMap = make(map[*ScorePart]*PartList)
	for partlist := range stage.PartLists {
		_ = partlist
		for _, _scorepart := range partlist.ScorePart {
			stage.PartList_ScorePart_reverseMap[_scorepart] = partlist
		}
	}

	// Compute reverse map for named struct ScorePart
	// insertion point per field
	stage.ScorePart_PartLink_reverseMap = make(map[*PartLink]*ScorePart)
	for scorepart := range stage.ScoreParts {
		_ = scorepart
		for _, _partlink := range scorepart.PartLink {
			stage.ScorePart_PartLink_reverseMap[_partlink] = scorepart
		}
	}
	stage.ScorePart_ScoreInstrument_reverseMap = make(map[*ScoreInstrument]*ScorePart)
	for scorepart := range stage.ScoreParts {
		_ = scorepart
		for _, _scoreinstrument := range scorepart.ScoreInstrument {
			stage.ScorePart_ScoreInstrument_reverseMap[_scoreinstrument] = scorepart
		}
	}
	stage.ScorePart_Player_reverseMap = make(map[*Player]*ScorePart)
	for scorepart := range stage.ScoreParts {
		_ = scorepart
		for _, _player := range scorepart.Player {
			stage.ScorePart_Player_reverseMap[_player] = scorepart
		}
	}

	// Compute reverse map for named struct ScorePartwise
	// insertion point per field
	stage.ScorePartwise_Credit_reverseMap = make(map[*Credit]*ScorePartwise)
	for scorepartwise := range stage.ScorePartwises {
		_ = scorepartwise
		for _, _credit := range scorepartwise.Credit {
			stage.ScorePartwise_Credit_reverseMap[_credit] = scorepartwise
		}
	}
	stage.ScorePartwise_Part_reverseMap = make(map[*Part]*ScorePartwise)
	for scorepartwise := range stage.ScorePartwises {
		_ = scorepartwise
		for _, _part := range scorepartwise.Part {
			stage.ScorePartwise_Part_reverseMap[_part] = scorepartwise
		}
	}

	// Compute reverse map for named struct ScoreTimewise
	// insertion point per field
	stage.ScoreTimewise_Credit_reverseMap = make(map[*Credit]*ScoreTimewise)
	for scoretimewise := range stage.ScoreTimewises {
		_ = scoretimewise
		for _, _credit := range scoretimewise.Credit {
			stage.ScoreTimewise_Credit_reverseMap[_credit] = scoretimewise
		}
	}
	stage.ScoreTimewise_Measure_reverseMap = make(map[*ScoreTimewiseMeasure]*ScoreTimewise)
	for scoretimewise := range stage.ScoreTimewises {
		_ = scoretimewise
		for _, _scoretimewisemeasure := range scoretimewise.Measure {
			stage.ScoreTimewise_Measure_reverseMap[_scoretimewisemeasure] = scoretimewise
		}
	}

	// Compute reverse map for named struct ScoreTimewiseMeasure
	// insertion point per field
	stage.ScoreTimewiseMeasure_Part_reverseMap = make(map[*ScoreTimewiseMeasurePart]*ScoreTimewiseMeasure)
	for scoretimewisemeasure := range stage.ScoreTimewiseMeasures {
		_ = scoretimewisemeasure
		for _, _scoretimewisemeasurepart := range scoretimewisemeasure.Part {
			stage.ScoreTimewiseMeasure_Part_reverseMap[_scoretimewisemeasurepart] = scoretimewisemeasure
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Credits)

	res = __gong__appendInstances(res, stage.Defaultss)

	res = __gong__appendInstances(res, stage.EmptyFonts)

	res = __gong__appendInstances(res, stage.GroupBarlines)

	res = __gong__appendInstances(res, stage.GroupNames)

	res = __gong__appendInstances(res, stage.GroupSymbols)

	res = __gong__appendInstances(res, stage.InstrumentLinks)

	res = __gong__appendInstances(res, stage.LyricFonts)

	res = __gong__appendInstances(res, stage.LyricLanguages)

	res = __gong__appendInstances(res, stage.Measures)

	res = __gong__appendInstances(res, stage.MeasureMusicDataChoices)

	res = __gong__appendInstances(res, stage.Opuss)

	res = __gong__appendInstances(res, stage.Parts)

	res = __gong__appendInstances(res, stage.PartGroups)

	res = __gong__appendInstances(res, stage.PartLinks)

	res = __gong__appendInstances(res, stage.PartLists)

	res = __gong__appendInstances(res, stage.PartNames)

	res = __gong__appendInstances(res, stage.Players)

	res = __gong__appendInstances(res, stage.ScoreInstruments)

	res = __gong__appendInstances(res, stage.ScoreParts)

	res = __gong__appendInstances(res, stage.ScorePartwises)

	res = __gong__appendInstances(res, stage.ScoreTimewises)

	res = __gong__appendInstances(res, stage.ScoreTimewiseMeasures)

	res = __gong__appendInstances(res, stage.ScoreTimewiseMeasureParts)

	res = __gong__appendInstances(res, stage.Works)

	return
}

// insertion point per named struct
func (credit *Credit) GongCopy() GongstructIF {
	newInstance := new(Credit)
	credit.GongCopyBasicFields(newInstance)
	return newInstance
}

func (defaults *Defaults) GongCopy() GongstructIF {
	newInstance := new(Defaults)
	defaults.GongCopyBasicFields(newInstance)
	return newInstance
}

func (emptyfont *EmptyFont) GongCopy() GongstructIF {
	newInstance := new(EmptyFont)
	emptyfont.GongCopyBasicFields(newInstance)
	return newInstance
}

func (groupbarline *GroupBarline) GongCopy() GongstructIF {
	newInstance := new(GroupBarline)
	groupbarline.GongCopyBasicFields(newInstance)
	return newInstance
}

func (groupname *GroupName) GongCopy() GongstructIF {
	newInstance := new(GroupName)
	groupname.GongCopyBasicFields(newInstance)
	return newInstance
}

func (groupsymbol *GroupSymbol) GongCopy() GongstructIF {
	newInstance := new(GroupSymbol)
	groupsymbol.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrumentlink *InstrumentLink) GongCopy() GongstructIF {
	newInstance := new(InstrumentLink)
	instrumentlink.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lyricfont *LyricFont) GongCopy() GongstructIF {
	newInstance := new(LyricFont)
	lyricfont.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lyriclanguage *LyricLanguage) GongCopy() GongstructIF {
	newInstance := new(LyricLanguage)
	lyriclanguage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measure *Measure) GongCopy() GongstructIF {
	newInstance := new(Measure)
	measure.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongCopy() GongstructIF {
	newInstance := new(MeasureMusicDataChoice)
	measuremusicdatachoice.GongCopyBasicFields(newInstance)
	return newInstance
}

func (opus *Opus) GongCopy() GongstructIF {
	newInstance := new(Opus)
	opus.GongCopyBasicFields(newInstance)
	return newInstance
}

func (part *Part) GongCopy() GongstructIF {
	newInstance := new(Part)
	part.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partgroup *PartGroup) GongCopy() GongstructIF {
	newInstance := new(PartGroup)
	partgroup.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partlink *PartLink) GongCopy() GongstructIF {
	newInstance := new(PartLink)
	partlink.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partlist *PartList) GongCopy() GongstructIF {
	newInstance := new(PartList)
	partlist.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partname *PartName) GongCopy() GongstructIF {
	newInstance := new(PartName)
	partname.GongCopyBasicFields(newInstance)
	return newInstance
}

func (player *Player) GongCopy() GongstructIF {
	newInstance := new(Player)
	player.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scoreinstrument *ScoreInstrument) GongCopy() GongstructIF {
	newInstance := new(ScoreInstrument)
	scoreinstrument.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scorepart *ScorePart) GongCopy() GongstructIF {
	newInstance := new(ScorePart)
	scorepart.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scorepartwise *ScorePartwise) GongCopy() GongstructIF {
	newInstance := new(ScorePartwise)
	scorepartwise.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scoretimewise *ScoreTimewise) GongCopy() GongstructIF {
	newInstance := new(ScoreTimewise)
	scoretimewise.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongCopy() GongstructIF {
	newInstance := new(ScoreTimewiseMeasure)
	scoretimewisemeasure.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongCopy() GongstructIF {
	newInstance := new(ScoreTimewiseMeasurePart)
	scoretimewisemeasurepart.GongCopyBasicFields(newInstance)
	return newInstance
}

func (work *Work) GongCopy() GongstructIF {
	newInstance := new(Work)
	work.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (credit *Credit) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, credit)
}

func (defaults *Defaults) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, defaults)
}

func (emptyfont *EmptyFont) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, emptyfont)
}

func (groupbarline *GroupBarline) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, groupbarline)
}

func (groupname *GroupName) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, groupname)
}

func (groupsymbol *GroupSymbol) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, groupsymbol)
}

func (instrumentlink *InstrumentLink) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrumentlink)
}

func (lyricfont *LyricFont) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lyricfont)
}

func (lyriclanguage *LyricLanguage) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lyriclanguage)
}

func (measure *Measure) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measure)
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measuremusicdatachoice)
}

func (opus *Opus) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, opus)
}

func (part *Part) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, part)
}

func (partgroup *PartGroup) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partgroup)
}

func (partlink *PartLink) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partlink)
}

func (partlist *PartList) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partlist)
}

func (partname *PartName) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partname)
}

func (player *Player) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, player)
}

func (scoreinstrument *ScoreInstrument) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scoreinstrument)
}

func (scorepart *ScorePart) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scorepart)
}

func (scorepartwise *ScorePartwise) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scorepartwise)
}

func (scoretimewise *ScoreTimewise) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scoretimewise)
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scoretimewisemeasure)
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scoretimewisemeasurepart)
}

func (work *Work) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, work)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.Credits,
		stage.Credit_stagedOrder,
		stage.Credits_reference,
		&stage.Credits_referenceOrder,
		stage.Credits_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Defaultss,
		stage.Defaults_stagedOrder,
		stage.Defaultss_reference,
		&stage.Defaultss_referenceOrder,
		stage.Defaultss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.EmptyFonts,
		stage.EmptyFont_stagedOrder,
		stage.EmptyFonts_reference,
		&stage.EmptyFonts_referenceOrder,
		stage.EmptyFonts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GroupBarlines,
		stage.GroupBarline_stagedOrder,
		stage.GroupBarlines_reference,
		&stage.GroupBarlines_referenceOrder,
		stage.GroupBarlines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GroupNames,
		stage.GroupName_stagedOrder,
		stage.GroupNames_reference,
		&stage.GroupNames_referenceOrder,
		stage.GroupNames_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GroupSymbols,
		stage.GroupSymbol_stagedOrder,
		stage.GroupSymbols_reference,
		&stage.GroupSymbols_referenceOrder,
		stage.GroupSymbols_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InstrumentLinks,
		stage.InstrumentLink_stagedOrder,
		stage.InstrumentLinks_reference,
		&stage.InstrumentLinks_referenceOrder,
		stage.InstrumentLinks_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LyricFonts,
		stage.LyricFont_stagedOrder,
		stage.LyricFonts_reference,
		&stage.LyricFonts_referenceOrder,
		stage.LyricFonts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LyricLanguages,
		stage.LyricLanguage_stagedOrder,
		stage.LyricLanguages_reference,
		&stage.LyricLanguages_referenceOrder,
		stage.LyricLanguages_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Measures,
		stage.Measure_stagedOrder,
		stage.Measures_reference,
		&stage.Measures_referenceOrder,
		stage.Measures_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MeasureMusicDataChoices,
		stage.MeasureMusicDataChoice_stagedOrder,
		stage.MeasureMusicDataChoices_reference,
		&stage.MeasureMusicDataChoices_referenceOrder,
		stage.MeasureMusicDataChoices_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Opuss,
		stage.Opus_stagedOrder,
		stage.Opuss_reference,
		&stage.Opuss_referenceOrder,
		stage.Opuss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Parts,
		stage.Part_stagedOrder,
		stage.Parts_reference,
		&stage.Parts_referenceOrder,
		stage.Parts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PartGroups,
		stage.PartGroup_stagedOrder,
		stage.PartGroups_reference,
		&stage.PartGroups_referenceOrder,
		stage.PartGroups_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PartLinks,
		stage.PartLink_stagedOrder,
		stage.PartLinks_reference,
		&stage.PartLinks_referenceOrder,
		stage.PartLinks_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PartLists,
		stage.PartList_stagedOrder,
		stage.PartLists_reference,
		&stage.PartLists_referenceOrder,
		stage.PartLists_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PartNames,
		stage.PartName_stagedOrder,
		stage.PartNames_reference,
		&stage.PartNames_referenceOrder,
		stage.PartNames_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Players,
		stage.Player_stagedOrder,
		stage.Players_reference,
		&stage.Players_referenceOrder,
		stage.Players_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ScoreInstruments,
		stage.ScoreInstrument_stagedOrder,
		stage.ScoreInstruments_reference,
		&stage.ScoreInstruments_referenceOrder,
		stage.ScoreInstruments_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ScoreParts,
		stage.ScorePart_stagedOrder,
		stage.ScoreParts_reference,
		&stage.ScoreParts_referenceOrder,
		stage.ScoreParts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ScorePartwises,
		stage.ScorePartwise_stagedOrder,
		stage.ScorePartwises_reference,
		&stage.ScorePartwises_referenceOrder,
		stage.ScorePartwises_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ScoreTimewises,
		stage.ScoreTimewise_stagedOrder,
		stage.ScoreTimewises_reference,
		&stage.ScoreTimewises_referenceOrder,
		stage.ScoreTimewises_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ScoreTimewiseMeasures,
		stage.ScoreTimewiseMeasure_stagedOrder,
		stage.ScoreTimewiseMeasures_reference,
		&stage.ScoreTimewiseMeasures_referenceOrder,
		stage.ScoreTimewiseMeasures_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ScoreTimewiseMeasureParts,
		stage.ScoreTimewiseMeasurePart_stagedOrder,
		stage.ScoreTimewiseMeasureParts_reference,
		&stage.ScoreTimewiseMeasureParts_referenceOrder,
		stage.ScoreTimewiseMeasureParts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Works,
		stage.Work_stagedOrder,
		stage.Works_reference,
		&stage.Works_referenceOrder,
		stage.Works_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.Credits, &stage.Credits_reference, &stage.Credits_referenceOrder, &stage.Credits_instance)

	__gong__computeReferencePass1(stage, stage.Defaultss, &stage.Defaultss_reference, &stage.Defaultss_referenceOrder, &stage.Defaultss_instance)

	__gong__computeReferencePass1(stage, stage.EmptyFonts, &stage.EmptyFonts_reference, &stage.EmptyFonts_referenceOrder, &stage.EmptyFonts_instance)

	__gong__computeReferencePass1(stage, stage.GroupBarlines, &stage.GroupBarlines_reference, &stage.GroupBarlines_referenceOrder, &stage.GroupBarlines_instance)

	__gong__computeReferencePass1(stage, stage.GroupNames, &stage.GroupNames_reference, &stage.GroupNames_referenceOrder, &stage.GroupNames_instance)

	__gong__computeReferencePass1(stage, stage.GroupSymbols, &stage.GroupSymbols_reference, &stage.GroupSymbols_referenceOrder, &stage.GroupSymbols_instance)

	__gong__computeReferencePass1(stage, stage.InstrumentLinks, &stage.InstrumentLinks_reference, &stage.InstrumentLinks_referenceOrder, &stage.InstrumentLinks_instance)

	__gong__computeReferencePass1(stage, stage.LyricFonts, &stage.LyricFonts_reference, &stage.LyricFonts_referenceOrder, &stage.LyricFonts_instance)

	__gong__computeReferencePass1(stage, stage.LyricLanguages, &stage.LyricLanguages_reference, &stage.LyricLanguages_referenceOrder, &stage.LyricLanguages_instance)

	__gong__computeReferencePass1(stage, stage.Measures, &stage.Measures_reference, &stage.Measures_referenceOrder, &stage.Measures_instance)

	__gong__computeReferencePass1(stage, stage.MeasureMusicDataChoices, &stage.MeasureMusicDataChoices_reference, &stage.MeasureMusicDataChoices_referenceOrder, &stage.MeasureMusicDataChoices_instance)

	__gong__computeReferencePass1(stage, stage.Opuss, &stage.Opuss_reference, &stage.Opuss_referenceOrder, &stage.Opuss_instance)

	__gong__computeReferencePass1(stage, stage.Parts, &stage.Parts_reference, &stage.Parts_referenceOrder, &stage.Parts_instance)

	__gong__computeReferencePass1(stage, stage.PartGroups, &stage.PartGroups_reference, &stage.PartGroups_referenceOrder, &stage.PartGroups_instance)

	__gong__computeReferencePass1(stage, stage.PartLinks, &stage.PartLinks_reference, &stage.PartLinks_referenceOrder, &stage.PartLinks_instance)

	__gong__computeReferencePass1(stage, stage.PartLists, &stage.PartLists_reference, &stage.PartLists_referenceOrder, &stage.PartLists_instance)

	__gong__computeReferencePass1(stage, stage.PartNames, &stage.PartNames_reference, &stage.PartNames_referenceOrder, &stage.PartNames_instance)

	__gong__computeReferencePass1(stage, stage.Players, &stage.Players_reference, &stage.Players_referenceOrder, &stage.Players_instance)

	__gong__computeReferencePass1(stage, stage.ScoreInstruments, &stage.ScoreInstruments_reference, &stage.ScoreInstruments_referenceOrder, &stage.ScoreInstruments_instance)

	__gong__computeReferencePass1(stage, stage.ScoreParts, &stage.ScoreParts_reference, &stage.ScoreParts_referenceOrder, &stage.ScoreParts_instance)

	__gong__computeReferencePass1(stage, stage.ScorePartwises, &stage.ScorePartwises_reference, &stage.ScorePartwises_referenceOrder, &stage.ScorePartwises_instance)

	__gong__computeReferencePass1(stage, stage.ScoreTimewises, &stage.ScoreTimewises_reference, &stage.ScoreTimewises_referenceOrder, &stage.ScoreTimewises_instance)

	__gong__computeReferencePass1(stage, stage.ScoreTimewiseMeasures, &stage.ScoreTimewiseMeasures_reference, &stage.ScoreTimewiseMeasures_referenceOrder, &stage.ScoreTimewiseMeasures_instance)

	__gong__computeReferencePass1(stage, stage.ScoreTimewiseMeasureParts, &stage.ScoreTimewiseMeasureParts_reference, &stage.ScoreTimewiseMeasureParts_referenceOrder, &stage.ScoreTimewiseMeasureParts_instance)

	__gong__computeReferencePass1(stage, stage.Works, &stage.Works_reference, &stage.Works_referenceOrder, &stage.Works_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Credits, stage.Credits_reference, stage)

	__gong__computeReferencePass2(stage.Defaultss, stage.Defaultss_reference, stage)

	__gong__computeReferencePass2(stage.EmptyFonts, stage.EmptyFonts_reference, stage)

	__gong__computeReferencePass2(stage.GroupBarlines, stage.GroupBarlines_reference, stage)

	__gong__computeReferencePass2(stage.GroupNames, stage.GroupNames_reference, stage)

	__gong__computeReferencePass2(stage.GroupSymbols, stage.GroupSymbols_reference, stage)

	__gong__computeReferencePass2(stage.InstrumentLinks, stage.InstrumentLinks_reference, stage)

	__gong__computeReferencePass2(stage.LyricFonts, stage.LyricFonts_reference, stage)

	__gong__computeReferencePass2(stage.LyricLanguages, stage.LyricLanguages_reference, stage)

	__gong__computeReferencePass2(stage.Measures, stage.Measures_reference, stage)

	__gong__computeReferencePass2(stage.MeasureMusicDataChoices, stage.MeasureMusicDataChoices_reference, stage)

	__gong__computeReferencePass2(stage.Opuss, stage.Opuss_reference, stage)

	__gong__computeReferencePass2(stage.Parts, stage.Parts_reference, stage)

	__gong__computeReferencePass2(stage.PartGroups, stage.PartGroups_reference, stage)

	__gong__computeReferencePass2(stage.PartLinks, stage.PartLinks_reference, stage)

	__gong__computeReferencePass2(stage.PartLists, stage.PartLists_reference, stage)

	__gong__computeReferencePass2(stage.PartNames, stage.PartNames_reference, stage)

	__gong__computeReferencePass2(stage.Players, stage.Players_reference, stage)

	__gong__computeReferencePass2(stage.ScoreInstruments, stage.ScoreInstruments_reference, stage)

	__gong__computeReferencePass2(stage.ScoreParts, stage.ScoreParts_reference, stage)

	__gong__computeReferencePass2(stage.ScorePartwises, stage.ScorePartwises_reference, stage)

	__gong__computeReferencePass2(stage.ScoreTimewises, stage.ScoreTimewises_reference, stage)

	__gong__computeReferencePass2(stage.ScoreTimewiseMeasures, stage.ScoreTimewiseMeasures_reference, stage)

	__gong__computeReferencePass2(stage.ScoreTimewiseMeasureParts, stage.ScoreTimewiseMeasureParts_reference, stage)

	__gong__computeReferencePass2(stage.Works, stage.Works_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (credit *Credit) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Credit_stagedOrder, stage.Credits_referenceOrder, credit, "Credit")
}

func (defaults *Defaults) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Defaults_stagedOrder, stage.Defaultss_referenceOrder, defaults, "Defaults")
}

func (emptyfont *EmptyFont) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EmptyFont_stagedOrder, stage.EmptyFonts_referenceOrder, emptyfont, "EmptyFont")
}

func (groupbarline *GroupBarline) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroupBarline_stagedOrder, stage.GroupBarlines_referenceOrder, groupbarline, "GroupBarline")
}

func (groupname *GroupName) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroupName_stagedOrder, stage.GroupNames_referenceOrder, groupname, "GroupName")
}

func (groupsymbol *GroupSymbol) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroupSymbol_stagedOrder, stage.GroupSymbols_referenceOrder, groupsymbol, "GroupSymbol")
}

func (instrumentlink *InstrumentLink) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InstrumentLink_stagedOrder, stage.InstrumentLinks_referenceOrder, instrumentlink, "InstrumentLink")
}

func (lyricfont *LyricFont) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LyricFont_stagedOrder, stage.LyricFonts_referenceOrder, lyricfont, "LyricFont")
}

func (lyriclanguage *LyricLanguage) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LyricLanguage_stagedOrder, stage.LyricLanguages_referenceOrder, lyriclanguage, "LyricLanguage")
}

func (measure *Measure) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Measure_stagedOrder, stage.Measures_referenceOrder, measure, "Measure")
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeasureMusicDataChoice_stagedOrder, stage.MeasureMusicDataChoices_referenceOrder, measuremusicdatachoice, "MeasureMusicDataChoice")
}

func (opus *Opus) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Opus_stagedOrder, stage.Opuss_referenceOrder, opus, "Opus")
}

func (part *Part) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Part_stagedOrder, stage.Parts_referenceOrder, part, "Part")
}

func (partgroup *PartGroup) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartGroup_stagedOrder, stage.PartGroups_referenceOrder, partgroup, "PartGroup")
}

func (partlink *PartLink) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartLink_stagedOrder, stage.PartLinks_referenceOrder, partlink, "PartLink")
}

func (partlist *PartList) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartList_stagedOrder, stage.PartLists_referenceOrder, partlist, "PartList")
}

func (partname *PartName) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartName_stagedOrder, stage.PartNames_referenceOrder, partname, "PartName")
}

func (player *Player) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Player_stagedOrder, stage.Players_referenceOrder, player, "Player")
}

func (scoreinstrument *ScoreInstrument) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ScoreInstrument_stagedOrder, stage.ScoreInstruments_referenceOrder, scoreinstrument, "ScoreInstrument")
}

func (scorepart *ScorePart) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ScorePart_stagedOrder, stage.ScoreParts_referenceOrder, scorepart, "ScorePart")
}

func (scorepartwise *ScorePartwise) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ScorePartwise_stagedOrder, stage.ScorePartwises_referenceOrder, scorepartwise, "ScorePartwise")
}

func (scoretimewise *ScoreTimewise) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ScoreTimewise_stagedOrder, stage.ScoreTimewises_referenceOrder, scoretimewise, "ScoreTimewise")
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ScoreTimewiseMeasure_stagedOrder, stage.ScoreTimewiseMeasures_referenceOrder, scoretimewisemeasure, "ScoreTimewiseMeasure")
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ScoreTimewiseMeasurePart_stagedOrder, stage.ScoreTimewiseMeasureParts_referenceOrder, scoretimewisemeasurepart, "ScoreTimewiseMeasurePart")
}

func (work *Work) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Work_stagedOrder, stage.Works_referenceOrder, work, "Work")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (credit *Credit) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(credit, credit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (credit *Credit) GongGetReferenceIdentifier(stage *Stage) string {
	return credit.GongGetIdentifier(stage)
}

func (defaults *Defaults) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(defaults, defaults.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (defaults *Defaults) GongGetReferenceIdentifier(stage *Stage) string {
	return defaults.GongGetIdentifier(stage)
}

func (emptyfont *EmptyFont) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(emptyfont, emptyfont.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (emptyfont *EmptyFont) GongGetReferenceIdentifier(stage *Stage) string {
	return emptyfont.GongGetIdentifier(stage)
}

func (groupbarline *GroupBarline) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(groupbarline, groupbarline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groupbarline *GroupBarline) GongGetReferenceIdentifier(stage *Stage) string {
	return groupbarline.GongGetIdentifier(stage)
}

func (groupname *GroupName) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(groupname, groupname.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groupname *GroupName) GongGetReferenceIdentifier(stage *Stage) string {
	return groupname.GongGetIdentifier(stage)
}

func (groupsymbol *GroupSymbol) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(groupsymbol, groupsymbol.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groupsymbol *GroupSymbol) GongGetReferenceIdentifier(stage *Stage) string {
	return groupsymbol.GongGetIdentifier(stage)
}

func (instrumentlink *InstrumentLink) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrumentlink, instrumentlink.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrumentlink *InstrumentLink) GongGetReferenceIdentifier(stage *Stage) string {
	return instrumentlink.GongGetIdentifier(stage)
}

func (lyricfont *LyricFont) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lyricfont, lyricfont.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lyricfont *LyricFont) GongGetReferenceIdentifier(stage *Stage) string {
	return lyricfont.GongGetIdentifier(stage)
}

func (lyriclanguage *LyricLanguage) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lyriclanguage, lyriclanguage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lyriclanguage *LyricLanguage) GongGetReferenceIdentifier(stage *Stage) string {
	return lyriclanguage.GongGetIdentifier(stage)
}

func (measure *Measure) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measure, measure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measure *Measure) GongGetReferenceIdentifier(stage *Stage) string {
	return measure.GongGetIdentifier(stage)
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measuremusicdatachoice, measuremusicdatachoice.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measuremusicdatachoice *MeasureMusicDataChoice) GongGetReferenceIdentifier(stage *Stage) string {
	return measuremusicdatachoice.GongGetIdentifier(stage)
}

func (opus *Opus) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(opus, opus.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (opus *Opus) GongGetReferenceIdentifier(stage *Stage) string {
	return opus.GongGetIdentifier(stage)
}

func (part *Part) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(part, part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part *Part) GongGetReferenceIdentifier(stage *Stage) string {
	return part.GongGetIdentifier(stage)
}

func (partgroup *PartGroup) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partgroup, partgroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partgroup *PartGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return partgroup.GongGetIdentifier(stage)
}

func (partlink *PartLink) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partlink, partlink.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partlink *PartLink) GongGetReferenceIdentifier(stage *Stage) string {
	return partlink.GongGetIdentifier(stage)
}

func (partlist *PartList) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partlist, partlist.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partlist *PartList) GongGetReferenceIdentifier(stage *Stage) string {
	return partlist.GongGetIdentifier(stage)
}

func (partname *PartName) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partname, partname.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partname *PartName) GongGetReferenceIdentifier(stage *Stage) string {
	return partname.GongGetIdentifier(stage)
}

func (player *Player) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(player, player.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (player *Player) GongGetReferenceIdentifier(stage *Stage) string {
	return player.GongGetIdentifier(stage)
}

func (scoreinstrument *ScoreInstrument) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scoreinstrument, scoreinstrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scoreinstrument *ScoreInstrument) GongGetReferenceIdentifier(stage *Stage) string {
	return scoreinstrument.GongGetIdentifier(stage)
}

func (scorepart *ScorePart) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scorepart, scorepart.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scorepart *ScorePart) GongGetReferenceIdentifier(stage *Stage) string {
	return scorepart.GongGetIdentifier(stage)
}

func (scorepartwise *ScorePartwise) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scorepartwise, scorepartwise.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scorepartwise *ScorePartwise) GongGetReferenceIdentifier(stage *Stage) string {
	return scorepartwise.GongGetIdentifier(stage)
}

func (scoretimewise *ScoreTimewise) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scoretimewise, scoretimewise.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scoretimewise *ScoreTimewise) GongGetReferenceIdentifier(stage *Stage) string {
	return scoretimewise.GongGetIdentifier(stage)
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scoretimewisemeasure, scoretimewisemeasure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scoretimewisemeasure *ScoreTimewiseMeasure) GongGetReferenceIdentifier(stage *Stage) string {
	return scoretimewisemeasure.GongGetIdentifier(stage)
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scoretimewisemeasurepart, scoretimewisemeasurepart.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongGetReferenceIdentifier(stage *Stage) string {
	return scoretimewisemeasurepart.GongGetIdentifier(stage)
}

func (work *Work) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(work, work.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (work *Work) GongGetReferenceIdentifier(stage *Stage) string {
	return work.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (credit *Credit) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(credit.GongGetIdentifier(stage), "Credit", credit.Name)
}

func (defaults *Defaults) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(defaults.GongGetIdentifier(stage), "Defaults", defaults.Name)
}

func (emptyfont *EmptyFont) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(emptyfont.GongGetIdentifier(stage), "EmptyFont", emptyfont.Name)
}

func (groupbarline *GroupBarline) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(groupbarline.GongGetIdentifier(stage), "GroupBarline", groupbarline.Name)
}

func (groupname *GroupName) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(groupname.GongGetIdentifier(stage), "GroupName", groupname.Name)
}

func (groupsymbol *GroupSymbol) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(groupsymbol.GongGetIdentifier(stage), "GroupSymbol", groupsymbol.Name)
}

func (instrumentlink *InstrumentLink) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrumentlink.GongGetIdentifier(stage), "InstrumentLink", instrumentlink.Name)
}

func (lyricfont *LyricFont) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lyricfont.GongGetIdentifier(stage), "LyricFont", lyricfont.Name)
}

func (lyriclanguage *LyricLanguage) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lyriclanguage.GongGetIdentifier(stage), "LyricLanguage", lyriclanguage.Name)
}

func (measure *Measure) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measure.GongGetIdentifier(stage), "Measure", measure.Name)
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measuremusicdatachoice.GongGetIdentifier(stage), "MeasureMusicDataChoice", measuremusicdatachoice.Name)
}

func (opus *Opus) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(opus.GongGetIdentifier(stage), "Opus", opus.Name)
}

func (part *Part) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(part.GongGetIdentifier(stage), "Part", part.Name)
}

func (partgroup *PartGroup) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partgroup.GongGetIdentifier(stage), "PartGroup", partgroup.Name)
}

func (partlink *PartLink) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partlink.GongGetIdentifier(stage), "PartLink", partlink.Name)
}

func (partlist *PartList) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partlist.GongGetIdentifier(stage), "PartList", partlist.Name)
}

func (partname *PartName) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partname.GongGetIdentifier(stage), "PartName", partname.Name)
}

func (player *Player) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(player.GongGetIdentifier(stage), "Player", player.Name)
}

func (scoreinstrument *ScoreInstrument) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scoreinstrument.GongGetIdentifier(stage), "ScoreInstrument", scoreinstrument.Name)
}

func (scorepart *ScorePart) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scorepart.GongGetIdentifier(stage), "ScorePart", scorepart.Name)
}

func (scorepartwise *ScorePartwise) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scorepartwise.GongGetIdentifier(stage), "ScorePartwise", scorepartwise.Name)
}

func (scoretimewise *ScoreTimewise) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scoretimewise.GongGetIdentifier(stage), "ScoreTimewise", scoretimewise.Name)
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scoretimewisemeasure.GongGetIdentifier(stage), "ScoreTimewiseMeasure", scoretimewisemeasure.Name)
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scoretimewisemeasurepart.GongGetIdentifier(stage), "ScoreTimewiseMeasurePart", scoretimewisemeasurepart.Name)
}

func (work *Work) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(work.GongGetIdentifier(stage), "Work", work.Name)
}

// insertion point for unstaging
func (credit *Credit) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(credit.GongGetReferenceIdentifier(stage))
}

func (defaults *Defaults) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(defaults.GongGetReferenceIdentifier(stage))
}

func (emptyfont *EmptyFont) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(emptyfont.GongGetReferenceIdentifier(stage))
}

func (groupbarline *GroupBarline) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(groupbarline.GongGetReferenceIdentifier(stage))
}

func (groupname *GroupName) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(groupname.GongGetReferenceIdentifier(stage))
}

func (groupsymbol *GroupSymbol) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(groupsymbol.GongGetReferenceIdentifier(stage))
}

func (instrumentlink *InstrumentLink) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrumentlink.GongGetReferenceIdentifier(stage))
}

func (lyricfont *LyricFont) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lyricfont.GongGetReferenceIdentifier(stage))
}

func (lyriclanguage *LyricLanguage) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lyriclanguage.GongGetReferenceIdentifier(stage))
}

func (measure *Measure) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measure.GongGetReferenceIdentifier(stage))
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measuremusicdatachoice.GongGetReferenceIdentifier(stage))
}

func (opus *Opus) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(opus.GongGetReferenceIdentifier(stage))
}

func (part *Part) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(part.GongGetReferenceIdentifier(stage))
}

func (partgroup *PartGroup) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partgroup.GongGetReferenceIdentifier(stage))
}

func (partlink *PartLink) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partlink.GongGetReferenceIdentifier(stage))
}

func (partlist *PartList) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partlist.GongGetReferenceIdentifier(stage))
}

func (partname *PartName) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partname.GongGetReferenceIdentifier(stage))
}

func (player *Player) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(player.GongGetReferenceIdentifier(stage))
}

func (scoreinstrument *ScoreInstrument) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scoreinstrument.GongGetReferenceIdentifier(stage))
}

func (scorepart *ScorePart) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scorepart.GongGetReferenceIdentifier(stage))
}

func (scorepartwise *ScorePartwise) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scorepartwise.GongGetReferenceIdentifier(stage))
}

func (scoretimewise *ScoreTimewise) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scoretimewise.GongGetReferenceIdentifier(stage))
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scoretimewisemeasure.GongGetReferenceIdentifier(stage))
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scoretimewisemeasurepart.GongGetReferenceIdentifier(stage))
}

func (work *Work) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(work.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
