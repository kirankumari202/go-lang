package main

import (
	"fmt"
	"go-lang/models"
)

func main() {
	var TaskModel models.TaskModel
	Tasks, err := TaskModel.FindAll()
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("Task Management List")
		for _, Task := range Tasks {
			fmt.Println("ID: ", Task.ID)
			fmt.Println("TITLE: ", Task.Title)
			fmt.Println("DESCRIPTION: ", Task.Description)
			fmt.Println("STATUS: ", Task.Status)
			fmt.Println("------------------")

		}
	}

}
