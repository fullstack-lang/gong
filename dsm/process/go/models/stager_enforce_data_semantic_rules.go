package models

import (
	"fmt"
	"slices"
	"time"
)

// collectReachableDatas returns the set of all non-nil Data instances reachable
// from roots via SubDatas (cycle-safe).
func collectReachableDatas(roots []*Data) map[*Data]bool {
	reachable := make(map[*Data]bool)
	var visit func(d *Data)
	visit = func(d *Data) {
		if d == nil || reachable[d] {
			return
		}
		reachable[d] = true
		for _, sub := range d.SubDatas {
			visit(sub)
		}
	}
	for _, root := range roots {
		visit(root)
	}
	return reachable
}

// isDataReachableFromDataFlow returns true if target is in dataFlow.Datas or is
// a descendant of any Data in dataFlow.Datas via SubDatas.
func isDataReachableFromDataFlow(dataFlow *DataFlow, target *Data) bool {
	if dataFlow == nil || target == nil {
		return false
	}
	visited := make(map[*Data]bool)
	var visit func(d *Data) bool
	visit = func(d *Data) bool {
		if d == nil || visited[d] {
			return false
		}
		if d == target {
			return true
		}
		visited[d] = true
		for _, sub := range d.SubDatas {
			if visit(sub) {
				return true
			}
		}
		return false
	}
	for _, root := range dataFlow.Datas {
		if visit(root) {
			return true
		}
	}
	return false
}

// topLevelDatasInSlice filters datas to preserve order while omitting elements
// whose ancestor (via parentData) is already present in datas.
func topLevelDatasInSlice(datas []*Data) []*Data {
	present := make(map[*Data]bool, len(datas))
	for _, d := range datas {
		if d != nil {
			present[d] = true
		}
	}
	seen := make(map[*Data]bool, len(datas))
	var res []*Data
	for _, d := range datas {
		if d == nil || seen[d] {
			continue
		}
		hasAnc := false
		visited := make(map[*Data]bool)
		visited[d] = true
		for anc := d.parentData; anc != nil && !visited[anc]; anc = anc.parentData {
			visited[anc] = true
			if present[anc] {
				hasAnc = true
				break
			}
		}
		if !hasAnc {
			seen[d] = true
			res = append(res, d)
		}
	}
	return res
}

// walkDatasPreOrder traverses roots and their SubDatas in pre-order DFS.
// Roots that are already descendants of another root in roots are visited under
// their ancestor so parents always precede descendants.
func walkDatasPreOrder(roots []*Data, fn func(d *Data)) {
	rootSet := make(map[*Data]bool, len(roots))
	for _, r := range roots {
		if r != nil {
			rootSet[r] = true
		}
	}

	hasAncestorInRoots := func(d *Data) bool {
		seen := make(map[*Data]bool)
		seen[d] = true
		for anc := d.parentData; anc != nil && !seen[anc]; anc = anc.parentData {
			seen[anc] = true
			if rootSet[anc] {
				return true
			}
		}
		return false
	}

	visited := make(map[*Data]bool)
	var visit func(d *Data)
	visit = func(d *Data) {
		if d == nil || visited[d] {
			return
		}
		visited[d] = true
		fn(d)
		for _, sub := range d.SubDatas {
			visit(sub)
		}
	}

	for _, root := range roots {
		if root == nil || hasAncestorInRoots(root) {
			continue
		}
		visit(root)
	}
	// Fallback pass in case of a cycle or uncomputed parentData
	for _, root := range roots {
		visit(root)
	}
}

// computeDataFields computes non-persisted parentData pointers and propagates
// owningLibrary from Library.RootDatas down through SubDatas.
func (stager *Stager) computeDataFields() {
	stage := stager.stage

	datas := stage.GetInstancesSorted[*Data]()
	for _, data := range datas {
		data.parentData = nil
	}
	for _, data := range datas {
		for _, sub := range data.SubDatas {
			if sub != nil && sub != data && sub.parentData == nil {
				sub.parentData = data
			}
		}
	}

	for _, library := range stage.GetInstancesSorted[*Library]() {
		visited := make(map[*Data]bool)
		var setLib func(d *Data)
		setLib = func(d *Data) {
			if d == nil || visited[d] {
				return
			}
			visited[d] = true
			d.SetOwningLibrary(library)
			for _, sub := range d.SubDatas {
				setLib(sub)
			}
		}
		for _, rootData := range library.RootDatas {
			setLib(rootData)
		}
	}
}

// enforceDataSemanticRules enforces strict tree (forest) invariants on Data.SubDatas:
// 1. No nil entries, self-references, or duplicate children in SubDatas.
// 2. Each Data has at most one parent Data.
// 3. No cycles in SubDatas.
// 4. A Data that is a sub-data of another Data is removed from Library.RootDatas
//    (and duplicate entries across Library.RootDatas are cleaned up).
func (stager *Stager) enforceDataSemanticRules() (needCommit bool) {
	stage := stager.stage
	datas := stage.GetInstancesSorted[*Data]()
	libraries := stage.GetInstancesSorted[*Library]()

	// 1. Clean nil, self-references, and duplicate children in SubDatas
	for _, data := range datas {
		seen := make(map[*Data]bool, len(data.SubDatas))
		var cleaned []*Data
		modified := false
		for _, sub := range data.SubDatas {
			if sub == nil {
				modified = true
				continue
			}
			if sub == data {
				modified = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Data \"%s\" cannot be a sub-data of itself, link broken", data.GetName()))
				}
				continue
			}
			if seen[sub] {
				modified = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Removed duplicate sub-data \"%s\" from Data \"%s\"", sub.GetName(), data.GetName()))
				}
				continue
			}
			seen[sub] = true
			cleaned = append(cleaned, sub)
		}
		if modified {
			data.SubDatas = cleaned
			needCommit = true
		}
	}

	// 2. Ensure at most one parent per Data
	parentMap := make(map[*Data]*Data)
	for _, data := range datas {
		childrenCopy := slices.Clone(data.SubDatas)
		for _, child := range childrenCopy {
			if existingParent, ok := parentMap[child]; ok {
				if idx := slices.Index(data.SubDatas, child); idx != -1 {
					data.SubDatas = slices.Delete(data.SubDatas, idx, idx+1)
				}
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Data \"%s\" has 2 parents: \"%s\" and \"%s\". Link from the latter is deleted",
							child.GetName(), existingParent.GetName(), data.GetName()))
				}
				needCommit = true
			} else {
				parentMap[child] = data
			}
		}
	}

	// 3. Ensure no cycles in SubDatas (3-color DFS)
	whiteSet := make(map[*Data]struct{}, len(datas))
	graySet := make(map[*Data]struct{})
	blackSet := make(map[*Data]struct{})
	for _, d := range datas {
		whiteSet[d] = struct{}{}
	}

	var dfs func(node *Data)
	dfs = func(node *Data) {
		delete(whiteSet, node)
		graySet[node] = struct{}{}

		childrenCopy := slices.Clone(node.SubDatas)
		for _, child := range childrenCopy {
			if _, ok := graySet[child]; ok {
				if idx := slices.Index(node.SubDatas, child); idx != -1 {
					node.SubDatas = slices.Delete(node.SubDatas, idx, idx+1)
				}
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Found loop involving Data \"%s\", breaking edge with \"%s\"",
							node.GetName(), child.GetName()))
				}
				needCommit = true
				continue
			}
			if _, ok := blackSet[child]; ok {
				continue
			}
			if _, ok := whiteSet[child]; ok {
				dfs(child)
			}
		}

		delete(graySet, node)
		blackSet[node] = struct{}{}
	}
	// Start DFS from existing Library.RootDatas first so back-edges pointing
	// to a root or ancestor are broken rather than established parent->child edges.
	for _, lib := range libraries {
		for _, d := range lib.RootDatas {
			if _, ok := whiteSet[d]; ok {
				dfs(d)
			}
		}
	}
	for _, d := range datas {
		if _, ok := whiteSet[d]; ok {
			dfs(d)
		}
	}

	// 4. Remove sub-datas (and duplicates) from Library.RootDatas
	isSubData := make(map[*Data]bool)
	for _, data := range datas {
		for _, sub := range data.SubDatas {
			if sub != nil {
				isSubData[sub] = true
			}
		}
	}
	seenRootData := make(map[*Data]bool)
	for _, lib := range libraries {
		var filtered []*Data
		modified := false
		for _, d := range lib.RootDatas {
			if d == nil {
				modified = true
				continue
			}
			if isSubData[d] {
				modified = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Data \"%s\" is a sub-data, removing it from library \"%s\" RootDatas", d.GetName(), lib.GetName()))
				}
				continue
			}
			if seenRootData[d] {
				modified = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Data \"%s\" is duplicated in RootDatas, removing extra entry from library \"%s\"", d.GetName(), lib.GetName()))
				}
				continue
			}
			seenRootData[d] = true
			filtered = append(filtered, d)
		}
		if modified {
			lib.RootDatas = filtered
			needCommit = true
		}
	}

	stager.computeDataFields()

	return
}
