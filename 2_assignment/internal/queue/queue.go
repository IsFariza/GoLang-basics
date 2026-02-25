package queue

import "github.com/FarizaIsmagambetova/Assignment2/internal/model"

type TaskQueue struct {
	queue chan *model.Task
}

func NewTaskQueue(size int) *TaskQueue {
	return &TaskQueue{
		queue: make(chan *model.Task, size),
	}
}
func (q *TaskQueue) AddTask(task *model.Task) {
	q.queue <- task
}
func (q *TaskQueue) GetTask() *model.Task {
	task, ok := <-q.queue
	if !ok {
		return nil
	}
	return task
}
func (q *TaskQueue) Close() {
	close(q.queue)
}
