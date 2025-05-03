package models

const SVGName string = `SVG`

func (stager *Stager) UpdateAndCommitSVGStage() {

	stager.UpdateSVGStage()

	stager.svgStage.Commit()
}

func (stager *Stager) UpdateSVGStage() {
	stager.svgStage.Reset()

	docSVG := NewDocSVGMapper(stager.svgStage)
	docSVG.GenerateSvg(stager.stage)

}
