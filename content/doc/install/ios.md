---
title: iOS, tvOS
---

## Dependencies

Xcode is required for Apple platforms.

## Building

Install `gogio`, if you already haven't:

``` sh
go install gioui.org/cmd/gogio@latest
```

The `-appid` flag specifies the iOS bundle id or Android package id. The flag is
required for creating signed .ipa files for iOS and tvOS devices, because the
bundle id must match an id previously provisioned in Xcode. For example,

``` sh
gogio -target ios -appid <bundle-id>	gioui.org/example/kitchen
```

Use the `Window->Devices and Simulators` option in Xcode to install the ipa file
to the device. If you have [ideviceinstaller](https://github.com/libimobiledevice/ideviceinstaller) installed, you can
install the app from the command line:

``` sh
ideviceinstaller -i kitchen.ipa
```

If you just want to run a program on the iOS simulator, use the `-o` flag to
specify a .app directory:

``` sh
gogio -o kitchen.app -target ios	gioui.org/example/kitchen
```

Install the app to a running simulator with simctl:

``` sh
xcrun simctl install booted kitchen.app
```

## Integrate

The `gogio` tool can also produce a framework ready to include in an iOS or tvOS
Xcode project. The command

``` sh
gogio -target ios -buildmode archive gioui.org/example/kitchen
```

outputs Kitchen.framework with the demo program built for iOS.

To run the Gio program, use the GioAppDelegate class from your program:

``` objc
@import UIKit;
@import Gio;

int main(int argc, char * argv[]) {
	@autoreleasepool {
		return UIApplicationMain(argc, argv, nil, NSStringFromClass([GioAppDelegate class]));
	}
}
```
