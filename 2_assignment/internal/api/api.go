package api

import (
	"encoding/json"

	"github.com/FarizaIsmagambetova/Assignment2/internal/model"
	"github.com/FarizaIsmagambetova/Assignment2/internal/store"
	"github.com/FarizaIsmagambetova/Assignment2/internal/worker"

	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// TaskHandler handles HTTP requests related to tasks(creates tasks, lists them etc)
type TaskHandler struct {
	repo    *store.Repository[string, *model.Task]
	wp      *worker.WorkerPool
	mu      sync.Mutex
	counter uint64
}

func NewTaskHandler(repo *store.Repository[string, *model.Task], wp *worker.WorkerPool) *TaskHandler {
	return &TaskHandler{
		repo: repo,
		wp:   wp,
	}
}

// StatsHandler handles /api/stats endpoint (task statistics, how many submitted/progress/completed)
type StatsHandler struct {
	repo *store.Repository[string, *model.Task]
}

func NewStatsHandler(repo *store.Repository[string, *model.Task]) *StatsHandler {
	return &StatsHandler{
		repo: repo,
	}
}

// POST /tasks
func (th *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Payload string `json:"payload"`
	}
	//Decoder.Decode(&req) reads HTTP request body json
	//and automatically fills the req struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	th.mu.Lock()
	th.counter++
	id := strconv.FormatUint(th.counter, 10)
	th.mu.Unlock()

	task := &model.Task{
		ID:      id,
		Payload: req.Payload,
		Status:  "PENDING",
	}

	th.repo.Add(id, task)
	th.wp.AddTaskToQueue(task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	errEncode := json.NewEncoder(w).Encode(map[string]string{"id": id})
	if errEncode != nil {
		return
	}
}

// GET /tasks
func (th *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	type TaskInfo struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}

	tasks := make([]TaskInfo, 0, len(th.repo.GetAll()))

	for _, task := range th.repo.GetAll() {
		tasks = append(tasks, TaskInfo{
			ID:     task.ID,
			Status: task.Status,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return
	}
}

// GET /tasks/{id}
func (th *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	//mux.Vars takes path variables from URL that matches {id} }pattern
	//for /tasks/1, mux will take "42" as "id"
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	task, ok := th.repo.Get(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		return
	}
}

// GET /stats
func (sh *StatsHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	sc := sh.repo.StatusCount
	stats := map[string]int{
		"submitted":   sc["PENDING"] + sc["IN_PROGRESS"] + sc["COMPLETED"],
		"completed":   sc["COMPLETED"],
		"in_progress": sc["IN_PROGRESS"],
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(stats)
	if err != nil {
		return
	}

}
