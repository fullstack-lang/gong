// File: xhtml/converter.go

package xhtml

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// ToMarkdown converts a given XHTML string to its Markdown equivalent.
func ToMarkdown(xhtmlString string) (string, error) {
	sanitizedXHTML := strings.ReplaceAll(xhtmlString, "reqif-xhtml:", "") // DOORS exports xhtml with the "reqif-" prefix
	sanitizedXHTML = strings.ReplaceAll(sanitizedXHTML, "xhtml:", "")

	doc, err := html.Parse(strings.NewReader(sanitizedXHTML))
	if err != nil {
		return "", fmt.Errorf("failed to parse xhtml: %w", err)
	}

	bodyNode := findBody(doc)
	if bodyNode == nil {
		return "", fmt.Errorf("could not find <body> tag in parsed document")
	}

	var mdBuilder strings.Builder
	for c := bodyNode.FirstChild; c != nil; c = c.NextSibling {
		// Pass the parent node to renderNode for context (important for lists)
		mdBuilder.WriteString(renderNode(c, c.Parent))
	}

	finalOutput := strings.TrimSpace(mdBuilder.String())
	re := regexp.MustCompile(`\n{3,}`)
	finalOutput = re.ReplaceAllString(finalOutput, "\n\n")

	return finalOutput, nil
}

// renderNode is the main recursive function for processing nodes.
// It now takes a parent node for context.
func renderNode(n *html.Node, parent *html.Node) string {
	switch n.Type {
	case html.TextNode:
		// Discard text nodes that are just formatting whitespace between block elements.
		if parent != nil && (parent.Data == "body" || parent.Data == "blockquote" || parent.Data == "div") {
			if strings.TrimSpace(n.Data) == "" {
				return ""
			}
		}
		return n.Data

	case html.ElementNode:
		var contentBuilder strings.Builder
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			contentBuilder.WriteString(renderNode(c, n))
		}
		content := contentBuilder.String()

		switch n.Data {
		// BLOCK-LEVEL ELEMENTS
		case "p":
			return "\n\n" + normalizeWhitespace(content)
		case "h1":
			return "\n\n# " + normalizeWhitespace(content)
		case "h2":
			return "\n\n## " + normalizeWhitespace(content)
		case "h3":
			return "\n\n### " + normalizeWhitespace(content)
		case "h4":
			return "\n\n#### " + normalizeWhitespace(content)
		case "blockquote":
			lines := strings.Split(normalizeWhitespace(content), "\n")
			for i, line := range lines {
				lines[i] = "> " + line
			}
			return "\n\n" + strings.Join(lines, "\n")
		case "ul":
			var result strings.Builder
			result.WriteString("\n")
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "li" {
					result.WriteString("- " + renderNode(c, n) + "\n")
				}
			}
			return result.String()
		case "ol":
			var result strings.Builder
			result.WriteString("\n")
			itemCounter := 1
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "li" {
					result.WriteString(fmt.Sprintf("%d. %s\n", itemCounter, renderNode(c, n)))
					itemCounter++
				}
			}
			return result.String()
		case "li":
			return normalizeWhitespace(content)

		// TABLE-SPECIFIC LOGIC
		case "table":
			return "\n\n" + renderTable(n) + "\n\n"

		// INLINE-LEVEL ELEMENTS
		case "strong", "b":
			return "**" + normalizeWhitespace(content) + "**"
		case "em", "i":
			return "*" + normalizeWhitespace(content) + "*"
		case "a":
			href := getAttr(n, "href")
			text := normalizeWhitespace(content)
			return fmt.Sprintf("[%s](%s)", text, href)
		case "br":
			return "  \n"
		case "object":
			if getAttr(n, "type") == "image/svg" {
				data := getAttr(n, "data")
				return fmt.Sprintf("\n\n![%s](svg:%s?width=750px)", data, data)
			}
			if getAttr(n, "type") == "image/jpg" {
				data := getAttr(n, "data")
				return fmt.Sprintf("\n\n![%s](jpg:%s?width=750px)", data, data)
			}
			if getAttr(n, "type") == "image/png" {
				data := getAttr(n, "data")
				return fmt.Sprintf("\n\n![%s](png:%s?width=750px)", data, data)
			}
			return content
		default:
			return content
		}
	}
	return ""
}

// findFirstRowNode finds the first <tr> element in a table,
// searching within thead, tbody, or direct children.
func findFirstRowNode(table *html.Node) *html.Node {
	for child := table.FirstChild; child != nil; child = child.NextSibling {
		if child.Type != html.ElementNode {
			continue
		}
		// Case 1: <tr> is a direct child of <table>
		if child.Data == "tr" {
			return child
		}
		// Case 2: <tr> is inside <thead> or <tbody>
		if child.Data == "thead" || child.Data == "tbody" {
			for tr := child.FirstChild; tr != nil; tr = tr.NextSibling {
				if tr.Type == html.ElementNode && tr.Data == "tr" {
					return tr
				}
			}
		}
	}
	return nil
}

// renderTable specifically handles the logic for converting a <table> node.
func renderTable(tableNode *html.Node) string {
	var b strings.Builder
	var headerRows [][]string
	var bodyRows [][]string

	// Find thead and tbody and parse their rows
	for child := tableNode.FirstChild; child != nil; child = child.NextSibling {
		if child.Type != html.ElementNode {
			continue
		}
		switch child.Data {
		case "thead":
			headerRows = parseTableRows(child, "th")
		case "tbody":
			bodyRows = parseTableRows(child, "td")
		}
	}

	// Fallback for tables without explicit thead/tbody
	if len(headerRows) == 0 && len(bodyRows) == 0 {
		bodyRows = parseTableRows(tableNode, "td")
		// Check if the first row looks like a header (contains only th)
		if len(bodyRows) > 0 {
			firstRowIsHeader := true
			firstTr := tableNode.FirstChild
			for firstTr != nil && (firstTr.Type != html.ElementNode || firstTr.Data != "tr") {
				firstTr = firstTr.NextSibling
			}

			if firstTr != nil {
				for cell := firstTr.FirstChild; cell != nil; cell = cell.NextSibling {
					if cell.Type == html.ElementNode && cell.Data != "th" {
						firstRowIsHeader = false
						break
					}
				}
				if firstRowIsHeader {
					headerRows = [][]string{bodyRows[0]}
					bodyRows = bodyRows[1:]
				}
			}
		}
	}

	// --- FIX: Handle header row inside a <tbody> ---
	if len(headerRows) == 0 && len(bodyRows) > 0 {
		firstTr := findFirstRowNode(tableNode)
		if firstTr != nil {
			isHeader := true
			cellCount := 0
			for cell := firstTr.FirstChild; cell != nil; cell = cell.NextSibling {
				if cell.Type == html.ElementNode {
					if cell.Data == "td" || cell.Data == "th" {
						cellCount++
						if cell.Data != "th" {
							isHeader = false
							break
						}
					}
				}
			}
			// If the first row has cells and all are <th>, treat it as a header
			if isHeader && cellCount > 0 {
				headerRows = append(headerRows, bodyRows[0])
				bodyRows = bodyRows[1:]
			}
		}
	}

	// --- RENDER LOGIC ---

	// 1. Determine the maximum number of columns in the entire table
	maxColumns := 0
	allRows := append(headerRows, bodyRows...)
	for _, row := range allRows {
		if len(row) > maxColumns {
			maxColumns = len(row)
		}
	}

	// If table is empty, return nothing
	if maxColumns == 0 {
		return ""
	}

	// 2. Render header, padded to maxColumns
	if len(headerRows) > 0 {
		header := headerRows[0]
		// Pad the header row with empty strings if it's shorter than maxColumns
		for len(header) < maxColumns {
			header = append(header, "")
		}
		b.WriteString("| " + strings.Join(header, " | ") + " |\n")
	} else {
		// If there's no header, we can't make a valid markdown table.
		// For this case, we'll add an empty header row to satisfy markdown syntax.
		var emptyHeader []string
		for i := 0; i < maxColumns; i++ {
			emptyHeader = append(emptyHeader, "")
		}
		b.WriteString("| " + strings.Join(emptyHeader, " | ") + " |\n")
	}

	// 3. Render the separator, which MUST match maxColumns
	var separator []string
	for i := 0; i < maxColumns; i++ {
		separator = append(separator, "---")
	}
	b.WriteString("| " + strings.Join(separator, " | ") + " |\n")

	// 4. Render body rows, each padded to maxColumns
	for _, row := range bodyRows {
		// Pad the current row with empty strings to match maxColumns
		for len(row) < maxColumns {
			row = append(row, "")
		}
		b.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}

	return b.String()
}

// parseTableRows is a helper to extract cell content from a <tr>.
func parseTableRows(node *html.Node, cellTag string) [][]string {
	var rows [][]string
	for tr := node.FirstChild; tr != nil; tr = tr.NextSibling {
		if tr.Type == html.ElementNode && tr.Data == "tr" {
			var row []string
			for cell := tr.FirstChild; cell != nil; cell = cell.NextSibling {
				if cell.Type == html.ElementNode && (cell.Data == "td" || cell.Data == "th") {
					row = append(row, normalizeWhitespace(renderNode(cell, tr)))
				}
			}
			rows = append(rows, row)
		}
	}
	return rows
}

// normalizeWhitespace collapses all forms of whitespace into a single space.
func normalizeWhitespace(s string) string {
	re := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(re.ReplaceAllString(s, " "))
}

// --- Helper functions ---

func getAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func findBody(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "body" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findBody(c); result != nil {
			return result
		}
	}
	return nil
}
