package models

type Choice struct {
	Name string
	Annotated

	ModelGroup

	IsDuplicatedInXSD bool

	OuterParticle Particle
}
