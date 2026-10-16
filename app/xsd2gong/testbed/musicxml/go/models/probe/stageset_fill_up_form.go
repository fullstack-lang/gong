// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.MusicXMLDocument:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ScorePartwise", inst.ScorePartwise, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScorePartwise](), probe.formStage)
		StageSetAssociationFieldToForm("ScoreTimewise", inst.ScoreTimewise, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScoreTimewise](), probe.formStage)
	case *score.Credit:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Page", inst.Page, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CreditType", inst.CreditType, probe.formStage, formGroup)

		{
			// Slice of pointers: Link
			div := (&form.FormDiv{Name: "Link"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Link {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Link",
				Label: "Link",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Bookmark
			div := (&form.FormDiv{Name: "Bookmark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bookmark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bookmark",
				Label: "Bookmark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("CreditImage", inst.CreditImage, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Image](), probe.formStage)

		{
			// Slice of pointers: CreditWords
			div := (&form.FormDiv{Name: "CreditWords"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.CreditWords {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "CreditWords",
				Label: "CreditWords",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: CreditSymbol
			div := (&form.FormDiv{Name: "CreditSymbol"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.CreditSymbol {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "CreditSymbol",
				Label: "CreditSymbol",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScorePartwises {
				for _, target := range src.Credit {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePartwise", "Credit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewises {
				for _, target := range src.Credit {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewise", "Credit", refNames, formGroup, probe.formStage)
		}
	case *score.Defaults:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Scaling", inst.Scaling, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Scaling](), probe.formStage)
		StageSetAssociationFieldToForm("ConcertScore", inst.ConcertScore, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("PageLayout", inst.PageLayout, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.PageLayout](), probe.formStage)
		StageSetAssociationFieldToForm("SystemLayout", inst.SystemLayout, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.SystemLayout](), probe.formStage)

		{
			// Slice of pointers: StaffLayout
			div := (&form.FormDiv{Name: "StaffLayout"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.StaffLayout {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "StaffLayout",
				Label: "StaffLayout",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Appearance", inst.Appearance, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Appearance](), probe.formStage)
		StageSetAssociationFieldToForm("MusicFont", inst.MusicFont, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.EmptyFont](), probe.formStage)
		StageSetAssociationFieldToForm("WordFont", inst.WordFont, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.EmptyFont](), probe.formStage)

		{
			// Slice of pointers: LyricFont
			div := (&form.FormDiv{Name: "LyricFont"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.LyricFont {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "LyricFont",
				Label: "LyricFont",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: LyricLanguage
			div := (&form.FormDiv{Name: "LyricLanguage"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.LyricLanguage {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "LyricLanguage",
				Label: "LyricLanguage",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScorePartwises {
				if src.Defaults == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePartwise", "Defaults", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewises {
				if src.Defaults == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewise", "Defaults", refNames, formGroup, probe.formStage)
		}
	case *score.EmptyFont:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				if src.MusicFont == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "MusicFont", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				if src.WordFont == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "WordFont", refNames, formGroup, probe.formStage)
		}
	case *score.GroupBarline:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartGroups {
				if src.GroupBarline == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartGroup", "GroupBarline", refNames, formGroup, probe.formStage)
		}
	case *score.GroupName:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartGroups {
				if src.GroupName == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartGroup", "GroupName", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartGroups {
				if src.GroupAbbreviation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartGroup", "GroupAbbreviation", refNames, formGroup, probe.formStage)
		}
	case *score.GroupSymbol:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartGroups {
				if src.GroupSymbol == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartGroup", "GroupSymbol", refNames, formGroup, probe.formStage)
		}
	case *score.InstrumentLink:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RefScoreInstrument", inst.RefScoreInstrument, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScoreInstrument](), probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartLinks {
				for _, target := range src.InstrumentLink {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartLink", "InstrumentLink", refNames, formGroup, probe.formStage)
		}
	case *score.LyricFont:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameAttr", inst.NameAttr, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				for _, target := range src.LyricFont {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "LyricFont", refNames, formGroup, probe.formStage)
		}
	case *score.LyricLanguage:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameAttr", inst.NameAttr, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				for _, target := range src.LyricLanguage {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "LyricLanguage", refNames, formGroup, probe.formStage)
		}
	case *score.Measure:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Implicit", inst.Implicit, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("NonControlling", inst.NonControlling, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: MusicData
			div := (&form.FormDiv{Name: "MusicData"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MusicData {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MusicData",
				Label: "MusicData",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Parts {
				for _, target := range src.Measure {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Part", "Measure", refNames, formGroup, probe.formStage)
		}
	case *score.MeasureMusicDataChoice:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Note](), probe.formStage)
		StageSetAssociationFieldToForm("Backup", inst.Backup, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Backup](), probe.formStage)
		StageSetAssociationFieldToForm("Forward", inst.Forward, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Forward](), probe.formStage)
		StageSetAssociationFieldToForm("Direction", inst.Direction, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Direction](), probe.formStage)
		StageSetAssociationFieldToForm("Attributes", inst.Attributes, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Attributes](), probe.formStage)
		StageSetAssociationFieldToForm("Harmony", inst.Harmony, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Harmony](), probe.formStage)
		StageSetAssociationFieldToForm("FiguredBass", inst.FiguredBass, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.FiguredBass](), probe.formStage)
		StageSetAssociationFieldToForm("Print", inst.Print, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Print](), probe.formStage)
		StageSetAssociationFieldToForm("Sound", inst.Sound, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Sound](), probe.formStage)
		StageSetAssociationFieldToForm("Listening", inst.Listening, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Listening](), probe.formStage)
		StageSetAssociationFieldToForm("Barline", inst.Barline, formGroup, probe.stageSet.BarlineStage.GetInstancesSet[*barline.Barline](), probe.formStage)
		StageSetAssociationFieldToForm("Grouping", inst.Grouping, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Grouping](), probe.formStage)
		StageSetAssociationFieldToForm("Link", inst.Link, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Link](), probe.formStage)
		StageSetAssociationFieldToForm("Bookmark", inst.Bookmark, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Bookmark](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Measures {
				for _, target := range src.MusicData {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Measure", "MusicData", refNames, formGroup, probe.formStage)
		}
	case *score.Opus:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Show", inst.Show, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Actuate", inst.Actuate, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Works {
				if src.Opus == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Work", "Opus", refNames, formGroup, probe.formStage)
		}
	case *score.Part:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RefScorePart", inst.RefScorePart, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScorePart](), probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Measure
			div := (&form.FormDiv{Name: "Measure"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Measure {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Measure",
				Label: "Measure",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScorePartwises {
				for _, target := range src.Part {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePartwise", "Part", refNames, formGroup, probe.formStage)
		}
	case *score.PartGroup:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("GroupName", inst.GroupName, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.GroupName](), probe.formStage)
		StageSetAssociationFieldToForm("GroupNameDisplay", inst.GroupNameDisplay, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), probe.formStage)
		StageSetAssociationFieldToForm("GroupAbbreviation", inst.GroupAbbreviation, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.GroupName](), probe.formStage)
		StageSetAssociationFieldToForm("GroupAbbreviationDisplay", inst.GroupAbbreviationDisplay, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), probe.formStage)
		StageSetAssociationFieldToForm("GroupSymbol", inst.GroupSymbol, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.GroupSymbol](), probe.formStage)
		StageSetAssociationFieldToForm("GroupBarline", inst.GroupBarline, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.GroupBarline](), probe.formStage)
		StageSetAssociationFieldToForm("GroupTime", inst.GroupTime, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartLists {
				for _, target := range src.PartGroup {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartList", "PartGroup", refNames, formGroup, probe.formStage)
		}
	case *score.PartLink:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Show", inst.Show, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Actuate", inst.Actuate, formGroup, probe.formStage)

		{
			// Slice of pointers: InstrumentLink
			div := (&form.FormDiv{Name: "InstrumentLink"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.InstrumentLink {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "InstrumentLink",
				Label: "InstrumentLink",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("GroupLink", inst.GroupLink, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				for _, target := range src.PartLink {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "PartLink", refNames, formGroup, probe.formStage)
		}
	case *score.PartList:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: PartGroup
			div := (&form.FormDiv{Name: "PartGroup"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PartGroup {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PartGroup",
				Label: "PartGroup",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ScorePart
			div := (&form.FormDiv{Name: "ScorePart"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ScorePart {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ScorePart",
				Label: "ScorePart",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScorePartwises {
				if src.PartList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePartwise", "PartList", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewises {
				if src.PartList == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewise", "PartList", refNames, formGroup, probe.formStage)
		}
	case *score.PartName:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				if src.PartName == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "PartName", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				if src.PartAbbreviation == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "PartAbbreviation", refNames, formGroup, probe.formStage)
		}
	case *score.Player:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PlayerName", inst.PlayerName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				for _, target := range src.Player {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "Player", refNames, formGroup, probe.formStage)
		}
	case *score.ScoreInstrument:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InstrumentName", inst.InstrumentName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InstrumentAbbreviation", inst.InstrumentAbbreviation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InstrumentSound", inst.InstrumentSound, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Solo", inst.Solo, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetBasicFieldtoForm("Ensemble", inst.Ensemble, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("VirtualInstrument", inst.VirtualInstrument, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.VirtualInstrument](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.InstrumentLinks {
				if src.RefScoreInstrument == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.InstrumentLink", "RefScoreInstrument", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				for _, target := range src.ScoreInstrument {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "ScoreInstrument", refNames, formGroup, probe.formStage)
		}
	case *score.ScorePart:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identification", inst.Identification, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Identification](), probe.formStage)

		{
			// Slice of pointers: PartLink
			div := (&form.FormDiv{Name: "PartLink"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PartLink {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PartLink",
				Label: "PartLink",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PartName", inst.PartName, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.PartName](), probe.formStage)
		StageSetAssociationFieldToForm("PartNameDisplay", inst.PartNameDisplay, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), probe.formStage)
		StageSetAssociationFieldToForm("PartAbbreviation", inst.PartAbbreviation, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.PartName](), probe.formStage)
		StageSetAssociationFieldToForm("PartAbbreviationDisplay", inst.PartAbbreviationDisplay, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), probe.formStage)
		StageSetBasicFieldtoForm("Group", inst.Group, probe.formStage, formGroup)

		{
			// Slice of pointers: ScoreInstrument
			div := (&form.FormDiv{Name: "ScoreInstrument"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ScoreInstrument {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ScoreInstrument",
				Label: "ScoreInstrument",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Player
			div := (&form.FormDiv{Name: "Player"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Player {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Player",
				Label: "Player",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: MidiDevice
			div := (&form.FormDiv{Name: "MidiDevice"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MidiDevice {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MidiDevice",
				Label: "MidiDevice",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: MidiInstrument
			div := (&form.FormDiv{Name: "MidiInstrument"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MidiInstrument {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MidiInstrument",
				Label: "MidiInstrument",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Parts {
				if src.RefScorePart == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Part", "RefScorePart", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartLists {
				for _, target := range src.ScorePart {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartList", "ScorePart", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				if src.RefScorePart == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "RefScorePart", refNames, formGroup, probe.formStage)
		}
	case *score.ScorePartwise:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Version", inst.Version, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Work", inst.Work, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.Work](), probe.formStage)
		StageSetBasicFieldtoForm("MovementNumber", inst.MovementNumber, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementTitle", inst.MovementTitle, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identification", inst.Identification, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Identification](), probe.formStage)
		StageSetAssociationFieldToForm("Defaults", inst.Defaults, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.Defaults](), probe.formStage)

		{
			// Slice of pointers: Credit
			div := (&form.FormDiv{Name: "Credit"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Credit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Credit",
				Label: "Credit",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PartList", inst.PartList, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.PartList](), probe.formStage)

		{
			// Slice of pointers: Part
			div := (&form.FormDiv{Name: "Part"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Part {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Part",
				Label: "Part",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.MusicXMLDocuments {
				if src.ScorePartwise == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.MusicXMLDocument", "ScorePartwise", refNames, formGroup, probe.formStage)
		}
	case *score.ScoreTimewise:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Version", inst.Version, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Work", inst.Work, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.Work](), probe.formStage)
		StageSetBasicFieldtoForm("MovementNumber", inst.MovementNumber, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MovementTitle", inst.MovementTitle, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Identification", inst.Identification, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Identification](), probe.formStage)
		StageSetAssociationFieldToForm("Defaults", inst.Defaults, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.Defaults](), probe.formStage)

		{
			// Slice of pointers: Credit
			div := (&form.FormDiv{Name: "Credit"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Credit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Credit",
				Label: "Credit",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PartList", inst.PartList, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.PartList](), probe.formStage)

		{
			// Slice of pointers: Measure
			div := (&form.FormDiv{Name: "Measure"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Measure {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Measure",
				Label: "Measure",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.MusicXMLDocuments {
				if src.ScoreTimewise == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.MusicXMLDocument", "ScoreTimewise", refNames, formGroup, probe.formStage)
		}
	case *score.ScoreTimewiseMeasure:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Implicit", inst.Implicit, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("NonControlling", inst.NonControlling, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Part
			div := (&form.FormDiv{Name: "Part"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Part {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Part",
				Label: "Part",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewises {
				for _, target := range src.Measure {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewise", "Measure", refNames, formGroup, probe.formStage)
		}
	case *score.ScoreTimewiseMeasurePart:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RefScorePart", inst.RefScorePart, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.ScorePart](), probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Note
			div := (&form.FormDiv{Name: "Note"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Note {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Note",
				Label: "Note",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Backup
			div := (&form.FormDiv{Name: "Backup"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Backup {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Backup",
				Label: "Backup",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Forward
			div := (&form.FormDiv{Name: "Forward"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Forward {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Forward",
				Label: "Forward",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Direction
			div := (&form.FormDiv{Name: "Direction"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Direction {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Direction",
				Label: "Direction",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Attributes
			div := (&form.FormDiv{Name: "Attributes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Attributes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Attributes",
				Label: "Attributes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Harmony
			div := (&form.FormDiv{Name: "Harmony"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Harmony {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Harmony",
				Label: "Harmony",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: FiguredBass
			div := (&form.FormDiv{Name: "FiguredBass"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FiguredBass {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FiguredBass",
				Label: "FiguredBass",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Print
			div := (&form.FormDiv{Name: "Print"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Print {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Print",
				Label: "Print",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Sound
			div := (&form.FormDiv{Name: "Sound"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sound {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sound",
				Label: "Sound",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Listening
			div := (&form.FormDiv{Name: "Listening"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Listening {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Listening",
				Label: "Listening",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Barline
			div := (&form.FormDiv{Name: "Barline"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Barline {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Barline",
				Label: "Barline",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Grouping
			div := (&form.FormDiv{Name: "Grouping"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Grouping {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Grouping",
				Label: "Grouping",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Link
			div := (&form.FormDiv{Name: "Link"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Link {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Link",
				Label: "Link",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Bookmark
			div := (&form.FormDiv{Name: "Bookmark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bookmark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bookmark",
				Label: "Bookmark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasures {
				for _, target := range src.Part {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasure", "Part", refNames, formGroup, probe.formStage)
		}
	case *score.Work:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("WorkNumber", inst.WorkNumber, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("WorkTitle", inst.WorkTitle, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Opus", inst.Opus, formGroup, probe.stageSet.ScoreStage.GetInstancesSet[*score.Opus](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScorePartwises {
				if src.Work == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePartwise", "Work", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewises {
				if src.Work == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewise", "Work", refNames, formGroup, probe.formStage)
		}
	case *note.Accidental:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Cautionary", inst.Cautionary, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Editorial", inst.Editorial, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Accidental == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Accidental", refNames, formGroup, probe.formStage)
		}
	case *note.AccidentalMark:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.AccidentalMark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "AccidentalMark", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.AccidentalMark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "AccidentalMark", refNames, formGroup, probe.formStage)
		}
	case *note.Arpeggiate:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Direction", inst.Direction, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Unbroken", inst.Unbroken, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Arpeggiate {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Arpeggiate", refNames, formGroup, probe.formStage)
		}
	case *note.Arrow:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ArrowDirection", inst.ArrowDirection, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ArrowStyle", inst.ArrowStyle, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("Arrowhead", inst.Arrowhead, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetEnumStringFieldToForm("CircularArrow", inst.CircularArrow, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Arrow {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Arrow", refNames, formGroup, probe.formStage)
		}
	case *note.Articulations:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Accent
			div := (&form.FormDiv{Name: "Accent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Accent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Accent",
				Label: "Accent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: StrongAccent
			div := (&form.FormDiv{Name: "StrongAccent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.StrongAccent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "StrongAccent",
				Label: "StrongAccent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Staccato
			div := (&form.FormDiv{Name: "Staccato"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Staccato {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Staccato",
				Label: "Staccato",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tenuto
			div := (&form.FormDiv{Name: "Tenuto"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tenuto {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tenuto",
				Label: "Tenuto",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DetachedLegato
			div := (&form.FormDiv{Name: "DetachedLegato"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DetachedLegato {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DetachedLegato",
				Label: "DetachedLegato",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Staccatissimo
			div := (&form.FormDiv{Name: "Staccatissimo"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Staccatissimo {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Staccatissimo",
				Label: "Staccatissimo",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Spiccato
			div := (&form.FormDiv{Name: "Spiccato"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Spiccato {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Spiccato",
				Label: "Spiccato",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Scoop
			div := (&form.FormDiv{Name: "Scoop"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Scoop {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Scoop",
				Label: "Scoop",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Plop
			div := (&form.FormDiv{Name: "Plop"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Plop {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Plop",
				Label: "Plop",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Doit
			div := (&form.FormDiv{Name: "Doit"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Doit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Doit",
				Label: "Doit",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Falloff
			div := (&form.FormDiv{Name: "Falloff"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Falloff {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Falloff",
				Label: "Falloff",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: BreathMark
			div := (&form.FormDiv{Name: "BreathMark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.BreathMark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "BreathMark",
				Label: "BreathMark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Caesura
			div := (&form.FormDiv{Name: "Caesura"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Caesura {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Caesura",
				Label: "Caesura",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Stress
			div := (&form.FormDiv{Name: "Stress"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Stress {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Stress",
				Label: "Stress",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Unstress
			div := (&form.FormDiv{Name: "Unstress"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Unstress {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Unstress",
				Label: "Unstress",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SoftAccent
			div := (&form.FormDiv{Name: "SoftAccent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SoftAccent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SoftAccent",
				Label: "SoftAccent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OtherArticulation
			div := (&form.FormDiv{Name: "OtherArticulation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OtherArticulation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OtherArticulation",
				Label: "OtherArticulation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Articulations {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Articulations", refNames, formGroup, probe.formStage)
		}
	case *note.Assess:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeOnly", inst.TimeOnly, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Listens {
				for _, target := range src.Assess {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Listen", "Assess", refNames, formGroup, probe.formStage)
		}
	case *note.Backup:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Backup == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Backup", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Backup {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Backup", refNames, formGroup, probe.formStage)
		}
	case *note.Beam:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Repeater", inst.Repeater, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Fan", inst.Fan, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				for _, target := range src.Beam {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Beam", refNames, formGroup, probe.formStage)
		}
	case *note.Bend:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Shape", inst.Shape, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FirstBeat", inst.FirstBeat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LastBeat", inst.LastBeat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BendAlter", inst.BendAlter, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PreBend", inst.PreBend, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("Release", inst.Release, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Release](), probe.formStage)
		StageSetAssociationFieldToForm("WithBar", inst.WithBar, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.PlacementText](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Bend {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Bend", refNames, formGroup, probe.formStage)
		}
	case *note.BreathMark:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.BreathMark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "BreathMark", refNames, formGroup, probe.formStage)
		}
	case *note.Caesura:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Caesura {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Caesura", refNames, formGroup, probe.formStage)
		}
	case *note.Elision:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Lyrics {
				for _, target := range src.Elision {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Lyric", "Elision", refNames, formGroup, probe.formStage)
		}
	case *note.EmptyLine:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("LineShape", inst.LineShape, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("LineType", inst.LineType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("LineLength", inst.LineLength, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DashLength", inst.DashLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SpaceLength", inst.SpaceLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Scoop {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Scoop", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Plop {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Plop", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Doit {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Doit", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Falloff {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Falloff", refNames, formGroup, probe.formStage)
		}
	case *note.Extend:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Figures {
				if src.Extend == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Figure", "Extend", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Lyrics {
				if src.Extend == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Lyric", "Extend", refNames, formGroup, probe.formStage)
		}
	case *note.Figure:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Prefix", inst.Prefix, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StyleText](), probe.formStage)
		StageSetAssociationFieldToForm("FigureNumber", inst.FigureNumber, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StyleText](), probe.formStage)
		StageSetAssociationFieldToForm("Suffix", inst.Suffix, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StyleText](), probe.formStage)
		StageSetAssociationFieldToForm("Extend", inst.Extend, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Extend](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.FiguredBasss {
				for _, target := range src.Figure {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.FiguredBass", "Figure", refNames, formGroup, probe.formStage)
		}
	case *note.FiguredBass:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintDot", inst.PrintDot, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintLyric", inst.PrintLyric, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintSpacing", inst.PrintSpacing, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Figure
			div := (&form.FormDiv{Name: "Figure"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Figure {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Figure",
				Label: "Figure",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.FiguredBass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "FiguredBass", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.FiguredBass {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "FiguredBass", refNames, formGroup, probe.formStage)
		}
	case *note.Forward:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Voice", inst.Voice, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Forward == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Forward", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Forward {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Forward", refNames, formGroup, probe.formStage)
		}
	case *note.Glissando:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("LineType", inst.LineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DashLength", inst.DashLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SpaceLength", inst.SpaceLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Glissando {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Glissando", refNames, formGroup, probe.formStage)
		}
	case *note.Grace:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StealTimePrevious", inst.StealTimePrevious, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StealTimeFollowing", inst.StealTimeFollowing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MakeTime", inst.MakeTime, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Slash", inst.Slash, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Grace == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Grace", refNames, formGroup, probe.formStage)
		}
	case *note.HammerOnPullOff:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.HammerOn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "HammerOn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.PullOff {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "PullOff", refNames, formGroup, probe.formStage)
		}
	case *note.Handbell:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Handbell {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Handbell", refNames, formGroup, probe.formStage)
		}
	case *note.HarmonClosed:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.HarmonMutes {
				if src.HarmonClosed == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.HarmonMute", "HarmonClosed", refNames, formGroup, probe.formStage)
		}
	case *note.HarmonMute:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("HarmonClosed", inst.HarmonClosed, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.HarmonClosed](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.HarmonMute {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "HarmonMute", refNames, formGroup, probe.formStage)
		}
	case *note.Harmonic:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("Natural", inst.Natural, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("Artificial", inst.Artificial, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("BasePitch", inst.BasePitch, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("TouchingPitch", inst.TouchingPitch, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("SoundingPitch", inst.SoundingPitch, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Harmonic {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Harmonic", refNames, formGroup, probe.formStage)
		}
	case *note.HeelToe:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Substitution", inst.Substitution, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Heel {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Heel", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Toe {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Toe", refNames, formGroup, probe.formStage)
		}
	case *note.Hole:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("HoleType", inst.HoleType, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("HoleClosed", inst.HoleClosed, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.HoleClosed](), probe.formStage)
		StageSetBasicFieldtoForm("HoleShape", inst.HoleShape, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Hole {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Hole", refNames, formGroup, probe.formStage)
		}
	case *note.HoleClosed:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Holes {
				if src.HoleClosed == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Hole", "HoleClosed", refNames, formGroup, probe.formStage)
		}
	case *note.Instrument:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				for _, target := range src.Instrument {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Instrument", refNames, formGroup, probe.formStage)
		}
	case *note.Listen:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Assess
			div := (&form.FormDiv{Name: "Assess"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Assess {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Assess",
				Label: "Assess",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Wait
			div := (&form.FormDiv{Name: "Wait"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Wait {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Wait",
				Label: "Wait",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OtherListen
			div := (&form.FormDiv{Name: "OtherListen"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OtherListen {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OtherListen",
				Label: "OtherListen",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Listen == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Listen", refNames, formGroup, probe.formStage)
		}
	case *note.Lyric:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameAttr", inst.NameAttr, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeOnly", inst.TimeOnly, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Syllabic", inst.Syllabic, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("Text", inst.Text, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.TextElementData](), probe.formStage)

		{
			// Slice of pointers: Elision
			div := (&form.FormDiv{Name: "Elision"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Elision {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Elision",
				Label: "Elision",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Extend", inst.Extend, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Extend](), probe.formStage)
		StageSetAssociationFieldToForm("Laughing", inst.Laughing, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("Humming", inst.Humming, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("EndLine", inst.EndLine, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("EndParagraph", inst.EndParagraph, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				for _, target := range src.Lyric {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Lyric", refNames, formGroup, probe.formStage)
		}
	case *note.Mordent:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("StartNote", inst.StartNote, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TrillStep", inst.TrillStep, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TwoNoteTurn", inst.TwoNoteTurn, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SecondBeat", inst.SecondBeat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LastBeat", inst.LastBeat, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Long", inst.Long, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Approach", inst.Approach, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Departure", inst.Departure, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.Mordent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "Mordent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.InvertedMordent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "InvertedMordent", refNames, formGroup, probe.formStage)
		}
	case *note.NonArpeggiate:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.NonArpeggiate {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "NonArpeggiate", refNames, formGroup, probe.formStage)
		}
	case *note.Notations:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)

		{
			// Slice of pointers: Tied
			div := (&form.FormDiv{Name: "Tied"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tied {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tied",
				Label: "Tied",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Slur
			div := (&form.FormDiv{Name: "Slur"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Slur {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Slur",
				Label: "Slur",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tuplet
			div := (&form.FormDiv{Name: "Tuplet"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tuplet {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tuplet",
				Label: "Tuplet",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Glissando
			div := (&form.FormDiv{Name: "Glissando"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Glissando {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Glissando",
				Label: "Glissando",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Slide
			div := (&form.FormDiv{Name: "Slide"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Slide {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Slide",
				Label: "Slide",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Ornaments
			div := (&form.FormDiv{Name: "Ornaments"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Ornaments {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Ornaments",
				Label: "Ornaments",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Technical
			div := (&form.FormDiv{Name: "Technical"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Technical {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Technical",
				Label: "Technical",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Articulations
			div := (&form.FormDiv{Name: "Articulations"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Articulations {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Articulations",
				Label: "Articulations",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Dynamics
			div := (&form.FormDiv{Name: "Dynamics"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Dynamics {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Dynamics",
				Label: "Dynamics",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fermata
			div := (&form.FormDiv{Name: "Fermata"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fermata {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fermata",
				Label: "Fermata",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Arpeggiate
			div := (&form.FormDiv{Name: "Arpeggiate"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Arpeggiate {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Arpeggiate",
				Label: "Arpeggiate",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: NonArpeggiate
			div := (&form.FormDiv{Name: "NonArpeggiate"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NonArpeggiate {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NonArpeggiate",
				Label: "NonArpeggiate",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AccidentalMark
			div := (&form.FormDiv{Name: "AccidentalMark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AccidentalMark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AccidentalMark",
				Label: "AccidentalMark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OtherNotation
			div := (&form.FormDiv{Name: "OtherNotation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OtherNotation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OtherNotation",
				Label: "OtherNotation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				for _, target := range src.Notations {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Notations", refNames, formGroup, probe.formStage)
		}
	case *note.Note:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintLeger", inst.PrintLeger, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Dynamics", inst.Dynamics, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndDynamics", inst.EndDynamics, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Attack", inst.Attack, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Release", inst.Release, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeOnly", inst.TimeOnly, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Pizzicato", inst.Pizzicato, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintDot", inst.PrintDot, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintLyric", inst.PrintLyric, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintSpacing", inst.PrintSpacing, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Grace", inst.Grace, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Grace](), probe.formStage)
		StageSetAssociationFieldToForm("Chord", inst.Chord, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetAssociationFieldToForm("Pitch", inst.Pitch, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Pitch](), probe.formStage)
		StageSetAssociationFieldToForm("Unpitched", inst.Unpitched, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Unpitched](), probe.formStage)
		StageSetAssociationFieldToForm("Rest", inst.Rest, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Rest](), probe.formStage)

		{
			// Slice of pointers: Tie
			div := (&form.FormDiv{Name: "Tie"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tie {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tie",
				Label: "Tie",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Cue", inst.Cue, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)

		{
			// Slice of pointers: Instrument
			div := (&form.FormDiv{Name: "Instrument"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Instrument {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Instrument",
				Label: "Instrument",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Voice", inst.Voice, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Type", inst.Type, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.NoteType](), probe.formStage)

		{
			// Slice of pointers: Dot
			div := (&form.FormDiv{Name: "Dot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Dot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Dot",
				Label: "Dot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Accidental", inst.Accidental, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Accidental](), probe.formStage)
		StageSetAssociationFieldToForm("TimeModification", inst.TimeModification, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.TimeModification](), probe.formStage)
		StageSetAssociationFieldToForm("Stem", inst.Stem, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Stem](), probe.formStage)
		StageSetAssociationFieldToForm("Notehead", inst.Notehead, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Notehead](), probe.formStage)
		StageSetAssociationFieldToForm("NoteheadText", inst.NoteheadText, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.NoteheadText](), probe.formStage)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)

		{
			// Slice of pointers: Beam
			div := (&form.FormDiv{Name: "Beam"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Beam {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Beam",
				Label: "Beam",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Notations
			div := (&form.FormDiv{Name: "Notations"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Notations {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Notations",
				Label: "Notations",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Lyric
			div := (&form.FormDiv{Name: "Lyric"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Lyric {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Lyric",
				Label: "Lyric",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Play", inst.Play, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Play](), probe.formStage)
		StageSetAssociationFieldToForm("Listen", inst.Listen, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.Listen](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Note", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Note {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Note", refNames, formGroup, probe.formStage)
		}
	case *note.NoteType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Type == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Type", refNames, formGroup, probe.formStage)
		}
	case *note.Notehead:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Filled", inst.Filled, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Notehead == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Notehead", refNames, formGroup, probe.formStage)
		}
	case *note.NoteheadText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: DisplayText
			div := (&form.FormDiv{Name: "DisplayText"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DisplayText {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DisplayText",
				Label: "DisplayText",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AccidentalText
			div := (&form.FormDiv{Name: "AccidentalText"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AccidentalText {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AccidentalText",
				Label: "AccidentalText",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.NoteheadText == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "NoteheadText", refNames, formGroup, probe.formStage)
		}
	case *note.Ornaments:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: TrillMark
			div := (&form.FormDiv{Name: "TrillMark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TrillMark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TrillMark",
				Label: "TrillMark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Turn
			div := (&form.FormDiv{Name: "Turn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Turn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Turn",
				Label: "Turn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DelayedTurn
			div := (&form.FormDiv{Name: "DelayedTurn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DelayedTurn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DelayedTurn",
				Label: "DelayedTurn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: InvertedTurn
			div := (&form.FormDiv{Name: "InvertedTurn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.InvertedTurn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "InvertedTurn",
				Label: "InvertedTurn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DelayedInvertedTurn
			div := (&form.FormDiv{Name: "DelayedInvertedTurn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DelayedInvertedTurn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DelayedInvertedTurn",
				Label: "DelayedInvertedTurn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: VerticalTurn
			div := (&form.FormDiv{Name: "VerticalTurn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.VerticalTurn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "VerticalTurn",
				Label: "VerticalTurn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: InvertedVerticalTurn
			div := (&form.FormDiv{Name: "InvertedVerticalTurn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.InvertedVerticalTurn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "InvertedVerticalTurn",
				Label: "InvertedVerticalTurn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Shake
			div := (&form.FormDiv{Name: "Shake"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Shake {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Shake",
				Label: "Shake",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: WavyLine
			div := (&form.FormDiv{Name: "WavyLine"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.WavyLine {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "WavyLine",
				Label: "WavyLine",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Mordent
			div := (&form.FormDiv{Name: "Mordent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Mordent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Mordent",
				Label: "Mordent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: InvertedMordent
			div := (&form.FormDiv{Name: "InvertedMordent"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.InvertedMordent {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "InvertedMordent",
				Label: "InvertedMordent",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Schleifer
			div := (&form.FormDiv{Name: "Schleifer"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Schleifer {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Schleifer",
				Label: "Schleifer",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tremolo
			div := (&form.FormDiv{Name: "Tremolo"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tremolo {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tremolo",
				Label: "Tremolo",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Haydn
			div := (&form.FormDiv{Name: "Haydn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Haydn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Haydn",
				Label: "Haydn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OtherOrnament
			div := (&form.FormDiv{Name: "OtherOrnament"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OtherOrnament {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OtherOrnament",
				Label: "OtherOrnament",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AccidentalMark
			div := (&form.FormDiv{Name: "AccidentalMark"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AccidentalMark {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AccidentalMark",
				Label: "AccidentalMark",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Ornaments {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Ornaments", refNames, formGroup, probe.formStage)
		}
	case *note.OtherNotation:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.OtherNotation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "OtherNotation", refNames, formGroup, probe.formStage)
		}
	case *note.OtherPlacementText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.OtherArticulation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "OtherArticulation", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.OtherOrnament {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "OtherOrnament", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.OtherTechnical {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "OtherTechnical", refNames, formGroup, probe.formStage)
		}
	case *note.Pitch:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Step", inst.Step, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Alter", inst.Alter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Octave", inst.Octave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Pitch == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Pitch", refNames, formGroup, probe.formStage)
		}
	case *note.PlacementText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Bends {
				if src.WithBar == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Bend", "WithBar", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Pluck {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Pluck", refNames, formGroup, probe.formStage)
		}
	case *note.Release:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Offset", inst.Offset, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Bends {
				if src.Release == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Bend", "Release", refNames, formGroup, probe.formStage)
		}
	case *note.Rest:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Measure", inst.Measure, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("DisplayStep", inst.DisplayStep, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DisplayOctave", inst.DisplayOctave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Rest == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Rest", refNames, formGroup, probe.formStage)
		}
	case *note.Slide:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("LineType", inst.LineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DashLength", inst.DashLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SpaceLength", inst.SpaceLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FirstBeat", inst.FirstBeat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LastBeat", inst.LastBeat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Slide {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Slide", refNames, formGroup, probe.formStage)
		}
	case *note.Slur:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("LineType", inst.LineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DashLength", inst.DashLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SpaceLength", inst.SpaceLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Orientation", inst.Orientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("BezierX", inst.BezierX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierY", inst.BezierY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierX2", inst.BezierX2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierY2", inst.BezierY2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierOffset", inst.BezierOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierOffset2", inst.BezierOffset2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Slur {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Slur", refNames, formGroup, probe.formStage)
		}
	case *note.Stem:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Stem == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Stem", refNames, formGroup, probe.formStage)
		}
	case *note.StrongAccent:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.StrongAccent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "StrongAccent", refNames, formGroup, probe.formStage)
		}
	case *note.Tap:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Hand", inst.Hand, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Tap {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Tap", refNames, formGroup, probe.formStage)
		}
	case *note.Technical:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: UpBow
			div := (&form.FormDiv{Name: "UpBow"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.UpBow {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "UpBow",
				Label: "UpBow",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DownBow
			div := (&form.FormDiv{Name: "DownBow"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DownBow {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DownBow",
				Label: "DownBow",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Harmonic
			div := (&form.FormDiv{Name: "Harmonic"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Harmonic {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Harmonic",
				Label: "Harmonic",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OpenString
			div := (&form.FormDiv{Name: "OpenString"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OpenString {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OpenString",
				Label: "OpenString",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ThumbPosition
			div := (&form.FormDiv{Name: "ThumbPosition"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ThumbPosition {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ThumbPosition",
				Label: "ThumbPosition",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fingering
			div := (&form.FormDiv{Name: "Fingering"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fingering {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fingering",
				Label: "Fingering",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Pluck
			div := (&form.FormDiv{Name: "Pluck"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pluck {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pluck",
				Label: "Pluck",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DoubleTongue
			div := (&form.FormDiv{Name: "DoubleTongue"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DoubleTongue {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DoubleTongue",
				Label: "DoubleTongue",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TripleTongue
			div := (&form.FormDiv{Name: "TripleTongue"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TripleTongue {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TripleTongue",
				Label: "TripleTongue",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Stopped
			div := (&form.FormDiv{Name: "Stopped"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Stopped {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Stopped",
				Label: "Stopped",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SnapPizzicato
			div := (&form.FormDiv{Name: "SnapPizzicato"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SnapPizzicato {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SnapPizzicato",
				Label: "SnapPizzicato",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fret
			div := (&form.FormDiv{Name: "Fret"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fret {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fret",
				Label: "Fret",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: String
			div := (&form.FormDiv{Name: "String"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.String {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "String",
				Label: "String",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: HammerOn
			div := (&form.FormDiv{Name: "HammerOn"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.HammerOn {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "HammerOn",
				Label: "HammerOn",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: PullOff
			div := (&form.FormDiv{Name: "PullOff"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PullOff {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PullOff",
				Label: "PullOff",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Bend
			div := (&form.FormDiv{Name: "Bend"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bend {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bend",
				Label: "Bend",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tap
			div := (&form.FormDiv{Name: "Tap"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tap {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tap",
				Label: "Tap",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Heel
			div := (&form.FormDiv{Name: "Heel"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Heel {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Heel",
				Label: "Heel",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Toe
			div := (&form.FormDiv{Name: "Toe"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Toe {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Toe",
				Label: "Toe",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fingernails
			div := (&form.FormDiv{Name: "Fingernails"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fingernails {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fingernails",
				Label: "Fingernails",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Hole
			div := (&form.FormDiv{Name: "Hole"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Hole {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Hole",
				Label: "Hole",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Arrow
			div := (&form.FormDiv{Name: "Arrow"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Arrow {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Arrow",
				Label: "Arrow",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Handbell
			div := (&form.FormDiv{Name: "Handbell"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Handbell {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Handbell",
				Label: "Handbell",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: BrassBend
			div := (&form.FormDiv{Name: "BrassBend"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.BrassBend {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "BrassBend",
				Label: "BrassBend",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Flip
			div := (&form.FormDiv{Name: "Flip"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Flip {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Flip",
				Label: "Flip",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Smear
			div := (&form.FormDiv{Name: "Smear"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Smear {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Smear",
				Label: "Smear",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Open
			div := (&form.FormDiv{Name: "Open"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Open {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Open",
				Label: "Open",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: HalfMuted
			div := (&form.FormDiv{Name: "HalfMuted"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.HalfMuted {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "HalfMuted",
				Label: "HalfMuted",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: HarmonMute
			div := (&form.FormDiv{Name: "HarmonMute"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.HarmonMute {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "HarmonMute",
				Label: "HarmonMute",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Golpe
			div := (&form.FormDiv{Name: "Golpe"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Golpe {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Golpe",
				Label: "Golpe",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OtherTechnical
			div := (&form.FormDiv{Name: "OtherTechnical"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OtherTechnical {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OtherTechnical",
				Label: "OtherTechnical",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Technical {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Technical", refNames, formGroup, probe.formStage)
		}
	case *note.TextElementData:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineThrough", inst.LineThrough, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LetterSpacing", inst.LetterSpacing, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Lyrics {
				if src.Text == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Lyric", "Text", refNames, formGroup, probe.formStage)
		}
	case *note.Tie:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("TimeOnly", inst.TimeOnly, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				for _, target := range src.Tie {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Tie", refNames, formGroup, probe.formStage)
		}
	case *note.Tied:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("LineType", inst.LineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DashLength", inst.DashLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SpaceLength", inst.SpaceLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Orientation", inst.Orientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("BezierX", inst.BezierX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierY", inst.BezierY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierX2", inst.BezierX2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierY2", inst.BezierY2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierOffset", inst.BezierOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BezierOffset2", inst.BezierOffset2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Tied {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Tied", refNames, formGroup, probe.formStage)
		}
	case *note.Tremolo:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.Tremolo {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "Tremolo", refNames, formGroup, probe.formStage)
		}
	case *note.Tuplet:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ShowNumber", inst.ShowNumber, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ShowType", inst.ShowType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("LineShape", inst.LineShape, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TupletActual", inst.TupletActual, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.TupletPortion](), probe.formStage)
		StageSetAssociationFieldToForm("TupletNormal", inst.TupletNormal, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.TupletPortion](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Tuplet {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Tuplet", refNames, formGroup, probe.formStage)
		}
	case *note.TupletDot:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.TupletPortions {
				for _, target := range src.TupletDot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.TupletPortion", "TupletDot", refNames, formGroup, probe.formStage)
		}
	case *note.TupletNumber:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.TupletPortions {
				if src.TupletNumber == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.TupletPortion", "TupletNumber", refNames, formGroup, probe.formStage)
		}
	case *note.TupletPortion:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TupletNumber", inst.TupletNumber, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.TupletNumber](), probe.formStage)
		StageSetAssociationFieldToForm("TupletType", inst.TupletType, formGroup, probe.stageSet.NoteStage.GetInstancesSet[*note.TupletType](), probe.formStage)

		{
			// Slice of pointers: TupletDot
			div := (&form.FormDiv{Name: "TupletDot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TupletDot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TupletDot",
				Label: "TupletDot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Tuplets {
				if src.TupletActual == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Tuplet", "TupletActual", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Tuplets {
				if src.TupletNormal == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Tuplet", "TupletNormal", refNames, formGroup, probe.formStage)
		}
	case *note.TupletType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.TupletPortions {
				if src.TupletType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.TupletPortion", "TupletType", refNames, formGroup, probe.formStage)
		}
	case *note.Unpitched:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("DisplayStep", inst.DisplayStep, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DisplayOctave", inst.DisplayOctave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Unpitched == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Unpitched", refNames, formGroup, probe.formStage)
		}
	case *note.Wait:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeOnly", inst.TimeOnly, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Listens {
				for _, target := range src.Wait {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Listen", "Wait", refNames, formGroup, probe.formStage)
		}
	case *direction.Accord:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("String", inst.String, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("TuningStep", inst.TuningStep, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("TuningAlter", inst.TuningAlter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TuningOctave", inst.TuningOctave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Scordaturas {
				for _, target := range src.Accord {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Scordatura", "Accord", refNames, formGroup, probe.formStage)
		}
	case *direction.AccordionRegistration:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AccordionHigh", inst.AccordionHigh, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetBasicFieldtoForm("AccordionMiddle", inst.AccordionMiddle, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AccordionLow", inst.AccordionLow, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.AccordionRegistration == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "AccordionRegistration", refNames, formGroup, probe.formStage)
		}
	case *direction.Barre:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.FrameNotes {
				if src.Barre == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.FrameNote", "Barre", refNames, formGroup, probe.formStage)
		}
	case *direction.Bass:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Arrangement", inst.Arrangement, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("BassSeparator", inst.BassSeparator, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StyleText](), probe.formStage)
		StageSetAssociationFieldToForm("BassStep", inst.BassStep, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.BassStep](), probe.formStage)
		StageSetAssociationFieldToForm("BassAlter", inst.BassAlter, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.HarmonyAlter](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				for _, target := range src.Bass {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Bass", refNames, formGroup, probe.formStage)
		}
	case *direction.BassStep:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Basss {
				if src.BassStep == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Bass", "BassStep", refNames, formGroup, probe.formStage)
		}
	case *direction.BeatUnitTied:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("BeatUnit", inst.BeatUnit, formGroup, probe.formStage)

		{
			// Slice of pointers: BeatUnitDot
			div := (&form.FormDiv{Name: "BeatUnitDot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.BeatUnitDot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "BeatUnitDot",
				Label: "BeatUnitDot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Metronomes {
				for _, target := range src.BeatUnitTied {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Metronome", "BeatUnitTied", refNames, formGroup, probe.formStage)
		}
	case *direction.Beater:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Tip", inst.Tip, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.Beater == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "Beater", refNames, formGroup, probe.formStage)
		}
	case *direction.Bracket:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("LineEnd", inst.LineEnd, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EndLength", inst.EndLength, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("LineType", inst.LineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DashLength", inst.DashLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SpaceLength", inst.SpaceLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.Bracket == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Bracket", refNames, formGroup, probe.formStage)
		}
	case *direction.Dashes:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DashLength", inst.DashLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SpaceLength", inst.SpaceLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.Dashes == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Dashes", refNames, formGroup, probe.formStage)
		}
	case *direction.Degree:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("DegreeValue", inst.DegreeValue, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.DegreeValue](), probe.formStage)
		StageSetAssociationFieldToForm("DegreeAlter", inst.DegreeAlter, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.DegreeAlter](), probe.formStage)
		StageSetAssociationFieldToForm("DegreeType", inst.DegreeType, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.DegreeType](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				for _, target := range src.Degree {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Degree", refNames, formGroup, probe.formStage)
		}
	case *direction.DegreeAlter:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PlusMinus", inst.PlusMinus, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Degrees {
				if src.DegreeAlter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Degree", "DegreeAlter", refNames, formGroup, probe.formStage)
		}
	case *direction.DegreeType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Degrees {
				if src.DegreeType == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Degree", "DegreeType", refNames, formGroup, probe.formStage)
		}
	case *direction.DegreeValue:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Symbol", inst.Symbol, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Degrees {
				if src.DegreeValue == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Degree", "DegreeValue", refNames, formGroup, probe.formStage)
		}
	case *direction.Direction:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Directive", inst.Directive, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("System", inst.System, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: DirectionType
			div := (&form.FormDiv{Name: "DirectionType"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DirectionType {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DirectionType",
				Label: "DirectionType",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Offset](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Voice", inst.Voice, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Sound", inst.Sound, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Sound](), probe.formStage)
		StageSetAssociationFieldToForm("Listening", inst.Listening, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Listening](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Direction == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Direction", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Direction {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Direction", refNames, formGroup, probe.formStage)
		}
	case *direction.DirectionType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Rehearsal
			div := (&form.FormDiv{Name: "Rehearsal"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Rehearsal {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Rehearsal",
				Label: "Rehearsal",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Segno
			div := (&form.FormDiv{Name: "Segno"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Segno {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Segno",
				Label: "Segno",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Coda
			div := (&form.FormDiv{Name: "Coda"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Coda {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Coda",
				Label: "Coda",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Words
			div := (&form.FormDiv{Name: "Words"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Words {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Words",
				Label: "Words",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Symbol
			div := (&form.FormDiv{Name: "Symbol"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Symbol {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Symbol",
				Label: "Symbol",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Wedge", inst.Wedge, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Wedge](), probe.formStage)

		{
			// Slice of pointers: Dynamics
			div := (&form.FormDiv{Name: "Dynamics"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Dynamics {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Dynamics",
				Label: "Dynamics",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Dashes", inst.Dashes, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Dashes](), probe.formStage)
		StageSetAssociationFieldToForm("Bracket", inst.Bracket, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Bracket](), probe.formStage)
		StageSetAssociationFieldToForm("Pedal", inst.Pedal, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Pedal](), probe.formStage)
		StageSetAssociationFieldToForm("Metronome", inst.Metronome, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Metronome](), probe.formStage)
		StageSetAssociationFieldToForm("OctaveShift", inst.OctaveShift, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.OctaveShift](), probe.formStage)
		StageSetAssociationFieldToForm("HarpPedals", inst.HarpPedals, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.HarpPedals](), probe.formStage)
		StageSetAssociationFieldToForm("Damp", inst.Damp, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintStyleAlignId](), probe.formStage)
		StageSetAssociationFieldToForm("DampAll", inst.DampAll, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintStyleAlignId](), probe.formStage)
		StageSetAssociationFieldToForm("Eyeglasses", inst.Eyeglasses, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintStyleAlignId](), probe.formStage)
		StageSetAssociationFieldToForm("StringMute", inst.StringMute, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StringMute](), probe.formStage)
		StageSetAssociationFieldToForm("Scordatura", inst.Scordatura, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Scordatura](), probe.formStage)
		StageSetAssociationFieldToForm("Image", inst.Image, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Image](), probe.formStage)
		StageSetAssociationFieldToForm("PrincipalVoice", inst.PrincipalVoice, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.PrincipalVoice](), probe.formStage)

		{
			// Slice of pointers: Percussion
			div := (&form.FormDiv{Name: "Percussion"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Percussion {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Percussion",
				Label: "Percussion",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("AccordionRegistration", inst.AccordionRegistration, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.AccordionRegistration](), probe.formStage)
		StageSetAssociationFieldToForm("StaffDivide", inst.StaffDivide, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.StaffDivide](), probe.formStage)
		StageSetAssociationFieldToForm("OtherDirection", inst.OtherDirection, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.OtherDirection](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Directions {
				for _, target := range src.DirectionType {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Direction", "DirectionType", refNames, formGroup, probe.formStage)
		}
	case *direction.Effect:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.Effect == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "Effect", refNames, formGroup, probe.formStage)
		}
	case *direction.Feature:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Groupings {
				for _, target := range src.Feature {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Grouping", "Feature", refNames, formGroup, probe.formStage)
		}
	case *direction.FirstFret:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Frames {
				if src.FirstFret == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Frame", "FirstFret", refNames, formGroup, probe.formStage)
		}
	case *direction.Frame:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Unplayed", inst.Unplayed, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FrameStrings", inst.FrameStrings, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FrameFrets", inst.FrameFrets, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FirstFret", inst.FirstFret, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.FirstFret](), probe.formStage)

		{
			// Slice of pointers: FrameNote
			div := (&form.FormDiv{Name: "FrameNote"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FrameNote {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FrameNote",
				Label: "FrameNote",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				if src.Frame == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Frame", refNames, formGroup, probe.formStage)
		}
	case *direction.FrameNote:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("String", inst.String, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.StringType](), probe.formStage)
		StageSetAssociationFieldToForm("Fret", inst.Fret, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Fret](), probe.formStage)
		StageSetAssociationFieldToForm("Fingering", inst.Fingering, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Fingering](), probe.formStage)
		StageSetAssociationFieldToForm("Barre", inst.Barre, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Barre](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Frames {
				for _, target := range src.FrameNote {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Frame", "FrameNote", refNames, formGroup, probe.formStage)
		}
	case *direction.Glass:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.Glass == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "Glass", refNames, formGroup, probe.formStage)
		}
	case *direction.Grouping:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MemberOf", inst.MemberOf, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Feature
			div := (&form.FormDiv{Name: "Feature"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Feature {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Feature",
				Label: "Feature",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Grouping == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Grouping", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Grouping {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Grouping", refNames, formGroup, probe.formStage)
		}
	case *direction.Harmony:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintFrame", inst.PrintFrame, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Arrangement", inst.Arrangement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("System", inst.System, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Root
			div := (&form.FormDiv{Name: "Root"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Root {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Root",
				Label: "Root",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Numeral
			div := (&form.FormDiv{Name: "Numeral"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Numeral {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Numeral",
				Label: "Numeral",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Function
			div := (&form.FormDiv{Name: "Function"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Function {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Function",
				Label: "Function",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Kind
			div := (&form.FormDiv{Name: "Kind"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Kind {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Kind",
				Label: "Kind",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Inversion
			div := (&form.FormDiv{Name: "Inversion"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Inversion {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Inversion",
				Label: "Inversion",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Bass
			div := (&form.FormDiv{Name: "Bass"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bass {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bass",
				Label: "Bass",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Degree
			div := (&form.FormDiv{Name: "Degree"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Degree {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Degree",
				Label: "Degree",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Frame", inst.Frame, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Frame](), probe.formStage)
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Offset](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Harmony == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Harmony", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Harmony {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Harmony", refNames, formGroup, probe.formStage)
		}
	case *direction.HarmonyAlter:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Basss {
				if src.BassAlter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Bass", "BassAlter", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Numerals {
				if src.NumeralAlter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Numeral", "NumeralAlter", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Roots {
				if src.RootAlter == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Root", "RootAlter", refNames, formGroup, probe.formStage)
		}
	case *direction.HarpPedals:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: PedalTuning
			div := (&form.FormDiv{Name: "PedalTuning"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PedalTuning {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PedalTuning",
				Label: "PedalTuning",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.HarpPedals == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "HarpPedals", refNames, formGroup, probe.formStage)
		}
	case *direction.Image:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Source", inst.Source, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Credits {
				if src.CreditImage == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Credit", "CreditImage", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.Image == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Image", refNames, formGroup, probe.formStage)
		}
	case *direction.InstrumentChange:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InstrumentSound", inst.InstrumentSound, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Solo", inst.Solo, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetBasicFieldtoForm("Ensemble", inst.Ensemble, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("VirtualInstrument", inst.VirtualInstrument, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.VirtualInstrument](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Sounds {
				for _, target := range src.InstrumentChange {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Sound", "InstrumentChange", refNames, formGroup, probe.formStage)
		}
	case *direction.Inversion:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				for _, target := range src.Inversion {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Inversion", refNames, formGroup, probe.formStage)
		}
	case *direction.Kind:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("UseSymbols", inst.UseSymbols, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StackDegrees", inst.StackDegrees, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ParenthesesDegrees", inst.ParenthesesDegrees, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("BracketDegrees", inst.BracketDegrees, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				for _, target := range src.Kind {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Kind", refNames, formGroup, probe.formStage)
		}
	case *direction.Listening:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Sync
			div := (&form.FormDiv{Name: "Sync"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sync {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sync",
				Label: "Sync",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OtherListening
			div := (&form.FormDiv{Name: "OtherListening"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OtherListening {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OtherListening",
				Label: "OtherListening",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Offset](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Listening == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Listening", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Listening {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Listening", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Directions {
				if src.Listening == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Direction", "Listening", refNames, formGroup, probe.formStage)
		}
	case *direction.MeasureNumbering:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("System", inst.System, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Staff", inst.Staff, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("MultipleRestAlways", inst.MultipleRestAlways, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MultipleRestRange", inst.MultipleRestRange, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Prints {
				if src.MeasureNumbering == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Print", "MeasureNumbering", refNames, formGroup, probe.formStage)
		}
	case *direction.Membrane:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.Membrane == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "Membrane", refNames, formGroup, probe.formStage)
		}
	case *direction.Metal:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.Metal == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "Metal", refNames, formGroup, probe.formStage)
		}
	case *direction.Metronome:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: BeatUnitDot
			div := (&form.FormDiv{Name: "BeatUnitDot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.BeatUnitDot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "BeatUnitDot",
				Label: "BeatUnitDot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: BeatUnitTied
			div := (&form.FormDiv{Name: "BeatUnitTied"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.BeatUnitTied {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "BeatUnitTied",
				Label: "BeatUnitTied",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("PerMinute", inst.PerMinute, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.PerMinute](), probe.formStage)
		StageSetAssociationFieldToForm("MetronomeArrows", inst.MetronomeArrows, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)

		{
			// Slice of pointers: MetronomeNote
			div := (&form.FormDiv{Name: "MetronomeNote"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MetronomeNote {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MetronomeNote",
				Label: "MetronomeNote",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("MetronomeRelation", inst.MetronomeRelation, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.Metronome == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Metronome", refNames, formGroup, probe.formStage)
		}
	case *direction.MetronomeBeam:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.MetronomeNotes {
				for _, target := range src.MetronomeBeam {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.MetronomeNote", "MetronomeBeam", refNames, formGroup, probe.formStage)
		}
	case *direction.MetronomeNote:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("MetronomeType", inst.MetronomeType, formGroup, probe.formStage)

		{
			// Slice of pointers: MetronomeDot
			div := (&form.FormDiv{Name: "MetronomeDot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MetronomeDot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MetronomeDot",
				Label: "MetronomeDot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: MetronomeBeam
			div := (&form.FormDiv{Name: "MetronomeBeam"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MetronomeBeam {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MetronomeBeam",
				Label: "MetronomeBeam",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("MetronomeTied", inst.MetronomeTied, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.MetronomeTied](), probe.formStage)
		StageSetAssociationFieldToForm("MetronomeTuplet", inst.MetronomeTuplet, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.MetronomeTuplet](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Metronomes {
				for _, target := range src.MetronomeNote {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Metronome", "MetronomeNote", refNames, formGroup, probe.formStage)
		}
	case *direction.MetronomeTied:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.MetronomeNotes {
				if src.MetronomeTied == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.MetronomeNote", "MetronomeTied", refNames, formGroup, probe.formStage)
		}
	case *direction.MetronomeTuplet:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("ShowNumber", inst.ShowNumber, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ActualNotes", inst.ActualNotes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NormalNotes", inst.NormalNotes, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("NormalType", inst.NormalType, formGroup, probe.formStage)

		{
			// Slice of pointers: NormalDot
			div := (&form.FormDiv{Name: "NormalDot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NormalDot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NormalDot",
				Label: "NormalDot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.MetronomeNotes {
				if src.MetronomeTuplet == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.MetronomeNote", "MetronomeTuplet", refNames, formGroup, probe.formStage)
		}
	case *direction.Numeral:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("NumeralRoot", inst.NumeralRoot, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.NumeralRoot](), probe.formStage)
		StageSetAssociationFieldToForm("NumeralAlter", inst.NumeralAlter, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.HarmonyAlter](), probe.formStage)
		StageSetAssociationFieldToForm("NumeralKey", inst.NumeralKey, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.NumeralKey](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				for _, target := range src.Numeral {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Numeral", refNames, formGroup, probe.formStage)
		}
	case *direction.NumeralKey:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("NumeralFifths", inst.NumeralFifths, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("NumeralMode", inst.NumeralMode, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Numerals {
				if src.NumeralKey == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Numeral", "NumeralKey", refNames, formGroup, probe.formStage)
		}
	case *direction.NumeralRoot:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Numerals {
				if src.NumeralRoot == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Numeral", "NumeralRoot", refNames, formGroup, probe.formStage)
		}
	case *direction.OctaveShift:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Size", inst.Size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DashLength", inst.DashLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SpaceLength", inst.SpaceLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.OctaveShift == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "OctaveShift", refNames, formGroup, probe.formStage)
		}
	case *direction.Offset:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Sound", inst.Sound, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Directions {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Direction", "Offset", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Offset", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Listenings {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Listening", "Offset", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Sounds {
				if src.Offset == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Sound", "Offset", refNames, formGroup, probe.formStage)
		}
	case *direction.OtherDirection:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.OtherDirection == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "OtherDirection", refNames, formGroup, probe.formStage)
		}
	case *direction.OtherListening:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeOnly", inst.TimeOnly, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Listens {
				for _, target := range src.OtherListen {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Listen", "OtherListen", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Listenings {
				for _, target := range src.OtherListening {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Listening", "OtherListening", refNames, formGroup, probe.formStage)
		}
	case *direction.Pedal:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Line", inst.Line, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Sign", inst.Sign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Abbreviated", inst.Abbreviated, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.Pedal == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Pedal", refNames, formGroup, probe.formStage)
		}
	case *direction.PedalTuning:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PedalStep", inst.PedalStep, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("PedalAlter", inst.PedalAlter, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.HarpPedalss {
				for _, target := range src.PedalTuning {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.HarpPedals", "PedalTuning", refNames, formGroup, probe.formStage)
		}
	case *direction.PerMinute:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Metronomes {
				if src.PerMinute == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Metronome", "PerMinute", refNames, formGroup, probe.formStage)
		}
	case *direction.Percussion:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Glass", inst.Glass, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Glass](), probe.formStage)
		StageSetAssociationFieldToForm("Metal", inst.Metal, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Metal](), probe.formStage)
		StageSetAssociationFieldToForm("Wood", inst.Wood, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Wood](), probe.formStage)
		StageSetAssociationFieldToForm("Pitched", inst.Pitched, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Pitched](), probe.formStage)
		StageSetAssociationFieldToForm("Membrane", inst.Membrane, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Membrane](), probe.formStage)
		StageSetAssociationFieldToForm("Effect", inst.Effect, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Effect](), probe.formStage)
		StageSetAssociationFieldToForm("Timpani", inst.Timpani, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Timpani](), probe.formStage)
		StageSetAssociationFieldToForm("Beater", inst.Beater, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Beater](), probe.formStage)
		StageSetAssociationFieldToForm("Stick", inst.Stick, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Stick](), probe.formStage)
		StageSetEnumStringFieldToForm("StickLocation", inst.StickLocation, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("OtherPercussion", inst.OtherPercussion, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.OtherText](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				for _, target := range src.Percussion {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Percussion", refNames, formGroup, probe.formStage)
		}
	case *direction.Pitched:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.Pitched == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "Pitched", refNames, formGroup, probe.formStage)
		}
	case *direction.PrincipalVoice:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Symbol", inst.Symbol, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.PrincipalVoice == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "PrincipalVoice", refNames, formGroup, probe.formStage)
		}
	case *direction.Print:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StaffSpacing", inst.StaffSpacing, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("NewSystem", inst.NewSystem, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("NewPage", inst.NewPage, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("BlankPage", inst.BlankPage, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PageNumber", inst.PageNumber, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PageLayout", inst.PageLayout, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.PageLayout](), probe.formStage)
		StageSetAssociationFieldToForm("SystemLayout", inst.SystemLayout, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.SystemLayout](), probe.formStage)

		{
			// Slice of pointers: StaffLayout
			div := (&form.FormDiv{Name: "StaffLayout"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.StaffLayout {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "StaffLayout",
				Label: "StaffLayout",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("MeasureLayout", inst.MeasureLayout, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.MeasureLayout](), probe.formStage)
		StageSetAssociationFieldToForm("MeasureNumbering", inst.MeasureNumbering, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.MeasureNumbering](), probe.formStage)
		StageSetAssociationFieldToForm("PartNameDisplay", inst.PartNameDisplay, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), probe.formStage)
		StageSetAssociationFieldToForm("PartAbbreviationDisplay", inst.PartAbbreviationDisplay, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.NameDisplay](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Print == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Print", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Print {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Print", refNames, formGroup, probe.formStage)
		}
	case *direction.Root:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("RootStep", inst.RootStep, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.RootStep](), probe.formStage)
		StageSetAssociationFieldToForm("RootAlter", inst.RootAlter, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.HarmonyAlter](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				for _, target := range src.Root {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Root", refNames, formGroup, probe.formStage)
		}
	case *direction.RootStep:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Text", inst.Text, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Roots {
				if src.RootStep == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Root", "RootStep", refNames, formGroup, probe.formStage)
		}
	case *direction.Scordatura:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: Accord
			div := (&form.FormDiv{Name: "Accord"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Accord {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Accord",
				Label: "Accord",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.Scordatura == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Scordatura", refNames, formGroup, probe.formStage)
		}
	case *direction.Sound:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tempo", inst.Tempo, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dynamics", inst.Dynamics, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dacapo", inst.Dacapo, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Segno", inst.Segno, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Dalsegno", inst.Dalsegno, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Coda", inst.Coda, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tocoda", inst.Tocoda, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Divisions", inst.Divisions, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ForwardRepeat", inst.ForwardRepeat, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Fine", inst.Fine, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeOnly", inst.TimeOnly, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Pizzicato", inst.Pizzicato, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Pan", inst.Pan, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Elevation", inst.Elevation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DamperPedal", inst.DamperPedal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SoftPedal", inst.SoftPedal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SostenutoPedal", inst.SostenutoPedal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: InstrumentChange
			div := (&form.FormDiv{Name: "InstrumentChange"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.InstrumentChange {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "InstrumentChange",
				Label: "InstrumentChange",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: MidiDevice
			div := (&form.FormDiv{Name: "MidiDevice"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MidiDevice {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MidiDevice",
				Label: "MidiDevice",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: MidiInstrument
			div := (&form.FormDiv{Name: "MidiInstrument"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MidiInstrument {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MidiInstrument",
				Label: "MidiInstrument",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Play
			div := (&form.FormDiv{Name: "Play"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Play {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Play",
				Label: "Play",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Swing", inst.Swing, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Swing](), probe.formStage)
		StageSetAssociationFieldToForm("Offset", inst.Offset, formGroup, probe.stageSet.DirectionStage.GetInstancesSet[*direction.Offset](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Sound == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Sound", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Sound {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Sound", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Directions {
				if src.Sound == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Direction", "Sound", refNames, formGroup, probe.formStage)
		}
	case *direction.StaffDivide:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.StaffDivide == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "StaffDivide", refNames, formGroup, probe.formStage)
		}
	case *direction.Stick:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Tip", inst.Tip, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("DashedCircle", inst.DashedCircle, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("StickType", inst.StickType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("StickMaterial", inst.StickMaterial, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.Stick == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "Stick", refNames, formGroup, probe.formStage)
		}
	case *direction.StringMute:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.StringMute == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "StringMute", refNames, formGroup, probe.formStage)
		}
	case *direction.StyleText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Figures {
				if src.Prefix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Figure", "Prefix", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Figures {
				if src.FigureNumber == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Figure", "FigureNumber", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Figures {
				if src.Suffix == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Figure", "Suffix", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Basss {
				if src.BassSeparator == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Bass", "BassSeparator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				for _, target := range src.Function {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Function", refNames, formGroup, probe.formStage)
		}
	case *direction.Swing:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Straight", inst.Straight, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Empty](), probe.formStage)
		StageSetBasicFieldtoForm("First", inst.First, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Second", inst.Second, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("SwingType", inst.SwingType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("SwingStyle", inst.SwingStyle, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Sounds {
				if src.Swing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Sound", "Swing", refNames, formGroup, probe.formStage)
		}
	case *direction.Sync:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Latency", inst.Latency, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Player", inst.Player, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeOnly", inst.TimeOnly, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Listenings {
				for _, target := range src.Sync {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Listening", "Sync", refNames, formGroup, probe.formStage)
		}
	case *direction.TimeModification:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ActualNotes", inst.ActualNotes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NormalNotes", inst.NormalNotes, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("NormalType", inst.NormalType, formGroup, probe.formStage)

		{
			// Slice of pointers: NormalDot
			div := (&form.FormDiv{Name: "NormalDot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NormalDot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NormalDot",
				Label: "NormalDot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.TimeModification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "TimeModification", refNames, formGroup, probe.formStage)
		}
	case *direction.Timpani:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.Timpani == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "Timpani", refNames, formGroup, probe.formStage)
		}
	case *direction.VirtualInstrument:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("VirtualLibrary", inst.VirtualLibrary, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("VirtualName", inst.VirtualName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreInstruments {
				if src.VirtualInstrument == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreInstrument", "VirtualInstrument", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.InstrumentChanges {
				if src.VirtualInstrument == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.InstrumentChange", "VirtualInstrument", refNames, formGroup, probe.formStage)
		}
	case *direction.Wedge:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Spread", inst.Spread, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Niente", inst.Niente, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("LineType", inst.LineType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DashLength", inst.DashLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SpaceLength", inst.SpaceLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.Wedge == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Wedge", refNames, formGroup, probe.formStage)
		}
	case *direction.Wood:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.Wood == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "Wood", refNames, formGroup, probe.formStage)
		}
	case *barline.BarStyleColor:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.BarlineStage.Barlines {
				if src.BarStyle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("barline.Barline", "BarStyle", refNames, formGroup, probe.formStage)
		}
	case *barline.Barline:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Segno", inst.Segno, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Coda", inst.Coda, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Divisions", inst.Divisions, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("BarStyle", inst.BarStyle, formGroup, probe.stageSet.BarlineStage.GetInstancesSet[*barline.BarStyleColor](), probe.formStage)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)
		StageSetAssociationFieldToForm("WavyLine", inst.WavyLine, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.WavyLine](), probe.formStage)
		StageSetAssociationFieldToForm("Segno_1", inst.Segno_1, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Segno](), probe.formStage)
		StageSetAssociationFieldToForm("Coda_1", inst.Coda_1, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Coda](), probe.formStage)

		{
			// Slice of pointers: Fermata
			div := (&form.FormDiv{Name: "Fermata"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fermata {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fermata",
				Label: "Fermata",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Ending", inst.Ending, formGroup, probe.stageSet.BarlineStage.GetInstancesSet[*barline.Ending](), probe.formStage)
		StageSetAssociationFieldToForm("Repeat", inst.Repeat, formGroup, probe.stageSet.BarlineStage.GetInstancesSet[*barline.Repeat](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Barline == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Barline", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Barline {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Barline", refNames, formGroup, probe.formStage)
		}
	case *barline.Ending:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EndLength", inst.EndLength, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TextX", inst.TextX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TextY", inst.TextY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("System", inst.System, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.BarlineStage.Barlines {
				if src.Ending == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("barline.Barline", "Ending", refNames, formGroup, probe.formStage)
		}
	case *barline.Repeat:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Direction", inst.Direction, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Times", inst.Times, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("AfterJump", inst.AfterJump, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Winged", inst.Winged, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.BarlineStage.Barlines {
				if src.Repeat == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("barline.Barline", "Repeat", refNames, formGroup, probe.formStage)
		}
	case *attributes.Attributes:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Footnote", inst.Footnote, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.FormattedText](), probe.formStage)
		StageSetAssociationFieldToForm("Level", inst.Level, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Level](), probe.formStage)
		StageSetBasicFieldtoForm("Divisions", inst.Divisions, probe.formStage, formGroup)

		{
			// Slice of pointers: Key
			div := (&form.FormDiv{Name: "Key"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Key {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Key",
				Label: "Key",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Time
			div := (&form.FormDiv{Name: "Time"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Time {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Time",
				Label: "Time",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Staves", inst.Staves, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PartSymbol", inst.PartSymbol, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.PartSymbol](), probe.formStage)
		StageSetBasicFieldtoForm("Instruments", inst.Instruments, probe.formStage, formGroup)

		{
			// Slice of pointers: Clef
			div := (&form.FormDiv{Name: "Clef"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Clef {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Clef",
				Label: "Clef",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: StaffDetails
			div := (&form.FormDiv{Name: "StaffDetails"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.StaffDetails {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "StaffDetails",
				Label: "StaffDetails",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Transpose
			div := (&form.FormDiv{Name: "Transpose"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Transpose {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Transpose",
				Label: "Transpose",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ForPart
			div := (&form.FormDiv{Name: "ForPart"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ForPart {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ForPart",
				Label: "ForPart",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Directive
			div := (&form.FormDiv{Name: "Directive"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Directive {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Directive",
				Label: "Directive",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: MeasureStyle
			div := (&form.FormDiv{Name: "MeasureStyle"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MeasureStyle {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MeasureStyle",
				Label: "MeasureStyle",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Attributes == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Attributes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Attributes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Attributes", refNames, formGroup, probe.formStage)
		}
	case *attributes.AttributesDirective:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				for _, target := range src.Directive {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "Directive", refNames, formGroup, probe.formStage)
		}
	case *attributes.BeatRepeat:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slashes", inst.Slashes, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("UseDots", inst.UseDots, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("SlashType", inst.SlashType, formGroup, probe.formStage)

		{
			// Slice of pointers: SlashDot
			div := (&form.FormDiv{Name: "SlashDot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SlashDot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SlashDot",
				Label: "SlashDot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ExceptVoice", inst.ExceptVoice, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.MeasureStyles {
				if src.BeatRepeat == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.MeasureStyle", "BeatRepeat", refNames, formGroup, probe.formStage)
		}
	case *attributes.Cancel:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Location", inst.Location, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Keys {
				if src.Cancel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Key", "Cancel", refNames, formGroup, probe.formStage)
		}
	case *attributes.Clef:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Additional", inst.Additional, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("AfterBarline", inst.AfterBarline, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Sign", inst.Sign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ClefOctaveChange", inst.ClefOctaveChange, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				for _, target := range src.Clef {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "Clef", refNames, formGroup, probe.formStage)
		}
	case *attributes.Double:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Above", inst.Above, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.PartTransposes {
				if src.Double == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.PartTranspose", "Double", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Transposes {
				if src.Double == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Transpose", "Double", refNames, formGroup, probe.formStage)
		}
	case *attributes.ForPart:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PartClef", inst.PartClef, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.PartClef](), probe.formStage)
		StageSetAssociationFieldToForm("PartTranspose", inst.PartTranspose, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.PartTranspose](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				for _, target := range src.ForPart {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "ForPart", refNames, formGroup, probe.formStage)
		}
	case *attributes.Interchangeable:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Symbol", inst.Symbol, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Separator", inst.Separator, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TimeRelation", inst.TimeRelation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BeatType", inst.BeatType, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Times {
				if src.Interchangeable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Time", "Interchangeable", refNames, formGroup, probe.formStage)
		}
	case *attributes.Key:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Cancel", inst.Cancel, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Cancel](), probe.formStage)
		StageSetBasicFieldtoForm("Fifths", inst.Fifths, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Mode", inst.Mode, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("KeyStep", inst.KeyStep, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("KeyAlter", inst.KeyAlter, probe.formStage, formGroup)

		{
			// Slice of pointers: KeyAccidental
			div := (&form.FormDiv{Name: "KeyAccidental"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.KeyAccidental {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "KeyAccidental",
				Label: "KeyAccidental",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: KeyOctave
			div := (&form.FormDiv{Name: "KeyOctave"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.KeyOctave {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "KeyOctave",
				Label: "KeyOctave",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				for _, target := range src.Key {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "Key", refNames, formGroup, probe.formStage)
		}
	case *attributes.KeyAccidental:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Keys {
				for _, target := range src.KeyAccidental {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Key", "KeyAccidental", refNames, formGroup, probe.formStage)
		}
	case *attributes.KeyOctave:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Cancel", inst.Cancel, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Keys {
				for _, target := range src.KeyOctave {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Key", "KeyOctave", refNames, formGroup, probe.formStage)
		}
	case *attributes.LineDetail:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("LineType", inst.LineType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.StaffDetailss {
				for _, target := range src.LineDetail {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.StaffDetails", "LineDetail", refNames, formGroup, probe.formStage)
		}
	case *attributes.MeasureRepeat:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Slashes", inst.Slashes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.MeasureStyles {
				if src.MeasureRepeat == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.MeasureStyle", "MeasureRepeat", refNames, formGroup, probe.formStage)
		}
	case *attributes.MeasureStyle:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("MultipleRest", inst.MultipleRest, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.MultipleRest](), probe.formStage)
		StageSetAssociationFieldToForm("MeasureRepeat", inst.MeasureRepeat, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.MeasureRepeat](), probe.formStage)
		StageSetAssociationFieldToForm("BeatRepeat", inst.BeatRepeat, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.BeatRepeat](), probe.formStage)
		StageSetAssociationFieldToForm("Slash", inst.Slash, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Slash](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				for _, target := range src.MeasureStyle {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "MeasureStyle", refNames, formGroup, probe.formStage)
		}
	case *attributes.MultipleRest:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("UseSymbols", inst.UseSymbols, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.MeasureStyles {
				if src.MultipleRest == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.MeasureStyle", "MultipleRest", refNames, formGroup, probe.formStage)
		}
	case *attributes.PartClef:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Sign", inst.Sign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ClefOctaveChange", inst.ClefOctaveChange, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.ForParts {
				if src.PartClef == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.ForPart", "PartClef", refNames, formGroup, probe.formStage)
		}
	case *attributes.PartSymbol:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TopStaff", inst.TopStaff, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomStaff", inst.BottomStaff, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				if src.PartSymbol == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "PartSymbol", refNames, formGroup, probe.formStage)
		}
	case *attributes.PartTranspose:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Diatonic", inst.Diatonic, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Chromatic", inst.Chromatic, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OctaveChange", inst.OctaveChange, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Double", inst.Double, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Double](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.ForParts {
				if src.PartTranspose == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.ForPart", "PartTranspose", refNames, formGroup, probe.formStage)
		}
	case *attributes.Slash:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("UseDots", inst.UseDots, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("UseStems", inst.UseStems, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("SlashType", inst.SlashType, formGroup, probe.formStage)

		{
			// Slice of pointers: SlashDot
			div := (&form.FormDiv{Name: "SlashDot"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SlashDot {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SlashDot",
				Label: "SlashDot",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ExceptVoice", inst.ExceptVoice, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.MeasureStyles {
				if src.Slash == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.MeasureStyle", "Slash", refNames, formGroup, probe.formStage)
		}
	case *attributes.StaffDetails:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ShowFrets", inst.ShowFrets, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintSpacing", inst.PrintSpacing, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("StaffType", inst.StaffType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("StaffLines", inst.StaffLines, probe.formStage, formGroup)

		{
			// Slice of pointers: LineDetail
			div := (&form.FormDiv{Name: "LineDetail"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.LineDetail {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "LineDetail",
				Label: "LineDetail",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: StaffTuning
			div := (&form.FormDiv{Name: "StaffTuning"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.StaffTuning {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "StaffTuning",
				Label: "StaffTuning",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Capo", inst.Capo, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("StaffSize", inst.StaffSize, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.StaffSize](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				for _, target := range src.StaffDetails {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "StaffDetails", refNames, formGroup, probe.formStage)
		}
	case *attributes.StaffSize:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Scaling", inst.Scaling, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.StaffDetailss {
				if src.StaffSize == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.StaffDetails", "StaffSize", refNames, formGroup, probe.formStage)
		}
	case *attributes.StaffTuning:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Line", inst.Line, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("TuningStep", inst.TuningStep, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("TuningAlter", inst.TuningAlter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TuningOctave", inst.TuningOctave, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.StaffDetailss {
				for _, target := range src.StaffTuning {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.StaffDetails", "StaffTuning", refNames, formGroup, probe.formStage)
		}
	case *attributes.Time:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Symbol", inst.Symbol, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Separator", inst.Separator, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BeatType", inst.BeatType, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Interchangeable", inst.Interchangeable, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Interchangeable](), probe.formStage)
		StageSetBasicFieldtoForm("SenzaMisura", inst.SenzaMisura, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				for _, target := range src.Time {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "Time", refNames, formGroup, probe.formStage)
		}
	case *attributes.Transpose:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Diatonic", inst.Diatonic, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Chromatic", inst.Chromatic, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OctaveChange", inst.OctaveChange, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Double", inst.Double, formGroup, probe.stageSet.AttributesStage.GetInstancesSet[*attributes.Double](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				for _, target := range src.Transpose {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "Transpose", refNames, formGroup, probe.formStage)
		}
	case *common.AccidentalText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space", inst.Space, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineThrough", inst.LineThrough, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LetterSpacing", inst.LetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineHeight", inst.LineHeight, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.NoteheadTexts {
				for _, target := range src.AccidentalText {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.NoteheadText", "AccidentalText", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.NameDisplays {
				for _, target := range src.AccidentalText {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.NameDisplay", "AccidentalText", refNames, formGroup, probe.formStage)
		}
	case *common.Appearance:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: LineWidth
			div := (&form.FormDiv{Name: "LineWidth"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.LineWidth {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "LineWidth",
				Label: "LineWidth",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: NoteSize
			div := (&form.FormDiv{Name: "NoteSize"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NoteSize {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NoteSize",
				Label: "NoteSize",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Distance
			div := (&form.FormDiv{Name: "Distance"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Distance {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Distance",
				Label: "Distance",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Glyph
			div := (&form.FormDiv{Name: "Glyph"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Glyph {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Glyph",
				Label: "Glyph",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OtherAppearance
			div := (&form.FormDiv{Name: "OtherAppearance"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OtherAppearance {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OtherAppearance",
				Label: "OtherAppearance",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				if src.Appearance == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "Appearance", refNames, formGroup, probe.formStage)
		}
	case *common.Bookmark:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameAttr", inst.NameAttr, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Element", inst.Element, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Position", inst.Position, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Credits {
				for _, target := range src.Bookmark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Credit", "Bookmark", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Bookmark == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Bookmark", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Bookmark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Bookmark", refNames, formGroup, probe.formStage)
		}
	case *common.Coda:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				for _, target := range src.Coda {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Coda", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.BarlineStage.Barlines {
				if src.Coda_1 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("barline.Barline", "Coda_1", refNames, formGroup, probe.formStage)
		}
	case *common.Distance:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Appearances {
				for _, target := range src.Distance {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Appearance", "Distance", refNames, formGroup, probe.formStage)
		}
	case *common.Dynamics:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineThrough", inst.LineThrough, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			// Slice of pointers: P
			div := (&form.FormDiv{Name: "P"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.P {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "P",
				Label: "P",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Pp
			div := (&form.FormDiv{Name: "Pp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pp",
				Label: "Pp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Ppp
			div := (&form.FormDiv{Name: "Ppp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Ppp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Ppp",
				Label: "Ppp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Pppp
			div := (&form.FormDiv{Name: "Pppp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pppp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pppp",
				Label: "Pppp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Ppppp
			div := (&form.FormDiv{Name: "Ppppp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Ppppp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Ppppp",
				Label: "Ppppp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Pppppp
			div := (&form.FormDiv{Name: "Pppppp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pppppp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pppppp",
				Label: "Pppppp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: F
			div := (&form.FormDiv{Name: "F"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.F {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "F",
				Label: "F",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Ff
			div := (&form.FormDiv{Name: "Ff"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Ff {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Ff",
				Label: "Ff",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fff
			div := (&form.FormDiv{Name: "Fff"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fff {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fff",
				Label: "Fff",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Ffff
			div := (&form.FormDiv{Name: "Ffff"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Ffff {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Ffff",
				Label: "Ffff",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fffff
			div := (&form.FormDiv{Name: "Fffff"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fffff {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fffff",
				Label: "Fffff",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Ffffff
			div := (&form.FormDiv{Name: "Ffffff"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Ffffff {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Ffffff",
				Label: "Ffffff",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Mp
			div := (&form.FormDiv{Name: "Mp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Mp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Mp",
				Label: "Mp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Mf
			div := (&form.FormDiv{Name: "Mf"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Mf {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Mf",
				Label: "Mf",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Sf
			div := (&form.FormDiv{Name: "Sf"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sf {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sf",
				Label: "Sf",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Sfp
			div := (&form.FormDiv{Name: "Sfp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sfp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sfp",
				Label: "Sfp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Sfpp
			div := (&form.FormDiv{Name: "Sfpp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sfpp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sfpp",
				Label: "Sfpp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fp
			div := (&form.FormDiv{Name: "Fp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fp",
				Label: "Fp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Rf
			div := (&form.FormDiv{Name: "Rf"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Rf {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Rf",
				Label: "Rf",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Rfz
			div := (&form.FormDiv{Name: "Rfz"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Rfz {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Rfz",
				Label: "Rfz",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Sfz
			div := (&form.FormDiv{Name: "Sfz"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sfz {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sfz",
				Label: "Sfz",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Sffz
			div := (&form.FormDiv{Name: "Sffz"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sffz {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sffz",
				Label: "Sffz",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Fz
			div := (&form.FormDiv{Name: "Fz"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fz {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fz",
				Label: "Fz",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: N
			div := (&form.FormDiv{Name: "N"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.N {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "N",
				Label: "N",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Pf
			div := (&form.FormDiv{Name: "Pf"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pf {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pf",
				Label: "Pf",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Sfzp
			div := (&form.FormDiv{Name: "Sfzp"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sfzp {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sfzp",
				Label: "Sfzp",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: OtherDynamics
			div := (&form.FormDiv{Name: "OtherDynamics"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OtherDynamics {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OtherDynamics",
				Label: "OtherDynamics",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Dynamics {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Dynamics", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				for _, target := range src.Dynamics {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Dynamics", refNames, formGroup, probe.formStage)
		}
	case *common.Empty:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				if src.ConcertScore == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "ConcertScore", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartGroups {
				if src.GroupTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartGroup", "GroupTime", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreInstruments {
				if src.Solo == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreInstrument", "Solo", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Arrows {
				if src.Arrowhead == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Arrow", "Arrowhead", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Bends {
				if src.PreBend == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Bend", "PreBend", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Harmonics {
				if src.Natural == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Harmonic", "Natural", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Harmonics {
				if src.Artificial == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Harmonic", "Artificial", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Harmonics {
				if src.BasePitch == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Harmonic", "BasePitch", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Harmonics {
				if src.TouchingPitch == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Harmonic", "TouchingPitch", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Harmonics {
				if src.SoundingPitch == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Harmonic", "SoundingPitch", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Lyrics {
				if src.Laughing == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Lyric", "Laughing", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Lyrics {
				if src.Humming == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Lyric", "Humming", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Lyrics {
				if src.EndLine == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Lyric", "EndLine", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Lyrics {
				if src.EndParagraph == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Lyric", "EndParagraph", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Chord == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Chord", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Cue == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Cue", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.AccordionRegistrations {
				if src.AccordionHigh == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.AccordionRegistration", "AccordionHigh", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.AccordionRegistrations {
				if src.AccordionLow == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.AccordionRegistration", "AccordionLow", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.BeatUnitTieds {
				for _, target := range src.BeatUnitDot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.BeatUnitTied", "BeatUnitDot", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.InstrumentChanges {
				if src.Solo == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.InstrumentChange", "Solo", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Metronomes {
				for _, target := range src.BeatUnitDot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Metronome", "BeatUnitDot", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Metronomes {
				if src.MetronomeArrows == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Metronome", "MetronomeArrows", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.MetronomeNotes {
				for _, target := range src.MetronomeDot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.MetronomeNote", "MetronomeDot", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.MetronomeTuplets {
				for _, target := range src.NormalDot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.MetronomeTuplet", "NormalDot", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Swings {
				if src.Straight == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Swing", "Straight", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.TimeModifications {
				for _, target := range src.NormalDot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.TimeModification", "NormalDot", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.BeatRepeats {
				for _, target := range src.SlashDot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.BeatRepeat", "SlashDot", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Slashs {
				for _, target := range src.SlashDot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Slash", "SlashDot", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.P {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "P", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Pp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Pp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Ppp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Ppp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Pppp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Pppp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Ppppp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Ppppp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Pppppp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Pppppp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.F {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "F", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Ff {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Ff", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Fff {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Fff", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Ffff {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Ffff", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Fffff {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Fffff", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Ffffff {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Ffffff", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Mp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Mp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Mf {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Mf", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Sf {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Sf", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Sfp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Sfp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Sfpp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Sfpp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Fp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Fp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Rf {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Rf", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Rfz {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Rfz", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Sfz {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Sfz", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Sffz {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Sffz", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Fz {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Fz", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.N {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "N", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Pf {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Pf", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.Sfzp {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "Sfzp", refNames, formGroup, probe.formStage)
		}
	case *common.EmptyPlacement:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Accent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Accent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Staccato {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Staccato", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Tenuto {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Tenuto", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.DetachedLegato {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "DetachedLegato", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Staccatissimo {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Staccatissimo", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Spiccato {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Spiccato", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Stress {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Stress", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.Unstress {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "Unstress", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Articulationss {
				for _, target := range src.SoftAccent {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Articulations", "SoftAccent", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				for _, target := range src.Dot {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Dot", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.Schleifer {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "Schleifer", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.UpBow {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "UpBow", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.DownBow {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "DownBow", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.OpenString {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "OpenString", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.ThumbPosition {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "ThumbPosition", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.DoubleTongue {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "DoubleTongue", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.TripleTongue {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "TripleTongue", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.SnapPizzicato {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "SnapPizzicato", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Fingernails {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Fingernails", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.BrassBend {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "BrassBend", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Flip {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Flip", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Smear {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Smear", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Golpe {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Golpe", refNames, formGroup, probe.formStage)
		}
	case *common.EmptyPlacementSmufl:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Stopped {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Stopped", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Open {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Open", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.HalfMuted {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "HalfMuted", refNames, formGroup, probe.formStage)
		}
	case *common.EmptyPrintObjectStyleAlign:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.SystemDividerss {
				if src.LeftDivider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.SystemDividers", "LeftDivider", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.SystemDividerss {
				if src.RightDivider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.SystemDividers", "RightDivider", refNames, formGroup, probe.formStage)
		}
	case *common.EmptyPrintStyle:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
	case *common.EmptyPrintStyleAlign:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
	case *common.EmptyPrintStyleAlignId:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.Damp == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Damp", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.DampAll == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "DampAll", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				if src.Eyeglasses == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Eyeglasses", refNames, formGroup, probe.formStage)
		}
	case *common.EmptyTrillSound:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("StartNote", inst.StartNote, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TrillStep", inst.TrillStep, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TwoNoteTurn", inst.TwoNoteTurn, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SecondBeat", inst.SecondBeat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LastBeat", inst.LastBeat, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.TrillMark {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "TrillMark", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.VerticalTurn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "VerticalTurn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.InvertedVerticalTurn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "InvertedVerticalTurn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.Shake {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "Shake", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.Haydn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "Haydn", refNames, formGroup, probe.formStage)
		}
	case *common.Encoding:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EncodingDate", inst.EncodingDate, probe.formStage, formGroup)

		{
			// Slice of pointers: Encoder
			div := (&form.FormDiv{Name: "Encoder"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Encoder {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Encoder",
				Label: "Encoder",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Software", inst.Software, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EncodingDescription", inst.EncodingDescription, probe.formStage, formGroup)

		{
			// Slice of pointers: Supports
			div := (&form.FormDiv{Name: "Supports"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Supports {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Supports",
				Label: "Supports",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Identifications {
				if src.Encoding == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Identification", "Encoding", refNames, formGroup, probe.formStage)
		}
	case *common.Fermata:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Value", inst.Value, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				for _, target := range src.Fermata {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Fermata", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.BarlineStage.Barlines {
				for _, target := range src.Fermata {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("barline.Barline", "Fermata", refNames, formGroup, probe.formStage)
		}
	case *common.Fingering:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Substitution", inst.Substitution, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Alternate", inst.Alternate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Fingering {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Fingering", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.FrameNotes {
				if src.Fingering == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.FrameNote", "Fingering", refNames, formGroup, probe.formStage)
		}
	case *common.FormattedSymbol:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineThrough", inst.LineThrough, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LetterSpacing", inst.LetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineHeight", inst.LineHeight, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)
	case *common.FormattedSymbolId:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineThrough", inst.LineThrough, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LetterSpacing", inst.LetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineHeight", inst.LineHeight, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Credits {
				for _, target := range src.CreditSymbol {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Credit", "CreditSymbol", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				for _, target := range src.Symbol {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Symbol", refNames, formGroup, probe.formStage)
		}
	case *common.FormattedText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space", inst.Space, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineThrough", inst.LineThrough, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LetterSpacing", inst.LetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineHeight", inst.LineHeight, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartGroups {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartGroup", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Backups {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Backup", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Figures {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Figure", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.FiguredBasss {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.FiguredBass", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Forwards {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Forward", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Lyrics {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Lyric", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.NoteheadTexts {
				for _, target := range src.DisplayText {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.NoteheadText", "DisplayText", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Directions {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Direction", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.BarlineStage.Barlines {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("barline.Barline", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				if src.Footnote == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "Footnote", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.NameDisplays {
				for _, target := range src.DisplayText {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.NameDisplay", "DisplayText", refNames, formGroup, probe.formStage)
		}
	case *common.FormattedTextId:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Space", inst.Space, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Justify", inst.Justify, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Underline", inst.Underline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Overline", inst.Overline, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineThrough", inst.LineThrough, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Rotation", inst.Rotation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LetterSpacing", inst.LetterSpacing, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LineHeight", inst.LineHeight, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Dir", inst.Dir, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Enclosure", inst.Enclosure, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Credits {
				for _, target := range src.CreditWords {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Credit", "CreditWords", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				for _, target := range src.Rehearsal {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Rehearsal", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				for _, target := range src.Words {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Words", refNames, formGroup, probe.formStage)
		}
	case *common.Fret:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.Fret {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "Fret", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.FrameNotes {
				if src.Fret == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.FrameNote", "Fret", refNames, formGroup, probe.formStage)
		}
	case *common.Glyph:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Appearances {
				for _, target := range src.Glyph {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Appearance", "Glyph", refNames, formGroup, probe.formStage)
		}
	case *common.HorizontalTurn:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Slash", inst.Slash, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("StartNote", inst.StartNote, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TrillStep", inst.TrillStep, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TwoNoteTurn", inst.TwoNoteTurn, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SecondBeat", inst.SecondBeat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LastBeat", inst.LastBeat, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.Turn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "Turn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.DelayedTurn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "DelayedTurn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.InvertedTurn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "InvertedTurn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.DelayedInvertedTurn {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "DelayedInvertedTurn", refNames, formGroup, probe.formStage)
		}
	case *common.Identification:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Creator
			div := (&form.FormDiv{Name: "Creator"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Creator {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Creator",
				Label: "Creator",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Rights
			div := (&form.FormDiv{Name: "Rights"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Rights {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Rights",
				Label: "Rights",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Encoding", inst.Encoding, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Encoding](), probe.formStage)
		StageSetBasicFieldtoForm("Source", inst.Source, probe.formStage, formGroup)

		{
			// Slice of pointers: Relation
			div := (&form.FormDiv{Name: "Relation"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Relation {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Relation",
				Label: "Relation",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Miscellaneous", inst.Miscellaneous, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.Miscellaneous](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				if src.Identification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "Identification", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScorePartwises {
				if src.Identification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePartwise", "Identification", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewises {
				if src.Identification == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewise", "Identification", refNames, formGroup, probe.formStage)
		}
	case *common.Level:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Reference", inst.Reference, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Parentheses", inst.Parentheses, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Bracket", inst.Bracket, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Size", inst.Size, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartGroups {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartGroup", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Backups {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Backup", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Figures {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Figure", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.FiguredBasss {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.FiguredBass", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Forwards {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Forward", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Lyrics {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Lyric", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notationss {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Notations", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Directions {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Direction", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Harmonys {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Harmony", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.BarlineStage.Barlines {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("barline.Barline", "Level", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.AttributesStage.Attributess {
				if src.Level == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("attributes.Attributes", "Level", refNames, formGroup, probe.formStage)
		}
	case *common.LineWidth:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Appearances {
				for _, target := range src.LineWidth {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Appearance", "LineWidth", refNames, formGroup, probe.formStage)
		}
	case *common.Link:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameAttr", inst.NameAttr, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Show", inst.Show, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Actuate", inst.Actuate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Element", inst.Element, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Position", inst.Position, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Credits {
				for _, target := range src.Link {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Credit", "Link", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.MeasureMusicDataChoices {
				if src.Link == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.MeasureMusicDataChoice", "Link", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreTimewiseMeasureParts {
				for _, target := range src.Link {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScoreTimewiseMeasurePart", "Link", refNames, formGroup, probe.formStage)
		}
	case *common.MeasureLayout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MeasureDistance", inst.MeasureDistance, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Prints {
				if src.MeasureLayout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Print", "MeasureLayout", refNames, formGroup, probe.formStage)
		}
	case *common.MidiDevice:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Port", inst.Port, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				for _, target := range src.MidiDevice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "MidiDevice", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Sounds {
				for _, target := range src.MidiDevice {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Sound", "MidiDevice", refNames, formGroup, probe.formStage)
		}
	case *common.MidiInstrument:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MidiChannel", inst.MidiChannel, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MidiName", inst.MidiName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MidiBank", inst.MidiBank, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MidiProgram", inst.MidiProgram, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MidiUnpitched", inst.MidiUnpitched, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Volume", inst.Volume, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Pan", inst.Pan, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Elevation", inst.Elevation, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				for _, target := range src.MidiInstrument {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "MidiInstrument", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Sounds {
				for _, target := range src.MidiInstrument {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Sound", "MidiInstrument", refNames, formGroup, probe.formStage)
		}
	case *common.Miscellaneous:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: MiscellaneousField
			div := (&form.FormDiv{Name: "MiscellaneousField"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.MiscellaneousField {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "MiscellaneousField",
				Label: "MiscellaneousField",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Identifications {
				if src.Miscellaneous == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Identification", "Miscellaneous", refNames, formGroup, probe.formStage)
		}
	case *common.MiscellaneousField:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameAttr", inst.NameAttr, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Miscellaneouss {
				for _, target := range src.MiscellaneousField {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Miscellaneous", "MiscellaneousField", refNames, formGroup, probe.formStage)
		}
	case *common.NameDisplay:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("PrintObject", inst.PrintObject, formGroup, probe.formStage)

		{
			// Slice of pointers: DisplayText
			div := (&form.FormDiv{Name: "DisplayText"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DisplayText {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DisplayText",
				Label: "DisplayText",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AccidentalText
			div := (&form.FormDiv{Name: "AccidentalText"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AccidentalText {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AccidentalText",
				Label: "AccidentalText",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartGroups {
				if src.GroupNameDisplay == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartGroup", "GroupNameDisplay", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.PartGroups {
				if src.GroupAbbreviationDisplay == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.PartGroup", "GroupAbbreviationDisplay", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				if src.PartNameDisplay == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "PartNameDisplay", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.ScoreParts {
				if src.PartAbbreviationDisplay == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.ScorePart", "PartAbbreviationDisplay", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Prints {
				if src.PartNameDisplay == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Print", "PartNameDisplay", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Prints {
				if src.PartAbbreviationDisplay == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Print", "PartAbbreviationDisplay", refNames, formGroup, probe.formStage)
		}
	case *common.NoteSize:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Appearances {
				for _, target := range src.NoteSize {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Appearance", "NoteSize", refNames, formGroup, probe.formStage)
		}
	case *common.OtherAppearance:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Appearances {
				for _, target := range src.OtherAppearance {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Appearance", "OtherAppearance", refNames, formGroup, probe.formStage)
		}
	case *common.OtherPlay:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Plays {
				for _, target := range src.OtherPlay {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Play", "OtherPlay", refNames, formGroup, probe.formStage)
		}
	case *common.OtherText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Percussions {
				if src.OtherPercussion == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Percussion", "OtherPercussion", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Dynamicss {
				for _, target := range src.OtherDynamics {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Dynamics", "OtherDynamics", refNames, formGroup, probe.formStage)
		}
	case *common.PageLayout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PageHeight", inst.PageHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PageWidth", inst.PageWidth, probe.formStage, formGroup)

		{
			// Slice of pointers: PageMargins
			div := (&form.FormDiv{Name: "PageMargins"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PageMargins {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PageMargins",
				Label: "PageMargins",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				if src.PageLayout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "PageLayout", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Prints {
				if src.PageLayout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Print", "PageLayout", refNames, formGroup, probe.formStage)
		}
	case *common.PageMargins:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("LeftMargin", inst.LeftMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RightMargin", inst.RightMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TopMargin", inst.TopMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BottomMargin", inst.BottomMargin, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.PageLayouts {
				for _, target := range src.PageMargins {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.PageLayout", "PageMargins", refNames, formGroup, probe.formStage)
		}
	case *common.Play:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Ipa", inst.Ipa, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Mute", inst.Mute, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("SemiPitched", inst.SemiPitched, formGroup, probe.formStage)

		{
			// Slice of pointers: OtherPlay
			div := (&form.FormDiv{Name: "OtherPlay"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.OtherPlay {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "OtherPlay",
				Label: "OtherPlay",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Notes {
				if src.Play == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Note", "Play", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Sounds {
				for _, target := range src.Play {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Sound", "Play", refNames, formGroup, probe.formStage)
		}
	case *common.Scaling:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Millimeters", inst.Millimeters, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Tenths", inst.Tenths, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				if src.Scaling == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "Scaling", refNames, formGroup, probe.formStage)
		}
	case *common.Segno:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Halign", inst.Halign, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Valign", inst.Valign, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.DirectionTypes {
				for _, target := range src.Segno {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.DirectionType", "Segno", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.BarlineStage.Barlines {
				if src.Segno_1 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("barline.Barline", "Segno_1", refNames, formGroup, probe.formStage)
		}
	case *common.StaffLayout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StaffDistance", inst.StaffDistance, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				for _, target := range src.StaffLayout {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "StaffLayout", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Prints {
				for _, target := range src.StaffLayout {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Print", "StaffLayout", refNames, formGroup, probe.formStage)
		}
	case *common.StringType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FontFamily", inst.FontFamily, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("FontSize", inst.FontSize, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontWeight", inst.FontWeight, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Technicals {
				for _, target := range src.String {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Technical", "String", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.FrameNotes {
				if src.String == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.FrameNote", "String", refNames, formGroup, probe.formStage)
		}
	case *common.Supports:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Element", inst.Element, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Attribute", inst.Attribute, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Encodings {
				for _, target := range src.Supports {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Encoding", "Supports", refNames, formGroup, probe.formStage)
		}
	case *common.SystemDividers:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("LeftDivider", inst.LeftDivider, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintObjectStyleAlign](), probe.formStage)
		StageSetAssociationFieldToForm("RightDivider", inst.RightDivider, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.EmptyPrintObjectStyleAlign](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.SystemLayouts {
				if src.SystemDividers == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.SystemLayout", "SystemDividers", refNames, formGroup, probe.formStage)
		}
	case *common.SystemLayout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SystemMargins", inst.SystemMargins, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.SystemMargins](), probe.formStage)
		StageSetBasicFieldtoForm("SystemDistance", inst.SystemDistance, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TopSystemDistance", inst.TopSystemDistance, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SystemDividers", inst.SystemDividers, formGroup, probe.stageSet.CommonStage.GetInstancesSet[*common.SystemDividers](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.ScoreStage.Defaultss {
				if src.SystemLayout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("score.Defaults", "SystemLayout", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.DirectionStage.Prints {
				if src.SystemLayout == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("direction.Print", "SystemLayout", refNames, formGroup, probe.formStage)
		}
	case *common.SystemMargins:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LeftMargin", inst.LeftMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RightMargin", inst.RightMargin, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.SystemLayouts {
				if src.SystemMargins == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.SystemLayout", "SystemMargins", refNames, formGroup, probe.formStage)
		}
	case *common.TypedText:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Type", inst.Type, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Encodings {
				for _, target := range src.Encoder {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Encoding", "Encoder", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Identifications {
				for _, target := range src.Creator {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Identification", "Creator", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Identifications {
				for _, target := range src.Rights {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Identification", "Rights", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.CommonStage.Identifications {
				for _, target := range src.Relation {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("common.Identification", "Relation", refNames, formGroup, probe.formStage)
		}
	case *common.WavyLine:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Number", inst.Number, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Smufl", inst.Smufl, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultX", inst.DefaultX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultY", inst.DefaultY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeX", inst.RelativeX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RelativeY", inst.RelativeY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Placement", inst.Placement, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Color", inst.Color, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartNote", inst.StartNote, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TrillStep", inst.TrillStep, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("TwoNoteTurn", inst.TwoNoteTurn, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Accelerate", inst.Accelerate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Beats", inst.Beats, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SecondBeat", inst.SecondBeat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LastBeat", inst.LastBeat, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.NoteStage.Ornamentss {
				for _, target := range src.WavyLine {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("note.Ornaments", "WavyLine", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.BarlineStage.Barlines {
				if src.WavyLine == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("barline.Barline", "WavyLine", refNames, formGroup, probe.formStage)
		}
	case *xlink.ArcType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Show", inst.Show, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Actuate", inst.Actuate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("From", inst.From, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("To", inst.To, probe.formStage, formGroup)

		{
			// Slice of pointers: Title_1
			div := (&form.FormDiv{Name: "Title_1"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Title_1 {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Title_1",
				Label: "Title_1",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.XlinkStage.Extendeds {
				for _, target := range src.Arc {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.Extended", "Arc", refNames, formGroup, probe.formStage)
		}
	case *xlink.Extended:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)

		{
			// Slice of pointers: Title_1
			div := (&form.FormDiv{Name: "Title_1"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Title_1 {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Title_1",
				Label: "Title_1",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Resource
			div := (&form.FormDiv{Name: "Resource"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Resource {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Resource",
				Label: "Resource",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Locator
			div := (&form.FormDiv{Name: "Locator"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Locator {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Locator",
				Label: "Locator",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Arc
			div := (&form.FormDiv{Name: "Arc"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Arc {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Arc",
				Label: "Arc",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *xlink.LocatorType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)

		{
			// Slice of pointers: Title_1
			div := (&form.FormDiv{Name: "Title_1"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Title_1 {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Title_1",
				Label: "Title_1",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.XlinkStage.Extendeds {
				for _, target := range src.Locator {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.Extended", "Locator", refNames, formGroup, probe.formStage)
		}
	case *xlink.ResourceType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.XlinkStage.Extendeds {
				for _, target := range src.Resource {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.Extended", "Resource", refNames, formGroup, probe.formStage)
		}
	case *xlink.Simple:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Href", inst.Href, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Role", inst.Role, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Arcrole", inst.Arcrole, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Show", inst.Show, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Actuate", inst.Actuate, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)
	case *xlink.TitleEltType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.XlinkStage.ArcTypes {
				for _, target := range src.Title_1 {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.ArcType", "Title_1", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.XlinkStage.Extendeds {
				for _, target := range src.Title_1 {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.Extended", "Title_1", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.XlinkStage.LocatorTypes {
				for _, target := range src.Title_1 {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("xlink.LocatorType", "Title_1", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
