package models

func computeIsBounded(particle Particle) (res bool) {

	// if element, ok := particle.(*Element); ok {
	// 	log.Println("element", element.Name)
	// }
	switch v := particle.(type) {
	case OccurrenceDefinition:
		if v.GetIsLocalyUnbounded() {
			return true
		} else {
			// check if outer particle is unbounded recursively
			if parent := particle.GetParent(); parent != nil {
				return computeIsBounded(parent)
			}
		}
	}

	return
}
