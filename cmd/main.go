package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Yuji5117/todo-app-go/infra/repository"

	"github.com/Yuji5117/todo-app-go/adapter/presenter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getTasks(repo *repository.TaskMySQL, c echo.Context) error {

	tasks, err := repo.List()
	if err != nil {
		log.Fatalf("getTask d.List error err:%v", err)
	}

	var tasksResponse []presenter.TaskDTO

	for _, task := range tasks {
		dto := presenter.TaskDTO{ID: task.ID, Title: task.Title, Done: task.Done}
		tasksResponse = append(tasksResponse, dto)
	}

	return c.JSON(http.StatusOK, tasksResponse)
}

func saveTask(d *sql.DB, c echo.Context) error {
	var reqBody struct {
		Title string `json:"title"`
	}

	if err := c.Bind(&reqBody); err != nil {
		log.Fatalf("insertTask db.Exec error err:%v", c)
	}

	title := reqBody.Title

	task, err := d.Exec(
		"INSERT INTO tasks (title) VALUES (?)",
		title,
	)
	if err != nil {
		log.Fatalf("insertTask db.Exec error err:%v", err)
	}

	id, err := task.LastInsertId()
	if err != nil {
		log.Fatalf("insertTask db.Exec error err:%v", err)
	}

	return c.JSON(http.StatusCreated, id)
}

func updateTask(d *sql.DB, c echo.Context) error {
	id := c.Param("id")

	var reqBody struct {
		Title *string `json:"title,omitempty"`
		Done  *bool   `json:"done,omitempty"`
	}

	if err := c.Bind(&reqBody); err != nil {
		log.Fatalf("insertTask c.Bind error err:%v", c)
	}

	title := reqBody.Title
	done := reqBody.Done

	if title != nil {
		_, err := d.Exec(
			"UPDATE tasks SET title = ? WHERE id = ?",
			title,
			id,
		)
		if err != nil {
			log.Fatalf("updateTask for title db.Exec error err:%v", err)
		}

		return c.JSON(http.StatusCreated, id)
	}

	if done != nil {
		doneValue := 0
		if *done == true {
			doneValue = 1
		}

		_, err := d.Exec(
			"UPDATE tasks SET done = ? WHERE id = ?",
			doneValue,
			id,
		)
		if err != nil {
			log.Fatalf("updateTask for done db.Exec error err:%v", err)
		}

		return c.JSON(http.StatusCreated, id)
	}

	_, err := d.Exec(
		"UPDATE tasks SET title = ?, done = ? WHERE id = ?",
		title,
		done,
		id,
	)
	if err != nil {
		log.Fatalf("updateTask for title & done db.Exec error err:%v", err)
	}

	return c.JSON(http.StatusCreated, id)
}

func deleteTask(d *sql.DB, c echo.Context) error {
	id := c.Param("id")

	_, err := d.Exec(
		"DELETE FROM tasks WHERE id = ?",
		id,
	)
	if err != nil {
		log.Fatalf("insertUser db.Exec error err:%v", err)
	}

	return c.JSON(http.StatusCreated, id)
}

func main() {
	db, err := sql.Open("mysql", "user:12345678@tcp(db:3306)/todo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

	repo := repository.NewTaskMySQL(db)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.GET("/tasks", func(c echo.Context) error { return getTasks(repo, c) })
	e.POST("/tasks", func(c echo.Context) error { return saveTask(db, c) })
	e.PATCH("/tasks/:id", func(c echo.Context) error { return updateTask(db, c) })
	e.DELETE("/tasks/:id", func(c echo.Context) error { return deleteTask(db, c) })

	e.Logger.Fatal(e.Start(":8080"))

}
