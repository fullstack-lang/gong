// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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
	case *Threejs:
		if stage.OnAfterThreejsCreateCallback != nil {
			stage.OnAfterThreejsCreateCallback.OnAfterCreate(stage, target)
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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AsSplit:
		newTarget := any(new).(*AsSplit)
		if stage.OnAfterAsSplitUpdateCallback != nil {
			stage.OnAfterAsSplitUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AsSplitArea:
		newTarget := any(new).(*AsSplitArea)
		if stage.OnAfterAsSplitAreaUpdateCallback != nil {
			stage.OnAfterAsSplitAreaUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Button:
		newTarget := any(new).(*Button)
		if stage.OnAfterButtonUpdateCallback != nil {
			stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Cursor:
		newTarget := any(new).(*Cursor)
		if stage.OnAfterCursorUpdateCallback != nil {
			stage.OnAfterCursorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FavIcon:
		newTarget := any(new).(*FavIcon)
		if stage.OnAfterFavIconUpdateCallback != nil {
			stage.OnAfterFavIconUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Form:
		newTarget := any(new).(*Form)
		if stage.OnAfterFormUpdateCallback != nil {
			stage.OnAfterFormUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Load:
		newTarget := any(new).(*Load)
		if stage.OnAfterLoadUpdateCallback != nil {
			stage.OnAfterLoadUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LogoOnTheLeft:
		newTarget := any(new).(*LogoOnTheLeft)
		if stage.OnAfterLogoOnTheLeftUpdateCallback != nil {
			stage.OnAfterLogoOnTheLeftUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LogoOnTheRight:
		newTarget := any(new).(*LogoOnTheRight)
		if stage.OnAfterLogoOnTheRightUpdateCallback != nil {
			stage.OnAfterLogoOnTheRightUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Markdown:
		newTarget := any(new).(*Markdown)
		if stage.OnAfterMarkdownUpdateCallback != nil {
			stage.OnAfterMarkdownUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Slider:
		newTarget := any(new).(*Slider)
		if stage.OnAfterSliderUpdateCallback != nil {
			stage.OnAfterSliderUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Split:
		newTarget := any(new).(*Split)
		if stage.OnAfterSplitUpdateCallback != nil {
			stage.OnAfterSplitUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Svg:
		newTarget := any(new).(*Svg)
		if stage.OnAfterSvgUpdateCallback != nil {
			stage.OnAfterSvgUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Table:
		newTarget := any(new).(*Table)
		if stage.OnAfterTableUpdateCallback != nil {
			stage.OnAfterTableUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Threejs:
		newTarget := any(new).(*Threejs)
		if stage.OnAfterThreejsUpdateCallback != nil {
			stage.OnAfterThreejsUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Title:
		newTarget := any(new).(*Title)
		if stage.OnAfterTitleUpdateCallback != nil {
			stage.OnAfterTitleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tone:
		newTarget := any(new).(*Tone)
		if stage.OnAfterToneUpdateCallback != nil {
			stage.OnAfterToneUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tree:
		newTarget := any(new).(*Tree)
		if stage.OnAfterTreeUpdateCallback != nil {
			stage.OnAfterTreeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *View:
		newTarget := any(new).(*View)
		if stage.OnAfterViewUpdateCallback != nil {
			stage.OnAfterViewUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xlsx:
		newTarget := any(new).(*Xlsx)
		if stage.OnAfterXlsxUpdateCallback != nil {
			stage.OnAfterXlsxUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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
	case *Threejs:
		if stage.OnAfterThreejsDeleteCallback != nil {
			staged := any(staged).(*Threejs)
			stage.OnAfterThreejsDeleteCallback.OnAfterDelete(stage, staged, front)
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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
