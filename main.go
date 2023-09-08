package main

import "fmt"

type Task struct {
	Description string
	Completed   bool
}

var tasks []Task

func addTask(description string) {
	newTask := Task{
		Description: description,
		Completed:   false,
	}

	tasks = append(tasks, newTask)
}

func markTaskCompleted(index int) {
	if index >= 0 && index < len(tasks) {
		tasks[index].Completed = true
	}
}

func listTasks() {
	fmt.Println("Tasks: ")
	for i, task := range tasks {
		fmt.Printf("%d: %s (Completed: %v)\n", i+1, task.Description, task.Completed)
	}
}

func main() {

	tasks = []Task{}
	addTask("Go for a run")
	addTask("Buy groceries")
	for {
		fmt.Println("\nTodo App menu")
		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as Completed")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task description: ")
			var description string
			fmt.Scanln(&description)
			addTask(description)
			fmt.Println("Task added succesfully")

		case 2:
			listTasks()
			fmt.Print("Enter the task number to mark as completed: ")
			var taskNumber int
			fmt.Scanln(&taskNumber)
			index := taskNumber - 1
			markTaskCompleted(index)
			fmt.Println("Task marked as completed!")

		case 3:
			listTasks()

		case 4:
			fmt.Println("Exiting todo app.")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
