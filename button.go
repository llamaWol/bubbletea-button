
package button

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


// Model is the Bubble Tea model for this text input element.
type Model struct {
	Label			string

	FocusStyle		lipgloss.Style
	BlurStyle		lipgloss.Style

	focus 			bool
}

// New creates a new model with default settings.
func New() Model {
	return Model{
		Label: "Press me!",

		FocusStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffeb")).
			Background(lipgloss.Color("#ff6b97")).
			Padding(0, 2),
		BlurStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffeb")).
			Background(lipgloss.Color("#303030")).
			Padding(0, 2),

		focus: false,
	}
}

// Focused returns the focus state on the model.
func (m Model) Focused() bool {
	return m.focus
}

// Focus sets the focus state on the model. When the model is in focus it can
// receive keyboard input and the cursor will be shown.
func (m *Model) Focus() {
	m.focus = true
}

// Blur removes the focus state on the model.  When the model is blurred it can
// not receive keyboard input and the cursor will be hidden.
func (m *Model) Blur() {
	m.focus = false
}

// Update is the Bubble Tea update loop.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}


func (m Model) View() string {
	if m.focus {
		return m.FocusStyle.Render(m.Label)
	} else {
		return m.BlurStyle.Render(m.Label)
	}
}
