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

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractNavaidEquipmentTypeCreateCallback != nil {
		stage.OnAfterAbstractNavaidEquipmentTypeCreateCallback.OnAfterCreate(stage, abstractnavaidequipmenttype)
	}
}

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractNavaidEquipmentTypeUpdateCallback != nil {
		var frontAbstractNavaidEquipmentType *AbstractNavaidEquipmentType
		if front != nil {
			frontAbstractNavaidEquipmentType, _ = front.(*AbstractNavaidEquipmentType)
		}
		stage.OnAfterAbstractNavaidEquipmentTypeUpdateCallback.OnAfterUpdate(stage, abstractnavaidequipmenttype, frontAbstractNavaidEquipmentType)
	}
}

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractNavaidEquipmentTypeDeleteCallback != nil {
		var frontAbstractNavaidEquipmentType *AbstractNavaidEquipmentType
		if front != nil {
			frontAbstractNavaidEquipmentType, _ = front.(*AbstractNavaidEquipmentType)
		}
		stage.OnAfterAbstractNavaidEquipmentTypeDeleteCallback.OnAfterDelete(stage, abstractnavaidequipmenttype, frontAbstractNavaidEquipmentType)
	}
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractNavigationSystemCheckpointTypeCreateCallback != nil {
		stage.OnAfterAbstractNavigationSystemCheckpointTypeCreateCallback.OnAfterCreate(stage, abstractnavigationsystemcheckpointtype)
	}
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractNavigationSystemCheckpointTypeUpdateCallback != nil {
		var frontAbstractNavigationSystemCheckpointType *AbstractNavigationSystemCheckpointType
		if front != nil {
			frontAbstractNavigationSystemCheckpointType, _ = front.(*AbstractNavigationSystemCheckpointType)
		}
		stage.OnAfterAbstractNavigationSystemCheckpointTypeUpdateCallback.OnAfterUpdate(stage, abstractnavigationsystemcheckpointtype, frontAbstractNavigationSystemCheckpointType)
	}
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractNavigationSystemCheckpointTypeDeleteCallback != nil {
		var frontAbstractNavigationSystemCheckpointType *AbstractNavigationSystemCheckpointType
		if front != nil {
			frontAbstractNavigationSystemCheckpointType, _ = front.(*AbstractNavigationSystemCheckpointType)
		}
		stage.OnAfterAbstractNavigationSystemCheckpointTypeDeleteCallback.OnAfterDelete(stage, abstractnavigationsystemcheckpointtype, frontAbstractNavigationSystemCheckpointType)
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

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractRadarEquipmentTypeCreateCallback != nil {
		stage.OnAfterAbstractRadarEquipmentTypeCreateCallback.OnAfterCreate(stage, abstractradarequipmenttype)
	}
}

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRadarEquipmentTypeUpdateCallback != nil {
		var frontAbstractRadarEquipmentType *AbstractRadarEquipmentType
		if front != nil {
			frontAbstractRadarEquipmentType, _ = front.(*AbstractRadarEquipmentType)
		}
		stage.OnAfterAbstractRadarEquipmentTypeUpdateCallback.OnAfterUpdate(stage, abstractradarequipmenttype, frontAbstractRadarEquipmentType)
	}
}

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRadarEquipmentTypeDeleteCallback != nil {
		var frontAbstractRadarEquipmentType *AbstractRadarEquipmentType
		if front != nil {
			frontAbstractRadarEquipmentType, _ = front.(*AbstractRadarEquipmentType)
		}
		stage.OnAfterAbstractRadarEquipmentTypeDeleteCallback.OnAfterDelete(stage, abstractradarequipmenttype, frontAbstractRadarEquipmentType)
	}
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractSurveillanceRadarTypeCreateCallback != nil {
		stage.OnAfterAbstractSurveillanceRadarTypeCreateCallback.OnAfterCreate(stage, abstractsurveillanceradartype)
	}
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurveillanceRadarTypeUpdateCallback != nil {
		var frontAbstractSurveillanceRadarType *AbstractSurveillanceRadarType
		if front != nil {
			frontAbstractSurveillanceRadarType, _ = front.(*AbstractSurveillanceRadarType)
		}
		stage.OnAfterAbstractSurveillanceRadarTypeUpdateCallback.OnAfterUpdate(stage, abstractsurveillanceradartype, frontAbstractSurveillanceRadarType)
	}
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurveillanceRadarTypeDeleteCallback != nil {
		var frontAbstractSurveillanceRadarType *AbstractSurveillanceRadarType
		if front != nil {
			frontAbstractSurveillanceRadarType, _ = front.(*AbstractSurveillanceRadarType)
		}
		stage.OnAfterAbstractSurveillanceRadarTypeDeleteCallback.OnAfterDelete(stage, abstractsurveillanceradartype, frontAbstractSurveillanceRadarType)
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

func (airspacepropertytype *AirspacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspacePropertyTypeCreateCallback != nil {
		stage.OnAfterAirspacePropertyTypeCreateCallback.OnAfterCreate(stage, airspacepropertytype)
	}
}

func (airspacepropertytype *AirspacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspacePropertyTypeUpdateCallback != nil {
		var frontAirspacePropertyType *AirspacePropertyType
		if front != nil {
			frontAirspacePropertyType, _ = front.(*AirspacePropertyType)
		}
		stage.OnAfterAirspacePropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacepropertytype, frontAirspacePropertyType)
	}
}

func (airspacepropertytype *AirspacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspacePropertyTypeDeleteCallback != nil {
		var frontAirspacePropertyType *AirspacePropertyType
		if front != nil {
			frontAirspacePropertyType, _ = front.(*AirspacePropertyType)
		}
		stage.OnAfterAirspacePropertyTypeDeleteCallback.OnAfterDelete(stage, airspacepropertytype, frontAirspacePropertyType)
	}
}

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngleIndicationTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAngleIndicationTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, angleindicationtimeslicepropertytype)
	}
}

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleIndicationTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAngleIndicationTimeSlicePropertyType *AngleIndicationTimeSlicePropertyType
		if front != nil {
			frontAngleIndicationTimeSlicePropertyType, _ = front.(*AngleIndicationTimeSlicePropertyType)
		}
		stage.OnAfterAngleIndicationTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, angleindicationtimeslicepropertytype, frontAngleIndicationTimeSlicePropertyType)
	}
}

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleIndicationTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAngleIndicationTimeSlicePropertyType *AngleIndicationTimeSlicePropertyType
		if front != nil {
			frontAngleIndicationTimeSlicePropertyType, _ = front.(*AngleIndicationTimeSlicePropertyType)
		}
		stage.OnAfterAngleIndicationTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, angleindicationtimeslicepropertytype, frontAngleIndicationTimeSlicePropertyType)
	}
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngleIndicationTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAngleIndicationTimeSliceTypeCreateCallback.OnAfterCreate(stage, angleindicationtimeslicetype)
	}
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleIndicationTimeSliceTypeUpdateCallback != nil {
		var frontAngleIndicationTimeSliceType *AngleIndicationTimeSliceType
		if front != nil {
			frontAngleIndicationTimeSliceType, _ = front.(*AngleIndicationTimeSliceType)
		}
		stage.OnAfterAngleIndicationTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, angleindicationtimeslicetype, frontAngleIndicationTimeSliceType)
	}
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleIndicationTimeSliceTypeDeleteCallback != nil {
		var frontAngleIndicationTimeSliceType *AngleIndicationTimeSliceType
		if front != nil {
			frontAngleIndicationTimeSliceType, _ = front.(*AngleIndicationTimeSliceType)
		}
		stage.OnAfterAngleIndicationTimeSliceTypeDeleteCallback.OnAfterDelete(stage, angleindicationtimeslicetype, frontAngleIndicationTimeSliceType)
	}
}

func (angleindicationtype *AngleIndicationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngleIndicationTypeCreateCallback != nil {
		stage.OnAfterAngleIndicationTypeCreateCallback.OnAfterCreate(stage, angleindicationtype)
	}
}

func (angleindicationtype *AngleIndicationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleIndicationTypeUpdateCallback != nil {
		var frontAngleIndicationType *AngleIndicationType
		if front != nil {
			frontAngleIndicationType, _ = front.(*AngleIndicationType)
		}
		stage.OnAfterAngleIndicationTypeUpdateCallback.OnAfterUpdate(stage, angleindicationtype, frontAngleIndicationType)
	}
}

func (angleindicationtype *AngleIndicationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleIndicationTypeDeleteCallback != nil {
		var frontAngleIndicationType *AngleIndicationType
		if front != nil {
			frontAngleIndicationType, _ = front.(*AngleIndicationType)
		}
		stage.OnAfterAngleIndicationTypeDeleteCallback.OnAfterDelete(stage, angleindicationtype, frontAngleIndicationType)
	}
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForNavaidEquipmentPropertyTypeCreateCallback != nil {
		stage.OnAfterAuthorityForNavaidEquipmentPropertyTypeCreateCallback.OnAfterCreate(stage, authorityfornavaidequipmentpropertytype)
	}
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForNavaidEquipmentPropertyTypeUpdateCallback != nil {
		var frontAuthorityForNavaidEquipmentPropertyType *AuthorityForNavaidEquipmentPropertyType
		if front != nil {
			frontAuthorityForNavaidEquipmentPropertyType, _ = front.(*AuthorityForNavaidEquipmentPropertyType)
		}
		stage.OnAfterAuthorityForNavaidEquipmentPropertyTypeUpdateCallback.OnAfterUpdate(stage, authorityfornavaidequipmentpropertytype, frontAuthorityForNavaidEquipmentPropertyType)
	}
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForNavaidEquipmentPropertyTypeDeleteCallback != nil {
		var frontAuthorityForNavaidEquipmentPropertyType *AuthorityForNavaidEquipmentPropertyType
		if front != nil {
			frontAuthorityForNavaidEquipmentPropertyType, _ = front.(*AuthorityForNavaidEquipmentPropertyType)
		}
		stage.OnAfterAuthorityForNavaidEquipmentPropertyTypeDeleteCallback.OnAfterDelete(stage, authorityfornavaidequipmentpropertytype, frontAuthorityForNavaidEquipmentPropertyType)
	}
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForNavaidEquipmentTypeCreateCallback != nil {
		stage.OnAfterAuthorityForNavaidEquipmentTypeCreateCallback.OnAfterCreate(stage, authorityfornavaidequipmenttype)
	}
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForNavaidEquipmentTypeUpdateCallback != nil {
		var frontAuthorityForNavaidEquipmentType *AuthorityForNavaidEquipmentType
		if front != nil {
			frontAuthorityForNavaidEquipmentType, _ = front.(*AuthorityForNavaidEquipmentType)
		}
		stage.OnAfterAuthorityForNavaidEquipmentTypeUpdateCallback.OnAfterUpdate(stage, authorityfornavaidequipmenttype, frontAuthorityForNavaidEquipmentType)
	}
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForNavaidEquipmentTypeDeleteCallback != nil {
		var frontAuthorityForNavaidEquipmentType *AuthorityForNavaidEquipmentType
		if front != nil {
			frontAuthorityForNavaidEquipmentType, _ = front.(*AuthorityForNavaidEquipmentType)
		}
		stage.OnAfterAuthorityForNavaidEquipmentTypeDeleteCallback.OnAfterDelete(stage, authorityfornavaidequipmenttype, frontAuthorityForNavaidEquipmentType)
	}
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForSpecialNavigationStationPropertyTypeCreateCallback != nil {
		stage.OnAfterAuthorityForSpecialNavigationStationPropertyTypeCreateCallback.OnAfterCreate(stage, authorityforspecialnavigationstationpropertytype)
	}
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForSpecialNavigationStationPropertyTypeUpdateCallback != nil {
		var frontAuthorityForSpecialNavigationStationPropertyType *AuthorityForSpecialNavigationStationPropertyType
		if front != nil {
			frontAuthorityForSpecialNavigationStationPropertyType, _ = front.(*AuthorityForSpecialNavigationStationPropertyType)
		}
		stage.OnAfterAuthorityForSpecialNavigationStationPropertyTypeUpdateCallback.OnAfterUpdate(stage, authorityforspecialnavigationstationpropertytype, frontAuthorityForSpecialNavigationStationPropertyType)
	}
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForSpecialNavigationStationPropertyTypeDeleteCallback != nil {
		var frontAuthorityForSpecialNavigationStationPropertyType *AuthorityForSpecialNavigationStationPropertyType
		if front != nil {
			frontAuthorityForSpecialNavigationStationPropertyType, _ = front.(*AuthorityForSpecialNavigationStationPropertyType)
		}
		stage.OnAfterAuthorityForSpecialNavigationStationPropertyTypeDeleteCallback.OnAfterDelete(stage, authorityforspecialnavigationstationpropertytype, frontAuthorityForSpecialNavigationStationPropertyType)
	}
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForSpecialNavigationStationTypeCreateCallback != nil {
		stage.OnAfterAuthorityForSpecialNavigationStationTypeCreateCallback.OnAfterCreate(stage, authorityforspecialnavigationstationtype)
	}
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForSpecialNavigationStationTypeUpdateCallback != nil {
		var frontAuthorityForSpecialNavigationStationType *AuthorityForSpecialNavigationStationType
		if front != nil {
			frontAuthorityForSpecialNavigationStationType, _ = front.(*AuthorityForSpecialNavigationStationType)
		}
		stage.OnAfterAuthorityForSpecialNavigationStationTypeUpdateCallback.OnAfterUpdate(stage, authorityforspecialnavigationstationtype, frontAuthorityForSpecialNavigationStationType)
	}
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForSpecialNavigationStationTypeDeleteCallback != nil {
		var frontAuthorityForSpecialNavigationStationType *AuthorityForSpecialNavigationStationType
		if front != nil {
			frontAuthorityForSpecialNavigationStationType, _ = front.(*AuthorityForSpecialNavigationStationType)
		}
		stage.OnAfterAuthorityForSpecialNavigationStationTypeDeleteCallback.OnAfterDelete(stage, authorityforspecialnavigationstationtype, frontAuthorityForSpecialNavigationStationType)
	}
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForSpecialNavigationSystemPropertyTypeCreateCallback != nil {
		stage.OnAfterAuthorityForSpecialNavigationSystemPropertyTypeCreateCallback.OnAfterCreate(stage, authorityforspecialnavigationsystempropertytype)
	}
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForSpecialNavigationSystemPropertyTypeUpdateCallback != nil {
		var frontAuthorityForSpecialNavigationSystemPropertyType *AuthorityForSpecialNavigationSystemPropertyType
		if front != nil {
			frontAuthorityForSpecialNavigationSystemPropertyType, _ = front.(*AuthorityForSpecialNavigationSystemPropertyType)
		}
		stage.OnAfterAuthorityForSpecialNavigationSystemPropertyTypeUpdateCallback.OnAfterUpdate(stage, authorityforspecialnavigationsystempropertytype, frontAuthorityForSpecialNavigationSystemPropertyType)
	}
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForSpecialNavigationSystemPropertyTypeDeleteCallback != nil {
		var frontAuthorityForSpecialNavigationSystemPropertyType *AuthorityForSpecialNavigationSystemPropertyType
		if front != nil {
			frontAuthorityForSpecialNavigationSystemPropertyType, _ = front.(*AuthorityForSpecialNavigationSystemPropertyType)
		}
		stage.OnAfterAuthorityForSpecialNavigationSystemPropertyTypeDeleteCallback.OnAfterDelete(stage, authorityforspecialnavigationsystempropertytype, frontAuthorityForSpecialNavigationSystemPropertyType)
	}
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForSpecialNavigationSystemTypeCreateCallback != nil {
		stage.OnAfterAuthorityForSpecialNavigationSystemTypeCreateCallback.OnAfterCreate(stage, authorityforspecialnavigationsystemtype)
	}
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForSpecialNavigationSystemTypeUpdateCallback != nil {
		var frontAuthorityForSpecialNavigationSystemType *AuthorityForSpecialNavigationSystemType
		if front != nil {
			frontAuthorityForSpecialNavigationSystemType, _ = front.(*AuthorityForSpecialNavigationSystemType)
		}
		stage.OnAfterAuthorityForSpecialNavigationSystemTypeUpdateCallback.OnAfterUpdate(stage, authorityforspecialnavigationsystemtype, frontAuthorityForSpecialNavigationSystemType)
	}
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForSpecialNavigationSystemTypeDeleteCallback != nil {
		var frontAuthorityForSpecialNavigationSystemType *AuthorityForSpecialNavigationSystemType
		if front != nil {
			frontAuthorityForSpecialNavigationSystemType, _ = front.(*AuthorityForSpecialNavigationSystemType)
		}
		stage.OnAfterAuthorityForSpecialNavigationSystemTypeDeleteCallback.OnAfterDelete(stage, authorityforspecialnavigationsystemtype, frontAuthorityForSpecialNavigationSystemType)
	}
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAzimuthTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAzimuthTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, azimuthtimeslicepropertytype)
	}
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAzimuthTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAzimuthTimeSlicePropertyType *AzimuthTimeSlicePropertyType
		if front != nil {
			frontAzimuthTimeSlicePropertyType, _ = front.(*AzimuthTimeSlicePropertyType)
		}
		stage.OnAfterAzimuthTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, azimuthtimeslicepropertytype, frontAzimuthTimeSlicePropertyType)
	}
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAzimuthTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAzimuthTimeSlicePropertyType *AzimuthTimeSlicePropertyType
		if front != nil {
			frontAzimuthTimeSlicePropertyType, _ = front.(*AzimuthTimeSlicePropertyType)
		}
		stage.OnAfterAzimuthTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, azimuthtimeslicepropertytype, frontAzimuthTimeSlicePropertyType)
	}
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAzimuthTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAzimuthTimeSliceTypeCreateCallback.OnAfterCreate(stage, azimuthtimeslicetype)
	}
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAzimuthTimeSliceTypeUpdateCallback != nil {
		var frontAzimuthTimeSliceType *AzimuthTimeSliceType
		if front != nil {
			frontAzimuthTimeSliceType, _ = front.(*AzimuthTimeSliceType)
		}
		stage.OnAfterAzimuthTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, azimuthtimeslicetype, frontAzimuthTimeSliceType)
	}
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAzimuthTimeSliceTypeDeleteCallback != nil {
		var frontAzimuthTimeSliceType *AzimuthTimeSliceType
		if front != nil {
			frontAzimuthTimeSliceType, _ = front.(*AzimuthTimeSliceType)
		}
		stage.OnAfterAzimuthTimeSliceTypeDeleteCallback.OnAfterDelete(stage, azimuthtimeslicetype, frontAzimuthTimeSliceType)
	}
}

func (azimuthtype *AzimuthType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAzimuthTypeCreateCallback != nil {
		stage.OnAfterAzimuthTypeCreateCallback.OnAfterCreate(stage, azimuthtype)
	}
}

func (azimuthtype *AzimuthType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAzimuthTypeUpdateCallback != nil {
		var frontAzimuthType *AzimuthType
		if front != nil {
			frontAzimuthType, _ = front.(*AzimuthType)
		}
		stage.OnAfterAzimuthTypeUpdateCallback.OnAfterUpdate(stage, azimuthtype, frontAzimuthType)
	}
}

func (azimuthtype *AzimuthType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAzimuthTypeDeleteCallback != nil {
		var frontAzimuthType *AzimuthType
		if front != nil {
			frontAzimuthType, _ = front.(*AzimuthType)
		}
		stage.OnAfterAzimuthTypeDeleteCallback.OnAfterDelete(stage, azimuthtype, frontAzimuthType)
	}
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCheckpointINSTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterCheckpointINSTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, checkpointinstimeslicepropertytype)
	}
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointINSTimeSlicePropertyTypeUpdateCallback != nil {
		var frontCheckpointINSTimeSlicePropertyType *CheckpointINSTimeSlicePropertyType
		if front != nil {
			frontCheckpointINSTimeSlicePropertyType, _ = front.(*CheckpointINSTimeSlicePropertyType)
		}
		stage.OnAfterCheckpointINSTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, checkpointinstimeslicepropertytype, frontCheckpointINSTimeSlicePropertyType)
	}
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointINSTimeSlicePropertyTypeDeleteCallback != nil {
		var frontCheckpointINSTimeSlicePropertyType *CheckpointINSTimeSlicePropertyType
		if front != nil {
			frontCheckpointINSTimeSlicePropertyType, _ = front.(*CheckpointINSTimeSlicePropertyType)
		}
		stage.OnAfterCheckpointINSTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, checkpointinstimeslicepropertytype, frontCheckpointINSTimeSlicePropertyType)
	}
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCheckpointINSTimeSliceTypeCreateCallback != nil {
		stage.OnAfterCheckpointINSTimeSliceTypeCreateCallback.OnAfterCreate(stage, checkpointinstimeslicetype)
	}
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointINSTimeSliceTypeUpdateCallback != nil {
		var frontCheckpointINSTimeSliceType *CheckpointINSTimeSliceType
		if front != nil {
			frontCheckpointINSTimeSliceType, _ = front.(*CheckpointINSTimeSliceType)
		}
		stage.OnAfterCheckpointINSTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, checkpointinstimeslicetype, frontCheckpointINSTimeSliceType)
	}
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointINSTimeSliceTypeDeleteCallback != nil {
		var frontCheckpointINSTimeSliceType *CheckpointINSTimeSliceType
		if front != nil {
			frontCheckpointINSTimeSliceType, _ = front.(*CheckpointINSTimeSliceType)
		}
		stage.OnAfterCheckpointINSTimeSliceTypeDeleteCallback.OnAfterDelete(stage, checkpointinstimeslicetype, frontCheckpointINSTimeSliceType)
	}
}

func (checkpointinstype *CheckpointINSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCheckpointINSTypeCreateCallback != nil {
		stage.OnAfterCheckpointINSTypeCreateCallback.OnAfterCreate(stage, checkpointinstype)
	}
}

func (checkpointinstype *CheckpointINSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointINSTypeUpdateCallback != nil {
		var frontCheckpointINSType *CheckpointINSType
		if front != nil {
			frontCheckpointINSType, _ = front.(*CheckpointINSType)
		}
		stage.OnAfterCheckpointINSTypeUpdateCallback.OnAfterUpdate(stage, checkpointinstype, frontCheckpointINSType)
	}
}

func (checkpointinstype *CheckpointINSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointINSTypeDeleteCallback != nil {
		var frontCheckpointINSType *CheckpointINSType
		if front != nil {
			frontCheckpointINSType, _ = front.(*CheckpointINSType)
		}
		stage.OnAfterCheckpointINSTypeDeleteCallback.OnAfterDelete(stage, checkpointinstype, frontCheckpointINSType)
	}
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCheckpointVORTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterCheckpointVORTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, checkpointvortimeslicepropertytype)
	}
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointVORTimeSlicePropertyTypeUpdateCallback != nil {
		var frontCheckpointVORTimeSlicePropertyType *CheckpointVORTimeSlicePropertyType
		if front != nil {
			frontCheckpointVORTimeSlicePropertyType, _ = front.(*CheckpointVORTimeSlicePropertyType)
		}
		stage.OnAfterCheckpointVORTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, checkpointvortimeslicepropertytype, frontCheckpointVORTimeSlicePropertyType)
	}
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointVORTimeSlicePropertyTypeDeleteCallback != nil {
		var frontCheckpointVORTimeSlicePropertyType *CheckpointVORTimeSlicePropertyType
		if front != nil {
			frontCheckpointVORTimeSlicePropertyType, _ = front.(*CheckpointVORTimeSlicePropertyType)
		}
		stage.OnAfterCheckpointVORTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, checkpointvortimeslicepropertytype, frontCheckpointVORTimeSlicePropertyType)
	}
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCheckpointVORTimeSliceTypeCreateCallback != nil {
		stage.OnAfterCheckpointVORTimeSliceTypeCreateCallback.OnAfterCreate(stage, checkpointvortimeslicetype)
	}
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointVORTimeSliceTypeUpdateCallback != nil {
		var frontCheckpointVORTimeSliceType *CheckpointVORTimeSliceType
		if front != nil {
			frontCheckpointVORTimeSliceType, _ = front.(*CheckpointVORTimeSliceType)
		}
		stage.OnAfterCheckpointVORTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, checkpointvortimeslicetype, frontCheckpointVORTimeSliceType)
	}
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointVORTimeSliceTypeDeleteCallback != nil {
		var frontCheckpointVORTimeSliceType *CheckpointVORTimeSliceType
		if front != nil {
			frontCheckpointVORTimeSliceType, _ = front.(*CheckpointVORTimeSliceType)
		}
		stage.OnAfterCheckpointVORTimeSliceTypeDeleteCallback.OnAfterDelete(stage, checkpointvortimeslicetype, frontCheckpointVORTimeSliceType)
	}
}

func (checkpointvortype *CheckpointVORType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCheckpointVORTypeCreateCallback != nil {
		stage.OnAfterCheckpointVORTypeCreateCallback.OnAfterCreate(stage, checkpointvortype)
	}
}

func (checkpointvortype *CheckpointVORType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointVORTypeUpdateCallback != nil {
		var frontCheckpointVORType *CheckpointVORType
		if front != nil {
			frontCheckpointVORType, _ = front.(*CheckpointVORType)
		}
		stage.OnAfterCheckpointVORTypeUpdateCallback.OnAfterUpdate(stage, checkpointvortype, frontCheckpointVORType)
	}
}

func (checkpointvortype *CheckpointVORType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckpointVORTypeDeleteCallback != nil {
		var frontCheckpointVORType *CheckpointVORType
		if front != nil {
			frontCheckpointVORType, _ = front.(*CheckpointVORType)
		}
		stage.OnAfterCheckpointVORTypeDeleteCallback.OnAfterDelete(stage, checkpointvortype, frontCheckpointVORType)
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

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDMETimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterDMETimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, dmetimeslicepropertytype)
	}
}

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMETimeSlicePropertyTypeUpdateCallback != nil {
		var frontDMETimeSlicePropertyType *DMETimeSlicePropertyType
		if front != nil {
			frontDMETimeSlicePropertyType, _ = front.(*DMETimeSlicePropertyType)
		}
		stage.OnAfterDMETimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, dmetimeslicepropertytype, frontDMETimeSlicePropertyType)
	}
}

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMETimeSlicePropertyTypeDeleteCallback != nil {
		var frontDMETimeSlicePropertyType *DMETimeSlicePropertyType
		if front != nil {
			frontDMETimeSlicePropertyType, _ = front.(*DMETimeSlicePropertyType)
		}
		stage.OnAfterDMETimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, dmetimeslicepropertytype, frontDMETimeSlicePropertyType)
	}
}

func (dmetimeslicetype *DMETimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDMETimeSliceTypeCreateCallback != nil {
		stage.OnAfterDMETimeSliceTypeCreateCallback.OnAfterCreate(stage, dmetimeslicetype)
	}
}

func (dmetimeslicetype *DMETimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMETimeSliceTypeUpdateCallback != nil {
		var frontDMETimeSliceType *DMETimeSliceType
		if front != nil {
			frontDMETimeSliceType, _ = front.(*DMETimeSliceType)
		}
		stage.OnAfterDMETimeSliceTypeUpdateCallback.OnAfterUpdate(stage, dmetimeslicetype, frontDMETimeSliceType)
	}
}

func (dmetimeslicetype *DMETimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMETimeSliceTypeDeleteCallback != nil {
		var frontDMETimeSliceType *DMETimeSliceType
		if front != nil {
			frontDMETimeSliceType, _ = front.(*DMETimeSliceType)
		}
		stage.OnAfterDMETimeSliceTypeDeleteCallback.OnAfterDelete(stage, dmetimeslicetype, frontDMETimeSliceType)
	}
}

func (dmetype *DMEType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDMETypeCreateCallback != nil {
		stage.OnAfterDMETypeCreateCallback.OnAfterCreate(stage, dmetype)
	}
}

func (dmetype *DMEType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMETypeUpdateCallback != nil {
		var frontDMEType *DMEType
		if front != nil {
			frontDMEType, _ = front.(*DMEType)
		}
		stage.OnAfterDMETypeUpdateCallback.OnAfterUpdate(stage, dmetype, frontDMEType)
	}
}

func (dmetype *DMEType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMETypeDeleteCallback != nil {
		var frontDMEType *DMEType
		if front != nil {
			frontDMEType, _ = front.(*DMEType)
		}
		stage.OnAfterDMETypeDeleteCallback.OnAfterDelete(stage, dmetype, frontDMEType)
	}
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDesignatedPointPropertyTypeCreateCallback != nil {
		stage.OnAfterDesignatedPointPropertyTypeCreateCallback.OnAfterCreate(stage, designatedpointpropertytype)
	}
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDesignatedPointPropertyTypeUpdateCallback != nil {
		var frontDesignatedPointPropertyType *DesignatedPointPropertyType
		if front != nil {
			frontDesignatedPointPropertyType, _ = front.(*DesignatedPointPropertyType)
		}
		stage.OnAfterDesignatedPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, designatedpointpropertytype, frontDesignatedPointPropertyType)
	}
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDesignatedPointPropertyTypeDeleteCallback != nil {
		var frontDesignatedPointPropertyType *DesignatedPointPropertyType
		if front != nil {
			frontDesignatedPointPropertyType, _ = front.(*DesignatedPointPropertyType)
		}
		stage.OnAfterDesignatedPointPropertyTypeDeleteCallback.OnAfterDelete(stage, designatedpointpropertytype, frontDesignatedPointPropertyType)
	}
}

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDesignatedPointTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterDesignatedPointTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, designatedpointtimeslicepropertytype)
	}
}

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDesignatedPointTimeSlicePropertyTypeUpdateCallback != nil {
		var frontDesignatedPointTimeSlicePropertyType *DesignatedPointTimeSlicePropertyType
		if front != nil {
			frontDesignatedPointTimeSlicePropertyType, _ = front.(*DesignatedPointTimeSlicePropertyType)
		}
		stage.OnAfterDesignatedPointTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, designatedpointtimeslicepropertytype, frontDesignatedPointTimeSlicePropertyType)
	}
}

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDesignatedPointTimeSlicePropertyTypeDeleteCallback != nil {
		var frontDesignatedPointTimeSlicePropertyType *DesignatedPointTimeSlicePropertyType
		if front != nil {
			frontDesignatedPointTimeSlicePropertyType, _ = front.(*DesignatedPointTimeSlicePropertyType)
		}
		stage.OnAfterDesignatedPointTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, designatedpointtimeslicepropertytype, frontDesignatedPointTimeSlicePropertyType)
	}
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDesignatedPointTimeSliceTypeCreateCallback != nil {
		stage.OnAfterDesignatedPointTimeSliceTypeCreateCallback.OnAfterCreate(stage, designatedpointtimeslicetype)
	}
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDesignatedPointTimeSliceTypeUpdateCallback != nil {
		var frontDesignatedPointTimeSliceType *DesignatedPointTimeSliceType
		if front != nil {
			frontDesignatedPointTimeSliceType, _ = front.(*DesignatedPointTimeSliceType)
		}
		stage.OnAfterDesignatedPointTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, designatedpointtimeslicetype, frontDesignatedPointTimeSliceType)
	}
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDesignatedPointTimeSliceTypeDeleteCallback != nil {
		var frontDesignatedPointTimeSliceType *DesignatedPointTimeSliceType
		if front != nil {
			frontDesignatedPointTimeSliceType, _ = front.(*DesignatedPointTimeSliceType)
		}
		stage.OnAfterDesignatedPointTimeSliceTypeDeleteCallback.OnAfterDelete(stage, designatedpointtimeslicetype, frontDesignatedPointTimeSliceType)
	}
}

func (designatedpointtype *DesignatedPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDesignatedPointTypeCreateCallback != nil {
		stage.OnAfterDesignatedPointTypeCreateCallback.OnAfterCreate(stage, designatedpointtype)
	}
}

func (designatedpointtype *DesignatedPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDesignatedPointTypeUpdateCallback != nil {
		var frontDesignatedPointType *DesignatedPointType
		if front != nil {
			frontDesignatedPointType, _ = front.(*DesignatedPointType)
		}
		stage.OnAfterDesignatedPointTypeUpdateCallback.OnAfterUpdate(stage, designatedpointtype, frontDesignatedPointType)
	}
}

func (designatedpointtype *DesignatedPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDesignatedPointTypeDeleteCallback != nil {
		var frontDesignatedPointType *DesignatedPointType
		if front != nil {
			frontDesignatedPointType, _ = front.(*DesignatedPointType)
		}
		stage.OnAfterDesignatedPointTypeDeleteCallback.OnAfterDelete(stage, designatedpointtype, frontDesignatedPointType)
	}
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionFinderTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterDirectionFinderTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, directionfindertimeslicepropertytype)
	}
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionFinderTimeSlicePropertyTypeUpdateCallback != nil {
		var frontDirectionFinderTimeSlicePropertyType *DirectionFinderTimeSlicePropertyType
		if front != nil {
			frontDirectionFinderTimeSlicePropertyType, _ = front.(*DirectionFinderTimeSlicePropertyType)
		}
		stage.OnAfterDirectionFinderTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, directionfindertimeslicepropertytype, frontDirectionFinderTimeSlicePropertyType)
	}
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionFinderTimeSlicePropertyTypeDeleteCallback != nil {
		var frontDirectionFinderTimeSlicePropertyType *DirectionFinderTimeSlicePropertyType
		if front != nil {
			frontDirectionFinderTimeSlicePropertyType, _ = front.(*DirectionFinderTimeSlicePropertyType)
		}
		stage.OnAfterDirectionFinderTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, directionfindertimeslicepropertytype, frontDirectionFinderTimeSlicePropertyType)
	}
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionFinderTimeSliceTypeCreateCallback != nil {
		stage.OnAfterDirectionFinderTimeSliceTypeCreateCallback.OnAfterCreate(stage, directionfindertimeslicetype)
	}
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionFinderTimeSliceTypeUpdateCallback != nil {
		var frontDirectionFinderTimeSliceType *DirectionFinderTimeSliceType
		if front != nil {
			frontDirectionFinderTimeSliceType, _ = front.(*DirectionFinderTimeSliceType)
		}
		stage.OnAfterDirectionFinderTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, directionfindertimeslicetype, frontDirectionFinderTimeSliceType)
	}
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionFinderTimeSliceTypeDeleteCallback != nil {
		var frontDirectionFinderTimeSliceType *DirectionFinderTimeSliceType
		if front != nil {
			frontDirectionFinderTimeSliceType, _ = front.(*DirectionFinderTimeSliceType)
		}
		stage.OnAfterDirectionFinderTimeSliceTypeDeleteCallback.OnAfterDelete(stage, directionfindertimeslicetype, frontDirectionFinderTimeSliceType)
	}
}

func (directionfindertype *DirectionFinderType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionFinderTypeCreateCallback != nil {
		stage.OnAfterDirectionFinderTypeCreateCallback.OnAfterCreate(stage, directionfindertype)
	}
}

func (directionfindertype *DirectionFinderType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionFinderTypeUpdateCallback != nil {
		var frontDirectionFinderType *DirectionFinderType
		if front != nil {
			frontDirectionFinderType, _ = front.(*DirectionFinderType)
		}
		stage.OnAfterDirectionFinderTypeUpdateCallback.OnAfterUpdate(stage, directionfindertype, frontDirectionFinderType)
	}
}

func (directionfindertype *DirectionFinderType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionFinderTypeDeleteCallback != nil {
		var frontDirectionFinderType *DirectionFinderType
		if front != nil {
			frontDirectionFinderType, _ = front.(*DirectionFinderType)
		}
		stage.OnAfterDirectionFinderTypeDeleteCallback.OnAfterDelete(stage, directionfindertype, frontDirectionFinderType)
	}
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDistanceIndicationTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterDistanceIndicationTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, distanceindicationtimeslicepropertytype)
	}
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceIndicationTimeSlicePropertyTypeUpdateCallback != nil {
		var frontDistanceIndicationTimeSlicePropertyType *DistanceIndicationTimeSlicePropertyType
		if front != nil {
			frontDistanceIndicationTimeSlicePropertyType, _ = front.(*DistanceIndicationTimeSlicePropertyType)
		}
		stage.OnAfterDistanceIndicationTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, distanceindicationtimeslicepropertytype, frontDistanceIndicationTimeSlicePropertyType)
	}
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceIndicationTimeSlicePropertyTypeDeleteCallback != nil {
		var frontDistanceIndicationTimeSlicePropertyType *DistanceIndicationTimeSlicePropertyType
		if front != nil {
			frontDistanceIndicationTimeSlicePropertyType, _ = front.(*DistanceIndicationTimeSlicePropertyType)
		}
		stage.OnAfterDistanceIndicationTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, distanceindicationtimeslicepropertytype, frontDistanceIndicationTimeSlicePropertyType)
	}
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDistanceIndicationTimeSliceTypeCreateCallback != nil {
		stage.OnAfterDistanceIndicationTimeSliceTypeCreateCallback.OnAfterCreate(stage, distanceindicationtimeslicetype)
	}
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceIndicationTimeSliceTypeUpdateCallback != nil {
		var frontDistanceIndicationTimeSliceType *DistanceIndicationTimeSliceType
		if front != nil {
			frontDistanceIndicationTimeSliceType, _ = front.(*DistanceIndicationTimeSliceType)
		}
		stage.OnAfterDistanceIndicationTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, distanceindicationtimeslicetype, frontDistanceIndicationTimeSliceType)
	}
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceIndicationTimeSliceTypeDeleteCallback != nil {
		var frontDistanceIndicationTimeSliceType *DistanceIndicationTimeSliceType
		if front != nil {
			frontDistanceIndicationTimeSliceType, _ = front.(*DistanceIndicationTimeSliceType)
		}
		stage.OnAfterDistanceIndicationTimeSliceTypeDeleteCallback.OnAfterDelete(stage, distanceindicationtimeslicetype, frontDistanceIndicationTimeSliceType)
	}
}

func (distanceindicationtype *DistanceIndicationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDistanceIndicationTypeCreateCallback != nil {
		stage.OnAfterDistanceIndicationTypeCreateCallback.OnAfterCreate(stage, distanceindicationtype)
	}
}

func (distanceindicationtype *DistanceIndicationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceIndicationTypeUpdateCallback != nil {
		var frontDistanceIndicationType *DistanceIndicationType
		if front != nil {
			frontDistanceIndicationType, _ = front.(*DistanceIndicationType)
		}
		stage.OnAfterDistanceIndicationTypeUpdateCallback.OnAfterUpdate(stage, distanceindicationtype, frontDistanceIndicationType)
	}
}

func (distanceindicationtype *DistanceIndicationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceIndicationTypeDeleteCallback != nil {
		var frontDistanceIndicationType *DistanceIndicationType
		if front != nil {
			frontDistanceIndicationType, _ = front.(*DistanceIndicationType)
		}
		stage.OnAfterDistanceIndicationTypeDeleteCallback.OnAfterDelete(stage, distanceindicationtype, frontDistanceIndicationType)
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

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevationTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterElevationTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, elevationtimeslicepropertytype)
	}
}

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevationTimeSlicePropertyTypeUpdateCallback != nil {
		var frontElevationTimeSlicePropertyType *ElevationTimeSlicePropertyType
		if front != nil {
			frontElevationTimeSlicePropertyType, _ = front.(*ElevationTimeSlicePropertyType)
		}
		stage.OnAfterElevationTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, elevationtimeslicepropertytype, frontElevationTimeSlicePropertyType)
	}
}

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevationTimeSlicePropertyTypeDeleteCallback != nil {
		var frontElevationTimeSlicePropertyType *ElevationTimeSlicePropertyType
		if front != nil {
			frontElevationTimeSlicePropertyType, _ = front.(*ElevationTimeSlicePropertyType)
		}
		stage.OnAfterElevationTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, elevationtimeslicepropertytype, frontElevationTimeSlicePropertyType)
	}
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevationTimeSliceTypeCreateCallback != nil {
		stage.OnAfterElevationTimeSliceTypeCreateCallback.OnAfterCreate(stage, elevationtimeslicetype)
	}
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevationTimeSliceTypeUpdateCallback != nil {
		var frontElevationTimeSliceType *ElevationTimeSliceType
		if front != nil {
			frontElevationTimeSliceType, _ = front.(*ElevationTimeSliceType)
		}
		stage.OnAfterElevationTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, elevationtimeslicetype, frontElevationTimeSliceType)
	}
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevationTimeSliceTypeDeleteCallback != nil {
		var frontElevationTimeSliceType *ElevationTimeSliceType
		if front != nil {
			frontElevationTimeSliceType, _ = front.(*ElevationTimeSliceType)
		}
		stage.OnAfterElevationTimeSliceTypeDeleteCallback.OnAfterDelete(stage, elevationtimeslicetype, frontElevationTimeSliceType)
	}
}

func (elevationtype *ElevationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevationTypeCreateCallback != nil {
		stage.OnAfterElevationTypeCreateCallback.OnAfterCreate(stage, elevationtype)
	}
}

func (elevationtype *ElevationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevationTypeUpdateCallback != nil {
		var frontElevationType *ElevationType
		if front != nil {
			frontElevationType, _ = front.(*ElevationType)
		}
		stage.OnAfterElevationTypeUpdateCallback.OnAfterUpdate(stage, elevationtype, frontElevationType)
	}
}

func (elevationtype *ElevationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevationTypeDeleteCallback != nil {
		var frontElevationType *ElevationType
		if front != nil {
			frontElevationType, _ = front.(*ElevationType)
		}
		stage.OnAfterElevationTypeDeleteCallback.OnAfterDelete(stage, elevationtype, frontElevationType)
	}
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGlidepathTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterGlidepathTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, glidepathtimeslicepropertytype)
	}
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlidepathTimeSlicePropertyTypeUpdateCallback != nil {
		var frontGlidepathTimeSlicePropertyType *GlidepathTimeSlicePropertyType
		if front != nil {
			frontGlidepathTimeSlicePropertyType, _ = front.(*GlidepathTimeSlicePropertyType)
		}
		stage.OnAfterGlidepathTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, glidepathtimeslicepropertytype, frontGlidepathTimeSlicePropertyType)
	}
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlidepathTimeSlicePropertyTypeDeleteCallback != nil {
		var frontGlidepathTimeSlicePropertyType *GlidepathTimeSlicePropertyType
		if front != nil {
			frontGlidepathTimeSlicePropertyType, _ = front.(*GlidepathTimeSlicePropertyType)
		}
		stage.OnAfterGlidepathTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, glidepathtimeslicepropertytype, frontGlidepathTimeSlicePropertyType)
	}
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGlidepathTimeSliceTypeCreateCallback != nil {
		stage.OnAfterGlidepathTimeSliceTypeCreateCallback.OnAfterCreate(stage, glidepathtimeslicetype)
	}
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlidepathTimeSliceTypeUpdateCallback != nil {
		var frontGlidepathTimeSliceType *GlidepathTimeSliceType
		if front != nil {
			frontGlidepathTimeSliceType, _ = front.(*GlidepathTimeSliceType)
		}
		stage.OnAfterGlidepathTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, glidepathtimeslicetype, frontGlidepathTimeSliceType)
	}
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlidepathTimeSliceTypeDeleteCallback != nil {
		var frontGlidepathTimeSliceType *GlidepathTimeSliceType
		if front != nil {
			frontGlidepathTimeSliceType, _ = front.(*GlidepathTimeSliceType)
		}
		stage.OnAfterGlidepathTimeSliceTypeDeleteCallback.OnAfterDelete(stage, glidepathtimeslicetype, frontGlidepathTimeSliceType)
	}
}

func (glidepathtype *GlidepathType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGlidepathTypeCreateCallback != nil {
		stage.OnAfterGlidepathTypeCreateCallback.OnAfterCreate(stage, glidepathtype)
	}
}

func (glidepathtype *GlidepathType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlidepathTypeUpdateCallback != nil {
		var frontGlidepathType *GlidepathType
		if front != nil {
			frontGlidepathType, _ = front.(*GlidepathType)
		}
		stage.OnAfterGlidepathTypeUpdateCallback.OnAfterUpdate(stage, glidepathtype, frontGlidepathType)
	}
}

func (glidepathtype *GlidepathType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGlidepathTypeDeleteCallback != nil {
		var frontGlidepathType *GlidepathType
		if front != nil {
			frontGlidepathType, _ = front.(*GlidepathType)
		}
		stage.OnAfterGlidepathTypeDeleteCallback.OnAfterDelete(stage, glidepathtype, frontGlidepathType)
	}
}

func (informationservicepropertytype *InformationServicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInformationServicePropertyTypeCreateCallback != nil {
		stage.OnAfterInformationServicePropertyTypeCreateCallback.OnAfterCreate(stage, informationservicepropertytype)
	}
}

func (informationservicepropertytype *InformationServicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInformationServicePropertyTypeUpdateCallback != nil {
		var frontInformationServicePropertyType *InformationServicePropertyType
		if front != nil {
			frontInformationServicePropertyType, _ = front.(*InformationServicePropertyType)
		}
		stage.OnAfterInformationServicePropertyTypeUpdateCallback.OnAfterUpdate(stage, informationservicepropertytype, frontInformationServicePropertyType)
	}
}

func (informationservicepropertytype *InformationServicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInformationServicePropertyTypeDeleteCallback != nil {
		var frontInformationServicePropertyType *InformationServicePropertyType
		if front != nil {
			frontInformationServicePropertyType, _ = front.(*InformationServicePropertyType)
		}
		stage.OnAfterInformationServicePropertyTypeDeleteCallback.OnAfterDelete(stage, informationservicepropertytype, frontInformationServicePropertyType)
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

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLocalizerTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterLocalizerTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, localizertimeslicepropertytype)
	}
}

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalizerTimeSlicePropertyTypeUpdateCallback != nil {
		var frontLocalizerTimeSlicePropertyType *LocalizerTimeSlicePropertyType
		if front != nil {
			frontLocalizerTimeSlicePropertyType, _ = front.(*LocalizerTimeSlicePropertyType)
		}
		stage.OnAfterLocalizerTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, localizertimeslicepropertytype, frontLocalizerTimeSlicePropertyType)
	}
}

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalizerTimeSlicePropertyTypeDeleteCallback != nil {
		var frontLocalizerTimeSlicePropertyType *LocalizerTimeSlicePropertyType
		if front != nil {
			frontLocalizerTimeSlicePropertyType, _ = front.(*LocalizerTimeSlicePropertyType)
		}
		stage.OnAfterLocalizerTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, localizertimeslicepropertytype, frontLocalizerTimeSlicePropertyType)
	}
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLocalizerTimeSliceTypeCreateCallback != nil {
		stage.OnAfterLocalizerTimeSliceTypeCreateCallback.OnAfterCreate(stage, localizertimeslicetype)
	}
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalizerTimeSliceTypeUpdateCallback != nil {
		var frontLocalizerTimeSliceType *LocalizerTimeSliceType
		if front != nil {
			frontLocalizerTimeSliceType, _ = front.(*LocalizerTimeSliceType)
		}
		stage.OnAfterLocalizerTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, localizertimeslicetype, frontLocalizerTimeSliceType)
	}
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalizerTimeSliceTypeDeleteCallback != nil {
		var frontLocalizerTimeSliceType *LocalizerTimeSliceType
		if front != nil {
			frontLocalizerTimeSliceType, _ = front.(*LocalizerTimeSliceType)
		}
		stage.OnAfterLocalizerTimeSliceTypeDeleteCallback.OnAfterDelete(stage, localizertimeslicetype, frontLocalizerTimeSliceType)
	}
}

func (localizertype *LocalizerType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLocalizerTypeCreateCallback != nil {
		stage.OnAfterLocalizerTypeCreateCallback.OnAfterCreate(stage, localizertype)
	}
}

func (localizertype *LocalizerType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalizerTypeUpdateCallback != nil {
		var frontLocalizerType *LocalizerType
		if front != nil {
			frontLocalizerType, _ = front.(*LocalizerType)
		}
		stage.OnAfterLocalizerTypeUpdateCallback.OnAfterUpdate(stage, localizertype, frontLocalizerType)
	}
}

func (localizertype *LocalizerType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalizerTypeDeleteCallback != nil {
		var frontLocalizerType *LocalizerType
		if front != nil {
			frontLocalizerType, _ = front.(*LocalizerType)
		}
		stage.OnAfterLocalizerTypeDeleteCallback.OnAfterDelete(stage, localizertype, frontLocalizerType)
	}
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkerBeaconTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterMarkerBeaconTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, markerbeacontimeslicepropertytype)
	}
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerBeaconTimeSlicePropertyTypeUpdateCallback != nil {
		var frontMarkerBeaconTimeSlicePropertyType *MarkerBeaconTimeSlicePropertyType
		if front != nil {
			frontMarkerBeaconTimeSlicePropertyType, _ = front.(*MarkerBeaconTimeSlicePropertyType)
		}
		stage.OnAfterMarkerBeaconTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, markerbeacontimeslicepropertytype, frontMarkerBeaconTimeSlicePropertyType)
	}
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerBeaconTimeSlicePropertyTypeDeleteCallback != nil {
		var frontMarkerBeaconTimeSlicePropertyType *MarkerBeaconTimeSlicePropertyType
		if front != nil {
			frontMarkerBeaconTimeSlicePropertyType, _ = front.(*MarkerBeaconTimeSlicePropertyType)
		}
		stage.OnAfterMarkerBeaconTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, markerbeacontimeslicepropertytype, frontMarkerBeaconTimeSlicePropertyType)
	}
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkerBeaconTimeSliceTypeCreateCallback != nil {
		stage.OnAfterMarkerBeaconTimeSliceTypeCreateCallback.OnAfterCreate(stage, markerbeacontimeslicetype)
	}
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerBeaconTimeSliceTypeUpdateCallback != nil {
		var frontMarkerBeaconTimeSliceType *MarkerBeaconTimeSliceType
		if front != nil {
			frontMarkerBeaconTimeSliceType, _ = front.(*MarkerBeaconTimeSliceType)
		}
		stage.OnAfterMarkerBeaconTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, markerbeacontimeslicetype, frontMarkerBeaconTimeSliceType)
	}
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerBeaconTimeSliceTypeDeleteCallback != nil {
		var frontMarkerBeaconTimeSliceType *MarkerBeaconTimeSliceType
		if front != nil {
			frontMarkerBeaconTimeSliceType, _ = front.(*MarkerBeaconTimeSliceType)
		}
		stage.OnAfterMarkerBeaconTimeSliceTypeDeleteCallback.OnAfterDelete(stage, markerbeacontimeslicetype, frontMarkerBeaconTimeSliceType)
	}
}

func (markerbeacontype *MarkerBeaconType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkerBeaconTypeCreateCallback != nil {
		stage.OnAfterMarkerBeaconTypeCreateCallback.OnAfterCreate(stage, markerbeacontype)
	}
}

func (markerbeacontype *MarkerBeaconType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerBeaconTypeUpdateCallback != nil {
		var frontMarkerBeaconType *MarkerBeaconType
		if front != nil {
			frontMarkerBeaconType, _ = front.(*MarkerBeaconType)
		}
		stage.OnAfterMarkerBeaconTypeUpdateCallback.OnAfterUpdate(stage, markerbeacontype, frontMarkerBeaconType)
	}
}

func (markerbeacontype *MarkerBeaconType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerBeaconTypeDeleteCallback != nil {
		var frontMarkerBeaconType *MarkerBeaconType
		if front != nil {
			frontMarkerBeaconType, _ = front.(*MarkerBeaconType)
		}
		stage.OnAfterMarkerBeaconTypeDeleteCallback.OnAfterDelete(stage, markerbeacontype, frontMarkerBeaconType)
	}
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNDBTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterNDBTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, ndbtimeslicepropertytype)
	}
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNDBTimeSlicePropertyTypeUpdateCallback != nil {
		var frontNDBTimeSlicePropertyType *NDBTimeSlicePropertyType
		if front != nil {
			frontNDBTimeSlicePropertyType, _ = front.(*NDBTimeSlicePropertyType)
		}
		stage.OnAfterNDBTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, ndbtimeslicepropertytype, frontNDBTimeSlicePropertyType)
	}
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNDBTimeSlicePropertyTypeDeleteCallback != nil {
		var frontNDBTimeSlicePropertyType *NDBTimeSlicePropertyType
		if front != nil {
			frontNDBTimeSlicePropertyType, _ = front.(*NDBTimeSlicePropertyType)
		}
		stage.OnAfterNDBTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, ndbtimeslicepropertytype, frontNDBTimeSlicePropertyType)
	}
}

func (ndbtimeslicetype *NDBTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNDBTimeSliceTypeCreateCallback != nil {
		stage.OnAfterNDBTimeSliceTypeCreateCallback.OnAfterCreate(stage, ndbtimeslicetype)
	}
}

func (ndbtimeslicetype *NDBTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNDBTimeSliceTypeUpdateCallback != nil {
		var frontNDBTimeSliceType *NDBTimeSliceType
		if front != nil {
			frontNDBTimeSliceType, _ = front.(*NDBTimeSliceType)
		}
		stage.OnAfterNDBTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, ndbtimeslicetype, frontNDBTimeSliceType)
	}
}

func (ndbtimeslicetype *NDBTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNDBTimeSliceTypeDeleteCallback != nil {
		var frontNDBTimeSliceType *NDBTimeSliceType
		if front != nil {
			frontNDBTimeSliceType, _ = front.(*NDBTimeSliceType)
		}
		stage.OnAfterNDBTimeSliceTypeDeleteCallback.OnAfterDelete(stage, ndbtimeslicetype, frontNDBTimeSliceType)
	}
}

func (ndbtype *NDBType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNDBTypeCreateCallback != nil {
		stage.OnAfterNDBTypeCreateCallback.OnAfterCreate(stage, ndbtype)
	}
}

func (ndbtype *NDBType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNDBTypeUpdateCallback != nil {
		var frontNDBType *NDBType
		if front != nil {
			frontNDBType, _ = front.(*NDBType)
		}
		stage.OnAfterNDBTypeUpdateCallback.OnAfterUpdate(stage, ndbtype, frontNDBType)
	}
}

func (ndbtype *NDBType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNDBTypeDeleteCallback != nil {
		var frontNDBType *NDBType
		if front != nil {
			frontNDBType, _ = front.(*NDBType)
		}
		stage.OnAfterNDBTypeDeleteCallback.OnAfterDelete(stage, ndbtype, frontNDBType)
	}
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidComponentPropertyTypeCreateCallback != nil {
		stage.OnAfterNavaidComponentPropertyTypeCreateCallback.OnAfterCreate(stage, navaidcomponentpropertytype)
	}
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidComponentPropertyTypeUpdateCallback != nil {
		var frontNavaidComponentPropertyType *NavaidComponentPropertyType
		if front != nil {
			frontNavaidComponentPropertyType, _ = front.(*NavaidComponentPropertyType)
		}
		stage.OnAfterNavaidComponentPropertyTypeUpdateCallback.OnAfterUpdate(stage, navaidcomponentpropertytype, frontNavaidComponentPropertyType)
	}
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidComponentPropertyTypeDeleteCallback != nil {
		var frontNavaidComponentPropertyType *NavaidComponentPropertyType
		if front != nil {
			frontNavaidComponentPropertyType, _ = front.(*NavaidComponentPropertyType)
		}
		stage.OnAfterNavaidComponentPropertyTypeDeleteCallback.OnAfterDelete(stage, navaidcomponentpropertytype, frontNavaidComponentPropertyType)
	}
}

func (navaidcomponenttype *NavaidComponentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidComponentTypeCreateCallback != nil {
		stage.OnAfterNavaidComponentTypeCreateCallback.OnAfterCreate(stage, navaidcomponenttype)
	}
}

func (navaidcomponenttype *NavaidComponentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidComponentTypeUpdateCallback != nil {
		var frontNavaidComponentType *NavaidComponentType
		if front != nil {
			frontNavaidComponentType, _ = front.(*NavaidComponentType)
		}
		stage.OnAfterNavaidComponentTypeUpdateCallback.OnAfterUpdate(stage, navaidcomponenttype, frontNavaidComponentType)
	}
}

func (navaidcomponenttype *NavaidComponentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidComponentTypeDeleteCallback != nil {
		var frontNavaidComponentType *NavaidComponentType
		if front != nil {
			frontNavaidComponentType, _ = front.(*NavaidComponentType)
		}
		stage.OnAfterNavaidComponentTypeDeleteCallback.OnAfterDelete(stage, navaidcomponenttype, frontNavaidComponentType)
	}
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidEquipmentMonitoringPropertyTypeCreateCallback != nil {
		stage.OnAfterNavaidEquipmentMonitoringPropertyTypeCreateCallback.OnAfterCreate(stage, navaidequipmentmonitoringpropertytype)
	}
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentMonitoringPropertyTypeUpdateCallback != nil {
		var frontNavaidEquipmentMonitoringPropertyType *NavaidEquipmentMonitoringPropertyType
		if front != nil {
			frontNavaidEquipmentMonitoringPropertyType, _ = front.(*NavaidEquipmentMonitoringPropertyType)
		}
		stage.OnAfterNavaidEquipmentMonitoringPropertyTypeUpdateCallback.OnAfterUpdate(stage, navaidequipmentmonitoringpropertytype, frontNavaidEquipmentMonitoringPropertyType)
	}
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentMonitoringPropertyTypeDeleteCallback != nil {
		var frontNavaidEquipmentMonitoringPropertyType *NavaidEquipmentMonitoringPropertyType
		if front != nil {
			frontNavaidEquipmentMonitoringPropertyType, _ = front.(*NavaidEquipmentMonitoringPropertyType)
		}
		stage.OnAfterNavaidEquipmentMonitoringPropertyTypeDeleteCallback.OnAfterDelete(stage, navaidequipmentmonitoringpropertytype, frontNavaidEquipmentMonitoringPropertyType)
	}
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidEquipmentMonitoringTypeCreateCallback != nil {
		stage.OnAfterNavaidEquipmentMonitoringTypeCreateCallback.OnAfterCreate(stage, navaidequipmentmonitoringtype)
	}
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentMonitoringTypeUpdateCallback != nil {
		var frontNavaidEquipmentMonitoringType *NavaidEquipmentMonitoringType
		if front != nil {
			frontNavaidEquipmentMonitoringType, _ = front.(*NavaidEquipmentMonitoringType)
		}
		stage.OnAfterNavaidEquipmentMonitoringTypeUpdateCallback.OnAfterUpdate(stage, navaidequipmentmonitoringtype, frontNavaidEquipmentMonitoringType)
	}
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentMonitoringTypeDeleteCallback != nil {
		var frontNavaidEquipmentMonitoringType *NavaidEquipmentMonitoringType
		if front != nil {
			frontNavaidEquipmentMonitoringType, _ = front.(*NavaidEquipmentMonitoringType)
		}
		stage.OnAfterNavaidEquipmentMonitoringTypeDeleteCallback.OnAfterDelete(stage, navaidequipmentmonitoringtype, frontNavaidEquipmentMonitoringType)
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

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidOperationalStatusPropertyTypeCreateCallback != nil {
		stage.OnAfterNavaidOperationalStatusPropertyTypeCreateCallback.OnAfterCreate(stage, navaidoperationalstatuspropertytype)
	}
}

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidOperationalStatusPropertyTypeUpdateCallback != nil {
		var frontNavaidOperationalStatusPropertyType *NavaidOperationalStatusPropertyType
		if front != nil {
			frontNavaidOperationalStatusPropertyType, _ = front.(*NavaidOperationalStatusPropertyType)
		}
		stage.OnAfterNavaidOperationalStatusPropertyTypeUpdateCallback.OnAfterUpdate(stage, navaidoperationalstatuspropertytype, frontNavaidOperationalStatusPropertyType)
	}
}

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidOperationalStatusPropertyTypeDeleteCallback != nil {
		var frontNavaidOperationalStatusPropertyType *NavaidOperationalStatusPropertyType
		if front != nil {
			frontNavaidOperationalStatusPropertyType, _ = front.(*NavaidOperationalStatusPropertyType)
		}
		stage.OnAfterNavaidOperationalStatusPropertyTypeDeleteCallback.OnAfterDelete(stage, navaidoperationalstatuspropertytype, frontNavaidOperationalStatusPropertyType)
	}
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidOperationalStatusTypeCreateCallback != nil {
		stage.OnAfterNavaidOperationalStatusTypeCreateCallback.OnAfterCreate(stage, navaidoperationalstatustype)
	}
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidOperationalStatusTypeUpdateCallback != nil {
		var frontNavaidOperationalStatusType *NavaidOperationalStatusType
		if front != nil {
			frontNavaidOperationalStatusType, _ = front.(*NavaidOperationalStatusType)
		}
		stage.OnAfterNavaidOperationalStatusTypeUpdateCallback.OnAfterUpdate(stage, navaidoperationalstatustype, frontNavaidOperationalStatusType)
	}
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidOperationalStatusTypeDeleteCallback != nil {
		var frontNavaidOperationalStatusType *NavaidOperationalStatusType
		if front != nil {
			frontNavaidOperationalStatusType, _ = front.(*NavaidOperationalStatusType)
		}
		stage.OnAfterNavaidOperationalStatusTypeDeleteCallback.OnAfterDelete(stage, navaidoperationalstatustype, frontNavaidOperationalStatusType)
	}
}

func (navaidpropertytype *NavaidPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidPropertyTypeCreateCallback != nil {
		stage.OnAfterNavaidPropertyTypeCreateCallback.OnAfterCreate(stage, navaidpropertytype)
	}
}

func (navaidpropertytype *NavaidPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidPropertyTypeUpdateCallback != nil {
		var frontNavaidPropertyType *NavaidPropertyType
		if front != nil {
			frontNavaidPropertyType, _ = front.(*NavaidPropertyType)
		}
		stage.OnAfterNavaidPropertyTypeUpdateCallback.OnAfterUpdate(stage, navaidpropertytype, frontNavaidPropertyType)
	}
}

func (navaidpropertytype *NavaidPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidPropertyTypeDeleteCallback != nil {
		var frontNavaidPropertyType *NavaidPropertyType
		if front != nil {
			frontNavaidPropertyType, _ = front.(*NavaidPropertyType)
		}
		stage.OnAfterNavaidPropertyTypeDeleteCallback.OnAfterDelete(stage, navaidpropertytype, frontNavaidPropertyType)
	}
}

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterNavaidTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, navaidtimeslicepropertytype)
	}
}

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidTimeSlicePropertyTypeUpdateCallback != nil {
		var frontNavaidTimeSlicePropertyType *NavaidTimeSlicePropertyType
		if front != nil {
			frontNavaidTimeSlicePropertyType, _ = front.(*NavaidTimeSlicePropertyType)
		}
		stage.OnAfterNavaidTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, navaidtimeslicepropertytype, frontNavaidTimeSlicePropertyType)
	}
}

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidTimeSlicePropertyTypeDeleteCallback != nil {
		var frontNavaidTimeSlicePropertyType *NavaidTimeSlicePropertyType
		if front != nil {
			frontNavaidTimeSlicePropertyType, _ = front.(*NavaidTimeSlicePropertyType)
		}
		stage.OnAfterNavaidTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, navaidtimeslicepropertytype, frontNavaidTimeSlicePropertyType)
	}
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidTimeSliceTypeCreateCallback != nil {
		stage.OnAfterNavaidTimeSliceTypeCreateCallback.OnAfterCreate(stage, navaidtimeslicetype)
	}
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidTimeSliceTypeUpdateCallback != nil {
		var frontNavaidTimeSliceType *NavaidTimeSliceType
		if front != nil {
			frontNavaidTimeSliceType, _ = front.(*NavaidTimeSliceType)
		}
		stage.OnAfterNavaidTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, navaidtimeslicetype, frontNavaidTimeSliceType)
	}
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidTimeSliceTypeDeleteCallback != nil {
		var frontNavaidTimeSliceType *NavaidTimeSliceType
		if front != nil {
			frontNavaidTimeSliceType, _ = front.(*NavaidTimeSliceType)
		}
		stage.OnAfterNavaidTimeSliceTypeDeleteCallback.OnAfterDelete(stage, navaidtimeslicetype, frontNavaidTimeSliceType)
	}
}

func (navaidtype *NavaidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidTypeCreateCallback != nil {
		stage.OnAfterNavaidTypeCreateCallback.OnAfterCreate(stage, navaidtype)
	}
}

func (navaidtype *NavaidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidTypeUpdateCallback != nil {
		var frontNavaidType *NavaidType
		if front != nil {
			frontNavaidType, _ = front.(*NavaidType)
		}
		stage.OnAfterNavaidTypeUpdateCallback.OnAfterUpdate(stage, navaidtype, frontNavaidType)
	}
}

func (navaidtype *NavaidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidTypeDeleteCallback != nil {
		var frontNavaidType *NavaidType
		if front != nil {
			frontNavaidType, _ = front.(*NavaidType)
		}
		stage.OnAfterNavaidTypeDeleteCallback.OnAfterDelete(stage, navaidtype, frontNavaidType)
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

func (pointpropertytype *PointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointPropertyTypeCreateCallback != nil {
		stage.OnAfterPointPropertyTypeCreateCallback.OnAfterCreate(stage, pointpropertytype)
	}
}

func (pointpropertytype *PointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointPropertyTypeUpdateCallback != nil {
		var frontPointPropertyType *PointPropertyType
		if front != nil {
			frontPointPropertyType, _ = front.(*PointPropertyType)
		}
		stage.OnAfterPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, pointpropertytype, frontPointPropertyType)
	}
}

func (pointpropertytype *PointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointPropertyTypeDeleteCallback != nil {
		var frontPointPropertyType *PointPropertyType
		if front != nil {
			frontPointPropertyType, _ = front.(*PointPropertyType)
		}
		stage.OnAfterPointPropertyTypeDeleteCallback.OnAfterDelete(stage, pointpropertytype, frontPointPropertyType)
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

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrecisionApproachRadarTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterPrecisionApproachRadarTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, precisionapproachradartimeslicepropertytype)
	}
}

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrecisionApproachRadarTimeSlicePropertyTypeUpdateCallback != nil {
		var frontPrecisionApproachRadarTimeSlicePropertyType *PrecisionApproachRadarTimeSlicePropertyType
		if front != nil {
			frontPrecisionApproachRadarTimeSlicePropertyType, _ = front.(*PrecisionApproachRadarTimeSlicePropertyType)
		}
		stage.OnAfterPrecisionApproachRadarTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, precisionapproachradartimeslicepropertytype, frontPrecisionApproachRadarTimeSlicePropertyType)
	}
}

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrecisionApproachRadarTimeSlicePropertyTypeDeleteCallback != nil {
		var frontPrecisionApproachRadarTimeSlicePropertyType *PrecisionApproachRadarTimeSlicePropertyType
		if front != nil {
			frontPrecisionApproachRadarTimeSlicePropertyType, _ = front.(*PrecisionApproachRadarTimeSlicePropertyType)
		}
		stage.OnAfterPrecisionApproachRadarTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, precisionapproachradartimeslicepropertytype, frontPrecisionApproachRadarTimeSlicePropertyType)
	}
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrecisionApproachRadarTimeSliceTypeCreateCallback != nil {
		stage.OnAfterPrecisionApproachRadarTimeSliceTypeCreateCallback.OnAfterCreate(stage, precisionapproachradartimeslicetype)
	}
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrecisionApproachRadarTimeSliceTypeUpdateCallback != nil {
		var frontPrecisionApproachRadarTimeSliceType *PrecisionApproachRadarTimeSliceType
		if front != nil {
			frontPrecisionApproachRadarTimeSliceType, _ = front.(*PrecisionApproachRadarTimeSliceType)
		}
		stage.OnAfterPrecisionApproachRadarTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, precisionapproachradartimeslicetype, frontPrecisionApproachRadarTimeSliceType)
	}
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrecisionApproachRadarTimeSliceTypeDeleteCallback != nil {
		var frontPrecisionApproachRadarTimeSliceType *PrecisionApproachRadarTimeSliceType
		if front != nil {
			frontPrecisionApproachRadarTimeSliceType, _ = front.(*PrecisionApproachRadarTimeSliceType)
		}
		stage.OnAfterPrecisionApproachRadarTimeSliceTypeDeleteCallback.OnAfterDelete(stage, precisionapproachradartimeslicetype, frontPrecisionApproachRadarTimeSliceType)
	}
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrecisionApproachRadarTypeCreateCallback != nil {
		stage.OnAfterPrecisionApproachRadarTypeCreateCallback.OnAfterCreate(stage, precisionapproachradartype)
	}
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrecisionApproachRadarTypeUpdateCallback != nil {
		var frontPrecisionApproachRadarType *PrecisionApproachRadarType
		if front != nil {
			frontPrecisionApproachRadarType, _ = front.(*PrecisionApproachRadarType)
		}
		stage.OnAfterPrecisionApproachRadarTypeUpdateCallback.OnAfterUpdate(stage, precisionapproachradartype, frontPrecisionApproachRadarType)
	}
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrecisionApproachRadarTypeDeleteCallback != nil {
		var frontPrecisionApproachRadarType *PrecisionApproachRadarType
		if front != nil {
			frontPrecisionApproachRadarType, _ = front.(*PrecisionApproachRadarType)
		}
		stage.OnAfterPrecisionApproachRadarTypeDeleteCallback.OnAfterDelete(stage, precisionapproachradartype, frontPrecisionApproachRadarType)
	}
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrimarySurveillanceRadarTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterPrimarySurveillanceRadarTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, primarysurveillanceradartimeslicepropertytype)
	}
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimarySurveillanceRadarTimeSlicePropertyTypeUpdateCallback != nil {
		var frontPrimarySurveillanceRadarTimeSlicePropertyType *PrimarySurveillanceRadarTimeSlicePropertyType
		if front != nil {
			frontPrimarySurveillanceRadarTimeSlicePropertyType, _ = front.(*PrimarySurveillanceRadarTimeSlicePropertyType)
		}
		stage.OnAfterPrimarySurveillanceRadarTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, primarysurveillanceradartimeslicepropertytype, frontPrimarySurveillanceRadarTimeSlicePropertyType)
	}
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimarySurveillanceRadarTimeSlicePropertyTypeDeleteCallback != nil {
		var frontPrimarySurveillanceRadarTimeSlicePropertyType *PrimarySurveillanceRadarTimeSlicePropertyType
		if front != nil {
			frontPrimarySurveillanceRadarTimeSlicePropertyType, _ = front.(*PrimarySurveillanceRadarTimeSlicePropertyType)
		}
		stage.OnAfterPrimarySurveillanceRadarTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, primarysurveillanceradartimeslicepropertytype, frontPrimarySurveillanceRadarTimeSlicePropertyType)
	}
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrimarySurveillanceRadarTimeSliceTypeCreateCallback != nil {
		stage.OnAfterPrimarySurveillanceRadarTimeSliceTypeCreateCallback.OnAfterCreate(stage, primarysurveillanceradartimeslicetype)
	}
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimarySurveillanceRadarTimeSliceTypeUpdateCallback != nil {
		var frontPrimarySurveillanceRadarTimeSliceType *PrimarySurveillanceRadarTimeSliceType
		if front != nil {
			frontPrimarySurveillanceRadarTimeSliceType, _ = front.(*PrimarySurveillanceRadarTimeSliceType)
		}
		stage.OnAfterPrimarySurveillanceRadarTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, primarysurveillanceradartimeslicetype, frontPrimarySurveillanceRadarTimeSliceType)
	}
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimarySurveillanceRadarTimeSliceTypeDeleteCallback != nil {
		var frontPrimarySurveillanceRadarTimeSliceType *PrimarySurveillanceRadarTimeSliceType
		if front != nil {
			frontPrimarySurveillanceRadarTimeSliceType, _ = front.(*PrimarySurveillanceRadarTimeSliceType)
		}
		stage.OnAfterPrimarySurveillanceRadarTimeSliceTypeDeleteCallback.OnAfterDelete(stage, primarysurveillanceradartimeslicetype, frontPrimarySurveillanceRadarTimeSliceType)
	}
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrimarySurveillanceRadarTypeCreateCallback != nil {
		stage.OnAfterPrimarySurveillanceRadarTypeCreateCallback.OnAfterCreate(stage, primarysurveillanceradartype)
	}
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimarySurveillanceRadarTypeUpdateCallback != nil {
		var frontPrimarySurveillanceRadarType *PrimarySurveillanceRadarType
		if front != nil {
			frontPrimarySurveillanceRadarType, _ = front.(*PrimarySurveillanceRadarType)
		}
		stage.OnAfterPrimarySurveillanceRadarTypeUpdateCallback.OnAfterUpdate(stage, primarysurveillanceradartype, frontPrimarySurveillanceRadarType)
	}
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimarySurveillanceRadarTypeDeleteCallback != nil {
		var frontPrimarySurveillanceRadarType *PrimarySurveillanceRadarType
		if front != nil {
			frontPrimarySurveillanceRadarType, _ = front.(*PrimarySurveillanceRadarType)
		}
		stage.OnAfterPrimarySurveillanceRadarTypeDeleteCallback.OnAfterDelete(stage, primarysurveillanceradartype, frontPrimarySurveillanceRadarType)
	}
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadarComponentPropertyTypeCreateCallback != nil {
		stage.OnAfterRadarComponentPropertyTypeCreateCallback.OnAfterCreate(stage, radarcomponentpropertytype)
	}
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarComponentPropertyTypeUpdateCallback != nil {
		var frontRadarComponentPropertyType *RadarComponentPropertyType
		if front != nil {
			frontRadarComponentPropertyType, _ = front.(*RadarComponentPropertyType)
		}
		stage.OnAfterRadarComponentPropertyTypeUpdateCallback.OnAfterUpdate(stage, radarcomponentpropertytype, frontRadarComponentPropertyType)
	}
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarComponentPropertyTypeDeleteCallback != nil {
		var frontRadarComponentPropertyType *RadarComponentPropertyType
		if front != nil {
			frontRadarComponentPropertyType, _ = front.(*RadarComponentPropertyType)
		}
		stage.OnAfterRadarComponentPropertyTypeDeleteCallback.OnAfterDelete(stage, radarcomponentpropertytype, frontRadarComponentPropertyType)
	}
}

func (radarcomponenttype *RadarComponentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadarComponentTypeCreateCallback != nil {
		stage.OnAfterRadarComponentTypeCreateCallback.OnAfterCreate(stage, radarcomponenttype)
	}
}

func (radarcomponenttype *RadarComponentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarComponentTypeUpdateCallback != nil {
		var frontRadarComponentType *RadarComponentType
		if front != nil {
			frontRadarComponentType, _ = front.(*RadarComponentType)
		}
		stage.OnAfterRadarComponentTypeUpdateCallback.OnAfterUpdate(stage, radarcomponenttype, frontRadarComponentType)
	}
}

func (radarcomponenttype *RadarComponentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarComponentTypeDeleteCallback != nil {
		var frontRadarComponentType *RadarComponentType
		if front != nil {
			frontRadarComponentType, _ = front.(*RadarComponentType)
		}
		stage.OnAfterRadarComponentTypeDeleteCallback.OnAfterDelete(stage, radarcomponenttype, frontRadarComponentType)
	}
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadarEquipmentPropertyTypeCreateCallback != nil {
		stage.OnAfterRadarEquipmentPropertyTypeCreateCallback.OnAfterCreate(stage, radarequipmentpropertytype)
	}
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarEquipmentPropertyTypeUpdateCallback != nil {
		var frontRadarEquipmentPropertyType *RadarEquipmentPropertyType
		if front != nil {
			frontRadarEquipmentPropertyType, _ = front.(*RadarEquipmentPropertyType)
		}
		stage.OnAfterRadarEquipmentPropertyTypeUpdateCallback.OnAfterUpdate(stage, radarequipmentpropertytype, frontRadarEquipmentPropertyType)
	}
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarEquipmentPropertyTypeDeleteCallback != nil {
		var frontRadarEquipmentPropertyType *RadarEquipmentPropertyType
		if front != nil {
			frontRadarEquipmentPropertyType, _ = front.(*RadarEquipmentPropertyType)
		}
		stage.OnAfterRadarEquipmentPropertyTypeDeleteCallback.OnAfterDelete(stage, radarequipmentpropertytype, frontRadarEquipmentPropertyType)
	}
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadarSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRadarSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, radarsystemtimeslicepropertytype)
	}
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRadarSystemTimeSlicePropertyType *RadarSystemTimeSlicePropertyType
		if front != nil {
			frontRadarSystemTimeSlicePropertyType, _ = front.(*RadarSystemTimeSlicePropertyType)
		}
		stage.OnAfterRadarSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, radarsystemtimeslicepropertytype, frontRadarSystemTimeSlicePropertyType)
	}
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRadarSystemTimeSlicePropertyType *RadarSystemTimeSlicePropertyType
		if front != nil {
			frontRadarSystemTimeSlicePropertyType, _ = front.(*RadarSystemTimeSlicePropertyType)
		}
		stage.OnAfterRadarSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, radarsystemtimeslicepropertytype, frontRadarSystemTimeSlicePropertyType)
	}
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadarSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRadarSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, radarsystemtimeslicetype)
	}
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarSystemTimeSliceTypeUpdateCallback != nil {
		var frontRadarSystemTimeSliceType *RadarSystemTimeSliceType
		if front != nil {
			frontRadarSystemTimeSliceType, _ = front.(*RadarSystemTimeSliceType)
		}
		stage.OnAfterRadarSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, radarsystemtimeslicetype, frontRadarSystemTimeSliceType)
	}
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarSystemTimeSliceTypeDeleteCallback != nil {
		var frontRadarSystemTimeSliceType *RadarSystemTimeSliceType
		if front != nil {
			frontRadarSystemTimeSliceType, _ = front.(*RadarSystemTimeSliceType)
		}
		stage.OnAfterRadarSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, radarsystemtimeslicetype, frontRadarSystemTimeSliceType)
	}
}

func (radarsystemtype *RadarSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadarSystemTypeCreateCallback != nil {
		stage.OnAfterRadarSystemTypeCreateCallback.OnAfterCreate(stage, radarsystemtype)
	}
}

func (radarsystemtype *RadarSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarSystemTypeUpdateCallback != nil {
		var frontRadarSystemType *RadarSystemType
		if front != nil {
			frontRadarSystemType, _ = front.(*RadarSystemType)
		}
		stage.OnAfterRadarSystemTypeUpdateCallback.OnAfterUpdate(stage, radarsystemtype, frontRadarSystemType)
	}
}

func (radarsystemtype *RadarSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarSystemTypeDeleteCallback != nil {
		var frontRadarSystemType *RadarSystemType
		if front != nil {
			frontRadarSystemType, _ = front.(*RadarSystemType)
		}
		stage.OnAfterRadarSystemTypeDeleteCallback.OnAfterDelete(stage, radarsystemtype, frontRadarSystemType)
	}
}

func (reflectorpropertytype *ReflectorPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterReflectorPropertyTypeCreateCallback != nil {
		stage.OnAfterReflectorPropertyTypeCreateCallback.OnAfterCreate(stage, reflectorpropertytype)
	}
}

func (reflectorpropertytype *ReflectorPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReflectorPropertyTypeUpdateCallback != nil {
		var frontReflectorPropertyType *ReflectorPropertyType
		if front != nil {
			frontReflectorPropertyType, _ = front.(*ReflectorPropertyType)
		}
		stage.OnAfterReflectorPropertyTypeUpdateCallback.OnAfterUpdate(stage, reflectorpropertytype, frontReflectorPropertyType)
	}
}

func (reflectorpropertytype *ReflectorPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReflectorPropertyTypeDeleteCallback != nil {
		var frontReflectorPropertyType *ReflectorPropertyType
		if front != nil {
			frontReflectorPropertyType, _ = front.(*ReflectorPropertyType)
		}
		stage.OnAfterReflectorPropertyTypeDeleteCallback.OnAfterDelete(stage, reflectorpropertytype, frontReflectorPropertyType)
	}
}

func (reflectortype *ReflectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterReflectorTypeCreateCallback != nil {
		stage.OnAfterReflectorTypeCreateCallback.OnAfterCreate(stage, reflectortype)
	}
}

func (reflectortype *ReflectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReflectorTypeUpdateCallback != nil {
		var frontReflectorType *ReflectorType
		if front != nil {
			frontReflectorType, _ = front.(*ReflectorType)
		}
		stage.OnAfterReflectorTypeUpdateCallback.OnAfterUpdate(stage, reflectortype, frontReflectorType)
	}
}

func (reflectortype *ReflectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReflectorTypeDeleteCallback != nil {
		var frontReflectorType *ReflectorType
		if front != nil {
			frontReflectorType, _ = front.(*ReflectorType)
		}
		stage.OnAfterReflectorTypeDeleteCallback.OnAfterDelete(stage, reflectortype, frontReflectorType)
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

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSDFTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSDFTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, sdftimeslicepropertytype)
	}
}

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSDFTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSDFTimeSlicePropertyType *SDFTimeSlicePropertyType
		if front != nil {
			frontSDFTimeSlicePropertyType, _ = front.(*SDFTimeSlicePropertyType)
		}
		stage.OnAfterSDFTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, sdftimeslicepropertytype, frontSDFTimeSlicePropertyType)
	}
}

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSDFTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSDFTimeSlicePropertyType *SDFTimeSlicePropertyType
		if front != nil {
			frontSDFTimeSlicePropertyType, _ = front.(*SDFTimeSlicePropertyType)
		}
		stage.OnAfterSDFTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, sdftimeslicepropertytype, frontSDFTimeSlicePropertyType)
	}
}

func (sdftimeslicetype *SDFTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSDFTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSDFTimeSliceTypeCreateCallback.OnAfterCreate(stage, sdftimeslicetype)
	}
}

func (sdftimeslicetype *SDFTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSDFTimeSliceTypeUpdateCallback != nil {
		var frontSDFTimeSliceType *SDFTimeSliceType
		if front != nil {
			frontSDFTimeSliceType, _ = front.(*SDFTimeSliceType)
		}
		stage.OnAfterSDFTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, sdftimeslicetype, frontSDFTimeSliceType)
	}
}

func (sdftimeslicetype *SDFTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSDFTimeSliceTypeDeleteCallback != nil {
		var frontSDFTimeSliceType *SDFTimeSliceType
		if front != nil {
			frontSDFTimeSliceType, _ = front.(*SDFTimeSliceType)
		}
		stage.OnAfterSDFTimeSliceTypeDeleteCallback.OnAfterDelete(stage, sdftimeslicetype, frontSDFTimeSliceType)
	}
}

func (sdftype *SDFType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSDFTypeCreateCallback != nil {
		stage.OnAfterSDFTypeCreateCallback.OnAfterCreate(stage, sdftype)
	}
}

func (sdftype *SDFType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSDFTypeUpdateCallback != nil {
		var frontSDFType *SDFType
		if front != nil {
			frontSDFType, _ = front.(*SDFType)
		}
		stage.OnAfterSDFTypeUpdateCallback.OnAfterUpdate(stage, sdftype, frontSDFType)
	}
}

func (sdftype *SDFType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSDFTypeDeleteCallback != nil {
		var frontSDFType *SDFType
		if front != nil {
			frontSDFType, _ = front.(*SDFType)
		}
		stage.OnAfterSDFTypeDeleteCallback.OnAfterDelete(stage, sdftype, frontSDFType)
	}
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSecondarySurveillanceRadarTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSecondarySurveillanceRadarTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, secondarysurveillanceradartimeslicepropertytype)
	}
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondarySurveillanceRadarTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSecondarySurveillanceRadarTimeSlicePropertyType *SecondarySurveillanceRadarTimeSlicePropertyType
		if front != nil {
			frontSecondarySurveillanceRadarTimeSlicePropertyType, _ = front.(*SecondarySurveillanceRadarTimeSlicePropertyType)
		}
		stage.OnAfterSecondarySurveillanceRadarTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, secondarysurveillanceradartimeslicepropertytype, frontSecondarySurveillanceRadarTimeSlicePropertyType)
	}
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondarySurveillanceRadarTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSecondarySurveillanceRadarTimeSlicePropertyType *SecondarySurveillanceRadarTimeSlicePropertyType
		if front != nil {
			frontSecondarySurveillanceRadarTimeSlicePropertyType, _ = front.(*SecondarySurveillanceRadarTimeSlicePropertyType)
		}
		stage.OnAfterSecondarySurveillanceRadarTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, secondarysurveillanceradartimeslicepropertytype, frontSecondarySurveillanceRadarTimeSlicePropertyType)
	}
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSecondarySurveillanceRadarTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSecondarySurveillanceRadarTimeSliceTypeCreateCallback.OnAfterCreate(stage, secondarysurveillanceradartimeslicetype)
	}
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondarySurveillanceRadarTimeSliceTypeUpdateCallback != nil {
		var frontSecondarySurveillanceRadarTimeSliceType *SecondarySurveillanceRadarTimeSliceType
		if front != nil {
			frontSecondarySurveillanceRadarTimeSliceType, _ = front.(*SecondarySurveillanceRadarTimeSliceType)
		}
		stage.OnAfterSecondarySurveillanceRadarTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, secondarysurveillanceradartimeslicetype, frontSecondarySurveillanceRadarTimeSliceType)
	}
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondarySurveillanceRadarTimeSliceTypeDeleteCallback != nil {
		var frontSecondarySurveillanceRadarTimeSliceType *SecondarySurveillanceRadarTimeSliceType
		if front != nil {
			frontSecondarySurveillanceRadarTimeSliceType, _ = front.(*SecondarySurveillanceRadarTimeSliceType)
		}
		stage.OnAfterSecondarySurveillanceRadarTimeSliceTypeDeleteCallback.OnAfterDelete(stage, secondarysurveillanceradartimeslicetype, frontSecondarySurveillanceRadarTimeSliceType)
	}
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSecondarySurveillanceRadarTypeCreateCallback != nil {
		stage.OnAfterSecondarySurveillanceRadarTypeCreateCallback.OnAfterCreate(stage, secondarysurveillanceradartype)
	}
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondarySurveillanceRadarTypeUpdateCallback != nil {
		var frontSecondarySurveillanceRadarType *SecondarySurveillanceRadarType
		if front != nil {
			frontSecondarySurveillanceRadarType, _ = front.(*SecondarySurveillanceRadarType)
		}
		stage.OnAfterSecondarySurveillanceRadarTypeUpdateCallback.OnAfterUpdate(stage, secondarysurveillanceradartype, frontSecondarySurveillanceRadarType)
	}
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondarySurveillanceRadarTypeDeleteCallback != nil {
		var frontSecondarySurveillanceRadarType *SecondarySurveillanceRadarType
		if front != nil {
			frontSecondarySurveillanceRadarType, _ = front.(*SecondarySurveillanceRadarType)
		}
		stage.OnAfterSecondarySurveillanceRadarTypeDeleteCallback.OnAfterDelete(stage, secondarysurveillanceradartype, frontSecondarySurveillanceRadarType)
	}
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSignificantPointInAirspaceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSignificantPointInAirspaceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, significantpointinairspacetimeslicepropertytype)
	}
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSignificantPointInAirspaceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSignificantPointInAirspaceTimeSlicePropertyType *SignificantPointInAirspaceTimeSlicePropertyType
		if front != nil {
			frontSignificantPointInAirspaceTimeSlicePropertyType, _ = front.(*SignificantPointInAirspaceTimeSlicePropertyType)
		}
		stage.OnAfterSignificantPointInAirspaceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, significantpointinairspacetimeslicepropertytype, frontSignificantPointInAirspaceTimeSlicePropertyType)
	}
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSignificantPointInAirspaceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSignificantPointInAirspaceTimeSlicePropertyType *SignificantPointInAirspaceTimeSlicePropertyType
		if front != nil {
			frontSignificantPointInAirspaceTimeSlicePropertyType, _ = front.(*SignificantPointInAirspaceTimeSlicePropertyType)
		}
		stage.OnAfterSignificantPointInAirspaceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, significantpointinairspacetimeslicepropertytype, frontSignificantPointInAirspaceTimeSlicePropertyType)
	}
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSignificantPointInAirspaceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSignificantPointInAirspaceTimeSliceTypeCreateCallback.OnAfterCreate(stage, significantpointinairspacetimeslicetype)
	}
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSignificantPointInAirspaceTimeSliceTypeUpdateCallback != nil {
		var frontSignificantPointInAirspaceTimeSliceType *SignificantPointInAirspaceTimeSliceType
		if front != nil {
			frontSignificantPointInAirspaceTimeSliceType, _ = front.(*SignificantPointInAirspaceTimeSliceType)
		}
		stage.OnAfterSignificantPointInAirspaceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, significantpointinairspacetimeslicetype, frontSignificantPointInAirspaceTimeSliceType)
	}
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSignificantPointInAirspaceTimeSliceTypeDeleteCallback != nil {
		var frontSignificantPointInAirspaceTimeSliceType *SignificantPointInAirspaceTimeSliceType
		if front != nil {
			frontSignificantPointInAirspaceTimeSliceType, _ = front.(*SignificantPointInAirspaceTimeSliceType)
		}
		stage.OnAfterSignificantPointInAirspaceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, significantpointinairspacetimeslicetype, frontSignificantPointInAirspaceTimeSliceType)
	}
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSignificantPointInAirspaceTypeCreateCallback != nil {
		stage.OnAfterSignificantPointInAirspaceTypeCreateCallback.OnAfterCreate(stage, significantpointinairspacetype)
	}
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSignificantPointInAirspaceTypeUpdateCallback != nil {
		var frontSignificantPointInAirspaceType *SignificantPointInAirspaceType
		if front != nil {
			frontSignificantPointInAirspaceType, _ = front.(*SignificantPointInAirspaceType)
		}
		stage.OnAfterSignificantPointInAirspaceTypeUpdateCallback.OnAfterUpdate(stage, significantpointinairspacetype, frontSignificantPointInAirspaceType)
	}
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSignificantPointInAirspaceTypeDeleteCallback != nil {
		var frontSignificantPointInAirspaceType *SignificantPointInAirspaceType
		if front != nil {
			frontSignificantPointInAirspaceType, _ = front.(*SignificantPointInAirspaceType)
		}
		stage.OnAfterSignificantPointInAirspaceTypeDeleteCallback.OnAfterDelete(stage, significantpointinairspacetype, frontSignificantPointInAirspaceType)
	}
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationStationStatusPropertyTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationStationStatusPropertyTypeCreateCallback.OnAfterCreate(stage, specialnavigationstationstatuspropertytype)
	}
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationStatusPropertyTypeUpdateCallback != nil {
		var frontSpecialNavigationStationStatusPropertyType *SpecialNavigationStationStatusPropertyType
		if front != nil {
			frontSpecialNavigationStationStatusPropertyType, _ = front.(*SpecialNavigationStationStatusPropertyType)
		}
		stage.OnAfterSpecialNavigationStationStatusPropertyTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationstationstatuspropertytype, frontSpecialNavigationStationStatusPropertyType)
	}
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationStatusPropertyTypeDeleteCallback != nil {
		var frontSpecialNavigationStationStatusPropertyType *SpecialNavigationStationStatusPropertyType
		if front != nil {
			frontSpecialNavigationStationStatusPropertyType, _ = front.(*SpecialNavigationStationStatusPropertyType)
		}
		stage.OnAfterSpecialNavigationStationStatusPropertyTypeDeleteCallback.OnAfterDelete(stage, specialnavigationstationstatuspropertytype, frontSpecialNavigationStationStatusPropertyType)
	}
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationStationStatusTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationStationStatusTypeCreateCallback.OnAfterCreate(stage, specialnavigationstationstatustype)
	}
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationStatusTypeUpdateCallback != nil {
		var frontSpecialNavigationStationStatusType *SpecialNavigationStationStatusType
		if front != nil {
			frontSpecialNavigationStationStatusType, _ = front.(*SpecialNavigationStationStatusType)
		}
		stage.OnAfterSpecialNavigationStationStatusTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationstationstatustype, frontSpecialNavigationStationStatusType)
	}
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationStatusTypeDeleteCallback != nil {
		var frontSpecialNavigationStationStatusType *SpecialNavigationStationStatusType
		if front != nil {
			frontSpecialNavigationStationStatusType, _ = front.(*SpecialNavigationStationStatusType)
		}
		stage.OnAfterSpecialNavigationStationStatusTypeDeleteCallback.OnAfterDelete(stage, specialnavigationstationstatustype, frontSpecialNavigationStationStatusType)
	}
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationStationTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationStationTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, specialnavigationstationtimeslicepropertytype)
	}
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSpecialNavigationStationTimeSlicePropertyType *SpecialNavigationStationTimeSlicePropertyType
		if front != nil {
			frontSpecialNavigationStationTimeSlicePropertyType, _ = front.(*SpecialNavigationStationTimeSlicePropertyType)
		}
		stage.OnAfterSpecialNavigationStationTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationstationtimeslicepropertytype, frontSpecialNavigationStationTimeSlicePropertyType)
	}
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSpecialNavigationStationTimeSlicePropertyType *SpecialNavigationStationTimeSlicePropertyType
		if front != nil {
			frontSpecialNavigationStationTimeSlicePropertyType, _ = front.(*SpecialNavigationStationTimeSlicePropertyType)
		}
		stage.OnAfterSpecialNavigationStationTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, specialnavigationstationtimeslicepropertytype, frontSpecialNavigationStationTimeSlicePropertyType)
	}
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationStationTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationStationTimeSliceTypeCreateCallback.OnAfterCreate(stage, specialnavigationstationtimeslicetype)
	}
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationTimeSliceTypeUpdateCallback != nil {
		var frontSpecialNavigationStationTimeSliceType *SpecialNavigationStationTimeSliceType
		if front != nil {
			frontSpecialNavigationStationTimeSliceType, _ = front.(*SpecialNavigationStationTimeSliceType)
		}
		stage.OnAfterSpecialNavigationStationTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationstationtimeslicetype, frontSpecialNavigationStationTimeSliceType)
	}
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationTimeSliceTypeDeleteCallback != nil {
		var frontSpecialNavigationStationTimeSliceType *SpecialNavigationStationTimeSliceType
		if front != nil {
			frontSpecialNavigationStationTimeSliceType, _ = front.(*SpecialNavigationStationTimeSliceType)
		}
		stage.OnAfterSpecialNavigationStationTimeSliceTypeDeleteCallback.OnAfterDelete(stage, specialnavigationstationtimeslicetype, frontSpecialNavigationStationTimeSliceType)
	}
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationStationTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationStationTypeCreateCallback.OnAfterCreate(stage, specialnavigationstationtype)
	}
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationTypeUpdateCallback != nil {
		var frontSpecialNavigationStationType *SpecialNavigationStationType
		if front != nil {
			frontSpecialNavigationStationType, _ = front.(*SpecialNavigationStationType)
		}
		stage.OnAfterSpecialNavigationStationTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationstationtype, frontSpecialNavigationStationType)
	}
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationTypeDeleteCallback != nil {
		var frontSpecialNavigationStationType *SpecialNavigationStationType
		if front != nil {
			frontSpecialNavigationStationType, _ = front.(*SpecialNavigationStationType)
		}
		stage.OnAfterSpecialNavigationStationTypeDeleteCallback.OnAfterDelete(stage, specialnavigationstationtype, frontSpecialNavigationStationType)
	}
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationSystemPropertyTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationSystemPropertyTypeCreateCallback.OnAfterCreate(stage, specialnavigationsystempropertytype)
	}
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationSystemPropertyTypeUpdateCallback != nil {
		var frontSpecialNavigationSystemPropertyType *SpecialNavigationSystemPropertyType
		if front != nil {
			frontSpecialNavigationSystemPropertyType, _ = front.(*SpecialNavigationSystemPropertyType)
		}
		stage.OnAfterSpecialNavigationSystemPropertyTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationsystempropertytype, frontSpecialNavigationSystemPropertyType)
	}
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationSystemPropertyTypeDeleteCallback != nil {
		var frontSpecialNavigationSystemPropertyType *SpecialNavigationSystemPropertyType
		if front != nil {
			frontSpecialNavigationSystemPropertyType, _ = front.(*SpecialNavigationSystemPropertyType)
		}
		stage.OnAfterSpecialNavigationSystemPropertyTypeDeleteCallback.OnAfterDelete(stage, specialnavigationsystempropertytype, frontSpecialNavigationSystemPropertyType)
	}
}

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationSystemTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationSystemTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, specialnavigationsystemtimeslicepropertytype)
	}
}

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationSystemTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSpecialNavigationSystemTimeSlicePropertyType *SpecialNavigationSystemTimeSlicePropertyType
		if front != nil {
			frontSpecialNavigationSystemTimeSlicePropertyType, _ = front.(*SpecialNavigationSystemTimeSlicePropertyType)
		}
		stage.OnAfterSpecialNavigationSystemTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationsystemtimeslicepropertytype, frontSpecialNavigationSystemTimeSlicePropertyType)
	}
}

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationSystemTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSpecialNavigationSystemTimeSlicePropertyType *SpecialNavigationSystemTimeSlicePropertyType
		if front != nil {
			frontSpecialNavigationSystemTimeSlicePropertyType, _ = front.(*SpecialNavigationSystemTimeSlicePropertyType)
		}
		stage.OnAfterSpecialNavigationSystemTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, specialnavigationsystemtimeslicepropertytype, frontSpecialNavigationSystemTimeSlicePropertyType)
	}
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationSystemTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationSystemTimeSliceTypeCreateCallback.OnAfterCreate(stage, specialnavigationsystemtimeslicetype)
	}
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationSystemTimeSliceTypeUpdateCallback != nil {
		var frontSpecialNavigationSystemTimeSliceType *SpecialNavigationSystemTimeSliceType
		if front != nil {
			frontSpecialNavigationSystemTimeSliceType, _ = front.(*SpecialNavigationSystemTimeSliceType)
		}
		stage.OnAfterSpecialNavigationSystemTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationsystemtimeslicetype, frontSpecialNavigationSystemTimeSliceType)
	}
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationSystemTimeSliceTypeDeleteCallback != nil {
		var frontSpecialNavigationSystemTimeSliceType *SpecialNavigationSystemTimeSliceType
		if front != nil {
			frontSpecialNavigationSystemTimeSliceType, _ = front.(*SpecialNavigationSystemTimeSliceType)
		}
		stage.OnAfterSpecialNavigationSystemTimeSliceTypeDeleteCallback.OnAfterDelete(stage, specialnavigationsystemtimeslicetype, frontSpecialNavigationSystemTimeSliceType)
	}
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationSystemTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationSystemTypeCreateCallback.OnAfterCreate(stage, specialnavigationsystemtype)
	}
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationSystemTypeUpdateCallback != nil {
		var frontSpecialNavigationSystemType *SpecialNavigationSystemType
		if front != nil {
			frontSpecialNavigationSystemType, _ = front.(*SpecialNavigationSystemType)
		}
		stage.OnAfterSpecialNavigationSystemTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationsystemtype, frontSpecialNavigationSystemType)
	}
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationSystemTypeDeleteCallback != nil {
		var frontSpecialNavigationSystemType *SpecialNavigationSystemType
		if front != nil {
			frontSpecialNavigationSystemType, _ = front.(*SpecialNavigationSystemType)
		}
		stage.OnAfterSpecialNavigationSystemTypeDeleteCallback.OnAfterDelete(stage, specialnavigationsystemtype, frontSpecialNavigationSystemType)
	}
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurveillanceGroundStationPropertyTypeCreateCallback != nil {
		stage.OnAfterSurveillanceGroundStationPropertyTypeCreateCallback.OnAfterCreate(stage, surveillancegroundstationpropertytype)
	}
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveillanceGroundStationPropertyTypeUpdateCallback != nil {
		var frontSurveillanceGroundStationPropertyType *SurveillanceGroundStationPropertyType
		if front != nil {
			frontSurveillanceGroundStationPropertyType, _ = front.(*SurveillanceGroundStationPropertyType)
		}
		stage.OnAfterSurveillanceGroundStationPropertyTypeUpdateCallback.OnAfterUpdate(stage, surveillancegroundstationpropertytype, frontSurveillanceGroundStationPropertyType)
	}
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveillanceGroundStationPropertyTypeDeleteCallback != nil {
		var frontSurveillanceGroundStationPropertyType *SurveillanceGroundStationPropertyType
		if front != nil {
			frontSurveillanceGroundStationPropertyType, _ = front.(*SurveillanceGroundStationPropertyType)
		}
		stage.OnAfterSurveillanceGroundStationPropertyTypeDeleteCallback.OnAfterDelete(stage, surveillancegroundstationpropertytype, frontSurveillanceGroundStationPropertyType)
	}
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurveillanceGroundStationTypeCreateCallback != nil {
		stage.OnAfterSurveillanceGroundStationTypeCreateCallback.OnAfterCreate(stage, surveillancegroundstationtype)
	}
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveillanceGroundStationTypeUpdateCallback != nil {
		var frontSurveillanceGroundStationType *SurveillanceGroundStationType
		if front != nil {
			frontSurveillanceGroundStationType, _ = front.(*SurveillanceGroundStationType)
		}
		stage.OnAfterSurveillanceGroundStationTypeUpdateCallback.OnAfterUpdate(stage, surveillancegroundstationtype, frontSurveillanceGroundStationType)
	}
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurveillanceGroundStationTypeDeleteCallback != nil {
		var frontSurveillanceGroundStationType *SurveillanceGroundStationType
		if front != nil {
			frontSurveillanceGroundStationType, _ = front.(*SurveillanceGroundStationType)
		}
		stage.OnAfterSurveillanceGroundStationTypeDeleteCallback.OnAfterDelete(stage, surveillancegroundstationtype, frontSurveillanceGroundStationType)
	}
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTACANTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTACANTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, tacantimeslicepropertytype)
	}
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTACANTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTACANTimeSlicePropertyType *TACANTimeSlicePropertyType
		if front != nil {
			frontTACANTimeSlicePropertyType, _ = front.(*TACANTimeSlicePropertyType)
		}
		stage.OnAfterTACANTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, tacantimeslicepropertytype, frontTACANTimeSlicePropertyType)
	}
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTACANTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTACANTimeSlicePropertyType *TACANTimeSlicePropertyType
		if front != nil {
			frontTACANTimeSlicePropertyType, _ = front.(*TACANTimeSlicePropertyType)
		}
		stage.OnAfterTACANTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, tacantimeslicepropertytype, frontTACANTimeSlicePropertyType)
	}
}

func (tacantimeslicetype *TACANTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTACANTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTACANTimeSliceTypeCreateCallback.OnAfterCreate(stage, tacantimeslicetype)
	}
}

func (tacantimeslicetype *TACANTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTACANTimeSliceTypeUpdateCallback != nil {
		var frontTACANTimeSliceType *TACANTimeSliceType
		if front != nil {
			frontTACANTimeSliceType, _ = front.(*TACANTimeSliceType)
		}
		stage.OnAfterTACANTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, tacantimeslicetype, frontTACANTimeSliceType)
	}
}

func (tacantimeslicetype *TACANTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTACANTimeSliceTypeDeleteCallback != nil {
		var frontTACANTimeSliceType *TACANTimeSliceType
		if front != nil {
			frontTACANTimeSliceType, _ = front.(*TACANTimeSliceType)
		}
		stage.OnAfterTACANTimeSliceTypeDeleteCallback.OnAfterDelete(stage, tacantimeslicetype, frontTACANTimeSliceType)
	}
}

func (tacantype *TACANType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTACANTypeCreateCallback != nil {
		stage.OnAfterTACANTypeCreateCallback.OnAfterCreate(stage, tacantype)
	}
}

func (tacantype *TACANType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTACANTypeUpdateCallback != nil {
		var frontTACANType *TACANType
		if front != nil {
			frontTACANType, _ = front.(*TACANType)
		}
		stage.OnAfterTACANTypeUpdateCallback.OnAfterUpdate(stage, tacantype, frontTACANType)
	}
}

func (tacantype *TACANType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTACANTypeDeleteCallback != nil {
		var frontTACANType *TACANType
		if front != nil {
			frontTACANType, _ = front.(*TACANType)
		}
		stage.OnAfterTACANTypeDeleteCallback.OnAfterDelete(stage, tacantype, frontTACANType)
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

func (unitpropertytype *UnitPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitPropertyTypeCreateCallback != nil {
		stage.OnAfterUnitPropertyTypeCreateCallback.OnAfterCreate(stage, unitpropertytype)
	}
}

func (unitpropertytype *UnitPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitPropertyTypeUpdateCallback != nil {
		var frontUnitPropertyType *UnitPropertyType
		if front != nil {
			frontUnitPropertyType, _ = front.(*UnitPropertyType)
		}
		stage.OnAfterUnitPropertyTypeUpdateCallback.OnAfterUpdate(stage, unitpropertytype, frontUnitPropertyType)
	}
}

func (unitpropertytype *UnitPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitPropertyTypeDeleteCallback != nil {
		var frontUnitPropertyType *UnitPropertyType
		if front != nil {
			frontUnitPropertyType, _ = front.(*UnitPropertyType)
		}
		stage.OnAfterUnitPropertyTypeDeleteCallback.OnAfterDelete(stage, unitpropertytype, frontUnitPropertyType)
	}
}

func (vorpropertytype *VORPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVORPropertyTypeCreateCallback != nil {
		stage.OnAfterVORPropertyTypeCreateCallback.OnAfterCreate(stage, vorpropertytype)
	}
}

func (vorpropertytype *VORPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVORPropertyTypeUpdateCallback != nil {
		var frontVORPropertyType *VORPropertyType
		if front != nil {
			frontVORPropertyType, _ = front.(*VORPropertyType)
		}
		stage.OnAfterVORPropertyTypeUpdateCallback.OnAfterUpdate(stage, vorpropertytype, frontVORPropertyType)
	}
}

func (vorpropertytype *VORPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVORPropertyTypeDeleteCallback != nil {
		var frontVORPropertyType *VORPropertyType
		if front != nil {
			frontVORPropertyType, _ = front.(*VORPropertyType)
		}
		stage.OnAfterVORPropertyTypeDeleteCallback.OnAfterDelete(stage, vorpropertytype, frontVORPropertyType)
	}
}

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVORTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterVORTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, vortimeslicepropertytype)
	}
}

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVORTimeSlicePropertyTypeUpdateCallback != nil {
		var frontVORTimeSlicePropertyType *VORTimeSlicePropertyType
		if front != nil {
			frontVORTimeSlicePropertyType, _ = front.(*VORTimeSlicePropertyType)
		}
		stage.OnAfterVORTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, vortimeslicepropertytype, frontVORTimeSlicePropertyType)
	}
}

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVORTimeSlicePropertyTypeDeleteCallback != nil {
		var frontVORTimeSlicePropertyType *VORTimeSlicePropertyType
		if front != nil {
			frontVORTimeSlicePropertyType, _ = front.(*VORTimeSlicePropertyType)
		}
		stage.OnAfterVORTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, vortimeslicepropertytype, frontVORTimeSlicePropertyType)
	}
}

func (vortimeslicetype *VORTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVORTimeSliceTypeCreateCallback != nil {
		stage.OnAfterVORTimeSliceTypeCreateCallback.OnAfterCreate(stage, vortimeslicetype)
	}
}

func (vortimeslicetype *VORTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVORTimeSliceTypeUpdateCallback != nil {
		var frontVORTimeSliceType *VORTimeSliceType
		if front != nil {
			frontVORTimeSliceType, _ = front.(*VORTimeSliceType)
		}
		stage.OnAfterVORTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, vortimeslicetype, frontVORTimeSliceType)
	}
}

func (vortimeslicetype *VORTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVORTimeSliceTypeDeleteCallback != nil {
		var frontVORTimeSliceType *VORTimeSliceType
		if front != nil {
			frontVORTimeSliceType, _ = front.(*VORTimeSliceType)
		}
		stage.OnAfterVORTimeSliceTypeDeleteCallback.OnAfterDelete(stage, vortimeslicetype, frontVORTimeSliceType)
	}
}

func (vortype *VORType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVORTypeCreateCallback != nil {
		stage.OnAfterVORTypeCreateCallback.OnAfterCreate(stage, vortype)
	}
}

func (vortype *VORType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVORTypeUpdateCallback != nil {
		var frontVORType *VORType
		if front != nil {
			frontVORType, _ = front.(*VORType)
		}
		stage.OnAfterVORTypeUpdateCallback.OnAfterUpdate(stage, vortype, frontVORType)
	}
}

func (vortype *VORType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVORTypeDeleteCallback != nil {
		var frontVORType *VORType
		if front != nil {
			frontVORType, _ = front.(*VORType)
		}
		stage.OnAfterVORTypeDeleteCallback.OnAfterDelete(stage, vortype, frontVORType)
	}
}

