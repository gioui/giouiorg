---
title: Windows
---

## Dependencies

The default windows setup does not require extra dependencies.

<!-- TODO mention special requirements for glfw -->

## Building

To test whether the installation works, run:

``` sh
go run gioui.org/example/hello@latest
```

### Avoiding console

To avoid the console appearing when running Gio programs, use the `-H windowsgui` linker flag:

``` sh
go run -ldflags="-H windowsgui" gioui.org/example/hello@latest
```