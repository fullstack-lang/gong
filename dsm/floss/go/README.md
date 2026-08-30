# FLOSS: Domain-Specific Modeling for System Complexity & Trade-Offs

**FLOSS** is a Domain-Specific Modeling (DSM) application built on the [Gong](https://github.com/fullstack-lang/gong) framework, designed for modeling, visualizing, and analyzing system architectures according to the **De Weck Complexity Equation**.

---

## 📐 The Theoretical Framework

The De Weck Complexity framework models the fundamental trade-off between architectural complexity, delivered performance, and engineering effort:

$$C = \mu P - \varepsilon E$$

When comparing two systems (or evaluating system evolution from baseline $S_1$ to variant $S_2$):

$$\Delta C = \mu \Delta P - \varepsilon \Delta E$$

Where:
* **$C$ (Complexity)**: Structural and operational complexity of the system architecture.
* **$P$ (Performance)**: Functional utility, capability, or capacity provided by the system.
* **$E$ (Effort)**: Engineering, standardisation, or tooling effort invested in architecting the system.
* **$\mu$ (Performance Multiplier)**: The marginal complexity cost per unit of performance gain.
* **$\varepsilon$ (Effort Multiplier)**: The complexity mitigation yield per unit of engineering effort.

---

## 🧮 Dynamic Pairwise $\mu$ & $\varepsilon$ Calibration

While a single system's parameters leave the De Weck equation underdetermined, comparing two system architectures $S_1 = (C_1, P_1, E_1)$ and $S_2 = (C_2, P_2, E_2)$ operating in the same evaluation frame yields a closed $2 \times 2$ linear system:

$$\begin{cases}
C_1 = \mu P_1 - \varepsilon E_1 \\
C_2 = \mu P_2 - \varepsilon E_2
\end{cases}$$

When $D = P_1 E_2 - P_2 E_1 \neq 0$, FLOSS automatically solves for the exact transition parameters:

$$\mu = \frac{C_1 E_2 - C_2 E_1}{P_1 E_2 - P_2 E_1}, \qquad \varepsilon = \frac{C_1 P_2 - C_2 P_1}{P_1 E_2 - P_2 E_1}$$

This guarantees exact local equilibrium ($\text{Diff} = 0$) and gives immediate insight into the architectural exchange rates.

---

## ✨ Core Features

* **Hierarchical System Modeling**:
  * Define Systems and Subsystems with compounded Complexity, Performance, and Effort items.
  * Weight strengths, categories, and descriptions for each architectural element.
* **Comparative Evolution Analysis (`CompareAnalysis`)**:
  * Map transitions from $S_1$ (`FromSystem`) to $S_2$ (`ToSystem`).
  * Automatic filtering of common elements to focus purely on net architectural mutations.
* **Interactive SVG Visualization (`DiagramFlossEquation`)**:
  * **6-Column Delta Pair View**: Side-by-side comparative stacks for $(C_2, C_1)$, $(P_2, P_1)$, and $(E_2, E_1)$ with difference block arrows and alignment level lines.
  * **3-Column Net Delta View**: Summarized representation of net changes $(\Delta C, \mu \Delta P, \varepsilon \Delta E)$.
  * **Dynamic Typography & Layout**: Responsive font scaling (Small, Normal, Big, Very Big) matching box fonts with floating $\mu$ and $\varepsilon$ metrics.
  * **Notes & Annotations**: Attach floating rectangular notes and connector shapes directly onto the diagram canvas.
* **Semantic Enforcement Engine**:
  * Self-healing data integrity rules that clean up orphaned shapes, enforce naming conventions, and compute extents automatically.

---

## 🚀 Getting Started

### Prerequisites

* **Go** (1.26+ recommended)

### Running the Application

To launch FLOSS in interactive edit mode:

```bash
cd dsm/floss/go/cmd/floss
go run . edit data/stage.go
```

Then open your browser to the local web interface (default: `http://localhost:8080`) or open the browser on [Floss](https://fullstack-lang.github.io/gong/floss-app-portable.html) to use it without installing it (the server runs in the browser).

---

## 🏛 Architecture & Codebase Structure

```
dsm/floss/go/
├── cmd/floss/                    # Main application executable
├── models/
│   ├── models_abstract.go        # Domain models (Library, System, Complexity, Performance, Effort, CompareAnalysis)
│   ├── diagram_floss_equation.go # Diagram and visual shape domain structures
│   ├── stager_enforce_floss_equation.go # Semantic rules & Cramer's rule computation for mu/epsilon
│   ├── stager_enforce_diagram_size.go   # Diagram extents & automatic canvas sizing
│   ├── stager_ux_svg_diagram_floss_equation.go # SVG rendering pipeline
│   ├── stager_ux_tree_*.go       # UI hierarchy tree builders and button callbacks
│   └── yyy_*.go                  # DSM generated glue code and callbacks
└── data/                         # Saved stages and scenario definitions (stage.go)
```

---

## 📄 License

This project is part of the Fullstack-lang Gong ecosystem and is licensed under the same terms as the parent repository.
