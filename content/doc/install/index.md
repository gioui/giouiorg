---
title: Install
subtitle: All the dependencies
children:
    - ./linux
    - ./windows
    - ./macos
    - ./android
    - ./ios
    - ./wasm
---

Gio is designed to work with very few dependencies. It depends only on the
platform libraries for window management, input and GPU drawing.

Currently Gio targets the latest released version of [Go](https://golang.org/dl)
in module mode. Earlier versions of Go and `GOPATH` mode might work, but no
effort is made to keep them working.

For desktop builds using `go` tool works directly. For mobile and some
additional desktop feature support, Gio uses a separate tool `gogio`.

To install the latest version of the tool use:

    go install gioui.org/cmd/gogio@latest

For the platforms some additional dependencies may be necessary.

<div class="big-links">
    <a href="/doc/install/linux">Linux</a>
    <a href="/doc/install/windows">Windows</a>
    <a href="/doc/install/macos">macOS</a>
    <a href="/doc/install/android">Android</a>
    <a href="/doc/install/ios">iOS / tvOS</a>
    <a href="/doc/install/wasm">WebAssembly</a>
</div>

## App Icon

The `gogio` tool will use the `appicon.png` file in your main package directory,
if present, as the app icon.
