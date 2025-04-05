package models

var GeneratedModelFiles = []string{

	"../..embed.go",
	"../embed.go",
	string(GeneratedGongGoFilePath),
	string(GeneratedGongCoderGoFilePath),
	string(GeneratedGongAstGoFilePath),
	string(GeneratedGongSerializeGoFilePath),
	string(GeneratedGongSlicesGoFilePath),
	string(GeneratedGongMarshallGoFilePath),
	string(GeneratedGongGraphGoFilePath),
	string(GeneratedGongEnumGoFilePath),
	string(GeneratedGongCallbacksGoFilePath),
	string(GeneratedGongOrchestratorGoFilePath),
	string(GeneratedGongWopGoFilePath),
}

type GeneratedGoFilePath string

const (
	GeneratedGongGoFilePath             GeneratedGoFilePath = "gong.go"
	GeneratedGongEnumGoFilePath         GeneratedGoFilePath = "gong_enum.go"
	GeneratedGongMarshallGoFilePath     GeneratedGoFilePath = "gong_marshall.go"
	GeneratedGongGraphGoFilePath        GeneratedGoFilePath = "gong_graph.go"
	GeneratedGongSlicesGoFilePath       GeneratedGoFilePath = "gong_slices.go"
	GeneratedGongCoderGoFilePath        GeneratedGoFilePath = "gong_coder.go"
	GeneratedGongAstGoFilePath          GeneratedGoFilePath = "gong_ast.go"
	GeneratedGongCallbacksGoFilePath    GeneratedGoFilePath = "gong_callbacks.go"
	GeneratedGongOrchestratorGoFilePath GeneratedGoFilePath = "gong_orchestrator.go"
	GeneratedGongSerializeGoFilePath    GeneratedGoFilePath = "gong_serialize.go"
	GeneratedGongWopGoFilePath          GeneratedGoFilePath = "gong_wop.go"
	StagerGoFilePath                    GeneratedGoFilePath = "stager.go"
	DocsGoFilePath                      GeneratedGoFilePath = "docs.go"
)
