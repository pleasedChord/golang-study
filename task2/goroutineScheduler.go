package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID      string
	Handler func()
	Time    time.Duration
	error   error
}

type TaskScheduler struct {
	MaxWorkers int
	Tasks      []Task
	Results    []Task
	Wg         sync.WaitGroup
	Semophore  chan struct{}
	mu         sync.Mutex
}

func NewTaskScheduler(maxWorkers int) *TaskScheduler {
	return &TaskScheduler{
		MaxWorkers: maxWorkers,
		Semophore:  make(chan struct{}, maxWorkers),
	}
}

func (s *TaskScheduler) AddTask(id string, handler func()) {
	s.Tasks = append(s.Tasks, Task{
		ID:      id,
		Handler: handler,
	})
}

func (s *TaskScheduler) Run() {
	for _, task := range s.Tasks {
		s.Wg.Add(1)
		go s.ExecuteTask(task)
	}
	s.Wg.Wait()
}

func (s *TaskScheduler) ExecuteTask(task Task) {
	defer s.Wg.Done()

	s.Semophore <- struct{}{}
	defer func() {
		<-s.Semophore
	}()

	start := time.Now()
	defer func() {
		task.Time = time.Since(start)

		s.mu.Lock()
		s.Results = append(s.Results, task)
		s.mu.Unlock()
	}()

	defer func() {
		if err := recover(); err != nil {
			task.error = fmt.Errorf("panic:%v", err)
		}
	}()

	task.Handler()
}

func (s *TaskScheduler) GetResult() []Task {
	return s.Results
}

func main() {
	ts := NewTaskScheduler(3)

	ts.AddTask("task1", func() {
		time.Sleep(2 * time.Second)
		fmt.Println("task1 completed")
	})

	ts.AddTask("task2", func() {
		time.Sleep(1 * time.Second)
		fmt.Println("task2 completed")
	})

	ts.AddTask("task3", func() {
		time.Sleep(3 * time.Second)
		fmt.Println("task3 completed")
	})

	ts.Run()

	for _, result := range ts.Results {
		if result.error != nil {
			fmt.Printf("任务 %s 失败: %v，耗时: %v\n", result.ID, result.error, result.Time)
		} else {
			fmt.Printf("任务 %s 成功，耗时: %v\n", result.ID, result.Time)
		}
	}
}
