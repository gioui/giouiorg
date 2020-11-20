---
title: Contributing
---

Commit messages follow [the Go project style](https://golang.org/doc/contribute.html#commit_messages):
the first line is prefixed with the package and a short summary. The rest of the message provides context
for the change and what it does. See
[an example](https://gioui.org/commit/abb9d291e954f3b80384046d7d4487e1ead6bd6a).
Add `Fixes gio#nnn` or `Updates gio#nnn` if the change fixes or updates an existing
issue.

Contributors must agree to the [developer certificate of origin](https://developercertificate.org/),
to ensure their work is compatible with the MIT license and the UNLICENSE. Sign your commits with Signed-off-by
statements to show your agreement. The `git commit --signoff` (or `-s`) command signs a commit with
your name and email address.

Patches should be sent to
[~eliasnaur/gio-patches@lists.sr.ht](mailto:~eliasnaur/gio-patches@lists.sr.ht)
mailing list with the `git send-email` command. See
[git-send-email.io](https://git-send-email.io) for a thorough setup guide.

If you have a [sourcehut](https://sr.ht) account, you can also fork
the Gio repository, push your changes to that and use the web-based
flow for emailing the patch. Start the process by click the "Prepare a
patchset" button on the front page of your fork.


## git send-email setup

With `git send-email` configured, you can clone the project and set it up for submitting your changes:

    $ git clone https://git.sr.ht/~eliasnaur/gio
    $ cd gio
    $ git config sendemail.to '~eliasnaur/gio-patches@lists.sr.ht'
    $ git config sendemail.annotate yes

Include the project name in the mail subject:

	$ git config format.subjectPrefix "PATCH gio"

If you're contributing for a different project, replace "gio" with the
repository name from https://sr.ht/~eliasnaur/gio/sources. For
example, "giouiorg" is the repository for the
[gioui.org](https://gioui.org) website.

Configure your name and email address if you have not done so already:

    $ git config --global user.email "you@example.com"
    $ git config --global user.name "Your Name"

Whenever you want to submit your work for review, use `git send-email` with the base revision of your
changes. For example, to submit the most recent commit use

    $ git send-email HEAD^

## Automatic patch testing

Patches with the project name "gio" in the subject will be picked up by the automatic testers at
[builds.sr.ht](https://builds.sr.ht). A report with the testing results will be sent to you, CC'ed
to the mailing list.
