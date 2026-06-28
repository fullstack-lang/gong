package models

type OuterParticleOwner interface {
	GetOuterParticle() Particle
	SetOuterParticle(particle Particle)
}

type OuterParticleOwnerAbstract struct {
	// OuterParticle is the particle that reference the element
	// this is used in the factoring process of non normalized xsd
	OuterParticle Particle `xml:"-"`
}

func (outerParticleOwnerAbstract *OuterParticleOwnerAbstract) GetOuterParticle() Particle {
	return outerParticleOwnerAbstract.OuterParticle
}

func (outerParticleOwnerAbstract *OuterParticleOwnerAbstract) SetOuterParticle(particle Particle) {
	outerParticleOwnerAbstract.OuterParticle = particle
}
