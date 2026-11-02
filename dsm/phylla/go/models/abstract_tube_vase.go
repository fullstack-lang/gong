package models

type TubeVaseAbstract struct {
	Name string

	// Z_Ribbon is the Z coordinate for the start of the 3D Ribbon (for Vase Trapeze).
	// The tiled floor will be on Z=0.
	Z_Ribbon float64

	// RibbonVerticalScale is the vertical scale applied to the 3D ribbon with origin at Z_Ribbon (for Vase Trapeze).
	RibbonVerticalScale float64

	// TopPlaneHeight is the height of the first projection plane relative to Z_Ribbon (for Vase Trapeze).
	TopPlaneHeight float64

	// BottomPlaneHeight is the height of the second projection plane relative to Z_Ribbon (for Vase Trapeze).
	BottomPlaneHeight float64

	// BaseRingBottomPlaneHeight is the height of the bottom projection plane for the base ring (h == 0) relative to Z_Ribbon (for Vase Trapeze).
	BaseRingBottomPlaneHeight float64

	// ProjectionAngle is the constant projection angle in degrees (for Vase Trapeze).
	ProjectionAngle float64

	// BasePlateHeight is the height of the plancher (base plate) for the vase wall (for Vase Trapeze).
	BasePlateHeight float64

	// CarvedOutTopRingsParameter is the depth ratio [0.0, 1.0] for the carved out top curves (for Vase Trapeze).
	// If 0.0, the carved out top curve is the same as the original top curve.
	// If 1.0, the carved out top curve reaches the bottom curve at the minimum radius.
	CarvedOutTopRingsParameter float64

	// BulbousStartTangentMagnitude is the magnitude of the tangent vector to the bulbous base curve at the start (extending from the ring).
	BulbousStartTangentMagnitude float64

	// BulbousEndAngle is the angle in degrees relative to the vertical of the tangent vector at the bottom of the bulbous base curve.
	// Default value: 30 degrees.
	BulbousEndAngle float64

	// BulbousEndTangentMagnitude is the magnitude of the tangent vector to the bulbous base curve at the end (at the floor).
	BulbousEndTangentMagnitude float64

	// HasConstantThickness when true adjusts the 3D ribbon curves (for Vase Trapeze) so that
	// the projected 3D rings have constant 3D wall thickness and surface-to-surface rotated contact.
	// When false, the original variable thickness from the 2D growth curves is used.
	HasConstantThickness bool

	// RelativeVerticalThickness of the growth curve. when growth curve are stacked, each is separate from the next
	// the vertical thickness is RelativeVerticalThickness x RhombusSideLength
	RelativeVerticalThickness float64

	// thickness allong the radius
	// radial with of the solid torus is RelativeRadialThickness x RhombusSideLength
	RelativeRadialThickness float64

	// For laser cutting the torus forms, they will be stack on top of another, without the rotation
	// RelativeCuttedStackFloorHeight x RhombusSideLength is the distance between each ribbon
	RelativeCuttedStackFloorHeight float64

	// 3D rotated torus stack on top of another, with 	 rotation
	// RelativeRotatedTorusSeparation x RhombusSideLength is an additionnal distance between each ribbon
	RelativeRotatedTorusSeparation float64

	// RotationRatio is the ratio of rotation between 0.0 and 1.0.
	// 0.0 means no rotation
	// 1.0 means the rotation of 2xPi x (GrowthVectorShape.X /  PlantCircumferenceShape.Length)
	RotationRatio float64

	heightAtRotRatio0 float64
	heightAtRotRatio1 float64

	// RadialRepetitions impacts the conversion from 2D ribbons to 3D torus.
	// The 3D torus  is construed by 2Pi/RadialRepetitions each time it ranges over the  GrowthCurve2DRibbon
	// If ThreeDModule is 1, the 3D torus wraps 360 degrees while ranging over the GrowthCurve2DRibbon
	RadialRepetitions int

	// Transparency of the 3D torus material (between 0.0 for opaque and 1.0 for transparent)
	Transparency float64

	// ColorCombination selects one of 10 designer color combinations (from 1 to 10) for the rings in the stack
	ColorCombination int

	// RelativeTrajectoryOffsetX x PlantCircumferenceShape.Length
	// RelativeTrajectoryOffsetY x PlantCircumferenceShape.Length
	// is taken into account for drawing the PartiallyGrowthCurve2DTrajectory
	RelativeTrajectoryOffsetX, RelativeTrajectoryOffsetY float64

	// NbStepP1P2 is the number of interpolation points for P1, P2 pairs
	NbStepP1P2 int

	// ChosenStep is
	ChosenStep int

	// RelativeHorizontalRingsHeight x RhombusSideLength is the manual height of the bottom/top horizontal rings.
	// If 0.0, the height automatically matches the wave amplitude of the bottom/top ring.
	RelativeHorizontalRingsHeight float64

	OffsetKeyX float64
	OffsetKeyY float64
	HeightKey  float64
	WidthKey   float64

	// 1.0 is full hey hole occupency
	RelativeKeySize float64

	// MovieNbFrames is the number of frames for movie recording.
	// The rotation increment per frame is 1.0 / MovieNbFrames.
	MovieNbFrames int

	PerpendicularVectorGridHalfway *PerpendicularVectorGridHalfway

	// used to construe the xxxHalfwayArcShapeGrid (x)
	TopStartArcShapeGrid *TopStartArcShapeGrid
	TopEndArcShapeGrid   *TopEndArcShapeGrid

	ShiftedBottomTopStartArcShapeGrid *ShiftedBottomTopStartArcShapeGrid
	TopMidArcVectorShapeGrid          *TopMidArcVectorShapeGrid

	StartHalfwayArcShapeGrid    *StartHalfwayArcShapeGrid
	TopStartHalfwayArcShapeGrid *TopStartHalfwayArcShapeGrid
	EndHalfwayArcShapeGrid      *EndHalfwayArcShapeGrid
	TopEndHalfwayArcShapeGrid   *TopEndHalfwayArcShapeGrid

	StackOfRotatedGrowthCurve2D    *StackOfRotatedGrowthCurve2D
	TopStackOfRotatedGrowthCurve2D *TopStackOfRotatedGrowthCurve2D

	TopGrowthCurve2D *TopGrowthCurve2D

	StackOfGrowthCurve2D    *StackOfGrowthCurve2D
	TopStackOfGrowthCurve2D *TopStackOfGrowthCurve2D

	StackOfGrowthCurve2DRibbon        *StackOfGrowthCurve2DRibbon
	StackOfRotatedGrowthCurve2DRibbon *StackOfRotatedGrowthCurve2DRibbon
	GrowthCurve2DRibbon               *GrowthCurve2DRibbon
	ShiftedRightGrowthCurve2DRibbon   *ShiftedRightGrowthCurve2DRibbon
	ShiftedLeftGrowthCurve2DRibbon    *ShiftedLeftGrowthCurve2DRibbon

	PartiallyGrowthCurve2DRibbon            *PartiallyGrowthCurve2DRibbon
	ShiftedLeftPartiallyGrowthCurve2DRibbon *ShiftedLeftPartiallyGrowthCurve2DRibbon
	PartiallyGrowthCurve2DTrajectory        *PartiallyGrowthCurve2DTrajectory
	PartiallyGrowthCurve2DTrajectoryP1P2    *PartiallyGrowthCurve2DTrajectoryP1P2
	PxShape                                 *PxShape
	ChosenP1P2PairShape                     *ChosenP1P2PairShape
	KeyHoleShape                            *KeyHoleShape
}

// VaseColorPalettes defines 10 designer color combinations (1-indexed: 1 to 10),
// mixing homogeneous palettes (white, black, Hella Jongerius tones) and
// heterogeneous palettes (Hella Jongerius multi-tone and bright crystal/glass colors).
var VaseColorPalettes = [10][]string{
	// 1. Porcelain White (Homogeneous) — warm gallery bisque porcelain
	{"#f4f1ea"},
	// 2. Anthracite Black (Homogeneous) — satin architectural charcoal/black
	{"#4a4e54"},
	// 3. Jongerius Celadon (Homogeneous) — Hella Jongerius mineral celadon green
	{"#9caf9e"},
	// 4. Jongerius Terracotta (Homogeneous) — Hella Jongerius warm oxblood/terracotta red
	{"#b84a39"},
	// 5. Jongerius Red & White Vase (Heterogeneous 2-tone) — porcelain white & vermilion terracotta
	{"#f8f6f0", "#c83e2b"},
	// 6. Jongerius Coloured Vases (Heterogeneous 4-tone) — chalk white, ochre amber, cobalt blue, celadon sage
	{"#f4efe6", "#d48c46", "#3b5978", "#8a9a86"},
	// 7. Jongerius Breathing Colour (Heterogeneous 4-tone) — petrol blue, dusty coral rose, mustard ochre, oxide green
	{"#2b3a4a", "#e07a5f", "#e9c46a", "#457b63"},
	// 8. Murano Crystal Jewel Glass (Heterogeneous 5-tone) — bright lagoon cyan, rubino fuchsia, amber citrine, amethyst, emerald aqua
	{"#00b4d8", "#ff007f", "#ffbe0b", "#7209b7", "#06d6a0"},
	// 9. Nordic Bright Art Glass (Heterogeneous 4-tone) — cobalt crystal blue, mandarin orange, cranberry rose, chartreuse green
	{"#0077b6", "#fb5607", "#ff006e", "#8ac926"},
	// 10. Bauhaus / Primary Pop Crystal (Heterogeneous 4-tone) — primary red, ultramarine blue, cadmium yellow, porcelain white
	{"#e63946", "#1d3557", "#ffb703", "#f8f6f0"},
}

// GetVaseRingColor returns the hex color for a given 1-based colorCombination (1..10) and ring index h.
func GetVaseRingColor(colorCombination int, ringIndex int) string {
	idx := colorCombination - 1
	if idx < 0 || idx >= len(VaseColorPalettes) {
		idx = 0
	}
	palette := VaseColorPalettes[idx]
	if ringIndex < 0 {
		ringIndex = 0
	}
	return palette[ringIndex%len(palette)]
}
