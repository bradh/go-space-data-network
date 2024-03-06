// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package HYP

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Score struct {
	_tab flatbuffers.Table
}

func GetRootAsScore(buf []byte, offset flatbuffers.UOffsetT) *Score {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Score{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsScore(buf []byte, offset flatbuffers.UOffsetT) *Score {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Score{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Score) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Score) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Score) NORAD_CAT_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Score) TYPE() ScoreType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return ScoreType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Score) MutateTYPE(n ScoreType) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func (rcv *Score) TAG() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Score) SCORE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Score) MutateSCORE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func ScoreStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ScoreAddNORAD_CAT_ID(builder *flatbuffers.Builder, NORAD_CAT_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(NORAD_CAT_ID), 0)
}
func ScoreAddTYPE(builder *flatbuffers.Builder, TYPE ScoreType) {
	builder.PrependInt8Slot(1, int8(TYPE), 0)
}
func ScoreAddTAG(builder *flatbuffers.Builder, TAG flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(TAG), 0)
}
func ScoreAddSCORE(builder *flatbuffers.Builder, SCORE float32) {
	builder.PrependFloat32Slot(3, SCORE, 0.0)
}
func ScoreEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
