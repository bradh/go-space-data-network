// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package LDM

import "strconv"

/// Simple polarization types
type SimplePolarization int8

const (
	SimplePolarizationvertical          SimplePolarization = 0
	SimplePolarizationhorizontal        SimplePolarization = 1
	SimplePolarizationleftHandCircular  SimplePolarization = 2
	SimplePolarizationrightHandCircular SimplePolarization = 3
)

var EnumNamesSimplePolarization = map[SimplePolarization]string{
	SimplePolarizationvertical:          "vertical",
	SimplePolarizationhorizontal:        "horizontal",
	SimplePolarizationleftHandCircular:  "leftHandCircular",
	SimplePolarizationrightHandCircular: "rightHandCircular",
}

var EnumValuesSimplePolarization = map[string]SimplePolarization{
	"vertical":          SimplePolarizationvertical,
	"horizontal":        SimplePolarizationhorizontal,
	"leftHandCircular":  SimplePolarizationleftHandCircular,
	"rightHandCircular": SimplePolarizationrightHandCircular,
}

func (v SimplePolarization) String() string {
	if s, ok := EnumNamesSimplePolarization[v]; ok {
		return s
	}
	return "SimplePolarization(" + strconv.FormatInt(int64(v), 10) + ")"
}
