package icons

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

const titleOffSVG = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg
   height="24px"
   viewBox="0 -960 960 960"
   width="24px"
   fill="gray"
   version="1.1"
   id="svg1"
   sodipodi:docname="title_off_24dp_E3E3E3_FILL0_wght400_GRAD0_opsz24.svg"
   inkscape:version="1.3.2 (091e20e, 2023-11-25)"
   xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
   xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
   xmlns:xlink="http://www.w3.org/1999/xlink"
   xmlns="http://www.w3.org/2000/svg"
   xmlns:svg="http://www.w3.org/2000/svg">
  <defs
     id="defs1">
    <linearGradient
       id="swatch1"
       inkscape:swatch="solid">
      <stop
         style="stop-color:gray;stop-opacity:1;"
         offset="0"
         id="stop1" />
    </linearGradient>
    <linearGradient
       inkscape:collect="always"
       xlink:href="#swatch1"
       id="linearGradient1"
       x1="216.27119"
       y1="-549.23831"
       x2="776.27119"
       y2="-549.23831"
       gradientUnits="userSpaceOnUse" />
  </defs>
  <sodipodi:namedview
     id="namedview1"
     pagecolor="gray"
     bordercolor="#666666"
     borderopacity="1.0"
     inkscape:showpageshadow="2"
     inkscape:pageopacity="0.0"
     inkscape:pagecheckerboard="0"
     inkscape:deskcolor="#d1d1d1"
     inkscape:zoom="9.8333333"
     inkscape:cx="12"
     inkscape:cy="11.949153"
     inkscape:window-width="1792"
     inkscape:window-height="1120"
     inkscape:window-x="0"
     inkscape:window-y="0"
     inkscape:window-maximized="0"
     inkscape:current-layer="svg1" />
  <path
     d="m 436.27119,-229.23831 v -520 h -220 v -120 h 560 v 120 h -220 v 520 z"
     id="path1"
     style="stroke:url(#linearGradient1)" />
  <line
     x1="56.673222"
     y1="-843.8067"
     x2="847.91577"
     y2="-135.01297"
     stroke="#1f1f1f"
     stroke-width="40"
     id="line1"
     style="stroke-width:74.8884;stroke-dasharray:none;stroke:#e3e3e3;stroke-opacity:1" />
  <line
     x1="112.08426"
     y1="-893.89038"
     x2="903.32678"
     y2="-185.09668"
     stroke="#1f1f1f"
     stroke-width="40"
     id="line1-0"
     style="fill:#000000;fill-opacity:1;stroke:gray;stroke-width:74.8884;stroke-dasharray:none;stroke-opacity:1" />
</svg>
`

var SvgIconTitleOff = &tree.SVGIcon{
	Name: "T letter with a diagonal line",
	SVG:  titleOffSVG,
}

const titleSVG = `<svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="gray"><path d="M420-160v-520H200v-120h560v120H540v520H420Z"/></svg>`

var SvgIconTitle = &tree.SVGIcon{
	Name: "T letter",
	SVG:  titleSVG,
}
