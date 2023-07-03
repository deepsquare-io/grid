// Copyright (C) 2023 DeepSquare
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
