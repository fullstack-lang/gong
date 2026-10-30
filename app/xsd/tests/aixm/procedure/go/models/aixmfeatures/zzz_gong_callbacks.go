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

func (abstractapproachlegtype *AbstractApproachLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractApproachLegTypeCreateCallback != nil {
		stage.OnAfterAbstractApproachLegTypeCreateCallback.OnAfterCreate(stage, abstractapproachlegtype)
	}
}

func (abstractapproachlegtype *AbstractApproachLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractApproachLegTypeUpdateCallback != nil {
		var frontAbstractApproachLegType *AbstractApproachLegType
		if front != nil {
			frontAbstractApproachLegType, _ = front.(*AbstractApproachLegType)
		}
		stage.OnAfterAbstractApproachLegTypeUpdateCallback.OnAfterUpdate(stage, abstractapproachlegtype, frontAbstractApproachLegType)
	}
}

func (abstractapproachlegtype *AbstractApproachLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractApproachLegTypeDeleteCallback != nil {
		var frontAbstractApproachLegType *AbstractApproachLegType
		if front != nil {
			frontAbstractApproachLegType, _ = front.(*AbstractApproachLegType)
		}
		stage.OnAfterAbstractApproachLegTypeDeleteCallback.OnAfterDelete(stage, abstractapproachlegtype, frontAbstractApproachLegType)
	}
}

func (abstractproceduretype *AbstractProcedureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractProcedureTypeCreateCallback != nil {
		stage.OnAfterAbstractProcedureTypeCreateCallback.OnAfterCreate(stage, abstractproceduretype)
	}
}

func (abstractproceduretype *AbstractProcedureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractProcedureTypeUpdateCallback != nil {
		var frontAbstractProcedureType *AbstractProcedureType
		if front != nil {
			frontAbstractProcedureType, _ = front.(*AbstractProcedureType)
		}
		stage.OnAfterAbstractProcedureTypeUpdateCallback.OnAfterUpdate(stage, abstractproceduretype, frontAbstractProcedureType)
	}
}

func (abstractproceduretype *AbstractProcedureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractProcedureTypeDeleteCallback != nil {
		var frontAbstractProcedureType *AbstractProcedureType
		if front != nil {
			frontAbstractProcedureType, _ = front.(*AbstractProcedureType)
		}
		stage.OnAfterAbstractProcedureTypeDeleteCallback.OnAfterDelete(stage, abstractproceduretype, frontAbstractProcedureType)
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

func (abstractsegmentlegtype *AbstractSegmentLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractSegmentLegTypeCreateCallback != nil {
		stage.OnAfterAbstractSegmentLegTypeCreateCallback.OnAfterCreate(stage, abstractsegmentlegtype)
	}
}

func (abstractsegmentlegtype *AbstractSegmentLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSegmentLegTypeUpdateCallback != nil {
		var frontAbstractSegmentLegType *AbstractSegmentLegType
		if front != nil {
			frontAbstractSegmentLegType, _ = front.(*AbstractSegmentLegType)
		}
		stage.OnAfterAbstractSegmentLegTypeUpdateCallback.OnAfterUpdate(stage, abstractsegmentlegtype, frontAbstractSegmentLegType)
	}
}

func (abstractsegmentlegtype *AbstractSegmentLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSegmentLegTypeDeleteCallback != nil {
		var frontAbstractSegmentLegType *AbstractSegmentLegType
		if front != nil {
			frontAbstractSegmentLegType, _ = front.(*AbstractSegmentLegType)
		}
		stage.OnAfterAbstractSegmentLegTypeDeleteCallback.OnAfterDelete(stage, abstractsegmentlegtype, frontAbstractSegmentLegType)
	}
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractSegmentPointTypeCreateCallback != nil {
		stage.OnAfterAbstractSegmentPointTypeCreateCallback.OnAfterCreate(stage, abstractsegmentpointtype)
	}
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSegmentPointTypeUpdateCallback != nil {
		var frontAbstractSegmentPointType *AbstractSegmentPointType
		if front != nil {
			frontAbstractSegmentPointType, _ = front.(*AbstractSegmentPointType)
		}
		stage.OnAfterAbstractSegmentPointTypeUpdateCallback.OnAfterUpdate(stage, abstractsegmentpointtype, frontAbstractSegmentPointType)
	}
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSegmentPointTypeDeleteCallback != nil {
		var frontAbstractSegmentPointType *AbstractSegmentPointType
		if front != nil {
			frontAbstractSegmentPointType, _ = front.(*AbstractSegmentPointType)
		}
		stage.OnAfterAbstractSegmentPointTypeDeleteCallback.OnAfterDelete(stage, abstractsegmentpointtype, frontAbstractSegmentPointType)
	}
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingPointTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingPointTypeCreateCallback.OnAfterCreate(stage, aerialrefuellingpointtype)
	}
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingPointTypeUpdateCallback != nil {
		var frontAerialRefuellingPointType *AerialRefuellingPointType
		if front != nil {
			frontAerialRefuellingPointType, _ = front.(*AerialRefuellingPointType)
		}
		stage.OnAfterAerialRefuellingPointTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellingpointtype, frontAerialRefuellingPointType)
	}
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingPointTypeDeleteCallback != nil {
		var frontAerialRefuellingPointType *AerialRefuellingPointType
		if front != nil {
			frontAerialRefuellingPointType, _ = front.(*AerialRefuellingPointType)
		}
		stage.OnAfterAerialRefuellingPointTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellingpointtype, frontAerialRefuellingPointType)
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

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAltitudeAdjustmentPropertyTypeCreateCallback != nil {
		stage.OnAfterAltitudeAdjustmentPropertyTypeCreateCallback.OnAfterCreate(stage, altitudeadjustmentpropertytype)
	}
}

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltitudeAdjustmentPropertyTypeUpdateCallback != nil {
		var frontAltitudeAdjustmentPropertyType *AltitudeAdjustmentPropertyType
		if front != nil {
			frontAltitudeAdjustmentPropertyType, _ = front.(*AltitudeAdjustmentPropertyType)
		}
		stage.OnAfterAltitudeAdjustmentPropertyTypeUpdateCallback.OnAfterUpdate(stage, altitudeadjustmentpropertytype, frontAltitudeAdjustmentPropertyType)
	}
}

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltitudeAdjustmentPropertyTypeDeleteCallback != nil {
		var frontAltitudeAdjustmentPropertyType *AltitudeAdjustmentPropertyType
		if front != nil {
			frontAltitudeAdjustmentPropertyType, _ = front.(*AltitudeAdjustmentPropertyType)
		}
		stage.OnAfterAltitudeAdjustmentPropertyTypeDeleteCallback.OnAfterDelete(stage, altitudeadjustmentpropertytype, frontAltitudeAdjustmentPropertyType)
	}
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAltitudeAdjustmentTypeCreateCallback != nil {
		stage.OnAfterAltitudeAdjustmentTypeCreateCallback.OnAfterCreate(stage, altitudeadjustmenttype)
	}
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltitudeAdjustmentTypeUpdateCallback != nil {
		var frontAltitudeAdjustmentType *AltitudeAdjustmentType
		if front != nil {
			frontAltitudeAdjustmentType, _ = front.(*AltitudeAdjustmentType)
		}
		stage.OnAfterAltitudeAdjustmentTypeUpdateCallback.OnAfterUpdate(stage, altitudeadjustmenttype, frontAltitudeAdjustmentType)
	}
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAltitudeAdjustmentTypeDeleteCallback != nil {
		var frontAltitudeAdjustmentType *AltitudeAdjustmentType
		if front != nil {
			frontAltitudeAdjustmentType, _ = front.(*AltitudeAdjustmentType)
		}
		stage.OnAfterAltitudeAdjustmentTypeDeleteCallback.OnAfterDelete(stage, altitudeadjustmenttype, frontAltitudeAdjustmentType)
	}
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngleIndicationPropertyTypeCreateCallback != nil {
		stage.OnAfterAngleIndicationPropertyTypeCreateCallback.OnAfterCreate(stage, angleindicationpropertytype)
	}
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleIndicationPropertyTypeUpdateCallback != nil {
		var frontAngleIndicationPropertyType *AngleIndicationPropertyType
		if front != nil {
			frontAngleIndicationPropertyType, _ = front.(*AngleIndicationPropertyType)
		}
		stage.OnAfterAngleIndicationPropertyTypeUpdateCallback.OnAfterUpdate(stage, angleindicationpropertytype, frontAngleIndicationPropertyType)
	}
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleIndicationPropertyTypeDeleteCallback != nil {
		var frontAngleIndicationPropertyType *AngleIndicationPropertyType
		if front != nil {
			frontAngleIndicationPropertyType, _ = front.(*AngleIndicationPropertyType)
		}
		stage.OnAfterAngleIndicationPropertyTypeDeleteCallback.OnAfterDelete(stage, angleindicationpropertytype, frontAngleIndicationPropertyType)
	}
}

func (angleusepropertytype *AngleUsePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngleUsePropertyTypeCreateCallback != nil {
		stage.OnAfterAngleUsePropertyTypeCreateCallback.OnAfterCreate(stage, angleusepropertytype)
	}
}

func (angleusepropertytype *AngleUsePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleUsePropertyTypeUpdateCallback != nil {
		var frontAngleUsePropertyType *AngleUsePropertyType
		if front != nil {
			frontAngleUsePropertyType, _ = front.(*AngleUsePropertyType)
		}
		stage.OnAfterAngleUsePropertyTypeUpdateCallback.OnAfterUpdate(stage, angleusepropertytype, frontAngleUsePropertyType)
	}
}

func (angleusepropertytype *AngleUsePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleUsePropertyTypeDeleteCallback != nil {
		var frontAngleUsePropertyType *AngleUsePropertyType
		if front != nil {
			frontAngleUsePropertyType, _ = front.(*AngleUsePropertyType)
		}
		stage.OnAfterAngleUsePropertyTypeDeleteCallback.OnAfterDelete(stage, angleusepropertytype, frontAngleUsePropertyType)
	}
}

func (angleusetype *AngleUseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngleUseTypeCreateCallback != nil {
		stage.OnAfterAngleUseTypeCreateCallback.OnAfterCreate(stage, angleusetype)
	}
}

func (angleusetype *AngleUseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleUseTypeUpdateCallback != nil {
		var frontAngleUseType *AngleUseType
		if front != nil {
			frontAngleUseType, _ = front.(*AngleUseType)
		}
		stage.OnAfterAngleUseTypeUpdateCallback.OnAfterUpdate(stage, angleusetype, frontAngleUseType)
	}
}

func (angleusetype *AngleUseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleUseTypeDeleteCallback != nil {
		var frontAngleUseType *AngleUseType
		if front != nil {
			frontAngleUseType, _ = front.(*AngleUseType)
		}
		stage.OnAfterAngleUseTypeDeleteCallback.OnAfterDelete(stage, angleusetype, frontAngleUseType)
	}
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachAltitudeTablePropertyTypeCreateCallback != nil {
		stage.OnAfterApproachAltitudeTablePropertyTypeCreateCallback.OnAfterCreate(stage, approachaltitudetablepropertytype)
	}
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachAltitudeTablePropertyTypeUpdateCallback != nil {
		var frontApproachAltitudeTablePropertyType *ApproachAltitudeTablePropertyType
		if front != nil {
			frontApproachAltitudeTablePropertyType, _ = front.(*ApproachAltitudeTablePropertyType)
		}
		stage.OnAfterApproachAltitudeTablePropertyTypeUpdateCallback.OnAfterUpdate(stage, approachaltitudetablepropertytype, frontApproachAltitudeTablePropertyType)
	}
}

func (approachaltitudetablepropertytype *ApproachAltitudeTablePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachAltitudeTablePropertyTypeDeleteCallback != nil {
		var frontApproachAltitudeTablePropertyType *ApproachAltitudeTablePropertyType
		if front != nil {
			frontApproachAltitudeTablePropertyType, _ = front.(*ApproachAltitudeTablePropertyType)
		}
		stage.OnAfterApproachAltitudeTablePropertyTypeDeleteCallback.OnAfterDelete(stage, approachaltitudetablepropertytype, frontApproachAltitudeTablePropertyType)
	}
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachAltitudeTableTypeCreateCallback != nil {
		stage.OnAfterApproachAltitudeTableTypeCreateCallback.OnAfterCreate(stage, approachaltitudetabletype)
	}
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachAltitudeTableTypeUpdateCallback != nil {
		var frontApproachAltitudeTableType *ApproachAltitudeTableType
		if front != nil {
			frontApproachAltitudeTableType, _ = front.(*ApproachAltitudeTableType)
		}
		stage.OnAfterApproachAltitudeTableTypeUpdateCallback.OnAfterUpdate(stage, approachaltitudetabletype, frontApproachAltitudeTableType)
	}
}

func (approachaltitudetabletype *ApproachAltitudeTableType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachAltitudeTableTypeDeleteCallback != nil {
		var frontApproachAltitudeTableType *ApproachAltitudeTableType
		if front != nil {
			frontApproachAltitudeTableType, _ = front.(*ApproachAltitudeTableType)
		}
		stage.OnAfterApproachAltitudeTableTypeDeleteCallback.OnAfterDelete(stage, approachaltitudetabletype, frontApproachAltitudeTableType)
	}
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachConditionPropertyTypeCreateCallback != nil {
		stage.OnAfterApproachConditionPropertyTypeCreateCallback.OnAfterCreate(stage, approachconditionpropertytype)
	}
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachConditionPropertyTypeUpdateCallback != nil {
		var frontApproachConditionPropertyType *ApproachConditionPropertyType
		if front != nil {
			frontApproachConditionPropertyType, _ = front.(*ApproachConditionPropertyType)
		}
		stage.OnAfterApproachConditionPropertyTypeUpdateCallback.OnAfterUpdate(stage, approachconditionpropertytype, frontApproachConditionPropertyType)
	}
}

func (approachconditionpropertytype *ApproachConditionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachConditionPropertyTypeDeleteCallback != nil {
		var frontApproachConditionPropertyType *ApproachConditionPropertyType
		if front != nil {
			frontApproachConditionPropertyType, _ = front.(*ApproachConditionPropertyType)
		}
		stage.OnAfterApproachConditionPropertyTypeDeleteCallback.OnAfterDelete(stage, approachconditionpropertytype, frontApproachConditionPropertyType)
	}
}

func (approachconditiontype *ApproachConditionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachConditionTypeCreateCallback != nil {
		stage.OnAfterApproachConditionTypeCreateCallback.OnAfterCreate(stage, approachconditiontype)
	}
}

func (approachconditiontype *ApproachConditionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachConditionTypeUpdateCallback != nil {
		var frontApproachConditionType *ApproachConditionType
		if front != nil {
			frontApproachConditionType, _ = front.(*ApproachConditionType)
		}
		stage.OnAfterApproachConditionTypeUpdateCallback.OnAfterUpdate(stage, approachconditiontype, frontApproachConditionType)
	}
}

func (approachconditiontype *ApproachConditionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachConditionTypeDeleteCallback != nil {
		var frontApproachConditionType *ApproachConditionType
		if front != nil {
			frontApproachConditionType, _ = front.(*ApproachConditionType)
		}
		stage.OnAfterApproachConditionTypeDeleteCallback.OnAfterDelete(stage, approachconditiontype, frontApproachConditionType)
	}
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachDistanceTablePropertyTypeCreateCallback != nil {
		stage.OnAfterApproachDistanceTablePropertyTypeCreateCallback.OnAfterCreate(stage, approachdistancetablepropertytype)
	}
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachDistanceTablePropertyTypeUpdateCallback != nil {
		var frontApproachDistanceTablePropertyType *ApproachDistanceTablePropertyType
		if front != nil {
			frontApproachDistanceTablePropertyType, _ = front.(*ApproachDistanceTablePropertyType)
		}
		stage.OnAfterApproachDistanceTablePropertyTypeUpdateCallback.OnAfterUpdate(stage, approachdistancetablepropertytype, frontApproachDistanceTablePropertyType)
	}
}

func (approachdistancetablepropertytype *ApproachDistanceTablePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachDistanceTablePropertyTypeDeleteCallback != nil {
		var frontApproachDistanceTablePropertyType *ApproachDistanceTablePropertyType
		if front != nil {
			frontApproachDistanceTablePropertyType, _ = front.(*ApproachDistanceTablePropertyType)
		}
		stage.OnAfterApproachDistanceTablePropertyTypeDeleteCallback.OnAfterDelete(stage, approachdistancetablepropertytype, frontApproachDistanceTablePropertyType)
	}
}

func (approachdistancetabletype *ApproachDistanceTableType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachDistanceTableTypeCreateCallback != nil {
		stage.OnAfterApproachDistanceTableTypeCreateCallback.OnAfterCreate(stage, approachdistancetabletype)
	}
}

func (approachdistancetabletype *ApproachDistanceTableType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachDistanceTableTypeUpdateCallback != nil {
		var frontApproachDistanceTableType *ApproachDistanceTableType
		if front != nil {
			frontApproachDistanceTableType, _ = front.(*ApproachDistanceTableType)
		}
		stage.OnAfterApproachDistanceTableTypeUpdateCallback.OnAfterUpdate(stage, approachdistancetabletype, frontApproachDistanceTableType)
	}
}

func (approachdistancetabletype *ApproachDistanceTableType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachDistanceTableTypeDeleteCallback != nil {
		var frontApproachDistanceTableType *ApproachDistanceTableType
		if front != nil {
			frontApproachDistanceTableType, _ = front.(*ApproachDistanceTableType)
		}
		stage.OnAfterApproachDistanceTableTypeDeleteCallback.OnAfterDelete(stage, approachdistancetabletype, frontApproachDistanceTableType)
	}
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachTimingTablePropertyTypeCreateCallback != nil {
		stage.OnAfterApproachTimingTablePropertyTypeCreateCallback.OnAfterCreate(stage, approachtimingtablepropertytype)
	}
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachTimingTablePropertyTypeUpdateCallback != nil {
		var frontApproachTimingTablePropertyType *ApproachTimingTablePropertyType
		if front != nil {
			frontApproachTimingTablePropertyType, _ = front.(*ApproachTimingTablePropertyType)
		}
		stage.OnAfterApproachTimingTablePropertyTypeUpdateCallback.OnAfterUpdate(stage, approachtimingtablepropertytype, frontApproachTimingTablePropertyType)
	}
}

func (approachtimingtablepropertytype *ApproachTimingTablePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachTimingTablePropertyTypeDeleteCallback != nil {
		var frontApproachTimingTablePropertyType *ApproachTimingTablePropertyType
		if front != nil {
			frontApproachTimingTablePropertyType, _ = front.(*ApproachTimingTablePropertyType)
		}
		stage.OnAfterApproachTimingTablePropertyTypeDeleteCallback.OnAfterDelete(stage, approachtimingtablepropertytype, frontApproachTimingTablePropertyType)
	}
}

func (approachtimingtabletype *ApproachTimingTableType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterApproachTimingTableTypeCreateCallback != nil {
		stage.OnAfterApproachTimingTableTypeCreateCallback.OnAfterCreate(stage, approachtimingtabletype)
	}
}

func (approachtimingtabletype *ApproachTimingTableType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachTimingTableTypeUpdateCallback != nil {
		var frontApproachTimingTableType *ApproachTimingTableType
		if front != nil {
			frontApproachTimingTableType, _ = front.(*ApproachTimingTableType)
		}
		stage.OnAfterApproachTimingTableTypeUpdateCallback.OnAfterUpdate(stage, approachtimingtabletype, frontApproachTimingTableType)
	}
}

func (approachtimingtabletype *ApproachTimingTableType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterApproachTimingTableTypeDeleteCallback != nil {
		var frontApproachTimingTableType *ApproachTimingTableType
		if front != nil {
			frontApproachTimingTableType, _ = front.(*ApproachTimingTableType)
		}
		stage.OnAfterApproachTimingTableTypeDeleteCallback.OnAfterDelete(stage, approachtimingtabletype, frontApproachTimingTableType)
	}
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrivalFeederLegTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterArrivalFeederLegTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, arrivalfeederlegtimeslicepropertytype)
	}
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalFeederLegTimeSlicePropertyTypeUpdateCallback != nil {
		var frontArrivalFeederLegTimeSlicePropertyType *ArrivalFeederLegTimeSlicePropertyType
		if front != nil {
			frontArrivalFeederLegTimeSlicePropertyType, _ = front.(*ArrivalFeederLegTimeSlicePropertyType)
		}
		stage.OnAfterArrivalFeederLegTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, arrivalfeederlegtimeslicepropertytype, frontArrivalFeederLegTimeSlicePropertyType)
	}
}

func (arrivalfeederlegtimeslicepropertytype *ArrivalFeederLegTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalFeederLegTimeSlicePropertyTypeDeleteCallback != nil {
		var frontArrivalFeederLegTimeSlicePropertyType *ArrivalFeederLegTimeSlicePropertyType
		if front != nil {
			frontArrivalFeederLegTimeSlicePropertyType, _ = front.(*ArrivalFeederLegTimeSlicePropertyType)
		}
		stage.OnAfterArrivalFeederLegTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, arrivalfeederlegtimeslicepropertytype, frontArrivalFeederLegTimeSlicePropertyType)
	}
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrivalFeederLegTimeSliceTypeCreateCallback != nil {
		stage.OnAfterArrivalFeederLegTimeSliceTypeCreateCallback.OnAfterCreate(stage, arrivalfeederlegtimeslicetype)
	}
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalFeederLegTimeSliceTypeUpdateCallback != nil {
		var frontArrivalFeederLegTimeSliceType *ArrivalFeederLegTimeSliceType
		if front != nil {
			frontArrivalFeederLegTimeSliceType, _ = front.(*ArrivalFeederLegTimeSliceType)
		}
		stage.OnAfterArrivalFeederLegTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, arrivalfeederlegtimeslicetype, frontArrivalFeederLegTimeSliceType)
	}
}

func (arrivalfeederlegtimeslicetype *ArrivalFeederLegTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalFeederLegTimeSliceTypeDeleteCallback != nil {
		var frontArrivalFeederLegTimeSliceType *ArrivalFeederLegTimeSliceType
		if front != nil {
			frontArrivalFeederLegTimeSliceType, _ = front.(*ArrivalFeederLegTimeSliceType)
		}
		stage.OnAfterArrivalFeederLegTimeSliceTypeDeleteCallback.OnAfterDelete(stage, arrivalfeederlegtimeslicetype, frontArrivalFeederLegTimeSliceType)
	}
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrivalFeederLegTypeCreateCallback != nil {
		stage.OnAfterArrivalFeederLegTypeCreateCallback.OnAfterCreate(stage, arrivalfeederlegtype)
	}
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalFeederLegTypeUpdateCallback != nil {
		var frontArrivalFeederLegType *ArrivalFeederLegType
		if front != nil {
			frontArrivalFeederLegType, _ = front.(*ArrivalFeederLegType)
		}
		stage.OnAfterArrivalFeederLegTypeUpdateCallback.OnAfterUpdate(stage, arrivalfeederlegtype, frontArrivalFeederLegType)
	}
}

func (arrivalfeederlegtype *ArrivalFeederLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalFeederLegTypeDeleteCallback != nil {
		var frontArrivalFeederLegType *ArrivalFeederLegType
		if front != nil {
			frontArrivalFeederLegType, _ = front.(*ArrivalFeederLegType)
		}
		stage.OnAfterArrivalFeederLegTypeDeleteCallback.OnAfterDelete(stage, arrivalfeederlegtype, frontArrivalFeederLegType)
	}
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrivalLegTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterArrivalLegTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, arrivallegtimeslicepropertytype)
	}
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalLegTimeSlicePropertyTypeUpdateCallback != nil {
		var frontArrivalLegTimeSlicePropertyType *ArrivalLegTimeSlicePropertyType
		if front != nil {
			frontArrivalLegTimeSlicePropertyType, _ = front.(*ArrivalLegTimeSlicePropertyType)
		}
		stage.OnAfterArrivalLegTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, arrivallegtimeslicepropertytype, frontArrivalLegTimeSlicePropertyType)
	}
}

func (arrivallegtimeslicepropertytype *ArrivalLegTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalLegTimeSlicePropertyTypeDeleteCallback != nil {
		var frontArrivalLegTimeSlicePropertyType *ArrivalLegTimeSlicePropertyType
		if front != nil {
			frontArrivalLegTimeSlicePropertyType, _ = front.(*ArrivalLegTimeSlicePropertyType)
		}
		stage.OnAfterArrivalLegTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, arrivallegtimeslicepropertytype, frontArrivalLegTimeSlicePropertyType)
	}
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrivalLegTimeSliceTypeCreateCallback != nil {
		stage.OnAfterArrivalLegTimeSliceTypeCreateCallback.OnAfterCreate(stage, arrivallegtimeslicetype)
	}
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalLegTimeSliceTypeUpdateCallback != nil {
		var frontArrivalLegTimeSliceType *ArrivalLegTimeSliceType
		if front != nil {
			frontArrivalLegTimeSliceType, _ = front.(*ArrivalLegTimeSliceType)
		}
		stage.OnAfterArrivalLegTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, arrivallegtimeslicetype, frontArrivalLegTimeSliceType)
	}
}

func (arrivallegtimeslicetype *ArrivalLegTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalLegTimeSliceTypeDeleteCallback != nil {
		var frontArrivalLegTimeSliceType *ArrivalLegTimeSliceType
		if front != nil {
			frontArrivalLegTimeSliceType, _ = front.(*ArrivalLegTimeSliceType)
		}
		stage.OnAfterArrivalLegTimeSliceTypeDeleteCallback.OnAfterDelete(stage, arrivallegtimeslicetype, frontArrivalLegTimeSliceType)
	}
}

func (arrivallegtype *ArrivalLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrivalLegTypeCreateCallback != nil {
		stage.OnAfterArrivalLegTypeCreateCallback.OnAfterCreate(stage, arrivallegtype)
	}
}

func (arrivallegtype *ArrivalLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalLegTypeUpdateCallback != nil {
		var frontArrivalLegType *ArrivalLegType
		if front != nil {
			frontArrivalLegType, _ = front.(*ArrivalLegType)
		}
		stage.OnAfterArrivalLegTypeUpdateCallback.OnAfterUpdate(stage, arrivallegtype, frontArrivalLegType)
	}
}

func (arrivallegtype *ArrivalLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrivalLegTypeDeleteCallback != nil {
		var frontArrivalLegType *ArrivalLegType
		if front != nil {
			frontArrivalLegType, _ = front.(*ArrivalLegType)
		}
		stage.OnAfterArrivalLegTypeDeleteCallback.OnAfterDelete(stage, arrivallegtype, frontArrivalLegType)
	}
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCircleSectorPropertyTypeCreateCallback != nil {
		stage.OnAfterCircleSectorPropertyTypeCreateCallback.OnAfterCreate(stage, circlesectorpropertytype)
	}
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleSectorPropertyTypeUpdateCallback != nil {
		var frontCircleSectorPropertyType *CircleSectorPropertyType
		if front != nil {
			frontCircleSectorPropertyType, _ = front.(*CircleSectorPropertyType)
		}
		stage.OnAfterCircleSectorPropertyTypeUpdateCallback.OnAfterUpdate(stage, circlesectorpropertytype, frontCircleSectorPropertyType)
	}
}

func (circlesectorpropertytype *CircleSectorPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleSectorPropertyTypeDeleteCallback != nil {
		var frontCircleSectorPropertyType *CircleSectorPropertyType
		if front != nil {
			frontCircleSectorPropertyType, _ = front.(*CircleSectorPropertyType)
		}
		stage.OnAfterCircleSectorPropertyTypeDeleteCallback.OnAfterDelete(stage, circlesectorpropertytype, frontCircleSectorPropertyType)
	}
}

func (circlesectortype *CircleSectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCircleSectorTypeCreateCallback != nil {
		stage.OnAfterCircleSectorTypeCreateCallback.OnAfterCreate(stage, circlesectortype)
	}
}

func (circlesectortype *CircleSectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleSectorTypeUpdateCallback != nil {
		var frontCircleSectorType *CircleSectorType
		if front != nil {
			frontCircleSectorType, _ = front.(*CircleSectorType)
		}
		stage.OnAfterCircleSectorTypeUpdateCallback.OnAfterUpdate(stage, circlesectortype, frontCircleSectorType)
	}
}

func (circlesectortype *CircleSectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleSectorTypeDeleteCallback != nil {
		var frontCircleSectorType *CircleSectorType
		if front != nil {
			frontCircleSectorType, _ = front.(*CircleSectorType)
		}
		stage.OnAfterCircleSectorTypeDeleteCallback.OnAfterDelete(stage, circlesectortype, frontCircleSectorType)
	}
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCirclingAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterCirclingAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, circlingareatimeslicepropertytype)
	}
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontCirclingAreaTimeSlicePropertyType *CirclingAreaTimeSlicePropertyType
		if front != nil {
			frontCirclingAreaTimeSlicePropertyType, _ = front.(*CirclingAreaTimeSlicePropertyType)
		}
		stage.OnAfterCirclingAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, circlingareatimeslicepropertytype, frontCirclingAreaTimeSlicePropertyType)
	}
}

func (circlingareatimeslicepropertytype *CirclingAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontCirclingAreaTimeSlicePropertyType *CirclingAreaTimeSlicePropertyType
		if front != nil {
			frontCirclingAreaTimeSlicePropertyType, _ = front.(*CirclingAreaTimeSlicePropertyType)
		}
		stage.OnAfterCirclingAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, circlingareatimeslicepropertytype, frontCirclingAreaTimeSlicePropertyType)
	}
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCirclingAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterCirclingAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, circlingareatimeslicetype)
	}
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingAreaTimeSliceTypeUpdateCallback != nil {
		var frontCirclingAreaTimeSliceType *CirclingAreaTimeSliceType
		if front != nil {
			frontCirclingAreaTimeSliceType, _ = front.(*CirclingAreaTimeSliceType)
		}
		stage.OnAfterCirclingAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, circlingareatimeslicetype, frontCirclingAreaTimeSliceType)
	}
}

func (circlingareatimeslicetype *CirclingAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingAreaTimeSliceTypeDeleteCallback != nil {
		var frontCirclingAreaTimeSliceType *CirclingAreaTimeSliceType
		if front != nil {
			frontCirclingAreaTimeSliceType, _ = front.(*CirclingAreaTimeSliceType)
		}
		stage.OnAfterCirclingAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, circlingareatimeslicetype, frontCirclingAreaTimeSliceType)
	}
}

func (circlingareatype *CirclingAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCirclingAreaTypeCreateCallback != nil {
		stage.OnAfterCirclingAreaTypeCreateCallback.OnAfterCreate(stage, circlingareatype)
	}
}

func (circlingareatype *CirclingAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingAreaTypeUpdateCallback != nil {
		var frontCirclingAreaType *CirclingAreaType
		if front != nil {
			frontCirclingAreaType, _ = front.(*CirclingAreaType)
		}
		stage.OnAfterCirclingAreaTypeUpdateCallback.OnAfterUpdate(stage, circlingareatype, frontCirclingAreaType)
	}
}

func (circlingareatype *CirclingAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingAreaTypeDeleteCallback != nil {
		var frontCirclingAreaType *CirclingAreaType
		if front != nil {
			frontCirclingAreaType, _ = front.(*CirclingAreaType)
		}
		stage.OnAfterCirclingAreaTypeDeleteCallback.OnAfterDelete(stage, circlingareatype, frontCirclingAreaType)
	}
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCirclingRestrictionPropertyTypeCreateCallback != nil {
		stage.OnAfterCirclingRestrictionPropertyTypeCreateCallback.OnAfterCreate(stage, circlingrestrictionpropertytype)
	}
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingRestrictionPropertyTypeUpdateCallback != nil {
		var frontCirclingRestrictionPropertyType *CirclingRestrictionPropertyType
		if front != nil {
			frontCirclingRestrictionPropertyType, _ = front.(*CirclingRestrictionPropertyType)
		}
		stage.OnAfterCirclingRestrictionPropertyTypeUpdateCallback.OnAfterUpdate(stage, circlingrestrictionpropertytype, frontCirclingRestrictionPropertyType)
	}
}

func (circlingrestrictionpropertytype *CirclingRestrictionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingRestrictionPropertyTypeDeleteCallback != nil {
		var frontCirclingRestrictionPropertyType *CirclingRestrictionPropertyType
		if front != nil {
			frontCirclingRestrictionPropertyType, _ = front.(*CirclingRestrictionPropertyType)
		}
		stage.OnAfterCirclingRestrictionPropertyTypeDeleteCallback.OnAfterDelete(stage, circlingrestrictionpropertytype, frontCirclingRestrictionPropertyType)
	}
}

func (circlingrestrictiontype *CirclingRestrictionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCirclingRestrictionTypeCreateCallback != nil {
		stage.OnAfterCirclingRestrictionTypeCreateCallback.OnAfterCreate(stage, circlingrestrictiontype)
	}
}

func (circlingrestrictiontype *CirclingRestrictionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingRestrictionTypeUpdateCallback != nil {
		var frontCirclingRestrictionType *CirclingRestrictionType
		if front != nil {
			frontCirclingRestrictionType, _ = front.(*CirclingRestrictionType)
		}
		stage.OnAfterCirclingRestrictionTypeUpdateCallback.OnAfterUpdate(stage, circlingrestrictiontype, frontCirclingRestrictionType)
	}
}

func (circlingrestrictiontype *CirclingRestrictionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCirclingRestrictionTypeDeleteCallback != nil {
		var frontCirclingRestrictionType *CirclingRestrictionType
		if front != nil {
			frontCirclingRestrictionType, _ = front.(*CirclingRestrictionType)
		}
		stage.OnAfterCirclingRestrictionTypeDeleteCallback.OnAfterDelete(stage, circlingrestrictiontype, frontCirclingRestrictionType)
	}
}

func (curvepropertytype *CurvePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurvePropertyTypeCreateCallback != nil {
		stage.OnAfterCurvePropertyTypeCreateCallback.OnAfterCreate(stage, curvepropertytype)
	}
}

func (curvepropertytype *CurvePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurvePropertyTypeUpdateCallback != nil {
		var frontCurvePropertyType *CurvePropertyType
		if front != nil {
			frontCurvePropertyType, _ = front.(*CurvePropertyType)
		}
		stage.OnAfterCurvePropertyTypeUpdateCallback.OnAfterUpdate(stage, curvepropertytype, frontCurvePropertyType)
	}
}

func (curvepropertytype *CurvePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurvePropertyTypeDeleteCallback != nil {
		var frontCurvePropertyType *CurvePropertyType
		if front != nil {
			frontCurvePropertyType, _ = front.(*CurvePropertyType)
		}
		stage.OnAfterCurvePropertyTypeDeleteCallback.OnAfterDelete(stage, curvepropertytype, frontCurvePropertyType)
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

func (dmepropertytype *DMEPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDMEPropertyTypeCreateCallback != nil {
		stage.OnAfterDMEPropertyTypeCreateCallback.OnAfterCreate(stage, dmepropertytype)
	}
}

func (dmepropertytype *DMEPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMEPropertyTypeUpdateCallback != nil {
		var frontDMEPropertyType *DMEPropertyType
		if front != nil {
			frontDMEPropertyType, _ = front.(*DMEPropertyType)
		}
		stage.OnAfterDMEPropertyTypeUpdateCallback.OnAfterUpdate(stage, dmepropertytype, frontDMEPropertyType)
	}
}

func (dmepropertytype *DMEPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMEPropertyTypeDeleteCallback != nil {
		var frontDMEPropertyType *DMEPropertyType
		if front != nil {
			frontDMEPropertyType, _ = front.(*DMEPropertyType)
		}
		stage.OnAfterDMEPropertyTypeDeleteCallback.OnAfterDelete(stage, dmepropertytype, frontDMEPropertyType)
	}
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDepartureArrivalConditionPropertyTypeCreateCallback != nil {
		stage.OnAfterDepartureArrivalConditionPropertyTypeCreateCallback.OnAfterCreate(stage, departurearrivalconditionpropertytype)
	}
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureArrivalConditionPropertyTypeUpdateCallback != nil {
		var frontDepartureArrivalConditionPropertyType *DepartureArrivalConditionPropertyType
		if front != nil {
			frontDepartureArrivalConditionPropertyType, _ = front.(*DepartureArrivalConditionPropertyType)
		}
		stage.OnAfterDepartureArrivalConditionPropertyTypeUpdateCallback.OnAfterUpdate(stage, departurearrivalconditionpropertytype, frontDepartureArrivalConditionPropertyType)
	}
}

func (departurearrivalconditionpropertytype *DepartureArrivalConditionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureArrivalConditionPropertyTypeDeleteCallback != nil {
		var frontDepartureArrivalConditionPropertyType *DepartureArrivalConditionPropertyType
		if front != nil {
			frontDepartureArrivalConditionPropertyType, _ = front.(*DepartureArrivalConditionPropertyType)
		}
		stage.OnAfterDepartureArrivalConditionPropertyTypeDeleteCallback.OnAfterDelete(stage, departurearrivalconditionpropertytype, frontDepartureArrivalConditionPropertyType)
	}
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDepartureArrivalConditionTypeCreateCallback != nil {
		stage.OnAfterDepartureArrivalConditionTypeCreateCallback.OnAfterCreate(stage, departurearrivalconditiontype)
	}
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureArrivalConditionTypeUpdateCallback != nil {
		var frontDepartureArrivalConditionType *DepartureArrivalConditionType
		if front != nil {
			frontDepartureArrivalConditionType, _ = front.(*DepartureArrivalConditionType)
		}
		stage.OnAfterDepartureArrivalConditionTypeUpdateCallback.OnAfterUpdate(stage, departurearrivalconditiontype, frontDepartureArrivalConditionType)
	}
}

func (departurearrivalconditiontype *DepartureArrivalConditionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureArrivalConditionTypeDeleteCallback != nil {
		var frontDepartureArrivalConditionType *DepartureArrivalConditionType
		if front != nil {
			frontDepartureArrivalConditionType, _ = front.(*DepartureArrivalConditionType)
		}
		stage.OnAfterDepartureArrivalConditionTypeDeleteCallback.OnAfterDelete(stage, departurearrivalconditiontype, frontDepartureArrivalConditionType)
	}
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDepartureLegTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterDepartureLegTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, departurelegtimeslicepropertytype)
	}
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureLegTimeSlicePropertyTypeUpdateCallback != nil {
		var frontDepartureLegTimeSlicePropertyType *DepartureLegTimeSlicePropertyType
		if front != nil {
			frontDepartureLegTimeSlicePropertyType, _ = front.(*DepartureLegTimeSlicePropertyType)
		}
		stage.OnAfterDepartureLegTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, departurelegtimeslicepropertytype, frontDepartureLegTimeSlicePropertyType)
	}
}

func (departurelegtimeslicepropertytype *DepartureLegTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureLegTimeSlicePropertyTypeDeleteCallback != nil {
		var frontDepartureLegTimeSlicePropertyType *DepartureLegTimeSlicePropertyType
		if front != nil {
			frontDepartureLegTimeSlicePropertyType, _ = front.(*DepartureLegTimeSlicePropertyType)
		}
		stage.OnAfterDepartureLegTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, departurelegtimeslicepropertytype, frontDepartureLegTimeSlicePropertyType)
	}
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDepartureLegTimeSliceTypeCreateCallback != nil {
		stage.OnAfterDepartureLegTimeSliceTypeCreateCallback.OnAfterCreate(stage, departurelegtimeslicetype)
	}
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureLegTimeSliceTypeUpdateCallback != nil {
		var frontDepartureLegTimeSliceType *DepartureLegTimeSliceType
		if front != nil {
			frontDepartureLegTimeSliceType, _ = front.(*DepartureLegTimeSliceType)
		}
		stage.OnAfterDepartureLegTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, departurelegtimeslicetype, frontDepartureLegTimeSliceType)
	}
}

func (departurelegtimeslicetype *DepartureLegTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureLegTimeSliceTypeDeleteCallback != nil {
		var frontDepartureLegTimeSliceType *DepartureLegTimeSliceType
		if front != nil {
			frontDepartureLegTimeSliceType, _ = front.(*DepartureLegTimeSliceType)
		}
		stage.OnAfterDepartureLegTimeSliceTypeDeleteCallback.OnAfterDelete(stage, departurelegtimeslicetype, frontDepartureLegTimeSliceType)
	}
}

func (departurelegtype *DepartureLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDepartureLegTypeCreateCallback != nil {
		stage.OnAfterDepartureLegTypeCreateCallback.OnAfterCreate(stage, departurelegtype)
	}
}

func (departurelegtype *DepartureLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureLegTypeUpdateCallback != nil {
		var frontDepartureLegType *DepartureLegType
		if front != nil {
			frontDepartureLegType, _ = front.(*DepartureLegType)
		}
		stage.OnAfterDepartureLegTypeUpdateCallback.OnAfterUpdate(stage, departurelegtype, frontDepartureLegType)
	}
}

func (departurelegtype *DepartureLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDepartureLegTypeDeleteCallback != nil {
		var frontDepartureLegType *DepartureLegType
		if front != nil {
			frontDepartureLegType, _ = front.(*DepartureLegType)
		}
		stage.OnAfterDepartureLegTypeDeleteCallback.OnAfterDelete(stage, departurelegtype, frontDepartureLegType)
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

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDistanceIndicationPropertyTypeCreateCallback != nil {
		stage.OnAfterDistanceIndicationPropertyTypeCreateCallback.OnAfterCreate(stage, distanceindicationpropertytype)
	}
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceIndicationPropertyTypeUpdateCallback != nil {
		var frontDistanceIndicationPropertyType *DistanceIndicationPropertyType
		if front != nil {
			frontDistanceIndicationPropertyType, _ = front.(*DistanceIndicationPropertyType)
		}
		stage.OnAfterDistanceIndicationPropertyTypeUpdateCallback.OnAfterUpdate(stage, distanceindicationpropertytype, frontDistanceIndicationPropertyType)
	}
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistanceIndicationPropertyTypeDeleteCallback != nil {
		var frontDistanceIndicationPropertyType *DistanceIndicationPropertyType
		if front != nil {
			frontDistanceIndicationPropertyType, _ = front.(*DistanceIndicationPropertyType)
		}
		stage.OnAfterDistanceIndicationPropertyTypeDeleteCallback.OnAfterDelete(stage, distanceindicationpropertytype, frontDistanceIndicationPropertyType)
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

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEnRouteSegmentPointTypeCreateCallback != nil {
		stage.OnAfterEnRouteSegmentPointTypeCreateCallback.OnAfterCreate(stage, enroutesegmentpointtype)
	}
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnRouteSegmentPointTypeUpdateCallback != nil {
		var frontEnRouteSegmentPointType *EnRouteSegmentPointType
		if front != nil {
			frontEnRouteSegmentPointType, _ = front.(*EnRouteSegmentPointType)
		}
		stage.OnAfterEnRouteSegmentPointTypeUpdateCallback.OnAfterUpdate(stage, enroutesegmentpointtype, frontEnRouteSegmentPointType)
	}
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnRouteSegmentPointTypeDeleteCallback != nil {
		var frontEnRouteSegmentPointType *EnRouteSegmentPointType
		if front != nil {
			frontEnRouteSegmentPointType, _ = front.(*EnRouteSegmentPointType)
		}
		stage.OnAfterEnRouteSegmentPointTypeDeleteCallback.OnAfterDelete(stage, enroutesegmentpointtype, frontEnRouteSegmentPointType)
	}
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEquipmentUnavailableAdjustmentColumnPropertyTypeCreateCallback != nil {
		stage.OnAfterEquipmentUnavailableAdjustmentColumnPropertyTypeCreateCallback.OnAfterCreate(stage, equipmentunavailableadjustmentcolumnpropertytype)
	}
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEquipmentUnavailableAdjustmentColumnPropertyTypeUpdateCallback != nil {
		var frontEquipmentUnavailableAdjustmentColumnPropertyType *EquipmentUnavailableAdjustmentColumnPropertyType
		if front != nil {
			frontEquipmentUnavailableAdjustmentColumnPropertyType, _ = front.(*EquipmentUnavailableAdjustmentColumnPropertyType)
		}
		stage.OnAfterEquipmentUnavailableAdjustmentColumnPropertyTypeUpdateCallback.OnAfterUpdate(stage, equipmentunavailableadjustmentcolumnpropertytype, frontEquipmentUnavailableAdjustmentColumnPropertyType)
	}
}

func (equipmentunavailableadjustmentcolumnpropertytype *EquipmentUnavailableAdjustmentColumnPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEquipmentUnavailableAdjustmentColumnPropertyTypeDeleteCallback != nil {
		var frontEquipmentUnavailableAdjustmentColumnPropertyType *EquipmentUnavailableAdjustmentColumnPropertyType
		if front != nil {
			frontEquipmentUnavailableAdjustmentColumnPropertyType, _ = front.(*EquipmentUnavailableAdjustmentColumnPropertyType)
		}
		stage.OnAfterEquipmentUnavailableAdjustmentColumnPropertyTypeDeleteCallback.OnAfterDelete(stage, equipmentunavailableadjustmentcolumnpropertytype, frontEquipmentUnavailableAdjustmentColumnPropertyType)
	}
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEquipmentUnavailableAdjustmentColumnTypeCreateCallback != nil {
		stage.OnAfterEquipmentUnavailableAdjustmentColumnTypeCreateCallback.OnAfterCreate(stage, equipmentunavailableadjustmentcolumntype)
	}
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEquipmentUnavailableAdjustmentColumnTypeUpdateCallback != nil {
		var frontEquipmentUnavailableAdjustmentColumnType *EquipmentUnavailableAdjustmentColumnType
		if front != nil {
			frontEquipmentUnavailableAdjustmentColumnType, _ = front.(*EquipmentUnavailableAdjustmentColumnType)
		}
		stage.OnAfterEquipmentUnavailableAdjustmentColumnTypeUpdateCallback.OnAfterUpdate(stage, equipmentunavailableadjustmentcolumntype, frontEquipmentUnavailableAdjustmentColumnType)
	}
}

func (equipmentunavailableadjustmentcolumntype *EquipmentUnavailableAdjustmentColumnType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEquipmentUnavailableAdjustmentColumnTypeDeleteCallback != nil {
		var frontEquipmentUnavailableAdjustmentColumnType *EquipmentUnavailableAdjustmentColumnType
		if front != nil {
			frontEquipmentUnavailableAdjustmentColumnType, _ = front.(*EquipmentUnavailableAdjustmentColumnType)
		}
		stage.OnAfterEquipmentUnavailableAdjustmentColumnTypeDeleteCallback.OnAfterDelete(stage, equipmentunavailableadjustmentcolumntype, frontEquipmentUnavailableAdjustmentColumnType)
	}
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEquipmentUnavailableAdjustmentPropertyTypeCreateCallback != nil {
		stage.OnAfterEquipmentUnavailableAdjustmentPropertyTypeCreateCallback.OnAfterCreate(stage, equipmentunavailableadjustmentpropertytype)
	}
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEquipmentUnavailableAdjustmentPropertyTypeUpdateCallback != nil {
		var frontEquipmentUnavailableAdjustmentPropertyType *EquipmentUnavailableAdjustmentPropertyType
		if front != nil {
			frontEquipmentUnavailableAdjustmentPropertyType, _ = front.(*EquipmentUnavailableAdjustmentPropertyType)
		}
		stage.OnAfterEquipmentUnavailableAdjustmentPropertyTypeUpdateCallback.OnAfterUpdate(stage, equipmentunavailableadjustmentpropertytype, frontEquipmentUnavailableAdjustmentPropertyType)
	}
}

func (equipmentunavailableadjustmentpropertytype *EquipmentUnavailableAdjustmentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEquipmentUnavailableAdjustmentPropertyTypeDeleteCallback != nil {
		var frontEquipmentUnavailableAdjustmentPropertyType *EquipmentUnavailableAdjustmentPropertyType
		if front != nil {
			frontEquipmentUnavailableAdjustmentPropertyType, _ = front.(*EquipmentUnavailableAdjustmentPropertyType)
		}
		stage.OnAfterEquipmentUnavailableAdjustmentPropertyTypeDeleteCallback.OnAfterDelete(stage, equipmentunavailableadjustmentpropertytype, frontEquipmentUnavailableAdjustmentPropertyType)
	}
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEquipmentUnavailableAdjustmentTypeCreateCallback != nil {
		stage.OnAfterEquipmentUnavailableAdjustmentTypeCreateCallback.OnAfterCreate(stage, equipmentunavailableadjustmenttype)
	}
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEquipmentUnavailableAdjustmentTypeUpdateCallback != nil {
		var frontEquipmentUnavailableAdjustmentType *EquipmentUnavailableAdjustmentType
		if front != nil {
			frontEquipmentUnavailableAdjustmentType, _ = front.(*EquipmentUnavailableAdjustmentType)
		}
		stage.OnAfterEquipmentUnavailableAdjustmentTypeUpdateCallback.OnAfterUpdate(stage, equipmentunavailableadjustmenttype, frontEquipmentUnavailableAdjustmentType)
	}
}

func (equipmentunavailableadjustmenttype *EquipmentUnavailableAdjustmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEquipmentUnavailableAdjustmentTypeDeleteCallback != nil {
		var frontEquipmentUnavailableAdjustmentType *EquipmentUnavailableAdjustmentType
		if front != nil {
			frontEquipmentUnavailableAdjustmentType, _ = front.(*EquipmentUnavailableAdjustmentType)
		}
		stage.OnAfterEquipmentUnavailableAdjustmentTypeDeleteCallback.OnAfterDelete(stage, equipmentunavailableadjustmenttype, frontEquipmentUnavailableAdjustmentType)
	}
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFASDataBlockPropertyTypeCreateCallback != nil {
		stage.OnAfterFASDataBlockPropertyTypeCreateCallback.OnAfterCreate(stage, fasdatablockpropertytype)
	}
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFASDataBlockPropertyTypeUpdateCallback != nil {
		var frontFASDataBlockPropertyType *FASDataBlockPropertyType
		if front != nil {
			frontFASDataBlockPropertyType, _ = front.(*FASDataBlockPropertyType)
		}
		stage.OnAfterFASDataBlockPropertyTypeUpdateCallback.OnAfterUpdate(stage, fasdatablockpropertytype, frontFASDataBlockPropertyType)
	}
}

func (fasdatablockpropertytype *FASDataBlockPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFASDataBlockPropertyTypeDeleteCallback != nil {
		var frontFASDataBlockPropertyType *FASDataBlockPropertyType
		if front != nil {
			frontFASDataBlockPropertyType, _ = front.(*FASDataBlockPropertyType)
		}
		stage.OnAfterFASDataBlockPropertyTypeDeleteCallback.OnAfterDelete(stage, fasdatablockpropertytype, frontFASDataBlockPropertyType)
	}
}

func (fasdatablocktype *FASDataBlockType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFASDataBlockTypeCreateCallback != nil {
		stage.OnAfterFASDataBlockTypeCreateCallback.OnAfterCreate(stage, fasdatablocktype)
	}
}

func (fasdatablocktype *FASDataBlockType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFASDataBlockTypeUpdateCallback != nil {
		var frontFASDataBlockType *FASDataBlockType
		if front != nil {
			frontFASDataBlockType, _ = front.(*FASDataBlockType)
		}
		stage.OnAfterFASDataBlockTypeUpdateCallback.OnAfterUpdate(stage, fasdatablocktype, frontFASDataBlockType)
	}
}

func (fasdatablocktype *FASDataBlockType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFASDataBlockTypeDeleteCallback != nil {
		var frontFASDataBlockType *FASDataBlockType
		if front != nil {
			frontFASDataBlockType, _ = front.(*FASDataBlockType)
		}
		stage.OnAfterFASDataBlockTypeDeleteCallback.OnAfterDelete(stage, fasdatablocktype, frontFASDataBlockType)
	}
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFinalLegTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterFinalLegTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, finallegtimeslicepropertytype)
	}
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalLegTimeSlicePropertyTypeUpdateCallback != nil {
		var frontFinalLegTimeSlicePropertyType *FinalLegTimeSlicePropertyType
		if front != nil {
			frontFinalLegTimeSlicePropertyType, _ = front.(*FinalLegTimeSlicePropertyType)
		}
		stage.OnAfterFinalLegTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, finallegtimeslicepropertytype, frontFinalLegTimeSlicePropertyType)
	}
}

func (finallegtimeslicepropertytype *FinalLegTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalLegTimeSlicePropertyTypeDeleteCallback != nil {
		var frontFinalLegTimeSlicePropertyType *FinalLegTimeSlicePropertyType
		if front != nil {
			frontFinalLegTimeSlicePropertyType, _ = front.(*FinalLegTimeSlicePropertyType)
		}
		stage.OnAfterFinalLegTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, finallegtimeslicepropertytype, frontFinalLegTimeSlicePropertyType)
	}
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFinalLegTimeSliceTypeCreateCallback != nil {
		stage.OnAfterFinalLegTimeSliceTypeCreateCallback.OnAfterCreate(stage, finallegtimeslicetype)
	}
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalLegTimeSliceTypeUpdateCallback != nil {
		var frontFinalLegTimeSliceType *FinalLegTimeSliceType
		if front != nil {
			frontFinalLegTimeSliceType, _ = front.(*FinalLegTimeSliceType)
		}
		stage.OnAfterFinalLegTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, finallegtimeslicetype, frontFinalLegTimeSliceType)
	}
}

func (finallegtimeslicetype *FinalLegTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalLegTimeSliceTypeDeleteCallback != nil {
		var frontFinalLegTimeSliceType *FinalLegTimeSliceType
		if front != nil {
			frontFinalLegTimeSliceType, _ = front.(*FinalLegTimeSliceType)
		}
		stage.OnAfterFinalLegTimeSliceTypeDeleteCallback.OnAfterDelete(stage, finallegtimeslicetype, frontFinalLegTimeSliceType)
	}
}

func (finallegtype *FinalLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFinalLegTypeCreateCallback != nil {
		stage.OnAfterFinalLegTypeCreateCallback.OnAfterCreate(stage, finallegtype)
	}
}

func (finallegtype *FinalLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalLegTypeUpdateCallback != nil {
		var frontFinalLegType *FinalLegType
		if front != nil {
			frontFinalLegType, _ = front.(*FinalLegType)
		}
		stage.OnAfterFinalLegTypeUpdateCallback.OnAfterUpdate(stage, finallegtype, frontFinalLegType)
	}
}

func (finallegtype *FinalLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalLegTypeDeleteCallback != nil {
		var frontFinalLegType *FinalLegType
		if front != nil {
			frontFinalLegType, _ = front.(*FinalLegType)
		}
		stage.OnAfterFinalLegTypeDeleteCallback.OnAfterDelete(stage, finallegtype, frontFinalLegType)
	}
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFinalProfilePropertyTypeCreateCallback != nil {
		stage.OnAfterFinalProfilePropertyTypeCreateCallback.OnAfterCreate(stage, finalprofilepropertytype)
	}
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalProfilePropertyTypeUpdateCallback != nil {
		var frontFinalProfilePropertyType *FinalProfilePropertyType
		if front != nil {
			frontFinalProfilePropertyType, _ = front.(*FinalProfilePropertyType)
		}
		stage.OnAfterFinalProfilePropertyTypeUpdateCallback.OnAfterUpdate(stage, finalprofilepropertytype, frontFinalProfilePropertyType)
	}
}

func (finalprofilepropertytype *FinalProfilePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalProfilePropertyTypeDeleteCallback != nil {
		var frontFinalProfilePropertyType *FinalProfilePropertyType
		if front != nil {
			frontFinalProfilePropertyType, _ = front.(*FinalProfilePropertyType)
		}
		stage.OnAfterFinalProfilePropertyTypeDeleteCallback.OnAfterDelete(stage, finalprofilepropertytype, frontFinalProfilePropertyType)
	}
}

func (finalprofiletype *FinalProfileType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFinalProfileTypeCreateCallback != nil {
		stage.OnAfterFinalProfileTypeCreateCallback.OnAfterCreate(stage, finalprofiletype)
	}
}

func (finalprofiletype *FinalProfileType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalProfileTypeUpdateCallback != nil {
		var frontFinalProfileType *FinalProfileType
		if front != nil {
			frontFinalProfileType, _ = front.(*FinalProfileType)
		}
		stage.OnAfterFinalProfileTypeUpdateCallback.OnAfterUpdate(stage, finalprofiletype, frontFinalProfileType)
	}
}

func (finalprofiletype *FinalProfileType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFinalProfileTypeDeleteCallback != nil {
		var frontFinalProfileType *FinalProfileType
		if front != nil {
			frontFinalProfileType, _ = front.(*FinalProfileType)
		}
		stage.OnAfterFinalProfileTypeDeleteCallback.OnAfterDelete(stage, finalprofiletype, frontFinalProfileType)
	}
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingAssessmentTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterHoldingAssessmentTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, holdingassessmenttimeslicepropertytype)
	}
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingAssessmentTimeSlicePropertyTypeUpdateCallback != nil {
		var frontHoldingAssessmentTimeSlicePropertyType *HoldingAssessmentTimeSlicePropertyType
		if front != nil {
			frontHoldingAssessmentTimeSlicePropertyType, _ = front.(*HoldingAssessmentTimeSlicePropertyType)
		}
		stage.OnAfterHoldingAssessmentTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, holdingassessmenttimeslicepropertytype, frontHoldingAssessmentTimeSlicePropertyType)
	}
}

func (holdingassessmenttimeslicepropertytype *HoldingAssessmentTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingAssessmentTimeSlicePropertyTypeDeleteCallback != nil {
		var frontHoldingAssessmentTimeSlicePropertyType *HoldingAssessmentTimeSlicePropertyType
		if front != nil {
			frontHoldingAssessmentTimeSlicePropertyType, _ = front.(*HoldingAssessmentTimeSlicePropertyType)
		}
		stage.OnAfterHoldingAssessmentTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, holdingassessmenttimeslicepropertytype, frontHoldingAssessmentTimeSlicePropertyType)
	}
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingAssessmentTimeSliceTypeCreateCallback != nil {
		stage.OnAfterHoldingAssessmentTimeSliceTypeCreateCallback.OnAfterCreate(stage, holdingassessmenttimeslicetype)
	}
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingAssessmentTimeSliceTypeUpdateCallback != nil {
		var frontHoldingAssessmentTimeSliceType *HoldingAssessmentTimeSliceType
		if front != nil {
			frontHoldingAssessmentTimeSliceType, _ = front.(*HoldingAssessmentTimeSliceType)
		}
		stage.OnAfterHoldingAssessmentTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, holdingassessmenttimeslicetype, frontHoldingAssessmentTimeSliceType)
	}
}

func (holdingassessmenttimeslicetype *HoldingAssessmentTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingAssessmentTimeSliceTypeDeleteCallback != nil {
		var frontHoldingAssessmentTimeSliceType *HoldingAssessmentTimeSliceType
		if front != nil {
			frontHoldingAssessmentTimeSliceType, _ = front.(*HoldingAssessmentTimeSliceType)
		}
		stage.OnAfterHoldingAssessmentTimeSliceTypeDeleteCallback.OnAfterDelete(stage, holdingassessmenttimeslicetype, frontHoldingAssessmentTimeSliceType)
	}
}

func (holdingassessmenttype *HoldingAssessmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingAssessmentTypeCreateCallback != nil {
		stage.OnAfterHoldingAssessmentTypeCreateCallback.OnAfterCreate(stage, holdingassessmenttype)
	}
}

func (holdingassessmenttype *HoldingAssessmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingAssessmentTypeUpdateCallback != nil {
		var frontHoldingAssessmentType *HoldingAssessmentType
		if front != nil {
			frontHoldingAssessmentType, _ = front.(*HoldingAssessmentType)
		}
		stage.OnAfterHoldingAssessmentTypeUpdateCallback.OnAfterUpdate(stage, holdingassessmenttype, frontHoldingAssessmentType)
	}
}

func (holdingassessmenttype *HoldingAssessmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingAssessmentTypeDeleteCallback != nil {
		var frontHoldingAssessmentType *HoldingAssessmentType
		if front != nil {
			frontHoldingAssessmentType, _ = front.(*HoldingAssessmentType)
		}
		stage.OnAfterHoldingAssessmentTypeDeleteCallback.OnAfterDelete(stage, holdingassessmenttype, frontHoldingAssessmentType)
	}
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingPatternDistancePropertyTypeCreateCallback != nil {
		stage.OnAfterHoldingPatternDistancePropertyTypeCreateCallback.OnAfterCreate(stage, holdingpatterndistancepropertytype)
	}
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternDistancePropertyTypeUpdateCallback != nil {
		var frontHoldingPatternDistancePropertyType *HoldingPatternDistancePropertyType
		if front != nil {
			frontHoldingPatternDistancePropertyType, _ = front.(*HoldingPatternDistancePropertyType)
		}
		stage.OnAfterHoldingPatternDistancePropertyTypeUpdateCallback.OnAfterUpdate(stage, holdingpatterndistancepropertytype, frontHoldingPatternDistancePropertyType)
	}
}

func (holdingpatterndistancepropertytype *HoldingPatternDistancePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternDistancePropertyTypeDeleteCallback != nil {
		var frontHoldingPatternDistancePropertyType *HoldingPatternDistancePropertyType
		if front != nil {
			frontHoldingPatternDistancePropertyType, _ = front.(*HoldingPatternDistancePropertyType)
		}
		stage.OnAfterHoldingPatternDistancePropertyTypeDeleteCallback.OnAfterDelete(stage, holdingpatterndistancepropertytype, frontHoldingPatternDistancePropertyType)
	}
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingPatternDistanceTypeCreateCallback != nil {
		stage.OnAfterHoldingPatternDistanceTypeCreateCallback.OnAfterCreate(stage, holdingpatterndistancetype)
	}
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternDistanceTypeUpdateCallback != nil {
		var frontHoldingPatternDistanceType *HoldingPatternDistanceType
		if front != nil {
			frontHoldingPatternDistanceType, _ = front.(*HoldingPatternDistanceType)
		}
		stage.OnAfterHoldingPatternDistanceTypeUpdateCallback.OnAfterUpdate(stage, holdingpatterndistancetype, frontHoldingPatternDistanceType)
	}
}

func (holdingpatterndistancetype *HoldingPatternDistanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternDistanceTypeDeleteCallback != nil {
		var frontHoldingPatternDistanceType *HoldingPatternDistanceType
		if front != nil {
			frontHoldingPatternDistanceType, _ = front.(*HoldingPatternDistanceType)
		}
		stage.OnAfterHoldingPatternDistanceTypeDeleteCallback.OnAfterDelete(stage, holdingpatterndistancetype, frontHoldingPatternDistanceType)
	}
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingPatternDurationPropertyTypeCreateCallback != nil {
		stage.OnAfterHoldingPatternDurationPropertyTypeCreateCallback.OnAfterCreate(stage, holdingpatterndurationpropertytype)
	}
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternDurationPropertyTypeUpdateCallback != nil {
		var frontHoldingPatternDurationPropertyType *HoldingPatternDurationPropertyType
		if front != nil {
			frontHoldingPatternDurationPropertyType, _ = front.(*HoldingPatternDurationPropertyType)
		}
		stage.OnAfterHoldingPatternDurationPropertyTypeUpdateCallback.OnAfterUpdate(stage, holdingpatterndurationpropertytype, frontHoldingPatternDurationPropertyType)
	}
}

func (holdingpatterndurationpropertytype *HoldingPatternDurationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternDurationPropertyTypeDeleteCallback != nil {
		var frontHoldingPatternDurationPropertyType *HoldingPatternDurationPropertyType
		if front != nil {
			frontHoldingPatternDurationPropertyType, _ = front.(*HoldingPatternDurationPropertyType)
		}
		stage.OnAfterHoldingPatternDurationPropertyTypeDeleteCallback.OnAfterDelete(stage, holdingpatterndurationpropertytype, frontHoldingPatternDurationPropertyType)
	}
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingPatternDurationTypeCreateCallback != nil {
		stage.OnAfterHoldingPatternDurationTypeCreateCallback.OnAfterCreate(stage, holdingpatterndurationtype)
	}
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternDurationTypeUpdateCallback != nil {
		var frontHoldingPatternDurationType *HoldingPatternDurationType
		if front != nil {
			frontHoldingPatternDurationType, _ = front.(*HoldingPatternDurationType)
		}
		stage.OnAfterHoldingPatternDurationTypeUpdateCallback.OnAfterUpdate(stage, holdingpatterndurationtype, frontHoldingPatternDurationType)
	}
}

func (holdingpatterndurationtype *HoldingPatternDurationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternDurationTypeDeleteCallback != nil {
		var frontHoldingPatternDurationType *HoldingPatternDurationType
		if front != nil {
			frontHoldingPatternDurationType, _ = front.(*HoldingPatternDurationType)
		}
		stage.OnAfterHoldingPatternDurationTypeDeleteCallback.OnAfterDelete(stage, holdingpatterndurationtype, frontHoldingPatternDurationType)
	}
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingPatternPropertyTypeCreateCallback != nil {
		stage.OnAfterHoldingPatternPropertyTypeCreateCallback.OnAfterCreate(stage, holdingpatternpropertytype)
	}
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternPropertyTypeUpdateCallback != nil {
		var frontHoldingPatternPropertyType *HoldingPatternPropertyType
		if front != nil {
			frontHoldingPatternPropertyType, _ = front.(*HoldingPatternPropertyType)
		}
		stage.OnAfterHoldingPatternPropertyTypeUpdateCallback.OnAfterUpdate(stage, holdingpatternpropertytype, frontHoldingPatternPropertyType)
	}
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternPropertyTypeDeleteCallback != nil {
		var frontHoldingPatternPropertyType *HoldingPatternPropertyType
		if front != nil {
			frontHoldingPatternPropertyType, _ = front.(*HoldingPatternPropertyType)
		}
		stage.OnAfterHoldingPatternPropertyTypeDeleteCallback.OnAfterDelete(stage, holdingpatternpropertytype, frontHoldingPatternPropertyType)
	}
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingPatternTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterHoldingPatternTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, holdingpatterntimeslicepropertytype)
	}
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternTimeSlicePropertyTypeUpdateCallback != nil {
		var frontHoldingPatternTimeSlicePropertyType *HoldingPatternTimeSlicePropertyType
		if front != nil {
			frontHoldingPatternTimeSlicePropertyType, _ = front.(*HoldingPatternTimeSlicePropertyType)
		}
		stage.OnAfterHoldingPatternTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, holdingpatterntimeslicepropertytype, frontHoldingPatternTimeSlicePropertyType)
	}
}

func (holdingpatterntimeslicepropertytype *HoldingPatternTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternTimeSlicePropertyTypeDeleteCallback != nil {
		var frontHoldingPatternTimeSlicePropertyType *HoldingPatternTimeSlicePropertyType
		if front != nil {
			frontHoldingPatternTimeSlicePropertyType, _ = front.(*HoldingPatternTimeSlicePropertyType)
		}
		stage.OnAfterHoldingPatternTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, holdingpatterntimeslicepropertytype, frontHoldingPatternTimeSlicePropertyType)
	}
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingPatternTimeSliceTypeCreateCallback != nil {
		stage.OnAfterHoldingPatternTimeSliceTypeCreateCallback.OnAfterCreate(stage, holdingpatterntimeslicetype)
	}
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternTimeSliceTypeUpdateCallback != nil {
		var frontHoldingPatternTimeSliceType *HoldingPatternTimeSliceType
		if front != nil {
			frontHoldingPatternTimeSliceType, _ = front.(*HoldingPatternTimeSliceType)
		}
		stage.OnAfterHoldingPatternTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, holdingpatterntimeslicetype, frontHoldingPatternTimeSliceType)
	}
}

func (holdingpatterntimeslicetype *HoldingPatternTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternTimeSliceTypeDeleteCallback != nil {
		var frontHoldingPatternTimeSliceType *HoldingPatternTimeSliceType
		if front != nil {
			frontHoldingPatternTimeSliceType, _ = front.(*HoldingPatternTimeSliceType)
		}
		stage.OnAfterHoldingPatternTimeSliceTypeDeleteCallback.OnAfterDelete(stage, holdingpatterntimeslicetype, frontHoldingPatternTimeSliceType)
	}
}

func (holdingpatterntype *HoldingPatternType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingPatternTypeCreateCallback != nil {
		stage.OnAfterHoldingPatternTypeCreateCallback.OnAfterCreate(stage, holdingpatterntype)
	}
}

func (holdingpatterntype *HoldingPatternType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternTypeUpdateCallback != nil {
		var frontHoldingPatternType *HoldingPatternType
		if front != nil {
			frontHoldingPatternType, _ = front.(*HoldingPatternType)
		}
		stage.OnAfterHoldingPatternTypeUpdateCallback.OnAfterUpdate(stage, holdingpatterntype, frontHoldingPatternType)
	}
}

func (holdingpatterntype *HoldingPatternType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingPatternTypeDeleteCallback != nil {
		var frontHoldingPatternType *HoldingPatternType
		if front != nil {
			frontHoldingPatternType, _ = front.(*HoldingPatternType)
		}
		stage.OnAfterHoldingPatternTypeDeleteCallback.OnAfterDelete(stage, holdingpatterntype, frontHoldingPatternType)
	}
}

func (holdingusepropertytype *HoldingUsePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingUsePropertyTypeCreateCallback != nil {
		stage.OnAfterHoldingUsePropertyTypeCreateCallback.OnAfterCreate(stage, holdingusepropertytype)
	}
}

func (holdingusepropertytype *HoldingUsePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingUsePropertyTypeUpdateCallback != nil {
		var frontHoldingUsePropertyType *HoldingUsePropertyType
		if front != nil {
			frontHoldingUsePropertyType, _ = front.(*HoldingUsePropertyType)
		}
		stage.OnAfterHoldingUsePropertyTypeUpdateCallback.OnAfterUpdate(stage, holdingusepropertytype, frontHoldingUsePropertyType)
	}
}

func (holdingusepropertytype *HoldingUsePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingUsePropertyTypeDeleteCallback != nil {
		var frontHoldingUsePropertyType *HoldingUsePropertyType
		if front != nil {
			frontHoldingUsePropertyType, _ = front.(*HoldingUsePropertyType)
		}
		stage.OnAfterHoldingUsePropertyTypeDeleteCallback.OnAfterDelete(stage, holdingusepropertytype, frontHoldingUsePropertyType)
	}
}

func (holdingusetype *HoldingUseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHoldingUseTypeCreateCallback != nil {
		stage.OnAfterHoldingUseTypeCreateCallback.OnAfterCreate(stage, holdingusetype)
	}
}

func (holdingusetype *HoldingUseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingUseTypeUpdateCallback != nil {
		var frontHoldingUseType *HoldingUseType
		if front != nil {
			frontHoldingUseType, _ = front.(*HoldingUseType)
		}
		stage.OnAfterHoldingUseTypeUpdateCallback.OnAfterUpdate(stage, holdingusetype, frontHoldingUseType)
	}
}

func (holdingusetype *HoldingUseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHoldingUseTypeDeleteCallback != nil {
		var frontHoldingUseType *HoldingUseType
		if front != nil {
			frontHoldingUseType, _ = front.(*HoldingUseType)
		}
		stage.OnAfterHoldingUseTypeDeleteCallback.OnAfterDelete(stage, holdingusetype, frontHoldingUseType)
	}
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInitialLegTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterInitialLegTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, initiallegtimeslicepropertytype)
	}
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialLegTimeSlicePropertyTypeUpdateCallback != nil {
		var frontInitialLegTimeSlicePropertyType *InitialLegTimeSlicePropertyType
		if front != nil {
			frontInitialLegTimeSlicePropertyType, _ = front.(*InitialLegTimeSlicePropertyType)
		}
		stage.OnAfterInitialLegTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, initiallegtimeslicepropertytype, frontInitialLegTimeSlicePropertyType)
	}
}

func (initiallegtimeslicepropertytype *InitialLegTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialLegTimeSlicePropertyTypeDeleteCallback != nil {
		var frontInitialLegTimeSlicePropertyType *InitialLegTimeSlicePropertyType
		if front != nil {
			frontInitialLegTimeSlicePropertyType, _ = front.(*InitialLegTimeSlicePropertyType)
		}
		stage.OnAfterInitialLegTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, initiallegtimeslicepropertytype, frontInitialLegTimeSlicePropertyType)
	}
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInitialLegTimeSliceTypeCreateCallback != nil {
		stage.OnAfterInitialLegTimeSliceTypeCreateCallback.OnAfterCreate(stage, initiallegtimeslicetype)
	}
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialLegTimeSliceTypeUpdateCallback != nil {
		var frontInitialLegTimeSliceType *InitialLegTimeSliceType
		if front != nil {
			frontInitialLegTimeSliceType, _ = front.(*InitialLegTimeSliceType)
		}
		stage.OnAfterInitialLegTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, initiallegtimeslicetype, frontInitialLegTimeSliceType)
	}
}

func (initiallegtimeslicetype *InitialLegTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialLegTimeSliceTypeDeleteCallback != nil {
		var frontInitialLegTimeSliceType *InitialLegTimeSliceType
		if front != nil {
			frontInitialLegTimeSliceType, _ = front.(*InitialLegTimeSliceType)
		}
		stage.OnAfterInitialLegTimeSliceTypeDeleteCallback.OnAfterDelete(stage, initiallegtimeslicetype, frontInitialLegTimeSliceType)
	}
}

func (initiallegtype *InitialLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInitialLegTypeCreateCallback != nil {
		stage.OnAfterInitialLegTypeCreateCallback.OnAfterCreate(stage, initiallegtype)
	}
}

func (initiallegtype *InitialLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialLegTypeUpdateCallback != nil {
		var frontInitialLegType *InitialLegType
		if front != nil {
			frontInitialLegType, _ = front.(*InitialLegType)
		}
		stage.OnAfterInitialLegTypeUpdateCallback.OnAfterUpdate(stage, initiallegtype, frontInitialLegType)
	}
}

func (initiallegtype *InitialLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialLegTypeDeleteCallback != nil {
		var frontInitialLegType *InitialLegType
		if front != nil {
			frontInitialLegType, _ = front.(*InitialLegType)
		}
		stage.OnAfterInitialLegTypeDeleteCallback.OnAfterDelete(stage, initiallegtype, frontInitialLegType)
	}
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInstrumentApproachProcedurePropertyTypeCreateCallback != nil {
		stage.OnAfterInstrumentApproachProcedurePropertyTypeCreateCallback.OnAfterCreate(stage, instrumentapproachprocedurepropertytype)
	}
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentApproachProcedurePropertyTypeUpdateCallback != nil {
		var frontInstrumentApproachProcedurePropertyType *InstrumentApproachProcedurePropertyType
		if front != nil {
			frontInstrumentApproachProcedurePropertyType, _ = front.(*InstrumentApproachProcedurePropertyType)
		}
		stage.OnAfterInstrumentApproachProcedurePropertyTypeUpdateCallback.OnAfterUpdate(stage, instrumentapproachprocedurepropertytype, frontInstrumentApproachProcedurePropertyType)
	}
}

func (instrumentapproachprocedurepropertytype *InstrumentApproachProcedurePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentApproachProcedurePropertyTypeDeleteCallback != nil {
		var frontInstrumentApproachProcedurePropertyType *InstrumentApproachProcedurePropertyType
		if front != nil {
			frontInstrumentApproachProcedurePropertyType, _ = front.(*InstrumentApproachProcedurePropertyType)
		}
		stage.OnAfterInstrumentApproachProcedurePropertyTypeDeleteCallback.OnAfterDelete(stage, instrumentapproachprocedurepropertytype, frontInstrumentApproachProcedurePropertyType)
	}
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInstrumentApproachProcedureTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterInstrumentApproachProcedureTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, instrumentapproachproceduretimeslicepropertytype)
	}
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentApproachProcedureTimeSlicePropertyTypeUpdateCallback != nil {
		var frontInstrumentApproachProcedureTimeSlicePropertyType *InstrumentApproachProcedureTimeSlicePropertyType
		if front != nil {
			frontInstrumentApproachProcedureTimeSlicePropertyType, _ = front.(*InstrumentApproachProcedureTimeSlicePropertyType)
		}
		stage.OnAfterInstrumentApproachProcedureTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, instrumentapproachproceduretimeslicepropertytype, frontInstrumentApproachProcedureTimeSlicePropertyType)
	}
}

func (instrumentapproachproceduretimeslicepropertytype *InstrumentApproachProcedureTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentApproachProcedureTimeSlicePropertyTypeDeleteCallback != nil {
		var frontInstrumentApproachProcedureTimeSlicePropertyType *InstrumentApproachProcedureTimeSlicePropertyType
		if front != nil {
			frontInstrumentApproachProcedureTimeSlicePropertyType, _ = front.(*InstrumentApproachProcedureTimeSlicePropertyType)
		}
		stage.OnAfterInstrumentApproachProcedureTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, instrumentapproachproceduretimeslicepropertytype, frontInstrumentApproachProcedureTimeSlicePropertyType)
	}
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInstrumentApproachProcedureTimeSliceTypeCreateCallback != nil {
		stage.OnAfterInstrumentApproachProcedureTimeSliceTypeCreateCallback.OnAfterCreate(stage, instrumentapproachproceduretimeslicetype)
	}
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentApproachProcedureTimeSliceTypeUpdateCallback != nil {
		var frontInstrumentApproachProcedureTimeSliceType *InstrumentApproachProcedureTimeSliceType
		if front != nil {
			frontInstrumentApproachProcedureTimeSliceType, _ = front.(*InstrumentApproachProcedureTimeSliceType)
		}
		stage.OnAfterInstrumentApproachProcedureTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, instrumentapproachproceduretimeslicetype, frontInstrumentApproachProcedureTimeSliceType)
	}
}

func (instrumentapproachproceduretimeslicetype *InstrumentApproachProcedureTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentApproachProcedureTimeSliceTypeDeleteCallback != nil {
		var frontInstrumentApproachProcedureTimeSliceType *InstrumentApproachProcedureTimeSliceType
		if front != nil {
			frontInstrumentApproachProcedureTimeSliceType, _ = front.(*InstrumentApproachProcedureTimeSliceType)
		}
		stage.OnAfterInstrumentApproachProcedureTimeSliceTypeDeleteCallback.OnAfterDelete(stage, instrumentapproachproceduretimeslicetype, frontInstrumentApproachProcedureTimeSliceType)
	}
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInstrumentApproachProcedureTypeCreateCallback != nil {
		stage.OnAfterInstrumentApproachProcedureTypeCreateCallback.OnAfterCreate(stage, instrumentapproachproceduretype)
	}
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentApproachProcedureTypeUpdateCallback != nil {
		var frontInstrumentApproachProcedureType *InstrumentApproachProcedureType
		if front != nil {
			frontInstrumentApproachProcedureType, _ = front.(*InstrumentApproachProcedureType)
		}
		stage.OnAfterInstrumentApproachProcedureTypeUpdateCallback.OnAfterUpdate(stage, instrumentapproachproceduretype, frontInstrumentApproachProcedureType)
	}
}

func (instrumentapproachproceduretype *InstrumentApproachProcedureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInstrumentApproachProcedureTypeDeleteCallback != nil {
		var frontInstrumentApproachProcedureType *InstrumentApproachProcedureType
		if front != nil {
			frontInstrumentApproachProcedureType, _ = front.(*InstrumentApproachProcedureType)
		}
		stage.OnAfterInstrumentApproachProcedureTypeDeleteCallback.OnAfterDelete(stage, instrumentapproachproceduretype, frontInstrumentApproachProcedureType)
	}
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterIntermediateLegTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterIntermediateLegTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, intermediatelegtimeslicepropertytype)
	}
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIntermediateLegTimeSlicePropertyTypeUpdateCallback != nil {
		var frontIntermediateLegTimeSlicePropertyType *IntermediateLegTimeSlicePropertyType
		if front != nil {
			frontIntermediateLegTimeSlicePropertyType, _ = front.(*IntermediateLegTimeSlicePropertyType)
		}
		stage.OnAfterIntermediateLegTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, intermediatelegtimeslicepropertytype, frontIntermediateLegTimeSlicePropertyType)
	}
}

func (intermediatelegtimeslicepropertytype *IntermediateLegTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIntermediateLegTimeSlicePropertyTypeDeleteCallback != nil {
		var frontIntermediateLegTimeSlicePropertyType *IntermediateLegTimeSlicePropertyType
		if front != nil {
			frontIntermediateLegTimeSlicePropertyType, _ = front.(*IntermediateLegTimeSlicePropertyType)
		}
		stage.OnAfterIntermediateLegTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, intermediatelegtimeslicepropertytype, frontIntermediateLegTimeSlicePropertyType)
	}
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterIntermediateLegTimeSliceTypeCreateCallback != nil {
		stage.OnAfterIntermediateLegTimeSliceTypeCreateCallback.OnAfterCreate(stage, intermediatelegtimeslicetype)
	}
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIntermediateLegTimeSliceTypeUpdateCallback != nil {
		var frontIntermediateLegTimeSliceType *IntermediateLegTimeSliceType
		if front != nil {
			frontIntermediateLegTimeSliceType, _ = front.(*IntermediateLegTimeSliceType)
		}
		stage.OnAfterIntermediateLegTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, intermediatelegtimeslicetype, frontIntermediateLegTimeSliceType)
	}
}

func (intermediatelegtimeslicetype *IntermediateLegTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIntermediateLegTimeSliceTypeDeleteCallback != nil {
		var frontIntermediateLegTimeSliceType *IntermediateLegTimeSliceType
		if front != nil {
			frontIntermediateLegTimeSliceType, _ = front.(*IntermediateLegTimeSliceType)
		}
		stage.OnAfterIntermediateLegTimeSliceTypeDeleteCallback.OnAfterDelete(stage, intermediatelegtimeslicetype, frontIntermediateLegTimeSliceType)
	}
}

func (intermediatelegtype *IntermediateLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterIntermediateLegTypeCreateCallback != nil {
		stage.OnAfterIntermediateLegTypeCreateCallback.OnAfterCreate(stage, intermediatelegtype)
	}
}

func (intermediatelegtype *IntermediateLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIntermediateLegTypeUpdateCallback != nil {
		var frontIntermediateLegType *IntermediateLegType
		if front != nil {
			frontIntermediateLegType, _ = front.(*IntermediateLegType)
		}
		stage.OnAfterIntermediateLegTypeUpdateCallback.OnAfterUpdate(stage, intermediatelegtype, frontIntermediateLegType)
	}
}

func (intermediatelegtype *IntermediateLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIntermediateLegTypeDeleteCallback != nil {
		var frontIntermediateLegType *IntermediateLegType
		if front != nil {
			frontIntermediateLegType, _ = front.(*IntermediateLegType)
		}
		stage.OnAfterIntermediateLegTypeDeleteCallback.OnAfterDelete(stage, intermediatelegtype, frontIntermediateLegType)
	}
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLandingTakeoffAreaCollectionPropertyTypeCreateCallback != nil {
		stage.OnAfterLandingTakeoffAreaCollectionPropertyTypeCreateCallback.OnAfterCreate(stage, landingtakeoffareacollectionpropertytype)
	}
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLandingTakeoffAreaCollectionPropertyTypeUpdateCallback != nil {
		var frontLandingTakeoffAreaCollectionPropertyType *LandingTakeoffAreaCollectionPropertyType
		if front != nil {
			frontLandingTakeoffAreaCollectionPropertyType, _ = front.(*LandingTakeoffAreaCollectionPropertyType)
		}
		stage.OnAfterLandingTakeoffAreaCollectionPropertyTypeUpdateCallback.OnAfterUpdate(stage, landingtakeoffareacollectionpropertytype, frontLandingTakeoffAreaCollectionPropertyType)
	}
}

func (landingtakeoffareacollectionpropertytype *LandingTakeoffAreaCollectionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLandingTakeoffAreaCollectionPropertyTypeDeleteCallback != nil {
		var frontLandingTakeoffAreaCollectionPropertyType *LandingTakeoffAreaCollectionPropertyType
		if front != nil {
			frontLandingTakeoffAreaCollectionPropertyType, _ = front.(*LandingTakeoffAreaCollectionPropertyType)
		}
		stage.OnAfterLandingTakeoffAreaCollectionPropertyTypeDeleteCallback.OnAfterDelete(stage, landingtakeoffareacollectionpropertytype, frontLandingTakeoffAreaCollectionPropertyType)
	}
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLandingTakeoffAreaCollectionTypeCreateCallback != nil {
		stage.OnAfterLandingTakeoffAreaCollectionTypeCreateCallback.OnAfterCreate(stage, landingtakeoffareacollectiontype)
	}
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLandingTakeoffAreaCollectionTypeUpdateCallback != nil {
		var frontLandingTakeoffAreaCollectionType *LandingTakeoffAreaCollectionType
		if front != nil {
			frontLandingTakeoffAreaCollectionType, _ = front.(*LandingTakeoffAreaCollectionType)
		}
		stage.OnAfterLandingTakeoffAreaCollectionTypeUpdateCallback.OnAfterUpdate(stage, landingtakeoffareacollectiontype, frontLandingTakeoffAreaCollectionType)
	}
}

func (landingtakeoffareacollectiontype *LandingTakeoffAreaCollectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLandingTakeoffAreaCollectionTypeDeleteCallback != nil {
		var frontLandingTakeoffAreaCollectionType *LandingTakeoffAreaCollectionType
		if front != nil {
			frontLandingTakeoffAreaCollectionType, _ = front.(*LandingTakeoffAreaCollectionType)
		}
		stage.OnAfterLandingTakeoffAreaCollectionTypeDeleteCallback.OnAfterDelete(stage, landingtakeoffareacollectiontype, frontLandingTakeoffAreaCollectionType)
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

func (minimapropertytype *MinimaPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMinimaPropertyTypeCreateCallback != nil {
		stage.OnAfterMinimaPropertyTypeCreateCallback.OnAfterCreate(stage, minimapropertytype)
	}
}

func (minimapropertytype *MinimaPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMinimaPropertyTypeUpdateCallback != nil {
		var frontMinimaPropertyType *MinimaPropertyType
		if front != nil {
			frontMinimaPropertyType, _ = front.(*MinimaPropertyType)
		}
		stage.OnAfterMinimaPropertyTypeUpdateCallback.OnAfterUpdate(stage, minimapropertytype, frontMinimaPropertyType)
	}
}

func (minimapropertytype *MinimaPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMinimaPropertyTypeDeleteCallback != nil {
		var frontMinimaPropertyType *MinimaPropertyType
		if front != nil {
			frontMinimaPropertyType, _ = front.(*MinimaPropertyType)
		}
		stage.OnAfterMinimaPropertyTypeDeleteCallback.OnAfterDelete(stage, minimapropertytype, frontMinimaPropertyType)
	}
}

func (minimatype *MinimaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMinimaTypeCreateCallback != nil {
		stage.OnAfterMinimaTypeCreateCallback.OnAfterCreate(stage, minimatype)
	}
}

func (minimatype *MinimaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMinimaTypeUpdateCallback != nil {
		var frontMinimaType *MinimaType
		if front != nil {
			frontMinimaType, _ = front.(*MinimaType)
		}
		stage.OnAfterMinimaTypeUpdateCallback.OnAfterUpdate(stage, minimatype, frontMinimaType)
	}
}

func (minimatype *MinimaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMinimaTypeDeleteCallback != nil {
		var frontMinimaType *MinimaType
		if front != nil {
			frontMinimaType, _ = front.(*MinimaType)
		}
		stage.OnAfterMinimaTypeDeleteCallback.OnAfterDelete(stage, minimatype, frontMinimaType)
	}
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMissedApproachGroupPropertyTypeCreateCallback != nil {
		stage.OnAfterMissedApproachGroupPropertyTypeCreateCallback.OnAfterCreate(stage, missedapproachgrouppropertytype)
	}
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachGroupPropertyTypeUpdateCallback != nil {
		var frontMissedApproachGroupPropertyType *MissedApproachGroupPropertyType
		if front != nil {
			frontMissedApproachGroupPropertyType, _ = front.(*MissedApproachGroupPropertyType)
		}
		stage.OnAfterMissedApproachGroupPropertyTypeUpdateCallback.OnAfterUpdate(stage, missedapproachgrouppropertytype, frontMissedApproachGroupPropertyType)
	}
}

func (missedapproachgrouppropertytype *MissedApproachGroupPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachGroupPropertyTypeDeleteCallback != nil {
		var frontMissedApproachGroupPropertyType *MissedApproachGroupPropertyType
		if front != nil {
			frontMissedApproachGroupPropertyType, _ = front.(*MissedApproachGroupPropertyType)
		}
		stage.OnAfterMissedApproachGroupPropertyTypeDeleteCallback.OnAfterDelete(stage, missedapproachgrouppropertytype, frontMissedApproachGroupPropertyType)
	}
}

func (missedapproachgrouptype *MissedApproachGroupType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMissedApproachGroupTypeCreateCallback != nil {
		stage.OnAfterMissedApproachGroupTypeCreateCallback.OnAfterCreate(stage, missedapproachgrouptype)
	}
}

func (missedapproachgrouptype *MissedApproachGroupType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachGroupTypeUpdateCallback != nil {
		var frontMissedApproachGroupType *MissedApproachGroupType
		if front != nil {
			frontMissedApproachGroupType, _ = front.(*MissedApproachGroupType)
		}
		stage.OnAfterMissedApproachGroupTypeUpdateCallback.OnAfterUpdate(stage, missedapproachgrouptype, frontMissedApproachGroupType)
	}
}

func (missedapproachgrouptype *MissedApproachGroupType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachGroupTypeDeleteCallback != nil {
		var frontMissedApproachGroupType *MissedApproachGroupType
		if front != nil {
			frontMissedApproachGroupType, _ = front.(*MissedApproachGroupType)
		}
		stage.OnAfterMissedApproachGroupTypeDeleteCallback.OnAfterDelete(stage, missedapproachgrouptype, frontMissedApproachGroupType)
	}
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMissedApproachLegTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterMissedApproachLegTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, missedapproachlegtimeslicepropertytype)
	}
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachLegTimeSlicePropertyTypeUpdateCallback != nil {
		var frontMissedApproachLegTimeSlicePropertyType *MissedApproachLegTimeSlicePropertyType
		if front != nil {
			frontMissedApproachLegTimeSlicePropertyType, _ = front.(*MissedApproachLegTimeSlicePropertyType)
		}
		stage.OnAfterMissedApproachLegTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, missedapproachlegtimeslicepropertytype, frontMissedApproachLegTimeSlicePropertyType)
	}
}

func (missedapproachlegtimeslicepropertytype *MissedApproachLegTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachLegTimeSlicePropertyTypeDeleteCallback != nil {
		var frontMissedApproachLegTimeSlicePropertyType *MissedApproachLegTimeSlicePropertyType
		if front != nil {
			frontMissedApproachLegTimeSlicePropertyType, _ = front.(*MissedApproachLegTimeSlicePropertyType)
		}
		stage.OnAfterMissedApproachLegTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, missedapproachlegtimeslicepropertytype, frontMissedApproachLegTimeSlicePropertyType)
	}
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMissedApproachLegTimeSliceTypeCreateCallback != nil {
		stage.OnAfterMissedApproachLegTimeSliceTypeCreateCallback.OnAfterCreate(stage, missedapproachlegtimeslicetype)
	}
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachLegTimeSliceTypeUpdateCallback != nil {
		var frontMissedApproachLegTimeSliceType *MissedApproachLegTimeSliceType
		if front != nil {
			frontMissedApproachLegTimeSliceType, _ = front.(*MissedApproachLegTimeSliceType)
		}
		stage.OnAfterMissedApproachLegTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, missedapproachlegtimeslicetype, frontMissedApproachLegTimeSliceType)
	}
}

func (missedapproachlegtimeslicetype *MissedApproachLegTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachLegTimeSliceTypeDeleteCallback != nil {
		var frontMissedApproachLegTimeSliceType *MissedApproachLegTimeSliceType
		if front != nil {
			frontMissedApproachLegTimeSliceType, _ = front.(*MissedApproachLegTimeSliceType)
		}
		stage.OnAfterMissedApproachLegTimeSliceTypeDeleteCallback.OnAfterDelete(stage, missedapproachlegtimeslicetype, frontMissedApproachLegTimeSliceType)
	}
}

func (missedapproachlegtype *MissedApproachLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMissedApproachLegTypeCreateCallback != nil {
		stage.OnAfterMissedApproachLegTypeCreateCallback.OnAfterCreate(stage, missedapproachlegtype)
	}
}

func (missedapproachlegtype *MissedApproachLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachLegTypeUpdateCallback != nil {
		var frontMissedApproachLegType *MissedApproachLegType
		if front != nil {
			frontMissedApproachLegType, _ = front.(*MissedApproachLegType)
		}
		stage.OnAfterMissedApproachLegTypeUpdateCallback.OnAfterUpdate(stage, missedapproachlegtype, frontMissedApproachLegType)
	}
}

func (missedapproachlegtype *MissedApproachLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMissedApproachLegTypeDeleteCallback != nil {
		var frontMissedApproachLegType *MissedApproachLegType
		if front != nil {
			frontMissedApproachLegType, _ = front.(*MissedApproachLegType)
		}
		stage.OnAfterMissedApproachLegTypeDeleteCallback.OnAfterDelete(stage, missedapproachlegtype, frontMissedApproachLegType)
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

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavigationAreaRestrictionTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterNavigationAreaRestrictionTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, navigationarearestrictiontimeslicepropertytype)
	}
}

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaRestrictionTimeSlicePropertyTypeUpdateCallback != nil {
		var frontNavigationAreaRestrictionTimeSlicePropertyType *NavigationAreaRestrictionTimeSlicePropertyType
		if front != nil {
			frontNavigationAreaRestrictionTimeSlicePropertyType, _ = front.(*NavigationAreaRestrictionTimeSlicePropertyType)
		}
		stage.OnAfterNavigationAreaRestrictionTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, navigationarearestrictiontimeslicepropertytype, frontNavigationAreaRestrictionTimeSlicePropertyType)
	}
}

func (navigationarearestrictiontimeslicepropertytype *NavigationAreaRestrictionTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaRestrictionTimeSlicePropertyTypeDeleteCallback != nil {
		var frontNavigationAreaRestrictionTimeSlicePropertyType *NavigationAreaRestrictionTimeSlicePropertyType
		if front != nil {
			frontNavigationAreaRestrictionTimeSlicePropertyType, _ = front.(*NavigationAreaRestrictionTimeSlicePropertyType)
		}
		stage.OnAfterNavigationAreaRestrictionTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, navigationarearestrictiontimeslicepropertytype, frontNavigationAreaRestrictionTimeSlicePropertyType)
	}
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavigationAreaRestrictionTimeSliceTypeCreateCallback != nil {
		stage.OnAfterNavigationAreaRestrictionTimeSliceTypeCreateCallback.OnAfterCreate(stage, navigationarearestrictiontimeslicetype)
	}
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaRestrictionTimeSliceTypeUpdateCallback != nil {
		var frontNavigationAreaRestrictionTimeSliceType *NavigationAreaRestrictionTimeSliceType
		if front != nil {
			frontNavigationAreaRestrictionTimeSliceType, _ = front.(*NavigationAreaRestrictionTimeSliceType)
		}
		stage.OnAfterNavigationAreaRestrictionTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, navigationarearestrictiontimeslicetype, frontNavigationAreaRestrictionTimeSliceType)
	}
}

func (navigationarearestrictiontimeslicetype *NavigationAreaRestrictionTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaRestrictionTimeSliceTypeDeleteCallback != nil {
		var frontNavigationAreaRestrictionTimeSliceType *NavigationAreaRestrictionTimeSliceType
		if front != nil {
			frontNavigationAreaRestrictionTimeSliceType, _ = front.(*NavigationAreaRestrictionTimeSliceType)
		}
		stage.OnAfterNavigationAreaRestrictionTimeSliceTypeDeleteCallback.OnAfterDelete(stage, navigationarearestrictiontimeslicetype, frontNavigationAreaRestrictionTimeSliceType)
	}
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavigationAreaRestrictionTypeCreateCallback != nil {
		stage.OnAfterNavigationAreaRestrictionTypeCreateCallback.OnAfterCreate(stage, navigationarearestrictiontype)
	}
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaRestrictionTypeUpdateCallback != nil {
		var frontNavigationAreaRestrictionType *NavigationAreaRestrictionType
		if front != nil {
			frontNavigationAreaRestrictionType, _ = front.(*NavigationAreaRestrictionType)
		}
		stage.OnAfterNavigationAreaRestrictionTypeUpdateCallback.OnAfterUpdate(stage, navigationarearestrictiontype, frontNavigationAreaRestrictionType)
	}
}

func (navigationarearestrictiontype *NavigationAreaRestrictionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaRestrictionTypeDeleteCallback != nil {
		var frontNavigationAreaRestrictionType *NavigationAreaRestrictionType
		if front != nil {
			frontNavigationAreaRestrictionType, _ = front.(*NavigationAreaRestrictionType)
		}
		stage.OnAfterNavigationAreaRestrictionTypeDeleteCallback.OnAfterDelete(stage, navigationarearestrictiontype, frontNavigationAreaRestrictionType)
	}
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavigationAreaSectorPropertyTypeCreateCallback != nil {
		stage.OnAfterNavigationAreaSectorPropertyTypeCreateCallback.OnAfterCreate(stage, navigationareasectorpropertytype)
	}
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaSectorPropertyTypeUpdateCallback != nil {
		var frontNavigationAreaSectorPropertyType *NavigationAreaSectorPropertyType
		if front != nil {
			frontNavigationAreaSectorPropertyType, _ = front.(*NavigationAreaSectorPropertyType)
		}
		stage.OnAfterNavigationAreaSectorPropertyTypeUpdateCallback.OnAfterUpdate(stage, navigationareasectorpropertytype, frontNavigationAreaSectorPropertyType)
	}
}

func (navigationareasectorpropertytype *NavigationAreaSectorPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaSectorPropertyTypeDeleteCallback != nil {
		var frontNavigationAreaSectorPropertyType *NavigationAreaSectorPropertyType
		if front != nil {
			frontNavigationAreaSectorPropertyType, _ = front.(*NavigationAreaSectorPropertyType)
		}
		stage.OnAfterNavigationAreaSectorPropertyTypeDeleteCallback.OnAfterDelete(stage, navigationareasectorpropertytype, frontNavigationAreaSectorPropertyType)
	}
}

func (navigationareasectortype *NavigationAreaSectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavigationAreaSectorTypeCreateCallback != nil {
		stage.OnAfterNavigationAreaSectorTypeCreateCallback.OnAfterCreate(stage, navigationareasectortype)
	}
}

func (navigationareasectortype *NavigationAreaSectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaSectorTypeUpdateCallback != nil {
		var frontNavigationAreaSectorType *NavigationAreaSectorType
		if front != nil {
			frontNavigationAreaSectorType, _ = front.(*NavigationAreaSectorType)
		}
		stage.OnAfterNavigationAreaSectorTypeUpdateCallback.OnAfterUpdate(stage, navigationareasectortype, frontNavigationAreaSectorType)
	}
}

func (navigationareasectortype *NavigationAreaSectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaSectorTypeDeleteCallback != nil {
		var frontNavigationAreaSectorType *NavigationAreaSectorType
		if front != nil {
			frontNavigationAreaSectorType, _ = front.(*NavigationAreaSectorType)
		}
		stage.OnAfterNavigationAreaSectorTypeDeleteCallback.OnAfterDelete(stage, navigationareasectortype, frontNavigationAreaSectorType)
	}
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavigationAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterNavigationAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, navigationareatimeslicepropertytype)
	}
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontNavigationAreaTimeSlicePropertyType *NavigationAreaTimeSlicePropertyType
		if front != nil {
			frontNavigationAreaTimeSlicePropertyType, _ = front.(*NavigationAreaTimeSlicePropertyType)
		}
		stage.OnAfterNavigationAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, navigationareatimeslicepropertytype, frontNavigationAreaTimeSlicePropertyType)
	}
}

func (navigationareatimeslicepropertytype *NavigationAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontNavigationAreaTimeSlicePropertyType *NavigationAreaTimeSlicePropertyType
		if front != nil {
			frontNavigationAreaTimeSlicePropertyType, _ = front.(*NavigationAreaTimeSlicePropertyType)
		}
		stage.OnAfterNavigationAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, navigationareatimeslicepropertytype, frontNavigationAreaTimeSlicePropertyType)
	}
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavigationAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterNavigationAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, navigationareatimeslicetype)
	}
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaTimeSliceTypeUpdateCallback != nil {
		var frontNavigationAreaTimeSliceType *NavigationAreaTimeSliceType
		if front != nil {
			frontNavigationAreaTimeSliceType, _ = front.(*NavigationAreaTimeSliceType)
		}
		stage.OnAfterNavigationAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, navigationareatimeslicetype, frontNavigationAreaTimeSliceType)
	}
}

func (navigationareatimeslicetype *NavigationAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaTimeSliceTypeDeleteCallback != nil {
		var frontNavigationAreaTimeSliceType *NavigationAreaTimeSliceType
		if front != nil {
			frontNavigationAreaTimeSliceType, _ = front.(*NavigationAreaTimeSliceType)
		}
		stage.OnAfterNavigationAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, navigationareatimeslicetype, frontNavigationAreaTimeSliceType)
	}
}

func (navigationareatype *NavigationAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavigationAreaTypeCreateCallback != nil {
		stage.OnAfterNavigationAreaTypeCreateCallback.OnAfterCreate(stage, navigationareatype)
	}
}

func (navigationareatype *NavigationAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaTypeUpdateCallback != nil {
		var frontNavigationAreaType *NavigationAreaType
		if front != nil {
			frontNavigationAreaType, _ = front.(*NavigationAreaType)
		}
		stage.OnAfterNavigationAreaTypeUpdateCallback.OnAfterUpdate(stage, navigationareatype, frontNavigationAreaType)
	}
}

func (navigationareatype *NavigationAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavigationAreaTypeDeleteCallback != nil {
		var frontNavigationAreaType *NavigationAreaType
		if front != nil {
			frontNavigationAreaType, _ = front.(*NavigationAreaType)
		}
		stage.OnAfterNavigationAreaTypeDeleteCallback.OnAfterDelete(stage, navigationareatype, frontNavigationAreaType)
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

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObstacleAssessmentAreaPropertyTypeCreateCallback != nil {
		stage.OnAfterObstacleAssessmentAreaPropertyTypeCreateCallback.OnAfterCreate(stage, obstacleassessmentareapropertytype)
	}
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAssessmentAreaPropertyTypeUpdateCallback != nil {
		var frontObstacleAssessmentAreaPropertyType *ObstacleAssessmentAreaPropertyType
		if front != nil {
			frontObstacleAssessmentAreaPropertyType, _ = front.(*ObstacleAssessmentAreaPropertyType)
		}
		stage.OnAfterObstacleAssessmentAreaPropertyTypeUpdateCallback.OnAfterUpdate(stage, obstacleassessmentareapropertytype, frontObstacleAssessmentAreaPropertyType)
	}
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAssessmentAreaPropertyTypeDeleteCallback != nil {
		var frontObstacleAssessmentAreaPropertyType *ObstacleAssessmentAreaPropertyType
		if front != nil {
			frontObstacleAssessmentAreaPropertyType, _ = front.(*ObstacleAssessmentAreaPropertyType)
		}
		stage.OnAfterObstacleAssessmentAreaPropertyTypeDeleteCallback.OnAfterDelete(stage, obstacleassessmentareapropertytype, frontObstacleAssessmentAreaPropertyType)
	}
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObstacleAssessmentAreaTypeCreateCallback != nil {
		stage.OnAfterObstacleAssessmentAreaTypeCreateCallback.OnAfterCreate(stage, obstacleassessmentareatype)
	}
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAssessmentAreaTypeUpdateCallback != nil {
		var frontObstacleAssessmentAreaType *ObstacleAssessmentAreaType
		if front != nil {
			frontObstacleAssessmentAreaType, _ = front.(*ObstacleAssessmentAreaType)
		}
		stage.OnAfterObstacleAssessmentAreaTypeUpdateCallback.OnAfterUpdate(stage, obstacleassessmentareatype, frontObstacleAssessmentAreaType)
	}
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAssessmentAreaTypeDeleteCallback != nil {
		var frontObstacleAssessmentAreaType *ObstacleAssessmentAreaType
		if front != nil {
			frontObstacleAssessmentAreaType, _ = front.(*ObstacleAssessmentAreaType)
		}
		stage.OnAfterObstacleAssessmentAreaTypeDeleteCallback.OnAfterDelete(stage, obstacleassessmentareatype, frontObstacleAssessmentAreaType)
	}
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObstaclePlacementPropertyTypeCreateCallback != nil {
		stage.OnAfterObstaclePlacementPropertyTypeCreateCallback.OnAfterCreate(stage, obstacleplacementpropertytype)
	}
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstaclePlacementPropertyTypeUpdateCallback != nil {
		var frontObstaclePlacementPropertyType *ObstaclePlacementPropertyType
		if front != nil {
			frontObstaclePlacementPropertyType, _ = front.(*ObstaclePlacementPropertyType)
		}
		stage.OnAfterObstaclePlacementPropertyTypeUpdateCallback.OnAfterUpdate(stage, obstacleplacementpropertytype, frontObstaclePlacementPropertyType)
	}
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstaclePlacementPropertyTypeDeleteCallback != nil {
		var frontObstaclePlacementPropertyType *ObstaclePlacementPropertyType
		if front != nil {
			frontObstaclePlacementPropertyType, _ = front.(*ObstaclePlacementPropertyType)
		}
		stage.OnAfterObstaclePlacementPropertyTypeDeleteCallback.OnAfterDelete(stage, obstacleplacementpropertytype, frontObstaclePlacementPropertyType)
	}
}

func (obstacleplacementtype *ObstaclePlacementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObstaclePlacementTypeCreateCallback != nil {
		stage.OnAfterObstaclePlacementTypeCreateCallback.OnAfterCreate(stage, obstacleplacementtype)
	}
}

func (obstacleplacementtype *ObstaclePlacementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstaclePlacementTypeUpdateCallback != nil {
		var frontObstaclePlacementType *ObstaclePlacementType
		if front != nil {
			frontObstaclePlacementType, _ = front.(*ObstaclePlacementType)
		}
		stage.OnAfterObstaclePlacementTypeUpdateCallback.OnAfterUpdate(stage, obstacleplacementtype, frontObstaclePlacementType)
	}
}

func (obstacleplacementtype *ObstaclePlacementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstaclePlacementTypeDeleteCallback != nil {
		var frontObstaclePlacementType *ObstaclePlacementType
		if front != nil {
			frontObstaclePlacementType, _ = front.(*ObstaclePlacementType)
		}
		stage.OnAfterObstaclePlacementTypeDeleteCallback.OnAfterDelete(stage, obstacleplacementtype, frontObstaclePlacementType)
	}
}

func (obstructionpropertytype *ObstructionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObstructionPropertyTypeCreateCallback != nil {
		stage.OnAfterObstructionPropertyTypeCreateCallback.OnAfterCreate(stage, obstructionpropertytype)
	}
}

func (obstructionpropertytype *ObstructionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstructionPropertyTypeUpdateCallback != nil {
		var frontObstructionPropertyType *ObstructionPropertyType
		if front != nil {
			frontObstructionPropertyType, _ = front.(*ObstructionPropertyType)
		}
		stage.OnAfterObstructionPropertyTypeUpdateCallback.OnAfterUpdate(stage, obstructionpropertytype, frontObstructionPropertyType)
	}
}

func (obstructionpropertytype *ObstructionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstructionPropertyTypeDeleteCallback != nil {
		var frontObstructionPropertyType *ObstructionPropertyType
		if front != nil {
			frontObstructionPropertyType, _ = front.(*ObstructionPropertyType)
		}
		stage.OnAfterObstructionPropertyTypeDeleteCallback.OnAfterDelete(stage, obstructionpropertytype, frontObstructionPropertyType)
	}
}

func (obstructiontype *ObstructionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObstructionTypeCreateCallback != nil {
		stage.OnAfterObstructionTypeCreateCallback.OnAfterCreate(stage, obstructiontype)
	}
}

func (obstructiontype *ObstructionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstructionTypeUpdateCallback != nil {
		var frontObstructionType *ObstructionType
		if front != nil {
			frontObstructionType, _ = front.(*ObstructionType)
		}
		stage.OnAfterObstructionTypeUpdateCallback.OnAfterUpdate(stage, obstructiontype, frontObstructionType)
	}
}

func (obstructiontype *ObstructionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstructionTypeDeleteCallback != nil {
		var frontObstructionType *ObstructionType
		if front != nil {
			frontObstructionType, _ = front.(*ObstructionType)
		}
		stage.OnAfterObstructionTypeDeleteCallback.OnAfterDelete(stage, obstructiontype, frontObstructionType)
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

func (pointreferencepropertytype *PointReferencePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointReferencePropertyTypeCreateCallback != nil {
		stage.OnAfterPointReferencePropertyTypeCreateCallback.OnAfterCreate(stage, pointreferencepropertytype)
	}
}

func (pointreferencepropertytype *PointReferencePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointReferencePropertyTypeUpdateCallback != nil {
		var frontPointReferencePropertyType *PointReferencePropertyType
		if front != nil {
			frontPointReferencePropertyType, _ = front.(*PointReferencePropertyType)
		}
		stage.OnAfterPointReferencePropertyTypeUpdateCallback.OnAfterUpdate(stage, pointreferencepropertytype, frontPointReferencePropertyType)
	}
}

func (pointreferencepropertytype *PointReferencePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointReferencePropertyTypeDeleteCallback != nil {
		var frontPointReferencePropertyType *PointReferencePropertyType
		if front != nil {
			frontPointReferencePropertyType, _ = front.(*PointReferencePropertyType)
		}
		stage.OnAfterPointReferencePropertyTypeDeleteCallback.OnAfterDelete(stage, pointreferencepropertytype, frontPointReferencePropertyType)
	}
}

func (pointreferencetype *PointReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointReferenceTypeCreateCallback != nil {
		stage.OnAfterPointReferenceTypeCreateCallback.OnAfterCreate(stage, pointreferencetype)
	}
}

func (pointreferencetype *PointReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointReferenceTypeUpdateCallback != nil {
		var frontPointReferenceType *PointReferenceType
		if front != nil {
			frontPointReferenceType, _ = front.(*PointReferenceType)
		}
		stage.OnAfterPointReferenceTypeUpdateCallback.OnAfterUpdate(stage, pointreferencetype, frontPointReferenceType)
	}
}

func (pointreferencetype *PointReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointReferenceTypeDeleteCallback != nil {
		var frontPointReferenceType *PointReferenceType
		if front != nil {
			frontPointReferenceType, _ = front.(*PointReferenceType)
		}
		stage.OnAfterPointReferenceTypeDeleteCallback.OnAfterDelete(stage, pointreferencetype, frontPointReferenceType)
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

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedureAvailabilityPropertyTypeCreateCallback != nil {
		stage.OnAfterProcedureAvailabilityPropertyTypeCreateCallback.OnAfterCreate(stage, procedureavailabilitypropertytype)
	}
}

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureAvailabilityPropertyTypeUpdateCallback != nil {
		var frontProcedureAvailabilityPropertyType *ProcedureAvailabilityPropertyType
		if front != nil {
			frontProcedureAvailabilityPropertyType, _ = front.(*ProcedureAvailabilityPropertyType)
		}
		stage.OnAfterProcedureAvailabilityPropertyTypeUpdateCallback.OnAfterUpdate(stage, procedureavailabilitypropertytype, frontProcedureAvailabilityPropertyType)
	}
}

func (procedureavailabilitypropertytype *ProcedureAvailabilityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureAvailabilityPropertyTypeDeleteCallback != nil {
		var frontProcedureAvailabilityPropertyType *ProcedureAvailabilityPropertyType
		if front != nil {
			frontProcedureAvailabilityPropertyType, _ = front.(*ProcedureAvailabilityPropertyType)
		}
		stage.OnAfterProcedureAvailabilityPropertyTypeDeleteCallback.OnAfterDelete(stage, procedureavailabilitypropertytype, frontProcedureAvailabilityPropertyType)
	}
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedureAvailabilityTypeCreateCallback != nil {
		stage.OnAfterProcedureAvailabilityTypeCreateCallback.OnAfterCreate(stage, procedureavailabilitytype)
	}
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureAvailabilityTypeUpdateCallback != nil {
		var frontProcedureAvailabilityType *ProcedureAvailabilityType
		if front != nil {
			frontProcedureAvailabilityType, _ = front.(*ProcedureAvailabilityType)
		}
		stage.OnAfterProcedureAvailabilityTypeUpdateCallback.OnAfterUpdate(stage, procedureavailabilitytype, frontProcedureAvailabilityType)
	}
}

func (procedureavailabilitytype *ProcedureAvailabilityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureAvailabilityTypeDeleteCallback != nil {
		var frontProcedureAvailabilityType *ProcedureAvailabilityType
		if front != nil {
			frontProcedureAvailabilityType, _ = front.(*ProcedureAvailabilityType)
		}
		stage.OnAfterProcedureAvailabilityTypeDeleteCallback.OnAfterDelete(stage, procedureavailabilitytype, frontProcedureAvailabilityType)
	}
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedureDMETimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterProcedureDMETimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, proceduredmetimeslicepropertytype)
	}
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureDMETimeSlicePropertyTypeUpdateCallback != nil {
		var frontProcedureDMETimeSlicePropertyType *ProcedureDMETimeSlicePropertyType
		if front != nil {
			frontProcedureDMETimeSlicePropertyType, _ = front.(*ProcedureDMETimeSlicePropertyType)
		}
		stage.OnAfterProcedureDMETimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, proceduredmetimeslicepropertytype, frontProcedureDMETimeSlicePropertyType)
	}
}

func (proceduredmetimeslicepropertytype *ProcedureDMETimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureDMETimeSlicePropertyTypeDeleteCallback != nil {
		var frontProcedureDMETimeSlicePropertyType *ProcedureDMETimeSlicePropertyType
		if front != nil {
			frontProcedureDMETimeSlicePropertyType, _ = front.(*ProcedureDMETimeSlicePropertyType)
		}
		stage.OnAfterProcedureDMETimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, proceduredmetimeslicepropertytype, frontProcedureDMETimeSlicePropertyType)
	}
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedureDMETimeSliceTypeCreateCallback != nil {
		stage.OnAfterProcedureDMETimeSliceTypeCreateCallback.OnAfterCreate(stage, proceduredmetimeslicetype)
	}
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureDMETimeSliceTypeUpdateCallback != nil {
		var frontProcedureDMETimeSliceType *ProcedureDMETimeSliceType
		if front != nil {
			frontProcedureDMETimeSliceType, _ = front.(*ProcedureDMETimeSliceType)
		}
		stage.OnAfterProcedureDMETimeSliceTypeUpdateCallback.OnAfterUpdate(stage, proceduredmetimeslicetype, frontProcedureDMETimeSliceType)
	}
}

func (proceduredmetimeslicetype *ProcedureDMETimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureDMETimeSliceTypeDeleteCallback != nil {
		var frontProcedureDMETimeSliceType *ProcedureDMETimeSliceType
		if front != nil {
			frontProcedureDMETimeSliceType, _ = front.(*ProcedureDMETimeSliceType)
		}
		stage.OnAfterProcedureDMETimeSliceTypeDeleteCallback.OnAfterDelete(stage, proceduredmetimeslicetype, frontProcedureDMETimeSliceType)
	}
}

func (proceduredmetype *ProcedureDMEType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedureDMETypeCreateCallback != nil {
		stage.OnAfterProcedureDMETypeCreateCallback.OnAfterCreate(stage, proceduredmetype)
	}
}

func (proceduredmetype *ProcedureDMEType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureDMETypeUpdateCallback != nil {
		var frontProcedureDMEType *ProcedureDMEType
		if front != nil {
			frontProcedureDMEType, _ = front.(*ProcedureDMEType)
		}
		stage.OnAfterProcedureDMETypeUpdateCallback.OnAfterUpdate(stage, proceduredmetype, frontProcedureDMEType)
	}
}

func (proceduredmetype *ProcedureDMEType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureDMETypeDeleteCallback != nil {
		var frontProcedureDMEType *ProcedureDMEType
		if front != nil {
			frontProcedureDMEType, _ = front.(*ProcedureDMEType)
		}
		stage.OnAfterProcedureDMETypeDeleteCallback.OnAfterDelete(stage, proceduredmetype, frontProcedureDMEType)
	}
}

func (procedurepropertytype *ProcedurePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedurePropertyTypeCreateCallback != nil {
		stage.OnAfterProcedurePropertyTypeCreateCallback.OnAfterCreate(stage, procedurepropertytype)
	}
}

func (procedurepropertytype *ProcedurePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedurePropertyTypeUpdateCallback != nil {
		var frontProcedurePropertyType *ProcedurePropertyType
		if front != nil {
			frontProcedurePropertyType, _ = front.(*ProcedurePropertyType)
		}
		stage.OnAfterProcedurePropertyTypeUpdateCallback.OnAfterUpdate(stage, procedurepropertytype, frontProcedurePropertyType)
	}
}

func (procedurepropertytype *ProcedurePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedurePropertyTypeDeleteCallback != nil {
		var frontProcedurePropertyType *ProcedurePropertyType
		if front != nil {
			frontProcedurePropertyType, _ = front.(*ProcedurePropertyType)
		}
		stage.OnAfterProcedurePropertyTypeDeleteCallback.OnAfterDelete(stage, procedurepropertytype, frontProcedurePropertyType)
	}
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedureTransitionLegPropertyTypeCreateCallback != nil {
		stage.OnAfterProcedureTransitionLegPropertyTypeCreateCallback.OnAfterCreate(stage, proceduretransitionlegpropertytype)
	}
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureTransitionLegPropertyTypeUpdateCallback != nil {
		var frontProcedureTransitionLegPropertyType *ProcedureTransitionLegPropertyType
		if front != nil {
			frontProcedureTransitionLegPropertyType, _ = front.(*ProcedureTransitionLegPropertyType)
		}
		stage.OnAfterProcedureTransitionLegPropertyTypeUpdateCallback.OnAfterUpdate(stage, proceduretransitionlegpropertytype, frontProcedureTransitionLegPropertyType)
	}
}

func (proceduretransitionlegpropertytype *ProcedureTransitionLegPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureTransitionLegPropertyTypeDeleteCallback != nil {
		var frontProcedureTransitionLegPropertyType *ProcedureTransitionLegPropertyType
		if front != nil {
			frontProcedureTransitionLegPropertyType, _ = front.(*ProcedureTransitionLegPropertyType)
		}
		stage.OnAfterProcedureTransitionLegPropertyTypeDeleteCallback.OnAfterDelete(stage, proceduretransitionlegpropertytype, frontProcedureTransitionLegPropertyType)
	}
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedureTransitionLegTypeCreateCallback != nil {
		stage.OnAfterProcedureTransitionLegTypeCreateCallback.OnAfterCreate(stage, proceduretransitionlegtype)
	}
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureTransitionLegTypeUpdateCallback != nil {
		var frontProcedureTransitionLegType *ProcedureTransitionLegType
		if front != nil {
			frontProcedureTransitionLegType, _ = front.(*ProcedureTransitionLegType)
		}
		stage.OnAfterProcedureTransitionLegTypeUpdateCallback.OnAfterUpdate(stage, proceduretransitionlegtype, frontProcedureTransitionLegType)
	}
}

func (proceduretransitionlegtype *ProcedureTransitionLegType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureTransitionLegTypeDeleteCallback != nil {
		var frontProcedureTransitionLegType *ProcedureTransitionLegType
		if front != nil {
			frontProcedureTransitionLegType, _ = front.(*ProcedureTransitionLegType)
		}
		stage.OnAfterProcedureTransitionLegTypeDeleteCallback.OnAfterDelete(stage, proceduretransitionlegtype, frontProcedureTransitionLegType)
	}
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedureTransitionPropertyTypeCreateCallback != nil {
		stage.OnAfterProcedureTransitionPropertyTypeCreateCallback.OnAfterCreate(stage, proceduretransitionpropertytype)
	}
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureTransitionPropertyTypeUpdateCallback != nil {
		var frontProcedureTransitionPropertyType *ProcedureTransitionPropertyType
		if front != nil {
			frontProcedureTransitionPropertyType, _ = front.(*ProcedureTransitionPropertyType)
		}
		stage.OnAfterProcedureTransitionPropertyTypeUpdateCallback.OnAfterUpdate(stage, proceduretransitionpropertytype, frontProcedureTransitionPropertyType)
	}
}

func (proceduretransitionpropertytype *ProcedureTransitionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureTransitionPropertyTypeDeleteCallback != nil {
		var frontProcedureTransitionPropertyType *ProcedureTransitionPropertyType
		if front != nil {
			frontProcedureTransitionPropertyType, _ = front.(*ProcedureTransitionPropertyType)
		}
		stage.OnAfterProcedureTransitionPropertyTypeDeleteCallback.OnAfterDelete(stage, proceduretransitionpropertytype, frontProcedureTransitionPropertyType)
	}
}

func (proceduretransitiontype *ProcedureTransitionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedureTransitionTypeCreateCallback != nil {
		stage.OnAfterProcedureTransitionTypeCreateCallback.OnAfterCreate(stage, proceduretransitiontype)
	}
}

func (proceduretransitiontype *ProcedureTransitionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureTransitionTypeUpdateCallback != nil {
		var frontProcedureTransitionType *ProcedureTransitionType
		if front != nil {
			frontProcedureTransitionType, _ = front.(*ProcedureTransitionType)
		}
		stage.OnAfterProcedureTransitionTypeUpdateCallback.OnAfterUpdate(stage, proceduretransitiontype, frontProcedureTransitionType)
	}
}

func (proceduretransitiontype *ProcedureTransitionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedureTransitionTypeDeleteCallback != nil {
		var frontProcedureTransitionType *ProcedureTransitionType
		if front != nil {
			frontProcedureTransitionType, _ = front.(*ProcedureTransitionType)
		}
		stage.OnAfterProcedureTransitionTypeDeleteCallback.OnAfterDelete(stage, proceduretransitiontype, frontProcedureTransitionType)
	}
}

func (radarsystempropertytype *RadarSystemPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadarSystemPropertyTypeCreateCallback != nil {
		stage.OnAfterRadarSystemPropertyTypeCreateCallback.OnAfterCreate(stage, radarsystempropertytype)
	}
}

func (radarsystempropertytype *RadarSystemPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarSystemPropertyTypeUpdateCallback != nil {
		var frontRadarSystemPropertyType *RadarSystemPropertyType
		if front != nil {
			frontRadarSystemPropertyType, _ = front.(*RadarSystemPropertyType)
		}
		stage.OnAfterRadarSystemPropertyTypeUpdateCallback.OnAfterUpdate(stage, radarsystempropertytype, frontRadarSystemPropertyType)
	}
}

func (radarsystempropertytype *RadarSystemPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadarSystemPropertyTypeDeleteCallback != nil {
		var frontRadarSystemPropertyType *RadarSystemPropertyType
		if front != nil {
			frontRadarSystemPropertyType, _ = front.(*RadarSystemPropertyType)
		}
		stage.OnAfterRadarSystemPropertyTypeDeleteCallback.OnAfterDelete(stage, radarsystempropertytype, frontRadarSystemPropertyType)
	}
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadioFrequencyAreaPropertyTypeCreateCallback != nil {
		stage.OnAfterRadioFrequencyAreaPropertyTypeCreateCallback.OnAfterCreate(stage, radiofrequencyareapropertytype)
	}
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioFrequencyAreaPropertyTypeUpdateCallback != nil {
		var frontRadioFrequencyAreaPropertyType *RadioFrequencyAreaPropertyType
		if front != nil {
			frontRadioFrequencyAreaPropertyType, _ = front.(*RadioFrequencyAreaPropertyType)
		}
		stage.OnAfterRadioFrequencyAreaPropertyTypeUpdateCallback.OnAfterUpdate(stage, radiofrequencyareapropertytype, frontRadioFrequencyAreaPropertyType)
	}
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioFrequencyAreaPropertyTypeDeleteCallback != nil {
		var frontRadioFrequencyAreaPropertyType *RadioFrequencyAreaPropertyType
		if front != nil {
			frontRadioFrequencyAreaPropertyType, _ = front.(*RadioFrequencyAreaPropertyType)
		}
		stage.OnAfterRadioFrequencyAreaPropertyTypeDeleteCallback.OnAfterDelete(stage, radiofrequencyareapropertytype, frontRadioFrequencyAreaPropertyType)
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

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSafeAltitudeAreaPropertyTypeCreateCallback != nil {
		stage.OnAfterSafeAltitudeAreaPropertyTypeCreateCallback.OnAfterCreate(stage, safealtitudeareapropertytype)
	}
}

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaPropertyTypeUpdateCallback != nil {
		var frontSafeAltitudeAreaPropertyType *SafeAltitudeAreaPropertyType
		if front != nil {
			frontSafeAltitudeAreaPropertyType, _ = front.(*SafeAltitudeAreaPropertyType)
		}
		stage.OnAfterSafeAltitudeAreaPropertyTypeUpdateCallback.OnAfterUpdate(stage, safealtitudeareapropertytype, frontSafeAltitudeAreaPropertyType)
	}
}

func (safealtitudeareapropertytype *SafeAltitudeAreaPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaPropertyTypeDeleteCallback != nil {
		var frontSafeAltitudeAreaPropertyType *SafeAltitudeAreaPropertyType
		if front != nil {
			frontSafeAltitudeAreaPropertyType, _ = front.(*SafeAltitudeAreaPropertyType)
		}
		stage.OnAfterSafeAltitudeAreaPropertyTypeDeleteCallback.OnAfterDelete(stage, safealtitudeareapropertytype, frontSafeAltitudeAreaPropertyType)
	}
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSafeAltitudeAreaSectorPropertyTypeCreateCallback != nil {
		stage.OnAfterSafeAltitudeAreaSectorPropertyTypeCreateCallback.OnAfterCreate(stage, safealtitudeareasectorpropertytype)
	}
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaSectorPropertyTypeUpdateCallback != nil {
		var frontSafeAltitudeAreaSectorPropertyType *SafeAltitudeAreaSectorPropertyType
		if front != nil {
			frontSafeAltitudeAreaSectorPropertyType, _ = front.(*SafeAltitudeAreaSectorPropertyType)
		}
		stage.OnAfterSafeAltitudeAreaSectorPropertyTypeUpdateCallback.OnAfterUpdate(stage, safealtitudeareasectorpropertytype, frontSafeAltitudeAreaSectorPropertyType)
	}
}

func (safealtitudeareasectorpropertytype *SafeAltitudeAreaSectorPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaSectorPropertyTypeDeleteCallback != nil {
		var frontSafeAltitudeAreaSectorPropertyType *SafeAltitudeAreaSectorPropertyType
		if front != nil {
			frontSafeAltitudeAreaSectorPropertyType, _ = front.(*SafeAltitudeAreaSectorPropertyType)
		}
		stage.OnAfterSafeAltitudeAreaSectorPropertyTypeDeleteCallback.OnAfterDelete(stage, safealtitudeareasectorpropertytype, frontSafeAltitudeAreaSectorPropertyType)
	}
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSafeAltitudeAreaSectorTypeCreateCallback != nil {
		stage.OnAfterSafeAltitudeAreaSectorTypeCreateCallback.OnAfterCreate(stage, safealtitudeareasectortype)
	}
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaSectorTypeUpdateCallback != nil {
		var frontSafeAltitudeAreaSectorType *SafeAltitudeAreaSectorType
		if front != nil {
			frontSafeAltitudeAreaSectorType, _ = front.(*SafeAltitudeAreaSectorType)
		}
		stage.OnAfterSafeAltitudeAreaSectorTypeUpdateCallback.OnAfterUpdate(stage, safealtitudeareasectortype, frontSafeAltitudeAreaSectorType)
	}
}

func (safealtitudeareasectortype *SafeAltitudeAreaSectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaSectorTypeDeleteCallback != nil {
		var frontSafeAltitudeAreaSectorType *SafeAltitudeAreaSectorType
		if front != nil {
			frontSafeAltitudeAreaSectorType, _ = front.(*SafeAltitudeAreaSectorType)
		}
		stage.OnAfterSafeAltitudeAreaSectorTypeDeleteCallback.OnAfterDelete(stage, safealtitudeareasectortype, frontSafeAltitudeAreaSectorType)
	}
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSafeAltitudeAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSafeAltitudeAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, safealtitudeareatimeslicepropertytype)
	}
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSafeAltitudeAreaTimeSlicePropertyType *SafeAltitudeAreaTimeSlicePropertyType
		if front != nil {
			frontSafeAltitudeAreaTimeSlicePropertyType, _ = front.(*SafeAltitudeAreaTimeSlicePropertyType)
		}
		stage.OnAfterSafeAltitudeAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, safealtitudeareatimeslicepropertytype, frontSafeAltitudeAreaTimeSlicePropertyType)
	}
}

func (safealtitudeareatimeslicepropertytype *SafeAltitudeAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSafeAltitudeAreaTimeSlicePropertyType *SafeAltitudeAreaTimeSlicePropertyType
		if front != nil {
			frontSafeAltitudeAreaTimeSlicePropertyType, _ = front.(*SafeAltitudeAreaTimeSlicePropertyType)
		}
		stage.OnAfterSafeAltitudeAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, safealtitudeareatimeslicepropertytype, frontSafeAltitudeAreaTimeSlicePropertyType)
	}
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSafeAltitudeAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSafeAltitudeAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, safealtitudeareatimeslicetype)
	}
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaTimeSliceTypeUpdateCallback != nil {
		var frontSafeAltitudeAreaTimeSliceType *SafeAltitudeAreaTimeSliceType
		if front != nil {
			frontSafeAltitudeAreaTimeSliceType, _ = front.(*SafeAltitudeAreaTimeSliceType)
		}
		stage.OnAfterSafeAltitudeAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, safealtitudeareatimeslicetype, frontSafeAltitudeAreaTimeSliceType)
	}
}

func (safealtitudeareatimeslicetype *SafeAltitudeAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaTimeSliceTypeDeleteCallback != nil {
		var frontSafeAltitudeAreaTimeSliceType *SafeAltitudeAreaTimeSliceType
		if front != nil {
			frontSafeAltitudeAreaTimeSliceType, _ = front.(*SafeAltitudeAreaTimeSliceType)
		}
		stage.OnAfterSafeAltitudeAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, safealtitudeareatimeslicetype, frontSafeAltitudeAreaTimeSliceType)
	}
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSafeAltitudeAreaTypeCreateCallback != nil {
		stage.OnAfterSafeAltitudeAreaTypeCreateCallback.OnAfterCreate(stage, safealtitudeareatype)
	}
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaTypeUpdateCallback != nil {
		var frontSafeAltitudeAreaType *SafeAltitudeAreaType
		if front != nil {
			frontSafeAltitudeAreaType, _ = front.(*SafeAltitudeAreaType)
		}
		stage.OnAfterSafeAltitudeAreaTypeUpdateCallback.OnAfterUpdate(stage, safealtitudeareatype, frontSafeAltitudeAreaType)
	}
}

func (safealtitudeareatype *SafeAltitudeAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSafeAltitudeAreaTypeDeleteCallback != nil {
		var frontSafeAltitudeAreaType *SafeAltitudeAreaType
		if front != nil {
			frontSafeAltitudeAreaType, _ = front.(*SafeAltitudeAreaType)
		}
		stage.OnAfterSafeAltitudeAreaTypeDeleteCallback.OnAfterDelete(stage, safealtitudeareatype, frontSafeAltitudeAreaType)
	}
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSectorDesignPropertyTypeCreateCallback != nil {
		stage.OnAfterSectorDesignPropertyTypeCreateCallback.OnAfterCreate(stage, sectordesignpropertytype)
	}
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSectorDesignPropertyTypeUpdateCallback != nil {
		var frontSectorDesignPropertyType *SectorDesignPropertyType
		if front != nil {
			frontSectorDesignPropertyType, _ = front.(*SectorDesignPropertyType)
		}
		stage.OnAfterSectorDesignPropertyTypeUpdateCallback.OnAfterUpdate(stage, sectordesignpropertytype, frontSectorDesignPropertyType)
	}
}

func (sectordesignpropertytype *SectorDesignPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSectorDesignPropertyTypeDeleteCallback != nil {
		var frontSectorDesignPropertyType *SectorDesignPropertyType
		if front != nil {
			frontSectorDesignPropertyType, _ = front.(*SectorDesignPropertyType)
		}
		stage.OnAfterSectorDesignPropertyTypeDeleteCallback.OnAfterDelete(stage, sectordesignpropertytype, frontSectorDesignPropertyType)
	}
}

func (sectordesigntype *SectorDesignType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSectorDesignTypeCreateCallback != nil {
		stage.OnAfterSectorDesignTypeCreateCallback.OnAfterCreate(stage, sectordesigntype)
	}
}

func (sectordesigntype *SectorDesignType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSectorDesignTypeUpdateCallback != nil {
		var frontSectorDesignType *SectorDesignType
		if front != nil {
			frontSectorDesignType, _ = front.(*SectorDesignType)
		}
		stage.OnAfterSectorDesignTypeUpdateCallback.OnAfterUpdate(stage, sectordesigntype, frontSectorDesignType)
	}
}

func (sectordesigntype *SectorDesignType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSectorDesignTypeDeleteCallback != nil {
		var frontSectorDesignType *SectorDesignType
		if front != nil {
			frontSectorDesignType, _ = front.(*SectorDesignType)
		}
		stage.OnAfterSectorDesignTypeDeleteCallback.OnAfterDelete(stage, sectordesigntype, frontSectorDesignType)
	}
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSegmentLegPropertyTypeCreateCallback != nil {
		stage.OnAfterSegmentLegPropertyTypeCreateCallback.OnAfterCreate(stage, segmentlegpropertytype)
	}
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSegmentLegPropertyTypeUpdateCallback != nil {
		var frontSegmentLegPropertyType *SegmentLegPropertyType
		if front != nil {
			frontSegmentLegPropertyType, _ = front.(*SegmentLegPropertyType)
		}
		stage.OnAfterSegmentLegPropertyTypeUpdateCallback.OnAfterUpdate(stage, segmentlegpropertytype, frontSegmentLegPropertyType)
	}
}

func (segmentlegpropertytype *SegmentLegPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSegmentLegPropertyTypeDeleteCallback != nil {
		var frontSegmentLegPropertyType *SegmentLegPropertyType
		if front != nil {
			frontSegmentLegPropertyType, _ = front.(*SegmentLegPropertyType)
		}
		stage.OnAfterSegmentLegPropertyTypeDeleteCallback.OnAfterDelete(stage, segmentlegpropertytype, frontSegmentLegPropertyType)
	}
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSegmentPointPropertyTypeCreateCallback != nil {
		stage.OnAfterSegmentPointPropertyTypeCreateCallback.OnAfterCreate(stage, segmentpointpropertytype)
	}
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSegmentPointPropertyTypeUpdateCallback != nil {
		var frontSegmentPointPropertyType *SegmentPointPropertyType
		if front != nil {
			frontSegmentPointPropertyType, _ = front.(*SegmentPointPropertyType)
		}
		stage.OnAfterSegmentPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, segmentpointpropertytype, frontSegmentPointPropertyType)
	}
}

func (segmentpointpropertytype *SegmentPointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSegmentPointPropertyTypeDeleteCallback != nil {
		var frontSegmentPointPropertyType *SegmentPointPropertyType
		if front != nil {
			frontSegmentPointPropertyType, _ = front.(*SegmentPointPropertyType)
		}
		stage.OnAfterSegmentPointPropertyTypeDeleteCallback.OnAfterDelete(stage, segmentpointpropertytype, frontSegmentPointPropertyType)
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

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardInstrumentArrivalPropertyTypeCreateCallback != nil {
		stage.OnAfterStandardInstrumentArrivalPropertyTypeCreateCallback.OnAfterCreate(stage, standardinstrumentarrivalpropertytype)
	}
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentArrivalPropertyTypeUpdateCallback != nil {
		var frontStandardInstrumentArrivalPropertyType *StandardInstrumentArrivalPropertyType
		if front != nil {
			frontStandardInstrumentArrivalPropertyType, _ = front.(*StandardInstrumentArrivalPropertyType)
		}
		stage.OnAfterStandardInstrumentArrivalPropertyTypeUpdateCallback.OnAfterUpdate(stage, standardinstrumentarrivalpropertytype, frontStandardInstrumentArrivalPropertyType)
	}
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentArrivalPropertyTypeDeleteCallback != nil {
		var frontStandardInstrumentArrivalPropertyType *StandardInstrumentArrivalPropertyType
		if front != nil {
			frontStandardInstrumentArrivalPropertyType, _ = front.(*StandardInstrumentArrivalPropertyType)
		}
		stage.OnAfterStandardInstrumentArrivalPropertyTypeDeleteCallback.OnAfterDelete(stage, standardinstrumentarrivalpropertytype, frontStandardInstrumentArrivalPropertyType)
	}
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardInstrumentArrivalTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterStandardInstrumentArrivalTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, standardinstrumentarrivaltimeslicepropertytype)
	}
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentArrivalTimeSlicePropertyTypeUpdateCallback != nil {
		var frontStandardInstrumentArrivalTimeSlicePropertyType *StandardInstrumentArrivalTimeSlicePropertyType
		if front != nil {
			frontStandardInstrumentArrivalTimeSlicePropertyType, _ = front.(*StandardInstrumentArrivalTimeSlicePropertyType)
		}
		stage.OnAfterStandardInstrumentArrivalTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, standardinstrumentarrivaltimeslicepropertytype, frontStandardInstrumentArrivalTimeSlicePropertyType)
	}
}

func (standardinstrumentarrivaltimeslicepropertytype *StandardInstrumentArrivalTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentArrivalTimeSlicePropertyTypeDeleteCallback != nil {
		var frontStandardInstrumentArrivalTimeSlicePropertyType *StandardInstrumentArrivalTimeSlicePropertyType
		if front != nil {
			frontStandardInstrumentArrivalTimeSlicePropertyType, _ = front.(*StandardInstrumentArrivalTimeSlicePropertyType)
		}
		stage.OnAfterStandardInstrumentArrivalTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, standardinstrumentarrivaltimeslicepropertytype, frontStandardInstrumentArrivalTimeSlicePropertyType)
	}
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardInstrumentArrivalTimeSliceTypeCreateCallback != nil {
		stage.OnAfterStandardInstrumentArrivalTimeSliceTypeCreateCallback.OnAfterCreate(stage, standardinstrumentarrivaltimeslicetype)
	}
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentArrivalTimeSliceTypeUpdateCallback != nil {
		var frontStandardInstrumentArrivalTimeSliceType *StandardInstrumentArrivalTimeSliceType
		if front != nil {
			frontStandardInstrumentArrivalTimeSliceType, _ = front.(*StandardInstrumentArrivalTimeSliceType)
		}
		stage.OnAfterStandardInstrumentArrivalTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, standardinstrumentarrivaltimeslicetype, frontStandardInstrumentArrivalTimeSliceType)
	}
}

func (standardinstrumentarrivaltimeslicetype *StandardInstrumentArrivalTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentArrivalTimeSliceTypeDeleteCallback != nil {
		var frontStandardInstrumentArrivalTimeSliceType *StandardInstrumentArrivalTimeSliceType
		if front != nil {
			frontStandardInstrumentArrivalTimeSliceType, _ = front.(*StandardInstrumentArrivalTimeSliceType)
		}
		stage.OnAfterStandardInstrumentArrivalTimeSliceTypeDeleteCallback.OnAfterDelete(stage, standardinstrumentarrivaltimeslicetype, frontStandardInstrumentArrivalTimeSliceType)
	}
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardInstrumentArrivalTypeCreateCallback != nil {
		stage.OnAfterStandardInstrumentArrivalTypeCreateCallback.OnAfterCreate(stage, standardinstrumentarrivaltype)
	}
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentArrivalTypeUpdateCallback != nil {
		var frontStandardInstrumentArrivalType *StandardInstrumentArrivalType
		if front != nil {
			frontStandardInstrumentArrivalType, _ = front.(*StandardInstrumentArrivalType)
		}
		stage.OnAfterStandardInstrumentArrivalTypeUpdateCallback.OnAfterUpdate(stage, standardinstrumentarrivaltype, frontStandardInstrumentArrivalType)
	}
}

func (standardinstrumentarrivaltype *StandardInstrumentArrivalType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentArrivalTypeDeleteCallback != nil {
		var frontStandardInstrumentArrivalType *StandardInstrumentArrivalType
		if front != nil {
			frontStandardInstrumentArrivalType, _ = front.(*StandardInstrumentArrivalType)
		}
		stage.OnAfterStandardInstrumentArrivalTypeDeleteCallback.OnAfterDelete(stage, standardinstrumentarrivaltype, frontStandardInstrumentArrivalType)
	}
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardInstrumentDeparturePropertyTypeCreateCallback != nil {
		stage.OnAfterStandardInstrumentDeparturePropertyTypeCreateCallback.OnAfterCreate(stage, standardinstrumentdeparturepropertytype)
	}
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentDeparturePropertyTypeUpdateCallback != nil {
		var frontStandardInstrumentDeparturePropertyType *StandardInstrumentDeparturePropertyType
		if front != nil {
			frontStandardInstrumentDeparturePropertyType, _ = front.(*StandardInstrumentDeparturePropertyType)
		}
		stage.OnAfterStandardInstrumentDeparturePropertyTypeUpdateCallback.OnAfterUpdate(stage, standardinstrumentdeparturepropertytype, frontStandardInstrumentDeparturePropertyType)
	}
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentDeparturePropertyTypeDeleteCallback != nil {
		var frontStandardInstrumentDeparturePropertyType *StandardInstrumentDeparturePropertyType
		if front != nil {
			frontStandardInstrumentDeparturePropertyType, _ = front.(*StandardInstrumentDeparturePropertyType)
		}
		stage.OnAfterStandardInstrumentDeparturePropertyTypeDeleteCallback.OnAfterDelete(stage, standardinstrumentdeparturepropertytype, frontStandardInstrumentDeparturePropertyType)
	}
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardInstrumentDepartureTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterStandardInstrumentDepartureTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, standardinstrumentdeparturetimeslicepropertytype)
	}
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentDepartureTimeSlicePropertyTypeUpdateCallback != nil {
		var frontStandardInstrumentDepartureTimeSlicePropertyType *StandardInstrumentDepartureTimeSlicePropertyType
		if front != nil {
			frontStandardInstrumentDepartureTimeSlicePropertyType, _ = front.(*StandardInstrumentDepartureTimeSlicePropertyType)
		}
		stage.OnAfterStandardInstrumentDepartureTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, standardinstrumentdeparturetimeslicepropertytype, frontStandardInstrumentDepartureTimeSlicePropertyType)
	}
}

func (standardinstrumentdeparturetimeslicepropertytype *StandardInstrumentDepartureTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentDepartureTimeSlicePropertyTypeDeleteCallback != nil {
		var frontStandardInstrumentDepartureTimeSlicePropertyType *StandardInstrumentDepartureTimeSlicePropertyType
		if front != nil {
			frontStandardInstrumentDepartureTimeSlicePropertyType, _ = front.(*StandardInstrumentDepartureTimeSlicePropertyType)
		}
		stage.OnAfterStandardInstrumentDepartureTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, standardinstrumentdeparturetimeslicepropertytype, frontStandardInstrumentDepartureTimeSlicePropertyType)
	}
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardInstrumentDepartureTimeSliceTypeCreateCallback != nil {
		stage.OnAfterStandardInstrumentDepartureTimeSliceTypeCreateCallback.OnAfterCreate(stage, standardinstrumentdeparturetimeslicetype)
	}
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentDepartureTimeSliceTypeUpdateCallback != nil {
		var frontStandardInstrumentDepartureTimeSliceType *StandardInstrumentDepartureTimeSliceType
		if front != nil {
			frontStandardInstrumentDepartureTimeSliceType, _ = front.(*StandardInstrumentDepartureTimeSliceType)
		}
		stage.OnAfterStandardInstrumentDepartureTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, standardinstrumentdeparturetimeslicetype, frontStandardInstrumentDepartureTimeSliceType)
	}
}

func (standardinstrumentdeparturetimeslicetype *StandardInstrumentDepartureTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentDepartureTimeSliceTypeDeleteCallback != nil {
		var frontStandardInstrumentDepartureTimeSliceType *StandardInstrumentDepartureTimeSliceType
		if front != nil {
			frontStandardInstrumentDepartureTimeSliceType, _ = front.(*StandardInstrumentDepartureTimeSliceType)
		}
		stage.OnAfterStandardInstrumentDepartureTimeSliceTypeDeleteCallback.OnAfterDelete(stage, standardinstrumentdeparturetimeslicetype, frontStandardInstrumentDepartureTimeSliceType)
	}
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardInstrumentDepartureTypeCreateCallback != nil {
		stage.OnAfterStandardInstrumentDepartureTypeCreateCallback.OnAfterCreate(stage, standardinstrumentdeparturetype)
	}
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentDepartureTypeUpdateCallback != nil {
		var frontStandardInstrumentDepartureType *StandardInstrumentDepartureType
		if front != nil {
			frontStandardInstrumentDepartureType, _ = front.(*StandardInstrumentDepartureType)
		}
		stage.OnAfterStandardInstrumentDepartureTypeUpdateCallback.OnAfterUpdate(stage, standardinstrumentdeparturetype, frontStandardInstrumentDepartureType)
	}
}

func (standardinstrumentdeparturetype *StandardInstrumentDepartureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardInstrumentDepartureTypeDeleteCallback != nil {
		var frontStandardInstrumentDepartureType *StandardInstrumentDepartureType
		if front != nil {
			frontStandardInstrumentDepartureType, _ = front.(*StandardInstrumentDepartureType)
		}
		stage.OnAfterStandardInstrumentDepartureTypeDeleteCallback.OnAfterDelete(stage, standardinstrumentdeparturetype, frontStandardInstrumentDepartureType)
	}
}

func (surfacepropertytype *SurfacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfacePropertyTypeCreateCallback != nil {
		stage.OnAfterSurfacePropertyTypeCreateCallback.OnAfterCreate(stage, surfacepropertytype)
	}
}

func (surfacepropertytype *SurfacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePropertyTypeUpdateCallback != nil {
		var frontSurfacePropertyType *SurfacePropertyType
		if front != nil {
			frontSurfacePropertyType, _ = front.(*SurfacePropertyType)
		}
		stage.OnAfterSurfacePropertyTypeUpdateCallback.OnAfterUpdate(stage, surfacepropertytype, frontSurfacePropertyType)
	}
}

func (surfacepropertytype *SurfacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePropertyTypeDeleteCallback != nil {
		var frontSurfacePropertyType *SurfacePropertyType
		if front != nil {
			frontSurfacePropertyType, _ = front.(*SurfacePropertyType)
		}
		stage.OnAfterSurfacePropertyTypeDeleteCallback.OnAfterDelete(stage, surfacepropertytype, frontSurfacePropertyType)
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

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTerminalArrivalAreaSectorPropertyTypeCreateCallback != nil {
		stage.OnAfterTerminalArrivalAreaSectorPropertyTypeCreateCallback.OnAfterCreate(stage, terminalarrivalareasectorpropertytype)
	}
}

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaSectorPropertyTypeUpdateCallback != nil {
		var frontTerminalArrivalAreaSectorPropertyType *TerminalArrivalAreaSectorPropertyType
		if front != nil {
			frontTerminalArrivalAreaSectorPropertyType, _ = front.(*TerminalArrivalAreaSectorPropertyType)
		}
		stage.OnAfterTerminalArrivalAreaSectorPropertyTypeUpdateCallback.OnAfterUpdate(stage, terminalarrivalareasectorpropertytype, frontTerminalArrivalAreaSectorPropertyType)
	}
}

func (terminalarrivalareasectorpropertytype *TerminalArrivalAreaSectorPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaSectorPropertyTypeDeleteCallback != nil {
		var frontTerminalArrivalAreaSectorPropertyType *TerminalArrivalAreaSectorPropertyType
		if front != nil {
			frontTerminalArrivalAreaSectorPropertyType, _ = front.(*TerminalArrivalAreaSectorPropertyType)
		}
		stage.OnAfterTerminalArrivalAreaSectorPropertyTypeDeleteCallback.OnAfterDelete(stage, terminalarrivalareasectorpropertytype, frontTerminalArrivalAreaSectorPropertyType)
	}
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTerminalArrivalAreaSectorTypeCreateCallback != nil {
		stage.OnAfterTerminalArrivalAreaSectorTypeCreateCallback.OnAfterCreate(stage, terminalarrivalareasectortype)
	}
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaSectorTypeUpdateCallback != nil {
		var frontTerminalArrivalAreaSectorType *TerminalArrivalAreaSectorType
		if front != nil {
			frontTerminalArrivalAreaSectorType, _ = front.(*TerminalArrivalAreaSectorType)
		}
		stage.OnAfterTerminalArrivalAreaSectorTypeUpdateCallback.OnAfterUpdate(stage, terminalarrivalareasectortype, frontTerminalArrivalAreaSectorType)
	}
}

func (terminalarrivalareasectortype *TerminalArrivalAreaSectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaSectorTypeDeleteCallback != nil {
		var frontTerminalArrivalAreaSectorType *TerminalArrivalAreaSectorType
		if front != nil {
			frontTerminalArrivalAreaSectorType, _ = front.(*TerminalArrivalAreaSectorType)
		}
		stage.OnAfterTerminalArrivalAreaSectorTypeDeleteCallback.OnAfterDelete(stage, terminalarrivalareasectortype, frontTerminalArrivalAreaSectorType)
	}
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTerminalArrivalAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterTerminalArrivalAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, terminalarrivalareatimeslicepropertytype)
	}
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontTerminalArrivalAreaTimeSlicePropertyType *TerminalArrivalAreaTimeSlicePropertyType
		if front != nil {
			frontTerminalArrivalAreaTimeSlicePropertyType, _ = front.(*TerminalArrivalAreaTimeSlicePropertyType)
		}
		stage.OnAfterTerminalArrivalAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, terminalarrivalareatimeslicepropertytype, frontTerminalArrivalAreaTimeSlicePropertyType)
	}
}

func (terminalarrivalareatimeslicepropertytype *TerminalArrivalAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontTerminalArrivalAreaTimeSlicePropertyType *TerminalArrivalAreaTimeSlicePropertyType
		if front != nil {
			frontTerminalArrivalAreaTimeSlicePropertyType, _ = front.(*TerminalArrivalAreaTimeSlicePropertyType)
		}
		stage.OnAfterTerminalArrivalAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, terminalarrivalareatimeslicepropertytype, frontTerminalArrivalAreaTimeSlicePropertyType)
	}
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTerminalArrivalAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterTerminalArrivalAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, terminalarrivalareatimeslicetype)
	}
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaTimeSliceTypeUpdateCallback != nil {
		var frontTerminalArrivalAreaTimeSliceType *TerminalArrivalAreaTimeSliceType
		if front != nil {
			frontTerminalArrivalAreaTimeSliceType, _ = front.(*TerminalArrivalAreaTimeSliceType)
		}
		stage.OnAfterTerminalArrivalAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, terminalarrivalareatimeslicetype, frontTerminalArrivalAreaTimeSliceType)
	}
}

func (terminalarrivalareatimeslicetype *TerminalArrivalAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaTimeSliceTypeDeleteCallback != nil {
		var frontTerminalArrivalAreaTimeSliceType *TerminalArrivalAreaTimeSliceType
		if front != nil {
			frontTerminalArrivalAreaTimeSliceType, _ = front.(*TerminalArrivalAreaTimeSliceType)
		}
		stage.OnAfterTerminalArrivalAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, terminalarrivalareatimeslicetype, frontTerminalArrivalAreaTimeSliceType)
	}
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTerminalArrivalAreaTypeCreateCallback != nil {
		stage.OnAfterTerminalArrivalAreaTypeCreateCallback.OnAfterCreate(stage, terminalarrivalareatype)
	}
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaTypeUpdateCallback != nil {
		var frontTerminalArrivalAreaType *TerminalArrivalAreaType
		if front != nil {
			frontTerminalArrivalAreaType, _ = front.(*TerminalArrivalAreaType)
		}
		stage.OnAfterTerminalArrivalAreaTypeUpdateCallback.OnAfterUpdate(stage, terminalarrivalareatype, frontTerminalArrivalAreaType)
	}
}

func (terminalarrivalareatype *TerminalArrivalAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalArrivalAreaTypeDeleteCallback != nil {
		var frontTerminalArrivalAreaType *TerminalArrivalAreaType
		if front != nil {
			frontTerminalArrivalAreaType, _ = front.(*TerminalArrivalAreaType)
		}
		stage.OnAfterTerminalArrivalAreaTypeDeleteCallback.OnAfterDelete(stage, terminalarrivalareatype, frontTerminalArrivalAreaType)
	}
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTerminalSegmentPointPropertyTypeCreateCallback != nil {
		stage.OnAfterTerminalSegmentPointPropertyTypeCreateCallback.OnAfterCreate(stage, terminalsegmentpointpropertytype)
	}
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalSegmentPointPropertyTypeUpdateCallback != nil {
		var frontTerminalSegmentPointPropertyType *TerminalSegmentPointPropertyType
		if front != nil {
			frontTerminalSegmentPointPropertyType, _ = front.(*TerminalSegmentPointPropertyType)
		}
		stage.OnAfterTerminalSegmentPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, terminalsegmentpointpropertytype, frontTerminalSegmentPointPropertyType)
	}
}

func (terminalsegmentpointpropertytype *TerminalSegmentPointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalSegmentPointPropertyTypeDeleteCallback != nil {
		var frontTerminalSegmentPointPropertyType *TerminalSegmentPointPropertyType
		if front != nil {
			frontTerminalSegmentPointPropertyType, _ = front.(*TerminalSegmentPointPropertyType)
		}
		stage.OnAfterTerminalSegmentPointPropertyTypeDeleteCallback.OnAfterDelete(stage, terminalsegmentpointpropertytype, frontTerminalSegmentPointPropertyType)
	}
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTerminalSegmentPointTypeCreateCallback != nil {
		stage.OnAfterTerminalSegmentPointTypeCreateCallback.OnAfterCreate(stage, terminalsegmentpointtype)
	}
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalSegmentPointTypeUpdateCallback != nil {
		var frontTerminalSegmentPointType *TerminalSegmentPointType
		if front != nil {
			frontTerminalSegmentPointType, _ = front.(*TerminalSegmentPointType)
		}
		stage.OnAfterTerminalSegmentPointTypeUpdateCallback.OnAfterUpdate(stage, terminalsegmentpointtype, frontTerminalSegmentPointType)
	}
}

func (terminalsegmentpointtype *TerminalSegmentPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTerminalSegmentPointTypeDeleteCallback != nil {
		var frontTerminalSegmentPointType *TerminalSegmentPointType
		if front != nil {
			frontTerminalSegmentPointType, _ = front.(*TerminalSegmentPointType)
		}
		stage.OnAfterTerminalSegmentPointTypeDeleteCallback.OnAfterDelete(stage, terminalsegmentpointtype, frontTerminalSegmentPointType)
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

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnplannedHoldingPropertyTypeCreateCallback != nil {
		stage.OnAfterUnplannedHoldingPropertyTypeCreateCallback.OnAfterCreate(stage, unplannedholdingpropertytype)
	}
}

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnplannedHoldingPropertyTypeUpdateCallback != nil {
		var frontUnplannedHoldingPropertyType *UnplannedHoldingPropertyType
		if front != nil {
			frontUnplannedHoldingPropertyType, _ = front.(*UnplannedHoldingPropertyType)
		}
		stage.OnAfterUnplannedHoldingPropertyTypeUpdateCallback.OnAfterUpdate(stage, unplannedholdingpropertytype, frontUnplannedHoldingPropertyType)
	}
}

func (unplannedholdingpropertytype *UnplannedHoldingPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnplannedHoldingPropertyTypeDeleteCallback != nil {
		var frontUnplannedHoldingPropertyType *UnplannedHoldingPropertyType
		if front != nil {
			frontUnplannedHoldingPropertyType, _ = front.(*UnplannedHoldingPropertyType)
		}
		stage.OnAfterUnplannedHoldingPropertyTypeDeleteCallback.OnAfterDelete(stage, unplannedholdingpropertytype, frontUnplannedHoldingPropertyType)
	}
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnplannedHoldingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterUnplannedHoldingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, unplannedholdingtimeslicepropertytype)
	}
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnplannedHoldingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontUnplannedHoldingTimeSlicePropertyType *UnplannedHoldingTimeSlicePropertyType
		if front != nil {
			frontUnplannedHoldingTimeSlicePropertyType, _ = front.(*UnplannedHoldingTimeSlicePropertyType)
		}
		stage.OnAfterUnplannedHoldingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, unplannedholdingtimeslicepropertytype, frontUnplannedHoldingTimeSlicePropertyType)
	}
}

func (unplannedholdingtimeslicepropertytype *UnplannedHoldingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnplannedHoldingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontUnplannedHoldingTimeSlicePropertyType *UnplannedHoldingTimeSlicePropertyType
		if front != nil {
			frontUnplannedHoldingTimeSlicePropertyType, _ = front.(*UnplannedHoldingTimeSlicePropertyType)
		}
		stage.OnAfterUnplannedHoldingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, unplannedholdingtimeslicepropertytype, frontUnplannedHoldingTimeSlicePropertyType)
	}
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnplannedHoldingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterUnplannedHoldingTimeSliceTypeCreateCallback.OnAfterCreate(stage, unplannedholdingtimeslicetype)
	}
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnplannedHoldingTimeSliceTypeUpdateCallback != nil {
		var frontUnplannedHoldingTimeSliceType *UnplannedHoldingTimeSliceType
		if front != nil {
			frontUnplannedHoldingTimeSliceType, _ = front.(*UnplannedHoldingTimeSliceType)
		}
		stage.OnAfterUnplannedHoldingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, unplannedholdingtimeslicetype, frontUnplannedHoldingTimeSliceType)
	}
}

func (unplannedholdingtimeslicetype *UnplannedHoldingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnplannedHoldingTimeSliceTypeDeleteCallback != nil {
		var frontUnplannedHoldingTimeSliceType *UnplannedHoldingTimeSliceType
		if front != nil {
			frontUnplannedHoldingTimeSliceType, _ = front.(*UnplannedHoldingTimeSliceType)
		}
		stage.OnAfterUnplannedHoldingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, unplannedholdingtimeslicetype, frontUnplannedHoldingTimeSliceType)
	}
}

func (unplannedholdingtype *UnplannedHoldingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnplannedHoldingTypeCreateCallback != nil {
		stage.OnAfterUnplannedHoldingTypeCreateCallback.OnAfterCreate(stage, unplannedholdingtype)
	}
}

func (unplannedholdingtype *UnplannedHoldingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnplannedHoldingTypeUpdateCallback != nil {
		var frontUnplannedHoldingType *UnplannedHoldingType
		if front != nil {
			frontUnplannedHoldingType, _ = front.(*UnplannedHoldingType)
		}
		stage.OnAfterUnplannedHoldingTypeUpdateCallback.OnAfterUpdate(stage, unplannedholdingtype, frontUnplannedHoldingType)
	}
}

func (unplannedholdingtype *UnplannedHoldingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnplannedHoldingTypeDeleteCallback != nil {
		var frontUnplannedHoldingType *UnplannedHoldingType
		if front != nil {
			frontUnplannedHoldingType, _ = front.(*UnplannedHoldingType)
		}
		stage.OnAfterUnplannedHoldingTypeDeleteCallback.OnAfterDelete(stage, unplannedholdingtype, frontUnplannedHoldingType)
	}
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalStructurePropertyTypeCreateCallback != nil {
		stage.OnAfterVerticalStructurePropertyTypeCreateCallback.OnAfterCreate(stage, verticalstructurepropertytype)
	}
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructurePropertyTypeUpdateCallback != nil {
		var frontVerticalStructurePropertyType *VerticalStructurePropertyType
		if front != nil {
			frontVerticalStructurePropertyType, _ = front.(*VerticalStructurePropertyType)
		}
		stage.OnAfterVerticalStructurePropertyTypeUpdateCallback.OnAfterUpdate(stage, verticalstructurepropertytype, frontVerticalStructurePropertyType)
	}
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructurePropertyTypeDeleteCallback != nil {
		var frontVerticalStructurePropertyType *VerticalStructurePropertyType
		if front != nil {
			frontVerticalStructurePropertyType, _ = front.(*VerticalStructurePropertyType)
		}
		stage.OnAfterVerticalStructurePropertyTypeDeleteCallback.OnAfterDelete(stage, verticalstructurepropertytype, frontVerticalStructurePropertyType)
	}
}

