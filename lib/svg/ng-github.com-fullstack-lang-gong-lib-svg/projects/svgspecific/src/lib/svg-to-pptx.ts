export function convertSvgToPptx(svgString: string): Uint8Array {
  const parser = new DOMParser();
  const doc = parser.parseFromString(svgString, 'image/svg+xml');
  const svgEl = doc.documentElement;

  // Parse viewBox or width/height
  let vbX = 0;
  let vbY = 0;
  let vbW = parseFloat(svgEl.getAttribute('width') || '800') || 800;
  let vbH = parseFloat(svgEl.getAttribute('height') || '600') || 600;

  const viewBoxAttr = svgEl.getAttribute('viewBox');
  if (viewBoxAttr) {
    const parts = viewBoxAttr.trim().split(/[\s,]+/).map(parseFloat);
    if (parts.length === 4 && parts.every(n => !isNaN(n))) {
      vbX = parts[0];
      vbY = parts[1];
      vbW = parts[2] || vbW;
      vbH = parts[3] || vbH;
    }
  }

  // Standard 16:9 widescreen slide in EMUs (13.333 in x 7.5 in)
  // 1 inch = 914400 EMUs; 1 pt = 12700 EMUs; 1 px (96 DPI) = 9525 EMUs
  const SLIDE_W = 12192000;
  const SLIDE_H = 6858000;
  const MARGIN = 457200; // 0.5 inch margin

  const availW = SLIDE_W - 2 * MARGIN;
  const availH = SLIDE_H - 2 * MARGIN;

  // Scale factor in EMUs per SVG user unit.
  // Default 1 SVG px = 9525 EMUs (1:1 at 96 DPI), scaled down if diagram is larger than slide.
  const naturalW = vbW * 9525;
  const naturalH = vbH * 9525;
  let scaleEmu = 9525;
  if (naturalW > availW || naturalH > availH) {
    scaleEmu = Math.min(availW / Math.max(vbW, 1), availH / Math.max(vbH, 1));
  }

  const renderedW = vbW * scaleEmu;
  const renderedH = vbH * scaleEmu;
  const offsetX = Math.round((SLIDE_W - renderedW) / 2);
  const offsetY = Math.round((SLIDE_H - renderedH) / 2);

  const toEmuX = (x: number): number => Math.round(offsetX + (x - vbX) * scaleEmu);
  const toEmuY = (y: number): number => Math.round(offsetY + (y - vbY) * scaleEmu);
  const toEmuLen = (len: number): number => Math.round(len * scaleEmu);
  const pxToPt = scaleEmu / 12700; // converts SVG px to PowerPoint points after slide scaling

  let shapeIdCounter = 2; // 1 is reserved for nvGrpSpPr of spTree
  const shapesXml: string[] = [];
  const mediaFiles: { fileName: string; rId: string; data: Uint8Array }[] = [];

  // Pre-collect RectAnchoredText elements that belong inside their parent <rect> shape
  const rectEmbeddedTexts = new Map<string, Element>();
  const consumedTextElements = new Set<Element>();

  const allRectsById = new Map<string, Element>();
  for (const r of Array.from(svgEl.querySelectorAll('rect[data-rect-id]'))) {
    const id = r.getAttribute('data-rect-id');
    if (id) {
      allRectsById.set(id, r);
    }
  }

  for (const textEl of Array.from(svgEl.querySelectorAll('text[data-parent-rect-id]'))) {
    const parentId = textEl.getAttribute('data-parent-rect-id') || '';
    if (!parentId || rectEmbeddedTexts.has(parentId)) {
      continue;
    }
    const parentRect = allRectsById.get(parentId);
    if (!parentRect) {
      continue;
    }
    const rectW = parseFloat(parentRect.getAttribute('width') || '0');
    const rectH = parseFloat(parentRect.getAttribute('height') || '0');
    if (rectW <= 0 || rectH <= 0) {
      continue;
    }

    const rectX = parseFloat(parentRect.getAttribute('x') || '0');
    const rectY = parseFloat(parentRect.getAttribute('y') || '0');
    const firstTspan = textEl.querySelector('tspan');
    const anchorX = parseFloat(firstTspan?.getAttribute('x') || textEl.getAttribute('x') || '0');
    const anchorY = parseFloat(firstTspan?.getAttribute('y') || textEl.getAttribute('y') || '0');

    // Only embed the text inside the rect shape if its anchor point lies within the rect bounds
    // (excludes lane header labels or external date/milestone labels anchored outside the bar)
    const tol = 2;
    if (
      anchorX >= rectX - tol &&
      anchorX <= rectX + rectW + tol &&
      anchorY >= rectY - tol &&
      anchorY <= rectY + rectH + tol
    ) {
      rectEmbeddedTexts.set(parentId, textEl);
      consumedTextElements.add(textEl);
    }
  }

  interface TransformState {
    tx: number;
    ty: number;
    sx: number;
    sy: number;
    rotDeg: number;
  }

  function parseTransform(attr: string | null, parent: TransformState): TransformState {
    let { tx, ty, sx, sy, rotDeg } = parent;
    if (!attr) {
      return { tx, ty, sx, sy, rotDeg };
    }
    const regex = /(translate|scale|rotate)\s*\(([^)]+)\)/g;
    let match: RegExpExecArray | null;
    while ((match = regex.exec(attr)) !== null) {
      const fn = match[1];
      const args = match[2].trim().split(/[\s,]+/).map(parseFloat);
      if (fn === 'translate') {
        const dx = args[0] || 0;
        const dy = args.length > 1 ? (args[1] || 0) : 0;
        tx += dx * sx;
        ty += dy * sy;
      } else if (fn === 'scale') {
        const dsx = args[0] !== undefined && !isNaN(args[0]) ? args[0] : 1;
        const dsy = args.length > 1 && !isNaN(args[1]) ? args[1] : dsx;
        sx *= dsx;
        sy *= dsy;
      } else if (fn === 'rotate') {
        const angle = args[0] || 0;
        rotDeg = (rotDeg + angle) % 360;
      }
    }
    return { tx, ty, sx, sy, rotDeg };
  }

  function walkElement(el: Element, parentTransform: TransformState): void {
    const tag = el.tagName.replace(/^.*:/, '').toLowerCase();
    const curTransform = parseTransform(el.getAttribute('transform'), parentTransform);

    if (tag === 'style' || tag === 'defs') {
      return;
    }

    if (tag === 'svg' || tag === 'g') {
      for (const child of Array.from(el.children)) {
        walkElement(child, curTransform);
      }
      return;
    }

    if (tag === 'rect') {
      emitRect(el, curTransform);
    } else if (tag === 'circle') {
      emitCircle(el, curTransform);
    } else if (tag === 'ellipse') {
      emitEllipse(el, curTransform);
    } else if (tag === 'line') {
      emitLine(el, curTransform);
    } else if (tag === 'polygon' || tag === 'polyline') {
      emitPoly(el, curTransform, tag === 'polygon');
    } else if (tag === 'path') {
      emitPath(el, curTransform);
    } else if (tag === 'text') {
      if (!consumedTextElements.has(el)) {
        emitText(el, curTransform);
      }
    } else if (tag === 'image') {
      emitImage(el, curTransform);
    }
  }

  function buildFillXml(el: Element, defaultFill = 'none'): string {
    const rawFill = (el.getAttribute('fill') ?? defaultFill).trim();
    const fillOpacity = parseOpacity(el.getAttribute('fill-opacity'), el.getAttribute('opacity'));
    if (!rawFill || rawFill.toLowerCase() === 'none' || rawFill.toLowerCase() === 'transparent' || fillOpacity <= 0) {
      return '<a:noFill/>';
    }
    const hex = toHex6(rawFill);
    if (!hex) {
      return '<a:noFill/>';
    }
    if (fillOpacity < 1) {
      const alphaVal = Math.round(fillOpacity * 100000);
      return `<a:solidFill><a:srgbClr val="${hex}"><a:alpha val="${alphaVal}"/></a:srgbClr></a:solidFill>`;
    }
    return `<a:solidFill><a:srgbClr val="${hex}"/></a:solidFill>`;
  }

  function buildLineXml(el: Element): string {
    const rawStroke = (el.getAttribute('stroke') || '').trim();
    const strokeWidthPx = parseFloat(el.getAttribute('stroke-width') || '1');
    const strokeOpacity = parseOpacity(el.getAttribute('stroke-opacity'), el.getAttribute('opacity'));

    if (
      !rawStroke ||
      rawStroke.toLowerCase() === 'none' ||
      rawStroke.toLowerCase() === 'transparent' ||
      isNaN(strokeWidthPx) ||
      strokeWidthPx <= 0 ||
      strokeOpacity <= 0
    ) {
      return '<a:ln><a:noFill/></a:ln>';
    }

    const hex = toHex6(rawStroke) || '000000';
    const widthEmu = Math.max(Math.round(strokeWidthPx * scaleEmu), 6350); // min 0.5pt
    const alphaXml = strokeOpacity < 1 ? `<a:alpha val="${Math.round(strokeOpacity * 100000)}"/>` : '';
    const dashAttr = (el.getAttribute('stroke-dasharray') || '').trim();
    let dashXml = '';
    if (dashAttr && dashAttr !== 'none') {
      const nums = dashAttr.split(/[\s,]+/).map(parseFloat).filter(n => !isNaN(n));
      if (nums.length > 0 && nums[0] <= 3) {
        dashXml = '<a:prstDash val="sysDot"/>';
      } else if (nums.length > 0) {
        dashXml = '<a:prstDash val="dash"/>';
      }
    }

    return `<a:ln w="${widthEmu}"><a:solidFill><a:srgbClr val="${hex}">${alphaXml}</a:srgbClr></a:solidFill>${dashXml}</a:ln>`;
  }

  function extractTextLines(el: Element): string[] {
    const tspans = Array.from(el.querySelectorAll('tspan'));
    const lines: string[] = [];
    if (tspans.length > 0) {
      for (const ts of tspans) {
        lines.push(ts.textContent ?? '');
      }
    } else {
      const txt = el.textContent ?? '';
      if (txt.trim().length > 0) {
        lines.push(txt);
      }
    }
    return lines;
  }

  function buildParagraphsXml(el: Element, lines: string[], sy: number, algn: string): string {
    const rawFontSize = parseFloat(el.getAttribute('font-size') || '14') || 14;
    const scaledFontSizePx = rawFontSize * Math.abs(sy);
    const fontPt = Math.max(scaledFontSizePx * pxToPt, 4);
    const szHundredths = Math.round(fontPt * 100);

    const fontWeight = (el.getAttribute('font-weight') || '').trim().toLowerCase();
    const isBold = fontWeight === 'bold' || parseInt(fontWeight, 10) >= 600;
    const fontStyle = (el.getAttribute('font-style') || '').trim().toLowerCase();
    const isItalic = fontStyle === 'italic' || fontStyle === 'oblique';

    const fontFamilyAttr =
      (el.getAttribute('font-family') || 'Arial')
        .split(',')[0]
        .replace(/['"]/g, '')
        .trim() || 'Arial';

    const rawFill = (el.getAttribute('fill') || '#000000').trim();
    const fillHex = toHex6(rawFill) || '000000';
    const fillOpacity = parseOpacity(el.getAttribute('fill-opacity'), el.getAttribute('opacity'));
    const alphaXml = fillOpacity < 1 ? `<a:alpha val="${Math.round(fillOpacity * 100000)}"/>` : '';

    const boldAttr = isBold ? ' b="1"' : '';
    const italicAttr = isItalic ? ' i="1"' : '';

    return lines
      .map(lineText => {
        const escaped = escapeXml(lineText);
        return `<a:p>
          <a:pPr algn="${algn}"/>
          <a:r>
            <a:rPr lang="en-US" sz="${szHundredths}"${boldAttr}${italicAttr} dirty="0">
              <a:solidFill><a:srgbClr val="${fillHex}">${alphaXml}</a:srgbClr></a:solidFill>
              <a:latin typeface="${escapeXml(fontFamilyAttr)}"/>
              <a:cs typeface="${escapeXml(fontFamilyAttr)}"/>
            </a:rPr>
            <a:t>${escaped}</a:t>
          </a:r>
        </a:p>`;
      })
      .join('');
  }

  function buildEmbeddedTxBodyXml(rectEl: Element, textEl: Element, t: TransformState): string {
    const lines = extractTextLines(textEl);
    if (lines.length === 0 || lines.every(l => l.trim().length === 0)) {
      return '';
    }

    const firstTspan = textEl.querySelector('tspan');
    const textAnchor = (
      firstTspan?.getAttribute('text-anchor') ||
      textEl.getAttribute('text-anchor') ||
      'start'
    )
      .trim()
      .toLowerCase();

    let algn = 'l';
    if (textAnchor === 'middle' || textAnchor === 'center') {
      algn = 'ctr';
    } else if (textAnchor === 'end') {
      algn = 'r';
    }

    const rectAnchorType = (textEl.getAttribute('data-rect-anchor-type') || '').trim();
    const dominantBaseline = (textEl.getAttribute('dominant-baseline') || '').trim().toLowerCase();

    let anchor = 'ctr'; // vertical alignment: 't' (top), 'ctr' (middle), 'b' (bottom)
    if (rectAnchorType.includes('TOP')) {
      anchor = 't';
    } else if (rectAnchorType.includes('BOTTOM')) {
      anchor = 'b';
    } else if (
      rectAnchorType.includes('MIDDLE') ||
      rectAnchorType.includes('CENTER') ||
      dominantBaseline === 'central' ||
      dominantBaseline === 'middle'
    ) {
      anchor = 'ctr';
    }

    // Compute horizontal & vertical insets based on text anchor offset relative to rect
    const rectX = parseFloat(rectEl.getAttribute('x') || '0');
    const rectY = parseFloat(rectEl.getAttribute('y') || '0');
    const rectW = parseFloat(rectEl.getAttribute('width') || '0');
    const anchorX = parseFloat(firstTspan?.getAttribute('x') || textEl.getAttribute('x') || '0');
    const anchorY = parseFloat(firstTspan?.getAttribute('y') || textEl.getAttribute('y') || '0');
    const rawFontSize = parseFloat(textEl.getAttribute('font-size') || '14') || 14;

    let lIns = toEmuLen(4);
    let rIns = toEmuLen(4);
    let tIns = toEmuLen(2);
    let bIns = toEmuLen(2);

    if (algn === 'l') {
      const leftPadPx = Math.max(anchorX - rectX, 4);
      lIns = toEmuLen(leftPadPx);
    } else if (algn === 'r') {
      const rightPadPx = Math.max(rectX + rectW - anchorX, 4);
      rIns = toEmuLen(rightPadPx);
    } else if (algn === 'ctr') {
      // Check if centered text was shifted horizontally (e.g. for a left completion/person icon)
      const shiftFromCenterPx = anchorX - (rectX + rectW / 2);
      if (shiftFromCenterPx > 2) {
        lIns = toEmuLen(shiftFromCenterPx * 2);
      } else if (shiftFromCenterPx < -2) {
        rIns = toEmuLen(-shiftFromCenterPx * 2);
      }
    }

    if (anchor === 't') {
      // In SVG, anchorY is the first line's baseline; top of text is ~ anchorY - 0.85 * fontSize
      const topPadPx = Math.max(anchorY - rectY - rawFontSize * 0.85, 2);
      tIns = toEmuLen(topPadPx);
    }

    const paragraphsXml = buildParagraphsXml(textEl, lines, t.sy, algn);

    return `
        <p:txBody>
          <a:bodyPr wrap="square" lIns="${lIns}" tIns="${tIns}" rIns="${rIns}" bIns="${bIns}" anchor="${anchor}"/>
          <a:lstStyle/>
          ${paragraphsXml}
        </p:txBody>`;
  }

  function emitRect(el: Element, t: TransformState): void {
    const x = parseFloat(el.getAttribute('x') || '0') * t.sx + t.tx;
    const y = parseFloat(el.getAttribute('y') || '0') * t.sy + t.ty;
    const w = parseFloat(el.getAttribute('width') || '0') * Math.abs(t.sx);
    const h = parseFloat(el.getAttribute('height') || '0') * Math.abs(t.sy);
    if (w <= 0 || h <= 0) {
      return;
    }

    // Invisible placeholder rects (e.g. laneHeaderRect with fill-opacity <= 0.001 and stroke="none")
    // that have no embedded text can be skipped to avoid invisible clutter shapes on the slide
    const fillOp = parseOpacity(el.getAttribute('fill-opacity'), el.getAttribute('opacity'));
    const rawFill = (el.getAttribute('fill') || '').trim().toLowerCase();
    const rawStroke = (el.getAttribute('stroke') || '').trim().toLowerCase();
    const strokeW = parseFloat(el.getAttribute('stroke-width') || '1');
    const strokeOp = parseOpacity(el.getAttribute('stroke-opacity'), el.getAttribute('opacity'));
    const rectId = el.getAttribute('data-rect-id') || '';
    const embeddedTextEl = rectId ? rectEmbeddedTexts.get(rectId) : undefined;

    const isFillInvisible = !rawFill || rawFill === 'none' || rawFill === 'transparent' || fillOp <= 0.001;
    const isStrokeInvisible = !rawStroke || rawStroke === 'none' || rawStroke === 'transparent' || strokeW <= 0 || strokeOp <= 0.001;
    if (isFillInvisible && isStrokeInvisible && !embeddedTextEl) {
      return;
    }

    const rx = parseFloat(el.getAttribute('rx') || '0') * Math.abs(t.sx);
    const emuX = toEmuX(x);
    const emuY = toEmuY(y);
    const emuW = Math.max(toEmuLen(w), 1);
    const emuH = Math.max(toEmuLen(h), 1);

    let prst = 'rect';
    let avLst = '<a:avLst/>';
    if (rx > 0) {
      prst = 'roundRect';
      const minDim = Math.min(w, h);
      const ratio = minDim > 0 ? Math.min(Math.max(rx / minDim, 0), 0.5) : 0;
      const adj = Math.round(ratio * 100000);
      avLst = `<a:avLst><a:gd name="adj" fmla="val ${adj}"/></a:avLst>`;
    }

    // PowerPoint rotation is in 60,000ths of a degree
    const rotAttr = t.rotDeg !== 0 ? ` rot="${Math.round(t.rotDeg * 60000)}"` : '';

    const fillXml = isFillInvisible ? '<a:noFill/>' : buildFillXml(el, '#000000');
    const lnXml = buildLineXml(el);
    const txBodyXml = embeddedTextEl ? buildEmbeddedTxBodyXml(el, embeddedTextEl, t) : '';
    const id = shapeIdCounter++;

    shapesXml.push(`
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="${id}" name="Rect ${id}"/>
          <p:cNvSpPr/>
          <p:nvPr/>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm${rotAttr}>
            <a:off x="${emuX}" y="${emuY}"/>
            <a:ext cx="${emuW}" cy="${emuH}"/>
          </a:xfrm>
          <a:prstGeom prst="${prst}">${avLst}</a:prstGeom>
          ${fillXml}
          ${lnXml}
        </p:spPr>${txBodyXml}
      </p:sp>`);
  }

  function emitCircle(el: Element, t: TransformState): void {
    const cx = parseFloat(el.getAttribute('cx') || '0') * t.sx + t.tx;
    const cy = parseFloat(el.getAttribute('cy') || '0') * t.sy + t.ty;
    const r = parseFloat(el.getAttribute('r') || '0') * Math.abs(t.sx);
    if (r <= 0) {
      return;
    }
    emitOval(el, cx - r, cy - r, r * 2, r * 2);
  }

  function emitEllipse(el: Element, t: TransformState): void {
    const cx = parseFloat(el.getAttribute('cx') || '0') * t.sx + t.tx;
    const cy = parseFloat(el.getAttribute('cy') || '0') * t.sy + t.ty;
    const rx = parseFloat(el.getAttribute('rx') || '0') * Math.abs(t.sx);
    const ry = parseFloat(el.getAttribute('ry') || '0') * Math.abs(t.sy);
    if (rx <= 0 || ry <= 0) {
      return;
    }
    emitOval(el, cx - rx, cy - ry, rx * 2, ry * 2);
  }

  function emitOval(el: Element, x: number, y: number, w: number, h: number): void {
    const emuX = toEmuX(x);
    const emuY = toEmuY(y);
    const emuW = Math.max(toEmuLen(w), 1);
    const emuH = Math.max(toEmuLen(h), 1);
    const fillXml = buildFillXml(el, '#000000');
    const lnXml = buildLineXml(el);
    const id = shapeIdCounter++;

    shapesXml.push(`
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="${id}" name="Ellipse ${id}"/>
          <p:cNvSpPr/>
          <p:nvPr/>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="${emuX}" y="${emuY}"/>
            <a:ext cx="${emuW}" cy="${emuH}"/>
          </a:xfrm>
          <a:prstGeom prst="ellipse"><a:avLst/></a:prstGeom>
          ${fillXml}
          ${lnXml}
        </p:spPr>
      </p:sp>`);
  }

  function emitLine(el: Element, t: TransformState): void {
    const x1 = parseFloat(el.getAttribute('x1') || '0') * t.sx + t.tx;
    const y1 = parseFloat(el.getAttribute('y1') || '0') * t.sy + t.ty;
    const x2 = parseFloat(el.getAttribute('x2') || '0') * t.sx + t.tx;
    const y2 = parseFloat(el.getAttribute('y2') || '0') * t.sy + t.ty;

    const ex1 = toEmuX(x1);
    const ey1 = toEmuY(y1);
    const ex2 = toEmuX(x2);
    const ey2 = toEmuY(y2);

    const offX = Math.min(ex1, ex2);
    const offY = Math.min(ey1, ey2);
    const extW = Math.abs(ex2 - ex1);
    const extH = Math.abs(ey2 - ey1);
    const flipH = ex2 < ex1 ? ' flipH="1"' : '';
    const flipV = ey2 < ey1 ? ' flipV="1"' : '';

    const lnXml = buildLineXml(el);
    const id = shapeIdCounter++;

    shapesXml.push(`
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="${id}" name="Line ${id}"/>
          <p:cNvSpPr/>
          <p:nvPr/>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm${flipH}${flipV}>
            <a:off x="${offX}" y="${offY}"/>
            <a:ext cx="${extW}" cy="${extH}"/>
          </a:xfrm>
          <a:prstGeom prst="line"><a:avLst/></a:prstGeom>
          <a:noFill/>
          ${lnXml}
        </p:spPr>
      </p:sp>`);
  }

  function emitPoly(el: Element, t: TransformState, close: boolean): void {
    const pointsAttr = (el.getAttribute('points') || '').trim();
    if (!pointsAttr) {
      return;
    }
    const nums = pointsAttr.split(/[\s,]+/).map(parseFloat).filter(n => !isNaN(n));
    if (nums.length < 4) {
      return;
    }
    const pts: { x: number; y: number }[] = [];
    for (let i = 0; i < nums.length - 1; i += 2) {
      pts.push({
        x: toEmuX(nums[i] * t.sx + t.tx),
        y: toEmuY(nums[i + 1] * t.sy + t.ty)
      });
    }
    let minX = Infinity, minY = Infinity, maxX = -Infinity, maxY = -Infinity;
    for (const p of pts) {
      minX = Math.min(minX, p.x);
      minY = Math.min(minY, p.y);
      maxX = Math.max(maxX, p.x);
      maxY = Math.max(maxY, p.y);
    }
    const extW = Math.max(maxX - minX, 1);
    const extH = Math.max(maxY - minY, 1);

    let pathCmds = `<a:moveTo><a:pt x="${pts[0].x - minX}" y="${pts[0].y - minY}"/></a:moveTo>`;
    for (let i = 1; i < pts.length; i++) {
      pathCmds += `<a:lnTo><a:pt x="${pts[i].x - minX}" y="${pts[i].y - minY}"/></a:lnTo>`;
    }
    if (close) {
      pathCmds += '<a:close/>';
    }

    const fillXml = close ? buildFillXml(el, '#000000') : '<a:noFill/>';
    const lnXml = buildLineXml(el);
    const id = shapeIdCounter++;

    shapesXml.push(`
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="${id}" name="${close ? 'Polygon' : 'Polyline'} ${id}"/>
          <p:cNvSpPr/>
          <p:nvPr/>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="${minX}" y="${minY}"/>
            <a:ext cx="${extW}" cy="${extH}"/>
          </a:xfrm>
          <a:custGeom>
            <a:avLst/>
            <a:gdLst/>
            <a:ahLst/>
            <a:cxnLst/>
            <a:rect l="0" t="0" r="r" b="b"/>
            <a:pathLst>
              <a:path w="${extW}" h="${extH}">
                ${pathCmds}
              </a:path>
            </a:pathLst>
          </a:custGeom>
          ${fillXml}
          ${lnXml}
        </p:spPr>
      </p:sp>`);
  }

  function emitPath(el: Element, t: TransformState): void {
    const d = (el.getAttribute('d') || '').trim();
    if (!d) {
      return;
    }

    interface PathCmd {
      op: 'M' | 'L' | 'C' | 'Q' | 'Z';
      pts: { x: number; y: number }[];
    }

    const cmds: PathCmd[] = [];
    let curX = 0;
    let curY = 0;
    let startX = 0;
    let startY = 0;

    const toAbsEmu = (x: number, y: number) => ({
      x: toEmuX(x * t.sx + t.tx),
      y: toEmuY(y * t.sy + t.ty)
    });

    const tokens = d.matchAll(/([MmLlHhVvCcSsQqTtAaZz])([^MmLlHhVvCcSsQqTtAaZz]*)/g);
    for (const token of tokens) {
      const type = token[1];
      const args = (token[2].match(/-?\d*\.?\d+(?:e[-+]?\d+)?/gi) || []).map(parseFloat);

      switch (type) {
        case 'M':
          for (let i = 0; i < args.length - 1; i += 2) {
            curX = args[i];
            curY = args[i + 1];
            if (i === 0) {
              startX = curX;
              startY = curY;
              cmds.push({ op: 'M', pts: [toAbsEmu(curX, curY)] });
            } else {
              cmds.push({ op: 'L', pts: [toAbsEmu(curX, curY)] });
            }
          }
          break;
        case 'm':
          for (let i = 0; i < args.length - 1; i += 2) {
            curX += args[i];
            curY += args[i + 1];
            if (i === 0) {
              startX = curX;
              startY = curY;
              cmds.push({ op: 'M', pts: [toAbsEmu(curX, curY)] });
            } else {
              cmds.push({ op: 'L', pts: [toAbsEmu(curX, curY)] });
            }
          }
          break;
        case 'L':
        case 'T':
          for (let i = 0; i < args.length - 1; i += 2) {
            curX = args[i];
            curY = args[i + 1];
            cmds.push({ op: 'L', pts: [toAbsEmu(curX, curY)] });
          }
          break;
        case 'l':
        case 't':
          for (let i = 0; i < args.length - 1; i += 2) {
            curX += args[i];
            curY += args[i + 1];
            cmds.push({ op: 'L', pts: [toAbsEmu(curX, curY)] });
          }
          break;
        case 'H':
          for (const x of args) {
            curX = x;
            cmds.push({ op: 'L', pts: [toAbsEmu(curX, curY)] });
          }
          break;
        case 'h':
          for (const dx of args) {
            curX += dx;
            cmds.push({ op: 'L', pts: [toAbsEmu(curX, curY)] });
          }
          break;
        case 'V':
          for (const y of args) {
            curY = y;
            cmds.push({ op: 'L', pts: [toAbsEmu(curX, curY)] });
          }
          break;
        case 'v':
          for (const dy of args) {
            curY += dy;
            cmds.push({ op: 'L', pts: [toAbsEmu(curX, curY)] });
          }
          break;
        case 'C':
          for (let i = 0; i < args.length - 5; i += 6) {
            const p1 = toAbsEmu(args[i], args[i + 1]);
            const p2 = toAbsEmu(args[i + 2], args[i + 3]);
            curX = args[i + 4];
            curY = args[i + 5];
            const p3 = toAbsEmu(curX, curY);
            cmds.push({ op: 'C', pts: [p1, p2, p3] });
          }
          break;
        case 'c':
          for (let i = 0; i < args.length - 5; i += 6) {
            const p1 = toAbsEmu(curX + args[i], curY + args[i + 1]);
            const p2 = toAbsEmu(curX + args[i + 2], curY + args[i + 3]);
            curX += args[i + 4];
            curY += args[i + 5];
            const p3 = toAbsEmu(curX, curY);
            cmds.push({ op: 'C', pts: [p1, p2, p3] });
          }
          break;
        case 'Q':
        case 'S':
          for (let i = 0; i < args.length - 3; i += 4) {
            const p1 = toAbsEmu(args[i], args[i + 1]);
            curX = args[i + 2];
            curY = args[i + 3];
            const p2 = toAbsEmu(curX, curY);
            cmds.push({ op: 'Q', pts: [p1, p2] });
          }
          break;
        case 'q':
        case 's':
          for (let i = 0; i < args.length - 3; i += 4) {
            const p1 = toAbsEmu(curX + args[i], curY + args[i + 1]);
            curX += args[i + 2];
            curY += args[i + 3];
            const p2 = toAbsEmu(curX, curY);
            cmds.push({ op: 'Q', pts: [p1, p2] });
          }
          break;
        case 'A':
          for (let i = 0; i < args.length - 6; i += 7) {
            const rx = args[i];
            const ry = args[i + 1];
            const xAxisRot = args[i + 2];
            const largeArcFlag = args[i + 3];
            const sweepFlag = args[i + 4];
            const endX = args[i + 5];
            const endY = args[i + 6];
            const arcPts = approximateArc(curX, curY, rx, ry, xAxisRot, largeArcFlag, sweepFlag, endX, endY);
            for (const pt of arcPts) {
              cmds.push({ op: 'L', pts: [toAbsEmu(pt.x, pt.y)] });
            }
            curX = endX;
            curY = endY;
          }
          break;
        case 'a':
          for (let i = 0; i < args.length - 6; i += 7) {
            const rx = args[i];
            const ry = args[i + 1];
            const xAxisRot = args[i + 2];
            const largeArcFlag = args[i + 3];
            const sweepFlag = args[i + 4];
            const endX = curX + args[i + 5];
            const endY = curY + args[i + 6];
            const arcPts = approximateArc(curX, curY, rx, ry, xAxisRot, largeArcFlag, sweepFlag, endX, endY);
            for (const pt of arcPts) {
              cmds.push({ op: 'L', pts: [toAbsEmu(pt.x, pt.y)] });
            }
            curX = endX;
            curY = endY;
          }
          break;
        case 'Z':
        case 'z':
          curX = startX;
          curY = startY;
          cmds.push({ op: 'Z', pts: [] });
          break;
      }
    }

    if (cmds.length === 0) {
      return;
    }

    let minX = Infinity, minY = Infinity, maxX = -Infinity, maxY = -Infinity;
    for (const c of cmds) {
      for (const p of c.pts) {
        minX = Math.min(minX, p.x);
        minY = Math.min(minY, p.y);
        maxX = Math.max(maxX, p.x);
        maxY = Math.max(maxY, p.y);
      }
    }
    if (!isFinite(minX) || !isFinite(minY)) {
      return;
    }

    const extW = Math.max(maxX - minX, 1);
    const extH = Math.max(maxY - minY, 1);

    let pathXml = '';
    for (const c of cmds) {
      if (c.op === 'M') {
        const p = c.pts[0];
        pathXml += `<a:moveTo><a:pt x="${p.x - minX}" y="${p.y - minY}"/></a:moveTo>`;
      } else if (c.op === 'L') {
        const p = c.pts[0];
        pathXml += `<a:lnTo><a:pt x="${p.x - minX}" y="${p.y - minY}"/></a:lnTo>`;
      } else if (c.op === 'C') {
        const [p1, p2, p3] = c.pts;
        pathXml += `<a:cubicBezTo><a:pt x="${p1.x - minX}" y="${p1.y - minY}"/><a:pt x="${p2.x - minX}" y="${p2.y - minY}"/><a:pt x="${p3.x - minX}" y="${p3.y - minY}"/></a:cubicBezTo>`;
      } else if (c.op === 'Q') {
        const [p1, p2] = c.pts;
        pathXml += `<a:quadBezTo><a:pt x="${p1.x - minX}" y="${p1.y - minY}"/><a:pt x="${p2.x - minX}" y="${p2.y - minY}"/></a:quadBezTo>`;
      } else if (c.op === 'Z') {
        pathXml += '<a:close/>';
      }
    }

    const fillXml = buildFillXml(el, 'none');
    const lnXml = buildLineXml(el);
    const id = shapeIdCounter++;

    shapesXml.push(`
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="${id}" name="Path ${id}"/>
          <p:cNvSpPr/>
          <p:nvPr/>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="${minX}" y="${minY}"/>
            <a:ext cx="${extW}" cy="${extH}"/>
          </a:xfrm>
          <a:custGeom>
            <a:avLst/>
            <a:gdLst/>
            <a:ahLst/>
            <a:cxnLst/>
            <a:rect l="0" t="0" r="r" b="b"/>
            <a:pathLst>
              <a:path w="${extW}" h="${extH}">
                ${pathXml}
              </a:path>
            </a:pathLst>
          </a:custGeom>
          ${fillXml}
          ${lnXml}
        </p:spPr>
      </p:sp>`);
  }

  function emitText(el: Element, t: TransformState): void {
    const lines = extractTextLines(el);
    if (lines.length === 0 || lines.every(l => l.trim().length === 0)) {
      return;
    }

    const tspans = Array.from(el.querySelectorAll('tspan'));
    const firstTspan = tspans[0];
    const rawX = parseFloat(firstTspan?.getAttribute('x') || el.getAttribute('x') || '0');
    const rawY = parseFloat(firstTspan?.getAttribute('y') || el.getAttribute('y') || '0');
    const x = rawX * t.sx + t.tx;
    const y = rawY * t.sy + t.ty;

    const rawFontSize = parseFloat(el.getAttribute('font-size') || '14') || 14;
    const scaledFontSizePx = rawFontSize * Math.abs(t.sy);

    // Handle initial dy offset on first tspan (e.g., "-0.5em", "1em", "0.35em")
    let firstDyPx = 0;
    const dyAttr = firstTspan?.getAttribute('dy') || el.getAttribute('dy') || '';
    if (dyAttr) {
      if (dyAttr.endsWith('em')) {
        firstDyPx = (parseFloat(dyAttr) || 0) * scaledFontSizePx;
      } else {
        firstDyPx = (parseFloat(dyAttr) || 0) * Math.abs(t.sy);
      }
    }

    const dominantBaseline = (el.getAttribute('dominant-baseline') || '').trim().toLowerCase();
    const textAnchor = (
      firstTspan?.getAttribute('text-anchor') ||
      el.getAttribute('text-anchor') ||
      'start'
    )
      .trim()
      .toLowerCase();

    const fontWeight = (el.getAttribute('font-weight') || '').trim().toLowerCase();
    const isBold = fontWeight === 'bold' || parseInt(fontWeight, 10) >= 600;
    const fontStyle = (el.getAttribute('font-style') || '').trim().toLowerCase();
    const isItalic = fontStyle === 'italic' || fontStyle === 'oblique';

    // Measure widest line in SVG pixels
    const fontFamilyAttr =
      (el.getAttribute('font-family') || 'Arial')
        .split(',')[0]
        .replace(/['"]/g, '')
        .trim() || 'Arial';
    let maxLineWidthPx = 20;
    for (const line of lines) {
      const w = measureTextWidthPx(line, scaledFontSizePx, fontFamilyAttr, isBold, isItalic);
      if (w > maxLineWidthPx) {
        maxLineWidthPx = w;
      }
    }

    // Add safety padding so PowerPoint never wraps a single line unintentionally
    const boxWidthPx = maxLineWidthPx + scaledFontSizePx * 1.2 + 12;
    const lineHeightPx = scaledFontSizePx * 1.22;
    const boxHeightPx = Math.max(lines.length * lineHeightPx + 6, scaledFontSizePx * 1.4);

    let boxX = x;
    let algn = 'l';
    if (textAnchor === 'middle' || textAnchor === 'center') {
      boxX = x - boxWidthPx / 2;
      algn = 'ctr';
    } else if (textAnchor === 'end') {
      boxX = x - boxWidthPx;
      algn = 'r';
    }

    // In SVG, y is the baseline of the first line (unless dominant-baseline="central"/"middle").
    // Top of the first line is approximately baselineY - 0.85 * fontSize.
    let firstBaselineY = y + firstDyPx;
    if (dominantBaseline === 'central' || dominantBaseline === 'middle') {
      firstBaselineY += scaledFontSizePx * 0.32;
    }
    const boxY = firstBaselineY - scaledFontSizePx * 0.85;

    const emuX = toEmuX(boxX);
    const emuY = toEmuY(boxY);
    const emuW = Math.max(toEmuLen(boxWidthPx), 91440);
    const emuH = Math.max(toEmuLen(boxHeightPx), 45720);

    const paragraphsXml = buildParagraphsXml(el, lines, t.sy, algn);

    const id = shapeIdCounter++;
    shapesXml.push(`
      <p:sp>
        <p:nvSpPr>
          <p:cNvPr id="${id}" name="TextBox ${id}"/>
          <p:cNvSpPr txBox="1"/>
          <p:nvPr/>
        </p:nvSpPr>
        <p:spPr>
          <a:xfrm>
            <a:off x="${emuX}" y="${emuY}"/>
            <a:ext cx="${emuW}" cy="${emuH}"/>
          </a:xfrm>
          <a:prstGeom prst="rect"><a:avLst/></a:prstGeom>
          <a:noFill/>
          <a:ln><a:noFill/></a:ln>
        </p:spPr>
        <p:txBody>
          <a:bodyPr wrap="none" lIns="0" tIns="0" rIns="0" bIns="0" anchor="t">
            <a:spAutoFit/>
          </a:bodyPr>
          <a:lstStyle/>
          ${paragraphsXml}
        </p:txBody>
      </p:sp>`);
  }

  function emitImage(el: Element, t: TransformState): void {
    const href = el.getAttribute('href') || el.getAttribute('xlink:href') || '';
    const prefix = 'data:image/png;base64,';
    if (!href.startsWith(prefix)) {
      return;
    }
    const base64Data = href.slice(prefix.length).trim();
    if (!base64Data) {
      return;
    }

    let bytes: Uint8Array;
    try {
      const binStr = atob(base64Data);
      bytes = new Uint8Array(binStr.length);
      for (let i = 0; i < binStr.length; i++) {
        bytes[i] = binStr.charCodeAt(i);
      }
    } catch {
      return;
    }

    const x = parseFloat(el.getAttribute('x') || '0') * t.sx + t.tx;
    const y = parseFloat(el.getAttribute('y') || '0') * t.sy + t.ty;
    const w = parseFloat(el.getAttribute('width') || '0') * Math.abs(t.sx);
    const h = parseFloat(el.getAttribute('height') || '0') * Math.abs(t.sy);
    if (w <= 0 || h <= 0) {
      return;
    }

    const imgIndex = mediaFiles.length + 1;
    const rId = `rIdImg${imgIndex}`;
    const fileName = `image${imgIndex}.png`;
    mediaFiles.push({ fileName, rId, data: bytes });

    const emuX = toEmuX(x);
    const emuY = toEmuY(y);
    const emuW = Math.max(toEmuLen(w), 1);
    const emuH = Math.max(toEmuLen(h), 1);
    const id = shapeIdCounter++;

    shapesXml.push(`
      <p:pic>
        <p:nvPicPr>
          <p:cNvPr id="${id}" name="Picture ${id}"/>
          <p:cNvPicPr><a:picLocks noChangeAspect="1"/></p:cNvPicPr>
          <p:nvPr/>
        </p:nvPicPr>
        <p:blipFill>
          <a:blip r:embed="${rId}"/>
          <a:stretch><a:fillRect/></a:stretch>
        </p:blipFill>
        <p:spPr>
          <a:xfrm>
            <a:off x="${emuX}" y="${emuY}"/>
            <a:ext cx="${emuW}" cy="${emuH}"/>
          </a:xfrm>
          <a:prstGeom prst="rect"><a:avLst/></a:prstGeom>
        </p:spPr>
      </p:pic>`);
  }

  // Walk the SVG DOM tree
  walkElement(svgEl, { tx: 0, ty: 0, sx: 1, sy: 1, rotDeg: 0 });

  // Build OpenXML package
  return buildPptxZip(shapesXml.join('\n'), SLIDE_W, SLIDE_H, mediaFiles);
}

function measureTextWidthPx(
  text: string,
  fontSizePx: number,
  fontFamily: string,
  isBold: boolean,
  isItalic: boolean
): number {
  if (typeof document !== 'undefined') {
    try {
      const canvas = document.createElement('canvas');
      const ctx = canvas.getContext('2d');
      if (ctx) {
        const style = `${isItalic ? 'italic ' : ''}${isBold ? 'bold ' : ''}${fontSizePx}px ${fontFamily}, Arial, sans-serif`;
        ctx.font = style;
        return ctx.measureText(text).width;
      }
    } catch {
      // Fallback below
    }
  }
  return text.length * fontSizePx * (isBold ? 0.65 : 0.58);
}

function approximateArc(
  x1: number,
  y1: number,
  rx: number,
  ry: number,
  phiDeg: number,
  largeArcFlag: number,
  sweepFlag: number,
  x2: number,
  y2: number
): { x: number; y: number }[] {
  if ((x1 === x2 && y1 === y2) || rx === 0 || ry === 0) {
    return [{ x: x2, y: y2 }];
  }
  rx = Math.abs(rx);
  ry = Math.abs(ry);
  const phi = (phiDeg * Math.PI) / 180;
  const cosPhi = Math.cos(phi);
  const sinPhi = Math.sin(phi);

  const dx2 = (x1 - x2) / 2;
  const dy2 = (y1 - y2) / 2;
  const x1p = cosPhi * dx2 + sinPhi * dy2;
  const y1p = -sinPhi * dx2 + cosPhi * dy2;

  const lambda = (x1p * x1p) / (rx * rx) + (y1p * y1p) / (ry * ry);
  if (lambda > 1) {
    const sqrtLambda = Math.sqrt(lambda);
    rx *= sqrtLambda;
    ry *= sqrtLambda;
  }

  const rx2 = rx * rx;
  const ry2 = ry * ry;
  const x1p2 = x1p * x1p;
  const y1p2 = y1p * y1p;

  const num = Math.max(0, rx2 * ry2 - rx2 * y1p2 - ry2 * x1p2);
  const den = rx2 * y1p2 + ry2 * x1p2;
  const sign = largeArcFlag === sweepFlag ? -1 : 1;
  const coef = den > 0 ? sign * Math.sqrt(num / den) : 0;

  const cxp = coef * ((rx * y1p) / ry);
  const cyp = coef * (-(ry * x1p) / rx);

  const cx = cosPhi * cxp - sinPhi * cyp + (x1 + x2) / 2;
  const cy = sinPhi * cxp + cosPhi * cyp + (y1 + y2) / 2;

  const angle = (ux: number, uy: number, vx: number, vy: number): number => {
    const dot = ux * vx + uy * vy;
    const len = Math.sqrt(ux * ux + uy * uy) * Math.sqrt(vx * vx + vy * vy);
    const cos = len > 0 ? Math.min(Math.max(dot / len, -1), 1) : 1;
    const a = Math.acos(cos);
    return ux * vy - uy * vx < 0 ? -a : a;
  };

  const theta1 = angle(1, 0, (x1p - cxp) / rx, (y1p - cyp) / ry);
  let dTheta = angle((x1p - cxp) / rx, (y1p - cyp) / ry, (-x1p - cxp) / rx, (-y1p - cyp) / ry);

  if (!sweepFlag && dTheta > 0) {
    dTheta -= 2 * Math.PI;
  } else if (sweepFlag && dTheta < 0) {
    dTheta += 2 * Math.PI;
  }

  const steps = Math.max(4, Math.ceil((Math.abs(dTheta) / (Math.PI / 2)) * 4));
  const pts: { x: number; y: number }[] = [];
  for (let i = 1; i <= steps; i++) {
    const theta = theta1 + (dTheta * i) / steps;
    const cosT = Math.cos(theta);
    const sinT = Math.sin(theta);
    pts.push({
      x: cosPhi * rx * cosT - sinPhi * ry * sinT + cx,
      y: sinPhi * rx * cosT + cosPhi * ry * sinT + cy
    });
  }
  return pts;
}

function parseOpacity(specificAttr: string | null, generalAttr: string | null): number {
  let op = 1;
  if (specificAttr !== null && specificAttr !== '') {
    const v = parseFloat(specificAttr);
    if (!isNaN(v)) {
      op *= v > 1 ? v / 100 : v;
    }
  }
  if (generalAttr !== null && generalAttr !== '') {
    const v = parseFloat(generalAttr);
    if (!isNaN(v)) {
      op *= v > 1 ? v / 100 : v;
    }
  }
  return Math.min(Math.max(op, 0), 1);
}

function toHex6(color: string): string | null {
  const c = color.trim().toLowerCase();
  if (!c || c === 'none' || c === 'transparent') {
    return null;
  }
  if (c.startsWith('#')) {
    const hex = c.slice(1);
    if (hex.length === 3) {
      return (hex[0] + hex[0] + hex[1] + hex[1] + hex[2] + hex[2]).toUpperCase();
    }
    if (hex.length >= 6) {
      return hex.slice(0, 6).toUpperCase();
    }
  }
  const rgbMatch = c.match(/^rgba?\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)/);
  if (rgbMatch) {
    const r = Math.min(255, parseInt(rgbMatch[1], 10)).toString(16).padStart(2, '0');
    const g = Math.min(255, parseInt(rgbMatch[2], 10)).toString(16).padStart(2, '0');
    const b = Math.min(255, parseInt(rgbMatch[3], 10)).toString(16).padStart(2, '0');
    return (r + g + b).toUpperCase();
  }
  return '000000';
}

function escapeXml(str: string): string {
  return str
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&apos;');
}

function buildPptxZip(
  shapesXml: string,
  slideW: number,
  slideH: number,
  mediaFiles: { fileName: string; rId: string; data: Uint8Array }[]
): Uint8Array {
  const contentTypesXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
  <Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
  <Default Extension="xml" ContentType="application/xml"/>
  <Default Extension="png" ContentType="image/png"/>
  <Override PartName="/ppt/presentation.xml" ContentType="application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"/>
  <Override PartName="/ppt/slideMasters/slideMaster1.xml" ContentType="application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"/>
  <Override PartName="/ppt/slideLayouts/slideLayout1.xml" ContentType="application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml"/>
  <Override PartName="/ppt/slides/slide1.xml" ContentType="application/vnd.openxmlformats-officedocument.presentationml.slide+xml"/>
  <Override PartName="/ppt/theme/theme1.xml" ContentType="application/vnd.openxmlformats-officedocument.theme+xml"/>
</Types>`;

  const rootRelsXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="ppt/presentation.xml"/>
</Relationships>`;

  const presentationXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:presentation xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main" saveSubsetFonts="1">
  <p:sldMasterIdLst>
    <p:sldMasterId id="2147483648" r:id="rId1"/>
  </p:sldMasterIdLst>
  <p:sldIdLst>
    <p:sldId id="256" r:id="rId2"/>
  </p:sldIdLst>
  <p:sldSz cx="${slideW}" cy="${slideH}" type="screen16x9"/>
  <p:notesSz cx="${slideH}" cy="${slideW}"/>
</p:presentation>`;

  const presentationRelsXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster" Target="slideMasters/slideMaster1.xml"/>
  <Relationship Id="rId2" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/slide" Target="slides/slide1.xml"/>
  <Relationship Id="rId3" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme" Target="theme/theme1.xml"/>
</Relationships>`;

  const slideMasterXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:sldMaster xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
  <p:cSld>
    <p:bg>
      <p:bgRef idx="1001">
        <a:schemeClr val="bg1"/>
      </p:bgRef>
    </p:bg>
    <p:spTree>
      <p:nvGrpSpPr>
        <p:cNvPr id="1" name=""/>
        <p:cNvGrpSpPr/>
        <p:nvPr/>
      </p:nvGrpSpPr>
      <p:grpSpPr>
        <a:xfrm>
          <a:off x="0" y="0"/>
          <a:ext cx="0" cy="0"/>
          <a:chOff x="0" y="0"/>
          <a:chExt cx="0" cy="0"/>
        </a:xfrm>
      </p:grpSpPr>
    </p:spTree>
  </p:cSld>
  <p:clrMap bg1="lt1" tx1="dk1" bg2="lt2" tx2="dk2" accent1="accent1" accent2="accent2" accent3="accent3" accent4="accent4" accent5="accent5" accent6="accent6" hlink="hlink" folHlink="folHlink"/>
  <p:sldLayoutIdLst>
    <p:sldLayoutId id="2147483649" r:id="rId1"/>
  </p:sldLayoutIdLst>
</p:sldMaster>`;

  const slideMasterRelsXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout" Target="../slideLayouts/slideLayout1.xml"/>
  <Relationship Id="rId2" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme" Target="../theme/theme1.xml"/>
</Relationships>`;

  const slideLayoutXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:sldLayout xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main" type="blank" preserve="1">
  <p:cSld name="Blank">
    <p:spTree>
      <p:nvGrpSpPr>
        <p:cNvPr id="1" name=""/>
        <p:cNvGrpSpPr/>
        <p:nvPr/>
      </p:nvGrpSpPr>
      <p:grpSpPr>
        <a:xfrm>
          <a:off x="0" y="0"/>
          <a:ext cx="0" cy="0"/>
          <a:chOff x="0" y="0"/>
          <a:chExt cx="0" cy="0"/>
        </a:xfrm>
      </p:grpSpPr>
    </p:spTree>
  </p:cSld>
  <p:clrMapOvr>
    <a:masterClrMapping/>
  </p:clrMapOvr>
</p:sldLayout>`;

  const slideLayoutRelsXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster" Target="../slideMasters/slideMaster1.xml"/>
</Relationships>`;

  const slideXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:sld xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
  <p:cSld>
    <p:spTree>
      <p:nvGrpSpPr>
        <p:cNvPr id="1" name=""/>
        <p:cNvGrpSpPr/>
        <p:nvPr/>
      </p:nvGrpSpPr>
      <p:grpSpPr>
        <a:xfrm>
          <a:off x="0" y="0"/>
          <a:ext cx="0" cy="0"/>
          <a:chOff x="0" y="0"/>
          <a:chExt cx="0" cy="0"/>
        </a:xfrm>
      </p:grpSpPr>
      ${shapesXml}
    </p:spTree>
  </p:cSld>
  <p:clrMapOvr>
    <a:masterClrMapping/>
  </p:clrMapOvr>
</p:sld>`;

  const imgRels = mediaFiles
    .map(
      m =>
        `  <Relationship Id="${m.rId}" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/image" Target="../media/${m.fileName}"/>`
    )
    .join('\n');

  const slideRelsXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout" Target="../slideLayouts/slideLayout1.xml"/>
${imgRels}
</Relationships>`;

  const themeXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<a:theme xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" name="Office Theme">
  <a:themeElements>
    <a:clrScheme name="Office">
      <a:dk1><a:sysClr val="windowText" lastClr="000000"/></a:dk1>
      <a:lt1><a:sysClr val="window" lastClr="FFFFFF"/></a:lt1>
      <a:dk2><a:srgbClr val="1F497D"/></a:dk2>
      <a:lt2><a:srgbClr val="EEECE1"/></a:lt2>
      <a:accent1><a:srgbClr val="4F81BD"/></a:accent1>
      <a:accent2><a:srgbClr val="C0504D"/></a:accent2>
      <a:accent3><a:srgbClr val="9BBB59"/></a:accent3>
      <a:accent4><a:srgbClr val="8064A2"/></a:accent4>
      <a:accent5><a:srgbClr val="4BACC6"/></a:accent5>
      <a:accent6><a:srgbClr val="F79646"/></a:accent6>
      <a:hlink><a:srgbClr val="0000FF"/></a:hlink>
      <a:folHlink><a:srgbClr val="800080"/></a:folHlink>
    </a:clrScheme>
    <a:fontScheme name="Office">
      <a:majorFont>
        <a:latin typeface="Arial"/>
        <a:ea typeface=""/>
        <a:cs typeface=""/>
      </a:majorFont>
      <a:minorFont>
        <a:latin typeface="Arial"/>
        <a:ea typeface=""/>
        <a:cs typeface=""/>
      </a:minorFont>
    </a:fontScheme>
    <a:fmtScheme name="Office">
      <a:fillStyleLst>
        <a:solidFill><a:schemeClr val="phClr"/></a:solidFill>
        <a:solidFill><a:schemeClr val="phClr"/></a:solidFill>
        <a:solidFill><a:schemeClr val="phClr"/></a:solidFill>
      </a:fillStyleLst>
      <a:lnStyleLst>
        <a:ln w="9525"><a:solidFill><a:schemeClr val="phClr"/></a:solidFill></a:ln>
        <a:ln w="25400"><a:solidFill><a:schemeClr val="phClr"/></a:solidFill></a:ln>
        <a:ln w="38100"><a:solidFill><a:schemeClr val="phClr"/></a:solidFill></a:ln>
      </a:lnStyleLst>
      <a:effectStyleLst>
        <a:effectStyle><a:effectLst/></a:effectStyle>
        <a:effectStyle><a:effectLst/></a:effectStyle>
        <a:effectStyle><a:effectLst/></a:effectStyle>
      </a:effectStyleLst>
      <a:bgFillStyleLst>
        <a:solidFill><a:schemeClr val="phClr"/></a:solidFill>
        <a:solidFill><a:schemeClr val="phClr"/></a:solidFill>
        <a:solidFill><a:schemeClr val="phClr"/></a:solidFill>
      </a:bgFillStyleLst>
    </a:fmtScheme>
  </a:themeElements>
  <a:objectDefaults/>
  <a:extraClrSchemeLst/>
</a:theme>`;

  const encoder = new TextEncoder();
  const entries: { name: string; data: Uint8Array }[] = [
    { name: '[Content_Types].xml', data: encoder.encode(contentTypesXml) },
    { name: '_rels/.rels', data: encoder.encode(rootRelsXml) },
    { name: 'ppt/presentation.xml', data: encoder.encode(presentationXml) },
    { name: 'ppt/_rels/presentation.xml.rels', data: encoder.encode(presentationRelsXml) },
    { name: 'ppt/slideMasters/slideMaster1.xml', data: encoder.encode(slideMasterXml) },
    { name: 'ppt/slideMasters/_rels/slideMaster1.xml.rels', data: encoder.encode(slideMasterRelsXml) },
    { name: 'ppt/slideLayouts/slideLayout1.xml', data: encoder.encode(slideLayoutXml) },
    { name: 'ppt/slideLayouts/_rels/slideLayout1.xml.rels', data: encoder.encode(slideLayoutRelsXml) },
    { name: 'ppt/slides/slide1.xml', data: encoder.encode(slideXml) },
    { name: 'ppt/slides/_rels/slide1.xml.rels', data: encoder.encode(slideRelsXml) },
    { name: 'ppt/theme/theme1.xml', data: encoder.encode(themeXml) }
  ];

  for (const m of mediaFiles) {
    entries.push({ name: `ppt/media/${m.fileName}`, data: m.data });
  }

  return createZipArchive(entries);
}

const CRC32_TABLE = (() => {
  const table = new Uint32Array(256);
  for (let i = 0; i < 256; i++) {
    let c = i;
    for (let k = 0; k < 8; k++) {
      c = c & 1 ? 0xedb88320 ^ (c >>> 1) : c >>> 1;
    }
    table[i] = c >>> 0;
  }
  return table;
})();

function crc32(data: Uint8Array): number {
  let crc = 0xffffffff;
  for (let i = 0; i < data.length; i++) {
    crc = CRC32_TABLE[(crc ^ data[i]) & 0xff] ^ (crc >>> 8);
  }
  return (crc ^ 0xffffffff) >>> 0;
}

function createZipArchive(entries: { name: string; data: Uint8Array }[]): Uint8Array {
  const encoder = new TextEncoder();
  const localHeaders: Uint8Array[] = [];
  const centralHeaders: Uint8Array[] = [];
  let offset = 0;

  for (const entry of entries) {
    const nameBytes = encoder.encode(entry.name);
    const data = entry.data;
    const crc = crc32(data);
    const size = data.length;

    // Local file header (30 bytes + filename)
    const localHeader = new Uint8Array(30 + nameBytes.length);
    const lv = new DataView(localHeader.buffer);
    lv.setUint32(0, 0x04034b50, true); // signature
    lv.setUint16(4, 20, true); // version needed to extract (2.0)
    lv.setUint16(6, 0, true); // general purpose bit flag
    lv.setUint16(8, 0, true); // compression method: 0 = STORED
    lv.setUint16(10, 0, true); // last mod file time
    lv.setUint16(12, 0x21, true); // last mod file date (1980-01-01)
    lv.setUint32(14, crc, true); // crc-32
    lv.setUint32(18, size, true); // compressed size
    lv.setUint32(22, size, true); // uncompressed size
    lv.setUint16(26, nameBytes.length, true); // file name length
    lv.setUint16(28, 0, true); // extra field length
    localHeader.set(nameBytes, 30);

    // Central directory file header (46 bytes + filename)
    const centralHeader = new Uint8Array(46 + nameBytes.length);
    const cv = new DataView(centralHeader.buffer);
    cv.setUint32(0, 0x02014b50, true); // signature
    cv.setUint16(4, 20, true); // version made by
    cv.setUint16(6, 20, true); // version needed to extract
    cv.setUint16(8, 0, true); // general purpose bit flag
    cv.setUint16(10, 0, true); // compression method: 0 = STORED
    cv.setUint16(12, 0, true); // last mod file time
    cv.setUint16(14, 0x21, true); // last mod file date
    cv.setUint32(16, crc, true); // crc-32
    cv.setUint32(20, size, true); // compressed size
    cv.setUint32(24, size, true); // uncompressed size
    cv.setUint16(28, nameBytes.length, true); // file name length
    cv.setUint16(30, 0, true); // extra field length
    cv.setUint16(32, 0, true); // file comment length
    cv.setUint16(34, 0, true); // disk number start
    cv.setUint16(36, 0, true); // internal file attributes
    cv.setUint32(38, 0, true); // external file attributes
    cv.setUint32(42, offset, true); // relative offset of local header
    centralHeader.set(nameBytes, 46);

    localHeaders.push(localHeader, data);
    centralHeaders.push(centralHeader);
    offset += localHeader.length + data.length;
  }

  let centralSize = 0;
  for (const ch of centralHeaders) {
    centralSize += ch.length;
  }

  // End of central directory record (22 bytes)
  const eocd = new Uint8Array(22);
  const ev = new DataView(eocd.buffer);
  ev.setUint32(0, 0x06054b50, true); // signature
  ev.setUint16(4, 0, true); // disk number
  ev.setUint16(6, 0, true); // disk where central directory starts
  ev.setUint16(8, entries.length, true); // number of central directory records on this disk
  ev.setUint16(10, entries.length, true); // total number of central directory records
  ev.setUint32(12, centralSize, true); // size of central directory
  ev.setUint32(16, offset, true); // offset of start of central directory
  ev.setUint16(20, 0, true); // comment length

  const totalSize = offset + centralSize + eocd.length;
  const out = new Uint8Array(totalSize);
  let pos = 0;
  for (const chunk of localHeaders) {
    out.set(chunk, pos);
    pos += chunk.length;
  }
  for (const chunk of centralHeaders) {
    out.set(chunk, pos);
    pos += chunk.length;
  }
  out.set(eocd, pos);
  return out;
}
