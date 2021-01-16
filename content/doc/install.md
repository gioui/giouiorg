---
title: Installation - Gio documentation
---

# Installation

## Linux {#linux}

For Linux you need Wayland and the wayland, x11, xkbcommon, GLES, EGL, libXcursor development packages.

On Fedora 28 and newer, install the dependencies with the command

    $ dnf install wayland-devel libX11-devel libxkbcommon-x11-devel mesa-libGLES-devel mesa-libEGL-devel libXcursor-devel

On Ubuntu 18.04 and newer, use

    $ apt install libwayland-dev libx11-dev libx11-xcb-dev libxkbcommon-x11-dev libgles2-mesa-dev libegl1-mesa-dev libffi-dev libxcursor-dev

You can build Gio programs without X11 support with the `nox11` build tag, and
without Wayland support with the `nowayland` build tag.

## macOS, iOS, tvOS {#apple}

Xcode is required for Apple platforms.

## Windows {#windows}

To avoid the console appearing when running Gio programs, use the `-H windowsgui` linker flag:

	$ go build -ldflags="-H windowsgui" gioui.org/example/hello

## Android {#android}

For Android you need the [Android SDK](https://developer.android.com/studio#command-tools) with the NDK bundle installed.

Point the ANDROID_HOME to the SDK root directory. To install the NDK bundle use the `sdkmanager`
command that comes with the SDK:

	$ sdkmanager ndk-bundle

To run Gio programs on the emulator, you may need to [enable OpenGL ES 3](https://developer.android.com/studio/run/emulator-acceleration).

## Webassembly/WebGL {#wasm}

To run Gio in a browser you need support for WebAssembly and WebGL.

Building for webassembly requires Go 1.14.
