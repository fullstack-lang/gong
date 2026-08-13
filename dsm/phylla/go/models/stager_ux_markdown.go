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

![Unrolling Cylindrical Stem to 2D Rhombus Lattice](svg:cylinder_to_lattice?width=760px)

1. **The Rhombus Grid:** The helical spiral paths in the 3D cylinder unfold into a periodic grid of diamond (rhombus) shapes.
2. **Inside Angle & Side Length:**
   - The **inside angle** (<b>RhombusInsideAngle</b>) of the rhombus determines the steepness of the spirals and the overall geometry of the plant.
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
		th := float64(k) * 137.507764 * math.Pi / 180.0
		pts[k] = pt{
			x: cx + r*math.Cos(th),
			y: cy + r*math.Sin(th),
		}
	}

	var sb strings.Builder
	sb.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 740 380" width="100%" height="auto">`)
	sb.WriteString(`<rect width="740" height="380" rx="12" fill="#f8fafc" stroke="#cbd5e1" stroke-width="1.5"/>`)

	// Background circle for seed head
	sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="145" fill="#f1f5f9" stroke="#e2e8f0" stroke-width="1.5"/>`, cx, cy))

	// 8 clockwise spirals (step +8)
	for start := 1; start <= 8; start++ {
		var chain []string
		k := start
		for k <= numPts {
			p := pts[k]
			chain = append(chain, fmt.Sprintf("%.1f,%.1f", p.x, p.y))
			k += 8
		}
		if len(chain) > 1 {
			sb.WriteString(fmt.Sprintf(`<path d="M %s" fill="none" stroke="#2563eb" stroke-width="2.5" stroke-linecap="round" opacity="0.85"/>`, strings.Join(chain, " L ")))
		}
	}

	// 13 counter-clockwise spirals (step +13)
	for start := 1; start <= 13; start++ {
		var chain []string
		k := start
		for k <= numPts {
			p := pts[k]
			chain = append(chain, fmt.Sprintf("%.1f,%.1f", p.x, p.y))
			k += 13
		}
		if len(chain) > 1 {
			sb.WriteString(fmt.Sprintf(`<path d="M %s" fill="none" stroke="#ea580c" stroke-width="2.2" stroke-linecap="round" opacity="0.85"/>`, strings.Join(chain, " L ")))
		}
	}

	// Points
	for k := 1; k <= numPts; k++ {
		p := pts[k]
		sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="3.2" fill="#334155"/>`, p.x, p.y))
	}
	sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="5" fill="#0f172a"/>`, cx, cy))

	// Right panel: Title and Cards
	sb.WriteString(`
	<g transform="translate(365, 30)">
		<text x="0" y="24" font-family="system-ui, sans-serif" font-size="18" font-weight="bold" fill="#0f172a">Contact Parastichy Pair (N, M)</text>
		<text x="0" y="48" font-family="system-ui, sans-serif" font-size="13" fill="#64748b">Pinecone / Sunflower Fibonacci Spirals</text>

		<!-- Blue Card (Clockwise N=8) -->
		<rect x="0" y="70" width="345" height="95" rx="8" fill="#eff6ff" stroke="#bfdbfe" stroke-width="1.5"/>
		<circle cx="22" cy="98" r="8" fill="#2563eb"/>
		<text x="38" y="103" font-family="system-ui, sans-serif" font-size="15" font-weight="bold" fill="#1e40af">8 Clockwise Spirals (N = 8)</text>
		<text x="22" y="128" font-family="system-ui, sans-serif" font-size="12.5" fill="#3b82f6">• Family of 8 spirals curving clockwise</text>
		<text x="22" y="148" font-family="system-ui, sans-serif" font-size="12.5" fill="#3b82f6">• Connects node k to node k + 8</text>

		<!-- Orange Card (Counter-Clockwise M=13) -->
		<rect x="0" y="180" width="345" height="95" rx="8" fill="#fff7ed" stroke="#fed7aa" stroke-width="1.5"/>
		<circle cx="22" cy="208" r="8" fill="#ea580c"/>
		<text x="38" y="213" font-family="system-ui, sans-serif" font-size="15" font-weight="bold" fill="#9a3412">13 Counter-Clockwise Spirals (M = 13)</text>
		<text x="22" y="238" font-family="system-ui, sans-serif" font-size="12.5" fill="#c2410c">• Opposing family of 13 spirals</text>
		<text x="22" y="258" font-family="system-ui, sans-serif" font-size="12.5" fill="#c2410c">• Connects node k to node k + 13</text>

		<!-- Fibonacci bottom banner -->
		<rect x="0" y="290" width="345" height="42" rx="6" fill="#f1f5f9" stroke="#cbd5e1"/>
		<text x="172" y="316" text-anchor="middle" font-family="system-ui, sans-serif" font-size="13" font-weight="600" fill="#334155">Fibonacci Pair: (N, M) = (8, 13) with N &lt; M</text>
	</g>
	`)

	sb.WriteString(`</svg>`)

	(&markdown.SvgImage{
		Name:    "spiral_parastichies",
		Content: sb.String(),
	}).Stage(stage)
}

func stageSvgGoldenAngle(stage *markdown.Stage) {
	var sb strings.Builder
	sb.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 740 360" width="100%" height="auto">`)
	sb.WriteString(`<rect width="740" height="360" rx="12" fill="#f8fafc" stroke="#cbd5e1" stroke-width="1.5"/>`)

	cx, cy := 180.0, 180.0

	// Concentric guide rings
	for _, r := range []float64{45, 80, 115, 145} {
		sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="%.1f" fill="none" stroke="#e2e8f0" stroke-dasharray="4 4"/>`, cx, cy, r))
	}

	// Calculate primordia 0..8
	type prim struct {
		k    int
		x, y float64
		r    float64
		deg  float64
	}
	var prims []prim
	c := 44.0
	for k := 0; k <= 8; k++ {
		r := c * math.Sqrt(float64(k)+0.4)
		deg := float64(k) * 137.507764
		rad := deg * math.Pi / 180.0
		prims = append(prims, prim{
			k:   k,
			x:   cx + r*math.Cos(rad),
			y:   cy + r*math.Sin(rad),
			r:   r,
			deg: deg,
		})
	}

	// Divergence angle arc between 0 and 1
	p0 := prims[0]
	p1 := prims[1]
	sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#94a3b8" stroke-width="1.5" stroke-dasharray="3 3"/>`, cx, cy, p0.x, p0.y))
	sb.WriteString(fmt.Sprintf(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#94a3b8" stroke-width="1.5" stroke-dasharray="3 3"/>`, cx, cy, p1.x, p1.y))

	// Arc for 137.5 deg
	arcR := 52.0
	arcStartX := cx + arcR
	arcStartY := cy
	arcEndRad := 137.507764 * math.Pi / 180.0
	arcEndX := cx + arcR*math.Cos(arcEndRad)
	arcEndY := cy + arcR*math.Sin(arcEndRad)
	sb.WriteString(fmt.Sprintf(`<path d="M %.1f %.1f A %.1f %.1f 0 0 1 %.1f %.1f" fill="none" stroke="#059669" stroke-width="3"/>`, arcStartX, arcStartY, arcR, arcR, arcEndX, arcEndY))
	sb.WriteString(fmt.Sprintf(`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="#059669"/>`,
		arcEndX+1, arcEndY-7, arcEndX-6, arcEndY+1, arcEndX+4, arcEndY+4))

	// Label for the angle
	sb.WriteString(fmt.Sprintf(`<rect x="%.1f" y="%.1f" width="90" height="26" rx="4" fill="#ecfdf5" stroke="#a7f3d0"/>`, cx-15, cy-65))
	sb.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="12" font-weight="bold" fill="#065f46" text-anchor="middle">α ≈ 137.5°</text>`, cx+30, cy-48))

	// Center apex
	sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="6" fill="#0f172a"/>`, cx, cy))
	sb.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#64748b" text-anchor="middle">Apex</text>`, cx, cy+18))

	// Draw primordia badges
	for _, p := range prims {
		sb.WriteString(fmt.Sprintf(`<circle cx="%.1f" cy="%.1f" r="14" fill="#3b82f6" stroke="#ffffff" stroke-width="2"/>`, p.x, p.y))
		sb.WriteString(fmt.Sprintf(`<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="12" font-weight="bold" fill="#ffffff" text-anchor="middle" dominant-baseline="central">%d</text>`, p.x, p.y, p.k))
	}

	// Right Panel
	sb.WriteString(`
	<g transform="translate(365, 25)">
		<text x="0" y="24" font-family="system-ui, sans-serif" font-size="18" font-weight="bold" fill="#0f172a">Divergence Angle (α ≈ 137.5°)</text>
		<text x="0" y="46" font-family="system-ui, sans-serif" font-size="13" fill="#64748b">Growth Tip Primordia Generation (0 → 1 → 2 → ...)</text>

		<!-- Card 1: Math Formula -->
		<rect x="0" y="65" width="345" height="110" rx="8" fill="#f0fdf4" stroke="#bbf7d0" stroke-width="1.5"/>
		<text x="18" y="90" font-family="system-ui, sans-serif" font-size="14" font-weight="bold" fill="#166534">Mathematical Derivation</text>
		<text x="18" y="115" font-family="monospace" font-size="14" font-weight="bold" fill="#14532d">α = 360° × (1 − 1/φ) ≈ 137.5077°</text>
		<text x="18" y="140" font-family="system-ui, sans-serif" font-size="12" fill="#15803d">where φ = (1 + √5) / 2 ≈ 1.618034 (Golden Ratio)</text>
		<text x="18" y="158" font-family="system-ui, sans-serif" font-size="12" fill="#15803d">Each step advances by exactly 1 − 1/φ ≈ 0.381966 turns</text>

		<!-- Card 2: Biological Advantages -->
		<rect x="0" y="190" width="345" height="115" rx="8" fill="#f8fafc" stroke="#cbd5e1" stroke-width="1.5"/>
		<text x="18" y="215" font-family="system-ui, sans-serif" font-size="14" font-weight="bold" fill="#0f172a">Biological Significance</text>
		<text x="18" y="240" font-family="system-ui, sans-serif" font-size="12.5" fill="#334155">1. <tspan font-weight="bold">No Shadowing:</tspan> Leaves never align directly vertically</text>
		<text x="18" y="262" font-family="system-ui, sans-serif" font-size="12.5" fill="#334155">2. <tspan font-weight="bold">Optimal Packing:</tspan> Maximizes seed and floret density</text>
		<text x="18" y="284" font-family="system-ui, sans-serif" font-size="12.5" fill="#334155">3. <tspan font-weight="bold">Irrational Angle:</tspan> Avoids periodic alignment patterns</text>
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
	sb.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 760 360" width="100%" height="auto">`)
	sb.WriteString(`<rect width="760" height="360" rx="12" fill="#f8fafc" stroke="#cbd5e1" stroke-width="1.5"/>`)

	// Left: 3D Cylinder
	sb.WriteString(`
	<g transform="translate(40, 45)">
		<text x="75" y="0" font-family="system-ui, sans-serif" font-size="16" font-weight="bold" fill="#0f172a" text-anchor="middle">3D Plant Stem Cylinder</text>
		
		<!-- Cylinder Body -->
		<defs>
			<linearGradient id="cylGrad" x1="0" y1="0" x2="1" y2="0">
				<stop offset="0%" stop-color="#cbd5e1"/>
				<stop offset="35%" stop-color="#f8fafc"/>
				<stop offset="70%" stop-color="#e2e8f0"/>
				<stop offset="100%" stop-color="#94a3b8"/>
			</linearGradient>
		</defs>

		<!-- Bottom Ellipse -->
		<ellipse cx="75" cy="225" rx="65" ry="24" fill="#cbd5e1" stroke="#64748b" stroke-width="2"/>
		<!-- Body rect -->
		<rect x="10" y="55" width="130" height="170" fill="url(#cylGrad)"/>
		<!-- Side lines -->
		<line x1="10" y1="55" x2="10" y2="225" stroke="#64748b" stroke-width="2"/>
		<line x1="140" y1="55" x2="140" y2="225" stroke="#64748b" stroke-width="2"/>
		<!-- Top Ellipse -->
		<ellipse cx="75" cy="55" rx="65" ry="24" fill="#e2e8f0" stroke="#64748b" stroke-width="2"/>

		<!-- Helical parastichy curves on cylinder -->
		<path d="M 10 90 Q 75 125 140 160" fill="none" stroke="#2563eb" stroke-width="2.5" stroke-linecap="round"/>
		<path d="M 10 135 Q 75 170 140 205" fill="none" stroke="#2563eb" stroke-width="2.5" stroke-linecap="round"/>
		<path d="M 140 85 Q 75 120 10 155" fill="none" stroke="#ea580c" stroke-width="2.5" stroke-linecap="round"/>
		<path d="M 140 130 Q 75 165 10 200" fill="none" stroke="#ea580c" stroke-width="2.5" stroke-linecap="round"/>

		<!-- Labels -->
		<text x="75" y="270" font-family="system-ui, sans-serif" font-size="12" font-weight="600" fill="#475569" text-anchor="middle">Circumference C = 2πR</text>
	</g>
	`)

	// Center: Transition Arrow
	sb.WriteString(`
	<g transform="translate(230, 155)">
		<rect x="0" y="0" width="85" height="36" rx="6" fill="#eff6ff" stroke="#bfdbfe"/>
		<text x="42" y="16" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#1d4ed8" text-anchor="middle">UNROLL</text>
		<text x="42" y="28" font-family="system-ui, sans-serif" font-size="9" fill="#3b82f6" text-anchor="middle">Cylinder → 2D</text>
		<path d="M 92 18 L 115 18 M 107 11 L 115 18 L 107 25" fill="none" stroke="#1d4ed8" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
	</g>
	`)

	// Right: 2D Periodic Rhombus Grid
	sb.WriteString(`
	<g transform="translate(375, 45)">
		<text x="175" y="0" font-family="system-ui, sans-serif" font-size="16" font-weight="bold" fill="#0f172a" text-anchor="middle">2D Rhombus (Diamond) Lattice</text>

		<!-- Grid of Rhombuses -->
		<g transform="translate(35, 30)">
	`)

	// Draw rhombuses grid
	rW := 65.0
	rH := 42.0
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			xBase := float64(col)*rW*1.2 + float64(row)*rW*0.35
			yBase := float64(row)*rH*1.15 + 20.0
			x0, y0 := xBase, yBase
			x1, y1 := xBase + rW*0.6, yBase - rH*0.55
			x2, y2 := xBase + rW*1.2, yBase
			x3, y3 := xBase + rW*0.6, yBase + rH*0.55

			fillCol := "#f1f5f9"
			strokeCol := "#cbd5e1"
			strokeW := "1.2"

			// Highlight one reference rhombus
			if row == 1 && col == 1 {
				fillCol = "#e0f2fe"
				strokeCol = "#0284c7"
				strokeW = "2.5"
			}

			sb.WriteString(fmt.Sprintf(`<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="%s" stroke="%s" stroke-width="%s"/>`,
				x0, y0, x1, y1, x2, y2, x3, y3, fillCol, strokeCol, strokeW))
		}
	}

	// Rhombus parameters annotation on the highlighted rhombus (row=1, col=1)
	refX := 1.0*rW*1.2 + 1.0*rW*0.35
	refY := 1.0*rH*1.15 + 20.0
	sb.WriteString(fmt.Sprintf(`
		<!-- Side Length L -->
		<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="12" font-weight="bold" fill="#0369a1">Side Length (L)</text>
		<!-- Inside Angle α -->
		<path d="M %.1f %.1f A 14 14 0 0 1 %.1f %.1f" fill="none" stroke="#dc2626" stroke-width="2"/>
		<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="11" font-weight="bold" fill="#dc2626">α (Inside Angle)</text>

		<!-- Growth Vector Arrow -->
		<path d="M %.1f %.1f L %.1f %.1f" stroke="#059669" stroke-width="3" stroke-linecap="round"/>
		<polygon points="%.1f,%.1f %.1f,%.1f %.1f,%.1f" fill="#059669"/>
		<text x="%.1f" y="%.1f" font-family="system-ui, sans-serif" font-size="12" font-weight="bold" fill="#059669">Growth Vector G</text>
	`,
		refX-40, refY-15,
		refX+14, refY-3, refX+9, refY+9,
		refX-25, refY+16,
		refX+rW*1.2+30, refY+50, refX+rW*1.2+30, refY-35,
		refX+rW*1.2+30, refY-43, refX+rW*1.2+25, refY-33, refX+rW*1.2+35, refY-33,
		refX+rW*1.2+36, refY+8,
	))

	sb.WriteString(`
		</g>

		<!-- Bottom summary line -->
		<text x="175" y="270" font-family="system-ui, sans-serif" font-size="12" font-weight="600" fill="#475569" text-anchor="middle">Periodic lattice paths form continuous 3D spirals on cylinder</text>
	</g>
	`)

	sb.WriteString(`</svg>`)

	(&markdown.SvgImage{
		Name:    "cylinder_to_lattice",
		Content: sb.String(),
	}).Stage(stage)
}
