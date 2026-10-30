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

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAirportGroundServiceTypeCreateCallback != nil {
		stage.OnAfterAbstractAirportGroundServiceTypeCreateCallback.OnAfterCreate(stage, abstractairportgroundservicetype)
	}
}

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAirportGroundServiceTypeUpdateCallback != nil {
		var frontAbstractAirportGroundServiceType *AbstractAirportGroundServiceType
		if front != nil {
			frontAbstractAirportGroundServiceType, _ = front.(*AbstractAirportGroundServiceType)
		}
		stage.OnAfterAbstractAirportGroundServiceTypeUpdateCallback.OnAfterUpdate(stage, abstractairportgroundservicetype, frontAbstractAirportGroundServiceType)
	}
}

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAirportGroundServiceTypeDeleteCallback != nil {
		var frontAbstractAirportGroundServiceType *AbstractAirportGroundServiceType
		if front != nil {
			frontAbstractAirportGroundServiceType, _ = front.(*AbstractAirportGroundServiceType)
		}
		stage.OnAfterAbstractAirportGroundServiceTypeDeleteCallback.OnAfterDelete(stage, abstractairportgroundservicetype, frontAbstractAirportGroundServiceType)
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

func (abstractservicetype *AbstractServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractServiceTypeCreateCallback != nil {
		stage.OnAfterAbstractServiceTypeCreateCallback.OnAfterCreate(stage, abstractservicetype)
	}
}

func (abstractservicetype *AbstractServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractServiceTypeUpdateCallback != nil {
		var frontAbstractServiceType *AbstractServiceType
		if front != nil {
			frontAbstractServiceType, _ = front.(*AbstractServiceType)
		}
		stage.OnAfterAbstractServiceTypeUpdateCallback.OnAfterUpdate(stage, abstractservicetype, frontAbstractServiceType)
	}
}

func (abstractservicetype *AbstractServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractServiceTypeDeleteCallback != nil {
		var frontAbstractServiceType *AbstractServiceType
		if front != nil {
			frontAbstractServiceType, _ = front.(*AbstractServiceType)
		}
		stage.OnAfterAbstractServiceTypeDeleteCallback.OnAfterDelete(stage, abstractservicetype, frontAbstractServiceType)
	}
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTrafficSeparationServiceTypeCreateCallback != nil {
		stage.OnAfterAbstractTrafficSeparationServiceTypeCreateCallback.OnAfterCreate(stage, abstracttrafficseparationservicetype)
	}
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTrafficSeparationServiceTypeUpdateCallback != nil {
		var frontAbstractTrafficSeparationServiceType *AbstractTrafficSeparationServiceType
		if front != nil {
			frontAbstractTrafficSeparationServiceType, _ = front.(*AbstractTrafficSeparationServiceType)
		}
		stage.OnAfterAbstractTrafficSeparationServiceTypeUpdateCallback.OnAfterUpdate(stage, abstracttrafficseparationservicetype, frontAbstractTrafficSeparationServiceType)
	}
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTrafficSeparationServiceTypeDeleteCallback != nil {
		var frontAbstractTrafficSeparationServiceType *AbstractTrafficSeparationServiceType
		if front != nil {
			frontAbstractTrafficSeparationServiceType, _ = front.(*AbstractTrafficSeparationServiceType)
		}
		stage.OnAfterAbstractTrafficSeparationServiceTypeDeleteCallback.OnAfterDelete(stage, abstracttrafficseparationservicetype, frontAbstractTrafficSeparationServiceType)
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

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirTrafficControlServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirTrafficControlServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airtrafficcontrolservicetimeslicepropertytype)
	}
}

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficControlServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirTrafficControlServiceTimeSlicePropertyType *AirTrafficControlServiceTimeSlicePropertyType
		if front != nil {
			frontAirTrafficControlServiceTimeSlicePropertyType, _ = front.(*AirTrafficControlServiceTimeSlicePropertyType)
		}
		stage.OnAfterAirTrafficControlServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airtrafficcontrolservicetimeslicepropertytype, frontAirTrafficControlServiceTimeSlicePropertyType)
	}
}

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficControlServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirTrafficControlServiceTimeSlicePropertyType *AirTrafficControlServiceTimeSlicePropertyType
		if front != nil {
			frontAirTrafficControlServiceTimeSlicePropertyType, _ = front.(*AirTrafficControlServiceTimeSlicePropertyType)
		}
		stage.OnAfterAirTrafficControlServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airtrafficcontrolservicetimeslicepropertytype, frontAirTrafficControlServiceTimeSlicePropertyType)
	}
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirTrafficControlServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirTrafficControlServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, airtrafficcontrolservicetimeslicetype)
	}
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficControlServiceTimeSliceTypeUpdateCallback != nil {
		var frontAirTrafficControlServiceTimeSliceType *AirTrafficControlServiceTimeSliceType
		if front != nil {
			frontAirTrafficControlServiceTimeSliceType, _ = front.(*AirTrafficControlServiceTimeSliceType)
		}
		stage.OnAfterAirTrafficControlServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airtrafficcontrolservicetimeslicetype, frontAirTrafficControlServiceTimeSliceType)
	}
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficControlServiceTimeSliceTypeDeleteCallback != nil {
		var frontAirTrafficControlServiceTimeSliceType *AirTrafficControlServiceTimeSliceType
		if front != nil {
			frontAirTrafficControlServiceTimeSliceType, _ = front.(*AirTrafficControlServiceTimeSliceType)
		}
		stage.OnAfterAirTrafficControlServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airtrafficcontrolservicetimeslicetype, frontAirTrafficControlServiceTimeSliceType)
	}
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirTrafficControlServiceTypeCreateCallback != nil {
		stage.OnAfterAirTrafficControlServiceTypeCreateCallback.OnAfterCreate(stage, airtrafficcontrolservicetype)
	}
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficControlServiceTypeUpdateCallback != nil {
		var frontAirTrafficControlServiceType *AirTrafficControlServiceType
		if front != nil {
			frontAirTrafficControlServiceType, _ = front.(*AirTrafficControlServiceType)
		}
		stage.OnAfterAirTrafficControlServiceTypeUpdateCallback.OnAfterUpdate(stage, airtrafficcontrolservicetype, frontAirTrafficControlServiceType)
	}
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficControlServiceTypeDeleteCallback != nil {
		var frontAirTrafficControlServiceType *AirTrafficControlServiceType
		if front != nil {
			frontAirTrafficControlServiceType, _ = front.(*AirTrafficControlServiceType)
		}
		stage.OnAfterAirTrafficControlServiceTypeDeleteCallback.OnAfterDelete(stage, airtrafficcontrolservicetype, frontAirTrafficControlServiceType)
	}
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirTrafficManagementServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirTrafficManagementServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airtrafficmanagementservicetimeslicepropertytype)
	}
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficManagementServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirTrafficManagementServiceTimeSlicePropertyType *AirTrafficManagementServiceTimeSlicePropertyType
		if front != nil {
			frontAirTrafficManagementServiceTimeSlicePropertyType, _ = front.(*AirTrafficManagementServiceTimeSlicePropertyType)
		}
		stage.OnAfterAirTrafficManagementServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airtrafficmanagementservicetimeslicepropertytype, frontAirTrafficManagementServiceTimeSlicePropertyType)
	}
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficManagementServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirTrafficManagementServiceTimeSlicePropertyType *AirTrafficManagementServiceTimeSlicePropertyType
		if front != nil {
			frontAirTrafficManagementServiceTimeSlicePropertyType, _ = front.(*AirTrafficManagementServiceTimeSlicePropertyType)
		}
		stage.OnAfterAirTrafficManagementServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airtrafficmanagementservicetimeslicepropertytype, frontAirTrafficManagementServiceTimeSlicePropertyType)
	}
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirTrafficManagementServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirTrafficManagementServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, airtrafficmanagementservicetimeslicetype)
	}
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficManagementServiceTimeSliceTypeUpdateCallback != nil {
		var frontAirTrafficManagementServiceTimeSliceType *AirTrafficManagementServiceTimeSliceType
		if front != nil {
			frontAirTrafficManagementServiceTimeSliceType, _ = front.(*AirTrafficManagementServiceTimeSliceType)
		}
		stage.OnAfterAirTrafficManagementServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airtrafficmanagementservicetimeslicetype, frontAirTrafficManagementServiceTimeSliceType)
	}
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficManagementServiceTimeSliceTypeDeleteCallback != nil {
		var frontAirTrafficManagementServiceTimeSliceType *AirTrafficManagementServiceTimeSliceType
		if front != nil {
			frontAirTrafficManagementServiceTimeSliceType, _ = front.(*AirTrafficManagementServiceTimeSliceType)
		}
		stage.OnAfterAirTrafficManagementServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airtrafficmanagementservicetimeslicetype, frontAirTrafficManagementServiceTimeSliceType)
	}
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirTrafficManagementServiceTypeCreateCallback != nil {
		stage.OnAfterAirTrafficManagementServiceTypeCreateCallback.OnAfterCreate(stage, airtrafficmanagementservicetype)
	}
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficManagementServiceTypeUpdateCallback != nil {
		var frontAirTrafficManagementServiceType *AirTrafficManagementServiceType
		if front != nil {
			frontAirTrafficManagementServiceType, _ = front.(*AirTrafficManagementServiceType)
		}
		stage.OnAfterAirTrafficManagementServiceTypeUpdateCallback.OnAfterUpdate(stage, airtrafficmanagementservicetype, frontAirTrafficManagementServiceType)
	}
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirTrafficManagementServiceTypeDeleteCallback != nil {
		var frontAirTrafficManagementServiceType *AirTrafficManagementServiceType
		if front != nil {
			frontAirTrafficManagementServiceType, _ = front.(*AirTrafficManagementServiceType)
		}
		stage.OnAfterAirTrafficManagementServiceTypeDeleteCallback.OnAfterDelete(stage, airtrafficmanagementservicetype, frontAirTrafficManagementServiceType)
	}
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftGroundServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAircraftGroundServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, aircraftgroundservicetimeslicepropertytype)
	}
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftGroundServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAircraftGroundServiceTimeSlicePropertyType *AircraftGroundServiceTimeSlicePropertyType
		if front != nil {
			frontAircraftGroundServiceTimeSlicePropertyType, _ = front.(*AircraftGroundServiceTimeSlicePropertyType)
		}
		stage.OnAfterAircraftGroundServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, aircraftgroundservicetimeslicepropertytype, frontAircraftGroundServiceTimeSlicePropertyType)
	}
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftGroundServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAircraftGroundServiceTimeSlicePropertyType *AircraftGroundServiceTimeSlicePropertyType
		if front != nil {
			frontAircraftGroundServiceTimeSlicePropertyType, _ = front.(*AircraftGroundServiceTimeSlicePropertyType)
		}
		stage.OnAfterAircraftGroundServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, aircraftgroundservicetimeslicepropertytype, frontAircraftGroundServiceTimeSlicePropertyType)
	}
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftGroundServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAircraftGroundServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, aircraftgroundservicetimeslicetype)
	}
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftGroundServiceTimeSliceTypeUpdateCallback != nil {
		var frontAircraftGroundServiceTimeSliceType *AircraftGroundServiceTimeSliceType
		if front != nil {
			frontAircraftGroundServiceTimeSliceType, _ = front.(*AircraftGroundServiceTimeSliceType)
		}
		stage.OnAfterAircraftGroundServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, aircraftgroundservicetimeslicetype, frontAircraftGroundServiceTimeSliceType)
	}
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftGroundServiceTimeSliceTypeDeleteCallback != nil {
		var frontAircraftGroundServiceTimeSliceType *AircraftGroundServiceTimeSliceType
		if front != nil {
			frontAircraftGroundServiceTimeSliceType, _ = front.(*AircraftGroundServiceTimeSliceType)
		}
		stage.OnAfterAircraftGroundServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, aircraftgroundservicetimeslicetype, frontAircraftGroundServiceTimeSliceType)
	}
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftGroundServiceTypeCreateCallback != nil {
		stage.OnAfterAircraftGroundServiceTypeCreateCallback.OnAfterCreate(stage, aircraftgroundservicetype)
	}
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftGroundServiceTypeUpdateCallback != nil {
		var frontAircraftGroundServiceType *AircraftGroundServiceType
		if front != nil {
			frontAircraftGroundServiceType, _ = front.(*AircraftGroundServiceType)
		}
		stage.OnAfterAircraftGroundServiceTypeUpdateCallback.OnAfterUpdate(stage, aircraftgroundservicetype, frontAircraftGroundServiceType)
	}
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftGroundServiceTypeDeleteCallback != nil {
		var frontAircraftGroundServiceType *AircraftGroundServiceType
		if front != nil {
			frontAircraftGroundServiceType, _ = front.(*AircraftGroundServiceType)
		}
		stage.OnAfterAircraftGroundServiceTypeDeleteCallback.OnAfterDelete(stage, aircraftgroundservicetype, frontAircraftGroundServiceType)
	}
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportClearanceServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirportClearanceServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airportclearanceservicetimeslicepropertytype)
	}
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportClearanceServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirportClearanceServiceTimeSlicePropertyType *AirportClearanceServiceTimeSlicePropertyType
		if front != nil {
			frontAirportClearanceServiceTimeSlicePropertyType, _ = front.(*AirportClearanceServiceTimeSlicePropertyType)
		}
		stage.OnAfterAirportClearanceServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airportclearanceservicetimeslicepropertytype, frontAirportClearanceServiceTimeSlicePropertyType)
	}
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportClearanceServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirportClearanceServiceTimeSlicePropertyType *AirportClearanceServiceTimeSlicePropertyType
		if front != nil {
			frontAirportClearanceServiceTimeSlicePropertyType, _ = front.(*AirportClearanceServiceTimeSlicePropertyType)
		}
		stage.OnAfterAirportClearanceServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airportclearanceservicetimeslicepropertytype, frontAirportClearanceServiceTimeSlicePropertyType)
	}
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportClearanceServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirportClearanceServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, airportclearanceservicetimeslicetype)
	}
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportClearanceServiceTimeSliceTypeUpdateCallback != nil {
		var frontAirportClearanceServiceTimeSliceType *AirportClearanceServiceTimeSliceType
		if front != nil {
			frontAirportClearanceServiceTimeSliceType, _ = front.(*AirportClearanceServiceTimeSliceType)
		}
		stage.OnAfterAirportClearanceServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airportclearanceservicetimeslicetype, frontAirportClearanceServiceTimeSliceType)
	}
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportClearanceServiceTimeSliceTypeDeleteCallback != nil {
		var frontAirportClearanceServiceTimeSliceType *AirportClearanceServiceTimeSliceType
		if front != nil {
			frontAirportClearanceServiceTimeSliceType, _ = front.(*AirportClearanceServiceTimeSliceType)
		}
		stage.OnAfterAirportClearanceServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airportclearanceservicetimeslicetype, frontAirportClearanceServiceTimeSliceType)
	}
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportClearanceServiceTypeCreateCallback != nil {
		stage.OnAfterAirportClearanceServiceTypeCreateCallback.OnAfterCreate(stage, airportclearanceservicetype)
	}
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportClearanceServiceTypeUpdateCallback != nil {
		var frontAirportClearanceServiceType *AirportClearanceServiceType
		if front != nil {
			frontAirportClearanceServiceType, _ = front.(*AirportClearanceServiceType)
		}
		stage.OnAfterAirportClearanceServiceTypeUpdateCallback.OnAfterUpdate(stage, airportclearanceservicetype, frontAirportClearanceServiceType)
	}
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportClearanceServiceTypeDeleteCallback != nil {
		var frontAirportClearanceServiceType *AirportClearanceServiceType
		if front != nil {
			frontAirportClearanceServiceType, _ = front.(*AirportClearanceServiceType)
		}
		stage.OnAfterAirportClearanceServiceTypeDeleteCallback.OnAfterDelete(stage, airportclearanceservicetype, frontAirportClearanceServiceType)
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

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportSuppliesServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirportSuppliesServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airportsuppliesservicetimeslicepropertytype)
	}
}

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportSuppliesServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirportSuppliesServiceTimeSlicePropertyType *AirportSuppliesServiceTimeSlicePropertyType
		if front != nil {
			frontAirportSuppliesServiceTimeSlicePropertyType, _ = front.(*AirportSuppliesServiceTimeSlicePropertyType)
		}
		stage.OnAfterAirportSuppliesServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airportsuppliesservicetimeslicepropertytype, frontAirportSuppliesServiceTimeSlicePropertyType)
	}
}

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportSuppliesServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirportSuppliesServiceTimeSlicePropertyType *AirportSuppliesServiceTimeSlicePropertyType
		if front != nil {
			frontAirportSuppliesServiceTimeSlicePropertyType, _ = front.(*AirportSuppliesServiceTimeSlicePropertyType)
		}
		stage.OnAfterAirportSuppliesServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airportsuppliesservicetimeslicepropertytype, frontAirportSuppliesServiceTimeSlicePropertyType)
	}
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportSuppliesServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirportSuppliesServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, airportsuppliesservicetimeslicetype)
	}
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportSuppliesServiceTimeSliceTypeUpdateCallback != nil {
		var frontAirportSuppliesServiceTimeSliceType *AirportSuppliesServiceTimeSliceType
		if front != nil {
			frontAirportSuppliesServiceTimeSliceType, _ = front.(*AirportSuppliesServiceTimeSliceType)
		}
		stage.OnAfterAirportSuppliesServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airportsuppliesservicetimeslicetype, frontAirportSuppliesServiceTimeSliceType)
	}
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportSuppliesServiceTimeSliceTypeDeleteCallback != nil {
		var frontAirportSuppliesServiceTimeSliceType *AirportSuppliesServiceTimeSliceType
		if front != nil {
			frontAirportSuppliesServiceTimeSliceType, _ = front.(*AirportSuppliesServiceTimeSliceType)
		}
		stage.OnAfterAirportSuppliesServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airportsuppliesservicetimeslicetype, frontAirportSuppliesServiceTimeSliceType)
	}
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportSuppliesServiceTypeCreateCallback != nil {
		stage.OnAfterAirportSuppliesServiceTypeCreateCallback.OnAfterCreate(stage, airportsuppliesservicetype)
	}
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportSuppliesServiceTypeUpdateCallback != nil {
		var frontAirportSuppliesServiceType *AirportSuppliesServiceType
		if front != nil {
			frontAirportSuppliesServiceType, _ = front.(*AirportSuppliesServiceType)
		}
		stage.OnAfterAirportSuppliesServiceTypeUpdateCallback.OnAfterUpdate(stage, airportsuppliesservicetype, frontAirportSuppliesServiceType)
	}
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportSuppliesServiceTypeDeleteCallback != nil {
		var frontAirportSuppliesServiceType *AirportSuppliesServiceType
		if front != nil {
			frontAirportSuppliesServiceType, _ = front.(*AirportSuppliesServiceType)
		}
		stage.OnAfterAirportSuppliesServiceTypeDeleteCallback.OnAfterDelete(stage, airportsuppliesservicetype, frontAirportSuppliesServiceType)
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

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCallsignDetailPropertyTypeCreateCallback != nil {
		stage.OnAfterCallsignDetailPropertyTypeCreateCallback.OnAfterCreate(stage, callsigndetailpropertytype)
	}
}

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCallsignDetailPropertyTypeUpdateCallback != nil {
		var frontCallsignDetailPropertyType *CallsignDetailPropertyType
		if front != nil {
			frontCallsignDetailPropertyType, _ = front.(*CallsignDetailPropertyType)
		}
		stage.OnAfterCallsignDetailPropertyTypeUpdateCallback.OnAfterUpdate(stage, callsigndetailpropertytype, frontCallsignDetailPropertyType)
	}
}

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCallsignDetailPropertyTypeDeleteCallback != nil {
		var frontCallsignDetailPropertyType *CallsignDetailPropertyType
		if front != nil {
			frontCallsignDetailPropertyType, _ = front.(*CallsignDetailPropertyType)
		}
		stage.OnAfterCallsignDetailPropertyTypeDeleteCallback.OnAfterDelete(stage, callsigndetailpropertytype, frontCallsignDetailPropertyType)
	}
}

func (callsigndetailtype *CallsignDetailType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCallsignDetailTypeCreateCallback != nil {
		stage.OnAfterCallsignDetailTypeCreateCallback.OnAfterCreate(stage, callsigndetailtype)
	}
}

func (callsigndetailtype *CallsignDetailType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCallsignDetailTypeUpdateCallback != nil {
		var frontCallsignDetailType *CallsignDetailType
		if front != nil {
			frontCallsignDetailType, _ = front.(*CallsignDetailType)
		}
		stage.OnAfterCallsignDetailTypeUpdateCallback.OnAfterUpdate(stage, callsigndetailtype, frontCallsignDetailType)
	}
}

func (callsigndetailtype *CallsignDetailType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCallsignDetailTypeDeleteCallback != nil {
		var frontCallsignDetailType *CallsignDetailType
		if front != nil {
			frontCallsignDetailType, _ = front.(*CallsignDetailType)
		}
		stage.OnAfterCallsignDetailTypeDeleteCallback.OnAfterDelete(stage, callsigndetailtype, frontCallsignDetailType)
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

func (directionfinderpropertytype *DirectionFinderPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionFinderPropertyTypeCreateCallback != nil {
		stage.OnAfterDirectionFinderPropertyTypeCreateCallback.OnAfterCreate(stage, directionfinderpropertytype)
	}
}

func (directionfinderpropertytype *DirectionFinderPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionFinderPropertyTypeUpdateCallback != nil {
		var frontDirectionFinderPropertyType *DirectionFinderPropertyType
		if front != nil {
			frontDirectionFinderPropertyType, _ = front.(*DirectionFinderPropertyType)
		}
		stage.OnAfterDirectionFinderPropertyTypeUpdateCallback.OnAfterUpdate(stage, directionfinderpropertytype, frontDirectionFinderPropertyType)
	}
}

func (directionfinderpropertytype *DirectionFinderPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionFinderPropertyTypeDeleteCallback != nil {
		var frontDirectionFinderPropertyType *DirectionFinderPropertyType
		if front != nil {
			frontDirectionFinderPropertyType, _ = front.(*DirectionFinderPropertyType)
		}
		stage.OnAfterDirectionFinderPropertyTypeDeleteCallback.OnAfterDelete(stage, directionfinderpropertytype, frontDirectionFinderPropertyType)
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

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFireFightingServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterFireFightingServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, firefightingservicetimeslicepropertytype)
	}
}

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFireFightingServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontFireFightingServiceTimeSlicePropertyType *FireFightingServiceTimeSlicePropertyType
		if front != nil {
			frontFireFightingServiceTimeSlicePropertyType, _ = front.(*FireFightingServiceTimeSlicePropertyType)
		}
		stage.OnAfterFireFightingServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, firefightingservicetimeslicepropertytype, frontFireFightingServiceTimeSlicePropertyType)
	}
}

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFireFightingServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontFireFightingServiceTimeSlicePropertyType *FireFightingServiceTimeSlicePropertyType
		if front != nil {
			frontFireFightingServiceTimeSlicePropertyType, _ = front.(*FireFightingServiceTimeSlicePropertyType)
		}
		stage.OnAfterFireFightingServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, firefightingservicetimeslicepropertytype, frontFireFightingServiceTimeSlicePropertyType)
	}
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFireFightingServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterFireFightingServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, firefightingservicetimeslicetype)
	}
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFireFightingServiceTimeSliceTypeUpdateCallback != nil {
		var frontFireFightingServiceTimeSliceType *FireFightingServiceTimeSliceType
		if front != nil {
			frontFireFightingServiceTimeSliceType, _ = front.(*FireFightingServiceTimeSliceType)
		}
		stage.OnAfterFireFightingServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, firefightingservicetimeslicetype, frontFireFightingServiceTimeSliceType)
	}
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFireFightingServiceTimeSliceTypeDeleteCallback != nil {
		var frontFireFightingServiceTimeSliceType *FireFightingServiceTimeSliceType
		if front != nil {
			frontFireFightingServiceTimeSliceType, _ = front.(*FireFightingServiceTimeSliceType)
		}
		stage.OnAfterFireFightingServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, firefightingservicetimeslicetype, frontFireFightingServiceTimeSliceType)
	}
}

func (firefightingservicetype *FireFightingServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFireFightingServiceTypeCreateCallback != nil {
		stage.OnAfterFireFightingServiceTypeCreateCallback.OnAfterCreate(stage, firefightingservicetype)
	}
}

func (firefightingservicetype *FireFightingServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFireFightingServiceTypeUpdateCallback != nil {
		var frontFireFightingServiceType *FireFightingServiceType
		if front != nil {
			frontFireFightingServiceType, _ = front.(*FireFightingServiceType)
		}
		stage.OnAfterFireFightingServiceTypeUpdateCallback.OnAfterUpdate(stage, firefightingservicetype, frontFireFightingServiceType)
	}
}

func (firefightingservicetype *FireFightingServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFireFightingServiceTypeDeleteCallback != nil {
		var frontFireFightingServiceType *FireFightingServiceType
		if front != nil {
			frontFireFightingServiceType, _ = front.(*FireFightingServiceType)
		}
		stage.OnAfterFireFightingServiceTypeDeleteCallback.OnAfterDelete(stage, firefightingservicetype, frontFireFightingServiceType)
	}
}

func (fuelpropertytype *FuelPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFuelPropertyTypeCreateCallback != nil {
		stage.OnAfterFuelPropertyTypeCreateCallback.OnAfterCreate(stage, fuelpropertytype)
	}
}

func (fuelpropertytype *FuelPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFuelPropertyTypeUpdateCallback != nil {
		var frontFuelPropertyType *FuelPropertyType
		if front != nil {
			frontFuelPropertyType, _ = front.(*FuelPropertyType)
		}
		stage.OnAfterFuelPropertyTypeUpdateCallback.OnAfterUpdate(stage, fuelpropertytype, frontFuelPropertyType)
	}
}

func (fuelpropertytype *FuelPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFuelPropertyTypeDeleteCallback != nil {
		var frontFuelPropertyType *FuelPropertyType
		if front != nil {
			frontFuelPropertyType, _ = front.(*FuelPropertyType)
		}
		stage.OnAfterFuelPropertyTypeDeleteCallback.OnAfterDelete(stage, fuelpropertytype, frontFuelPropertyType)
	}
}

func (fueltype *FuelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFuelTypeCreateCallback != nil {
		stage.OnAfterFuelTypeCreateCallback.OnAfterCreate(stage, fueltype)
	}
}

func (fueltype *FuelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFuelTypeUpdateCallback != nil {
		var frontFuelType *FuelType
		if front != nil {
			frontFuelType, _ = front.(*FuelType)
		}
		stage.OnAfterFuelTypeUpdateCallback.OnAfterUpdate(stage, fueltype, frontFuelType)
	}
}

func (fueltype *FuelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFuelTypeDeleteCallback != nil {
		var frontFuelType *FuelType
		if front != nil {
			frontFuelType, _ = front.(*FuelType)
		}
		stage.OnAfterFuelTypeDeleteCallback.OnAfterDelete(stage, fueltype, frontFuelType)
	}
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroundLightSystemPropertyTypeCreateCallback != nil {
		stage.OnAfterGroundLightSystemPropertyTypeCreateCallback.OnAfterCreate(stage, groundlightsystempropertytype)
	}
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundLightSystemPropertyTypeUpdateCallback != nil {
		var frontGroundLightSystemPropertyType *GroundLightSystemPropertyType
		if front != nil {
			frontGroundLightSystemPropertyType, _ = front.(*GroundLightSystemPropertyType)
		}
		stage.OnAfterGroundLightSystemPropertyTypeUpdateCallback.OnAfterUpdate(stage, groundlightsystempropertytype, frontGroundLightSystemPropertyType)
	}
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundLightSystemPropertyTypeDeleteCallback != nil {
		var frontGroundLightSystemPropertyType *GroundLightSystemPropertyType
		if front != nil {
			frontGroundLightSystemPropertyType, _ = front.(*GroundLightSystemPropertyType)
		}
		stage.OnAfterGroundLightSystemPropertyTypeDeleteCallback.OnAfterDelete(stage, groundlightsystempropertytype, frontGroundLightSystemPropertyType)
	}
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroundTrafficControlServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterGroundTrafficControlServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, groundtrafficcontrolservicetimeslicepropertytype)
	}
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundTrafficControlServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontGroundTrafficControlServiceTimeSlicePropertyType *GroundTrafficControlServiceTimeSlicePropertyType
		if front != nil {
			frontGroundTrafficControlServiceTimeSlicePropertyType, _ = front.(*GroundTrafficControlServiceTimeSlicePropertyType)
		}
		stage.OnAfterGroundTrafficControlServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, groundtrafficcontrolservicetimeslicepropertytype, frontGroundTrafficControlServiceTimeSlicePropertyType)
	}
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundTrafficControlServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontGroundTrafficControlServiceTimeSlicePropertyType *GroundTrafficControlServiceTimeSlicePropertyType
		if front != nil {
			frontGroundTrafficControlServiceTimeSlicePropertyType, _ = front.(*GroundTrafficControlServiceTimeSlicePropertyType)
		}
		stage.OnAfterGroundTrafficControlServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, groundtrafficcontrolservicetimeslicepropertytype, frontGroundTrafficControlServiceTimeSlicePropertyType)
	}
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroundTrafficControlServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterGroundTrafficControlServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, groundtrafficcontrolservicetimeslicetype)
	}
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundTrafficControlServiceTimeSliceTypeUpdateCallback != nil {
		var frontGroundTrafficControlServiceTimeSliceType *GroundTrafficControlServiceTimeSliceType
		if front != nil {
			frontGroundTrafficControlServiceTimeSliceType, _ = front.(*GroundTrafficControlServiceTimeSliceType)
		}
		stage.OnAfterGroundTrafficControlServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, groundtrafficcontrolservicetimeslicetype, frontGroundTrafficControlServiceTimeSliceType)
	}
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundTrafficControlServiceTimeSliceTypeDeleteCallback != nil {
		var frontGroundTrafficControlServiceTimeSliceType *GroundTrafficControlServiceTimeSliceType
		if front != nil {
			frontGroundTrafficControlServiceTimeSliceType, _ = front.(*GroundTrafficControlServiceTimeSliceType)
		}
		stage.OnAfterGroundTrafficControlServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, groundtrafficcontrolservicetimeslicetype, frontGroundTrafficControlServiceTimeSliceType)
	}
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroundTrafficControlServiceTypeCreateCallback != nil {
		stage.OnAfterGroundTrafficControlServiceTypeCreateCallback.OnAfterCreate(stage, groundtrafficcontrolservicetype)
	}
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundTrafficControlServiceTypeUpdateCallback != nil {
		var frontGroundTrafficControlServiceType *GroundTrafficControlServiceType
		if front != nil {
			frontGroundTrafficControlServiceType, _ = front.(*GroundTrafficControlServiceType)
		}
		stage.OnAfterGroundTrafficControlServiceTypeUpdateCallback.OnAfterUpdate(stage, groundtrafficcontrolservicetype, frontGroundTrafficControlServiceType)
	}
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundTrafficControlServiceTypeDeleteCallback != nil {
		var frontGroundTrafficControlServiceType *GroundTrafficControlServiceType
		if front != nil {
			frontGroundTrafficControlServiceType, _ = front.(*GroundTrafficControlServiceType)
		}
		stage.OnAfterGroundTrafficControlServiceTypeDeleteCallback.OnAfterDelete(stage, groundtrafficcontrolservicetype, frontGroundTrafficControlServiceType)
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

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInformationServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterInformationServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, informationservicetimeslicepropertytype)
	}
}

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInformationServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontInformationServiceTimeSlicePropertyType *InformationServiceTimeSlicePropertyType
		if front != nil {
			frontInformationServiceTimeSlicePropertyType, _ = front.(*InformationServiceTimeSlicePropertyType)
		}
		stage.OnAfterInformationServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, informationservicetimeslicepropertytype, frontInformationServiceTimeSlicePropertyType)
	}
}

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInformationServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontInformationServiceTimeSlicePropertyType *InformationServiceTimeSlicePropertyType
		if front != nil {
			frontInformationServiceTimeSlicePropertyType, _ = front.(*InformationServiceTimeSlicePropertyType)
		}
		stage.OnAfterInformationServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, informationservicetimeslicepropertytype, frontInformationServiceTimeSlicePropertyType)
	}
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInformationServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterInformationServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, informationservicetimeslicetype)
	}
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInformationServiceTimeSliceTypeUpdateCallback != nil {
		var frontInformationServiceTimeSliceType *InformationServiceTimeSliceType
		if front != nil {
			frontInformationServiceTimeSliceType, _ = front.(*InformationServiceTimeSliceType)
		}
		stage.OnAfterInformationServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, informationservicetimeslicetype, frontInformationServiceTimeSliceType)
	}
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInformationServiceTimeSliceTypeDeleteCallback != nil {
		var frontInformationServiceTimeSliceType *InformationServiceTimeSliceType
		if front != nil {
			frontInformationServiceTimeSliceType, _ = front.(*InformationServiceTimeSliceType)
		}
		stage.OnAfterInformationServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, informationservicetimeslicetype, frontInformationServiceTimeSliceType)
	}
}

func (informationservicetype *InformationServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInformationServiceTypeCreateCallback != nil {
		stage.OnAfterInformationServiceTypeCreateCallback.OnAfterCreate(stage, informationservicetype)
	}
}

func (informationservicetype *InformationServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInformationServiceTypeUpdateCallback != nil {
		var frontInformationServiceType *InformationServiceType
		if front != nil {
			frontInformationServiceType, _ = front.(*InformationServiceType)
		}
		stage.OnAfterInformationServiceTypeUpdateCallback.OnAfterUpdate(stage, informationservicetype, frontInformationServiceType)
	}
}

func (informationservicetype *InformationServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInformationServiceTypeDeleteCallback != nil {
		var frontInformationServiceType *InformationServiceType
		if front != nil {
			frontInformationServiceType, _ = front.(*InformationServiceType)
		}
		stage.OnAfterInformationServiceTypeDeleteCallback.OnAfterDelete(stage, informationservicetype, frontInformationServiceType)
	}
}

func (lightactivationpropertytype *LightActivationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightActivationPropertyTypeCreateCallback != nil {
		stage.OnAfterLightActivationPropertyTypeCreateCallback.OnAfterCreate(stage, lightactivationpropertytype)
	}
}

func (lightactivationpropertytype *LightActivationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightActivationPropertyTypeUpdateCallback != nil {
		var frontLightActivationPropertyType *LightActivationPropertyType
		if front != nil {
			frontLightActivationPropertyType, _ = front.(*LightActivationPropertyType)
		}
		stage.OnAfterLightActivationPropertyTypeUpdateCallback.OnAfterUpdate(stage, lightactivationpropertytype, frontLightActivationPropertyType)
	}
}

func (lightactivationpropertytype *LightActivationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightActivationPropertyTypeDeleteCallback != nil {
		var frontLightActivationPropertyType *LightActivationPropertyType
		if front != nil {
			frontLightActivationPropertyType, _ = front.(*LightActivationPropertyType)
		}
		stage.OnAfterLightActivationPropertyTypeDeleteCallback.OnAfterDelete(stage, lightactivationpropertytype, frontLightActivationPropertyType)
	}
}

func (lightactivationtype *LightActivationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightActivationTypeCreateCallback != nil {
		stage.OnAfterLightActivationTypeCreateCallback.OnAfterCreate(stage, lightactivationtype)
	}
}

func (lightactivationtype *LightActivationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightActivationTypeUpdateCallback != nil {
		var frontLightActivationType *LightActivationType
		if front != nil {
			frontLightActivationType, _ = front.(*LightActivationType)
		}
		stage.OnAfterLightActivationTypeUpdateCallback.OnAfterUpdate(stage, lightactivationtype, frontLightActivationType)
	}
}

func (lightactivationtype *LightActivationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightActivationTypeDeleteCallback != nil {
		var frontLightActivationType *LightActivationType
		if front != nil {
			frontLightActivationType, _ = front.(*LightActivationType)
		}
		stage.OnAfterLightActivationTypeDeleteCallback.OnAfterDelete(stage, lightactivationtype, frontLightActivationType)
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

func (nitrogenpropertytype *NitrogenPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNitrogenPropertyTypeCreateCallback != nil {
		stage.OnAfterNitrogenPropertyTypeCreateCallback.OnAfterCreate(stage, nitrogenpropertytype)
	}
}

func (nitrogenpropertytype *NitrogenPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNitrogenPropertyTypeUpdateCallback != nil {
		var frontNitrogenPropertyType *NitrogenPropertyType
		if front != nil {
			frontNitrogenPropertyType, _ = front.(*NitrogenPropertyType)
		}
		stage.OnAfterNitrogenPropertyTypeUpdateCallback.OnAfterUpdate(stage, nitrogenpropertytype, frontNitrogenPropertyType)
	}
}

func (nitrogenpropertytype *NitrogenPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNitrogenPropertyTypeDeleteCallback != nil {
		var frontNitrogenPropertyType *NitrogenPropertyType
		if front != nil {
			frontNitrogenPropertyType, _ = front.(*NitrogenPropertyType)
		}
		stage.OnAfterNitrogenPropertyTypeDeleteCallback.OnAfterDelete(stage, nitrogenpropertytype, frontNitrogenPropertyType)
	}
}

func (nitrogentype *NitrogenType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNitrogenTypeCreateCallback != nil {
		stage.OnAfterNitrogenTypeCreateCallback.OnAfterCreate(stage, nitrogentype)
	}
}

func (nitrogentype *NitrogenType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNitrogenTypeUpdateCallback != nil {
		var frontNitrogenType *NitrogenType
		if front != nil {
			frontNitrogenType, _ = front.(*NitrogenType)
		}
		stage.OnAfterNitrogenTypeUpdateCallback.OnAfterUpdate(stage, nitrogentype, frontNitrogenType)
	}
}

func (nitrogentype *NitrogenType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNitrogenTypeDeleteCallback != nil {
		var frontNitrogenType *NitrogenType
		if front != nil {
			frontNitrogenType, _ = front.(*NitrogenType)
		}
		stage.OnAfterNitrogenTypeDeleteCallback.OnAfterDelete(stage, nitrogentype, frontNitrogenType)
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

func (oilpropertytype *OilPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOilPropertyTypeCreateCallback != nil {
		stage.OnAfterOilPropertyTypeCreateCallback.OnAfterCreate(stage, oilpropertytype)
	}
}

func (oilpropertytype *OilPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOilPropertyTypeUpdateCallback != nil {
		var frontOilPropertyType *OilPropertyType
		if front != nil {
			frontOilPropertyType, _ = front.(*OilPropertyType)
		}
		stage.OnAfterOilPropertyTypeUpdateCallback.OnAfterUpdate(stage, oilpropertytype, frontOilPropertyType)
	}
}

func (oilpropertytype *OilPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOilPropertyTypeDeleteCallback != nil {
		var frontOilPropertyType *OilPropertyType
		if front != nil {
			frontOilPropertyType, _ = front.(*OilPropertyType)
		}
		stage.OnAfterOilPropertyTypeDeleteCallback.OnAfterDelete(stage, oilpropertytype, frontOilPropertyType)
	}
}

func (oiltype *OilType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOilTypeCreateCallback != nil {
		stage.OnAfterOilTypeCreateCallback.OnAfterCreate(stage, oiltype)
	}
}

func (oiltype *OilType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOilTypeUpdateCallback != nil {
		var frontOilType *OilType
		if front != nil {
			frontOilType, _ = front.(*OilType)
		}
		stage.OnAfterOilTypeUpdateCallback.OnAfterUpdate(stage, oiltype, frontOilType)
	}
}

func (oiltype *OilType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOilTypeDeleteCallback != nil {
		var frontOilType *OilType
		if front != nil {
			frontOilType, _ = front.(*OilType)
		}
		stage.OnAfterOilTypeDeleteCallback.OnAfterDelete(stage, oiltype, frontOilType)
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

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrganisationAuthorityAssociationPropertyTypeCreateCallback != nil {
		stage.OnAfterOrganisationAuthorityAssociationPropertyTypeCreateCallback.OnAfterCreate(stage, organisationauthorityassociationpropertytype)
	}
}

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityAssociationPropertyTypeUpdateCallback != nil {
		var frontOrganisationAuthorityAssociationPropertyType *OrganisationAuthorityAssociationPropertyType
		if front != nil {
			frontOrganisationAuthorityAssociationPropertyType, _ = front.(*OrganisationAuthorityAssociationPropertyType)
		}
		stage.OnAfterOrganisationAuthorityAssociationPropertyTypeUpdateCallback.OnAfterUpdate(stage, organisationauthorityassociationpropertytype, frontOrganisationAuthorityAssociationPropertyType)
	}
}

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityAssociationPropertyTypeDeleteCallback != nil {
		var frontOrganisationAuthorityAssociationPropertyType *OrganisationAuthorityAssociationPropertyType
		if front != nil {
			frontOrganisationAuthorityAssociationPropertyType, _ = front.(*OrganisationAuthorityAssociationPropertyType)
		}
		stage.OnAfterOrganisationAuthorityAssociationPropertyTypeDeleteCallback.OnAfterDelete(stage, organisationauthorityassociationpropertytype, frontOrganisationAuthorityAssociationPropertyType)
	}
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrganisationAuthorityAssociationTypeCreateCallback != nil {
		stage.OnAfterOrganisationAuthorityAssociationTypeCreateCallback.OnAfterCreate(stage, organisationauthorityassociationtype)
	}
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityAssociationTypeUpdateCallback != nil {
		var frontOrganisationAuthorityAssociationType *OrganisationAuthorityAssociationType
		if front != nil {
			frontOrganisationAuthorityAssociationType, _ = front.(*OrganisationAuthorityAssociationType)
		}
		stage.OnAfterOrganisationAuthorityAssociationTypeUpdateCallback.OnAfterUpdate(stage, organisationauthorityassociationtype, frontOrganisationAuthorityAssociationType)
	}
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityAssociationTypeDeleteCallback != nil {
		var frontOrganisationAuthorityAssociationType *OrganisationAuthorityAssociationType
		if front != nil {
			frontOrganisationAuthorityAssociationType, _ = front.(*OrganisationAuthorityAssociationType)
		}
		stage.OnAfterOrganisationAuthorityAssociationTypeDeleteCallback.OnAfterDelete(stage, organisationauthorityassociationtype, frontOrganisationAuthorityAssociationType)
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

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrganisationAuthorityTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterOrganisationAuthorityTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, organisationauthoritytimeslicepropertytype)
	}
}

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityTimeSlicePropertyTypeUpdateCallback != nil {
		var frontOrganisationAuthorityTimeSlicePropertyType *OrganisationAuthorityTimeSlicePropertyType
		if front != nil {
			frontOrganisationAuthorityTimeSlicePropertyType, _ = front.(*OrganisationAuthorityTimeSlicePropertyType)
		}
		stage.OnAfterOrganisationAuthorityTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, organisationauthoritytimeslicepropertytype, frontOrganisationAuthorityTimeSlicePropertyType)
	}
}

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityTimeSlicePropertyTypeDeleteCallback != nil {
		var frontOrganisationAuthorityTimeSlicePropertyType *OrganisationAuthorityTimeSlicePropertyType
		if front != nil {
			frontOrganisationAuthorityTimeSlicePropertyType, _ = front.(*OrganisationAuthorityTimeSlicePropertyType)
		}
		stage.OnAfterOrganisationAuthorityTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, organisationauthoritytimeslicepropertytype, frontOrganisationAuthorityTimeSlicePropertyType)
	}
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrganisationAuthorityTimeSliceTypeCreateCallback != nil {
		stage.OnAfterOrganisationAuthorityTimeSliceTypeCreateCallback.OnAfterCreate(stage, organisationauthoritytimeslicetype)
	}
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityTimeSliceTypeUpdateCallback != nil {
		var frontOrganisationAuthorityTimeSliceType *OrganisationAuthorityTimeSliceType
		if front != nil {
			frontOrganisationAuthorityTimeSliceType, _ = front.(*OrganisationAuthorityTimeSliceType)
		}
		stage.OnAfterOrganisationAuthorityTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, organisationauthoritytimeslicetype, frontOrganisationAuthorityTimeSliceType)
	}
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityTimeSliceTypeDeleteCallback != nil {
		var frontOrganisationAuthorityTimeSliceType *OrganisationAuthorityTimeSliceType
		if front != nil {
			frontOrganisationAuthorityTimeSliceType, _ = front.(*OrganisationAuthorityTimeSliceType)
		}
		stage.OnAfterOrganisationAuthorityTimeSliceTypeDeleteCallback.OnAfterDelete(stage, organisationauthoritytimeslicetype, frontOrganisationAuthorityTimeSliceType)
	}
}

func (organisationauthoritytype *OrganisationAuthorityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrganisationAuthorityTypeCreateCallback != nil {
		stage.OnAfterOrganisationAuthorityTypeCreateCallback.OnAfterCreate(stage, organisationauthoritytype)
	}
}

func (organisationauthoritytype *OrganisationAuthorityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityTypeUpdateCallback != nil {
		var frontOrganisationAuthorityType *OrganisationAuthorityType
		if front != nil {
			frontOrganisationAuthorityType, _ = front.(*OrganisationAuthorityType)
		}
		stage.OnAfterOrganisationAuthorityTypeUpdateCallback.OnAfterUpdate(stage, organisationauthoritytype, frontOrganisationAuthorityType)
	}
}

func (organisationauthoritytype *OrganisationAuthorityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityTypeDeleteCallback != nil {
		var frontOrganisationAuthorityType *OrganisationAuthorityType
		if front != nil {
			frontOrganisationAuthorityType, _ = front.(*OrganisationAuthorityType)
		}
		stage.OnAfterOrganisationAuthorityTypeDeleteCallback.OnAfterDelete(stage, organisationauthoritytype, frontOrganisationAuthorityType)
	}
}

func (oxygenpropertytype *OxygenPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOxygenPropertyTypeCreateCallback != nil {
		stage.OnAfterOxygenPropertyTypeCreateCallback.OnAfterCreate(stage, oxygenpropertytype)
	}
}

func (oxygenpropertytype *OxygenPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOxygenPropertyTypeUpdateCallback != nil {
		var frontOxygenPropertyType *OxygenPropertyType
		if front != nil {
			frontOxygenPropertyType, _ = front.(*OxygenPropertyType)
		}
		stage.OnAfterOxygenPropertyTypeUpdateCallback.OnAfterUpdate(stage, oxygenpropertytype, frontOxygenPropertyType)
	}
}

func (oxygenpropertytype *OxygenPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOxygenPropertyTypeDeleteCallback != nil {
		var frontOxygenPropertyType *OxygenPropertyType
		if front != nil {
			frontOxygenPropertyType, _ = front.(*OxygenPropertyType)
		}
		stage.OnAfterOxygenPropertyTypeDeleteCallback.OnAfterDelete(stage, oxygenpropertytype, frontOxygenPropertyType)
	}
}

func (oxygentype *OxygenType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOxygenTypeCreateCallback != nil {
		stage.OnAfterOxygenTypeCreateCallback.OnAfterCreate(stage, oxygentype)
	}
}

func (oxygentype *OxygenType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOxygenTypeUpdateCallback != nil {
		var frontOxygenType *OxygenType
		if front != nil {
			frontOxygenType, _ = front.(*OxygenType)
		}
		stage.OnAfterOxygenTypeUpdateCallback.OnAfterUpdate(stage, oxygentype, frontOxygenType)
	}
}

func (oxygentype *OxygenType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOxygenTypeDeleteCallback != nil {
		var frontOxygenType *OxygenType
		if front != nil {
			frontOxygenType, _ = front.(*OxygenType)
		}
		stage.OnAfterOxygenTypeDeleteCallback.OnAfterDelete(stage, oxygentype, frontOxygenType)
	}
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPassengerServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterPassengerServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, passengerservicetimeslicepropertytype)
	}
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontPassengerServiceTimeSlicePropertyType *PassengerServiceTimeSlicePropertyType
		if front != nil {
			frontPassengerServiceTimeSlicePropertyType, _ = front.(*PassengerServiceTimeSlicePropertyType)
		}
		stage.OnAfterPassengerServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, passengerservicetimeslicepropertytype, frontPassengerServiceTimeSlicePropertyType)
	}
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontPassengerServiceTimeSlicePropertyType *PassengerServiceTimeSlicePropertyType
		if front != nil {
			frontPassengerServiceTimeSlicePropertyType, _ = front.(*PassengerServiceTimeSlicePropertyType)
		}
		stage.OnAfterPassengerServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, passengerservicetimeslicepropertytype, frontPassengerServiceTimeSlicePropertyType)
	}
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPassengerServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterPassengerServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, passengerservicetimeslicetype)
	}
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerServiceTimeSliceTypeUpdateCallback != nil {
		var frontPassengerServiceTimeSliceType *PassengerServiceTimeSliceType
		if front != nil {
			frontPassengerServiceTimeSliceType, _ = front.(*PassengerServiceTimeSliceType)
		}
		stage.OnAfterPassengerServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, passengerservicetimeslicetype, frontPassengerServiceTimeSliceType)
	}
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerServiceTimeSliceTypeDeleteCallback != nil {
		var frontPassengerServiceTimeSliceType *PassengerServiceTimeSliceType
		if front != nil {
			frontPassengerServiceTimeSliceType, _ = front.(*PassengerServiceTimeSliceType)
		}
		stage.OnAfterPassengerServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, passengerservicetimeslicetype, frontPassengerServiceTimeSliceType)
	}
}

func (passengerservicetype *PassengerServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPassengerServiceTypeCreateCallback != nil {
		stage.OnAfterPassengerServiceTypeCreateCallback.OnAfterCreate(stage, passengerservicetype)
	}
}

func (passengerservicetype *PassengerServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerServiceTypeUpdateCallback != nil {
		var frontPassengerServiceType *PassengerServiceType
		if front != nil {
			frontPassengerServiceType, _ = front.(*PassengerServiceType)
		}
		stage.OnAfterPassengerServiceTypeUpdateCallback.OnAfterUpdate(stage, passengerservicetype, frontPassengerServiceType)
	}
}

func (passengerservicetype *PassengerServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerServiceTypeDeleteCallback != nil {
		var frontPassengerServiceType *PassengerServiceType
		if front != nil {
			frontPassengerServiceType, _ = front.(*PassengerServiceType)
		}
		stage.OnAfterPassengerServiceTypeDeleteCallback.OnAfterDelete(stage, passengerservicetype, frontPassengerServiceType)
	}
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPilotControlledLightingTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterPilotControlledLightingTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, pilotcontrolledlightingtimeslicepropertytype)
	}
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPilotControlledLightingTimeSlicePropertyTypeUpdateCallback != nil {
		var frontPilotControlledLightingTimeSlicePropertyType *PilotControlledLightingTimeSlicePropertyType
		if front != nil {
			frontPilotControlledLightingTimeSlicePropertyType, _ = front.(*PilotControlledLightingTimeSlicePropertyType)
		}
		stage.OnAfterPilotControlledLightingTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, pilotcontrolledlightingtimeslicepropertytype, frontPilotControlledLightingTimeSlicePropertyType)
	}
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPilotControlledLightingTimeSlicePropertyTypeDeleteCallback != nil {
		var frontPilotControlledLightingTimeSlicePropertyType *PilotControlledLightingTimeSlicePropertyType
		if front != nil {
			frontPilotControlledLightingTimeSlicePropertyType, _ = front.(*PilotControlledLightingTimeSlicePropertyType)
		}
		stage.OnAfterPilotControlledLightingTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, pilotcontrolledlightingtimeslicepropertytype, frontPilotControlledLightingTimeSlicePropertyType)
	}
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPilotControlledLightingTimeSliceTypeCreateCallback != nil {
		stage.OnAfterPilotControlledLightingTimeSliceTypeCreateCallback.OnAfterCreate(stage, pilotcontrolledlightingtimeslicetype)
	}
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPilotControlledLightingTimeSliceTypeUpdateCallback != nil {
		var frontPilotControlledLightingTimeSliceType *PilotControlledLightingTimeSliceType
		if front != nil {
			frontPilotControlledLightingTimeSliceType, _ = front.(*PilotControlledLightingTimeSliceType)
		}
		stage.OnAfterPilotControlledLightingTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, pilotcontrolledlightingtimeslicetype, frontPilotControlledLightingTimeSliceType)
	}
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPilotControlledLightingTimeSliceTypeDeleteCallback != nil {
		var frontPilotControlledLightingTimeSliceType *PilotControlledLightingTimeSliceType
		if front != nil {
			frontPilotControlledLightingTimeSliceType, _ = front.(*PilotControlledLightingTimeSliceType)
		}
		stage.OnAfterPilotControlledLightingTimeSliceTypeDeleteCallback.OnAfterDelete(stage, pilotcontrolledlightingtimeslicetype, frontPilotControlledLightingTimeSliceType)
	}
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPilotControlledLightingTypeCreateCallback != nil {
		stage.OnAfterPilotControlledLightingTypeCreateCallback.OnAfterCreate(stage, pilotcontrolledlightingtype)
	}
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPilotControlledLightingTypeUpdateCallback != nil {
		var frontPilotControlledLightingType *PilotControlledLightingType
		if front != nil {
			frontPilotControlledLightingType, _ = front.(*PilotControlledLightingType)
		}
		stage.OnAfterPilotControlledLightingTypeUpdateCallback.OnAfterUpdate(stage, pilotcontrolledlightingtype, frontPilotControlledLightingType)
	}
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPilotControlledLightingTypeDeleteCallback != nil {
		var frontPilotControlledLightingType *PilotControlledLightingType
		if front != nil {
			frontPilotControlledLightingType, _ = front.(*PilotControlledLightingType)
		}
		stage.OnAfterPilotControlledLightingTypeDeleteCallback.OnAfterDelete(stage, pilotcontrolledlightingtype, frontPilotControlledLightingType)
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

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadioCommunicationChannelPropertyTypeCreateCallback != nil {
		stage.OnAfterRadioCommunicationChannelPropertyTypeCreateCallback.OnAfterCreate(stage, radiocommunicationchannelpropertytype)
	}
}

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationChannelPropertyTypeUpdateCallback != nil {
		var frontRadioCommunicationChannelPropertyType *RadioCommunicationChannelPropertyType
		if front != nil {
			frontRadioCommunicationChannelPropertyType, _ = front.(*RadioCommunicationChannelPropertyType)
		}
		stage.OnAfterRadioCommunicationChannelPropertyTypeUpdateCallback.OnAfterUpdate(stage, radiocommunicationchannelpropertytype, frontRadioCommunicationChannelPropertyType)
	}
}

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationChannelPropertyTypeDeleteCallback != nil {
		var frontRadioCommunicationChannelPropertyType *RadioCommunicationChannelPropertyType
		if front != nil {
			frontRadioCommunicationChannelPropertyType, _ = front.(*RadioCommunicationChannelPropertyType)
		}
		stage.OnAfterRadioCommunicationChannelPropertyTypeDeleteCallback.OnAfterDelete(stage, radiocommunicationchannelpropertytype, frontRadioCommunicationChannelPropertyType)
	}
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadioCommunicationChannelTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterRadioCommunicationChannelTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, radiocommunicationchanneltimeslicepropertytype)
	}
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationChannelTimeSlicePropertyTypeUpdateCallback != nil {
		var frontRadioCommunicationChannelTimeSlicePropertyType *RadioCommunicationChannelTimeSlicePropertyType
		if front != nil {
			frontRadioCommunicationChannelTimeSlicePropertyType, _ = front.(*RadioCommunicationChannelTimeSlicePropertyType)
		}
		stage.OnAfterRadioCommunicationChannelTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, radiocommunicationchanneltimeslicepropertytype, frontRadioCommunicationChannelTimeSlicePropertyType)
	}
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationChannelTimeSlicePropertyTypeDeleteCallback != nil {
		var frontRadioCommunicationChannelTimeSlicePropertyType *RadioCommunicationChannelTimeSlicePropertyType
		if front != nil {
			frontRadioCommunicationChannelTimeSlicePropertyType, _ = front.(*RadioCommunicationChannelTimeSlicePropertyType)
		}
		stage.OnAfterRadioCommunicationChannelTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, radiocommunicationchanneltimeslicepropertytype, frontRadioCommunicationChannelTimeSlicePropertyType)
	}
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadioCommunicationChannelTimeSliceTypeCreateCallback != nil {
		stage.OnAfterRadioCommunicationChannelTimeSliceTypeCreateCallback.OnAfterCreate(stage, radiocommunicationchanneltimeslicetype)
	}
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationChannelTimeSliceTypeUpdateCallback != nil {
		var frontRadioCommunicationChannelTimeSliceType *RadioCommunicationChannelTimeSliceType
		if front != nil {
			frontRadioCommunicationChannelTimeSliceType, _ = front.(*RadioCommunicationChannelTimeSliceType)
		}
		stage.OnAfterRadioCommunicationChannelTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, radiocommunicationchanneltimeslicetype, frontRadioCommunicationChannelTimeSliceType)
	}
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationChannelTimeSliceTypeDeleteCallback != nil {
		var frontRadioCommunicationChannelTimeSliceType *RadioCommunicationChannelTimeSliceType
		if front != nil {
			frontRadioCommunicationChannelTimeSliceType, _ = front.(*RadioCommunicationChannelTimeSliceType)
		}
		stage.OnAfterRadioCommunicationChannelTimeSliceTypeDeleteCallback.OnAfterDelete(stage, radiocommunicationchanneltimeslicetype, frontRadioCommunicationChannelTimeSliceType)
	}
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadioCommunicationChannelTypeCreateCallback != nil {
		stage.OnAfterRadioCommunicationChannelTypeCreateCallback.OnAfterCreate(stage, radiocommunicationchanneltype)
	}
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationChannelTypeUpdateCallback != nil {
		var frontRadioCommunicationChannelType *RadioCommunicationChannelType
		if front != nil {
			frontRadioCommunicationChannelType, _ = front.(*RadioCommunicationChannelType)
		}
		stage.OnAfterRadioCommunicationChannelTypeUpdateCallback.OnAfterUpdate(stage, radiocommunicationchanneltype, frontRadioCommunicationChannelType)
	}
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationChannelTypeDeleteCallback != nil {
		var frontRadioCommunicationChannelType *RadioCommunicationChannelType
		if front != nil {
			frontRadioCommunicationChannelType, _ = front.(*RadioCommunicationChannelType)
		}
		stage.OnAfterRadioCommunicationChannelTypeDeleteCallback.OnAfterDelete(stage, radiocommunicationchanneltype, frontRadioCommunicationChannelType)
	}
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadioCommunicationOperationalStatusPropertyTypeCreateCallback != nil {
		stage.OnAfterRadioCommunicationOperationalStatusPropertyTypeCreateCallback.OnAfterCreate(stage, radiocommunicationoperationalstatuspropertytype)
	}
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationOperationalStatusPropertyTypeUpdateCallback != nil {
		var frontRadioCommunicationOperationalStatusPropertyType *RadioCommunicationOperationalStatusPropertyType
		if front != nil {
			frontRadioCommunicationOperationalStatusPropertyType, _ = front.(*RadioCommunicationOperationalStatusPropertyType)
		}
		stage.OnAfterRadioCommunicationOperationalStatusPropertyTypeUpdateCallback.OnAfterUpdate(stage, radiocommunicationoperationalstatuspropertytype, frontRadioCommunicationOperationalStatusPropertyType)
	}
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationOperationalStatusPropertyTypeDeleteCallback != nil {
		var frontRadioCommunicationOperationalStatusPropertyType *RadioCommunicationOperationalStatusPropertyType
		if front != nil {
			frontRadioCommunicationOperationalStatusPropertyType, _ = front.(*RadioCommunicationOperationalStatusPropertyType)
		}
		stage.OnAfterRadioCommunicationOperationalStatusPropertyTypeDeleteCallback.OnAfterDelete(stage, radiocommunicationoperationalstatuspropertytype, frontRadioCommunicationOperationalStatusPropertyType)
	}
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRadioCommunicationOperationalStatusTypeCreateCallback != nil {
		stage.OnAfterRadioCommunicationOperationalStatusTypeCreateCallback.OnAfterCreate(stage, radiocommunicationoperationalstatustype)
	}
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationOperationalStatusTypeUpdateCallback != nil {
		var frontRadioCommunicationOperationalStatusType *RadioCommunicationOperationalStatusType
		if front != nil {
			frontRadioCommunicationOperationalStatusType, _ = front.(*RadioCommunicationOperationalStatusType)
		}
		stage.OnAfterRadioCommunicationOperationalStatusTypeUpdateCallback.OnAfterUpdate(stage, radiocommunicationoperationalstatustype, frontRadioCommunicationOperationalStatusType)
	}
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRadioCommunicationOperationalStatusTypeDeleteCallback != nil {
		var frontRadioCommunicationOperationalStatusType *RadioCommunicationOperationalStatusType
		if front != nil {
			frontRadioCommunicationOperationalStatusType, _ = front.(*RadioCommunicationOperationalStatusType)
		}
		stage.OnAfterRadioCommunicationOperationalStatusTypeDeleteCallback.OnAfterDelete(stage, radiocommunicationoperationalstatustype, frontRadioCommunicationOperationalStatusType)
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

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSearchRescueServiceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSearchRescueServiceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, searchrescueservicetimeslicepropertytype)
	}
}

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSearchRescueServiceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSearchRescueServiceTimeSlicePropertyType *SearchRescueServiceTimeSlicePropertyType
		if front != nil {
			frontSearchRescueServiceTimeSlicePropertyType, _ = front.(*SearchRescueServiceTimeSlicePropertyType)
		}
		stage.OnAfterSearchRescueServiceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, searchrescueservicetimeslicepropertytype, frontSearchRescueServiceTimeSlicePropertyType)
	}
}

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSearchRescueServiceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSearchRescueServiceTimeSlicePropertyType *SearchRescueServiceTimeSlicePropertyType
		if front != nil {
			frontSearchRescueServiceTimeSlicePropertyType, _ = front.(*SearchRescueServiceTimeSlicePropertyType)
		}
		stage.OnAfterSearchRescueServiceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, searchrescueservicetimeslicepropertytype, frontSearchRescueServiceTimeSlicePropertyType)
	}
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSearchRescueServiceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSearchRescueServiceTimeSliceTypeCreateCallback.OnAfterCreate(stage, searchrescueservicetimeslicetype)
	}
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSearchRescueServiceTimeSliceTypeUpdateCallback != nil {
		var frontSearchRescueServiceTimeSliceType *SearchRescueServiceTimeSliceType
		if front != nil {
			frontSearchRescueServiceTimeSliceType, _ = front.(*SearchRescueServiceTimeSliceType)
		}
		stage.OnAfterSearchRescueServiceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, searchrescueservicetimeslicetype, frontSearchRescueServiceTimeSliceType)
	}
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSearchRescueServiceTimeSliceTypeDeleteCallback != nil {
		var frontSearchRescueServiceTimeSliceType *SearchRescueServiceTimeSliceType
		if front != nil {
			frontSearchRescueServiceTimeSliceType, _ = front.(*SearchRescueServiceTimeSliceType)
		}
		stage.OnAfterSearchRescueServiceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, searchrescueservicetimeslicetype, frontSearchRescueServiceTimeSliceType)
	}
}

func (searchrescueservicetype *SearchRescueServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSearchRescueServiceTypeCreateCallback != nil {
		stage.OnAfterSearchRescueServiceTypeCreateCallback.OnAfterCreate(stage, searchrescueservicetype)
	}
}

func (searchrescueservicetype *SearchRescueServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSearchRescueServiceTypeUpdateCallback != nil {
		var frontSearchRescueServiceType *SearchRescueServiceType
		if front != nil {
			frontSearchRescueServiceType, _ = front.(*SearchRescueServiceType)
		}
		stage.OnAfterSearchRescueServiceTypeUpdateCallback.OnAfterUpdate(stage, searchrescueservicetype, frontSearchRescueServiceType)
	}
}

func (searchrescueservicetype *SearchRescueServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSearchRescueServiceTypeDeleteCallback != nil {
		var frontSearchRescueServiceType *SearchRescueServiceType
		if front != nil {
			frontSearchRescueServiceType, _ = front.(*SearchRescueServiceType)
		}
		stage.OnAfterSearchRescueServiceTypeDeleteCallback.OnAfterDelete(stage, searchrescueservicetype, frontSearchRescueServiceType)
	}
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterServiceOperationalStatusPropertyTypeCreateCallback != nil {
		stage.OnAfterServiceOperationalStatusPropertyTypeCreateCallback.OnAfterCreate(stage, serviceoperationalstatuspropertytype)
	}
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterServiceOperationalStatusPropertyTypeUpdateCallback != nil {
		var frontServiceOperationalStatusPropertyType *ServiceOperationalStatusPropertyType
		if front != nil {
			frontServiceOperationalStatusPropertyType, _ = front.(*ServiceOperationalStatusPropertyType)
		}
		stage.OnAfterServiceOperationalStatusPropertyTypeUpdateCallback.OnAfterUpdate(stage, serviceoperationalstatuspropertytype, frontServiceOperationalStatusPropertyType)
	}
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterServiceOperationalStatusPropertyTypeDeleteCallback != nil {
		var frontServiceOperationalStatusPropertyType *ServiceOperationalStatusPropertyType
		if front != nil {
			frontServiceOperationalStatusPropertyType, _ = front.(*ServiceOperationalStatusPropertyType)
		}
		stage.OnAfterServiceOperationalStatusPropertyTypeDeleteCallback.OnAfterDelete(stage, serviceoperationalstatuspropertytype, frontServiceOperationalStatusPropertyType)
	}
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterServiceOperationalStatusTypeCreateCallback != nil {
		stage.OnAfterServiceOperationalStatusTypeCreateCallback.OnAfterCreate(stage, serviceoperationalstatustype)
	}
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterServiceOperationalStatusTypeUpdateCallback != nil {
		var frontServiceOperationalStatusType *ServiceOperationalStatusType
		if front != nil {
			frontServiceOperationalStatusType, _ = front.(*ServiceOperationalStatusType)
		}
		stage.OnAfterServiceOperationalStatusTypeUpdateCallback.OnAfterUpdate(stage, serviceoperationalstatustype, frontServiceOperationalStatusType)
	}
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterServiceOperationalStatusTypeDeleteCallback != nil {
		var frontServiceOperationalStatusType *ServiceOperationalStatusType
		if front != nil {
			frontServiceOperationalStatusType, _ = front.(*ServiceOperationalStatusType)
		}
		stage.OnAfterServiceOperationalStatusTypeDeleteCallback.OnAfterDelete(stage, serviceoperationalstatustype, frontServiceOperationalStatusType)
	}
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialDateTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterSpecialDateTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, specialdatetimeslicepropertytype)
	}
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialDateTimeSlicePropertyTypeUpdateCallback != nil {
		var frontSpecialDateTimeSlicePropertyType *SpecialDateTimeSlicePropertyType
		if front != nil {
			frontSpecialDateTimeSlicePropertyType, _ = front.(*SpecialDateTimeSlicePropertyType)
		}
		stage.OnAfterSpecialDateTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, specialdatetimeslicepropertytype, frontSpecialDateTimeSlicePropertyType)
	}
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialDateTimeSlicePropertyTypeDeleteCallback != nil {
		var frontSpecialDateTimeSlicePropertyType *SpecialDateTimeSlicePropertyType
		if front != nil {
			frontSpecialDateTimeSlicePropertyType, _ = front.(*SpecialDateTimeSlicePropertyType)
		}
		stage.OnAfterSpecialDateTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, specialdatetimeslicepropertytype, frontSpecialDateTimeSlicePropertyType)
	}
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialDateTimeSliceTypeCreateCallback != nil {
		stage.OnAfterSpecialDateTimeSliceTypeCreateCallback.OnAfterCreate(stage, specialdatetimeslicetype)
	}
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialDateTimeSliceTypeUpdateCallback != nil {
		var frontSpecialDateTimeSliceType *SpecialDateTimeSliceType
		if front != nil {
			frontSpecialDateTimeSliceType, _ = front.(*SpecialDateTimeSliceType)
		}
		stage.OnAfterSpecialDateTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, specialdatetimeslicetype, frontSpecialDateTimeSliceType)
	}
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialDateTimeSliceTypeDeleteCallback != nil {
		var frontSpecialDateTimeSliceType *SpecialDateTimeSliceType
		if front != nil {
			frontSpecialDateTimeSliceType, _ = front.(*SpecialDateTimeSliceType)
		}
		stage.OnAfterSpecialDateTimeSliceTypeDeleteCallback.OnAfterDelete(stage, specialdatetimeslicetype, frontSpecialDateTimeSliceType)
	}
}

func (specialdatetype *SpecialDateType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialDateTypeCreateCallback != nil {
		stage.OnAfterSpecialDateTypeCreateCallback.OnAfterCreate(stage, specialdatetype)
	}
}

func (specialdatetype *SpecialDateType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialDateTypeUpdateCallback != nil {
		var frontSpecialDateType *SpecialDateType
		if front != nil {
			frontSpecialDateType, _ = front.(*SpecialDateType)
		}
		stage.OnAfterSpecialDateTypeUpdateCallback.OnAfterUpdate(stage, specialdatetype, frontSpecialDateType)
	}
}

func (specialdatetype *SpecialDateType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialDateTypeDeleteCallback != nil {
		var frontSpecialDateType *SpecialDateType
		if front != nil {
			frontSpecialDateType, _ = front.(*SpecialDateType)
		}
		stage.OnAfterSpecialDateTypeDeleteCallback.OnAfterDelete(stage, specialdatetype, frontSpecialDateType)
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

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelColumnTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterStandardLevelColumnTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, standardlevelcolumntimeslicepropertytype)
	}
}

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnTimeSlicePropertyTypeUpdateCallback != nil {
		var frontStandardLevelColumnTimeSlicePropertyType *StandardLevelColumnTimeSlicePropertyType
		if front != nil {
			frontStandardLevelColumnTimeSlicePropertyType, _ = front.(*StandardLevelColumnTimeSlicePropertyType)
		}
		stage.OnAfterStandardLevelColumnTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, standardlevelcolumntimeslicepropertytype, frontStandardLevelColumnTimeSlicePropertyType)
	}
}

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnTimeSlicePropertyTypeDeleteCallback != nil {
		var frontStandardLevelColumnTimeSlicePropertyType *StandardLevelColumnTimeSlicePropertyType
		if front != nil {
			frontStandardLevelColumnTimeSlicePropertyType, _ = front.(*StandardLevelColumnTimeSlicePropertyType)
		}
		stage.OnAfterStandardLevelColumnTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, standardlevelcolumntimeslicepropertytype, frontStandardLevelColumnTimeSlicePropertyType)
	}
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelColumnTimeSliceTypeCreateCallback != nil {
		stage.OnAfterStandardLevelColumnTimeSliceTypeCreateCallback.OnAfterCreate(stage, standardlevelcolumntimeslicetype)
	}
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnTimeSliceTypeUpdateCallback != nil {
		var frontStandardLevelColumnTimeSliceType *StandardLevelColumnTimeSliceType
		if front != nil {
			frontStandardLevelColumnTimeSliceType, _ = front.(*StandardLevelColumnTimeSliceType)
		}
		stage.OnAfterStandardLevelColumnTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, standardlevelcolumntimeslicetype, frontStandardLevelColumnTimeSliceType)
	}
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnTimeSliceTypeDeleteCallback != nil {
		var frontStandardLevelColumnTimeSliceType *StandardLevelColumnTimeSliceType
		if front != nil {
			frontStandardLevelColumnTimeSliceType, _ = front.(*StandardLevelColumnTimeSliceType)
		}
		stage.OnAfterStandardLevelColumnTimeSliceTypeDeleteCallback.OnAfterDelete(stage, standardlevelcolumntimeslicetype, frontStandardLevelColumnTimeSliceType)
	}
}

func (standardlevelcolumntype *StandardLevelColumnType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelColumnTypeCreateCallback != nil {
		stage.OnAfterStandardLevelColumnTypeCreateCallback.OnAfterCreate(stage, standardlevelcolumntype)
	}
}

func (standardlevelcolumntype *StandardLevelColumnType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnTypeUpdateCallback != nil {
		var frontStandardLevelColumnType *StandardLevelColumnType
		if front != nil {
			frontStandardLevelColumnType, _ = front.(*StandardLevelColumnType)
		}
		stage.OnAfterStandardLevelColumnTypeUpdateCallback.OnAfterUpdate(stage, standardlevelcolumntype, frontStandardLevelColumnType)
	}
}

func (standardlevelcolumntype *StandardLevelColumnType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnTypeDeleteCallback != nil {
		var frontStandardLevelColumnType *StandardLevelColumnType
		if front != nil {
			frontStandardLevelColumnType, _ = front.(*StandardLevelColumnType)
		}
		stage.OnAfterStandardLevelColumnTypeDeleteCallback.OnAfterDelete(stage, standardlevelcolumntype, frontStandardLevelColumnType)
	}
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelPropertyTypeCreateCallback != nil {
		stage.OnAfterStandardLevelPropertyTypeCreateCallback.OnAfterCreate(stage, standardlevelpropertytype)
	}
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelPropertyTypeUpdateCallback != nil {
		var frontStandardLevelPropertyType *StandardLevelPropertyType
		if front != nil {
			frontStandardLevelPropertyType, _ = front.(*StandardLevelPropertyType)
		}
		stage.OnAfterStandardLevelPropertyTypeUpdateCallback.OnAfterUpdate(stage, standardlevelpropertytype, frontStandardLevelPropertyType)
	}
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelPropertyTypeDeleteCallback != nil {
		var frontStandardLevelPropertyType *StandardLevelPropertyType
		if front != nil {
			frontStandardLevelPropertyType, _ = front.(*StandardLevelPropertyType)
		}
		stage.OnAfterStandardLevelPropertyTypeDeleteCallback.OnAfterDelete(stage, standardlevelpropertytype, frontStandardLevelPropertyType)
	}
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelSectorTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterStandardLevelSectorTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, standardlevelsectortimeslicepropertytype)
	}
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelSectorTimeSlicePropertyTypeUpdateCallback != nil {
		var frontStandardLevelSectorTimeSlicePropertyType *StandardLevelSectorTimeSlicePropertyType
		if front != nil {
			frontStandardLevelSectorTimeSlicePropertyType, _ = front.(*StandardLevelSectorTimeSlicePropertyType)
		}
		stage.OnAfterStandardLevelSectorTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, standardlevelsectortimeslicepropertytype, frontStandardLevelSectorTimeSlicePropertyType)
	}
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelSectorTimeSlicePropertyTypeDeleteCallback != nil {
		var frontStandardLevelSectorTimeSlicePropertyType *StandardLevelSectorTimeSlicePropertyType
		if front != nil {
			frontStandardLevelSectorTimeSlicePropertyType, _ = front.(*StandardLevelSectorTimeSlicePropertyType)
		}
		stage.OnAfterStandardLevelSectorTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, standardlevelsectortimeslicepropertytype, frontStandardLevelSectorTimeSlicePropertyType)
	}
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelSectorTimeSliceTypeCreateCallback != nil {
		stage.OnAfterStandardLevelSectorTimeSliceTypeCreateCallback.OnAfterCreate(stage, standardlevelsectortimeslicetype)
	}
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelSectorTimeSliceTypeUpdateCallback != nil {
		var frontStandardLevelSectorTimeSliceType *StandardLevelSectorTimeSliceType
		if front != nil {
			frontStandardLevelSectorTimeSliceType, _ = front.(*StandardLevelSectorTimeSliceType)
		}
		stage.OnAfterStandardLevelSectorTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, standardlevelsectortimeslicetype, frontStandardLevelSectorTimeSliceType)
	}
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelSectorTimeSliceTypeDeleteCallback != nil {
		var frontStandardLevelSectorTimeSliceType *StandardLevelSectorTimeSliceType
		if front != nil {
			frontStandardLevelSectorTimeSliceType, _ = front.(*StandardLevelSectorTimeSliceType)
		}
		stage.OnAfterStandardLevelSectorTimeSliceTypeDeleteCallback.OnAfterDelete(stage, standardlevelsectortimeslicetype, frontStandardLevelSectorTimeSliceType)
	}
}

func (standardlevelsectortype *StandardLevelSectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelSectorTypeCreateCallback != nil {
		stage.OnAfterStandardLevelSectorTypeCreateCallback.OnAfterCreate(stage, standardlevelsectortype)
	}
}

func (standardlevelsectortype *StandardLevelSectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelSectorTypeUpdateCallback != nil {
		var frontStandardLevelSectorType *StandardLevelSectorType
		if front != nil {
			frontStandardLevelSectorType, _ = front.(*StandardLevelSectorType)
		}
		stage.OnAfterStandardLevelSectorTypeUpdateCallback.OnAfterUpdate(stage, standardlevelsectortype, frontStandardLevelSectorType)
	}
}

func (standardlevelsectortype *StandardLevelSectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelSectorTypeDeleteCallback != nil {
		var frontStandardLevelSectorType *StandardLevelSectorType
		if front != nil {
			frontStandardLevelSectorType, _ = front.(*StandardLevelSectorType)
		}
		stage.OnAfterStandardLevelSectorTypeDeleteCallback.OnAfterDelete(stage, standardlevelsectortype, frontStandardLevelSectorType)
	}
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelTablePropertyTypeCreateCallback != nil {
		stage.OnAfterStandardLevelTablePropertyTypeCreateCallback.OnAfterCreate(stage, standardleveltablepropertytype)
	}
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTablePropertyTypeUpdateCallback != nil {
		var frontStandardLevelTablePropertyType *StandardLevelTablePropertyType
		if front != nil {
			frontStandardLevelTablePropertyType, _ = front.(*StandardLevelTablePropertyType)
		}
		stage.OnAfterStandardLevelTablePropertyTypeUpdateCallback.OnAfterUpdate(stage, standardleveltablepropertytype, frontStandardLevelTablePropertyType)
	}
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTablePropertyTypeDeleteCallback != nil {
		var frontStandardLevelTablePropertyType *StandardLevelTablePropertyType
		if front != nil {
			frontStandardLevelTablePropertyType, _ = front.(*StandardLevelTablePropertyType)
		}
		stage.OnAfterStandardLevelTablePropertyTypeDeleteCallback.OnAfterDelete(stage, standardleveltablepropertytype, frontStandardLevelTablePropertyType)
	}
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelTableTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterStandardLevelTableTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, standardleveltabletimeslicepropertytype)
	}
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTableTimeSlicePropertyTypeUpdateCallback != nil {
		var frontStandardLevelTableTimeSlicePropertyType *StandardLevelTableTimeSlicePropertyType
		if front != nil {
			frontStandardLevelTableTimeSlicePropertyType, _ = front.(*StandardLevelTableTimeSlicePropertyType)
		}
		stage.OnAfterStandardLevelTableTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, standardleveltabletimeslicepropertytype, frontStandardLevelTableTimeSlicePropertyType)
	}
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTableTimeSlicePropertyTypeDeleteCallback != nil {
		var frontStandardLevelTableTimeSlicePropertyType *StandardLevelTableTimeSlicePropertyType
		if front != nil {
			frontStandardLevelTableTimeSlicePropertyType, _ = front.(*StandardLevelTableTimeSlicePropertyType)
		}
		stage.OnAfterStandardLevelTableTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, standardleveltabletimeslicepropertytype, frontStandardLevelTableTimeSlicePropertyType)
	}
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelTableTimeSliceTypeCreateCallback != nil {
		stage.OnAfterStandardLevelTableTimeSliceTypeCreateCallback.OnAfterCreate(stage, standardleveltabletimeslicetype)
	}
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTableTimeSliceTypeUpdateCallback != nil {
		var frontStandardLevelTableTimeSliceType *StandardLevelTableTimeSliceType
		if front != nil {
			frontStandardLevelTableTimeSliceType, _ = front.(*StandardLevelTableTimeSliceType)
		}
		stage.OnAfterStandardLevelTableTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, standardleveltabletimeslicetype, frontStandardLevelTableTimeSliceType)
	}
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTableTimeSliceTypeDeleteCallback != nil {
		var frontStandardLevelTableTimeSliceType *StandardLevelTableTimeSliceType
		if front != nil {
			frontStandardLevelTableTimeSliceType, _ = front.(*StandardLevelTableTimeSliceType)
		}
		stage.OnAfterStandardLevelTableTimeSliceTypeDeleteCallback.OnAfterDelete(stage, standardleveltabletimeslicetype, frontStandardLevelTableTimeSliceType)
	}
}

func (standardleveltabletype *StandardLevelTableType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelTableTypeCreateCallback != nil {
		stage.OnAfterStandardLevelTableTypeCreateCallback.OnAfterCreate(stage, standardleveltabletype)
	}
}

func (standardleveltabletype *StandardLevelTableType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTableTypeUpdateCallback != nil {
		var frontStandardLevelTableType *StandardLevelTableType
		if front != nil {
			frontStandardLevelTableType, _ = front.(*StandardLevelTableType)
		}
		stage.OnAfterStandardLevelTableTypeUpdateCallback.OnAfterUpdate(stage, standardleveltabletype, frontStandardLevelTableType)
	}
}

func (standardleveltabletype *StandardLevelTableType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTableTypeDeleteCallback != nil {
		var frontStandardLevelTableType *StandardLevelTableType
		if front != nil {
			frontStandardLevelTableType, _ = front.(*StandardLevelTableType)
		}
		stage.OnAfterStandardLevelTableTypeDeleteCallback.OnAfterDelete(stage, standardleveltabletype, frontStandardLevelTableType)
	}
}

func (standardleveltype *StandardLevelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelTypeCreateCallback != nil {
		stage.OnAfterStandardLevelTypeCreateCallback.OnAfterCreate(stage, standardleveltype)
	}
}

func (standardleveltype *StandardLevelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTypeUpdateCallback != nil {
		var frontStandardLevelType *StandardLevelType
		if front != nil {
			frontStandardLevelType, _ = front.(*StandardLevelType)
		}
		stage.OnAfterStandardLevelTypeUpdateCallback.OnAfterUpdate(stage, standardleveltype, frontStandardLevelType)
	}
}

func (standardleveltype *StandardLevelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelTypeDeleteCallback != nil {
		var frontStandardLevelType *StandardLevelType
		if front != nil {
			frontStandardLevelType, _ = front.(*StandardLevelType)
		}
		stage.OnAfterStandardLevelTypeDeleteCallback.OnAfterDelete(stage, standardleveltype, frontStandardLevelType)
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

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitAvailabilityPropertyTypeCreateCallback != nil {
		stage.OnAfterUnitAvailabilityPropertyTypeCreateCallback.OnAfterCreate(stage, unitavailabilitypropertytype)
	}
}

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitAvailabilityPropertyTypeUpdateCallback != nil {
		var frontUnitAvailabilityPropertyType *UnitAvailabilityPropertyType
		if front != nil {
			frontUnitAvailabilityPropertyType, _ = front.(*UnitAvailabilityPropertyType)
		}
		stage.OnAfterUnitAvailabilityPropertyTypeUpdateCallback.OnAfterUpdate(stage, unitavailabilitypropertytype, frontUnitAvailabilityPropertyType)
	}
}

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitAvailabilityPropertyTypeDeleteCallback != nil {
		var frontUnitAvailabilityPropertyType *UnitAvailabilityPropertyType
		if front != nil {
			frontUnitAvailabilityPropertyType, _ = front.(*UnitAvailabilityPropertyType)
		}
		stage.OnAfterUnitAvailabilityPropertyTypeDeleteCallback.OnAfterDelete(stage, unitavailabilitypropertytype, frontUnitAvailabilityPropertyType)
	}
}

func (unitavailabilitytype *UnitAvailabilityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitAvailabilityTypeCreateCallback != nil {
		stage.OnAfterUnitAvailabilityTypeCreateCallback.OnAfterCreate(stage, unitavailabilitytype)
	}
}

func (unitavailabilitytype *UnitAvailabilityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitAvailabilityTypeUpdateCallback != nil {
		var frontUnitAvailabilityType *UnitAvailabilityType
		if front != nil {
			frontUnitAvailabilityType, _ = front.(*UnitAvailabilityType)
		}
		stage.OnAfterUnitAvailabilityTypeUpdateCallback.OnAfterUpdate(stage, unitavailabilitytype, frontUnitAvailabilityType)
	}
}

func (unitavailabilitytype *UnitAvailabilityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitAvailabilityTypeDeleteCallback != nil {
		var frontUnitAvailabilityType *UnitAvailabilityType
		if front != nil {
			frontUnitAvailabilityType, _ = front.(*UnitAvailabilityType)
		}
		stage.OnAfterUnitAvailabilityTypeDeleteCallback.OnAfterDelete(stage, unitavailabilitytype, frontUnitAvailabilityType)
	}
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitDependencyPropertyTypeCreateCallback != nil {
		stage.OnAfterUnitDependencyPropertyTypeCreateCallback.OnAfterCreate(stage, unitdependencypropertytype)
	}
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitDependencyPropertyTypeUpdateCallback != nil {
		var frontUnitDependencyPropertyType *UnitDependencyPropertyType
		if front != nil {
			frontUnitDependencyPropertyType, _ = front.(*UnitDependencyPropertyType)
		}
		stage.OnAfterUnitDependencyPropertyTypeUpdateCallback.OnAfterUpdate(stage, unitdependencypropertytype, frontUnitDependencyPropertyType)
	}
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitDependencyPropertyTypeDeleteCallback != nil {
		var frontUnitDependencyPropertyType *UnitDependencyPropertyType
		if front != nil {
			frontUnitDependencyPropertyType, _ = front.(*UnitDependencyPropertyType)
		}
		stage.OnAfterUnitDependencyPropertyTypeDeleteCallback.OnAfterDelete(stage, unitdependencypropertytype, frontUnitDependencyPropertyType)
	}
}

func (unitdependencytype *UnitDependencyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitDependencyTypeCreateCallback != nil {
		stage.OnAfterUnitDependencyTypeCreateCallback.OnAfterCreate(stage, unitdependencytype)
	}
}

func (unitdependencytype *UnitDependencyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitDependencyTypeUpdateCallback != nil {
		var frontUnitDependencyType *UnitDependencyType
		if front != nil {
			frontUnitDependencyType, _ = front.(*UnitDependencyType)
		}
		stage.OnAfterUnitDependencyTypeUpdateCallback.OnAfterUpdate(stage, unitdependencytype, frontUnitDependencyType)
	}
}

func (unitdependencytype *UnitDependencyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitDependencyTypeDeleteCallback != nil {
		var frontUnitDependencyType *UnitDependencyType
		if front != nil {
			frontUnitDependencyType, _ = front.(*UnitDependencyType)
		}
		stage.OnAfterUnitDependencyTypeDeleteCallback.OnAfterDelete(stage, unitdependencytype, frontUnitDependencyType)
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

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterUnitTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, unittimeslicepropertytype)
	}
}

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitTimeSlicePropertyTypeUpdateCallback != nil {
		var frontUnitTimeSlicePropertyType *UnitTimeSlicePropertyType
		if front != nil {
			frontUnitTimeSlicePropertyType, _ = front.(*UnitTimeSlicePropertyType)
		}
		stage.OnAfterUnitTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, unittimeslicepropertytype, frontUnitTimeSlicePropertyType)
	}
}

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitTimeSlicePropertyTypeDeleteCallback != nil {
		var frontUnitTimeSlicePropertyType *UnitTimeSlicePropertyType
		if front != nil {
			frontUnitTimeSlicePropertyType, _ = front.(*UnitTimeSlicePropertyType)
		}
		stage.OnAfterUnitTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, unittimeslicepropertytype, frontUnitTimeSlicePropertyType)
	}
}

func (unittimeslicetype *UnitTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitTimeSliceTypeCreateCallback != nil {
		stage.OnAfterUnitTimeSliceTypeCreateCallback.OnAfterCreate(stage, unittimeslicetype)
	}
}

func (unittimeslicetype *UnitTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitTimeSliceTypeUpdateCallback != nil {
		var frontUnitTimeSliceType *UnitTimeSliceType
		if front != nil {
			frontUnitTimeSliceType, _ = front.(*UnitTimeSliceType)
		}
		stage.OnAfterUnitTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, unittimeslicetype, frontUnitTimeSliceType)
	}
}

func (unittimeslicetype *UnitTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitTimeSliceTypeDeleteCallback != nil {
		var frontUnitTimeSliceType *UnitTimeSliceType
		if front != nil {
			frontUnitTimeSliceType, _ = front.(*UnitTimeSliceType)
		}
		stage.OnAfterUnitTimeSliceTypeDeleteCallback.OnAfterDelete(stage, unittimeslicetype, frontUnitTimeSliceType)
	}
}

func (unittype *UnitType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitTypeCreateCallback != nil {
		stage.OnAfterUnitTypeCreateCallback.OnAfterCreate(stage, unittype)
	}
}

func (unittype *UnitType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitTypeUpdateCallback != nil {
		var frontUnitType *UnitType
		if front != nil {
			frontUnitType, _ = front.(*UnitType)
		}
		stage.OnAfterUnitTypeUpdateCallback.OnAfterUpdate(stage, unittype, frontUnitType)
	}
}

func (unittype *UnitType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitTypeDeleteCallback != nil {
		var frontUnitType *UnitType
		if front != nil {
			frontUnitType, _ = front.(*UnitType)
		}
		stage.OnAfterUnitTypeDeleteCallback.OnAfterDelete(stage, unittype, frontUnitType)
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

