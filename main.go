package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	todos := Todos{}

	storage := NewStorage[Todos]()
	err := storage.Load(&todos)

	if err != nil {

		fmt.Println("No existing todos, starting fresh.")
	}

	if cmdFlags := NewCmdFlags(); cmdFlags != nil {
		cmdFlags.Execute(&todos)
		storage.Save(todos)
	} else {


			todoTable := NewTodoTable(&todos, storage)
			todoTable.updateRows()
			if err := tea.NewProgram(todoTable).Start(); err != nil {
				fmt.Println("Error running todo table:", err)
			}
	}

}

type TodoModel struct {
	todos  Todos
	cursor int
}

func NewTodoModel(todos Todos) TodoModel {
	return TodoModel{todos: todos, cursor: 0}
}

func (m TodoModel) Init() tea.Cmd {
	return nil
}

func (m TodoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m TodoModel) View() string {
	s := "Interactive Todo List:\n\n"
	for i, t := range m.todos {
		cursor := "  "
		if m.cursor == i {
			cursor = ">>"
		}
		status := "❌"
		if t.Completed {
			status = "✅"
		}
		s += fmt.Sprintf("%s [%d] %s - %s (Created: %s)\n", cursor, i, t.Title, status, t.CreatedAt.Format(time.RFC1123))
	}
	s += "\nPress q to quit the interactive view.\n"
	return s
}
