// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// START EXAMPLE OMIT
var isChecked widget.Bool

func themedApplication(gtx layout.Context, th *material.Theme) layout.Dimensions {
	var checkboxLabel string
	if isChecked.Value {
		checkboxLabel = "checked"
	} else {
		checkboxLabel = "not-checked"
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(material.H3(th, "Hello, World!").Layout),
		layout.Rigid(material.CheckBox(th, &isChecked, checkboxLabel).Layout),
	)
}

// END EXAMPLE OMIT
