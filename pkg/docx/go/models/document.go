package models

// GONGNOTE(NoteOnDocument)
// In the structure of a .docx file, the word/document.xml file is one of the most
// crucial components. It contains the main content of the document, including the
// text and its organization into paragraphs and other structures, as well as
// references to other components of the document such as images, styles, formatting
// instructions, etc.
//
//	-
//
// This XML file primarily houses the textual content and its associated XML tags
// denote various properties of the text such as font size, style, alignment, and
// more. All these pieces of information are represented using WordProcessingML, a
// type of XML developed by Microsoft for Word documents.
// -
// The document.xml file refers to other files in the .docx structure to help render
// the final document. For instance, it references styles from the styles.xml file,
// numbered lists from the numbering.xml file, relationships from the
// document.xml.rels file, and more. These references help assemble the complete,
// rendered document that you see when you open a .docx file in a word processor.
// -
// There could be multiple document*.xml files in some situations. This is typically
// seen when a Word document has been split into separate sections, perhaps for
// editing or collaboration purposes. Each document*.xml file would contain a
// different section of the overall document.
const NoteOnDocument = ""

type Document struct {
	Name string
	File *File
	Root *Node
	Body *Body
}
