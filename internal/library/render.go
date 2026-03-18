package library

import (
	"dantes/internal/books"
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	choices  []books.Book
	cursor   int
	selected int
}

func (l *Library) Model() model {
	return model{
		choices:  l.books,
		selected: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", "space":
			m.selected = m.cursor
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := "Choose a book to read?\n\n"

	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " "
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if m.selected == i {
			checked = "x" // selected!
		}

		s += fmt.Sprintf("%s [%s] %s, by %s\n", cursor, checked, choice.Title, choice.Author)
	}

	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}
