package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Task struct {
	Name_task        string    `json:"name"`
	Description_task string    `json:"description"`
	Sprint_task      int       `json:"sprint"`
	Assignee_task    string    `json:"assignee"`
	Story_point      int       `json:"story_point"`
	Status           string    `json:"status"`
	Created_at       time.Time `json:"created_at"`
	Last_update_at   time.Time `json:"last_update_at"`
}

var tasks []Task

func init() {
	tasks = make([]Task, 0)
	file, _ := ioutil.ReadFile("tasks.json")
	_ = json.Unmarshal([]byte(file), &tasks)
}

func main() {

}
