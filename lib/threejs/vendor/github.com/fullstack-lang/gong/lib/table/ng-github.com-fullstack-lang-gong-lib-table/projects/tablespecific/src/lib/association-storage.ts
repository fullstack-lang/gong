/**
 * Encodes an array of integers into a JSON string.
 *
 * @param data - The array of integers to encode.
 * @returns A JSON string representation of the integer array.
 */
export function encodeIntArrayToString_json(data: number[]): string {
  return JSON.stringify(data);
}

/**
 * Decodes a JSON string (previously encoded with encodeIntArrayToString_json)
 * back into an array of integers.
 *
 * @param str - The JSON string to decode.
 * @returns An array of integers.
 * @throws Error if the string is not valid JSON or does not represent an array of numbers.
 */
export function decodeStringToIntArray_json(str: string): number[] {
  try {
    const parsedData = JSON.parse(str);
    if (Array.isArray(parsedData) && parsedData.every(item => typeof item === 'number')) {
      return parsedData as number[];
    } else {
      // Or handle more gracefully, e.g., return [] or throw a custom error
      console.error("Decoded data is not an array of numbers:", parsedData);
      return [];
    }
  } catch (error) {
    console.error("Failed to decode JSON string:", error);
    // Or re-throw, or return [], depending on desired error handling
    return [];
  }
}