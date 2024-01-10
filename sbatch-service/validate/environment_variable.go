// Copyright (C) 2024 DeepSquare Association
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

package validate

import "regexp"

var (
	// regexEnvVarName is a matcher from the POSIX standard
	regexEnvVarName = regexp.MustCompilePOSIX(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
)

func EnvVarNameValidator(env string) bool {
	return regexEnvVarName.MatchString(env)
}
