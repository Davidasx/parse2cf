package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TestCase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type Task struct {
	Name    string     `json:"name"`
	Contest string     `json:"group"`
	Url     string     `json:"url"`
	Memory  int        `json:"memoryLimit"`
	Time    int        `json:"timeLimit"`
	Tests   []TestCase `json:"tests"`
}

const PORT = 10043

func Handle(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	decoder := json.NewDecoder(req.Body)
	task := new(Task)
	if err := decoder.Decode(task); err != nil {
		panic(err)
	}
	for i, x := range task.Tests {
		os.WriteFile(fmt.Sprintf("in%d.txt", i+1), []byte(x.Input), 0600)
		os.WriteFile(fmt.Sprintf("ans%d.txt", i+1), []byte(x.Output), 0600)
	}
	println(fmt.Sprintf("parsed %d testcases form %d\n", len(task.Tests), PORT))
	println("parsed from " + task.Contest)
}

func main() {
	println(fmt.Sprintf("listening form port %d\n", PORT))
	http.HandleFunc("/", Handle)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", PORT), nil)
}
