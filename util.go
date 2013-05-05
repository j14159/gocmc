package gocmc

/*
#cgo LDFLAGS: -framework CoreMIDI -framework CoreFoundation -framework CoreServices
#include <CoreFoundation/CoreFoundation.h>
#include <CoreMIDI/CoreMIDI.h>
*/
import "C"
import "unsafe"

func cstrToStr(cstr C.CFStringRef) (goString string) {
    len := C.CFStringGetLength(cstr)
    buf := make([]byte, len * 2)
    cbuf := (*C.char)(unsafe.Pointer(&buf[0]))
    C.CFStringGetCString(cstr, cbuf, len * 2, C.CFStringGetSystemEncoding())
    return string(buf)
}

/*
Based on youpy's version in go-coremidi
*/
func strToCfstr(str string) (C.CFStringRef) {
    cStr := C.CString(str)
    cfStr := C.CFStringCreateWithCString(nil, cStr, C.kCFStringEncodingMacRoman)
    /*
    The next two lines cause C.MIDICreateClient to hang in client.go's MakeClient method.
    Removing them makes things work just fine.  Potential memory leak by not freeing these?
    Need to understand better.  youpy has all of this in a callback.
    */
    //defer C.free(unsafe.Pointer(cStr))
    //defer C.CFRelease((C.CFTypeRef)(cfStr))

    return cfStr
}