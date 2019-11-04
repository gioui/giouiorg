---
title: Running on mobiles - Gio documentation
---

## Running on mobiles

For Android, iOS, tvOS the `gogio` tool can build and package a Gio program for you.

### Android

To build an Android .apk file from the `gophers` example:

	$ go run gioui.org/cmd/gogio -target android gioui.org/example/kitchen

The apk can be installed to a running emulator or attached device with adb:

	$ adb install gophers.apk

The `gogio` tool passes command line arguments to os.Args at runtime:

	$ go run gioui.org/cmd/gogio -target android gioui.org/example/gophers -token <github token>


### iOS, tvOS

The `-appid` flag specifies the iOS bundle id or Android package id. The flag is required
for creating signed .ipa files for iOS and tvOS devices, because the bundle id must match an id
previously provisioned in Xcode. For example,

	$ go run gioui.org/cmd/gogio -target ios -appid <bundle-id> gioui.org/example/gophers

Use the `Window->Devices and Simulators` option in Xcode to install the ipa file to the device.
If you have [ideviceinstaller](https://github.com/libimobiledevice/ideviceinstaller) installed,
you can install the app from the command line:

	$ ideviceinstaller -i gophers.ipa

If you just want to run a program on the iOS simulator, use the `-o` flag to specify a .app
directory:

	$ go run gioui.org/cmd/gogio -o gophers.app -target ios gioui.org/example/gophers

Install the app to a running simulator with simctl:

	$ xcrun simctl install booted gophers.app


### App Icon

The `gogio` tool will use the `appicon.png` file in your main package
directory, if present, as the app icon.
