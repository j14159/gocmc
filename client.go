package gocmc

/*
#cgo LDFLAGS: -framework CoreMIDI -framework CoreFoundation -framework CoreServices
#include <CoreFoundation/CoreFoundation.h>
#include <CoreMIDI/CoreMIDI.h>
*/
import "C"
import "fmt"

type Client struct {
    client C.MIDIClientRef
}

func MakeClient(name string) (Client) {
    cfName := strToCfstr(name)
    var client C.MIDIClientRef
    err := C.MIDIClientCreate(cfName, nil, nil, &client)
    //TODO:  actual error handling (DERP)
    if err != C.noErr {
        fmt.Println("Error creating client:  ", int(err))
    }

    return Client{client}
}

func (client Client) NewOutput(name string, midiChannel int, destination CoreMidiEndpoint) (Output) {
    return makeOutput(client.client, name, midiChannel, (C.MIDIEndpointRef)(destination))
}