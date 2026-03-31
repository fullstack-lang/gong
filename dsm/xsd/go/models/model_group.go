package models

import (
	"cmp"
	"fmt"
	"log"
	"slices"
)

// ModelGroup is a construct that allows for
// the specification of complex structures within an XML document.
// Model groups define how elements within an XML schema are composed and ordered.
// There are three main types of model groups in XSD
type ModelGroup struct {
	OuterElementName string

	Sequences []*Sequence `xml:"sequence"`
	Alls      []*All      `xml:"all"`
	Choices   []*Choice   `xml:"choice"`
	Groups    []*Group    `xml:"group"`

	Elements []*Element `xml:"element"`

	ParticleAbstract
	OccurrenceDefinitionAbstract
}

func (modelGroup *ModelGroup) nameRecursively(name string) {
	modelGroup.OuterElementName = name
	for _, redefinable := range modelGroup.Groups {
		redefinable.Name = name + "_G_"
		redefinable.ModelGroup.nameRecursively(name + "_G_")
	}
	for _, redefinable := range modelGroup.Alls {
		redefinable.Name = name + "_A_"
		redefinable.ModelGroup.nameRecursively(name + "_A_")
	}
	for _, redefinable := range modelGroup.Sequences {
		redefinable.Name = name + "_S_"
		redefinable.ModelGroup.nameRecursively(name + "_S_")
	}
	for _, redefinable := range modelGroup.Choices {
		redefinable.Name = name + "_C_"
		redefinable.ModelGroup.nameRecursively(name + "_C_")
	}

}

func (modelGroup *ModelGroup) getParticles(
	groupMap map[string]*Group,
	map_Name_Elems map[string]*Element) (particles []Particle) {

	// log.Println("modelGroup.getElements", modelGroup.OuterElementName)
	for _, referenceGroup := range modelGroup.Groups {

		if referenceGroup.Ref != "" {
			// log.Println("Processing group", gRef.Name, gRef.Ref)
			if _, ok := groupMap[referenceGroup.Ref]; ok {
				particles = append(particles, referenceGroup)
			}
		}
	}
	for _, s := range modelGroup.Sequences {
		particles = append(particles, s.getParticles(groupMap, map_Name_Elems)...)
		for _, e := range s.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			particles = append(particles, e)
		}
	}
	for _, c := range modelGroup.Choices {
		particles = append(particles, c.getParticles(groupMap, map_Name_Elems)...)
		for _, e := range c.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			particles = append(particles, e)
		}
	}
	for _, a := range modelGroup.Alls {
		particles = append(particles, a.ModelGroup.getParticles(groupMap, map_Name_Elems)...)
		for _, e := range a.Elements {
			if _, ok := map_Name_Elems[e.Name]; ok {
				continue
			}
			map_Name_Elems[e.Name] = e
			particles = append(particles, e)
		}
	}

	// append the model group elements

	// reoder elements according to their rank
	slices.SortFunc(particles, func(a, b Particle) int {
		return cmp.Compare(a.GetOrder(), b.GetOrder())
	})

	return
}

func (modelGroup *ModelGroup) generateElements(
	map_Name_Elems map[string]*Element,
	stMap map[string]*SimpleType,
	ctMap map[string]*ComplexType,
	groupMap map[string]*Group,
	setOfFieldGoIdentifiers map[string]any,
	fields *string,
) {
	particles := modelGroup.getParticles(groupMap, map_Name_Elems)

	for _, particle := range particles {

		if elem, ok := particle.(*Element); ok {

			computeGoIdentifier(elem.GoIdentifier, &elem.WithGoIdentifier, setOfFieldGoIdentifiers)

			// remove namespace from type
			if NsPrefix(elem.Type) != "" {
				elem.Type = Name(elem.Type)
			}

			// compute if the element is unbounded
			unbounded := computeIsBounded(elem)

			sliceOrPointer := " *"
			if unbounded {
				sliceOrPointer = " []*"
			}

			// an element can be of 3 types:
			// 1. a simple type
			// 2. a named complex type
			// 3. an anonmous complex type
			goType := generateGoTypeFromType(elem.Type, stMap)
			if goType != "" {
				// 1. a simple type
				*fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + elem.Type +
					" order " + fmt.Sprintf("%d", elem.Order) + " depth " + fmt.Sprintf("%d", elem.Depth) +

					"\n\t" + elem.GoIdentifier + " " + goType + " " + "`" + `xml:"` + elem.NameXSD + `,omitempty"` + "`"
			} else {
				if elem.Type != "" {
					if ct, ok := ctMap[elem.Type]; ok {
						*fields += "\n\n\t// generated from element \"" + elem.NameXSD + "\" of type " + ct.Name +
							" order " + fmt.Sprintf("%d", elem.Order) + " depth " + fmt.Sprintf("%d", elem.Depth) +
							"\n\t" + elem.GoIdentifier + sliceOrPointer + ct.GoIdentifier + " " + "`" + `xml:"` +
							elem.NameXSD + `,omitempty"` + "`"
					} else {
						log.Println("element", elem.NameXSD, "unknown type", elem.Type)
					}
				} else {
					if elem.ComplexType == nil {
						log.Println("element", elem.Name, elem.NameXSD, "should have an anonymous complex type", elem.Type)
					} else {
						ct := elem.ComplexType
						*fields += "\n\n\t// generated from anonymous type within outer element \"" + elem.NameXSD +
							"\" of type " + ct.Name + "." +
							"\n\t" + elem.GoIdentifier + sliceOrPointer + ct.GoIdentifier + " " + "`" + `xml:"` + elem.NameXSD +
							`,omitempty"` + "`"
					}
				}

			}
		}

		if referenceGroup, ok := particle.(*Group); ok {

			computeGoIdentifier(referenceGroup.Ref, &referenceGroup.WithGoIdentifier, setOfFieldGoIdentifiers)

			if namedGroup, ok := groupMap[referenceGroup.Ref]; ok {

				// *fields += "\n\n\t// generated from group " +
				// 	"\n\t" + referenceGroup.GoIdentifier + " " + namedGroup.GoIdentifier // + " `xml:\",inline\"`"

				if !referenceGroup.HasNameConflict {
					*fields += "\n\n\t// generated from group with order " + fmt.Sprintf("%d", referenceGroup.Order) +
						" depth " + fmt.Sprintf("%d", referenceGroup.Depth) +
						"\n\t" + namedGroup.GoIdentifier // + " `xml:\",inline\"`"
				} else {
					log.Println("name conflict, ref", referenceGroup.GoIdentifier, " ", namedGroup.GoIdentifier)
				}
			}

		}

	}
}

func (modelGroup *ModelGroup) SetParentAndChildren(parent Particle) {

	modelGroup.Parent = parent
	for _, p := range modelGroup.Alls {
		modelGroup.Children = append(modelGroup.Children, p)
		p.SetParentAndChildren(modelGroup)
	}
	for _, p := range modelGroup.Sequences {
		modelGroup.Children = append(modelGroup.Children, p)
		p.SetParentAndChildren(modelGroup)
	}
	for _, p := range modelGroup.Choices {
		modelGroup.Children = append(modelGroup.Children, p)
		p.SetParentAndChildren(modelGroup)
	}
	for _, p := range modelGroup.Elements {
		modelGroup.Children = append(modelGroup.Children, p)
		p.SetParentAndChildren(modelGroup)
	}
}
