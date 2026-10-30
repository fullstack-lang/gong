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

func (abstractdirectflighttype *AbstractDirectFlightType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDirectFlightTypeCreateCallback != nil {
		stage.OnAfterAbstractDirectFlightTypeCreateCallback.OnAfterCreate(stage, abstractdirectflighttype)
	}
}

func (abstractdirectflighttype *AbstractDirectFlightType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDirectFlightTypeUpdateCallback != nil {
		var frontAbstractDirectFlightType *AbstractDirectFlightType
		if front != nil {
			frontAbstractDirectFlightType, _ = front.(*AbstractDirectFlightType)
		}
		stage.OnAfterAbstractDirectFlightTypeUpdateCallback.OnAfterUpdate(stage, abstractdirectflighttype, frontAbstractDirectFlightType)
	}
}

func (abstractdirectflighttype *AbstractDirectFlightType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDirectFlightTypeDeleteCallback != nil {
		var frontAbstractDirectFlightType *AbstractDirectFlightType
		if front != nil {
			frontAbstractDirectFlightType, _ = front.(*AbstractDirectFlightType)
		}
		stage.OnAfterAbstractDirectFlightTypeDeleteCallback.OnAfterDelete(stage, abstractdirectflighttype, frontAbstractDirectFlightType)
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

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingAnchorPropertyTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingAnchorPropertyTypeCreateCallback.OnAfterCreate(stage, aerialrefuellinganchorpropertytype)
	}
}

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingAnchorPropertyTypeUpdateCallback != nil {
		var frontAerialRefuellingAnchorPropertyType *AerialRefuellingAnchorPropertyType
		if front != nil {
			frontAerialRefuellingAnchorPropertyType, _ = front.(*AerialRefuellingAnchorPropertyType)
		}
		stage.OnAfterAerialRefuellingAnchorPropertyTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellinganchorpropertytype, frontAerialRefuellingAnchorPropertyType)
	}
}

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingAnchorPropertyTypeDeleteCallback != nil {
		var frontAerialRefuellingAnchorPropertyType *AerialRefuellingAnchorPropertyType
		if front != nil {
			frontAerialRefuellingAnchorPropertyType, _ = front.(*AerialRefuellingAnchorPropertyType)
		}
		stage.OnAfterAerialRefuellingAnchorPropertyTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellinganchorpropertytype, frontAerialRefuellingAnchorPropertyType)
	}
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingAnchorTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingAnchorTypeCreateCallback.OnAfterCreate(stage, aerialrefuellinganchortype)
	}
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingAnchorTypeUpdateCallback != nil {
		var frontAerialRefuellingAnchorType *AerialRefuellingAnchorType
		if front != nil {
			frontAerialRefuellingAnchorType, _ = front.(*AerialRefuellingAnchorType)
		}
		stage.OnAfterAerialRefuellingAnchorTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellinganchortype, frontAerialRefuellingAnchorType)
	}
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingAnchorTypeDeleteCallback != nil {
		var frontAerialRefuellingAnchorType *AerialRefuellingAnchorType
		if front != nil {
			frontAerialRefuellingAnchorType, _ = front.(*AerialRefuellingAnchorType)
		}
		stage.OnAfterAerialRefuellingAnchorTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellinganchortype, frontAerialRefuellingAnchorType)
	}
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingPointPropertyTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingPointPropertyTypeCreateCallback.OnAfterCreate(stage, aerialrefuellingpointpropertytype)
	}
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingPointPropertyTypeUpdateCallback != nil {
		var frontAerialRefuellingPointPropertyType *AerialRefuellingPointPropertyType
		if front != nil {
			frontAerialRefuellingPointPropertyType, _ = front.(*AerialRefuellingPointPropertyType)
		}
		stage.OnAfterAerialRefuellingPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellingpointpropertytype, frontAerialRefuellingPointPropertyType)
	}
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingPointPropertyTypeDeleteCallback != nil {
		var frontAerialRefuellingPointPropertyType *AerialRefuellingPointPropertyType
		if front != nil {
			frontAerialRefuellingPointPropertyType, _ = front.(*AerialRefuellingPointPropertyType)
		}
		stage.OnAfterAerialRefuellingPointPropertyTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellingpointpropertytype, frontAerialRefuellingPointPropertyType)
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

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingPropertyTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingPropertyTypeCreateCallback.OnAfterCreate(stage, aerialrefuellingpropertytype)
	}
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingPropertyTypeUpdateCallback != nil {
		var frontAerialRefuellingPropertyType *AerialRefuellingPropertyType
		if front != nil {
			frontAerialRefuellingPropertyType, _ = front.(*AerialRefuellingPropertyType)
		}
		stage.OnAfterAerialRefuellingPropertyTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellingpropertytype, frontAerialRefuellingPropertyType)
	}
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingPropertyTypeDeleteCallback != nil {
		var frontAerialRefuellingPropertyType *AerialRefuellingPropertyType
		if front != nil {
			frontAerialRefuellingPropertyType, _ = front.(*AerialRefuellingPropertyType)
		}
		stage.OnAfterAerialRefuellingPropertyTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellingpropertytype, frontAerialRefuellingPropertyType)
	}
}

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, aerialrefuellingtimeslicepropertytype)
	}
}

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAerialRefuellingTimeSlicePropertyType *AerialRefuellingTimeSlicePropertyType
		if front != nil {
			frontAerialRefuellingTimeSlicePropertyType, _ = front.(*AerialRefuellingTimeSlicePropertyType)
		}
		stage.OnAfterAerialRefuellingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellingtimeslicepropertytype, frontAerialRefuellingTimeSlicePropertyType)
	}
}

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAerialRefuellingTimeSlicePropertyType *AerialRefuellingTimeSlicePropertyType
		if front != nil {
			frontAerialRefuellingTimeSlicePropertyType, _ = front.(*AerialRefuellingTimeSlicePropertyType)
		}
		stage.OnAfterAerialRefuellingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellingtimeslicepropertytype, frontAerialRefuellingTimeSlicePropertyType)
	}
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingTimeSliceTypeCreateCallback.OnAfterCreate(stage, aerialrefuellingtimeslicetype)
	}
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTimeSliceTypeUpdateCallback != nil {
		var frontAerialRefuellingTimeSliceType *AerialRefuellingTimeSliceType
		if front != nil {
			frontAerialRefuellingTimeSliceType, _ = front.(*AerialRefuellingTimeSliceType)
		}
		stage.OnAfterAerialRefuellingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellingtimeslicetype, frontAerialRefuellingTimeSliceType)
	}
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTimeSliceTypeDeleteCallback != nil {
		var frontAerialRefuellingTimeSliceType *AerialRefuellingTimeSliceType
		if front != nil {
			frontAerialRefuellingTimeSliceType, _ = front.(*AerialRefuellingTimeSliceType)
		}
		stage.OnAfterAerialRefuellingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellingtimeslicetype, frontAerialRefuellingTimeSliceType)
	}
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingTrackPropertyTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingTrackPropertyTypeCreateCallback.OnAfterCreate(stage, aerialrefuellingtrackpropertytype)
	}
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTrackPropertyTypeUpdateCallback != nil {
		var frontAerialRefuellingTrackPropertyType *AerialRefuellingTrackPropertyType
		if front != nil {
			frontAerialRefuellingTrackPropertyType, _ = front.(*AerialRefuellingTrackPropertyType)
		}
		stage.OnAfterAerialRefuellingTrackPropertyTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellingtrackpropertytype, frontAerialRefuellingTrackPropertyType)
	}
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTrackPropertyTypeDeleteCallback != nil {
		var frontAerialRefuellingTrackPropertyType *AerialRefuellingTrackPropertyType
		if front != nil {
			frontAerialRefuellingTrackPropertyType, _ = front.(*AerialRefuellingTrackPropertyType)
		}
		stage.OnAfterAerialRefuellingTrackPropertyTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellingtrackpropertytype, frontAerialRefuellingTrackPropertyType)
	}
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingTrackTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingTrackTypeCreateCallback.OnAfterCreate(stage, aerialrefuellingtracktype)
	}
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTrackTypeUpdateCallback != nil {
		var frontAerialRefuellingTrackType *AerialRefuellingTrackType
		if front != nil {
			frontAerialRefuellingTrackType, _ = front.(*AerialRefuellingTrackType)
		}
		stage.OnAfterAerialRefuellingTrackTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellingtracktype, frontAerialRefuellingTrackType)
	}
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTrackTypeDeleteCallback != nil {
		var frontAerialRefuellingTrackType *AerialRefuellingTrackType
		if front != nil {
			frontAerialRefuellingTrackType, _ = front.(*AerialRefuellingTrackType)
		}
		stage.OnAfterAerialRefuellingTrackTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellingtracktype, frontAerialRefuellingTrackType)
	}
}

func (aerialrefuellingtype *AerialRefuellingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAerialRefuellingTypeCreateCallback != nil {
		stage.OnAfterAerialRefuellingTypeCreateCallback.OnAfterCreate(stage, aerialrefuellingtype)
	}
}

func (aerialrefuellingtype *AerialRefuellingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTypeUpdateCallback != nil {
		var frontAerialRefuellingType *AerialRefuellingType
		if front != nil {
			frontAerialRefuellingType, _ = front.(*AerialRefuellingType)
		}
		stage.OnAfterAerialRefuellingTypeUpdateCallback.OnAfterUpdate(stage, aerialrefuellingtype, frontAerialRefuellingType)
	}
}

func (aerialrefuellingtype *AerialRefuellingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAerialRefuellingTypeDeleteCallback != nil {
		var frontAerialRefuellingType *AerialRefuellingType
		if front != nil {
			frontAerialRefuellingType, _ = front.(*AerialRefuellingType)
		}
		stage.OnAfterAerialRefuellingTypeDeleteCallback.OnAfterDelete(stage, aerialrefuellingtype, frontAerialRefuellingType)
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

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceBorderCrossingPropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceBorderCrossingPropertyTypeCreateCallback.OnAfterCreate(stage, airspacebordercrossingpropertytype)
	}
}

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceBorderCrossingPropertyTypeUpdateCallback != nil {
		var frontAirspaceBorderCrossingPropertyType *AirspaceBorderCrossingPropertyType
		if front != nil {
			frontAirspaceBorderCrossingPropertyType, _ = front.(*AirspaceBorderCrossingPropertyType)
		}
		stage.OnAfterAirspaceBorderCrossingPropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacebordercrossingpropertytype, frontAirspaceBorderCrossingPropertyType)
	}
}

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceBorderCrossingPropertyTypeDeleteCallback != nil {
		var frontAirspaceBorderCrossingPropertyType *AirspaceBorderCrossingPropertyType
		if front != nil {
			frontAirspaceBorderCrossingPropertyType, _ = front.(*AirspaceBorderCrossingPropertyType)
		}
		stage.OnAfterAirspaceBorderCrossingPropertyTypeDeleteCallback.OnAfterDelete(stage, airspacebordercrossingpropertytype, frontAirspaceBorderCrossingPropertyType)
	}
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceBorderCrossingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceBorderCrossingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airspacebordercrossingtimeslicepropertytype)
	}
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceBorderCrossingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirspaceBorderCrossingTimeSlicePropertyType *AirspaceBorderCrossingTimeSlicePropertyType
		if front != nil {
			frontAirspaceBorderCrossingTimeSlicePropertyType, _ = front.(*AirspaceBorderCrossingTimeSlicePropertyType)
		}
		stage.OnAfterAirspaceBorderCrossingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacebordercrossingtimeslicepropertytype, frontAirspaceBorderCrossingTimeSlicePropertyType)
	}
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceBorderCrossingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirspaceBorderCrossingTimeSlicePropertyType *AirspaceBorderCrossingTimeSlicePropertyType
		if front != nil {
			frontAirspaceBorderCrossingTimeSlicePropertyType, _ = front.(*AirspaceBorderCrossingTimeSlicePropertyType)
		}
		stage.OnAfterAirspaceBorderCrossingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airspacebordercrossingtimeslicepropertytype, frontAirspaceBorderCrossingTimeSlicePropertyType)
	}
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceBorderCrossingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirspaceBorderCrossingTimeSliceTypeCreateCallback.OnAfterCreate(stage, airspacebordercrossingtimeslicetype)
	}
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceBorderCrossingTimeSliceTypeUpdateCallback != nil {
		var frontAirspaceBorderCrossingTimeSliceType *AirspaceBorderCrossingTimeSliceType
		if front != nil {
			frontAirspaceBorderCrossingTimeSliceType, _ = front.(*AirspaceBorderCrossingTimeSliceType)
		}
		stage.OnAfterAirspaceBorderCrossingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airspacebordercrossingtimeslicetype, frontAirspaceBorderCrossingTimeSliceType)
	}
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceBorderCrossingTimeSliceTypeDeleteCallback != nil {
		var frontAirspaceBorderCrossingTimeSliceType *AirspaceBorderCrossingTimeSliceType
		if front != nil {
			frontAirspaceBorderCrossingTimeSliceType, _ = front.(*AirspaceBorderCrossingTimeSliceType)
		}
		stage.OnAfterAirspaceBorderCrossingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airspacebordercrossingtimeslicetype, frontAirspaceBorderCrossingTimeSliceType)
	}
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceBorderCrossingTypeCreateCallback != nil {
		stage.OnAfterAirspaceBorderCrossingTypeCreateCallback.OnAfterCreate(stage, airspacebordercrossingtype)
	}
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceBorderCrossingTypeUpdateCallback != nil {
		var frontAirspaceBorderCrossingType *AirspaceBorderCrossingType
		if front != nil {
			frontAirspaceBorderCrossingType, _ = front.(*AirspaceBorderCrossingType)
		}
		stage.OnAfterAirspaceBorderCrossingTypeUpdateCallback.OnAfterUpdate(stage, airspacebordercrossingtype, frontAirspaceBorderCrossingType)
	}
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceBorderCrossingTypeDeleteCallback != nil {
		var frontAirspaceBorderCrossingType *AirspaceBorderCrossingType
		if front != nil {
			frontAirspaceBorderCrossingType, _ = front.(*AirspaceBorderCrossingType)
		}
		stage.OnAfterAirspaceBorderCrossingTypeDeleteCallback.OnAfterDelete(stage, airspacebordercrossingtype, frontAirspaceBorderCrossingType)
	}
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceLayerPropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceLayerPropertyTypeCreateCallback.OnAfterCreate(stage, airspacelayerpropertytype)
	}
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerPropertyTypeUpdateCallback != nil {
		var frontAirspaceLayerPropertyType *AirspaceLayerPropertyType
		if front != nil {
			frontAirspaceLayerPropertyType, _ = front.(*AirspaceLayerPropertyType)
		}
		stage.OnAfterAirspaceLayerPropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacelayerpropertytype, frontAirspaceLayerPropertyType)
	}
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerPropertyTypeDeleteCallback != nil {
		var frontAirspaceLayerPropertyType *AirspaceLayerPropertyType
		if front != nil {
			frontAirspaceLayerPropertyType, _ = front.(*AirspaceLayerPropertyType)
		}
		stage.OnAfterAirspaceLayerPropertyTypeDeleteCallback.OnAfterDelete(stage, airspacelayerpropertytype, frontAirspaceLayerPropertyType)
	}
}

func (airspacelayertype *AirspaceLayerType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceLayerTypeCreateCallback != nil {
		stage.OnAfterAirspaceLayerTypeCreateCallback.OnAfterCreate(stage, airspacelayertype)
	}
}

func (airspacelayertype *AirspaceLayerType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerTypeUpdateCallback != nil {
		var frontAirspaceLayerType *AirspaceLayerType
		if front != nil {
			frontAirspaceLayerType, _ = front.(*AirspaceLayerType)
		}
		stage.OnAfterAirspaceLayerTypeUpdateCallback.OnAfterUpdate(stage, airspacelayertype, frontAirspaceLayerType)
	}
}

func (airspacelayertype *AirspaceLayerType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerTypeDeleteCallback != nil {
		var frontAirspaceLayerType *AirspaceLayerType
		if front != nil {
			frontAirspaceLayerType, _ = front.(*AirspaceLayerType)
		}
		stage.OnAfterAirspaceLayerTypeDeleteCallback.OnAfterDelete(stage, airspacelayertype, frontAirspaceLayerType)
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

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForAerialRefuellingPropertyTypeCreateCallback != nil {
		stage.OnAfterAuthorityForAerialRefuellingPropertyTypeCreateCallback.OnAfterCreate(stage, authorityforaerialrefuellingpropertytype)
	}
}

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAerialRefuellingPropertyTypeUpdateCallback != nil {
		var frontAuthorityForAerialRefuellingPropertyType *AuthorityForAerialRefuellingPropertyType
		if front != nil {
			frontAuthorityForAerialRefuellingPropertyType, _ = front.(*AuthorityForAerialRefuellingPropertyType)
		}
		stage.OnAfterAuthorityForAerialRefuellingPropertyTypeUpdateCallback.OnAfterUpdate(stage, authorityforaerialrefuellingpropertytype, frontAuthorityForAerialRefuellingPropertyType)
	}
}

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAerialRefuellingPropertyTypeDeleteCallback != nil {
		var frontAuthorityForAerialRefuellingPropertyType *AuthorityForAerialRefuellingPropertyType
		if front != nil {
			frontAuthorityForAerialRefuellingPropertyType, _ = front.(*AuthorityForAerialRefuellingPropertyType)
		}
		stage.OnAfterAuthorityForAerialRefuellingPropertyTypeDeleteCallback.OnAfterDelete(stage, authorityforaerialrefuellingpropertytype, frontAuthorityForAerialRefuellingPropertyType)
	}
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForAerialRefuellingTypeCreateCallback != nil {
		stage.OnAfterAuthorityForAerialRefuellingTypeCreateCallback.OnAfterCreate(stage, authorityforaerialrefuellingtype)
	}
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAerialRefuellingTypeUpdateCallback != nil {
		var frontAuthorityForAerialRefuellingType *AuthorityForAerialRefuellingType
		if front != nil {
			frontAuthorityForAerialRefuellingType, _ = front.(*AuthorityForAerialRefuellingType)
		}
		stage.OnAfterAuthorityForAerialRefuellingTypeUpdateCallback.OnAfterUpdate(stage, authorityforaerialrefuellingtype, frontAuthorityForAerialRefuellingType)
	}
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAerialRefuellingTypeDeleteCallback != nil {
		var frontAuthorityForAerialRefuellingType *AuthorityForAerialRefuellingType
		if front != nil {
			frontAuthorityForAerialRefuellingType, _ = front.(*AuthorityForAerialRefuellingType)
		}
		stage.OnAfterAuthorityForAerialRefuellingTypeDeleteCallback.OnAfterDelete(stage, authorityforaerialrefuellingtype, frontAuthorityForAerialRefuellingType)
	}
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterChangeOverPointTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterChangeOverPointTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, changeoverpointtimeslicepropertytype)
	}
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChangeOverPointTimeSlicePropertyTypeUpdateCallback != nil {
		var frontChangeOverPointTimeSlicePropertyType *ChangeOverPointTimeSlicePropertyType
		if front != nil {
			frontChangeOverPointTimeSlicePropertyType, _ = front.(*ChangeOverPointTimeSlicePropertyType)
		}
		stage.OnAfterChangeOverPointTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, changeoverpointtimeslicepropertytype, frontChangeOverPointTimeSlicePropertyType)
	}
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChangeOverPointTimeSlicePropertyTypeDeleteCallback != nil {
		var frontChangeOverPointTimeSlicePropertyType *ChangeOverPointTimeSlicePropertyType
		if front != nil {
			frontChangeOverPointTimeSlicePropertyType, _ = front.(*ChangeOverPointTimeSlicePropertyType)
		}
		stage.OnAfterChangeOverPointTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, changeoverpointtimeslicepropertytype, frontChangeOverPointTimeSlicePropertyType)
	}
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterChangeOverPointTimeSliceTypeCreateCallback != nil {
		stage.OnAfterChangeOverPointTimeSliceTypeCreateCallback.OnAfterCreate(stage, changeoverpointtimeslicetype)
	}
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChangeOverPointTimeSliceTypeUpdateCallback != nil {
		var frontChangeOverPointTimeSliceType *ChangeOverPointTimeSliceType
		if front != nil {
			frontChangeOverPointTimeSliceType, _ = front.(*ChangeOverPointTimeSliceType)
		}
		stage.OnAfterChangeOverPointTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, changeoverpointtimeslicetype, frontChangeOverPointTimeSliceType)
	}
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChangeOverPointTimeSliceTypeDeleteCallback != nil {
		var frontChangeOverPointTimeSliceType *ChangeOverPointTimeSliceType
		if front != nil {
			frontChangeOverPointTimeSliceType, _ = front.(*ChangeOverPointTimeSliceType)
		}
		stage.OnAfterChangeOverPointTimeSliceTypeDeleteCallback.OnAfterDelete(stage, changeoverpointtimeslicetype, frontChangeOverPointTimeSliceType)
	}
}

func (changeoverpointtype *ChangeOverPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterChangeOverPointTypeCreateCallback != nil {
		stage.OnAfterChangeOverPointTypeCreateCallback.OnAfterCreate(stage, changeoverpointtype)
	}
}

func (changeoverpointtype *ChangeOverPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChangeOverPointTypeUpdateCallback != nil {
		var frontChangeOverPointType *ChangeOverPointType
		if front != nil {
			frontChangeOverPointType, _ = front.(*ChangeOverPointType)
		}
		stage.OnAfterChangeOverPointTypeUpdateCallback.OnAfterUpdate(stage, changeoverpointtype, frontChangeOverPointType)
	}
}

func (changeoverpointtype *ChangeOverPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChangeOverPointTypeDeleteCallback != nil {
		var frontChangeOverPointType *ChangeOverPointType
		if front != nil {
			frontChangeOverPointType, _ = front.(*ChangeOverPointType)
		}
		stage.OnAfterChangeOverPointTypeDeleteCallback.OnAfterDelete(stage, changeoverpointtype, frontChangeOverPointType)
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

func (directflightclasstype *DirectFlightClassType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectFlightClassTypeCreateCallback != nil {
		stage.OnAfterDirectFlightClassTypeCreateCallback.OnAfterCreate(stage, directflightclasstype)
	}
}

func (directflightclasstype *DirectFlightClassType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectFlightClassTypeUpdateCallback != nil {
		var frontDirectFlightClassType *DirectFlightClassType
		if front != nil {
			frontDirectFlightClassType, _ = front.(*DirectFlightClassType)
		}
		stage.OnAfterDirectFlightClassTypeUpdateCallback.OnAfterUpdate(stage, directflightclasstype, frontDirectFlightClassType)
	}
}

func (directflightclasstype *DirectFlightClassType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectFlightClassTypeDeleteCallback != nil {
		var frontDirectFlightClassType *DirectFlightClassType
		if front != nil {
			frontDirectFlightClassType, _ = front.(*DirectFlightClassType)
		}
		stage.OnAfterDirectFlightClassTypeDeleteCallback.OnAfterDelete(stage, directflightclasstype, frontDirectFlightClassType)
	}
}

func (directflightpropertytype *DirectFlightPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectFlightPropertyTypeCreateCallback != nil {
		stage.OnAfterDirectFlightPropertyTypeCreateCallback.OnAfterCreate(stage, directflightpropertytype)
	}
}

func (directflightpropertytype *DirectFlightPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectFlightPropertyTypeUpdateCallback != nil {
		var frontDirectFlightPropertyType *DirectFlightPropertyType
		if front != nil {
			frontDirectFlightPropertyType, _ = front.(*DirectFlightPropertyType)
		}
		stage.OnAfterDirectFlightPropertyTypeUpdateCallback.OnAfterUpdate(stage, directflightpropertytype, frontDirectFlightPropertyType)
	}
}

func (directflightpropertytype *DirectFlightPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectFlightPropertyTypeDeleteCallback != nil {
		var frontDirectFlightPropertyType *DirectFlightPropertyType
		if front != nil {
			frontDirectFlightPropertyType, _ = front.(*DirectFlightPropertyType)
		}
		stage.OnAfterDirectFlightPropertyTypeDeleteCallback.OnAfterDelete(stage, directflightpropertytype, frontDirectFlightPropertyType)
	}
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectFlightSegmentPropertyTypeCreateCallback != nil {
		stage.OnAfterDirectFlightSegmentPropertyTypeCreateCallback.OnAfterCreate(stage, directflightsegmentpropertytype)
	}
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectFlightSegmentPropertyTypeUpdateCallback != nil {
		var frontDirectFlightSegmentPropertyType *DirectFlightSegmentPropertyType
		if front != nil {
			frontDirectFlightSegmentPropertyType, _ = front.(*DirectFlightSegmentPropertyType)
		}
		stage.OnAfterDirectFlightSegmentPropertyTypeUpdateCallback.OnAfterUpdate(stage, directflightsegmentpropertytype, frontDirectFlightSegmentPropertyType)
	}
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectFlightSegmentPropertyTypeDeleteCallback != nil {
		var frontDirectFlightSegmentPropertyType *DirectFlightSegmentPropertyType
		if front != nil {
			frontDirectFlightSegmentPropertyType, _ = front.(*DirectFlightSegmentPropertyType)
		}
		stage.OnAfterDirectFlightSegmentPropertyTypeDeleteCallback.OnAfterDelete(stage, directflightsegmentpropertytype, frontDirectFlightSegmentPropertyType)
	}
}

func (directflightsegmenttype *DirectFlightSegmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectFlightSegmentTypeCreateCallback != nil {
		stage.OnAfterDirectFlightSegmentTypeCreateCallback.OnAfterCreate(stage, directflightsegmenttype)
	}
}

func (directflightsegmenttype *DirectFlightSegmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectFlightSegmentTypeUpdateCallback != nil {
		var frontDirectFlightSegmentType *DirectFlightSegmentType
		if front != nil {
			frontDirectFlightSegmentType, _ = front.(*DirectFlightSegmentType)
		}
		stage.OnAfterDirectFlightSegmentTypeUpdateCallback.OnAfterUpdate(stage, directflightsegmenttype, frontDirectFlightSegmentType)
	}
}

func (directflightsegmenttype *DirectFlightSegmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectFlightSegmentTypeDeleteCallback != nil {
		var frontDirectFlightSegmentType *DirectFlightSegmentType
		if front != nil {
			frontDirectFlightSegmentType, _ = front.(*DirectFlightSegmentType)
		}
		stage.OnAfterDirectFlightSegmentTypeDeleteCallback.OnAfterDelete(stage, directflightsegmenttype, frontDirectFlightSegmentType)
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

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEnRouteSegmentPointPropertyTypeCreateCallback != nil {
		stage.OnAfterEnRouteSegmentPointPropertyTypeCreateCallback.OnAfterCreate(stage, enroutesegmentpointpropertytype)
	}
}

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnRouteSegmentPointPropertyTypeUpdateCallback != nil {
		var frontEnRouteSegmentPointPropertyType *EnRouteSegmentPointPropertyType
		if front != nil {
			frontEnRouteSegmentPointPropertyType, _ = front.(*EnRouteSegmentPointPropertyType)
		}
		stage.OnAfterEnRouteSegmentPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, enroutesegmentpointpropertytype, frontEnRouteSegmentPointPropertyType)
	}
}

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnRouteSegmentPointPropertyTypeDeleteCallback != nil {
		var frontEnRouteSegmentPointPropertyType *EnRouteSegmentPointPropertyType
		if front != nil {
			frontEnRouteSegmentPointPropertyType, _ = front.(*EnRouteSegmentPointPropertyType)
		}
		stage.OnAfterEnRouteSegmentPointPropertyTypeDeleteCallback.OnAfterDelete(stage, enroutesegmentpointpropertytype, frontEnRouteSegmentPointPropertyType)
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

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightConditionCircumstancePropertyTypeCreateCallback != nil {
		stage.OnAfterFlightConditionCircumstancePropertyTypeCreateCallback.OnAfterCreate(stage, flightconditioncircumstancepropertytype)
	}
}

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionCircumstancePropertyTypeUpdateCallback != nil {
		var frontFlightConditionCircumstancePropertyType *FlightConditionCircumstancePropertyType
		if front != nil {
			frontFlightConditionCircumstancePropertyType, _ = front.(*FlightConditionCircumstancePropertyType)
		}
		stage.OnAfterFlightConditionCircumstancePropertyTypeUpdateCallback.OnAfterUpdate(stage, flightconditioncircumstancepropertytype, frontFlightConditionCircumstancePropertyType)
	}
}

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionCircumstancePropertyTypeDeleteCallback != nil {
		var frontFlightConditionCircumstancePropertyType *FlightConditionCircumstancePropertyType
		if front != nil {
			frontFlightConditionCircumstancePropertyType, _ = front.(*FlightConditionCircumstancePropertyType)
		}
		stage.OnAfterFlightConditionCircumstancePropertyTypeDeleteCallback.OnAfterDelete(stage, flightconditioncircumstancepropertytype, frontFlightConditionCircumstancePropertyType)
	}
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightConditionCircumstanceTypeCreateCallback != nil {
		stage.OnAfterFlightConditionCircumstanceTypeCreateCallback.OnAfterCreate(stage, flightconditioncircumstancetype)
	}
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionCircumstanceTypeUpdateCallback != nil {
		var frontFlightConditionCircumstanceType *FlightConditionCircumstanceType
		if front != nil {
			frontFlightConditionCircumstanceType, _ = front.(*FlightConditionCircumstanceType)
		}
		stage.OnAfterFlightConditionCircumstanceTypeUpdateCallback.OnAfterUpdate(stage, flightconditioncircumstancetype, frontFlightConditionCircumstanceType)
	}
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionCircumstanceTypeDeleteCallback != nil {
		var frontFlightConditionCircumstanceType *FlightConditionCircumstanceType
		if front != nil {
			frontFlightConditionCircumstanceType, _ = front.(*FlightConditionCircumstanceType)
		}
		stage.OnAfterFlightConditionCircumstanceTypeDeleteCallback.OnAfterDelete(stage, flightconditioncircumstancetype, frontFlightConditionCircumstanceType)
	}
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightConditionCombinationPropertyTypeCreateCallback != nil {
		stage.OnAfterFlightConditionCombinationPropertyTypeCreateCallback.OnAfterCreate(stage, flightconditioncombinationpropertytype)
	}
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionCombinationPropertyTypeUpdateCallback != nil {
		var frontFlightConditionCombinationPropertyType *FlightConditionCombinationPropertyType
		if front != nil {
			frontFlightConditionCombinationPropertyType, _ = front.(*FlightConditionCombinationPropertyType)
		}
		stage.OnAfterFlightConditionCombinationPropertyTypeUpdateCallback.OnAfterUpdate(stage, flightconditioncombinationpropertytype, frontFlightConditionCombinationPropertyType)
	}
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionCombinationPropertyTypeDeleteCallback != nil {
		var frontFlightConditionCombinationPropertyType *FlightConditionCombinationPropertyType
		if front != nil {
			frontFlightConditionCombinationPropertyType, _ = front.(*FlightConditionCombinationPropertyType)
		}
		stage.OnAfterFlightConditionCombinationPropertyTypeDeleteCallback.OnAfterDelete(stage, flightconditioncombinationpropertytype, frontFlightConditionCombinationPropertyType)
	}
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightConditionCombinationTypeCreateCallback != nil {
		stage.OnAfterFlightConditionCombinationTypeCreateCallback.OnAfterCreate(stage, flightconditioncombinationtype)
	}
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionCombinationTypeUpdateCallback != nil {
		var frontFlightConditionCombinationType *FlightConditionCombinationType
		if front != nil {
			frontFlightConditionCombinationType, _ = front.(*FlightConditionCombinationType)
		}
		stage.OnAfterFlightConditionCombinationTypeUpdateCallback.OnAfterUpdate(stage, flightconditioncombinationtype, frontFlightConditionCombinationType)
	}
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionCombinationTypeDeleteCallback != nil {
		var frontFlightConditionCombinationType *FlightConditionCombinationType
		if front != nil {
			frontFlightConditionCombinationType, _ = front.(*FlightConditionCombinationType)
		}
		stage.OnAfterFlightConditionCombinationTypeDeleteCallback.OnAfterDelete(stage, flightconditioncombinationtype, frontFlightConditionCombinationType)
	}
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightConditionElementPropertyTypeCreateCallback != nil {
		stage.OnAfterFlightConditionElementPropertyTypeCreateCallback.OnAfterCreate(stage, flightconditionelementpropertytype)
	}
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionElementPropertyTypeUpdateCallback != nil {
		var frontFlightConditionElementPropertyType *FlightConditionElementPropertyType
		if front != nil {
			frontFlightConditionElementPropertyType, _ = front.(*FlightConditionElementPropertyType)
		}
		stage.OnAfterFlightConditionElementPropertyTypeUpdateCallback.OnAfterUpdate(stage, flightconditionelementpropertytype, frontFlightConditionElementPropertyType)
	}
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionElementPropertyTypeDeleteCallback != nil {
		var frontFlightConditionElementPropertyType *FlightConditionElementPropertyType
		if front != nil {
			frontFlightConditionElementPropertyType, _ = front.(*FlightConditionElementPropertyType)
		}
		stage.OnAfterFlightConditionElementPropertyTypeDeleteCallback.OnAfterDelete(stage, flightconditionelementpropertytype, frontFlightConditionElementPropertyType)
	}
}

func (flightconditionelementtype *FlightConditionElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightConditionElementTypeCreateCallback != nil {
		stage.OnAfterFlightConditionElementTypeCreateCallback.OnAfterCreate(stage, flightconditionelementtype)
	}
}

func (flightconditionelementtype *FlightConditionElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionElementTypeUpdateCallback != nil {
		var frontFlightConditionElementType *FlightConditionElementType
		if front != nil {
			frontFlightConditionElementType, _ = front.(*FlightConditionElementType)
		}
		stage.OnAfterFlightConditionElementTypeUpdateCallback.OnAfterUpdate(stage, flightconditionelementtype, frontFlightConditionElementType)
	}
}

func (flightconditionelementtype *FlightConditionElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightConditionElementTypeDeleteCallback != nil {
		var frontFlightConditionElementType *FlightConditionElementType
		if front != nil {
			frontFlightConditionElementType, _ = front.(*FlightConditionElementType)
		}
		stage.OnAfterFlightConditionElementTypeDeleteCallback.OnAfterDelete(stage, flightconditionelementtype, frontFlightConditionElementType)
	}
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightRestrictionLevelPropertyTypeCreateCallback != nil {
		stage.OnAfterFlightRestrictionLevelPropertyTypeCreateCallback.OnAfterCreate(stage, flightrestrictionlevelpropertytype)
	}
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionLevelPropertyTypeUpdateCallback != nil {
		var frontFlightRestrictionLevelPropertyType *FlightRestrictionLevelPropertyType
		if front != nil {
			frontFlightRestrictionLevelPropertyType, _ = front.(*FlightRestrictionLevelPropertyType)
		}
		stage.OnAfterFlightRestrictionLevelPropertyTypeUpdateCallback.OnAfterUpdate(stage, flightrestrictionlevelpropertytype, frontFlightRestrictionLevelPropertyType)
	}
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionLevelPropertyTypeDeleteCallback != nil {
		var frontFlightRestrictionLevelPropertyType *FlightRestrictionLevelPropertyType
		if front != nil {
			frontFlightRestrictionLevelPropertyType, _ = front.(*FlightRestrictionLevelPropertyType)
		}
		stage.OnAfterFlightRestrictionLevelPropertyTypeDeleteCallback.OnAfterDelete(stage, flightrestrictionlevelpropertytype, frontFlightRestrictionLevelPropertyType)
	}
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightRestrictionLevelTypeCreateCallback != nil {
		stage.OnAfterFlightRestrictionLevelTypeCreateCallback.OnAfterCreate(stage, flightrestrictionleveltype)
	}
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionLevelTypeUpdateCallback != nil {
		var frontFlightRestrictionLevelType *FlightRestrictionLevelType
		if front != nil {
			frontFlightRestrictionLevelType, _ = front.(*FlightRestrictionLevelType)
		}
		stage.OnAfterFlightRestrictionLevelTypeUpdateCallback.OnAfterUpdate(stage, flightrestrictionleveltype, frontFlightRestrictionLevelType)
	}
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionLevelTypeDeleteCallback != nil {
		var frontFlightRestrictionLevelType *FlightRestrictionLevelType
		if front != nil {
			frontFlightRestrictionLevelType, _ = front.(*FlightRestrictionLevelType)
		}
		stage.OnAfterFlightRestrictionLevelTypeDeleteCallback.OnAfterDelete(stage, flightrestrictionleveltype, frontFlightRestrictionLevelType)
	}
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightRestrictionRoutePropertyTypeCreateCallback != nil {
		stage.OnAfterFlightRestrictionRoutePropertyTypeCreateCallback.OnAfterCreate(stage, flightrestrictionroutepropertytype)
	}
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionRoutePropertyTypeUpdateCallback != nil {
		var frontFlightRestrictionRoutePropertyType *FlightRestrictionRoutePropertyType
		if front != nil {
			frontFlightRestrictionRoutePropertyType, _ = front.(*FlightRestrictionRoutePropertyType)
		}
		stage.OnAfterFlightRestrictionRoutePropertyTypeUpdateCallback.OnAfterUpdate(stage, flightrestrictionroutepropertytype, frontFlightRestrictionRoutePropertyType)
	}
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionRoutePropertyTypeDeleteCallback != nil {
		var frontFlightRestrictionRoutePropertyType *FlightRestrictionRoutePropertyType
		if front != nil {
			frontFlightRestrictionRoutePropertyType, _ = front.(*FlightRestrictionRoutePropertyType)
		}
		stage.OnAfterFlightRestrictionRoutePropertyTypeDeleteCallback.OnAfterDelete(stage, flightrestrictionroutepropertytype, frontFlightRestrictionRoutePropertyType)
	}
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightRestrictionRouteTypeCreateCallback != nil {
		stage.OnAfterFlightRestrictionRouteTypeCreateCallback.OnAfterCreate(stage, flightrestrictionroutetype)
	}
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionRouteTypeUpdateCallback != nil {
		var frontFlightRestrictionRouteType *FlightRestrictionRouteType
		if front != nil {
			frontFlightRestrictionRouteType, _ = front.(*FlightRestrictionRouteType)
		}
		stage.OnAfterFlightRestrictionRouteTypeUpdateCallback.OnAfterUpdate(stage, flightrestrictionroutetype, frontFlightRestrictionRouteType)
	}
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionRouteTypeDeleteCallback != nil {
		var frontFlightRestrictionRouteType *FlightRestrictionRouteType
		if front != nil {
			frontFlightRestrictionRouteType, _ = front.(*FlightRestrictionRouteType)
		}
		stage.OnAfterFlightRestrictionRouteTypeDeleteCallback.OnAfterDelete(stage, flightrestrictionroutetype, frontFlightRestrictionRouteType)
	}
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightRestrictionTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterFlightRestrictionTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, flightrestrictiontimeslicepropertytype)
	}
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionTimeSlicePropertyTypeUpdateCallback != nil {
		var frontFlightRestrictionTimeSlicePropertyType *FlightRestrictionTimeSlicePropertyType
		if front != nil {
			frontFlightRestrictionTimeSlicePropertyType, _ = front.(*FlightRestrictionTimeSlicePropertyType)
		}
		stage.OnAfterFlightRestrictionTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, flightrestrictiontimeslicepropertytype, frontFlightRestrictionTimeSlicePropertyType)
	}
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionTimeSlicePropertyTypeDeleteCallback != nil {
		var frontFlightRestrictionTimeSlicePropertyType *FlightRestrictionTimeSlicePropertyType
		if front != nil {
			frontFlightRestrictionTimeSlicePropertyType, _ = front.(*FlightRestrictionTimeSlicePropertyType)
		}
		stage.OnAfterFlightRestrictionTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, flightrestrictiontimeslicepropertytype, frontFlightRestrictionTimeSlicePropertyType)
	}
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightRestrictionTimeSliceTypeCreateCallback != nil {
		stage.OnAfterFlightRestrictionTimeSliceTypeCreateCallback.OnAfterCreate(stage, flightrestrictiontimeslicetype)
	}
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionTimeSliceTypeUpdateCallback != nil {
		var frontFlightRestrictionTimeSliceType *FlightRestrictionTimeSliceType
		if front != nil {
			frontFlightRestrictionTimeSliceType, _ = front.(*FlightRestrictionTimeSliceType)
		}
		stage.OnAfterFlightRestrictionTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, flightrestrictiontimeslicetype, frontFlightRestrictionTimeSliceType)
	}
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionTimeSliceTypeDeleteCallback != nil {
		var frontFlightRestrictionTimeSliceType *FlightRestrictionTimeSliceType
		if front != nil {
			frontFlightRestrictionTimeSliceType, _ = front.(*FlightRestrictionTimeSliceType)
		}
		stage.OnAfterFlightRestrictionTimeSliceTypeDeleteCallback.OnAfterDelete(stage, flightrestrictiontimeslicetype, frontFlightRestrictionTimeSliceType)
	}
}

func (flightrestrictiontype *FlightRestrictionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightRestrictionTypeCreateCallback != nil {
		stage.OnAfterFlightRestrictionTypeCreateCallback.OnAfterCreate(stage, flightrestrictiontype)
	}
}

func (flightrestrictiontype *FlightRestrictionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionTypeUpdateCallback != nil {
		var frontFlightRestrictionType *FlightRestrictionType
		if front != nil {
			frontFlightRestrictionType, _ = front.(*FlightRestrictionType)
		}
		stage.OnAfterFlightRestrictionTypeUpdateCallback.OnAfterUpdate(stage, flightrestrictiontype, frontFlightRestrictionType)
	}
}

func (flightrestrictiontype *FlightRestrictionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRestrictionTypeDeleteCallback != nil {
		var frontFlightRestrictionType *FlightRestrictionType
		if front != nil {
			frontFlightRestrictionType, _ = front.(*FlightRestrictionType)
		}
		stage.OnAfterFlightRestrictionTypeDeleteCallback.OnAfterDelete(stage, flightrestrictiontype, frontFlightRestrictionType)
	}
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightRoutingElementPropertyTypeCreateCallback != nil {
		stage.OnAfterFlightRoutingElementPropertyTypeCreateCallback.OnAfterCreate(stage, flightroutingelementpropertytype)
	}
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRoutingElementPropertyTypeUpdateCallback != nil {
		var frontFlightRoutingElementPropertyType *FlightRoutingElementPropertyType
		if front != nil {
			frontFlightRoutingElementPropertyType, _ = front.(*FlightRoutingElementPropertyType)
		}
		stage.OnAfterFlightRoutingElementPropertyTypeUpdateCallback.OnAfterUpdate(stage, flightroutingelementpropertytype, frontFlightRoutingElementPropertyType)
	}
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRoutingElementPropertyTypeDeleteCallback != nil {
		var frontFlightRoutingElementPropertyType *FlightRoutingElementPropertyType
		if front != nil {
			frontFlightRoutingElementPropertyType, _ = front.(*FlightRoutingElementPropertyType)
		}
		stage.OnAfterFlightRoutingElementPropertyTypeDeleteCallback.OnAfterDelete(stage, flightroutingelementpropertytype, frontFlightRoutingElementPropertyType)
	}
}

func (flightroutingelementtype *FlightRoutingElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFlightRoutingElementTypeCreateCallback != nil {
		stage.OnAfterFlightRoutingElementTypeCreateCallback.OnAfterCreate(stage, flightroutingelementtype)
	}
}

func (flightroutingelementtype *FlightRoutingElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRoutingElementTypeUpdateCallback != nil {
		var frontFlightRoutingElementType *FlightRoutingElementType
		if front != nil {
			frontFlightRoutingElementType, _ = front.(*FlightRoutingElementType)
		}
		stage.OnAfterFlightRoutingElementTypeUpdateCallback.OnAfterUpdate(stage, flightroutingelementtype, frontFlightRoutingElementType)
	}
}

func (flightroutingelementtype *FlightRoutingElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFlightRoutingElementTypeDeleteCallback != nil {
		var frontFlightRoutingElementType *FlightRoutingElementType
		if front != nil {
			frontFlightRoutingElementType, _ = front.(*FlightRoutingElementType)
		}
		stage.OnAfterFlightRoutingElementTypeDeleteCallback.OnAfterDelete(stage, flightroutingelementtype, frontFlightRoutingElementType)
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

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteAvailabilityPropertyTypeCreateCallback != nil {
		stage.OnAfterRouteAvailabilityPropertyTypeCreateCallback.OnAfterCreate(stage, routeavailabilitypropertytype)
	}
}

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteAvailabilityPropertyTypeUpdateCallback != nil {
		var frontRouteAvailabilityPropertyType *RouteAvailabilityPropertyType
		if front != nil {
			frontRouteAvailabilityPropertyType, _ = front.(*RouteAvailabilityPropertyType)
		}
		stage.OnAfterRouteAvailabilityPropertyTypeUpdateCallback.OnAfterUpdate(stage, routeavailabilitypropertytype, frontRouteAvailabilityPropertyType)
	}
}

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteAvailabilityPropertyTypeDeleteCallback != nil {
		var frontRouteAvailabilityPropertyType *RouteAvailabilityPropertyType
		if front != nil {
			frontRouteAvailabilityPropertyType, _ = front.(*RouteAvailabilityPropertyType)
		}
		stage.OnAfterRouteAvailabilityPropertyTypeDeleteCallback.OnAfterDelete(stage, routeavailabilitypropertytype, frontRouteAvailabilityPropertyType)
	}
}

func (routeavailabilitytype *RouteAvailabilityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteAvailabilityTypeCreateCallback != nil {
		stage.OnAfterRouteAvailabilityTypeCreateCallback.OnAfterCreate(stage, routeavailabilitytype)
	}
}

func (routeavailabilitytype *RouteAvailabilityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteAvailabilityTypeUpdateCallback != nil {
		var frontRouteAvailabilityType *RouteAvailabilityType
		if front != nil {
			frontRouteAvailabilityType, _ = front.(*RouteAvailabilityType)
		}
		stage.OnAfterRouteAvailabilityTypeUpdateCallback.OnAfterUpdate(stage, routeavailabilitytype, frontRouteAvailabilityType)
	}
}

func (routeavailabilitytype *RouteAvailabilityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteAvailabilityTypeDeleteCallback != nil {
		var frontRouteAvailabilityType *RouteAvailabilityType
		if front != nil {
			frontRouteAvailabilityType, _ = front.(*RouteAvailabilityType)
		}
		stage.OnAfterRouteAvailabilityTypeDeleteCallback.OnAfterDelete(stage, routeavailabilitytype, frontRouteAvailabilityType)
	}
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteDMETimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRouteDMETimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, routedmetimeslicepropertytype)
	}
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteDMETimeSlicePropertyTypeUpdateCallback != nil {
		var frontRouteDMETimeSlicePropertyType *RouteDMETimeSlicePropertyType
		if front != nil {
			frontRouteDMETimeSlicePropertyType, _ = front.(*RouteDMETimeSlicePropertyType)
		}
		stage.OnAfterRouteDMETimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, routedmetimeslicepropertytype, frontRouteDMETimeSlicePropertyType)
	}
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteDMETimeSlicePropertyTypeDeleteCallback != nil {
		var frontRouteDMETimeSlicePropertyType *RouteDMETimeSlicePropertyType
		if front != nil {
			frontRouteDMETimeSlicePropertyType, _ = front.(*RouteDMETimeSlicePropertyType)
		}
		stage.OnAfterRouteDMETimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, routedmetimeslicepropertytype, frontRouteDMETimeSlicePropertyType)
	}
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteDMETimeSliceTypeCreateCallback != nil {
		stage.OnAfterRouteDMETimeSliceTypeCreateCallback.OnAfterCreate(stage, routedmetimeslicetype)
	}
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteDMETimeSliceTypeUpdateCallback != nil {
		var frontRouteDMETimeSliceType *RouteDMETimeSliceType
		if front != nil {
			frontRouteDMETimeSliceType, _ = front.(*RouteDMETimeSliceType)
		}
		stage.OnAfterRouteDMETimeSliceTypeUpdateCallback.OnAfterUpdate(stage, routedmetimeslicetype, frontRouteDMETimeSliceType)
	}
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteDMETimeSliceTypeDeleteCallback != nil {
		var frontRouteDMETimeSliceType *RouteDMETimeSliceType
		if front != nil {
			frontRouteDMETimeSliceType, _ = front.(*RouteDMETimeSliceType)
		}
		stage.OnAfterRouteDMETimeSliceTypeDeleteCallback.OnAfterDelete(stage, routedmetimeslicetype, frontRouteDMETimeSliceType)
	}
}

func (routedmetype *RouteDMEType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteDMETypeCreateCallback != nil {
		stage.OnAfterRouteDMETypeCreateCallback.OnAfterCreate(stage, routedmetype)
	}
}

func (routedmetype *RouteDMEType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteDMETypeUpdateCallback != nil {
		var frontRouteDMEType *RouteDMEType
		if front != nil {
			frontRouteDMEType, _ = front.(*RouteDMEType)
		}
		stage.OnAfterRouteDMETypeUpdateCallback.OnAfterUpdate(stage, routedmetype, frontRouteDMEType)
	}
}

func (routedmetype *RouteDMEType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteDMETypeDeleteCallback != nil {
		var frontRouteDMEType *RouteDMEType
		if front != nil {
			frontRouteDMEType, _ = front.(*RouteDMEType)
		}
		stage.OnAfterRouteDMETypeDeleteCallback.OnAfterDelete(stage, routedmetype, frontRouteDMEType)
	}
}

func (routeportionpropertytype *RoutePortionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRoutePortionPropertyTypeCreateCallback != nil {
		stage.OnAfterRoutePortionPropertyTypeCreateCallback.OnAfterCreate(stage, routeportionpropertytype)
	}
}

func (routeportionpropertytype *RoutePortionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoutePortionPropertyTypeUpdateCallback != nil {
		var frontRoutePortionPropertyType *RoutePortionPropertyType
		if front != nil {
			frontRoutePortionPropertyType, _ = front.(*RoutePortionPropertyType)
		}
		stage.OnAfterRoutePortionPropertyTypeUpdateCallback.OnAfterUpdate(stage, routeportionpropertytype, frontRoutePortionPropertyType)
	}
}

func (routeportionpropertytype *RoutePortionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoutePortionPropertyTypeDeleteCallback != nil {
		var frontRoutePortionPropertyType *RoutePortionPropertyType
		if front != nil {
			frontRoutePortionPropertyType, _ = front.(*RoutePortionPropertyType)
		}
		stage.OnAfterRoutePortionPropertyTypeDeleteCallback.OnAfterDelete(stage, routeportionpropertytype, frontRoutePortionPropertyType)
	}
}

func (routeportiontype *RoutePortionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRoutePortionTypeCreateCallback != nil {
		stage.OnAfterRoutePortionTypeCreateCallback.OnAfterCreate(stage, routeportiontype)
	}
}

func (routeportiontype *RoutePortionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoutePortionTypeUpdateCallback != nil {
		var frontRoutePortionType *RoutePortionType
		if front != nil {
			frontRoutePortionType, _ = front.(*RoutePortionType)
		}
		stage.OnAfterRoutePortionTypeUpdateCallback.OnAfterUpdate(stage, routeportiontype, frontRoutePortionType)
	}
}

func (routeportiontype *RoutePortionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoutePortionTypeDeleteCallback != nil {
		var frontRoutePortionType *RoutePortionType
		if front != nil {
			frontRoutePortionType, _ = front.(*RoutePortionType)
		}
		stage.OnAfterRoutePortionTypeDeleteCallback.OnAfterDelete(stage, routeportiontype, frontRoutePortionType)
	}
}

func (routepropertytype *RoutePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRoutePropertyTypeCreateCallback != nil {
		stage.OnAfterRoutePropertyTypeCreateCallback.OnAfterCreate(stage, routepropertytype)
	}
}

func (routepropertytype *RoutePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoutePropertyTypeUpdateCallback != nil {
		var frontRoutePropertyType *RoutePropertyType
		if front != nil {
			frontRoutePropertyType, _ = front.(*RoutePropertyType)
		}
		stage.OnAfterRoutePropertyTypeUpdateCallback.OnAfterUpdate(stage, routepropertytype, frontRoutePropertyType)
	}
}

func (routepropertytype *RoutePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoutePropertyTypeDeleteCallback != nil {
		var frontRoutePropertyType *RoutePropertyType
		if front != nil {
			frontRoutePropertyType, _ = front.(*RoutePropertyType)
		}
		stage.OnAfterRoutePropertyTypeDeleteCallback.OnAfterDelete(stage, routepropertytype, frontRoutePropertyType)
	}
}

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteSegmentTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRouteSegmentTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, routesegmenttimeslicepropertytype)
	}
}

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteSegmentTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRouteSegmentTimeSlicePropertyType *RouteSegmentTimeSlicePropertyType
		if front != nil {
			frontRouteSegmentTimeSlicePropertyType, _ = front.(*RouteSegmentTimeSlicePropertyType)
		}
		stage.OnAfterRouteSegmentTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, routesegmenttimeslicepropertytype, frontRouteSegmentTimeSlicePropertyType)
	}
}

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteSegmentTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRouteSegmentTimeSlicePropertyType *RouteSegmentTimeSlicePropertyType
		if front != nil {
			frontRouteSegmentTimeSlicePropertyType, _ = front.(*RouteSegmentTimeSlicePropertyType)
		}
		stage.OnAfterRouteSegmentTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, routesegmenttimeslicepropertytype, frontRouteSegmentTimeSlicePropertyType)
	}
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteSegmentTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRouteSegmentTimeSliceTypeCreateCallback.OnAfterCreate(stage, routesegmenttimeslicetype)
	}
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteSegmentTimeSliceTypeUpdateCallback != nil {
		var frontRouteSegmentTimeSliceType *RouteSegmentTimeSliceType
		if front != nil {
			frontRouteSegmentTimeSliceType, _ = front.(*RouteSegmentTimeSliceType)
		}
		stage.OnAfterRouteSegmentTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, routesegmenttimeslicetype, frontRouteSegmentTimeSliceType)
	}
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteSegmentTimeSliceTypeDeleteCallback != nil {
		var frontRouteSegmentTimeSliceType *RouteSegmentTimeSliceType
		if front != nil {
			frontRouteSegmentTimeSliceType, _ = front.(*RouteSegmentTimeSliceType)
		}
		stage.OnAfterRouteSegmentTimeSliceTypeDeleteCallback.OnAfterDelete(stage, routesegmenttimeslicetype, frontRouteSegmentTimeSliceType)
	}
}

func (routesegmenttype *RouteSegmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteSegmentTypeCreateCallback != nil {
		stage.OnAfterRouteSegmentTypeCreateCallback.OnAfterCreate(stage, routesegmenttype)
	}
}

func (routesegmenttype *RouteSegmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteSegmentTypeUpdateCallback != nil {
		var frontRouteSegmentType *RouteSegmentType
		if front != nil {
			frontRouteSegmentType, _ = front.(*RouteSegmentType)
		}
		stage.OnAfterRouteSegmentTypeUpdateCallback.OnAfterUpdate(stage, routesegmenttype, frontRouteSegmentType)
	}
}

func (routesegmenttype *RouteSegmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteSegmentTypeDeleteCallback != nil {
		var frontRouteSegmentType *RouteSegmentType
		if front != nil {
			frontRouteSegmentType, _ = front.(*RouteSegmentType)
		}
		stage.OnAfterRouteSegmentTypeDeleteCallback.OnAfterDelete(stage, routesegmenttype, frontRouteSegmentType)
	}
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRouteTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, routetimeslicepropertytype)
	}
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRouteTimeSlicePropertyType *RouteTimeSlicePropertyType
		if front != nil {
			frontRouteTimeSlicePropertyType, _ = front.(*RouteTimeSlicePropertyType)
		}
		stage.OnAfterRouteTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, routetimeslicepropertytype, frontRouteTimeSlicePropertyType)
	}
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRouteTimeSlicePropertyType *RouteTimeSlicePropertyType
		if front != nil {
			frontRouteTimeSlicePropertyType, _ = front.(*RouteTimeSlicePropertyType)
		}
		stage.OnAfterRouteTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, routetimeslicepropertytype, frontRouteTimeSlicePropertyType)
	}
}

func (routetimeslicetype *RouteTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRouteTimeSliceTypeCreateCallback.OnAfterCreate(stage, routetimeslicetype)
	}
}

func (routetimeslicetype *RouteTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteTimeSliceTypeUpdateCallback != nil {
		var frontRouteTimeSliceType *RouteTimeSliceType
		if front != nil {
			frontRouteTimeSliceType, _ = front.(*RouteTimeSliceType)
		}
		stage.OnAfterRouteTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, routetimeslicetype, frontRouteTimeSliceType)
	}
}

func (routetimeslicetype *RouteTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteTimeSliceTypeDeleteCallback != nil {
		var frontRouteTimeSliceType *RouteTimeSliceType
		if front != nil {
			frontRouteTimeSliceType, _ = front.(*RouteTimeSliceType)
		}
		stage.OnAfterRouteTimeSliceTypeDeleteCallback.OnAfterDelete(stage, routetimeslicetype, frontRouteTimeSliceType)
	}
}

func (routetype *RouteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRouteTypeCreateCallback != nil {
		stage.OnAfterRouteTypeCreateCallback.OnAfterCreate(stage, routetype)
	}
}

func (routetype *RouteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteTypeUpdateCallback != nil {
		var frontRouteType *RouteType
		if front != nil {
			frontRouteType, _ = front.(*RouteType)
		}
		stage.OnAfterRouteTypeUpdateCallback.OnAfterUpdate(stage, routetype, frontRouteType)
	}
}

func (routetype *RouteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRouteTypeDeleteCallback != nil {
		var frontRouteType *RouteType
		if front != nil {
			frontRouteType, _ = front.(*RouteType)
		}
		stage.OnAfterRouteTypeDeleteCallback.OnAfterDelete(stage, routetype, frontRouteType)
	}
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRulesProceduresTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRulesProceduresTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, rulesprocedurestimeslicepropertytype)
	}
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRulesProceduresTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRulesProceduresTimeSlicePropertyType *RulesProceduresTimeSlicePropertyType
		if front != nil {
			frontRulesProceduresTimeSlicePropertyType, _ = front.(*RulesProceduresTimeSlicePropertyType)
		}
		stage.OnAfterRulesProceduresTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, rulesprocedurestimeslicepropertytype, frontRulesProceduresTimeSlicePropertyType)
	}
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRulesProceduresTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRulesProceduresTimeSlicePropertyType *RulesProceduresTimeSlicePropertyType
		if front != nil {
			frontRulesProceduresTimeSlicePropertyType, _ = front.(*RulesProceduresTimeSlicePropertyType)
		}
		stage.OnAfterRulesProceduresTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, rulesprocedurestimeslicepropertytype, frontRulesProceduresTimeSlicePropertyType)
	}
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRulesProceduresTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRulesProceduresTimeSliceTypeCreateCallback.OnAfterCreate(stage, rulesprocedurestimeslicetype)
	}
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRulesProceduresTimeSliceTypeUpdateCallback != nil {
		var frontRulesProceduresTimeSliceType *RulesProceduresTimeSliceType
		if front != nil {
			frontRulesProceduresTimeSliceType, _ = front.(*RulesProceduresTimeSliceType)
		}
		stage.OnAfterRulesProceduresTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, rulesprocedurestimeslicetype, frontRulesProceduresTimeSliceType)
	}
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRulesProceduresTimeSliceTypeDeleteCallback != nil {
		var frontRulesProceduresTimeSliceType *RulesProceduresTimeSliceType
		if front != nil {
			frontRulesProceduresTimeSliceType, _ = front.(*RulesProceduresTimeSliceType)
		}
		stage.OnAfterRulesProceduresTimeSliceTypeDeleteCallback.OnAfterDelete(stage, rulesprocedurestimeslicetype, frontRulesProceduresTimeSliceType)
	}
}

func (rulesprocedurestype *RulesProceduresType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRulesProceduresTypeCreateCallback != nil {
		stage.OnAfterRulesProceduresTypeCreateCallback.OnAfterCreate(stage, rulesprocedurestype)
	}
}

func (rulesprocedurestype *RulesProceduresType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRulesProceduresTypeUpdateCallback != nil {
		var frontRulesProceduresType *RulesProceduresType
		if front != nil {
			frontRulesProceduresType, _ = front.(*RulesProceduresType)
		}
		stage.OnAfterRulesProceduresTypeUpdateCallback.OnAfterUpdate(stage, rulesprocedurestype, frontRulesProceduresType)
	}
}

func (rulesprocedurestype *RulesProceduresType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRulesProceduresTypeDeleteCallback != nil {
		var frontRulesProceduresType *RulesProceduresType
		if front != nil {
			frontRulesProceduresType, _ = front.(*RulesProceduresType)
		}
		stage.OnAfterRulesProceduresTypeDeleteCallback.OnAfterDelete(stage, rulesprocedurestype, frontRulesProceduresType)
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

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelColumnPropertyTypeCreateCallback != nil {
		stage.OnAfterStandardLevelColumnPropertyTypeCreateCallback.OnAfterCreate(stage, standardlevelcolumnpropertytype)
	}
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnPropertyTypeUpdateCallback != nil {
		var frontStandardLevelColumnPropertyType *StandardLevelColumnPropertyType
		if front != nil {
			frontStandardLevelColumnPropertyType, _ = front.(*StandardLevelColumnPropertyType)
		}
		stage.OnAfterStandardLevelColumnPropertyTypeUpdateCallback.OnAfterUpdate(stage, standardlevelcolumnpropertytype, frontStandardLevelColumnPropertyType)
	}
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnPropertyTypeDeleteCallback != nil {
		var frontStandardLevelColumnPropertyType *StandardLevelColumnPropertyType
		if front != nil {
			frontStandardLevelColumnPropertyType, _ = front.(*StandardLevelColumnPropertyType)
		}
		stage.OnAfterStandardLevelColumnPropertyTypeDeleteCallback.OnAfterDelete(stage, standardlevelcolumnpropertytype, frontStandardLevelColumnPropertyType)
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

