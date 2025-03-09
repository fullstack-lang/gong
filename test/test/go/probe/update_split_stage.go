package probe

import (
	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

func updateSplitStage(probe *Probe) {

	probe.splitStage.Reset()

	mainView := (&split.View{
		Name: "Main view",
	}).Stage(probe.splitStage)

	topSplitArea := (&split.AsSplitArea{
		Name: "Top",
		Size: 50,
	}).Stage(probe.splitStage)
	mainView.RootAsSplitAreas = append(mainView.RootAsSplitAreas, topSplitArea)

	horizontalSplit := (&split.AsSplit{
		Name:      "Top, sidebar, table & form",
		Direction: split.Horizontal,
	}).Stage(probe.splitStage)
	topSplitArea.AsSplits = append(topSplitArea.AsSplits, horizontalSplit)

	sidebarArea := (&split.AsSplitArea{
		Name: "sidebar tree",
		Size: 20,
	}).Stage(probe.splitStage)
	horizontalSplit.AsSplitAreas = append(horizontalSplit.AsSplitAreas, sidebarArea)

	tree := (&split.Tree{
		Name:      "Sidebar",
		StackName: probe.stackPath + ProbeTreeSidebarSuffix,
		TreeName:  SideBarTreeName,
	}).Stage(probe.splitStage)
	sidebarArea.Tree = tree

	tableArea := (&split.AsSplitArea{
		Name: "table",
		Size: 50,
	}).Stage(probe.splitStage)
	horizontalSplit.AsSplitAreas = append(horizontalSplit.AsSplitAreas, tableArea)

	table := (&split.Table{
		Name:      "Table",
		StackName: probe.stackPath + ProbeTableSuffix,
		TableName: TableName,
	}).Stage(probe.splitStage)
	tableArea.Table = table

	formArea := (&split.AsSplitArea{
		Name: "form",
		Size: 30,
	}).Stage(probe.splitStage)
	horizontalSplit.AsSplitAreas = append(horizontalSplit.AsSplitAreas, formArea)

	form := (&split.Form{
		Name:      "Form",
		StackName: probe.stackPath + ProbeFormSuffix,
		FormName:  FormName,
	}).Stage(probe.splitStage)
	formArea.Form = form

	bottomSplitArea := (&split.AsSplitArea{
		Name: "Bottom",
		Size: 50,
	}).Stage(probe.splitStage)
	mainView.RootAsSplitAreas = append(mainView.RootAsSplitAreas, bottomSplitArea)

	doc := (&split.Doc{
		Name:      "Doc",
		StackName: probe.stageOfInterest.GetType(),
	}).Stage(probe.splitStage)
	bottomSplitArea.Doc = doc

	probe.splitStage.Commit()
}
