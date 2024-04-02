// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package CDM

import "strconv"

type objectType int8

const (
	///0
	objectTypePAYLOAD     objectType = 0
	///1
	objectTypeROCKET_BODY objectType = 1
	///2
	objectTypeDEBRIS      objectType = 2
	///3
	objectTypeUNKNOWN     objectType = 3
)

var EnumNamesobjectType = map[objectType]string{
	objectTypePAYLOAD:     "PAYLOAD",
	objectTypeROCKET_BODY: "ROCKET_BODY",
	objectTypeDEBRIS:      "DEBRIS",
	objectTypeUNKNOWN:     "UNKNOWN",
}

var EnumValuesobjectType = map[string]objectType{
	"PAYLOAD":     objectTypePAYLOAD,
	"ROCKET_BODY": objectTypeROCKET_BODY,
	"DEBRIS":      objectTypeDEBRIS,
	"UNKNOWN":     objectTypeUNKNOWN,
}

func (v objectType) String() string {
	if s, ok := EnumNamesobjectType[v]; ok {
		return s
	}
	return "objectType(" + strconv.FormatInt(int64(v), 10) + ")"
}
