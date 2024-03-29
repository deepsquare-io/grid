// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package log initialize the logger for the TUI.
package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// I is the instance of the default logger.
var I *zap.Logger

var atom = zap.NewAtomicLevel()

func init() {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	I = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.Lock(os.Stdout),
		atom,
	), zap.AddCaller())
	atom.SetLevel(zap.InfoLevel)
}

// EnableDebug changes the log level to Debug.
func EnableDebug() {
	atom.SetLevel(zap.DebugLevel)
	I.Debug("Enabled debug logging.")
}
