package models

// GetEffectiveComplexities returns all complexities for this system.
// If AreCPEsCompoundedFromSubSystems is true, it includes complexities from its subsystems recursively.
func (s *System) GetEffectiveComplexities() ([]*Complexity, map[*Complexity]*System) {
	complexities := make([]*Complexity, 0)
	sysMap := make(map[*Complexity]*System)
	visited := make(map[*System]bool)
	s.collectComplexities(&complexities, sysMap, visited)
	return complexities, sysMap
}

func (s *System) collectComplexities(list *[]*Complexity, sysMap map[*Complexity]*System, visited map[*System]bool) {
	if s == nil || visited[s] {
		return
	}
	visited[s] = true

	if s.AreCPEsCompoundedFromSubSystems && len(s.SubSystemes) > 0 {
		for _, sub := range s.SubSystemes {
			sub.collectComplexities(list, sysMap, visited)
		}
	} else {
		for _, c := range s.Complexities {
			*list = append(*list, c)
			sysMap[c] = s
		}
	}
}

// GetEffectivePerformances returns all performances for this system.
// If AreCPEsCompoundedFromSubSystems is true, it includes performances from its subsystems recursively.
func (s *System) GetEffectivePerformances() ([]*Performance, map[*Performance]*System) {
	performances := make([]*Performance, 0)
	sysMap := make(map[*Performance]*System)
	visited := make(map[*System]bool)
	s.collectPerformances(&performances, sysMap, visited)
	return performances, sysMap
}

func (s *System) collectPerformances(list *[]*Performance, sysMap map[*Performance]*System, visited map[*System]bool) {
	if s == nil || visited[s] {
		return
	}
	visited[s] = true

	if s.AreCPEsCompoundedFromSubSystems && len(s.SubSystemes) > 0 {
		for _, sub := range s.SubSystemes {
			sub.collectPerformances(list, sysMap, visited)
		}
	} else {
		for _, p := range s.Performances {
			*list = append(*list, p)
			sysMap[p] = s
		}
	}
}

// GetEffectiveEfforts returns all efforts for this system.
// If AreCPEsCompoundedFromSubSystems is true, it includes efforts from its subsystems recursively.
func (s *System) GetEffectiveEfforts() ([]*Effort, map[*Effort]*System) {
	efforts := make([]*Effort, 0)
	sysMap := make(map[*Effort]*System)
	visited := make(map[*System]bool)
	s.collectEfforts(&efforts, sysMap, visited)
	return efforts, sysMap
}

func (s *System) collectEfforts(list *[]*Effort, sysMap map[*Effort]*System, visited map[*System]bool) {
	if s == nil || visited[s] {
		return
	}
	visited[s] = true

	if s.AreCPEsCompoundedFromSubSystems && len(s.SubSystemes) > 0 {
		for _, sub := range s.SubSystemes {
			sub.collectEfforts(list, sysMap, visited)
		}
	} else {
		for _, e := range s.Efforts {
			*list = append(*list, e)
			sysMap[e] = s
		}
	}
}
