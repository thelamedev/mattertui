package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/thelamedev/mattertui/internal/tui/components"
	"github.com/thelamedev/mattertui/internal/tui/msg"
	"github.com/thelamedev/mattertui/internal/tui/styles"
)

type Session struct {
	Token    string
	UserName string
}

type State int

const (
	StateLogin State = iota
	StateRegister
	StateMain
)

type Model struct {
	Session  Session
	State    State
	Login    components.LoginModel
	Register components.RegisterModel
}

func NewModel() Model {
	return Model{
		Session:  Session{},
		State:    StateLogin,
		Login:    components.NewLoginModel(),
		Register: components.NewRegisterModel(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.Login.Init(), m.Register.Init())
}

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch message := message.(type) {
	case tea.KeyMsg:
		switch message.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "ctrl+r":
			m.State = StateRegister
			return m, nil
		case "ctrl+l":
			m.State = StateLogin
			return m, nil
		}
	case msg.LoginSuccessMsg:
		m.Session.Token = message.Token
		m.Session.UserName = message.UserName
		m.State = StateMain
		return m, nil
	case msg.SwitchToRegisterMsg:
		m.State = StateRegister
		return m, nil
	case msg.SwitchToLoginMsg:
		m.State = StateLogin
		return m, nil
	}

	switch m.State {
	case StateLogin:
		m.Login, cmd = m.Login.Update(message)
		cmds = append(cmds, cmd)
	case StateRegister:
		m.Register, cmd = m.Register.Update(message)
		cmds = append(cmds, cmd)
	case StateMain:
		// Handle main app update
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	switch m.State {
	case StateLogin:
		return styles.AppContainer.Render(m.Login.View())
	case StateRegister:
		return styles.AppContainer.Render(m.Register.View())
	case StateMain:
		return styles.AppContainer.Render(
			styles.Title.Render("MatterTUI Dashboard") + "\n\n" +
				"You are logged in!\nUser Name: " + m.Session.UserName + "\n\nPress 'q' to quit.",
		)
	}
	return ""
}
