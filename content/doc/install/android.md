---
title: Android
---

## Dependencies

For Android you need the [Android SDK](https://developer.android.com/studio#command-tools)
with the NDK bundle installed.

Point the `ANDROID_HOME` to the SDK root directory. To install the NDK bundle use
the `sdkmanager` command that comes with the SDK:

    $ sdkmanager ndk-bundle

To run Gio programs on the emulator, you may need to [enable OpenGL ES 3](https://developer.android.com/studio/run/emulator-acceleration).

## Building

Install `gogio`, if you already haven't:

    go install gioui.org/cmd/gogio@latest

To build an Android .apk file from the `kitchen` example:

	$ gogio -target android gioui.org/example/kitchen

The apk can be installed to a running emulator or attached device with adb:

	$ adb install kitchen.apk

The `gogio` tool passes command line arguments to os.Args at runtime:

	$ gogio -target android gioui.org/example/gophers -token <github token>

## Integrate

To build a Gio program as an .aar package, use the `-buildmode=archive` flag. For example,

    $ gogio -target android -buildmode archive gioui.org/example/kitchen

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
