// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package PLD

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Integrated Device Message
type IDM struct {
	_tab flatbuffers.Table
}

func GetRootAsIDM(buf []byte, offset flatbuffers.UOffsetT) *IDM {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IDM{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIDM(buf []byte, offset flatbuffers.UOffsetT) *IDM {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IDM{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IDM) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IDM) Table() flatbuffers.Table {
	return rcv._tab
}

/// Unique identifier for the EMT
func (rcv *IDM) ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Unique identifier for the EMT
/// Name of the EMT
func (rcv *IDM) NAME() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of the EMT
/// Mode of the data (real, simulated, synthetic)
func (rcv *IDM) DATA_MODE() DataMode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return DataMode(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Mode of the data (real, simulated, synthetic)
func (rcv *IDM) MutateDATA_MODE(n DataMode) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

/// Uplink frequency range
func (rcv *IDM) UPLINK(obj *FrequencyRange) *FrequencyRange {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FrequencyRange)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Uplink frequency range
/// Downlink frequency range
func (rcv *IDM) DOWNLINK(obj *FrequencyRange) *FrequencyRange {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FrequencyRange)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Downlink frequency range
/// Beacon frequency range
func (rcv *IDM) BEACON(obj *FrequencyRange) *FrequencyRange {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FrequencyRange)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Beacon frequency range
/// Bands associated with the EMT
func (rcv *IDM) BAND(obj *Band, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *IDM) BANDLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Bands associated with the EMT
/// Type of polarization used
func (rcv *IDM) POLARIZATION_TYPE() PolarizationType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return PolarizationType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Type of polarization used
func (rcv *IDM) MutatePOLARIZATION_TYPE(n PolarizationType) bool {
	return rcv._tab.MutateInt8Slot(18, int8(n))
}

/// Simple polarization configuration
func (rcv *IDM) SIMPLE_POLARIZATION() SimplePolarization {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return SimplePolarization(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Simple polarization configuration
func (rcv *IDM) MutateSIMPLE_POLARIZATION(n SimplePolarization) bool {
	return rcv._tab.MutateInt8Slot(20, int8(n))
}

/// Stokes parameters for polarization characterization
func (rcv *IDM) STOKES_PARAMETERS(obj *StokesParameters) *StokesParameters {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(StokesParameters)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Stokes parameters for polarization characterization
/// Power required in Watts
func (rcv *IDM) POWER_REQUIRED() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Power required in Watts
func (rcv *IDM) MutatePOWER_REQUIRED(n float64) bool {
	return rcv._tab.MutateFloat64Slot(24, n)
}

/// Type of power (eg. AC or DC)
func (rcv *IDM) POWER_TYPE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Type of power (eg. AC or DC)
/// Indicates if the EMT can transmit
func (rcv *IDM) TRANSMIT() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Indicates if the EMT can transmit
func (rcv *IDM) MutateTRANSMIT(n bool) bool {
	return rcv._tab.MutateBoolSlot(28, n)
}

/// Indicates if the EMT can receive
func (rcv *IDM) RECEIVE() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Indicates if the EMT can receive
func (rcv *IDM) MutateRECEIVE(n bool) bool {
	return rcv._tab.MutateBoolSlot(30, n)
}

/// Type of the sensor
func (rcv *IDM) SENSOR_TYPE() DeviceType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return DeviceType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Type of the sensor
func (rcv *IDM) MutateSENSOR_TYPE(n DeviceType) bool {
	return rcv._tab.MutateInt8Slot(32, int8(n))
}

/// Source of the data
func (rcv *IDM) SOURCE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Source of the data
/// Timestamp of the last observation
func (rcv *IDM) LAST_OB_TIME() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Timestamp of the last observation
/// Lower left elevation limit
func (rcv *IDM) LOWER_LEFT_ELEVATION_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Lower left elevation limit
func (rcv *IDM) MutateLOWER_LEFT_ELEVATION_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(38, n)
}

/// Upper left azimuth limit
func (rcv *IDM) UPPER_LEFT_AZIMUTH_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Upper left azimuth limit
func (rcv *IDM) MutateUPPER_LEFT_AZIMUTH_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(40, n)
}

/// Lower right elevation limit
func (rcv *IDM) LOWER_RIGHT_ELEVATION_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Lower right elevation limit
func (rcv *IDM) MutateLOWER_RIGHT_ELEVATION_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(42, n)
}

/// Lower left azimuth limit
func (rcv *IDM) LOWER_LEFT_AZIMUTH_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Lower left azimuth limit
func (rcv *IDM) MutateLOWER_LEFT_AZIMUTH_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(44, n)
}

/// Upper right elevation limit
func (rcv *IDM) UPPER_RIGHT_ELEVATION_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Upper right elevation limit
func (rcv *IDM) MutateUPPER_RIGHT_ELEVATION_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(46, n)
}

/// Upper right azimuth limit
func (rcv *IDM) UPPER_RIGHT_AZIMUTH_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Upper right azimuth limit
func (rcv *IDM) MutateUPPER_RIGHT_AZIMUTH_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(48, n)
}

/// Lower right azimuth limit
func (rcv *IDM) LOWER_RIGHT_AZIMUTH_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Lower right azimuth limit
func (rcv *IDM) MutateLOWER_RIGHT_AZIMUTH_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(50, n)
}

/// Upper left elevation limit
func (rcv *IDM) UPPER_LEFT_ELEVATION_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Upper left elevation limit
func (rcv *IDM) MutateUPPER_LEFT_ELEVATION_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(52, n)
}

/// Right geostationary belt limit
func (rcv *IDM) RIGHT_GEO_BELT_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Right geostationary belt limit
func (rcv *IDM) MutateRIGHT_GEO_BELT_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(54, n)
}

/// Left geostationary belt limit
func (rcv *IDM) LEFT_GEO_BELT_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Left geostationary belt limit
func (rcv *IDM) MutateLEFT_GEO_BELT_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(56, n)
}

/// Magnitude limit of the sensor
func (rcv *IDM) MAGNITUDE_LIMIT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Magnitude limit of the sensor
func (rcv *IDM) MutateMAGNITUDE_LIMIT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(58, n)
}

/// Indicates if the site is taskable
func (rcv *IDM) TASKABLE() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Indicates if the site is taskable
func (rcv *IDM) MutateTASKABLE(n bool) bool {
	return rcv._tab.MutateBoolSlot(60, n)
}

func IDMStart(builder *flatbuffers.Builder) {
	builder.StartObject(29)
}
func IDMAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ID), 0)
}
func IDMAddNAME(builder *flatbuffers.Builder, NAME flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(NAME), 0)
}
func IDMAddDATA_MODE(builder *flatbuffers.Builder, DATA_MODE DataMode) {
	builder.PrependInt8Slot(2, int8(DATA_MODE), 0)
}
func IDMAddUPLINK(builder *flatbuffers.Builder, UPLINK flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(UPLINK), 0)
}
func IDMAddDOWNLINK(builder *flatbuffers.Builder, DOWNLINK flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(DOWNLINK), 0)
}
func IDMAddBEACON(builder *flatbuffers.Builder, BEACON flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(BEACON), 0)
}
func IDMAddBAND(builder *flatbuffers.Builder, BAND flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(BAND), 0)
}
func IDMStartBANDVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func IDMAddPOLARIZATION_TYPE(builder *flatbuffers.Builder, POLARIZATION_TYPE PolarizationType) {
	builder.PrependInt8Slot(7, int8(POLARIZATION_TYPE), 0)
}
func IDMAddSIMPLE_POLARIZATION(builder *flatbuffers.Builder, SIMPLE_POLARIZATION SimplePolarization) {
	builder.PrependInt8Slot(8, int8(SIMPLE_POLARIZATION), 0)
}
func IDMAddSTOKES_PARAMETERS(builder *flatbuffers.Builder, STOKES_PARAMETERS flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(STOKES_PARAMETERS), 0)
}
func IDMAddPOWER_REQUIRED(builder *flatbuffers.Builder, POWER_REQUIRED float64) {
	builder.PrependFloat64Slot(10, POWER_REQUIRED, 0.0)
}
func IDMAddPOWER_TYPE(builder *flatbuffers.Builder, POWER_TYPE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(POWER_TYPE), 0)
}
func IDMAddTRANSMIT(builder *flatbuffers.Builder, TRANSMIT bool) {
	builder.PrependBoolSlot(12, TRANSMIT, false)
}
func IDMAddRECEIVE(builder *flatbuffers.Builder, RECEIVE bool) {
	builder.PrependBoolSlot(13, RECEIVE, false)
}
func IDMAddSENSOR_TYPE(builder *flatbuffers.Builder, SENSOR_TYPE DeviceType) {
	builder.PrependInt8Slot(14, int8(SENSOR_TYPE), 0)
}
func IDMAddSOURCE(builder *flatbuffers.Builder, SOURCE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(15, flatbuffers.UOffsetT(SOURCE), 0)
}
func IDMAddLAST_OB_TIME(builder *flatbuffers.Builder, LAST_OB_TIME flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(16, flatbuffers.UOffsetT(LAST_OB_TIME), 0)
}
func IDMAddLOWER_LEFT_ELEVATION_LIMIT(builder *flatbuffers.Builder, LOWER_LEFT_ELEVATION_LIMIT float64) {
	builder.PrependFloat64Slot(17, LOWER_LEFT_ELEVATION_LIMIT, 0.0)
}
func IDMAddUPPER_LEFT_AZIMUTH_LIMIT(builder *flatbuffers.Builder, UPPER_LEFT_AZIMUTH_LIMIT float64) {
	builder.PrependFloat64Slot(18, UPPER_LEFT_AZIMUTH_LIMIT, 0.0)
}
func IDMAddLOWER_RIGHT_ELEVATION_LIMIT(builder *flatbuffers.Builder, LOWER_RIGHT_ELEVATION_LIMIT float64) {
	builder.PrependFloat64Slot(19, LOWER_RIGHT_ELEVATION_LIMIT, 0.0)
}
func IDMAddLOWER_LEFT_AZIMUTH_LIMIT(builder *flatbuffers.Builder, LOWER_LEFT_AZIMUTH_LIMIT float64) {
	builder.PrependFloat64Slot(20, LOWER_LEFT_AZIMUTH_LIMIT, 0.0)
}
func IDMAddUPPER_RIGHT_ELEVATION_LIMIT(builder *flatbuffers.Builder, UPPER_RIGHT_ELEVATION_LIMIT float64) {
	builder.PrependFloat64Slot(21, UPPER_RIGHT_ELEVATION_LIMIT, 0.0)
}
func IDMAddUPPER_RIGHT_AZIMUTH_LIMIT(builder *flatbuffers.Builder, UPPER_RIGHT_AZIMUTH_LIMIT float64) {
	builder.PrependFloat64Slot(22, UPPER_RIGHT_AZIMUTH_LIMIT, 0.0)
}
func IDMAddLOWER_RIGHT_AZIMUTH_LIMIT(builder *flatbuffers.Builder, LOWER_RIGHT_AZIMUTH_LIMIT float64) {
	builder.PrependFloat64Slot(23, LOWER_RIGHT_AZIMUTH_LIMIT, 0.0)
}
func IDMAddUPPER_LEFT_ELEVATION_LIMIT(builder *flatbuffers.Builder, UPPER_LEFT_ELEVATION_LIMIT float64) {
	builder.PrependFloat64Slot(24, UPPER_LEFT_ELEVATION_LIMIT, 0.0)
}
func IDMAddRIGHT_GEO_BELT_LIMIT(builder *flatbuffers.Builder, RIGHT_GEO_BELT_LIMIT float64) {
	builder.PrependFloat64Slot(25, RIGHT_GEO_BELT_LIMIT, 0.0)
}
func IDMAddLEFT_GEO_BELT_LIMIT(builder *flatbuffers.Builder, LEFT_GEO_BELT_LIMIT float64) {
	builder.PrependFloat64Slot(26, LEFT_GEO_BELT_LIMIT, 0.0)
}
func IDMAddMAGNITUDE_LIMIT(builder *flatbuffers.Builder, MAGNITUDE_LIMIT float64) {
	builder.PrependFloat64Slot(27, MAGNITUDE_LIMIT, 0.0)
}
func IDMAddTASKABLE(builder *flatbuffers.Builder, TASKABLE bool) {
	builder.PrependBoolSlot(28, TASKABLE, false)
}
func IDMEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
