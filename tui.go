package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type TodoTable struct {
	table      table.Model
	todos      *Todos
	storage    *Storage[Todos]
	inputMode  inputMode
	textInput  textinput.Model
	inputIndex int
}

type inputMode int

const (
	normalMode inputMode = iota
	addMode
	editMode
)

func NewTodoTable(todos *Todos, storage *Storage[Todos]) TodoTable {
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Title", Width: 40},
		{Title: "Status", Width: 8},
		{Title: "Created", Width: 30},
		{Title: "Completed", Width: 30},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(true)
	t.SetStyles(s)

	ti := textinput.New()

	ti.Placeholder = "Enter todo title..."
	ti.Focus()

	return TodoTable{
		table:     t,
		todos:     todos,
		storage:   storage,
		textInput: ti,
	}
}

func (m TodoTable) Init() tea.Cmd {
	return nil
}

func (m TodoTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case m.inputMode == normalMode:
			switch msg.String() {
			case "q", "ctrl+c":

				// os.Exit(0)

				return m, tea.Quit
			case "t":
				if len(*m.todos) > 0 {
					index := m.table.Cursor()
					m.todos.toggle(index)
					m.storage.Save(*m.todos)
					m.updateRows()
				}
			case "d":
				if len(*m.todos) > 0 {
					index := m.table.Cursor()
					m.todos.delete(index)
					m.storage.Save(*m.todos)
					m.updateRows()
				}
			case "e":
				if len(*m.todos) > 0 {
					m.inputMode = editMode
					m.inputIndex = m.table.Cursor()
					m.textInput.Reset()
					m.textInput.SetValue((*m.todos)[m.inputIndex].Title)
					m.textInput.Focus()
					return m, nil
				}
			case "a":
				m.inputMode = addMode
				m.textInput.Reset()
				m.textInput.Focus()
				return m, nil
			}

		case m.inputMode == addMode || m.inputMode == editMode:
			switch msg.String() {
			case "enter":
				if m.inputMode == addMode {
					m.todos.add(m.textInput.Value())
				} else {
					m.todos.edit(m.inputIndex, m.textInput.Value())
				}
				m.storage.Save(*m.todos)
				m.updateRows()
				m.inputMode = normalMode
				m.textInput.Blur()
			case "esc":
				m.inputMode = normalMode
				m.textInput.Blur()
			}
		}
	}

	if m.inputMode == addMode || m.inputMode == editMode {
		var tiCmd tea.Cmd
		m.textInput, tiCmd = m.textInput.Update(msg)
		return m, tiCmd
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *TodoTable) updateRows() {
	var rows []table.Row
	for i, todo := range *m.todos {
		status := "❌"
		completedAt := ""
		if todo.Completed {
			status = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		rows = append(rows, table.Row{
			fmt.Sprint(i),
			todo.Title,
			status,
			todo.CreatedAt.Format(time.RFC1123),
			completedAt,
		})
	}
	m.table.SetRows(rows)
}

func (m TodoTable) View() string {
	var view string
	if m.inputMode == addMode {
		view = "ADD MODE - Enter new todo title:\n" + m.textInput.View()
	} else if m.inputMode == editMode {
		view = "EDIT MODE - Edit todo title:\n" + m.textInput.View()
	} else {
		view = m.table.View()
	}

	helpText := "\n(↑/↓) navigate • (a) add • (e) edit • (t) toggle • (d) delete • (q) quit"
	if m.inputMode != normalMode {
		helpText = "\n(enter) confirm • (esc) cancel"
	}

	return baseStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			view,
			helpText,
		),
	)
}
