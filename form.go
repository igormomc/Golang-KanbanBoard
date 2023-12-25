package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Form struct {
	focused     status
	title       textinput.Model
	description textarea.Model
}

func NewForm(focused status) *Form {
	form := &Form{focused: focused}
	form.title = textinput.New()
	form.title.Focus()
	form.description = textarea.New()
	return form
}

func (f *Form) Init() tea.Cmd {
	return nil
}

func (f *Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return f, tea.Quit
		case "enter":
			if f.title.Focused() {
				f.title.Blur()
				f.description.Focus()
				return f, textarea.Blink
			} else {
				// Make sure we're passing a pointer to Form here
				models[form] = f
				return models[model], func() tea.Msg {
					return f.CreateTask()
				}
			}
		}
	}

	if f.title.Focused() {
		f.title, cmd = f.title.Update(msg)
	} else {
		f.description, cmd = f.description.Update(msg)
	}
	return f, cmd
}
func (f *Form) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		f.title.View(),
		f.description.View(),
	)
}

func (f *Form) CreateTask() tea.Msg {
	task := NewTask(f.focused, f.title.Value(), f.description.Value())
	return task
}
