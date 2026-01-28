package task

import "time"

func AddTask(tasks *[]Task, task string) {
	var id int
	if len(*tasks) == 0 { 
		id = 1
	}	else { id = (*tasks)[len(*tasks) - 1].ID + 1}
	*tasks = append(*tasks, Task{id, task, false, time.Now(), time.Time{}})
}