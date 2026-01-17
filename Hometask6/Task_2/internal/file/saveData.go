package file

import (
	"Task_2/internal/task"
	"encoding/json"
	"os"
	"fmt"
)

func SaveData (tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return fmt.Errorf("Marshal data error: %s", err)
	}

	if err := os.WriteFile("data/tasks.json", data, 0644); err != nil {
		return fmt.Errorf("Write file error: %s", err)
	}
	return nil
}