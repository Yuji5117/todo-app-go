package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Yuji5117/todo-app-go/domain/entity"

	_ "github.com/go-sql-driver/mysql"
)

type TaskMySQL struct {
	db *sql.DB
}

func NewTaskMySQL(db *sql.DB) *TaskMySQL {
	return &TaskMySQL{
		db: db,
	}
}

func (r *TaskMySQL) List() ([]entity.Task, error) {
	stmt, err := r.db.Query("SELECT * FROM tasks ORDER BY id DESC")
	if err != nil {
		fmt.Println("Err2")
		panic(err.Error())
	}

	var tasks []entity.Task

	for stmt.Next() {
		var task entity.Task

		var createdAt, updatedAt string
		var deletedAt sql.NullString

		err := stmt.Scan(&task.ID, &task.Title, &task.Done, &createdAt, &updatedAt, &deletedAt)
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

		tasks = append(tasks, entity.Task{
			ID:        task.ID,
			Title:     task.Title,
			Done:      task.Done,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
			DeletedAt: task.DeletedAt,
		})
	}

	return tasks, nil
}
