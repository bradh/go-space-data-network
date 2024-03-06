// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package CAT

import "strconv"

type orbitType int8

const (
	///0
	orbitTypeORBIT     orbitType = 0
	///1
	orbitTypeLANDING   orbitType = 1
	///2
	orbitTypeIMPACT    orbitType = 2
	///3
	orbitTypeDOCKED    orbitType = 3
	///4
	orbitTypeROUNDTRIP orbitType = 4
)

var EnumNamesorbitType = map[orbitType]string{
	orbitTypeORBIT:     "ORBIT",
	orbitTypeLANDING:   "LANDING",
	orbitTypeIMPACT:    "IMPACT",
	orbitTypeDOCKED:    "DOCKED",
	orbitTypeROUNDTRIP: "ROUNDTRIP",
}

var EnumValuesorbitType = map[string]orbitType{
	"ORBIT":     orbitTypeORBIT,
	"LANDING":   orbitTypeLANDING,
	"IMPACT":    orbitTypeIMPACT,
	"DOCKED":    orbitTypeDOCKED,
	"ROUNDTRIP": orbitTypeROUNDTRIP,
}

func (v orbitType) String() string {
	if s, ok := EnumNamesorbitType[v]; ok {
		return s
	}
	return "orbitType(" + strconv.FormatInt(int64(v), 10) + ")"
}
