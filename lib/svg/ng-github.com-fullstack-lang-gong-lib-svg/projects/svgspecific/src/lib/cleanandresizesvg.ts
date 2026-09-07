export function processSVG(svgString: string): string {
  const parser = new DOMParser();
  const doc = parser.parseFromString(svgString, "image/svg+xml");
  const svg = doc.documentElement;

  // Find all rect elements
  const rects = svg.getElementsByTagName('rect');

  // Remove the first rect if it exists
  if (rects.length > 0) {
    rects[0].remove();
  }

  // Remove interactive helper elements that should not appear in exported SVGs
  // (e.g. hit areas for mouse dragging, overlay handles, and control points)
  const hitAreas = Array.from(svg.querySelectorAll('.hit-area'));
  hitAreas.forEach(el => el.remove());

  const overlays = Array.from(svg.querySelectorAll('.overlay'));
  overlays.forEach(el => el.remove());

  // Remove invisible lines (in SVG a line without a visible stroke has no visual appearance,
  // but PowerPoint/Office converts lines with stroke="none" or stroke="transparent" into default solid black lines)
  const lines = Array.from(svg.querySelectorAll('line'));
  for (const line of lines) {
    const stroke = (line.getAttribute('stroke') || '').trim().toLowerCase();
    const strokeWidth = parseFloat(line.getAttribute('stroke-width') || '1');
    const strokeOpacity = line.hasAttribute('stroke-opacity') ? parseFloat(line.getAttribute('stroke-opacity') || '1') : 1;
    if (stroke === 'none' || stroke === 'transparent' || stroke === '' || strokeWidth === 0 || strokeOpacity === 0) {
      line.remove();
    }
  }

  // Remove empty paths that have no data
  const paths = Array.from(svg.querySelectorAll('path'));
  for (const path of paths) {
    const d = (path.getAttribute('d') || '').trim();
    if (!d) {
      path.remove();
    }
  }

  // Remove empty groups
  const groups = Array.from(svg.querySelectorAll('g'));
  for (const g of groups) {
    if (g.children.length === 0) {
      g.remove();
    }
  }

  // Get all remaining elements and convert to array for iteration
  const elements = Array.from(svg.getElementsByTagName('*'));

  // Initialize boundaries
  let minX = Infinity;
  let minY = Infinity;
  let maxX = -Infinity;
  let maxY = -Infinity;

  // Process each element to find boundaries
  for (const element of elements) {
    // Special handling for rect elements
    if (element.tagName.toLowerCase() === 'rect') {
      const x = parseFloat(element.getAttribute('x') || '0');
      const y = parseFloat(element.getAttribute('y') || '0');
      const width = parseFloat(element.getAttribute('width') || '0');
      const height = parseFloat(element.getAttribute('height') || '0');
      const rx = parseFloat(element.getAttribute('rx') || '0');

      // Consider transform attributes
      const transform = element.getAttribute('transform');
      let tx = 0, ty = 0;

      if (transform) {
        const match = transform.match(/translate\(([-\d.]+)\s*([-\d.]+)?\)/);
        if (match) {
          tx = parseFloat(match[1]) || 0;
          ty = parseFloat(match[2]) || 0;
        }
      }

      minX = Math.min(minX, x + tx);
      minY = Math.min(minY, y + ty);
      maxX = Math.max(maxX, x + width + tx);
      maxY = Math.max(maxY, y + height + ty);
      continue;
    }

    // Handle other SVG elements
    try {
      const bbox = (element as SVGGraphicsElement).getBBox?.();
      if (bbox) {
        // Consider transform attributes
        const transform = element.getAttribute('transform');
        let tx = 0, ty = 0;

        if (transform) {
          const match = transform.match(/translate\(([-\d.]+)\s*([-\d.]+)?\)/);
          if (match) {
            tx = parseFloat(match[1]) || 0;
            ty = parseFloat(match[2]) || 0;
          }
        }

        minX = Math.min(minX, bbox.x + tx);
        minY = Math.min(minY, bbox.y + ty);
        maxX = Math.max(maxX, bbox.x + bbox.width + tx);
        maxY = Math.max(maxY, bbox.y + bbox.height + ty);
      }
    } catch (e) {
      // Some elements might not support getBBox
    }

    // Handle explicit coordinates for elements like lines
    ['x', 'x1', 'x2', 'cx'].forEach(attr => {
      const val = parseFloat(element.getAttribute(attr) || '');
      if (!isNaN(val)) {
        minX = Math.min(minX, val);
        maxX = Math.max(maxX, val);
      }
    });

    ['y', 'y1', 'y2', 'cy'].forEach(attr => {
      const val = parseFloat(element.getAttribute(attr) || '');
      if (!isNaN(val)) {
        minY = Math.min(minY, val);
        maxY = Math.max(maxY, val);
      }
    });

    // Handle path elements
    if (element.tagName.toLowerCase() === 'path') {
      const d = element.getAttribute('d');
      if (d) {
        // Split path into commands
        const commands = d.match(/[MmLlHhVvCcSsQqTtAaZz][^MmLlHhVvCcSsQqTtAaZz]*/g) || [];
        for (const cmd of commands) {
          // Extract numbers from command
          const numbers = cmd.slice(1).trim().split(/[\s,]+/).map(parseFloat);
          for (let i = 0; i < numbers.length; i += 2) {
            if (!isNaN(numbers[i])) {
              minX = Math.min(minX, numbers[i]);
              maxX = Math.max(maxX, numbers[i]);
            }
            if (!isNaN(numbers[i + 1])) {
              minY = Math.min(minY, numbers[i + 1]);
              maxY = Math.max(maxY, numbers[i + 1]);
            }
          }
        }
      }
    }
  }

  // Add padding
  const padding = 10;
  minX -= padding;
  minY -= padding;
  maxX += padding;
  maxY += padding;

  // Calculate new dimensions
  const width = maxX - minX;
  const height = maxY - minY;

  // Update SVG attributes
  svg.setAttribute('width', width.toString());
  svg.setAttribute('height', height.toString());
  svg.setAttribute('viewBox', `${minX} ${minY} ${width} ${height}`);

  // Ensure font-family is declared on the SVG root for PowerPoint compatibility
  if (!svg.getAttribute('font-family')) {
    svg.setAttribute('font-family', 'Roboto, Arial, sans-serif');
  }

  // Ensure all text elements have valid font-size and remove empty attributes
  const textElements = Array.from(svg.querySelectorAll('text'));
  for (const textEl of textElements) {
    const fontSize = textEl.getAttribute('font-size');
    if (!fontSize || fontSize === '') {
      textEl.setAttribute('font-size', '16px');
    }
    ['font-family', 'font-weight', 'font-style', 'letter-spacing', 'white-space', 'transform', 'stroke-dasharray', 'writing-mode', 'dominant-baseline'].forEach(attr => {
      if (textEl.getAttribute(attr) === '') {
        textEl.removeAttribute(attr);
      }
    });
  }

  // Normalize colors and clean attributes across all SVG elements for PowerPoint compatibility
  const allElements = Array.from(svg.querySelectorAll('*'));
  for (const el of allElements) {
    // Normalize stroke
    if (el.hasAttribute('stroke')) {
      const stroke = el.getAttribute('stroke') || '';
      const trimmed = stroke.trim();
      if (trimmed === '') {
        el.removeAttribute('stroke');
      } else if (trimmed.toLowerCase() === 'transparent') {
        el.setAttribute('stroke', 'none');
      } else {
        el.setAttribute('stroke', normalizeColor(trimmed));
      }
    }

    // Normalize fill
    if (el.hasAttribute('fill')) {
      const fill = el.getAttribute('fill') || '';
      const trimmed = fill.trim();
      if (trimmed === '') {
        el.removeAttribute('fill');
      } else if (trimmed.toLowerCase() === 'transparent') {
        el.setAttribute('fill', 'none');
      } else {
        el.setAttribute('fill', normalizeColor(trimmed));
      }
    }

    // Normalize opacity if specified as integer > 1 (e.g. 100 -> 1)
    ['fill-opacity', 'stroke-opacity', 'opacity'].forEach(attr => {
      if (el.hasAttribute(attr)) {
        const val = parseFloat(el.getAttribute(attr) || '1');
        if (!isNaN(val) && val > 1) {
          el.setAttribute(attr, (val / 100).toString());
        }
      }
    });

    // Remove empty attributes that can confuse strict parsers
    ['stroke-dasharray', 'transform'].forEach(attr => {
      if (el.getAttribute(attr) === '') {
        el.removeAttribute(attr);
      }
    });
  }

  // Convert back to string
  return new XMLSerializer().serializeToString(svg);
}

const CSS_COLOR_NAMES: Record<string, string> = {
  aliceblue: '#f0f8ff',
  antiquewhite: '#faebd7',
  aqua: '#00ffff',
  aquamarine: '#7fffd4',
  azure: '#f0ffff',
  beige: '#f5f5dc',
  bisque: '#ffe4c4',
  black: '#000000',
  blanchedalmond: '#ffebcd',
  blue: '#0000ff',
  blueviolet: '#8a2be2',
  brown: '#a52a2a',
  burlywood: '#deb887',
  cadetblue: '#5f9ea0',
  chartreuse: '#7fff00',
  chocolate: '#d2691e',
  coral: '#ff7f50',
  cornflowerblue: '#6495ed',
  cornsilk: '#fff8dc',
  crimson: '#dc143c',
  cyan: '#00ffff',
  darkblue: '#00008b',
  darkcyan: '#008b8b',
  darkgoldenrod: '#b8860b',
  darkgray: '#a9a9a9',
  darkgreen: '#006400',
  darkgrey: '#a9a9a9',
  darkkhaki: '#bdb76b',
  darkmagenta: '#8b008b',
  darkolivegreen: '#556b2f',
  darkorange: '#ff8c00',
  darkorchid: '#9932cc',
  darkred: '#8b0000',
  darksalmon: '#e9967a',
  darkseagreen: '#8fbc8f',
  darkslateblue: '#483d8b',
  darkslategray: '#2f4f4f',
  darkslategrey: '#2f4f4f',
  darkturquoise: '#00ced1',
  darkviolet: '#9400d3',
  deeppink: '#ff1493',
  deepskyblue: '#00bfff',
  dimgray: '#696969',
  dimgrey: '#696969',
  dodgerblue: '#1e90ff',
  firebrick: '#b22222',
  floralwhite: '#fffaf0',
  forestgreen: '#228b22',
  fuchsia: '#ff00ff',
  gainsboro: '#dcdcdc',
  ghostwhite: '#f8f8ff',
  gold: '#ffd700',
  goldenrod: '#daa520',
  gray: '#808080',
  green: '#008000',
  greenyellow: '#adff2f',
  grey: '#808080',
  honeydew: '#f0fff0',
  hotpink: '#ff69b4',
  indianred: '#cd5c5c',
  indigo: '#4b0082',
  ivory: '#fffff0',
  khaki: '#f0e68c',
  lavender: '#e6e6fa',
  lavenderblush: '#fff0f5',
  lawngreen: '#7cfc00',
  lemonchiffon: '#fffacd',
  lightblue: '#add8e6',
  lightcoral: '#f08080',
  lightcyan: '#e0ffff',
  lightgoldenrodyellow: '#fafad2',
  lightgray: '#d3d3d3',
  lightgreen: '#90ee90',
  lightgrey: '#d3d3d3',
  lightpink: '#ffb6c1',
  lightsalmon: '#ffa07a',
  lightseagreen: '#20b2aa',
  lightskyblue: '#87cefa',
  lightslategray: '#778899',
  lightslategrey: '#778899',
  lightsteelblue: '#b0c4de',
  lightyellow: '#ffffe0',
  lime: '#00ff00',
  limegreen: '#32cd32',
  linen: '#faf0e6',
  magenta: '#ff00ff',
  maroon: '#800000',
  mediumaquamarine: '#66cdaa',
  mediumblue: '#0000cd',
  mediumorchid: '#ba55d3',
  mediumpurple: '#9370db',
  mediumseagreen: '#3cb371',
  mediumslateblue: '#7b68ee',
  mediumspringgreen: '#00fa9a',
  mediumturquoise: '#48d1cc',
  mediumvioletred: '#c71585',
  midnightblue: '#191970',
  mintcream: '#f5fffa',
  mistyrose: '#ffe4e1',
  moccasin: '#ffe4b5',
  navajowhite: '#ffdead',
  navy: '#000080',
  oldlace: '#fdf5e6',
  olive: '#808000',
  olivedrab: '#6b8e23',
  orange: '#ffa500',
  orangered: '#ff4500',
  orchid: '#da70d6',
  palegoldenrod: '#eee8aa',
  palegreen: '#98fb98',
  paleturquoise: '#afeeee',
  palevioletred: '#db7093',
  papayawhip: '#ffefd5',
  peachpuff: '#ffdab9',
  peru: '#cd853f',
  pink: '#ffc0cb',
  plum: '#dda0dd',
  powderblue: '#b0e0e6',
  purple: '#800080',
  red: '#ff0000',
  rosybrown: '#bc8f8f',
  royalblue: '#4169e1',
  saddlebrown: '#8b4513',
  salmon: '#fa8072',
  sandybrown: '#f4a460',
  seagreen: '#2e8b57',
  seashell: '#fff5ee',
  sienna: '#a0522d',
  silver: '#c0c0c0',
  skyblue: '#87ceeb',
  slateblue: '#6a5acd',
  slategray: '#708090',
  slategrey: '#708090',
  snow: '#fffafa',
  springgreen: '#00ff7f',
  steelblue: '#4682b4',
  tan: '#d2b48c',
  teal: '#008080',
  thistle: '#d8bfd8',
  tomato: '#ff6347',
  turquoise: '#40e0d0',
  violet: '#ee82ee',
  wheat: '#f5deb3',
  white: '#ffffff',
  whitesmoke: '#f5f5f5',
  yellow: '#ffff00',
  yellowgreen: '#9acd32',
};

function normalizeColor(color: string): string {
  const trimmed = color.trim().toLowerCase();
  if (CSS_COLOR_NAMES[trimmed]) {
    return CSS_COLOR_NAMES[trimmed];
  }
  return color.trim();
}

/**
* A simple pretty-print/format function for SVG strings.
*
* @param svgString The raw SVG markup as a string
* @returns A formatted version of the SVG with indentation and line breaks
*/
export function formatSVG(svgString: string): string {
  // Remove any existing newlines or extra spaces between tags
  // so we start from a predictable single-line "minified" form.
  let minified = svgString
    // Remove line breaks
    .replace(/\r?\n|\r/g, '')
    // Remove spaces between tags
    .replace(/>\s+</g, '><')
    // Trim start/end
    .trim();

  // We'll split on the boundaries between tags: ">"
  // and the start of the next tag "<".
  // Then we insert newlines and indentation accordingly.
  const tokens = minified.split(/></);

  let formatted = '';
  let indentLevel = 0;
  const indentSize = 2; // Adjust spaces per indent if desired
  const indentChar = ' ';

  for (let i = 0; i < tokens.length; i++) {
    let token = tokens[i];

    // If this is not the first token, prepend "<" back
    // because we split by "><".
    if (i > 0) {
      token = '<' + token;
    }
    // If this is not the last token, append ">" back
    // because we split by "><".
    if (i < tokens.length - 1) {
      token += '>';
    }

    // Check if it's a closing tag
    if (token.match(/^<\/\w/)) {
      // Decrease indent level for closing tag
      indentLevel = Math.max(indentLevel - 1, 0);
    }

    // Create the current indentation string
    const currentIndent = indentChar.repeat(indentLevel * indentSize);

    // Add the token with indentation
    formatted += currentIndent + token + '\n';

    // Check if it's an opening tag (and not a self-closing tag like <rect ... />)
    // If so, increase indent level
    if (
      token.match(/^<\w[^>]*[^/]>$/) && // starts with <tag ...> and doesn't end with />
      !token.match(/<\/\w/) // not a closing tag
    ) {
      indentLevel++;
    }
  }

  // Trim extra whitespace at the end
  return formatted.trim();
}
