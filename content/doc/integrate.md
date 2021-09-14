---
title: Integrate with Mobile and Browser
---

Gio includes very basic support for integrating with existing mobile and browser projects.

## Android

To build a Gio program as an .aar package, use the `-buildmode=archive` flag. For example,

    $ go run gioui.org/cmd/gogio -target android -buildmode archive gioui.org/example/kitchen

produces kitchen.aar, ready to include in an Android project.

To display the Gio Android Activity, declare it in your AndroidManifest.xml:

	<?xml version="1.0" encoding="utf-8"?>
	<manifest xmlns:android="http://schemas.android.com/apk/res/android">
		...
		<uses-sdk android:minSdkVersion="16" android:targetSdkVersion="28" />
		<uses-feature android:glEsVersion="0x00030000"/>
		...
		<application android:label="Gio">
			<activity android:name="org.gioui.GioActivity"
			  android:theme="@style/Theme.GioApp"
			  android:configChanges="orientation|keyboardHidden"
			  android:windowSoftInputMode="adjustResize">
			</activity>
		</application>
		...
	</manifest>

and launch it from another Activity with

	startActivity(new Intent(this, GioActivity.class));


## iOS/tvOS

The `gogio` tool can also produce a framework ready to include in an iOS or tvOS Xcode project.
The command

    $ go run gioui.org/cmd/gogio -target ios -buildmode archive gioui.org/example/kitchen

outputs Kitchen.framework with the demo program built for iOS.

To run the Gio program, use the GioAppDelegate class from your program:

	@import UIKit;
	@import Gio;

	int main(int argc, char * argv[]) {
		@autoreleasepool {
			return UIApplicationMain(argc, argv, nil, NSStringFromClass([GioAppDelegate class]));
		}
	}


## WebAssembly

If the embedding HTML page for the Gio program contains a `<div id="giowindow">` element, Gio
will run in that instead of creating its own container.
