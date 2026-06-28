package icons

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

const tableSVG = `<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="gray"><path d="M760-360v-80H200v80h560Zm0-160v-80H200v80h560Zm0-160v-80H200v80h560ZM200-120q-33 0-56.5-23.5T120-200v-560q0-33 23.5-56.5T200-840h560q33 0 56.5 23.5T840-760v560q0 33-23.5 56.5T760-120H200Zm560-80v-80H200v80h560Z"/></svg>`

var SvgIconTable = &tree.SVGIcon{
	Name: "Table",
	SVG:  tableSVG,
}

const tableOffSVG = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg
   height="24px"
   viewBox="0 -960 960 960"
   width="24px"
   fill="#1f1f1f"
   version="1.1"
   id="svg1"
   sodipodi:docname="table_off_rows_narrow_24dp_1F1F1F_FILL0_wght400_GRAD0_opsz24.svg"
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
     d="m 702.27142,-413.53013 v -71.07831 H 192.78393 v 71.07831 z m 0,-142.15662 v -71.07831 H 192.78393 v 71.07831 z m 0,-142.15663 v -71.07831 H 192.78393 v 71.07831 z M 192.78393,-200.2952 q -30.02337,0 -51.40365,-20.87926 Q 120,-242.05371 120,-271.37351 v -497.54818 q 0,-29.3198 21.38028,-50.19905 Q 162.76056,-840 192.78393,-840 h 509.48749 q 30.02337,0 51.40365,20.87926 21.38028,20.87925 21.38028,50.19905 v 497.54818 q 0,29.3198 -21.38028,50.19905 -21.38028,20.87926 -51.40365,20.87926 z m 509.48749,-71.07831 v -71.07831 H 192.78393 v 71.07831 z"
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

var SvgIconTableOff = &tree.SVGIcon{
	Name: "Table with a diagonal line",
	SVG:  tableOffSVG,
}
