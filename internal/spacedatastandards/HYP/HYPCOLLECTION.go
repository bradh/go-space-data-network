// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package HYP

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Collection of HYP records
type HYPCOLLECTION struct {
	_tab flatbuffers.Table
}

func GetRootAsHYPCOLLECTION(buf []byte, offset flatbuffers.UOffsetT) *HYPCOLLECTION {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &HYPCOLLECTION{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsHYPCOLLECTION(buf []byte, offset flatbuffers.UOffsetT) *HYPCOLLECTION {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &HYPCOLLECTION{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *HYPCOLLECTION) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HYPCOLLECTION) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *HYPCOLLECTION) RECORDS(obj *HYP, j int) bool {
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

func (rcv *HYPCOLLECTION) RECORDSLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func HYPCOLLECTIONStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func HYPCOLLECTIONAddRECORDS(builder *flatbuffers.Builder, RECORDS flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(RECORDS), 0)
}
func HYPCOLLECTIONStartRECORDSVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func HYPCOLLECTIONEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
