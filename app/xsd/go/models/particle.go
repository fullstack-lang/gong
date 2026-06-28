package models

// Particle in XSD is a component that defines the structure of XML content.
// Particles can include individual elements, element groups (such as <xs:group>),
// and other complex structures like sequences (<xs:sequence>), choices (<xs:choice>),
// and all (<xs:all>). These particles dictate how XML elements are organized within a particular complex type.
//
// In essence, a particle is a generic term that encompasses any component that
// can occur in the content model of an XML schema, whether it's an individual element or a grouping of elements.
type Particle interface {
	GetOrder() int
	GetParent() Particle
	SetParentAndChildren(Particle) //
}

var _ Particle = (*Element)(nil)
var _ Particle = (*Group)(nil)
var _ Particle = (*All)(nil)
var _ Particle = (*Sequence)(nil)
var _ Particle = (*Choice)(nil)
var _ Particle = (*ComplexType)(nil)

type ParticleAbstract struct {
	// Order is the order at wich the particle was unmarshalled in the xsd
	// It is important to preserve the order output that is defined in the xsd
	Order int `xml:"-"`
	Depth int `xml:"-"`

	Parent   Particle
	Children []Particle
}

// GetOrder implements Particle.
func (p *ParticleAbstract) GetOrder() int {
	return p.Order
}

// GetOrder implements Particle.
func (p *ParticleAbstract) GetParent() Particle {
	return p.Parent
}

func (p *ParticleAbstract) GetChildren() Particle {
	return p.GetChildren()
}
