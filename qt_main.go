/*
 * Copyright (C) 2020 Nicolas SCHWARTZ
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301, USA
 */

// +build !GTK

package main

import (
	"os"

	"github.com/therecipe/qt/widgets"

	"github.com/StarAurryon/lpedit/qt/qtctrl"
	"github.com/StarAurryon/lpedit/qt/ui"
)

func main() {
    app := widgets.NewQApplication(len(os.Args), os.Args)
    app.SetWheelScrollLines(1)
    c := qtctrl.NewController(app)
    ui.NewLPEdit(c, nil).Show()
    widgets.QApplication_Exec()
}
