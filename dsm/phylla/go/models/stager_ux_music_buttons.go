package models

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) ux_button_music() {
	stager.buttonMusicStage.Reset()

	plant := stager.GetCurrentPlant()
	if plant == nil || plant.PlantType != Music || plant.MusicAbstract == nil {
		stager.buttonMusicStage.Commit()
		return
	}

	layout := new(button.Layout).Stage(stager.buttonMusicStage)

	group1 := new(button.Group).Stage(stager.buttonMusicStage)
	group1.Percentage = 100
	group1.NbColumns = 2
	layout.Groups = append(layout.Groups, group1)

	buttonExportToMuseScore := button.NewButton(
		&ExportToMuseScoreButtonProxy{
			stager: stager,
		},
		"Export to Musescore",
		string(buttons.BUTTON_music_note),
		"Export to Musescore",
	)
	buttonExportToMuseScore.MatButtonType = button.MatButtonTypeBasic
	buttonExportToMuseScore.MatButtonAppearance = button.MatButtonAppearanceOutlined
	buttonExportToMuseScore.Color = button.MatButtonPaletteTypePrimary
	group1.Buttons = append(group1.Buttons, buttonExportToMuseScore)

	buttonExportStaticSite := button.NewButton(
		&ExportStaticSiteButtonProxy{
			stager: stager,
		},
		"Export Static Web Site",
		string(buttons.BUTTON_web),
		"Export Static Web Site",
	)
	buttonExportStaticSite.MatButtonType = button.MatButtonTypeBasic
	buttonExportStaticSite.MatButtonAppearance = button.MatButtonAppearanceOutlined
	buttonExportStaticSite.Color = button.MatButtonPaletteTypePrimary
	group1.Buttons = append(group1.Buttons, buttonExportStaticSite)

	stager.buttonMusicStage.Commit()
}

type ExportToMuseScoreButtonProxy struct {
	stager *Stager
}

func (e *ExportToMuseScoreButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonMusicStage
}

func (e *ExportToMuseScoreButtonProxy) OnAfterUpdateButton() {
	e.stager.exportMusicXML()
}

type ExportStaticSiteButtonProxy struct {
	stager *Stager
}

func (e *ExportStaticSiteButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonMusicStage
}

func (e *ExportStaticSiteButtonProxy) OnAfterUpdateButton() {
	e.stager.exportWebsite()
}

func (stager *Stager) exportMusicXML() {
	plant := stager.GetCurrentPlant()
	if plant == nil || plant.MusicAbstract == nil {
		return
	}

	stager.loadStage.Reset()
	fileToDownload := new(load.FileToDownload).Stage(stager.loadStage)

	xmlContent := generateSimpleMusicXML(plant, stager.musicNotes)

	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(xmlContent))
	fileToDownload.Name = time.Now().Format("20060102 1504 ") + "phylla-" + plant.Name + ".musicxml"
	stager.loadStage.Commit()
}

func generateSimpleMusicXML(plant *PlantAbstract, notes []*MusicNoteData) string {
	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	sb.WriteString(`<!DOCTYPE score-partwise PUBLIC "-//Recordare//DTD MusicXML 4.0 Partwise//EN" "http://www.musicxml.org/dtds/partwise.dtd">` + "\n")
	sb.WriteString(`<score-partwise version="4.0">` + "\n")
	sb.WriteString(`  <work><work-title>` + plant.Name + `</work-title></work>` + "\n")
	sb.WriteString(`  <part-list>` + "\n")
	sb.WriteString(`    <score-part id="P1"><part-name>Piano</part-name></score-part>` + "\n")
	sb.WriteString(`  </part-list>` + "\n")
	sb.WriteString(`  <part id="P1">` + "\n")

	stepNames := []string{"C", "C", "D", "D", "E", "F", "F", "G", "G", "A", "A", "B"}
	alterVals := []int{0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0}

	nbBeats := plant.MusicAbstract.NbOfBeatsInTheme
	if nbBeats <= 0 {
		nbBeats = 16
	}

	notesByBeat := make(map[int][]*MusicNoteData)
	for _, nd := range notes {
		if nd.IsKept {
			notesByBeat[nd.BeatNb] = append(notesByBeat[nd.BeatNb], nd)
		}
	}

	measureNb := 1
	sb.WriteString(fmt.Sprintf(`    <measure number="%d">`+"\n", measureNb))
	sb.WriteString(`      <attributes>` + "\n")
	sb.WriteString(`        <divisions>1</divisions>` + "\n")
	sb.WriteString(`        <key><fifths>0</fifths></key>` + "\n")
	sb.WriteString(fmt.Sprintf(`        <time><beats>%d</beats><beat-type>4</beat-type></time>`+"\n", nbBeats))
	sb.WriteString(`        <clef><sign>G</sign><line>2</line></clef>` + "\n")
	sb.WriteString(`      </attributes>` + "\n")

	for beat := 0; beat < nbBeats; beat++ {
		beatNotes := notesByBeat[beat]
		if len(beatNotes) == 0 {
			sb.WriteString(`      <note><rest/><duration>1</duration><type>quarter</type></note>` + "\n")
		} else {
			for i, n := range beatNotes {
				pitch := n.Pitch + 36 // start around C3
				step := stepNames[pitch%12]
				alter := alterVals[pitch%12]
				octave := (pitch / 12) - 1

				sb.WriteString(`      <note>` + "\n")
				if i > 0 {
					sb.WriteString(`        <chord/>` + "\n")
				}
				sb.WriteString(`        <pitch>` + "\n")
				sb.WriteString(fmt.Sprintf(`          <step>%s</step>`+"\n", step))
				if alter != 0 {
					sb.WriteString(fmt.Sprintf(`          <alter>%d</alter>`+"\n", alter))
				}
				sb.WriteString(fmt.Sprintf(`          <octave>%d</octave>`+"\n", octave))
				sb.WriteString(`        </pitch>` + "\n")
				sb.WriteString(`        <duration>1</duration>` + "\n")
				sb.WriteString(`        <type>quarter</type>` + "\n")
				sb.WriteString(`      </note>` + "\n")
			}
		}
	}
	sb.WriteString(`    </measure>` + "\n")
	sb.WriteString(`  </part>` + "\n")
	sb.WriteString(`</score-partwise>` + "\n")

	return sb.String()
}
