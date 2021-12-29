---
title: WebAssembly
---

## Building

Install `gogio`, if you already haven't:

    go install gioui.org/cmd/gogio@latest

To build WebAssembly from the `kitchen` example:

	$ gogio -target js gioui.org/example/kitchen

This will create an `index.html`, `.wasm` and `.js` needed to start up the
project inside a browser. These need to be served as a website, directly opening
the `index.html` will not work.

One way to quickly setup a server is to use:

    $ go get github.com/shurcooL/goexec
	$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("kitchen")))'

Open http://localhost:8080 in a browser to run the program.

## Integrate

If the embedding HTML page for the Gio program contains a `<div id="giowindow">`
element, Gio will run in that instead of creating its own container.
