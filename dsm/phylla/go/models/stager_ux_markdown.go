package models

import (
	"fmt"
	"math"
	"strings"

	markdown "github.com/fullstack-lang/gong/lib/markdown/go/models"
)

func (stager *Stager) ux_markdown() {

	stage := stager.markdownStage
	stage.Reset()

	// 1. Stage the SVG illustrations
	stageSvgParastichies(stage)
	stageSvgGoldenAngle(stage)
	stageSvgCylinderToLattice(stage)

	// 2. Stage the formatted Markdown content
	content := markdown.Content{
		Name: "About Spiral Plants",
		Content: `# About Spiral Plants

## 1. Spiral Phyllotaxis

Observation of sunflower heads, pinecone scales, or pineapple skins reveals one of nature's prominent geometric phenomena: **Spiral Phyllotaxis** (from Ancient Greek *phýllon* meaning "leaf" and *táxis* meaning "arrangement").

In biological systems, organ primordia (leaves, petals, florets, or seeds) do not emerge at random. Instead, they organize into interlocking spiral trajectories radiating from the apical growth center (*shoot apical meristem*) of the plant stem.

---

## 2. Fibonacci Parastichy Numbers ($N$ and $M$)

Examination of a spiral plant head reveals two opposing families of spiral curves winding around the central axis:
- One family turns **clockwise** (illustrated in blue below).
- The opposing family turns **counter-clockwise** (illustrated in orange below).

![Phyllotaxis Parastichies in Spiral Plants](svg:spiral_parastichies?width=740px)

Counting the number of spirals in each direction consistently produces two consecutive numbers from the **Fibonacci Sequence**:

$$\mathbf{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, \dots}$$

Typical biological configurations include:
- **Pinecone:** $(N, M) = (8, 13)$
- **Pineapple:** $(N, M) = (8, 13)$ and $(13, 21)$
- **Sunflower Head:** $(N, M) = (34, 55)$ or $(55, 89)$

In botanical mathematics, these spiral pairs are designated as **contact parastichies**, represented by $(N, M)$ with $N < M$. Consecutive Fibonacci numbers guarantee optimal tangential packing density between neighboring florets.

---

## 3. The Golden Divergence Angle ($\alpha \approx 137.5^\circ$)

This geometric configuration results from growth dynamics at the shoot apical meristem. As new organ primordia are generated sequentially, each subsequent primordium is placed at a fixed divergence angle $\alpha$ relative to the preceding one.

![The Golden Angle Divergence in Plant Meristems](svg:golden_angle?width=740px)

To prevent structural overlap and maximize light interception, this angle must avoid rational fractions of a full rotation.

The optimal divergence is the **Golden Angle** ($\alpha \approx 137.5077^\circ$), derived from the Golden Ratio ($\phi = \frac{1 + \sqrt{5}}{2} \approx 1.618034$):

$$\alpha = 360^\circ \times \left(1 - \frac{1}{\phi}\right) \approx 137.507764^\circ$$

This irrational angular spacing ensures:
1. **Maximized Exposure:** Upper foliage avoids shadowing lower leaves.
2. **Optimal Packing Density:** Floret primordia fill the available area uniformly without gaps.

---

## 4. Cylindrical Surface and the Rotated Rhombus Lattice

Phyllotactic arrangements on cylindrical plant stems unfold into a periodic two-dimensional geometric lattice:

![Unrolling Cylindrical Stem to 2D Rotated Rhombus Lattice](svg:cylinder_to_lattice?width=860px)

1. **Parastichy Vectors and Unit Rhombuses:**
   - Helical trajectories on the 3D cylinder project onto straight vector directions in 2D: vector $\mathbf{v}_1$ ($N$-parastichy) and vector $\mathbf{v}_2$ ($M$-parastichy).
   - Each rhombus unit cell is parameterized by:
     - **Inside Angle ($\alpha$, <b>RhombusInsideAngle</b>):** Governs spiral steepness and divergence pitch.
     - **Side Length ($L$, <b>RhombusSideLength</b>):** Sets the physical distance between adjacent primordia.
   - The geometric center of each rhombus (marked by $\mathbf{+}$) represents the spatial location of an organ primordium (leaf, seed, or floret).

2. **Circumference Closure Vector:**
   - One complete revolution around the cylinder perimeter corresponds to taking $N$ discrete steps along $\mathbf{v}_1$ followed by $M$ discrete steps along $\mathbf{v}_2$:
     $$\mathbf{C} = N \mathbf{v}_1 + M \mathbf{v}_2, \quad \|\mathbf{C}\| = 2\pi R$$
   - In Phylla, the lattice frame is rotated by $-\phi = -\operatorname{atan2}(C_y, C_x)$ so that the circumference vector $\mathbf{C}$ aligns with the horizontal $X$-axis (**green line**).

3. **Fundamental Domain and Grid Paths:**
   - **Rotated Grid Path (Dark Red Dashed Line):** Traces the $N$ steps along $\mathbf{v}_1$ and $M$ steps along $\mathbf{v}_2$ connecting the origin $(0,0)$ to the circumference endpoint $(C, 0)$, with circular markers at each intermediate vertex.
   - **Periodic Boundary (Dashed Blue Lines):** Encloses the fundamental surface domain unrolled from one complete revolution of the cylinder.
   - Continuous traversal across neighboring rhombus cells defines growth trajectories and ribbons that can be extruded into full 3D geometries.

---

## 5. Generative Design with Phylla

**Phylla** utilizes these mathematical phyllotaxis principles to construct parametric 2D diagrams and 3D models:

- **Plant 2D & 3D:** Exploration of fundamental rhombus grids, computation of $(N, M)$ contact parastichy pairs, and visualization of spiral trajectories.
- **Vases:** Parametric rotation profiles and stacked ring transformations producing organic, twisting 3D vase forms.
- **Stools & Furniture:** Phyllotactic spiral vectors mapped to structural lattice legs and self-supporting geometry.
- **Clocks:** Spiral distribution of hour and minute indicators for balanced dial layouts.
- **Additive Manufacturing (STL Export):** Direct export of generated 3D geometry to STL format via the **Download STL** action in the tree menu.
`,
	}
	content.Stage(stage)

	stage.Commit()
}

func stageSvgParastichies(stage *markdown.Stage) {
	cx, cy := 180.0, 195.0
	c := 13.5
	numPts := 95

	type pt struct {
		x, y float64
	}
	pts := make(map[int]pt)
	for k := 1; k <= numPts; k++ {
		r := c * math.Sqrt(float64(k))
		theta := float64(k) * 137.507764 * math.Pi / 180.0
		pts[k] = pt{
			x: cx + r*math.Cos(theta),
			y: cy + r*math.Sin(theta),
		}
	}

	var sb strings.Builder
	sb.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 780 400" width="100%" height="auto">`)
	sb.WriteString(`<rect width="780" height="400" rx="14" fill="#f8fafc" stroke="#cbd5e1" stroke-width="1.5"/>`)

	// Title
	sb.WriteString(`<text x="390" y="34" font-family="system-ui, sans-serif" font-size="18" font-weight="bold" fill="#0f172a" text-anchor="middle">Sunflower Head Phyllotaxis: Contact Parastichies (8, 13)</text>`)

	// Draw clockwise spirals (step 8) in Blue
	for start := 1; start <= 8; start++ {
		var pathParts []string
		for k := start; k <= numPts; k += 8 {
			if p, ok := pts[k]; ok {
				if len(pathParts) == 0 {
					pathParts = append(pathParts, fmt.Sprintf("M %.1f %.1f", p.x, p.y))
				} else {
					pathParts = append(pathParts, fmt.Sprintf("L %.1f %.1f", p.x, p.y))
				}
			}
		}
		if len(pathParts) > 1 {
			sb.WriteString(fmt.Sprintf(`<path d="%s" fill="none" stroke="#2563eb" stroke-width="2.6" stroke-opacity="0.85" stroke-linecap="round"/>`, strings.Join(pathParts, " ")))
		}
	}

	// Draw counter-clockwise spirals (step 13) in Orange
	for start := 1; start <= 13; start++ {
		var pathParts []string
		for k := start; k <= numPts; k += 13 {
			if p, ok := pts[k]; ok {
				if len(pathParts) == 0 {
					pathParts = append(pathParts, fmt.Sprintf("M %.1f %.1f", p.x, p.y))
				} else {
					pathParts = append(pathParts, fmt.Sprintf("L %.1f %.1f", p.x, p.y))
				}
			}
		}
		if len(pathParts) > 1 {
			sb.WriteString(fmt.Sprintf(`<path d="%s" fill="none" stroke="#ea580c" stroke-width="2.6" stroke-opacity="0.85" stroke-linecap="round"/>`, strings.Join(pathParts, " ")))
		}
	}

	// Draw seeds / florets
	for k := 1; k <= numPts; k++ {
		p := pts[k]
		sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="4.0" fill="#1e293b" stroke="#ffffff" stroke-width="1.0"/>`, p.x, p.y))
	}

	// Legend Box on the right
	sb.WriteString(`
	<g transform="translate(410, 75)">
		<rect width="345" height="280" rx="10" fill="#ffffff" stroke="#cbd5e1" stroke-width="1.5"/>
		<text x="22" y="34" font-family="system-ui, sans-serif" font-size="16" font-weight="bold" fill="#0f172a">Parastichy Families</text>
		
		<!-- Blue Spiral Legend -->
		<line x1="22" y1="72" x2="68" y2="72" stroke="#2563eb" stroke-width="3.5" stroke-linecap="round"/>
		<text x="80" y="78" font-family="system-ui, sans-serif" font-size="15" font-weight="bold" fill="#1e293b">8 Clockwise Spirals (N = 8)</text>
		<text x="80" y="98" font-family="system-ui, sans-serif" font-size="13" fill="#64748b">Steps along primordium index +8</text>

		<!-- Orange Spiral Legend -->
		<line x1="22" y1="140" x2="68" y2="140" stroke="#ea580c" stroke-width="3.5" stroke-linecap="round"/>
		<text x="80" y="146" font-family="system-ui, sans-serif" font-size="15" font-weight="bold" fill="#1e293b">13 Counter-Clockwise (M = 13)</text>
		<text x="80" y="166" font-family="system-ui, sans-serif" font-size="13" fill="#64748b">Steps along primordium index +13</text>

		<!-- Fibonacci Pair Annotation -->
		<rect x="18" y="200" width="310" height="58" rx="8" fill="#f1f5f9" stroke="#e2e8f0"/>
		<text x="30" y="224" font-family="system-ui, sans-serif" font-size="14" font-weight="bold" fill="#0f172a">Contact Parastichy Pair:</text>
		<text x="30" y="244" font-family="system-ui, sans-serif" font-size="13" font-weight="600" fill="#2563eb">(N, M) = (8, 13) (Fibonacci numbers)</text>
	</g>
	`)

	sb.WriteString(`</svg>`)

	(&markdown.SvgImage{
		Name:    "spiral_parastichies",
		Content: sb.String(),
	}).Stage(stage)
}

func stageSvgGoldenAngle(stage *markdown.Stage) {
	cx, cy := 185.0, 205.0
	rOuter := 130.0

	var sb strings.Builder
	sb.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 780 400" width="100%" height="auto">`)
	sb.WriteString(`<rect width="780" height="400" rx="14" fill="#f8fafc" stroke="#cbd5e1" stroke-width="1.5"/>`)

	// Title
	sb.WriteString(`<text x="390" y="34" font-family="system-ui, sans-serif" font-size="18" font-weight="bold" fill="#0f172a" text-anchor="middle">Shoot Apical Meristem: Primordia Divergence Angle α ≈ 137.5°</text>`)

	// Meristem Circle
	sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="%.1f" fill="#f1f5f9" stroke="#94a3b8" stroke-width="1.8" stroke-dasharray="5,5"/>`, cx, cy, rOuter))
	sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="22" fill="#e2e8f0" stroke="#64748b" stroke-width="1.8"/>`, cx, cy))
	sb.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="13" font-weight="bold" fill="#334155" text-anchor="middle">Apex</text>`, cx, cy+4))

	// Successive Primordia (0 to 5)
	type prim struct {
		num int
		deg float64
		rad float64
	}
	prims := []prim{
		{num: 0, deg: 0, rad: 110},
		{num: 1, deg: 137.5077, rad: 95},
		{num: 2, deg: 275.0154, rad: 80},
		{num: 3, deg: 52.5231, rad: 65},
		{num: 4, deg: 190.0308, rad: 50},
		{num: 5, deg: 327.5385, rad: 35},
	}

	// Reference radii lines for Node 0 and Node 1
	p0X := cx + prims[0].rad*math.Cos(prims[0].deg*math.Pi/180.0)
	p0Y := cy - prims[0].rad*math.Sin(prims[0].deg*math.Pi/180.0)
	p1X := cx + prims[1].rad*math.Cos(prims[1].deg*math.Pi/180.0)
	p1Y := cy - prims[1].rad*math.Sin(prims[1].deg*math.Pi/180.0)

	sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#94a3b8" stroke-width="1.8"/>`, cx, cy, p0X+20, p0Y))
	sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#94a3b8" stroke-width="1.8"/>`, cx, cy, p1X-15, p1Y-15))

	// Arc for divergence angle α (counter-clockwise from 0° to 137.5°)
	arcR := 54.0
	arcX0 := cx + arcR
	arcY0 := cy
	arcX1 := cx + arcR*math.Cos(137.5077*math.Pi/180.0)
	arcY1 := cy - arcR*math.Sin(137.5077*math.Pi/180.0)
	sb.WriteString(fmt.Sprintf(`<path d="M %.1f %.1f A %.1f %.1f 0 0 0 %.1f %.1f" fill="none" stroke="#2563eb" stroke-width="3.0"/>`, arcX0, arcY0, arcR, arcR, arcX1, arcY1))
	sb.WriteString(fmt.Sprintf(`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="#2563eb"/>`, arcX1+3, arcY1+12, arcX1-10, arcY1+3, arcX1-3, arcY1+14))
	sb.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="16" font-weight="bold" fill="#2563eb">α ≈ 137.5°</text>`, cx+32, cy-52))

	// Draw and label each primordium
	for _, pr := range prims {
		px := cx + pr.rad*math.Cos(pr.deg*math.Pi/180.0)
		py := cy - pr.rad*math.Sin(pr.deg*math.Pi/180.0)

		fillCol := "#10b981"
		if pr.num == 0 || pr.num == 1 {
			fillCol = "#2563eb"
		}
		sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="13" fill="%s" stroke="#ffffff" stroke-width="2.0"/>`, px, py, fillCol))
		sb.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="13" font-weight="bold" fill="#ffffff" text-anchor="middle">P%d</text>`, px, py+5, pr.num))
	}

	// Right Explanation Box
	sb.WriteString(`
	<g transform="translate(400, 75)">
		<rect width="355" height="280" rx="10" fill="#ffffff" stroke="#cbd5e1" stroke-width="1.5"/>
		<text x="22" y="34" font-family="system-ui, sans-serif" font-size="16" font-weight="bold" fill="#0f172a">Golden Divergence Principle</text>
		
		<text x="22" y="68" font-family="system-ui, sans-serif" font-size="14" fill="#334155">• Each new organ primordium (P₀, P₁, P₂...)</text>
		<text x="32" y="88" font-family="system-ui, sans-serif" font-size="14" fill="#334155">is initiated sequentially at the apex.</text>
		
		<text x="22" y="122" font-family="system-ui, sans-serif" font-size="14" fill="#334155">• Divergence angle between successive</text>
		<text x="32" y="142" font-family="system-ui, sans-serif" font-size="14" fill="#334155">primordia is strictly constant:</text>

		<rect x="22" y="160" width="310" height="46" rx="8" fill="#eff6ff" stroke="#bfdbfe"/>
		<text x="35" y="189" font-family="monospace" font-size="15" font-weight="bold" fill="#1d4ed8">α = 360° × (1 − 1/φ) ≈ 137.5°</text>

		<text x="22" y="234" font-family="system-ui, sans-serif" font-size="13" fill="#475569">Older primordia (P₀) are displaced radially</text>
		<text x="22" y="252" font-family="system-ui, sans-serif" font-size="13" fill="#475569">outward as newer ones (P₁, P₂, ...) emerge.</text>
	</g>
	`)

	sb.WriteString(`</svg>`)

	(&markdown.SvgImage{
		Name:    "golden_angle",
		Content: sb.String(),
	}).Stage(stage)
}

func stageSvgCylinderToLattice(stage *markdown.Stage) {
	var sb strings.Builder
	sb.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 960 520" width="100%" height="auto">`)
	sb.WriteString(`<rect width="960" height="520" rx="14" fill="#ffffff" stroke="#cbd5e1" stroke-width="1.5"/>`)

	// Left: 3D Plant Stem Cylinder
	sb.WriteString(`
	<g transform="translate(25, 45)">
		<text x="95" y="15" font-family="system-ui, sans-serif" font-size="16" font-weight="bold" fill="#0f172a" text-anchor="middle">3D Plant Stem Cylinder</text>
		
		<defs>
			<linearGradient id="cylGrad" x1="0" y1="0" x2="1" y2="0">
				<stop offset="0%" stop-color="#94a3b8"/>
				<stop offset="25%" stop-color="#f1f5f9"/>
				<stop offset="60%" stop-color="#e2e8f0"/>
				<stop offset="100%" stop-color="#64748b"/>
			</linearGradient>
			<marker id="arrow" viewBox="0 0 10 10" refX="6" refY="5" markerWidth="6" markerHeight="6" orient="auto-start-reverse">
				<path d="M 0 1 L 8 5 L 0 9 z" fill="#0f172a"/>
			</marker>
		</defs>

		<!-- Bottom Ellipse -->
		<ellipse cx="95" cy="285" rx="65" ry="24" fill="#cbd5e1" stroke="#64748b" stroke-width="2.2"/>
		<!-- Body rect -->
		<rect x="30" y="75" width="130" height="210" fill="url(#cylGrad)"/>
		<!-- Side lines -->
		<line x1="30" y1="75" x2="30" y2="285" stroke="#64748b" stroke-width="2.2"/>
		<line x1="160" y1="75" x2="160" y2="285" stroke="#64748b" stroke-width="2.2"/>
		<!-- Top Ellipse -->
		<ellipse cx="95" cy="75" rx="65" ry="24" fill="#f8fafc" stroke="#64748b" stroke-width="2.2"/>

		<!-- Helical parastichy curves on cylinder -->
		<!-- Blue spirals (up and right) -->
		<path d="M 30 245 Q 95 280 160 210" fill="none" stroke="#2563eb" stroke-width="3.0" stroke-linecap="round"/>
		<path d="M 30 185 Q 95 220 160 150" fill="none" stroke="#2563eb" stroke-width="3.0" stroke-linecap="round"/>
		<path d="M 30 125 Q 95 160 160 90" fill="none" stroke="#2563eb" stroke-width="3.0" stroke-linecap="round"/>
		<!-- Blue dashed back -->
		<path d="M 160 210 Q 95 185 30 185" fill="none" stroke="#93c5fd" stroke-width="2.0" stroke-dasharray="4,4"/>
		<path d="M 160 150 Q 95 125 30 125" fill="none" stroke="#93c5fd" stroke-width="2.0" stroke-dasharray="4,4"/>

		<!-- Orange spirals (up and left) -->
		<path d="M 160 245 Q 95 280 30 210" fill="none" stroke="#ea580c" stroke-width="3.0" stroke-linecap="round"/>
		<path d="M 160 185 Q 95 220 30 150" fill="none" stroke="#ea580c" stroke-width="3.0" stroke-linecap="round"/>
		<path d="M 160 125 Q 95 160 30 90" fill="none" stroke="#ea580c" stroke-width="3.0" stroke-linecap="round"/>
		<!-- Orange dashed back -->
		<path d="M 30 210 Q 95 185 160 185" fill="none" stroke="#fdba74" stroke-width="2.0" stroke-dasharray="4,4"/>
		<path d="M 30 150 Q 95 125 160 125" fill="none" stroke="#fdba74" stroke-width="2.0" stroke-dasharray="4,4"/>

		<!-- Cut line -->
		<line x1="30" y1="75" x2="30" y2="285" stroke="#db2777" stroke-width="2.5" stroke-dasharray="5,5"/>
		<rect x="5" y="165" width="56" height="24" rx="5" fill="#fdf2f8" stroke="#f472b6" stroke-width="1.2"/>
		<text x="33" y="181" font-family="system-ui, sans-serif" font-size="12" font-weight="bold" fill="#db2777" text-anchor="middle">Cut Line</text>

		<!-- Bottom Circumference Label -->
		<line x1="30" y1="330" x2="160" y2="330" stroke="#334155" stroke-width="2"/>
		<rect x="10" y="342" width="170" height="28" rx="6" fill="#f1f5f9" stroke="#cbd5e1" stroke-width="1"/>
		<text x="95" y="361" font-family="system-ui, sans-serif" font-size="14" font-weight="bold" fill="#1e293b" text-anchor="middle">Circumference C = 2πR</text>
	</g>
	`)

	// Center: Transition Arrow
	sb.WriteString(`
	<g transform="translate(205, 210)">
		<rect x="0" y="0" width="85" height="50" rx="8" fill="#eff6ff" stroke="#93c5fd" stroke-width="1.5"/>
		<text x="42" y="20" font-family="system-ui, sans-serif" font-size="12" font-weight="bold" fill="#1d4ed8" text-anchor="middle">UNROLL</text>
		<text x="42" y="36" font-family="system-ui, sans-serif" font-size="11" font-weight="600" fill="#3b82f6" text-anchor="middle">Cylinder → 2D</text>
		<path d="M 95 25 L 118 25 M 110 18 L 118 25 L 110 32" fill="none" stroke="#1d4ed8" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
	</g>
	`)

	// Right: 2D Rotated Rhombus Lattice (matching Phylla 2D diagram)
	sb.WriteString(`
	<g transform="translate(325, 45)">
		<text x="290" y="15" font-family="system-ui, sans-serif" font-size="16" font-weight="bold" fill="#0f172a" text-anchor="middle">2D Rotated Rhombus Lattice (Fundamental Domain)</text>
	`)

	ox, oy := 120.0, 290.0
	N, M := 2, 3
	insideAngleDeg := 69.0
	sideLength := 65.0

	insideAngleRad := insideAngleDeg * math.Pi / 180.0
	halfAngle := insideAngleRad / 2.0
	v1_orig_x := sideLength * math.Cos(halfAngle)
	v1_orig_y := sideLength * math.Sin(halfAngle)
	v2_orig_x := -sideLength * math.Cos(halfAngle)
	v2_orig_y := sideLength * math.Sin(halfAngle)

	Wx := float64(N)*v1_orig_x + float64(M)*v2_orig_x
	Wy := float64(N)*v1_orig_y + float64(M)*v2_orig_y
	C := math.Sqrt(Wx*Wx + Wy*Wy)
	phi := math.Atan2(Wy, Wx)

	cosRot := math.Cos(-phi)
	sinRot := math.Sin(-phi)

	v1_rot_x := v1_orig_x*cosRot - v1_orig_y*sinRot
	v1_rot_y := v1_orig_x*sinRot + v1_orig_y*cosRot
	v2_rot_x := v2_orig_x*cosRot - v2_orig_y*sinRot
	v2_rot_y := v2_orig_x*sinRot + v2_orig_y*cosRot

	// SVG coordinate conversion (SVG Y is inverted)
	v1x := v1_rot_x
	v1y := -v1_rot_y
	v2x := v2_rot_x
	v2y := -v2_rot_y

	// Draw Axes with Arrows
	sb.WriteString(fmt.Sprintf(`
		<!-- Coordinate Axes -->
		<!-- Y-Axis (Stem Vertical Axis) -->
		<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#0f172a" stroke-width="2.0" marker-end="url(#arrow)"/>
		<rect x="%.1f" y="18" width="110" height="26" rx="6" fill="#f8fafc" stroke="#94a3b8" stroke-width="1.2"/>
		<text x="%.1f" y="35" font-family="system-ui, sans-serif" font-size="13" font-weight="bold" fill="#0f172a" text-anchor="middle">Stem Axis (Y)</text>

		<!-- X-Axis (Horizontal Circumference) -->
		<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#0f172a" stroke-width="2.0" marker-end="url(#arrow)"/>
		<rect x="%.1f" y="%.1f" width="135" height="26" rx="6" fill="#f8fafc" stroke="#94a3b8" stroke-width="1.2"/>
		<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="13" font-weight="bold" fill="#0f172a" text-anchor="middle">Circumference (X)</text>
	`,
		ox, oy+155.0, ox, 48.0,
		ox-55.0,
		ox,
		ox-55.0, oy, ox+460.0, oy,
		ox+340.0, oy-32.0,
		ox+407.0, oy-15.0,
	))

	// Draw the 9 rhombuses for i in [-1, N-1] and j in [0, M-1]
	for i := -1; i < N; i++ {
		for j := 0; j < M; j++ {
			x0 := ox + float64(i)*v1x + float64(j)*v2x
			y0 := oy + float64(i)*v1y + float64(j)*v2y
			x1 := x0 + v1x
			y1 := y0 + v1y
			x2 := x0 + v1x + v2x
			y2 := y0 + v1y + v2y
			x3 := x0 + v2x
			y3 := y0 + v2y

			fillCol := "#ffffff"
			fillOp := "0.0"
			if i >= 0 {
				fillCol = "#e0f2fe"
				fillOp = "0.20"
			}

			// Rhombus polygon
			sb.WriteString(fmt.Sprintf(`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="%s" fill-opacity="%s" stroke="#3b82f6" stroke-width="1.8" stroke-opacity="0.85"/>`,
				x0, y0, x1, y1, x2, y2, x3, y3, fillCol, fillOp))

			// Center cross '+' (primordium location)
			cx := (x0 + x2) / 2.0
			cy := (y0 + y2) / 2.0
			sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#0f172a" stroke-width="1.8"/>`, cx-5.0, cy, cx+5.0, cy))
			sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#0f172a" stroke-width="1.8"/>`, cx, cy-5.0, cx, cy+5.0))
		}
	}

	// Blue dashed outer periodic boundary around the rhombus set
	pTopLeft_x := ox + float64(-1)*v1x + float64(M)*v2x
	pTopLeft_y := oy + float64(-1)*v1y + float64(M)*v2y
	pTopRight_x := ox + float64(N)*v1x + float64(M)*v2x
	pTopRight_y := oy + float64(N)*v1y + float64(M)*v2y
	pFarRight_x := ox + float64(N)*v1x
	pFarRight_y := oy + float64(N)*v1y
	pFarLeft_x := ox + float64(-1)*v1x
	pFarLeft_y := oy + float64(-1)*v1y

	sb.WriteString(fmt.Sprintf(`<polyline points="%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="none" stroke="#1d4ed8" stroke-width="2.2" stroke-dasharray="6,6"/>`,
		pFarLeft_x, pFarLeft_y, pTopLeft_x, pTopLeft_y, pTopRight_x, pTopRight_y, pFarRight_x, pFarRight_y))

	// Dark red dashed path for the fundamental domain closure (0,0) -> N*v1 -> N*v1 + M*v2 = (C, 0)
	var pathPts []string
	currX, currY := ox, oy
	pathPts = append(pathPts, fmt.Sprintf("%.1f,%.1f", currX, currY))

	type circNode struct {
		x, y float64
	}
	var nodes []circNode
	nodes = append(nodes, circNode{currX, currY})

	for k := 0; k < N; k++ {
		currX += v1x
		currY += v1y
		pathPts = append(pathPts, fmt.Sprintf("%.1f,%.1f", currX, currY))
		nodes = append(nodes, circNode{currX, currY})
	}
	for k := 0; k < M; k++ {
		currX += v2x
		currY += v2y
		pathPts = append(pathPts, fmt.Sprintf("%.1f,%.1f", currX, currY))
		nodes = append(nodes, circNode{currX, currY})
	}

	sb.WriteString(fmt.Sprintf(`<polyline points="%s" fill="none" stroke="#991b1b" stroke-width="2.8" stroke-dasharray="6,6"/>`,
		strings.Join(pathPts, " ")))

	// Green circumference line along horizontal axis from (0,0) to (C,0)
	sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#16a34a" stroke-width="4.0" stroke-linecap="round"/>`,
		ox, oy, ox+C, oy))

	// Red circle markers along the grid path
	for _, n := range nodes {
		sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="6.0" fill="#ffffff" stroke="#991b1b" stroke-width="2.2"/>`, n.x, n.y))
	}

	// High-Contrast Annotations with Pill Backgrounds
	sb.WriteString(fmt.Sprintf(`
		<!-- Annotation: Circumference vector C -->
		<rect x="%.1f" y="%.1f" width="200" height="28" rx="6" fill="#f0fdf4" stroke="#86efac" stroke-width="1.2"/>
		<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="13" font-weight="bold" fill="#15803d" text-anchor="middle">Vector C = N v₁ + M v₂</text>

		<!-- Annotation: N Parastichy directions -->
		<rect x="%.1f" y="%.1f" width="130" height="26" rx="6" fill="#fef2f2" stroke="#fca5a5" stroke-width="1.2"/>
		<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="13" font-weight="bold" fill="#991b1b" text-anchor="middle">N = 2 steps (v₁)</text>

		<!-- Annotation: M Parastichy directions -->
		<rect x="%.1f" y="%.1f" width="130" height="26" rx="6" fill="#fef2f2" stroke="#fca5a5" stroke-width="1.2"/>
		<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="13" font-weight="bold" fill="#991b1b" text-anchor="middle">M = 3 steps (v₂)</text>

		<!-- Annotation: Primordium Center '+' -->
		<rect x="360" y="145" width="155" height="26" rx="6" fill="#ffffff" stroke="#cbd5e1" stroke-width="1.2"/>
		<text x="437" y="162" font-family="system-ui, sans-serif" font-size="13" font-weight="bold" fill="#334155" text-anchor="middle">Primordium Node (+)</text>

		<!-- Annotation: Periodic Boundary -->
		<rect x="235" y="45" width="205" height="26" rx="6" fill="#eff6ff" stroke="#93c5fd" stroke-width="1.2"/>
		<text x="337" y="62" font-family="system-ui, sans-serif" font-size="13" font-weight="bold" fill="#1d4ed8" text-anchor="middle">Periodic Boundary (Dashed)</text>
	`,
		ox+20.0, oy+10.0,
		ox+120.0, oy+29.0,
		ox-105.0, oy+95.0,
		ox-40.0, oy+112.0,
		ox+120.0, oy+135.0,
		ox+185.0, oy+152.0,
	))

	sb.WriteString(`
	</g>
	</svg>`)

	(&markdown.SvgImage{
		Name:    "cylinder_to_lattice",
		Content: sb.String(),
	}).Stage(stage)
}
