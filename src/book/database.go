package main

import (
	//"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

//Task is the struct used to identify tasks
type Task struct {
	Id      int
	Title   string
	Content string
	Created string
}

//Context is the struct passed to templates
type Context struct {
	Tasks      []Task
	Navigation string
	Search     string
	Message    string
}

func init() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func main() {
	http.HandleFunc("/", ShowAllTasksFunc)
	http.HandleFunc("/add/", AddTaskFunc)
	fmt.Println("running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//ShowAllTasksFunc is used to handle the "/" URL which is the default one
func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		context := GetTasks() //true when you want non deleted notes
		w.Write([]byte(context.Tasks[0].Title))
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func GetTasks() Context {
	var task []Task
	var context Context
	var TaskID int
	var TaskTitle string
	var TaskContent string
	var TaskCreated time.Time
	var getTasksql string

	getTasksql = "select id, title, content, created_date from task;"

	rows, err := db.Query(getTasksql)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&TaskID, &TaskTitle, &TaskContent, &TaskCreated)
		if err != nil {
			fmt.Println(err)
		}
		TaskCreated = TaskCreated.Local()
		a := Task{Id: TaskID, Title: TaskTitle, Content: TaskContent,
			Created: TaskCreated.Format(time.UnixDate)[0:20]}
		task = append(task, a)
	}
	context = Context{Tasks: task}
	return context
}

//AddTaskFunc is used to handle the addition of new task, "/add" URL
func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	title := "random title"
	content := "random content"
	truth := AddTask(title, content)
	if truth != nil {
		log.Fatal("Error adding task")
	}
	w.Write([]byte("Added task"))
}

//AddTask is used to add the task in the database
func AddTask(title, content string) error {
	query := "insert into task(title, content, created_date, last_modified_at) values(?,?,datetime(), datetime())"
	restoreSQL, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	tx, err := db.Begin()
	_, err = tx.Stmt(restoreSQL).Exec(title, content)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	} else {
		log.Print("insert successful")
		tx.Commit()
	}
	return err
}
