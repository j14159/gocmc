package gocmc

/*
#cgo LDFLAGS: -framework CoreMIDI -framework CoreFoundation -framework CoreServices
#include <CoreFoundation/CoreFoundation.h>
#include <CoreMIDI/CoreMIDI.h>
*/
import "C"
import "fmt"
import "strings"

type Destination struct {
    Name string
    Endpoint C.MIDIEndpointRef
}

func GetDestinations() (ret map[string] Destination) {
    destCount := C.MIDIGetNumberOfDestinations()
    fmt.Println("Found destination count:  ", destCount)

    ret = make(map[string] Destination)
    var x C.ItemCount = 0
    for ; x < destCount; x++ {
        dest := C.MIDIGetDestination(x)
        var destName C.CFStringRef
        C.MIDIObjectGetStringProperty((C.MIDIObjectRef)(dest), C.kMIDIPropertyName, &destName)
        cleanName := strings.Trim(cstrToStr(destName), "\u0000")
        ret[cleanName] = Destination{ cleanName, dest }
    }

    return
}
