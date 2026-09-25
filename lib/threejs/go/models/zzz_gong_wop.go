// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AmbiantLight_WOP struct {
	// insertion point

	Name string

	Intensity float64
}

func (from *AmbiantLight) GongCopyBasicFields(to *AmbiantLight) {
	// insertion point
	*to = *from
}

type BoxGeometry_WOP struct {
	// insertion point

	Name string

	Width float64

	Height float64

	Depth float64

	WidthSegments int

	HeightSegments int

	DepthSegments int
}

func (from *BoxGeometry) GongCopyBasicFields(to *BoxGeometry) {
	// insertion point
	*to = *from
}

type BufferGeometry_WOP struct {
	// insertion point

	Name string
}

func (from *BufferGeometry) GongCopyBasicFields(to *BufferGeometry) {
	// insertion point
	to.Name = from.Name
}

type Camera_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Z float64

	TargetX float64

	TargetY float64

	TargetZ float64

	Fov float64
}

func (from *Camera) GongCopyBasicFields(to *Camera) {
	// insertion point
	*to = *from
}

type Canvas_WOP struct {
	// insertion point

	Name string

	IsWithLastRenderingUpdate bool

	LastRendering time.Time

	Frame64BitsEncoded string
}

func (from *Canvas) GongCopyBasicFields(to *Canvas) {
	// insertion point
	to.Name = from.Name
	to.IsWithLastRenderingUpdate = from.IsWithLastRenderingUpdate
	to.LastRendering = from.LastRendering
	to.Frame64BitsEncoded = from.Frame64BitsEncoded
}

type Curve_WOP struct {
	// insertion point

	Name string
}

func (from *Curve) GongCopyBasicFields(to *Curve) {
	// insertion point
	to.Name = from.Name
}

type CylinderGeometry_WOP struct {
	// insertion point

	Name string

	RadiusTop float64

	RadiusBottom float64

	Height float64

	RadialSegments int

	HeightSegments int

	OpenEnded bool

	ThetaStart float64

	ThetaLength float64
}

func (from *CylinderGeometry) GongCopyBasicFields(to *CylinderGeometry) {
	// insertion point
	*to = *from
}

type DirectionalLight_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Z float64

	Intensity float64

	IsWithCastShadow bool
}

func (from *DirectionalLight) GongCopyBasicFields(to *DirectionalLight) {
	// insertion point
	*to = *from
}

type ExtrudeGeometry_WOP struct {
	// insertion point

	Name string

	Steps int
}

func (from *ExtrudeGeometry) GongCopyBasicFields(to *ExtrudeGeometry) {
	// insertion point
	to.Name = from.Name
	to.Steps = from.Steps
}

type Mesh_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Z float64
}

func (from *Mesh) GongCopyBasicFields(to *Mesh) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Z = from.Z
}

type MeshMaterialBasic_WOP struct {
	// insertion point

	Name string

	Color string
}

func (from *MeshMaterialBasic) GongCopyBasicFields(to *MeshMaterialBasic) {
	// insertion point
	*to = *from
}

type MeshPhysicalMaterial_WOP struct {
	// insertion point

	Name string

	Color string

	Wireframe bool

	Opacity float64

	Transparent bool

	Visible bool
}

func (from *MeshPhysicalMaterial) GongCopyBasicFields(to *MeshPhysicalMaterial) {
	// insertion point
	*to = *from
}

type PlaneGeometry_WOP struct {
	// insertion point

	Name string

	Width float64

	Height float64

	WidthSegments int

	HeightSegments int
}

func (from *PlaneGeometry) GongCopyBasicFields(to *PlaneGeometry) {
	// insertion point
	*to = *from
}

type Shape_WOP struct {
	// insertion point

	Name string
}

func (from *Shape) GongCopyBasicFields(to *Shape) {
	// insertion point
	to.Name = from.Name
}

type SphereGeometry_WOP struct {
	// insertion point

	Name string

	Radius float64

	WidthSegments int

	HeightSegments int

	PhiStart float64

	PhiLength float64

	ThetaStart float64

	ThetaLength float64
}

func (from *SphereGeometry) GongCopyBasicFields(to *SphereGeometry) {
	// insertion point
	*to = *from
}

type TorusGeometry_WOP struct {
	// insertion point

	Name string

	Radius float64

	Tube float64

	RadialSegments int

	TubularSegments int

	Arc float64
}

func (from *TorusGeometry) GongCopyBasicFields(to *TorusGeometry) {
	// insertion point
	*to = *from
}

type Triangle_WOP struct {
	// insertion point

	Name string

	V1 int

	V2 int

	V3 int
}

func (from *Triangle) GongCopyBasicFields(to *Triangle) {
	// insertion point
	*to = *from
}

type TubeGeometry_WOP struct {
	// insertion point

	Name string

	TubularSegments int

	Radius float64

	RadialSegments int

	Closed bool
}

func (from *TubeGeometry) GongCopyBasicFields(to *TubeGeometry) {
	// insertion point
	to.Name = from.Name
	to.TubularSegments = from.TubularSegments
	to.Radius = from.Radius
	to.RadialSegments = from.RadialSegments
	to.Closed = from.Closed
}

type Vector2_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64
}

func (from *Vector2) GongCopyBasicFields(to *Vector2) {
	// insertion point
	*to = *from
}

type Vector3_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Z float64
}

func (from *Vector3) GongCopyBasicFields(to *Vector3) {
	// insertion point
	*to = *from
}

// end of insertion point
