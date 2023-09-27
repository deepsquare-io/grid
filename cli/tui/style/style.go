// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package styles provides object for styling the TUI.
package style

import "github.com/charmbracelet/lipgloss"

var (
	StandardHeight       = 13
	primaryDarkestColor  = lipgloss.Color("#9202de")
	primaryColor         = lipgloss.Color("#BD43FD")
	primaryLightestColor = lipgloss.Color("#dea2fe")
	errorColor           = lipgloss.Color("#ff3333")
	green                = lipgloss.Color("#04B575")
	white                = lipgloss.Color("#FAFAFA")
	Base                 = Box.Copy()
	Error                = lipgloss.NewStyle().Foreground(errorColor)
	Foreground           = lipgloss.NewStyle().Foreground(primaryDarkestColor)
	NoError              = lipgloss.NewStyle().Foreground(green)
	AccentForeground     = lipgloss.NewStyle().Foreground(primaryColor)

	Title1 = lipgloss.NewStyle().
		Bold(true).
		Foreground(white).
		Background(primaryDarkestColor).
		Align(lipgloss.Center).
		Width(22)

	Title2 = lipgloss.NewStyle().
		Bold(true).
		Foreground(primaryDarkestColor).
		Width(22)

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

func BoolToYN(b bool) string {
	if b {
		return NoError.Render("yes")
	}
	return Error.Render("no")
}

func BoolToYNColorReverted(b bool) string {
	if b {
		return Error.Render("yes")
	}
	return NoError.Render("no")
}

func StyleOnError(errorStyle lipgloss.Style, v string, err error) string {
	if err != nil {
		return errorStyle.Render(v)
	}
	return v
}
