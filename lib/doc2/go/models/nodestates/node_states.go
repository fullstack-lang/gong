// nodestates/nodestates.go
package nodestates

import (
	"encoding/json"
	"fmt"
)

// internalUnmarshal converts a JSON string representation of boolean states
// into a slice of booleans. It treats an empty string or "null" as an empty slice.
func internalUnmarshal(nodeStatesJSON string) ([]bool, error) {
	if nodeStatesJSON == "" || nodeStatesJSON == "null" {
		return []bool{}, nil // Represent empty string or "null" as an empty slice
	}
	var states []bool
	err := json.Unmarshal([]byte(nodeStatesJSON), &states)
	if err != nil {
		return nil, fmt.Errorf("nodestates: failed to unmarshal from JSON string '%s': %w", nodeStatesJSON, err)
	}
	// If unmarshal results in a nil slice (e.g. from "null" JSON), ensure it's an empty slice
	if states == nil {
		return []bool{}, nil
	}
	return states, nil
}

// internalMarshal converts a slice of booleans into a JSON string representation.
func internalMarshal(states []bool) (string, error) {
	// Ensure that a nil slice is marshaled as an empty JSON array "[]"
	// rather than "null", to maintain consistency.
	if states == nil {
		states = []bool{}
	}
	jsonData, err := json.Marshal(states)
	if err != nil {
		return "", fmt.Errorf("nodestates: failed to marshal to JSON string: %w", err)
	}
	return string(jsonData), nil
}

// IsNodeExpanded checks if the node at the given index is true (expanded)
// based on a JSON string representation of node states.
// The nodeStatesJSON string can be empty (""), which is treated as an empty list of states.
// If the index is out of bounds for the deserialized states, it returns false.
func IsNodeExpanded(nodeStatesJSON string, index int) (bool, error) {
	if index < 0 {
		return false, fmt.Errorf("nodestates: index cannot be negative, got %d", index)
	}

	states, err := internalUnmarshal(nodeStatesJSON)
	if err != nil {
		return false, err // Error already contains context
	}

	if index >= len(states) {
		return false, nil // Index is out of bounds, so node is considered not expanded
	}
	return states[index], nil
}

// ToggleNodeExpanded changes the state of the node at the given index.
// The nodeStatesJSON is a pointer to a string containing the JSON array of booleans.
// If the string is empty (""), it's treated as an empty list.
// If the index is out of bounds, the list of states is automatically expanded
// with 'false' values up to the required index before toggling.
// The string pointed to by nodeStatesJSON is updated with the new JSON state.
func ToggleNodeExpanded(nodeStatesJSON *string, index int) error {
	if nodeStatesJSON == nil {
		return fmt.Errorf("nodestates: nodeStatesJSON pointer cannot be nil")
	}
	if index < 0 {
		return fmt.Errorf("nodestates: index cannot be negative, got %d", index)
	}

	states, err := internalUnmarshal(*nodeStatesJSON)
	if err != nil {
		return err // Error already contains context
	}

	// Expand the slice if the index is out of bounds
	if index >= len(states) {
		// Create a new slice with enough capacity. New elements will be 'false'.
		newStates := make([]bool, index+1)
		copy(newStates, states) // Copy existing elements
		states = newStates
	}

	// Toggle the state at the given index
	states[index] = !states[index]

	// Marshal back to JSON and update the original string
	newJSON, err := internalMarshal(states)
	if err != nil {
		return err // Error already contains context
	}
	*nodeStatesJSON = newJSON
	return nil
}
