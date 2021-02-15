// SPDX-License-Identifier: Unlicense OR MIT

// Command appletoolchain implements the `xcrun` command and wraps
// `clang` to invoke a cross compiler for macOS and iOS. It uses
// os.Args[0] for selecting the tool.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var sdkRoot string

func main() {
	if err := runMain(); err != nil {
		fmt.Fprintf(os.Stderr, "appletoolchain: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func runMain() error {
	sdkRoot = os.Getenv("APPLE_TOOLCHAIN_ROOT")
	if sdkRoot == "" {
		return errors.New("the APPLE_TOOLCHAIN_ROOT environment variable is not specified")
	}
	tool := filepath.Base(os.Args[0])
	switch tool {
	case "xcrun":
		return xcrun()
	default:
		return runTool(tool)
	}
}

func xcrun() error {
	var (
		sdk      = flag.String("sdk", "", "select an SDK.")
		showPath = flag.Bool("show-sdk-path", false, "display SDK install path")
		find     = flag.Bool("find", false, "display tool path")
	)

	flag.Parse()
	os.Args = flag.Args()
	if *showPath {
		var sdkPath string
		switch *sdk {
		case "iphonesimulator":
			sdkPath = "iPhoneSimulator.sdk"
		case "iphoneos":
			sdkPath = "iPhoneOS.sdk"
		default:
			return fmt.Errorf("unsupported sdk: %s", *sdk)
		}
		fmt.Println(filepath.Join(sdkRoot, sdkPath))
		return nil
	}
	tool := os.Args[0]
	if *find {
		var toolPath string
		switch tool {
		case "":
			return errors.New("no tool specified")
		default:
			toolPath = filepath.Join(sdkRoot, "tools", tool+"-ios")
		}
		fmt.Println(toolPath)
		return nil
	}
	return runTool(tool)
}

func runTool(tool string) error {
	switch tool {
	case "clang-ios", "clang-macos":
		s := strings.Split(tool, "-")
		return clang(s[1])
	case "lipo":
		return lipo()
	default:
		return errors.New("unsupported tool: " + tool)
	}
}

func lipo() error {
	return exe(exec.Command(
		"lipo",
		os.Args[1:]...))
}

func clang(platform string) error {
	var rtlib, defArch, defSDK string
	switch platform {
	case "ios":
		rtlib = "clang_rt.ios"
		defArch = "arm64"
		defSDK = "iPhoneOS.sdk"
	case "macos":
		rtlib = "clang_rt.osx"
		defArch = "x86_64"
		defSDK = "MacOSX.sdk"
	default:
		return fmt.Errorf("unsupported platform: %s", platform)
	}
	clangArgs := os.Args[1:]
	hasArch := false
	hasSysroot := false
	for _, a := range clangArgs {
		switch a {
		case "-arch":
			hasArch = true
		case "-isysroot":
			hasSysroot = true
		}
	}
	clangArgs = append(clangArgs,
		"-target", "unknown-apple-darwin19",
		"-L"+filepath.Join(sdkRoot, "XcodeDefault.xctoolchain/usr/lib/clang/11.0.0/lib/darwin"),
		"-B", filepath.Join(sdkRoot, "tools"),
		// Link the clang runtime for runtime symbols such as
		// __isOSVersionAtLeast.
		"-l"+rtlib,
		"-Wno-unused-command-line-argument",
	)
	if !hasArch {
		clangArgs = append(clangArgs, "-arch", defArch)
	}
	if !hasSysroot {
		clangArgs = append(clangArgs, "-isysroot", filepath.Join(sdkRoot, defSDK))
	}
	cmd := exec.Command(
		"clang",
		clangArgs...)
	toolchain := filepath.Join(sdkRoot, "bin")
	cmd.Env = append(os.Environ(), "COMPILER_PATH="+toolchain)
	return exe(cmd)
}

func exe(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			os.Exit(err.ExitCode())
		}
		return err
	}
	return nil
}
