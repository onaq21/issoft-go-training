package file

import (
	"text/template"
	"os"
	"fmt"
)

func SaveReport(t *template.Template, data any) error {
	file, err := os.Create("report.md")
	if err != nil {
		return fmt.Errorf("Create file error: %s", err)
	}
	defer file.Close()

	err = t.ExecuteTemplate(file, "report.md", data)
	if err != nil {
		return fmt.Errorf("Execute template error: %s", err)
	}

	return nil
}