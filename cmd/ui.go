package cmd

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

func setupTable(tasks []Task) *list.List {
	l := list.New()
	itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("212")).MarginRight(1)
	enumaratorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)

	l.ItemStyle(itemStyle)
	// l.Enumerator(list.)
	l.EnumeratorStyle(enumaratorStyle)

	for _, task := range tasks {
		tempList := l.Item(task.Name)

		subList := list.New(task.Project, task.Status, task.CreatedAt)
		subList.Enumerator(list.Bullet)
		subList.ItemStyle(lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#0000ff")))

		tempList.Item(subList)
	}

	return l
}
