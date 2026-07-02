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

  // Convert back to string
  return new XMLSerializer().serializeToString(svg);
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
