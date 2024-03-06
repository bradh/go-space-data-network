// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package MPE

import "strconv"

type referenceFrame int8

const (
	/// Earth Mean Equator and Equinox of J2000
	referenceFrameEME2000  referenceFrame = 0
	/// Geocentric Celestial Reference Frame
	referenceFrameGCRF     referenceFrame = 1
	/// Greenwich Rotating Coordinates
	referenceFrameGRC      referenceFrame = 2
	/// International Celestial Reference Frame
	referenceFrameICRF     referenceFrame = 3
	/// International Terrestrial Reference Frame 2000
	referenceFrameITRF2000 referenceFrame = 4
	/// International Terrestrial Reference Frame 1993
	referenceFrameITRF93   referenceFrame = 5
	/// International Terrestrial Reference Frame 1997
	referenceFrameITRF97   referenceFrame = 6
	/// Mars Centered Inertial
	referenceFrameMCI      referenceFrame = 7
	/// True of Date, Rotating
	referenceFrameTDR      referenceFrame = 8
	/// True Equator Mean Equinox
	referenceFrameTEME     referenceFrame = 9
	/// True of Date
	referenceFrameTOD      referenceFrame = 10
)

var EnumNamesreferenceFrame = map[referenceFrame]string{
	referenceFrameEME2000:  "EME2000",
	referenceFrameGCRF:     "GCRF",
	referenceFrameGRC:      "GRC",
	referenceFrameICRF:     "ICRF",
	referenceFrameITRF2000: "ITRF2000",
	referenceFrameITRF93:   "ITRF93",
	referenceFrameITRF97:   "ITRF97",
	referenceFrameMCI:      "MCI",
	referenceFrameTDR:      "TDR",
	referenceFrameTEME:     "TEME",
	referenceFrameTOD:      "TOD",
}

var EnumValuesreferenceFrame = map[string]referenceFrame{
	"EME2000":  referenceFrameEME2000,
	"GCRF":     referenceFrameGCRF,
	"GRC":      referenceFrameGRC,
	"ICRF":     referenceFrameICRF,
	"ITRF2000": referenceFrameITRF2000,
	"ITRF93":   referenceFrameITRF93,
	"ITRF97":   referenceFrameITRF97,
	"MCI":      referenceFrameMCI,
	"TDR":      referenceFrameTDR,
	"TEME":     referenceFrameTEME,
	"TOD":      referenceFrameTOD,
}

func (v referenceFrame) String() string {
	if s, ok := EnumNamesreferenceFrame[v]; ok {
		return s
	}
	return "referenceFrame(" + strconv.FormatInt(int64(v), 10) + ")"
}
