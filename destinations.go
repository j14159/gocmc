package gocmc

/*
#cgo LDFLAGS: -framework CoreMIDI -framework CoreFoundation -framework CoreServices
#include <CoreFoundation/CoreFoundation.h>
#include <CoreMIDI/CoreMIDI.h>
*/
import "C"
import "fmt"

type Destination struct {
    Name string
    Endpoint C.MIDIEndpointRef
}

func GetDestinations() (ret map[int] Destination) {
    destCount := C.MIDIGetNumberOfDestinations()
    fmt.Println("Found destination count:  ", destCount)

    ret = make(map[int] Destination)

    var x C.ItemCount = 0
    for ; x < destCount; x++ {
        dest := C.MIDIGetDestination(x)
        var destName C.CFStringRef
        C.MIDIObjectGetStringProperty((C.MIDIObjectRef)(dest), C.kMIDIPropertyName, &destName)
        ret[int(x)] = Destination{ cstrToStr(destName), dest }
    }

    return
}
