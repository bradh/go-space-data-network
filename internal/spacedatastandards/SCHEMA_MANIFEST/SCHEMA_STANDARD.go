// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package SCHEMA_MANIFEST

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Schema Standard Definition
type SCHEMA_STANDARD struct {
	_tab flatbuffers.Table
}

func GetRootAsSCHEMA_STANDARD(buf []byte, offset flatbuffers.UOffsetT) *SCHEMA_STANDARD {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SCHEMA_STANDARD{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSCHEMA_STANDARD(buf []byte, offset flatbuffers.UOffsetT) *SCHEMA_STANDARD {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SCHEMA_STANDARD{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SCHEMA_STANDARD) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SCHEMA_STANDARD) Table() flatbuffers.Table {
	return rcv._tab
}

/// Unique identifier for the standard
func (rcv *SCHEMA_STANDARD) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Unique identifier for the standard
/// IDL
func (rcv *SCHEMA_STANDARD) Idl() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// IDL
/// List Of File Paths
func (rcv *SCHEMA_STANDARD) Files(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *SCHEMA_STANDARD) FilesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// List Of File Paths
func SCHEMA_STANDARDStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func SCHEMA_STANDARDAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func SCHEMA_STANDARDAddIdl(builder *flatbuffers.Builder, idl flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(idl), 0)
}
func SCHEMA_STANDARDAddFiles(builder *flatbuffers.Builder, files flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(files), 0)
}
func SCHEMA_STANDARDStartFilesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SCHEMA_STANDARDEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
