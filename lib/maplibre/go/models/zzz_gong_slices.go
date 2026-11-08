// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
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
	// Compute reverse map for named struct Layer
	// insertion point per field
	stage.Layer_Polygons_reverseMap = make(map[*Polygon]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _polygon := range layer.Polygons {
			stage.Layer_Polygons_reverseMap[_polygon] = layer
		}
	}
	stage.Layer_Polylines_reverseMap = make(map[*Polyline]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _polyline := range layer.Polylines {
			stage.Layer_Polylines_reverseMap[_polyline] = layer
		}
	}
	stage.Layer_Markers_reverseMap = make(map[*Marker]*Layer)
	for layer := range stage.Layers {
		_ = layer
		for _, _marker := range layer.Markers {
			stage.Layer_Markers_reverseMap[_marker] = layer
		}
	}

	// Compute reverse map for named struct MapView
	// insertion point per field
	stage.MapView_Layers_reverseMap = make(map[*Layer]*MapView)
	for mapview := range stage.MapViews {
		_ = mapview
		for _, _layer := range mapview.Layers {
			stage.MapView_Layers_reverseMap[_layer] = mapview
		}
	}

	// Compute reverse map for named struct Polygon
	// insertion point per field
	stage.Polygon_Coordinates_reverseMap = make(map[*Coordinate]*Polygon)
	for polygon := range stage.Polygons {
		_ = polygon
		for _, _coordinate := range polygon.Coordinates {
			stage.Polygon_Coordinates_reverseMap[_coordinate] = polygon
		}
	}

	// Compute reverse map for named struct Polyline
	// insertion point per field
	stage.Polyline_Coordinates_reverseMap = make(map[*Coordinate]*Polyline)
	for polyline := range stage.Polylines {
		_ = polyline
		for _, _coordinate := range polyline.Coordinates {
			stage.Polyline_Coordinates_reverseMap[_coordinate] = polyline
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Coordinates)

	res = __gong__appendInstances(res, stage.Layers)

	res = __gong__appendInstances(res, stage.MapViews)

	res = __gong__appendInstances(res, stage.Markers)

	res = __gong__appendInstances(res, stage.Polygons)

	res = __gong__appendInstances(res, stage.Polylines)

	return
}

// insertion point per named struct
func (coordinate *Coordinate) GongCopy() GongstructIF {
	newInstance := new(Coordinate)
	coordinate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (layer *Layer) GongCopy() GongstructIF {
	newInstance := new(Layer)
	layer.GongCopyBasicFields(newInstance)
	return newInstance
}

func (mapview *MapView) GongCopy() GongstructIF {
	newInstance := new(MapView)
	mapview.GongCopyBasicFields(newInstance)
	return newInstance
}

func (marker *Marker) GongCopy() GongstructIF {
	newInstance := new(Marker)
	marker.GongCopyBasicFields(newInstance)
	return newInstance
}

func (polygon *Polygon) GongCopy() GongstructIF {
	newInstance := new(Polygon)
	polygon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (polyline *Polyline) GongCopy() GongstructIF {
	newInstance := new(Polyline)
	polyline.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (coordinate *Coordinate) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coordinate)
}

func (layer *Layer) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, layer)
}

func (mapview *MapView) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, mapview)
}

func (marker *Marker) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, marker)
}

func (polygon *Polygon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, polygon)
}

func (polyline *Polyline) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, polyline)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.Coordinates,
		stage.Coordinate_stagedOrder,
		stage.Coordinates_reference,
		&stage.Coordinates_referenceOrder,
		stage.Coordinates_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Layers,
		stage.Layer_stagedOrder,
		stage.Layers_reference,
		&stage.Layers_referenceOrder,
		stage.Layers_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MapViews,
		stage.MapView_stagedOrder,
		stage.MapViews_reference,
		&stage.MapViews_referenceOrder,
		stage.MapViews_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Markers,
		stage.Marker_stagedOrder,
		stage.Markers_reference,
		&stage.Markers_referenceOrder,
		stage.Markers_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Polygons,
		stage.Polygon_stagedOrder,
		stage.Polygons_reference,
		&stage.Polygons_referenceOrder,
		stage.Polygons_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Polylines,
		stage.Polyline_stagedOrder,
		stage.Polylines_reference,
		&stage.Polylines_referenceOrder,
		stage.Polylines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.Coordinates, &stage.Coordinates_reference, &stage.Coordinates_referenceOrder, &stage.Coordinates_instance)

	__gong__computeReferencePass1(stage, stage.Layers, &stage.Layers_reference, &stage.Layers_referenceOrder, &stage.Layers_instance)

	__gong__computeReferencePass1(stage, stage.MapViews, &stage.MapViews_reference, &stage.MapViews_referenceOrder, &stage.MapViews_instance)

	__gong__computeReferencePass1(stage, stage.Markers, &stage.Markers_reference, &stage.Markers_referenceOrder, &stage.Markers_instance)

	__gong__computeReferencePass1(stage, stage.Polygons, &stage.Polygons_reference, &stage.Polygons_referenceOrder, &stage.Polygons_instance)

	__gong__computeReferencePass1(stage, stage.Polylines, &stage.Polylines_reference, &stage.Polylines_referenceOrder, &stage.Polylines_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Coordinates, stage.Coordinates_reference, stage)

	__gong__computeReferencePass2(stage.Layers, stage.Layers_reference, stage)

	__gong__computeReferencePass2(stage.MapViews, stage.MapViews_reference, stage)

	__gong__computeReferencePass2(stage.Markers, stage.Markers_reference, stage)

	__gong__computeReferencePass2(stage.Polygons, stage.Polygons_reference, stage)

	__gong__computeReferencePass2(stage.Polylines, stage.Polylines_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (coordinate *Coordinate) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Coordinate_stagedOrder, stage.Coordinates_referenceOrder, coordinate, "Coordinate")
}

func (layer *Layer) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Layer_stagedOrder, stage.Layers_referenceOrder, layer, "Layer")
}

func (mapview *MapView) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MapView_stagedOrder, stage.MapViews_referenceOrder, mapview, "MapView")
}

func (marker *Marker) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Marker_stagedOrder, stage.Markers_referenceOrder, marker, "Marker")
}

func (polygon *Polygon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Polygon_stagedOrder, stage.Polygons_referenceOrder, polygon, "Polygon")
}

func (polyline *Polyline) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Polyline_stagedOrder, stage.Polylines_referenceOrder, polyline, "Polyline")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (coordinate *Coordinate) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coordinate, coordinate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coordinate *Coordinate) GongGetReferenceIdentifier(stage *Stage) string {
	return coordinate.GongGetIdentifier(stage)
}

func (layer *Layer) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(layer, layer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (layer *Layer) GongGetReferenceIdentifier(stage *Stage) string {
	return layer.GongGetIdentifier(stage)
}

func (mapview *MapView) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(mapview, mapview.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mapview *MapView) GongGetReferenceIdentifier(stage *Stage) string {
	return mapview.GongGetIdentifier(stage)
}

func (marker *Marker) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(marker, marker.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (marker *Marker) GongGetReferenceIdentifier(stage *Stage) string {
	return marker.GongGetIdentifier(stage)
}

func (polygon *Polygon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(polygon, polygon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (polygon *Polygon) GongGetReferenceIdentifier(stage *Stage) string {
	return polygon.GongGetIdentifier(stage)
}

func (polyline *Polyline) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(polyline, polyline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (polyline *Polyline) GongGetReferenceIdentifier(stage *Stage) string {
	return polyline.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (coordinate *Coordinate) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coordinate.GongGetIdentifier(stage), "Coordinate", coordinate.Name)
}

func (layer *Layer) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(layer.GongGetIdentifier(stage), "Layer", layer.Name)
}

func (mapview *MapView) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(mapview.GongGetIdentifier(stage), "MapView", mapview.Name)
}

func (marker *Marker) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(marker.GongGetIdentifier(stage), "Marker", marker.Name)
}

func (polygon *Polygon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(polygon.GongGetIdentifier(stage), "Polygon", polygon.Name)
}

func (polyline *Polyline) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(polyline.GongGetIdentifier(stage), "Polyline", polyline.Name)
}

// insertion point for unstaging
func (coordinate *Coordinate) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coordinate.GongGetReferenceIdentifier(stage))
}

func (layer *Layer) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(layer.GongGetReferenceIdentifier(stage))
}

func (mapview *MapView) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(mapview.GongGetReferenceIdentifier(stage))
}

func (marker *Marker) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(marker.GongGetReferenceIdentifier(stage))
}

func (polygon *Polygon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(polygon.GongGetReferenceIdentifier(stage))
}

func (polyline *Polyline) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(polyline.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__models__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
