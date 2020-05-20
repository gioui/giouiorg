// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// START EXAMPLE OMIT
var isChecked widget.Bool

func themedApplication(gtx *layout.Context, th *material.Theme) {
	var checkboxLabel string
	if isChecked.Value {
		checkboxLabel = "checked"
	} else {
		checkboxLabel = "not-checked"
	}

	layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func() { material.H3(th, "Hello, World!").Layout(gtx) }),
		layout.Rigid(func() { material.CheckBox(th, checkboxLabel).Layout(gtx, &isChecked) }),
	)
}

// END EXAMPLE OMIT
