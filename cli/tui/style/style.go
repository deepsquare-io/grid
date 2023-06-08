package style

import "github.com/charmbracelet/lipgloss"

var (
	StandardHeight  = 13
	foregroundColor = lipgloss.Color("#9202de")
	Base            = Box.Copy()
	Foreground      = lipgloss.NewStyle().Foreground(foregroundColor)

	LogTitle = Box.Copy().
			Padding(0, 1)

	LogInfo               = LogTitle.Copy()
	LeftVerticalSeparator = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderLeft(true).
				BorderForeground(foregroundColor)
	Box = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(foregroundColor)
)
