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

// Package style provides object for styling the TUI.
package style

import "github.com/charmbracelet/lipgloss"

var (
	// StandardHeight is the height of the main window.
	StandardHeight       = 13
	primaryDarkestColor  = lipgloss.Color("#9202de")
	primaryColor         = lipgloss.Color("#BD43FD")
	primaryLightestColor = lipgloss.Color("#dea2fe")
	errorColor           = lipgloss.Color("#ff3333")
	green                = lipgloss.Color("#04B575")
	white                = lipgloss.Color("#FAFAFA")
	// Base is the box of the main window.
	Base = Box.Copy()
	// Error is the style for errors.
	Error = lipgloss.NewStyle().Foreground(errorColor)
	// Foreground is the main style.
	Foreground = lipgloss.NewStyle().Foreground(primaryDarkestColor)
	// NoError is the style for positive results.
	NoError = lipgloss.NewStyle().Foreground(green)
	// AccentForeground is the secondary style.
	AccentForeground = lipgloss.NewStyle().Foreground(primaryColor)

	// Title1 is the most accentuated style for titles.
	Title1 = lipgloss.NewStyle().
		Bold(true).
		Foreground(white).
		Background(primaryDarkestColor).
		Align(lipgloss.Center).
		Width(22)

	// Title2 is the second accent style for titles.
	Title2 = lipgloss.NewStyle().
		Bold(true).
		Foreground(primaryDarkestColor).
		Width(22)

	// LogTitle is the style for the log page title.
	LogTitle = Box.Copy().
			Padding(0, 1)

	// LogInfo is the style for the log page info.
	LogInfo = LogTitle.Copy()
	// LeftVerticalSeparator is the style to make a vertical separator on the left.
	LeftVerticalSeparator = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderLeft(true).
				BorderForeground(primaryDarkestColor)
	// Box is the style for surrounding an area.
	Box = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(primaryDarkestColor)
	// FocusBox is the style for surrounding an area with accent.
	FocusBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryLightestColor)
)

// BoolToYN convert a boolean to a stylized yes/no. "yes" being positive.
func BoolToYN(b bool) string {
	if b {
		return NoError.Render("yes")
	}
	return Error.Render("no")
}

// BoolToYNColorReverted convert a boolean to a stylized yes/no. "no" being positive.
func BoolToYNColorReverted(b bool) string {
	if b {
		return Error.Render("yes")
	}
	return NoError.Render("no")
}

// OnError prints with the error style if the error is not nil.
func OnError(errorStyle lipgloss.Style, v string, err error) string {
	if err != nil {
		return errorStyle.Render(v)
	}
	return v
}
