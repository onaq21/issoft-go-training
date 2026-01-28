package main 

import (
	"Task_3/internal/file"
	"Task_3/internal/task"
	"flag"
	"fmt"
	"os"
)

func main() {
	showTasks := flag.Bool("list", false, "show all the tasks")
	completeTask := flag.Int("complete", -1, "this task will be completed")
	newTask := flag.String("task", "", "add this task")

	flag.Parse()

	tasks, err := file.LoadData()
	if err != nil {
		fmt.Printf("Load data error: %s", err)
		os.Exit(1)
	}
	fmt.Println("Load data: success")

	if *newTask != "" {
		task.AddTask(&tasks, *newTask)
	}

	if *completeTask > 0{
		if *completeTask <= len(tasks) {
			task.CompleteTask(tasks, *completeTask)
		}	else {
			fmt.Println("Value of -complete is out of range")
			os.Exit(1)
		}
	}

	if *showTasks {
		task.ListTasks(tasks)
	}

	if *newTask != "" || *completeTask > 0 {
		if err = file.SaveData(tasks); err != nil {
			fmt.Printf("Save data error: %s", err)
			os.Exit(1)
		}
		fmt.Println("Save data: success")
	}
}