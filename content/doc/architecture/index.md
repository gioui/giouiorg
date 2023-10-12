---
title: Architecture
subtitle: Internals of Gio
after: ./window
children:
    - ./window
    - ./drawing
    - ./input
    - ./widget
    - ./layout
    - ./theme
    - ./units
    - ./text
    - ./color
---

Gio is a library for implementing [immediate mode user interfaces](https://eliasnaur.com/blog/immediate-mode-gui-programming). This approach can be implemented in multiple ways, however the overarching similarity is that the program:

1. Listens for events such as mouse or keyboard input.
2. Updates its internal state based on the event.
3. Runs code that lays out and redraws the user interface state.

A minimal immediate mode command-line UI in pseudo-code:

```
main() {
	checked = false
	for every keypress {
		clear screen
		layoutCheckbox(keypress, &checked)
		if checked {
			print("info")
		}
	}
}

layoutCheckbox(keypress, checked) {
	if keypress == SPACE {
		*checked = !*checked
	}

	if *checked {
		print("[x]")
	} else {
		print("[ ]")
	}
}
```

In the immediate mode model, the program is in control of clearing and updating the display, and directly draws widgets and handles input during the updates.

In contrast, traditional "retained mode" libraries own the widgets through implicit library-managed state, typically arranged in a tree-like structure such as a
browser's [DOM](https://en.wikipedia.org/wiki/Document_Object_Model). As a result, the program must use the facilities given by the library to manipulate
its widgets.

Actual GUI programming has several concerns in addition to the simple example above:

1. How to get the events?
2. When to redraw the state?
3. What do the widget structures look like?
4. How to track the focus?
5. How to structure the events?
6. How to communicate with the graphics card?
7. How to handle input?
8. How to draw text?
9. Where does the widget state belong?
10. And many more.

The rest of this document tries to answer how Gio does it. If you wish to know more about immediate mode UI, these references are a good start:

* https://caseymuratori.com/blog_0001
* http://sol.gfxile.net/imgui/
* http://www.johno.se/book/imgui.html
* https://github.com/ocornut/imgui
* https://eliasnaur.com/blog/immediate-mode-gui-programming
