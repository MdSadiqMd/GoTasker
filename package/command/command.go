package command

import (
	"flag"

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
