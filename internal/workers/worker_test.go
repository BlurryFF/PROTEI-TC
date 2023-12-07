package workers

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWorkerPool_NewPool(t *testing.T) {
	if _, err := NewWorkerPool(0, 0); err != incorrectWorkersSize {
		t.Fatalf("expected error when creating pool with 0 workers, got: %v", err)
	}
	if _, err := NewWorkerPool(-1, 0); err != incorrectWorkersSize {
		t.Fatalf("expected error when creating pool with -1 workers, got: %v", err)
	}
	if _, err := NewWorkerPool(1, -1); err != incorrectChannelSize {
		t.Fatalf("expected error when creating pool with -1 channel size, got: %v", err)
	}

	p, err := NewWorkerPool(5, 0)
	if err != nil {
		t.Fatalf("expected no error creating pool, got: %v", err)
	}
	if p == nil {
		t.Fatal("NewWorkerPool returned nil pool for valid input")
	}
}

func TestWorkerPool_MultipleStartStopDontPanic(t *testing.T) {
	p, err := NewWorkerPool(5, 0)
	if err != nil {
		t.Fatal("error creating pool:", err)
	}

	p.Start()
	p.Start()

	p.Stop()
	p.Stop()
}

type testTask struct {
	executeFunc func() error

	shouldErr bool
	wg        *sync.WaitGroup

	mFailure       *sync.Mutex
	failureHandled bool
}

func newTestTask(executeFunc func() error, shouldErr bool, wg *sync.WaitGroup) *testTask {
	return &testTask{
		executeFunc: executeFunc,
		shouldErr:   shouldErr,
		wg:          wg,
		mFailure:    &sync.Mutex{},
	}
}

func (t *testTask) Execute() error {
	if t.wg != nil {
		defer t.wg.Done()
	}

	if t.executeFunc != nil {
		return t.executeFunc()
	}

	time.Sleep(50 * time.Millisecond)
	if t.shouldErr {
		return fmt.Errorf("planned Execute() error")
	}
	return nil
}

func (t *testTask) OnFailure(e error) {
	t.mFailure.Lock()
	defer t.mFailure.Unlock()

	t.failureHandled = true
}

func (t *testTask) hitFailureCase() bool {
	t.mFailure.Lock()
	defer t.mFailure.Unlock()

	return t.failureHandled
}

func TestWorkerPool_Work(t *testing.T) {
	var tasks []*testTask
	wg := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		tasks = append(tasks, newTestTask(nil, false, wg))
	}

	p, err := NewWorkerPool(5, len(tasks))
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}
	p.Start()

	for _, j := range tasks {
		p.AddWork(j)
	}

	wg.Wait()

	for taskNum, task := range tasks {
		if task.hitFailureCase() {
			t.Fatalf("error function called on task %d when it shouldn't be", taskNum)
		}
	}
}

func TestWorkerPool_BlockedAddWorkReleaseAfterStop(t *testing.T) {
	p, err := NewWorkerPool(1, 0)
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}

	p.Start()

	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			p.AddWork(newTestTask(func() error {
				time.Sleep(20 * time.Second)
				return nil
			}, false, nil))
			wg.Done()
		}()
	}

	done := make(chan struct{})
	p.Stop()
	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	select {
	case <-time.After(1 * time.Second):
		t.Fatal("failed because still hanging on AddWork")
	case <-done:
	}
}
