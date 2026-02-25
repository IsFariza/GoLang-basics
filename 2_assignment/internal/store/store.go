package store

import (
	"sync"

	"github.com/FarizaIsmagambetova/Assignment2/internal/model"
)

type Repository[K comparable, V any] struct {
	mu          sync.RWMutex
	data        map[K]V
	StatusCount map[string]int
}

func NewRepository[K comparable, V any]() *Repository[K, V] {
	return &Repository[K, V]{
		data: make(map[K]V),
		StatusCount: map[string]int{
			"PENDING":     0,
			"IN_PROGRESS": 0,
			"COMPLETED":   0,
		},
	} //returns pointer to new Repository (TaskStore previously) (in-memory storage)
}
func (r *Repository[K, V]) Add(key K, value V) {
	r.mu.Lock()         //to avoid multiple goroutines accessing one shared map at teh same time
	defer r.mu.Unlock() //when method returns, unlock the mutex

	r.data[key] = value
	task := any(value).(*model.Task)
	r.StatusCount[task.Status]++
}
func (r *Repository[K, V]) Get(key K) (V, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	value, ok := r.data[key]
	return value, ok
}
func (r *Repository[K, V]) GetAll() []V {
	r.mu.RLock()
	defer r.mu.RUnlock()

	allTasks := make([]V, 0, len(r.data))
	for _, v := range r.data {
		allTasks = append(allTasks, v)
	}
	return allTasks
}

func (r *Repository[K, V]) UpdateStatus(k K, newStatus string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	v, ok := r.data[k]
	if !ok {
		return false
	}

	task, ok := any(v).(*model.Task)
	if !ok {
		return false
	}
	if task.Status != "PENDING" { //to ensure pending is not decremented
		r.StatusCount[task.Status]--
	}
	task.Status = newStatus
	//when add new task, pending is incremented there, so
	if newStatus != "PENDING" { //to ensure pending is not incremented again
		r.StatusCount[newStatus]++
	}

	return true
}



//type TaskStore struct {
//	repo *Repository[string, *model.Task]
//
//}
//
//func NewTaskStore() *TaskStore {
//	return &TaskStore{
//		repo: NewRepository[string, *model.Task](),
//	}
//}
//func (s *TaskStore) Add(task *model.Task) {
//	s.repo.Add(task.ID, task)
//}
//func (s *TaskStore) GetByID(id string) (*model.Task, bool) {
//	return s.repo.Get(id)
//}
//func (s *TaskStore) GetAll() []*model.Task {
//	return s.repo.GetAll()
//}
//
//func (s *TaskStore) UpdateStoreStatus(id string, status string) {
//	s.repo.Update(id, func(t *model.Task) {
//		t.Status = status
//	})
//}
