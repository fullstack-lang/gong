package models

import (
	"fmt"
	"html"
	"strconv"
	"strings"
)

func generateTextElement(sb *strings.Builder, content string, x, y float64,
	explicitFontSize, explicitFontWeight, explicitFontStyle, textAnchor string,
	p *Presentation) {

	lines := strings.Split(content, "\n")
	baseX := formatFloat(x)

	sb.WriteString(fmt.Sprintf(`  <text x="%s" y="%s"`, baseX, formatFloat(y)))

	if textAnchor != "" {
		sb.WriteString(fmt.Sprintf(` text-anchor="%s"`, html.EscapeString(textAnchor)))
	}

	if textAnchor == string(TEXT_ANCHOR_CENTER) { // Use models.TEXT_ANCHOR_MIDDLE
		sb.WriteString(` dominant-baseline="middle"`)
	}

	clonedPres := *p
	if clonedPres.Color == "" {
		clonedPres.Color = "black"
	}

	finalFontSize := explicitFontSize
	if finalFontSize == "" {
		finalFontSize = "12px"
	} else {
		if _, errConv := strconv.Atoi(strings.TrimSuffix(finalFontSize, "px")); errConv == nil && !strings.HasSuffix(finalFontSize, "px") {
			finalFontSize += "px"
		}
	}
	sb.WriteString(fmt.Sprintf(` font-size="%s"`, html.EscapeString(finalFontSize)))

	if explicitFontWeight != "" {
		sb.WriteString(fmt.Sprintf(` font-weight="%s"`, html.EscapeString(explicitFontWeight)))
	}
	if explicitFontStyle != "" {
		sb.WriteString(fmt.Sprintf(` font-style="%s"`, html.EscapeString(explicitFontStyle)))
	}

	appendPresentationAttributes(sb, &clonedPres, true)
	sb.WriteString(">\n")

	for i, line := range lines {
		if i == 0 {
			sb.WriteString(fmt.Sprintf("    <tspan>%s</tspan>\n", html.EscapeString(line)))
		} else {
			sb.WriteString(fmt.Sprintf("    <tspan x=\"%s\" dy=\"1.2em\">%s</tspan>\n", baseX, html.EscapeString(line)))
		}
	}
	sb.WriteString("  </text>\n")
}
