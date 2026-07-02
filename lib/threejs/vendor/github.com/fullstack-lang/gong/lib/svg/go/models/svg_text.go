package models

type SvgText struct {
	Name string

	//gong:text
	//gong:width 600 gong:height 300
	Text string

	Proxy SvgTextProxyInterface
}

type SvgTextProxyInterface interface {
	Updated()
}

func (svgtext *SvgText) OnAfterUpdate(
	stage *Stage,
	stageSvgText, frontSvgText *SvgText) {

	if svgtext.Proxy != nil {
		svgtext.Proxy.Updated()
	}
}
