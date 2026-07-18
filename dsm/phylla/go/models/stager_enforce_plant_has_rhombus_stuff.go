package models

func (stager *Stager) enforcePlantHasRhombusStuff() (needCommit bool) {
	return enforcePlantHasShape[*RhombusStuff](
		stager,
		func() *RhombusStuff { return new(RhombusStuff) },
		func(p *Plant) *RhombusStuff { return p.RhombusStuff },
		func(p *Plant, shape *RhombusStuff) { p.RhombusStuff = shape },
		func(p *Plant, shape *RhombusStuff) bool {
			return p.RhombusStuff == shape
		},
		"RhombusStuff",
	)
}

func (stager *Stager) enforceRhombusStuffName() (needCommit bool) {
	return enforcePlantShapeName[*RhombusStuff](
		stager,
		func(p *Plant) *RhombusStuff { return p.RhombusStuff },
		"RhombusStuff",
	)
}
