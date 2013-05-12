package gocmc

/*
#cgo LDFLAGS: -framework CoreMIDI -framework CoreFoundation -framework CoreServices
#include <CoreFoundation/CoreFoundation.h>
#include <CoreMIDI/CoreMIDI.h>

//
//
// MIDIPacket and packet list functions here because C structs
// in Go are all kinds of not awesome.
//
//

//Helper for MIDI messages using 3 bytes (cmd + chan, note on/of, etc and velocity)
//TODO:  generalize this more, feels like a hack right now.
MIDIPacketList _3PacketList(int command, int note, int channel, int velocity) {
    MIDIPacketList packetList;
    
    packetList.numPackets = 1;
    
    MIDIPacket* firstPacket = &packetList.packet[0];
    
    firstPacket->timeStamp = 0; // send immediately
    firstPacket->length = 3;
    firstPacket->data[0] = command + channel; //0x90 + channel;
    firstPacket->data[1] = note;
    firstPacket->data[2] = velocity;
    
    // TODO: add end note sequence
    return packetList;
}

MIDIPacketList MidiNoteOn(int note, int channel, int velocity){
    return _3PacketList(9 << 4, note, channel, velocity);
}

MIDIPacketList MidiNoteOff(int note, int channel, int velocity) {
    return _3PacketList(8 << 4, note, channel, velocity);
}
*/
import "C"

type Output struct {
    cmPort C.MIDIPortRef
    midiChannel int
    destination C.MIDIEndpointRef
    notesIn chan NoteEvent
}

type NoteEvent struct {
    On bool
    Note int
    Velocity int
}

func makeOutput(client C.MIDIClientRef, name string, midiChannel int, destination C.MIDIEndpointRef) (ret Output) {
    var outPort C.MIDIPortRef
    C.MIDIOutputPortCreate(client, strToCfstr(name), &outPort)
    ret = Output{outPort, midiChannel, destination, make(chan NoteEvent, 100)}
    go outputHandler(ret)
    return
}

func (output Output) EventOut(ne NoteEvent) {
    output.notesIn <- ne
}

func outputHandler(output Output) {
    //maybe gross to do infinite loop here, possibly need start/stop channel:
    for {
        nextNote := <-output.notesIn
        handleNoteEvent(nextNote, output)
    }
}

func handleNoteEvent(n NoteEvent, output Output) {
    var packetList C.MIDIPacketList
    if n.On {
        packetList = C.MidiNoteOn((C.int)(n.Note), (C.int)(output.midiChannel), (C.int)(n.Velocity))
    } else {
        packetList = C.MidiNoteOff((C.int)(n.Note), (C.int)(output.midiChannel), (C.int)(0))
    }

    C.MIDISend(output.cmPort, output.destination, &packetList)   
}
