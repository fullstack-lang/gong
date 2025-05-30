package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/load/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__FileToUpload__000000_DOORS_export_sample_1_reqif := (&models.FileToUpload{}).Stage(stage)

	// Setup of values

	__FileToUpload__000000_DOORS_export_sample_1_reqif.Name = `DOORS export sample 1.reqif`
	__FileToUpload__000000_DOORS_export_sample_1_reqif.Content = `<?xml version="1.0" encoding="UTF-8"?>
<REQ-IF xmlns="http://www.omg.org/spec/ReqIF/20110401/reqif.xsd"
    xmlns:xhtml="http://www.w3.org/1999/xhtml">
  <THE-HEADER>
    <REQ-IF-HEADER
        IDENTIFIER="_8da68f26-585c-4193-b6ca-a07cedff2a5c">
      <CREATION-TIME>2017-04-25T15:44:26.000+02:00</CREATION-TIME>
      <REPOSITORY-ID>524d2a940b0233d5</REPOSITORY-ID>
      <REQ-IF-TOOL-ID>IBM Rational DOORS</REQ-IF-TOOL-ID>
      <REQ-IF-VERSION>1.0</REQ-IF-VERSION>
      <SOURCE-TOOL-ID>IBM Rational DOORS</SOURCE-TOOL-ID>
      <TITLE>IBM Rational DOORS, 25 April 2017</TITLE>
    </REQ-IF-HEADER>
  </THE-HEADER>
  <CORE-CONTENT>
    <REQ-IF-CONTENT>
      <DATATYPES>
        <DATATYPE-DEFINITION-XHTML
            IDENTIFIER="_i6FNMMkgEee8KsfWrp9EJQ"
            LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
            LONG-NAME="String"/>
        <DATATYPE-DEFINITION-ENUMERATION
            IDENTIFIER="_ahyOEMk3Eee5A_N9aQFa1w"
            LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
            LONG-NAME="IE Object Type">
          <SPECIFIED-VALUES>
            <ENUM-VALUE
                IDENTIFIER="_gXZ9oMk3Eee5A_N9aQFa1w"
                LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
                LONG-NAME="Requirement">
              <PROPERTIES>
                <EMBEDDED-VALUE
                    KEY="0"
                    OTHER-CONTENT="Navy"/>
              </PROPERTIES>
            </ENUM-VALUE>
            <ENUM-VALUE
                IDENTIFIER="_hGeDEMk3Eee5A_N9aQFa1w"
                LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
                LONG-NAME="Function">
              <PROPERTIES>
                <EMBEDDED-VALUE
                    KEY="1"
                    OTHER-CONTENT="Purple"/>
              </PROPERTIES>
            </ENUM-VALUE>
            <ENUM-VALUE
                IDENTIFIER="_hZP1IMk3Eee5A_N9aQFa1w"
                LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
                LONG-NAME="Capability">
              <PROPERTIES>
                <EMBEDDED-VALUE
                    KEY="2"
                    OTHER-CONTENT="Sea_Green"/>
              </PROPERTIES>
            </ENUM-VALUE>
          </SPECIFIED-VALUES>
        </DATATYPE-DEFINITION-ENUMERATION>
      </DATATYPES>
      <SPEC-TYPES>
        <SPEC-OBJECT-TYPE
            IDENTIFIER="_LaeSsMkhEee8KsfWrp9EJQ"
            LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
            LONG-NAME="Requirement Type">
          <SPEC-ATTRIBUTES>
            <ATTRIBUTE-DEFINITION-XHTML
                IDENTIFIER="_cf7-0MkhEee8KsfWrp9EJQ"
                LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
                LONG-NAME="ReqIF.Text"
                IS-EDITABLE="true">
              <TYPE>
                <DATATYPE-DEFINITION-XHTML-REF>_i6FNMMkgEee8KsfWrp9EJQ</DATATYPE-DEFINITION-XHTML-REF>
              </TYPE>
            </ATTRIBUTE-DEFINITION-XHTML>
            <ATTRIBUTE-DEFINITION-XHTML
                IDENTIFIER="_25YvEMkyEee5A_N9aQFa1w"
                LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
                LONG-NAME="IE PUID"
                IS-EDITABLE="true">
              <TYPE>
                <DATATYPE-DEFINITION-XHTML-REF>_i6FNMMkgEee8KsfWrp9EJQ</DATATYPE-DEFINITION-XHTML-REF>
              </TYPE>
            </ATTRIBUTE-DEFINITION-XHTML>
            <ATTRIBUTE-DEFINITION-ENUMERATION
                IDENTIFIER="_0VDX0Mk3Eee5A_N9aQFa1w"
                LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
                LONG-NAME="IE Object Type"
                IS-EDITABLE="true"
                MULTI-VALUED="false">
              <TYPE>
                <DATATYPE-DEFINITION-ENUMERATION-REF>_ahyOEMk3Eee5A_N9aQFa1w</DATATYPE-DEFINITION-ENUMERATION-REF>
              </TYPE>
            </ATTRIBUTE-DEFINITION-ENUMERATION>
          </SPEC-ATTRIBUTES>
        </SPEC-OBJECT-TYPE>
        <SPECIFICATION-TYPE
            IDENTIFIER="_NAXScMkhEee8KsfWrp9EJQ"
            LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
            LONG-NAME="Module Type">
          <SPEC-ATTRIBUTES>
            <ATTRIBUTE-DEFINITION-XHTML
                IDENTIFIER="_jLB90MkhEee8KsfWrp9EJQ"
                LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
                LONG-NAME="ReqIF.Name"
                IS-EDITABLE="true">
              <TYPE>
                <DATATYPE-DEFINITION-XHTML-REF>_i6FNMMkgEee8KsfWrp9EJQ</DATATYPE-DEFINITION-XHTML-REF>
              </TYPE>
            </ATTRIBUTE-DEFINITION-XHTML>
          </SPEC-ATTRIBUTES>
        </SPECIFICATION-TYPE>
      </SPEC-TYPES>
      <SPEC-OBJECTS>
        <SPEC-OBJECT
            IDENTIFIER="_xen_QMkhEee8KsfWrp9EJQ"
            LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
            LONG-NAME="Requirement-1">
          <VALUES>
            <ATTRIBUTE-VALUE-XHTML>
              <DEFINITION>
                <ATTRIBUTE-DEFINITION-XHTML-REF>_25YvEMkyEee5A_N9aQFa1w</ATTRIBUTE-DEFINITION-XHTML-REF>
              </DEFINITION>
              <THE-VALUE>
                <xhtml:div>PUID-1</xhtml:div>
              </THE-VALUE>
            </ATTRIBUTE-VALUE-XHTML>
            <ATTRIBUTE-VALUE-XHTML>
              <DEFINITION>
                <ATTRIBUTE-DEFINITION-XHTML-REF>_cf7-0MkhEee8KsfWrp9EJQ</ATTRIBUTE-DEFINITION-XHTML-REF>
              </DEFINITION>
              <THE-VALUE>
                <xhtml:div>Requirement-1</xhtml:div>
              </THE-VALUE>
            </ATTRIBUTE-VALUE-XHTML>
            <ATTRIBUTE-VALUE-ENUMERATION>
              <VALUES>
                <ENUM-VALUE-REF>_gXZ9oMk3Eee5A_N9aQFa1w</ENUM-VALUE-REF>
              </VALUES>
              <DEFINITION>
                <ATTRIBUTE-DEFINITION-ENUMERATION-REF>_0VDX0Mk3Eee5A_N9aQFa1w</ATTRIBUTE-DEFINITION-ENUMERATION-REF>
              </DEFINITION>
            </ATTRIBUTE-VALUE-ENUMERATION>
          </VALUES>
          <TYPE>
            <SPEC-OBJECT-TYPE-REF>_LaeSsMkhEee8KsfWrp9EJQ</SPEC-OBJECT-TYPE-REF>
          </TYPE>
        </SPEC-OBJECT>
      </SPEC-OBJECTS>
      <SPEC-RELATIONS/>
      <SPECIFICATIONS>
        <SPECIFICATION
            IDENTIFIER="_dESzoMkiEee8KsfWrp9EJQ"
            LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
            LONG-NAME="MODULE-1">
          <VALUES/>
          <TYPE>
            <SPECIFICATION-TYPE-REF>_NAXScMkhEee8KsfWrp9EJQ</SPECIFICATION-TYPE-REF>
          </TYPE>
          <CHILDREN>
            <SPEC-HIERARCHY
                IDENTIFIER="_Ob6YgMkjEee8KsfWrp9EJQ"
                LAST-CHANGE="2017-11-14T15:44:26.000+02:00"
                LONG-NAME="Requirement-1"
                IS-EDITABLE="true">
              <OBJECT>
                <SPEC-OBJECT-REF>_xen_QMkhEee8KsfWrp9EJQ</SPEC-OBJECT-REF>
              </OBJECT>
            </SPEC-HIERARCHY>
          </CHILDREN>
        </SPECIFICATION>
      </SPECIFICATIONS>
      <SPEC-RELATION-GROUPS/>
    </REQ-IF-CONTENT>
  </CORE-CONTENT>
  <TOOL-EXTENSIONS/>
</REQ-IF>
`

	// Setup of pointers
	// setup of FileToUpload instances pointers
}
