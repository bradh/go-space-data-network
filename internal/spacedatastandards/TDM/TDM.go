// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package TDM

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Tracking Data Message
type TDM struct {
	_tab flatbuffers.Table
}

func GetRootAsTDM(buf []byte, offset flatbuffers.UOffsetT) *TDM {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TDM{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTDM(buf []byte, offset flatbuffers.UOffsetT) *TDM {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TDM{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TDM) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TDM) Table() flatbuffers.Table {
	return rcv._tab
}

/// Unique identifier for the observation OBSERVER -  [Specific CCSDS Document]
func (rcv *TDM) OBSERVER_ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Unique identifier for the observation OBSERVER -  [Specific CCSDS Document]
/// Cartesian X coordinate of the OBSERVER location in chosen reference frame
func (rcv *TDM) OBSERVER_X() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Cartesian X coordinate of the OBSERVER location in chosen reference frame
func (rcv *TDM) MutateOBSERVER_X(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// Cartesian Y coordinate of the OBSERVER location in chosen reference frame 
func (rcv *TDM) OBSERVER_Y() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Cartesian Y coordinate of the OBSERVER location in chosen reference frame 
func (rcv *TDM) MutateOBSERVER_Y(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// Cartesian Z coordinate of the OBSERVER location in chosen reference frame 
func (rcv *TDM) OBSERVER_Z() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Cartesian Z coordinate of the OBSERVER location in chosen reference frame 
func (rcv *TDM) MutateOBSERVER_Z(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

/// Cartesian X coordinate of the OBSERVER velocity in chosen reference frame
func (rcv *TDM) OBSERVER_VX() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Cartesian X coordinate of the OBSERVER velocity in chosen reference frame
func (rcv *TDM) MutateOBSERVER_VX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

/// Cartesian Y coordinate of the OBSERVER velocity in chosen reference frame 
func (rcv *TDM) OBSERVER_VY() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Cartesian Y coordinate of the OBSERVER velocity in chosen reference frame 
func (rcv *TDM) MutateOBSERVER_VY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(14, n)
}

/// Cartesian Z coordinate of the OBSERVER velocity in chosen reference frame 
func (rcv *TDM) OBSERVER_VZ() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Cartesian Z coordinate of the OBSERVER velocity in chosen reference frame 
func (rcv *TDM) MutateOBSERVER_VZ(n float64) bool {
	return rcv._tab.MutateFloat64Slot(16, n)
}

/// Reference frame used for OBSERVER location Cartesian coordinates (e.g., ECEF, ECI)
func (rcv *TDM) OBSERVER_POSITION_REFERENCE_FRAME() referenceFrame {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return referenceFrame(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Reference frame used for OBSERVER location Cartesian coordinates (e.g., ECEF, ECI)
func (rcv *TDM) MutateOBSERVER_POSITION_REFERENCE_FRAME(n referenceFrame) bool {
	return rcv._tab.MutateInt8Slot(18, int8(n))
}

/// Reference frame used for obs location Cartesian coordinates (e.g., ECEF, ECI)
func (rcv *TDM) OBS_REFERENCE_FRAME() referenceFrame {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return referenceFrame(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Reference frame used for obs location Cartesian coordinates (e.g., ECEF, ECI)
func (rcv *TDM) MutateOBS_REFERENCE_FRAME(n referenceFrame) bool {
	return rcv._tab.MutateInt8Slot(20, int8(n))
}

/// Epoch or observation time -  CCSDS 503.0-B-1
func (rcv *TDM) EPOCH() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Epoch or observation time -  CCSDS 503.0-B-1
/// TDM version number -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) CCSDS_TDM_VERS() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// TDM version number -  CCSDS 503.0-B-1, Page D-9
/// Comments regarding TDM -  various sections, e.g., Page D-9
func (rcv *TDM) COMMENT(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *TDM) COMMENTLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Comments regarding TDM -  various sections, e.g., Page D-9
/// Date of TDM creation -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) CREATION_DATE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Date of TDM creation -  CCSDS 503.0-B-1, Page D-9
/// Originator of the TDM -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) ORIGINATOR() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Originator of the TDM -  CCSDS 503.0-B-1, Page D-9
/// Start of metadata section -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) META_START() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Start of metadata section -  CCSDS 503.0-B-1, Page D-9
/// Time system used -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) TIME_SYSTEM() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Time system used -  CCSDS 503.0-B-1, Page D-9
/// Start time of the data -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) START_TIME() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Start time of the data -  CCSDS 503.0-B-1, Page D-9
/// Stop time of the data -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) STOP_TIME() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Stop time of the data -  CCSDS 503.0-B-1, Page D-9
/// First participant in the TDM -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) PARTICIPANT_1() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// First participant in the TDM -  CCSDS 503.0-B-1, Page D-9
/// Second participant in the TDM -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) PARTICIPANT_2() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Second participant in the TDM -  CCSDS 503.0-B-1, Page D-9
/// Third participant in the TDM (if applicable) -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) PARTICIPANT_3() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Third participant in the TDM (if applicable) -  CCSDS 503.0-B-1, Page D-9
/// Fourth participant in the TDM (if applicable) -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) PARTICIPANT_4() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Fourth participant in the TDM (if applicable) -  CCSDS 503.0-B-1, Page D-9
/// Fifth participant in the TDM (if applicable) -  CCSDS 503.0-B-1, Page D-9, max participants
func (rcv *TDM) PARTICIPANT_5() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Fifth participant in the TDM (if applicable) -  CCSDS 503.0-B-1, Page D-9, max participants
/// Mode of TDM -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) MODE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Mode of TDM -  CCSDS 503.0-B-1, Page D-9
/// First path in TDM -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) PATH_1() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

/// First path in TDM -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) MutatePATH_1(n uint16) bool {
	return rcv._tab.MutateUint16Slot(52, n)
}

/// Second path in TDM (if applicable) -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) PATH_2() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

/// Second path in TDM (if applicable) -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) MutatePATH_2(n uint16) bool {
	return rcv._tab.MutateUint16Slot(54, n)
}

/// Transmit band -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) TRANSMIT_BAND() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Transmit band -  CCSDS 503.0-B-1, Page D-9
/// Receive band -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) RECEIVE_BAND() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Receive band -  CCSDS 503.0-B-1, Page D-9
/// Integration interval -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) INTEGRATION_INTERVAL() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Integration interval -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) MutateINTEGRATION_INTERVAL(n float32) bool {
	return rcv._tab.MutateFloat32Slot(60, n)
}

/// Integration reference -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) INTEGRATION_REF() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Integration reference -  CCSDS 503.0-B-1, Page D-9
/// Receive delay for second participant -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) RECEIVE_DELAY_2() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Receive delay for second participant -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) MutateRECEIVE_DELAY_2(n float64) bool {
	return rcv._tab.MutateFloat64Slot(64, n)
}

/// Receive delay for third participant -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) RECEIVE_DELAY_3() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Receive delay for third participant -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) MutateRECEIVE_DELAY_3(n float64) bool {
	return rcv._tab.MutateFloat64Slot(66, n)
}

/// Data quality -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) DATA_QUALITY() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(68))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Data quality -  CCSDS 503.0-B-1, Page D-9
/// End of metadata section -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) META_STOP() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(70))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// End of metadata section -  CCSDS 503.0-B-1, Page D-9
/// Start of data section -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) DATA_START() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(72))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Start of data section -  CCSDS 503.0-B-1, Page D-9
/// Transmit frequency for first participant -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) TRANSMIT_FREQ_1() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(74))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Transmit frequency for first participant -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) MutateTRANSMIT_FREQ_1(n float64) bool {
	return rcv._tab.MutateFloat64Slot(74, n)
}

/// Receive frequency -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) RECEIVE_FREQ(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(76))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *TDM) RECEIVE_FREQLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(76))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Receive frequency -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) MutateRECEIVE_FREQ(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(76))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// End of data section -  CCSDS 503.0-B-1, Page D-9
func (rcv *TDM) DATA_STOP() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(78))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// End of data section -  CCSDS 503.0-B-1, Page D-9
/// Additional properties as required by the specific application of the TDM...
/// Reference for time tagging -  CCSDS 503.0-B-1, Page D-10
func (rcv *TDM) TIMETAG_REF() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(80))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Additional properties as required by the specific application of the TDM...
/// Reference for time tagging -  CCSDS 503.0-B-1, Page D-10
/// Type of angle data -  CCSDS 503.0-B-1, Page D-12
/// Can be AZEL, RADEC, XEYN, XSYE, or another value with provided ICD
func (rcv *TDM) ANGLE_TYPE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(82))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Type of angle data -  CCSDS 503.0-B-1, Page D-12
/// Can be AZEL, RADEC, XEYN, XSYE, or another value with provided ICD
/// First angle value -  CCSDS 503.0-B-1, Page D-12
func (rcv *TDM) ANGLE_1(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(84))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *TDM) ANGLE_1Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(84))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// First angle value -  CCSDS 503.0-B-1, Page D-12
func (rcv *TDM) MutateANGLE_1(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(84))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

/// Second angle value -  CCSDS 503.0-B-1, Page D-12
func (rcv *TDM) ANGLE_2(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(86))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *TDM) ANGLE_2Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(86))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Second angle value -  CCSDS 503.0-B-1, Page D-12
func (rcv *TDM) MutateANGLE_2(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(86))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

/// Uncertainty of first angle -  CCSDS 503.0-B-1
func (rcv *TDM) ANGLE_UNCERTAINTY_1() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(88))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty of first angle -  CCSDS 503.0-B-1
func (rcv *TDM) MutateANGLE_UNCERTAINTY_1(n float32) bool {
	return rcv._tab.MutateFloat32Slot(88, n)
}

/// Uncertainty of second angle -  CCSDS 503.0-B-1
func (rcv *TDM) ANGLE_UNCERTAINTY_2() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(90))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty of second angle -  CCSDS 503.0-B-1
func (rcv *TDM) MutateANGLE_UNCERTAINTY_2(n float32) bool {
	return rcv._tab.MutateFloat32Slot(90, n)
}

/// Rate of change of range -  CCSDS 503.0-B-1
func (rcv *TDM) RANGE_RATE() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(92))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Rate of change of range -  CCSDS 503.0-B-1
func (rcv *TDM) MutateRANGE_RATE(n float64) bool {
	return rcv._tab.MutateFloat64Slot(92, n)
}

/// Uncertainty in range -  CCSDS 503.0-B-1
func (rcv *TDM) RANGE_UNCERTAINTY() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(94))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Uncertainty in range -  CCSDS 503.0-B-1
func (rcv *TDM) MutateRANGE_UNCERTAINTY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(94, n)
}

/// Mode of range data -  CCSDS 503.0-B-1, Page D-10
func (rcv *TDM) RANGE_MODE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(96))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Mode of range data -  CCSDS 503.0-B-1, Page D-10
/// Modulus value for range data -  CCSDS 503.0-B-1, Page D-10
func (rcv *TDM) RANGE_MODULUS() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(98))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Modulus value for range data -  CCSDS 503.0-B-1, Page D-10
func (rcv *TDM) MutateRANGE_MODULUS(n float64) bool {
	return rcv._tab.MutateFloat64Slot(98, n)
}

/// First correction angle -  CCSDS 503.0-B-1, Page D-12
func (rcv *TDM) CORRECTION_ANGLE_1() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(100))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// First correction angle -  CCSDS 503.0-B-1, Page D-12
func (rcv *TDM) MutateCORRECTION_ANGLE_1(n float32) bool {
	return rcv._tab.MutateFloat32Slot(100, n)
}

/// Second correction angle -  CCSDS 503.0-B-1, Page D-12
func (rcv *TDM) CORRECTION_ANGLE_2() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(102))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Second correction angle -  CCSDS 503.0-B-1, Page D-12
func (rcv *TDM) MutateCORRECTION_ANGLE_2(n float32) bool {
	return rcv._tab.MutateFloat32Slot(102, n)
}

/// Indicator of corrections applied -  CCSDS 503.0-B-1, Page D-12
func (rcv *TDM) CORRECTIONS_APPLIED() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(104))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Indicator of corrections applied -  CCSDS 503.0-B-1, Page D-12
/// Dry component of tropospheric delay -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) TROPO_DRY(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(106))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *TDM) TROPO_DRYLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(106))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Dry component of tropospheric delay -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) MutateTROPO_DRY(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(106))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// Wet component of tropospheric delay -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) TROPO_WET(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(108))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *TDM) TROPO_WETLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(108))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Wet component of tropospheric delay -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) MutateTROPO_WET(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(108))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// Slant total electron content -  CCSDS 503.0-B-1, Page D-13
func (rcv *TDM) STEC(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(110))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *TDM) STECLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(110))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Slant total electron content -  CCSDS 503.0-B-1, Page D-13
func (rcv *TDM) MutateSTEC(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(110))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// Atmospheric pressure -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) PRESSURE(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(112))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *TDM) PRESSURELength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(112))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Atmospheric pressure -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) MutatePRESSURE(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(112))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// Relative humidity -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) RHUMIDITY(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(114))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *TDM) RHUMIDITYLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(114))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Relative humidity -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) MutateRHUMIDITY(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(114))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// Ambient temperature -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) TEMPERATURE(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(116))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *TDM) TEMPERATURELength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(116))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Ambient temperature -  CCSDS 503.0-B-1, Page D-14
func (rcv *TDM) MutateTEMPERATURE(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(116))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// Clock bias values -  CCSDS 503.0-B-1, Page D-15
func (rcv *TDM) CLOCK_BIAS(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(118))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *TDM) CLOCK_BIASLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(118))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Clock bias values -  CCSDS 503.0-B-1, Page D-15
func (rcv *TDM) MutateCLOCK_BIAS(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(118))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// Clock drift values -  CCSDS 503.0-B-1, Page D-15
func (rcv *TDM) CLOCK_DRIFT(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(120))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *TDM) CLOCK_DRIFTLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(120))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Clock drift values -  CCSDS 503.0-B-1, Page D-15
func (rcv *TDM) MutateCLOCK_DRIFT(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(120))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func TDMStart(builder *flatbuffers.Builder) {
	builder.StartObject(59)
}
func TDMAddOBSERVER_ID(builder *flatbuffers.Builder, OBSERVER_ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(OBSERVER_ID), 0)
}
func TDMAddOBSERVER_X(builder *flatbuffers.Builder, OBSERVER_X float64) {
	builder.PrependFloat64Slot(1, OBSERVER_X, 0.0)
}
func TDMAddOBSERVER_Y(builder *flatbuffers.Builder, OBSERVER_Y float64) {
	builder.PrependFloat64Slot(2, OBSERVER_Y, 0.0)
}
func TDMAddOBSERVER_Z(builder *flatbuffers.Builder, OBSERVER_Z float64) {
	builder.PrependFloat64Slot(3, OBSERVER_Z, 0.0)
}
func TDMAddOBSERVER_VX(builder *flatbuffers.Builder, OBSERVER_VX float64) {
	builder.PrependFloat64Slot(4, OBSERVER_VX, 0.0)
}
func TDMAddOBSERVER_VY(builder *flatbuffers.Builder, OBSERVER_VY float64) {
	builder.PrependFloat64Slot(5, OBSERVER_VY, 0.0)
}
func TDMAddOBSERVER_VZ(builder *flatbuffers.Builder, OBSERVER_VZ float64) {
	builder.PrependFloat64Slot(6, OBSERVER_VZ, 0.0)
}
func TDMAddOBSERVER_POSITION_REFERENCE_FRAME(builder *flatbuffers.Builder, OBSERVER_POSITION_REFERENCE_FRAME referenceFrame) {
	builder.PrependInt8Slot(7, int8(OBSERVER_POSITION_REFERENCE_FRAME), 0)
}
func TDMAddOBS_REFERENCE_FRAME(builder *flatbuffers.Builder, OBS_REFERENCE_FRAME referenceFrame) {
	builder.PrependInt8Slot(8, int8(OBS_REFERENCE_FRAME), 0)
}
func TDMAddEPOCH(builder *flatbuffers.Builder, EPOCH flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(EPOCH), 0)
}
func TDMAddCCSDS_TDM_VERS(builder *flatbuffers.Builder, CCSDS_TDM_VERS flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(CCSDS_TDM_VERS), 0)
}
func TDMAddCOMMENT(builder *flatbuffers.Builder, COMMENT flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(COMMENT), 0)
}
func TDMStartCOMMENTVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TDMAddCREATION_DATE(builder *flatbuffers.Builder, CREATION_DATE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(CREATION_DATE), 0)
}
func TDMAddORIGINATOR(builder *flatbuffers.Builder, ORIGINATOR flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(ORIGINATOR), 0)
}
func TDMAddMETA_START(builder *flatbuffers.Builder, META_START flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(META_START), 0)
}
func TDMAddTIME_SYSTEM(builder *flatbuffers.Builder, TIME_SYSTEM flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(15, flatbuffers.UOffsetT(TIME_SYSTEM), 0)
}
func TDMAddSTART_TIME(builder *flatbuffers.Builder, START_TIME flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(16, flatbuffers.UOffsetT(START_TIME), 0)
}
func TDMAddSTOP_TIME(builder *flatbuffers.Builder, STOP_TIME flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(17, flatbuffers.UOffsetT(STOP_TIME), 0)
}
func TDMAddPARTICIPANT_1(builder *flatbuffers.Builder, PARTICIPANT_1 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(18, flatbuffers.UOffsetT(PARTICIPANT_1), 0)
}
func TDMAddPARTICIPANT_2(builder *flatbuffers.Builder, PARTICIPANT_2 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(19, flatbuffers.UOffsetT(PARTICIPANT_2), 0)
}
func TDMAddPARTICIPANT_3(builder *flatbuffers.Builder, PARTICIPANT_3 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(20, flatbuffers.UOffsetT(PARTICIPANT_3), 0)
}
func TDMAddPARTICIPANT_4(builder *flatbuffers.Builder, PARTICIPANT_4 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(21, flatbuffers.UOffsetT(PARTICIPANT_4), 0)
}
func TDMAddPARTICIPANT_5(builder *flatbuffers.Builder, PARTICIPANT_5 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(22, flatbuffers.UOffsetT(PARTICIPANT_5), 0)
}
func TDMAddMODE(builder *flatbuffers.Builder, MODE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(23, flatbuffers.UOffsetT(MODE), 0)
}
func TDMAddPATH_1(builder *flatbuffers.Builder, PATH_1 uint16) {
	builder.PrependUint16Slot(24, PATH_1, 0)
}
func TDMAddPATH_2(builder *flatbuffers.Builder, PATH_2 uint16) {
	builder.PrependUint16Slot(25, PATH_2, 0)
}
func TDMAddTRANSMIT_BAND(builder *flatbuffers.Builder, TRANSMIT_BAND flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(26, flatbuffers.UOffsetT(TRANSMIT_BAND), 0)
}
func TDMAddRECEIVE_BAND(builder *flatbuffers.Builder, RECEIVE_BAND flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(27, flatbuffers.UOffsetT(RECEIVE_BAND), 0)
}
func TDMAddINTEGRATION_INTERVAL(builder *flatbuffers.Builder, INTEGRATION_INTERVAL float32) {
	builder.PrependFloat32Slot(28, INTEGRATION_INTERVAL, 0.0)
}
func TDMAddINTEGRATION_REF(builder *flatbuffers.Builder, INTEGRATION_REF flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(29, flatbuffers.UOffsetT(INTEGRATION_REF), 0)
}
func TDMAddRECEIVE_DELAY_2(builder *flatbuffers.Builder, RECEIVE_DELAY_2 float64) {
	builder.PrependFloat64Slot(30, RECEIVE_DELAY_2, 0.0)
}
func TDMAddRECEIVE_DELAY_3(builder *flatbuffers.Builder, RECEIVE_DELAY_3 float64) {
	builder.PrependFloat64Slot(31, RECEIVE_DELAY_3, 0.0)
}
func TDMAddDATA_QUALITY(builder *flatbuffers.Builder, DATA_QUALITY flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(32, flatbuffers.UOffsetT(DATA_QUALITY), 0)
}
func TDMAddMETA_STOP(builder *flatbuffers.Builder, META_STOP flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(33, flatbuffers.UOffsetT(META_STOP), 0)
}
func TDMAddDATA_START(builder *flatbuffers.Builder, DATA_START flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(34, flatbuffers.UOffsetT(DATA_START), 0)
}
func TDMAddTRANSMIT_FREQ_1(builder *flatbuffers.Builder, TRANSMIT_FREQ_1 float64) {
	builder.PrependFloat64Slot(35, TRANSMIT_FREQ_1, 0.0)
}
func TDMAddRECEIVE_FREQ(builder *flatbuffers.Builder, RECEIVE_FREQ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(36, flatbuffers.UOffsetT(RECEIVE_FREQ), 0)
}
func TDMStartRECEIVE_FREQVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TDMAddDATA_STOP(builder *flatbuffers.Builder, DATA_STOP flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(37, flatbuffers.UOffsetT(DATA_STOP), 0)
}
func TDMAddTIMETAG_REF(builder *flatbuffers.Builder, TIMETAG_REF flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(38, flatbuffers.UOffsetT(TIMETAG_REF), 0)
}
func TDMAddANGLE_TYPE(builder *flatbuffers.Builder, ANGLE_TYPE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(39, flatbuffers.UOffsetT(ANGLE_TYPE), 0)
}
func TDMAddANGLE_1(builder *flatbuffers.Builder, ANGLE_1 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(40, flatbuffers.UOffsetT(ANGLE_1), 0)
}
func TDMStartANGLE_1Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TDMAddANGLE_2(builder *flatbuffers.Builder, ANGLE_2 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(41, flatbuffers.UOffsetT(ANGLE_2), 0)
}
func TDMStartANGLE_2Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TDMAddANGLE_UNCERTAINTY_1(builder *flatbuffers.Builder, ANGLE_UNCERTAINTY_1 float32) {
	builder.PrependFloat32Slot(42, ANGLE_UNCERTAINTY_1, 0.0)
}
func TDMAddANGLE_UNCERTAINTY_2(builder *flatbuffers.Builder, ANGLE_UNCERTAINTY_2 float32) {
	builder.PrependFloat32Slot(43, ANGLE_UNCERTAINTY_2, 0.0)
}
func TDMAddRANGE_RATE(builder *flatbuffers.Builder, RANGE_RATE float64) {
	builder.PrependFloat64Slot(44, RANGE_RATE, 0.0)
}
func TDMAddRANGE_UNCERTAINTY(builder *flatbuffers.Builder, RANGE_UNCERTAINTY float64) {
	builder.PrependFloat64Slot(45, RANGE_UNCERTAINTY, 0.0)
}
func TDMAddRANGE_MODE(builder *flatbuffers.Builder, RANGE_MODE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(46, flatbuffers.UOffsetT(RANGE_MODE), 0)
}
func TDMAddRANGE_MODULUS(builder *flatbuffers.Builder, RANGE_MODULUS float64) {
	builder.PrependFloat64Slot(47, RANGE_MODULUS, 0.0)
}
func TDMAddCORRECTION_ANGLE_1(builder *flatbuffers.Builder, CORRECTION_ANGLE_1 float32) {
	builder.PrependFloat32Slot(48, CORRECTION_ANGLE_1, 0.0)
}
func TDMAddCORRECTION_ANGLE_2(builder *flatbuffers.Builder, CORRECTION_ANGLE_2 float32) {
	builder.PrependFloat32Slot(49, CORRECTION_ANGLE_2, 0.0)
}
func TDMAddCORRECTIONS_APPLIED(builder *flatbuffers.Builder, CORRECTIONS_APPLIED flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(50, flatbuffers.UOffsetT(CORRECTIONS_APPLIED), 0)
}
func TDMAddTROPO_DRY(builder *flatbuffers.Builder, TROPO_DRY flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(51, flatbuffers.UOffsetT(TROPO_DRY), 0)
}
func TDMStartTROPO_DRYVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TDMAddTROPO_WET(builder *flatbuffers.Builder, TROPO_WET flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(52, flatbuffers.UOffsetT(TROPO_WET), 0)
}
func TDMStartTROPO_WETVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TDMAddSTEC(builder *flatbuffers.Builder, STEC flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(53, flatbuffers.UOffsetT(STEC), 0)
}
func TDMStartSTECVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TDMAddPRESSURE(builder *flatbuffers.Builder, PRESSURE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(54, flatbuffers.UOffsetT(PRESSURE), 0)
}
func TDMStartPRESSUREVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TDMAddRHUMIDITY(builder *flatbuffers.Builder, RHUMIDITY flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(55, flatbuffers.UOffsetT(RHUMIDITY), 0)
}
func TDMStartRHUMIDITYVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TDMAddTEMPERATURE(builder *flatbuffers.Builder, TEMPERATURE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(56, flatbuffers.UOffsetT(TEMPERATURE), 0)
}
func TDMStartTEMPERATUREVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TDMAddCLOCK_BIAS(builder *flatbuffers.Builder, CLOCK_BIAS flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(57, flatbuffers.UOffsetT(CLOCK_BIAS), 0)
}
func TDMStartCLOCK_BIASVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TDMAddCLOCK_DRIFT(builder *flatbuffers.Builder, CLOCK_DRIFT flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(58, flatbuffers.UOffsetT(CLOCK_DRIFT), 0)
}
func TDMStartCLOCK_DRIFTVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TDMEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
