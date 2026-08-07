package models

func (stager *Stager) enforcePlantHasVaseAbstract() (needCommit bool) {
	return enforcePlantHasShape[*VaseAbstract](
		stager,
		func() *VaseAbstract { return new(VaseAbstract) },
		func(p *PlantAbstract) *VaseAbstract { return p.VaseAbstract },
		func(p *PlantAbstract, shape *VaseAbstract) { p.VaseAbstract = shape },
		func(p *PlantAbstract, shape *VaseAbstract) bool {
			return p.VaseAbstract == shape
		},
		"VaseAbstract",
	)
}

func (stager *Stager) enforceVaseAbstractName() (needCommit bool) {
	return enforcePlantShapeName[*VaseAbstract](
		stager,
		func(p *PlantAbstract) *VaseAbstract { return p.VaseAbstract },
		"VaseAbstract",
	)
}
