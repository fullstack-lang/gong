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

  // Reset pan/zoom transform on top-level <g> so content coordinates are 1:1 in SVG space
  const topGroups = Array.from(svg.querySelectorAll('g'));
  for (const g of topGroups) {
    if (g.parentElement === svg && g.hasAttribute('transform')) {
      g.removeAttribute('transform');
    }
  }

  // Convert emojis (such as note memo 📝 or link 🔗) into clean vector SVG icons for PowerPoint compatibility.
  // PowerPoint and Microsoft Office SVG renderers do not support color emoji glyphs in SVG text,
  // causing missing glyphs to render as solid black rectangles (■).
  const allTextElements = Array.from(svg.querySelectorAll('text'));
  for (const textEl of allTextElements) {
    const tspans = Array.from(textEl.querySelectorAll('tspan'));
    const targetNodes = tspans.length > 0 ? tspans : [textEl];
    const firstNode = targetNodes[0];
    const textContent = firstNode.textContent || '';

    if (textContent.includes('📝')) {
      const textX = parseFloat(firstNode.getAttribute('x') || textEl.getAttribute('x') || '0');
      const textY = parseFloat(firstNode.getAttribute('y') || textEl.getAttribute('y') || '0');
      const fontSize = parseFloat(textEl.getAttribute('font-size') || '16');

      const iconX = textX;
      const iconY = textY - fontSize * 0.82;

      // Create a crisp vector note/memo icon
      const iconG = doc.createElementNS('http://www.w3.org/2000/svg', 'g');
      iconG.setAttribute('class', 'note-icon');
      iconG.setAttribute('transform', `translate(${iconX} ${iconY})`);

      // Note page background with folded top-right corner
      const paper = doc.createElementNS('http://www.w3.org/2000/svg', 'path');
      paper.setAttribute('d', 'M 0 0 L 8 0 L 12 4 L 12 14 L 0 14 Z');
      paper.setAttribute('fill', '#FFFFFF');
      paper.setAttribute('stroke', '#E65100');
      paper.setAttribute('stroke-width', '1.2');
      paper.setAttribute('stroke-linejoin', 'round');

      // Folded corner flap
      const fold = doc.createElementNS('http://www.w3.org/2000/svg', 'path');
      fold.setAttribute('d', 'M 8 0 L 8 4 L 12 4 Z');
      fold.setAttribute('fill', '#FFB300');
      fold.setAttribute('stroke', '#E65100');
      fold.setAttribute('stroke-width', '1.2');
      fold.setAttribute('stroke-linejoin', 'round');

      // Note text lines
      const lines = doc.createElementNS('http://www.w3.org/2000/svg', 'path');
      lines.setAttribute('d', 'M 2.5 5.5 L 9.5 5.5 M 2.5 8.5 L 9.5 8.5 M 2.5 11.5 L 7 11.5');
      lines.setAttribute('fill', 'none');
      lines.setAttribute('stroke', '#757575');
      lines.setAttribute('stroke-width', '1.0');
      lines.setAttribute('stroke-linecap', 'round');

      iconG.appendChild(paper);
      iconG.appendChild(fold);
      iconG.appendChild(lines);

      textEl.parentNode?.insertBefore(iconG, textEl);

      // Remove the emoji from the text node
      firstNode.textContent = textContent.replace(/📝\s*/g, '');

      // Shift text to the right of the icon (icon width 12px + 6px spacing = 18px)
      const shiftAmount = 18;
      for (const node of targetNodes) {
        const curX = parseFloat(node.getAttribute('x') || String(textX));
        node.setAttribute('x', String(curX + shiftAmount));
      }
      if (textEl.hasAttribute('x')) {
        const curX = parseFloat(textEl.getAttribute('x') || '0');
        textEl.setAttribute('x', String(curX + shiftAmount));
      }
    } else if (textContent.includes('🔗')) {
      const textX = parseFloat(firstNode.getAttribute('x') || textEl.getAttribute('x') || '0');
      const textY = parseFloat(firstNode.getAttribute('y') || textEl.getAttribute('y') || '0');
      const fontSize = parseFloat(textEl.getAttribute('font-size') || '16');

      const iconX = textX;
      const iconY = textY - fontSize * 0.75;

      // Create a crisp vector link icon
      const iconG = doc.createElementNS('http://www.w3.org/2000/svg', 'g');
      iconG.setAttribute('class', 'link-icon');
      iconG.setAttribute('transform', `translate(${iconX} ${iconY})`);

      const link = doc.createElementNS('http://www.w3.org/2000/svg', 'path');
      link.setAttribute('d', 'M 4.5 7.5 L 7.5 4.5 M 6 3 L 7.5 1.5 C 8.9 0.1 11.1 0.1 12.5 1.5 C 13.9 2.9 13.9 5.1 12.5 6.5 L 11 8 M 8 11 L 6.5 12.5 C 5.1 13.9 2.9 13.9 1.5 12.5 C 0.1 11.1 0.1 8.9 1.5 7.5 L 3 6');
      link.setAttribute('fill', 'none');
      link.setAttribute('stroke', '#666666');
      link.setAttribute('stroke-width', '1.2');
      link.setAttribute('stroke-linecap', 'round');

      iconG.appendChild(link);
      textEl.parentNode?.insertBefore(iconG, textEl);

      firstNode.textContent = textContent.replace(/🔗\s*/g, '');

      const shiftAmount = 18;
      for (const node of targetNodes) {
        const curX = parseFloat(node.getAttribute('x') || String(textX));
        node.setAttribute('x', String(curX + shiftAmount));
      }
      if (textEl.hasAttribute('x')) {
        const curX = parseFloat(textEl.getAttribute('x') || '0');
        textEl.setAttribute('x', String(curX + shiftAmount));
      }
    }
  }

  // Strip any remaining emoji characters in text elements to prevent PowerPoint from displaying black rectangles
  const allTextAndTspans = Array.from(svg.querySelectorAll('text, tspan'));
  for (const el of allTextAndTspans) {
    for (const child of Array.from(el.childNodes)) {
      if (child.nodeType === 3 /* Node.TEXT_NODE */) {
        const original = child.textContent || '';
        const cleaned = original.replace(/[\u{1F300}-\u{1FAFF}\u{2600}-\u{27BF}]/gu, '');
        if (cleaned !== original) {
          child.textContent = cleaned;
        }
      }
    }
  }

  // Initialize boundaries
  let minX = Infinity;
  let minY = Infinity;
  let maxX = -Infinity;
  let maxY = -Infinity;
  let measuredByDOM = false;

  // When running in a browser environment, measure the clean SVG using the browser's live layout engine
  if (typeof document !== 'undefined' && document.body) {
    try {
      const tempSvg = svg.cloneNode(true) as SVGSVGElement;
      tempSvg.style.position = 'absolute';
      tempSvg.style.visibility = 'hidden';
      tempSvg.style.left = '-99999px';
      tempSvg.style.top = '-99999px';
      document.body.appendChild(tempSvg);

      const bbox = tempSvg.getBBox();
      if (bbox && !isNaN(bbox.x) && !isNaN(bbox.y) && bbox.width > 0 && bbox.height > 0) {
        minX = bbox.x;
        minY = bbox.y;
        maxX = bbox.x + bbox.width;
        maxY = bbox.y + bbox.height;
        measuredByDOM = true;
      }
      document.body.removeChild(tempSvg);
    } catch (_) {
      // In case getBBox fails, fall back to manual measurement
    }
  }

  // Manual fallback measurement when DOM measurement is unavailable
  if (!measuredByDOM) {
    const elements = Array.from(svg.getElementsByTagName('*'));

    for (const element of elements) {
      const tagName = element.tagName.toLowerCase();

      // Rect elements
      if (tagName === 'rect') {
        const x = parseFloat(element.getAttribute('x') || '0');
        const y = parseFloat(element.getAttribute('y') || '0');
        const width = parseFloat(element.getAttribute('width') || '0');
        const height = parseFloat(element.getAttribute('height') || '0');

        let tx = 0, ty = 0;
        const transform = element.getAttribute('transform');
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

      // Circle elements
      if (tagName === 'circle') {
        const cx = parseFloat(element.getAttribute('cx') || '0');
        const cy = parseFloat(element.getAttribute('cy') || '0');
        const r = parseFloat(element.getAttribute('r') || '0');
        minX = Math.min(minX, cx - r);
        maxX = Math.max(maxX, cx + r);
        minY = Math.min(minY, cy - r);
        maxY = Math.max(maxY, cy + r);
        continue;
      }

      // Ellipse elements
      if (tagName === 'ellipse') {
        const cx = parseFloat(element.getAttribute('cx') || '0');
        const cy = parseFloat(element.getAttribute('cy') || '0');
        const rx = parseFloat(element.getAttribute('rx') || '0');
        const ry = parseFloat(element.getAttribute('ry') || '0');
        minX = Math.min(minX, cx - rx);
        maxX = Math.max(maxX, cx + rx);
        minY = Math.min(minY, cy - ry);
        maxY = Math.max(maxY, cy + ry);
        continue;
      }

      // Line elements
      if (tagName === 'line') {
        const x1 = parseFloat(element.getAttribute('x1') || '');
        const y1 = parseFloat(element.getAttribute('y1') || '');
        const x2 = parseFloat(element.getAttribute('x2') || '');
        const y2 = parseFloat(element.getAttribute('y2') || '');
        if (!isNaN(x1)) { minX = Math.min(minX, x1); maxX = Math.max(maxX, x1); }
        if (!isNaN(x2)) { minX = Math.min(minX, x2); maxX = Math.max(maxX, x2); }
        if (!isNaN(y1)) { minY = Math.min(minY, y1); maxY = Math.max(maxY, y1); }
        if (!isNaN(y2)) { minY = Math.min(minY, y2); maxY = Math.max(maxY, y2); }
        continue;
      }

      // Text and TSpan elements
      if (tagName === 'text' || tagName === 'tspan') {
        const x = parseFloat(element.getAttribute('x') || '0');
        const y = parseFloat(element.getAttribute('y') || '0');
        const text = (element.textContent || '').trim();
        if (text.length > 0) {
          const fontSize = parseFloat(element.getAttribute('font-size') || '16') || 16;
          const textAnchor = element.getAttribute('text-anchor') || 'start';
          let textWidth = text.length * fontSize * 0.65;
          if (typeof document !== 'undefined') {
            try {
              const canvas = document.createElement('canvas');
              const ctx = canvas.getContext('2d');
              if (ctx) {
                const fontFamily = element.getAttribute('font-family') || 'Roboto, Arial, sans-serif';
                ctx.font = `${fontSize}px ${fontFamily}`;
                textWidth = ctx.measureText(text).width;
              }
            } catch (_) {}
          }
          let tMinX = x;
          let tMaxX = x + textWidth;
          if (textAnchor === 'middle') {
            tMinX = x - textWidth / 2;
            tMaxX = x + textWidth / 2;
          } else if (textAnchor === 'end') {
            tMinX = x - textWidth;
            tMaxX = x;
          }
          minX = Math.min(minX, tMinX);
          maxX = Math.max(maxX, tMaxX);
          minY = Math.min(minY, y - fontSize * 0.85);
          maxY = Math.max(maxY, y + fontSize * 0.25);
        }
        continue;
      }

      // Path elements
      if (tagName === 'path') {
        const d = element.getAttribute('d');
        if (d) {
          const matches = d.matchAll(/([MmLlHhVvCcSsQqTtAaZz])([^MmLlHhVvCcSsQqTtAaZz]*)/g);
          let curX = 0, curY = 0;
          for (const match of matches) {
            const type = match[1];
            const nums = match[2].trim().split(/[\s,]+/).filter(Boolean).map(parseFloat);
            if (type === 'M' || type === 'L' || type === 'T') {
              for (let i = 0; i < nums.length - 1; i += 2) {
                curX = nums[i]; curY = nums[i + 1];
                minX = Math.min(minX, curX); maxX = Math.max(maxX, curX);
                minY = Math.min(minY, curY); maxY = Math.max(maxY, curY);
              }
            } else if (type === 'm' || type === 'l' || type === 't') {
              for (let i = 0; i < nums.length - 1; i += 2) {
                curX += nums[i]; curY += nums[i + 1];
                minX = Math.min(minX, curX); maxX = Math.max(maxX, curX);
                minY = Math.min(minY, curY); maxY = Math.max(maxY, curY);
              }
            } else if (type === 'H') {
              for (const n of nums) { curX = n; minX = Math.min(minX, curX); maxX = Math.max(maxX, curX); }
            } else if (type === 'h') {
              for (const n of nums) { curX += n; minX = Math.min(minX, curX); maxX = Math.max(maxX, curX); }
            } else if (type === 'V') {
              for (const n of nums) { curY = n; minY = Math.min(minY, curY); maxY = Math.max(maxY, curY); }
            } else if (type === 'v') {
              for (const n of nums) { curY += n; minY = Math.min(minY, curY); maxY = Math.max(maxY, curY); }
            } else if (type === 'C') {
              for (let i = 0; i < nums.length - 5; i += 6) {
                [nums[i], nums[i + 2], nums[i + 4]].forEach(x => { minX = Math.min(minX, x); maxX = Math.max(maxX, x); });
                [nums[i + 1], nums[i + 3], nums[i + 5]].forEach(y => { minY = Math.min(minY, y); maxY = Math.max(maxY, y); });
                curX = nums[i + 4]; curY = nums[i + 5];
              }
            } else if (type === 'c') {
              for (let i = 0; i < nums.length - 5; i += 6) {
                [curX + nums[i], curX + nums[i + 2], curX + nums[i + 4]].forEach(x => { minX = Math.min(minX, x); maxX = Math.max(maxX, x); });
                [curY + nums[i + 1], curY + nums[i + 3], curY + nums[i + 5]].forEach(y => { minY = Math.min(minY, y); maxY = Math.max(maxY, y); });
                curX += nums[i + 4]; curY += nums[i + 5];
              }
            } else if (type === 'S' || type === 'Q') {
              for (let i = 0; i < nums.length - 3; i += 4) {
                [nums[i], nums[i + 2]].forEach(x => { minX = Math.min(minX, x); maxX = Math.max(maxX, x); });
                [nums[i + 1], nums[i + 3]].forEach(y => { minY = Math.min(minY, y); maxY = Math.max(maxY, y); });
                curX = nums[i + 2]; curY = nums[i + 3];
              }
            } else if (type === 's' || type === 'q') {
              for (let i = 0; i < nums.length - 3; i += 4) {
                [curX + nums[i], curX + nums[i + 2]].forEach(x => { minX = Math.min(minX, x); maxX = Math.max(maxX, x); });
                [curY + nums[i + 1], curY + nums[i + 3]].forEach(y => { minY = Math.min(minY, y); maxY = Math.max(maxY, y); });
                curX += nums[i + 2]; curY += nums[i + 3];
              }
            } else if (type === 'A') {
              // A rx ry x-axis-rotation large-arc-flag sweep-flag x y
              for (let i = 0; i < nums.length - 6; i += 7) {
                const x = nums[i + 5], y = nums[i + 6];
                minX = Math.min(minX, x); maxX = Math.max(maxX, x);
                minY = Math.min(minY, y); maxY = Math.max(maxY, y);
                curX = x; curY = y;
              }
            } else if (type === 'a') {
              for (let i = 0; i < nums.length - 6; i += 7) {
                curX += nums[i + 5]; curY += nums[i + 6];
                minX = Math.min(minX, curX); maxX = Math.max(maxX, curX);
                minY = Math.min(minY, curY); maxY = Math.max(maxY, curY);
              }
            }
          }
        }
        continue;
      }
    }
  }

  // Add padding
  if (isFinite(minX) && isFinite(maxX) && isFinite(minY) && isFinite(maxY)) {
    const padding = 10;
    minX -= padding;
    minY -= padding;
    maxX += padding;
    maxY += padding;

    // Calculate new dimensions
    const width = Math.max(maxX - minX, 100);
    const height = Math.max(maxY - minY, 100);

    // Update SVG attributes
    svg.setAttribute('width', width.toString());
    svg.setAttribute('height', height.toString());
    svg.setAttribute('viewBox', `${minX} ${minY} ${width} ${height}`);
  }

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
