package style

import "github.com/charmbracelet/lipgloss"

var (
	StandardHeight   = 13
	primaryDarkest   = lipgloss.Color("#9202de")
	primary          = lipgloss.Color("#BD43FD")
	primaryLightest  = lipgloss.Color("#dea2fe")
	Base             = Box.Copy()
	Foreground       = lipgloss.NewStyle().Foreground(primaryDarkest)
	AccentForeground = lipgloss.NewStyle().Foreground(primary)

	LogTitle = Box.Copy().
			Padding(0, 1)

	LogInfo               = LogTitle.Copy()
	LeftVerticalSeparator = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderLeft(true).
				BorderForeground(primaryDarkest)
	Box = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(primaryDarkest)
	FocusBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryLightest)
)
