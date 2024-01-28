package models

func (stage *StageStruct) subscribe() <-chan int {
	stage.rwMutex.Lock()
	defer stage.rwMutex.Unlock()

	ch := make(chan int)
	stage.subscribers = append(stage.subscribers, ch)
	return ch
}

func (stage *StageStruct) broadcastNbCommitToBack(nbCommitToBack int) {
	stage.rwMutex.RLock()
	defer stage.rwMutex.RUnlock()

	for _, ch := range stage.subscribers {
		ch <- nbCommitToBack
	}
}
