package gocmc

/*
#cgo LDFLAGS: -framework CoreMIDI -framework CoreFoundation -framework CoreServices
#include <CoreFoundation/CoreFoundation.h>
#include <CoreMIDI/CoreMIDI.h>
*/
import "C"
import "fmt"

func MakeClient(name string) (C.MIDIClientRef) {
    cfName := strToCfstr(name)
    var client C.MIDIClientRef
    err := C.MIDIClientCreate(cfName, nil, nil, &client)
    //TODO:  actual error handling (DERP)
    if err != C.noErr {
        fmt.Println("Error creating client:  ", int(err))
    }

    return client
}