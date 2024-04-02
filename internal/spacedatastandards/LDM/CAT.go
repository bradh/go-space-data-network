// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package LDM

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Catalog Entity Message
type CAT struct {
	_tab flatbuffers.Table
}

func GetRootAsCAT(buf []byte, offset flatbuffers.UOffsetT) *CAT {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CAT{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCAT(buf []byte, offset flatbuffers.UOffsetT) *CAT {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CAT{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CAT) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CAT) Table() flatbuffers.Table {
	return rcv._tab
}

/// Satellite Name(s)
func (rcv *CAT) OBJECT_NAME() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Satellite Name(s)
/// International Designator (YYYY-NNNAAA)
func (rcv *CAT) OBJECT_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// International Designator (YYYY-NNNAAA)
/// NORAD Catalog Number
func (rcv *CAT) NORAD_CAT_ID() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// NORAD Catalog Number
func (rcv *CAT) MutateNORAD_CAT_ID(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

/// Object type (Payload, Rocket body, Debris, Unknown)
func (rcv *CAT) OBJECT_TYPE() objectType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return objectType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 3
}

/// Object type (Payload, Rocket body, Debris, Unknown)
func (rcv *CAT) MutateOBJECT_TYPE(n objectType) bool {
	return rcv._tab.MutateInt8Slot(10, int8(n))
}

/// Operational Status Code
func (rcv *CAT) OPS_STATUS_CODE() opsStatusCode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return opsStatusCode(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 7
}

/// Operational Status Code
func (rcv *CAT) MutateOPS_STATUS_CODE(n opsStatusCode) bool {
	return rcv._tab.MutateInt8Slot(12, int8(n))
}

/// Ownership, typically country or company
func (rcv *CAT) OWNER() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Ownership, typically country or company
/// Launch Date [year-month-day] (ISO 8601)
func (rcv *CAT) LAUNCH_DATE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Launch Date [year-month-day] (ISO 8601)
/// Launch Site
func (rcv *CAT) LAUNCH_SITE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Launch Site
/// Decay Date, if applicable [year-month-day] (ISO 8601)
func (rcv *CAT) DECAY_DATE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Decay Date, if applicable [year-month-day] (ISO 8601)
/// Orbital period [minutes]
func (rcv *CAT) PERIOD() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Orbital period [minutes]
func (rcv *CAT) MutatePERIOD(n float64) bool {
	return rcv._tab.MutateFloat64Slot(22, n)
}

/// Inclination [degrees]
func (rcv *CAT) INCLINATION() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Inclination [degrees]
func (rcv *CAT) MutateINCLINATION(n float64) bool {
	return rcv._tab.MutateFloat64Slot(24, n)
}

/// Apogee Altitude [kilometers]
func (rcv *CAT) APOGEE() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Apogee Altitude [kilometers]
func (rcv *CAT) MutateAPOGEE(n float64) bool {
	return rcv._tab.MutateFloat64Slot(26, n)
}

/// Perigee Altitude [kilometers]
func (rcv *CAT) PERIGEE() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Perigee Altitude [kilometers]
func (rcv *CAT) MutatePERIGEE(n float64) bool {
	return rcv._tab.MutateFloat64Slot(28, n)
}

/// Radar Cross Section [meters2]; blank if no data available
func (rcv *CAT) RCS() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Radar Cross Section [meters2]; blank if no data available
func (rcv *CAT) MutateRCS(n float64) bool {
	return rcv._tab.MutateFloat64Slot(30, n)
}

/// Data status code; blank otherwise
func (rcv *CAT) DATA_STATUS_CODE() dataStatusCode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return dataStatusCode(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Data status code; blank otherwise
func (rcv *CAT) MutateDATA_STATUS_CODE(n dataStatusCode) bool {
	return rcv._tab.MutateInt8Slot(32, int8(n))
}

/// Orbit center
func (rcv *CAT) ORBIT_CENTER() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Orbit center
/// Orbit type (Orbit, Landing, Impact, Docked to RSO, roundtrip)
func (rcv *CAT) ORBIT_TYPE() orbitType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return orbitType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Orbit type (Orbit, Landing, Impact, Docked to RSO, roundtrip)
func (rcv *CAT) MutateORBIT_TYPE(n orbitType) bool {
	return rcv._tab.MutateInt8Slot(36, int8(n))
}

/// Deployment Date [year-month-day] (ISO 8601)
func (rcv *CAT) DEPLOYMENT_DATE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Deployment Date [year-month-day] (ISO 8601)
/// Indicates if the object is maneuverable
func (rcv *CAT) MANEUVERABLE() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Indicates if the object is maneuverable
func (rcv *CAT) MutateMANEUVERABLE(n bool) bool {
	return rcv._tab.MutateBoolSlot(40, n)
}

/// Size [meters]; blank if no data available
func (rcv *CAT) SIZE() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Size [meters]; blank if no data available
func (rcv *CAT) MutateSIZE(n float64) bool {
	return rcv._tab.MutateFloat64Slot(42, n)
}

/// Mass [kilograms]; blank if no data available
func (rcv *CAT) MASS() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Mass [kilograms]; blank if no data available
func (rcv *CAT) MutateMASS(n float64) bool {
	return rcv._tab.MutateFloat64Slot(44, n)
}

/// Mass type (Dry, Wet)
func (rcv *CAT) MASS_TYPE() massType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return massType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Mass type (Dry, Wet)
func (rcv *CAT) MutateMASS_TYPE(n massType) bool {
	return rcv._tab.MutateInt8Slot(46, int8(n))
}

/// Vector of PAYLOADS
func (rcv *CAT) PAYLOADS(obj *PLD, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *CAT) PAYLOADSLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Vector of PAYLOADS
func CATStart(builder *flatbuffers.Builder) {
	builder.StartObject(23)
}
func CATAddOBJECT_NAME(builder *flatbuffers.Builder, OBJECT_NAME flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(OBJECT_NAME), 0)
}
func CATAddOBJECT_ID(builder *flatbuffers.Builder, OBJECT_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(OBJECT_ID), 0)
}
func CATAddNORAD_CAT_ID(builder *flatbuffers.Builder, NORAD_CAT_ID uint32) {
	builder.PrependUint32Slot(2, NORAD_CAT_ID, 0)
}
func CATAddOBJECT_TYPE(builder *flatbuffers.Builder, OBJECT_TYPE objectType) {
	builder.PrependInt8Slot(3, int8(OBJECT_TYPE), 3)
}
func CATAddOPS_STATUS_CODE(builder *flatbuffers.Builder, OPS_STATUS_CODE opsStatusCode) {
	builder.PrependInt8Slot(4, int8(OPS_STATUS_CODE), 7)
}
func CATAddOWNER(builder *flatbuffers.Builder, OWNER flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(OWNER), 0)
}
func CATAddLAUNCH_DATE(builder *flatbuffers.Builder, LAUNCH_DATE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(LAUNCH_DATE), 0)
}
func CATAddLAUNCH_SITE(builder *flatbuffers.Builder, LAUNCH_SITE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(LAUNCH_SITE), 0)
}
func CATAddDECAY_DATE(builder *flatbuffers.Builder, DECAY_DATE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(DECAY_DATE), 0)
}
func CATAddPERIOD(builder *flatbuffers.Builder, PERIOD float64) {
	builder.PrependFloat64Slot(9, PERIOD, 0.0)
}
func CATAddINCLINATION(builder *flatbuffers.Builder, INCLINATION float64) {
	builder.PrependFloat64Slot(10, INCLINATION, 0.0)
}
func CATAddAPOGEE(builder *flatbuffers.Builder, APOGEE float64) {
	builder.PrependFloat64Slot(11, APOGEE, 0.0)
}
func CATAddPERIGEE(builder *flatbuffers.Builder, PERIGEE float64) {
	builder.PrependFloat64Slot(12, PERIGEE, 0.0)
}
func CATAddRCS(builder *flatbuffers.Builder, RCS float64) {
	builder.PrependFloat64Slot(13, RCS, 0.0)
}
func CATAddDATA_STATUS_CODE(builder *flatbuffers.Builder, DATA_STATUS_CODE dataStatusCode) {
	builder.PrependInt8Slot(14, int8(DATA_STATUS_CODE), 0)
}
func CATAddORBIT_CENTER(builder *flatbuffers.Builder, ORBIT_CENTER flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(15, flatbuffers.UOffsetT(ORBIT_CENTER), 0)
}
func CATAddORBIT_TYPE(builder *flatbuffers.Builder, ORBIT_TYPE orbitType) {
	builder.PrependInt8Slot(16, int8(ORBIT_TYPE), 0)
}
func CATAddDEPLOYMENT_DATE(builder *flatbuffers.Builder, DEPLOYMENT_DATE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(17, flatbuffers.UOffsetT(DEPLOYMENT_DATE), 0)
}
func CATAddMANEUVERABLE(builder *flatbuffers.Builder, MANEUVERABLE bool) {
	builder.PrependBoolSlot(18, MANEUVERABLE, false)
}
func CATAddSIZE(builder *flatbuffers.Builder, SIZE float64) {
	builder.PrependFloat64Slot(19, SIZE, 0.0)
}
func CATAddMASS(builder *flatbuffers.Builder, MASS float64) {
	builder.PrependFloat64Slot(20, MASS, 0.0)
}
func CATAddMASS_TYPE(builder *flatbuffers.Builder, MASS_TYPE massType) {
	builder.PrependInt8Slot(21, int8(MASS_TYPE), 0)
}
func CATAddPAYLOADS(builder *flatbuffers.Builder, PAYLOADS flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(22, flatbuffers.UOffsetT(PAYLOADS), 0)
}
func CATStartPAYLOADSVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CATEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
