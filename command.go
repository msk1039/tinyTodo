package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
	Help   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.BoolVar(&cf.Help, "h", false, "Show help information")

	flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Todo CLI Application\n\n")
        fmt.Fprintf(os.Stderr, "Usage:\n")
        fmt.Fprintf(os.Stderr, "  todo [flags]\n\n")
        fmt.Fprintf(os.Stderr, "If NO FLAGS are provided, an interactive TUI will be launched.\n\n")
        fmt.Fprintf(os.Stderr, "Flags:\n")
        fmt.Fprintf(os.Stderr, "  -add string\n")
        fmt.Fprintf(os.Stderr, "        Add a new todo with the specified title\n")
        fmt.Fprintf(os.Stderr, "  -edit string\n")
        fmt.Fprintf(os.Stderr, "        Edit a todo by providing index and new title (Format: id:new_title)\n")
        fmt.Fprintf(os.Stderr, "  -del int\n")
        fmt.Fprintf(os.Stderr, "        Delete a todo by its index\n")
        fmt.Fprintf(os.Stderr, "  -toggle int\n")
        fmt.Fprintf(os.Stderr, "        Toggle completion status of a todo by its index\n")
        fmt.Fprintf(os.Stderr, "  -list\n")
        fmt.Fprintf(os.Stderr, "        List all todos\n")
        fmt.Fprintf(os.Stderr, "  -h\n")
        fmt.Fprintf(os.Stderr, "        Show this help message\n\n")
        fmt.Fprintf(os.Stderr, "Examples:\n")
        fmt.Fprintf(os.Stderr, "  todo                          # Launch interactive TUI\n")
        fmt.Fprintf(os.Stderr, "  todo -add \"Buy groceries\"     # Add new todo\n")
        fmt.Fprintf(os.Stderr, "  todo -edit \"0:Buy milk\"       # Edit todo at index 0\n")
        fmt.Fprintf(os.Stderr, "  todo -toggle 1                # Toggle completion of todo at index 1\n")
        fmt.Fprintf(os.Stderr, "  todo -del 2                   # Delete todo at index 2\n")
        fmt.Fprintf(os.Stderr, "  todo -list                    # Show all todos\n")
    }

	flag.Parse()

	if cf.Help {
        flag.Usage()
        os.Exit(0)
    }

	if cf.Add == "" && cf.Del == -1 && cf.Edit == "" && cf.Toggle == -1 && cf.List == false {
		return nil
	}

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	// fmt.Println("Executing command")
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
