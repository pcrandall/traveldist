package workbook

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct{}
type errMsg error

type model struct {
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.NewModel()
	ti.Placeholder = writeFileName
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 80

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
	// return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			fallthrough
		case tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			writeFileName = m.textInput.Value()
			return m, tea.Quit
		case tea.KeyTab:
			m.textInput.SetValue(writeFileName)
			// put cursor to date position in filename
			m.textInput.SetCursor(len(writeFileName) - 10)
			return m, cmd
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"\nWhat would you like to save workbook as?\n\n%s\n\n%s",
		m.textInput.View(),
		"[Press tab to autocomplete, Enter when finished]",
	) + "\n"
}
