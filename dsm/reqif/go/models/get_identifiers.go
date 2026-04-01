package models

func (d *DATATYPE_DEFINITION_XHTML) GetIdentifier() string       { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_STRING) GetIdentifier() string      { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_BOOLEAN) GetIdentifier() string     { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_INTEGER) GetIdentifier() string     { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_REAL) GetIdentifier() string        { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_DATE) GetIdentifier() string        { return d.IDENTIFIER }
func (d *DATATYPE_DEFINITION_ENUMERATION) GetIdentifier() string { return d.IDENTIFIER }

func (s *SPEC_OBJECT_TYPE) GetIdentifier() string   { return s.IDENTIFIER }
func (s *SPECIFICATION_TYPE) GetIdentifier() string { return s.IDENTIFIER }
func (s *SPEC_OBJECT) GetIdentifier() string        { return s.IDENTIFIER }

func (a *ATTRIBUTE_DEFINITION_XHTML) GetIdentifier() string       { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_STRING) GetIdentifier() string      { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_BOOLEAN) GetIdentifier() string     { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_INTEGER) GetIdentifier() string     { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_DATE) GetIdentifier() string        { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_REAL) GetIdentifier() string        { return a.IDENTIFIER }
func (a *ATTRIBUTE_DEFINITION_ENUMERATION) GetIdentifier() string { return a.IDENTIFIER }

func (e *ENUM_VALUE) GetIdentifier() string         { return e.IDENTIFIER }
func (s *SPEC_RELATION_TYPE) GetIdentifier() string { return s.IDENTIFIER }
func (s *SPECIFICATION) GetIdentifier() string      { return s.IDENTIFIER }
