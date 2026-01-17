package models

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"
)

func walk(
	zf *file,
	parentNode any,
	node *Node,
	gongdocxStage *Stage,
	node_ *Node_,
	w io.Writer) error {

	var nodeCounter = 0
	var dummyNode *Node

	switch node_.XMLName.Local {
	case "hyperlink":
		// foo.Fprint(w, "[")
		var cbuf bytes.Buffer
		for _, n := range node_.Nodes {
			if err := walk(zf, dummyNode, node, gongdocxStage, &n, &cbuf); err != nil {
				return err
			}
		}
		// foo.Fprint(w, escape(cbuf.String(), "[]"))
		// foo.Fprint(w, "]")

		// foo.Fprint(w, "(")
		if id, ok := attr(node_.Attrs, "id"); ok {
			for _, rel := range zf.rels.Relationship {
				if id == rel.ID {
					// foo.Fprint(w, escape(rel.Target, "()"))
					break
				}
			}
		}
		// foo.Fprint(w, ")")
	case "t":
		node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter)}).Stage(gongdocxStage)
		node.Nodes = append(node.Nodes, node__)
		nodeCounter = nodeCounter + 1

		text := (&Text{Name: node__.Name}).Stage(gongdocxStage)
		text.Node = node__
		text.Content = string(node_.Content)

		if _, ok := attr(node_.Attrs, "space"); ok {
			text.PreserveWhiteSpace = true
		}

		// check if the parent node is a rune
		switch rune_ := parentNode.(type) {
		case *Rune:
			rune_.Text = text
			text.EnclosingRune = rune_
			if rune_.EnclosingParagraph != nil {
				rune_.EnclosingParagraph.CollatedText = rune_.EnclosingParagraph.CollatedText + text.Content
			}
		}

		// foo.Fprint(w, string(node_.Content))
	case "pPr":
		code := false

		paragraphProperties := (&ParagraphProperties{Name: node.Name}).Stage(gongdocxStage)
		paragraphProperties.Node = node
		paragraphProperties.Content = string(node_.Content)

		// check if the parent node is a paragraph
		switch paragraph := parentNode.(type) {
		case *Paragraph:
			paragraph.ParagraphProperties = paragraphProperties
		}

		for _, n := range node_.Nodes {

			switch n.XMLName.Local {
			case "ind":
				if left, ok := attr(n.Attrs, "left"); ok {
					if i, err := strconv.Atoi(left); err == nil && i > 0 {
						// foo.Fprint(w, strings.Repeat("  ", i/360))
					}
				}
			case "pStyle":

				nodeCounter_ := 0
				node___ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter_)}).Stage(gongdocxStage)
				node.Nodes = append(node.Nodes, node___)
				nodeCounter_ = nodeCounter_ + 1

				paragraphStyle := (&ParagraphStyle{Name: node___.Name}).Stage(gongdocxStage)
				paragraphStyle.Node = node
				paragraphStyle.Content = string(n.Content)
				paragraphProperties.ParagraphStyle = paragraphStyle

				if val, ok := attr(n.Attrs, "val"); ok {

					paragraphStyle.ValAttr = val
					paragraphStyle.Name = paragraphStyle.Name + " " + paragraphStyle.ValAttr

					if strings.HasPrefix(val, "Heading") {
						if i, err := strconv.Atoi(val[7:]); err == nil && i > 0 {
							// foo.Fprint(w, strings.Repeat("#", i)+" ")
						}
					} else if val == "Code" {
						code = true
					} else {
						if i, err := strconv.Atoi(val); err == nil && i > 0 {
							// foo.Fprint(w, strings.Repeat("#", i)+" ")
						}
					}
				}
			case "numPr":
				numID := ""
				ilvl := ""
				numFmt := ""
				start := 1
				ind := 0
				for _, nn := range n.Nodes {
					if nn.XMLName.Local == "numId" {
						if val, ok := attr(nn.Attrs, "val"); ok {
							numID = val
						}
					}
					if nn.XMLName.Local == "ilvl" {
						if val, ok := attr(nn.Attrs, "val"); ok {
							ilvl = val
						}
					}
				}
				for _, num := range zf.num.Num {
					if numID != num.NumID {
						continue
					}
					for _, abnum := range zf.num.AbstractNum {
						if abnum.AbstractNumID != num.AbstractNumID.Val {
							continue
						}
						for _, ablvl := range abnum.Lvl {
							if ablvl.Ilvl != ilvl {
								continue
							}
							if i, err := strconv.Atoi(ablvl.Start.Val); err == nil {
								start = i
							}
							if i, err := strconv.Atoi(ablvl.PPr.Ind.Left); err == nil {
								ind = i / 360
							}
							numFmt = ablvl.NumFmt.Val
							break
						}
						break
					}
					break
				}

				// foo.Fprint(w, strings.Repeat("  ", ind))
				switch numFmt {
				case "decimal", "aiueoFullWidth":
					key := fmt.Sprintf("%s:%d", numID, ind)
					cur, ok := zf.list[key]
					if !ok {
						zf.list[key] = start
					} else {
						zf.list[key] = cur + 1
					}
					// foo.Fprintf(w, "%d. ", zf.list[key])
				case "bullet":
					// foo.Fprint(w, "* ")
				}
			}
		}
		if code {
			// foo.Fprint(w, "`")
		}
		for _, n := range node_.Nodes {
			if err := walk(zf, dummyNode, node, gongdocxStage, &n, w); err != nil {
				return err
			}
		}
		if code {
			// foo.Fprint(w, "`")
		}
	case "tbl":

		table := (&Table{Name: node.Name}).Stage(gongdocxStage)
		table.Node = node
		// table.Content = string(node_.Content)

		var rows [][]string
		nodeRowCounter_ := 0
		for _, tableNode := range node_.Nodes {
			if tableNode.XMLName.Local == "tblPr" {
				nodeCounter_ := 0
				node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter_)}).Stage(gongdocxStage)
				node.Nodes = append(node.Nodes, node__)
				nodeCounter_ = nodeCounter_ + 1

				tableProperties := (&TableProperties{Name: node.Name}).Stage(gongdocxStage)
				tableProperties.Node = node
				tableProperties.Content = string(tableNode.Content)
				table.TableProperties = tableProperties

				for _, tblStyle := range tableNode.Nodes {

					if tblStyle.XMLName.Local == "tblStyle" {
						node___ := (&Node{Name: fmt.Sprintf(node__.Name+".%d", nodeCounter_)}).Stage(gongdocxStage)
						node.Nodes = append(node.Nodes, node___)
						nodeCounter_ = nodeCounter_ + 1

						tableStyle := (&TableStyle{Name: node.Name}).Stage(gongdocxStage)
						tableStyle.Node = node
						tableStyle.Content = string(tblStyle.Content)
						tableProperties.TableStyle = tableStyle

						if val, ok := attr(tblStyle.Attrs, "val"); ok {
							tableStyle.Val = val
						}
					}
				}
			}
			if tableNode.XMLName.Local != "tr" {
				continue
			}
			node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeRowCounter_)}).Stage(gongdocxStage)
			node.Nodes = append(node.Nodes, node__)
			nodeRowCounter_ = nodeRowCounter_ + 1

			tableRow := (&TableRow{Name: node__.Name}).Stage(gongdocxStage)
			tableRow.Node = node
			// tableRow.Content = string(tableNode.Content)
			table.TableRows = append(table.TableRows, tableRow)

			var cols []string
			nodeColumnCounter__ := 0
			for _, tc := range tableNode.Nodes {
				if tc.XMLName.Local != "tc" {
					continue
				}
				node__ := (&Node{Name: fmt.Sprintf(node__.Name+".%d", nodeColumnCounter__)}).Stage(gongdocxStage)
				node.Nodes = append(node__.Nodes, node__)
				nodeColumnCounter__ = nodeColumnCounter__ + 1

				tableColumn := (&TableColumn{Name: node__.Name}).Stage(gongdocxStage)
				tableColumn.Node = node__
				// tableColumn.Content = string(tc.Content)
				tableRow.TableColumns = append(tableRow.TableColumns, tableColumn)

				var cbuf bytes.Buffer
				if err := walk(zf, tableColumn, node__, gongdocxStage, &tc, &cbuf); err != nil {
					return err
				}
				cols = append(cols, strings.Replace(cbuf.String(), "\n", "", -1))
			}
			rows = append(rows, cols)
		}
		maxcol := 0
		for _, cols := range rows {
			if len(cols) > maxcol {
				maxcol = len(cols)
			}
		}
		widths := make([]int, maxcol)
		for _, row := range rows {
			for i := 0; i < maxcol; i++ {
				if i < len(row) {
					width := runewidth.StringWidth(row[i])
					if widths[i] < width {
						widths[i] = width
					}
				}
			}
		}
		for i, row := range rows {
			if i == 0 {
				for j := 0; j < maxcol; j++ {
					// // foo.Fprint(w, "|")
					// // foo.Fprint(w, strings.Repeat(" ", widths[j]))
				}
				// // foo.Fprint(w, "|\n")
				for j := 0; j < maxcol; j++ {
					// // foo.Fprint(w, "|")
					// // foo.Fprint(w, strings.Repeat("-", widths[j]))
				}
				// // foo.Fprint(w, "|\n")
			}
			for j := 0; j < maxcol; j++ {
				// // foo.Fprint(w, "|")
				if j < len(row) {
					// width := runewidth.StringWidth(row[j])
					// // foo.Fprint(w, escape(row[j], "|"))
					// // foo.Fprint(w, strings.Repeat(" ", widths[j]-width))
				} else {
					// // foo.Fprint(w, strings.Repeat(" ", widths[j]))
				}
			}
			// // foo.Fprint(w, "|\n")
		}
		// // foo.Fprint(w, "\n")
	case "r":

		rune_ := (&Rune{Name: node.Name}).Stage(gongdocxStage)
		rune_.Node = node
		rune_.Content = string(node_.Content)

		// check if the parent node is a paragraph
		switch paragraph := parentNode.(type) {
		case *Paragraph:
			paragraph.Runes = append(paragraph.Runes, rune_)
			rune_.EnclosingParagraph = paragraph
		}

		bold := false
		italic := false
		strike := false
		for _, n := range node_.Nodes {
			if n.XMLName.Local != "rPr" {
				continue
			}

			nodeCounter_ := 0
			node___ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter_)}).Stage(gongdocxStage)
			node.Nodes = append(node.Nodes, node___)
			nodeCounter_ = nodeCounter_ + 1

			runeProperties := (&RuneProperties{Name: node___.Name}).Stage(gongdocxStage)
			runeProperties.Node = node___
			runeProperties.Content = string(n.Content)
			rune_.RuneProperties = runeProperties

			for _, nn := range n.Nodes {

				switch nn.XMLName.Local {
				case "b":
					bold = true
					runeProperties.IsBold = true
				case "i":
					italic = true
					runeProperties.IsItalic = true
				case "strike":
					strike = true
					runeProperties.IsStrike = true
				}
			}
		}
		if strike {
			// // foo.Fprint(w, "~~")
		}
		if bold {
			// // foo.Fprint(w, "**")
		}
		if italic {
			// // foo.Fprint(w, "*")
		}
		var cbuf bytes.Buffer
		for _, n := range node_.Nodes {
			if err := walk(zf, rune_, node, gongdocxStage, &n, &cbuf); err != nil {
				return err
			}
		}
		// // foo.Fprint(w, escape(cbuf.String(), `*~\`))
		if italic {
			// // foo.Fprint(w, "*")
		}
		if bold {
			// // foo.Fprint(w, "**")
		}
		if strike {
			// // foo.Fprint(w, "~~")
		}
	case "p":
		paragraph := (&Paragraph{Name: node.Name}).Stage(gongdocxStage)
		paragraph.Node = node
		// paragraph.Content = string(node_.Content)

		// check if the parent node is a paragraph
		switch parentNodeTyped := parentNode.(type) {
		case *TableColumn:
			parentNodeTyped.Paragraphs = append(parentNodeTyped.Paragraphs, paragraph)
			paragraph.EnclosingTableColumn = parentNodeTyped
		case *Body:
			parentNodeTyped.Paragraphs = append(parentNodeTyped.Paragraphs, paragraph)
			paragraph.EnclosingBody = parentNodeTyped

			// make the link
			if parentNodeTyped.LastParagraph != nil {
				parentNodeTyped.LastParagraph.Next = paragraph
				paragraph.Previous = parentNodeTyped.LastParagraph
			}

			parentNodeTyped.LastParagraph = paragraph
		}

		for _, n := range node_.Nodes {

			node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter)}).Stage(gongdocxStage)
			node.Nodes = append(node.Nodes, node__)
			nodeCounter = nodeCounter + 1

			if err := walk(zf, paragraph, node__, gongdocxStage, &n, w); err != nil {
				return err
			}
		}
		// // foo.Fprintln(w)
	case "blip":
		if id, ok := attr(node_.Attrs, "embed"); ok {
			for _, rel := range zf.rels.Relationship {
				if id != rel.ID {
					continue
				}
				if err := zf.extract(&rel, w); err != nil {
					return err
				}
			}
		}
	case "Fallback":
	case "txbxContent":
		var cbuf bytes.Buffer
		for _, n := range node_.Nodes {
			if err := walk(zf, dummyNode, node, gongdocxStage, &n, &cbuf); err != nil {
				return err
			}
		}
		// foo.Fprintln(w, "\n```\n"+cbuf.String()+"```")
	default:
		for _, n := range node_.Nodes {
			node__ := (&Node{Name: fmt.Sprintf(node.Name+".%d", nodeCounter)}).Stage(gongdocxStage)
			node.Nodes = append(node.Nodes, node__)
			nodeCounter = nodeCounter + 1

			if err := walk(zf, parentNode, node__, gongdocxStage, &n, w); err != nil {
				return err
			}
		}
	}

	return nil
}
