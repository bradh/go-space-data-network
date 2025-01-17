// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package EOO

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Electro-Optical Observation
type EOO struct {
	_tab flatbuffers.Table
}

func GetRootAsEOO(buf []byte, offset flatbuffers.UOffsetT) *EOO {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EOO{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEOO(buf []byte, offset flatbuffers.UOffsetT) *EOO {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EOO{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EOO) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EOO) Table() flatbuffers.Table {
	return rcv._tab
}

/// Unique identifier for Earth Observation Observation
func (rcv *EOO) EOBSERVATION_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Unique identifier for Earth Observation Observation
/// Classification marking of the data
func (rcv *EOO) CLASSIFICATION() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Classification marking of the data
/// Observation time in UTC
func (rcv *EOO) OB_TIME() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Observation time in UTC
/// Quality of the correlation
func (rcv *EOO) CORR_QUALITY() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Quality of the correlation
func (rcv *EOO) MutateCORR_QUALITY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

/// Identifier for the satellite on orbit
func (rcv *EOO) ID_ON_ORBIT() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Identifier for the satellite on orbit
/// Identifier for the sensor
func (rcv *EOO) SENSOR_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Identifier for the sensor
/// Method of data collection
func (rcv *EOO) COLLECT_METHOD() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Method of data collection
/// NORAD catalog identifier for the satellite
func (rcv *EOO) NORAD_CAT_ID() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

/// NORAD catalog identifier for the satellite
func (rcv *EOO) MutateNORAD_CAT_ID(n int32) bool {
	return rcv._tab.MutateInt32Slot(18, n)
}

/// Identifier for the task
func (rcv *EOO) TASK_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Identifier for the task
/// Identifier for the transaction
func (rcv *EOO) TRANSACTION_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Identifier for the transaction
/// Identifier for the track
func (rcv *EOO) TRACK_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Identifier for the track
/// Position of the observation
func (rcv *EOO) OB_POSITION() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Position of the observation
/// Original object identifier
func (rcv *EOO) ORIG_OBJECT_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Original object identifier
/// Original sensor identifier
func (rcv *EOO) ORIG_SENSOR_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Original sensor identifier
/// Universal Coordinated Time flag
func (rcv *EOO) UCT() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Universal Coordinated Time flag
func (rcv *EOO) MutateUCT(n bool) bool {
	return rcv._tab.MutateBoolSlot(32, n)
}

/// Azimuth angle
func (rcv *EOO) AZIMUTH() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Azimuth angle
func (rcv *EOO) MutateAZIMUTH(n float32) bool {
	return rcv._tab.MutateFloat32Slot(34, n)
}

/// Uncertainty in azimuth angle
func (rcv *EOO) AZIMUTH_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in azimuth angle
func (rcv *EOO) MutateAZIMUTH_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(36, n)
}

/// Bias in azimuth angle
func (rcv *EOO) AZIMUTH_BIAS() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Bias in azimuth angle
func (rcv *EOO) MutateAZIMUTH_BIAS(n float32) bool {
	return rcv._tab.MutateFloat32Slot(38, n)
}

/// Rate of change in azimuth
func (rcv *EOO) AZIMUTH_RATE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Rate of change in azimuth
func (rcv *EOO) MutateAZIMUTH_RATE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(40, n)
}

/// Elevation angle
func (rcv *EOO) ELEVATION() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Elevation angle
func (rcv *EOO) MutateELEVATION(n float32) bool {
	return rcv._tab.MutateFloat32Slot(42, n)
}

/// Uncertainty in elevation angle
func (rcv *EOO) ELEVATION_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in elevation angle
func (rcv *EOO) MutateELEVATION_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(44, n)
}

/// Bias in elevation angle
func (rcv *EOO) ELEVATION_BIAS() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Bias in elevation angle
func (rcv *EOO) MutateELEVATION_BIAS(n float32) bool {
	return rcv._tab.MutateFloat32Slot(46, n)
}

/// Rate of change in elevation
func (rcv *EOO) ELEVATION_RATE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Rate of change in elevation
func (rcv *EOO) MutateELEVATION_RATE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(48, n)
}

/// Range to the target
func (rcv *EOO) RANGE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Range to the target
func (rcv *EOO) MutateRANGE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(50, n)
}

/// Uncertainty in range
func (rcv *EOO) RANGE_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in range
func (rcv *EOO) MutateRANGE_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(52, n)
}

/// Bias in range measurement
func (rcv *EOO) RANGE_BIAS() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Bias in range measurement
func (rcv *EOO) MutateRANGE_BIAS(n float32) bool {
	return rcv._tab.MutateFloat32Slot(54, n)
}

/// Rate of change in range
func (rcv *EOO) RANGE_RATE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Rate of change in range
func (rcv *EOO) MutateRANGE_RATE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(56, n)
}

/// Uncertainty in range rate
func (rcv *EOO) RANGE_RATE_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in range rate
func (rcv *EOO) MutateRANGE_RATE_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(58, n)
}

/// Right ascension
func (rcv *EOO) RA() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Right ascension
func (rcv *EOO) MutateRA(n float32) bool {
	return rcv._tab.MutateFloat32Slot(60, n)
}

/// Rate of change in right ascension
func (rcv *EOO) RA_RATE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Rate of change in right ascension
func (rcv *EOO) MutateRA_RATE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(62, n)
}

/// Uncertainty in right ascension
func (rcv *EOO) RA_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in right ascension
func (rcv *EOO) MutateRA_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(64, n)
}

/// Bias in right ascension
func (rcv *EOO) RA_BIAS() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Bias in right ascension
func (rcv *EOO) MutateRA_BIAS(n float32) bool {
	return rcv._tab.MutateFloat32Slot(66, n)
}

/// Declination angle
func (rcv *EOO) DECLINATION() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(68))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Declination angle
func (rcv *EOO) MutateDECLINATION(n float32) bool {
	return rcv._tab.MutateFloat32Slot(68, n)
}

/// Rate of change in declination
func (rcv *EOO) DECLINATION_RATE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(70))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Rate of change in declination
func (rcv *EOO) MutateDECLINATION_RATE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(70, n)
}

/// Uncertainty in declination
func (rcv *EOO) DECLINATION_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(72))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in declination
func (rcv *EOO) MutateDECLINATION_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(72, n)
}

/// Bias in declination
func (rcv *EOO) DECLINATION_BIAS() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(74))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Bias in declination
func (rcv *EOO) MutateDECLINATION_BIAS(n float32) bool {
	return rcv._tab.MutateFloat32Slot(74, n)
}

/// X-component of line-of-sight vector
func (rcv *EOO) LOSX() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(76))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// X-component of line-of-sight vector
func (rcv *EOO) MutateLOSX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(76, n)
}

/// Y-component of line-of-sight vector
func (rcv *EOO) LOSY() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(78))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Y-component of line-of-sight vector
func (rcv *EOO) MutateLOSY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(78, n)
}

/// Z-component of line-of-sight vector
func (rcv *EOO) LOSZ() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(80))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Z-component of line-of-sight vector
func (rcv *EOO) MutateLOSZ(n float32) bool {
	return rcv._tab.MutateFloat32Slot(80, n)
}

/// Uncertainty in line-of-sight vector
func (rcv *EOO) LOS_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(82))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in line-of-sight vector
func (rcv *EOO) MutateLOS_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(82, n)
}

/// X-component of line-of-sight velocity
func (rcv *EOO) LOSXVEL() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(84))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// X-component of line-of-sight velocity
func (rcv *EOO) MutateLOSXVEL(n float32) bool {
	return rcv._tab.MutateFloat32Slot(84, n)
}

/// Y-component of line-of-sight velocity
func (rcv *EOO) LOSYVEL() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(86))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Y-component of line-of-sight velocity
func (rcv *EOO) MutateLOSYVEL(n float32) bool {
	return rcv._tab.MutateFloat32Slot(86, n)
}

/// Z-component of line-of-sight velocity
func (rcv *EOO) LOSZVEL() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(88))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Z-component of line-of-sight velocity
func (rcv *EOO) MutateLOSZVEL(n float32) bool {
	return rcv._tab.MutateFloat32Slot(88, n)
}

/// Latitude of sensor
func (rcv *EOO) SENLAT() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(90))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Latitude of sensor
func (rcv *EOO) MutateSENLAT(n float32) bool {
	return rcv._tab.MutateFloat32Slot(90, n)
}

/// Longitude of sensor
func (rcv *EOO) SENLON() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(92))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Longitude of sensor
func (rcv *EOO) MutateSENLON(n float32) bool {
	return rcv._tab.MutateFloat32Slot(92, n)
}

/// Altitude of sensor
func (rcv *EOO) SENALT() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(94))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Altitude of sensor
func (rcv *EOO) MutateSENALT(n float32) bool {
	return rcv._tab.MutateFloat32Slot(94, n)
}

/// X-coordinate of sensor position
func (rcv *EOO) SENX() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(96))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// X-coordinate of sensor position
func (rcv *EOO) MutateSENX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(96, n)
}

/// Y-coordinate of sensor position
func (rcv *EOO) SENY() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(98))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Y-coordinate of sensor position
func (rcv *EOO) MutateSENY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(98, n)
}

/// Z-coordinate of sensor position
func (rcv *EOO) SENZ() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(100))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Z-coordinate of sensor position
func (rcv *EOO) MutateSENZ(n float32) bool {
	return rcv._tab.MutateFloat32Slot(100, n)
}

/// Number of fields of view
func (rcv *EOO) FOV_COUNT() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(102))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

/// Number of fields of view
func (rcv *EOO) MutateFOV_COUNT(n int32) bool {
	return rcv._tab.MutateInt32Slot(102, n)
}

/// Duration of the exposure
func (rcv *EOO) EXP_DURATION() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(104))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Duration of the exposure
func (rcv *EOO) MutateEXP_DURATION(n float32) bool {
	return rcv._tab.MutateFloat32Slot(104, n)
}

/// Zero-point displacement
func (rcv *EOO) ZEROPTD() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(106))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Zero-point displacement
func (rcv *EOO) MutateZEROPTD(n float32) bool {
	return rcv._tab.MutateFloat32Slot(106, n)
}

/// Net object signal
func (rcv *EOO) NET_OBJ_SIG() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(108))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Net object signal
func (rcv *EOO) MutateNET_OBJ_SIG(n float32) bool {
	return rcv._tab.MutateFloat32Slot(108, n)
}

/// Uncertainty in net object signal
func (rcv *EOO) NET_OBJ_SIG_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(110))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in net object signal
func (rcv *EOO) MutateNET_OBJ_SIG_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(110, n)
}

/// Magnitude of the observation
func (rcv *EOO) MAG() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(112))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Magnitude of the observation
func (rcv *EOO) MutateMAG(n float32) bool {
	return rcv._tab.MutateFloat32Slot(112, n)
}

/// Uncertainty in magnitude
func (rcv *EOO) MAG_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(114))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in magnitude
func (rcv *EOO) MutateMAG_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(114, n)
}

/// Normalized range for magnitude
func (rcv *EOO) MAG_NORM_RANGE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(116))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Normalized range for magnitude
func (rcv *EOO) MutateMAG_NORM_RANGE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(116, n)
}

/// Geocentric latitude
func (rcv *EOO) GEOLAT() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(118))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Geocentric latitude
func (rcv *EOO) MutateGEOLAT(n float32) bool {
	return rcv._tab.MutateFloat32Slot(118, n)
}

/// Geocentric longitude
func (rcv *EOO) GEOLON() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(120))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Geocentric longitude
func (rcv *EOO) MutateGEOLON(n float32) bool {
	return rcv._tab.MutateFloat32Slot(120, n)
}

/// Geocentric altitude
func (rcv *EOO) GEOALT() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(122))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Geocentric altitude
func (rcv *EOO) MutateGEOALT(n float32) bool {
	return rcv._tab.MutateFloat32Slot(122, n)
}

/// Geocentric range
func (rcv *EOO) GEORANGE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(124))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Geocentric range
func (rcv *EOO) MutateGEORANGE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(124, n)
}

/// Sky background level
func (rcv *EOO) SKY_BKGRND() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(126))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Sky background level
func (rcv *EOO) MutateSKY_BKGRND(n float32) bool {
	return rcv._tab.MutateFloat32Slot(126, n)
}

/// Primary extinction
func (rcv *EOO) PRIMARY_EXTINCTION() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(128))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Primary extinction
func (rcv *EOO) MutatePRIMARY_EXTINCTION(n float32) bool {
	return rcv._tab.MutateFloat32Slot(128, n)
}

/// Uncertainty in primary extinction
func (rcv *EOO) PRIMARY_EXTINCTION_UNC() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(130))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in primary extinction
func (rcv *EOO) MutatePRIMARY_EXTINCTION_UNC(n float32) bool {
	return rcv._tab.MutateFloat32Slot(130, n)
}

/// Solar phase angle
func (rcv *EOO) SOLAR_PHASE_ANGLE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(132))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Solar phase angle
func (rcv *EOO) MutateSOLAR_PHASE_ANGLE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(132, n)
}

/// Solar equatorial phase angle
func (rcv *EOO) SOLAR_EQ_PHASE_ANGLE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(134))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Solar equatorial phase angle
func (rcv *EOO) MutateSOLAR_EQ_PHASE_ANGLE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(134, n)
}

/// Solar declination angle
func (rcv *EOO) SOLAR_DEC_ANGLE() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(136))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Solar declination angle
func (rcv *EOO) MutateSOLAR_DEC_ANGLE(n float32) bool {
	return rcv._tab.MutateFloat32Slot(136, n)
}

/// Shutter delay
func (rcv *EOO) SHUTTER_DELAY() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(138))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Shutter delay
func (rcv *EOO) MutateSHUTTER_DELAY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(138, n)
}

/// Timing bias
func (rcv *EOO) TIMING_BIAS() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(140))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Timing bias
func (rcv *EOO) MutateTIMING_BIAS(n float32) bool {
	return rcv._tab.MutateFloat32Slot(140, n)
}

/// URI of the raw data file
func (rcv *EOO) RAW_FILE_URI() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(142))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// URI of the raw data file
/// Intensity of the observation
func (rcv *EOO) INTENSITY() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(144))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Intensity of the observation
func (rcv *EOO) MutateINTENSITY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(144, n)
}

/// Background intensity
func (rcv *EOO) BG_INTENSITY() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(146))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Background intensity
func (rcv *EOO) MutateBG_INTENSITY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(146, n)
}

/// Descriptor of the provided data
func (rcv *EOO) DESCRIPTOR() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(148))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Descriptor of the provided data
/// Source of the data
func (rcv *EOO) SOURCE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(150))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Source of the data
/// Origin of the data
func (rcv *EOO) ORIGIN() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(152))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Origin of the data
/// Mode of the data
func (rcv *EOO) DATA_MODE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(154))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Mode of the data
/// Creation time of the record
func (rcv *EOO) CREATED_AT() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(156))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Creation time of the record
/// User who created the record
func (rcv *EOO) CREATED_BY() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(158))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// User who created the record
/// Reference frame of the observation
func (rcv *EOO) REFERENCE_FRAME() referenceFrame {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(160))
	if o != 0 {
		return referenceFrame(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Reference frame of the observation
func (rcv *EOO) MutateREFERENCE_FRAME(n referenceFrame) bool {
	return rcv._tab.MutateInt8Slot(160, int8(n))
}

/// Reference frame of the sensor
func (rcv *EOO) SEN_REFERENCE_FRAME() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(162))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Reference frame of the sensor
/// Flag for umbra (total eclipse)
func (rcv *EOO) UMBRA() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(164))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Flag for umbra (total eclipse)
func (rcv *EOO) MutateUMBRA(n bool) bool {
	return rcv._tab.MutateBoolSlot(164, n)
}

/// Flag for penumbra (partial eclipse)
func (rcv *EOO) PENUMBRA() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(166))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Flag for penumbra (partial eclipse)
func (rcv *EOO) MutatePENUMBRA(n bool) bool {
	return rcv._tab.MutateBoolSlot(166, n)
}

/// Original network identifier
func (rcv *EOO) ORIG_NETWORK() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(168))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Original network identifier
/// Data link source
func (rcv *EOO) SOURCE_DL() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(170))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Data link source
/// Type of the observation
func (rcv *EOO) TYPE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(172))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Type of the observation
func EOOStart(builder *flatbuffers.Builder) {
	builder.StartObject(85)
}
func EOOAddEOBSERVATION_ID(builder *flatbuffers.Builder, EOBSERVATION_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(EOBSERVATION_ID), 0)
}
func EOOAddCLASSIFICATION(builder *flatbuffers.Builder, CLASSIFICATION flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(CLASSIFICATION), 0)
}
func EOOAddOB_TIME(builder *flatbuffers.Builder, OB_TIME flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(OB_TIME), 0)
}
func EOOAddCORR_QUALITY(builder *flatbuffers.Builder, CORR_QUALITY float32) {
	builder.PrependFloat32Slot(3, CORR_QUALITY, 0.0)
}
func EOOAddID_ON_ORBIT(builder *flatbuffers.Builder, ID_ON_ORBIT flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(ID_ON_ORBIT), 0)
}
func EOOAddSENSOR_ID(builder *flatbuffers.Builder, SENSOR_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(SENSOR_ID), 0)
}
func EOOAddCOLLECT_METHOD(builder *flatbuffers.Builder, COLLECT_METHOD flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(COLLECT_METHOD), 0)
}
func EOOAddNORAD_CAT_ID(builder *flatbuffers.Builder, NORAD_CAT_ID int32) {
	builder.PrependInt32Slot(7, NORAD_CAT_ID, 0)
}
func EOOAddTASK_ID(builder *flatbuffers.Builder, TASK_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(TASK_ID), 0)
}
func EOOAddTRANSACTION_ID(builder *flatbuffers.Builder, TRANSACTION_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(TRANSACTION_ID), 0)
}
func EOOAddTRACK_ID(builder *flatbuffers.Builder, TRACK_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(TRACK_ID), 0)
}
func EOOAddOB_POSITION(builder *flatbuffers.Builder, OB_POSITION flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(OB_POSITION), 0)
}
func EOOAddORIG_OBJECT_ID(builder *flatbuffers.Builder, ORIG_OBJECT_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(ORIG_OBJECT_ID), 0)
}
func EOOAddORIG_SENSOR_ID(builder *flatbuffers.Builder, ORIG_SENSOR_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(ORIG_SENSOR_ID), 0)
}
func EOOAddUCT(builder *flatbuffers.Builder, UCT bool) {
	builder.PrependBoolSlot(14, UCT, false)
}
func EOOAddAZIMUTH(builder *flatbuffers.Builder, AZIMUTH float32) {
	builder.PrependFloat32Slot(15, AZIMUTH, 0.0)
}
func EOOAddAZIMUTH_UNC(builder *flatbuffers.Builder, AZIMUTH_UNC float32) {
	builder.PrependFloat32Slot(16, AZIMUTH_UNC, 0.0)
}
func EOOAddAZIMUTH_BIAS(builder *flatbuffers.Builder, AZIMUTH_BIAS float32) {
	builder.PrependFloat32Slot(17, AZIMUTH_BIAS, 0.0)
}
func EOOAddAZIMUTH_RATE(builder *flatbuffers.Builder, AZIMUTH_RATE float32) {
	builder.PrependFloat32Slot(18, AZIMUTH_RATE, 0.0)
}
func EOOAddELEVATION(builder *flatbuffers.Builder, ELEVATION float32) {
	builder.PrependFloat32Slot(19, ELEVATION, 0.0)
}
func EOOAddELEVATION_UNC(builder *flatbuffers.Builder, ELEVATION_UNC float32) {
	builder.PrependFloat32Slot(20, ELEVATION_UNC, 0.0)
}
func EOOAddELEVATION_BIAS(builder *flatbuffers.Builder, ELEVATION_BIAS float32) {
	builder.PrependFloat32Slot(21, ELEVATION_BIAS, 0.0)
}
func EOOAddELEVATION_RATE(builder *flatbuffers.Builder, ELEVATION_RATE float32) {
	builder.PrependFloat32Slot(22, ELEVATION_RATE, 0.0)
}
func EOOAddRANGE(builder *flatbuffers.Builder, RANGE float32) {
	builder.PrependFloat32Slot(23, RANGE, 0.0)
}
func EOOAddRANGE_UNC(builder *flatbuffers.Builder, RANGE_UNC float32) {
	builder.PrependFloat32Slot(24, RANGE_UNC, 0.0)
}
func EOOAddRANGE_BIAS(builder *flatbuffers.Builder, RANGE_BIAS float32) {
	builder.PrependFloat32Slot(25, RANGE_BIAS, 0.0)
}
func EOOAddRANGE_RATE(builder *flatbuffers.Builder, RANGE_RATE float32) {
	builder.PrependFloat32Slot(26, RANGE_RATE, 0.0)
}
func EOOAddRANGE_RATE_UNC(builder *flatbuffers.Builder, RANGE_RATE_UNC float32) {
	builder.PrependFloat32Slot(27, RANGE_RATE_UNC, 0.0)
}
func EOOAddRA(builder *flatbuffers.Builder, RA float32) {
	builder.PrependFloat32Slot(28, RA, 0.0)
}
func EOOAddRA_RATE(builder *flatbuffers.Builder, RA_RATE float32) {
	builder.PrependFloat32Slot(29, RA_RATE, 0.0)
}
func EOOAddRA_UNC(builder *flatbuffers.Builder, RA_UNC float32) {
	builder.PrependFloat32Slot(30, RA_UNC, 0.0)
}
func EOOAddRA_BIAS(builder *flatbuffers.Builder, RA_BIAS float32) {
	builder.PrependFloat32Slot(31, RA_BIAS, 0.0)
}
func EOOAddDECLINATION(builder *flatbuffers.Builder, DECLINATION float32) {
	builder.PrependFloat32Slot(32, DECLINATION, 0.0)
}
func EOOAddDECLINATION_RATE(builder *flatbuffers.Builder, DECLINATION_RATE float32) {
	builder.PrependFloat32Slot(33, DECLINATION_RATE, 0.0)
}
func EOOAddDECLINATION_UNC(builder *flatbuffers.Builder, DECLINATION_UNC float32) {
	builder.PrependFloat32Slot(34, DECLINATION_UNC, 0.0)
}
func EOOAddDECLINATION_BIAS(builder *flatbuffers.Builder, DECLINATION_BIAS float32) {
	builder.PrependFloat32Slot(35, DECLINATION_BIAS, 0.0)
}
func EOOAddLOSX(builder *flatbuffers.Builder, LOSX float32) {
	builder.PrependFloat32Slot(36, LOSX, 0.0)
}
func EOOAddLOSY(builder *flatbuffers.Builder, LOSY float32) {
	builder.PrependFloat32Slot(37, LOSY, 0.0)
}
func EOOAddLOSZ(builder *flatbuffers.Builder, LOSZ float32) {
	builder.PrependFloat32Slot(38, LOSZ, 0.0)
}
func EOOAddLOS_UNC(builder *flatbuffers.Builder, LOS_UNC float32) {
	builder.PrependFloat32Slot(39, LOS_UNC, 0.0)
}
func EOOAddLOSXVEL(builder *flatbuffers.Builder, LOSXVEL float32) {
	builder.PrependFloat32Slot(40, LOSXVEL, 0.0)
}
func EOOAddLOSYVEL(builder *flatbuffers.Builder, LOSYVEL float32) {
	builder.PrependFloat32Slot(41, LOSYVEL, 0.0)
}
func EOOAddLOSZVEL(builder *flatbuffers.Builder, LOSZVEL float32) {
	builder.PrependFloat32Slot(42, LOSZVEL, 0.0)
}
func EOOAddSENLAT(builder *flatbuffers.Builder, SENLAT float32) {
	builder.PrependFloat32Slot(43, SENLAT, 0.0)
}
func EOOAddSENLON(builder *flatbuffers.Builder, SENLON float32) {
	builder.PrependFloat32Slot(44, SENLON, 0.0)
}
func EOOAddSENALT(builder *flatbuffers.Builder, SENALT float32) {
	builder.PrependFloat32Slot(45, SENALT, 0.0)
}
func EOOAddSENX(builder *flatbuffers.Builder, SENX float32) {
	builder.PrependFloat32Slot(46, SENX, 0.0)
}
func EOOAddSENY(builder *flatbuffers.Builder, SENY float32) {
	builder.PrependFloat32Slot(47, SENY, 0.0)
}
func EOOAddSENZ(builder *flatbuffers.Builder, SENZ float32) {
	builder.PrependFloat32Slot(48, SENZ, 0.0)
}
func EOOAddFOV_COUNT(builder *flatbuffers.Builder, FOV_COUNT int32) {
	builder.PrependInt32Slot(49, FOV_COUNT, 0)
}
func EOOAddEXP_DURATION(builder *flatbuffers.Builder, EXP_DURATION float32) {
	builder.PrependFloat32Slot(50, EXP_DURATION, 0.0)
}
func EOOAddZEROPTD(builder *flatbuffers.Builder, ZEROPTD float32) {
	builder.PrependFloat32Slot(51, ZEROPTD, 0.0)
}
func EOOAddNET_OBJ_SIG(builder *flatbuffers.Builder, NET_OBJ_SIG float32) {
	builder.PrependFloat32Slot(52, NET_OBJ_SIG, 0.0)
}
func EOOAddNET_OBJ_SIG_UNC(builder *flatbuffers.Builder, NET_OBJ_SIG_UNC float32) {
	builder.PrependFloat32Slot(53, NET_OBJ_SIG_UNC, 0.0)
}
func EOOAddMAG(builder *flatbuffers.Builder, MAG float32) {
	builder.PrependFloat32Slot(54, MAG, 0.0)
}
func EOOAddMAG_UNC(builder *flatbuffers.Builder, MAG_UNC float32) {
	builder.PrependFloat32Slot(55, MAG_UNC, 0.0)
}
func EOOAddMAG_NORM_RANGE(builder *flatbuffers.Builder, MAG_NORM_RANGE float32) {
	builder.PrependFloat32Slot(56, MAG_NORM_RANGE, 0.0)
}
func EOOAddGEOLAT(builder *flatbuffers.Builder, GEOLAT float32) {
	builder.PrependFloat32Slot(57, GEOLAT, 0.0)
}
func EOOAddGEOLON(builder *flatbuffers.Builder, GEOLON float32) {
	builder.PrependFloat32Slot(58, GEOLON, 0.0)
}
func EOOAddGEOALT(builder *flatbuffers.Builder, GEOALT float32) {
	builder.PrependFloat32Slot(59, GEOALT, 0.0)
}
func EOOAddGEORANGE(builder *flatbuffers.Builder, GEORANGE float32) {
	builder.PrependFloat32Slot(60, GEORANGE, 0.0)
}
func EOOAddSKY_BKGRND(builder *flatbuffers.Builder, SKY_BKGRND float32) {
	builder.PrependFloat32Slot(61, SKY_BKGRND, 0.0)
}
func EOOAddPRIMARY_EXTINCTION(builder *flatbuffers.Builder, PRIMARY_EXTINCTION float32) {
	builder.PrependFloat32Slot(62, PRIMARY_EXTINCTION, 0.0)
}
func EOOAddPRIMARY_EXTINCTION_UNC(builder *flatbuffers.Builder, PRIMARY_EXTINCTION_UNC float32) {
	builder.PrependFloat32Slot(63, PRIMARY_EXTINCTION_UNC, 0.0)
}
func EOOAddSOLAR_PHASE_ANGLE(builder *flatbuffers.Builder, SOLAR_PHASE_ANGLE float32) {
	builder.PrependFloat32Slot(64, SOLAR_PHASE_ANGLE, 0.0)
}
func EOOAddSOLAR_EQ_PHASE_ANGLE(builder *flatbuffers.Builder, SOLAR_EQ_PHASE_ANGLE float32) {
	builder.PrependFloat32Slot(65, SOLAR_EQ_PHASE_ANGLE, 0.0)
}
func EOOAddSOLAR_DEC_ANGLE(builder *flatbuffers.Builder, SOLAR_DEC_ANGLE float32) {
	builder.PrependFloat32Slot(66, SOLAR_DEC_ANGLE, 0.0)
}
func EOOAddSHUTTER_DELAY(builder *flatbuffers.Builder, SHUTTER_DELAY float32) {
	builder.PrependFloat32Slot(67, SHUTTER_DELAY, 0.0)
}
func EOOAddTIMING_BIAS(builder *flatbuffers.Builder, TIMING_BIAS float32) {
	builder.PrependFloat32Slot(68, TIMING_BIAS, 0.0)
}
func EOOAddRAW_FILE_URI(builder *flatbuffers.Builder, RAW_FILE_URI flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(69, flatbuffers.UOffsetT(RAW_FILE_URI), 0)
}
func EOOAddINTENSITY(builder *flatbuffers.Builder, INTENSITY float32) {
	builder.PrependFloat32Slot(70, INTENSITY, 0.0)
}
func EOOAddBG_INTENSITY(builder *flatbuffers.Builder, BG_INTENSITY float32) {
	builder.PrependFloat32Slot(71, BG_INTENSITY, 0.0)
}
func EOOAddDESCRIPTOR(builder *flatbuffers.Builder, DESCRIPTOR flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(72, flatbuffers.UOffsetT(DESCRIPTOR), 0)
}
func EOOAddSOURCE(builder *flatbuffers.Builder, SOURCE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(73, flatbuffers.UOffsetT(SOURCE), 0)
}
func EOOAddORIGIN(builder *flatbuffers.Builder, ORIGIN flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(74, flatbuffers.UOffsetT(ORIGIN), 0)
}
func EOOAddDATA_MODE(builder *flatbuffers.Builder, DATA_MODE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(75, flatbuffers.UOffsetT(DATA_MODE), 0)
}
func EOOAddCREATED_AT(builder *flatbuffers.Builder, CREATED_AT flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(76, flatbuffers.UOffsetT(CREATED_AT), 0)
}
func EOOAddCREATED_BY(builder *flatbuffers.Builder, CREATED_BY flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(77, flatbuffers.UOffsetT(CREATED_BY), 0)
}
func EOOAddREFERENCE_FRAME(builder *flatbuffers.Builder, REFERENCE_FRAME referenceFrame) {
	builder.PrependInt8Slot(78, int8(REFERENCE_FRAME), 0)
}
func EOOAddSEN_REFERENCE_FRAME(builder *flatbuffers.Builder, SEN_REFERENCE_FRAME flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(79, flatbuffers.UOffsetT(SEN_REFERENCE_FRAME), 0)
}
func EOOAddUMBRA(builder *flatbuffers.Builder, UMBRA bool) {
	builder.PrependBoolSlot(80, UMBRA, false)
}
func EOOAddPENUMBRA(builder *flatbuffers.Builder, PENUMBRA bool) {
	builder.PrependBoolSlot(81, PENUMBRA, false)
}
func EOOAddORIG_NETWORK(builder *flatbuffers.Builder, ORIG_NETWORK flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(82, flatbuffers.UOffsetT(ORIG_NETWORK), 0)
}
func EOOAddSOURCE_DL(builder *flatbuffers.Builder, SOURCE_DL flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(83, flatbuffers.UOffsetT(SOURCE_DL), 0)
}
func EOOAddTYPE(builder *flatbuffers.Builder, TYPE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(84, flatbuffers.UOffsetT(TYPE), 0)
}
func EOOEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
