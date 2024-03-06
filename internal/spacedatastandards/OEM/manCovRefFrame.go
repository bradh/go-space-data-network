// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package OEM

import "strconv"

type manCovRefFrame int8

const (
	/// Another name for 'Radial, Transverse, Normal'
	manCovRefFrameRSW manCovRefFrame = 0
	/// Radial, Transverse, Normal
	manCovRefFrameRTN manCovRefFrame = 1
	/// A local orbital coordinate frame
	manCovRefFrameTNW manCovRefFrame = 2
)

var EnumNamesmanCovRefFrame = map[manCovRefFrame]string{
	manCovRefFrameRSW: "RSW",
	manCovRefFrameRTN: "RTN",
	manCovRefFrameTNW: "TNW",
}

var EnumValuesmanCovRefFrame = map[string]manCovRefFrame{
	"RSW": manCovRefFrameRSW,
	"RTN": manCovRefFrameRTN,
	"TNW": manCovRefFrameTNW,
}

func (v manCovRefFrame) String() string {
	if s, ok := EnumNamesmanCovRefFrame[v]; ok {
		return s
	}
	return "manCovRefFrame(" + strconv.FormatInt(int64(v), 10) + ")"
}
