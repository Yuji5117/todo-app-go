package main

import (
  "net/http"
	"github.com/Yuji5117/todo-app-go/domain/entity"
	"github.com/labstack/echo/v4"
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
	e := echo.New()
	e.GET("/tasks", getTasks)
	e.POST("/tasks", saveTask)


	e.Logger.Fatal(e.Start(":8080"))

}