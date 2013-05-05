gocmc
=====

CoreMIDI Channels for Go

#What
The beginnings of a [Go](http://golang.org) library to use CoreMIDI on OSX via Go channels.

This is super early right now (only note on/note off) so you should probably try to use [youpy](https://github.com/youpy)'s
[go-coremidi](https://github.com/youpy/go-coremidi) instead.

#Why
go-coremidi looks pretty good but its approach to output was not immediately clear.  Further, my heavy use of actors
for other stuff makes me think of channels as a decent abstraction for dealing with MIDI input and output.

I've got the rough beginnings of a sequencer coming together on top of this for sharing later.

#Next
Will tackle sysex and input as it's necessary for my use.  As noted above, probably best you use [go-coremidi](https://github.com/youpy/go-coremidi) instead.