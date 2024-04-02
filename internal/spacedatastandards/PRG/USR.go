// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package PRG

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type USR struct {
	_tab flatbuffers.Table
}

func GetRootAsUSR(buf []byte, offset flatbuffers.UOffsetT) *USR {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &USR{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUSR(buf []byte, offset flatbuffers.UOffsetT) *USR {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &USR{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *USR) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *USR) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *USR) ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *USR) MESSAGE_TYPES(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *USR) MESSAGE_TYPESLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func USRStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func USRAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ID), 0)
}
func USRAddMESSAGE_TYPES(builder *flatbuffers.Builder, MESSAGE_TYPES flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(MESSAGE_TYPES), 0)
}
func USRStartMESSAGE_TYPESVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func USREnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
