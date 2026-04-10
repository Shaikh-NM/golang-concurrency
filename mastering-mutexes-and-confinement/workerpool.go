package main

import (
	"fmt"
	"sync"
	"time"
)

// task definition
type Task struct {
	ID int
}

// way to process the task
func (t *Task) Process() {
	fmt.Println("Processing task %d\n", t.ID)
	time.Sleep(2 * time.Second)
	fmt.Println("Processed task %d\n", t.ID)
}

// worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// functions to execute the worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)
	wp.wg.Wait()
}

func main() {
	tasks := make([]Task, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = Task{ID: i + 1}
	}

	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5,
	}

	wp.Run()
	fmt.Println("All tasks have been processed")
}
