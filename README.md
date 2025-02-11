# tinyTodo 📝

A minimalist CLI todo application with both command-line flags and an interactive Terminal User Interface (TUI) along with persistent storage in a json file. Built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Features ✨

- Interactive TUI mode with keyboard navigation
- Command-line flag operations for quick tasks
- Persistent storage of todos in a json file
- Cross-platform support

## Demo

https://github.com/user-attachments/assets/83ffbf32-4b80-44c7-b72d-1cae6684bf8f

## Installation 🚀

```bash
# Clone the repository
git clone https://github.com/yourusername/tinyTodo.git

# Navigate to the project directory
cd tinyTodo

# Build the application
go build -o tinytodo
```

## Usage 💻

### TUI Mode

Simply run the application without any flags to enter the interactive TUI mode:

```bash
./tinytodo
```

TUI Navigation:
- `↑/k`: Move cursor up
- `↓/j`: Move cursor down
- `a`: Add new todo
- `e`: Edit selected todo
- `t`: Toggle completion status
- `d`: Delete selected todo
- `q`: Quit application

### Command-line Mode

```bash
# Add a new todo
./tinytodo -add "Buy groceries"

# Edit a todo (format: index:new_title)
./tinytodo -edit "0:Buy milk"

# Toggle todo completion status
./tinytodo -toggle 1

# Delete a todo
./tinytodo -del 2

# List all todos
./tinytodo -list

# Show help
./tinytodo -h
```

## JSON file storage locations 

- Windows → Stores in %APPDATA%\todo\todos.json
- macOS → Stores in ~/Library/Application Support/todo/todos.json
- Linux → Stores in ~/.config/todo/todos.json

## Dependencies 📦

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions

## Building from Source 🔨

```bash
# Ensure you have Go 1.23+ installed
go version

# Clone the repository
git clone https://github.com/yourusername/tinyTodo.git

# Navigate to project directory
cd tinyTodo

# Install dependencies
go mod download

# Build the project
go build -o tinytodo
```

## Contributing 🤝

Contributions are welcome! Feel free to:
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Acknowledgments 🙏

- [Charm](https://charm.sh) for their amazing TUI libraries
- The Go community for inspiration and support
