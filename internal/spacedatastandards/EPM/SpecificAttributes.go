// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package EPM

import "strconv"

/// Union for specific attributes, distinguishing between Person and Organization
type SpecificAttributes byte

const (
	SpecificAttributesNONE                   SpecificAttributes = 0
	SpecificAttributesPersonAttributes       SpecificAttributes = 1
	SpecificAttributesOrganizationAttributes SpecificAttributes = 2
)

var EnumNamesSpecificAttributes = map[SpecificAttributes]string{
	SpecificAttributesNONE:                   "NONE",
	SpecificAttributesPersonAttributes:       "PersonAttributes",
	SpecificAttributesOrganizationAttributes: "OrganizationAttributes",
}

var EnumValuesSpecificAttributes = map[string]SpecificAttributes{
	"NONE":                   SpecificAttributesNONE,
	"PersonAttributes":       SpecificAttributesPersonAttributes,
	"OrganizationAttributes": SpecificAttributesOrganizationAttributes,
}

func (v SpecificAttributes) String() string {
	if s, ok := EnumNamesSpecificAttributes[v]; ok {
		return s
	}
	return "SpecificAttributes(" + strconv.FormatInt(int64(v), 10) + ")"
}
