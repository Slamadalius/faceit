package listener

import (
	"fmt"
	"sync"
)

var (
	que = make(chan UserAction, 100)
	wg  = &sync.WaitGroup{}
)

type UserAction struct {
	UserID string
	Action string
}

type Listener struct {
	workersCount int
}

// this can be used with rabbitMQ or other messaging platform to notify
// other services about changes to user

func NewListener(workersCount int) *Listener {
	return &Listener{
		workersCount: workersCount,
	}
}

func (l *Listener) Start() {
	for i := 0; i < l.workersCount; i++ {
		wg.Add(1)
		go work(i + 1)
	}
}

func (l *Listener) Stop() {
	close(que)
	fmt.Println("channel [que] closed")
}

func (l *Listener) Wait() {
	wg.Wait()
	fmt.Println("waiting for listener to shutdown")
}

func AddRequest(request UserAction) {
	que <- request
}
