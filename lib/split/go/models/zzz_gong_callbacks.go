// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AsSplit:
		if stage.OnAfterAsSplitCreateCallback != nil {
			stage.OnAfterAsSplitCreateCallback.OnAfterCreate(stage, target)
		}
	case *AsSplitArea:
		if stage.OnAfterAsSplitAreaCreateCallback != nil {
			stage.OnAfterAsSplitAreaCreateCallback.OnAfterCreate(stage, target)
		}
	case *Button:
		if stage.OnAfterButtonCreateCallback != nil {
			stage.OnAfterButtonCreateCallback.OnAfterCreate(stage, target)
		}
	case *Cursor:
		if stage.OnAfterCursorCreateCallback != nil {
			stage.OnAfterCursorCreateCallback.OnAfterCreate(stage, target)
		}
	case *FavIcon:
		if stage.OnAfterFavIconCreateCallback != nil {
			stage.OnAfterFavIconCreateCallback.OnAfterCreate(stage, target)
		}
	case *Form:
		if stage.OnAfterFormCreateCallback != nil {
			stage.OnAfterFormCreateCallback.OnAfterCreate(stage, target)
		}
	case *Load:
		if stage.OnAfterLoadCreateCallback != nil {
			stage.OnAfterLoadCreateCallback.OnAfterCreate(stage, target)
		}
	case *LogoOnTheLeft:
		if stage.OnAfterLogoOnTheLeftCreateCallback != nil {
			stage.OnAfterLogoOnTheLeftCreateCallback.OnAfterCreate(stage, target)
		}
	case *LogoOnTheRight:
		if stage.OnAfterLogoOnTheRightCreateCallback != nil {
			stage.OnAfterLogoOnTheRightCreateCallback.OnAfterCreate(stage, target)
		}
	case *Markdown:
		if stage.OnAfterMarkdownCreateCallback != nil {
			stage.OnAfterMarkdownCreateCallback.OnAfterCreate(stage, target)
		}
	case *Slider:
		if stage.OnAfterSliderCreateCallback != nil {
			stage.OnAfterSliderCreateCallback.OnAfterCreate(stage, target)
		}
	case *Split:
		if stage.OnAfterSplitCreateCallback != nil {
			stage.OnAfterSplitCreateCallback.OnAfterCreate(stage, target)
		}
	case *Svg:
		if stage.OnAfterSvgCreateCallback != nil {
			stage.OnAfterSvgCreateCallback.OnAfterCreate(stage, target)
		}
	case *Table:
		if stage.OnAfterTableCreateCallback != nil {
			stage.OnAfterTableCreateCallback.OnAfterCreate(stage, target)
		}
	case *Title:
		if stage.OnAfterTitleCreateCallback != nil {
			stage.OnAfterTitleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tone:
		if stage.OnAfterToneCreateCallback != nil {
			stage.OnAfterToneCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeCreateCallback != nil {
			stage.OnAfterTreeCreateCallback.OnAfterCreate(stage, target)
		}
	case *View:
		if stage.OnAfterViewCreateCallback != nil {
			stage.OnAfterViewCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xlsx:
		if stage.OnAfterXlsxCreateCallback != nil {
			stage.OnAfterXlsxCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type, mouseEvent *Gong__MouseEvent) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AsSplit:
		newTarget := any(new).(*AsSplit)
		if stage.OnAfterAsSplitUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterAsSplitUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterAsSplitUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterAsSplitUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *AsSplitArea:
		newTarget := any(new).(*AsSplitArea)
		if stage.OnAfterAsSplitAreaUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterAsSplitAreaUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterAsSplitAreaUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterAsSplitAreaUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Button:
		newTarget := any(new).(*Button)
		if stage.OnAfterButtonUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterButtonUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterButtonUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Cursor:
		newTarget := any(new).(*Cursor)
		if stage.OnAfterCursorUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterCursorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterCursorUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterCursorUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *FavIcon:
		newTarget := any(new).(*FavIcon)
		if stage.OnAfterFavIconUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterFavIconUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterFavIconUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterFavIconUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Form:
		newTarget := any(new).(*Form)
		if stage.OnAfterFormUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterFormUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterFormUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterFormUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Load:
		newTarget := any(new).(*Load)
		if stage.OnAfterLoadUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterLoadUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterLoadUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterLoadUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *LogoOnTheLeft:
		newTarget := any(new).(*LogoOnTheLeft)
		if stage.OnAfterLogoOnTheLeftUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterLogoOnTheLeftUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterLogoOnTheLeftUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterLogoOnTheLeftUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *LogoOnTheRight:
		newTarget := any(new).(*LogoOnTheRight)
		if stage.OnAfterLogoOnTheRightUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterLogoOnTheRightUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterLogoOnTheRightUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterLogoOnTheRightUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Markdown:
		newTarget := any(new).(*Markdown)
		if stage.OnAfterMarkdownUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMarkdownUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMarkdownUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMarkdownUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Slider:
		newTarget := any(new).(*Slider)
		if stage.OnAfterSliderUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSliderUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSliderUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSliderUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Split:
		newTarget := any(new).(*Split)
		if stage.OnAfterSplitUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSplitUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSplitUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSplitUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Svg:
		newTarget := any(new).(*Svg)
		if stage.OnAfterSvgUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSvgUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSvgUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSvgUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Table:
		newTarget := any(new).(*Table)
		if stage.OnAfterTableUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterTableUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterTableUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterTableUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Title:
		newTarget := any(new).(*Title)
		if stage.OnAfterTitleUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterTitleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterTitleUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterTitleUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Tone:
		newTarget := any(new).(*Tone)
		if stage.OnAfterToneUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterToneUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterToneUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterToneUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Tree:
		newTarget := any(new).(*Tree)
		if stage.OnAfterTreeUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterTreeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterTreeUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterTreeUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *View:
		newTarget := any(new).(*View)
		if stage.OnAfterViewUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterViewUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterViewUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterViewUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Xlsx:
		newTarget := any(new).(*Xlsx)
		if stage.OnAfterXlsxUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterXlsxUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterXlsxUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterXlsxUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *AsSplit:
		if stage.OnAfterAsSplitDeleteCallback != nil {
			staged := any(staged).(*AsSplit)
			stage.OnAfterAsSplitDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *AsSplitArea:
		if stage.OnAfterAsSplitAreaDeleteCallback != nil {
			staged := any(staged).(*AsSplitArea)
			stage.OnAfterAsSplitAreaDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Button:
		if stage.OnAfterButtonDeleteCallback != nil {
			staged := any(staged).(*Button)
			stage.OnAfterButtonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Cursor:
		if stage.OnAfterCursorDeleteCallback != nil {
			staged := any(staged).(*Cursor)
			stage.OnAfterCursorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FavIcon:
		if stage.OnAfterFavIconDeleteCallback != nil {
			staged := any(staged).(*FavIcon)
			stage.OnAfterFavIconDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Form:
		if stage.OnAfterFormDeleteCallback != nil {
			staged := any(staged).(*Form)
			stage.OnAfterFormDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Load:
		if stage.OnAfterLoadDeleteCallback != nil {
			staged := any(staged).(*Load)
			stage.OnAfterLoadDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LogoOnTheLeft:
		if stage.OnAfterLogoOnTheLeftDeleteCallback != nil {
			staged := any(staged).(*LogoOnTheLeft)
			stage.OnAfterLogoOnTheLeftDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LogoOnTheRight:
		if stage.OnAfterLogoOnTheRightDeleteCallback != nil {
			staged := any(staged).(*LogoOnTheRight)
			stage.OnAfterLogoOnTheRightDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Markdown:
		if stage.OnAfterMarkdownDeleteCallback != nil {
			staged := any(staged).(*Markdown)
			stage.OnAfterMarkdownDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Slider:
		if stage.OnAfterSliderDeleteCallback != nil {
			staged := any(staged).(*Slider)
			stage.OnAfterSliderDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Split:
		if stage.OnAfterSplitDeleteCallback != nil {
			staged := any(staged).(*Split)
			stage.OnAfterSplitDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Svg:
		if stage.OnAfterSvgDeleteCallback != nil {
			staged := any(staged).(*Svg)
			stage.OnAfterSvgDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Table:
		if stage.OnAfterTableDeleteCallback != nil {
			staged := any(staged).(*Table)
			stage.OnAfterTableDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Title:
		if stage.OnAfterTitleDeleteCallback != nil {
			staged := any(staged).(*Title)
			stage.OnAfterTitleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tone:
		if stage.OnAfterToneDeleteCallback != nil {
			staged := any(staged).(*Tone)
			stage.OnAfterToneDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tree:
		if stage.OnAfterTreeDeleteCallback != nil {
			staged := any(staged).(*Tree)
			stage.OnAfterTreeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *View:
		if stage.OnAfterViewDeleteCallback != nil {
			staged := any(staged).(*View)
			stage.OnAfterViewDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xlsx:
		if stage.OnAfterXlsxDeleteCallback != nil {
			staged := any(staged).(*Xlsx)
			stage.OnAfterXlsxDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AsSplit:
		if stage.OnAfterAsSplitReadCallback != nil {
			stage.OnAfterAsSplitReadCallback.OnAfterRead(stage, target)
		}
	case *AsSplitArea:
		if stage.OnAfterAsSplitAreaReadCallback != nil {
			stage.OnAfterAsSplitAreaReadCallback.OnAfterRead(stage, target)
		}
	case *Button:
		if stage.OnAfterButtonReadCallback != nil {
			stage.OnAfterButtonReadCallback.OnAfterRead(stage, target)
		}
	case *Cursor:
		if stage.OnAfterCursorReadCallback != nil {
			stage.OnAfterCursorReadCallback.OnAfterRead(stage, target)
		}
	case *FavIcon:
		if stage.OnAfterFavIconReadCallback != nil {
			stage.OnAfterFavIconReadCallback.OnAfterRead(stage, target)
		}
	case *Form:
		if stage.OnAfterFormReadCallback != nil {
			stage.OnAfterFormReadCallback.OnAfterRead(stage, target)
		}
	case *Load:
		if stage.OnAfterLoadReadCallback != nil {
			stage.OnAfterLoadReadCallback.OnAfterRead(stage, target)
		}
	case *LogoOnTheLeft:
		if stage.OnAfterLogoOnTheLeftReadCallback != nil {
			stage.OnAfterLogoOnTheLeftReadCallback.OnAfterRead(stage, target)
		}
	case *LogoOnTheRight:
		if stage.OnAfterLogoOnTheRightReadCallback != nil {
			stage.OnAfterLogoOnTheRightReadCallback.OnAfterRead(stage, target)
		}
	case *Markdown:
		if stage.OnAfterMarkdownReadCallback != nil {
			stage.OnAfterMarkdownReadCallback.OnAfterRead(stage, target)
		}
	case *Slider:
		if stage.OnAfterSliderReadCallback != nil {
			stage.OnAfterSliderReadCallback.OnAfterRead(stage, target)
		}
	case *Split:
		if stage.OnAfterSplitReadCallback != nil {
			stage.OnAfterSplitReadCallback.OnAfterRead(stage, target)
		}
	case *Svg:
		if stage.OnAfterSvgReadCallback != nil {
			stage.OnAfterSvgReadCallback.OnAfterRead(stage, target)
		}
	case *Table:
		if stage.OnAfterTableReadCallback != nil {
			stage.OnAfterTableReadCallback.OnAfterRead(stage, target)
		}
	case *Title:
		if stage.OnAfterTitleReadCallback != nil {
			stage.OnAfterTitleReadCallback.OnAfterRead(stage, target)
		}
	case *Tone:
		if stage.OnAfterToneReadCallback != nil {
			stage.OnAfterToneReadCallback.OnAfterRead(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeReadCallback != nil {
			stage.OnAfterTreeReadCallback.OnAfterRead(stage, target)
		}
	case *View:
		if stage.OnAfterViewReadCallback != nil {
			stage.OnAfterViewReadCallback.OnAfterRead(stage, target)
		}
	case *Xlsx:
		if stage.OnAfterXlsxReadCallback != nil {
			stage.OnAfterXlsxReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AsSplit:
		stage.OnAfterAsSplitUpdateCallback = any(callback).(OnAfterUpdateInterface[AsSplit])
	
	case *AsSplitArea:
		stage.OnAfterAsSplitAreaUpdateCallback = any(callback).(OnAfterUpdateInterface[AsSplitArea])
	
	case *Button:
		stage.OnAfterButtonUpdateCallback = any(callback).(OnAfterUpdateInterface[Button])
	
	case *Cursor:
		stage.OnAfterCursorUpdateCallback = any(callback).(OnAfterUpdateInterface[Cursor])
	
	case *FavIcon:
		stage.OnAfterFavIconUpdateCallback = any(callback).(OnAfterUpdateInterface[FavIcon])
	
	case *Form:
		stage.OnAfterFormUpdateCallback = any(callback).(OnAfterUpdateInterface[Form])
	
	case *Load:
		stage.OnAfterLoadUpdateCallback = any(callback).(OnAfterUpdateInterface[Load])
	
	case *LogoOnTheLeft:
		stage.OnAfterLogoOnTheLeftUpdateCallback = any(callback).(OnAfterUpdateInterface[LogoOnTheLeft])
	
	case *LogoOnTheRight:
		stage.OnAfterLogoOnTheRightUpdateCallback = any(callback).(OnAfterUpdateInterface[LogoOnTheRight])
	
	case *Markdown:
		stage.OnAfterMarkdownUpdateCallback = any(callback).(OnAfterUpdateInterface[Markdown])
	
	case *Slider:
		stage.OnAfterSliderUpdateCallback = any(callback).(OnAfterUpdateInterface[Slider])
	
	case *Split:
		stage.OnAfterSplitUpdateCallback = any(callback).(OnAfterUpdateInterface[Split])
	
	case *Svg:
		stage.OnAfterSvgUpdateCallback = any(callback).(OnAfterUpdateInterface[Svg])
	
	case *Table:
		stage.OnAfterTableUpdateCallback = any(callback).(OnAfterUpdateInterface[Table])
	
	case *Title:
		stage.OnAfterTitleUpdateCallback = any(callback).(OnAfterUpdateInterface[Title])
	
	case *Tone:
		stage.OnAfterToneUpdateCallback = any(callback).(OnAfterUpdateInterface[Tone])
	
	case *Tree:
		stage.OnAfterTreeUpdateCallback = any(callback).(OnAfterUpdateInterface[Tree])
	
	case *View:
		stage.OnAfterViewUpdateCallback = any(callback).(OnAfterUpdateInterface[View])
	
	case *Xlsx:
		stage.OnAfterXlsxUpdateCallback = any(callback).(OnAfterUpdateInterface[Xlsx])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AsSplit:
		stage.OnAfterAsSplitCreateCallback = any(callback).(OnAfterCreateInterface[AsSplit])
	
	case *AsSplitArea:
		stage.OnAfterAsSplitAreaCreateCallback = any(callback).(OnAfterCreateInterface[AsSplitArea])
	
	case *Button:
		stage.OnAfterButtonCreateCallback = any(callback).(OnAfterCreateInterface[Button])
	
	case *Cursor:
		stage.OnAfterCursorCreateCallback = any(callback).(OnAfterCreateInterface[Cursor])
	
	case *FavIcon:
		stage.OnAfterFavIconCreateCallback = any(callback).(OnAfterCreateInterface[FavIcon])
	
	case *Form:
		stage.OnAfterFormCreateCallback = any(callback).(OnAfterCreateInterface[Form])
	
	case *Load:
		stage.OnAfterLoadCreateCallback = any(callback).(OnAfterCreateInterface[Load])
	
	case *LogoOnTheLeft:
		stage.OnAfterLogoOnTheLeftCreateCallback = any(callback).(OnAfterCreateInterface[LogoOnTheLeft])
	
	case *LogoOnTheRight:
		stage.OnAfterLogoOnTheRightCreateCallback = any(callback).(OnAfterCreateInterface[LogoOnTheRight])
	
	case *Markdown:
		stage.OnAfterMarkdownCreateCallback = any(callback).(OnAfterCreateInterface[Markdown])
	
	case *Slider:
		stage.OnAfterSliderCreateCallback = any(callback).(OnAfterCreateInterface[Slider])
	
	case *Split:
		stage.OnAfterSplitCreateCallback = any(callback).(OnAfterCreateInterface[Split])
	
	case *Svg:
		stage.OnAfterSvgCreateCallback = any(callback).(OnAfterCreateInterface[Svg])
	
	case *Table:
		stage.OnAfterTableCreateCallback = any(callback).(OnAfterCreateInterface[Table])
	
	case *Title:
		stage.OnAfterTitleCreateCallback = any(callback).(OnAfterCreateInterface[Title])
	
	case *Tone:
		stage.OnAfterToneCreateCallback = any(callback).(OnAfterCreateInterface[Tone])
	
	case *Tree:
		stage.OnAfterTreeCreateCallback = any(callback).(OnAfterCreateInterface[Tree])
	
	case *View:
		stage.OnAfterViewCreateCallback = any(callback).(OnAfterCreateInterface[View])
	
	case *Xlsx:
		stage.OnAfterXlsxCreateCallback = any(callback).(OnAfterCreateInterface[Xlsx])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AsSplit:
		stage.OnAfterAsSplitDeleteCallback = any(callback).(OnAfterDeleteInterface[AsSplit])
	
	case *AsSplitArea:
		stage.OnAfterAsSplitAreaDeleteCallback = any(callback).(OnAfterDeleteInterface[AsSplitArea])
	
	case *Button:
		stage.OnAfterButtonDeleteCallback = any(callback).(OnAfterDeleteInterface[Button])
	
	case *Cursor:
		stage.OnAfterCursorDeleteCallback = any(callback).(OnAfterDeleteInterface[Cursor])
	
	case *FavIcon:
		stage.OnAfterFavIconDeleteCallback = any(callback).(OnAfterDeleteInterface[FavIcon])
	
	case *Form:
		stage.OnAfterFormDeleteCallback = any(callback).(OnAfterDeleteInterface[Form])
	
	case *Load:
		stage.OnAfterLoadDeleteCallback = any(callback).(OnAfterDeleteInterface[Load])
	
	case *LogoOnTheLeft:
		stage.OnAfterLogoOnTheLeftDeleteCallback = any(callback).(OnAfterDeleteInterface[LogoOnTheLeft])
	
	case *LogoOnTheRight:
		stage.OnAfterLogoOnTheRightDeleteCallback = any(callback).(OnAfterDeleteInterface[LogoOnTheRight])
	
	case *Markdown:
		stage.OnAfterMarkdownDeleteCallback = any(callback).(OnAfterDeleteInterface[Markdown])
	
	case *Slider:
		stage.OnAfterSliderDeleteCallback = any(callback).(OnAfterDeleteInterface[Slider])
	
	case *Split:
		stage.OnAfterSplitDeleteCallback = any(callback).(OnAfterDeleteInterface[Split])
	
	case *Svg:
		stage.OnAfterSvgDeleteCallback = any(callback).(OnAfterDeleteInterface[Svg])
	
	case *Table:
		stage.OnAfterTableDeleteCallback = any(callback).(OnAfterDeleteInterface[Table])
	
	case *Title:
		stage.OnAfterTitleDeleteCallback = any(callback).(OnAfterDeleteInterface[Title])
	
	case *Tone:
		stage.OnAfterToneDeleteCallback = any(callback).(OnAfterDeleteInterface[Tone])
	
	case *Tree:
		stage.OnAfterTreeDeleteCallback = any(callback).(OnAfterDeleteInterface[Tree])
	
	case *View:
		stage.OnAfterViewDeleteCallback = any(callback).(OnAfterDeleteInterface[View])
	
	case *Xlsx:
		stage.OnAfterXlsxDeleteCallback = any(callback).(OnAfterDeleteInterface[Xlsx])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AsSplit:
		stage.OnAfterAsSplitReadCallback = any(callback).(OnAfterReadInterface[AsSplit])
	
	case *AsSplitArea:
		stage.OnAfterAsSplitAreaReadCallback = any(callback).(OnAfterReadInterface[AsSplitArea])
	
	case *Button:
		stage.OnAfterButtonReadCallback = any(callback).(OnAfterReadInterface[Button])
	
	case *Cursor:
		stage.OnAfterCursorReadCallback = any(callback).(OnAfterReadInterface[Cursor])
	
	case *FavIcon:
		stage.OnAfterFavIconReadCallback = any(callback).(OnAfterReadInterface[FavIcon])
	
	case *Form:
		stage.OnAfterFormReadCallback = any(callback).(OnAfterReadInterface[Form])
	
	case *Load:
		stage.OnAfterLoadReadCallback = any(callback).(OnAfterReadInterface[Load])
	
	case *LogoOnTheLeft:
		stage.OnAfterLogoOnTheLeftReadCallback = any(callback).(OnAfterReadInterface[LogoOnTheLeft])
	
	case *LogoOnTheRight:
		stage.OnAfterLogoOnTheRightReadCallback = any(callback).(OnAfterReadInterface[LogoOnTheRight])
	
	case *Markdown:
		stage.OnAfterMarkdownReadCallback = any(callback).(OnAfterReadInterface[Markdown])
	
	case *Slider:
		stage.OnAfterSliderReadCallback = any(callback).(OnAfterReadInterface[Slider])
	
	case *Split:
		stage.OnAfterSplitReadCallback = any(callback).(OnAfterReadInterface[Split])
	
	case *Svg:
		stage.OnAfterSvgReadCallback = any(callback).(OnAfterReadInterface[Svg])
	
	case *Table:
		stage.OnAfterTableReadCallback = any(callback).(OnAfterReadInterface[Table])
	
	case *Title:
		stage.OnAfterTitleReadCallback = any(callback).(OnAfterReadInterface[Title])
	
	case *Tone:
		stage.OnAfterToneReadCallback = any(callback).(OnAfterReadInterface[Tone])
	
	case *Tree:
		stage.OnAfterTreeReadCallback = any(callback).(OnAfterReadInterface[Tree])
	
	case *View:
		stage.OnAfterViewReadCallback = any(callback).(OnAfterReadInterface[View])
	
	case *Xlsx:
		stage.OnAfterXlsxReadCallback = any(callback).(OnAfterReadInterface[Xlsx])
	
	}
}
