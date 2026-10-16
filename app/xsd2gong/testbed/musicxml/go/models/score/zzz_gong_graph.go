// generated code - do not edit
package score

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (credit *Credit) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Credits[credit]
	return ok
}

func (defaults *Defaults) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Defaultss[defaults]
	return ok
}

func (emptyfont *EmptyFont) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EmptyFonts[emptyfont]
	return ok
}

func (groupbarline *GroupBarline) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GroupBarlines[groupbarline]
	return ok
}

func (groupname *GroupName) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GroupNames[groupname]
	return ok
}

func (groupsymbol *GroupSymbol) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GroupSymbols[groupsymbol]
	return ok
}

func (instrumentlink *InstrumentLink) GongIsStaged(stage *Stage) bool {
	_, ok := stage.InstrumentLinks[instrumentlink]
	return ok
}

func (lyricfont *LyricFont) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LyricFonts[lyricfont]
	return ok
}

func (lyriclanguage *LyricLanguage) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LyricLanguages[lyriclanguage]
	return ok
}

func (measure *Measure) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Measures[measure]
	return ok
}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeasureMusicDataChoices[measuremusicdatachoice]
	return ok
}

func (opus *Opus) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Opuss[opus]
	return ok
}

func (part *Part) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Parts[part]
	return ok
}

func (partgroup *PartGroup) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartGroups[partgroup]
	return ok
}

func (partlink *PartLink) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartLinks[partlink]
	return ok
}

func (partlist *PartList) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartLists[partlist]
	return ok
}

func (partname *PartName) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartNames[partname]
	return ok
}

func (player *Player) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Players[player]
	return ok
}

func (scoreinstrument *ScoreInstrument) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ScoreInstruments[scoreinstrument]
	return ok
}

func (scorepart *ScorePart) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ScoreParts[scorepart]
	return ok
}

func (scorepartwise *ScorePartwise) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ScorePartwises[scorepartwise]
	return ok
}

func (scoretimewise *ScoreTimewise) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ScoreTimewises[scoretimewise]
	return ok
}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ScoreTimewiseMeasures[scoretimewisemeasure]
	return ok
}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ScoreTimewiseMeasureParts[scoretimewisemeasurepart]
	return ok
}

func (work *Work) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Works[work]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (credit *Credit) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(credit) {
		return
	}

	credit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (defaults *Defaults) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(defaults) {
		return
	}

	defaults.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if defaults.MusicFont != nil {
		stage.StageBranch(defaults.MusicFont)
	}
	if defaults.WordFont != nil {
		stage.StageBranch(defaults.WordFont)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lyricfont := range defaults.LyricFont {
		stage.StageBranch(_lyricfont)
	}
	for _, _lyriclanguage := range defaults.LyricLanguage {
		stage.StageBranch(_lyriclanguage)
	}

}

func (emptyfont *EmptyFont) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(emptyfont) {
		return
	}

	emptyfont.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (groupbarline *GroupBarline) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(groupbarline) {
		return
	}

	groupbarline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (groupname *GroupName) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(groupname) {
		return
	}

	groupname.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (groupsymbol *GroupSymbol) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(groupsymbol) {
		return
	}

	groupsymbol.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrumentlink *InstrumentLink) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(instrumentlink) {
		return
	}

	instrumentlink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if instrumentlink.RefScoreInstrument != nil {
		stage.StageBranch(instrumentlink.RefScoreInstrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lyricfont *LyricFont) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lyricfont) {
		return
	}

	lyricfont.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lyriclanguage *LyricLanguage) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lyriclanguage) {
		return
	}

	lyriclanguage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure *Measure) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measure) {
		return
	}

	measure.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _measuremusicdatachoice := range measure.MusicData {
		stage.StageBranch(_measuremusicdatachoice)
	}

}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measuremusicdatachoice) {
		return
	}

	measuremusicdatachoice.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (opus *Opus) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(opus) {
		return
	}

	opus.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part *Part) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(part) {
		return
	}

	part.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part.RefScorePart != nil {
		stage.StageBranch(part.RefScorePart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _measure := range part.Measure {
		stage.StageBranch(_measure)
	}

}

func (partgroup *PartGroup) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partgroup) {
		return
	}

	partgroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if partgroup.GroupName != nil {
		stage.StageBranch(partgroup.GroupName)
	}
	if partgroup.GroupAbbreviation != nil {
		stage.StageBranch(partgroup.GroupAbbreviation)
	}
	if partgroup.GroupSymbol != nil {
		stage.StageBranch(partgroup.GroupSymbol)
	}
	if partgroup.GroupBarline != nil {
		stage.StageBranch(partgroup.GroupBarline)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partlink *PartLink) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partlink) {
		return
	}

	partlink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrumentlink := range partlink.InstrumentLink {
		stage.StageBranch(_instrumentlink)
	}

}

func (partlist *PartList) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partlist) {
		return
	}

	partlist.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _partgroup := range partlist.PartGroup {
		stage.StageBranch(_partgroup)
	}
	for _, _scorepart := range partlist.ScorePart {
		stage.StageBranch(_scorepart)
	}

}

func (partname *PartName) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(partname) {
		return
	}

	partname.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (player *Player) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(player) {
		return
	}

	player.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scoreinstrument *ScoreInstrument) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scoreinstrument) {
		return
	}

	scoreinstrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scorepart *ScorePart) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scorepart) {
		return
	}

	scorepart.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scorepart.PartName != nil {
		stage.StageBranch(scorepart.PartName)
	}
	if scorepart.PartAbbreviation != nil {
		stage.StageBranch(scorepart.PartAbbreviation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _partlink := range scorepart.PartLink {
		stage.StageBranch(_partlink)
	}
	for _, _scoreinstrument := range scorepart.ScoreInstrument {
		stage.StageBranch(_scoreinstrument)
	}
	for _, _player := range scorepart.Player {
		stage.StageBranch(_player)
	}

}

func (scorepartwise *ScorePartwise) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scorepartwise) {
		return
	}

	scorepartwise.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scorepartwise.Work != nil {
		stage.StageBranch(scorepartwise.Work)
	}
	if scorepartwise.Defaults != nil {
		stage.StageBranch(scorepartwise.Defaults)
	}
	if scorepartwise.PartList != nil {
		stage.StageBranch(scorepartwise.PartList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range scorepartwise.Credit {
		stage.StageBranch(_credit)
	}
	for _, _part := range scorepartwise.Part {
		stage.StageBranch(_part)
	}

}

func (scoretimewise *ScoreTimewise) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scoretimewise) {
		return
	}

	scoretimewise.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scoretimewise.Work != nil {
		stage.StageBranch(scoretimewise.Work)
	}
	if scoretimewise.Defaults != nil {
		stage.StageBranch(scoretimewise.Defaults)
	}
	if scoretimewise.PartList != nil {
		stage.StageBranch(scoretimewise.PartList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range scoretimewise.Credit {
		stage.StageBranch(_credit)
	}
	for _, _scoretimewisemeasure := range scoretimewise.Measure {
		stage.StageBranch(_scoretimewisemeasure)
	}

}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scoretimewisemeasure) {
		return
	}

	scoretimewisemeasure.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _scoretimewisemeasurepart := range scoretimewisemeasure.Part {
		stage.StageBranch(_scoretimewisemeasurepart)
	}

}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scoretimewisemeasurepart) {
		return
	}

	scoretimewisemeasurepart.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scoretimewisemeasurepart.RefScorePart != nil {
		stage.StageBranch(scoretimewisemeasurepart.RefScorePart)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (work *Work) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(work) {
		return
	}

	work.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if work.Opus != nil {
		stage.StageBranch(work.Opus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Credit:
		toT := GongCopyBranchCredit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Defaults:
		toT := GongCopyBranchDefaults(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EmptyFont:
		toT := GongCopyBranchEmptyFont(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GroupBarline:
		toT := GongCopyBranchGroupBarline(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GroupName:
		toT := GongCopyBranchGroupName(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GroupSymbol:
		toT := GongCopyBranchGroupSymbol(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *InstrumentLink:
		toT := GongCopyBranchInstrumentLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LyricFont:
		toT := GongCopyBranchLyricFont(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LyricLanguage:
		toT := GongCopyBranchLyricLanguage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Measure:
		toT := GongCopyBranchMeasure(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeasureMusicDataChoice:
		toT := GongCopyBranchMeasureMusicDataChoice(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Opus:
		toT := GongCopyBranchOpus(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part:
		toT := GongCopyBranchPart(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartGroup:
		toT := GongCopyBranchPartGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartLink:
		toT := GongCopyBranchPartLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartList:
		toT := GongCopyBranchPartList(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartName:
		toT := GongCopyBranchPartName(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Player:
		toT := GongCopyBranchPlayer(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ScoreInstrument:
		toT := GongCopyBranchScoreInstrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ScorePart:
		toT := GongCopyBranchScorePart(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ScorePartwise:
		toT := GongCopyBranchScorePartwise(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ScoreTimewise:
		toT := GongCopyBranchScoreTimewise(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ScoreTimewiseMeasure:
		toT := GongCopyBranchScoreTimewiseMeasure(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ScoreTimewiseMeasurePart:
		toT := GongCopyBranchScoreTimewiseMeasurePart(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Work:
		toT := GongCopyBranchWork(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCredit(mapOrigCopy map[any]any, creditFrom *Credit) (creditTo *Credit) {
	var alreadyCopied bool
	creditTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, creditFrom)
	if alreadyCopied {
		return
	}
	creditFrom.GongCopyBasicFields(creditTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDefaults(mapOrigCopy map[any]any, defaultsFrom *Defaults) (defaultsTo *Defaults) {
	var alreadyCopied bool
	defaultsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, defaultsFrom)
	if alreadyCopied {
		return
	}
	defaultsFrom.GongCopyBasicFields(defaultsTo)

	//insertion point for the staging of instances referenced by pointers
	if defaultsFrom.MusicFont != nil {
		defaultsTo.MusicFont = GongCopyBranchEmptyFont(mapOrigCopy, defaultsFrom.MusicFont)
	}
	if defaultsFrom.WordFont != nil {
		defaultsTo.WordFont = GongCopyBranchEmptyFont(mapOrigCopy, defaultsFrom.WordFont)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lyricfont := range defaultsFrom.LyricFont {
		defaultsTo.LyricFont = append(defaultsTo.LyricFont, GongCopyBranchLyricFont(mapOrigCopy, _lyricfont))
	}
	for _, _lyriclanguage := range defaultsFrom.LyricLanguage {
		defaultsTo.LyricLanguage = append(defaultsTo.LyricLanguage, GongCopyBranchLyricLanguage(mapOrigCopy, _lyriclanguage))
	}

	return
}

func GongCopyBranchEmptyFont(mapOrigCopy map[any]any, emptyfontFrom *EmptyFont) (emptyfontTo *EmptyFont) {
	var alreadyCopied bool
	emptyfontTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptyfontFrom)
	if alreadyCopied {
		return
	}
	emptyfontFrom.GongCopyBasicFields(emptyfontTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGroupBarline(mapOrigCopy map[any]any, groupbarlineFrom *GroupBarline) (groupbarlineTo *GroupBarline) {
	var alreadyCopied bool
	groupbarlineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupbarlineFrom)
	if alreadyCopied {
		return
	}
	groupbarlineFrom.GongCopyBasicFields(groupbarlineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGroupName(mapOrigCopy map[any]any, groupnameFrom *GroupName) (groupnameTo *GroupName) {
	var alreadyCopied bool
	groupnameTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupnameFrom)
	if alreadyCopied {
		return
	}
	groupnameFrom.GongCopyBasicFields(groupnameTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGroupSymbol(mapOrigCopy map[any]any, groupsymbolFrom *GroupSymbol) (groupsymbolTo *GroupSymbol) {
	var alreadyCopied bool
	groupsymbolTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupsymbolFrom)
	if alreadyCopied {
		return
	}
	groupsymbolFrom.GongCopyBasicFields(groupsymbolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInstrumentLink(mapOrigCopy map[any]any, instrumentlinkFrom *InstrumentLink) (instrumentlinkTo *InstrumentLink) {
	var alreadyCopied bool
	instrumentlinkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, instrumentlinkFrom)
	if alreadyCopied {
		return
	}
	instrumentlinkFrom.GongCopyBasicFields(instrumentlinkTo)

	//insertion point for the staging of instances referenced by pointers
	if instrumentlinkFrom.RefScoreInstrument != nil {
		instrumentlinkTo.RefScoreInstrument = GongCopyBranchScoreInstrument(mapOrigCopy, instrumentlinkFrom.RefScoreInstrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLyricFont(mapOrigCopy map[any]any, lyricfontFrom *LyricFont) (lyricfontTo *LyricFont) {
	var alreadyCopied bool
	lyricfontTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lyricfontFrom)
	if alreadyCopied {
		return
	}
	lyricfontFrom.GongCopyBasicFields(lyricfontTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLyricLanguage(mapOrigCopy map[any]any, lyriclanguageFrom *LyricLanguage) (lyriclanguageTo *LyricLanguage) {
	var alreadyCopied bool
	lyriclanguageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lyriclanguageFrom)
	if alreadyCopied {
		return
	}
	lyriclanguageFrom.GongCopyBasicFields(lyriclanguageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasure(mapOrigCopy map[any]any, measureFrom *Measure) (measureTo *Measure) {
	var alreadyCopied bool
	measureTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measureFrom)
	if alreadyCopied {
		return
	}
	measureFrom.GongCopyBasicFields(measureTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _measuremusicdatachoice := range measureFrom.MusicData {
		measureTo.MusicData = append(measureTo.MusicData, GongCopyBranchMeasureMusicDataChoice(mapOrigCopy, _measuremusicdatachoice))
	}

	return
}

func GongCopyBranchMeasureMusicDataChoice(mapOrigCopy map[any]any, measuremusicdatachoiceFrom *MeasureMusicDataChoice) (measuremusicdatachoiceTo *MeasureMusicDataChoice) {
	var alreadyCopied bool
	measuremusicdatachoiceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measuremusicdatachoiceFrom)
	if alreadyCopied {
		return
	}
	measuremusicdatachoiceFrom.GongCopyBasicFields(measuremusicdatachoiceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOpus(mapOrigCopy map[any]any, opusFrom *Opus) (opusTo *Opus) {
	var alreadyCopied bool
	opusTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, opusFrom)
	if alreadyCopied {
		return
	}
	opusFrom.GongCopyBasicFields(opusTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPart(mapOrigCopy map[any]any, partFrom *Part) (partTo *Part) {
	var alreadyCopied bool
	partTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partFrom)
	if alreadyCopied {
		return
	}
	partFrom.GongCopyBasicFields(partTo)

	//insertion point for the staging of instances referenced by pointers
	if partFrom.RefScorePart != nil {
		partTo.RefScorePart = GongCopyBranchScorePart(mapOrigCopy, partFrom.RefScorePart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _measure := range partFrom.Measure {
		partTo.Measure = append(partTo.Measure, GongCopyBranchMeasure(mapOrigCopy, _measure))
	}

	return
}

func GongCopyBranchPartGroup(mapOrigCopy map[any]any, partgroupFrom *PartGroup) (partgroupTo *PartGroup) {
	var alreadyCopied bool
	partgroupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partgroupFrom)
	if alreadyCopied {
		return
	}
	partgroupFrom.GongCopyBasicFields(partgroupTo)

	//insertion point for the staging of instances referenced by pointers
	if partgroupFrom.GroupName != nil {
		partgroupTo.GroupName = GongCopyBranchGroupName(mapOrigCopy, partgroupFrom.GroupName)
	}
	if partgroupFrom.GroupAbbreviation != nil {
		partgroupTo.GroupAbbreviation = GongCopyBranchGroupName(mapOrigCopy, partgroupFrom.GroupAbbreviation)
	}
	if partgroupFrom.GroupSymbol != nil {
		partgroupTo.GroupSymbol = GongCopyBranchGroupSymbol(mapOrigCopy, partgroupFrom.GroupSymbol)
	}
	if partgroupFrom.GroupBarline != nil {
		partgroupTo.GroupBarline = GongCopyBranchGroupBarline(mapOrigCopy, partgroupFrom.GroupBarline)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartLink(mapOrigCopy map[any]any, partlinkFrom *PartLink) (partlinkTo *PartLink) {
	var alreadyCopied bool
	partlinkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partlinkFrom)
	if alreadyCopied {
		return
	}
	partlinkFrom.GongCopyBasicFields(partlinkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrumentlink := range partlinkFrom.InstrumentLink {
		partlinkTo.InstrumentLink = append(partlinkTo.InstrumentLink, GongCopyBranchInstrumentLink(mapOrigCopy, _instrumentlink))
	}

	return
}

func GongCopyBranchPartList(mapOrigCopy map[any]any, partlistFrom *PartList) (partlistTo *PartList) {
	var alreadyCopied bool
	partlistTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partlistFrom)
	if alreadyCopied {
		return
	}
	partlistFrom.GongCopyBasicFields(partlistTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _partgroup := range partlistFrom.PartGroup {
		partlistTo.PartGroup = append(partlistTo.PartGroup, GongCopyBranchPartGroup(mapOrigCopy, _partgroup))
	}
	for _, _scorepart := range partlistFrom.ScorePart {
		partlistTo.ScorePart = append(partlistTo.ScorePart, GongCopyBranchScorePart(mapOrigCopy, _scorepart))
	}

	return
}

func GongCopyBranchPartName(mapOrigCopy map[any]any, partnameFrom *PartName) (partnameTo *PartName) {
	var alreadyCopied bool
	partnameTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partnameFrom)
	if alreadyCopied {
		return
	}
	partnameFrom.GongCopyBasicFields(partnameTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlayer(mapOrigCopy map[any]any, playerFrom *Player) (playerTo *Player) {
	var alreadyCopied bool
	playerTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, playerFrom)
	if alreadyCopied {
		return
	}
	playerFrom.GongCopyBasicFields(playerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchScoreInstrument(mapOrigCopy map[any]any, scoreinstrumentFrom *ScoreInstrument) (scoreinstrumentTo *ScoreInstrument) {
	var alreadyCopied bool
	scoreinstrumentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scoreinstrumentFrom)
	if alreadyCopied {
		return
	}
	scoreinstrumentFrom.GongCopyBasicFields(scoreinstrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchScorePart(mapOrigCopy map[any]any, scorepartFrom *ScorePart) (scorepartTo *ScorePart) {
	var alreadyCopied bool
	scorepartTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scorepartFrom)
	if alreadyCopied {
		return
	}
	scorepartFrom.GongCopyBasicFields(scorepartTo)

	//insertion point for the staging of instances referenced by pointers
	if scorepartFrom.PartName != nil {
		scorepartTo.PartName = GongCopyBranchPartName(mapOrigCopy, scorepartFrom.PartName)
	}
	if scorepartFrom.PartAbbreviation != nil {
		scorepartTo.PartAbbreviation = GongCopyBranchPartName(mapOrigCopy, scorepartFrom.PartAbbreviation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _partlink := range scorepartFrom.PartLink {
		scorepartTo.PartLink = append(scorepartTo.PartLink, GongCopyBranchPartLink(mapOrigCopy, _partlink))
	}
	for _, _scoreinstrument := range scorepartFrom.ScoreInstrument {
		scorepartTo.ScoreInstrument = append(scorepartTo.ScoreInstrument, GongCopyBranchScoreInstrument(mapOrigCopy, _scoreinstrument))
	}
	for _, _player := range scorepartFrom.Player {
		scorepartTo.Player = append(scorepartTo.Player, GongCopyBranchPlayer(mapOrigCopy, _player))
	}

	return
}

func GongCopyBranchScorePartwise(mapOrigCopy map[any]any, scorepartwiseFrom *ScorePartwise) (scorepartwiseTo *ScorePartwise) {
	var alreadyCopied bool
	scorepartwiseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scorepartwiseFrom)
	if alreadyCopied {
		return
	}
	scorepartwiseFrom.GongCopyBasicFields(scorepartwiseTo)

	//insertion point for the staging of instances referenced by pointers
	if scorepartwiseFrom.Work != nil {
		scorepartwiseTo.Work = GongCopyBranchWork(mapOrigCopy, scorepartwiseFrom.Work)
	}
	if scorepartwiseFrom.Defaults != nil {
		scorepartwiseTo.Defaults = GongCopyBranchDefaults(mapOrigCopy, scorepartwiseFrom.Defaults)
	}
	if scorepartwiseFrom.PartList != nil {
		scorepartwiseTo.PartList = GongCopyBranchPartList(mapOrigCopy, scorepartwiseFrom.PartList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range scorepartwiseFrom.Credit {
		scorepartwiseTo.Credit = append(scorepartwiseTo.Credit, GongCopyBranchCredit(mapOrigCopy, _credit))
	}
	for _, _part := range scorepartwiseFrom.Part {
		scorepartwiseTo.Part = append(scorepartwiseTo.Part, GongCopyBranchPart(mapOrigCopy, _part))
	}

	return
}

func GongCopyBranchScoreTimewise(mapOrigCopy map[any]any, scoretimewiseFrom *ScoreTimewise) (scoretimewiseTo *ScoreTimewise) {
	var alreadyCopied bool
	scoretimewiseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scoretimewiseFrom)
	if alreadyCopied {
		return
	}
	scoretimewiseFrom.GongCopyBasicFields(scoretimewiseTo)

	//insertion point for the staging of instances referenced by pointers
	if scoretimewiseFrom.Work != nil {
		scoretimewiseTo.Work = GongCopyBranchWork(mapOrigCopy, scoretimewiseFrom.Work)
	}
	if scoretimewiseFrom.Defaults != nil {
		scoretimewiseTo.Defaults = GongCopyBranchDefaults(mapOrigCopy, scoretimewiseFrom.Defaults)
	}
	if scoretimewiseFrom.PartList != nil {
		scoretimewiseTo.PartList = GongCopyBranchPartList(mapOrigCopy, scoretimewiseFrom.PartList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range scoretimewiseFrom.Credit {
		scoretimewiseTo.Credit = append(scoretimewiseTo.Credit, GongCopyBranchCredit(mapOrigCopy, _credit))
	}
	for _, _scoretimewisemeasure := range scoretimewiseFrom.Measure {
		scoretimewiseTo.Measure = append(scoretimewiseTo.Measure, GongCopyBranchScoreTimewiseMeasure(mapOrigCopy, _scoretimewisemeasure))
	}

	return
}

func GongCopyBranchScoreTimewiseMeasure(mapOrigCopy map[any]any, scoretimewisemeasureFrom *ScoreTimewiseMeasure) (scoretimewisemeasureTo *ScoreTimewiseMeasure) {
	var alreadyCopied bool
	scoretimewisemeasureTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scoretimewisemeasureFrom)
	if alreadyCopied {
		return
	}
	scoretimewisemeasureFrom.GongCopyBasicFields(scoretimewisemeasureTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _scoretimewisemeasurepart := range scoretimewisemeasureFrom.Part {
		scoretimewisemeasureTo.Part = append(scoretimewisemeasureTo.Part, GongCopyBranchScoreTimewiseMeasurePart(mapOrigCopy, _scoretimewisemeasurepart))
	}

	return
}

func GongCopyBranchScoreTimewiseMeasurePart(mapOrigCopy map[any]any, scoretimewisemeasurepartFrom *ScoreTimewiseMeasurePart) (scoretimewisemeasurepartTo *ScoreTimewiseMeasurePart) {
	var alreadyCopied bool
	scoretimewisemeasurepartTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scoretimewisemeasurepartFrom)
	if alreadyCopied {
		return
	}
	scoretimewisemeasurepartFrom.GongCopyBasicFields(scoretimewisemeasurepartTo)

	//insertion point for the staging of instances referenced by pointers
	if scoretimewisemeasurepartFrom.RefScorePart != nil {
		scoretimewisemeasurepartTo.RefScorePart = GongCopyBranchScorePart(mapOrigCopy, scoretimewisemeasurepartFrom.RefScorePart)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWork(mapOrigCopy map[any]any, workFrom *Work) (workTo *Work) {
	var alreadyCopied bool
	workTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, workFrom)
	if alreadyCopied {
		return
	}
	workFrom.GongCopyBasicFields(workTo)

	//insertion point for the staging of instances referenced by pointers
	if workFrom.Opus != nil {
		workTo.Opus = GongCopyBranchOpus(mapOrigCopy, workFrom.Opus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (credit *Credit) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(credit) {
		return
	}

	credit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (defaults *Defaults) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(defaults) {
		return
	}

	defaults.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if defaults.MusicFont != nil {
		stage.UnstageBranch(defaults.MusicFont)
	}
	if defaults.WordFont != nil {
		stage.UnstageBranch(defaults.WordFont)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lyricfont := range defaults.LyricFont {
		stage.UnstageBranch(_lyricfont)
	}
	for _, _lyriclanguage := range defaults.LyricLanguage {
		stage.UnstageBranch(_lyriclanguage)
	}

}

func (emptyfont *EmptyFont) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(emptyfont) {
		return
	}

	emptyfont.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (groupbarline *GroupBarline) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(groupbarline) {
		return
	}

	groupbarline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (groupname *GroupName) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(groupname) {
		return
	}

	groupname.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (groupsymbol *GroupSymbol) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(groupsymbol) {
		return
	}

	groupsymbol.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrumentlink *InstrumentLink) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(instrumentlink) {
		return
	}

	instrumentlink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if instrumentlink.RefScoreInstrument != nil {
		stage.UnstageBranch(instrumentlink.RefScoreInstrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lyricfont *LyricFont) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lyricfont) {
		return
	}

	lyricfont.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lyriclanguage *LyricLanguage) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lyriclanguage) {
		return
	}

	lyriclanguage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure *Measure) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measure) {
		return
	}

	measure.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _measuremusicdatachoice := range measure.MusicData {
		stage.UnstageBranch(_measuremusicdatachoice)
	}

}

func (measuremusicdatachoice *MeasureMusicDataChoice) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measuremusicdatachoice) {
		return
	}

	measuremusicdatachoice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (opus *Opus) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(opus) {
		return
	}

	opus.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part *Part) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(part) {
		return
	}

	part.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part.RefScorePart != nil {
		stage.UnstageBranch(part.RefScorePart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _measure := range part.Measure {
		stage.UnstageBranch(_measure)
	}

}

func (partgroup *PartGroup) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partgroup) {
		return
	}

	partgroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if partgroup.GroupName != nil {
		stage.UnstageBranch(partgroup.GroupName)
	}
	if partgroup.GroupAbbreviation != nil {
		stage.UnstageBranch(partgroup.GroupAbbreviation)
	}
	if partgroup.GroupSymbol != nil {
		stage.UnstageBranch(partgroup.GroupSymbol)
	}
	if partgroup.GroupBarline != nil {
		stage.UnstageBranch(partgroup.GroupBarline)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partlink *PartLink) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partlink) {
		return
	}

	partlink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrumentlink := range partlink.InstrumentLink {
		stage.UnstageBranch(_instrumentlink)
	}

}

func (partlist *PartList) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partlist) {
		return
	}

	partlist.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _partgroup := range partlist.PartGroup {
		stage.UnstageBranch(_partgroup)
	}
	for _, _scorepart := range partlist.ScorePart {
		stage.UnstageBranch(_scorepart)
	}

}

func (partname *PartName) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(partname) {
		return
	}

	partname.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (player *Player) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(player) {
		return
	}

	player.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scoreinstrument *ScoreInstrument) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scoreinstrument) {
		return
	}

	scoreinstrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scorepart *ScorePart) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scorepart) {
		return
	}

	scorepart.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scorepart.PartName != nil {
		stage.UnstageBranch(scorepart.PartName)
	}
	if scorepart.PartAbbreviation != nil {
		stage.UnstageBranch(scorepart.PartAbbreviation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _partlink := range scorepart.PartLink {
		stage.UnstageBranch(_partlink)
	}
	for _, _scoreinstrument := range scorepart.ScoreInstrument {
		stage.UnstageBranch(_scoreinstrument)
	}
	for _, _player := range scorepart.Player {
		stage.UnstageBranch(_player)
	}

}

func (scorepartwise *ScorePartwise) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scorepartwise) {
		return
	}

	scorepartwise.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scorepartwise.Work != nil {
		stage.UnstageBranch(scorepartwise.Work)
	}
	if scorepartwise.Defaults != nil {
		stage.UnstageBranch(scorepartwise.Defaults)
	}
	if scorepartwise.PartList != nil {
		stage.UnstageBranch(scorepartwise.PartList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range scorepartwise.Credit {
		stage.UnstageBranch(_credit)
	}
	for _, _part := range scorepartwise.Part {
		stage.UnstageBranch(_part)
	}

}

func (scoretimewise *ScoreTimewise) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scoretimewise) {
		return
	}

	scoretimewise.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scoretimewise.Work != nil {
		stage.UnstageBranch(scoretimewise.Work)
	}
	if scoretimewise.Defaults != nil {
		stage.UnstageBranch(scoretimewise.Defaults)
	}
	if scoretimewise.PartList != nil {
		stage.UnstageBranch(scoretimewise.PartList)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range scoretimewise.Credit {
		stage.UnstageBranch(_credit)
	}
	for _, _scoretimewisemeasure := range scoretimewise.Measure {
		stage.UnstageBranch(_scoretimewisemeasure)
	}

}

func (scoretimewisemeasure *ScoreTimewiseMeasure) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scoretimewisemeasure) {
		return
	}

	scoretimewisemeasure.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _scoretimewisemeasurepart := range scoretimewisemeasure.Part {
		stage.UnstageBranch(_scoretimewisemeasurepart)
	}

}

func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scoretimewisemeasurepart) {
		return
	}

	scoretimewisemeasurepart.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if scoretimewisemeasurepart.RefScorePart != nil {
		stage.UnstageBranch(scoretimewisemeasurepart.RefScorePart)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (work *Work) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(work) {
		return
	}

	work.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if work.Opus != nil {
		stage.UnstageBranch(work.Opus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Credit) GongReconstructPointersFromReferences(stage *Stage, instance *Credit) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Defaults) GongReconstructPointersFromReferences(stage *Stage, instance *Defaults) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MusicFont, stage.EmptyFonts_reference, instance.MusicFont)
	__gong__reconstructPointer(&reference.WordFont, stage.EmptyFonts_reference, instance.WordFont)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.LyricFont, stage.LyricFonts_reference, instance.LyricFont)
	__gong__reconstructSliceOfPointersFromReferences(&reference.LyricLanguage, stage.LyricLanguages_reference, instance.LyricLanguage)
}

func (reference *EmptyFont) GongReconstructPointersFromReferences(stage *Stage, instance *EmptyFont) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GroupBarline) GongReconstructPointersFromReferences(stage *Stage, instance *GroupBarline) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GroupName) GongReconstructPointersFromReferences(stage *Stage, instance *GroupName) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GroupSymbol) GongReconstructPointersFromReferences(stage *Stage, instance *GroupSymbol) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *InstrumentLink) GongReconstructPointersFromReferences(stage *Stage, instance *InstrumentLink) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.RefScoreInstrument, stage.ScoreInstruments_reference, instance.RefScoreInstrument)
	// insertion point for slice of pointers field
}

func (reference *LyricFont) GongReconstructPointersFromReferences(stage *Stage, instance *LyricFont) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LyricLanguage) GongReconstructPointersFromReferences(stage *Stage, instance *LyricLanguage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Measure) GongReconstructPointersFromReferences(stage *Stage, instance *Measure) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.MusicData, stage.MeasureMusicDataChoices_reference, instance.MusicData)
}

func (reference *MeasureMusicDataChoice) GongReconstructPointersFromReferences(stage *Stage, instance *MeasureMusicDataChoice) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Opus) GongReconstructPointersFromReferences(stage *Stage, instance *Opus) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Part) GongReconstructPointersFromReferences(stage *Stage, instance *Part) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.RefScorePart, stage.ScoreParts_reference, instance.RefScorePart)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Measure, stage.Measures_reference, instance.Measure)
}

func (reference *PartGroup) GongReconstructPointersFromReferences(stage *Stage, instance *PartGroup) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.GroupName, stage.GroupNames_reference, instance.GroupName)
	__gong__reconstructPointer(&reference.GroupAbbreviation, stage.GroupNames_reference, instance.GroupAbbreviation)
	__gong__reconstructPointer(&reference.GroupSymbol, stage.GroupSymbols_reference, instance.GroupSymbol)
	__gong__reconstructPointer(&reference.GroupBarline, stage.GroupBarlines_reference, instance.GroupBarline)
	// insertion point for slice of pointers field
}

func (reference *PartLink) GongReconstructPointersFromReferences(stage *Stage, instance *PartLink) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.InstrumentLink, stage.InstrumentLinks_reference, instance.InstrumentLink)
}

func (reference *PartList) GongReconstructPointersFromReferences(stage *Stage, instance *PartList) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.PartGroup, stage.PartGroups_reference, instance.PartGroup)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ScorePart, stage.ScoreParts_reference, instance.ScorePart)
}

func (reference *PartName) GongReconstructPointersFromReferences(stage *Stage, instance *PartName) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Player) GongReconstructPointersFromReferences(stage *Stage, instance *Player) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ScoreInstrument) GongReconstructPointersFromReferences(stage *Stage, instance *ScoreInstrument) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ScorePart) GongReconstructPointersFromReferences(stage *Stage, instance *ScorePart) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.PartName, stage.PartNames_reference, instance.PartName)
	__gong__reconstructPointer(&reference.PartAbbreviation, stage.PartNames_reference, instance.PartAbbreviation)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.PartLink, stage.PartLinks_reference, instance.PartLink)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ScoreInstrument, stage.ScoreInstruments_reference, instance.ScoreInstrument)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Player, stage.Players_reference, instance.Player)
}

func (reference *ScorePartwise) GongReconstructPointersFromReferences(stage *Stage, instance *ScorePartwise) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Work, stage.Works_reference, instance.Work)
	__gong__reconstructPointer(&reference.Defaults, stage.Defaultss_reference, instance.Defaults)
	__gong__reconstructPointer(&reference.PartList, stage.PartLists_reference, instance.PartList)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Credit, stage.Credits_reference, instance.Credit)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Part, stage.Parts_reference, instance.Part)
}

func (reference *ScoreTimewise) GongReconstructPointersFromReferences(stage *Stage, instance *ScoreTimewise) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Work, stage.Works_reference, instance.Work)
	__gong__reconstructPointer(&reference.Defaults, stage.Defaultss_reference, instance.Defaults)
	__gong__reconstructPointer(&reference.PartList, stage.PartLists_reference, instance.PartList)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Credit, stage.Credits_reference, instance.Credit)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Measure, stage.ScoreTimewiseMeasures_reference, instance.Measure)
}

func (reference *ScoreTimewiseMeasure) GongReconstructPointersFromReferences(stage *Stage, instance *ScoreTimewiseMeasure) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Part, stage.ScoreTimewiseMeasureParts_reference, instance.Part)
}

func (reference *ScoreTimewiseMeasurePart) GongReconstructPointersFromReferences(stage *Stage, instance *ScoreTimewiseMeasurePart) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.RefScorePart, stage.ScoreParts_reference, instance.RefScorePart)
	// insertion point for slice of pointers field
}

func (reference *Work) GongReconstructPointersFromReferences(stage *Stage, instance *Work) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Opus, stage.Opuss_reference, instance.Opus)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Credit) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Defaults) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MusicFont, stage.EmptyFonts_instance)
	__gong__reconstructPointerFromInstance(&reference.WordFont, stage.EmptyFonts_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.LyricFont, stage.LyricFonts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.LyricLanguage, stage.LyricLanguages_instance)
}

func (reference *EmptyFont) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GroupBarline) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GroupName) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GroupSymbol) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *InstrumentLink) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.RefScoreInstrument, stage.ScoreInstruments_instance)
	// insertion point for slice of pointers fields
}

func (reference *LyricFont) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LyricLanguage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Measure) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.MusicData, stage.MeasureMusicDataChoices_instance)
}

func (reference *MeasureMusicDataChoice) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Opus) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Part) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.RefScorePart, stage.ScoreParts_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Measure, stage.Measures_instance)
}

func (reference *PartGroup) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.GroupName, stage.GroupNames_instance)
	__gong__reconstructPointerFromInstance(&reference.GroupAbbreviation, stage.GroupNames_instance)
	__gong__reconstructPointerFromInstance(&reference.GroupSymbol, stage.GroupSymbols_instance)
	__gong__reconstructPointerFromInstance(&reference.GroupBarline, stage.GroupBarlines_instance)
	// insertion point for slice of pointers fields
}

func (reference *PartLink) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.InstrumentLink, stage.InstrumentLinks_instance)
}

func (reference *PartList) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.PartGroup, stage.PartGroups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ScorePart, stage.ScoreParts_instance)
}

func (reference *PartName) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Player) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ScoreInstrument) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ScorePart) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.PartName, stage.PartNames_instance)
	__gong__reconstructPointerFromInstance(&reference.PartAbbreviation, stage.PartNames_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.PartLink, stage.PartLinks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ScoreInstrument, stage.ScoreInstruments_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Player, stage.Players_instance)
}

func (reference *ScorePartwise) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Work, stage.Works_instance)
	__gong__reconstructPointerFromInstance(&reference.Defaults, stage.Defaultss_instance)
	__gong__reconstructPointerFromInstance(&reference.PartList, stage.PartLists_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Credit, stage.Credits_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Part, stage.Parts_instance)
}

func (reference *ScoreTimewise) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Work, stage.Works_instance)
	__gong__reconstructPointerFromInstance(&reference.Defaults, stage.Defaultss_instance)
	__gong__reconstructPointerFromInstance(&reference.PartList, stage.PartLists_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Credit, stage.Credits_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Measure, stage.ScoreTimewiseMeasures_instance)
}

func (reference *ScoreTimewiseMeasure) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Part, stage.ScoreTimewiseMeasureParts_instance)
}

func (reference *ScoreTimewiseMeasurePart) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.RefScorePart, stage.ScoreParts_instance)
	// insertion point for slice of pointers fields
}

func (reference *Work) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Opus, stage.Opuss_instance)
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (credit *Credit) GongDiff(stage *Stage, creditOther *Credit) (diffs []string) {
	// insertion point for field diffs
	if credit.Name != creditOther.Name {
		diffs = append(diffs, credit.GongMarshallField(stage, "Name"))
	}
	if credit.Page != creditOther.Page {
		diffs = append(diffs, credit.GongMarshallField(stage, "Page"))
	}
	if credit.Id != creditOther.Id {
		diffs = append(diffs, credit.GongMarshallField(stage, "Id"))
	}
	if credit.CreditType != creditOther.CreditType {
		diffs = append(diffs, credit.GongMarshallField(stage, "CreditType"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (defaults *Defaults) GongDiff(stage *Stage, defaultsOther *Defaults) (diffs []string) {
	// insertion point for field diffs
	if defaults.Name != defaultsOther.Name {
		diffs = append(diffs, defaults.GongMarshallField(stage, "Name"))
	}
	if defaults.MusicFont != defaultsOther.MusicFont {
		diffs = append(diffs, defaults.GongMarshallField(stage, "MusicFont"))
	}
	if defaults.WordFont != defaultsOther.WordFont {
		diffs = append(diffs, defaults.GongMarshallField(stage, "WordFont"))
	}
	if ops := __gong__diffSliceOfPointers(stage, defaults, "LyricFont", defaultsOther.LyricFont, defaults.LyricFont); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, defaults, "LyricLanguage", defaultsOther.LyricLanguage, defaults.LyricLanguage); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (emptyfont *EmptyFont) GongDiff(stage *Stage, emptyfontOther *EmptyFont) (diffs []string) {
	// insertion point for field diffs
	if emptyfont.Name != emptyfontOther.Name {
		diffs = append(diffs, emptyfont.GongMarshallField(stage, "Name"))
	}
	if emptyfont.FontFamily != emptyfontOther.FontFamily {
		diffs = append(diffs, emptyfont.GongMarshallField(stage, "FontFamily"))
	}
	if emptyfont.FontSize != emptyfontOther.FontSize {
		diffs = append(diffs, emptyfont.GongMarshallField(stage, "FontSize"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (groupbarline *GroupBarline) GongDiff(stage *Stage, groupbarlineOther *GroupBarline) (diffs []string) {
	// insertion point for field diffs
	if groupbarline.Name != groupbarlineOther.Name {
		diffs = append(diffs, groupbarline.GongMarshallField(stage, "Name"))
	}
	if groupbarline.Color != groupbarlineOther.Color {
		diffs = append(diffs, groupbarline.GongMarshallField(stage, "Color"))
	}
	if groupbarline.Value != groupbarlineOther.Value {
		diffs = append(diffs, groupbarline.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (groupname *GroupName) GongDiff(stage *Stage, groupnameOther *GroupName) (diffs []string) {
	// insertion point for field diffs
	if groupname.Name != groupnameOther.Name {
		diffs = append(diffs, groupname.GongMarshallField(stage, "Name"))
	}
	if groupname.DefaultX != groupnameOther.DefaultX {
		diffs = append(diffs, groupname.GongMarshallField(stage, "DefaultX"))
	}
	if groupname.DefaultY != groupnameOther.DefaultY {
		diffs = append(diffs, groupname.GongMarshallField(stage, "DefaultY"))
	}
	if groupname.RelativeX != groupnameOther.RelativeX {
		diffs = append(diffs, groupname.GongMarshallField(stage, "RelativeX"))
	}
	if groupname.RelativeY != groupnameOther.RelativeY {
		diffs = append(diffs, groupname.GongMarshallField(stage, "RelativeY"))
	}
	if groupname.FontFamily != groupnameOther.FontFamily {
		diffs = append(diffs, groupname.GongMarshallField(stage, "FontFamily"))
	}
	if groupname.FontSize != groupnameOther.FontSize {
		diffs = append(diffs, groupname.GongMarshallField(stage, "FontSize"))
	}
	if groupname.Color != groupnameOther.Color {
		diffs = append(diffs, groupname.GongMarshallField(stage, "Color"))
	}
	if groupname.Value != groupnameOther.Value {
		diffs = append(diffs, groupname.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (groupsymbol *GroupSymbol) GongDiff(stage *Stage, groupsymbolOther *GroupSymbol) (diffs []string) {
	// insertion point for field diffs
	if groupsymbol.Name != groupsymbolOther.Name {
		diffs = append(diffs, groupsymbol.GongMarshallField(stage, "Name"))
	}
	if groupsymbol.DefaultX != groupsymbolOther.DefaultX {
		diffs = append(diffs, groupsymbol.GongMarshallField(stage, "DefaultX"))
	}
	if groupsymbol.DefaultY != groupsymbolOther.DefaultY {
		diffs = append(diffs, groupsymbol.GongMarshallField(stage, "DefaultY"))
	}
	if groupsymbol.RelativeX != groupsymbolOther.RelativeX {
		diffs = append(diffs, groupsymbol.GongMarshallField(stage, "RelativeX"))
	}
	if groupsymbol.RelativeY != groupsymbolOther.RelativeY {
		diffs = append(diffs, groupsymbol.GongMarshallField(stage, "RelativeY"))
	}
	if groupsymbol.Color != groupsymbolOther.Color {
		diffs = append(diffs, groupsymbol.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (instrumentlink *InstrumentLink) GongDiff(stage *Stage, instrumentlinkOther *InstrumentLink) (diffs []string) {
	// insertion point for field diffs
	if instrumentlink.Name != instrumentlinkOther.Name {
		diffs = append(diffs, instrumentlink.GongMarshallField(stage, "Name"))
	}
	if instrumentlink.RefScoreInstrument != instrumentlinkOther.RefScoreInstrument {
		diffs = append(diffs, instrumentlink.GongMarshallField(stage, "RefScoreInstrument"))
	}
	if instrumentlink.Id != instrumentlinkOther.Id {
		diffs = append(diffs, instrumentlink.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lyricfont *LyricFont) GongDiff(stage *Stage, lyricfontOther *LyricFont) (diffs []string) {
	// insertion point for field diffs
	if lyricfont.Name != lyricfontOther.Name {
		diffs = append(diffs, lyricfont.GongMarshallField(stage, "Name"))
	}
	if lyricfont.Number != lyricfontOther.Number {
		diffs = append(diffs, lyricfont.GongMarshallField(stage, "Number"))
	}
	if lyricfont.NameAttr != lyricfontOther.NameAttr {
		diffs = append(diffs, lyricfont.GongMarshallField(stage, "NameAttr"))
	}
	if lyricfont.FontFamily != lyricfontOther.FontFamily {
		diffs = append(diffs, lyricfont.GongMarshallField(stage, "FontFamily"))
	}
	if lyricfont.FontSize != lyricfontOther.FontSize {
		diffs = append(diffs, lyricfont.GongMarshallField(stage, "FontSize"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lyriclanguage *LyricLanguage) GongDiff(stage *Stage, lyriclanguageOther *LyricLanguage) (diffs []string) {
	// insertion point for field diffs
	if lyriclanguage.Name != lyriclanguageOther.Name {
		diffs = append(diffs, lyriclanguage.GongMarshallField(stage, "Name"))
	}
	if lyriclanguage.Number != lyriclanguageOther.Number {
		diffs = append(diffs, lyriclanguage.GongMarshallField(stage, "Number"))
	}
	if lyriclanguage.NameAttr != lyriclanguageOther.NameAttr {
		diffs = append(diffs, lyriclanguage.GongMarshallField(stage, "NameAttr"))
	}
	if lyriclanguage.Lang != lyriclanguageOther.Lang {
		diffs = append(diffs, lyriclanguage.GongMarshallField(stage, "Lang"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measure *Measure) GongDiff(stage *Stage, measureOther *Measure) (diffs []string) {
	// insertion point for field diffs
	if measure.Name != measureOther.Name {
		diffs = append(diffs, measure.GongMarshallField(stage, "Name"))
	}
	if measure.Number != measureOther.Number {
		diffs = append(diffs, measure.GongMarshallField(stage, "Number"))
	}
	if measure.Text != measureOther.Text {
		diffs = append(diffs, measure.GongMarshallField(stage, "Text"))
	}
	if measure.Width != measureOther.Width {
		diffs = append(diffs, measure.GongMarshallField(stage, "Width"))
	}
	if measure.Id != measureOther.Id {
		diffs = append(diffs, measure.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, measure, "MusicData", measureOther.MusicData, measure.MusicData); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measuremusicdatachoice *MeasureMusicDataChoice) GongDiff(stage *Stage, measuremusicdatachoiceOther *MeasureMusicDataChoice) (diffs []string) {
	// insertion point for field diffs
	if measuremusicdatachoice.Name != measuremusicdatachoiceOther.Name {
		diffs = append(diffs, measuremusicdatachoice.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (opus *Opus) GongDiff(stage *Stage, opusOther *Opus) (diffs []string) {
	// insertion point for field diffs
	if opus.Name != opusOther.Name {
		diffs = append(diffs, opus.GongMarshallField(stage, "Name"))
	}
	if opus.Href != opusOther.Href {
		diffs = append(diffs, opus.GongMarshallField(stage, "Href"))
	}
	if opus.Role != opusOther.Role {
		diffs = append(diffs, opus.GongMarshallField(stage, "Role"))
	}
	if opus.Title != opusOther.Title {
		diffs = append(diffs, opus.GongMarshallField(stage, "Title"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (part *Part) GongDiff(stage *Stage, partOther *Part) (diffs []string) {
	// insertion point for field diffs
	if part.Name != partOther.Name {
		diffs = append(diffs, part.GongMarshallField(stage, "Name"))
	}
	if part.RefScorePart != partOther.RefScorePart {
		diffs = append(diffs, part.GongMarshallField(stage, "RefScorePart"))
	}
	if part.Id != partOther.Id {
		diffs = append(diffs, part.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, part, "Measure", partOther.Measure, part.Measure); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partgroup *PartGroup) GongDiff(stage *Stage, partgroupOther *PartGroup) (diffs []string) {
	// insertion point for field diffs
	if partgroup.Name != partgroupOther.Name {
		diffs = append(diffs, partgroup.GongMarshallField(stage, "Name"))
	}
	if partgroup.Number != partgroupOther.Number {
		diffs = append(diffs, partgroup.GongMarshallField(stage, "Number"))
	}
	if partgroup.GroupName != partgroupOther.GroupName {
		diffs = append(diffs, partgroup.GongMarshallField(stage, "GroupName"))
	}
	if partgroup.GroupAbbreviation != partgroupOther.GroupAbbreviation {
		diffs = append(diffs, partgroup.GongMarshallField(stage, "GroupAbbreviation"))
	}
	if partgroup.GroupSymbol != partgroupOther.GroupSymbol {
		diffs = append(diffs, partgroup.GongMarshallField(stage, "GroupSymbol"))
	}
	if partgroup.GroupBarline != partgroupOther.GroupBarline {
		diffs = append(diffs, partgroup.GongMarshallField(stage, "GroupBarline"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partlink *PartLink) GongDiff(stage *Stage, partlinkOther *PartLink) (diffs []string) {
	// insertion point for field diffs
	if partlink.Name != partlinkOther.Name {
		diffs = append(diffs, partlink.GongMarshallField(stage, "Name"))
	}
	if partlink.Href != partlinkOther.Href {
		diffs = append(diffs, partlink.GongMarshallField(stage, "Href"))
	}
	if partlink.Role != partlinkOther.Role {
		diffs = append(diffs, partlink.GongMarshallField(stage, "Role"))
	}
	if partlink.Title != partlinkOther.Title {
		diffs = append(diffs, partlink.GongMarshallField(stage, "Title"))
	}
	if ops := __gong__diffSliceOfPointers(stage, partlink, "InstrumentLink", partlinkOther.InstrumentLink, partlink.InstrumentLink); ops != "" {
		diffs = append(diffs, ops)
	}
	if partlink.GroupLink != partlinkOther.GroupLink {
		diffs = append(diffs, partlink.GongMarshallField(stage, "GroupLink"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partlist *PartList) GongDiff(stage *Stage, partlistOther *PartList) (diffs []string) {
	// insertion point for field diffs
	if partlist.Name != partlistOther.Name {
		diffs = append(diffs, partlist.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, partlist, "PartGroup", partlistOther.PartGroup, partlist.PartGroup); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, partlist, "ScorePart", partlistOther.ScorePart, partlist.ScorePart); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partname *PartName) GongDiff(stage *Stage, partnameOther *PartName) (diffs []string) {
	// insertion point for field diffs
	if partname.Name != partnameOther.Name {
		diffs = append(diffs, partname.GongMarshallField(stage, "Name"))
	}
	if partname.DefaultX != partnameOther.DefaultX {
		diffs = append(diffs, partname.GongMarshallField(stage, "DefaultX"))
	}
	if partname.DefaultY != partnameOther.DefaultY {
		diffs = append(diffs, partname.GongMarshallField(stage, "DefaultY"))
	}
	if partname.RelativeX != partnameOther.RelativeX {
		diffs = append(diffs, partname.GongMarshallField(stage, "RelativeX"))
	}
	if partname.RelativeY != partnameOther.RelativeY {
		diffs = append(diffs, partname.GongMarshallField(stage, "RelativeY"))
	}
	if partname.FontFamily != partnameOther.FontFamily {
		diffs = append(diffs, partname.GongMarshallField(stage, "FontFamily"))
	}
	if partname.FontSize != partnameOther.FontSize {
		diffs = append(diffs, partname.GongMarshallField(stage, "FontSize"))
	}
	if partname.Color != partnameOther.Color {
		diffs = append(diffs, partname.GongMarshallField(stage, "Color"))
	}
	if partname.Value != partnameOther.Value {
		diffs = append(diffs, partname.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (player *Player) GongDiff(stage *Stage, playerOther *Player) (diffs []string) {
	// insertion point for field diffs
	if player.Name != playerOther.Name {
		diffs = append(diffs, player.GongMarshallField(stage, "Name"))
	}
	if player.Id != playerOther.Id {
		diffs = append(diffs, player.GongMarshallField(stage, "Id"))
	}
	if player.PlayerName != playerOther.PlayerName {
		diffs = append(diffs, player.GongMarshallField(stage, "PlayerName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scoreinstrument *ScoreInstrument) GongDiff(stage *Stage, scoreinstrumentOther *ScoreInstrument) (diffs []string) {
	// insertion point for field diffs
	if scoreinstrument.Name != scoreinstrumentOther.Name {
		diffs = append(diffs, scoreinstrument.GongMarshallField(stage, "Name"))
	}
	if scoreinstrument.Id != scoreinstrumentOther.Id {
		diffs = append(diffs, scoreinstrument.GongMarshallField(stage, "Id"))
	}
	if scoreinstrument.InstrumentName != scoreinstrumentOther.InstrumentName {
		diffs = append(diffs, scoreinstrument.GongMarshallField(stage, "InstrumentName"))
	}
	if scoreinstrument.InstrumentAbbreviation != scoreinstrumentOther.InstrumentAbbreviation {
		diffs = append(diffs, scoreinstrument.GongMarshallField(stage, "InstrumentAbbreviation"))
	}
	if scoreinstrument.InstrumentSound != scoreinstrumentOther.InstrumentSound {
		diffs = append(diffs, scoreinstrument.GongMarshallField(stage, "InstrumentSound"))
	}
	if scoreinstrument.Ensemble != scoreinstrumentOther.Ensemble {
		diffs = append(diffs, scoreinstrument.GongMarshallField(stage, "Ensemble"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scorepart *ScorePart) GongDiff(stage *Stage, scorepartOther *ScorePart) (diffs []string) {
	// insertion point for field diffs
	if scorepart.Name != scorepartOther.Name {
		diffs = append(diffs, scorepart.GongMarshallField(stage, "Name"))
	}
	if scorepart.Id != scorepartOther.Id {
		diffs = append(diffs, scorepart.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scorepart, "PartLink", scorepartOther.PartLink, scorepart.PartLink); ops != "" {
		diffs = append(diffs, ops)
	}
	if scorepart.PartName != scorepartOther.PartName {
		diffs = append(diffs, scorepart.GongMarshallField(stage, "PartName"))
	}
	if scorepart.PartAbbreviation != scorepartOther.PartAbbreviation {
		diffs = append(diffs, scorepart.GongMarshallField(stage, "PartAbbreviation"))
	}
	if scorepart.Group != scorepartOther.Group {
		diffs = append(diffs, scorepart.GongMarshallField(stage, "Group"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scorepart, "ScoreInstrument", scorepartOther.ScoreInstrument, scorepart.ScoreInstrument); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, scorepart, "Player", scorepartOther.Player, scorepart.Player); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scorepartwise *ScorePartwise) GongDiff(stage *Stage, scorepartwiseOther *ScorePartwise) (diffs []string) {
	// insertion point for field diffs
	if scorepartwise.Name != scorepartwiseOther.Name {
		diffs = append(diffs, scorepartwise.GongMarshallField(stage, "Name"))
	}
	if scorepartwise.Version != scorepartwiseOther.Version {
		diffs = append(diffs, scorepartwise.GongMarshallField(stage, "Version"))
	}
	if scorepartwise.Work != scorepartwiseOther.Work {
		diffs = append(diffs, scorepartwise.GongMarshallField(stage, "Work"))
	}
	if scorepartwise.MovementNumber != scorepartwiseOther.MovementNumber {
		diffs = append(diffs, scorepartwise.GongMarshallField(stage, "MovementNumber"))
	}
	if scorepartwise.MovementTitle != scorepartwiseOther.MovementTitle {
		diffs = append(diffs, scorepartwise.GongMarshallField(stage, "MovementTitle"))
	}
	if scorepartwise.Defaults != scorepartwiseOther.Defaults {
		diffs = append(diffs, scorepartwise.GongMarshallField(stage, "Defaults"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scorepartwise, "Credit", scorepartwiseOther.Credit, scorepartwise.Credit); ops != "" {
		diffs = append(diffs, ops)
	}
	if scorepartwise.PartList != scorepartwiseOther.PartList {
		diffs = append(diffs, scorepartwise.GongMarshallField(stage, "PartList"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scorepartwise, "Part", scorepartwiseOther.Part, scorepartwise.Part); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scoretimewise *ScoreTimewise) GongDiff(stage *Stage, scoretimewiseOther *ScoreTimewise) (diffs []string) {
	// insertion point for field diffs
	if scoretimewise.Name != scoretimewiseOther.Name {
		diffs = append(diffs, scoretimewise.GongMarshallField(stage, "Name"))
	}
	if scoretimewise.Version != scoretimewiseOther.Version {
		diffs = append(diffs, scoretimewise.GongMarshallField(stage, "Version"))
	}
	if scoretimewise.Work != scoretimewiseOther.Work {
		diffs = append(diffs, scoretimewise.GongMarshallField(stage, "Work"))
	}
	if scoretimewise.MovementNumber != scoretimewiseOther.MovementNumber {
		diffs = append(diffs, scoretimewise.GongMarshallField(stage, "MovementNumber"))
	}
	if scoretimewise.MovementTitle != scoretimewiseOther.MovementTitle {
		diffs = append(diffs, scoretimewise.GongMarshallField(stage, "MovementTitle"))
	}
	if scoretimewise.Defaults != scoretimewiseOther.Defaults {
		diffs = append(diffs, scoretimewise.GongMarshallField(stage, "Defaults"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scoretimewise, "Credit", scoretimewiseOther.Credit, scoretimewise.Credit); ops != "" {
		diffs = append(diffs, ops)
	}
	if scoretimewise.PartList != scoretimewiseOther.PartList {
		diffs = append(diffs, scoretimewise.GongMarshallField(stage, "PartList"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scoretimewise, "Measure", scoretimewiseOther.Measure, scoretimewise.Measure); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scoretimewisemeasure *ScoreTimewiseMeasure) GongDiff(stage *Stage, scoretimewisemeasureOther *ScoreTimewiseMeasure) (diffs []string) {
	// insertion point for field diffs
	if scoretimewisemeasure.Name != scoretimewisemeasureOther.Name {
		diffs = append(diffs, scoretimewisemeasure.GongMarshallField(stage, "Name"))
	}
	if scoretimewisemeasure.Number != scoretimewisemeasureOther.Number {
		diffs = append(diffs, scoretimewisemeasure.GongMarshallField(stage, "Number"))
	}
	if scoretimewisemeasure.Text != scoretimewisemeasureOther.Text {
		diffs = append(diffs, scoretimewisemeasure.GongMarshallField(stage, "Text"))
	}
	if scoretimewisemeasure.Width != scoretimewisemeasureOther.Width {
		diffs = append(diffs, scoretimewisemeasure.GongMarshallField(stage, "Width"))
	}
	if scoretimewisemeasure.Id != scoretimewisemeasureOther.Id {
		diffs = append(diffs, scoretimewisemeasure.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scoretimewisemeasure, "Part", scoretimewisemeasureOther.Part, scoretimewisemeasure.Part); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scoretimewisemeasurepart *ScoreTimewiseMeasurePart) GongDiff(stage *Stage, scoretimewisemeasurepartOther *ScoreTimewiseMeasurePart) (diffs []string) {
	// insertion point for field diffs
	if scoretimewisemeasurepart.Name != scoretimewisemeasurepartOther.Name {
		diffs = append(diffs, scoretimewisemeasurepart.GongMarshallField(stage, "Name"))
	}
	if scoretimewisemeasurepart.RefScorePart != scoretimewisemeasurepartOther.RefScorePart {
		diffs = append(diffs, scoretimewisemeasurepart.GongMarshallField(stage, "RefScorePart"))
	}
	if scoretimewisemeasurepart.Id != scoretimewisemeasurepartOther.Id {
		diffs = append(diffs, scoretimewisemeasurepart.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (work *Work) GongDiff(stage *Stage, workOther *Work) (diffs []string) {
	// insertion point for field diffs
	if work.Name != workOther.Name {
		diffs = append(diffs, work.GongMarshallField(stage, "Name"))
	}
	if work.WorkNumber != workOther.WorkNumber {
		diffs = append(diffs, work.GongMarshallField(stage, "WorkNumber"))
	}
	if work.WorkTitle != workOther.WorkTitle {
		diffs = append(diffs, work.GongMarshallField(stage, "WorkTitle"))
	}
	if work.Opus != workOther.Opus {
		diffs = append(diffs, work.GongMarshallField(stage, "Opus"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
