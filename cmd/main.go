package main

import (
	"fmt"

	"github.com/MdSadiqMd/GoTasker/package/database"
	"github.com/MdSadiqMd/GoTasker/package/handlers"
)

func main() {
	todos := handlers.Todos{}
	storage :=database.NewStorage[handlers.Todos]("todos.json")
	storage.Load(&todos)
	todos.Add("Buy Milk")
	todos.Add("Buy Bread")
	fmt.Printf("%+v\n\n", todos)
	todos.Toggle(0)
	fmt.Printf("%+v", todos)
	todos.Print()
	storage.Save(todos)
}
