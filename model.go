package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	focused  status
	lists    []list.Model
	err      error
	loaded   bool
	quitting bool
}

func New() *Model {
	return &Model{}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			columnStyle.Width(msg.Width / divisor)
			focusedStyle.Width(msg.Width / divisor)
			columnStyle.Height(msg.Height - divisor)
			focusedStyle.Height(msg.Height - divisor)
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			m.Prev()
		case "right", "l":
			m.Next()
		case "enter":
			// Wrap the tea.Msg in a function to return it as tea.Cmd
			return m, func() tea.Msg {
				return m.MoveToNext()
			}
		case "n":
			models[model] = m
			models[form] = NewForm(m.focused)
			return models[form], nil
		}
	case Task:
		task := msg
		m.lists[task.status].InsertItem(len(m.lists[task.status].Items()), task)
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	if m.quitting {
		return ""
	}
	if m.loaded {
		todoView := m.lists[todo].View()
		inProgView := m.lists[inProgress].View()
		doneView := m.lists[done].View()

		switch m.focused {
		case inProgress:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(todoView),
				focusedStyle.Render(inProgView),
				columnStyle.Render(doneView),
			)
		case done:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(todoView),
				columnStyle.Render(inProgView),
				focusedStyle.Render(doneView),
			)
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				focusedStyle.Render(todoView),
				columnStyle.Render(inProgView),
				columnStyle.Render(doneView),
			)
		}
	} else {
		return "Loading..."
	}
}

func (m *Model) Next() {
	if m.focused == done {
		m.focused = todo
	} else {
		m.focused++
	}
}

func (m *Model) Prev() {
	if m.focused == todo {
		m.focused = done
	} else {
		m.focused--
	}
}

func (m *Model) MoveToNext() tea.Msg {
	selectedItem := m.lists[m.focused].SelectedItem()
	selectedTask, ok := selectedItem.(Task)
	if !ok {
		return nil
	}
	currentIndex := m.lists[m.focused].Index()

	m.lists[selectedTask.status].RemoveItem(currentIndex)
	selectedTask.Next()
	m.lists[selectedTask.status].InsertItem(currentIndex, selectedTask)

	return nil
}

func (m *Model) initLists(width, height int) {
	d := list.NewDefaultDelegate()

	// Change colors
	blue := lipgloss.Color("#2396a6")       // Normal color
	darkerBlue := lipgloss.Color("#00e3ff") // Selected color

	d.Styles.NormalTitle = d.Styles.NormalTitle.Foreground(blue)
	d.Styles.NormalDesc = d.Styles.NormalTitle.Copy() // reuse the title style here

	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(darkerBlue)
	d.Styles.SelectedDesc = d.Styles.SelectedTitle.Copy() // reuse the title style here

	// Initialize the list model with our delegate
	defaultList := list.New([]list.Item{}, d, width/divisor, height/2)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}
	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "Play tennis", description: "Its fun"},
		Task{status: todo, title: "Buy Gifts", description: "Christmas is coming soon"},
		Task{status: todo, title: "Go Sleep", description: "Its important"},
	})

	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{status: todo, title: "Stay cool", description: "Cooool"},
	})
	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{status: todo, title: "Idk man", description: "Just bec"},
	})
}
