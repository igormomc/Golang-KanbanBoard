package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type status int

const divisor = 4

const (
	todo status = iota
	inProgress
	done
)

var models []tea.Model

const (
	model status = iota
	form
)
