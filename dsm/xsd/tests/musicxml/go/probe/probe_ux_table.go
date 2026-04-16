// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"time"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/musicxml/go/models"
)

const TableName = "Table"

// update the current table if there is one
func (probe *Probe) ux_table() {
	var tableName string
	for table := range probe.tableStage.Tables {
		tableName = table.Name
	}
	switch tableName {
	// insertion point
	case "A_directive":
		updateProbeTable[*models.A_directive](probe)
	case "A_measure":
		updateProbeTable[*models.A_measure](probe)
	case "A_measure_1":
		updateProbeTable[*models.A_measure_1](probe)
	case "A_part":
		updateProbeTable[*models.A_part](probe)
	case "A_part_1":
		updateProbeTable[*models.A_part_1](probe)
	case "Accidental":
		updateProbeTable[*models.Accidental](probe)
	case "Accidental_mark":
		updateProbeTable[*models.Accidental_mark](probe)
	case "Accidental_text":
		updateProbeTable[*models.Accidental_text](probe)
	case "Accord":
		updateProbeTable[*models.Accord](probe)
	case "Accordion_registration":
		updateProbeTable[*models.Accordion_registration](probe)
	case "Appearance":
		updateProbeTable[*models.Appearance](probe)
	case "Arpeggiate":
		updateProbeTable[*models.Arpeggiate](probe)
	case "Arrow":
		updateProbeTable[*models.Arrow](probe)
	case "Articulations":
		updateProbeTable[*models.Articulations](probe)
	case "Assess":
		updateProbeTable[*models.Assess](probe)
	case "Attributes":
		updateProbeTable[*models.Attributes](probe)
	case "Backup":
		updateProbeTable[*models.Backup](probe)
	case "Bar_style_color":
		updateProbeTable[*models.Bar_style_color](probe)
	case "Barline":
		updateProbeTable[*models.Barline](probe)
	case "Barre":
		updateProbeTable[*models.Barre](probe)
	case "Bass":
		updateProbeTable[*models.Bass](probe)
	case "Bass_step":
		updateProbeTable[*models.Bass_step](probe)
	case "Beam":
		updateProbeTable[*models.Beam](probe)
	case "Beat_repeat":
		updateProbeTable[*models.Beat_repeat](probe)
	case "Beat_unit_tied":
		updateProbeTable[*models.Beat_unit_tied](probe)
	case "Beater":
		updateProbeTable[*models.Beater](probe)
	case "Bend":
		updateProbeTable[*models.Bend](probe)
	case "Bookmark":
		updateProbeTable[*models.Bookmark](probe)
	case "Bracket":
		updateProbeTable[*models.Bracket](probe)
	case "Breath_mark":
		updateProbeTable[*models.Breath_mark](probe)
	case "Caesura":
		updateProbeTable[*models.Caesura](probe)
	case "Cancel":
		updateProbeTable[*models.Cancel](probe)
	case "Clef":
		updateProbeTable[*models.Clef](probe)
	case "Coda":
		updateProbeTable[*models.Coda](probe)
	case "Credit":
		updateProbeTable[*models.Credit](probe)
	case "Dashes":
		updateProbeTable[*models.Dashes](probe)
	case "Defaults":
		updateProbeTable[*models.Defaults](probe)
	case "Degree":
		updateProbeTable[*models.Degree](probe)
	case "Degree_alter":
		updateProbeTable[*models.Degree_alter](probe)
	case "Degree_type":
		updateProbeTable[*models.Degree_type](probe)
	case "Degree_value":
		updateProbeTable[*models.Degree_value](probe)
	case "Direction":
		updateProbeTable[*models.Direction](probe)
	case "Direction_type":
		updateProbeTable[*models.Direction_type](probe)
	case "Distance":
		updateProbeTable[*models.Distance](probe)
	case "Double":
		updateProbeTable[*models.Double](probe)
	case "Dynamics":
		updateProbeTable[*models.Dynamics](probe)
	case "Effect":
		updateProbeTable[*models.Effect](probe)
	case "Elision":
		updateProbeTable[*models.Elision](probe)
	case "Empty":
		updateProbeTable[*models.Empty](probe)
	case "Empty_font":
		updateProbeTable[*models.Empty_font](probe)
	case "Empty_line":
		updateProbeTable[*models.Empty_line](probe)
	case "Empty_placement":
		updateProbeTable[*models.Empty_placement](probe)
	case "Empty_placement_smufl":
		updateProbeTable[*models.Empty_placement_smufl](probe)
	case "Empty_print_object_style_align":
		updateProbeTable[*models.Empty_print_object_style_align](probe)
	case "Empty_print_style":
		updateProbeTable[*models.Empty_print_style](probe)
	case "Empty_print_style_align":
		updateProbeTable[*models.Empty_print_style_align](probe)
	case "Empty_print_style_align_id":
		updateProbeTable[*models.Empty_print_style_align_id](probe)
	case "Empty_trill_sound":
		updateProbeTable[*models.Empty_trill_sound](probe)
	case "Encoding":
		updateProbeTable[*models.Encoding](probe)
	case "Ending":
		updateProbeTable[*models.Ending](probe)
	case "Extend":
		updateProbeTable[*models.Extend](probe)
	case "Feature":
		updateProbeTable[*models.Feature](probe)
	case "Fermata":
		updateProbeTable[*models.Fermata](probe)
	case "Figure":
		updateProbeTable[*models.Figure](probe)
	case "Figured_bass":
		updateProbeTable[*models.Figured_bass](probe)
	case "Fingering":
		updateProbeTable[*models.Fingering](probe)
	case "First_fret":
		updateProbeTable[*models.First_fret](probe)
	case "For_part":
		updateProbeTable[*models.For_part](probe)
	case "Formatted_symbol":
		updateProbeTable[*models.Formatted_symbol](probe)
	case "Formatted_symbol_id":
		updateProbeTable[*models.Formatted_symbol_id](probe)
	case "Formatted_text":
		updateProbeTable[*models.Formatted_text](probe)
	case "Formatted_text_id":
		updateProbeTable[*models.Formatted_text_id](probe)
	case "Forward":
		updateProbeTable[*models.Forward](probe)
	case "Frame":
		updateProbeTable[*models.Frame](probe)
	case "Frame_note":
		updateProbeTable[*models.Frame_note](probe)
	case "Fret":
		updateProbeTable[*models.Fret](probe)
	case "Glass":
		updateProbeTable[*models.Glass](probe)
	case "Glissando":
		updateProbeTable[*models.Glissando](probe)
	case "Glyph":
		updateProbeTable[*models.Glyph](probe)
	case "Grace":
		updateProbeTable[*models.Grace](probe)
	case "Group_barline":
		updateProbeTable[*models.Group_barline](probe)
	case "Group_name":
		updateProbeTable[*models.Group_name](probe)
	case "Group_symbol":
		updateProbeTable[*models.Group_symbol](probe)
	case "Grouping":
		updateProbeTable[*models.Grouping](probe)
	case "Hammer_on_pull_off":
		updateProbeTable[*models.Hammer_on_pull_off](probe)
	case "Handbell":
		updateProbeTable[*models.Handbell](probe)
	case "Harmon_closed":
		updateProbeTable[*models.Harmon_closed](probe)
	case "Harmon_mute":
		updateProbeTable[*models.Harmon_mute](probe)
	case "Harmonic":
		updateProbeTable[*models.Harmonic](probe)
	case "Harmony":
		updateProbeTable[*models.Harmony](probe)
	case "Harmony_alter":
		updateProbeTable[*models.Harmony_alter](probe)
	case "Harp_pedals":
		updateProbeTable[*models.Harp_pedals](probe)
	case "Heel_toe":
		updateProbeTable[*models.Heel_toe](probe)
	case "Hole":
		updateProbeTable[*models.Hole](probe)
	case "Hole_closed":
		updateProbeTable[*models.Hole_closed](probe)
	case "Horizontal_turn":
		updateProbeTable[*models.Horizontal_turn](probe)
	case "Identification":
		updateProbeTable[*models.Identification](probe)
	case "Image":
		updateProbeTable[*models.Image](probe)
	case "Instrument":
		updateProbeTable[*models.Instrument](probe)
	case "Instrument_change":
		updateProbeTable[*models.Instrument_change](probe)
	case "Instrument_link":
		updateProbeTable[*models.Instrument_link](probe)
	case "Interchangeable":
		updateProbeTable[*models.Interchangeable](probe)
	case "Inversion":
		updateProbeTable[*models.Inversion](probe)
	case "Key":
		updateProbeTable[*models.Key](probe)
	case "Key_accidental":
		updateProbeTable[*models.Key_accidental](probe)
	case "Key_octave":
		updateProbeTable[*models.Key_octave](probe)
	case "Kind":
		updateProbeTable[*models.Kind](probe)
	case "Level":
		updateProbeTable[*models.Level](probe)
	case "Line_detail":
		updateProbeTable[*models.Line_detail](probe)
	case "Line_width":
		updateProbeTable[*models.Line_width](probe)
	case "Link":
		updateProbeTable[*models.Link](probe)
	case "Listen":
		updateProbeTable[*models.Listen](probe)
	case "Listening":
		updateProbeTable[*models.Listening](probe)
	case "Lyric":
		updateProbeTable[*models.Lyric](probe)
	case "Lyric_font":
		updateProbeTable[*models.Lyric_font](probe)
	case "Lyric_language":
		updateProbeTable[*models.Lyric_language](probe)
	case "Measure_layout":
		updateProbeTable[*models.Measure_layout](probe)
	case "Measure_numbering":
		updateProbeTable[*models.Measure_numbering](probe)
	case "Measure_repeat":
		updateProbeTable[*models.Measure_repeat](probe)
	case "Measure_style":
		updateProbeTable[*models.Measure_style](probe)
	case "Membrane":
		updateProbeTable[*models.Membrane](probe)
	case "Metal":
		updateProbeTable[*models.Metal](probe)
	case "Metronome":
		updateProbeTable[*models.Metronome](probe)
	case "Metronome_beam":
		updateProbeTable[*models.Metronome_beam](probe)
	case "Metronome_note":
		updateProbeTable[*models.Metronome_note](probe)
	case "Metronome_tied":
		updateProbeTable[*models.Metronome_tied](probe)
	case "Metronome_tuplet":
		updateProbeTable[*models.Metronome_tuplet](probe)
	case "Midi_device":
		updateProbeTable[*models.Midi_device](probe)
	case "Midi_instrument":
		updateProbeTable[*models.Midi_instrument](probe)
	case "Miscellaneous":
		updateProbeTable[*models.Miscellaneous](probe)
	case "Miscellaneous_field":
		updateProbeTable[*models.Miscellaneous_field](probe)
	case "Mordent":
		updateProbeTable[*models.Mordent](probe)
	case "Multiple_rest":
		updateProbeTable[*models.Multiple_rest](probe)
	case "Name_display":
		updateProbeTable[*models.Name_display](probe)
	case "Non_arpeggiate":
		updateProbeTable[*models.Non_arpeggiate](probe)
	case "Notations":
		updateProbeTable[*models.Notations](probe)
	case "Note":
		updateProbeTable[*models.Note](probe)
	case "Note_size":
		updateProbeTable[*models.Note_size](probe)
	case "Note_type":
		updateProbeTable[*models.Note_type](probe)
	case "Notehead":
		updateProbeTable[*models.Notehead](probe)
	case "Notehead_text":
		updateProbeTable[*models.Notehead_text](probe)
	case "Numeral":
		updateProbeTable[*models.Numeral](probe)
	case "Numeral_key":
		updateProbeTable[*models.Numeral_key](probe)
	case "Numeral_root":
		updateProbeTable[*models.Numeral_root](probe)
	case "Octave_shift":
		updateProbeTable[*models.Octave_shift](probe)
	case "Offset":
		updateProbeTable[*models.Offset](probe)
	case "Opus":
		updateProbeTable[*models.Opus](probe)
	case "Ornaments":
		updateProbeTable[*models.Ornaments](probe)
	case "Other_appearance":
		updateProbeTable[*models.Other_appearance](probe)
	case "Other_direction":
		updateProbeTable[*models.Other_direction](probe)
	case "Other_listening":
		updateProbeTable[*models.Other_listening](probe)
	case "Other_notation":
		updateProbeTable[*models.Other_notation](probe)
	case "Other_placement_text":
		updateProbeTable[*models.Other_placement_text](probe)
	case "Other_play":
		updateProbeTable[*models.Other_play](probe)
	case "Other_text":
		updateProbeTable[*models.Other_text](probe)
	case "Page_layout":
		updateProbeTable[*models.Page_layout](probe)
	case "Page_margins":
		updateProbeTable[*models.Page_margins](probe)
	case "Part_clef":
		updateProbeTable[*models.Part_clef](probe)
	case "Part_group":
		updateProbeTable[*models.Part_group](probe)
	case "Part_link":
		updateProbeTable[*models.Part_link](probe)
	case "Part_list":
		updateProbeTable[*models.Part_list](probe)
	case "Part_name":
		updateProbeTable[*models.Part_name](probe)
	case "Part_symbol":
		updateProbeTable[*models.Part_symbol](probe)
	case "Part_transpose":
		updateProbeTable[*models.Part_transpose](probe)
	case "Pedal":
		updateProbeTable[*models.Pedal](probe)
	case "Pedal_tuning":
		updateProbeTable[*models.Pedal_tuning](probe)
	case "Per_minute":
		updateProbeTable[*models.Per_minute](probe)
	case "Percussion":
		updateProbeTable[*models.Percussion](probe)
	case "Pitch":
		updateProbeTable[*models.Pitch](probe)
	case "Pitched":
		updateProbeTable[*models.Pitched](probe)
	case "Placement_text":
		updateProbeTable[*models.Placement_text](probe)
	case "Play":
		updateProbeTable[*models.Play](probe)
	case "Player":
		updateProbeTable[*models.Player](probe)
	case "Principal_voice":
		updateProbeTable[*models.Principal_voice](probe)
	case "Print":
		updateProbeTable[*models.Print](probe)
	case "Release":
		updateProbeTable[*models.Release](probe)
	case "Repeat":
		updateProbeTable[*models.Repeat](probe)
	case "Rest":
		updateProbeTable[*models.Rest](probe)
	case "Root":
		updateProbeTable[*models.Root](probe)
	case "Root_step":
		updateProbeTable[*models.Root_step](probe)
	case "Scaling":
		updateProbeTable[*models.Scaling](probe)
	case "Scordatura":
		updateProbeTable[*models.Scordatura](probe)
	case "Score_instrument":
		updateProbeTable[*models.Score_instrument](probe)
	case "Score_part":
		updateProbeTable[*models.Score_part](probe)
	case "Score_partwise":
		updateProbeTable[*models.Score_partwise](probe)
	case "Score_timewise":
		updateProbeTable[*models.Score_timewise](probe)
	case "Segno":
		updateProbeTable[*models.Segno](probe)
	case "Slash":
		updateProbeTable[*models.Slash](probe)
	case "Slide":
		updateProbeTable[*models.Slide](probe)
	case "Slur":
		updateProbeTable[*models.Slur](probe)
	case "Sound":
		updateProbeTable[*models.Sound](probe)
	case "Staff_details":
		updateProbeTable[*models.Staff_details](probe)
	case "Staff_divide":
		updateProbeTable[*models.Staff_divide](probe)
	case "Staff_layout":
		updateProbeTable[*models.Staff_layout](probe)
	case "Staff_size":
		updateProbeTable[*models.Staff_size](probe)
	case "Staff_tuning":
		updateProbeTable[*models.Staff_tuning](probe)
	case "Stem":
		updateProbeTable[*models.Stem](probe)
	case "Stick":
		updateProbeTable[*models.Stick](probe)
	case "String_mute":
		updateProbeTable[*models.String_mute](probe)
	case "String_type":
		updateProbeTable[*models.String_type](probe)
	case "Strong_accent":
		updateProbeTable[*models.Strong_accent](probe)
	case "Style_text":
		updateProbeTable[*models.Style_text](probe)
	case "Supports":
		updateProbeTable[*models.Supports](probe)
	case "Swing":
		updateProbeTable[*models.Swing](probe)
	case "Sync":
		updateProbeTable[*models.Sync](probe)
	case "System_dividers":
		updateProbeTable[*models.System_dividers](probe)
	case "System_layout":
		updateProbeTable[*models.System_layout](probe)
	case "System_margins":
		updateProbeTable[*models.System_margins](probe)
	case "Tap":
		updateProbeTable[*models.Tap](probe)
	case "Technical":
		updateProbeTable[*models.Technical](probe)
	case "Text_element_data":
		updateProbeTable[*models.Text_element_data](probe)
	case "Tie":
		updateProbeTable[*models.Tie](probe)
	case "Tied":
		updateProbeTable[*models.Tied](probe)
	case "Time":
		updateProbeTable[*models.Time](probe)
	case "Time_modification":
		updateProbeTable[*models.Time_modification](probe)
	case "Timpani":
		updateProbeTable[*models.Timpani](probe)
	case "Transpose":
		updateProbeTable[*models.Transpose](probe)
	case "Tremolo":
		updateProbeTable[*models.Tremolo](probe)
	case "Tuplet":
		updateProbeTable[*models.Tuplet](probe)
	case "Tuplet_dot":
		updateProbeTable[*models.Tuplet_dot](probe)
	case "Tuplet_number":
		updateProbeTable[*models.Tuplet_number](probe)
	case "Tuplet_portion":
		updateProbeTable[*models.Tuplet_portion](probe)
	case "Tuplet_type":
		updateProbeTable[*models.Tuplet_type](probe)
	case "Typed_text":
		updateProbeTable[*models.Typed_text](probe)
	case "Unpitched":
		updateProbeTable[*models.Unpitched](probe)
	case "Virtual_instrument":
		updateProbeTable[*models.Virtual_instrument](probe)
	case "Wait":
		updateProbeTable[*models.Wait](probe)
	case "Wavy_line":
		updateProbeTable[*models.Wavy_line](probe)
	case "Wedge":
		updateProbeTable[*models.Wedge](probe)
	case "Wood":
		updateProbeTable[*models.Wood](probe)
	case "Work":
		updateProbeTable[*models.Work](probe)
	}
}

func updateProbeTable[T models.PointerToGongstruct](
	probe *Probe,
) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = models.GetPointerToGongstructName[T]()
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = probe.bulkDeleteMode
	table.HasBulkDeleteButton = probe.bulkDeleteMode
	table.BulkDeleteButtonTooltip = "Select rows to delete and click to delete them"
	table.HasSaveButton = false

	fields := models.GetFieldsFromPointer[T]()
	reverseFields := models.GetReverseFields[T]()

	table.NbOfStickyColumns = 3

	setOfStructs := (*models.GetGongstructInstancesSetFromPointerType[T](probe.stageOfInterest))
	sliceOfGongStructsSorted := make([]T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return models.GetOrderPointerGongstruct(probe.stageOfInterest, sliceOfGongStructsSorted[i]) <
			models.GetOrderPointerGongstruct(probe.stageOfInterest, sliceOfGongStructsSorted[j])
	})

	// add a button for bulk delete
	bulkDeleteModeButton := &table_models.Button{
		Name: "BulkDeleteButton",
		Icon: func() string {
			if probe.bulkDeleteMode {
				return string(maticons.BUTTON_deselect)
			}
			return string(maticons.BUTTON_delete_sweep)
		}(),
		HasToolTip: true,
		ToolTipText: func() string {
			if probe.bulkDeleteMode {
				return "Cancel bulk delete"
			}
			return "Select instances to delete"
		}(),
		ToolTipPosition: table_models.Below,
	}

	bulkDeleteModeButton.OnClick = func() {
		probe.bulkDeleteMode = !probe.bulkDeleteMode
		updateProbeTable[T](probe)
	}
	table.Buttons = append(table.Buttons, bulkDeleteModeButton)

	column := new(table_models.DisplayedColumn)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	if !probe.bulkDeleteMode {
		column = new(table_models.DisplayedColumn)
		column.Name = "Delete"
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	} else {
		table.OnUpdate = func(stage *table_models.Stage, updatedTable *table_models.Table) {
			// log.Println("Table updated, selectedRowIDsForBulkDelete:", updatedTable.BulkDeleteSelectedRowsIDsJson)

			for _, row := range updatedTable.RowsSelectedForBulkDelete {
				cellID := row.Cells[0]
				id := cellID.CellInt.Value
				instance := models.GongGetInstanceFromOrder[T](probe.stageOfInterest, uint(id))
				var zeroInstance T
				if instance == zeroInstance {
					continue
				}
				instance.UnstageVoid(probe.stageOfInterest)
			}
			if len(updatedTable.RowsSelectedForBulkDelete) > 0 {
				// after a delete of an instance, the stage might be dirty if a pointer or a slice of pointer
				// reference the deleted instance.
				// therefore, it is mandatory to clean the stage of interest
				probe.bulkDeleteMode = false
				probe.stageOfInterest.Commit()
				probe.Refresh()
			}

			probe.stageOfInterest.Commit()
		}
	}

	for _, fieldName := range fields {
		column := new(table_models.DisplayedColumn)
		column.Name = fieldName.Name
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(table_models.DisplayedColumn)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(table_models.Row)
		value := models.GetFieldStringValueFromPointer(structInstance, "Name", probe.stageOfInterest)
		row.Name = value.GetValueString()

		updater := NewRowUpdate(structInstance, probe)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := &table_models.Cell{
			Name: "ID",
		}
		row.Cells = append(row.Cells, cell)
		cellInt := &table_models.CellInt{
			Name: "ID",
			Value: int(models.GetOrderPointerGongstruct(
				probe.stageOfInterest,
				structInstance,
			)),
		}
		cell.CellInt = cellInt

		cell = &table_models.Cell{
			Name: "Delete Icon",
		}
		row.Cells = append(row.Cells, cell)
		cellIcon := &table_models.CellIcon{
			Name: fmt.Sprintf("Delete Icon %d", models.GetOrderPointerGongstruct(
				probe.stageOfInterest,
				structInstance,
			)),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm tou want to delete this instance ?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, cellIcon, updatedCellIcon *table_models.CellIcon) {
				structInstance.UnstageVoid(probe.stageOfInterest)

				// after a delete of an instance, the stage might be dirty if a pointer or a slice of pointer
				// reference the deleted instance.
				// therefore, it is mandatory to clean the stage of interest
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()

				updateProbeTable[T](probe)
				probe.ux_tree()
			},
		}
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValueFromPointer(structInstance, fieldName.Name, probe.stageOfInterest)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value.GetValueString()
			fieldIndex++
			// log.Println(fieldName, value)
			cell := &table_models.Cell{
				Name: name,
			}
			row.Cells = append(row.Cells, cell)

			switch value.GongFieldValueType {
			case models.GongFieldValueTypeInt:
				cellInt := &table_models.CellInt{
					Name:  name,
					Value: value.GetValueInt(),
				}
				cell.CellInt = cellInt
			case models.GongFieldValueTypeFloat:
				cellFloat := &table_models.CellFloat64{
					Name:  name,
					Value: value.GetValueFloat(),
				}
				cell.CellFloat64 = cellFloat
			case models.GongFieldValueTypeBool:
				cellBool := &table_models.CellBoolean{
					Name:  name,
					Value: value.GetValueBool(),
				}
				cell.CellBool = cellBool
			default:
				cellString := &table_models.CellString{
					Name:  name,
					Value: value.GetValueString(),
				}
				cell.CellString = cellString

			}
		}
		for _, reverseField := range reverseFields {

			value := structInstance.GongGetReverseFieldOwnerName(
				probe.stageOfInterest,
				&reverseField)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := &table_models.Cell{
				Name: name,
			}
			row.Cells = append(row.Cells, cell)

			cellString := &table_models.CellString{
				Name:  name,
				Value: value,
			}
			cell.CellString = cellString
		}
	}

	table_models.StageBranch(probe.tableStage, table)

	probe.tableStage.Commit()
}

func NewRowUpdate[T models.PointerToGongstruct](
	Instance T,
	probe *Probe,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.probe = probe
	return
}

type RowUpdate[T models.PointerToGongstruct] struct {
	Instance T
	probe    *Probe
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *table_models.Stage, row, updatedRow *table_models.Row) {
	// log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}

func (probe *Probe) UpdateAndCommitNotificationTable() {
	probe.notificationTableStage.Reset()

	table := new(table_models.Table)
	table.Name = TableName
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	notificationsResetButton := &table_models.Button{
		Name:            "NotificationsResetButton",
		Icon:            string(tree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Reset notification table",
		ToolTipPosition: table_models.Below,
		OnClick: func() {
			probe.ResetNotifications()
		},
	}
	table.Buttons = append(table.Buttons, notificationsResetButton)

	column := new(table_models.DisplayedColumn)
	column.Name = "Date"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(table_models.DisplayedColumn)
	column.Name = "Message"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	// sort notifications by date
	sort.Slice(probe.notification, func(i, j int) bool {
		return probe.notification[i].Date.After(probe.notification[j].Date)
	})

	for _, notification := range probe.notification {
		row := new(table_models.Row)
		table.Rows = append(table.Rows, row)

		{
			cell := new(table_models.Cell)
			cellString := new(table_models.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Date.Format(time.StampMicro)
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := new(table_models.Cell)
			cellString := new(table_models.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Message
			row.Cells = append(row.Cells, cell)
		}
	}

	table_models.StageBranch(probe.notificationTableStage, table)

	probe.notificationTableStage.Commit()
}
