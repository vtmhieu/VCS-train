package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Task `json:"data"`
	Message string `json:"message"`
}

var tasks []Task

func init() {
	tasks = make([]Task, 0)
	file, _ := ioutil.ReadFile("tasks.json")
	_ = json.Unmarshal([]byte(file), &tasks)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "Task"
)

func setupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting all tasks...")

	rows, err := db.Query("SELECT * FROM tasks")

	checkErr(err)

	var tasks []Task

	for rows.Next() {
		var Name string
		var Description string
		var Sprint int
		var Assignee string
		var Story_point int
		var Status string
		var Created time.Time
		var Last time.Time

		err = rows.Scan(&Name, &Description, &Sprint, &Assignee, &Story_point, &Status, &Created, &Last)

		checkErr(err)
		tasks = append(tasks, Task{Name_task: Name, Description_task: Description, Sprint_task: Sprint, Assignee_task: Assignee,
			Status: Status, Created_at: Created, Last_update_at: Last})

	}
	var response = JsonResponse{Type: "Success", Data: tasks}
	json.NewEncoder(w).Encode(response)
}

func addTask(w http.ResponseWriter, r *http.Request) {

}

func updateTask(w http.ResponseWriter, r *http.Request) {

}

func deleteTask(w http.ResponseWriter, r *http.Request) {

}

func getTask(w http.ResponseWriter, r *http.Request) {

}
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", addTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))

}
