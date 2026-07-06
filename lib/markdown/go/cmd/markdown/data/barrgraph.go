package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Content__00000000_ := (&models.Content{Name: `bargraph`}).Stage(stage)

	__SvgImage__00000000_ := (&models.SvgImage{Name: `barrgraph`}).Stage(stage)

	// insertion point for initialization of values

	__Content__00000000_.Name = `bargraph`
	__Content__00000000_.Content = `# BarrGraph

**BarrGraph** is a specialized visualization tool designed for art curators, historians, and researchers. It enables the creation of chronological flowcharts to map the genealogy and evolution of art movements, inspired by the methodology of **Alfred Barr**.

## Inspiration

Alfred H. Barr Jr., the first director of the Museum of Modern Art (MoMA), famously used flowcharts to explain the origins and development of abstract art. His diagram for the 1936 exhibition *Cubism and Abstract Art* remains a canonical example of data visualization in art history.

**BarrGraph** digitizes this approach, allowing users to construct directed graphs where nodes represent art movements and edges represent influence, all plotted against a chronological axis.

## Features

- **Directed Influence Maps:** Model the flow of influence between movements (e.g., Cézanne → Cubism).
- **Chronological Alignment:** Automatically align movements by their active decades.
- **Curatorial Metadata:** Attach descriptions, key artists, and representative works to each node.
- **Visual Styling:** Export charts that pay homage to the classic mid-century aesthetic of Barr's original posters.

## Abstract Syntax

The application is built around a domain-specific model that captures the relationships between art movements.

![Abstract Syntax](images/barr%20absract%20syntax.png)
*The abstract syntax defining ArtMovements and their Influences.*

The model consists of:
- **ArtMovement**: Represents a distinct artistic movement or period.
- **ArtefactType**: A categorization of art (e.g., Painting, Sculpture, Architecture).
- **Artist**: Individual artists associated with movements.
- **Influence**: Represents a directed relationship where one movement influences another. 
An "Influence" links a **Source** "ArtMovement/ArtefactType/Artist" 
to a **Target** "ArtMovement/ArtefactType/Artist".

## Gallery

Below are examples of the tool in use, including the original inspiration and the digital recreation.

![Alfred Barr's Original Poster, redrawn](svg:barrgraph?width=800px)

## Installation

go run github.com/fullstack-lang/gong/dsm/barrgraph/go/cmd/barrgraph@latest

then open a browser on http://localhost:8080/.

## License

Distributed under the MIT License.`

	__SvgImage__00000000_.Name = `barrgraph`
	__SvgImage__00000000_.Content = `<svg xmlns="http://www.w3.org/2000/svg" width="790" height="956.3636363636363" viewBox="-10 -10 790 956.3636363636363"><style>text { font-family: Roboto, Arial, sans-serif !important; }</style>
  <g transform="translate(0 0) scale(1)">
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="0" y="0" width="770" height="786.3636363636363" rx="0" fill="#DED6CA" fill-opacity="1" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="40" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="40" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1890</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="121.81818181818181" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="121.81818181818181" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1895</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="203.63636363636363" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="203.63636363636363" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1900</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="285.45454545454544" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="285.45454545454544" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1905</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="367.27272727272725" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="367.27272727272725" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1910</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="449.09090909090907" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="449.09090909090907" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1915</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="530.9090909090909" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="530.9090909090909" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1920</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="612.7272727272727" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="612.7272727272727" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1925</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="694.5454545454545" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="694.5454545454545" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1930</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="20" y="776.3636363636363" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="20" y="776.3636363636363" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="20" dy="0" text-anchor="">1935</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="40" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="40" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1890</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="121.81818181818181" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="121.81818181818181" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1895</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="203.63636363636363" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="203.63636363636363" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1900</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="285.45454545454544" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="285.45454545454544" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1905</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="367.27272727272725" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="367.27272727272725" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1910</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="449.09090909090907" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="449.09090909090907" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1915</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="530.9090909090909" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="530.9090909090909" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1920</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="612.7272727272727" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="612.7272727272727" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1925</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="694.5454545454545" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="694.5454545454545" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1930</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="720" y="776.3636363636363" width="0" height="0" rx="0" fill="" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="720" y="776.3636363636363" fill="#D23B22" fill-opacity="1" stroke="#D23B22" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="" text-anchor="" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="720" dy="0" text-anchor="">1935</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="0" y="786.3636363636363" width="770" height="150" rx="0" fill="#D23B22" fill-opacity="1" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="8" y="861.3636363636363" fill="#debdaaff" fill-opacity="1" stroke="#debdaaff" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="translate(0, 861.363636) scale(0.649, 1.2) translate(0, -861.363636)" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="500" font-style="" font-size="74px" font-family="ChunkFive, sans-serif" style="white-space: pre;">
      <tspan x="8" dy="0" text-anchor="start">CUBISM   AND   ABSTRACT   ART</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="158.219546" y="442.062032" width="91.094118" height="44.724138" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="249.31366400000002" y="464.424101" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="249.31366400000002" dy="0" text-anchor="end">DADAISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="249.31366400000002" y="486.78616999999997" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="249.31366400000002" dy="0" text-anchor="end">Zurich</tspan>
      <tspan x="249.31366400000002" dy="1em" text-anchor="end">Paris</tspan>
      <tspan x="249.31366400000002" dy="1em" text-anchor="end">Cologne</tspan>
      <tspan x="249.31366400000002" dy="1em" text-anchor="end">Berlin</tspan>
      <tspan x="249.31366400000002" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="158.219546" y="486.78616999999997" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="158.219546" dy="0" text-anchor="start">1916</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="249.31366400000002" y="442.062032" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="12px" font-family="Futura, sans serif">
      <tspan x="249.31366400000002" dy="0" text-anchor="end">(ABSTRACT)</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="89.5" y="367.913793" width="112.500031" height="34.086207" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="202.000031" y="384.95689649999997" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="202.000031" dy="0" text-anchor="end">EXPRESSIONISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="202.000031" y="402" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="202.000031" dy="0" text-anchor="end">Munich</tspan>
      <tspan x="202.000031" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="89.5" y="402" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="89.5" dy="0" text-anchor="start">1911</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="202.000031" y="367.913793" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="12px" font-family="Futura, sans serif">
      <tspan x="202.000031" dy="0" text-anchor="end">(ABSTRACT)</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="122.813468" y="590.434344" width="88.2" height="36.2" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="211.013468" y="608.534344" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="211.013468" dy="0" text-anchor="end">SURREALISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="211.013468" y="626.634344" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="211.013468" dy="0" text-anchor="end">Paris</tspan>
      <tspan x="211.013468" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="122.813468" y="626.634344" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="122.813468" dy="0" text-anchor="start">1924</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="211.013468" y="590.434344" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="12px" font-family="Futura, sans serif">
      <tspan x="211.013468" dy="0" text-anchor="end">(ABSTRACT)</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="270.420875" y="356.713805" width="69.333333" height="29.090909" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="339.754208" y="371.2592595" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="339.754208" dy="0" text-anchor="end">FUTURISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="339.754208" y="385.804714" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="339.754208" dy="0" text-anchor="end">Milan</tspan>
      <tspan x="339.754208" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="270.420875" y="385.804714" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="270.420875" dy="0" text-anchor="start">1910</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="281.023424" y="241.051724" width="63.13075" height="41.908046" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="344.15417399999995" y="262.005747" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="344.15417399999995" dy="0" text-anchor="end">FAUVISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="344.15417399999995" y="282.95977" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="344.15417399999995" dy="0" text-anchor="end">Paris</tspan>
      <tspan x="344.15417399999995" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="281.023424" y="282.95977" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="281.023424" dy="0" text-anchor="start">1905</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="84.385185" y="740.444444" width="240" height="23.2" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="324.385185" y="752.044444" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="324.385185" dy="0" text-anchor="end">NON-GEOMETRICAL ABSTRACT ART</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="324.385185" y="763.644444" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="324.385185" dy="0" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="416.599327" y="449.518519" width="114.19798" height="50.340067" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="530.797307" y="474.6885525" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="530.797307" dy="0" text-anchor="end">DE STIJL and</tspan>
      <tspan x="530.797307" dy="1em" text-anchor="end">NEOPLASTICISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="530.797307" y="499.858586" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="530.797307" dy="0" text-anchor="end">Leyden</tspan>
      <tspan x="530.797307" dy="1em" text-anchor="end">Berlin</tspan>
      <tspan x="530.797307" dy="1em" text-anchor="end">Paris</tspan>
      <tspan x="530.797307" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="416.599327" y="499.858586" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="416.599327" dy="0" text-anchor="start">1916</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="553.786833" y="568.328805" width="71.909122" height="28.181818" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="625.695955" y="582.419714" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="625.695955" dy="0" text-anchor="end">BAUHAUS</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="625.695955" y="596.510623" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="625.695955" dy="0" text-anchor="end">Weimar</tspan>
      <tspan x="625.695955" dy="1em" text-anchor="end">Dessau</tspan>
      <tspan x="625.695955" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="553.786833" y="596.510623" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="553.786833" dy="0" text-anchor="start">1919</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="596.169697" y="426.4" width="125.363636" height="40" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="721.5333330000001" y="446.4" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="721.5333330000001" dy="0" text-anchor="end">CONSTRUCTIVISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="721.5333330000001" y="466.4" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="721.5333330000001" dy="0" text-anchor="end">Moscow</tspan>
      <tspan x="721.5333330000001" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="596.169697" y="466.4" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="596.169697" dy="0" text-anchor="start">1914</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="493.727304" y="734.558923" width="208.181818" height="32.545455" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="701.909122" y="750.8316505" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="701.909122" dy="0" text-anchor="end">GEOMETRICAL ABSTRACT ART</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="701.909122" y="767.104378" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="701.909122" dy="0" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="556.142096" y="55.409613" width="148.757576" height="34.090909" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="704.899672" y="72.4550675" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="704.899672" dy="0" text-anchor="end">NEO-IMPRESSIONISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="704.899672" y="89.500522" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="704.899672" dy="0" text-anchor="end">Paris</tspan>
      <tspan x="704.899672" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="556.142096" y="89.500522" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="556.142096" dy="0" text-anchor="start">1886</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="438.55451" y="349" width="42.564263" height="36.839827" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="481.118773" y="367.4199135" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="10px" font-family="Futura, sans serif">
      <tspan x="481.118773" dy="0" text-anchor="end">ORPHISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="481.118773" y="385.839827" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="481.118773" dy="0" text-anchor="end">Paris</tspan>
      <tspan x="481.118773" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="438.55451" y="385.839827" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="438.55451" dy="0" text-anchor="start">1912</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="303.138046" y="482.222222" width="51.000031" height="32" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="354.13807699999995" y="498.222222" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="354.13807699999995" dy="0" text-anchor="end">PURISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="354.13807699999995" y="514.222222" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="354.13807699999995" dy="0" text-anchor="end">Paris</tspan>
      <tspan x="354.13807699999995" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="303.138046" y="514.222222" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="303.138046" dy="0" text-anchor="start">1918</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="579.787879" y="370.921212" width="99.090909" height="27.636364" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="678.878788" y="384.739394" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="678.878788" dy="0" text-anchor="end">SUPREMATISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="678.878788" y="398.55757600000004" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="678.878788" dy="0" text-anchor="end">Moscow</tspan>
      <tspan x="678.878788" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="579.787879" y="398.55757600000004" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="579.787879" dy="0" text-anchor="start">1913</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="284.666681" y="54.2" width="89.666667" height="31" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="374.333348" y="69.7" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="14px" font-family="Futura, sans serif">
      <tspan x="374.333348" dy="0" text-anchor="end">SYNTHETISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="374.333348" y="85.2" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="374.333348" dy="0" text-anchor="end">Pont-Aven</tspan>
      <tspan x="374.333348" dy="1em" text-anchor="end">Paris</tspan>
      <tspan x="374.333348" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="284.666681" y="85.2" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="284.666681" dy="0" text-anchor="start">1888</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="612.064814" y="275.523148" width="75" height="48.636364" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="687.064814" y="299.84132999999997" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="18px" font-family="Futura, sans serif">
      <tspan x="687.064814" dy="0" text-anchor="end">CUBISM</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="687.064814" y="324.159512" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="687.064814" dy="0" text-anchor="end">Paris</tspan>
      <tspan x="687.064814" dy="1em" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="612.064814" y="324.159512" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="612.064814" dy="0" text-anchor="start">1906</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="249.070707" y="600.747475" width="147.222222" height="49.555556" rx="0" fill="white" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="390.29292899999996" y="615.525253" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="Thin" font-style="" font-size="18px" font-family="Futura, sans serif">
      <tspan x="390.29292899999996" dy="0" text-anchor="end">MODERN</tspan>
      <tspan x="390.29292899999996" dy="1em" text-anchor="end">ARCHITECTURE</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="396.29292899999996" y="650.303031" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="" font-style="" font-size="9px" font-family="Futura">
      <tspan x="396.29292899999996" dy="0" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="80" y="37" width="119" height="21" rx="0" fill="white" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="139.5" y="47.5" fill="#D23B22" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="" font-style="" font-size="12px" font-family="Futura">
      <tspan x="139.5" dy="0" text-anchor="middle">JAPANESE PRINTS</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="101" y="238" width="130" height="21" rx="0" fill="white" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="166" y="248.5" fill="#D23B22" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="" font-style="" font-size="12px" font-family="Futura">
      <tspan x="166" dy="0" text-anchor="middle">NEAR-EASTERN ART</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="405" y="319.333333" width="139" height="21" rx="0" fill="white" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="474.5" y="329.833333" fill="#D23B22" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="" font-style="" font-size="12px" font-family="Futura">
      <tspan x="474.5" dy="0" text-anchor="middle">MACHINE ESTHETIC</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="397" y="238" width="132" height="20" rx="0" fill="white" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="463" y="248" fill="#D23B22" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="" font-style="" font-size="12px" font-family="Futura">
      <tspan x="463" dy="0" text-anchor="middle">AFRICAN FIRST ART</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="296.680251" y="413.927203" width="54" height="28" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="323.680251" y="427.927203" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="100" font-style="" font-size="14px" font-family="Futura">
      <tspan x="323.680251" dy="0" text-anchor="middle">Brancusi</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="350.680251" y="441.927203" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="350.680251" dy="0" text-anchor="end">Paris</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="445.888889" y="26.333333" width="84" height="33" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="487.888889" y="42.833332999999996" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="100" font-style="" font-size="14px" font-family="Futura">
      <tspan x="487.888889" dy="0" text-anchor="middle">Cézanne</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="529.8888890000001" y="59.333332999999996" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="529.8888890000001" dy="0" text-anchor="end">Provence</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="445.888889" y="59.333332999999996" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="445.888889" dy="0" text-anchor="start">d. 1906</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="293" y="19" width="74" height="31" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="330" y="34.5" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="100" font-style="" font-size="14px" font-family="Futura">
      <tspan x="330" dy="0" text-anchor="middle">Gauguin</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="367" y="50" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="367" dy="0" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="293" y="50" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="293" dy="0" text-anchor="start">d. 1903</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="57" y="116" width="73" height="30" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="93.5" y="131" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="100" font-style="" font-size="14px" font-family="Futura">
      <tspan x="93.5" dy="0" text-anchor="middle">Redon</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="130" y="146" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="130" dy="0" text-anchor="end">Paris</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="57" y="146" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="57" dy="0" text-anchor="start">d. 1916</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="662.222222" y="126.555556" width="69" height="30" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="696.722222" y="141.555556" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="100" font-style="" font-size="14px" font-family="Futura">
      <tspan x="696.722222" dy="0" text-anchor="middle">Rousseau</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="731.222222" y="156.555556" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="731.222222" dy="0" text-anchor="end">Paris</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="662.222222" y="156.555556" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="662.222222" dy="0" text-anchor="start">d. 1910</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="655" y="17" width="47" height="30" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="678.5" y="32" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="100" font-style="" font-size="14px" font-family="Futura">
      <tspan x="678.5" dy="0" text-anchor="middle">Seurat</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="702" y="47" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="702" dy="0" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="655" y="47" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="655" dy="0" text-anchor="start">d. 1891</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <rect class="mat-mdc-tooltip-trigger mat-mdc-tooltip-disabled" x="156" y="75" width="70" height="30" rx="0" fill="white" fill-opacity="0" stroke="" stroke-opacity="0" stroke-width="0" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <text class="click-through" x="191" y="90" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="middle" font-weight="100" font-style="" font-size="14px" font-family="Futura">
      <tspan x="191" dy="0" text-anchor="middle">Van Gogh</tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="226" y="105" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="end" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="226" dy="0" text-anchor="end">
      </tspan>
      <!---->
    </text>
    <!---->
    <text class="click-through" x="156" y="105" fill="#343434" fill-opacity="1" stroke="" stroke-opacity="0" stroke-dasharray="" transform="" writing-mode="" dominant-baseline="central" text-anchor="start" font-weight="100" font-style="" font-size="10px" font-family="Futura">
      <tspan x="156" dy="0" text-anchor="start">d. 1890</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <g>
    <line class="moveable-line" x1="523.2198587831674" y1="601.5548986426808" x2="396.29292899999996" y2="622.9952603336228" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
    <line class="hit-area" x1="523.2198587831674" y1="601.5548986426808" x2="396.29292899999996" y2="622.9952603336228" stroke="transparent" stroke-width="13"/>
    <!---->
  </g>
  <!---->
  <path d="M 395.07042184026204 622.1260789698083 L 403.20847382342265 627.9120926263935 M 395.4237476361855 624.2177674933607 L 401.20976129277057 616.0797155102001" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <path d="M 395.07042184026204 622.1260789698083 L 403.20847382342265 627.9120926263935 M 395.4237476361855 624.2177674933607 L 401.20976129277057 616.0797155102001" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <g>
  <line class="moveable-line" x1="311.96463444307693" y1="459.75780121639093" x2="171.47595164521414" y2="732.4438408963499" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
  <line class="hit-area" x1="311.96463444307693" y1="459.75780121639093" x2="171.47595164521414" y2="732.4438408963499" stroke="transparent" stroke-width="13"/>
  <!---->
</g>
<!---->
<path d="M 170.0472968029975 732.9009462896935 L 179.55764386259727 729.8580623088662 M 171.9330570385577 733.8724957385665 L 168.89017305773046 724.3621486789668" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 170.0472968029975 732.9009462896935 L 179.55764386259727 729.8580623088662 M 171.9330570385577 733.8724957385665 L 168.89017305773046 724.3621486789668" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="637.3653967983341" y1="488.6880992027315" x2="600.977591642526" y2="560.3050892870166" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="637.3653967983341" y1="488.6880992027315" x2="600.977591642526" y2="560.3050892870166" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 599.5515367009787 560.7702422752915 L 609.0445766007564 557.6737866279358 M 601.4427446308009 561.731144228564 L 598.3462889834452 552.2381043287862" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 599.5515367009787 560.7702422752915 L 609.0445766007564 557.6737866279358 M 601.4427446308009 561.731144228564 L 598.3462889834452 552.2381043287862" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="649.9275779164958" y1="490.912238503651" x2="675.1302992218403" y2="725.6103690305822" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="649.9275779164958" y1="490.912238503651" x2="675.1302992218403" y2="725.6103690305822" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 674.1889484270991 726.7782126331475 L 680.4553834667403 719.004047982925 M 676.2981428244057 726.5517198253234 L 668.5239781741832 720.2852847856823" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 674.1889484270991 726.7782126331475 L 680.4553834667403 719.004047982925 M 676.2981428244057 726.5517198253234 L 668.5239781741832 720.2852847856823" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="681.7358981508892" y1="348.1075091366596" x2="702.6125135898865" y2="417.7787210308025" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="681.7358981508892" y1="348.1075091366596" x2="702.6125135898865" y2="417.7787210308025" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 701.9009334545873 419.0991958355074 L 706.63781870317 410.30898751848935 M 703.9329883945913 418.4903011661017 L 695.1427800775733 413.753415917519" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 701.9009334545873 419.0991958355074 L 706.63781870317 410.30898751848935 M 703.9329883945913 418.4903011661017 L 695.1427800775733 413.753415917519" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="607.3758977683208" y1="341.7857307200204" x2="475.8523705540136" y2="443.9959445778426" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="607.3758977683208" y1="341.7857307200204" x2="475.8523705540136" y2="443.9959445778426" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 474.3640295659788 443.8092868639693 L 484.27169859909316 445.05184006004555 M 475.6657128401403 445.48428556587737 L 476.90826603621656 435.576616532763" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 474.3640295659788 443.8092868639693 L 484.27169859909316 445.05184006004555 M 475.6657128401403 445.48428556587737 L 476.90826603621656 435.576616532763" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="587.1052352484409" y1="298.06548115639725" x2="379.9973099724562" y2="309.8571075922562" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="587.1052352484409" y1="298.06548115639725" x2="379.9973099724562" y2="309.8571075922562" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="364.54389021821765" y1="318.65504110387116" x2="343.94645210651663" y2="349.24818235460003" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="364.54389021821765" y1="318.65504110387116" x2="343.94645210651663" y2="349.24818235460003" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 379.9973099724562,309.8571075922562 A 20,20 0 0,0 364.54389021821765,318.65504110387116" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 342.4742556805543 349.5356500609542 L 352.2744527170063 347.6220194378898 M 344.2339198128708 350.7203787805624 L 342.3202891898064 340.9201817441104" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 342.4742556805543 349.5356500609542 L 352.2744527170063 347.6220194378898 M 344.2339198128708 350.7203787805624 L 342.3202891898064 340.9201817441104" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="588.5141395625212" y1="321.58674632821806" x2="489.19340967059094" y2="370.4799833222666" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="588.5141395625212" y1="321.58674632821806" x2="489.19340967059094" y2="370.4799833222666" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 487.7733520390493 369.9968308391795 L 497.2264687202205 373.2131065000916 M 488.7102571875038 371.90004095380823 L 491.92653284841595 362.44692427263703" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 487.7733520390493 369.9968308391795 L 497.2264687202205 373.2131065000916 M 488.7102571875038 371.90004095380823 L 491.92653284841595 362.44692427263703" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="590.7996477324307" y1="336.1670211799171" x2="463.7219131201865" y2="432.5367742506763" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="590.7996477324307" y1="336.1670211799171" x2="463.7219131201865" y2="432.5367742506763" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="459.6516182943686" y1="434.92476262694885" x2="361.79353685632486" y2="477.7271628482337" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="459.6516182943686" y1="434.92476262694885" x2="361.79353685632486" y2="477.7271628482337" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 463.7219131201865,432.5367742506763 A 20,20 0 0,1 459.6516182943686,434.92476262694885" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 360.39672197034645 477.18043789624966 L 369.69511508191056 480.8199062169841 M 361.2468119043408 479.12397773421213 L 364.88628022507527 469.825584622648" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 360.39672197034645 477.18043789624966 L 369.69511508191056 480.8199062169841 M 361.2468119043408 479.12397773421213 L 364.88628022507527 469.825584622648" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="637.9744141356682" y1="348.4785273354901" x2="634.7125582384174" y2="362.1663664792236" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="637.9744141356682" y1="348.4785273354901" x2="634.7125582384174" y2="362.1663664792236" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 633.4349160765461 362.9522611954684 L 641.9399937330367 357.7206746128076 M 635.4984529546622 363.44400864109497 L 630.2668663720015 354.93893098460427" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 633.4349160765461 362.9522611954684 L 641.9399937330367 357.7206746128076 M 635.4984529546622 363.44400864109497 L 630.2668663720015 354.93893098460427" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="508.38553972224497" y1="75.4158164299403" x2="629.4747202607863" y2="267.90512953318614" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="508.38553972224497" y1="75.4158164299403" x2="629.4747202607863" y2="267.90512953318614" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 629.1417019081383 269.36769535242723 L 631.358556544619 259.631607861267 M 630.9372860800274 268.2381478858341 L 621.2011985888671 266.0212932493535" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 629.1417019081383 269.36769535242723 L 631.358556544619 259.631607861267 M 630.9372860800274 268.2381478858341 L 621.2011985888671 266.0212932493535" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="187.39448083649913" y1="528.4455380795159" x2="176.00176339910925" y2="572.995538757412" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="187.39448083649913" y1="528.4455380795159" x2="176.00176339910925" y2="572.995538757412" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 174.71138723446796 573.760345499366 L 183.30123329142972 568.6691384880071 M 176.76657014106337 574.2859149220533 L 171.67536312970432 565.6960688650915" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 174.71138723446796 573.760345499366 L 183.30123329142972 568.6691384880071 M 176.76657014106337 574.2859149220533 L 171.67536312970432 565.6960688650915" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="526.0677559122878" y1="524.5797456539959" x2="567.9675149067576" y2="562.2054509373653" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="526.0677559122878" y1="524.5797456539959" x2="567.9675149067576" y2="562.2054509373653" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 568.0480154298747 563.7032892679573 L 567.5121351803563 553.7323978096679 M 569.4653532373496 562.1249504142482 L 559.4944617790602 562.6608306637667" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 568.0480154298747 563.7032892679573 L 567.5121351803563 553.7323978096679 M 569.4653532373496 562.1249504142482 L 559.4944617790602 562.6608306637667" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="500.1804674677454" y1="533.6062869449705" x2="586.8142910850635" y2="726.3500227701423" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="500.1804674677454" y1="533.6062869449705" x2="586.8142910850635" y2="726.3500227701423" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 586.2816993632362 727.7522873873434 L 589.8270848307214 718.4176162093233 M 588.2165557022645 726.8826144919697 L 578.8818845242445 723.3372290244844" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 586.2816993632362 727.7522873873434 L 589.8270848307214 718.4176162093233 M 588.2165557022645 726.8826144919697 L 578.8818845242445 723.3372290244844" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="422.3197504371509" y1="526.0059483347411" x2="376.4275515482314" y2="600.747475" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="422.3197504371509" y1="526.0059483347411" x2="376.4275515482314" y2="600.747475" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 374.96868904480027 601.0963592141677 L 384.68012410236815 598.7738878499838 M 376.77643576239916 602.2063375034312 L 374.45396439821513 592.4949024458633" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 374.96868904480027 601.0963592141677 L 384.68012410236815 598.7738878499838 M 376.77643576239916 602.2063375034312 L 374.45396439821513 592.4949024458633" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="98.9447745817297" y1="426.9950318571717" x2="92.87427466274842" y2="731.4462325314182" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="98.9447745817297" y1="426.9950318571717" x2="92.87427466274842" y2="731.4462325314182" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 91.79268080608779 732.4855374568597 L 98.99269346863811 725.5670360458655 M 93.91357958818985 732.5278263880789 L 86.99507817719571 725.3278137255286" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 91.79268080608779 732.4855374568597 L 98.99269346863811 725.5670360458655 M 93.91357958818985 732.5278263880789 L 86.99507817719571 725.3278137255286" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="281.13031250173003" y1="296.1994066896503" x2="183.3669725690725" y2="357.2352360124994" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="281.13031250173003" y1="296.1994066896503" x2="183.3669725690725" y2="357.2352360124994" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 181.90555038219833 356.89723425137265 L 191.63402487937495 359.1472627117697 M 183.0289708079457 358.69665819937353 L 185.2789992683428 348.96818370219694" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 181.90555038219833 356.89723425137265 L 191.63402487937495 359.1472627117697 M 183.0289708079457 358.69665819937353 L 185.2789992683428 348.96818370219694" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="270.8658708392752" y1="402.726166864912" x2="241.33639937389304" y2="429.8785859372633" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="270.8658708392752" y1="402.726166864912" x2="241.33639937389304" y2="429.8785859372633" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 239.83771665319472 429.8157361043831 L 249.8142290933528 430.2341182815653 M 241.27354954101284 431.3772686579617 L 241.69193171819504 421.4007562178036" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 239.83771665319472 429.8157361043831 L 249.8142290933528 430.2341182815653 M 241.27354954101284 431.3772686579617 L 241.69193171819504 421.4007562178036" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="245.51274340110436" y1="376.38067677810534" x2="219.93388575120488" y2="378.57959322772433" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="245.51274340110436" y1="376.38067677810534" x2="219.93388575120488" y2="378.57959322772433" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 218.7862773979535 377.61367653030135 L 226.42573894276057 384.04364320363857 M 218.96796905378193 379.7272015809757 L 225.39793572711912 372.08774003616867" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 218.7862773979535 377.61367653030135 L 226.42573894276057 384.04364320363857 M 218.96796905378193 379.7272015809757 L 225.39793572711912 372.08774003616867" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="101.34308794338061" y1="58" x2="159.44692188419071" y2="157.23075222128477" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="101.34308794338061" y1="58" x2="159.44692188419071" y2="157.23075222128477" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="165.2499546841467" y1="163.51885082905412" x2="273.64614439366113" y2="239.26498998125635" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="165.2499546841467" y1="163.51885082905412" x2="273.64614439366113" y2="239.26498998125635" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 159.44692188419071,157.23075222128477 A 20,20 0 0,0 165.2499546841467,163.51885082905412" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 273.908022889255 240.7419529639876 L 272.1647359126196 230.91002565384608 M 275.12310737639234 239.00311148566252 L 265.29118006625083 240.7463984622979" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 273.908022889255 240.7419529639876 L 272.1647359126196 230.91002565384608 M 275.12310737639234 239.00311148566252 L 265.29118006625083 240.7463984622979" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="199" y1="54.452104732602535" x2="275.7274933818241" y2="63.41710590672873" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="199" y1="54.452104732602535" x2="275.7274933818241" y2="63.41710590672873" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 276.65789437231115 64.59369164338948 L 270.464350583667 56.76133188065132 M 276.90407911848484 62.48670491624167 L 269.07171935574667 68.68024870488586" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 276.65789437231115 64.59369164338948 L 270.464350583667 56.76133188065132 M 276.90407911848484 62.48670491624167 L 269.07171935574667 68.68024870488586" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="526.0789932162817" y1="340.333333" x2="546.6567146872432" y2="457.0438379271026" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="526.0789932162817" y1="340.333333" x2="546.6567146872432" y2="457.0438379271026" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="547.3110747418176" y1="459.687359322573" x2="579.5767213606489" y2="560.1407586098609" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="547.3110747418176" y1="459.687359322573" x2="579.5767213606489" y2="560.1407586098609" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 546.6567146872432,457.0438379271026 A 20,20 0 0,0 547.3110747418176,459.687359322573" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 578.8912380732072 561.4749672372693 L 583.4544004091922 552.5933348642051 M 580.9109299880573 560.8262418973026 L 572.0292976149931 556.2630795613176" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 578.8912380732072 561.4749672372693 L 583.4544004091922 552.5933348642051 M 580.9109299880573 560.8262418973026 L 572.0292976149931 556.2630795613176" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="516.2663370571294" y1="340.333333" x2="568.2976689918963" y2="417.69248707215957" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="516.2663370571294" y1="340.333333" x2="568.2976689918963" y2="417.69248707215957" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="572.929324375767" y1="422.5575516306669" x2="588.9575109935478" y2="434.5221858639606" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="572.929324375767" y1="422.5575516306669" x2="588.9575109935478" y2="434.5221858639606" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 568.2976689918963,417.69248707215957 A 20,20 0 0,0 572.929324375767,422.5575516306669" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 589.1729985234784 436.0066268776951 L 587.738529443825 426.12491940508033 M 590.4419520072823 434.3066983340299 L 580.5602445346675 435.74116741368346" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 589.1729985234784 436.0066268776951 L 587.738529443825 426.12491940508033 M 590.4419520072823 434.3066983340299 L 580.5602445346675 435.74116741368346" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="501.7306089956303" y1="340.333333" x2="519.3220025817496" y2="440.6537752055967" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="501.7306089956303" y1="340.333333" x2="519.3220025817496" y2="440.6537752055967" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 518.4604762969528 441.88168907041654 L 524.1955312082379 433.70764543954715 M 520.5499164465696 441.51530149039354 L 512.3758728157001 435.7802465791084" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 518.4604762969528 441.88168907041654 L 524.1955312082379 433.70764543954715 M 520.5499164465696 441.51530149039354 L 512.3758728157001 435.7802465791084" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="405" y1="332.05072673313003" x2="347.3344401137955" y2="368.9602117814845" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="405" y1="332.05072673313003" x2="347.3344401137955" y2="368.9602117814845" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 345.8693096474855 368.63866440472157 L 355.62246962038836 370.7791584266189 M 347.0128927370326 370.42534224779445 L 349.1533867589299 360.6721822748916" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 345.8693096474855 368.63866440472157 L 355.62246962038836 370.7791584266189 M 347.0128927370326 370.42534224779445 L 349.1533867589299 360.6721822748916" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="438.1564309064393" y1="340.333333" x2="378.872694534692" y2="509.71367086493444" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="438.1564309064393" y1="340.333333" x2="378.872694534692" y2="509.71367086493444" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="390.46356155499" y1="534.9462647162202" x2="439.87998396690193" y2="554.2778992367005" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="390.46356155499" y1="534.9462647162202" x2="439.87998396690193" y2="554.2778992367005" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="451.2358496520822" y1="565.6597739056336" x2="513.6041838090745" y2="726.1699580593947" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="2"/>
<line class="hit-area" x1="451.2358496520822" y1="565.6597739056336" x2="513.6041838090745" y2="726.1699580593947" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 378.872694534692,509.71367086493444 A 20,20 0 0,0 390.46356155499,534.9462647162202" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 439.87998396690193,554.2778992367005 A 20,20 0 0,1 451.2358496520822,565.6597739056336" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 512.9996873262598 727.5427595749379 L 517.0237323075894 718.404219970436 M 514.9769853246177 726.7744545422094 L 505.83844572011577 722.7504095608798" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 512.9996873262598 727.5427595749379 L 517.0237323075894 718.404219970436 M 514.9769853246177 726.7744545422094 L 505.83844572011577 722.7504095608798" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="430.0054405978366" y1="340.333333" x2="348.39023453574595" y2="475.4195189300348" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="430.0054405978366" y1="340.333333" x2="348.39023453574595" y2="475.4195189300348" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 346.9339138086499 475.77886547397696 L 356.6284286316166 473.38674790535896 M 348.7495810796881 476.8758396571309 L 346.3574635110701 467.1813248341642" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 346.9339138086499 475.77886547397696 L 356.6284286316166 473.38674790535896 M 348.7495810796881 476.8758396571309 L 346.3574635110701 467.1813248341642" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="231" y1="254.48868099737962" x2="272.0613811807766" y2="258.2718119697408" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="231" y1="254.48868099737962" x2="272.0613811807766" y2="258.2718119697408" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 273.0202578053614 259.4253090948511 L 266.6371558708245 251.74664685406177 M 273.21487830588694 257.312935345156 L 265.53621606509756 263.6960372796929" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 273.0202578053614 259.4253090948511 L 266.6371558708245 251.74664685406177 M 273.21487830588694 257.312935345156 L 265.53621606509756 263.6960372796929" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="634.0349471065396" y1="114.41330175023931" x2="646.7769975642727" y2="266.55454728991384" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="634.0349471065396" y1="114.41330175023931" x2="646.7769975642727" y2="266.55454728991384" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 645.8085597482736 267.7000289333092 L 652.2553091405214 260.07472458604764 M 647.9224792076681 267.5229851059129 L 640.2971748604066 261.07623571366514" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 645.8085597482736 267.7000289333092 L 652.2553091405214 260.07472458604764 M 647.9224792076681 267.5229851059129 L 640.2971748604066 261.07623571366514" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="551.0643481479744" y1="103.53021326551493" x2="351.88454443015394" y2="238.57767666936726" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="551.0643481479744" y1="103.53021326551493" x2="351.88454443015394" y2="238.57767666936726" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 350.41141870655537 238.29500972527055 L 360.2178019419276 240.17668237369372 M 351.6018774860572 240.05080239296583 L 353.48355013448037 230.24441915759365" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 350.41141870655537 238.29500972527055 L 360.2178019419276 240.17668237369372 M 351.6018774860572 240.05080239296583 L 353.48355013448037 230.24441915759365" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="593.541326781214" y1="106.4087079917609" x2="593.9836112414657" y2="110.84470878232177" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="593.541326781214" y1="106.4087079917609" x2="593.9836112414657" y2="110.84470878232177" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="593.3481431644844" y1="118.19798636627834" x2="548.7829234395505" y2="278.11225577594706" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="593.3481431644844" y1="118.19798636627834" x2="548.7829234395505" y2="278.11225577594706" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="530.5837474877877" y1="292.7147430467705" x2="363.67387290488597" y2="301.62942994960076" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="2"/>
<line class="hit-area" x1="530.5837474877877" y1="292.7147430467705" x2="363.67387290488597" y2="301.62942994960076" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="349.5129462735515" y1="308.6348587116108" x2="314.4091158130337" y2="349.8613805297644" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="3"/>
<line class="hit-area" x1="349.5129462735515" y1="308.6348587116108" x2="314.4091158130337" y2="349.8613805297644" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 593.9836112414657,110.84470878232177 A 20,20 0 0,1 593.3481431644844,118.19798636627834" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 548.7829234395505,278.11225577594706 A 20,20 0 0,1 530.5837474877877,292.7147430467705" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<path d="M 363.67387290488597,301.62942994960076 A 20,20 0 0,0 349.5129462735515,308.6348587116108" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 312.9139182527643 349.981314905893 L 322.8672304879055 349.18292924432205 M 314.52905018916226 351.3565780900338 L 313.73066452759133 341.4032658548926" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 312.9139182527643 349.981314905893 L 322.8672304879055 349.18292924432205 M 314.52905018916226 351.3565780900338 L 313.73066452759133 341.4032658548926" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="326.7210351506834" y1="539.194902981707" x2="323.84111859164847" y2="600.747475" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="326.7210351506834" y1="539.194902981707" x2="323.84111859164847" y2="600.747475" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 322.73204561852987 601.7574042743112 L 330.11498275428676 595.0344522914189 M 324.85104786595963 601.8565479731186 L 318.12809588306743 594.4736108373617" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 322.73204561852987 601.7574042743112 L 330.11498275428676 595.0344522914189 M 324.85104786595963 601.8565479731186 L 318.12809588306743 594.4736108373617" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="100.41508345488857" y1="164.61019353653586" x2="76.5373197620361" y2="232.15704936642985" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="7 3" segment-number="0"/>
<line class="hit-area" x1="100.41508345488857" y1="164.61019353653586" x2="76.5373197620361" y2="232.15704936642985" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="78.21929654380226" y1="249.0715012647896" x2="138.6160990742351" y2="350.2830833338081" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="7 3" segment-number="1"/>
<line class="hit-area" x1="78.21929654380226" y1="249.0715012647896" x2="138.6160990742351" y2="350.2830833338081" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 76.5373197620361,232.15704936642985 A 20,20 0 0,0 78.21929654380226,249.0715012647896" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="7 3"/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 138.24880135102583 351.7374190520445 L 140.69384876114856 342.05611814374373 M 140.07043479247147 350.6503810570174 L 130.38913388417072 348.2053336468947" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="7 3"/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 138.24880135102583 351.7374190520445 L 140.69384876114856 342.05611814374373 M 140.07043479247147 350.6503810570174 L 130.38913388417072 348.2053336468947" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="7 3"/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="686.8283873183386" y1="174.76461788089918" x2="659.3795390332394" y2="266.89780289852143" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="686.8283873183386" y1="174.76461788089918" x2="659.3795390332394" y2="266.89780289852143" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 658.0601888529949 267.6114660548504 L 666.8429107088934 262.86071443887073 M 660.0932021895684 268.217153078766 L 655.3424505735887 259.4344312228674" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 658.0601888529949 267.6114660548504 L 666.8429107088934 262.86071443887073 M 660.0932021895684 268.217153078766 L 655.3424505735887 259.4344312228674" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="625.9995933293734" y1="423.46545445245346" x2="616.8798729700372" y2="417.3853451282747" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="625.9995933293734" y1="423.46545445245346" x2="616.8798729700372" y2="417.3853451282747" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="603.3495304639723" y1="421.3393496715605" x2="585.8383256940951" y2="461.17713599106156" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="603.3495304639723" y1="421.3393496715605" x2="585.8383256940951" y2="461.17713599106156" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="590.3926078716938" y1="483.7441815589377" x2="640.0318677521412" y2="530.7713406673045" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="2"/>
<line class="hit-area" x1="590.3926078716938" y1="483.7441815589377" x2="640.0318677521412" y2="530.7713406673045" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="646.2292071533686" y1="543.909960803996" x2="658.7977227956571" y2="725.5803843209674" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="3"/>
<line class="hit-area" x1="646.2292071533686" y1="543.909960803996" x2="658.7977227956571" y2="725.5803843209674" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 616.8798729700372,417.3853451282747 A 20,20 0 0,0 603.3495304639723,421.3393496715605" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 585.8383256940951,461.17713599106156 A 20,20 0 0,0 590.3926078716938,483.7441815589377" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<path d="M 640.0318677521412,530.7713406673045 A 20,20 0 0,1 646.2292071533686,543.909960803996" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 657.8127965997662 726.7117199862148 L 664.3693067337887 719.180583353722 M 659.9290584609046 726.5653105168583 L 652.3979218284118 720.0088003828358" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 657.8127965997662 726.7117199862148 L 664.3693067337887 719.180583353722 M 659.9290584609046 726.5653105168583 L 652.3979218284118 720.0088003828358" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="177.95549331588734" y1="650.8233631363669" x2="150.67932593398126" y2="731.9140871375552" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="177.95549331588734" y1="650.8233631363669" x2="150.67932593398126" y2="731.9140871375552" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 149.33586083367075 732.5812464311729 L 158.2791121977225 728.14006425137 M 151.34648522759895 733.2575522378656 L 146.90530304779602 724.3143008738139" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 149.33586083367075 732.5812464311729 L 158.2791121977225 728.14006425137 M 151.34648522759895 733.2575522378656 L 146.90530304779602 724.3143008738139" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="325.42132142271345" y1="116.08082454873316" x2="315.2198894863968" y2="232.08632332456136" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="325.42132142271345" y1="116.08082454873316" x2="315.2198894863968" y2="232.08632332456136" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 314.07039184491316 233.0499908870726 L 321.72243010612056 226.63499637703356 M 316.1835570489081 233.23582096604503 L 309.768562538869 225.58378270483763" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 314.07039184491316 233.0499908870726 L 321.72243010612056 226.63499637703356 M 316.1835570489081 233.23582096604503 L 309.768562538869 225.58378270483763" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="212.57071214188937" y1="120.51503498515231" x2="292.58157105987533" y2="233.70249690176996" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="212.57071214188937" y1="120.51503498515231" x2="292.58157105987533" y2="233.70249690176996" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 292.32770244484055 235.180857725491 L 294.01766881402636 225.33962519161426 M 294.0599318835964 233.95636551680474 L 284.21869934971966 232.26639914761893" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 292.32770244484055 235.180857725491 L 294.01766881402636 225.33962519161426 M 294.0599318835964 233.95636551680474 L 284.21869934971966 232.26639914761893" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="164.4418168469045" y1="259" x2="150.92142050148865" y2="350.1087771324163" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="164.4418168469045" y1="259" x2="150.92142050148865" y2="350.1087771324163" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 149.71655464333014 351.002252516888 L 157.737171053079 345.05451710561755 M 151.8148958859603 351.3136429905748 L 145.86716047468988 343.293026580826" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 149.71655464333014 351.002252516888 L 157.737171053079 345.05451710561755 M 151.8148958859603 351.3136429905748 L 145.86716047468988 343.293026580826" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="414.4113236686246" y1="340.333333" x2="352.7499086391214" y2="409.0201617384708" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="7 3" segment-number="0"/>
<line class="hit-area" x1="414.4113236686246" y1="340.333333" x2="352.7499086391214" y2="409.0201617384708" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 351.25208259060565 409.1008904644644 L 361.2228922889033 408.56349110165735 M 352.83063736511497 410.5179877869865 L 352.29323800230793 400.54717808868884" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="7 3"/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 351.25208259060565 409.1008904644644 L 361.2228922889033 408.56349110165735 M 352.83063736511497 410.5179877869865 L 352.29323800230793 400.54717808868884" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="7 3"/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="529" y1="249.5093395413728" x2="603.3933658474177" y2="287.01152984134893" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="529" y1="249.5093395413728" x2="603.3933658474177" y2="287.01152984134893" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 603.8630375467775 288.43610293306216 L 600.7365014983271 278.9529274916526 M 604.817938939131 286.54185814198917 L 595.3347634977214 289.6683941904395" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 603.8630375467775 288.43610293306216 L 600.7365014983271 278.9529274916526 M 604.817938939131 286.54185814198917 L 595.3347634977214 289.6683941904395" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="114.89653152903436" y1="426.94285124228776" x2="121.4967089107211" y2="524.3883181078097" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="114.89653152903436" y1="426.94285124228776" x2="121.4967089107211" y2="524.3883181078097" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="139.8176926013392" y1="542.9699714704111" x2="544.8168943807193" y2="576.1550198264129" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="139.8176926013392" y1="542.9699714704111" x2="544.8168943807193" y2="576.1550198264129" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 121.4967089107211,524.3883181078097 A 20,20 0 0,0 139.8176926013392,542.9699714704111" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 545.7873931146404 577.2987559001075 L 539.3269244921626 569.6850715559287 M 545.960630454414 575.1845210924918 L 538.3469461102352 581.6449897149696" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 545.7873931146404 577.2987559001075 L 539.3269244921626 569.6850715559287 M 545.960630454414 575.1845210924918 L 538.3469461102352 581.6449897149696" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="126.96600738535592" y1="426.9836545991781" x2="132.22868475054673" y2="572.4461126885918" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="126.96600738535592" y1="426.9836545991781" x2="132.22868475054673" y2="572.4461126885918" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 131.20706652582905 573.544427853623 L 138.0078301479244 566.233103878364 M 133.32699991557794 573.4677309133095 L 126.01567594031891 566.6669672912142" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 131.20706652582905 573.544427853623 L 138.0078301479244 566.233103878364 M 133.32699991557794 573.4677309133095 L 126.01567594031891 566.6669672912142" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="427.0953444307178" y1="58.50220506287683" x2="383.2355533222916" y2="65.02118776995367" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="427.0953444307178" y1="58.50220506287683" x2="383.2355533222916" y2="65.02118776995367" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 382.03048299647025 64.12798817935945 L 390.0524605176496 70.07388767098443 M 382.34235373169736 66.22625809577501 L 388.28825322332233 58.204280574595664" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 382.03048299647025 64.12798817935945 L 390.0524605176496 70.07388767098443 M 382.34235373169736 66.22625809577501 L 388.28825322332233 58.204280574595664" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="462.82409404927444" y1="74.17108996159388" x2="334.96991504836916" y2="234.02331280766606" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="462.82409404927444" y1="74.17108996159388" x2="334.96991504836916" y2="234.02331280766606" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 333.47910689540015 234.18911705928227 L 343.4031994858945 233.08538232207948 M 335.1357192999854 235.51412096063504 L 334.0319845627826 225.59002837014071" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 333.47910689540015 234.18911705928227 L 343.4031994858945 233.08538232207948 M 335.1357192999854 235.51412096063504 L 334.0319845627826 225.59002837014071" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="535.7890225104338" y1="319.333333" x2="603.2734672829822" y2="304.543431091969" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="535.7890225104338" y1="319.333333" x2="603.2734672829822" y2="304.543431091969" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 604.5366027312824 305.352435936244 L 596.1280941526958 299.967008599565 M 604.0824721272571 303.2802956436688 L 598.6970447905782 311.6888042222554" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 604.5366027312824 305.352435936244 L 596.1280941526958 299.967008599565 M 604.0824721272571 303.2802956436688 L 598.6970447905782 311.6888042222554" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="544" y1="333.75423308030724" x2="597.4442445350307" y2="366.2458540204047" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="544" y1="333.75423308030724" x2="597.4442445350307" y2="366.2458540204047" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 597.7995612629803 367.7031631920124 L 595.434269592083 358.00206843799646 M 598.9015537066384 365.89053729245506 L 589.2004589526225 368.2558289633525" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 597.7995612629803 367.7031631920124 L 595.434269592083 358.00206843799646 M 598.9015537066384 365.89053729245506 L 589.2004589526225 368.2558289633525" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="437.14698376984154" y1="340.333333" x2="348.08929029236816" y2="583.8425682212782" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="437.14698376984154" y1="340.333333" x2="348.08929029236816" y2="583.8425682212782" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<!---->
<path d="M 346.7288478912897 584.4743881894342 L 355.78511467212866 580.2684547485574 M 348.72111026052426 585.2030106223566 L 344.5151768196474 576.1467438415177" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 346.7288478912897 584.4743881894342 L 355.78511467212866 580.2684547485574 M 348.72111026052426 585.2030106223566 L 344.5151768196474 576.1467438415177" fill="" fill-opacity="0" stroke="#D23B22" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="621.6562258097946" y1="423.07075952641605" x2="607.3026796715873" y2="411.788681230576" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="0"/>
<line class="hit-area" x1="621.6562258097946" y1="423.07075952641605" x2="607.3026796715873" y2="411.788681230576" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="587.6728355558912" y1="417.98431036642904" x2="573.1931922303008" y2="465.9480476003108" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="1"/>
<line class="hit-area" x1="587.6728355558912" y1="417.98431036642904" x2="573.1931922303008" y2="465.9480476003108" stroke="transparent" stroke-width="13"/>
<line class="moveable-line" x1="572.9111908361272" y1="476.4748664632646" x2="593.2166099570652" y2="559.5859561950784" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" segment-number="2"/>
<line class="hit-area" x1="572.9111908361272" y1="476.4748664632646" x2="593.2166099570652" y2="559.5859561950784" stroke="transparent" stroke-width="13"/>
<!---->
</g>
<path d="M 607.3026796715873,411.788681230576 A 20,20 0 0,0 587.6728355558912,417.98431036642904" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 573.1931922303008,465.9480476003108 A 20,20 0 0,0 572.9111908361272,476.4748664632646" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 592.437988083424 560.8680435465097 L 597.6211604128824 552.3333749110003 M 594.4986973084965 560.3645780687197 L 585.9640286729871 555.1814057392612" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 592.437988083424 560.8680435465097 L 597.6211604128824 552.3333749110003 M 594.4986973084965 560.3645780687197 L 585.9640286729871 555.1814057392612" fill="" fill-opacity="0" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path class="click-through" d="M 158.219546 516.786170 A 45.547059 22.362069 0 0 0 249.313664 516.786170" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 89.500000 414.000000 A 56.250016 17.043104 0 0 0 202.000031 414.000000" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 122.813468 638.634344 A 44.100000 18.100000 0 0 0 211.013468 638.634344" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 270.420875 397.804714 A 34.666666 14.545454 0 0 0 339.754208 397.804714" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 281.023424 294.959770 A 31.565375 20.954023 0 0 0 344.154174 294.959770" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 416.599327 523.858586 A 57.098990 25.170033 0 0 0 530.797307 523.858586" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 553.786833 614.510623 A 35.954561 14.090909 0 0 0 625.695955 614.510623" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 596.169697 478.400000 A 62.681818 20.000000 0 0 0 721.533333 478.400000" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 556.142096 101.500522 A 74.378788 17.045455 0 0 0 704.899672 101.500522" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 438.554510 397.839827 A 21.282131 18.419913 0 0 0 481.118773 397.839827" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 303.138046 526.222222 A 25.500015 16.000000 0 0 0 354.138077 526.222222" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 579.787879 410.557576 A 49.545454 13.818182 0 0 0 678.878788 410.557576" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 284.666681 103.200000 A 44.833334 15.500000 0 0 0 374.333348 103.200000" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 612.064814 336.159512 A 37.500000 24.318182 0 0 0 687.064814 336.159512" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 296.680251 453.927203 A 27.000000 14.000000 0 0 0 350.680251 453.927203" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 445.888889 71.333333 A 42.000000 16.500000 0 0 0 529.888889 71.333333" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 57.000000 158.000000 A 36.500000 15.000000 0 0 0 130.000000 158.000000" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 662.222222 168.555556 A 34.500000 15.000000 0 0 0 731.222222 168.555556" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<path class="click-through" d="M 156.000000 117.000000 A 35.000000 15.000000 0 0 0 226.000000 117.000000" fill="#DED6CA" fill-opacity="1" stroke="#343434" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
</g>
</svg>`

	// insertion point for setup of pointers
}
