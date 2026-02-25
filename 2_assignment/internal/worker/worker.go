package worker

import (
	"sync"
	"time"

	"github.com/FarizaIsmagambetova/Assignment2/internal/model"
	"github.com/FarizaIsmagambetova/Assignment2/internal/queue"
	"github.com/FarizaIsmagambetova/Assignment2/internal/store"
)

// worker==goroutine
type WorkerPool struct {
	repo  *store.Repository[string, *model.Task]
	queue *queue.TaskQueue
	wg    sync.WaitGroup
}

func NewWorkerPool(repo *store.Repository[string, *model.Task], workers int, queueSize int) *WorkerPool {
	wp := &WorkerPool{
		repo:  repo,
		queue: queue.NewTaskQueue(queueSize), //buffered channel

	}
	for i := 0; i < workers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
	return wp
}
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for {
		task := wp.queue.GetTask()
		if task == nil {
			return
		}
		wp.repo.UpdateStatus(task.ID, "PENDING")
		time.Sleep(3 * time.Second) //simulate some work
		wp.repo.UpdateStatus(task.ID, "IN_PROGRESS")
		time.Sleep(3 * time.Second)
		wp.repo.UpdateStatus(task.ID, "COMPLETED")
	}

}

func (wp *WorkerPool) AddTaskToQueue(task *model.Task) {
	wp.queue.AddTask(task)
}

func (wp *WorkerPool) Stop() {
	wp.queue.Close()
	wp.wg.Wait()
}

//we have queue of tasks to be done,
//we have fixed number of workers(goroutines) that process the tasks from the queue
