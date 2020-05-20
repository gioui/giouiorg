// SPDX-License-Identifier: Unlicense OR MIT

package wasm

import (
	_ "gioui.org/example/kitchen"
)

//go:generate gogio -target js -o kitchen gioui.org/example/kitchen
//go:generate gogio -target js -o architecture ../../include/files/architecture
//go:generate go run gioui.org/example/kitchen -screenshot kitchen.png
