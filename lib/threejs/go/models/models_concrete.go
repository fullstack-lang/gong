package models

type Canvas struct {
	Name              string
	DirectionalLights []*DirectionalLight
	AmbiantLight      *AmbiantLight

	Meshs []*Mesh
}

type Position struct {
	X float64
	Y float64
	Z float64
}

type LightAbstract struct {
	Intensity float64
}

type DirectionalLight struct {
	Name string

	Position

	LightAbstract

	IsWithCastShadow bool
}

type AmbiantLight struct {
	Name string

	LightAbstract
}

type MeshMaterialAbstract struct {
	Color string // Color of the material, e.g., "hotpink", "#ff0000". Default is white.
}

type MeshMaterialBasic struct {
	Name string

	MeshMaterialAbstract
}

type Mesh struct {
	Name string

	Position

	MeshMaterialBasic    *MeshMaterialBasic
	MeshPhysicalMaterial *MeshPhysicalMaterial

	// one pointer per type of Geometry
	CylinderGeometry *CylinderGeometry
	BoxGeometry      *BoxGeometry
	SphereGeometry   *SphereGeometry
	TorusGeometry    *TorusGeometry
	PlaneGeometry    *PlaneGeometry
	TubeGeometry     *TubeGeometry
	ExtrudeGeometry  *ExtrudeGeometry
}

type CylinderGeometry struct {
	Name string

	RadiusTop      float64 // Radius of the cylinder at the top. Default is 1.
	RadiusBottom   float64 // Radius of the cylinder at the bottom. Default is 1.
	Height         float64 // Height of the cylinder. Default is 1.
	RadialSegments int     // Number of segmented faces around the circumference. Default is 32
	HeightSegments int     // Number of rows of faces along the height. Default is 1.
	OpenEnded      bool    // A Boolean indicating whether the ends of the cylinder are open or capped. Default is false
	ThetaStart     float64 // Start angle for first segment, default = 0 (three o'clock figure).
	ThetaLength    float64 // The central angle, often called theta, of the circular sector. Default is 2*Pi
}

type MeshPhysicalMaterial struct {
	Name string

	MeshMaterialAbstract

	Wireframe   bool    // Render geometry as wireframe. Default is false (i.e. render as flat polygons).
	Opacity     float64 // Float in the range of 0.0 - 1.0 indicating how transparent the material is. Default is 1.0.
	Transparent bool    // Defines whether this material is transparent. Default is false.
	Visible     bool    // Defines whether this material is visible. Default is true.
}

type BoxGeometry struct {
	Name string

	Width          float64 // Width; that is, the length of the edges parallel to the X axis. Default is 1.
	Height         float64 // Height; that is, the length of the edges parallel to the Y axis. Default is 1.
	Depth          float64 // Depth; that is, the length of the edges parallel to the Z axis. Default is 1.
	WidthSegments  int     // Number of segmented rectangular faces along the width of the sides. Default is 1.
	HeightSegments int     // Number of segmented rectangular faces along the height of the sides. Default is 1.
	DepthSegments  int     // Number of segmented rectangular faces along the depth of the sides. Default is 1.
}

type SphereGeometry struct {
	Name string

	Radius         float64 // sphere radius. Default is 1.
	WidthSegments  int     // number of horizontal segments. Minimum value is 3, and the default is 32.
	HeightSegments int     // number of vertical segments. Minimum value is 2, and the default is 16.
	PhiStart       float64 // specify horizontal starting angle. Default is 0.
	PhiLength      float64 // specify horizontal sweep angle size. Default is Math.PI * 2.
	ThetaStart     float64 // specify vertical starting angle. Default is 0.
	ThetaLength    float64 // specify vertical sweep angle size. Default is Math.PI.
}

type TorusGeometry struct {
	Name string

	Radius          float64 // Radius of the torus, from the center of the torus to the center of the tube. Default is 1.
	Tube            float64 // Radius of the tube. Default is 0.4.
	RadialSegments  int     // Default is 12
	TubularSegments int     // Default is 48.
	Arc             float64 // Central angle. Default is Math.PI * 2.
}

type PlaneGeometry struct {
	Name string

	Width          float64 // Width along the X axis. Default is 1.
	Height         float64 // Height along the Y axis. Default is 1.
	WidthSegments  int     // Number of segmented rectangular faces along the width. Default is 1.
	HeightSegments int     // Number of segmented rectangular faces along the height. Default is 1.
}

type TubeGeometry struct {
	Name string

	Path *Curve

	TubularSegments int     // The number of segments that make up the tube. Default is 64.
	Radius          float64 // The radius of the tube. Default is 1.
	RadialSegments  int     // The number of segments that make up the cross-section. Default is 8.
	Closed          bool    // Is the tube open or closed. Default is false.
}

type Vector3 struct {
	Name string
	X, Y, Z float64
}

type Curve struct {
	Name string
	Points []*Vector3
}

type ExtrudeGeometry struct {
	Name string

	Shape       *Shape
	ExtrudePath *Curve
	Steps       int
}

type Vector2 struct {
	Name string
	X, Y float64
}

type Shape struct {
	Name string
	Points []*Vector2
}
