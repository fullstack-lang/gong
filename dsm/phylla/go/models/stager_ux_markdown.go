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

## 1. What are Spiral Plants?

Have you ever looked closely at the face of a sunflower, the scales of a pinecone, or the skin of a pineapple? If so, you have witnessed one of nature's most mesmerizing mathematical phenomena: **Spiral Phyllotaxis** (from Ancient Greek *phýllon* meaning "leaf" and *táxis* meaning "arrangement").

In nature, plants do not place their leaves, petals, or seeds at random. Instead, they organize them in interlocking geometric spirals that emerge from the growing tip (*meristem*) of the plant stem.

---

## 2. The Mystery of Fibonacci Numbers ($N$ and $M$)

When observing a sunflower or a pinecone, you can immediately spot two sets of spiral arms winding around the center in opposite directions:
- One family of spirals turns **clockwise** (shown in blue below).
- The opposing family of spirals turns **counter-clockwise** (shown in orange below).

![Phyllotaxis Parastichies in Spiral Plants](svg:spiral_parastichies?width=740px)

If you count the number of spirals in each direction, you will almost always find two consecutive numbers from the famous **Fibonacci Sequence**:

$$\mathbf{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, \dots}$$

For example:
- A typical **pinecone** displays **8** spirals winding one way and **13** winding the other: $(N, M) = (8, 13)$.
- A **pineapple** shows **8**, **13**, and **21** spirals: $(8, 13)$ and $(13, 21)$.
- A large **sunflower head** often features **34** and **55**, or even **55** and **89** spirals!

In botany and mathematics, these opposing spiral families are called **contact parastichies**, denoted as the pair $(N, M)$ where $N < M$. Consecutive Fibonacci numbers ensure the tightest possible tangential contact between adjacent seeds or scales.

---

## 3. The Secret Ingredient: The Golden Angle ($\alpha \approx 137.5^\circ$)

Why does nature follow this strict mathematical rule? The answer lies in how a plant grows at its microscopic tip (the *shoot apical meristem*).

As the plant stem grows, it produces new organ primordia (leaves, florets, petals) one after another. Each new organ is placed at a specific angular turn $\alpha$ from the previous one.

![The Golden Angle Divergence in Plant Meristems](svg:golden_angle?width=740px)

To ensure the plant thrives, this divergence angle must avoid repeating in simple fractions of a circle (which would cause leaves to stack directly on top of each other and block out light).

Nature's optimal solution is the **Golden Angle** ($\alpha \approx 137.5077^\circ$), derived from the Golden Ratio ($\phi = \frac{1 + \sqrt{5}}{2} \approx 1.618034$):

$$\alpha = 360^\circ \times \left(1 - \frac{1}{\phi}\right) \approx 137.507764^\circ$$

Because the Golden Ratio is the "most irrational" number, this angle guarantees:
1. **Maximum Sunlight & Rain Collection:** No leaf directly overshadows the leaf below it.
2. **Optimal Packing Efficiency:** Seeds and florets pack with maximum density, leaving zero wasted space across the head.

---

## 4. From Cylinder to Diamond Lattice

Imagine taking the cylindrical surface of a growing plant stem and unrolling it flat into a 2D sheet:

![Unrolling Cylindrical Stem to 2D Rhombus Lattice](svg:cylinder_to_lattice?width=780px)

1. **Diagonal Helical Paths:** The helical spiral paths in the 3D cylinder unfold into straight diagonal lines inclined across the 2D sheet:
   - **Blue lines:** The family of $N$ parastichies (clockwise spirals).
   - **Orange lines:** The family of $M$ parastichies (counter-clockwise spirals).
2. **The Rhombus (Diamond) Grid:** The intersection of these two families of diagonal lines tiles the 2D plane into a periodic **rhombus lattice**:
   - Each diamond/rhombus cell is bounded by two blue edges and two orange edges aligned with the diagonal spiral directions.
   - The **inside angle** (<b>RhombusInsideAngle</b>) determines the pitch/steepness of the spirals and the overall geometry of the plant.
   - The **side length** (<b>RhombusSideLength</b>) defines the physical distance or scale between adjacent nodes along the lattice.
3. **Growth Curves:** Stepping continuously across the rhombus grid generates continuous growth curves and ribbons that trace the biological development of the plant.

---

## 5. Generative Design with Phylla

**Phylla** harnesses these natural growth algorithms to create parametric 2D and 3D models:

- **Plant 2D & 3D:** Explore fundamental rhombus grids, calculate contact parastichy pairs **(<i>N</i>, <i>M</i>)**, and visualize the emerging spiral curves.
- **Vases:** Apply variable rotation curves and stack transformations to produce organic, twisting 3D vase geometries.
- **Stools & Furniture:** Use phyllotactic spiral vectors to generate parametric structural designs, such as self-supporting stool leg lattices.
- **Clocks:** Arrange numerals, ticks, and indicators following phyllotaxis spirals for harmonious dial layouts.
- **3D Printing (STL Export):** Export any generated 3D model directly to STL format for additive manufacturing using the **Download STL** button in the tree menu.
`,
	}
	content.Stage(stage)

	stage.Commit()
}

func stageSvgParastichies(stage *markdown.Stage) {
	cx, cy := 175.0, 190.0
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
	sb.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 740 380" width="100%" height="auto">`)
	sb.WriteString(`<rect width="740" height="380" rx="12" fill="#f8fafc" stroke="#cbd5e1" stroke-width="1.5"/>`)

	// Title
	sb.WriteString(`<text x="370" y="32" font-family="system-ui, sans-serif" font-size="16" font-weight="bold" fill="#0f172a" text-anchor="middle">Sunflower Head Phyllotaxis: Contact Parastichies (8, 13)</text>`)

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
			sb.WriteString(fmt.Sprintf(`<path d="%s" fill="none" stroke="#2563eb" stroke-width="2.2" stroke-opacity="0.85" stroke-linecap="round"/>`, strings.Join(pathParts, " ")))
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
			sb.WriteString(fmt.Sprintf(`<path d="%s" fill="none" stroke="#ea580c" stroke-width="2.2" stroke-opacity="0.85" stroke-linecap="round"/>`, strings.Join(pathParts, " ")))
		}
	}

	// Draw seeds / florets
	for k := 1; k <= numPts; k++ {
		p := pts[k]
		sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="3.2" fill="#1e293b" stroke="#ffffff" stroke-width="0.8"/>`, p.x, p.y))
	}

	// Legend Box on the right
	sb.WriteString(`
	<g transform="translate(420, 85)">
		<rect width="290" height="235" rx="8" fill="#ffffff" stroke="#e2e8f0" stroke-width="1.5"/>
		<text x="20" y="32" font-family="system-ui, sans-serif" font-size="14" font-weight="bold" fill="#0f172a">Parastichy Families</text>
		
		<!-- Blue Spiral Legend -->
		<line x1="20" y1="65" x2="60" y2="65" stroke="#2563eb" stroke-width="3" stroke-linecap="round"/>
		<text x="75" y="70" font-family="system-ui, sans-serif" font-size="13" font-weight="600" fill="#1e293b">8 Clockwise Spirals (N = 8)</text>
		<text x="75" y="88" font-family="system-ui, sans-serif" font-size="11" fill="#64748b">Steps along primordium index +8</text>

		<!-- Orange Spiral Legend -->
		<line x1="20" y1="125" x2="60" y2="125" stroke="#ea580c" stroke-width="3" stroke-linecap="round"/>
		<text x="75" y="130" font-family="system-ui, sans-serif" font-size="13" font-weight="600" fill="#1e293b">13 Counter-Clockwise (M = 13)</text>
		<text x="75" y="148" font-family="system-ui, sans-serif" font-size="11" fill="#64748b">Steps along primordium index +13</text>

		<!-- Fibonacci Pair Annotation -->
		<rect x="15" y="175" width="260" height="42" rx="6" fill="#f1f5f9"/>
		<text x="25" y="195" font-family="system-ui, sans-serif" font-size="12" font-weight="bold" fill="#0f172a">Contact Parastichy Pair:</text>
		<text x="25" y="209" font-family="system-ui, sans-serif" font-size="11" fill="#475569">(N, M) = (8, 13) (Fibonacci numbers)</text>
	</g>
	`)

	sb.WriteString(`</svg>`)

	(&markdown.SvgImage{
		Name:    "spiral_parastichies",
		Content: sb.String(),
	}).Stage(stage)
}

func stageSvgGoldenAngle(stage *markdown.Stage) {
	cx, cy := 190.0, 190.0
	rOuter := 130.0

	var sb strings.Builder
	sb.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 740 380" width="100%" height="auto">`)
	sb.WriteString(`<rect width="740" height="380" rx="12" fill="#f8fafc" stroke="#cbd5e1" stroke-width="1.5"/>`)

	// Title
	sb.WriteString(`<text x="370" y="32" font-family="system-ui, sans-serif" font-size="16" font-weight="bold" fill="#0f172a" text-anchor="middle">Shoot Apical Meristem: Primordia Divergence Angle α ≈ 137.5°</text>`)

	// Meristem Circle
	sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="%.1f" fill="#f1f5f9" stroke="#94a3b8" stroke-width="1.5" stroke-dasharray="4,4"/>`, cx, cy, rOuter))
	sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="18" fill="#e2e8f0" stroke="#64748b" stroke-width="1.5"/>`, cx, cy))
	sb.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="10" font-weight="bold" fill="#475569" text-anchor="middle">Apex</text>`, cx, cy+3))

	// Successive Primordia (0 to 6)
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

	sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#94a3b8" stroke-width="1.5"/>`, cx, cy, p0X+20, p0Y))
	sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#94a3b8" stroke-width="1.5"/>`, cx, cy, p1X-15, p1Y-15))

	// Arc for divergence angle α (counter-clockwise from 0° to 137.5°)
	arcR := 52.0
	arcX0 := cx + arcR
	arcY0 := cy
	arcX1 := cx + arcR*math.Cos(137.5077*math.Pi/180.0)
	arcY1 := cy - arcR*math.Sin(137.5077*math.Pi/180.0)
	sb.WriteString(fmt.Sprintf(`<path d="M %.1f %.1f A %.1f %.1f 0 0 0 %.1f %.1f" fill="none" stroke="#2563eb" stroke-width="2.5"/>`, arcX0, arcY0, arcR, arcR, arcX1, arcY1))
	sb.WriteString(fmt.Sprintf(`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="#2563eb"/>`, arcX1+3, arcY1+10, arcX1-8, arcY1+2, arcX1-2, arcY1+12))
	sb.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="14" font-weight="bold" fill="#2563eb">α ≈ 137.5°</text>`, cx+30, cy-50))

	// Draw and label each primordium
	for _, pr := range prims {
		px := cx + pr.rad*math.Cos(pr.deg*math.Pi/180.0)
		py := cy - pr.rad*math.Sin(pr.deg*math.Pi/180.0)

		fillCol := "#10b981"
		if pr.num == 0 || pr.num == 1 {
			fillCol = "#2563eb"
		}
		sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="11" fill="%s" stroke="#ffffff" stroke-width="1.5"/>`, px, py, fillCol))
		sb.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#ffffff" text-anchor="middle">P%d</text>`, px, py+4, pr.num))
	}

	// Right Explanation Box
	sb.WriteString(`
	<g transform="translate(420, 75)">
		<rect width="290" height="255" rx="8" fill="#ffffff" stroke="#e2e8f0" stroke-width="1.5"/>
		<text x="20" y="32" font-family="system-ui, sans-serif" font-size="14" font-weight="bold" fill="#0f172a">Golden Divergence Principle</text>
		
		<text x="20" y="62" font-family="system-ui, sans-serif" font-size="12" fill="#334155">• Each new organ primordium (P₀, P₁, P₂...)</text>
		<text x="30" y="80" font-family="system-ui, sans-serif" font-size="12" fill="#334155">is initiated at the apex.</text>
		
		<text x="20" y="110" font-family="system-ui, sans-serif" font-size="12" fill="#334155">• Divergence angle between successive</text>
		<text x="30" y="128" font-family="system-ui, sans-serif" font-size="12" fill="#334155">primordia is strictly constant:</text>

		<rect x="20" y="145" width="250" height="40" rx="6" fill="#eff6ff"/>
		<text x="30" y="170" font-family="monospace" font-size="13" font-weight="bold" fill="#1d4ed8">α = 360° × (1 − 1/φ) ≈ 137.5°</text>

		<text x="20" y="210" font-family="system-ui, sans-serif" font-size="11" fill="#475569">Older primordia (P₀) are pushed radially</text>
		<text x="20" y="226" font-family="system-ui, sans-serif" font-size="11" fill="#475569">outward as new ones (P₁, P₂, ...) emerge.</text>
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
	sb.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 820 370" width="100%" height="auto">`)
	sb.WriteString(`<rect width="820" height="370" rx="12" fill="#ffffff" stroke="#e2e8f0" stroke-width="1.5"/>`)

	// Left: 3D Cylinder
	sb.WriteString(`
	<g transform="translate(35, 45)">
		<text x="80" y="0" font-family="system-ui, sans-serif" font-size="15" font-weight="bold" fill="#0f172a" text-anchor="middle">3D Plant Stem Cylinder</text>
		
		<defs>
			<linearGradient id="cylGrad" x1="0" y1="0" x2="1" y2="0">
				<stop offset="0%" stop-color="#94a3b8"/>
				<stop offset="25%" stop-color="#f1f5f9"/>
				<stop offset="60%" stop-color="#e2e8f0"/>
				<stop offset="100%" stop-color="#64748b"/>
			</linearGradient>
			<clipPath id="sheetClip">
				<rect x="25" y="30" width="430" height="215" rx="4"/>
			</clipPath>
		</defs>

		<!-- Bottom Ellipse -->
		<ellipse cx="80" cy="225" rx="60" ry="22" fill="#cbd5e1" stroke="#64748b" stroke-width="2"/>
		<!-- Body rect -->
		<rect x="20" y="55" width="120" height="170" fill="url(#cylGrad)"/>
		<!-- Side lines -->
		<line x1="20" y1="55" x2="20" y2="225" stroke="#64748b" stroke-width="2"/>
		<line x1="140" y1="55" x2="140" y2="225" stroke="#64748b" stroke-width="2"/>
		<!-- Top Ellipse -->
		<ellipse cx="80" cy="55" rx="60" ry="22" fill="#f8fafc" stroke="#64748b" stroke-width="2"/>

		<!-- Helical parastichy curves on cylinder -->
		<!-- Blue spirals (up and right) -->
		<path d="M 20 190 Q 80 220 140 160" fill="none" stroke="#2563eb" stroke-width="2.5" stroke-linecap="round"/>
		<path d="M 20 140 Q 80 170 140 110" fill="none" stroke="#2563eb" stroke-width="2.5" stroke-linecap="round"/>
		<path d="M 20 90 Q 80 120 140 60" fill="none" stroke="#2563eb" stroke-width="2.5" stroke-linecap="round"/>
		<!-- Blue dashed back -->
		<path d="M 140 160 Q 80 140 20 140" fill="none" stroke="#93c5fd" stroke-width="1.5" stroke-dasharray="3,3"/>
		<path d="M 140 110 Q 80 90 20 90" fill="none" stroke="#93c5fd" stroke-width="1.5" stroke-dasharray="3,3"/>

		<!-- Orange spirals (up and left) -->
		<path d="M 140 190 Q 80 220 20 160" fill="none" stroke="#ea580c" stroke-width="2.5" stroke-linecap="round"/>
		<path d="M 140 140 Q 80 170 20 110" fill="none" stroke="#ea580c" stroke-width="2.5" stroke-linecap="round"/>
		<path d="M 140 90 Q 80 120 20 60" fill="none" stroke="#ea580c" stroke-width="2.5" stroke-linecap="round"/>
		<!-- Orange dashed back -->
		<path d="M 20 160 Q 80 140 140 140" fill="none" stroke="#fdba74" stroke-width="1.5" stroke-dasharray="3,3"/>
		<path d="M 20 110 Q 80 90 140 90" fill="none" stroke="#fdba74" stroke-width="1.5" stroke-dasharray="3,3"/>

		<!-- Cut line -->
		<line x1="20" y1="55" x2="20" y2="225" stroke="#db2777" stroke-width="2" stroke-dasharray="4,4"/>
		<rect x="0" y="130" width="40" height="18" rx="4" fill="#fdf2f8" stroke="#fbcfe8"/>
		<text x="20" y="142" font-family="system-ui, sans-serif" font-size="9" font-weight="bold" fill="#db2777" text-anchor="middle">Cut Line</text>

		<!-- Bottom Circumference Label -->
		<line x1="20" y1="260" x2="140" y2="260" stroke="#475569" stroke-width="1.5"/>
		<text x="80" y="278" font-family="system-ui, sans-serif" font-size="12" font-weight="600" fill="#475569" text-anchor="middle">Circumference C = 2πR</text>
	</g>
	`)

	// Center: Transition Arrow
	sb.WriteString(`
	<g transform="translate(205, 150)">
		<rect x="0" y="0" width="85" height="42" rx="8" fill="#eff6ff" stroke="#bfdbfe" stroke-width="1.5"/>
		<text x="42" y="18" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#1d4ed8" text-anchor="middle">UNROLL</text>
		<text x="42" y="32" font-family="system-ui, sans-serif" font-size="9" fill="#3b82f6" text-anchor="middle">Cylinder → 2D</text>
		<path d="M 94 21 L 115 21 M 108 14 L 115 21 L 108 28" fill="none" stroke="#1d4ed8" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
	</g>
	`)

	// Right: 2D Periodic Rhombus Grid
	sb.WriteString(`
	<g transform="translate(315, 45)">
		<text x="240" y="0" font-family="system-ui, sans-serif" font-size="15" font-weight="bold" fill="#0f172a" text-anchor="middle">2D Rhombus Lattice (Diagonal Spiral Lines)</text>

		<!-- Unrolled 2D Sheet Background -->
		<rect x="25" y="30" width="430" height="215" rx="4" fill="#f8fafc" stroke="#94a3b8" stroke-width="1.5"/>

		<!-- Grid of Rhombuses (Clipped to Sheet) -->
		<g clip-path="url(#sheetClip)">
	`)

	// Lattice vectors (diagonal along spiral directions):
	// Blue vector (up and right): ux, uy = 48.0, -24.6 (length = 54)
	// Orange vector (up and left): vx, vy = -32.0, -43.5 (length = 54)
	ux, uy := 48.0, -24.6
	vx, vy := -32.0, -43.5
	ox, oy := 140.0, 230.0

	// Draw regular rhombus cells
	for j := -2; j <= 6; j++ {
		for i := -2; i <= 9; i++ {
			x0 := ox + float64(i)*ux + float64(j)*vx
			y0 := oy + float64(i)*uy + float64(j)*vy
			x1 := x0 + ux
			y1 := y0 + uy
			x2 := x0 + ux + vx
			y2 := y0 + uy + vy
			x3 := x0 + vx
			y3 := y0 + vy

			isRef := (i == 2 && j == 1)

			fillCol := "#ffffff"
			strokeCol := "#cbd5e1"
			strokeW := "1.2"

			if isRef {
				fillCol = "#e0f2fe"
				strokeCol = "#0284c7"
				strokeW = "2.5"
			}

			sb.WriteString(fmt.Sprintf(`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="%s" stroke="%s" stroke-width="%s"/>`,
				x0, y0, x1, y1, x2, y2, x3, y3, fillCol, strokeCol, strokeW))
		}
	}

	// Draw diagonal parastichy guidelines
	// Blue diagonal lines (along +u)
	for j := -1; j <= 5; j++ {
		lx1 := ox + float64(-2)*ux + float64(j)*vx
		ly1 := oy + float64(-2)*uy + float64(j)*vy
		lx2 := ox + float64(9)*ux + float64(j)*vx
		ly2 := oy + float64(9)*uy + float64(j)*vy
		sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#2563eb" stroke-width="1.8" stroke-opacity="0.7"/>`,
			lx1, ly1, lx2, ly2))
	}
	// Orange diagonal lines (along +v)
	for i := -1; i <= 8; i++ {
		lx1 := ox + float64(i)*ux + float64(-2)*vx
		ly1 := oy + float64(i)*uy + float64(-2)*vy
		lx2 := ox + float64(i)*ux + float64(6)*vx
		ly2 := oy + float64(i)*uy + float64(6)*vy
		sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#ea580c" stroke-width="1.8" stroke-opacity="0.7"/>`,
			lx1, ly1, lx2, ly2))
	}

	// Draw primordia / lattice node dots
	for j := -1; j <= 5; j++ {
		for i := -1; i <= 8; i++ {
			nx := ox + float64(i)*ux + float64(j)*vx
			ny := oy + float64(i)*uy + float64(j)*vy
			if nx >= 20 && nx <= 460 && ny >= 25 && ny <= 250 {
				sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="3" fill="#334155"/>`, nx, ny))
			}
		}
	}

	sb.WriteString(`
		</g>

		<!-- Periodic Boundary Markers -->
		<line x1="25" y1="30" x2="25" y2="245" stroke="#db2777" stroke-width="2" stroke-dasharray="4,4"/>
		<line x1="455" y1="30" x2="455" y2="245" stroke="#db2777" stroke-width="2" stroke-dasharray="4,4"/>
		<text x="25" y="24" font-family="system-ui, sans-serif" font-size="10" font-weight="bold" fill="#db2777" text-anchor="middle">x = 0</text>
		<text x="455" y="24" font-family="system-ui, sans-serif" font-size="10" font-weight="bold" fill="#db2777" text-anchor="middle">x = C (Periodic)</text>

		<!-- Reference Rhombus Annotations (i=2, j=1) -->
	`)

	refA_x := ox + 2.0*ux + 1.0*vx
	refA_y := oy + 2.0*uy + 1.0*vy
	refB_x := refA_x + ux
	refB_y := refA_y + uy
	refC_x := refA_x + ux + vx
	refC_y := refA_y + uy + vy
	refD_x := refA_x + vx
	refD_y := refA_y + vy

	// Draw highlighted rhombus on top
	sb.WriteString(fmt.Sprintf(`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="#e0f2fe" fill-opacity="0.85" stroke="#0284c7" stroke-width="2.5"/>`,
		refA_x, refA_y, refB_x, refB_y, refC_x, refC_y, refD_x, refD_y))

	sb.WriteString(fmt.Sprintf(`
		<!-- Side Length L along diagonal blue edge -->
		<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#0369a1">Side Length (L)</text>
		<!-- Inside Angle α at bottom corner -->
		<path d="M %.1f %.1f A 16 16 0 0 1 %.1f %.1f" fill="none" stroke="#dc2626" stroke-width="2"/>
		<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#dc2626" text-anchor="middle">α (Inside Angle)</text>

		<!-- Growth Vector Arrow (Stem Axis) -->
		<line x1="380" y1="210" x2="380" y2="85" stroke="#059669" stroke-width="3" stroke-linecap="round"/>
		<polygon points="380,75 374,88 386,88" fill="#059669"/>
		<text x="390" y="145" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#059669">Growth Vector G</text>

		<!-- Spiral direction labels -->
		<text x="130" y="60" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#2563eb">Blue Parastichy (N)</text>
		<text x="30" y="110" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#ea580c">Orange Parastichy (M)</text>

		<!-- Circumference bottom dimension -->
		<line x1="25" y1="260" x2="455" y2="260" stroke="#475569" stroke-width="1.5"/>
		<text x="240" y="278" font-family="system-ui, sans-serif" font-size="12" font-weight="600" fill="#475569" text-anchor="middle">Stem Circumference C = 2πR (Periodic Width)</text>
	`,
		(refA_x+refB_x)/2.0-10, (refA_y+refB_y)/2.0+16,
		refA_x+14, refA_y-8, refA_x-10, refA_y-13,
		refA_x+5, refA_y+20,
	))

	sb.WriteString(`
	</g>
	</svg>`)

	(&markdown.SvgImage{
		Name:    "cylinder_to_lattice",
		Content: sb.String(),
	}).Stage(stage)
}
