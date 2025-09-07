package ui

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) updateDatabaseSelector(msg tea.Msg) (*Model, tea.Cmd) {
	var cmd tea.Cmd
	m.dbSelector, cmd = m.dbSelector.Update(msg)
	return m, cmd
}

func (m *Model) updateConfigForm(msg tea.Msg) (*Model, tea.Cmd) {
	var cmd tea.Cmd
	m.configForm, cmd = m.configForm.Update(msg)
	return m, cmd
}
