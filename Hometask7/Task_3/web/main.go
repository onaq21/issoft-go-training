package main

import (
	"log"
	"text/template"
	"net/http"
	"Task_3/internal/file"
	"Task_3/internal/task"
)

type Server struct {
	tasks []task.Task
	t *template.Template
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := s.t.ExecuteTemplate(w, "base.html", s.tasks); err != nil {
			http.Error(w, "Template execution error: " + err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	tasks, err := file.LoadData()
	if err != nil {
		log.Fatalf("Load data error: %s", err)
	}
	log.Println("Load data: success")

	t, err := template.ParseFiles("web/templates/base.html", "web/templates/list.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %s", err)
	}
	log.Println("Parse templates: success")

	if err := http.ListenAndServe(":5001", &Server{tasks, t}); err != nil {
		log.Fatalf("Listen and serve server error: %s", err)
	}
}