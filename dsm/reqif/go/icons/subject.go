package icons

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

const subjectSVG = `<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="gray"><path d="M160-200v-80h400v80H160Zm0-160v-80h640v80H160Zm0-160v-80h640v80H160Zm0-160v-80h640v80H160Z"/></svg>`

var SvgIconSubject = &tree.SVGIcon{
	Name: "Subject",
	SVG:  subjectSVG,
}

const subjectOffSVG = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg
   height="24px"
   viewBox="0 -960 960 960"
   width="24px"
   fill="#1f1f1f"
   version="1.1"
   id="svg1"
   sodipodi:docname="subject_off_rows_narrow_24dp_1F1F1F_FILL0_wght400_GRAD0_opsz24.svg"
   inkscape:version="1.3.2 (091e20e, 2023-11-25, custom)"
   xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
   xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
   xmlns="http://www.w3.org/2000/svg"
   xmlns:svg="http://www.w3.org/2000/svg">
  <defs
     id="defs1" />
  <sodipodi:namedview
     id="namedview1"
     pagecolor="#ffffff"
     bordercolor="#000000"
     borderopacity="0.25"
     inkscape:showpageshadow="2"
     inkscape:pageopacity="0.0"
     inkscape:pagecheckerboard="0"
     inkscape:deskcolor="#d1d1d1"
     inkscape:zoom="33.875"
     inkscape:cx="12"
     inkscape:cy="12"
     inkscape:window-width="1920"
     inkscape:window-height="1009"
     inkscape:window-x="-8"
     inkscape:window-y="148"
     inkscape:window-maximized="1"
     inkscape:current-layer="svg1" />
  <path
     d="M160-200v-80h400v80H160Zm0-160v-80h640v80H160Zm0-160v-80h640v80H160Zm0-160v-80h640v80H160Z"
     id="path1"
     style="stroke-width:0.999997;fill:gray;fill-opacity:1" />
  <line
     x1="28.378874"
     y1="-812.10974"
     x2="819.6214"
     y2="-103.31602"
     stroke="#1f1f1f"
     stroke-width="40"
     id="line1"
     style="stroke:gray;stroke-width:74.8884;stroke-dasharray:none;stroke-opacity:1" />
</svg>
`

var SvgIconSubjectOff = &tree.SVGIcon{
	Name: "Subject with a diagonal line",
	SVG:  subjectOffSVG,
}
