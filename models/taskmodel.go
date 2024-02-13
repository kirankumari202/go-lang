package models

import (
	"go-lang/configure"
	"go-lang/entites"
)

type TaskModel struct {
}

func (*TaskModel) FindAll() ([]entites.Task, error) {
	db, err := configure.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from Task")
		if err2 != nil {
			return nil, err
		} else {
			var Tasks []entites.Task
			for rows.Next() {
				var Task entites.Task
				rows.Scan(&Task.ID, &Task.Title, &Task.Description, &Task.Status)
				Tasks = append(Tasks, Task)
			}

			return Tasks, nil
		}

	}
}
