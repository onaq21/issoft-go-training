package task

import "time"

func CompleteTask(tasks []Task, id int) {
	tasks[id - 1].Completed = true
	tasks[id - 1].CompletedAt = time.Now()
}