package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/pkg/docx/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__AttributeShape__00000000_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000003_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000004_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000005_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000006_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000007_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000008_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000010_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000011_ := (&models.AttributeShape{Name: `IsBold`}).Stage(stage)
	__AttributeShape__00000012_ := (&models.AttributeShape{Name: `IsStrike`}).Stage(stage)
	__AttributeShape__00000013_ := (&models.AttributeShape{Name: `IsItalic`}).Stage(stage)
	__AttributeShape__00000014_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000015_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000016_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000019_ := (&models.AttributeShape{Name: `Text`}).Stage(stage)
	__AttributeShape__00000020_ := (&models.AttributeShape{Name: `PreserveWhiteSpace`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `01 - Document`}).Stage(stage)
	__Classdiagram__00000001_ := (&models.Classdiagram{Name: `04 - Table`}).Stage(stage)
	__Classdiagram__00000002_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)
	__Classdiagram__00000003_ := (&models.Classdiagram{Name: `02 - Node`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-12-19T11:01:18Z`}).Stage(stage)

	__GongNoteShape__00000000_ := (&models.GongNoteShape{Name: `Default-NoteOnDocument`}).Stage(stage)
	__GongNoteShape__00000001_ := (&models.GongNoteShape{Name: `02 - Body-NoteOnTable`}).Stage(stage)
	__GongNoteShape__00000002_ := (&models.GongNoteShape{Name: `02 - Body-NoteOnColumn`}).Stage(stage)
	__GongNoteShape__00000003_ := (&models.GongNoteShape{Name: `Default-NoteOnRune`}).Stage(stage)
	__GongNoteShape__00000004_ := (&models.GongNoteShape{Name: `Default-NoteOnRunProperties`}).Stage(stage)
	__GongNoteShape__00000006_ := (&models.GongNoteShape{Name: `Default-NoteOnParagraph`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-Docx`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Document`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-File`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Default-Node`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Default-Body`}).Stage(stage)
	__GongStructShape__00000005_ := (&models.GongStructShape{Name: `02 - Body-Body`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `02 - Body-Paragraph`}).Stage(stage)
	__GongStructShape__00000007_ := (&models.GongStructShape{Name: `02 - Body-Table`}).Stage(stage)
	__GongStructShape__00000008_ := (&models.GongStructShape{Name: `02 - Body-TableRow`}).Stage(stage)
	__GongStructShape__00000009_ := (&models.GongStructShape{Name: `02 - Body-TableProperties`}).Stage(stage)
	__GongStructShape__00000010_ := (&models.GongStructShape{Name: `02 - Body-TableStyle`}).Stage(stage)
	__GongStructShape__00000011_ := (&models.GongStructShape{Name: `02 - Body-TableColumn`}).Stage(stage)
	__GongStructShape__00000013_ := (&models.GongStructShape{Name: `Default-Document`}).Stage(stage)
	__GongStructShape__00000014_ := (&models.GongStructShape{Name: `Default-Paragraph`}).Stage(stage)
	__GongStructShape__00000015_ := (&models.GongStructShape{Name: `Default-Body`}).Stage(stage)
	__GongStructShape__00000016_ := (&models.GongStructShape{Name: `02 - Node-Node`}).Stage(stage)
	__GongStructShape__00000017_ := (&models.GongStructShape{Name: `02 - Node-Paragraph`}).Stage(stage)
	__GongStructShape__00000018_ := (&models.GongStructShape{Name: `02 - Node-Document`}).Stage(stage)
	__GongStructShape__00000019_ := (&models.GongStructShape{Name: `02 - Node-ParagraphProperties`}).Stage(stage)
	__GongStructShape__00000020_ := (&models.GongStructShape{Name: `02 - Node-ParagraphStyle`}).Stage(stage)
	__GongStructShape__00000021_ := (&models.GongStructShape{Name: `02 - Node-Rune`}).Stage(stage)
	__GongStructShape__00000022_ := (&models.GongStructShape{Name: `02 - Node-RuneProperties`}).Stage(stage)
	__GongStructShape__00000023_ := (&models.GongStructShape{Name: `02 - Node-Table`}).Stage(stage)
	__GongStructShape__00000024_ := (&models.GongStructShape{Name: `02 - Node-TableColumn`}).Stage(stage)
	__GongStructShape__00000025_ := (&models.GongStructShape{Name: `02 - Node-TableProperties`}).Stage(stage)
	__GongStructShape__00000026_ := (&models.GongStructShape{Name: `02 - Node-TableRow`}).Stage(stage)
	__GongStructShape__00000027_ := (&models.GongStructShape{Name: `02 - Node-TableStyle`}).Stage(stage)
	__GongStructShape__00000028_ := (&models.GongStructShape{Name: `02 - Node-Text`}).Stage(stage)
	__GongStructShape__00000029_ := (&models.GongStructShape{Name: `Default-ParagraphProperties`}).Stage(stage)
	__GongStructShape__00000030_ := (&models.GongStructShape{Name: `Default-ParagraphStyle`}).Stage(stage)
	__GongStructShape__00000031_ := (&models.GongStructShape{Name: `Default-Rune`}).Stage(stage)
	__GongStructShape__00000032_ := (&models.GongStructShape{Name: `Default-RuneProperties`}).Stage(stage)
	__GongStructShape__00000033_ := (&models.GongStructShape{Name: `Default-Text`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `Document`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `Files`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `Root`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `File`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `Body`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `Paragraphs`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `TableRows`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `TableProperties`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `TableStyle`}).Stage(stage)
	__LinkShape__00000009_ := (&models.LinkShape{Name: `LastParagraph`}).Stage(stage)
	__LinkShape__00000010_ := (&models.LinkShape{Name: `Tables`}).Stage(stage)
	__LinkShape__00000011_ := (&models.LinkShape{Name: `Paragraphs`}).Stage(stage)
	__LinkShape__00000012_ := (&models.LinkShape{Name: `TableColumns`}).Stage(stage)
	__LinkShape__00000015_ := (&models.LinkShape{Name: `Body`}).Stage(stage)
	__LinkShape__00000016_ := (&models.LinkShape{Name: `Nodes`}).Stage(stage)
	__LinkShape__00000017_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000018_ := (&models.LinkShape{Name: `Root`}).Stage(stage)
	__LinkShape__00000019_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000020_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000021_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000022_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000023_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000024_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000025_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000026_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000027_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000028_ := (&models.LinkShape{Name: `Node`}).Stage(stage)
	__LinkShape__00000029_ := (&models.LinkShape{Name: `Paragraphs`}).Stage(stage)
	__LinkShape__00000030_ := (&models.LinkShape{Name: `LastParagraph`}).Stage(stage)
	__LinkShape__00000031_ := (&models.LinkShape{Name: `ParagraphStyle`}).Stage(stage)
	__LinkShape__00000032_ := (&models.LinkShape{Name: `ParagraphProperties`}).Stage(stage)
	__LinkShape__00000033_ := (&models.LinkShape{Name: `Runes`}).Stage(stage)
	__LinkShape__00000034_ := (&models.LinkShape{Name: `RuneProperties`}).Stage(stage)
	__LinkShape__00000035_ := (&models.LinkShape{Name: `EnclosingParagraph`}).Stage(stage)
	__LinkShape__00000037_ := (&models.LinkShape{Name: `Next`}).Stage(stage)
	__LinkShape__00000038_ := (&models.LinkShape{Name: `Previous`}).Stage(stage)
	__LinkShape__00000039_ := (&models.LinkShape{Name: `Text`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `Name`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.Docx{}.Name
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `Docx`
	__AttributeShape__00000000_.Fieldtypename = `string`

	__AttributeShape__00000001_.Name = `Name`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.File{}.Name
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `File`
	__AttributeShape__00000001_.Fieldtypename = `string`

	__AttributeShape__00000002_.Name = `Name`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.Body{}.Name
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `Body`
	__AttributeShape__00000002_.Fieldtypename = `string`

	__AttributeShape__00000003_.Name = `Name`
	__AttributeShape__00000003_.IdentifierMeta = ref_models.Table{}.Name
	__AttributeShape__00000003_.FieldTypeAsString = ``
	__AttributeShape__00000003_.Structname = `Table`
	__AttributeShape__00000003_.Fieldtypename = `string`

	__AttributeShape__00000004_.Name = `Content`
	__AttributeShape__00000004_.IdentifierMeta = ref_models.Table{}.Content
	__AttributeShape__00000004_.FieldTypeAsString = ``
	__AttributeShape__00000004_.Structname = `Table`
	__AttributeShape__00000004_.Fieldtypename = `string`

	__AttributeShape__00000005_.Name = `Content`
	__AttributeShape__00000005_.IdentifierMeta = ref_models.TableRow{}.Content
	__AttributeShape__00000005_.FieldTypeAsString = ``
	__AttributeShape__00000005_.Structname = `TableRow`
	__AttributeShape__00000005_.Fieldtypename = `string`

	__AttributeShape__00000006_.Name = `Name`
	__AttributeShape__00000006_.IdentifierMeta = ref_models.TableProperties{}.Name
	__AttributeShape__00000006_.FieldTypeAsString = ``
	__AttributeShape__00000006_.Structname = `TableProperties`
	__AttributeShape__00000006_.Fieldtypename = `string`

	__AttributeShape__00000007_.Name = `Content`
	__AttributeShape__00000007_.IdentifierMeta = ref_models.TableProperties{}.Content
	__AttributeShape__00000007_.FieldTypeAsString = ``
	__AttributeShape__00000007_.Structname = `TableProperties`
	__AttributeShape__00000007_.Fieldtypename = `string`

	__AttributeShape__00000008_.Name = `Content`
	__AttributeShape__00000008_.IdentifierMeta = ref_models.TableColumn{}.Content
	__AttributeShape__00000008_.FieldTypeAsString = ``
	__AttributeShape__00000008_.Structname = `TableColumn`
	__AttributeShape__00000008_.Fieldtypename = `string`

	__AttributeShape__00000010_.Name = `Content`
	__AttributeShape__00000010_.IdentifierMeta = ref_models.Paragraph{}.Content
	__AttributeShape__00000010_.FieldTypeAsString = ``
	__AttributeShape__00000010_.Structname = `Paragraph`
	__AttributeShape__00000010_.Fieldtypename = `string`

	__AttributeShape__00000011_.Name = `IsBold`
	__AttributeShape__00000011_.IdentifierMeta = ref_models.RuneProperties{}.IsBold
	__AttributeShape__00000011_.FieldTypeAsString = ``
	__AttributeShape__00000011_.Structname = `RuneProperties`
	__AttributeShape__00000011_.Fieldtypename = `bool`

	__AttributeShape__00000012_.Name = `IsStrike`
	__AttributeShape__00000012_.IdentifierMeta = ref_models.RuneProperties{}.IsStrike
	__AttributeShape__00000012_.FieldTypeAsString = ``
	__AttributeShape__00000012_.Structname = `RuneProperties`
	__AttributeShape__00000012_.Fieldtypename = `bool`

	__AttributeShape__00000013_.Name = `IsItalic`
	__AttributeShape__00000013_.IdentifierMeta = ref_models.RuneProperties{}.IsItalic
	__AttributeShape__00000013_.FieldTypeAsString = ``
	__AttributeShape__00000013_.Structname = `RuneProperties`
	__AttributeShape__00000013_.Fieldtypename = `bool`

	__AttributeShape__00000014_.Name = `Content`
	__AttributeShape__00000014_.IdentifierMeta = ref_models.RuneProperties{}.Content
	__AttributeShape__00000014_.FieldTypeAsString = ``
	__AttributeShape__00000014_.Structname = `RuneProperties`
	__AttributeShape__00000014_.Fieldtypename = `string`

	__AttributeShape__00000015_.Name = `Content`
	__AttributeShape__00000015_.IdentifierMeta = ref_models.Rune{}.Content
	__AttributeShape__00000015_.FieldTypeAsString = ``
	__AttributeShape__00000015_.Structname = `Rune`
	__AttributeShape__00000015_.Fieldtypename = `string`

	__AttributeShape__00000016_.Name = `Name`
	__AttributeShape__00000016_.IdentifierMeta = ref_models.Text{}.Name
	__AttributeShape__00000016_.FieldTypeAsString = ``
	__AttributeShape__00000016_.Structname = `Text`
	__AttributeShape__00000016_.Fieldtypename = `string`

	__AttributeShape__00000019_.Name = `Text`
	__AttributeShape__00000019_.IdentifierMeta = ref_models.Paragraph{}.CollatedText
	__AttributeShape__00000019_.FieldTypeAsString = ``
	__AttributeShape__00000019_.Structname = `Paragraph`
	__AttributeShape__00000019_.Fieldtypename = `string`

	__AttributeShape__00000020_.Name = `PreserveWhiteSpace`
	__AttributeShape__00000020_.IdentifierMeta = ref_models.Text{}.PreserveWhiteSpace
	__AttributeShape__00000020_.FieldTypeAsString = ``
	__AttributeShape__00000020_.Structname = `Text`
	__AttributeShape__00000020_.Fieldtypename = `bool`

	__Classdiagram__00000000_.Name = `01 - Document`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = false
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = false
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,false,false,false]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = true
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__Classdiagram__00000001_.Name = `04 - Table`
	__Classdiagram__00000001_.Description = ``
	__Classdiagram__00000001_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000001_.ShowNbInstances = true
	__Classdiagram__00000001_.ShowMultiplicity = true
	__Classdiagram__00000001_.ShowLinkNames = true
	__Classdiagram__00000001_.IsInRenameMode = false
	__Classdiagram__00000001_.IsExpanded = false
	__Classdiagram__00000001_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000001_.NodeGongStructNodeExpansion = `[true,false,false,false,false,false,false,false,false,false,true,true,true,true]`
	__Classdiagram__00000001_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000001_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000001_.NodeGongNotesIsExpanded = true
	__Classdiagram__00000001_.NodeGongNoteNodeExpansion = ``

	__Classdiagram__00000002_.Name = `Default`
	__Classdiagram__00000002_.Description = ``
	__Classdiagram__00000002_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000002_.ShowNbInstances = true
	__Classdiagram__00000002_.ShowMultiplicity = true
	__Classdiagram__00000002_.ShowLinkNames = true
	__Classdiagram__00000002_.IsInRenameMode = false
	__Classdiagram__00000002_.IsExpanded = true
	__Classdiagram__00000002_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000002_.NodeGongStructNodeExpansion = `[false,false,false,false,false,true,true,false,false,false,false,false,false,false,false,true]`
	__Classdiagram__00000002_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000002_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000002_.NodeGongNotesIsExpanded = true
	__Classdiagram__00000002_.NodeGongNoteNodeExpansion = `[true]`

	__Classdiagram__00000003_.Name = `02 - Node`
	__Classdiagram__00000003_.Description = ``
	__Classdiagram__00000003_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000003_.ShowNbInstances = true
	__Classdiagram__00000003_.ShowMultiplicity = true
	__Classdiagram__00000003_.ShowLinkNames = true
	__Classdiagram__00000003_.IsInRenameMode = false
	__Classdiagram__00000003_.IsExpanded = false
	__Classdiagram__00000003_.NodeGongStructsIsExpanded = false
	__Classdiagram__00000003_.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false]`
	__Classdiagram__00000003_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000003_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000003_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000003_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2025-12-19T11:01:18Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongNoteShape__00000000_.Name = `Default-NoteOnDocument`
	__GongNoteShape__00000000_.Identifier = `ref_models.NoteOnDocument`
	__GongNoteShape__00000000_.Body = `
In the structure of a .docx file, the word/document.xml file is one of the most
crucial components. It contains the main content of the document, including the
text and its organization into paragraphs and other structures, as well as
references to other components of the document such as images, styles, formatting
instructions, etc.

 -

This XML file primarily houses the textual content and its associated XML tags
denote various properties of the text such as font size, style, alignment, and
more. All these pieces of information are represented using WordProcessingML, a
type of XML developed by Microsoft for Word documents.
-
The document.xml file refers to other files in the .docx structure to help render
the final document. For instance, it references styles from the styles.xml file,
numbered lists from the numbering.xml file, relationships from the
document.xml.rels file, and more. These references help assemble the complete,
rendered document that you see when you open a .docx file in a word processor.
-
There could be multiple document*.xml files in some situations. This is typically
seen when a Word document has been split into separate sections, perhaps for
editing or collaboration purposes. Each document*.xml file would contain a
different section of the overall document.
`
	__GongNoteShape__00000000_.BodyHTML = `<p>In the structure of a .docx file, the word/document.xml file is one of the most
crucial components. It contains the main content of the document, including the
text and its organization into paragraphs and other structures, as well as
references to other components of the document such as images, styles, formatting
instructions, etc.
<pre>-
</pre>
<p>This XML file primarily houses the textual content and its associated XML tags
denote various properties of the text such as font size, style, alignment, and
more. All these pieces of information are represented using WordProcessingML, a
type of XML developed by Microsoft for Word documents.
-
The document.xml file refers to other files in the .docx structure to help render
the final document. For instance, it references styles from the styles.xml file,
numbered lists from the numbering.xml file, relationships from the
document.xml.rels file, and more. These references help assemble the complete,
rendered document that you see when you open a .docx file in a word processor.
-
There could be multiple document*.xml files in some situations. This is typically
seen when a Word document has been split into separate sections, perhaps for
editing or collaboration purposes. Each document*.xml file would contain a
different section of the overall document.
`
	__GongNoteShape__00000000_.X = 1186.000000
	__GongNoteShape__00000000_.Y = 59.000000
	__GongNoteShape__00000000_.Width = 691.000000
	__GongNoteShape__00000000_.Height = 307.000000
	__GongNoteShape__00000000_.Matched = false
	__GongNoteShape__00000000_.IsExpanded = false

	__GongNoteShape__00000001_.Name = `02 - Body-NoteOnTable`
	__GongNoteShape__00000001_.Identifier = `ref_models.NoteOnTable`
	__GongNoteShape__00000001_.Body = `
The "w:tbl" node represents a table within a Word document's XML structure. It
is found in the document.xml file.
-
This node defines the structure and formatting of a table in the document. It
contains child elements that represent the table's properties ("w:tblPr"),
grid ("w:tblGrid"), and rows ("w:tr").
-
The "w:tblPr" node defines the table's overall properties, such as its width,
alignment, borders, and shading.
-
The "w:tblGrid" node defines the table's grid structure - specifically, the
number and width of the columns.
-
The "w:tr" nodes represent table rows, and each row contains "w:tc" nodes that
represent the individual cells within that row.
-
When parsing a "w:tbl" node, your code should handle the table structure and
formatting information it provides to correctly represent the table in your
data structure or output format.
`
	__GongNoteShape__00000001_.BodyHTML = `<p>The &quot;w:tbl&quot; node represents a table within a Word document&apos;s XML structure. It
is found in the document.xml file.
-
This node defines the structure and formatting of a table in the document. It
contains child elements that represent the table&apos;s properties (&quot;w:tblPr&quot;),
grid (&quot;w:tblGrid&quot;), and rows (&quot;w:tr&quot;).
-
The &quot;w:tblPr&quot; node defines the table&apos;s overall properties, such as its width,
alignment, borders, and shading.
-
The &quot;w:tblGrid&quot; node defines the table&apos;s grid structure - specifically, the
number and width of the columns.
-
The &quot;w:tr&quot; nodes represent table rows, and each row contains &quot;w:tc&quot; nodes that
represent the individual cells within that row.
-
When parsing a &quot;w:tbl&quot; node, your code should handle the table structure and
formatting information it provides to correctly represent the table in your
data structure or output format.
`
	__GongNoteShape__00000001_.X = 586.000000
	__GongNoteShape__00000001_.Y = 362.000000
	__GongNoteShape__00000001_.Width = 594.000000
	__GongNoteShape__00000001_.Height = 249.000000
	__GongNoteShape__00000001_.Matched = false
	__GongNoteShape__00000001_.IsExpanded = false

	__GongNoteShape__00000002_.Name = `02 - Body-NoteOnColumn`
	__GongNoteShape__00000002_.Identifier = `ref_models.NoteOnColumn`
	__GongNoteShape__00000002_.Body = `
The "w:tr" node represents a table row within a Word document's XML structure.
It is found as a child of the "w:tbl" (table) node in the document.xml file.
-
This node contains child elements that represent the individual cells ("w:tc")
in the row, as well as the row properties ("w:trPr"), which include attributes
like height, header status, and more.
-
The "w:tc" node, on the other hand, represents an individual table cell. It is
found as a child of the "w:tr" (table row) node.
-
This node contains the content of the cell, which can include text,
paragraphs, or even other tables. It also contains the cell properties
("w:tcPr"), which include attributes like cell width, vertical alignment,
borders, shading, and more.
-
When parsing "w:tr" and "w:tc" nodes, your code should correctly map the
structure of the table, row, and cell, and apply the appropriate properties
to each element.
`
	__GongNoteShape__00000002_.BodyHTML = `<p>The &quot;w:tr&quot; node represents a table row within a Word document&apos;s XML structure.
It is found as a child of the &quot;w:tbl&quot; (table) node in the document.xml file.
-
This node contains child elements that represent the individual cells (&quot;w:tc&quot;)
in the row, as well as the row properties (&quot;w:trPr&quot;), which include attributes
like height, header status, and more.
-
The &quot;w:tc&quot; node, on the other hand, represents an individual table cell. It is
found as a child of the &quot;w:tr&quot; (table row) node.
-
This node contains the content of the cell, which can include text,
paragraphs, or even other tables. It also contains the cell properties
(&quot;w:tcPr&quot;), which include attributes like cell width, vertical alignment,
borders, shading, and more.
-
When parsing &quot;w:tr&quot; and &quot;w:tc&quot; nodes, your code should correctly map the
structure of the table, row, and cell, and apply the appropriate properties
to each element.
`
	__GongNoteShape__00000002_.X = 1235.000000
	__GongNoteShape__00000002_.Y = 74.000000
	__GongNoteShape__00000002_.Width = 492.000000
	__GongNoteShape__00000002_.Height = 318.000000
	__GongNoteShape__00000002_.Matched = false
	__GongNoteShape__00000002_.IsExpanded = false

	__GongNoteShape__00000003_.Name = `Default-NoteOnRune`
	__GongNoteShape__00000003_.Identifier = `ref_models.NoteOnRune`
	__GongNoteShape__00000003_.Body = `
for [models.Rune]
-
The "w:r" node, known as a run, represents a continuous run of text within a
paragraph ("w:p" node) in a Word document's XML structure. It is found within
the document.xml file.
-
Runs are segments of text within a paragraph that share the same formatting.
This can include properties like bolding, italics, underlining, color, size,
font, and more. The specific formatting is defined in a "w:rPr" (Run Properties)
element within the "w:r" node.
-
A "w:r" node contains one or more "w:t" nodes, which hold the actual text
content of the run. It can also contain other types of nodes like "w:br" for a
line break or "w:tab" for a tab character.
-
When parsing a "w:r" node, your code should handle the formatting information
provided in the "w:rPr" node (if present) and apply it to the text found within
the "w:t" nodes.
`
	__GongNoteShape__00000003_.BodyHTML = `<p>for <a href="/models#Rune">models.Rune</a>
-
The &quot;w:r&quot; node, known as a run, represents a continuous run of text within a
paragraph (&quot;w:p&quot; node) in a Word document&apos;s XML structure. It is found within
the document.xml file.
-
Runs are segments of text within a paragraph that share the same formatting.
This can include properties like bolding, italics, underlining, color, size,
font, and more. The specific formatting is defined in a &quot;w:rPr&quot; (Run Properties)
element within the &quot;w:r&quot; node.
-
A &quot;w:r&quot; node contains one or more &quot;w:t&quot; nodes, which hold the actual text
content of the run. It can also contain other types of nodes like &quot;w:br&quot; for a
line break or &quot;w:tab&quot; for a tab character.
-
When parsing a &quot;w:r&quot; node, your code should handle the formatting information
provided in the &quot;w:rPr&quot; node (if present) and apply it to the text found within
the &quot;w:t&quot; nodes.
`
	__GongNoteShape__00000003_.X = 584.000000
	__GongNoteShape__00000003_.Y = 767.000000
	__GongNoteShape__00000003_.Width = 786.000000
	__GongNoteShape__00000003_.Height = 189.000000
	__GongNoteShape__00000003_.Matched = false
	__GongNoteShape__00000003_.IsExpanded = false

	__GongNoteShape__00000004_.Name = `Default-NoteOnRunProperties`
	__GongNoteShape__00000004_.Identifier = `ref_models.NoteOnRunProperties`
	__GongNoteShape__00000004_.Body = `
The "w:rPr" node represents run properties in a Word document's XML structure.
It is found within a "w:r" (run) node in the document.xml file.
-
This node defines the formatting for a specific run of text within a paragraph.
It can include properties like font size, font type, color, highlighting,
bolding, italics, underlining, and more.
-
For example, a "w:rPr" node might contain a "w:sz" element for font size, a
"w:color" element for text color, or "w:b" for bold formatting. The presence of
elements like "w:b" (bold), "w:i" (italic), and "w:u" (underline) indicate that
the formatting is applied, as they are toggled by their presence alone.
-
When parsing a "w:rPr" node, your code should use the information it provides to
apply the appropriate formatting to the text in the run ("w:r") that contains
this "w:rPr" node.
`
	__GongNoteShape__00000004_.BodyHTML = `<p>The &quot;w:rPr&quot; node represents run properties in a Word document&apos;s XML structure.
It is found within a &quot;w:r&quot; (run) node in the document.xml file.
-
This node defines the formatting for a specific run of text within a paragraph.
It can include properties like font size, font type, color, highlighting,
bolding, italics, underlining, and more.
-
For example, a &quot;w:rPr&quot; node might contain a &quot;w:sz&quot; element for font size, a
&quot;w:color&quot; element for text color, or &quot;w:b&quot; for bold formatting. The presence of
elements like &quot;w:b&quot; (bold), &quot;w:i&quot; (italic), and &quot;w:u&quot; (underline) indicate that
the formatting is applied, as they are toggled by their presence alone.
-
When parsing a &quot;w:rPr&quot; node, your code should use the information it provides to
apply the appropriate formatting to the text in the run (&quot;w:r&quot;) that contains
this &quot;w:rPr&quot; node.
`
	__GongNoteShape__00000004_.X = 1517.000000
	__GongNoteShape__00000004_.Y = 117.000000
	__GongNoteShape__00000004_.Width = 395.000000
	__GongNoteShape__00000004_.Height = 347.000000
	__GongNoteShape__00000004_.Matched = false
	__GongNoteShape__00000004_.IsExpanded = false

	__GongNoteShape__00000006_.Name = `Default-NoteOnParagraph`
	__GongNoteShape__00000006_.Identifier = `ref_models.NoteOnParagraph`
	__GongNoteShape__00000006_.Body = `
The "w:p" node represents a [models.Paragraph] in a Word document's XML structure.
-
It is one of the primary building blocks of a document's content within the
document.xml file.
-
Each "w:p" element contains a series of "w:r" (run) elements, which represent
sections of text within the paragraph that have consistent formatting. These
"w:r" elements, in turn, contain "w:t" elements that hold the actual text
content.
-
The "w:p" element may also contain various other child elements that provide
additional information about the paragraph. For example, "w:pPr" specifies
paragraph properties like alignment, indentation, and spacing.
-
As you parse a "w:p" node, you would typically create a new paragraph object in
your code, then parse the child nodes to fill in the content and properties of
that paragraph.
`
	__GongNoteShape__00000006_.BodyHTML = `<p>The &quot;w:p&quot; node represents a <a href="/models#Paragraph">models.Paragraph</a> in a Word document&apos;s XML structure.
-
It is one of the primary building blocks of a document&apos;s content within the
document.xml file.
-
Each &quot;w:p&quot; element contains a series of &quot;w:r&quot; (run) elements, which represent
sections of text within the paragraph that have consistent formatting. These
&quot;w:r&quot; elements, in turn, contain &quot;w:t&quot; elements that hold the actual text
content.
-
The &quot;w:p&quot; element may also contain various other child elements that provide
additional information about the paragraph. For example, &quot;w:pPr&quot; specifies
paragraph properties like alignment, indentation, and spacing.
-
As you parse a &quot;w:p&quot; node, you would typically create a new paragraph object in
your code, then parse the child nodes to fill in the content and properties of
that paragraph.
`
	__GongNoteShape__00000006_.X = 61.000000
	__GongNoteShape__00000006_.Y = 484.000000
	__GongNoteShape__00000006_.Width = 509.000000
	__GongNoteShape__00000006_.Height = 336.000000
	__GongNoteShape__00000006_.Matched = false
	__GongNoteShape__00000006_.IsExpanded = false

	__GongStructShape__00000000_.Name = `Default-Docx`
	__GongStructShape__00000000_.X = 67.000000
	__GongStructShape__00000000_.Y = 42.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Docx{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 83.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Document`
	__GongStructShape__00000001_.X = 426.000000
	__GongStructShape__00000001_.Y = 197.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Document{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 63.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-File`
	__GongStructShape__00000002_.X = 424.000000
	__GongStructShape__00000002_.Y = 35.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.File{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 83.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Default-Node`
	__GongStructShape__00000003_.X = 836.000000
	__GongStructShape__00000003_.Y = 196.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.Node{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 63.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Default-Body`
	__GongStructShape__00000004_.X = 431.000000
	__GongStructShape__00000004_.Y = 377.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.Body{}
	__GongStructShape__00000004_.Width = 240.000000
	__GongStructShape__00000004_.Height = 63.000000
	__GongStructShape__00000004_.IsSelected = false

	__GongStructShape__00000005_.Name = `02 - Body-Body`
	__GongStructShape__00000005_.X = 50.000000
	__GongStructShape__00000005_.Y = 58.000000
	__GongStructShape__00000005_.IdentifierMeta = ref_models.Body{}
	__GongStructShape__00000005_.Width = 240.000000
	__GongStructShape__00000005_.Height = 83.000000
	__GongStructShape__00000005_.IsSelected = false

	__GongStructShape__00000006_.Name = `02 - Body-Paragraph`
	__GongStructShape__00000006_.X = 564.000000
	__GongStructShape__00000006_.Y = 41.000000
	__GongStructShape__00000006_.IdentifierMeta = ref_models.Paragraph{}
	__GongStructShape__00000006_.Width = 240.000000
	__GongStructShape__00000006_.Height = 121.000000
	__GongStructShape__00000006_.IsSelected = false

	__GongStructShape__00000007_.Name = `02 - Body-Table`
	__GongStructShape__00000007_.X = 235.000000
	__GongStructShape__00000007_.Y = 239.000000
	__GongStructShape__00000007_.IdentifierMeta = ref_models.Table{}
	__GongStructShape__00000007_.Width = 240.000000
	__GongStructShape__00000007_.Height = 103.000000
	__GongStructShape__00000007_.IsSelected = false

	__GongStructShape__00000008_.Name = `02 - Body-TableRow`
	__GongStructShape__00000008_.X = 595.000000
	__GongStructShape__00000008_.Y = 239.000000
	__GongStructShape__00000008_.IdentifierMeta = ref_models.TableRow{}
	__GongStructShape__00000008_.Width = 240.000000
	__GongStructShape__00000008_.Height = 83.000000
	__GongStructShape__00000008_.IsSelected = false

	__GongStructShape__00000009_.Name = `02 - Body-TableProperties`
	__GongStructShape__00000009_.X = 232.000000
	__GongStructShape__00000009_.Y = 406.000000
	__GongStructShape__00000009_.IdentifierMeta = ref_models.TableProperties{}
	__GongStructShape__00000009_.Width = 240.000000
	__GongStructShape__00000009_.Height = 103.000000
	__GongStructShape__00000009_.IsSelected = false

	__GongStructShape__00000010_.Name = `02 - Body-TableStyle`
	__GongStructShape__00000010_.X = 235.000000
	__GongStructShape__00000010_.Y = 566.000000
	__GongStructShape__00000010_.IdentifierMeta = ref_models.TableStyle{}
	__GongStructShape__00000010_.Width = 240.000000
	__GongStructShape__00000010_.Height = 63.000000
	__GongStructShape__00000010_.IsSelected = false

	__GongStructShape__00000011_.Name = `02 - Body-TableColumn`
	__GongStructShape__00000011_.X = 956.000000
	__GongStructShape__00000011_.Y = 237.000000
	__GongStructShape__00000011_.IdentifierMeta = ref_models.TableColumn{}
	__GongStructShape__00000011_.Width = 240.000000
	__GongStructShape__00000011_.Height = 83.000000
	__GongStructShape__00000011_.IsSelected = false

	__GongStructShape__00000013_.Name = `Default-Document`
	__GongStructShape__00000013_.X = 107.000000
	__GongStructShape__00000013_.Y = 145.000000
	__GongStructShape__00000013_.IdentifierMeta = ref_models.Document{}
	__GongStructShape__00000013_.Width = 240.000000
	__GongStructShape__00000013_.Height = 63.000000
	__GongStructShape__00000013_.IsSelected = false

	__GongStructShape__00000014_.Name = `Default-Paragraph`
	__GongStructShape__00000014_.X = 625.000000
	__GongStructShape__00000014_.Y = 119.000000
	__GongStructShape__00000014_.IdentifierMeta = ref_models.Paragraph{}
	__GongStructShape__00000014_.Width = 240.000000
	__GongStructShape__00000014_.Height = 361.000000
	__GongStructShape__00000014_.IsSelected = false

	__GongStructShape__00000015_.Name = `Default-Body`
	__GongStructShape__00000015_.X = 109.000000
	__GongStructShape__00000015_.Y = 381.000000
	__GongStructShape__00000015_.IdentifierMeta = ref_models.Body{}
	__GongStructShape__00000015_.Width = 240.000000
	__GongStructShape__00000015_.Height = 63.000000
	__GongStructShape__00000015_.IsSelected = false

	__GongStructShape__00000016_.Name = `02 - Node-Node`
	__GongStructShape__00000016_.X = 575.000000
	__GongStructShape__00000016_.Y = 215.000000
	__GongStructShape__00000016_.IdentifierMeta = ref_models.Node{}
	__GongStructShape__00000016_.Width = 240.000000
	__GongStructShape__00000016_.Height = 697.000000
	__GongStructShape__00000016_.IsSelected = false

	__GongStructShape__00000017_.Name = `02 - Node-Paragraph`
	__GongStructShape__00000017_.X = 1107.000000
	__GongStructShape__00000017_.Y = 289.000000
	__GongStructShape__00000017_.IdentifierMeta = ref_models.Paragraph{}
	__GongStructShape__00000017_.Width = 240.000000
	__GongStructShape__00000017_.Height = 63.000000
	__GongStructShape__00000017_.IsSelected = false

	__GongStructShape__00000018_.Name = `02 - Node-Document`
	__GongStructShape__00000018_.X = 1107.000000
	__GongStructShape__00000018_.Y = 206.000000
	__GongStructShape__00000018_.IdentifierMeta = ref_models.Document{}
	__GongStructShape__00000018_.Width = 240.000000
	__GongStructShape__00000018_.Height = 63.000000
	__GongStructShape__00000018_.IsSelected = false

	__GongStructShape__00000019_.Name = `02 - Node-ParagraphProperties`
	__GongStructShape__00000019_.X = 1110.000000
	__GongStructShape__00000019_.Y = 369.000000
	__GongStructShape__00000019_.IdentifierMeta = ref_models.ParagraphProperties{}
	__GongStructShape__00000019_.Width = 240.000000
	__GongStructShape__00000019_.Height = 63.000000
	__GongStructShape__00000019_.IsSelected = false

	__GongStructShape__00000020_.Name = `02 - Node-ParagraphStyle`
	__GongStructShape__00000020_.X = 1110.000000
	__GongStructShape__00000020_.Y = 450.000000
	__GongStructShape__00000020_.IdentifierMeta = ref_models.ParagraphStyle{}
	__GongStructShape__00000020_.Width = 240.000000
	__GongStructShape__00000020_.Height = 63.000000
	__GongStructShape__00000020_.IsSelected = false

	__GongStructShape__00000021_.Name = `02 - Node-Rune`
	__GongStructShape__00000021_.X = 1108.000000
	__GongStructShape__00000021_.Y = 529.000000
	__GongStructShape__00000021_.IdentifierMeta = ref_models.Rune{}
	__GongStructShape__00000021_.Width = 240.000000
	__GongStructShape__00000021_.Height = 63.000000
	__GongStructShape__00000021_.IsSelected = false

	__GongStructShape__00000022_.Name = `02 - Node-RuneProperties`
	__GongStructShape__00000022_.X = 1110.000000
	__GongStructShape__00000022_.Y = 605.000000
	__GongStructShape__00000022_.IdentifierMeta = ref_models.RuneProperties{}
	__GongStructShape__00000022_.Width = 240.000000
	__GongStructShape__00000022_.Height = 63.000000
	__GongStructShape__00000022_.IsSelected = false

	__GongStructShape__00000023_.Name = `02 - Node-Table`
	__GongStructShape__00000023_.X = 191.000000
	__GongStructShape__00000023_.Y = 224.000000
	__GongStructShape__00000023_.IdentifierMeta = ref_models.Table{}
	__GongStructShape__00000023_.Width = 240.000000
	__GongStructShape__00000023_.Height = 63.000000
	__GongStructShape__00000023_.IsSelected = false

	__GongStructShape__00000024_.Name = `02 - Node-TableColumn`
	__GongStructShape__00000024_.X = 189.000000
	__GongStructShape__00000024_.Y = 454.000000
	__GongStructShape__00000024_.IdentifierMeta = ref_models.TableColumn{}
	__GongStructShape__00000024_.Width = 240.000000
	__GongStructShape__00000024_.Height = 63.000000
	__GongStructShape__00000024_.IsSelected = false

	__GongStructShape__00000025_.Name = `02 - Node-TableProperties`
	__GongStructShape__00000025_.X = 190.000000
	__GongStructShape__00000025_.Y = 527.000000
	__GongStructShape__00000025_.IdentifierMeta = ref_models.TableProperties{}
	__GongStructShape__00000025_.Width = 240.000000
	__GongStructShape__00000025_.Height = 63.000000
	__GongStructShape__00000025_.IsSelected = false

	__GongStructShape__00000026_.Name = `02 - Node-TableRow`
	__GongStructShape__00000026_.X = 192.000000
	__GongStructShape__00000026_.Y = 297.000000
	__GongStructShape__00000026_.IdentifierMeta = ref_models.TableRow{}
	__GongStructShape__00000026_.Width = 240.000000
	__GongStructShape__00000026_.Height = 63.000000
	__GongStructShape__00000026_.IsSelected = false

	__GongStructShape__00000027_.Name = `02 - Node-TableStyle`
	__GongStructShape__00000027_.X = 184.000000
	__GongStructShape__00000027_.Y = 382.000000
	__GongStructShape__00000027_.IdentifierMeta = ref_models.TableStyle{}
	__GongStructShape__00000027_.Width = 240.000000
	__GongStructShape__00000027_.Height = 63.000000
	__GongStructShape__00000027_.IsSelected = false

	__GongStructShape__00000028_.Name = `02 - Node-Text`
	__GongStructShape__00000028_.X = 193.000000
	__GongStructShape__00000028_.Y = 613.000000
	__GongStructShape__00000028_.IdentifierMeta = ref_models.Text{}
	__GongStructShape__00000028_.Width = 240.000000
	__GongStructShape__00000028_.Height = 63.000000
	__GongStructShape__00000028_.IsSelected = false

	__GongStructShape__00000029_.Name = `Default-ParagraphProperties`
	__GongStructShape__00000029_.X = 1085.000000
	__GongStructShape__00000029_.Y = 116.000000
	__GongStructShape__00000029_.IdentifierMeta = ref_models.ParagraphProperties{}
	__GongStructShape__00000029_.Width = 240.000000
	__GongStructShape__00000029_.Height = 63.000000
	__GongStructShape__00000029_.IsSelected = false

	__GongStructShape__00000030_.Name = `Default-ParagraphStyle`
	__GongStructShape__00000030_.X = 1089.000000
	__GongStructShape__00000030_.Y = 228.000000
	__GongStructShape__00000030_.IdentifierMeta = ref_models.ParagraphStyle{}
	__GongStructShape__00000030_.Width = 240.000000
	__GongStructShape__00000030_.Height = 63.000000
	__GongStructShape__00000030_.IsSelected = false

	__GongStructShape__00000031_.Name = `Default-Rune`
	__GongStructShape__00000031_.X = 1097.000000
	__GongStructShape__00000031_.Y = 346.000000
	__GongStructShape__00000031_.IdentifierMeta = ref_models.Rune{}
	__GongStructShape__00000031_.Width = 240.000000
	__GongStructShape__00000031_.Height = 83.000000
	__GongStructShape__00000031_.IsSelected = false

	__GongStructShape__00000032_.Name = `Default-RuneProperties`
	__GongStructShape__00000032_.X = 1094.000000
	__GongStructShape__00000032_.Y = 588.000000
	__GongStructShape__00000032_.IdentifierMeta = ref_models.RuneProperties{}
	__GongStructShape__00000032_.Width = 240.000000
	__GongStructShape__00000032_.Height = 143.000000
	__GongStructShape__00000032_.IsSelected = false

	__GongStructShape__00000033_.Name = `Default-Text`
	__GongStructShape__00000033_.X = 631.000000
	__GongStructShape__00000033_.Y = 525.000000
	__GongStructShape__00000033_.IdentifierMeta = ref_models.Text{}
	__GongStructShape__00000033_.Width = 240.000000
	__GongStructShape__00000033_.Height = 103.000000
	__GongStructShape__00000033_.IsSelected = false

	__LinkShape__00000000_.Name = `Document`
	__LinkShape__00000000_.IdentifierMeta = ref_models.Docx{}.Document
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.Document{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 613.000000
	__LinkShape__00000000_.Y = 84.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000000_.StartRatio = 0.461491
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.500000
	__LinkShape__00000000_.CornerOffsetRatio = 1.072454

	__LinkShape__00000001_.Name = `Files`
	__LinkShape__00000001_.IdentifierMeta = ref_models.Docx{}.Files
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.File{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 602.500000
	__LinkShape__00000001_.Y = 81.000000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.500000
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.500000
	__LinkShape__00000001_.CornerOffsetRatio = 1.380000

	__LinkShape__00000002_.Name = `Root`
	__LinkShape__00000002_.IdentifierMeta = ref_models.Document{}.Root
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 991.000000
	__LinkShape__00000002_.Y = 228.000000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.500000
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.500000
	__LinkShape__00000002_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `File`
	__LinkShape__00000003_.IdentifierMeta = ref_models.Document{}.File
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.File{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 785.000000
	__LinkShape__00000003_.Y = 147.500000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000003_.StartRatio = 0.469824
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000003_.EndRatio = 0.457324
	__LinkShape__00000003_.CornerOffsetRatio = -0.650577

	__LinkShape__00000004_.Name = `Body`
	__LinkShape__00000004_.IdentifierMeta = ref_models.Document{}.Body
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.Body{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 788.500000
	__LinkShape__00000004_.Y = 318.500000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_.StartRatio = 0.523991
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_.EndRatio = 0.503158
	__LinkShape__00000004_.CornerOffsetRatio = 1.254185

	__LinkShape__00000005_.Name = `Paragraphs`
	__LinkShape__00000005_.IdentifierMeta = ref_models.Body{}.Paragraphs
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.Paragraph{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.MANY
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 669.500000
	__LinkShape__00000005_.Y = 138.500000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.StartRatio = 0.500000
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.EndRatio = 0.139576
	__LinkShape__00000005_.CornerOffsetRatio = 1.380000

	__LinkShape__00000006_.Name = `TableRows`
	__LinkShape__00000006_.IdentifierMeta = ref_models.Table{}.TableRows
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.TableRow{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.MANY
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 1103.000000
	__LinkShape__00000006_.Y = 299.500000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.StartRatio = 0.500000
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.EndRatio = 0.500000
	__LinkShape__00000006_.CornerOffsetRatio = 1.380000

	__LinkShape__00000007_.Name = `TableProperties`
	__LinkShape__00000007_.IdentifierMeta = ref_models.Table{}.TableProperties
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.TableProperties{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 921.500000
	__LinkShape__00000007_.Y = 383.000000
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.StartRatio = 0.500000
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.EndRatio = 0.261055
	__LinkShape__00000007_.CornerOffsetRatio = -0.705176

	__LinkShape__00000008_.Name = `TableStyle`
	__LinkShape__00000008_.IdentifierMeta = ref_models.TableProperties{}.TableStyle
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.TableStyle{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 921.500000
	__LinkShape__00000008_.Y = 546.500000
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.StartRatio = 0.765909
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.EndRatio = 0.500000
	__LinkShape__00000008_.CornerOffsetRatio = -0.663509

	__LinkShape__00000009_.Name = `LastParagraph`
	__LinkShape__00000009_.IdentifierMeta = ref_models.Body{}.LastParagraph
	__LinkShape__00000009_.FieldTypeIdentifierMeta = ref_models.Paragraph{}
	__LinkShape__00000009_.FieldOffsetX = 0.000000
	__LinkShape__00000009_.FieldOffsetY = 0.000000
	__LinkShape__00000009_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000009_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.SourceMultiplicity = models.MANY
	__LinkShape__00000009_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.X = 669.500000
	__LinkShape__00000009_.Y = 138.500000
	__LinkShape__00000009_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.StartRatio = 0.500000
	__LinkShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.EndRatio = 0.932964
	__LinkShape__00000009_.CornerOffsetRatio = 1.380000

	__LinkShape__00000010_.Name = `Tables`
	__LinkShape__00000010_.IdentifierMeta = ref_models.Body{}.Tables
	__LinkShape__00000010_.FieldTypeIdentifierMeta = ref_models.Table{}
	__LinkShape__00000010_.FieldOffsetX = 0.000000
	__LinkShape__00000010_.FieldOffsetY = 0.000000
	__LinkShape__00000010_.TargetMultiplicity = models.MANY
	__LinkShape__00000010_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.SourceMultiplicity = models.MANY
	__LinkShape__00000010_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.X = 669.500000
	__LinkShape__00000010_.Y = 214.000000
	__LinkShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000010_.StartRatio = 0.319824
	__LinkShape__00000010_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000010_.EndRatio = 0.144550
	__LinkShape__00000010_.CornerOffsetRatio = 1.010707

	__LinkShape__00000011_.Name = `Paragraphs`
	__LinkShape__00000011_.IdentifierMeta = ref_models.TableColumn{}.Paragraphs
	__LinkShape__00000011_.FieldTypeIdentifierMeta = ref_models.Paragraph{}
	__LinkShape__00000011_.FieldOffsetX = 0.000000
	__LinkShape__00000011_.FieldOffsetY = 0.000000
	__LinkShape__00000011_.TargetMultiplicity = models.MANY
	__LinkShape__00000011_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.SourceMultiplicity = models.MANY
	__LinkShape__00000011_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.X = 1284.000000
	__LinkShape__00000011_.Y = 175.000000
	__LinkShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000011_.StartRatio = 0.494824
	__LinkShape__00000011_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.EndRatio = 0.500000
	__LinkShape__00000011_.CornerOffsetRatio = -0.477958

	__LinkShape__00000012_.Name = `TableColumns`
	__LinkShape__00000012_.IdentifierMeta = ref_models.TableRow{}.TableColumns
	__LinkShape__00000012_.FieldTypeIdentifierMeta = ref_models.TableColumn{}
	__LinkShape__00000012_.FieldOffsetX = 0.000000
	__LinkShape__00000012_.FieldOffsetY = 0.000000
	__LinkShape__00000012_.TargetMultiplicity = models.MANY
	__LinkShape__00000012_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.SourceMultiplicity = models.MANY
	__LinkShape__00000012_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.X = 1463.500000
	__LinkShape__00000012_.Y = 288.500000
	__LinkShape__00000012_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000012_.StartRatio = 0.500000
	__LinkShape__00000012_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000012_.EndRatio = 0.500000
	__LinkShape__00000012_.CornerOffsetRatio = 1.380000

	__LinkShape__00000015_.Name = `Body`
	__LinkShape__00000015_.IdentifierMeta = ref_models.Document{}.Body
	__LinkShape__00000015_.FieldTypeIdentifierMeta = ref_models.Body{}
	__LinkShape__00000015_.FieldOffsetX = 0.000000
	__LinkShape__00000015_.FieldOffsetY = 0.000000
	__LinkShape__00000015_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000015_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000015_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000015_.SourceMultiplicity = models.MANY
	__LinkShape__00000015_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000015_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000015_.X = 426.500000
	__LinkShape__00000015_.Y = 157.500000
	__LinkShape__00000015_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000015_.StartRatio = 0.319824
	__LinkShape__00000015_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000015_.EndRatio = 0.311491
	__LinkShape__00000015_.CornerOffsetRatio = 2.395058

	__LinkShape__00000016_.Name = `Nodes`
	__LinkShape__00000016_.IdentifierMeta = ref_models.Node{}.Nodes
	__LinkShape__00000016_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000016_.FieldOffsetX = 0.000000
	__LinkShape__00000016_.FieldOffsetY = 0.000000
	__LinkShape__00000016_.TargetMultiplicity = models.MANY
	__LinkShape__00000016_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000016_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000016_.SourceMultiplicity = models.MANY
	__LinkShape__00000016_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000016_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000016_.X = 703.000000
	__LinkShape__00000016_.Y = 119.500000
	__LinkShape__00000016_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000016_.StartRatio = 0.182324
	__LinkShape__00000016_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000016_.EndRatio = 0.782324
	__LinkShape__00000016_.CornerOffsetRatio = 1.128592

	__LinkShape__00000017_.Name = `Node`
	__LinkShape__00000017_.IdentifierMeta = ref_models.Paragraph{}.Node
	__LinkShape__00000017_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000017_.FieldOffsetX = 0.000000
	__LinkShape__00000017_.FieldOffsetY = 0.000000
	__LinkShape__00000017_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000017_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000017_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000017_.SourceMultiplicity = models.MANY
	__LinkShape__00000017_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000017_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000017_.X = 503.500000
	__LinkShape__00000017_.Y = 117.000000
	__LinkShape__00000017_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000017_.StartRatio = 0.500000
	__LinkShape__00000017_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000017_.EndRatio = 0.164998
	__LinkShape__00000017_.CornerOffsetRatio = -0.371842

	__LinkShape__00000018_.Name = `Root`
	__LinkShape__00000018_.IdentifierMeta = ref_models.Document{}.Root
	__LinkShape__00000018_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000018_.FieldOffsetX = 0.000000
	__LinkShape__00000018_.FieldOffsetY = 0.000000
	__LinkShape__00000018_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000018_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000018_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000018_.SourceMultiplicity = models.MANY
	__LinkShape__00000018_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000018_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000018_.X = 526.500000
	__LinkShape__00000018_.Y = 123.500000
	__LinkShape__00000018_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000018_.StartRatio = 0.500000
	__LinkShape__00000018_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000018_.EndRatio = 0.037308
	__LinkShape__00000018_.CornerOffsetRatio = -0.530176

	__LinkShape__00000019_.Name = `Node`
	__LinkShape__00000019_.IdentifierMeta = ref_models.ParagraphProperties{}.Node
	__LinkShape__00000019_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000019_.FieldOffsetX = 0.000000
	__LinkShape__00000019_.FieldOffsetY = 0.000000
	__LinkShape__00000019_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000019_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000019_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000019_.SourceMultiplicity = models.MANY
	__LinkShape__00000019_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000019_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000019_.X = 493.500000
	__LinkShape__00000019_.Y = 94.000000
	__LinkShape__00000019_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000019_.StartRatio = 0.500000
	__LinkShape__00000019_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000019_.EndRatio = 0.282645
	__LinkShape__00000019_.CornerOffsetRatio = -0.534342

	__LinkShape__00000020_.Name = `Node`
	__LinkShape__00000020_.IdentifierMeta = ref_models.ParagraphStyle{}.Node
	__LinkShape__00000020_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000020_.FieldOffsetX = 0.000000
	__LinkShape__00000020_.FieldOffsetY = 0.000000
	__LinkShape__00000020_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000020_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000020_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000020_.SourceMultiplicity = models.MANY
	__LinkShape__00000020_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000020_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000020_.X = 535.500000
	__LinkShape__00000020_.Y = 72.500000
	__LinkShape__00000020_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000020_.StartRatio = 0.500000
	__LinkShape__00000020_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000020_.EndRatio = 0.398858
	__LinkShape__00000020_.CornerOffsetRatio = -0.584342

	__LinkShape__00000021_.Name = `Node`
	__LinkShape__00000021_.IdentifierMeta = ref_models.Rune{}.Node
	__LinkShape__00000021_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000021_.FieldOffsetX = 0.000000
	__LinkShape__00000021_.FieldOffsetY = 0.000000
	__LinkShape__00000021_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000021_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000021_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000021_.SourceMultiplicity = models.MANY
	__LinkShape__00000021_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000021_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000021_.X = 526.000000
	__LinkShape__00000021_.Y = 70.500000
	__LinkShape__00000021_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000021_.StartRatio = 0.500000
	__LinkShape__00000021_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000021_.EndRatio = 0.504310
	__LinkShape__00000021_.CornerOffsetRatio = -0.730176

	__LinkShape__00000022_.Name = `Node`
	__LinkShape__00000022_.IdentifierMeta = ref_models.RuneProperties{}.Node
	__LinkShape__00000022_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000022_.FieldOffsetX = 0.000000
	__LinkShape__00000022_.FieldOffsetY = 0.000000
	__LinkShape__00000022_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000022_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000022_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000022_.SourceMultiplicity = models.MANY
	__LinkShape__00000022_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000022_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000022_.X = 499.500000
	__LinkShape__00000022_.Y = 110.000000
	__LinkShape__00000022_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000022_.StartRatio = 0.500000
	__LinkShape__00000022_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000022_.EndRatio = 0.620701
	__LinkShape__00000022_.CornerOffsetRatio = -0.646842

	__LinkShape__00000023_.Name = `Node`
	__LinkShape__00000023_.IdentifierMeta = ref_models.Table{}.Node
	__LinkShape__00000023_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000023_.FieldOffsetX = 0.000000
	__LinkShape__00000023_.FieldOffsetY = 0.000000
	__LinkShape__00000023_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000023_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000023_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000023_.SourceMultiplicity = models.MANY
	__LinkShape__00000023_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000023_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000023_.X = 506.000000
	__LinkShape__00000023_.Y = 98.000000
	__LinkShape__00000023_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000023_.StartRatio = 0.500000
	__LinkShape__00000023_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000023_.EndRatio = 0.066720
	__LinkShape__00000023_.CornerOffsetRatio = 1.290658

	__LinkShape__00000024_.Name = `Node`
	__LinkShape__00000024_.IdentifierMeta = ref_models.TableColumn{}.Node
	__LinkShape__00000024_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000024_.FieldOffsetX = 0.000000
	__LinkShape__00000024_.FieldOffsetY = 0.000000
	__LinkShape__00000024_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000024_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000024_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000024_.SourceMultiplicity = models.MANY
	__LinkShape__00000024_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000024_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000024_.X = 524.500000
	__LinkShape__00000024_.Y = 111.000000
	__LinkShape__00000024_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000024_.StartRatio = 0.500000
	__LinkShape__00000024_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000024_.EndRatio = 0.399575
	__LinkShape__00000024_.CornerOffsetRatio = 1.290658

	__LinkShape__00000025_.Name = `Node`
	__LinkShape__00000025_.IdentifierMeta = ref_models.TableRow{}.Node
	__LinkShape__00000025_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000025_.FieldOffsetX = 0.000000
	__LinkShape__00000025_.FieldOffsetY = 0.000000
	__LinkShape__00000025_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000025_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000025_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000025_.SourceMultiplicity = models.MANY
	__LinkShape__00000025_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000025_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000025_.X = 743.500000
	__LinkShape__00000025_.Y = 287.500000
	__LinkShape__00000025_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000025_.StartRatio = 0.500000
	__LinkShape__00000025_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000025_.EndRatio = 0.152803
	__LinkShape__00000025_.CornerOffsetRatio = 1.380000

	__LinkShape__00000026_.Name = `Node`
	__LinkShape__00000026_.IdentifierMeta = ref_models.TableProperties{}.Node
	__LinkShape__00000026_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000026_.FieldOffsetX = 0.000000
	__LinkShape__00000026_.FieldOffsetY = 0.000000
	__LinkShape__00000026_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000026_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000026_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000026_.SourceMultiplicity = models.MANY
	__LinkShape__00000026_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000026_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000026_.X = 742.500000
	__LinkShape__00000026_.Y = 402.500000
	__LinkShape__00000026_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000026_.StartRatio = 0.500000
	__LinkShape__00000026_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000026_.EndRatio = 0.500000
	__LinkShape__00000026_.CornerOffsetRatio = 1.380000

	__LinkShape__00000027_.Name = `Node`
	__LinkShape__00000027_.IdentifierMeta = ref_models.TableStyle{}.Node
	__LinkShape__00000027_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000027_.FieldOffsetX = 0.000000
	__LinkShape__00000027_.FieldOffsetY = 0.000000
	__LinkShape__00000027_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000027_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000027_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000027_.SourceMultiplicity = models.MANY
	__LinkShape__00000027_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000027_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000027_.X = 739.500000
	__LinkShape__00000027_.Y = 330.000000
	__LinkShape__00000027_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000027_.StartRatio = 0.500000
	__LinkShape__00000027_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000027_.EndRatio = 0.296275
	__LinkShape__00000027_.CornerOffsetRatio = 1.380000

	__LinkShape__00000028_.Name = `Node`
	__LinkShape__00000028_.IdentifierMeta = ref_models.Text{}.Node
	__LinkShape__00000028_.FieldTypeIdentifierMeta = ref_models.Node{}
	__LinkShape__00000028_.FieldOffsetX = 0.000000
	__LinkShape__00000028_.FieldOffsetY = 0.000000
	__LinkShape__00000028_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000028_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000028_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000028_.SourceMultiplicity = models.MANY
	__LinkShape__00000028_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000028_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000028_.X = 744.000000
	__LinkShape__00000028_.Y = 445.500000
	__LinkShape__00000028_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000028_.StartRatio = 0.500000
	__LinkShape__00000028_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000028_.EndRatio = 0.619087
	__LinkShape__00000028_.CornerOffsetRatio = 1.380000

	__LinkShape__00000029_.Name = `Paragraphs`
	__LinkShape__00000029_.IdentifierMeta = ref_models.Body{}.Paragraphs
	__LinkShape__00000029_.FieldTypeIdentifierMeta = ref_models.Paragraph{}
	__LinkShape__00000029_.FieldOffsetX = 0.000000
	__LinkShape__00000029_.FieldOffsetY = 0.000000
	__LinkShape__00000029_.TargetMultiplicity = models.MANY
	__LinkShape__00000029_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000029_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000029_.SourceMultiplicity = models.MANY
	__LinkShape__00000029_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000029_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000029_.X = 689.000000
	__LinkShape__00000029_.Y = 415.000000
	__LinkShape__00000029_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000029_.StartRatio = 0.500000
	__LinkShape__00000029_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000029_.EndRatio = 0.141367
	__LinkShape__00000029_.CornerOffsetRatio = 1.380000

	__LinkShape__00000030_.Name = `LastParagraph`
	__LinkShape__00000030_.IdentifierMeta = ref_models.Body{}.LastParagraph
	__LinkShape__00000030_.FieldTypeIdentifierMeta = ref_models.Paragraph{}
	__LinkShape__00000030_.FieldOffsetX = 0.000000
	__LinkShape__00000030_.FieldOffsetY = 0.000000
	__LinkShape__00000030_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000030_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000030_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000030_.SourceMultiplicity = models.MANY
	__LinkShape__00000030_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000030_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000030_.X = 689.000000
	__LinkShape__00000030_.Y = 415.000000
	__LinkShape__00000030_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000030_.StartRatio = 0.500000
	__LinkShape__00000030_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000030_.EndRatio = 0.935760
	__LinkShape__00000030_.CornerOffsetRatio = 1.380000

	__LinkShape__00000031_.Name = `ParagraphStyle`
	__LinkShape__00000031_.IdentifierMeta = ref_models.ParagraphProperties{}.ParagraphStyle
	__LinkShape__00000031_.FieldTypeIdentifierMeta = ref_models.ParagraphStyle{}
	__LinkShape__00000031_.FieldOffsetX = 0.000000
	__LinkShape__00000031_.FieldOffsetY = 0.000000
	__LinkShape__00000031_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000031_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000031_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000031_.SourceMultiplicity = models.MANY
	__LinkShape__00000031_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000031_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000031_.X = 1382.500000
	__LinkShape__00000031_.Y = 283.000000
	__LinkShape__00000031_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000031_.StartRatio = 0.500000
	__LinkShape__00000031_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000031_.EndRatio = 0.500000
	__LinkShape__00000031_.CornerOffsetRatio = 1.557324

	__LinkShape__00000032_.Name = `ParagraphProperties`
	__LinkShape__00000032_.IdentifierMeta = ref_models.Paragraph{}.ParagraphProperties
	__LinkShape__00000032_.FieldTypeIdentifierMeta = ref_models.ParagraphProperties{}
	__LinkShape__00000032_.FieldOffsetX = 0.000000
	__LinkShape__00000032_.FieldOffsetY = 0.000000
	__LinkShape__00000032_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000032_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000032_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000032_.SourceMultiplicity = models.MANY
	__LinkShape__00000032_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000032_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000032_.X = 1178.000000
	__LinkShape__00000032_.Y = 278.500000
	__LinkShape__00000032_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000032_.StartRatio = 0.092020
	__LinkShape__00000032_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000032_.EndRatio = 0.500000
	__LinkShape__00000032_.CornerOffsetRatio = 1.380000

	__LinkShape__00000033_.Name = `Runes`
	__LinkShape__00000033_.IdentifierMeta = ref_models.Paragraph{}.Runes
	__LinkShape__00000033_.FieldTypeIdentifierMeta = ref_models.Rune{}
	__LinkShape__00000033_.FieldOffsetX = 0.000000
	__LinkShape__00000033_.FieldOffsetY = 0.000000
	__LinkShape__00000033_.TargetMultiplicity = models.MANY
	__LinkShape__00000033_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000033_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000033_.SourceMultiplicity = models.MANY
	__LinkShape__00000033_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000033_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000033_.X = 685.000000
	__LinkShape__00000033_.Y = 282.000000
	__LinkShape__00000033_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000033_.StartRatio = 0.778237
	__LinkShape__00000033_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000033_.EndRatio = 0.500000
	__LinkShape__00000033_.CornerOffsetRatio = 1.380000

	__LinkShape__00000034_.Name = `RuneProperties`
	__LinkShape__00000034_.IdentifierMeta = ref_models.Rune{}.RuneProperties
	__LinkShape__00000034_.FieldTypeIdentifierMeta = ref_models.RuneProperties{}
	__LinkShape__00000034_.FieldOffsetX = 0.000000
	__LinkShape__00000034_.FieldOffsetY = 0.000000
	__LinkShape__00000034_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000034_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000034_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000034_.SourceMultiplicity = models.MANY
	__LinkShape__00000034_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000034_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000034_.X = 1458.500000
	__LinkShape__00000034_.Y = 451.000000
	__LinkShape__00000034_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000034_.StartRatio = 0.500000
	__LinkShape__00000034_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000034_.EndRatio = 0.500000
	__LinkShape__00000034_.CornerOffsetRatio = 1.380000

	__LinkShape__00000035_.Name = `EnclosingParagraph`
	__LinkShape__00000035_.IdentifierMeta = ref_models.Rune{}.EnclosingParagraph
	__LinkShape__00000035_.FieldTypeIdentifierMeta = ref_models.Paragraph{}
	__LinkShape__00000035_.FieldOffsetX = 0.000000
	__LinkShape__00000035_.FieldOffsetY = 0.000000
	__LinkShape__00000035_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000035_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000035_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000035_.SourceMultiplicity = models.MANY
	__LinkShape__00000035_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000035_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000035_.X = 1221.000000
	__LinkShape__00000035_.Y = 274.000000
	__LinkShape__00000035_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000035_.StartRatio = 0.453158
	__LinkShape__00000035_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000035_.EndRatio = 0.564161
	__LinkShape__00000035_.CornerOffsetRatio = -0.079772

	__LinkShape__00000037_.Name = `Next`
	__LinkShape__00000037_.IdentifierMeta = ref_models.Paragraph{}.Next
	__LinkShape__00000037_.FieldTypeIdentifierMeta = ref_models.Paragraph{}
	__LinkShape__00000037_.FieldOffsetX = 0.000000
	__LinkShape__00000037_.FieldOffsetY = 0.000000
	__LinkShape__00000037_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000037_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000037_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000037_.SourceMultiplicity = models.MANY
	__LinkShape__00000037_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000037_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000037_.X = 985.000000
	__LinkShape__00000037_.Y = 299.500000
	__LinkShape__00000037_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000037_.StartRatio = 0.304720
	__LinkShape__00000037_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000037_.EndRatio = 0.465385
	__LinkShape__00000037_.CornerOffsetRatio = -0.471842

	__LinkShape__00000038_.Name = `Previous`
	__LinkShape__00000038_.IdentifierMeta = ref_models.Paragraph{}.Previous
	__LinkShape__00000038_.FieldTypeIdentifierMeta = ref_models.Paragraph{}
	__LinkShape__00000038_.FieldOffsetX = 0.000000
	__LinkShape__00000038_.FieldOffsetY = 0.000000
	__LinkShape__00000038_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000038_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000038_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000038_.SourceMultiplicity = models.MANY
	__LinkShape__00000038_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000038_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000038_.X = 985.000000
	__LinkShape__00000038_.Y = 299.500000
	__LinkShape__00000038_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000038_.StartRatio = 0.576188
	__LinkShape__00000038_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000038_.EndRatio = 0.745163
	__LinkShape__00000038_.CornerOffsetRatio = -0.434342

	__LinkShape__00000039_.Name = `Text`
	__LinkShape__00000039_.IdentifierMeta = ref_models.Rune{}.Text
	__LinkShape__00000039_.FieldTypeIdentifierMeta = ref_models.Text{}
	__LinkShape__00000039_.FieldOffsetX = 0.000000
	__LinkShape__00000039_.FieldOffsetY = 0.000000
	__LinkShape__00000039_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000039_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000039_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000039_.SourceMultiplicity = models.MANY
	__LinkShape__00000039_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000039_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000039_.X = 1224.000000
	__LinkShape__00000039_.Y = 477.000000
	__LinkShape__00000039_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000039_.StartRatio = 0.178158
	__LinkShape__00000039_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000039_.EndRatio = 0.500000
	__LinkShape__00000039_.CornerOffsetRatio = 1.067865

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__Classdiagram__00000000_.GongNoteShapes = append(__Classdiagram__00000000_.GongNoteShapes, __GongNoteShape__00000000_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000005_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000006_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000007_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000008_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000009_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000010_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000011_)
	__Classdiagram__00000001_.GongNoteShapes = append(__Classdiagram__00000001_.GongNoteShapes, __GongNoteShape__00000001_)
	__Classdiagram__00000001_.GongNoteShapes = append(__Classdiagram__00000001_.GongNoteShapes, __GongNoteShape__00000002_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000013_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000014_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000015_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000029_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000030_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000031_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000032_)
	__Classdiagram__00000002_.GongStructShapes = append(__Classdiagram__00000002_.GongStructShapes, __GongStructShape__00000033_)
	__Classdiagram__00000002_.GongNoteShapes = append(__Classdiagram__00000002_.GongNoteShapes, __GongNoteShape__00000003_)
	__Classdiagram__00000002_.GongNoteShapes = append(__Classdiagram__00000002_.GongNoteShapes, __GongNoteShape__00000004_)
	__Classdiagram__00000002_.GongNoteShapes = append(__Classdiagram__00000002_.GongNoteShapes, __GongNoteShape__00000006_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000016_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000017_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000018_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000019_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000020_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000021_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000022_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000023_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000024_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000025_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000026_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000027_)
	__Classdiagram__00000003_.GongStructShapes = append(__Classdiagram__00000003_.GongStructShapes, __GongStructShape__00000028_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000001_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000002_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000003_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000002_
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000001_)
	__GongStructShape__00000005_.AttributeShapes = append(__GongStructShape__00000005_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000009_)
	__GongStructShape__00000005_.LinkShapes = append(__GongStructShape__00000005_.LinkShapes, __LinkShape__00000010_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000003_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000004_)
	__GongStructShape__00000007_.LinkShapes = append(__GongStructShape__00000007_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000007_.LinkShapes = append(__GongStructShape__00000007_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000008_.AttributeShapes = append(__GongStructShape__00000008_.AttributeShapes, __AttributeShape__00000005_)
	__GongStructShape__00000008_.LinkShapes = append(__GongStructShape__00000008_.LinkShapes, __LinkShape__00000012_)
	__GongStructShape__00000009_.AttributeShapes = append(__GongStructShape__00000009_.AttributeShapes, __AttributeShape__00000006_)
	__GongStructShape__00000009_.AttributeShapes = append(__GongStructShape__00000009_.AttributeShapes, __AttributeShape__00000007_)
	__GongStructShape__00000009_.LinkShapes = append(__GongStructShape__00000009_.LinkShapes, __LinkShape__00000008_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000008_)
	__GongStructShape__00000011_.LinkShapes = append(__GongStructShape__00000011_.LinkShapes, __LinkShape__00000011_)
	__GongStructShape__00000013_.LinkShapes = append(__GongStructShape__00000013_.LinkShapes, __LinkShape__00000015_)
	__GongStructShape__00000014_.AttributeShapes = append(__GongStructShape__00000014_.AttributeShapes, __AttributeShape__00000010_)
	__GongStructShape__00000014_.AttributeShapes = append(__GongStructShape__00000014_.AttributeShapes, __AttributeShape__00000019_)
	__GongStructShape__00000014_.LinkShapes = append(__GongStructShape__00000014_.LinkShapes, __LinkShape__00000032_)
	__GongStructShape__00000014_.LinkShapes = append(__GongStructShape__00000014_.LinkShapes, __LinkShape__00000033_)
	__GongStructShape__00000014_.LinkShapes = append(__GongStructShape__00000014_.LinkShapes, __LinkShape__00000037_)
	__GongStructShape__00000014_.LinkShapes = append(__GongStructShape__00000014_.LinkShapes, __LinkShape__00000038_)
	__GongStructShape__00000015_.LinkShapes = append(__GongStructShape__00000015_.LinkShapes, __LinkShape__00000029_)
	__GongStructShape__00000015_.LinkShapes = append(__GongStructShape__00000015_.LinkShapes, __LinkShape__00000030_)
	__GongStructShape__00000016_.LinkShapes = append(__GongStructShape__00000016_.LinkShapes, __LinkShape__00000016_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000017_)
	__GongStructShape__00000018_.LinkShapes = append(__GongStructShape__00000018_.LinkShapes, __LinkShape__00000018_)
	__GongStructShape__00000019_.LinkShapes = append(__GongStructShape__00000019_.LinkShapes, __LinkShape__00000019_)
	__GongStructShape__00000020_.LinkShapes = append(__GongStructShape__00000020_.LinkShapes, __LinkShape__00000020_)
	__GongStructShape__00000021_.LinkShapes = append(__GongStructShape__00000021_.LinkShapes, __LinkShape__00000021_)
	__GongStructShape__00000022_.LinkShapes = append(__GongStructShape__00000022_.LinkShapes, __LinkShape__00000022_)
	__GongStructShape__00000023_.LinkShapes = append(__GongStructShape__00000023_.LinkShapes, __LinkShape__00000023_)
	__GongStructShape__00000024_.LinkShapes = append(__GongStructShape__00000024_.LinkShapes, __LinkShape__00000024_)
	__GongStructShape__00000025_.LinkShapes = append(__GongStructShape__00000025_.LinkShapes, __LinkShape__00000026_)
	__GongStructShape__00000026_.LinkShapes = append(__GongStructShape__00000026_.LinkShapes, __LinkShape__00000025_)
	__GongStructShape__00000027_.LinkShapes = append(__GongStructShape__00000027_.LinkShapes, __LinkShape__00000027_)
	__GongStructShape__00000028_.LinkShapes = append(__GongStructShape__00000028_.LinkShapes, __LinkShape__00000028_)
	__GongStructShape__00000029_.LinkShapes = append(__GongStructShape__00000029_.LinkShapes, __LinkShape__00000031_)
	__GongStructShape__00000031_.AttributeShapes = append(__GongStructShape__00000031_.AttributeShapes, __AttributeShape__00000015_)
	__GongStructShape__00000031_.LinkShapes = append(__GongStructShape__00000031_.LinkShapes, __LinkShape__00000034_)
	__GongStructShape__00000031_.LinkShapes = append(__GongStructShape__00000031_.LinkShapes, __LinkShape__00000035_)
	__GongStructShape__00000031_.LinkShapes = append(__GongStructShape__00000031_.LinkShapes, __LinkShape__00000039_)
	__GongStructShape__00000032_.AttributeShapes = append(__GongStructShape__00000032_.AttributeShapes, __AttributeShape__00000011_)
	__GongStructShape__00000032_.AttributeShapes = append(__GongStructShape__00000032_.AttributeShapes, __AttributeShape__00000012_)
	__GongStructShape__00000032_.AttributeShapes = append(__GongStructShape__00000032_.AttributeShapes, __AttributeShape__00000013_)
	__GongStructShape__00000032_.AttributeShapes = append(__GongStructShape__00000032_.AttributeShapes, __AttributeShape__00000014_)
	__GongStructShape__00000033_.AttributeShapes = append(__GongStructShape__00000033_.AttributeShapes, __AttributeShape__00000016_)
	__GongStructShape__00000033_.AttributeShapes = append(__GongStructShape__00000033_.AttributeShapes, __AttributeShape__00000020_)
}
