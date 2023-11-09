package main

import (
  "net/http"
	"log"
	"database/sql"

	"github.com/Yuji5117/todo-app-go/domain/entity"
	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)

func getTasks(c echo.Context) error {
	var tasks [3]entity.Task = [3]entity.Task{entity.NewTask("散歩"), entity.NewTask("掃除"), entity.NewTask("宿題")}
	return c.JSON(http.StatusOK, tasks)
}

func saveTask(c echo.Context) error {
	title := c.FormValue("title")
	task := entity.NewTask(title)
	return c.JSON(http.StatusCreated, task)
}


func main()  {
	db, err := sql.Open("mysql", "user:12345678@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

	e := echo.New()
	e.GET("/tasks", getTasks)
	e.POST("/tasks", saveTask)


	e.Logger.Fatal(e.Start(":8080"))

}