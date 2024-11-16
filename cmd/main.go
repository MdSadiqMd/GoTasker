package main

import (
	"fmt"

	"github.com/MdSadiqMd/GoTasker/package/handlers"
)

func main() {
	todos := handlers.Todos{}
	todos.Add("Buy Milk")
	todos.Add("Buy Bread")
	fmt.Printf("%+v\n\n", todos)
	todos.Delete(0)
	fmt.Printf("%+v", todos)
}
