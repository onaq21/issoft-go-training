package file

import (
	"fmt"
	"os"
	"encoding/json"
	"Task_3/internal/task"
)

func LoadData() ([]task.Task, error) {
	data, err := os.ReadFile("data/tasks.json")
	if os.IsNotExist(err) {
		_, err = os.Create("data/tasks.json")
			if err != nil {
				return nil, fmt.Errorf("Create file error: %s", err)
			}
		return []task.Task{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("Read file error: %s", err)
	} 

	if len(data) == 0 { return []task.Task{}, nil }

	var tasks []task.Task
	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("Unmarshal error: %s", err)
	}
	return tasks, nil
}