// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for GeneratedGoFilePath
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (generatedgofilepath GeneratedGoFilePath) ToString() (res string) {

	// migration of former implementation of enum
	switch generatedgofilepath {
	// insertion code per enum code
	case LegacyGeneratedGongGoFilePath:
		res = "gong.go"
	case LegacyGeneratedGongEnumGoFilePath:
		res = "gong_enum.go"
	case LegacyGeneratedGongMarshallGoFilePath:
		res = "gong_marshall.go"
	case LegacyGeneratedGongGraphGoFilePath:
		res = "gong_graph.go"
	case LegacyGeneratedGongSlicesGoFilePath:
		res = "gong_slices.go"
	case LegacyGeneratedGongCoderGoFilePath:
		res = "gong_coder.go"
	case LegacyGeneratedGongAstGoFilePath:
		res = "gong_ast.go"
	case LegacyGeneratedGongCallbacksGoFilePath:
		res = "gong_callbacks.go"
	case LegacyGeneratedGongOrchestratorGoFilePath:
		res = "gong_orchestrator.go"
	case LegacyGeneratedGongReverseGoFilePath:
		res = "gong_reverse.go"
	case LegacyGeneratedGongSerializeGoFilePath:
		res = "gong_serialize.go"
	case LegacyGeneratedGongWopGoFilePath:
		res = "gong_wop.go"
	case GeneratedGongGoFilePath:
		res = "zzz_gong.go"
	case GeneratedGongEnumGoFilePath:
		res = "zzz_gong_enum.go"
	case GeneratedGongMarshallGoFilePath:
		res = "zzz_gong_marshall.go"
	case GeneratedGongGraphGoFilePath:
		res = "zzz_gong_graph.go"
	case GeneratedGongSlicesGoFilePath:
		res = "zzz_gong_slices.go"
	case GeneratedGongCoderGoFilePath:
		res = "zzz_gong_coder.go"
	case GeneratedGongAstGoFilePath:
		res = "zzz_gong_ast.go"
	case GeneratedGongCallbacksGoFilePath:
		res = "zzz_gong_callbacks.go"
	case GeneratedGongOrchestratorGoFilePath:
		res = "zzz_gong_orchestrator.go"
	case GeneratedGongReverseGoFilePath:
		res = "zzz_gong_reverse.go"
	case GeneratedGongSerializeGoFilePath:
		res = "zzz_gong_serialize.go"
	case GeneratedGongWopGoFilePath:
		res = "zzz_gong_wop.go"
	case StagerGoFilePath:
		res = "stager.go"
	case DocsGoFilePath:
		res = "docs.go"
	}
	return
}

func (generatedgofilepath *GeneratedGoFilePath) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "gong.go":
		*generatedgofilepath = LegacyGeneratedGongGoFilePath
		return
	case "gong_enum.go":
		*generatedgofilepath = LegacyGeneratedGongEnumGoFilePath
		return
	case "gong_marshall.go":
		*generatedgofilepath = LegacyGeneratedGongMarshallGoFilePath
		return
	case "gong_graph.go":
		*generatedgofilepath = LegacyGeneratedGongGraphGoFilePath
		return
	case "gong_slices.go":
		*generatedgofilepath = LegacyGeneratedGongSlicesGoFilePath
		return
	case "gong_coder.go":
		*generatedgofilepath = LegacyGeneratedGongCoderGoFilePath
		return
	case "gong_ast.go":
		*generatedgofilepath = LegacyGeneratedGongAstGoFilePath
		return
	case "gong_callbacks.go":
		*generatedgofilepath = LegacyGeneratedGongCallbacksGoFilePath
		return
	case "gong_orchestrator.go":
		*generatedgofilepath = LegacyGeneratedGongOrchestratorGoFilePath
		return
	case "gong_reverse.go":
		*generatedgofilepath = LegacyGeneratedGongReverseGoFilePath
		return
	case "gong_serialize.go":
		*generatedgofilepath = LegacyGeneratedGongSerializeGoFilePath
		return
	case "gong_wop.go":
		*generatedgofilepath = LegacyGeneratedGongWopGoFilePath
		return
	case "zzz_gong.go":
		*generatedgofilepath = GeneratedGongGoFilePath
		return
	case "zzz_gong_enum.go":
		*generatedgofilepath = GeneratedGongEnumGoFilePath
		return
	case "zzz_gong_marshall.go":
		*generatedgofilepath = GeneratedGongMarshallGoFilePath
		return
	case "zzz_gong_graph.go":
		*generatedgofilepath = GeneratedGongGraphGoFilePath
		return
	case "zzz_gong_slices.go":
		*generatedgofilepath = GeneratedGongSlicesGoFilePath
		return
	case "zzz_gong_coder.go":
		*generatedgofilepath = GeneratedGongCoderGoFilePath
		return
	case "zzz_gong_ast.go":
		*generatedgofilepath = GeneratedGongAstGoFilePath
		return
	case "zzz_gong_callbacks.go":
		*generatedgofilepath = GeneratedGongCallbacksGoFilePath
		return
	case "zzz_gong_orchestrator.go":
		*generatedgofilepath = GeneratedGongOrchestratorGoFilePath
		return
	case "zzz_gong_reverse.go":
		*generatedgofilepath = GeneratedGongReverseGoFilePath
		return
	case "zzz_gong_serialize.go":
		*generatedgofilepath = GeneratedGongSerializeGoFilePath
		return
	case "zzz_gong_wop.go":
		*generatedgofilepath = GeneratedGongWopGoFilePath
		return
	case "stager.go":
		*generatedgofilepath = StagerGoFilePath
		return
	case "docs.go":
		*generatedgofilepath = DocsGoFilePath
		return
	default:
		return errUnkownEnum
	}
}

func (generatedgofilepath *GeneratedGoFilePath) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LegacyGeneratedGongGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongGoFilePath
	case "LegacyGeneratedGongEnumGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongEnumGoFilePath
	case "LegacyGeneratedGongMarshallGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongMarshallGoFilePath
	case "LegacyGeneratedGongGraphGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongGraphGoFilePath
	case "LegacyGeneratedGongSlicesGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongSlicesGoFilePath
	case "LegacyGeneratedGongCoderGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongCoderGoFilePath
	case "LegacyGeneratedGongAstGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongAstGoFilePath
	case "LegacyGeneratedGongCallbacksGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongCallbacksGoFilePath
	case "LegacyGeneratedGongOrchestratorGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongOrchestratorGoFilePath
	case "LegacyGeneratedGongReverseGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongReverseGoFilePath
	case "LegacyGeneratedGongSerializeGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongSerializeGoFilePath
	case "LegacyGeneratedGongWopGoFilePath":
		*generatedgofilepath = LegacyGeneratedGongWopGoFilePath
	case "GeneratedGongGoFilePath":
		*generatedgofilepath = GeneratedGongGoFilePath
	case "GeneratedGongEnumGoFilePath":
		*generatedgofilepath = GeneratedGongEnumGoFilePath
	case "GeneratedGongMarshallGoFilePath":
		*generatedgofilepath = GeneratedGongMarshallGoFilePath
	case "GeneratedGongGraphGoFilePath":
		*generatedgofilepath = GeneratedGongGraphGoFilePath
	case "GeneratedGongSlicesGoFilePath":
		*generatedgofilepath = GeneratedGongSlicesGoFilePath
	case "GeneratedGongCoderGoFilePath":
		*generatedgofilepath = GeneratedGongCoderGoFilePath
	case "GeneratedGongAstGoFilePath":
		*generatedgofilepath = GeneratedGongAstGoFilePath
	case "GeneratedGongCallbacksGoFilePath":
		*generatedgofilepath = GeneratedGongCallbacksGoFilePath
	case "GeneratedGongOrchestratorGoFilePath":
		*generatedgofilepath = GeneratedGongOrchestratorGoFilePath
	case "GeneratedGongReverseGoFilePath":
		*generatedgofilepath = GeneratedGongReverseGoFilePath
	case "GeneratedGongSerializeGoFilePath":
		*generatedgofilepath = GeneratedGongSerializeGoFilePath
	case "GeneratedGongWopGoFilePath":
		*generatedgofilepath = GeneratedGongWopGoFilePath
	case "StagerGoFilePath":
		*generatedgofilepath = StagerGoFilePath
	case "DocsGoFilePath":
		*generatedgofilepath = DocsGoFilePath
	default:
		err = errUnkownEnum
	}
	return
}

func (generatedgofilepath *GeneratedGoFilePath) ToCodeString() (res string) {

	switch *generatedgofilepath {
	// insertion code per enum code
	case LegacyGeneratedGongGoFilePath:
		res = "LegacyGeneratedGongGoFilePath"
	case LegacyGeneratedGongEnumGoFilePath:
		res = "LegacyGeneratedGongEnumGoFilePath"
	case LegacyGeneratedGongMarshallGoFilePath:
		res = "LegacyGeneratedGongMarshallGoFilePath"
	case LegacyGeneratedGongGraphGoFilePath:
		res = "LegacyGeneratedGongGraphGoFilePath"
	case LegacyGeneratedGongSlicesGoFilePath:
		res = "LegacyGeneratedGongSlicesGoFilePath"
	case LegacyGeneratedGongCoderGoFilePath:
		res = "LegacyGeneratedGongCoderGoFilePath"
	case LegacyGeneratedGongAstGoFilePath:
		res = "LegacyGeneratedGongAstGoFilePath"
	case LegacyGeneratedGongCallbacksGoFilePath:
		res = "LegacyGeneratedGongCallbacksGoFilePath"
	case LegacyGeneratedGongOrchestratorGoFilePath:
		res = "LegacyGeneratedGongOrchestratorGoFilePath"
	case LegacyGeneratedGongReverseGoFilePath:
		res = "LegacyGeneratedGongReverseGoFilePath"
	case LegacyGeneratedGongSerializeGoFilePath:
		res = "LegacyGeneratedGongSerializeGoFilePath"
	case LegacyGeneratedGongWopGoFilePath:
		res = "LegacyGeneratedGongWopGoFilePath"
	case GeneratedGongGoFilePath:
		res = "GeneratedGongGoFilePath"
	case GeneratedGongEnumGoFilePath:
		res = "GeneratedGongEnumGoFilePath"
	case GeneratedGongMarshallGoFilePath:
		res = "GeneratedGongMarshallGoFilePath"
	case GeneratedGongGraphGoFilePath:
		res = "GeneratedGongGraphGoFilePath"
	case GeneratedGongSlicesGoFilePath:
		res = "GeneratedGongSlicesGoFilePath"
	case GeneratedGongCoderGoFilePath:
		res = "GeneratedGongCoderGoFilePath"
	case GeneratedGongAstGoFilePath:
		res = "GeneratedGongAstGoFilePath"
	case GeneratedGongCallbacksGoFilePath:
		res = "GeneratedGongCallbacksGoFilePath"
	case GeneratedGongOrchestratorGoFilePath:
		res = "GeneratedGongOrchestratorGoFilePath"
	case GeneratedGongReverseGoFilePath:
		res = "GeneratedGongReverseGoFilePath"
	case GeneratedGongSerializeGoFilePath:
		res = "GeneratedGongSerializeGoFilePath"
	case GeneratedGongWopGoFilePath:
		res = "GeneratedGongWopGoFilePath"
	case StagerGoFilePath:
		res = "StagerGoFilePath"
	case DocsGoFilePath:
		res = "DocsGoFilePath"
	}
	return
}

func (generatedgofilepath GeneratedGoFilePath) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LegacyGeneratedGongGoFilePath")
	res = append(res, "LegacyGeneratedGongEnumGoFilePath")
	res = append(res, "LegacyGeneratedGongMarshallGoFilePath")
	res = append(res, "LegacyGeneratedGongGraphGoFilePath")
	res = append(res, "LegacyGeneratedGongSlicesGoFilePath")
	res = append(res, "LegacyGeneratedGongCoderGoFilePath")
	res = append(res, "LegacyGeneratedGongAstGoFilePath")
	res = append(res, "LegacyGeneratedGongCallbacksGoFilePath")
	res = append(res, "LegacyGeneratedGongOrchestratorGoFilePath")
	res = append(res, "LegacyGeneratedGongReverseGoFilePath")
	res = append(res, "LegacyGeneratedGongSerializeGoFilePath")
	res = append(res, "LegacyGeneratedGongWopGoFilePath")
	res = append(res, "GeneratedGongGoFilePath")
	res = append(res, "GeneratedGongEnumGoFilePath")
	res = append(res, "GeneratedGongMarshallGoFilePath")
	res = append(res, "GeneratedGongGraphGoFilePath")
	res = append(res, "GeneratedGongSlicesGoFilePath")
	res = append(res, "GeneratedGongCoderGoFilePath")
	res = append(res, "GeneratedGongAstGoFilePath")
	res = append(res, "GeneratedGongCallbacksGoFilePath")
	res = append(res, "GeneratedGongOrchestratorGoFilePath")
	res = append(res, "GeneratedGongReverseGoFilePath")
	res = append(res, "GeneratedGongSerializeGoFilePath")
	res = append(res, "GeneratedGongWopGoFilePath")
	res = append(res, "StagerGoFilePath")
	res = append(res, "DocsGoFilePath")

	return
}

func (generatedgofilepath GeneratedGoFilePath) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "gong.go")
	res = append(res, "gong_enum.go")
	res = append(res, "gong_marshall.go")
	res = append(res, "gong_graph.go")
	res = append(res, "gong_slices.go")
	res = append(res, "gong_coder.go")
	res = append(res, "gong_ast.go")
	res = append(res, "gong_callbacks.go")
	res = append(res, "gong_orchestrator.go")
	res = append(res, "gong_reverse.go")
	res = append(res, "gong_serialize.go")
	res = append(res, "gong_wop.go")
	res = append(res, "zzz_gong.go")
	res = append(res, "zzz_gong_enum.go")
	res = append(res, "zzz_gong_marshall.go")
	res = append(res, "zzz_gong_graph.go")
	res = append(res, "zzz_gong_slices.go")
	res = append(res, "zzz_gong_coder.go")
	res = append(res, "zzz_gong_ast.go")
	res = append(res, "zzz_gong_callbacks.go")
	res = append(res, "zzz_gong_orchestrator.go")
	res = append(res, "zzz_gong_reverse.go")
	res = append(res, "zzz_gong_serialize.go")
	res = append(res, "zzz_gong_wop.go")
	res = append(res, "stager.go")
	res = append(res, "docs.go")

	return
}

// Utility function for GongEnumType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (gongenumtype GongEnumType) ToInt() (res int) {

	// migration of former implementation of enum
	switch gongenumtype {
	// insertion code per enum code
	case Int:
		res = 0
	case String:
		res = 1
	}
	return
}

func (gongenumtype *GongEnumType) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*gongenumtype = Int
		return
	case 1:
		*gongenumtype = String
		return
	default:
		return errUnkownEnum
	}
}

func (gongenumtype *GongEnumType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Int":
		*gongenumtype = Int
	case "String":
		*gongenumtype = String
	default:
		err = errUnkownEnum
	}
	return
}

func (gongenumtype *GongEnumType) ToCodeString() (res string) {

	switch *gongenumtype {
	// insertion code per enum code
	case Int:
		res = "Int"
	case String:
		res = "String"
	}
	return
}

func (gongenumtype GongEnumType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Int")
	res = append(res, "String")

	return
}

func (gongenumtype GongEnumType) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int | GongEnumType
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*GongEnumType
	FromCodeString(input string) (err error)
}

// Last line of the template
