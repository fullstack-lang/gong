package diagrammer

func IndexOf(slice []PortfolioNode, value PortfolioNode) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1 // Not found
}

func Remove(slice []PortfolioNode, index int) []PortfolioNode {
	// Perform a bounds check to avoid runtime panic
	if index < 0 || index >= len(slice) {
		return slice // Optionally, return an error or handle this case as needed
	}
	return append(slice[:index], slice[index+1:]...)
}
