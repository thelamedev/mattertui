package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thelamedev/mattertui/internal/tui/api"
	"github.com/thelamedev/mattertui/internal/tui/msg"
	"github.com/thelamedev/mattertui/internal/tui/styles"
)

type LoginModel struct {
	inputs  []textinput.Model
	focused int
	err     error
	client  *api.Client
}

func NewLoginModel() LoginModel {
	m := LoginModel{
		inputs: make([]textinput.Model, 2),
		client: api.NewClient(),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = styles.SubTitle
		t.CharLimit = 64

		switch i {
		case 0:
			t.Placeholder = "Username"
			t.Focus()
			t.PromptStyle = styles.SubTitle
			t.TextStyle = styles.SubTitle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.PromptStyle = styles.SubTitle
			t.TextStyle = styles.SubTitle
		}

		m.inputs[i] = t
	}

	return m
}

func (m LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m LoginModel) Update(message tea.Msg) (LoginModel, tea.Cmd) {
	switch message := message.(type) {
	case tea.KeyMsg:
		switch message.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				username := m.inputs[0].Value()
				password := m.inputs[1].Value()
				return m, func() tea.Msg {
					token, err := m.client.Login(username, password)
					if err != nil {
						return msg.LoginFailMsg(err)
					}
					return msg.LoginSuccessMsg{Token: token, UserName: username}
				}
			}
			m.NextInput()
		case tea.KeyTab, tea.KeyDown:
			m.NextInput()
		case tea.KeyShiftTab, tea.KeyUp:
			m.PrevInput()
		}
	case msg.LoginFailMsg:
		m.err = message
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(message)
	}

	return m, tea.Batch(cmds...)
}

func (m *LoginModel) NextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
	m.updateFocus()
}

func (m *LoginModel) PrevInput() {
	m.focused--
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
	m.updateFocus()
}

func (m *LoginModel) updateFocus() {
	for i := 0; i < len(m.inputs); i++ {
		if i == m.focused {
			m.inputs[i].Focus()
			m.inputs[i].PromptStyle = styles.SubTitle
			m.inputs[i].TextStyle = styles.SubTitle
			continue
		}
		m.inputs[i].Blur()
		m.inputs[i].PromptStyle = lipgloss.NewStyle()
		m.inputs[i].TextStyle = lipgloss.NewStyle()
	}
}

func (m LoginModel) View() string {
	var b strings.Builder

	b.WriteString(styles.Title.Render("Login") + "\n\n")

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := styles.SubTitle.Render("[ Submit ]")
	if m.focused == len(m.inputs)-1 {
		button = styles.Title.Render("[ Submit ]")
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", button)

	if m.err != nil {
		b.WriteString(styles.SubTitle.Render(fmt.Sprintf("Error: %v", m.err)))
	}

	b.WriteString("\n\nPress 'ctrl+r' to Register")

	return b.String()
}
