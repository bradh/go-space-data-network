// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package PNM

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Collection of Publish Notification Messages
/// This table groups multiple PNM records for batch processing and management.
type PNMCOLLECTION struct {
	_tab flatbuffers.Table
}

func GetRootAsPNMCOLLECTION(buf []byte, offset flatbuffers.UOffsetT) *PNMCOLLECTION {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PNMCOLLECTION{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPNMCOLLECTION(buf []byte, offset flatbuffers.UOffsetT) *PNMCOLLECTION {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PNMCOLLECTION{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PNMCOLLECTION) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PNMCOLLECTION) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PNMCOLLECTION) RECORDS(obj *PNM, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *PNMCOLLECTION) RECORDSLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func PNMCOLLECTIONStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func PNMCOLLECTIONAddRECORDS(builder *flatbuffers.Builder, RECORDS flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(RECORDS), 0)
}
func PNMCOLLECTIONStartRECORDSVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PNMCOLLECTIONEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
