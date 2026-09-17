// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (assplit *AsSplit) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAsSplitCreateCallback != nil {
		stage.OnAfterAsSplitCreateCallback.OnAfterCreate(stage, assplit)
	}
}

func (assplit *AsSplit) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAsSplitUpdateCallback != nil {
		var frontAsSplit *AsSplit
		if front != nil {
			frontAsSplit, _ = front.(*AsSplit)
		}
		stage.OnAfterAsSplitUpdateCallback.OnAfterUpdate(stage, assplit, frontAsSplit)
	}
}

func (assplit *AsSplit) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAsSplitDeleteCallback != nil {
		var frontAsSplit *AsSplit
		if front != nil {
			frontAsSplit, _ = front.(*AsSplit)
		}
		stage.OnAfterAsSplitDeleteCallback.OnAfterDelete(stage, assplit, frontAsSplit)
	}
}

func (assplitarea *AsSplitArea) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAsSplitAreaCreateCallback != nil {
		stage.OnAfterAsSplitAreaCreateCallback.OnAfterCreate(stage, assplitarea)
	}
}

func (assplitarea *AsSplitArea) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAsSplitAreaUpdateCallback != nil {
		var frontAsSplitArea *AsSplitArea
		if front != nil {
			frontAsSplitArea, _ = front.(*AsSplitArea)
		}
		stage.OnAfterAsSplitAreaUpdateCallback.OnAfterUpdate(stage, assplitarea, frontAsSplitArea)
	}
}

func (assplitarea *AsSplitArea) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAsSplitAreaDeleteCallback != nil {
		var frontAsSplitArea *AsSplitArea
		if front != nil {
			frontAsSplitArea, _ = front.(*AsSplitArea)
		}
		stage.OnAfterAsSplitAreaDeleteCallback.OnAfterDelete(stage, assplitarea, frontAsSplitArea)
	}
}

func (button *Button) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterButtonCreateCallback != nil {
		stage.OnAfterButtonCreateCallback.OnAfterCreate(stage, button)
	}
}

func (button *Button) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonUpdateCallback != nil {
		var frontButton *Button
		if front != nil {
			frontButton, _ = front.(*Button)
		}
		stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, button, frontButton)
	}
}

func (button *Button) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonDeleteCallback != nil {
		var frontButton *Button
		if front != nil {
			frontButton, _ = front.(*Button)
		}
		stage.OnAfterButtonDeleteCallback.OnAfterDelete(stage, button, frontButton)
	}
}

func (cursor *Cursor) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCursorCreateCallback != nil {
		stage.OnAfterCursorCreateCallback.OnAfterCreate(stage, cursor)
	}
}

func (cursor *Cursor) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCursorUpdateCallback != nil {
		var frontCursor *Cursor
		if front != nil {
			frontCursor, _ = front.(*Cursor)
		}
		stage.OnAfterCursorUpdateCallback.OnAfterUpdate(stage, cursor, frontCursor)
	}
}

func (cursor *Cursor) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCursorDeleteCallback != nil {
		var frontCursor *Cursor
		if front != nil {
			frontCursor, _ = front.(*Cursor)
		}
		stage.OnAfterCursorDeleteCallback.OnAfterDelete(stage, cursor, frontCursor)
	}
}

func (favicon *FavIcon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFavIconCreateCallback != nil {
		stage.OnAfterFavIconCreateCallback.OnAfterCreate(stage, favicon)
	}
}

func (favicon *FavIcon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFavIconUpdateCallback != nil {
		var frontFavIcon *FavIcon
		if front != nil {
			frontFavIcon, _ = front.(*FavIcon)
		}
		stage.OnAfterFavIconUpdateCallback.OnAfterUpdate(stage, favicon, frontFavIcon)
	}
}

func (favicon *FavIcon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFavIconDeleteCallback != nil {
		var frontFavIcon *FavIcon
		if front != nil {
			frontFavIcon, _ = front.(*FavIcon)
		}
		stage.OnAfterFavIconDeleteCallback.OnAfterDelete(stage, favicon, frontFavIcon)
	}
}

func (form *Form) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormCreateCallback != nil {
		stage.OnAfterFormCreateCallback.OnAfterCreate(stage, form)
	}
}

func (form *Form) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormUpdateCallback != nil {
		var frontForm *Form
		if front != nil {
			frontForm, _ = front.(*Form)
		}
		stage.OnAfterFormUpdateCallback.OnAfterUpdate(stage, form, frontForm)
	}
}

func (form *Form) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormDeleteCallback != nil {
		var frontForm *Form
		if front != nil {
			frontForm, _ = front.(*Form)
		}
		stage.OnAfterFormDeleteCallback.OnAfterDelete(stage, form, frontForm)
	}
}

func (load *Load) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLoadCreateCallback != nil {
		stage.OnAfterLoadCreateCallback.OnAfterCreate(stage, load)
	}
}

func (load *Load) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLoadUpdateCallback != nil {
		var frontLoad *Load
		if front != nil {
			frontLoad, _ = front.(*Load)
		}
		stage.OnAfterLoadUpdateCallback.OnAfterUpdate(stage, load, frontLoad)
	}
}

func (load *Load) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLoadDeleteCallback != nil {
		var frontLoad *Load
		if front != nil {
			frontLoad, _ = front.(*Load)
		}
		stage.OnAfterLoadDeleteCallback.OnAfterDelete(stage, load, frontLoad)
	}
}

func (logoontheleft *LogoOnTheLeft) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLogoOnTheLeftCreateCallback != nil {
		stage.OnAfterLogoOnTheLeftCreateCallback.OnAfterCreate(stage, logoontheleft)
	}
}

func (logoontheleft *LogoOnTheLeft) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLogoOnTheLeftUpdateCallback != nil {
		var frontLogoOnTheLeft *LogoOnTheLeft
		if front != nil {
			frontLogoOnTheLeft, _ = front.(*LogoOnTheLeft)
		}
		stage.OnAfterLogoOnTheLeftUpdateCallback.OnAfterUpdate(stage, logoontheleft, frontLogoOnTheLeft)
	}
}

func (logoontheleft *LogoOnTheLeft) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLogoOnTheLeftDeleteCallback != nil {
		var frontLogoOnTheLeft *LogoOnTheLeft
		if front != nil {
			frontLogoOnTheLeft, _ = front.(*LogoOnTheLeft)
		}
		stage.OnAfterLogoOnTheLeftDeleteCallback.OnAfterDelete(stage, logoontheleft, frontLogoOnTheLeft)
	}
}

func (logoontheright *LogoOnTheRight) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLogoOnTheRightCreateCallback != nil {
		stage.OnAfterLogoOnTheRightCreateCallback.OnAfterCreate(stage, logoontheright)
	}
}

func (logoontheright *LogoOnTheRight) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLogoOnTheRightUpdateCallback != nil {
		var frontLogoOnTheRight *LogoOnTheRight
		if front != nil {
			frontLogoOnTheRight, _ = front.(*LogoOnTheRight)
		}
		stage.OnAfterLogoOnTheRightUpdateCallback.OnAfterUpdate(stage, logoontheright, frontLogoOnTheRight)
	}
}

func (logoontheright *LogoOnTheRight) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLogoOnTheRightDeleteCallback != nil {
		var frontLogoOnTheRight *LogoOnTheRight
		if front != nil {
			frontLogoOnTheRight, _ = front.(*LogoOnTheRight)
		}
		stage.OnAfterLogoOnTheRightDeleteCallback.OnAfterDelete(stage, logoontheright, frontLogoOnTheRight)
	}
}

func (markdown *Markdown) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkdownCreateCallback != nil {
		stage.OnAfterMarkdownCreateCallback.OnAfterCreate(stage, markdown)
	}
}

func (markdown *Markdown) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkdownUpdateCallback != nil {
		var frontMarkdown *Markdown
		if front != nil {
			frontMarkdown, _ = front.(*Markdown)
		}
		stage.OnAfterMarkdownUpdateCallback.OnAfterUpdate(stage, markdown, frontMarkdown)
	}
}

func (markdown *Markdown) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkdownDeleteCallback != nil {
		var frontMarkdown *Markdown
		if front != nil {
			frontMarkdown, _ = front.(*Markdown)
		}
		stage.OnAfterMarkdownDeleteCallback.OnAfterDelete(stage, markdown, frontMarkdown)
	}
}

func (slider *Slider) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSliderCreateCallback != nil {
		stage.OnAfterSliderCreateCallback.OnAfterCreate(stage, slider)
	}
}

func (slider *Slider) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSliderUpdateCallback != nil {
		var frontSlider *Slider
		if front != nil {
			frontSlider, _ = front.(*Slider)
		}
		stage.OnAfterSliderUpdateCallback.OnAfterUpdate(stage, slider, frontSlider)
	}
}

func (slider *Slider) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSliderDeleteCallback != nil {
		var frontSlider *Slider
		if front != nil {
			frontSlider, _ = front.(*Slider)
		}
		stage.OnAfterSliderDeleteCallback.OnAfterDelete(stage, slider, frontSlider)
	}
}

func (split *Split) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSplitCreateCallback != nil {
		stage.OnAfterSplitCreateCallback.OnAfterCreate(stage, split)
	}
}

func (split *Split) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSplitUpdateCallback != nil {
		var frontSplit *Split
		if front != nil {
			frontSplit, _ = front.(*Split)
		}
		stage.OnAfterSplitUpdateCallback.OnAfterUpdate(stage, split, frontSplit)
	}
}

func (split *Split) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSplitDeleteCallback != nil {
		var frontSplit *Split
		if front != nil {
			frontSplit, _ = front.(*Split)
		}
		stage.OnAfterSplitDeleteCallback.OnAfterDelete(stage, split, frontSplit)
	}
}

func (svg *Svg) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSvgCreateCallback != nil {
		stage.OnAfterSvgCreateCallback.OnAfterCreate(stage, svg)
	}
}

func (svg *Svg) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSvgUpdateCallback != nil {
		var frontSvg *Svg
		if front != nil {
			frontSvg, _ = front.(*Svg)
		}
		stage.OnAfterSvgUpdateCallback.OnAfterUpdate(stage, svg, frontSvg)
	}
}

func (svg *Svg) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSvgDeleteCallback != nil {
		var frontSvg *Svg
		if front != nil {
			frontSvg, _ = front.(*Svg)
		}
		stage.OnAfterSvgDeleteCallback.OnAfterDelete(stage, svg, frontSvg)
	}
}

func (table *Table) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTableCreateCallback != nil {
		stage.OnAfterTableCreateCallback.OnAfterCreate(stage, table)
	}
}

func (table *Table) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableUpdateCallback != nil {
		var frontTable *Table
		if front != nil {
			frontTable, _ = front.(*Table)
		}
		stage.OnAfterTableUpdateCallback.OnAfterUpdate(stage, table, frontTable)
	}
}

func (table *Table) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableDeleteCallback != nil {
		var frontTable *Table
		if front != nil {
			frontTable, _ = front.(*Table)
		}
		stage.OnAfterTableDeleteCallback.OnAfterDelete(stage, table, frontTable)
	}
}

func (threejs *Threejs) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterThreejsCreateCallback != nil {
		stage.OnAfterThreejsCreateCallback.OnAfterCreate(stage, threejs)
	}
}

func (threejs *Threejs) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterThreejsUpdateCallback != nil {
		var frontThreejs *Threejs
		if front != nil {
			frontThreejs, _ = front.(*Threejs)
		}
		stage.OnAfterThreejsUpdateCallback.OnAfterUpdate(stage, threejs, frontThreejs)
	}
}

func (threejs *Threejs) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterThreejsDeleteCallback != nil {
		var frontThreejs *Threejs
		if front != nil {
			frontThreejs, _ = front.(*Threejs)
		}
		stage.OnAfterThreejsDeleteCallback.OnAfterDelete(stage, threejs, frontThreejs)
	}
}

func (title *Title) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTitleCreateCallback != nil {
		stage.OnAfterTitleCreateCallback.OnAfterCreate(stage, title)
	}
}

func (title *Title) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTitleUpdateCallback != nil {
		var frontTitle *Title
		if front != nil {
			frontTitle, _ = front.(*Title)
		}
		stage.OnAfterTitleUpdateCallback.OnAfterUpdate(stage, title, frontTitle)
	}
}

func (title *Title) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTitleDeleteCallback != nil {
		var frontTitle *Title
		if front != nil {
			frontTitle, _ = front.(*Title)
		}
		stage.OnAfterTitleDeleteCallback.OnAfterDelete(stage, title, frontTitle)
	}
}

func (tone *Tone) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterToneCreateCallback != nil {
		stage.OnAfterToneCreateCallback.OnAfterCreate(stage, tone)
	}
}

func (tone *Tone) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterToneUpdateCallback != nil {
		var frontTone *Tone
		if front != nil {
			frontTone, _ = front.(*Tone)
		}
		stage.OnAfterToneUpdateCallback.OnAfterUpdate(stage, tone, frontTone)
	}
}

func (tone *Tone) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterToneDeleteCallback != nil {
		var frontTone *Tone
		if front != nil {
			frontTone, _ = front.(*Tone)
		}
		stage.OnAfterToneDeleteCallback.OnAfterDelete(stage, tone, frontTone)
	}
}

func (tree *Tree) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTreeCreateCallback != nil {
		stage.OnAfterTreeCreateCallback.OnAfterCreate(stage, tree)
	}
}

func (tree *Tree) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTreeUpdateCallback != nil {
		var frontTree *Tree
		if front != nil {
			frontTree, _ = front.(*Tree)
		}
		stage.OnAfterTreeUpdateCallback.OnAfterUpdate(stage, tree, frontTree)
	}
}

func (tree *Tree) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTreeDeleteCallback != nil {
		var frontTree *Tree
		if front != nil {
			frontTree, _ = front.(*Tree)
		}
		stage.OnAfterTreeDeleteCallback.OnAfterDelete(stage, tree, frontTree)
	}
}

func (view *View) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterViewCreateCallback != nil {
		stage.OnAfterViewCreateCallback.OnAfterCreate(stage, view)
	}
}

func (view *View) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterViewUpdateCallback != nil {
		var frontView *View
		if front != nil {
			frontView, _ = front.(*View)
		}
		stage.OnAfterViewUpdateCallback.OnAfterUpdate(stage, view, frontView)
	}
}

func (view *View) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterViewDeleteCallback != nil {
		var frontView *View
		if front != nil {
			frontView, _ = front.(*View)
		}
		stage.OnAfterViewDeleteCallback.OnAfterDelete(stage, view, frontView)
	}
}

func (xlsx *Xlsx) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXlsxCreateCallback != nil {
		stage.OnAfterXlsxCreateCallback.OnAfterCreate(stage, xlsx)
	}
}

func (xlsx *Xlsx) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXlsxUpdateCallback != nil {
		var frontXlsx *Xlsx
		if front != nil {
			frontXlsx, _ = front.(*Xlsx)
		}
		stage.OnAfterXlsxUpdateCallback.OnAfterUpdate(stage, xlsx, frontXlsx)
	}
}

func (xlsx *Xlsx) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXlsxDeleteCallback != nil {
		var frontXlsx *Xlsx
		if front != nil {
			frontXlsx, _ = front.(*Xlsx)
		}
		stage.OnAfterXlsxDeleteCallback.OnAfterDelete(stage, xlsx, frontXlsx)
	}
}

