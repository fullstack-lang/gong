package models

// GongDiff checks for differences between the 'a' instance and the 'ref' instance.
// It returns a slice of strings containing the names of the fields that have changed.
func (a *A) GongDiff(ref *A) (res []string) {

	if a.Name != ref.Name {
		res = append(res, "Name")
	}

	if len(a.Bs) != len(ref.Bs) {
		res = append(res, "Bs")
	} else {
		for i, _b := range a.Bs {
			if _b != ref.Bs[i] {
				res = append(res, "Bs")
				break
			}
		}
	}

	return
}

// GongDiff checks for differences between the 'b' instance and the 'ref' instance.
// It returns a slice of strings containing the names of the fields that have changed.
func (b *B) GongDiff(ref *B) (res []string) {

	if b.Name != ref.Name {
		res = append(res, "Name")
	}

	return
}
