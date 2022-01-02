---
title: FAQ
subtitle: Frequently Asked Questions
---

See also [Common Errors](/doc/learn/common-errors) for usual mistakes.

## What is the difference between Gio and Gomobile?

[Go Mobile](https://github.com/golang/mobile) produces either standalone programs
(`gomobile build`) or libraries suitable for calling from Java or Objective-C/Swift
(`gomobile bind`).

`gomobile build` is similar to using the [gioui.org/app](https://gioui.org/app)
package and the [gioui.org/cmd/gogio](https://gioui.org/cmd/gogio) tool to produce an
Android or iOS app. The difference is the abstraction level: `gomobile build` programs
have access to a raw OpenGL ES context while the Gio `app` package exposes a higher
level drawing interface. Gomobile programs also don't have any GUI packages available. 

`gomobile bind` exports a set of Go packages for convenient access from Java or
Objective-C/Swift code. There is no counterpart in Gio, and could be used for
interfacing with native code from Gio programs.

## Why Sourcehut?

The most important feature of Sourcehut is that email contributions are first
class citizens. Email has many problems, but it is permissionless, ubiquitous and
decentralized. For example, people banned from GitHub can contribute patches,
use the mailing list and file issues if they can send and receive email.

Second, Sourcehut encourages using Git in a more decentralized way. There is a
canonical Gio repository, but contributors can work on their changes in ways
that suit them. For example, local clones for smaller changes or pushing to a
Git host of their choice for larger changes. The Sourcehut author wrote a [blog
post](https://drewdevault.com/2019/05/24/What-is-a-fork.html) about how GitHub
changed the meaning of forks and pull requests to be more self-serving and
centralized.

Finally, the code for Sourcehut itself is open source, keeping it honest by the
threat of self-hosting or even a complete fork. Project owners pay for hosting,
lowering the incentive to extract indirect value from users and keeps the
feature set focused on its projects.

Note that even ignoring the above arguments, there is not a clear alternative.
For example, the Go project itself supports GitHub contributions, but only as a
bridge to its preferred code review tool Gerrit.
