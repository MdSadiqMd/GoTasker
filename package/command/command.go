package command

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MdSadiqMd/GoTasker/package/handlers"
	"github.com/MdSadiqMd/GoTasker/package/types"
)

func NewCmdFlags() *types.CmdFlags {
	cf := types.CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo by index to toggle complete true/false")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.Parse()

	return &cf
}

func Execute(cf *types.CmdFlags, todos *handlers.Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}
		todos.Edit(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Del != -1:
		todos.Delete(cf.Del)
	default:
		fmt.Println("Invalid command")
	}
}
