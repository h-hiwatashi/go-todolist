package main

import (
	"fmt"
	"todo_app/app/models"
)

func main(){
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLdriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	fmt.Println(models.Db)
	fmt.Println("Database setup complete")

	// CREATE USER
	// u:= &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@ecample.com"
	// u.Password = "test"
	// u.CreateUser()

	//GET USER
	user, _ := models.GetUser(3)
	fmt.Println(user)

	//UPDATE USER
	// user.Name = "test2"
	// user.Email = "test2@exmaple.com"
	// user.UpdateUser()
	// fmt.Println(user)

	//DELETE USER
	// user.DeleteUser()

	//CREATE TODO
	// user.CreateTodo("Third todo")

	//GET TODO
	// todo, _ := models.GetTodo(3)
	// fmt.Println(todo)

	// todos, _ :=models.GetTodos()
	// fmt.Println(todos)

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// fmt.Println(todos)

	todo, _ := models.GetTodo(3)
	todo.Content = "Updated content"
	todo.UpdateTodo()

}