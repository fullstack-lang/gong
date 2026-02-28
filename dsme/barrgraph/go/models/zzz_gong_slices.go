// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct ArtefactType
	// insertion point per field

	// Compute reverse map for named struct ArtefactTypeShape
	// insertion point per field

	// Compute reverse map for named struct Artist
	// insertion point per field

	// Compute reverse map for named struct ArtistShape
	// insertion point per field

	// Compute reverse map for named struct ControlPointShape
	// insertion point per field

	// Compute reverse map for named struct Desk
	// insertion point per field

	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_MovementShapes_reverseMap = make(map[*MovementShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _movementshape := range diagram.MovementShapes {
			stage.Diagram_MovementShapes_reverseMap[_movementshape] = diagram
		}
	}
	stage.Diagram_ArtefactTypeShapes_reverseMap = make(map[*ArtefactTypeShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
			stage.Diagram_ArtefactTypeShapes_reverseMap[_artefacttypeshape] = diagram
		}
	}
	stage.Diagram_ArtistShapes_reverseMap = make(map[*ArtistShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _artistshape := range diagram.ArtistShapes {
			stage.Diagram_ArtistShapes_reverseMap[_artistshape] = diagram
		}
	}
	stage.Diagram_InfluenceShapes_reverseMap = make(map[*InfluenceShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _influenceshape := range diagram.InfluenceShapes {
			stage.Diagram_InfluenceShapes_reverseMap[_influenceshape] = diagram
		}
	}

	// Compute reverse map for named struct Influence
	// insertion point per field

	// Compute reverse map for named struct InfluenceShape
	// insertion point per field
	stage.InfluenceShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*InfluenceShape)
	for influenceshape := range stage.InfluenceShapes {
		_ = influenceshape
		for _, _controlpointshape := range influenceshape.ControlPointShapes {
			stage.InfluenceShape_ControlPointShapes_reverseMap[_controlpointshape] = influenceshape
		}
	}

	// Compute reverse map for named struct Movement
	// insertion point per field
	stage.Movement_Places_reverseMap = make(map[*Place]*Movement)
	for movement := range stage.Movements {
		_ = movement
		for _, _place := range movement.Places {
			stage.Movement_Places_reverseMap[_place] = movement
		}
	}

	// Compute reverse map for named struct MovementShape
	// insertion point per field

	// Compute reverse map for named struct Place
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.ArtefactTypes {
		res = append(res, instance)
	}

	for instance := range stage.ArtefactTypeShapes {
		res = append(res, instance)
	}

	for instance := range stage.Artists {
		res = append(res, instance)
	}

	for instance := range stage.ArtistShapes {
		res = append(res, instance)
	}

	for instance := range stage.ControlPointShapes {
		res = append(res, instance)
	}

	for instance := range stage.Desks {
		res = append(res, instance)
	}

	for instance := range stage.Diagrams {
		res = append(res, instance)
	}

	for instance := range stage.Influences {
		res = append(res, instance)
	}

	for instance := range stage.InfluenceShapes {
		res = append(res, instance)
	}

	for instance := range stage.Movements {
		res = append(res, instance)
	}

	for instance := range stage.MovementShapes {
		res = append(res, instance)
	}

	for instance := range stage.Places {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (artefacttype *ArtefactType) GongCopy() GongstructIF {
	newInstance := new(ArtefactType)
	artefacttype.CopyBasicFields(newInstance)
	return newInstance
}

func (artefacttypeshape *ArtefactTypeShape) GongCopy() GongstructIF {
	newInstance := new(ArtefactTypeShape)
	artefacttypeshape.CopyBasicFields(newInstance)
	return newInstance
}

func (artist *Artist) GongCopy() GongstructIF {
	newInstance := new(Artist)
	artist.CopyBasicFields(newInstance)
	return newInstance
}

func (artistshape *ArtistShape) GongCopy() GongstructIF {
	newInstance := new(ArtistShape)
	artistshape.CopyBasicFields(newInstance)
	return newInstance
}

func (controlpointshape *ControlPointShape) GongCopy() GongstructIF {
	newInstance := new(ControlPointShape)
	controlpointshape.CopyBasicFields(newInstance)
	return newInstance
}

func (desk *Desk) GongCopy() GongstructIF {
	newInstance := new(Desk)
	desk.CopyBasicFields(newInstance)
	return newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := new(Diagram)
	diagram.CopyBasicFields(newInstance)
	return newInstance
}

func (influence *Influence) GongCopy() GongstructIF {
	newInstance := new(Influence)
	influence.CopyBasicFields(newInstance)
	return newInstance
}

func (influenceshape *InfluenceShape) GongCopy() GongstructIF {
	newInstance := new(InfluenceShape)
	influenceshape.CopyBasicFields(newInstance)
	return newInstance
}

func (movement *Movement) GongCopy() GongstructIF {
	newInstance := new(Movement)
	movement.CopyBasicFields(newInstance)
	return newInstance
}

func (movementshape *MovementShape) GongCopy() GongstructIF {
	newInstance := new(MovementShape)
	movementshape.CopyBasicFields(newInstance)
	return newInstance
}

func (place *Place) GongCopy() GongstructIF {
	newInstance := new(Place)
	place.CopyBasicFields(newInstance)
	return newInstance
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var artefacttypes_newInstances []*ArtefactType
	var artefacttypes_deletedInstances []*ArtefactType

	// parse all staged instances and check if they have a reference
	for artefacttype := range stage.ArtefactTypes {
		if ref, ok := stage.ArtefactTypes_reference[artefacttype]; !ok {
			artefacttypes_newInstances = append(artefacttypes_newInstances, artefacttype)
			newInstancesSlice = append(newInstancesSlice, artefacttype.GongMarshallIdentifier(stage))
			if stage.ArtefactTypes_referenceOrder == nil {
				stage.ArtefactTypes_referenceOrder = make(map[*ArtefactType]uint)
			}
			stage.ArtefactTypes_referenceOrder[artefacttype] = stage.ArtefactType_stagedOrder[artefacttype]
			newInstancesReverseSlice = append(newInstancesReverseSlice, artefacttype.GongMarshallUnstaging(stage))
			// delete(stage.ArtefactTypes_referenceOrder, artefacttype)
			fieldInitializers, pointersInitializations := artefacttype.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ArtefactType_stagedOrder[ref] = stage.ArtefactType_stagedOrder[artefacttype]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := artefacttype.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, artefacttype)
			// delete(stage.ArtefactType_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", artefacttype.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ArtefactTypes_reference {
		instance := stage.ArtefactTypes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ArtefactTypes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			artefacttypes_deletedInstances = append(artefacttypes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(artefacttypes_newInstances)
	lenDeletedInstances += len(artefacttypes_deletedInstances)
	var artefacttypeshapes_newInstances []*ArtefactTypeShape
	var artefacttypeshapes_deletedInstances []*ArtefactTypeShape

	// parse all staged instances and check if they have a reference
	for artefacttypeshape := range stage.ArtefactTypeShapes {
		if ref, ok := stage.ArtefactTypeShapes_reference[artefacttypeshape]; !ok {
			artefacttypeshapes_newInstances = append(artefacttypeshapes_newInstances, artefacttypeshape)
			newInstancesSlice = append(newInstancesSlice, artefacttypeshape.GongMarshallIdentifier(stage))
			if stage.ArtefactTypeShapes_referenceOrder == nil {
				stage.ArtefactTypeShapes_referenceOrder = make(map[*ArtefactTypeShape]uint)
			}
			stage.ArtefactTypeShapes_referenceOrder[artefacttypeshape] = stage.ArtefactTypeShape_stagedOrder[artefacttypeshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, artefacttypeshape.GongMarshallUnstaging(stage))
			// delete(stage.ArtefactTypeShapes_referenceOrder, artefacttypeshape)
			fieldInitializers, pointersInitializations := artefacttypeshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ArtefactTypeShape_stagedOrder[ref] = stage.ArtefactTypeShape_stagedOrder[artefacttypeshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := artefacttypeshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, artefacttypeshape)
			// delete(stage.ArtefactTypeShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", artefacttypeshape.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ArtefactTypeShapes_reference {
		instance := stage.ArtefactTypeShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ArtefactTypeShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			artefacttypeshapes_deletedInstances = append(artefacttypeshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(artefacttypeshapes_newInstances)
	lenDeletedInstances += len(artefacttypeshapes_deletedInstances)
	var artists_newInstances []*Artist
	var artists_deletedInstances []*Artist

	// parse all staged instances and check if they have a reference
	for artist := range stage.Artists {
		if ref, ok := stage.Artists_reference[artist]; !ok {
			artists_newInstances = append(artists_newInstances, artist)
			newInstancesSlice = append(newInstancesSlice, artist.GongMarshallIdentifier(stage))
			if stage.Artists_referenceOrder == nil {
				stage.Artists_referenceOrder = make(map[*Artist]uint)
			}
			stage.Artists_referenceOrder[artist] = stage.Artist_stagedOrder[artist]
			newInstancesReverseSlice = append(newInstancesReverseSlice, artist.GongMarshallUnstaging(stage))
			// delete(stage.Artists_referenceOrder, artist)
			fieldInitializers, pointersInitializations := artist.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Artist_stagedOrder[ref] = stage.Artist_stagedOrder[artist]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := artist.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, artist)
			// delete(stage.Artist_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", artist.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Artists_reference {
		instance := stage.Artists_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Artists[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			artists_deletedInstances = append(artists_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(artists_newInstances)
	lenDeletedInstances += len(artists_deletedInstances)
	var artistshapes_newInstances []*ArtistShape
	var artistshapes_deletedInstances []*ArtistShape

	// parse all staged instances and check if they have a reference
	for artistshape := range stage.ArtistShapes {
		if ref, ok := stage.ArtistShapes_reference[artistshape]; !ok {
			artistshapes_newInstances = append(artistshapes_newInstances, artistshape)
			newInstancesSlice = append(newInstancesSlice, artistshape.GongMarshallIdentifier(stage))
			if stage.ArtistShapes_referenceOrder == nil {
				stage.ArtistShapes_referenceOrder = make(map[*ArtistShape]uint)
			}
			stage.ArtistShapes_referenceOrder[artistshape] = stage.ArtistShape_stagedOrder[artistshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, artistshape.GongMarshallUnstaging(stage))
			// delete(stage.ArtistShapes_referenceOrder, artistshape)
			fieldInitializers, pointersInitializations := artistshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ArtistShape_stagedOrder[ref] = stage.ArtistShape_stagedOrder[artistshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := artistshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, artistshape)
			// delete(stage.ArtistShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", artistshape.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ArtistShapes_reference {
		instance := stage.ArtistShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ArtistShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			artistshapes_deletedInstances = append(artistshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(artistshapes_newInstances)
	lenDeletedInstances += len(artistshapes_deletedInstances)
	var controlpointshapes_newInstances []*ControlPointShape
	var controlpointshapes_deletedInstances []*ControlPointShape

	// parse all staged instances and check if they have a reference
	for controlpointshape := range stage.ControlPointShapes {
		if ref, ok := stage.ControlPointShapes_reference[controlpointshape]; !ok {
			controlpointshapes_newInstances = append(controlpointshapes_newInstances, controlpointshape)
			newInstancesSlice = append(newInstancesSlice, controlpointshape.GongMarshallIdentifier(stage))
			if stage.ControlPointShapes_referenceOrder == nil {
				stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint)
			}
			stage.ControlPointShapes_referenceOrder[controlpointshape] = stage.ControlPointShape_stagedOrder[controlpointshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, controlpointshape.GongMarshallUnstaging(stage))
			// delete(stage.ControlPointShapes_referenceOrder, controlpointshape)
			fieldInitializers, pointersInitializations := controlpointshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ControlPointShape_stagedOrder[ref] = stage.ControlPointShape_stagedOrder[controlpointshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := controlpointshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, controlpointshape)
			// delete(stage.ControlPointShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", controlpointshape.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ControlPointShapes_reference {
		instance := stage.ControlPointShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ControlPointShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			controlpointshapes_deletedInstances = append(controlpointshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(controlpointshapes_newInstances)
	lenDeletedInstances += len(controlpointshapes_deletedInstances)
	var desks_newInstances []*Desk
	var desks_deletedInstances []*Desk

	// parse all staged instances and check if they have a reference
	for desk := range stage.Desks {
		if ref, ok := stage.Desks_reference[desk]; !ok {
			desks_newInstances = append(desks_newInstances, desk)
			newInstancesSlice = append(newInstancesSlice, desk.GongMarshallIdentifier(stage))
			if stage.Desks_referenceOrder == nil {
				stage.Desks_referenceOrder = make(map[*Desk]uint)
			}
			stage.Desks_referenceOrder[desk] = stage.Desk_stagedOrder[desk]
			newInstancesReverseSlice = append(newInstancesReverseSlice, desk.GongMarshallUnstaging(stage))
			// delete(stage.Desks_referenceOrder, desk)
			fieldInitializers, pointersInitializations := desk.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Desk_stagedOrder[ref] = stage.Desk_stagedOrder[desk]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := desk.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, desk)
			// delete(stage.Desk_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", desk.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Desks_reference {
		instance := stage.Desks_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Desks[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			desks_deletedInstances = append(desks_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(desks_newInstances)
	lenDeletedInstances += len(desks_deletedInstances)
	var diagrams_newInstances []*Diagram
	var diagrams_deletedInstances []*Diagram

	// parse all staged instances and check if they have a reference
	for diagram := range stage.Diagrams {
		if ref, ok := stage.Diagrams_reference[diagram]; !ok {
			diagrams_newInstances = append(diagrams_newInstances, diagram)
			newInstancesSlice = append(newInstancesSlice, diagram.GongMarshallIdentifier(stage))
			if stage.Diagrams_referenceOrder == nil {
				stage.Diagrams_referenceOrder = make(map[*Diagram]uint)
			}
			stage.Diagrams_referenceOrder[diagram] = stage.Diagram_stagedOrder[diagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagram.GongMarshallUnstaging(stage))
			// delete(stage.Diagrams_referenceOrder, diagram)
			fieldInitializers, pointersInitializations := diagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Diagram_stagedOrder[ref] = stage.Diagram_stagedOrder[diagram]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagram)
			// delete(stage.Diagram_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", diagram.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Diagrams_reference {
		instance := stage.Diagrams_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Diagrams[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagrams_deletedInstances = append(diagrams_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagrams_newInstances)
	lenDeletedInstances += len(diagrams_deletedInstances)
	var influences_newInstances []*Influence
	var influences_deletedInstances []*Influence

	// parse all staged instances and check if they have a reference
	for influence := range stage.Influences {
		if ref, ok := stage.Influences_reference[influence]; !ok {
			influences_newInstances = append(influences_newInstances, influence)
			newInstancesSlice = append(newInstancesSlice, influence.GongMarshallIdentifier(stage))
			if stage.Influences_referenceOrder == nil {
				stage.Influences_referenceOrder = make(map[*Influence]uint)
			}
			stage.Influences_referenceOrder[influence] = stage.Influence_stagedOrder[influence]
			newInstancesReverseSlice = append(newInstancesReverseSlice, influence.GongMarshallUnstaging(stage))
			// delete(stage.Influences_referenceOrder, influence)
			fieldInitializers, pointersInitializations := influence.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Influence_stagedOrder[ref] = stage.Influence_stagedOrder[influence]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := influence.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, influence)
			// delete(stage.Influence_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", influence.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Influences_reference {
		instance := stage.Influences_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Influences[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			influences_deletedInstances = append(influences_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(influences_newInstances)
	lenDeletedInstances += len(influences_deletedInstances)
	var influenceshapes_newInstances []*InfluenceShape
	var influenceshapes_deletedInstances []*InfluenceShape

	// parse all staged instances and check if they have a reference
	for influenceshape := range stage.InfluenceShapes {
		if ref, ok := stage.InfluenceShapes_reference[influenceshape]; !ok {
			influenceshapes_newInstances = append(influenceshapes_newInstances, influenceshape)
			newInstancesSlice = append(newInstancesSlice, influenceshape.GongMarshallIdentifier(stage))
			if stage.InfluenceShapes_referenceOrder == nil {
				stage.InfluenceShapes_referenceOrder = make(map[*InfluenceShape]uint)
			}
			stage.InfluenceShapes_referenceOrder[influenceshape] = stage.InfluenceShape_stagedOrder[influenceshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, influenceshape.GongMarshallUnstaging(stage))
			// delete(stage.InfluenceShapes_referenceOrder, influenceshape)
			fieldInitializers, pointersInitializations := influenceshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.InfluenceShape_stagedOrder[ref] = stage.InfluenceShape_stagedOrder[influenceshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := influenceshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, influenceshape)
			// delete(stage.InfluenceShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", influenceshape.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.InfluenceShapes_reference {
		instance := stage.InfluenceShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.InfluenceShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			influenceshapes_deletedInstances = append(influenceshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(influenceshapes_newInstances)
	lenDeletedInstances += len(influenceshapes_deletedInstances)
	var movements_newInstances []*Movement
	var movements_deletedInstances []*Movement

	// parse all staged instances and check if they have a reference
	for movement := range stage.Movements {
		if ref, ok := stage.Movements_reference[movement]; !ok {
			movements_newInstances = append(movements_newInstances, movement)
			newInstancesSlice = append(newInstancesSlice, movement.GongMarshallIdentifier(stage))
			if stage.Movements_referenceOrder == nil {
				stage.Movements_referenceOrder = make(map[*Movement]uint)
			}
			stage.Movements_referenceOrder[movement] = stage.Movement_stagedOrder[movement]
			newInstancesReverseSlice = append(newInstancesReverseSlice, movement.GongMarshallUnstaging(stage))
			// delete(stage.Movements_referenceOrder, movement)
			fieldInitializers, pointersInitializations := movement.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Movement_stagedOrder[ref] = stage.Movement_stagedOrder[movement]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := movement.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, movement)
			// delete(stage.Movement_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", movement.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Movements_reference {
		instance := stage.Movements_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Movements[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			movements_deletedInstances = append(movements_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(movements_newInstances)
	lenDeletedInstances += len(movements_deletedInstances)
	var movementshapes_newInstances []*MovementShape
	var movementshapes_deletedInstances []*MovementShape

	// parse all staged instances and check if they have a reference
	for movementshape := range stage.MovementShapes {
		if ref, ok := stage.MovementShapes_reference[movementshape]; !ok {
			movementshapes_newInstances = append(movementshapes_newInstances, movementshape)
			newInstancesSlice = append(newInstancesSlice, movementshape.GongMarshallIdentifier(stage))
			if stage.MovementShapes_referenceOrder == nil {
				stage.MovementShapes_referenceOrder = make(map[*MovementShape]uint)
			}
			stage.MovementShapes_referenceOrder[movementshape] = stage.MovementShape_stagedOrder[movementshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, movementshape.GongMarshallUnstaging(stage))
			// delete(stage.MovementShapes_referenceOrder, movementshape)
			fieldInitializers, pointersInitializations := movementshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MovementShape_stagedOrder[ref] = stage.MovementShape_stagedOrder[movementshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := movementshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, movementshape)
			// delete(stage.MovementShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", movementshape.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.MovementShapes_reference {
		instance := stage.MovementShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MovementShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			movementshapes_deletedInstances = append(movementshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(movementshapes_newInstances)
	lenDeletedInstances += len(movementshapes_deletedInstances)
	var places_newInstances []*Place
	var places_deletedInstances []*Place

	// parse all staged instances and check if they have a reference
	for place := range stage.Places {
		if ref, ok := stage.Places_reference[place]; !ok {
			places_newInstances = append(places_newInstances, place)
			newInstancesSlice = append(newInstancesSlice, place.GongMarshallIdentifier(stage))
			if stage.Places_referenceOrder == nil {
				stage.Places_referenceOrder = make(map[*Place]uint)
			}
			stage.Places_referenceOrder[place] = stage.Place_stagedOrder[place]
			newInstancesReverseSlice = append(newInstancesReverseSlice, place.GongMarshallUnstaging(stage))
			// delete(stage.Places_referenceOrder, place)
			fieldInitializers, pointersInitializations := place.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Place_stagedOrder[ref] = stage.Place_stagedOrder[place]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := place.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, place)
			// delete(stage.Place_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", place.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Places_reference {
		instance := stage.Places_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Places[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			places_deletedInstances = append(places_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(places_newInstances)
	lenDeletedInstances += len(places_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.ArtefactTypes_reference = make(map[*ArtefactType]*ArtefactType)
	stage.ArtefactTypes_referenceOrder = make(map[*ArtefactType]uint) // diff Unstage needs the reference order
	stage.ArtefactTypes_instance = make(map[*ArtefactType]*ArtefactType)
	for instance := range stage.ArtefactTypes {
		_copy := instance.GongCopy().(*ArtefactType)
		stage.ArtefactTypes_reference[instance] = _copy
		stage.ArtefactTypes_instance[_copy] = instance
		stage.ArtefactTypes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ArtefactTypeShapes_reference = make(map[*ArtefactTypeShape]*ArtefactTypeShape)
	stage.ArtefactTypeShapes_referenceOrder = make(map[*ArtefactTypeShape]uint) // diff Unstage needs the reference order
	stage.ArtefactTypeShapes_instance = make(map[*ArtefactTypeShape]*ArtefactTypeShape)
	for instance := range stage.ArtefactTypeShapes {
		_copy := instance.GongCopy().(*ArtefactTypeShape)
		stage.ArtefactTypeShapes_reference[instance] = _copy
		stage.ArtefactTypeShapes_instance[_copy] = instance
		stage.ArtefactTypeShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Artists_reference = make(map[*Artist]*Artist)
	stage.Artists_referenceOrder = make(map[*Artist]uint) // diff Unstage needs the reference order
	stage.Artists_instance = make(map[*Artist]*Artist)
	for instance := range stage.Artists {
		_copy := instance.GongCopy().(*Artist)
		stage.Artists_reference[instance] = _copy
		stage.Artists_instance[_copy] = instance
		stage.Artists_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ArtistShapes_reference = make(map[*ArtistShape]*ArtistShape)
	stage.ArtistShapes_referenceOrder = make(map[*ArtistShape]uint) // diff Unstage needs the reference order
	stage.ArtistShapes_instance = make(map[*ArtistShape]*ArtistShape)
	for instance := range stage.ArtistShapes {
		_copy := instance.GongCopy().(*ArtistShape)
		stage.ArtistShapes_reference[instance] = _copy
		stage.ArtistShapes_instance[_copy] = instance
		stage.ArtistShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ControlPointShapes_reference = make(map[*ControlPointShape]*ControlPointShape)
	stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint) // diff Unstage needs the reference order
	stage.ControlPointShapes_instance = make(map[*ControlPointShape]*ControlPointShape)
	for instance := range stage.ControlPointShapes {
		_copy := instance.GongCopy().(*ControlPointShape)
		stage.ControlPointShapes_reference[instance] = _copy
		stage.ControlPointShapes_instance[_copy] = instance
		stage.ControlPointShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Desks_reference = make(map[*Desk]*Desk)
	stage.Desks_referenceOrder = make(map[*Desk]uint) // diff Unstage needs the reference order
	stage.Desks_instance = make(map[*Desk]*Desk)
	for instance := range stage.Desks {
		_copy := instance.GongCopy().(*Desk)
		stage.Desks_reference[instance] = _copy
		stage.Desks_instance[_copy] = instance
		stage.Desks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint) // diff Unstage needs the reference order
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	for instance := range stage.Diagrams {
		_copy := instance.GongCopy().(*Diagram)
		stage.Diagrams_reference[instance] = _copy
		stage.Diagrams_instance[_copy] = instance
		stage.Diagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Influences_reference = make(map[*Influence]*Influence)
	stage.Influences_referenceOrder = make(map[*Influence]uint) // diff Unstage needs the reference order
	stage.Influences_instance = make(map[*Influence]*Influence)
	for instance := range stage.Influences {
		_copy := instance.GongCopy().(*Influence)
		stage.Influences_reference[instance] = _copy
		stage.Influences_instance[_copy] = instance
		stage.Influences_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.InfluenceShapes_reference = make(map[*InfluenceShape]*InfluenceShape)
	stage.InfluenceShapes_referenceOrder = make(map[*InfluenceShape]uint) // diff Unstage needs the reference order
	stage.InfluenceShapes_instance = make(map[*InfluenceShape]*InfluenceShape)
	for instance := range stage.InfluenceShapes {
		_copy := instance.GongCopy().(*InfluenceShape)
		stage.InfluenceShapes_reference[instance] = _copy
		stage.InfluenceShapes_instance[_copy] = instance
		stage.InfluenceShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Movements_reference = make(map[*Movement]*Movement)
	stage.Movements_referenceOrder = make(map[*Movement]uint) // diff Unstage needs the reference order
	stage.Movements_instance = make(map[*Movement]*Movement)
	for instance := range stage.Movements {
		_copy := instance.GongCopy().(*Movement)
		stage.Movements_reference[instance] = _copy
		stage.Movements_instance[_copy] = instance
		stage.Movements_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MovementShapes_reference = make(map[*MovementShape]*MovementShape)
	stage.MovementShapes_referenceOrder = make(map[*MovementShape]uint) // diff Unstage needs the reference order
	stage.MovementShapes_instance = make(map[*MovementShape]*MovementShape)
	for instance := range stage.MovementShapes {
		_copy := instance.GongCopy().(*MovementShape)
		stage.MovementShapes_reference[instance] = _copy
		stage.MovementShapes_instance[_copy] = instance
		stage.MovementShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Places_reference = make(map[*Place]*Place)
	stage.Places_referenceOrder = make(map[*Place]uint) // diff Unstage needs the reference order
	stage.Places_instance = make(map[*Place]*Place)
	for instance := range stage.Places {
		_copy := instance.GongCopy().(*Place)
		stage.Places_reference[instance] = _copy
		stage.Places_instance[_copy] = instance
		stage.Places_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.ArtefactTypes {
		reference := stage.ArtefactTypes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ArtefactTypeShapes {
		reference := stage.ArtefactTypeShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Artists {
		reference := stage.Artists_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ArtistShapes {
		reference := stage.ArtistShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ControlPointShapes {
		reference := stage.ControlPointShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Desks {
		reference := stage.Desks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Diagrams {
		reference := stage.Diagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Influences {
		reference := stage.Influences_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.InfluenceShapes {
		reference := stage.InfluenceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Movements {
		reference := stage.Movements_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MovementShapes {
		reference := stage.MovementShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Places {
		reference := stage.Places_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (artefacttype *ArtefactType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ArtefactType_stagedOrder[artefacttype]; ok {
		return order
	}
	if order, ok := stage.ArtefactTypes_referenceOrder[artefacttype]; ok {
		return order
	} else {
		log.Printf("instance %p of type ArtefactType was not staged and does not have a reference order", artefacttype)
		return 0
	}
}

func (artefacttypeshape *ArtefactTypeShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ArtefactTypeShape_stagedOrder[artefacttypeshape]; ok {
		return order
	}
	if order, ok := stage.ArtefactTypeShapes_referenceOrder[artefacttypeshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ArtefactTypeShape was not staged and does not have a reference order", artefacttypeshape)
		return 0
	}
}

func (artist *Artist) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Artist_stagedOrder[artist]; ok {
		return order
	}
	if order, ok := stage.Artists_referenceOrder[artist]; ok {
		return order
	} else {
		log.Printf("instance %p of type Artist was not staged and does not have a reference order", artist)
		return 0
	}
}

func (artistshape *ArtistShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ArtistShape_stagedOrder[artistshape]; ok {
		return order
	}
	if order, ok := stage.ArtistShapes_referenceOrder[artistshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ArtistShape was not staged and does not have a reference order", artistshape)
		return 0
	}
}

func (controlpointshape *ControlPointShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ControlPointShape_stagedOrder[controlpointshape]; ok {
		return order
	}
	if order, ok := stage.ControlPointShapes_referenceOrder[controlpointshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ControlPointShape was not staged and does not have a reference order", controlpointshape)
		return 0
	}
}

func (desk *Desk) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Desk_stagedOrder[desk]; ok {
		return order
	}
	if order, ok := stage.Desks_referenceOrder[desk]; ok {
		return order
	} else {
		log.Printf("instance %p of type Desk was not staged and does not have a reference order", desk)
		return 0
	}
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Diagram_stagedOrder[diagram]; ok {
		return order
	}
	if order, ok := stage.Diagrams_referenceOrder[diagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Diagram was not staged and does not have a reference order", diagram)
		return 0
	}
}

func (influence *Influence) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Influence_stagedOrder[influence]; ok {
		return order
	}
	if order, ok := stage.Influences_referenceOrder[influence]; ok {
		return order
	} else {
		log.Printf("instance %p of type Influence was not staged and does not have a reference order", influence)
		return 0
	}
}

func (influenceshape *InfluenceShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.InfluenceShape_stagedOrder[influenceshape]; ok {
		return order
	}
	if order, ok := stage.InfluenceShapes_referenceOrder[influenceshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type InfluenceShape was not staged and does not have a reference order", influenceshape)
		return 0
	}
}

func (movement *Movement) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Movement_stagedOrder[movement]; ok {
		return order
	}
	if order, ok := stage.Movements_referenceOrder[movement]; ok {
		return order
	} else {
		log.Printf("instance %p of type Movement was not staged and does not have a reference order", movement)
		return 0
	}
}

func (movementshape *MovementShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MovementShape_stagedOrder[movementshape]; ok {
		return order
	}
	if order, ok := stage.MovementShapes_referenceOrder[movementshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type MovementShape was not staged and does not have a reference order", movementshape)
		return 0
	}
}

func (place *Place) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Place_stagedOrder[place]; ok {
		return order
	}
	if order, ok := stage.Places_referenceOrder[place]; ok {
		return order
	} else {
		log.Printf("instance %p of type Place was not staged and does not have a reference order", place)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (artefacttype *ArtefactType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artefacttype.GongGetGongstructName(), artefacttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artefacttype *ArtefactType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artefacttype.GongGetGongstructName(), artefacttype.GongGetOrder(stage))
}

func (artefacttypeshape *ArtefactTypeShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artefacttypeshape.GongGetGongstructName(), artefacttypeshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artefacttypeshape *ArtefactTypeShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artefacttypeshape.GongGetGongstructName(), artefacttypeshape.GongGetOrder(stage))
}

func (artist *Artist) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artist.GongGetGongstructName(), artist.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artist *Artist) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artist.GongGetGongstructName(), artist.GongGetOrder(stage))
}

func (artistshape *ArtistShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artistshape.GongGetGongstructName(), artistshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (artistshape *ArtistShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", artistshape.GongGetGongstructName(), artistshape.GongGetOrder(stage))
}

func (controlpointshape *ControlPointShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpointshape *ControlPointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
}

func (desk *Desk) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", desk.GongGetGongstructName(), desk.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (desk *Desk) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", desk.GongGetGongstructName(), desk.GongGetOrder(stage))
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

func (influence *Influence) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influence.GongGetGongstructName(), influence.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (influence *Influence) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influence.GongGetGongstructName(), influence.GongGetOrder(stage))
}

func (influenceshape *InfluenceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influenceshape.GongGetGongstructName(), influenceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (influenceshape *InfluenceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influenceshape.GongGetGongstructName(), influenceshape.GongGetOrder(stage))
}

func (movement *Movement) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", movement.GongGetGongstructName(), movement.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (movement *Movement) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", movement.GongGetGongstructName(), movement.GongGetOrder(stage))
}

func (movementshape *MovementShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", movementshape.GongGetGongstructName(), movementshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (movementshape *MovementShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", movementshape.GongGetGongstructName(), movementshape.GongGetOrder(stage))
}

func (place *Place) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", place.GongGetGongstructName(), place.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (place *Place) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", place.GongGetGongstructName(), place.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (artefacttype *ArtefactType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artefacttype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArtefactType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(artefacttype.Name))
	return
}

func (artefacttypeshape *ArtefactTypeShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artefacttypeshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArtefactTypeShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(artefacttypeshape.Name))
	return
}

func (artist *Artist) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artist.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Artist")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(artist.Name))
	return
}

func (artistshape *ArtistShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artistshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ArtistShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(artistshape.Name))
	return
}

func (controlpointshape *ControlPointShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlPointShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(controlpointshape.Name))
	return
}

func (desk *Desk) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", desk.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Desk")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(desk.Name))
	return
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.Name))
	return
}

func (influence *Influence) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influence.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Influence")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(influence.Name))
	return
}

func (influenceshape *InfluenceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influenceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "InfluenceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(influenceshape.Name))
	return
}

func (movement *Movement) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", movement.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Movement")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(movement.Name))
	return
}

func (movementshape *MovementShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", movementshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MovementShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(movementshape.Name))
	return
}

func (place *Place) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", place.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Place")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(place.Name))
	return
}

// insertion point for unstaging
func (artefacttype *ArtefactType) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artefacttype.GongGetReferenceIdentifier(stage))
	return
}

func (artefacttypeshape *ArtefactTypeShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artefacttypeshape.GongGetReferenceIdentifier(stage))
	return
}

func (artist *Artist) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artist.GongGetReferenceIdentifier(stage))
	return
}

func (artistshape *ArtistShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", artistshape.GongGetReferenceIdentifier(stage))
	return
}

func (controlpointshape *ControlPointShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetReferenceIdentifier(stage))
	return
}

func (desk *Desk) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", desk.GongGetReferenceIdentifier(stage))
	return
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetReferenceIdentifier(stage))
	return
}

func (influence *Influence) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influence.GongGetReferenceIdentifier(stage))
	return
}

func (influenceshape *InfluenceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influenceshape.GongGetReferenceIdentifier(stage))
	return
}

func (movement *Movement) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", movement.GongGetReferenceIdentifier(stage))
	return
}

func (movementshape *MovementShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", movementshape.GongGetReferenceIdentifier(stage))
	return
}

func (place *Place) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", place.GongGetReferenceIdentifier(stage))
	return
}

// end of template
