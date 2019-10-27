---
title: Installation
---

# Installation

## Linux {#linux}

For Linux you need Wayland and the wayland, x11, xkbcommon, GLES, EGL development packages.

On Fedora 28 and newer, install the dependencies with the command

    $ sudo dnf install wayland-devel libX11-devel libxkbcommon-devel mesa-libGLES-devel mesa-libEGL-devel

On Ubuntu 18.04 and newer, use

    $ sudo apt install libwayland-dev libx11-dev libxkbcommon-dev libgles2-mesa-dev libegl1-mesa-dev

You can build Gio programs without X11 support with the `nox11` build tag, and
without Wayland support with the `nowayland` build tag.

## macOS, iOS, tvOS {#apple}

Xcode is required for Apple platforms.

Building for tvOS requires Go 1.13.

## Windows {#windows}

For Windows you need the ANGLE drivers for emulating OpenGL ES.

You can build ANGLE yourself or use
[a prebuilt version](https://drive.google.com/file/d/1k2950mHNtR2iwhweHS1rJ7reChTa3rki/view?usp=sharing).
Leave the DLLs in the same directory as the Gio program.

To avoid the console appearing when running Gio programs, use the `-H windowsgui` linker flag:

	$ go build -ldflags="-H windowsgui" gioui.org/example/hello

## Android {#android}

For Android you need the [Android SDK](https://developer.android.com/studio#command-tools) with the NDK bundle installed.

Point the ANDROID_HOME to the SDK root directory. To install the NDK bundle use the `sdkmanager`
command that comes with the SDK:

	$ sdkmanager ndk-bundle

## Webassembly/WebGL {#wasm}

To run Gio in a browser you need support for WebAssembly and WebGL.

Building for webassembly requires Go 1.13.
