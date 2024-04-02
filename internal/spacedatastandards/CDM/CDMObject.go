// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package CDM

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CDMObject struct {
	_tab flatbuffers.Table
}

func GetRootAsCDMObject(buf []byte, offset flatbuffers.UOffsetT) *CDMObject {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CDMObject{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCDMObject(buf []byte, offset flatbuffers.UOffsetT) *CDMObject {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CDMObject{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CDMObject) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CDMObject) Table() flatbuffers.Table {
	return rcv._tab
}

/// A comment
func (rcv *CDMObject) COMMENT() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// A comment
func (rcv *CDMObject) OBJECT(obj *CAT) *CAT {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(CAT)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Point of Contact
func (rcv *CDMObject) POC(obj *EPM) *EPM {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(EPM)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Point of Contact
/// Operator contact position
func (rcv *CDMObject) OPERATOR_CONTACT_POSITION() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Operator contact position
/// Operator organization
func (rcv *CDMObject) OPERATOR_ORGANIZATION() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Operator organization
/// Ephemeris name
func (rcv *CDMObject) EPHEMERIS_NAME() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Ephemeris name
/// Covariance method
func (rcv *CDMObject) COVARIANCE_METHOD() covarianceMethod {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return covarianceMethod(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Covariance method
func (rcv *CDMObject) MutateCOVARIANCE_METHOD(n covarianceMethod) bool {
	return rcv._tab.MutateInt8Slot(16, int8(n))
}

/// Reference Frame in which the object position is defined
func (rcv *CDMObject) REF_FRAME() referenceFrame {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return referenceFrame(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Reference Frame in which the object position is defined
func (rcv *CDMObject) MutateREF_FRAME(n referenceFrame) bool {
	return rcv._tab.MutateInt8Slot(18, int8(n))
}

/// Gravity model
func (rcv *CDMObject) GRAVITY_MODEL() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Gravity model
/// Atmospheric model
func (rcv *CDMObject) ATMOSPHERIC_MODEL() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Atmospheric model
/// N-body perturbations
func (rcv *CDMObject) N_BODY_PERTURBATIONS() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// N-body perturbations
/// Solar radiation pressure
func (rcv *CDMObject) SOLAR_RAD_PRESSURE() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Solar radiation pressure
func (rcv *CDMObject) MutateSOLAR_RAD_PRESSURE(n bool) bool {
	return rcv._tab.MutateBoolSlot(26, n)
}

/// Earth tides
func (rcv *CDMObject) EARTH_TIDES() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Earth tides
func (rcv *CDMObject) MutateEARTH_TIDES(n bool) bool {
	return rcv._tab.MutateBoolSlot(28, n)
}

/// Intrack thrust
func (rcv *CDMObject) INTRACK_THRUST() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Intrack thrust
func (rcv *CDMObject) MutateINTRACK_THRUST(n bool) bool {
	return rcv._tab.MutateBoolSlot(30, n)
}

/// Time of last observation start
func (rcv *CDMObject) TIME_LASTOB_START() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Time of last observation start
/// Time of last observation end
func (rcv *CDMObject) TIME_LASTOB_END() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Time of last observation end
/// Recommended observation data span
func (rcv *CDMObject) RECOMMENDED_OD_SPAN() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Recommended observation data span
func (rcv *CDMObject) MutateRECOMMENDED_OD_SPAN(n float64) bool {
	return rcv._tab.MutateFloat64Slot(36, n)
}

/// Actual observation data span
func (rcv *CDMObject) ACTUAL_OD_SPAN() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Actual observation data span
func (rcv *CDMObject) MutateACTUAL_OD_SPAN(n float64) bool {
	return rcv._tab.MutateFloat64Slot(38, n)
}

/// Number of observations available
func (rcv *CDMObject) OBS_AVAILABLE() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Number of observations available
func (rcv *CDMObject) MutateOBS_AVAILABLE(n uint32) bool {
	return rcv._tab.MutateUint32Slot(40, n)
}

/// Number of observations used
func (rcv *CDMObject) OBS_USED() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Number of observations used
func (rcv *CDMObject) MutateOBS_USED(n uint32) bool {
	return rcv._tab.MutateUint32Slot(42, n)
}

/// Number of tracks available
func (rcv *CDMObject) TRACKS_AVAILABLE() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Number of tracks available
func (rcv *CDMObject) MutateTRACKS_AVAILABLE(n uint32) bool {
	return rcv._tab.MutateUint32Slot(44, n)
}

/// Number of tracks used
func (rcv *CDMObject) TRACKS_USED() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Number of tracks used
func (rcv *CDMObject) MutateTRACKS_USED(n uint32) bool {
	return rcv._tab.MutateUint32Slot(46, n)
}

/// Residuals accepted
func (rcv *CDMObject) RESIDUALS_ACCEPTED() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Residuals accepted
func (rcv *CDMObject) MutateRESIDUALS_ACCEPTED(n float64) bool {
	return rcv._tab.MutateFloat64Slot(48, n)
}

/// Weighted root mean square
func (rcv *CDMObject) WEIGHTED_RMS() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Weighted root mean square
func (rcv *CDMObject) MutateWEIGHTED_RMS(n float64) bool {
	return rcv._tab.MutateFloat64Slot(50, n)
}

/// Area of the object
func (rcv *CDMObject) AREA_PC() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Area of the object
func (rcv *CDMObject) MutateAREA_PC(n float64) bool {
	return rcv._tab.MutateFloat64Slot(52, n)
}

/// Area of the object drag
func (rcv *CDMObject) AREA_DRG() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Area of the object drag
func (rcv *CDMObject) MutateAREA_DRG(n float64) bool {
	return rcv._tab.MutateFloat64Slot(54, n)
}

/// Area of the object solar radiation pressure
func (rcv *CDMObject) AREA_SRP() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Area of the object solar radiation pressure
func (rcv *CDMObject) MutateAREA_SRP(n float64) bool {
	return rcv._tab.MutateFloat64Slot(56, n)
}

/// Object's area-to-mass ratio
func (rcv *CDMObject) CR_AREA_OVER_MASS() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Object's area-to-mass ratio
func (rcv *CDMObject) MutateCR_AREA_OVER_MASS(n float64) bool {
	return rcv._tab.MutateFloat64Slot(58, n)
}

/// Object's thrust acceleration
func (rcv *CDMObject) THRUST_ACCELERATION() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Object's thrust acceleration
func (rcv *CDMObject) MutateTHRUST_ACCELERATION(n float64) bool {
	return rcv._tab.MutateFloat64Slot(60, n)
}

/// Object's solar flux
func (rcv *CDMObject) SEDR() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Object's solar flux
func (rcv *CDMObject) MutateSEDR(n float64) bool {
	return rcv._tab.MutateFloat64Slot(62, n)
}

/// X-coordinate of the object's position in RTN coordinates
func (rcv *CDMObject) X() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// X-coordinate of the object's position in RTN coordinates
func (rcv *CDMObject) MutateX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(64, n)
}

/// Y-coordinate of the object's position in RTN
func (rcv *CDMObject) Y() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Y-coordinate of the object's position in RTN
func (rcv *CDMObject) MutateY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(66, n)
}

/// Z-coordinate of the object's position in RTN
func (rcv *CDMObject) Z() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(68))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Z-coordinate of the object's position in RTN
func (rcv *CDMObject) MutateZ(n float64) bool {
	return rcv._tab.MutateFloat64Slot(68, n)
}

/// X-coordinate of the object's position in RTN coordinates
func (rcv *CDMObject) X_DOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(70))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// X-coordinate of the object's position in RTN coordinates
func (rcv *CDMObject) MutateX_DOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(70, n)
}

/// Y-coordinate of the object's position in RTN
func (rcv *CDMObject) Y_DOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(72))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Y-coordinate of the object's position in RTN
func (rcv *CDMObject) MutateY_DOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(72, n)
}

/// Z-coordinate of the object's position in RTN
func (rcv *CDMObject) Z_DOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(74))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Z-coordinate of the object's position in RTN
func (rcv *CDMObject) MutateZ_DOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(74, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CR_R() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(76))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCR_R(n float64) bool {
	return rcv._tab.MutateFloat64Slot(76, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CT_R() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(78))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCT_R(n float64) bool {
	return rcv._tab.MutateFloat64Slot(78, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CT_T() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(80))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCT_T(n float64) bool {
	return rcv._tab.MutateFloat64Slot(80, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CN_R() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(82))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCN_R(n float64) bool {
	return rcv._tab.MutateFloat64Slot(82, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CN_T() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(84))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCN_T(n float64) bool {
	return rcv._tab.MutateFloat64Slot(84, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CN_N() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(86))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCN_N(n float64) bool {
	return rcv._tab.MutateFloat64Slot(86, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CRDOT_R() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(88))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCRDOT_R(n float64) bool {
	return rcv._tab.MutateFloat64Slot(88, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CRDOT_T() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(90))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCRDOT_T(n float64) bool {
	return rcv._tab.MutateFloat64Slot(90, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CRDOT_N() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(92))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCRDOT_N(n float64) bool {
	return rcv._tab.MutateFloat64Slot(92, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CRDOT_RDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(94))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCRDOT_RDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(94, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTDOT_R() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(96))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTDOT_R(n float64) bool {
	return rcv._tab.MutateFloat64Slot(96, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTDOT_T() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(98))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTDOT_T(n float64) bool {
	return rcv._tab.MutateFloat64Slot(98, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTDOT_N() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(100))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTDOT_N(n float64) bool {
	return rcv._tab.MutateFloat64Slot(100, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTDOT_RDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(102))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTDOT_RDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(102, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTDOT_TDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(104))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTDOT_TDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(104, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CNDOT_R() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(106))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCNDOT_R(n float64) bool {
	return rcv._tab.MutateFloat64Slot(106, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CNDOT_T() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(108))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCNDOT_T(n float64) bool {
	return rcv._tab.MutateFloat64Slot(108, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CNDOT_N() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(110))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCNDOT_N(n float64) bool {
	return rcv._tab.MutateFloat64Slot(110, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CNDOT_RDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(112))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCNDOT_RDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(112, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CNDOT_TDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(114))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCNDOT_TDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(114, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CNDOT_NDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(116))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCNDOT_NDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(116, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CDRG_R() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(118))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCDRG_R(n float64) bool {
	return rcv._tab.MutateFloat64Slot(118, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CDRG_T() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(120))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCDRG_T(n float64) bool {
	return rcv._tab.MutateFloat64Slot(120, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CDRG_N() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(122))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCDRG_N(n float64) bool {
	return rcv._tab.MutateFloat64Slot(122, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CDRG_RDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(124))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCDRG_RDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(124, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CDRG_TDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(126))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCDRG_TDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(126, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CDRG_NDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(128))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCDRG_NDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(128, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CDRG_DRG() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(130))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCDRG_DRG(n float64) bool {
	return rcv._tab.MutateFloat64Slot(130, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CSRP_R() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(132))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCSRP_R(n float64) bool {
	return rcv._tab.MutateFloat64Slot(132, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CSRP_T() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(134))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCSRP_T(n float64) bool {
	return rcv._tab.MutateFloat64Slot(134, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CSRP_N() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(136))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCSRP_N(n float64) bool {
	return rcv._tab.MutateFloat64Slot(136, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CSRP_RDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(138))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCSRP_RDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(138, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CSRP_TDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(140))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCSRP_TDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(140, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CSRP_NDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(142))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCSRP_NDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(142, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CSRP_DRG() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(144))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCSRP_DRG(n float64) bool {
	return rcv._tab.MutateFloat64Slot(144, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CSRP_SRP() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(146))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCSRP_SRP(n float64) bool {
	return rcv._tab.MutateFloat64Slot(146, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTHR_R() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(148))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTHR_R(n float64) bool {
	return rcv._tab.MutateFloat64Slot(148, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTHR_T() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(150))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTHR_T(n float64) bool {
	return rcv._tab.MutateFloat64Slot(150, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTHR_N() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(152))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTHR_N(n float64) bool {
	return rcv._tab.MutateFloat64Slot(152, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTHR_RDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(154))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTHR_RDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(154, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTHR_TDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(156))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTHR_TDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(156, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTHR_NDOT() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(158))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTHR_NDOT(n float64) bool {
	return rcv._tab.MutateFloat64Slot(158, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTHR_DRG() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(160))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTHR_DRG(n float64) bool {
	return rcv._tab.MutateFloat64Slot(160, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTHR_SRP() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(162))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTHR_SRP(n float64) bool {
	return rcv._tab.MutateFloat64Slot(162, n)
}

/// Covariance Matrix component
func (rcv *CDMObject) CTHR_THR() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(164))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Covariance Matrix component
func (rcv *CDMObject) MutateCTHR_THR(n float64) bool {
	return rcv._tab.MutateFloat64Slot(164, n)
}

func CDMObjectStart(builder *flatbuffers.Builder) {
	builder.StartObject(81)
}
func CDMObjectAddCOMMENT(builder *flatbuffers.Builder, COMMENT flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(COMMENT), 0)
}
func CDMObjectAddOBJECT(builder *flatbuffers.Builder, OBJECT flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(OBJECT), 0)
}
func CDMObjectAddPOC(builder *flatbuffers.Builder, POC flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(POC), 0)
}
func CDMObjectAddOPERATOR_CONTACT_POSITION(builder *flatbuffers.Builder, OPERATOR_CONTACT_POSITION flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(OPERATOR_CONTACT_POSITION), 0)
}
func CDMObjectAddOPERATOR_ORGANIZATION(builder *flatbuffers.Builder, OPERATOR_ORGANIZATION flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(OPERATOR_ORGANIZATION), 0)
}
func CDMObjectAddEPHEMERIS_NAME(builder *flatbuffers.Builder, EPHEMERIS_NAME flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(EPHEMERIS_NAME), 0)
}
func CDMObjectAddCOVARIANCE_METHOD(builder *flatbuffers.Builder, COVARIANCE_METHOD covarianceMethod) {
	builder.PrependInt8Slot(6, int8(COVARIANCE_METHOD), 0)
}
func CDMObjectAddREF_FRAME(builder *flatbuffers.Builder, REF_FRAME referenceFrame) {
	builder.PrependInt8Slot(7, int8(REF_FRAME), 0)
}
func CDMObjectAddGRAVITY_MODEL(builder *flatbuffers.Builder, GRAVITY_MODEL flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(GRAVITY_MODEL), 0)
}
func CDMObjectAddATMOSPHERIC_MODEL(builder *flatbuffers.Builder, ATMOSPHERIC_MODEL flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(ATMOSPHERIC_MODEL), 0)
}
func CDMObjectAddN_BODY_PERTURBATIONS(builder *flatbuffers.Builder, N_BODY_PERTURBATIONS flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(N_BODY_PERTURBATIONS), 0)
}
func CDMObjectAddSOLAR_RAD_PRESSURE(builder *flatbuffers.Builder, SOLAR_RAD_PRESSURE bool) {
	builder.PrependBoolSlot(11, SOLAR_RAD_PRESSURE, false)
}
func CDMObjectAddEARTH_TIDES(builder *flatbuffers.Builder, EARTH_TIDES bool) {
	builder.PrependBoolSlot(12, EARTH_TIDES, false)
}
func CDMObjectAddINTRACK_THRUST(builder *flatbuffers.Builder, INTRACK_THRUST bool) {
	builder.PrependBoolSlot(13, INTRACK_THRUST, false)
}
func CDMObjectAddTIME_LASTOB_START(builder *flatbuffers.Builder, TIME_LASTOB_START flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(TIME_LASTOB_START), 0)
}
func CDMObjectAddTIME_LASTOB_END(builder *flatbuffers.Builder, TIME_LASTOB_END flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(15, flatbuffers.UOffsetT(TIME_LASTOB_END), 0)
}
func CDMObjectAddRECOMMENDED_OD_SPAN(builder *flatbuffers.Builder, RECOMMENDED_OD_SPAN float64) {
	builder.PrependFloat64Slot(16, RECOMMENDED_OD_SPAN, 0.0)
}
func CDMObjectAddACTUAL_OD_SPAN(builder *flatbuffers.Builder, ACTUAL_OD_SPAN float64) {
	builder.PrependFloat64Slot(17, ACTUAL_OD_SPAN, 0.0)
}
func CDMObjectAddOBS_AVAILABLE(builder *flatbuffers.Builder, OBS_AVAILABLE uint32) {
	builder.PrependUint32Slot(18, OBS_AVAILABLE, 0)
}
func CDMObjectAddOBS_USED(builder *flatbuffers.Builder, OBS_USED uint32) {
	builder.PrependUint32Slot(19, OBS_USED, 0)
}
func CDMObjectAddTRACKS_AVAILABLE(builder *flatbuffers.Builder, TRACKS_AVAILABLE uint32) {
	builder.PrependUint32Slot(20, TRACKS_AVAILABLE, 0)
}
func CDMObjectAddTRACKS_USED(builder *flatbuffers.Builder, TRACKS_USED uint32) {
	builder.PrependUint32Slot(21, TRACKS_USED, 0)
}
func CDMObjectAddRESIDUALS_ACCEPTED(builder *flatbuffers.Builder, RESIDUALS_ACCEPTED float64) {
	builder.PrependFloat64Slot(22, RESIDUALS_ACCEPTED, 0.0)
}
func CDMObjectAddWEIGHTED_RMS(builder *flatbuffers.Builder, WEIGHTED_RMS float64) {
	builder.PrependFloat64Slot(23, WEIGHTED_RMS, 0.0)
}
func CDMObjectAddAREA_PC(builder *flatbuffers.Builder, AREA_PC float64) {
	builder.PrependFloat64Slot(24, AREA_PC, 0.0)
}
func CDMObjectAddAREA_DRG(builder *flatbuffers.Builder, AREA_DRG float64) {
	builder.PrependFloat64Slot(25, AREA_DRG, 0.0)
}
func CDMObjectAddAREA_SRP(builder *flatbuffers.Builder, AREA_SRP float64) {
	builder.PrependFloat64Slot(26, AREA_SRP, 0.0)
}
func CDMObjectAddCR_AREA_OVER_MASS(builder *flatbuffers.Builder, CR_AREA_OVER_MASS float64) {
	builder.PrependFloat64Slot(27, CR_AREA_OVER_MASS, 0.0)
}
func CDMObjectAddTHRUST_ACCELERATION(builder *flatbuffers.Builder, THRUST_ACCELERATION float64) {
	builder.PrependFloat64Slot(28, THRUST_ACCELERATION, 0.0)
}
func CDMObjectAddSEDR(builder *flatbuffers.Builder, SEDR float64) {
	builder.PrependFloat64Slot(29, SEDR, 0.0)
}
func CDMObjectAddX(builder *flatbuffers.Builder, X float64) {
	builder.PrependFloat64Slot(30, X, 0.0)
}
func CDMObjectAddY(builder *flatbuffers.Builder, Y float64) {
	builder.PrependFloat64Slot(31, Y, 0.0)
}
func CDMObjectAddZ(builder *flatbuffers.Builder, Z float64) {
	builder.PrependFloat64Slot(32, Z, 0.0)
}
func CDMObjectAddX_DOT(builder *flatbuffers.Builder, X_DOT float64) {
	builder.PrependFloat64Slot(33, X_DOT, 0.0)
}
func CDMObjectAddY_DOT(builder *flatbuffers.Builder, Y_DOT float64) {
	builder.PrependFloat64Slot(34, Y_DOT, 0.0)
}
func CDMObjectAddZ_DOT(builder *flatbuffers.Builder, Z_DOT float64) {
	builder.PrependFloat64Slot(35, Z_DOT, 0.0)
}
func CDMObjectAddCR_R(builder *flatbuffers.Builder, CR_R float64) {
	builder.PrependFloat64Slot(36, CR_R, 0.0)
}
func CDMObjectAddCT_R(builder *flatbuffers.Builder, CT_R float64) {
	builder.PrependFloat64Slot(37, CT_R, 0.0)
}
func CDMObjectAddCT_T(builder *flatbuffers.Builder, CT_T float64) {
	builder.PrependFloat64Slot(38, CT_T, 0.0)
}
func CDMObjectAddCN_R(builder *flatbuffers.Builder, CN_R float64) {
	builder.PrependFloat64Slot(39, CN_R, 0.0)
}
func CDMObjectAddCN_T(builder *flatbuffers.Builder, CN_T float64) {
	builder.PrependFloat64Slot(40, CN_T, 0.0)
}
func CDMObjectAddCN_N(builder *flatbuffers.Builder, CN_N float64) {
	builder.PrependFloat64Slot(41, CN_N, 0.0)
}
func CDMObjectAddCRDOT_R(builder *flatbuffers.Builder, CRDOT_R float64) {
	builder.PrependFloat64Slot(42, CRDOT_R, 0.0)
}
func CDMObjectAddCRDOT_T(builder *flatbuffers.Builder, CRDOT_T float64) {
	builder.PrependFloat64Slot(43, CRDOT_T, 0.0)
}
func CDMObjectAddCRDOT_N(builder *flatbuffers.Builder, CRDOT_N float64) {
	builder.PrependFloat64Slot(44, CRDOT_N, 0.0)
}
func CDMObjectAddCRDOT_RDOT(builder *flatbuffers.Builder, CRDOT_RDOT float64) {
	builder.PrependFloat64Slot(45, CRDOT_RDOT, 0.0)
}
func CDMObjectAddCTDOT_R(builder *flatbuffers.Builder, CTDOT_R float64) {
	builder.PrependFloat64Slot(46, CTDOT_R, 0.0)
}
func CDMObjectAddCTDOT_T(builder *flatbuffers.Builder, CTDOT_T float64) {
	builder.PrependFloat64Slot(47, CTDOT_T, 0.0)
}
func CDMObjectAddCTDOT_N(builder *flatbuffers.Builder, CTDOT_N float64) {
	builder.PrependFloat64Slot(48, CTDOT_N, 0.0)
}
func CDMObjectAddCTDOT_RDOT(builder *flatbuffers.Builder, CTDOT_RDOT float64) {
	builder.PrependFloat64Slot(49, CTDOT_RDOT, 0.0)
}
func CDMObjectAddCTDOT_TDOT(builder *flatbuffers.Builder, CTDOT_TDOT float64) {
	builder.PrependFloat64Slot(50, CTDOT_TDOT, 0.0)
}
func CDMObjectAddCNDOT_R(builder *flatbuffers.Builder, CNDOT_R float64) {
	builder.PrependFloat64Slot(51, CNDOT_R, 0.0)
}
func CDMObjectAddCNDOT_T(builder *flatbuffers.Builder, CNDOT_T float64) {
	builder.PrependFloat64Slot(52, CNDOT_T, 0.0)
}
func CDMObjectAddCNDOT_N(builder *flatbuffers.Builder, CNDOT_N float64) {
	builder.PrependFloat64Slot(53, CNDOT_N, 0.0)
}
func CDMObjectAddCNDOT_RDOT(builder *flatbuffers.Builder, CNDOT_RDOT float64) {
	builder.PrependFloat64Slot(54, CNDOT_RDOT, 0.0)
}
func CDMObjectAddCNDOT_TDOT(builder *flatbuffers.Builder, CNDOT_TDOT float64) {
	builder.PrependFloat64Slot(55, CNDOT_TDOT, 0.0)
}
func CDMObjectAddCNDOT_NDOT(builder *flatbuffers.Builder, CNDOT_NDOT float64) {
	builder.PrependFloat64Slot(56, CNDOT_NDOT, 0.0)
}
func CDMObjectAddCDRG_R(builder *flatbuffers.Builder, CDRG_R float64) {
	builder.PrependFloat64Slot(57, CDRG_R, 0.0)
}
func CDMObjectAddCDRG_T(builder *flatbuffers.Builder, CDRG_T float64) {
	builder.PrependFloat64Slot(58, CDRG_T, 0.0)
}
func CDMObjectAddCDRG_N(builder *flatbuffers.Builder, CDRG_N float64) {
	builder.PrependFloat64Slot(59, CDRG_N, 0.0)
}
func CDMObjectAddCDRG_RDOT(builder *flatbuffers.Builder, CDRG_RDOT float64) {
	builder.PrependFloat64Slot(60, CDRG_RDOT, 0.0)
}
func CDMObjectAddCDRG_TDOT(builder *flatbuffers.Builder, CDRG_TDOT float64) {
	builder.PrependFloat64Slot(61, CDRG_TDOT, 0.0)
}
func CDMObjectAddCDRG_NDOT(builder *flatbuffers.Builder, CDRG_NDOT float64) {
	builder.PrependFloat64Slot(62, CDRG_NDOT, 0.0)
}
func CDMObjectAddCDRG_DRG(builder *flatbuffers.Builder, CDRG_DRG float64) {
	builder.PrependFloat64Slot(63, CDRG_DRG, 0.0)
}
func CDMObjectAddCSRP_R(builder *flatbuffers.Builder, CSRP_R float64) {
	builder.PrependFloat64Slot(64, CSRP_R, 0.0)
}
func CDMObjectAddCSRP_T(builder *flatbuffers.Builder, CSRP_T float64) {
	builder.PrependFloat64Slot(65, CSRP_T, 0.0)
}
func CDMObjectAddCSRP_N(builder *flatbuffers.Builder, CSRP_N float64) {
	builder.PrependFloat64Slot(66, CSRP_N, 0.0)
}
func CDMObjectAddCSRP_RDOT(builder *flatbuffers.Builder, CSRP_RDOT float64) {
	builder.PrependFloat64Slot(67, CSRP_RDOT, 0.0)
}
func CDMObjectAddCSRP_TDOT(builder *flatbuffers.Builder, CSRP_TDOT float64) {
	builder.PrependFloat64Slot(68, CSRP_TDOT, 0.0)
}
func CDMObjectAddCSRP_NDOT(builder *flatbuffers.Builder, CSRP_NDOT float64) {
	builder.PrependFloat64Slot(69, CSRP_NDOT, 0.0)
}
func CDMObjectAddCSRP_DRG(builder *flatbuffers.Builder, CSRP_DRG float64) {
	builder.PrependFloat64Slot(70, CSRP_DRG, 0.0)
}
func CDMObjectAddCSRP_SRP(builder *flatbuffers.Builder, CSRP_SRP float64) {
	builder.PrependFloat64Slot(71, CSRP_SRP, 0.0)
}
func CDMObjectAddCTHR_R(builder *flatbuffers.Builder, CTHR_R float64) {
	builder.PrependFloat64Slot(72, CTHR_R, 0.0)
}
func CDMObjectAddCTHR_T(builder *flatbuffers.Builder, CTHR_T float64) {
	builder.PrependFloat64Slot(73, CTHR_T, 0.0)
}
func CDMObjectAddCTHR_N(builder *flatbuffers.Builder, CTHR_N float64) {
	builder.PrependFloat64Slot(74, CTHR_N, 0.0)
}
func CDMObjectAddCTHR_RDOT(builder *flatbuffers.Builder, CTHR_RDOT float64) {
	builder.PrependFloat64Slot(75, CTHR_RDOT, 0.0)
}
func CDMObjectAddCTHR_TDOT(builder *flatbuffers.Builder, CTHR_TDOT float64) {
	builder.PrependFloat64Slot(76, CTHR_TDOT, 0.0)
}
func CDMObjectAddCTHR_NDOT(builder *flatbuffers.Builder, CTHR_NDOT float64) {
	builder.PrependFloat64Slot(77, CTHR_NDOT, 0.0)
}
func CDMObjectAddCTHR_DRG(builder *flatbuffers.Builder, CTHR_DRG float64) {
	builder.PrependFloat64Slot(78, CTHR_DRG, 0.0)
}
func CDMObjectAddCTHR_SRP(builder *flatbuffers.Builder, CTHR_SRP float64) {
	builder.PrependFloat64Slot(79, CTHR_SRP, 0.0)
}
func CDMObjectAddCTHR_THR(builder *flatbuffers.Builder, CTHR_THR float64) {
	builder.PrependFloat64Slot(80, CTHR_THR, 0.0)
}
func CDMObjectEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
