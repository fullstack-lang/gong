package exporter

import "hash/fnv"

// ConvertStringToUniqueNumbers is the core function that hashes a string
// to generate a unique int64 and float64.
func ConvertStringToUniqueNumbers(s string) (int64, float64) {
	// Create a new 64-bit FNV-1a hash object.
	h := fnv.New64a()

	// Write the byte representation of the string to the hasher.
	h.Write([]byte(s))

	// Get the resulting 64-bit hash value as a uint64.
	hashValue := h.Sum64()

	// Convert the uint64 to an int64.
	intValue := int64(hashValue)

	// Convert the int64 to a float64.
	floatValue := float64(intValue)

	return intValue, floatValue
}

// ------------------------------------------------------------------
// NEW: Direct function to get only the int64
// ------------------------------------------------------------------
func HashStringToInt(s string) int {
	// We use the blank identifier (_) to ignore the float64 return value.
	intValue, _ := ConvertStringToUniqueNumbers(s)
	return int(intValue)
}

// ------------------------------------------------------------------
// NEW: Direct function to get only the float64
// ------------------------------------------------------------------
func HashStringToFloat64(s string) float64 {
	// We use the blank identifier (_) to ignore the int64 return value.
	_, floatValue := ConvertStringToUniqueNumbers(s)
	return floatValue
}
