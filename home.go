// +build !windows

/*
 * Logcat - a small utility useful for exposing logs to http requests.
 * Copyright (C) 2020  Nicolò Santamaria
 *
 * Logcat is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Logcat is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with logcat.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import "os"

var HOME = os.Getenv("HOME")
