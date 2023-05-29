package handler

import (
	client "client/internal/app"
	"client/internal/pkg/models"
	"context"
	"fmt"
	"log"
)

func printListOfCommands() {
	fmt.Println("Available commands:")
	fmt.Println(`	'help' to print list of commands`)
	fmt.Println(`	'create-user' to create new user`)
	fmt.Println(`	'create-task' to create new task`)
	fmt.Println(`	'get-user-by-id' to get user by id`)
	fmt.Println(`	'get-task-by-id' to get task by id`)
	fmt.Println(`	'get-all-users' to get all users`)
	fmt.Println(`	'get-all-tasks' to get all tasks for specific user`)
	fmt.Println(`	'update-user' to update user`)
	fmt.Println(`	'update-task' to update task for specific user`)
	fmt.Println(`	'delete-user' to delete user by id`)
	fmt.Println(`	'delete-task' to delete task by id`)
	fmt.Println(`	'exit'`)
}

func scanUser() models.User {
	fmt.Println("Enter user name: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Enter user email: ")
	var email string
	fmt.Scan(&email)
	fmt.Println("Enter user password: ")
	var password string
	fmt.Scan(&password)
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return user
}

func scanTask() models.Task {
	fmt.Println("Enter task title: ")
	var title string
	fmt.Scan(&title)
	fmt.Println("Enter task description: ")
	var description string
	fmt.Scan(&description)
	fmt.Println("Enter user id of user which should do this task: ")
	var userId int64
	fmt.Scan(&userId)
	task := models.Task{
		UserID:      userId,
		Title:       title,
		Description: description,
	}
	return task
}

func scanId() int64 {
	fmt.Println("Enter id: ")
	var id int64
	fmt.Scan(&id)
	return id
}

func updateUser() models.User {
	fmt.Println("Enter user id which you want to update: ")
	var id int64
	fmt.Scan(&id)
	user := scanUser()
	user.ID = id
	return user
}

func updateTask() models.Task {
	fmt.Println("Enter task id which you want to update: ")
	var id int64
	fmt.Scan(&id)
	task := scanTask()
	task.ID = id
	return task
}

func printUser(user *models.User) {
	fmt.Println("User id: ", user.ID)
	fmt.Println("User name: ", user.Name)
	fmt.Println("User email: ", user.Email)
	fmt.Println("User password: ", user.Password)
}

func printTask(task *models.Task) {
	fmt.Println("Task id: ", task.ID)
	fmt.Println("Task user id: ", task.UserID)
	fmt.Println("Task title: ", task.Title)
	fmt.Println("Task description: ", task.Description)
}

func printUsers(users []*models.User) {
	for _, user := range users {
		printUser(user)
		fmt.Println("--------------------------------------------------")
	}
}

func printTasks(tasks []*models.Task) {
	for _, task := range tasks {
		printTask(task)
		fmt.Println("--------------------------------------------------")
	}
}

func CommandHandler(ctx context.Context, client *client.Client) {
	printListOfCommands()
	var command string
	for {
		fmt.Println("Enter command: ")
		fmt.Scan(&command)
		switch command {
		case "help":
			printListOfCommands()
		case "create-user":
			user := scanUser()
			id, err := client.CreateUser(ctx, user)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("User created with id: ", id)
			}
		case "create-task":
			task := scanTask()
			id, err := client.CreateTask(ctx, task)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Task created with id: ", id)
			}
		case "get-user-by-id":
			id := scanId()
			user, err := client.GetUser(ctx, id)
			if err != nil {
				log.Fatal(err)
			} else {
				printUser(user)
			}
		case "get-task-by-id":
			id := scanId()
			task, err := client.GetTask(ctx, id)
			if err != nil {
				log.Fatal(err)
			} else {
				printTask(task)
			}
		case "get-all-users":
			users, err := client.ListUsers(ctx)
			if err != nil {
				log.Fatal(err)
			} else {
				printUsers(users)
			}
		case "get-all-tasks":
			tasks, err := client.ListTasks(ctx)
			if err != nil {
				log.Fatal(err)
			} else {
				printTasks(tasks)
			}
		case "update-user":
			user := updateUser()
			status, err := client.UpdateUser(ctx, user)
			if err != nil {
				log.Fatal(err)
			} else {
				if status {
					fmt.Println("User updated")
				} else {
					fmt.Println("User not found")
				}
			}
		case "update-task":
			task := updateTask()
			status, err := client.UpdateTask(ctx, task)
			if err != nil {
				log.Fatal(err)
			} else {
				if status {
					fmt.Println("Task updated")
				} else {
					fmt.Println("Task not found")
				}
			}
		case "delete-user":
			id := scanId()
			status, err := client.DeleteUser(ctx, id)
			if err != nil {
				log.Fatal(err)
			} else {
				if status {
					fmt.Println("User deleted")
				} else {
					fmt.Println("User not found")
				}
			}
		case "delete-task":
			id := scanId()
			status, err := client.DeleteTask(ctx, id)
			if err != nil {
				log.Fatal(err)
			} else {
				if status {
					fmt.Println("Task deleted")
				} else {
					fmt.Println("Task not found")
				}
			}
		case "exit":
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Unknown command, type 'help' to print list of commands")
		}
	}

}
