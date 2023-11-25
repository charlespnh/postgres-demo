package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle           = lipgloss.NewStyle().Margin(0, 2)
	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

/**************************************************************************************/

type KeyMap struct {
	// get key.Binding			// get is performed by default to list all entries
	create key.Binding
	update key.Binding
	delete key.Binding

	enter key.Binding
	back  key.Binding
}

var CustomKeyMap = KeyMap{
	create: key.NewBinding(
		key.WithKeys("c"),           // actual keybindings
		key.WithHelp("c", "create"), // corresponding help text
	),
	update: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "update email"),
	),
	delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),

	enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "enter"),
	),
	back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
}

/**************************************************************************************/

type refreshMsg struct{}
type mode int

const (
	nav mode = iota
	create
	update
)

type Model struct {
	mode  mode
	input textinput.Model
	list  list.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case refreshMsg:
		students, err := GetAllStudents()
		if err != nil {
			display(&m, err.Error())
		}
		num := len(students)
		items := make([]list.Item, num)
		for i, s := range students {
			items[i] = list.Item(s)
		}
		m.list.SetItems(items)
		m.mode = nav
	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case tea.KeyMsg:
		if m.input.Focused() {
			switch {
			case key.Matches(msg, CustomKeyMap.enter):
				if m.mode == create {
					raw := m.input.Value()
					parsed := strings.Split(raw, ",")
					student := Student{
						first_name:      parsed[0],
						last_name:       parsed[1],
						email:           parsed[2],
						enrollment_date: parsed[3],
					}
					_, err := AddStudent(student)
					if err != nil {
						display(&m, err.Error())
					} else {
						display(&m, "Query successful")
						cmd = func() tea.Msg {
							return refreshMsg{}
						}
					}
				} else if m.mode == update {
					email := m.input.Value()
					id := m.cursorStudentID()
					_, err := UpdateStudentEmail(id, email)
					if err != nil {
						display(&m, err.Error())
					} else {
						display(&m, "Query successful")
						cmd = func() tea.Msg {
							return refreshMsg{}
						}
					}
				}
				m.input.SetValue("")
				m.mode = nav
				m.input.Blur()
				cmds = append(cmds, cmd)
			case key.Matches(msg, CustomKeyMap.back):
				m.input.SetValue("")
				m.mode = nav
				m.input.Blur()
			}
			m.input, cmd = m.input.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			switch {
			case key.Matches(msg, CustomKeyMap.create):
				display(&m, "Creating a new student...")
				m.mode = create
				m.input.Focus()
				cmd = textinput.Blink
			case key.Matches(msg, CustomKeyMap.update):
				display(&m, "Updating student email...")
				m.mode = update
				m.input.Focus()
				cmd = textinput.Blink
			case key.Matches(msg, CustomKeyMap.delete):
				display(&m, "Deleting student email...")
				if items := m.list.Items(); len(items) > 0 {
					id := m.cursorStudentID()
					_, err := DeleteStudent(id)
					if err != nil {
						display(&m, err.Error())
					} else {
						display(&m, "Query successful")
					}
					cmd = func() tea.Msg {
						return refreshMsg{}
					}
					m.list.ResetSelected()
				} else {
					display(&m, "No entries to delete")
				}
			default:
				m.list, cmd = m.list.Update(msg)
			}
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.input.Focused() {
		return docStyle.Render(m.list.View() + "\n" + m.input.View())
	}
	return docStyle.Render(m.list.View() + "\n")
}

func (m Model) cursorStudentID() int {
	items := m.list.Items()
	item := items[m.list.Index()]
	return item.(Student).ID()
}

func InitList() tea.Model {
	students, err := GetAllStudents()
	if err != nil {
		panic(err)
	}
	num := len(students)
	items := make([]list.Item, num)
	for i, s := range students {
		items[i] = list.Item(s)
	}

	m := Model{
		mode:  nav,
		input: textinput.New(),
		list:  list.New(items, list.NewDefaultDelegate(), 8, 8),
	}
	m.input.Prompt = "$ "
	m.input.Placeholder = "Enter..."
	m.input.CharLimit = 500
	m.input.Width = 100
	m.list.Title = "List of Students"
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			CustomKeyMap.create,
			CustomKeyMap.update,
			CustomKeyMap.delete,
			CustomKeyMap.back,
		}
	}
	return m
}

func display(m *Model, msg string) {
	m.list.NewStatusMessage(statusMessageStyle(msg))
}

/**************************************************************************************/
