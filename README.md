# Task_Tracker


A simple command-line task tracker written in Go. Stores tasks in `tasks.json` and provides a minimal interface to add, list, and manage tasks.

**Features**
- Add tasks with a title and optional details.
- Persist tasks to `tasks.json` in the project directory.
- List current tasks.

**Prerequisites**
- Go 1.18+ installed (see `go version`).

**Quick Start**

1. Clone the repo (if you haven't already):

```
git clone https://github.com/naveen-93/Task_Tracker.git
cd Task_Tracker
```

2. Run directly with `go run`:

```
go run main.go
```

3. Or build a binary and run it:

```
go build -o task_tracker
./task_tracker
```

Note: The project stores tasks in `tasks.json` located in the project folder. If the file doesn't exist, the program will create it when you add your first task.

**Usage (examples)**

The project is a small CLI. Typical flows include:

- Add a task (example):

```
go run main.go add "Buy groceries"
```

- List tasks (example):

```
go run main.go list
```

Replace `go run main.go` with the built binary (`./task_tracker`) if you built it.

