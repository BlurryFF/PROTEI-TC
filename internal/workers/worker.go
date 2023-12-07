package workers

import (
	"errors"
	"log"
	"sync"
)

var (
	incorrectWorkersSize = errors.New("cant create worker pool with less than 1 worker")
	incorrectChannelSize = errors.New("channel size less than 1")
)

type Pool interface {
	Start()
	Stop()
	AddWork(Task)
}

type Task interface {
	Execute() error
	OnFailure(error)
}

type WorkerPool struct {
	numWorkers int

	tasks chan Task
	quit  chan struct{}

	start sync.Once
	stop  sync.Once
}

func NewWorkerPool(numWorkers int, channelSize int) (Pool, error) {
	if numWorkers <= 0 {
		return nil, incorrectWorkersSize
	}
	if channelSize < 0 {
		return nil, incorrectChannelSize
	}

	tasks := make(chan Task, channelSize)

	return &WorkerPool{
		numWorkers: numWorkers,
		tasks:      tasks,

		start: sync.Once{},
		stop:  sync.Once{},

		quit: make(chan struct{}),
	}, nil
}

func (p *WorkerPool) Start() {
	p.start.Do(func() {
		log.Print("starting worker pool")
		p.startWorkers()
	})
}

func (p *WorkerPool) Stop() {
	p.stop.Do(func() {
		log.Print("stopping worker pool")
		close(p.quit)
	})
}

func (p *WorkerPool) AddWork(t Task) {
	select {
	case p.tasks <- t:
	case <-p.quit:
	}
}

func (p *WorkerPool) startWorkers() {
	for i := 0; i < p.numWorkers; i++ {
		go func(workerNum int) {
			log.Printf("starting worker %d", workerNum)

			for {
				select {
				case <-p.quit:
					log.Printf("stopping worker %d with quit channel\n", workerNum)
					return
				case task, ok := <-p.tasks:
					if !ok {
						log.Printf("stopping worker %d with closed tasks channel\n", workerNum)
						return
					}
					if err := task.Execute(); err != nil {
						task.OnFailure(err)
					}
				}
			}
		}(i)
	}
}
