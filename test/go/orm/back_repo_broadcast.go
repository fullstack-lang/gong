package orm

func (backRepoStruct *BackRepoStruct) Subscribe() <-chan int {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()

	ch := make(chan int)
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	return ch
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack(nbCommitToBack int) {
	backRepoStruct.rwMutex.RLock()
	defer backRepoStruct.rwMutex.RUnlock()

	for _, ch := range backRepoStruct.subscribers {
		ch <- nbCommitToBack
	}
}
