package runtime

import (
	"errors"
	"log"
	"sync"
)

// StageCore holds common state and operations for any Gong Stage
type StageCore struct {
	name string

	// isInDeltaMode is true when the stage is used to compute difference between successive commits
	isInDeltaMode bool

	// gongMarshallingMode sets the marshalling mode
	gongMarshallingMode GongMarshallingMode

	// isWithGenesisCommit indicates if the genesis commit (#0) cannot be rolled back
	isWithGenesisCommit bool

	forwardCommits  []string
	backwardCommits []string

	// navigationMode is set when navigating commit history
	navigationMode StageNavigationMode
	commitsBehind  int

	isApplyingBackwardCommit bool
	isApplyingForwardCommit  bool
	isSquashing              bool
	modified                 bool

	beforeCommitHooks []func()
	afterCommitHooks  []func()

	lock sync.RWMutex
}

func (sc *StageCore) SetName(name string) {
	sc.name = name
}

func (sc *StageCore) GetName() string {
	return sc.name
}

func (sc *StageCore) SetGongMarshallingMode(mode GongMarshallingMode) {
	sc.gongMarshallingMode = mode
}

func (sc *StageCore) GetGongMarshallingMode() GongMarshallingMode {
	return sc.gongMarshallingMode
}

func (sc *StageCore) SetIsWithGenesisCommit(isWithGenesisCommit bool) {
	sc.isWithGenesisCommit = isWithGenesisCommit
}

func (sc *StageCore) GetIsWithGenesisCommit() bool {
	return sc.isWithGenesisCommit
}

func (sc *StageCore) SetDeltaMode(inDeltaMode bool) {
	sc.isInDeltaMode = inDeltaMode
}

func (sc *StageCore) IsInDeltaMode() bool {
	return sc.isInDeltaMode
}

func (sc *StageCore) Lock() {
	sc.lock.Lock()
}

func (sc *StageCore) Unlock() {
	sc.lock.Unlock()
}

func (sc *StageCore) RLock() {
	sc.lock.RLock()
}

func (sc *StageCore) RUnlock() {
	sc.lock.RUnlock()
}

func (sc *StageCore) RegisterBeforeCommitHook(hook func()) {
	sc.beforeCommitHooks = append(sc.beforeCommitHooks, hook)
}

func (sc *StageCore) RegisterAfterCommitHook(hook func()) {
	sc.afterCommitHooks = append(sc.afterCommitHooks, hook)
}

func (sc *StageCore) GetBeforeCommitHooks() []func() {
	return sc.beforeCommitHooks
}

func (sc *StageCore) SetBeforeCommitHooks(hooks []func()) {
	sc.beforeCommitHooks = hooks
}

func (sc *StageCore) ClearBeforeCommitHooks() {
	sc.beforeCommitHooks = nil
}

func (sc *StageCore) GetAfterCommitHooks() []func() {
	return sc.afterCommitHooks
}

func (sc *StageCore) SetAfterCommitHooks(hooks []func()) {
	sc.afterCommitHooks = hooks
}

func (sc *StageCore) ClearAfterCommitHooks() {
	sc.afterCommitHooks = nil
}

func (sc *StageCore) RunBeforeCommitHooks() {
	for _, hook := range sc.beforeCommitHooks {
		hook()
	}
}

func (sc *StageCore) RunAfterCommitHooks() {
	for _, hook := range sc.afterCommitHooks {
		hook()
	}
}

func (sc *StageCore) GetForwardCommits() []string {
	return sc.forwardCommits
}

func (sc *StageCore) GetBackwardCommits() []string {
	return sc.backwardCommits
}

func (sc *StageCore) SetForwardCommits(commits []string) {
	sc.forwardCommits = commits
}

func (sc *StageCore) SetBackwardCommits(commits []string) {
	sc.backwardCommits = commits
}

func (sc *StageCore) AppendForwardCommit(commit string) {
	sc.forwardCommits = append(sc.forwardCommits, commit)
}

func (sc *StageCore) AppendBackwardCommit(commit string) {
	sc.backwardCommits = append(sc.backwardCommits, commit)
}

func (sc *StageCore) GetCommitsBehind() int {
	return sc.commitsBehind
}

func (sc *StageCore) SetCommitsBehind(commitsBehind int) {
	sc.commitsBehind = commitsBehind
}

func (sc *StageCore) GetNavigationMode() StageNavigationMode {
	return sc.navigationMode
}

func (sc *StageCore) SetNavigationMode(mode StageNavigationMode) {
	sc.navigationMode = mode
}

func (sc *StageCore) IsApplyingBackwardCommit() bool {
	return sc.isApplyingBackwardCommit
}

func (sc *StageCore) IsApplyingForwardCommit() bool {
	return sc.isApplyingForwardCommit
}

func (sc *StageCore) IsSquashing() bool {
	return sc.isSquashing
}

func (sc *StageCore) SetModified(modified bool) {
	sc.modified = modified
}

func (sc *StageCore) IsModified() bool {
	return sc.modified
}

// ApplyBackwardCommit applies the commit before the current one
func (sc *StageCore) ApplyBackwardCommit(parseAstString func(string, bool) error, computeRefAndOrders func()) error {
	if len(sc.backwardCommits) == 0 {
		return errors.New("no backward commit to apply")
	}

	if sc.navigationMode == GongNavigationModeNormal && sc.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if sc.navigationMode == GongNavigationModeNormal {
		sc.navigationMode = GongNavigationModeNavigating
	}

	if sc.isWithGenesisCommit && sc.commitsBehind >= len(sc.backwardCommits)-1 {
		return errors.New("cannot rollback genesis commit")
	}

	if sc.commitsBehind >= len(sc.backwardCommits) {
		return errors.New("no more backward commit to apply")
	}

	commitToApply := sc.backwardCommits[len(sc.backwardCommits)-1-sc.commitsBehind]

	sc.commitsBehind++
	sc.isApplyingBackwardCommit = true
	err := parseAstString(commitToApply, true)
	sc.isApplyingBackwardCommit = false
	if err != nil {
		log.Println("error during ApplyBackwardCommit: ", err)
		return err
	}

	if computeRefAndOrders != nil {
		computeRefAndOrders()
	}

	return nil
}

// ApplyForwardCommit applies the commit after the current one
func (sc *StageCore) ApplyForwardCommit(parseAstString func(string, bool) error, computeRefAndOrders func()) error {
	if sc.navigationMode == GongNavigationModeNormal && sc.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if sc.commitsBehind == 0 {
		return errors.New("no more forward commit to apply")
	}

	if sc.navigationMode == GongNavigationModeNormal {
		sc.navigationMode = GongNavigationModeNavigating
	}

	commitToApply := sc.forwardCommits[len(sc.forwardCommits)-1-sc.commitsBehind+1]

	sc.commitsBehind--
	sc.isApplyingForwardCommit = true
	err := parseAstString(commitToApply, true)
	sc.isApplyingForwardCommit = false
	if err != nil {
		log.Println("error during ApplyForwardCommit: ", err)
		return err
	}

	if computeRefAndOrders != nil {
		computeRefAndOrders()
	}

	return nil
}

// ResetHardCore resets the commits arrays to the current commitsBehind position
func (sc *StageCore) ResetHardCore() {
	newCommitsLen := len(sc.forwardCommits) - sc.commitsBehind
	if newCommitsLen < 0 {
		newCommitsLen = 0
	}

	sc.forwardCommits = sc.forwardCommits[:newCommitsLen]
	sc.backwardCommits = sc.backwardCommits[:newCommitsLen]
	sc.commitsBehind = 0
	sc.navigationMode = GongNavigationModeNormal
}

// SquashCore clears all history commits and marks stage as squashing
func (sc *StageCore) SquashCore() {
	sc.forwardCommits = sc.forwardCommits[:0]
	sc.backwardCommits = sc.backwardCommits[:0]
	sc.commitsBehind = 0
	sc.navigationMode = GongNavigationModeNormal

	sc.modified = true
	sc.isSquashing = true
}

func (sc *StageCore) EndSquash() {
	sc.isSquashing = false
}
