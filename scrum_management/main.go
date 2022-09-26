package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Name_task        string    `json:"name"`
	Description_task string    `json:"description"`
	Sprint_task      string    `json:"sprint"`
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

// func getTasks(w http.ResponseWriter, r *http.Request) {
// 	db := setupDB()

// 	printMessage("Getting all tasks...")

// 	rows, err := db.Query("SELECT * FROM tasks")

// 	checkErr(err)

// 	var tasks []Task

// 	for rows.Next() {
// 		var Name string
// 		var Description string
// 		var Sprint int
// 		var Assignee string
// 		var Story_point int
// 		var Status string
// 		var Created time.Time
// 		var Last time.Time

// 		err = rows.Scan(&Name, &Description, &Sprint, &Assignee, &Story_point, &Status, &Created, &Last)

// 		checkErr(err)
// 		tasks = append(tasks, Task{Name_task: Name, Description_task: Description, Sprint_task: Sprint, Assignee_task: Assignee,
// 			Status: Status, Created_at: Created, Last_update_at: Last})

// 	}
// 	var response = JsonResponse{Type: "Success", Data: tasks}
// 	json.NewEncoder(w).Encode(response)
// }

// GET ALL THE TASKS = GET /tasks
func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// ADD TASK = POST /tasks
func addTask(c *gin.Context) {
	var task Task
	err := c.BindJSON(&task)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.Created_at = time.Now()
	task.Last_update_at = time.Now()
	c.JSON(http.StatusOK, tasks)

}

// PUT /tasks/{id}
func updateTask(c *gin.Context) {
	Sprint_task := c.Param("sprint")
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	index := -1
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Sprint_task == Sprint_task {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Task not found"})
		return
	}

	tasks[index] = task
	c.JSON(http.StatusOK, tasks)
}

// DELETE /tasks/{sprint}
func deleteTask(c *gin.Context) {
	Sprint_task := c.Param("sprint")
	index := -1
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Sprint_task == Sprint_task {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Task not found"})
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully"})
}

// GET /tasks/search
func getTask(c *gin.Context) {
	status := c.Query("status")
	ListofTasks := make([]Task, 0)

	for i := 0; i < len(tasks); i++ {
		found := false
		for _, s := range tasks[i].Status {
			if string(s) == status {
				found = true
			}
		}
		if found {
			ListofTasks = append(ListofTasks, tasks[i])
		}
	}
	c.JSON(http.StatusOK, ListofTasks)

}
func main() {
	router := gin.Default()
	router.POST("/tasks", addTask)
	router.PUT("/tasks/:sprint", updateTask)
	router.GET("/tasks", getTasks)
	router.DELETE("/tasks/:sprint", deleteTask)
	router.GET("/tasks/search", getTask)
	router.Run()

}
