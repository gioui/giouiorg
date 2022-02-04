---
title: Sointu
subtitle: A modular software synthesizer to produce music for 4k intros.
images:
    -
        source: ./1.png
links:
    -
        title: Presentation
        url: https://youtu.be/fokSSaz3mbs?t=349
    -
        title: Source
        url: https://github.com/vsariola/sointu
---

Sointu is a fork and an evolution of 4klang, a modular software synthesizer intended to easily produce music for 4k intros — small executables with a maximum filesize of 4096 bytes containing realtime audio and visuals. Like 4klang, the sound is produced by a virtual machine that executes small bytecode to produce the audio; however, by now the internal virtual machine has been heavily rewritten and extended. It is actually extended so much that you will never fit all the features at the same time in a 4k intro, but a fairly capable synthesis engine can already be fitted in 600 bytes (386, compressed), with another few hundred bytes for the patch and pattern data.