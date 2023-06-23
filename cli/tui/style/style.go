package style

import "github.com/charmbracelet/lipgloss"

var (
	StandardHeight       = 13
	primaryDarkestColor  = lipgloss.Color("#9202de")
	primaryColor         = lipgloss.Color("#BD43FD")
	primaryLightestColor = lipgloss.Color("#dea2fe")
	errorColor           = lipgloss.Color("#4b1113")
	Base                 = Box.Copy()
	Error                = lipgloss.NewStyle().Foreground(errorColor)
	Foreground           = lipgloss.NewStyle().Foreground(primaryDarkestColor)
	AccentForeground     = lipgloss.NewStyle().Foreground(primaryColor)

	LogTitle = Box.Copy().
			Padding(0, 1)

	LogInfo               = LogTitle.Copy()
	LeftVerticalSeparator = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderLeft(true).
				BorderForeground(primaryDarkestColor)
	Box = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(primaryDarkestColor)
	FocusBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryLightestColor)
)
