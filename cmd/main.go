package main

import (
	"fmt"
  "net/http"
	"log"
	"database/sql"
	"time"

	"github.com/Yuji5117/todo-app-go/domain/entity"
	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)

func getTasks(d *sql.DB, c echo.Context) error {
  tasks, err := d.Query("SELECT * FROM tasks")
	if err != nil {
			fmt.Println("Err2")
			panic(err.Error())
	}

	var tasksResponse []entity.Task

	for tasks.Next() {
		var task entity.Task

		var createdAt, updatedAt string
		var deletedAt sql.NullString

		err := tasks.Scan(&task.ID, &task.Title, &task.Done, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
				fmt.Println("スキャンが失敗しました。")
				panic(err.Error())
		}

		if createdAt != "" {
				task.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
				if err != nil {
						fmt.Println("パースに失敗")
				}
		}
		if updatedAt != "" {
				task.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
				if err != nil {
						fmt.Println("パースに失敗")
				}
		}
		if deletedAt.Valid {
				task.DeletedAt, err = time.Parse("2006-01-02 15:04:05", deletedAt.String)
				if err != nil {
						fmt.Println("deleted_at パースに失敗")
				}
		} else {
				task.DeletedAt, err = time.Parse("2006-01-02 15:04:05", "0000-00-00 00:00:00")
				if err != nil {
					fmt.Println("deleted_at パースに失敗")
				}
		}

		tasksResponse = append(tasksResponse, task)
	}

	return c.JSON(http.StatusOK, tasksResponse)
}

func saveTask(d *sql.DB, c echo.Context) error {
	title := c.FormValue("title")
	task, err := d.Exec(
		"INSERT INTO tasks (title) VALUES (?)",
		title,
	)
	if err != nil {
		log.Fatalf("insertUser db.Exec error err:%v", err)
	}

	id, err := task.LastInsertId()
	if err != nil {
		log.Fatalf("insertUser db.Exec error err:%v", err)
	}

	return c.JSON(http.StatusCreated, id)
}

func updateTask(d *sql.DB, c echo.Context) error {
	id := c.Param("id")
	title := c.FormValue("title")
	done := c.FormValue("done")

	doneValue := 0
	if done == "true" {
			doneValue = 1
  }

	_, err := d.Exec(
		"UPDATE tasks SET title = ?, done = ? WHERE id = ?",
		title,
		doneValue,
		id,
	)
	if err != nil {
		log.Fatalf("insertUser db.Exec error err:%v", err)
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

func main()  {
	db, err := sql.Open("mysql", "user:12345678@tcp(db:3306)/todo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()


	e := echo.New()
	e.GET("/tasks", func(c echo.Context) error { return getTasks(db, c) })
	e.POST("/tasks", func(c echo.Context) error { return saveTask(db, c) })
	e.POST("/tasks/:id", func(c echo.Context) error { return updateTask(db, c) })
	e.DELETE("/tasks/:id", func(c echo.Context) error { return deleteTask(db, c) })


	e.Logger.Fatal(e.Start(":8080"))

}