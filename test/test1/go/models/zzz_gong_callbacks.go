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
func (astruct *Astruct) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAstructCreateCallback != nil {
		stage.OnAfterAstructCreateCallback.OnAfterCreate(stage, astruct)
	}
}

func (astruct *Astruct) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAstructUpdateCallback != nil {
		var frontAstruct *Astruct
		if front != nil {
			frontAstruct, _ = front.(*Astruct)
		}
		stage.OnAfterAstructUpdateCallback.OnAfterUpdate(stage, astruct, frontAstruct)
	}
}

func (astruct *Astruct) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAstructDeleteCallback != nil {
		var frontAstruct *Astruct
		if front != nil {
			frontAstruct, _ = front.(*Astruct)
		}
		stage.OnAfterAstructDeleteCallback.OnAfterDelete(stage, astruct, frontAstruct)
	}
}

func (astructbstruct2use *AstructBstruct2Use) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAstructBstruct2UseCreateCallback != nil {
		stage.OnAfterAstructBstruct2UseCreateCallback.OnAfterCreate(stage, astructbstruct2use)
	}
}

func (astructbstruct2use *AstructBstruct2Use) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAstructBstruct2UseUpdateCallback != nil {
		var frontAstructBstruct2Use *AstructBstruct2Use
		if front != nil {
			frontAstructBstruct2Use, _ = front.(*AstructBstruct2Use)
		}
		stage.OnAfterAstructBstruct2UseUpdateCallback.OnAfterUpdate(stage, astructbstruct2use, frontAstructBstruct2Use)
	}
}

func (astructbstruct2use *AstructBstruct2Use) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAstructBstruct2UseDeleteCallback != nil {
		var frontAstructBstruct2Use *AstructBstruct2Use
		if front != nil {
			frontAstructBstruct2Use, _ = front.(*AstructBstruct2Use)
		}
		stage.OnAfterAstructBstruct2UseDeleteCallback.OnAfterDelete(stage, astructbstruct2use, frontAstructBstruct2Use)
	}
}

func (astructbstructuse *AstructBstructUse) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAstructBstructUseCreateCallback != nil {
		stage.OnAfterAstructBstructUseCreateCallback.OnAfterCreate(stage, astructbstructuse)
	}
}

func (astructbstructuse *AstructBstructUse) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAstructBstructUseUpdateCallback != nil {
		var frontAstructBstructUse *AstructBstructUse
		if front != nil {
			frontAstructBstructUse, _ = front.(*AstructBstructUse)
		}
		stage.OnAfterAstructBstructUseUpdateCallback.OnAfterUpdate(stage, astructbstructuse, frontAstructBstructUse)
	}
}

func (astructbstructuse *AstructBstructUse) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAstructBstructUseDeleteCallback != nil {
		var frontAstructBstructUse *AstructBstructUse
		if front != nil {
			frontAstructBstructUse, _ = front.(*AstructBstructUse)
		}
		stage.OnAfterAstructBstructUseDeleteCallback.OnAfterDelete(stage, astructbstructuse, frontAstructBstructUse)
	}
}

func (bstruct *Bstruct) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBstructCreateCallback != nil {
		stage.OnAfterBstructCreateCallback.OnAfterCreate(stage, bstruct)
	}
}

func (bstruct *Bstruct) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBstructUpdateCallback != nil {
		var frontBstruct *Bstruct
		if front != nil {
			frontBstruct, _ = front.(*Bstruct)
		}
		stage.OnAfterBstructUpdateCallback.OnAfterUpdate(stage, bstruct, frontBstruct)
	}
}

func (bstruct *Bstruct) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBstructDeleteCallback != nil {
		var frontBstruct *Bstruct
		if front != nil {
			frontBstruct, _ = front.(*Bstruct)
		}
		stage.OnAfterBstructDeleteCallback.OnAfterDelete(stage, bstruct, frontBstruct)
	}
}

func (dstruct *Dstruct) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDstructCreateCallback != nil {
		stage.OnAfterDstructCreateCallback.OnAfterCreate(stage, dstruct)
	}
}

func (dstruct *Dstruct) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDstructUpdateCallback != nil {
		var frontDstruct *Dstruct
		if front != nil {
			frontDstruct, _ = front.(*Dstruct)
		}
		stage.OnAfterDstructUpdateCallback.OnAfterUpdate(stage, dstruct, frontDstruct)
	}
}

func (dstruct *Dstruct) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDstructDeleteCallback != nil {
		var frontDstruct *Dstruct
		if front != nil {
			frontDstruct, _ = front.(*Dstruct)
		}
		stage.OnAfterDstructDeleteCallback.OnAfterDelete(stage, dstruct, frontDstruct)
	}
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterF0123456789012345678901234567890CreateCallback != nil {
		stage.OnAfterF0123456789012345678901234567890CreateCallback.OnAfterCreate(stage, f0123456789012345678901234567890)
	}
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterF0123456789012345678901234567890UpdateCallback != nil {
		var frontF0123456789012345678901234567890 *F0123456789012345678901234567890
		if front != nil {
			frontF0123456789012345678901234567890, _ = front.(*F0123456789012345678901234567890)
		}
		stage.OnAfterF0123456789012345678901234567890UpdateCallback.OnAfterUpdate(stage, f0123456789012345678901234567890, frontF0123456789012345678901234567890)
	}
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterF0123456789012345678901234567890DeleteCallback != nil {
		var frontF0123456789012345678901234567890 *F0123456789012345678901234567890
		if front != nil {
			frontF0123456789012345678901234567890, _ = front.(*F0123456789012345678901234567890)
		}
		stage.OnAfterF0123456789012345678901234567890DeleteCallback.OnAfterDelete(stage, f0123456789012345678901234567890, frontF0123456789012345678901234567890)
	}
}

func (gstruct *Gstruct) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGstructCreateCallback != nil {
		stage.OnAfterGstructCreateCallback.OnAfterCreate(stage, gstruct)
	}
}

func (gstruct *Gstruct) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGstructUpdateCallback != nil {
		var frontGstruct *Gstruct
		if front != nil {
			frontGstruct, _ = front.(*Gstruct)
		}
		stage.OnAfterGstructUpdateCallback.OnAfterUpdate(stage, gstruct, frontGstruct)
	}
}

func (gstruct *Gstruct) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGstructDeleteCallback != nil {
		var frontGstruct *Gstruct
		if front != nil {
			frontGstruct, _ = front.(*Gstruct)
		}
		stage.OnAfterGstructDeleteCallback.OnAfterDelete(stage, gstruct, frontGstruct)
	}
}

