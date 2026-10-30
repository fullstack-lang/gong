// generated code - do not edit
package aixmfeatures

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
func (a_extension *A_extension) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extensionCreateCallback != nil {
		stage.OnAfterA_extensionCreateCallback.OnAfterCreate(stage, a_extension)
	}
}

func (a_extension *A_extension) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extensionUpdateCallback != nil {
		var frontA_extension *A_extension
		if front != nil {
			frontA_extension, _ = front.(*A_extension)
		}
		stage.OnAfterA_extensionUpdateCallback.OnAfterUpdate(stage, a_extension, frontA_extension)
	}
}

func (a_extension *A_extension) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extensionDeleteCallback != nil {
		var frontA_extension *A_extension
		if front != nil {
			frontA_extension, _ = front.(*A_extension)
		}
		stage.OnAfterA_extensionDeleteCallback.OnAfterDelete(stage, a_extension, frontA_extension)
	}
}

func (a_extension_1 *A_extension_1) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_1CreateCallback != nil {
		stage.OnAfterA_extension_1CreateCallback.OnAfterCreate(stage, a_extension_1)
	}
}

func (a_extension_1 *A_extension_1) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_1UpdateCallback != nil {
		var frontA_extension_1 *A_extension_1
		if front != nil {
			frontA_extension_1, _ = front.(*A_extension_1)
		}
		stage.OnAfterA_extension_1UpdateCallback.OnAfterUpdate(stage, a_extension_1, frontA_extension_1)
	}
}

func (a_extension_1 *A_extension_1) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_1DeleteCallback != nil {
		var frontA_extension_1 *A_extension_1
		if front != nil {
			frontA_extension_1, _ = front.(*A_extension_1)
		}
		stage.OnAfterA_extension_1DeleteCallback.OnAfterDelete(stage, a_extension_1, frontA_extension_1)
	}
}

func (a_extension_10 *A_extension_10) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_10CreateCallback != nil {
		stage.OnAfterA_extension_10CreateCallback.OnAfterCreate(stage, a_extension_10)
	}
}

func (a_extension_10 *A_extension_10) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_10UpdateCallback != nil {
		var frontA_extension_10 *A_extension_10
		if front != nil {
			frontA_extension_10, _ = front.(*A_extension_10)
		}
		stage.OnAfterA_extension_10UpdateCallback.OnAfterUpdate(stage, a_extension_10, frontA_extension_10)
	}
}

func (a_extension_10 *A_extension_10) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_10DeleteCallback != nil {
		var frontA_extension_10 *A_extension_10
		if front != nil {
			frontA_extension_10, _ = front.(*A_extension_10)
		}
		stage.OnAfterA_extension_10DeleteCallback.OnAfterDelete(stage, a_extension_10, frontA_extension_10)
	}
}

func (a_extension_11 *A_extension_11) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_11CreateCallback != nil {
		stage.OnAfterA_extension_11CreateCallback.OnAfterCreate(stage, a_extension_11)
	}
}

func (a_extension_11 *A_extension_11) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_11UpdateCallback != nil {
		var frontA_extension_11 *A_extension_11
		if front != nil {
			frontA_extension_11, _ = front.(*A_extension_11)
		}
		stage.OnAfterA_extension_11UpdateCallback.OnAfterUpdate(stage, a_extension_11, frontA_extension_11)
	}
}

func (a_extension_11 *A_extension_11) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_11DeleteCallback != nil {
		var frontA_extension_11 *A_extension_11
		if front != nil {
			frontA_extension_11, _ = front.(*A_extension_11)
		}
		stage.OnAfterA_extension_11DeleteCallback.OnAfterDelete(stage, a_extension_11, frontA_extension_11)
	}
}

func (a_extension_12 *A_extension_12) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_12CreateCallback != nil {
		stage.OnAfterA_extension_12CreateCallback.OnAfterCreate(stage, a_extension_12)
	}
}

func (a_extension_12 *A_extension_12) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_12UpdateCallback != nil {
		var frontA_extension_12 *A_extension_12
		if front != nil {
			frontA_extension_12, _ = front.(*A_extension_12)
		}
		stage.OnAfterA_extension_12UpdateCallback.OnAfterUpdate(stage, a_extension_12, frontA_extension_12)
	}
}

func (a_extension_12 *A_extension_12) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_12DeleteCallback != nil {
		var frontA_extension_12 *A_extension_12
		if front != nil {
			frontA_extension_12, _ = front.(*A_extension_12)
		}
		stage.OnAfterA_extension_12DeleteCallback.OnAfterDelete(stage, a_extension_12, frontA_extension_12)
	}
}

func (a_extension_13 *A_extension_13) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_13CreateCallback != nil {
		stage.OnAfterA_extension_13CreateCallback.OnAfterCreate(stage, a_extension_13)
	}
}

func (a_extension_13 *A_extension_13) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_13UpdateCallback != nil {
		var frontA_extension_13 *A_extension_13
		if front != nil {
			frontA_extension_13, _ = front.(*A_extension_13)
		}
		stage.OnAfterA_extension_13UpdateCallback.OnAfterUpdate(stage, a_extension_13, frontA_extension_13)
	}
}

func (a_extension_13 *A_extension_13) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_13DeleteCallback != nil {
		var frontA_extension_13 *A_extension_13
		if front != nil {
			frontA_extension_13, _ = front.(*A_extension_13)
		}
		stage.OnAfterA_extension_13DeleteCallback.OnAfterDelete(stage, a_extension_13, frontA_extension_13)
	}
}

func (a_extension_14 *A_extension_14) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_14CreateCallback != nil {
		stage.OnAfterA_extension_14CreateCallback.OnAfterCreate(stage, a_extension_14)
	}
}

func (a_extension_14 *A_extension_14) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_14UpdateCallback != nil {
		var frontA_extension_14 *A_extension_14
		if front != nil {
			frontA_extension_14, _ = front.(*A_extension_14)
		}
		stage.OnAfterA_extension_14UpdateCallback.OnAfterUpdate(stage, a_extension_14, frontA_extension_14)
	}
}

func (a_extension_14 *A_extension_14) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_14DeleteCallback != nil {
		var frontA_extension_14 *A_extension_14
		if front != nil {
			frontA_extension_14, _ = front.(*A_extension_14)
		}
		stage.OnAfterA_extension_14DeleteCallback.OnAfterDelete(stage, a_extension_14, frontA_extension_14)
	}
}

func (a_extension_15 *A_extension_15) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_15CreateCallback != nil {
		stage.OnAfterA_extension_15CreateCallback.OnAfterCreate(stage, a_extension_15)
	}
}

func (a_extension_15 *A_extension_15) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_15UpdateCallback != nil {
		var frontA_extension_15 *A_extension_15
		if front != nil {
			frontA_extension_15, _ = front.(*A_extension_15)
		}
		stage.OnAfterA_extension_15UpdateCallback.OnAfterUpdate(stage, a_extension_15, frontA_extension_15)
	}
}

func (a_extension_15 *A_extension_15) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_15DeleteCallback != nil {
		var frontA_extension_15 *A_extension_15
		if front != nil {
			frontA_extension_15, _ = front.(*A_extension_15)
		}
		stage.OnAfterA_extension_15DeleteCallback.OnAfterDelete(stage, a_extension_15, frontA_extension_15)
	}
}

func (a_extension_16 *A_extension_16) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_16CreateCallback != nil {
		stage.OnAfterA_extension_16CreateCallback.OnAfterCreate(stage, a_extension_16)
	}
}

func (a_extension_16 *A_extension_16) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_16UpdateCallback != nil {
		var frontA_extension_16 *A_extension_16
		if front != nil {
			frontA_extension_16, _ = front.(*A_extension_16)
		}
		stage.OnAfterA_extension_16UpdateCallback.OnAfterUpdate(stage, a_extension_16, frontA_extension_16)
	}
}

func (a_extension_16 *A_extension_16) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_16DeleteCallback != nil {
		var frontA_extension_16 *A_extension_16
		if front != nil {
			frontA_extension_16, _ = front.(*A_extension_16)
		}
		stage.OnAfterA_extension_16DeleteCallback.OnAfterDelete(stage, a_extension_16, frontA_extension_16)
	}
}

func (a_extension_17 *A_extension_17) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_17CreateCallback != nil {
		stage.OnAfterA_extension_17CreateCallback.OnAfterCreate(stage, a_extension_17)
	}
}

func (a_extension_17 *A_extension_17) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_17UpdateCallback != nil {
		var frontA_extension_17 *A_extension_17
		if front != nil {
			frontA_extension_17, _ = front.(*A_extension_17)
		}
		stage.OnAfterA_extension_17UpdateCallback.OnAfterUpdate(stage, a_extension_17, frontA_extension_17)
	}
}

func (a_extension_17 *A_extension_17) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_17DeleteCallback != nil {
		var frontA_extension_17 *A_extension_17
		if front != nil {
			frontA_extension_17, _ = front.(*A_extension_17)
		}
		stage.OnAfterA_extension_17DeleteCallback.OnAfterDelete(stage, a_extension_17, frontA_extension_17)
	}
}

func (a_extension_18 *A_extension_18) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_18CreateCallback != nil {
		stage.OnAfterA_extension_18CreateCallback.OnAfterCreate(stage, a_extension_18)
	}
}

func (a_extension_18 *A_extension_18) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_18UpdateCallback != nil {
		var frontA_extension_18 *A_extension_18
		if front != nil {
			frontA_extension_18, _ = front.(*A_extension_18)
		}
		stage.OnAfterA_extension_18UpdateCallback.OnAfterUpdate(stage, a_extension_18, frontA_extension_18)
	}
}

func (a_extension_18 *A_extension_18) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_18DeleteCallback != nil {
		var frontA_extension_18 *A_extension_18
		if front != nil {
			frontA_extension_18, _ = front.(*A_extension_18)
		}
		stage.OnAfterA_extension_18DeleteCallback.OnAfterDelete(stage, a_extension_18, frontA_extension_18)
	}
}

func (a_extension_19 *A_extension_19) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_19CreateCallback != nil {
		stage.OnAfterA_extension_19CreateCallback.OnAfterCreate(stage, a_extension_19)
	}
}

func (a_extension_19 *A_extension_19) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_19UpdateCallback != nil {
		var frontA_extension_19 *A_extension_19
		if front != nil {
			frontA_extension_19, _ = front.(*A_extension_19)
		}
		stage.OnAfterA_extension_19UpdateCallback.OnAfterUpdate(stage, a_extension_19, frontA_extension_19)
	}
}

func (a_extension_19 *A_extension_19) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_19DeleteCallback != nil {
		var frontA_extension_19 *A_extension_19
		if front != nil {
			frontA_extension_19, _ = front.(*A_extension_19)
		}
		stage.OnAfterA_extension_19DeleteCallback.OnAfterDelete(stage, a_extension_19, frontA_extension_19)
	}
}

func (a_extension_2 *A_extension_2) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_2CreateCallback != nil {
		stage.OnAfterA_extension_2CreateCallback.OnAfterCreate(stage, a_extension_2)
	}
}

func (a_extension_2 *A_extension_2) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_2UpdateCallback != nil {
		var frontA_extension_2 *A_extension_2
		if front != nil {
			frontA_extension_2, _ = front.(*A_extension_2)
		}
		stage.OnAfterA_extension_2UpdateCallback.OnAfterUpdate(stage, a_extension_2, frontA_extension_2)
	}
}

func (a_extension_2 *A_extension_2) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_2DeleteCallback != nil {
		var frontA_extension_2 *A_extension_2
		if front != nil {
			frontA_extension_2, _ = front.(*A_extension_2)
		}
		stage.OnAfterA_extension_2DeleteCallback.OnAfterDelete(stage, a_extension_2, frontA_extension_2)
	}
}

func (a_extension_20 *A_extension_20) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_20CreateCallback != nil {
		stage.OnAfterA_extension_20CreateCallback.OnAfterCreate(stage, a_extension_20)
	}
}

func (a_extension_20 *A_extension_20) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_20UpdateCallback != nil {
		var frontA_extension_20 *A_extension_20
		if front != nil {
			frontA_extension_20, _ = front.(*A_extension_20)
		}
		stage.OnAfterA_extension_20UpdateCallback.OnAfterUpdate(stage, a_extension_20, frontA_extension_20)
	}
}

func (a_extension_20 *A_extension_20) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_20DeleteCallback != nil {
		var frontA_extension_20 *A_extension_20
		if front != nil {
			frontA_extension_20, _ = front.(*A_extension_20)
		}
		stage.OnAfterA_extension_20DeleteCallback.OnAfterDelete(stage, a_extension_20, frontA_extension_20)
	}
}

func (a_extension_21 *A_extension_21) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_21CreateCallback != nil {
		stage.OnAfterA_extension_21CreateCallback.OnAfterCreate(stage, a_extension_21)
	}
}

func (a_extension_21 *A_extension_21) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_21UpdateCallback != nil {
		var frontA_extension_21 *A_extension_21
		if front != nil {
			frontA_extension_21, _ = front.(*A_extension_21)
		}
		stage.OnAfterA_extension_21UpdateCallback.OnAfterUpdate(stage, a_extension_21, frontA_extension_21)
	}
}

func (a_extension_21 *A_extension_21) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_21DeleteCallback != nil {
		var frontA_extension_21 *A_extension_21
		if front != nil {
			frontA_extension_21, _ = front.(*A_extension_21)
		}
		stage.OnAfterA_extension_21DeleteCallback.OnAfterDelete(stage, a_extension_21, frontA_extension_21)
	}
}

func (a_extension_22 *A_extension_22) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_22CreateCallback != nil {
		stage.OnAfterA_extension_22CreateCallback.OnAfterCreate(stage, a_extension_22)
	}
}

func (a_extension_22 *A_extension_22) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_22UpdateCallback != nil {
		var frontA_extension_22 *A_extension_22
		if front != nil {
			frontA_extension_22, _ = front.(*A_extension_22)
		}
		stage.OnAfterA_extension_22UpdateCallback.OnAfterUpdate(stage, a_extension_22, frontA_extension_22)
	}
}

func (a_extension_22 *A_extension_22) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_22DeleteCallback != nil {
		var frontA_extension_22 *A_extension_22
		if front != nil {
			frontA_extension_22, _ = front.(*A_extension_22)
		}
		stage.OnAfterA_extension_22DeleteCallback.OnAfterDelete(stage, a_extension_22, frontA_extension_22)
	}
}

func (a_extension_23 *A_extension_23) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_23CreateCallback != nil {
		stage.OnAfterA_extension_23CreateCallback.OnAfterCreate(stage, a_extension_23)
	}
}

func (a_extension_23 *A_extension_23) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_23UpdateCallback != nil {
		var frontA_extension_23 *A_extension_23
		if front != nil {
			frontA_extension_23, _ = front.(*A_extension_23)
		}
		stage.OnAfterA_extension_23UpdateCallback.OnAfterUpdate(stage, a_extension_23, frontA_extension_23)
	}
}

func (a_extension_23 *A_extension_23) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_23DeleteCallback != nil {
		var frontA_extension_23 *A_extension_23
		if front != nil {
			frontA_extension_23, _ = front.(*A_extension_23)
		}
		stage.OnAfterA_extension_23DeleteCallback.OnAfterDelete(stage, a_extension_23, frontA_extension_23)
	}
}

func (a_extension_24 *A_extension_24) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_24CreateCallback != nil {
		stage.OnAfterA_extension_24CreateCallback.OnAfterCreate(stage, a_extension_24)
	}
}

func (a_extension_24 *A_extension_24) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_24UpdateCallback != nil {
		var frontA_extension_24 *A_extension_24
		if front != nil {
			frontA_extension_24, _ = front.(*A_extension_24)
		}
		stage.OnAfterA_extension_24UpdateCallback.OnAfterUpdate(stage, a_extension_24, frontA_extension_24)
	}
}

func (a_extension_24 *A_extension_24) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_24DeleteCallback != nil {
		var frontA_extension_24 *A_extension_24
		if front != nil {
			frontA_extension_24, _ = front.(*A_extension_24)
		}
		stage.OnAfterA_extension_24DeleteCallback.OnAfterDelete(stage, a_extension_24, frontA_extension_24)
	}
}

func (a_extension_25 *A_extension_25) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_25CreateCallback != nil {
		stage.OnAfterA_extension_25CreateCallback.OnAfterCreate(stage, a_extension_25)
	}
}

func (a_extension_25 *A_extension_25) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_25UpdateCallback != nil {
		var frontA_extension_25 *A_extension_25
		if front != nil {
			frontA_extension_25, _ = front.(*A_extension_25)
		}
		stage.OnAfterA_extension_25UpdateCallback.OnAfterUpdate(stage, a_extension_25, frontA_extension_25)
	}
}

func (a_extension_25 *A_extension_25) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_25DeleteCallback != nil {
		var frontA_extension_25 *A_extension_25
		if front != nil {
			frontA_extension_25, _ = front.(*A_extension_25)
		}
		stage.OnAfterA_extension_25DeleteCallback.OnAfterDelete(stage, a_extension_25, frontA_extension_25)
	}
}

func (a_extension_26 *A_extension_26) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_26CreateCallback != nil {
		stage.OnAfterA_extension_26CreateCallback.OnAfterCreate(stage, a_extension_26)
	}
}

func (a_extension_26 *A_extension_26) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_26UpdateCallback != nil {
		var frontA_extension_26 *A_extension_26
		if front != nil {
			frontA_extension_26, _ = front.(*A_extension_26)
		}
		stage.OnAfterA_extension_26UpdateCallback.OnAfterUpdate(stage, a_extension_26, frontA_extension_26)
	}
}

func (a_extension_26 *A_extension_26) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_26DeleteCallback != nil {
		var frontA_extension_26 *A_extension_26
		if front != nil {
			frontA_extension_26, _ = front.(*A_extension_26)
		}
		stage.OnAfterA_extension_26DeleteCallback.OnAfterDelete(stage, a_extension_26, frontA_extension_26)
	}
}

func (a_extension_27 *A_extension_27) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_27CreateCallback != nil {
		stage.OnAfterA_extension_27CreateCallback.OnAfterCreate(stage, a_extension_27)
	}
}

func (a_extension_27 *A_extension_27) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_27UpdateCallback != nil {
		var frontA_extension_27 *A_extension_27
		if front != nil {
			frontA_extension_27, _ = front.(*A_extension_27)
		}
		stage.OnAfterA_extension_27UpdateCallback.OnAfterUpdate(stage, a_extension_27, frontA_extension_27)
	}
}

func (a_extension_27 *A_extension_27) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_27DeleteCallback != nil {
		var frontA_extension_27 *A_extension_27
		if front != nil {
			frontA_extension_27, _ = front.(*A_extension_27)
		}
		stage.OnAfterA_extension_27DeleteCallback.OnAfterDelete(stage, a_extension_27, frontA_extension_27)
	}
}

func (a_extension_28 *A_extension_28) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_28CreateCallback != nil {
		stage.OnAfterA_extension_28CreateCallback.OnAfterCreate(stage, a_extension_28)
	}
}

func (a_extension_28 *A_extension_28) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_28UpdateCallback != nil {
		var frontA_extension_28 *A_extension_28
		if front != nil {
			frontA_extension_28, _ = front.(*A_extension_28)
		}
		stage.OnAfterA_extension_28UpdateCallback.OnAfterUpdate(stage, a_extension_28, frontA_extension_28)
	}
}

func (a_extension_28 *A_extension_28) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_28DeleteCallback != nil {
		var frontA_extension_28 *A_extension_28
		if front != nil {
			frontA_extension_28, _ = front.(*A_extension_28)
		}
		stage.OnAfterA_extension_28DeleteCallback.OnAfterDelete(stage, a_extension_28, frontA_extension_28)
	}
}

func (a_extension_29 *A_extension_29) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_29CreateCallback != nil {
		stage.OnAfterA_extension_29CreateCallback.OnAfterCreate(stage, a_extension_29)
	}
}

func (a_extension_29 *A_extension_29) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_29UpdateCallback != nil {
		var frontA_extension_29 *A_extension_29
		if front != nil {
			frontA_extension_29, _ = front.(*A_extension_29)
		}
		stage.OnAfterA_extension_29UpdateCallback.OnAfterUpdate(stage, a_extension_29, frontA_extension_29)
	}
}

func (a_extension_29 *A_extension_29) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_29DeleteCallback != nil {
		var frontA_extension_29 *A_extension_29
		if front != nil {
			frontA_extension_29, _ = front.(*A_extension_29)
		}
		stage.OnAfterA_extension_29DeleteCallback.OnAfterDelete(stage, a_extension_29, frontA_extension_29)
	}
}

func (a_extension_3 *A_extension_3) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_3CreateCallback != nil {
		stage.OnAfterA_extension_3CreateCallback.OnAfterCreate(stage, a_extension_3)
	}
}

func (a_extension_3 *A_extension_3) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_3UpdateCallback != nil {
		var frontA_extension_3 *A_extension_3
		if front != nil {
			frontA_extension_3, _ = front.(*A_extension_3)
		}
		stage.OnAfterA_extension_3UpdateCallback.OnAfterUpdate(stage, a_extension_3, frontA_extension_3)
	}
}

func (a_extension_3 *A_extension_3) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_3DeleteCallback != nil {
		var frontA_extension_3 *A_extension_3
		if front != nil {
			frontA_extension_3, _ = front.(*A_extension_3)
		}
		stage.OnAfterA_extension_3DeleteCallback.OnAfterDelete(stage, a_extension_3, frontA_extension_3)
	}
}

func (a_extension_30 *A_extension_30) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_30CreateCallback != nil {
		stage.OnAfterA_extension_30CreateCallback.OnAfterCreate(stage, a_extension_30)
	}
}

func (a_extension_30 *A_extension_30) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_30UpdateCallback != nil {
		var frontA_extension_30 *A_extension_30
		if front != nil {
			frontA_extension_30, _ = front.(*A_extension_30)
		}
		stage.OnAfterA_extension_30UpdateCallback.OnAfterUpdate(stage, a_extension_30, frontA_extension_30)
	}
}

func (a_extension_30 *A_extension_30) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_30DeleteCallback != nil {
		var frontA_extension_30 *A_extension_30
		if front != nil {
			frontA_extension_30, _ = front.(*A_extension_30)
		}
		stage.OnAfterA_extension_30DeleteCallback.OnAfterDelete(stage, a_extension_30, frontA_extension_30)
	}
}

func (a_extension_31 *A_extension_31) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_31CreateCallback != nil {
		stage.OnAfterA_extension_31CreateCallback.OnAfterCreate(stage, a_extension_31)
	}
}

func (a_extension_31 *A_extension_31) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_31UpdateCallback != nil {
		var frontA_extension_31 *A_extension_31
		if front != nil {
			frontA_extension_31, _ = front.(*A_extension_31)
		}
		stage.OnAfterA_extension_31UpdateCallback.OnAfterUpdate(stage, a_extension_31, frontA_extension_31)
	}
}

func (a_extension_31 *A_extension_31) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_31DeleteCallback != nil {
		var frontA_extension_31 *A_extension_31
		if front != nil {
			frontA_extension_31, _ = front.(*A_extension_31)
		}
		stage.OnAfterA_extension_31DeleteCallback.OnAfterDelete(stage, a_extension_31, frontA_extension_31)
	}
}

func (a_extension_32 *A_extension_32) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_32CreateCallback != nil {
		stage.OnAfterA_extension_32CreateCallback.OnAfterCreate(stage, a_extension_32)
	}
}

func (a_extension_32 *A_extension_32) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_32UpdateCallback != nil {
		var frontA_extension_32 *A_extension_32
		if front != nil {
			frontA_extension_32, _ = front.(*A_extension_32)
		}
		stage.OnAfterA_extension_32UpdateCallback.OnAfterUpdate(stage, a_extension_32, frontA_extension_32)
	}
}

func (a_extension_32 *A_extension_32) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_32DeleteCallback != nil {
		var frontA_extension_32 *A_extension_32
		if front != nil {
			frontA_extension_32, _ = front.(*A_extension_32)
		}
		stage.OnAfterA_extension_32DeleteCallback.OnAfterDelete(stage, a_extension_32, frontA_extension_32)
	}
}

func (a_extension_33 *A_extension_33) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_33CreateCallback != nil {
		stage.OnAfterA_extension_33CreateCallback.OnAfterCreate(stage, a_extension_33)
	}
}

func (a_extension_33 *A_extension_33) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_33UpdateCallback != nil {
		var frontA_extension_33 *A_extension_33
		if front != nil {
			frontA_extension_33, _ = front.(*A_extension_33)
		}
		stage.OnAfterA_extension_33UpdateCallback.OnAfterUpdate(stage, a_extension_33, frontA_extension_33)
	}
}

func (a_extension_33 *A_extension_33) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_33DeleteCallback != nil {
		var frontA_extension_33 *A_extension_33
		if front != nil {
			frontA_extension_33, _ = front.(*A_extension_33)
		}
		stage.OnAfterA_extension_33DeleteCallback.OnAfterDelete(stage, a_extension_33, frontA_extension_33)
	}
}

func (a_extension_34 *A_extension_34) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_34CreateCallback != nil {
		stage.OnAfterA_extension_34CreateCallback.OnAfterCreate(stage, a_extension_34)
	}
}

func (a_extension_34 *A_extension_34) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_34UpdateCallback != nil {
		var frontA_extension_34 *A_extension_34
		if front != nil {
			frontA_extension_34, _ = front.(*A_extension_34)
		}
		stage.OnAfterA_extension_34UpdateCallback.OnAfterUpdate(stage, a_extension_34, frontA_extension_34)
	}
}

func (a_extension_34 *A_extension_34) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_34DeleteCallback != nil {
		var frontA_extension_34 *A_extension_34
		if front != nil {
			frontA_extension_34, _ = front.(*A_extension_34)
		}
		stage.OnAfterA_extension_34DeleteCallback.OnAfterDelete(stage, a_extension_34, frontA_extension_34)
	}
}

func (a_extension_35 *A_extension_35) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_35CreateCallback != nil {
		stage.OnAfterA_extension_35CreateCallback.OnAfterCreate(stage, a_extension_35)
	}
}

func (a_extension_35 *A_extension_35) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_35UpdateCallback != nil {
		var frontA_extension_35 *A_extension_35
		if front != nil {
			frontA_extension_35, _ = front.(*A_extension_35)
		}
		stage.OnAfterA_extension_35UpdateCallback.OnAfterUpdate(stage, a_extension_35, frontA_extension_35)
	}
}

func (a_extension_35 *A_extension_35) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_35DeleteCallback != nil {
		var frontA_extension_35 *A_extension_35
		if front != nil {
			frontA_extension_35, _ = front.(*A_extension_35)
		}
		stage.OnAfterA_extension_35DeleteCallback.OnAfterDelete(stage, a_extension_35, frontA_extension_35)
	}
}

func (a_extension_36 *A_extension_36) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_36CreateCallback != nil {
		stage.OnAfterA_extension_36CreateCallback.OnAfterCreate(stage, a_extension_36)
	}
}

func (a_extension_36 *A_extension_36) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_36UpdateCallback != nil {
		var frontA_extension_36 *A_extension_36
		if front != nil {
			frontA_extension_36, _ = front.(*A_extension_36)
		}
		stage.OnAfterA_extension_36UpdateCallback.OnAfterUpdate(stage, a_extension_36, frontA_extension_36)
	}
}

func (a_extension_36 *A_extension_36) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_36DeleteCallback != nil {
		var frontA_extension_36 *A_extension_36
		if front != nil {
			frontA_extension_36, _ = front.(*A_extension_36)
		}
		stage.OnAfterA_extension_36DeleteCallback.OnAfterDelete(stage, a_extension_36, frontA_extension_36)
	}
}

func (a_extension_37 *A_extension_37) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_37CreateCallback != nil {
		stage.OnAfterA_extension_37CreateCallback.OnAfterCreate(stage, a_extension_37)
	}
}

func (a_extension_37 *A_extension_37) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_37UpdateCallback != nil {
		var frontA_extension_37 *A_extension_37
		if front != nil {
			frontA_extension_37, _ = front.(*A_extension_37)
		}
		stage.OnAfterA_extension_37UpdateCallback.OnAfterUpdate(stage, a_extension_37, frontA_extension_37)
	}
}

func (a_extension_37 *A_extension_37) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_37DeleteCallback != nil {
		var frontA_extension_37 *A_extension_37
		if front != nil {
			frontA_extension_37, _ = front.(*A_extension_37)
		}
		stage.OnAfterA_extension_37DeleteCallback.OnAfterDelete(stage, a_extension_37, frontA_extension_37)
	}
}

func (a_extension_38 *A_extension_38) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_38CreateCallback != nil {
		stage.OnAfterA_extension_38CreateCallback.OnAfterCreate(stage, a_extension_38)
	}
}

func (a_extension_38 *A_extension_38) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_38UpdateCallback != nil {
		var frontA_extension_38 *A_extension_38
		if front != nil {
			frontA_extension_38, _ = front.(*A_extension_38)
		}
		stage.OnAfterA_extension_38UpdateCallback.OnAfterUpdate(stage, a_extension_38, frontA_extension_38)
	}
}

func (a_extension_38 *A_extension_38) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_38DeleteCallback != nil {
		var frontA_extension_38 *A_extension_38
		if front != nil {
			frontA_extension_38, _ = front.(*A_extension_38)
		}
		stage.OnAfterA_extension_38DeleteCallback.OnAfterDelete(stage, a_extension_38, frontA_extension_38)
	}
}

func (a_extension_39 *A_extension_39) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_39CreateCallback != nil {
		stage.OnAfterA_extension_39CreateCallback.OnAfterCreate(stage, a_extension_39)
	}
}

func (a_extension_39 *A_extension_39) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_39UpdateCallback != nil {
		var frontA_extension_39 *A_extension_39
		if front != nil {
			frontA_extension_39, _ = front.(*A_extension_39)
		}
		stage.OnAfterA_extension_39UpdateCallback.OnAfterUpdate(stage, a_extension_39, frontA_extension_39)
	}
}

func (a_extension_39 *A_extension_39) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_39DeleteCallback != nil {
		var frontA_extension_39 *A_extension_39
		if front != nil {
			frontA_extension_39, _ = front.(*A_extension_39)
		}
		stage.OnAfterA_extension_39DeleteCallback.OnAfterDelete(stage, a_extension_39, frontA_extension_39)
	}
}

func (a_extension_4 *A_extension_4) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_4CreateCallback != nil {
		stage.OnAfterA_extension_4CreateCallback.OnAfterCreate(stage, a_extension_4)
	}
}

func (a_extension_4 *A_extension_4) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_4UpdateCallback != nil {
		var frontA_extension_4 *A_extension_4
		if front != nil {
			frontA_extension_4, _ = front.(*A_extension_4)
		}
		stage.OnAfterA_extension_4UpdateCallback.OnAfterUpdate(stage, a_extension_4, frontA_extension_4)
	}
}

func (a_extension_4 *A_extension_4) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_4DeleteCallback != nil {
		var frontA_extension_4 *A_extension_4
		if front != nil {
			frontA_extension_4, _ = front.(*A_extension_4)
		}
		stage.OnAfterA_extension_4DeleteCallback.OnAfterDelete(stage, a_extension_4, frontA_extension_4)
	}
}

func (a_extension_40 *A_extension_40) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_40CreateCallback != nil {
		stage.OnAfterA_extension_40CreateCallback.OnAfterCreate(stage, a_extension_40)
	}
}

func (a_extension_40 *A_extension_40) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_40UpdateCallback != nil {
		var frontA_extension_40 *A_extension_40
		if front != nil {
			frontA_extension_40, _ = front.(*A_extension_40)
		}
		stage.OnAfterA_extension_40UpdateCallback.OnAfterUpdate(stage, a_extension_40, frontA_extension_40)
	}
}

func (a_extension_40 *A_extension_40) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_40DeleteCallback != nil {
		var frontA_extension_40 *A_extension_40
		if front != nil {
			frontA_extension_40, _ = front.(*A_extension_40)
		}
		stage.OnAfterA_extension_40DeleteCallback.OnAfterDelete(stage, a_extension_40, frontA_extension_40)
	}
}

func (a_extension_41 *A_extension_41) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_41CreateCallback != nil {
		stage.OnAfterA_extension_41CreateCallback.OnAfterCreate(stage, a_extension_41)
	}
}

func (a_extension_41 *A_extension_41) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_41UpdateCallback != nil {
		var frontA_extension_41 *A_extension_41
		if front != nil {
			frontA_extension_41, _ = front.(*A_extension_41)
		}
		stage.OnAfterA_extension_41UpdateCallback.OnAfterUpdate(stage, a_extension_41, frontA_extension_41)
	}
}

func (a_extension_41 *A_extension_41) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_41DeleteCallback != nil {
		var frontA_extension_41 *A_extension_41
		if front != nil {
			frontA_extension_41, _ = front.(*A_extension_41)
		}
		stage.OnAfterA_extension_41DeleteCallback.OnAfterDelete(stage, a_extension_41, frontA_extension_41)
	}
}

func (a_extension_42 *A_extension_42) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_42CreateCallback != nil {
		stage.OnAfterA_extension_42CreateCallback.OnAfterCreate(stage, a_extension_42)
	}
}

func (a_extension_42 *A_extension_42) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_42UpdateCallback != nil {
		var frontA_extension_42 *A_extension_42
		if front != nil {
			frontA_extension_42, _ = front.(*A_extension_42)
		}
		stage.OnAfterA_extension_42UpdateCallback.OnAfterUpdate(stage, a_extension_42, frontA_extension_42)
	}
}

func (a_extension_42 *A_extension_42) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_42DeleteCallback != nil {
		var frontA_extension_42 *A_extension_42
		if front != nil {
			frontA_extension_42, _ = front.(*A_extension_42)
		}
		stage.OnAfterA_extension_42DeleteCallback.OnAfterDelete(stage, a_extension_42, frontA_extension_42)
	}
}

func (a_extension_43 *A_extension_43) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_43CreateCallback != nil {
		stage.OnAfterA_extension_43CreateCallback.OnAfterCreate(stage, a_extension_43)
	}
}

func (a_extension_43 *A_extension_43) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_43UpdateCallback != nil {
		var frontA_extension_43 *A_extension_43
		if front != nil {
			frontA_extension_43, _ = front.(*A_extension_43)
		}
		stage.OnAfterA_extension_43UpdateCallback.OnAfterUpdate(stage, a_extension_43, frontA_extension_43)
	}
}

func (a_extension_43 *A_extension_43) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_43DeleteCallback != nil {
		var frontA_extension_43 *A_extension_43
		if front != nil {
			frontA_extension_43, _ = front.(*A_extension_43)
		}
		stage.OnAfterA_extension_43DeleteCallback.OnAfterDelete(stage, a_extension_43, frontA_extension_43)
	}
}

func (a_extension_44 *A_extension_44) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_44CreateCallback != nil {
		stage.OnAfterA_extension_44CreateCallback.OnAfterCreate(stage, a_extension_44)
	}
}

func (a_extension_44 *A_extension_44) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_44UpdateCallback != nil {
		var frontA_extension_44 *A_extension_44
		if front != nil {
			frontA_extension_44, _ = front.(*A_extension_44)
		}
		stage.OnAfterA_extension_44UpdateCallback.OnAfterUpdate(stage, a_extension_44, frontA_extension_44)
	}
}

func (a_extension_44 *A_extension_44) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_44DeleteCallback != nil {
		var frontA_extension_44 *A_extension_44
		if front != nil {
			frontA_extension_44, _ = front.(*A_extension_44)
		}
		stage.OnAfterA_extension_44DeleteCallback.OnAfterDelete(stage, a_extension_44, frontA_extension_44)
	}
}

func (a_extension_45 *A_extension_45) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_45CreateCallback != nil {
		stage.OnAfterA_extension_45CreateCallback.OnAfterCreate(stage, a_extension_45)
	}
}

func (a_extension_45 *A_extension_45) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_45UpdateCallback != nil {
		var frontA_extension_45 *A_extension_45
		if front != nil {
			frontA_extension_45, _ = front.(*A_extension_45)
		}
		stage.OnAfterA_extension_45UpdateCallback.OnAfterUpdate(stage, a_extension_45, frontA_extension_45)
	}
}

func (a_extension_45 *A_extension_45) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_45DeleteCallback != nil {
		var frontA_extension_45 *A_extension_45
		if front != nil {
			frontA_extension_45, _ = front.(*A_extension_45)
		}
		stage.OnAfterA_extension_45DeleteCallback.OnAfterDelete(stage, a_extension_45, frontA_extension_45)
	}
}

func (a_extension_46 *A_extension_46) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_46CreateCallback != nil {
		stage.OnAfterA_extension_46CreateCallback.OnAfterCreate(stage, a_extension_46)
	}
}

func (a_extension_46 *A_extension_46) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_46UpdateCallback != nil {
		var frontA_extension_46 *A_extension_46
		if front != nil {
			frontA_extension_46, _ = front.(*A_extension_46)
		}
		stage.OnAfterA_extension_46UpdateCallback.OnAfterUpdate(stage, a_extension_46, frontA_extension_46)
	}
}

func (a_extension_46 *A_extension_46) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_46DeleteCallback != nil {
		var frontA_extension_46 *A_extension_46
		if front != nil {
			frontA_extension_46, _ = front.(*A_extension_46)
		}
		stage.OnAfterA_extension_46DeleteCallback.OnAfterDelete(stage, a_extension_46, frontA_extension_46)
	}
}

func (a_extension_47 *A_extension_47) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_47CreateCallback != nil {
		stage.OnAfterA_extension_47CreateCallback.OnAfterCreate(stage, a_extension_47)
	}
}

func (a_extension_47 *A_extension_47) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_47UpdateCallback != nil {
		var frontA_extension_47 *A_extension_47
		if front != nil {
			frontA_extension_47, _ = front.(*A_extension_47)
		}
		stage.OnAfterA_extension_47UpdateCallback.OnAfterUpdate(stage, a_extension_47, frontA_extension_47)
	}
}

func (a_extension_47 *A_extension_47) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_47DeleteCallback != nil {
		var frontA_extension_47 *A_extension_47
		if front != nil {
			frontA_extension_47, _ = front.(*A_extension_47)
		}
		stage.OnAfterA_extension_47DeleteCallback.OnAfterDelete(stage, a_extension_47, frontA_extension_47)
	}
}

func (a_extension_48 *A_extension_48) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_48CreateCallback != nil {
		stage.OnAfterA_extension_48CreateCallback.OnAfterCreate(stage, a_extension_48)
	}
}

func (a_extension_48 *A_extension_48) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_48UpdateCallback != nil {
		var frontA_extension_48 *A_extension_48
		if front != nil {
			frontA_extension_48, _ = front.(*A_extension_48)
		}
		stage.OnAfterA_extension_48UpdateCallback.OnAfterUpdate(stage, a_extension_48, frontA_extension_48)
	}
}

func (a_extension_48 *A_extension_48) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_48DeleteCallback != nil {
		var frontA_extension_48 *A_extension_48
		if front != nil {
			frontA_extension_48, _ = front.(*A_extension_48)
		}
		stage.OnAfterA_extension_48DeleteCallback.OnAfterDelete(stage, a_extension_48, frontA_extension_48)
	}
}

func (a_extension_49 *A_extension_49) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_49CreateCallback != nil {
		stage.OnAfterA_extension_49CreateCallback.OnAfterCreate(stage, a_extension_49)
	}
}

func (a_extension_49 *A_extension_49) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_49UpdateCallback != nil {
		var frontA_extension_49 *A_extension_49
		if front != nil {
			frontA_extension_49, _ = front.(*A_extension_49)
		}
		stage.OnAfterA_extension_49UpdateCallback.OnAfterUpdate(stage, a_extension_49, frontA_extension_49)
	}
}

func (a_extension_49 *A_extension_49) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_49DeleteCallback != nil {
		var frontA_extension_49 *A_extension_49
		if front != nil {
			frontA_extension_49, _ = front.(*A_extension_49)
		}
		stage.OnAfterA_extension_49DeleteCallback.OnAfterDelete(stage, a_extension_49, frontA_extension_49)
	}
}

func (a_extension_5 *A_extension_5) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_5CreateCallback != nil {
		stage.OnAfterA_extension_5CreateCallback.OnAfterCreate(stage, a_extension_5)
	}
}

func (a_extension_5 *A_extension_5) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_5UpdateCallback != nil {
		var frontA_extension_5 *A_extension_5
		if front != nil {
			frontA_extension_5, _ = front.(*A_extension_5)
		}
		stage.OnAfterA_extension_5UpdateCallback.OnAfterUpdate(stage, a_extension_5, frontA_extension_5)
	}
}

func (a_extension_5 *A_extension_5) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_5DeleteCallback != nil {
		var frontA_extension_5 *A_extension_5
		if front != nil {
			frontA_extension_5, _ = front.(*A_extension_5)
		}
		stage.OnAfterA_extension_5DeleteCallback.OnAfterDelete(stage, a_extension_5, frontA_extension_5)
	}
}

func (a_extension_50 *A_extension_50) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_50CreateCallback != nil {
		stage.OnAfterA_extension_50CreateCallback.OnAfterCreate(stage, a_extension_50)
	}
}

func (a_extension_50 *A_extension_50) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_50UpdateCallback != nil {
		var frontA_extension_50 *A_extension_50
		if front != nil {
			frontA_extension_50, _ = front.(*A_extension_50)
		}
		stage.OnAfterA_extension_50UpdateCallback.OnAfterUpdate(stage, a_extension_50, frontA_extension_50)
	}
}

func (a_extension_50 *A_extension_50) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_50DeleteCallback != nil {
		var frontA_extension_50 *A_extension_50
		if front != nil {
			frontA_extension_50, _ = front.(*A_extension_50)
		}
		stage.OnAfterA_extension_50DeleteCallback.OnAfterDelete(stage, a_extension_50, frontA_extension_50)
	}
}

func (a_extension_51 *A_extension_51) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_51CreateCallback != nil {
		stage.OnAfterA_extension_51CreateCallback.OnAfterCreate(stage, a_extension_51)
	}
}

func (a_extension_51 *A_extension_51) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_51UpdateCallback != nil {
		var frontA_extension_51 *A_extension_51
		if front != nil {
			frontA_extension_51, _ = front.(*A_extension_51)
		}
		stage.OnAfterA_extension_51UpdateCallback.OnAfterUpdate(stage, a_extension_51, frontA_extension_51)
	}
}

func (a_extension_51 *A_extension_51) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_51DeleteCallback != nil {
		var frontA_extension_51 *A_extension_51
		if front != nil {
			frontA_extension_51, _ = front.(*A_extension_51)
		}
		stage.OnAfterA_extension_51DeleteCallback.OnAfterDelete(stage, a_extension_51, frontA_extension_51)
	}
}

func (a_extension_52 *A_extension_52) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_52CreateCallback != nil {
		stage.OnAfterA_extension_52CreateCallback.OnAfterCreate(stage, a_extension_52)
	}
}

func (a_extension_52 *A_extension_52) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_52UpdateCallback != nil {
		var frontA_extension_52 *A_extension_52
		if front != nil {
			frontA_extension_52, _ = front.(*A_extension_52)
		}
		stage.OnAfterA_extension_52UpdateCallback.OnAfterUpdate(stage, a_extension_52, frontA_extension_52)
	}
}

func (a_extension_52 *A_extension_52) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_52DeleteCallback != nil {
		var frontA_extension_52 *A_extension_52
		if front != nil {
			frontA_extension_52, _ = front.(*A_extension_52)
		}
		stage.OnAfterA_extension_52DeleteCallback.OnAfterDelete(stage, a_extension_52, frontA_extension_52)
	}
}

func (a_extension_53 *A_extension_53) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_53CreateCallback != nil {
		stage.OnAfterA_extension_53CreateCallback.OnAfterCreate(stage, a_extension_53)
	}
}

func (a_extension_53 *A_extension_53) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_53UpdateCallback != nil {
		var frontA_extension_53 *A_extension_53
		if front != nil {
			frontA_extension_53, _ = front.(*A_extension_53)
		}
		stage.OnAfterA_extension_53UpdateCallback.OnAfterUpdate(stage, a_extension_53, frontA_extension_53)
	}
}

func (a_extension_53 *A_extension_53) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_53DeleteCallback != nil {
		var frontA_extension_53 *A_extension_53
		if front != nil {
			frontA_extension_53, _ = front.(*A_extension_53)
		}
		stage.OnAfterA_extension_53DeleteCallback.OnAfterDelete(stage, a_extension_53, frontA_extension_53)
	}
}

func (a_extension_54 *A_extension_54) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_54CreateCallback != nil {
		stage.OnAfterA_extension_54CreateCallback.OnAfterCreate(stage, a_extension_54)
	}
}

func (a_extension_54 *A_extension_54) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_54UpdateCallback != nil {
		var frontA_extension_54 *A_extension_54
		if front != nil {
			frontA_extension_54, _ = front.(*A_extension_54)
		}
		stage.OnAfterA_extension_54UpdateCallback.OnAfterUpdate(stage, a_extension_54, frontA_extension_54)
	}
}

func (a_extension_54 *A_extension_54) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_54DeleteCallback != nil {
		var frontA_extension_54 *A_extension_54
		if front != nil {
			frontA_extension_54, _ = front.(*A_extension_54)
		}
		stage.OnAfterA_extension_54DeleteCallback.OnAfterDelete(stage, a_extension_54, frontA_extension_54)
	}
}

func (a_extension_55 *A_extension_55) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_55CreateCallback != nil {
		stage.OnAfterA_extension_55CreateCallback.OnAfterCreate(stage, a_extension_55)
	}
}

func (a_extension_55 *A_extension_55) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_55UpdateCallback != nil {
		var frontA_extension_55 *A_extension_55
		if front != nil {
			frontA_extension_55, _ = front.(*A_extension_55)
		}
		stage.OnAfterA_extension_55UpdateCallback.OnAfterUpdate(stage, a_extension_55, frontA_extension_55)
	}
}

func (a_extension_55 *A_extension_55) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_55DeleteCallback != nil {
		var frontA_extension_55 *A_extension_55
		if front != nil {
			frontA_extension_55, _ = front.(*A_extension_55)
		}
		stage.OnAfterA_extension_55DeleteCallback.OnAfterDelete(stage, a_extension_55, frontA_extension_55)
	}
}

func (a_extension_56 *A_extension_56) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_56CreateCallback != nil {
		stage.OnAfterA_extension_56CreateCallback.OnAfterCreate(stage, a_extension_56)
	}
}

func (a_extension_56 *A_extension_56) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_56UpdateCallback != nil {
		var frontA_extension_56 *A_extension_56
		if front != nil {
			frontA_extension_56, _ = front.(*A_extension_56)
		}
		stage.OnAfterA_extension_56UpdateCallback.OnAfterUpdate(stage, a_extension_56, frontA_extension_56)
	}
}

func (a_extension_56 *A_extension_56) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_56DeleteCallback != nil {
		var frontA_extension_56 *A_extension_56
		if front != nil {
			frontA_extension_56, _ = front.(*A_extension_56)
		}
		stage.OnAfterA_extension_56DeleteCallback.OnAfterDelete(stage, a_extension_56, frontA_extension_56)
	}
}

func (a_extension_57 *A_extension_57) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_57CreateCallback != nil {
		stage.OnAfterA_extension_57CreateCallback.OnAfterCreate(stage, a_extension_57)
	}
}

func (a_extension_57 *A_extension_57) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_57UpdateCallback != nil {
		var frontA_extension_57 *A_extension_57
		if front != nil {
			frontA_extension_57, _ = front.(*A_extension_57)
		}
		stage.OnAfterA_extension_57UpdateCallback.OnAfterUpdate(stage, a_extension_57, frontA_extension_57)
	}
}

func (a_extension_57 *A_extension_57) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_57DeleteCallback != nil {
		var frontA_extension_57 *A_extension_57
		if front != nil {
			frontA_extension_57, _ = front.(*A_extension_57)
		}
		stage.OnAfterA_extension_57DeleteCallback.OnAfterDelete(stage, a_extension_57, frontA_extension_57)
	}
}

func (a_extension_58 *A_extension_58) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_58CreateCallback != nil {
		stage.OnAfterA_extension_58CreateCallback.OnAfterCreate(stage, a_extension_58)
	}
}

func (a_extension_58 *A_extension_58) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_58UpdateCallback != nil {
		var frontA_extension_58 *A_extension_58
		if front != nil {
			frontA_extension_58, _ = front.(*A_extension_58)
		}
		stage.OnAfterA_extension_58UpdateCallback.OnAfterUpdate(stage, a_extension_58, frontA_extension_58)
	}
}

func (a_extension_58 *A_extension_58) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_58DeleteCallback != nil {
		var frontA_extension_58 *A_extension_58
		if front != nil {
			frontA_extension_58, _ = front.(*A_extension_58)
		}
		stage.OnAfterA_extension_58DeleteCallback.OnAfterDelete(stage, a_extension_58, frontA_extension_58)
	}
}

func (a_extension_59 *A_extension_59) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_59CreateCallback != nil {
		stage.OnAfterA_extension_59CreateCallback.OnAfterCreate(stage, a_extension_59)
	}
}

func (a_extension_59 *A_extension_59) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_59UpdateCallback != nil {
		var frontA_extension_59 *A_extension_59
		if front != nil {
			frontA_extension_59, _ = front.(*A_extension_59)
		}
		stage.OnAfterA_extension_59UpdateCallback.OnAfterUpdate(stage, a_extension_59, frontA_extension_59)
	}
}

func (a_extension_59 *A_extension_59) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_59DeleteCallback != nil {
		var frontA_extension_59 *A_extension_59
		if front != nil {
			frontA_extension_59, _ = front.(*A_extension_59)
		}
		stage.OnAfterA_extension_59DeleteCallback.OnAfterDelete(stage, a_extension_59, frontA_extension_59)
	}
}

func (a_extension_6 *A_extension_6) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_6CreateCallback != nil {
		stage.OnAfterA_extension_6CreateCallback.OnAfterCreate(stage, a_extension_6)
	}
}

func (a_extension_6 *A_extension_6) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_6UpdateCallback != nil {
		var frontA_extension_6 *A_extension_6
		if front != nil {
			frontA_extension_6, _ = front.(*A_extension_6)
		}
		stage.OnAfterA_extension_6UpdateCallback.OnAfterUpdate(stage, a_extension_6, frontA_extension_6)
	}
}

func (a_extension_6 *A_extension_6) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_6DeleteCallback != nil {
		var frontA_extension_6 *A_extension_6
		if front != nil {
			frontA_extension_6, _ = front.(*A_extension_6)
		}
		stage.OnAfterA_extension_6DeleteCallback.OnAfterDelete(stage, a_extension_6, frontA_extension_6)
	}
}

func (a_extension_60 *A_extension_60) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_60CreateCallback != nil {
		stage.OnAfterA_extension_60CreateCallback.OnAfterCreate(stage, a_extension_60)
	}
}

func (a_extension_60 *A_extension_60) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_60UpdateCallback != nil {
		var frontA_extension_60 *A_extension_60
		if front != nil {
			frontA_extension_60, _ = front.(*A_extension_60)
		}
		stage.OnAfterA_extension_60UpdateCallback.OnAfterUpdate(stage, a_extension_60, frontA_extension_60)
	}
}

func (a_extension_60 *A_extension_60) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_60DeleteCallback != nil {
		var frontA_extension_60 *A_extension_60
		if front != nil {
			frontA_extension_60, _ = front.(*A_extension_60)
		}
		stage.OnAfterA_extension_60DeleteCallback.OnAfterDelete(stage, a_extension_60, frontA_extension_60)
	}
}

func (a_extension_61 *A_extension_61) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_61CreateCallback != nil {
		stage.OnAfterA_extension_61CreateCallback.OnAfterCreate(stage, a_extension_61)
	}
}

func (a_extension_61 *A_extension_61) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_61UpdateCallback != nil {
		var frontA_extension_61 *A_extension_61
		if front != nil {
			frontA_extension_61, _ = front.(*A_extension_61)
		}
		stage.OnAfterA_extension_61UpdateCallback.OnAfterUpdate(stage, a_extension_61, frontA_extension_61)
	}
}

func (a_extension_61 *A_extension_61) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_61DeleteCallback != nil {
		var frontA_extension_61 *A_extension_61
		if front != nil {
			frontA_extension_61, _ = front.(*A_extension_61)
		}
		stage.OnAfterA_extension_61DeleteCallback.OnAfterDelete(stage, a_extension_61, frontA_extension_61)
	}
}

func (a_extension_62 *A_extension_62) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_62CreateCallback != nil {
		stage.OnAfterA_extension_62CreateCallback.OnAfterCreate(stage, a_extension_62)
	}
}

func (a_extension_62 *A_extension_62) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_62UpdateCallback != nil {
		var frontA_extension_62 *A_extension_62
		if front != nil {
			frontA_extension_62, _ = front.(*A_extension_62)
		}
		stage.OnAfterA_extension_62UpdateCallback.OnAfterUpdate(stage, a_extension_62, frontA_extension_62)
	}
}

func (a_extension_62 *A_extension_62) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_62DeleteCallback != nil {
		var frontA_extension_62 *A_extension_62
		if front != nil {
			frontA_extension_62, _ = front.(*A_extension_62)
		}
		stage.OnAfterA_extension_62DeleteCallback.OnAfterDelete(stage, a_extension_62, frontA_extension_62)
	}
}

func (a_extension_63 *A_extension_63) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_63CreateCallback != nil {
		stage.OnAfterA_extension_63CreateCallback.OnAfterCreate(stage, a_extension_63)
	}
}

func (a_extension_63 *A_extension_63) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_63UpdateCallback != nil {
		var frontA_extension_63 *A_extension_63
		if front != nil {
			frontA_extension_63, _ = front.(*A_extension_63)
		}
		stage.OnAfterA_extension_63UpdateCallback.OnAfterUpdate(stage, a_extension_63, frontA_extension_63)
	}
}

func (a_extension_63 *A_extension_63) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_63DeleteCallback != nil {
		var frontA_extension_63 *A_extension_63
		if front != nil {
			frontA_extension_63, _ = front.(*A_extension_63)
		}
		stage.OnAfterA_extension_63DeleteCallback.OnAfterDelete(stage, a_extension_63, frontA_extension_63)
	}
}

func (a_extension_64 *A_extension_64) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_64CreateCallback != nil {
		stage.OnAfterA_extension_64CreateCallback.OnAfterCreate(stage, a_extension_64)
	}
}

func (a_extension_64 *A_extension_64) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_64UpdateCallback != nil {
		var frontA_extension_64 *A_extension_64
		if front != nil {
			frontA_extension_64, _ = front.(*A_extension_64)
		}
		stage.OnAfterA_extension_64UpdateCallback.OnAfterUpdate(stage, a_extension_64, frontA_extension_64)
	}
}

func (a_extension_64 *A_extension_64) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_64DeleteCallback != nil {
		var frontA_extension_64 *A_extension_64
		if front != nil {
			frontA_extension_64, _ = front.(*A_extension_64)
		}
		stage.OnAfterA_extension_64DeleteCallback.OnAfterDelete(stage, a_extension_64, frontA_extension_64)
	}
}

func (a_extension_65 *A_extension_65) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_65CreateCallback != nil {
		stage.OnAfterA_extension_65CreateCallback.OnAfterCreate(stage, a_extension_65)
	}
}

func (a_extension_65 *A_extension_65) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_65UpdateCallback != nil {
		var frontA_extension_65 *A_extension_65
		if front != nil {
			frontA_extension_65, _ = front.(*A_extension_65)
		}
		stage.OnAfterA_extension_65UpdateCallback.OnAfterUpdate(stage, a_extension_65, frontA_extension_65)
	}
}

func (a_extension_65 *A_extension_65) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_65DeleteCallback != nil {
		var frontA_extension_65 *A_extension_65
		if front != nil {
			frontA_extension_65, _ = front.(*A_extension_65)
		}
		stage.OnAfterA_extension_65DeleteCallback.OnAfterDelete(stage, a_extension_65, frontA_extension_65)
	}
}

func (a_extension_66 *A_extension_66) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_66CreateCallback != nil {
		stage.OnAfterA_extension_66CreateCallback.OnAfterCreate(stage, a_extension_66)
	}
}

func (a_extension_66 *A_extension_66) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_66UpdateCallback != nil {
		var frontA_extension_66 *A_extension_66
		if front != nil {
			frontA_extension_66, _ = front.(*A_extension_66)
		}
		stage.OnAfterA_extension_66UpdateCallback.OnAfterUpdate(stage, a_extension_66, frontA_extension_66)
	}
}

func (a_extension_66 *A_extension_66) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_66DeleteCallback != nil {
		var frontA_extension_66 *A_extension_66
		if front != nil {
			frontA_extension_66, _ = front.(*A_extension_66)
		}
		stage.OnAfterA_extension_66DeleteCallback.OnAfterDelete(stage, a_extension_66, frontA_extension_66)
	}
}

func (a_extension_67 *A_extension_67) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_67CreateCallback != nil {
		stage.OnAfterA_extension_67CreateCallback.OnAfterCreate(stage, a_extension_67)
	}
}

func (a_extension_67 *A_extension_67) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_67UpdateCallback != nil {
		var frontA_extension_67 *A_extension_67
		if front != nil {
			frontA_extension_67, _ = front.(*A_extension_67)
		}
		stage.OnAfterA_extension_67UpdateCallback.OnAfterUpdate(stage, a_extension_67, frontA_extension_67)
	}
}

func (a_extension_67 *A_extension_67) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_67DeleteCallback != nil {
		var frontA_extension_67 *A_extension_67
		if front != nil {
			frontA_extension_67, _ = front.(*A_extension_67)
		}
		stage.OnAfterA_extension_67DeleteCallback.OnAfterDelete(stage, a_extension_67, frontA_extension_67)
	}
}

func (a_extension_68 *A_extension_68) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_68CreateCallback != nil {
		stage.OnAfterA_extension_68CreateCallback.OnAfterCreate(stage, a_extension_68)
	}
}

func (a_extension_68 *A_extension_68) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_68UpdateCallback != nil {
		var frontA_extension_68 *A_extension_68
		if front != nil {
			frontA_extension_68, _ = front.(*A_extension_68)
		}
		stage.OnAfterA_extension_68UpdateCallback.OnAfterUpdate(stage, a_extension_68, frontA_extension_68)
	}
}

func (a_extension_68 *A_extension_68) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_68DeleteCallback != nil {
		var frontA_extension_68 *A_extension_68
		if front != nil {
			frontA_extension_68, _ = front.(*A_extension_68)
		}
		stage.OnAfterA_extension_68DeleteCallback.OnAfterDelete(stage, a_extension_68, frontA_extension_68)
	}
}

func (a_extension_69 *A_extension_69) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_69CreateCallback != nil {
		stage.OnAfterA_extension_69CreateCallback.OnAfterCreate(stage, a_extension_69)
	}
}

func (a_extension_69 *A_extension_69) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_69UpdateCallback != nil {
		var frontA_extension_69 *A_extension_69
		if front != nil {
			frontA_extension_69, _ = front.(*A_extension_69)
		}
		stage.OnAfterA_extension_69UpdateCallback.OnAfterUpdate(stage, a_extension_69, frontA_extension_69)
	}
}

func (a_extension_69 *A_extension_69) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_69DeleteCallback != nil {
		var frontA_extension_69 *A_extension_69
		if front != nil {
			frontA_extension_69, _ = front.(*A_extension_69)
		}
		stage.OnAfterA_extension_69DeleteCallback.OnAfterDelete(stage, a_extension_69, frontA_extension_69)
	}
}

func (a_extension_7 *A_extension_7) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_7CreateCallback != nil {
		stage.OnAfterA_extension_7CreateCallback.OnAfterCreate(stage, a_extension_7)
	}
}

func (a_extension_7 *A_extension_7) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_7UpdateCallback != nil {
		var frontA_extension_7 *A_extension_7
		if front != nil {
			frontA_extension_7, _ = front.(*A_extension_7)
		}
		stage.OnAfterA_extension_7UpdateCallback.OnAfterUpdate(stage, a_extension_7, frontA_extension_7)
	}
}

func (a_extension_7 *A_extension_7) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_7DeleteCallback != nil {
		var frontA_extension_7 *A_extension_7
		if front != nil {
			frontA_extension_7, _ = front.(*A_extension_7)
		}
		stage.OnAfterA_extension_7DeleteCallback.OnAfterDelete(stage, a_extension_7, frontA_extension_7)
	}
}

func (a_extension_70 *A_extension_70) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_70CreateCallback != nil {
		stage.OnAfterA_extension_70CreateCallback.OnAfterCreate(stage, a_extension_70)
	}
}

func (a_extension_70 *A_extension_70) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_70UpdateCallback != nil {
		var frontA_extension_70 *A_extension_70
		if front != nil {
			frontA_extension_70, _ = front.(*A_extension_70)
		}
		stage.OnAfterA_extension_70UpdateCallback.OnAfterUpdate(stage, a_extension_70, frontA_extension_70)
	}
}

func (a_extension_70 *A_extension_70) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_70DeleteCallback != nil {
		var frontA_extension_70 *A_extension_70
		if front != nil {
			frontA_extension_70, _ = front.(*A_extension_70)
		}
		stage.OnAfterA_extension_70DeleteCallback.OnAfterDelete(stage, a_extension_70, frontA_extension_70)
	}
}

func (a_extension_71 *A_extension_71) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_71CreateCallback != nil {
		stage.OnAfterA_extension_71CreateCallback.OnAfterCreate(stage, a_extension_71)
	}
}

func (a_extension_71 *A_extension_71) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_71UpdateCallback != nil {
		var frontA_extension_71 *A_extension_71
		if front != nil {
			frontA_extension_71, _ = front.(*A_extension_71)
		}
		stage.OnAfterA_extension_71UpdateCallback.OnAfterUpdate(stage, a_extension_71, frontA_extension_71)
	}
}

func (a_extension_71 *A_extension_71) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_71DeleteCallback != nil {
		var frontA_extension_71 *A_extension_71
		if front != nil {
			frontA_extension_71, _ = front.(*A_extension_71)
		}
		stage.OnAfterA_extension_71DeleteCallback.OnAfterDelete(stage, a_extension_71, frontA_extension_71)
	}
}

func (a_extension_72 *A_extension_72) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_72CreateCallback != nil {
		stage.OnAfterA_extension_72CreateCallback.OnAfterCreate(stage, a_extension_72)
	}
}

func (a_extension_72 *A_extension_72) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_72UpdateCallback != nil {
		var frontA_extension_72 *A_extension_72
		if front != nil {
			frontA_extension_72, _ = front.(*A_extension_72)
		}
		stage.OnAfterA_extension_72UpdateCallback.OnAfterUpdate(stage, a_extension_72, frontA_extension_72)
	}
}

func (a_extension_72 *A_extension_72) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_72DeleteCallback != nil {
		var frontA_extension_72 *A_extension_72
		if front != nil {
			frontA_extension_72, _ = front.(*A_extension_72)
		}
		stage.OnAfterA_extension_72DeleteCallback.OnAfterDelete(stage, a_extension_72, frontA_extension_72)
	}
}

func (a_extension_73 *A_extension_73) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_73CreateCallback != nil {
		stage.OnAfterA_extension_73CreateCallback.OnAfterCreate(stage, a_extension_73)
	}
}

func (a_extension_73 *A_extension_73) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_73UpdateCallback != nil {
		var frontA_extension_73 *A_extension_73
		if front != nil {
			frontA_extension_73, _ = front.(*A_extension_73)
		}
		stage.OnAfterA_extension_73UpdateCallback.OnAfterUpdate(stage, a_extension_73, frontA_extension_73)
	}
}

func (a_extension_73 *A_extension_73) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_73DeleteCallback != nil {
		var frontA_extension_73 *A_extension_73
		if front != nil {
			frontA_extension_73, _ = front.(*A_extension_73)
		}
		stage.OnAfterA_extension_73DeleteCallback.OnAfterDelete(stage, a_extension_73, frontA_extension_73)
	}
}

func (a_extension_74 *A_extension_74) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_74CreateCallback != nil {
		stage.OnAfterA_extension_74CreateCallback.OnAfterCreate(stage, a_extension_74)
	}
}

func (a_extension_74 *A_extension_74) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_74UpdateCallback != nil {
		var frontA_extension_74 *A_extension_74
		if front != nil {
			frontA_extension_74, _ = front.(*A_extension_74)
		}
		stage.OnAfterA_extension_74UpdateCallback.OnAfterUpdate(stage, a_extension_74, frontA_extension_74)
	}
}

func (a_extension_74 *A_extension_74) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_74DeleteCallback != nil {
		var frontA_extension_74 *A_extension_74
		if front != nil {
			frontA_extension_74, _ = front.(*A_extension_74)
		}
		stage.OnAfterA_extension_74DeleteCallback.OnAfterDelete(stage, a_extension_74, frontA_extension_74)
	}
}

func (a_extension_75 *A_extension_75) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_75CreateCallback != nil {
		stage.OnAfterA_extension_75CreateCallback.OnAfterCreate(stage, a_extension_75)
	}
}

func (a_extension_75 *A_extension_75) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_75UpdateCallback != nil {
		var frontA_extension_75 *A_extension_75
		if front != nil {
			frontA_extension_75, _ = front.(*A_extension_75)
		}
		stage.OnAfterA_extension_75UpdateCallback.OnAfterUpdate(stage, a_extension_75, frontA_extension_75)
	}
}

func (a_extension_75 *A_extension_75) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_75DeleteCallback != nil {
		var frontA_extension_75 *A_extension_75
		if front != nil {
			frontA_extension_75, _ = front.(*A_extension_75)
		}
		stage.OnAfterA_extension_75DeleteCallback.OnAfterDelete(stage, a_extension_75, frontA_extension_75)
	}
}

func (a_extension_76 *A_extension_76) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_76CreateCallback != nil {
		stage.OnAfterA_extension_76CreateCallback.OnAfterCreate(stage, a_extension_76)
	}
}

func (a_extension_76 *A_extension_76) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_76UpdateCallback != nil {
		var frontA_extension_76 *A_extension_76
		if front != nil {
			frontA_extension_76, _ = front.(*A_extension_76)
		}
		stage.OnAfterA_extension_76UpdateCallback.OnAfterUpdate(stage, a_extension_76, frontA_extension_76)
	}
}

func (a_extension_76 *A_extension_76) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_76DeleteCallback != nil {
		var frontA_extension_76 *A_extension_76
		if front != nil {
			frontA_extension_76, _ = front.(*A_extension_76)
		}
		stage.OnAfterA_extension_76DeleteCallback.OnAfterDelete(stage, a_extension_76, frontA_extension_76)
	}
}

func (a_extension_77 *A_extension_77) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_77CreateCallback != nil {
		stage.OnAfterA_extension_77CreateCallback.OnAfterCreate(stage, a_extension_77)
	}
}

func (a_extension_77 *A_extension_77) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_77UpdateCallback != nil {
		var frontA_extension_77 *A_extension_77
		if front != nil {
			frontA_extension_77, _ = front.(*A_extension_77)
		}
		stage.OnAfterA_extension_77UpdateCallback.OnAfterUpdate(stage, a_extension_77, frontA_extension_77)
	}
}

func (a_extension_77 *A_extension_77) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_77DeleteCallback != nil {
		var frontA_extension_77 *A_extension_77
		if front != nil {
			frontA_extension_77, _ = front.(*A_extension_77)
		}
		stage.OnAfterA_extension_77DeleteCallback.OnAfterDelete(stage, a_extension_77, frontA_extension_77)
	}
}

func (a_extension_78 *A_extension_78) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_78CreateCallback != nil {
		stage.OnAfterA_extension_78CreateCallback.OnAfterCreate(stage, a_extension_78)
	}
}

func (a_extension_78 *A_extension_78) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_78UpdateCallback != nil {
		var frontA_extension_78 *A_extension_78
		if front != nil {
			frontA_extension_78, _ = front.(*A_extension_78)
		}
		stage.OnAfterA_extension_78UpdateCallback.OnAfterUpdate(stage, a_extension_78, frontA_extension_78)
	}
}

func (a_extension_78 *A_extension_78) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_78DeleteCallback != nil {
		var frontA_extension_78 *A_extension_78
		if front != nil {
			frontA_extension_78, _ = front.(*A_extension_78)
		}
		stage.OnAfterA_extension_78DeleteCallback.OnAfterDelete(stage, a_extension_78, frontA_extension_78)
	}
}

func (a_extension_79 *A_extension_79) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_79CreateCallback != nil {
		stage.OnAfterA_extension_79CreateCallback.OnAfterCreate(stage, a_extension_79)
	}
}

func (a_extension_79 *A_extension_79) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_79UpdateCallback != nil {
		var frontA_extension_79 *A_extension_79
		if front != nil {
			frontA_extension_79, _ = front.(*A_extension_79)
		}
		stage.OnAfterA_extension_79UpdateCallback.OnAfterUpdate(stage, a_extension_79, frontA_extension_79)
	}
}

func (a_extension_79 *A_extension_79) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_79DeleteCallback != nil {
		var frontA_extension_79 *A_extension_79
		if front != nil {
			frontA_extension_79, _ = front.(*A_extension_79)
		}
		stage.OnAfterA_extension_79DeleteCallback.OnAfterDelete(stage, a_extension_79, frontA_extension_79)
	}
}

func (a_extension_8 *A_extension_8) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_8CreateCallback != nil {
		stage.OnAfterA_extension_8CreateCallback.OnAfterCreate(stage, a_extension_8)
	}
}

func (a_extension_8 *A_extension_8) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_8UpdateCallback != nil {
		var frontA_extension_8 *A_extension_8
		if front != nil {
			frontA_extension_8, _ = front.(*A_extension_8)
		}
		stage.OnAfterA_extension_8UpdateCallback.OnAfterUpdate(stage, a_extension_8, frontA_extension_8)
	}
}

func (a_extension_8 *A_extension_8) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_8DeleteCallback != nil {
		var frontA_extension_8 *A_extension_8
		if front != nil {
			frontA_extension_8, _ = front.(*A_extension_8)
		}
		stage.OnAfterA_extension_8DeleteCallback.OnAfterDelete(stage, a_extension_8, frontA_extension_8)
	}
}

func (a_extension_80 *A_extension_80) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_80CreateCallback != nil {
		stage.OnAfterA_extension_80CreateCallback.OnAfterCreate(stage, a_extension_80)
	}
}

func (a_extension_80 *A_extension_80) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_80UpdateCallback != nil {
		var frontA_extension_80 *A_extension_80
		if front != nil {
			frontA_extension_80, _ = front.(*A_extension_80)
		}
		stage.OnAfterA_extension_80UpdateCallback.OnAfterUpdate(stage, a_extension_80, frontA_extension_80)
	}
}

func (a_extension_80 *A_extension_80) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_80DeleteCallback != nil {
		var frontA_extension_80 *A_extension_80
		if front != nil {
			frontA_extension_80, _ = front.(*A_extension_80)
		}
		stage.OnAfterA_extension_80DeleteCallback.OnAfterDelete(stage, a_extension_80, frontA_extension_80)
	}
}

func (a_extension_81 *A_extension_81) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_81CreateCallback != nil {
		stage.OnAfterA_extension_81CreateCallback.OnAfterCreate(stage, a_extension_81)
	}
}

func (a_extension_81 *A_extension_81) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_81UpdateCallback != nil {
		var frontA_extension_81 *A_extension_81
		if front != nil {
			frontA_extension_81, _ = front.(*A_extension_81)
		}
		stage.OnAfterA_extension_81UpdateCallback.OnAfterUpdate(stage, a_extension_81, frontA_extension_81)
	}
}

func (a_extension_81 *A_extension_81) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_81DeleteCallback != nil {
		var frontA_extension_81 *A_extension_81
		if front != nil {
			frontA_extension_81, _ = front.(*A_extension_81)
		}
		stage.OnAfterA_extension_81DeleteCallback.OnAfterDelete(stage, a_extension_81, frontA_extension_81)
	}
}

func (a_extension_82 *A_extension_82) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_82CreateCallback != nil {
		stage.OnAfterA_extension_82CreateCallback.OnAfterCreate(stage, a_extension_82)
	}
}

func (a_extension_82 *A_extension_82) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_82UpdateCallback != nil {
		var frontA_extension_82 *A_extension_82
		if front != nil {
			frontA_extension_82, _ = front.(*A_extension_82)
		}
		stage.OnAfterA_extension_82UpdateCallback.OnAfterUpdate(stage, a_extension_82, frontA_extension_82)
	}
}

func (a_extension_82 *A_extension_82) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_82DeleteCallback != nil {
		var frontA_extension_82 *A_extension_82
		if front != nil {
			frontA_extension_82, _ = front.(*A_extension_82)
		}
		stage.OnAfterA_extension_82DeleteCallback.OnAfterDelete(stage, a_extension_82, frontA_extension_82)
	}
}

func (a_extension_83 *A_extension_83) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_83CreateCallback != nil {
		stage.OnAfterA_extension_83CreateCallback.OnAfterCreate(stage, a_extension_83)
	}
}

func (a_extension_83 *A_extension_83) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_83UpdateCallback != nil {
		var frontA_extension_83 *A_extension_83
		if front != nil {
			frontA_extension_83, _ = front.(*A_extension_83)
		}
		stage.OnAfterA_extension_83UpdateCallback.OnAfterUpdate(stage, a_extension_83, frontA_extension_83)
	}
}

func (a_extension_83 *A_extension_83) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_83DeleteCallback != nil {
		var frontA_extension_83 *A_extension_83
		if front != nil {
			frontA_extension_83, _ = front.(*A_extension_83)
		}
		stage.OnAfterA_extension_83DeleteCallback.OnAfterDelete(stage, a_extension_83, frontA_extension_83)
	}
}

func (a_extension_84 *A_extension_84) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_84CreateCallback != nil {
		stage.OnAfterA_extension_84CreateCallback.OnAfterCreate(stage, a_extension_84)
	}
}

func (a_extension_84 *A_extension_84) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_84UpdateCallback != nil {
		var frontA_extension_84 *A_extension_84
		if front != nil {
			frontA_extension_84, _ = front.(*A_extension_84)
		}
		stage.OnAfterA_extension_84UpdateCallback.OnAfterUpdate(stage, a_extension_84, frontA_extension_84)
	}
}

func (a_extension_84 *A_extension_84) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_84DeleteCallback != nil {
		var frontA_extension_84 *A_extension_84
		if front != nil {
			frontA_extension_84, _ = front.(*A_extension_84)
		}
		stage.OnAfterA_extension_84DeleteCallback.OnAfterDelete(stage, a_extension_84, frontA_extension_84)
	}
}

func (a_extension_85 *A_extension_85) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_85CreateCallback != nil {
		stage.OnAfterA_extension_85CreateCallback.OnAfterCreate(stage, a_extension_85)
	}
}

func (a_extension_85 *A_extension_85) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_85UpdateCallback != nil {
		var frontA_extension_85 *A_extension_85
		if front != nil {
			frontA_extension_85, _ = front.(*A_extension_85)
		}
		stage.OnAfterA_extension_85UpdateCallback.OnAfterUpdate(stage, a_extension_85, frontA_extension_85)
	}
}

func (a_extension_85 *A_extension_85) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_85DeleteCallback != nil {
		var frontA_extension_85 *A_extension_85
		if front != nil {
			frontA_extension_85, _ = front.(*A_extension_85)
		}
		stage.OnAfterA_extension_85DeleteCallback.OnAfterDelete(stage, a_extension_85, frontA_extension_85)
	}
}

func (a_extension_86 *A_extension_86) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_86CreateCallback != nil {
		stage.OnAfterA_extension_86CreateCallback.OnAfterCreate(stage, a_extension_86)
	}
}

func (a_extension_86 *A_extension_86) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_86UpdateCallback != nil {
		var frontA_extension_86 *A_extension_86
		if front != nil {
			frontA_extension_86, _ = front.(*A_extension_86)
		}
		stage.OnAfterA_extension_86UpdateCallback.OnAfterUpdate(stage, a_extension_86, frontA_extension_86)
	}
}

func (a_extension_86 *A_extension_86) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_86DeleteCallback != nil {
		var frontA_extension_86 *A_extension_86
		if front != nil {
			frontA_extension_86, _ = front.(*A_extension_86)
		}
		stage.OnAfterA_extension_86DeleteCallback.OnAfterDelete(stage, a_extension_86, frontA_extension_86)
	}
}

func (a_extension_87 *A_extension_87) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_87CreateCallback != nil {
		stage.OnAfterA_extension_87CreateCallback.OnAfterCreate(stage, a_extension_87)
	}
}

func (a_extension_87 *A_extension_87) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_87UpdateCallback != nil {
		var frontA_extension_87 *A_extension_87
		if front != nil {
			frontA_extension_87, _ = front.(*A_extension_87)
		}
		stage.OnAfterA_extension_87UpdateCallback.OnAfterUpdate(stage, a_extension_87, frontA_extension_87)
	}
}

func (a_extension_87 *A_extension_87) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_87DeleteCallback != nil {
		var frontA_extension_87 *A_extension_87
		if front != nil {
			frontA_extension_87, _ = front.(*A_extension_87)
		}
		stage.OnAfterA_extension_87DeleteCallback.OnAfterDelete(stage, a_extension_87, frontA_extension_87)
	}
}

func (a_extension_88 *A_extension_88) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_88CreateCallback != nil {
		stage.OnAfterA_extension_88CreateCallback.OnAfterCreate(stage, a_extension_88)
	}
}

func (a_extension_88 *A_extension_88) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_88UpdateCallback != nil {
		var frontA_extension_88 *A_extension_88
		if front != nil {
			frontA_extension_88, _ = front.(*A_extension_88)
		}
		stage.OnAfterA_extension_88UpdateCallback.OnAfterUpdate(stage, a_extension_88, frontA_extension_88)
	}
}

func (a_extension_88 *A_extension_88) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_88DeleteCallback != nil {
		var frontA_extension_88 *A_extension_88
		if front != nil {
			frontA_extension_88, _ = front.(*A_extension_88)
		}
		stage.OnAfterA_extension_88DeleteCallback.OnAfterDelete(stage, a_extension_88, frontA_extension_88)
	}
}

func (a_extension_89 *A_extension_89) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_89CreateCallback != nil {
		stage.OnAfterA_extension_89CreateCallback.OnAfterCreate(stage, a_extension_89)
	}
}

func (a_extension_89 *A_extension_89) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_89UpdateCallback != nil {
		var frontA_extension_89 *A_extension_89
		if front != nil {
			frontA_extension_89, _ = front.(*A_extension_89)
		}
		stage.OnAfterA_extension_89UpdateCallback.OnAfterUpdate(stage, a_extension_89, frontA_extension_89)
	}
}

func (a_extension_89 *A_extension_89) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_89DeleteCallback != nil {
		var frontA_extension_89 *A_extension_89
		if front != nil {
			frontA_extension_89, _ = front.(*A_extension_89)
		}
		stage.OnAfterA_extension_89DeleteCallback.OnAfterDelete(stage, a_extension_89, frontA_extension_89)
	}
}

func (a_extension_9 *A_extension_9) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_9CreateCallback != nil {
		stage.OnAfterA_extension_9CreateCallback.OnAfterCreate(stage, a_extension_9)
	}
}

func (a_extension_9 *A_extension_9) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_9UpdateCallback != nil {
		var frontA_extension_9 *A_extension_9
		if front != nil {
			frontA_extension_9, _ = front.(*A_extension_9)
		}
		stage.OnAfterA_extension_9UpdateCallback.OnAfterUpdate(stage, a_extension_9, frontA_extension_9)
	}
}

func (a_extension_9 *A_extension_9) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_9DeleteCallback != nil {
		var frontA_extension_9 *A_extension_9
		if front != nil {
			frontA_extension_9, _ = front.(*A_extension_9)
		}
		stage.OnAfterA_extension_9DeleteCallback.OnAfterDelete(stage, a_extension_9, frontA_extension_9)
	}
}

func (abstractairportheliportprotectionareatype *AbstractAirportHeliportProtectionAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAirportHeliportProtectionAreaTypeCreateCallback != nil {
		stage.OnAfterAbstractAirportHeliportProtectionAreaTypeCreateCallback.OnAfterCreate(stage, abstractairportheliportprotectionareatype)
	}
}

func (abstractairportheliportprotectionareatype *AbstractAirportHeliportProtectionAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAirportHeliportProtectionAreaTypeUpdateCallback != nil {
		var frontAbstractAirportHeliportProtectionAreaType *AbstractAirportHeliportProtectionAreaType
		if front != nil {
			frontAbstractAirportHeliportProtectionAreaType, _ = front.(*AbstractAirportHeliportProtectionAreaType)
		}
		stage.OnAfterAbstractAirportHeliportProtectionAreaTypeUpdateCallback.OnAfterUpdate(stage, abstractairportheliportprotectionareatype, frontAbstractAirportHeliportProtectionAreaType)
	}
}

func (abstractairportheliportprotectionareatype *AbstractAirportHeliportProtectionAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAirportHeliportProtectionAreaTypeDeleteCallback != nil {
		var frontAbstractAirportHeliportProtectionAreaType *AbstractAirportHeliportProtectionAreaType
		if front != nil {
			frontAbstractAirportHeliportProtectionAreaType, _ = front.(*AbstractAirportHeliportProtectionAreaType)
		}
		stage.OnAfterAbstractAirportHeliportProtectionAreaTypeDeleteCallback.OnAfterDelete(stage, abstractairportheliportprotectionareatype, frontAbstractAirportHeliportProtectionAreaType)
	}
}

func (abstractgroundlightsystemtype *AbstractGroundLightSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGroundLightSystemTypeCreateCallback != nil {
		stage.OnAfterAbstractGroundLightSystemTypeCreateCallback.OnAfterCreate(stage, abstractgroundlightsystemtype)
	}
}

func (abstractgroundlightsystemtype *AbstractGroundLightSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGroundLightSystemTypeUpdateCallback != nil {
		var frontAbstractGroundLightSystemType *AbstractGroundLightSystemType
		if front != nil {
			frontAbstractGroundLightSystemType, _ = front.(*AbstractGroundLightSystemType)
		}
		stage.OnAfterAbstractGroundLightSystemTypeUpdateCallback.OnAfterUpdate(stage, abstractgroundlightsystemtype, frontAbstractGroundLightSystemType)
	}
}

func (abstractgroundlightsystemtype *AbstractGroundLightSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGroundLightSystemTypeDeleteCallback != nil {
		var frontAbstractGroundLightSystemType *AbstractGroundLightSystemType
		if front != nil {
			frontAbstractGroundLightSystemType, _ = front.(*AbstractGroundLightSystemType)
		}
		stage.OnAfterAbstractGroundLightSystemTypeDeleteCallback.OnAfterDelete(stage, abstractgroundlightsystemtype, frontAbstractGroundLightSystemType)
	}
}

func (abstractmarkingtype *AbstractMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractMarkingTypeCreateCallback != nil {
		stage.OnAfterAbstractMarkingTypeCreateCallback.OnAfterCreate(stage, abstractmarkingtype)
	}
}

func (abstractmarkingtype *AbstractMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMarkingTypeUpdateCallback != nil {
		var frontAbstractMarkingType *AbstractMarkingType
		if front != nil {
			frontAbstractMarkingType, _ = front.(*AbstractMarkingType)
		}
		stage.OnAfterAbstractMarkingTypeUpdateCallback.OnAfterUpdate(stage, abstractmarkingtype, frontAbstractMarkingType)
	}
}

func (abstractmarkingtype *AbstractMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMarkingTypeDeleteCallback != nil {
		var frontAbstractMarkingType *AbstractMarkingType
		if front != nil {
			frontAbstractMarkingType, _ = front.(*AbstractMarkingType)
		}
		stage.OnAfterAbstractMarkingTypeDeleteCallback.OnAfterDelete(stage, abstractmarkingtype, frontAbstractMarkingType)
	}
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractPropertiesWithScheduleTypeCreateCallback != nil {
		stage.OnAfterAbstractPropertiesWithScheduleTypeCreateCallback.OnAfterCreate(stage, abstractpropertieswithscheduletype)
	}
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractPropertiesWithScheduleTypeUpdateCallback != nil {
		var frontAbstractPropertiesWithScheduleType *AbstractPropertiesWithScheduleType
		if front != nil {
			frontAbstractPropertiesWithScheduleType, _ = front.(*AbstractPropertiesWithScheduleType)
		}
		stage.OnAfterAbstractPropertiesWithScheduleTypeUpdateCallback.OnAfterUpdate(stage, abstractpropertieswithscheduletype, frontAbstractPropertiesWithScheduleType)
	}
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractPropertiesWithScheduleTypeDeleteCallback != nil {
		var frontAbstractPropertiesWithScheduleType *AbstractPropertiesWithScheduleType
		if front != nil {
			frontAbstractPropertiesWithScheduleType, _ = front.(*AbstractPropertiesWithScheduleType)
		}
		stage.OnAfterAbstractPropertiesWithScheduleTypeDeleteCallback.OnAfterDelete(stage, abstractpropertieswithscheduletype, frontAbstractPropertiesWithScheduleType)
	}
}

func (abstractsurfacecontaminationtype *AbstractSurfaceContaminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractSurfaceContaminationTypeCreateCallback != nil {
		stage.OnAfterAbstractSurfaceContaminationTypeCreateCallback.OnAfterCreate(stage, abstractsurfacecontaminationtype)
	}
}

func (abstractsurfacecontaminationtype *AbstractSurfaceContaminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfaceContaminationTypeUpdateCallback != nil {
		var frontAbstractSurfaceContaminationType *AbstractSurfaceContaminationType
		if front != nil {
			frontAbstractSurfaceContaminationType, _ = front.(*AbstractSurfaceContaminationType)
		}
		stage.OnAfterAbstractSurfaceContaminationTypeUpdateCallback.OnAfterUpdate(stage, abstractsurfacecontaminationtype, frontAbstractSurfaceContaminationType)
	}
}

func (abstractsurfacecontaminationtype *AbstractSurfaceContaminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfaceContaminationTypeDeleteCallback != nil {
		var frontAbstractSurfaceContaminationType *AbstractSurfaceContaminationType
		if front != nil {
			frontAbstractSurfaceContaminationType, _ = front.(*AbstractSurfaceContaminationType)
		}
		stage.OnAfterAbstractSurfaceContaminationTypeDeleteCallback.OnAfterDelete(stage, abstractsurfacecontaminationtype, frontAbstractSurfaceContaminationType)
	}
}

func (abstractusageconditiontype *AbstractUsageConditionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractUsageConditionTypeCreateCallback != nil {
		stage.OnAfterAbstractUsageConditionTypeCreateCallback.OnAfterCreate(stage, abstractusageconditiontype)
	}
}

func (abstractusageconditiontype *AbstractUsageConditionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractUsageConditionTypeUpdateCallback != nil {
		var frontAbstractUsageConditionType *AbstractUsageConditionType
		if front != nil {
			frontAbstractUsageConditionType, _ = front.(*AbstractUsageConditionType)
		}
		stage.OnAfterAbstractUsageConditionTypeUpdateCallback.OnAfterUpdate(stage, abstractusageconditiontype, frontAbstractUsageConditionType)
	}
}

func (abstractusageconditiontype *AbstractUsageConditionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractUsageConditionTypeDeleteCallback != nil {
		var frontAbstractUsageConditionType *AbstractUsageConditionType
		if front != nil {
			frontAbstractUsageConditionType, _ = front.(*AbstractUsageConditionType)
		}
		stage.OnAfterAbstractUsageConditionTypeDeleteCallback.OnAfterDelete(stage, abstractusageconditiontype, frontAbstractUsageConditionType)
	}
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftCharacteristicPropertyTypeCreateCallback != nil {
		stage.OnAfterAircraftCharacteristicPropertyTypeCreateCallback.OnAfterCreate(stage, aircraftcharacteristicpropertytype)
	}
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftCharacteristicPropertyTypeUpdateCallback != nil {
		var frontAircraftCharacteristicPropertyType *AircraftCharacteristicPropertyType
		if front != nil {
			frontAircraftCharacteristicPropertyType, _ = front.(*AircraftCharacteristicPropertyType)
		}
		stage.OnAfterAircraftCharacteristicPropertyTypeUpdateCallback.OnAfterUpdate(stage, aircraftcharacteristicpropertytype, frontAircraftCharacteristicPropertyType)
	}
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftCharacteristicPropertyTypeDeleteCallback != nil {
		var frontAircraftCharacteristicPropertyType *AircraftCharacteristicPropertyType
		if front != nil {
			frontAircraftCharacteristicPropertyType, _ = front.(*AircraftCharacteristicPropertyType)
		}
		stage.OnAfterAircraftCharacteristicPropertyTypeDeleteCallback.OnAfterDelete(stage, aircraftcharacteristicpropertytype, frontAircraftCharacteristicPropertyType)
	}
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftCharacteristicTypeCreateCallback != nil {
		stage.OnAfterAircraftCharacteristicTypeCreateCallback.OnAfterCreate(stage, aircraftcharacteristictype)
	}
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftCharacteristicTypeUpdateCallback != nil {
		var frontAircraftCharacteristicType *AircraftCharacteristicType
		if front != nil {
			frontAircraftCharacteristicType, _ = front.(*AircraftCharacteristicType)
		}
		stage.OnAfterAircraftCharacteristicTypeUpdateCallback.OnAfterUpdate(stage, aircraftcharacteristictype, frontAircraftCharacteristicType)
	}
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftCharacteristicTypeDeleteCallback != nil {
		var frontAircraftCharacteristicType *AircraftCharacteristicType
		if front != nil {
			frontAircraftCharacteristicType, _ = front.(*AircraftCharacteristicType)
		}
		stage.OnAfterAircraftCharacteristicTypeDeleteCallback.OnAfterDelete(stage, aircraftcharacteristictype, frontAircraftCharacteristicType)
	}
}

func (aircraftstandcontaminationpropertytype *AircraftStandContaminationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftStandContaminationPropertyTypeCreateCallback != nil {
		stage.OnAfterAircraftStandContaminationPropertyTypeCreateCallback.OnAfterCreate(stage, aircraftstandcontaminationpropertytype)
	}
}

func (aircraftstandcontaminationpropertytype *AircraftStandContaminationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandContaminationPropertyTypeUpdateCallback != nil {
		var frontAircraftStandContaminationPropertyType *AircraftStandContaminationPropertyType
		if front != nil {
			frontAircraftStandContaminationPropertyType, _ = front.(*AircraftStandContaminationPropertyType)
		}
		stage.OnAfterAircraftStandContaminationPropertyTypeUpdateCallback.OnAfterUpdate(stage, aircraftstandcontaminationpropertytype, frontAircraftStandContaminationPropertyType)
	}
}

func (aircraftstandcontaminationpropertytype *AircraftStandContaminationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandContaminationPropertyTypeDeleteCallback != nil {
		var frontAircraftStandContaminationPropertyType *AircraftStandContaminationPropertyType
		if front != nil {
			frontAircraftStandContaminationPropertyType, _ = front.(*AircraftStandContaminationPropertyType)
		}
		stage.OnAfterAircraftStandContaminationPropertyTypeDeleteCallback.OnAfterDelete(stage, aircraftstandcontaminationpropertytype, frontAircraftStandContaminationPropertyType)
	}
}

func (aircraftstandcontaminationtype *AircraftStandContaminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftStandContaminationTypeCreateCallback != nil {
		stage.OnAfterAircraftStandContaminationTypeCreateCallback.OnAfterCreate(stage, aircraftstandcontaminationtype)
	}
}

func (aircraftstandcontaminationtype *AircraftStandContaminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandContaminationTypeUpdateCallback != nil {
		var frontAircraftStandContaminationType *AircraftStandContaminationType
		if front != nil {
			frontAircraftStandContaminationType, _ = front.(*AircraftStandContaminationType)
		}
		stage.OnAfterAircraftStandContaminationTypeUpdateCallback.OnAfterUpdate(stage, aircraftstandcontaminationtype, frontAircraftStandContaminationType)
	}
}

func (aircraftstandcontaminationtype *AircraftStandContaminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandContaminationTypeDeleteCallback != nil {
		var frontAircraftStandContaminationType *AircraftStandContaminationType
		if front != nil {
			frontAircraftStandContaminationType, _ = front.(*AircraftStandContaminationType)
		}
		stage.OnAfterAircraftStandContaminationTypeDeleteCallback.OnAfterDelete(stage, aircraftstandcontaminationtype, frontAircraftStandContaminationType)
	}
}

func (aircraftstandpropertytype *AircraftStandPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftStandPropertyTypeCreateCallback != nil {
		stage.OnAfterAircraftStandPropertyTypeCreateCallback.OnAfterCreate(stage, aircraftstandpropertytype)
	}
}

func (aircraftstandpropertytype *AircraftStandPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandPropertyTypeUpdateCallback != nil {
		var frontAircraftStandPropertyType *AircraftStandPropertyType
		if front != nil {
			frontAircraftStandPropertyType, _ = front.(*AircraftStandPropertyType)
		}
		stage.OnAfterAircraftStandPropertyTypeUpdateCallback.OnAfterUpdate(stage, aircraftstandpropertytype, frontAircraftStandPropertyType)
	}
}

func (aircraftstandpropertytype *AircraftStandPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandPropertyTypeDeleteCallback != nil {
		var frontAircraftStandPropertyType *AircraftStandPropertyType
		if front != nil {
			frontAircraftStandPropertyType, _ = front.(*AircraftStandPropertyType)
		}
		stage.OnAfterAircraftStandPropertyTypeDeleteCallback.OnAfterDelete(stage, aircraftstandpropertytype, frontAircraftStandPropertyType)
	}
}

func (aircraftstandtimeslicepropertytype *AircraftStandTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftStandTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAircraftStandTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, aircraftstandtimeslicepropertytype)
	}
}

func (aircraftstandtimeslicepropertytype *AircraftStandTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAircraftStandTimeSlicePropertyType *AircraftStandTimeSlicePropertyType
		if front != nil {
			frontAircraftStandTimeSlicePropertyType, _ = front.(*AircraftStandTimeSlicePropertyType)
		}
		stage.OnAfterAircraftStandTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, aircraftstandtimeslicepropertytype, frontAircraftStandTimeSlicePropertyType)
	}
}

func (aircraftstandtimeslicepropertytype *AircraftStandTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAircraftStandTimeSlicePropertyType *AircraftStandTimeSlicePropertyType
		if front != nil {
			frontAircraftStandTimeSlicePropertyType, _ = front.(*AircraftStandTimeSlicePropertyType)
		}
		stage.OnAfterAircraftStandTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, aircraftstandtimeslicepropertytype, frontAircraftStandTimeSlicePropertyType)
	}
}

func (aircraftstandtimeslicetype *AircraftStandTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftStandTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAircraftStandTimeSliceTypeCreateCallback.OnAfterCreate(stage, aircraftstandtimeslicetype)
	}
}

func (aircraftstandtimeslicetype *AircraftStandTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandTimeSliceTypeUpdateCallback != nil {
		var frontAircraftStandTimeSliceType *AircraftStandTimeSliceType
		if front != nil {
			frontAircraftStandTimeSliceType, _ = front.(*AircraftStandTimeSliceType)
		}
		stage.OnAfterAircraftStandTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, aircraftstandtimeslicetype, frontAircraftStandTimeSliceType)
	}
}

func (aircraftstandtimeslicetype *AircraftStandTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandTimeSliceTypeDeleteCallback != nil {
		var frontAircraftStandTimeSliceType *AircraftStandTimeSliceType
		if front != nil {
			frontAircraftStandTimeSliceType, _ = front.(*AircraftStandTimeSliceType)
		}
		stage.OnAfterAircraftStandTimeSliceTypeDeleteCallback.OnAfterDelete(stage, aircraftstandtimeslicetype, frontAircraftStandTimeSliceType)
	}
}

func (aircraftstandtype *AircraftStandType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftStandTypeCreateCallback != nil {
		stage.OnAfterAircraftStandTypeCreateCallback.OnAfterCreate(stage, aircraftstandtype)
	}
}

func (aircraftstandtype *AircraftStandType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandTypeUpdateCallback != nil {
		var frontAircraftStandType *AircraftStandType
		if front != nil {
			frontAircraftStandType, _ = front.(*AircraftStandType)
		}
		stage.OnAfterAircraftStandTypeUpdateCallback.OnAfterUpdate(stage, aircraftstandtype, frontAircraftStandType)
	}
}

func (aircraftstandtype *AircraftStandType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftStandTypeDeleteCallback != nil {
		var frontAircraftStandType *AircraftStandType
		if front != nil {
			frontAircraftStandType, _ = front.(*AircraftStandType)
		}
		stage.OnAfterAircraftStandTypeDeleteCallback.OnAfterDelete(stage, aircraftstandtype, frontAircraftStandType)
	}
}

func (airportheliportavailabilitypropertytype *AirportHeliportAvailabilityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportAvailabilityPropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportAvailabilityPropertyTypeCreateCallback.OnAfterCreate(stage, airportheliportavailabilitypropertytype)
	}
}

func (airportheliportavailabilitypropertytype *AirportHeliportAvailabilityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportAvailabilityPropertyTypeUpdateCallback != nil {
		var frontAirportHeliportAvailabilityPropertyType *AirportHeliportAvailabilityPropertyType
		if front != nil {
			frontAirportHeliportAvailabilityPropertyType, _ = front.(*AirportHeliportAvailabilityPropertyType)
		}
		stage.OnAfterAirportHeliportAvailabilityPropertyTypeUpdateCallback.OnAfterUpdate(stage, airportheliportavailabilitypropertytype, frontAirportHeliportAvailabilityPropertyType)
	}
}

func (airportheliportavailabilitypropertytype *AirportHeliportAvailabilityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportAvailabilityPropertyTypeDeleteCallback != nil {
		var frontAirportHeliportAvailabilityPropertyType *AirportHeliportAvailabilityPropertyType
		if front != nil {
			frontAirportHeliportAvailabilityPropertyType, _ = front.(*AirportHeliportAvailabilityPropertyType)
		}
		stage.OnAfterAirportHeliportAvailabilityPropertyTypeDeleteCallback.OnAfterDelete(stage, airportheliportavailabilitypropertytype, frontAirportHeliportAvailabilityPropertyType)
	}
}

func (airportheliportavailabilitytype *AirportHeliportAvailabilityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportAvailabilityTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportAvailabilityTypeCreateCallback.OnAfterCreate(stage, airportheliportavailabilitytype)
	}
}

func (airportheliportavailabilitytype *AirportHeliportAvailabilityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportAvailabilityTypeUpdateCallback != nil {
		var frontAirportHeliportAvailabilityType *AirportHeliportAvailabilityType
		if front != nil {
			frontAirportHeliportAvailabilityType, _ = front.(*AirportHeliportAvailabilityType)
		}
		stage.OnAfterAirportHeliportAvailabilityTypeUpdateCallback.OnAfterUpdate(stage, airportheliportavailabilitytype, frontAirportHeliportAvailabilityType)
	}
}

func (airportheliportavailabilitytype *AirportHeliportAvailabilityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportAvailabilityTypeDeleteCallback != nil {
		var frontAirportHeliportAvailabilityType *AirportHeliportAvailabilityType
		if front != nil {
			frontAirportHeliportAvailabilityType, _ = front.(*AirportHeliportAvailabilityType)
		}
		stage.OnAfterAirportHeliportAvailabilityTypeDeleteCallback.OnAfterDelete(stage, airportheliportavailabilitytype, frontAirportHeliportAvailabilityType)
	}
}

func (airportheliportcollocationtimeslicepropertytype *AirportHeliportCollocationTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportCollocationTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportCollocationTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airportheliportcollocationtimeslicepropertytype)
	}
}

func (airportheliportcollocationtimeslicepropertytype *AirportHeliportCollocationTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportCollocationTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirportHeliportCollocationTimeSlicePropertyType *AirportHeliportCollocationTimeSlicePropertyType
		if front != nil {
			frontAirportHeliportCollocationTimeSlicePropertyType, _ = front.(*AirportHeliportCollocationTimeSlicePropertyType)
		}
		stage.OnAfterAirportHeliportCollocationTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airportheliportcollocationtimeslicepropertytype, frontAirportHeliportCollocationTimeSlicePropertyType)
	}
}

func (airportheliportcollocationtimeslicepropertytype *AirportHeliportCollocationTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportCollocationTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirportHeliportCollocationTimeSlicePropertyType *AirportHeliportCollocationTimeSlicePropertyType
		if front != nil {
			frontAirportHeliportCollocationTimeSlicePropertyType, _ = front.(*AirportHeliportCollocationTimeSlicePropertyType)
		}
		stage.OnAfterAirportHeliportCollocationTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airportheliportcollocationtimeslicepropertytype, frontAirportHeliportCollocationTimeSlicePropertyType)
	}
}

func (airportheliportcollocationtimeslicetype *AirportHeliportCollocationTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportCollocationTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportCollocationTimeSliceTypeCreateCallback.OnAfterCreate(stage, airportheliportcollocationtimeslicetype)
	}
}

func (airportheliportcollocationtimeslicetype *AirportHeliportCollocationTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportCollocationTimeSliceTypeUpdateCallback != nil {
		var frontAirportHeliportCollocationTimeSliceType *AirportHeliportCollocationTimeSliceType
		if front != nil {
			frontAirportHeliportCollocationTimeSliceType, _ = front.(*AirportHeliportCollocationTimeSliceType)
		}
		stage.OnAfterAirportHeliportCollocationTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airportheliportcollocationtimeslicetype, frontAirportHeliportCollocationTimeSliceType)
	}
}

func (airportheliportcollocationtimeslicetype *AirportHeliportCollocationTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportCollocationTimeSliceTypeDeleteCallback != nil {
		var frontAirportHeliportCollocationTimeSliceType *AirportHeliportCollocationTimeSliceType
		if front != nil {
			frontAirportHeliportCollocationTimeSliceType, _ = front.(*AirportHeliportCollocationTimeSliceType)
		}
		stage.OnAfterAirportHeliportCollocationTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airportheliportcollocationtimeslicetype, frontAirportHeliportCollocationTimeSliceType)
	}
}

func (airportheliportcollocationtype *AirportHeliportCollocationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportCollocationTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportCollocationTypeCreateCallback.OnAfterCreate(stage, airportheliportcollocationtype)
	}
}

func (airportheliportcollocationtype *AirportHeliportCollocationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportCollocationTypeUpdateCallback != nil {
		var frontAirportHeliportCollocationType *AirportHeliportCollocationType
		if front != nil {
			frontAirportHeliportCollocationType, _ = front.(*AirportHeliportCollocationType)
		}
		stage.OnAfterAirportHeliportCollocationTypeUpdateCallback.OnAfterUpdate(stage, airportheliportcollocationtype, frontAirportHeliportCollocationType)
	}
}

func (airportheliportcollocationtype *AirportHeliportCollocationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportCollocationTypeDeleteCallback != nil {
		var frontAirportHeliportCollocationType *AirportHeliportCollocationType
		if front != nil {
			frontAirportHeliportCollocationType, _ = front.(*AirportHeliportCollocationType)
		}
		stage.OnAfterAirportHeliportCollocationTypeDeleteCallback.OnAfterDelete(stage, airportheliportcollocationtype, frontAirportHeliportCollocationType)
	}
}

func (airportheliportcontaminationpropertytype *AirportHeliportContaminationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportContaminationPropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportContaminationPropertyTypeCreateCallback.OnAfterCreate(stage, airportheliportcontaminationpropertytype)
	}
}

func (airportheliportcontaminationpropertytype *AirportHeliportContaminationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportContaminationPropertyTypeUpdateCallback != nil {
		var frontAirportHeliportContaminationPropertyType *AirportHeliportContaminationPropertyType
		if front != nil {
			frontAirportHeliportContaminationPropertyType, _ = front.(*AirportHeliportContaminationPropertyType)
		}
		stage.OnAfterAirportHeliportContaminationPropertyTypeUpdateCallback.OnAfterUpdate(stage, airportheliportcontaminationpropertytype, frontAirportHeliportContaminationPropertyType)
	}
}

func (airportheliportcontaminationpropertytype *AirportHeliportContaminationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportContaminationPropertyTypeDeleteCallback != nil {
		var frontAirportHeliportContaminationPropertyType *AirportHeliportContaminationPropertyType
		if front != nil {
			frontAirportHeliportContaminationPropertyType, _ = front.(*AirportHeliportContaminationPropertyType)
		}
		stage.OnAfterAirportHeliportContaminationPropertyTypeDeleteCallback.OnAfterDelete(stage, airportheliportcontaminationpropertytype, frontAirportHeliportContaminationPropertyType)
	}
}

func (airportheliportcontaminationtype *AirportHeliportContaminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportContaminationTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportContaminationTypeCreateCallback.OnAfterCreate(stage, airportheliportcontaminationtype)
	}
}

func (airportheliportcontaminationtype *AirportHeliportContaminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportContaminationTypeUpdateCallback != nil {
		var frontAirportHeliportContaminationType *AirportHeliportContaminationType
		if front != nil {
			frontAirportHeliportContaminationType, _ = front.(*AirportHeliportContaminationType)
		}
		stage.OnAfterAirportHeliportContaminationTypeUpdateCallback.OnAfterUpdate(stage, airportheliportcontaminationtype, frontAirportHeliportContaminationType)
	}
}

func (airportheliportcontaminationtype *AirportHeliportContaminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportContaminationTypeDeleteCallback != nil {
		var frontAirportHeliportContaminationType *AirportHeliportContaminationType
		if front != nil {
			frontAirportHeliportContaminationType, _ = front.(*AirportHeliportContaminationType)
		}
		stage.OnAfterAirportHeliportContaminationTypeDeleteCallback.OnAfterDelete(stage, airportheliportcontaminationtype, frontAirportHeliportContaminationType)
	}
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportPropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportPropertyTypeCreateCallback.OnAfterCreate(stage, airportheliportpropertytype)
	}
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportPropertyTypeUpdateCallback != nil {
		var frontAirportHeliportPropertyType *AirportHeliportPropertyType
		if front != nil {
			frontAirportHeliportPropertyType, _ = front.(*AirportHeliportPropertyType)
		}
		stage.OnAfterAirportHeliportPropertyTypeUpdateCallback.OnAfterUpdate(stage, airportheliportpropertytype, frontAirportHeliportPropertyType)
	}
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportPropertyTypeDeleteCallback != nil {
		var frontAirportHeliportPropertyType *AirportHeliportPropertyType
		if front != nil {
			frontAirportHeliportPropertyType, _ = front.(*AirportHeliportPropertyType)
		}
		stage.OnAfterAirportHeliportPropertyTypeDeleteCallback.OnAfterDelete(stage, airportheliportpropertytype, frontAirportHeliportPropertyType)
	}
}

func (airportheliportprotectionareapropertytype *AirportHeliportProtectionAreaPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportProtectionAreaPropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportProtectionAreaPropertyTypeCreateCallback.OnAfterCreate(stage, airportheliportprotectionareapropertytype)
	}
}

func (airportheliportprotectionareapropertytype *AirportHeliportProtectionAreaPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportProtectionAreaPropertyTypeUpdateCallback != nil {
		var frontAirportHeliportProtectionAreaPropertyType *AirportHeliportProtectionAreaPropertyType
		if front != nil {
			frontAirportHeliportProtectionAreaPropertyType, _ = front.(*AirportHeliportProtectionAreaPropertyType)
		}
		stage.OnAfterAirportHeliportProtectionAreaPropertyTypeUpdateCallback.OnAfterUpdate(stage, airportheliportprotectionareapropertytype, frontAirportHeliportProtectionAreaPropertyType)
	}
}

func (airportheliportprotectionareapropertytype *AirportHeliportProtectionAreaPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportProtectionAreaPropertyTypeDeleteCallback != nil {
		var frontAirportHeliportProtectionAreaPropertyType *AirportHeliportProtectionAreaPropertyType
		if front != nil {
			frontAirportHeliportProtectionAreaPropertyType, _ = front.(*AirportHeliportProtectionAreaPropertyType)
		}
		stage.OnAfterAirportHeliportProtectionAreaPropertyTypeDeleteCallback.OnAfterDelete(stage, airportheliportprotectionareapropertytype, frontAirportHeliportProtectionAreaPropertyType)
	}
}

func (airportheliportresponsibilityorganisationpropertytype *AirportHeliportResponsibilityOrganisationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportResponsibilityOrganisationPropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportResponsibilityOrganisationPropertyTypeCreateCallback.OnAfterCreate(stage, airportheliportresponsibilityorganisationpropertytype)
	}
}

func (airportheliportresponsibilityorganisationpropertytype *AirportHeliportResponsibilityOrganisationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportResponsibilityOrganisationPropertyTypeUpdateCallback != nil {
		var frontAirportHeliportResponsibilityOrganisationPropertyType *AirportHeliportResponsibilityOrganisationPropertyType
		if front != nil {
			frontAirportHeliportResponsibilityOrganisationPropertyType, _ = front.(*AirportHeliportResponsibilityOrganisationPropertyType)
		}
		stage.OnAfterAirportHeliportResponsibilityOrganisationPropertyTypeUpdateCallback.OnAfterUpdate(stage, airportheliportresponsibilityorganisationpropertytype, frontAirportHeliportResponsibilityOrganisationPropertyType)
	}
}

func (airportheliportresponsibilityorganisationpropertytype *AirportHeliportResponsibilityOrganisationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportResponsibilityOrganisationPropertyTypeDeleteCallback != nil {
		var frontAirportHeliportResponsibilityOrganisationPropertyType *AirportHeliportResponsibilityOrganisationPropertyType
		if front != nil {
			frontAirportHeliportResponsibilityOrganisationPropertyType, _ = front.(*AirportHeliportResponsibilityOrganisationPropertyType)
		}
		stage.OnAfterAirportHeliportResponsibilityOrganisationPropertyTypeDeleteCallback.OnAfterDelete(stage, airportheliportresponsibilityorganisationpropertytype, frontAirportHeliportResponsibilityOrganisationPropertyType)
	}
}

func (airportheliportresponsibilityorganisationtype *AirportHeliportResponsibilityOrganisationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportResponsibilityOrganisationTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportResponsibilityOrganisationTypeCreateCallback.OnAfterCreate(stage, airportheliportresponsibilityorganisationtype)
	}
}

func (airportheliportresponsibilityorganisationtype *AirportHeliportResponsibilityOrganisationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportResponsibilityOrganisationTypeUpdateCallback != nil {
		var frontAirportHeliportResponsibilityOrganisationType *AirportHeliportResponsibilityOrganisationType
		if front != nil {
			frontAirportHeliportResponsibilityOrganisationType, _ = front.(*AirportHeliportResponsibilityOrganisationType)
		}
		stage.OnAfterAirportHeliportResponsibilityOrganisationTypeUpdateCallback.OnAfterUpdate(stage, airportheliportresponsibilityorganisationtype, frontAirportHeliportResponsibilityOrganisationType)
	}
}

func (airportheliportresponsibilityorganisationtype *AirportHeliportResponsibilityOrganisationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportResponsibilityOrganisationTypeDeleteCallback != nil {
		var frontAirportHeliportResponsibilityOrganisationType *AirportHeliportResponsibilityOrganisationType
		if front != nil {
			frontAirportHeliportResponsibilityOrganisationType, _ = front.(*AirportHeliportResponsibilityOrganisationType)
		}
		stage.OnAfterAirportHeliportResponsibilityOrganisationTypeDeleteCallback.OnAfterDelete(stage, airportheliportresponsibilityorganisationtype, frontAirportHeliportResponsibilityOrganisationType)
	}
}

func (airportheliporttimeslicepropertytype *AirportHeliportTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airportheliporttimeslicepropertytype)
	}
}

func (airportheliporttimeslicepropertytype *AirportHeliportTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirportHeliportTimeSlicePropertyType *AirportHeliportTimeSlicePropertyType
		if front != nil {
			frontAirportHeliportTimeSlicePropertyType, _ = front.(*AirportHeliportTimeSlicePropertyType)
		}
		stage.OnAfterAirportHeliportTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airportheliporttimeslicepropertytype, frontAirportHeliportTimeSlicePropertyType)
	}
}

func (airportheliporttimeslicepropertytype *AirportHeliportTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirportHeliportTimeSlicePropertyType *AirportHeliportTimeSlicePropertyType
		if front != nil {
			frontAirportHeliportTimeSlicePropertyType, _ = front.(*AirportHeliportTimeSlicePropertyType)
		}
		stage.OnAfterAirportHeliportTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airportheliporttimeslicepropertytype, frontAirportHeliportTimeSlicePropertyType)
	}
}

func (airportheliporttimeslicetype *AirportHeliportTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportTimeSliceTypeCreateCallback.OnAfterCreate(stage, airportheliporttimeslicetype)
	}
}

func (airportheliporttimeslicetype *AirportHeliportTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportTimeSliceTypeUpdateCallback != nil {
		var frontAirportHeliportTimeSliceType *AirportHeliportTimeSliceType
		if front != nil {
			frontAirportHeliportTimeSliceType, _ = front.(*AirportHeliportTimeSliceType)
		}
		stage.OnAfterAirportHeliportTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airportheliporttimeslicetype, frontAirportHeliportTimeSliceType)
	}
}

func (airportheliporttimeslicetype *AirportHeliportTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportTimeSliceTypeDeleteCallback != nil {
		var frontAirportHeliportTimeSliceType *AirportHeliportTimeSliceType
		if front != nil {
			frontAirportHeliportTimeSliceType, _ = front.(*AirportHeliportTimeSliceType)
		}
		stage.OnAfterAirportHeliportTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airportheliporttimeslicetype, frontAirportHeliportTimeSliceType)
	}
}

func (airportheliporttype *AirportHeliportType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportTypeCreateCallback.OnAfterCreate(stage, airportheliporttype)
	}
}

func (airportheliporttype *AirportHeliportType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportTypeUpdateCallback != nil {
		var frontAirportHeliportType *AirportHeliportType
		if front != nil {
			frontAirportHeliportType, _ = front.(*AirportHeliportType)
		}
		stage.OnAfterAirportHeliportTypeUpdateCallback.OnAfterUpdate(stage, airportheliporttype, frontAirportHeliportType)
	}
}

func (airportheliporttype *AirportHeliportType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportTypeDeleteCallback != nil {
		var frontAirportHeliportType *AirportHeliportType
		if front != nil {
			frontAirportHeliportType, _ = front.(*AirportHeliportType)
		}
		stage.OnAfterAirportHeliportTypeDeleteCallback.OnAfterDelete(stage, airportheliporttype, frontAirportHeliportType)
	}
}

func (airportheliportusagepropertytype *AirportHeliportUsagePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportUsagePropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportUsagePropertyTypeCreateCallback.OnAfterCreate(stage, airportheliportusagepropertytype)
	}
}

func (airportheliportusagepropertytype *AirportHeliportUsagePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportUsagePropertyTypeUpdateCallback != nil {
		var frontAirportHeliportUsagePropertyType *AirportHeliportUsagePropertyType
		if front != nil {
			frontAirportHeliportUsagePropertyType, _ = front.(*AirportHeliportUsagePropertyType)
		}
		stage.OnAfterAirportHeliportUsagePropertyTypeUpdateCallback.OnAfterUpdate(stage, airportheliportusagepropertytype, frontAirportHeliportUsagePropertyType)
	}
}

func (airportheliportusagepropertytype *AirportHeliportUsagePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportUsagePropertyTypeDeleteCallback != nil {
		var frontAirportHeliportUsagePropertyType *AirportHeliportUsagePropertyType
		if front != nil {
			frontAirportHeliportUsagePropertyType, _ = front.(*AirportHeliportUsagePropertyType)
		}
		stage.OnAfterAirportHeliportUsagePropertyTypeDeleteCallback.OnAfterDelete(stage, airportheliportusagepropertytype, frontAirportHeliportUsagePropertyType)
	}
}

func (airportheliportusagetype *AirportHeliportUsageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportUsageTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportUsageTypeCreateCallback.OnAfterCreate(stage, airportheliportusagetype)
	}
}

func (airportheliportusagetype *AirportHeliportUsageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportUsageTypeUpdateCallback != nil {
		var frontAirportHeliportUsageType *AirportHeliportUsageType
		if front != nil {
			frontAirportHeliportUsageType, _ = front.(*AirportHeliportUsageType)
		}
		stage.OnAfterAirportHeliportUsageTypeUpdateCallback.OnAfterUpdate(stage, airportheliportusagetype, frontAirportHeliportUsageType)
	}
}

func (airportheliportusagetype *AirportHeliportUsageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportUsageTypeDeleteCallback != nil {
		var frontAirportHeliportUsageType *AirportHeliportUsageType
		if front != nil {
			frontAirportHeliportUsageType, _ = front.(*AirportHeliportUsageType)
		}
		stage.OnAfterAirportHeliportUsageTypeDeleteCallback.OnAfterDelete(stage, airportheliportusagetype, frontAirportHeliportUsageType)
	}
}

func (airporthotspottimeslicepropertytype *AirportHotSpotTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHotSpotTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHotSpotTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airporthotspottimeslicepropertytype)
	}
}

func (airporthotspottimeslicepropertytype *AirportHotSpotTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHotSpotTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirportHotSpotTimeSlicePropertyType *AirportHotSpotTimeSlicePropertyType
		if front != nil {
			frontAirportHotSpotTimeSlicePropertyType, _ = front.(*AirportHotSpotTimeSlicePropertyType)
		}
		stage.OnAfterAirportHotSpotTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airporthotspottimeslicepropertytype, frontAirportHotSpotTimeSlicePropertyType)
	}
}

func (airporthotspottimeslicepropertytype *AirportHotSpotTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHotSpotTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirportHotSpotTimeSlicePropertyType *AirportHotSpotTimeSlicePropertyType
		if front != nil {
			frontAirportHotSpotTimeSlicePropertyType, _ = front.(*AirportHotSpotTimeSlicePropertyType)
		}
		stage.OnAfterAirportHotSpotTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airporthotspottimeslicepropertytype, frontAirportHotSpotTimeSlicePropertyType)
	}
}

func (airporthotspottimeslicetype *AirportHotSpotTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHotSpotTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirportHotSpotTimeSliceTypeCreateCallback.OnAfterCreate(stage, airporthotspottimeslicetype)
	}
}

func (airporthotspottimeslicetype *AirportHotSpotTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHotSpotTimeSliceTypeUpdateCallback != nil {
		var frontAirportHotSpotTimeSliceType *AirportHotSpotTimeSliceType
		if front != nil {
			frontAirportHotSpotTimeSliceType, _ = front.(*AirportHotSpotTimeSliceType)
		}
		stage.OnAfterAirportHotSpotTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airporthotspottimeslicetype, frontAirportHotSpotTimeSliceType)
	}
}

func (airporthotspottimeslicetype *AirportHotSpotTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHotSpotTimeSliceTypeDeleteCallback != nil {
		var frontAirportHotSpotTimeSliceType *AirportHotSpotTimeSliceType
		if front != nil {
			frontAirportHotSpotTimeSliceType, _ = front.(*AirportHotSpotTimeSliceType)
		}
		stage.OnAfterAirportHotSpotTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airporthotspottimeslicetype, frontAirportHotSpotTimeSliceType)
	}
}

func (airporthotspottype *AirportHotSpotType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHotSpotTypeCreateCallback != nil {
		stage.OnAfterAirportHotSpotTypeCreateCallback.OnAfterCreate(stage, airporthotspottype)
	}
}

func (airporthotspottype *AirportHotSpotType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHotSpotTypeUpdateCallback != nil {
		var frontAirportHotSpotType *AirportHotSpotType
		if front != nil {
			frontAirportHotSpotType, _ = front.(*AirportHotSpotType)
		}
		stage.OnAfterAirportHotSpotTypeUpdateCallback.OnAfterUpdate(stage, airporthotspottype, frontAirportHotSpotType)
	}
}

func (airporthotspottype *AirportHotSpotType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHotSpotTypeDeleteCallback != nil {
		var frontAirportHotSpotType *AirportHotSpotType
		if front != nil {
			frontAirportHotSpotType, _ = front.(*AirportHotSpotType)
		}
		stage.OnAfterAirportHotSpotTypeDeleteCallback.OnAfterDelete(stage, airporthotspottype, frontAirportHotSpotType)
	}
}

func (airportprotectionareamarkingtimeslicepropertytype *AirportProtectionAreaMarkingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportProtectionAreaMarkingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirportProtectionAreaMarkingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airportprotectionareamarkingtimeslicepropertytype)
	}
}

func (airportprotectionareamarkingtimeslicepropertytype *AirportProtectionAreaMarkingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportProtectionAreaMarkingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirportProtectionAreaMarkingTimeSlicePropertyType *AirportProtectionAreaMarkingTimeSlicePropertyType
		if front != nil {
			frontAirportProtectionAreaMarkingTimeSlicePropertyType, _ = front.(*AirportProtectionAreaMarkingTimeSlicePropertyType)
		}
		stage.OnAfterAirportProtectionAreaMarkingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airportprotectionareamarkingtimeslicepropertytype, frontAirportProtectionAreaMarkingTimeSlicePropertyType)
	}
}

func (airportprotectionareamarkingtimeslicepropertytype *AirportProtectionAreaMarkingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportProtectionAreaMarkingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirportProtectionAreaMarkingTimeSlicePropertyType *AirportProtectionAreaMarkingTimeSlicePropertyType
		if front != nil {
			frontAirportProtectionAreaMarkingTimeSlicePropertyType, _ = front.(*AirportProtectionAreaMarkingTimeSlicePropertyType)
		}
		stage.OnAfterAirportProtectionAreaMarkingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airportprotectionareamarkingtimeslicepropertytype, frontAirportProtectionAreaMarkingTimeSlicePropertyType)
	}
}

func (airportprotectionareamarkingtimeslicetype *AirportProtectionAreaMarkingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportProtectionAreaMarkingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirportProtectionAreaMarkingTimeSliceTypeCreateCallback.OnAfterCreate(stage, airportprotectionareamarkingtimeslicetype)
	}
}

func (airportprotectionareamarkingtimeslicetype *AirportProtectionAreaMarkingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportProtectionAreaMarkingTimeSliceTypeUpdateCallback != nil {
		var frontAirportProtectionAreaMarkingTimeSliceType *AirportProtectionAreaMarkingTimeSliceType
		if front != nil {
			frontAirportProtectionAreaMarkingTimeSliceType, _ = front.(*AirportProtectionAreaMarkingTimeSliceType)
		}
		stage.OnAfterAirportProtectionAreaMarkingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airportprotectionareamarkingtimeslicetype, frontAirportProtectionAreaMarkingTimeSliceType)
	}
}

func (airportprotectionareamarkingtimeslicetype *AirportProtectionAreaMarkingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportProtectionAreaMarkingTimeSliceTypeDeleteCallback != nil {
		var frontAirportProtectionAreaMarkingTimeSliceType *AirportProtectionAreaMarkingTimeSliceType
		if front != nil {
			frontAirportProtectionAreaMarkingTimeSliceType, _ = front.(*AirportProtectionAreaMarkingTimeSliceType)
		}
		stage.OnAfterAirportProtectionAreaMarkingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airportprotectionareamarkingtimeslicetype, frontAirportProtectionAreaMarkingTimeSliceType)
	}
}

func (airportprotectionareamarkingtype *AirportProtectionAreaMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportProtectionAreaMarkingTypeCreateCallback != nil {
		stage.OnAfterAirportProtectionAreaMarkingTypeCreateCallback.OnAfterCreate(stage, airportprotectionareamarkingtype)
	}
}

func (airportprotectionareamarkingtype *AirportProtectionAreaMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportProtectionAreaMarkingTypeUpdateCallback != nil {
		var frontAirportProtectionAreaMarkingType *AirportProtectionAreaMarkingType
		if front != nil {
			frontAirportProtectionAreaMarkingType, _ = front.(*AirportProtectionAreaMarkingType)
		}
		stage.OnAfterAirportProtectionAreaMarkingTypeUpdateCallback.OnAfterUpdate(stage, airportprotectionareamarkingtype, frontAirportProtectionAreaMarkingType)
	}
}

func (airportprotectionareamarkingtype *AirportProtectionAreaMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportProtectionAreaMarkingTypeDeleteCallback != nil {
		var frontAirportProtectionAreaMarkingType *AirportProtectionAreaMarkingType
		if front != nil {
			frontAirportProtectionAreaMarkingType, _ = front.(*AirportProtectionAreaMarkingType)
		}
		stage.OnAfterAirportProtectionAreaMarkingTypeDeleteCallback.OnAfterDelete(stage, airportprotectionareamarkingtype, frontAirportProtectionAreaMarkingType)
	}
}

func (airportsuppliesservicepropertytype *AirportSuppliesServicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportSuppliesServicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirportSuppliesServicePropertyTypeCreateCallback.OnAfterCreate(stage, airportsuppliesservicepropertytype)
	}
}

func (airportsuppliesservicepropertytype *AirportSuppliesServicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportSuppliesServicePropertyTypeUpdateCallback != nil {
		var frontAirportSuppliesServicePropertyType *AirportSuppliesServicePropertyType
		if front != nil {
			frontAirportSuppliesServicePropertyType, _ = front.(*AirportSuppliesServicePropertyType)
		}
		stage.OnAfterAirportSuppliesServicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airportsuppliesservicepropertytype, frontAirportSuppliesServicePropertyType)
	}
}

func (airportsuppliesservicepropertytype *AirportSuppliesServicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportSuppliesServicePropertyTypeDeleteCallback != nil {
		var frontAirportSuppliesServicePropertyType *AirportSuppliesServicePropertyType
		if front != nil {
			frontAirportSuppliesServicePropertyType, _ = front.(*AirportSuppliesServicePropertyType)
		}
		stage.OnAfterAirportSuppliesServicePropertyTypeDeleteCallback.OnAfterDelete(stage, airportsuppliesservicepropertytype, frontAirportSuppliesServicePropertyType)
	}
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAltimeterSourcePropertyTypeCreateCallback != nil {
		stage.OnAfterAltimeterSourcePropertyTypeCreateCallback.OnAfterCreate(stage, altimetersourcepropertytype)
	}
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourcePropertyTypeUpdateCallback != nil {
		var frontAltimeterSourcePropertyType *AltimeterSourcePropertyType
		if front != nil {
			frontAltimeterSourcePropertyType, _ = front.(*AltimeterSourcePropertyType)
		}
		stage.OnAfterAltimeterSourcePropertyTypeUpdateCallback.OnAfterUpdate(stage, altimetersourcepropertytype, frontAltimeterSourcePropertyType)
	}
}

func (altimetersourcepropertytype *AltimeterSourcePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourcePropertyTypeDeleteCallback != nil {
		var frontAltimeterSourcePropertyType *AltimeterSourcePropertyType
		if front != nil {
			frontAltimeterSourcePropertyType, _ = front.(*AltimeterSourcePropertyType)
		}
		stage.OnAfterAltimeterSourcePropertyTypeDeleteCallback.OnAfterDelete(stage, altimetersourcepropertytype, frontAltimeterSourcePropertyType)
	}
}

func (altimetersourcestatuspropertytype *AltimeterSourceStatusPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAltimeterSourceStatusPropertyTypeCreateCallback != nil {
		stage.OnAfterAltimeterSourceStatusPropertyTypeCreateCallback.OnAfterCreate(stage, altimetersourcestatuspropertytype)
	}
}

func (altimetersourcestatuspropertytype *AltimeterSourceStatusPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceStatusPropertyTypeUpdateCallback != nil {
		var frontAltimeterSourceStatusPropertyType *AltimeterSourceStatusPropertyType
		if front != nil {
			frontAltimeterSourceStatusPropertyType, _ = front.(*AltimeterSourceStatusPropertyType)
		}
		stage.OnAfterAltimeterSourceStatusPropertyTypeUpdateCallback.OnAfterUpdate(stage, altimetersourcestatuspropertytype, frontAltimeterSourceStatusPropertyType)
	}
}

func (altimetersourcestatuspropertytype *AltimeterSourceStatusPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceStatusPropertyTypeDeleteCallback != nil {
		var frontAltimeterSourceStatusPropertyType *AltimeterSourceStatusPropertyType
		if front != nil {
			frontAltimeterSourceStatusPropertyType, _ = front.(*AltimeterSourceStatusPropertyType)
		}
		stage.OnAfterAltimeterSourceStatusPropertyTypeDeleteCallback.OnAfterDelete(stage, altimetersourcestatuspropertytype, frontAltimeterSourceStatusPropertyType)
	}
}

func (altimetersourcestatustype *AltimeterSourceStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAltimeterSourceStatusTypeCreateCallback != nil {
		stage.OnAfterAltimeterSourceStatusTypeCreateCallback.OnAfterCreate(stage, altimetersourcestatustype)
	}
}

func (altimetersourcestatustype *AltimeterSourceStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceStatusTypeUpdateCallback != nil {
		var frontAltimeterSourceStatusType *AltimeterSourceStatusType
		if front != nil {
			frontAltimeterSourceStatusType, _ = front.(*AltimeterSourceStatusType)
		}
		stage.OnAfterAltimeterSourceStatusTypeUpdateCallback.OnAfterUpdate(stage, altimetersourcestatustype, frontAltimeterSourceStatusType)
	}
}

func (altimetersourcestatustype *AltimeterSourceStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceStatusTypeDeleteCallback != nil {
		var frontAltimeterSourceStatusType *AltimeterSourceStatusType
		if front != nil {
			frontAltimeterSourceStatusType, _ = front.(*AltimeterSourceStatusType)
		}
		stage.OnAfterAltimeterSourceStatusTypeDeleteCallback.OnAfterDelete(stage, altimetersourcestatustype, frontAltimeterSourceStatusType)
	}
}

func (altimetersourcetimeslicepropertytype *AltimeterSourceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAltimeterSourceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAltimeterSourceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, altimetersourcetimeslicepropertytype)
	}
}

func (altimetersourcetimeslicepropertytype *AltimeterSourceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAltimeterSourceTimeSlicePropertyType *AltimeterSourceTimeSlicePropertyType
		if front != nil {
			frontAltimeterSourceTimeSlicePropertyType, _ = front.(*AltimeterSourceTimeSlicePropertyType)
		}
		stage.OnAfterAltimeterSourceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, altimetersourcetimeslicepropertytype, frontAltimeterSourceTimeSlicePropertyType)
	}
}

func (altimetersourcetimeslicepropertytype *AltimeterSourceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAltimeterSourceTimeSlicePropertyType *AltimeterSourceTimeSlicePropertyType
		if front != nil {
			frontAltimeterSourceTimeSlicePropertyType, _ = front.(*AltimeterSourceTimeSlicePropertyType)
		}
		stage.OnAfterAltimeterSourceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, altimetersourcetimeslicepropertytype, frontAltimeterSourceTimeSlicePropertyType)
	}
}

func (altimetersourcetimeslicetype *AltimeterSourceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAltimeterSourceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAltimeterSourceTimeSliceTypeCreateCallback.OnAfterCreate(stage, altimetersourcetimeslicetype)
	}
}

func (altimetersourcetimeslicetype *AltimeterSourceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceTimeSliceTypeUpdateCallback != nil {
		var frontAltimeterSourceTimeSliceType *AltimeterSourceTimeSliceType
		if front != nil {
			frontAltimeterSourceTimeSliceType, _ = front.(*AltimeterSourceTimeSliceType)
		}
		stage.OnAfterAltimeterSourceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, altimetersourcetimeslicetype, frontAltimeterSourceTimeSliceType)
	}
}

func (altimetersourcetimeslicetype *AltimeterSourceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceTimeSliceTypeDeleteCallback != nil {
		var frontAltimeterSourceTimeSliceType *AltimeterSourceTimeSliceType
		if front != nil {
			frontAltimeterSourceTimeSliceType, _ = front.(*AltimeterSourceTimeSliceType)
		}
		stage.OnAfterAltimeterSourceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, altimetersourcetimeslicetype, frontAltimeterSourceTimeSliceType)
	}
}

func (altimetersourcetype *AltimeterSourceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAltimeterSourceTypeCreateCallback != nil {
		stage.OnAfterAltimeterSourceTypeCreateCallback.OnAfterCreate(stage, altimetersourcetype)
	}
}

func (altimetersourcetype *AltimeterSourceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceTypeUpdateCallback != nil {
		var frontAltimeterSourceType *AltimeterSourceType
		if front != nil {
			frontAltimeterSourceType, _ = front.(*AltimeterSourceType)
		}
		stage.OnAfterAltimeterSourceTypeUpdateCallback.OnAfterUpdate(stage, altimetersourcetype, frontAltimeterSourceType)
	}
}

func (altimetersourcetype *AltimeterSourceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltimeterSourceTypeDeleteCallback != nil {
		var frontAltimeterSourceType *AltimeterSourceType
		if front != nil {
			frontAltimeterSourceType, _ = front.(*AltimeterSourceType)
		}
		stage.OnAfterAltimeterSourceTypeDeleteCallback.OnAfterDelete(stage, altimetersourcetype, frontAltimeterSourceType)
	}
}

func (approachlightingsystemtimeslicepropertytype *ApproachLightingSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachLightingSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterApproachLightingSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, approachlightingsystemtimeslicepropertytype)
	}
}

func (approachlightingsystemtimeslicepropertytype *ApproachLightingSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachLightingSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontApproachLightingSystemTimeSlicePropertyType *ApproachLightingSystemTimeSlicePropertyType
		if front != nil {
			frontApproachLightingSystemTimeSlicePropertyType, _ = front.(*ApproachLightingSystemTimeSlicePropertyType)
		}
		stage.OnAfterApproachLightingSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, approachlightingsystemtimeslicepropertytype, frontApproachLightingSystemTimeSlicePropertyType)
	}
}

func (approachlightingsystemtimeslicepropertytype *ApproachLightingSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachLightingSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontApproachLightingSystemTimeSlicePropertyType *ApproachLightingSystemTimeSlicePropertyType
		if front != nil {
			frontApproachLightingSystemTimeSlicePropertyType, _ = front.(*ApproachLightingSystemTimeSlicePropertyType)
		}
		stage.OnAfterApproachLightingSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, approachlightingsystemtimeslicepropertytype, frontApproachLightingSystemTimeSlicePropertyType)
	}
}

func (approachlightingsystemtimeslicetype *ApproachLightingSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachLightingSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterApproachLightingSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, approachlightingsystemtimeslicetype)
	}
}

func (approachlightingsystemtimeslicetype *ApproachLightingSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachLightingSystemTimeSliceTypeUpdateCallback != nil {
		var frontApproachLightingSystemTimeSliceType *ApproachLightingSystemTimeSliceType
		if front != nil {
			frontApproachLightingSystemTimeSliceType, _ = front.(*ApproachLightingSystemTimeSliceType)
		}
		stage.OnAfterApproachLightingSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, approachlightingsystemtimeslicetype, frontApproachLightingSystemTimeSliceType)
	}
}

func (approachlightingsystemtimeslicetype *ApproachLightingSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachLightingSystemTimeSliceTypeDeleteCallback != nil {
		var frontApproachLightingSystemTimeSliceType *ApproachLightingSystemTimeSliceType
		if front != nil {
			frontApproachLightingSystemTimeSliceType, _ = front.(*ApproachLightingSystemTimeSliceType)
		}
		stage.OnAfterApproachLightingSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, approachlightingsystemtimeslicetype, frontApproachLightingSystemTimeSliceType)
	}
}

func (approachlightingsystemtype *ApproachLightingSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachLightingSystemTypeCreateCallback != nil {
		stage.OnAfterApproachLightingSystemTypeCreateCallback.OnAfterCreate(stage, approachlightingsystemtype)
	}
}

func (approachlightingsystemtype *ApproachLightingSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachLightingSystemTypeUpdateCallback != nil {
		var frontApproachLightingSystemType *ApproachLightingSystemType
		if front != nil {
			frontApproachLightingSystemType, _ = front.(*ApproachLightingSystemType)
		}
		stage.OnAfterApproachLightingSystemTypeUpdateCallback.OnAfterUpdate(stage, approachlightingsystemtype, frontApproachLightingSystemType)
	}
}

func (approachlightingsystemtype *ApproachLightingSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachLightingSystemTypeDeleteCallback != nil {
		var frontApproachLightingSystemType *ApproachLightingSystemType
		if front != nil {
			frontApproachLightingSystemType, _ = front.(*ApproachLightingSystemType)
		}
		stage.OnAfterApproachLightingSystemTypeDeleteCallback.OnAfterDelete(stage, approachlightingsystemtype, frontApproachLightingSystemType)
	}
}

func (apronareaavailabilitypropertytype *ApronAreaAvailabilityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronAreaAvailabilityPropertyTypeCreateCallback != nil {
		stage.OnAfterApronAreaAvailabilityPropertyTypeCreateCallback.OnAfterCreate(stage, apronareaavailabilitypropertytype)
	}
}

func (apronareaavailabilitypropertytype *ApronAreaAvailabilityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronAreaAvailabilityPropertyTypeUpdateCallback != nil {
		var frontApronAreaAvailabilityPropertyType *ApronAreaAvailabilityPropertyType
		if front != nil {
			frontApronAreaAvailabilityPropertyType, _ = front.(*ApronAreaAvailabilityPropertyType)
		}
		stage.OnAfterApronAreaAvailabilityPropertyTypeUpdateCallback.OnAfterUpdate(stage, apronareaavailabilitypropertytype, frontApronAreaAvailabilityPropertyType)
	}
}

func (apronareaavailabilitypropertytype *ApronAreaAvailabilityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronAreaAvailabilityPropertyTypeDeleteCallback != nil {
		var frontApronAreaAvailabilityPropertyType *ApronAreaAvailabilityPropertyType
		if front != nil {
			frontApronAreaAvailabilityPropertyType, _ = front.(*ApronAreaAvailabilityPropertyType)
		}
		stage.OnAfterApronAreaAvailabilityPropertyTypeDeleteCallback.OnAfterDelete(stage, apronareaavailabilitypropertytype, frontApronAreaAvailabilityPropertyType)
	}
}

func (apronareaavailabilitytype *ApronAreaAvailabilityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronAreaAvailabilityTypeCreateCallback != nil {
		stage.OnAfterApronAreaAvailabilityTypeCreateCallback.OnAfterCreate(stage, apronareaavailabilitytype)
	}
}

func (apronareaavailabilitytype *ApronAreaAvailabilityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronAreaAvailabilityTypeUpdateCallback != nil {
		var frontApronAreaAvailabilityType *ApronAreaAvailabilityType
		if front != nil {
			frontApronAreaAvailabilityType, _ = front.(*ApronAreaAvailabilityType)
		}
		stage.OnAfterApronAreaAvailabilityTypeUpdateCallback.OnAfterUpdate(stage, apronareaavailabilitytype, frontApronAreaAvailabilityType)
	}
}

func (apronareaavailabilitytype *ApronAreaAvailabilityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronAreaAvailabilityTypeDeleteCallback != nil {
		var frontApronAreaAvailabilityType *ApronAreaAvailabilityType
		if front != nil {
			frontApronAreaAvailabilityType, _ = front.(*ApronAreaAvailabilityType)
		}
		stage.OnAfterApronAreaAvailabilityTypeDeleteCallback.OnAfterDelete(stage, apronareaavailabilitytype, frontApronAreaAvailabilityType)
	}
}

func (apronareausagepropertytype *ApronAreaUsagePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronAreaUsagePropertyTypeCreateCallback != nil {
		stage.OnAfterApronAreaUsagePropertyTypeCreateCallback.OnAfterCreate(stage, apronareausagepropertytype)
	}
}

func (apronareausagepropertytype *ApronAreaUsagePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronAreaUsagePropertyTypeUpdateCallback != nil {
		var frontApronAreaUsagePropertyType *ApronAreaUsagePropertyType
		if front != nil {
			frontApronAreaUsagePropertyType, _ = front.(*ApronAreaUsagePropertyType)
		}
		stage.OnAfterApronAreaUsagePropertyTypeUpdateCallback.OnAfterUpdate(stage, apronareausagepropertytype, frontApronAreaUsagePropertyType)
	}
}

func (apronareausagepropertytype *ApronAreaUsagePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronAreaUsagePropertyTypeDeleteCallback != nil {
		var frontApronAreaUsagePropertyType *ApronAreaUsagePropertyType
		if front != nil {
			frontApronAreaUsagePropertyType, _ = front.(*ApronAreaUsagePropertyType)
		}
		stage.OnAfterApronAreaUsagePropertyTypeDeleteCallback.OnAfterDelete(stage, apronareausagepropertytype, frontApronAreaUsagePropertyType)
	}
}

func (apronareausagetype *ApronAreaUsageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronAreaUsageTypeCreateCallback != nil {
		stage.OnAfterApronAreaUsageTypeCreateCallback.OnAfterCreate(stage, apronareausagetype)
	}
}

func (apronareausagetype *ApronAreaUsageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronAreaUsageTypeUpdateCallback != nil {
		var frontApronAreaUsageType *ApronAreaUsageType
		if front != nil {
			frontApronAreaUsageType, _ = front.(*ApronAreaUsageType)
		}
		stage.OnAfterApronAreaUsageTypeUpdateCallback.OnAfterUpdate(stage, apronareausagetype, frontApronAreaUsageType)
	}
}

func (apronareausagetype *ApronAreaUsageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronAreaUsageTypeDeleteCallback != nil {
		var frontApronAreaUsageType *ApronAreaUsageType
		if front != nil {
			frontApronAreaUsageType, _ = front.(*ApronAreaUsageType)
		}
		stage.OnAfterApronAreaUsageTypeDeleteCallback.OnAfterDelete(stage, apronareausagetype, frontApronAreaUsageType)
	}
}

func (aproncontaminationpropertytype *ApronContaminationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronContaminationPropertyTypeCreateCallback != nil {
		stage.OnAfterApronContaminationPropertyTypeCreateCallback.OnAfterCreate(stage, aproncontaminationpropertytype)
	}
}

func (aproncontaminationpropertytype *ApronContaminationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronContaminationPropertyTypeUpdateCallback != nil {
		var frontApronContaminationPropertyType *ApronContaminationPropertyType
		if front != nil {
			frontApronContaminationPropertyType, _ = front.(*ApronContaminationPropertyType)
		}
		stage.OnAfterApronContaminationPropertyTypeUpdateCallback.OnAfterUpdate(stage, aproncontaminationpropertytype, frontApronContaminationPropertyType)
	}
}

func (aproncontaminationpropertytype *ApronContaminationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronContaminationPropertyTypeDeleteCallback != nil {
		var frontApronContaminationPropertyType *ApronContaminationPropertyType
		if front != nil {
			frontApronContaminationPropertyType, _ = front.(*ApronContaminationPropertyType)
		}
		stage.OnAfterApronContaminationPropertyTypeDeleteCallback.OnAfterDelete(stage, aproncontaminationpropertytype, frontApronContaminationPropertyType)
	}
}

func (aproncontaminationtype *ApronContaminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronContaminationTypeCreateCallback != nil {
		stage.OnAfterApronContaminationTypeCreateCallback.OnAfterCreate(stage, aproncontaminationtype)
	}
}

func (aproncontaminationtype *ApronContaminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronContaminationTypeUpdateCallback != nil {
		var frontApronContaminationType *ApronContaminationType
		if front != nil {
			frontApronContaminationType, _ = front.(*ApronContaminationType)
		}
		stage.OnAfterApronContaminationTypeUpdateCallback.OnAfterUpdate(stage, aproncontaminationtype, frontApronContaminationType)
	}
}

func (aproncontaminationtype *ApronContaminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronContaminationTypeDeleteCallback != nil {
		var frontApronContaminationType *ApronContaminationType
		if front != nil {
			frontApronContaminationType, _ = front.(*ApronContaminationType)
		}
		stage.OnAfterApronContaminationTypeDeleteCallback.OnAfterDelete(stage, aproncontaminationtype, frontApronContaminationType)
	}
}

func (apronelementpropertytype *ApronElementPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronElementPropertyTypeCreateCallback != nil {
		stage.OnAfterApronElementPropertyTypeCreateCallback.OnAfterCreate(stage, apronelementpropertytype)
	}
}

func (apronelementpropertytype *ApronElementPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronElementPropertyTypeUpdateCallback != nil {
		var frontApronElementPropertyType *ApronElementPropertyType
		if front != nil {
			frontApronElementPropertyType, _ = front.(*ApronElementPropertyType)
		}
		stage.OnAfterApronElementPropertyTypeUpdateCallback.OnAfterUpdate(stage, apronelementpropertytype, frontApronElementPropertyType)
	}
}

func (apronelementpropertytype *ApronElementPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronElementPropertyTypeDeleteCallback != nil {
		var frontApronElementPropertyType *ApronElementPropertyType
		if front != nil {
			frontApronElementPropertyType, _ = front.(*ApronElementPropertyType)
		}
		stage.OnAfterApronElementPropertyTypeDeleteCallback.OnAfterDelete(stage, apronelementpropertytype, frontApronElementPropertyType)
	}
}

func (apronelementtimeslicepropertytype *ApronElementTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronElementTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterApronElementTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, apronelementtimeslicepropertytype)
	}
}

func (apronelementtimeslicepropertytype *ApronElementTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronElementTimeSlicePropertyTypeUpdateCallback != nil {
		var frontApronElementTimeSlicePropertyType *ApronElementTimeSlicePropertyType
		if front != nil {
			frontApronElementTimeSlicePropertyType, _ = front.(*ApronElementTimeSlicePropertyType)
		}
		stage.OnAfterApronElementTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, apronelementtimeslicepropertytype, frontApronElementTimeSlicePropertyType)
	}
}

func (apronelementtimeslicepropertytype *ApronElementTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronElementTimeSlicePropertyTypeDeleteCallback != nil {
		var frontApronElementTimeSlicePropertyType *ApronElementTimeSlicePropertyType
		if front != nil {
			frontApronElementTimeSlicePropertyType, _ = front.(*ApronElementTimeSlicePropertyType)
		}
		stage.OnAfterApronElementTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, apronelementtimeslicepropertytype, frontApronElementTimeSlicePropertyType)
	}
}

func (apronelementtimeslicetype *ApronElementTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronElementTimeSliceTypeCreateCallback != nil {
		stage.OnAfterApronElementTimeSliceTypeCreateCallback.OnAfterCreate(stage, apronelementtimeslicetype)
	}
}

func (apronelementtimeslicetype *ApronElementTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronElementTimeSliceTypeUpdateCallback != nil {
		var frontApronElementTimeSliceType *ApronElementTimeSliceType
		if front != nil {
			frontApronElementTimeSliceType, _ = front.(*ApronElementTimeSliceType)
		}
		stage.OnAfterApronElementTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, apronelementtimeslicetype, frontApronElementTimeSliceType)
	}
}

func (apronelementtimeslicetype *ApronElementTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronElementTimeSliceTypeDeleteCallback != nil {
		var frontApronElementTimeSliceType *ApronElementTimeSliceType
		if front != nil {
			frontApronElementTimeSliceType, _ = front.(*ApronElementTimeSliceType)
		}
		stage.OnAfterApronElementTimeSliceTypeDeleteCallback.OnAfterDelete(stage, apronelementtimeslicetype, frontApronElementTimeSliceType)
	}
}

func (apronelementtype *ApronElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronElementTypeCreateCallback != nil {
		stage.OnAfterApronElementTypeCreateCallback.OnAfterCreate(stage, apronelementtype)
	}
}

func (apronelementtype *ApronElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronElementTypeUpdateCallback != nil {
		var frontApronElementType *ApronElementType
		if front != nil {
			frontApronElementType, _ = front.(*ApronElementType)
		}
		stage.OnAfterApronElementTypeUpdateCallback.OnAfterUpdate(stage, apronelementtype, frontApronElementType)
	}
}

func (apronelementtype *ApronElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronElementTypeDeleteCallback != nil {
		var frontApronElementType *ApronElementType
		if front != nil {
			frontApronElementType, _ = front.(*ApronElementType)
		}
		stage.OnAfterApronElementTypeDeleteCallback.OnAfterDelete(stage, apronelementtype, frontApronElementType)
	}
}

func (apronlightsystemtimeslicepropertytype *ApronLightSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronLightSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterApronLightSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, apronlightsystemtimeslicepropertytype)
	}
}

func (apronlightsystemtimeslicepropertytype *ApronLightSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronLightSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontApronLightSystemTimeSlicePropertyType *ApronLightSystemTimeSlicePropertyType
		if front != nil {
			frontApronLightSystemTimeSlicePropertyType, _ = front.(*ApronLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterApronLightSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, apronlightsystemtimeslicepropertytype, frontApronLightSystemTimeSlicePropertyType)
	}
}

func (apronlightsystemtimeslicepropertytype *ApronLightSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronLightSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontApronLightSystemTimeSlicePropertyType *ApronLightSystemTimeSlicePropertyType
		if front != nil {
			frontApronLightSystemTimeSlicePropertyType, _ = front.(*ApronLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterApronLightSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, apronlightsystemtimeslicepropertytype, frontApronLightSystemTimeSlicePropertyType)
	}
}

func (apronlightsystemtimeslicetype *ApronLightSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronLightSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterApronLightSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, apronlightsystemtimeslicetype)
	}
}

func (apronlightsystemtimeslicetype *ApronLightSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronLightSystemTimeSliceTypeUpdateCallback != nil {
		var frontApronLightSystemTimeSliceType *ApronLightSystemTimeSliceType
		if front != nil {
			frontApronLightSystemTimeSliceType, _ = front.(*ApronLightSystemTimeSliceType)
		}
		stage.OnAfterApronLightSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, apronlightsystemtimeslicetype, frontApronLightSystemTimeSliceType)
	}
}

func (apronlightsystemtimeslicetype *ApronLightSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronLightSystemTimeSliceTypeDeleteCallback != nil {
		var frontApronLightSystemTimeSliceType *ApronLightSystemTimeSliceType
		if front != nil {
			frontApronLightSystemTimeSliceType, _ = front.(*ApronLightSystemTimeSliceType)
		}
		stage.OnAfterApronLightSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, apronlightsystemtimeslicetype, frontApronLightSystemTimeSliceType)
	}
}

func (apronlightsystemtype *ApronLightSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronLightSystemTypeCreateCallback != nil {
		stage.OnAfterApronLightSystemTypeCreateCallback.OnAfterCreate(stage, apronlightsystemtype)
	}
}

func (apronlightsystemtype *ApronLightSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronLightSystemTypeUpdateCallback != nil {
		var frontApronLightSystemType *ApronLightSystemType
		if front != nil {
			frontApronLightSystemType, _ = front.(*ApronLightSystemType)
		}
		stage.OnAfterApronLightSystemTypeUpdateCallback.OnAfterUpdate(stage, apronlightsystemtype, frontApronLightSystemType)
	}
}

func (apronlightsystemtype *ApronLightSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronLightSystemTypeDeleteCallback != nil {
		var frontApronLightSystemType *ApronLightSystemType
		if front != nil {
			frontApronLightSystemType, _ = front.(*ApronLightSystemType)
		}
		stage.OnAfterApronLightSystemTypeDeleteCallback.OnAfterDelete(stage, apronlightsystemtype, frontApronLightSystemType)
	}
}

func (apronmarkingtimeslicepropertytype *ApronMarkingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronMarkingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterApronMarkingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, apronmarkingtimeslicepropertytype)
	}
}

func (apronmarkingtimeslicepropertytype *ApronMarkingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronMarkingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontApronMarkingTimeSlicePropertyType *ApronMarkingTimeSlicePropertyType
		if front != nil {
			frontApronMarkingTimeSlicePropertyType, _ = front.(*ApronMarkingTimeSlicePropertyType)
		}
		stage.OnAfterApronMarkingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, apronmarkingtimeslicepropertytype, frontApronMarkingTimeSlicePropertyType)
	}
}

func (apronmarkingtimeslicepropertytype *ApronMarkingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronMarkingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontApronMarkingTimeSlicePropertyType *ApronMarkingTimeSlicePropertyType
		if front != nil {
			frontApronMarkingTimeSlicePropertyType, _ = front.(*ApronMarkingTimeSlicePropertyType)
		}
		stage.OnAfterApronMarkingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, apronmarkingtimeslicepropertytype, frontApronMarkingTimeSlicePropertyType)
	}
}

func (apronmarkingtimeslicetype *ApronMarkingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronMarkingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterApronMarkingTimeSliceTypeCreateCallback.OnAfterCreate(stage, apronmarkingtimeslicetype)
	}
}

func (apronmarkingtimeslicetype *ApronMarkingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronMarkingTimeSliceTypeUpdateCallback != nil {
		var frontApronMarkingTimeSliceType *ApronMarkingTimeSliceType
		if front != nil {
			frontApronMarkingTimeSliceType, _ = front.(*ApronMarkingTimeSliceType)
		}
		stage.OnAfterApronMarkingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, apronmarkingtimeslicetype, frontApronMarkingTimeSliceType)
	}
}

func (apronmarkingtimeslicetype *ApronMarkingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronMarkingTimeSliceTypeDeleteCallback != nil {
		var frontApronMarkingTimeSliceType *ApronMarkingTimeSliceType
		if front != nil {
			frontApronMarkingTimeSliceType, _ = front.(*ApronMarkingTimeSliceType)
		}
		stage.OnAfterApronMarkingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, apronmarkingtimeslicetype, frontApronMarkingTimeSliceType)
	}
}

func (apronmarkingtype *ApronMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronMarkingTypeCreateCallback != nil {
		stage.OnAfterApronMarkingTypeCreateCallback.OnAfterCreate(stage, apronmarkingtype)
	}
}

func (apronmarkingtype *ApronMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronMarkingTypeUpdateCallback != nil {
		var frontApronMarkingType *ApronMarkingType
		if front != nil {
			frontApronMarkingType, _ = front.(*ApronMarkingType)
		}
		stage.OnAfterApronMarkingTypeUpdateCallback.OnAfterUpdate(stage, apronmarkingtype, frontApronMarkingType)
	}
}

func (apronmarkingtype *ApronMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronMarkingTypeDeleteCallback != nil {
		var frontApronMarkingType *ApronMarkingType
		if front != nil {
			frontApronMarkingType, _ = front.(*ApronMarkingType)
		}
		stage.OnAfterApronMarkingTypeDeleteCallback.OnAfterDelete(stage, apronmarkingtype, frontApronMarkingType)
	}
}

func (apronpropertytype *ApronPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronPropertyTypeCreateCallback != nil {
		stage.OnAfterApronPropertyTypeCreateCallback.OnAfterCreate(stage, apronpropertytype)
	}
}

func (apronpropertytype *ApronPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronPropertyTypeUpdateCallback != nil {
		var frontApronPropertyType *ApronPropertyType
		if front != nil {
			frontApronPropertyType, _ = front.(*ApronPropertyType)
		}
		stage.OnAfterApronPropertyTypeUpdateCallback.OnAfterUpdate(stage, apronpropertytype, frontApronPropertyType)
	}
}

func (apronpropertytype *ApronPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronPropertyTypeDeleteCallback != nil {
		var frontApronPropertyType *ApronPropertyType
		if front != nil {
			frontApronPropertyType, _ = front.(*ApronPropertyType)
		}
		stage.OnAfterApronPropertyTypeDeleteCallback.OnAfterDelete(stage, apronpropertytype, frontApronPropertyType)
	}
}

func (aprontimeslicepropertytype *ApronTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterApronTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, aprontimeslicepropertytype)
	}
}

func (aprontimeslicepropertytype *ApronTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronTimeSlicePropertyTypeUpdateCallback != nil {
		var frontApronTimeSlicePropertyType *ApronTimeSlicePropertyType
		if front != nil {
			frontApronTimeSlicePropertyType, _ = front.(*ApronTimeSlicePropertyType)
		}
		stage.OnAfterApronTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, aprontimeslicepropertytype, frontApronTimeSlicePropertyType)
	}
}

func (aprontimeslicepropertytype *ApronTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronTimeSlicePropertyTypeDeleteCallback != nil {
		var frontApronTimeSlicePropertyType *ApronTimeSlicePropertyType
		if front != nil {
			frontApronTimeSlicePropertyType, _ = front.(*ApronTimeSlicePropertyType)
		}
		stage.OnAfterApronTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, aprontimeslicepropertytype, frontApronTimeSlicePropertyType)
	}
}

func (aprontimeslicetype *ApronTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronTimeSliceTypeCreateCallback != nil {
		stage.OnAfterApronTimeSliceTypeCreateCallback.OnAfterCreate(stage, aprontimeslicetype)
	}
}

func (aprontimeslicetype *ApronTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronTimeSliceTypeUpdateCallback != nil {
		var frontApronTimeSliceType *ApronTimeSliceType
		if front != nil {
			frontApronTimeSliceType, _ = front.(*ApronTimeSliceType)
		}
		stage.OnAfterApronTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, aprontimeslicetype, frontApronTimeSliceType)
	}
}

func (aprontimeslicetype *ApronTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronTimeSliceTypeDeleteCallback != nil {
		var frontApronTimeSliceType *ApronTimeSliceType
		if front != nil {
			frontApronTimeSliceType, _ = front.(*ApronTimeSliceType)
		}
		stage.OnAfterApronTimeSliceTypeDeleteCallback.OnAfterDelete(stage, aprontimeslicetype, frontApronTimeSliceType)
	}
}

func (aprontype *ApronType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApronTypeCreateCallback != nil {
		stage.OnAfterApronTypeCreateCallback.OnAfterCreate(stage, aprontype)
	}
}

func (aprontype *ApronType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronTypeUpdateCallback != nil {
		var frontApronType *ApronType
		if front != nil {
			frontApronType, _ = front.(*ApronType)
		}
		stage.OnAfterApronTypeUpdateCallback.OnAfterUpdate(stage, aprontype, frontApronType)
	}
}

func (aprontype *ApronType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApronTypeDeleteCallback != nil {
		var frontApronType *ApronType
		if front != nil {
			frontApronType, _ = front.(*ApronType)
		}
		stage.OnAfterApronTypeDeleteCallback.OnAfterDelete(stage, aprontype, frontApronType)
	}
}

func (arrestinggeartimeslicepropertytype *ArrestingGearTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrestingGearTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterArrestingGearTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, arrestinggeartimeslicepropertytype)
	}
}

func (arrestinggeartimeslicepropertytype *ArrestingGearTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrestingGearTimeSlicePropertyTypeUpdateCallback != nil {
		var frontArrestingGearTimeSlicePropertyType *ArrestingGearTimeSlicePropertyType
		if front != nil {
			frontArrestingGearTimeSlicePropertyType, _ = front.(*ArrestingGearTimeSlicePropertyType)
		}
		stage.OnAfterArrestingGearTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, arrestinggeartimeslicepropertytype, frontArrestingGearTimeSlicePropertyType)
	}
}

func (arrestinggeartimeslicepropertytype *ArrestingGearTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrestingGearTimeSlicePropertyTypeDeleteCallback != nil {
		var frontArrestingGearTimeSlicePropertyType *ArrestingGearTimeSlicePropertyType
		if front != nil {
			frontArrestingGearTimeSlicePropertyType, _ = front.(*ArrestingGearTimeSlicePropertyType)
		}
		stage.OnAfterArrestingGearTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, arrestinggeartimeslicepropertytype, frontArrestingGearTimeSlicePropertyType)
	}
}

func (arrestinggeartimeslicetype *ArrestingGearTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrestingGearTimeSliceTypeCreateCallback != nil {
		stage.OnAfterArrestingGearTimeSliceTypeCreateCallback.OnAfterCreate(stage, arrestinggeartimeslicetype)
	}
}

func (arrestinggeartimeslicetype *ArrestingGearTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrestingGearTimeSliceTypeUpdateCallback != nil {
		var frontArrestingGearTimeSliceType *ArrestingGearTimeSliceType
		if front != nil {
			frontArrestingGearTimeSliceType, _ = front.(*ArrestingGearTimeSliceType)
		}
		stage.OnAfterArrestingGearTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, arrestinggeartimeslicetype, frontArrestingGearTimeSliceType)
	}
}

func (arrestinggeartimeslicetype *ArrestingGearTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrestingGearTimeSliceTypeDeleteCallback != nil {
		var frontArrestingGearTimeSliceType *ArrestingGearTimeSliceType
		if front != nil {
			frontArrestingGearTimeSliceType, _ = front.(*ArrestingGearTimeSliceType)
		}
		stage.OnAfterArrestingGearTimeSliceTypeDeleteCallback.OnAfterDelete(stage, arrestinggeartimeslicetype, frontArrestingGearTimeSliceType)
	}
}

func (arrestinggeartype *ArrestingGearType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrestingGearTypeCreateCallback != nil {
		stage.OnAfterArrestingGearTypeCreateCallback.OnAfterCreate(stage, arrestinggeartype)
	}
}

func (arrestinggeartype *ArrestingGearType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrestingGearTypeUpdateCallback != nil {
		var frontArrestingGearType *ArrestingGearType
		if front != nil {
			frontArrestingGearType, _ = front.(*ArrestingGearType)
		}
		stage.OnAfterArrestingGearTypeUpdateCallback.OnAfterUpdate(stage, arrestinggeartype, frontArrestingGearType)
	}
}

func (arrestinggeartype *ArrestingGearType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrestingGearTypeDeleteCallback != nil {
		var frontArrestingGearType *ArrestingGearType
		if front != nil {
			frontArrestingGearType, _ = front.(*ArrestingGearType)
		}
		stage.OnAfterArrestingGearTypeDeleteCallback.OnAfterDelete(stage, arrestinggeartype, frontArrestingGearType)
	}
}

func (citypropertytype *CityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCityPropertyTypeCreateCallback != nil {
		stage.OnAfterCityPropertyTypeCreateCallback.OnAfterCreate(stage, citypropertytype)
	}
}

func (citypropertytype *CityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCityPropertyTypeUpdateCallback != nil {
		var frontCityPropertyType *CityPropertyType
		if front != nil {
			frontCityPropertyType, _ = front.(*CityPropertyType)
		}
		stage.OnAfterCityPropertyTypeUpdateCallback.OnAfterUpdate(stage, citypropertytype, frontCityPropertyType)
	}
}

func (citypropertytype *CityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCityPropertyTypeDeleteCallback != nil {
		var frontCityPropertyType *CityPropertyType
		if front != nil {
			frontCityPropertyType, _ = front.(*CityPropertyType)
		}
		stage.OnAfterCityPropertyTypeDeleteCallback.OnAfterDelete(stage, citypropertytype, frontCityPropertyType)
	}
}

func (citytype *CityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCityTypeCreateCallback != nil {
		stage.OnAfterCityTypeCreateCallback.OnAfterCreate(stage, citytype)
	}
}

func (citytype *CityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCityTypeUpdateCallback != nil {
		var frontCityType *CityType
		if front != nil {
			frontCityType, _ = front.(*CityType)
		}
		stage.OnAfterCityTypeUpdateCallback.OnAfterUpdate(stage, citytype, frontCityType)
	}
}

func (citytype *CityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCityTypeDeleteCallback != nil {
		var frontCityType *CityType
		if front != nil {
			frontCityType, _ = front.(*CityType)
		}
		stage.OnAfterCityTypeDeleteCallback.OnAfterDelete(stage, citytype, frontCityType)
	}
}

func (conditioncombinationpropertytype *ConditionCombinationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConditionCombinationPropertyTypeCreateCallback != nil {
		stage.OnAfterConditionCombinationPropertyTypeCreateCallback.OnAfterCreate(stage, conditioncombinationpropertytype)
	}
}

func (conditioncombinationpropertytype *ConditionCombinationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConditionCombinationPropertyTypeUpdateCallback != nil {
		var frontConditionCombinationPropertyType *ConditionCombinationPropertyType
		if front != nil {
			frontConditionCombinationPropertyType, _ = front.(*ConditionCombinationPropertyType)
		}
		stage.OnAfterConditionCombinationPropertyTypeUpdateCallback.OnAfterUpdate(stage, conditioncombinationpropertytype, frontConditionCombinationPropertyType)
	}
}

func (conditioncombinationpropertytype *ConditionCombinationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConditionCombinationPropertyTypeDeleteCallback != nil {
		var frontConditionCombinationPropertyType *ConditionCombinationPropertyType
		if front != nil {
			frontConditionCombinationPropertyType, _ = front.(*ConditionCombinationPropertyType)
		}
		stage.OnAfterConditionCombinationPropertyTypeDeleteCallback.OnAfterDelete(stage, conditioncombinationpropertytype, frontConditionCombinationPropertyType)
	}
}

func (conditioncombinationtype *ConditionCombinationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConditionCombinationTypeCreateCallback != nil {
		stage.OnAfterConditionCombinationTypeCreateCallback.OnAfterCreate(stage, conditioncombinationtype)
	}
}

func (conditioncombinationtype *ConditionCombinationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConditionCombinationTypeUpdateCallback != nil {
		var frontConditionCombinationType *ConditionCombinationType
		if front != nil {
			frontConditionCombinationType, _ = front.(*ConditionCombinationType)
		}
		stage.OnAfterConditionCombinationTypeUpdateCallback.OnAfterUpdate(stage, conditioncombinationtype, frontConditionCombinationType)
	}
}

func (conditioncombinationtype *ConditionCombinationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConditionCombinationTypeDeleteCallback != nil {
		var frontConditionCombinationType *ConditionCombinationType
		if front != nil {
			frontConditionCombinationType, _ = front.(*ConditionCombinationType)
		}
		stage.OnAfterConditionCombinationTypeDeleteCallback.OnAfterDelete(stage, conditioncombinationtype, frontConditionCombinationType)
	}
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterContactInformationPropertyTypeCreateCallback != nil {
		stage.OnAfterContactInformationPropertyTypeCreateCallback.OnAfterCreate(stage, contactinformationpropertytype)
	}
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterContactInformationPropertyTypeUpdateCallback != nil {
		var frontContactInformationPropertyType *ContactInformationPropertyType
		if front != nil {
			frontContactInformationPropertyType, _ = front.(*ContactInformationPropertyType)
		}
		stage.OnAfterContactInformationPropertyTypeUpdateCallback.OnAfterUpdate(stage, contactinformationpropertytype, frontContactInformationPropertyType)
	}
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterContactInformationPropertyTypeDeleteCallback != nil {
		var frontContactInformationPropertyType *ContactInformationPropertyType
		if front != nil {
			frontContactInformationPropertyType, _ = front.(*ContactInformationPropertyType)
		}
		stage.OnAfterContactInformationPropertyTypeDeleteCallback.OnAfterDelete(stage, contactinformationpropertytype, frontContactInformationPropertyType)
	}
}

func (contactinformationtype *ContactInformationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterContactInformationTypeCreateCallback != nil {
		stage.OnAfterContactInformationTypeCreateCallback.OnAfterCreate(stage, contactinformationtype)
	}
}

func (contactinformationtype *ContactInformationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterContactInformationTypeUpdateCallback != nil {
		var frontContactInformationType *ContactInformationType
		if front != nil {
			frontContactInformationType, _ = front.(*ContactInformationType)
		}
		stage.OnAfterContactInformationTypeUpdateCallback.OnAfterUpdate(stage, contactinformationtype, frontContactInformationType)
	}
}

func (contactinformationtype *ContactInformationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterContactInformationTypeDeleteCallback != nil {
		var frontContactInformationType *ContactInformationType
		if front != nil {
			frontContactInformationType, _ = front.(*ContactInformationType)
		}
		stage.OnAfterContactInformationTypeDeleteCallback.OnAfterDelete(stage, contactinformationtype, frontContactInformationType)
	}
}

func (curvetype *CurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurveTypeCreateCallback != nil {
		stage.OnAfterCurveTypeCreateCallback.OnAfterCreate(stage, curvetype)
	}
}

func (curvetype *CurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveTypeUpdateCallback != nil {
		var frontCurveType *CurveType
		if front != nil {
			frontCurveType, _ = front.(*CurveType)
		}
		stage.OnAfterCurveTypeUpdateCallback.OnAfterUpdate(stage, curvetype, frontCurveType)
	}
}

func (curvetype *CurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveTypeDeleteCallback != nil {
		var frontCurveType *CurveType
		if front != nil {
			frontCurveType, _ = front.(*CurveType)
		}
		stage.OnAfterCurveTypeDeleteCallback.OnAfterDelete(stage, curvetype, frontCurveType)
	}
}

func (deicingareamarkingtimeslicepropertytype *DeicingAreaMarkingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeicingAreaMarkingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterDeicingAreaMarkingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, deicingareamarkingtimeslicepropertytype)
	}
}

func (deicingareamarkingtimeslicepropertytype *DeicingAreaMarkingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaMarkingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontDeicingAreaMarkingTimeSlicePropertyType *DeicingAreaMarkingTimeSlicePropertyType
		if front != nil {
			frontDeicingAreaMarkingTimeSlicePropertyType, _ = front.(*DeicingAreaMarkingTimeSlicePropertyType)
		}
		stage.OnAfterDeicingAreaMarkingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, deicingareamarkingtimeslicepropertytype, frontDeicingAreaMarkingTimeSlicePropertyType)
	}
}

func (deicingareamarkingtimeslicepropertytype *DeicingAreaMarkingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaMarkingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontDeicingAreaMarkingTimeSlicePropertyType *DeicingAreaMarkingTimeSlicePropertyType
		if front != nil {
			frontDeicingAreaMarkingTimeSlicePropertyType, _ = front.(*DeicingAreaMarkingTimeSlicePropertyType)
		}
		stage.OnAfterDeicingAreaMarkingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, deicingareamarkingtimeslicepropertytype, frontDeicingAreaMarkingTimeSlicePropertyType)
	}
}

func (deicingareamarkingtimeslicetype *DeicingAreaMarkingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeicingAreaMarkingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterDeicingAreaMarkingTimeSliceTypeCreateCallback.OnAfterCreate(stage, deicingareamarkingtimeslicetype)
	}
}

func (deicingareamarkingtimeslicetype *DeicingAreaMarkingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaMarkingTimeSliceTypeUpdateCallback != nil {
		var frontDeicingAreaMarkingTimeSliceType *DeicingAreaMarkingTimeSliceType
		if front != nil {
			frontDeicingAreaMarkingTimeSliceType, _ = front.(*DeicingAreaMarkingTimeSliceType)
		}
		stage.OnAfterDeicingAreaMarkingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, deicingareamarkingtimeslicetype, frontDeicingAreaMarkingTimeSliceType)
	}
}

func (deicingareamarkingtimeslicetype *DeicingAreaMarkingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaMarkingTimeSliceTypeDeleteCallback != nil {
		var frontDeicingAreaMarkingTimeSliceType *DeicingAreaMarkingTimeSliceType
		if front != nil {
			frontDeicingAreaMarkingTimeSliceType, _ = front.(*DeicingAreaMarkingTimeSliceType)
		}
		stage.OnAfterDeicingAreaMarkingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, deicingareamarkingtimeslicetype, frontDeicingAreaMarkingTimeSliceType)
	}
}

func (deicingareamarkingtype *DeicingAreaMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeicingAreaMarkingTypeCreateCallback != nil {
		stage.OnAfterDeicingAreaMarkingTypeCreateCallback.OnAfterCreate(stage, deicingareamarkingtype)
	}
}

func (deicingareamarkingtype *DeicingAreaMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaMarkingTypeUpdateCallback != nil {
		var frontDeicingAreaMarkingType *DeicingAreaMarkingType
		if front != nil {
			frontDeicingAreaMarkingType, _ = front.(*DeicingAreaMarkingType)
		}
		stage.OnAfterDeicingAreaMarkingTypeUpdateCallback.OnAfterUpdate(stage, deicingareamarkingtype, frontDeicingAreaMarkingType)
	}
}

func (deicingareamarkingtype *DeicingAreaMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaMarkingTypeDeleteCallback != nil {
		var frontDeicingAreaMarkingType *DeicingAreaMarkingType
		if front != nil {
			frontDeicingAreaMarkingType, _ = front.(*DeicingAreaMarkingType)
		}
		stage.OnAfterDeicingAreaMarkingTypeDeleteCallback.OnAfterDelete(stage, deicingareamarkingtype, frontDeicingAreaMarkingType)
	}
}

func (deicingareapropertytype *DeicingAreaPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeicingAreaPropertyTypeCreateCallback != nil {
		stage.OnAfterDeicingAreaPropertyTypeCreateCallback.OnAfterCreate(stage, deicingareapropertytype)
	}
}

func (deicingareapropertytype *DeicingAreaPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaPropertyTypeUpdateCallback != nil {
		var frontDeicingAreaPropertyType *DeicingAreaPropertyType
		if front != nil {
			frontDeicingAreaPropertyType, _ = front.(*DeicingAreaPropertyType)
		}
		stage.OnAfterDeicingAreaPropertyTypeUpdateCallback.OnAfterUpdate(stage, deicingareapropertytype, frontDeicingAreaPropertyType)
	}
}

func (deicingareapropertytype *DeicingAreaPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaPropertyTypeDeleteCallback != nil {
		var frontDeicingAreaPropertyType *DeicingAreaPropertyType
		if front != nil {
			frontDeicingAreaPropertyType, _ = front.(*DeicingAreaPropertyType)
		}
		stage.OnAfterDeicingAreaPropertyTypeDeleteCallback.OnAfterDelete(stage, deicingareapropertytype, frontDeicingAreaPropertyType)
	}
}

func (deicingareatimeslicepropertytype *DeicingAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeicingAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterDeicingAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, deicingareatimeslicepropertytype)
	}
}

func (deicingareatimeslicepropertytype *DeicingAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontDeicingAreaTimeSlicePropertyType *DeicingAreaTimeSlicePropertyType
		if front != nil {
			frontDeicingAreaTimeSlicePropertyType, _ = front.(*DeicingAreaTimeSlicePropertyType)
		}
		stage.OnAfterDeicingAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, deicingareatimeslicepropertytype, frontDeicingAreaTimeSlicePropertyType)
	}
}

func (deicingareatimeslicepropertytype *DeicingAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontDeicingAreaTimeSlicePropertyType *DeicingAreaTimeSlicePropertyType
		if front != nil {
			frontDeicingAreaTimeSlicePropertyType, _ = front.(*DeicingAreaTimeSlicePropertyType)
		}
		stage.OnAfterDeicingAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, deicingareatimeslicepropertytype, frontDeicingAreaTimeSlicePropertyType)
	}
}

func (deicingareatimeslicetype *DeicingAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeicingAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterDeicingAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, deicingareatimeslicetype)
	}
}

func (deicingareatimeslicetype *DeicingAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaTimeSliceTypeUpdateCallback != nil {
		var frontDeicingAreaTimeSliceType *DeicingAreaTimeSliceType
		if front != nil {
			frontDeicingAreaTimeSliceType, _ = front.(*DeicingAreaTimeSliceType)
		}
		stage.OnAfterDeicingAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, deicingareatimeslicetype, frontDeicingAreaTimeSliceType)
	}
}

func (deicingareatimeslicetype *DeicingAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaTimeSliceTypeDeleteCallback != nil {
		var frontDeicingAreaTimeSliceType *DeicingAreaTimeSliceType
		if front != nil {
			frontDeicingAreaTimeSliceType, _ = front.(*DeicingAreaTimeSliceType)
		}
		stage.OnAfterDeicingAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, deicingareatimeslicetype, frontDeicingAreaTimeSliceType)
	}
}

func (deicingareatype *DeicingAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeicingAreaTypeCreateCallback != nil {
		stage.OnAfterDeicingAreaTypeCreateCallback.OnAfterCreate(stage, deicingareatype)
	}
}

func (deicingareatype *DeicingAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaTypeUpdateCallback != nil {
		var frontDeicingAreaType *DeicingAreaType
		if front != nil {
			frontDeicingAreaType, _ = front.(*DeicingAreaType)
		}
		stage.OnAfterDeicingAreaTypeUpdateCallback.OnAfterUpdate(stage, deicingareatype, frontDeicingAreaType)
	}
}

func (deicingareatype *DeicingAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeicingAreaTypeDeleteCallback != nil {
		var frontDeicingAreaType *DeicingAreaType
		if front != nil {
			frontDeicingAreaType, _ = front.(*DeicingAreaType)
		}
		stage.OnAfterDeicingAreaTypeDeleteCallback.OnAfterDelete(stage, deicingareatype, frontDeicingAreaType)
	}
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedCurvePropertyTypeCreateCallback != nil {
		stage.OnAfterElevatedCurvePropertyTypeCreateCallback.OnAfterCreate(stage, elevatedcurvepropertytype)
	}
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedCurvePropertyTypeUpdateCallback != nil {
		var frontElevatedCurvePropertyType *ElevatedCurvePropertyType
		if front != nil {
			frontElevatedCurvePropertyType, _ = front.(*ElevatedCurvePropertyType)
		}
		stage.OnAfterElevatedCurvePropertyTypeUpdateCallback.OnAfterUpdate(stage, elevatedcurvepropertytype, frontElevatedCurvePropertyType)
	}
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedCurvePropertyTypeDeleteCallback != nil {
		var frontElevatedCurvePropertyType *ElevatedCurvePropertyType
		if front != nil {
			frontElevatedCurvePropertyType, _ = front.(*ElevatedCurvePropertyType)
		}
		stage.OnAfterElevatedCurvePropertyTypeDeleteCallback.OnAfterDelete(stage, elevatedcurvepropertytype, frontElevatedCurvePropertyType)
	}
}

func (elevatedcurvetype *ElevatedCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedCurveTypeCreateCallback != nil {
		stage.OnAfterElevatedCurveTypeCreateCallback.OnAfterCreate(stage, elevatedcurvetype)
	}
}

func (elevatedcurvetype *ElevatedCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedCurveTypeUpdateCallback != nil {
		var frontElevatedCurveType *ElevatedCurveType
		if front != nil {
			frontElevatedCurveType, _ = front.(*ElevatedCurveType)
		}
		stage.OnAfterElevatedCurveTypeUpdateCallback.OnAfterUpdate(stage, elevatedcurvetype, frontElevatedCurveType)
	}
}

func (elevatedcurvetype *ElevatedCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedCurveTypeDeleteCallback != nil {
		var frontElevatedCurveType *ElevatedCurveType
		if front != nil {
			frontElevatedCurveType, _ = front.(*ElevatedCurveType)
		}
		stage.OnAfterElevatedCurveTypeDeleteCallback.OnAfterDelete(stage, elevatedcurvetype, frontElevatedCurveType)
	}
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedPointPropertyTypeCreateCallback != nil {
		stage.OnAfterElevatedPointPropertyTypeCreateCallback.OnAfterCreate(stage, elevatedpointpropertytype)
	}
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedPointPropertyTypeUpdateCallback != nil {
		var frontElevatedPointPropertyType *ElevatedPointPropertyType
		if front != nil {
			frontElevatedPointPropertyType, _ = front.(*ElevatedPointPropertyType)
		}
		stage.OnAfterElevatedPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, elevatedpointpropertytype, frontElevatedPointPropertyType)
	}
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedPointPropertyTypeDeleteCallback != nil {
		var frontElevatedPointPropertyType *ElevatedPointPropertyType
		if front != nil {
			frontElevatedPointPropertyType, _ = front.(*ElevatedPointPropertyType)
		}
		stage.OnAfterElevatedPointPropertyTypeDeleteCallback.OnAfterDelete(stage, elevatedpointpropertytype, frontElevatedPointPropertyType)
	}
}

func (elevatedpointtype *ElevatedPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedPointTypeCreateCallback != nil {
		stage.OnAfterElevatedPointTypeCreateCallback.OnAfterCreate(stage, elevatedpointtype)
	}
}

func (elevatedpointtype *ElevatedPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedPointTypeUpdateCallback != nil {
		var frontElevatedPointType *ElevatedPointType
		if front != nil {
			frontElevatedPointType, _ = front.(*ElevatedPointType)
		}
		stage.OnAfterElevatedPointTypeUpdateCallback.OnAfterUpdate(stage, elevatedpointtype, frontElevatedPointType)
	}
}

func (elevatedpointtype *ElevatedPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedPointTypeDeleteCallback != nil {
		var frontElevatedPointType *ElevatedPointType
		if front != nil {
			frontElevatedPointType, _ = front.(*ElevatedPointType)
		}
		stage.OnAfterElevatedPointTypeDeleteCallback.OnAfterDelete(stage, elevatedpointtype, frontElevatedPointType)
	}
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedSurfacePropertyTypeCreateCallback != nil {
		stage.OnAfterElevatedSurfacePropertyTypeCreateCallback.OnAfterCreate(stage, elevatedsurfacepropertytype)
	}
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedSurfacePropertyTypeUpdateCallback != nil {
		var frontElevatedSurfacePropertyType *ElevatedSurfacePropertyType
		if front != nil {
			frontElevatedSurfacePropertyType, _ = front.(*ElevatedSurfacePropertyType)
		}
		stage.OnAfterElevatedSurfacePropertyTypeUpdateCallback.OnAfterUpdate(stage, elevatedsurfacepropertytype, frontElevatedSurfacePropertyType)
	}
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedSurfacePropertyTypeDeleteCallback != nil {
		var frontElevatedSurfacePropertyType *ElevatedSurfacePropertyType
		if front != nil {
			frontElevatedSurfacePropertyType, _ = front.(*ElevatedSurfacePropertyType)
		}
		stage.OnAfterElevatedSurfacePropertyTypeDeleteCallback.OnAfterDelete(stage, elevatedsurfacepropertytype, frontElevatedSurfacePropertyType)
	}
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedSurfaceTypeCreateCallback != nil {
		stage.OnAfterElevatedSurfaceTypeCreateCallback.OnAfterCreate(stage, elevatedsurfacetype)
	}
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedSurfaceTypeUpdateCallback != nil {
		var frontElevatedSurfaceType *ElevatedSurfaceType
		if front != nil {
			frontElevatedSurfaceType, _ = front.(*ElevatedSurfaceType)
		}
		stage.OnAfterElevatedSurfaceTypeUpdateCallback.OnAfterUpdate(stage, elevatedsurfacetype, frontElevatedSurfaceType)
	}
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedSurfaceTypeDeleteCallback != nil {
		var frontElevatedSurfaceType *ElevatedSurfaceType
		if front != nil {
			frontElevatedSurfaceType, _ = front.(*ElevatedSurfaceType)
		}
		stage.OnAfterElevatedSurfaceTypeDeleteCallback.OnAfterDelete(stage, elevatedsurfacetype, frontElevatedSurfaceType)
	}
}

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightCharacteristicPropertyTypeCreateCallback != nil {
		stage.OnAfterFlightCharacteristicPropertyTypeCreateCallback.OnAfterCreate(stage, flightcharacteristicpropertytype)
	}
}

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightCharacteristicPropertyTypeUpdateCallback != nil {
		var frontFlightCharacteristicPropertyType *FlightCharacteristicPropertyType
		if front != nil {
			frontFlightCharacteristicPropertyType, _ = front.(*FlightCharacteristicPropertyType)
		}
		stage.OnAfterFlightCharacteristicPropertyTypeUpdateCallback.OnAfterUpdate(stage, flightcharacteristicpropertytype, frontFlightCharacteristicPropertyType)
	}
}

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightCharacteristicPropertyTypeDeleteCallback != nil {
		var frontFlightCharacteristicPropertyType *FlightCharacteristicPropertyType
		if front != nil {
			frontFlightCharacteristicPropertyType, _ = front.(*FlightCharacteristicPropertyType)
		}
		stage.OnAfterFlightCharacteristicPropertyTypeDeleteCallback.OnAfterDelete(stage, flightcharacteristicpropertytype, frontFlightCharacteristicPropertyType)
	}
}

func (flightcharacteristictype *FlightCharacteristicType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightCharacteristicTypeCreateCallback != nil {
		stage.OnAfterFlightCharacteristicTypeCreateCallback.OnAfterCreate(stage, flightcharacteristictype)
	}
}

func (flightcharacteristictype *FlightCharacteristicType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightCharacteristicTypeUpdateCallback != nil {
		var frontFlightCharacteristicType *FlightCharacteristicType
		if front != nil {
			frontFlightCharacteristicType, _ = front.(*FlightCharacteristicType)
		}
		stage.OnAfterFlightCharacteristicTypeUpdateCallback.OnAfterUpdate(stage, flightcharacteristictype, frontFlightCharacteristicType)
	}
}

func (flightcharacteristictype *FlightCharacteristicType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightCharacteristicTypeDeleteCallback != nil {
		var frontFlightCharacteristicType *FlightCharacteristicType
		if front != nil {
			frontFlightCharacteristicType, _ = front.(*FlightCharacteristicType)
		}
		stage.OnAfterFlightCharacteristicTypeDeleteCallback.OnAfterDelete(stage, flightcharacteristictype, frontFlightCharacteristicType)
	}
}

func (floatingdocksitepropertytype *FloatingDockSitePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFloatingDockSitePropertyTypeCreateCallback != nil {
		stage.OnAfterFloatingDockSitePropertyTypeCreateCallback.OnAfterCreate(stage, floatingdocksitepropertytype)
	}
}

func (floatingdocksitepropertytype *FloatingDockSitePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFloatingDockSitePropertyTypeUpdateCallback != nil {
		var frontFloatingDockSitePropertyType *FloatingDockSitePropertyType
		if front != nil {
			frontFloatingDockSitePropertyType, _ = front.(*FloatingDockSitePropertyType)
		}
		stage.OnAfterFloatingDockSitePropertyTypeUpdateCallback.OnAfterUpdate(stage, floatingdocksitepropertytype, frontFloatingDockSitePropertyType)
	}
}

func (floatingdocksitepropertytype *FloatingDockSitePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFloatingDockSitePropertyTypeDeleteCallback != nil {
		var frontFloatingDockSitePropertyType *FloatingDockSitePropertyType
		if front != nil {
			frontFloatingDockSitePropertyType, _ = front.(*FloatingDockSitePropertyType)
		}
		stage.OnAfterFloatingDockSitePropertyTypeDeleteCallback.OnAfterDelete(stage, floatingdocksitepropertytype, frontFloatingDockSitePropertyType)
	}
}

func (floatingdocksitetimeslicepropertytype *FloatingDockSiteTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFloatingDockSiteTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterFloatingDockSiteTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, floatingdocksitetimeslicepropertytype)
	}
}

func (floatingdocksitetimeslicepropertytype *FloatingDockSiteTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFloatingDockSiteTimeSlicePropertyTypeUpdateCallback != nil {
		var frontFloatingDockSiteTimeSlicePropertyType *FloatingDockSiteTimeSlicePropertyType
		if front != nil {
			frontFloatingDockSiteTimeSlicePropertyType, _ = front.(*FloatingDockSiteTimeSlicePropertyType)
		}
		stage.OnAfterFloatingDockSiteTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, floatingdocksitetimeslicepropertytype, frontFloatingDockSiteTimeSlicePropertyType)
	}
}

func (floatingdocksitetimeslicepropertytype *FloatingDockSiteTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFloatingDockSiteTimeSlicePropertyTypeDeleteCallback != nil {
		var frontFloatingDockSiteTimeSlicePropertyType *FloatingDockSiteTimeSlicePropertyType
		if front != nil {
			frontFloatingDockSiteTimeSlicePropertyType, _ = front.(*FloatingDockSiteTimeSlicePropertyType)
		}
		stage.OnAfterFloatingDockSiteTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, floatingdocksitetimeslicepropertytype, frontFloatingDockSiteTimeSlicePropertyType)
	}
}

func (floatingdocksitetimeslicetype *FloatingDockSiteTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFloatingDockSiteTimeSliceTypeCreateCallback != nil {
		stage.OnAfterFloatingDockSiteTimeSliceTypeCreateCallback.OnAfterCreate(stage, floatingdocksitetimeslicetype)
	}
}

func (floatingdocksitetimeslicetype *FloatingDockSiteTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFloatingDockSiteTimeSliceTypeUpdateCallback != nil {
		var frontFloatingDockSiteTimeSliceType *FloatingDockSiteTimeSliceType
		if front != nil {
			frontFloatingDockSiteTimeSliceType, _ = front.(*FloatingDockSiteTimeSliceType)
		}
		stage.OnAfterFloatingDockSiteTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, floatingdocksitetimeslicetype, frontFloatingDockSiteTimeSliceType)
	}
}

func (floatingdocksitetimeslicetype *FloatingDockSiteTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFloatingDockSiteTimeSliceTypeDeleteCallback != nil {
		var frontFloatingDockSiteTimeSliceType *FloatingDockSiteTimeSliceType
		if front != nil {
			frontFloatingDockSiteTimeSliceType, _ = front.(*FloatingDockSiteTimeSliceType)
		}
		stage.OnAfterFloatingDockSiteTimeSliceTypeDeleteCallback.OnAfterDelete(stage, floatingdocksitetimeslicetype, frontFloatingDockSiteTimeSliceType)
	}
}

func (floatingdocksitetype *FloatingDockSiteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFloatingDockSiteTypeCreateCallback != nil {
		stage.OnAfterFloatingDockSiteTypeCreateCallback.OnAfterCreate(stage, floatingdocksitetype)
	}
}

func (floatingdocksitetype *FloatingDockSiteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFloatingDockSiteTypeUpdateCallback != nil {
		var frontFloatingDockSiteType *FloatingDockSiteType
		if front != nil {
			frontFloatingDockSiteType, _ = front.(*FloatingDockSiteType)
		}
		stage.OnAfterFloatingDockSiteTypeUpdateCallback.OnAfterUpdate(stage, floatingdocksitetype, frontFloatingDockSiteType)
	}
}

func (floatingdocksitetype *FloatingDockSiteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFloatingDockSiteTypeDeleteCallback != nil {
		var frontFloatingDockSiteType *FloatingDockSiteType
		if front != nil {
			frontFloatingDockSiteType, _ = front.(*FloatingDockSiteType)
		}
		stage.OnAfterFloatingDockSiteTypeDeleteCallback.OnAfterDelete(stage, floatingdocksitetype, frontFloatingDockSiteType)
	}
}

func (groundlightingavailabilitypropertytype *GroundLightingAvailabilityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroundLightingAvailabilityPropertyTypeCreateCallback != nil {
		stage.OnAfterGroundLightingAvailabilityPropertyTypeCreateCallback.OnAfterCreate(stage, groundlightingavailabilitypropertytype)
	}
}

func (groundlightingavailabilitypropertytype *GroundLightingAvailabilityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundLightingAvailabilityPropertyTypeUpdateCallback != nil {
		var frontGroundLightingAvailabilityPropertyType *GroundLightingAvailabilityPropertyType
		if front != nil {
			frontGroundLightingAvailabilityPropertyType, _ = front.(*GroundLightingAvailabilityPropertyType)
		}
		stage.OnAfterGroundLightingAvailabilityPropertyTypeUpdateCallback.OnAfterUpdate(stage, groundlightingavailabilitypropertytype, frontGroundLightingAvailabilityPropertyType)
	}
}

func (groundlightingavailabilitypropertytype *GroundLightingAvailabilityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundLightingAvailabilityPropertyTypeDeleteCallback != nil {
		var frontGroundLightingAvailabilityPropertyType *GroundLightingAvailabilityPropertyType
		if front != nil {
			frontGroundLightingAvailabilityPropertyType, _ = front.(*GroundLightingAvailabilityPropertyType)
		}
		stage.OnAfterGroundLightingAvailabilityPropertyTypeDeleteCallback.OnAfterDelete(stage, groundlightingavailabilitypropertytype, frontGroundLightingAvailabilityPropertyType)
	}
}

func (groundlightingavailabilitytype *GroundLightingAvailabilityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroundLightingAvailabilityTypeCreateCallback != nil {
		stage.OnAfterGroundLightingAvailabilityTypeCreateCallback.OnAfterCreate(stage, groundlightingavailabilitytype)
	}
}

func (groundlightingavailabilitytype *GroundLightingAvailabilityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundLightingAvailabilityTypeUpdateCallback != nil {
		var frontGroundLightingAvailabilityType *GroundLightingAvailabilityType
		if front != nil {
			frontGroundLightingAvailabilityType, _ = front.(*GroundLightingAvailabilityType)
		}
		stage.OnAfterGroundLightingAvailabilityTypeUpdateCallback.OnAfterUpdate(stage, groundlightingavailabilitytype, frontGroundLightingAvailabilityType)
	}
}

func (groundlightingavailabilitytype *GroundLightingAvailabilityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundLightingAvailabilityTypeDeleteCallback != nil {
		var frontGroundLightingAvailabilityType *GroundLightingAvailabilityType
		if front != nil {
			frontGroundLightingAvailabilityType, _ = front.(*GroundLightingAvailabilityType)
		}
		stage.OnAfterGroundLightingAvailabilityTypeDeleteCallback.OnAfterDelete(stage, groundlightingavailabilitytype, frontGroundLightingAvailabilityType)
	}
}

func (guidancelinelightsystemtimeslicepropertytype *GuidanceLineLightSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLineLightSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterGuidanceLineLightSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, guidancelinelightsystemtimeslicepropertytype)
	}
}

func (guidancelinelightsystemtimeslicepropertytype *GuidanceLineLightSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineLightSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontGuidanceLineLightSystemTimeSlicePropertyType *GuidanceLineLightSystemTimeSlicePropertyType
		if front != nil {
			frontGuidanceLineLightSystemTimeSlicePropertyType, _ = front.(*GuidanceLineLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterGuidanceLineLightSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, guidancelinelightsystemtimeslicepropertytype, frontGuidanceLineLightSystemTimeSlicePropertyType)
	}
}

func (guidancelinelightsystemtimeslicepropertytype *GuidanceLineLightSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineLightSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontGuidanceLineLightSystemTimeSlicePropertyType *GuidanceLineLightSystemTimeSlicePropertyType
		if front != nil {
			frontGuidanceLineLightSystemTimeSlicePropertyType, _ = front.(*GuidanceLineLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterGuidanceLineLightSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, guidancelinelightsystemtimeslicepropertytype, frontGuidanceLineLightSystemTimeSlicePropertyType)
	}
}

func (guidancelinelightsystemtimeslicetype *GuidanceLineLightSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLineLightSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterGuidanceLineLightSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, guidancelinelightsystemtimeslicetype)
	}
}

func (guidancelinelightsystemtimeslicetype *GuidanceLineLightSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineLightSystemTimeSliceTypeUpdateCallback != nil {
		var frontGuidanceLineLightSystemTimeSliceType *GuidanceLineLightSystemTimeSliceType
		if front != nil {
			frontGuidanceLineLightSystemTimeSliceType, _ = front.(*GuidanceLineLightSystemTimeSliceType)
		}
		stage.OnAfterGuidanceLineLightSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, guidancelinelightsystemtimeslicetype, frontGuidanceLineLightSystemTimeSliceType)
	}
}

func (guidancelinelightsystemtimeslicetype *GuidanceLineLightSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineLightSystemTimeSliceTypeDeleteCallback != nil {
		var frontGuidanceLineLightSystemTimeSliceType *GuidanceLineLightSystemTimeSliceType
		if front != nil {
			frontGuidanceLineLightSystemTimeSliceType, _ = front.(*GuidanceLineLightSystemTimeSliceType)
		}
		stage.OnAfterGuidanceLineLightSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, guidancelinelightsystemtimeslicetype, frontGuidanceLineLightSystemTimeSliceType)
	}
}

func (guidancelinelightsystemtype *GuidanceLineLightSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLineLightSystemTypeCreateCallback != nil {
		stage.OnAfterGuidanceLineLightSystemTypeCreateCallback.OnAfterCreate(stage, guidancelinelightsystemtype)
	}
}

func (guidancelinelightsystemtype *GuidanceLineLightSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineLightSystemTypeUpdateCallback != nil {
		var frontGuidanceLineLightSystemType *GuidanceLineLightSystemType
		if front != nil {
			frontGuidanceLineLightSystemType, _ = front.(*GuidanceLineLightSystemType)
		}
		stage.OnAfterGuidanceLineLightSystemTypeUpdateCallback.OnAfterUpdate(stage, guidancelinelightsystemtype, frontGuidanceLineLightSystemType)
	}
}

func (guidancelinelightsystemtype *GuidanceLineLightSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineLightSystemTypeDeleteCallback != nil {
		var frontGuidanceLineLightSystemType *GuidanceLineLightSystemType
		if front != nil {
			frontGuidanceLineLightSystemType, _ = front.(*GuidanceLineLightSystemType)
		}
		stage.OnAfterGuidanceLineLightSystemTypeDeleteCallback.OnAfterDelete(stage, guidancelinelightsystemtype, frontGuidanceLineLightSystemType)
	}
}

func (guidancelinemarkingtimeslicepropertytype *GuidanceLineMarkingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLineMarkingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterGuidanceLineMarkingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, guidancelinemarkingtimeslicepropertytype)
	}
}

func (guidancelinemarkingtimeslicepropertytype *GuidanceLineMarkingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineMarkingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontGuidanceLineMarkingTimeSlicePropertyType *GuidanceLineMarkingTimeSlicePropertyType
		if front != nil {
			frontGuidanceLineMarkingTimeSlicePropertyType, _ = front.(*GuidanceLineMarkingTimeSlicePropertyType)
		}
		stage.OnAfterGuidanceLineMarkingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, guidancelinemarkingtimeslicepropertytype, frontGuidanceLineMarkingTimeSlicePropertyType)
	}
}

func (guidancelinemarkingtimeslicepropertytype *GuidanceLineMarkingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineMarkingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontGuidanceLineMarkingTimeSlicePropertyType *GuidanceLineMarkingTimeSlicePropertyType
		if front != nil {
			frontGuidanceLineMarkingTimeSlicePropertyType, _ = front.(*GuidanceLineMarkingTimeSlicePropertyType)
		}
		stage.OnAfterGuidanceLineMarkingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, guidancelinemarkingtimeslicepropertytype, frontGuidanceLineMarkingTimeSlicePropertyType)
	}
}

func (guidancelinemarkingtimeslicetype *GuidanceLineMarkingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLineMarkingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterGuidanceLineMarkingTimeSliceTypeCreateCallback.OnAfterCreate(stage, guidancelinemarkingtimeslicetype)
	}
}

func (guidancelinemarkingtimeslicetype *GuidanceLineMarkingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineMarkingTimeSliceTypeUpdateCallback != nil {
		var frontGuidanceLineMarkingTimeSliceType *GuidanceLineMarkingTimeSliceType
		if front != nil {
			frontGuidanceLineMarkingTimeSliceType, _ = front.(*GuidanceLineMarkingTimeSliceType)
		}
		stage.OnAfterGuidanceLineMarkingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, guidancelinemarkingtimeslicetype, frontGuidanceLineMarkingTimeSliceType)
	}
}

func (guidancelinemarkingtimeslicetype *GuidanceLineMarkingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineMarkingTimeSliceTypeDeleteCallback != nil {
		var frontGuidanceLineMarkingTimeSliceType *GuidanceLineMarkingTimeSliceType
		if front != nil {
			frontGuidanceLineMarkingTimeSliceType, _ = front.(*GuidanceLineMarkingTimeSliceType)
		}
		stage.OnAfterGuidanceLineMarkingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, guidancelinemarkingtimeslicetype, frontGuidanceLineMarkingTimeSliceType)
	}
}

func (guidancelinemarkingtype *GuidanceLineMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLineMarkingTypeCreateCallback != nil {
		stage.OnAfterGuidanceLineMarkingTypeCreateCallback.OnAfterCreate(stage, guidancelinemarkingtype)
	}
}

func (guidancelinemarkingtype *GuidanceLineMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineMarkingTypeUpdateCallback != nil {
		var frontGuidanceLineMarkingType *GuidanceLineMarkingType
		if front != nil {
			frontGuidanceLineMarkingType, _ = front.(*GuidanceLineMarkingType)
		}
		stage.OnAfterGuidanceLineMarkingTypeUpdateCallback.OnAfterUpdate(stage, guidancelinemarkingtype, frontGuidanceLineMarkingType)
	}
}

func (guidancelinemarkingtype *GuidanceLineMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineMarkingTypeDeleteCallback != nil {
		var frontGuidanceLineMarkingType *GuidanceLineMarkingType
		if front != nil {
			frontGuidanceLineMarkingType, _ = front.(*GuidanceLineMarkingType)
		}
		stage.OnAfterGuidanceLineMarkingTypeDeleteCallback.OnAfterDelete(stage, guidancelinemarkingtype, frontGuidanceLineMarkingType)
	}
}

func (guidancelinepropertytype *GuidanceLinePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLinePropertyTypeCreateCallback != nil {
		stage.OnAfterGuidanceLinePropertyTypeCreateCallback.OnAfterCreate(stage, guidancelinepropertytype)
	}
}

func (guidancelinepropertytype *GuidanceLinePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLinePropertyTypeUpdateCallback != nil {
		var frontGuidanceLinePropertyType *GuidanceLinePropertyType
		if front != nil {
			frontGuidanceLinePropertyType, _ = front.(*GuidanceLinePropertyType)
		}
		stage.OnAfterGuidanceLinePropertyTypeUpdateCallback.OnAfterUpdate(stage, guidancelinepropertytype, frontGuidanceLinePropertyType)
	}
}

func (guidancelinepropertytype *GuidanceLinePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLinePropertyTypeDeleteCallback != nil {
		var frontGuidanceLinePropertyType *GuidanceLinePropertyType
		if front != nil {
			frontGuidanceLinePropertyType, _ = front.(*GuidanceLinePropertyType)
		}
		stage.OnAfterGuidanceLinePropertyTypeDeleteCallback.OnAfterDelete(stage, guidancelinepropertytype, frontGuidanceLinePropertyType)
	}
}

func (guidancelinetimeslicepropertytype *GuidanceLineTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLineTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterGuidanceLineTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, guidancelinetimeslicepropertytype)
	}
}

func (guidancelinetimeslicepropertytype *GuidanceLineTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineTimeSlicePropertyTypeUpdateCallback != nil {
		var frontGuidanceLineTimeSlicePropertyType *GuidanceLineTimeSlicePropertyType
		if front != nil {
			frontGuidanceLineTimeSlicePropertyType, _ = front.(*GuidanceLineTimeSlicePropertyType)
		}
		stage.OnAfterGuidanceLineTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, guidancelinetimeslicepropertytype, frontGuidanceLineTimeSlicePropertyType)
	}
}

func (guidancelinetimeslicepropertytype *GuidanceLineTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineTimeSlicePropertyTypeDeleteCallback != nil {
		var frontGuidanceLineTimeSlicePropertyType *GuidanceLineTimeSlicePropertyType
		if front != nil {
			frontGuidanceLineTimeSlicePropertyType, _ = front.(*GuidanceLineTimeSlicePropertyType)
		}
		stage.OnAfterGuidanceLineTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, guidancelinetimeslicepropertytype, frontGuidanceLineTimeSlicePropertyType)
	}
}

func (guidancelinetimeslicetype *GuidanceLineTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLineTimeSliceTypeCreateCallback != nil {
		stage.OnAfterGuidanceLineTimeSliceTypeCreateCallback.OnAfterCreate(stage, guidancelinetimeslicetype)
	}
}

func (guidancelinetimeslicetype *GuidanceLineTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineTimeSliceTypeUpdateCallback != nil {
		var frontGuidanceLineTimeSliceType *GuidanceLineTimeSliceType
		if front != nil {
			frontGuidanceLineTimeSliceType, _ = front.(*GuidanceLineTimeSliceType)
		}
		stage.OnAfterGuidanceLineTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, guidancelinetimeslicetype, frontGuidanceLineTimeSliceType)
	}
}

func (guidancelinetimeslicetype *GuidanceLineTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineTimeSliceTypeDeleteCallback != nil {
		var frontGuidanceLineTimeSliceType *GuidanceLineTimeSliceType
		if front != nil {
			frontGuidanceLineTimeSliceType, _ = front.(*GuidanceLineTimeSliceType)
		}
		stage.OnAfterGuidanceLineTimeSliceTypeDeleteCallback.OnAfterDelete(stage, guidancelinetimeslicetype, frontGuidanceLineTimeSliceType)
	}
}

func (guidancelinetype *GuidanceLineType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGuidanceLineTypeCreateCallback != nil {
		stage.OnAfterGuidanceLineTypeCreateCallback.OnAfterCreate(stage, guidancelinetype)
	}
}

func (guidancelinetype *GuidanceLineType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineTypeUpdateCallback != nil {
		var frontGuidanceLineType *GuidanceLineType
		if front != nil {
			frontGuidanceLineType, _ = front.(*GuidanceLineType)
		}
		stage.OnAfterGuidanceLineTypeUpdateCallback.OnAfterUpdate(stage, guidancelinetype, frontGuidanceLineType)
	}
}

func (guidancelinetype *GuidanceLineType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGuidanceLineTypeDeleteCallback != nil {
		var frontGuidanceLineType *GuidanceLineType
		if front != nil {
			frontGuidanceLineType, _ = front.(*GuidanceLineType)
		}
		stage.OnAfterGuidanceLineTypeDeleteCallback.OnAfterDelete(stage, guidancelinetype, frontGuidanceLineType)
	}
}

func (lightelementpropertytype *LightElementPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightElementPropertyTypeCreateCallback != nil {
		stage.OnAfterLightElementPropertyTypeCreateCallback.OnAfterCreate(stage, lightelementpropertytype)
	}
}

func (lightelementpropertytype *LightElementPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementPropertyTypeUpdateCallback != nil {
		var frontLightElementPropertyType *LightElementPropertyType
		if front != nil {
			frontLightElementPropertyType, _ = front.(*LightElementPropertyType)
		}
		stage.OnAfterLightElementPropertyTypeUpdateCallback.OnAfterUpdate(stage, lightelementpropertytype, frontLightElementPropertyType)
	}
}

func (lightelementpropertytype *LightElementPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementPropertyTypeDeleteCallback != nil {
		var frontLightElementPropertyType *LightElementPropertyType
		if front != nil {
			frontLightElementPropertyType, _ = front.(*LightElementPropertyType)
		}
		stage.OnAfterLightElementPropertyTypeDeleteCallback.OnAfterDelete(stage, lightelementpropertytype, frontLightElementPropertyType)
	}
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightElementStatusPropertyTypeCreateCallback != nil {
		stage.OnAfterLightElementStatusPropertyTypeCreateCallback.OnAfterCreate(stage, lightelementstatuspropertytype)
	}
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementStatusPropertyTypeUpdateCallback != nil {
		var frontLightElementStatusPropertyType *LightElementStatusPropertyType
		if front != nil {
			frontLightElementStatusPropertyType, _ = front.(*LightElementStatusPropertyType)
		}
		stage.OnAfterLightElementStatusPropertyTypeUpdateCallback.OnAfterUpdate(stage, lightelementstatuspropertytype, frontLightElementStatusPropertyType)
	}
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementStatusPropertyTypeDeleteCallback != nil {
		var frontLightElementStatusPropertyType *LightElementStatusPropertyType
		if front != nil {
			frontLightElementStatusPropertyType, _ = front.(*LightElementStatusPropertyType)
		}
		stage.OnAfterLightElementStatusPropertyTypeDeleteCallback.OnAfterDelete(stage, lightelementstatuspropertytype, frontLightElementStatusPropertyType)
	}
}

func (lightelementstatustype *LightElementStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightElementStatusTypeCreateCallback != nil {
		stage.OnAfterLightElementStatusTypeCreateCallback.OnAfterCreate(stage, lightelementstatustype)
	}
}

func (lightelementstatustype *LightElementStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementStatusTypeUpdateCallback != nil {
		var frontLightElementStatusType *LightElementStatusType
		if front != nil {
			frontLightElementStatusType, _ = front.(*LightElementStatusType)
		}
		stage.OnAfterLightElementStatusTypeUpdateCallback.OnAfterUpdate(stage, lightelementstatustype, frontLightElementStatusType)
	}
}

func (lightelementstatustype *LightElementStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementStatusTypeDeleteCallback != nil {
		var frontLightElementStatusType *LightElementStatusType
		if front != nil {
			frontLightElementStatusType, _ = front.(*LightElementStatusType)
		}
		stage.OnAfterLightElementStatusTypeDeleteCallback.OnAfterDelete(stage, lightelementstatustype, frontLightElementStatusType)
	}
}

func (lightelementtype *LightElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightElementTypeCreateCallback != nil {
		stage.OnAfterLightElementTypeCreateCallback.OnAfterCreate(stage, lightelementtype)
	}
}

func (lightelementtype *LightElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementTypeUpdateCallback != nil {
		var frontLightElementType *LightElementType
		if front != nil {
			frontLightElementType, _ = front.(*LightElementType)
		}
		stage.OnAfterLightElementTypeUpdateCallback.OnAfterUpdate(stage, lightelementtype, frontLightElementType)
	}
}

func (lightelementtype *LightElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementTypeDeleteCallback != nil {
		var frontLightElementType *LightElementType
		if front != nil {
			frontLightElementType, _ = front.(*LightElementType)
		}
		stage.OnAfterLightElementTypeDeleteCallback.OnAfterDelete(stage, lightelementtype, frontLightElementType)
	}
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinguisticNotePropertyTypeCreateCallback != nil {
		stage.OnAfterLinguisticNotePropertyTypeCreateCallback.OnAfterCreate(stage, linguisticnotepropertytype)
	}
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinguisticNotePropertyTypeUpdateCallback != nil {
		var frontLinguisticNotePropertyType *LinguisticNotePropertyType
		if front != nil {
			frontLinguisticNotePropertyType, _ = front.(*LinguisticNotePropertyType)
		}
		stage.OnAfterLinguisticNotePropertyTypeUpdateCallback.OnAfterUpdate(stage, linguisticnotepropertytype, frontLinguisticNotePropertyType)
	}
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinguisticNotePropertyTypeDeleteCallback != nil {
		var frontLinguisticNotePropertyType *LinguisticNotePropertyType
		if front != nil {
			frontLinguisticNotePropertyType, _ = front.(*LinguisticNotePropertyType)
		}
		stage.OnAfterLinguisticNotePropertyTypeDeleteCallback.OnAfterDelete(stage, linguisticnotepropertytype, frontLinguisticNotePropertyType)
	}
}

func (linguisticnotetype *LinguisticNoteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinguisticNoteTypeCreateCallback != nil {
		stage.OnAfterLinguisticNoteTypeCreateCallback.OnAfterCreate(stage, linguisticnotetype)
	}
}

func (linguisticnotetype *LinguisticNoteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinguisticNoteTypeUpdateCallback != nil {
		var frontLinguisticNoteType *LinguisticNoteType
		if front != nil {
			frontLinguisticNoteType, _ = front.(*LinguisticNoteType)
		}
		stage.OnAfterLinguisticNoteTypeUpdateCallback.OnAfterUpdate(stage, linguisticnotetype, frontLinguisticNoteType)
	}
}

func (linguisticnotetype *LinguisticNoteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinguisticNoteTypeDeleteCallback != nil {
		var frontLinguisticNoteType *LinguisticNoteType
		if front != nil {
			frontLinguisticNoteType, _ = front.(*LinguisticNoteType)
		}
		stage.OnAfterLinguisticNoteTypeDeleteCallback.OnAfterDelete(stage, linguisticnotetype, frontLinguisticNoteType)
	}
}

func (manoeuvringareaavailabilitypropertytype *ManoeuvringAreaAvailabilityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterManoeuvringAreaAvailabilityPropertyTypeCreateCallback != nil {
		stage.OnAfterManoeuvringAreaAvailabilityPropertyTypeCreateCallback.OnAfterCreate(stage, manoeuvringareaavailabilitypropertytype)
	}
}

func (manoeuvringareaavailabilitypropertytype *ManoeuvringAreaAvailabilityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterManoeuvringAreaAvailabilityPropertyTypeUpdateCallback != nil {
		var frontManoeuvringAreaAvailabilityPropertyType *ManoeuvringAreaAvailabilityPropertyType
		if front != nil {
			frontManoeuvringAreaAvailabilityPropertyType, _ = front.(*ManoeuvringAreaAvailabilityPropertyType)
		}
		stage.OnAfterManoeuvringAreaAvailabilityPropertyTypeUpdateCallback.OnAfterUpdate(stage, manoeuvringareaavailabilitypropertytype, frontManoeuvringAreaAvailabilityPropertyType)
	}
}

func (manoeuvringareaavailabilitypropertytype *ManoeuvringAreaAvailabilityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterManoeuvringAreaAvailabilityPropertyTypeDeleteCallback != nil {
		var frontManoeuvringAreaAvailabilityPropertyType *ManoeuvringAreaAvailabilityPropertyType
		if front != nil {
			frontManoeuvringAreaAvailabilityPropertyType, _ = front.(*ManoeuvringAreaAvailabilityPropertyType)
		}
		stage.OnAfterManoeuvringAreaAvailabilityPropertyTypeDeleteCallback.OnAfterDelete(stage, manoeuvringareaavailabilitypropertytype, frontManoeuvringAreaAvailabilityPropertyType)
	}
}

func (manoeuvringareaavailabilitytype *ManoeuvringAreaAvailabilityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterManoeuvringAreaAvailabilityTypeCreateCallback != nil {
		stage.OnAfterManoeuvringAreaAvailabilityTypeCreateCallback.OnAfterCreate(stage, manoeuvringareaavailabilitytype)
	}
}

func (manoeuvringareaavailabilitytype *ManoeuvringAreaAvailabilityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterManoeuvringAreaAvailabilityTypeUpdateCallback != nil {
		var frontManoeuvringAreaAvailabilityType *ManoeuvringAreaAvailabilityType
		if front != nil {
			frontManoeuvringAreaAvailabilityType, _ = front.(*ManoeuvringAreaAvailabilityType)
		}
		stage.OnAfterManoeuvringAreaAvailabilityTypeUpdateCallback.OnAfterUpdate(stage, manoeuvringareaavailabilitytype, frontManoeuvringAreaAvailabilityType)
	}
}

func (manoeuvringareaavailabilitytype *ManoeuvringAreaAvailabilityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterManoeuvringAreaAvailabilityTypeDeleteCallback != nil {
		var frontManoeuvringAreaAvailabilityType *ManoeuvringAreaAvailabilityType
		if front != nil {
			frontManoeuvringAreaAvailabilityType, _ = front.(*ManoeuvringAreaAvailabilityType)
		}
		stage.OnAfterManoeuvringAreaAvailabilityTypeDeleteCallback.OnAfterDelete(stage, manoeuvringareaavailabilitytype, frontManoeuvringAreaAvailabilityType)
	}
}

func (manoeuvringareausagepropertytype *ManoeuvringAreaUsagePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterManoeuvringAreaUsagePropertyTypeCreateCallback != nil {
		stage.OnAfterManoeuvringAreaUsagePropertyTypeCreateCallback.OnAfterCreate(stage, manoeuvringareausagepropertytype)
	}
}

func (manoeuvringareausagepropertytype *ManoeuvringAreaUsagePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterManoeuvringAreaUsagePropertyTypeUpdateCallback != nil {
		var frontManoeuvringAreaUsagePropertyType *ManoeuvringAreaUsagePropertyType
		if front != nil {
			frontManoeuvringAreaUsagePropertyType, _ = front.(*ManoeuvringAreaUsagePropertyType)
		}
		stage.OnAfterManoeuvringAreaUsagePropertyTypeUpdateCallback.OnAfterUpdate(stage, manoeuvringareausagepropertytype, frontManoeuvringAreaUsagePropertyType)
	}
}

func (manoeuvringareausagepropertytype *ManoeuvringAreaUsagePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterManoeuvringAreaUsagePropertyTypeDeleteCallback != nil {
		var frontManoeuvringAreaUsagePropertyType *ManoeuvringAreaUsagePropertyType
		if front != nil {
			frontManoeuvringAreaUsagePropertyType, _ = front.(*ManoeuvringAreaUsagePropertyType)
		}
		stage.OnAfterManoeuvringAreaUsagePropertyTypeDeleteCallback.OnAfterDelete(stage, manoeuvringareausagepropertytype, frontManoeuvringAreaUsagePropertyType)
	}
}

func (manoeuvringareausagetype *ManoeuvringAreaUsageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterManoeuvringAreaUsageTypeCreateCallback != nil {
		stage.OnAfterManoeuvringAreaUsageTypeCreateCallback.OnAfterCreate(stage, manoeuvringareausagetype)
	}
}

func (manoeuvringareausagetype *ManoeuvringAreaUsageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterManoeuvringAreaUsageTypeUpdateCallback != nil {
		var frontManoeuvringAreaUsageType *ManoeuvringAreaUsageType
		if front != nil {
			frontManoeuvringAreaUsageType, _ = front.(*ManoeuvringAreaUsageType)
		}
		stage.OnAfterManoeuvringAreaUsageTypeUpdateCallback.OnAfterUpdate(stage, manoeuvringareausagetype, frontManoeuvringAreaUsageType)
	}
}

func (manoeuvringareausagetype *ManoeuvringAreaUsageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterManoeuvringAreaUsageTypeDeleteCallback != nil {
		var frontManoeuvringAreaUsageType *ManoeuvringAreaUsageType
		if front != nil {
			frontManoeuvringAreaUsageType, _ = front.(*ManoeuvringAreaUsageType)
		}
		stage.OnAfterManoeuvringAreaUsageTypeDeleteCallback.OnAfterDelete(stage, manoeuvringareausagetype, frontManoeuvringAreaUsageType)
	}
}

func (markingbuoytimeslicepropertytype *MarkingBuoyTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkingBuoyTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterMarkingBuoyTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, markingbuoytimeslicepropertytype)
	}
}

func (markingbuoytimeslicepropertytype *MarkingBuoyTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingBuoyTimeSlicePropertyTypeUpdateCallback != nil {
		var frontMarkingBuoyTimeSlicePropertyType *MarkingBuoyTimeSlicePropertyType
		if front != nil {
			frontMarkingBuoyTimeSlicePropertyType, _ = front.(*MarkingBuoyTimeSlicePropertyType)
		}
		stage.OnAfterMarkingBuoyTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, markingbuoytimeslicepropertytype, frontMarkingBuoyTimeSlicePropertyType)
	}
}

func (markingbuoytimeslicepropertytype *MarkingBuoyTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingBuoyTimeSlicePropertyTypeDeleteCallback != nil {
		var frontMarkingBuoyTimeSlicePropertyType *MarkingBuoyTimeSlicePropertyType
		if front != nil {
			frontMarkingBuoyTimeSlicePropertyType, _ = front.(*MarkingBuoyTimeSlicePropertyType)
		}
		stage.OnAfterMarkingBuoyTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, markingbuoytimeslicepropertytype, frontMarkingBuoyTimeSlicePropertyType)
	}
}

func (markingbuoytimeslicetype *MarkingBuoyTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkingBuoyTimeSliceTypeCreateCallback != nil {
		stage.OnAfterMarkingBuoyTimeSliceTypeCreateCallback.OnAfterCreate(stage, markingbuoytimeslicetype)
	}
}

func (markingbuoytimeslicetype *MarkingBuoyTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingBuoyTimeSliceTypeUpdateCallback != nil {
		var frontMarkingBuoyTimeSliceType *MarkingBuoyTimeSliceType
		if front != nil {
			frontMarkingBuoyTimeSliceType, _ = front.(*MarkingBuoyTimeSliceType)
		}
		stage.OnAfterMarkingBuoyTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, markingbuoytimeslicetype, frontMarkingBuoyTimeSliceType)
	}
}

func (markingbuoytimeslicetype *MarkingBuoyTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingBuoyTimeSliceTypeDeleteCallback != nil {
		var frontMarkingBuoyTimeSliceType *MarkingBuoyTimeSliceType
		if front != nil {
			frontMarkingBuoyTimeSliceType, _ = front.(*MarkingBuoyTimeSliceType)
		}
		stage.OnAfterMarkingBuoyTimeSliceTypeDeleteCallback.OnAfterDelete(stage, markingbuoytimeslicetype, frontMarkingBuoyTimeSliceType)
	}
}

func (markingbuoytype *MarkingBuoyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkingBuoyTypeCreateCallback != nil {
		stage.OnAfterMarkingBuoyTypeCreateCallback.OnAfterCreate(stage, markingbuoytype)
	}
}

func (markingbuoytype *MarkingBuoyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingBuoyTypeUpdateCallback != nil {
		var frontMarkingBuoyType *MarkingBuoyType
		if front != nil {
			frontMarkingBuoyType, _ = front.(*MarkingBuoyType)
		}
		stage.OnAfterMarkingBuoyTypeUpdateCallback.OnAfterUpdate(stage, markingbuoytype, frontMarkingBuoyType)
	}
}

func (markingbuoytype *MarkingBuoyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingBuoyTypeDeleteCallback != nil {
		var frontMarkingBuoyType *MarkingBuoyType
		if front != nil {
			frontMarkingBuoyType, _ = front.(*MarkingBuoyType)
		}
		stage.OnAfterMarkingBuoyTypeDeleteCallback.OnAfterDelete(stage, markingbuoytype, frontMarkingBuoyType)
	}
}

func (markingelementpropertytype *MarkingElementPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkingElementPropertyTypeCreateCallback != nil {
		stage.OnAfterMarkingElementPropertyTypeCreateCallback.OnAfterCreate(stage, markingelementpropertytype)
	}
}

func (markingelementpropertytype *MarkingElementPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingElementPropertyTypeUpdateCallback != nil {
		var frontMarkingElementPropertyType *MarkingElementPropertyType
		if front != nil {
			frontMarkingElementPropertyType, _ = front.(*MarkingElementPropertyType)
		}
		stage.OnAfterMarkingElementPropertyTypeUpdateCallback.OnAfterUpdate(stage, markingelementpropertytype, frontMarkingElementPropertyType)
	}
}

func (markingelementpropertytype *MarkingElementPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingElementPropertyTypeDeleteCallback != nil {
		var frontMarkingElementPropertyType *MarkingElementPropertyType
		if front != nil {
			frontMarkingElementPropertyType, _ = front.(*MarkingElementPropertyType)
		}
		stage.OnAfterMarkingElementPropertyTypeDeleteCallback.OnAfterDelete(stage, markingelementpropertytype, frontMarkingElementPropertyType)
	}
}

func (markingelementtype *MarkingElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkingElementTypeCreateCallback != nil {
		stage.OnAfterMarkingElementTypeCreateCallback.OnAfterCreate(stage, markingelementtype)
	}
}

func (markingelementtype *MarkingElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingElementTypeUpdateCallback != nil {
		var frontMarkingElementType *MarkingElementType
		if front != nil {
			frontMarkingElementType, _ = front.(*MarkingElementType)
		}
		stage.OnAfterMarkingElementTypeUpdateCallback.OnAfterUpdate(stage, markingelementtype, frontMarkingElementType)
	}
}

func (markingelementtype *MarkingElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkingElementTypeDeleteCallback != nil {
		var frontMarkingElementType *MarkingElementType
		if front != nil {
			frontMarkingElementType, _ = front.(*MarkingElementType)
		}
		stage.OnAfterMarkingElementTypeDeleteCallback.OnAfterDelete(stage, markingelementtype, frontMarkingElementType)
	}
}

func (meteorologypropertytype *MeteorologyPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeteorologyPropertyTypeCreateCallback != nil {
		stage.OnAfterMeteorologyPropertyTypeCreateCallback.OnAfterCreate(stage, meteorologypropertytype)
	}
}

func (meteorologypropertytype *MeteorologyPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeteorologyPropertyTypeUpdateCallback != nil {
		var frontMeteorologyPropertyType *MeteorologyPropertyType
		if front != nil {
			frontMeteorologyPropertyType, _ = front.(*MeteorologyPropertyType)
		}
		stage.OnAfterMeteorologyPropertyTypeUpdateCallback.OnAfterUpdate(stage, meteorologypropertytype, frontMeteorologyPropertyType)
	}
}

func (meteorologypropertytype *MeteorologyPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeteorologyPropertyTypeDeleteCallback != nil {
		var frontMeteorologyPropertyType *MeteorologyPropertyType
		if front != nil {
			frontMeteorologyPropertyType, _ = front.(*MeteorologyPropertyType)
		}
		stage.OnAfterMeteorologyPropertyTypeDeleteCallback.OnAfterDelete(stage, meteorologypropertytype, frontMeteorologyPropertyType)
	}
}

func (meteorologytype *MeteorologyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeteorologyTypeCreateCallback != nil {
		stage.OnAfterMeteorologyTypeCreateCallback.OnAfterCreate(stage, meteorologytype)
	}
}

func (meteorologytype *MeteorologyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeteorologyTypeUpdateCallback != nil {
		var frontMeteorologyType *MeteorologyType
		if front != nil {
			frontMeteorologyType, _ = front.(*MeteorologyType)
		}
		stage.OnAfterMeteorologyTypeUpdateCallback.OnAfterUpdate(stage, meteorologytype, frontMeteorologyType)
	}
}

func (meteorologytype *MeteorologyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeteorologyTypeDeleteCallback != nil {
		var frontMeteorologyType *MeteorologyType
		if front != nil {
			frontMeteorologyType, _ = front.(*MeteorologyType)
		}
		stage.OnAfterMeteorologyTypeDeleteCallback.OnAfterDelete(stage, meteorologytype, frontMeteorologyType)
	}
}

func (navaidequipmentdistancepropertytype *NavaidEquipmentDistancePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidEquipmentDistancePropertyTypeCreateCallback != nil {
		stage.OnAfterNavaidEquipmentDistancePropertyTypeCreateCallback.OnAfterCreate(stage, navaidequipmentdistancepropertytype)
	}
}

func (navaidequipmentdistancepropertytype *NavaidEquipmentDistancePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentDistancePropertyTypeUpdateCallback != nil {
		var frontNavaidEquipmentDistancePropertyType *NavaidEquipmentDistancePropertyType
		if front != nil {
			frontNavaidEquipmentDistancePropertyType, _ = front.(*NavaidEquipmentDistancePropertyType)
		}
		stage.OnAfterNavaidEquipmentDistancePropertyTypeUpdateCallback.OnAfterUpdate(stage, navaidequipmentdistancepropertytype, frontNavaidEquipmentDistancePropertyType)
	}
}

func (navaidequipmentdistancepropertytype *NavaidEquipmentDistancePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentDistancePropertyTypeDeleteCallback != nil {
		var frontNavaidEquipmentDistancePropertyType *NavaidEquipmentDistancePropertyType
		if front != nil {
			frontNavaidEquipmentDistancePropertyType, _ = front.(*NavaidEquipmentDistancePropertyType)
		}
		stage.OnAfterNavaidEquipmentDistancePropertyTypeDeleteCallback.OnAfterDelete(stage, navaidequipmentdistancepropertytype, frontNavaidEquipmentDistancePropertyType)
	}
}

func (navaidequipmentdistancetype *NavaidEquipmentDistanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidEquipmentDistanceTypeCreateCallback != nil {
		stage.OnAfterNavaidEquipmentDistanceTypeCreateCallback.OnAfterCreate(stage, navaidequipmentdistancetype)
	}
}

func (navaidequipmentdistancetype *NavaidEquipmentDistanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentDistanceTypeUpdateCallback != nil {
		var frontNavaidEquipmentDistanceType *NavaidEquipmentDistanceType
		if front != nil {
			frontNavaidEquipmentDistanceType, _ = front.(*NavaidEquipmentDistanceType)
		}
		stage.OnAfterNavaidEquipmentDistanceTypeUpdateCallback.OnAfterUpdate(stage, navaidequipmentdistancetype, frontNavaidEquipmentDistanceType)
	}
}

func (navaidequipmentdistancetype *NavaidEquipmentDistanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentDistanceTypeDeleteCallback != nil {
		var frontNavaidEquipmentDistanceType *NavaidEquipmentDistanceType
		if front != nil {
			frontNavaidEquipmentDistanceType, _ = front.(*NavaidEquipmentDistanceType)
		}
		stage.OnAfterNavaidEquipmentDistanceTypeDeleteCallback.OnAfterDelete(stage, navaidequipmentdistancetype, frontNavaidEquipmentDistanceType)
	}
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidEquipmentPropertyTypeCreateCallback != nil {
		stage.OnAfterNavaidEquipmentPropertyTypeCreateCallback.OnAfterCreate(stage, navaidequipmentpropertytype)
	}
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentPropertyTypeUpdateCallback != nil {
		var frontNavaidEquipmentPropertyType *NavaidEquipmentPropertyType
		if front != nil {
			frontNavaidEquipmentPropertyType, _ = front.(*NavaidEquipmentPropertyType)
		}
		stage.OnAfterNavaidEquipmentPropertyTypeUpdateCallback.OnAfterUpdate(stage, navaidequipmentpropertytype, frontNavaidEquipmentPropertyType)
	}
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentPropertyTypeDeleteCallback != nil {
		var frontNavaidEquipmentPropertyType *NavaidEquipmentPropertyType
		if front != nil {
			frontNavaidEquipmentPropertyType, _ = front.(*NavaidEquipmentPropertyType)
		}
		stage.OnAfterNavaidEquipmentPropertyTypeDeleteCallback.OnAfterDelete(stage, navaidequipmentpropertytype, frontNavaidEquipmentPropertyType)
	}
}

func (nonmovementareatimeslicepropertytype *NonMovementAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNonMovementAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterNonMovementAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, nonmovementareatimeslicepropertytype)
	}
}

func (nonmovementareatimeslicepropertytype *NonMovementAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNonMovementAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontNonMovementAreaTimeSlicePropertyType *NonMovementAreaTimeSlicePropertyType
		if front != nil {
			frontNonMovementAreaTimeSlicePropertyType, _ = front.(*NonMovementAreaTimeSlicePropertyType)
		}
		stage.OnAfterNonMovementAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, nonmovementareatimeslicepropertytype, frontNonMovementAreaTimeSlicePropertyType)
	}
}

func (nonmovementareatimeslicepropertytype *NonMovementAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNonMovementAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontNonMovementAreaTimeSlicePropertyType *NonMovementAreaTimeSlicePropertyType
		if front != nil {
			frontNonMovementAreaTimeSlicePropertyType, _ = front.(*NonMovementAreaTimeSlicePropertyType)
		}
		stage.OnAfterNonMovementAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, nonmovementareatimeslicepropertytype, frontNonMovementAreaTimeSlicePropertyType)
	}
}

func (nonmovementareatimeslicetype *NonMovementAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNonMovementAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterNonMovementAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, nonmovementareatimeslicetype)
	}
}

func (nonmovementareatimeslicetype *NonMovementAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNonMovementAreaTimeSliceTypeUpdateCallback != nil {
		var frontNonMovementAreaTimeSliceType *NonMovementAreaTimeSliceType
		if front != nil {
			frontNonMovementAreaTimeSliceType, _ = front.(*NonMovementAreaTimeSliceType)
		}
		stage.OnAfterNonMovementAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, nonmovementareatimeslicetype, frontNonMovementAreaTimeSliceType)
	}
}

func (nonmovementareatimeslicetype *NonMovementAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNonMovementAreaTimeSliceTypeDeleteCallback != nil {
		var frontNonMovementAreaTimeSliceType *NonMovementAreaTimeSliceType
		if front != nil {
			frontNonMovementAreaTimeSliceType, _ = front.(*NonMovementAreaTimeSliceType)
		}
		stage.OnAfterNonMovementAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, nonmovementareatimeslicetype, frontNonMovementAreaTimeSliceType)
	}
}

func (nonmovementareatype *NonMovementAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNonMovementAreaTypeCreateCallback != nil {
		stage.OnAfterNonMovementAreaTypeCreateCallback.OnAfterCreate(stage, nonmovementareatype)
	}
}

func (nonmovementareatype *NonMovementAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNonMovementAreaTypeUpdateCallback != nil {
		var frontNonMovementAreaType *NonMovementAreaType
		if front != nil {
			frontNonMovementAreaType, _ = front.(*NonMovementAreaType)
		}
		stage.OnAfterNonMovementAreaTypeUpdateCallback.OnAfterUpdate(stage, nonmovementareatype, frontNonMovementAreaType)
	}
}

func (nonmovementareatype *NonMovementAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNonMovementAreaTypeDeleteCallback != nil {
		var frontNonMovementAreaType *NonMovementAreaType
		if front != nil {
			frontNonMovementAreaType, _ = front.(*NonMovementAreaType)
		}
		stage.OnAfterNonMovementAreaTypeDeleteCallback.OnAfterDelete(stage, nonmovementareatype, frontNonMovementAreaType)
	}
}

func (notepropertytype *NotePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNotePropertyTypeCreateCallback != nil {
		stage.OnAfterNotePropertyTypeCreateCallback.OnAfterCreate(stage, notepropertytype)
	}
}

func (notepropertytype *NotePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePropertyTypeUpdateCallback != nil {
		var frontNotePropertyType *NotePropertyType
		if front != nil {
			frontNotePropertyType, _ = front.(*NotePropertyType)
		}
		stage.OnAfterNotePropertyTypeUpdateCallback.OnAfterUpdate(stage, notepropertytype, frontNotePropertyType)
	}
}

func (notepropertytype *NotePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePropertyTypeDeleteCallback != nil {
		var frontNotePropertyType *NotePropertyType
		if front != nil {
			frontNotePropertyType, _ = front.(*NotePropertyType)
		}
		stage.OnAfterNotePropertyTypeDeleteCallback.OnAfterDelete(stage, notepropertytype, frontNotePropertyType)
	}
}

func (notetype *NoteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteTypeCreateCallback != nil {
		stage.OnAfterNoteTypeCreateCallback.OnAfterCreate(stage, notetype)
	}
}

func (notetype *NoteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTypeUpdateCallback != nil {
		var frontNoteType *NoteType
		if front != nil {
			frontNoteType, _ = front.(*NoteType)
		}
		stage.OnAfterNoteTypeUpdateCallback.OnAfterUpdate(stage, notetype, frontNoteType)
	}
}

func (notetype *NoteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTypeDeleteCallback != nil {
		var frontNoteType *NoteType
		if front != nil {
			frontNoteType, _ = front.(*NoteType)
		}
		stage.OnAfterNoteTypeDeleteCallback.OnAfterDelete(stage, notetype, frontNoteType)
	}
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOnlineContactPropertyTypeCreateCallback != nil {
		stage.OnAfterOnlineContactPropertyTypeCreateCallback.OnAfterCreate(stage, onlinecontactpropertytype)
	}
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOnlineContactPropertyTypeUpdateCallback != nil {
		var frontOnlineContactPropertyType *OnlineContactPropertyType
		if front != nil {
			frontOnlineContactPropertyType, _ = front.(*OnlineContactPropertyType)
		}
		stage.OnAfterOnlineContactPropertyTypeUpdateCallback.OnAfterUpdate(stage, onlinecontactpropertytype, frontOnlineContactPropertyType)
	}
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOnlineContactPropertyTypeDeleteCallback != nil {
		var frontOnlineContactPropertyType *OnlineContactPropertyType
		if front != nil {
			frontOnlineContactPropertyType, _ = front.(*OnlineContactPropertyType)
		}
		stage.OnAfterOnlineContactPropertyTypeDeleteCallback.OnAfterDelete(stage, onlinecontactpropertytype, frontOnlineContactPropertyType)
	}
}

func (onlinecontacttype *OnlineContactType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOnlineContactTypeCreateCallback != nil {
		stage.OnAfterOnlineContactTypeCreateCallback.OnAfterCreate(stage, onlinecontacttype)
	}
}

func (onlinecontacttype *OnlineContactType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOnlineContactTypeUpdateCallback != nil {
		var frontOnlineContactType *OnlineContactType
		if front != nil {
			frontOnlineContactType, _ = front.(*OnlineContactType)
		}
		stage.OnAfterOnlineContactTypeUpdateCallback.OnAfterUpdate(stage, onlinecontacttype, frontOnlineContactType)
	}
}

func (onlinecontacttype *OnlineContactType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOnlineContactTypeDeleteCallback != nil {
		var frontOnlineContactType *OnlineContactType
		if front != nil {
			frontOnlineContactType, _ = front.(*OnlineContactType)
		}
		stage.OnAfterOnlineContactTypeDeleteCallback.OnAfterDelete(stage, onlinecontacttype, frontOnlineContactType)
	}
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrganisationAuthorityPropertyTypeCreateCallback != nil {
		stage.OnAfterOrganisationAuthorityPropertyTypeCreateCallback.OnAfterCreate(stage, organisationauthoritypropertytype)
	}
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityPropertyTypeUpdateCallback != nil {
		var frontOrganisationAuthorityPropertyType *OrganisationAuthorityPropertyType
		if front != nil {
			frontOrganisationAuthorityPropertyType, _ = front.(*OrganisationAuthorityPropertyType)
		}
		stage.OnAfterOrganisationAuthorityPropertyTypeUpdateCallback.OnAfterUpdate(stage, organisationauthoritypropertytype, frontOrganisationAuthorityPropertyType)
	}
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityPropertyTypeDeleteCallback != nil {
		var frontOrganisationAuthorityPropertyType *OrganisationAuthorityPropertyType
		if front != nil {
			frontOrganisationAuthorityPropertyType, _ = front.(*OrganisationAuthorityPropertyType)
		}
		stage.OnAfterOrganisationAuthorityPropertyTypeDeleteCallback.OnAfterDelete(stage, organisationauthoritypropertytype, frontOrganisationAuthorityPropertyType)
	}
}

func (passengerloadingbridgetimeslicepropertytype *PassengerLoadingBridgeTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPassengerLoadingBridgeTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterPassengerLoadingBridgeTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, passengerloadingbridgetimeslicepropertytype)
	}
}

func (passengerloadingbridgetimeslicepropertytype *PassengerLoadingBridgeTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerLoadingBridgeTimeSlicePropertyTypeUpdateCallback != nil {
		var frontPassengerLoadingBridgeTimeSlicePropertyType *PassengerLoadingBridgeTimeSlicePropertyType
		if front != nil {
			frontPassengerLoadingBridgeTimeSlicePropertyType, _ = front.(*PassengerLoadingBridgeTimeSlicePropertyType)
		}
		stage.OnAfterPassengerLoadingBridgeTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, passengerloadingbridgetimeslicepropertytype, frontPassengerLoadingBridgeTimeSlicePropertyType)
	}
}

func (passengerloadingbridgetimeslicepropertytype *PassengerLoadingBridgeTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerLoadingBridgeTimeSlicePropertyTypeDeleteCallback != nil {
		var frontPassengerLoadingBridgeTimeSlicePropertyType *PassengerLoadingBridgeTimeSlicePropertyType
		if front != nil {
			frontPassengerLoadingBridgeTimeSlicePropertyType, _ = front.(*PassengerLoadingBridgeTimeSlicePropertyType)
		}
		stage.OnAfterPassengerLoadingBridgeTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, passengerloadingbridgetimeslicepropertytype, frontPassengerLoadingBridgeTimeSlicePropertyType)
	}
}

func (passengerloadingbridgetimeslicetype *PassengerLoadingBridgeTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPassengerLoadingBridgeTimeSliceTypeCreateCallback != nil {
		stage.OnAfterPassengerLoadingBridgeTimeSliceTypeCreateCallback.OnAfterCreate(stage, passengerloadingbridgetimeslicetype)
	}
}

func (passengerloadingbridgetimeslicetype *PassengerLoadingBridgeTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerLoadingBridgeTimeSliceTypeUpdateCallback != nil {
		var frontPassengerLoadingBridgeTimeSliceType *PassengerLoadingBridgeTimeSliceType
		if front != nil {
			frontPassengerLoadingBridgeTimeSliceType, _ = front.(*PassengerLoadingBridgeTimeSliceType)
		}
		stage.OnAfterPassengerLoadingBridgeTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, passengerloadingbridgetimeslicetype, frontPassengerLoadingBridgeTimeSliceType)
	}
}

func (passengerloadingbridgetimeslicetype *PassengerLoadingBridgeTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerLoadingBridgeTimeSliceTypeDeleteCallback != nil {
		var frontPassengerLoadingBridgeTimeSliceType *PassengerLoadingBridgeTimeSliceType
		if front != nil {
			frontPassengerLoadingBridgeTimeSliceType, _ = front.(*PassengerLoadingBridgeTimeSliceType)
		}
		stage.OnAfterPassengerLoadingBridgeTimeSliceTypeDeleteCallback.OnAfterDelete(stage, passengerloadingbridgetimeslicetype, frontPassengerLoadingBridgeTimeSliceType)
	}
}

func (passengerloadingbridgetype *PassengerLoadingBridgeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPassengerLoadingBridgeTypeCreateCallback != nil {
		stage.OnAfterPassengerLoadingBridgeTypeCreateCallback.OnAfterCreate(stage, passengerloadingbridgetype)
	}
}

func (passengerloadingbridgetype *PassengerLoadingBridgeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerLoadingBridgeTypeUpdateCallback != nil {
		var frontPassengerLoadingBridgeType *PassengerLoadingBridgeType
		if front != nil {
			frontPassengerLoadingBridgeType, _ = front.(*PassengerLoadingBridgeType)
		}
		stage.OnAfterPassengerLoadingBridgeTypeUpdateCallback.OnAfterUpdate(stage, passengerloadingbridgetype, frontPassengerLoadingBridgeType)
	}
}

func (passengerloadingbridgetype *PassengerLoadingBridgeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerLoadingBridgeTypeDeleteCallback != nil {
		var frontPassengerLoadingBridgeType *PassengerLoadingBridgeType
		if front != nil {
			frontPassengerLoadingBridgeType, _ = front.(*PassengerLoadingBridgeType)
		}
		stage.OnAfterPassengerLoadingBridgeTypeDeleteCallback.OnAfterDelete(stage, passengerloadingbridgetype, frontPassengerLoadingBridgeType)
	}
}

func (pointtype *PointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointTypeCreateCallback != nil {
		stage.OnAfterPointTypeCreateCallback.OnAfterCreate(stage, pointtype)
	}
}

func (pointtype *PointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointTypeUpdateCallback != nil {
		var frontPointType *PointType
		if front != nil {
			frontPointType, _ = front.(*PointType)
		}
		stage.OnAfterPointTypeUpdateCallback.OnAfterUpdate(stage, pointtype, frontPointType)
	}
}

func (pointtype *PointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointTypeDeleteCallback != nil {
		var frontPointType *PointType
		if front != nil {
			frontPointType, _ = front.(*PointType)
		}
		stage.OnAfterPointTypeDeleteCallback.OnAfterDelete(stage, pointtype, frontPointType)
	}
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPostalAddressPropertyTypeCreateCallback != nil {
		stage.OnAfterPostalAddressPropertyTypeCreateCallback.OnAfterCreate(stage, postaladdresspropertytype)
	}
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPostalAddressPropertyTypeUpdateCallback != nil {
		var frontPostalAddressPropertyType *PostalAddressPropertyType
		if front != nil {
			frontPostalAddressPropertyType, _ = front.(*PostalAddressPropertyType)
		}
		stage.OnAfterPostalAddressPropertyTypeUpdateCallback.OnAfterUpdate(stage, postaladdresspropertytype, frontPostalAddressPropertyType)
	}
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPostalAddressPropertyTypeDeleteCallback != nil {
		var frontPostalAddressPropertyType *PostalAddressPropertyType
		if front != nil {
			frontPostalAddressPropertyType, _ = front.(*PostalAddressPropertyType)
		}
		stage.OnAfterPostalAddressPropertyTypeDeleteCallback.OnAfterDelete(stage, postaladdresspropertytype, frontPostalAddressPropertyType)
	}
}

func (postaladdresstype *PostalAddressType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPostalAddressTypeCreateCallback != nil {
		stage.OnAfterPostalAddressTypeCreateCallback.OnAfterCreate(stage, postaladdresstype)
	}
}

func (postaladdresstype *PostalAddressType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPostalAddressTypeUpdateCallback != nil {
		var frontPostalAddressType *PostalAddressType
		if front != nil {
			frontPostalAddressType, _ = front.(*PostalAddressType)
		}
		stage.OnAfterPostalAddressTypeUpdateCallback.OnAfterUpdate(stage, postaladdresstype, frontPostalAddressType)
	}
}

func (postaladdresstype *PostalAddressType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPostalAddressTypeDeleteCallback != nil {
		var frontPostalAddressType *PostalAddressType
		if front != nil {
			frontPostalAddressType, _ = front.(*PostalAddressType)
		}
		stage.OnAfterPostalAddressTypeDeleteCallback.OnAfterDelete(stage, postaladdresstype, frontPostalAddressType)
	}
}

func (ridgepropertytype *RidgePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRidgePropertyTypeCreateCallback != nil {
		stage.OnAfterRidgePropertyTypeCreateCallback.OnAfterCreate(stage, ridgepropertytype)
	}
}

func (ridgepropertytype *RidgePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRidgePropertyTypeUpdateCallback != nil {
		var frontRidgePropertyType *RidgePropertyType
		if front != nil {
			frontRidgePropertyType, _ = front.(*RidgePropertyType)
		}
		stage.OnAfterRidgePropertyTypeUpdateCallback.OnAfterUpdate(stage, ridgepropertytype, frontRidgePropertyType)
	}
}

func (ridgepropertytype *RidgePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRidgePropertyTypeDeleteCallback != nil {
		var frontRidgePropertyType *RidgePropertyType
		if front != nil {
			frontRidgePropertyType, _ = front.(*RidgePropertyType)
		}
		stage.OnAfterRidgePropertyTypeDeleteCallback.OnAfterDelete(stage, ridgepropertytype, frontRidgePropertyType)
	}
}

func (ridgetype *RidgeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRidgeTypeCreateCallback != nil {
		stage.OnAfterRidgeTypeCreateCallback.OnAfterCreate(stage, ridgetype)
	}
}

func (ridgetype *RidgeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRidgeTypeUpdateCallback != nil {
		var frontRidgeType *RidgeType
		if front != nil {
			frontRidgeType, _ = front.(*RidgeType)
		}
		stage.OnAfterRidgeTypeUpdateCallback.OnAfterUpdate(stage, ridgetype, frontRidgeType)
	}
}

func (ridgetype *RidgeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRidgeTypeDeleteCallback != nil {
		var frontRidgeType *RidgeType
		if front != nil {
			frontRidgeType, _ = front.(*RidgeType)
		}
		stage.OnAfterRidgeTypeDeleteCallback.OnAfterDelete(stage, ridgetype, frontRidgeType)
	}
}

func (roadtimeslicepropertytype *RoadTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRoadTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRoadTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, roadtimeslicepropertytype)
	}
}

func (roadtimeslicepropertytype *RoadTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoadTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRoadTimeSlicePropertyType *RoadTimeSlicePropertyType
		if front != nil {
			frontRoadTimeSlicePropertyType, _ = front.(*RoadTimeSlicePropertyType)
		}
		stage.OnAfterRoadTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, roadtimeslicepropertytype, frontRoadTimeSlicePropertyType)
	}
}

func (roadtimeslicepropertytype *RoadTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoadTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRoadTimeSlicePropertyType *RoadTimeSlicePropertyType
		if front != nil {
			frontRoadTimeSlicePropertyType, _ = front.(*RoadTimeSlicePropertyType)
		}
		stage.OnAfterRoadTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, roadtimeslicepropertytype, frontRoadTimeSlicePropertyType)
	}
}

func (roadtimeslicetype *RoadTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRoadTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRoadTimeSliceTypeCreateCallback.OnAfterCreate(stage, roadtimeslicetype)
	}
}

func (roadtimeslicetype *RoadTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoadTimeSliceTypeUpdateCallback != nil {
		var frontRoadTimeSliceType *RoadTimeSliceType
		if front != nil {
			frontRoadTimeSliceType, _ = front.(*RoadTimeSliceType)
		}
		stage.OnAfterRoadTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, roadtimeslicetype, frontRoadTimeSliceType)
	}
}

func (roadtimeslicetype *RoadTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoadTimeSliceTypeDeleteCallback != nil {
		var frontRoadTimeSliceType *RoadTimeSliceType
		if front != nil {
			frontRoadTimeSliceType, _ = front.(*RoadTimeSliceType)
		}
		stage.OnAfterRoadTimeSliceTypeDeleteCallback.OnAfterDelete(stage, roadtimeslicetype, frontRoadTimeSliceType)
	}
}

func (roadtype *RoadType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRoadTypeCreateCallback != nil {
		stage.OnAfterRoadTypeCreateCallback.OnAfterCreate(stage, roadtype)
	}
}

func (roadtype *RoadType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoadTypeUpdateCallback != nil {
		var frontRoadType *RoadType
		if front != nil {
			frontRoadType, _ = front.(*RoadType)
		}
		stage.OnAfterRoadTypeUpdateCallback.OnAfterUpdate(stage, roadtype, frontRoadType)
	}
}

func (roadtype *RoadType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoadTypeDeleteCallback != nil {
		var frontRoadType *RoadType
		if front != nil {
			frontRoadType, _ = front.(*RoadType)
		}
		stage.OnAfterRoadTypeDeleteCallback.OnAfterDelete(stage, roadtype, frontRoadType)
	}
}

func (runwayblastpadtimeslicepropertytype *RunwayBlastPadTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayBlastPadTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayBlastPadTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwayblastpadtimeslicepropertytype)
	}
}

func (runwayblastpadtimeslicepropertytype *RunwayBlastPadTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayBlastPadTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayBlastPadTimeSlicePropertyType *RunwayBlastPadTimeSlicePropertyType
		if front != nil {
			frontRunwayBlastPadTimeSlicePropertyType, _ = front.(*RunwayBlastPadTimeSlicePropertyType)
		}
		stage.OnAfterRunwayBlastPadTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwayblastpadtimeslicepropertytype, frontRunwayBlastPadTimeSlicePropertyType)
	}
}

func (runwayblastpadtimeslicepropertytype *RunwayBlastPadTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayBlastPadTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayBlastPadTimeSlicePropertyType *RunwayBlastPadTimeSlicePropertyType
		if front != nil {
			frontRunwayBlastPadTimeSlicePropertyType, _ = front.(*RunwayBlastPadTimeSlicePropertyType)
		}
		stage.OnAfterRunwayBlastPadTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwayblastpadtimeslicepropertytype, frontRunwayBlastPadTimeSlicePropertyType)
	}
}

func (runwayblastpadtimeslicetype *RunwayBlastPadTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayBlastPadTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayBlastPadTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwayblastpadtimeslicetype)
	}
}

func (runwayblastpadtimeslicetype *RunwayBlastPadTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayBlastPadTimeSliceTypeUpdateCallback != nil {
		var frontRunwayBlastPadTimeSliceType *RunwayBlastPadTimeSliceType
		if front != nil {
			frontRunwayBlastPadTimeSliceType, _ = front.(*RunwayBlastPadTimeSliceType)
		}
		stage.OnAfterRunwayBlastPadTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwayblastpadtimeslicetype, frontRunwayBlastPadTimeSliceType)
	}
}

func (runwayblastpadtimeslicetype *RunwayBlastPadTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayBlastPadTimeSliceTypeDeleteCallback != nil {
		var frontRunwayBlastPadTimeSliceType *RunwayBlastPadTimeSliceType
		if front != nil {
			frontRunwayBlastPadTimeSliceType, _ = front.(*RunwayBlastPadTimeSliceType)
		}
		stage.OnAfterRunwayBlastPadTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwayblastpadtimeslicetype, frontRunwayBlastPadTimeSliceType)
	}
}

func (runwayblastpadtype *RunwayBlastPadType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayBlastPadTypeCreateCallback != nil {
		stage.OnAfterRunwayBlastPadTypeCreateCallback.OnAfterCreate(stage, runwayblastpadtype)
	}
}

func (runwayblastpadtype *RunwayBlastPadType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayBlastPadTypeUpdateCallback != nil {
		var frontRunwayBlastPadType *RunwayBlastPadType
		if front != nil {
			frontRunwayBlastPadType, _ = front.(*RunwayBlastPadType)
		}
		stage.OnAfterRunwayBlastPadTypeUpdateCallback.OnAfterUpdate(stage, runwayblastpadtype, frontRunwayBlastPadType)
	}
}

func (runwayblastpadtype *RunwayBlastPadType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayBlastPadTypeDeleteCallback != nil {
		var frontRunwayBlastPadType *RunwayBlastPadType
		if front != nil {
			frontRunwayBlastPadType, _ = front.(*RunwayBlastPadType)
		}
		stage.OnAfterRunwayBlastPadTypeDeleteCallback.OnAfterDelete(stage, runwayblastpadtype, frontRunwayBlastPadType)
	}
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayCentrelinePointPropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayCentrelinePointPropertyTypeCreateCallback.OnAfterCreate(stage, runwaycentrelinepointpropertytype)
	}
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayCentrelinePointPropertyTypeUpdateCallback != nil {
		var frontRunwayCentrelinePointPropertyType *RunwayCentrelinePointPropertyType
		if front != nil {
			frontRunwayCentrelinePointPropertyType, _ = front.(*RunwayCentrelinePointPropertyType)
		}
		stage.OnAfterRunwayCentrelinePointPropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaycentrelinepointpropertytype, frontRunwayCentrelinePointPropertyType)
	}
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayCentrelinePointPropertyTypeDeleteCallback != nil {
		var frontRunwayCentrelinePointPropertyType *RunwayCentrelinePointPropertyType
		if front != nil {
			frontRunwayCentrelinePointPropertyType, _ = front.(*RunwayCentrelinePointPropertyType)
		}
		stage.OnAfterRunwayCentrelinePointPropertyTypeDeleteCallback.OnAfterDelete(stage, runwaycentrelinepointpropertytype, frontRunwayCentrelinePointPropertyType)
	}
}

func (runwaycentrelinepointtimeslicepropertytype *RunwayCentrelinePointTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayCentrelinePointTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayCentrelinePointTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwaycentrelinepointtimeslicepropertytype)
	}
}

func (runwaycentrelinepointtimeslicepropertytype *RunwayCentrelinePointTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayCentrelinePointTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayCentrelinePointTimeSlicePropertyType *RunwayCentrelinePointTimeSlicePropertyType
		if front != nil {
			frontRunwayCentrelinePointTimeSlicePropertyType, _ = front.(*RunwayCentrelinePointTimeSlicePropertyType)
		}
		stage.OnAfterRunwayCentrelinePointTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaycentrelinepointtimeslicepropertytype, frontRunwayCentrelinePointTimeSlicePropertyType)
	}
}

func (runwaycentrelinepointtimeslicepropertytype *RunwayCentrelinePointTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayCentrelinePointTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayCentrelinePointTimeSlicePropertyType *RunwayCentrelinePointTimeSlicePropertyType
		if front != nil {
			frontRunwayCentrelinePointTimeSlicePropertyType, _ = front.(*RunwayCentrelinePointTimeSlicePropertyType)
		}
		stage.OnAfterRunwayCentrelinePointTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwaycentrelinepointtimeslicepropertytype, frontRunwayCentrelinePointTimeSlicePropertyType)
	}
}

func (runwaycentrelinepointtimeslicetype *RunwayCentrelinePointTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayCentrelinePointTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayCentrelinePointTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwaycentrelinepointtimeslicetype)
	}
}

func (runwaycentrelinepointtimeslicetype *RunwayCentrelinePointTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayCentrelinePointTimeSliceTypeUpdateCallback != nil {
		var frontRunwayCentrelinePointTimeSliceType *RunwayCentrelinePointTimeSliceType
		if front != nil {
			frontRunwayCentrelinePointTimeSliceType, _ = front.(*RunwayCentrelinePointTimeSliceType)
		}
		stage.OnAfterRunwayCentrelinePointTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwaycentrelinepointtimeslicetype, frontRunwayCentrelinePointTimeSliceType)
	}
}

func (runwaycentrelinepointtimeslicetype *RunwayCentrelinePointTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayCentrelinePointTimeSliceTypeDeleteCallback != nil {
		var frontRunwayCentrelinePointTimeSliceType *RunwayCentrelinePointTimeSliceType
		if front != nil {
			frontRunwayCentrelinePointTimeSliceType, _ = front.(*RunwayCentrelinePointTimeSliceType)
		}
		stage.OnAfterRunwayCentrelinePointTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwaycentrelinepointtimeslicetype, frontRunwayCentrelinePointTimeSliceType)
	}
}

func (runwaycentrelinepointtype *RunwayCentrelinePointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayCentrelinePointTypeCreateCallback != nil {
		stage.OnAfterRunwayCentrelinePointTypeCreateCallback.OnAfterCreate(stage, runwaycentrelinepointtype)
	}
}

func (runwaycentrelinepointtype *RunwayCentrelinePointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayCentrelinePointTypeUpdateCallback != nil {
		var frontRunwayCentrelinePointType *RunwayCentrelinePointType
		if front != nil {
			frontRunwayCentrelinePointType, _ = front.(*RunwayCentrelinePointType)
		}
		stage.OnAfterRunwayCentrelinePointTypeUpdateCallback.OnAfterUpdate(stage, runwaycentrelinepointtype, frontRunwayCentrelinePointType)
	}
}

func (runwaycentrelinepointtype *RunwayCentrelinePointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayCentrelinePointTypeDeleteCallback != nil {
		var frontRunwayCentrelinePointType *RunwayCentrelinePointType
		if front != nil {
			frontRunwayCentrelinePointType, _ = front.(*RunwayCentrelinePointType)
		}
		stage.OnAfterRunwayCentrelinePointTypeDeleteCallback.OnAfterDelete(stage, runwaycentrelinepointtype, frontRunwayCentrelinePointType)
	}
}

func (runwaycontaminationpropertytype *RunwayContaminationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayContaminationPropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayContaminationPropertyTypeCreateCallback.OnAfterCreate(stage, runwaycontaminationpropertytype)
	}
}

func (runwaycontaminationpropertytype *RunwayContaminationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayContaminationPropertyTypeUpdateCallback != nil {
		var frontRunwayContaminationPropertyType *RunwayContaminationPropertyType
		if front != nil {
			frontRunwayContaminationPropertyType, _ = front.(*RunwayContaminationPropertyType)
		}
		stage.OnAfterRunwayContaminationPropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaycontaminationpropertytype, frontRunwayContaminationPropertyType)
	}
}

func (runwaycontaminationpropertytype *RunwayContaminationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayContaminationPropertyTypeDeleteCallback != nil {
		var frontRunwayContaminationPropertyType *RunwayContaminationPropertyType
		if front != nil {
			frontRunwayContaminationPropertyType, _ = front.(*RunwayContaminationPropertyType)
		}
		stage.OnAfterRunwayContaminationPropertyTypeDeleteCallback.OnAfterDelete(stage, runwaycontaminationpropertytype, frontRunwayContaminationPropertyType)
	}
}

func (runwaycontaminationtype *RunwayContaminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayContaminationTypeCreateCallback != nil {
		stage.OnAfterRunwayContaminationTypeCreateCallback.OnAfterCreate(stage, runwaycontaminationtype)
	}
}

func (runwaycontaminationtype *RunwayContaminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayContaminationTypeUpdateCallback != nil {
		var frontRunwayContaminationType *RunwayContaminationType
		if front != nil {
			frontRunwayContaminationType, _ = front.(*RunwayContaminationType)
		}
		stage.OnAfterRunwayContaminationTypeUpdateCallback.OnAfterUpdate(stage, runwaycontaminationtype, frontRunwayContaminationType)
	}
}

func (runwaycontaminationtype *RunwayContaminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayContaminationTypeDeleteCallback != nil {
		var frontRunwayContaminationType *RunwayContaminationType
		if front != nil {
			frontRunwayContaminationType, _ = front.(*RunwayContaminationType)
		}
		stage.OnAfterRunwayContaminationTypeDeleteCallback.OnAfterDelete(stage, runwaycontaminationtype, frontRunwayContaminationType)
	}
}

func (runwaydeclareddistancepropertytype *RunwayDeclaredDistancePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDeclaredDistancePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayDeclaredDistancePropertyTypeCreateCallback.OnAfterCreate(stage, runwaydeclareddistancepropertytype)
	}
}

func (runwaydeclareddistancepropertytype *RunwayDeclaredDistancePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDeclaredDistancePropertyTypeUpdateCallback != nil {
		var frontRunwayDeclaredDistancePropertyType *RunwayDeclaredDistancePropertyType
		if front != nil {
			frontRunwayDeclaredDistancePropertyType, _ = front.(*RunwayDeclaredDistancePropertyType)
		}
		stage.OnAfterRunwayDeclaredDistancePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaydeclareddistancepropertytype, frontRunwayDeclaredDistancePropertyType)
	}
}

func (runwaydeclareddistancepropertytype *RunwayDeclaredDistancePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDeclaredDistancePropertyTypeDeleteCallback != nil {
		var frontRunwayDeclaredDistancePropertyType *RunwayDeclaredDistancePropertyType
		if front != nil {
			frontRunwayDeclaredDistancePropertyType, _ = front.(*RunwayDeclaredDistancePropertyType)
		}
		stage.OnAfterRunwayDeclaredDistancePropertyTypeDeleteCallback.OnAfterDelete(stage, runwaydeclareddistancepropertytype, frontRunwayDeclaredDistancePropertyType)
	}
}

func (runwaydeclareddistancetype *RunwayDeclaredDistanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDeclaredDistanceTypeCreateCallback != nil {
		stage.OnAfterRunwayDeclaredDistanceTypeCreateCallback.OnAfterCreate(stage, runwaydeclareddistancetype)
	}
}

func (runwaydeclareddistancetype *RunwayDeclaredDistanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDeclaredDistanceTypeUpdateCallback != nil {
		var frontRunwayDeclaredDistanceType *RunwayDeclaredDistanceType
		if front != nil {
			frontRunwayDeclaredDistanceType, _ = front.(*RunwayDeclaredDistanceType)
		}
		stage.OnAfterRunwayDeclaredDistanceTypeUpdateCallback.OnAfterUpdate(stage, runwaydeclareddistancetype, frontRunwayDeclaredDistanceType)
	}
}

func (runwaydeclareddistancetype *RunwayDeclaredDistanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDeclaredDistanceTypeDeleteCallback != nil {
		var frontRunwayDeclaredDistanceType *RunwayDeclaredDistanceType
		if front != nil {
			frontRunwayDeclaredDistanceType, _ = front.(*RunwayDeclaredDistanceType)
		}
		stage.OnAfterRunwayDeclaredDistanceTypeDeleteCallback.OnAfterDelete(stage, runwaydeclareddistancetype, frontRunwayDeclaredDistanceType)
	}
}

func (runwaydeclareddistancevaluepropertytype *RunwayDeclaredDistanceValuePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDeclaredDistanceValuePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayDeclaredDistanceValuePropertyTypeCreateCallback.OnAfterCreate(stage, runwaydeclareddistancevaluepropertytype)
	}
}

func (runwaydeclareddistancevaluepropertytype *RunwayDeclaredDistanceValuePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDeclaredDistanceValuePropertyTypeUpdateCallback != nil {
		var frontRunwayDeclaredDistanceValuePropertyType *RunwayDeclaredDistanceValuePropertyType
		if front != nil {
			frontRunwayDeclaredDistanceValuePropertyType, _ = front.(*RunwayDeclaredDistanceValuePropertyType)
		}
		stage.OnAfterRunwayDeclaredDistanceValuePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaydeclareddistancevaluepropertytype, frontRunwayDeclaredDistanceValuePropertyType)
	}
}

func (runwaydeclareddistancevaluepropertytype *RunwayDeclaredDistanceValuePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDeclaredDistanceValuePropertyTypeDeleteCallback != nil {
		var frontRunwayDeclaredDistanceValuePropertyType *RunwayDeclaredDistanceValuePropertyType
		if front != nil {
			frontRunwayDeclaredDistanceValuePropertyType, _ = front.(*RunwayDeclaredDistanceValuePropertyType)
		}
		stage.OnAfterRunwayDeclaredDistanceValuePropertyTypeDeleteCallback.OnAfterDelete(stage, runwaydeclareddistancevaluepropertytype, frontRunwayDeclaredDistanceValuePropertyType)
	}
}

func (runwaydeclareddistancevaluetype *RunwayDeclaredDistanceValueType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDeclaredDistanceValueTypeCreateCallback != nil {
		stage.OnAfterRunwayDeclaredDistanceValueTypeCreateCallback.OnAfterCreate(stage, runwaydeclareddistancevaluetype)
	}
}

func (runwaydeclareddistancevaluetype *RunwayDeclaredDistanceValueType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDeclaredDistanceValueTypeUpdateCallback != nil {
		var frontRunwayDeclaredDistanceValueType *RunwayDeclaredDistanceValueType
		if front != nil {
			frontRunwayDeclaredDistanceValueType, _ = front.(*RunwayDeclaredDistanceValueType)
		}
		stage.OnAfterRunwayDeclaredDistanceValueTypeUpdateCallback.OnAfterUpdate(stage, runwaydeclareddistancevaluetype, frontRunwayDeclaredDistanceValueType)
	}
}

func (runwaydeclareddistancevaluetype *RunwayDeclaredDistanceValueType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDeclaredDistanceValueTypeDeleteCallback != nil {
		var frontRunwayDeclaredDistanceValueType *RunwayDeclaredDistanceValueType
		if front != nil {
			frontRunwayDeclaredDistanceValueType, _ = front.(*RunwayDeclaredDistanceValueType)
		}
		stage.OnAfterRunwayDeclaredDistanceValueTypeDeleteCallback.OnAfterDelete(stage, runwaydeclareddistancevaluetype, frontRunwayDeclaredDistanceValueType)
	}
}

func (runwaydirectionlightsystemtimeslicepropertytype *RunwayDirectionLightSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDirectionLightSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayDirectionLightSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwaydirectionlightsystemtimeslicepropertytype)
	}
}

func (runwaydirectionlightsystemtimeslicepropertytype *RunwayDirectionLightSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionLightSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayDirectionLightSystemTimeSlicePropertyType *RunwayDirectionLightSystemTimeSlicePropertyType
		if front != nil {
			frontRunwayDirectionLightSystemTimeSlicePropertyType, _ = front.(*RunwayDirectionLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterRunwayDirectionLightSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaydirectionlightsystemtimeslicepropertytype, frontRunwayDirectionLightSystemTimeSlicePropertyType)
	}
}

func (runwaydirectionlightsystemtimeslicepropertytype *RunwayDirectionLightSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionLightSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayDirectionLightSystemTimeSlicePropertyType *RunwayDirectionLightSystemTimeSlicePropertyType
		if front != nil {
			frontRunwayDirectionLightSystemTimeSlicePropertyType, _ = front.(*RunwayDirectionLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterRunwayDirectionLightSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwaydirectionlightsystemtimeslicepropertytype, frontRunwayDirectionLightSystemTimeSlicePropertyType)
	}
}

func (runwaydirectionlightsystemtimeslicetype *RunwayDirectionLightSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDirectionLightSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayDirectionLightSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwaydirectionlightsystemtimeslicetype)
	}
}

func (runwaydirectionlightsystemtimeslicetype *RunwayDirectionLightSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionLightSystemTimeSliceTypeUpdateCallback != nil {
		var frontRunwayDirectionLightSystemTimeSliceType *RunwayDirectionLightSystemTimeSliceType
		if front != nil {
			frontRunwayDirectionLightSystemTimeSliceType, _ = front.(*RunwayDirectionLightSystemTimeSliceType)
		}
		stage.OnAfterRunwayDirectionLightSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwaydirectionlightsystemtimeslicetype, frontRunwayDirectionLightSystemTimeSliceType)
	}
}

func (runwaydirectionlightsystemtimeslicetype *RunwayDirectionLightSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionLightSystemTimeSliceTypeDeleteCallback != nil {
		var frontRunwayDirectionLightSystemTimeSliceType *RunwayDirectionLightSystemTimeSliceType
		if front != nil {
			frontRunwayDirectionLightSystemTimeSliceType, _ = front.(*RunwayDirectionLightSystemTimeSliceType)
		}
		stage.OnAfterRunwayDirectionLightSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwaydirectionlightsystemtimeslicetype, frontRunwayDirectionLightSystemTimeSliceType)
	}
}

func (runwaydirectionlightsystemtype *RunwayDirectionLightSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDirectionLightSystemTypeCreateCallback != nil {
		stage.OnAfterRunwayDirectionLightSystemTypeCreateCallback.OnAfterCreate(stage, runwaydirectionlightsystemtype)
	}
}

func (runwaydirectionlightsystemtype *RunwayDirectionLightSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionLightSystemTypeUpdateCallback != nil {
		var frontRunwayDirectionLightSystemType *RunwayDirectionLightSystemType
		if front != nil {
			frontRunwayDirectionLightSystemType, _ = front.(*RunwayDirectionLightSystemType)
		}
		stage.OnAfterRunwayDirectionLightSystemTypeUpdateCallback.OnAfterUpdate(stage, runwaydirectionlightsystemtype, frontRunwayDirectionLightSystemType)
	}
}

func (runwaydirectionlightsystemtype *RunwayDirectionLightSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionLightSystemTypeDeleteCallback != nil {
		var frontRunwayDirectionLightSystemType *RunwayDirectionLightSystemType
		if front != nil {
			frontRunwayDirectionLightSystemType, _ = front.(*RunwayDirectionLightSystemType)
		}
		stage.OnAfterRunwayDirectionLightSystemTypeDeleteCallback.OnAfterDelete(stage, runwaydirectionlightsystemtype, frontRunwayDirectionLightSystemType)
	}
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDirectionPropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayDirectionPropertyTypeCreateCallback.OnAfterCreate(stage, runwaydirectionpropertytype)
	}
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionPropertyTypeUpdateCallback != nil {
		var frontRunwayDirectionPropertyType *RunwayDirectionPropertyType
		if front != nil {
			frontRunwayDirectionPropertyType, _ = front.(*RunwayDirectionPropertyType)
		}
		stage.OnAfterRunwayDirectionPropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaydirectionpropertytype, frontRunwayDirectionPropertyType)
	}
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionPropertyTypeDeleteCallback != nil {
		var frontRunwayDirectionPropertyType *RunwayDirectionPropertyType
		if front != nil {
			frontRunwayDirectionPropertyType, _ = front.(*RunwayDirectionPropertyType)
		}
		stage.OnAfterRunwayDirectionPropertyTypeDeleteCallback.OnAfterDelete(stage, runwaydirectionpropertytype, frontRunwayDirectionPropertyType)
	}
}

func (runwaydirectiontimeslicepropertytype *RunwayDirectionTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDirectionTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayDirectionTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwaydirectiontimeslicepropertytype)
	}
}

func (runwaydirectiontimeslicepropertytype *RunwayDirectionTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayDirectionTimeSlicePropertyType *RunwayDirectionTimeSlicePropertyType
		if front != nil {
			frontRunwayDirectionTimeSlicePropertyType, _ = front.(*RunwayDirectionTimeSlicePropertyType)
		}
		stage.OnAfterRunwayDirectionTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaydirectiontimeslicepropertytype, frontRunwayDirectionTimeSlicePropertyType)
	}
}

func (runwaydirectiontimeslicepropertytype *RunwayDirectionTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayDirectionTimeSlicePropertyType *RunwayDirectionTimeSlicePropertyType
		if front != nil {
			frontRunwayDirectionTimeSlicePropertyType, _ = front.(*RunwayDirectionTimeSlicePropertyType)
		}
		stage.OnAfterRunwayDirectionTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwaydirectiontimeslicepropertytype, frontRunwayDirectionTimeSlicePropertyType)
	}
}

func (runwaydirectiontimeslicetype *RunwayDirectionTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDirectionTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayDirectionTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwaydirectiontimeslicetype)
	}
}

func (runwaydirectiontimeslicetype *RunwayDirectionTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionTimeSliceTypeUpdateCallback != nil {
		var frontRunwayDirectionTimeSliceType *RunwayDirectionTimeSliceType
		if front != nil {
			frontRunwayDirectionTimeSliceType, _ = front.(*RunwayDirectionTimeSliceType)
		}
		stage.OnAfterRunwayDirectionTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwaydirectiontimeslicetype, frontRunwayDirectionTimeSliceType)
	}
}

func (runwaydirectiontimeslicetype *RunwayDirectionTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionTimeSliceTypeDeleteCallback != nil {
		var frontRunwayDirectionTimeSliceType *RunwayDirectionTimeSliceType
		if front != nil {
			frontRunwayDirectionTimeSliceType, _ = front.(*RunwayDirectionTimeSliceType)
		}
		stage.OnAfterRunwayDirectionTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwaydirectiontimeslicetype, frontRunwayDirectionTimeSliceType)
	}
}

func (runwaydirectiontype *RunwayDirectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDirectionTypeCreateCallback != nil {
		stage.OnAfterRunwayDirectionTypeCreateCallback.OnAfterCreate(stage, runwaydirectiontype)
	}
}

func (runwaydirectiontype *RunwayDirectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionTypeUpdateCallback != nil {
		var frontRunwayDirectionType *RunwayDirectionType
		if front != nil {
			frontRunwayDirectionType, _ = front.(*RunwayDirectionType)
		}
		stage.OnAfterRunwayDirectionTypeUpdateCallback.OnAfterUpdate(stage, runwaydirectiontype, frontRunwayDirectionType)
	}
}

func (runwaydirectiontype *RunwayDirectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionTypeDeleteCallback != nil {
		var frontRunwayDirectionType *RunwayDirectionType
		if front != nil {
			frontRunwayDirectionType, _ = front.(*RunwayDirectionType)
		}
		stage.OnAfterRunwayDirectionTypeDeleteCallback.OnAfterDelete(stage, runwaydirectiontype, frontRunwayDirectionType)
	}
}

func (runwayelementpropertytype *RunwayElementPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayElementPropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayElementPropertyTypeCreateCallback.OnAfterCreate(stage, runwayelementpropertytype)
	}
}

func (runwayelementpropertytype *RunwayElementPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayElementPropertyTypeUpdateCallback != nil {
		var frontRunwayElementPropertyType *RunwayElementPropertyType
		if front != nil {
			frontRunwayElementPropertyType, _ = front.(*RunwayElementPropertyType)
		}
		stage.OnAfterRunwayElementPropertyTypeUpdateCallback.OnAfterUpdate(stage, runwayelementpropertytype, frontRunwayElementPropertyType)
	}
}

func (runwayelementpropertytype *RunwayElementPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayElementPropertyTypeDeleteCallback != nil {
		var frontRunwayElementPropertyType *RunwayElementPropertyType
		if front != nil {
			frontRunwayElementPropertyType, _ = front.(*RunwayElementPropertyType)
		}
		stage.OnAfterRunwayElementPropertyTypeDeleteCallback.OnAfterDelete(stage, runwayelementpropertytype, frontRunwayElementPropertyType)
	}
}

func (runwayelementtimeslicepropertytype *RunwayElementTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayElementTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayElementTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwayelementtimeslicepropertytype)
	}
}

func (runwayelementtimeslicepropertytype *RunwayElementTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayElementTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayElementTimeSlicePropertyType *RunwayElementTimeSlicePropertyType
		if front != nil {
			frontRunwayElementTimeSlicePropertyType, _ = front.(*RunwayElementTimeSlicePropertyType)
		}
		stage.OnAfterRunwayElementTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwayelementtimeslicepropertytype, frontRunwayElementTimeSlicePropertyType)
	}
}

func (runwayelementtimeslicepropertytype *RunwayElementTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayElementTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayElementTimeSlicePropertyType *RunwayElementTimeSlicePropertyType
		if front != nil {
			frontRunwayElementTimeSlicePropertyType, _ = front.(*RunwayElementTimeSlicePropertyType)
		}
		stage.OnAfterRunwayElementTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwayelementtimeslicepropertytype, frontRunwayElementTimeSlicePropertyType)
	}
}

func (runwayelementtimeslicetype *RunwayElementTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayElementTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayElementTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwayelementtimeslicetype)
	}
}

func (runwayelementtimeslicetype *RunwayElementTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayElementTimeSliceTypeUpdateCallback != nil {
		var frontRunwayElementTimeSliceType *RunwayElementTimeSliceType
		if front != nil {
			frontRunwayElementTimeSliceType, _ = front.(*RunwayElementTimeSliceType)
		}
		stage.OnAfterRunwayElementTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwayelementtimeslicetype, frontRunwayElementTimeSliceType)
	}
}

func (runwayelementtimeslicetype *RunwayElementTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayElementTimeSliceTypeDeleteCallback != nil {
		var frontRunwayElementTimeSliceType *RunwayElementTimeSliceType
		if front != nil {
			frontRunwayElementTimeSliceType, _ = front.(*RunwayElementTimeSliceType)
		}
		stage.OnAfterRunwayElementTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwayelementtimeslicetype, frontRunwayElementTimeSliceType)
	}
}

func (runwayelementtype *RunwayElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayElementTypeCreateCallback != nil {
		stage.OnAfterRunwayElementTypeCreateCallback.OnAfterCreate(stage, runwayelementtype)
	}
}

func (runwayelementtype *RunwayElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayElementTypeUpdateCallback != nil {
		var frontRunwayElementType *RunwayElementType
		if front != nil {
			frontRunwayElementType, _ = front.(*RunwayElementType)
		}
		stage.OnAfterRunwayElementTypeUpdateCallback.OnAfterUpdate(stage, runwayelementtype, frontRunwayElementType)
	}
}

func (runwayelementtype *RunwayElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayElementTypeDeleteCallback != nil {
		var frontRunwayElementType *RunwayElementType
		if front != nil {
			frontRunwayElementType, _ = front.(*RunwayElementType)
		}
		stage.OnAfterRunwayElementTypeDeleteCallback.OnAfterDelete(stage, runwayelementtype, frontRunwayElementType)
	}
}

func (runwaymarkingtimeslicepropertytype *RunwayMarkingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayMarkingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayMarkingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwaymarkingtimeslicepropertytype)
	}
}

func (runwaymarkingtimeslicepropertytype *RunwayMarkingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayMarkingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayMarkingTimeSlicePropertyType *RunwayMarkingTimeSlicePropertyType
		if front != nil {
			frontRunwayMarkingTimeSlicePropertyType, _ = front.(*RunwayMarkingTimeSlicePropertyType)
		}
		stage.OnAfterRunwayMarkingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaymarkingtimeslicepropertytype, frontRunwayMarkingTimeSlicePropertyType)
	}
}

func (runwaymarkingtimeslicepropertytype *RunwayMarkingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayMarkingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayMarkingTimeSlicePropertyType *RunwayMarkingTimeSlicePropertyType
		if front != nil {
			frontRunwayMarkingTimeSlicePropertyType, _ = front.(*RunwayMarkingTimeSlicePropertyType)
		}
		stage.OnAfterRunwayMarkingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwaymarkingtimeslicepropertytype, frontRunwayMarkingTimeSlicePropertyType)
	}
}

func (runwaymarkingtimeslicetype *RunwayMarkingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayMarkingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayMarkingTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwaymarkingtimeslicetype)
	}
}

func (runwaymarkingtimeslicetype *RunwayMarkingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayMarkingTimeSliceTypeUpdateCallback != nil {
		var frontRunwayMarkingTimeSliceType *RunwayMarkingTimeSliceType
		if front != nil {
			frontRunwayMarkingTimeSliceType, _ = front.(*RunwayMarkingTimeSliceType)
		}
		stage.OnAfterRunwayMarkingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwaymarkingtimeslicetype, frontRunwayMarkingTimeSliceType)
	}
}

func (runwaymarkingtimeslicetype *RunwayMarkingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayMarkingTimeSliceTypeDeleteCallback != nil {
		var frontRunwayMarkingTimeSliceType *RunwayMarkingTimeSliceType
		if front != nil {
			frontRunwayMarkingTimeSliceType, _ = front.(*RunwayMarkingTimeSliceType)
		}
		stage.OnAfterRunwayMarkingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwaymarkingtimeslicetype, frontRunwayMarkingTimeSliceType)
	}
}

func (runwaymarkingtype *RunwayMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayMarkingTypeCreateCallback != nil {
		stage.OnAfterRunwayMarkingTypeCreateCallback.OnAfterCreate(stage, runwaymarkingtype)
	}
}

func (runwaymarkingtype *RunwayMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayMarkingTypeUpdateCallback != nil {
		var frontRunwayMarkingType *RunwayMarkingType
		if front != nil {
			frontRunwayMarkingType, _ = front.(*RunwayMarkingType)
		}
		stage.OnAfterRunwayMarkingTypeUpdateCallback.OnAfterUpdate(stage, runwaymarkingtype, frontRunwayMarkingType)
	}
}

func (runwaymarkingtype *RunwayMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayMarkingTypeDeleteCallback != nil {
		var frontRunwayMarkingType *RunwayMarkingType
		if front != nil {
			frontRunwayMarkingType, _ = front.(*RunwayMarkingType)
		}
		stage.OnAfterRunwayMarkingTypeDeleteCallback.OnAfterDelete(stage, runwaymarkingtype, frontRunwayMarkingType)
	}
}

func (runwaypropertytype *RunwayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayPropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayPropertyTypeCreateCallback.OnAfterCreate(stage, runwaypropertytype)
	}
}

func (runwaypropertytype *RunwayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayPropertyTypeUpdateCallback != nil {
		var frontRunwayPropertyType *RunwayPropertyType
		if front != nil {
			frontRunwayPropertyType, _ = front.(*RunwayPropertyType)
		}
		stage.OnAfterRunwayPropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaypropertytype, frontRunwayPropertyType)
	}
}

func (runwaypropertytype *RunwayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayPropertyTypeDeleteCallback != nil {
		var frontRunwayPropertyType *RunwayPropertyType
		if front != nil {
			frontRunwayPropertyType, _ = front.(*RunwayPropertyType)
		}
		stage.OnAfterRunwayPropertyTypeDeleteCallback.OnAfterDelete(stage, runwaypropertytype, frontRunwayPropertyType)
	}
}

func (runwayprotectarealightsystemtimeslicepropertytype *RunwayProtectAreaLightSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayProtectAreaLightSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayProtectAreaLightSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwayprotectarealightsystemtimeslicepropertytype)
	}
}

func (runwayprotectarealightsystemtimeslicepropertytype *RunwayProtectAreaLightSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaLightSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayProtectAreaLightSystemTimeSlicePropertyType *RunwayProtectAreaLightSystemTimeSlicePropertyType
		if front != nil {
			frontRunwayProtectAreaLightSystemTimeSlicePropertyType, _ = front.(*RunwayProtectAreaLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterRunwayProtectAreaLightSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwayprotectarealightsystemtimeslicepropertytype, frontRunwayProtectAreaLightSystemTimeSlicePropertyType)
	}
}

func (runwayprotectarealightsystemtimeslicepropertytype *RunwayProtectAreaLightSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaLightSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayProtectAreaLightSystemTimeSlicePropertyType *RunwayProtectAreaLightSystemTimeSlicePropertyType
		if front != nil {
			frontRunwayProtectAreaLightSystemTimeSlicePropertyType, _ = front.(*RunwayProtectAreaLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterRunwayProtectAreaLightSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwayprotectarealightsystemtimeslicepropertytype, frontRunwayProtectAreaLightSystemTimeSlicePropertyType)
	}
}

func (runwayprotectarealightsystemtimeslicetype *RunwayProtectAreaLightSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayProtectAreaLightSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayProtectAreaLightSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwayprotectarealightsystemtimeslicetype)
	}
}

func (runwayprotectarealightsystemtimeslicetype *RunwayProtectAreaLightSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaLightSystemTimeSliceTypeUpdateCallback != nil {
		var frontRunwayProtectAreaLightSystemTimeSliceType *RunwayProtectAreaLightSystemTimeSliceType
		if front != nil {
			frontRunwayProtectAreaLightSystemTimeSliceType, _ = front.(*RunwayProtectAreaLightSystemTimeSliceType)
		}
		stage.OnAfterRunwayProtectAreaLightSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwayprotectarealightsystemtimeslicetype, frontRunwayProtectAreaLightSystemTimeSliceType)
	}
}

func (runwayprotectarealightsystemtimeslicetype *RunwayProtectAreaLightSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaLightSystemTimeSliceTypeDeleteCallback != nil {
		var frontRunwayProtectAreaLightSystemTimeSliceType *RunwayProtectAreaLightSystemTimeSliceType
		if front != nil {
			frontRunwayProtectAreaLightSystemTimeSliceType, _ = front.(*RunwayProtectAreaLightSystemTimeSliceType)
		}
		stage.OnAfterRunwayProtectAreaLightSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwayprotectarealightsystemtimeslicetype, frontRunwayProtectAreaLightSystemTimeSliceType)
	}
}

func (runwayprotectarealightsystemtype *RunwayProtectAreaLightSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayProtectAreaLightSystemTypeCreateCallback != nil {
		stage.OnAfterRunwayProtectAreaLightSystemTypeCreateCallback.OnAfterCreate(stage, runwayprotectarealightsystemtype)
	}
}

func (runwayprotectarealightsystemtype *RunwayProtectAreaLightSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaLightSystemTypeUpdateCallback != nil {
		var frontRunwayProtectAreaLightSystemType *RunwayProtectAreaLightSystemType
		if front != nil {
			frontRunwayProtectAreaLightSystemType, _ = front.(*RunwayProtectAreaLightSystemType)
		}
		stage.OnAfterRunwayProtectAreaLightSystemTypeUpdateCallback.OnAfterUpdate(stage, runwayprotectarealightsystemtype, frontRunwayProtectAreaLightSystemType)
	}
}

func (runwayprotectarealightsystemtype *RunwayProtectAreaLightSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaLightSystemTypeDeleteCallback != nil {
		var frontRunwayProtectAreaLightSystemType *RunwayProtectAreaLightSystemType
		if front != nil {
			frontRunwayProtectAreaLightSystemType, _ = front.(*RunwayProtectAreaLightSystemType)
		}
		stage.OnAfterRunwayProtectAreaLightSystemTypeDeleteCallback.OnAfterDelete(stage, runwayprotectarealightsystemtype, frontRunwayProtectAreaLightSystemType)
	}
}

func (runwayprotectareapropertytype *RunwayProtectAreaPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayProtectAreaPropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayProtectAreaPropertyTypeCreateCallback.OnAfterCreate(stage, runwayprotectareapropertytype)
	}
}

func (runwayprotectareapropertytype *RunwayProtectAreaPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaPropertyTypeUpdateCallback != nil {
		var frontRunwayProtectAreaPropertyType *RunwayProtectAreaPropertyType
		if front != nil {
			frontRunwayProtectAreaPropertyType, _ = front.(*RunwayProtectAreaPropertyType)
		}
		stage.OnAfterRunwayProtectAreaPropertyTypeUpdateCallback.OnAfterUpdate(stage, runwayprotectareapropertytype, frontRunwayProtectAreaPropertyType)
	}
}

func (runwayprotectareapropertytype *RunwayProtectAreaPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaPropertyTypeDeleteCallback != nil {
		var frontRunwayProtectAreaPropertyType *RunwayProtectAreaPropertyType
		if front != nil {
			frontRunwayProtectAreaPropertyType, _ = front.(*RunwayProtectAreaPropertyType)
		}
		stage.OnAfterRunwayProtectAreaPropertyTypeDeleteCallback.OnAfterDelete(stage, runwayprotectareapropertytype, frontRunwayProtectAreaPropertyType)
	}
}

func (runwayprotectareatimeslicepropertytype *RunwayProtectAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayProtectAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayProtectAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwayprotectareatimeslicepropertytype)
	}
}

func (runwayprotectareatimeslicepropertytype *RunwayProtectAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayProtectAreaTimeSlicePropertyType *RunwayProtectAreaTimeSlicePropertyType
		if front != nil {
			frontRunwayProtectAreaTimeSlicePropertyType, _ = front.(*RunwayProtectAreaTimeSlicePropertyType)
		}
		stage.OnAfterRunwayProtectAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwayprotectareatimeslicepropertytype, frontRunwayProtectAreaTimeSlicePropertyType)
	}
}

func (runwayprotectareatimeslicepropertytype *RunwayProtectAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayProtectAreaTimeSlicePropertyType *RunwayProtectAreaTimeSlicePropertyType
		if front != nil {
			frontRunwayProtectAreaTimeSlicePropertyType, _ = front.(*RunwayProtectAreaTimeSlicePropertyType)
		}
		stage.OnAfterRunwayProtectAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwayprotectareatimeslicepropertytype, frontRunwayProtectAreaTimeSlicePropertyType)
	}
}

func (runwayprotectareatimeslicetype *RunwayProtectAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayProtectAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayProtectAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwayprotectareatimeslicetype)
	}
}

func (runwayprotectareatimeslicetype *RunwayProtectAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaTimeSliceTypeUpdateCallback != nil {
		var frontRunwayProtectAreaTimeSliceType *RunwayProtectAreaTimeSliceType
		if front != nil {
			frontRunwayProtectAreaTimeSliceType, _ = front.(*RunwayProtectAreaTimeSliceType)
		}
		stage.OnAfterRunwayProtectAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwayprotectareatimeslicetype, frontRunwayProtectAreaTimeSliceType)
	}
}

func (runwayprotectareatimeslicetype *RunwayProtectAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaTimeSliceTypeDeleteCallback != nil {
		var frontRunwayProtectAreaTimeSliceType *RunwayProtectAreaTimeSliceType
		if front != nil {
			frontRunwayProtectAreaTimeSliceType, _ = front.(*RunwayProtectAreaTimeSliceType)
		}
		stage.OnAfterRunwayProtectAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwayprotectareatimeslicetype, frontRunwayProtectAreaTimeSliceType)
	}
}

func (runwayprotectareatype *RunwayProtectAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayProtectAreaTypeCreateCallback != nil {
		stage.OnAfterRunwayProtectAreaTypeCreateCallback.OnAfterCreate(stage, runwayprotectareatype)
	}
}

func (runwayprotectareatype *RunwayProtectAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaTypeUpdateCallback != nil {
		var frontRunwayProtectAreaType *RunwayProtectAreaType
		if front != nil {
			frontRunwayProtectAreaType, _ = front.(*RunwayProtectAreaType)
		}
		stage.OnAfterRunwayProtectAreaTypeUpdateCallback.OnAfterUpdate(stage, runwayprotectareatype, frontRunwayProtectAreaType)
	}
}

func (runwayprotectareatype *RunwayProtectAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayProtectAreaTypeDeleteCallback != nil {
		var frontRunwayProtectAreaType *RunwayProtectAreaType
		if front != nil {
			frontRunwayProtectAreaType, _ = front.(*RunwayProtectAreaType)
		}
		stage.OnAfterRunwayProtectAreaTypeDeleteCallback.OnAfterDelete(stage, runwayprotectareatype, frontRunwayProtectAreaType)
	}
}

func (runwaysectioncontaminationpropertytype *RunwaySectionContaminationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwaySectionContaminationPropertyTypeCreateCallback != nil {
		stage.OnAfterRunwaySectionContaminationPropertyTypeCreateCallback.OnAfterCreate(stage, runwaysectioncontaminationpropertytype)
	}
}

func (runwaysectioncontaminationpropertytype *RunwaySectionContaminationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwaySectionContaminationPropertyTypeUpdateCallback != nil {
		var frontRunwaySectionContaminationPropertyType *RunwaySectionContaminationPropertyType
		if front != nil {
			frontRunwaySectionContaminationPropertyType, _ = front.(*RunwaySectionContaminationPropertyType)
		}
		stage.OnAfterRunwaySectionContaminationPropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaysectioncontaminationpropertytype, frontRunwaySectionContaminationPropertyType)
	}
}

func (runwaysectioncontaminationpropertytype *RunwaySectionContaminationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwaySectionContaminationPropertyTypeDeleteCallback != nil {
		var frontRunwaySectionContaminationPropertyType *RunwaySectionContaminationPropertyType
		if front != nil {
			frontRunwaySectionContaminationPropertyType, _ = front.(*RunwaySectionContaminationPropertyType)
		}
		stage.OnAfterRunwaySectionContaminationPropertyTypeDeleteCallback.OnAfterDelete(stage, runwaysectioncontaminationpropertytype, frontRunwaySectionContaminationPropertyType)
	}
}

func (runwaysectioncontaminationtype *RunwaySectionContaminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwaySectionContaminationTypeCreateCallback != nil {
		stage.OnAfterRunwaySectionContaminationTypeCreateCallback.OnAfterCreate(stage, runwaysectioncontaminationtype)
	}
}

func (runwaysectioncontaminationtype *RunwaySectionContaminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwaySectionContaminationTypeUpdateCallback != nil {
		var frontRunwaySectionContaminationType *RunwaySectionContaminationType
		if front != nil {
			frontRunwaySectionContaminationType, _ = front.(*RunwaySectionContaminationType)
		}
		stage.OnAfterRunwaySectionContaminationTypeUpdateCallback.OnAfterUpdate(stage, runwaysectioncontaminationtype, frontRunwaySectionContaminationType)
	}
}

func (runwaysectioncontaminationtype *RunwaySectionContaminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwaySectionContaminationTypeDeleteCallback != nil {
		var frontRunwaySectionContaminationType *RunwaySectionContaminationType
		if front != nil {
			frontRunwaySectionContaminationType, _ = front.(*RunwaySectionContaminationType)
		}
		stage.OnAfterRunwaySectionContaminationTypeDeleteCallback.OnAfterDelete(stage, runwaysectioncontaminationtype, frontRunwaySectionContaminationType)
	}
}

func (runwaytimeslicepropertytype *RunwayTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwaytimeslicepropertytype)
	}
}

func (runwaytimeslicepropertytype *RunwayTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayTimeSlicePropertyType *RunwayTimeSlicePropertyType
		if front != nil {
			frontRunwayTimeSlicePropertyType, _ = front.(*RunwayTimeSlicePropertyType)
		}
		stage.OnAfterRunwayTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaytimeslicepropertytype, frontRunwayTimeSlicePropertyType)
	}
}

func (runwaytimeslicepropertytype *RunwayTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayTimeSlicePropertyType *RunwayTimeSlicePropertyType
		if front != nil {
			frontRunwayTimeSlicePropertyType, _ = front.(*RunwayTimeSlicePropertyType)
		}
		stage.OnAfterRunwayTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwaytimeslicepropertytype, frontRunwayTimeSlicePropertyType)
	}
}

func (runwaytimeslicetype *RunwayTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwaytimeslicetype)
	}
}

func (runwaytimeslicetype *RunwayTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayTimeSliceTypeUpdateCallback != nil {
		var frontRunwayTimeSliceType *RunwayTimeSliceType
		if front != nil {
			frontRunwayTimeSliceType, _ = front.(*RunwayTimeSliceType)
		}
		stage.OnAfterRunwayTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwaytimeslicetype, frontRunwayTimeSliceType)
	}
}

func (runwaytimeslicetype *RunwayTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayTimeSliceTypeDeleteCallback != nil {
		var frontRunwayTimeSliceType *RunwayTimeSliceType
		if front != nil {
			frontRunwayTimeSliceType, _ = front.(*RunwayTimeSliceType)
		}
		stage.OnAfterRunwayTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwaytimeslicetype, frontRunwayTimeSliceType)
	}
}

func (runwaytype *RunwayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayTypeCreateCallback != nil {
		stage.OnAfterRunwayTypeCreateCallback.OnAfterCreate(stage, runwaytype)
	}
}

func (runwaytype *RunwayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayTypeUpdateCallback != nil {
		var frontRunwayType *RunwayType
		if front != nil {
			frontRunwayType, _ = front.(*RunwayType)
		}
		stage.OnAfterRunwayTypeUpdateCallback.OnAfterUpdate(stage, runwaytype, frontRunwayType)
	}
}

func (runwaytype *RunwayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayTypeDeleteCallback != nil {
		var frontRunwayType *RunwayType
		if front != nil {
			frontRunwayType, _ = front.(*RunwayType)
		}
		stage.OnAfterRunwayTypeDeleteCallback.OnAfterDelete(stage, runwaytype, frontRunwayType)
	}
}

func (runwayvisualrangetimeslicepropertytype *RunwayVisualRangeTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayVisualRangeTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayVisualRangeTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, runwayvisualrangetimeslicepropertytype)
	}
}

func (runwayvisualrangetimeslicepropertytype *RunwayVisualRangeTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayVisualRangeTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRunwayVisualRangeTimeSlicePropertyType *RunwayVisualRangeTimeSlicePropertyType
		if front != nil {
			frontRunwayVisualRangeTimeSlicePropertyType, _ = front.(*RunwayVisualRangeTimeSlicePropertyType)
		}
		stage.OnAfterRunwayVisualRangeTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, runwayvisualrangetimeslicepropertytype, frontRunwayVisualRangeTimeSlicePropertyType)
	}
}

func (runwayvisualrangetimeslicepropertytype *RunwayVisualRangeTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayVisualRangeTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRunwayVisualRangeTimeSlicePropertyType *RunwayVisualRangeTimeSlicePropertyType
		if front != nil {
			frontRunwayVisualRangeTimeSlicePropertyType, _ = front.(*RunwayVisualRangeTimeSlicePropertyType)
		}
		stage.OnAfterRunwayVisualRangeTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, runwayvisualrangetimeslicepropertytype, frontRunwayVisualRangeTimeSlicePropertyType)
	}
}

func (runwayvisualrangetimeslicetype *RunwayVisualRangeTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayVisualRangeTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRunwayVisualRangeTimeSliceTypeCreateCallback.OnAfterCreate(stage, runwayvisualrangetimeslicetype)
	}
}

func (runwayvisualrangetimeslicetype *RunwayVisualRangeTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayVisualRangeTimeSliceTypeUpdateCallback != nil {
		var frontRunwayVisualRangeTimeSliceType *RunwayVisualRangeTimeSliceType
		if front != nil {
			frontRunwayVisualRangeTimeSliceType, _ = front.(*RunwayVisualRangeTimeSliceType)
		}
		stage.OnAfterRunwayVisualRangeTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, runwayvisualrangetimeslicetype, frontRunwayVisualRangeTimeSliceType)
	}
}

func (runwayvisualrangetimeslicetype *RunwayVisualRangeTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayVisualRangeTimeSliceTypeDeleteCallback != nil {
		var frontRunwayVisualRangeTimeSliceType *RunwayVisualRangeTimeSliceType
		if front != nil {
			frontRunwayVisualRangeTimeSliceType, _ = front.(*RunwayVisualRangeTimeSliceType)
		}
		stage.OnAfterRunwayVisualRangeTimeSliceTypeDeleteCallback.OnAfterDelete(stage, runwayvisualrangetimeslicetype, frontRunwayVisualRangeTimeSliceType)
	}
}

func (runwayvisualrangetype *RunwayVisualRangeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayVisualRangeTypeCreateCallback != nil {
		stage.OnAfterRunwayVisualRangeTypeCreateCallback.OnAfterCreate(stage, runwayvisualrangetype)
	}
}

func (runwayvisualrangetype *RunwayVisualRangeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayVisualRangeTypeUpdateCallback != nil {
		var frontRunwayVisualRangeType *RunwayVisualRangeType
		if front != nil {
			frontRunwayVisualRangeType, _ = front.(*RunwayVisualRangeType)
		}
		stage.OnAfterRunwayVisualRangeTypeUpdateCallback.OnAfterUpdate(stage, runwayvisualrangetype, frontRunwayVisualRangeType)
	}
}

func (runwayvisualrangetype *RunwayVisualRangeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayVisualRangeTypeDeleteCallback != nil {
		var frontRunwayVisualRangeType *RunwayVisualRangeType
		if front != nil {
			frontRunwayVisualRangeType, _ = front.(*RunwayVisualRangeType)
		}
		stage.OnAfterRunwayVisualRangeTypeDeleteCallback.OnAfterDelete(stage, runwayvisualrangetype, frontRunwayVisualRangeType)
	}
}

func (seaplanelandingareapropertytype *SeaplaneLandingAreaPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeaplaneLandingAreaPropertyTypeCreateCallback != nil {
		stage.OnAfterSeaplaneLandingAreaPropertyTypeCreateCallback.OnAfterCreate(stage, seaplanelandingareapropertytype)
	}
}

func (seaplanelandingareapropertytype *SeaplaneLandingAreaPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneLandingAreaPropertyTypeUpdateCallback != nil {
		var frontSeaplaneLandingAreaPropertyType *SeaplaneLandingAreaPropertyType
		if front != nil {
			frontSeaplaneLandingAreaPropertyType, _ = front.(*SeaplaneLandingAreaPropertyType)
		}
		stage.OnAfterSeaplaneLandingAreaPropertyTypeUpdateCallback.OnAfterUpdate(stage, seaplanelandingareapropertytype, frontSeaplaneLandingAreaPropertyType)
	}
}

func (seaplanelandingareapropertytype *SeaplaneLandingAreaPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneLandingAreaPropertyTypeDeleteCallback != nil {
		var frontSeaplaneLandingAreaPropertyType *SeaplaneLandingAreaPropertyType
		if front != nil {
			frontSeaplaneLandingAreaPropertyType, _ = front.(*SeaplaneLandingAreaPropertyType)
		}
		stage.OnAfterSeaplaneLandingAreaPropertyTypeDeleteCallback.OnAfterDelete(stage, seaplanelandingareapropertytype, frontSeaplaneLandingAreaPropertyType)
	}
}

func (seaplanelandingareatimeslicepropertytype *SeaplaneLandingAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeaplaneLandingAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSeaplaneLandingAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, seaplanelandingareatimeslicepropertytype)
	}
}

func (seaplanelandingareatimeslicepropertytype *SeaplaneLandingAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneLandingAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSeaplaneLandingAreaTimeSlicePropertyType *SeaplaneLandingAreaTimeSlicePropertyType
		if front != nil {
			frontSeaplaneLandingAreaTimeSlicePropertyType, _ = front.(*SeaplaneLandingAreaTimeSlicePropertyType)
		}
		stage.OnAfterSeaplaneLandingAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, seaplanelandingareatimeslicepropertytype, frontSeaplaneLandingAreaTimeSlicePropertyType)
	}
}

func (seaplanelandingareatimeslicepropertytype *SeaplaneLandingAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneLandingAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSeaplaneLandingAreaTimeSlicePropertyType *SeaplaneLandingAreaTimeSlicePropertyType
		if front != nil {
			frontSeaplaneLandingAreaTimeSlicePropertyType, _ = front.(*SeaplaneLandingAreaTimeSlicePropertyType)
		}
		stage.OnAfterSeaplaneLandingAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, seaplanelandingareatimeslicepropertytype, frontSeaplaneLandingAreaTimeSlicePropertyType)
	}
}

func (seaplanelandingareatimeslicetype *SeaplaneLandingAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeaplaneLandingAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSeaplaneLandingAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, seaplanelandingareatimeslicetype)
	}
}

func (seaplanelandingareatimeslicetype *SeaplaneLandingAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneLandingAreaTimeSliceTypeUpdateCallback != nil {
		var frontSeaplaneLandingAreaTimeSliceType *SeaplaneLandingAreaTimeSliceType
		if front != nil {
			frontSeaplaneLandingAreaTimeSliceType, _ = front.(*SeaplaneLandingAreaTimeSliceType)
		}
		stage.OnAfterSeaplaneLandingAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, seaplanelandingareatimeslicetype, frontSeaplaneLandingAreaTimeSliceType)
	}
}

func (seaplanelandingareatimeslicetype *SeaplaneLandingAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneLandingAreaTimeSliceTypeDeleteCallback != nil {
		var frontSeaplaneLandingAreaTimeSliceType *SeaplaneLandingAreaTimeSliceType
		if front != nil {
			frontSeaplaneLandingAreaTimeSliceType, _ = front.(*SeaplaneLandingAreaTimeSliceType)
		}
		stage.OnAfterSeaplaneLandingAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, seaplanelandingareatimeslicetype, frontSeaplaneLandingAreaTimeSliceType)
	}
}

func (seaplanelandingareatype *SeaplaneLandingAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeaplaneLandingAreaTypeCreateCallback != nil {
		stage.OnAfterSeaplaneLandingAreaTypeCreateCallback.OnAfterCreate(stage, seaplanelandingareatype)
	}
}

func (seaplanelandingareatype *SeaplaneLandingAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneLandingAreaTypeUpdateCallback != nil {
		var frontSeaplaneLandingAreaType *SeaplaneLandingAreaType
		if front != nil {
			frontSeaplaneLandingAreaType, _ = front.(*SeaplaneLandingAreaType)
		}
		stage.OnAfterSeaplaneLandingAreaTypeUpdateCallback.OnAfterUpdate(stage, seaplanelandingareatype, frontSeaplaneLandingAreaType)
	}
}

func (seaplanelandingareatype *SeaplaneLandingAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneLandingAreaTypeDeleteCallback != nil {
		var frontSeaplaneLandingAreaType *SeaplaneLandingAreaType
		if front != nil {
			frontSeaplaneLandingAreaType, _ = front.(*SeaplaneLandingAreaType)
		}
		stage.OnAfterSeaplaneLandingAreaTypeDeleteCallback.OnAfterDelete(stage, seaplanelandingareatype, frontSeaplaneLandingAreaType)
	}
}

func (seaplanerampsitepropertytype *SeaplaneRampSitePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeaplaneRampSitePropertyTypeCreateCallback != nil {
		stage.OnAfterSeaplaneRampSitePropertyTypeCreateCallback.OnAfterCreate(stage, seaplanerampsitepropertytype)
	}
}

func (seaplanerampsitepropertytype *SeaplaneRampSitePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneRampSitePropertyTypeUpdateCallback != nil {
		var frontSeaplaneRampSitePropertyType *SeaplaneRampSitePropertyType
		if front != nil {
			frontSeaplaneRampSitePropertyType, _ = front.(*SeaplaneRampSitePropertyType)
		}
		stage.OnAfterSeaplaneRampSitePropertyTypeUpdateCallback.OnAfterUpdate(stage, seaplanerampsitepropertytype, frontSeaplaneRampSitePropertyType)
	}
}

func (seaplanerampsitepropertytype *SeaplaneRampSitePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneRampSitePropertyTypeDeleteCallback != nil {
		var frontSeaplaneRampSitePropertyType *SeaplaneRampSitePropertyType
		if front != nil {
			frontSeaplaneRampSitePropertyType, _ = front.(*SeaplaneRampSitePropertyType)
		}
		stage.OnAfterSeaplaneRampSitePropertyTypeDeleteCallback.OnAfterDelete(stage, seaplanerampsitepropertytype, frontSeaplaneRampSitePropertyType)
	}
}

func (seaplanerampsitetimeslicepropertytype *SeaplaneRampSiteTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeaplaneRampSiteTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSeaplaneRampSiteTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, seaplanerampsitetimeslicepropertytype)
	}
}

func (seaplanerampsitetimeslicepropertytype *SeaplaneRampSiteTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneRampSiteTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSeaplaneRampSiteTimeSlicePropertyType *SeaplaneRampSiteTimeSlicePropertyType
		if front != nil {
			frontSeaplaneRampSiteTimeSlicePropertyType, _ = front.(*SeaplaneRampSiteTimeSlicePropertyType)
		}
		stage.OnAfterSeaplaneRampSiteTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, seaplanerampsitetimeslicepropertytype, frontSeaplaneRampSiteTimeSlicePropertyType)
	}
}

func (seaplanerampsitetimeslicepropertytype *SeaplaneRampSiteTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneRampSiteTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSeaplaneRampSiteTimeSlicePropertyType *SeaplaneRampSiteTimeSlicePropertyType
		if front != nil {
			frontSeaplaneRampSiteTimeSlicePropertyType, _ = front.(*SeaplaneRampSiteTimeSlicePropertyType)
		}
		stage.OnAfterSeaplaneRampSiteTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, seaplanerampsitetimeslicepropertytype, frontSeaplaneRampSiteTimeSlicePropertyType)
	}
}

func (seaplanerampsitetimeslicetype *SeaplaneRampSiteTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeaplaneRampSiteTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSeaplaneRampSiteTimeSliceTypeCreateCallback.OnAfterCreate(stage, seaplanerampsitetimeslicetype)
	}
}

func (seaplanerampsitetimeslicetype *SeaplaneRampSiteTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneRampSiteTimeSliceTypeUpdateCallback != nil {
		var frontSeaplaneRampSiteTimeSliceType *SeaplaneRampSiteTimeSliceType
		if front != nil {
			frontSeaplaneRampSiteTimeSliceType, _ = front.(*SeaplaneRampSiteTimeSliceType)
		}
		stage.OnAfterSeaplaneRampSiteTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, seaplanerampsitetimeslicetype, frontSeaplaneRampSiteTimeSliceType)
	}
}

func (seaplanerampsitetimeslicetype *SeaplaneRampSiteTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneRampSiteTimeSliceTypeDeleteCallback != nil {
		var frontSeaplaneRampSiteTimeSliceType *SeaplaneRampSiteTimeSliceType
		if front != nil {
			frontSeaplaneRampSiteTimeSliceType, _ = front.(*SeaplaneRampSiteTimeSliceType)
		}
		stage.OnAfterSeaplaneRampSiteTimeSliceTypeDeleteCallback.OnAfterDelete(stage, seaplanerampsitetimeslicetype, frontSeaplaneRampSiteTimeSliceType)
	}
}

func (seaplanerampsitetype *SeaplaneRampSiteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeaplaneRampSiteTypeCreateCallback != nil {
		stage.OnAfterSeaplaneRampSiteTypeCreateCallback.OnAfterCreate(stage, seaplanerampsitetype)
	}
}

func (seaplanerampsitetype *SeaplaneRampSiteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneRampSiteTypeUpdateCallback != nil {
		var frontSeaplaneRampSiteType *SeaplaneRampSiteType
		if front != nil {
			frontSeaplaneRampSiteType, _ = front.(*SeaplaneRampSiteType)
		}
		stage.OnAfterSeaplaneRampSiteTypeUpdateCallback.OnAfterUpdate(stage, seaplanerampsitetype, frontSeaplaneRampSiteType)
	}
}

func (seaplanerampsitetype *SeaplaneRampSiteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeaplaneRampSiteTypeDeleteCallback != nil {
		var frontSeaplaneRampSiteType *SeaplaneRampSiteType
		if front != nil {
			frontSeaplaneRampSiteType, _ = front.(*SeaplaneRampSiteType)
		}
		stage.OnAfterSeaplaneRampSiteTypeDeleteCallback.OnAfterDelete(stage, seaplanerampsitetype, frontSeaplaneRampSiteType)
	}
}

func (standmarkingtimeslicepropertytype *StandMarkingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandMarkingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterStandMarkingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, standmarkingtimeslicepropertytype)
	}
}

func (standmarkingtimeslicepropertytype *StandMarkingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandMarkingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontStandMarkingTimeSlicePropertyType *StandMarkingTimeSlicePropertyType
		if front != nil {
			frontStandMarkingTimeSlicePropertyType, _ = front.(*StandMarkingTimeSlicePropertyType)
		}
		stage.OnAfterStandMarkingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, standmarkingtimeslicepropertytype, frontStandMarkingTimeSlicePropertyType)
	}
}

func (standmarkingtimeslicepropertytype *StandMarkingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandMarkingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontStandMarkingTimeSlicePropertyType *StandMarkingTimeSlicePropertyType
		if front != nil {
			frontStandMarkingTimeSlicePropertyType, _ = front.(*StandMarkingTimeSlicePropertyType)
		}
		stage.OnAfterStandMarkingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, standmarkingtimeslicepropertytype, frontStandMarkingTimeSlicePropertyType)
	}
}

func (standmarkingtimeslicetype *StandMarkingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandMarkingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterStandMarkingTimeSliceTypeCreateCallback.OnAfterCreate(stage, standmarkingtimeslicetype)
	}
}

func (standmarkingtimeslicetype *StandMarkingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandMarkingTimeSliceTypeUpdateCallback != nil {
		var frontStandMarkingTimeSliceType *StandMarkingTimeSliceType
		if front != nil {
			frontStandMarkingTimeSliceType, _ = front.(*StandMarkingTimeSliceType)
		}
		stage.OnAfterStandMarkingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, standmarkingtimeslicetype, frontStandMarkingTimeSliceType)
	}
}

func (standmarkingtimeslicetype *StandMarkingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandMarkingTimeSliceTypeDeleteCallback != nil {
		var frontStandMarkingTimeSliceType *StandMarkingTimeSliceType
		if front != nil {
			frontStandMarkingTimeSliceType, _ = front.(*StandMarkingTimeSliceType)
		}
		stage.OnAfterStandMarkingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, standmarkingtimeslicetype, frontStandMarkingTimeSliceType)
	}
}

func (standmarkingtype *StandMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandMarkingTypeCreateCallback != nil {
		stage.OnAfterStandMarkingTypeCreateCallback.OnAfterCreate(stage, standmarkingtype)
	}
}

func (standmarkingtype *StandMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandMarkingTypeUpdateCallback != nil {
		var frontStandMarkingType *StandMarkingType
		if front != nil {
			frontStandMarkingType, _ = front.(*StandMarkingType)
		}
		stage.OnAfterStandMarkingTypeUpdateCallback.OnAfterUpdate(stage, standmarkingtype, frontStandMarkingType)
	}
}

func (standmarkingtype *StandMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandMarkingTypeDeleteCallback != nil {
		var frontStandMarkingType *StandMarkingType
		if front != nil {
			frontStandMarkingType, _ = front.(*StandMarkingType)
		}
		stage.OnAfterStandMarkingTypeDeleteCallback.OnAfterDelete(stage, standmarkingtype, frontStandMarkingType)
	}
}

func (surfacecharacteristicspropertytype *SurfaceCharacteristicsPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfaceCharacteristicsPropertyTypeCreateCallback != nil {
		stage.OnAfterSurfaceCharacteristicsPropertyTypeCreateCallback.OnAfterCreate(stage, surfacecharacteristicspropertytype)
	}
}

func (surfacecharacteristicspropertytype *SurfaceCharacteristicsPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceCharacteristicsPropertyTypeUpdateCallback != nil {
		var frontSurfaceCharacteristicsPropertyType *SurfaceCharacteristicsPropertyType
		if front != nil {
			frontSurfaceCharacteristicsPropertyType, _ = front.(*SurfaceCharacteristicsPropertyType)
		}
		stage.OnAfterSurfaceCharacteristicsPropertyTypeUpdateCallback.OnAfterUpdate(stage, surfacecharacteristicspropertytype, frontSurfaceCharacteristicsPropertyType)
	}
}

func (surfacecharacteristicspropertytype *SurfaceCharacteristicsPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceCharacteristicsPropertyTypeDeleteCallback != nil {
		var frontSurfaceCharacteristicsPropertyType *SurfaceCharacteristicsPropertyType
		if front != nil {
			frontSurfaceCharacteristicsPropertyType, _ = front.(*SurfaceCharacteristicsPropertyType)
		}
		stage.OnAfterSurfaceCharacteristicsPropertyTypeDeleteCallback.OnAfterDelete(stage, surfacecharacteristicspropertytype, frontSurfaceCharacteristicsPropertyType)
	}
}

func (surfacecharacteristicstype *SurfaceCharacteristicsType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfaceCharacteristicsTypeCreateCallback != nil {
		stage.OnAfterSurfaceCharacteristicsTypeCreateCallback.OnAfterCreate(stage, surfacecharacteristicstype)
	}
}

func (surfacecharacteristicstype *SurfaceCharacteristicsType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceCharacteristicsTypeUpdateCallback != nil {
		var frontSurfaceCharacteristicsType *SurfaceCharacteristicsType
		if front != nil {
			frontSurfaceCharacteristicsType, _ = front.(*SurfaceCharacteristicsType)
		}
		stage.OnAfterSurfaceCharacteristicsTypeUpdateCallback.OnAfterUpdate(stage, surfacecharacteristicstype, frontSurfaceCharacteristicsType)
	}
}

func (surfacecharacteristicstype *SurfaceCharacteristicsType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceCharacteristicsTypeDeleteCallback != nil {
		var frontSurfaceCharacteristicsType *SurfaceCharacteristicsType
		if front != nil {
			frontSurfaceCharacteristicsType, _ = front.(*SurfaceCharacteristicsType)
		}
		stage.OnAfterSurfaceCharacteristicsTypeDeleteCallback.OnAfterDelete(stage, surfacecharacteristicstype, frontSurfaceCharacteristicsType)
	}
}

func (surfacecontaminationlayerpropertytype *SurfaceContaminationLayerPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfaceContaminationLayerPropertyTypeCreateCallback != nil {
		stage.OnAfterSurfaceContaminationLayerPropertyTypeCreateCallback.OnAfterCreate(stage, surfacecontaminationlayerpropertytype)
	}
}

func (surfacecontaminationlayerpropertytype *SurfaceContaminationLayerPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceContaminationLayerPropertyTypeUpdateCallback != nil {
		var frontSurfaceContaminationLayerPropertyType *SurfaceContaminationLayerPropertyType
		if front != nil {
			frontSurfaceContaminationLayerPropertyType, _ = front.(*SurfaceContaminationLayerPropertyType)
		}
		stage.OnAfterSurfaceContaminationLayerPropertyTypeUpdateCallback.OnAfterUpdate(stage, surfacecontaminationlayerpropertytype, frontSurfaceContaminationLayerPropertyType)
	}
}

func (surfacecontaminationlayerpropertytype *SurfaceContaminationLayerPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceContaminationLayerPropertyTypeDeleteCallback != nil {
		var frontSurfaceContaminationLayerPropertyType *SurfaceContaminationLayerPropertyType
		if front != nil {
			frontSurfaceContaminationLayerPropertyType, _ = front.(*SurfaceContaminationLayerPropertyType)
		}
		stage.OnAfterSurfaceContaminationLayerPropertyTypeDeleteCallback.OnAfterDelete(stage, surfacecontaminationlayerpropertytype, frontSurfaceContaminationLayerPropertyType)
	}
}

func (surfacecontaminationlayertype *SurfaceContaminationLayerType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfaceContaminationLayerTypeCreateCallback != nil {
		stage.OnAfterSurfaceContaminationLayerTypeCreateCallback.OnAfterCreate(stage, surfacecontaminationlayertype)
	}
}

func (surfacecontaminationlayertype *SurfaceContaminationLayerType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceContaminationLayerTypeUpdateCallback != nil {
		var frontSurfaceContaminationLayerType *SurfaceContaminationLayerType
		if front != nil {
			frontSurfaceContaminationLayerType, _ = front.(*SurfaceContaminationLayerType)
		}
		stage.OnAfterSurfaceContaminationLayerTypeUpdateCallback.OnAfterUpdate(stage, surfacecontaminationlayertype, frontSurfaceContaminationLayerType)
	}
}

func (surfacecontaminationlayertype *SurfaceContaminationLayerType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceContaminationLayerTypeDeleteCallback != nil {
		var frontSurfaceContaminationLayerType *SurfaceContaminationLayerType
		if front != nil {
			frontSurfaceContaminationLayerType, _ = front.(*SurfaceContaminationLayerType)
		}
		stage.OnAfterSurfaceContaminationLayerTypeDeleteCallback.OnAfterDelete(stage, surfacecontaminationlayertype, frontSurfaceContaminationLayerType)
	}
}

func (surfacetype *SurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfaceTypeCreateCallback != nil {
		stage.OnAfterSurfaceTypeCreateCallback.OnAfterCreate(stage, surfacetype)
	}
}

func (surfacetype *SurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceTypeUpdateCallback != nil {
		var frontSurfaceType *SurfaceType
		if front != nil {
			frontSurfaceType, _ = front.(*SurfaceType)
		}
		stage.OnAfterSurfaceTypeUpdateCallback.OnAfterUpdate(stage, surfacetype, frontSurfaceType)
	}
}

func (surfacetype *SurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceTypeDeleteCallback != nil {
		var frontSurfaceType *SurfaceType
		if front != nil {
			frontSurfaceType, _ = front.(*SurfaceType)
		}
		stage.OnAfterSurfaceTypeDeleteCallback.OnAfterDelete(stage, surfacetype, frontSurfaceType)
	}
}

func (surveycontrolpointtimeslicepropertytype *SurveyControlPointTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurveyControlPointTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSurveyControlPointTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, surveycontrolpointtimeslicepropertytype)
	}
}

func (surveycontrolpointtimeslicepropertytype *SurveyControlPointTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveyControlPointTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSurveyControlPointTimeSlicePropertyType *SurveyControlPointTimeSlicePropertyType
		if front != nil {
			frontSurveyControlPointTimeSlicePropertyType, _ = front.(*SurveyControlPointTimeSlicePropertyType)
		}
		stage.OnAfterSurveyControlPointTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, surveycontrolpointtimeslicepropertytype, frontSurveyControlPointTimeSlicePropertyType)
	}
}

func (surveycontrolpointtimeslicepropertytype *SurveyControlPointTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveyControlPointTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSurveyControlPointTimeSlicePropertyType *SurveyControlPointTimeSlicePropertyType
		if front != nil {
			frontSurveyControlPointTimeSlicePropertyType, _ = front.(*SurveyControlPointTimeSlicePropertyType)
		}
		stage.OnAfterSurveyControlPointTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, surveycontrolpointtimeslicepropertytype, frontSurveyControlPointTimeSlicePropertyType)
	}
}

func (surveycontrolpointtimeslicetype *SurveyControlPointTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurveyControlPointTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSurveyControlPointTimeSliceTypeCreateCallback.OnAfterCreate(stage, surveycontrolpointtimeslicetype)
	}
}

func (surveycontrolpointtimeslicetype *SurveyControlPointTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveyControlPointTimeSliceTypeUpdateCallback != nil {
		var frontSurveyControlPointTimeSliceType *SurveyControlPointTimeSliceType
		if front != nil {
			frontSurveyControlPointTimeSliceType, _ = front.(*SurveyControlPointTimeSliceType)
		}
		stage.OnAfterSurveyControlPointTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, surveycontrolpointtimeslicetype, frontSurveyControlPointTimeSliceType)
	}
}

func (surveycontrolpointtimeslicetype *SurveyControlPointTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveyControlPointTimeSliceTypeDeleteCallback != nil {
		var frontSurveyControlPointTimeSliceType *SurveyControlPointTimeSliceType
		if front != nil {
			frontSurveyControlPointTimeSliceType, _ = front.(*SurveyControlPointTimeSliceType)
		}
		stage.OnAfterSurveyControlPointTimeSliceTypeDeleteCallback.OnAfterDelete(stage, surveycontrolpointtimeslicetype, frontSurveyControlPointTimeSliceType)
	}
}

func (surveycontrolpointtype *SurveyControlPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurveyControlPointTypeCreateCallback != nil {
		stage.OnAfterSurveyControlPointTypeCreateCallback.OnAfterCreate(stage, surveycontrolpointtype)
	}
}

func (surveycontrolpointtype *SurveyControlPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveyControlPointTypeUpdateCallback != nil {
		var frontSurveyControlPointType *SurveyControlPointType
		if front != nil {
			frontSurveyControlPointType, _ = front.(*SurveyControlPointType)
		}
		stage.OnAfterSurveyControlPointTypeUpdateCallback.OnAfterUpdate(stage, surveycontrolpointtype, frontSurveyControlPointType)
	}
}

func (surveycontrolpointtype *SurveyControlPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveyControlPointTypeDeleteCallback != nil {
		var frontSurveyControlPointType *SurveyControlPointType
		if front != nil {
			frontSurveyControlPointType, _ = front.(*SurveyControlPointType)
		}
		stage.OnAfterSurveyControlPointTypeDeleteCallback.OnAfterDelete(stage, surveycontrolpointtype, frontSurveyControlPointType)
	}
}

func (taxiholdingpositionlightsystemtimeslicepropertytype *TaxiHoldingPositionLightSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionLightSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionLightSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositionlightsystemtimeslicepropertytype)
	}
}

func (taxiholdingpositionlightsystemtimeslicepropertytype *TaxiHoldingPositionLightSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionLightSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionLightSystemTimeSlicePropertyType *TaxiHoldingPositionLightSystemTimeSlicePropertyType
		if front != nil {
			frontTaxiHoldingPositionLightSystemTimeSlicePropertyType, _ = front.(*TaxiHoldingPositionLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterTaxiHoldingPositionLightSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositionlightsystemtimeslicepropertytype, frontTaxiHoldingPositionLightSystemTimeSlicePropertyType)
	}
}

func (taxiholdingpositionlightsystemtimeslicepropertytype *TaxiHoldingPositionLightSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionLightSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionLightSystemTimeSlicePropertyType *TaxiHoldingPositionLightSystemTimeSlicePropertyType
		if front != nil {
			frontTaxiHoldingPositionLightSystemTimeSlicePropertyType, _ = front.(*TaxiHoldingPositionLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterTaxiHoldingPositionLightSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositionlightsystemtimeslicepropertytype, frontTaxiHoldingPositionLightSystemTimeSlicePropertyType)
	}
}

func (taxiholdingpositionlightsystemtimeslicetype *TaxiHoldingPositionLightSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionLightSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionLightSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositionlightsystemtimeslicetype)
	}
}

func (taxiholdingpositionlightsystemtimeslicetype *TaxiHoldingPositionLightSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionLightSystemTimeSliceTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionLightSystemTimeSliceType *TaxiHoldingPositionLightSystemTimeSliceType
		if front != nil {
			frontTaxiHoldingPositionLightSystemTimeSliceType, _ = front.(*TaxiHoldingPositionLightSystemTimeSliceType)
		}
		stage.OnAfterTaxiHoldingPositionLightSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositionlightsystemtimeslicetype, frontTaxiHoldingPositionLightSystemTimeSliceType)
	}
}

func (taxiholdingpositionlightsystemtimeslicetype *TaxiHoldingPositionLightSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionLightSystemTimeSliceTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionLightSystemTimeSliceType *TaxiHoldingPositionLightSystemTimeSliceType
		if front != nil {
			frontTaxiHoldingPositionLightSystemTimeSliceType, _ = front.(*TaxiHoldingPositionLightSystemTimeSliceType)
		}
		stage.OnAfterTaxiHoldingPositionLightSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositionlightsystemtimeslicetype, frontTaxiHoldingPositionLightSystemTimeSliceType)
	}
}

func (taxiholdingpositionlightsystemtype *TaxiHoldingPositionLightSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionLightSystemTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionLightSystemTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositionlightsystemtype)
	}
}

func (taxiholdingpositionlightsystemtype *TaxiHoldingPositionLightSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionLightSystemTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionLightSystemType *TaxiHoldingPositionLightSystemType
		if front != nil {
			frontTaxiHoldingPositionLightSystemType, _ = front.(*TaxiHoldingPositionLightSystemType)
		}
		stage.OnAfterTaxiHoldingPositionLightSystemTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositionlightsystemtype, frontTaxiHoldingPositionLightSystemType)
	}
}

func (taxiholdingpositionlightsystemtype *TaxiHoldingPositionLightSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionLightSystemTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionLightSystemType *TaxiHoldingPositionLightSystemType
		if front != nil {
			frontTaxiHoldingPositionLightSystemType, _ = front.(*TaxiHoldingPositionLightSystemType)
		}
		stage.OnAfterTaxiHoldingPositionLightSystemTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositionlightsystemtype, frontTaxiHoldingPositionLightSystemType)
	}
}

func (taxiholdingpositionmarkingtimeslicepropertytype *TaxiHoldingPositionMarkingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionMarkingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionMarkingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositionmarkingtimeslicepropertytype)
	}
}

func (taxiholdingpositionmarkingtimeslicepropertytype *TaxiHoldingPositionMarkingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionMarkingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionMarkingTimeSlicePropertyType *TaxiHoldingPositionMarkingTimeSlicePropertyType
		if front != nil {
			frontTaxiHoldingPositionMarkingTimeSlicePropertyType, _ = front.(*TaxiHoldingPositionMarkingTimeSlicePropertyType)
		}
		stage.OnAfterTaxiHoldingPositionMarkingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositionmarkingtimeslicepropertytype, frontTaxiHoldingPositionMarkingTimeSlicePropertyType)
	}
}

func (taxiholdingpositionmarkingtimeslicepropertytype *TaxiHoldingPositionMarkingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionMarkingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionMarkingTimeSlicePropertyType *TaxiHoldingPositionMarkingTimeSlicePropertyType
		if front != nil {
			frontTaxiHoldingPositionMarkingTimeSlicePropertyType, _ = front.(*TaxiHoldingPositionMarkingTimeSlicePropertyType)
		}
		stage.OnAfterTaxiHoldingPositionMarkingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositionmarkingtimeslicepropertytype, frontTaxiHoldingPositionMarkingTimeSlicePropertyType)
	}
}

func (taxiholdingpositionmarkingtimeslicetype *TaxiHoldingPositionMarkingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionMarkingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionMarkingTimeSliceTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositionmarkingtimeslicetype)
	}
}

func (taxiholdingpositionmarkingtimeslicetype *TaxiHoldingPositionMarkingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionMarkingTimeSliceTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionMarkingTimeSliceType *TaxiHoldingPositionMarkingTimeSliceType
		if front != nil {
			frontTaxiHoldingPositionMarkingTimeSliceType, _ = front.(*TaxiHoldingPositionMarkingTimeSliceType)
		}
		stage.OnAfterTaxiHoldingPositionMarkingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositionmarkingtimeslicetype, frontTaxiHoldingPositionMarkingTimeSliceType)
	}
}

func (taxiholdingpositionmarkingtimeslicetype *TaxiHoldingPositionMarkingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionMarkingTimeSliceTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionMarkingTimeSliceType *TaxiHoldingPositionMarkingTimeSliceType
		if front != nil {
			frontTaxiHoldingPositionMarkingTimeSliceType, _ = front.(*TaxiHoldingPositionMarkingTimeSliceType)
		}
		stage.OnAfterTaxiHoldingPositionMarkingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositionmarkingtimeslicetype, frontTaxiHoldingPositionMarkingTimeSliceType)
	}
}

func (taxiholdingpositionmarkingtype *TaxiHoldingPositionMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionMarkingTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionMarkingTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositionmarkingtype)
	}
}

func (taxiholdingpositionmarkingtype *TaxiHoldingPositionMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionMarkingTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionMarkingType *TaxiHoldingPositionMarkingType
		if front != nil {
			frontTaxiHoldingPositionMarkingType, _ = front.(*TaxiHoldingPositionMarkingType)
		}
		stage.OnAfterTaxiHoldingPositionMarkingTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositionmarkingtype, frontTaxiHoldingPositionMarkingType)
	}
}

func (taxiholdingpositionmarkingtype *TaxiHoldingPositionMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionMarkingTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionMarkingType *TaxiHoldingPositionMarkingType
		if front != nil {
			frontTaxiHoldingPositionMarkingType, _ = front.(*TaxiHoldingPositionMarkingType)
		}
		stage.OnAfterTaxiHoldingPositionMarkingTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositionmarkingtype, frontTaxiHoldingPositionMarkingType)
	}
}

func (taxiholdingpositionpropertytype *TaxiHoldingPositionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionPropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionPropertyTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositionpropertytype)
	}
}

func (taxiholdingpositionpropertytype *TaxiHoldingPositionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionPropertyTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionPropertyType *TaxiHoldingPositionPropertyType
		if front != nil {
			frontTaxiHoldingPositionPropertyType, _ = front.(*TaxiHoldingPositionPropertyType)
		}
		stage.OnAfterTaxiHoldingPositionPropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositionpropertytype, frontTaxiHoldingPositionPropertyType)
	}
}

func (taxiholdingpositionpropertytype *TaxiHoldingPositionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionPropertyTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionPropertyType *TaxiHoldingPositionPropertyType
		if front != nil {
			frontTaxiHoldingPositionPropertyType, _ = front.(*TaxiHoldingPositionPropertyType)
		}
		stage.OnAfterTaxiHoldingPositionPropertyTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositionpropertytype, frontTaxiHoldingPositionPropertyType)
	}
}

func (taxiholdingpositiontimeslicepropertytype *TaxiHoldingPositionTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositiontimeslicepropertytype)
	}
}

func (taxiholdingpositiontimeslicepropertytype *TaxiHoldingPositionTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionTimeSlicePropertyType *TaxiHoldingPositionTimeSlicePropertyType
		if front != nil {
			frontTaxiHoldingPositionTimeSlicePropertyType, _ = front.(*TaxiHoldingPositionTimeSlicePropertyType)
		}
		stage.OnAfterTaxiHoldingPositionTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositiontimeslicepropertytype, frontTaxiHoldingPositionTimeSlicePropertyType)
	}
}

func (taxiholdingpositiontimeslicepropertytype *TaxiHoldingPositionTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionTimeSlicePropertyType *TaxiHoldingPositionTimeSlicePropertyType
		if front != nil {
			frontTaxiHoldingPositionTimeSlicePropertyType, _ = front.(*TaxiHoldingPositionTimeSlicePropertyType)
		}
		stage.OnAfterTaxiHoldingPositionTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositiontimeslicepropertytype, frontTaxiHoldingPositionTimeSlicePropertyType)
	}
}

func (taxiholdingpositiontimeslicetype *TaxiHoldingPositionTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionTimeSliceTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositiontimeslicetype)
	}
}

func (taxiholdingpositiontimeslicetype *TaxiHoldingPositionTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionTimeSliceTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionTimeSliceType *TaxiHoldingPositionTimeSliceType
		if front != nil {
			frontTaxiHoldingPositionTimeSliceType, _ = front.(*TaxiHoldingPositionTimeSliceType)
		}
		stage.OnAfterTaxiHoldingPositionTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositiontimeslicetype, frontTaxiHoldingPositionTimeSliceType)
	}
}

func (taxiholdingpositiontimeslicetype *TaxiHoldingPositionTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionTimeSliceTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionTimeSliceType *TaxiHoldingPositionTimeSliceType
		if front != nil {
			frontTaxiHoldingPositionTimeSliceType, _ = front.(*TaxiHoldingPositionTimeSliceType)
		}
		stage.OnAfterTaxiHoldingPositionTimeSliceTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositiontimeslicetype, frontTaxiHoldingPositionTimeSliceType)
	}
}

func (taxiholdingpositiontype *TaxiHoldingPositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiHoldingPositionTypeCreateCallback != nil {
		stage.OnAfterTaxiHoldingPositionTypeCreateCallback.OnAfterCreate(stage, taxiholdingpositiontype)
	}
}

func (taxiholdingpositiontype *TaxiHoldingPositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionTypeUpdateCallback != nil {
		var frontTaxiHoldingPositionType *TaxiHoldingPositionType
		if front != nil {
			frontTaxiHoldingPositionType, _ = front.(*TaxiHoldingPositionType)
		}
		stage.OnAfterTaxiHoldingPositionTypeUpdateCallback.OnAfterUpdate(stage, taxiholdingpositiontype, frontTaxiHoldingPositionType)
	}
}

func (taxiholdingpositiontype *TaxiHoldingPositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiHoldingPositionTypeDeleteCallback != nil {
		var frontTaxiHoldingPositionType *TaxiHoldingPositionType
		if front != nil {
			frontTaxiHoldingPositionType, _ = front.(*TaxiHoldingPositionType)
		}
		stage.OnAfterTaxiHoldingPositionTypeDeleteCallback.OnAfterDelete(stage, taxiholdingpositiontype, frontTaxiHoldingPositionType)
	}
}

func (taxiwaycontaminationpropertytype *TaxiwayContaminationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayContaminationPropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiwayContaminationPropertyTypeCreateCallback.OnAfterCreate(stage, taxiwaycontaminationpropertytype)
	}
}

func (taxiwaycontaminationpropertytype *TaxiwayContaminationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayContaminationPropertyTypeUpdateCallback != nil {
		var frontTaxiwayContaminationPropertyType *TaxiwayContaminationPropertyType
		if front != nil {
			frontTaxiwayContaminationPropertyType, _ = front.(*TaxiwayContaminationPropertyType)
		}
		stage.OnAfterTaxiwayContaminationPropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiwaycontaminationpropertytype, frontTaxiwayContaminationPropertyType)
	}
}

func (taxiwaycontaminationpropertytype *TaxiwayContaminationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayContaminationPropertyTypeDeleteCallback != nil {
		var frontTaxiwayContaminationPropertyType *TaxiwayContaminationPropertyType
		if front != nil {
			frontTaxiwayContaminationPropertyType, _ = front.(*TaxiwayContaminationPropertyType)
		}
		stage.OnAfterTaxiwayContaminationPropertyTypeDeleteCallback.OnAfterDelete(stage, taxiwaycontaminationpropertytype, frontTaxiwayContaminationPropertyType)
	}
}

func (taxiwaycontaminationtype *TaxiwayContaminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayContaminationTypeCreateCallback != nil {
		stage.OnAfterTaxiwayContaminationTypeCreateCallback.OnAfterCreate(stage, taxiwaycontaminationtype)
	}
}

func (taxiwaycontaminationtype *TaxiwayContaminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayContaminationTypeUpdateCallback != nil {
		var frontTaxiwayContaminationType *TaxiwayContaminationType
		if front != nil {
			frontTaxiwayContaminationType, _ = front.(*TaxiwayContaminationType)
		}
		stage.OnAfterTaxiwayContaminationTypeUpdateCallback.OnAfterUpdate(stage, taxiwaycontaminationtype, frontTaxiwayContaminationType)
	}
}

func (taxiwaycontaminationtype *TaxiwayContaminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayContaminationTypeDeleteCallback != nil {
		var frontTaxiwayContaminationType *TaxiwayContaminationType
		if front != nil {
			frontTaxiwayContaminationType, _ = front.(*TaxiwayContaminationType)
		}
		stage.OnAfterTaxiwayContaminationTypeDeleteCallback.OnAfterDelete(stage, taxiwaycontaminationtype, frontTaxiwayContaminationType)
	}
}

func (taxiwayelementpropertytype *TaxiwayElementPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayElementPropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiwayElementPropertyTypeCreateCallback.OnAfterCreate(stage, taxiwayelementpropertytype)
	}
}

func (taxiwayelementpropertytype *TaxiwayElementPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayElementPropertyTypeUpdateCallback != nil {
		var frontTaxiwayElementPropertyType *TaxiwayElementPropertyType
		if front != nil {
			frontTaxiwayElementPropertyType, _ = front.(*TaxiwayElementPropertyType)
		}
		stage.OnAfterTaxiwayElementPropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiwayelementpropertytype, frontTaxiwayElementPropertyType)
	}
}

func (taxiwayelementpropertytype *TaxiwayElementPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayElementPropertyTypeDeleteCallback != nil {
		var frontTaxiwayElementPropertyType *TaxiwayElementPropertyType
		if front != nil {
			frontTaxiwayElementPropertyType, _ = front.(*TaxiwayElementPropertyType)
		}
		stage.OnAfterTaxiwayElementPropertyTypeDeleteCallback.OnAfterDelete(stage, taxiwayelementpropertytype, frontTaxiwayElementPropertyType)
	}
}

func (taxiwayelementtimeslicepropertytype *TaxiwayElementTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayElementTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiwayElementTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, taxiwayelementtimeslicepropertytype)
	}
}

func (taxiwayelementtimeslicepropertytype *TaxiwayElementTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayElementTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTaxiwayElementTimeSlicePropertyType *TaxiwayElementTimeSlicePropertyType
		if front != nil {
			frontTaxiwayElementTimeSlicePropertyType, _ = front.(*TaxiwayElementTimeSlicePropertyType)
		}
		stage.OnAfterTaxiwayElementTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiwayelementtimeslicepropertytype, frontTaxiwayElementTimeSlicePropertyType)
	}
}

func (taxiwayelementtimeslicepropertytype *TaxiwayElementTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayElementTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTaxiwayElementTimeSlicePropertyType *TaxiwayElementTimeSlicePropertyType
		if front != nil {
			frontTaxiwayElementTimeSlicePropertyType, _ = front.(*TaxiwayElementTimeSlicePropertyType)
		}
		stage.OnAfterTaxiwayElementTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, taxiwayelementtimeslicepropertytype, frontTaxiwayElementTimeSlicePropertyType)
	}
}

func (taxiwayelementtimeslicetype *TaxiwayElementTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayElementTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTaxiwayElementTimeSliceTypeCreateCallback.OnAfterCreate(stage, taxiwayelementtimeslicetype)
	}
}

func (taxiwayelementtimeslicetype *TaxiwayElementTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayElementTimeSliceTypeUpdateCallback != nil {
		var frontTaxiwayElementTimeSliceType *TaxiwayElementTimeSliceType
		if front != nil {
			frontTaxiwayElementTimeSliceType, _ = front.(*TaxiwayElementTimeSliceType)
		}
		stage.OnAfterTaxiwayElementTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, taxiwayelementtimeslicetype, frontTaxiwayElementTimeSliceType)
	}
}

func (taxiwayelementtimeslicetype *TaxiwayElementTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayElementTimeSliceTypeDeleteCallback != nil {
		var frontTaxiwayElementTimeSliceType *TaxiwayElementTimeSliceType
		if front != nil {
			frontTaxiwayElementTimeSliceType, _ = front.(*TaxiwayElementTimeSliceType)
		}
		stage.OnAfterTaxiwayElementTimeSliceTypeDeleteCallback.OnAfterDelete(stage, taxiwayelementtimeslicetype, frontTaxiwayElementTimeSliceType)
	}
}

func (taxiwayelementtype *TaxiwayElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayElementTypeCreateCallback != nil {
		stage.OnAfterTaxiwayElementTypeCreateCallback.OnAfterCreate(stage, taxiwayelementtype)
	}
}

func (taxiwayelementtype *TaxiwayElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayElementTypeUpdateCallback != nil {
		var frontTaxiwayElementType *TaxiwayElementType
		if front != nil {
			frontTaxiwayElementType, _ = front.(*TaxiwayElementType)
		}
		stage.OnAfterTaxiwayElementTypeUpdateCallback.OnAfterUpdate(stage, taxiwayelementtype, frontTaxiwayElementType)
	}
}

func (taxiwayelementtype *TaxiwayElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayElementTypeDeleteCallback != nil {
		var frontTaxiwayElementType *TaxiwayElementType
		if front != nil {
			frontTaxiwayElementType, _ = front.(*TaxiwayElementType)
		}
		stage.OnAfterTaxiwayElementTypeDeleteCallback.OnAfterDelete(stage, taxiwayelementtype, frontTaxiwayElementType)
	}
}

func (taxiwaylightsystemtimeslicepropertytype *TaxiwayLightSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayLightSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiwayLightSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, taxiwaylightsystemtimeslicepropertytype)
	}
}

func (taxiwaylightsystemtimeslicepropertytype *TaxiwayLightSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayLightSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTaxiwayLightSystemTimeSlicePropertyType *TaxiwayLightSystemTimeSlicePropertyType
		if front != nil {
			frontTaxiwayLightSystemTimeSlicePropertyType, _ = front.(*TaxiwayLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterTaxiwayLightSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiwaylightsystemtimeslicepropertytype, frontTaxiwayLightSystemTimeSlicePropertyType)
	}
}

func (taxiwaylightsystemtimeslicepropertytype *TaxiwayLightSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayLightSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTaxiwayLightSystemTimeSlicePropertyType *TaxiwayLightSystemTimeSlicePropertyType
		if front != nil {
			frontTaxiwayLightSystemTimeSlicePropertyType, _ = front.(*TaxiwayLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterTaxiwayLightSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, taxiwaylightsystemtimeslicepropertytype, frontTaxiwayLightSystemTimeSlicePropertyType)
	}
}

func (taxiwaylightsystemtimeslicetype *TaxiwayLightSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayLightSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTaxiwayLightSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, taxiwaylightsystemtimeslicetype)
	}
}

func (taxiwaylightsystemtimeslicetype *TaxiwayLightSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayLightSystemTimeSliceTypeUpdateCallback != nil {
		var frontTaxiwayLightSystemTimeSliceType *TaxiwayLightSystemTimeSliceType
		if front != nil {
			frontTaxiwayLightSystemTimeSliceType, _ = front.(*TaxiwayLightSystemTimeSliceType)
		}
		stage.OnAfterTaxiwayLightSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, taxiwaylightsystemtimeslicetype, frontTaxiwayLightSystemTimeSliceType)
	}
}

func (taxiwaylightsystemtimeslicetype *TaxiwayLightSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayLightSystemTimeSliceTypeDeleteCallback != nil {
		var frontTaxiwayLightSystemTimeSliceType *TaxiwayLightSystemTimeSliceType
		if front != nil {
			frontTaxiwayLightSystemTimeSliceType, _ = front.(*TaxiwayLightSystemTimeSliceType)
		}
		stage.OnAfterTaxiwayLightSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, taxiwaylightsystemtimeslicetype, frontTaxiwayLightSystemTimeSliceType)
	}
}

func (taxiwaylightsystemtype *TaxiwayLightSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayLightSystemTypeCreateCallback != nil {
		stage.OnAfterTaxiwayLightSystemTypeCreateCallback.OnAfterCreate(stage, taxiwaylightsystemtype)
	}
}

func (taxiwaylightsystemtype *TaxiwayLightSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayLightSystemTypeUpdateCallback != nil {
		var frontTaxiwayLightSystemType *TaxiwayLightSystemType
		if front != nil {
			frontTaxiwayLightSystemType, _ = front.(*TaxiwayLightSystemType)
		}
		stage.OnAfterTaxiwayLightSystemTypeUpdateCallback.OnAfterUpdate(stage, taxiwaylightsystemtype, frontTaxiwayLightSystemType)
	}
}

func (taxiwaylightsystemtype *TaxiwayLightSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayLightSystemTypeDeleteCallback != nil {
		var frontTaxiwayLightSystemType *TaxiwayLightSystemType
		if front != nil {
			frontTaxiwayLightSystemType, _ = front.(*TaxiwayLightSystemType)
		}
		stage.OnAfterTaxiwayLightSystemTypeDeleteCallback.OnAfterDelete(stage, taxiwaylightsystemtype, frontTaxiwayLightSystemType)
	}
}

func (taxiwaymarkingtimeslicepropertytype *TaxiwayMarkingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayMarkingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiwayMarkingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, taxiwaymarkingtimeslicepropertytype)
	}
}

func (taxiwaymarkingtimeslicepropertytype *TaxiwayMarkingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayMarkingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTaxiwayMarkingTimeSlicePropertyType *TaxiwayMarkingTimeSlicePropertyType
		if front != nil {
			frontTaxiwayMarkingTimeSlicePropertyType, _ = front.(*TaxiwayMarkingTimeSlicePropertyType)
		}
		stage.OnAfterTaxiwayMarkingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiwaymarkingtimeslicepropertytype, frontTaxiwayMarkingTimeSlicePropertyType)
	}
}

func (taxiwaymarkingtimeslicepropertytype *TaxiwayMarkingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayMarkingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTaxiwayMarkingTimeSlicePropertyType *TaxiwayMarkingTimeSlicePropertyType
		if front != nil {
			frontTaxiwayMarkingTimeSlicePropertyType, _ = front.(*TaxiwayMarkingTimeSlicePropertyType)
		}
		stage.OnAfterTaxiwayMarkingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, taxiwaymarkingtimeslicepropertytype, frontTaxiwayMarkingTimeSlicePropertyType)
	}
}

func (taxiwaymarkingtimeslicetype *TaxiwayMarkingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayMarkingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTaxiwayMarkingTimeSliceTypeCreateCallback.OnAfterCreate(stage, taxiwaymarkingtimeslicetype)
	}
}

func (taxiwaymarkingtimeslicetype *TaxiwayMarkingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayMarkingTimeSliceTypeUpdateCallback != nil {
		var frontTaxiwayMarkingTimeSliceType *TaxiwayMarkingTimeSliceType
		if front != nil {
			frontTaxiwayMarkingTimeSliceType, _ = front.(*TaxiwayMarkingTimeSliceType)
		}
		stage.OnAfterTaxiwayMarkingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, taxiwaymarkingtimeslicetype, frontTaxiwayMarkingTimeSliceType)
	}
}

func (taxiwaymarkingtimeslicetype *TaxiwayMarkingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayMarkingTimeSliceTypeDeleteCallback != nil {
		var frontTaxiwayMarkingTimeSliceType *TaxiwayMarkingTimeSliceType
		if front != nil {
			frontTaxiwayMarkingTimeSliceType, _ = front.(*TaxiwayMarkingTimeSliceType)
		}
		stage.OnAfterTaxiwayMarkingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, taxiwaymarkingtimeslicetype, frontTaxiwayMarkingTimeSliceType)
	}
}

func (taxiwaymarkingtype *TaxiwayMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayMarkingTypeCreateCallback != nil {
		stage.OnAfterTaxiwayMarkingTypeCreateCallback.OnAfterCreate(stage, taxiwaymarkingtype)
	}
}

func (taxiwaymarkingtype *TaxiwayMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayMarkingTypeUpdateCallback != nil {
		var frontTaxiwayMarkingType *TaxiwayMarkingType
		if front != nil {
			frontTaxiwayMarkingType, _ = front.(*TaxiwayMarkingType)
		}
		stage.OnAfterTaxiwayMarkingTypeUpdateCallback.OnAfterUpdate(stage, taxiwaymarkingtype, frontTaxiwayMarkingType)
	}
}

func (taxiwaymarkingtype *TaxiwayMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayMarkingTypeDeleteCallback != nil {
		var frontTaxiwayMarkingType *TaxiwayMarkingType
		if front != nil {
			frontTaxiwayMarkingType, _ = front.(*TaxiwayMarkingType)
		}
		stage.OnAfterTaxiwayMarkingTypeDeleteCallback.OnAfterDelete(stage, taxiwaymarkingtype, frontTaxiwayMarkingType)
	}
}

func (taxiwaypropertytype *TaxiwayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayPropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiwayPropertyTypeCreateCallback.OnAfterCreate(stage, taxiwaypropertytype)
	}
}

func (taxiwaypropertytype *TaxiwayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayPropertyTypeUpdateCallback != nil {
		var frontTaxiwayPropertyType *TaxiwayPropertyType
		if front != nil {
			frontTaxiwayPropertyType, _ = front.(*TaxiwayPropertyType)
		}
		stage.OnAfterTaxiwayPropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiwaypropertytype, frontTaxiwayPropertyType)
	}
}

func (taxiwaypropertytype *TaxiwayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayPropertyTypeDeleteCallback != nil {
		var frontTaxiwayPropertyType *TaxiwayPropertyType
		if front != nil {
			frontTaxiwayPropertyType, _ = front.(*TaxiwayPropertyType)
		}
		stage.OnAfterTaxiwayPropertyTypeDeleteCallback.OnAfterDelete(stage, taxiwaypropertytype, frontTaxiwayPropertyType)
	}
}

func (taxiwaytimeslicepropertytype *TaxiwayTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTaxiwayTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, taxiwaytimeslicepropertytype)
	}
}

func (taxiwaytimeslicepropertytype *TaxiwayTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTaxiwayTimeSlicePropertyType *TaxiwayTimeSlicePropertyType
		if front != nil {
			frontTaxiwayTimeSlicePropertyType, _ = front.(*TaxiwayTimeSlicePropertyType)
		}
		stage.OnAfterTaxiwayTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, taxiwaytimeslicepropertytype, frontTaxiwayTimeSlicePropertyType)
	}
}

func (taxiwaytimeslicepropertytype *TaxiwayTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTaxiwayTimeSlicePropertyType *TaxiwayTimeSlicePropertyType
		if front != nil {
			frontTaxiwayTimeSlicePropertyType, _ = front.(*TaxiwayTimeSlicePropertyType)
		}
		stage.OnAfterTaxiwayTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, taxiwaytimeslicepropertytype, frontTaxiwayTimeSlicePropertyType)
	}
}

func (taxiwaytimeslicetype *TaxiwayTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTaxiwayTimeSliceTypeCreateCallback.OnAfterCreate(stage, taxiwaytimeslicetype)
	}
}

func (taxiwaytimeslicetype *TaxiwayTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayTimeSliceTypeUpdateCallback != nil {
		var frontTaxiwayTimeSliceType *TaxiwayTimeSliceType
		if front != nil {
			frontTaxiwayTimeSliceType, _ = front.(*TaxiwayTimeSliceType)
		}
		stage.OnAfterTaxiwayTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, taxiwaytimeslicetype, frontTaxiwayTimeSliceType)
	}
}

func (taxiwaytimeslicetype *TaxiwayTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayTimeSliceTypeDeleteCallback != nil {
		var frontTaxiwayTimeSliceType *TaxiwayTimeSliceType
		if front != nil {
			frontTaxiwayTimeSliceType, _ = front.(*TaxiwayTimeSliceType)
		}
		stage.OnAfterTaxiwayTimeSliceTypeDeleteCallback.OnAfterDelete(stage, taxiwaytimeslicetype, frontTaxiwayTimeSliceType)
	}
}

func (taxiwaytype *TaxiwayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaxiwayTypeCreateCallback != nil {
		stage.OnAfterTaxiwayTypeCreateCallback.OnAfterCreate(stage, taxiwaytype)
	}
}

func (taxiwaytype *TaxiwayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayTypeUpdateCallback != nil {
		var frontTaxiwayType *TaxiwayType
		if front != nil {
			frontTaxiwayType, _ = front.(*TaxiwayType)
		}
		stage.OnAfterTaxiwayTypeUpdateCallback.OnAfterUpdate(stage, taxiwaytype, frontTaxiwayType)
	}
}

func (taxiwaytype *TaxiwayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaxiwayTypeDeleteCallback != nil {
		var frontTaxiwayType *TaxiwayType
		if front != nil {
			frontTaxiwayType, _ = front.(*TaxiwayType)
		}
		stage.OnAfterTaxiwayTypeDeleteCallback.OnAfterDelete(stage, taxiwaytype, frontTaxiwayType)
	}
}

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTelephoneContactPropertyTypeCreateCallback != nil {
		stage.OnAfterTelephoneContactPropertyTypeCreateCallback.OnAfterCreate(stage, telephonecontactpropertytype)
	}
}

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTelephoneContactPropertyTypeUpdateCallback != nil {
		var frontTelephoneContactPropertyType *TelephoneContactPropertyType
		if front != nil {
			frontTelephoneContactPropertyType, _ = front.(*TelephoneContactPropertyType)
		}
		stage.OnAfterTelephoneContactPropertyTypeUpdateCallback.OnAfterUpdate(stage, telephonecontactpropertytype, frontTelephoneContactPropertyType)
	}
}

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTelephoneContactPropertyTypeDeleteCallback != nil {
		var frontTelephoneContactPropertyType *TelephoneContactPropertyType
		if front != nil {
			frontTelephoneContactPropertyType, _ = front.(*TelephoneContactPropertyType)
		}
		stage.OnAfterTelephoneContactPropertyTypeDeleteCallback.OnAfterDelete(stage, telephonecontactpropertytype, frontTelephoneContactPropertyType)
	}
}

func (telephonecontacttype *TelephoneContactType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTelephoneContactTypeCreateCallback != nil {
		stage.OnAfterTelephoneContactTypeCreateCallback.OnAfterCreate(stage, telephonecontacttype)
	}
}

func (telephonecontacttype *TelephoneContactType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTelephoneContactTypeUpdateCallback != nil {
		var frontTelephoneContactType *TelephoneContactType
		if front != nil {
			frontTelephoneContactType, _ = front.(*TelephoneContactType)
		}
		stage.OnAfterTelephoneContactTypeUpdateCallback.OnAfterUpdate(stage, telephonecontacttype, frontTelephoneContactType)
	}
}

func (telephonecontacttype *TelephoneContactType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTelephoneContactTypeDeleteCallback != nil {
		var frontTelephoneContactType *TelephoneContactType
		if front != nil {
			frontTelephoneContactType, _ = front.(*TelephoneContactType)
		}
		stage.OnAfterTelephoneContactTypeDeleteCallback.OnAfterDelete(stage, telephonecontacttype, frontTelephoneContactType)
	}
}

func (timesheetpropertytype *TimesheetPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimesheetPropertyTypeCreateCallback != nil {
		stage.OnAfterTimesheetPropertyTypeCreateCallback.OnAfterCreate(stage, timesheetpropertytype)
	}
}

func (timesheetpropertytype *TimesheetPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimesheetPropertyTypeUpdateCallback != nil {
		var frontTimesheetPropertyType *TimesheetPropertyType
		if front != nil {
			frontTimesheetPropertyType, _ = front.(*TimesheetPropertyType)
		}
		stage.OnAfterTimesheetPropertyTypeUpdateCallback.OnAfterUpdate(stage, timesheetpropertytype, frontTimesheetPropertyType)
	}
}

func (timesheetpropertytype *TimesheetPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimesheetPropertyTypeDeleteCallback != nil {
		var frontTimesheetPropertyType *TimesheetPropertyType
		if front != nil {
			frontTimesheetPropertyType, _ = front.(*TimesheetPropertyType)
		}
		stage.OnAfterTimesheetPropertyTypeDeleteCallback.OnAfterDelete(stage, timesheetpropertytype, frontTimesheetPropertyType)
	}
}

func (timesheettype *TimesheetType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimesheetTypeCreateCallback != nil {
		stage.OnAfterTimesheetTypeCreateCallback.OnAfterCreate(stage, timesheettype)
	}
}

func (timesheettype *TimesheetType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimesheetTypeUpdateCallback != nil {
		var frontTimesheetType *TimesheetType
		if front != nil {
			frontTimesheetType, _ = front.(*TimesheetType)
		}
		stage.OnAfterTimesheetTypeUpdateCallback.OnAfterUpdate(stage, timesheettype, frontTimesheetType)
	}
}

func (timesheettype *TimesheetType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimesheetTypeDeleteCallback != nil {
		var frontTimesheetType *TimesheetType
		if front != nil {
			frontTimesheetType, _ = front.(*TimesheetType)
		}
		stage.OnAfterTimesheetTypeDeleteCallback.OnAfterDelete(stage, timesheettype, frontTimesheetType)
	}
}

func (touchdownliftoffcontaminationpropertytype *TouchDownLiftOffContaminationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffContaminationPropertyTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffContaminationPropertyTypeCreateCallback.OnAfterCreate(stage, touchdownliftoffcontaminationpropertytype)
	}
}

func (touchdownliftoffcontaminationpropertytype *TouchDownLiftOffContaminationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffContaminationPropertyTypeUpdateCallback != nil {
		var frontTouchDownLiftOffContaminationPropertyType *TouchDownLiftOffContaminationPropertyType
		if front != nil {
			frontTouchDownLiftOffContaminationPropertyType, _ = front.(*TouchDownLiftOffContaminationPropertyType)
		}
		stage.OnAfterTouchDownLiftOffContaminationPropertyTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftoffcontaminationpropertytype, frontTouchDownLiftOffContaminationPropertyType)
	}
}

func (touchdownliftoffcontaminationpropertytype *TouchDownLiftOffContaminationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffContaminationPropertyTypeDeleteCallback != nil {
		var frontTouchDownLiftOffContaminationPropertyType *TouchDownLiftOffContaminationPropertyType
		if front != nil {
			frontTouchDownLiftOffContaminationPropertyType, _ = front.(*TouchDownLiftOffContaminationPropertyType)
		}
		stage.OnAfterTouchDownLiftOffContaminationPropertyTypeDeleteCallback.OnAfterDelete(stage, touchdownliftoffcontaminationpropertytype, frontTouchDownLiftOffContaminationPropertyType)
	}
}

func (touchdownliftoffcontaminationtype *TouchDownLiftOffContaminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffContaminationTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffContaminationTypeCreateCallback.OnAfterCreate(stage, touchdownliftoffcontaminationtype)
	}
}

func (touchdownliftoffcontaminationtype *TouchDownLiftOffContaminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffContaminationTypeUpdateCallback != nil {
		var frontTouchDownLiftOffContaminationType *TouchDownLiftOffContaminationType
		if front != nil {
			frontTouchDownLiftOffContaminationType, _ = front.(*TouchDownLiftOffContaminationType)
		}
		stage.OnAfterTouchDownLiftOffContaminationTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftoffcontaminationtype, frontTouchDownLiftOffContaminationType)
	}
}

func (touchdownliftoffcontaminationtype *TouchDownLiftOffContaminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffContaminationTypeDeleteCallback != nil {
		var frontTouchDownLiftOffContaminationType *TouchDownLiftOffContaminationType
		if front != nil {
			frontTouchDownLiftOffContaminationType, _ = front.(*TouchDownLiftOffContaminationType)
		}
		stage.OnAfterTouchDownLiftOffContaminationTypeDeleteCallback.OnAfterDelete(stage, touchdownliftoffcontaminationtype, frontTouchDownLiftOffContaminationType)
	}
}

func (touchdownliftofflightsystemtimeslicepropertytype *TouchDownLiftOffLightSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffLightSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffLightSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, touchdownliftofflightsystemtimeslicepropertytype)
	}
}

func (touchdownliftofflightsystemtimeslicepropertytype *TouchDownLiftOffLightSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffLightSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTouchDownLiftOffLightSystemTimeSlicePropertyType *TouchDownLiftOffLightSystemTimeSlicePropertyType
		if front != nil {
			frontTouchDownLiftOffLightSystemTimeSlicePropertyType, _ = front.(*TouchDownLiftOffLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterTouchDownLiftOffLightSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftofflightsystemtimeslicepropertytype, frontTouchDownLiftOffLightSystemTimeSlicePropertyType)
	}
}

func (touchdownliftofflightsystemtimeslicepropertytype *TouchDownLiftOffLightSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffLightSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTouchDownLiftOffLightSystemTimeSlicePropertyType *TouchDownLiftOffLightSystemTimeSlicePropertyType
		if front != nil {
			frontTouchDownLiftOffLightSystemTimeSlicePropertyType, _ = front.(*TouchDownLiftOffLightSystemTimeSlicePropertyType)
		}
		stage.OnAfterTouchDownLiftOffLightSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, touchdownliftofflightsystemtimeslicepropertytype, frontTouchDownLiftOffLightSystemTimeSlicePropertyType)
	}
}

func (touchdownliftofflightsystemtimeslicetype *TouchDownLiftOffLightSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffLightSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffLightSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, touchdownliftofflightsystemtimeslicetype)
	}
}

func (touchdownliftofflightsystemtimeslicetype *TouchDownLiftOffLightSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffLightSystemTimeSliceTypeUpdateCallback != nil {
		var frontTouchDownLiftOffLightSystemTimeSliceType *TouchDownLiftOffLightSystemTimeSliceType
		if front != nil {
			frontTouchDownLiftOffLightSystemTimeSliceType, _ = front.(*TouchDownLiftOffLightSystemTimeSliceType)
		}
		stage.OnAfterTouchDownLiftOffLightSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftofflightsystemtimeslicetype, frontTouchDownLiftOffLightSystemTimeSliceType)
	}
}

func (touchdownliftofflightsystemtimeslicetype *TouchDownLiftOffLightSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffLightSystemTimeSliceTypeDeleteCallback != nil {
		var frontTouchDownLiftOffLightSystemTimeSliceType *TouchDownLiftOffLightSystemTimeSliceType
		if front != nil {
			frontTouchDownLiftOffLightSystemTimeSliceType, _ = front.(*TouchDownLiftOffLightSystemTimeSliceType)
		}
		stage.OnAfterTouchDownLiftOffLightSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, touchdownliftofflightsystemtimeslicetype, frontTouchDownLiftOffLightSystemTimeSliceType)
	}
}

func (touchdownliftofflightsystemtype *TouchDownLiftOffLightSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffLightSystemTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffLightSystemTypeCreateCallback.OnAfterCreate(stage, touchdownliftofflightsystemtype)
	}
}

func (touchdownliftofflightsystemtype *TouchDownLiftOffLightSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffLightSystemTypeUpdateCallback != nil {
		var frontTouchDownLiftOffLightSystemType *TouchDownLiftOffLightSystemType
		if front != nil {
			frontTouchDownLiftOffLightSystemType, _ = front.(*TouchDownLiftOffLightSystemType)
		}
		stage.OnAfterTouchDownLiftOffLightSystemTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftofflightsystemtype, frontTouchDownLiftOffLightSystemType)
	}
}

func (touchdownliftofflightsystemtype *TouchDownLiftOffLightSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffLightSystemTypeDeleteCallback != nil {
		var frontTouchDownLiftOffLightSystemType *TouchDownLiftOffLightSystemType
		if front != nil {
			frontTouchDownLiftOffLightSystemType, _ = front.(*TouchDownLiftOffLightSystemType)
		}
		stage.OnAfterTouchDownLiftOffLightSystemTypeDeleteCallback.OnAfterDelete(stage, touchdownliftofflightsystemtype, frontTouchDownLiftOffLightSystemType)
	}
}

func (touchdownliftoffmarkingtimeslicepropertytype *TouchDownLiftOffMarkingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffMarkingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffMarkingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, touchdownliftoffmarkingtimeslicepropertytype)
	}
}

func (touchdownliftoffmarkingtimeslicepropertytype *TouchDownLiftOffMarkingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffMarkingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTouchDownLiftOffMarkingTimeSlicePropertyType *TouchDownLiftOffMarkingTimeSlicePropertyType
		if front != nil {
			frontTouchDownLiftOffMarkingTimeSlicePropertyType, _ = front.(*TouchDownLiftOffMarkingTimeSlicePropertyType)
		}
		stage.OnAfterTouchDownLiftOffMarkingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftoffmarkingtimeslicepropertytype, frontTouchDownLiftOffMarkingTimeSlicePropertyType)
	}
}

func (touchdownliftoffmarkingtimeslicepropertytype *TouchDownLiftOffMarkingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffMarkingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTouchDownLiftOffMarkingTimeSlicePropertyType *TouchDownLiftOffMarkingTimeSlicePropertyType
		if front != nil {
			frontTouchDownLiftOffMarkingTimeSlicePropertyType, _ = front.(*TouchDownLiftOffMarkingTimeSlicePropertyType)
		}
		stage.OnAfterTouchDownLiftOffMarkingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, touchdownliftoffmarkingtimeslicepropertytype, frontTouchDownLiftOffMarkingTimeSlicePropertyType)
	}
}

func (touchdownliftoffmarkingtimeslicetype *TouchDownLiftOffMarkingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffMarkingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffMarkingTimeSliceTypeCreateCallback.OnAfterCreate(stage, touchdownliftoffmarkingtimeslicetype)
	}
}

func (touchdownliftoffmarkingtimeslicetype *TouchDownLiftOffMarkingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffMarkingTimeSliceTypeUpdateCallback != nil {
		var frontTouchDownLiftOffMarkingTimeSliceType *TouchDownLiftOffMarkingTimeSliceType
		if front != nil {
			frontTouchDownLiftOffMarkingTimeSliceType, _ = front.(*TouchDownLiftOffMarkingTimeSliceType)
		}
		stage.OnAfterTouchDownLiftOffMarkingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftoffmarkingtimeslicetype, frontTouchDownLiftOffMarkingTimeSliceType)
	}
}

func (touchdownliftoffmarkingtimeslicetype *TouchDownLiftOffMarkingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffMarkingTimeSliceTypeDeleteCallback != nil {
		var frontTouchDownLiftOffMarkingTimeSliceType *TouchDownLiftOffMarkingTimeSliceType
		if front != nil {
			frontTouchDownLiftOffMarkingTimeSliceType, _ = front.(*TouchDownLiftOffMarkingTimeSliceType)
		}
		stage.OnAfterTouchDownLiftOffMarkingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, touchdownliftoffmarkingtimeslicetype, frontTouchDownLiftOffMarkingTimeSliceType)
	}
}

func (touchdownliftoffmarkingtype *TouchDownLiftOffMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffMarkingTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffMarkingTypeCreateCallback.OnAfterCreate(stage, touchdownliftoffmarkingtype)
	}
}

func (touchdownliftoffmarkingtype *TouchDownLiftOffMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffMarkingTypeUpdateCallback != nil {
		var frontTouchDownLiftOffMarkingType *TouchDownLiftOffMarkingType
		if front != nil {
			frontTouchDownLiftOffMarkingType, _ = front.(*TouchDownLiftOffMarkingType)
		}
		stage.OnAfterTouchDownLiftOffMarkingTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftoffmarkingtype, frontTouchDownLiftOffMarkingType)
	}
}

func (touchdownliftoffmarkingtype *TouchDownLiftOffMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffMarkingTypeDeleteCallback != nil {
		var frontTouchDownLiftOffMarkingType *TouchDownLiftOffMarkingType
		if front != nil {
			frontTouchDownLiftOffMarkingType, _ = front.(*TouchDownLiftOffMarkingType)
		}
		stage.OnAfterTouchDownLiftOffMarkingTypeDeleteCallback.OnAfterDelete(stage, touchdownliftoffmarkingtype, frontTouchDownLiftOffMarkingType)
	}
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffPropertyTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffPropertyTypeCreateCallback.OnAfterCreate(stage, touchdownliftoffpropertytype)
	}
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffPropertyTypeUpdateCallback != nil {
		var frontTouchDownLiftOffPropertyType *TouchDownLiftOffPropertyType
		if front != nil {
			frontTouchDownLiftOffPropertyType, _ = front.(*TouchDownLiftOffPropertyType)
		}
		stage.OnAfterTouchDownLiftOffPropertyTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftoffpropertytype, frontTouchDownLiftOffPropertyType)
	}
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffPropertyTypeDeleteCallback != nil {
		var frontTouchDownLiftOffPropertyType *TouchDownLiftOffPropertyType
		if front != nil {
			frontTouchDownLiftOffPropertyType, _ = front.(*TouchDownLiftOffPropertyType)
		}
		stage.OnAfterTouchDownLiftOffPropertyTypeDeleteCallback.OnAfterDelete(stage, touchdownliftoffpropertytype, frontTouchDownLiftOffPropertyType)
	}
}

func (touchdownliftoffsafeareatimeslicepropertytype *TouchDownLiftOffSafeAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffSafeAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffSafeAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, touchdownliftoffsafeareatimeslicepropertytype)
	}
}

func (touchdownliftoffsafeareatimeslicepropertytype *TouchDownLiftOffSafeAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffSafeAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTouchDownLiftOffSafeAreaTimeSlicePropertyType *TouchDownLiftOffSafeAreaTimeSlicePropertyType
		if front != nil {
			frontTouchDownLiftOffSafeAreaTimeSlicePropertyType, _ = front.(*TouchDownLiftOffSafeAreaTimeSlicePropertyType)
		}
		stage.OnAfterTouchDownLiftOffSafeAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftoffsafeareatimeslicepropertytype, frontTouchDownLiftOffSafeAreaTimeSlicePropertyType)
	}
}

func (touchdownliftoffsafeareatimeslicepropertytype *TouchDownLiftOffSafeAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffSafeAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTouchDownLiftOffSafeAreaTimeSlicePropertyType *TouchDownLiftOffSafeAreaTimeSlicePropertyType
		if front != nil {
			frontTouchDownLiftOffSafeAreaTimeSlicePropertyType, _ = front.(*TouchDownLiftOffSafeAreaTimeSlicePropertyType)
		}
		stage.OnAfterTouchDownLiftOffSafeAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, touchdownliftoffsafeareatimeslicepropertytype, frontTouchDownLiftOffSafeAreaTimeSlicePropertyType)
	}
}

func (touchdownliftoffsafeareatimeslicetype *TouchDownLiftOffSafeAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffSafeAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffSafeAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, touchdownliftoffsafeareatimeslicetype)
	}
}

func (touchdownliftoffsafeareatimeslicetype *TouchDownLiftOffSafeAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffSafeAreaTimeSliceTypeUpdateCallback != nil {
		var frontTouchDownLiftOffSafeAreaTimeSliceType *TouchDownLiftOffSafeAreaTimeSliceType
		if front != nil {
			frontTouchDownLiftOffSafeAreaTimeSliceType, _ = front.(*TouchDownLiftOffSafeAreaTimeSliceType)
		}
		stage.OnAfterTouchDownLiftOffSafeAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftoffsafeareatimeslicetype, frontTouchDownLiftOffSafeAreaTimeSliceType)
	}
}

func (touchdownliftoffsafeareatimeslicetype *TouchDownLiftOffSafeAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffSafeAreaTimeSliceTypeDeleteCallback != nil {
		var frontTouchDownLiftOffSafeAreaTimeSliceType *TouchDownLiftOffSafeAreaTimeSliceType
		if front != nil {
			frontTouchDownLiftOffSafeAreaTimeSliceType, _ = front.(*TouchDownLiftOffSafeAreaTimeSliceType)
		}
		stage.OnAfterTouchDownLiftOffSafeAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, touchdownliftoffsafeareatimeslicetype, frontTouchDownLiftOffSafeAreaTimeSliceType)
	}
}

func (touchdownliftoffsafeareatype *TouchDownLiftOffSafeAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffSafeAreaTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffSafeAreaTypeCreateCallback.OnAfterCreate(stage, touchdownliftoffsafeareatype)
	}
}

func (touchdownliftoffsafeareatype *TouchDownLiftOffSafeAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffSafeAreaTypeUpdateCallback != nil {
		var frontTouchDownLiftOffSafeAreaType *TouchDownLiftOffSafeAreaType
		if front != nil {
			frontTouchDownLiftOffSafeAreaType, _ = front.(*TouchDownLiftOffSafeAreaType)
		}
		stage.OnAfterTouchDownLiftOffSafeAreaTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftoffsafeareatype, frontTouchDownLiftOffSafeAreaType)
	}
}

func (touchdownliftoffsafeareatype *TouchDownLiftOffSafeAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffSafeAreaTypeDeleteCallback != nil {
		var frontTouchDownLiftOffSafeAreaType *TouchDownLiftOffSafeAreaType
		if front != nil {
			frontTouchDownLiftOffSafeAreaType, _ = front.(*TouchDownLiftOffSafeAreaType)
		}
		stage.OnAfterTouchDownLiftOffSafeAreaTypeDeleteCallback.OnAfterDelete(stage, touchdownliftoffsafeareatype, frontTouchDownLiftOffSafeAreaType)
	}
}

func (touchdownliftofftimeslicepropertytype *TouchDownLiftOffTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, touchdownliftofftimeslicepropertytype)
	}
}

func (touchdownliftofftimeslicepropertytype *TouchDownLiftOffTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTouchDownLiftOffTimeSlicePropertyType *TouchDownLiftOffTimeSlicePropertyType
		if front != nil {
			frontTouchDownLiftOffTimeSlicePropertyType, _ = front.(*TouchDownLiftOffTimeSlicePropertyType)
		}
		stage.OnAfterTouchDownLiftOffTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftofftimeslicepropertytype, frontTouchDownLiftOffTimeSlicePropertyType)
	}
}

func (touchdownliftofftimeslicepropertytype *TouchDownLiftOffTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTouchDownLiftOffTimeSlicePropertyType *TouchDownLiftOffTimeSlicePropertyType
		if front != nil {
			frontTouchDownLiftOffTimeSlicePropertyType, _ = front.(*TouchDownLiftOffTimeSlicePropertyType)
		}
		stage.OnAfterTouchDownLiftOffTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, touchdownliftofftimeslicepropertytype, frontTouchDownLiftOffTimeSlicePropertyType)
	}
}

func (touchdownliftofftimeslicetype *TouchDownLiftOffTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffTimeSliceTypeCreateCallback.OnAfterCreate(stage, touchdownliftofftimeslicetype)
	}
}

func (touchdownliftofftimeslicetype *TouchDownLiftOffTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffTimeSliceTypeUpdateCallback != nil {
		var frontTouchDownLiftOffTimeSliceType *TouchDownLiftOffTimeSliceType
		if front != nil {
			frontTouchDownLiftOffTimeSliceType, _ = front.(*TouchDownLiftOffTimeSliceType)
		}
		stage.OnAfterTouchDownLiftOffTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftofftimeslicetype, frontTouchDownLiftOffTimeSliceType)
	}
}

func (touchdownliftofftimeslicetype *TouchDownLiftOffTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffTimeSliceTypeDeleteCallback != nil {
		var frontTouchDownLiftOffTimeSliceType *TouchDownLiftOffTimeSliceType
		if front != nil {
			frontTouchDownLiftOffTimeSliceType, _ = front.(*TouchDownLiftOffTimeSliceType)
		}
		stage.OnAfterTouchDownLiftOffTimeSliceTypeDeleteCallback.OnAfterDelete(stage, touchdownliftofftimeslicetype, frontTouchDownLiftOffTimeSliceType)
	}
}

func (touchdownliftofftype *TouchDownLiftOffType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTouchDownLiftOffTypeCreateCallback != nil {
		stage.OnAfterTouchDownLiftOffTypeCreateCallback.OnAfterCreate(stage, touchdownliftofftype)
	}
}

func (touchdownliftofftype *TouchDownLiftOffType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffTypeUpdateCallback != nil {
		var frontTouchDownLiftOffType *TouchDownLiftOffType
		if front != nil {
			frontTouchDownLiftOffType, _ = front.(*TouchDownLiftOffType)
		}
		stage.OnAfterTouchDownLiftOffTypeUpdateCallback.OnAfterUpdate(stage, touchdownliftofftype, frontTouchDownLiftOffType)
	}
}

func (touchdownliftofftype *TouchDownLiftOffType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTouchDownLiftOffTypeDeleteCallback != nil {
		var frontTouchDownLiftOffType *TouchDownLiftOffType
		if front != nil {
			frontTouchDownLiftOffType, _ = front.(*TouchDownLiftOffType)
		}
		stage.OnAfterTouchDownLiftOffTypeDeleteCallback.OnAfterDelete(stage, touchdownliftofftype, frontTouchDownLiftOffType)
	}
}

func (visualglideslopeindicatortimeslicepropertytype *VisualGlideSlopeIndicatorTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVisualGlideSlopeIndicatorTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterVisualGlideSlopeIndicatorTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, visualglideslopeindicatortimeslicepropertytype)
	}
}

func (visualglideslopeindicatortimeslicepropertytype *VisualGlideSlopeIndicatorTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVisualGlideSlopeIndicatorTimeSlicePropertyTypeUpdateCallback != nil {
		var frontVisualGlideSlopeIndicatorTimeSlicePropertyType *VisualGlideSlopeIndicatorTimeSlicePropertyType
		if front != nil {
			frontVisualGlideSlopeIndicatorTimeSlicePropertyType, _ = front.(*VisualGlideSlopeIndicatorTimeSlicePropertyType)
		}
		stage.OnAfterVisualGlideSlopeIndicatorTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, visualglideslopeindicatortimeslicepropertytype, frontVisualGlideSlopeIndicatorTimeSlicePropertyType)
	}
}

func (visualglideslopeindicatortimeslicepropertytype *VisualGlideSlopeIndicatorTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVisualGlideSlopeIndicatorTimeSlicePropertyTypeDeleteCallback != nil {
		var frontVisualGlideSlopeIndicatorTimeSlicePropertyType *VisualGlideSlopeIndicatorTimeSlicePropertyType
		if front != nil {
			frontVisualGlideSlopeIndicatorTimeSlicePropertyType, _ = front.(*VisualGlideSlopeIndicatorTimeSlicePropertyType)
		}
		stage.OnAfterVisualGlideSlopeIndicatorTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, visualglideslopeindicatortimeslicepropertytype, frontVisualGlideSlopeIndicatorTimeSlicePropertyType)
	}
}

func (visualglideslopeindicatortimeslicetype *VisualGlideSlopeIndicatorTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVisualGlideSlopeIndicatorTimeSliceTypeCreateCallback != nil {
		stage.OnAfterVisualGlideSlopeIndicatorTimeSliceTypeCreateCallback.OnAfterCreate(stage, visualglideslopeindicatortimeslicetype)
	}
}

func (visualglideslopeindicatortimeslicetype *VisualGlideSlopeIndicatorTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVisualGlideSlopeIndicatorTimeSliceTypeUpdateCallback != nil {
		var frontVisualGlideSlopeIndicatorTimeSliceType *VisualGlideSlopeIndicatorTimeSliceType
		if front != nil {
			frontVisualGlideSlopeIndicatorTimeSliceType, _ = front.(*VisualGlideSlopeIndicatorTimeSliceType)
		}
		stage.OnAfterVisualGlideSlopeIndicatorTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, visualglideslopeindicatortimeslicetype, frontVisualGlideSlopeIndicatorTimeSliceType)
	}
}

func (visualglideslopeindicatortimeslicetype *VisualGlideSlopeIndicatorTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVisualGlideSlopeIndicatorTimeSliceTypeDeleteCallback != nil {
		var frontVisualGlideSlopeIndicatorTimeSliceType *VisualGlideSlopeIndicatorTimeSliceType
		if front != nil {
			frontVisualGlideSlopeIndicatorTimeSliceType, _ = front.(*VisualGlideSlopeIndicatorTimeSliceType)
		}
		stage.OnAfterVisualGlideSlopeIndicatorTimeSliceTypeDeleteCallback.OnAfterDelete(stage, visualglideslopeindicatortimeslicetype, frontVisualGlideSlopeIndicatorTimeSliceType)
	}
}

func (visualglideslopeindicatortype *VisualGlideSlopeIndicatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVisualGlideSlopeIndicatorTypeCreateCallback != nil {
		stage.OnAfterVisualGlideSlopeIndicatorTypeCreateCallback.OnAfterCreate(stage, visualglideslopeindicatortype)
	}
}

func (visualglideslopeindicatortype *VisualGlideSlopeIndicatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVisualGlideSlopeIndicatorTypeUpdateCallback != nil {
		var frontVisualGlideSlopeIndicatorType *VisualGlideSlopeIndicatorType
		if front != nil {
			frontVisualGlideSlopeIndicatorType, _ = front.(*VisualGlideSlopeIndicatorType)
		}
		stage.OnAfterVisualGlideSlopeIndicatorTypeUpdateCallback.OnAfterUpdate(stage, visualglideslopeindicatortype, frontVisualGlideSlopeIndicatorType)
	}
}

func (visualglideslopeindicatortype *VisualGlideSlopeIndicatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVisualGlideSlopeIndicatorTypeDeleteCallback != nil {
		var frontVisualGlideSlopeIndicatorType *VisualGlideSlopeIndicatorType
		if front != nil {
			frontVisualGlideSlopeIndicatorType, _ = front.(*VisualGlideSlopeIndicatorType)
		}
		stage.OnAfterVisualGlideSlopeIndicatorTypeDeleteCallback.OnAfterDelete(stage, visualglideslopeindicatortype, frontVisualGlideSlopeIndicatorType)
	}
}

func (workareatimeslicepropertytype *WorkAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWorkAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterWorkAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, workareatimeslicepropertytype)
	}
}

func (workareatimeslicepropertytype *WorkAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontWorkAreaTimeSlicePropertyType *WorkAreaTimeSlicePropertyType
		if front != nil {
			frontWorkAreaTimeSlicePropertyType, _ = front.(*WorkAreaTimeSlicePropertyType)
		}
		stage.OnAfterWorkAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, workareatimeslicepropertytype, frontWorkAreaTimeSlicePropertyType)
	}
}

func (workareatimeslicepropertytype *WorkAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontWorkAreaTimeSlicePropertyType *WorkAreaTimeSlicePropertyType
		if front != nil {
			frontWorkAreaTimeSlicePropertyType, _ = front.(*WorkAreaTimeSlicePropertyType)
		}
		stage.OnAfterWorkAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, workareatimeslicepropertytype, frontWorkAreaTimeSlicePropertyType)
	}
}

func (workareatimeslicetype *WorkAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWorkAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterWorkAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, workareatimeslicetype)
	}
}

func (workareatimeslicetype *WorkAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkAreaTimeSliceTypeUpdateCallback != nil {
		var frontWorkAreaTimeSliceType *WorkAreaTimeSliceType
		if front != nil {
			frontWorkAreaTimeSliceType, _ = front.(*WorkAreaTimeSliceType)
		}
		stage.OnAfterWorkAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, workareatimeslicetype, frontWorkAreaTimeSliceType)
	}
}

func (workareatimeslicetype *WorkAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkAreaTimeSliceTypeDeleteCallback != nil {
		var frontWorkAreaTimeSliceType *WorkAreaTimeSliceType
		if front != nil {
			frontWorkAreaTimeSliceType, _ = front.(*WorkAreaTimeSliceType)
		}
		stage.OnAfterWorkAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, workareatimeslicetype, frontWorkAreaTimeSliceType)
	}
}

func (workareatype *WorkAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWorkAreaTypeCreateCallback != nil {
		stage.OnAfterWorkAreaTypeCreateCallback.OnAfterCreate(stage, workareatype)
	}
}

func (workareatype *WorkAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkAreaTypeUpdateCallback != nil {
		var frontWorkAreaType *WorkAreaType
		if front != nil {
			frontWorkAreaType, _ = front.(*WorkAreaType)
		}
		stage.OnAfterWorkAreaTypeUpdateCallback.OnAfterUpdate(stage, workareatype, frontWorkAreaType)
	}
}

func (workareatype *WorkAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkAreaTypeDeleteCallback != nil {
		var frontWorkAreaType *WorkAreaType
		if front != nil {
			frontWorkAreaType, _ = front.(*WorkAreaType)
		}
		stage.OnAfterWorkAreaTypeDeleteCallback.OnAfterDelete(stage, workareatype, frontWorkAreaType)
	}
}

func (workareaactivitypropertytype *WorkareaActivityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWorkareaActivityPropertyTypeCreateCallback != nil {
		stage.OnAfterWorkareaActivityPropertyTypeCreateCallback.OnAfterCreate(stage, workareaactivitypropertytype)
	}
}

func (workareaactivitypropertytype *WorkareaActivityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkareaActivityPropertyTypeUpdateCallback != nil {
		var frontWorkareaActivityPropertyType *WorkareaActivityPropertyType
		if front != nil {
			frontWorkareaActivityPropertyType, _ = front.(*WorkareaActivityPropertyType)
		}
		stage.OnAfterWorkareaActivityPropertyTypeUpdateCallback.OnAfterUpdate(stage, workareaactivitypropertytype, frontWorkareaActivityPropertyType)
	}
}

func (workareaactivitypropertytype *WorkareaActivityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkareaActivityPropertyTypeDeleteCallback != nil {
		var frontWorkareaActivityPropertyType *WorkareaActivityPropertyType
		if front != nil {
			frontWorkareaActivityPropertyType, _ = front.(*WorkareaActivityPropertyType)
		}
		stage.OnAfterWorkareaActivityPropertyTypeDeleteCallback.OnAfterDelete(stage, workareaactivitypropertytype, frontWorkareaActivityPropertyType)
	}
}

func (workareaactivitytype *WorkareaActivityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWorkareaActivityTypeCreateCallback != nil {
		stage.OnAfterWorkareaActivityTypeCreateCallback.OnAfterCreate(stage, workareaactivitytype)
	}
}

func (workareaactivitytype *WorkareaActivityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkareaActivityTypeUpdateCallback != nil {
		var frontWorkareaActivityType *WorkareaActivityType
		if front != nil {
			frontWorkareaActivityType, _ = front.(*WorkareaActivityType)
		}
		stage.OnAfterWorkareaActivityTypeUpdateCallback.OnAfterUpdate(stage, workareaactivitytype, frontWorkareaActivityType)
	}
}

func (workareaactivitytype *WorkareaActivityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWorkareaActivityTypeDeleteCallback != nil {
		var frontWorkareaActivityType *WorkareaActivityType
		if front != nil {
			frontWorkareaActivityType, _ = front.(*WorkareaActivityType)
		}
		stage.OnAfterWorkareaActivityTypeDeleteCallback.OnAfterDelete(stage, workareaactivitytype, frontWorkareaActivityType)
	}
}

