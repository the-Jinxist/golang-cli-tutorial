package cmd

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/charmbracelet/lipgloss/table"
)

func setupList(tasks []Task) *list.List {
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

func setupTable(tasks []Task) *table.Table {
	HeaderStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("99")).Width(15).Height(20).AlignHorizontal(lipgloss.Center)
	EvenRowStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Padding(2).AlignHorizontal(lipgloss.Center)
	OddRowStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#A9A9A9")).Padding(2).AlignHorizontal(lipgloss.Center)

	table := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99")).BorderBottom(true)).
		BorderBottom(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				if col == 1 {
					//Can change color depending on status
					return EvenRowStyle.Italic(true)
				}
				return EvenRowStyle
			default:
				if col == 1 {
					//Can change color depending on status
					return OddRowStyle.Italic(true)
				}
				return OddRowStyle
			}
		}).
		Headers("ID", "NAME", "STATUS", "PROJECT", "TIME OF CREATION")

	for _, task := range tasks {
		table.Row(strconv.Itoa(task.ID), strings.ToUpper(task.Name), task.Status, task.Project, task.CreatedAt.Local().String())
	}

	return table
}
