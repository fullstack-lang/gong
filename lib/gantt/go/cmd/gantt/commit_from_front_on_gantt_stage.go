package main

import (
	"fmt"
	"log"
	"os"

	gantt_models "github.com/fullstack-lang/gong/lib/gantt/go/models"

	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// hook marhalling to stage
type CommitFromFrontOnGanttStage struct {
	gongsvgStage   *svg_models.Stage
	ganttSVGMapper *gantt_models.GanttSVGMapper
}

// BeforeCommit meets the interface for the commit on the gantt stage
// It performs 2 tasks
// 1 - update the SVG stack
// 2 - persists the data to the gantt file
func (beforeCommitFromFrontOnGanttStage *CommitFromFrontOnGanttStage) BeforeCommit(
	ganttStage *gantt_models.Stage) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	// update the gantt stage with the back repo data that was updated from the front
	ganttStage.Checkout()

	// marshall to the file
	ganttStage.Marshall(file, "github.com/fullstack-lang/gong/lib/gantt/go/models", "main")
	beforeCommitFromFrontOnGanttStage.ganttSVGMapper.GenerateSvg(ganttStage, beforeCommitFromFrontOnGanttStage.gongsvgStage)
}
