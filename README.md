# go-eiscp

This started as a fork of github.com/reddec/go-eiscp, but has mutated into its own thing.

eISCP is a 1980's serial-line protocol with a 1990's update to stream over TCP/IP. It does its job, but not in any modern sort of way. It really expects the listener to keep track of all state and doesn't do well with short-lived (one-shot) commands. I'm trying to work around some of these limitations.

For a small CLI tool, it is still one-shot. But for longer-lived processes I am trying to make persistent connections work more sanely.

For now this is more feature-full than the project I initially forked this from, but I am no where near where I intend to be when I get it working as I envision.
