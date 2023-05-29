package database_command_handler

import (
	"context"
	"fmt"
	"homework/internal/pkg/repository"
	database "homework/internal/pkg/repository/postgres"
	"log"
)

type Handler interface {
	GetListOfCommands() []string
	IsMyCommand(command string) bool
	HandleCommand(ctx context.Context, command string)
	GetGoalOfHandler() string
}

type DatabaseCommandHandler struct {
	users *database.UsersRepo
	tasks *database.TasksRepo
	// key - command, value - description
	commands map[string]string
}

func CreateDataBaseCommandHandler(users *database.UsersRepo, tasks *database.TasksRepo) Handler {
	databaseHandler := DatabaseCommandHandler{}
	databaseHandler.users = users
	databaseHandler.tasks = tasks
	databaseHandler.commands = map[string]string{
		"create-user":    "to create new user",
		"create-task":    "to create new task",
		"get-user-by-id": "to get user by id",
		"get-task-by-id": "to get task by id",
		"get-all-users":  "to get all users",
		"get-all-tasks":  "to get all tasks for specific user",
		"update-user":    "to update user",
		"update-task":    "to update task for specific user",
		"delete-user":    "to delete user by id",
		"delete-task":    "to delete task by id",
	}
	return databaseHandler
}

func (handler DatabaseCommandHandler) GetListOfCommands() []string {
	var commands []string
	for key, value := range handler.commands {
		commands = append(commands, fmt.Sprintf("'%s' %s", key, value))
	}
	return commands
}

func (handler DatabaseCommandHandler) IsMyCommand(command string) bool {
	_, ok := handler.commands[command]
	return ok
}

func (handler DatabaseCommandHandler) HandleCommand(ctx context.Context, command string) {
	switch command {
	case "create-user":
		user := scanUser()
		id, err := handler.users.Add(ctx, &user)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("User created with id: ", id)
		}
	case "create-task":
		task := scanTask()
		id, err := handler.tasks.Add(ctx, &task)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Task created with id: ", id)
		}
	case "get-user-by-id":
		id := scanId()
		user, err := handler.users.GetById(ctx, id)
		if err != nil {
			log.Fatal(err)
		} else {
			printUser(user)
		}
	case "get-task-by-id":
		id := scanId()
		task, err := handler.tasks.GetById(ctx, id)
		if err != nil {
			log.Fatal(err)
		} else {
			printTask(task)
		}
	case "get-all-users":
		users, err := handler.users.List(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			printUsers(users)
		}
	case "get-all-tasks":
		tasks, err := handler.tasks.List(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			printTasks(tasks)
		}
	case "update-user":
		user := updateUser()
		status, err := handler.users.Update(ctx, &user)
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
		status, err := handler.tasks.Update(ctx, &task)
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
		status, err := handler.users.Delete(ctx, id)
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
		status, err := handler.tasks.Delete(ctx, id)
		if err != nil {
			log.Fatal(err)
		} else {
			if status {
				fmt.Println("Task deleted")
			} else {
				fmt.Println("Task not found")
			}
		}
	default:
		fmt.Println("Unknown command, type 'help' to print list of commands")
	}
}

func (handler DatabaseCommandHandler) GetGoalOfHandler() string {
	return "to work with database"
}

func scanUser() repository.User {
	fmt.Println("Enter user name: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Enter user email: ")
	var email string
	fmt.Scan(&email)
	fmt.Println("Enter user password: ")
	var password string
	fmt.Scan(&password)
	user := repository.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return user
}

func scanTask() repository.Task {
	fmt.Println("Enter task title: ")
	var title string
	fmt.Scan(&title)
	fmt.Println("Enter task description: ")
	var description string
	fmt.Scan(&description)
	fmt.Println("Enter user id of user which should do this task: ")
	var userId int64
	fmt.Scan(&userId)
	task := repository.Task{
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

func updateUser() repository.User {
	fmt.Println("Enter user id which you want to update: ")
	var id int64
	fmt.Scan(&id)
	user := scanUser()
	user.ID = id
	return user
}

func updateTask() repository.Task {
	fmt.Println("Enter task id which you want to update: ")
	var id int64
	fmt.Scan(&id)
	task := scanTask()
	task.ID = id
	return task
}

func printUser(user *repository.User) {
	fmt.Println("User id: ", user.ID)
	fmt.Println("User name: ", user.Name)
	fmt.Println("User email: ", user.Email)
	fmt.Println("User password: ", user.Password)
	fmt.Println("User created at: ", user.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("User updated at: ", user.UpdatedAt.Time.Format("2006-01-02 15:04:05"))
}

func printTask(task *repository.Task) {
	fmt.Println("Task id: ", task.ID)
	fmt.Println("Task user id: ", task.UserID)
	fmt.Println("Task title: ", task.Title)
	fmt.Println("Task description: ", task.Description)
	fmt.Println("Task created at: ", task.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("Task updated at: ", task.UpdatedAt.Time.Format("2006-01-02 15:04:05"))
}

func printUsers(users []*repository.User) {
	for _, user := range users {
		printUser(user)
		fmt.Println("--------------------------------------------------")
	}
}

func printTasks(tasks []*repository.Task) {
	for _, task := range tasks {
		printTask(task)
		fmt.Println("--------------------------------------------------")
	}
}
