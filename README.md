# 🎮 Tic Tac Toe CLI (Go)

A clean, modular Tic Tac Toe implementation in Go with CLI interface, featuring proper separation of concerns and input validation.

## ✨ Features

- � Human vs. Human gameplay in terminal
- ✅ Robust input validation with retry logic
- 🖥️ Persistent board state between moves
- 🏆 Win/draw detection
- 🧩 Modular architecture (game, ui packages)
- 🚫 No external dependencies

## 📦 Installation

```bash
git clone https://github.com/lacriman/tic-tac-toe.git
cd tic-tac-toe
go run main.go



🏗️ Project Structure

tic-tac-toe/
├── main.go            # Entry point & game loop
├── go.mod             # Go module definition
├── game/              # Core game logic
│   ├── board.go       # Board state management
│   ├── rules.go       # Win/draw detection
│   └── validate.go    # Move validation
└── ui/                # User interface
    └── input.go       # Input handling & prompts



🎮 How to Play

Players alternate turns entering coordinates (1-3 for row and column)

Valid moves update the board automatically

Game continues until:

A player gets 3 in a row (win)

All cells are filled (draw)

Example move:

text
Player X's turn
Enter row (1-3): 2
Enter column (1-3): 2



🧠 Design Philosophy

Modularity: Clean separation between game logic and UI

Testability: Packages structured for easy unit testing

Error Handling: Graceful recovery from invalid inputs

Simplicity: Minimal code for maximum clarity



Built with ❤️ by Yaroslav