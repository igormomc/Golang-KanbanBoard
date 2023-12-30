# GoLang Kanban Project

![GoLang Logo](https://github.com/rfyiamcool/golang_logo/blob/master/png/golang_20.png)


## Introduction

This small project is a Kanban board implementation in Go, utilizing the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework for building modern command-line applications and the [Lip Gloss](https://github.com/charmbracelet/lipgloss) styling library for layout and styling. It's designed to provide a simple and intuitive interface for task management directly from your command line.

## Features

- Simple and intuitive CLI interface.
- Manage tasks across three columns: To Do, In Progress, and Done.
- Navigate and modify tasks using keyboard shortcuts.

## Installation

To install this project, you need to have Go installed on your system. Follow these steps:

1. Clone the repository:

2. Navigate to the project directory and run build
    ```bash
    go build .
3. Run it
    ```bash
    go run .
    

## Looks
<img width="1168" alt="Screenshot 2023-12-29 at 22 25 20" src="https://github.com/igormomc/Golang-KanbanBoard/assets/60653284/532a2ba8-c5e6-4ab6-b836-93fdcb752a01">



## Usage

Run the application using the following command:

### Keyboard Shortcuts

- `Left` or `h`: Move focus to the left column.
- `Right` or `l`: Move focus to the right column.
- `Enter`: Move the selected task to the next column.
- `n`: Create a new task.
- `Ctrl+C` or `q`: Quit the application.

---

For more information, please refer to the [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss) documentation.
