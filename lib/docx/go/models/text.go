package models

// GONGNOTE(NodeOnText) :
// [models.Text]
// -
// t: This stands for 'text'. This element is used to represent a run of text within
// a paragraph. The t element will contain the actual string of text as its content.
// -
// When parsing these nodes, your code should handle each type of node appropriately
// based on its name. For example, when you encounter a t node, you might simply
// extract and store the text content.
// -
// In XML, "xml:space="preserve"" instructs the XML processor to preserve
// whitespace in the content of the element it's applied to. This includes spaces,
// line breaks, and tabs.
// -
// By default, XML processors may ignore or normalize such whitespace. This
// attribute ensures that the whitespace is kept intact, which can be important
// for maintaining the correct formatting or interpretation of the data.
// -
// For example, in a WordProcessingML document, if "xml:space="preserve"" is
// applied to a text ("w:t") element, it means the spaces within that text should
// be preserved when the document is displayed or processed.
const NodeOnText = ""

type Text struct {
	Name               string
	Content            string
	Node               *Node
	PreserveWhiteSpace bool

	// navigation
	EnclosingRune *Rune
}
