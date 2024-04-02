// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package OMM

import "strconv"

type referenceFrame int8

const (
	/// Earth-Centered-Earth-Fixed (ECEF) frame: Rotates with Earth. Origin at Earth's center. X-axis towards prime meridian, Y-axis eastward, Z-axis towards North Pole. Ideal for terrestrial points.
	referenceFrameECEF     referenceFrame = 0
	/// International Celestial Reference Frame (ICRF): An inertial frame fixed relative to distant stars. Based on quasars. Used for precision astronomy and unaffected by Earth's rotation.
	referenceFrameICRF     referenceFrame = 1
	/// True Equator Mean Equinox (TEME): Used in SGP4 model for satellite tracking. Accounts for Earth's precession and nutation. Dynamic frame useful for orbit prediction.
	referenceFrameTEME     referenceFrame = 2
	/// East-North-Up (ENU): Local tangent plane system for surface points. "East" eastward, "North" northward, "Up" perpendicular to Earth's surface. Suited for stationary or slow-moving objects at low altitudes.
	referenceFrameENU      referenceFrame = 3
	/// North-East-Down (NED): Common in aviation and navigation. "North" northward, "East" eastward, "Down" towards Earth's center. Aligns with gravity, intuitive for aircraft and vehicles.
	referenceFrameNED      referenceFrame = 4
	/// North-East-Up (NEU): Similar to NED but "Up" axis is opposite to gravity. Suited for applications preferring a conventional "Up" direction.
	referenceFrameNEU      referenceFrame = 5
	/// Radial-Intrack-Cross-track (RIC): Aligned with spacecraft's UVW system. "Radial" axis towards spacecraft, "In-track" perpendicular to radial and cross-track, "Cross-track" normal to orbit plane. Used for spacecraft orientation and tracking.
	referenceFrameRIC      referenceFrame = 6
	/// Earth Mean Equator and Equinox of J2000 (J2000): An Earth-Centered Inertial (ECI) frame defined by Earth's mean equator and equinox at the start of the year 2000. Fixed relative to distant stars, used for celestial mechanics and space navigation.
	referenceFrameJ2000    referenceFrame = 7
	/// Geocentric Celestial Reference Frame
	referenceFrameGCRF     referenceFrame = 8
	/// Greenwich Rotating Coordinates
	referenceFrameGRC      referenceFrame = 9
	/// International Terrestrial Reference Frame 2000
	referenceFrameITRF2000 referenceFrame = 10
	/// International Terrestrial Reference Frame 1993
	referenceFrameITRF93   referenceFrame = 11
	/// International Terrestrial Reference Frame 1997
	referenceFrameITRF97   referenceFrame = 12
	/// True of Date, Rotating
	referenceFrameTDR      referenceFrame = 13
	/// True of Date
	referenceFrameTOD      referenceFrame = 14
	/// Radial, Transverse, Normal
	referenceFrameRTN      referenceFrame = 15
	/// Transverse, Velocity, Normal
	referenceFrameTVN      referenceFrame = 16
	/// Vehicle-Body-Local-Horizontal (VVLH): An orbit reference frame with X-axis pointing from the center of the central body to the vehicle, Z-axis oppoOBSERVER to the orbital angular momentum vector, and Y-axis completing the right-handed system.
	referenceFrameVVLH     referenceFrame = 17
	/// Vehicle-Local-Vertical-Local-Horizontal (VLVH): An orbit reference frame similar to VVLH, often used in close proximity operations or surface-oriented missions.
	referenceFrameVLVH     referenceFrame = 18
	/// Local Tangent Plane (LTP): A local, surface-fixed reference frame often used for terrestrial applications, aligned with the local horizon.
	referenceFrameLTP      referenceFrame = 19
	/// Local Vertical-Local Horizontal (LVLH): An orbit reference frame with the Z-axis pointing towards the center of the central body (oppoOBSERVER to local vertical), the X-axis in the velocity direction (local horizontal), and the Y-axis completing the right-hand system.
	referenceFrameLVLH     referenceFrame = 20
	/// Polar-North-East (PNE): A variation of local coordinate systems typically used in polar regions, with axes aligned toward the geographic North Pole, Eastward, and perpendicular to the Earth's surface.
	referenceFramePNE      referenceFrame = 21
	/// Body-Fixed Reference Frame (BRF): A reference frame fixed to the body of a spacecraft or celestial object, oriented according to the body's principal axes.
	referenceFrameBRF      referenceFrame = 22
)

var EnumNamesreferenceFrame = map[referenceFrame]string{
	referenceFrameECEF:     "ECEF",
	referenceFrameICRF:     "ICRF",
	referenceFrameTEME:     "TEME",
	referenceFrameENU:      "ENU",
	referenceFrameNED:      "NED",
	referenceFrameNEU:      "NEU",
	referenceFrameRIC:      "RIC",
	referenceFrameJ2000:    "J2000",
	referenceFrameGCRF:     "GCRF",
	referenceFrameGRC:      "GRC",
	referenceFrameITRF2000: "ITRF2000",
	referenceFrameITRF93:   "ITRF93",
	referenceFrameITRF97:   "ITRF97",
	referenceFrameTDR:      "TDR",
	referenceFrameTOD:      "TOD",
	referenceFrameRTN:      "RTN",
	referenceFrameTVN:      "TVN",
	referenceFrameVVLH:     "VVLH",
	referenceFrameVLVH:     "VLVH",
	referenceFrameLTP:      "LTP",
	referenceFrameLVLH:     "LVLH",
	referenceFramePNE:      "PNE",
	referenceFrameBRF:      "BRF",
}

var EnumValuesreferenceFrame = map[string]referenceFrame{
	"ECEF":     referenceFrameECEF,
	"ICRF":     referenceFrameICRF,
	"TEME":     referenceFrameTEME,
	"ENU":      referenceFrameENU,
	"NED":      referenceFrameNED,
	"NEU":      referenceFrameNEU,
	"RIC":      referenceFrameRIC,
	"J2000":    referenceFrameJ2000,
	"GCRF":     referenceFrameGCRF,
	"GRC":      referenceFrameGRC,
	"ITRF2000": referenceFrameITRF2000,
	"ITRF93":   referenceFrameITRF93,
	"ITRF97":   referenceFrameITRF97,
	"TDR":      referenceFrameTDR,
	"TOD":      referenceFrameTOD,
	"RTN":      referenceFrameRTN,
	"TVN":      referenceFrameTVN,
	"VVLH":     referenceFrameVVLH,
	"VLVH":     referenceFrameVLVH,
	"LTP":      referenceFrameLTP,
	"LVLH":     referenceFrameLVLH,
	"PNE":      referenceFramePNE,
	"BRF":      referenceFrameBRF,
}

func (v referenceFrame) String() string {
	if s, ok := EnumNamesreferenceFrame[v]; ok {
		return s
	}
	return "referenceFrame(" + strconv.FormatInt(int64(v), 10) + ")"
}
