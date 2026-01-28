package task

import "fmt"

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("There aren't any tasks")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s:\n", i + 1, task.Name)
		fmt.Printf("\tCompleted: %t\n", task.Completed)
		fmt.Printf("\tCreated at: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
		if task.Completed {
			fmt.Printf("\tCompleted at: %s\n", task.CompletedAt.Format("2006-01-02 15:04:05"))
		}
	}
}